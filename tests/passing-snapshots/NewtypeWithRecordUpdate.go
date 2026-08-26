package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_NewType gopurs_runtime.Value
var once_Main_NewType sync.Once

func Get_Main_NewType() gopurs_runtime.Value {
	once_Main_NewType.Do(func() {
		cache_Main_NewType = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NewType(x_0_box)
		})
	})
	return cache_Main_NewType
}

var cache_Main_rec1 gopurs_runtime.Value
var once_Main_rec1 sync.Once

func Get_Main_rec1() gopurs_runtime.Value {
	once_Main_rec1.Do(func() {
		cache_Main_rec1 = gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0), gopurs_runtime.Float(0.0))
	})
	return cache_Main_rec1
}

var cache_Main_rec2 gopurs_runtime.Value
var once_Main_rec2 sync.Once

func Get_Main_rec2() gopurs_runtime.Value {
	once_Main_rec2.Do(func() {
		cache_Main_rec2 = gopurs_runtime.RecordUpdate1(Get_Main_rec1(), "a", gopurs_runtime.Float(1.0))
	})
	return cache_Main_rec2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_NewType(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
