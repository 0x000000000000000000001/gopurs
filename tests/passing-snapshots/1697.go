package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_y gopurs_runtime.Value
var once_Main_y sync.Once

func Get_Main_y() gopurs_runtime.Value {
	once_Main_y.Do(func() {
		cache_Main_y = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_y(dictMonad_0_box)
		})
	})
	return cache_Main_y
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_x(dictMonad_0_box)
		})
	})
	return cache_Main_x
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main__2 gopurs_runtime.Value
var once_Main__2 sync.Once

func Get_Main__2() gopurs_runtime.Value {
	once_Main__2.Do(func() {
		cache_Main__2 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main__2(a_0_box)
		})
	})
	return cache_Main__2
}

var cache_Main_wtf gopurs_runtime.Value
var once_Main_wtf sync.Once

func Get_Main_wtf() gopurs_runtime.Value {
	once_Main_wtf.Do(func() {
		cache_Main_wtf = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wtf(dictMonad_0_box)
		})
	})
	return cache_Main_wtf
}

var cache_Main__2__2674436160 gopurs_runtime.Value
var once_Main__2__2674436160 sync.Once

func Get_Main__2__2674436160() gopurs_runtime.Value {
	once_Main__2__2674436160.Do(func() {
		cache_Main__2__2674436160 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main__2__2674436160(a_0_box)
		})
	})
	return cache_Main__2__2674436160
}

func Call_Main_y(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit()), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
	}))
}

func Call_Main_x(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit()), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
	}))
}

func Call_Main__2(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return a_0
}

func Call_Main_wtf(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Applicative0_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
	Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_1_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit()), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
	}))
}

func Call_Main__2__2674436160(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return a_0
}
