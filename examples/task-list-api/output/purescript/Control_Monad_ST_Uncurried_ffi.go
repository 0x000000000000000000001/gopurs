//gopurs:missing-ffi-v1:{"binding":"_Gopurs_Control_Monad_ST_Uncurried_MkSTFn10","module":"Control.Monad.ST.Uncurried","path":"gopurs-st/src/Control/Monad/ST/Uncurried.go","symbol":"mkSTFn10"}
//gopurs:missing-ffi-v1:{"binding":"_Gopurs_Control_Monad_ST_Uncurried_RunSTFn10","module":"Control.Monad.ST.Uncurried","path":"gopurs-st/src/Control/Monad/ST/Uncurried.go","symbol":"runSTFn10"}
package purescript

import "gopurs/output/gopurs_runtime"

// STFn values are ordinary gopurs_runtime.Value functions; the foreign bridge
// passes them without adapters, so runSTFn{N} applies them directly.

func Control_Monad_ST_Uncurried_MkSTFn1(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn2(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn3(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn4(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn5(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn6(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn7(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn8(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}
func Control_Monad_ST_Uncurried_MkSTFn9(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return fn
}

func Control_Monad_ST_Uncurried_RunSTFn1(fn gopurs_runtime.Value, a gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(fn, a)
}
func Control_Monad_ST_Uncurried_RunSTFn2(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply2(fn, a, b)
}
func Control_Monad_ST_Uncurried_RunSTFn3(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply3(fn, a, b, c)
}
func Control_Monad_ST_Uncurried_RunSTFn4(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply4(fn, a, b, c, d)
}
func Control_Monad_ST_Uncurried_RunSTFn5(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply5(fn, a, b, c, d, e)
}
func Control_Monad_ST_Uncurried_RunSTFn6(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, f gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply6(fn, a, b, c, d, e, f)
}
func Control_Monad_ST_Uncurried_RunSTFn7(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, f gopurs_runtime.Value, g gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply7(fn, a, b, c, d, e, f, g)
}
func Control_Monad_ST_Uncurried_RunSTFn8(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, f gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply8(fn, a, b, c, d, e, f, g, h)
}
func Control_Monad_ST_Uncurried_RunSTFn9(fn gopurs_runtime.Value, a gopurs_runtime.Value, b gopurs_runtime.Value, c gopurs_runtime.Value, d gopurs_runtime.Value, e gopurs_runtime.Value, f gopurs_runtime.Value, g gopurs_runtime.Value, h gopurs_runtime.Value, i gopurs_runtime.Value, _ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply9(fn, a, b, c, d, e, f, g, h, i)
}

// STFn10 and runSTFn10 are not implemented in this runtime; the bridge keeps
// its not-implemented fallback for them.

// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn1 = // TAST: (ForAll [a, t, r] (Func [(Func [(TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn1"] [(TypeVar a), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn1(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	panic("FFI not implemented: mkSTFn10")
	return gopurs_runtime.Value{}
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn2 = // TAST: (ForAll [a, b, t, r] (Func [(Func [(TypeVar a), (TypeVar b)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn2"] [(TypeVar a), (TypeVar b), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn2(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn3 = // TAST: (ForAll [a, b, c, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn3(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn4 = // TAST: (ForAll [a, b, c, d, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn4(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn5 = // TAST: (ForAll [a, b, c, d, e, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn5(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn6 = // TAST: (ForAll [a, b, c, d, e, f, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn6(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn7(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn8(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_MkSTFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, t, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)]))] (ADT ["Control","Monad","ST","Uncurried","STFn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar t), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Control_Monad_ST_Uncurried_MkSTFn9(go_arg0)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn1 = // TAST: (ForAll [a, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn1"] [(TypeVar a), (TypeVar t), (TypeVar r)]), (TypeVar a)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Control_Monad_ST_Uncurried_RunSTFn1(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn10 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	panic("FFI not implemented: runSTFn10")
	return gopurs_runtime.Value{}
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn2 = // TAST: (ForAll [a, b, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn2"] [(TypeVar a), (TypeVar b), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Control_Monad_ST_Uncurried_RunSTFn2(go_arg0, go_arg1, go_arg2, go_arg3)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn3 = // TAST: (ForAll [a, b, c, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Control_Monad_ST_Uncurried_RunSTFn3(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn4 = // TAST: (ForAll [a, b, c, d, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_res := Control_Monad_ST_Uncurried_RunSTFn4(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn5 = // TAST: (ForAll [a, b, c, d, e, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_res := Control_Monad_ST_Uncurried_RunSTFn5(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn6 = // TAST: (ForAll [a, b, c, d, e, f, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_res := Control_Monad_ST_Uncurried_RunSTFn6(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func9(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_res := Control_Monad_ST_Uncurried_RunSTFn7(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func10(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_arg9 := arg9
	go_res := Control_Monad_ST_Uncurried_RunSTFn8(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9)
	return go_res
})
var _Gopurs_Control_Monad_ST_Uncurried_RunSTFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, t, r] (Func [(ADT ["Control","Monad","ST","Uncurried","STFn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar t), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar t), (TypeVar r)])))
gopurs_runtime.Func11(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value, arg8 gopurs_runtime.Value, arg9 gopurs_runtime.Value, arg10 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_arg8 := arg8
	go_arg9 := arg9
	go_arg10 := arg10
	go_res := Control_Monad_ST_Uncurried_RunSTFn9(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9, go_arg10)
	return go_res
})
