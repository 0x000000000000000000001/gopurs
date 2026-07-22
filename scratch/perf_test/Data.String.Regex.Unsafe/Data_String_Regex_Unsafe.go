package Data_String_Regex_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_String_Regex "gopurs/output/Data.String.Regex"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var unsafeRegex gopurs_runtime.Value
var once_unsafeRegex sync.Once
func Get_unsafeRegex() gopurs_runtime.Value {
	once_unsafeRegex.Do(func() {
		unsafeRegex = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Regex.Get_regexImpl(), pkg_Data_Either.Get_Left()), pkg_Data_Either.Get_Right()), s_0), gopurs_runtime.Apply(pkg_Data_String_Regex.Get_renderFlags(), f_1))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__local_var_3_2 := __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t1 = gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), __local_var_3_2)
}))
} else {
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
	})
	return unsafeRegex
}


