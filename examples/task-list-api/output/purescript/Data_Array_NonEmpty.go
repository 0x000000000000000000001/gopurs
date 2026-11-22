package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_NonEmpty_coerce gopurs_runtime.Value
var once_Data_Array_NonEmpty_coerce sync.Once

func Get_Data_Array_NonEmpty_coerce() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_coerce.Do(func() {
		cache_Data_Array_NonEmpty_coerce = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Data_Array_NonEmpty_coerce
}

var cache_Data_Array_NonEmpty_compose gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose sync.Once

func Get_Data_Array_NonEmpty_compose() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose.Do(func() {
		cache_Data_Array_NonEmpty_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose
}

var cache_Data_Array_NonEmpty_compose1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose1 sync.Once

func Get_Data_Array_NonEmpty_compose1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose1.Do(func() {
		cache_Data_Array_NonEmpty_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose1
}

var cache_Data_Array_NonEmpty_compose2 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose2 sync.Once

func Get_Data_Array_NonEmpty_compose2() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose2.Do(func() {
		cache_Data_Array_NonEmpty_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose2
}

var cache_Data_Array_NonEmpty_compose3 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose3 sync.Once

func Get_Data_Array_NonEmpty_compose3() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose3.Do(func() {
		cache_Data_Array_NonEmpty_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose3
}

var cache_Data_Array_NonEmpty_compose4 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose4 sync.Once

func Get_Data_Array_NonEmpty_compose4() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose4.Do(func() {
		cache_Data_Array_NonEmpty_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose4
}

var cache_Data_Array_NonEmpty_max gopurs_runtime.Value
var once_Data_Array_NonEmpty_max sync.Once

func Get_Data_Array_NonEmpty_max() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_max.Do(func() {
		cache_Data_Array_NonEmpty_max = Get_Data_Ord_max__4009478382()
	})
	return cache_Data_Array_NonEmpty_max
}

var cache_Data_Array_NonEmpty_go__append gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__append sync.Once

func Get_Data_Array_NonEmpty_go__append() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__append.Do(func() {
		cache_Data_Array_NonEmpty_go__append = Get_Data_Semigroup_concatArray()
	})
	return cache_Data_Array_NonEmpty_go__append
}

var cache_Data_Array_NonEmpty_compose5 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose5 sync.Once

func Get_Data_Array_NonEmpty_compose5() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose5.Do(func() {
		cache_Data_Array_NonEmpty_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose5
}

var cache_Data_Array_NonEmpty_compose6 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose6 sync.Once

func Get_Data_Array_NonEmpty_compose6() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose6.Do(func() {
		cache_Data_Array_NonEmpty_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose6
}

var cache_Data_Array_NonEmpty_compose7 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose7 sync.Once

func Get_Data_Array_NonEmpty_compose7() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose7.Do(func() {
		cache_Data_Array_NonEmpty_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose7
}

var cache_Data_Array_NonEmpty_intercalate1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_intercalate1 sync.Once

func Get_Data_Array_NonEmpty_intercalate1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intercalate1.Do(func() {
		cache_Data_Array_NonEmpty_intercalate1 = gopurs_runtime.Apply(Get_Data_Semigroup_Foldable_intercalate(), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_3016437835_4151366573(Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))))})
	})
	return cache_Data_Array_NonEmpty_intercalate1
}

var cache_Data_Array_NonEmpty_compose8 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose8 sync.Once

func Get_Data_Array_NonEmpty_compose8() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose8.Do(func() {
		cache_Data_Array_NonEmpty_compose8 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose8
}

var cache_Data_Array_NonEmpty_compose9 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose9 sync.Once

func Get_Data_Array_NonEmpty_compose9() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose9.Do(func() {
		cache_Data_Array_NonEmpty_compose9 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose9
}

var cache_Data_Array_NonEmpty_compose10 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose10 sync.Once

func Get_Data_Array_NonEmpty_compose10() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose10.Do(func() {
		cache_Data_Array_NonEmpty_compose10 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose10
}

var cache_Data_Array_NonEmpty_greaterThan gopurs_runtime.Value
var once_Data_Array_NonEmpty_greaterThan sync.Once

func Get_Data_Array_NonEmpty_greaterThan() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_greaterThan.Do(func() {
		cache_Data_Array_NonEmpty_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_Array_NonEmpty_greaterThan
}

var cache_Data_Array_NonEmpty_compose11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose11 sync.Once

func Get_Data_Array_NonEmpty_compose11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose11.Do(func() {
		cache_Data_Array_NonEmpty_compose11 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose11
}

var cache_Data_Array_NonEmpty_foldMap11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldMap11 sync.Once

func Get_Data_Array_NonEmpty_foldMap11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldMap11.Do(func() {
		cache_Data_Array_NonEmpty_foldMap11 = Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray())).V1
	})
	return cache_Data_Array_NonEmpty_foldMap11
}

var cache_Data_Array_NonEmpty_fold11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fold11 sync.Once

func Get_Data_Array_NonEmpty_fold11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fold11.Do(func() {
		cache_Data_Array_NonEmpty_fold11 = gopurs_runtime.Apply(Get_Data_Semigroup_Foldable_fold1(), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_3016437835_4151366573(Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))))})
	})
	return cache_Data_Array_NonEmpty_fold11
}

var cache_Data_Array_NonEmpty_append1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_append1 sync.Once

func Get_Data_Array_NonEmpty_append1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_append1.Do(func() {
		cache_Data_Array_NonEmpty_append1 = Get_Data_Semigroup_concatArray()
	})
	return cache_Data_Array_NonEmpty_append1
}

var cache_Data_Array_NonEmpty_compose12 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose12 sync.Once

func Get_Data_Array_NonEmpty_compose12() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose12.Do(func() {
		cache_Data_Array_NonEmpty_compose12 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose12
}

var cache_Data_Array_NonEmpty_compose13 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose13 sync.Once

func Get_Data_Array_NonEmpty_compose13() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose13.Do(func() {
		cache_Data_Array_NonEmpty_compose13 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose13
}

var cache_Data_Array_NonEmpty_fromJust gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromJust sync.Once

func Get_Data_Array_NonEmpty_fromJust() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromJust.Do(func() {
		cache_Data_Array_NonEmpty_fromJust = gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})
	})
	return cache_Data_Array_NonEmpty_fromJust
}

var cache_Data_Array_NonEmpty_compose14 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose14 sync.Once

func Get_Data_Array_NonEmpty_compose14() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose14.Do(func() {
		cache_Data_Array_NonEmpty_compose14 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose14
}

var cache_Data_Array_NonEmpty_compose15 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose15 sync.Once

func Get_Data_Array_NonEmpty_compose15() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose15.Do(func() {
		cache_Data_Array_NonEmpty_compose15 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose15
}

var cache_Data_Array_NonEmpty_compose16 gopurs_runtime.Value
var once_Data_Array_NonEmpty_compose16 sync.Once

func Get_Data_Array_NonEmpty_compose16() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_compose16.Do(func() {
		cache_Data_Array_NonEmpty_compose16 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Array_NonEmpty_compose16
}

var cache_Data_Array_NonEmpty_unsafeIndex1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeIndex1 sync.Once

func Get_Data_Array_NonEmpty_unsafeIndex1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeIndex1.Do(func() {
		cache_Data_Array_NonEmpty_unsafeIndex1 = gopurs_runtime.Apply(Get_Data_Array_unsafeIndex(), gopurs_runtime.Value{})
	})
	return cache_Data_Array_NonEmpty_unsafeIndex1
}

var cache_Data_Array_NonEmpty_lessThan gopurs_runtime.Value
var once_Data_Array_NonEmpty_lessThan sync.Once

func Get_Data_Array_NonEmpty_lessThan() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_lessThan.Do(func() {
		cache_Data_Array_NonEmpty_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_Array_NonEmpty_lessThan
}

var cache_Data_Array_NonEmpty_sub gopurs_runtime.Value
var once_Data_Array_NonEmpty_sub sync.Once

func Get_Data_Array_NonEmpty_sub() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sub.Do(func() {
		cache_Data_Array_NonEmpty_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Array_NonEmpty_sub
}

var cache_Data_Array_NonEmpty_add gopurs_runtime.Value
var once_Data_Array_NonEmpty_add sync.Once

func Get_Data_Array_NonEmpty_add() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_add.Do(func() {
		cache_Data_Array_NonEmpty_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Array_NonEmpty_add
}

var cache_Data_Array_NonEmpty_unsafeFromArrayF gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArrayF sync.Once

func Get_Data_Array_NonEmpty_unsafeFromArrayF() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArrayF.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArrayF = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArrayF
}

var cache_Data_Array_NonEmpty_unsafeFromArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray sync.Once

func Get_Data_Array_NonEmpty_unsafeFromArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray
}

var cache_Data_Array_NonEmpty_unsafeFromArray__4190607066 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__4190607066 sync.Once

func Get_Data_Array_NonEmpty_unsafeFromArray__4190607066() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__4190607066.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__4190607066 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unsafeFromArray__4190607066(*((*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__4190607066
}

var cache_Data_Array_NonEmpty_unsafeFromArray__4020493786 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__4020493786 sync.Once

func Get_Data_Array_NonEmpty_unsafeFromArray__4020493786() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__4020493786.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__4020493786 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_NonEmpty_unsafeFromArray__4020493786(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__4020493786
}

var cache_Data_Array_NonEmpty_transpose gopurs_runtime.Value
var once_Data_Array_NonEmpty_transpose sync.Once

func Get_Data_Array_NonEmpty_transpose() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_transpose.Do(func() {
		cache_Data_Array_NonEmpty_transpose = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_transpose(), Call_Safe_Coerce_coerce(gopurs_runtime.Value{})))
	})
	return cache_Data_Array_NonEmpty_transpose
}

var cache_Data_Array_NonEmpty_toArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray sync.Once

func Get_Data_Array_NonEmpty_toArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray.Do(func() {
		cache_Data_Array_NonEmpty_toArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_toArray(*((*[]gopurs_runtime.Value)(v_0_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_toArray
}

var cache_Data_Array_NonEmpty_unionBy_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy_prime_ sync.Once

func Get_Data_Array_NonEmpty_unionBy_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy_prime_.Do(func() {
		cache_Data_Array_NonEmpty_unionBy_prime_ = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_unionBy_prime_(eq_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_NonEmpty_unionBy_prime_
}

var cache_Data_Array_NonEmpty_union_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_union_prime_ sync.Once

func Get_Data_Array_NonEmpty_union_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_union_prime_.Do(func() {
		cache_Data_Array_NonEmpty_union_prime_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_union_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_union_prime_
}

var cache_Data_Array_NonEmpty_unionBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy sync.Once

func Get_Data_Array_NonEmpty_unionBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy.Do(func() {
		cache_Data_Array_NonEmpty_unionBy = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_unionBy(eq_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_NonEmpty_unionBy
}

var cache_Data_Array_NonEmpty_union gopurs_runtime.Value
var once_Data_Array_NonEmpty_union sync.Once

func Get_Data_Array_NonEmpty_union() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_union.Do(func() {
		cache_Data_Array_NonEmpty_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_union
}

var cache_Data_Array_NonEmpty_unzip gopurs_runtime.Value
var once_Data_Array_NonEmpty_unzip sync.Once

func Get_Data_Array_NonEmpty_unzip() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unzip.Do(func() {
		cache_Data_Array_NonEmpty_unzip = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0), gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_unzip(), Get_Data_Array_NonEmpty_toArray()))
	})
	return cache_Data_Array_NonEmpty_unzip
}

var cache_Data_Array_NonEmpty_updateAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_updateAt sync.Once

func Get_Data_Array_NonEmpty_updateAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_updateAt.Do(func() {
		cache_Data_Array_NonEmpty_updateAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_updateAt(i_0_box.IntVal, x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_updateAt
}

var cache_Data_Array_NonEmpty_zip gopurs_runtime.Value
var once_Data_Array_NonEmpty_zip sync.Once

func Get_Data_Array_NonEmpty_zip() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_zip.Do(func() {
		cache_Data_Array_NonEmpty_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_NonEmpty_zip(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)))
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Array_NonEmpty_zip
}

var cache_Data_Array_NonEmpty_zipWith gopurs_runtime.Value
var once_Data_Array_NonEmpty_zipWith sync.Once

func Get_Data_Array_NonEmpty_zipWith() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_zipWith.Do(func() {
		cache_Data_Array_NonEmpty_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_zipWith(f_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_zipWith
}

var cache_Data_Array_NonEmpty_zipWithA gopurs_runtime.Value
var once_Data_Array_NonEmpty_zipWithA sync.Once

func Get_Data_Array_NonEmpty_zipWithA() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_zipWithA.Do(func() {
		cache_Data_Array_NonEmpty_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_zipWithA
}

var cache_Data_Array_NonEmpty_splitAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_splitAt sync.Once

func Get_Data_Array_NonEmpty_splitAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_splitAt.Do(func() {
		cache_Data_Array_NonEmpty_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_splitAt(i_0_box.IntVal, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(orig.after), gopurs_runtime.Array(orig.before))
			}()
		})
	})
	return cache_Data_Array_NonEmpty_splitAt
}

var cache_Data_Array_NonEmpty_some gopurs_runtime.Value
var once_Data_Array_NonEmpty_some sync.Once

func Get_Data_Array_NonEmpty_some() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_some.Do(func() {
		cache_Data_Array_NonEmpty_some = gopurs_runtime.Func2(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_1_box))
		})
	})
	return cache_Data_Array_NonEmpty_some
}

var cache_Data_Array_NonEmpty_snoc_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_snoc_prime_ sync.Once

func Get_Data_Array_NonEmpty_snoc_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_snoc_prime_.Do(func() {
		cache_Data_Array_NonEmpty_snoc_prime_ = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_snoc_prime_(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)), x_1_box))
		})
	})
	return cache_Data_Array_NonEmpty_snoc_prime_
}

var cache_Data_Array_NonEmpty_snoc gopurs_runtime.Value
var once_Data_Array_NonEmpty_snoc sync.Once

func Get_Data_Array_NonEmpty_snoc() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_snoc.Do(func() {
		cache_Data_Array_NonEmpty_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_snoc(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)), x_1_box))
		})
	})
	return cache_Data_Array_NonEmpty_snoc
}

var cache_Data_Array_NonEmpty_singleton gopurs_runtime.Value
var once_Data_Array_NonEmpty_singleton sync.Once

func Get_Data_Array_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_singleton.Do(func() {
		cache_Data_Array_NonEmpty_singleton = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), Get_Data_Array_singleton())
	})
	return cache_Data_Array_NonEmpty_singleton
}

var cache_Data_Array_NonEmpty_replicate gopurs_runtime.Value
var once_Data_Array_NonEmpty_replicate sync.Once

func Get_Data_Array_NonEmpty_replicate() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_replicate.Do(func() {
		cache_Data_Array_NonEmpty_replicate = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_replicate(i_0_box.IntVal, x_1_box))
		})
	})
	return cache_Data_Array_NonEmpty_replicate
}

var cache_Data_Array_NonEmpty_go__range gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__range sync.Once

func Get_Data_Array_NonEmpty_go__range() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__range.Do(func() {
		cache_Data_Array_NonEmpty_go__range = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_NonEmpty_go__range(x_0_box.IntVal, y_1_box.IntVal)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Array_NonEmpty_go__range
}

var cache_Data_Array_NonEmpty_prependArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_prependArray sync.Once

func Get_Data_Array_NonEmpty_prependArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_prependArray.Do(func() {
		cache_Data_Array_NonEmpty_prependArray = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_prependArray(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_prependArray
}

var cache_Data_Array_NonEmpty_modifyAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_modifyAt sync.Once

func Get_Data_Array_NonEmpty_modifyAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_modifyAt.Do(func() {
		cache_Data_Array_NonEmpty_modifyAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_modifyAt(i_0_box.IntVal, f_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_modifyAt
}

var cache_Data_Array_NonEmpty_intersectBy_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy_prime_ sync.Once

func Get_Data_Array_NonEmpty_intersectBy_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy_prime_.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy_prime_ = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_intersectBy_prime_(eq_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_NonEmpty_intersectBy_prime_
}

var cache_Data_Array_NonEmpty_intersectBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy sync.Once

func Get_Data_Array_NonEmpty_intersectBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_intersectBy(eq_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_NonEmpty_intersectBy
}

var cache_Data_Array_NonEmpty_intersect_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersect_prime_ sync.Once

func Get_Data_Array_NonEmpty_intersect_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersect_prime_.Do(func() {
		cache_Data_Array_NonEmpty_intersect_prime_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_intersect_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_intersect_prime_
}

var cache_Data_Array_NonEmpty_intersect gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersect sync.Once

func Get_Data_Array_NonEmpty_intersect() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersect.Do(func() {
		cache_Data_Array_NonEmpty_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_intersect
}

var cache_Data_Array_NonEmpty_intercalate gopurs_runtime.Value
var once_Data_Array_NonEmpty_intercalate sync.Once

func Get_Data_Array_NonEmpty_intercalate() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intercalate.Do(func() {
		cache_Data_Array_NonEmpty_intercalate = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_intercalate
}

var cache_Data_Array_NonEmpty_insertAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_insertAt sync.Once

func Get_Data_Array_NonEmpty_insertAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insertAt.Do(func() {
		cache_Data_Array_NonEmpty_insertAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_insertAt(i_0_box.IntVal, x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_insertAt
}

var cache_Data_Array_NonEmpty_fromFoldable1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromFoldable1 sync.Once

func Get_Data_Array_NonEmpty_fromFoldable1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromFoldable1.Do(func() {
		cache_Data_Array_NonEmpty_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_fromFoldable1
}

var cache_Data_Array_NonEmpty_fromArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray sync.Once

func Get_Data_Array_NonEmpty_fromArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray.Do(func() {
		cache_Data_Array_NonEmpty_fromArray = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_fromArray(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Array_NonEmpty_fromArray
}

var cache_Data_Array_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromFoldable sync.Once

func Get_Data_Array_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromFoldable.Do(func() {
		cache_Data_Array_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_fromFoldable
}

var cache_Data_Array_NonEmpty_transpose_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_transpose_prime_ sync.Once

func Get_Data_Array_NonEmpty_transpose_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_transpose_prime_.Do(func() {
		cache_Data_Array_NonEmpty_transpose_prime_ = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_fromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_transpose(), Call_Safe_Coerce_coerce(gopurs_runtime.Value{})))
	})
	return cache_Data_Array_NonEmpty_transpose_prime_
}

var cache_Data_Array_NonEmpty_foldr1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldr1 sync.Once

func Get_Data_Array_NonEmpty_foldr1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldr1.Do(func() {
		cache_Data_Array_NonEmpty_foldr1 = Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray())).V3
	})
	return cache_Data_Array_NonEmpty_foldr1
}

var cache_Data_Array_NonEmpty_foldl1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldl1 sync.Once

func Get_Data_Array_NonEmpty_foldl1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldl1.Do(func() {
		cache_Data_Array_NonEmpty_foldl1 = Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray())).V2
	})
	return cache_Data_Array_NonEmpty_foldl1
}

var cache_Data_Array_NonEmpty_foldMap1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldMap1 sync.Once

func Get_Data_Array_NonEmpty_foldMap1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldMap1.Do(func() {
		cache_Data_Array_NonEmpty_foldMap1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_foldMap1
}

var cache_Data_Array_NonEmpty_fold1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fold1 sync.Once

func Get_Data_Array_NonEmpty_fold1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fold1.Do(func() {
		cache_Data_Array_NonEmpty_fold1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_fold1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_fold1
}

var cache_Data_Array_NonEmpty_difference_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_difference_prime_ sync.Once

func Get_Data_Array_NonEmpty_difference_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_difference_prime_.Do(func() {
		cache_Data_Array_NonEmpty_difference_prime_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_difference_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_difference_prime_
}

var cache_Data_Array_NonEmpty_cons_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons_prime_ sync.Once

func Get_Data_Array_NonEmpty_cons_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons_prime_.Do(func() {
		cache_Data_Array_NonEmpty_cons_prime_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_cons_prime_(x_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_cons_prime_
}

var cache_Data_Array_NonEmpty_cons_prime___3391444928 gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons_prime___3391444928 sync.Once

func Get_Data_Array_NonEmpty_cons_prime___3391444928() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons_prime___3391444928.Do(func() {
		cache_Data_Array_NonEmpty_cons_prime___3391444928 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_cons_prime___3391444928(x_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_cons_prime___3391444928
}

var cache_Data_Array_NonEmpty_fromNonEmpty gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromNonEmpty sync.Once

func Get_Data_Array_NonEmpty_fromNonEmpty() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromNonEmpty.Do(func() {
		cache_Data_Array_NonEmpty_fromNonEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_fromNonEmpty(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
		})
	})
	return cache_Data_Array_NonEmpty_fromNonEmpty
}

var cache_Data_Array_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_Array_NonEmpty_concatMap sync.Once

func Get_Data_Array_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_concatMap.Do(func() {
		cache_Data_Array_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_concatMap(b_0_box, *((*[]gopurs_runtime.Value)(a_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_concatMap
}

var cache_Data_Array_NonEmpty_concat gopurs_runtime.Value
var once_Data_Array_NonEmpty_concat sync.Once

func Get_Data_Array_NonEmpty_concat() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_concat.Do(func() {
		cache_Data_Array_NonEmpty_concat = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_concat(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_toArray(), gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Array_NonEmpty_toArray()))))
	})
	return cache_Data_Array_NonEmpty_concat
}

var cache_Data_Array_NonEmpty_appendArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_appendArray sync.Once

func Get_Data_Array_NonEmpty_appendArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_appendArray.Do(func() {
		cache_Data_Array_NonEmpty_appendArray = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_appendArray(*((*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)), *((*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_appendArray
}

var cache_Data_Array_NonEmpty_alterAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_alterAt sync.Once

func Get_Data_Array_NonEmpty_alterAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_alterAt.Do(func() {
		cache_Data_Array_NonEmpty_alterAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_alterAt(i_0_box.IntVal, f_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_alterAt
}

var cache_Data_Array_NonEmpty_adaptMaybe gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe sync.Once

func Get_Data_Array_NonEmpty_adaptMaybe() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_adaptMaybe(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe
}

var cache_Data_Array_NonEmpty_head gopurs_runtime.Value
var once_Data_Array_NonEmpty_head sync.Once

func Get_Data_Array_NonEmpty_head() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_head.Do(func() {
		cache_Data_Array_NonEmpty_head = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_head())
	})
	return cache_Data_Array_NonEmpty_head
}

var cache_Data_Array_NonEmpty_go__init gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__init sync.Once

func Get_Data_Array_NonEmpty_go__init() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__init.Do(func() {
		cache_Data_Array_NonEmpty_go__init = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_go__init())
	})
	return cache_Data_Array_NonEmpty_go__init
}

var cache_Data_Array_NonEmpty_last gopurs_runtime.Value
var once_Data_Array_NonEmpty_last sync.Once

func Get_Data_Array_NonEmpty_last() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_last.Do(func() {
		cache_Data_Array_NonEmpty_last = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_last())
	})
	return cache_Data_Array_NonEmpty_last
}

var cache_Data_Array_NonEmpty_tail gopurs_runtime.Value
var once_Data_Array_NonEmpty_tail sync.Once

func Get_Data_Array_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_tail.Do(func() {
		cache_Data_Array_NonEmpty_tail = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_tail())
	})
	return cache_Data_Array_NonEmpty_tail
}

var cache_Data_Array_NonEmpty_uncons gopurs_runtime.Value
var once_Data_Array_NonEmpty_uncons sync.Once

func Get_Data_Array_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_uncons.Do(func() {
		cache_Data_Array_NonEmpty_uncons = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_uncons())
	})
	return cache_Data_Array_NonEmpty_uncons
}

var cache_Data_Array_NonEmpty_toNonEmpty gopurs_runtime.Value
var once_Data_Array_NonEmpty_toNonEmpty sync.Once

func Get_Data_Array_NonEmpty_toNonEmpty() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toNonEmpty.Do(func() {
		cache_Data_Array_NonEmpty_toNonEmpty = gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_uncons(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(v_0, "head"), gopurs_runtime.RecordGet(v_0, "tail")}))}
		}))
	})
	return cache_Data_Array_NonEmpty_toNonEmpty
}

var cache_Data_Array_NonEmpty_unsnoc gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsnoc sync.Once

func Get_Data_Array_NonEmpty_unsnoc() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsnoc.Do(func() {
		cache_Data_Array_NonEmpty_unsnoc = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_unsnoc())
	})
	return cache_Data_Array_NonEmpty_unsnoc
}

var cache_Data_Array_NonEmpty_adaptAny gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny sync.Once

func Get_Data_Array_NonEmpty_adaptAny() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_adaptAny(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_adaptAny
}

var cache_Data_Array_NonEmpty_all gopurs_runtime.Value
var once_Data_Array_NonEmpty_all sync.Once

func Get_Data_Array_NonEmpty_all() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_all.Do(func() {
		cache_Data_Array_NonEmpty_all = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_all(p_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_all
}

var cache_Data_Array_NonEmpty_any gopurs_runtime.Value
var once_Data_Array_NonEmpty_any sync.Once

func Get_Data_Array_NonEmpty_any() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_any.Do(func() {
		cache_Data_Array_NonEmpty_any = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_any(p_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_any
}

var cache_Data_Array_NonEmpty_catMaybes gopurs_runtime.Value
var once_Data_Array_NonEmpty_catMaybes sync.Once

func Get_Data_Array_NonEmpty_catMaybes() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_catMaybes.Do(func() {
		cache_Data_Array_NonEmpty_catMaybes = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_catMaybes(), Get_Data_Array_NonEmpty_toArray())
	})
	return cache_Data_Array_NonEmpty_catMaybes
}

var cache_Data_Array_NonEmpty_go__delete gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__delete sync.Once

func Get_Data_Array_NonEmpty_go__delete() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__delete.Do(func() {
		cache_Data_Array_NonEmpty_go__delete = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_go__delete
}

var cache_Data_Array_NonEmpty_deleteAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_deleteAt sync.Once

func Get_Data_Array_NonEmpty_deleteAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_deleteAt.Do(func() {
		cache_Data_Array_NonEmpty_deleteAt = gopurs_runtime.Func(func(i_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_deleteAt(i_0_box.IntVal)
		})
	})
	return cache_Data_Array_NonEmpty_deleteAt
}

var cache_Data_Array_NonEmpty_deleteBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_deleteBy sync.Once

func Get_Data_Array_NonEmpty_deleteBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_deleteBy.Do(func() {
		cache_Data_Array_NonEmpty_deleteBy = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_deleteBy(f_0_box, x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_deleteBy
}

var cache_Data_Array_NonEmpty_difference gopurs_runtime.Value
var once_Data_Array_NonEmpty_difference sync.Once

func Get_Data_Array_NonEmpty_difference() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_difference.Do(func() {
		cache_Data_Array_NonEmpty_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_difference
}

var cache_Data_Array_NonEmpty_drop gopurs_runtime.Value
var once_Data_Array_NonEmpty_drop sync.Once

func Get_Data_Array_NonEmpty_drop() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_drop.Do(func() {
		cache_Data_Array_NonEmpty_drop = gopurs_runtime.Func(func(i_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_drop(i_0_box.IntVal)
		})
	})
	return cache_Data_Array_NonEmpty_drop
}

var cache_Data_Array_NonEmpty_dropEnd gopurs_runtime.Value
var once_Data_Array_NonEmpty_dropEnd sync.Once

func Get_Data_Array_NonEmpty_dropEnd() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_dropEnd.Do(func() {
		cache_Data_Array_NonEmpty_dropEnd = gopurs_runtime.Func(func(i_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_dropEnd(i_0_box.IntVal)
		})
	})
	return cache_Data_Array_NonEmpty_dropEnd
}

var cache_Data_Array_NonEmpty_dropWhile gopurs_runtime.Value
var once_Data_Array_NonEmpty_dropWhile sync.Once

func Get_Data_Array_NonEmpty_dropWhile() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_dropWhile.Do(func() {
		cache_Data_Array_NonEmpty_dropWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_dropWhile(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_dropWhile
}

var cache_Data_Array_NonEmpty_elem gopurs_runtime.Value
var once_Data_Array_NonEmpty_elem sync.Once

func Get_Data_Array_NonEmpty_elem() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elem.Do(func() {
		cache_Data_Array_NonEmpty_elem = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_elem
}

var cache_Data_Array_NonEmpty_elemIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_elemIndex sync.Once

func Get_Data_Array_NonEmpty_elemIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elemIndex.Do(func() {
		cache_Data_Array_NonEmpty_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_elemIndex
}

var cache_Data_Array_NonEmpty_elemLastIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_elemLastIndex sync.Once

func Get_Data_Array_NonEmpty_elemLastIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elemLastIndex.Do(func() {
		cache_Data_Array_NonEmpty_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_elemLastIndex
}

var cache_Data_Array_NonEmpty_filter gopurs_runtime.Value
var once_Data_Array_NonEmpty_filter sync.Once

func Get_Data_Array_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_filter.Do(func() {
		cache_Data_Array_NonEmpty_filter = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_filter(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_filter
}

var cache_Data_Array_NonEmpty_filterA gopurs_runtime.Value
var once_Data_Array_NonEmpty_filterA sync.Once

func Get_Data_Array_NonEmpty_filterA() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_filterA.Do(func() {
		cache_Data_Array_NonEmpty_filterA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_filterA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_filterA
}

var cache_Data_Array_NonEmpty_find gopurs_runtime.Value
var once_Data_Array_NonEmpty_find sync.Once

func Get_Data_Array_NonEmpty_find() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_find.Do(func() {
		cache_Data_Array_NonEmpty_find = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_find(p_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_find
}

var cache_Data_Array_NonEmpty_findIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_findIndex sync.Once

func Get_Data_Array_NonEmpty_findIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findIndex.Do(func() {
		cache_Data_Array_NonEmpty_findIndex = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_findIndex(p_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_findIndex
}

var cache_Data_Array_NonEmpty_findLastIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_findLastIndex sync.Once

func Get_Data_Array_NonEmpty_findLastIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findLastIndex.Do(func() {
		cache_Data_Array_NonEmpty_findLastIndex = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_findLastIndex(x_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_findLastIndex
}

var cache_Data_Array_NonEmpty_findMap gopurs_runtime.Value
var once_Data_Array_NonEmpty_findMap sync.Once

func Get_Data_Array_NonEmpty_findMap() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findMap.Do(func() {
		cache_Data_Array_NonEmpty_findMap = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_findMap(p_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_findMap
}

var cache_Data_Array_NonEmpty_foldM gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldM sync.Once

func Get_Data_Array_NonEmpty_foldM() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldM.Do(func() {
		cache_Data_Array_NonEmpty_foldM = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, acc_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, acc_2_box)
		})
	})
	return cache_Data_Array_NonEmpty_foldM
}

var cache_Data_Array_NonEmpty_foldRecM gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldRecM sync.Once

func Get_Data_Array_NonEmpty_foldRecM() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldRecM.Do(func() {
		cache_Data_Array_NonEmpty_foldRecM = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_foldRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_foldRecM
}

var cache_Data_Array_NonEmpty_index gopurs_runtime.Value
var once_Data_Array_NonEmpty_index sync.Once

func Get_Data_Array_NonEmpty_index() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_index.Do(func() {
		cache_Data_Array_NonEmpty_index = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_index(), Get_Data_Array_NonEmpty_toArray())
	})
	return cache_Data_Array_NonEmpty_index
}

var cache_Data_Array_NonEmpty_length gopurs_runtime.Value
var once_Data_Array_NonEmpty_length sync.Once

func Get_Data_Array_NonEmpty_length() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_length.Do(func() {
		cache_Data_Array_NonEmpty_length = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_length(), Get_Data_Array_NonEmpty_toArray())
	})
	return cache_Data_Array_NonEmpty_length
}

var cache_Data_Array_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_Array_NonEmpty_mapMaybe sync.Once

func Get_Data_Array_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_mapMaybe.Do(func() {
		cache_Data_Array_NonEmpty_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_mapMaybe(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_mapMaybe
}

var cache_Data_Array_NonEmpty_notElem gopurs_runtime.Value
var once_Data_Array_NonEmpty_notElem sync.Once

func Get_Data_Array_NonEmpty_notElem() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_notElem.Do(func() {
		cache_Data_Array_NonEmpty_notElem = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_notElem
}

var cache_Data_Array_NonEmpty_partition gopurs_runtime.Value
var once_Data_Array_NonEmpty_partition sync.Once

func Get_Data_Array_NonEmpty_partition() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_partition.Do(func() {
		cache_Data_Array_NonEmpty_partition = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_partition(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_partition
}

var cache_Data_Array_NonEmpty_slice gopurs_runtime.Value
var once_Data_Array_NonEmpty_slice sync.Once

func Get_Data_Array_NonEmpty_slice() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_slice.Do(func() {
		cache_Data_Array_NonEmpty_slice = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_slice(start_0_box.IntVal, end_1_box.IntVal)
		})
	})
	return cache_Data_Array_NonEmpty_slice
}

var cache_Data_Array_NonEmpty_span gopurs_runtime.Value
var once_Data_Array_NonEmpty_span sync.Once

func Get_Data_Array_NonEmpty_span() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_span.Do(func() {
		cache_Data_Array_NonEmpty_span = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_span(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_span
}

var cache_Data_Array_NonEmpty_take gopurs_runtime.Value
var once_Data_Array_NonEmpty_take sync.Once

func Get_Data_Array_NonEmpty_take() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_take.Do(func() {
		cache_Data_Array_NonEmpty_take = gopurs_runtime.Func(func(i_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_take(i_0_box.IntVal)
		})
	})
	return cache_Data_Array_NonEmpty_take
}

var cache_Data_Array_NonEmpty_takeEnd gopurs_runtime.Value
var once_Data_Array_NonEmpty_takeEnd sync.Once

func Get_Data_Array_NonEmpty_takeEnd() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_takeEnd.Do(func() {
		cache_Data_Array_NonEmpty_takeEnd = gopurs_runtime.Func(func(i_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_takeEnd(i_0_box.IntVal)
		})
	})
	return cache_Data_Array_NonEmpty_takeEnd
}

var cache_Data_Array_NonEmpty_takeWhile gopurs_runtime.Value
var once_Data_Array_NonEmpty_takeWhile sync.Once

func Get_Data_Array_NonEmpty_takeWhile() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_takeWhile.Do(func() {
		cache_Data_Array_NonEmpty_takeWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_takeWhile(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_takeWhile
}

var cache_Data_Array_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_Array_NonEmpty_toUnfoldable sync.Once

func Get_Data_Array_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_Array_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_toUnfoldable
}

var cache_Data_Array_NonEmpty_unsafeAdapt gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeAdapt sync.Once

func Get_Data_Array_NonEmpty_unsafeAdapt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeAdapt.Do(func() {
		cache_Data_Array_NonEmpty_unsafeAdapt = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_unsafeAdapt(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_unsafeAdapt
}

var cache_Data_Array_NonEmpty_cons gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons sync.Once

func Get_Data_Array_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons.Do(func() {
		cache_Data_Array_NonEmpty_cons = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_cons(x_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_cons
}

var cache_Data_Array_NonEmpty_group gopurs_runtime.Value
var once_Data_Array_NonEmpty_group sync.Once

func Get_Data_Array_NonEmpty_group() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_group.Do(func() {
		cache_Data_Array_NonEmpty_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_group
}

var cache_Data_Array_NonEmpty_groupAllBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAllBy sync.Once

func Get_Data_Array_NonEmpty_groupAllBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAllBy.Do(func() {
		cache_Data_Array_NonEmpty_groupAllBy = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_groupAllBy(op_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_groupAllBy
}

var cache_Data_Array_NonEmpty_groupAll gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAll sync.Once

func Get_Data_Array_NonEmpty_groupAll() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAll.Do(func() {
		cache_Data_Array_NonEmpty_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_groupAll
}

var cache_Data_Array_NonEmpty_groupBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupBy sync.Once

func Get_Data_Array_NonEmpty_groupBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupBy.Do(func() {
		cache_Data_Array_NonEmpty_groupBy = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_groupBy(op_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_groupBy
}

var cache_Data_Array_NonEmpty_insert gopurs_runtime.Value
var once_Data_Array_NonEmpty_insert sync.Once

func Get_Data_Array_NonEmpty_insert() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insert.Do(func() {
		cache_Data_Array_NonEmpty_insert = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_insert
}

var cache_Data_Array_NonEmpty_insertBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_insertBy sync.Once

func Get_Data_Array_NonEmpty_insertBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insertBy.Do(func() {
		cache_Data_Array_NonEmpty_insertBy = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_insertBy(f_0_box, x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_insertBy
}

var cache_Data_Array_NonEmpty_intersperse gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersperse sync.Once

func Get_Data_Array_NonEmpty_intersperse() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersperse.Do(func() {
		cache_Data_Array_NonEmpty_intersperse = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_intersperse(x_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_intersperse
}

var cache_Data_Array_NonEmpty_mapWithIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_mapWithIndex sync.Once

func Get_Data_Array_NonEmpty_mapWithIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_mapWithIndex.Do(func() {
		cache_Data_Array_NonEmpty_mapWithIndex = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_mapWithIndex(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_mapWithIndex
}

var cache_Data_Array_NonEmpty_modifyAtIndices gopurs_runtime.Value
var once_Data_Array_NonEmpty_modifyAtIndices sync.Once

func Get_Data_Array_NonEmpty_modifyAtIndices() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_modifyAtIndices.Do(func() {
		cache_Data_Array_NonEmpty_modifyAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_modifyAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_modifyAtIndices
}

var cache_Data_Array_NonEmpty_nub gopurs_runtime.Value
var once_Data_Array_NonEmpty_nub sync.Once

func Get_Data_Array_NonEmpty_nub() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nub.Do(func() {
		cache_Data_Array_NonEmpty_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_nub
}

var cache_Data_Array_NonEmpty_nubBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_nubBy sync.Once

func Get_Data_Array_NonEmpty_nubBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nubBy.Do(func() {
		cache_Data_Array_NonEmpty_nubBy = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_nubBy(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_nubBy
}

var cache_Data_Array_NonEmpty_nubByEq gopurs_runtime.Value
var once_Data_Array_NonEmpty_nubByEq sync.Once

func Get_Data_Array_NonEmpty_nubByEq() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nubByEq.Do(func() {
		cache_Data_Array_NonEmpty_nubByEq = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_nubByEq(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_nubByEq
}

var cache_Data_Array_NonEmpty_nubEq gopurs_runtime.Value
var once_Data_Array_NonEmpty_nubEq sync.Once

func Get_Data_Array_NonEmpty_nubEq() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nubEq.Do(func() {
		cache_Data_Array_NonEmpty_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_nubEq
}

var cache_Data_Array_NonEmpty_reverse gopurs_runtime.Value
var once_Data_Array_NonEmpty_reverse sync.Once

func Get_Data_Array_NonEmpty_reverse() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_reverse.Do(func() {
		cache_Data_Array_NonEmpty_reverse = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_reverse(), Get_Data_Array_NonEmpty_toArray()))
	})
	return cache_Data_Array_NonEmpty_reverse
}

var cache_Data_Array_NonEmpty_scanl gopurs_runtime.Value
var once_Data_Array_NonEmpty_scanl sync.Once

func Get_Data_Array_NonEmpty_scanl() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_scanl.Do(func() {
		cache_Data_Array_NonEmpty_scanl = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_scanl(f_0_box, x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_scanl
}

var cache_Data_Array_NonEmpty_scanr gopurs_runtime.Value
var once_Data_Array_NonEmpty_scanr sync.Once

func Get_Data_Array_NonEmpty_scanr() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_scanr.Do(func() {
		cache_Data_Array_NonEmpty_scanr = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_scanr(f_0_box, x_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_scanr
}

var cache_Data_Array_NonEmpty_sort gopurs_runtime.Value
var once_Data_Array_NonEmpty_sort sync.Once

func Get_Data_Array_NonEmpty_sort() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sort.Do(func() {
		cache_Data_Array_NonEmpty_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_sort
}

var cache_Data_Array_NonEmpty_sortBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_sortBy sync.Once

func Get_Data_Array_NonEmpty_sortBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sortBy.Do(func() {
		cache_Data_Array_NonEmpty_sortBy = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_sortBy(f_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_sortBy
}

var cache_Data_Array_NonEmpty_sortWith gopurs_runtime.Value
var once_Data_Array_NonEmpty_sortWith sync.Once

func Get_Data_Array_NonEmpty_sortWith() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sortWith.Do(func() {
		cache_Data_Array_NonEmpty_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
		})
	})
	return cache_Data_Array_NonEmpty_sortWith
}

var cache_Data_Array_NonEmpty_updateAtIndices gopurs_runtime.Value
var once_Data_Array_NonEmpty_updateAtIndices sync.Once

func Get_Data_Array_NonEmpty_updateAtIndices() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_updateAtIndices.Do(func() {
		cache_Data_Array_NonEmpty_updateAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_updateAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
		})
	})
	return cache_Data_Array_NonEmpty_updateAtIndices
}

var cache_Data_Array_NonEmpty_unsafeIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeIndex sync.Once

func Get_Data_Array_NonEmpty_unsafeIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeIndex.Do(func() {
		cache_Data_Array_NonEmpty_unsafeIndex = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_unsafeIndex(_dollar___unused_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_unsafeIndex
}

var cache_Data_Array_NonEmpty_unsafeIndex2 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeIndex2 sync.Once

func Get_Data_Array_NonEmpty_unsafeIndex2() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeIndex2.Do(func() {
		cache_Data_Array_NonEmpty_unsafeIndex2 = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_unsafeIndex(), gopurs_runtime.Value{}), Get_Data_Array_NonEmpty_toArray())
	})
	return cache_Data_Array_NonEmpty_unsafeIndex2
}

var cache_Data_Array_NonEmpty_toUnfoldable1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toUnfoldable1 sync.Once

func Get_Data_Array_NonEmpty_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toUnfoldable1.Do(func() {
		cache_Data_Array_NonEmpty_toUnfoldable1 = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_NonEmpty_toUnfoldable1
}

func Call_Data_Array_NonEmpty_unsafeFromArray__4190607066(__eta_norm_0_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
unsafeFromArray__4190607066:
	for {
		if false {
			continue unsafeFromArray__4190607066
		}
		var __eta_norm_0_0 []gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Data_Array_NonEmpty_unsafeFromArray__4020493786(__eta_norm_0_0_loop []int64) []int64 {
unsafeFromArray__4020493786:
	for {
		if false {
			continue unsafeFromArray__4020493786
		}
		var __eta_norm_0_0 []int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Data_Array_NonEmpty_toArray(v_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var v_0 []gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_Array_NonEmpty_unionBy_prime_(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var eq_0 gopurs_runtime.Value = eq_0_loop
	_ = eq_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1)))
}

func Call_Data_Array_NonEmpty_union_prime_(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unionBy_prime_(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_NonEmpty_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var eq_0 gopurs_runtime.Value = eq_0_loop
	_ = eq_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1))), Get_Data_Array_NonEmpty_toArray())
}

func Call_Data_Array_NonEmpty_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unionBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_NonEmpty_updateAt(i_0_loop int64, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArrayF(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Array_updateAt(), gopurs_runtime.Int(i_0), x_1), Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_zip(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var ys_1 []gopurs_runtime.Value = ys_1_loop
	_ = ys_1
	return func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), func() gopurs_runtime.Value {
			arr := func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
				unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v)
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)}
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr)
		unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
		for i, v := range arr {
			unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v)
		}
		return unboxed
	}()
}

func Call_Data_Array_NonEmpty_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	var ys_2 []gopurs_runtime.Value = ys_2_loop
	_ = ys_2
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(ys_2)).UnsafePtr)))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): zipWithA1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope50), (TypeVar b$scope51)] (TypeApp (TypeVar m$scope49) [(TypeVar c$scope52)])), (Array (TypeVar a$scope50)), (Array (TypeVar b$scope51))] (TypeApp (TypeVar m$scope49) [(Array (TypeVar c$scope52))]))
	zipWithA1_1_0 := Call_Data_Array_zipWithA(dictApplicative_0)
	_ = zipWithA1_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArrayF(), gopurs_runtime.Apply3(zipWithA1_1_0, f_2, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(xs_3.UnsafePtr))), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(ys_4.UnsafePtr)))))
	})
}

func Call_Data_Array_NonEmpty_splitAt(i_0_loop int64, xs_1_loop []gopurs_runtime.Value) struct {
	after  []gopurs_runtime.Value
	before []gopurs_runtime.Value
	/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
} {
	var i_0 int64 = i_0_loop
	_ = i_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return func() struct {
		after  []gopurs_runtime.Value
		before []gopurs_runtime.Value
		/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
	} {
		orig := func() gopurs_runtime.Value {
			orig := Call_Data_Array_splitAt(i_0, xs_1)
			_ = orig
			return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(orig.after), gopurs_runtime.Array(orig.before))
		}()
		_ = orig
		clone := struct {
			after  []gopurs_runtime.Value
			before []gopurs_runtime.Value
			/*gopurs:record-v1:2:00610066007400650072,006200650066006f00720065*/
		}{}
		clone.after = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "after").UnsafePtr))
		clone.before = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "before").UnsafePtr))
		return clone
	}()
}

func Call_Data_Array_NonEmpty_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value], dictLazy_1_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
	_ = dictAlternative_0
	var dictLazy_1 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_1_loop
	_ = dictLazy_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArrayF(), gopurs_runtime.Apply2(Get_Data_Array_some(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(dictLazy_1)}))
}

func Call_Data_Array_NonEmpty_snoc_prime_(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(Call_Data_Array_snoc(xs_0, x_1))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_snoc(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(Call_Data_Array_snoc(xs_0, x_1))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_replicate(i_0_loop int64, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(Call_Data_Ord_max__4009478382(int64(1), i_0)), x_1).UnsafePtr)))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_go__range(x_0_loop int64, y_1_loop int64) []int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	var y_1 int64 = y_1_loop
	_ = y_1
	return func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1)).UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
}

func Call_Data_Array_NonEmpty_prependArray(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var ys_1 []gopurs_runtime.Value = ys_1_loop
	_ = ys_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArrayF(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Array_modifyAt(), gopurs_runtime.Int(i_0), f_1), Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_intersectBy_prime_(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var eq_0 gopurs_runtime.Value = eq_0_loop
	_ = eq_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Apply2(Get_Data_Array_intersectBy(), eq_0, gopurs_runtime.Array(xs_1))
}

func Call_Data_Array_NonEmpty_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var eq_0 gopurs_runtime.Value = eq_0_loop
	_ = eq_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Array_NonEmpty_intersectBy_prime_(), eq_0, gopurs_runtime.Array(xs_1)), Get_Data_Array_NonEmpty_toArray())
}

func Call_Data_Array_NonEmpty_intersect_prime_(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_intersectBy_prime_(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_NonEmpty_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_intersectBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_Array_NonEmpty_intercalate(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return Call_Data_Semigroup_Foldable_intercalate(Rebox_Data_Array_NonEmpty_3016437835_4151366573(Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))), dictSemigroup_0)
}

func Call_Data_Array_NonEmpty_insertAt(i_0_loop int64, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArrayF(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Array_insertAt(), gopurs_runtime.Int(i_0), x_1), Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
	_ = dictFoldable1_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), Call_Data_Array_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))))
}

func Call_Data_Array_NonEmpty_fromArray(xs_0_loop []gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var __t0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
	{
		if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (int64(0)) {
			__t0 = Rebox_Data_Array_NonEmpty_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Array(xs_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Rebox_Data_Array_NonEmpty_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_2807397954_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Array_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_fromArray(), Call_Data_Array_fromFoldable(dictFoldable_0))
}

func Call_Data_Array_NonEmpty_foldMap1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return Call_Data_Semigroup_Foldable_foldMap1DefaultL(Rebox_Data_Array_NonEmpty_3016437835_4151366573(Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()), dictSemigroup_0)
}

func Call_Data_Array_NonEmpty_fold1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return Call_Data_Semigroup_Foldable_fold1(Rebox_Data_Array_NonEmpty_3016437835_4151366573(Rebox_Data_Array_NonEmpty_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))), dictSemigroup_0)
}

func Call_Data_Array_NonEmpty_difference_prime_(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): difference1_1_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope120)), (Array (TypeVar a$scope120))] (Array (TypeVar a$scope120)))
	difference1_1_0 := Call_Data_Array_difference(dictEq_0)
	_ = difference1_1_0
	return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(difference1_1_0, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(xs_2.UnsafePtr))))
	})
}

func Call_Data_Array_NonEmpty_cons_prime_(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr)))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_cons_prime___3391444928(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
cons_prime___3391444928:
	for {
		if false {
			continue cons_prime___3391444928
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		var xs_1 []gopurs_runtime.Value = xs_1_loop
		_ = xs_1
		return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr))
	}
}

func Call_Data_Array_NonEmpty_fromNonEmpty(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) []gopurs_runtime.Value {
	var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{(v_0).V0}), (v_0).V1).UnsafePtr)))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 []gopurs_runtime.Value = a_1_loop
	_ = a_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Bind_arrayBind(), gopurs_runtime.Array(a_1), b_0).UnsafePtr))
}

func Call_Data_Array_NonEmpty_appendArray(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var xs_0 []gopurs_runtime.Value = xs_0_loop
	_ = xs_0
	var ys_1 []gopurs_runtime.Value = ys_1_loop
	_ = ys_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)))).UnsafePtr))
}

func Call_Data_Array_NonEmpty_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Array_alterAt(), gopurs_runtime.Int(i_0), f_1), Get_Data_Array_NonEmpty_toArray())
}

func Call_Data_Array_NonEmpty_adaptMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Get_Data_Array_NonEmpty_toArray()))
	}))
}

func Call_Data_Array_NonEmpty_adaptAny(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Get_Data_Array_NonEmpty_toArray())
}

func Call_Data_Array_NonEmpty_all(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_all(), p_0))
}

func Call_Data_Array_NonEmpty_any(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_any(), p_0))
}

func Call_Data_Array_NonEmpty_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Call_Data_Array_go__delete(dictEq_0), x_1))
}

func Call_Data_Array_NonEmpty_deleteAt(i_0_loop int64) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_deleteAt(), gopurs_runtime.Int(i_0)))
}

func Call_Data_Array_NonEmpty_deleteBy(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply2(Get_Data_Array_deleteBy(), f_0, x_1))
}

func Call_Data_Array_NonEmpty_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): difference1_1_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a$scope181)), (Array (TypeVar a$scope181))] (Array (TypeVar a$scope181)))
	difference1_1_0 := Call_Data_Array_difference(dictEq_0)
	_ = difference1_1_0
	return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(difference1_1_0, gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(xs_2.UnsafePtr)))))
	})
}

func Call_Data_Array_NonEmpty_drop(i_0_loop int64) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_drop(), gopurs_runtime.Int(i_0)))
}

func Call_Data_Array_NonEmpty_dropEnd(i_0_loop int64) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_dropEnd(), gopurs_runtime.Int(i_0)))
}

func Call_Data_Array_NonEmpty_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_dropWhile(), f_0))
}

func Call_Data_Array_NonEmpty_elem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply2(Get_Data_Array_elem(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1))
}

func Call_Data_Array_NonEmpty_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_adaptAny(Call_Data_Array_elemIndex(dictEq_0, x_1))
}

func Call_Data_Array_NonEmpty_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_adaptAny(Call_Data_Array_elemLastIndex(dictEq_0, x_1))
}

func Call_Data_Array_NonEmpty_filter(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_filter(), f_0))
}

func Call_Data_Array_NonEmpty_filterA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): filterA1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope198)] (TypeApp (TypeVar f$scope199) [Boolean])), (Array (TypeVar a$scope198))] (TypeApp (TypeVar f$scope199) [(Array (TypeVar a$scope198))]))
	filterA1_1_0 := Call_Data_Array_filterA(dictApplicative_0)
	_ = filterA1_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(filterA1_1_0, f_2))
	})
}

func Call_Data_Array_NonEmpty_find(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_find(), p_0))
}

func Call_Data_Array_NonEmpty_findIndex(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_findIndex(), p_0))
}

func Call_Data_Array_NonEmpty_findLastIndex(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_findLastIndex(), x_0))
}

func Call_Data_Array_NonEmpty_findMap(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var p_0 gopurs_runtime.Value = p_0_loop
	_ = p_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_findMap(), p_0))
}

func Call_Data_Array_NonEmpty_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var acc_2 gopurs_runtime.Value = acc_2_loop
	_ = acc_2
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply3(Get_Data_Array_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_1, acc_2))
}

func Call_Data_Array_NonEmpty_foldRecM(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): foldRecM1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope221), (TypeVar a$scope220)] (TypeApp (TypeVar m$scope219) [(TypeVar b$scope221)])), (TypeVar b$scope221), (Array (TypeVar a$scope220))] (TypeApp (TypeVar m$scope219) [(TypeVar b$scope221)]))
	foldRecM1_1_0 := Call_Data_Array_foldRecM(dictMonadRec_0)
	_ = foldRecM1_1_0
	return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply2(foldRecM1_1_0, f_2, acc_3))
	})
}

func Call_Data_Array_NonEmpty_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_adaptAny(Call_Data_Array_mapMaybe(f_0))
}

func Call_Data_Array_NonEmpty_notElem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply2(Get_Data_Array_notElem(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1))
}

func Call_Data_Array_NonEmpty_partition(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_partition(), f_0))
}

func Call_Data_Array_NonEmpty_slice(start_0_loop int64, end_1_loop int64) gopurs_runtime.Value {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply2(Get_Data_Array_slice(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)))
}

func Call_Data_Array_NonEmpty_span(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_span(), f_0))
}

func Call_Data_Array_NonEmpty_take(i_0_loop int64) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_take(), gopurs_runtime.Int(i_0)))
}

func Call_Data_Array_NonEmpty_takeEnd(i_0_loop int64) gopurs_runtime.Value {
	var i_0 int64 = i_0_loop
	_ = i_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_takeEnd(), gopurs_runtime.Int(i_0)))
}

func Call_Data_Array_NonEmpty_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_adaptAny(gopurs_runtime.Apply(Get_Data_Array_takeWhile(), f_0))
}

func Call_Data_Array_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)}), Get_Data_Array_NonEmpty_toArray())
}

func Call_Data_Array_NonEmpty_unsafeAdapt(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_cons(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_Array_cons(), x_0))
}

func Call_Data_Array_NonEmpty_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(Call_Data_Array_group(dictEq_0))
}

func Call_Data_Array_NonEmpty_groupAllBy(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var op_0 gopurs_runtime.Value = op_0_loop
	_ = op_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(Call_Data_Array_groupAllBy(op_0))
}

func Call_Data_Array_NonEmpty_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(Call_Data_Array_groupAllBy(Call_Data_Ord_compare(dictOrd_0)))
}

func Call_Data_Array_NonEmpty_groupBy(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var op_0 gopurs_runtime.Value = op_0_loop
	_ = op_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_Array_groupBy(), op_0))
}

func Call_Data_Array_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Call_Data_Array_insert(dictOrd_0), x_1))
}

func Call_Data_Array_NonEmpty_insertBy(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply2(Get_Data_Array_insertBy(), f_0, x_1))
}

func Call_Data_Array_NonEmpty_intersperse(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_Array_intersperse(), x_0))
}

func Call_Data_Array_NonEmpty_mapWithIndex(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), f_0))
}

func Call_Data_Array_NonEmpty_modifyAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): modifyAtIndices1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar t$scope276) [Int]), (Func [(TypeVar a$scope277)] (TypeVar a$scope277)), (Array (TypeVar a$scope277))] (Array (TypeVar a$scope277)))
	modifyAtIndices1_1_0 := Call_Data_Array_modifyAtIndices(dictFoldable_0)
	_ = modifyAtIndices1_1_0
	return gopurs_runtime.Func2(func(is_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply2(modifyAtIndices1_1_0, is_2, f_3))
	})
}

func Call_Data_Array_NonEmpty_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Array_nub(dictOrd_0), Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_nubBy(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_Array_nubBy(), f_0))
}

func Call_Data_Array_NonEmpty_nubByEq(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_Array_nubByEq(), f_0))
}

func Call_Data_Array_NonEmpty_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Array_nubEq(dictEq_0), Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_scanl(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply2(Get_Data_Array_scanl(), f_0, x_1))
}

func Call_Data_Array_NonEmpty_scanr(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply2(Get_Data_Array_scanr(), f_0, x_1))
}

func Call_Data_Array_NonEmpty_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Array_NonEmpty_unsafeFromArray(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Array_sort(dictOrd_0), Get_Data_Array_NonEmpty_toArray()))
}

func Call_Data_Array_NonEmpty_sortBy(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(Get_Data_Array_sortBy(), f_0))
}

func Call_Data_Array_NonEmpty_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return Call_Data_Array_NonEmpty_unsafeAdapt(Call_Data_Array_sortWith(dictOrd_0, f_1))
}

func Call_Data_Array_NonEmpty_updateAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
	_ = dictFoldable_0
	// TAST (Let): updateAtIndices1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar t$scope307) [(ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a$scope308)])]), (Array (TypeVar a$scope308))] (Array (TypeVar a$scope308)))
	updateAtIndices1_1_0 := Call_Data_Array_updateAtIndices(dictFoldable_0)
	_ = updateAtIndices1_1_0
	return gopurs_runtime.Func(func(pairs_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Array_NonEmpty_unsafeAdapt(gopurs_runtime.Apply(updateAtIndices1_1_0, pairs_2))
	})
}

func Call_Data_Array_NonEmpty_unsafeIndex(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_unsafeIndex(), gopurs_runtime.Value{}), Get_Data_Array_NonEmpty_toArray())
}

func Call_Data_Array_NonEmpty_toUnfoldable1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
	_ = dictUnfoldable1_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	// TAST (Let): len_2_0 shape=App(Var) bindingType=Int
	len_2_0 := gopurs_runtime.Apply(Get_Data_Array_NonEmpty_length(), gopurs_runtime.Array(xs_1)).IntVal
	_ = len_2_0
	return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (i_3.IntVal) < (gopurs_runtime.IntSub(len_2_0, int64(1))) {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(gopurs_runtime.IntAdd(i_3.IntVal, int64(1))), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}
	end_branch_1:
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply4(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Array_unsafeIndex(), gopurs_runtime.Value{}), Get_Data_Array_NonEmpty_toArray(), gopurs_runtime.Array(xs_1), gopurs_runtime.Int(i_3.IntVal)), __t1}))}
	}), gopurs_runtime.Int(int64(0)))
}

func Rebox_Data_Array_NonEmpty_2807397954_3094389156(in *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Array(in.V0)
	return out
}

func Rebox_Data_Array_NonEmpty_3016437835_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_3094389156_2807397954(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]{}
	out.V0 = *((*[]gopurs_runtime.Value)(in.V0.UnsafePtr))
	return out
}

func Rebox_Data_Array_NonEmpty_4151366573_3016437835(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}
