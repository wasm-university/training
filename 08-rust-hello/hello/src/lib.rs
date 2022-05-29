use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn hello(s: String) -> String {
  let r = String::from("(from 🦀) 👋 hello ");
  
  return r + &s;
}
