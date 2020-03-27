# ovh-cli functionnal testing


## Requirement

- tested with [bats-core](https://github.com/bats-core/bats-core)
- a valid `~/.ovh.conf` enviroment, with valid credential (`ovh-cli GET /auth/currentCredential`)
- a network access to OVH API

## run tests

```
cd ./tests
bats suite.bats
```
