package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Category_Category_dollar_Dict gopurs_runtime.Value
var once_Control_Category_Category_dollar_Dict sync.Once

func Get_Control_Category_Category_dollar_Dict() gopurs_runtime.Value {
	once_Control_Category_Category_dollar_Dict.Do(func() {
		cache_Control_Category_Category_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(Call_Control_Category_Category_dollar_Dict(func() struct {
				Semigroupoid0 gopurs_runtime.Value
				identity      gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900670072006f00750070006f006900640030,006900640065006e0074006900740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semigroupoid0 gopurs_runtime.Value
					identity      gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900670072006f00750070006f006900640030,006900640065006e0074006900740079*/
				}{}
				clone.Semigroupoid0 = gopurs_runtime.RecordGet(orig, "Semigroupoid0")
				clone.identity = gopurs_runtime.RecordGet(orig, "identity")
				return clone
			}()))}
		})
	})
	return cache_Control_Category_Category_dollar_Dict
}

var cache_Control_Category_identity gopurs_runtime.Value
var once_Control_Category_identity sync.Once

func Get_Control_Category_identity() gopurs_runtime.Value {
	once_Control_Category_identity.Do(func() {
		cache_Control_Category_identity = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Category_identity(dict_0_box)
		})
	})
	return cache_Control_Category_identity
}

var cache_Control_Category_categoryFn gopurs_runtime.Value
var once_Control_Category_categoryFn sync.Once

func Get_Control_Category_categoryFn() gopurs_runtime.Value {
	once_Control_Category_categoryFn.Do(func() {
		cache_Control_Category_categoryFn = gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer((&Constructor_Control_Category_Category[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Control_Category_categoryFn
}

var cache_Control_Category_identity__1909583636 gopurs_runtime.Value
var once_Control_Category_identity__1909583636 sync.Once

func Get_Control_Category_identity__1909583636() gopurs_runtime.Value {
	once_Control_Category_identity__1909583636.Do(func() {
		cache_Control_Category_identity__1909583636 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Category_identity__1909583636(x_0_box)
		})
	})
	return cache_Control_Category_identity__1909583636
}

type Constructor_Control_Category_Category[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f00430061007400650067006f00720079:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f00430061007400650067006f00720079005f00430061007400650067006f00720079
func init() {
	gopurs_runtime.StructGetters[784524589] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Category_Category[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Semigroupoid0":
			return gopurs_runtime.Box(c.V0)
		case "identity":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Category_Category: " + key)
		}
	}
}

func Call_Control_Category_Category_dollar_Dict(x_0_loop struct {
	Semigroupoid0 gopurs_runtime.Value
	identity      gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900670072006f00750070006f006900640030,006900640065006e0074006900740079*/
}) *Constructor_Control_Category_Category[gopurs_runtime.Value] {
	var x_0 struct {
		Semigroupoid0 gopurs_runtime.Value
		identity      gopurs_runtime.Value
		/*gopurs:record-v1:2:00530065006d006900670072006f00750070006f006900640030,006900640065006e0074006900740079*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Semigroupoid0", "identity", orig.Semigroupoid0, orig.identity)
	}())
}

func Call_Control_Category_identity(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "identity")
}

func Call_Control_Category_identity__1909583636(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
identity__1909583636:
	for {
		if false {
			continue identity__1909583636
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}
