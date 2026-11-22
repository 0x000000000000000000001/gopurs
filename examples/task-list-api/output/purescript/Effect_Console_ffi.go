package purescript

import "gopurs/output/gopurs_runtime"

import (
	"fmt"
	"os"
	"strings"
	"sync/atomic"
)

var indentation int32 = 0

func getIndent() string {
	ind := atomic.LoadInt32(&indentation)
	if ind < 0 {
		ind = 0
	}
	return strings.Repeat("  ", int(ind))
}

func Effect_Console_Log(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_Warn(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_Error(s string, _ interface{}) interface{} {
	fmt.Fprintf(os.Stderr, "%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_Info(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_Debug(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_Time(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_TimeLog(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_TimeEnd(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	return nil
}

func Effect_Console_Clear(_ interface{}) interface{} {
	return nil
}

func Effect_Console_Group(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	atomic.AddInt32(&indentation, 1)
	return nil
}

func Effect_Console_GroupCollapsed(s string, _ interface{}) interface{} {
	fmt.Printf("%s%s\n", getIndent(), s)
	atomic.AddInt32(&indentation, 1)
	return nil
}

func Effect_Console_GroupEnd(_ interface{}) interface{} {
	ind := atomic.LoadInt32(&indentation)
	if ind > 0 {
		atomic.AddInt32(&indentation, -1)
	}
	return nil
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Effect_Console_Clear = // TAST: (ADT ["Effect","Effect"] [Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Console_Clear(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Debug = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Debug(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Error = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Error(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Group = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Group(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_GroupCollapsed = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_GroupCollapsed(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_GroupEnd = // TAST: (ADT ["Effect","Effect"] [Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Console_GroupEnd(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Info = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Info(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Log = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Log(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Time = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Time(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_TimeEnd = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_TimeEnd(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_TimeLog = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_TimeLog(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Console_Warn = // TAST: (Func [String] (ADT ["Effect","Effect"] [Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := arg1
	go_res := Effect_Console_Warn(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
