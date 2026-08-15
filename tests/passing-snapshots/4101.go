package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ClassA_dollar_Dict gopurs_runtime.Value
var once_Main_ClassA_dollar_Dict sync.Once

func Get_Main_ClassA_dollar_Dict() gopurs_runtime.Value {
	once_Main_ClassA_dollar_Dict.Do(func() {
		cache_Main_ClassA_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ClassA_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ClassA_dollar_Dict
}

var cache_Main_ClassB_dollar_Dict gopurs_runtime.Value
var once_Main_ClassB_dollar_Dict sync.Once

func Get_Main_ClassB_dollar_Dict() gopurs_runtime.Value {
	once_Main_ClassB_dollar_Dict.Do(func() {
		cache_Main_ClassB_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ClassB_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ClassB_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_a
}

var cache_Main_b gopurs_runtime.Value
var once_Main_b sync.Once

func Get_Main_b() gopurs_runtime.Value {
	once_Main_b.Do(func() {
		cache_Main_b = gopurs_runtime.Value{Type: 9, IntVal: 1454680786, UnsafePtr: unsafe.Pointer((&Constructor_Main_ClassB{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_b
}

type Constructor_Main_ClassA struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3328906353] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ClassA)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_ClassA: " + key)
		}
	}
}

type Constructor_Main_ClassB struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1454680786] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ClassB)(ptr)
		_ = c
		switch key {
		case "ClassA0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ClassB: " + key)
		}
	}
}

func Call_Main_ClassA_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_ClassB_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
