const fs = require('fs')
require("./wasm_exec")

const fetch = require('node-fetch')

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

function host_fetch(url) {
  console.log("🤖", url)
  return new Promise((resolve, reject) => {
    
    fetch(url)
      .then(response => response.text())
      .then(data => {
        console.log("🤖 data:", data)
        resolve(data)
      })
      .catch(error => reject(error) )
  })

}


global.host_fetch = host_fetch


fs.readFile('./main.wasm', (err, wasmFile) => {
  if (err) {
    console.error(err)
    return
  }
  runWasm(wasmFile).then(wasm => {
    // foo
  }).catch(error => {
    console.log("ouch", error)
  }) 
})

