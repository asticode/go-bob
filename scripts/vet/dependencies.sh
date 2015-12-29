#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/../..

# Get dependencies
echo "Getting dependencies..."
go get -u github.com/golang/lint/golint
go install github.com/golang/lint/golint
go get golang.org/x/tools/cmd/goimports
go install golang.org/x/tools/cmd/goimports