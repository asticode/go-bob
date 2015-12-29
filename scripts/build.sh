#!/usr/bin/env bash

BASH_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ROOT_DIR=$BASH_DIR/..

# Get dependencies
$ROOT_DIR/scripts/build/dependencies.sh

# Bind resources
$ROOT_DIR/scripts/build/bind_clean.sh
$ROOT_DIR/scripts/build/bind_execute.sh

# Core
$ROOT_DIR/scripts/build/core.sh