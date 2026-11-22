package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_composeFlipped gopurs_runtime.Value
var once_Data_Profunctor_composeFlipped sync.Once

func Get_Data_Profunctor_composeFlipped() gopurs_runtime.Value {
	once_Data_Profunctor_composeFlipped.Do(func() {
		cache_Data_Profunctor_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Profunctor_composeFlipped
}

var cache_Data_Profunctor_composeFlipped1 gopurs_runtime.Value
var once_Data_Profunctor_composeFlipped1 sync.Once

func Get_Data_Profunctor_composeFlipped1() gopurs_runtime.Value {
	once_Data_Profunctor_composeFlipped1.Do(func() {
		cache_Data_Profunctor_composeFlipped1 = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Profunctor_composeFlipped1
}

var cache_Data_Profunctor_identity gopurs_runtime.Value
var once_Data_Profunctor_identity sync.Once

func Get_Data_Profunctor_identity() gopurs_runtime.Value {
	once_Data_Profunctor_identity.Do(func() {
		cache_Data_Profunctor_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_identity(x_0_box)
		})
	})
	return cache_Data_Profunctor_identity
}

var cache_Data_Profunctor_identity1 gopurs_runtime.Value
var once_Data_Profunctor_identity1 sync.Once

func Get_Data_Profunctor_identity1() gopurs_runtime.Value {
	once_Data_Profunctor_identity1.Do(func() {
		cache_Data_Profunctor_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_identity1(x_0_box)
		})
	})
	return cache_Data_Profunctor_identity1
}

var cache_Data_Profunctor_wrap gopurs_runtime.Value
var once_Data_Profunctor_wrap sync.Once

func Get_Data_Profunctor_wrap() gopurs_runtime.Value {
	once_Data_Profunctor_wrap.Do(func() {
		cache_Data_Profunctor_wrap = Call_Data_Newtype_wrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Profunctor_wrap
}

var cache_Data_Profunctor_unwrap gopurs_runtime.Value
var once_Data_Profunctor_unwrap sync.Once

func Get_Data_Profunctor_unwrap() gopurs_runtime.Value {
	once_Data_Profunctor_unwrap.Do(func() {
		cache_Data_Profunctor_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Profunctor_unwrap
}

var cache_Data_Profunctor_unwrap1 gopurs_runtime.Value
var once_Data_Profunctor_unwrap1 sync.Once

func Get_Data_Profunctor_unwrap1() gopurs_runtime.Value {
	once_Data_Profunctor_unwrap1.Do(func() {
		cache_Data_Profunctor_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Profunctor_unwrap1
}

var cache_Data_Profunctor_wrap1 gopurs_runtime.Value
var once_Data_Profunctor_wrap1 sync.Once

func Get_Data_Profunctor_wrap1() gopurs_runtime.Value {
	once_Data_Profunctor_wrap1.Do(func() {
		cache_Data_Profunctor_wrap1 = Call_Data_Newtype_wrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Profunctor_wrap1
}

var cache_Data_Profunctor_Profunctor_dollar_Dict gopurs_runtime.Value
var once_Data_Profunctor_Profunctor_dollar_Dict sync.Once

func Get_Data_Profunctor_Profunctor_dollar_Dict() gopurs_runtime.Value {
	once_Data_Profunctor_Profunctor_dollar_Dict.Do(func() {
		cache_Data_Profunctor_Profunctor_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(Call_Data_Profunctor_Profunctor_dollar_Dict(func() struct {
				dimap gopurs_runtime.Value
				/*gopurs:record-v1:1:00640069006d00610070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					dimap gopurs_runtime.Value
					/*gopurs:record-v1:1:00640069006d00610070*/
				}{}
				clone.dimap = gopurs_runtime.RecordGet(orig, "dimap")
				return clone
			}()))}
		})
	})
	return cache_Data_Profunctor_Profunctor_dollar_Dict
}

var cache_Data_Profunctor_profunctorFn gopurs_runtime.Value
var once_Data_Profunctor_profunctorFn sync.Once

func Get_Data_Profunctor_profunctorFn() gopurs_runtime.Value {
	once_Data_Profunctor_profunctorFn.Do(func() {
		cache_Data_Profunctor_profunctorFn = gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(a2b_0 gopurs_runtime.Value, c2d_1 gopurs_runtime.Value, b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), a2b_0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), b2c_2, c2d_1))
		})}))}
	})
	return cache_Data_Profunctor_profunctorFn
}

var cache_Data_Profunctor_dimap gopurs_runtime.Value
var once_Data_Profunctor_dimap sync.Once

func Get_Data_Profunctor_dimap() gopurs_runtime.Value {
	once_Data_Profunctor_dimap.Do(func() {
		cache_Data_Profunctor_dimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_dimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_dimap
}

var cache_Data_Profunctor_lcmap gopurs_runtime.Value
var once_Data_Profunctor_lcmap sync.Once

func Get_Data_Profunctor_lcmap() gopurs_runtime.Value {
	once_Data_Profunctor_lcmap.Do(func() {
		cache_Data_Profunctor_lcmap = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_lcmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), a2b_1_box)
		})
	})
	return cache_Data_Profunctor_lcmap
}

var cache_Data_Profunctor_rmap gopurs_runtime.Value
var once_Data_Profunctor_rmap sync.Once

func Get_Data_Profunctor_rmap() gopurs_runtime.Value {
	once_Data_Profunctor_rmap.Do(func() {
		cache_Data_Profunctor_rmap = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_rmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), b2c_1_box)
		})
	})
	return cache_Data_Profunctor_rmap
}

var cache_Data_Profunctor_unwrapIso gopurs_runtime.Value
var once_Data_Profunctor_unwrapIso sync.Once

func Get_Data_Profunctor_unwrapIso() gopurs_runtime.Value {
	once_Data_Profunctor_unwrapIso.Do(func() {
		cache_Data_Profunctor_unwrapIso = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_unwrapIso(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar___unused_1_box))
		})
	})
	return cache_Data_Profunctor_unwrapIso
}

var cache_Data_Profunctor_wrapIso gopurs_runtime.Value
var once_Data_Profunctor_wrapIso sync.Once

func Get_Data_Profunctor_wrapIso() gopurs_runtime.Value {
	once_Data_Profunctor_wrapIso.Do(func() {
		cache_Data_Profunctor_wrapIso = gopurs_runtime.Func3(func(dictProfunctor_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_wrapIso(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar___unused_1_box), v_2_box)
		})
	})
	return cache_Data_Profunctor_wrapIso
}

var cache_Data_Profunctor_arr gopurs_runtime.Value
var once_Data_Profunctor_arr sync.Once

func Get_Data_Profunctor_arr() gopurs_runtime.Value {
	once_Data_Profunctor_arr.Do(func() {
		cache_Data_Profunctor_arr = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_arr(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](dictCategory_0_box))
		})
	})
	return cache_Data_Profunctor_arr
}

type Constructor_Data_Profunctor_Profunctor[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00500072006f00660075006e00630074006f0072:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00500072006f00660075006e00630074006f0072005f00500072006f00660075006e00630074006f0072
func init() {
	gopurs_runtime.StructGetters[2367018778] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "dimap":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Data_Profunctor_Profunctor: " + key)
		}
	}
}

func Call_Data_Profunctor_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Profunctor_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Profunctor_Profunctor_dollar_Dict(x_0_loop struct {
	dimap gopurs_runtime.Value
	/*gopurs:record-v1:1:00640069006d00610070*/
}) *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] {
	var x_0 struct {
		dimap gopurs_runtime.Value
		/*gopurs:record-v1:1:00640069006d00610070*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("dimap", orig.dimap)
	}())
}

func Call_Data_Profunctor_dimap(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Profunctor_lcmap(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
	_ = dictProfunctor_0
	var a2b_1 gopurs_runtime.Value = a2b_1_loop
	_ = a2b_1
	return gopurs_runtime.Apply2(dictProfunctor_0.V0, a2b_1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}))
}

func Call_Data_Profunctor_rmap(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
	_ = dictProfunctor_0
	var b2c_1 gopurs_runtime.Value = b2c_1_loop
	_ = b2c_1
	return gopurs_runtime.Apply2(dictProfunctor_0.V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}), b2c_1)
}

func Call_Data_Profunctor_unwrapIso(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], _dollar___unused_1_loop *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
	_ = dictProfunctor_0
	var _dollar___unused_1 *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_1_loop
	_ = _dollar___unused_1
	return gopurs_runtime.Apply2(dictProfunctor_0.V0, Call_Data_Newtype_wrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})))
}

func Call_Data_Profunctor_wrapIso(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], _dollar___unused_1_loop *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
	_ = dictProfunctor_0
	var _dollar___unused_1 *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictProfunctor_0.V0, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Call_Data_Newtype_wrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})))
}

func Call_Data_Profunctor_arr(dictCategory_0_loop *Constructor_Control_Category_Category[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictCategory_0 *Constructor_Control_Category_Category[gopurs_runtime.Value] = dictCategory_0_loop
	_ = dictCategory_0
	// TAST (Let): identity2_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar p$scope51) [(TypeVar a$scope49), (TypeVar a$scope49)])
	identity2_1_0 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(dictCategory_0)})
	_ = identity2_1_0
	return gopurs_runtime.Func2(func(dictProfunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_2, "dimap"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_4
		}), f_3, identity2_1_0)
	})
}
