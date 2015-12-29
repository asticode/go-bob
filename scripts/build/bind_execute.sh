#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/../..

# Bind resources
echo "Binding resources..."
cd $ROOT_DIR
go-bindata -o resources.go resources/...