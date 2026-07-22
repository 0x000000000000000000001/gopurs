package Effect_Console

import (
	"fmt"
	"gopurs/output/gopurs_runtime"
)

var Log = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		fmt.Println(s.StrVal)
		return gopurs_runtime.Value{}
	})
})

var Warn = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		fmt.Println("[WARN]", s.StrVal)
		return gopurs_runtime.Value{}
	})
})

var Error = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		fmt.Println("[ERROR]", s.StrVal)
		return gopurs_runtime.Value{}
	})
})

var Info = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		fmt.Println("[INFO]", s.StrVal)
		return gopurs_runtime.Value{}
	})
})

// Unimplemented dummies for now
var Debug = Log
var Time = Log
var TimeLog = Log
var TimeEnd = Log
var Clear = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})
var Group = Log
var GroupCollapsed = Log
var GroupEnd = Clear
