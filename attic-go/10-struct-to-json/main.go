package main

import (
	"encoding/json"
	"log"
	"syscall/js"
)

type Human struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var human Human = Human{"Bob", "Morane"}

func GiveMeHumanJsonString(this js.Value, args []js.Value) interface{} {

	jsonHuman, err := json.Marshal(human)

	if err != nil {
		log.Fatalf(
			"Error occured during marshaling: %s",
			err.Error())
	}

	return string(jsonHuman)

}

func GiveMeHumanJsonObject(this js.Value, args []js.Value) interface{} {

	jsonHuman, err := json.Marshal(human)

	if err != nil {
		log.Fatalf(
			"Error occured during marshaling: %s",
			err.Error())
	}

	JSON := js.Global().Get("JSON")
	jsonString := string(jsonHuman)

	return JSON.Call("parse", jsonString)

}

func main() {

	js.Global().Set(
		"GiveMeHumanJsonString",
		js.FuncOf(GiveMeHumanJsonString))

	js.Global().Set(
		"GiveMeHumanJsonObject",
		js.FuncOf(GiveMeHumanJsonObject))

	<-make(chan bool)
}
