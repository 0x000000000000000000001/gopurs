package Data_Functor

import "gopurs/output/gopurs_runtime"

var ArrayMap = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(arr gopurs_runtime.Value) gopurs_runtime.Value {
		a := arr.PtrVal.([]gopurs_runtime.Value)
		result := make([]gopurs_runtime.Value, len(a))
		for i, v := range a {
			result[i] = gopurs_runtime.Apply(f, v)
		}
		return gopurs_runtime.Array(result)
	})
})
