package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(x_0_box)
		})
	})
	return cache_Main_f
}

var cache_Main_f__3508275643 gopurs_runtime.Value
var once_Main_f__3508275643 sync.Once

func Get_Main_f__3508275643() gopurs_runtime.Value {
	once_Main_f__3508275643.Do(func() {
		cache_Main_f__3508275643 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f__3508275643(x_0_box)
		})
	})
	return cache_Main_f__3508275643
}

var cache_Main_f__2001193531 gopurs_runtime.Value
var once_Main_f__2001193531 sync.Once

func Get_Main_f__2001193531() gopurs_runtime.Value {
	once_Main_f__2001193531.Do(func() {
		cache_Main_f__2001193531 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f__2001193531(x_0_box)
		})
	})
	return cache_Main_f__2001193531
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
	})
	return cache_Main_main
}

func Call_Main_f(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_f__3508275643(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_f__2001193531(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
