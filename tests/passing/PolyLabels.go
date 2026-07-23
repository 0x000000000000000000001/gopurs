package Main
import "gopurs/output/gopurs_runtime"
var UnsafeGet = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(o gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordGet(o, s.StrVal)
	})
})
var UnsafeSet = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
        return gopurs_runtime.Func(func(o gopurs_runtime.Value) gopurs_runtime.Value {
		    return gopurs_runtime.RecordUpdate(o, map[string]gopurs_runtime.Value{s.StrVal: a})
		})
	})
})
