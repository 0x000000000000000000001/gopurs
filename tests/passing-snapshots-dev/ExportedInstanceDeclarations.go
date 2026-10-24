package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

type Constructor_Main_NonexportedType struct {
	Rc uint32
}

type Constructor_Main_Const[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_NonexportedClass[T_a any] struct {
	Rc uint32
	V0 T_a
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004e006f006e006500780070006f00720074006500640043006c006100730073
func init() {
	gopurs_runtime.StructGetters[888703674] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_NonexportedClass[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "notExported":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_NonexportedClass: " + key)
		}
	}
}

type Constructor_Main_Foo[T_a any] struct {
	Rc uint32
	V0 T_a
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

func Call_Main_nonExportedFoo2(dictNonexportedClass_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictNonexportedClass_0 gopurs_runtime.Value = dictNonexportedClass_0_loop
	_ = dictNonexportedClass_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictNonexportedClass_0, "notExported")}))}
}

func Call_Main_nonExportedFoo(dictFoo_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoo_0 gopurs_runtime.Value = dictFoo_0_loop
	_ = dictFoo_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo[gopurs_runtime.Value]{1, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})}))}
}

func Call_Main_foo(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "foo")
}
