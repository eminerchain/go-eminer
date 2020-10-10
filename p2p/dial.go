// Copyright 2018 The go-eminer Authors
// This file is part of the go-eminer library.
//
// The the go-eminer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The the go-eminer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-eminer library. If not, see <http://www.gnu.org/licenses/>.

package p2p

import (
	"container/heap"
	"crypto/rand"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/eminer-pro/go-eminer/log"
	"github.com/eminer-pro/go-eminer/p2p/discover"
	"github.com/eminer-pro/go-eminer/p2p/netutil"
)

const (
	// This is the amount of time spent waiting in between
	// redialing a certain node.
	dialHistoryExpiration = 30 * time.Second

	// Discovery lookups are throttled and can only run
	// once every few seconds.
	lookupInterval = 4 * time.Second

	// If no peers are found for this amount of time, the initial bootnodes are
	// attempted to be connected.
	fallbackInterval = 20 * time.Second

	// Endpoint resolution is throttled with bounded backoff.
	initialResolveDelay = 60 * time.Second
	maxResolveDelay     = time.Hour
)

// NodeDialer is used to connect to nodes in the network, typically by using
// an underlying net.Dialer but also using net.Pipe in tests
type NodeDialer interface {
	Dial(*discover.Node) (net.Conn, error)
}

// TCPDialer implements the NodeDialer interface by using a net.Dialer to
// create TCP connections to nodes in the network
type TCPDialer struct {
	*net.Dialer
}

// Dial creates a TCP connection to the node
func (t TCPDialer) Dial(dest *discover.Node) (net.Conn, error) {
	addr := &net.TCPAddr{IP: dest.IP, Port: int(dest.TCP)}
	return t.Dialer.Dial("tcp", addr.String())
}

// dialstate schedules dials and discovery lookups.
// it get's a chance to compute new tasks on every iteration
// of the main loop in Server.run.
type dialstate struct {
	maxDynDials int

	ntab discoverTable

	netrestrict *netutil.Netlist

	commonLookupRunning bool

	topLookupRunning bool

	commonDialing map[discover.NodeID]connFlag

	commonLookupBuf []*discover.Node // current discovery lookup results

	topLookupBuf []*discover.Node // current discovery lookup results

	commonRandomNodes []*discover.Node // filled from Table

	consRandomNodes []*discover.Node //

	static map[discover.NodeID]*dialTask

	commonHist *dialHistory
	topHist    *dialHistory

	start time.Time // time when the dialer was first used

	bootnodes []*discover.Node // default dials when there are no peers
}

type discoverTable interface {
	Self() *discover.Node
	Close()
	Resolve(target discover.NodeID) *discover.Node
	Delete(target discover.NodeID)
	Lookup(target discover.NodeID, netType byte) []*discover.Node
	ReadRandomNodes(nodes []*discover.Node, netType byte) int
	OpenTopNet()
}

// the dial history remembers recent dials.
type dialHistory []pastDial

// pastDial is an entry in the dial history.
type pastDial struct {
	id  discover.NodeID
	exp time.Time
}

type task interface {
	Do(*Server)
	GetNetType() byte
}

// A dialTask is generated for each node that is dialed. Its
// fields cannot be accessed while the task is running.
type dialTask struct {
	flags        connFlag
	ds           *dialstate
	netType      byte
	dest         *discover.Node
	lastResolved time.Time
	resolveDelay time.Duration
}

// discoverTask runs discovery table operations.
// Only one discoverTask is active at any time.
// discoverTask.Do performs a random lookup.
type discoverTask struct {
	results []*discover.Node
	netType byte
}

// A waitExpireTask is generated if there are no other tasks
// to keep the loop in Server.run ticking.
type waitExpireTask struct {
	time.Duration
	netType byte
}

func newDialState(static []*discover.Node, bootnodes []*discover.Node, ntab discoverTable, maxdyn int, netrestrict *netutil.Netlist) *dialstate {
	s := &dialstate{
		maxDynDials: maxdyn,
		ntab:        ntab,
		netrestrict: netrestrict,

		static: make(map[discover.NodeID]*dialTask),

		commonDialing:     make(map[discover.NodeID]connFlag),
		bootnodes:         make([]*discover.Node, len(bootnodes)),
		commonRandomNodes: make([]*discover.Node, maxdyn/2),
		consRandomNodes:   make([]*discover.Node, maxdyn/2),
		commonHist:        new(dialHistory),
		topHist:           new(dialHistory),
	}
	copy(s.bootnodes, bootnodes)
	for _, n := range static {
		s.addStatic(n)
	}
	return s
}

func (s *dialstate) addStatic(n *discover.Node) {
	// This overwites the task instead of updating an existing
	// entry, giving users the opportunity to force a resolve operation.
	s.static[n.ID] = &dialTask{flags: staticDialedConn, dest: n, ds: s}
}

func (s *dialstate) removeStatic(n *discover.Node) {
	// This removes a task so future attempts to connect will not be made.
	delete(s.static, n.ID)
}

func (s *dialstate) newTasks(nRunning int, peers map[discover.NodeID]*Peer, now time.Time, openTopNet bool) []task {
	if s.start.IsZero() {
		s.start = now
	}

	var newtasks []task
	addDial := func(flag connFlag, n *discover.Node) bool {
		if err := s.checkDial(n, peers); err != nil {
			log.Trace("Skipping dial candidate", "id", n.ID, "addr", &net.TCPAddr{IP: n.IP, Port: int(n.TCP)}, "err", err)
			return false
		}

		s.commonDialing[n.ID] = flag

		newtasks = append(newtasks, &dialTask{flags: flag, dest: n, ds: s})
		return true
	}

	// Compute number of dynamic dials necessary at this point.
	needDynDials := s.maxDynDials
	for _, p := range peers {
		if p.rw.is(dynDialedConn) {
			needDynDials--
		}
	}

	for _, flag := range s.commonDialing {
		if flag&dynDialedConn != 0 {
			needDynDials--
		}
	}
	s.commonHist.expire(now)

	// Create dials for static nodes if they are not connected.
	for id, t := range s.static {
		err := s.checkDial(t.dest, peers)
		switch err {
		case errNotWhitelisted, errSelf:
			log.Warn("Removing static dial candidate", "id", t.dest.ID, "addr", &net.TCPAddr{IP: t.dest.IP, Port: int(t.dest.TCP)}, "err", err)
			delete(s.static, t.dest.ID)
		case nil:
			s.commonDialing[id] = t.flags
			newtasks = append(newtasks, t)
		}

	}

	// If we don't have any peers whatsoever, try to dial a random bootnode. This
	// scenario is useful for the testnet (and private networks) where the discovery
	// table might be full of mostly bad peers, making it hard to find good ones.
	if len(peers) == 0 && len(s.bootnodes) > 0 && needDynDials > 0 && now.Sub(s.start) > fallbackInterval {
		bootnode := s.bootnodes[0]
		s.bootnodes = append(s.bootnodes[:0], s.bootnodes[1:]...)
		s.bootnodes = append(s.bootnodes, bootnode)

		if addDial(dynDialedConn, bootnode) {
			needDynDials--
		}
	}
	// Use random nodes from the table for half of the necessary
	// dynamic dials.
	randomCandidates := needDynDials / 2
	if openTopNet {
		if randomCandidates > 0 {
			n := s.ntab.ReadRandomNodes(s.commonRandomNodes, discover.ConsNet)
			for i := 0; i < randomCandidates && i < n; i++ {
				if addDial(dynDialedConn, s.commonRandomNodes[i]) {
					needDynDials--
				}
			}
		}
	}

	if randomCandidates > 0 {
		n := s.ntab.ReadRandomNodes(s.commonRandomNodes, discover.CommNet)
		for i := 0; i < randomCandidates && i < n; i++ {
			if addDial(dynDialedConn, s.commonRandomNodes[i]) {
				needDynDials--
			}
		}
	}

	// Create dynamic dials from random lookup results, removing tried
	// items from the result buffer.
	i := 0
	for ; i < len(s.commonLookupBuf) && needDynDials > 0; i++ {
		if addDial(dynDialedConn, s.commonLookupBuf[i]) {
			s.commonLookupBuf = append(s.commonLookupBuf[:i], s.commonLookupBuf[i+1:]...)
			needDynDials--
		}

	}

	// Launch a discovery lookup if more candidates are needed.
	if len(s.commonLookupBuf) < needDynDials && !s.commonLookupRunning {
		s.commonLookupRunning = true
		newtasks = append(newtasks, &discoverTask{netType: discover.CommNet})
	}
	// Launch a timer to wait for the next node to expire if all
	// candidates have been tried and no task is currently active.
	// This should prevent cases where the dialer logic is not ticked
	// because there are no pending events.
	if nRunning == 0 && len(newtasks) == 0 && s.commonHist.Len() > 0 {
		t := &waitExpireTask{s.commonHist.min().exp.Sub(now), discover.CommNet}
		newtasks = append(newtasks, t)
	}
	if openTopNet {
		i := 0
		for ; i < len(s.topLookupBuf) && needDynDials > 0; i++ {
			if addDial(dynDialedConn, s.topLookupBuf[i]) {
				needDynDials--
			}
		}

		// Launch a discovery lookup if more candidates are needed.
		if len(s.topLookupBuf) < needDynDials && !s.topLookupRunning {
			s.topLookupRunning = true
			newtasks = append(newtasks, &discoverTask{netType: discover.ConsNet})
		}

		// Launch a timer to wait for the next node to expire if all
		// candidates have been tried and no task is currently active.
		// This should prevent cases where the dialer logic is not ticked
		// because there are no pending events.
		if nRunning == 0 && len(newtasks) == 0 && s.topHist.Len() > 0 {
			t := &waitExpireTask{s.topHist.min().exp.Sub(now), discover.ConsNet}
			newtasks = append(newtasks, t)
		}

	}

	return newtasks
}

var (
	errSelf             = errors.New("is self")
	errAlreadyDialing   = errors.New("already dialing")
	errAlreadyConnected = errors.New("already connected")
	errRecentlyDialed   = errors.New("recently dialed")
	errNotWhitelisted   = errors.New("not contained in netrestrict whitelist")
)

func (s *dialstate) checkDial(n *discover.Node, peers map[discover.NodeID]*Peer) error {

	_, dialing := s.commonDialing[n.ID]
	switch {
	case dialing:
		return errAlreadyDialing
	case peers[n.ID] != nil:
		return errAlreadyConnected
	case s.ntab != nil && n.ID == s.ntab.Self().ID:
		return errSelf
	case s.netrestrict != nil && !s.netrestrict.Contains(n.IP):
		return errNotWhitelisted
	case s.commonHist.contains(n.ID):
		return errRecentlyDialed
	}
	return nil

}

func (s *dialstate) taskDone(t task, now time.Time) {
	switch t := t.(type) {
	case *dialTask:
		delete(s.commonDialing, t.dest.ID)
	case *discoverTask:
		if t.netType == discover.CommNet {
			s.commonLookupRunning = false
			s.commonLookupBuf = append(s.commonLookupBuf, t.results...)
		} else {
			s.topLookupRunning = false
			s.topLookupBuf = append(s.topLookupBuf, t.results...)
		}

	}
}

func (t *dialTask) Do(srv *Server) {
	log.Debug("remote node start mask", "node", t.dest, "netType", t.netType)
	if t.dest.Incomplete() {
		if !t.resolve(srv) {
			return
		}
	}
	err := t.dial(srv, t.dest)
	if err != nil {
		if t.ds.ntab != nil {
			t.ds.ntab.Delete(t.dest.ID)
		}

		log.Info("failed to connect remote node", "err", err.Error())

	}

}

func (t *dialTask) GetNetType() byte {
	return t.netType

}

// resolve attempts to find the current endpoint for the destination
// using discovery.
//
// Resolve operations are throttled with backoff to avoid flooding the
// discovery network with useless queries for nodes that don't exist.
// The backoff delay resets when the node is found.
func (t *dialTask) resolve(srv *Server) bool {
	if srv.ntab == nil {
		log.Debug("Can't resolve node", "id", t.dest.ID, "err", "discovery is disabled")
		return false
	}
	if t.resolveDelay == 0 {
		t.resolveDelay = initialResolveDelay
	}
	if time.Since(t.lastResolved) < t.resolveDelay {
		return false
	}
	resolved := srv.ntab.Resolve(t.dest.ID)
	t.lastResolved = time.Now()
	if resolved == nil {
		t.resolveDelay *= 2
		if t.resolveDelay > maxResolveDelay {
			t.resolveDelay = maxResolveDelay
		}
		log.Debug("Resolving node failed", "id", t.dest.ID, "newdelay", t.resolveDelay)
		return false
	}
	// The node was found.
	t.resolveDelay = initialResolveDelay
	t.dest = resolved
	log.Debug("Resolved node", "id", t.dest.ID, "addr", &net.TCPAddr{IP: t.dest.IP, Port: int(t.dest.TCP)})
	return true
}

type dialError struct {
	error
}

// dial performs the actual connection attempt.
func (t *dialTask) dial(srv *Server, dest *discover.Node) error {
	fd, err := srv.Dialer.Dial(dest)

	if err != nil {
		log.Info("tcp connect failed", "tcpconnectionErr", err)
		return &dialError{err}
	}
	mfd := newMeteredConn(fd, false)
	return srv.SetupConn(mfd, t.flags, dest, t.netType)
}

func (t *dialTask) String() string {
	return fmt.Sprintf("%v %x %v:%d", t.flags, t.dest.ID[:8], t.dest.IP, t.dest.TCP)
}

func (t *discoverTask) Do(srv *Server) {
	// newTasks generates a lookup task whenever dynamic dials are
	// necessary. Lookups need to take some time, otherwise the
	// event loop spins too fast.

	next := srv.lastLookup.Add(lookupInterval)
	if now := time.Now(); now.Before(next) {
		time.Sleep(next.Sub(now))
	}
	srv.lastLookup = time.Now()
	var target discover.NodeID
	rand.Read(target[:])
	t.results = srv.ntab.Lookup(target, t.netType)
}

func (t *discoverTask) GetNetType() byte {
	return t.netType
}

func (t *discoverTask) String() string {
	s := "discovery lookup"
	if len(t.results) > 0 {
		s += fmt.Sprintf(" (%d results)", len(t.results))
	}
	return s
}

func (t waitExpireTask) Do(*Server) {
	time.Sleep(t.Duration)
}

func (t waitExpireTask) GetNetType() byte {
	return t.netType
}
func (t waitExpireTask) String() string {
	return fmt.Sprintf("wait for dial hist expire (%v)", t.Duration)
}

// Use only these methods to access or modify dialHistory.
func (h dialHistory) min() pastDial {
	return h[0]
}
func (h *dialHistory) add(id discover.NodeID, exp time.Time) {
	heap.Push(h, pastDial{id, exp})
}
func (h dialHistory) contains(id discover.NodeID) bool {
	for _, v := range h {
		if v.id == id {
			return true
		}
	}
	return false
}
func (h *dialHistory) expire(now time.Time) {
	for h.Len() > 0 && h.min().exp.Before(now) {
		heap.Pop(h)
	}
}

// heap.Interface boilerplate
func (h dialHistory) Len() int           { return len(h) }
func (h dialHistory) Less(i, j int) bool { return h[i].exp.Before(h[j].exp) }
func (h dialHistory) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *dialHistory) Push(x interface{}) {
	*h = append(*h, x.(pastDial))
}
func (h *dialHistory) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
