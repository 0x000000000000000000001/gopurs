package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.RecordDict1("baz", gopurs_runtime.Str("baz"))
	})
	return cache_Main_x
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_blah gopurs_runtime.Value
var once_Main_blah sync.Once

func Get_Main_blah() gopurs_runtime.Value {
	once_Main_blah.Do(func() {
		cache_Main_blah = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_blah(x1_0_box)
		})
	})
	return cache_Main_blah
}

var cache_Main_blah__335853299 gopurs_runtime.Value
var once_Main_blah__335853299 sync.Once

func Get_Main_blah__335853299() gopurs_runtime.Value {
	once_Main_blah__335853299.Do(func() {
		cache_Main_blah__335853299 = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_blah__335853299(x1_0_box)
		})
	})
	return cache_Main_blah__335853299
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.RecordUpdate1(Get_Main_x(), "baz", gopurs_runtime.Str("blah"))
	})
	return cache_Main_test
}

func Call_Main_blah(x1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x1_0 gopurs_runtime.Value = x1_0_loop
	_ = x1_0
	return x1_0
}

func Call_Main_blah__335853299(x1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x1_0 gopurs_runtime.Value = x1_0_loop
	_ = x1_0
	return x1_0
}
