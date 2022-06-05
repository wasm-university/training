package main

import (
	"fmt"
	"log"

	"github.com/suborbital/sat/sat"
)

func main() {
	wasmModuleConfig, _ := sat.ConfigFromRunnableArg("hello/hello.wasm")
	
	satFunction, _ := sat.New(wasmModuleConfig, nil)

	result, err := satFunction.Exec([]byte("Bob!"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result.Output))

	/* 🤔
	jsonResult, err := result.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonResult)) 
	*/

}