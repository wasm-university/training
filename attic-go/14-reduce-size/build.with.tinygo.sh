#!/bin/bash
#tinygo_version="0.21.0"
#wget https://raw.githubusercontent.com/tinygo-org/tinygo/v${tinygo_version}/targets/wasm_exec.js
tinygo build -o main.wasm -target wasm ./main.go

ls -lh *.wasm
