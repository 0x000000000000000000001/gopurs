package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Foo_dollar_Dict gopurs_runtime.Value
var once_Main_Foo_dollar_Dict sync.Once

func Get_Main_Foo_dollar_Dict() gopurs_runtime.Value {
	once_Main_Foo_dollar_Dict.Do(func() {
		cache_Main_Foo_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Foo_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Foo_dollar_Dict
}

var cache_Main_fooLogEff gopurs_runtime.Value
var once_Main_fooLogEff sync.Once

func Get_Main_fooLogEff() gopurs_runtime.Value {
	once_Main_fooLogEff.Do(func() {
		cache_Main_fooLogEff = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, Get_Effect_Console_log()}))}
	})
	return cache_Main_fooLogEff
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo](dict_0_box))
		})
	})
	return cache_Main_foo
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_foo__519845775 gopurs_runtime.Value
var once_Main_foo__519845775 sync.Once

func Get_Main_foo__519845775() gopurs_runtime.Value {
	once_Main_foo__519845775.Do(func() {
		cache_Main_foo__519845775 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo__519845775(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo](dict_0_box))
		})
	})
	return cache_Main_foo__519845775
}

type Constructor_Main_Foo struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2763139640] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Foo)(ptr)
		_ = c
		switch key {
		case "foo":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Foo: " + key)
		}
	}
}

func Call_Main_Foo_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_foo(dict_0_loop *Constructor_Main_Foo) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_foo__519845775(dict_0_loop *Constructor_Main_Foo) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
