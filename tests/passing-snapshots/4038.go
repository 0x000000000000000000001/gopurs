package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_A_dollar_Dict gopurs_runtime.Value
var once_Main_A_dollar_Dict sync.Once

func Get_Main_A_dollar_Dict() gopurs_runtime.Value {
	once_Main_A_dollar_Dict.Do(func() {
		cache_Main_A_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_A_dollar_Dict(func() struct {

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
	return cache_Main_A_dollar_Dict
}

var cache_Main_B_dollar_Dict gopurs_runtime.Value
var once_Main_B_dollar_Dict sync.Once

func Get_Main_B_dollar_Dict() gopurs_runtime.Value {
	once_Main_B_dollar_Dict.Do(func() {
		cache_Main_B_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(Call_Main_B_dollar_Dict(func() struct {
				A0 gopurs_runtime.Value
				/*gopurs:record-v1:1:00410030*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					A0 gopurs_runtime.Value
					/*gopurs:record-v1:1:00410030*/
				}{}
				clone.A0 = gopurs_runtime.RecordGet(orig, "A0")
				return clone
			}()))}
		})
	})
	return cache_Main_B_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_A struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0041
func init() {
	gopurs_runtime.StructGetters[4219254943] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_A)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_A: " + key)
		}
	}
}

type Constructor_Main_B struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0042
func init() {
	gopurs_runtime.StructGetters[4250879068] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_B)(ptr)
		_ = c
		switch key {
		case "A0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_B: " + key)
		}
	}
}

func Call_Main_A_dollar_Dict(x_0_loop struct {

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

func Call_Main_B_dollar_Dict(x_0_loop struct {
	A0 gopurs_runtime.Value
	/*gopurs:record-v1:1:00410030*/
}) *Constructor_Main_B {
	var x_0 struct {
		A0 gopurs_runtime.Value
		/*gopurs:record-v1:1:00410030*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_B](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("A0", orig.A0)
	}())
}
