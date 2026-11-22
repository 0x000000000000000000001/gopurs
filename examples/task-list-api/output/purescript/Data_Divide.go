package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Divide_conj gopurs_runtime.Value
var once_Data_Divide_conj sync.Once

func Get_Data_Divide_conj() gopurs_runtime.Value {
	once_Data_Divide_conj.Do(func() {
		cache_Data_Divide_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Divide_conj
}

var cache_Data_Divide_go__append gopurs_runtime.Value
var once_Data_Divide_go__append sync.Once

func Get_Data_Divide_go__append() gopurs_runtime.Value {
	once_Data_Divide_go__append.Do(func() {
		cache_Data_Divide_go__append = Get_Data_Semigroup_append__3720349872()
	})
	return cache_Data_Divide_go__append
}

var cache_Data_Divide_identity gopurs_runtime.Value
var once_Data_Divide_identity sync.Once

func Get_Data_Divide_identity() gopurs_runtime.Value {
	once_Data_Divide_identity.Do(func() {
		cache_Data_Divide_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Divide_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Divide_identity
}

var cache_Data_Divide_Divide_dollar_Dict gopurs_runtime.Value
var once_Data_Divide_Divide_dollar_Dict sync.Once

func Get_Data_Divide_Divide_dollar_Dict() gopurs_runtime.Value {
	once_Data_Divide_Divide_dollar_Dict.Do(func() {
		cache_Data_Divide_Divide_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer(Call_Data_Divide_Divide_dollar_Dict(func() struct {
				Contravariant0 gopurs_runtime.Value
				divide         gopurs_runtime.Value
				/*gopurs:record-v1:2:0043006f006e00740072006100760061007200690061006e00740030,006400690076006900640065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Contravariant0 gopurs_runtime.Value
					divide         gopurs_runtime.Value
					/*gopurs:record-v1:2:0043006f006e00740072006100760061007200690061006e00740030,006400690076006900640065*/
				}{}
				clone.Contravariant0 = gopurs_runtime.RecordGet(orig, "Contravariant0")
				clone.divide = gopurs_runtime.RecordGet(orig, "divide")
				return clone
			}()))}
		})
	})
	return cache_Data_Divide_Divide_dollar_Dict
}

var cache_Data_Divide_dividePredicate gopurs_runtime.Value
var once_Data_Divide_dividePredicate sync.Once

func Get_Data_Divide_dividePredicate() gopurs_runtime.Value {
	once_Data_Divide_dividePredicate.Do(func() {
		cache_Data_Divide_dividePredicate = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divide_Divide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]](Get_Data_Predicate_contravariantPredicate()))}
		}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_4_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope5), (TypeVar c$scope6)])
			v2_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a_3))
			_ = v2_4_0
			return gopurs_runtime.Bool(((gopurs_runtime.Apply(v_1, (v2_4_0).V0).IntVal) != (0)) && ((gopurs_runtime.Apply(v1_2, (v2_4_0).V1).IntVal) != (0)))
		})}))}
	})
	return cache_Data_Divide_dividePredicate
}

var cache_Data_Divide_divideOp gopurs_runtime.Value
var once_Data_Divide_divideOp sync.Once

func Get_Data_Divide_divideOp() gopurs_runtime.Value {
	once_Data_Divide_divideOp.Do(func() {
		cache_Data_Divide_divideOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Divide_divideOp(dictSemigroup_0_box)
		})
	})
	return cache_Data_Divide_divideOp
}

var cache_Data_Divide_divideEquivalence gopurs_runtime.Value
var once_Data_Divide_divideEquivalence sync.Once

func Get_Data_Divide_divideEquivalence() gopurs_runtime.Value {
	once_Data_Divide_divideEquivalence.Do(func() {
		cache_Data_Divide_divideEquivalence = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divide_Divide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]](Get_Data_Equivalence_contravariantEquivalence()))}
		}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_5_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope24), (TypeVar c$scope25)])
			v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a_3))
			_ = v2_5_0
			// TAST (Let): v3_6_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope24), (TypeVar c$scope25)])
			v3_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_4))
			_ = v3_6_1
			return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_1, (v2_5_0).V0, (v3_6_1).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(v1_2, (v2_5_0).V1, (v3_6_1).V1).IntVal) != (0)))
		})}))}
	})
	return cache_Data_Divide_divideEquivalence
}

var cache_Data_Divide_divideComparison gopurs_runtime.Value
var once_Data_Divide_divideComparison sync.Once

func Get_Data_Divide_divideComparison() gopurs_runtime.Value {
	once_Data_Divide_divideComparison.Do(func() {
		cache_Data_Divide_divideComparison = gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divide_Divide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]](Get_Data_Comparison_contravariantComparison()))}
		}), gopurs_runtime.Func5(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_5_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope32), (TypeVar c$scope33)])
			v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a_3))
			_ = v2_5_0
			// TAST (Let): v3_6_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope32), (TypeVar c$scope33)])
			v3_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_4))
			_ = v3_6_1
			// TAST (Let): __local_var_7_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			__local_var_7_2 := uint32(gopurs_runtime.Apply2(v_1, (v2_5_0).V0, (v3_6_1).V0).IntVal)
			_ = __local_var_7_2
			var __t3 gopurs_runtime.Value
			{
				if __local_var_7_2 == 1527465420 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_7_2 == 380165415 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				if __local_var_7_2 == 902936544 {
					__t3 = gopurs_runtime.Apply2(v1_2, (v2_5_0).V1, (v3_6_1).V1)
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_3:
			return __t3
		})}))}
	})
	return cache_Data_Divide_divideComparison
}

var cache_Data_Divide_divide gopurs_runtime.Value
var once_Data_Divide_divide sync.Once

func Get_Data_Divide_divide() gopurs_runtime.Value {
	once_Data_Divide_divide.Do(func() {
		cache_Data_Divide_divide = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Divide_divide(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Divide_divide
}

var cache_Data_Divide_divided gopurs_runtime.Value
var once_Data_Divide_divided sync.Once

func Get_Data_Divide_divided() gopurs_runtime.Value {
	once_Data_Divide_divided.Do(func() {
		cache_Data_Divide_divided = gopurs_runtime.Func(func(dictDivide_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Divide_divided(gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](dictDivide_0_box))
		})
	})
	return cache_Data_Divide_divided
}

type Constructor_Data_Divide_Divide[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004400690076006900640065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004400690076006900640065005f004400690076006900640065
func init() {
	gopurs_runtime.StructGetters[2642321722] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Divide_Divide[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Contravariant0":
			return gopurs_runtime.Box(c.V0)
		case "divide":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Divide_Divide: " + key)
		}
	}
}

func Call_Data_Divide_identity(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var x_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_0)}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Data_Divide_Divide_dollar_Dict(x_0_loop struct {
	Contravariant0 gopurs_runtime.Value
	divide         gopurs_runtime.Value
	/*gopurs:record-v1:2:0043006f006e00740072006100760061007200690061006e00740030,006400690076006900640065*/
}) *Constructor_Data_Divide_Divide[gopurs_runtime.Value] {
	var x_0 struct {
		Contravariant0 gopurs_runtime.Value
		divide         gopurs_runtime.Value
		/*gopurs:record-v1:2:0043006f006e00740072006100760061007200690061006e00740030,006400690076006900640065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Divide_Divide[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Contravariant0", "divide", orig.Contravariant0, orig.divide)
	}())
}

func Call_Data_Divide_divideOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2642321722, UnsafePtr: unsafe.Pointer((&Constructor_Data_Divide_Divide[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Op_contravariantOp()
	}), gopurs_runtime.Func4(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v2_5_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope15), (TypeVar c$scope16)])
		v2_5_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, a_4))
		_ = v2_5_0
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_2, (v2_5_0).V0), gopurs_runtime.Apply(v1_3, (v2_5_0).V1))
	})}))}
}

func Call_Data_Divide_divide(dict_0_loop *Constructor_Data_Divide_Divide[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Divide_Divide[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Divide_divided(dictDivide_0_loop *Constructor_Data_Divide_Divide[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDivide_0 *Constructor_Data_Divide_Divide[gopurs_runtime.Value] = dictDivide_0_loop
	_ = dictDivide_0
	return gopurs_runtime.Apply(dictDivide_0.V1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}
