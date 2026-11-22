package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_State_Class_MonadState_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_State_Class_MonadState_dollar_Dict sync.Once

func Get_Control_Monad_State_Class_MonadState_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_State_Class_MonadState_dollar_Dict.Do(func() {
		cache_Control_Monad_State_Class_MonadState_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer(Call_Control_Monad_State_Class_MonadState_dollar_Dict(func() struct {
				Monad0 gopurs_runtime.Value
				state  gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640030,00730074006100740065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0 gopurs_runtime.Value
					state  gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640030,00730074006100740065*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.state = gopurs_runtime.RecordGet(orig, "state")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_State_Class_MonadState_dollar_Dict
}

var cache_Control_Monad_State_Class_state gopurs_runtime.Value
var once_Control_Monad_State_Class_state sync.Once

func Get_Control_Monad_State_Class_state() gopurs_runtime.Value {
	once_Control_Monad_State_Class_state.Do(func() {
		cache_Control_Monad_State_Class_state = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Class_state(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_State_Class_state
}

var cache_Control_Monad_State_Class_put gopurs_runtime.Value
var once_Control_Monad_State_Class_put sync.Once

func Get_Control_Monad_State_Class_put() gopurs_runtime.Value {
	once_Control_Monad_State_Class_put.Do(func() {
		cache_Control_Monad_State_Class_put = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Class_put(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), s_1_box)
		})
	})
	return cache_Control_Monad_State_Class_put
}

var cache_Control_Monad_State_Class_modify_ gopurs_runtime.Value
var once_Control_Monad_State_Class_modify_ sync.Once

func Get_Control_Monad_State_Class_modify_() gopurs_runtime.Value {
	once_Control_Monad_State_Class_modify_.Do(func() {
		cache_Control_Monad_State_Class_modify_ = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Class_modify_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), f_1_box)
		})
	})
	return cache_Control_Monad_State_Class_modify_
}

var cache_Control_Monad_State_Class_modify gopurs_runtime.Value
var once_Control_Monad_State_Class_modify sync.Once

func Get_Control_Monad_State_Class_modify() gopurs_runtime.Value {
	once_Control_Monad_State_Class_modify.Do(func() {
		cache_Control_Monad_State_Class_modify = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Class_modify(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), f_1_box)
		})
	})
	return cache_Control_Monad_State_Class_modify
}

var cache_Control_Monad_State_Class_gets gopurs_runtime.Value
var once_Control_Monad_State_Class_gets sync.Once

func Get_Control_Monad_State_Class_gets() gopurs_runtime.Value {
	once_Control_Monad_State_Class_gets.Do(func() {
		cache_Control_Monad_State_Class_gets = gopurs_runtime.Func2(func(dictMonadState_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Class_gets(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadState_0_box), f_1_box)
		})
	})
	return cache_Control_Monad_State_Class_gets
}

var cache_Control_Monad_State_Class_get gopurs_runtime.Value
var once_Control_Monad_State_Class_get sync.Once

func Get_Control_Monad_State_Class_get() gopurs_runtime.Value {
	once_Control_Monad_State_Class_get.Do(func() {
		cache_Control_Monad_State_Class_get = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_Class_get(dictMonadState_0_box)
		})
	})
	return cache_Control_Monad_State_Class_get
}

type Constructor_Control_Monad_State_Class_MonadState[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f00530074006100740065005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f00530074006100740065005f0043006c006100730073005f004d006f006e0061006400530074006100740065
func init() {
	gopurs_runtime.StructGetters[2100320995] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "state":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_State_Class_MonadState: " + key)
		}
	}
}

func Call_Control_Monad_State_Class_MonadState_dollar_Dict(x_0_loop struct {
	Monad0 gopurs_runtime.Value
	state  gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640030,00730074006100740065*/
}) *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad0 gopurs_runtime.Value
		state  gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e006100640030,00730074006100740065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "state", orig.Monad0, orig.state)
	}())
}

func Call_Control_Monad_State_Class_state(dict_0_loop *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Monad_State_Class_put(dictMonadState_0_loop *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
	_ = dictMonadState_0
	var s_1 gopurs_runtime.Value = s_1_loop
	_ = s_1
	return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{Get_Data_Unit_unit(), s_1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}))
}

func Call_Control_Monad_State_Class_modify_(dictMonadState_0_loop *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
	_ = dictMonadState_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{Get_Data_Unit_unit(), gopurs_runtime.Apply(f_1, s_2)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}))
}

func Call_Control_Monad_State_Class_modify(dictMonadState_0_loop *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
	_ = dictMonadState_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): s_prime__3_0 shape=App(Other) bindingType=(TypeVar s$scope17)
		s_prime__3_0 := gopurs_runtime.Apply(f_1, s_2)
		_ = s_prime__3_0
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{s_prime__3_0, s_prime__3_0}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}))
}

func Call_Control_Monad_State_Class_gets(dictMonadState_0_loop *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 *Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadState_0_loop
	_ = dictMonadState_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(dictMonadState_0.V1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply(f_1, s_2), s_2}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}))
}

func Call_Control_Monad_State_Class_get(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
	_ = dictMonadState_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{s_1, s_1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}))
}
