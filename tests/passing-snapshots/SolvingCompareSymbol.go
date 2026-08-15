package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_symB gopurs_runtime.Value
var once_Main_symB sync.Once

func Get_Main_symB() gopurs_runtime.Value {
	once_Main_symB.Do(func() {
		cache_Main_symB = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_symB
}

var cache_Main_symA gopurs_runtime.Value
var once_Main_symA sync.Once

func Get_Main_symA() gopurs_runtime.Value {
	once_Main_symA.Do(func() {
		cache_Main_symA = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_symA
}

var cache_Main_egLT gopurs_runtime.Value
var once_Main_egLT sync.Once

func Get_Main_egLT() gopurs_runtime.Value {
	once_Main_egLT.Do(func() {
		cache_Main_egLT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_egLT
}

var cache_Main_egGT gopurs_runtime.Value
var once_Main_egGT sync.Once

func Get_Main_egGT() gopurs_runtime.Value {
	once_Main_egGT.Do(func() {
		cache_Main_egGT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_egGT
}

var cache_Main_egEQ gopurs_runtime.Value
var once_Main_egEQ sync.Once

func Get_Main_egEQ() gopurs_runtime.Value {
	once_Main_egEQ.Do(func() {
		cache_Main_egEQ = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_egEQ
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := Get_Data_Unit_unit()
			_ = _dollar___unused_0_0
			_dollar___unused_1_1 := Get_Data_Unit_unit()
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := Get_Data_Unit_unit()
			_ = _dollar___unused_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
