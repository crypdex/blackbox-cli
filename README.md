<img src="https://raw.githubusercontent.com/crypdex/blackbox/master/resources/images/logo.png" width=250>

# Blackbox CLI

A command line utility for interfacing with the devices running the **BlackboxOS**.

## Installation

Downloadable binaries for various platforms will be available soon. Until then, you will need to have Go installed and you can `go get` it.

```shell
$ go get github.com/crypdex/blackbox-cli
```

If you are using modules with Go, you will want to have Go version `>= 1.12` to avoid frustration with the above command.

## Usage

Once installed and available in your `PATH`, you may run the following command

```shell
$ blackbox
```

Which will print out the available commands, flags and usage instructions.

## Getting Started

For new devices, you may want to begin by checking status.

```shell
$ blackbox status --host crypdex-0000.local
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
