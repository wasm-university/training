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
pub fn hello(value: &JsValue) -> JsValue {
    // deserialize value (parameter) to question
    let question: Question = value.into_serde().unwrap();

    // serialize answer to JsValue
    let answer = Answer {
        text: String::from(format!("🍊 hello {}", question.author)),
        author: String::from("@k33g_org"),
    };

    return JsValue::from_serde(&answer).unwrap()
}
