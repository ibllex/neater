#!/bin/bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/neater-darwin-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/neater-linux-amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/neater-windows-amd64.exe

CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ./bin/neater-darwin-386
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./bin/neater-linux-386
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./bin/neater-windows-386.exe
