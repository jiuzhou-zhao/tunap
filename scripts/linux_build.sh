#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tunaps -v -ldflags "-w -s" cmd/tunaps/tunaps.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tunapc -v -ldflags "-w -s" cmd/tunapc/tunapc.go
