from web3 import Web3, HTTPProvider
from web3.middleware import construct_sign_and_send_raw_middleware
from time import time
import os
import argparse

PRIVATE_KEY = os.environ.get('PRIVATE_KEY')  # add private key here
RPC_URL = os.environ.get('RPC_URL')  # add rpc url here

# Connect to local Ethereum node
w3 = Web3(HTTPProvider(RPC_URL))

argparser = argparse.ArgumentParser("Run Send Transaction Tests", add_help=False)
argparser.add_argument(dest="num_txs", nargs='?', type=int, const=1, default=1000, help="Number of transactions to send")
argparser.add_argument('--no-wait', action="store_true", default=False, help="Don't wait for tx receipt")
argparser.add_argument('--gas', type=int, default=21000)
argparser.add_argument('--max-gas-fee', type=int, help="in gwei")
argparser.add_argument('--max-gas-priority-fee', type=int, help="in gwei")
argparser.add_argument('--value', type=int, default=1, help="in ether")
argparser.add_argument('--help', '-h', action="store_true", help="show this help message and exit")

args = argparser.parse_args()

if args.help:
    argparser.print_help()
    exit(0)

if not PRIVATE_KEY:
    raise ValueError("Missing PRIVATE_KEY environment variable")


acct1 = w3.eth.account.from_key(PRIVATE_KEY)
w3.middleware_onion.add(construct_sign_and_send_raw_middleware(acct1))

params = {
    "from": acct1.address,
    "to": acct1.address,
    "value": Web3.to_wei(args.value, 'ether'),
    'gas': args.gas
}

if args.max_gas_fee and args.max_gas_priority_fee:
    params["maxFeePerGas"] = Web3.to_wei(args.max_gas_fee, 'gwei')
    params["maxPriorityFeePerGas"] = Web3.to_wei(args.max_gas_priority_fee, 'gwei')

for i in range(args.num_txs):
    tx_hash = w3.eth.send_transaction(params)

    if args.no_wait:
        continue

    start_time = time()
    tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    print(time() - start_time)
