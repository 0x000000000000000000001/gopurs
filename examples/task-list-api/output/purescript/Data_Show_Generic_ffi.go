package purescript

import "gopurs/output/gopurs_runtime"

import "strings"

func Data_Show_Generic_Intercalate(separator string, arr []string) string {
	return strings.Join(arr, separator)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Show_Generic_Intercalate = // TAST: (Func [String, (Array String)] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := func() []string {
		_source := arg1
		_out := make([]string, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.Unbox[string](gopurs_runtime.ArrayAccess(_source, _i))
		}
		return _out
	}()
	go_res := Data_Show_Generic_Intercalate(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
