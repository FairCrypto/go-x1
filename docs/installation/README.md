# Installation

## System Packages

### Install

#### Ubuntu

Download the latest deb release from [here](https://github.com/FairCrypto/go-x1/releases).

> Install the deb package
```shell
sudo dpkg -i go-x1_x.x.x_amd64.deb
```

#### Redhat/Centos
Download the latest rpm release from [here](https://github.com/FairCrypto/go-x1/releases).

> Install the rpm package
```shell
sudo yum install go-x1_x.x.x_amd64.rpm
```

### Config

Modify the configuration file `/etc/default/x1` to your needs.

By default, the service will run as the user `x1` and group `x1` with the home directory `/var/lib/x1`.

> You may change this to your needs by running the following commands:
```shell
sudo systemctl edit x1.service
```

> Add the following lines
```unit file (systemd)
[Service]
User=<user>
Group=<group>
```

### Start

> Start the service
```shell
sudo systemctl start x1.service
```

### Logs

> Tail the logs
```shell
sudo journalctl -t x1 -f
```

> See the last 1000 lines of logs
```shell
sudo journalctl -t x1 -n 1000
```

## Binaries

Download the latest binary release from [here](https://github.com/FairCrypto/go-x1/releases).

> Extract and install the binary
```shell
tar xvf x1_x.x.x_linux_amd64.tar.gz
sudo cp x1_x.x.x_linux_amd64/x1 /usr/local/bin
sudo chmod +x /usr/local/bin/x1
```

## Docker

See the Docker packages [here](https://github.com/nibty/faircrypto/pkgs/container/go-x1).

> Pull the latest image
```shell
docker pull ghcr.io/faircrypto/go-x1:latest
```

> Run the container
```shell
mkdir -p $HOME/.x1
docker run -d --name x1 -p 5050:5050 -v $HOME/.x1:/root/.x1 ghcr.io/faircrypto/go-x1:latest --testnet
```

> See the logs
```shell
docker logs -f x1
```

> Stop the container
```shell
docker stop x1
```

## From Source

```shell
# Install dependencies (ex: ubuntu)
apt update -y
apt install -y golang wget git make

# Clone and build the X1 binary
git clone --branch x1 https://github.com/FairCrypto/go-x1
cd go-x1
make x1
cp build/x1 /usr/local/bin
```
