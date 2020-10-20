#!/usr/bin/env bash
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#

set -e

THIS_DIR="""$(dirname "$0")/.."""

for f in $(find . -name "main.go"); do
    directory="""$(dirname "$f")"""
    (
        echo "Testing $f"
        cd "${directory}" || exit 1
        go test -v -i main.go || {
            echo " ";echo " "
            echo "Test failed in $f"
            echo " ";echo " "
            exit 2
        }
    )
done