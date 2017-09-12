#!/usr/bin/env bash

set -e

echo "--- PB Test Procedure ---"

echo "---> Run Golang Tests"
/go/bin/goverage -v -coverprofile=coverage.out ./...
/go/bin/gometalinter --exclude \w*_test\w*
echo "---> Export Golang Test Coverage to Codacy"
godacov -t ${CODACY_TOKEN} -r ./coverage.out -c ${CI_COMMIT_ID}

pushd ./c/macros
    echo "---> Run C Macro Tests"
    export CC=/usr/bin/clang
    cmake --target pbmacro_test -- -j 4 ./
    make
    ./pbmacro_test
popd