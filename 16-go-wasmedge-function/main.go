package main

import (
	"fmt"
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

  //vm.RunWasmFile(os.Args[1], "_start")
  err := vm.LoadWasmFile(os.Args[1])
  if err != nil {
    fmt.Println("failed to load wasm")
  }
  vm.Validate()
  vm.Instantiate()

  result, _ := vm.Execute("add", int32(12), int32(30))

  fmt.Println(result)

  vm.Release()
  conf.Release()
}