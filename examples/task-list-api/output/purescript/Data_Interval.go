package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Interval_go__append gopurs_runtime.Value
var once_Data_Interval_go__append sync.Once

func Get_Data_Interval_go__append() gopurs_runtime.Value {
	once_Data_Interval_go__append.Do(func() {
		cache_Data_Interval_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Interval_go__append
}

var cache_Data_Interval_show gopurs_runtime.Value
var once_Data_Interval_show sync.Once

func Get_Data_Interval_show() gopurs_runtime.Value {
	once_Data_Interval_show.Do(func() {
		cache_Data_Interval_show = Get_Data_Show_show__2514826429()
	})
	return cache_Data_Interval_show
}

var cache_Data_Interval_compose gopurs_runtime.Value
var once_Data_Interval_compose sync.Once

func Get_Data_Interval_compose() gopurs_runtime.Value {
	once_Data_Interval_compose.Do(func() {
		cache_Data_Interval_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Interval_compose
}

var cache_Data_Interval_compose1 gopurs_runtime.Value
var once_Data_Interval_compose1 sync.Once

func Get_Data_Interval_compose1() gopurs_runtime.Value {
	once_Data_Interval_compose1.Do(func() {
		cache_Data_Interval_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Interval_compose1
}

var cache_Data_Interval_conj gopurs_runtime.Value
var once_Data_Interval_conj sync.Once

func Get_Data_Interval_conj() gopurs_runtime.Value {
	once_Data_Interval_conj.Do(func() {
		cache_Data_Interval_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Interval_conj
}

var cache_Data_Interval_eq gopurs_runtime.Value
var once_Data_Interval_eq sync.Once

func Get_Data_Interval_eq() gopurs_runtime.Value {
	once_Data_Interval_eq.Do(func() {
		cache_Data_Interval_eq = Get_Data_Eq_eq__1199147798()
	})
	return cache_Data_Interval_eq
}

var cache_Data_Interval_compare gopurs_runtime.Value
var once_Data_Interval_compare sync.Once

func Get_Data_Interval_compare() gopurs_runtime.Value {
	once_Data_Interval_compare.Do(func() {
		cache_Data_Interval_compare = Get_Data_Ord_compare__1384306305()
	})
	return cache_Data_Interval_compare
}

var cache_Data_Interval_compose2 gopurs_runtime.Value
var once_Data_Interval_compose2 sync.Once

func Get_Data_Interval_compose2() gopurs_runtime.Value {
	once_Data_Interval_compose2.Do(func() {
		cache_Data_Interval_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Interval_compose2
}

var cache_Data_Interval_compose3 gopurs_runtime.Value
var once_Data_Interval_compose3 sync.Once

func Get_Data_Interval_compose3() gopurs_runtime.Value {
	once_Data_Interval_compose3.Do(func() {
		cache_Data_Interval_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Interval_compose3
}

var cache_Data_Interval_StartEnd gopurs_runtime.Value
var once_Data_Interval_StartEnd sync.Once

func Get_Data_Interval_StartEnd() gopurs_runtime.Value {
	once_Data_Interval_StartEnd.Do(func() {
		cache_Data_Interval_StartEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_Interval_StartEnd
}

var cache_Data_Interval_DurationEnd gopurs_runtime.Value
var once_Data_Interval_DurationEnd sync.Once

func Get_Data_Interval_DurationEnd() gopurs_runtime.Value {
	once_Data_Interval_DurationEnd.Do(func() {
		cache_Data_Interval_DurationEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_Interval_DurationEnd
}

var cache_Data_Interval_StartDuration gopurs_runtime.Value
var once_Data_Interval_StartDuration sync.Once

func Get_Data_Interval_StartDuration() gopurs_runtime.Value {
	once_Data_Interval_StartDuration.Do(func() {
		cache_Data_Interval_StartDuration = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_Interval_StartDuration
}

var cache_Data_Interval_DurationOnly gopurs_runtime.Value
var once_Data_Interval_DurationOnly sync.Once

func Get_Data_Interval_DurationOnly() gopurs_runtime.Value {
	once_Data_Interval_DurationOnly.Do(func() {
		cache_Data_Interval_DurationOnly = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Data_Interval_DurationOnly
}

var cache_Data_Interval_RecurringInterval gopurs_runtime.Value
var once_Data_Interval_RecurringInterval sync.Once

func Get_Data_Interval_RecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_RecurringInterval.Do(func() {
		cache_Data_Interval_RecurringInterval = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, Rebox_Data_Interval_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](value0)), value1}))}
			})
		})
	})
	return cache_Data_Interval_RecurringInterval
}

var cache_Data_Interval_showInterval gopurs_runtime.Value
var once_Data_Interval_showInterval sync.Once

func Get_Data_Interval_showInterval() gopurs_runtime.Value {
	once_Data_Interval_showInterval.Do(func() {
		cache_Data_Interval_showInterval = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_showInterval(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_Interval_showInterval
}

var cache_Data_Interval_showRecurringInterval gopurs_runtime.Value
var once_Data_Interval_showRecurringInterval sync.Once

func Get_Data_Interval_showRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_showRecurringInterval.Do(func() {
		cache_Data_Interval_showRecurringInterval = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_showRecurringInterval(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_Interval_showRecurringInterval
}

var cache_Data_Interval_over gopurs_runtime.Value
var once_Data_Interval_over sync.Once

func Get_Data_Interval_over() gopurs_runtime.Value {
	once_Data_Interval_over.Do(func() {
		cache_Data_Interval_over = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_over(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box))
		})
	})
	return cache_Data_Interval_over
}

var cache_Data_Interval_interval gopurs_runtime.Value
var once_Data_Interval_interval sync.Once

func Get_Data_Interval_interval() gopurs_runtime.Value {
	once_Data_Interval_interval.Do(func() {
		cache_Data_Interval_interval = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_interval(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Interval_interval
}

var cache_Data_Interval_foldableInterval gopurs_runtime.Value
var once_Data_Interval_foldableInterval sync.Once

func Get_Data_Interval_foldableInterval() gopurs_runtime.Value {
	once_Data_Interval_foldableInterval.Do(func() {
		cache_Data_Interval_foldableInterval = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Foldable_foldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableInterval()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v2_2.Type == 9 && v2_2.IntVal == 237113226 {
					__t0 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 1992629780 {
					__t0 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 2020675835 {
					__t0 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = v1_1
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Data_Foldable_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableInterval())), x_0)
		})}))}
	})
	return cache_Data_Interval_foldableInterval
}

var cache_Data_Interval_foldl gopurs_runtime.Value
var once_Data_Interval_foldl sync.Once

func Get_Data_Interval_foldl() gopurs_runtime.Value {
	once_Data_Interval_foldl.Do(func() {
		cache_Data_Interval_foldl = gopurs_runtime.RecordGet(Get_Data_Interval_foldableInterval(), "foldl")
	})
	return cache_Data_Interval_foldl
}

var cache_Data_Interval_foldr gopurs_runtime.Value
var once_Data_Interval_foldr sync.Once

func Get_Data_Interval_foldr() gopurs_runtime.Value {
	once_Data_Interval_foldr.Do(func() {
		cache_Data_Interval_foldr = gopurs_runtime.RecordGet(Get_Data_Interval_foldableInterval(), "foldr")
	})
	return cache_Data_Interval_foldr
}

var cache_Data_Interval_foldableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_foldableRecurringInterval sync.Once

func Get_Data_Interval_foldableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_foldableRecurringInterval.Do(func() {
		cache_Data_Interval_foldableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_170409538_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Foldable_foldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableRecurringInterval()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if v2_2.Type == 9 && v2_2.IntVal == 237113226 {
						__t0 = gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
						goto end_branch_0
					} else {

					}
				}
				{
					if v2_2.Type == 9 && v2_2.IntVal == 1992629780 {
						__t0 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
						goto end_branch_0
					} else {

					}
				}
				{
					if v2_2.Type == 9 && v2_2.IntVal == 2020675835 {
						__t0 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = i_1
				}
			end_branch_0:
				return __t0
			}), Get_Data_Interval_interval())
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Call_Data_Foldable_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableInterval())), f_0, i_1), Get_Data_Interval_interval())
		})})))}
	})
	return cache_Data_Interval_foldableRecurringInterval
}

var cache_Data_Interval_eqInterval gopurs_runtime.Value
var once_Data_Interval_eqInterval sync.Once

func Get_Data_Interval_eqInterval() gopurs_runtime.Value {
	once_Data_Interval_eqInterval.Do(func() {
		cache_Data_Interval_eqInterval = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_eqInterval(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Data_Interval_eqInterval
}

var cache_Data_Interval_eqRecurringInterval gopurs_runtime.Value
var once_Data_Interval_eqRecurringInterval sync.Once

func Get_Data_Interval_eqRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_eqRecurringInterval.Do(func() {
		cache_Data_Interval_eqRecurringInterval = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_eqRecurringInterval(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Data_Interval_eqRecurringInterval
}

var cache_Data_Interval_ordInterval gopurs_runtime.Value
var once_Data_Interval_ordInterval sync.Once

func Get_Data_Interval_ordInterval() gopurs_runtime.Value {
	once_Data_Interval_ordInterval.Do(func() {
		cache_Data_Interval_ordInterval = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_ordInterval(dictOrd_0_box)
		})
	})
	return cache_Data_Interval_ordInterval
}

var cache_Data_Interval_ordRecurringInterval gopurs_runtime.Value
var once_Data_Interval_ordRecurringInterval sync.Once

func Get_Data_Interval_ordRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_ordRecurringInterval.Do(func() {
		cache_Data_Interval_ordRecurringInterval = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Interval_ordRecurringInterval(dictOrd_0_box)
		})
	})
	return cache_Data_Interval_ordRecurringInterval
}

var cache_Data_Interval_bifunctorInterval gopurs_runtime.Value
var once_Data_Interval_bifunctorInterval sync.Once

func Get_Data_Interval_bifunctorInterval() gopurs_runtime.Value {
	once_Data_Interval_bifunctorInterval.Do(func() {
		cache_Data_Interval_bifunctorInterval = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v2_2.Type == 9 && v2_2.IntVal == 237113226 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 1992629780 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 2020675835 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_2.Type == 9 && v2_2.IntVal == 2281256335 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)}))}
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
	return cache_Data_Interval_bifunctorInterval
}

var cache_Data_Interval_bimap gopurs_runtime.Value
var once_Data_Interval_bimap sync.Once

func Get_Data_Interval_bimap() gopurs_runtime.Value {
	once_Data_Interval_bimap.Do(func() {
		cache_Data_Interval_bimap = gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Interval_bifunctorInterval()).V0
	})
	return cache_Data_Interval_bimap
}

var cache_Data_Interval_bifunctorRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bifunctorRecurringInterval sync.Once

func Get_Data_Interval_bifunctorRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bifunctorRecurringInterval.Do(func() {
		cache_Data_Interval_bifunctorRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2718763394_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
				_ = __t_tag_0
				if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V1)}))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
				_ = __t_tag_1
				if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V1)}))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
				_ = __t_tag_2
				if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V1)}))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
				_ = __t_tag_3
				if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0)}))}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, __t4}))}
		})})))}
	})
	return cache_Data_Interval_bifunctorRecurringInterval
}

var cache_Data_Interval_functorInterval gopurs_runtime.Value
var once_Data_Interval_functorInterval sync.Once

func Get_Data_Interval_functorInterval() gopurs_runtime.Value {
	once_Data_Interval_functorInterval.Do(func() {
		cache_Data_Interval_functorInterval = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v1_0 gopurs_runtime.Value, v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v2_1.Type == 9 && v2_1.IntVal == 237113226 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_1.Type == 9 && v2_1.IntVal == 1992629780 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_1.Type == 9 && v2_1.IntVal == 2020675835 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V1}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v2_1.Type == 9 && v2_1.IntVal == 2281256335 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0}))}
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
	return cache_Data_Interval_functorInterval
}

var cache_Data_Interval_go__map gopurs_runtime.Value
var once_Data_Interval_go__map sync.Once

func Get_Data_Interval_go__map() gopurs_runtime.Value {
	once_Data_Interval_go__map.Do(func() {
		cache_Data_Interval_go__map = gopurs_runtime.RecordGet(Get_Data_Interval_functorInterval(), "map")
	})
	return cache_Data_Interval_go__map
}

var cache_Data_Interval_extendInterval gopurs_runtime.Value
var once_Data_Interval_extendInterval sync.Once

func Get_Data_Interval_extendInterval() gopurs_runtime.Value {
	once_Data_Interval_extendInterval.Do(func() {
		cache_Data_Interval_extendInterval = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Interval_functorInterval()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v1_1.Type == 9 && v1_1.IntVal == 237113226 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v1_1.Type == 9 && v1_1.IntVal == 1992629780 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v1_1.Type == 9 && v1_1.IntVal == 2020675835 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if v1_1.Type == 9 && v1_1.IntVal == 2281256335 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0}))}
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
	return cache_Data_Interval_extendInterval
}

var cache_Data_Interval_extend gopurs_runtime.Value
var once_Data_Interval_extend sync.Once

func Get_Data_Interval_extend() gopurs_runtime.Value {
	once_Data_Interval_extend.Do(func() {
		cache_Data_Interval_extend = gopurs_runtime.RecordGet(Get_Data_Interval_extendInterval(), "extend")
	})
	return cache_Data_Interval_extend
}

var cache_Data_Interval_functorRecurringInterval gopurs_runtime.Value
var once_Data_Interval_functorRecurringInterval sync.Once

func Get_Data_Interval_functorRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_functorRecurringInterval.Do(func() {
		cache_Data_Interval_functorRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_749713730_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_0
				if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1)}))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_1
				if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1)}))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_2
				if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1}))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_3
				if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0}))}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, __t4}))}
		})})))}
	})
	return cache_Data_Interval_functorRecurringInterval
}

var cache_Data_Interval_extendRecurringInterval gopurs_runtime.Value
var once_Data_Interval_extendRecurringInterval sync.Once

func Get_Data_Interval_extendRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_extendRecurringInterval.Do(func() {
		cache_Data_Interval_extendRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_844300469_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Interval_functorRecurringInterval()
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(TypeVar b$scope124)
			__local_var_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))})
			_ = __local_var_2_0
			var __t5 gopurs_runtime.Value
			{
				var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_1
				if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 237113226 {
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_0, __local_var_2_0}))}
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_2
				if __t_tag_2.Type == 9 && __t_tag_2.IntVal == 1992629780 {
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0, __local_var_2_0}))}
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_3
				if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 2020675835 {
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1}))}
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
				_ = __t_tag_4
				if __t_tag_4.Type == 9 && __t_tag_4.IntVal == 2281256335 {
					__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0}))}
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_5:
			return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, __t5}))}
		})})))}
	})
	return cache_Data_Interval_extendRecurringInterval
}

var cache_Data_Interval_traversableInterval gopurs_runtime.Value
var once_Data_Interval_traversableInterval sync.Once

func Get_Data_Interval_traversableInterval() gopurs_runtime.Value {
	once_Data_Interval_traversableInterval.Do(func() {
		cache_Data_Interval_traversableInterval = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Interval_foldableInterval()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Interval_functorInterval()
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Interval_traversableInterval()), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope10)])
			Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_1_0
			return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if v1_3.Type == 9 && v1_3.IntVal == 237113226 {
						__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)), gopurs_runtime.Apply(v_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))
						goto end_branch_2
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 1992629780 {
						__t2 = gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply(v_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))
						goto end_branch_2
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 2020675835 {
						// TAST (Let): __local_var_4_1 shape=Other bindingType=Any
						__local_var_4_1 := (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
						_ = __local_var_4_1
						__t2 = gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_5, __local_var_4_1}))}
						}), gopurs_runtime.Apply(v_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
						goto end_branch_2
					} else {

					}
				}
				{
					if v1_3.Type == 9 && v1_3.IntVal == 2281256335 {
						__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0}))})
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			})
		})}))}
	})
	return cache_Data_Interval_traversableInterval
}

var cache_Data_Interval_traverse gopurs_runtime.Value
var once_Data_Interval_traverse sync.Once

func Get_Data_Interval_traverse() gopurs_runtime.Value {
	once_Data_Interval_traverse.Do(func() {
		cache_Data_Interval_traverse = gopurs_runtime.RecordGet(Get_Data_Interval_traversableInterval(), "traverse")
	})
	return cache_Data_Interval_traverse
}

var cache_Data_Interval_traversableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_traversableRecurringInterval sync.Once

func Get_Data_Interval_traversableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_traversableRecurringInterval.Do(func() {
		cache_Data_Interval_traversableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_185619522_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Interval_foldableRecurringInterval()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Interval_functorRecurringInterval()
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Interval_traversableRecurringInterval()), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): over1_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Interval","Interval"] [(TypeVar d$scope44), (TypeVar a$scope51)])] (TypeApp (TypeVar m$scope53) [(ADT ["Data","Interval","Interval"] [(TypeVar d$scope44), (TypeVar b$scope52)])])), (ADT ["Data","Interval","RecurringInterval"] [(TypeVar d$scope44), (TypeVar a$scope51)])] (TypeApp (TypeVar m$scope53) [(ADT ["Data","Interval","RecurringInterval"] [(TypeVar d$scope44), (TypeVar b$scope52)])]))
			over1_1_0 := gopurs_runtime.Apply(Get_Data_Interval_over(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))})
			_ = over1_1_0
			// TAST (Let): traverse1_2_1 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope51)] (TypeApp (TypeVar m$scope53) [(TypeVar b$scope52)])), (ADT ["Data","Interval","Interval"] [(TypeVar d$scope44), (TypeVar a$scope51)])] (TypeApp (TypeVar m$scope53) [(ADT ["Data","Interval","Interval"] [(TypeVar d$scope44), (TypeVar b$scope52)])]))
			traverse1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Interval_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))})
			_ = traverse1_2_1
			return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(over1_1_0, gopurs_runtime.Apply(traverse1_2_1, f_3), gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](i_4))})
			})
		})})))}
	})
	return cache_Data_Interval_traversableRecurringInterval
}

var cache_Data_Interval_bifoldableInterval gopurs_runtime.Value
var once_Data_Interval_bifoldableInterval sync.Once

func Get_Data_Interval_bifoldableInterval() gopurs_runtime.Value {
	once_Data_Interval_bifoldableInterval.Do(func() {
		cache_Data_Interval_bifoldableInterval = gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifoldable_bifoldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
		}), gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v3_3.Type == 9 && v3_3.IntVal == 237113226 {
					__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
					goto end_branch_0
				} else {

				}
			}
			{
				if v3_3.Type == 9 && v3_3.IntVal == 1992629780 {
					__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
					goto end_branch_0
				} else {

				}
			}
			{
				if v3_3.Type == 9 && v3_3.IntVal == 2020675835 {
					__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				if v3_3.Type == 9 && v3_3.IntVal == 2281256335 {
					__t0 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Data_Bifoldable_bifoldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval())), x_0)
		})}))}
	})
	return cache_Data_Interval_bifoldableInterval
}

var cache_Data_Interval_bifoldl gopurs_runtime.Value
var once_Data_Interval_bifoldl sync.Once

func Get_Data_Interval_bifoldl() gopurs_runtime.Value {
	once_Data_Interval_bifoldl.Do(func() {
		cache_Data_Interval_bifoldl = gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()).V1
	})
	return cache_Data_Interval_bifoldl
}

var cache_Data_Interval_bifoldr gopurs_runtime.Value
var once_Data_Interval_bifoldr sync.Once

func Get_Data_Interval_bifoldr() gopurs_runtime.Value {
	once_Data_Interval_bifoldr.Do(func() {
		cache_Data_Interval_bifoldr = gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()).V2
	})
	return cache_Data_Interval_bifoldr
}

var cache_Data_Interval_bifoldableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bifoldableRecurringInterval sync.Once

func Get_Data_Interval_bifoldableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bifoldableRecurringInterval.Do(func() {
		cache_Data_Interval_bifoldableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1868148450_3566843086((&Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifoldable_bifoldMapDefaultL(Rebox_Data_Interval_1868148450_3566843086(Rebox_Data_Interval_3566843086_1868148450(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableRecurringInterval()))), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if v3_3.Type == 9 && v3_3.IntVal == 237113226 {
						__t0 = gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(g_1, i_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
						goto end_branch_0
					} else {

					}
				}
				{
					if v3_3.Type == 9 && v3_3.IntVal == 1992629780 {
						__t0 = gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, i_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
						goto end_branch_0
					} else {

					}
				}
				{
					if v3_3.Type == 9 && v3_3.IntVal == 2020675835 {
						__t0 = gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, i_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					if v3_3.Type == 9 && v3_3.IntVal == 2281256335 {
						__t0 = gopurs_runtime.Apply2(f_0, i_2, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_0:
				return __t0
			}), Get_Data_Interval_interval())
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply3(Call_Data_Bifoldable_bifoldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval())), f_0, g_1, i_2), Get_Data_Interval_interval())
		})})))}
	})
	return cache_Data_Interval_bifoldableRecurringInterval
}

var cache_Data_Interval_bitraversableInterval gopurs_runtime.Value
var once_Data_Interval_bitraversableInterval sync.Once

func Get_Data_Interval_bitraversableInterval() gopurs_runtime.Value {
	once_Data_Interval_bitraversableInterval.Do(func() {
		cache_Data_Interval_bitraversableInterval = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Interval_bifunctorInterval()))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bisequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Interval_bitraversableInterval()), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f$scope179)])
			Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_1_0
			// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope179)])
			Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_2_1
			return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if v2_5.Type == 9 && v2_5.IntVal == 237113226 {
						__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
						goto end_branch_2
					} else {

					}
				}
				{
					if v2_5.Type == 9 && v2_5.IntVal == 1992629780 {
						__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
						goto end_branch_2
					} else {

					}
				}
				{
					if v2_5.Type == 9 && v2_5.IntVal == 2020675835 {
						__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
						goto end_branch_2
					} else {

					}
				}
				{
					if v2_5.Type == 9 && v2_5.IntVal == 2281256335 {
						__t2 = gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			})
		})}))}
	})
	return cache_Data_Interval_bitraversableInterval
}

var cache_Data_Interval_bitraverse gopurs_runtime.Value
var once_Data_Interval_bitraverse sync.Once

func Get_Data_Interval_bitraverse() gopurs_runtime.Value {
	once_Data_Interval_bitraverse.Do(func() {
		cache_Data_Interval_bitraverse = gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Interval_bitraversableInterval()).V3
	})
	return cache_Data_Interval_bitraverse
}

var cache_Data_Interval_bitraversableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bitraversableRecurringInterval sync.Once

func Get_Data_Interval_bitraversableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bitraversableRecurringInterval.Do(func() {
		cache_Data_Interval_bitraversableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1937664898_3561684974((&Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1868148450_3566843086(Rebox_Data_Interval_3566843086_1868148450(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableRecurringInterval()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2718763394_1688994542(Rebox_Data_Interval_1688994542_2718763394(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Interval_bifunctorRecurringInterval()))))}
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bitraversable_bisequenceDefault(Rebox_Data_Interval_1937664898_3561684974(Rebox_Data_Interval_3561684974_1937664898(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Interval_bitraversableRecurringInterval()))), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): over1_1_0 shape=App(Var) bindingType=(Func [(Func [(ADT ["Data","Interval","Interval"] [(TypeVar a$scope163), (TypeVar b$scope164)])] (TypeApp (TypeVar f$scope162) [(ADT ["Data","Interval","Interval"] [(TypeVar c$scope165), (TypeVar d$scope166)])])), (ADT ["Data","Interval","RecurringInterval"] [(TypeVar a$scope163), (TypeVar b$scope164)])] (TypeApp (TypeVar f$scope162) [(ADT ["Data","Interval","RecurringInterval"] [(TypeVar c$scope165), (TypeVar d$scope166)])]))
			over1_1_0 := gopurs_runtime.Apply(Get_Data_Interval_over(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))})
			_ = over1_1_0
			// TAST (Let): bitraverse1_2_1 shape=App(Other) bindingType=(Func [(Func [(TypeVar a$scope163)] (TypeApp (TypeVar f$scope162) [(TypeVar c$scope165)])), (Func [(TypeVar b$scope164)] (TypeApp (TypeVar f$scope162) [(TypeVar d$scope166)])), (ADT ["Data","Interval","Interval"] [(TypeVar a$scope163), (TypeVar b$scope164)])] (TypeApp (TypeVar f$scope162) [(ADT ["Data","Interval","Interval"] [(TypeVar c$scope165), (TypeVar d$scope166)])]))
			bitraverse1_2_1 := gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Interval_bitraversableInterval()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))})
			_ = bitraverse1_2_1
			return gopurs_runtime.Func3(func(l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(over1_1_0, gopurs_runtime.Apply2(bitraverse1_2_1, l_3, r_4), gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](i_5))})
			})
		})})))}
	})
	return cache_Data_Interval_bitraversableRecurringInterval
}

type Constructor_Data_Interval_StartEnd[T_d any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_a
}

type Constructor_Data_Interval_DurationEnd[T_d any, T_a any] struct {
	Rc uint32
	V0 T_d
	V1 T_a
}

type Constructor_Data_Interval_StartDuration[T_d any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_d
}

type Constructor_Data_Interval_DurationOnly[T_d any, T_a any] struct {
	Rc uint32
	V0 T_d
}

type Constructor_Data_Interval_RecurringInterval[T_d any, T_a any] struct {
	Rc uint32
	V0 *Constructor_Data_Maybe_Just[int64]
	V1 gopurs_runtime.Value
}

func Call_Data_Interval_showInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if v_2.Type == 9 && v_2.IntVal == 237113226 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(StartEnd ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_2.Type == 9 && v_2.IntVal == 1992629780 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(DurationEnd ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_2.Type == 9 && v_2.IntVal == 2020675835 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(StartDuration ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			if v_2.Type == 9 && v_2.IntVal == 2281256335 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(DurationOnly ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), ")"))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
		}
	end_branch_0:
		return __t0
	})}))}
}

func Call_Data_Interval_showRecurringInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_951668290_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(RecurringInterval ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Maybe_showMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1636311157_1386611502(Rebox_Data_Interval_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), "show"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0))}).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Interval_showInterval(dictShow_0, dictShow1_1), "show"), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
	})})))}
}

func Call_Data_Interval_over(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var v_2 *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply(Get_Data_Interval_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((v_2).V0))}), gopurs_runtime.Apply(f_1, (v_2).V1))
}

func Call_Data_Interval_interval(v_0_loop *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0).V1
}

func Call_Data_Interval_eqInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 bool
		{
			if x_2.Type == 9 && x_2.IntVal == 237113226 {
				__t0 = (y_3.Type == 9 && y_3.IntVal == 237113226) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
				goto end_branch_0
			} else {

			}
		}
		{
			if x_2.Type == 9 && x_2.IntVal == 1992629780 {
				__t0 = (y_3.Type == 9 && y_3.IntVal == 1992629780) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
				goto end_branch_0
			} else {

			}
		}
		{
			if x_2.Type == 9 && x_2.IntVal == 2020675835 {
				__t0 = (y_3.Type == 9 && y_3.IntVal == 2020675835) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (x_2.Type == 9 && x_2.IntVal == 2281256335) && ((y_3.Type == 9 && y_3.IntVal == 2281256335) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)))
		}
	end_branch_0:
		return gopurs_runtime.Bool(__t0)
	})}))}
}

func Call_Data_Interval_eqRecurringInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1142559074_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1053099733_3790796878(Rebox_Data_Interval_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0))}).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Interval_eqInterval(dictEq_0, dictEq1_1), "eq"), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
	})})))}
}

func Call_Data_Interval_ordInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqInterval1_1_0 shape=App(Var) bindingType=Any
	eqInterval1_1_0 := gopurs_runtime.Apply(Get_Data_Interval_eqInterval(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqInterval1_1_0
	return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqInterval2_3_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","Interval"] [(TypeVar d$scope63), (TypeVar a$scope64)])])
		eqInterval2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqInterval1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
		_ = eqInterval2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqInterval2_3_1)}
		}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t11 gopurs_runtime.Value
			{
				if x_4.Type == 9 && x_4.IntVal == 237113226 {
					var __t4 gopurs_runtime.Value
					{
						if y_5.Type == 9 && y_5.IntVal == 237113226 {
							// TAST (Let): v_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
							v_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
							_ = v_6_2
							var __t3 gopurs_runtime.Value
							{
								if v_6_2 == 1527465420 {
									__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
									goto end_branch_3
								} else {

								}
							}
							{
								if v_6_2 == 380165415 {
									__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
									goto end_branch_3
								} else {

								}
							}
							{
								__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1)
							}
						end_branch_3:
							__t4 = __t3
							goto end_branch_4
						} else {

						}
					}
					{
						__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_4:
					__t11 = __t4
					goto end_branch_11
				} else {

				}
			}
			{
				if y_5.Type == 9 && y_5.IntVal == 237113226 {
					__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_11
				} else {

				}
			}
			{
				if x_4.Type == 9 && x_4.IntVal == 1992629780 {
					var __t7 gopurs_runtime.Value
					{
						if y_5.Type == 9 && y_5.IntVal == 1992629780 {
							// TAST (Let): v_6_5 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
							v_6_5 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
							_ = v_6_5
							var __t6 gopurs_runtime.Value
							{
								if v_6_5 == 1527465420 {
									__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
									goto end_branch_6
								} else {

								}
							}
							{
								if v_6_5 == 380165415 {
									__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
									goto end_branch_6
								} else {

								}
							}
							{
								__t6 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1)
							}
						end_branch_6:
							__t7 = __t6
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_7:
					__t11 = __t7
					goto end_branch_11
				} else {

				}
			}
			{
				if y_5.Type == 9 && y_5.IntVal == 1992629780 {
					__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_11
				} else {

				}
			}
			{
				if x_4.Type == 9 && x_4.IntVal == 2020675835 {
					var __t10 gopurs_runtime.Value
					{
						if y_5.Type == 9 && y_5.IntVal == 2020675835 {
							// TAST (Let): v_6_8 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
							v_6_8 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
							_ = v_6_8
							var __t9 gopurs_runtime.Value
							{
								if v_6_8 == 1527465420 {
									__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
									goto end_branch_9
								} else {

								}
							}
							{
								if v_6_8 == 380165415 {
									__t9 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
									goto end_branch_9
								} else {

								}
							}
							{
								__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1)
							}
						end_branch_9:
							__t10 = __t9
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_10:
					__t11 = __t10
					goto end_branch_11
				} else {

				}
			}
			{
				if y_5.Type == 9 && y_5.IntVal == 2020675835 {
					__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_11
				} else {

				}
			}
			{
				if (x_4.Type == 9 && x_4.IntVal == 2281256335) && (y_5.Type == 9 && y_5.IntVal == 2281256335) {
					__t11 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_11:
			return __t11
		})}))}
	})
}

func Call_Data_Interval_ordRecurringInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): ordInterval1_1_0 shape=App(Var) bindingType=Any
	ordInterval1_1_0 := Call_Data_Interval_ordInterval(dictOrd_0)
	_ = ordInterval1_1_0
	// TAST (Let): eqRecurringInterval1_2_1 shape=App(Var) bindingType=Any
	eqRecurringInterval1_2_1 := gopurs_runtime.Apply(Get_Data_Interval_eqRecurringInterval(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqRecurringInterval1_2_1
	return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqRecurringInterval2_4_2 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","RecurringInterval"] [(TypeVar d$scope58), (TypeVar a$scope59)])])
		eqRecurringInterval2_4_2 := Rebox_Data_Interval_3790796878_1142559074(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqRecurringInterval1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))))
		_ = eqRecurringInterval2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2740699138_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1142559074_3790796878(eqRecurringInterval2_4_2))}
		}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_7_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			v_7_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_ordMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_3308271157_4177771502(Rebox_Data_Interval_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0))}).IntVal)
			_ = v_7_3
			var __t4 gopurs_runtime.Value
			{
				if v_7_3 == 1527465420 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_4
				} else {

				}
			}
			{
				if v_7_3 == 380165415 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordInterval1_1_0, dictOrd1_3), "compare"), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1)
			}
		end_branch_4:
			return __t4
		})})))}
	})
}

func Rebox_Data_Interval_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_1142559074_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Interval_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_1688994542_2718763394(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_170409538_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_185619522_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_1868148450_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_1937664898_3561684974(in *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_2718763394_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_2740699138_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Interval_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_3561684974_1937664898(in *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_3566843086_1868148450(in *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_3790796878_1142559074(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_749713730_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_844300469_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Interval_951668290_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
