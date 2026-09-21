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
		cache_Main_coerce = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Main_coerce
}

var cache_Main_coerce1 gopurs_runtime.Value
var once_Main_coerce1 sync.Once

func Get_Main_coerce1() gopurs_runtime.Value {
	once_Main_coerce1.Do(func() {
		cache_Main_coerce1 = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
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

var cache_Main_Pair__1394429419 gopurs_runtime.Value
var once_Main_Pair__1394429419 sync.Once

func Get_Main_Pair__1394429419() gopurs_runtime.Value {
	once_Main_Pair__1394429419.Do(func() {
		cache_Main_Pair__1394429419 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512(Call_Main_Pair__1394429419(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal)))}
		})
	})
	return cache_Main_Pair__1394429419
}

var cache_Main_Newtype_dollar_Dict gopurs_runtime.Value
var once_Main_Newtype_dollar_Dict sync.Once

func Get_Main_Newtype_dollar_Dict() gopurs_runtime.Value {
	once_Main_Newtype_dollar_Dict.Do(func() {
		cache_Main_Newtype_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1443761658, UnsafePtr: unsafe.Pointer(Call_Main_Newtype_dollar_Dict(func() struct {
				Coercible0 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Coercible0 gopurs_runtime.Value
				}{}
				clone.Coercible0 = gopurs_runtime.RecordGet(orig, "Coercible0")
				return clone
			}()))}
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
		cache_Main_wrap1 = Call_Main_wrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
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
		cache_Main_unwrap1 = Call_Main_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
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

var cache_Main_ala__754736784 gopurs_runtime.Value
var once_Main_ala__754736784 sync.Once

func Get_Main_ala__754736784() gopurs_runtime.Value {
	once_Main_ala__754736784.Do(func() {
		cache_Main_ala__754736784 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ala__754736784(v_unused_0_box, f_unused_1_box, __eta_norm_0_2_box)
		})
	})
	return cache_Main_ala__754736784
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

var cache_Main_test__1409993861 gopurs_runtime.Value
var once_Main_test__1409993861 sync.Once

func Get_Main_test__1409993861() gopurs_runtime.Value {
	once_Main_test__1409993861.Do(func() {
		cache_Main_test__1409993861 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test__1409993861(Rebox_Main_791404512_1557125915(gopurs_runtime.CoerceToStruct[Constructor_Main_Pair[gopurs_runtime.Value]](__eta_norm_0_0_box))))
		})
	})
	return cache_Main_test__1409993861
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}), Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512((&Constructor_Main_Pair[int64]{1, int64(2), int64(3)})))}).IntVal)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Int(Call_Main_ala__754736784(Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))})), gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512((&Constructor_Main_Pair[int64]{1, int64(2), int64(3)})))}).IntVal)
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

func Call_Main_Pair__1394429419(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) *Constructor_Main_Pair[int64] {
Pair__1394429419:
	for {
		if false {
			continue Pair__1394429419
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Pair[int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_Newtype_dollar_Dict(x_0_loop struct {
	Coercible0 gopurs_runtime.Value
}) *Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Coercible0 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("Coercible0", orig.Coercible0)
	}())
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
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
	})}))}
}

func Call_Main_foldPair(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Main_Pair[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Main_Pair[gopurs_runtime.Value] = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictSemigroup_0.V0, gopurs_runtime.Apply(f_1, (v_2).V0), gopurs_runtime.Apply(f_1, (v_2).V1))
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
	return gopurs_runtime.Apply2(dictFunctor_0.V0, Call_Main_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply(f_3, Call_Main_wrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}

func Call_Main_ala__754736784(v_unused_0_loop gopurs_runtime.Value, f_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
ala__754736784:
	for {
		if false {
			continue ala__754736784
		}
		var v_unused_0 gopurs_runtime.Value = v_unused_0_loop
		_ = v_unused_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), Call_Main_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}), Call_Main_wrap(gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))), __eta_norm_0_2)
	}
}

func Call_Main_test(dictSemiring_0_loop *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemiring_0 *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] = dictSemiring_0_loop
	_ = dictSemiring_0
	return Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}), Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(dictSemiring_0)})))
}

func Call_Main_test__1409993861(__eta_norm_0_0_loop *Constructor_Main_Pair[int64]) int64 {
test__1409993861:
	for {
		if false {
			continue test__1409993861
		}
		var __eta_norm_0_0 *Constructor_Main_Pair[int64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Call_Main_ala(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()), gopurs_runtime.CoerceToStruct[Constructor_Main_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}), Get_Main_Multiplicative(), gopurs_runtime.Apply(Get_Main_foldPair(), Call_Main_semiringMultiplicative(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}))), gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512(__eta_norm_0_0))}).IntVal
	}
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

func Rebox_Main_791404512_1557125915(in *Constructor_Main_Pair[gopurs_runtime.Value]) *Constructor_Main_Pair[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Pair[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.IntVal
	return out
}
