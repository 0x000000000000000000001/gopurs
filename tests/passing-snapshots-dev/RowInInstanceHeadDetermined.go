package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Empty gopurs_runtime.Value
var once_Main_Empty sync.Once

func Get_Main_Empty() gopurs_runtime.Value {
	once_Main_Empty.Do(func() {
		cache_Main_Empty = gopurs_runtime.Value{Type: 9, IntVal: int64(2697813931), UnsafePtr: nil}
	})
	return cache_Main_Empty
}

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Value{Type: 9, IntVal: int64(322902991), UnsafePtr: nil}
	})
	return cache_Main_Cons
}

var cache_Main_transitive gopurs_runtime.Value
var once_Main_transitive sync.Once

func Get_Main_transitive() gopurs_runtime.Value {
	once_Main_transitive.Do(func() {
		cache_Main_transitive = gopurs_runtime.Value{Type: 9, IntVal: 4241826851, UnsafePtr: unsafe.Pointer(Rebox_Main_3392908154_1025015703((&Constructor_Main_Transitive[uint32, struct {

			/*gopurs:record-v1:0:*/
		}, struct {

			/*gopurs:record-v1:0:*/
		}]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})})))}
	})
	return cache_Main_transitive
}

var cache_Main_simple1 gopurs_runtime.Value
var once_Main_simple1 sync.Once

func Get_Main_simple1() gopurs_runtime.Value {
	once_Main_simple1.Do(func() {
		cache_Main_simple1 = gopurs_runtime.Value{Type: 9, IntVal: 2086942032, UnsafePtr: unsafe.Pointer(Rebox_Main_866733052_373320770((&Constructor_Main_Simple[uint32, struct {
			foo uint32
			/*gopurs:record-v1:1:0066006f006f*/
		}]{1, gopurs_runtime.Func(func(cons_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict1("foo", cons_0)
		})})))}
	})
	return cache_Main_simple1
}

var cache_Main_simple0 gopurs_runtime.Value
var once_Main_simple0 sync.Once

func Get_Main_simple0() gopurs_runtime.Value {
	once_Main_simple0.Do(func() {
		cache_Main_simple0 = gopurs_runtime.Value{Type: 9, IntVal: 2086942032, UnsafePtr: unsafe.Pointer(Rebox_Main_2568002420_373320770((&Constructor_Main_Simple[uint32, struct {

			/*gopurs:record-v1:0:*/
		}]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})})))}
	})
	return cache_Main_simple0
}

var cache_Main_multipleDeterminers gopurs_runtime.Value
var once_Main_multipleDeterminers sync.Once

func Get_Main_multipleDeterminers() gopurs_runtime.Value {
	once_Main_multipleDeterminers.Do(func() {
		cache_Main_multipleDeterminers = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_multipleDeterminers
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_determinedCycle gopurs_runtime.Value
var once_Main_determinedCycle sync.Once

func Get_Main_determinedCycle() gopurs_runtime.Value {
	once_Main_determinedCycle.Do(func() {
		cache_Main_determinedCycle = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_determinedCycle
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_d(gopurs_runtime.CoerceToStruct[Constructor_Main_Transitive[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_d
}

var cache_Main_cyclic gopurs_runtime.Value
var once_Main_cyclic sync.Once

func Get_Main_cyclic() gopurs_runtime.Value {
	once_Main_cyclic.Do(func() {
		cache_Main_cyclic = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_cyclic
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c(gopurs_runtime.CoerceToStruct[Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_c
}

type Constructor_Main_Empty struct {
	Rc uint32
}

type Constructor_Main_Cons struct {
	Rc uint32
}

type Constructor_Main_Transitive[T_a any, T_b any, T_c any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f005400720061006e007300690074006900760065
func init() {
	gopurs_runtime.StructGetters[4241826851] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Transitive[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "d":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Transitive: " + key)
		}
	}
}

type Constructor_Main_Simple[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00530069006d0070006c0065
func init() {
	gopurs_runtime.StructGetters[2086942032] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "c":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Simple: " + key)
		}
	}
}

type Constructor_Main_MultipleDeterminers[T_a any, T_b any, T_c any, T_d any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d0075006c007400690070006c006500440065007400650072006d0069006e006500720073
func init() {
	gopurs_runtime.StructGetters[1290835010] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleDeterminers[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleDeterminers: " + key)
		}
	}
}

type Constructor_Main_DeterminedCycle[T_a any, T_b any, T_c any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00440065007400650072006d0069006e00650064004300790063006c0065
func init() {
	gopurs_runtime.StructGetters[1744313767] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_DeterminedCycle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_DeterminedCycle: " + key)
		}
	}
}

type Constructor_Main_Cyclic[T_a any, T_b any, T_c any, T_d any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004300790063006c00690063
func init() {
	gopurs_runtime.StructGetters[1751764097] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Cyclic[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Cyclic: " + key)
		}
	}
}

func Call_Main_d(dict_0_loop *Constructor_Main_Transitive[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Transitive[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_c(dict_0_loop *Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_2568002420_373320770(in *Constructor_Main_Simple[uint32, struct {

	/*gopurs:record-v1:0:*/
}]) *Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3392908154_1025015703(in *Constructor_Main_Transitive[uint32, struct {

	/*gopurs:record-v1:0:*/
}, struct {

	/*gopurs:record-v1:0:*/
}]) *Constructor_Main_Transitive[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Transitive[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_866733052_373320770(in *Constructor_Main_Simple[uint32, struct {
	foo uint32
	/*gopurs:record-v1:1:0066006f006f*/
}]) *Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Simple[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
