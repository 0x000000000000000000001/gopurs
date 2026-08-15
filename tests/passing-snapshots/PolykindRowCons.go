package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_Identity gopurs_runtime.Value
var once_Main_Identity sync.Once

func Get_Main_Identity() gopurs_runtime.Value {
	once_Main_Identity.Do(func() {
		cache_Main_Identity = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 390884680, UnsafePtr: unsafe.Pointer((&Constructor_Main_Identity{1, value0}))}
		})
	})
	return cache_Main_Identity
}

var cache_Main_App gopurs_runtime.Value
var once_Main_App sync.Once

func Get_Main_App() gopurs_runtime.Value {
	once_Main_App.Do(func() {
		cache_Main_App = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4100404127, UnsafePtr: unsafe.Pointer((&Constructor_Main_App{1, value0}))}
		})
	})
	return cache_Main_App
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_lookup gopurs_runtime.Value
var once_Main_lookup sync.Once

func Get_Main_lookup() gopurs_runtime.Value {
	once_Main_lookup.Do(func() {
		cache_Main_lookup = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup
}

var cache_Main_lookup1 gopurs_runtime.Value
var once_Main_lookup1 sync.Once

func Get_Main_lookup1() gopurs_runtime.Value {
	once_Main_lookup1.Do(func() {
		cache_Main_lookup1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup1
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test1
}

var cache_Main_lookup2 gopurs_runtime.Value
var once_Main_lookup2 sync.Once

func Get_Main_lookup2() gopurs_runtime.Value {
	once_Main_lookup2.Do(func() {
		cache_Main_lookup2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup2
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test2
}

var cache_Main_lookup3 gopurs_runtime.Value
var once_Main_lookup3 sync.Once

func Get_Main_lookup3() gopurs_runtime.Value {
	once_Main_lookup3.Do(func() {
		cache_Main_lookup3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup3
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test3
}

var cache_Main_lookup4 gopurs_runtime.Value
var once_Main_lookup4 sync.Once

func Get_Main_lookup4() gopurs_runtime.Value {
	once_Main_lookup4.Do(func() {
		cache_Main_lookup4 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup4
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test4
}

var cache_Main_lookup5 gopurs_runtime.Value
var once_Main_lookup5 sync.Once

func Get_Main_lookup5() gopurs_runtime.Value {
	once_Main_lookup5.Do(func() {
		cache_Main_lookup5 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup5
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test5
}

var cache_Main_lookup6 gopurs_runtime.Value
var once_Main_lookup6 sync.Once

func Get_Main_lookup6() gopurs_runtime.Value {
	once_Main_lookup6.Do(func() {
		cache_Main_lookup6 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup6
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test6
}

var cache_Main_lookup7 gopurs_runtime.Value
var once_Main_lookup7 sync.Once

func Get_Main_lookup7() gopurs_runtime.Value {
	once_Main_lookup7.Do(func() {
		cache_Main_lookup7 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup7
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test7
}

var cache_Main_lookup8 gopurs_runtime.Value
var once_Main_lookup8 sync.Once

func Get_Main_lookup8() gopurs_runtime.Value {
	once_Main_lookup8.Do(func() {
		cache_Main_lookup8 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup8
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test8
}

var cache_Main_lookup9 gopurs_runtime.Value
var once_Main_lookup9 sync.Once

func Get_Main_lookup9() gopurs_runtime.Value {
	once_Main_lookup9.Do(func() {
		cache_Main_lookup9 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_lookup9
}

var cache_Main_test9 gopurs_runtime.Value
var once_Main_test9 sync.Once

func Get_Main_test9() gopurs_runtime.Value {
	once_Main_test9.Do(func() {
		cache_Main_test9 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test9
}

var cache_Main_lookup__2933592557 gopurs_runtime.Value
var once_Main_lookup__2933592557 sync.Once

func Get_Main_lookup__2933592557() gopurs_runtime.Value {
	once_Main_lookup__2933592557.Do(func() {
		cache_Main_lookup__2933592557 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__2933592557(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__2933592557
}

var cache_Main_lookup__1504059307 gopurs_runtime.Value
var once_Main_lookup__1504059307 sync.Once

func Get_Main_lookup__1504059307() gopurs_runtime.Value {
	once_Main_lookup__1504059307.Do(func() {
		cache_Main_lookup__1504059307 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__1504059307(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__1504059307
}

type Constructor_Main_Proxy struct {
	Rc uint32
}

type Constructor_Main_Identity struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_App struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_lookup(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_lookup__2933592557(v_0_loop uint32, v1_1_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 uint32 = v1_1_loop
	_ = v1_1
	return 227768594
}

func Call_Main_lookup__1504059307(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}
