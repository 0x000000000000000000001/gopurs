package purescript

import "gopurs/output/gopurs_runtime"

type r1 interface{}
type r2 interface{}

func Record_Unsafe_Union__UnsafeUnionFn(r1, r2 interface{}) interface{} {
	rec1 := r1.(gopurs_runtime.Value)
	rec2 := r2.(gopurs_runtime.Value)
	m1 := gopurs_runtime.RecordToMap(rec1)
	m2 := gopurs_runtime.RecordToMap(rec2)

	m3 := make(map[string]gopurs_runtime.Value)

	for k, v := range m2 {
		m3[k] = v
	}
	for k, v := range m1 {
		m3[k] = v
	}

	return gopurs_runtime.Record(m3)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Record_Unsafe_Union__UnsafeUnionFn = // TAST: (ForAll [a, b, c] (ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar a), (TypeVar b), (TypeVar c)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Record_Unsafe_Union__UnsafeUnionFn(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
