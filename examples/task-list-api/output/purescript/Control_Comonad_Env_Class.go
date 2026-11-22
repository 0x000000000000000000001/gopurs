package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Env_Class_compose gopurs_runtime.Value
var once_Control_Comonad_Env_Class_compose sync.Once

func Get_Control_Comonad_Env_Class_compose() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_compose.Do(func() {
		cache_Control_Comonad_Env_Class_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Env_Class_compose
}

var cache_Control_Comonad_Env_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Env_Class_lower sync.Once

func Get_Control_Comonad_Env_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_lower.Do(func() {
		cache_Control_Comonad_Env_Class_lower = gopurs_runtime.RecordGet(Get_Control_Comonad_Store_Trans_comonadTransStoreT(), "lower")
	})
	return cache_Control_Comonad_Env_Class_lower
}

var cache_Control_Comonad_Env_Class_compose1 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_compose1 sync.Once

func Get_Control_Comonad_Env_Class_compose1() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_compose1.Do(func() {
		cache_Control_Comonad_Env_Class_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Comonad_Env_Class_compose1
}

var cache_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict sync.Once

func Get_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict.Do(func() {
		cache_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict(func() struct {
				Comonad0 gopurs_runtime.Value
				ask      gopurs_runtime.Value
				/*gopurs:record-v1:2:0043006f006d006f006e006100640030,00610073006b*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Comonad0 gopurs_runtime.Value
					ask      gopurs_runtime.Value
					/*gopurs:record-v1:2:0043006f006d006f006e006100640030,00610073006b*/
				}{}
				clone.Comonad0 = gopurs_runtime.RecordGet(orig, "Comonad0")
				clone.ask = gopurs_runtime.RecordGet(orig, "ask")
				return clone
			}()))}
		})
	})
	return cache_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict
}

var cache_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict sync.Once

func Get_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict.Do(func() {
		cache_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict(func() struct {
				ComonadAsk0 gopurs_runtime.Value
				local       gopurs_runtime.Value
				/*gopurs:record-v1:2:0043006f006d006f006e0061006400410073006b0030,006c006f00630061006c*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					ComonadAsk0 gopurs_runtime.Value
					local       gopurs_runtime.Value
					/*gopurs:record-v1:2:0043006f006d006f006e0061006400410073006b0030,006c006f00630061006c*/
				}{}
				clone.ComonadAsk0 = gopurs_runtime.RecordGet(orig, "ComonadAsk0")
				clone.local = gopurs_runtime.RecordGet(orig, "local")
				return clone
			}()))}
		})
	})
	return cache_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict
}

var cache_Control_Comonad_Env_Class_local gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local sync.Once

func Get_Control_Comonad_Env_Class_local() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local.Do(func() {
		cache_Control_Comonad_Env_Class_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Comonad_Env_Class_local
}

var cache_Control_Comonad_Env_Class_comonadAskTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTuple sync.Once

func Get_Control_Comonad_Env_Class_comonadAskTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTuple = gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Tuple_comonadTuple()
		}), Get_Data_Tuple_fst()})))}
	})
	return cache_Control_Comonad_Env_Class_comonadAskTuple
}

var cache_Control_Comonad_Env_Class_comonadEnvTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTuple sync.Once

func Get_Control_Comonad_Env_Class_comonadEnvTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTuple = gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_144036252_155172497((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Comonad_Env_Class_comonadAskTuple()
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
	return cache_Control_Comonad_Env_Class_comonadEnvTuple
}

var cache_Control_Comonad_Env_Class_comonadAskEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskEnvT sync.Once

func Get_Control_Comonad_Env_Class_comonadAskEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskEnvT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Env_Class_comonadAskEnvT
}

var cache_Control_Comonad_Env_Class_comonadEnvEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvEnvT sync.Once

func Get_Control_Comonad_Env_Class_comonadEnvEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvEnvT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvEnvT
}

var cache_Control_Comonad_Env_Class_ask gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ask sync.Once

func Get_Control_Comonad_Env_Class_ask() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ask.Do(func() {
		cache_Control_Comonad_Env_Class_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_ask(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Comonad_Env_Class_ask
}

var cache_Control_Comonad_Env_Class_asks gopurs_runtime.Value
var once_Control_Comonad_Env_Class_asks sync.Once

func Get_Control_Comonad_Env_Class_asks() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_asks.Do(func() {
		cache_Control_Comonad_Env_Class_asks = gopurs_runtime.Func3(func(dictComonadAsk_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadAsk_0_box), f_1_box, x_2_box)
		})
	})
	return cache_Control_Comonad_Env_Class_asks
}

var cache_Control_Comonad_Env_Class_comonadAskStoreT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskStoreT sync.Once

func Get_Control_Comonad_Env_Class_comonadAskStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskStoreT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskStoreT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_comonadAskStoreT(dictComonadAsk_0_box)
		})
	})
	return cache_Control_Comonad_Env_Class_comonadAskStoreT
}

var cache_Control_Comonad_Env_Class_comonadEnvStoreT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvStoreT sync.Once

func Get_Control_Comonad_Env_Class_comonadEnvStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvStoreT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvStoreT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_box)
		})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvStoreT
}

var cache_Control_Comonad_Env_Class_comonadAskTracedT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTracedT sync.Once

func Get_Control_Comonad_Env_Class_comonadAskTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTracedT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTracedT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_box)
		})
	})
	return cache_Control_Comonad_Env_Class_comonadAskTracedT
}

var cache_Control_Comonad_Env_Class_comonadEnvTracedT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTracedT sync.Once

func Get_Control_Comonad_Env_Class_comonadEnvTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTracedT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTracedT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_box)
		})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvTracedT
}

type Constructor_Control_Comonad_Env_Class_ComonadAsk[T_e any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f0045006e0076005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f0045006e0076005f0043006c006100730073005f0043006f006d006f006e0061006400410073006b
func init() {
	gopurs_runtime.StructGetters[1424092807] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Comonad0":
			return gopurs_runtime.Box(c.V0)
		case "ask":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadAsk: " + key)
		}
	}
}

type Constructor_Control_Comonad_Env_Class_ComonadEnv[T_e any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f0045006e0076005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0043006f006d006f006e00610064005f0045006e0076005f0043006c006100730073005f0043006f006d006f006e006100640045006e0076
func init() {
	gopurs_runtime.StructGetters[3863290147] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "ComonadAsk0":
			return gopurs_runtime.Box(c.V0)
		case "local":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadEnv: " + key)
		}
	}
}

func Call_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict(x_0_loop struct {
	Comonad0 gopurs_runtime.Value
	ask      gopurs_runtime.Value
	/*gopurs:record-v1:2:0043006f006d006f006e006100640030,00610073006b*/
}) *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Comonad0 gopurs_runtime.Value
		ask      gopurs_runtime.Value
		/*gopurs:record-v1:2:0043006f006d006f006e006100640030,00610073006b*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Comonad0", "ask", orig.Comonad0, orig.ask)
	}())
}

func Call_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict(x_0_loop struct {
	ComonadAsk0 gopurs_runtime.Value
	local       gopurs_runtime.Value
	/*gopurs:record-v1:2:0043006f006d006f006e0061006400410073006b0030,006c006f00630061006c*/
}) *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		ComonadAsk0 gopurs_runtime.Value
		local       gopurs_runtime.Value
		/*gopurs:record-v1:2:0043006f006d006f006e0061006400410073006b0030,006c006f00630061006c*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("ComonadAsk0", "local", orig.ComonadAsk0, orig.local)
	}())
}

func Call_Control_Comonad_Env_Class_local(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): comonadEnvT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope62), (TypeVar w$scope61)])])
	comonadEnvT_1_0 := Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Comonad_Env_Trans_comonadEnvT(dictComonad_0)))
	_ = comonadEnvT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(comonadEnvT_1_0))}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
	})})))}
}

func Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): comonadAskEnvT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Env","Class","ComonadAsk"] [(TypeVar e$scope42), (TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e$scope42), (TypeVar w$scope41)])])
	comonadAskEnvT1_1_0 := Rebox_Control_Comonad_Env_Class_2829962037_1649752824(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0)))
	_ = comonadAskEnvT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_144036252_155172497((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037(comonadAskEnvT1_1_0))}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})})))}
}

func Call_Control_Comonad_Env_Class_ask(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Comonad_Env_Class_asks(dictComonadAsk_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadAsk_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadAsk_0_loop
	_ = dictComonadAsk_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(dictComonadAsk_0.V1, x_2))
}

func Call_Control_Comonad_Env_Class_comonadAskStoreT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
	_ = dictComonadAsk_0
	// TAST (Let): Comonad0_1_0 shape=App(Other) bindingType=Any
	Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
	_ = Comonad0_1_0
	// TAST (Let): comonadStoreT_2_1 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w$scope82) [(Func [(TypeVar s$scope83)] (TypeVar a))]), (TypeVar s$scope83)])])
	comonadStoreT_2_1 := Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Call_Control_Comonad_Store_Trans_comonadStoreT(Comonad0_1_0)))
	_ = comonadStoreT_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(comonadStoreT_2_1))}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Comonad_Env_Class_ask(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadAsk_0)), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_2 shape=Other bindingType=Any
		__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1
		_ = __local_var_4_2
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v1_5, __local_var_4_2)
		}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)
	}))})))}
}

func Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
	_ = dictComonadEnv_0
	// TAST (Let): comonadAskStoreT1_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Env","Class","ComonadAsk"] [(TypeVar e$scope27), (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w$scope28) [(Func [(TypeVar s$scope29)] (TypeVar a))]), (TypeVar s$scope29)])])
	comonadAskStoreT1_1_0 := Rebox_Control_Comonad_Env_Class_2829962037_1649752824(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Comonad_Env_Class_comonadAskStoreT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{}))))
	_ = comonadAskStoreT1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_144036252_155172497((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037(comonadAskStoreT1_1_0))}
	}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})})))}
}

func Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
	_ = dictComonadAsk_0
	// TAST (Let): ask1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar w$scope93) [(TypeVar a$scope98)])] (TypeVar e$scope92))
	ask1_1_0 := Call_Control_Comonad_Env_Class_ask(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadAsk_0))
	_ = ask1_1_0
	// TAST (Let): Comonad0_2_1 shape=App(Other) bindingType=Any
	Comonad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
	_ = Comonad0_2_1
	// TAST (Let): comonadTracedT_3_2 shape=App(Var) bindingType=Any
	comonadTracedT_3_2 := Call_Control_Comonad_Traced_Trans_comonadTracedT(Comonad0_2_1)
	_ = comonadTracedT_3_2
	return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): comonadTracedT1_5_3 shape=App(Other) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeVar w$scope93) [(Func [(TypeVar t$scope94)] (TypeVar a))])])
		comonadTracedT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(comonadTracedT_3_2, dictMonoid_4))
		_ = comonadTracedT1_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_5_3)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), ask1_1_0, gopurs_runtime.Apply(Call_Control_Comonad_Trans_Class_lower(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]](Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_4))), Comonad0_2_1))}))}
	})
}

func Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
	_ = dictComonadEnv_0
	// TAST (Let): comonadAskTracedT1_1_0 shape=App(Var) bindingType=Any
	comonadAskTracedT1_1_0 := Call_Control_Comonad_Env_Class_comonadAskTracedT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{}))
	_ = comonadAskTracedT1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): comonadAskTracedT2_3_1 shape=App(Other) bindingType=(ADT ["Control","Comonad","Env","Class","ComonadAsk"] [(TypeVar e$scope14), (TypeApp (TypeVar w$scope15) [(Func [(TypeVar t$scope16)] (TypeVar a))])])
		comonadAskTracedT2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(comonadAskTracedT1_1_0, dictMonoid_2))
		_ = comonadAskTracedT2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(comonadAskTracedT2_3_1)}
		}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_4, v_5)
		})}))}
	})
}

func Rebox_Control_Comonad_Env_Class_144036252_155172497(in *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Class_1649752824_2829962037(in *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Class_2550391993_3056445460(in *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Class_2829962037_1649752824(in *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Comonad_Env_Class_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(unsafe.Pointer(in))
}
