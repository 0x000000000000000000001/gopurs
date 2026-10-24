package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Pair_prime_(uint32(x_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_Pair_prime_
}

var cache_Main_to2 gopurs_runtime.Value
var once_Main_to2 sync.Once

func Get_Main_to2() gopurs_runtime.Value {
	once_Main_to2.Do(func() {
		cache_Main_to2 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_to2
}

var cache_Main_to1 gopurs_runtime.Value
var once_Main_to1 sync.Once

func Get_Main_to1() gopurs_runtime.Value {
	once_Main_to1.Do(func() {
		cache_Main_to1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
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
		cache_Main_test5 = gopurs_runtime.Int(int64(42))
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0054006f
func init() {
	gopurs_runtime.StructGetters[3555036389] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_To[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_To: " + key)
		}
	}
}

func Call_Main_Pair_prime_(x_0_loop uint32) uint32 {
	var x_0 uint32 = x_0_loop
	_ = x_0
	return x_0
}
