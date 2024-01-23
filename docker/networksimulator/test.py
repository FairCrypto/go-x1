from web3 import Web3, HTTPProvider
from web3.middleware import construct_sign_and_send_raw_middleware
from time import time
import os

# Connect to local Ethereum node
w3 = Web3(HTTPProvider('http://localhost:8545'))

pk = os.environ.get('PRIVATE_KEY')  # add private key here
acct1 = w3.eth.account.from_key(pk)

w3.middleware_onion.add(construct_sign_and_send_raw_middleware(acct1))

while True:
    tx_hash = w3.eth.send_transaction({
        "from": acct1.address,
        "to": acct1.address,
        "value": Web3.to_wei(1, 'ether'),
        'gas': 9000000,
        # "maxFeePerGas": Web3.to_wei(240, 'gwei'),
        # "maxPriorityFeePerGas": Web3.to_wei(120, 'gwei'),
    })

    start_time = time()
    tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    print(time() - start_time)
