#!/bin/sh

set -ex
go get github.com/alecthomas/gometalinter
gometalinter --install --update
