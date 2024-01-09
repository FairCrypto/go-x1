package verifier

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func NewTransactionOpts(context context.Context, conn ethclient.Client, chainId uint64, ks *keystore.KeyStore, account accounts.Account, gasLimit uint64) (*bind.TransactOpts, error) {
	ci := new(big.Int).SetUint64(chainId)
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, ci)
	if err != nil {
		return nil, err
	}

	auth.GasTipCap, _ = conn.SuggestGasTipCap(context)
	auth.GasFeeCap, _ = conn.SuggestGasPrice(context)
	auth.GasLimit = gasLimit

	return auth, nil
}

func WaitTxConfirmed(context context.Context, conn ethclient.Client, hash common.Hash) <-chan *types.Transaction {
	ch := make(chan *types.Transaction)
	go func() {
		for {
			tx, pending, _ := conn.TransactionByHash(context, hash)
			if !pending {
				ch <- tx
			}

			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}
