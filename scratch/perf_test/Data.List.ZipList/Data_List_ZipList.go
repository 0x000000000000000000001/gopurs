package Data_List_ZipList

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_List_Lazy "gopurs/output/Data.List.Lazy"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var ZipList gopurs_runtime.Value
var once_ZipList sync.Once
func Get_ZipList() gopurs_runtime.Value {
	once_ZipList.Do(func() {
		ZipList = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return ZipList
}

var traversableZipList gopurs_runtime.Value
var once_traversableZipList sync.Once
func Get_traversableZipList() gopurs_runtime.Value {
	once_traversableZipList.Do(func() {
		traversableZipList = pkg_Data_List_Lazy_Types.Get_traversableList()
	})
	return traversableZipList
}

var showZipList gopurs_runtime.Value
var once_showZipList sync.Once
func Get_showZipList() gopurs_runtime.Value {
	once_showZipList.Do(func() {
		showZipList = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(ZipList ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_showList(), dictShow_0).PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showZipList
}

var semigroupZipList gopurs_runtime.Value
var once_semigroupZipList sync.Once
func Get_semigroupZipList() gopurs_runtime.Value {
	once_semigroupZipList.Do(func() {
		semigroupZipList = pkg_Data_List_Lazy_Types.Get_semigroupList()
	})
	return semigroupZipList
}

var ordZipList gopurs_runtime.Value
var once_ordZipList sync.Once
func Get_ordZipList() gopurs_runtime.Value {
	once_ordZipList.Do(func() {
		ordZipList = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_ordList(), dictOrd_0)
})
	})
	return ordZipList
}

var newtypeZipList gopurs_runtime.Value
var once_newtypeZipList sync.Once
func Get_newtypeZipList() gopurs_runtime.Value {
	once_newtypeZipList.Do(func() {
		newtypeZipList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeZipList
}

var monoidZipList gopurs_runtime.Value
var once_monoidZipList sync.Once
func Get_monoidZipList() gopurs_runtime.Value {
	once_monoidZipList.Do(func() {
		monoidZipList = pkg_Data_List_Lazy_Types.Get_monoidList()
	})
	return monoidZipList
}

var functorZipList gopurs_runtime.Value
var once_functorZipList sync.Once
func Get_functorZipList() gopurs_runtime.Value {
	once_functorZipList.Do(func() {
		functorZipList = pkg_Data_List_Lazy_Types.Get_functorList()
	})
	return functorZipList
}

var foldableZipList gopurs_runtime.Value
var once_foldableZipList sync.Once
func Get_foldableZipList() gopurs_runtime.Value {
	once_foldableZipList.Do(func() {
		foldableZipList = pkg_Data_List_Lazy_Types.Get_foldableList()
	})
	return foldableZipList
}

var eqZipList gopurs_runtime.Value
var once_eqZipList sync.Once
func Get_eqZipList() gopurs_runtime.Value {
	once_eqZipList.Do(func() {
		eqZipList = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_eq1List().PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_0)})
})
	})
	return eqZipList
}

var applyZipList gopurs_runtime.Value
var once_applyZipList sync.Once
func Get_applyZipList() gopurs_runtime.Value {
	once_applyZipList.Do(func() {
		applyZipList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_zipWith(), pkg_Data_Function.Get_apply()), v_0), v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
})})
	})
	return applyZipList
}

var zipListIsNotBind gopurs_runtime.Value
var once_zipListIsNotBind sync.Once
func Get_zipListIsNotBind() gopurs_runtime.Value {
	once_zipListIsNotBind.Do(func() {
		zipListIsNotBind = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("bind: unreachable"))
})), "Apply0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyZipList()
})})
})
	})
	return zipListIsNotBind
}

var applicativeZipList gopurs_runtime.Value
var once_applicativeZipList sync.Once
func Get_applicativeZipList() gopurs_runtime.Value {
	once_applicativeZipList.Do(func() {
		applicativeZipList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_fix(), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": x_0, "value1": xs_1})
}))
}))
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyZipList()
})})
	})
	return applicativeZipList
}

var altZipList gopurs_runtime.Value
var once_altZipList sync.Once
func Get_altZipList() gopurs_runtime.Value {
	once_altZipList.Do(func() {
		altZipList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy_Types.Get_semigroupList().PtrVal.(map[string]gopurs_runtime.Value)["append"], v_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_drop(), gopurs_runtime.Apply(pkg_Data_List_Lazy.Get_length(), v_0)), v1_1))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Lazy_Types.Get_functorList()
})})
	})
	return altZipList
}

var plusZipList gopurs_runtime.Value
var once_plusZipList sync.Once
func Get_plusZipList() gopurs_runtime.Value {
	once_plusZipList.Do(func() {
		plusZipList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": pkg_Data_List_Lazy_Types.Get_nil(), "Alt0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altZipList()
})})
	})
	return plusZipList
}

var alternativeZipList gopurs_runtime.Value
var once_alternativeZipList sync.Once
func Get_alternativeZipList() gopurs_runtime.Value {
	once_alternativeZipList.Do(func() {
		alternativeZipList = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeZipList()
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusZipList()
})})
	})
	return alternativeZipList
}


