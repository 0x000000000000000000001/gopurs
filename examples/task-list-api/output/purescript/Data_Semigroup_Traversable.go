package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_Traversable_identity gopurs_runtime.Value
var once_Data_Semigroup_Traversable_identity sync.Once

func Get_Data_Semigroup_Traversable_identity() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_identity.Do(func() {
		cache_Data_Semigroup_Traversable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_identity(x_0_box)
		})
	})
	return cache_Data_Semigroup_Traversable_identity
}

var cache_Data_Semigroup_Traversable_Traversable1_dollar_Dict gopurs_runtime.Value
var once_Data_Semigroup_Traversable_Traversable1_dollar_Dict sync.Once

func Get_Data_Semigroup_Traversable_Traversable1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_Traversable1_dollar_Dict.Do(func() {
		cache_Data_Semigroup_Traversable_Traversable1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Call_Data_Semigroup_Traversable_Traversable1_dollar_Dict(func() struct {
				Foldable10   gopurs_runtime.Value
				Traversable1 gopurs_runtime.Value
				sequence1    gopurs_runtime.Value
				traverse1    gopurs_runtime.Value
				/*gopurs:record-v1:4:0046006f006c006400610062006c006500310030,005400720061007600650072007300610062006c00650031,00730065007100750065006e006300650031,007400720061007600650072007300650031*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Foldable10   gopurs_runtime.Value
					Traversable1 gopurs_runtime.Value
					sequence1    gopurs_runtime.Value
					traverse1    gopurs_runtime.Value
					/*gopurs:record-v1:4:0046006f006c006400610062006c006500310030,005400720061007600650072007300610062006c00650031,00730065007100750065006e006300650031,007400720061007600650072007300650031*/
				}{}
				clone.Foldable10 = gopurs_runtime.RecordGet(orig, "Foldable10")
				clone.Traversable1 = gopurs_runtime.RecordGet(orig, "Traversable1")
				clone.sequence1 = gopurs_runtime.RecordGet(orig, "sequence1")
				clone.traverse1 = gopurs_runtime.RecordGet(orig, "traverse1")
				return clone
			}()))}
		})
	})
	return cache_Data_Semigroup_Traversable_Traversable1_dollar_Dict
}

var cache_Data_Semigroup_Traversable_traverse1 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1 sync.Once

func Get_Data_Semigroup_Traversable_traverse1() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_traverse1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Semigroup_Traversable_traverse1
}

var cache_Data_Semigroup_Traversable_traversableTuple gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableTuple sync.Once

func Get_Data_Semigroup_Traversable_traversableTuple() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableTuple.Do(func() {
		cache_Data_Semigroup_Traversable_traversableTuple = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Traversable_1487937088_306175789((&Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Semigroup_Foldable_foldableTuple()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Traversable_traversableTuple()
		}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
		}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
		})})))}
	})
	return cache_Data_Semigroup_Traversable_traversableTuple
}

var cache_Data_Semigroup_Traversable_traversableIdentity gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableIdentity sync.Once

func Get_Data_Semigroup_Traversable_traversableIdentity() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableIdentity.Do(func() {
		cache_Data_Semigroup_Traversable_traversableIdentity = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableIdentity()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableIdentity()))}
		}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Identity_Identity(), v_1)
		}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Semigroup_Traversable_traversableIdentity
}

var cache_Data_Semigroup_Traversable_sequence1Default gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1Default sync.Once

func Get_Data_Semigroup_Traversable_sequence1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1Default.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_sequence1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dictTraversable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_1_box))
		})
	})
	return cache_Data_Semigroup_Traversable_sequence1Default
}

var cache_Data_Semigroup_Traversable_traversableDual gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableDual sync.Once

func Get_Data_Semigroup_Traversable_traversableDual() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableDual.Do(func() {
		cache_Data_Semigroup_Traversable_traversableDual = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableDual()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDual()))}
		}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_sequence1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_Semigroup_Traversable_traversableDual()), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))
		}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Semigroup_Traversable_traversableDual
}

var cache_Data_Semigroup_Traversable_traversableMultiplicative gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traversableMultiplicative sync.Once

func Get_Data_Semigroup_Traversable_traversableMultiplicative() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traversableMultiplicative.Do(func() {
		cache_Data_Semigroup_Traversable_traversableMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldableMultiplicative()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMultiplicative()))}
		}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_sequence1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_Semigroup_Traversable_traversableMultiplicative()), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))
		}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_1, v_2))
		})}))}
	})
	return cache_Data_Semigroup_Traversable_traversableMultiplicative
}

var cache_Data_Semigroup_Traversable_sequence1 gopurs_runtime.Value
var once_Data_Semigroup_Traversable_sequence1 sync.Once

func Get_Data_Semigroup_Traversable_sequence1() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_sequence1.Do(func() {
		cache_Data_Semigroup_Traversable_sequence1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_sequence1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Semigroup_Traversable_sequence1
}

var cache_Data_Semigroup_Traversable_traverse1Default gopurs_runtime.Value
var once_Data_Semigroup_Traversable_traverse1Default sync.Once

func Get_Data_Semigroup_Traversable_traverse1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Traversable_traverse1Default.Do(func() {
		cache_Data_Semigroup_Traversable_traverse1Default = gopurs_runtime.Func2(func(dictTraversable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Traversable_traverse1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](dictTraversable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_1_box))
		})
	})
	return cache_Data_Semigroup_Traversable_traverse1Default
}

type Constructor_Data_Semigroup_Traversable_Traversable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00530065006d006900670072006f00750070005f005400720061007600650072007300610062006c0065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00530065006d006900670072006f00750070005f005400720061007600650072007300610062006c0065005f005400720061007600650072007300610062006c00650031
func init() {
	gopurs_runtime.StructGetters[1596088409] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Foldable10":
			return gopurs_runtime.Box(c.V0)
		case "Traversable1":
			return gopurs_runtime.Box(c.V1)
		case "sequence1":
			return gopurs_runtime.Box(c.V2)
		case "traverse1":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Data_Semigroup_Traversable_Traversable1: " + key)
		}
	}
}

func Call_Data_Semigroup_Traversable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Semigroup_Traversable_Traversable1_dollar_Dict(x_0_loop struct {
	Foldable10   gopurs_runtime.Value
	Traversable1 gopurs_runtime.Value
	sequence1    gopurs_runtime.Value
	traverse1    gopurs_runtime.Value
	/*gopurs:record-v1:4:0046006f006c006400610062006c006500310030,005400720061007600650072007300610062006c00650031,00730065007100750065006e006300650031,007400720061007600650072007300650031*/
}) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
	var x_0 struct {
		Foldable10   gopurs_runtime.Value
		Traversable1 gopurs_runtime.Value
		sequence1    gopurs_runtime.Value
		traverse1    gopurs_runtime.Value
		/*gopurs:record-v1:4:0046006f006c006400610062006c006500310030,005400720061007600650072007300610062006c00650031,00730065007100750065006e006300650031,007400720061007600650072007300650031*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", orig.Foldable10, orig.Traversable1, orig.sequence1, orig.traverse1)
	}())
}

func Call_Data_Semigroup_Traversable_traverse1(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Data_Semigroup_Traversable_sequence1Default(dictTraversable1_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value], dictApply_1_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable1_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dictTraversable1_0_loop
	_ = dictTraversable1_0
	var dictApply_1 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_1_loop
	_ = dictApply_1
	return gopurs_runtime.Apply2(dictTraversable1_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}))
}

func Call_Data_Semigroup_Traversable_sequence1(dict_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Semigroup_Traversable_traverse1Default(dictTraversable1_0_loop *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value], dictApply_1_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictTraversable1_0 *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] = dictTraversable1_0_loop
	_ = dictTraversable1_0
	var dictApply_1 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_1_loop
	_ = dictApply_1
	// TAST (Let): sequence12_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t$scope72) [(TypeApp (TypeVar m$scope75) [(TypeVar b$scope74)])])] (TypeApp (TypeVar m$scope75) [(TypeApp (TypeVar t$scope72) [(TypeVar b$scope74)])]))
	sequence12_2_0 := gopurs_runtime.Apply(dictTraversable1_0.V2, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_1)})
	_ = sequence12_2_0
	return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(sequence12_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictTraversable1_0.V1, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), f_3, ta_4))
	})
}

func Rebox_Data_Semigroup_Traversable_1487937088_306175789(in *Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}
