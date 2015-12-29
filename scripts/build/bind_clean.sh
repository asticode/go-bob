#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/../..

# Clean binded resources
echo "Cleaning binded resources..."
rm -f $ROOT_DIR/resources.go