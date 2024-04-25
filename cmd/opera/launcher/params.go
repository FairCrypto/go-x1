package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

var (
	Bootnodes = map[string][]string{
		"main": {},
		"x1-testnet": {
			"enode://5d1a3b86c8ee297142fb916db77785cc679d4ed10f51cca84c9ee9442d66280a90378d9ca364620bd8cebec0aafb2cf22cd033905afe2028c22b804a2e18d981@bootnode0.x1-testnet.xen.network:5050",
			"enode://703e4f01760a1f432d5e5e9a8a9c86b5dca301563bc9559f8c991de052e438d8072b2e789f7722762ea44c5c06ed87428a25606f278a128d7a39bd5996be9536@bootnode1.x1-testnet.xen.network:5050",
			"enode://873de92942769e57c2360dd3198af2fda9602a9826b1f1a6d36326b6a75c9178c4f13f910c9e715a5ca6344aa62452930fcee18636ab0a95f269dded52f45960@bootnode2.x1-testnet.xen.network:5050",
			"enode://125e59370d9c7f5a195e2b99da9bb4172c8ac7dde74bb47a5b97fdfbeb3e876ee64289fd05b96b2b24006e00c4a6ad990ae159008f006f809966636797ab0054@bootnode3.x1-testnet.xen.network:5050",
		},
		"x1-alphanet": {
			"enode://fab4851576e588468139e884dba7b6518435512e04ec094d3de654e70ea68ba61a40a446b472994fd2e9eeb7de596041bd2ec7678d13cb7e7e7e180958e88f66@bootnode0.x1-alphanet.xen.network:5050",
			"enode://978bf7631880ee70a440f09312022b655f7663c4bdd5fbae3dd5877a4988a5c1e5fbcf3e1f67feb58a4a6480de2a0dc4132d56cf42b00e1326357d58e1c138c4@bootnode1.x1-alphanet.xen.network:5050",
			"enode://32297a35eded4275a197e16324ec8fcfbb26e0440e3156c8c969946c5d30c4f5ab3ac626690f98cbd164549a66788d93bea908e73daa12f3bcc34c3de8e2f48f@bootnode2.x1-alphanet.xen.network:5050",
			"enode://e099e469ffdf146141ce7bb3c8f086ebdc1b16af96eba5ca0d8206d350c17b6af8d1a450666b67aed08b564183eb3c339bfbf5c09208adc085e083c781a7deb9@bootnode3.x1-alphanet.xen.network:5050",
		},
	}

	testnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0x4c4fdf6346c8851355eb305399d05036a0512aea15d1cb9b364a353704d5fbcb"),
		NetworkID:   opera.TestNetworkID,
		NetworkName: "x1-testnet",
	}

	AllowedOperaGenesis = []GenesisTemplate{
		{
			Name:   "x1-testnet with full MPT",
			Header: testnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xe4c9d47ea5aac4beaac9655c8e63257762f1a4bc4e55028765d7b791392beba7"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xc3163030010b8a02bdce6dbb8d6dacb2e9d9a136c3afe03a4700df6473e8252a"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x48c9563f4c42333b1c8c0a6feccd839c5feb3c9507e334ddb088fc2ee67b4641"),
			},
		},
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
