package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Console_discard gopurs_runtime.Value
var once_Effect_Console_discard sync.Once

func Get_Effect_Console_discard() gopurs_runtime.Value {
	once_Effect_Console_discard.Do(func() {
		cache_Effect_Console_discard = Get_Effect_bindE()
	})
	return cache_Effect_Console_discard
}

var cache_Effect_Console_bind gopurs_runtime.Value
var once_Effect_Console_bind sync.Once

func Get_Effect_Console_bind() gopurs_runtime.Value {
	once_Effect_Console_bind.Do(func() {
		cache_Effect_Console_bind = Get_Effect_bindE()
	})
	return cache_Effect_Console_bind
}

var cache_Effect_Console_pure gopurs_runtime.Value
var once_Effect_Console_pure sync.Once

func Get_Effect_Console_pure() gopurs_runtime.Value {
	once_Effect_Console_pure.Do(func() {
		cache_Effect_Console_pure = Get_Effect_pureE()
	})
	return cache_Effect_Console_pure
}

var cache_Effect_Console_warnShow gopurs_runtime.Value
var once_Effect_Console_warnShow sync.Once

func Get_Effect_Console_warnShow() gopurs_runtime.Value {
	once_Effect_Console_warnShow.Do(func() {
		cache_Effect_Console_warnShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Console_warnShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
		})
	})
	return cache_Effect_Console_warnShow
}

var cache_Effect_Console_logShow gopurs_runtime.Value
var once_Effect_Console_logShow sync.Once

func Get_Effect_Console_logShow() gopurs_runtime.Value {
	once_Effect_Console_logShow.Do(func() {
		cache_Effect_Console_logShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Console_logShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
		})
	})
	return cache_Effect_Console_logShow
}

var cache_Effect_Console_infoShow gopurs_runtime.Value
var once_Effect_Console_infoShow sync.Once

func Get_Effect_Console_infoShow() gopurs_runtime.Value {
	once_Effect_Console_infoShow.Do(func() {
		cache_Effect_Console_infoShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Console_infoShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
		})
	})
	return cache_Effect_Console_infoShow
}

var cache_Effect_Console_grouped gopurs_runtime.Value
var once_Effect_Console_grouped sync.Once

func Get_Effect_Console_grouped() gopurs_runtime.Value {
	once_Effect_Console_grouped.Do(func() {
		cache_Effect_Console_grouped = gopurs_runtime.Func2(func(name_0_box gopurs_runtime.Value, inner_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Console_grouped(name_0_box.StrVal(), inner_1_box)
		})
	})
	return cache_Effect_Console_grouped
}

var cache_Effect_Console_errorShow gopurs_runtime.Value
var once_Effect_Console_errorShow sync.Once

func Get_Effect_Console_errorShow() gopurs_runtime.Value {
	once_Effect_Console_errorShow.Do(func() {
		cache_Effect_Console_errorShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Console_errorShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
		})
	})
	return cache_Effect_Console_errorShow
}

var cache_Effect_Console_debugShow gopurs_runtime.Value
var once_Effect_Console_debugShow sync.Once

func Get_Effect_Console_debugShow() gopurs_runtime.Value {
	once_Effect_Console_debugShow.Do(func() {
		cache_Effect_Console_debugShow = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Console_debugShow(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
		})
	})
	return cache_Effect_Console_debugShow
}

func Call_Effect_Console_warnShow(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply(Get_Effect_Console_warn(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictShow_0.V0, a_1)))
}

func Call_Effect_Console_logShow(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictShow_0.V0, a_1)))
}

func Call_Effect_Console_infoShow(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply(Get_Effect_Console_info(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictShow_0.V0, a_1)))
}

func Call_Effect_Console_grouped(name_0_loop string, inner_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var name_0 string = name_0_loop
	_ = name_0
	var inner_1 gopurs_runtime.Value = inner_1_loop
	_ = inner_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Console_group(), gopurs_runtime.Str(name_0))
		_ = __local_var_2_0
		_dollar___unused_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
		_ = _dollar___unused_3_1
		result_4_2 := gopurs_runtime.Apply(inner_1, gopurs_runtime.Value{})
		_ = result_4_2
		_dollar___unused_5_3 := gopurs_runtime.Apply(Get_Effect_Console_groupEnd(), gopurs_runtime.Value{})
		_ = _dollar___unused_5_3
		return result_4_2
	})
}

func Call_Effect_Console_errorShow(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictShow_0.V0, a_1)))
}

func Call_Effect_Console_debugShow(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	return gopurs_runtime.Apply(Get_Effect_Console_debug(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictShow_0.V0, a_1)))
}

func Get_Effect_Console_clear() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Clear
}

func Get_Effect_Console_debug() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Debug
}

func Get_Effect_Console_error() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Error
}

func Get_Effect_Console_group() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Group
}

func Get_Effect_Console_groupCollapsed() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_GroupCollapsed
}

func Get_Effect_Console_groupEnd() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_GroupEnd
}

func Get_Effect_Console_info() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Info
}

func Get_Effect_Console_log() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Log
}

func Get_Effect_Console_time() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Time
}

func Get_Effect_Console_timeEnd() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_TimeEnd
}

func Get_Effect_Console_timeLog() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_TimeLog
}

func Get_Effect_Console_warn() gopurs_runtime.Value {
	return _Gopurs_Effect_Console_Warn
}
