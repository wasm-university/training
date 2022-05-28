package main

import (
	"syscall/js"
)

func Hello(this js.Value, args []js.Value) interface{} {
  message := args[0].String() // get the parameters
  return "😃 Hello " + message
}

// ! The `Hello` function takes two parameters (`this` and `args`) and returns an `interface{}` type.
// ! The first parameter (`this`) refers to the JavaScript's global object.
// ! The second parameter is a slice of `[]js.Value` representing 
// ! the arguments passed to the Javascript function call.

func main() {

	// ? To export the `Hello` function to the `global` JavaScript context, 
	// ? we use the `FuncOf` Go method to create a `Func` type):

	js.Global().Set("Hello", js.FuncOf(Hello))

	// ? The `Hello` function is attached to the `global` object (JavaScript). 
	// ? You can call the function like that: `global.Hello("Bob")` 
	// ? or directly with `Hello("Bob")` like any JavaScript function.
	
	// ! The Go runtime (JavaScript side) uses the wasm file as an application, 
	// ! runs this application, then exits. 
	// ! As the Go application has been completed and cleaned up, 
	// ! you cannot call the `Hello` function, and you get an error.

	// ? We want to tell the Go application that we don't want to exit 
	// ? and the way to do this is to use a channel 
	// ? (A channel waits for data and will pause the execution until it receives data). 
  // ? You just need to add the line below at the end of the `main` Go function:

	// make sure that the go program won't exit
	<-make(chan bool)
}
