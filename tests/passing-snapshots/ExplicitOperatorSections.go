package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_subtractOne gopurs_runtime.Value
var once_Main_subtractOne sync.Once

func Get_Main_subtractOne() gopurs_runtime.Value {
	once_Main_subtractOne.Do(func() {
		cache_Main_subtractOne = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_subtractOne(v_0_box.IntVal))
		})
	})
	return cache_Main_subtractOne
}

var cache_Main_named gopurs_runtime.Value
var once_Main_named sync.Once

func Get_Main_named() gopurs_runtime.Value {
	once_Main_named.Do(func() {
		cache_Main_named = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_named(v_0_box.IntVal))
		})
	})
	return cache_Main_named
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_addOne gopurs_runtime.Value
var once_Main_addOne sync.Once

func Get_Main_addOne() gopurs_runtime.Value {
	once_Main_addOne.Do(func() {
		cache_Main_addOne = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_addOne(v_0_box.IntVal))
		})
	})
	return cache_Main_addOne
}

func Call_Main_subtractOne(v_0_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	return (v_0) - (1)
}

func Call_Main_named(v_0_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	return (v_0) - (1)
}

func Call_Main_addOne(v_0_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	return (1) + (v_0)
}
