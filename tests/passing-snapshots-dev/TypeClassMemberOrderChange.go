package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
		cache_Main_testBoolean = gopurs_runtime.Value{Type: 9, IntVal: 3423238824, UnsafePtr: unsafe.Pointer(Rebox_Main_1531651672_2549914428((&Constructor_Main_Test[bool]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return y_1
		}), true})))}
	})
	return cache_Main_testBoolean
}

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn(gopurs_runtime.CoerceToStruct[Constructor_Main_Test[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_fn
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true"))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Test[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0054006500730074
func init() {
	gopurs_runtime.StructGetters[3423238824] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test[gopurs_runtime.Value])(ptr)
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

func Call_Main_val(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "val")
}

func Call_Main_fn(dict_0_loop *Constructor_Main_Test[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Test[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1531651672_2549914428(in *Constructor_Main_Test[bool]) *Constructor_Main_Test[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Bool(in.V1)
	return out
}
