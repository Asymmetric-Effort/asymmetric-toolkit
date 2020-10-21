#!/usr/bin/env bash
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#

set -e

echo "$0 starting"

#
# Lint Shell code
#

# shellcheck disable=SC2044
for f in $(find . -name "*.sh"); do
    (
        echo "Testing $f"
        shellcheck "${f}" || {
            echo " "
            echo " "
            echo "Linter failed in $f"
            echo " "
            echo " "
            exit 2
        }
    )
done

#
# Lint python code
#
# shellcheck disable=SC2044
for f in $(find . -name "*.py"); do
    (
        echo "Testing $f"
        flake8 "${f}" || {
            echo " "
            echo " "
            echo "Linter failed in $f"
            echo " "
            echo " "
            exit 2
        }
    )
done

#
# Lint golang code
#
# shellcheck disable=SC2044
golangci-lint run --enable-all --disable gomodguard --disable bodyclose \
                  --disable funlen --disable wsl --disable golint --disable stylecheck \
                  --disable nlreturn --disable gocognit --disable gofumpt

#
# Lint yaml code
#
# shellcheck disable=SC2044
for f in $(find . -name "*.y?ml"); do
    (
        echo "Testing $f"
        yamllint "${f}" || {
            echo " "
            echo " "
            echo "Linter failed in $f"
            echo " "
            echo " "
            exit 2
        }
    )
done


#
# Lint json code
#
# shellcheck disable=SC2044
for f in $(find . -name "*.json"); do
    (
        echo "Testing $f"
        jsonlint "${f}" || {
            echo " "
            echo " "
            echo "Linter failed in $f"
            echo " "
            echo " "
            exit 2
        }
    )
done