package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Test_dollar_Dict gopurs_runtime.Value
var once_Main_Test_dollar_Dict sync.Once

func Get_Main_Test_dollar_Dict() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict.Do(func() {
		cache_Main_Test_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Test_dollar_Dict
}

var cache_Main_testUnit gopurs_runtime.Value
var once_Main_testUnit sync.Once

func Get_Main_testUnit() gopurs_runtime.Value {
	once_Main_testUnit.Do(func() {
		cache_Main_testUnit = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
			orig := func() *struct {
			} {
				orig := gopurs_runtime.RecordDict0()
				_ = orig
				clone := struct {
				}{}

				return &clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
		}().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_testUnit
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(uint32(_dollar___unused_0_box.IntVal), a_1_box)
		})
	})
	return cache_Main_test
}

var cache_Main_test__3373544390 gopurs_runtime.Value
var once_Main_test__3373544390 sync.Once

func Get_Main_test__3373544390() gopurs_runtime.Value {
	once_Main_test__3373544390.Do(func() {
		cache_Main_test__3373544390 = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test__3373544390(uint32(_dollar___unused_0_box.IntVal), a_1_box)
		})
	})
	return cache_Main_test__3373544390
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
	})
	return cache_Main_main
}

type Constructor_Main_Test[T_a any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3423238824] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test[any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Test: " + key)
		}
	}
}

func Call_Main_Test_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_test(_dollar___unused_0_loop uint32, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return a_1
}

func Call_Main_test__3373544390(_dollar___unused_0_loop uint32, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return a_1
}
