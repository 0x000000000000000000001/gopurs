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
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_To_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_To_dollar_Dict
}

var cache_Main_To_dollar_Dict__379880752 gopurs_runtime.Value
var once_Main_To_dollar_Dict__379880752 sync.Once

func Get_Main_To_dollar_Dict__379880752() gopurs_runtime.Value {
	once_Main_To_dollar_Dict__379880752.Do(func() {
		cache_Main_To_dollar_Dict__379880752 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_To_dollar_Dict__379880752(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_To_dollar_Dict__379880752
}

var cache_Main_To_dollar_Dict__2888513633 gopurs_runtime.Value
var once_Main_To_dollar_Dict__2888513633 sync.Once

func Get_Main_To_dollar_Dict__2888513633() gopurs_runtime.Value {
	once_Main_To_dollar_Dict__2888513633.Do(func() {
		cache_Main_To_dollar_Dict__2888513633 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_To_dollar_Dict__2888513633(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_To_dollar_Dict__2888513633
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
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Pair_prime_(uint32(x_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_Pair_prime_
}

var cache_Main_Pair_prime___745695676 gopurs_runtime.Value
var once_Main_Pair_prime___745695676 sync.Once

func Get_Main_Pair_prime___745695676() gopurs_runtime.Value {
	once_Main_Pair_prime___745695676.Do(func() {
		cache_Main_Pair_prime___745695676 = gopurs_runtime.Func(func(x_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Pair_prime___745695676(uint32(x_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_Pair_prime___745695676
}

var cache_Main_Pair_prime___1296481677 gopurs_runtime.Value
var once_Main_Pair_prime___1296481677 sync.Once

func Get_Main_Pair_prime___1296481677() gopurs_runtime.Value {
	once_Main_Pair_prime___1296481677.Do(func() {
		cache_Main_Pair_prime___1296481677 = gopurs_runtime.Func(func(x_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Pair_prime___1296481677(uint32(x_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_Pair_prime___1296481677
}

var cache_Main_to2 gopurs_runtime.Value
var once_Main_to2 sync.Once

func Get_Main_to2() gopurs_runtime.Value {
	once_Main_to2.Do(func() {
		cache_Main_to2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_to2
}

var cache_Main_to1 gopurs_runtime.Value
var once_Main_to1 sync.Once

func Get_Main_to1() gopurs_runtime.Value {
	once_Main_to1.Do(func() {
		cache_Main_to1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)), UnsafePtr: nil}
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

func Call_Main_To_dollar_Dict(x_0_loop struct {
}) uint32 {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return uint32(func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}().IntVal)
}

func Call_Main_To_dollar_Dict__379880752(x_0_loop struct {
}) uint32 {
To_dollar_Dict__379880752:
	for {
		if false {
			continue To_dollar_Dict__379880752
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return uint32(func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)
	}
}

func Call_Main_To_dollar_Dict__2888513633(x_0_loop struct {
}) uint32 {
To_dollar_Dict__2888513633:
	for {
		if false {
			continue To_dollar_Dict__2888513633
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return uint32(func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)
	}
}

func Call_Main_Pair_prime_(x_0_loop uint32) uint32 {
	var x_0 uint32 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Pair_prime___745695676(x_unused_0_loop uint32) uint32 {
Pair_prime___745695676:
	for {
		if false {
			continue Pair_prime___745695676
		}
		var x_unused_0 uint32 = x_unused_0_loop
		_ = x_unused_0
		return 893478516
	}
}

func Call_Main_Pair_prime___1296481677(x_unused_0_loop uint32) uint32 {
Pair_prime___1296481677:
	for {
		if false {
			continue Pair_prime___1296481677
		}
		var x_unused_0 uint32 = x_unused_0_loop
		_ = x_unused_0
		return 893478516
	}
}
