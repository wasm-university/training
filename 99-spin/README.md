# Spin

- https://spin.fermyon.dev/

## GoLang

```bash
tinygo build -wasm-abi=generic -target=wasi -no-debug -o main.wasm main.go

spin up --file spin.toml

curl -i localhost:3000/hello
```

