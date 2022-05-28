package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("[go]", js.Global().Call("sayHello", "Bill"))

	message := js.Global().Get("message").String()
	
	fmt.Println("[go] message (before):", message)

	js.Global().Set("message", "🚀 this is a message from GoLang")

	<-make(chan bool)
}
