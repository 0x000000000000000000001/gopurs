package ScratchFFI
import "gopurs/output/gopurs_runtime"
var ShowImpl = gopurs_runtime.Func(func(showFn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(showFn, val)
	})
})
