package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Component_Gen_genSecond gopurs_runtime.Value
var once_Data_Time_Component_Gen_genSecond sync.Once

func Get_Data_Time_Component_Gen_genSecond() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genSecond.Do(func() {
		cache_Data_Time_Component_Gen_genSecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Component_Gen_genSecond
}

var cache_Data_Time_Component_Gen_genMinute gopurs_runtime.Value
var once_Data_Time_Component_Gen_genMinute sync.Once

func Get_Data_Time_Component_Gen_genMinute() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genMinute.Do(func() {
		cache_Data_Time_Component_Gen_genMinute = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Component_Gen_genMinute
}

var cache_Data_Time_Component_Gen_genMillisecond gopurs_runtime.Value
var once_Data_Time_Component_Gen_genMillisecond sync.Once

func Get_Data_Time_Component_Gen_genMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genMillisecond.Do(func() {
		cache_Data_Time_Component_Gen_genMillisecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Component_Gen_genMillisecond
}

var cache_Data_Time_Component_Gen_genHour gopurs_runtime.Value
var once_Data_Time_Component_Gen_genHour sync.Once

func Get_Data_Time_Component_Gen_genHour() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genHour.Do(func() {
		cache_Data_Time_Component_Gen_genHour = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Component_Gen_genHour(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Component_Gen_genHour
}

func Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_Gen_1306125126_123048125(Rebox_Data_Time_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond()))))})
}

func Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_Gen_1306125126_123048125(Rebox_Data_Time_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute()))))})
}

func Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_Gen_1306125126_123048125(Rebox_Data_Time_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond()))))})
}

func Call_Data_Time_Component_Gen_genHour(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_Component_Gen_1306125126_123048125(Rebox_Data_Time_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour()))))})
}

func Rebox_Data_Time_Component_Gen_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Time_Component_Gen_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}
