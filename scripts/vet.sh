#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/..

# Clean binded resources
$ROOT_DIR/scripts/build/bind_clean.sh

# Vet
echo "Vetting..."
go fmt github.com/asticode/go-bob/...
goimports -w $ROOT_DIR
golint github.com/asticode/go-bob/...
go vet github.com/asticode/go-bob/...