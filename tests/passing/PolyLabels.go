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
		    if _, ok := o.PtrVal.(map[string]gopurs_runtime.Value); ok {
                return gopurs_runtime.RecordUpdate(o, map[string]gopurs_runtime.Value{s.StrVal: a})
            }
            r := o.PtrVal.(gopurs_runtime.RecordData)
            for j, k := range r.Keys {
                if k == s.StrVal {
                    newVals := make([]gopurs_runtime.Value, len(r.Vals))
                    copy(newVals, r.Vals)
                    newVals[j] = a
                    return gopurs_runtime.RecordDict(r.Keys, newVals)
                }
            }
            newKeys := make([]string, len(r.Keys)+1)
            newVals := make([]gopurs_runtime.Value, len(r.Vals)+1)
            copy(newKeys, r.Keys)
            copy(newVals, r.Vals)
            newKeys[len(r.Keys)] = s.StrVal
            newVals[len(r.Vals)] = a
            return gopurs_runtime.RecordDict(newKeys, newVals)
		})
	})
})
