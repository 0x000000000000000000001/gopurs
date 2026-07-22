package Data_Ring

import (
	"gopurs/output/gopurs_runtime"
	"math"
)

var IntSub = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(int(x.IntVal - y.IntVal))
	})
})

var NumSub = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(math.Float64frombits(uint64(x.IntVal)) - math.Float64frombits(uint64(y.IntVal)))
	})
})
