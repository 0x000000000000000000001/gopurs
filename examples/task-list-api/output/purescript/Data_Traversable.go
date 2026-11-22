package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Traversable_identity gopurs_runtime.Value
var once_Data_Traversable_identity sync.Once

func Get_Data_Traversable_identity() gopurs_runtime.Value {
	once_Data_Traversable_identity.Do(func() {
		cache_Data_Traversable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_identity(x_0_box)
		})
	})
	return cache_Data_Traversable_identity
}

var cache_Data_Traversable_identity1 gopurs_runtime.Value
var once_Data_Traversable_identity1 sync.Once

func Get_Data_Traversable_identity1() gopurs_runtime.Value {
	once_Data_Traversable_identity1.Do(func() {
		cache_Data_Traversable_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_identity1(x_0_box)
		})
	})
	return cache_Data_Traversable_identity1
}

var cache_Data_Traversable_go__append gopurs_runtime.Value
var once_Data_Traversable_go__append sync.Once

func Get_Data_Traversable_go__append() gopurs_runtime.Value {
	once_Data_Traversable_go__append.Do(func() {
		cache_Data_Traversable_go__append = Get_Data_Semigroup_concatArray()
	})
	return cache_Data_Traversable_go__append
}

var cache_Data_Traversable_compose gopurs_runtime.Value
var once_Data_Traversable_compose sync.Once

func Get_Data_Traversable_compose() gopurs_runtime.Value {
	once_Data_Traversable_compose.Do(func() {
		cache_Data_Traversable_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose
}

var cache_Data_Traversable_compose1 gopurs_runtime.Value
var once_Data_Traversable_compose1 sync.Once

func Get_Data_Traversable_compose1() gopurs_runtime.Value {
	once_Data_Traversable_compose1.Do(func() {
		cache_Data_Traversable_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose1
}

var cache_Data_Traversable_compose2 gopurs_runtime.Value
var once_Data_Traversable_compose2 sync.Once

func Get_Data_Traversable_compose2() gopurs_runtime.Value {
	once_Data_Traversable_compose2.Do(func() {
		cache_Data_Traversable_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose2
}

var cache_Data_Traversable_compose3 gopurs_runtime.Value
var once_Data_Traversable_compose3 sync.Once

func Get_Data_Traversable_compose3() gopurs_runtime.Value {
	once_Data_Traversable_compose3.Do(func() {
		cache_Data_Traversable_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose3
}

var cache_Data_Traversable_compose4 gopurs_runtime.Value
var once_Data_Traversable_compose4 sync.Once

func Get_Data_Traversable_compose4() gopurs_runtime.Value {
	once_Data_Traversable_compose4.Do(func() {
		cache_Data_Traversable_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose4
}

var cache_Data_Traversable_compose5 gopurs_runtime.Value
var once_Data_Traversable_compose5 sync.Once

func Get_Data_Traversable_compose5() gopurs_runtime.Value {
	once_Data_Traversable_compose5.Do(func() {
		cache_Data_Traversable_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose5
}

var cache_Data_Traversable_compose6 gopurs_runtime.Value
var once_Data_Traversable_compose6 sync.Once

func Get_Data_Traversable_compose6() gopurs_runtime.Value {
	once_Data_Traversable_compose6.Do(func() {
		cache_Data_Traversable_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose6
}

var cache_Data_Traversable_compose7 gopurs_runtime.Value
var once_Data_Traversable_compose7 sync.Once

func Get_Data_Traversable_compose7() gopurs_runtime.Value {
	once_Data_Traversable_compose7.Do(func() {
		cache_Data_Traversable_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Traversable_compose7
}

var cache_Data_Traversable_Traversable_dollar_Dict gopurs_runtime.Value
var once_Data_Traversable_Traversable_dollar_Dict sync.Once

func Get_Data_Traversable_Traversable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Traversable_Traversable_dollar_Dict.Do(func() {
		cache_Data_Traversable_Traversable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Call_Data_Traversable_Traversable_dollar_Dict(func() struct {
				Foldable1 gopurs_runtime.Value
				Functor0  gopurs_runtime.Value
				sequence  gopurs_runtime.Value
				traverse  gopurs_runtime.Value
				/*gopurs:record-v1:4:0046006f006c006400610062006c00650031,00460075006e00630074006f00720030,00730065007100750065006e00630065,00740072006100760065007200730065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Foldable1 gopurs_runtime.Value
					Functor0  gopurs_runtime.Value
					sequence  gopurs_runtime.Value
					traverse  gopurs_runtime.Value
					/*gopurs:record-v1:4:0046006f006c006400610062006c00650031,00460075006e00630074006f00720030,00730065007100750065006e00630065,00740072006100760065007200730065*/
				}{}
				clone.Foldable1 = gopurs_runtime.RecordGet(orig, "Foldable1")
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.sequence = gopurs_runtime.RecordGet(orig, "sequence")
				clone.traverse = gopurs_runtime.RecordGet(orig, "traverse")
				return clone
			}()))}
		})
	})
	return cache_Data_Traversable_Traversable_dollar_Dict
}

var cache_Data_Traversable_traverse gopurs_runtime.Value
var once_Data_Traversable_traverse sync.Once

func Get_Data_Traversable_traverse() gopurs_runtime.Value {
	once_Data_Traversable_traverse.Do(func() {
		cache_Data_Traversable_traverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Traversable_traverse
}

var cache_Data_Traversable_sequenceDefault gopurs_runtime.Value
var once_Data_Traversable_sequenceDefault sync.Once

func Get_Data_Traversable_sequenceDefault() gopurs_runtime.Value {
	once_Data_Traversable_sequenceDefault.Do(func() {
		cache_Data_Traversable_sequenceDefault = gopurs_runtime.Func2(func(dictTraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Traversable_sequenceDefault
}

var cache_Data_Traversable_traversableArray gopurs_runtime.Value
var once_Data_Traversable_traversableArray sync.Once

func Get_Data_Traversable_traversableArray() gopurs_runtime.Value {
	once_Data_Traversable_traversableArray.Do(func() {
		cache_Data_Traversable_traversableArray = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=Any
			Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
			_ = Apply0_1_0
			return gopurs_runtime.Apply4(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_1_0)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0)), Get_Data_Semigroup_concatArray())
		})}))}
	})
	return cache_Data_Traversable_traversableArray
}

var cache_Data_Traversable_traverse__2498349992 gopurs_runtime.Value
var once_Data_Traversable_traverse__2498349992 sync.Once

func Get_Data_Traversable_traverse__2498349992() gopurs_runtime.Value {
	once_Data_Traversable_traverse__2498349992.Do(func() {
		cache_Data_Traversable_traverse__2498349992 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traverse__2498349992(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Traversable_traverse__2498349992
}

var cache_Data_Traversable_traversableTuple gopurs_runtime.Value
var once_Data_Traversable_traversableTuple sync.Once

func Get_Data_Traversable_traversableTuple() gopurs_runtime.Value {
	once_Data_Traversable_traversableTuple.Do(func() {
		cache_Data_Traversable_traversableTuple = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3543431075_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Foldable_foldableTuple()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Tuple_functorTuple()
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
		})})))}
	})
	return cache_Data_Traversable_traversableTuple
}

var cache_Data_Traversable_traversableMultiplicative gopurs_runtime.Value
var once_Data_Traversable_traversableMultiplicative sync.Once

func Get_Data_Traversable_traversableMultiplicative() gopurs_runtime.Value {
	once_Data_Traversable_traversableMultiplicative.Do(func() {
		cache_Data_Traversable_traversableMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Multiplicative_functorMultiplicative()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Multiplicative_Multiplicative(), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Traversable_traversableMultiplicative
}

var cache_Data_Traversable_traversableMaybe gopurs_runtime.Value
var once_Data_Traversable_traversableMaybe sync.Once

func Get_Data_Traversable_traversableMaybe() gopurs_runtime.Value {
	once_Data_Traversable_traversableMaybe.Do(func() {
		cache_Data_Traversable_traversableMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3188237647_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_1146820559_1680800814(Rebox_Data_Traversable_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3689823567_2812149806(Rebox_Data_Traversable_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil {
					__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
					goto end_branch_0
				} else {

				}
			}
			{
				if v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil {
					__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil {
					__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
					goto end_branch_1
				} else {

				}
			}
			{
				if v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil {
					__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		})})))}
	})
	return cache_Data_Traversable_traversableMaybe
}

var cache_Data_Traversable_traverse1 gopurs_runtime.Value
var once_Data_Traversable_traverse1 sync.Once

func Get_Data_Traversable_traverse1() gopurs_runtime.Value {
	once_Data_Traversable_traverse1.Do(func() {
		cache_Data_Traversable_traverse1 = Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V3
	})
	return cache_Data_Traversable_traverse1
}

var cache_Data_Traversable_traverse2 gopurs_runtime.Value
var once_Data_Traversable_traverse2 sync.Once

func Get_Data_Traversable_traverse2() gopurs_runtime.Value {
	once_Data_Traversable_traverse2.Do(func() {
		cache_Data_Traversable_traverse2 = Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V3
	})
	return cache_Data_Traversable_traverse2
}

var cache_Data_Traversable_traversableIdentity gopurs_runtime.Value
var once_Data_Traversable_traversableIdentity sync.Once

func Get_Data_Traversable_traversableIdentity() gopurs_runtime.Value {
	once_Data_Traversable_traversableIdentity.Do(func() {
		cache_Data_Traversable_traversableIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableIdentity()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Identity_Identity(), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Traversable_traversableIdentity
}

var cache_Data_Traversable_traversableEither gopurs_runtime.Value
var once_Data_Traversable_traversableEither sync.Once

func Get_Data_Traversable_traversableEither() gopurs_runtime.Value {
	once_Data_Traversable_traversableEither.Do(func() {
		cache_Data_Traversable_traversableEither = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Foldable_foldableEither()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_functorEither()
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v_1.Type == 9 && v_1.IntVal == 3711209382 {
					__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}())
					goto end_branch_0
				} else {

				}
			}
			{
				if v_1.Type == 9 && v_1.IntVal == 2465973597 {
					__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Either_Right(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 gopurs_runtime.Value
			{
				if v1_2.Type == 9 && v1_2.IntVal == 3711209382 {
					__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
							V2 bool
						}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
						if _v.V2 {
							return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
					}())
					goto end_branch_1
				} else {

				}
			}
			{
				if v1_2.Type == 9 && v1_2.IntVal == 2465973597 {
					__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Either_Right(), gopurs_runtime.Apply(v_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		})}))}
	})
	return cache_Data_Traversable_traversableEither
}

var cache_Data_Traversable_traversableDual gopurs_runtime.Value
var once_Data_Traversable_traversableDual sync.Once

func Get_Data_Traversable_traversableDual() gopurs_runtime.Value {
	once_Data_Traversable_traversableDual.Do(func() {
		cache_Data_Traversable_traversableDual = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Dual_functorDual()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Dual_Dual(), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Traversable_traversableDual
}

var cache_Data_Traversable_traversableDisj gopurs_runtime.Value
var once_Data_Traversable_traversableDisj sync.Once

func Get_Data_Traversable_traversableDisj() gopurs_runtime.Value {
	once_Data_Traversable_traversableDisj.Do(func() {
		cache_Data_Traversable_traversableDisj = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Disj_functorDisj()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Disj_Disj(), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Disj_Disj(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Traversable_traversableDisj
}

var cache_Data_Traversable_traversableConst gopurs_runtime.Value
var once_Data_Traversable_traversableConst sync.Once

func Get_Data_Traversable_traversableConst() gopurs_runtime.Value {
	once_Data_Traversable_traversableConst.Do(func() {
		cache_Data_Traversable_traversableConst = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Foldable_foldableConst()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Const_functorConst()
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
		})}))}
	})
	return cache_Data_Traversable_traversableConst
}

var cache_Data_Traversable_traversableConj gopurs_runtime.Value
var once_Data_Traversable_traversableConj sync.Once

func Get_Data_Traversable_traversableConj() gopurs_runtime.Value {
	once_Data_Traversable_traversableConj.Do(func() {
		cache_Data_Traversable_traversableConj = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Conj_functorConj()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Conj_Conj(), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Conj_Conj(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Traversable_traversableConj
}

var cache_Data_Traversable_traversableCompose gopurs_runtime.Value
var once_Data_Traversable_traversableCompose sync.Once

func Get_Data_Traversable_traversableCompose() gopurs_runtime.Value {
	once_Data_Traversable_traversableCompose.Do(func() {
		cache_Data_Traversable_traversableCompose = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traversableCompose(dictTraversable_0_box)
		})
	})
	return cache_Data_Traversable_traversableCompose
}

var cache_Data_Traversable_traversableAdditive gopurs_runtime.Value
var once_Data_Traversable_traversableAdditive sync.Once

func Get_Data_Traversable_traversableAdditive() gopurs_runtime.Value {
	once_Data_Traversable_traversableAdditive.Do(func() {
		cache_Data_Traversable_traversableAdditive = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Additive_functorAdditive()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Additive_Additive(), v_1)
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Additive_Additive(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Traversable_traversableAdditive
}

var cache_Data_Traversable_sequence gopurs_runtime.Value
var once_Data_Traversable_sequence sync.Once

func Get_Data_Traversable_sequence() gopurs_runtime.Value {
	once_Data_Traversable_sequence.Do(func() {
		cache_Data_Traversable_sequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Traversable_sequence
}

var cache_Data_Traversable_sequence1 gopurs_runtime.Value
var once_Data_Traversable_sequence1 sync.Once

func Get_Data_Traversable_sequence1() gopurs_runtime.Value {
	once_Data_Traversable_sequence1.Do(func() {
		cache_Data_Traversable_sequence1 = Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V2
	})
	return cache_Data_Traversable_sequence1
}

var cache_Data_Traversable_sequence2 gopurs_runtime.Value
var once_Data_Traversable_sequence2 sync.Once

func Get_Data_Traversable_sequence2() gopurs_runtime.Value {
	once_Data_Traversable_sequence2.Do(func() {
		cache_Data_Traversable_sequence2 = Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V2
	})
	return cache_Data_Traversable_sequence2
}

var cache_Data_Traversable_traversableApp gopurs_runtime.Value
var once_Data_Traversable_traversableApp sync.Once

func Get_Data_Traversable_traversableApp() gopurs_runtime.Value {
	once_Data_Traversable_traversableApp.Do(func() {
		cache_Data_Traversable_traversableApp = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traversableApp(dictTraversable_0_box)
		})
	})
	return cache_Data_Traversable_traversableApp
}

var cache_Data_Traversable_traversableCoproduct gopurs_runtime.Value
var once_Data_Traversable_traversableCoproduct sync.Once

func Get_Data_Traversable_traversableCoproduct() gopurs_runtime.Value {
	once_Data_Traversable_traversableCoproduct.Do(func() {
		cache_Data_Traversable_traversableCoproduct = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traversableCoproduct(dictTraversable_0_box)
		})
	})
	return cache_Data_Traversable_traversableCoproduct
}

var cache_Data_Traversable_traversableFirst gopurs_runtime.Value
var once_Data_Traversable_traversableFirst sync.Once

func Get_Data_Traversable_traversableFirst() gopurs_runtime.Value {
	once_Data_Traversable_traversableFirst.Do(func() {
		cache_Data_Traversable_traversableFirst = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3188237647_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_1146820559_1680800814(Rebox_Data_Traversable_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3689823567_2812149806(Rebox_Data_Traversable_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_First_First(), gopurs_runtime.Apply2(Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1))}))
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_First_First(), gopurs_runtime.Apply3(Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2))}))
		})})))}
	})
	return cache_Data_Traversable_traversableFirst
}

var cache_Data_Traversable_traversableLast gopurs_runtime.Value
var once_Data_Traversable_traversableLast sync.Once

func Get_Data_Traversable_traversableLast() gopurs_runtime.Value {
	once_Data_Traversable_traversableLast.Do(func() {
		cache_Data_Traversable_traversableLast = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3188237647_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_1146820559_1680800814(Rebox_Data_Traversable_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3689823567_2812149806(Rebox_Data_Traversable_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_Last_Last(), gopurs_runtime.Apply2(Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1))}))
		}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Maybe_Last_Last(), gopurs_runtime.Apply3(Rebox_Data_Traversable_3043886126_3188237647(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMaybe())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2))}))
		})})))}
	})
	return cache_Data_Traversable_traversableLast
}

var cache_Data_Traversable_traversableProduct gopurs_runtime.Value
var once_Data_Traversable_traversableProduct sync.Once

func Get_Data_Traversable_traversableProduct() gopurs_runtime.Value {
	once_Data_Traversable_traversableProduct.Do(func() {
		cache_Data_Traversable_traversableProduct = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traversableProduct(dictTraversable_0_box)
		})
	})
	return cache_Data_Traversable_traversableProduct
}

var cache_Data_Traversable_traverseDefault gopurs_runtime.Value
var once_Data_Traversable_traverseDefault sync.Once

func Get_Data_Traversable_traverseDefault() gopurs_runtime.Value {
	once_Data_Traversable_traverseDefault.Do(func() {
		cache_Data_Traversable_traverseDefault = gopurs_runtime.Func2(func(dictTraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Traversable_traverseDefault
}

var cache_Data_Traversable_mapAccumR gopurs_runtime.Value
var once_Data_Traversable_mapAccumR sync.Once

func Get_Data_Traversable_mapAccumR() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumR.Do(func() {
		cache_Data_Traversable_mapAccumR = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_mapAccumR(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
		})
	})
	return cache_Data_Traversable_mapAccumR
}

var cache_Data_Traversable_scanr gopurs_runtime.Value
var once_Data_Traversable_scanr sync.Once

func Get_Data_Traversable_scanr() gopurs_runtime.Value {
	once_Data_Traversable_scanr.Do(func() {
		cache_Data_Traversable_scanr = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_scanr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
		})
	})
	return cache_Data_Traversable_scanr
}

var cache_Data_Traversable_mapAccumL gopurs_runtime.Value
var once_Data_Traversable_mapAccumL sync.Once

func Get_Data_Traversable_mapAccumL() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumL.Do(func() {
		cache_Data_Traversable_mapAccumL = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_mapAccumL(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
		})
	})
	return cache_Data_Traversable_mapAccumL
}

var cache_Data_Traversable_scanl gopurs_runtime.Value
var once_Data_Traversable_scanl sync.Once

func Get_Data_Traversable_scanl() gopurs_runtime.Value {
	once_Data_Traversable_scanl.Do(func() {
		cache_Data_Traversable_scanl = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_scanl(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
		})
	})
	return cache_Data_Traversable_scanl
}

var cache_Data_Traversable_go__for gopurs_runtime.Value
var once_Data_Traversable_go__for sync.Once

func Get_Data_Traversable_go__for() gopurs_runtime.Value {
	once_Data_Traversable_go__for.Do(func() {
		cache_Data_Traversable_go__for = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_go__for(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_1_box))
		})
	})
	return cache_Data_Traversable_go__for
}

type Constructor_Data_Traversable_Traversable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f005400720061007600650072007300610062006c0065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f005400720061007600650072007300610062006c0065005f005400720061007600650072007300610062006c0065
func init() {
	gopurs_runtime.StructGetters[3941073978] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Foldable1":
			return gopurs_runtime.Box(c.V0)
		case "Functor0":
			return gopurs_runtime.Box(c.V1)
		case "sequence":
			return gopurs_runtime.Box(c.V2)
		case "traverse":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Data_Traversable_Traversable: " + key)
		}
	}
}

func Call_Data_Traversable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Traversable_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Traversable_Traversable_dollar_Dict(x_0_loop struct {
	Foldable1 gopurs_runtime.Value
	Functor0  gopurs_runtime.Value
	sequence  gopurs_runtime.Value
	traverse  gopurs_runtime.Value
	/*gopurs:record-v1:4:0046006f006c006400610062006c00650031,00460075006e00630074006f00720030,00730065007100750065006e00630065,00740072006100760065007200730065*/
}) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	var x_0 struct {
		Foldable1 gopurs_runtime.Value
		Functor0  gopurs_runtime.Value
		sequence  gopurs_runtime.Value
		traverse  gopurs_runtime.Value
		/*gopurs:record-v1:4:0046006f006c006400610062006c00650031,00460075006e00630074006f00720030,00730065007100750065006e00630065,00740072006100760065007200730065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", orig.Foldable1, orig.Functor0, orig.sequence, orig.traverse)
	}())
}

func Call_Data_Traversable_traverse(dict_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Data_Traversable_sequenceDefault(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	return gopurs_runtime.Apply2(dictTraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}))
}

func Call_Data_Traversable_traverse__2498349992(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
traverse__2498349992:
	for {
		if false {
			continue traverse__2498349992
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()))}, __eta0_0, __eta1_1)
	}
}

func Call_Data_Traversable_traversableCompose(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traversableCompose:
	for {
		if false {
			continue traversableCompose
		}
		var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
		_ = dictTraversable_0
		// TAST (Let): functorCompose_1_0 shape=App(Var) bindingType=Any
		functorCompose_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Compose_functorCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
		_ = functorCompose_1_0
		// TAST (Let): foldableCompose_2_1 shape=App(Var) bindingType=Any
		foldableCompose_2_1 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
		_ = foldableCompose_2_1
		return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): functorCompose1_4_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope143), (TypeVar g$scope144)])])
			functorCompose1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})))
			_ = functorCompose1_4_2
			// TAST (Let): foldableCompose1_5_3 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope143), (TypeVar g$scope144)])])
			foldableCompose1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCompose_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{})))
			_ = foldableCompose1_5_3
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCompose1_5_3)}
			}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose1_4_2)}
			}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Traversable_traversableCompose(dictTraversable_0), dictTraversable1_3), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_7
				}))
			}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): traverse5_7_4 shape=App(Other) bindingType=(Func [(Func [(TypeApp (TypeVar g$scope144) [(TypeVar a$scope151)])] (TypeApp (TypeVar m$scope153) [(TypeApp (TypeVar g$scope144) [(TypeVar b$scope152)])])), (TypeApp (TypeVar f$scope143) [(TypeApp (TypeVar g$scope144) [(TypeVar a$scope151)])])] (TypeApp (TypeVar m$scope153) [(TypeApp (TypeVar f$scope143) [(TypeApp (TypeVar g$scope144) [(TypeVar b$scope152)])])]))
				traverse5_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
				_ = traverse5_7_4
				// TAST (Let): traverse6_8_5 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope151)] (TypeApp (TypeVar m$scope153) [(TypeVar b$scope152)])), (TypeApp (TypeVar g$scope144) [(TypeVar a$scope151)])] (TypeApp (TypeVar m$scope153) [(TypeApp (TypeVar g$scope144) [(TypeVar b$scope152)])]))
				traverse6_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
				_ = traverse6_8_5
				return gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Compose_Compose(), gopurs_runtime.Apply2(traverse5_7_4, gopurs_runtime.Apply(traverse6_8_5, f_9), v_10))
				})
			})}))}
		})
	}
}

func Call_Data_Traversable_sequence(dict_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Traversable_traversableApp(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): functorApp_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope210)])])
	functorApp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorApp_1_0
	// TAST (Let): foldableApp_2_1 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope210)])])
	foldableApp_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_Foldable_foldableApp(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})))
	_ = foldableApp_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableApp_2_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorApp_1_0)}
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): sequence4_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope210) [(TypeApp (TypeVar m$scope224) [(TypeVar a$scope223)])])] (TypeApp (TypeVar m$scope224) [(TypeApp (TypeVar f$scope210) [(TypeVar a$scope223)])]))
		sequence4_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = sequence4_4_2
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_App_App(), gopurs_runtime.Apply(sequence4_4_2, v_5))
		})
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse4_4_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope217)] (TypeApp (TypeVar m$scope219) [(TypeVar b$scope218)])), (TypeApp (TypeVar f$scope210) [(TypeVar a$scope217)])] (TypeApp (TypeVar m$scope219) [(TypeApp (TypeVar f$scope210) [(TypeVar b$scope218)])]))
		traverse4_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = traverse4_4_3
		return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_App_App(), gopurs_runtime.Apply2(traverse4_4_3, f_5, v_6))
		})
	})}))}
}

func Call_Data_Traversable_traversableCoproduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): sequence3_1_0 shape=App(Var) bindingType=Any
	sequence3_1_0 := Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0))
	_ = sequence3_1_0
	// TAST (Let): functorCoproduct_2_1 shape=App(Var) bindingType=Any
	functorCoproduct_2_1 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorCoproduct_2_1
	// TAST (Let): foldableCoproduct_3_2 shape=App(Var) bindingType=Any
	foldableCoproduct_3_2 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
	_ = foldableCoproduct_3_2
	return gopurs_runtime.Func(func(dictTraversable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): sequence4_5_3 shape=App(Var) bindingType=Any
		sequence4_5_3 := Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable1_4))
		_ = sequence4_5_3
		// TAST (Let): functorCoproduct1_6_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope230), (TypeVar g$scope231)])])
		functorCoproduct1_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCoproduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_4, "Functor0"), gopurs_runtime.Value{})))
		_ = functorCoproduct1_6_4
		// TAST (Let): foldableCoproduct1_7_5 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope230), (TypeVar g$scope231)])])
		foldableCoproduct1_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCoproduct_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_4, "Foldable1"), gopurs_runtime.Value{})))
		_ = foldableCoproduct1_7_5
		return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCoproduct1_7_5)}
		}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCoproduct1_6_4)}
		}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Functor0_9_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope259)])
			Functor0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_9_6
			return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_9_6.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Left())), gopurs_runtime.Apply(sequence3_1_0, dictApplicative_8)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_9_6.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Right())), gopurs_runtime.Apply(sequence4_5_3, dictApplicative_8)))
		}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Functor0_9_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope240)])
			Functor0_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_9_7
			// TAST (Let): traverse5_10_8 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope238)] (TypeApp (TypeVar m$scope240) [(TypeVar b$scope239)])), (TypeApp (TypeVar f$scope230) [(TypeVar a$scope238)])] (TypeApp (TypeVar m$scope240) [(TypeApp (TypeVar f$scope230) [(TypeVar b$scope239)])]))
			traverse5_10_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))})
			_ = traverse5_10_8
			// TAST (Let): traverse6_11_9 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope238)] (TypeApp (TypeVar m$scope240) [(TypeVar b$scope239)])), (TypeApp (TypeVar g$scope231) [(TypeVar a$scope238)])] (TypeApp (TypeVar m$scope240) [(TypeApp (TypeVar g$scope231) [(TypeVar b$scope239)])]))
			traverse6_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_4, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))})
			_ = traverse6_11_9
			return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_9_7.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Left())), gopurs_runtime.Apply(traverse5_10_8, f_12)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_9_7.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Right())), gopurs_runtime.Apply(traverse6_11_9, f_12)))
			})
		})}))}
	})
}

func Call_Data_Traversable_traversableProduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): functorProduct_1_0 shape=App(Var) bindingType=Any
	functorProduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorProduct_1_0
	// TAST (Let): foldableProduct_2_1 shape=App(Var) bindingType=Any
	foldableProduct_2_1 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
	_ = foldableProduct_2_1
	return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): functorProduct1_4_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope305), (TypeVar g$scope306)])])
		functorProduct1_4_2 := Rebox_Data_Traversable_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{}))))
		_ = functorProduct1_4_2
		// TAST (Let): foldableProduct1_5_3 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope305), (TypeVar g$scope306)])])
		foldableProduct1_5_3 := Rebox_Data_Traversable_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableProduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))))
		_ = foldableProduct1_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3543431075_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_4173511203_1680800814(foldableProduct1_5_3))}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_2363162019_2812149806(functorProduct1_4_2))}
		}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_7_4 shape=App(Other) bindingType=Any
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
			_ = __local_var_7_4
			// TAST (Let): sequence5_8_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope305) [(TypeApp (TypeVar m$scope318) [(TypeVar a$scope317)])])] (TypeApp (TypeVar m$scope318) [(TypeApp (TypeVar f$scope305) [(TypeVar a$scope317)])]))
			sequence5_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = sequence5_8_5
			// TAST (Let): sequence6_9_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope306) [(TypeApp (TypeVar m$scope318) [(TypeVar a$scope317)])])] (TypeApp (TypeVar m$scope318) [(TypeApp (TypeVar g$scope306) [(TypeVar a$scope317)])]))
			sequence6_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = sequence6_9_6
			return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_4, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply(sequence5_8_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0)), gopurs_runtime.Apply(sequence6_9_6, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1))
			})
		}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_7_7 shape=App(Other) bindingType=Any
			__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
			_ = __local_var_7_7
			// TAST (Let): traverse5_8_8 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope313)] (TypeApp (TypeVar m$scope315) [(TypeVar b$scope314)])), (TypeApp (TypeVar f$scope305) [(TypeVar a$scope313)])] (TypeApp (TypeVar m$scope315) [(TypeApp (TypeVar f$scope305) [(TypeVar b$scope314)])]))
			traverse5_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = traverse5_8_8
			// TAST (Let): traverse6_9_9 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope313)] (TypeApp (TypeVar m$scope315) [(TypeVar b$scope314)])), (TypeApp (TypeVar g$scope306) [(TypeVar a$scope313)])] (TypeApp (TypeVar m$scope315) [(TypeApp (TypeVar g$scope306) [(TypeVar b$scope314)])]))
			traverse6_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = traverse6_9_9
			return gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply2(traverse5_8_8, f_10, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply2(traverse6_9_9, f_10, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
			})
		})})))}
	})
}

func Call_Data_Traversable_traverseDefault(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	// TAST (Let): sequence4_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t$scope324) [(TypeApp (TypeVar m$scope327) [(TypeVar b$scope326)])])] (TypeApp (TypeVar m$scope327) [(TypeApp (TypeVar t$scope324) [(TypeVar b$scope326)])]))
	sequence4_2_0 := gopurs_runtime.Apply(dictTraversable_0.V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
	_ = sequence4_2_0
	return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence4_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictTraversable_0.V1, gopurs_runtime.Value{}), "map"), f_3, ta_4))
	})
}

func Call_Data_Traversable_mapAccumR(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): traverse3_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope332), (TypeVar s$scope334)] (Record (Row [accum: (TypeVar s$scope334), value: (TypeVar b$scope333)] Empty))), (TypeApp (TypeVar f$scope335) [(TypeVar a$scope332)]), (TypeVar s$scope334)] (Record (Row [accum: (TypeVar s$scope334), value: (TypeApp (TypeVar f$scope335) [(TypeVar b$scope333)])] Empty)))
	traverse3_1_0 := gopurs_runtime.Apply(dictTraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Traversable_Accum_Internal_applicativeStateR()))})
	_ = traverse3_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, s0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(traverse3_1_0, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_2, s_6, a_5)
		}), xs_4, s0_3)
	})
}

func Call_Data_Traversable_scanr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): mapAccumR1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope342), (TypeVar a$scope341)] (Record (Row [accum: (TypeVar b$scope342), value: (TypeVar b$scope342)] Empty))), (TypeVar b$scope342), (TypeApp (TypeVar f$scope343) [(TypeVar a$scope341)])] (Record (Row [accum: (TypeVar b$scope342), value: (TypeApp (TypeVar f$scope343) [(TypeVar b$scope342)])] Empty)))
	mapAccumR1_1_0 := Call_Data_Traversable_mapAccumR(dictTraversable_0)
	_ = mapAccumR1_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumR1_1_0, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): b_prime__7_1 shape=App(Other) bindingType=(TypeVar b$scope342)
			b_prime__7_1 := gopurs_runtime.Apply2(f_2, a_6, b_5)
			_ = b_prime__7_1
			return gopurs_runtime.RecordDict2("accum", "value", b_prime__7_1, b_prime__7_1)
		}), b0_3, xs_4), "value")
	})
}

func Call_Data_Traversable_mapAccumL(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): traverse3_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope348), (TypeVar s$scope350)] (Record (Row [accum: (TypeVar s$scope350), value: (TypeVar b$scope349)] Empty))), (TypeApp (TypeVar f$scope351) [(TypeVar a$scope348)]), (TypeVar s$scope350)] (Record (Row [accum: (TypeVar s$scope350), value: (TypeApp (TypeVar f$scope351) [(TypeVar b$scope349)])] Empty)))
	traverse3_1_0 := gopurs_runtime.Apply(dictTraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Traversable_Accum_Internal_applicativeStateL()))})
	_ = traverse3_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, s0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(traverse3_1_0, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_2, s_6, a_5)
		}), xs_4, s0_3)
	})
}

func Call_Data_Traversable_scanl(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): mapAccumL1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope358), (TypeVar a$scope357)] (Record (Row [accum: (TypeVar b$scope358), value: (TypeVar b$scope358)] Empty))), (TypeVar b$scope358), (TypeApp (TypeVar f$scope359) [(TypeVar a$scope357)])] (Record (Row [accum: (TypeVar b$scope358), value: (TypeApp (TypeVar f$scope359) [(TypeVar b$scope358)])] Empty)))
	mapAccumL1_1_0 := Call_Data_Traversable_mapAccumL(dictTraversable_0)
	_ = mapAccumL1_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumL1_1_0, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): b_prime__7_1 shape=App(Other) bindingType=(TypeVar b$scope358)
			b_prime__7_1 := gopurs_runtime.Apply2(f_2, b_5, a_6)
			_ = b_prime__7_1
			return gopurs_runtime.RecordDict2("accum", "value", b_prime__7_1, b_prime__7_1)
		}), b0_3, xs_4), "value")
	})
}

func Call_Data_Traversable_go__for(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictTraversable_1_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictTraversable_1 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_1_loop
	_ = dictTraversable_1
	// TAST (Let): traverse3_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope364)] (TypeApp (TypeVar m$scope366) [(TypeVar b$scope365)])), (TypeApp (TypeVar t$scope367) [(TypeVar a$scope364)])] (TypeApp (TypeVar m$scope366) [(TypeApp (TypeVar t$scope367) [(TypeVar b$scope365)])]))
	traverse3_2_0 := gopurs_runtime.Apply(dictTraversable_1.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
	_ = traverse3_2_0
	return gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(traverse3_2_0, f_4, x_3)
	})
}

func Rebox_Data_Traversable_1146820559_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_1680800814_1146820559(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_1680800814_4173511203(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_3043886126_3188237647(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_3188237647_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_3543431075_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Traversable_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_Traversable_traverseArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Traversable_TraverseArrayImpl
}
