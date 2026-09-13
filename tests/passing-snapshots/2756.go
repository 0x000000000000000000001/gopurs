package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Id gopurs_runtime.Value
var once_Main_Id sync.Once

func Get_Main_Id() gopurs_runtime.Value {
	once_Main_Id.Do(func() {
		cache_Main_Id = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Id(x_0_box)
		})
	})
	return cache_Main_Id
}

var cache_Main_Id__468248410 gopurs_runtime.Value
var once_Main_Id__468248410 sync.Once

func Get_Main_Id__468248410() gopurs_runtime.Value {
	once_Main_Id__468248410.Do(func() {
		cache_Main_Id__468248410 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_Id__468248410(func() struct {
					pu gopurs_runtime.Value
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						pu gopurs_runtime.Value
					}{}
					clone.pu = gopurs_runtime.RecordGet(orig, "pu")
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict1("pu", orig.pu)
			}()
		})
	})
	return cache_Main_Id__468248410
}

var cache_Main_pu gopurs_runtime.Value
var once_Main_pu sync.Once

func Get_Main_pu() gopurs_runtime.Value {
	once_Main_pu.Do(func() {
		cache_Main_pu = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_pu(v_0_box)
		})
	})
	return cache_Main_pu
}

var cache_Main_sampleC gopurs_runtime.Value
var once_Main_sampleC sync.Once

func Get_Main_sampleC() gopurs_runtime.Value {
	once_Main_sampleC.Do(func() {
		cache_Main_sampleC = func() gopurs_runtime.Value {
			orig := struct {
				pu gopurs_runtime.Value
			}{Get_Main_pu()}
			_ = orig
			return gopurs_runtime.RecordDict1("pu", orig.pu)
		}()
	})
	return cache_Main_sampleC
}

var cache_Main_sampleIdC gopurs_runtime.Value
var once_Main_sampleIdC sync.Once

func Get_Main_sampleIdC() gopurs_runtime.Value {
	once_Main_sampleIdC.Do(func() {
		cache_Main_sampleIdC = func() gopurs_runtime.Value {
			orig := struct {
				pu gopurs_runtime.Value
			}{Get_Main_pu()}
			_ = orig
			return gopurs_runtime.RecordDict1("pu", orig.pu)
		}()
	})
	return cache_Main_sampleIdC
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_Id(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Id__468248410(x_0_loop struct {
	pu gopurs_runtime.Value
}) struct {
	pu gopurs_runtime.Value
} {
Id__468248410:
	for {
		if false {
			continue Id__468248410
		}
		var x_0 struct {
			pu gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_pu(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Unit_unit()
	})
}
