package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_apply gopurs_runtime.Value
var once_Data_Enum_apply sync.Once

func Get_Data_Enum_apply() gopurs_runtime.Value {
	once_Data_Enum_apply.Do(func() {
		cache_Data_Enum_apply = gopurs_runtime.RecordGet(Get_Control_Apply_applyFn(), "apply")
	})
	return cache_Data_Enum_apply
}

var cache_Data_Enum_go__append gopurs_runtime.Value
var once_Data_Enum_go__append sync.Once

func Get_Data_Enum_go__append() gopurs_runtime.Value {
	once_Data_Enum_go__append.Do(func() {
		cache_Data_Enum_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Enum_go__append
}

var cache_Data_Enum_show gopurs_runtime.Value
var once_Data_Enum_show sync.Once

func Get_Data_Enum_show() gopurs_runtime.Value {
	once_Data_Enum_show.Do(func() {
		cache_Data_Enum_show = Get_Data_Show_show__1986550628()
	})
	return cache_Data_Enum_show
}

var cache_Data_Enum_lessThan gopurs_runtime.Value
var once_Data_Enum_lessThan sync.Once

func Get_Data_Enum_lessThan() gopurs_runtime.Value {
	once_Data_Enum_lessThan.Do(func() {
		cache_Data_Enum_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_Enum_lessThan
}

var cache_Data_Enum_go__map gopurs_runtime.Value
var once_Data_Enum_go__map sync.Once

func Get_Data_Enum_go__map() gopurs_runtime.Value {
	once_Data_Enum_go__map.Do(func() {
		cache_Data_Enum_go__map = Rebox_Data_Enum_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Enum_go__map
}

var cache_Data_Enum_compose gopurs_runtime.Value
var once_Data_Enum_compose sync.Once

func Get_Data_Enum_compose() gopurs_runtime.Value {
	once_Data_Enum_compose.Do(func() {
		cache_Data_Enum_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose
}

var cache_Data_Enum_compose1 gopurs_runtime.Value
var once_Data_Enum_compose1 sync.Once

func Get_Data_Enum_compose1() gopurs_runtime.Value {
	once_Data_Enum_compose1.Do(func() {
		cache_Data_Enum_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose1
}

var cache_Data_Enum_map1 gopurs_runtime.Value
var once_Data_Enum_map1 sync.Once

func Get_Data_Enum_map1() gopurs_runtime.Value {
	once_Data_Enum_map1.Do(func() {
		cache_Data_Enum_map1 = Rebox_Data_Enum_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Enum_map1
}

var cache_Data_Enum_top gopurs_runtime.Value
var once_Data_Enum_top sync.Once

func Get_Data_Enum_top() gopurs_runtime.Value {
	once_Data_Enum_top.Do(func() {
		cache_Data_Enum_top = gopurs_runtime.Int(gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal).IntVal)
	})
	return cache_Data_Enum_top
}

var cache_Data_Enum_add gopurs_runtime.Value
var once_Data_Enum_add sync.Once

func Get_Data_Enum_add() gopurs_runtime.Value {
	once_Data_Enum_add.Do(func() {
		cache_Data_Enum_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Enum_add
}

var cache_Data_Enum_greaterThan gopurs_runtime.Value
var once_Data_Enum_greaterThan sync.Once

func Get_Data_Enum_greaterThan() gopurs_runtime.Value {
	once_Data_Enum_greaterThan.Do(func() {
		cache_Data_Enum_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_Enum_greaterThan
}

var cache_Data_Enum_bottom gopurs_runtime.Value
var once_Data_Enum_bottom sync.Once

func Get_Data_Enum_bottom() gopurs_runtime.Value {
	once_Data_Enum_bottom.Do(func() {
		cache_Data_Enum_bottom = gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)
	})
	return cache_Data_Enum_bottom
}

var cache_Data_Enum_sub gopurs_runtime.Value
var once_Data_Enum_sub sync.Once

func Get_Data_Enum_sub() gopurs_runtime.Value {
	once_Data_Enum_sub.Do(func() {
		cache_Data_Enum_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Enum_sub
}

var cache_Data_Enum_bind gopurs_runtime.Value
var once_Data_Enum_bind sync.Once

func Get_Data_Enum_bind() gopurs_runtime.Value {
	once_Data_Enum_bind.Do(func() {
		cache_Data_Enum_bind = Rebox_Data_Enum_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1
	})
	return cache_Data_Enum_bind
}

var cache_Data_Enum_voidLeft gopurs_runtime.Value
var once_Data_Enum_voidLeft sync.Once

func Get_Data_Enum_voidLeft() gopurs_runtime.Value {
	once_Data_Enum_voidLeft.Do(func() {
		cache_Data_Enum_voidLeft = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_voidLeft(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](f_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_voidLeft
}

var cache_Data_Enum_guard gopurs_runtime.Value
var once_Data_Enum_guard sync.Once

func Get_Data_Enum_guard() gopurs_runtime.Value {
	once_Data_Enum_guard.Do(func() {
		cache_Data_Enum_guard = Get_Control_Alternative_guard__3382145287()
	})
	return cache_Data_Enum_guard
}

var cache_Data_Enum_lessThanOrEq gopurs_runtime.Value
var once_Data_Enum_lessThanOrEq sync.Once

func Get_Data_Enum_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Enum_lessThanOrEq.Do(func() {
		cache_Data_Enum_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Enum_lessThanOrEq
}

var cache_Data_Enum_composeFlipped gopurs_runtime.Value
var once_Data_Enum_composeFlipped sync.Once

func Get_Data_Enum_composeFlipped() gopurs_runtime.Value {
	once_Data_Enum_composeFlipped.Do(func() {
		cache_Data_Enum_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_composeFlipped
}

var cache_Data_Enum_fromJust gopurs_runtime.Value
var once_Data_Enum_fromJust sync.Once

func Get_Data_Enum_fromJust() gopurs_runtime.Value {
	once_Data_Enum_fromJust.Do(func() {
		cache_Data_Enum_fromJust = gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})
	})
	return cache_Data_Enum_fromJust
}

var cache_Data_Enum_compose2 gopurs_runtime.Value
var once_Data_Enum_compose2 sync.Once

func Get_Data_Enum_compose2() gopurs_runtime.Value {
	once_Data_Enum_compose2.Do(func() {
		cache_Data_Enum_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose2
}

var cache_Data_Enum_compose3 gopurs_runtime.Value
var once_Data_Enum_compose3 sync.Once

func Get_Data_Enum_compose3() gopurs_runtime.Value {
	once_Data_Enum_compose3.Do(func() {
		cache_Data_Enum_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose3
}

var cache_Data_Enum_compose4 gopurs_runtime.Value
var once_Data_Enum_compose4 sync.Once

func Get_Data_Enum_compose4() gopurs_runtime.Value {
	once_Data_Enum_compose4.Do(func() {
		cache_Data_Enum_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose4
}

var cache_Data_Enum_compose5 gopurs_runtime.Value
var once_Data_Enum_compose5 sync.Once

func Get_Data_Enum_compose5() gopurs_runtime.Value {
	once_Data_Enum_compose5.Do(func() {
		cache_Data_Enum_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose5
}

var cache_Data_Enum_apply1 gopurs_runtime.Value
var once_Data_Enum_apply1 sync.Once

func Get_Data_Enum_apply1() gopurs_runtime.Value {
	once_Data_Enum_apply1.Do(func() {
		cache_Data_Enum_apply1 = gopurs_runtime.RecordGet(Get_Control_Apply_applyFn(), "apply")
	})
	return cache_Data_Enum_apply1
}

var cache_Data_Enum_compose6 gopurs_runtime.Value
var once_Data_Enum_compose6 sync.Once

func Get_Data_Enum_compose6() gopurs_runtime.Value {
	once_Data_Enum_compose6.Do(func() {
		cache_Data_Enum_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose6
}

var cache_Data_Enum_map2 gopurs_runtime.Value
var once_Data_Enum_map2 sync.Once

func Get_Data_Enum_map2() gopurs_runtime.Value {
	once_Data_Enum_map2.Do(func() {
		cache_Data_Enum_map2 = Rebox_Data_Enum_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Enum_map2
}

var cache_Data_Enum_compose7 gopurs_runtime.Value
var once_Data_Enum_compose7 sync.Once

func Get_Data_Enum_compose7() gopurs_runtime.Value {
	once_Data_Enum_compose7.Do(func() {
		cache_Data_Enum_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Enum_compose7
}

var cache_Data_Enum_map3 gopurs_runtime.Value
var once_Data_Enum_map3 sync.Once

func Get_Data_Enum_map3() gopurs_runtime.Value {
	once_Data_Enum_map3.Do(func() {
		cache_Data_Enum_map3 = Rebox_Data_Enum_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0
	})
	return cache_Data_Enum_map3
}

var cache_Data_Enum_eq gopurs_runtime.Value
var once_Data_Enum_eq sync.Once

func Get_Data_Enum_eq() gopurs_runtime.Value {
	once_Data_Enum_eq.Do(func() {
		cache_Data_Enum_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_Enum_eq
}

var cache_Data_Enum_conj gopurs_runtime.Value
var once_Data_Enum_conj sync.Once

func Get_Data_Enum_conj() gopurs_runtime.Value {
	once_Data_Enum_conj.Do(func() {
		cache_Data_Enum_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Enum_conj
}

var cache_Data_Enum_greaterThanOrEq gopurs_runtime.Value
var once_Data_Enum_greaterThanOrEq sync.Once

func Get_Data_Enum_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Enum_greaterThanOrEq.Do(func() {
		cache_Data_Enum_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_Enum_greaterThanOrEq
}

var cache_Data_Enum_bottom1 gopurs_runtime.Value
var once_Data_Enum_bottom1 sync.Once

func Get_Data_Enum_bottom1() gopurs_runtime.Value {
	once_Data_Enum_bottom1.Do(func() {
		cache_Data_Enum_bottom1 = gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())
	})
	return cache_Data_Enum_bottom1
}

var cache_Data_Enum_top1 gopurs_runtime.Value
var once_Data_Enum_top1 sync.Once

func Get_Data_Enum_top1() gopurs_runtime.Value {
	once_Data_Enum_top1.Do(func() {
		cache_Data_Enum_top1 = gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar())
	})
	return cache_Data_Enum_top1
}

var cache_Data_Enum_Enum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict sync.Once

func Get_Data_Enum_Enum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Enum_dollar_Dict(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}()))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict
}

var cache_Data_Enum_Enum_dollar_Dict__2787884439 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__2787884439 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__2787884439() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__2787884439.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__2787884439 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Call_Data_Enum_Enum_dollar_Dict__2787884439(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__2787884439
}

var cache_Data_Enum_Enum_dollar_Dict__479058327 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__479058327 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__479058327() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__479058327.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__479058327 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Call_Data_Enum_Enum_dollar_Dict__479058327(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__479058327
}

var cache_Data_Enum_Enum_dollar_Dict__3712690007 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__3712690007 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__3712690007() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__3712690007.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__3712690007 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3494563625_556578094(Call_Data_Enum_Enum_dollar_Dict__3712690007(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__3712690007
}

var cache_Data_Enum_Enum_dollar_Dict__2123951319 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__2123951319 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__2123951319() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__2123951319.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__2123951319 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Call_Data_Enum_Enum_dollar_Dict__2123951319(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__2123951319
}

var cache_Data_Enum_Enum_dollar_Dict__1439164631 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__1439164631 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__1439164631() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__1439164631.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__1439164631 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094(Call_Data_Enum_Enum_dollar_Dict__1439164631(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__1439164631
}

var cache_Data_Enum_Enum_dollar_Dict__2220371927 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__2220371927 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__2220371927() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__2220371927.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__2220371927 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094(Call_Data_Enum_Enum_dollar_Dict__2220371927(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__2220371927
}

var cache_Data_Enum_Enum_dollar_Dict__3674609207 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__3674609207 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__3674609207() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__3674609207.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__3674609207 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4060049525_556578094(Call_Data_Enum_Enum_dollar_Dict__3674609207(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__3674609207
}

var cache_Data_Enum_Enum_dollar_Dict__4138904023 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__4138904023 sync.Once

func Get_Data_Enum_Enum_dollar_Dict__4138904023() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__4138904023.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__4138904023 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Enum_dollar_Dict__4138904023(func() struct {
				Ord0 gopurs_runtime.Value
				pred gopurs_runtime.Value
				succ gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0 gopurs_runtime.Value
					pred gopurs_runtime.Value
					succ gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.pred = gopurs_runtime.RecordGet(orig, "pred")
				clone.succ = gopurs_runtime.RecordGet(orig, "succ")
				return clone
			}()))}
		})
	})
	return cache_Data_Enum_Enum_dollar_Dict__4138904023
}

var cache_Data_Enum_Cardinality gopurs_runtime.Value
var once_Data_Enum_Cardinality sync.Once

func Get_Data_Enum_Cardinality() gopurs_runtime.Value {
	once_Data_Enum_Cardinality.Do(func() {
		cache_Data_Enum_Cardinality = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality
}

var cache_Data_Enum_Cardinality__1400572777 gopurs_runtime.Value
var once_Data_Enum_Cardinality__1400572777 sync.Once

func Get_Data_Enum_Cardinality__1400572777() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__1400572777.Do(func() {
		cache_Data_Enum_Cardinality__1400572777 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__1400572777(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__1400572777
}

var cache_Data_Enum_Cardinality__2274214265 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2274214265 sync.Once

func Get_Data_Enum_Cardinality__2274214265() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2274214265.Do(func() {
		cache_Data_Enum_Cardinality__2274214265 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__2274214265(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__2274214265
}

var cache_Data_Enum_Cardinality__123016565 gopurs_runtime.Value
var once_Data_Enum_Cardinality__123016565 sync.Once

func Get_Data_Enum_Cardinality__123016565() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__123016565.Do(func() {
		cache_Data_Enum_Cardinality__123016565 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__123016565(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__123016565
}

var cache_Data_Enum_Cardinality__1784762877 gopurs_runtime.Value
var once_Data_Enum_Cardinality__1784762877 sync.Once

func Get_Data_Enum_Cardinality__1784762877() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__1784762877.Do(func() {
		cache_Data_Enum_Cardinality__1784762877 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__1784762877(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__1784762877
}

var cache_Data_Enum_Cardinality__4276004359 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4276004359 sync.Once

func Get_Data_Enum_Cardinality__4276004359() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4276004359.Do(func() {
		cache_Data_Enum_Cardinality__4276004359 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__4276004359(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__4276004359
}

var cache_Data_Enum_Cardinality__2427762903 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2427762903 sync.Once

func Get_Data_Enum_Cardinality__2427762903() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2427762903.Do(func() {
		cache_Data_Enum_Cardinality__2427762903 = gopurs_runtime.Func(func(x_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__2427762903(x_unused_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__2427762903
}

var cache_Data_Enum_Cardinality__4110088112 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4110088112 sync.Once

func Get_Data_Enum_Cardinality__4110088112() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4110088112.Do(func() {
		cache_Data_Enum_Cardinality__4110088112 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__4110088112(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__4110088112
}

var cache_Data_Enum_Cardinality__4100296677 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4100296677 sync.Once

func Get_Data_Enum_Cardinality__4100296677() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4100296677.Do(func() {
		cache_Data_Enum_Cardinality__4100296677 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_Cardinality__4100296677(x_0_box.IntVal)
		})
	})
	return cache_Data_Enum_Cardinality__4100296677
}

var cache_Data_Enum_BoundedEnum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Call_Data_Enum_BoundedEnum_dollar_Dict(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}()))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__4237532871 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__4237532871 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__4237532871() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__4237532871.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__4237532871 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__4237532871(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__4237532871
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__2081071559 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__2081071559 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__2081071559() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__2081071559.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__2081071559 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__2081071559(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__2081071559
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__2130790279 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__2130790279 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__2130790279() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__2130790279.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__2130790279 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__2130790279(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__2130790279
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__1429463943 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__1429463943 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__1429463943() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__1429463943.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__1429463943 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3083804185_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__1429463943(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__1429463943
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__882291527 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__882291527 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__882291527() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__882291527.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__882291527 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3569500834_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__882291527(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__882291527
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__2897860935 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__2897860935 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__2897860935() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__2897860935.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__2897860935 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1306125126_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__2897860935(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}())))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__2897860935
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__3240687239 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__3240687239 sync.Once

func Get_Data_Enum_BoundedEnum_dollar_Dict__3240687239() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__3240687239.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__3240687239 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Call_Data_Enum_BoundedEnum_dollar_Dict__3240687239(func() struct {
				Bounded0    gopurs_runtime.Value
				Enum1       gopurs_runtime.Value
				cardinality gopurs_runtime.Value
				fromEnum    gopurs_runtime.Value
				toEnum      gopurs_runtime.Value
				/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Bounded0    gopurs_runtime.Value
					Enum1       gopurs_runtime.Value
					cardinality gopurs_runtime.Value
					fromEnum    gopurs_runtime.Value
					toEnum      gopurs_runtime.Value
					/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
				}{}
				clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
				clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
				clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
				clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
				clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
				return clone
			}()))}
		})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__3240687239
}

var cache_Data_Enum_toEnum gopurs_runtime.Value
var once_Data_Enum_toEnum sync.Once

func Get_Data_Enum_toEnum() gopurs_runtime.Value {
	once_Data_Enum_toEnum.Do(func() {
		cache_Data_Enum_toEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_toEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Enum_toEnum
}

var cache_Data_Enum_toEnum__4195956870 gopurs_runtime.Value
var once_Data_Enum_toEnum__4195956870 sync.Once

func Get_Data_Enum_toEnum__4195956870() gopurs_runtime.Value {
	once_Data_Enum_toEnum__4195956870.Do(func() {
		cache_Data_Enum_toEnum__4195956870 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__4195956870(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__4195956870
}

var cache_Data_Enum_toEnum__890235542 gopurs_runtime.Value
var once_Data_Enum_toEnum__890235542 sync.Once

func Get_Data_Enum_toEnum__890235542() gopurs_runtime.Value {
	once_Data_Enum_toEnum__890235542.Do(func() {
		cache_Data_Enum_toEnum__890235542 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__890235542(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__890235542
}

var cache_Data_Enum_charToEnum gopurs_runtime.Value
var once_Data_Enum_charToEnum sync.Once

func Get_Data_Enum_charToEnum() gopurs_runtime.Value {
	once_Data_Enum_charToEnum.Do(func() {
		cache_Data_Enum_charToEnum = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(v_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_charToEnum
}

var cache_Data_Enum_defaultSucc__3855932694 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__3855932694 sync.Once

func Get_Data_Enum_defaultSucc__3855932694() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__3855932694.Do(func() {
		cache_Data_Enum_defaultSucc__3855932694 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc__3855932694(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_defaultSucc__3855932694
}

var cache_Data_Enum_defaultPred__3855932694 gopurs_runtime.Value
var once_Data_Enum_defaultPred__3855932694 sync.Once

func Get_Data_Enum_defaultPred__3855932694() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__3855932694.Do(func() {
		cache_Data_Enum_defaultPred__3855932694 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred__3855932694(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_defaultPred__3855932694
}

var cache_Data_Enum_enumChar gopurs_runtime.Value
var once_Data_Enum_enumChar sync.Once

func Get_Data_Enum_enumChar() gopurs_runtime.Value {
	once_Data_Enum_enumChar.Do(func() {
		cache_Data_Enum_enumChar = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094((&Constructor_Data_Enum_Enum[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2406510097_4177771502(Rebox_Data_Enum_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar()))))}
		}), gopurs_runtime.Apply2(Get_Data_Enum_defaultPred__3855932694(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()), gopurs_runtime.Apply2(Get_Data_Enum_defaultSucc__3855932694(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode())})))}
	})
	return cache_Data_Enum_enumChar
}

var cache_Data_Enum_boundedEnumChar gopurs_runtime.Value
var once_Data_Enum_boundedEnumChar sync.Once

func Get_Data_Enum_boundedEnumChar() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumChar.Do(func() {
		cache_Data_Enum_boundedEnumChar = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3569500834_123048125((&Constructor_Data_Enum_BoundedEnum[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094(Rebox_Data_Enum_556578094_3409577041(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumChar()))))}
		}), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar())).IntVal, gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())).IntVal)).IntVal, Get_Data_Enum_toCharCode(), Get_Data_Enum_charToEnum()})))}
	})
	return cache_Data_Enum_boundedEnumChar
}

var cache_Data_Enum_toEnum__1787937256 gopurs_runtime.Value
var once_Data_Enum_toEnum__1787937256 sync.Once

func Get_Data_Enum_toEnum__1787937256() gopurs_runtime.Value {
	once_Data_Enum_toEnum__1787937256.Do(func() {
		cache_Data_Enum_toEnum__1787937256 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__1787937256(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__1787937256
}

var cache_Data_Enum_toEnum__706605770 gopurs_runtime.Value
var once_Data_Enum_toEnum__706605770 sync.Once

func Get_Data_Enum_toEnum__706605770() gopurs_runtime.Value {
	once_Data_Enum_toEnum__706605770.Do(func() {
		cache_Data_Enum_toEnum__706605770 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__706605770(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__706605770
}

var cache_Data_Enum_toEnum__1693208633 gopurs_runtime.Value
var once_Data_Enum_toEnum__1693208633 sync.Once

func Get_Data_Enum_toEnum__1693208633() gopurs_runtime.Value {
	once_Data_Enum_toEnum__1693208633.Do(func() {
		cache_Data_Enum_toEnum__1693208633 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__1693208633(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__1693208633
}

var cache_Data_Enum_toEnum__6210423 gopurs_runtime.Value
var once_Data_Enum_toEnum__6210423 sync.Once

func Get_Data_Enum_toEnum__6210423() gopurs_runtime.Value {
	once_Data_Enum_toEnum__6210423.Do(func() {
		cache_Data_Enum_toEnum__6210423 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__6210423(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__6210423
}

var cache_Data_Enum_toEnum__2413733450 gopurs_runtime.Value
var once_Data_Enum_toEnum__2413733450 sync.Once

func Get_Data_Enum_toEnum__2413733450() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2413733450.Do(func() {
		cache_Data_Enum_toEnum__2413733450 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__2413733450(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__2413733450
}

var cache_Data_Enum_toEnum__72020153 gopurs_runtime.Value
var once_Data_Enum_toEnum__72020153 sync.Once

func Get_Data_Enum_toEnum__72020153() gopurs_runtime.Value {
	once_Data_Enum_toEnum__72020153.Do(func() {
		cache_Data_Enum_toEnum__72020153 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__72020153(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__72020153
}

var cache_Data_Enum_toEnum__3521034119 gopurs_runtime.Value
var once_Data_Enum_toEnum__3521034119 sync.Once

func Get_Data_Enum_toEnum__3521034119() gopurs_runtime.Value {
	once_Data_Enum_toEnum__3521034119.Do(func() {
		cache_Data_Enum_toEnum__3521034119 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__3521034119(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_toEnum__3521034119
}

var cache_Data_Enum_succ gopurs_runtime.Value
var once_Data_Enum_succ sync.Once

func Get_Data_Enum_succ() gopurs_runtime.Value {
	once_Data_Enum_succ.Do(func() {
		cache_Data_Enum_succ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_succ(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Enum_succ
}

var cache_Data_Enum_succ__77857599 gopurs_runtime.Value
var once_Data_Enum_succ__77857599 sync.Once

func Get_Data_Enum_succ__77857599() gopurs_runtime.Value {
	once_Data_Enum_succ__77857599.Do(func() {
		cache_Data_Enum_succ__77857599 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__77857599(uint32(__eta0_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_succ__77857599
}

var cache_Data_Enum_succ__545117358 gopurs_runtime.Value
var once_Data_Enum_succ__545117358 sync.Once

func Get_Data_Enum_succ__545117358() gopurs_runtime.Value {
	once_Data_Enum_succ__545117358.Do(func() {
		cache_Data_Enum_succ__545117358 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__545117358(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__eta0_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_succ__545117358
}

var cache_Data_Enum_succ__3573123786 gopurs_runtime.Value
var once_Data_Enum_succ__3573123786 sync.Once

func Get_Data_Enum_succ__3573123786() gopurs_runtime.Value {
	once_Data_Enum_succ__3573123786.Do(func() {
		cache_Data_Enum_succ__3573123786 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__3573123786(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_succ__3573123786
}

var cache_Data_Enum_succ__2729625849 gopurs_runtime.Value
var once_Data_Enum_succ__2729625849 sync.Once

func Get_Data_Enum_succ__2729625849() gopurs_runtime.Value {
	once_Data_Enum_succ__2729625849.Do(func() {
		cache_Data_Enum_succ__2729625849 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__2729625849(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_succ__2729625849
}

var cache_Data_Enum_upFromIncluding gopurs_runtime.Value
var once_Data_Enum_upFromIncluding sync.Once

func Get_Data_Enum_upFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_upFromIncluding.Do(func() {
		cache_Data_Enum_upFromIncluding = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_upFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
		})
	})
	return cache_Data_Enum_upFromIncluding
}

var cache_Data_Enum_showCardinality gopurs_runtime.Value
var once_Data_Enum_showCardinality sync.Once

func Get_Data_Enum_showCardinality() gopurs_runtime.Value {
	once_Data_Enum_showCardinality.Do(func() {
		cache_Data_Enum_showCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Cardinality ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})}))}
	})
	return cache_Data_Enum_showCardinality
}

var cache_Data_Enum_pred gopurs_runtime.Value
var once_Data_Enum_pred sync.Once

func Get_Data_Enum_pred() gopurs_runtime.Value {
	once_Data_Enum_pred.Do(func() {
		cache_Data_Enum_pred = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_pred(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Enum_pred
}

var cache_Data_Enum_pred__77857599 gopurs_runtime.Value
var once_Data_Enum_pred__77857599 sync.Once

func Get_Data_Enum_pred__77857599() gopurs_runtime.Value {
	once_Data_Enum_pred__77857599.Do(func() {
		cache_Data_Enum_pred__77857599 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred__77857599(uint32(__eta0_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_pred__77857599
}

var cache_Data_Enum_pred__3573123786 gopurs_runtime.Value
var once_Data_Enum_pred__3573123786 sync.Once

func Get_Data_Enum_pred__3573123786() gopurs_runtime.Value {
	once_Data_Enum_pred__3573123786.Do(func() {
		cache_Data_Enum_pred__3573123786 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred__3573123786(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_pred__3573123786
}

var cache_Data_Enum_pred__2729625849 gopurs_runtime.Value
var once_Data_Enum_pred__2729625849 sync.Once

func Get_Data_Enum_pred__2729625849() gopurs_runtime.Value {
	once_Data_Enum_pred__2729625849.Do(func() {
		cache_Data_Enum_pred__2729625849 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred__2729625849(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_pred__2729625849
}

var cache_Data_Enum_ordCardinality gopurs_runtime.Value
var once_Data_Enum_ordCardinality sync.Once

func Get_Data_Enum_ordCardinality() gopurs_runtime.Value {
	once_Data_Enum_ordCardinality.Do(func() {
		cache_Data_Enum_ordCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3308271157_4177771502(Rebox_Data_Enum_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Enum_ordCardinality
}

var cache_Data_Enum_newtypeCardinality gopurs_runtime.Value
var once_Data_Enum_newtypeCardinality sync.Once

func Get_Data_Enum_newtypeCardinality() gopurs_runtime.Value {
	once_Data_Enum_newtypeCardinality.Do(func() {
		cache_Data_Enum_newtypeCardinality = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2621187955_385277032((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Enum_newtypeCardinality
}

var cache_Data_Enum_fromEnum gopurs_runtime.Value
var once_Data_Enum_fromEnum sync.Once

func Get_Data_Enum_fromEnum() gopurs_runtime.Value {
	once_Data_Enum_fromEnum.Do(func() {
		cache_Data_Enum_fromEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_fromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Enum_fromEnum
}

var cache_Data_Enum_fromEnum__798184167 gopurs_runtime.Value
var once_Data_Enum_fromEnum__798184167 sync.Once

func Get_Data_Enum_fromEnum__798184167() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__798184167.Do(func() {
		cache_Data_Enum_fromEnum__798184167 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__798184167(uint32(__eta0_0_box.IntVal)))
		})
	})
	return cache_Data_Enum_fromEnum__798184167
}

var cache_Data_Enum_fromEnum__1526214825 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1526214825 sync.Once

func Get_Data_Enum_fromEnum__1526214825() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1526214825.Do(func() {
		cache_Data_Enum_fromEnum__1526214825 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__1526214825(__eta0_0_box.StrVal()))
		})
	})
	return cache_Data_Enum_fromEnum__1526214825
}

var cache_Data_Enum_fromEnum__1795542731 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1795542731 sync.Once

func Get_Data_Enum_fromEnum__1795542731() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1795542731.Do(func() {
		cache_Data_Enum_fromEnum__1795542731 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__1795542731(__eta0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__1795542731
}

var cache_Data_Enum_fromEnum__2307775192 gopurs_runtime.Value
var once_Data_Enum_fromEnum__2307775192 sync.Once

func Get_Data_Enum_fromEnum__2307775192() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__2307775192.Do(func() {
		cache_Data_Enum_fromEnum__2307775192 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__2307775192(__eta0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__2307775192
}

var cache_Data_Enum_fromEnum__3185865430 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3185865430 sync.Once

func Get_Data_Enum_fromEnum__3185865430() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3185865430.Do(func() {
		cache_Data_Enum_fromEnum__3185865430 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3185865430(__eta_norm_0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__3185865430
}

var cache_Data_Enum_fromEnum__862809814 gopurs_runtime.Value
var once_Data_Enum_fromEnum__862809814 sync.Once

func Get_Data_Enum_fromEnum__862809814() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__862809814.Do(func() {
		cache_Data_Enum_fromEnum__862809814 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__862809814(__eta0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__862809814
}

var cache_Data_Enum_fromEnum__3415298635 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3415298635 sync.Once

func Get_Data_Enum_fromEnum__3415298635() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3415298635.Do(func() {
		cache_Data_Enum_fromEnum__3415298635 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3415298635(__eta0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__3415298635
}

var cache_Data_Enum_fromEnum__1091906392 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1091906392 sync.Once

func Get_Data_Enum_fromEnum__1091906392() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1091906392.Do(func() {
		cache_Data_Enum_fromEnum__1091906392 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__1091906392(__eta0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__1091906392
}

var cache_Data_Enum_fromEnum__4269580326 gopurs_runtime.Value
var once_Data_Enum_fromEnum__4269580326 sync.Once

func Get_Data_Enum_fromEnum__4269580326() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__4269580326.Do(func() {
		cache_Data_Enum_fromEnum__4269580326 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Enum_fromEnum__4269580326(__eta0_0_box.IntVal))
		})
	})
	return cache_Data_Enum_fromEnum__4269580326
}

var cache_Data_Enum_toEnumWithDefaults gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults sync.Once

func Get_Data_Enum_toEnumWithDefaults() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults.Do(func() {
		cache_Data_Enum_toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_toEnumWithDefaults(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_0_box))
		})
	})
	return cache_Data_Enum_toEnumWithDefaults
}

var cache_Data_Enum_toEnumWithDefaults__3799330569 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__3799330569 sync.Once

func Get_Data_Enum_toEnumWithDefaults__3799330569() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__3799330569.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__3799330569 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Enum_toEnumWithDefaults__3799330569(__eta0_0_box.StrVal(), __eta1_1_box.StrVal(), __eta2_2_box.IntVal))
		})
	})
	return cache_Data_Enum_toEnumWithDefaults__3799330569
}

var cache_Data_Enum_toEnumWithDefaults__4060784579 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__4060784579 sync.Once

func Get_Data_Enum_toEnumWithDefaults__4060784579() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__4060784579.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__4060784579 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Enum_toEnumWithDefaults__4060784579(__eta_norm_1_unused_0_box.StrVal(), __eta_norm_0_unused_1_box.StrVal(), x_2_box.IntVal))
		})
	})
	return cache_Data_Enum_toEnumWithDefaults__4060784579
}

var cache_Data_Enum_toEnumWithDefaults__92386535 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__92386535 sync.Once

func Get_Data_Enum_toEnumWithDefaults__92386535() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__92386535.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__92386535 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Enum_toEnumWithDefaults__92386535(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box.StrVal(), x_2_box.IntVal))
		})
	})
	return cache_Data_Enum_toEnumWithDefaults__92386535
}

var cache_Data_Enum_eqCardinality gopurs_runtime.Value
var once_Data_Enum_eqCardinality sync.Once

func Get_Data_Enum_eqCardinality() gopurs_runtime.Value {
	once_Data_Enum_eqCardinality.Do(func() {
		cache_Data_Enum_eqCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1053099733_3790796878(Rebox_Data_Enum_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Enum_eqCardinality
}

var cache_Data_Enum_enumUnit gopurs_runtime.Value
var once_Data_Enum_enumUnit sync.Once

func Get_Data_Enum_enumUnit() gopurs_runtime.Value {
	once_Data_Enum_enumUnit.Do(func() {
		cache_Data_Enum_enumUnit = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Enum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordUnit()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		})}))}
	})
	return cache_Data_Enum_enumUnit
}

var cache_Data_Enum_enumTuple gopurs_runtime.Value
var once_Data_Enum_enumTuple sync.Once

func Get_Data_Enum_enumTuple() gopurs_runtime.Value {
	once_Data_Enum_enumTuple.Do(func() {
		cache_Data_Enum_enumTuple = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_enumTuple(dictEnum_0_box)
		})
	})
	return cache_Data_Enum_enumTuple
}

var cache_Data_Enum_enumOrdering gopurs_runtime.Value
var once_Data_Enum_enumOrdering sync.Once

func Get_Data_Enum_enumOrdering() gopurs_runtime.Value {
	once_Data_Enum_enumOrdering.Do(func() {
		cache_Data_Enum_enumOrdering = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094((&Constructor_Data_Enum_Enum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3730953251_4177771502(Rebox_Data_Enum_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordOrdering()))))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1527465420 {
					__t3 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 902936544 {
					__t3 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, true}
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 380165415 {
					__t3 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() struct {
					V0 gopurs_runtime.Value
					V1 bool
				} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(func() *Constructor_Data_Maybe_Just[uint32] { panic("Failed pattern match") }()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{V0: gopurs_runtime.Value{}, V1: false} }()
			}
		end_branch_3:
			return func() gopurs_runtime.Value {
				_v := __t3
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t7 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				var __t_tag_4 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 1527465420 {
					__t7 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 902936544 {
					__t7 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, true}
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_6 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_6
				if uint32(__t_tag_6) == 380165415 {
					__t7 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = func() struct {
					V0 gopurs_runtime.Value
					V1 bool
				} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(func() *Constructor_Data_Maybe_Just[uint32] { panic("Failed pattern match") }()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
				}; return struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{V0: gopurs_runtime.Value{}, V1: false} }()
			}
		end_branch_7:
			return func() gopurs_runtime.Value {
				_v := __t7
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_Enum_enumOrdering
}

var cache_Data_Enum_enumMaybe gopurs_runtime.Value
var once_Data_Enum_enumMaybe sync.Once

func Get_Data_Enum_enumMaybe() gopurs_runtime.Value {
	once_Data_Enum_enumMaybe.Do(func() {
		cache_Data_Enum_enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_enumMaybe(dictBoundedEnum_0_box)
		})
	})
	return cache_Data_Enum_enumMaybe
}

var cache_Data_Enum_enumInt gopurs_runtime.Value
var once_Data_Enum_enumInt sync.Once

func Get_Data_Enum_enumInt() gopurs_runtime.Value {
	once_Data_Enum_enumInt.Do(func() {
		cache_Data_Enum_enumInt = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3308271157_4177771502(Rebox_Data_Enum_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (n_0.IntVal) > (Get_Data_Bounded_bottomInt().IntVal) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(gopurs_runtime.IntSub(n_0.IntVal, int64(1))), true}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (n_0.IntVal) < (gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal).IntVal) {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(gopurs_runtime.IntAdd(n_0.IntVal, int64(1))), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_Enum_enumInt
}

var cache_Data_Enum_enumFromTo gopurs_runtime.Value
var once_Data_Enum_enumFromTo sync.Once

func Get_Data_Enum_enumFromTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromTo.Do(func() {
		cache_Data_Enum_enumFromTo = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_enumFromTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
		})
	})
	return cache_Data_Enum_enumFromTo
}

var cache_Data_Enum_enumFromThenTo gopurs_runtime.Value
var once_Data_Enum_enumFromThenTo sync.Once

func Get_Data_Enum_enumFromThenTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromThenTo.Do(func() {
		cache_Data_Enum_enumFromThenTo = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictBoundedEnum_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_enumFromThenTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_2_box))
		})
	})
	return cache_Data_Enum_enumFromThenTo
}

var cache_Data_Enum_enumEither gopurs_runtime.Value
var once_Data_Enum_enumEither sync.Once

func Get_Data_Enum_enumEither() gopurs_runtime.Value {
	once_Data_Enum_enumEither.Do(func() {
		cache_Data_Enum_enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_enumEither(dictBoundedEnum_0_box)
		})
	})
	return cache_Data_Enum_enumEither
}

var cache_Data_Enum_enumBoolean gopurs_runtime.Value
var once_Data_Enum_enumBoolean sync.Once

func Get_Data_Enum_enumBoolean() gopurs_runtime.Value {
	once_Data_Enum_enumBoolean.Do(func() {
		cache_Data_Enum_enumBoolean = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094((&Constructor_Data_Enum_Enum[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_219188042_4177771502(Rebox_Data_Enum_4177771502_219188042(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordBoolean()))))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (v_0.IntVal) != (0) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Bool(false), true}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((v_0.IntVal) != (0)) != (true) {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Bool(true), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_Enum_enumBoolean
}

var cache_Data_Enum_downFromIncluding gopurs_runtime.Value
var once_Data_Enum_downFromIncluding sync.Once

func Get_Data_Enum_downFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_downFromIncluding.Do(func() {
		cache_Data_Enum_downFromIncluding = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_downFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
		})
	})
	return cache_Data_Enum_downFromIncluding
}

var cache_Data_Enum_diag gopurs_runtime.Value
var once_Data_Enum_diag sync.Once

func Get_Data_Enum_diag() gopurs_runtime.Value {
	once_Data_Enum_diag.Do(func() {
		cache_Data_Enum_diag = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_diag(a_0_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Enum_diag
}

var cache_Data_Enum_downFrom gopurs_runtime.Value
var once_Data_Enum_downFrom sync.Once

func Get_Data_Enum_downFrom() gopurs_runtime.Value {
	once_Data_Enum_downFrom.Do(func() {
		cache_Data_Enum_downFrom = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_downFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
		})
	})
	return cache_Data_Enum_downFrom
}

var cache_Data_Enum_upFrom gopurs_runtime.Value
var once_Data_Enum_upFrom sync.Once

func Get_Data_Enum_upFrom() gopurs_runtime.Value {
	once_Data_Enum_upFrom.Do(func() {
		cache_Data_Enum_upFrom = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_upFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
		})
	})
	return cache_Data_Enum_upFrom
}

var cache_Data_Enum_defaultToEnum gopurs_runtime.Value
var once_Data_Enum_defaultToEnum sync.Once

func Get_Data_Enum_defaultToEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultToEnum.Do(func() {
		cache_Data_Enum_defaultToEnum = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_defaultToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0_box))
		})
	})
	return cache_Data_Enum_defaultToEnum
}

var cache_Data_Enum_defaultSucc gopurs_runtime.Value
var once_Data_Enum_defaultSucc sync.Once

func Get_Data_Enum_defaultSucc() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc.Do(func() {
		cache_Data_Enum_defaultSucc = gopurs_runtime.Func3(func(toEnum_prime__0_box gopurs_runtime.Value, fromEnum_prime__1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc(toEnum_prime__0_box, fromEnum_prime__1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_defaultSucc
}

var cache_Data_Enum_defaultSucc__2406132389 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__2406132389 sync.Once

func Get_Data_Enum_defaultSucc__2406132389() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__2406132389.Do(func() {
		cache_Data_Enum_defaultSucc__2406132389 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc__2406132389(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_defaultSucc__2406132389
}

var cache_Data_Enum_defaultPred gopurs_runtime.Value
var once_Data_Enum_defaultPred sync.Once

func Get_Data_Enum_defaultPred() gopurs_runtime.Value {
	once_Data_Enum_defaultPred.Do(func() {
		cache_Data_Enum_defaultPred = gopurs_runtime.Func3(func(toEnum_prime__0_box gopurs_runtime.Value, fromEnum_prime__1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred(toEnum_prime__0_box, fromEnum_prime__1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_defaultPred
}

var cache_Data_Enum_defaultPred__2406132389 gopurs_runtime.Value
var once_Data_Enum_defaultPred__2406132389 sync.Once

func Get_Data_Enum_defaultPred__2406132389() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__2406132389.Do(func() {
		cache_Data_Enum_defaultPred__2406132389 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred__2406132389(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Enum_defaultPred__2406132389
}

var cache_Data_Enum_defaultFromEnum gopurs_runtime.Value
var once_Data_Enum_defaultFromEnum sync.Once

func Get_Data_Enum_defaultFromEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultFromEnum.Do(func() {
		cache_Data_Enum_defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_defaultFromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
		})
	})
	return cache_Data_Enum_defaultFromEnum
}

var cache_Data_Enum_defaultCardinality gopurs_runtime.Value
var once_Data_Enum_defaultCardinality sync.Once

func Get_Data_Enum_defaultCardinality() gopurs_runtime.Value {
	once_Data_Enum_defaultCardinality.Do(func() {
		cache_Data_Enum_defaultCardinality = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_defaultCardinality(dictBounded_0_box)
		})
	})
	return cache_Data_Enum_defaultCardinality
}

var cache_Data_Enum_cardinality gopurs_runtime.Value
var once_Data_Enum_cardinality sync.Once

func Get_Data_Enum_cardinality() gopurs_runtime.Value {
	once_Data_Enum_cardinality.Do(func() {
		cache_Data_Enum_cardinality = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Enum_cardinality(dict_0_box)
		})
	})
	return cache_Data_Enum_cardinality
}

var cache_Data_Enum_boundedEnumUnit gopurs_runtime.Value
var once_Data_Enum_boundedEnumUnit sync.Once

func Get_Data_Enum_boundedEnumUnit() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumUnit.Do(func() {
		cache_Data_Enum_boundedEnumUnit = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedUnit()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumUnit()))}
		}), gopurs_runtime.Int(int64(1)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (v_0.IntVal) == (int64(0)) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{Get_Data_Unit_unit(), true}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})}))}
	})
	return cache_Data_Enum_boundedEnumUnit
}

var cache_Data_Enum_boundedEnumOrdering gopurs_runtime.Value
var once_Data_Enum_boundedEnumOrdering sync.Once

func Get_Data_Enum_boundedEnumOrdering() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumOrdering.Do(func() {
		cache_Data_Enum_boundedEnumOrdering = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125((&Constructor_Data_Enum_BoundedEnum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_832288803_2094947566(Rebox_Data_Enum_2094947566_832288803(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedOrdering()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Rebox_Data_Enum_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumOrdering()))))}
		}), gopurs_runtime.Int(int64(3)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1527465420 {
					__t3 = gopurs_runtime.Int(int64(0))
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 902936544 {
					__t3 = gopurs_runtime.Int(int64(1))
					goto end_branch_3
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 380165415 {
					__t3 = gopurs_runtime.Int(int64(2))
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
			}
		end_branch_3:
			return __t3
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (v_0.IntVal) == (int64(0)) {
					__t4 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, true}
					goto end_branch_4
				} else {

				}
			}
			{
				if (v_0.IntVal) == (int64(1)) {
					__t4 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
					goto end_branch_4
				} else {

				}
			}
			{
				if (v_0.IntVal) == (int64(2)) {
					__t4 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, true}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_4:
			return func() gopurs_runtime.Value {
				_v := __t4
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_Enum_boundedEnumOrdering
}

var cache_Data_Enum_boundedEnumBoolean gopurs_runtime.Value
var once_Data_Enum_boundedEnumBoolean sync.Once

func Get_Data_Enum_boundedEnumBoolean() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumBoolean.Do(func() {
		cache_Data_Enum_boundedEnumBoolean = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3083804185_123048125((&Constructor_Data_Enum_BoundedEnum[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3728870730_2094947566(Rebox_Data_Enum_2094947566_3728870730(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedBoolean()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094(Rebox_Data_Enum_556578094_342165130(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumBoolean()))))}
		}), gopurs_runtime.Int(int64(2)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if ((v_0.IntVal) != (0)) != (true) {
					__t0 = gopurs_runtime.Int(int64(0))
					goto end_branch_0
				} else {

				}
			}
			{
				if (v_0.IntVal) != (0) {
					__t0 = gopurs_runtime.Int(int64(1))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Int(func() int64 { panic("Failed pattern match") }())
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if (v_0.IntVal) == (int64(0)) {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Bool(false), true}
					goto end_branch_1
				} else {

				}
			}
			{
				if (v_0.IntVal) == (int64(1)) {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Bool(true), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})})))}
	})
	return cache_Data_Enum_boundedEnumBoolean
}

type Constructor_Data_Enum_Enum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0045006e0075006d:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0045006e0075006d005f0045006e0075006d
func init() {
	gopurs_runtime.StructGetters[4075786298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Ord0":
			return gopurs_runtime.Box(c.V0)
		case "pred":
			return gopurs_runtime.Box(c.V1)
		case "succ":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Enum_Enum: " + key)
		}
	}
}

type Constructor_Data_Enum_BoundedEnum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 int64
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0045006e0075006d:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0045006e0075006d005f0042006f0075006e0064006500640045006e0075006d
func init() {
	gopurs_runtime.StructGetters[287434377] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Bounded0":
			return gopurs_runtime.Box(c.V0)
		case "Enum1":
			return gopurs_runtime.Box(c.V1)
		case "cardinality":
			return gopurs_runtime.Box(c.V2)
		case "fromEnum":
			return gopurs_runtime.Box(c.V3)
		case "toEnum":
			return gopurs_runtime.Box(c.V4)
		default:
			panic("Key not found in dictionary Constructor_Data_Enum_BoundedEnum: " + key)
		}
	}
}

func Call_Data_Enum_voidLeft(f_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var f_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	{
		if f_0 != nil {
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{x_1, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Enum_Enum_dollar_Dict(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	var x_0 struct {
		Ord0 gopurs_runtime.Value
		pred gopurs_runtime.Value
		succ gopurs_runtime.Value
		/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
	}())
}

func Call_Data_Enum_Enum_dollar_Dict__2787884439(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[uint32] {
Enum_dollar_Dict__2787884439:
	for {
		if false {
			continue Enum_dollar_Dict__2787884439
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__479058327(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[uint32] {
Enum_dollar_Dict__479058327:
	for {
		if false {
			continue Enum_dollar_Dict__479058327
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__3712690007(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date] {
Enum_dollar_Dict__3712690007:
	for {
		if false {
			continue Enum_dollar_Dict__3712690007
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__2123951319(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[uint32] {
Enum_dollar_Dict__2123951319:
	for {
		if false {
			continue Enum_dollar_Dict__2123951319
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__1439164631(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[bool] {
Enum_dollar_Dict__1439164631:
	for {
		if false {
			continue Enum_dollar_Dict__1439164631
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[bool]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__2220371927(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[string] {
Enum_dollar_Dict__2220371927:
	for {
		if false {
			continue Enum_dollar_Dict__2220371927
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__3674609207(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[int64] {
Enum_dollar_Dict__3674609207:
	for {
		if false {
			continue Enum_dollar_Dict__3674609207
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[int64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Enum_dollar_Dict__4138904023(x_0_loop struct {
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
}) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
Enum_dollar_Dict__4138904023:
	for {
		if false {
			continue Enum_dollar_Dict__4138904023
		}
		var x_0 struct {
			Ord0 gopurs_runtime.Value
			pred gopurs_runtime.Value
			succ gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0070007200650064,0073007500630063*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
		}())
	}
}

func Call_Data_Enum_Cardinality(x_0_loop int64) gopurs_runtime.Value {
	var x_0 int64 = x_0_loop
	_ = x_0
	return gopurs_runtime.Int(x_0)
}

func Call_Data_Enum_Cardinality__1400572777(x_0_loop int64) gopurs_runtime.Value {
Cardinality__1400572777:
	for {
		if false {
			continue Cardinality__1400572777
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_Cardinality__2274214265(x_0_loop int64) gopurs_runtime.Value {
Cardinality__2274214265:
	for {
		if false {
			continue Cardinality__2274214265
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_Cardinality__123016565(x_0_loop int64) gopurs_runtime.Value {
Cardinality__123016565:
	for {
		if false {
			continue Cardinality__123016565
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_Cardinality__1784762877(x_0_loop int64) gopurs_runtime.Value {
Cardinality__1784762877:
	for {
		if false {
			continue Cardinality__1784762877
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_Cardinality__4276004359(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4276004359:
	for {
		if false {
			continue Cardinality__4276004359
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_Cardinality__2427762903(x_unused_0_loop int64) gopurs_runtime.Value {
Cardinality__2427762903:
	for {
		if false {
			continue Cardinality__2427762903
		}
		var x_unused_0 int64 = x_unused_0_loop
		_ = x_unused_0
		return gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar())).IntVal, gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())).IntVal))
	}
}

func Call_Data_Enum_Cardinality__4110088112(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4110088112:
	for {
		if false {
			continue Cardinality__4110088112
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_Cardinality__4100296677(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4100296677:
	for {
		if false {
			continue Cardinality__4100296677
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return gopurs_runtime.Int(x_0)
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	var x_0 struct {
		Bounded0    gopurs_runtime.Value
		Enum1       gopurs_runtime.Value
		cardinality gopurs_runtime.Value
		fromEnum    gopurs_runtime.Value
		toEnum      gopurs_runtime.Value
		/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
	}())
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__4237532871(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[uint32] {
BoundedEnum_dollar_Dict__4237532871:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__4237532871
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__2081071559(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[uint32] {
BoundedEnum_dollar_Dict__2081071559:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__2081071559
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__2130790279(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[uint32] {
BoundedEnum_dollar_Dict__2130790279:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__2130790279
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__1429463943(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[bool] {
BoundedEnum_dollar_Dict__1429463943:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__1429463943
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[bool]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__882291527(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[string] {
BoundedEnum_dollar_Dict__882291527:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__882291527
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__2897860935(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[int64] {
BoundedEnum_dollar_Dict__2897860935:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__2897860935
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__3240687239(x_0_loop struct {
	Bounded0    gopurs_runtime.Value
	Enum1       gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum    gopurs_runtime.Value
	toEnum      gopurs_runtime.Value
	/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
}) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
BoundedEnum_dollar_Dict__3240687239:
	for {
		if false {
			continue BoundedEnum_dollar_Dict__3240687239
		}
		var x_0 struct {
			Bounded0    gopurs_runtime.Value
			Enum1       gopurs_runtime.Value
			cardinality gopurs_runtime.Value
			fromEnum    gopurs_runtime.Value
			toEnum      gopurs_runtime.Value
			/*gopurs:record-v1:5:0042006f0075006e0064006500640030,0045006e0075006d0031,00630061007200640069006e0061006c006900740079,00660072006f006d0045006e0075006d,0074006f0045006e0075006d*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
		}())
	}
}

func Call_Data_Enum_toEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V4
}

func Call_Data_Enum_toEnum__4195956870(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__4195956870:
	for {
		if false {
			continue toEnum__4195956870
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__890235542(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__890235542:
	for {
		if false {
			continue toEnum__890235542
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumWeekday())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_charToEnum(v_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 int64 = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[string]
	{
		if ((v_0) >= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())).IntVal)) && ((v_0) <= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar())).IntVal)) {
			__t0 = Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Apply(Get_Data_Enum_fromCharCode(), gopurs_runtime.Int(v_0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Enum_defaultSucc__3855932694(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
defaultSucc__3855932694:
	for {
		if false {
			continue defaultSucc__3855932694
		}
		var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
		_ = toEnum_prime__unused_0
		var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
		_ = fromEnum_prime__unused_1
		var a_2 string = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Enum_charToEnum(gopurs_runtime.IntAdd(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(a_2)).IntVal, int64(1)))
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_defaultPred__3855932694(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
defaultPred__3855932694:
	for {
		if false {
			continue defaultPred__3855932694
		}
		var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
		_ = toEnum_prime__unused_0
		var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
		_ = fromEnum_prime__unused_1
		var a_2 string = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Enum_charToEnum(gopurs_runtime.IntSub(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(a_2)).IntVal, int64(1)))
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__1787937256(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__1787937256:
	for {
		if false {
			continue toEnum__1787937256
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Enum_charToEnum(__eta0_0)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__706605770(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__706605770:
	for {
		if false {
			continue toEnum__706605770
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__1693208633(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__1693208633:
	for {
		if false {
			continue toEnum__1693208633
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__6210423(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__6210423:
	for {
		if false {
			continue toEnum__6210423
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__2413733450(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__2413733450:
	for {
		if false {
			continue toEnum__2413733450
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__72020153(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__72020153:
	for {
		if false {
			continue toEnum__72020153
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_toEnum__3521034119(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
toEnum__3521034119:
	for {
		if false {
			continue toEnum__3521034119
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V4, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_succ(dict_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Enum_succ__77857599(__eta0_0_loop uint32) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
succ__77857599:
	for {
		if false {
			continue succ__77857599
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V2, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_succ__545117358(__eta0_0_loop *Constructor_Data_Date_Date) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
succ__545117358:
	for {
		if false {
			continue succ__545117358
		}
		var __eta0_0 *Constructor_Data_Date_Date = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate())).V2, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__eta0_0)})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_succ__3573123786(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
succ__3573123786:
	for {
		if false {
			continue succ__3573123786
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V2, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_succ__2729625849(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
succ__2729625849:
	for {
		if false {
			continue succ__2729625849
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_upFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
	_ = dictEnum_0
	// TAST (Let): succ1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope4)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope4)]))
	succ1_1_0 := Call_Data_Enum_succ(dictEnum_0)
	_ = succ1_1_0
	return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_3, gopurs_runtime.Apply(succ1_1_0, x_3)}))}
		}))
	})
}

func Call_Data_Enum_pred(dict_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Enum_pred__77857599(__eta0_0_loop uint32) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
pred__77857599:
	for {
		if false {
			continue pred__77857599
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_pred__3573123786(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
pred__3573123786:
	for {
		if false {
			continue pred__3573123786
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V1, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_pred__2729625849(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
pred__2729625849:
	for {
		if false {
			continue pred__2729625849
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V1, gopurs_runtime.Int(__eta0_0))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_fromEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Data_Enum_fromEnum__798184167(__eta0_0_loop uint32) int64 {
fromEnum__798184167:
	for {
		if false {
			continue fromEnum__798184167
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth())).V3, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil}).IntVal
	}
}

func Call_Data_Enum_fromEnum__1526214825(__eta0_0_loop string) int64 {
fromEnum__1526214825:
	for {
		if false {
			continue fromEnum__1526214825
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__1795542731(__eta0_0_loop int64) int64 {
fromEnum__1795542731:
	for {
		if false {
			continue fromEnum__1795542731
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())).V3, gopurs_runtime.Int(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__2307775192(__eta0_0_loop int64) int64 {
fromEnum__2307775192:
	for {
		if false {
			continue fromEnum__2307775192
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear())).V3, gopurs_runtime.Int(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__3185865430(__eta_norm_0_0_loop int64) int64 {
fromEnum__3185865430:
	for {
		if false {
			continue fromEnum__3185865430
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__862809814(__eta0_0_loop int64) int64 {
fromEnum__862809814:
	for {
		if false {
			continue fromEnum__862809814
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V3, gopurs_runtime.Int(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__3415298635(__eta0_0_loop int64) int64 {
fromEnum__3415298635:
	for {
		if false {
			continue fromEnum__3415298635
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V3, gopurs_runtime.Int(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__1091906392(__eta0_0_loop int64) int64 {
fromEnum__1091906392:
	for {
		if false {
			continue fromEnum__1091906392
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V3, gopurs_runtime.Int(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_fromEnum__4269580326(__eta0_0_loop int64) int64 {
fromEnum__4269580326:
	for {
		if false {
			continue fromEnum__4269580326
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V3, gopurs_runtime.Int(__eta0_0)).IntVal
	}
}

func Call_Data_Enum_toEnumWithDefaults(dictBoundedEnum_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBoundedEnum_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_0_loop
	_ = dictBoundedEnum_0
	// TAST (Let): bottom2_1_0 shape=App(Var) bindingType=(TypeVar a$scope21)
	bottom2_1_0 := Call_Data_Bounded_bottom(gopurs_runtime.Apply(dictBoundedEnum_0.V0, gopurs_runtime.Value{}))
	_ = bottom2_1_0
	return gopurs_runtime.Func3(func(low_2 gopurs_runtime.Value, high_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope21)])
		v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBoundedEnum_0.V4, gopurs_runtime.Int(x_4.IntVal)))
		_ = v_5_1
		var __t3 gopurs_runtime.Value
		{
			if v_5_1 != nil {
				__t3 = (v_5_1).V0
				goto end_branch_3
			} else {

			}
		}
		{
			if v_5_1 == nil {
				var __t2 gopurs_runtime.Value
				{
					if (x_4.IntVal) < (gopurs_runtime.Apply(dictBoundedEnum_0.V3, bottom2_1_0).IntVal) {
						__t2 = low_2
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = high_3
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
	})
}

func Call_Data_Enum_toEnumWithDefaults__3799330569(__eta0_0_loop string, __eta1_1_loop string, __eta2_2_loop int64) string {
toEnumWithDefaults__3799330569:
	for {
		if false {
			continue toEnumWithDefaults__3799330569
		}
		var __eta0_0 string = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 int64 = __eta2_2_loop
		_ = __eta2_2
		// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Char])
		v_3_0 := Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Enum_charToEnum(__eta2_2)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_3_0
		var __t2 string
		{
			if v_3_0 != nil {
				__t2 = (v_3_0).V0
				goto end_branch_2
			} else {

			}
		}
		{
			if v_3_0 == nil {
				var __t1 string
				{
					if (__eta2_2) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())).IntVal) {
						__t1 = __eta0_0
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = __eta1_1
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() string { panic("Failed pattern match") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Data_Enum_toEnumWithDefaults__4060784579(__eta_norm_1_unused_0_loop string, __eta_norm_0_unused_1_loop string, x_2_loop int64) string {
toEnumWithDefaults__4060784579:
	for {
		if false {
			continue toEnumWithDefaults__4060784579
		}
		var __eta_norm_1_unused_0 string = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 string = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 int64 = x_2_loop
		_ = x_2
		// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Char])
		v_3_0 := Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Enum_charToEnum(x_2)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_3_0
		var __t2 string
		{
			if v_3_0 != nil {
				__t2 = (v_3_0).V0
				goto end_branch_2
			} else {

			}
		}
		{
			if v_3_0 == nil {
				var __t1 string
				{
					if (x_2) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())).IntVal) {
						__t1 = gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar()).StrVal()
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()).StrVal()
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() string { panic("Failed pattern match") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Data_Enum_toEnumWithDefaults__92386535(__eta_norm_1_0_loop string, __eta_norm_0_1_loop string, x_2_loop int64) string {
toEnumWithDefaults__92386535:
	for {
		if false {
			continue toEnumWithDefaults__92386535
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 int64 = x_2_loop
		_ = x_2
		// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Char])
		v_3_0 := Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := Call_Data_Enum_charToEnum(x_2)
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
		_ = v_3_0
		var __t2 string
		{
			if v_3_0 != nil {
				__t2 = (v_3_0).V0
				goto end_branch_2
			} else {

			}
		}
		{
			if v_3_0 == nil {
				var __t1 string
				{
					if (x_2) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())).IntVal) {
						__t1 = __eta_norm_1_0
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = __eta_norm_0_1
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() string { panic("Failed pattern match") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Data_Enum_enumTuple(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
	_ = dictEnum_0
	// TAST (Let): ordTuple_1_0 shape=App(Var) bindingType=Any
	ordTuple_1_0 := Call_Data_Tuple_ordTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{}))
	_ = ordTuple_1_0
	return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): Bounded0_3_1 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar b$scope30)])
		Bounded0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
		_ = Bounded0_3_1
		// TAST (Let): bottom2_4_2 shape=Other bindingType=(TypeVar b$scope30)
		bottom2_4_2 := Bounded0_3_1.V1
		_ = bottom2_4_2
		// TAST (Let): Enum1_5_3 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar b$scope30)])
		Enum1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
		_ = Enum1_5_3
		// TAST (Let): top2_6_4 shape=Other bindingType=(TypeVar b$scope30)
		top2_6_4 := Bounded0_3_1.V2
		_ = top2_6_4
		// TAST (Let): ordTuple1_7_5 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])])
		ordTuple1_7_5 := Rebox_Data_Enum_4177771502_1535415139(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordTuple_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))))
		_ = ordTuple1_7_5
		return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3481421219_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1535415139_4177771502(ordTuple1_7_5))}
		}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_9_6 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope29)])
			__local_var_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
			_ = __local_var_9_6
			var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
			{
				if __local_var_9_6 != nil {
					__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_9_6).V0, top2_6_4}))}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}())
					goto end_branch_8
				} else {

				}
			}
			{
				__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
			}
		end_branch_8:
			// TAST (Let): __local_var_10_7 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])])
			var __local_var_10_7 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Data_Enum_3094389156_4010058633(__t8)
			// TAST (Let): __local_var_11_9 shape=App(Var) bindingType=(Func [(TypeVar b$scope30)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])]))
			__local_var_11_9 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
			_ = __local_var_11_9
			// TAST (Let): __local_var_12_10 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope30)])
			__local_var_12_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_5_3.V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
			_ = __local_var_12_10
			var __t11 gopurs_runtime.Value
			{
				if __local_var_12_10 == nil {
					__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(__local_var_10_7))}
					goto end_branch_11
				} else {

				}
			}
			{
				if __local_var_12_10 != nil {
					__t11 = gopurs_runtime.Apply(__local_var_11_9, (__local_var_12_10).V0)
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_11:
			return __t11
		}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_9_12 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope29)])
			__local_var_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
			_ = __local_var_9_12
			var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
			{
				if __local_var_9_12 != nil {
					__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_9_12).V0, bottom2_4_2}))}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}())
					goto end_branch_14
				} else {

				}
			}
			{
				__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}())
			}
		end_branch_14:
			// TAST (Let): __local_var_10_13 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])])
			var __local_var_10_13 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Data_Enum_3094389156_4010058633(__t14)
			// TAST (Let): __local_var_11_15 shape=App(Var) bindingType=(Func [(TypeVar b$scope30)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])]))
			__local_var_11_15 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
			_ = __local_var_11_15
			// TAST (Let): __local_var_12_16 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope30)])
			__local_var_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_5_3.V2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
			_ = __local_var_12_16
			var __t17 gopurs_runtime.Value
			{
				if __local_var_12_16 == nil {
					__t17 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(__local_var_10_13))}
					goto end_branch_17
				} else {

				}
			}
			{
				if __local_var_12_16 != nil {
					__t17 = gopurs_runtime.Apply(__local_var_11_15, (__local_var_12_16).V0)
					goto end_branch_17
				} else {

				}
			}
			{
				__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_17:
			return __t17
		})})))}
	})
}

func Call_Data_Enum_enumMaybe(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
	_ = dictBoundedEnum_0
	// TAST (Let): bottom2_1_0 shape=Other bindingType=(TypeVar a$scope52)
	bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom")
	_ = bottom2_1_0
	// TAST (Let): Enum1_2_1 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a$scope52)])
	Enum1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
	_ = Enum1_2_1
	// TAST (Let): ordMaybe_3_2 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope52)])])
	ordMaybe_3_2 := Rebox_Data_Enum_4177771502_2155612431(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))))
	_ = ordMaybe_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2522609743_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2155612431_4177771502(ordMaybe_3_2))}
	}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				goto end_branch_3
			} else {

			}
		}
		{
			if v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil {
				__t3 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Apply(Enum1_2_1.V1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0), true}
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() struct {
				V0 gopurs_runtime.Value
				V1 bool
			} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
				panic("Failed pattern match")
			}()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
			}; return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: gopurs_runtime.Value{}, V1: false} }()
		}
	end_branch_3:
		return func() gopurs_runtime.Value {
			_v := __t3
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t6 struct {
			V0 gopurs_runtime.Value
			V1 bool
		}
		{
			if v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil {
				__t6 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, bottom2_1_0}))}, true}
				goto end_branch_6
			} else {

			}
		}
		{
			if v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil {
				// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope52)])
				__local_var_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_2_1.V2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0))
				_ = __local_var_5_4
				var __t5 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if __local_var_5_4 != nil {
						__t5 = struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (__local_var_5_4).V0}))}, true}
						goto end_branch_5
					} else {

					}
				}
				{
					__t5 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
				}
			end_branch_5:
				__t6 = __t5
				goto end_branch_6
			} else {

			}
		}
		{
			__t6 = func() struct {
				V0 gopurs_runtime.Value
				V1 bool
			} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
				panic("Failed pattern match")
			}()))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
			}; return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: gopurs_runtime.Value{}, V1: false} }()
		}
	end_branch_6:
		return func() gopurs_runtime.Value {
			_v := __t6
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()
	})})))}
}

func Call_Data_Enum_enumFromTo(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
	_ = dictEnum_0
	// TAST (Let): Ord0_1_0 shape=App(Other) bindingType=Any
	Ord0_1_0 := gopurs_runtime.Apply(dictEnum_0.V0, gopurs_runtime.Value{})
	_ = Ord0_1_0
	// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(TypeVar a$scope66)])
	__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V0, gopurs_runtime.Value{}))
	_ = __local_var_2_1
	// TAST (Let): succ1_3_2 shape=App(Var) bindingType=(Func [(TypeVar a$scope66)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope66)]))
	succ1_3_2 := Call_Data_Enum_succ(dictEnum_0)
	_ = succ1_3_2
	// TAST (Let): pred1_4_3 shape=App(Var) bindingType=(Func [(TypeVar a$scope66)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope66)]))
	pred1_4_3 := Call_Data_Enum_pred(dictEnum_0)
	_ = pred1_4_3
	return gopurs_runtime.Func3(func(dictUnfoldable1_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t14 gopurs_runtime.Value
		{
			if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}), "eq"), v_6, v1_7).IntVal) != (0) {
				__t14 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_5, "unfoldr1"), gopurs_runtime.Func(func(i_8 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t8 struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}
					{
						if (i_8.IntVal) <= (int64(0)) {
							__t8 = struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
							}{v_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{v_6, func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Int(gopurs_runtime.IntSub(i_8.IntVal, int64(1))), true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()}
					}
				end_branch_8:
					return func() gopurs_runtime.Value {
						_v := __t8
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				}), gopurs_runtime.Int(int64(0)))
				goto end_branch_14
			} else {

			}
		}
		{
			var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_2_1.V1, v_6, v1_7)
			_ = __t_tag_9
			if uint32(__t_tag_9.IntVal) == 1527465420 {
				__t14 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_5, "unfoldr1"), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_9_10 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope66)])
					__local_var_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(succ1_3_2, a_8))
					_ = __local_var_9_10
					var __t13 struct {
						V0 gopurs_runtime.Value
						V1 bool
					}
					{
						if __local_var_9_10 != nil {
							var __t12 struct {
								V0 gopurs_runtime.Value
								V1 bool
							}
							{
								var __t_tag_11 struct {
									V0 gopurs_runtime.Value
									V1 bool
								} = Call_Control_Alternative_guard__3382145287(Call_Data_Ord_lessThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Ord0_1_0), (__local_var_9_10).V0, v1_7))
								_ = __t_tag_11
								if __t_tag_11.V1 {
									__t12 = struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{(__local_var_9_10).V0, true}
									goto end_branch_12
								} else {

								}
							}
							{
								__t12 = struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
							}
						end_branch_12:
							__t13 = __t12
							goto end_branch_13
						} else {

						}
					}
					{
						if __local_var_9_10 == nil {
							__t13 = struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							goto end_branch_13
						} else {

						}
					}
					{
						__t13 = func() struct {
							V0 gopurs_runtime.Value
							V1 bool
						} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
							return struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
						}; return struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{V0: gopurs_runtime.Value{}, V1: false} }()
					}
				end_branch_13:
					return func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{a_8, func() gopurs_runtime.Value {
							_v := __t13
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()
				}), v_6)
				goto end_branch_14
			} else {

			}
		}
		{
			__t14 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_5, "unfoldr1"), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_9_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope66)])
				__local_var_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pred1_4_3, a_8))
				_ = __local_var_9_4
				var __t7 struct {
					V0 gopurs_runtime.Value
					V1 bool
				}
				{
					if __local_var_9_4 != nil {
						var __t6 struct {
							V0 gopurs_runtime.Value
							V1 bool
						}
						{
							var __t_tag_5 struct {
								V0 gopurs_runtime.Value
								V1 bool
							} = Call_Control_Alternative_guard__3382145287(Call_Data_Ord_greaterThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Ord0_1_0), (__local_var_9_4).V0, v1_7))
							_ = __t_tag_5
							if __t_tag_5.V1 {
								__t6 = struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{(__local_var_9_4).V0, true}
								goto end_branch_6
							} else {

							}
						}
						{
							__t6 = struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
						}
					end_branch_6:
						__t7 = __t6
						goto end_branch_7
					} else {

					}
				}
				{
					if __local_var_9_4 == nil {
						__t7 = struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						goto end_branch_7
					} else {

					}
				}
				{
					__t7 = func() struct {
						V0 gopurs_runtime.Value
						V1 bool
					} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
						return struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
					}; return struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{V0: gopurs_runtime.Value{}, V1: false} }()
				}
			end_branch_7:
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{a_8, func() gopurs_runtime.Value {
						_v := __t7
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			}), v_6)
		}
	end_branch_14:
		return __t14
	})
}

func Call_Data_Enum_enumFromThenTo(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], dictBoundedEnum_2_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
	_ = dictUnfoldable_0
	var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
	_ = dictFunctor_1
	var dictBoundedEnum_2 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_2_loop
	_ = dictBoundedEnum_2
	// TAST (Let): toEnum1_3_0 shape=App(Var) bindingType=(Func [Int] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope73)]))
	toEnum1_3_0 := Call_Data_Enum_toEnum(dictBoundedEnum_2)
	_ = toEnum1_3_0
	return gopurs_runtime.Func3(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, c_6 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): a_prime__7_1 shape=App(Other) bindingType=Int
		a_prime__7_1 := gopurs_runtime.Apply(dictBoundedEnum_2.V3, a_4).IntVal
		_ = a_prime__7_1
		// TAST (Let): __local_var_8_2 shape=Other bindingType=Int
		__local_var_8_2 := gopurs_runtime.IntSub(gopurs_runtime.Apply(dictBoundedEnum_2.V3, b_5).IntVal, a_prime__7_1)
		_ = __local_var_8_2
		// TAST (Let): __local_var_9_3 shape=App(Other) bindingType=Int
		__local_var_9_3 := gopurs_runtime.Apply(dictBoundedEnum_2.V3, c_6).IntVal
		_ = __local_var_9_3
		return gopurs_runtime.Apply2(dictFunctor_1.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), toEnum1_3_0, gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})), gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(e_10 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if (e_10.IntVal) <= (__local_var_9_3) {
					__t4 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
							}{e_10, gopurs_runtime.Int(gopurs_runtime.IntAdd(e_10.IntVal, __local_var_8_2))}
							return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
						}(), true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
			}
		end_branch_4:
			return __t4
		}), gopurs_runtime.Int(a_prime__7_1)))
	})
}

func Call_Data_Enum_enumEither(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
	_ = dictBoundedEnum_0
	// TAST (Let): Enum1_1_0 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a$scope81)])
	Enum1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
	_ = Enum1_1_0
	// TAST (Let): top2_2_1 shape=Other bindingType=(TypeVar a$scope81)
	top2_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "top")
	_ = top2_2_1
	// TAST (Let): ordEither_3_2 shape=App(Var) bindingType=Any
	ordEither_3_2 := Call_Data_Either_ordEither(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))
	_ = ordEither_3_2
	return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): bottom2_5_3 shape=Other bindingType=(TypeVar b$scope82)
		bottom2_5_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Bounded0"), gopurs_runtime.Value{}), "bottom")
		_ = bottom2_5_3
		// TAST (Let): Enum11_6_4 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar b$scope82)])
		Enum11_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}))
		_ = Enum11_6_4
		// TAST (Let): ordEither1_7_5 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])])
		ordEither1_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordEither_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{})))
		_ = ordEither1_7_5
		return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Enum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordEither1_7_5)}
		}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t12 gopurs_runtime.Value
			{
				if v_8.Type == 9 && v_8.IntVal == 3711209382 {
					// TAST (Let): __local_var_9_6 shape=App(Var) bindingType=(Func [(TypeVar a$scope81)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
					__local_var_9_6 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Left())
					_ = __local_var_9_6
					// TAST (Let): __local_var_10_7 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope81)])
					__local_var_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_1_0.V1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
					_ = __local_var_10_7
					var __t8 gopurs_runtime.Value
					{
						if __local_var_10_7 == nil {
							__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
							goto end_branch_8
						} else {

						}
					}
					{
						if __local_var_10_7 != nil {
							__t8 = gopurs_runtime.Apply(__local_var_9_6, (__local_var_10_7).V0)
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_8:
					__t12 = __t8
					goto end_branch_12
				} else {

				}
			}
			{
				if v_8.Type == 9 && v_8.IntVal == 2465973597 {
					// TAST (Let): __local_var_9_9 shape=App(Var) bindingType=(Func [(TypeVar b$scope82)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
					__local_var_9_9 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Right())
					_ = __local_var_9_9
					// TAST (Let): __local_var_10_10 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope82)])
					__local_var_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum11_6_4.V1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
					_ = __local_var_10_10
					var __t11 gopurs_runtime.Value
					{
						if __local_var_10_10 == nil {
							__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, top2_2_1}))}}))}
							goto end_branch_11
						} else {

						}
					}
					{
						if __local_var_10_10 != nil {
							__t11 = gopurs_runtime.Apply(__local_var_9_9, (__local_var_10_10).V0)
							goto end_branch_11
						} else {

						}
					}
					{
						__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_11:
					__t12 = __t11
					goto end_branch_12
				} else {

				}
			}
			{
				__t12 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_12:
			return __t12
		}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t19 gopurs_runtime.Value
			{
				if v_8.Type == 9 && v_8.IntVal == 3711209382 {
					// TAST (Let): __local_var_9_13 shape=App(Var) bindingType=(Func [(TypeVar a$scope81)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
					__local_var_9_13 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Left())
					_ = __local_var_9_13
					// TAST (Let): __local_var_10_14 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope81)])
					__local_var_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_1_0.V2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
					_ = __local_var_10_14
					var __t15 gopurs_runtime.Value
					{
						if __local_var_10_14 == nil {
							__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, bottom2_5_3}))}}))}
							goto end_branch_15
						} else {

						}
					}
					{
						if __local_var_10_14 != nil {
							__t15 = gopurs_runtime.Apply(__local_var_9_13, (__local_var_10_14).V0)
							goto end_branch_15
						} else {

						}
					}
					{
						__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_15:
					__t19 = __t15
					goto end_branch_19
				} else {

				}
			}
			{
				if v_8.Type == 9 && v_8.IntVal == 2465973597 {
					// TAST (Let): __local_var_9_16 shape=App(Var) bindingType=(Func [(TypeVar b$scope82)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
					__local_var_9_16 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Right())
					_ = __local_var_9_16
					// TAST (Let): __local_var_10_17 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope82)])
					__local_var_10_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum11_6_4.V2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
					_ = __local_var_10_17
					var __t18 gopurs_runtime.Value
					{
						if __local_var_10_17 == nil {
							__t18 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
							goto end_branch_18
						} else {

						}
					}
					{
						if __local_var_10_17 != nil {
							__t18 = gopurs_runtime.Apply(__local_var_9_16, (__local_var_10_17).V0)
							goto end_branch_18
						} else {

						}
					}
					{
						__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_18:
					__t19 = __t18
					goto end_branch_19
				} else {

				}
			}
			{
				__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_19:
			return __t19
		})}))}
	})
}

func Call_Data_Enum_downFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
	_ = dictEnum_0
	// TAST (Let): pred1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope115)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope115)]))
	pred1_1_0 := Call_Data_Enum_pred(dictEnum_0)
	_ = pred1_1_0
	return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_3, gopurs_runtime.Apply(pred1_1_0, x_3)}))}
		}))
	})
}

func Call_Data_Enum_diag(a_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{a_0, a_0}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Data_Enum_downFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
	_ = dictEnum_0
	// TAST (Let): pred1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope125)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope125)]))
	pred1_1_0 := Call_Data_Enum_pred(dictEnum_0)
	_ = pred1_1_0
	return gopurs_runtime.Func(func(dictUnfoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_2, "unfoldr"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{func() gopurs_runtime.Value {
						_v := Call_Data_Enum_diag((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}(), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), pred1_1_0))
	})
}

func Call_Data_Enum_upFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
	_ = dictEnum_0
	// TAST (Let): succ1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope130)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope130)]))
	succ1_1_0 := Call_Data_Enum_succ(dictEnum_0)
	_ = succ1_1_0
	return gopurs_runtime.Func(func(dictUnfoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_2, "unfoldr"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil {
					__t1 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{func() gopurs_runtime.Value {
						_v := Call_Data_Enum_diag((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}(), true}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
			}
		end_branch_1:
			return func() gopurs_runtime.Value {
				_v := __t1
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		}), succ1_1_0))
	})
}

func Call_Data_Enum_defaultToEnum(dictBounded_0_loop *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBounded_0 *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] = dictBounded_0_loop
	_ = dictBounded_0
	// TAST (Let): bottom2_1_0 shape=Other bindingType=(TypeVar a$scope134)
	bottom2_1_0 := dictBounded_0.V1
	_ = bottom2_1_0
	return gopurs_runtime.Func2(func(dictEnum_2 gopurs_runtime.Value, i_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
		var Call_local_Data_Enum_go__go_4_1_0 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
		_ = Call_local_Data_Enum_go__go_4_1_0
		var go__go_4_1_0 gopurs_runtime.Value
		_ = go__go_4_1_0
		Call_local_Data_Enum_go__go_4_1_0 = func(i_5_loop int64, x_6_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
		go__go_4_1_0:
			for {
				if false {
					continue go__go_4_1_0
				}
				var i_5 int64 = i_5_loop
				_ = i_5
				var x_6 gopurs_runtime.Value = x_6_loop
				_ = x_6
				var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				{
					if (i_5) == (int64(0)) {
						__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{x_6, true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}())
						goto end_branch_4
					} else {

					}
				}
				{
					// TAST (Let): v_7_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope134)])
					v_7_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_6))
					_ = v_7_2
					var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if v_7_2 != nil {
							i_5_loop = gopurs_runtime.IntSub(i_5, int64(1))
							x_6_loop = (v_7_2).V0
							continue go__go_4_1_0
							__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
							goto end_branch_3
						} else {

						}
					}
					{
						if v_7_2 == nil {
							__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
							goto end_branch_3
						} else {

						}
					}
					{
						__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
					}
				end_branch_3:
					__t4 = __t3
				}
			end_branch_4:
				return __t4
			}
		}
		go__go_4_1_0 = gopurs_runtime.Func(func(i_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(x_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Enum_go__go_4_1_0(i_5_loop_val.IntVal, x_6_loop_val))}
			})
		})
		var __t5 gopurs_runtime.Value
		{
			if (i_prime__3.IntVal) < (int64(0)) {
				__t5 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()
				goto end_branch_5
			} else {

			}
		}
		{
			__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Enum_go__go_4_1_0(i_prime__3.IntVal, bottom2_1_0))}
		}
	end_branch_5:
		return __t5
	})
}

func Call_Data_Enum_defaultSucc(toEnum_prime__0_loop gopurs_runtime.Value, fromEnum_prime__1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var toEnum_prime__0 gopurs_runtime.Value = toEnum_prime__0_loop
	_ = toEnum_prime__0
	var fromEnum_prime__1 gopurs_runtime.Value = fromEnum_prime__1_loop
	_ = fromEnum_prime__1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime__0, gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply(fromEnum_prime__1, a_2).IntVal, int64(1))))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Enum_defaultSucc__2406132389(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
defaultSucc__2406132389:
	for {
		if false {
			continue defaultSucc__2406132389
		}
		var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
		_ = toEnum_prime__unused_0
		var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
		_ = fromEnum_prime__unused_1
		var a_2 int64 = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V4, gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V3, gopurs_runtime.Int(a_2)).IntVal, int64(1))))))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_defaultPred(toEnum_prime__0_loop gopurs_runtime.Value, fromEnum_prime__1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var toEnum_prime__0 gopurs_runtime.Value = toEnum_prime__0_loop
	_ = toEnum_prime__0
	var fromEnum_prime__1 gopurs_runtime.Value = fromEnum_prime__1_loop
	_ = fromEnum_prime__1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime__0, gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(fromEnum_prime__1, a_2).IntVal, int64(1))))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Enum_defaultPred__2406132389(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
defaultPred__2406132389:
	for {
		if false {
			continue defaultPred__2406132389
		}
		var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
		_ = toEnum_prime__unused_0
		var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
		_ = fromEnum_prime__unused_1
		var a_2 int64 = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V4, gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V3, gopurs_runtime.Int(a_2)).IntVal, int64(1))))))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Enum_defaultFromEnum(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
	_ = dictEnum_0
	var Call_local_Data_Enum_go__118764402_1_0_1 func(int64, gopurs_runtime.Value) int64
	_ = Call_local_Data_Enum_go__118764402_1_0_1
	var go__118764402_1_0_1 gopurs_runtime.Value
	_ = go__118764402_1_0_1
	var Call_local_Data_Enum_go__go_1_1_2 func(int64, gopurs_runtime.Value) int64
	_ = Call_local_Data_Enum_go__go_1_1_2
	var go__go_1_1_2 gopurs_runtime.Value
	_ = go__go_1_1_2
	Call_local_Data_Enum_go__118764402_1_0_1 = func(i_2_loop int64, x_3_loop gopurs_runtime.Value) int64 {
	go__118764402_1_0_1:
		for {
			if false {
				continue go__118764402_1_0_1
			}
			var i_2 int64 = i_2_loop
			_ = i_2
			var x_3 gopurs_runtime.Value = x_3_loop
			_ = x_3
			// TAST (Let): v_4_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope143)])
			v_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V1, x_3))
			_ = v_4_2
			var __t3 int64
			{
				if v_4_2 != nil {
					i_2_loop = gopurs_runtime.IntAdd(i_2, int64(1))
					x_3_loop = (v_4_2).V0
					continue go__118764402_1_0_1
					__t3 = func() int64 { panic("unreachable") }()
					goto end_branch_3
				} else {

				}
			}
			{
				if v_4_2 == nil {
					__t3 = i_2
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() int64 { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}
	}
	go__118764402_1_0_1 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_1_0_1(i_2_loop_val.IntVal, x_3_loop_val))
		})
	})
	Call_local_Data_Enum_go__go_1_1_2 = func(i_2_loop int64, x_3_loop gopurs_runtime.Value) int64 {
	go__go_1_1_2:
		for {
			if false {
				continue go__go_1_1_2
			}
			var i_2 int64 = i_2_loop
			_ = i_2
			var x_3 gopurs_runtime.Value = x_3_loop
			_ = x_3
			// TAST (Let): v_4_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope143)])
			v_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V1, x_3))
			_ = v_4_4
			var __t5 int64
			{
				if v_4_4 != nil {
					__t5 = Call_local_Data_Enum_go__118764402_1_0_1(gopurs_runtime.IntAdd(i_2, int64(1)), (v_4_4).V0)
					goto end_branch_5
				} else {

				}
			}
			{
				if v_4_4 == nil {
					__t5 = i_2
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() int64 { panic("Failed pattern match") }()
			}
		end_branch_5:
			return __t5
		}
	}
	go__go_1_1_2 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Data_Enum_go__go_1_1_2(i_2_loop_val.IntVal, x_3_loop_val))
		})
	})
	return gopurs_runtime.Apply(go__118764402_1_0_1, gopurs_runtime.Int(int64(0)))
}

func Call_Data_Enum_defaultCardinality(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	// TAST (Let): bottom2_1_0 shape=App(Var) bindingType=(TypeVar a$scope145)
	bottom2_1_0 := Call_Data_Bounded_bottom(dictBounded_0)
	_ = bottom2_1_0
	return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var Call_local_Data_Enum_go__go_3_1_3 func(int64, gopurs_runtime.Value) int64
		_ = Call_local_Data_Enum_go__go_3_1_3
		var go__go_3_1_3 gopurs_runtime.Value
		_ = go__go_3_1_3
		Call_local_Data_Enum_go__go_3_1_3 = func(i_4_loop int64, x_5_loop gopurs_runtime.Value) int64 {
		go__go_3_1_3:
			for {
				if false {
					continue go__go_3_1_3
				}
				var i_4 int64 = i_4_loop
				_ = i_4
				var x_5 gopurs_runtime.Value = x_5_loop
				_ = x_5
				// TAST (Let): v_6_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope145)])
				v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5))
				_ = v_6_2
				var __t3 int64
				{
					if v_6_2 != nil {
						i_4_loop = gopurs_runtime.IntAdd(i_4, int64(1))
						x_5_loop = (v_6_2).V0
						continue go__go_3_1_3
						__t3 = func() int64 { panic("unreachable") }()
						goto end_branch_3
					} else {

					}
				}
				{
					if v_6_2 == nil {
						__t3 = i_4
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() int64 { panic("Failed pattern match") }()
				}
			end_branch_3:
				return __t3
			}
		}
		go__go_3_1_3 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Data_Enum_go__go_3_1_3(i_4_loop_val.IntVal, x_5_loop_val))
			})
		})
		return gopurs_runtime.Int(Call_local_Data_Enum_go__go_3_1_3(int64(1), bottom2_1_0))
	})
}

func Call_Data_Enum_cardinality(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Int(gopurs_runtime.RecordGet(dict_0, "cardinality").IntVal)
}

func Rebox_Data_Enum_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Enum_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_123048125_4021906832(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_1535415139_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2094947566_2420955921(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[string]{}
	out.V0 = in.V0
	out.V1 = in.V1.StrVal()
	out.V2 = in.V2.StrVal()
	return out
}

func Rebox_Data_Enum_2094947566_3728870730(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[bool]{}
	out.V0 = in.V0
	out.V1 = (in.V1.IntVal) != (0)
	out.V2 = (in.V2.IntVal) != (0)
	return out
}

func Rebox_Data_Enum_2094947566_832288803(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[uint32]{}
	out.V0 = in.V0
	out.V1 = uint32(in.V1.IntVal)
	out.V2 = uint32(in.V2.IntVal)
	return out
}

func Rebox_Data_Enum_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_215731685_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2359585123_556578094(in *Constructor_Data_Enum_Enum[uint32]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2420955921_2094947566(in *Constructor_Data_Bounded_Bounded[string]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	out.V2 = gopurs_runtime.Str(in.V2)
	return out
}

func Rebox_Data_Enum_2522609743_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2621187955_385277032(in *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, int64]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3083804185_123048125(in *Constructor_Data_Enum_BoundedEnum[bool]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Enum_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Enum_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Enum_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3409577041_556578094(in *Constructor_Data_Enum_Enum[string]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_342165130_556578094(in *Constructor_Data_Enum_Enum[bool]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3481421219_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3494563625_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3569500834_123048125(in *Constructor_Data_Enum_BoundedEnum[string]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3728870730_2094947566(in *Constructor_Data_Bounded_Bounded[bool]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Bool(in.V1)
	out.V2 = gopurs_runtime.Bool(in.V2)
	return out
}

func Rebox_Data_Enum_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4060049525_556578094(in *Constructor_Data_Enum_Enum[int64]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4177771502_1535415139(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4177771502_2155612431(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4177771502_219188042(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[bool])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_556578094_2359585123(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_556578094_3409577041(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[string])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_556578094_342165130(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[bool])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_556578094_3494563625(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_556578094_4060049525(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Enum_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Enum_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Rebox_Data_Enum_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}

func Get_Data_Enum_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_FromCharCode
}

func Get_Data_Enum_toCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_ToCharCode
}
