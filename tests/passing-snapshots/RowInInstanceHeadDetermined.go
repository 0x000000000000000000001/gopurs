package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Transitive_dollar_Dict gopurs_runtime.Value
var once_Main_Transitive_dollar_Dict sync.Once

func Get_Main_Transitive_dollar_Dict() gopurs_runtime.Value {
	once_Main_Transitive_dollar_Dict.Do(func() {
		cache_Main_Transitive_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Transitive_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Transitive_dollar_Dict
}

var cache_Main_Simple_dollar_Dict gopurs_runtime.Value
var once_Main_Simple_dollar_Dict sync.Once

func Get_Main_Simple_dollar_Dict() gopurs_runtime.Value {
	once_Main_Simple_dollar_Dict.Do(func() {
		cache_Main_Simple_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Simple_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Simple_dollar_Dict
}

var cache_Main_MultipleDeterminers_dollar_Dict gopurs_runtime.Value
var once_Main_MultipleDeterminers_dollar_Dict sync.Once

func Get_Main_MultipleDeterminers_dollar_Dict() gopurs_runtime.Value {
	once_Main_MultipleDeterminers_dollar_Dict.Do(func() {
		cache_Main_MultipleDeterminers_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_MultipleDeterminers_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MultipleDeterminers_dollar_Dict
}

var cache_Main_Empty gopurs_runtime.Value
var once_Main_Empty sync.Once

func Get_Main_Empty() gopurs_runtime.Value {
	once_Main_Empty.Do(func() {
		cache_Main_Empty = gopurs_runtime.Value{Type: 9, IntVal: int64(2697813931), UnsafePtr: nil}
	})
	return cache_Main_Empty
}

var cache_Main_DeterminedCycle_dollar_Dict gopurs_runtime.Value
var once_Main_DeterminedCycle_dollar_Dict sync.Once

func Get_Main_DeterminedCycle_dollar_Dict() gopurs_runtime.Value {
	once_Main_DeterminedCycle_dollar_Dict.Do(func() {
		cache_Main_DeterminedCycle_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_DeterminedCycle_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_DeterminedCycle_dollar_Dict
}

var cache_Main_Cyclic_dollar_Dict gopurs_runtime.Value
var once_Main_Cyclic_dollar_Dict sync.Once

func Get_Main_Cyclic_dollar_Dict() gopurs_runtime.Value {
	once_Main_Cyclic_dollar_Dict.Do(func() {
		cache_Main_Cyclic_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Cyclic_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Cyclic_dollar_Dict
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
		cache_Main_transitive = gopurs_runtime.Value{Type: 9, IntVal: 4241826851, UnsafePtr: unsafe.Pointer((&Constructor_Main_Transitive{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})}))}
	})
	return cache_Main_transitive
}

var cache_Main_simple1 gopurs_runtime.Value
var once_Main_simple1 sync.Once

func Get_Main_simple1() gopurs_runtime.Value {
	once_Main_simple1.Do(func() {
		cache_Main_simple1 = gopurs_runtime.Value{Type: 9, IntVal: 2086942032, UnsafePtr: unsafe.Pointer((&Constructor_Main_Simple{1, gopurs_runtime.Func(func(cons_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict1("foo", gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(cons_0.IntVal)), UnsafePtr: nil})
		})}))}
	})
	return cache_Main_simple1
}

var cache_Main_simple0 gopurs_runtime.Value
var once_Main_simple0 sync.Once

func Get_Main_simple0() gopurs_runtime.Value {
	once_Main_simple0.Do(func() {
		cache_Main_simple0 = gopurs_runtime.Value{Type: 9, IntVal: 2086942032, UnsafePtr: unsafe.Pointer((&Constructor_Main_Simple{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})}))}
	})
	return cache_Main_simple0
}

var cache_Main_multipleDeterminers gopurs_runtime.Value
var once_Main_multipleDeterminers sync.Once

func Get_Main_multipleDeterminers() gopurs_runtime.Value {
	once_Main_multipleDeterminers.Do(func() {
		cache_Main_multipleDeterminers = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
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
		cache_Main_determinedCycle = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_determinedCycle
}

var cache_Main_d gopurs_runtime.Value
var once_Main_d sync.Once

func Get_Main_d() gopurs_runtime.Value {
	once_Main_d.Do(func() {
		cache_Main_d = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_d(gopurs_runtime.CoerceToStruct[Constructor_Main_Transitive](dict_0_box))
		})
	})
	return cache_Main_d
}

var cache_Main_cyclic gopurs_runtime.Value
var once_Main_cyclic sync.Once

func Get_Main_cyclic() gopurs_runtime.Value {
	once_Main_cyclic.Do(func() {
		cache_Main_cyclic = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_cyclic
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c(gopurs_runtime.CoerceToStruct[Constructor_Main_Simple](dict_0_box))
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

type Constructor_Main_Transitive struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4241826851] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Transitive)(ptr)
		_ = c
		switch key {
		case "d":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Transitive: " + key)
		}
	}
}

type Constructor_Main_Simple struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2086942032] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Simple)(ptr)
		_ = c
		switch key {
		case "c":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Simple: " + key)
		}
	}
}

type Constructor_Main_MultipleDeterminers struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1290835010] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MultipleDeterminers)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_MultipleDeterminers: " + key)
		}
	}
}

type Constructor_Main_DeterminedCycle struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1744313767] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_DeterminedCycle)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_DeterminedCycle: " + key)
		}
	}
}

type Constructor_Main_Cyclic struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1751764097] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Cyclic)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Cyclic: " + key)
		}
	}
}

func Call_Main_Transitive_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Simple_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_MultipleDeterminers_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_DeterminedCycle_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Cyclic_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_d(dict_0_loop *Constructor_Main_Transitive) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Transitive = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_c(dict_0_loop *Constructor_Main_Simple) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Simple = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
