package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Traced_Class_compose gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_compose sync.Once

func Get_Control_Comonad_Traced_Class_compose() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_compose.Do(func() {
		cache_Control_Comonad_Traced_Class_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Traced_Class_compose
}

var cache_Control_Comonad_Traced_Class_composeFlipped gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_composeFlipped sync.Once

func Get_Control_Comonad_Traced_Class_composeFlipped() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_composeFlipped.Do(func() {
		cache_Control_Comonad_Traced_Class_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Traced_Class_composeFlipped
}

var cache_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict sync.Once

func Get_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict.Do(func() {
		cache_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 433046755, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict(func() struct {
				Comonad0 gopurs_runtime.Value
				track    gopurs_runtime.Value
				/*gopurs:record-v1:2:0043006f006d006f006e006100640030,0074007200610063006b*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Comonad0 gopurs_runtime.Value
					track    gopurs_runtime.Value
					/*gopurs:record-v1:2:0043006f006d006f006e006100640030,0074007200610063006b*/
				}{}
				clone.Comonad0 = gopurs_runtime.RecordGet(orig, "Comonad0")
				clone.track = gopurs_runtime.RecordGet(orig, "track")
				return clone
			}()))}
		})
	})
	return cache_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict
}

var cache_Control_Comonad_Traced_Class_track gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_track sync.Once

func Get_Control_Comonad_Traced_Class_track() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_track.Do(func() {
		cache_Control_Comonad_Traced_Class_track = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_track(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Comonad_Traced_Class_track
}

var cache_Control_Comonad_Traced_Class_tracks gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_tracks sync.Once

func Get_Control_Comonad_Traced_Class_tracks() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_tracks.Do(func() {
		cache_Control_Comonad_Traced_Class_tracks = gopurs_runtime.Func3(func(dictComonadTraced_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, w_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_tracks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadTraced_0_box), f_1_box, w_2_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_tracks
}

var cache_Control_Comonad_Traced_Class_lowerTrack gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack sync.Once

func Get_Control_Comonad_Traced_Class_lowerTrack() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack = gopurs_runtime.Func(func(dictComonadTrans_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](dictComonadTrans_0_box))
		})
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack
}

var cache_Control_Comonad_Traced_Class_lowerTrack1 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack1 sync.Once

func Get_Control_Comonad_Traced_Class_lowerTrack1() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack1.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack1 = Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Get_Control_Comonad_Store_Trans_comonadTransStoreT()))
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack1
}

var cache_Control_Comonad_Traced_Class_lowerTrack2 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack2 sync.Once

func Get_Control_Comonad_Traced_Class_lowerTrack2() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack2.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack2 = Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Get_Control_Comonad_Trans_Class_comonadTransIdentityT()))
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack2
}

var cache_Control_Comonad_Traced_Class_lowerTrack3 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack3 sync.Once

func Get_Control_Comonad_Traced_Class_lowerTrack3() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack3.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack3 = Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Get_Control_Comonad_Env_Trans_comonadTransEnvT()))
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack3
}

var cache_Control_Comonad_Traced_Class_listens gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_listens sync.Once

func Get_Control_Comonad_Traced_Class_listens() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_listens.Do(func() {
		cache_Control_Comonad_Traced_Class_listens = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_listens(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_listens
}

var cache_Control_Comonad_Traced_Class_listen gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_listen sync.Once

func Get_Control_Comonad_Traced_Class_listen() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_listen.Do(func() {
		cache_Control_Comonad_Traced_Class_listen = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_listen(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_listen
}

var cache_Control_Comonad_Traced_Class_comonadTracedTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedTracedT sync.Once

func Get_Control_Comonad_Traced_Class_comonadTracedTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedTracedT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_comonadTracedTracedT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedTracedT
}

var cache_Control_Comonad_Traced_Class_comonadTracedStoreT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedStoreT sync.Once

func Get_Control_Comonad_Traced_Class_comonadTracedStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedStoreT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedStoreT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_comonadTracedStoreT(dictComonadTraced_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedStoreT
}

var cache_Control_Comonad_Traced_Class_comonadTracedIdentityT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedIdentityT sync.Once

func Get_Control_Comonad_Traced_Class_comonadTracedIdentityT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedIdentityT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedIdentityT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_comonadTracedIdentityT(dictComonadTraced_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedIdentityT
}

var cache_Control_Comonad_Traced_Class_comonadTracedEnvT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedEnvT sync.Once

func Get_Control_Comonad_Traced_Class_comonadTracedEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedEnvT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedEnvT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_comonadTracedEnvT(dictComonadTraced_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedEnvT
}

var cache_Control_Comonad_Traced_Class_censor gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_censor sync.Once

func Get_Control_Comonad_Traced_Class_censor() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_censor.Do(func() {
		cache_Control_Comonad_Traced_Class_censor = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Class_censor(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, v_2_box)
		})
	})
	return cache_Control_Comonad_Traced_Class_censor
}

type Constructor_Control_Comonad_Traced_Class_ComonadTraced[T_t any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f005400720061006300650064005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f005400720061006300650064005f0043006c006100730073005f0043006f006d006f006e00610064005400720061006300650064
func init() {
	gopurs_runtime.StructGetters[433046755] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Comonad0":
			return gopurs_runtime.Box(c.V0)
		case "track":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Comonad_Traced_Class_ComonadTraced: " + key)
		}
	}
}

func Call_Control_Comonad_Traced_Class_ComonadTraced_dollar_Dict(x_0_loop struct {
	Comonad0 gopurs_runtime.Value
	track    gopurs_runtime.Value
	/*gopurs:record-v1:2:0043006f006d006f006e006100640030,0074007200610063006b*/
}) *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Comonad0 gopurs_runtime.Value
		track    gopurs_runtime.Value
		/*gopurs:record-v1:2:0043006f006d006f006e006100640030,0074007200610063006b*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Comonad0", "track", orig.Comonad0, orig.track)
	}())
}

func Call_Control_Comonad_Traced_Class_track(dict_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Comonad_Traced_Class_tracks(dictComonadTraced_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, w_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadTraced_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadTraced_0_loop
	_ = dictComonadTraced_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var w_2 gopurs_runtime.Value = w_2_loop
	_ = w_2
	return gopurs_runtime.Apply2(dictComonadTraced_0.V1, gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictComonadTraced_0.V0, gopurs_runtime.Value{}), "extract"), w_2)), w_2)
}

func Call_Control_Comonad_Traced_Class_lowerTrack(dictComonadTrans_0_loop *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictComonadTrans_0 *Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value] = dictComonadTrans_0_loop
	_ = dictComonadTrans_0
	// TAST (Let): lower_1_0 shape=App(Var) bindingType=Any
	lower_1_0 := Call_Control_Comonad_Trans_Class_lower(dictComonadTrans_0)
	_ = lower_1_0
	return gopurs_runtime.Func(func(dictComonadTraced_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): lower1_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t$scope13) [(TypeVar w$scope15), (TypeVar a$scope16)])] (TypeApp (TypeVar w$scope15) [(TypeVar a$scope16)]))
		lower1_3_1 := gopurs_runtime.Apply(lower_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_2, "Comonad0"), gopurs_runtime.Value{}))
		_ = lower1_3_1
		return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_2, "track"), m_4), lower1_3_1)
		})
	})
}

func Call_Control_Comonad_Traced_Class_listens(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func2(func(g_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply(g_3, t_4), gopurs_runtime.Apply(f_1, t_4)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), v_2)
}

func Call_Control_Comonad_Traced_Class_listen(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, t_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply(f_2, t_3), t_3}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), v_1)
}

func Call_Control_Comonad_Traced_Class_comonadTracedTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): comonadTracedT_1_0 shape=App(Var) bindingType=Any
	comonadTracedT_1_0 := Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0)
	_ = comonadTracedT_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): comonadTracedT1_3_1 shape=App(Other) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeVar w$scope44) [(Func [(TypeVar t$scope45)] (TypeVar a))])])
		comonadTracedT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(comonadTracedT_1_0, dictMonoid_2))
		_ = comonadTracedT1_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 433046755, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_3_1)}
		}), gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_5, t_4)
		})}))}
	})
}

func Call_Control_Comonad_Traced_Class_comonadTracedStoreT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
	_ = dictComonadTraced_0
	// TAST (Let): comonadStoreT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w$scope54) [(Func [(TypeVar s$scope55)] (TypeVar a))]), (TypeVar s$scope55)])])
	comonadStoreT_1_0 := Rebox_Control_Comonad_Traced_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Comonad_Store_Trans_comonadStoreT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))))
	_ = comonadStoreT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 433046755, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Traced_Class_843532572_2535844177((&Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Traced_Class_3056445460_2550391993(comonadStoreT_1_0))}
	}), gopurs_runtime.Apply(Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Get_Control_Comonad_Store_Trans_comonadTransStoreT())), dictComonadTraced_0)})))}
}

func Call_Control_Comonad_Traced_Class_comonadTracedIdentityT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
	_ = dictComonadTraced_0
	// TAST (Let): comonadIdentityT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeApp (TypeVar m) [(TypeVar a)]) [(TypeVar w$scope63)])])
	comonadIdentityT_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Monad_Identity_Trans_comonadIdentityT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{})))
	_ = comonadIdentityT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 433046755, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadIdentityT_1_0)}
	}), gopurs_runtime.Apply(Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Get_Control_Comonad_Trans_Class_comonadTransIdentityT())), dictComonadTraced_0)}))}
}

func Call_Control_Comonad_Traced_Class_comonadTracedEnvT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
	_ = dictComonadTraced_0
	// TAST (Let): comonadEnvT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope73), (TypeVar w$scope72)])])
	comonadEnvT_1_0 := Rebox_Control_Comonad_Traced_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_comonadEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))))
	_ = comonadEnvT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 433046755, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Traced_Class_843532572_2535844177((&Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Traced_Class_3056445460_2550391993(comonadEnvT_1_0))}
	}), gopurs_runtime.Apply(Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Get_Control_Comonad_Env_Trans_comonadTransEnvT())), dictComonadTraced_0)})))}
}

func Call_Control_Comonad_Traced_Class_censor(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_1, v1_3)
	}), v_2)
}

func Rebox_Control_Comonad_Traced_Class_2550391993_3056445460(in *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Traced_Class_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Traced_Class_843532572_2535844177(in *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Traced_Class_ComonadTraced[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
