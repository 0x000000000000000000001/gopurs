package Effect_Class_Console

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var warnShow gopurs_runtime.Value
var once_warnShow sync.Once
func Get_warnShow() gopurs_runtime.Value {
	once_warnShow.Do(func() {
		warnShow = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], x_2)))
})
})
})
	})
	return warnShow
}

var warn gopurs_runtime.Value
var once_warn sync.Once
func Get_warn() gopurs_runtime.Value {
	once_warn.Do(func() {
		warn = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_warn(), x_1))
})
})
	})
	return warn
}

var timeLog gopurs_runtime.Value
var once_timeLog sync.Once
func Get_timeLog() gopurs_runtime.Value {
	once_timeLog.Do(func() {
		timeLog = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_timeLog(), x_1))
})
})
	})
	return timeLog
}

var timeEnd gopurs_runtime.Value
var once_timeEnd sync.Once
func Get_timeEnd() gopurs_runtime.Value {
	once_timeEnd.Do(func() {
		timeEnd = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_timeEnd(), x_1))
})
})
	})
	return timeEnd
}

var time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		time = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_time(), x_1))
})
})
	})
	return time
}

var logShow gopurs_runtime.Value
var once_logShow sync.Once
func Get_logShow() gopurs_runtime.Value {
	once_logShow.Do(func() {
		logShow = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], x_2)))
})
})
})
	})
	return logShow
}

var log gopurs_runtime.Value
var once_log sync.Once
func Get_log() gopurs_runtime.Value {
	once_log.Do(func() {
		log = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), x_1))
})
})
	})
	return log
}

var infoShow gopurs_runtime.Value
var once_infoShow sync.Once
func Get_infoShow() gopurs_runtime.Value {
	once_infoShow.Do(func() {
		infoShow = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], x_2)))
})
})
})
	})
	return infoShow
}

var info gopurs_runtime.Value
var once_info sync.Once
func Get_info() gopurs_runtime.Value {
	once_info.Do(func() {
		info = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_info(), x_1))
})
})
	})
	return info
}

var groupEnd gopurs_runtime.Value
var once_groupEnd sync.Once
func Get_groupEnd() gopurs_runtime.Value {
	once_groupEnd.Do(func() {
		groupEnd = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], pkg_Effect_Console.Get_groupEnd())
})
	})
	return groupEnd
}

var groupCollapsed gopurs_runtime.Value
var once_groupCollapsed sync.Once
func Get_groupCollapsed() gopurs_runtime.Value {
	once_groupCollapsed.Do(func() {
		groupCollapsed = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_groupCollapsed(), x_1))
})
})
	})
	return groupCollapsed
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), x_1))
})
})
	})
	return group
}

var grouped gopurs_runtime.Value
var once_grouped sync.Once
func Get_grouped() gopurs_runtime.Value {
	once_grouped.Do(func() {
		grouped = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
Bind1_2_1 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
groupEnd1_3_2 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], pkg_Effect_Console.Get_groupEnd())
return gopurs_runtime.Func(func(name_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(inner_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_group(), name_4))), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], inner_5), gopurs_runtime.Func(func(result_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], groupEnd1_3_2), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], result_7)
}))
}))
}))
})
})
})
	})
	return grouped
}

var errorShow gopurs_runtime.Value
var once_errorShow sync.Once
func Get_errorShow() gopurs_runtime.Value {
	once_errorShow.Do(func() {
		errorShow = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], x_2)))
})
})
})
	})
	return errorShow
}

var error gopurs_runtime.Value
var once_error sync.Once
func Get_error() gopurs_runtime.Value {
	once_error.Do(func() {
		error = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_error(), x_1))
})
})
	})
	return error
}

var debugShow gopurs_runtime.Value
var once_debugShow sync.Once
func Get_debugShow() gopurs_runtime.Value {
	once_debugShow.Do(func() {
		debugShow = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], x_2)))
})
})
})
	})
	return debugShow
}

var debug gopurs_runtime.Value
var once_debug sync.Once
func Get_debug() gopurs_runtime.Value {
	once_debug.Do(func() {
		debug = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], gopurs_runtime.Apply(pkg_Effect_Console.Get_debug(), x_1))
})
})
	})
	return debug
}

var clear gopurs_runtime.Value
var once_clear sync.Once
func Get_clear() gopurs_runtime.Value {
	once_clear.Do(func() {
		clear = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], pkg_Effect_Console.Get_clear())
})
	})
	return clear
}


