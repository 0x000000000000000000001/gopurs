package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_lessThanOrEq gopurs_runtime.Value
var once_Control_Monad_Gen_lessThanOrEq sync.Once

func Get_Control_Monad_Gen_lessThanOrEq() gopurs_runtime.Value {
	once_Control_Monad_Gen_lessThanOrEq.Do(func() {
		cache_Control_Monad_Gen_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Control_Monad_Gen_lessThanOrEq
}

var cache_Control_Monad_Gen_sub gopurs_runtime.Value
var once_Control_Monad_Gen_sub sync.Once

func Get_Control_Monad_Gen_sub() gopurs_runtime.Value {
	once_Control_Monad_Gen_sub.Do(func() {
		cache_Control_Monad_Gen_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Control_Monad_Gen_sub
}

var cache_Control_Monad_Gen_compose gopurs_runtime.Value
var once_Control_Monad_Gen_compose sync.Once

func Get_Control_Monad_Gen_compose() gopurs_runtime.Value {
	once_Control_Monad_Gen_compose.Do(func() {
		cache_Control_Monad_Gen_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Gen_compose
}

var cache_Control_Monad_Gen_compose1 gopurs_runtime.Value
var once_Control_Monad_Gen_compose1 sync.Once

func Get_Control_Monad_Gen_compose1() gopurs_runtime.Value {
	once_Control_Monad_Gen_compose1.Do(func() {
		cache_Control_Monad_Gen_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Gen_compose1
}

var cache_Control_Monad_Gen_un gopurs_runtime.Value
var once_Control_Monad_Gen_un sync.Once

func Get_Control_Monad_Gen_un() gopurs_runtime.Value {
	once_Control_Monad_Gen_un.Do(func() {
		cache_Control_Monad_Gen_un = gopurs_runtime.Apply(Get_Data_Newtype_un(), gopurs_runtime.Value{})
	})
	return cache_Control_Monad_Gen_un
}

var cache_Control_Monad_Gen_un__3878100693 gopurs_runtime.Value
var once_Control_Monad_Gen_un__3878100693 sync.Once

func Get_Control_Monad_Gen_un__3878100693() gopurs_runtime.Value {
	once_Control_Monad_Gen_un__3878100693.Do(func() {
		cache_Control_Monad_Gen_un__3878100693 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_un__3878100693(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Monad_Gen_un__3878100693
}

var cache_Control_Monad_Gen_greaterThanOrEq gopurs_runtime.Value
var once_Control_Monad_Gen_greaterThanOrEq sync.Once

func Get_Control_Monad_Gen_greaterThanOrEq() gopurs_runtime.Value {
	once_Control_Monad_Gen_greaterThanOrEq.Do(func() {
		cache_Control_Monad_Gen_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__756116985()
	})
	return cache_Control_Monad_Gen_greaterThanOrEq
}

var cache_Control_Monad_Gen_sub1 gopurs_runtime.Value
var once_Control_Monad_Gen_sub1 sync.Once

func Get_Control_Monad_Gen_sub1() gopurs_runtime.Value {
	once_Control_Monad_Gen_sub1.Do(func() {
		cache_Control_Monad_Gen_sub1 = Get_Data_Ring_sub__179436318()
	})
	return cache_Control_Monad_Gen_sub1
}

var cache_Control_Monad_Gen_alaF gopurs_runtime.Value
var once_Control_Monad_Gen_alaF sync.Once

func Get_Control_Monad_Gen_alaF() gopurs_runtime.Value {
	once_Control_Monad_Gen_alaF.Do(func() {
		cache_Control_Monad_Gen_alaF = gopurs_runtime.Apply4(Get_Data_Newtype_alaF(), gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Value{})
	})
	return cache_Control_Monad_Gen_alaF
}

var cache_Control_Monad_Gen_monoidAdditive gopurs_runtime.Value
var once_Control_Monad_Gen_monoidAdditive sync.Once

func Get_Control_Monad_Gen_monoidAdditive() gopurs_runtime.Value {
	once_Control_Monad_Gen_monoidAdditive.Do(func() {
		cache_Control_Monad_Gen_monoidAdditive = Call_Data_Monoid_Additive_monoidAdditive(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_602713622_2826095630(Rebox_Control_Monad_Gen_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber()))))})
	})
	return cache_Control_Monad_Gen_monoidAdditive
}

var cache_Control_Monad_Gen_Cons gopurs_runtime.Value
var once_Control_Monad_Gen_Cons sync.Once

func Get_Control_Monad_Gen_Cons() gopurs_runtime.Value {
	once_Control_Monad_Gen_Cons.Do(func() {
		cache_Control_Monad_Gen_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Control_Monad_Gen_Cons
}

var cache_Control_Monad_Gen_Nil gopurs_runtime.Value
var once_Control_Monad_Gen_Nil sync.Once

func Get_Control_Monad_Gen_Nil() gopurs_runtime.Value {
	once_Control_Monad_Gen_Nil.Do(func() {
		cache_Control_Monad_Gen_Nil = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Control_Monad_Gen_Nil
}

var cache_Control_Monad_Gen_FreqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_FreqSemigroup sync.Once

func Get_Control_Monad_Gen_FreqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_FreqSemigroup.Do(func() {
		cache_Control_Monad_Gen_FreqSemigroup = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_FreqSemigroup(x_0_box)
		})
	})
	return cache_Control_Monad_Gen_FreqSemigroup
}

var cache_Control_Monad_Gen_unfoldable gopurs_runtime.Value
var once_Control_Monad_Gen_unfoldable sync.Once

func Get_Control_Monad_Gen_unfoldable() gopurs_runtime.Value {
	once_Control_Monad_Gen_unfoldable.Do(func() {
		cache_Control_Monad_Gen_unfoldable = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_unfoldable(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Control_Monad_Gen_unfoldable
}

var cache_Control_Monad_Gen_semigroupFreqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_semigroupFreqSemigroup sync.Once

func Get_Control_Monad_Gen_semigroupFreqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_semigroupFreqSemigroup.Do(func() {
		cache_Control_Monad_Gen_semigroupFreqSemigroup = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, pos_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v2_3_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Maybe","Maybe"] [Number]), (TypeVar a$scope22)])
			v2_3_0 := Rebox_Control_Monad_Gen_138441832_1087394609(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, gopurs_runtime.Float(pos_2.FloatVal()))))
			_ = v2_3_0
			var __t2 gopurs_runtime.Value
			{
				var __t_tag_1 *Constructor_Data_Maybe_Just[float64] = (v2_3_0).V0
				_ = __t_tag_1
				if __t_tag_1 != nil {
					__t2 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Float(((v2_3_0).V0).V0))
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1087394609_138441832(v2_3_0))}
			}
		end_branch_2:
			return __t2
		})}))}
	})
	return cache_Control_Monad_Gen_semigroupFreqSemigroup
}

var cache_Control_Monad_Gen_getFreqVal gopurs_runtime.Value
var once_Control_Monad_Gen_getFreqVal sync.Once

func Get_Control_Monad_Gen_getFreqVal() gopurs_runtime.Value {
	once_Control_Monad_Gen_getFreqVal.Do(func() {
		cache_Control_Monad_Gen_getFreqVal = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_getFreqVal(v_0_box)
		})
	})
	return cache_Control_Monad_Gen_getFreqVal
}

var cache_Control_Monad_Gen_fromIndex gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex sync.Once

func Get_Control_Monad_Gen_fromIndex() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex.Do(func() {
		cache_Control_Monad_Gen_fromIndex = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
		})
	})
	return cache_Control_Monad_Gen_fromIndex
}

var cache_Control_Monad_Gen_fromIndex__10227597 gopurs_runtime.Value
var once_Control_Monad_Gen_fromIndex__10227597 sync.Once

func Get_Control_Monad_Gen_fromIndex__10227597() gopurs_runtime.Value {
	once_Control_Monad_Gen_fromIndex__10227597.Do(func() {
		cache_Control_Monad_Gen_fromIndex__10227597 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_fromIndex__10227597(Rebox_Control_Monad_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box)))
		})
	})
	return cache_Control_Monad_Gen_fromIndex__10227597
}

var cache_Control_Monad_Gen_oneOf gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf sync.Once

func Get_Control_Monad_Gen_oneOf() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf.Do(func() {
		cache_Control_Monad_Gen_oneOf = gopurs_runtime.Func2(func(dictMonadGen_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_1_box))
		})
	})
	return cache_Control_Monad_Gen_oneOf
}

var cache_Control_Monad_Gen_oneOf__1997668141 gopurs_runtime.Value
var once_Control_Monad_Gen_oneOf__1997668141 sync.Once

func Get_Control_Monad_Gen_oneOf__1997668141() gopurs_runtime.Value {
	once_Control_Monad_Gen_oneOf__1997668141.Do(func() {
		cache_Control_Monad_Gen_oneOf__1997668141 = gopurs_runtime.Func2(func(dictMonadGen_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_oneOf__1997668141(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_1_box))
		})
	})
	return cache_Control_Monad_Gen_oneOf__1997668141
}

var cache_Control_Monad_Gen_freqSemigroup gopurs_runtime.Value
var once_Control_Monad_Gen_freqSemigroup sync.Once

func Get_Control_Monad_Gen_freqSemigroup() gopurs_runtime.Value {
	once_Control_Monad_Gen_freqSemigroup.Do(func() {
		cache_Control_Monad_Gen_freqSemigroup = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_freqSemigroup(Rebox_Control_Monad_Gen_138441832_3854293424(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
		})
	})
	return cache_Control_Monad_Gen_freqSemigroup
}

var cache_Control_Monad_Gen_frequency gopurs_runtime.Value
var once_Control_Monad_Gen_frequency sync.Once

func Get_Control_Monad_Gen_frequency() gopurs_runtime.Value {
	once_Control_Monad_Gen_frequency.Do(func() {
		cache_Control_Monad_Gen_frequency = gopurs_runtime.Func2(func(dictMonadGen_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_frequency(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_1_box))
		})
	})
	return cache_Control_Monad_Gen_frequency
}

var cache_Control_Monad_Gen_filtered gopurs_runtime.Value
var once_Control_Monad_Gen_filtered sync.Once

func Get_Control_Monad_Gen_filtered() gopurs_runtime.Value {
	once_Control_Monad_Gen_filtered.Do(func() {
		cache_Control_Monad_Gen_filtered = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_filtered(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Control_Monad_Gen_filtered
}

var cache_Control_Monad_Gen_suchThat gopurs_runtime.Value
var once_Control_Monad_Gen_suchThat sync.Once

func Get_Control_Monad_Gen_suchThat() gopurs_runtime.Value {
	once_Control_Monad_Gen_suchThat.Do(func() {
		cache_Control_Monad_Gen_suchThat = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_suchThat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Control_Monad_Gen_suchThat
}

var cache_Control_Monad_Gen_elements gopurs_runtime.Value
var once_Control_Monad_Gen_elements sync.Once

func Get_Control_Monad_Gen_elements() gopurs_runtime.Value {
	once_Control_Monad_Gen_elements.Do(func() {
		cache_Control_Monad_Gen_elements = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_elements(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
		})
	})
	return cache_Control_Monad_Gen_elements
}

var cache_Control_Monad_Gen_choose gopurs_runtime.Value
var once_Control_Monad_Gen_choose sync.Once

func Get_Control_Monad_Gen_choose() gopurs_runtime.Value {
	once_Control_Monad_Gen_choose.Do(func() {
		cache_Control_Monad_Gen_choose = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_choose(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
		})
	})
	return cache_Control_Monad_Gen_choose
}

type Constructor_Control_Monad_Gen_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Control_Monad_Gen_Cons[T_a]
}

type Constructor_Control_Monad_Gen_Nil[T_a any] struct {
	Rc uint32
}

func Call_Control_Monad_Gen_un__3878100693(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
un__3878100693:
	for {
		if false {
			continue un__3878100693
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), __eta_norm_0_1)
	}
}

func Call_Control_Monad_Gen_FreqSemigroup(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Monad_Gen_unfoldable(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
	_ = dictMonadGen_1
	// TAST (Let): Monad0_2_0 shape=App(Other) bindingType=Any
	Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
	_ = Monad0_2_0
	// TAST (Let): pure_3_1 shape=App(Var) bindingType=(Func [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Control","Monad","Gen","LL"] [(TypeVar a$scope5)]), Int]), (ADT ["Control","Monad","Gen","LL"] [(TypeVar a$scope5)])])] (TypeApp (TypeVar m$scope3) [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Control","Monad","Gen","LL"] [(TypeVar a$scope5)]), Int]), (ADT ["Control","Monad","Gen","LL"] [(TypeVar a$scope5)])])]))
	pure_3_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_3_1
	return gopurs_runtime.Func2(func(dictUnfoldable_4 gopurs_runtime.Value, gen_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_4, "unfoldr"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if v_6.Type == 9 && v_6.IntVal == 759514854 && v_6.UnsafePtr == nil {
					__t2 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					goto end_branch_2
				} else {

				}
			}
			{
				if v_6.Type == 9 && v_6.IntVal == 759514854 && v_6.UnsafePtr != nil {
					__t2 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}}))}, true}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() struct {
					V0 gopurs_runtime.Value
					V1 bool
				} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_3081497822_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]] {
					panic("Failed pattern match")
				}()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{V0: gopurs_runtime.Value{}, V1: false} }()
			}
		end_branch_2:
			return func() gopurs_runtime.Value {
				_v := __t2
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})), gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 gopurs_runtime.Value
			{
				if ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1.IntVal) <= (int64(0)) {
					__t5 = gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1429920250_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)})))})
					goto end_branch_5
				} else {

				}
			}
			{
				// TAST (Let): __local_var_7_3 shape=Other bindingType=Any
				__local_var_7_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0
				_ = __local_var_7_3
				// TAST (Let): __local_var_8_4 shape=Other bindingType=Any
				__local_var_8_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1
				_ = __local_var_8_4
				__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gen_5, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_3474068486_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]{1, (&Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64]{1, (&Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]{1, x_9, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](__local_var_7_3)}), gopurs_runtime.IntSub(__local_var_8_4.IntVal, int64(1))})})))})
				}))
			}
		end_branch_5:
			return __t5
		})), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}))}))))
	})
}

func Call_Control_Monad_Gen_getFreqVal(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), v_0)
}

func Call_Control_Monad_Gen_fromIndex(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
	_ = dictFoldable1_0
	// TAST (Let): foldMap1_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope31)] (TypeVar a$scope31)), (TypeApp (TypeVar f$scope30) [(TypeVar a$scope31)])] (TypeVar a$scope31))
	foldMap1_1_0 := gopurs_runtime.Apply(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Last_semigroupLast()))})
	_ = foldMap1_1_0
	return gopurs_runtime.Func2(func(i_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var Call_local_Control_Monad_Gen_go__go_4_1_0 func(int64, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
		_ = Call_local_Control_Monad_Gen_go__go_4_1_0
		var go__go_4_1_0 gopurs_runtime.Value
		_ = go__go_4_1_0
		Call_local_Control_Monad_Gen_go__go_4_1_0 = func(v_5_loop int64, v1_6_loop *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
		go__go_4_1_0:
			for {
				if false {
					continue go__go_4_1_0
				}
				var v_5 int64 = v_5_loop
				_ = v_5
				var v1_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = v1_6_loop
				_ = v1_6
				var __t4 gopurs_runtime.Value
				{
					if v1_6 != nil {
						var __t3 gopurs_runtime.Value
						{
							var __t_tag_2 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = (v1_6).V1
							_ = __t_tag_2
							if __t_tag_2 == nil {
								__t3 = (v1_6).V0
								goto end_branch_3
							} else {

							}
						}
						{
							if (v_5) <= (int64(0)) {
								__t3 = (v1_6).V0
								goto end_branch_3
							} else {

							}
						}
						{
							v_5_loop = gopurs_runtime.IntSub(v_5, int64(1))
							v1_6_loop = (v1_6).V1
							continue go__go_4_1_0
							__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
						}
					end_branch_3:
						__t4 = __t3
						goto end_branch_4
					} else {

					}
				}
				{
					if v1_6 == nil {
						__t4 = gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(foldMap1_1_0, Get_Data_Semigroup_Last_Last(), xs_3))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			}
		}
		go__go_4_1_0 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_local_Control_Monad_Gen_go__go_4_1_0(v_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_6_loop_val))
			})
		})
		return Call_local_Control_Monad_Gen_go__go_4_1_0(i_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}, xs_3)))
	})
}

func Call_Control_Monad_Gen_fromIndex__10227597(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
fromIndex__10227597:
	for {
		if false {
			continue fromIndex__10227597
		}
		var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] = dictFoldable1_0_loop
		_ = dictFoldable1_0
		// TAST (Let): foldMap1__2504427995_1_0 shape=App(Other) bindingType=(Func [(Func [(ADT ["Prim","Array"] [Char])] (ADT ["Prim","Array"] [Char])), (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] [Char])])] (ADT ["Prim","Array"] [Char]))
		foldMap1__2504427995_1_0 := gopurs_runtime.Apply(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Last_semigroupLast()))})
		_ = foldMap1__2504427995_1_0
		return gopurs_runtime.Func2(func(i_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var Call_local_Control_Monad_Gen_go__3236717243_4_1_1 func(int64, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
			_ = Call_local_Control_Monad_Gen_go__3236717243_4_1_1
			var go__3236717243_4_1_1 gopurs_runtime.Value
			_ = go__3236717243_4_1_1
			var Call_local_Control_Monad_Gen_go__go_4_2_2 func(int64, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
			_ = Call_local_Control_Monad_Gen_go__go_4_2_2
			var go__go_4_2_2 gopurs_runtime.Value
			_ = go__go_4_2_2
			Call_local_Control_Monad_Gen_go__3236717243_4_1_1 = func(v_5_loop int64, v1_6_loop *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
			go__3236717243_4_1_1:
				for {
					if false {
						continue go__3236717243_4_1_1
					}
					var v_5 int64 = v_5_loop
					_ = v_5
					var v1_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = v1_6_loop
					_ = v1_6
					var __t5 gopurs_runtime.Value
					{
						if v1_6 != nil {
							var __t4 gopurs_runtime.Value
							{
								var __t_tag_3 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = (v1_6).V1
								_ = __t_tag_3
								if __t_tag_3 == nil {
									__t4 = (v1_6).V0
									goto end_branch_4
								} else {

								}
							}
							{
								if (v_5) <= (int64(0)) {
									__t4 = (v1_6).V0
									goto end_branch_4
								} else {

								}
							}
							{
								v_5_loop = gopurs_runtime.IntSub(v_5, int64(1))
								v1_6_loop = (v1_6).V1
								continue go__3236717243_4_1_1
								__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
							}
						end_branch_4:
							__t5 = __t4
							goto end_branch_5
						} else {

						}
					}
					{
						if v1_6 == nil {
							__t5 = gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(foldMap1__2504427995_1_0, Get_Data_Semigroup_Last_Last(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))}))
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_5:
					return __t5
				}
			}
			go__3236717243_4_1_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Control_Monad_Gen_go__3236717243_4_1_1(v_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_6_loop_val))
				})
			})
			Call_local_Control_Monad_Gen_go__go_4_2_2 = func(v_5_loop int64, v1_6_loop *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
			go__go_4_2_2:
				for {
					if false {
						continue go__go_4_2_2
					}
					var v_5 int64 = v_5_loop
					_ = v_5
					var v1_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = v1_6_loop
					_ = v1_6
					var __t8 gopurs_runtime.Value
					{
						if v1_6 != nil {
							var __t7 gopurs_runtime.Value
							{
								var __t_tag_6 *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value] = (v1_6).V1
								_ = __t_tag_6
								if __t_tag_6 == nil {
									__t7 = (v1_6).V0
									goto end_branch_7
								} else {

								}
							}
							{
								if (v_5) <= (int64(0)) {
									__t7 = (v1_6).V0
									goto end_branch_7
								} else {

								}
							}
							{
								__t7 = Call_local_Control_Monad_Gen_go__3236717243_4_1_1(gopurs_runtime.IntSub(v_5, int64(1)), (v1_6).V1)
							}
						end_branch_7:
							__t8 = __t7
							goto end_branch_8
						} else {

						}
					}
					{
						if v1_6 == nil {
							__t8 = gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(foldMap1__2504427995_1_0, Get_Data_Semigroup_Last_Last(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))}))
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_8:
					return __t8
				}
			}
			go__go_4_2_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_local_Control_Monad_Gen_go__go_4_2_2(v_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](v1_6_loop_val))
				})
			})
			return Call_local_Control_Monad_Gen_go__3236717243_4_1_1(i_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3))})))
		})
	}
}

func Call_Control_Monad_Gen_oneOf(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value], dictFoldable1_1_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	var dictFoldable1_1 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_1_loop
	_ = dictFoldable1_1
	// TAST (Let): length_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope40) [(TypeApp (TypeVar m$scope39) [(TypeVar a$scope41)])])] Int)
	length_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictFoldable1_1.V0, gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func2(func(c_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), c_2.IntVal))
	}), gopurs_runtime.Int(int64(0)))
	_ = length_2_0
	// TAST (Let): fromIndex1_3_1 shape=App(Var) bindingType=(Func [Int, (TypeApp (TypeVar f$scope40) [(TypeApp (TypeVar m$scope39) [(TypeVar a$scope41)])])] (TypeApp (TypeVar m$scope39) [(TypeVar a$scope41)]))
	fromIndex1_3_1 := Call_Control_Monad_Gen_fromIndex(dictFoldable1_1)
	_ = fromIndex1_3_1
	return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(length_2_0, xs_4).IntVal, int64(1)))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(fromIndex1_3_1, gopurs_runtime.Int(n_5.IntVal), xs_4)
		}))
	})
}

func Call_Control_Monad_Gen_oneOf__1997668141(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
oneOf__1997668141:
	for {
		if false {
			continue oneOf__1997668141
		}
		var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
		_ = dictMonadGen_0
		var __eta_norm_0_1 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): fromIndex1__244441942_2_0 shape=App(Var) bindingType=(Func [Int, (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Prim","Array"] [Char])])] (ADT ["Prim","Array"] [Char]))
		fromIndex1__244441942_2_0 := Call_Control_Monad_Gen_fromIndex__10227597(Rebox_Control_Monad_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Char_Gen_foldable1NonEmpty())))
		_ = fromIndex1__244441942_2_0
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Rebox_Control_Monad_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Char_Gen_foldable1NonEmpty())).V0, gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func2(func(c_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), c_3.IntVal))
		}), gopurs_runtime.Int(gopurs_runtime.Int(int64(0)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__eta_norm_0_1)}).IntVal, int64(1)))), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(fromIndex1__244441942_2_0, gopurs_runtime.Int(n_3.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__eta_norm_0_1)})
		}))
	}
}

func Call_Control_Monad_Gen_freqSemigroup(v_0_loop *Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
	__local_var_1_0 := (v_0).V0
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
	__local_var_2_1 := (v_0).V1
	_ = __local_var_2_1
	return gopurs_runtime.Func(func(pos_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t2 struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}
		{
			if (pos_3.FloatVal()) >= (__local_var_1_0) {
				__t2 = struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Float((pos_3.FloatVal()) - (__local_var_1_0)), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}(), __local_var_2_1}
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, __local_var_2_1}
		}
	end_branch_2:
		return func() gopurs_runtime.Value {
			_v := __t2
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})
}

func Call_Control_Monad_Gen_frequency(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value], dictFoldable1_1_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	var dictFoldable1_1 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_1_loop
	_ = dictFoldable1_1
	// TAST (Let): foldMap_2_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Tuple","Tuple"] [Number, (TypeApp (TypeVar m$scope54) [(TypeVar a$scope56)])])] Number), (TypeApp (TypeVar f$scope55) [(ADT ["Data","Tuple","Tuple"] [Number, (TypeApp (TypeVar m$scope54) [(TypeVar a$scope56)])])])] Number)
	foldMap_2_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_1.V0, gopurs_runtime.Value{}))), Get_Control_Monad_Gen_monoidAdditive())
	_ = foldMap_2_0
	// TAST (Let): foldMap1_3_1 shape=App(Other) bindingType=(Func [(Func [(ADT ["Data","Tuple","Tuple"] [Number, (TypeApp (TypeVar m$scope54) [(TypeVar a$scope56)])]), Number] (ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Maybe","Maybe"] [Number]), (TypeApp (TypeVar m$scope54) [(TypeVar a$scope56)])])), (TypeApp (TypeVar f$scope55) [(ADT ["Data","Tuple","Tuple"] [Number, (TypeApp (TypeVar m$scope54) [(TypeVar a$scope56)])])]), Number] (ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Maybe","Maybe"] [Number]), (TypeApp (TypeVar m$scope54) [(TypeVar a$scope56)])]))
	foldMap1_3_1 := gopurs_runtime.Apply(dictFoldable1_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Control_Monad_Gen_semigroupFreqSemigroup()))})
	_ = foldMap1_3_1
	return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(gopurs_runtime.Apply3(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), foldMap_2_0, Get_Data_Tuple_fst(), xs_4).FloatVal())), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply2(foldMap1_3_1, Get_Control_Monad_Gen_freqSemigroup(), xs_4)))
	})
}

func Call_Control_Monad_Gen_filtered(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
	_ = dictMonadGen_1
	// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
	__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_2_0
	return gopurs_runtime.Func(func(gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "map"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if a_5.Type == 9 && a_5.IntVal == 930809136 && a_5.UnsafePtr == nil {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))}
						goto end_branch_1
					} else {

					}
				}
				{
					if a_5.Type == 9 && a_5.IntVal == 930809136 && a_5.UnsafePtr != nil {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_5.UnsafePtr).V0}))}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_1:
				return __t1
			}), gen_3)
		}), Get_Data_Unit_unit())
	})
}

func Call_Control_Monad_Gen_suchThat(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
	_ = dictMonadGen_1
	// TAST (Let): filtered2_2_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope69)])])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope69)]))
	filtered2_2_0 := Call_Control_Monad_Gen_filtered(dictMonadRec_0, dictMonadGen_1)
	_ = filtered2_2_0
	// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
	__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_3_1
	return gopurs_runtime.Func2(func(gen_4 gopurs_runtime.Value, pred_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(filtered2_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if (gopurs_runtime.Apply(pred_5, a_6).IntVal) != (0) {
					__t2 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{a_6, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
			}
		end_branch_2:
			return __t2
		}), gen_4))
	})
}

func Call_Control_Monad_Gen_elements(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): pure_2_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope77)] (TypeApp (TypeVar m$scope75) [(TypeVar a$scope77)]))
	pure_2_1 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_2_1
	return gopurs_runtime.Func(func(dictFoldable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): length_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope76) [(TypeVar a$scope77)])] Int)
		length_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_3, "Foldable0"), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func2(func(c_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(int64(1), c_4.IntVal))
		}), gopurs_runtime.Int(int64(0)))
		_ = length_4_2
		// TAST (Let): fromIndex1_5_3 shape=App(Var) bindingType=(Func [Int, (TypeApp (TypeVar f$scope76) [(TypeVar a$scope77)])] (TypeVar a$scope77))
		fromIndex1_5_3 := Call_Control_Monad_Gen_fromIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_3))
		_ = fromIndex1_5_3
		return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(length_4_2, xs_6).IntVal, int64(1)))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(fromIndex1_5_3, gopurs_runtime.Int(n_7.IntVal), xs_6))
			}))
		})
	})
}

func Call_Control_Monad_Gen_choose(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): chooseBool_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar m$scope80) [Boolean])
	chooseBool_1_0 := Call_Control_Monad_Gen_Class_chooseBool(gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)})
	_ = chooseBool_1_0
	return gopurs_runtime.Func2(func(genA_2 gopurs_runtime.Value, genB_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), chooseBool_1_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if (v_4.IntVal) != (0) {
					__t1 = genA_2
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = genB_3
			}
		end_branch_1:
			return __t1
		}))
	})
}

func Rebox_Control_Monad_Gen_1080505919_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Control_Monad_Gen_1087394609_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_3240988860_3094389156(in.V0))}
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Gen_138441832_1087394609(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[float64], gopurs_runtime.Value]{}
	out.V0 = Rebox_Control_Monad_Gen_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0))
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Gen_138441832_3854293424(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[float64, gopurs_runtime.Value]{}
	out.V0 = in.V0.FloatVal()
	out.V1 = in.V1
	return out
}

func Rebox_Control_Monad_Gen_1429920250_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_Gen_2826095630_602713622(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[float64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.FloatVal()
	out.V3 = in.V3.FloatVal()
	return out
}

func Rebox_Control_Monad_Gen_3081497822_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_1080505919_138441832(in.V0))}
	return out
}

func Rebox_Control_Monad_Gen_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Control_Monad_Gen_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Control_Monad_Gen_3474068486_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64], *Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Gen_4216845252_138441832(in.V0))}
	return out
}

func Rebox_Control_Monad_Gen_4151366573_4217626592(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_Gen_4216845252_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Control_Monad_Gen_Cons[gopurs_runtime.Value], int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(in.V0)}
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Control_Monad_Gen_602713622_2826095630(in *Constructor_Data_Semiring_Semiring[float64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Float(in.V2)
	out.V3 = gopurs_runtime.Float(in.V3)
	return out
}
