#!/usr/bin/env bash
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#

set -e

echo "Pre-commit starting..."

THIS_DIR="""$(dirname "$0")/.."""

make lint