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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/eminer-pro/go-eminer/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("em/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("em/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("em/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("em/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("em/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("em/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("em/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("em/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("em/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("em/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("em/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("em/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("em/downloader/states/in")
	stateDropMeter = metrics.NewMeter("em/downloader/states/drop")
)
