let fastifyServerOptions = {
  logger: process.env.LOGGING || true
}
const fastify = require('fastify')(fastifyServerOptions)


const fs = require('fs')
require("./wasm_exec")

async function runWasm(wasmFile, args) {
  const go = new Go()
  try {
    let { instance } = await WebAssembly.instantiate(wasmFile, go.importObject)
    if(args) go.argv = args
    go.run(instance) 
  } catch(err) {
    throw err
  }
}

const loadWasmFile = async () => {

  const wasmFile = fs.readFileSync('./main.wasm')
  await runWasm(wasmFile)  
}
loadWasmFile()

// Declare a route
fastify.get('/human/:first/:last', async (request, reply) => {
  return Hello(request.params.first, request.params.last)
})


const startService = async () => {
  try {
    await fastify.listen(8080, "0.0.0.0")
    console.log("server listening on:", fastify.server.address().port)

  } catch (error) {
    fastify.log.error(error)
  }
}
startService()
