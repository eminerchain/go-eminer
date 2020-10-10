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

package types

import (
	"github.com/eminer-pro/go-eminer/common"
	"github.com/eminer-pro/go-eminer/common/hexutil"
	"math/big"
)

//go:generate gencodec -type InnerTx -field-override innertxMarshaling -out gen_innertx_json.go

type InnerTx struct {
	From    common.Address  `json:"from" gencodec:"required"`
	To      common.Address  `json:"to" gencodec:"to" gencodec:"required"`
	AssetID *common.Address `json:"assetid" rlp:"nil"`
	Value   *big.Int        `json:"value" gencodec:"required"`
}

type innertxMarshaling struct {
	Value *hexutil.Big
}
