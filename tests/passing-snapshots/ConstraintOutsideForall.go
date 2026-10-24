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
			return Call_Main_Test_dollar_Dict(func() struct {

				/*gopurs:record-v1:0:*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {

					/*gopurs:record-v1:0:*/
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Test_dollar_Dict
}

var cache_Main_Test_dollar_Dict__3889889207 gopurs_runtime.Value
var once_Main_Test_dollar_Dict__3889889207 sync.Once

func Get_Main_Test_dollar_Dict__3889889207() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict__3889889207.Do(func() {
		cache_Main_Test_dollar_Dict__3889889207 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test_dollar_Dict__3889889207(func() struct {

				/*gopurs:record-v1:0:*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {

					/*gopurs:record-v1:0:*/
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Test_dollar_Dict__3889889207
}

var cache_Main_testUnit gopurs_runtime.Value
var once_Main_testUnit sync.Once

func Get_Main_testUnit() gopurs_runtime.Value {
	once_Main_testUnit.Do(func() {
		cache_Main_testUnit = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_testUnit
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(_dollar___unused_0_box, a_1_box)
		})
	})
	return cache_Main_test
}

var cache_Main_test__2521650171 gopurs_runtime.Value
var once_Main_test__2521650171 sync.Once

func Get_Main_test__2521650171() gopurs_runtime.Value {
	once_Main_test__2521650171.Do(func() {
		cache_Main_test__2521650171 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test__2521650171(a_0_box.StrVal()))
		})
	})
	return cache_Main_test__2521650171
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0054006500730074
func init() {
	gopurs_runtime.StructGetters[3423238824] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Test: " + key)
		}
	}
}

func Call_Main_Test_dollar_Dict(x_0_loop struct {

	/*gopurs:record-v1:0:*/
}) gopurs_runtime.Value {
	var x_0 struct {

		/*gopurs:record-v1:0:*/
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_Test_dollar_Dict__3889889207(x_0_loop struct {

	/*gopurs:record-v1:0:*/
}) gopurs_runtime.Value {
Test_dollar_Dict__3889889207:
	for {
		if false {
			continue Test_dollar_Dict__3889889207
		}
		var x_0 struct {

			/*gopurs:record-v1:0:*/
		} = x_0_loop
		_ = x_0
		return func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	}
}

func Call_Main_test(_dollar___unused_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return a_1
}

func Call_Main_test__2521650171(a_0_loop string) string {
test__2521650171:
	for {
		if false {
			continue test__2521650171
		}
		var a_0 string = a_0_loop
		_ = a_0
		return gopurs_runtime.Str(a_0).StrVal()
	}
}
