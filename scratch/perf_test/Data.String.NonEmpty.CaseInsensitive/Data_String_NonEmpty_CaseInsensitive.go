package Data_String_NonEmpty_CaseInsensitive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var CaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_CaseInsensitiveNonEmptyString sync.Once
func Get_CaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_CaseInsensitiveNonEmptyString.Do(func() {
		CaseInsensitiveNonEmptyString = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return CaseInsensitiveNonEmptyString
}

var showCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_showCaseInsensitiveNonEmptyString sync.Once
func Get_showCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_showCaseInsensitiveNonEmptyString.Do(func() {
		showCaseInsensitiveNonEmptyString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(CaseInsensitiveNonEmptyString ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmptyString.unsafeFromString ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0)), gopurs_runtime.Str(")")))), gopurs_runtime.Str(")")))
})})
	})
	return showCaseInsensitiveNonEmptyString
}

var newtypeCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_newtypeCaseInsensitiveNonEmptyString sync.Once
func Get_newtypeCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_newtypeCaseInsensitiveNonEmptyString.Do(func() {
		newtypeCaseInsensitiveNonEmptyString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeCaseInsensitiveNonEmptyString
}

var eqCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_eqCaseInsensitiveNonEmptyString sync.Once
func Get_eqCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_eqCaseInsensitiveNonEmptyString.Do(func() {
		eqCaseInsensitiveNonEmptyString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0)), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1))
})
})})
	})
	return eqCaseInsensitiveNonEmptyString
}

var ordCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_ordCaseInsensitiveNonEmptyString sync.Once
func Get_ordCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_ordCaseInsensitiveNonEmptyString.Do(func() {
		ordCaseInsensitiveNonEmptyString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordString().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0)), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1))
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCaseInsensitiveNonEmptyString()
})})
	})
	return ordCaseInsensitiveNonEmptyString
}


