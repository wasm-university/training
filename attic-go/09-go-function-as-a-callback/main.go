package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	callback :=
		func(this js.Value, args []js.Value) interface{} {
			fmt.Println(
				"🎉 Hello",
				args[0].String(),
				args[1].String(),
				"from Go")
			return ""
		}

	js.Global().Call(
		"sayHello",
		"Bob",
		"Morane",
		js.FuncOf(callback))

	<-make(chan bool)

}
