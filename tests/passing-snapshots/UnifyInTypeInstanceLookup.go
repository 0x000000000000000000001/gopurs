package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Value{Type: 9, IntVal: int64(1714575428), UnsafePtr: nil}
	})
	return cache_Main_Z
}

var cache_Main_S gopurs_runtime.Value
var once_Main_S sync.Once

func Get_Main_S() gopurs_runtime.Value {
	once_Main_S.Do(func() {
		cache_Main_S = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_S
}

var cache_Main_S__2743305020 gopurs_runtime.Value
var once_Main_S__2743305020 sync.Once

func Get_Main_S__2743305020() gopurs_runtime.Value {
	once_Main_S__2743305020.Do(func() {
		cache_Main_S__2743305020 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_S__2743305020(uint32(__eta_norm_0_unused_0_box.IntVal))
		})
	})
	return cache_Main_S__2743305020
}

var cache_Main_EQ_dollar_Dict gopurs_runtime.Value
var once_Main_EQ_dollar_Dict sync.Once

func Get_Main_EQ_dollar_Dict() gopurs_runtime.Value {
	once_Main_EQ_dollar_Dict.Do(func() {
		cache_Main_EQ_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_EQ_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_EQ_dollar_Dict
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(_dollar___unused_0_box, a_1_box, v_2_box)
		})
	})
	return cache_Main_test
}

var cache_Main_spin gopurs_runtime.Value
var once_Main_spin sync.Once

func Get_Main_spin() gopurs_runtime.Value {
	once_Main_spin.Do(func() {
		cache_Main_spin = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_spin(a_0_box)
		})
	})
	return cache_Main_spin
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(dictEQ_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1(dictEQ_0_box)
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_eqT
}

var cache_Main_eqF gopurs_runtime.Value
var once_Main_eqF sync.Once

func Get_Main_eqF() gopurs_runtime.Value {
	once_Main_eqF.Do(func() {
		cache_Main_eqF = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_eqF
}

type Constructor_Main_Z struct {
	Rc uint32
}

type Constructor_Main_S[T_n any] struct {
	Rc uint32
	V0 T_n
}

type Constructor_Main_EQ[T_x any, T_y any, T_b any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3323825930] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_EQ[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_EQ: " + key)
		}
	}
}

func Call_Main_S__2743305020(__eta_norm_0_unused_0_loop uint32) gopurs_runtime.Value {
S__2743305020:
	for {
		if false {
			continue S__2743305020
		}
		var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Value{Type: 9, IntVal: int64(1714575428), UnsafePtr: nil}
	}
}

func Call_Main_EQ_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_test(_dollar___unused_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return a_1
}

func Call_Main_spin(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
spin:
	for {
		if false {
			continue spin
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		a_0_loop = a_0
		continue spin
		return func() gopurs_runtime.Value { panic("unreachable") }()
	}
}

func Call_Main_test1(dictEQ_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEQ_0 gopurs_runtime.Value = dictEQ_0_loop
	_ = dictEQ_0
	return Call_Main_spin(gopurs_runtime.Int(int64(1)))
}
