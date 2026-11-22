package purescript

import (
	"gopurs/output/gopurs_runtime"
	"sort"
)

func Foreign_Object__CopyST(m map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		newMap := make(map[string]interface{})
		for k, v := range m {
			newMap[k] = v
		}
		return newMap
	}
}

var Foreign_Object_Empty = map[string]interface{}{}

func Foreign_Object_RunST(f func(interface{}) interface{}) interface{} {
	return f(nil)
}

func Foreign_Object__FmapObject(m0 map[string]interface{}, f func(interface{}) interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range m0 {
		m[k] = f(v)
	}
	return m
}

func Foreign_Object__MapWithKey(m0 map[string]interface{}, f func(string) func(interface{}) interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range m0 {
		m[k] = f(k)(v)
	}
	return m
}

func Foreign_Object__FoldM(bind func(interface{}) func(interface{}) interface{}) func(func(interface{}) func(string) func(interface{}) interface{}) func(interface{}) func(map[string]interface{}) interface{} {
	return func(f func(interface{}) func(string) func(interface{}) interface{}) func(interface{}) func(map[string]interface{}) interface{} {
		return func(mz interface{}) func(map[string]interface{}) interface{} {
			return func(m map[string]interface{}) interface{} {
				acc := mz
				g := func(k string) func(interface{}) interface{} {
					return func(z interface{}) interface{} {
						return f(z)(k)(m[k])
					}
				}
				keys := make([]string, 0, len(m))
				for k := range m {
					keys = append(keys, k)
				}
				sort.Strings(keys)

				for _, k := range keys {
					acc = bind(acc)(g(k))
				}
				return acc
			}
		}
	}
}

func Foreign_Object__FoldSCObject(m map[string]interface{}, z interface{}, f func(interface{}) func(string) func(interface{}) interface{}, fromMaybe func(interface{}) func(interface{}) interface{}) interface{} {
	acc := z
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		maybeR := f(acc)(k)(v)
		r := fromMaybe(nil)(maybeR)
		if r == nil {
			return acc
		} else {
			acc = r
		}
	}
	return acc
}

func Foreign_Object_All(f func(string) func(interface{}) bool) func(map[string]interface{}) bool {
	return func(m map[string]interface{}) bool {
		for k, v := range m {
			if !f(k)(v) {
				return false
			}
		}
		return true
	}
}

func Foreign_Object_Size(m map[string]interface{}) int64 {
	return int64(len(m))
}

func Foreign_Object__Lookup(no interface{}, yes func(interface{}) interface{}, k string, m gopurs_runtime.Value) interface{} {
	object, ok := gopurs_runtime.ReadJSONObjectValue(m)
	if !ok {
		return no
	}
	if val, ok := object.Lookup(k); ok {
		return yes(val)
	}
	return no
}

func Foreign_Object__LookupST(no interface{}, yes func(interface{}) interface{}, k string, m map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		if val, ok := m[k]; ok {
			return yes(val)
		}
		return no
	}
}

func Foreign_Object_ToArrayWithKey(f func(string) func(interface{}) interface{}) func(map[string]interface{}) []interface{} {
	return func(m map[string]interface{}) []interface{} {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		r := make([]interface{}, 0, len(m))
		for _, k := range keys {
			r = append(r, f(k)(m[k]))
		}
		return r
	}
}

func Foreign_Object_Keys(m map[string]interface{}) []interface{} {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	r := make([]interface{}, len(keys))
	for i, k := range keys {
		r[i] = k
	}
	return r
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Foreign_Object__CopyST = // TAST: (ForAll [a, b, r] (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Foreign_Object__CopyST(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res := go_res(arg)
		return gopurs_runtime.Box(inner_res)
	})
})
var _Gopurs_Foreign_Object__FmapObject = // TAST: (ForAll [a, b] (ADT ["Data","Function","Uncurried","Fn2"] [(ADT ["Foreign","Object","Object"] [(TypeVar a)]), (Func [(TypeVar a)] (TypeVar b)), (ADT ["Foreign","Object","Object"] [(TypeVar b)])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_res := Foreign_Object__FmapObject(go_arg0, go_arg1)
	return gopurs_runtime.Any(go_res)
})
var _Gopurs_Foreign_Object__FoldM = // TAST: (ForAll [a, m, z] (Func [(Func [(TypeVar m), (Func [(TypeVar z)] (TypeVar m))] (TypeVar m)), (Func [(TypeVar z), String, (TypeVar a)] (TypeVar m)), (TypeVar m), (ADT ["Foreign","Object","Object"] [(TypeVar a)])] (TypeVar m)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_res := Foreign_Object__FoldM(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res := go_res(func(p99_0 any) func(string) func(any) any {
			inner_res99 := gopurs_runtime.Apply(arg, gopurs_runtime.Box(p99_0))
			return func(p100_0 string) func(any) any {
				inner_res100 := gopurs_runtime.Apply(inner_res99, gopurs_runtime.Box(p100_0))
				return func(p101_0 any) any {
					return gopurs_runtime.Apply(inner_res100, gopurs_runtime.Box(p101_0))
				}
			}
		})
		return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := inner_res(arg)
			return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := inner_res(gopurs_runtime.Unbox[map[string]any](arg))
				return gopurs_runtime.Box(inner_res)
			})
		})
	})
})
var _Gopurs_Foreign_Object__FoldSCObject = // TAST: (ForAll [a, z] (ADT ["Data","Function","Uncurried","Fn4"] [(ADT ["Foreign","Object","Object"] [(TypeVar a)]), (TypeVar z), (Func [(TypeVar z), String, (TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar z)])), (ForAll [b] (Func [(TypeVar b), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])] (TypeVar b))), (TypeVar z)]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_arg1 := arg1
	go_arg2 := func(p0_0 any) func(string) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		return func(p1_0 string) func(any) any {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 any) any {
				return gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			}
		}
	}
	go_arg3 := func(p0_0 any) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_res := Foreign_Object__FoldSCObject(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_Object__Lookup = // TAST: (ForAll [a, z] (ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar z), (Func [(TypeVar a)] (TypeVar z)), String, (ADT ["Foreign","Object","Object"] [(TypeVar a)]), (TypeVar z)]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := arg3
	go_res := Foreign_Object__Lookup(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_Object__LookupST = // TAST: (ForAll [a, r, z] (ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar z), (Func [(TypeVar a)] (TypeVar z)), String, (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)]), (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar z)])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.UnboxObject(arg3)
	go_res := Foreign_Object__LookupST(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res := go_res(arg)
		return gopurs_runtime.Box(inner_res)
	})
})
var _Gopurs_Foreign_Object__MapWithKey = // TAST: (ForAll [a, b] (ADT ["Data","Function","Uncurried","Fn2"] [(ADT ["Foreign","Object","Object"] [(TypeVar a)]), (Func [String, (TypeVar a)] (TypeVar b)), (ADT ["Foreign","Object","Object"] [(TypeVar b)])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_arg1 := func(p0_0 string) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_res := Foreign_Object__MapWithKey(go_arg0, go_arg1)
	return gopurs_runtime.Any(go_res)
})
var _Gopurs_Foreign_Object_All = // TAST: (ForAll [a] (Func [(Func [String, (TypeVar a)] Boolean), (ADT ["Foreign","Object","Object"] [(TypeVar a)])] Boolean))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) func(any) bool {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) bool {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[bool](inner_res1)
		}
	}
	go_res := Foreign_Object_All(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res := go_res(gopurs_runtime.Unbox[map[string]any](arg))
		return gopurs_runtime.Box(inner_res)
	})
})
var _Gopurs_Foreign_Object_Empty = // TAST: (ForAll [a] (ADT ["Foreign","Object","Object"] [(TypeVar a)]))
gopurs_runtime.Box(Foreign_Object_Empty)
var _Gopurs_Foreign_Object_Keys = // TAST: (ForAll [a] (Func [(ADT ["Foreign","Object","Object"] [(TypeVar a)])] (Array String)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Foreign_Object_Keys(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Foreign_Object_RunST = // TAST: (ForAll [a] (Func [(ForAll [r] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])]))] (ADT ["Foreign","Object","Object"] [(TypeVar a)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_res := Foreign_Object_RunST(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_Object_Size = // TAST: (ForAll [a] (Func [(ADT ["Foreign","Object","Object"] [(TypeVar a)])] Int))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Foreign_Object_Size(go_arg0)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Foreign_Object_ToArrayWithKey = // TAST: (ForAll [a, b] (Func [(Func [String, (TypeVar a)] (TypeVar b)), (ADT ["Foreign","Object","Object"] [(TypeVar a)])] (Array (TypeVar b))))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_res := Foreign_Object_ToArrayWithKey(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
		inner_res := go_res(gopurs_runtime.Unbox[map[string]any](arg))
		return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(inner_res))
			for i, v := range inner_res {
				res_arr[i] = gopurs_runtime.Box(v)
			}
			return gopurs_runtime.Array(res_arr)
		}()
	})
})
