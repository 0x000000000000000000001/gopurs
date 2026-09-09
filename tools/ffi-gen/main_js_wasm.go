package main

import "syscall/js"

func main() {
	js.Global().Set("parseFFI", js.FuncOf(func(this js.Value, args []js.Value) any {
		result, err := parseFFI(args[0].String())
		diagnostic := ""
		if err != nil {
			diagnostic = err.Error()
		}
		return map[string]any{"json": result, "error": diagnostic}
	}))
	select {}
}
