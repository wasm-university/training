const fs = require('fs/promises')
require("./wasm_exec")

function runWasm(wasmFile, args) {
  const go = new Go()
  
  // 🖐️ hack for tiny go
  go.importObject.env["syscall/js.finalizeRef"] = () => {}

  return new Promise((resolve, reject) => {
    WebAssembly.instantiate(wasmFile, go.importObject)
    .then(result => {
      if(args) go.argv = args
      go.run(result.instance) 
      resolve(result.instance)
    })
    .catch(error => {
      reject(error)
    })
  })
}


fs.readFile('./main.wasm')
  .then(wasmFile => runWasm(wasmFile))
  .then(wasm => {

    console.log(Hello("Bob", "Morane"))

  })
  .catch(error => {
    console.log("ouch", error)
  }) 
