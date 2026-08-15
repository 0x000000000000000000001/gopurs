package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Person gopurs_runtime.Value
var once_Main_Person sync.Once

func Get_Main_Person() gopurs_runtime.Value {
	once_Main_Person.Do(func() {
		cache_Main_Person = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1621815371, UnsafePtr: unsafe.Pointer((&Constructor_Main_Person{1, value0.StrVal(), (value1.IntVal) != (0)}))}
			})
		})
	})
	return cache_Main_Person
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_getName gopurs_runtime.Value
var once_Main_getName sync.Once

func Get_Main_getName() gopurs_runtime.Value {
	once_Main_getName.Do(func() {
		cache_Main_getName = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_getName(gopurs_runtime.CoerceToStruct[Constructor_Main_Person](p_0_box)))
		})
	})
	return cache_Main_getName
}

var cache_Main_name gopurs_runtime.Value
var once_Main_name sync.Once

func Get_Main_name() gopurs_runtime.Value {
	once_Main_name.Do(func() {
		cache_Main_name = gopurs_runtime.Str("John Smith")
	})
	return cache_Main_name
}

type Constructor_Main_Person struct {
	Rc uint32
	V0 string
	V1 bool
}

func Call_Main_getName(p_0_loop *Constructor_Main_Person) string {
	var p_0 *Constructor_Main_Person = p_0_loop
	_ = p_0
	var __t0 string
	{
		if (p_0).V1 {
			__t0 = (p_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = "Unknown"
	}
end_branch_0:
	return __t0
}
