package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ClassA_dollar_Dict gopurs_runtime.Value
var once_Main_ClassA_dollar_Dict sync.Once

func Get_Main_ClassA_dollar_Dict() gopurs_runtime.Value {
	once_Main_ClassA_dollar_Dict.Do(func() {
		cache_Main_ClassA_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ClassA_dollar_Dict(func() struct {

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
	return cache_Main_ClassA_dollar_Dict
}

var cache_Main_ClassB_dollar_Dict gopurs_runtime.Value
var once_Main_ClassB_dollar_Dict sync.Once

func Get_Main_ClassB_dollar_Dict() gopurs_runtime.Value {
	once_Main_ClassB_dollar_Dict.Do(func() {
		cache_Main_ClassB_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1454680786, UnsafePtr: unsafe.Pointer(Call_Main_ClassB_dollar_Dict(func() struct {
				ClassA0 gopurs_runtime.Value
				/*gopurs:record-v1:1:0043006c00610073007300410030*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					ClassA0 gopurs_runtime.Value
					/*gopurs:record-v1:1:0043006c00610073007300410030*/
				}{}
				clone.ClassA0 = gopurs_runtime.RecordGet(orig, "ClassA0")
				return clone
			}()))}
		})
	})
	return cache_Main_ClassB_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_a
}

var cache_Main_b gopurs_runtime.Value
var once_Main_b sync.Once

func Get_Main_b() gopurs_runtime.Value {
	once_Main_b.Do(func() {
		cache_Main_b = gopurs_runtime.Value{Type: 9, IntVal: 1454680786, UnsafePtr: unsafe.Pointer((&Constructor_Main_ClassB[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_b
}

type Constructor_Main_ClassA[T_t any, T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043006c0061007300730041
func init() {
	gopurs_runtime.StructGetters[3328906353] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ClassA[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_ClassA: " + key)
		}
	}
}

type Constructor_Main_ClassB[T_t any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043006c0061007300730042
func init() {
	gopurs_runtime.StructGetters[1454680786] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ClassB[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "ClassA0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ClassB: " + key)
		}
	}
}

func Call_Main_ClassA_dollar_Dict(x_0_loop struct {

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

func Call_Main_ClassB_dollar_Dict(x_0_loop struct {
	ClassA0 gopurs_runtime.Value
	/*gopurs:record-v1:1:0043006c00610073007300410030*/
}) *Constructor_Main_ClassB[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		ClassA0 gopurs_runtime.Value
		/*gopurs:record-v1:1:0043006c00610073007300410030*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_ClassB[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("ClassA0", orig.ClassA0)
	}())
}
