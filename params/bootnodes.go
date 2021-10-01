// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes//TODO
	"enode://cc1ea6bd31782dbabdbe5367b1ff9e355d97be2802331e0bc6bd92c2700e11b306757bcecd757c9958d4c3faf898fdc57f0c86c1739164b460f91a05fcb9bbd5@80.85.84.23:38990",
	"enode://97fbf6c6eb489f6c494af47df33f71f0ac1afa7f9dd4db6faba85d8561b8af330fac4fd556338823c9fc4aa6cb5b601807bc06adfaeba7044754178b87074ba4@139.162.197.220:38992",
	"enode://f048462cd27830f94a5165d2b2a84506e0ebe48d87e8215d4eff1204a107c7db15f7cb3b584b5ebc1f727afe91373ea4dd3a5741c2f5903b3161d61d802f9e29@139.162.200.12:38993",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://f994a9a1000f299dc5fc611cfee816819b63aad2f91c25f7d55c4420700e6b93434df00e9941021b22c03af9cb8901ed97c997c3dff0bf463e19e472e67361c0@147.182.181.76:38991",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
