#!/bin/bash
tinygo build -wasm-abi=generic -target=wasi -no-debug -o main.wasm main.go
ls -lh *.wasm

