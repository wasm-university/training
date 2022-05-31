# 👋 hello world 🌍

## How to create a "wasm rust function"

```bash
cargo new --lib hello
```

Add this to `Cargo.toml`:


```toml
[lib]
name = "hello"
path = "src/lib.rs"
crate-type =["cdylib"]

[dependencies]
serde = { version = "1.0", features = ["derive"] }
wasm-bindgen = { version = "0.2", features = ["serde-serialize"] }

```

Change `./src/libs.rs`

```rust
use wasm_bindgen::prelude::*;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct Question {
    pub text: String,
    pub author: String,
}

#[derive(Serialize, Deserialize)]
pub struct Answer {
    pub text: String,
    pub author: String,
}

#[wasm_bindgen]
pub fn handle(value: &JsValue) -> JsValue {
    // deserialize value (parameter) to question
    let question: Question = value.into_serde().unwrap();

    // serialize answer to JsValue
    let answer = Answer {
        text: String::from(format!("hello {}", question.author)),
        author: String::from("@k33g_org"),
    };

    return JsValue::from_serde(&answer).unwrap()
}

```

```bash
cd hello; wasm-pack build --target nodejs # 🖐🖐🖐
```

Then, the "distribution" files are located in `./hello/pkg`

## Test it

### Query

```bash
http POST http://0.0.0.0:8080/ text="what is the meaning of life?" author="bob"
```

### Query from outside

```bash
# type the command below to get a "public" url
gp url 8080
# use the terminal of the browser host
curl \
    -d '{"text":"hey", "author":"sam"}' \
    -H "Content-Type: application/json" \
    -X POST <UR>

```



### Load testing

```bash

```
