package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	thenFunc :=
		func(this js.Value, args []js.Value) interface{} {
			fmt.Println("🎉 All good:", args[0].String())
			return ""
		}

	catchFunc :=
		func(this js.Value, args []js.Value) interface{} {
			fmt.Println("😡 Ouch:", args[0].Get("message"))
			return ""
		}

	js.Global().Call("compute", false).Call("then", js.FuncOf(thenFunc)).Call("catch", js.FuncOf(catchFunc))

	js.Global().Call("compute", true).Call("then", js.FuncOf(thenFunc)).Call("catch", js.FuncOf(catchFunc))

	<-make(chan bool)
}
