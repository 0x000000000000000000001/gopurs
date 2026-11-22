package purescript

import "gopurs/output/gopurs_runtime"

func Record_Builder_CopyRecord(r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m2 := make(map[string]gopurs_runtime.Value, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return gopurs_runtime.Record(m2)
}

func Record_Builder_UnsafeInsert(l string, a gopurs_runtime.Value, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m[l] = a
	return gopurs_runtime.Record(m)
}

func Record_Builder_UnsafeModify(l string, f gopurs_runtime.Value, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m[l] = gopurs_runtime.Apply(f, m[l])
	return gopurs_runtime.Record(m)
}

func Record_Builder_UnsafeDelete(l string, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	delete(m, l)
	return gopurs_runtime.Record(m)
}

func Record_Builder_UnsafeRename(l1 string, l2 string, r gopurs_runtime.Value) gopurs_runtime.Value {
	m := gopurs_runtime.RecordToMap(r)
	m[l2] = m[l1]
	delete(m, l1)
	return gopurs_runtime.Record(m)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Record_Builder_CopyRecord = // TAST: (ForAll [r1] (Func [(Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r1)))))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Record_Builder_CopyRecord(go_arg0)
	return go_res
})
var _Gopurs_Record_Builder_UnsafeDelete = // TAST: (ForAll [r1, r2] (Func [String, (Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r2)))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Record_Builder_UnsafeDelete(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Record_Builder_UnsafeInsert = // TAST: (ForAll [a, r1, r2] (Func [String, (TypeVar a), (Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r2)))))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Record_Builder_UnsafeInsert(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Record_Builder_UnsafeModify = // TAST: (ForAll [a, b, r1, r2] (Func [String, (Func [(TypeVar a)] (TypeVar b)), (Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r2)))))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Record_Builder_UnsafeModify(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Record_Builder_UnsafeRename = // TAST: (ForAll [r1, r2] (Func [String, String, (Record (Row [] (TypeVar r1)))] (Record (Row [] (TypeVar r2)))))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := arg2
	go_res := Record_Builder_UnsafeRename(go_arg0, go_arg1, go_arg2)
	return go_res
})
