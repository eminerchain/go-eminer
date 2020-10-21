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

package params

import "math/big"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main eminer-pro network.
var MainnetBootnodes = []string{
	"enode://d6120a69fc324f739780c1c15a7ece40fc04d57a0dd242f249223b065b79cdb849b486580ec591760470f2ace0424c698ce8f022229d36eb38aa2230caa3a5d3@54.249.211.177:30303",
}

// TestnetBootnodes
var TestnetBootnodes = []string{
	// "enode://7baae2fac6c271737672ad6f15200b60a5b971cd802f85854999536c47bfa644e04eb9dcc8a57333dbd755d77f4797a4dadc0e8c2d0da4f38dd9f422ee593f7f@172.16.20.76:30303",
}

// block reward
var (
	AnnulProfit = 1.15
	AnnulBlockAmount = big.NewInt(3153600)
)
