#!/bin/bash
#GOOS=js GOARCH=wasm go build -o main.wasm
tinygo build -o main.wasm -target wasi ./main.go

ls -lh *.wasm