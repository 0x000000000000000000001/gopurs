package Effect_Ref

import "gopurs/output/gopurs_runtime"

var X_New = gopurs_runtime.Func(func(val gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"value": val})
	})
})

var NewWithSelf = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		ref := gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
		ref.PtrVal.(map[string]gopurs_runtime.Value)["value"] = gopurs_runtime.Apply(f, ref)
		return ref
	})
})

var Read = gopurs_runtime.Func(func(ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return ref.PtrVal.(map[string]gopurs_runtime.Value)["value"]
	})
})

var ModifyImpl = gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(ref gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			t := gopurs_runtime.Apply(f, ref.PtrVal.(map[string]gopurs_runtime.Value)["value"])
			ref.PtrVal.(map[string]gopurs_runtime.Value)["value"] = t.PtrVal.(map[string]gopurs_runtime.Value)["state"]
			return t.PtrVal.(map[string]gopurs_runtime.Value)["value"]
		})
	})
})

var Write = gopurs_runtime.Func(func(val gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(ref gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			ref.PtrVal.(map[string]gopurs_runtime.Value)["value"] = val
			return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
		})
	})
})
