package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_lacksX gopurs_runtime.Value
var once_Main_lacksX sync.Once

func Get_Main_lacksX() gopurs_runtime.Value {
	once_Main_lacksX.Do(func() {
		cache_Main_lacksX = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lacksX(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lacksX
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lacksX(gopurs_runtime.Value{}, 513803634)), UnsafePtr: nil}
	})
	return cache_Main_test1
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_test2(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_test2
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_test2(gopurs_runtime.Value{}, 513803634)), UnsafePtr: nil}
	})
	return cache_Main_test3
}

var cache_Main_lacksSym gopurs_runtime.Value
var once_Main_lacksSym sync.Once

func Get_Main_lacksSym() gopurs_runtime.Value {
	once_Main_lacksSym.Do(func() {
		cache_Main_lacksSym = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lacksSym(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lacksSym
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Apply(Get_Main_lacksSym(), gopurs_runtime.Value{})
	})
	return cache_Main_test4
}

func Call_Main_lacksX(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 513803634
}

func Call_Main_test2(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return Call_Main_lacksX(gopurs_runtime.Value{}, 513803634)
}

func Call_Main_lacksSym(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 513803634
}
