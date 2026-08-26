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

var cache_Main_CtorKind_dollar_Dict gopurs_runtime.Value
var once_Main_CtorKind_dollar_Dict sync.Once

func Get_Main_CtorKind_dollar_Dict() gopurs_runtime.Value {
	once_Main_CtorKind_dollar_Dict.Do(func() {
		cache_Main_CtorKind_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_CtorKind_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_CtorKind_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_ctorKind1 gopurs_runtime.Value
var once_Main_ctorKind1 sync.Once

func Get_Main_ctorKind1() gopurs_runtime.Value {
	once_Main_ctorKind1.Do(func() {
		cache_Main_ctorKind1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_ctorKind1
}

var cache_Main_ctorKind0 gopurs_runtime.Value
var once_Main_ctorKind0 sync.Once

func Get_Main_ctorKind0() gopurs_runtime.Value {
	once_Main_ctorKind0.Do(func() {
		cache_Main_ctorKind0 = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ctorKind0(_dollar___unused_0_box)
		})
	})
	return cache_Main_ctorKind0
}

var cache_Main_ctorKind gopurs_runtime.Value
var once_Main_ctorKind sync.Once

func Get_Main_ctorKind() gopurs_runtime.Value {
	once_Main_ctorKind.Do(func() {
		cache_Main_ctorKind = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_ctorKind(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_ctorKind
}

var cache_Main_ctorKind__1288446513 gopurs_runtime.Value
var once_Main_ctorKind__1288446513 sync.Once

func Get_Main_ctorKind__1288446513() gopurs_runtime.Value {
	once_Main_ctorKind__1288446513.Do(func() {
		cache_Main_ctorKind__1288446513 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_ctorKind__1288446513(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_ctorKind__1288446513
}

var cache_Main_testCtor1 gopurs_runtime.Value
var once_Main_testCtor1 sync.Once

func Get_Main_testCtor1() gopurs_runtime.Value {
	once_Main_testCtor1.Do(func() {
		cache_Main_testCtor1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testCtor1
}

var cache_Main_testCtor2 gopurs_runtime.Value
var once_Main_testCtor2 sync.Once

func Get_Main_testCtor2() gopurs_runtime.Value {
	once_Main_testCtor2.Do(func() {
		cache_Main_testCtor2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testCtor2
}

var cache_Main_testCtor3 gopurs_runtime.Value
var once_Main_testCtor3 sync.Once

func Get_Main_testCtor3() gopurs_runtime.Value {
	once_Main_testCtor3.Do(func() {
		cache_Main_testCtor3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_testCtor3
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_CtorKind[T_ctor any, T_kind any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[133020252] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_CtorKind[any, any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_CtorKind: " + key)
		}
	}
}

func Call_Main_CtorKind_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_ctorKind0(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
}

func Call_Main_ctorKind(_dollar___unused_0_loop uint32, v_1_loop uint32) uint32 {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 227768594
}

func Call_Main_ctorKind__1288446513(_dollar___unused_0_loop uint32, v_1_loop uint32) uint32 {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 227768594
}
