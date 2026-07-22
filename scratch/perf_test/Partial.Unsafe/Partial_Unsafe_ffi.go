package Partial_Unsafe

import "gopurs/output/gopurs_runtime"

var X_UnsafePartial = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(f, gopurs_runtime.Value{})
})
