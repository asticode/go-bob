#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/../..

# Build
echo "Building..."
cd $GOPATH/bin
go build github.com/asticode/go-bob