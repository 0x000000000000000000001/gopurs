package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Time_Gen_genTime gopurs_runtime.Value
var once_Data_Time_Gen_genTime sync.Once

func Get_Data_Time_Gen_genTime() gopurs_runtime.Value {
	once_Data_Time_Gen_genTime.Do(func() {
		cache_Data_Time_Gen_genTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Gen_genTime(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Gen_genTime
}

func Call_Data_Time_Gen_genTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=Any
	Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
	_ = Bind1_1_0
	// TAST (Let): Apply0_2_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope1)])
	Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}))
	_ = Apply0_2_1
	return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Time_Time(), Call_Data_Time_Component_Gen_genHour(dictMonadGen_0)), Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0)), Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0)), Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0))
}
