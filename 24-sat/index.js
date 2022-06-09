const http = require('http');

const requestListener = function (req, res) {
  console.log("👋, request from wasm")
  res.writeHead(200);
  res.end("👋 Hello, World! 🌍");
}

const server = http.createServer(requestListener);
server.listen(8080);