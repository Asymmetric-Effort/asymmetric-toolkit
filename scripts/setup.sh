#!/usr/bin/env bash
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#

set -e

echo "installing git hooks"
(
    mkdir -p $GOPATH/src/github.com/
    cd $GOPATH/src/github.com/
    [[ -d git-hooks ]] && rm -rf git-hooks
    git clone git@github.com:sam-caldwell/git-hooks.git
    cd git-hooks
    make compress
    go install
)
echo "git hooks installed."
git hooks install
git hooks
mkdir githooks

