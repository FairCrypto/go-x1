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

## Docker


```shell
docker pull ghcr.io/faircrypto/go-x1:0.0.3-arm64
```
