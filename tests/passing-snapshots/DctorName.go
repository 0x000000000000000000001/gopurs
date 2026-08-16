package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Baz_prime__prime_ gopurs_runtime.Value
var once_Main_Baz_prime__prime_ sync.Once

func Get_Main_Baz_prime__prime_() gopurs_runtime.Value {
	once_Main_Baz_prime__prime_.Do(func() {
		cache_Main_Baz_prime__prime_ = gopurs_runtime.Value{Type: 9, IntVal: int64(905033287), UnsafePtr: nil}
	})
	return cache_Main_Baz_prime__prime_
}

var cache_Main_Baz_prime_ gopurs_runtime.Value
var once_Main_Baz_prime_ sync.Once

func Get_Main_Baz_prime_() gopurs_runtime.Value {
	once_Main_Baz_prime_.Do(func() {
		cache_Main_Baz_prime_ = gopurs_runtime.Value{Type: 9, IntVal: int64(2412272388), UnsafePtr: nil}
	})
	return cache_Main_Baz_prime_
}

var cache_Main_Bar_prime_ gopurs_runtime.Value
var once_Main_Bar_prime_ sync.Once

func Get_Main_Bar_prime_() gopurs_runtime.Value {
	once_Main_Bar_prime_.Do(func() {
		cache_Main_Bar_prime_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Bar_prime_(x_0_box)
		})
	})
	return cache_Main_Bar_prime_
}

var cache_Main_Foo_prime_ gopurs_runtime.Value
var once_Main_Foo_prime_ sync.Once

func Get_Main_Foo_prime_() gopurs_runtime.Value {
	once_Main_Foo_prime_.Do(func() {
		cache_Main_Foo_prime_ = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Foo_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_h gopurs_runtime.Value
var once_Main_h sync.Once

func Get_Main_h() gopurs_runtime.Value {
	once_Main_h.Do(func() {
		cache_Main_h = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_h(v_0_box.IntVal))
		})
	})
	return cache_Main_h
}

var cache_Main_h_prime_ gopurs_runtime.Value
var once_Main_h_prime_ sync.Once

func Get_Main_h_prime_() gopurs_runtime.Value {
	once_Main_h_prime_.Do(func() {
		cache_Main_h_prime_ = gopurs_runtime.Int(8)
	})
	return cache_Main_h_prime_
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(uint32(v_0_box.IntVal)))
		})
	})
	return cache_Main_g
}

var cache_Main_g_prime_ gopurs_runtime.Value
var once_Main_g_prime_ sync.Once

func Get_Main_g_prime_() gopurs_runtime.Value {
	once_Main_g_prime_.Do(func() {
		cache_Main_g_prime_ = gopurs_runtime.Int(0)
	})
	return cache_Main_g_prime_
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_f(a_0_box))
		})
	})
	return cache_Main_f
}

var cache_Main_f_prime_ gopurs_runtime.Value
var once_Main_f_prime_ sync.Once

func Get_Main_f_prime_() gopurs_runtime.Value {
	once_Main_f_prime_.Do(func() {
		cache_Main_f_prime_ = gopurs_runtime.Bool(true)
	})
	return cache_Main_f_prime_
}

type Constructor_Main_Baz_prime__prime_ struct {
	Rc uint32
}

type Constructor_Main_Baz_prime_ struct {
	Rc uint32
}

type Constructor_Main_Foo_prime_ struct {
	Rc uint32
	V0 int64
}

func Call_Main_Bar_prime_(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_h(v_0_loop int64) int64 {
	var v_0 int64 = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if (v_0) <= (10) {
			__t0 = (v_0) * (2)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = 10
	}
end_branch_0:
	return __t0
}

func Call_Main_g(v_0_loop uint32) int64 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if v_0 == 905033287 {
			__t0 = 0
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 2412272388 {
			__t0 = 1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
	}
end_branch_0:
	return __t0
}

func Call_Main_f(a_0_loop gopurs_runtime.Value) bool {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return true
}
