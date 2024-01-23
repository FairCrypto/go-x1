import jinja2
import ipaddress
import sys

ip_base = '172.16.239.0'
starting_port = 6050

if len(sys.argv) > 1:
    num_nodes = int(sys.argv[1])
else:
    num_nodes = 3

environment = jinja2.Environment()
environment.filters['ip_address'] = ipaddress.ip_address
environment.loader = jinja2.FileSystemLoader('templates')
template = environment.get_template('docker-compose.yaml.j2')

template.stream({
    "num_nodes": num_nodes,
    "ip_base": ip_base,
    "starting_port": starting_port,
}).dump('docker-compose.yaml')
