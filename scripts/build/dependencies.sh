#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/../..

# Get dependencies
echo "Getting dependencies..."
go get github.com/asticode/go-bob/...

# Setting up go-bindata
echo "Setting up go-bindata..."
go get -u github.com/jteeuwen/go-bindata/...
go install github.com/jteeuwen/go-bindata/...