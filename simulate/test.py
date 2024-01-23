from web3 import Web3, HTTPProvider
from web3.middleware import construct_sign_and_send_raw_middleware
from time import sleep, time

# Connect to local Ethereum node
w3 = Web3(HTTPProvider('http://localhost:8545'))

pk = ""  # add private key here
acct1 = w3.eth.account.from_key(pk)

w3.middleware_onion.add(construct_sign_and_send_raw_middleware(acct1))

while True:
    tx_hash = w3.eth.send_transaction({
        "from": acct1.address,
        "to": acct1.address,
        "value": 123123123123123,
        'gas': 21000,
    })

    start_time = time()
    while True:
        sleep(.1)
        tx = w3.eth.get_transaction(tx_hash)
        if tx.blockNumber is not None:
            break

    print(time() - start_time)

