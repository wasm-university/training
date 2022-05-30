# Rust and HTML and JSON

> See https://rustwasm.github.io/wasm-bindgen/reference/arbitrary-data-with-serde.html

<hr>

> create library structure
```bash
cargo new --lib hello
```

> update `Cargo.toml`
```toml
[dependencies]
serde = { version = "1.0", features = ["derive"] }
wasm-bindgen = { version = "0.2.80", features = ["serde-serialize"] }

[dependencies.web-sys]
version = "0.3.4"
features = [
  'Document',
  'Element',
  'HtmlElement',
  'Node',
  'Window',
]

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

// Called by our JS entry point to run the example
#[wasm_bindgen(start)]
pub fn run() -> Result<(), JsValue> {
    // Use `web_sys`'s global `window` function to get a handle on the global
    // window object.
    let window = web_sys::window().expect("no global `window` exists");
    let document = window.document().expect("should have a document on window");
    let body = document.body().expect("document should have a body");

    // Manufacture the element we're gonna append
    let h1 = document.create_element("h1")?;
    h1.set_text_content(Some("👋 Hello from Rust! 🦀"));

    let h2 = document.create_element("h2")?;
    h2.set_text_content(Some("😍 I'm a subtitle"));

    body.append_child(&h1)?;
    body.append_child(&h2)?;

    Ok(())
}
```

> build
```bash
cd hello
wasm-pack build --release --target web
```

> the output is in: `hello/pkg/`