package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_X_dollar_Dict gopurs_runtime.Value
var once_Main_X_dollar_Dict sync.Once

func Get_Main_X_dollar_Dict() gopurs_runtime.Value {
	once_Main_X_dollar_Dict.Do(func() {
		cache_Main_X_dollar_Dict = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_X_dollar_Dict(func() struct {

				/*gopurs:record-v1:0:*/
			} {
				orig := x1_0_box
				_ = orig
				clone := struct {

					/*gopurs:record-v1:0:*/
				}{}

				return clone
			}())
		})
	})
	return cache_Main_X_dollar_Dict
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_x(_dollar___unused_0_box)
		})
	})
	return cache_Main_x
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_X[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0058
func init() {
	gopurs_runtime.StructGetters[1409933510] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_X[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_X: " + key)
		}
	}
}

func Call_Main_X_dollar_Dict(x1_0_loop struct {

	/*gopurs:record-v1:0:*/
}) gopurs_runtime.Value {
	var x1_0 struct {

		/*gopurs:record-v1:0:*/
	} = x1_0_loop
	_ = x1_0
	return func() gopurs_runtime.Value {
		orig := x1_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_x(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return func() gopurs_runtime.Value {
		orig := struct {

			/*gopurs:record-v1:0:*/
		}{}
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}
