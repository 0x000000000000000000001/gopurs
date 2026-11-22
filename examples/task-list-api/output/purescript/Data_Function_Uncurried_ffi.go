package purescript

import "gopurs/output/gopurs_runtime"

func identity(fn any) any {
	return fn
}

var Data_Function_Uncurried_RunFn2 = identity
var Data_Function_Uncurried_RunFn3 = identity
var Data_Function_Uncurried_RunFn4 = identity
var Data_Function_Uncurried_RunFn5 = identity
var Data_Function_Uncurried_RunFn6 = identity
var Data_Function_Uncurried_RunFn7 = identity
var Data_Function_Uncurried_RunFn8 = identity
var Data_Function_Uncurried_RunFn9 = identity
var Data_Function_Uncurried_RunFn10 = identity

var Data_Function_Uncurried_MkFn2 = identity
var Data_Function_Uncurried_MkFn3 = identity
var Data_Function_Uncurried_MkFn4 = identity
var Data_Function_Uncurried_MkFn5 = identity
var Data_Function_Uncurried_MkFn6 = identity
var Data_Function_Uncurried_MkFn7 = identity
var Data_Function_Uncurried_MkFn8 = identity
var Data_Function_Uncurried_MkFn9 = identity
var Data_Function_Uncurried_MkFn10 = identity

func Data_Function_Uncurried_MkFn0(f any) any {
	return f
}

func Data_Function_Uncurried_RunFn0(f func(any) any) any {
	return f(nil)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Function_Uncurried_MkFn0 = // TAST: (ForAll [a] (Func [(Func [Unit] (TypeVar a))] (ADT ["Data","Function","Uncurried","Fn0"] [(TypeVar a)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Function_Uncurried_MkFn0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Function_Uncurried_MkFn10 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, j, k] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)] (TypeVar k))] (ADT ["Data","Function","Uncurried","Fn10"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j), (TypeVar k)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn10)
var _Gopurs_Data_Function_Uncurried_MkFn2 = // TAST: (ForAll [a, b, c] (Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar c))] (ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar a), (TypeVar b), (TypeVar c)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn2)
var _Gopurs_Data_Function_Uncurried_MkFn3 = // TAST: (ForAll [a, b, c, d] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c)] (TypeVar d))] (ADT ["Data","Function","Uncurried","Fn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn3)
var _Gopurs_Data_Function_Uncurried_MkFn4 = // TAST: (ForAll [a, b, c, d, e] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (TypeVar e))] (ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn4)
var _Gopurs_Data_Function_Uncurried_MkFn5 = // TAST: (ForAll [a, b, c, d, e, f] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (TypeVar f))] (ADT ["Data","Function","Uncurried","Fn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn5)
var _Gopurs_Data_Function_Uncurried_MkFn6 = // TAST: (ForAll [a, b, c, d, e, f, g] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (TypeVar g))] (ADT ["Data","Function","Uncurried","Fn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn6)
var _Gopurs_Data_Function_Uncurried_MkFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, h] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (TypeVar h))] (ADT ["Data","Function","Uncurried","Fn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn7)
var _Gopurs_Data_Function_Uncurried_MkFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (TypeVar i))] (ADT ["Data","Function","Uncurried","Fn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn8)
var _Gopurs_Data_Function_Uncurried_MkFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, j] (Func [(Func [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (TypeVar j))] (ADT ["Data","Function","Uncurried","Fn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)])))
gopurs_runtime.Box(Data_Function_Uncurried_MkFn9)
var _Gopurs_Data_Function_Uncurried_RunFn0 = // TAST: (ForAll [a] (Func [(ADT ["Data","Function","Uncurried","Fn0"] [(TypeVar a)])] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_res := Data_Function_Uncurried_RunFn0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Function_Uncurried_RunFn10 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, j, k] (Func [(ADT ["Data","Function","Uncurried","Fn10"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j), (TypeVar k)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)] (TypeVar k)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn10)
var _Gopurs_Data_Function_Uncurried_RunFn2 = // TAST: (ForAll [a, b, c] (Func [(ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar a), (TypeVar b), (TypeVar c)]), (TypeVar a), (TypeVar b)] (TypeVar c)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn2)
var _Gopurs_Data_Function_Uncurried_RunFn3 = // TAST: (ForAll [a, b, c, d] (Func [(ADT ["Data","Function","Uncurried","Fn3"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)]), (TypeVar a), (TypeVar b), (TypeVar c)] (TypeVar d)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn3)
var _Gopurs_Data_Function_Uncurried_RunFn4 = // TAST: (ForAll [a, b, c, d, e] (Func [(ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d)] (TypeVar e)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn4)
var _Gopurs_Data_Function_Uncurried_RunFn5 = // TAST: (ForAll [a, b, c, d, e, f] (Func [(ADT ["Data","Function","Uncurried","Fn5"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e)] (TypeVar f)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn5)
var _Gopurs_Data_Function_Uncurried_RunFn6 = // TAST: (ForAll [a, b, c, d, e, f, g] (Func [(ADT ["Data","Function","Uncurried","Fn6"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f)] (TypeVar g)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn6)
var _Gopurs_Data_Function_Uncurried_RunFn7 = // TAST: (ForAll [a, b, c, d, e, f, g, h] (Func [(ADT ["Data","Function","Uncurried","Fn7"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g)] (TypeVar h)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn7)
var _Gopurs_Data_Function_Uncurried_RunFn8 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i] (Func [(ADT ["Data","Function","Uncurried","Fn8"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h)] (TypeVar i)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn8)
var _Gopurs_Data_Function_Uncurried_RunFn9 = // TAST: (ForAll [a, b, c, d, e, f, g, h, i, j] (Func [(ADT ["Data","Function","Uncurried","Fn9"] [(TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i), (TypeVar j)]), (TypeVar a), (TypeVar b), (TypeVar c), (TypeVar d), (TypeVar e), (TypeVar f), (TypeVar g), (TypeVar h), (TypeVar i)] (TypeVar j)))
gopurs_runtime.Box(Data_Function_Uncurried_RunFn9)
