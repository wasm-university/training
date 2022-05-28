package main

import (
	"syscall/js"
)

func Hello(this js.Value, args []js.Value) interface{} {

	// get items of an array
	firstName := args[0].Index(0)
	lastName := args[0].Index(1)

	// cast the array into interface
	return []interface{} {
		"Hello",
		firstName,
		lastName,
		"by @k33g_org",
	}

}

func main() {
	js.Global().Set("Hello", js.FuncOf(Hello))

	<-make(chan bool)
}
