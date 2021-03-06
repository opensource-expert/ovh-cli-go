#
# Makefile for managing ovh-cli-go build
#

# dependancies
GOVVV=${GOPATH}/bin/govvv

#= = = = = = = = = = = = = = = = = = = =  keep this as first target for development
# build 64 bits version
# govvv define main.Version with the contents of ./VERSION file, if exists
BUILD_FLAGS=$(shell ./get_ldflags.sh)
ovh-cli: ovh-cli.go Makefile ${GOVVV} VERSION
	go build -o $@ -ldflags "${BUILD_FLAGS} ${LDFLAGS}"

${GOVVV}:
	go get github.com/ahmetb/govvv

test: ovh-cli
	./ovh-cli --version
	cd tests/ && ./bats/bin/bats .

# can be overrided:
# PREFIX=~/.local make install
PREFIX ?= /usr/local
install: ovh-cli
	install -m 0755 $< ${PREFIX}/bin

clean:
	rm -f ovh-cli-go ovh-cli build/*
