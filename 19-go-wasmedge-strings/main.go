package main

import (
	"fmt"
	"github.com/wasm-university/training/19-go-wasmedge-strings/vmhelpers"
)

func main() {

	wasmVM, _ := vmhelpers.InitializeVMFromWasmFile("hello-function/hello.wasm")

	wasmVM.ExecuteMainFunction()

	result, _ := wasmVM.ExecuteFunction("hello", "Bob Morane")

	fmt.Println(result)

	wasmVM.Release()

}
