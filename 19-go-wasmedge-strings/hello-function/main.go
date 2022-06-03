package main

import (
	"gitlab.com/wasmuniversity/demos-fight-strings-limitation/hello-function/helpers"
)

var message string = ""

func main() {
	message = "😃 Hello"
}


//export hello
func hello(parameters *int32) *byte {
	name := helpers.FromInt32PtrToString(parameters)
	returnValue := message + " " + name
	return helpers.FromStringToBytePtr(returnValue)
}
