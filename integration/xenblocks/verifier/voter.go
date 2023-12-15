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
	"time"
)

var (
	voteManagerAddr = common.HexToAddress("0x84B3519B57E8017324F082d2e0F0C95051687aE2")
	BATCH_SIZE      = 50
	GAS_LIMIT       = uint64(8000000)
)

type Voter struct {
	vm      *votemanager.Votemanager
	conn    *ethclient.Client
	queue   []votemanager.VoteManagerPayload
	ks      *keystore.KeyStore
	account accounts.Account
	chainId uint64
}

func NewVoter(conn *ethclient.Client, ks *keystore.KeyStore, account accounts.Account, chainId uint64) *Voter {
	vm, err := votemanager.NewVotemanager(voteManagerAddr, conn)
	if err != nil {
		panic(err)
	}

	var queue []votemanager.VoteManagerPayload

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

func (v *Voter) AddToQueue(hashId *big.Int, currencyType *big.Int) {
	v.queue = append(v.queue, votemanager.VoteManagerPayload{HashId: hashId, CurrencyType: currencyType})

	if len(v.queue) >= BATCH_SIZE {
		_ = v.Vote()
		v.queue = []votemanager.VoteManagerPayload{}
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
	auth.GasLimit = GAS_LIMIT

	tx, err := v.vm.VoteBatch(auth, v.queue)
	if err != nil {
		log.Error("vote error", "err", err)
	}

	v.waitTxConfirmed(tx.Hash())
	log.Info("vote success!", "tx", tx.Hash().Hex())

	return nil
}
