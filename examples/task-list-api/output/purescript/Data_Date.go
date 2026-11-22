package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_fromEnum gopurs_runtime.Value
var once_Data_Date_fromEnum sync.Once

func Get_Data_Date_fromEnum() gopurs_runtime.Value {
	once_Data_Date_fromEnum.Do(func() {
		cache_Data_Date_fromEnum = Get_Data_Enum_fromEnum__798184167()
	})
	return cache_Data_Date_fromEnum
}

var cache_Data_Date_eq gopurs_runtime.Value
var once_Data_Date_eq sync.Once

func Get_Data_Date_eq() gopurs_runtime.Value {
	once_Data_Date_eq.Do(func() {
		cache_Data_Date_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Data_Date_eq
}

var cache_Data_Date_fromJust gopurs_runtime.Value
var once_Data_Date_fromJust sync.Once

func Get_Data_Date_fromJust() gopurs_runtime.Value {
	once_Data_Date_fromJust.Do(func() {
		cache_Data_Date_fromJust = Get_Data_Maybe_fromJust__1345383707()
	})
	return cache_Data_Date_fromJust
}

var cache_Data_Date_toEnum gopurs_runtime.Value
var once_Data_Date_toEnum sync.Once

func Get_Data_Date_toEnum() gopurs_runtime.Value {
	once_Data_Date_toEnum.Do(func() {
		cache_Data_Date_toEnum = Get_Data_Enum_toEnum__890235542()
	})
	return cache_Data_Date_toEnum
}

var cache_Data_Date_go__append gopurs_runtime.Value
var once_Data_Date_go__append sync.Once

func Get_Data_Date_go__append() gopurs_runtime.Value {
	once_Data_Date_go__append.Do(func() {
		cache_Data_Date_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Date_go__append
}

var cache_Data_Date_show gopurs_runtime.Value
var once_Data_Date_show sync.Once

func Get_Data_Date_show() gopurs_runtime.Value {
	once_Data_Date_show.Do(func() {
		cache_Data_Date_show = Get_Data_Show_show__1148161774()
	})
	return cache_Data_Date_show
}

var cache_Data_Date_show1 gopurs_runtime.Value
var once_Data_Date_show1 sync.Once

func Get_Data_Date_show1() gopurs_runtime.Value {
	once_Data_Date_show1.Do(func() {
		cache_Data_Date_show1 = Get_Data_Show_show__1423355633()
	})
	return cache_Data_Date_show1
}

var cache_Data_Date_show2 gopurs_runtime.Value
var once_Data_Date_show2 sync.Once

func Get_Data_Date_show2() gopurs_runtime.Value {
	once_Data_Date_show2.Do(func() {
		cache_Data_Date_show2 = Get_Data_Show_show__3736159005()
	})
	return cache_Data_Date_show2
}

var cache_Data_Date_fromEnum1 gopurs_runtime.Value
var once_Data_Date_fromEnum1 sync.Once

func Get_Data_Date_fromEnum1() gopurs_runtime.Value {
	once_Data_Date_fromEnum1.Do(func() {
		cache_Data_Date_fromEnum1 = Get_Data_Enum_fromEnum__2307775192()
	})
	return cache_Data_Date_fromEnum1
}

var cache_Data_Date_conj gopurs_runtime.Value
var once_Data_Date_conj sync.Once

func Get_Data_Date_conj() gopurs_runtime.Value {
	once_Data_Date_conj.Do(func() {
		cache_Data_Date_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Date_conj
}

var cache_Data_Date_mod gopurs_runtime.Value
var once_Data_Date_mod sync.Once

func Get_Data_Date_mod() gopurs_runtime.Value {
	once_Data_Date_mod.Do(func() {
		cache_Data_Date_mod = Get_Data_EuclideanRing_mod__478053998()
	})
	return cache_Data_Date_mod
}

var cache_Data_Date_disj gopurs_runtime.Value
var once_Data_Date_disj sync.Once

func Get_Data_Date_disj() gopurs_runtime.Value {
	once_Data_Date_disj.Do(func() {
		cache_Data_Date_disj = Get_Data_HeytingAlgebra_disj__2421783737()
	})
	return cache_Data_Date_disj
}

var cache_Data_Date_not gopurs_runtime.Value
var once_Data_Date_not sync.Once

func Get_Data_Date_not() gopurs_runtime.Value {
	once_Data_Date_not.Do(func() {
		cache_Data_Date_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Data_Date_not
}

var cache_Data_Date_compose gopurs_runtime.Value
var once_Data_Date_compose sync.Once

func Get_Data_Date_compose() gopurs_runtime.Value {
	once_Data_Date_compose.Do(func() {
		cache_Data_Date_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Date_compose
}

var cache_Data_Date_fromJust1 gopurs_runtime.Value
var once_Data_Date_fromJust1 sync.Once

func Get_Data_Date_fromJust1() gopurs_runtime.Value {
	once_Data_Date_fromJust1.Do(func() {
		cache_Data_Date_fromJust1 = Get_Data_Maybe_fromJust__3338668507()
	})
	return cache_Data_Date_fromJust1
}

var cache_Data_Date_toEnum1 gopurs_runtime.Value
var once_Data_Date_toEnum1 sync.Once

func Get_Data_Date_toEnum1() gopurs_runtime.Value {
	once_Data_Date_toEnum1.Do(func() {
		cache_Data_Date_toEnum1 = Rebox_Data_Date_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())).V4
	})
	return cache_Data_Date_toEnum1
}

var cache_Data_Date_eq1 gopurs_runtime.Value
var once_Data_Date_eq1 sync.Once

func Get_Data_Date_eq1() gopurs_runtime.Value {
	once_Data_Date_eq1.Do(func() {
		cache_Data_Date_eq1 = Get_Data_Eq_eq__2921765156()
	})
	return cache_Data_Date_eq1
}

var cache_Data_Date_eq2 gopurs_runtime.Value
var once_Data_Date_eq2 sync.Once

func Get_Data_Date_eq2() gopurs_runtime.Value {
	once_Data_Date_eq2.Do(func() {
		cache_Data_Date_eq2 = Get_Data_Eq_eq__2904741602()
	})
	return cache_Data_Date_eq2
}

var cache_Data_Date_eq3 gopurs_runtime.Value
var once_Data_Date_eq3 sync.Once

func Get_Data_Date_eq3() gopurs_runtime.Value {
	once_Data_Date_eq3.Do(func() {
		cache_Data_Date_eq3 = Get_Data_Eq_eq__2807846199()
	})
	return cache_Data_Date_eq3
}

var cache_Data_Date_compare gopurs_runtime.Value
var once_Data_Date_compare sync.Once

func Get_Data_Date_compare() gopurs_runtime.Value {
	once_Data_Date_compare.Do(func() {
		cache_Data_Date_compare = Get_Data_Ord_compare__1452131379()
	})
	return cache_Data_Date_compare
}

var cache_Data_Date_compare1 gopurs_runtime.Value
var once_Data_Date_compare1 sync.Once

func Get_Data_Date_compare1() gopurs_runtime.Value {
	once_Data_Date_compare1.Do(func() {
		cache_Data_Date_compare1 = Get_Data_Ord_compare__4139028341()
	})
	return cache_Data_Date_compare1
}

var cache_Data_Date_compare2 gopurs_runtime.Value
var once_Data_Date_compare2 sync.Once

func Get_Data_Date_compare2() gopurs_runtime.Value {
	once_Data_Date_compare2.Do(func() {
		cache_Data_Date_compare2 = Get_Data_Ord_compare__1544596384()
	})
	return cache_Data_Date_compare2
}

var cache_Data_Date_succ gopurs_runtime.Value
var once_Data_Date_succ sync.Once

func Get_Data_Date_succ() gopurs_runtime.Value {
	once_Data_Date_succ.Do(func() {
		cache_Data_Date_succ = Get_Data_Enum_succ__77857599()
	})
	return cache_Data_Date_succ
}

var cache_Data_Date_succ1 gopurs_runtime.Value
var once_Data_Date_succ1 sync.Once

func Get_Data_Date_succ1() gopurs_runtime.Value {
	once_Data_Date_succ1.Do(func() {
		cache_Data_Date_succ1 = Get_Data_Enum_succ__3573123786()
	})
	return cache_Data_Date_succ1
}

var cache_Data_Date_greaterThan gopurs_runtime.Value
var once_Data_Date_greaterThan sync.Once

func Get_Data_Date_greaterThan() gopurs_runtime.Value {
	once_Data_Date_greaterThan.Do(func() {
		cache_Data_Date_greaterThan = Get_Data_Ord_greaterThan__2273689272()
	})
	return cache_Data_Date_greaterThan
}

var cache_Data_Date_succ2 gopurs_runtime.Value
var once_Data_Date_succ2 sync.Once

func Get_Data_Date_succ2() gopurs_runtime.Value {
	once_Data_Date_succ2.Do(func() {
		cache_Data_Date_succ2 = Get_Data_Enum_succ__2729625849()
	})
	return cache_Data_Date_succ2
}

var cache_Data_Date_toEnum2 gopurs_runtime.Value
var once_Data_Date_toEnum2 sync.Once

func Get_Data_Date_toEnum2() gopurs_runtime.Value {
	once_Data_Date_toEnum2.Do(func() {
		cache_Data_Date_toEnum2 = Get_Data_Enum_toEnum__706605770()
	})
	return cache_Data_Date_toEnum2
}

var cache_Data_Date_apply gopurs_runtime.Value
var once_Data_Date_apply sync.Once

func Get_Data_Date_apply() gopurs_runtime.Value {
	once_Data_Date_apply.Do(func() {
		cache_Data_Date_apply = Get_Control_Apply_apply__2672146207()
	})
	return cache_Data_Date_apply
}

var cache_Data_Date_apply1 gopurs_runtime.Value
var once_Data_Date_apply1 sync.Once

func Get_Data_Date_apply1() gopurs_runtime.Value {
	once_Data_Date_apply1.Do(func() {
		cache_Data_Date_apply1 = Get_Control_Apply_apply__1827729951()
	})
	return cache_Data_Date_apply1
}

var cache_Data_Date_go__map gopurs_runtime.Value
var once_Data_Date_go__map sync.Once

func Get_Data_Date_go__map() gopurs_runtime.Value {
	once_Data_Date_go__map.Do(func() {
		cache_Data_Date_go__map = Get_Data_Functor_map__3216300106()
	})
	return cache_Data_Date_go__map
}

var cache_Data_Date_pure gopurs_runtime.Value
var once_Data_Date_pure sync.Once

func Get_Data_Date_pure() gopurs_runtime.Value {
	once_Data_Date_pure.Do(func() {
		cache_Data_Date_pure = Get_Control_Applicative_pure__1137267138()
	})
	return cache_Data_Date_pure
}

var cache_Data_Date_pred gopurs_runtime.Value
var once_Data_Date_pred sync.Once

func Get_Data_Date_pred() gopurs_runtime.Value {
	once_Data_Date_pred.Do(func() {
		cache_Data_Date_pred = Get_Data_Enum_pred__77857599()
	})
	return cache_Data_Date_pred
}

var cache_Data_Date_pred1 gopurs_runtime.Value
var once_Data_Date_pred1 sync.Once

func Get_Data_Date_pred1() gopurs_runtime.Value {
	once_Data_Date_pred1.Do(func() {
		cache_Data_Date_pred1 = Get_Data_Enum_pred__3573123786()
	})
	return cache_Data_Date_pred1
}

var cache_Data_Date_pred2 gopurs_runtime.Value
var once_Data_Date_pred2 sync.Once

func Get_Data_Date_pred2() gopurs_runtime.Value {
	once_Data_Date_pred2.Do(func() {
		cache_Data_Date_pred2 = Get_Data_Enum_pred__2729625849()
	})
	return cache_Data_Date_pred2
}

var cache_Data_Date_fromJust2 gopurs_runtime.Value
var once_Data_Date_fromJust2 sync.Once

func Get_Data_Date_fromJust2() gopurs_runtime.Value {
	once_Data_Date_fromJust2.Do(func() {
		cache_Data_Date_fromJust2 = Get_Data_Maybe_fromJust__875433499()
	})
	return cache_Data_Date_fromJust2
}

var cache_Data_Date_toEnum3 gopurs_runtime.Value
var once_Data_Date_toEnum3 sync.Once

func Get_Data_Date_toEnum3() gopurs_runtime.Value {
	once_Data_Date_toEnum3.Do(func() {
		cache_Data_Date_toEnum3 = Get_Data_Enum_toEnum__4195956870()
	})
	return cache_Data_Date_toEnum3
}

var cache_Data_Date_add gopurs_runtime.Value
var once_Data_Date_add sync.Once

func Get_Data_Date_add() gopurs_runtime.Value {
	once_Data_Date_add.Do(func() {
		cache_Data_Date_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Data_Date_add
}

var cache_Data_Date_fromEnum2 gopurs_runtime.Value
var once_Data_Date_fromEnum2 sync.Once

func Get_Data_Date_fromEnum2() gopurs_runtime.Value {
	once_Data_Date_fromEnum2.Do(func() {
		cache_Data_Date_fromEnum2 = Get_Data_Enum_fromEnum__1795542731()
	})
	return cache_Data_Date_fromEnum2
}

var cache_Data_Date_lessThan gopurs_runtime.Value
var once_Data_Date_lessThan sync.Once

func Get_Data_Date_lessThan() gopurs_runtime.Value {
	once_Data_Date_lessThan.Do(func() {
		cache_Data_Date_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Data_Date_lessThan
}

var cache_Data_Date_greaterThan1 gopurs_runtime.Value
var once_Data_Date_greaterThan1 sync.Once

func Get_Data_Date_greaterThan1() gopurs_runtime.Value {
	once_Data_Date_greaterThan1.Do(func() {
		cache_Data_Date_greaterThan1 = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Data_Date_greaterThan1
}

var cache_Data_Date_sub gopurs_runtime.Value
var once_Data_Date_sub sync.Once

func Get_Data_Date_sub() gopurs_runtime.Value {
	once_Data_Date_sub.Do(func() {
		cache_Data_Date_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Date_sub
}

var cache_Data_Date_bindFlipped gopurs_runtime.Value
var once_Data_Date_bindFlipped sync.Once

func Get_Data_Date_bindFlipped() gopurs_runtime.Value {
	once_Data_Date_bindFlipped.Do(func() {
		cache_Data_Date_bindFlipped = Get_Control_Bind_bindFlipped__3340339946()
	})
	return cache_Data_Date_bindFlipped
}

var cache_Data_Date_map1 gopurs_runtime.Value
var once_Data_Date_map1 sync.Once

func Get_Data_Date_map1() gopurs_runtime.Value {
	once_Data_Date_map1.Do(func() {
		cache_Data_Date_map1 = Get_Data_Functor_map__1167843562()
	})
	return cache_Data_Date_map1
}

var cache_Data_Date_bind gopurs_runtime.Value
var once_Data_Date_bind sync.Once

func Get_Data_Date_bind() gopurs_runtime.Value {
	once_Data_Date_bind.Do(func() {
		cache_Data_Date_bind = Get_Control_Bind_bind__2848886154()
	})
	return cache_Data_Date_bind
}

var cache_Data_Date_Date gopurs_runtime.Value
var once_Data_Date_Date sync.Once

func Get_Data_Date_Date() gopurs_runtime.Value {
	once_Data_Date_Date.Do(func() {
		cache_Data_Date_Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, value0.IntVal, uint32(value1.IntVal), value2.IntVal}))}
				})
			})
		})
	})
	return cache_Data_Date_Date
}

var cache_Data_Date_year gopurs_runtime.Value
var once_Data_Date_year sync.Once

func Get_Data_Date_year() gopurs_runtime.Value {
	once_Data_Date_year.Do(func() {
		cache_Data_Date_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Date_year(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))
		})
	})
	return cache_Data_Date_year
}

var cache_Data_Date_weekday gopurs_runtime.Value
var once_Data_Date_weekday sync.Once

func Get_Data_Date_weekday() gopurs_runtime.Value {
	once_Data_Date_weekday.Do(func() {
		cache_Data_Date_weekday = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Date_weekday(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box))), UnsafePtr: nil}
		})
	})
	return cache_Data_Date_weekday
}

var cache_Data_Date_showDate gopurs_runtime.Value
var once_Data_Date_showDate sync.Once

func Get_Data_Date_showDate() gopurs_runtime.Value {
	once_Data_Date_showDate.Do(func() {
		cache_Data_Date_showDate = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1726078761_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t12 string
			{
				var __t_tag_0 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1908470532 {
					__t12 = "January"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_1
				if uint32(__t_tag_1) == 2455627378 {
					__t12 = "February"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_2
				if uint32(__t_tag_2) == 4162469099 {
					__t12 = "March"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_3
				if uint32(__t_tag_3) == 1692989816 {
					__t12 = "April"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_4
				if uint32(__t_tag_4) == 330658827 {
					__t12 = "May"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_5
				if uint32(__t_tag_5) == 4067355978 {
					__t12 = "June"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_6
				if uint32(__t_tag_6) == 2276710548 {
					__t12 = "July"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_7
				if uint32(__t_tag_7) == 243771071 {
					__t12 = "August"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_8
				if uint32(__t_tag_8) == 215731793 {
					__t12 = "September"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_9
				if uint32(__t_tag_9) == 8639228 {
					__t12 = "October"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_10
				if uint32(__t_tag_10) == 49471444 {
					__t12 = "November"
					goto end_branch_12
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
				_ = __t_tag_11
				if uint32(__t_tag_11) == 3889233761 {
					__t12 = "December"
					goto end_branch_12
				} else {

				}
			}
			{
				__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_12:
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Date (Year ", Data_Show_ShowIntImpl((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)), ") "), __t12), " (Day "), Data_Show_ShowIntImpl((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2)), "))"))
		})})))}
	})
	return cache_Data_Date_showDate
}

var cache_Data_Date_month gopurs_runtime.Value
var once_Data_Date_month sync.Once

func Get_Data_Date_month() gopurs_runtime.Value {
	once_Data_Date_month.Do(func() {
		cache_Data_Date_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Date_month(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box))), UnsafePtr: nil}
		})
	})
	return cache_Data_Date_month
}

var cache_Data_Date_isLeapYear gopurs_runtime.Value
var once_Data_Date_isLeapYear sync.Once

func Get_Data_Date_isLeapYear() gopurs_runtime.Value {
	once_Data_Date_isLeapYear.Do(func() {
		cache_Data_Date_isLeapYear = gopurs_runtime.Func(func(y_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Date_isLeapYear(y_0_box.IntVal))
		})
	})
	return cache_Data_Date_isLeapYear
}

var cache_Data_Date_lastDayOfMonth gopurs_runtime.Value
var once_Data_Date_lastDayOfMonth sync.Once

func Get_Data_Date_lastDayOfMonth() gopurs_runtime.Value {
	once_Data_Date_lastDayOfMonth.Do(func() {
		cache_Data_Date_lastDayOfMonth = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth(y_0_box.IntVal, uint32(m_1_box.IntVal)))
		})
	})
	return cache_Data_Date_lastDayOfMonth
}

var cache_Data_Date_eqDate gopurs_runtime.Value
var once_Data_Date_eqDate sync.Once

func Get_Data_Date_eqDate() gopurs_runtime.Value {
	once_Data_Date_eqDate.Do(func() {
		cache_Data_Date_eqDate = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1957390985_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t_and_26 bool = false
			if ((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0) {

				var __t25 bool
				{
					var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_3
					if uint32(__t_tag_3) == 1908470532 {
						var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_4
						__t25 = (uint32(__t_tag_4) == 1908470532)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_5
					if uint32(__t_tag_5) == 2455627378 {
						var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_6
						__t25 = (uint32(__t_tag_6) == 2455627378)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_7
					if uint32(__t_tag_7) == 4162469099 {
						var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_8
						__t25 = (uint32(__t_tag_8) == 4162469099)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_9
					if uint32(__t_tag_9) == 1692989816 {
						var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_10
						__t25 = (uint32(__t_tag_10) == 1692989816)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_11
					if uint32(__t_tag_11) == 330658827 {
						var __t_tag_12 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_12
						__t25 = (uint32(__t_tag_12) == 330658827)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_13 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_13
					if uint32(__t_tag_13) == 4067355978 {
						var __t_tag_14 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_14
						__t25 = (uint32(__t_tag_14) == 4067355978)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_15 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_15
					if uint32(__t_tag_15) == 2276710548 {
						var __t_tag_16 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_16
						__t25 = (uint32(__t_tag_16) == 2276710548)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_17 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_17
					if uint32(__t_tag_17) == 243771071 {
						var __t_tag_18 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_18
						__t25 = (uint32(__t_tag_18) == 243771071)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_19 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_19
					if uint32(__t_tag_19) == 215731793 {
						var __t_tag_20 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_20
						__t25 = (uint32(__t_tag_20) == 215731793)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_21 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_21
					if uint32(__t_tag_21) == 8639228 {
						var __t_tag_22 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_22
						__t25 = (uint32(__t_tag_22) == 8639228)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_23 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_23
					if uint32(__t_tag_23) == 49471444 {
						var __t_tag_24 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_24
						__t25 = (uint32(__t_tag_24) == 49471444)
						goto end_branch_25
					} else {

					}
				}
				{
					var __t_tag_0 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
					_ = __t_tag_0
					var __t_and_2 bool = false
					if uint32(__t_tag_0) == 3889233761 {

						var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
						_ = __t_tag_1
						__t_and_2 = (uint32(__t_tag_1) == 3889233761)
					}
					__t25 = __t_and_2
				}
			end_branch_25:
				__t_and_26 = __t25
			}
			return gopurs_runtime.Bool((__t_and_26) && (((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2)))
		})})))}
	})
	return cache_Data_Date_eqDate
}

var cache_Data_Date_eq4 gopurs_runtime.Value
var once_Data_Date_eq4 sync.Once

func Get_Data_Date_eq4() gopurs_runtime.Value {
	once_Data_Date_eq4.Do(func() {
		cache_Data_Date_eq4 = Get_Data_Eq_eq__4220630099()
	})
	return cache_Data_Date_eq4
}

var cache_Data_Date_ordDate gopurs_runtime.Value
var once_Data_Date_ordDate sync.Once

func Get_Data_Date_ordDate() gopurs_runtime.Value {
	once_Data_Date_ordDate.Do(func() {
		cache_Data_Date_ordDate = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1957390985_3790796878(Rebox_Data_Date_3790796878_1957390985(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Date_eqDate()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			v_2_0 := uint32(gopurs_runtime.Apply2(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0)).IntVal)
			_ = v_2_0
			var __t3 gopurs_runtime.Value
			{
				if v_2_0 == 1527465420 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				if v_2_0 == 380165415 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				// TAST (Let): v1_3_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
				v1_3_1 := uint32(gopurs_runtime.Apply2(Rebox_Data_Date_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil}).IntVal)
				_ = v1_3_1
				var __t2 gopurs_runtime.Value
				{
					if v1_3_1 == 1527465420 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
						goto end_branch_2
					} else {

					}
				}
				{
					if v1_3_1 == 380165415 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = gopurs_runtime.Apply2(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2))
				}
			end_branch_2:
				__t3 = __t2
			}
		end_branch_3:
			return __t3
		})})))}
	})
	return cache_Data_Date_ordDate
}

var cache_Data_Date_enumDate gopurs_runtime.Value
var once_Data_Date_enumDate sync.Once

func Get_Data_Date_enumDate() gopurs_runtime.Value {
	once_Data_Date_enumDate.Do(func() {
		cache_Data_Date_enumDate = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3494563625_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502(Rebox_Data_Date_4177771502_758368489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_ordDate()))))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): pm_1_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
			var pm_1_0 *Constructor_Data_Maybe_Just[uint32] = Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})))
			// TAST (Let): pd_2_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
			var pd_2_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2))))
			var __t5 uint32
			{
				var __t3 bool
				{
					if pd_2_1 == nil {
						__t3 = true
						goto end_branch_3
					} else {

					}
				}
				{
					if pd_2_1 != nil {
						__t3 = false
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = func() bool { panic("Failed pattern match") }()
				}
			end_branch_3:
				if __t3 {
					var __t4 uint32
					{
						if pm_1_0 == nil {
							__t4 = 3889233761
							goto end_branch_4
						} else {

						}
					}
					{
						if pm_1_0 != nil {
							__t4 = (pm_1_0).V0
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = func() uint32 { panic("Failed pattern match") }()
					}
				end_branch_4:
					__t5 = __t4
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
			}
		end_branch_5:
			// TAST (Let): m_prime__3_2 shape=Branch(Branch(Other, Other, def=Other), def=Other) bindingType=Any
			m_prime__3_2 := __t5
			_ = m_prime__3_2
			// TAST (Let): l_4_6 shape=App(Var) bindingType=Int
			l_4_6 := Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, m_prime__3_2)
			_ = l_4_6
			var __t11 gopurs_runtime.Value
			{
				var __t8 bool
				{
					if pd_2_1 == nil {
						__t8 = true
						goto end_branch_8
					} else {

					}
				}
				{
					if pd_2_1 != nil {
						__t8 = false
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
				}
			end_branch_8:
				var __t_and_10 bool = false
				if __t8 {

					var __t9 bool
					{
						if pm_1_0 == nil {
							__t9 = true
							goto end_branch_9
						} else {

						}
					}
					{
						if pm_1_0 != nil {
							__t9 = false
							goto end_branch_9
						} else {

						}
					}
					{
						__t9 = func() bool { panic("Failed pattern match") }()
					}
				end_branch_9:
					__t_and_10 = __t9
				}
				if __t_and_10 {
					__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0))))}
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))))}
			}
		end_branch_11:
			// TAST (Let): __local_var_5_7 shape=Branch(App(Other), def=Other) bindingType=Any
			__local_var_5_7 := __t11
			_ = __local_var_5_7
			var __t17 gopurs_runtime.Value
			{
				if __local_var_5_7.Type == 9 && __local_var_5_7.IntVal == 930809136 && __local_var_5_7.UnsafePtr != nil {
					var __t16 gopurs_runtime.Value
					{
						var __t15 bool
						{
							if pd_2_1 == nil {
								__t15 = true
								goto end_branch_15
							} else {

							}
						}
						{
							if pd_2_1 != nil {
								__t15 = false
								goto end_branch_15
							} else {

							}
						}
						{
							__t15 = func() bool { panic("Failed pattern match") }()
						}
					end_branch_15:
						if __t15 {
							__t16 = func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_5_7.UnsafePtr).V0.IntVal, m_prime__3_2, l_4_6}))}, true}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()
							goto end_branch_16
						} else {

						}
					}
					{
						var __t14 struct {
							V0 gopurs_runtime.Value
							V1 bool
						}
						{
							if pd_2_1 != nil {
								__t14 = struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_5_7.UnsafePtr).V0.IntVal, m_prime__3_2, (pd_2_1).V0}))}, true}
								goto end_branch_14
							} else {

							}
						}
						{
							__t14 = struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
						}
					end_branch_14:
						__t16 = func() gopurs_runtime.Value {
							_v := __t14
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()
					}
				end_branch_16:
					__t17 = __t16
					goto end_branch_17
				} else {

				}
			}
			{
				var __t13 gopurs_runtime.Value
				{
					var __t12 bool
					{
						if pd_2_1 == nil {
							__t12 = true
							goto end_branch_12
						} else {

						}
					}
					{
						if pd_2_1 != nil {
							__t12 = false
							goto end_branch_12
						} else {

						}
					}
					{
						__t12 = func() bool { panic("Failed pattern match") }()
					}
				end_branch_12:
					if __t12 {
						__t13 = func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()
						goto end_branch_13
					} else {

					}
				}
				{
					__t13 = func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{}, false}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()
				}
			end_branch_13:
				__t17 = __t13
			}
		end_branch_17:
			return __t17
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): sm_1_18 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
			var sm_1_18 *Constructor_Data_Maybe_Just[uint32] = Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V2, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})))
			// TAST (Let): v1_2_19 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
			var v1_2_19 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2))))
			var __t60 gopurs_runtime.Value
			{
				var __t_tag_35 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_ordMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3308271157_4177771502(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_2_19))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1)), true}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))))})
				_ = __t_tag_35
				if uint32(__t_tag_35.IntVal) == 380165415 {
					var __t42 gopurs_runtime.Value
					{
						var __t39 bool
						{
							var __t_tag_37 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
							_ = __t_tag_37
							if __t_tag_37.Type == 9 && __t_tag_37.IntVal == 930809136 && __t_tag_37.UnsafePtr == nil {
								__t39 = true
								goto end_branch_39
							} else {

							}
						}
						{
							var __t_tag_38 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
							_ = __t_tag_38
							if __t_tag_38.Type == 9 && __t_tag_38.IntVal == 930809136 && __t_tag_38.UnsafePtr != nil {
								__t39 = false
								goto end_branch_39
							} else {

							}
						}
						{
							__t39 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
						}
					end_branch_39:
						var __t_and_41 bool = false
						if __t39 {

							var __t40 bool
							{
								if sm_1_18 == nil {
									__t40 = true
									goto end_branch_40
								} else {

								}
							}
							{
								if sm_1_18 != nil {
									__t40 = false
									goto end_branch_40
								} else {

								}
							}
							{
								__t40 = func() bool { panic("Failed pattern match") }()
							}
						end_branch_40:
							__t_and_41 = __t40
						}
						if __t_and_41 {
							__t42 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0))))}
							goto end_branch_42
						} else {

						}
					}
					{
						__t42 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), true}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))))}
					}
				end_branch_42:
					// TAST (Let): __local_var_3_36 shape=Branch(App(Other), def=Other) bindingType=Any
					__local_var_3_36 := __t42
					_ = __local_var_3_36
					var __t59 gopurs_runtime.Value
					{
						if __local_var_3_36.Type == 9 && __local_var_3_36.IntVal == 930809136 && __local_var_3_36.UnsafePtr != nil {
							var __t52 uint32
							{
								var __t50 bool
								{
									var __t_tag_48 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
									_ = __t_tag_48
									if __t_tag_48.Type == 9 && __t_tag_48.IntVal == 930809136 && __t_tag_48.UnsafePtr == nil {
										__t50 = true
										goto end_branch_50
									} else {

									}
								}
								{
									var __t_tag_49 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
									_ = __t_tag_49
									if __t_tag_49.Type == 9 && __t_tag_49.IntVal == 930809136 && __t_tag_49.UnsafePtr != nil {
										__t50 = false
										goto end_branch_50
									} else {

									}
								}
								{
									__t50 = func() bool { panic("Failed pattern match") }()
								}
							end_branch_50:
								if __t50 {
									var __t51 uint32
									{
										if sm_1_18 == nil {
											__t51 = 1908470532
											goto end_branch_51
										} else {

										}
									}
									{
										if sm_1_18 != nil {
											__t51 = (sm_1_18).V0
											goto end_branch_51
										} else {

										}
									}
									{
										__t51 = func() uint32 { panic("Failed pattern match") }()
									}
								end_branch_51:
									__t52 = __t51
									goto end_branch_52
								} else {

								}
							}
							{
								__t52 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
							}
						end_branch_52:
							// TAST (Let): __local_var_4_47 shape=Branch(Branch(Other, Other, def=Other), def=Other) bindingType=Any
							__local_var_4_47 := __t52
							_ = __local_var_4_47
							var __t58 gopurs_runtime.Value
							{
								var __t57 bool
								{
									var __t_tag_55 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
									_ = __t_tag_55
									if __t_tag_55.Type == 9 && __t_tag_55.IntVal == 930809136 && __t_tag_55.UnsafePtr == nil {
										__t57 = true
										goto end_branch_57
									} else {

									}
								}
								{
									var __t_tag_56 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
									_ = __t_tag_56
									if __t_tag_56.Type == 9 && __t_tag_56.IntVal == 930809136 && __t_tag_56.UnsafePtr != nil {
										__t57 = false
										goto end_branch_57
									} else {

									}
								}
								{
									__t57 = func() bool { panic("Failed pattern match") }()
								}
							end_branch_57:
								if __t57 {
									__t58 = func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 bool
										}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_36.UnsafePtr).V0.IntVal, __local_var_4_47, int64(1)}))}, true}
										if _v.V1 {
											return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
									}()
									goto end_branch_58
								} else {

								}
							}
							{
								var __t54 struct {
									V0 gopurs_runtime.Value
									V1 bool
								}
								{
									var __t_tag_53 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
									_ = __t_tag_53
									if __t_tag_53.Type == 9 && __t_tag_53.IntVal == 930809136 && __t_tag_53.UnsafePtr != nil {
										__t54 = struct {
											V0 gopurs_runtime.Value
											V1 bool
										}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_36.UnsafePtr).V0.IntVal, __local_var_4_47, ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0.IntVal}))}, true}
										goto end_branch_54
									} else {

									}
								}
								{
									__t54 = struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Value{}, false}
								}
							end_branch_54:
								__t58 = func() gopurs_runtime.Value {
									_v := __t54
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()
							}
						end_branch_58:
							__t59 = __t58
							goto end_branch_59
						} else {

						}
					}
					{
						var __t46 gopurs_runtime.Value
						{
							var __t45 bool
							{
								var __t_tag_43 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
								_ = __t_tag_43
								if __t_tag_43.Type == 9 && __t_tag_43.IntVal == 930809136 && __t_tag_43.UnsafePtr == nil {
									__t45 = true
									goto end_branch_45
								} else {

								}
							}
							{
								var __t_tag_44 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
								_ = __t_tag_44
								if __t_tag_44.Type == 9 && __t_tag_44.IntVal == 930809136 && __t_tag_44.UnsafePtr != nil {
									__t45 = false
									goto end_branch_45
								} else {

								}
							}
							{
								__t45 = func() bool { panic("Failed pattern match") }()
							}
						end_branch_45:
							if __t45 {
								__t46 = func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Value{}, false}
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()
								goto end_branch_46
							} else {

							}
						}
						{
							__t46 = func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()
						}
					end_branch_46:
						__t59 = __t46
					}
				end_branch_59:
					__t60 = __t59
					goto end_branch_60
				} else {

				}
			}
			{
				var __t24 gopurs_runtime.Value
				{
					var __t21 bool
					{
						if v1_2_19 == nil {
							__t21 = true
							goto end_branch_21
						} else {

						}
					}
					{
						if v1_2_19 != nil {
							__t21 = false
							goto end_branch_21
						} else {

						}
					}
					{
						__t21 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
					}
				end_branch_21:
					var __t_and_23 bool = false
					if __t21 {

						var __t22 bool
						{
							if sm_1_18 == nil {
								__t22 = true
								goto end_branch_22
							} else {

							}
						}
						{
							if sm_1_18 != nil {
								__t22 = false
								goto end_branch_22
							} else {

							}
						}
						{
							__t22 = func() bool { panic("Failed pattern match") }()
						}
					end_branch_22:
						__t_and_23 = __t22
					}
					if __t_and_23 {
						__t24 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0))))}
						goto end_branch_24
					} else {

					}
				}
				{
					__t24 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))))}
				}
			end_branch_24:
				// TAST (Let): __local_var_3_20 shape=Branch(App(Other), def=Other) bindingType=Any
				__local_var_3_20 := __t24
				_ = __local_var_3_20
				var __t34 gopurs_runtime.Value
				{
					if __local_var_3_20.Type == 9 && __local_var_3_20.IntVal == 930809136 && __local_var_3_20.UnsafePtr != nil {
						var __t30 uint32
						{
							var __t28 bool
							{
								if v1_2_19 == nil {
									__t28 = true
									goto end_branch_28
								} else {

								}
							}
							{
								if v1_2_19 != nil {
									__t28 = false
									goto end_branch_28
								} else {

								}
							}
							{
								__t28 = func() bool { panic("Failed pattern match") }()
							}
						end_branch_28:
							if __t28 {
								var __t29 uint32
								{
									if sm_1_18 == nil {
										__t29 = 1908470532
										goto end_branch_29
									} else {

									}
								}
								{
									if sm_1_18 != nil {
										__t29 = (sm_1_18).V0
										goto end_branch_29
									} else {

									}
								}
								{
									__t29 = func() uint32 { panic("Failed pattern match") }()
								}
							end_branch_29:
								__t30 = __t29
								goto end_branch_30
							} else {

							}
						}
						{
							__t30 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
						}
					end_branch_30:
						// TAST (Let): __local_var_4_27 shape=Branch(Branch(Other, Other, def=Other), def=Other) bindingType=Any
						__local_var_4_27 := __t30
						_ = __local_var_4_27
						var __t33 gopurs_runtime.Value
						{
							var __t32 bool
							{
								if v1_2_19 == nil {
									__t32 = true
									goto end_branch_32
								} else {

								}
							}
							{
								if v1_2_19 != nil {
									__t32 = false
									goto end_branch_32
								} else {

								}
							}
							{
								__t32 = func() bool { panic("Failed pattern match") }()
							}
						end_branch_32:
							if __t32 {
								__t33 = func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_20.UnsafePtr).V0.IntVal, __local_var_4_27, int64(1)}))}, true}
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()
								goto end_branch_33
							} else {

							}
						}
						{
							var __t31 struct {
								V0 gopurs_runtime.Value
								V1 bool
							}
							{
								if v1_2_19 != nil {
									__t31 = struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_20.UnsafePtr).V0.IntVal, __local_var_4_27, (v1_2_19).V0}))}, true}
									goto end_branch_31
								} else {

								}
							}
							{
								__t31 = struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
							}
						end_branch_31:
							__t33 = func() gopurs_runtime.Value {
								_v := __t31
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()
						}
					end_branch_33:
						__t34 = __t33
						goto end_branch_34
					} else {

					}
				}
				{
					var __t26 gopurs_runtime.Value
					{
						var __t25 bool
						{
							if v1_2_19 == nil {
								__t25 = true
								goto end_branch_25
							} else {

							}
						}
						{
							if v1_2_19 != nil {
								__t25 = false
								goto end_branch_25
							} else {

							}
						}
						{
							__t25 = func() bool { panic("Failed pattern match") }()
						}
					end_branch_25:
						if __t25 {
							__t26 = func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()
							goto end_branch_26
						} else {

						}
					}
					{
						__t26 = func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()
					}
				end_branch_26:
					__t34 = __t26
				}
			end_branch_34:
				__t60 = __t34
			}
		end_branch_60:
			return __t60
		})})))}
	})
	return cache_Data_Date_enumDate
}

var cache_Data_Date_pred3 gopurs_runtime.Value
var once_Data_Date_pred3 sync.Once

func Get_Data_Date_pred3() gopurs_runtime.Value {
	once_Data_Date_pred3.Do(func() {
		cache_Data_Date_pred3 = Rebox_Data_Date_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate())).V1
	})
	return cache_Data_Date_pred3
}

var cache_Data_Date_succ3 gopurs_runtime.Value
var once_Data_Date_succ3 sync.Once

func Get_Data_Date_succ3() gopurs_runtime.Value {
	once_Data_Date_succ3.Do(func() {
		cache_Data_Date_succ3 = Get_Data_Enum_succ__545117358()
	})
	return cache_Data_Date_succ3
}

var cache_Data_Date_diff gopurs_runtime.Value
var once_Data_Date_diff sync.Once

func Get_Data_Date_diff() gopurs_runtime.Value {
	once_Data_Date_diff.Do(func() {
		cache_Data_Date_diff = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Date_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box))
		})
	})
	return cache_Data_Date_diff
}

var cache_Data_Date_day gopurs_runtime.Value
var once_Data_Date_day sync.Once

func Get_Data_Date_day() gopurs_runtime.Value {
	once_Data_Date_day.Do(func() {
		cache_Data_Date_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Date_day(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))
		})
	})
	return cache_Data_Date_day
}

var cache_Data_Date_canonicalDate gopurs_runtime.Value
var once_Data_Date_canonicalDate sync.Once

func Get_Data_Date_canonicalDate() gopurs_runtime.Value {
	once_Data_Date_canonicalDate.Do(func() {
		cache_Data_Date_canonicalDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Date_canonicalDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal))}
		})
	})
	return cache_Data_Date_canonicalDate
}

var cache_Data_Date_exactDate gopurs_runtime.Value
var once_Data_Date_exactDate sync.Once

func Get_Data_Date_exactDate() gopurs_runtime.Value {
	once_Data_Date_exactDate.Do(func() {
		cache_Data_Date_exactDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Date_exactDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Date_exactDate
}

var cache_Data_Date_boundedDate gopurs_runtime.Value
var once_Data_Date_boundedDate sync.Once

func Get_Data_Date_boundedDate() gopurs_runtime.Value {
	once_Data_Date_boundedDate.Do(func() {
		cache_Data_Date_boundedDate = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3523628265_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502(Rebox_Data_Date_4177771502_758368489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_ordDate()))))}
		}), (&Constructor_Data_Date_Date{1, int64(-271820), 1908470532, int64(1)}), (&Constructor_Data_Date_Date{1, int64(275759), 3889233761, int64(31)})})))}
	})
	return cache_Data_Date_boundedDate
}

var cache_Data_Date_adjust gopurs_runtime.Value
var once_Data_Date_adjust sync.Once

func Get_Data_Date_adjust() gopurs_runtime.Value {
	once_Data_Date_adjust.Do(func() {
		cache_Data_Date_adjust = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, date_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Date_adjust(v_0_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](date_1_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Date_adjust
}

type Constructor_Data_Date_Date struct {
	Rc uint32
	V0 int64
	V1 uint32
	V2 int64
}

func Call_Data_Date_year(v_0_loop *Constructor_Data_Date_Date) int64 {
	var v_0 *Constructor_Data_Date_Date = v_0_loop
	_ = v_0
	return (v_0).V0
}

func Call_Data_Date_weekday(v_0_loop *Constructor_Data_Date_Date) uint32 {
	var v_0 *Constructor_Data_Date_Date = v_0_loop
	_ = v_0
	var __t13 int64
	{
		var __t_tag_1 uint32 = (v_0).V1
		_ = __t_tag_1
		if uint32(__t_tag_1) == 1908470532 {
			__t13 = int64(1)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_2 uint32 = (v_0).V1
		_ = __t_tag_2
		if uint32(__t_tag_2) == 2455627378 {
			__t13 = int64(2)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_3 uint32 = (v_0).V1
		_ = __t_tag_3
		if uint32(__t_tag_3) == 4162469099 {
			__t13 = int64(3)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_4 uint32 = (v_0).V1
		_ = __t_tag_4
		if uint32(__t_tag_4) == 1692989816 {
			__t13 = int64(4)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_5 uint32 = (v_0).V1
		_ = __t_tag_5
		if uint32(__t_tag_5) == 330658827 {
			__t13 = int64(5)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_6 uint32 = (v_0).V1
		_ = __t_tag_6
		if uint32(__t_tag_6) == 4067355978 {
			__t13 = int64(6)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_7 uint32 = (v_0).V1
		_ = __t_tag_7
		if uint32(__t_tag_7) == 2276710548 {
			__t13 = int64(7)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_8 uint32 = (v_0).V1
		_ = __t_tag_8
		if uint32(__t_tag_8) == 243771071 {
			__t13 = int64(8)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_9 uint32 = (v_0).V1
		_ = __t_tag_9
		if uint32(__t_tag_9) == 215731793 {
			__t13 = int64(9)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_10 uint32 = (v_0).V1
		_ = __t_tag_10
		if uint32(__t_tag_10) == 8639228 {
			__t13 = int64(10)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_11 uint32 = (v_0).V1
		_ = __t_tag_11
		if uint32(__t_tag_11) == 49471444 {
			__t13 = int64(11)
			goto end_branch_13
		} else {

		}
	}
	{
		var __t_tag_12 uint32 = (v_0).V1
		_ = __t_tag_12
		if uint32(__t_tag_12) == 3889233761 {
			__t13 = int64(12)
			goto end_branch_13
		} else {

		}
	}
	{
		__t13 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_13:
	// TAST (Let): n_1_0 shape=UncurriedApp(Var) bindingType=Int
	n_1_0 := gopurs_runtime.UncurriedApp3(Get_Data_Date_calcWeekday(), gopurs_runtime.Int((v_0).V0), gopurs_runtime.Int(__t13), gopurs_runtime.Int((v_0).V2)).IntVal
	_ = n_1_0
	var __t15 uint32
	{
		if (n_1_0) == (int64(0)) {
			__t15 = 1326716170
			goto end_branch_15
		} else {

		}
	}
	{
		var __t14 uint32
		{
			if (n_1_0) == (int64(1)) {
				__t14 = 2900196686
				goto end_branch_14
			} else {

			}
		}
		{
			if (n_1_0) == (int64(2)) {
				__t14 = 20457557
				goto end_branch_14
			} else {

			}
		}
		{
			if (n_1_0) == (int64(3)) {
				__t14 = 4227105004
				goto end_branch_14
			} else {

			}
		}
		{
			if (n_1_0) == (int64(4)) {
				__t14 = 3818857258
				goto end_branch_14
			} else {

			}
		}
		{
			if (n_1_0) == (int64(5)) {
				__t14 = 2946274527
				goto end_branch_14
			} else {

			}
		}
		{
			if (n_1_0) == (int64(6)) {
				__t14 = 1070786179
				goto end_branch_14
			} else {

			}
		}
		{
			if (n_1_0) == (int64(7)) {
				__t14 = 1326716170
				goto end_branch_14
			} else {

			}
		}
		{
			__t14 = func() uint32 { panic("Failed pattern match") }()
		}
	end_branch_14:
		__t15 = __t14
	}
end_branch_15:
	return __t15
}

func Call_Data_Date_month(v_0_loop *Constructor_Data_Date_Date) uint32 {
	var v_0 *Constructor_Data_Date_Date = v_0_loop
	_ = v_0
	return (v_0).V1
}

func Call_Data_Date_isLeapYear(y_0_loop int64) bool {
	var y_0 int64 = y_0_loop
	_ = y_0
	return ((gopurs_runtime.IntMod(y_0, int64(4))) == (int64(0))) && (((gopurs_runtime.IntMod(y_0, int64(400))) == (int64(0))) || (((gopurs_runtime.IntMod(y_0, int64(100))) == (int64(0))) != (true)))
}

func Call_Data_Date_lastDayOfMonth(y_0_loop int64, m_1_loop uint32) int64 {
	var y_0 int64 = y_0_loop
	_ = y_0
	var m_1 uint32 = m_1_loop
	_ = m_1
	// TAST (Let): unsafeDay__3466805691_2_0 shape=App(Var) bindingType=(Func [Int] Int)
	unsafeDay__3466805691_2_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__3338668507(), Rebox_Data_Date_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())).V4)
	_ = unsafeDay__3466805691_2_0
	var __t2 int64
	{
		if m_1 == 1908470532 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 2455627378 {
			var __t1 int64
			{
				if ((gopurs_runtime.IntMod(y_0, int64(4))) == (int64(0))) && (((gopurs_runtime.IntMod(y_0, int64(400))) == (int64(0))) || (((gopurs_runtime.IntMod(y_0, int64(100))) == (int64(0))) != (true))) {
					__t1 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(29))).IntVal
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(28))).IntVal
			}
		end_branch_1:
			__t2 = __t1
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 4162469099 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 1692989816 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 330658827 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 4067355978 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 2276710548 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 243771071 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 215731793 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 8639228 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 49471444 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		if m_1 == 3889233761 {
			__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_2:
	return __t2
}

func Call_Data_Date_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
	_ = dictDuration_0
	// TAST (Let): toDuration_1_0 shape=App(Var) bindingType=(Func [Number] (TypeVar d$scope15))
	toDuration_1_0 := Call_Data_Time_Duration_toDuration(dictDuration_0)
	_ = toDuration_1_0
	return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t13 int64
		{
			var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_1
			if uint32(__t_tag_1) == 1908470532 {
				__t13 = int64(1)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_2 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_2
			if uint32(__t_tag_2) == 2455627378 {
				__t13 = int64(2)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_3
			if uint32(__t_tag_3) == 4162469099 {
				__t13 = int64(3)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_4
			if uint32(__t_tag_4) == 1692989816 {
				__t13 = int64(4)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_5
			if uint32(__t_tag_5) == 330658827 {
				__t13 = int64(5)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_6
			if uint32(__t_tag_6) == 4067355978 {
				__t13 = int64(6)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_7
			if uint32(__t_tag_7) == 2276710548 {
				__t13 = int64(7)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_8
			if uint32(__t_tag_8) == 243771071 {
				__t13 = int64(8)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_9
			if uint32(__t_tag_9) == 215731793 {
				__t13 = int64(9)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_10
			if uint32(__t_tag_10) == 8639228 {
				__t13 = int64(10)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_11
			if uint32(__t_tag_11) == 49471444 {
				__t13 = int64(11)
				goto end_branch_13
			} else {

			}
		}
		{
			var __t_tag_12 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
			_ = __t_tag_12
			if uint32(__t_tag_12) == 3889233761 {
				__t13 = int64(12)
				goto end_branch_13
			} else {

			}
		}
		{
			__t13 = func() int64 { panic("Failed pattern match") }()
		}
	end_branch_13:
		var __t26 int64
		{
			var __t_tag_14 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_14
			if uint32(__t_tag_14) == 1908470532 {
				__t26 = int64(1)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_15 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_15
			if uint32(__t_tag_15) == 2455627378 {
				__t26 = int64(2)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_16 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_16
			if uint32(__t_tag_16) == 4162469099 {
				__t26 = int64(3)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_17 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_17
			if uint32(__t_tag_17) == 1692989816 {
				__t26 = int64(4)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_18 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_18
			if uint32(__t_tag_18) == 330658827 {
				__t26 = int64(5)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_19 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_19
			if uint32(__t_tag_19) == 4067355978 {
				__t26 = int64(6)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_20 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_20
			if uint32(__t_tag_20) == 2276710548 {
				__t26 = int64(7)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_21 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_21
			if uint32(__t_tag_21) == 243771071 {
				__t26 = int64(8)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_22 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_22
			if uint32(__t_tag_22) == 215731793 {
				__t26 = int64(9)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_23 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_23
			if uint32(__t_tag_23) == 8639228 {
				__t26 = int64(10)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_24 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_24
			if uint32(__t_tag_24) == 49471444 {
				__t26 = int64(11)
				goto end_branch_26
			} else {

			}
		}
		{
			var __t_tag_25 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
			_ = __t_tag_25
			if uint32(__t_tag_25) == 3889233761 {
				__t26 = int64(12)
				goto end_branch_26
			} else {

			}
		}
		{
			__t26 = func() int64 { panic("Failed pattern match") }()
		}
	end_branch_26:
		return gopurs_runtime.Apply(toDuration_1_0, gopurs_runtime.Float(gopurs_runtime.UncurriedApp6(Get_Data_Date_calcDiff(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_2.UnsafePtr).V0), gopurs_runtime.Int(__t13), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_2.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V0), gopurs_runtime.Int(__t26), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V2)).FloatVal()))
	})
}

func Call_Data_Date_day(v_0_loop *Constructor_Data_Date_Date) int64 {
	var v_0 *Constructor_Data_Date_Date = v_0_loop
	_ = v_0
	return (v_0).V2
}

func Call_Data_Date_canonicalDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) *Constructor_Data_Date_Date {
	var y_0 int64 = y_0_loop
	_ = y_0
	var m_1 uint32 = m_1_loop
	_ = m_1
	var d_2 int64 = d_2_loop
	_ = d_2
	var __t1 int64
	{
		if m_1 == 1908470532 {
			__t1 = int64(1)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 2455627378 {
			__t1 = int64(2)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 4162469099 {
			__t1 = int64(3)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 1692989816 {
			__t1 = int64(4)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 330658827 {
			__t1 = int64(5)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 4067355978 {
			__t1 = int64(6)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 2276710548 {
			__t1 = int64(7)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 243771071 {
			__t1 = int64(8)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 215731793 {
			__t1 = int64(9)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 8639228 {
			__t1 = int64(10)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 49471444 {
			__t1 = int64(11)
			goto end_branch_1
		} else {

		}
	}
	{
		if m_1 == 3889233761 {
			__t1 = int64(12)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_1:
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.UncurriedApp4(Get_Data_Date_canonicalDateImpl(), gopurs_runtime.Func3(func(y_prime__3 gopurs_runtime.Value, m_prime__4 gopurs_runtime.Value, d_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 uint32
		{
			if (m_prime__4.IntVal) == (int64(1)) {
				__t0 = 1908470532
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(2)) {
				__t0 = 2455627378
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(3)) {
				__t0 = 4162469099
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(4)) {
				__t0 = 1692989816
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(5)) {
				__t0 = 330658827
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(6)) {
				__t0 = 4067355978
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(7)) {
				__t0 = 2276710548
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(8)) {
				__t0 = 243771071
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(9)) {
				__t0 = 215731793
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(10)) {
				__t0 = 8639228
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(11)) {
				__t0 = 49471444
				goto end_branch_0
			} else {

			}
		}
		{
			if (m_prime__4.IntVal) == (int64(12)) {
				__t0 = 3889233761
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() uint32 { panic("Failed pattern match") }()
		}
	end_branch_0:
		return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, y_prime__3.IntVal, __t0, d_prime__5.IntVal}))}
	}), gopurs_runtime.Int(y_0), gopurs_runtime.Int(__t1), gopurs_runtime.Int(d_2)))
}

func Call_Data_Date_exactDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var y_0 int64 = y_0_loop
	_ = y_0
	var m_1 uint32 = m_1_loop
	_ = m_1
	var d_2 int64 = d_2_loop
	_ = d_2
	var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]
	{
		if (gopurs_runtime.Apply2(Rebox_Data_Date_3790796878_1957390985(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Date_eqDate())).V0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Date_canonicalDate(y_0, m_1, d_2))}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, y_0, m_1, d_2}))}).IntVal) != (0) {
			__t0 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, y_0, m_1, d_2}))}, true}
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
		__t0 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Date_adjust(v_0_loop float64, date_1_loop *Constructor_Data_Date_Date) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 float64 = v_0_loop
	_ = v_0
	var date_1 *Constructor_Data_Date_Date = date_1_loop
	_ = date_1
	var Call_local_Data_Date_adj_2_0_0 func(int64, *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]
	_ = Call_local_Data_Date_adj_2_0_0
	var adj_2_0_0 gopurs_runtime.Value
	_ = adj_2_0_0
	Call_local_Data_Date_adj_2_0_0 = func(v1_3_loop int64, v2_4_loop *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	adj_2_0_0:
		for {
			if false {
				continue adj_2_0_0
			}
			var v1_3 int64 = v1_3_loop
			_ = v1_3
			var v2_4 *Constructor_Data_Date_Date = v2_4_loop
			_ = v2_4
			var __t14 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]
			{
				if (v1_3) == (int64(0)) {
					__t14 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v2_4)}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_14
				} else {

				}
			}
			{
				// TAST (Let): j_5_1 shape=Other bindingType=Int
				j_5_1 := gopurs_runtime.IntAdd(v1_3, (v2_4).V2)
				_ = j_5_1
				// TAST (Let): low_6_2 shape=Other bindingType=Boolean
				low_6_2 := (j_5_1) < (int64(1))
				_ = low_6_2
				var __t6 uint32
				{
					if low_6_2 {
						// TAST (Let): __local_var_7_4 shape=App(Other) bindingType=Any
						__local_var_7_4 := gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil})
						_ = __local_var_7_4
						var __t5 uint32
						{
							if __local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 930809136 && __local_var_7_4.UnsafePtr == nil {
								__t5 = 3889233761
								goto end_branch_5
							} else {

							}
						}
						{
							if __local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 930809136 && __local_var_7_4.UnsafePtr != nil {
								__t5 = uint32((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_7_4.UnsafePtr).V0.IntVal)
								goto end_branch_5
							} else {

							}
						}
						{
							__t5 = func() uint32 { panic("Failed pattern match") }()
						}
					end_branch_5:
						__t6 = __t5
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = (v2_4).V1
				}
			end_branch_6:
				// TAST (Let): l_7_3 shape=App(Var) bindingType=Int
				l_7_3 := Call_Data_Date_lastDayOfMonth((v2_4).V0, __t6)
				_ = l_7_3
				// TAST (Let): hi_8_7 shape=Other bindingType=Boolean
				hi_8_7 := (j_5_1) > (l_7_3)
				_ = hi_8_7
				var __t11 gopurs_runtime.Value
				{
					if low_6_2 {
						__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate())).V1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (v2_4).V0, (v2_4).V1, int64(1)}))})))))}
						goto end_branch_11
					} else {

					}
				}
				{
					if hi_8_7 {
						__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate())).V2, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (v2_4).V0, (v2_4).V1, l_7_3}))})))}
						goto end_branch_11
					} else {

					}
				}
				{
					// TAST (Let): __local_var_9_9 shape=App(Var) bindingType=Any
					__local_var_9_9 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((v2_4).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil})
					_ = __local_var_9_9
					var __t10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					{
						if ((j_5_1) >= (int64(1))) && ((j_5_1) <= (int64(31))) {
							__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Apply(__local_var_9_9, gopurs_runtime.Int(j_5_1)), true}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}())
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
				end_branch_10:
					__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
				}
			end_branch_11:
				// TAST (Let): __local_var_9_8 shape=Branch(App(Other), App(Other), def=Let(Branch(Other, def=Other))) bindingType=Any
				__local_var_9_8 := __t11
				_ = __local_var_9_8
				var __t13 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]
				{
					if __local_var_9_8.Type == 9 && __local_var_9_8.IntVal == 930809136 && __local_var_9_8.UnsafePtr != nil {
						var __t12 int64
						{
							if low_6_2 {
								__t12 = j_5_1
								goto end_branch_12
							} else {

							}
						}
						{
							if hi_8_7 {
								__t12 = gopurs_runtime.IntSub(gopurs_runtime.IntSub(j_5_1, l_7_3), int64(1))
								goto end_branch_12
							} else {

							}
						}
						{
							__t12 = int64(0)
						}
					end_branch_12:
						v1_3_loop = __t12
						v2_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_9_8.UnsafePtr).V0)
						continue adj_2_0_0
						__t13 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] { panic("unreachable") }()
						goto end_branch_13
					} else {

					}
				}
				{
					if __local_var_9_8.Type == 9 && __local_var_9_8.IntVal == 930809136 && __local_var_9_8.UnsafePtr == nil {
						__t13 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))
						goto end_branch_13
					} else {

					}
				}
				{
					__t13 = Rebox_Data_Date_3094389156_2280409795(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())
				}
			end_branch_13:
				__t14 = __t13
			}
		end_branch_14:
			return __t14
		}
	}
	adj_2_0_0 = gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(Call_local_Data_Date_adj_2_0_0(v1_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v2_4_loop_val))))}
		})
	})
	// TAST (Let): __local_var_3_15 shape=App(Var) bindingType=Any
	__local_var_3_15 := gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(v_0))
	_ = __local_var_3_15
	var __t16 gopurs_runtime.Value
	{
		if __local_var_3_15.Type == 9 && __local_var_3_15.IntVal == 930809136 && __local_var_3_15.UnsafePtr != nil {
			__t16 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(Call_local_Data_Date_adj_2_0_0((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_15.UnsafePtr).V0.IntVal, date_1)))}
			goto end_branch_16
		} else {

		}
	}
	{
		if __local_var_3_15.Type == 9 && __local_var_3_15.IntVal == 930809136 && __local_var_3_15.UnsafePtr == nil {
			__t16 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}
			goto end_branch_16
		} else {

		}
	}
	{
		__t16 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
	}
end_branch_16:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t16))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Rebox_Data_Date_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Date_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Date_1726078761_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Date_Date]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Date_1957390985_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Date_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Date_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Date_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Date_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[uint32]{}
	out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Date_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Date_3494563625_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Date_3523628265_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Date_3790796878_1957390985(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date])(unsafe.Pointer(in))
}

func Rebox_Data_Date_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Date_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Date_4177771502_758368489(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date])(unsafe.Pointer(in))
}

func Rebox_Data_Date_556578094_2359585123(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Date_556578094_3494563625(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date])(unsafe.Pointer(in))
}

func Rebox_Data_Date_556578094_4060049525(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_Enum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Date_758368489_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Get_Data_Date_calcDiff() gopurs_runtime.Value {
	return _Gopurs_Data_Date_CalcDiff
}

func Get_Data_Date_calcWeekday() gopurs_runtime.Value {
	return _Gopurs_Data_Date_CalcWeekday
}

func Get_Data_Date_canonicalDateImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Date_CanonicalDateImpl
}
