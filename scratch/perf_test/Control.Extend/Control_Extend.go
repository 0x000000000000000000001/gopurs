package Control_Extend

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Functor "gopurs/output/Data.Functor"
)

var extendFn gopurs_runtime.Value
var once_extendFn sync.Once
func Get_extendFn() gopurs_runtime.Value {
	once_extendFn.Do(func() {
		extendFn = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Func(func(w_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], w_3), w_prime_4))
}))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
})})
})
	})
	return extendFn
}

var extendArray gopurs_runtime.Value
var once_extendArray sync.Once
func Get_extendArray() gopurs_runtime.Value {
	once_extendArray.Do(func() {
		extendArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": Get_arrayExtend(), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
})})
	})
	return extendArray
}

var extend gopurs_runtime.Value
var once_extend sync.Once
func Get_extend() gopurs_runtime.Value {
	once_extend.Do(func() {
		extend = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"]
})
	})
	return extend
}

var extendFlipped gopurs_runtime.Value
var once_extendFlipped sync.Once
func Get_extendFlipped() gopurs_runtime.Value {
	once_extendFlipped.Do(func() {
		extendFlipped = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], f_2), w_1)
})
})
})
	})
	return extendFlipped
}

var duplicate gopurs_runtime.Value
var once_duplicate sync.Once
func Get_duplicate() gopurs_runtime.Value {
	once_duplicate.Do(func() {
		duplicate = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return duplicate
}

var composeCoKleisliFlipped gopurs_runtime.Value
var once_composeCoKleisliFlipped sync.Once
func Get_composeCoKleisliFlipped() gopurs_runtime.Value {
	once_composeCoKleisliFlipped.Do(func() {
		composeCoKleisliFlipped = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], g_2), w_3))
})
})
})
})
	})
	return composeCoKleisliFlipped
}

var composeCoKleisli gopurs_runtime.Value
var once_composeCoKleisli sync.Once
func Get_composeCoKleisli() gopurs_runtime.Value {
	once_composeCoKleisli.Do(func() {
		composeCoKleisli = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], f_1), w_3))
})
})
})
})
	})
	return composeCoKleisli
}

func Get_arrayExtend() gopurs_runtime.Value {
	return ArrayExtend
}
