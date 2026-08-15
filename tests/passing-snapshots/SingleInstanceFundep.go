package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_SingleInstanceFundep_dollar_Dict gopurs_runtime.Value
var once_Main_SingleInstanceFundep_dollar_Dict sync.Once

func Get_Main_SingleInstanceFundep_dollar_Dict() gopurs_runtime.Value {
	once_Main_SingleInstanceFundep_dollar_Dict.Do(func() {
		cache_Main_SingleInstanceFundep_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_SingleInstanceFundep_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_SingleInstanceFundep_dollar_Dict
}

var cache_Main_singleInstanceFundepRow gopurs_runtime.Value
var once_Main_singleInstanceFundepRow sync.Once

func Get_Main_singleInstanceFundepRow() gopurs_runtime.Value {
	once_Main_singleInstanceFundepRow.Do(func() {
		cache_Main_singleInstanceFundepRow = gopurs_runtime.Value{Type: 9, IntVal: 4179949665, UnsafePtr: unsafe.Pointer((&Constructor_Main_SingleInstanceFundep{1, 513803634}))}
	})
	return cache_Main_singleInstanceFundepRow
}

var cache_Main_unified gopurs_runtime.Value
var once_Main_unified sync.Once

func Get_Main_unified() gopurs_runtime.Value {
	once_Main_unified.Do(func() {
		cache_Main_unified = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unified(dict_0_box)
		})
	})
	return cache_Main_unified
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

type Constructor_Main_SingleInstanceFundep struct {
	Rc uint32
	V0 uint32
}

func init() {
	gopurs_runtime.StructGetters[4179949665] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_SingleInstanceFundep)(ptr)
		_ = c
		switch key {
		case "unified":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_SingleInstanceFundep: " + key)
		}
	}
}

func Call_Main_SingleInstanceFundep_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_unified(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "unified")
}
