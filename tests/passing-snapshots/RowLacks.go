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
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
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
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
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
		cache_Main_test4 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_test4(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_test4
}

var cache_Main_lacksX__2190080542 gopurs_runtime.Value
var once_Main_lacksX__2190080542 sync.Once

func Get_Main_lacksX__2190080542() gopurs_runtime.Value {
	once_Main_lacksX__2190080542.Do(func() {
		cache_Main_lacksX__2190080542 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lacksX__2190080542(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lacksX__2190080542
}

var cache_Main_lacksX__2300874310 gopurs_runtime.Value
var once_Main_lacksX__2300874310 sync.Once

func Get_Main_lacksX__2300874310() gopurs_runtime.Value {
	once_Main_lacksX__2300874310.Do(func() {
		cache_Main_lacksX__2300874310 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lacksX__2300874310(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lacksX__2300874310
}

var cache_Main_lacksX__3616945734 gopurs_runtime.Value
var once_Main_lacksX__3616945734 sync.Once

func Get_Main_lacksX__3616945734() gopurs_runtime.Value {
	once_Main_lacksX__3616945734.Do(func() {
		cache_Main_lacksX__3616945734 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lacksX__3616945734(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lacksX__3616945734
}

var cache_Main_test2__2538753484 gopurs_runtime.Value
var once_Main_test2__2538753484 sync.Once

func Get_Main_test2__2538753484() gopurs_runtime.Value {
	once_Main_test2__2538753484.Do(func() {
		cache_Main_test2__2538753484 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_test2__2538753484(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_test2__2538753484
}

var cache_Main_test2__3443127756 gopurs_runtime.Value
var once_Main_test2__3443127756 sync.Once

func Get_Main_test2__3443127756() gopurs_runtime.Value {
	once_Main_test2__3443127756.Do(func() {
		cache_Main_test2__3443127756 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_test2__3443127756(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_test2__3443127756
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
	return 513803634
}

func Call_Main_lacksSym(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 513803634
}

func Call_Main_test4(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_lacksX__2190080542(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 513803634
}

func Call_Main_lacksX__2300874310(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_lacksX__3616945734(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 513803634
}

func Call_Main_test2__2538753484(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_test2__3443127756(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 513803634
}
