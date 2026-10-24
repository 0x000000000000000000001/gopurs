package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_0_box))
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}), Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512((&Constructor_Main_Pair[int64]{1, int64(2), int64(3)})))}).IntVal))
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
		cache_Main_test1 = gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}), Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512((&Constructor_Main_Pair[int64]{1, int64(2), int64(3)})))}).IntVal)
	})
	return cache_Main_test1
}

type Constructor_Main_Pair[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_a
}

type Constructor_Main_Newtype[T_t any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004e006500770074007900700065
func init() {
	gopurs_runtime.StructGetters[1443761658] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Coercible0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Newtype: " + key)
		}
	}
}

func Call_Main_Multiplicative(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_wrap(_dollar___unused_0_loop *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var _dollar___unused_0 *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
}

func Call_Main_unwrap(_dollar___unused_0_loop *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var _dollar___unused_0 *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
}

func Call_Main_semiringMultiplicative(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
	_ = dictSemiring_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Data_Semiring_mul(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_0)), v_1, v1_2)
	})}))}
}

func Call_Main_foldPair(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Main_Pair[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Main_Pair[gopurs_runtime.Value] = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(Call_Data_Semigroup_go__append(dictSemigroup_0), gopurs_runtime.Apply(f_1, (v_2).V0), gopurs_runtime.Apply(f_1, (v_2).V1))
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
	return gopurs_runtime.Apply2(Call_Data_Functor_go__map(dictFunctor_0), Call_Main_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply(f_3, Call_Main_wrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}

func Call_Main_test(dictSemiring_0_loop *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemiring_0 *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] = dictSemiring_0_loop
	_ = dictSemiring_0
	return Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}), Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(dictSemiring_0)})))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1557125915_791404512(in *Constructor_Main_Pair[int64]) *Constructor_Main_Pair[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Pair[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2826095630_348932501(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.IntVal
	out.V3 = in.V3.IntVal
	return out
}

func Rebox_Main_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Int(in.V2)
	out.V3 = gopurs_runtime.Int(in.V3)
	return out
}
