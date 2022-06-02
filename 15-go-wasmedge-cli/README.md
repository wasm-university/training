# ebook go samples template

Ref: https://wasmedge.org/book/en/embed/go.html

Developers must install the WasmEdge shared library with the same WasmEdge-go release version.

```bash
curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.10.0
source /home/gitpod/.wasmedge/env 
go get github.com/second-state/WasmEdge-go/wasmedge@v0.10.0
```

```bash
go build
./wasmedge-cli hello-app/hello.wasm bob morane
```
