package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fooInt gopurs_runtime.Value
var once_Main_fooInt sync.Once

func Get_Main_fooInt() gopurs_runtime.Value {
	once_Main_fooInt.Do(func() {
		cache_Main_fooInt = gopurs_runtime.Value{}
	})
	return cache_Main_fooInt
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_foo
}

var cache_Main_bar gopurs_runtime.Value
var once_Main_bar sync.Once

func Get_Main_bar() gopurs_runtime.Value {
	once_Main_bar.Do(func() {
		cache_Main_bar = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bar(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_bar
}

type Constructor_Main_Foo[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0046006f006f
func init() {
	gopurs_runtime.StructGetters[2763139640] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Foo[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "bar":
			return gopurs_runtime.Box(c.V0)
		case "foo":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Foo: " + key)
		}
	}
}

func Call_Main_foo(dict_0_loop *Constructor_Main_Foo[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_bar(dict_0_loop *Constructor_Main_Foo[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}
