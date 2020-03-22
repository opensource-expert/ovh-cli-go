# Simple ovh-cli client for shell

Here'is a basic ovh API cli. No fancy argument. You can copy paste from [OVH console](https://eu.api.ovh.com/console).

It is compatible with [`jq`](https://stedolan.github.io/jq/) JSON parser.

It is based on official OVH [github.com/ovh/go-ovh/ovh](https://github.com/ovh/go-ovh).


## Installation

You can download precompiled binaries in [releases](https://github.com/opensource-expert/ovh-cli-go/releases).

And put it in `/usr/local/bin`

## get the code

```
go get github.com/opensource-expert/ovh-cli-go
```

## Configuration

It uses default `~/.ovh.conf` mecanism provided by [go-ovh](https://github.com/ovh/go-ovh#configuration)

## Usage

```
./ovh-cli -h
Usage:
  ovh-cli [--debug] METHOD URL_API [JSON_INPUT]
```

## Examples

```
./ovh-cli GET /me
```

pretty print

```
./ovh-cli GET /me | jq .
```

## Have fun

`:-)`
