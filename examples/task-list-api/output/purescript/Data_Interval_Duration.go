package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Interval_Duration_go__append gopurs_runtime.Value
var once_Data_Interval_Duration_go__append sync.Once

func Get_Data_Interval_Duration_go__append() gopurs_runtime.Value {
	once_Data_Interval_Duration_go__append.Do(func() {
		cache_Data_Interval_Duration_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Interval_Duration_go__append
}

var cache_Data_Interval_Duration_add gopurs_runtime.Value
var once_Data_Interval_Duration_add sync.Once

func Get_Data_Interval_Duration_add() gopurs_runtime.Value {
	once_Data_Interval_Duration_add.Do(func() {
		cache_Data_Interval_Duration_add = Get_Data_Semiring_numAdd()
	})
	return cache_Data_Interval_Duration_add
}

var cache_Data_Interval_Duration_div gopurs_runtime.Value
var once_Data_Interval_Duration_div sync.Once

func Get_Data_Interval_Duration_div() gopurs_runtime.Value {
	once_Data_Interval_Duration_div.Do(func() {
		cache_Data_Interval_Duration_div = Get_Data_EuclideanRing_div__266314078()
	})
	return cache_Data_Interval_Duration_div
}

var cache_Data_Interval_Duration_Second gopurs_runtime.Value
var once_Data_Interval_Duration_Second sync.Once

func Get_Data_Interval_Duration_Second() gopurs_runtime.Value {
	once_Data_Interval_Duration_Second.Do(func() {
		cache_Data_Interval_Duration_Second = gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Second
}

var cache_Data_Interval_Duration_Minute gopurs_runtime.Value
var once_Data_Interval_Duration_Minute sync.Once

func Get_Data_Interval_Duration_Minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_Minute.Do(func() {
		cache_Data_Interval_Duration_Minute = gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Minute
}

var cache_Data_Interval_Duration_Hour gopurs_runtime.Value
var once_Data_Interval_Duration_Hour sync.Once

func Get_Data_Interval_Duration_Hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_Hour.Do(func() {
		cache_Data_Interval_Duration_Hour = gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Hour
}

var cache_Data_Interval_Duration_Day gopurs_runtime.Value
var once_Data_Interval_Duration_Day sync.Once

func Get_Data_Interval_Duration_Day() gopurs_runtime.Value {
	once_Data_Interval_Duration_Day.Do(func() {
		cache_Data_Interval_Duration_Day = gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Day
}

var cache_Data_Interval_Duration_Week gopurs_runtime.Value
var once_Data_Interval_Duration_Week sync.Once

func Get_Data_Interval_Duration_Week() gopurs_runtime.Value {
	once_Data_Interval_Duration_Week.Do(func() {
		cache_Data_Interval_Duration_Week = gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Week
}

var cache_Data_Interval_Duration_Month gopurs_runtime.Value
var once_Data_Interval_Duration_Month sync.Once

func Get_Data_Interval_Duration_Month() gopurs_runtime.Value {
	once_Data_Interval_Duration_Month.Do(func() {
		cache_Data_Interval_Duration_Month = gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Month
}

var cache_Data_Interval_Duration_Year gopurs_runtime.Value
var once_Data_Interval_Duration_Year sync.Once

func Get_Data_Interval_Duration_Year() gopurs_runtime.Value {
	once_Data_Interval_Duration_Year.Do(func() {
		cache_Data_Interval_Duration_Year = gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Year
}

var cache_Data_Interval_Duration_Duration gopurs_runtime.Value
var once_Data_Interval_Duration_Duration sync.Once

func Get_Data_Interval_Duration_Duration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Duration.Do(func() {
		cache_Data_Interval_Duration_Duration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_Duration(x_0_box)
		})
	})
	return cache_Data_Interval_Duration_Duration
}

var cache_Data_Interval_Duration_showDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_showDurationComponent sync.Once

func Get_Data_Interval_Duration_showDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDurationComponent.Do(func() {
		cache_Data_Interval_Duration_showDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t7 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 217821258 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Minute", 6)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 3908053364 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Second", 6)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 1292308612 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Hour", 4)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 2311060696 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Day", 3)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 401302776 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Week", 4)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 3327533908 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Month", 5)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_6 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_6
				if uint32(__t_tag_6) == 3631736139 {
					__t7 = gopurs_runtime.StrWithCodeUnits("Year", 4)
					goto end_branch_7
				} else {

				}
			}
			{
				__t7 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
			}
		end_branch_7:
			return __t7
		})})))}
	})
	return cache_Data_Interval_Duration_showDurationComponent
}

var cache_Data_Interval_Duration_show gopurs_runtime.Value
var once_Data_Interval_Duration_show sync.Once

func Get_Data_Interval_Duration_show() gopurs_runtime.Value {
	once_Data_Interval_Duration_show.Do(func() {
		cache_Data_Interval_Duration_show = Get_Data_Show_show__2692513315()
	})
	return cache_Data_Interval_Duration_show
}

var cache_Data_Interval_Duration_showDuration gopurs_runtime.Value
var once_Data_Interval_Duration_showDuration sync.Once

func Get_Data_Interval_Duration_showDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDuration.Do(func() {
		cache_Data_Interval_Duration_showDuration = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Duration ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Map_Internal_showMap(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1209612131_1386611502(Rebox_Data_Interval_Duration_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Interval_Duration_showDurationComponent()))))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3263178038_1386611502(Rebox_Data_Interval_Duration_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))))}), "show"), v_0).StrVal()), ")"))
		})}))}
	})
	return cache_Data_Interval_Duration_showDuration
}

var cache_Data_Interval_Duration_newtypeDuration gopurs_runtime.Value
var once_Data_Interval_Duration_newtypeDuration sync.Once

func Get_Data_Interval_Duration_newtypeDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_newtypeDuration.Do(func() {
		cache_Data_Interval_Duration_newtypeDuration = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Interval_Duration_newtypeDuration
}

var cache_Data_Interval_Duration_eqDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_eqDurationComponent sync.Once

func Get_Data_Interval_Duration_eqDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDurationComponent.Do(func() {
		cache_Data_Interval_Duration_eqDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t15 bool
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 3908053364 {
					var __t_tag_4 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_4
					__t15 = (uint32(__t_tag_4) == 3908053364)
					goto end_branch_15
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 217821258 {
					var __t_tag_6 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_6
					__t15 = (uint32(__t_tag_6) == 217821258)
					goto end_branch_15
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 1292308612 {
					var __t_tag_8 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_8
					__t15 = (uint32(__t_tag_8) == 1292308612)
					goto end_branch_15
				} else {

				}
			}
			{
				var __t_tag_9 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_9
				if uint32(__t_tag_9) == 2311060696 {
					var __t_tag_10 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_10
					__t15 = (uint32(__t_tag_10) == 2311060696)
					goto end_branch_15
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_11
				if uint32(__t_tag_11) == 401302776 {
					var __t_tag_12 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_12
					__t15 = (uint32(__t_tag_12) == 401302776)
					goto end_branch_15
				} else {

				}
			}
			{
				var __t_tag_13 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_13
				if uint32(__t_tag_13) == 3327533908 {
					var __t_tag_14 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_14
					__t15 = (uint32(__t_tag_14) == 3327533908)
					goto end_branch_15
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 3631736139 {

					var __t_tag_1 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 3631736139)
				}
				__t15 = __t_and_2
			}
		end_branch_15:
			return gopurs_runtime.Bool(__t15)
		})})))}
	})
	return cache_Data_Interval_Duration_eqDurationComponent
}

var cache_Data_Interval_Duration_eq gopurs_runtime.Value
var once_Data_Interval_Duration_eq sync.Once

func Get_Data_Interval_Duration_eq() gopurs_runtime.Value {
	once_Data_Interval_Duration_eq.Do(func() {
		cache_Data_Interval_Duration_eq = Get_Data_Eq_eq__587294585()
	})
	return cache_Data_Interval_Duration_eq
}

var cache_Data_Interval_Duration_ordDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_ordDurationComponent sync.Once

func Get_Data_Interval_Duration_ordDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDurationComponent.Do(func() {
		cache_Data_Interval_Duration_ordDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878(Rebox_Data_Interval_Duration_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Interval_Duration_eqDurationComponent()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t27 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 3908053364 {
					var __t2 gopurs_runtime.Value
					{
						var __t_tag_1 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_1
						if uint32(__t_tag_1) == 3908053364 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_2:
					__t27 = __t2
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 3908053364 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 217821258 {
					var __t6 gopurs_runtime.Value
					{
						var __t_tag_5 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_5
						if uint32(__t_tag_5) == 217821258 {
							__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_6:
					__t27 = __t6
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 217821258 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_8
				if uint32(__t_tag_8) == 1292308612 {
					var __t10 gopurs_runtime.Value
					{
						var __t_tag_9 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_9
						if uint32(__t_tag_9) == 1292308612 {
							__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_10:
					__t27 = __t10
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_11
				if uint32(__t_tag_11) == 1292308612 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_12 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_12
				if uint32(__t_tag_12) == 2311060696 {
					var __t14 gopurs_runtime.Value
					{
						var __t_tag_13 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_13
						if uint32(__t_tag_13) == 2311060696 {
							__t14 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_14
						} else {

						}
					}
					{
						__t14 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_14:
					__t27 = __t14
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_15 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_15
				if uint32(__t_tag_15) == 2311060696 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_16 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_16
				if uint32(__t_tag_16) == 401302776 {
					var __t18 gopurs_runtime.Value
					{
						var __t_tag_17 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_17
						if uint32(__t_tag_17) == 401302776 {
							__t18 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_18
						} else {

						}
					}
					{
						__t18 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_18:
					__t27 = __t18
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_19 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_19
				if uint32(__t_tag_19) == 401302776 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_20 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_20
				if uint32(__t_tag_20) == 3327533908 {
					var __t22 gopurs_runtime.Value
					{
						var __t_tag_21 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_21
						if uint32(__t_tag_21) == 3327533908 {
							__t22 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_22
						} else {

						}
					}
					{
						__t22 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_22:
					__t27 = __t22
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_23 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_23
				if uint32(__t_tag_23) == 3327533908 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				var __t_tag_24 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_24
				var __t_and_26 bool = false
				if uint32(__t_tag_24) == 3631736139 {

					var __t_tag_25 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_25
					__t_and_26 = (uint32(__t_tag_25) == 3631736139)
				}
				if __t_and_26 {
					__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
					goto end_branch_27
				} else {

				}
			}
			{
				__t27 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_27:
			return __t27
		})})))}
	})
	return cache_Data_Interval_Duration_ordDurationComponent
}

var cache_Data_Interval_Duration_unionWith gopurs_runtime.Value
var once_Data_Interval_Duration_unionWith sync.Once

func Get_Data_Interval_Duration_unionWith() gopurs_runtime.Value {
	once_Data_Interval_Duration_unionWith.Do(func() {
		cache_Data_Interval_Duration_unionWith = Get_Data_Map_Internal_unionWith__2491407586()
	})
	return cache_Data_Interval_Duration_unionWith
}

var cache_Data_Interval_Duration_compare gopurs_runtime.Value
var once_Data_Interval_Duration_compare sync.Once

func Get_Data_Interval_Duration_compare() gopurs_runtime.Value {
	once_Data_Interval_Duration_compare.Do(func() {
		cache_Data_Interval_Duration_compare = Get_Data_Ord_compare__78251971()
	})
	return cache_Data_Interval_Duration_compare
}

var cache_Data_Interval_Duration_semigroupDuration gopurs_runtime.Value
var once_Data_Interval_Duration_semigroupDuration sync.Once

func Get_Data_Interval_Duration_semigroupDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_semigroupDuration.Do(func() {
		cache_Data_Interval_Duration_semigroupDuration = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply4(Get_Data_Map_Internal_unionWithImpl(), gopurs_runtime.Apply(Get_Data_Map_Internal_compareInt(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3730953251_4177771502(Rebox_Data_Interval_Duration_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDurationComponent()))))}), Get_Data_Semiring_numAdd(), v_0, v1_1)
		})}))}
	})
	return cache_Data_Interval_Duration_semigroupDuration
}

var cache_Data_Interval_Duration_monoidDuration gopurs_runtime.Value
var once_Data_Interval_Duration_monoidDuration sync.Once

func Get_Data_Interval_Duration_monoidDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_monoidDuration.Do(func() {
		cache_Data_Interval_Duration_monoidDuration = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Interval_Duration_semigroupDuration()))}
		}), Get_Data_Map_Internal_empty()}))}
	})
	return cache_Data_Interval_Duration_monoidDuration
}

var cache_Data_Interval_Duration_eqDuration gopurs_runtime.Value
var once_Data_Interval_Duration_eqDuration sync.Once

func Get_Data_Interval_Duration_eqDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDuration.Do(func() {
		cache_Data_Interval_Duration_eqDuration = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Map_Internal_eqMap(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878(Rebox_Data_Interval_Duration_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Interval_Duration_eqDurationComponent()))))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_687527510_3790796878(Rebox_Data_Interval_Duration_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}), "eq"), x_0, y_1)
		})}))}
	})
	return cache_Data_Interval_Duration_eqDuration
}

var cache_Data_Interval_Duration_ordDuration gopurs_runtime.Value
var once_Data_Interval_Duration_ordDuration sync.Once

func Get_Data_Interval_Duration_ordDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDuration.Do(func() {
		cache_Data_Interval_Duration_ordDuration = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Interval_Duration_eqDuration()))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Map_Internal_ordMap(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3730953251_4177771502(Rebox_Data_Interval_Duration_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Interval_Duration_ordDurationComponent()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3047586294_4177771502(Rebox_Data_Interval_Duration_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}), "compare"), x_0, y_1)
		})}))}
	})
	return cache_Data_Interval_Duration_ordDuration
}

var cache_Data_Interval_Duration_durationFromComponent gopurs_runtime.Value
var once_Data_Interval_Duration_durationFromComponent sync.Once

func Get_Data_Interval_Duration_durationFromComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_durationFromComponent.Do(func() {
		cache_Data_Interval_Duration_durationFromComponent = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_durationFromComponent(uint32(k_0_box.IntVal), v_1_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_durationFromComponent
}

var cache_Data_Interval_Duration_hour gopurs_runtime.Value
var once_Data_Interval_Duration_hour sync.Once

func Get_Data_Interval_Duration_hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_hour.Do(func() {
		cache_Data_Interval_Duration_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_hour(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_hour
}

var cache_Data_Interval_Duration_millisecond gopurs_runtime.Value
var once_Data_Interval_Duration_millisecond sync.Once

func Get_Data_Interval_Duration_millisecond() gopurs_runtime.Value {
	once_Data_Interval_Duration_millisecond.Do(func() {
		cache_Data_Interval_Duration_millisecond = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(v_0.FloatVal()))
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) / (1000.0))
		}))
	})
	return cache_Data_Interval_Duration_millisecond
}

var cache_Data_Interval_Duration_minute gopurs_runtime.Value
var once_Data_Interval_Duration_minute sync.Once

func Get_Data_Interval_Duration_minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_minute.Do(func() {
		cache_Data_Interval_Duration_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_minute(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_minute
}

var cache_Data_Interval_Duration_month gopurs_runtime.Value
var once_Data_Interval_Duration_month sync.Once

func Get_Data_Interval_Duration_month() gopurs_runtime.Value {
	once_Data_Interval_Duration_month.Do(func() {
		cache_Data_Interval_Duration_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_month(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_month
}

var cache_Data_Interval_Duration_second gopurs_runtime.Value
var once_Data_Interval_Duration_second sync.Once

func Get_Data_Interval_Duration_second() gopurs_runtime.Value {
	once_Data_Interval_Duration_second.Do(func() {
		cache_Data_Interval_Duration_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_second(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_second
}

var cache_Data_Interval_Duration_week gopurs_runtime.Value
var once_Data_Interval_Duration_week sync.Once

func Get_Data_Interval_Duration_week() gopurs_runtime.Value {
	once_Data_Interval_Duration_week.Do(func() {
		cache_Data_Interval_Duration_week = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_week(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_week
}

var cache_Data_Interval_Duration_year gopurs_runtime.Value
var once_Data_Interval_Duration_year sync.Once

func Get_Data_Interval_Duration_year() gopurs_runtime.Value {
	once_Data_Interval_Duration_year.Do(func() {
		cache_Data_Interval_Duration_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_year(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_year
}

var cache_Data_Interval_Duration_day gopurs_runtime.Value
var once_Data_Interval_Duration_day sync.Once

func Get_Data_Interval_Duration_day() gopurs_runtime.Value {
	once_Data_Interval_Duration_day.Do(func() {
		cache_Data_Interval_Duration_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_Duration_day(v_0_box.FloatVal())
		})
	})
	return cache_Data_Interval_Duration_day
}

type Constructor_Data_Interval_Duration_Second struct {
	Rc uint32
}

type Constructor_Data_Interval_Duration_Minute struct {
	Rc uint32
}

type Constructor_Data_Interval_Duration_Hour struct {
	Rc uint32
}

type Constructor_Data_Interval_Duration_Day struct {
	Rc uint32
}

type Constructor_Data_Interval_Duration_Week struct {
	Rc uint32
}

type Constructor_Data_Interval_Duration_Month struct {
	Rc uint32
}

type Constructor_Data_Interval_Duration_Year struct {
	Rc uint32
}

func Call_Data_Interval_Duration_Duration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Interval_Duration_durationFromComponent(k_0_loop uint32, v_1_loop float64) gopurs_runtime.Value {
	var k_0 uint32 = k_0_loop
	_ = k_0
	var v_1 float64 = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Float(v_1))
}

func Call_Data_Interval_Duration_hour(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Call_Data_Interval_Duration_minute(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Call_Data_Interval_Duration_month(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Call_Data_Interval_Duration_second(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Call_Data_Interval_Duration_week(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Call_Data_Interval_Duration_year(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Call_Data_Interval_Duration_day(v_0_loop float64) gopurs_runtime.Value {
	var v_0 float64 = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Get_Data_Map_Internal_singleton(), gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}, gopurs_runtime.Float(v_0))
}

func Rebox_Data_Interval_Duration_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_1386611502_1209612131(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_4177771502_3047586294(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_Duration_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
