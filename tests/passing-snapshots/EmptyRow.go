package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Foo
}

var cache_Main_Foo__500064833 gopurs_runtime.Value
var once_Main_Foo__500064833 sync.Once

func Get_Main_Foo__500064833() gopurs_runtime.Value {
	once_Main_Foo__500064833.Do(func() {
		cache_Main_Foo__500064833 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Foo__500064833(__eta_norm_0_0_box)
		})
	})
	return cache_Main_Foo__500064833
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.RecordDict0()
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

type Constructor_Main_Foo[T_r any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_Foo__500064833(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Foo__500064833:
	for {
		if false {
			continue Foo__500064833
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}
