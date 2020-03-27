#!/usr/bin/env bash
#
# This code is for bats-core test framework
#

# OVH_CLI must point to a binary to test (See: suite.bats)
# Usage:
#   export OVH_CLI=../ovh-cli
#   bats ovh-cli.bats

@test "JSON_INPUT is read" {
  our_data="some data here"
  [[ -x $OVH_CLI ]]
  run $OVH_CLI --debug GET /me "\"$our_data\""

  grep -F "$our_data" <<< "$output"
}

@test "INPUT is read from stdin FILE" {
  our_data="some data from stdin here"
  [[ -x $OVH_CLI ]]
  run $OVH_CLI --debug GET /me <<< "\"$our_data\""

  grep -F "$our_data" <<< "$output"
}

@test "INPUT is read from stdin PIPE" {
  our_data="some data from stdin here"
  [[ -x $OVH_CLI ]]
  echo "\"$our_data\"" | {
    run $OVH_CLI --debug GET /me
    grep -F "$our_data" <<< "$output"
  }
}
