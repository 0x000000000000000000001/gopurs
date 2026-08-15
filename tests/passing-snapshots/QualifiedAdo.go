package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testIApplicative gopurs_runtime.Value
var once_Main_testIApplicative sync.Once

func Get_Main_testIApplicative() gopurs_runtime.Value {
	once_Main_testIApplicative.Do(func() {
		cache_Main_testIApplicative = gopurs_runtime.Func(func(dictIxApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testIApplicative(dictIxApplicative_0_box)
		})
	})
	return cache_Main_testIApplicative
}

var cache_Main_testApplicative gopurs_runtime.Value
var once_Main_testApplicative sync.Once

func Get_Main_testApplicative() gopurs_runtime.Value {
	once_Main_testApplicative.Do(func() {
		cache_Main_testApplicative = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_testApplicative(dictApplicative_0_box)
		})
	})
	return cache_Main_testApplicative
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_testIApplicative(dictIxApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIxApplicative_0 gopurs_runtime.Value = dictIxApplicative_0_loop
	_ = dictIxApplicative_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictIxApplicative_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIxApplicative_0, "IxFunctor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str((v_1.StrVal()) + (v1_2.StrVal()))
		})
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIxApplicative_0, "pure"), gopurs_runtime.Str("test"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIxApplicative_0, "pure"), gopurs_runtime.Str("test")))
}

func Call_Main_testApplicative(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str((v_1.StrVal()) + (v1_2.StrVal()))
		})
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Str("test"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Str("test")))
}
