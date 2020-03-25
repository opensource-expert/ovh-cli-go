#!/usr/bin/env bash
#
# This code is for bats-core test framework
#

OVH_CLI=../ovh-cli

@test "JSON_INPUT is read" {
  our_data="some data here"
  run $OVH_CLI --debug GET /me \""$our_data"\"

  # search our input in the collected lines
  found=0
  for i in $(seq 0 ${#lines})
  do
    if [[ ${lines[$i]} =~ $our_data ]] ; then
      found=1
      break
    fi
  done

  [[ $found -eq 1 ]]
}
