package Control_Bind

import "gopurs/output/gopurs_runtime"

var ArrayBind = gopurs_runtime.Func(func(arr gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
		input := arr.PtrVal.([]gopurs_runtime.Value)
		var result []gopurs_runtime.Value
		for _, v := range input {
			res := gopurs_runtime.Apply(f, v).PtrVal.([]gopurs_runtime.Value)
			result = append(result, res...)
		}
		return gopurs_runtime.Array(result)
	})
})
