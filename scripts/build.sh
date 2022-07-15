#!/bin/bash

d=$(cd "$(dirname "$0")"; pwd)
pushd $d/..

dist=dist
rm -rf ${dist} && true
mkdir ${dist} && true

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${dist}/tunaps -v -ldflags "-w -s" cmd/tunaps/tunaps.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${dist}/tunapc -v -ldflags "-w -s" cmd/tunapc/tunapc.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${dist}/tunaps.exe -v -ldflags "-w -s" cmd/tunaps/tunaps.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${dist}/tunapc.exe -v -ldflags "-w -s" cmd/tunapc/tunapc.go

popd
