package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Pointed_dollar_Dict gopurs_runtime.Value
var once_Main_Pointed_dollar_Dict sync.Once

func Get_Main_Pointed_dollar_Dict() gopurs_runtime.Value {
	once_Main_Pointed_dollar_Dict.Do(func() {
		cache_Main_Pointed_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Pointed_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Pointed_dollar_Dict
}

var cache_Main_pointedArray gopurs_runtime.Value
var once_Main_pointedArray sync.Once

func Get_Main_pointedArray() gopurs_runtime.Value {
	once_Main_pointedArray.Do(func() {
		cache_Main_pointedArray = gopurs_runtime.Value{Type: 9, IntVal: 4236620371, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pointed{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})}))}
	})
	return cache_Main_pointedArray
}

var cache_Main_point gopurs_runtime.Value
var once_Main_point sync.Once

func Get_Main_point() gopurs_runtime.Value {
	once_Main_point.Do(func() {
		cache_Main_point = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_point(gopurs_runtime.CoerceToStruct[Constructor_Main_Pointed](dict_0_box))
		})
	})
	return cache_Main_point
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Pointed struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4236620371] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Pointed)(ptr)
		_ = c
		switch key {
		case "point":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Pointed: " + key)
		}
	}
}

func Call_Main_Pointed_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_point(dict_0_loop *Constructor_Main_Pointed) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Pointed = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
