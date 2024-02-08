import logging

import jinja2
import ipaddress
import argparse
import socket

DEFAULT_IP_BASE = '172.232.0.0'
DEFAULT_NUM_NODES = 3
DEFAULT_STARTING_PORT = 17050
DEFAULT_PROMETHEUS_PORT = 9090
DEFAULT_GRAFANA_PORT = 3000
DEFAULT_RPC_PORT = 8545
DEFAULT_WS_PORT = 8546
DEFAULT_FLAGS = "--testnet --syncmode full --gcmode archive --metrics --metrics.addr 0.0.0.0 --http --http.port 8545 --http.addr 0.0.0.0 --http.vhosts '*' --http.corsdomain '*' --ws --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins '*' --verbosity 3 "
BOOT_NODE_KEY = "4b20a091f6389ca9ee1492187cc2d775511fa1e0801bf1f787b3a14f961530b1"
BOOT_NODE_PUB_KEY = "d06482f636e5c68586215f9ab9dfda270d38bf468195fc2e767d5d74b5fc7ab4faffc46028aa360b723ce53ded022949a8a6b1c96013d8ec1771f4ed448518b4"
DEFAULT_NODE_EXPORTER_PORT = 9100


def is_port_in_use(port: int) -> bool:
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        return s.connect_ex(('127.0.0.1', port)) == 0


def get_next_open_port(starting_port: int) -> int:
    port = starting_port
    while is_port_in_use(port):
        logging.info(f"Port {port} is in use, trying next port")
        port += 1
    return port


argparser = argparse.ArgumentParser()
argparser.add_argument(type=int, dest="num_nodes", default=DEFAULT_NUM_NODES)
argparser.add_argument('--ip-base', type=str, default=DEFAULT_IP_BASE)
argparser.add_argument('--flags', type=str, default=DEFAULT_FLAGS)
argparser.add_argument('--starting-port', type=int, default=DEFAULT_STARTING_PORT)
argparser.add_argument('--prometheus-port', type=int, default=DEFAULT_PROMETHEUS_PORT)
argparser.add_argument('--grafana-port', type=int, default=DEFAULT_GRAFANA_PORT)
argparser.add_argument('--rpc-port', type=str, default=DEFAULT_RPC_PORT)
argparser.add_argument('--ws-port', type=str, default=DEFAULT_WS_PORT)
argparser.add_argument('--node-exporter-port', type=str, default=DEFAULT_NODE_EXPORTER_PORT)
args = argparser.parse_args()

prometheus_ip = ipaddress.ip_address(args.ip_base) + 2
grafana_ip = ipaddress.ip_address(args.ip_base) + 3
node_exporter_ip = ipaddress.ip_address(args.ip_base) + 4

grafana_port = get_next_open_port(args.grafana_port)
prometheus_port = get_next_open_port(args.prometheus_port,)
main_rpc_port = get_next_open_port(args.rpc_port)
main_ws_port = get_next_open_port(args.ws_port)
node_exporter_port = get_next_open_port(args.node_exporter_port)

print(f"""Generating a {args.num_nodes} node X1 network

RPC: http://127.0.0.1:{main_rpc_port}
WS: http://127.0.0.1:{main_ws_port}/ws
Grafana: http://127.0.0.1:{grafana_port}
Prometheus: http://127.0.0.1:{prometheus_port}
""")

nodes = []
for i in range(1, args.num_nodes + 1):
    rpc_port = get_next_open_port(args.starting_port + i + 2000)
    ws_port = get_next_open_port(args.starting_port + i + 3000)
    metrics_port = get_next_open_port(args.starting_port + i + 1000)
    p2p_port = get_next_open_port(args.starting_port + i)

    nodes.append({
        "id": i,
        "ip": ipaddress.ip_address(args.ip_base) + i + 4,
        "p2p_port": p2p_port,
        "metrics_port": metrics_port,
        "rpc_port": rpc_port,
        "ws_port": ws_port,
    })

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
    "nodes": nodes,
    "prometheus_ip": prometheus_ip,
    "grafana_ip": grafana_ip,
    "grafana_port": grafana_port,
    "prometheus_port": prometheus_port,
    "rpc_port": main_rpc_port,
    "ws_port": main_ws_port,
    "node_exporter_ip": node_exporter_ip,
    "node_exporter_port": node_exporter_port,
}).dump('docker-compose.yaml')

prometheus = environment.get_template('prometheus.yml.j2')
prometheus.stream({
    "nodes": nodes,
    "node_exporter_ip": node_exporter_ip,
    "node_exporter_port": node_exporter_port,
}).dump('prometheus/prometheus.yml')
