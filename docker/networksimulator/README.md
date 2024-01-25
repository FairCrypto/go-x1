# Simulate a Network

> install deps

```shell
pip3 install -r requirements.txt
```

> run the simulated network. Replace NUM_OF_NODES with the number of nodes you want to simulate.

```shell
# be in the docker/networksimulator directory
cd docker/networksimulator

# install deps
pip3 install -r requirements.txt

# Generate the docker compose file
# Replace NUM_OF_NODES with the number of nodes you want to simulate.
python3 generate.py NUM_OF_NODES

# Run the network
# use --build to rebuild the docker image from source ex: docker-compose up --build
docker-compose up

# Example: build and run a network of 10 nodes
python3 generate.py 10 && docker compose up --build
```

Give it a few minutes for all the nodes to start up.

> Test the RPC endpoint
```shell
curl -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}' http://127.0.0.1:8545
```

Use the fakenet private keys from [../evmcore/apply_fake_genesis.go](../evmcore/apply_fake_genesis.go) to access the validator accounts and their funds

See test.py for an example of how to interact with the network.

## Lachesis Base

This is a hacky way to include a custom Lachesis base inside the docker build context.

> Rsync your lachesis base files into the `lachesis-base` directory inside go-x1. This will be copied into the docker build context.
```shell
rsync -rav --exclude .git ../lachesis-base/ lachesis-base/
```
 
> Update the go.mod file to point to the local lachesis-base directory
```shell
echo "replace github.com/Fantom-foundation/lachesis-base => ./lachesis-base" >> go.mod
```

> Now build the docker image as usual
```shell
cd docker/networksimulator

# Example: build and run a network of 10 nodes
python3 generate.py 10 && docker compose up --build
```

## Include a Custom SFC contract

The custom SFC contract will be built and included in the genesis block of the simulated network. 

Note: The genesis will no longer be compatible with the testnet genesis.

```shell
# clone the sfc contract repo into the parent directory of go-x1
# so both repos are in the same directory
cd ../../..
git clone git@github.com:nibty/x1-sfc.git

cd x1-sfc
# make any changes to the SFC contract here

# run go generate ./... inside the go-x1 directory
# This will build the SFC contract and generate the go bindings.
# This requires Docker, takes a long time, and will not exit with ctrl-c, so be careful.
cd ../go-x1
go generate ./...
```

> Now build the docker image as usual
```shell
cd docker/networksimulator

# Example: build and run a network of 10 nodes
python3 generate.py 10 && docker compose up --build
```
