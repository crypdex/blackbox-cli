<img src="https://raw.githubusercontent.com/crypdex/blackbox-cli/master/resources/images/logo.png" width=250>

# Blackbox CLI

A command line utility for interfacing with the devices running the **BlackboxOS**.

## Installation

Download the most recent [release](https://github.com/crypdex/blackbox-cli/releases) for your platform. If you are running Linux, you may install via `apt`.

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

## Configuration

There are some variables that you may use consistently across commands like `--chain` and `--host`. This is especially true if you are running a single chain. These variables can be saved in a config file, formatted in YAML and located at `~/.crypdex/blackbox.yaml` that will be used across commands. Flags override defaults.

You may create this file manually if you'd like or use the following command:

```shell
$ blackbox config set --chain pivx --host crypdex-0000.local
```

Interested in what's in the config?

```shell
$ blackbox config get
```

... or remove it altogether

```shell
$ blackbox config rm
```
