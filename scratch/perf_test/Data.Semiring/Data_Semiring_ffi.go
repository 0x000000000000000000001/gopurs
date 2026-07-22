package Data_Semiring

import (
	"math"
	"gopurs/output/gopurs_runtime"
)

var IntAdd = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(int(x.IntVal + y.IntVal))
	})
})

var IntMul = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(int(x.IntVal * y.IntVal))
	})
})

var NumAdd = gopurs_runtime.Func(func(n1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(n2 gopurs_runtime.Value) gopurs_runtime.Value {
		f1 := math.Float64frombits(uint64(n1.IntVal))
		f2 := math.Float64frombits(uint64(n2.IntVal))
		return gopurs_runtime.Float(f1 + f2)
	})
})

var NumMul = gopurs_runtime.Func(func(n1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(n2 gopurs_runtime.Value) gopurs_runtime.Value {
		f1 := math.Float64frombits(uint64(n1.IntVal))
		f2 := math.Float64frombits(uint64(n2.IntVal))
		return gopurs_runtime.Float(f1 * f2)
	})
})
