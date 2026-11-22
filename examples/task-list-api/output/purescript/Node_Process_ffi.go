package purescript

import (
	"gopurs/output/gopurs_runtime"
	"os"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

var startTime = time.Now()
var exitCode int64 = 0

func Node_Process_ExitImpl(code int64) interface{} {
	os.Exit(int(code))
	return nil
}

func Node_Process_Process(_ interface{}) interface{} {
	panic("Not implemented: process")
}

func Node_Process_AbortImpl(_ interface{}) interface{} {
	panic("Not implemented: abortImpl")
}

func Node_Process_Argv(_ interface{}) interface{} {
	arr := make([]gopurs_runtime.Value, len(os.Args))
	for i, arg := range os.Args {
		arr[i] = gopurs_runtime.Str(arg)
	}
	return gopurs_runtime.Value{Type: gopurs_runtime.TypeArray, UnsafePtr: unsafe.Pointer(&arr)}
}

func Node_Process_Argv0(_ interface{}) interface{} {
	if len(os.Args) > 0 {
		return gopurs_runtime.Str(os.Args[0])
	}
	return gopurs_runtime.Str("")
}

func Node_Process_ChannelRefImpl(_ interface{}) interface{} {
	panic("Not implemented: channelRefImpl")
}

func Node_Process_ChannelUnrefImpl(_ interface{}) interface{} {
	panic("Not implemented: channelUnrefImpl")
}

func Node_Process_ChdirImpl(dirVal interface{}) interface{} {
	dir := gopurs_runtime.Unbox[string](dirVal)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	return nil
}

func Node_Process_Config(_ interface{}) interface{} {
	panic("Not implemented: config")
}

func Node_Process_Connected(_ interface{}) interface{} {
	panic("Not implemented: connected")
}

func Node_Process_CpuUsage(_ interface{}) interface{} {
	panic("Not implemented: cpuUsage")
}

func Node_Process_CpuUsageDiffImpl(_ interface{}) interface{} {
	panic("Not implemented: cpuUsageDiffImpl")
}

func Node_Process_Cwd(_ interface{}) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func Node_Process_DebugPort(_ interface{}) interface{} {
	panic("Not implemented: debugPort")
}

func Node_Process_DisconnectImpl(_ interface{}) interface{} {
	panic("Not implemented: disconnectImpl")
}

func Node_Process_GetEnv(_ interface{}) interface{} {
	env := make(map[string]interface{})
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}
	return env
}

func Node_Process_UnsafeGetEnv(_ interface{}) interface{} {
	return Node_Process_GetEnv(nil)
}

func Node_Process_SetEnvImpl(keyVal interface{}, valVal interface{}) interface{} {
	key := gopurs_runtime.Unbox[string](keyVal)
	val := gopurs_runtime.Unbox[string](valVal)
	os.Setenv(key, val)
	return nil
}

func Node_Process_UnsetEnvImpl(keyVal interface{}) interface{} {
	key := gopurs_runtime.Unbox[string](keyVal)
	os.Unsetenv(key)
	return nil
}

// ExecArgv reports Node engine options, which do not exist in a native Go
// executable. Application arguments remain available through Argv.
func Node_Process_ExecArgv() []string {
	return []string{}
}

func Node_Process_ExecPath(_ interface{}) string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return exe
}

func Node_Process_Exit(_ interface{}) interface{} {
	os.Exit(int(exitCode))
	return nil
}

func Node_Process_SetExitCodeImpl(codeVal interface{}) interface{} {
	code := gopurs_runtime.Unbox[int64](codeVal)
	exitCode = code
	return nil
}

func Node_Process_GetExitCodeImpl(_ interface{}) interface{} {
	return exitCode
}

func Node_Process_GetGidImpl(_ interface{}) interface{} {
	return int64(os.Getgid())
}

func Node_Process_GetUidImpl(_ interface{}) interface{} {
	return int64(os.Getuid())
}

func Node_Process_HasUncaughtExceptionCaptureCallback(_ interface{}) interface{} {
	panic("Not implemented: hasUncaughtExceptionCaptureCallback")
}

func Node_Process_KillImpl(pidVal interface{}) interface{} {
	pid := int(gopurs_runtime.Unbox[int64](pidVal))
	p, err := os.FindProcess(pid)
	if err == nil {
		p.Kill()
	}
	return nil
}

func Node_Process_KillStrImpl(pidVal interface{}, sigVal interface{}) interface{} {
	panic("Not implemented: killStrImpl")
}

func Node_Process_KillIntImpl(pidVal interface{}, sigVal interface{}) interface{} {
	panic("Not implemented: killIntImpl")
}

func Node_Process_MemoryUsage(_ interface{}) interface{} {
	panic("Not implemented: memoryUsage")
}

func Node_Process_MemoryUsageRss(_ interface{}) interface{} {
	panic("Not implemented: memoryUsageRss")
}

func Node_Process_NextTickImpl(_ interface{}) interface{} {
	panic("Not implemented: nextTickImpl")
}

func Node_Process_NextTickCbImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: nextTickCbImpl")
}

var Node_Process_Pid = int64(os.Getpid())

var Node_Process_PlatformStr = runtime.GOOS

func Node_Process_Ppid(_ interface{}) interface{} {
	return int64(os.Getppid())
}

func Node_Process_ResourceUsage(_ interface{}) interface{} {
	panic("Not implemented: resourceUsage")
}

func Node_Process_SendImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendImpl")
}

func Node_Process_SendOptsImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendOptsImpl")
}

func Node_Process_SendCbImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendCbImpl")
}

func Node_Process_SendOptsCbImpl(_ interface{}, _ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendOptsCbImpl")
}

func Node_Process_SetUncaughtExceptionCaptureCallbackImpl(_ interface{}) interface{} {
	panic("Not implemented: setUncaughtExceptionCaptureCallbackImpl")
}

func Node_Process_ClearUncaughtExceptionCaptureCallback(_ interface{}) interface{} {
	panic("Not implemented: clearUncaughtExceptionCaptureCallback")
}

func isTTY(f *os.File) bool {
	stat, err := f.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func Node_Process_Stdin(_ interface{}) interface{} {
	panic("Not implemented: stdin")
}

func Node_Process_Stdout(_ interface{}) interface{} {
	panic("Not implemented: stdout")
}

func Node_Process_Stderr(_ interface{}) interface{} {
	panic("Not implemented: stderr")
}

func Node_Process_StdinIsTTY(_ interface{}) interface{} {
	return isTTY(os.Stdin)
}

func Node_Process_StdoutIsTTY(_ interface{}) interface{} {
	return isTTY(os.Stdout)
}

func Node_Process_StderrIsTTY(_ interface{}) interface{} {
	return isTTY(os.Stderr)
}

func Node_Process_GetTitle(_ interface{}) interface{} {
	if len(os.Args) > 0 {
		return os.Args[0]
	}
	return ""
}

func Node_Process_SetTitleImpl(_ interface{}) interface{} {
	return nil
}

func Node_Process_Uptime(_ interface{}) interface{} {
	return time.Since(startTime).Seconds()
}

var Node_Process_Version = runtime.Version()

// --- Auto-generated FFI wrappers ---
var _Gopurs_Node_Process_AbortImpl = // TAST: (ADT ["Data","Nullable","Nullable"] [(ADT ["Effect","Effect"] [Unit])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_AbortImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Argv = // TAST: (ADT ["Effect","Effect"] [(Array String)])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Argv(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Argv0 = // TAST: (ADT ["Effect","Effect"] [String])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Argv0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ChannelRefImpl = // TAST: (ADT ["Data","Nullable","Nullable"] [(ADT ["Effect","Effect"] [Unit])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_ChannelRefImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ChannelUnrefImpl = // TAST: (ADT ["Data","Nullable","Nullable"] [(ADT ["Effect","Effect"] [Unit])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_ChannelUnrefImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ChdirImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [String, Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_ChdirImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ClearUncaughtExceptionCaptureCallback = // TAST: (ADT ["Effect","Effect"] [Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_ClearUncaughtExceptionCaptureCallback(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Config = // TAST: (ADT ["Effect","Effect"] [(ADT ["Foreign","Foreign"] [])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Config(go_arg0)
	return gopurs_runtime.Any(go_res)
})
var _Gopurs_Node_Process_Connected = // TAST: (ADT ["Effect","Effect"] [Boolean])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Connected(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_CpuUsage = // TAST: (ADT ["Effect","Effect"] [(Record (Row [user: Int, system: Int] Empty))])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_CpuUsage(go_arg0)
	return func() gopurs_runtime.Value {
		_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(go_res))
		res_map := make(map[string]gopurs_runtime.Value)
		res_map["user"] = gopurs_runtime.Box(_raw["user"])
		res_map["system"] = gopurs_runtime.Box(_raw["system"])
		return gopurs_runtime.Record(res_map)
	}()
})
var _Gopurs_Node_Process_CpuUsageDiffImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [(Record (Row [user: Int, system: Int] Empty)), (Record (Row [user: Int, system: Int] Empty))])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_CpuUsageDiffImpl(go_arg0)
	return func() gopurs_runtime.Value {
		_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(go_res))
		res_map := make(map[string]gopurs_runtime.Value)
		res_map["user"] = gopurs_runtime.Box(_raw["user"])
		res_map["system"] = gopurs_runtime.Box(_raw["system"])
		return gopurs_runtime.Record(res_map)
	}()
})
var _Gopurs_Node_Process_Cwd = // TAST: (ADT ["Effect","Effect"] [String])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Cwd(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_DebugPort = // TAST: Int
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_DebugPort(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_DisconnectImpl = // TAST: (ADT ["Data","Nullable","Nullable"] [(ADT ["Effect","Effect"] [Unit])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_DisconnectImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ExecArgv = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {

	go_res := Node_Process_ExecArgv()
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Node_Process_ExecPath = // TAST: (ADT ["Effect","Effect"] [String])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_ExecPath(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Exit = // TAST: (ForAll [a] (ADT ["Effect","Effect"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Exit(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ExitImpl = // TAST: (ForAll [a] (ADT ["Effect","Uncurried","EffectFn1"] [Int, (TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_res := Node_Process_ExitImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_GetEnv = // TAST: (ADT ["Effect","Effect"] [(ADT ["Foreign","Object","Object"] [String])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_GetEnv(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_GetExitCodeImpl = // TAST: (ADT ["Effect","Effect"] [(ADT ["Data","Nullable","Nullable"] [Int])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_GetExitCodeImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_GetGidImpl = // TAST: (ADT ["Effect","Effect"] [(ADT ["Data","Nullable","Nullable"] [Int])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_GetGidImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_GetTitle = // TAST: (ADT ["Effect","Effect"] [String])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_GetTitle(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_GetUidImpl = // TAST: (ADT ["Effect","Effect"] [(ADT ["Data","Nullable","Nullable"] [Int])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_GetUidImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_HasUncaughtExceptionCaptureCallback = // TAST: (ADT ["Effect","Effect"] [Boolean])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_HasUncaughtExceptionCaptureCallback(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_KillImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [Int, Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_KillImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_KillIntImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn2"] [Int, Int, Unit])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Node_Process_KillIntImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_KillStrImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn2"] [Int, String, Unit])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Node_Process_KillStrImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_MemoryUsage = // TAST: (ADT ["Effect","Effect"] [(Record (Row [rss: Int, heapTotal: Int, heapUsed: Int, external: Int, arrayBuffers: Int] Empty))])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_MemoryUsage(go_arg0)
	return func() gopurs_runtime.Value {
		_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(go_res))
		res_map := make(map[string]gopurs_runtime.Value)
		res_map["rss"] = gopurs_runtime.Box(_raw["rss"])
		res_map["heapTotal"] = gopurs_runtime.Box(_raw["heapTotal"])
		res_map["heapUsed"] = gopurs_runtime.Box(_raw["heapUsed"])
		res_map["external"] = gopurs_runtime.Box(_raw["external"])
		res_map["arrayBuffers"] = gopurs_runtime.Box(_raw["arrayBuffers"])
		return gopurs_runtime.Record(res_map)
	}()
})
var _Gopurs_Node_Process_MemoryUsageRss = // TAST: (ADT ["Effect","Effect"] [Int])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_MemoryUsageRss(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_NextTickCbImpl = // TAST: (ForAll [r] (ADT ["Effect","Uncurried","EffectFn2"] [(ADT ["Effect","Uncurried","EffectFn1"] [(Record (Row [] (TypeVar r))), Unit]), (Record (Row [] (TypeVar r))), Unit]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Node_Process_NextTickCbImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_NextTickImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [(ADT ["Effect","Effect"] [Unit]), Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_NextTickImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Pid = // TAST: Int
gopurs_runtime.Box(Node_Process_Pid)
var _Gopurs_Node_Process_PlatformStr = // TAST: String
gopurs_runtime.Box(Node_Process_PlatformStr)
var _Gopurs_Node_Process_Ppid = // TAST: Int
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Ppid(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Process = // TAST: (ADT ["Node","Process","Process"] [])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Process(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_ResourceUsage = // TAST: (ADT ["Effect","Effect"] [(Record (Row [userCPUTime: Int, systemCPUTime: Int, maxRSS: Int, sharedMemorySize: Int, unsharedDataSize: Int, unsharedStackSize: Int, minorPageFault: Int, majorPageFault: Int, swappedOut: Int, fsRead: Int, fsWrite: Int, ipcSent: Int, ipcReceived: Int, signalsCount: Int, voluntaryContextSwitches: Int, involuntaryContextSwitches: Int] Empty))])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_ResourceUsage(go_arg0)
	return func() gopurs_runtime.Value {
		_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(go_res))
		res_map := make(map[string]gopurs_runtime.Value)
		res_map["userCPUTime"] = gopurs_runtime.Box(_raw["userCPUTime"])
		res_map["systemCPUTime"] = gopurs_runtime.Box(_raw["systemCPUTime"])
		res_map["maxRSS"] = gopurs_runtime.Box(_raw["maxRSS"])
		res_map["sharedMemorySize"] = gopurs_runtime.Box(_raw["sharedMemorySize"])
		res_map["unsharedDataSize"] = gopurs_runtime.Box(_raw["unsharedDataSize"])
		res_map["unsharedStackSize"] = gopurs_runtime.Box(_raw["unsharedStackSize"])
		res_map["minorPageFault"] = gopurs_runtime.Box(_raw["minorPageFault"])
		res_map["majorPageFault"] = gopurs_runtime.Box(_raw["majorPageFault"])
		res_map["swappedOut"] = gopurs_runtime.Box(_raw["swappedOut"])
		res_map["fsRead"] = gopurs_runtime.Box(_raw["fsRead"])
		res_map["fsWrite"] = gopurs_runtime.Box(_raw["fsWrite"])
		res_map["ipcSent"] = gopurs_runtime.Box(_raw["ipcSent"])
		res_map["ipcReceived"] = gopurs_runtime.Box(_raw["ipcReceived"])
		res_map["signalsCount"] = gopurs_runtime.Box(_raw["signalsCount"])
		res_map["voluntaryContextSwitches"] = gopurs_runtime.Box(_raw["voluntaryContextSwitches"])
		res_map["involuntaryContextSwitches"] = gopurs_runtime.Box(_raw["involuntaryContextSwitches"])
		return gopurs_runtime.Record(res_map)
	}()
})
var _Gopurs_Node_Process_SendCbImpl = // TAST: (ForAll [messageRows] (ADT ["Effect","Uncurried","EffectFn3"] [(Record (Row [] (TypeVar messageRows))), (ADT ["Data","Nullable","Nullable"] [(ADT ["Foreign","Foreign"] [])]), (ADT ["Effect","Uncurried","EffectFn1"] [(ADT ["Data","Nullable","Nullable"] [(ADT ["Effect","Exception","Error"] [])]), Unit]), Boolean]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Node_Process_SendCbImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SendImpl = // TAST: (ForAll [messageRows] (ADT ["Effect","Uncurried","EffectFn2"] [(Record (Row [] (TypeVar messageRows))), (ADT ["Data","Nullable","Nullable"] [(ADT ["Foreign","Foreign"] [])]), Boolean]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Node_Process_SendImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SendOptsCbImpl = // TAST: (ForAll [messageRows, r] (ADT ["Effect","Uncurried","EffectFn4"] [(Record (Row [] (TypeVar messageRows))), (ADT ["Data","Nullable","Nullable"] [(ADT ["Foreign","Foreign"] [])]), (Record (Row [] (TypeVar r))), (ADT ["Effect","Uncurried","EffectFn1"] [(ADT ["Data","Nullable","Nullable"] [(ADT ["Effect","Exception","Error"] [])]), Unit]), Boolean]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Node_Process_SendOptsCbImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SendOptsImpl = // TAST: (ForAll [messageRows, r] (ADT ["Effect","Uncurried","EffectFn3"] [(Record (Row [] (TypeVar messageRows))), (ADT ["Data","Nullable","Nullable"] [(ADT ["Foreign","Foreign"] [])]), (Record (Row [] (TypeVar r))), Boolean]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Node_Process_SendOptsImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SetEnvImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn2"] [String, String, Unit])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Node_Process_SetEnvImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SetExitCodeImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [Int, Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_SetExitCodeImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SetTitleImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [String, Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_SetTitleImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_SetUncaughtExceptionCaptureCallbackImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [(ADT ["Effect","Effect"] [Unit]), Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_SetUncaughtExceptionCaptureCallbackImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Stderr = // TAST: (ADT ["Node","Stream","Stream"] [(Row [write: (ADT ["Node","Stream","Write"] [])] Empty)])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Stderr(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_StderrIsTTY = // TAST: Boolean
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_StderrIsTTY(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Stdin = // TAST: (ADT ["Node","Stream","Stream"] [(Row [read: (ADT ["Node","Stream","Read"] [])] Empty)])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Stdin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_StdinIsTTY = // TAST: Boolean
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_StdinIsTTY(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Stdout = // TAST: (ADT ["Node","Stream","Stream"] [(Row [write: (ADT ["Node","Stream","Write"] [])] Empty)])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Stdout(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_StdoutIsTTY = // TAST: Boolean
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_StdoutIsTTY(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_UnsafeGetEnv = // TAST: (ADT ["Effect","Effect"] [(ADT ["Foreign","Object","Object"] [String])])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_UnsafeGetEnv(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_UnsetEnvImpl = // TAST: (ADT ["Effect","Uncurried","EffectFn1"] [String, Unit])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_UnsetEnvImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Uptime = // TAST: (ADT ["Effect","Effect"] [Number])
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Uptime(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Node_Process_Version = // TAST: String
gopurs_runtime.Box(Node_Process_Version)

// --- Boxed-result FFI workers ---
func Node_Process_DebugPort_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_DebugPort(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_Ppid_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Ppid(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_Process_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Process(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_Stderr_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Stderr(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_StderrIsTTY_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_StderrIsTTY(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_Stdin_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Stdin(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_StdinIsTTY_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_StdinIsTTY(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_Stdout_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_Stdout(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Node_Process_StdoutIsTTY_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Node_Process_StdoutIsTTY(go_arg0)
	return gopurs_runtime.Box(go_res)
}
