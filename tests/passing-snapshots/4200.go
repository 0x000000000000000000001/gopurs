package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_NewA gopurs_runtime.Value
var once_Main_NewA sync.Once

func Get_Main_NewA() gopurs_runtime.Value {
	once_Main_NewA.Do(func() {
		cache_Main_NewA = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NewA(x_0_box)
		})
	})
	return cache_Main_NewA
}

var cache_Main_newtypeNewA_ gopurs_runtime.Value
var once_Main_newtypeNewA_ sync.Once

func Get_Main_newtypeNewA_() gopurs_runtime.Value {
	once_Main_newtypeNewA_.Do(func() {
		cache_Main_newtypeNewA_ = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_newtypeNewA_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_NewA(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
