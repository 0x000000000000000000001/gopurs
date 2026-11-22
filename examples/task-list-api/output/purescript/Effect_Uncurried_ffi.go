package purescript

import "gopurs/output/gopurs_runtime"

func Effect_Uncurried_MkEffectFn1(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn2(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func2(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn3(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func3(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn4(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func4(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn5(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func5(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn6(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func6(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn7(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func7(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn8(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func8(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), a8), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn9(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func9(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), a8), a9), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_MkEffectFn10(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func10(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value, a10 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), a8), a9), a10), gopurs_runtime.Value{})
	})
}

func Effect_Uncurried_RunEffectFn1(f gopurs_runtime.Value, a1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f, a1)
	})
}

func Effect_Uncurried_RunEffectFn2(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f, a1, a2)
	})
}

func Effect_Uncurried_RunEffectFn3(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(f, a1, a2, a3)
	})
}

func Effect_Uncurried_RunEffectFn4(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply4(f, a1, a2, a3, a4)
	})
}

func Effect_Uncurried_RunEffectFn5(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply5(f, a1, a2, a3, a4, a5)
	})
}

func Effect_Uncurried_RunEffectFn6(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply6(f, a1, a2, a3, a4, a5, a6)
	})
}

func Effect_Uncurried_RunEffectFn7(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply7(f, a1, a2, a3, a4, a5, a6, a7)
	})
}

func Effect_Uncurried_RunEffectFn8(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply8(f, a1, a2, a3, a4, a5, a6, a7, a8)
	})
}

func Effect_Uncurried_RunEffectFn9(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply9(f, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	})
}

func Effect_Uncurried_RunEffectFn10(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value, a10 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply10(f, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	})
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Effect_Uncurried_MkEffectFn1 = // TAST: (ForAll [a, r] (Func [(Func [(TypeVar a)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn1"] [(TypeVar a), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn1(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn10 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, j, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn10"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn10(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn2 = // TAST: (ForAll [a, b, r] (Func [(Func [(TypeVar a), (TypeVar b)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn2"] [(TypeVar a), (TypeVar b), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn2(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn3 = // TAST: (ForAll [a, b, c, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn3(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn4 = // TAST: (ForAll [a, b, c, d, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn4(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn5 = // TAST: (ForAll [a, b, c, d, e, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn5(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn6 = // TAST: (ForAll [a, b, c, d, e, f, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn6(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn7(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn8(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_MkEffectFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, r] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (ADT ["Effect","Effect"] [(TypeVar r)]))] (ADT ["Effect","Uncurried","EffectFn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar r)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Uncurried_MkEffectFn9(go_arg0)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn1 = // TAST: (ForAll [a, r] (Func [(ADT ["Effect","Uncurried","EffectFn1"] [(TypeVar a), (TypeVar r)]), (TypeVar a)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Effect_Uncurried_RunEffectFn1(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn10 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, j, r] (Func [(ADT ["Effect","Uncurried","EffectFn10"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)] (ADT ["Effect","Effect"] [(TypeVar r)])))
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
	go_res := Effect_Uncurried_RunEffectFn10(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9, go_arg10)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn2 = // TAST: (ForAll [a, b, r] (Func [(ADT ["Effect","Uncurried","EffectFn2"] [(TypeVar a), (TypeVar b), (TypeVar r)]), (TypeVar a), (TypeVar b)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Effect_Uncurried_RunEffectFn2(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn3 = // TAST: (ForAll [a, b, c, r] (Func [(ADT ["Effect","Uncurried","EffectFn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Effect_Uncurried_RunEffectFn3(go_arg0, go_arg1, go_arg2, go_arg3)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn4 = // TAST: (ForAll [a, b, c, d, r] (Func [(ADT ["Effect","Uncurried","EffectFn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Effect_Uncurried_RunEffectFn4(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn5 = // TAST: (ForAll [a, b, c, d, e, r] (Func [(ADT ["Effect","Uncurried","EffectFn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_res := Effect_Uncurried_RunEffectFn5(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn6 = // TAST: (ForAll [a, b, c, d, e, f, r] (Func [(ADT ["Effect","Uncurried","EffectFn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_res := Effect_Uncurried_RunEffectFn6(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, r] (Func [(ADT ["Effect","Uncurried","EffectFn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (ADT ["Effect","Effect"] [(TypeVar r)])))
gopurs_runtime.Func8(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value, arg7 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_arg7 := arg7
	go_res := Effect_Uncurried_RunEffectFn7(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, r] (Func [(ADT ["Effect","Uncurried","EffectFn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (ADT ["Effect","Effect"] [(TypeVar r)])))
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
	go_res := Effect_Uncurried_RunEffectFn8(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8)
	return go_res
})
var _Gopurs_Effect_Uncurried_RunEffectFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, r] (Func [(ADT ["Effect","Uncurried","EffectFn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar r)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (ADT ["Effect","Effect"] [(TypeVar r)])))
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
	go_res := Effect_Uncurried_RunEffectFn9(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6, go_arg7, go_arg8, go_arg9)
	return go_res
})
