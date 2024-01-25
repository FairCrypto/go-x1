import jinja2
import ipaddress
import argparse

DEFAULT_IP_BASE = '172.16.239.0'
DEFAULT_NUM_NODES = 3
DEFAULT_STARTING_PORT = 7050
DEFAULT_FLAGS = "--testnet --metrics --metrics.addr 0.0.0.0 --http --http.port 8545 --http.addr 0.0.0.0 --http.vhosts '*' --http.corsdomain '*' --ws --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins '*' --verbosity 3 "
BOOT_NODE_KEY = "4b20a091f6389ca9ee1492187cc2d775511fa1e0801bf1f787b3a14f961530b1"
BOOT_NODE_PUB_KEY = "d06482f636e5c68586215f9ab9dfda270d38bf468195fc2e767d5d74b5fc7ab4faffc46028aa360b723ce53ded022949a8a6b1c96013d8ec1771f4ed448518b4"


argparser = argparse.ArgumentParser()
argparser.add_argument(type=int, dest="num_nodes", default=DEFAULT_NUM_NODES)
argparser.add_argument('--ip-base', type=str, default=DEFAULT_IP_BASE)
argparser.add_argument('--flags', type=str, default=DEFAULT_FLAGS)
argparser.add_argument('--starting-port', type=int, default=DEFAULT_STARTING_PORT)
args = argparser.parse_args()

environment = jinja2.Environment()
environment.filters['ip_address'] = ipaddress.ip_address
environment.loader = jinja2.FileSystemLoader('templates')
template = environment.get_template('docker-compose.yaml.j2')

template.stream({
    "num_nodes": args.num_nodes,
    "ip_base": args.ip_base,
    "starting_port": args.starting_port,
    "flags": args.flags,
    "bootnode_key": BOOT_NODE_KEY,
    "bootnode_pub_key": BOOT_NODE_PUB_KEY,
}).dump('docker-compose.yaml')
