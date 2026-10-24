package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Check gopurs_runtime.Value
var once_Main_Check sync.Once

func Get_Main_Check() gopurs_runtime.Value {
	once_Main_Check.Do(func() {
		cache_Main_Check = gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}
	})
	return cache_Main_Check
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
		cache_Main_withArgEmptyCheck = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_withArgEmptyCheck
}

var cache_Main_withArgHasEmptySuperCheck gopurs_runtime.Value
var once_Main_withArgHasEmptySuperCheck sync.Once

func Get_Main_withArgHasEmptySuperCheck() gopurs_runtime.Value {
	once_Main_withArgHasEmptySuperCheck.Do(func() {
		cache_Main_withArgHasEmptySuperCheck = gopurs_runtime.Value{Type: 9, IntVal: 785089286, UnsafePtr: unsafe.Pointer(Rebox_Main_761192255_1292520722((&Constructor_Main_WithArgHasEmptySuper[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		}), 60647608})))}
	})
	return cache_Main_withArgHasEmptySuperCheck
}

var cache_Main_whenAccessingSuperDict gopurs_runtime.Value
var once_Main_whenAccessingSuperDict sync.Once

func Get_Main_whenAccessingSuperDict() gopurs_runtime.Value {
	once_Main_whenAccessingSuperDict.Do(func() {
		cache_Main_whenAccessingSuperDict = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Func2(func(dictWithArgHasEmptySuper_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.Func2(func(_dollar___unused_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_3
			}), gopurs_runtime.Value{}, x_1)
		}), gopurs_runtime.Value{Type: 9, IntVal: 785089286, UnsafePtr: unsafe.Pointer(Rebox_Main_761192255_1292520722(Rebox_Main_1292520722_761192255(gopurs_runtime.CoerceToStruct[Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value]](Get_Main_withArgHasEmptySuperCheck()))))}, gopurs_runtime.Value{Type: 9, IntVal: int64(60647608), UnsafePtr: nil}).IntVal)), UnsafePtr: nil}
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
		cache_Main_eqCheck = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})})))}
	})
	return cache_Main_eqCheck
}

var cache_Main_emptyDictInst gopurs_runtime.Value
var once_Main_emptyDictInst sync.Once

func Get_Main_emptyDictInst() gopurs_runtime.Value {
	once_Main_emptyDictInst.Do(func() {
		cache_Main_emptyDictInst = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00570069007400680041007200670045006d007000740079
func init() {
	gopurs_runtime.StructGetters[207373949] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_WithArgEmpty[gopurs_runtime.Value])(ptr)
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
	V1 T_t
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00570069007400680041007200670048006100730045006d00700074007900530075007000650072
func init() {
	gopurs_runtime.StructGetters[785089286] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value])(ptr)
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0045006d0070007400790043006c006100730073
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0048006100730045006d00700074007900530075007000650072
func init() {
	gopurs_runtime.StructGetters[2786347248] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_HasEmptySuper)(ptr)
		_ = c
		switch key {
		case "EmptyClass0":
			return gopurs_runtime.Box(c.V0)
		case "hasEmptySuper":
			return gopurs_runtime.Value{Type: 9, IntVal: int64(c.V1)}
		default:
			panic("Key not found in dictionary Constructor_Main_HasEmptySuper: " + key)
		}
	}
}

type Constructor_Main_HasNonEmptySuper struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004800610073004e006f006e0045006d00700074007900530075007000650072
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0041006c0069006100730045006d0070007400790043006c006100730073
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

func Call_Main_withArgHasEmptySuper(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "withArgHasEmptySuper")
}

func Call_Main_hasNonEmptySuperInst(dictHasEmptySuper_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHasEmptySuper_0 gopurs_runtime.Value = dictHasEmptySuper_0_loop
	_ = dictHasEmptySuper_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3506164383, UnsafePtr: unsafe.Pointer((&Constructor_Main_HasNonEmptySuper{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return dictHasEmptySuper_0
	})}))}
}

func Call_Main_hasEmptySuper(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(dict_0, "hasEmptySuper").IntVal)), UnsafePtr: nil}
}

func Rebox_Main_1292520722_761192255(in *Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value]) *Constructor_Main_WithArgHasEmptySuper[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_WithArgHasEmptySuper[uint32]{}
	out.V0 = in.V0
	out.V1 = uint32(in.V1.IntVal)
	return out
}

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_761192255_1292520722(in *Constructor_Main_WithArgHasEmptySuper[uint32]) *Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_WithArgHasEmptySuper[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
}
