package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_DateTime_Instant_fromJust gopurs_runtime.Value
var once_Data_DateTime_Instant_fromJust sync.Once

func Get_Data_DateTime_Instant_fromJust() gopurs_runtime.Value {
	once_Data_DateTime_Instant_fromJust.Do(func() {
		cache_Data_DateTime_Instant_fromJust = Get_Data_Maybe_fromJust__875433499()
	})
	return cache_Data_DateTime_Instant_fromJust
}

var cache_Data_DateTime_Instant_toEnum gopurs_runtime.Value
var once_Data_DateTime_Instant_toEnum sync.Once

func Get_Data_DateTime_Instant_toEnum() gopurs_runtime.Value {
	once_Data_DateTime_Instant_toEnum.Do(func() {
		cache_Data_DateTime_Instant_toEnum = Get_Data_Enum_toEnum__4195956870()
	})
	return cache_Data_DateTime_Instant_toEnum
}

var cache_Data_DateTime_Instant_go__append gopurs_runtime.Value
var once_Data_DateTime_Instant_go__append sync.Once

func Get_Data_DateTime_Instant_go__append() gopurs_runtime.Value {
	once_Data_DateTime_Instant_go__append.Do(func() {
		cache_Data_DateTime_Instant_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_DateTime_Instant_go__append
}

var cache_Data_DateTime_Instant_show gopurs_runtime.Value
var once_Data_DateTime_Instant_show sync.Once

func Get_Data_DateTime_Instant_show() gopurs_runtime.Value {
	once_Data_DateTime_Instant_show.Do(func() {
		cache_Data_DateTime_Instant_show = Get_Data_Show_show__2672875391()
	})
	return cache_Data_DateTime_Instant_show
}

var cache_Data_DateTime_Instant_conj gopurs_runtime.Value
var once_Data_DateTime_Instant_conj sync.Once

func Get_Data_DateTime_Instant_conj() gopurs_runtime.Value {
	once_Data_DateTime_Instant_conj.Do(func() {
		cache_Data_DateTime_Instant_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_DateTime_Instant_conj
}

var cache_Data_DateTime_Instant_greaterThanOrEq gopurs_runtime.Value
var once_Data_DateTime_Instant_greaterThanOrEq sync.Once

func Get_Data_DateTime_Instant_greaterThanOrEq() gopurs_runtime.Value {
	once_Data_DateTime_Instant_greaterThanOrEq.Do(func() {
		cache_Data_DateTime_Instant_greaterThanOrEq = Get_Data_Ord_greaterThanOrEq__756116985()
	})
	return cache_Data_DateTime_Instant_greaterThanOrEq
}

var cache_Data_DateTime_Instant_negate gopurs_runtime.Value
var once_Data_DateTime_Instant_negate sync.Once

func Get_Data_DateTime_Instant_negate() gopurs_runtime.Value {
	once_Data_DateTime_Instant_negate.Do(func() {
		cache_Data_DateTime_Instant_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_DateTime_Instant_negate(a_0_box.FloatVal()))
		})
	})
	return cache_Data_DateTime_Instant_negate
}

var cache_Data_DateTime_Instant_lessThanOrEq gopurs_runtime.Value
var once_Data_DateTime_Instant_lessThanOrEq sync.Once

func Get_Data_DateTime_Instant_lessThanOrEq() gopurs_runtime.Value {
	once_Data_DateTime_Instant_lessThanOrEq.Do(func() {
		cache_Data_DateTime_Instant_lessThanOrEq = Get_Data_Ord_lessThanOrEq__756116985()
	})
	return cache_Data_DateTime_Instant_lessThanOrEq
}

var cache_Data_DateTime_Instant_fromEnum gopurs_runtime.Value
var once_Data_DateTime_Instant_fromEnum sync.Once

func Get_Data_DateTime_Instant_fromEnum() gopurs_runtime.Value {
	once_Data_DateTime_Instant_fromEnum.Do(func() {
		cache_Data_DateTime_Instant_fromEnum = Get_Data_Enum_fromEnum__798184167()
	})
	return cache_Data_DateTime_Instant_fromEnum
}

var cache_Data_DateTime_Instant_bottom gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom sync.Once

func Get_Data_DateTime_Instant_bottom() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom.Do(func() {
		cache_Data_DateTime_Instant_bottom = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_DateTime_Instant_bottom
}

var cache_Data_DateTime_Instant_bottom1 gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom1 sync.Once

func Get_Data_DateTime_Instant_bottom1() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom1.Do(func() {
		cache_Data_DateTime_Instant_bottom1 = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_DateTime_Instant_bottom1
}

var cache_Data_DateTime_Instant_bottom2 gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom2 sync.Once

func Get_Data_DateTime_Instant_bottom2() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom2.Do(func() {
		cache_Data_DateTime_Instant_bottom2 = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_DateTime_Instant_bottom2
}

var cache_Data_DateTime_Instant_bottom3 gopurs_runtime.Value
var once_Data_DateTime_Instant_bottom3 sync.Once

func Get_Data_DateTime_Instant_bottom3() gopurs_runtime.Value {
	once_Data_DateTime_Instant_bottom3.Do(func() {
		cache_Data_DateTime_Instant_bottom3 = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_DateTime_Instant_bottom3
}

var cache_Data_DateTime_Instant_append1 gopurs_runtime.Value
var once_Data_DateTime_Instant_append1 sync.Once

func Get_Data_DateTime_Instant_append1() gopurs_runtime.Value {
	once_Data_DateTime_Instant_append1.Do(func() {
		cache_Data_DateTime_Instant_append1 = Get_Data_Semigroup_append__2418119035()
	})
	return cache_Data_DateTime_Instant_append1
}

var cache_Data_DateTime_Instant_negateDuration gopurs_runtime.Value
var once_Data_DateTime_Instant_negateDuration sync.Once

func Get_Data_DateTime_Instant_negateDuration() gopurs_runtime.Value {
	once_Data_DateTime_Instant_negateDuration.Do(func() {
		cache_Data_DateTime_Instant_negateDuration = Get_Data_Time_Duration_negateDuration__389498750()
	})
	return cache_Data_DateTime_Instant_negateDuration
}

var cache_Data_DateTime_Instant_Instant gopurs_runtime.Value
var once_Data_DateTime_Instant_Instant sync.Once

func Get_Data_DateTime_Instant_Instant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_Instant.Do(func() {
		cache_Data_DateTime_Instant_Instant = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_DateTime_Instant_Instant(x_0_box.FloatVal()))
		})
	})
	return cache_Data_DateTime_Instant_Instant
}

var cache_Data_DateTime_Instant_unInstant gopurs_runtime.Value
var once_Data_DateTime_Instant_unInstant sync.Once

func Get_Data_DateTime_Instant_unInstant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_unInstant.Do(func() {
		cache_Data_DateTime_Instant_unInstant = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_DateTime_Instant_unInstant(v_0_box.FloatVal()))
		})
	})
	return cache_Data_DateTime_Instant_unInstant
}

var cache_Data_DateTime_Instant_toDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_toDateTime sync.Once

func Get_Data_DateTime_Instant_toDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_toDateTime.Do(func() {
		cache_Data_DateTime_Instant_toDateTime = gopurs_runtime.Apply(Get_Data_DateTime_Instant_toDateTimeImpl(), gopurs_runtime.Func5(func(y_0 gopurs_runtime.Value, mo_1 gopurs_runtime.Value, d_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, mi_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func2(func(s_5 gopurs_runtime.Value, ms_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 uint32
				{
					if (mo_1.IntVal) == (int64(1)) {
						__t0 = 1908470532
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(2)) {
						__t0 = 2455627378
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(3)) {
						__t0 = 4162469099
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(4)) {
						__t0 = 1692989816
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(5)) {
						__t0 = 330658827
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(6)) {
						__t0 = 4067355978
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(7)) {
						__t0 = 2276710548
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(8)) {
						__t0 = 243771071
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(9)) {
						__t0 = 215731793
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(10)) {
						__t0 = 8639228
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(11)) {
						__t0 = 49471444
						goto end_branch_0
					} else {

					}
				}
				{
					if (mo_1.IntVal) == (int64(12)) {
						__t0 = 3889233761
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() uint32 { panic("Failed pattern match") }()
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer((&Constructor_Data_DateTime_DateTime{1, Call_Data_Date_canonicalDate(y_0.IntVal, __t0, d_2.IntVal), (&Constructor_Data_Time_Time{1, h_3.IntVal, mi_4.IntVal, s_5.IntVal, ms_6.IntVal})}))}
			})
		}))
	})
	return cache_Data_DateTime_Instant_toDateTime
}

var cache_Data_DateTime_Instant_showInstant gopurs_runtime.Value
var once_Data_DateTime_Instant_showInstant sync.Once

func Get_Data_DateTime_Instant_showInstant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_showInstant.Do(func() {
		cache_Data_DateTime_Instant_showInstant = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_Instant_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Instant (Milliseconds ", Data_Show_ShowNumberImpl(v_0.FloatVal())), "))"))
		})})))}
	})
	return cache_Data_DateTime_Instant_showInstant
}

var cache_Data_DateTime_Instant_ordDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_ordDateTime sync.Once

func Get_Data_DateTime_Instant_ordDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_ordDateTime.Do(func() {
		cache_Data_DateTime_Instant_ordDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_Instant_3047586294_4177771502(Rebox_Data_DateTime_Instant_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
	})
	return cache_Data_DateTime_Instant_ordDateTime
}

var cache_Data_DateTime_Instant_instant gopurs_runtime.Value
var once_Data_DateTime_Instant_instant sync.Once

func Get_Data_DateTime_Instant_instant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_instant.Do(func() {
		cache_Data_DateTime_Instant_instant = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_DateTime_Instant_instant(v_0_box.FloatVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_DateTime_Instant_instant
}

var cache_Data_DateTime_Instant_fromDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_fromDateTime sync.Once

func Get_Data_DateTime_Instant_fromDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_fromDateTime.Do(func() {
		cache_Data_DateTime_Instant_fromDateTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_DateTime_Instant_fromDateTime(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_0_box)))
		})
	})
	return cache_Data_DateTime_Instant_fromDateTime
}

var cache_Data_DateTime_Instant_fromDate gopurs_runtime.Value
var once_Data_DateTime_Instant_fromDate sync.Once

func Get_Data_DateTime_Instant_fromDate() gopurs_runtime.Value {
	once_Data_DateTime_Instant_fromDate.Do(func() {
		cache_Data_DateTime_Instant_fromDate = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_DateTime_Instant_fromDate(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](d_0_box)))
		})
	})
	return cache_Data_DateTime_Instant_fromDate
}

var cache_Data_DateTime_Instant_eqDateTime gopurs_runtime.Value
var once_Data_DateTime_Instant_eqDateTime sync.Once

func Get_Data_DateTime_Instant_eqDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Instant_eqDateTime.Do(func() {
		cache_Data_DateTime_Instant_eqDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_Instant_687527510_3790796878(Rebox_Data_DateTime_Instant_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
	})
	return cache_Data_DateTime_Instant_eqDateTime
}

var cache_Data_DateTime_Instant_diff gopurs_runtime.Value
var once_Data_DateTime_Instant_diff sync.Once

func Get_Data_DateTime_Instant_diff() gopurs_runtime.Value {
	once_Data_DateTime_Instant_diff.Do(func() {
		cache_Data_DateTime_Instant_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_DateTime_Instant_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box), dt1_1_box.FloatVal(), dt2_2_box.FloatVal())
		})
	})
	return cache_Data_DateTime_Instant_diff
}

var cache_Data_DateTime_Instant_boundedInstant gopurs_runtime.Value
var once_Data_DateTime_Instant_boundedInstant sync.Once

func Get_Data_DateTime_Instant_boundedInstant() gopurs_runtime.Value {
	once_Data_DateTime_Instant_boundedInstant.Do(func() {
		cache_Data_DateTime_Instant_boundedInstant = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_Instant_4294747894_2094947566((&Constructor_Data_Bounded_Bounded[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_Instant_3047586294_4177771502(Rebox_Data_DateTime_Instant_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
		}), -8639977881600000.0, 8639977881599999.0})))}
	})
	return cache_Data_DateTime_Instant_boundedInstant
}

func Call_Data_DateTime_Instant_negate(a_0_loop float64) float64 {
	var a_0 float64 = a_0_loop
	_ = a_0
	return -(a_0)
}

func Call_Data_DateTime_Instant_Instant(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_DateTime_Instant_unInstant(v_0_loop float64) float64 {
	var v_0 float64 = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_DateTime_Instant_instant(v_0_loop float64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 float64 = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[float64]
	{
		if ((v_0) >= (-8639977881600000.0)) && ((v_0) <= (8639977881599999.0)) {
			__t0 = Rebox_Data_DateTime_Instant_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Float(v_0), true}
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
		__t0 = Rebox_Data_DateTime_Instant_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_Instant_3240988860_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_DateTime_Instant_fromDateTime(v_0_loop *Constructor_Data_DateTime_DateTime) float64 {
	var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
	_ = v_0
	var __t12 int64
	{
		var __t_tag_0 uint32 = ((v_0).V0).V1
		_ = __t_tag_0
		if uint32(__t_tag_0) == 1908470532 {
			__t12 = int64(1)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_1 uint32 = ((v_0).V0).V1
		_ = __t_tag_1
		if uint32(__t_tag_1) == 2455627378 {
			__t12 = int64(2)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_2 uint32 = ((v_0).V0).V1
		_ = __t_tag_2
		if uint32(__t_tag_2) == 4162469099 {
			__t12 = int64(3)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_3 uint32 = ((v_0).V0).V1
		_ = __t_tag_3
		if uint32(__t_tag_3) == 1692989816 {
			__t12 = int64(4)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_4 uint32 = ((v_0).V0).V1
		_ = __t_tag_4
		if uint32(__t_tag_4) == 330658827 {
			__t12 = int64(5)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_5 uint32 = ((v_0).V0).V1
		_ = __t_tag_5
		if uint32(__t_tag_5) == 4067355978 {
			__t12 = int64(6)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_6 uint32 = ((v_0).V0).V1
		_ = __t_tag_6
		if uint32(__t_tag_6) == 2276710548 {
			__t12 = int64(7)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_7 uint32 = ((v_0).V0).V1
		_ = __t_tag_7
		if uint32(__t_tag_7) == 243771071 {
			__t12 = int64(8)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_8 uint32 = ((v_0).V0).V1
		_ = __t_tag_8
		if uint32(__t_tag_8) == 215731793 {
			__t12 = int64(9)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_9 uint32 = ((v_0).V0).V1
		_ = __t_tag_9
		if uint32(__t_tag_9) == 8639228 {
			__t12 = int64(10)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_10 uint32 = ((v_0).V0).V1
		_ = __t_tag_10
		if uint32(__t_tag_10) == 49471444 {
			__t12 = int64(11)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_11 uint32 = ((v_0).V0).V1
		_ = __t_tag_11
		if uint32(__t_tag_11) == 3889233761 {
			__t12 = int64(12)
			goto end_branch_12
		} else {

		}
	}
	{
		__t12 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_12:
	return gopurs_runtime.UncurriedApp7(Get_Data_DateTime_Instant_fromDateTimeImpl(), gopurs_runtime.Int(((v_0).V0).V0), gopurs_runtime.Int(__t12), gopurs_runtime.Int(((v_0).V0).V2), gopurs_runtime.Int(((v_0).V1).V0), gopurs_runtime.Int(((v_0).V1).V1), gopurs_runtime.Int(((v_0).V1).V2), gopurs_runtime.Int(((v_0).V1).V3)).FloatVal()
}

func Call_Data_DateTime_Instant_fromDate(d_0_loop *Constructor_Data_Date_Date) float64 {
	var d_0 *Constructor_Data_Date_Date = d_0_loop
	_ = d_0
	var __t12 int64
	{
		var __t_tag_0 uint32 = (d_0).V1
		_ = __t_tag_0
		if uint32(__t_tag_0) == 1908470532 {
			__t12 = int64(1)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_1 uint32 = (d_0).V1
		_ = __t_tag_1
		if uint32(__t_tag_1) == 2455627378 {
			__t12 = int64(2)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_2 uint32 = (d_0).V1
		_ = __t_tag_2
		if uint32(__t_tag_2) == 4162469099 {
			__t12 = int64(3)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_3 uint32 = (d_0).V1
		_ = __t_tag_3
		if uint32(__t_tag_3) == 1692989816 {
			__t12 = int64(4)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_4 uint32 = (d_0).V1
		_ = __t_tag_4
		if uint32(__t_tag_4) == 330658827 {
			__t12 = int64(5)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_5 uint32 = (d_0).V1
		_ = __t_tag_5
		if uint32(__t_tag_5) == 4067355978 {
			__t12 = int64(6)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_6 uint32 = (d_0).V1
		_ = __t_tag_6
		if uint32(__t_tag_6) == 2276710548 {
			__t12 = int64(7)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_7 uint32 = (d_0).V1
		_ = __t_tag_7
		if uint32(__t_tag_7) == 243771071 {
			__t12 = int64(8)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_8 uint32 = (d_0).V1
		_ = __t_tag_8
		if uint32(__t_tag_8) == 215731793 {
			__t12 = int64(9)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_9 uint32 = (d_0).V1
		_ = __t_tag_9
		if uint32(__t_tag_9) == 8639228 {
			__t12 = int64(10)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_10 uint32 = (d_0).V1
		_ = __t_tag_10
		if uint32(__t_tag_10) == 49471444 {
			__t12 = int64(11)
			goto end_branch_12
		} else {

		}
	}
	{
		var __t_tag_11 uint32 = (d_0).V1
		_ = __t_tag_11
		if uint32(__t_tag_11) == 3889233761 {
			__t12 = int64(12)
			goto end_branch_12
		} else {

		}
	}
	{
		__t12 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_12:
	return gopurs_runtime.UncurriedApp7(Get_Data_DateTime_Instant_fromDateTimeImpl(), gopurs_runtime.Int((d_0).V0), gopurs_runtime.Int(__t12), gopurs_runtime.Int((d_0).V2), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(0))).FloatVal()
}

func Call_Data_DateTime_Instant_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value], dt1_1_loop float64, dt2_2_loop float64) gopurs_runtime.Value {
	var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
	_ = dictDuration_0
	var dt1_1 float64 = dt1_1_loop
	_ = dt1_1
	var dt2_2 float64 = dt2_2_loop
	_ = dt2_2
	return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Float(gopurs_runtime.Float((dt1_1)+(Call_Data_Time_Duration_negateDuration__389498750(dt2_2))).FloatVal()))
}

func Rebox_Data_DateTime_Instant_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_DateTime_Instant_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_DateTime_Instant_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Data_DateTime_Instant_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_DateTime_Instant_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[float64])(unsafe.Pointer(in))
}

func Rebox_Data_DateTime_Instant_4177771502_3047586294(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[float64])(unsafe.Pointer(in))
}

func Rebox_Data_DateTime_Instant_4294747894_2094947566(in *Constructor_Data_Bounded_Bounded[float64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Float(in.V1)
	out.V2 = gopurs_runtime.Float(in.V2)
	return out
}

func Rebox_Data_DateTime_Instant_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_DateTime_Instant_fromDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_Instant_FromDateTimeImpl
}

func Get_Data_DateTime_Instant_toDateTimeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_Instant_ToDateTimeImpl
}
