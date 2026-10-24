package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test(_dollar___unused_0_box, v_1_box, v1_2_box))
		})
	})
	return cache_Main_test
}

var cache_Main_eqAA gopurs_runtime.Value
var once_Main_eqAA sync.Once

func Get_Main_eqAA() gopurs_runtime.Value {
	once_Main_eqAA.Do(func() {
		cache_Main_eqAA = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_eqAA
}

var cache_Main_runTest gopurs_runtime.Value
var once_Main_runTest sync.Once

func Get_Main_runTest() gopurs_runtime.Value {
	once_Main_runTest.Do(func() {
		cache_Main_runTest = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_runTest(a_0_box))
		})
	})
	return cache_Main_runTest
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_runTest(gopurs_runtime.Float(0.0))))
	})
	return cache_Main_main
}

type Constructor_Main_EQ[T_a any, T_b any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00450051
func init() {
	gopurs_runtime.StructGetters[3323825930] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_EQ[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_EQ: " + key)
		}
	}
}

func Call_Main_test(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) string {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var v1_2 gopurs_runtime.Value = v1_2_loop
	_ = v1_2
	return "Done"
}

func Call_Main_runTest(a_0_loop gopurs_runtime.Value) string {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return "Done"
}
