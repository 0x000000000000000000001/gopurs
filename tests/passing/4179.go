package Main

import "gopurs/output/gopurs_runtime"

var RuntimeImportImpl = gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(moduleName gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(body gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					errStr := gopurs_runtime.Str("dynamic import not supported in Go")
					errVal := gopurs_runtime.Apply(just, errStr)
					action := gopurs_runtime.Apply(body, errVal)
					return gopurs_runtime.Apply(action, gopurs_runtime.Value{})
				})
			})
		})
	})
})
