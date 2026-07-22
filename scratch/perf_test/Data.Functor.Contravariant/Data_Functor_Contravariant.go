package Data_Functor_Contravariant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Void "gopurs/output/Data.Void"
)

var contravariantConst gopurs_runtime.Value
var once_contravariantConst sync.Once
func Get_contravariantConst() gopurs_runtime.Value {
	once_contravariantConst.Do(func() {
		contravariantConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cmap": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
})})
	})
	return contravariantConst
}

var cmap gopurs_runtime.Value
var once_cmap sync.Once
func Get_cmap() gopurs_runtime.Value {
	once_cmap.Do(func() {
		cmap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["cmap"]
})
	})
	return cmap
}

var cmapFlipped gopurs_runtime.Value
var once_cmapFlipped sync.Once
func Get_cmapFlipped() gopurs_runtime.Value {
	once_cmapFlipped.Do(func() {
		cmapFlipped = gopurs_runtime.Func(func(dictContravariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictContravariant_0.PtrVal.(map[string]gopurs_runtime.Value)["cmap"], f_2), x_1)
})
})
})
	})
	return cmapFlipped
}

var coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		coerce = gopurs_runtime.Func(func(dictContravariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Void.Get_absurd()), gopurs_runtime.Apply(gopurs_runtime.Apply(dictContravariant_0.PtrVal.(map[string]gopurs_runtime.Value)["cmap"], pkg_Data_Void.Get_absurd()), a_2))
})
})
})
	})
	return coerce
}

var imapC gopurs_runtime.Value
var once_imapC sync.Once
func Get_imapC() gopurs_runtime.Value {
	once_imapC.Do(func() {
		imapC = gopurs_runtime.Func(func(dictContravariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictContravariant_0.PtrVal.(map[string]gopurs_runtime.Value)["cmap"], f_2)
})
})
})
	})
	return imapC
}


