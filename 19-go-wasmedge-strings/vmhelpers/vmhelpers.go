package vmhelpers

import (
	"fmt"
	"strings"
	"github.com/second-state/WasmEdge-go/wasmedge"
)

type WasmVM struct {
	config *wasmedge.Configure
	store *wasmedge.Store
	vm *wasmedge.VM
}

func InitializeVMFromWasmFile(wasmPathFile string) (WasmVM, error) {
	wasmedge.SetLogErrorLevel()
	conf := wasmedge.NewConfigure(wasmedge.WASI)
	store := wasmedge.NewStore()

	vm := wasmedge.NewVMWithConfigAndStore(conf, store)

	err := vm.LoadWasmFile(wasmPathFile)
	if err != nil {
		fmt.Println("failed to load wasm")
		return WasmVM{}, err
	}
	vm.Validate()
	vm.Instantiate()
	return WasmVM{conf, store, vm}, nil
}

func (wasmVM WasmVM) ExecuteMainFunction() (interface{}, error) {
	// run the main function
	return wasmVM.vm.Execute("_start")
}

func (wasmVM WasmVM) writeStringValueToMemory(value string) (int32, *wasmedge.Memory, []byte) {
	lengthOfStringParam := len(value)
	// Allocate memory for the value, and get a pointer to it.
	// Include a byte for the NULL terminator we add below.
	allocateResult, _ := wasmVM.vm.Execute("malloc", int32(lengthOfStringParam+1))
	inputPointer := allocateResult[0].(int32)

	// Write the string value into the memory.
	mod := wasmVM.vm.GetActiveModule()
	mem := mod.FindMemory("memory")
	//mem := wasmVM.store.FindMemory("memory")
	memData, _ := mem.GetData(uint(inputPointer), uint(lengthOfStringParam+1))
	copy(memData, value)

	// C-string terminates by NULL.
	memData[lengthOfStringParam] = 0
	return inputPointer, mem, memData
}

func (wasmVM WasmVM) getStringValueFromResult(wasmFunctionResult []interface{}, mem *wasmedge.Memory, memData []byte) (string, int32) {
	outputPointer := wasmFunctionResult[0].(int32)

	pageSize := mem.GetPageSize()
	// Read the result of the wasm function.
	memData, _ = mem.GetData(uint(0), uint(pageSize * 65536))
	nth := 0
	var output strings.Builder

	for {
		if memData[int(outputPointer) + nth] == 0 {
			break
		}

		output.WriteByte(memData[int(outputPointer) + nth])
		nth++
	}

	return output.String(), outputPointer
}

func (wasmVM WasmVM) ExecuteFunction(functionName, stringParam string) (string, error) {

	inputPointer, mem, memData := wasmVM.writeStringValueToMemory(stringParam)

	// Run the "wasm" function. Given the pointer to the string value.
	wasmFunctionResult, err := wasmVM.vm.Execute(functionName, inputPointer)
	if err != nil {
		return "", err
	}

	result, outputPointer := wasmVM.getStringValueFromResult(wasmFunctionResult, mem, memData)

	// Deallocate the input, and the output.
	wasmVM.vm.Execute("free", outputPointer)
	wasmVM.vm.Execute("free", inputPointer)
	
	return result, err

}

func (wasmVM WasmVM) Release() {
	wasmVM.vm.Release()
	wasmVM.store.Release()
	wasmVM.config.Release()
}

