package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_A gopurs_runtime.Value
var once_Main_A sync.Once

func Get_Main_A() gopurs_runtime.Value {
	once_Main_A.Do(func() {
		cache_Main_A = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_A
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_B
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_g(b_0_box)
		})
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(a_0_box)
		})
	})
	return cache_Main_f
}

var cache_Main_showN gopurs_runtime.Value
var once_Main_showN sync.Once

func Get_Main_showN() gopurs_runtime.Value {
	once_Main_showN.Do(func() {
		cache_Main_showN = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showN(a_0_box)
		})
	})
	return cache_Main_showN
}

type Constructor_Main_A struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_B struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_g(b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	return Call_Main_f(b_0)
}

func Call_Main_f(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return Call_Main_f(a_0)
}

func Call_Main_showN(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return Call_Main_f(a_0)
}
