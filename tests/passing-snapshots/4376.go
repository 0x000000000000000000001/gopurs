package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_v1 gopurs_runtime.Value
var once_Main_v1 sync.Once

func Get_Main_v1() gopurs_runtime.Value {
	once_Main_v1.Do(func() {
		cache_Main_v1 = gopurs_runtime.RecordDict1("a", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()})})
	})
	return cache_Main_v1
}

var cache_Main_v2 gopurs_runtime.Value
var once_Main_v2 sync.Once

func Get_Main_v2() gopurs_runtime.Value {
	once_Main_v2.Do(func() {
		cache_Main_v2 = func() gopurs_runtime.Value {
			origVal := Get_Main_v1()
			if origVal.Type != gopurs_runtime.TypeRecord1 {
				return gopurs_runtime.RecordUpdateDict(origVal, []string{"a"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})
			}
			clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
			clone.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
			return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
		}()
	})
	return cache_Main_v2
}

var cache_Main_union gopurs_runtime.Value
var once_Main_union sync.Once

func Get_Main_union() gopurs_runtime.Value {
	once_Main_union.Do(func() {
		cache_Main_union = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v3_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_union(_dollar___unused_0_box, v_1_box, v3_2_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_union
}

var cache_Main_shouldSolve gopurs_runtime.Value
var once_Main_shouldSolve sync.Once

func Get_Main_shouldSolve() gopurs_runtime.Value {
	once_Main_shouldSolve.Do(func() {
		cache_Main_shouldSolve = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_shouldSolve
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_asNothing gopurs_runtime.Value
var once_Main_asNothing sync.Once

func Get_Main_asNothing() gopurs_runtime.Value {
	once_Main_asNothing.Do(func() {
		cache_Main_asNothing = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_asNothing(v_0_box)
		})
	})
	return cache_Main_asNothing
}

var cache_Main_union__3545286655 gopurs_runtime.Value
var once_Main_union__3545286655 sync.Once

func Get_Main_union__3545286655() gopurs_runtime.Value {
	once_Main_union__3545286655.Do(func() {
		cache_Main_union__3545286655 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v3_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_union__3545286655(v_0_box, v3_1_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_union__3545286655
}

var cache_Main_union__2976016889 gopurs_runtime.Value
var once_Main_union__2976016889 sync.Once

func Get_Main_union__2976016889() gopurs_runtime.Value {
	once_Main_union__2976016889.Do(func() {
		cache_Main_union__2976016889 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v3_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_union__2976016889(_dollar___unused_0_box, v_1_box, v3_2_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_union__2976016889
}

func Call_Main_union(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v3_2_loop gopurs_runtime.Value) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var v3_2 gopurs_runtime.Value = v3_2_loop
	_ = v3_2
	return 513803634
}

func Call_Main_asNothing(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return func() gopurs_runtime.Value {
		origVal := v_0
		if origVal.Type != gopurs_runtime.TypeRecord1 {
			return gopurs_runtime.RecordUpdateDict(origVal, []string{"a"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}})
		}
		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
		clone.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	}()
}

func Call_Main_union__3545286655(v_0_loop gopurs_runtime.Value, v3_1_loop gopurs_runtime.Value) uint32 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v3_1 gopurs_runtime.Value = v3_1_loop
	_ = v3_1
	return 513803634
}

func Call_Main_union__2976016889(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v3_2_loop gopurs_runtime.Value) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var v3_2 gopurs_runtime.Value = v3_2_loop
	_ = v3_2
	return 513803634
}
