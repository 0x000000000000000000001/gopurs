package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_NonexportedType gopurs_runtime.Value
var once_Main_NonexportedType sync.Once

func Get_Main_NonexportedType() gopurs_runtime.Value {
	once_Main_NonexportedType.Do(func() {
		cache_Main_NonexportedType = gopurs_runtime.Value{Type: 9, IntVal: int64(3828477004), UnsafePtr: nil}
	})
	return cache_Main_NonexportedType
}

var cache_Main_NonexportedClass_dollar_Dict gopurs_runtime.Value
var once_Main_NonexportedClass_dollar_Dict sync.Once

func Get_Main_NonexportedClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_NonexportedClass_dollar_Dict.Do(func() {
		cache_Main_NonexportedClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NonexportedClass_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_NonexportedClass_dollar_Dict
}

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

var cache_Main_Const gopurs_runtime.Value
var once_Main_Const sync.Once

func Get_Main_Const() gopurs_runtime.Value {
	once_Main_Const.Do(func() {
		cache_Main_Const = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Const
}

var cache_Main_notExported gopurs_runtime.Value
var once_Main_notExported sync.Once

func Get_Main_notExported() gopurs_runtime.Value {
	once_Main_notExported.Do(func() {
		cache_Main_notExported = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_notExported(dict_0_box)
		})
	})
	return cache_Main_notExported
}

var cache_Main_nonExportedNonexportedType gopurs_runtime.Value
var once_Main_nonExportedNonexportedType sync.Once

func Get_Main_nonExportedNonexportedType() gopurs_runtime.Value {
	once_Main_nonExportedNonexportedType.Do(func() {
		cache_Main_nonExportedNonexportedType = gopurs_runtime.Value{Type: 9, IntVal: 888703674, UnsafePtr: unsafe.Pointer((&Constructor_Main_NonexportedClass{1, gopurs_runtime.Int(0)}))}
	})
	return cache_Main_nonExportedNonexportedType
}

var cache_Main_nonExportedFoo2 gopurs_runtime.Value
var once_Main_nonExportedFoo2 sync.Once

func Get_Main_nonExportedFoo2() gopurs_runtime.Value {
	once_Main_nonExportedFoo2.Do(func() {
		cache_Main_nonExportedFoo2 = gopurs_runtime.Func(func(dictNonexportedClass_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_nonExportedFoo2(dictNonexportedClass_0_box)
		})
	})
	return cache_Main_nonExportedFoo2
}

var cache_Main_nonExportedFoo gopurs_runtime.Value
var once_Main_nonExportedFoo sync.Once

func Get_Main_nonExportedFoo() gopurs_runtime.Value {
	once_Main_nonExportedFoo.Do(func() {
		cache_Main_nonExportedFoo = gopurs_runtime.Func(func(dictFoo_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_nonExportedFoo(dictFoo_0_box)
		})
	})
	return cache_Main_nonExportedFoo
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
			return Call_Main_foo(dict_0_box)
		})
	})
	return cache_Main_foo
}

var cache_Main_constFoo gopurs_runtime.Value
var once_Main_constFoo sync.Once

func Get_Main_constFoo() gopurs_runtime.Value {
	once_Main_constFoo.Do(func() {
		cache_Main_constFoo = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3828477004), UnsafePtr: nil}}))}
	})
	return cache_Main_constFoo
}

var cache_Main_notExported__1435057913 gopurs_runtime.Value
var once_Main_notExported__1435057913 sync.Once

func Get_Main_notExported__1435057913() gopurs_runtime.Value {
	once_Main_notExported__1435057913.Do(func() {
		cache_Main_notExported__1435057913 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_notExported__1435057913(dict_0_box)
		})
	})
	return cache_Main_notExported__1435057913
}

type Constructor_Main_NonexportedType struct {
	Rc uint32
}

type Constructor_Main_Const struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_NonexportedClass struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[888703674] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_NonexportedClass)(ptr)
		_ = c
		switch key {
		case "notExported":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_NonexportedClass: " + key)
		}
	}
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

func Call_Main_NonexportedClass_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Foo_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_notExported(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "notExported")
}

func Call_Main_nonExportedFoo2(dictNonexportedClass_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictNonexportedClass_0 gopurs_runtime.Value = dictNonexportedClass_0_loop
	_ = dictNonexportedClass_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, gopurs_runtime.RecordGet(dictNonexportedClass_0, "notExported")}))}
}

func Call_Main_nonExportedFoo(dictFoo_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoo_0 gopurs_runtime.Value = dictFoo_0_loop
	_ = dictFoo_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	})}))}
}

func Call_Main_foo(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "foo")
}

func Call_Main_notExported__1435057913(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "notExported")
}
