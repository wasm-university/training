# Spin

- https://spin.fermyon.dev/

## GoLang

```bash
tinygo build -wasm-abi=generic -target=wasi -no-debug -o main.wasm main.go

spin up --file spin.toml

curl -i localhost:3000/hello
```

## Refs

- https://www.fermyon.com/blog/tinygo-webassembly-favicon-server
- https://spin.fermyon.dev/go-components/
- Hippo https://spin.fermyon.dev/deploying-to-hippo/
