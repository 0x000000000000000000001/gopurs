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

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(v_0_box))
		})
	})
	return cache_Main_g
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Int(gopurs_runtime.Int(gopurs_runtime.Int(0).IntVal).IntVal)
	})
	return cache_Main_test2
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(v_0_box)
		})
	})
	return cache_Main_f
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)
	})
	return cache_Main_test1
}

var cache_Main_f__3152029599 gopurs_runtime.Value
var once_Main_f__3152029599 sync.Once

func Get_Main_f__3152029599() gopurs_runtime.Value {
	once_Main_f__3152029599.Do(func() {
		cache_Main_f__3152029599 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f__3152029599(v_0_box)
		})
	})
	return cache_Main_f__3152029599
}

var cache_Main_g__1930409696 gopurs_runtime.Value
var once_Main_g__1930409696 sync.Once

func Get_Main_g__1930409696() gopurs_runtime.Value {
	once_Main_g__1930409696.Do(func() {
		cache_Main_g__1930409696 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g__1930409696(v_0_box))
		})
	})
	return cache_Main_g__1930409696
}

func Call_Main_g(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return y_1
	})).IntVal
}

func Call_Main_f(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, v_0)
}

func Call_Main_f__3152029599(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, v_0)
}

func Call_Main_g__1930409696(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return y_1
	})).IntVal
}
