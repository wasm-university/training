#!/bin/bash
cargo build --target wasm32-wasi --release
cp target/wasm32-wasi/release/wasm_rust_demo.wasm wasm_rust_demo.wasm
ls -lh 