package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Wasm")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}

func getNthFibonacci(n int) int {
	// I know a recursive solution requires polynomial time, that's why I'm
	// using it, how far can WASM be pushed?
	if n <= 1 {
		return n
	}
	return getNthFibonacci(n-1) + getNthFibonacci(n-2)
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJSON(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
