#!/bin/sh

set -ex
gometalinter --exclude=vendor --exclude=repos --disable-all --enable=golint --enable=vet --enable=gofmt ./...
find . -name '*.go' | xargs gofmt -w -s
