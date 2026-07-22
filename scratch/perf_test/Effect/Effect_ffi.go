package Effect

import "gopurs/output/gopurs_runtime"

var PureE = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return a
	})
})

var BindE = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			resA := gopurs_runtime.Apply(a, gopurs_runtime.Value{})
			nextThunk := gopurs_runtime.Apply(f, resA)
			return gopurs_runtime.Apply(nextThunk, gopurs_runtime.Value{})
		})
	})
})

var UntilE = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		for {
			res := gopurs_runtime.Apply(f, gopurs_runtime.Value{})
			if res.IntVal != 0 {
				break
			}
		}
		return gopurs_runtime.Value{}
	})
})

var WhileE = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			for {
				res := gopurs_runtime.Apply(f, gopurs_runtime.Value{})
				if res.IntVal == 0 {
					break
				}
				gopurs_runtime.Apply(a, gopurs_runtime.Value{})
			}
			return gopurs_runtime.Value{}
		})
	})
})

var ForE = gopurs_runtime.Func(func(lo gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(hi gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				loInt := lo.IntVal
				hiInt := hi.IntVal
				for i := loInt; i < hiInt; i++ {
					thunk := gopurs_runtime.Apply(f, gopurs_runtime.Int(int(i)))
					gopurs_runtime.Apply(thunk, gopurs_runtime.Value{})
				}
				return gopurs_runtime.Value{}
			})
		})
	})
})

var ForeachE = gopurs_runtime.Func(func(as gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			arr := as.PtrVal.([]gopurs_runtime.Value)
			for _, v := range arr {
				thunk := gopurs_runtime.Apply(f, v)
				gopurs_runtime.Apply(thunk, gopurs_runtime.Value{})
			}
			return gopurs_runtime.Value{}
		})
	})
})
