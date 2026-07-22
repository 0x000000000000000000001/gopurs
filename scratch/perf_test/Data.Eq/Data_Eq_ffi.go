package Data_Eq

import "gopurs/output/gopurs_runtime"

// refEq for Int, Boolean, String are safe by value in Go!
var RefEq = gopurs_runtime.Func(func(r1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(r2 gopurs_runtime.Value) gopurs_runtime.Value {
		if r1.Type == r2.Type && r1.IntVal == r2.IntVal && r1.StrVal == r2.StrVal && r1.PtrVal == r2.PtrVal {
			return gopurs_runtime.Bool(true)
		}
		return gopurs_runtime.Bool(false)
	})
})

var EqBooleanImpl = RefEq
var EqIntImpl = RefEq
var EqNumberImpl = RefEq
var EqCharImpl = RefEq
var EqStringImpl = RefEq

var EqArrayImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(ys gopurs_runtime.Value) gopurs_runtime.Value {
			arr1 := xs.PtrVal.([]gopurs_runtime.Value)
			arr2 := ys.PtrVal.([]gopurs_runtime.Value)
			if len(arr1) != len(arr2) {
				return gopurs_runtime.Bool(false)
			}
			for i := range arr1 {
				thunk := gopurs_runtime.Apply(f, arr1[i])
				res := gopurs_runtime.Apply(thunk, arr2[i])
				if res.IntVal == 0 {
					return gopurs_runtime.Bool(false)
				}
			}
			return gopurs_runtime.Bool(true)
		})
	})
})
