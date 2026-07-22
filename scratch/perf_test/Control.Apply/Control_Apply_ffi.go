package Control_Apply

import "gopurs/output/gopurs_runtime"

var ArrayApply = gopurs_runtime.Func(func(fs gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		fsArr := fs.PtrVal.([]gopurs_runtime.Value)
		xsArr := xs.PtrVal.([]gopurs_runtime.Value)
		result := make([]gopurs_runtime.Value, 0, len(fsArr)*len(xsArr))
		for _, f := range fsArr {
			for _, x := range xsArr {
				result = append(result, gopurs_runtime.Apply(f, x))
			}
		}
		return gopurs_runtime.Array(result)
	})
})
