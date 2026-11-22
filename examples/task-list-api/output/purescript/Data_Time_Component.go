package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Component_go__append gopurs_runtime.Value
var once_Data_Time_Component_go__append sync.Once

func Get_Data_Time_Component_go__append() gopurs_runtime.Value {
	once_Data_Time_Component_go__append.Do(func() {
		cache_Data_Time_Component_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Time_Component_go__append
}

var cache_Data_Time_Component_show gopurs_runtime.Value
var once_Data_Time_Component_show sync.Once

func Get_Data_Time_Component_show() gopurs_runtime.Value {
	once_Data_Time_Component_show.Do(func() {
		cache_Data_Time_Component_show = Get_Data_Show_show__1986550628()
	})
	return cache_Data_Time_Component_show
}

var cache_Data_Time_Component_conj gopurs_runtime.Value
var once_Data_Time_Component_conj sync.Once

func Get_Data_Time_Component_conj() gopurs_runtime.Value {
	once_Data_Time_Component_conj.Do(func() {
		cache_Data_Time_Component_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Time_Component_conj
}

var cache_Data_Time_Component_greaterThanOrEq gopurs_runtime.Value
var once_Data_Time_Component_greaterThanOrEq sync.Once

func Get_Data_Time_Component_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_Time_Component_greaterThanOrEq.Do(func() {
		cache_Data_Time_Component_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__3348776441()
	})
	return cache_Data_Time_Component_greaterThanOrEq
}

var cache_Data_Time_Component_lessThanOrEq gopurs_runtime.Value
var once_Data_Time_Component_lessThanOrEq sync.Once

func Get_Data_Time_Component_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Time_Component_lessThanOrEq.Do(func() {
		cache_Data_Time_Component_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Time_Component_lessThanOrEq
}

var cache_Data_Time_Component_add gopurs_runtime.Value
var once_Data_Time_Component_add sync.Once

func Get_Data_Time_Component_add() gopurs_runtime.Value {
	once_Data_Time_Component_add.Do(func() {
		cache_Data_Time_Component_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Time_Component_add
}

var cache_Data_Time_Component_sub gopurs_runtime.Value
var once_Data_Time_Component_sub sync.Once

func Get_Data_Time_Component_sub() gopurs_runtime.Value {
	once_Data_Time_Component_sub.Do(func() {
		cache_Data_Time_Component_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Time_Component_sub
}

var cache_Data_Time_Component_Second gopurs_runtime.Value
var once_Data_Time_Component_Second sync.Once

func Get_Data_Time_Component_Second() gopurs_runtime.Value {
	once_Data_Time_Component_Second.Do(func() {
		cache_Data_Time_Component_Second = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Time_Component_Second(x_0_box.IntVal))
		})
	})
	return cache_Data_Time_Component_Second
}

var cache_Data_Time_Component_Minute gopurs_runtime.Value
var once_Data_Time_Component_Minute sync.Once

func Get_Data_Time_Component_Minute() gopurs_runtime.Value {
	once_Data_Time_Component_Minute.Do(func() {
		cache_Data_Time_Component_Minute = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Time_Component_Minute(x_0_box.IntVal))
		})
	})
	return cache_Data_Time_Component_Minute
}

var cache_Data_Time_Component_Millisecond gopurs_runtime.Value
var once_Data_Time_Component_Millisecond sync.Once

func Get_Data_Time_Component_Millisecond() gopurs_runtime.Value {
	once_Data_Time_Component_Millisecond.Do(func() {
		cache_Data_Time_Component_Millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Time_Component_Millisecond(x_0_box.IntVal))
		})
	})
	return cache_Data_Time_Component_Millisecond
}

var cache_Data_Time_Component_Hour gopurs_runtime.Value
var once_Data_Time_Component_Hour sync.Once

func Get_Data_Time_Component_Hour() gopurs_runtime.Value {
	once_Data_Time_Component_Hour.Do(func() {
		cache_Data_Time_Component_Hour = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Time_Component_Hour(x_0_box.IntVal))
		})
	})
	return cache_Data_Time_Component_Hour
}

var cache_Data_Time_Component_showSecond gopurs_runtime.Value
var once_Data_Time_Component_showSecond sync.Once

func Get_Data_Time_Component_showSecond() gopurs_runtime.Value {
	once_Data_Time_Component_showSecond.Do(func() {
		cache_Data_Time_Component_showSecond = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Second ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Time_Component_showSecond
}

var cache_Data_Time_Component_showMinute gopurs_runtime.Value
var once_Data_Time_Component_showMinute sync.Once

func Get_Data_Time_Component_showMinute() gopurs_runtime.Value {
	once_Data_Time_Component_showMinute.Do(func() {
		cache_Data_Time_Component_showMinute = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Minute ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Time_Component_showMinute
}

var cache_Data_Time_Component_showMillisecond gopurs_runtime.Value
var once_Data_Time_Component_showMillisecond sync.Once

func Get_Data_Time_Component_showMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_showMillisecond.Do(func() {
		cache_Data_Time_Component_showMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Millisecond ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Time_Component_showMillisecond
}

var cache_Data_Time_Component_showHour gopurs_runtime.Value
var once_Data_Time_Component_showHour sync.Once

func Get_Data_Time_Component_showHour() gopurs_runtime.Value {
	once_Data_Time_Component_showHour.Do(func() {
		cache_Data_Time_Component_showHour = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Hour ", Data_Show_ShowIntImpl(v_0.IntVal)), ")"))
		})})))}
	})
	return cache_Data_Time_Component_showHour
}

var cache_Data_Time_Component_ordSecond gopurs_runtime.Value
var once_Data_Time_Component_ordSecond sync.Once

func Get_Data_Time_Component_ordSecond() gopurs_runtime.Value {
	once_Data_Time_Component_ordSecond.Do(func() {
		cache_Data_Time_Component_ordSecond = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Time_Component_ordSecond
}

var cache_Data_Time_Component_ordMinute gopurs_runtime.Value
var once_Data_Time_Component_ordMinute sync.Once

func Get_Data_Time_Component_ordMinute() gopurs_runtime.Value {
	once_Data_Time_Component_ordMinute.Do(func() {
		cache_Data_Time_Component_ordMinute = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Time_Component_ordMinute
}

var cache_Data_Time_Component_ordMillisecond gopurs_runtime.Value
var once_Data_Time_Component_ordMillisecond sync.Once

func Get_Data_Time_Component_ordMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_ordMillisecond.Do(func() {
		cache_Data_Time_Component_ordMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Time_Component_ordMillisecond
}

var cache_Data_Time_Component_ordHour gopurs_runtime.Value
var once_Data_Time_Component_ordHour sync.Once

func Get_Data_Time_Component_ordHour() gopurs_runtime.Value {
	once_Data_Time_Component_ordHour.Do(func() {
		cache_Data_Time_Component_ordHour = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Time_Component_ordHour
}

var cache_Data_Time_Component_eqSecond gopurs_runtime.Value
var once_Data_Time_Component_eqSecond sync.Once

func Get_Data_Time_Component_eqSecond() gopurs_runtime.Value {
	once_Data_Time_Component_eqSecond.Do(func() {
		cache_Data_Time_Component_eqSecond = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1053099733_3790796878(Rebox_Data_Time_Component_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Time_Component_eqSecond
}

var cache_Data_Time_Component_eqMinute gopurs_runtime.Value
var once_Data_Time_Component_eqMinute sync.Once

func Get_Data_Time_Component_eqMinute() gopurs_runtime.Value {
	once_Data_Time_Component_eqMinute.Do(func() {
		cache_Data_Time_Component_eqMinute = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1053099733_3790796878(Rebox_Data_Time_Component_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Time_Component_eqMinute
}

var cache_Data_Time_Component_eqMillisecond gopurs_runtime.Value
var once_Data_Time_Component_eqMillisecond sync.Once

func Get_Data_Time_Component_eqMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_eqMillisecond.Do(func() {
		cache_Data_Time_Component_eqMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1053099733_3790796878(Rebox_Data_Time_Component_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Time_Component_eqMillisecond
}

var cache_Data_Time_Component_eqHour gopurs_runtime.Value
var once_Data_Time_Component_eqHour sync.Once

func Get_Data_Time_Component_eqHour() gopurs_runtime.Value {
	once_Data_Time_Component_eqHour.Do(func() {
		cache_Data_Time_Component_eqHour = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1053099733_3790796878(Rebox_Data_Time_Component_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Time_Component_eqHour
}

var cache_Data_Time_Component_boundedSecond gopurs_runtime.Value
var once_Data_Time_Component_boundedSecond sync.Once

func Get_Data_Time_Component_boundedSecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedSecond.Do(func() {
		cache_Data_Time_Component_boundedSecond = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), int64(0), int64(59)})))}
	})
	return cache_Data_Time_Component_boundedSecond
}

var cache_Data_Time_Component_boundedMinute gopurs_runtime.Value
var once_Data_Time_Component_boundedMinute sync.Once

func Get_Data_Time_Component_boundedMinute() gopurs_runtime.Value {
	once_Data_Time_Component_boundedMinute.Do(func() {
		cache_Data_Time_Component_boundedMinute = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), int64(0), int64(59)})))}
	})
	return cache_Data_Time_Component_boundedMinute
}

var cache_Data_Time_Component_boundedMillisecond gopurs_runtime.Value
var once_Data_Time_Component_boundedMillisecond sync.Once

func Get_Data_Time_Component_boundedMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedMillisecond.Do(func() {
		cache_Data_Time_Component_boundedMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), int64(0), int64(999)})))}
	})
	return cache_Data_Time_Component_boundedMillisecond
}

var cache_Data_Time_Component_boundedHour gopurs_runtime.Value
var once_Data_Time_Component_boundedHour sync.Once

func Get_Data_Time_Component_boundedHour() gopurs_runtime.Value {
	once_Data_Time_Component_boundedHour.Do(func() {
		cache_Data_Time_Component_boundedHour = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), int64(0), int64(23)})))}
	})
	return cache_Data_Time_Component_boundedHour
}

var cache_Data_Time_Component_boundedEnumSecond gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumSecond sync.Once

func Get_Data_Time_Component_boundedEnumSecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumSecond.Do(func() {
		cache_Data_Time_Component_boundedEnumSecond = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566(Rebox_Data_Time_Component_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedSecond()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094(Rebox_Data_Time_Component_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Time_Component_enumSecond()))))}
		}), gopurs_runtime.Int(int64(60)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return v_0
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((n_0.IntVal) >= (int64(0))) && ((n_0.IntVal) <= (int64(59))) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{n_0, true}
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
		})})))}
	})
	return cache_Data_Time_Component_boundedEnumSecond
}

var cache_Data_Time_Component_enumSecond gopurs_runtime.Value
var once_Data_Time_Component_enumSecond sync.Once

func Get_Data_Time_Component_enumSecond() gopurs_runtime.Value {
	once_Data_Time_Component_enumSecond.Do(func() {
		cache_Data_Time_Component_enumSecond = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntSub(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V3)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V3))})))}
	})
	return cache_Data_Time_Component_enumSecond
}

var cache_Data_Time_Component_boundedEnumMinute gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumMinute sync.Once

func Get_Data_Time_Component_boundedEnumMinute() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumMinute.Do(func() {
		cache_Data_Time_Component_boundedEnumMinute = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566(Rebox_Data_Time_Component_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedMinute()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094(Rebox_Data_Time_Component_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Time_Component_enumMinute()))))}
		}), gopurs_runtime.Int(int64(60)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return v_0
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((n_0.IntVal) >= (int64(0))) && ((n_0.IntVal) <= (int64(59))) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{n_0, true}
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
		})})))}
	})
	return cache_Data_Time_Component_boundedEnumMinute
}

var cache_Data_Time_Component_enumMinute gopurs_runtime.Value
var once_Data_Time_Component_enumMinute sync.Once

func Get_Data_Time_Component_enumMinute() gopurs_runtime.Value {
	once_Data_Time_Component_enumMinute.Do(func() {
		cache_Data_Time_Component_enumMinute = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntSub(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V3)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V3))})))}
	})
	return cache_Data_Time_Component_enumMinute
}

var cache_Data_Time_Component_boundedEnumMillisecond gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumMillisecond sync.Once

func Get_Data_Time_Component_boundedEnumMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumMillisecond.Do(func() {
		cache_Data_Time_Component_boundedEnumMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566(Rebox_Data_Time_Component_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedMillisecond()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094(Rebox_Data_Time_Component_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Time_Component_enumMillisecond()))))}
		}), gopurs_runtime.Int(int64(1000)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return v_0
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((n_0.IntVal) >= (int64(0))) && ((n_0.IntVal) <= (int64(999))) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{n_0, true}
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
		})})))}
	})
	return cache_Data_Time_Component_boundedEnumMillisecond
}

var cache_Data_Time_Component_enumMillisecond gopurs_runtime.Value
var once_Data_Time_Component_enumMillisecond sync.Once

func Get_Data_Time_Component_enumMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_enumMillisecond.Do(func() {
		cache_Data_Time_Component_enumMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntSub(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V3)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V3))})))}
	})
	return cache_Data_Time_Component_enumMillisecond
}

var cache_Data_Time_Component_boundedEnumHour gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumHour sync.Once

func Get_Data_Time_Component_boundedEnumHour() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumHour.Do(func() {
		cache_Data_Time_Component_boundedEnumHour = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3764732725_2094947566(Rebox_Data_Time_Component_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedHour()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094(Rebox_Data_Time_Component_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Time_Component_enumHour()))))}
		}), gopurs_runtime.Int(int64(24)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return v_0
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 bool
			}
			{
				if ((n_0.IntVal) >= (int64(0))) && ((n_0.IntVal) <= (int64(23))) {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{n_0, true}
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
		})})))}
	})
	return cache_Data_Time_Component_boundedEnumHour
}

var cache_Data_Time_Component_enumHour gopurs_runtime.Value
var once_Data_Time_Component_enumHour sync.Once

func Get_Data_Time_Component_enumHour() gopurs_runtime.Value {
	once_Data_Time_Component_enumHour.Do(func() {
		cache_Data_Time_Component_enumHour = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_3308271157_4177771502(Rebox_Data_Time_Component_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntSub(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V3)), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V4, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(v_0.IntVal, int64(1)))
		}), Rebox_Data_Time_Component_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V3))})))}
	})
	return cache_Data_Time_Component_enumHour
}

func Call_Data_Time_Component_Second(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Component_Minute(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Component_Millisecond(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Component_Hour(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Data_Time_Component_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_2094947566_3764732725(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	out.V2 = in.V2.IntVal
	return out
}

func Rebox_Data_Time_Component_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Time_Component_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_4060049525_556578094(in *Constructor_Data_Enum_Enum[int64]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_556578094_4060049525(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[int64])(unsafe.Pointer(in))
}
