package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_WithArgEmpty_dollar_Dict gopurs_runtime.Value
var once_Main_WithArgEmpty_dollar_Dict sync.Once

func Get_Main_WithArgEmpty_dollar_Dict() gopurs_runtime.Value {
	once_Main_WithArgEmpty_dollar_Dict.Do(func() {
		cache_Main_WithArgEmpty_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_WithArgEmpty_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_WithArgEmpty_dollar_Dict
}

var cache_Main_WithArgHasEmptySuper_dollar_Dict gopurs_runtime.Value
var once_Main_WithArgHasEmptySuper_dollar_Dict sync.Once

func Get_Main_WithArgHasEmptySuper_dollar_Dict() gopurs_runtime.Value {
	once_Main_WithArgHasEmptySuper_dollar_Dict.Do(func() {
		cache_Main_WithArgHasEmptySuper_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_WithArgHasEmptySuper_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_WithArgHasEmptySuper_dollar_Dict
}

var cache_Main_EmptyClass_dollar_Dict gopurs_runtime.Value
var once_Main_EmptyClass_dollar_Dict sync.Once

func Get_Main_EmptyClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_EmptyClass_dollar_Dict.Do(func() {
		cache_Main_EmptyClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_EmptyClass_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_EmptyClass_dollar_Dict
}

var cache_Main_Check gopurs_runtime.Value
var once_Main_Check sync.Once

func Get_Main_Check() gopurs_runtime.Value {
	once_Main_Check.Do(func() {
		cache_Main_Check = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_Check
}

var cache_Main_HasEmptySuper_dollar_Dict gopurs_runtime.Value
var once_Main_HasEmptySuper_dollar_Dict sync.Once

func Get_Main_HasEmptySuper_dollar_Dict() gopurs_runtime.Value {
	once_Main_HasEmptySuper_dollar_Dict.Do(func() {
		cache_Main_HasEmptySuper_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_HasEmptySuper_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_HasEmptySuper_dollar_Dict
}

var cache_Main_HasNonEmptySuper_dollar_Dict gopurs_runtime.Value
var once_Main_HasNonEmptySuper_dollar_Dict sync.Once

func Get_Main_HasNonEmptySuper_dollar_Dict() gopurs_runtime.Value {
	once_Main_HasNonEmptySuper_dollar_Dict.Do(func() {
		cache_Main_HasNonEmptySuper_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_HasNonEmptySuper_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_HasNonEmptySuper_dollar_Dict
}

var cache_Main_AliasEmptyClass_dollar_Dict gopurs_runtime.Value
var once_Main_AliasEmptyClass_dollar_Dict sync.Once

func Get_Main_AliasEmptyClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_AliasEmptyClass_dollar_Dict.Do(func() {
		cache_Main_AliasEmptyClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_AliasEmptyClass_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_AliasEmptyClass_dollar_Dict
}

var cache_Main_withArgHasEmptySuper gopurs_runtime.Value
var once_Main_withArgHasEmptySuper sync.Once

func Get_Main_withArgHasEmptySuper() gopurs_runtime.Value {
	once_Main_withArgHasEmptySuper.Do(func() {
		cache_Main_withArgHasEmptySuper = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_withArgHasEmptySuper(dict_0_box)
		})
	})
	return cache_Main_withArgHasEmptySuper
}

var cache_Main_withArgEmptyCheck gopurs_runtime.Value
var once_Main_withArgEmptyCheck sync.Once

func Get_Main_withArgEmptyCheck() gopurs_runtime.Value {
	once_Main_withArgEmptyCheck.Do(func() {
		cache_Main_withArgEmptyCheck = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_withArgEmptyCheck
}

var cache_Main_withArgHasEmptySuperCheck gopurs_runtime.Value
var once_Main_withArgHasEmptySuperCheck sync.Once

func Get_Main_withArgHasEmptySuperCheck() gopurs_runtime.Value {
	once_Main_withArgHasEmptySuperCheck.Do(func() {
		cache_Main_withArgHasEmptySuperCheck = gopurs_runtime.Value{Type: 9, IntVal: 785089286, UnsafePtr: unsafe.Pointer((&Constructor_Main_WithArgHasEmptySuper[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		}), gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}}))}
	})
	return cache_Main_withArgHasEmptySuperCheck
}

var cache_Main_whenAccessingSuperDict gopurs_runtime.Value
var once_Main_whenAccessingSuperDict sync.Once

func Get_Main_whenAccessingSuperDict() gopurs_runtime.Value {
	once_Main_whenAccessingSuperDict.Do(func() {
		cache_Main_whenAccessingSuperDict = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
	})
	return cache_Main_whenAccessingSuperDict
}

var cache_Main_hasNonEmptySuperInst gopurs_runtime.Value
var once_Main_hasNonEmptySuperInst sync.Once

func Get_Main_hasNonEmptySuperInst() gopurs_runtime.Value {
	once_Main_hasNonEmptySuperInst.Do(func() {
		cache_Main_hasNonEmptySuperInst = gopurs_runtime.Func(func(dictHasEmptySuper_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hasNonEmptySuperInst(dictHasEmptySuper_0_box)
		})
	})
	return cache_Main_hasNonEmptySuperInst
}

var cache_Main_hasEmptySuper gopurs_runtime.Value
var once_Main_hasEmptySuper sync.Once

func Get_Main_hasEmptySuper() gopurs_runtime.Value {
	once_Main_hasEmptySuper.Do(func() {
		cache_Main_hasEmptySuper = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_hasEmptySuper(dict_0_box)
		})
	})
	return cache_Main_hasEmptySuper
}

var cache_Main_eqCheck gopurs_runtime.Value
var once_Main_eqCheck sync.Once

func Get_Main_eqCheck() gopurs_runtime.Value {
	once_Main_eqCheck.Do(func() {
		cache_Main_eqCheck = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})}))}
	})
	return cache_Main_eqCheck
}

var cache_Main_emptyDictInst gopurs_runtime.Value
var once_Main_emptyDictInst sync.Once

func Get_Main_emptyDictInst() gopurs_runtime.Value {
	once_Main_emptyDictInst.Do(func() {
		cache_Main_emptyDictInst = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_emptyDictInst
}

var cache_Main_hasEmptySuperInst gopurs_runtime.Value
var once_Main_hasEmptySuperInst sync.Once

func Get_Main_hasEmptySuperInst() gopurs_runtime.Value {
	once_Main_hasEmptySuperInst.Do(func() {
		cache_Main_hasEmptySuperInst = gopurs_runtime.Value{Type: 9, IntVal: 2786347248, UnsafePtr: unsafe.Pointer((&Constructor_Main_HasEmptySuper{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		}), 60647608}))}
	})
	return cache_Main_hasEmptySuperInst
}

var cache_Main_whenHasEmptySuper gopurs_runtime.Value
var once_Main_whenHasEmptySuper sync.Once

func Get_Main_whenHasEmptySuper() gopurs_runtime.Value {
	once_Main_whenHasEmptySuper.Do(func() {
		cache_Main_whenHasEmptySuper = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenHasEmptySuper
}

var cache_Main_whenHasNonEmptySuper gopurs_runtime.Value
var once_Main_whenHasNonEmptySuper sync.Once

func Get_Main_whenHasNonEmptySuper() gopurs_runtime.Value {
	once_Main_whenHasNonEmptySuper.Do(func() {
		cache_Main_whenHasNonEmptySuper = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenHasNonEmptySuper
}

var cache_Main_whenEmpty gopurs_runtime.Value
var once_Main_whenEmpty sync.Once

func Get_Main_whenEmpty() gopurs_runtime.Value {
	once_Main_whenEmpty.Do(func() {
		cache_Main_whenEmpty = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenEmpty
}

var cache_Main_aliasEmptyClassInst gopurs_runtime.Value
var once_Main_aliasEmptyClassInst sync.Once

func Get_Main_aliasEmptyClassInst() gopurs_runtime.Value {
	once_Main_aliasEmptyClassInst.Do(func() {
		cache_Main_aliasEmptyClassInst = gopurs_runtime.Value{Type: 9, IntVal: 4080323731, UnsafePtr: unsafe.Pointer((&Constructor_Main_AliasEmptyClass{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_aliasEmptyClassInst
}

var cache_Main_whenAliasEmptyClass gopurs_runtime.Value
var once_Main_whenAliasEmptyClass sync.Once

func Get_Main_whenAliasEmptyClass() gopurs_runtime.Value {
	once_Main_whenAliasEmptyClass.Do(func() {
		cache_Main_whenAliasEmptyClass = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_whenAliasEmptyClass
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Check struct {
	Rc uint32
}

type Constructor_Main_WithArgEmpty[T_t any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[207373949] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_WithArgEmpty[any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_WithArgEmpty: " + key)
		}
	}
}

type Constructor_Main_WithArgHasEmptySuper[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[785089286] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_WithArgHasEmptySuper[any])(ptr)
		_ = c
		switch key {
		case "WithArgEmpty0":
			return gopurs_runtime.Box(c.V0)
		case "withArgHasEmptySuper":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_WithArgHasEmptySuper: " + key)
		}
	}
}

type Constructor_Main_EmptyClass struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[394228773] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_EmptyClass)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_EmptyClass: " + key)
		}
	}
}

type Constructor_Main_HasEmptySuper struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 uint32
}

func init() {
	gopurs_runtime.StructGetters[2786347248] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HasEmptySuper)(ptr)
		_ = c
		switch key {
		case "EmptyClass0":
			return gopurs_runtime.Box(c.V0)
		case "hasEmptySuper":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_HasEmptySuper: " + key)
		}
	}
}

type Constructor_Main_HasNonEmptySuper struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3506164383] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HasNonEmptySuper)(ptr)
		_ = c
		switch key {
		case "HasEmptySuper0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_HasNonEmptySuper: " + key)
		}
	}
}

type Constructor_Main_AliasEmptyClass struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4080323731] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_AliasEmptyClass)(ptr)
		_ = c
		switch key {
		case "EmptyClass0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_AliasEmptyClass: " + key)
		}
	}
}

func Call_Main_WithArgEmpty_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_WithArgHasEmptySuper_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_EmptyClass_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_HasEmptySuper_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_HasNonEmptySuper_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_AliasEmptyClass_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_withArgHasEmptySuper(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "withArgHasEmptySuper")
}

func Call_Main_hasNonEmptySuperInst(dictHasEmptySuper_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHasEmptySuper_0 gopurs_runtime.Value = dictHasEmptySuper_0_loop
	_ = dictHasEmptySuper_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3506164383, UnsafePtr: unsafe.Pointer((&Constructor_Main_HasNonEmptySuper{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2786347248, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_HasEmptySuper](dictHasEmptySuper_0))}
	})}))}
}

func Call_Main_hasEmptySuper(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "hasEmptySuper")
}
