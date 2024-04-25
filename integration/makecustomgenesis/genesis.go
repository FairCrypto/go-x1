package makecustomgenesis

import (
	"github.com/Fantom-foundation/go-opera/integration/makegenesis"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/drivertype"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/inter/ier"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driver"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driver/drivercall"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driverauth"
	"github.com/Fantom-foundation/go-opera/opera/contracts/evmwriter"
	"github.com/Fantom-foundation/go-opera/opera/contracts/netinit"
	netinitcall "github.com/Fantom-foundation/go-opera/opera/contracts/netinit/netinitcalls"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfclib"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesis/gpos"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	futils "github.com/Fantom-foundation/go-opera/utils"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/Fantom-foundation/lachesis-base/kvdb/memorydb"
	"github.com/Fantom-foundation/lachesis-base/lachesis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
)

type Config struct {
	Enabled              bool
	StartValBalance      uint64
	StartStake           uint64
	NetworkID            uint64
	NetworkName          string
	GenesisTime          uint64
	GenesisValidators    []GenesisValidator
	TreasuryAddress      common.Address
	SfcOwnerAddress      common.Address
	StartTreasuryBalance uint64
}

type GenesisValidator struct {
	AccountAddress  common.Address
	ValidatorPubKey string
}

func DefaultConfig() Config {
	return Config{
		Enabled:           false,
		StartValBalance:   1000000,
		StartStake:        1000000,
		NetworkID:         opera.FakeNetworkID,
		NetworkName:       "FakeNet",
		GenesisTime:       1700783108,
		GenesisValidators: []GenesisValidator{},
		TreasuryAddress:   common.HexToAddress("0x0"),
		SfcOwnerAddress:   common.HexToAddress("0x0"),
	}
}

func GenesisStore(cfg Config) *genesisstore.Store {
	validators := GetValidators(cfg)
	rules := opera.CustomNetRules(cfg.NetworkName, cfg.NetworkID)
	return GenesisStoreWithRules(cfg, futils.ToFtm(cfg.StartValBalance), futils.ToFtm(cfg.StartStake), rules, validators)
}

func GenesisStoreWithValidators(cfg Config, validators gpos.Validators) *genesisstore.Store {
	rules := opera.CustomNetRules(cfg.NetworkName, cfg.NetworkID)
	return GenesisStoreWithRules(cfg, futils.ToFtm(cfg.StartValBalance), futils.ToFtm(cfg.StartStake), rules, validators)
}

func GenesisStoreWithRules(cfg Config, balance, stake *big.Int, rules opera.Rules, validators gpos.Validators) *genesisstore.Store {
	return GenesisStoreWithRulesAndStart(cfg, balance, stake, rules, 2, 1, validators)
}

func GenesisStoreWithRulesAndStart(cfg Config, balance, stake *big.Int, rules opera.Rules, epoch idx.Epoch, block idx.Block, validators gpos.Validators) *genesisstore.Store {
	builder := makegenesis.NewGenesisBuilder(memorydb.NewProducer(""))

	genesisTimestamp := inter.Timestamp(cfg.GenesisTime)

	// add balance to treasury
	builder.AddBalance(cfg.TreasuryAddress, futils.ToFtm(cfg.StartTreasuryBalance))

	// add balances to validators
	var delegations []drivercall.Delegation
	for _, val := range validators {
		log.Info("Validator", "address", val.Address, "pk", val.PubKey, "id", val.ID)
		builder.AddBalance(val.Address, balance)
		delegations = append(delegations, drivercall.Delegation{
			Address:            val.Address,
			ValidatorID:        val.ID,
			Stake:              stake,
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
			Rewards:            new(big.Int),
		})
	}

	// deploy essential contracts
	// pre deploy NetworkInitializer
	builder.SetCode(netinit.ContractAddress, netinit.GetContractBin())
	// pre deploy NodeDriver
	builder.SetCode(driver.ContractAddress, driver.GetContractBin())
	// pre deploy NodeDriverAuth
	builder.SetCode(driverauth.ContractAddress, driverauth.GetContractBin())
	// pre deploy SFC
	builder.SetCode(sfc.ContractAddress, sfc.GetContractBin())
	// pre deploy SFCLib
	builder.SetCode(sfclib.ContractAddress, sfclib.GetContractBin())
	// set non-zero code for pre-compiled contracts
	builder.SetCode(evmwriter.ContractAddress, []byte{0})

	builder.SetCurrentEpoch(ier.LlrIdxFullEpochRecord{
		LlrFullEpochRecord: ier.LlrFullEpochRecord{
			BlockState: iblockproc.BlockState{
				LastBlock: iblockproc.BlockCtx{
					Idx:     block - 1,
					Time:    genesisTimestamp,
					Atropos: hash.Event{},
				},
				FinalizedStateRoot:    hash.Hash{},
				EpochGas:              0,
				EpochCheaters:         lachesis.Cheaters{},
				CheatersWritten:       0,
				ValidatorStates:       make([]iblockproc.ValidatorBlockState, 0),
				NextValidatorProfiles: make(map[idx.ValidatorID]drivertype.Validator),
				DirtyRules:            nil,
				AdvanceEpochs:         0,
			},
			EpochState: iblockproc.EpochState{
				Epoch:             epoch - 1,
				EpochStart:        genesisTimestamp,
				PrevEpochStart:    genesisTimestamp - 1,
				EpochStateRoot:    hash.Zero,
				Validators:        pos.NewBuilder().Build(),
				ValidatorStates:   make([]iblockproc.ValidatorEpochState, 0),
				ValidatorProfiles: make(map[idx.ValidatorID]drivertype.Validator),
				Rules:             rules,
			},
		},
		Idx: epoch - 1,
	})

	blockProc := makegenesis.DefaultBlockProc()
	genesisTxs := GetGenesisTxs(epoch-2, validators, builder.TotalSupply(), delegations, cfg.SfcOwnerAddress)
	err := builder.ExecuteGenesisTxs(blockProc, genesisTxs)
	if err != nil {
		panic(err)
	}

	return builder.Build(genesis.Header{
		GenesisID:   builder.CurrentHash(),
		NetworkID:   rules.NetworkID,
		NetworkName: rules.Name,
	})
}

func txBuilder() func(calldata []byte, addr common.Address) *types.Transaction {
	nonce := uint64(0)
	return func(calldata []byte, addr common.Address) *types.Transaction {
		tx := types.NewTransaction(nonce, addr, common.Big0, 1e10, common.Big0, calldata)
		nonce++
		return tx
	}
}

func GetGenesisTxs(sealedEpoch idx.Epoch, validators gpos.Validators, totalSupply *big.Int, delegations []drivercall.Delegation, driverOwner common.Address) types.Transactions {
	buildTx := txBuilder()
	internalTxs := make(types.Transactions, 0, 15)
	// initialization
	calldata := netinitcall.InitializeAll(sealedEpoch, totalSupply, sfc.ContractAddress, sfclib.ContractAddress, driverauth.ContractAddress, driver.ContractAddress, evmwriter.ContractAddress, driverOwner)
	internalTxs = append(internalTxs, buildTx(calldata, netinit.ContractAddress))
	// push genesis validators
	for _, v := range validators {
		calldata := drivercall.SetGenesisValidator(v)
		internalTxs = append(internalTxs, buildTx(calldata, driver.ContractAddress))
	}
	// push genesis delegations
	for _, delegation := range delegations {
		calldata := drivercall.SetGenesisDelegation(delegation)
		internalTxs = append(internalTxs, buildTx(calldata, driver.ContractAddress))
	}
	return internalTxs
}

func GetValidators(cfg Config) gpos.Validators {
	validators := make(gpos.Validators, 0, len(cfg.GenesisValidators)-1)

	for id, genesisValidator := range cfg.GenesisValidators {
		validators = append(validators, gpos.Validator{
			ID:      idx.ValidatorID(id + 1),
			Address: genesisValidator.AccountAddress,
			PubKey: validatorpk.PubKey{
				Raw:  common.Hex2Bytes(strings.TrimPrefix(genesisValidator.ValidatorPubKey, "0xc0")),
				Type: validatorpk.Types.Secp256k1,
			},
			CreationTime:     inter.Timestamp(cfg.GenesisTime),
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	return validators
}
