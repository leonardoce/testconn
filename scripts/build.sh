#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit
PATH=$HOME/go/bin:$PATH go generate
CGO_ENABLED=0 go build -o bin/testconn main.go
