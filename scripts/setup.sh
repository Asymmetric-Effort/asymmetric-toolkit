#!/usr/bin/env bash
#(c) 2018 Sam Caldwell.  See LICENSE.txt.
#

set -e

brew install git-hooks || brew upgrade git-hooks
brew install shellcheck || brew upgrade shellcheck
brew install flake8 || brew upgrade flake8
brew install yamllint || brew upgrade yamllint
brew install jsonlint || brew upgrade jsonlint || brew reinstall
brew install golangci/tap/golangci-lint || brew upgrade golangci/tap/golangci-lint
