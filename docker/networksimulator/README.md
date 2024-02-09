# Simulate a Network

## Install Dependencies

First, install [Docker](https://docs.docker.com/engine/install/) and [Docker Compose V2](https://docs.docker.com/compose/install/linux/).

> Install python Deps
```shell
# be in the docker/networksimulator directory
cd docker/networksimulator

pip3 install -r requirements.txt
```

## Run the Simulation

> run the simulated network. Replace NUM_OF_NODES with the number of nodes you want to simulate.

```shell
# Generate the docker compose file
# Replace NUM_OF_NODES with the number of nodes you want to simulate.
python3 generate.py NUM_OF_NODES

# Run the network
# use --build to rebuild the docker image from source ex: docker-compose up --build
docker-compose up

# Example: build and run a network of 10 nodes
python3 generate.py 10 && docker compose up --remove-orphans --build
```

## Interact with the Network

Give it a few minutes for all the nodes to start up.

> Test the RPC endpoint
```shell
curl -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}' http://127.0.0.1:8545
```

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
python3 generate.py 10 && docker compose up --remove-orphans --build
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
python3 generate.py 10 && docker compose up --remove-orphans --build
```

## Metrics and Monitoring

Prometheus and Granfana are included in the docker-compose file.

- Prometheus server at http://localhost:9090/targets
- Grafana dashboards at http://localhost:3000

## Test Script

The txblaster.py script is an example of how to interact with the simulated network. It uses the web3.py library to send transactions to the network.

> send 5 transactions to the network
```shell
python3 txblaster.py 5
```

> send 100 transactions and don't wait for a receipt
```shell
python3 txblaster.py 5 --no-wait
```

> send 1000 transactions from 10 workers (accounts)
```shell
python3 txblaster.py 1000 --workers 10
```

> see all the options
```shell
python3 txblaster.py--help
```

## Rolling out partial upgrades

First, build the new docker image with the changes you want to test.

```shell
docker-compose build
```

Then, you can upgrade the nodes one at a time.

```shell
docker-compose up -d --no-deps node2
docker-compose up -d --no-deps node3
docker-compose up -d --no-deps node4
```
