package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_To_dollar_Dict gopurs_runtime.Value
var once_Main_To_dollar_Dict sync.Once

func Get_Main_To_dollar_Dict() gopurs_runtime.Value {
	once_Main_To_dollar_Dict.Do(func() {
		cache_Main_To_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_To_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_To_dollar_Dict
}

var cache_Main_Pair gopurs_runtime.Value
var once_Main_Pair sync.Once

func Get_Main_Pair() gopurs_runtime.Value {
	once_Main_Pair.Do(func() {
		cache_Main_Pair = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
	})
	return cache_Main_Pair
}

var cache_Main_Pair_prime_ gopurs_runtime.Value
var once_Main_Pair_prime_ sync.Once

func Get_Main_Pair_prime_() gopurs_runtime.Value {
	once_Main_Pair_prime_.Do(func() {
		cache_Main_Pair_prime_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Pair_prime_(x_0_box)
		})
	})
	return cache_Main_Pair_prime_
}

var cache_Main_to2 gopurs_runtime.Value
var once_Main_to2 sync.Once

func Get_Main_to2() gopurs_runtime.Value {
	once_Main_to2.Do(func() {
		cache_Main_to2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_to2
}

var cache_Main_to1 gopurs_runtime.Value
var once_Main_to1 sync.Once

func Get_Main_to1() gopurs_runtime.Value {
	once_Main_to1.Do(func() {
		cache_Main_to1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_to1
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
	})
	return cache_Main_test6
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Int(42)
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(893478516), UnsafePtr: nil}
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

type Constructor_Main_Pair[T_a any, T_b any] struct {
	Rc uint32
}

type Constructor_Main_To[T_a any, T_b any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3555036389] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_To[any, any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_To: " + key)
		}
	}
}

func Call_Main_To_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Pair_prime_(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
