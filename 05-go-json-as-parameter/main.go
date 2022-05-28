package main

import (
	"log"
	"reflect"
	"syscall/js"
)

func Hello(this js.Value, args []js.Value) interface{} {

	// get the first parameter
	human := args[0]

	log.Println("🖐️ human:", reflect.TypeOf(human))

	// get members of human
	firstName := human.Get("firstName").String()
	lastName := human.Get("lastName").String()

	return map[string]interface{} {
		"message": "Hello " + firstName + " " + lastName,
		"author":  "@k33g_org",
	}

}

func main() {
	js.Global().Set("Hello", js.FuncOf(Hello))

	<-make(chan bool)
}
