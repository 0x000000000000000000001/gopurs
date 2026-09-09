package main

import "syscall/js"

func main() {
	js.Global().Set("parseFFI", js.FuncOf(func(this js.Value, args []js.Value) any {
		return parseFFI(args[0].String())
	}))
	select {}
}
