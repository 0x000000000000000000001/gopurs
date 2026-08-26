package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_nubUnion gopurs_runtime.Value
var once_Main_nubUnion sync.Once

func Get_Main_nubUnion() gopurs_runtime.Value {
	once_Main_nubUnion.Do(func() {
		cache_Main_nubUnion = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_nubUnion(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal), uint32(v1_3_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_nubUnion
}

var cache_Main_nubUnion__1620431907 gopurs_runtime.Value
var once_Main_nubUnion__1620431907 sync.Once

func Get_Main_nubUnion__1620431907() gopurs_runtime.Value {
	once_Main_nubUnion__1620431907.Do(func() {
		cache_Main_nubUnion__1620431907 = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_nubUnion__1620431907(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal), uint32(v1_3_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_nubUnion__1620431907
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_nubUnion(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32, v1_3_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	var v1_3 uint32 = v1_3_loop
	_ = v1_3
	return 513803634
}

func Call_Main_nubUnion__1620431907(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32, v1_3_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	var v1_3 uint32 = v1_3_loop
	_ = v1_3
	return 513803634
}
