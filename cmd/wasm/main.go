package main

import (
	"fmt"
	"math"
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
	if n < 0 {
		math.NaN()
	} else if n <= 1 {
		return n
	}
	return getNthFibonacci(n-1) + getNthFibonacci(n-2)
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputN := args[0].Int()
		fmt.Printf("input %d\n", inputN)
		fibN := getNthFibonacci(inputN)
		return fibN
	})
	return jsonFunc
}
