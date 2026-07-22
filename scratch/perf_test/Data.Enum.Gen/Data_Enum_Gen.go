package Data_Enum_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
)

var foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		foldable1NonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())
	})
	return foldable1NonEmpty
}

var genBoundedEnum gopurs_runtime.Value
var once_genBoundedEnum sync.Once
func Get_genBoundedEnum() gopurs_runtime.Value {
	once_genBoundedEnum.Do(func() {
		genBoundedEnum = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
elements_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_Gen.Get_elements(), dictMonadGen_0), Get_foldable1NonEmpty())
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
Enum1_3_1 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["Enum1"], gopurs_runtime.Value{})
Bounded0_4_2 := gopurs_runtime.Apply(dictBoundedEnum_2.PtrVal.(map[string]gopurs_runtime.Value)["Bounded0"], gopurs_runtime.Value{})
v_5_3 := gopurs_runtime.Apply(Enum1_3_1.PtrVal.(map[string]gopurs_runtime.Value)["succ"], Bounded0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["bottom"])
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(elements_1_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": Bounded0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["bottom"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum.Get_enumFromTo(), Enum1_3_1), pkg_Data_Unfoldable1.Get_unfoldable1Array()), v_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), Bounded0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["top"])}))
} else {
if (gopurs_runtime.Bool(v_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], Bounded0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["bottom"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
})
	})
	return genBoundedEnum
}


