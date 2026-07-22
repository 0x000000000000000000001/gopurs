package Data_Semigroup

import "gopurs/output/gopurs_runtime"

var ConcatString = gopurs_runtime.Func(func(s1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(s1.StrVal + s2.StrVal)
	})
})

var ConcatArray = gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(ys gopurs_runtime.Value) gopurs_runtime.Value {
		xsArr := xs.PtrVal.([]gopurs_runtime.Value)
		ysArr := ys.PtrVal.([]gopurs_runtime.Value)
		if len(xsArr) == 0 {
			return ys
		}
		if len(ysArr) == 0 {
			return xs
		}
		res := make([]gopurs_runtime.Value, 0, len(xsArr)+len(ysArr))
		res = append(res, xsArr...)
		res = append(res, ysArr...)
		return gopurs_runtime.Array(res)
	})
})
