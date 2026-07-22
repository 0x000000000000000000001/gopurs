package Effect_Console

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
)

var warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		warnShow = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_warn(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], a_1))
})
})
	})
	return warnShow
}

var logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		logShow = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_log(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], a_1))
})
})
	})
	return logShow
}

var infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		infoShow = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_info(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], a_1))
})
})
	})
	return infoShow
}

var grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		grouped = gopurs_runtime.Func(func(name_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(inner_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(Get_group(), name_0)), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), inner_1), gopurs_runtime.Func(func(result_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), Get_groupEnd()), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect.Get_pureE(), result_3)
}))
}))
}))
})
})
	})
	return grouped
}

var errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		errorShow = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_error(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], a_1))
})
})
	})
	return errorShow
}

var debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		debugShow = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_debug(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], a_1))
})
})
	})
	return debugShow
}

func Get_clear() gopurs_runtime.Value {
	return Clear
}

func Get_debug() gopurs_runtime.Value {
	return Debug
}

func Get_error() gopurs_runtime.Value {
	return Error
}

func Get_group() gopurs_runtime.Value {
	return Group
}

func Get_groupCollapsed() gopurs_runtime.Value {
	return GroupCollapsed
}

func Get_groupEnd() gopurs_runtime.Value {
	return GroupEnd
}

func Get_info() gopurs_runtime.Value {
	return Info
}

func Get_log() gopurs_runtime.Value {
	return Log
}

func Get_time() gopurs_runtime.Value {
	return Time
}

func Get_timeEnd() gopurs_runtime.Value {
	return TimeEnd
}

func Get_timeLog() gopurs_runtime.Value {
	return TimeLog
}

func Get_warn() gopurs_runtime.Value {
	return Warn
}
