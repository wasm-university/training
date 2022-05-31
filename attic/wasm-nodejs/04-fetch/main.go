package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	thenFunc :=
		func(this js.Value, args []js.Value) interface{} {
			fmt.Println("🎉 All good:", args)
			return ""
		}

	catchFunc :=
		func(this js.Value, args []js.Value) interface{} {
			fmt.Println("😡 Ouch:", args)
			return ""
		}

	js.Global().Call("host_fetch", "https://jsonplaceholder.typicode.com/todos/1").Call("then", js.FuncOf(thenFunc)).Call("catch", js.FuncOf(catchFunc))

	// https://github.com/golang/go/issues/29845


	<-make(chan bool)
}
