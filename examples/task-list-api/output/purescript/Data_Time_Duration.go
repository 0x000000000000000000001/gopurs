package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Duration_go__append gopurs_runtime.Value
var once_Data_Time_Duration_go__append sync.Once

func Get_Data_Time_Duration_go__append() gopurs_runtime.Value {
	once_Data_Time_Duration_go__append.Do(func() {
		cache_Data_Time_Duration_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Time_Duration_go__append
}

var cache_Data_Time_Duration_show gopurs_runtime.Value
var once_Data_Time_Duration_show sync.Once

func Get_Data_Time_Duration_show() gopurs_runtime.Value {
	once_Data_Time_Duration_show.Do(func() {
		cache_Data_Time_Duration_show = Get_Data_Show_show__14502932()
	})
	return cache_Data_Time_Duration_show
}

var cache_Data_Time_Duration_add gopurs_runtime.Value
var once_Data_Time_Duration_add sync.Once

func Get_Data_Time_Duration_add() gopurs_runtime.Value {
	once_Data_Time_Duration_add.Do(func() {
		cache_Data_Time_Duration_add = Get_Data_Semiring_add__2603109278()
	})
	return cache_Data_Time_Duration_add
}

var cache_Data_Time_Duration_compose gopurs_runtime.Value
var once_Data_Time_Duration_compose sync.Once

func Get_Data_Time_Duration_compose() gopurs_runtime.Value {
	once_Data_Time_Duration_compose.Do(func() {
		cache_Data_Time_Duration_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Time_Duration_compose
}

var cache_Data_Time_Duration_over gopurs_runtime.Value
var once_Data_Time_Duration_over sync.Once

func Get_Data_Time_Duration_over() gopurs_runtime.Value {
	once_Data_Time_Duration_over.Do(func() {
		cache_Data_Time_Duration_over = Get_Data_Newtype_over__3511533087()
	})
	return cache_Data_Time_Duration_over
}

var cache_Data_Time_Duration_negate gopurs_runtime.Value
var once_Data_Time_Duration_negate sync.Once

func Get_Data_Time_Duration_negate() gopurs_runtime.Value {
	once_Data_Time_Duration_negate.Do(func() {
		cache_Data_Time_Duration_negate = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_negate(__local_var_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_negate
}

var cache_Data_Time_Duration_Milliseconds gopurs_runtime.Value
var once_Data_Time_Duration_Milliseconds sync.Once

func Get_Data_Time_Duration_Milliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_Milliseconds.Do(func() {
		cache_Data_Time_Duration_Milliseconds = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_Milliseconds(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_Milliseconds
}

var cache_Data_Time_Duration_Duration_dollar_Dict gopurs_runtime.Value
var once_Data_Time_Duration_Duration_dollar_Dict sync.Once

func Get_Data_Time_Duration_Duration_dollar_Dict() gopurs_runtime.Value {
	once_Data_Time_Duration_Duration_dollar_Dict.Do(func() {
		cache_Data_Time_Duration_Duration_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Call_Data_Time_Duration_Duration_dollar_Dict(func() struct {
				fromDuration gopurs_runtime.Value
				toDuration   gopurs_runtime.Value
				/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					fromDuration gopurs_runtime.Value
					toDuration   gopurs_runtime.Value
					/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
				}{}
				clone.fromDuration = gopurs_runtime.RecordGet(orig, "fromDuration")
				clone.toDuration = gopurs_runtime.RecordGet(orig, "toDuration")
				return clone
			}()))}
		})
	})
	return cache_Data_Time_Duration_Duration_dollar_Dict
}

var cache_Data_Time_Duration_fromDuration gopurs_runtime.Value
var once_Data_Time_Duration_fromDuration sync.Once

func Get_Data_Time_Duration_fromDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_fromDuration.Do(func() {
		cache_Data_Time_Duration_fromDuration = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_fromDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Time_Duration_fromDuration
}

var cache_Data_Time_Duration_identity gopurs_runtime.Value
var once_Data_Time_Duration_identity sync.Once

func Get_Data_Time_Duration_identity() gopurs_runtime.Value {
	once_Data_Time_Duration_identity.Do(func() {
		cache_Data_Time_Duration_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_identity(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_identity
}

var cache_Data_Time_Duration_Duration_dollar_Dict__4143256941 gopurs_runtime.Value
var once_Data_Time_Duration_Duration_dollar_Dict__4143256941 sync.Once

func Get_Data_Time_Duration_Duration_dollar_Dict__4143256941() gopurs_runtime.Value {
	once_Data_Time_Duration_Duration_dollar_Dict__4143256941.Do(func() {
		cache_Data_Time_Duration_Duration_dollar_Dict__4143256941 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3362292604_1457896548(Call_Data_Time_Duration_Duration_dollar_Dict__4143256941(func() struct {
				fromDuration gopurs_runtime.Value
				toDuration   gopurs_runtime.Value
				/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					fromDuration gopurs_runtime.Value
					toDuration   gopurs_runtime.Value
					/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
				}{}
				clone.fromDuration = gopurs_runtime.RecordGet(orig, "fromDuration")
				clone.toDuration = gopurs_runtime.RecordGet(orig, "toDuration")
				return clone
			}())))}
		})
	})
	return cache_Data_Time_Duration_Duration_dollar_Dict__4143256941
}

var cache_Data_Time_Duration_durationMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_durationMilliseconds sync.Once

func Get_Data_Time_Duration_durationMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_durationMilliseconds.Do(func() {
		cache_Data_Time_Duration_durationMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3362292604_1457896548((&Constructor_Data_Time_Duration_Duration[float64]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})})))}
	})
	return cache_Data_Time_Duration_durationMilliseconds
}

var cache_Data_Time_Duration_compose1__1079398611 gopurs_runtime.Value
var once_Data_Time_Duration_compose1__1079398611 sync.Once

func Get_Data_Time_Duration_compose1__1079398611() gopurs_runtime.Value {
	once_Data_Time_Duration_compose1__1079398611.Do(func() {
		cache_Data_Time_Duration_compose1__1079398611 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_compose1__1079398611(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Time_Duration_compose1__1079398611
}

var cache_Data_Time_Duration_toDuration gopurs_runtime.Value
var once_Data_Time_Duration_toDuration sync.Once

func Get_Data_Time_Duration_toDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_toDuration.Do(func() {
		cache_Data_Time_Duration_toDuration = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_toDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Time_Duration_toDuration
}

var cache_Data_Time_Duration_compose__3290897600 gopurs_runtime.Value
var once_Data_Time_Duration_compose__3290897600 sync.Once

func Get_Data_Time_Duration_compose__3290897600() gopurs_runtime.Value {
	once_Data_Time_Duration_compose__3290897600.Do(func() {
		cache_Data_Time_Duration_compose__3290897600 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_compose__3290897600(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Time_Duration_compose__3290897600
}

var cache_Data_Time_Duration_compose__1298651147 gopurs_runtime.Value
var once_Data_Time_Duration_compose__1298651147 sync.Once

func Get_Data_Time_Duration_compose__1298651147() gopurs_runtime.Value {
	once_Data_Time_Duration_compose__1298651147.Do(func() {
		cache_Data_Time_Duration_compose__1298651147 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_compose__1298651147(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Time_Duration_compose__1298651147
}

var cache_Data_Time_Duration_compose1 gopurs_runtime.Value
var once_Data_Time_Duration_compose1 sync.Once

func Get_Data_Time_Duration_compose1() gopurs_runtime.Value {
	once_Data_Time_Duration_compose1.Do(func() {
		cache_Data_Time_Duration_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Time_Duration_compose1
}

var cache_Data_Time_Duration_mul gopurs_runtime.Value
var once_Data_Time_Duration_mul sync.Once

func Get_Data_Time_Duration_mul() gopurs_runtime.Value {
	once_Data_Time_Duration_mul.Do(func() {
		cache_Data_Time_Duration_mul = Get_Data_Semiring_mul__2603109278()
	})
	return cache_Data_Time_Duration_mul
}

var cache_Data_Time_Duration_div gopurs_runtime.Value
var once_Data_Time_Duration_div sync.Once

func Get_Data_Time_Duration_div() gopurs_runtime.Value {
	once_Data_Time_Duration_div.Do(func() {
		cache_Data_Time_Duration_div = Get_Data_EuclideanRing_div__266314078()
	})
	return cache_Data_Time_Duration_div
}

var cache_Data_Time_Duration_compose2 gopurs_runtime.Value
var once_Data_Time_Duration_compose2 sync.Once

func Get_Data_Time_Duration_compose2() gopurs_runtime.Value {
	once_Data_Time_Duration_compose2.Do(func() {
		cache_Data_Time_Duration_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Time_Duration_compose2
}

var cache_Data_Time_Duration_Seconds gopurs_runtime.Value
var once_Data_Time_Duration_Seconds sync.Once

func Get_Data_Time_Duration_Seconds() gopurs_runtime.Value {
	once_Data_Time_Duration_Seconds.Do(func() {
		cache_Data_Time_Duration_Seconds = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_Seconds(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_Seconds
}

var cache_Data_Time_Duration_Minutes gopurs_runtime.Value
var once_Data_Time_Duration_Minutes sync.Once

func Get_Data_Time_Duration_Minutes() gopurs_runtime.Value {
	once_Data_Time_Duration_Minutes.Do(func() {
		cache_Data_Time_Duration_Minutes = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_Minutes(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_Minutes
}

var cache_Data_Time_Duration_Hours gopurs_runtime.Value
var once_Data_Time_Duration_Hours sync.Once

func Get_Data_Time_Duration_Hours() gopurs_runtime.Value {
	once_Data_Time_Duration_Hours.Do(func() {
		cache_Data_Time_Duration_Hours = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_Hours(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_Hours
}

var cache_Data_Time_Duration_Days gopurs_runtime.Value
var once_Data_Time_Duration_Days sync.Once

func Get_Data_Time_Duration_Days() gopurs_runtime.Value {
	once_Data_Time_Duration_Days.Do(func() {
		cache_Data_Time_Duration_Days = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_Days(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_Days
}

var cache_Data_Time_Duration_showSeconds gopurs_runtime.Value
var once_Data_Time_Duration_showSeconds sync.Once

func Get_Data_Time_Duration_showSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_showSeconds.Do(func() {
		cache_Data_Time_Duration_showSeconds = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Seconds ", Data_Show_ShowNumberImpl(v_0.FloatVal())), ")"))
		})})))}
	})
	return cache_Data_Time_Duration_showSeconds
}

var cache_Data_Time_Duration_showMinutes gopurs_runtime.Value
var once_Data_Time_Duration_showMinutes sync.Once

func Get_Data_Time_Duration_showMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_showMinutes.Do(func() {
		cache_Data_Time_Duration_showMinutes = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Minutes ", Data_Show_ShowNumberImpl(v_0.FloatVal())), ")"))
		})})))}
	})
	return cache_Data_Time_Duration_showMinutes
}

var cache_Data_Time_Duration_showMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_showMilliseconds sync.Once

func Get_Data_Time_Duration_showMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_showMilliseconds.Do(func() {
		cache_Data_Time_Duration_showMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Milliseconds ", Data_Show_ShowNumberImpl(v_0.FloatVal())), ")"))
		})})))}
	})
	return cache_Data_Time_Duration_showMilliseconds
}

var cache_Data_Time_Duration_showHours gopurs_runtime.Value
var once_Data_Time_Duration_showHours sync.Once

func Get_Data_Time_Duration_showHours() gopurs_runtime.Value {
	once_Data_Time_Duration_showHours.Do(func() {
		cache_Data_Time_Duration_showHours = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Hours ", Data_Show_ShowNumberImpl(v_0.FloatVal())), ")"))
		})})))}
	})
	return cache_Data_Time_Duration_showHours
}

var cache_Data_Time_Duration_showDays gopurs_runtime.Value
var once_Data_Time_Duration_showDays sync.Once

func Get_Data_Time_Duration_showDays() gopurs_runtime.Value {
	once_Data_Time_Duration_showDays.Do(func() {
		cache_Data_Time_Duration_showDays = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3263178038_1386611502((&Constructor_Data_Show_Show[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Days ", Data_Show_ShowNumberImpl(v_0.FloatVal())), ")"))
		})})))}
	})
	return cache_Data_Time_Duration_showDays
}

var cache_Data_Time_Duration_semigroupSeconds gopurs_runtime.Value
var once_Data_Time_Duration_semigroupSeconds sync.Once

func Get_Data_Time_Duration_semigroupSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupSeconds.Do(func() {
		cache_Data_Time_Duration_semigroupSeconds = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454((&Constructor_Data_Semigroup_Semigroup[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
		})})))}
	})
	return cache_Data_Time_Duration_semigroupSeconds
}

var cache_Data_Time_Duration_semigroupMinutes gopurs_runtime.Value
var once_Data_Time_Duration_semigroupMinutes sync.Once

func Get_Data_Time_Duration_semigroupMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupMinutes.Do(func() {
		cache_Data_Time_Duration_semigroupMinutes = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454((&Constructor_Data_Semigroup_Semigroup[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
		})})))}
	})
	return cache_Data_Time_Duration_semigroupMinutes
}

var cache_Data_Time_Duration_semigroupMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_semigroupMilliseconds sync.Once

func Get_Data_Time_Duration_semigroupMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupMilliseconds.Do(func() {
		cache_Data_Time_Duration_semigroupMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454((&Constructor_Data_Semigroup_Semigroup[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
		})})))}
	})
	return cache_Data_Time_Duration_semigroupMilliseconds
}

var cache_Data_Time_Duration_semigroupHours gopurs_runtime.Value
var once_Data_Time_Duration_semigroupHours sync.Once

func Get_Data_Time_Duration_semigroupHours() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupHours.Do(func() {
		cache_Data_Time_Duration_semigroupHours = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454((&Constructor_Data_Semigroup_Semigroup[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
		})})))}
	})
	return cache_Data_Time_Duration_semigroupHours
}

var cache_Data_Time_Duration_semigroupDays gopurs_runtime.Value
var once_Data_Time_Duration_semigroupDays sync.Once

func Get_Data_Time_Duration_semigroupDays() gopurs_runtime.Value {
	once_Data_Time_Duration_semigroupDays.Do(func() {
		cache_Data_Time_Duration_semigroupDays = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454((&Constructor_Data_Semigroup_Semigroup[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
		})})))}
	})
	return cache_Data_Time_Duration_semigroupDays
}

var cache_Data_Time_Duration_ordSeconds gopurs_runtime.Value
var once_Data_Time_Duration_ordSeconds sync.Once

func Get_Data_Time_Duration_ordSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_ordSeconds.Do(func() {
		cache_Data_Time_Duration_ordSeconds = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3047586294_4177771502(Rebox_Data_Time_Duration_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
	})
	return cache_Data_Time_Duration_ordSeconds
}

var cache_Data_Time_Duration_ordMinutes gopurs_runtime.Value
var once_Data_Time_Duration_ordMinutes sync.Once

func Get_Data_Time_Duration_ordMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_ordMinutes.Do(func() {
		cache_Data_Time_Duration_ordMinutes = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3047586294_4177771502(Rebox_Data_Time_Duration_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
	})
	return cache_Data_Time_Duration_ordMinutes
}

var cache_Data_Time_Duration_ordMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_ordMilliseconds sync.Once

func Get_Data_Time_Duration_ordMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_ordMilliseconds.Do(func() {
		cache_Data_Time_Duration_ordMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3047586294_4177771502(Rebox_Data_Time_Duration_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
	})
	return cache_Data_Time_Duration_ordMilliseconds
}

var cache_Data_Time_Duration_ordHours gopurs_runtime.Value
var once_Data_Time_Duration_ordHours sync.Once

func Get_Data_Time_Duration_ordHours() gopurs_runtime.Value {
	once_Data_Time_Duration_ordHours.Do(func() {
		cache_Data_Time_Duration_ordHours = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3047586294_4177771502(Rebox_Data_Time_Duration_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
	})
	return cache_Data_Time_Duration_ordHours
}

var cache_Data_Time_Duration_ordDays gopurs_runtime.Value
var once_Data_Time_Duration_ordDays sync.Once

func Get_Data_Time_Duration_ordDays() gopurs_runtime.Value {
	once_Data_Time_Duration_ordDays.Do(func() {
		cache_Data_Time_Duration_ordDays = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3047586294_4177771502(Rebox_Data_Time_Duration_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
	})
	return cache_Data_Time_Duration_ordDays
}

var cache_Data_Time_Duration_newtypeSeconds gopurs_runtime.Value
var once_Data_Time_Duration_newtypeSeconds sync.Once

func Get_Data_Time_Duration_newtypeSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeSeconds.Do(func() {
		cache_Data_Time_Duration_newtypeSeconds = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_4289840808_385277032((&Constructor_Data_Newtype_Newtype[float64, float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Time_Duration_newtypeSeconds
}

var cache_Data_Time_Duration_newtypeMinutes gopurs_runtime.Value
var once_Data_Time_Duration_newtypeMinutes sync.Once

func Get_Data_Time_Duration_newtypeMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeMinutes.Do(func() {
		cache_Data_Time_Duration_newtypeMinutes = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_4289840808_385277032((&Constructor_Data_Newtype_Newtype[float64, float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Time_Duration_newtypeMinutes
}

var cache_Data_Time_Duration_newtypeMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_newtypeMilliseconds sync.Once

func Get_Data_Time_Duration_newtypeMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeMilliseconds.Do(func() {
		cache_Data_Time_Duration_newtypeMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_4289840808_385277032((&Constructor_Data_Newtype_Newtype[float64, float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Time_Duration_newtypeMilliseconds
}

var cache_Data_Time_Duration_newtypeHours gopurs_runtime.Value
var once_Data_Time_Duration_newtypeHours sync.Once

func Get_Data_Time_Duration_newtypeHours() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeHours.Do(func() {
		cache_Data_Time_Duration_newtypeHours = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_4289840808_385277032((&Constructor_Data_Newtype_Newtype[float64, float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Time_Duration_newtypeHours
}

var cache_Data_Time_Duration_newtypeDays gopurs_runtime.Value
var once_Data_Time_Duration_newtypeDays sync.Once

func Get_Data_Time_Duration_newtypeDays() gopurs_runtime.Value {
	once_Data_Time_Duration_newtypeDays.Do(func() {
		cache_Data_Time_Duration_newtypeDays = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_4289840808_385277032((&Constructor_Data_Newtype_Newtype[float64, float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Time_Duration_newtypeDays
}

var cache_Data_Time_Duration_monoidSeconds gopurs_runtime.Value
var once_Data_Time_Duration_monoidSeconds sync.Once

func Get_Data_Time_Duration_monoidSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidSeconds.Do(func() {
		cache_Data_Time_Duration_monoidSeconds = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454(Rebox_Data_Time_Duration_4179793454_1640832054(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Time_Duration_semigroupSeconds()))))}
		}), 0.0})))}
	})
	return cache_Data_Time_Duration_monoidSeconds
}

var cache_Data_Time_Duration_monoidMinutes gopurs_runtime.Value
var once_Data_Time_Duration_monoidMinutes sync.Once

func Get_Data_Time_Duration_monoidMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidMinutes.Do(func() {
		cache_Data_Time_Duration_monoidMinutes = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454(Rebox_Data_Time_Duration_4179793454_1640832054(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Time_Duration_semigroupMinutes()))))}
		}), 0.0})))}
	})
	return cache_Data_Time_Duration_monoidMinutes
}

var cache_Data_Time_Duration_monoidMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_monoidMilliseconds sync.Once

func Get_Data_Time_Duration_monoidMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidMilliseconds.Do(func() {
		cache_Data_Time_Duration_monoidMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454(Rebox_Data_Time_Duration_4179793454_1640832054(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Time_Duration_semigroupMilliseconds()))))}
		}), 0.0})))}
	})
	return cache_Data_Time_Duration_monoidMilliseconds
}

var cache_Data_Time_Duration_monoidHours gopurs_runtime.Value
var once_Data_Time_Duration_monoidHours sync.Once

func Get_Data_Time_Duration_monoidHours() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidHours.Do(func() {
		cache_Data_Time_Duration_monoidHours = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454(Rebox_Data_Time_Duration_4179793454_1640832054(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Time_Duration_semigroupHours()))))}
		}), 0.0})))}
	})
	return cache_Data_Time_Duration_monoidHours
}

var cache_Data_Time_Duration_monoidDays gopurs_runtime.Value
var once_Data_Time_Duration_monoidDays sync.Once

func Get_Data_Time_Duration_monoidDays() gopurs_runtime.Value {
	once_Data_Time_Duration_monoidDays.Do(func() {
		cache_Data_Time_Duration_monoidDays = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_1640832054_4179793454(Rebox_Data_Time_Duration_4179793454_1640832054(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Time_Duration_semigroupDays()))))}
		}), 0.0})))}
	})
	return cache_Data_Time_Duration_monoidDays
}

var cache_Data_Time_Duration_durationDays gopurs_runtime.Value
var once_Data_Time_Duration_durationDays sync.Once

func Get_Data_Time_Duration_durationDays() gopurs_runtime.Value {
	once_Data_Time_Duration_durationDays.Do(func() {
		cache_Data_Time_Duration_durationDays = gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3362292604_1457896548((&Constructor_Data_Time_Duration_Duration[float64]{1, gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) * (86400000.0))
		})), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) / (86400000.0))
		}))})))}
	})
	return cache_Data_Time_Duration_durationDays
}

var cache_Data_Time_Duration_fromDuration__928029503 gopurs_runtime.Value
var once_Data_Time_Duration_fromDuration__928029503 sync.Once

func Get_Data_Time_Duration_fromDuration__928029503() gopurs_runtime.Value {
	once_Data_Time_Duration_fromDuration__928029503.Do(func() {
		cache_Data_Time_Duration_fromDuration__928029503 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_fromDuration__928029503(__eta0_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_fromDuration__928029503
}

var cache_Data_Time_Duration_negateDuration gopurs_runtime.Value
var once_Data_Time_Duration_negateDuration sync.Once

func Get_Data_Time_Duration_negateDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_negateDuration.Do(func() {
		cache_Data_Time_Duration_negateDuration = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_negateDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box))
		})
	})
	return cache_Data_Time_Duration_negateDuration
}

var cache_Data_Time_Duration_negateDuration__389498750 gopurs_runtime.Value
var once_Data_Time_Duration_negateDuration__389498750 sync.Once

func Get_Data_Time_Duration_negateDuration__389498750() gopurs_runtime.Value {
	once_Data_Time_Duration_negateDuration__389498750.Do(func() {
		cache_Data_Time_Duration_negateDuration__389498750 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Time_Duration_negateDuration__389498750(__eta0_0_box.FloatVal()))
		})
	})
	return cache_Data_Time_Duration_negateDuration__389498750
}

var cache_Data_Time_Duration_eqSeconds gopurs_runtime.Value
var once_Data_Time_Duration_eqSeconds sync.Once

func Get_Data_Time_Duration_eqSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_eqSeconds.Do(func() {
		cache_Data_Time_Duration_eqSeconds = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_687527510_3790796878(Rebox_Data_Time_Duration_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
	})
	return cache_Data_Time_Duration_eqSeconds
}

var cache_Data_Time_Duration_eqMinutes gopurs_runtime.Value
var once_Data_Time_Duration_eqMinutes sync.Once

func Get_Data_Time_Duration_eqMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_eqMinutes.Do(func() {
		cache_Data_Time_Duration_eqMinutes = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_687527510_3790796878(Rebox_Data_Time_Duration_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
	})
	return cache_Data_Time_Duration_eqMinutes
}

var cache_Data_Time_Duration_eqMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_eqMilliseconds sync.Once

func Get_Data_Time_Duration_eqMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_eqMilliseconds.Do(func() {
		cache_Data_Time_Duration_eqMilliseconds = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_687527510_3790796878(Rebox_Data_Time_Duration_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
	})
	return cache_Data_Time_Duration_eqMilliseconds
}

var cache_Data_Time_Duration_eqHours gopurs_runtime.Value
var once_Data_Time_Duration_eqHours sync.Once

func Get_Data_Time_Duration_eqHours() gopurs_runtime.Value {
	once_Data_Time_Duration_eqHours.Do(func() {
		cache_Data_Time_Duration_eqHours = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_687527510_3790796878(Rebox_Data_Time_Duration_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
	})
	return cache_Data_Time_Duration_eqHours
}

var cache_Data_Time_Duration_eqDays gopurs_runtime.Value
var once_Data_Time_Duration_eqDays sync.Once

func Get_Data_Time_Duration_eqDays() gopurs_runtime.Value {
	once_Data_Time_Duration_eqDays.Do(func() {
		cache_Data_Time_Duration_eqDays = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_687527510_3790796878(Rebox_Data_Time_Duration_3790796878_687527510(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqNumber()))))}
	})
	return cache_Data_Time_Duration_eqDays
}

var cache_Data_Time_Duration_durationSeconds gopurs_runtime.Value
var once_Data_Time_Duration_durationSeconds sync.Once

func Get_Data_Time_Duration_durationSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_durationSeconds.Do(func() {
		cache_Data_Time_Duration_durationSeconds = gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3362292604_1457896548((&Constructor_Data_Time_Duration_Duration[float64]{1, gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) * (1000.0))
		})), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) / (1000.0))
		}))})))}
	})
	return cache_Data_Time_Duration_durationSeconds
}

var cache_Data_Time_Duration_durationMinutes gopurs_runtime.Value
var once_Data_Time_Duration_durationMinutes sync.Once

func Get_Data_Time_Duration_durationMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_durationMinutes.Do(func() {
		cache_Data_Time_Duration_durationMinutes = gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3362292604_1457896548((&Constructor_Data_Time_Duration_Duration[float64]{1, gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) * (60000.0))
		})), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) / (60000.0))
		}))})))}
	})
	return cache_Data_Time_Duration_durationMinutes
}

var cache_Data_Time_Duration_durationHours gopurs_runtime.Value
var once_Data_Time_Duration_durationHours sync.Once

func Get_Data_Time_Duration_durationHours() gopurs_runtime.Value {
	once_Data_Time_Duration_durationHours.Do(func() {
		cache_Data_Time_Duration_durationHours = gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Duration_3362292604_1457896548((&Constructor_Data_Time_Duration_Duration[float64]{1, gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) * (3600000.0))
		})), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) / (3600000.0))
		}))})))}
	})
	return cache_Data_Time_Duration_durationHours
}

var cache_Data_Time_Duration_convertDuration gopurs_runtime.Value
var once_Data_Time_Duration_convertDuration sync.Once

func Get_Data_Time_Duration_convertDuration() gopurs_runtime.Value {
	once_Data_Time_Duration_convertDuration.Do(func() {
		cache_Data_Time_Duration_convertDuration = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_convertDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box))
		})
	})
	return cache_Data_Time_Duration_convertDuration
}

type Constructor_Data_Time_Duration_Duration[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00540069006d0065005f004400750072006100740069006f006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00540069006d0065005f004400750072006100740069006f006e005f004400750072006100740069006f006e
func init() {
	gopurs_runtime.StructGetters[487663984] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "fromDuration":
			return gopurs_runtime.Box(c.V0)
		case "toDuration":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Time_Duration_Duration: " + key)
		}
	}
}

func Call_Data_Time_Duration_negate(__local_var_0_loop float64) float64 {
	var __local_var_0 float64 = __local_var_0_loop
	_ = __local_var_0
	return -(__local_var_0)
}

func Call_Data_Time_Duration_Milliseconds(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Duration_Duration_dollar_Dict(x_0_loop struct {
	fromDuration gopurs_runtime.Value
	toDuration   gopurs_runtime.Value
	/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
}) *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] {
	var x_0 struct {
		fromDuration gopurs_runtime.Value
		toDuration   gopurs_runtime.Value
		/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("fromDuration", "toDuration", orig.fromDuration, orig.toDuration)
	}())
}

func Call_Data_Time_Duration_fromDuration(dict_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Time_Duration_identity(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return gopurs_runtime.Float(x_0).FloatVal()
}

func Call_Data_Time_Duration_Duration_dollar_Dict__4143256941(x_0_loop struct {
	fromDuration gopurs_runtime.Value
	toDuration   gopurs_runtime.Value
	/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
}) *Constructor_Data_Time_Duration_Duration[float64] {
Duration_dollar_Dict__4143256941:
	for {
		if false {
			continue Duration_dollar_Dict__4143256941
		}
		var x_0 struct {
			fromDuration gopurs_runtime.Value
			toDuration   gopurs_runtime.Value
			/*gopurs:record-v1:2:00660072006f006d004400750072006100740069006f006e,0074006f004400750072006100740069006f006e*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("fromDuration", "toDuration", orig.fromDuration, orig.toDuration)
		}())
	}
}

func Call_Data_Time_Duration_compose1__1079398611(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose1__1079398611:
	for {
		if false {
			continue compose1__1079398611
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(-(__local_var_2.FloatVal()))
		})), Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V0)
	}
}

func Call_Data_Time_Duration_toDuration(dict_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Time_Duration_compose__3290897600(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3290897600:
	for {
		if false {
			continue compose__3290897600
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(-(__local_var_2.FloatVal()))
		})), Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V0))
	}
}

func Call_Data_Time_Duration_compose__1298651147(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1298651147:
	for {
		if false {
			continue compose__1298651147
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(-(__local_var_2.FloatVal()))
		})), Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V0))
	}
}

func Call_Data_Time_Duration_Seconds(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Duration_Minutes(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Duration_Hours(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Duration_Days(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Time_Duration_fromDuration__928029503(__eta0_0_loop float64) float64 {
fromDuration__928029503:
	for {
		if false {
			continue fromDuration__928029503
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationDays())).V0, gopurs_runtime.Float(__eta0_0)).FloatVal()
	}
}

func Call_Data_Time_Duration_negateDuration(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
	_ = dictDuration_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Time_Duration_toDuration(dictDuration_0), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(__local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(-(__local_var_1.FloatVal()))
	})), Call_Data_Time_Duration_fromDuration(dictDuration_0)))
}

func Call_Data_Time_Duration_negateDuration__389498750(__eta0_0_loop float64) float64 {
negateDuration__389498750:
	for {
		if false {
			continue negateDuration__389498750
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Call_Data_Time_Duration_compose__1298651147(Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(__local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(-(__local_var_1.FloatVal()))
		})), Rebox_Data_Time_Duration_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds())).V0)), gopurs_runtime.Float(__eta0_0)).FloatVal()
	}
}

func Call_Data_Time_Duration_convertDuration(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
	_ = dictDuration_0
	// TAST (Let): fromDuration1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope45)] Number)
	fromDuration1_1_0 := Call_Data_Time_Duration_fromDuration(dictDuration_0)
	_ = fromDuration1_1_0
	return gopurs_runtime.Func(func(dictDuration1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Time_Duration_toDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration1_2)), fromDuration1_1_0)
	})
}

func Rebox_Data_Time_Duration_1457896548_3362292604(in *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) *Constructor_Data_Time_Duration_Duration[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Time_Duration_Duration[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_1640832054_4179793454(in *Constructor_Data_Semigroup_Semigroup[float64]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_321927638_1201789390(in *Constructor_Data_Monoid_Monoid[float64]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Time_Duration_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_3362292604_1457896548(in *Constructor_Data_Time_Duration_Duration[float64]) *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_3790796878_687527510(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_4177771502_3047586294(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_4179793454_1640832054(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_4289840808_385277032(in *Constructor_Data_Newtype_Newtype[float64, float64]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Duration_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
