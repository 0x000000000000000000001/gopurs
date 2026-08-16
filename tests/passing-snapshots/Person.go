package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Person gopurs_runtime.Value
var once_Main_Person sync.Once

func Get_Main_Person() gopurs_runtime.Value {
	once_Main_Person.Do(func() {
		cache_Main_Person = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Person
}

var cache_Main_showPerson gopurs_runtime.Value
var once_Main_showPerson sync.Once

func Get_Main_showPerson() gopurs_runtime.Value {
	once_Main_showPerson.Do(func() {
		cache_Main_showPerson = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_showPerson(p_0_box))
		})
	})
	return cache_Main_showPerson
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Person struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_showPerson(p_0_loop gopurs_runtime.Value) string {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return ((gopurs_runtime.RecordGet(p_0, "name").StrVal()) + (", aged ")) + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.RecordGet(p_0, "age").FloatVal())).StrVal())
}
