package Data_String_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), s_0), gopurs_runtime.Str(""))
})
	})
	return null
}

var localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		localeCompare = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get__localeCompare(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")}))
	})
	return localeCompare
}

func Get__localeCompare() gopurs_runtime.Value {
	return X_LocaleCompare
}

func Get_joinWith() gopurs_runtime.Value {
	return JoinWith
}

func Get_replace() gopurs_runtime.Value {
	return Replace
}

func Get_replaceAll() gopurs_runtime.Value {
	return ReplaceAll
}

func Get_split() gopurs_runtime.Value {
	return Split
}

func Get_toLower() gopurs_runtime.Value {
	return ToLower
}

func Get_toUpper() gopurs_runtime.Value {
	return ToUpper
}

func Get_trim() gopurs_runtime.Value {
	return Trim
}
