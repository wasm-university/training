package main

import (
	"fmt"
)

func main() {

	fmt.Println("👋 Hello World from (Tiny)Go 😍")

	<-make(chan bool)
}
