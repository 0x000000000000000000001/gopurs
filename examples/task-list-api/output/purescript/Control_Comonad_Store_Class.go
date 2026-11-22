package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Store_Class_compose gopurs_runtime.Value
var once_Control_Comonad_Store_Class_compose sync.Once

func Get_Control_Comonad_Store_Class_compose() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_compose.Do(func() {
		cache_Control_Comonad_Store_Class_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Class_compose
}

var cache_Control_Comonad_Store_Class_compose1 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_compose1 sync.Once

func Get_Control_Comonad_Store_Class_compose1() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_compose1.Do(func() {
		cache_Control_Comonad_Store_Class_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Class_compose1
}

var cache_Control_Comonad_Store_Class_compose2 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_compose2 sync.Once

func Get_Control_Comonad_Store_Class_compose2() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_compose2.Do(func() {
		cache_Control_Comonad_Store_Class_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Class_compose2
}

var cache_Control_Comonad_Store_Class_compose3 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_compose3 sync.Once

func Get_Control_Comonad_Store_Class_compose3() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_compose3.Do(func() {
		cache_Control_Comonad_Store_Class_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Class_compose3
}

var cache_Control_Comonad_Store_Class_compose4 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_compose4 sync.Once

func Get_Control_Comonad_Store_Class_compose4() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_compose4.Do(func() {
		cache_Control_Comonad_Store_Class_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Class_compose4
}

var cache_Control_Comonad_Store_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Store_Class_lower sync.Once

func Get_Control_Comonad_Store_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_lower.Do(func() {
		cache_Control_Comonad_Store_Class_lower = gopurs_runtime.RecordGet(Get_Control_Comonad_Env_Trans_comonadTransEnvT(), "lower")
	})
	return cache_Control_Comonad_Store_Class_lower
}

var cache_Control_Comonad_Store_Class_compose5 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_compose5 sync.Once

func Get_Control_Comonad_Store_Class_compose5() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_compose5.Do(func() {
		cache_Control_Comonad_Store_Class_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Store_Class_compose5
}

var cache_Control_Comonad_Store_Class_ComonadStore_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Store_Class_ComonadStore_dollar_Dict sync.Once

func Get_Control_Comonad_Store_Class_ComonadStore_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_ComonadStore_dollar_Dict.Do(func() {
		cache_Control_Comonad_Store_Class_ComonadStore_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Store_Class_ComonadStore_dollar_Dict(func() struct {
				Comonad0 gopurs_runtime.Value
				peek     gopurs_runtime.Value
				pos      gopurs_runtime.Value
				/*gopurs:record-v1:3:0043006f006d006f006e006100640030,007000650065006b,0070006f0073*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Comonad0 gopurs_runtime.Value
					peek     gopurs_runtime.Value
					pos      gopurs_runtime.Value
					/*gopurs:record-v1:3:0043006f006d006f006e006100640030,007000650065006b,0070006f0073*/
				}{}
				clone.Comonad0 = gopurs_runtime.RecordGet(orig, "Comonad0")
				clone.peek = gopurs_runtime.RecordGet(orig, "peek")
				clone.pos = gopurs_runtime.RecordGet(orig, "pos")
				return clone
			}()))}
		})
	})
	return cache_Control_Comonad_Store_Class_ComonadStore_dollar_Dict
}

var cache_Control_Comonad_Store_Class_pos gopurs_runtime.Value
var once_Control_Comonad_Store_Class_pos sync.Once

func Get_Control_Comonad_Store_Class_pos() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_pos.Do(func() {
		cache_Control_Comonad_Store_Class_pos = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_pos(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Comonad_Store_Class_pos
}

var cache_Control_Comonad_Store_Class_peek gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peek sync.Once

func Get_Control_Comonad_Store_Class_peek() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peek.Do(func() {
		cache_Control_Comonad_Store_Class_peek = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_peek(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Comonad_Store_Class_peek
}

var cache_Control_Comonad_Store_Class_peeks gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peeks sync.Once

func Get_Control_Comonad_Store_Class_peeks() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peeks.Do(func() {
		cache_Control_Comonad_Store_Class_peeks = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_peeks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), f_1_box, x_2_box)
		})
	})
	return cache_Control_Comonad_Store_Class_peeks
}

var cache_Control_Comonad_Store_Class_seeks gopurs_runtime.Value
var once_Control_Comonad_Store_Class_seeks sync.Once

func Get_Control_Comonad_Store_Class_seeks() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_seeks.Do(func() {
		cache_Control_Comonad_Store_Class_seeks = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_seeks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
		})
	})
	return cache_Control_Comonad_Store_Class_seeks
}

var cache_Control_Comonad_Store_Class_seek gopurs_runtime.Value
var once_Control_Comonad_Store_Class_seek sync.Once

func Get_Control_Comonad_Store_Class_seek() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_seek.Do(func() {
		cache_Control_Comonad_Store_Class_seek = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_seek(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
		})
	})
	return cache_Control_Comonad_Store_Class_seek
}

var cache_Control_Comonad_Store_Class_experiment gopurs_runtime.Value
var once_Control_Comonad_Store_Class_experiment sync.Once

func Get_Control_Comonad_Store_Class_experiment() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_experiment.Do(func() {
		cache_Control_Comonad_Store_Class_experiment = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_experiment(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
		})
	})
	return cache_Control_Comonad_Store_Class_experiment
}

var cache_Control_Comonad_Store_Class_comonadStoreTracedT gopurs_runtime.Value
var once_Control_Comonad_Store_Class_comonadStoreTracedT sync.Once

func Get_Control_Comonad_Store_Class_comonadStoreTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_comonadStoreTracedT.Do(func() {
		cache_Control_Comonad_Store_Class_comonadStoreTracedT = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_comonadStoreTracedT(dictComonadStore_0_box)
		})
	})
	return cache_Control_Comonad_Store_Class_comonadStoreTracedT
}

var cache_Control_Comonad_Store_Class_comonadStoreStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Class_comonadStoreStoreT sync.Once

func Get_Control_Comonad_Store_Class_comonadStoreStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_comonadStoreStoreT.Do(func() {
		cache_Control_Comonad_Store_Class_comonadStoreStoreT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_comonadStoreStoreT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Store_Class_comonadStoreStoreT
}

var cache_Control_Comonad_Store_Class_comonadStoreEnvT gopurs_runtime.Value
var once_Control_Comonad_Store_Class_comonadStoreEnvT sync.Once

func Get_Control_Comonad_Store_Class_comonadStoreEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_comonadStoreEnvT.Do(func() {
		cache_Control_Comonad_Store_Class_comonadStoreEnvT = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Store_Class_comonadStoreEnvT(dictComonadStore_0_box)
		})
	})
	return cache_Control_Comonad_Store_Class_comonadStoreEnvT
}

type Constructor_Control_Comonad_Store_Class_ComonadStore[T_s any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f00530074006f00720065005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f00530074006f00720065005f0043006c006100730073005f0043006f006d006f006e0061006400530074006f00720065
func init() {
	gopurs_runtime.StructGetters[4264314723] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Comonad0":
			return gopurs_runtime.Box(c.V0)
		case "peek":
			return gopurs_runtime.Box(c.V1)
		case "pos":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Control_Comonad_Store_Class_ComonadStore: " + key)
		}
	}
}

func Call_Control_Comonad_Store_Class_ComonadStore_dollar_Dict(x_0_loop struct {
	Comonad0 gopurs_runtime.Value
	peek     gopurs_runtime.Value
	pos      gopurs_runtime.Value
	/*gopurs:record-v1:3:0043006f006d006f006e006100640030,007000650065006b,0070006f0073*/
}) *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Comonad0 gopurs_runtime.Value
		peek     gopurs_runtime.Value
		pos      gopurs_runtime.Value
		/*gopurs:record-v1:3:0043006f006d006f006e006100640030,007000650065006b,0070006f0073*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", orig.Comonad0, orig.peek, orig.pos)
	}())
}

func Call_Control_Comonad_Store_Class_pos(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Control_Comonad_Store_Class_peek(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Comonad_Store_Class_peeks(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
	_ = dictComonadStore_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	return gopurs_runtime.Apply2(dictComonadStore_0.V1, gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(dictComonadStore_0.V2, x_2)), x_2)
}

func Call_Control_Comonad_Store_Class_seeks(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
	_ = dictComonadStore_0
	// TAST (Let): duplicate_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar w$scope17) [(TypeVar a$scope16)])] (TypeApp (TypeVar w$scope17) [(TypeApp (TypeVar w$scope17) [(TypeVar a$scope16)])]))
	duplicate_1_0 := Call_Control_Extend_duplicate(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictComonadStore_0.V0, gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{})))
	_ = duplicate_1_0
	return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(dictComonadStore_0.V1, gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(dictComonadStore_0.V2, x_3)), x_3)
		}), duplicate_1_0)
	})
}

func Call_Control_Comonad_Store_Class_seek(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
	_ = dictComonadStore_0
	// TAST (Let): duplicate_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar w$scope24) [(TypeVar a$scope23)])] (TypeApp (TypeVar w$scope24) [(TypeApp (TypeVar w$scope24) [(TypeVar a$scope23)])]))
	duplicate_1_0 := Call_Control_Extend_duplicate(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictComonadStore_0.V0, gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{})))
	_ = duplicate_1_0
	return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(dictComonadStore_0.V1, s_2), duplicate_1_0)
	})
}

func Call_Control_Comonad_Store_Class_experiment(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
	_ = dictComonadStore_0
	// TAST (Let): peek1_1_0 shape=App(Var) bindingType=(Func [(TypeVar s$scope33), (TypeApp (TypeVar w$scope32) [(TypeVar a$scope31)])] (TypeVar a$scope31))
	peek1_1_0 := Call_Control_Comonad_Store_Class_peek(dictComonadStore_0)
	_ = peek1_1_0
	return gopurs_runtime.Func3(func(dictFunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_2, "map"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(peek1_1_0, a_5, x_4)
		}), gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(dictComonadStore_0.V2, x_4)))
	})
}

func Call_Control_Comonad_Store_Class_comonadStoreTracedT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
	_ = dictComonadStore_0
	// TAST (Let): pos1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar w$scope38) [(TypeVar a$scope44)])] (TypeVar s$scope37))
	pos1_1_0 := Call_Control_Comonad_Store_Class_pos(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0))
	_ = pos1_1_0
	// TAST (Let): Comonad0_2_1 shape=App(Other) bindingType=Any
	Comonad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
	_ = Comonad0_2_1
	// TAST (Let): comonadTracedT_3_2 shape=App(Var) bindingType=Any
	comonadTracedT_3_2 := Call_Control_Comonad_Traced_Trans_comonadTracedT(Comonad0_2_1)
	_ = comonadTracedT_3_2
	return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): lower1_5_3 shape=App(Var) bindingType=Any
		lower1_5_3 := Call_Control_Comonad_Trans_Class_lower(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_4)))
		_ = lower1_5_3
		// TAST (Let): lower2_6_4 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar w$scope38) [(Func [(TypeVar m$scope39)] (TypeVar a$scope46))])] (TypeApp (TypeVar w$scope38) [(TypeVar a$scope46)]))
		lower2_6_4 := gopurs_runtime.Apply(lower1_5_3, Comonad0_2_1)
		_ = lower2_6_4
		// TAST (Let): comonadTracedT1_7_5 shape=App(Other) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeVar w$scope38) [(Func [(TypeVar m$scope39)] (TypeVar a))])])
		comonadTracedT1_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(comonadTracedT_3_2, dictMonoid_4))
		_ = comonadTracedT1_7_5
		return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_7_5)}
		}), gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_8), lower2_6_4)
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), pos1_1_0, gopurs_runtime.Apply(lower1_5_3, Comonad0_2_1))}))}
	})
}

func Call_Control_Comonad_Store_Class_comonadStoreStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): comonadStoreT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w$scope50) [(Func [(TypeVar s$scope51)] (TypeVar a))]), (TypeVar s$scope51)])])
	comonadStoreT_1_0 := Rebox_Control_Comonad_Store_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Comonad_Store_Trans_comonadStoreT(dictComonad_0)))
	_ = comonadStoreT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_229239260_3433217233((&Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_3056445460_2550391993(comonadStoreT_1_0))}
	}), gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_2)
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
	})})))}
}

func Call_Control_Comonad_Store_Class_comonadStoreEnvT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
	_ = dictComonadStore_0
	// TAST (Let): comonadEnvT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope63), (TypeVar w$scope62)])])
	comonadEnvT_1_0 := Rebox_Control_Comonad_Store_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_comonadEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}))))
	_ = comonadEnvT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_229239260_3433217233((&Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_3056445460_2550391993(comonadEnvT_1_0))}
	}), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_2), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1
		}))
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Comonad_Store_Class_pos(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0)), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
	}))})))}
}

func Rebox_Control_Comonad_Store_Class_229239260_3433217233(in *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Class_2550391993_3056445460(in *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Store_Class_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(unsafe.Pointer(in))
}
