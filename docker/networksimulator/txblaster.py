import sys
import signal

from web3 import Web3, HTTPProvider
from web3.middleware import construct_sign_and_send_raw_middleware
from time import time
import os
import argparse
import logging
import datetime
from multiprocessing import Process
from dotenv import load_dotenv

load_dotenv()

# This is the fake private key from the first validator in the network
# It is used to fund the accounts of the other validators
# This is only for testing purposes.
FAKE_VALIDATOR_PRIVATE_KEY = "0x163f5f0f9a621d72fedd85ffca3d08d131ab4e812181e0d30ffd1c885d20aac7"

PRIVATE_KEY = os.environ.get('PRIVATE_KEY') or FAKE_VALIDATOR_PRIVATE_KEY
RPC_URL = os.environ.get('RPC_URL') or "http://127.0.0.1:8545"
MNEMONIC = os.environ.get('MEMMONIC') or "witness explain monitor check grid depend music purchase ready title bar federal"

logging.basicConfig(level=logging.INFO, format='%(asctime)s %(levelname)s %(message)s')
logging.Formatter.formatTime = (lambda self, record, datefmt=None:
                                datetime.datetime.fromtimestamp(record.created, datetime.timezone.utc)
                                .astimezone()
                                .isoformat(sep="T", timespec="milliseconds"))


class TxBlaster:
    def __init__(self, rpc: str, private_key: str):
        self.web3 = Web3(HTTPProvider(rpc))
        self.web3.eth.account.enable_unaudited_hdwallet_features()
        self.funding_account = self.web3.eth.account.from_key(private_key)
        self.web3.middleware_onion.add(construct_sign_and_send_raw_middleware(self.funding_account))
        self.stop = False

    def fund_account(self, account_address: str, fund_value: int):

        balance = self.web3.eth.get_balance(account_address)
        diff = Web3.to_wei(fund_value, 'ether') - balance
        logging.info("funding from=%s to=%s balance=%d diff=%d", self.funding_account.address, account_address, balance, diff)

        if diff > 0:
            logging.info("funding account=%s", account_address)
            return self.web3.eth.send_transaction({
                "from": self.funding_account.address,
                "to": account_address,
                "value": diff,
                'gas': 21000,
            })
        else:
            logging.debug("account already funded: %s", account_address)

    def generate_accounts(self, num_accounts: int, menmonic: str, fund_value: int):
        tx_hashes = []
        accounts = []
        for i in range(num_accounts):
            acc = self.web3.eth.account.from_mnemonic(menmonic, account_path=f"m/44'/60'/0'/0/{i}")
            logging.info("generated account=%s fund_value=%s", acc.address, fund_value)
            tx_hash = self.fund_account(acc.address, fund_value)
            if tx_hash:
                tx_hashes.append(tx_hash)
            accounts.append(acc)
        for tx_hash in tx_hashes:
            self.web3.eth.wait_for_transaction_receipt(tx_hash, timeout=60000)
        return accounts

    def run_transactions(self, private_key, args):
        account = self.web3.eth.account.from_key(private_key)
        logging.info("running transactions account=%s txs=%d gas=%d value=%d",
                     account.address, args.num_txs, args.gas, args.value)

        self.web3.middleware_onion.add(construct_sign_and_send_raw_middleware(account))

        tx_params = {
            "from": account.address,
            "to": account.address,
            "value": Web3.to_wei(args.value, 'ether'),
            'gas': args.gas,
        }

        if args.nonce:
            tx_params["nonce"] = args.nonce

        if args.max_gas_fee and args.max_gas_priority_fee:
            tx_params["maxFeePerGas"] = Web3.to_wei(args.max_gas_fee, 'gwei')
            tx_params["maxPriorityFeePerGas"] = Web3.to_wei(args.max_gas_priority_fee, 'gwei')

        for _ in range(args.num_txs):
            if self.stop:
                logging.info("stopping")
                return

            try:
                tx_hash = self.web3.eth.send_transaction(tx_params)
            except Exception as e:
                logging.error("error sending transaction: %s", e)
                continue

            if args.no_wait:
                continue

            start_time = time()
            self.web3.eth.wait_for_transaction_receipt(tx_hash, timeout=60000)
            logging.info("tx confirmed tx_receipt=%s time=%f", tx_hash.hex(), time() - start_time)

    def sig_int(self, signum, frame):
        logging.info("Caught interrupt, stopping")
        self.stop = True


def thread_worker(private_key: str, args: argparse.Namespace):
    try:
        tx_blaster = TxBlaster(args.rpc, PRIVATE_KEY)
        tx_blaster.run_transactions(private_key, args)
    except KeyboardInterrupt:
        sys.exit(0)


def sig_int(signum, frame):
    logging.info("Caught interrupt, stopping workers")
    sys.exit(0)


if __name__ == "__main__":
    argparser = argparse.ArgumentParser("Run Send Transactions Tests", add_help=False)
    argparser.add_argument(dest="num_txs", nargs='?', type=int, const=1, default=100000, help="Number of transactions to send")
    argparser.add_argument('--no-wait', action="store_true", default=False, help="Don't wait for tx receipt")
    argparser.add_argument('--gas', type=int, default=21000)
    argparser.add_argument('--max-gas-fee', type=int, help="in gwei")
    argparser.add_argument('--max-gas-priority-fee', type=int, help="in gwei")
    argparser.add_argument('--value', type=int, default=.0001, help="in ether")
    argparser.add_argument('--help', '-h', action="store_true", help="show this help message and exit")
    argparser.add_argument('--rpc', type=str, default=RPC_URL, help="The RPC URL to use")
    argparser.add_argument('--nonce', type=int, required=False, help="The nonce")
    argparser.add_argument('--workers', type=int, default=1, help="The number of workers to use. Each worker will use a different account.")
    argparser.add_argument('--fund-value', type=int, default=1, help="The amount to fund each account.")
    args = argparser.parse_args()

    if args.help:
        argparser.print_help()
        exit(0)

    signal.signal(signal.SIGINT, sig_int)
    tx_blaster = TxBlaster(args.rpc, PRIVATE_KEY)
    accounts = tx_blaster.generate_accounts(args.workers, MNEMONIC, args.fund_value)

    processes = []
    for acc in accounts:
        logging.info("Starting worker: account=%s private_key=%s", acc.address, acc.key.hex())
        t = Process(target=thread_worker, args=(acc.key.hex(), args))
        t.start()
        processes.append(t)

