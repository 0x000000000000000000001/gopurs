package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_snd gopurs_runtime.Value
var once_Main_snd sync.Once

func Get_Main_snd() gopurs_runtime.Value {
	once_Main_snd.Do(func() {
		cache_Main_snd = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snd(_dollar___unused_0_box, func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v_1_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})
	})
	return cache_Main_snd
}

var cache_Main_snd__3188547773 gopurs_runtime.Value
var once_Main_snd__3188547773 sync.Once

func Get_Main_snd__3188547773() gopurs_runtime.Value {
	once_Main_snd__3188547773.Do(func() {
		cache_Main_snd__3188547773 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_snd__3188547773(func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_snd__3188547773
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result from 'snd'."), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

func Call_Main_snd(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 []gopurs_runtime.Value = v_1_loop
	_ = v_1
	var __t0 gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(v_1))).IntVal) == (int64(2)) {
			__t0 = gopurs_runtime.ArrayAccess(gopurs_runtime.Array(v_1), 1)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_snd__3188547773(v_0_loop []float64) float64 {
snd__3188547773:
	for {
		if false {
			continue snd__3188547773
		}
		var v_0 []float64 = v_0_loop
		_ = v_0
		var __t0 float64
		{
			if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(2)) {
				__t0 = gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Float(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), 1).FloatVal()
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() float64 { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}
}
