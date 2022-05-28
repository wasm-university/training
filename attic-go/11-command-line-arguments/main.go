package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	argsWithoutCaller := os.Args[1:]


	fmt.Println("args:", args)
	fmt.Println("argsWithoutCaller:", argsWithoutCaller)

	fmt.Println(args[1], args[2], args[3])

	<-make(chan bool)
}
