#!/bin/bash
#tinygo build -o main.wasm -target wasm ./main.go

tinygo build -o hello.wasm -target wasi ./main.go

ls -lh *.wasm