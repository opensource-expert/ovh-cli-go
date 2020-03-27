#!/usr/bin/env bash
#
# our test suite

###################################################################### helpers

_check_suite_tap_output() {
  # we only count non ok statment
  local nb_err=$(awk 'BEGIN { r = 0; }
    NR >= 2 {if($1 != "ok") { r++ } }
    END { print r}')

  return $nb_err
}

_run_suite() {
  run bats -t ovh-cli.bats
  echo "$output"
  echo "status $status"
  [ $status -eq 0 ]
  echo "$output" |  _check_suite_tap_output
}

###################################################################### tests

@test "running on compiled ovh-cli (Makefile)" {
  export OVH_CLI=../ovh-cli
  _run_suite
}

@test "running on build/ linux_amd64 (./deploy.sh build)" {
  OVH_CLI=../build/ovh-cli_linux_amd64
  _run_suite
}
