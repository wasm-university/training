package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("content-type: text/plain;utf-8")
	fmt.Println("")

	args := os.Args
	argsWithoutCaller := os.Args[1:]

	fmt.Println(args)
	fmt.Println(argsWithoutCaller)

	var reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	fmt.Println(message)

}
