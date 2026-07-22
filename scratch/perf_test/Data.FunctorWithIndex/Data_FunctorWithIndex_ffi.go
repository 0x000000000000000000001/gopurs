package Data_FunctorWithIndex

import "gopurs/output/gopurs_runtime"

var MapWithIndexArray = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		xsArr := xs.PtrVal.([]gopurs_runtime.Value)
		result := make([]gopurs_runtime.Value, len(xsArr))
		for i, x := range xsArr {
			result[i] = gopurs_runtime.Apply(gopurs_runtime.Apply(f, gopurs_runtime.Int(i)), x)
		}
		return gopurs_runtime.Array(result)
	})
})
