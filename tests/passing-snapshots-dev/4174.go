package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Effect_Console gopurs_runtime.Value
var once_Main_Effect_Console sync.Once

func Get_Main_Effect_Console() gopurs_runtime.Value {
	once_Main_Effect_Console.Do(func() {
		cache_Main_Effect_Console = gopurs_runtime.Value{Type: 9, IntVal: int64(4123976513), UnsafePtr: nil}
	})
	return cache_Main_Effect_Console
}

var cache_Main_Data_Unit gopurs_runtime.Value
var once_Main_Data_Unit sync.Once

func Get_Main_Data_Unit() gopurs_runtime.Value {
	once_Main_Data_Unit.Do(func() {
		cache_Main_Data_Unit = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Data_Unit(x_0_box)
		})
	})
	return cache_Main_Data_Unit
}

var cache_Main_n gopurs_runtime.Value
var once_Main_n sync.Once

func Get_Main_n() gopurs_runtime.Value {
	once_Main_n.Do(func() {
		cache_Main_n = Get_Data_Unit_unit()
	})
	return cache_Main_n
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Value{Type: 9, IntVal: int64(4123976513), UnsafePtr: nil}
	})
	return cache_Main_d
}

type Constructor_Main_Effect_Console struct {
	Rc uint32
}

func Call_Main_Data_Unit(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
