package Data_String_Regex_Flags

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
)

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Apply(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray())
	})
	return guard
}

var eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		eq = gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), pkg_Data_Eq.Get_eqStringImpl())
	})
	return eq
}

var RegexFlags gopurs_runtime.Value
var once_RegexFlags sync.Once
func Get_RegexFlags() gopurs_runtime.Value {
	once_RegexFlags.Do(func() {
		RegexFlags = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return RegexFlags
}

var unicode gopurs_runtime.Value
var once_unicode sync.Once
func Get_unicode() gopurs_runtime.Value {
	once_unicode.Do(func() {
		unicode = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(false), "ignoreCase": gopurs_runtime.Bool(false), "multiline": gopurs_runtime.Bool(false), "dotAll": gopurs_runtime.Bool(false), "sticky": gopurs_runtime.Bool(false), "unicode": gopurs_runtime.Bool(true)})
	})
	return unicode
}

var sticky gopurs_runtime.Value
var once_sticky sync.Once
func Get_sticky() gopurs_runtime.Value {
	once_sticky.Do(func() {
		sticky = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(false), "ignoreCase": gopurs_runtime.Bool(false), "multiline": gopurs_runtime.Bool(false), "dotAll": gopurs_runtime.Bool(false), "sticky": gopurs_runtime.Bool(true), "unicode": gopurs_runtime.Bool(false)})
	})
	return sticky
}

var showRegexFlags gopurs_runtime.Value
var once_showRegexFlags sync.Once
func Get_showRegexFlags() gopurs_runtime.Value {
	once_showRegexFlags.Do(func() {
		showRegexFlags = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
usedFlags_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{})), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
})), gopurs_runtime.Apply(Get_guard(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["global"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
})), gopurs_runtime.Apply(Get_guard(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["ignoreCase"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
})), gopurs_runtime.Apply(Get_guard(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["multiline"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
})), gopurs_runtime.Apply(Get_guard(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["dotAll"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
})), gopurs_runtime.Apply(Get_guard(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["sticky"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
})), gopurs_runtime.Apply(Get_guard(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["unicode"]))))))))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eq(), usedFlags_1_0), gopurs_runtime.Array([]gopurs_runtime.Value{}))).IntVal != 0 {
__t1 = gopurs_runtime.Str("noFlags")
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Common.Get_joinWith(), gopurs_runtime.Str(" <> ")), usedFlags_1_0)), gopurs_runtime.Str(")")))
}
return __t1
})})
	})
	return showRegexFlags
}

var semigroupRegexFlags gopurs_runtime.Value
var once_semigroupRegexFlags sync.Once
func Get_semigroupRegexFlags() gopurs_runtime.Value {
	once_semigroupRegexFlags.Do(func() {
		semigroupRegexFlags = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["global"]), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["global"]), "ignoreCase": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["ignoreCase"]), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["ignoreCase"]), "multiline": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["multiline"]), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["multiline"]), "dotAll": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["dotAll"]), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["dotAll"]), "sticky": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["sticky"]), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["sticky"]), "unicode": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["unicode"]), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["unicode"])})
})
})})
	})
	return semigroupRegexFlags
}

var noFlags gopurs_runtime.Value
var once_noFlags sync.Once
func Get_noFlags() gopurs_runtime.Value {
	once_noFlags.Do(func() {
		noFlags = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(false), "ignoreCase": gopurs_runtime.Bool(false), "multiline": gopurs_runtime.Bool(false), "dotAll": gopurs_runtime.Bool(false), "sticky": gopurs_runtime.Bool(false), "unicode": gopurs_runtime.Bool(false)})
	})
	return noFlags
}

var newtypeRegexFlags gopurs_runtime.Value
var once_newtypeRegexFlags sync.Once
func Get_newtypeRegexFlags() gopurs_runtime.Value {
	once_newtypeRegexFlags.Do(func() {
		newtypeRegexFlags = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeRegexFlags
}

var multiline gopurs_runtime.Value
var once_multiline sync.Once
func Get_multiline() gopurs_runtime.Value {
	once_multiline.Do(func() {
		multiline = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(false), "ignoreCase": gopurs_runtime.Bool(false), "multiline": gopurs_runtime.Bool(true), "dotAll": gopurs_runtime.Bool(false), "sticky": gopurs_runtime.Bool(false), "unicode": gopurs_runtime.Bool(false)})
	})
	return multiline
}

var monoidRegexFlags gopurs_runtime.Value
var once_monoidRegexFlags sync.Once
func Get_monoidRegexFlags() gopurs_runtime.Value {
	once_monoidRegexFlags.Do(func() {
		monoidRegexFlags = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": Get_noFlags(), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupRegexFlags()
})})
	})
	return monoidRegexFlags
}

var ignoreCase gopurs_runtime.Value
var once_ignoreCase sync.Once
func Get_ignoreCase() gopurs_runtime.Value {
	once_ignoreCase.Do(func() {
		ignoreCase = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(false), "ignoreCase": gopurs_runtime.Bool(true), "multiline": gopurs_runtime.Bool(false), "dotAll": gopurs_runtime.Bool(false), "sticky": gopurs_runtime.Bool(false), "unicode": gopurs_runtime.Bool(false)})
	})
	return ignoreCase
}

var global gopurs_runtime.Value
var once_global sync.Once
func Get_global() gopurs_runtime.Value {
	once_global.Do(func() {
		global = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(true), "ignoreCase": gopurs_runtime.Bool(false), "multiline": gopurs_runtime.Bool(false), "dotAll": gopurs_runtime.Bool(false), "sticky": gopurs_runtime.Bool(false), "unicode": gopurs_runtime.Bool(false)})
	})
	return global
}

var eqRegexFlags gopurs_runtime.Value
var once_eqRegexFlags sync.Once
func Get_eqRegexFlags() gopurs_runtime.Value {
	once_eqRegexFlags.Do(func() {
		eqRegexFlags = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), pkg_Data_Eq.Get_eqRowNil()), gopurs_runtime.Value{}), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectSymbol": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
})})), pkg_Data_Eq.Get_eqBoolean())), gopurs_runtime.Value{}), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectSymbol": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
})})), pkg_Data_Eq.Get_eqBoolean())), gopurs_runtime.Value{}), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectSymbol": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
})})), pkg_Data_Eq.Get_eqBoolean())), gopurs_runtime.Value{}), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectSymbol": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
})})), pkg_Data_Eq.Get_eqBoolean())), gopurs_runtime.Value{}), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectSymbol": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
})})), pkg_Data_Eq.Get_eqBoolean())), gopurs_runtime.Value{}), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectSymbol": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
})})), pkg_Data_Eq.Get_eqBoolean()).PtrVal.(map[string]gopurs_runtime.Value)["eqRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
	})
	return eqRegexFlags
}

var dotAll gopurs_runtime.Value
var once_dotAll sync.Once
func Get_dotAll() gopurs_runtime.Value {
	once_dotAll.Do(func() {
		dotAll = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"global": gopurs_runtime.Bool(false), "ignoreCase": gopurs_runtime.Bool(false), "multiline": gopurs_runtime.Bool(false), "dotAll": gopurs_runtime.Bool(true), "sticky": gopurs_runtime.Bool(false), "unicode": gopurs_runtime.Bool(false)})
	})
	return dotAll
}


