# My Ansible inventory

## Prerequisites

- Ansible
- Golang >= 1.22

## Setup Fresh Ubuntu 22 Server Playbook

1. Make your own `secrets.yml` with variables `server_username` and `server_password`

```sh
ansible-vault create secrets.yml
```

2. Run `servers.yml` playbook from root project folder

```sh
# New Ubuntu Server Basic Setup
ansible-playbook -i inventory.yml playbooks/servers.yml --ask-vault-pass
```

## Install Xray Playbook

1. Change `inventory.yml` to use newly created user with root priveleges in this playbook

```sh
ansible-playbook -i inventory.yml playbooks/xray.yml --ask-become-pass
```

## Configure IPv6 on Server

[Article about IPv6 on Habr](https://habr.com/ru/articles/811487/)
[Digital Ocean Docs For IPv6](https://docs.digitalocean.com/products/networking/ipv6/how-to/enable/#on-existing-droplets)

## Edit Secrets

```
ansible-vault edit secrets.yml
```

Cheat: password is mac password


# Scripts

## Build binary to create xray configs

```
cd automation
go build -o ../bin/automation
```

## Build xray config from template

1. Check that you create your `secrets.yml` according template in `secrets.template.yml`
2. Run script to create xray reality config from template in xray folder

```sh
./bin/automation build-config \
    -secrets ./secrets.yml \
    -xray-config-template ./xray/reality_config_simple.template.json \
    -xray-config-save ./xray/reality_config_simple.json
```