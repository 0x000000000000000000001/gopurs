package Main
import "gopurs/output/gopurs_runtime"
var UnsafeGet = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(o gopurs_runtime.Value) gopurs_runtime.Value {
		return o.PtrVal.(map[string]gopurs_runtime.Value)[s.StrVal]
	})
})
var UnsafeSet = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(o gopurs_runtime.Value) gopurs_runtime.Value {
			oldMap := o.PtrVal.(map[string]gopurs_runtime.Value)
			newMap := make(map[string]gopurs_runtime.Value)
			for k, v := range oldMap {
				newMap[k] = v
			}
			newMap[s.StrVal] = a
			return gopurs_runtime.Record(newMap)
		})
	})
})
