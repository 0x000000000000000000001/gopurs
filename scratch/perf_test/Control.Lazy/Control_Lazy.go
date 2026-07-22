package Control_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var lazyUnit gopurs_runtime.Value
var once_lazyUnit sync.Once
func Get_lazyUnit() gopurs_runtime.Value {
	once_lazyUnit.Do(func() {
		lazyUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"defer": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})})
	})
	return lazyUnit
}

var lazyFn gopurs_runtime.Value
var once_lazyFn sync.Once
func Get_lazyFn() gopurs_runtime.Value {
	once_lazyFn.Do(func() {
		lazyFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"defer": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), x_1)
})
})})
	})
	return lazyFn
}

var defer_ gopurs_runtime.Value
var once_defer_ sync.Once
func Get_defer_() gopurs_runtime.Value {
	once_defer_.Do(func() {
		defer_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["defer"]
})
	})
	return defer_
}

var fix gopurs_runtime.Value
var once_fix sync.Once
func Get_fix() gopurs_runtime.Value {
	once_fix.Do(func() {
		fix = gopurs_runtime.Func(func(dictLazy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Apply(dictLazy_0.PtrVal.(map[string]gopurs_runtime.Value)["defer"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__2_0)
}))
return go__2_0
})
})
	})
	return fix
}


