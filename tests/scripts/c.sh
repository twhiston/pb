#!/usr/bin/env bash

set -e

echo "--- PB Test Procedure ---"

pushd ./c/macros
    echo "---> Run C Macro Tests"
    export CC=/usr/bin/clang
    cmake --target pbmacro_test -- -j 4 ./
    make
    ./pbmacro_test
popd