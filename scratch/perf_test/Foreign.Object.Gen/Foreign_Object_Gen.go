package Foreign_Object_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Foreign_Object "gopurs/output/Foreign.Object"
)

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Apply(pkg_Foreign_Object.Get_fromFoldable(), pkg_Data_List_Types.Get_foldableList())
	})
	return fromFoldable
}

var genForeignObject gopurs_runtime.Value
var once_genForeignObject sync.Once
func Get_genForeignObject() gopurs_runtime.Value {
	once_genForeignObject.Do(func() {
		genForeignObject = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
Apply0_3_1 := gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_4_2 := gopurs_runtime.Apply(Apply0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
unfoldable1_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0), dictMonadGen_1), pkg_Data_List_Types.Get_unfoldableList())
return gopurs_runtime.Func(func(genKey_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genValue_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["sized"], gopurs_runtime.Func(func(size_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(0)), size_8)), gopurs_runtime.Func(func(newSize_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["resize"], gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_9
})), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_fromFoldable()), gopurs_runtime.Apply(unfoldable1_5_3, gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_Tuple()), genKey_6)), genValue_7))))
}))
}))
})
})
})
})
	})
	return genForeignObject
}


