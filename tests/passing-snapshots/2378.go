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
			return Call_Main_Foo_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Foo_dollar_Dict
}

var cache_Main_Foo_dollar_Dict__125961192 gopurs_runtime.Value
var once_Main_Foo_dollar_Dict__125961192 sync.Once

func Get_Main_Foo_dollar_Dict__125961192() gopurs_runtime.Value {
	once_Main_Foo_dollar_Dict__125961192.Do(func() {
		cache_Main_Foo_dollar_Dict__125961192 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Foo_dollar_Dict__125961192(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Foo_dollar_Dict__125961192
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fooX gopurs_runtime.Value
var once_Main_fooX sync.Once

func Get_Main_fooX() gopurs_runtime.Value {
	once_Main_fooX.Do(func() {
		cache_Main_fooX = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_fooX
}

type Constructor_Main_Foo[T_a any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2763139640] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Foo[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Foo: " + key)
		}
	}
}

func Call_Main_Foo_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_Foo_dollar_Dict__125961192(x_0_loop struct {
}) gopurs_runtime.Value {
Foo_dollar_Dict__125961192:
	for {
		if false {
			continue Foo_dollar_Dict__125961192
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	}
}
