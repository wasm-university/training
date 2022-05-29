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
pub fn send(value: &JsValue) -> Result<JsValue, JsValue> {
    // deserialize value (parameter) to message
    let message: Message = value.into_serde().unwrap();

    let response = Response {
        text: String::from(format!("👋 hello {}", message.author)),
        author: String::from("🦀"),
        message_text: String::from(format!("📝 your message:{}", message.text)),
    };

    // serialize response to JsValue
    //return JsValue::from_serde(&response).unwrap()
    return Ok(JsValue::from_serde(&response).unwrap())
}
