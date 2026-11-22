package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_Common_compose gopurs_runtime.Value
var once_Control_Monad_Gen_Common_compose sync.Once

func Get_Control_Monad_Gen_Common_compose() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_compose.Do(func() {
		cache_Control_Monad_Gen_Common_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Gen_Common_compose
}

var cache_Control_Monad_Gen_Common_max gopurs_runtime.Value
var once_Control_Monad_Gen_Common_max sync.Once

func Get_Control_Monad_Gen_Common_max() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_max.Do(func() {
		cache_Control_Monad_Gen_Common_max = Get_Data_Ord_max__4009478382()
	})
	return cache_Control_Monad_Gen_Common_max
}

var cache_Control_Monad_Gen_Common_sub gopurs_runtime.Value
var once_Control_Monad_Gen_Common_sub sync.Once

func Get_Control_Monad_Gen_Common_sub() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_sub.Do(func() {
		cache_Control_Monad_Gen_Common_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Control_Monad_Gen_Common_sub
}

var cache_Control_Monad_Gen_Common_lessThan gopurs_runtime.Value
var once_Control_Monad_Gen_Common_lessThan sync.Once

func Get_Control_Monad_Gen_Common_lessThan() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_lessThan.Do(func() {
		cache_Control_Monad_Gen_Common_lessThan = Get_Data_Ord_lessThan__756116985()
	})
	return cache_Control_Monad_Gen_Common_lessThan
}

var cache_Control_Monad_Gen_Common_genTuple gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genTuple sync.Once

func Get_Control_Monad_Gen_Common_genTuple() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genTuple.Do(func() {
		cache_Control_Monad_Gen_Common_genTuple = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genTuple(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), a_1_box, b_2_box)
		})
	})
	return cache_Control_Monad_Gen_Common_genTuple
}

var cache_Control_Monad_Gen_Common_genNonEmpty gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genNonEmpty sync.Once

func Get_Control_Monad_Gen_Common_genNonEmpty() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genNonEmpty.Do(func() {
		cache_Control_Monad_Gen_Common_genNonEmpty = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genNonEmpty(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Control_Monad_Gen_Common_genNonEmpty
}

var cache_Control_Monad_Gen_Common_genMaybe_prime_ gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genMaybe_prime_ sync.Once

func Get_Control_Monad_Gen_Common_genMaybe_prime_() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genMaybe_prime_.Do(func() {
		cache_Control_Monad_Gen_Common_genMaybe_prime_ = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genMaybe_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
		})
	})
	return cache_Control_Monad_Gen_Common_genMaybe_prime_
}

var cache_Control_Monad_Gen_Common_genMaybe gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genMaybe sync.Once

func Get_Control_Monad_Gen_Common_genMaybe() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genMaybe.Do(func() {
		cache_Control_Monad_Gen_Common_genMaybe = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genMaybe(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
		})
	})
	return cache_Control_Monad_Gen_Common_genMaybe
}

var cache_Control_Monad_Gen_Common_genIdentity gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genIdentity sync.Once

func Get_Control_Monad_Gen_Common_genIdentity() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genIdentity.Do(func() {
		cache_Control_Monad_Gen_Common_genIdentity = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genIdentity(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
		})
	})
	return cache_Control_Monad_Gen_Common_genIdentity
}

var cache_Control_Monad_Gen_Common_genEither_prime_ gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genEither_prime_ sync.Once

func Get_Control_Monad_Gen_Common_genEither_prime_() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genEither_prime_.Do(func() {
		cache_Control_Monad_Gen_Common_genEither_prime_ = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genEither_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
		})
	})
	return cache_Control_Monad_Gen_Common_genEither_prime_
}

var cache_Control_Monad_Gen_Common_genEither gopurs_runtime.Value
var once_Control_Monad_Gen_Common_genEither sync.Once

func Get_Control_Monad_Gen_Common_genEither() gopurs_runtime.Value {
	once_Control_Monad_Gen_Common_genEither.Do(func() {
		cache_Control_Monad_Gen_Common_genEither = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Gen_Common_genEither(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
		})
	})
	return cache_Control_Monad_Gen_Common_genEither
}

func Call_Control_Monad_Gen_Common_genTuple(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), a_1), b_2)
}

func Call_Control_Monad_Gen_Common_genNonEmpty(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
	_ = dictMonadGen_1
	// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=Any
	Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
	_ = Bind1_2_0
	// TAST (Let): unfoldable1__763098789_3_1 shape=App(Var) bindingType=(ConstrainedType (Func [(TypeApp (TypeVar m$scope11) [(TypeVar a$scope12)])] (TypeApp (TypeVar m$scope11) [(TypeApp (TypeVar f$scope13) [(TypeVar a$scope12)])])))
	unfoldable1__763098789_3_1 := Call_Control_Monad_Gen_unfoldable(dictMonadRec_0, dictMonadGen_1)
	_ = unfoldable1__763098789_3_1
	return gopurs_runtime.Func(func(dictUnfoldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): unfoldable2_5_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar m$scope11) [(TypeVar a$scope12)])] (TypeApp (TypeVar m$scope11) [(TypeApp (TypeVar f$scope13) [(TypeVar a$scope12)])]))
		unfoldable2_5_2 := gopurs_runtime.Apply(unfoldable1__763098789_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_4))})
		_ = unfoldable2_5_2
		return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gen_6), gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Ord_max__4009478382(), gopurs_runtime.Int(int64(0))), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntSub(v_7.IntVal, int64(1)))
			})), gopurs_runtime.Apply(unfoldable2_5_2, gen_6)))
		})
	})
}

func Call_Control_Monad_Gen_Common_genMaybe_prime_(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
	_ = Monad0_1_0
	return gopurs_runtime.Func2(func(bias_2 gopurs_runtime.Value, gen_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if (n_4.FloatVal()) < (bias_2.FloatVal()) {
					__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_Just(), gen_3)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
			}
		end_branch_1:
			return __t1
		}))
	})
}

func Call_Control_Monad_Gen_Common_genMaybe(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(Get_Control_Monad_Gen_Common_genMaybe_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)}, gopurs_runtime.Float(0.75))
}

func Call_Control_Monad_Gen_Common_genIdentity(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Apply(dictFunctor_0.V0, Get_Data_Identity_Identity())
}

func Call_Control_Monad_Gen_Common_genEither_prime_(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope35)])
	Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_2_1
	return gopurs_runtime.Func3(func(bias_3 gopurs_runtime.Value, genA_4 gopurs_runtime.Value, genB_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_0.V2, gopurs_runtime.Float(0.0), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if (n_6.FloatVal()) < (bias_3.FloatVal()) {
					__t2 = gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Either_Left(), genA_4)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Either_Right(), genB_5)
			}
		end_branch_2:
			return __t2
		}))
	})
}

func Call_Control_Monad_Gen_Common_genEither(dictMonadGen_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadGen_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(Get_Control_Monad_Gen_Common_genEither_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_0)}, gopurs_runtime.Float(0.5))
}
