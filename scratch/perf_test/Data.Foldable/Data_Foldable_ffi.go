package Data_Foldable

import "gopurs/output/gopurs_runtime"

var FoldrArray = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(init gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
			arr := xs.PtrVal.([]gopurs_runtime.Value)
			acc := init
			for i := len(arr) - 1; i >= 0; i-- {
				acc = gopurs_runtime.Apply(gopurs_runtime.Apply(f, arr[i]), acc)
			}
			return acc
		})
	})
})

var FoldlArray = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(init gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
			arr := xs.PtrVal.([]gopurs_runtime.Value)
			acc := init
			for i := 0; i < len(arr); i++ {
				acc = gopurs_runtime.Apply(gopurs_runtime.Apply(f, acc), arr[i])
			}
			return acc
		})
	})
})
