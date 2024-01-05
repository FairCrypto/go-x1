package verifier

import (
	"context"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/votemanager"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sync"
	"time"
)

var (
	BatchSize = 50
	GasLimit  = uint64(8000000)
)

type Voter struct {
	vm      *votemanager.Votemanager
	conn    *ethclient.Client
	queue   []votemanager.VoteManagerVotePayload
	ks      *keystore.KeyStore
	account accounts.Account
	chainId uint64
	mu      sync.Mutex
	Syncing bool
}

func NewVoter(conn *ethclient.Client, ks *keystore.KeyStore, account accounts.Account, chainId uint64, vm *votemanager.Votemanager) *Voter {
	var queue []votemanager.VoteManagerVotePayload

	return &Voter{
		vm:      vm,
		conn:    conn,
		queue:   queue,
		account: account,
		ks:      ks,
		chainId: chainId,
	}
}
func (v *Voter) waitTxConfirmed(hash common.Hash) <-chan *types.Transaction {
	ch := make(chan *types.Transaction)
	go func() {
		for {
			tx, pending, _ := v.conn.TransactionByHash(context.TODO(), hash)
			if !pending {
				ch <- tx
			}

			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}

func (v *Voter) AddToQueue(hashId *big.Int, MintedBlockNumber uint64, currencyType *big.Int, version uint16) {
	mbn := big.NewInt(int64(MintedBlockNumber))
	v.queue = append(v.queue, votemanager.VoteManagerVotePayload{HashId: hashId, CurrencyType: currencyType, MintedBlockNumber: mbn, Version: version})

	v.mu.Lock()
	defer v.mu.Unlock()

	if len(v.queue) >= v.getBatchSize() {
		_ = v.Vote()
		v.queue = []votemanager.VoteManagerVotePayload{}
	}
}

func (v *Voter) Vote() error {
	log.Info("voting now!", "queue", len(v.queue))

	chainId := new(big.Int).SetUint64(v.chainId)
	auth, err := bind.NewKeyStoreTransactorWithChainID(v.ks, v.account, chainId)
	if err != nil {
		panic(err)
	}

	auth.GasTipCap, _ = v.conn.SuggestGasTipCap(context.TODO())
	auth.GasFeeCap, _ = v.conn.SuggestGasPrice(context.TODO())
	auth.GasLimit = GasLimit

	tx, err := v.vm.VoteBatch(auth, v.queue)
	if err != nil {
		log.Error("vote error", "err", err)
	}
	if tx == nil {
		log.Error("vote error", "err", "tx is nil")
	}

	v.waitTxConfirmed(tx.Hash())
	log.Info("vote success!", "tx", tx.Hash().Hex())

	return nil
}

func (v *Voter) getBatchSize() int {
	if v.Syncing {
		return BatchSize * 4
	}
	return BatchSize
}
