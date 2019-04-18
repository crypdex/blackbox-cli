<img src="https://raw.githubusercontent.com/crypdex/blackbox-cli/master/resources/images/logo.png" width=250>

# Blackbox CLI

A command line utility for interfacing with the devices running the **[BlackboxOS](https://github.com/crypdex/blackbox)**. This includes the [PIVX Staking Node](https://crypdex.io/products/pivx-staking-node) and the [Multichain Staking Node](https://crypdex.io/products/multichain-staking-node)

## Installation

Download the most recent [release](https://github.com/crypdex/blackbox-cli/releases) for your platform.
## Installation from `apt`

Add the following to `/etc/apt/sources.list`
```
deb [trusted=yes] https://apt.fury.io/crypdex/ /
```
And install normally
```shell
$ apt update && apt install
```

## Usage

Once installed and available in your `PATH`, you may run the following command

```shell
$ blackbox-cli
```

Which will print out the available commands, flags and usage instructions.

## Getting Started

For new devices, you may want to begin by checking status. By default, it will look for `crypdex.local`.

```shell
$ blackbox-cli status
```

### Initialization

Initialize your wallet with the following command.

```shell
$ blackbox-cli init
```


