package Data_Show_Generic

import (
	"gopurs/output/gopurs_runtime"
	"strings"
)

var Intercalate = gopurs_runtime.Func(func(separator gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(xs gopurs_runtime.Value) gopurs_runtime.Value {
		sep := separator.StrVal
		arr := xs.PtrVal.([]gopurs_runtime.Value)
		var strs []string
		for _, v := range arr {
			strs = append(strs, v.StrVal)
		}
		return gopurs_runtime.Str(strings.Join(strs, sep))
	})
})
