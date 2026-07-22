package Partial

import "gopurs/output/gopurs_runtime"

var X_CrashWith = gopurs_runtime.Func(func(msg gopurs_runtime.Value) gopurs_runtime.Value {
	panic(msg.StrVal)
})
