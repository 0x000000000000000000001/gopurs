package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0_box), x_1_box))
		})
	})
	return cache_Main_test
}

var cache_Main_test__1683400371 gopurs_runtime.Value
var once_Main_test__1683400371 sync.Once

func Get_Main_test__1683400371() gopurs_runtime.Value {
	once_Main_test__1683400371.Do(func() {
		cache_Main_test__1683400371 = gopurs_runtime.Func(func(x_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_test__1683400371(x_unused_0_box))
		})
	})
	return cache_Main_test__1683400371
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_test(dictShow_0_loop *Constructor_Data_Show_Show[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) string {
	var dictShow_0 *Constructor_Data_Show_Show[gopurs_runtime.Value] = dictShow_0_loop
	_ = dictShow_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(dictShow_0.V0, x_1).StrVal()
}

func Call_Main_test__1683400371(x_unused_0_loop gopurs_runtime.Value) string {
test__1683400371:
	for {
		if false {
			continue test__1683400371
		}
		var x_unused_0 gopurs_runtime.Value = x_unused_0_loop
		_ = x_unused_0
		return "unit"
	}
}
