package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_DateTime_Gen_genDateTime gopurs_runtime.Value
var once_Data_DateTime_Gen_genDateTime sync.Once

func Get_Data_DateTime_Gen_genDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Gen_genDateTime.Do(func() {
		cache_Data_DateTime_Gen_genDateTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_DateTime_Gen_genDateTime(dictMonadGen_0_box)
		})
	})
	return cache_Data_DateTime_Gen_genDateTime
}

func Call_Data_DateTime_Gen_genDateTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=Any
	Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
	_ = Bind1_1_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_DateTime_DateTime(), Call_Data_Date_Gen_genDate(dictMonadGen_0)), Call_Data_Time_Gen_genTime(dictMonadGen_0))
}
