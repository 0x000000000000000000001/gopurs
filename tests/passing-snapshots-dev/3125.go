package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_B
}

var cache_Main_memptyB gopurs_runtime.Value
var once_Main_memptyB sync.Once

func Get_Main_memptyB() gopurs_runtime.Value {
	once_Main_memptyB.Do(func() {
		cache_Main_memptyB = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_memptyB(dictMonoid_0_box)
		})
	})
	return cache_Main_memptyB
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_0_0 shape=App(Var) bindingType=(ADT ["Main","B"] [(Func [Int] (Array Unit))])
			v_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_B[gopurs_runtime.Value]](Call_Main_memptyB(Get_Data_Monoid_monoidArray()))
			_ = v_0_0
			// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_1_1 := Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqUnit()))}), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply((v_0_0).V0, gopurs_runtime.Int(int64(0)))).UnsafePtr))), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply((v_0_0).V1, gopurs_runtime.Int(int64(0)))).UnsafePtr)))).IntVal) != (0)))
			_ = __local_var_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_B[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_a
}

func Call_Main_memptyB(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
	}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
	})}))}
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
