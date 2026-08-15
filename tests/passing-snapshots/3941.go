package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_TwoParams_dollar_Dict gopurs_runtime.Value
var once_Main_TwoParams_dollar_Dict sync.Once

func Get_Main_TwoParams_dollar_Dict() gopurs_runtime.Value {
	once_Main_TwoParams_dollar_Dict.Do(func() {
		cache_Main_TwoParams_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_TwoParams_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_TwoParams_dollar_Dict
}

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
			return Call_Main_go__func(gopurs_runtime.CoerceToStruct[Constructor_Main_TwoParams](dict_0_box))
		})
	})
	return cache_Main_go__func
}

var cache_Main_equals gopurs_runtime.Value
var once_Main_equals sync.Once

func Get_Main_equals() gopurs_runtime.Value {
	once_Main_equals.Do(func() {
		cache_Main_equals = gopurs_runtime.Value{Type: 9, IntVal: 1199216238, UnsafePtr: unsafe.Pointer((&Constructor_Main_TwoParams{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
		cache_Main_any = gopurs_runtime.Value{Type: 9, IntVal: 1199216238, UnsafePtr: unsafe.Pointer((&Constructor_Main_TwoParams{1, Get_Unsafe_Coerce_unsafeCoerce()}))}
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

type Constructor_Main_TwoParams struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1199216238] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_TwoParams)(ptr)
		_ = c
		switch key {
		case "func":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_TwoParams: " + key)
		}
	}
}

func Call_Main_TwoParams_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_go__func(dict_0_loop *Constructor_Main_TwoParams) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_TwoParams = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_testEquals(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return a_0
}
