package verifier

import (
	"context"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/votemanager"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sync"
)

var (
	BatchSize = 50
	gasLimit  = uint64(8000000)
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

	auth, err := NewTransactionOpts(context.Background(), *v.conn, v.chainId, v.ks, v.account, gasLimit)
	if err != nil {
		panic(err)
	}

	tx, err := v.vm.VoteBatch(auth, v.queue)
	if err != nil {
		log.Error("vote error", "err", err)
	}
	if tx == nil {
		log.Error("vote error", "err", "tx is nil")
	}

	WaitTxConfirmed(context.Background(), *v.conn, tx.Hash())
	log.Info("vote success!", "tx", tx.Hash().Hex())

	return nil
}

func (v *Voter) getBatchSize() int {
	if v.Syncing {
		return BatchSize * 4
	}
	return BatchSize
}
