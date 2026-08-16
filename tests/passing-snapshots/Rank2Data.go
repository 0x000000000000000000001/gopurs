package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Nat gopurs_runtime.Value
var once_Main_Nat sync.Once

func Get_Main_Nat() gopurs_runtime.Value {
	once_Main_Nat.Do(func() {
		cache_Main_Nat = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Nat
}

var cache_Main_Id gopurs_runtime.Value
var once_Main_Id sync.Once

func Get_Main_Id() gopurs_runtime.Value {
	once_Main_Id.Do(func() {
		cache_Main_Id = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Id
}

var cache_Main_zero_prime_ gopurs_runtime.Value
var once_Main_zero_prime_ sync.Once

func Get_Main_zero_prime_() gopurs_runtime.Value {
	once_Main_zero_prime_.Do(func() {
		cache_Main_zero_prime_ = gopurs_runtime.Func(func(zero_prime_1_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return zero_prime_1_0
			})
		})
	})
	return cache_Main_zero_prime_
}

var cache_Main_succ gopurs_runtime.Value
var once_Main_succ sync.Once

func Get_Main_succ() gopurs_runtime.Value {
	once_Main_succ.Do(func() {
		cache_Main_succ = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_succ(n_0_box)
		})
	})
	return cache_Main_succ
}

var cache_Main_two gopurs_runtime.Value
var once_Main_two sync.Once

func Get_Main_two() gopurs_runtime.Value {
	once_Main_two.Do(func() {
		cache_Main_two = gopurs_runtime.Func(func(zero_prime_1_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(succ1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(succ1_1, gopurs_runtime.Apply2(Get_Main_zero_prime_(), zero_prime_1_0, succ1_1))
			})
		})
	})
	return cache_Main_two
}

var cache_Main_runNat gopurs_runtime.Value
var once_Main_runNat sync.Once

func Get_Main_runNat() gopurs_runtime.Value {
	once_Main_runNat.Do(func() {
		cache_Main_runNat = gopurs_runtime.Func(func(nat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_runNat(nat_0_box))
		})
	})
	return cache_Main_runNat
}

var cache_Main_runId gopurs_runtime.Value
var once_Main_runId sync.Once

func Get_Main_runId() gopurs_runtime.Value {
	once_Main_runId.Do(func() {
		cache_Main_runId = gopurs_runtime.Func2(func(id_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runId(id_0_box, a_1_box)
		})
	})
	return cache_Main_runId
}

var cache_Main_one_prime_ gopurs_runtime.Value
var once_Main_one_prime_ sync.Once

func Get_Main_one_prime_() gopurs_runtime.Value {
	once_Main_one_prime_.Do(func() {
		cache_Main_one_prime_ = gopurs_runtime.Func(func(zero_prime_1_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(succ1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(succ1_1, gopurs_runtime.Apply2(Get_Main_zero_prime_(), zero_prime_1_0, succ1_1))
			})
		})
	})
	return cache_Main_one_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_add(n_0_box, m_1_box)
		})
	})
	return cache_Main_add
}

var cache_Main_four gopurs_runtime.Value
var once_Main_four sync.Once

func Get_Main_four() gopurs_runtime.Value {
	once_Main_four.Do(func() {
		cache_Main_four = gopurs_runtime.Func(func(zero_prime_1_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(succ1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Main_two(), gopurs_runtime.Apply2(Get_Main_two(), zero_prime_1_0, succ1_1), succ1_1)
			})
		})
	})
	return cache_Main_four
}

var cache_Main_fourNumber gopurs_runtime.Value
var once_Main_fourNumber sync.Once

func Get_Main_fourNumber() gopurs_runtime.Value {
	once_Main_fourNumber.Do(func() {
		cache_Main_fourNumber = gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Main_four(), gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((n_0.FloatVal()) + (1.0))
		})).FloatVal())
	})
	return cache_Main_fourNumber
}

type Constructor_Main_Nat struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Id struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_succ(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var n_0 gopurs_runtime.Value = n_0_loop
	_ = n_0
	return gopurs_runtime.Func(func(zero_prime_1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(succ1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(succ1_2, gopurs_runtime.Apply2(n_0, zero_prime_1_1, succ1_2))
		})
	})
}

func Call_Main_runNat(nat_0_loop gopurs_runtime.Value) float64 {
	var nat_0 gopurs_runtime.Value = nat_0_loop
	_ = nat_0
	return gopurs_runtime.Apply2(nat_0, gopurs_runtime.Float(0.0), gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((n_1.FloatVal()) + (1.0))
	})).FloatVal()
}

func Call_Main_runId(id_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var id_0 gopurs_runtime.Value = id_0_loop
	_ = id_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply(id_0, a_1)
}

func Call_Main_add(n_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var n_0 gopurs_runtime.Value = n_0_loop
	_ = n_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	return gopurs_runtime.Func(func(zero_prime_1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(succ1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(m_1, gopurs_runtime.Apply2(n_0, zero_prime_1_2, succ1_3), succ1_3)
		})
	})
}
