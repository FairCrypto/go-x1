#!/bin/bash

cat ../xenblock-tokens/artifacts/contracts/VoteManager.sol/VoteManager.json | jq '.abi' > VoteManager.abi.json
go run github.com/ethereum/go-ethereum/cmd/abigen --abi VoteManager.abi.json --pkg votemanager > integration/xenblocks/contracts/votemanager/votemanager.go
rm VoteManager.abi.json

cat ../xenblock-tokens/artifacts/contracts/TokenRegistry.sol/TokenRegistry.json | jq '.abi' > TokenRegistry.abi.json
go run github.com/ethereum/go-ethereum/cmd/abigen --abi TokenRegistry.abi.json --pkg tokenregistry > integration/xenblocks/contracts/tokenregistry/tokenregistry.go
rm TokenRegistry.abi.json

cat ../x1-block-storage/build/contracts/BlockStorage.json | jq '.abi' > BlockStorage.abi.json
go run github.com/ethereum/go-ethereum/cmd/abigen --abi BlockStorage.abi.json --pkg block_storage > integration/xenblocks/contracts/block_storage/block_storage.go
rm BlockStorage.abi.json
