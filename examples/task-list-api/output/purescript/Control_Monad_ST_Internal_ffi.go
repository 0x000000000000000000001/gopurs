package purescript

import "gopurs/output/gopurs_runtime"

// The ST primitives operate on gopurs_runtime.Value functions directly. The
// foreign bridge passes such arguments through without building any-adapters
// or boxing, which matters for the monadic steps of generated ST loops.

var unit = gopurs_runtime.Value{}

func Control_Monad_ST_Internal_Map_(f gopurs_runtime.Value, a gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(f, gopurs_runtime.Apply(a, unit))
}

func Control_Monad_ST_Internal_Pure_(a gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return a
}

func Control_Monad_ST_Internal_Bind_(a gopurs_runtime.Value, f gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(gopurs_runtime.Apply(f, gopurs_runtime.Apply(a, unit)), unit)
}

func Control_Monad_ST_Internal_Run(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(f, unit)
}

func Control_Monad_ST_Internal_While(f gopurs_runtime.Value, a gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	for gopurs_runtime.Apply(f, unit).BoolVal() {
		gopurs_runtime.Apply(a, unit)
	}
	return unit
}

func Control_Monad_ST_Internal_ForImpl(lo int64, hi int64, f gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	for i := lo; i < hi; i++ {
		gopurs_runtime.Apply(gopurs_runtime.Apply(f, gopurs_runtime.Int(i)), unit)
	}
	return unit
}

func Control_Monad_ST_Internal_Foreach(as gopurs_runtime.Value, f gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	for _, item := range *(*[]gopurs_runtime.Value)(as.UnsafePtr) {
		gopurs_runtime.Apply(gopurs_runtime.Apply(f, item), unit)
	}
	return unit
}

func Control_Monad_ST_Internal_NewImpl(val gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.NewSTRef(val)
}

func Control_Monad_ST_Internal_Read(ref gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.ReadSTRef(ref)
}

func Control_Monad_ST_Internal_ModifyImpl(f gopurs_runtime.Value, ref gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	updated := gopurs_runtime.Apply(f, gopurs_runtime.ReadSTRef(ref))
	gopurs_runtime.WriteSTRef(ref, gopurs_runtime.RecordGet(updated, "state"))
	return gopurs_runtime.RecordGet(updated, "value")
}

func Control_Monad_ST_Internal_Write(a gopurs_runtime.Value, ref gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.WriteSTRef(ref, a)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Monad_ST_Internal_Bind_ = // TAST: (ForAll [r, a, b] (Func [(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)]), (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Internal_Bind_(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_ForImpl = // TAST: (ForAll [r, a] (Func [Int, Int, (Func [Int] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Control_Monad_ST_Internal_ForImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_Foreach = // TAST: (ForAll [r, a] (Func [(Array (TypeVar a)), (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit]))] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Internal_Foreach(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_Map_ = // TAST: (ForAll [r, a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Internal_Map_(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_ModifyImpl = // TAST: (ForAll [r, a, b] (Func [(Func [(TypeVar a)] (Record (Row [state: (TypeVar a), value: (TypeVar b)] Empty))), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar b)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Internal_ModifyImpl(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_NewImpl = // TAST: (ForAll [a, r] (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Control_Monad_ST_Internal_NewImpl(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_Pure_ = // TAST: (ForAll [r, a] (Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Control_Monad_ST_Internal_Pure_(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_Read = // TAST: (ForAll [a, r] (Func [(ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Control_Monad_ST_Internal_Read(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_Run = // TAST: (ForAll [a] (Func [(ForAll [r] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)]))] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Internal_Run(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_While = // TAST: (ForAll [r, a] (Func [(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Boolean]), (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), Unit])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Internal_While(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Control_Monad_ST_Internal_Write = // TAST: (ForAll [a, r] (Func [(TypeVar a), (ADT ["Control","Monad","ST","Internal","STRef"] [(TypeVar r), (TypeVar a)])] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (TypeVar a)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Internal_Write(go_arg0, go_arg1, go_arg2)
	return go_res
})
