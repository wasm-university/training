package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println(js.Global().Call("sayHello", "Bill"))

	message := js.Global().Get("message").String()
	fmt.Println("message (before):", message)

	js.Global().Set("message", "🚀 this is a message from GoLang")

	bill := js.Global().Get("bill")
	fmt.Println("👋 bill is an", bill)
	fmt.Println("👋 age of bill:", bill.Get("age"))

	bill.Set("firstName", "Bill")
	bill.Set("lastName", "Ballantine")

	<-make(chan bool)
}
