package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Value{Type: 9, IntVal: int64(2763139640), UnsafePtr: nil}
	})
	return cache_Main_Foo
}

var cache_Main_classNameFoo gopurs_runtime.Value
var once_Main_classNameFoo sync.Once

func Get_Main_classNameFoo() gopurs_runtime.Value {
	once_Main_classNameFoo.Do(func() {
		cache_Main_classNameFoo = gopurs_runtime.Value{Type: 9, IntVal: 1687123273, UnsafePtr: unsafe.Pointer(Rebox_Main_2145685968_2100274173((&Constructor_ImportedClassName_ClassName[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})})))}
	})
	return cache_Main_classNameFoo
}

var cache_Main_classNameFoo1 gopurs_runtime.Value
var once_Main_classNameFoo1 sync.Once

func Get_Main_classNameFoo1() gopurs_runtime.Value {
	once_Main_classNameFoo1.Do(func() {
		cache_Main_classNameFoo1 = gopurs_runtime.Value{Type: 9, IntVal: 1438891703, UnsafePtr: unsafe.Pointer(Rebox_Main_1049220494_954425347((&Constructor_Main_ClassName[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})})))}
	})
	return cache_Main_classNameFoo1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(gopurs_runtime.CoerceToStruct[Constructor_Main_ClassName[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_foo
}

type Constructor_Main_Foo struct {
	Rc uint32
}

type Constructor_Main_ClassName[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043006c006100730073004e0061006d0065
func init() {
	gopurs_runtime.StructGetters[1438891703] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ClassName[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "foo":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ClassName: " + key)
		}
	}
}

func Call_Main_foo(dict_0_loop *Constructor_Main_ClassName[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ClassName[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1049220494_954425347(in *Constructor_Main_ClassName[uint32]) *Constructor_Main_ClassName[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_ClassName[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2145685968_2100274173(in *Constructor_ImportedClassName_ClassName[uint32]) *Constructor_ImportedClassName_ClassName[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_ImportedClassName_ClassName[gopurs_runtime.Value])(unsafe.Pointer(in))
}
