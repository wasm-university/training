package main

import (
  "os"
  "github.com/second-state/WasmEdge-go/wasmedge"
)

// go run main.go hello-app/hello.wasm 
// go build
// ./wasmedge-cli hello-app/hello.wasm 
func main() {
  /// Expected Args[0]: program name (./wasmedge-cli)
  /// Expected Args[1]: wasm file (hello-app/hello.wasm)
  wasmedge.SetLogErrorLevel()

  var conf = wasmedge.NewConfigure(wasmedge.REFERENCE_TYPES)

  conf.AddConfig(wasmedge.WASI)

  var vm = wasmedge.NewVMWithConfig(conf)

  var wasi = vm.GetImportObject(wasmedge.WASI)
  wasi.InitWasi(
    os.Args[1:],     /// The args
    os.Environ(),    /// The envs
    []string{".:."}, /// The mapping directories
  )
  
  vm.RunWasmFile(os.Args[1], "_start") // 👋

  vm.Release()
  conf.Release()
}