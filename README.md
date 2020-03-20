# Simple ovh-cli client for shell

Here'is a basic ovh API cli. No fancy argument. You can copy paste from [OVH console](https://eu.api.ovh.com/console).

It is compatible with [`jq`](https://stedolan.github.io/jq/) JSON parser.

It is based on official OVH [github.com/ovh/go-ovh/ovh](https://github.com/ovh/go-ovh).


## Installation

```
go get URL
```

or download precompiled binaries in release.

## Configuration

It use default `~/.ovh.conf` mecanism provided by [go-ovh](https://github.com/ovh/go-ovh#configuration)

## Usage

```
./ovh-cli-go -h
Usage:
  ovh-cli-go [--debug] METHOD URL_API [JSON_INPUT]
```

## Examples

```
./ovh-cli-go GET /me
```

pretty print

```
./ovh-cli-go GET /me | jq .
```
