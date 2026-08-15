package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_z gopurs_runtime.Value
var once_Main_z sync.Once

func Get_Main_z() gopurs_runtime.Value {
	once_Main_z.Do(func() {
		cache_Main_z = gopurs_runtime.Float(0.5)
	})
	return cache_Main_z
}

var cache_Main_y gopurs_runtime.Value
var once_Main_y sync.Once

func Get_Main_y() gopurs_runtime.Value {
	once_Main_y.Do(func() {
		cache_Main_y = gopurs_runtime.Float(gopurs_runtime.Float(-0.5).FloatVal())
	})
	return cache_Main_y
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Float(gopurs_runtime.Float(-1.0).FloatVal())
	})
	return cache_Main_x
}

var cache_Main_w gopurs_runtime.Value
var once_Main_w sync.Once

func Get_Main_w() gopurs_runtime.Value {
	once_Main_w.Do(func() {
		cache_Main_w = gopurs_runtime.Float(1.0)
	})
	return cache_Main_w
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Float(1.0)
	})
	return cache_Main_test1
}

var cache_Main_q gopurs_runtime.Value
var once_Main_q sync.Once

func Get_Main_q() gopurs_runtime.Value {
	once_Main_q.Do(func() {
		cache_Main_q = gopurs_runtime.Float(1.0)
	})
	return cache_Main_q
}

var cache_Main_p gopurs_runtime.Value
var once_Main_p sync.Once

func Get_Main_p() gopurs_runtime.Value {
	once_Main_p.Do(func() {
		cache_Main_p = gopurs_runtime.Float(0.5)
	})
	return cache_Main_p
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_f(x1_0_box.FloatVal()))
		})
	})
	return cache_Main_f
}

func Call_Main_f(x1_0_loop float64) float64 {
	var x1_0 float64 = x1_0_loop
	_ = x1_0
	return gopurs_runtime.Float(-(gopurs_runtime.Float(x1_0).FloatVal())).FloatVal()
}
