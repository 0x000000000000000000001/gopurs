package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Type_Equality_TypeEquals_dollar_Dict gopurs_runtime.Value
var once_Type_Equality_TypeEquals_dollar_Dict sync.Once

func Get_Type_Equality_TypeEquals_dollar_Dict() gopurs_runtime.Value {
	once_Type_Equality_TypeEquals_dollar_Dict.Do(func() {
		cache_Type_Equality_TypeEquals_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3275391293, UnsafePtr: unsafe.Pointer(Call_Type_Equality_TypeEquals_dollar_Dict(func() struct {
				Coercible0 gopurs_runtime.Value
				proof      gopurs_runtime.Value
				/*gopurs:record-v1:2:0043006f00650072006300690062006c00650030,00700072006f006f0066*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Coercible0 gopurs_runtime.Value
					proof      gopurs_runtime.Value
					/*gopurs:record-v1:2:0043006f00650072006300690062006c00650030,00700072006f006f0066*/
				}{}
				clone.Coercible0 = gopurs_runtime.RecordGet(orig, "Coercible0")
				clone.proof = gopurs_runtime.RecordGet(orig, "proof")
				return clone
			}()))}
		})
	})
	return cache_Type_Equality_TypeEquals_dollar_Dict
}

var cache_Type_Equality_To gopurs_runtime.Value
var once_Type_Equality_To sync.Once

func Get_Type_Equality_To() gopurs_runtime.Value {
	once_Type_Equality_To.Do(func() {
		cache_Type_Equality_To = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Equality_To(x_0_box)
		})
	})
	return cache_Type_Equality_To
}

var cache_Type_Equality_From gopurs_runtime.Value
var once_Type_Equality_From sync.Once

func Get_Type_Equality_From() gopurs_runtime.Value {
	once_Type_Equality_From.Do(func() {
		cache_Type_Equality_From = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Equality_From(x_0_box)
		})
	})
	return cache_Type_Equality_From
}

var cache_Type_Equality_refl gopurs_runtime.Value
var once_Type_Equality_refl sync.Once

func Get_Type_Equality_refl() gopurs_runtime.Value {
	once_Type_Equality_refl.Do(func() {
		cache_Type_Equality_refl = gopurs_runtime.Value{Type: 9, IntVal: 3275391293, UnsafePtr: unsafe.Pointer((&Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return a_0
		})}))}
	})
	return cache_Type_Equality_refl
}

var cache_Type_Equality_proof gopurs_runtime.Value
var once_Type_Equality_proof sync.Once

func Get_Type_Equality_proof() gopurs_runtime.Value {
	once_Type_Equality_proof.Do(func() {
		cache_Type_Equality_proof = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Equality_proof(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Type_Equality_proof
}

var cache_Type_Equality_to gopurs_runtime.Value
var once_Type_Equality_to sync.Once

func Get_Type_Equality_to() gopurs_runtime.Value {
	once_Type_Equality_to.Do(func() {
		cache_Type_Equality_to = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Equality_to(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
		})
	})
	return cache_Type_Equality_to
}

var cache_Type_Equality_from gopurs_runtime.Value
var once_Type_Equality_from sync.Once

func Get_Type_Equality_from() gopurs_runtime.Value {
	once_Type_Equality_from.Do(func() {
		cache_Type_Equality_from = gopurs_runtime.Func(func(dictTypeEquals_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Type_Equality_from(gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](dictTypeEquals_0_box))
		})
	})
	return cache_Type_Equality_from
}

type Constructor_Type_Equality_TypeEquals[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0054007900700065005f0045007100750061006c006900740079:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0054007900700065005f0045007100750061006c006900740079005f00540079007000650045007100750061006c0073
func init() {
	gopurs_runtime.StructGetters[3275391293] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Coercible0":
			return gopurs_runtime.Box(c.V0)
		case "proof":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Type_Equality_TypeEquals: " + key)
		}
	}
}

func Call_Type_Equality_TypeEquals_dollar_Dict(x_0_loop struct {
	Coercible0 gopurs_runtime.Value
	proof      gopurs_runtime.Value
	/*gopurs:record-v1:2:0043006f00650072006300690062006c00650030,00700072006f006f0066*/
}) *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Coercible0 gopurs_runtime.Value
		proof      gopurs_runtime.Value
		/*gopurs:record-v1:2:0043006f00650072006300690062006c00650030,00700072006f006f0066*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Coercible0", "proof", orig.Coercible0, orig.proof)
	}())
}

func Call_Type_Equality_To(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Type_Equality_From(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Type_Equality_proof(dict_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Type_Equality_to(dictTypeEquals_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTypeEquals_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
	_ = dictTypeEquals_0
	return gopurs_runtime.Apply(Call_Type_Equality_proof(dictTypeEquals_0), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return a_1
	}))
}

func Call_Type_Equality_from(dictTypeEquals_0_loop *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTypeEquals_0 *Constructor_Type_Equality_TypeEquals[gopurs_runtime.Value, gopurs_runtime.Value] = dictTypeEquals_0_loop
	_ = dictTypeEquals_0
	return gopurs_runtime.Apply(Call_Type_Equality_proof(dictTypeEquals_0), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return a_1
	}))
}
