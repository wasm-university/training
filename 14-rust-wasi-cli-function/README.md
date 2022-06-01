# WASI & Rust

```bash
# Create
rustup target add wasm32-wasi
cargo new --lib hello

# Build
cd hello
cargo build --target wasm32-wasi

# see: target/wasm32-wasi/debug/add.wasm
wasmedge --reactor target/wasm32-wasi/debug/add.wasm add 2 2
wasmtime target/wasm32-wasi/debug/add.wasm --invoke add 1 2
```
