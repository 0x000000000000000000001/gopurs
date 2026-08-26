package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_coerce gopurs_runtime.Value
var once_Main_coerce sync.Once

func Get_Main_coerce() gopurs_runtime.Value {
	once_Main_coerce.Do(func() {
		cache_Main_coerce = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_coerce
}

var cache_Main_coerce1 gopurs_runtime.Value
var once_Main_coerce1 sync.Once

func Get_Main_coerce1() gopurs_runtime.Value {
	once_Main_coerce1.Do(func() {
		cache_Main_coerce1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_coerce1
}

var cache_Main_Pair gopurs_runtime.Value
var once_Main_Pair sync.Once

func Get_Main_Pair() gopurs_runtime.Value {
	once_Main_Pair.Do(func() {
		cache_Main_Pair = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pair[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Pair
}

var cache_Main_Newtype_dollar_Dict gopurs_runtime.Value
var once_Main_Newtype_dollar_Dict sync.Once

func Get_Main_Newtype_dollar_Dict() gopurs_runtime.Value {
	once_Main_Newtype_dollar_Dict.Do(func() {
		cache_Main_Newtype_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Newtype_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Newtype_dollar_Dict
}

var cache_Main_Multiplicative gopurs_runtime.Value
var once_Main_Multiplicative sync.Once

func Get_Main_Multiplicative() gopurs_runtime.Value {
	once_Main_Multiplicative.Do(func() {
		cache_Main_Multiplicative = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Multiplicative(x_0_box)
		})
	})
	return cache_Main_Multiplicative
}

var cache_Main_wrap gopurs_runtime.Value
var once_Main_wrap sync.Once

func Get_Main_wrap() gopurs_runtime.Value {
	once_Main_wrap.Do(func() {
		cache_Main_wrap = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar___unused_0_box))
		})
	})
	return cache_Main_wrap
}

var cache_Main_wrap1 gopurs_runtime.Value
var once_Main_wrap1 sync.Once

func Get_Main_wrap1() gopurs_runtime.Value {
	once_Main_wrap1.Do(func() {
		cache_Main_wrap1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_wrap1
}

var cache_Main_unwrap gopurs_runtime.Value
var once_Main_unwrap sync.Once

func Get_Main_unwrap() gopurs_runtime.Value {
	once_Main_unwrap.Do(func() {
		cache_Main_unwrap = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar___unused_0_box))
		})
	})
	return cache_Main_unwrap
}

var cache_Main_unwrap1 gopurs_runtime.Value
var once_Main_unwrap1 sync.Once

func Get_Main_unwrap1() gopurs_runtime.Value {
	once_Main_unwrap1.Do(func() {
		cache_Main_unwrap1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Main_unwrap1
}

var cache_Main_semiringMultiplicative gopurs_runtime.Value
var once_Main_semiringMultiplicative sync.Once

func Get_Main_semiringMultiplicative() gopurs_runtime.Value {
	once_Main_semiringMultiplicative.Do(func() {
		cache_Main_semiringMultiplicative = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_semiringMultiplicative(dictSemiring_0_box)
		})
	})
	return cache_Main_semiringMultiplicative
}

var cache_Main_newtypeMultiplicative gopurs_runtime.Value
var once_Main_newtypeMultiplicative sync.Once

func Get_Main_newtypeMultiplicative() gopurs_runtime.Value {
	once_Main_newtypeMultiplicative.Do(func() {
		cache_Main_newtypeMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 1443761658, UnsafePtr: unsafe.Pointer((&Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_newtypeMultiplicative
}

var cache_Main_foldPair gopurs_runtime.Value
var once_Main_foldPair sync.Once

func Get_Main_foldPair() gopurs_runtime.Value {
	once_Main_foldPair.Do(func() {
		cache_Main_foldPair = gopurs_runtime.Func3(func(dictSemigroup_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldPair(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Main_Pair[gopurs_runtime.Value]](v_2_box))
		})
	})
	return cache_Main_foldPair
}

var cache_Main_ala gopurs_runtime.Value
var once_Main_ala sync.Once

func Get_Main_ala() gopurs_runtime.Value {
	once_Main_ala.Do(func() {
		cache_Main_ala = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar___unused_1_box), v_2_box, f_3_box)
		})
	})
	return cache_Main_ala
}

var cache_Main_ala__162501229 gopurs_runtime.Value
var once_Main_ala__162501229 sync.Once

func Get_Main_ala__162501229() gopurs_runtime.Value {
	once_Main_ala__162501229.Do(func() {
		cache_Main_ala__162501229 = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ala__162501229(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar___unused_1_box), v_2_box, f_3_box)
		})
	})
	return cache_Main_ala__162501229
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(dictSemiring_0_box)
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Int(6).IntVal)).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Int(gopurs_runtime.Int(gopurs_runtime.Int(6).IntVal).IntVal)
	})
	return cache_Main_test1
}

type Constructor_Main_Pair[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

type Constructor_Main_Newtype[T_t any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1443761658] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Newtype[any, any])(ptr)
		_ = c
		switch key {
		case "Coercible0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Newtype: " + key)
		}
	}
}

func Call_Main_Newtype_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Multiplicative(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_wrap(_dollar___unused_0_loop *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var _dollar___unused_0 *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Unsafe_Coerce_unsafeCoerce()
}

func Call_Main_unwrap(_dollar___unused_0_loop *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var _dollar___unused_0 *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Unsafe_Coerce_unsafeCoerce()
}

func Call_Main_semiringMultiplicative(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
	_ = dictSemiring_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
		})
	})}))}
}

func Call_Main_foldPair(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Main_Pair[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Main_Pair[gopurs_runtime.Value] = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(f_1, (v_2).V0), gopurs_runtime.Apply(f_1, (v_2).V1))
}

func Call_Main_ala(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], _dollar___unused_1_loop *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var _dollar___unused_1 *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Unsafe_Coerce_unsafeCoerce(), gopurs_runtime.Apply(f_3, Get_Unsafe_Coerce_unsafeCoerce()))
}

func Call_Main_ala__162501229(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], _dollar___unused_1_loop *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var _dollar___unused_1 *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Unsafe_Coerce_unsafeCoerce(), gopurs_runtime.Apply(f_3, Get_Unsafe_Coerce_unsafeCoerce()))
}

func Call_Main_test(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
	_ = dictSemiring_0
	// TAST (Let): __local_var_1_0 shape=LitRecord bindingType=(Record (Row [append: (Func [(TypeVar a), (TypeVar a)] (TypeVar a))] Any))
	__local_var_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
		})
	}))
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*Constructor_Main_Pair[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Main_Pair[gopurs_runtime.Value])(x_2.UnsafePtr).V1)
	})
}
