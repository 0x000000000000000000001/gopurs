package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ErrorSemigroup_dollar_Dict gopurs_runtime.Value
var once_Main_ErrorSemigroup_dollar_Dict sync.Once

func Get_Main_ErrorSemigroup_dollar_Dict() gopurs_runtime.Value {
	once_Main_ErrorSemigroup_dollar_Dict.Do(func() {
		cache_Main_ErrorSemigroup_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ErrorSemigroup_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ErrorSemigroup_dollar_Dict
}

var cache_Main_errorSemigroupMaybeMaybe gopurs_runtime.Value
var once_Main_errorSemigroupMaybeMaybe sync.Once

func Get_Main_errorSemigroupMaybeMaybe() gopurs_runtime.Value {
	once_Main_errorSemigroupMaybeMaybe.Do(func() {
		cache_Main_errorSemigroupMaybeMaybe = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_errorSemigroupMaybeMaybe
}

var cache_Main_errorSemigroupIdentityIde gopurs_runtime.Value
var once_Main_errorSemigroupIdentityIde sync.Once

func Get_Main_errorSemigroupIdentityIde() gopurs_runtime.Value {
	once_Main_errorSemigroupIdentityIde.Do(func() {
		cache_Main_errorSemigroupIdentityIde = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_errorSemigroupIdentityIde
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_ErrorSemigroup struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[246409291] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ErrorSemigroup)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_ErrorSemigroup: " + key)
		}
	}
}

func Call_Main_ErrorSemigroup_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
