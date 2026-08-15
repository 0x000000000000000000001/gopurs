package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Product gopurs_runtime.Value
var once_Main_Product sync.Once

func Get_Main_Product() gopurs_runtime.Value {
	once_Main_Product.Do(func() {
		cache_Main_Product = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2670017141, UnsafePtr: unsafe.Pointer((&Constructor_Main_Product{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Product
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqMu gopurs_runtime.Value
var once_Main_eqMu sync.Once

func Get_Main_eqMu() gopurs_runtime.Value {
	once_Main_eqMu.Do(func() {
		cache_Main_eqMu = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqMu(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Main_eqMu
}

var cache_Main_eq1Mu gopurs_runtime.Value
var once_Main_eq1Mu sync.Once

func Get_Main_eq1Mu() gopurs_runtime.Value {
	once_Main_eq1Mu.Do(func() {
		cache_Main_eq1Mu = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eq1Mu(dictEq_0_box)
		})
	})
	return cache_Main_eq1Mu
}

type Constructor_Main_Product struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Main_eqMu(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V0, (*Constructor_Main_Product)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V1, (*Constructor_Main_Product)(y_3.UnsafePtr).V1).IntVal) != (0)))
		})
	})}))}
}

func Call_Main_eq1Mu(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V0, (*Constructor_Main_Product)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Main_Product)(x_2.UnsafePtr).V1, (*Constructor_Main_Product)(y_3.UnsafePtr).V1).IntVal) != (0)))
			})
		})
	})}))}
}
