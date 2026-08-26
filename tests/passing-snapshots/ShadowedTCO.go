package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_zero_prime_ gopurs_runtime.Value
var once_Main_zero_prime_ sync.Once

func Get_Main_zero_prime_() gopurs_runtime.Value {
	once_Main_zero_prime_.Do(func() {
		cache_Main_zero_prime_ = gopurs_runtime.Func2(func(z_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_zero_prime_(z_0_box, v_1_box)
		})
	})
	return cache_Main_zero_prime_
}

var cache_Main_succ gopurs_runtime.Value
var once_Main_succ sync.Once

func Get_Main_succ() gopurs_runtime.Value {
	once_Main_succ.Do(func() {
		cache_Main_succ = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, zero_prime_1_1_box gopurs_runtime.Value, succ1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_succ(f_0_box, zero_prime_1_1_box, succ1_2_box)
		})
	})
	return cache_Main_succ
}

var cache_Main_runNat gopurs_runtime.Value
var once_Main_runNat sync.Once

func Get_Main_runNat() gopurs_runtime.Value {
	once_Main_runNat.Do(func() {
		cache_Main_runNat = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runNat(f_0_box)
		})
	})
	return cache_Main_runNat
}

var cache_Main_one_prime_ gopurs_runtime.Value
var once_Main_one_prime_ sync.Once

func Get_Main_one_prime_() gopurs_runtime.Value {
	once_Main_one_prime_.Do(func() {
		cache_Main_one_prime_ = gopurs_runtime.Func2(func(zero_prime_1_0_box gopurs_runtime.Value, succ1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_one_prime_(zero_prime_1_0_box, succ1_1_box)
		})
	})
	return cache_Main_one_prime_
}

var cache_Main_two gopurs_runtime.Value
var once_Main_two sync.Once

func Get_Main_two() gopurs_runtime.Value {
	once_Main_two.Do(func() {
		cache_Main_two = gopurs_runtime.Func2(func(zero_prime_1_0_box gopurs_runtime.Value, succ1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_two(zero_prime_1_0_box, succ1_1_box)
		})
	})
	return cache_Main_two
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, zero_prime_1_2_box gopurs_runtime.Value, succ1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_add(f_0_box, g_1_box, zero_prime_1_2_box, succ1_3_box)
		})
	})
	return cache_Main_add
}

var cache_Main_four gopurs_runtime.Value
var once_Main_four sync.Once

func Get_Main_four() gopurs_runtime.Value {
	once_Main_four.Do(func() {
		cache_Main_four = gopurs_runtime.Func2(func(zero_prime_1_0_box gopurs_runtime.Value, succ1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_four(zero_prime_1_0_box, succ1_1_box)
		})
	})
	return cache_Main_four
}

var cache_Main_fourNumber gopurs_runtime.Value
var once_Main_fourNumber sync.Once

func Get_Main_fourNumber() gopurs_runtime.Value {
	once_Main_fourNumber.Do(func() {
		cache_Main_fourNumber = gopurs_runtime.Float(gopurs_runtime.Float(4.0).FloatVal())
	})
	return cache_Main_fourNumber
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Float(4.0).FloatVal())).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_zero_prime_(z_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var z_0 gopurs_runtime.Value = z_0_loop
	_ = z_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return z_0
}

func Call_Main_succ(f_0_loop gopurs_runtime.Value, zero_prime_1_1_loop gopurs_runtime.Value, succ1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var zero_prime_1_1 gopurs_runtime.Value = zero_prime_1_1_loop
	_ = zero_prime_1_1
	var succ1_2 gopurs_runtime.Value = succ1_2_loop
	_ = succ1_2
	return gopurs_runtime.Apply(succ1_2, gopurs_runtime.Apply2(f_0, zero_prime_1_1, succ1_2))
}

func Call_Main_runNat(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(f_0, gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((n_1.FloatVal()) + (1.0))
	}))
}

func Call_Main_one_prime_(zero_prime_1_0_loop gopurs_runtime.Value, succ1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var zero_prime_1_0 gopurs_runtime.Value = zero_prime_1_0_loop
	_ = zero_prime_1_0
	var succ1_1 gopurs_runtime.Value = succ1_1_loop
	_ = succ1_1
	return gopurs_runtime.Apply(succ1_1, zero_prime_1_0)
}

func Call_Main_two(zero_prime_1_0_loop gopurs_runtime.Value, succ1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var zero_prime_1_0 gopurs_runtime.Value = zero_prime_1_0_loop
	_ = zero_prime_1_0
	var succ1_1 gopurs_runtime.Value = succ1_1_loop
	_ = succ1_1
	return gopurs_runtime.Apply(succ1_1, gopurs_runtime.Apply(succ1_1, zero_prime_1_0))
}

func Call_Main_add(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, zero_prime_1_2_loop gopurs_runtime.Value, succ1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var zero_prime_1_2 gopurs_runtime.Value = zero_prime_1_2_loop
	_ = zero_prime_1_2
	var succ1_3 gopurs_runtime.Value = succ1_3_loop
	_ = succ1_3
	return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, zero_prime_1_2, succ1_3), succ1_3)
}

func Call_Main_four(zero_prime_1_0_loop gopurs_runtime.Value, succ1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var zero_prime_1_0 gopurs_runtime.Value = zero_prime_1_0_loop
	_ = zero_prime_1_0
	var succ1_1 gopurs_runtime.Value = succ1_1_loop
	_ = succ1_1
	return gopurs_runtime.Apply(succ1_1, gopurs_runtime.Apply(succ1_1, gopurs_runtime.Apply(succ1_1, gopurs_runtime.Apply(succ1_1, zero_prime_1_0))))
}
