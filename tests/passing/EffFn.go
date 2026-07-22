package Main
import "gopurs/output/gopurs_runtime"
var Add3 = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(a.IntVal + b.IntVal + c.IntVal)
		})
	})
})
