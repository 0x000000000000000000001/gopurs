package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_In gopurs_runtime.Value
var once_Main_In sync.Once

func Get_Main_In() gopurs_runtime.Value {
	once_Main_In.Do(func() {
		cache_Main_In = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_In(x_0_box)
		})
	})
	return cache_Main_In
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
		cache_Main_eqMu = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqMu(dictEq1_0_box)
		})
	})
	return cache_Main_eqMu
}

var cache_Main_ordMu gopurs_runtime.Value
var once_Main_ordMu sync.Once

func Get_Main_ordMu() gopurs_runtime.Value {
	once_Main_ordMu.Do(func() {
		cache_Main_ordMu = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ordMu(dictOrd1_0_box)
		})
	})
	return cache_Main_ordMu
}

func Call_Main_In(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_eqMu(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
eqMu:
	for {
		if false {
			continue eqMu
		}
		var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
		_ = dictEq1_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Call_Main_eqMu(dictEq1_0)))}, x_1, y_2).IntVal) != (0))
			})
		})}))}
	}
}

func Call_Main_ordMu(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
ordMu:
	for {
		if false {
			continue ordMu
		}
		var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
		_ = dictOrd1_0
		// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
		__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
		_ = __local_var_1_1
		// TAST (Let): eqMu1_1_0 -> *Constructor_Data_Eq_Eq
		eqMu1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Call_Main_eqMu(__local_var_1_1)))}, x_2, y_3).IntVal) != (0))
			})
		})))
		_ = eqMu1_1_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMu1_1_0)}
		}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Call_Main_ordMu(dictOrd1_0)))}, x_2, y_3).IntVal)), UnsafePtr: nil}
			})
		})}))}
	}
}
