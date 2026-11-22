package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Rec_Class_compose gopurs_runtime.Value
var once_Control_Monad_Rec_Class_compose sync.Once

func Get_Control_Monad_Rec_Class_compose() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_compose.Do(func() {
		cache_Control_Monad_Rec_Class_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Rec_Class_compose
}

var cache_Control_Monad_Rec_Class_compose1 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_compose1 sync.Once

func Get_Control_Monad_Rec_Class_compose1() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_compose1.Do(func() {
		cache_Control_Monad_Rec_Class_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Rec_Class_compose1
}

var cache_Control_Monad_Rec_Class_compose2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_compose2 sync.Once

func Get_Control_Monad_Rec_Class_compose2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_compose2.Do(func() {
		cache_Control_Monad_Rec_Class_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Rec_Class_compose2
}

var cache_Control_Monad_Rec_Class_bind gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bind sync.Once

func Get_Control_Monad_Rec_Class_bind() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bind.Do(func() {
		cache_Control_Monad_Rec_Class_bind = Get_Effect_bindE()
	})
	return cache_Control_Monad_Rec_Class_bind
}

var cache_Control_Monad_Rec_Class_bindFlipped gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bindFlipped sync.Once

func Get_Control_Monad_Rec_Class_bindFlipped() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bindFlipped.Do(func() {
		cache_Control_Monad_Rec_Class_bindFlipped = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_bindFlipped(b_0_box, a_1_box)
		})
	})
	return cache_Control_Monad_Rec_Class_bindFlipped
}

var cache_Control_Monad_Rec_Class_discard gopurs_runtime.Value
var once_Control_Monad_Rec_Class_discard sync.Once

func Get_Control_Monad_Rec_Class_discard() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_discard.Do(func() {
		cache_Control_Monad_Rec_Class_discard = Get_Effect_bindE()
	})
	return cache_Control_Monad_Rec_Class_discard
}

var cache_Control_Monad_Rec_Class_bind1 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bind1 sync.Once

func Get_Control_Monad_Rec_Class_bind1() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bind1.Do(func() {
		cache_Control_Monad_Rec_Class_bind1 = Get_Effect_bindE()
	})
	return cache_Control_Monad_Rec_Class_bind1
}

var cache_Control_Monad_Rec_Class_bind2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bind2 sync.Once

func Get_Control_Monad_Rec_Class_bind2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bind2.Do(func() {
		cache_Control_Monad_Rec_Class_bind2 = Get_Control_Bind_bind__4262485428()
	})
	return cache_Control_Monad_Rec_Class_bind2
}

var cache_Control_Monad_Rec_Class_pure gopurs_runtime.Value
var once_Control_Monad_Rec_Class_pure sync.Once

func Get_Control_Monad_Rec_Class_pure() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_pure.Do(func() {
		cache_Control_Monad_Rec_Class_pure = Get_Control_Applicative_pure__1379658716()
	})
	return cache_Control_Monad_Rec_Class_pure
}

var cache_Control_Monad_Rec_Class_go__map gopurs_runtime.Value
var once_Control_Monad_Rec_Class_go__map sync.Once

func Get_Control_Monad_Rec_Class_go__map() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_go__map.Do(func() {
		cache_Control_Monad_Rec_Class_go__map = gopurs_runtime.Apply(Get_Control_Applicative_liftA1(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))})
	})
	return cache_Control_Monad_Rec_Class_go__map
}

var cache_Control_Monad_Rec_Class_Loop gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Loop sync.Once

func Get_Control_Monad_Rec_Class_Loop() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Loop.Do(func() {
		cache_Control_Monad_Rec_Class_Loop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Control_Monad_Rec_Class_Loop
}

var cache_Control_Monad_Rec_Class_Loop__103919764 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Loop__103919764 sync.Once

func Get_Control_Monad_Rec_Class_Loop__103919764() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Loop__103919764.Do(func() {
		cache_Control_Monad_Rec_Class_Loop__103919764 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_Loop__103919764(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Control_Monad_Rec_Class_Loop__103919764
}

var cache_Control_Monad_Rec_Class_Done gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Done sync.Once

func Get_Control_Monad_Rec_Class_Done() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Done.Do(func() {
		cache_Control_Monad_Rec_Class_Done = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Control_Monad_Rec_Class_Done
}

var cache_Control_Monad_Rec_Class_Done__4129036734 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_Done__4129036734 sync.Once

func Get_Control_Monad_Rec_Class_Done__4129036734() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_Done__4129036734.Do(func() {
		cache_Control_Monad_Rec_Class_Done__4129036734 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_Done__4129036734(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Control_Monad_Rec_Class_Done__4129036734
}

var cache_Control_Monad_Rec_Class_MonadRec_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Rec_Class_MonadRec_dollar_Dict sync.Once

func Get_Control_Monad_Rec_Class_MonadRec_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_MonadRec_dollar_Dict.Do(func() {
		cache_Control_Monad_Rec_Class_MonadRec_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Rec_Class_MonadRec_dollar_Dict(func() struct {
				Monad0   gopurs_runtime.Value
				tailRecM gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640030,007400610069006c005200650063004d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0   gopurs_runtime.Value
					tailRecM gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640030,007400610069006c005200650063004d*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.tailRecM = gopurs_runtime.RecordGet(orig, "tailRecM")
				return clone
			}()))}
		})
	})
	return cache_Control_Monad_Rec_Class_MonadRec_dollar_Dict
}

var cache_Control_Monad_Rec_Class_tailRecM gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM sync.Once

func Get_Control_Monad_Rec_Class_tailRecM() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_tailRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Monad_Rec_Class_tailRecM
}

var cache_Control_Monad_Rec_Class_tailRecM2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM2 sync.Once

func Get_Control_Monad_Rec_Class_tailRecM2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM2.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM2 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_tailRecM2(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
		})
	})
	return cache_Control_Monad_Rec_Class_tailRecM2
}

var cache_Control_Monad_Rec_Class_tailRecM3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRecM3 sync.Once

func Get_Control_Monad_Rec_Class_tailRecM3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRecM3.Do(func() {
		cache_Control_Monad_Rec_Class_tailRecM3 = gopurs_runtime.Func5(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_tailRecM3(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box, c_4_box)
		})
	})
	return cache_Control_Monad_Rec_Class_tailRecM3
}

var cache_Control_Monad_Rec_Class_untilJust gopurs_runtime.Value
var once_Control_Monad_Rec_Class_untilJust sync.Once

func Get_Control_Monad_Rec_Class_untilJust() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_untilJust.Do(func() {
		cache_Control_Monad_Rec_Class_untilJust = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_untilJust(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Control_Monad_Rec_Class_untilJust
}

var cache_Control_Monad_Rec_Class_whileJust gopurs_runtime.Value
var once_Control_Monad_Rec_Class_whileJust sync.Once

func Get_Control_Monad_Rec_Class_whileJust() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_whileJust.Do(func() {
		cache_Control_Monad_Rec_Class_whileJust = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_whileJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Control_Monad_Rec_Class_whileJust
}

var cache_Control_Monad_Rec_Class_tailRec gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec sync.Once

func Get_Control_Monad_Rec_Class_tailRec() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_tailRec(f_0_box)
		})
	})
	return cache_Control_Monad_Rec_Class_tailRec
}

var cache_Control_Monad_Rec_Class_tailRec2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec2 sync.Once

func Get_Control_Monad_Rec_Class_tailRec2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec2.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_tailRec2(f_0_box, a_1_box, b_2_box)
		})
	})
	return cache_Control_Monad_Rec_Class_tailRec2
}

var cache_Control_Monad_Rec_Class_tailRec3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_tailRec3 sync.Once

func Get_Control_Monad_Rec_Class_tailRec3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_tailRec3.Do(func() {
		cache_Control_Monad_Rec_Class_tailRec3 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, c_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_tailRec3(f_0_box, a_1_box, b_2_box, c_3_box)
		})
	})
	return cache_Control_Monad_Rec_Class_tailRec3
}

var cache_Control_Monad_Rec_Class_monadRecMaybe gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecMaybe sync.Once

func Get_Control_Monad_Rec_Class_monadRecMaybe() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecMaybe.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_1542299734_4130553207((&Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_1642601656_2568689657(Rebox_Control_Monad_Rec_Class_2568689657_1642601656(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Maybe_monadMaybe()))))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}))}
						goto end_branch_3
					} else {

					}
				}
				{
					if v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil {
						var __t2 gopurs_runtime.Value
						{
							var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
							_ = __t_tag_0
							if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 525585346 {
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)}))}
								goto end_branch_2
							} else {

							}
						}
						{
							var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
							_ = __t_tag_1
							if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 60402430 {
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0}))}}))}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_2:
						__t3 = __t2
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			})), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, a0_1)))})
		})})))}
	})
	return cache_Control_Monad_Rec_Class_monadRecMaybe
}

var cache_Control_Monad_Rec_Class_monadRecIdentity gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecIdentity sync.Once

func Get_Control_Monad_Rec_Class_monadRecIdentity() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecIdentity.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Identity_monadIdentity()))}
		}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return v_1
			}), f_0)))
		})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecIdentity
}

var cache_Control_Monad_Rec_Class_monadRecFunction gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecFunction sync.Once

func Get_Control_Monad_Rec_Class_monadRecFunction() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecFunction.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecFunction = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_monadFn()
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value, e_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(f_0, a_3, e_2)
			})), a0_1)
		})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecFunction
}

var cache_Control_Monad_Rec_Class_monadRecEither gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecEither sync.Once

func Get_Control_Monad_Rec_Class_monadRecEither() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecEither.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecEither = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Either_monadEither()
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if v_2.Type == 9 && v_2.IntVal == 3711209382 {
						__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0}))}}))}
						goto end_branch_3
					} else {

					}
				}
				{
					if v_2.Type == 9 && v_2.IntVal == 2465973597 {
						var __t2 gopurs_runtime.Value
						{
							var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
							_ = __t_tag_0
							if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 525585346 {
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0)}))}
								goto end_branch_2
							} else {

							}
						}
						{
							var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
							_ = __t_tag_1
							if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 60402430 {
								__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0.UnsafePtr).V0}))}}))}
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_2:
						__t3 = __t2
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			})), gopurs_runtime.Apply(f_0, a0_1))
		})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecEither
}

var cache_Control_Monad_Rec_Class_monadRecEffect gopurs_runtime.Value
var once_Control_Monad_Rec_Class_monadRecEffect sync.Once

func Get_Control_Monad_Rec_Class_monadRecEffect() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_monadRecEffect.Do(func() {
		cache_Control_Monad_Rec_Class_monadRecEffect = gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(ADT ["Effect","Effect"] [(ADT ["Control","Monad","Rec","Class","Step"] [(TypeVar a$scope104), (TypeVar b$scope105)])])
				__local_var_2_0 := gopurs_runtime.Apply(f_0, a_1)
				_ = __local_var_2_0
				__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
				_ = __local_var_3_2
				r_3_1 := gopurs_runtime.Apply2(Get_Effect_Ref_go__new(), __local_var_3_2, gopurs_runtime.Value{})
				_ = r_3_1
				_dollar___unused_4_3 := gopurs_runtime.Apply2(Get_Effect_untilE(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_4_4 shape=App(Var) bindingType=Any
					__local_var_4_4 := gopurs_runtime.Apply(Get_Effect_Ref_read(), r_3_1)
					_ = __local_var_4_4
					v_5_5 := gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Value{})
					_ = v_5_5
					var __t8 gopurs_runtime.Value
					{
						if v_5_5.Type == 9 && v_5_5.IntVal == 525585346 {
							__t8 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								e_6_6 := gopurs_runtime.Apply2(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_5.UnsafePtr).V0, gopurs_runtime.Value{})
								_ = e_6_6
								__local_var_7_7 := gopurs_runtime.Apply3(Get_Effect_Ref_write(), e_6_6, r_3_1, gopurs_runtime.Value{})
								_ = __local_var_7_7
								return gopurs_runtime.Bool(false)
							})
							goto end_branch_8
						} else {

						}
					}
					{
						if v_5_5.Type == 9 && v_5_5.IntVal == 60402430 {
							__t8 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Bool(true)
							})
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_8:
					return gopurs_runtime.Apply(__t8, gopurs_runtime.Value{})
				}), gopurs_runtime.Value{})
				_ = _dollar___unused_4_3
				return gopurs_runtime.Apply(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t9 gopurs_runtime.Value
					{
						if v_5.Type == 9 && v_5.IntVal == 60402430 {
							__t9 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
							goto end_branch_9
						} else {

						}
					}
					{
						__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_9:
					return __t9
				}), gopurs_runtime.Apply(Get_Effect_Ref_read(), r_3_1)), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Control_Monad_Rec_Class_monadRecEffect
}

var cache_Control_Monad_Rec_Class_loop3 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_loop3 sync.Once

func Get_Control_Monad_Rec_Class_loop3() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_loop3.Do(func() {
		cache_Control_Monad_Rec_Class_loop3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_loop3(a_0_box, b_1_box, c_2_box)
		})
	})
	return cache_Control_Monad_Rec_Class_loop3
}

var cache_Control_Monad_Rec_Class_loop2 gopurs_runtime.Value
var once_Control_Monad_Rec_Class_loop2 sync.Once

func Get_Control_Monad_Rec_Class_loop2() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_loop2.Do(func() {
		cache_Control_Monad_Rec_Class_loop2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_loop2(a_0_box, b_1_box)
		})
	})
	return cache_Control_Monad_Rec_Class_loop2
}

var cache_Control_Monad_Rec_Class_functorStep gopurs_runtime.Value
var once_Control_Monad_Rec_Class_functorStep sync.Once

func Get_Control_Monad_Rec_Class_functorStep() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_functorStep.Do(func() {
		cache_Control_Monad_Rec_Class_functorStep = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if m_1.Type == 9 && m_1.IntVal == 525585346 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 60402430 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		})}))}
	})
	return cache_Control_Monad_Rec_Class_functorStep
}

var cache_Control_Monad_Rec_Class_forever gopurs_runtime.Value
var once_Control_Monad_Rec_Class_forever sync.Once

func Get_Control_Monad_Rec_Class_forever() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_forever.Do(func() {
		cache_Control_Monad_Rec_Class_forever = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Rec_Class_forever(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
		})
	})
	return cache_Control_Monad_Rec_Class_forever
}

var cache_Control_Monad_Rec_Class_bifunctorStep gopurs_runtime.Value
var once_Control_Monad_Rec_Class_bifunctorStep sync.Once

func Get_Control_Monad_Rec_Class_bifunctorStep() gopurs_runtime.Value {
	once_Control_Monad_Rec_Class_bifunctorStep.Do(func() {
		cache_Control_Monad_Rec_Class_bifunctorStep = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v2_2.Type == 9 && v2_2.IntVal == 525585346 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 60402430 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		})}))}
	})
	return cache_Control_Monad_Rec_Class_bifunctorStep
}

type Constructor_Control_Monad_Rec_Class_Loop[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Control_Monad_Rec_Class_Done[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}

type Constructor_Control_Monad_Rec_Class_MonadRec[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004d006f006e00610064005f005200650063005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004d006f006e00610064005f005200650063005f0043006c006100730073005f004d006f006e00610064005200650063
func init() {
	gopurs_runtime.StructGetters[3709389635] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "tailRecM":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Monad_Rec_Class_MonadRec: " + key)
		}
	}
}

func Call_Control_Monad_Rec_Class_bindFlipped(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_2_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
		_ = __local_var_2_0
		return gopurs_runtime.Apply2(b_0, __local_var_2_0, gopurs_runtime.Value{})
	})
}

func Call_Control_Monad_Rec_Class_Loop__103919764(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Loop__103919764:
	for {
		if false {
			continue Loop__103919764
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_45792715_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[int64, gopurs_runtime.Value]{1, __eta_norm_0_0})))}
	}
}

func Call_Control_Monad_Rec_Class_Done__4129036734(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Done__4129036734:
	for {
		if false {
			continue Done__4129036734
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_3558361591_3603546092((&Constructor_Control_Monad_Rec_Class_Done[int64, gopurs_runtime.Value]{1, Get_Data_Unit_unit()})))}
	}
}

func Call_Control_Monad_Rec_Class_MonadRec_dollar_Dict(x_0_loop struct {
	Monad0   gopurs_runtime.Value
	tailRecM gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640030,007400610069006c005200650063004d*/
}) *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] {
	var x_0 struct {
		Monad0   gopurs_runtime.Value
		tailRecM gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e006100640030,007400610069006c005200650063004d*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "tailRecM", orig.Monad0, orig.tailRecM)
	}())
}

func Call_Control_Monad_Rec_Class_tailRecM(dict_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Monad_Rec_Class_tailRecM2(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
	}), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			/*gopurs:record-v1:2:0061,0062*/
		}{a_2, b_3}
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
	}())
}

func Call_Control_Monad_Rec_Class_tailRecM3(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	var c_4 gopurs_runtime.Value = c_4_loop
	_ = c_4
	return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(f_1, gopurs_runtime.RecordGet(o_5, "a"), gopurs_runtime.RecordGet(o_5, "b"), gopurs_runtime.RecordGet(o_5, "c"))
	}), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			c gopurs_runtime.Value
			/*gopurs:record-v1:3:0061,0062,0063*/
		}{a_2, b_3, c_4}
		_ = orig
		return gopurs_runtime.RecordDict3("a", "b", "c", orig.a, orig.b, orig.c)
	}())
}

func Call_Control_Monad_Rec_Class_untilJust(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 gopurs_runtime.Value
				{
					if v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr == nil {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))}
						goto end_branch_1
					} else {

					}
				}
				{
					if v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil {
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_4.UnsafePtr).V0}))}
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_1:
				return __t1
			}), m_2)
		}), Get_Data_Unit_unit())
	})
}

func Call_Control_Monad_Rec_Class_whileJust(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar a$scope31)
	mempty_1_0 := dictMonoid_0.V1
	_ = mempty_1_0
	return gopurs_runtime.Func(func(dictMonadRec_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_2, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
		_ = __local_var_3_1
		return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_2, "tailRecM"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t2 gopurs_runtime.Value
					{
						if v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr == nil {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_5}))}
							goto end_branch_2
						} else {

						}
					}
					{
						if v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{}), "append"), v_5, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_6.UnsafePtr).V0)}))}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_2:
					return __t2
				}), m_4)
			}), mempty_1_0)
		})
	})
}

func Call_Control_Monad_Rec_Class_tailRec(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var Call_local_Control_Monad_Rec_Class_go__go_1_0_0 func(gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Control_Monad_Rec_Class_go__go_1_0_0
	var go__go_1_0_0 gopurs_runtime.Value
	_ = go__go_1_0_0
	Call_local_Control_Monad_Rec_Class_go__go_1_0_0 = func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	go__go_1_0_0:
		for {
			if false {
				continue go__go_1_0_0
			}
			var v_2 gopurs_runtime.Value = v_2_loop
			_ = v_2
			var __t1 gopurs_runtime.Value
			{
				if v_2.Type == 9 && v_2.IntVal == 525585346 {
					v_2_loop = gopurs_runtime.Apply(f_0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
					continue go__go_1_0_0
					__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
					goto end_branch_1
				} else {

				}
			}
			{
				if v_2.Type == 9 && v_2.IntVal == 60402430 {
					__t1 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Control_Monad_Rec_Class_go__go_1_0_0(v_2_loop_val)
	})
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_1_0_0, f_0)
}

func Call_Control_Monad_Rec_Class_tailRec2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply(Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(o_3, "a"), gopurs_runtime.RecordGet(o_3, "b"))
	})), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			/*gopurs:record-v1:2:0061,0062*/
		}{a_1, b_2}
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
	}())
}

func Call_Control_Monad_Rec_Class_tailRec3(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, c_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	var c_3 gopurs_runtime.Value = c_3_loop
	_ = c_3
	return gopurs_runtime.Apply(Call_Control_Monad_Rec_Class_tailRec(gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(f_0, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"), gopurs_runtime.RecordGet(o_4, "c"))
	})), func() gopurs_runtime.Value {
		orig := struct {
			a gopurs_runtime.Value
			b gopurs_runtime.Value
			c gopurs_runtime.Value
			/*gopurs:record-v1:3:0061,0062,0063*/
		}{a_1, b_2, c_3}
		_ = orig
		return gopurs_runtime.RecordDict3("a", "b", "c", orig.a, orig.b, orig.c)
	}())
}

func Call_Control_Monad_Rec_Class_loop3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var b_1 gopurs_runtime.Value = b_1_loop
	_ = b_1
	var c_2 gopurs_runtime.Value = c_2_loop
	_ = c_2
	return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_1030879592_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct {
		a gopurs_runtime.Value
		b gopurs_runtime.Value
		c gopurs_runtime.Value
		/*gopurs:record-v1:3:0061,0062,0063*/
	}, gopurs_runtime.Value]{1, struct {
		a gopurs_runtime.Value
		b gopurs_runtime.Value
		c gopurs_runtime.Value
		/*gopurs:record-v1:3:0061,0062,0063*/
	}{a_0, b_1, c_2}})))}
}

func Call_Control_Monad_Rec_Class_loop2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var b_1 gopurs_runtime.Value = b_1_loop
	_ = b_1
	return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_Rec_Class_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct {
		a gopurs_runtime.Value
		b gopurs_runtime.Value
		/*gopurs:record-v1:2:0061,0062*/
	}, gopurs_runtime.Value]{1, struct {
		a gopurs_runtime.Value
		b gopurs_runtime.Value
		/*gopurs:record-v1:2:0061,0062*/
	}{a_0, b_1}})))}
}

func Call_Control_Monad_Rec_Class_forever(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, u_3}))}
			}), ma_2)
		}), Get_Data_Unit_unit())
	})
}

func Rebox_Control_Monad_Rec_Class_1030879592_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct {
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	c gopurs_runtime.Value
	/*gopurs:record-v1:3:0061,0062,0063*/
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict3("a", "b", "c", orig.a, orig.b, orig.c)
	}()
	return out
}

func Rebox_Control_Monad_Rec_Class_1542299734_4130553207(in *Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_Rec_Class_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Monad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_Rec_Class_2568689657_1642601656(in *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_Rec_Class_3558361591_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[int64, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Monad_Rec_Class_45792715_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[int64, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Control_Monad_Rec_Class_784458114_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct {
	a gopurs_runtime.Value
	b gopurs_runtime.Value
	/*gopurs:record-v1:2:0061,0062*/
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
	}()
	return out
}
