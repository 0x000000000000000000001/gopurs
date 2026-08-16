package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testMonad gopurs_runtime.Value
var once_Main_testMonad sync.Once

func Get_Main_testMonad() gopurs_runtime.Value {
	once_Main_testMonad.Do(func() {
		cache_Main_testMonad = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testMonad(dictMonad_0_box)
		})
	})
	return cache_Main_testMonad
}

var cache_Main_testIMonad gopurs_runtime.Value
var once_Main_testIMonad sync.Once

func Get_Main_testIMonad() gopurs_runtime.Value {
	once_Main_testIMonad.Do(func() {
		cache_Main_testIMonad = gopurs_runtime.Func(func(dictIxMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testIMonad(dictIxMonad_0_box)
		})
	})
	return cache_Main_testIMonad
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_testMonad(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) expectedFromAst=*Constructor_Control_Bind_Bind actual=*Constructor_Control_Bind_Bind bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_1_0
	// TAST (Let): Applicative0_2_1 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
	Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_2_1
	return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Str("test")), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Str("test")), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Str((a_3.StrVal())+(b_4.StrVal())))
		}))
	}))
}

func Call_Main_testIMonad(dictIxMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIxMonad_0 gopurs_runtime.Value = dictIxMonad_0_loop
	_ = dictIxMonad_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictIxMonad_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIxMonad_0, "pure"), gopurs_runtime.Str("test")), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictIxMonad_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIxMonad_0, "pure"), gopurs_runtime.Str("test")), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIxMonad_0, "pure"), gopurs_runtime.Str((a_1.StrVal())+(b_2.StrVal())))
		}))
	}))
}
