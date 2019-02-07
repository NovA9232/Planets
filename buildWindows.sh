#!/bin/bash
export GOPATH="$PWD/"
GOOS=windows GOARCH=386 go build -o Planets.exe
