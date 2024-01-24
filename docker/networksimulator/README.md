# Simulate a Network

> install deps

```shell
pip3 install -r requirements.txt
```

> run the simulated network. Replace NUM_OF_NODES with the number of nodes you want to simulate.

```shell
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

```shell
curl -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}' http://127.0.0.1:8545
```

Use the fakenet private keys from [../evmcore/apply_fake_genesis.go](../evmcore/apply_fake_genesis.go) to access the validator accounts and their funds

See test.py for an example of how to interact with the network.
