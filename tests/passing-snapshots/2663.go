package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(_dollar___unused_0_box, x_1_box)
		})
	})
	return cache_Main_foo
}

var cache_Main_foo__530935803 gopurs_runtime.Value
var once_Main_foo__530935803 sync.Once

func Get_Main_foo__530935803() gopurs_runtime.Value {
	once_Main_foo__530935803.Do(func() {
		cache_Main_foo__530935803 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_foo__530935803(x_0_box.IntVal))
		})
	})
	return cache_Main_foo__530935803
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_foo(_dollar___unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return x_1
}

func Call_Main_foo__530935803(x_0_loop int64) int64 {
foo__530935803:
	for {
		if false {
			continue foo__530935803
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return x_0
	}
}
