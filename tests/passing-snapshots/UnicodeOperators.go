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

var cache_Main_emptySet gopurs_runtime.Value
var once_Main_emptySet sync.Once

func Get_Main_emptySet() gopurs_runtime.Value {
	once_Main_emptySet.Do(func() {
		cache_Main_emptySet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_emptySet(v_0_box))
		})
	})
	return cache_Main_emptySet
}

var cache_Main_elem gopurs_runtime.Value
var once_Main_elem sync.Once

func Get_Main_elem() gopurs_runtime.Value {
	once_Main_elem.Do(func() {
		cache_Main_elem = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_elem(x_0_box, f_1_box))
		})
	})
	return cache_Main_elem
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Bool((gopurs_runtime.Bool(true).IntVal) != (0))
	})
	return cache_Main_test2
}

var cache_Main_compose gopurs_runtime.Value
var once_Main_compose sync.Once

func Get_Main_compose() gopurs_runtime.Value {
	once_Main_compose.Do(func() {
		cache_Main_compose = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_compose(f_0_box, g_1_box, a_2_box)
		})
	})
	return cache_Main_compose
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1(a_0_box)
		})
	})
	return cache_Main_test1
}

var cache_Main_compose__1200091329 gopurs_runtime.Value
var once_Main_compose__1200091329 sync.Once

func Get_Main_compose__1200091329() gopurs_runtime.Value {
	once_Main_compose__1200091329.Do(func() {
		cache_Main_compose__1200091329 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_compose__1200091329(f_0_box, g_1_box, a_2_box)
		})
	})
	return cache_Main_compose__1200091329
}

var cache_Main_elem__2654190116 gopurs_runtime.Value
var once_Main_elem__2654190116 sync.Once

func Get_Main_elem__2654190116() gopurs_runtime.Value {
	once_Main_elem__2654190116.Do(func() {
		cache_Main_elem__2654190116 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_elem__2654190116(x_0_box, f_1_box))
		})
	})
	return cache_Main_elem__2654190116
}

var cache_Main_emptySet__1966714060 gopurs_runtime.Value
var once_Main_emptySet__1966714060 sync.Once

func Get_Main_emptySet__1966714060() gopurs_runtime.Value {
	once_Main_emptySet__1966714060.Do(func() {
		cache_Main_emptySet__1966714060 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_emptySet__1966714060(v_0_box.IntVal))
		})
	})
	return cache_Main_emptySet__1966714060
}

func Call_Main_emptySet(v_0_loop gopurs_runtime.Value) bool {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return true
}

func Call_Main_elem(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) bool {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return (gopurs_runtime.Apply(f_1, x_0).IntVal) != (0)
}

func Call_Main_compose(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, a_2))
}

func Call_Main_test1(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return a_0
}

func Call_Main_compose__1200091329(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, a_2))
}

func Call_Main_elem__2654190116(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) bool {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return (gopurs_runtime.Apply(f_1, x_0).IntVal) != (0)
}

func Call_Main_emptySet__1966714060(v_0_loop int64) bool {
	var v_0 int64 = v_0_loop
	_ = v_0
	return true
}
