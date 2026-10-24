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

var cache_Main_fooString gopurs_runtime.Value
var once_Main_fooString sync.Once

func Get_Main_fooString() gopurs_runtime.Value {
	once_Main_fooString.Do(func() {
		cache_Main_fooString = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Rebox_Main_752630355_3584142444((&Constructor_Main_Foo[string]{1, gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return s_0
		})})))}
	})
	return cache_Main_fooString
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

type Constructor_Main_Foo[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0046006f006f
func init() {
	gopurs_runtime.StructGetters[2763139640] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Foo[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "foo":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Foo: " + key)
		}
	}
}

func Call_Main_foo(dict_0_loop *Constructor_Main_Foo[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_752630355_3584142444(in *Constructor_Main_Foo[string]) *Constructor_Main_Foo[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Foo[gopurs_runtime.Value])(unsafe.Pointer(in))
}
