#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/..

# Vet
$ROOT_DIR/scripts/vet.sh

# Bind resources
$ROOT_DIR/scripts/build/bind_execute.sh

# Test
echo "Testing..."
cd $ROOT_DIR
go test -coverprofile=scripts/test/coverage/main.out github.com/asticode/go-bob
go tool cover -html=scripts/test/coverage/main.out -o scripts/test/coverage/main.html