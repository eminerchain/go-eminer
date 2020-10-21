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

package walletType

import (
	"crypto/ecdsa"
	"github.com/eminerchain/go-eminer/accounts"
)

type WalletWrapper struct {
	Wallet accounts.Wallet
	Pwd    string
}

type DelegateWalletInfo struct {
	Address    string
	PrivateKey *ecdsa.PrivateKey
}
