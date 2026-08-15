package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

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
		cache_Main_foo = gopurs_runtime.Func(func(r1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_foo(r1_0_box.FloatVal()))
		})
	})
	return cache_Main_foo
}

var cache_Main_r gopurs_runtime.Value
var once_Main_r sync.Once

func Get_Main_r() gopurs_runtime.Value {
	once_Main_r.Do(func() {
		cache_Main_r = gopurs_runtime.Float(3.0)
	})
	return cache_Main_r
}

func Call_Main_foo(r1_0_loop float64) float64 {
	var r1_0 float64 = r1_0_loop
	_ = r1_0
	return (r1_0) + (1.0)
}
