

```bash
subo create runnable as-hello-service --lang assemblyscript
subo build as-hello-service

SAT_HTTP_PORT=8081 sat ./as-hello-service/as-hello-service.wasm

# "👋 Hello World 🌍 From AssemblyScript"

subo create runnable js-hello-service --lang javascript
subo build js-hello-service

SAT_HTTP_PORT=8082 sat ./js-hello-service/js-hello-service.wasm

# "👋 Hello World 🌍 from JavaScript 😍"

subo create runnable hello --lang grain
subo build hello

SAT_HTTP_PORT=8080 sat ./hello/hello.wasm

curl -d "Bob Morane" http://localhost:8080; echo ""

# 👀 at the grain terminal
```
