package Control_Monad_ST_Internal

import "gopurs/output/gopurs_runtime"

var Map_ = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f, gopurs_runtime.Apply(a, gopurs_runtime.Value{}))
		})
	})
})

var Pure_ = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return a
	})
})

var Bind_ = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(f, gopurs_runtime.Apply(a, gopurs_runtime.Value{})), gopurs_runtime.Value{})
		})
	})
})

var Run = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(f, gopurs_runtime.Value{})
})

var While = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			for gopurs_runtime.Apply(f, gopurs_runtime.Value{}).IntVal != 0 {
				gopurs_runtime.Apply(a, gopurs_runtime.Value{})
			}
			return gopurs_runtime.Value{}
		})
	})
})

var For_ = gopurs_runtime.Func(func(lo gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(hi gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				start := lo.IntVal
				end := hi.IntVal
				for i := start; i < end; i++ {
					gopurs_runtime.Apply(gopurs_runtime.Apply(f, gopurs_runtime.Int(int(i))), gopurs_runtime.Value{})
				}
				return gopurs_runtime.Value{}
			})
		})
	})
})

var Foreach = gopurs_runtime.Func(func(as gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			arr := as.PtrVal.([]gopurs_runtime.Value)
			for _, item := range arr {
				gopurs_runtime.Apply(gopurs_runtime.Apply(f, item), gopurs_runtime.Value{})
			}
			return gopurs_runtime.Value{}
		})
	})
})

var New_ = gopurs_runtime.Func(func(val gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		ref := &val
		return gopurs_runtime.Value{PtrVal: ref}
	})
})

var Read = gopurs_runtime.Func(func(ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		ptr := ref.PtrVal.(*gopurs_runtime.Value)
		return *ptr
	})
})

var ModifyImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(ref gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			ptr := ref.PtrVal.(*gopurs_runtime.Value)
			t := gopurs_runtime.Apply(f, *ptr)
			
			// t is { state: s, value: v }
			dict := t.PtrVal.(map[string]gopurs_runtime.Value)
			*ptr = dict["state"]
			return dict["value"]
		})
	})
})

var Write = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(ref gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			ptr := ref.PtrVal.(*gopurs_runtime.Value)
			*ptr = a
			return a
		})
	})
})
