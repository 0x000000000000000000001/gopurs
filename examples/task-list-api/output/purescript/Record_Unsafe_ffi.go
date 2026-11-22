package purescript

import "gopurs/output/gopurs_runtime"

func Record_Unsafe_UnsafeHas(label string, recVal interface{}) bool {
	v := recVal.(gopurs_runtime.Value)
	m := gopurs_runtime.RecordToMap(v)
	_, ok := m[label]
	return ok
}

func Record_Unsafe_UnsafeGet(label string, recVal interface{}) interface{} {
	v := recVal.(gopurs_runtime.Value)
	return gopurs_runtime.RecordGet(v, label)
}

func Record_Unsafe_UnsafeSet(label string, value interface{}, recVal interface{}) interface{} {
	v := recVal.(gopurs_runtime.Value)
	val := value.(gopurs_runtime.Value)
	return gopurs_runtime.RecordSet(v, label, val)
}

func Record_Unsafe_UnsafeDelete(label string, recVal interface{}) interface{} {
	v := recVal.(gopurs_runtime.Value)
	m := gopurs_runtime.RecordToMap(v)
	newMap := make(map[string]gopurs_runtime.Value)
	for k, val := range m {
		if k != label {
			newMap[k] = val
		}
	}
	return gopurs_runtime.Record(newMap)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Record_Unsafe_UnsafeDelete = // TAST: (ForAll [r1, r2] (Func [String, (Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r2)))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Record_Unsafe_UnsafeDelete(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Record_Unsafe_UnsafeGet = // TAST: (ForAll [r, a] (Func [String, (Record (Row [] (TypeVar r)))] (TypeVar a)))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Record_Unsafe_UnsafeGet(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Record_Unsafe_UnsafeHas = // TAST: (ForAll [r1] (Func [String, (Record (Row [] (TypeVar r1)))] Boolean))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Record_Unsafe_UnsafeHas(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Record_Unsafe_UnsafeSet = // TAST: (ForAll [r1, r2, a] (Func [String, (TypeVar a), (Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r2)))))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Record_Unsafe_UnsafeSet(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
