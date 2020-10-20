#!/usr/bin/env bash
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#

set -e

echo "installing git hooks"
rm -f .git/hooks/*
for f in $(find githooks -type f); do
    filename="""$(basename "$f")"""
    ln -sf "./githooks/${filename}" ".git/hooks/${filename}"
done