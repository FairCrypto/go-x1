package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/tokenregistry"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/votemanager"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/verifier"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type payload struct {
	decoded decoded
	bn      uint64
}

type decoded struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}

var (
	jobs chan payload
	v    *verifier.Verifier
)

func main() {
	ctx := context.TODO()

	conn, _ := ethclient.Dial("http://localhost:8545")
	bs, _ := block_storage.NewBlockStorage(common.HexToAddress(verifier.BlockStorageAddr), conn)
	tr, _ := tokenregistry.NewTokenregistry(common.HexToAddress(verifier.TokenRegistryAddr), conn)
	vm, _ := votemanager.NewVotemanager(common.HexToAddress(verifier.VoteManagerAddr), conn)
	bn, _ := conn.BlockNumber(ctx)
	v = verifier.NewVerifier(0, conn, bs, vm, tr)

	startBlock := 832437
	toBlock := startBlock + 100

	jobs = make(chan payload, 1000)
	for w := 0; w <= 10; w++ {
		go worker()
	}

	for {
		fmt.Printf("%d => %d\n", startBlock, toBlock)
		blockNumber := new(big.Int).SetUint64(bn)
		if blockNumber.Cmp(big.NewInt(int64(startBlock))) == -1 {
			fmt.Println("caught up with the latest block")
			break
		}

		logTransferSig := []byte("NewHash(uint256,uint256,address,uint256,bytes)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

		q := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: []common.Address{
				common.HexToAddress("0xf7E0CF7453ac619fD64b3D46D7De3638510F15eA"),
			},
			Topics: [][]common.Hash{
				{logTransferSigHash},
			},
		}

		logs, err := conn.FilterLogs(context.Background(), q)
		if err != nil {
			panic(err)
		}

		for _, vLog := range logs {
			event, err := bs.ParseNewHash(vLog)
			if err != nil {
				panic(err)
			}
			dr, err := bs.DecodeRecordBytes0(nil, event.HashId)
			if err != nil {
				panic(err)
			}

			jobs <- payload{
				decoded: dr,
				bn:      vLog.BlockNumber,
			}
		}

		startBlock = toBlock + 1
		toBlock += 100
	}
}

func worker() {
	for pl := range jobs {
		event := pl.decoded

		key := []byte(common.Bytes2Hex(event.K[:]))
		argon2Result, err := verifier.Argon2Hash(event.C, event.M, event.T, event.S, key)
		if err != nil {
			panic(err)
		}

		tokens, _ := v.Tr.FindTokensFromArgon2Hash(nil, argon2Result, big.NewInt(0))
		j, _ := json.Marshal(tokens)
		log.Printf("%d %s %s\n", pl.bn, argon2Result, j)
	}
}
