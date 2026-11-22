package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Parallel_compose gopurs_runtime.Value
var once_Control_Parallel_compose sync.Once

func Get_Control_Parallel_compose() gopurs_runtime.Value {
	once_Control_Parallel_compose.Do(func() {
		cache_Control_Parallel_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose
}

var cache_Control_Parallel_compose1 gopurs_runtime.Value
var once_Control_Parallel_compose1 sync.Once

func Get_Control_Parallel_compose1() gopurs_runtime.Value {
	once_Control_Parallel_compose1.Do(func() {
		cache_Control_Parallel_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose1
}

var cache_Control_Parallel_compose2 gopurs_runtime.Value
var once_Control_Parallel_compose2 sync.Once

func Get_Control_Parallel_compose2() gopurs_runtime.Value {
	once_Control_Parallel_compose2.Do(func() {
		cache_Control_Parallel_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose2
}

var cache_Control_Parallel_compose3 gopurs_runtime.Value
var once_Control_Parallel_compose3 sync.Once

func Get_Control_Parallel_compose3() gopurs_runtime.Value {
	once_Control_Parallel_compose3.Do(func() {
		cache_Control_Parallel_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose3
}

var cache_Control_Parallel_identity gopurs_runtime.Value
var once_Control_Parallel_identity sync.Once

func Get_Control_Parallel_identity() gopurs_runtime.Value {
	once_Control_Parallel_identity.Do(func() {
		cache_Control_Parallel_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_identity(x_0_box)
		})
	})
	return cache_Control_Parallel_identity
}

var cache_Control_Parallel_identity1 gopurs_runtime.Value
var once_Control_Parallel_identity1 sync.Once

func Get_Control_Parallel_identity1() gopurs_runtime.Value {
	once_Control_Parallel_identity1.Do(func() {
		cache_Control_Parallel_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_identity1(x_0_box)
		})
	})
	return cache_Control_Parallel_identity1
}

var cache_Control_Parallel_compose4 gopurs_runtime.Value
var once_Control_Parallel_compose4 sync.Once

func Get_Control_Parallel_compose4() gopurs_runtime.Value {
	once_Control_Parallel_compose4.Do(func() {
		cache_Control_Parallel_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose4
}

var cache_Control_Parallel_compose5 gopurs_runtime.Value
var once_Control_Parallel_compose5 sync.Once

func Get_Control_Parallel_compose5() gopurs_runtime.Value {
	once_Control_Parallel_compose5.Do(func() {
		cache_Control_Parallel_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose5
}

var cache_Control_Parallel_compose6 gopurs_runtime.Value
var once_Control_Parallel_compose6 sync.Once

func Get_Control_Parallel_compose6() gopurs_runtime.Value {
	once_Control_Parallel_compose6.Do(func() {
		cache_Control_Parallel_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_compose6
}

var cache_Control_Parallel_parTraverse_ gopurs_runtime.Value
var once_Control_Parallel_parTraverse_ sync.Once

func Get_Control_Parallel_parTraverse_() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse_.Do(func() {
		cache_Control_Parallel_parTraverse_ = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parTraverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
		})
	})
	return cache_Control_Parallel_parTraverse_
}

var cache_Control_Parallel_parTraverse gopurs_runtime.Value
var once_Control_Parallel_parTraverse sync.Once

func Get_Control_Parallel_parTraverse() gopurs_runtime.Value {
	once_Control_Parallel_parTraverse.Do(func() {
		cache_Control_Parallel_parTraverse = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parTraverse(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
		})
	})
	return cache_Control_Parallel_parTraverse
}

var cache_Control_Parallel_parSequence_ gopurs_runtime.Value
var once_Control_Parallel_parSequence_ sync.Once

func Get_Control_Parallel_parSequence_() gopurs_runtime.Value {
	once_Control_Parallel_parSequence_.Do(func() {
		cache_Control_Parallel_parSequence_ = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parSequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
		})
	})
	return cache_Control_Parallel_parSequence_
}

var cache_Control_Parallel_parSequence gopurs_runtime.Value
var once_Control_Parallel_parSequence sync.Once

func Get_Control_Parallel_parSequence() gopurs_runtime.Value {
	once_Control_Parallel_parSequence.Do(func() {
		cache_Control_Parallel_parSequence = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parSequence(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
		})
	})
	return cache_Control_Parallel_parSequence
}

var cache_Control_Parallel_parOneOfMap gopurs_runtime.Value
var once_Control_Parallel_parOneOfMap sync.Once

func Get_Control_Parallel_parOneOfMap() gopurs_runtime.Value {
	once_Control_Parallel_parOneOfMap.Do(func() {
		cache_Control_Parallel_parOneOfMap = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parOneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
		})
	})
	return cache_Control_Parallel_parOneOfMap
}

var cache_Control_Parallel_parOneOf gopurs_runtime.Value
var once_Control_Parallel_parOneOf sync.Once

func Get_Control_Parallel_parOneOf() gopurs_runtime.Value {
	once_Control_Parallel_parOneOf.Do(func() {
		cache_Control_Parallel_parOneOf = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parOneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box))
		})
	})
	return cache_Control_Parallel_parOneOf
}

var cache_Control_Parallel_parApply gopurs_runtime.Value
var once_Control_Parallel_parApply sync.Once

func Get_Control_Parallel_parApply() gopurs_runtime.Value {
	once_Control_Parallel_parApply.Do(func() {
		cache_Control_Parallel_parApply = gopurs_runtime.Func3(func(dictParallel_0_box gopurs_runtime.Value, mf_1_box gopurs_runtime.Value, ma_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_parApply(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0_box), mf_1_box, ma_2_box)
		})
	})
	return cache_Control_Parallel_parApply
}

func Call_Control_Parallel_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Parallel_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Parallel_parTraverse_(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [Unit])] (TypeApp (TypeVar m$scope6) [Unit]))
	sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
	_ = sequential_1_0
	// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope6) [(TypeVar b$scope9)])] (TypeApp Any [(TypeVar b$scope9)]))
	parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
	_ = parallel_2_1
	return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse___84386644_4_2 shape=App(Var) bindingType=(ConstrainedType (Func [(Func [(TypeVar a$scope8)] (TypeApp Any [(TypeVar b$scope9)])), (TypeApp (TypeVar t$scope7) [(TypeVar a$scope8)])] (TypeApp Any [Unit])))
		traverse___84386644_4_2 := Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))
		_ = traverse___84386644_4_2
		return gopurs_runtime.Func(func(dictFoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): traverse_1_6_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope8)] (TypeApp Any [(TypeVar b$scope9)])), (TypeApp (TypeVar t$scope7) [(TypeVar a$scope8)])] (TypeApp Any [Unit]))
			traverse_1_6_3 := gopurs_runtime.Apply(traverse___84386644_4_2, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_5))})
			_ = traverse_1_6_3
			return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply(traverse_1_6_3, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel_2_1, f_7)))
			})
		})
	})
}

func Call_Control_Parallel_parTraverse(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [(TypeApp (TypeVar t$scope19) [(TypeVar b$scope21)])])] (TypeApp (TypeVar m$scope18) [(TypeApp (TypeVar t$scope19) [(TypeVar b$scope21)])]))
	sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
	_ = sequential_1_0
	// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope18) [(TypeVar b$scope21)])] (TypeApp Any [(TypeVar b$scope21)]))
	parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
	_ = parallel_2_1
	return gopurs_runtime.Func2(func(dictApplicative_3 gopurs_runtime.Value, dictTraversable_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse_5_2 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope20)] (TypeApp Any [(TypeVar b$scope21)])), (TypeApp (TypeVar t$scope19) [(TypeVar a$scope20)])] (TypeApp Any [(TypeApp (TypeVar t$scope19) [(TypeVar b$scope21)])]))
		traverse_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_4, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = traverse_5_2
		return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply(traverse_5_2, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel_2_1, f_6)))
		})
	})
}

func Call_Control_Parallel_parSequence_(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): parTraverse_1__858671173_1_0 shape=App(Var) bindingType=(ConstrainedType (Func [(Func [(TypeApp (TypeVar m$scope30) [(TypeVar a$scope28)])] (TypeApp (TypeVar m$scope30) [(TypeVar a$scope28)])), (TypeApp (TypeVar t$scope29) [(TypeApp (TypeVar m$scope30) [(TypeVar a$scope28)])])] (TypeApp (TypeVar m$scope30) [Unit])))
	parTraverse_1__858671173_1_0 := Call_Control_Parallel_parTraverse_(dictParallel_0)
	_ = parTraverse_1__858671173_1_0
	return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): parTraverse_2_3_1 shape=App(Other) bindingType=(ConstrainedType (Func [(Func [(TypeApp (TypeVar m$scope30) [(TypeVar a$scope28)])] (TypeApp (TypeVar m$scope30) [(TypeVar a$scope28)])), (TypeApp (TypeVar t$scope29) [(TypeApp (TypeVar m$scope30) [(TypeVar a$scope28)])])] (TypeApp (TypeVar m$scope30) [Unit])))
		parTraverse_2_3_1 := gopurs_runtime.Apply(parTraverse_1__858671173_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_2))})
		_ = parTraverse_2_3_1
		return gopurs_runtime.Func(func(dictFoldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(parTraverse_2_3_1, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_4))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_5
			}))
		})
	})
}

func Call_Control_Parallel_parSequence(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): parTraverse1__3560054810_1_0 shape=App(Var) bindingType=(ConstrainedType (Func [(Func [(TypeApp (TypeVar m$scope39) [(TypeVar a$scope37)])] (TypeApp (TypeVar m$scope39) [(TypeVar a$scope37)])), (TypeApp (TypeVar t$scope38) [(TypeApp (TypeVar m$scope39) [(TypeVar a$scope37)])])] (TypeApp (TypeVar m$scope39) [(TypeApp (TypeVar t$scope38) [(TypeVar a$scope37)])])))
	parTraverse1__3560054810_1_0 := Call_Control_Parallel_parTraverse(dictParallel_0)
	_ = parTraverse1__3560054810_1_0
	return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): parTraverse2_3_1 shape=App(Other) bindingType=(ConstrainedType (Func [(Func [(TypeApp (TypeVar m$scope39) [(TypeVar a$scope37)])] (TypeApp (TypeVar m$scope39) [(TypeVar a$scope37)])), (TypeApp (TypeVar t$scope38) [(TypeApp (TypeVar m$scope39) [(TypeVar a$scope37)])])] (TypeApp (TypeVar m$scope39) [(TypeApp (TypeVar t$scope38) [(TypeVar a$scope37)])])))
		parTraverse2_3_1 := gopurs_runtime.Apply(parTraverse1__3560054810_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_2))})
		_ = parTraverse2_3_1
		return gopurs_runtime.Func(func(dictTraversable_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(parTraverse2_3_1, gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_4))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_5
			}))
		})
	})
}

func Call_Control_Parallel_parOneOfMap(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [(TypeVar b$scope48)])] (TypeApp (TypeVar m$scope50) [(TypeVar b$scope48)]))
	sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
	_ = sequential_1_0
	// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope50) [(TypeVar b$scope48)])] (TypeApp Any [(TypeVar b$scope48)]))
	parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
	_ = parallel_2_1
	return gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Plus1_4_2 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [Any])
		Plus1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{}))
		_ = Plus1_4_2
		return gopurs_runtime.Func(func(dictFoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): oneOfMap_6_3 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope47)] (TypeApp Any [(TypeVar b$scope48)])), (TypeApp (TypeVar t$scope49) [(TypeVar a$scope47)])] (TypeApp Any [(TypeVar b$scope48)]))
			oneOfMap_6_3 := Call_Data_Foldable_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_5), Plus1_4_2)
			_ = oneOfMap_6_3
			return gopurs_runtime.Func2(func(dictFunctor_7 gopurs_runtime.Value, f_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply(oneOfMap_6_3, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel_2_1, f_8)))
			})
		})
	})
}

func Call_Control_Parallel_parOneOf(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): sequential_1_0 shape=App(Var) bindingType=(Func [(TypeApp Any [(TypeVar a$scope58)])] (TypeApp (TypeVar m$scope60) [(TypeVar a$scope58)]))
	sequential_1_0 := Call_Control_Parallel_Class_sequential(dictParallel_0)
	_ = sequential_1_0
	// TAST (Let): parallel_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope60) [(TypeVar a$scope58)])] (TypeApp Any [(TypeVar a$scope58)]))
	parallel_2_1 := Call_Control_Parallel_Class_parallel(dictParallel_0)
	_ = parallel_2_1
	return gopurs_runtime.Func(func(dictAlternative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Plus1_4_2 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [Any])
		Plus1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_3, "Plus1"), gopurs_runtime.Value{}))
		_ = Plus1_4_2
		return gopurs_runtime.Func(func(dictFoldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): oneOfMap_6_3 shape=App(Var) bindingType=(Func [(Func [(TypeApp (TypeVar m$scope60) [(TypeVar a$scope58)])] (TypeApp Any [(TypeVar a$scope58)])), (TypeApp (TypeVar t$scope59) [(TypeApp (TypeVar m$scope60) [(TypeVar a$scope58)])])] (TypeApp Any [(TypeVar a$scope58)]))
			oneOfMap_6_3 := Call_Data_Foldable_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_5), Plus1_4_2)
			_ = oneOfMap_6_3
			return gopurs_runtime.Func(func(dictFunctor_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential_1_0, gopurs_runtime.Apply(oneOfMap_6_3, parallel_2_1))
			})
		})
	})
}

func Call_Control_Parallel_parApply(dictParallel_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value], mf_1_loop gopurs_runtime.Value, ma_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictParallel_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dictParallel_0_loop
	_ = dictParallel_0
	var mf_1 gopurs_runtime.Value = mf_1_loop
	_ = mf_1
	var ma_2 gopurs_runtime.Value = ma_2_loop
	_ = ma_2
	return gopurs_runtime.Apply(dictParallel_0.V3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictParallel_0.V1, gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(dictParallel_0.V2, mf_1), gopurs_runtime.Apply(dictParallel_0.V2, ma_2)))
}
