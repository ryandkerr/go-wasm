package main

import (
	"syscall/js"

	"github.com/google/uuid"
)

func add(this js.Value, i []js.Value) interface{} {
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
	return js.ValueOf(i[0].Int() + i[1].Int())
}

func subtract(this js.Value, i []js.Value) interface{} {
	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
	return js.ValueOf(i[0].Int() - i[1].Int())
}

func myUuid(this js.Value, i []js.Value) interface{} {
	u := uuid.New()
	return js.ValueOf(u.String())
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("myUuid", js.FuncOf(myUuid))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
