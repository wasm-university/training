package main

import (
	"strings"

	"github.com/suborbital/reactr/api/tinygo/runnable"
)

type Hey struct{}

func (h Hey) Run(input []byte) ([]byte, error) {

	someStrings := strings.Split(string(input), ";")
	

	return []byte("Go: " + strings.Join(someStrings, "-")), nil
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(Hey{})
}
