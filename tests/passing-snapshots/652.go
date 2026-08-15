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

var cache_Main_Bar_dollar_Dict gopurs_runtime.Value
var once_Main_Bar_dollar_Dict sync.Once

func Get_Main_Bar_dollar_Dict() gopurs_runtime.Value {
	once_Main_Bar_dollar_Dict.Do(func() {
		cache_Main_Bar_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Bar_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Bar_dollar_Dict
}

var cache_Main_Baz_dollar_Dict gopurs_runtime.Value
var once_Main_Baz_dollar_Dict sync.Once

func Get_Main_Baz_dollar_Dict() gopurs_runtime.Value {
	once_Main_Baz_dollar_Dict.Do(func() {
		cache_Main_Baz_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Baz_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Baz_dollar_Dict
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
		cache_Main_foo = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_foo
}

var cache_Main_bar gopurs_runtime.Value
var once_Main_bar sync.Once

func Get_Main_bar() gopurs_runtime.Value {
	once_Main_bar.Do(func() {
		cache_Main_bar = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_bar
}

var cache_Main_baz gopurs_runtime.Value
var once_Main_baz sync.Once

func Get_Main_baz() gopurs_runtime.Value {
	once_Main_baz.Do(func() {
		cache_Main_baz = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_baz(dictEq_0_box)
		})
	})
	return cache_Main_baz
}

type Constructor_Main_Foo struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2763139640] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Foo)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Foo: " + key)
		}
	}
}

type Constructor_Main_Bar struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2512729583] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Bar)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Bar: " + key)
		}
	}
}

type Constructor_Main_Baz struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2012165095] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Baz)(ptr)
		_ = c
		switch key {
		case "Bar1":
			return gopurs_runtime.Box(c.V0)
		case "Foo0":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Baz: " + key)
		}
	}
}

func Call_Main_Foo_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Bar_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Baz_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_baz(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2012165095, UnsafePtr: unsafe.Pointer(&Constructor_Main_Baz{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})})}
}
