package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_getFoo gopurs_runtime.Value
var once_Main_getFoo sync.Once

func Get_Main_getFoo() gopurs_runtime.Value {
	once_Main_getFoo.Do(func() {
		cache_Main_getFoo = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_getFoo(o_0_box))
		})
	})
	return cache_Main_getFoo
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_foo(s_0_box.StrVal()))
		})
	})
	return cache_Main_foo
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(g_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f(g_0_box))
		})
	})
	return cache_Main_f
}

func Call_Main_getFoo(o_0_loop gopurs_runtime.Value) string {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	return gopurs_runtime.RecordGet(o_0, "foo").StrVal()
}

func Call_Main_foo(s_0_loop string) string {
	var s_0 string = s_0_loop
	_ = s_0
	return s_0
}

func Call_Main_f(g_0_loop gopurs_runtime.Value) string {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	return gopurs_runtime.RecordGet(gopurs_runtime.Apply(g_0, gopurs_runtime.RecordDict1("x", gopurs_runtime.Str("Hello"))), "x").StrVal()
}
