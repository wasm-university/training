#!/bin/bash
wasmedge --reactor target/wasm32-wasi/debug/add.wasm add 20 22

wasmtime target/wasm32-wasi/debug/add.wasm --invoke add 15 27


