package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_NonEmpty_Internal_go__append gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_go__append sync.Once

func Get_Data_Array_NonEmpty_Internal_go__append() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_go__append.Do(func() {
		cache_Data_Array_NonEmpty_Internal_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Array_NonEmpty_Internal_go__append
}

var cache_Data_Array_NonEmpty_Internal_NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_NonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_NonEmptyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_Internal_NonEmptyArray(*((*[]gopurs_runtime.Value)(x_0_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_Internal_NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066 sync.Once

func Get_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066.Do(func() {
		cache_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066(*((*[]gopurs_runtime.Value)(x_0_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066
}

var cache_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786 sync.Once

func Get_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786.Do(func() {
		cache_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
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
	return cache_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786
}

var cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_2955889203_1812164904(Rebox_Data_Array_NonEmpty_Internal_1812164904_2955889203(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()))))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_traversableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_showNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_showNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_showNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_showNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_showNonEmptyArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_Internal_showNonEmptyArray(dictShow_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_Internal_showNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_986430664_4179793454(Rebox_Data_Array_NonEmpty_Internal_4179793454_986430664(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupArray()))))}
	})
	return cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_ordNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_ordNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_ordNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_ordNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_ordNonEmptyArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_Internal_ordNonEmptyArray(dictOrd_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_Internal_ordNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_ord1NonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_ord1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_ord1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_Ord_ord1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_monadNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_monadNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_monadNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_monadArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_2773701683_2412140840(Rebox_Data_Array_NonEmpty_Internal_2412140840_2773701683(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexArray()))))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_functorNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_2491554675_3725484264(Rebox_Data_Array_NonEmpty_Internal_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexArray()))))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_foldableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573((&Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
		}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Foldable_foldMap1DefaultL(Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573(Rebox_Data_Array_NonEmpty_Internal_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))
		}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), __local_var_0, __local_var_1)
		}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldr1Impl(), __local_var_0, __local_var_1)
		})})))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_3516332491_306175789((&Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573(Rebox_Data_Array_NonEmpty_Internal_4151366573_3016437835(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}
		}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_sequence1Default(Rebox_Data_Array_NonEmpty_Internal_3516332491_306175789(Rebox_Data_Array_NonEmpty_Internal_306175789_3516332491(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray()))), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))
		}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): apply_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope8) [(Func [(TypeVar a'$scope20)] (TypeVar b'$scope21))]), (TypeApp (TypeVar f$scope8) [(TypeVar a'$scope20)])] (TypeApp (TypeVar f$scope8) [(TypeVar b'$scope21)]))
			apply_1_0 := Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))
			_ = apply_1_0
			// TAST (Let): go__map_2_1 shape=App(Var) bindingType=(Func [(Func [(TypeVar a'$scope22)] (TypeVar b'$scope23)), (TypeApp (TypeVar f$scope8) [(TypeVar a'$scope22)])] (TypeApp (TypeVar f$scope8) [(TypeVar b'$scope23)]))
			go__map_2_1 := Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})))
			_ = go__map_2_1
			return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.UncurriedApp3(Get_Data_Array_NonEmpty_Internal_traverse1Impl(), apply_1_0, go__map_2_1, f_3)
			})
		})})))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_eqNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_eqNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_eqNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_eqNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_eqNonEmptyArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_NonEmpty_Internal_eqNonEmptyArray(dictEq_0_box)
		})
	})
	return cache_Data_Array_NonEmpty_Internal_eqNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_eq1NonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_eq1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_eq1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_bindNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_bindNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_bindNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Bind_bindArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applyNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_applyNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applyNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_altNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_altNonEmptyArray sync.Once

func Get_Data_Array_NonEmpty_Internal_altNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_altNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_altNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Control_Alt_altArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_altNonEmptyArray
}

func Call_Data_Array_NonEmpty_Internal_NonEmptyArray(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 []gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Array_NonEmpty_Internal_NonEmptyArray__4190607066(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
NonEmptyArray__4190607066:
	for {
		if false {
			continue NonEmptyArray__4190607066
		}
		var x_0 []gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Data_Array_NonEmpty_Internal_NonEmptyArray__4020493786(x_0_loop []int64) []int64 {
NonEmptyArray__4020493786:
	for {
		if false {
			continue NonEmptyArray__4020493786
		}
		var x_0 []int64 = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Data_Array_NonEmpty_Internal_showNonEmptyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_1356436936_1386611502((&Constructor_Data_Show_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(NonEmptyArray ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(dictShow_0), "show"), gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(v_1.UnsafePtr)))).StrVal()), ")"))
	})})))}
}

func Call_Data_Array_NonEmpty_Internal_ordNonEmptyArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_4088624520_4177771502(Rebox_Data_Array_NonEmpty_Internal_4177771502_4088624520(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Ord_ordArray(dictOrd_0)))))}
}

func Call_Data_Array_NonEmpty_Internal_eqNonEmptyArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_1939691112_3790796878(Rebox_Data_Array_NonEmpty_Internal_3790796878_1939691112(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(dictEq_0)))))}
}

func Rebox_Data_Array_NonEmpty_Internal_1356436936_1386611502(in *Constructor_Data_Show_Show[[]gopurs_runtime.Value]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_1812164904_2955889203(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_1939691112_3790796878(in *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_2412140840_2773701683(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_2773701683_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_2955889203_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_306175789_3516332491(in *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_3516332491_306175789(in *Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_3725484264_2491554675(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_3790796878_1939691112(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_4088624520_4177771502(in *Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_4151366573_3016437835(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_4177771502_4088624520(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_4179793454_986430664(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Array_NonEmpty_Internal_986430664_4179793454(in *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_Array_NonEmpty_Internal_foldl1Impl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_NonEmpty_Internal_Foldl1Impl
}

func Get_Data_Array_NonEmpty_Internal_foldr1Impl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_NonEmpty_Internal_Foldr1Impl
}

func Get_Data_Array_NonEmpty_Internal_traverse1Impl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_NonEmpty_Internal_Traverse1Impl
}
