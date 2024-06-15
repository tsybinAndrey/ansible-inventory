# My Ansible inventory

## Prerequisites

- Ansible

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

## Edit Secrets

```
ansible-vault edit secrets.yml
```

Cheat: password is mac password
