
const wasm = require("./hello/pkg/hello")

const fastify = require('fastify')({ logger: true })

// 🧰 Initialize settings
httpPort = process.env.HTTPS_PORT || 8080
maxListeners = process.env.MAX_LISTENERS || 1000
// Avoid: `MaxListenersExceededWarning: Possible EventEmitter memory leak detected`
require('events').setMaxListeners(maxListeners)


fastify.post(`/`, async (request, reply) => {
  let jsonParameters = request.body
  //let headers = request.headers
  return wasm.hello(jsonParameters)

})

const start = async () => {
  try {
    await fastify.listen(httpPort, "0.0.0.0")
    fastify.log.info(`server listening on ${fastify.server.address().port}`)

  } catch (error) {
    fastify.log.error(error)
  }
}
start()
