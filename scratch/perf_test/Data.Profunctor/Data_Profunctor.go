package Data_Profunctor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var profunctorFn gopurs_runtime.Value
var once_profunctorFn sync.Once
func Get_profunctorFn() gopurs_runtime.Value {
	once_profunctorFn.Do(func() {
		profunctorFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
})})
	})
	return profunctorFn
}

var dimap gopurs_runtime.Value
var once_dimap sync.Once
func Get_dimap() gopurs_runtime.Value {
	once_dimap.Do(func() {
		dimap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"]
})
	})
	return dimap
}

var lcmap gopurs_runtime.Value
var once_lcmap sync.Once
func Get_lcmap() gopurs_runtime.Value {
	once_lcmap.Do(func() {
		lcmap = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], a2b_1), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
})
	})
	return lcmap
}

var rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		rmap = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), b2c_1)
})
})
	})
	return rmap
}

var unwrapIso gopurs_runtime.Value
var once_unwrapIso sync.Once
func Get_unwrapIso() gopurs_runtime.Value {
	once_unwrapIso.Do(func() {
		unwrapIso = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], pkg_Unsafe_Coerce.Get_unsafeCoerce()), pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
})
	})
	return unwrapIso
}

var wrapIso gopurs_runtime.Value
var once_wrapIso sync.Once
func Get_wrapIso() gopurs_runtime.Value {
	once_wrapIso.Do(func() {
		wrapIso = gopurs_runtime.Func(func(dictProfunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], pkg_Unsafe_Coerce.Get_unsafeCoerce()), pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
})
})
	})
	return wrapIso
}

var arr gopurs_runtime.Value
var once_arr sync.Once
func Get_arr() gopurs_runtime.Value {
	once_arr.Do(func() {
		arr = gopurs_runtime.Func(func(dictCategory_0 gopurs_runtime.Value) gopurs_runtime.Value {
identity1_1_0 := dictCategory_0.PtrVal.(map[string]gopurs_runtime.Value)["identity"]
return gopurs_runtime.Func(func(dictProfunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictProfunctor_2.PtrVal.(map[string]gopurs_runtime.Value)["dimap"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), f_3), identity1_1_0)
})
})
})
	})
	return arr
}


