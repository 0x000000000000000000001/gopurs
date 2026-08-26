package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_UseOperatorInDataParamKind gopurs_runtime.Value
var once_Main_UseOperatorInDataParamKind sync.Once

func Get_Main_UseOperatorInDataParamKind() gopurs_runtime.Value {
	once_Main_UseOperatorInDataParamKind.Do(func() {
		cache_Main_UseOperatorInDataParamKind = gopurs_runtime.Value{Type: 9, IntVal: int64(2094590957), UnsafePtr: nil}
	})
	return cache_Main_UseOperatorInDataParamKind
}

var cache_Main_UseOperatorInClassParamKind_dollar_Dict gopurs_runtime.Value
var once_Main_UseOperatorInClassParamKind_dollar_Dict sync.Once

func Get_Main_UseOperatorInClassParamKind_dollar_Dict() gopurs_runtime.Value {
	once_Main_UseOperatorInClassParamKind_dollar_Dict.Do(func() {
		cache_Main_UseOperatorInClassParamKind_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_UseOperatorInClassParamKind_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_UseOperatorInClassParamKind_dollar_Dict
}

var cache_Main_Compose gopurs_runtime.Value
var once_Main_Compose sync.Once

func Get_Main_Compose() gopurs_runtime.Value {
	once_Main_Compose.Do(func() {
		cache_Main_Compose = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Compose
}

var cache_Main_testPrecedence2 gopurs_runtime.Value
var once_Main_testPrecedence2 sync.Once

func Get_Main_testPrecedence2() gopurs_runtime.Value {
	once_Main_testPrecedence2.Do(func() {
		cache_Main_testPrecedence2 = gopurs_runtime.Func2(func(nat_0_box gopurs_runtime.Value, fx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testPrecedence2(nat_0_box, fx_1_box)
		})
	})
	return cache_Main_testPrecedence2
}

var cache_Main_testPrecedence1 gopurs_runtime.Value
var once_Main_testPrecedence1 sync.Once

func Get_Main_testPrecedence1() gopurs_runtime.Value {
	once_Main_testPrecedence1.Do(func() {
		cache_Main_testPrecedence1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testPrecedence1(x_0_box)
		})
	})
	return cache_Main_testPrecedence1
}

var cache_Main_testParens gopurs_runtime.Value
var once_Main_testParens sync.Once

func Get_Main_testParens() gopurs_runtime.Value {
	once_Main_testParens.Do(func() {
		cache_Main_testParens = gopurs_runtime.Func(func(nat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testParens(nat_0_box)
		})
	})
	return cache_Main_testParens
}

var cache_Main_swap gopurs_runtime.Value
var once_Main_swap sync.Once

func Get_Main_swap() gopurs_runtime.Value {
	once_Main_swap.Do(func() {
		cache_Main_swap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3190713900, UnsafePtr: unsafe.Pointer(Call_Main_swap(gopurs_runtime.CoerceToStruct[Constructor_A_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
		})
	})
	return cache_Main_swap
}

var cache_Main_natty gopurs_runtime.Value
var once_Main_natty sync.Once

func Get_Main_natty() gopurs_runtime.Value {
	once_Main_natty.Do(func() {
		cache_Main_natty = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_natty(x_0_box)
		})
	})
	return cache_Main_natty
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_UseOperatorInDataParamKind[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Compose[T_f any, T_g any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_UseOperatorInClassParamKind[T_a any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[40269555] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_UseOperatorInClassParamKind[any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_UseOperatorInClassParamKind: " + key)
		}
	}
}

func Call_Main_UseOperatorInClassParamKind_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_testPrecedence2(nat_0_loop gopurs_runtime.Value, fx_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var nat_0 gopurs_runtime.Value = nat_0_loop
	_ = nat_0
	var fx_1 gopurs_runtime.Value = fx_1_loop
	_ = fx_1
	return gopurs_runtime.Apply(nat_0, fx_1)
}

func Call_Main_testPrecedence1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_testParens(nat_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var nat_0 gopurs_runtime.Value = nat_0_loop
	_ = nat_0
	return nat_0
}

func Call_Main_swap(v_0_loop *Constructor_A_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_A_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	var v_0 *Constructor_A_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (&Constructor_A_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_0).V1, (v_0).V0})
}

func Call_Main_natty(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
