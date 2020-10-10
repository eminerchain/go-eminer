// Copyright 2020 The go-eminer Authors
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

// Package em implements the eminer-pro protocol.
package em

import (
	"fmt"
	"github.com/eminer-pro/go-eminer/accounts"
	"github.com/eminer-pro/go-eminer/common/hexutil"
	"github.com/eminer-pro/go-eminer/consensus"
	"github.com/eminer-pro/go-eminer/consensus/dpos"
	"github.com/eminer-pro/go-eminer/core"
	"github.com/eminer-pro/go-eminer/core/bloombits"
	"github.com/eminer-pro/go-eminer/core/types"
	"github.com/eminer-pro/go-eminer/core/vm"
	"github.com/eminer-pro/go-eminer/em/downloader"
	"github.com/eminer-pro/go-eminer/em/filters"
	"github.com/eminer-pro/go-eminer/em/gasprice"
	"github.com/eminer-pro/go-eminer/emdb"
	"github.com/eminer-pro/go-eminer/internal/emapi"
	"github.com/eminer-pro/go-eminer/log"
	"github.com/eminer-pro/go-eminer/node"
	"github.com/eminer-pro/go-eminer/p2p"
	"github.com/eminer-pro/go-eminer/params"
	"github.com/eminer-pro/go-eminer/rlp"
	"github.com/eminer-pro/go-eminer/rpc"
	"math/big"
	"runtime"
	"sync"
)

type LesServer interface {
	Start(srvr *p2p.Server)
	Stop()
	Protocols() []p2p.Protocol
	SetBloomBitsIndexer(bbIndexer *core.ChainIndexer)
}

// eminer-pro implements the eminer-pro full node service.
type Dacchain struct {
	config      *Config
	chainConfig *params.ChainConfig

	// Channel for shutting down the service
	shutdownChan  chan bool    // Channel for shutting down the dacchain
	stopDbUpgrade func() error // stop chain db sequential key upgrade

	// Handlers
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	lesServer       LesServer

	// DB interfaces
	chainDb   emdb.Database // Block chain database
	watcherDb emdb.Database // database for watch internal transactions

	dacEngine      consensus.Engine
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

	ApiBackend *DacApiBackend

	gasPrice *big.Int

	networkId     uint64
	netRPCService *emapi.PublicNetAPI

	lock            sync.RWMutex // Protects the variadic fields (e.g. gas price )
	dposTaskManager *DposTaskManager
	dposMiner       *core.DposMiner
}

func (dacchain *Dacchain) AddLesServer(ls LesServer) {
	dacchain.lesServer = ls
	ls.SetBloomBitsIndexer(dacchain.bloomIndexer)
}

// New creates a new eminer-pro object (including the
// initialisation of the common eminer-pro object)
func New(ctx *node.ServiceContext, config *Config) (*Dacchain, error) {
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	chainDb, err := CreateDB(ctx, config, "chaindata")
	if err != nil {
		return nil, err
	}
	var watcherDb emdb.Database
	if config.EnableInterTxWatching {
		watcherDb, err = CreateDB(ctx, config, "watchdata")
		if err != nil {
			return nil, err
		}
	}
	stopDbUpgrade := upgradeDeduplicateData(chainDb)
	chainConfig, genesisHash, _, genesisErr := core.SetupGenesisBlock(chainDb, config.Genesis)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	dac := &Dacchain{
		config:         config,
		chainDb:        chainDb,
		chainConfig:    chainConfig,
		accountManager: ctx.AccountManager,
		shutdownChan:   make(chan bool),
		stopDbUpgrade:  stopDbUpgrade,
		networkId:      config.NetworkId,
		gasPrice:       config.GasPrice,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(chainDb, params.BloomBitsBlocks),
		dacEngine:      CreateDacchainConsensusEngine(),
		watcherDb:      watcherDb,
	}

	log.Info("Initialising eminer-pro protocol", "versions", ProtocolVersions, "network", config.NetworkId)

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
		if bcVersion != core.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run em upgradedb.\n", bcVersion, core.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
	}

	vmConfig := vm.Config{EnablePreimageRecording: config.EnablePreimageRecording, WatchInnerTx: config.EnableInterTxWatching}
	dac.blockchain, err = core.NewBlockChain(chainDb, dac.chainConfig, dac.dacEngine, vmConfig, watcherDb)
	if err != nil {
		return nil, err
	}
	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		dac.blockchain.SetHead(compat.RewindTo)
		core.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}
	dac.bloomIndexer.Start(dac.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}

	dac.txPool = core.NewTxPool(config.TxPool, dac.chainConfig, dac.blockchain)
	dac.dposMiner = core.NewDposMiner(dac.chainConfig, dac, dac.dacEngine)
	dac.dposTaskManager = NewDposTaskManager(ctx, dac.blockchain, dac.accountManager, dac.dposMiner.GetProduceCallback(), dac.dposMiner.GetShuffleHashChan())
	if dac.protocolManager, err = NewProtocolManager(dac.chainConfig, config.SyncMode, config.NetworkId, dac.txPool, dac.dacEngine, dac.blockchain, chainDb, dac.dposTaskManager, dac.dposMiner.GetProduceBlockChan(), dac.dposMiner.AddDelegateWalletCallback, dac.dposMiner.GetDelegateWallets()); err != nil {
		return nil, err
	}

	dac.ApiBackend = &DacApiBackend{dac, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.GasPrice
	}
	dac.ApiBackend.gpo = gasprice.NewOracle(dac.ApiBackend, gpoParams)

	return dac, nil
}

func makeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		// create default extradata
		extra, _ = rlp.EncodeToBytes([]interface{}{
			uint(params.VersionMajor<<16 | params.VersionMinor<<8 | params.VersionPatch),
			"em",
			runtime.Version(),
			runtime.GOOS,
		})
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", params.MaximumExtraDataSize)
		extra = nil
	}
	return extra
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (emdb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := db.(*emdb.LDBDatabase); ok {
		db.Meter("em/db/chaindata/")
	}
	return db, nil
}

func CreateDacchainConsensusEngine() consensus.Engine {
	return dpos.New()
}

// APIs returns the collection of RPC services the dacchain package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (dacchain *Dacchain) APIs() []rpc.API {
	apis := emapi.GetAPIs(dacchain.ApiBackend)

	// Append any APIs exposed explicitly by the consensus engine
	apis = append(apis, dacchain.dacEngine.APIs(dacchain.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "em",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(dacchain.protocolManager.downloader),
			Public:    true,
		}, {
			Namespace: "em",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(dacchain.ApiBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(dacchain),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(dacchain),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(dacchain.chainConfig, dacchain),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   dacchain.netRPCService,
			Public:    true,
		},
	}...)
}

func (dacchain *Dacchain) ResetWithGenesisBlock(gb *types.Block) {
	dacchain.blockchain.ResetWithGenesisBlock(gb)
}

func (dacchain *Dacchain) AccountManager() *accounts.Manager { return dacchain.accountManager }
func (dacchain *Dacchain) BlockChain() *core.BlockChain      { return dacchain.blockchain }
func (dacchain *Dacchain) TxPool() *core.TxPool              { return dacchain.txPool }
func (dacchain *Dacchain) Engine() consensus.Engine          { return dacchain.dacEngine }
func (dacchain *Dacchain) ChainDb() emdb.Database            { return dacchain.chainDb }
func (dacchain *Dacchain) WatcherDb() emdb.Database          { return dacchain.watcherDb }
func (dacchain *Dacchain) IsListening() bool                 { return true } // Always listening
func (dacchain *Dacchain) EthVersion() int {
	return int(dacchain.protocolManager.SubProtocols[0].Version)
}
func (dacchain *Dacchain) NetVersion() uint64 { return dacchain.networkId }
func (dacchain *Dacchain) Downloader() *downloader.Downloader {
	return dacchain.protocolManager.downloader
}

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (dacchain *Dacchain) Protocols() []p2p.Protocol {
	if dacchain.lesServer == nil {
		return dacchain.protocolManager.SubProtocols
	}
	return append(dacchain.protocolManager.SubProtocols, dacchain.lesServer.Protocols()...)
}

// Start implements node.Service, starting all internal goroutines needed by the
// eminer-pro protocol implementation.
func (dacchain *Dacchain) Start(srvr *p2p.Server) error {
	// Start the bloom bits servicing goroutines
	dacchain.startBloomHandlers()

	// Start the RPC service
	dacchain.netRPCService = emapi.NewPublicNetAPI(srvr, dacchain.NetVersion())

	// Figure out a max peers count based on the server limits
	maxPeers := srvr.MaxPeers
	if dacchain.config.LightServ > 0 {
		maxPeers -= dacchain.config.LightPeers
		if maxPeers < srvr.MaxPeers/2 {
			maxPeers = srvr.MaxPeers / 2
		}
	}
	// Start the networking layer and the light server if requested
	dacchain.protocolManager.Start(maxPeers)
	if dacchain.lesServer != nil {
		dacchain.lesServer.Start(srvr)
	}
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// eminer-pro protocol.
func (dacchain *Dacchain) Stop() error {
	if dacchain.stopDbUpgrade != nil {
		dacchain.stopDbUpgrade()
	}
	dacchain.bloomIndexer.Close()
	dacchain.blockchain.Stop()
	dacchain.protocolManager.Stop()
	if dacchain.lesServer != nil {
		dacchain.lesServer.Stop()
	}
	dacchain.txPool.Stop()
	dacchain.chainDb.Close()
	if dacchain.watcherDb != nil {
		dacchain.watcherDb.Close()
	}
	close(dacchain.shutdownChan)

	return nil
}
