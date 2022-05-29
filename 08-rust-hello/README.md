# Rust and HTML

> create library structure
```bash
cargo new --lib hello
```

> update `Cargo.toml`
```toml
[lib]
name = "hello"
path = "src/lib.rs"
crate-type =["cdylib"]

[dependencies]
wasm-bindgen = "0.2.50"
```

> change the code source of `main.rs`
```rust
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn hello(s: String) -> String {
  let r = String::from("👋 hello ");
  
  return r + &s;
}

```

> build
```bash
cd hello
wasm-pack build --release --target web
```

> the output is in: `hello/pkg/`