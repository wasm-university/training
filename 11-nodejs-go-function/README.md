

> download `wasm_exec.js`
```bash
tinygo_version="0.23.0"
wget https://raw.githubusercontent.com/tinygo-org/tinygo/v${tinygo_version}/targets/wasm_exec.js
```

> build
```bash
tinygo build -o main.wasm -target wasm ./main.go
```