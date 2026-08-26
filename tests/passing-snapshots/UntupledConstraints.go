package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Nonsense_dollar_Dict gopurs_runtime.Value
var once_Main_Nonsense_dollar_Dict sync.Once

func Get_Main_Nonsense_dollar_Dict() gopurs_runtime.Value {
	once_Main_Nonsense_dollar_Dict.Do(func() {
		cache_Main_Nonsense_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Nonsense_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Nonsense_dollar_Dict
}

var cache_Main_Box gopurs_runtime.Value
var once_Main_Box sync.Once

func Get_Main_Box() gopurs_runtime.Value {
	once_Main_Box.Do(func() {
		cache_Main_Box = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Box
}

var cache_Main_strangeThing gopurs_runtime.Value
var once_Main_strangeThing sync.Once

func Get_Main_strangeThing() gopurs_runtime.Value {
	once_Main_strangeThing.Do(func() {
		cache_Main_strangeThing = gopurs_runtime.Func3(func(dictSemigroup_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_strangeThing(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box), x_1_box, y_2_box)
		})
	})
	return cache_Main_strangeThing
}

var cache_Main_showBox gopurs_runtime.Value
var once_Main_showBox sync.Once

func Get_Main_showBox() gopurs_runtime.Value {
	once_Main_showBox.Do(func() {
		cache_Main_showBox = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showBox(dictShow_0_box)
		})
	})
	return cache_Main_showBox
}

var cache_Main_method gopurs_runtime.Value
var once_Main_method sync.Once

func Get_Main_method() gopurs_runtime.Value {
	once_Main_method.Do(func() {
		cache_Main_method = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_method(gopurs_runtime.CoerceToStruct[Constructor_Main_Nonsense[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_method
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Box[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Nonsense[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2074529919] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Nonsense[any])(ptr)
		_ = c
		switch key {
		case "Show0":
			return gopurs_runtime.Box(c.V0)
		case "method":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Nonsense: " + key)
		}
	}
}

func Call_Main_Nonsense_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_strangeThing(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var y_2 gopurs_runtime.Value = y_2_loop
	_ = y_2
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), x_1, y_2)
}

func Call_Main_showBox(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(("Box ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()))
	})}))}
}

func Call_Main_method(dict_0_loop *Constructor_Main_Nonsense[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Nonsense[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}
