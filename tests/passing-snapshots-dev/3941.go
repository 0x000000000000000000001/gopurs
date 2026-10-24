package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_go__func gopurs_runtime.Value
var once_Main_go__func sync.Once

func Get_Main_go__func() gopurs_runtime.Value {
	once_Main_go__func.Do(func() {
		cache_Main_go__func = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_go__func(gopurs_runtime.CoerceToStruct[Constructor_Main_TwoParams[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_go__func
}

var cache_Main_equals gopurs_runtime.Value
var once_Main_equals sync.Once

func Get_Main_equals() gopurs_runtime.Value {
	once_Main_equals.Do(func() {
		cache_Main_equals = gopurs_runtime.Value{Type: 9, IntVal: 1199216238, UnsafePtr: unsafe.Pointer((&Constructor_Main_TwoParams[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return a_0
		})}))}
	})
	return cache_Main_equals
}

var cache_Main_testEquals gopurs_runtime.Value
var once_Main_testEquals sync.Once

func Get_Main_testEquals() gopurs_runtime.Value {
	once_Main_testEquals.Do(func() {
		cache_Main_testEquals = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testEquals(a_0_box)
		})
	})
	return cache_Main_testEquals
}

var cache_Main_any gopurs_runtime.Value
var once_Main_any sync.Once

func Get_Main_any() gopurs_runtime.Value {
	once_Main_any.Do(func() {
		cache_Main_any = gopurs_runtime.Value{Type: 9, IntVal: 1199216238, UnsafePtr: unsafe.Pointer((&Constructor_Main_TwoParams[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Unsafe_Coerce_unsafeCoerce()}))}
	})
	return cache_Main_any
}

var cache_Main_testAny gopurs_runtime.Value
var once_Main_testAny sync.Once

func Get_Main_testAny() gopurs_runtime.Value {
	once_Main_testAny.Do(func() {
		cache_Main_testAny = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_testAny
}

var cache_Main_thisShouldBeCompiled gopurs_runtime.Value
var once_Main_thisShouldBeCompiled sync.Once

func Get_Main_thisShouldBeCompiled() gopurs_runtime.Value {
	once_Main_thisShouldBeCompiled.Do(func() {
		cache_Main_thisShouldBeCompiled = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_thisShouldBeCompiled
}

type Constructor_Main_TwoParams[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00540077006f0050006100720061006d0073
func init() {
	gopurs_runtime.StructGetters[1199216238] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_TwoParams[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "func":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_TwoParams: " + key)
		}
	}
}

func Call_Main_go__func(dict_0_loop *Constructor_Main_TwoParams[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_TwoParams[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_testEquals(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return a_0
}
