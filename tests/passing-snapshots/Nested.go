package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Extend gopurs_runtime.Value
var once_Main_Extend sync.Once

func Get_Main_Extend() gopurs_runtime.Value {
	once_Main_Extend.Do(func() {
		cache_Main_Extend = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Extend
}

var cache_Main_Square gopurs_runtime.Value
var once_Main_Square sync.Once

func Get_Main_Square() gopurs_runtime.Value {
	once_Main_Square.Do(func() {
		cache_Main_Square = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3516456831, UnsafePtr: unsafe.Pointer((&Constructor_Main_Square[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Square
}

var cache_Main_Bigger gopurs_runtime.Value
var once_Main_Bigger sync.Once

func Get_Main_Bigger() gopurs_runtime.Value {
	once_Main_Bigger.Do(func() {
		cache_Main_Bigger = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1523285570, UnsafePtr: unsafe.Pointer((&Constructor_Main_Bigger[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Bigger
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Extend[T_r any, T_a any] struct {
	Rc uint32
	V0 *struct {
		next gopurs_runtime.Value
		prev gopurs_runtime.Value
	}
}

type Constructor_Main_Square[T_r any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Bigger[T_r any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}
