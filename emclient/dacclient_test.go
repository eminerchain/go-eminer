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

package emclient

import _ "github.com/eminerchain/go-eminer"

// Verify that Client implements the dacchain interfaces.
var (
	_ = dacchain.ChainReader(&Client{})
	_ = dacchain.TransactionReader(&Client{})
	_ = dacchain.ChainStateReader(&Client{})
	_ = dacchain.ChainSyncReader(&Client{})
	_ = dacchain.ContractCaller(&Client{})
	_ = dacchain.GasEstimator(&Client{})
	_ = dacchain.GasPricer(&Client{})
	_ = dacchain.LogFilterer(&Client{})
	
	_ = dacchain.PendingStateReader(&Client{})
	// _ = dacchain.PendingStateEventer(&Client{})
	_ = dacchain.PendingContractCaller(&Client{})
)
