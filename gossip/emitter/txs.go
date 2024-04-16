package emitter

import (
	"time"
        "fmt"
	"github.com/Fantom-foundation/lachesis-base/common/bigendian"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/eventcheck/epochcheck"
	//"github.com/Fantom-foundation/go-opera/eventcheck/gaspowercheck"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/utils"
	"github.com/Fantom-foundation/go-opera/utils/txtime"
)

const (
	TxTurnPeriod        = 2 * time.Second
	TxTurnPeriodLatency = 1 * time.Second
	TxTurnNonces        = 32
)

func max64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func (em *Emitter) maxGasPowerToUse(e *inter.MutableEventPayload) uint64 {
	rules := em.world.GetRules()
	maxGasToUse := rules.Economy.Gas.MaxEventGas
	return maxGasToUse
}

func getTxRoundIndex(now, txTime time.Time, validatorsNum idx.Validator) int {
	passed := now.Sub(txTime)
	if passed < 0 {
		passed = 0
	}
	//fmt.Printf("Transaction wait time and val turn: %v, %v, %v\n", passed, TxTurnPeriod, int((passed / TxTurnPeriod) % time.Duration(validatorsNum)))
	return int((passed / TxTurnPeriod) % time.Duration(validatorsNum))
}



func (em *Emitter) isMyTxTurn(txHash common.Hash, sender common.Address, accountNonce uint64, now time.Time, validators *pos.Validators, me idx.ValidatorID, epoch idx.Epoch) bool {
	txTime := txtime.Of(txHash)

	roundIndex := getTxRoundIndex(now, txTime, validators.Len())
	if roundIndex != getTxRoundIndex(now.Add(TxTurnPeriodLatency), txTime, validators.Len()) {
		// round is about to change, avoid originating the transaction to avoid racing with another validator
		return false
	}

	// generate seed for generating the validators sequence for the tx
	roundsHash := hash.Of(sender.Bytes(), bigendian.Uint64ToBytes(accountNonce/TxTurnNonces), epoch.Bytes())

	// generate the validators sequence for the tx
	rounds := utils.WeightedPermutation(int(validators.Len()), validators.SortedWeights(), roundsHash)

	// take a validator from the sequence, skip offline validators
	for ; roundIndex < len(rounds); roundIndex++ {
		chosenValidator := validators.GetID(idx.Validator(rounds[roundIndex]))
		if chosenValidator == me {
		        fmt.Printf("Validator chosen: %v after %v rounds\n", chosenValidator, roundIndex)
			return true // current validator is the chosen - emit
		}
		if !em.offlineValidators[chosenValidator] {
			return false // chosen validator is online - don't emit
		} else 
		{
		
		// otherwise try next validator in the sequence
		// skippedOfflineValidatorsCounter.Inc(1)
		fmt.Printf("Validator offine: %v\n", chosenValidator)
	}
	}
	return false
}

func (em *Emitter) addTxs(e *inter.MutableEventPayload, sorted *types.TransactionsByPriceAndNonce) {
	maxGasUsed := em.maxGasPowerToUse(e)
	if maxGasUsed <= e.GasPowerUsed() {
		return
	}

	// sort transactions by price and nonce
	rules := em.world.GetRules()
	for tx := sorted.Peek(); tx != nil; tx = sorted.Peek() {
		sender, _ := types.Sender(em.world.TxSigner, tx)
		// check transaction epoch rules
		if epochcheck.CheckTxs(types.Transactions{tx}, rules) != nil {
			sorted.Pop()
			continue
		}
		// check there's enough gas power to originate the transaction
		if tx.Gas() >= e.GasPowerLeft().Min() || e.GasPowerUsed()+tx.Gas() >= maxGasUsed {
			if params.TxGas >= e.GasPowerLeft().Min() || e.GasPowerUsed()+params.TxGas >= maxGasUsed {
				// stop if cannot originate even an empty transaction
				break
			}
			sorted.Pop()
			continue
		}
		// check not conflicted with already originated txs (in any connected event)
		if em.originatedTxs.TotalOf(sender) != 0 {
			sorted.Pop()
			continue
		}
		// my turn, i.e. try to not include the same tx simultaneously by different validators
		if !em.isMyTxTurn(tx.Hash(), sender, tx.Nonce(), time.Now(), em.validators, e.Creator(), em.epoch) {
			sorted.Pop()
			continue
		}
		// check transaction is not outdated
		if !em.world.TxPool.Has(tx.Hash()) {
			sorted.Pop()
			continue
		}
		// add
		e.SetGasPowerUsed(e.GasPowerUsed() + tx.Gas())
		e.SetGasPowerLeft(e.GasPowerLeft().Sub(tx.Gas()))
		e.SetTxs(append(e.Txs(), tx))
		sorted.Shift()
	}
}
