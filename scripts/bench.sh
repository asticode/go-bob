#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/..

# Bind resources
$ROOT_DIR/scripts/build/bind_clean.sh
$ROOT_DIR/scripts/build/bind_execute.sh

# Benchmark
echo "Benchmarking..."
cd $ROOT_DIR
go test -bench=.