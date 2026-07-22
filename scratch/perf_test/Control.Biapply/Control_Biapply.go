package Control_Biapply

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Function "gopurs/output/Data.Function"
)

var biapplyTuple gopurs_runtime.Value
var once_biapplyTuple sync.Once
func Get_biapplyTuple() gopurs_runtime.Value {
	once_biapplyTuple.Do(func() {
		biapplyTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"], v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
})})
	})
	return biapplyTuple
}

var biapply gopurs_runtime.Value
var once_biapply sync.Once
func Get_biapply() gopurs_runtime.Value {
	once_biapply.Do(func() {
		biapply = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"]
})
	})
	return biapply
}

var biapplyFirst gopurs_runtime.Value
var once_biapplyFirst sync.Once
func Get_biapplyFirst() gopurs_runtime.Value {
	once_biapplyFirst.Do(func() {
		biapplyFirst = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bimap"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]
})), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]
})), a_1)), b_2)
})
})
})
	})
	return biapplyFirst
}

var biapplySecond gopurs_runtime.Value
var once_biapplySecond sync.Once
func Get_biapplySecond() gopurs_runtime.Value {
	once_biapplySecond.Do(func() {
		biapplySecond = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bimap"], pkg_Data_Function.Get_const_()), pkg_Data_Function.Get_const_()), a_1)), b_2)
})
})
})
	})
	return biapplySecond
}

var bilift2 gopurs_runtime.Value
var once_bilift2 sync.Once
func Get_bilift2() gopurs_runtime.Value {
	once_bilift2.Do(func() {
		bilift2 = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_1), g_2), a_3)), b_4)
})
})
})
})
})
	})
	return bilift2
}

var bilift3 gopurs_runtime.Value
var once_bilift3 sync.Once
func Get_bilift3() gopurs_runtime.Value {
	once_bilift3.Do(func() {
		bilift3 = gopurs_runtime.Func(func(dictBiapply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["biapply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictBiapply_0.PtrVal.(map[string]gopurs_runtime.Value)["Bifunctor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bimap"], f_1), g_2), a_3)), b_4)), c_5)
})
})
})
})
})
})
	})
	return bilift3
}


