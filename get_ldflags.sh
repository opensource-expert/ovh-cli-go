#!/usr/bin/env bash
#
# Usage: ./get_ldflags.sh [BUILD_FLAGS]

set -eu
build_flags=${1:-}
if [[ -z $build_flags ]]; then
  # govvv define main.Version with the contents of ./VERSION file, if exists
  build_flags=$(govvv -flags)
fi

# add more flags
build_flags+=" -X 'main.GoBuildVersion=$(go version)' -X 'main.ByUser=${USER}'"
echo "$build_flags"
