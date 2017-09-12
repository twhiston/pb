#!/usr/bin/env bash

set -e

echo "---> Run Golang Tests"
/go/bin/goverage -v -coverprofile=coverage.out ./...
/go/bin/gometalinter --exclude \w*_test\w*
echo "---> Export Golang Test Coverage to Codacy"
godacov -t ${CODACY_TOKEN} -r ./coverage.out -c ${CI_COMMIT_ID}