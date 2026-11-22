package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_Component_Gen_compose gopurs_runtime.Value
var once_Data_Date_Component_Gen_compose sync.Once

func Get_Data_Date_Component_Gen_compose() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_compose.Do(func() {
		cache_Data_Date_Component_Gen_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Date_Component_Gen_compose
}

var cache_Data_Date_Component_Gen_fromJust gopurs_runtime.Value
var once_Data_Date_Component_Gen_fromJust sync.Once

func Get_Data_Date_Component_Gen_fromJust() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_fromJust.Do(func() {
		cache_Data_Date_Component_Gen_fromJust = Get_Data_Maybe_fromJust__3338668507()
	})
	return cache_Data_Date_Component_Gen_fromJust
}

var cache_Data_Date_Component_Gen_toEnum gopurs_runtime.Value
var once_Data_Date_Component_Gen_toEnum sync.Once

func Get_Data_Date_Component_Gen_toEnum() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_toEnum.Do(func() {
		cache_Data_Date_Component_Gen_toEnum = Rebox_Data_Date_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear())).V4
	})
	return cache_Data_Date_Component_Gen_toEnum
}

var cache_Data_Date_Component_Gen_genYear gopurs_runtime.Value
var once_Data_Date_Component_Gen_genYear sync.Once

func Get_Data_Date_Component_Gen_genYear() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genYear.Do(func() {
		cache_Data_Date_Component_Gen_genYear = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Date_Component_Gen_genYear(dictMonadGen_0_box)
		})
	})
	return cache_Data_Date_Component_Gen_genYear
}

var cache_Data_Date_Component_Gen_genWeekday gopurs_runtime.Value
var once_Data_Date_Component_Gen_genWeekday sync.Once

func Get_Data_Date_Component_Gen_genWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genWeekday.Do(func() {
		cache_Data_Date_Component_Gen_genWeekday = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Date_Component_Gen_genWeekday(dictMonadGen_0_box)
		})
	})
	return cache_Data_Date_Component_Gen_genWeekday
}

var cache_Data_Date_Component_Gen_genMonth gopurs_runtime.Value
var once_Data_Date_Component_Gen_genMonth sync.Once

func Get_Data_Date_Component_Gen_genMonth() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genMonth.Do(func() {
		cache_Data_Date_Component_Gen_genMonth = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Date_Component_Gen_genMonth(dictMonadGen_0_box)
		})
	})
	return cache_Data_Date_Component_Gen_genMonth
}

var cache_Data_Date_Component_Gen_genDay gopurs_runtime.Value
var once_Data_Date_Component_Gen_genDay sync.Once

func Get_Data_Date_Component_Gen_genDay() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genDay.Do(func() {
		cache_Data_Date_Component_Gen_genDay = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Date_Component_Gen_genDay(dictMonadGen_0_box)
		})
	})
	return cache_Data_Date_Component_Gen_genDay
}

func Call_Data_Date_Component_Gen_genYear(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__3338668507(), Rebox_Data_Date_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear())).V4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(1900)), gopurs_runtime.Int(int64(2100))))
}

func Call_Data_Date_Component_Gen_genWeekday(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_Gen_4021906832_123048125(Rebox_Data_Date_Component_Gen_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumWeekday()))))})
}

func Call_Data_Date_Component_Gen_genMonth(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_Gen_4021906832_123048125(Rebox_Data_Date_Component_Gen_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth()))))})
}

func Call_Data_Date_Component_Gen_genDay(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_Gen_1306125126_123048125(Rebox_Data_Date_Component_Gen_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay()))))})
}

func Rebox_Data_Date_Component_Gen_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Date_Component_Gen_123048125_4021906832(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Date_Component_Gen_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Date_Component_Gen_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}
