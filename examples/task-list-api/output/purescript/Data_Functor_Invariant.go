package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Invariant_compose gopurs_runtime.Value
var once_Data_Functor_Invariant_compose sync.Once

func Get_Data_Functor_Invariant_compose() gopurs_runtime.Value {
	once_Data_Functor_Invariant_compose.Do(func() {
		cache_Data_Functor_Invariant_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Invariant_compose
}

var cache_Data_Functor_Invariant_compose1 gopurs_runtime.Value
var once_Data_Functor_Invariant_compose1 sync.Once

func Get_Data_Functor_Invariant_compose1() gopurs_runtime.Value {
	once_Data_Functor_Invariant_compose1.Do(func() {
		cache_Data_Functor_Invariant_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Invariant_compose1
}

var cache_Data_Functor_Invariant_Invariant_dollar_Dict gopurs_runtime.Value
var once_Data_Functor_Invariant_Invariant_dollar_Dict sync.Once

func Get_Data_Functor_Invariant_Invariant_dollar_Dict() gopurs_runtime.Value {
	once_Data_Functor_Invariant_Invariant_dollar_Dict.Do(func() {
		cache_Data_Functor_Invariant_Invariant_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Invariant_Invariant_dollar_Dict(func() struct {
				imap gopurs_runtime.Value
				/*gopurs:record-v1:1:0069006d00610070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					imap gopurs_runtime.Value
					/*gopurs:record-v1:1:0069006d00610070*/
				}{}
				clone.imap = gopurs_runtime.RecordGet(orig, "imap")
				return clone
			}()))}
		})
	})
	return cache_Data_Functor_Invariant_Invariant_dollar_Dict
}

var cache_Data_Functor_Invariant_invariantMultiplicative gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantMultiplicative sync.Once

func Get_Data_Functor_Invariant_invariantMultiplicative() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantMultiplicative.Do(func() {
		cache_Data_Functor_Invariant_invariantMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, v1_2)
		})}))}
	})
	return cache_Data_Functor_Invariant_invariantMultiplicative
}

var cache_Data_Functor_Invariant_invariantEndo gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantEndo sync.Once

func Get_Data_Functor_Invariant_invariantEndo() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantEndo.Do(func() {
		cache_Data_Functor_Invariant_invariantEndo = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(ab_0 gopurs_runtime.Value, ba_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), ab_0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_2, ba_1))
		})}))}
	})
	return cache_Data_Functor_Invariant_invariantEndo
}

var cache_Data_Functor_Invariant_invariantDual gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantDual sync.Once

func Get_Data_Functor_Invariant_invariantDual() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantDual.Do(func() {
		cache_Data_Functor_Invariant_invariantDual = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, v1_2)
		})}))}
	})
	return cache_Data_Functor_Invariant_invariantDual
}

var cache_Data_Functor_Invariant_invariantDisj gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantDisj sync.Once

func Get_Data_Functor_Invariant_invariantDisj() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantDisj.Do(func() {
		cache_Data_Functor_Invariant_invariantDisj = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, v1_2)
		})}))}
	})
	return cache_Data_Functor_Invariant_invariantDisj
}

var cache_Data_Functor_Invariant_invariantConj gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantConj sync.Once

func Get_Data_Functor_Invariant_invariantConj() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantConj.Do(func() {
		cache_Data_Functor_Invariant_invariantConj = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, v1_2)
		})}))}
	})
	return cache_Data_Functor_Invariant_invariantConj
}

var cache_Data_Functor_Invariant_invariantAdditive gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantAdditive sync.Once

func Get_Data_Functor_Invariant_invariantAdditive() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantAdditive.Do(func() {
		cache_Data_Functor_Invariant_invariantAdditive = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, v1_2)
		})}))}
	})
	return cache_Data_Functor_Invariant_invariantAdditive
}

var cache_Data_Functor_Invariant_imapF gopurs_runtime.Value
var once_Data_Functor_Invariant_imapF sync.Once

func Get_Data_Functor_Invariant_imapF() gopurs_runtime.Value {
	once_Data_Functor_Invariant_imapF.Do(func() {
		cache_Data_Functor_Invariant_imapF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Invariant_imapF(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
		})
	})
	return cache_Data_Functor_Invariant_imapF
}

var cache_Data_Functor_Invariant_invariantArray gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantArray sync.Once

func Get_Data_Functor_Invariant_invariantArray() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantArray.Do(func() {
		cache_Data_Functor_Invariant_invariantArray = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))})}))}
	})
	return cache_Data_Functor_Invariant_invariantArray
}

var cache_Data_Functor_Invariant_invariantFn gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantFn sync.Once

func Get_Data_Functor_Invariant_invariantFn() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantFn.Do(func() {
		cache_Data_Functor_Invariant_invariantFn = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), Get_Data_Functor_functorFn())}))}
	})
	return cache_Data_Functor_Invariant_invariantFn
}

var cache_Data_Functor_Invariant_imap gopurs_runtime.Value
var once_Data_Functor_Invariant_imap sync.Once

func Get_Data_Functor_Invariant_imap() gopurs_runtime.Value {
	once_Data_Functor_Invariant_imap.Do(func() {
		cache_Data_Functor_Invariant_imap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Invariant_imap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Functor_Invariant_imap
}

var cache_Data_Functor_Invariant_invariantAlternate gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantAlternate sync.Once

func Get_Data_Functor_Invariant_invariantAlternate() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantAlternate.Do(func() {
		cache_Data_Functor_Invariant_invariantAlternate = gopurs_runtime.Func(func(dictInvariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Invariant_invariantAlternate(dictInvariant_0_box)
		})
	})
	return cache_Data_Functor_Invariant_invariantAlternate
}

type Constructor_Data_Functor_Invariant_Invariant[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00460075006e00630074006f0072005f0049006e00760061007200690061006e0074:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00460075006e00630074006f0072005f0049006e00760061007200690061006e0074005f0049006e00760061007200690061006e0074
func init() {
	gopurs_runtime.StructGetters[2396985522] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "imap":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Data_Functor_Invariant_Invariant: " + key)
		}
	}
}

func Call_Data_Functor_Invariant_Invariant_dollar_Dict(x_0_loop struct {
	imap gopurs_runtime.Value
	/*gopurs:record-v1:1:0069006d00610070*/
}) *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] {
	var x_0 struct {
		imap gopurs_runtime.Value
		/*gopurs:record-v1:1:0069006d00610070*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("imap", orig.imap)
	}())
}

func Call_Data_Functor_Invariant_imapF(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply(dictFunctor_0.V0, f_1)
}

func Call_Data_Functor_Invariant_imap(dict_0_loop *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Functor_Invariant_invariantAlternate(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
	_ = dictInvariant_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), f_1, g_2, v_3)
	})}))}
}
