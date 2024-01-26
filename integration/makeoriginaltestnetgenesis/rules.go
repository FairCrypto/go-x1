package makeoriginaltestnetgenesis

import (
	"encoding/json"
	"github.com/Fantom-foundation/go-opera/opera"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	ethparams "github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/opera/contracts/evmwriter"
)

const (
	TestNetworkID       uint64 = 0x31CE5
	DefaultEventGas     uint64 = 28000
	berlinBit                  = 1 << 0
	londonBit                  = 1 << 1
	llrBit                     = 1 << 2
	TestnetStartBalance        = 328333333
	TestnetStartStake          = 5000000
	TestnetGenesisTime         = inter.Timestamp(1700783108 * time.Second)
)

var DefaultVMConfig = vm.Config{
	StatePrecompiles: map[common.Address]vm.PrecompiledStateContract{
		evmwriter.ContractAddress: &evmwriter.PreCompiledContract{},
	},
}

var GenesisValidators = []GenesisValidator{
	{
		"0x1149aD69030084b780C5c375b252E73235AAe0d0",
		"046e4a62824c79b42995e1144d6650dfc673029d4670dcbbdadce57f630a87e613b10cacb66f0f65995dfeedb7339af34c7d5e2031adc621c6bc0df78549726060",
	},
	{
		"0x9c11DafF4913c68838ce7ce6969b12BaBff4318b",
		"04dfc5e6a7594905af4ea831847367e22e9a02c2669d5b00407800e616e1504ab8c847d1e42c118230e593c010e0466b33410450d01811700d562e14a98c521b8f",
	},
	{
		"0xa12f1025aF20f6C13385CdCFE5fc897496F87d98",
		"04b2366cb5269d81cd00e0db9f808c43f614b44dc3abd25f54b025af498a579e7dd13de4f7573360690b23b22b2527a518ecf37119b8a03c8afc9fdbdbac98667e",
	},
}

type GenesisValidator struct {
	AccountAddress  string
	ValidatorPubKey string
}

// Rules describes opera net.
// Note keep track of all the non-copiable variables in Copy()
type Rules opera.Rules

// EvmChainConfig returns ChainConfig for transactions signing and execution
func (r Rules) EvmChainConfig(hh []opera.UpgradeHeight) *ethparams.ChainConfig {
	cfg := *ethparams.AllEthashProtocolChanges
	cfg.ChainID = new(big.Int).SetUint64(r.NetworkID)
	cfg.BerlinBlock = nil
	cfg.LondonBlock = nil
	for i, h := range hh {
		height := new(big.Int)
		if i > 0 {
			height.SetUint64(uint64(h.Height))
		}
		if cfg.BerlinBlock == nil && h.Upgrades.Berlin {
			cfg.BerlinBlock = height
		}
		if !h.Upgrades.Berlin {
			cfg.BerlinBlock = nil
		}

		if cfg.LondonBlock == nil && h.Upgrades.London {
			cfg.LondonBlock = height
		}
		if !h.Upgrades.London {
			cfg.LondonBlock = nil
		}
	}
	return &cfg
}

func TestNetRules() Rules {
	return Rules{
		Name:      "x1-testnet",
		NetworkID: TestNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    DefaultEpochsRules(),
		Economy:   TestnetEconomyRules(),
		Blocks: opera.BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(1 * time.Minute),
		},
		Upgrades: opera.Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

func TestnetEconomyRules() opera.EconomyRules {
	return opera.EconomyRules{
		BlockMissedSlack: 50,
		Gas:              DefaultGasRules(),
		MinGasPrice:      big.NewInt(5e11),
		ShortGasPower:    DefaultShortGasPowerRules(),
		LongGasPower:     DefaulLongGasPowerRules(),
	}
}

func DefaultDagRules() opera.DagRules {
	return opera.DagRules{
		MaxParents:     10,
		MaxFreeParents: 3,
		MaxExtraData:   128,
	}
}

func DefaultGasRules() opera.GasRules {
	return opera.GasRules{
		MaxEventGas:          10000000 + DefaultEventGas,
		EventGas:             DefaultEventGas,
		ParentGas:            2400,
		ExtraDataGas:         25,
		BlockVotesBaseGas:    1024,
		BlockVoteGas:         512,
		EpochVoteGas:         1536,
		MisbehaviourProofGas: 71536,
	}
}

func DefaultEpochsRules() opera.EpochsRules {
	return opera.EpochsRules{
		MaxEpochGas:      1500000000,
		MaxEpochDuration: inter.Timestamp(4 * time.Hour),
	}
}

// DefaulLongGasPowerRules is long-window config
func DefaulLongGasPowerRules() opera.GasPowerRules {
	return opera.GasPowerRules{
		AllocPerSec:        100 * DefaultEventGas,
		MaxAllocPeriod:     inter.Timestamp(60 * time.Minute),
		StartupAllocPeriod: inter.Timestamp(5 * time.Second),
		MinStartupGas:      DefaultEventGas * 20,
	}
}

// DefaultShortGasPowerRules is short-window config
func DefaultShortGasPowerRules() opera.GasPowerRules {
	// 2x faster allocation rate, 6x lower max accumulated gas power
	cfg := DefaulLongGasPowerRules()
	cfg.AllocPerSec *= 2
	cfg.StartupAllocPeriod /= 2
	cfg.MaxAllocPeriod /= 2 * 6
	return cfg
}

func (r Rules) Copy() Rules {
	cp := r
	cp.Economy.MinGasPrice = new(big.Int).Set(r.Economy.MinGasPrice)
	return cp
}

func (r Rules) String() string {
	b, _ := json.Marshal(&r)
	return string(b)
}
