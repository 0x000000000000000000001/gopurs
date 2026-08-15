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

var cache_Main_val gopurs_runtime.Value
var once_Main_val sync.Once

func Get_Main_val() gopurs_runtime.Value {
	once_Main_val.Do(func() {
		cache_Main_val = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_val(dict_0_box)
		})
	})
	return cache_Main_val
}

var cache_Main_testBoolean gopurs_runtime.Value
var once_Main_testBoolean sync.Once

func Get_Main_testBoolean() gopurs_runtime.Value {
	once_Main_testBoolean.Do(func() {
		cache_Main_testBoolean = gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((y_1.IntVal) != (0))
			})
		}), gopurs_runtime.Bool(true)}))}
	})
	return cache_Main_testBoolean
}

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn(gopurs_runtime.CoerceToStruct[Constructor_Main_Test](dict_0_box))
		})
	})
	return cache_Main_fn
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true"))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_fn__88475274 gopurs_runtime.Value
var once_Main_fn__88475274 sync.Once

func Get_Main_fn__88475274() gopurs_runtime.Value {
	once_Main_fn__88475274.Do(func() {
		cache_Main_fn__88475274 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn__88475274(gopurs_runtime.CoerceToStruct[Constructor_Main_Test](dict_0_box))
		})
	})
	return cache_Main_fn__88475274
}

type Constructor_Main_Test struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3423238824] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test)(ptr)
		_ = c
		switch key {
		case "fn":
			return gopurs_runtime.Box(c.V0)
		case "val":
			return gopurs_runtime.Box(c.V1)
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

func Call_Main_val(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "val")
}

func Call_Main_fn(dict_0_loop *Constructor_Main_Test) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Test = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_fn__88475274(dict_0_loop *Constructor_Main_Test) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Test = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
