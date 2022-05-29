# Rust and HTML and JSON

> See https://rustwasm.github.io/wasm-bindgen/reference/arbitrary-data-with-serde.html

<hr>

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

#[dependencies]
#wasm-bindgen = "0.2.50"
[dependencies]
serde = { version = "1.0", features = ["derive"] }
wasm-bindgen = { version = "0.2.50", features = ["serde-serialize"] }
```

> change the code source of `main.rs`
```rust
use wasm_bindgen::prelude::*;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct Message {
    pub text: String,
    pub author: String,
}

#[derive(Serialize, Deserialize)]
pub struct Response {
    pub text: String,
    pub author: String,
    pub message_text: String,
}


#[wasm_bindgen]
pub fn send(value: &JsValue) -> JsValue {
    // deserialize value (parameter) to message
    let message: Message = value.into_serde().unwrap();

    let response = Response {
        text: String::from(format!("👋 hello {}", message.author)),
        author: String::from("🦀"),
        message_text: String::from(format!("📝 your message:{}", message.text)),
    };

    // serialize response to JsValue
    return JsValue::from_serde(&response).unwrap()
}

```

> build
```bash
cd hello
wasm-pack build --release --target web
```

> the output is in: `hello/pkg/`