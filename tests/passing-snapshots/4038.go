package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_A_dollar_Dict gopurs_runtime.Value
var once_Main_A_dollar_Dict sync.Once

func Get_Main_A_dollar_Dict() gopurs_runtime.Value {
	once_Main_A_dollar_Dict.Do(func() {
		cache_Main_A_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_A_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_A_dollar_Dict
}

var cache_Main_B_dollar_Dict gopurs_runtime.Value
var once_Main_B_dollar_Dict sync.Once

func Get_Main_B_dollar_Dict() gopurs_runtime.Value {
	once_Main_B_dollar_Dict.Do(func() {
		cache_Main_B_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_B_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_B_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_A struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[4219254943] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_A)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_A: " + key)
		}
	}
}

type Constructor_Main_B struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4250879068] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_B)(ptr)
		_ = c
		switch key {
		case "A0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_B: " + key)
		}
	}
}

func Call_Main_A_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_B_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
