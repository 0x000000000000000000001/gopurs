package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn(v_0_box)
		})
	})
	return cache_Main_fn
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.RecordDict1("b", gopurs_runtime.RecordDict1("c", gopurs_runtime.RecordDict1("d", gopurs_runtime.Int(2))))
	})
	return cache_Main_a
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Int(4)
	})
	return cache_Main_d
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_fn(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "b"), "c"), "d")
}
