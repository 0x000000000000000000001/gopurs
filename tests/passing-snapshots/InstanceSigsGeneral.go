package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Eq_dollar_Dict gopurs_runtime.Value
var once_Main_Eq_dollar_Dict sync.Once

func Get_Main_Eq_dollar_Dict() gopurs_runtime.Value {
	once_Main_Eq_dollar_Dict.Do(func() {
		cache_Main_Eq_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Eq_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Eq_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqNumber gopurs_runtime.Value
var once_Main_eqNumber sync.Once

func Get_Main_eqNumber() gopurs_runtime.Value {
	once_Main_eqNumber.Do(func() {
		cache_Main_eqNumber = gopurs_runtime.Value{Type: 9, IntVal: 61300330, UnsafePtr: unsafe.Pointer((&Constructor_Main_Eq{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})}))}
	})
	return cache_Main_eqNumber
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eq(gopurs_runtime.CoerceToStruct[Constructor_Main_Eq](dict_0_box))
		})
	})
	return cache_Main_eq
}

type Constructor_Main_Eq struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[61300330] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eq)(ptr)
		_ = c
		switch key {
		case "eq":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Eq: " + key)
		}
	}
}

func Call_Main_Eq_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_eq(dict_0_loop *Constructor_Main_Eq) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Eq = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
