package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Aff_Class_compose gopurs_runtime.Value
var once_Effect_Aff_Class_compose sync.Once

func Get_Effect_Aff_Class_compose() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose.Do(func() {
		cache_Effect_Aff_Class_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose
}

var cache_Effect_Aff_Class_lift gopurs_runtime.Value
var once_Effect_Aff_Class_lift sync.Once

func Get_Effect_Aff_Class_lift() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift.Do(func() {
		cache_Effect_Aff_Class_lift = gopurs_runtime.RecordGet(Get_Control_Monad_Cont_Trans_monadTransContT(), "lift")
	})
	return cache_Effect_Aff_Class_lift
}

var cache_Effect_Aff_Class_compose1 gopurs_runtime.Value
var once_Effect_Aff_Class_compose1 sync.Once

func Get_Effect_Aff_Class_compose1() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose1.Do(func() {
		cache_Effect_Aff_Class_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose1
}

var cache_Effect_Aff_Class_lift1 gopurs_runtime.Value
var once_Effect_Aff_Class_lift1 sync.Once

func Get_Effect_Aff_Class_lift1() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift1.Do(func() {
		cache_Effect_Aff_Class_lift1 = gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift")
	})
	return cache_Effect_Aff_Class_lift1
}

var cache_Effect_Aff_Class_compose2 gopurs_runtime.Value
var once_Effect_Aff_Class_compose2 sync.Once

func Get_Effect_Aff_Class_compose2() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose2.Do(func() {
		cache_Effect_Aff_Class_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose2
}

var cache_Effect_Aff_Class_lift2 gopurs_runtime.Value
var once_Effect_Aff_Class_lift2 sync.Once

func Get_Effect_Aff_Class_lift2() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift2.Do(func() {
		cache_Effect_Aff_Class_lift2 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_List_Trans_monadTransListT()).V0
	})
	return cache_Effect_Aff_Class_lift2
}

var cache_Effect_Aff_Class_compose3 gopurs_runtime.Value
var once_Effect_Aff_Class_compose3 sync.Once

func Get_Effect_Aff_Class_compose3() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose3.Do(func() {
		cache_Effect_Aff_Class_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose3
}

var cache_Effect_Aff_Class_lift3 gopurs_runtime.Value
var once_Effect_Aff_Class_lift3 sync.Once

func Get_Effect_Aff_Class_lift3() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift3.Do(func() {
		cache_Effect_Aff_Class_lift3 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Maybe_Trans_monadTransMaybeT()).V0
	})
	return cache_Effect_Aff_Class_lift3
}

var cache_Effect_Aff_Class_compose4 gopurs_runtime.Value
var once_Effect_Aff_Class_compose4 sync.Once

func Get_Effect_Aff_Class_compose4() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose4.Do(func() {
		cache_Effect_Aff_Class_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose4
}

var cache_Effect_Aff_Class_compose5 gopurs_runtime.Value
var once_Effect_Aff_Class_compose5 sync.Once

func Get_Effect_Aff_Class_compose5() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose5.Do(func() {
		cache_Effect_Aff_Class_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose5
}

var cache_Effect_Aff_Class_lift4 gopurs_runtime.Value
var once_Effect_Aff_Class_lift4 sync.Once

func Get_Effect_Aff_Class_lift4() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift4.Do(func() {
		cache_Effect_Aff_Class_lift4 = gopurs_runtime.RecordGet(Get_Control_Monad_Reader_Trans_monadTransReaderT(), "lift")
	})
	return cache_Effect_Aff_Class_lift4
}

var cache_Effect_Aff_Class_compose6 gopurs_runtime.Value
var once_Effect_Aff_Class_compose6 sync.Once

func Get_Effect_Aff_Class_compose6() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose6.Do(func() {
		cache_Effect_Aff_Class_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose6
}

var cache_Effect_Aff_Class_lift5 gopurs_runtime.Value
var once_Effect_Aff_Class_lift5 sync.Once

func Get_Effect_Aff_Class_lift5() gopurs_runtime.Value {
	once_Effect_Aff_Class_lift5.Do(func() {
		cache_Effect_Aff_Class_lift5 = gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift")
	})
	return cache_Effect_Aff_Class_lift5
}

var cache_Effect_Aff_Class_compose7 gopurs_runtime.Value
var once_Effect_Aff_Class_compose7 sync.Once

func Get_Effect_Aff_Class_compose7() gopurs_runtime.Value {
	once_Effect_Aff_Class_compose7.Do(func() {
		cache_Effect_Aff_Class_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Class_compose7
}

var cache_Effect_Aff_Class_MonadAff_dollar_Dict gopurs_runtime.Value
var once_Effect_Aff_Class_MonadAff_dollar_Dict sync.Once

func Get_Effect_Aff_Class_MonadAff_dollar_Dict() gopurs_runtime.Value {
	once_Effect_Aff_Class_MonadAff_dollar_Dict.Do(func() {
		cache_Effect_Aff_Class_MonadAff_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer(Call_Effect_Aff_Class_MonadAff_dollar_Dict(func() struct {
				MonadEffect0 gopurs_runtime.Value
				liftAff      gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					MonadEffect0 gopurs_runtime.Value
					liftAff      gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
				}{}
				clone.MonadEffect0 = gopurs_runtime.RecordGet(orig, "MonadEffect0")
				clone.liftAff = gopurs_runtime.RecordGet(orig, "liftAff")
				return clone
			}()))}
		})
	})
	return cache_Effect_Aff_Class_MonadAff_dollar_Dict
}

var cache_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029 gopurs_runtime.Value
var once_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029 sync.Once

func Get_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029() gopurs_runtime.Value {
	once_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029.Do(func() {
		cache_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer(Call_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029(func() struct {
				MonadEffect0 gopurs_runtime.Value
				liftAff      gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					MonadEffect0 gopurs_runtime.Value
					liftAff      gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
				}{}
				clone.MonadEffect0 = gopurs_runtime.RecordGet(orig, "MonadEffect0")
				clone.liftAff = gopurs_runtime.RecordGet(orig, "liftAff")
				return clone
			}()))}
		})
	})
	return cache_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029
}

var cache_Effect_Aff_Class_monadAffAff gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffAff sync.Once

func Get_Effect_Aff_Class_monadAffAff() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffAff.Do(func() {
		cache_Effect_Aff_Class_monadAffAff = gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Effect_Aff_Class_monadAffAff
}

var cache_Effect_Aff_Class_liftAff gopurs_runtime.Value
var once_Effect_Aff_Class_liftAff sync.Once

func Get_Effect_Aff_Class_liftAff() gopurs_runtime.Value {
	once_Effect_Aff_Class_liftAff.Do(func() {
		cache_Effect_Aff_Class_liftAff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Effect_Aff_Class_liftAff
}

var cache_Effect_Aff_Class_monadAffContT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffContT sync.Once

func Get_Effect_Aff_Class_monadAffContT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffContT.Do(func() {
		cache_Effect_Aff_Class_monadAffContT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffContT(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffContT
}

var cache_Effect_Aff_Class_monadAffExceptT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffExceptT sync.Once

func Get_Effect_Aff_Class_monadAffExceptT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffExceptT.Do(func() {
		cache_Effect_Aff_Class_monadAffExceptT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffExceptT(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffExceptT
}

var cache_Effect_Aff_Class_monadAffListT gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffListT sync.Once

func Get_Effect_Aff_Class_monadAffListT() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffListT.Do(func() {
		cache_Effect_Aff_Class_monadAffListT = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffListT(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffListT
}

var cache_Effect_Aff_Class_monadAffMaybe gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffMaybe sync.Once

func Get_Effect_Aff_Class_monadAffMaybe() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffMaybe.Do(func() {
		cache_Effect_Aff_Class_monadAffMaybe = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffMaybe(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffMaybe
}

var cache_Effect_Aff_Class_monadAffRWS gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffRWS sync.Once

func Get_Effect_Aff_Class_monadAffRWS() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffRWS.Do(func() {
		cache_Effect_Aff_Class_monadAffRWS = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffRWS(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffRWS
}

var cache_Effect_Aff_Class_monadAffReader gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffReader sync.Once

func Get_Effect_Aff_Class_monadAffReader() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffReader.Do(func() {
		cache_Effect_Aff_Class_monadAffReader = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffReader(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffReader
}

var cache_Effect_Aff_Class_monadAffState gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffState sync.Once

func Get_Effect_Aff_Class_monadAffState() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffState.Do(func() {
		cache_Effect_Aff_Class_monadAffState = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffState(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffState
}

var cache_Effect_Aff_Class_monadAffWriter gopurs_runtime.Value
var once_Effect_Aff_Class_monadAffWriter sync.Once

func Get_Effect_Aff_Class_monadAffWriter() gopurs_runtime.Value {
	once_Effect_Aff_Class_monadAffWriter.Do(func() {
		cache_Effect_Aff_Class_monadAffWriter = gopurs_runtime.Func(func(dictMonadAff_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Class_monadAffWriter(dictMonadAff_0_box)
		})
	})
	return cache_Effect_Aff_Class_monadAffWriter
}

type Constructor_Effect_Aff_Class_MonadAff[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004500660066006500630074005f004100660066005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004500660066006500630074005f004100660066005f0043006c006100730073005f004d006f006e00610064004100660066
func init() {
	gopurs_runtime.StructGetters[3183257445] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "MonadEffect0":
			return gopurs_runtime.Box(c.V0)
		case "liftAff":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Effect_Aff_Class_MonadAff: " + key)
		}
	}
}

func Call_Effect_Aff_Class_MonadAff_dollar_Dict(x_0_loop struct {
	MonadEffect0 gopurs_runtime.Value
	liftAff      gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
}) *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] {
	var x_0 struct {
		MonadEffect0 gopurs_runtime.Value
		liftAff      gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", orig.MonadEffect0, orig.liftAff)
	}())
}

func Call_Effect_Aff_Class_MonadAff_dollar_Dict__2712063029(x_0_loop struct {
	MonadEffect0 gopurs_runtime.Value
	liftAff      gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
}) *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] {
MonadAff_dollar_Dict__2712063029:
	for {
		if false {
			continue MonadAff_dollar_Dict__2712063029
		}
		var x_0 struct {
			MonadEffect0 gopurs_runtime.Value
			liftAff      gopurs_runtime.Value
			/*gopurs:record-v1:2:004d006f006e006100640045006600660065006300740030,006c006900660074004100660066*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("MonadEffect0", "liftAff", orig.MonadEffect0, orig.liftAff)
		}())
	}
}

func Call_Effect_Aff_Class_liftAff(dict_0_loop *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Effect_Aff_Class_monadAffContT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): monadEffectContT_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r$scope9), (TypeVar m$scope8)])])
	monadEffectContT_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_monadEffectContT(MonadEffect0_1_0))
	_ = monadEffectContT_2_1
	// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
	__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
	_ = __local_var_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectContT_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func2(func(m_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}), "bind"), m_4, k_5)
	}), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffExceptT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): monadEffectExceptT_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope16) [(ADT ["Data","Either","Either"] [(TypeVar e$scope17), (TypeVar a)])])])
	monadEffectExceptT_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_monadEffectExceptT(MonadEffect0_1_0))
	_ = monadEffectExceptT_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectExceptT_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffListT(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): monadEffectListT_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope23) [(ADT ["Control","Monad","List","Trans","Step"] [(TypeVar a), (ADT ["Control","Monad","List","Trans","ListT"] [(TypeVar m$scope23), (TypeVar a)])])])])
	monadEffectListT_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_List_Trans_monadEffectListT(MonadEffect0_1_0))
	_ = monadEffectListT_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectListT_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Monad_List_Trans_fromEffect(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffMaybe(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): monadEffectMaybe_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope29) [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])])
	monadEffectMaybe_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Maybe_Trans_monadEffectMaybe(MonadEffect0_1_0))
	_ = monadEffectMaybe_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectMaybe_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Get_Control_Monad_Maybe_Trans_monadTransMaybeT()).V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffRWS(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): Monad0_2_1 shape=App(Other) bindingType=Any
	Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_2_1
	// TAST (Let): liftAff1_3_2 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a$scope44)])] (TypeApp (TypeVar m$scope38) [(TypeVar a$scope44)]))
	liftAff1_3_2 := Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0))
	_ = liftAff1_3_2
	return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadEffectRWS_5_3 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(Func [(TypeVar r$scope40), (TypeVar s$scope41)] (TypeApp (TypeVar m$scope38) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s$scope41), (TypeVar a), (TypeVar w$scope39)])]))])
		monadEffectRWS_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_4), MonadEffect0_1_0))
		_ = monadEffectRWS_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectRWS_5_3)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_4))), Monad0_2_1), liftAff1_3_2)}))}
	})
}

func Call_Effect_Aff_Class_monadAffReader(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): monadEffectReader_1_0 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r$scope49), (TypeVar m$scope48)])])
	monadEffectReader_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_Reader_Trans_monadEffectReader(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})))
	_ = monadEffectReader_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectReader_1_0)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Control_Monad_Reader_Trans_ReaderT(), Get_Data_Function_go__const()), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffState(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): monadEffectState_2_1 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(Func [(TypeVar s$scope57)] (TypeApp (TypeVar m$scope56) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s$scope57)])]))])
	monadEffectState_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_monadEffectState(MonadEffect0_1_0))
	_ = monadEffectState_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectState_2_1)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})), Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0)))}))}
}

func Call_Effect_Aff_Class_monadAffWriter(dictMonadAff_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAff_0 gopurs_runtime.Value = dictMonadAff_0_loop
	_ = dictMonadAff_0
	// TAST (Let): MonadEffect0_1_0 shape=App(Other) bindingType=Any
	MonadEffect0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAff_0, "MonadEffect0"), gopurs_runtime.Value{})
	_ = MonadEffect0_1_0
	// TAST (Let): Monad0_2_1 shape=App(Other) bindingType=Any
	Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadEffect0_1_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_2_1
	// TAST (Let): liftAff1_3_2 shape=App(Var) bindingType=(Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a$scope68)])] (TypeApp (TypeVar m$scope64) [(TypeVar a$scope68)]))
	liftAff1_3_2 := Call_Effect_Aff_Class_liftAff(gopurs_runtime.CoerceToStruct[Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]](dictMonadAff_0))
	_ = liftAff1_3_2
	return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): monadEffectWriter_5_3 shape=App(Var) bindingType=(ADT ["Effect","Class","MonadEffect"] [(TypeApp (TypeVar m$scope64) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope65)])])])
		monadEffectWriter_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Control_Monad_Writer_Trans_monadEffectWriter(dictMonoid_4), MonadEffect0_1_0))
		_ = monadEffectWriter_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3183257445, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Aff_Class_MonadAff[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(monadEffectWriter_5_3)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](Call_Control_Monad_Writer_Trans_monadTransWriterT(dictMonoid_4))), Monad0_2_1), liftAff1_3_2)}))}
	})
}
