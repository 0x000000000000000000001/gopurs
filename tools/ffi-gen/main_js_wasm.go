package main

import "syscall/js"

func main() {
	js.Global().Set("packageGoSource", js.FuncOf(func(this js.Value, args []js.Value) any {
		result, err := Packages(args[0].String())
		diagnostic := ""
		if err != nil {
			diagnostic = err.Error()
		}
		return map[string]any{"json": result, "error": diagnostic}
	}))
	js.Global().Set("bindGoSource", js.FuncOf(func(this js.Value, args []js.Value) any {
		result, err := Bind(args[0].String())
		diagnostic := ""
		if err != nil {
			diagnostic = err.Error()
		}
		return map[string]any{"json": result, "error": diagnostic}
	}))
	js.Global().Set("parseFFI", js.FuncOf(func(this js.Value, args []js.Value) any {
		var prefixes []string
		if len(args) > 1 {
			prefixes = []string{args[1].String()}
		}
		result, err := parseFFI(args[0].String(), prefixes...)
		diagnostic := ""
		if err != nil {
			diagnostic = err.Error()
		}
		return map[string]any{"json": result, "error": diagnostic}
	}))
	select {}
}
