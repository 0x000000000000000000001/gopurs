package purescript

import "gopurs/output/gopurs_runtime"

func Foreign_Object_ST_NewImpl(_ interface{}) interface{} {
	return make(map[string]interface{})
}

func Foreign_Object_ST_PeekImpl(just func(interface{}) interface{}, nothing interface{}, k string, m map[string]interface{}, _ interface{}) interface{} {
	if val, ok := m[k]; ok {
		return just(val)
	}
	return nothing
}

func Foreign_Object_ST_Poke(k string, v interface{}, m map[string]interface{}, _ interface{}) interface{} {
	m[k] = v
	return m
}

func Foreign_Object_ST_DeleteImpl(k string, m map[string]interface{}, _ interface{}) interface{} {
	delete(m, k)
	return m
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Foreign_Object_ST_DeleteImpl = // TAST: (ForAll [a, r] (Func [String, (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.UnboxObject(arg1)
	go_arg2 := arg2
	go_res := Foreign_Object_ST_DeleteImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_Object_ST_NewImpl = // TAST: (ForAll [a, r] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_Object_ST_NewImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_Object_ST_PeekImpl = // TAST: (ForAll [a, b, r] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeVar b), String, (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.UnboxObject(arg3)
	go_arg4 := arg4
	go_res := Foreign_Object_ST_PeekImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_Object_ST_Poke = // TAST: (ForAll [a, r] (Func [String, (TypeVar a), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Foreign","Object","ST","STObject"] [(TypeVar r), (TypeVar a)])])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.UnboxObject(arg2)
	go_arg3 := arg3
	go_res := Foreign_Object_ST_Poke(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
