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

var cache_Main_Const__3270746063 gopurs_runtime.Value
var once_Main_Const__3270746063 sync.Once

func Get_Main_Const__3270746063() gopurs_runtime.Value {
	once_Main_Const__3270746063.Do(func() {
		cache_Main_Const__3270746063 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Const__3270746063(x_0_box.StrVal())
		})
	})
	return cache_Main_Const__3270746063
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
		cache_Main_functorConst = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return v1_1
		})}))}
	})
	return cache_Main_functorConst
}

var cache_Main_example1 gopurs_runtime.Value
var once_Main_example1 sync.Once

func Get_Main_example1() gopurs_runtime.Value {
	once_Main_example1.Do(func() {
		cache_Main_example1 = gopurs_runtime.Str(("Do notation") + (" for Semigroup"))
	})
	return cache_Main_example1
}

var cache_Main_applySecond gopurs_runtime.Value
var once_Main_applySecond sync.Once

func Get_Main_applySecond() gopurs_runtime.Value {
	once_Main_applySecond.Do(func() {
		cache_Main_applySecond = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
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
		cache_Main_applyConst1 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Main_applyConst(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Main_443971153_4179793454(Rebox_Main_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))})))}
	})
	return cache_Main_applyConst1
}

var cache_Main_example2 gopurs_runtime.Value
var once_Main_example2 sync.Once

func Get_Main_example2() gopurs_runtime.Value {
	once_Main_example2.Do(func() {
		cache_Main_example2 = gopurs_runtime.Apply2(Call_Main_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Main_applyConst(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Main_443971153_4179793454(Rebox_Main_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}))), gopurs_runtime.Str("Do"), gopurs_runtime.Apply2(Call_Main_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Main_applyConst(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Main_443971153_4179793454(Rebox_Main_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}))), gopurs_runtime.Str(" notation"), gopurs_runtime.Apply2(Call_Main_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Main_applyConst(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Main_443971153_4179793454(Rebox_Main_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}))), gopurs_runtime.Str(" for"), gopurs_runtime.Str(" Apply"))))
	})
	return cache_Main_example2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Get_Main_example1().StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_runConst(Get_Main_example2()).StrVal())), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_Const(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Const__3270746063(x_0_loop string) gopurs_runtime.Value {
Const__3270746063:
	for {
		if false {
			continue Const__3270746063
		}
		var x_0 string = x_0_loop
		_ = x_0
		return gopurs_runtime.Str(x_0)
	}
}

func Call_Main_runConst(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_applySecond(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope17)])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func2(func(fa_2 gopurs_runtime.Value, fb_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
		__local_var_4_1 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
		_ = __local_var_4_1
		return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return __local_var_4_1
		}), fa_2), fb_3)
	})
}

func Call_Main_applyConst(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorConst()))}
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
	})}))}
}

func Rebox_Main_4179793454_443971153(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semigroup_Semigroup[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
