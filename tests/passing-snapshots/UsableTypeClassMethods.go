package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_C4_dollar_Dict gopurs_runtime.Value
var once_Main_C4_dollar_Dict sync.Once

func Get_Main_C4_dollar_Dict() gopurs_runtime.Value {
	once_Main_C4_dollar_Dict.Do(func() {
		cache_Main_C4_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C4_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_C4_dollar_Dict
}

var cache_Main_C3_dollar_Dict gopurs_runtime.Value
var once_Main_C3_dollar_Dict sync.Once

func Get_Main_C3_dollar_Dict() gopurs_runtime.Value {
	once_Main_C3_dollar_Dict.Do(func() {
		cache_Main_C3_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C3_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_C3_dollar_Dict
}

var cache_Main_C2_dollar_Dict gopurs_runtime.Value
var once_Main_C2_dollar_Dict sync.Once

func Get_Main_C2_dollar_Dict() gopurs_runtime.Value {
	once_Main_C2_dollar_Dict.Do(func() {
		cache_Main_C2_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C2_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_C2_dollar_Dict
}

var cache_Main_C1_dollar_Dict gopurs_runtime.Value
var once_Main_C1_dollar_Dict sync.Once

func Get_Main_C1_dollar_Dict() gopurs_runtime.Value {
	once_Main_C1_dollar_Dict.Do(func() {
		cache_Main_C1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C1_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_C1_dollar_Dict
}

var cache_Main_C0_dollar_Dict gopurs_runtime.Value
var once_Main_C0_dollar_Dict sync.Once

func Get_Main_C0_dollar_Dict() gopurs_runtime.Value {
	once_Main_C0_dollar_Dict.Do(func() {
		cache_Main_C0_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C0_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_C0_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_c4 gopurs_runtime.Value
var once_Main_c4 sync.Once

func Get_Main_c4() gopurs_runtime.Value {
	once_Main_c4.Do(func() {
		cache_Main_c4 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c4(dict_0_box)
		})
	})
	return cache_Main_c4
}

var cache_Main_c3_prime__prime__prime__prime_ gopurs_runtime.Value
var once_Main_c3_prime__prime__prime__prime_ sync.Once

func Get_Main_c3_prime__prime__prime__prime_() gopurs_runtime.Value {
	once_Main_c3_prime__prime__prime__prime_.Do(func() {
		cache_Main_c3_prime__prime__prime__prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c3_prime__prime__prime__prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C3](dict_0_box))
		})
	})
	return cache_Main_c3_prime__prime__prime__prime_
}

var cache_Main_c3_prime__prime__prime_ gopurs_runtime.Value
var once_Main_c3_prime__prime__prime_ sync.Once

func Get_Main_c3_prime__prime__prime_() gopurs_runtime.Value {
	once_Main_c3_prime__prime__prime_.Do(func() {
		cache_Main_c3_prime__prime__prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c3_prime__prime__prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C3](dict_0_box))
		})
	})
	return cache_Main_c3_prime__prime__prime_
}

var cache_Main_c3_prime__prime_ gopurs_runtime.Value
var once_Main_c3_prime__prime_ sync.Once

func Get_Main_c3_prime__prime_() gopurs_runtime.Value {
	once_Main_c3_prime__prime_.Do(func() {
		cache_Main_c3_prime__prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c3_prime__prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C3](dict_0_box))
		})
	})
	return cache_Main_c3_prime__prime_
}

var cache_Main_c3_prime_ gopurs_runtime.Value
var once_Main_c3_prime_ sync.Once

func Get_Main_c3_prime_() gopurs_runtime.Value {
	once_Main_c3_prime_.Do(func() {
		cache_Main_c3_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c3_prime_(dict_0_box)
		})
	})
	return cache_Main_c3_prime_
}

var cache_Main_c3 gopurs_runtime.Value
var once_Main_c3 sync.Once

func Get_Main_c3() gopurs_runtime.Value {
	once_Main_c3.Do(func() {
		cache_Main_c3 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c3(dict_0_box)
		})
	})
	return cache_Main_c3
}

var cache_Main_c2_prime__prime__prime_ gopurs_runtime.Value
var once_Main_c2_prime__prime__prime_ sync.Once

func Get_Main_c2_prime__prime__prime_() gopurs_runtime.Value {
	once_Main_c2_prime__prime__prime_.Do(func() {
		cache_Main_c2_prime__prime__prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c2_prime__prime__prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C2](dict_0_box))
		})
	})
	return cache_Main_c2_prime__prime__prime_
}

var cache_Main_c2_prime__prime_ gopurs_runtime.Value
var once_Main_c2_prime__prime_ sync.Once

func Get_Main_c2_prime__prime_() gopurs_runtime.Value {
	once_Main_c2_prime__prime_.Do(func() {
		cache_Main_c2_prime__prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c2_prime__prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C2](dict_0_box))
		})
	})
	return cache_Main_c2_prime__prime_
}

var cache_Main_c2_prime_ gopurs_runtime.Value
var once_Main_c2_prime_ sync.Once

func Get_Main_c2_prime_() gopurs_runtime.Value {
	once_Main_c2_prime_.Do(func() {
		cache_Main_c2_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c2_prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C2](dict_0_box))
		})
	})
	return cache_Main_c2_prime_
}

var cache_Main_c2 gopurs_runtime.Value
var once_Main_c2 sync.Once

func Get_Main_c2() gopurs_runtime.Value {
	once_Main_c2.Do(func() {
		cache_Main_c2 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c2(dict_0_box)
		})
	})
	return cache_Main_c2
}

var cache_Main_c1_prime_ gopurs_runtime.Value
var once_Main_c1_prime_ sync.Once

func Get_Main_c1_prime_() gopurs_runtime.Value {
	once_Main_c1_prime_.Do(func() {
		cache_Main_c1_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c1_prime_(gopurs_runtime.CoerceToStruct[Constructor_Main_C1](dict_0_box))
		})
	})
	return cache_Main_c1_prime_
}

var cache_Main_c1 gopurs_runtime.Value
var once_Main_c1 sync.Once

func Get_Main_c1() gopurs_runtime.Value {
	once_Main_c1.Do(func() {
		cache_Main_c1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c1(dict_0_box)
		})
	})
	return cache_Main_c1
}

var cache_Main_c0 gopurs_runtime.Value
var once_Main_c0 sync.Once

func Get_Main_c0() gopurs_runtime.Value {
	once_Main_c0.Do(func() {
		cache_Main_c0 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c0(gopurs_runtime.CoerceToStruct[Constructor_Main_C0](dict_0_box))
		})
	})
	return cache_Main_c0
}

type Constructor_Main_C4 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[719882249] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C4)(ptr)
		_ = c
		switch key {
		case "c4":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_C4: " + key)
		}
	}
}

type Constructor_Main_C3 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2766571374] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C3)(ptr)
		_ = c
		switch key {
		case "c3":
			return gopurs_runtime.Box(c.V0)
		case "c3'":
			return gopurs_runtime.Box(c.V1)
		case "c3''":
			return gopurs_runtime.Box(c.V2)
		case "c3'''":
			return gopurs_runtime.Box(c.V3)
		case "c3''''":
			return gopurs_runtime.Box(c.V4)
		default:
			panic("Key not found in dictionary Constructor_Main_C3: " + key)
		}
	}
}

type Constructor_Main_C2 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2226553295] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C2)(ptr)
		_ = c
		switch key {
		case "c2":
			return gopurs_runtime.Box(c.V0)
		case "c2'":
			return gopurs_runtime.Box(c.V1)
		case "c2''":
			return gopurs_runtime.Box(c.V2)
		case "c2'''":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Main_C2: " + key)
		}
	}
}

type Constructor_Main_C1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4264042284] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C1)(ptr)
		_ = c
		switch key {
		case "c1":
			return gopurs_runtime.Box(c.V0)
		case "c1'":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_C1: " + key)
		}
	}
}

type Constructor_Main_C0 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1613519245] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C0)(ptr)
		_ = c
		switch key {
		case "c0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_C0: " + key)
		}
	}
}

func Call_Main_C4_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C3_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C2_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C1_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C0_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_c4(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "c4")
}

func Call_Main_c3_prime__prime__prime__prime_(dict_0_loop *Constructor_Main_C3) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C3 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V4)
}

func Call_Main_c3_prime__prime__prime_(dict_0_loop *Constructor_Main_C3) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C3 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V3)
}

func Call_Main_c3_prime__prime_(dict_0_loop *Constructor_Main_C3) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C3 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V2)
}

func Call_Main_c3_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "c3'")
}

func Call_Main_c3(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "c3")
}

func Call_Main_c2_prime__prime__prime_(dict_0_loop *Constructor_Main_C2) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C2 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V3)
}

func Call_Main_c2_prime__prime_(dict_0_loop *Constructor_Main_C2) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C2 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V2)
}

func Call_Main_c2_prime_(dict_0_loop *Constructor_Main_C2) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C2 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_c2(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "c2")
}

func Call_Main_c1_prime_(dict_0_loop *Constructor_Main_C1) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C1 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_c1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "c1")
}

func Call_Main_c0(dict_0_loop *Constructor_Main_C0) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C0 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
