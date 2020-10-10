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

// +build none

/*

   The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

       go run mkalloc.go genesis.json

*/
package main

import (
	"fmt"
	"github.com/eminer-pro/go-eminer/core/types"
	"github.com/eminer-pro/go-eminer/rlp"
	"strconv"
)

type genesisAgents []types.Candidate

const (
	// em 2.1 billion
	firstDelegateVoteNumber  = uint64(200_000_000)
	secondDelegateVoteNumber = uint64(100_000_000)
	thirdDelegateVoteNumber  = uint64(50_000_000)
	fourthDelegateVoteNumber = uint64(20_000_000)
	fifthDelegateVoteNumber  = uint64(10_000_000)
	sixthDelegateVoteNumber  = uint64(6_000_000)
)

func main() {
	var list genesisAgents
	candidateList := mainNetAgents()
	list = append(list, candidateList...)

	data, err := rlp.EncodeToBytes(list)
	if err != nil {
		panic(err)
	}
	result := strconv.QuoteToASCII(string(data))
	fmt.Println("const agentData =", result)
}

func mainNetAgents() []types.Candidate {
	return []types.Candidate{
		{"EMf6565d95149e0bf80d466b200dc0e4356b27a01f", firstDelegateVoteNumber, "eminer", 1589277429},
		{"EM7277671ddc67523cf7fd8e559d0c7760cf1949ed", firstDelegateVoteNumber, "Zeus", 1589277429},
		{"EM23edc4b3c7a95a748e1162d796178ae7395f2b3a", firstDelegateVoteNumber, "Hera", 1589277429},
		{"EM60c79b43fac8bc353f197b380b35a4c02eeb54c7", firstDelegateVoteNumber, "Poseidon", 1589277429},
		{"EM2ef5513bcafacb92d5f03ab62a1b676f0eeadd98", firstDelegateVoteNumber, "Hades", 1589277429},
		{"EM50cb30135bdf3eb9cd70501c55ad80c38cceca3f", firstDelegateVoteNumber, "Demeter", 1589277429},
		{"EM0d8c77b49c1651268ea7cefee4671892f36ad129", firstDelegateVoteNumber, "Ares", 1589277429},
		{"EM0d3e25cd5ed1f7a474e656b55ece6f1ade7eefb5", firstDelegateVoteNumber, "Athene", 1589277429},
		{"EM01e7d536e3c6133e33e075a21d06f67d79868b58", firstDelegateVoteNumber, "Apollo", 1589277429},
		{"EM4c6e5d52b6998c4403791261c2b1c17c4a34c7e5", firstDelegateVoteNumber, "Artemis", 1589277429},

		{"EMf69a04a0640876db79b40e6b155cf16eb1d08bbe", secondDelegateVoteNumber, "Aphrodite", 1589277429},
		{"EMc38a0331d9821c43bf5b133b95697fb9cc2ac54a", secondDelegateVoteNumber, "Hermes", 1589277429},
		{"EM17409df06614da0c5a1f14ffc4de20b69b13aed9", secondDelegateVoteNumber, "Hephaestus", 1589277429},
		{"EM203eb0dfadd511d721110f73ae4b208793627c0d", secondDelegateVoteNumber, "Andromeda", 1589277429},
		{"EMe499fa7a1e465ee3d532753b198ec898a3e536e6", secondDelegateVoteNumber, "Antlia", 1589277429},
		{"EM31df20a0d6f66ac7f41d307c496da8162c264b2a", secondDelegateVoteNumber, "Apus", 1589277429},
		{"EM22735d69cf1cd0436e917acd967b2c2c40db8cb2", secondDelegateVoteNumber, "Aquarius", 1589277429},
		{"EMf54cd0a6c0fdf41290282ab0b2d6b8327458d96f", secondDelegateVoteNumber, "Aquila", 1589277429},
		{"EM466ff4c91813fc83d8504f19e3b85f39b81ab635", secondDelegateVoteNumber, "Ara", 1589277429},
		{"EM843f0f8b0cbffbd6b4757d573cc929270feec30d", secondDelegateVoteNumber, "Aries", 1589277429},
		{"EM4a5aca61fe77292cea2bbaa6309120c1fd59f1c7", secondDelegateVoteNumber, "Auriga", 1589277429},
		{"EM66e9b19eabf2523f7afd7fcda9a18eb4af01262e", secondDelegateVoteNumber, "Bootes", 1589277429},
		{"EM629810040024b5aad277beecb018822d84ab675f", secondDelegateVoteNumber, "Caelum", 1589277429},
		{"EM6e87b594159c7030fd8ba701ca164d14af6f49da", secondDelegateVoteNumber, "Camelopardalis", 1589277429},
		{"EM4088526fadf6e448972fa7c5ece36211fb8173ba", secondDelegateVoteNumber, "Cancer", 1589277429},
		{"EMd56cff7e54ec44125b1bbb2940aeb379333e4c6e", secondDelegateVoteNumber, "Canes Venatici", 1589277429},
		{"EM8d8b0872a3c338efb7f290f7b62e246f4667fa3e", secondDelegateVoteNumber, "Canis Major", 1589277429},
		{"EMcf28dd4cd2f99f909a532c3d2f59b3e78df15e24", secondDelegateVoteNumber, "Canis Minor", 1589277429},
		{"EM7ff30ff09bc82e172debe4f45c187a6364f0196b", secondDelegateVoteNumber, "Capricorn", 1589277429},
		{"EMc7670bb139e0734e2ec1cf6bc60e2d7c2e4a00cf", secondDelegateVoteNumber, "Carina", 1589277429},

		{"EM10ddb37671881369471897c4db2c248477618b4f", thirdDelegateVoteNumber, "Cassiopeia", 1589277429},
		{"EMb6f4ed5ad25fab4cc20a3d425cb6c809aa804153", thirdDelegateVoteNumber, "Centaurus", 1589277429},
		{"EM74109d2b35e0b9896aa4131640de4a5136e53475", thirdDelegateVoteNumber, "Cepheus", 1589277429},
		{"EM9036d0e161d13d76171aba63facd223de368e098", thirdDelegateVoteNumber, "Cetus", 1589277429},
		{"EM1240ff5ded573356c5444476d9f2990bec79ddc7", thirdDelegateVoteNumber, "Chamaeleon", 1589277429},
		{"EMc6a5b9605ed0decdd5a7c837b7f7c9fe46030e12", thirdDelegateVoteNumber, "Circinus", 1589277429},
		{"EM32d78a9bf1fd0f38d2d5c75e2634d66fc3a51fa2", thirdDelegateVoteNumber, "Columba", 1589277429},
		{"EM9633b16c9bcbda7bc29b28a7cdf4372876f2b335", thirdDelegateVoteNumber, "Coma", 1589277429},
		{"EMa265b63d18c1b209a45b7fdadb5000e2e8df1748", thirdDelegateVoteNumber, "Corona Australis", 1589277429},
		{"EM0f85906a2c8c8340fa1105a661950694ed1d147c", thirdDelegateVoteNumber, "Corona Borealis", 1589277429},
		{"EM43f4fc95485110683f745cd4a6362497c832a774", thirdDelegateVoteNumber, "Corvus", 1589277429},
		{"EM9e17fd45a2ac3853cfa1683e8ab8d75388cf0023", thirdDelegateVoteNumber, "Crater", 1589277429},
		{"EMb0006161c2f6b08b411a4b818d6f2fdca451d30a", thirdDelegateVoteNumber, "Crux", 1589277429},
		{"EM6d6fcd0962ec76c9348741ae94286ef1e81e4719", thirdDelegateVoteNumber, "Cygnus", 1589277429},
		{"EM8c2e475ce5770fe3efea80d10216d2e84b450fcc", thirdDelegateVoteNumber, "Delphinus", 1589277429},
		{"EM2f5313b71fc8402e55cbbe7f6679180aaca66ea5", thirdDelegateVoteNumber, "Dorado", 1589277429},
		{"EMb8ad18f0ecb51311da15e6d545917d64f7438a57", thirdDelegateVoteNumber, "Draco", 1589277429},
		{"EMdab043561dcfd59b086607d4ed5cdd0c22f93b25", thirdDelegateVoteNumber, "Equuleus", 1589277429},
		{"EMf301dd752183e2af49f1f93e9e3d42bfb8da23f2", thirdDelegateVoteNumber, "Eridanus", 1589277429},
		{"EM94d69bb98b0bf54abccbac667b6bf73002c46437", thirdDelegateVoteNumber, "Fornax", 1589277429},
		{"EMd5557cc9faaf52ea88015cede16d1627e14c2e4b", thirdDelegateVoteNumber, "Gemini", 1589277429},
		{"EMc23aa6734c50d48446229b13b3e6eaf48806b4d4", thirdDelegateVoteNumber, "Grus", 1589277429},
		{"EMfc0906cbda5ae7d5e3b2133cb7086d1aad85a8ee", thirdDelegateVoteNumber, "Hercules", 1589277429},
		{"EMf935cda278b1c1093cb5c07fbeb38bb706752192", thirdDelegateVoteNumber, "Horologium", 1589277429},
		{"EM9b8699827c98aaf2b18d0cab0f0db09229847855", thirdDelegateVoteNumber, "Hydra", 1589277429},
		{"EM2d535f1a8006e743ccb18a560164926bdb8d2c94", thirdDelegateVoteNumber, "Hydrus", 1589277429},
		{"EMa493c6a36bf4f58123099b010660ae6f49cabb50", thirdDelegateVoteNumber, "Indus", 1589277429},
		{"EM355f60d8643fe31131846d4d26a7e379be9e8d19", thirdDelegateVoteNumber, "Lacerta", 1589277429},
		{"EM8f086cc87de326240a61503e04c62e815bbfd593", thirdDelegateVoteNumber, "Leo", 1589277429},
		{"EMd748bd0838350e2709a92dcd38390889c320447e", thirdDelegateVoteNumber, "Leo Minor", 1589277429},
		{"EM3feb07157313fb704d6c0f589c5e979527a66952", thirdDelegateVoteNumber, "Lepus", 1589277429},
		{"EM9450a8b2ada1a41515dd88943453cfc4150fe284", thirdDelegateVoteNumber, "Libra", 1589277429},
		{"EMe99b3971275bb67671ef0c0c9ffafaa4b9a10755", thirdDelegateVoteNumber, "Lupus", 1589277429},
		{"EM5eefa3b29ff8aad603bb0bb31485038a5585ee67", thirdDelegateVoteNumber, "Lynx", 1589277429},
		{"EM5cb0b6c80d392f2eee126aee4fd2afe4ec274856", thirdDelegateVoteNumber, "Lyra", 1589277429},

		{"EM01bf2a7963bb412ef63c41b30ff8f1cb2dbfbcb2", fourthDelegateVoteNumber, "Mensa", 1589277429},
		{"EMf95c423c445fedd8386ff84a0b40e4188e8c5ed9", fourthDelegateVoteNumber, "Microscopium", 1589277429},
		{"EM16b430f7c49b277fa5ccd60bc612adad74d929e5", fourthDelegateVoteNumber, "Monoceros", 1589277429},
		{"EM68b7da86c3d3bc4f0740cf982a501e953469457f", fourthDelegateVoteNumber, "Musca", 1589277429},
		{"EM76f4d178eb775c0d64a86481a4de10560653c3fd", fourthDelegateVoteNumber, "Norma", 1589277429},
		{"EM924f15a1ba7bc91964babad9ebb27667d957eec9", fourthDelegateVoteNumber, "Octans", 1589277429},
		{"EM705c5c68086881e91b1e83c0ce273ac34a368a83", fourthDelegateVoteNumber, "Ophiuchus", 1589277429},
		{"EM5d5e6e92951296384dcf488b9a8250403d44cec9", fourthDelegateVoteNumber, "Orion", 1589277429},
		{"EM8568e6cbadb0a771158ab9173f283a16323b42d8", fourthDelegateVoteNumber, "Pavo", 1589277429},
		{"EM22a07805d9609139f494ad734b48b8a690337f16", fourthDelegateVoteNumber, "Pegasus", 1589277429},
		{"EM071110bd008fcaf7b81631c5901ca93d8b3eab97", fourthDelegateVoteNumber, "Perseus", 1589277429},
		{"EM2709f926f771d2cda5d769600fee1191f931d932", fourthDelegateVoteNumber, "Phoenix", 1589277429},
		{"EM3fc0577434d6dfe23804c272f3c7ebb231241cd4", fourthDelegateVoteNumber, "Pictor", 1589277429},
		{"EM8e23ad4af71e945a584ad5b4426de0078e8c8719", fourthDelegateVoteNumber, "Pisces", 1589277429},
		{"EMf204dc6e1c34b3714364f86f561bcd23d6a5a832", fourthDelegateVoteNumber, "Piscis Austrinus", 1589277429},
		{"EM5e20799a50833e38615f0f107d140f856c286c18", fourthDelegateVoteNumber, "Puppis", 1589277429},
		{"EMfd840a498e5e7ee0d8135f89bf7ca24513b924f1", fourthDelegateVoteNumber, "Pyxis", 1589277429},
		{"EMf04f779088756369db50ff45e7cd7641ee969778", fourthDelegateVoteNumber, "Reticulum", 1589277429},
		{"EM31a3214c739c9bcd5b83fdd639dd0d591010623c", fourthDelegateVoteNumber, "Sagitta", 1589277429},
		{"EM11ad5c9fd68d4d2bcd48596f8065ea0812372878", fourthDelegateVoteNumber, "Sagittarius", 1589277429},
		{"EM9952d9754684429a7ce57bbaeaab38fdaee96adb", fourthDelegateVoteNumber, "Scorpius", 1589277429},
		{"EM26b27187d9257bb3d2a8e7ce05b0b4eeac9cf5ad", fourthDelegateVoteNumber, "Sculptor", 1589277429},
		{"EM2ff6569c80e11a04093cdcb7f80928fbb1b1df2a", fourthDelegateVoteNumber, "Scutum", 1589277429},
		{"EMfe2dd9129716efe48bc946abc16dcf4a4a7ac4e1", fourthDelegateVoteNumber, "Serpens", 1589277429},
		{"EM1decb9cdd2c7ee4d4ef56c377f68cad2408b439c", fourthDelegateVoteNumber, "Sextans", 1589277429},
		{"EM2dedcf81a471a778f19d86f28e15781af723b341", fourthDelegateVoteNumber, "Taurus", 1589277429},
		{"EM757073628db6c6a77f5fde7720ca5cca55a1a425", fourthDelegateVoteNumber, "Telescopium", 1589277429},
		{"EM7febe46dca288190f24776648b1a06756a87cf06", fourthDelegateVoteNumber, "Triangulum", 1589277429},
		{"EMeb948117458c5fced54c7a71c7cad915d72078cc", fourthDelegateVoteNumber, "Triangulum Australe", 1589277429},
		{"EM95b26f4edb47f9ed677fe5effd1f0020a816cb37", fourthDelegateVoteNumber, "Tucana", 1589277429},

		{"EM086524e66702ce21bb9ba127f346054d36a40f0e", fifthDelegateVoteNumber, "Ursa Major", 1589277429},
		{"EM2369c440541d962d26ae232f6733c3dd5569aaac", fifthDelegateVoteNumber, "Ursa Minor", 1589277429},
		{"EM34db0e7b9ce83f0627d16299e1c2567d9231921b", fifthDelegateVoteNumber, "Vela", 1589277429},
		{"EM267a2f33c924a3b1d751a86689d4a0667341178b", fifthDelegateVoteNumber, "Virgo", 1589277429},
		{"EM2b11346cc600cc4105280ad7dde4ae5f790ee68a", fifthDelegateVoteNumber, "Volans", 1589277429},
		{"EM46fc5280327aaa1ddc79d368ab60af307a54d1b3", fifthDelegateVoteNumber, "Vulpecula", 1589277429},

		{"EM8ea2354ba012628dd1dad9e44500a70075664a16", sixthDelegateVoteNumber, "Vulpecula", 1589277429},
	}
}

func mainTestNetAgents() []types.Candidate {
	candidateList := []types.Candidate{
		{"0xae6951c73938d009835ffe56649083a24b823c42", uint64(2000000), "em-node1", 1544421807},
		{"0x8f0ccc912c0ec8ce7238044c6976abe146914ff5", uint64(2000000), "em-node2", 1544421807},
		{"0xf37be5b5eaff9cf82ae59b7bb2c106e9a5cb1523", uint64(2000000), "em-node3", 1544421807},
		{"0x5844650159455c08ebad1e719c0ed85113d9a91a", uint64(2000000), "em-node4", 1544421807},
		{"0xb78ec5b82019b6adb3026443a850a3f25d66e813", uint64(2000000), "em-node5", 1544421807},
		{"0xfe07c4020a86a9f902bd24c762fc02df1e78312f", uint64(2000000), "em-node6", 1544421807},
		{"0x8e218afb0aeb913c3b9aefa03a3194b95f0be2a8", uint64(2000000), "em-node7", 1544421807},
		{"0xb3ccd4bd10a19f0f0b68d86ecf8c8ba707472d47", uint64(2000000), "em-node8", 1544421807},
		{"0xf848390987333c7ee7fd0b63d5ef1b3fa7d1cca2", uint64(2000000), "em-node9", 1544421807},
		{"0x75de393d3e2b4671257e130b9f37b6e55de1d116", uint64(2000000), "em-node10", 1544421807},
		{"0xc3fd54f8031a068aee81be222d68cc31a7a93c5f", uint64(2000000), "em-node11", 1544421807},
		{"0x9c2490702a9f779ff93697c3935c1686d0ec481a", uint64(2000000), "em-node12", 1544421807},
		{"0x0a484d0a330f662769846873a3985192febdbbda", uint64(2000000), "em-node13", 1544421807},
		{"0x62ca04d697f23b86d57d9d41d904b7a5e808873d", uint64(2000000), "em-node14", 1544421807},
		{"0x9b5bc9b6047e090837b1bdc7a5ce9ab3305e211d", uint64(2000000), "em-node15", 1544421807},
		{"0x5b04e79a854435d3c4c849f1deac066404a3b266", uint64(2000000), "em-node16", 1544421807},
		{"0x21e8e874b78e150f5978b2ad72e3e52b120b8c79", uint64(2000000), "em-node17", 1544421807},
		{"0x3f1cd2621dd662ed2a49f76fd3e2c9377ff8e1f8", uint64(2000000), "em-node18", 1544421807},
		{"0xc71c00c652208621b60339a5c4da14edd48b297c", uint64(2000000), "em-node19", 1544421807},
		{"0x58cc16fe47a25898102e9aa7bacd9d743f3834a7", uint64(2000000), "em-node20", 1544421807},
		{"0x290183e728c68bda990939faca666f71fdb0729a", uint64(2000000), "em-node21", 1544421807},
		{"0x9c218da1360bc9a6953f7b1bb403908a94204375", uint64(2000000), "em-node22", 1544421807},
		{"0xe32c0383f7b945a1c8b85a07135f376fe0e0a6d6", uint64(2000000), "em-node23", 1544421807},
		{"0xb4cfd6e65d5b77f5c5f1764cbe6ffca3f734f4ea", uint64(2000000), "em-node24", 1544421807},
		{"0x5e2e87067d5428b7d2a281ef0e27dcd889cb87a5", uint64(2000000), "em-node25", 1544421807},
		{"0x4aabbbcc6f07e88396f85eb99e777bb25c540720", uint64(2000000), "em-node26", 1544421807},
		{"0xf99c1fd103b9beeb6bf5c4d76568812898f7a7b9", uint64(2000000), "em-node27", 1544421807},
		{"0xe27636f648505609ecfd622c9cdef371f396a60b", uint64(2000000), "em-node28", 1544421807},
		{"0xeaf37aa9392b32287d6a168ff7986e33169c39cd", uint64(2000000), "em-node29", 1544421807},
		{"0xe6b8c9bdb76e30eb11fc890718413f1eb32bdf8e", uint64(2000000), "em-node30", 1544421807},
		{"0x7a6b0b4b34e76122a9a0085a3e9e4feafec11470", uint64(2000000), "em-node31", 1544421807},
		{"0x689d2223743547a3e9c016cefcbac6d04de009ef", uint64(2000000), "em-node32", 1544421807},
		{"0x1b6b92e6891cbaed468fcb21db47f1d6a330bb2c", uint64(2000000), "em-node33", 1544421807},
		{"0x1f9cf1b6d66130de4cef4b4ba2c27427d1d281be", uint64(2000000), "em-node34", 1544421807},
		{"0x255d1f8eeb0bc29937915ea5adf6774c7df220e4", uint64(2000000), "em-node35", 1544421807},
		{"0xebbbac11a3861b70678a83b48a3f96403232fcb9", uint64(2000000), "em-node36", 1544421807},
		{"0x2a4e995b17a703e65df494d151c7b6582ef2732e", uint64(2000000), "em-node37", 1544421807},
		{"0xdc6a079ad488dbd2b2aa73165bc8a3af33cb0996", uint64(2000000), "em-node38", 1544421807},
		{"0x6a00b9059a095473ea410d9aa185025cf395e2aa", uint64(2000000), "em-node39", 1544421807},
		{"0x91bf699159786134843764861967de8bfafb6323", uint64(2000000), "em-node40", 1544421807},
		{"0x434829c35f3ed551eacc238f6d9aee757dc4d14e", uint64(2000000), "em-node41", 1544421807},
		{"0xb69fb7df43a9b5131c0e9ab25dd005530cacdd86", uint64(2000000), "em-node42", 1544421807},
		{"0xef8c732b3ba657da11b00a7ee0871cda50e433dd", uint64(2000000), "em-node43", 1544421807},
		{"0xdecfb4a2af72474c60705f7fae82ee4c6ed306f4", uint64(2000000), "em-node44", 1544421807},
		{"0x48165ab216a3436cfa798279bd13ead7c66aa508", uint64(2000000), "em-node45", 1544421807},
		{"0xfb9b90ef7d29ca05a9dbf51382708bbe6b1cab1d", uint64(2000000), "em-node46", 1544421807},
		{"0xc3025d450debcf8d88d8f2109fdc6a672c102789", uint64(2000000), "em-node47", 1544421807},
		{"0x056d0911ad19e1f8ceeafb097364d699c59cc64e", uint64(2000000), "em-node48", 1544421807},
		{"0xffd7b90cd18aee781ecb824e4bff5af57f76a734", uint64(2000000), "em-node49", 1544421807},
		{"0x72823a6016ea77d30238df3bbe9f22cd429d4ae8", uint64(2000000), "em-node50", 1544421807},
		{"0x7db8ff022d12273e95d4d9e232c4b3cdf7158dfe", uint64(2000000), "em-node51", 1544421807},
		{"0xf205583e5f7cac53336f2288b8db1b247978df88", uint64(2000000), "em-node52", 1544421807},
		{"0xca6fa5fb62c68758e9790005412c36aa3e315d2e", uint64(2000000), "em-node53", 1544421807},
		{"0xde5d2f5ae46d9339b796c1dad47ae25a917ab5a3", uint64(2000000), "em-node54", 1544421807},
		{"0xb32cfef47545fdcfe109543971cef6626528aea4", uint64(2000000), "em-node55", 1544421807},
		{"0x3e9b3ce3a11f00331ec39cf3feb4b6dc42ba3671", uint64(2000000), "em-node56", 1544421807},
		{"0xa584bbe89bd46d86925294701bcbd1b4158abf12", uint64(2000000), "em-node57", 1544421807},
		{"0x13042f9b0f6da4f598e56a9bb8249baacb9f24ae", uint64(2000000), "em-node58", 1544421807},
		{"0xccdaf20553d819faf64e4947fd4e94f65f20df5d", uint64(2000000), "em-node59", 1544421807},
		{"0x967062fff19489237fdce95432e0d8c0a67ebae4", uint64(2000000), "em-node60", 1544421807},
		{"0xc5f4ae494d7bd558b2035b63318b2a2cc47b0c2b", uint64(2000000), "em-node61", 1544421807},
		{"0xc1407b1974cc0ed2f6c1006c222b4b9893174323", uint64(2000000), "em-node62", 1544421807},
		{"0xb4008ab135ec7ced2b5b05216ef74bf1cd835170", uint64(2000000), "em-node63", 1544421807},
		{"0xdfd00dba4589b55dbaf812b5634b374a774b02ab", uint64(2000000), "em-node64", 1544421807},
		{"0x580322da3b0878a9b9bf70d746c681d407058952", uint64(2000000), "em-node65", 1544421807},
		{"0xb85a21534edc46fc331f7d3f51ea15c219b9ac31", uint64(2000000), "em-node66", 1544421807},
		{"0xf9bc1507ae2ec3b937d5f6fca42afa3f8cf57e7d", uint64(2000000), "em-node67", 1544421807},
		{"0xc36e16ab11b20c67bce7801247eecc196e13be76", uint64(2000000), "em-node68", 1544421807},
		{"0x7428eff1d1037357983af28829fb59c4fee02d2e", uint64(2000000), "em-node69", 1544421807},
		{"0x2ce4b50f320f81167719d3ba559960172f0f6fe4", uint64(2000000), "em-node70", 1544421807},
		{"0x0e67571505632bb30cab56c915593a3cd2b309f2", uint64(2000000), "em-node71", 1544421807},
		{"0x092fed360946dd1f11de7da0d8a3b731ebd4e5e7", uint64(2000000), "em-node72", 1544421807},
		{"0x69fe4725fb71142725814687aca7190c2dd138d8", uint64(2000000), "em-node73", 1544421807},
		{"0xe409807b63f317d6f223fcab05373413e334f7cf", uint64(2000000), "em-node74", 1544421807},
		{"0xa82c697e57898cc14e9d4270aec4456393b28ae0", uint64(2000000), "em-node75", 1544421807},
		{"0xbf93f54597b481c1b17fe06b7e2caec4e06c938c", uint64(2000000), "em-node76", 1544421807},
		{"0x9e824f10c206d6f6ab9e1c1afd9d00ff8db4bdc8", uint64(2000000), "em-node77", 1544421807},
		{"0xea278feefa5697a35e71da435521f3276086e16b", uint64(2000000), "em-node78", 1544421807},
		{"0xf118a7719460a99c7222554ffdd44e9463aea674", uint64(2000000), "em-node79", 1544421807},
		{"0x62de91d4286dfdd36cbc06c99b39005bb493c552", uint64(2000000), "em-node80", 1544421807},
		{"0x116d54909c42b7588b006a1bdb6bbb5186c8eaf7", uint64(2000000), "em-node81", 1544421807},
		{"0x3386d282ef9334aedc6c664999f5b6a7e7084a2e", uint64(2000000), "em-node82", 1544421807},
		{"0x5454c44dfbc5d60f2575231069b1168c99e4dabe", uint64(2000000), "em-node83", 1544421807},
		{"0x56232f4a299d343aba2620cbc11f7ab731169f96", uint64(2000000), "em-node84", 1544421807},
		{"0xbeb755aae22eace8d8dbbe92e229edf9d086caef", uint64(2000000), "em-node85", 1544421807},
		{"0x26c452286fd9f25251213e056ef1f8533b5468cd", uint64(2000000), "em-node86", 1544421807},
		{"0xa0afac3fc553f41cdf901060db66ef8493a355e2", uint64(2000000), "em-node87", 1544421807},
		{"0x98f78fe0cb1b3f5643d1d5bf54d9fcd2241f6a42", uint64(2000000), "em-node88", 1544421807},
		{"0x649d293e167658e899d8c96cf5f5ac5c5624309e", uint64(2000000), "em-node89", 1544421807},
		{"0xabb431653084c2b7ccac452f94a373eb976aa08f", uint64(2000000), "em-node90", 1544421807},
		{"0xa5f3394b2f11b2e288fa2f4e436cb30dcde521c7", uint64(2000000), "em-node91", 1544421807},
		{"0xd5f15ade48ecfd7425807544b82ee0eb6e272e2d", uint64(2000000), "em-node92", 1544421807},
		{"0xd9d05e7dae8727bb90e3a51e5231cb28aa0b4c2f", uint64(2000000), "em-node93", 1544421807},
		{"0xcd45eecc7c899054fb7ffda9b03d27ad339fd3c5", uint64(2000000), "em-node94", 1544421807},
		{"0x17cf6befd3ebc0ebc4815ad1675f0a17e3ef0e3a", uint64(2000000), "em-node95", 1544421807},
		{"0xb2afa200ef65564393f3446dc779ba8127a793b8", uint64(2000000), "em-node96", 1544421807},
		{"0x09e3b312613cd9ccbe041cc01b3d4e2ac78e6cb4", uint64(2000000), "em-node97", 1544421807},
		{"0xcd75375984cacf5bb9bd03e505ba2a97e4f7a37c", uint64(2000000), "em-node98", 1544421807},
		{"0x409ae4e6c465026bb67e665d7f49cb9bafdf3033", uint64(2000000), "em-node99", 1544421807},
		{"0x1858803562fc7ff10b3f224d7b6d538fa2a69523", uint64(2000000), "em-node100", 1544421807},
		{"0xac735a15e1ec9ff3dda829202a4359dcac4b0cfb", uint64(2000000), "em-node101", 1544421807},
	}
	return candidateList
}
