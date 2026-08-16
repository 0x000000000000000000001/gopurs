package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Const gopurs_runtime.Value
var once_Main_Const sync.Once

func Get_Main_Const() gopurs_runtime.Value {
	once_Main_Const.Do(func() {
		cache_Main_Const = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Const(x_0_box)
		})
	})
	return cache_Main_Const
}

var cache_Main_runConst gopurs_runtime.Value
var once_Main_runConst sync.Once

func Get_Main_runConst() gopurs_runtime.Value {
	once_Main_runConst.Do(func() {
		cache_Main_runConst = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runConst(v_0_box)
		})
	})
	return cache_Main_runConst
}

var cache_Main_functorConst gopurs_runtime.Value
var once_Main_functorConst sync.Once

func Get_Main_functorConst() gopurs_runtime.Value {
	once_Main_functorConst.Do(func() {
		cache_Main_functorConst = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return v1_1
			})
		})}))}
	})
	return cache_Main_functorConst
}

var cache_Main_example1 gopurs_runtime.Value
var once_Main_example1 sync.Once

func Get_Main_example1() gopurs_runtime.Value {
	once_Main_example1.Do(func() {
		cache_Main_example1 = gopurs_runtime.Str(("Do notation for") + (" Semigroup"))
	})
	return cache_Main_example1
}

var cache_Main_applySecond gopurs_runtime.Value
var once_Main_applySecond sync.Once

func Get_Main_applySecond() gopurs_runtime.Value {
	once_Main_applySecond.Do(func() {
		cache_Main_applySecond = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
		})
	})
	return cache_Main_applySecond
}

var cache_Main_applyConst gopurs_runtime.Value
var once_Main_applyConst sync.Once

func Get_Main_applyConst() gopurs_runtime.Value {
	once_Main_applyConst.Do(func() {
		cache_Main_applyConst = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_applyConst(dictSemigroup_0_box)
		})
	})
	return cache_Main_applyConst
}

var cache_Main_applyConst1 gopurs_runtime.Value
var once_Main_applyConst1 sync.Once

func Get_Main_applyConst1() gopurs_runtime.Value {
	once_Main_applyConst1.Do(func() {
		cache_Main_applyConst1 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorConst()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str((v_0.StrVal()) + (v1_1.StrVal()))
			})
		})}))}
	})
	return cache_Main_applyConst1
}

var cache_Main_example2 gopurs_runtime.Value
var once_Main_example2 sync.Once

func Get_Main_example2() gopurs_runtime.Value {
	once_Main_example2.Do(func() {
		cache_Main_example2 = gopurs_runtime.Str(("Do notation for") + (" Apply"))
	})
	return cache_Main_example2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Get_Main_example1().StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Get_Main_example2().StrVal())), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_applySecond__3603561348 gopurs_runtime.Value
var once_Main_applySecond__3603561348 sync.Once

func Get_Main_applySecond__3603561348() gopurs_runtime.Value {
	once_Main_applySecond__3603561348.Do(func() {
		cache_Main_applySecond__3603561348 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_applySecond__3603561348(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
		})
	})
	return cache_Main_applySecond__3603561348
}

var cache_Main_example2__2478757062 gopurs_runtime.Value
var once_Main_example2__2478757062 sync.Once

func Get_Main_example2__2478757062() gopurs_runtime.Value {
	once_Main_example2__2478757062.Do(func() {
		cache_Main_example2__2478757062 = gopurs_runtime.Str(("Do notation for") + (" Apply"))
	})
	return cache_Main_example2__2478757062
}

var cache_Main_functorConst__3854454365 gopurs_runtime.Value
var once_Main_functorConst__3854454365 sync.Once

func Get_Main_functorConst__3854454365() gopurs_runtime.Value {
	once_Main_functorConst__3854454365.Do(func() {
		cache_Main_functorConst__3854454365 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return v1_1
			})
		})}))}
	})
	return cache_Main_functorConst__3854454365
}

var cache_Main_runConst__3081633047 gopurs_runtime.Value
var once_Main_runConst__3081633047 sync.Once

func Get_Main_runConst__3081633047() gopurs_runtime.Value {
	once_Main_runConst__3081633047.Do(func() {
		cache_Main_runConst__3081633047 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runConst__3081633047(v_0_box)
		})
	})
	return cache_Main_runConst__3081633047
}

func Call_Main_Const(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_runConst(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_applySecond(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 shape=App(Other) expectedFromAst=*Constructor_Data_Functor_Functor actual=*Constructor_Data_Functor_Functor bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(fb_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				})
			}), fa_2), fb_3)
		})
	})
}

func Call_Main_applyConst(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorConst()))}
	}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
		})
	})}))}
}

func Call_Main_applySecond__3603561348(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 shape=App(Other) expectedFromAst=*Constructor_Data_Functor_Functor actual=*Constructor_Data_Functor_Functor bindingType=(ADT ["Data","Functor","Functor"] [Any])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(fb_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				})
			}), fa_2), fb_3)
		})
	})
}

func Call_Main_runConst__3081633047(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}
