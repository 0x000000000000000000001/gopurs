package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bitraversable_identity gopurs_runtime.Value
var once_Data_Bitraversable_identity sync.Once

func Get_Data_Bitraversable_identity() gopurs_runtime.Value {
	once_Data_Bitraversable_identity.Do(func() {
		cache_Data_Bitraversable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_identity(x_0_box)
		})
	})
	return cache_Data_Bitraversable_identity
}

var cache_Data_Bitraversable_identity1 gopurs_runtime.Value
var once_Data_Bitraversable_identity1 sync.Once

func Get_Data_Bitraversable_identity1() gopurs_runtime.Value {
	once_Data_Bitraversable_identity1.Do(func() {
		cache_Data_Bitraversable_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_identity1(x_0_box)
		})
	})
	return cache_Data_Bitraversable_identity1
}

var cache_Data_Bitraversable_Bitraversable_dollar_Dict gopurs_runtime.Value
var once_Data_Bitraversable_Bitraversable_dollar_Dict sync.Once

func Get_Data_Bitraversable_Bitraversable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bitraversable_Bitraversable_dollar_Dict.Do(func() {
		cache_Data_Bitraversable_Bitraversable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Call_Data_Bitraversable_Bitraversable_dollar_Dict(func() struct {
				Bifoldable1 gopurs_runtime.Value
				Bifunctor0  gopurs_runtime.Value
				bisequence  gopurs_runtime.Value
				bitraverse  gopurs_runtime.Value
				/*gopurs:record-v1:4:004200690066006f006c006400610062006c00650031,0042006900660075006e00630074006f00720030,0062006900730065007100750065006e00630065,0062006900740072006100760065007200730065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bifoldable1 gopurs_runtime.Value
					Bifunctor0  gopurs_runtime.Value
					bisequence  gopurs_runtime.Value
					bitraverse  gopurs_runtime.Value
					/*gopurs:record-v1:4:004200690066006f006c006400610062006c00650031,0042006900660075006e00630074006f00720030,0062006900730065007100750065006e00630065,0062006900740072006100760065007200730065*/
				}{}
				clone.Bifoldable1 = gopurs_runtime.RecordGet(orig, "Bifoldable1")
				clone.Bifunctor0 = gopurs_runtime.RecordGet(orig, "Bifunctor0")
				clone.bisequence = gopurs_runtime.RecordGet(orig, "bisequence")
				clone.bitraverse = gopurs_runtime.RecordGet(orig, "bitraverse")
				return clone
			}()))}
		})
	})
	return cache_Data_Bitraversable_Bitraversable_dollar_Dict
}

var cache_Data_Bitraversable_bitraverse gopurs_runtime.Value
var once_Data_Bitraversable_bitraverse sync.Once

func Get_Data_Bitraversable_bitraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverse.Do(func() {
		cache_Data_Bitraversable_bitraverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bitraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Bitraversable_bitraverse
}

var cache_Data_Bitraversable_lfor gopurs_runtime.Value
var once_Data_Bitraversable_lfor sync.Once

func Get_Data_Bitraversable_lfor() gopurs_runtime.Value {
	once_Data_Bitraversable_lfor.Do(func() {
		cache_Data_Bitraversable_lfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_lfor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_lfor
}

var cache_Data_Bitraversable_ltraverse gopurs_runtime.Value
var once_Data_Bitraversable_ltraverse sync.Once

func Get_Data_Bitraversable_ltraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_ltraverse.Do(func() {
		cache_Data_Bitraversable_ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_ltraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_ltraverse
}

var cache_Data_Bitraversable_rfor gopurs_runtime.Value
var once_Data_Bitraversable_rfor sync.Once

func Get_Data_Bitraversable_rfor() gopurs_runtime.Value {
	once_Data_Bitraversable_rfor.Do(func() {
		cache_Data_Bitraversable_rfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_rfor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_rfor
}

var cache_Data_Bitraversable_rtraverse gopurs_runtime.Value
var once_Data_Bitraversable_rtraverse sync.Once

func Get_Data_Bitraversable_rtraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_rtraverse.Do(func() {
		cache_Data_Bitraversable_rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_rtraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_rtraverse
}

var cache_Data_Bitraversable_bitraversableTuple gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableTuple sync.Once

func Get_Data_Bitraversable_bitraversableTuple() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableTuple.Do(func() {
		cache_Data_Bitraversable_bitraversableTuple = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2243670499_3561684974((&Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2812820739_3566843086(Rebox_Data_Bitraversable_3566843086_2812820739(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableTuple()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_1495429347_1688994542(Rebox_Data_Bitraversable_1688994542_1495429347(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorTuple()))))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
		}), gopurs_runtime.Func4(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Apply(g_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
		})})))}
	})
	return cache_Data_Bitraversable_bitraversableTuple
}

var cache_Data_Bitraversable_bitraversableJoker gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableJoker sync.Once

func Get_Data_Bitraversable_bitraversableJoker() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableJoker.Do(func() {
		cache_Data_Bitraversable_bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bitraversableJoker(dictTraversable_0_box)
		})
	})
	return cache_Data_Bitraversable_bitraversableJoker
}

var cache_Data_Bitraversable_bitraversableEither gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableEither sync.Once

func Get_Data_Bitraversable_bitraversableEither() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableEither.Do(func() {
		cache_Data_Bitraversable_bitraversableEither = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableEither()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorEither()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope110)])
			Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_1_0
			return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if v_2.Type == 9 && v_2.IntVal == 3711209382 {
						__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Either_Left(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
						goto end_branch_1
					} else {

					}
				}
				{
					if v_2.Type == 9 && v_2.IntVal == 2465973597 {
						__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Either_Right(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_1:
				return __t1
			})
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Functor0_1_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope101)])
			Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_1_2
			return gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v2_4.Type == 9 && v2_4.IntVal == 3711209382 {
						__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, Get_Data_Either_Left(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
						goto end_branch_3
					} else {

					}
				}
				{
					if v2_4.Type == 9 && v2_4.IntVal == 2465973597 {
						__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, Get_Data_Either_Right(), gopurs_runtime.Apply(v1_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			})
		})}))}
	})
	return cache_Data_Bitraversable_bitraversableEither
}

var cache_Data_Bitraversable_bitraversableConst gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableConst sync.Once

func Get_Data_Bitraversable_bitraversableConst() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableConst.Do(func() {
		cache_Data_Bitraversable_bitraversableConst = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableConst()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorConst()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Const_Const(), v_1)
		}), gopurs_runtime.Func4(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Const_Const(), gopurs_runtime.Apply(f_1, v1_3))
		})}))}
	})
	return cache_Data_Bitraversable_bitraversableConst
}

var cache_Data_Bitraversable_bitraversableClown gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableClown sync.Once

func Get_Data_Bitraversable_bitraversableClown() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableClown.Do(func() {
		cache_Data_Bitraversable_bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bitraversableClown(dictTraversable_0_box)
		})
	})
	return cache_Data_Bitraversable_bitraversableClown
}

var cache_Data_Bitraversable_bisequenceDefault gopurs_runtime.Value
var once_Data_Bitraversable_bisequenceDefault sync.Once

func Get_Data_Bitraversable_bisequenceDefault() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequenceDefault.Do(func() {
		cache_Data_Bitraversable_bisequenceDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bisequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_bisequenceDefault
}

var cache_Data_Bitraversable_bisequence gopurs_runtime.Value
var once_Data_Bitraversable_bisequence sync.Once

func Get_Data_Bitraversable_bisequence() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequence.Do(func() {
		cache_Data_Bitraversable_bisequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bisequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Bitraversable_bisequence
}

var cache_Data_Bitraversable_bitraversableFlip gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableFlip sync.Once

func Get_Data_Bitraversable_bitraversableFlip() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableFlip.Do(func() {
		cache_Data_Bitraversable_bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bitraversableFlip(dictBitraversable_0_box)
		})
	})
	return cache_Data_Bitraversable_bitraversableFlip
}

var cache_Data_Bitraversable_bitraversableProduct2 gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableProduct2 sync.Once

func Get_Data_Bitraversable_bitraversableProduct2() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableProduct2.Do(func() {
		cache_Data_Bitraversable_bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bitraversableProduct2(dictBitraversable_0_box)
		})
	})
	return cache_Data_Bitraversable_bitraversableProduct2
}

var cache_Data_Bitraversable_bitraverseDefault gopurs_runtime.Value
var once_Data_Bitraversable_bitraverseDefault sync.Once

func Get_Data_Bitraversable_bitraverseDefault() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverseDefault.Do(func() {
		cache_Data_Bitraversable_bitraverseDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bitraverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_bitraverseDefault
}

var cache_Data_Bitraversable_bifor gopurs_runtime.Value
var once_Data_Bitraversable_bifor sync.Once

func Get_Data_Bitraversable_bifor() gopurs_runtime.Value {
	once_Data_Bitraversable_bifor.Do(func() {
		cache_Data_Bitraversable_bifor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bifor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
		})
	})
	return cache_Data_Bitraversable_bifor
}

type Constructor_Data_Bitraversable_Bitraversable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00420069007400720061007600650072007300610062006c0065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00420069007400720061007600650072007300610062006c0065005f00420069007400720061007600650072007300610062006c0065
func init() {
	gopurs_runtime.StructGetters[3704227322] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Bifoldable1":
			return gopurs_runtime.Box(c.V0)
		case "Bifunctor0":
			return gopurs_runtime.Box(c.V1)
		case "bisequence":
			return gopurs_runtime.Box(c.V2)
		case "bitraverse":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Data_Bitraversable_Bitraversable: " + key)
		}
	}
}

func Call_Data_Bitraversable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Bitraversable_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Bitraversable_Bitraversable_dollar_Dict(x_0_loop struct {
	Bifoldable1 gopurs_runtime.Value
	Bifunctor0  gopurs_runtime.Value
	bisequence  gopurs_runtime.Value
	bitraverse  gopurs_runtime.Value
	/*gopurs:record-v1:4:004200690066006f006c006400610062006c00650031,0042006900660075006e00630074006f00720030,0062006900730065007100750065006e00630065,0062006900740072006100760065007200730065*/
}) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	var x_0 struct {
		Bifoldable1 gopurs_runtime.Value
		Bifunctor0  gopurs_runtime.Value
		bisequence  gopurs_runtime.Value
		bitraverse  gopurs_runtime.Value
		/*gopurs:record-v1:4:004200690066006f006c006400610062006c00650031,0042006900660075006e00630074006f00720030,0062006900730065007100750065006e00630065,0062006900740072006100760065007200730065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", orig.Bifoldable1, orig.Bifunctor0, orig.bisequence, orig.bitraverse)
	}())
}

func Call_Data_Bitraversable_bitraverse(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Data_Bitraversable_lfor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	// TAST (Let): bitraverse2_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope14)] (TypeApp (TypeVar f$scope15) [(TypeVar c$scope13)])), (Func [(TypeVar b$scope12)] (TypeApp (TypeVar f$scope15) [(TypeVar b$scope12)])), (TypeApp (TypeVar t$scope11) [(TypeVar a$scope14), (TypeVar b$scope12)])] (TypeApp (TypeVar f$scope15) [(TypeApp (TypeVar t$scope11) [(TypeVar c$scope13), (TypeVar b$scope12)])]))
	bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
	_ = bitraverse2_2_0
	// TAST (Let): pure_3_1 shape=App(Var) bindingType=(Func [(TypeVar b$scope12)] (TypeApp (TypeVar f$scope15) [(TypeVar b$scope12)]))
	pure_3_1 := Call_Control_Applicative_pure(dictApplicative_1)
	_ = pure_3_1
	return gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(bitraverse2_2_0, f_5, pure_3_1, t_4)
	})
}

func Call_Data_Bitraversable_ltraverse(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	// TAST (Let): bitraverse2_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope24)] (TypeApp (TypeVar f$scope25) [(TypeVar c$scope23)])), (Func [(TypeVar b$scope22)] (TypeApp (TypeVar f$scope25) [(TypeVar b$scope22)])), (TypeApp (TypeVar t$scope21) [(TypeVar a$scope24), (TypeVar b$scope22)])] (TypeApp (TypeVar f$scope25) [(TypeApp (TypeVar t$scope21) [(TypeVar c$scope23), (TypeVar b$scope22)])]))
	bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
	_ = bitraverse2_2_0
	// TAST (Let): pure_3_1 shape=App(Var) bindingType=(Func [(TypeVar b$scope22)] (TypeApp (TypeVar f$scope25) [(TypeVar b$scope22)]))
	pure_3_1 := Call_Control_Applicative_pure(dictApplicative_1)
	_ = pure_3_1
	return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(bitraverse2_2_0, f_4, pure_3_1)
	})
}

func Call_Data_Bitraversable_rfor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	// TAST (Let): bitraverse2_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope34)] (TypeApp (TypeVar f$scope35) [(TypeVar a$scope34)])), (Func [(TypeVar b$scope32)] (TypeApp (TypeVar f$scope35) [(TypeVar c$scope33)])), (TypeApp (TypeVar t$scope31) [(TypeVar a$scope34), (TypeVar b$scope32)])] (TypeApp (TypeVar f$scope35) [(TypeApp (TypeVar t$scope31) [(TypeVar a$scope34), (TypeVar c$scope33)])]))
	bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
	_ = bitraverse2_2_0
	// TAST (Let): pure_3_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope34)] (TypeApp (TypeVar f$scope35) [(TypeVar a$scope34)]))
	pure_3_1 := Call_Control_Applicative_pure(dictApplicative_1)
	_ = pure_3_1
	return gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(bitraverse2_2_0, pure_3_1, f_5, t_4)
	})
}

func Call_Data_Bitraversable_rtraverse(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	return gopurs_runtime.Apply2(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, Call_Control_Applicative_pure(dictApplicative_1))
}

func Call_Data_Bitraversable_bitraversableJoker(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): bifunctorJoker_1_0 shape=App(Var) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeVar f$scope68) [(TypeVar b)])])
	bifunctorJoker_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Call_Data_Functor_Joker_bifunctorJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})))
	_ = bifunctorJoker_1_0
	// TAST (Let): bifoldableJoker_2_1 shape=App(Var) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(TypeApp (TypeVar f$scope68) [(TypeVar b)])])
	bifoldableJoker_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Call_Data_Bifoldable_bifoldableJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})))
	_ = bifoldableJoker_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableJoker_2_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorJoker_1_0)}
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): sequence1_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope68) [(TypeApp (TypeVar f$scope86) [(TypeVar b$scope88)])])] (TypeApp (TypeVar f$scope86) [(TypeApp (TypeVar f$scope68) [(TypeVar b$scope88)])]))
		sequence1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = sequence1_4_2
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Joker_Joker(), gopurs_runtime.Apply(sequence1_4_2, v_5))
		})
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse1_4_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar b$scope80)] (TypeApp (TypeVar f$scope78) [(TypeVar d$scope82)])), (TypeApp (TypeVar f$scope68) [(TypeVar b$scope80)])] (TypeApp (TypeVar f$scope78) [(TypeApp (TypeVar f$scope68) [(TypeVar d$scope82)])]))
		traverse1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = traverse1_4_3
		return gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Joker_Joker(), gopurs_runtime.Apply2(traverse1_4_3, r_6, v1_7))
		})
	})}))}
}

func Call_Data_Bitraversable_bitraversableClown(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
	_ = dictTraversable_0
	// TAST (Let): bifunctorClown_1_0 shape=App(Var) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeVar f$scope141) [(TypeVar a)])])
	bifunctorClown_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Call_Data_Functor_Clown_bifunctorClown(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})))
	_ = bifunctorClown_1_0
	// TAST (Let): bifoldableClown_2_1 shape=App(Var) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(TypeApp (TypeVar f$scope141) [(TypeVar a)])])
	bifoldableClown_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Call_Data_Bifoldable_bifoldableClown(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})))
	_ = bifoldableClown_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableClown_2_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorClown_1_0)}
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): sequence1_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope141) [(TypeApp (TypeVar f$scope159) [(TypeVar a$scope160)])])] (TypeApp (TypeVar f$scope159) [(TypeApp (TypeVar f$scope141) [(TypeVar a$scope160)])]))
		sequence1_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = sequence1_4_2
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Clown_Clown(), gopurs_runtime.Apply(sequence1_4_2, v_5))
		})
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): traverse1_4_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope152)] (TypeApp (TypeVar f$scope151) [(TypeVar c$scope154)])), (TypeApp (TypeVar f$scope141) [(TypeVar a$scope152)])] (TypeApp (TypeVar f$scope151) [(TypeApp (TypeVar f$scope141) [(TypeVar c$scope154)])]))
		traverse1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = traverse1_4_3
		return gopurs_runtime.Func3(func(l_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Clown_Clown(), gopurs_runtime.Apply2(traverse1_4_3, l_5, v1_7))
		})
	})}))}
}

func Call_Data_Bitraversable_bisequenceDefault(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	return gopurs_runtime.Apply3(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}))
}

func Call_Data_Bitraversable_bisequence(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Bitraversable_bitraversableFlip(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
	_ = dictBitraversable_0
	// TAST (Let): bifunctorFlip_1_0 shape=App(Var) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p$scope180)])])
	bifunctorFlip_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Call_Data_Functor_Flip_bifunctorFlip(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})))
	_ = bifunctorFlip_1_0
	// TAST (Let): bifoldableFlip_2_1 shape=App(Var) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p$scope180)])])
	bifoldableFlip_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Call_Data_Bifoldable_bifoldableFlip(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})))
	_ = bifoldableFlip_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableFlip_2_1)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorFlip_1_0)}
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): bisequence2_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar p$scope180) [(TypeApp (TypeVar f$scope200) [(TypeVar b$scope202)]), (TypeApp (TypeVar f$scope200) [(TypeVar a$scope201)])])] (TypeApp (TypeVar f$scope200) [(TypeApp (TypeVar p$scope180) [(TypeVar b$scope202), (TypeVar a$scope201)])]))
		bisequence2_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = bisequence2_4_2
		return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Flip_Flip(), gopurs_runtime.Apply(bisequence2_4_2, v_5))
		})
	}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): bitraverse2_4_3 shape=App(Other) bindingType=(Func [(Func [(TypeVar b$scope192)] (TypeApp (TypeVar f$scope190) [(TypeVar d$scope194)])), (Func [(TypeVar a$scope191)] (TypeApp (TypeVar f$scope190) [(TypeVar c$scope193)])), (TypeApp (TypeVar p$scope180) [(TypeVar b$scope192), (TypeVar a$scope191)])] (TypeApp (TypeVar f$scope190) [(TypeApp (TypeVar p$scope180) [(TypeVar d$scope194), (TypeVar c$scope193)])]))
		bitraverse2_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))})
		_ = bitraverse2_4_3
		return gopurs_runtime.Func3(func(r_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Flip_Flip(), gopurs_runtime.Apply3(bitraverse2_4_3, l_6, r_5, v_7))
		})
	})}))}
}

func Call_Data_Bitraversable_bitraversableProduct2(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
	_ = dictBitraversable_0
	// TAST (Let): bifunctorProduct2_1_0 shape=App(Var) bindingType=Any
	bifunctorProduct2_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product2_bifunctorProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{}))
	_ = bifunctorProduct2_1_0
	// TAST (Let): bifoldableProduct2_2_1 shape=App(Var) bindingType=Any
	bifoldableProduct2_2_1 := gopurs_runtime.Apply(Get_Data_Bifoldable_bifoldableProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{}))
	_ = bifoldableProduct2_2_1
	return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): bifunctorProduct21_4_2 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f$scope210), (TypeVar g$scope211)])])
		bifunctorProduct21_4_2 := Rebox_Data_Bitraversable_1688994542_2864580459(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(bifunctorProduct2_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifunctor0"), gopurs_runtime.Value{}))))
		_ = bifunctorProduct21_4_2
		// TAST (Let): bifoldableProduct21_5_3 shape=App(Other) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f$scope210), (TypeVar g$scope211)])])
		bifoldableProduct21_5_3 := Rebox_Data_Bitraversable_3566843086_4280144779(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](gopurs_runtime.Apply(bifoldableProduct2_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifoldable1"), gopurs_runtime.Value{}))))
		_ = bifoldableProduct21_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_4211574891_3561684974((&Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_4280144779_3566843086(bifoldableProduct21_5_3))}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2864580459_1688994542(bifunctorProduct21_4_2))}
		}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): bisequence3_7_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope210) [(TypeApp (TypeVar f$scope230) [(TypeVar a$scope231)]), (TypeApp (TypeVar f$scope230) [(TypeVar b$scope232)])])] (TypeApp (TypeVar f$scope230) [(TypeApp (TypeVar f$scope210) [(TypeVar a$scope231), (TypeVar b$scope232)])]))
			bisequence3_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = bisequence3_7_4
			// TAST (Let): bisequence4_8_5 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g$scope211) [(TypeApp (TypeVar f$scope230) [(TypeVar a$scope231)]), (TypeApp (TypeVar f$scope230) [(TypeVar b$scope232)])])] (TypeApp (TypeVar f$scope230) [(TypeApp (TypeVar g$scope211) [(TypeVar a$scope231), (TypeVar b$scope232)])]))
			bisequence4_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = bisequence4_8_5
			return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product2_Product2(), gopurs_runtime.Apply(bisequence3_7_4, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply(bisequence4_8_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
			})
		}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): bitraverse3_7_6 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope222)] (TypeApp (TypeVar f$scope221) [(TypeVar c$scope224)])), (Func [(TypeVar b$scope223)] (TypeApp (TypeVar f$scope221) [(TypeVar d$scope225)])), (TypeApp (TypeVar f$scope210) [(TypeVar a$scope222), (TypeVar b$scope223)])] (TypeApp (TypeVar f$scope221) [(TypeApp (TypeVar f$scope210) [(TypeVar c$scope224), (TypeVar d$scope225)])]))
			bitraverse3_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = bitraverse3_7_6
			// TAST (Let): bitraverse4_8_7 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope222)] (TypeApp (TypeVar f$scope221) [(TypeVar c$scope224)])), (Func [(TypeVar b$scope223)] (TypeApp (TypeVar f$scope221) [(TypeVar d$scope225)])), (TypeApp (TypeVar g$scope211) [(TypeVar a$scope222), (TypeVar b$scope223)])] (TypeApp (TypeVar f$scope221) [(TypeApp (TypeVar g$scope211) [(TypeVar c$scope224), (TypeVar d$scope225)])]))
			bitraverse4_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))})
			_ = bitraverse4_8_7
			return gopurs_runtime.Func3(func(l_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product2_Product2(), gopurs_runtime.Apply3(bitraverse3_7_6, l_9, r_10, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply3(bitraverse4_8_7, l_9, r_10, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
			})
		})})))}
	})
}

func Call_Data_Bitraversable_bitraverseDefault(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	// TAST (Let): bisequence2_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t$scope243) [(TypeApp (TypeVar f$scope244) [(TypeVar c$scope247)]), (TypeApp (TypeVar f$scope244) [(TypeVar d$scope248)])])] (TypeApp (TypeVar f$scope244) [(TypeApp (TypeVar t$scope243) [(TypeVar c$scope247), (TypeVar d$scope248)])]))
	bisequence2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
	_ = bisequence2_2_0
	return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(bisequence2_2_0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictBitraversable_0.V1, gopurs_runtime.Value{}), "bimap"), f_3, g_4, t_5))
	})
}

func Call_Data_Bitraversable_bifor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
	_ = dictBitraversable_0
	var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
	_ = dictApplicative_1
	// TAST (Let): bitraverse2_2_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope257)] (TypeApp (TypeVar f$scope256) [(TypeVar c$scope259)])), (Func [(TypeVar b$scope258)] (TypeApp (TypeVar f$scope256) [(TypeVar d$scope260)])), (TypeApp (TypeVar t$scope255) [(TypeVar a$scope257), (TypeVar b$scope258)])] (TypeApp (TypeVar f$scope256) [(TypeApp (TypeVar t$scope255) [(TypeVar c$scope259), (TypeVar d$scope260)])]))
	bitraverse2_2_0 := gopurs_runtime.Apply(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)})
	_ = bitraverse2_2_0
	return gopurs_runtime.Func3(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(bitraverse2_2_0, f_4, g_5, t_3)
	})
}

func Rebox_Data_Bitraversable_1495429347_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_1688994542_1495429347(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_1688994542_2864580459(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_2243670499_3561684974(in *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_2812820739_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_2864580459_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_3566843086_2812820739(in *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_3566843086_4280144779(in *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_4211574891_3561684974(in *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bitraversable_4280144779_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}
