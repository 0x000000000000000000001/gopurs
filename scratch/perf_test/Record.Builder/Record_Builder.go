package Record_Builder

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Record_Unsafe_Union "gopurs/output/Record.Unsafe.Union"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), pkg_Record_Unsafe_Union.Get_unsafeUnionFn()), r1_1), r2_2)
})
})
})
	})
	return union
}

var semigroupoidBuilder gopurs_runtime.Value
var once_semigroupoidBuilder sync.Once
func Get_semigroupoidBuilder() gopurs_runtime.Value {
	once_semigroupoidBuilder.Do(func() {
		semigroupoidBuilder = pkg_Control_Semigroupoid.Get_semigroupoidFn()
	})
	return semigroupoidBuilder
}

var rename gopurs_runtime.Value
var once_rename sync.Once
func Get_rename() gopurs_runtime.Value {
	once_rename.Do(func() {
		rename = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeRename(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], l1_6)), gopurs_runtime.Apply(dictIsSymbol1_1.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], l2_7)), r1_8)
})
})
})
})
})
})
})
})
})
	})
	return rename
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
})
	})
	return nub
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeModify(), gopurs_runtime.Apply(dictIsSymbol_2.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], l_3)), f_4), r1_5)
})
})
})
})
})
})
	})
	return modify
}

var merge gopurs_runtime.Value
var once_merge sync.Once
func Get_merge() gopurs_runtime.Value {
	once_merge.Do(func() {
		merge = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), pkg_Record_Unsafe_Union.Get_unsafeUnionFn()), r1_2), r2_3)
})
})
})
})
	})
	return merge
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeInsert(), gopurs_runtime.Apply(dictIsSymbol_2.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], l_3)), a_4), r1_5)
})
})
})
})
})
})
	})
	return insert
}

var disjointUnion gopurs_runtime.Value
var once_disjointUnion sync.Once
func Get_disjointUnion() gopurs_runtime.Value {
	once_disjointUnion.Do(func() {
		disjointUnion = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), pkg_Record_Unsafe_Union.Get_unsafeUnionFn()), r1_2), r2_3)
})
})
})
})
	})
	return disjointUnion
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeDelete(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], l_3)), r2_4)
})
})
})
})
})
	})
	return delete_
}

var categoryBuilder gopurs_runtime.Value
var once_categoryBuilder sync.Once
func Get_categoryBuilder() gopurs_runtime.Value {
	once_categoryBuilder.Do(func() {
		categoryBuilder = pkg_Control_Category.Get_categoryFn()
	})
	return categoryBuilder
}

var build gopurs_runtime.Value
var once_build sync.Once
func Get_build() gopurs_runtime.Value {
	once_build.Do(func() {
		build = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(Get_copyRecord(), r1_1))
})
})
	})
	return build
}

var buildFromScratch gopurs_runtime.Value
var once_buildFromScratch sync.Once
func Get_buildFromScratch() gopurs_runtime.Value {
	once_buildFromScratch.Do(func() {
		buildFromScratch = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(a_0, gopurs_runtime.Apply(Get_copyRecord(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{})))
})
	})
	return buildFromScratch
}

var flip gopurs_runtime.Value
var once_flip sync.Once
func Get_flip() gopurs_runtime.Value {
	once_flip.Do(func() {
		flip = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(Get_copyRecord(), b_1))
})
})
})
	})
	return flip
}

func Get_copyRecord() gopurs_runtime.Value {
	return CopyRecord
}

func Get_unsafeDelete() gopurs_runtime.Value {
	return UnsafeDelete
}

func Get_unsafeInsert() gopurs_runtime.Value {
	return UnsafeInsert
}

func Get_unsafeModify() gopurs_runtime.Value {
	return UnsafeModify
}

func Get_unsafeRename() gopurs_runtime.Value {
	return UnsafeRename
}
