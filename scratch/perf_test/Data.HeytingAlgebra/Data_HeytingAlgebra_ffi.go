package Data_HeytingAlgebra

import "gopurs/output/gopurs_runtime"

var BoolConj = gopurs_runtime.Func(func(b1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b2 gopurs_runtime.Value) gopurs_runtime.Value {
		if b1.IntVal != 0 && b2.IntVal != 0 {
			return gopurs_runtime.Bool(true)
		}
		return gopurs_runtime.Bool(false)
	})
})

var BoolDisj = gopurs_runtime.Func(func(b1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(b2 gopurs_runtime.Value) gopurs_runtime.Value {
		if b1.IntVal != 0 || b2.IntVal != 0 {
			return gopurs_runtime.Bool(true)
		}
		return gopurs_runtime.Bool(false)
	})
})

var BoolNot = gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
	if b.IntVal != 0 {
		return gopurs_runtime.Bool(false)
	}
	return gopurs_runtime.Bool(true)
})
