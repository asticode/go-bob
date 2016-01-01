#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/..

# Bind resources
$ROOT_DIR/scripts/build/bind_clean.sh
$ROOT_DIR/scripts/build/bind_execute.sh

# Build
$ROOT_DIR/scripts/build/core.sh

# Run
echo "Running..."
$GOPATH/bin/go-bob -config="$GOPATH/bin/bob.json"