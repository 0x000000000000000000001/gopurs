package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Parser_jsonParser gopurs_runtime.Value
var once_Data_Argonaut_Parser_jsonParser sync.Once

func Get_Data_Argonaut_Parser_jsonParser() gopurs_runtime.Value {
	once_Data_Argonaut_Parser_jsonParser.Do(func() {
		cache_Data_Argonaut_Parser_jsonParser = gopurs_runtime.Func(func(j_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Parser_jsonParser(j_0_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Parser_jsonParser
}

func Call_Data_Argonaut_Parser_jsonParser(j_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var j_0 string = j_0_loop
	_ = j_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := gopurs_runtime.UncurriedApp3(Get_Data_Argonaut_Parser__jsonParser(), Get_Data_Either_Left(), Get_Data_Either_Right(), gopurs_runtime.Str(j_0)); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
}

func Get_Data_Argonaut_Parser__jsonParser() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Parser__JsonParser
}
