package Data_Show

import (
	"fmt"
	"math"
	"gopurs/output/gopurs_runtime"
)

var ShowIntImpl = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(fmt.Sprintf("%v", n.IntVal))
})

var ShowNumberImpl = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	// For now, assume Number is FloatVal in Value
	return gopurs_runtime.Str(fmt.Sprintf("%f", math.Float64frombits(uint64(n.IntVal))))
})

var ShowCharImpl = gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(fmt.Sprintf("'%s'", c.StrVal))
})

var ShowStringImpl = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(fmt.Sprintf("%q", s.StrVal))
})

var ShowArrayImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		// xs is an Array (Wait, what is Array in Value?)
		// Let's assume xs.ArrayVal is []gopurs_runtime.Value
		arr := xs.PtrVal.([]gopurs_runtime.Value)
		res := "["
		for i, v := range arr {
			if i > 0 {
				res += ","
			}
			strVal := gopurs_runtime.Apply(f, v)
			res += strVal.StrVal
		}
		res += "]"
		return gopurs_runtime.Str(res)
	})
})
