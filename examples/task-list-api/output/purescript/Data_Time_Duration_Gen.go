package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Time_Duration_Gen_genSeconds gopurs_runtime.Value
var once_Data_Time_Duration_Gen_genSeconds sync.Once

func Get_Data_Time_Duration_Gen_genSeconds() gopurs_runtime.Value {
	once_Data_Time_Duration_Gen_genSeconds.Do(func() {
		cache_Data_Time_Duration_Gen_genSeconds = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_Gen_genSeconds(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Duration_Gen_genSeconds
}

var cache_Data_Time_Duration_Gen_genMinutes gopurs_runtime.Value
var once_Data_Time_Duration_Gen_genMinutes sync.Once

func Get_Data_Time_Duration_Gen_genMinutes() gopurs_runtime.Value {
	once_Data_Time_Duration_Gen_genMinutes.Do(func() {
		cache_Data_Time_Duration_Gen_genMinutes = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_Gen_genMinutes(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Duration_Gen_genMinutes
}

var cache_Data_Time_Duration_Gen_genMilliseconds gopurs_runtime.Value
var once_Data_Time_Duration_Gen_genMilliseconds sync.Once

func Get_Data_Time_Duration_Gen_genMilliseconds() gopurs_runtime.Value {
	once_Data_Time_Duration_Gen_genMilliseconds.Do(func() {
		cache_Data_Time_Duration_Gen_genMilliseconds = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_Gen_genMilliseconds(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Duration_Gen_genMilliseconds
}

var cache_Data_Time_Duration_Gen_genHours gopurs_runtime.Value
var once_Data_Time_Duration_Gen_genHours sync.Once

func Get_Data_Time_Duration_Gen_genHours() gopurs_runtime.Value {
	once_Data_Time_Duration_Gen_genHours.Do(func() {
		cache_Data_Time_Duration_Gen_genHours = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_Gen_genHours(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Duration_Gen_genHours
}

var cache_Data_Time_Duration_Gen_genDays gopurs_runtime.Value
var once_Data_Time_Duration_Gen_genDays sync.Once

func Get_Data_Time_Duration_Gen_genDays() gopurs_runtime.Value {
	once_Data_Time_Duration_Gen_genDays.Do(func() {
		cache_Data_Time_Duration_Gen_genDays = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Time_Duration_Gen_genDays(dictMonadGen_0_box)
		})
	})
	return cache_Data_Time_Duration_Gen_genDays
}

func Call_Data_Time_Duration_Gen_genSeconds(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(600.0)))
}

func Call_Data_Time_Duration_Gen_genMinutes(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(600.0)))
}

func Call_Data_Time_Duration_Gen_genMilliseconds(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(600000.0)))
}

func Call_Data_Time_Duration_Gen_genHours(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(240.0)))
}

func Call_Data_Time_Duration_Gen_genDays(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseFloat"), gopurs_runtime.Float(0.0), gopurs_runtime.Float(42.0)))
}
