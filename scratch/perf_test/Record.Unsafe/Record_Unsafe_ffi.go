package Record_Unsafe

import "gopurs/output/gopurs_runtime"

var UnsafeHas = gopurs_runtime.Func(func(label gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
		_, ok := rec.PtrVal.(map[string]gopurs_runtime.Value)[label.StrVal]
		return gopurs_runtime.Bool(ok)
	})
})

var UnsafeGet = gopurs_runtime.Func(func(label gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
		return rec.PtrVal.(map[string]gopurs_runtime.Value)[label.StrVal]
	})
})

var UnsafeSet = gopurs_runtime.Func(func(label gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
			old := rec.PtrVal.(map[string]gopurs_runtime.Value)
			newMap := make(map[string]gopurs_runtime.Value, len(old)+1)
			for k, v := range old {
				newMap[k] = v
			}
			newMap[label.StrVal] = value
			return gopurs_runtime.Record(newMap)
		})
	})
})

var UnsafeDelete = gopurs_runtime.Func(func(label gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
		old := rec.PtrVal.(map[string]gopurs_runtime.Value)
		newMap := make(map[string]gopurs_runtime.Value)
		for k, v := range old {
			if k != label.StrVal {
				newMap[k] = v
			}
		}
		return gopurs_runtime.Record(newMap)
	})
})
