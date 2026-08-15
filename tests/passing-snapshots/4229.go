package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Prim gopurs_runtime.Value
var once_Main_Prim sync.Once

func Get_Main_Prim() gopurs_runtime.Value {
	once_Main_Prim.Do(func() {
		cache_Main_Prim = gopurs_runtime.Value{Type: 9, IntVal: int64(1858774840), UnsafePtr: nil}
	})
	return cache_Main_Prim
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
		cache_Main_f = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_f(_dollar___unused_0_box, v_1_box.IntVal))
		})
	})
	return cache_Main_f
}

var cache_Main_f_prime_ gopurs_runtime.Value
var once_Main_f_prime_ sync.Once

func Get_Main_f_prime_() gopurs_runtime.Value {
	once_Main_f_prime_.Do(func() {
		cache_Main_f_prime_ = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_f_prime_(v_0_box.IntVal))
		})
	})
	return cache_Main_f_prime_
}

type Constructor_Main_Prim struct {
	Rc uint32
}

func Call_Main_f(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop int64) int64 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 int64 = v_1_loop
	_ = v_1
	var __t0 int64
	{
		if (v_1) == (0) {
			__t0 = 0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
	}
end_branch_0:
	return gopurs_runtime.Int(__t0).IntVal
}

func Call_Main_f_prime_(v_0_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if (v_0) == (0) {
			__t0 = 0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
	}
end_branch_0:
	return gopurs_runtime.Int(__t0).IntVal
}
