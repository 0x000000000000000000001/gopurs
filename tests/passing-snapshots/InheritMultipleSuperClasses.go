package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

var cache_Main_Eg2_dollar_Dict gopurs_runtime.Value
var once_Main_Eg2_dollar_Dict sync.Once

func Get_Main_Eg2_dollar_Dict() gopurs_runtime.Value {
	once_Main_Eg2_dollar_Dict.Do(func() {
		cache_Main_Eg2_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Eg2_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Eg2_dollar_Dict
}

var cache_Main_Eg1_dollar_Dict gopurs_runtime.Value
var once_Main_Eg1_dollar_Dict sync.Once

func Get_Main_Eg1_dollar_Dict() gopurs_runtime.Value {
	once_Main_Eg1_dollar_Dict.Do(func() {
		cache_Main_Eg1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Eg1_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Eg1_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_g2 gopurs_runtime.Value
var once_Main_g2 sync.Once

func Get_Main_g2() gopurs_runtime.Value {
	once_Main_g2.Do(func() {
		cache_Main_g2 = gopurs_runtime.Func(func(dictEg2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_g2(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg2](dictEg2_0_box))
		})
	})
	return cache_Main_g2
}

var cache_Main_g1 gopurs_runtime.Value
var once_Main_g1 sync.Once

func Get_Main_g1() gopurs_runtime.Value {
	once_Main_g1.Do(func() {
		cache_Main_g1 = gopurs_runtime.Func(func(dictEg1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_g1(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg1](dictEg1_0_box))
		})
	})
	return cache_Main_g1
}

var cache_Main_f2 gopurs_runtime.Value
var once_Main_f2 sync.Once

func Get_Main_f2() gopurs_runtime.Value {
	once_Main_f2.Do(func() {
		cache_Main_f2 = gopurs_runtime.Func(func(dictEg2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f2(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg2](dictEg2_0_box))
		})
	})
	return cache_Main_f2
}

var cache_Main_f1 gopurs_runtime.Value
var once_Main_f1 sync.Once

func Get_Main_f1() gopurs_runtime.Value {
	once_Main_f1.Do(func() {
		cache_Main_f1 = gopurs_runtime.Func(func(dictEg1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f1(gopurs_runtime.CoerceToStruct[Constructor_Main_Eg1](dictEg1_0_box))
		})
	})
	return cache_Main_f1
}

type Constructor_Main_Eg2 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2207147950] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eg2)(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "Functor1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Eg2: " + key)
		}
	}
}

type Constructor_Main_Eg1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1828296749] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eg1)(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "Functor1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Eg1: " + key)
		}
	}
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Eg2_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Eg1_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_g2(dictEg2_0_loop *Constructor_Main_Eg2) gopurs_runtime.Value {
	var dictEg2_0 *Constructor_Main_Eg2 = dictEg2_0_loop
	_ = dictEg2_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEg2_0.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Main_g1(dictEg1_0_loop *Constructor_Main_Eg1) gopurs_runtime.Value {
	var dictEg1_0 *Constructor_Main_Eg1 = dictEg1_0_loop
	_ = dictEg1_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEg1_0.V1), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Main_f2(dictEg2_0_loop *Constructor_Main_Eg2) gopurs_runtime.Value {
	var dictEg2_0 *Constructor_Main_Eg2 = dictEg2_0_loop
	_ = dictEg2_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEg2_0.V1), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Main_f1(dictEg1_0_loop *Constructor_Main_Eg1) gopurs_runtime.Value {
	var dictEg1_0 *Constructor_Main_Eg1 = dictEg1_0_loop
	_ = dictEg1_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictEg1_0.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}
