package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = Call_Data_Semiring_add(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))
	})
	return cache_Main_f
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(a_0_box.IntVal, b_1_box.IntVal))
		})
	})
	return cache_Main_g
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if (gopurs_runtime.Apply2(Call_Data_Semiring_add(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt())))), gopurs_runtime.Int(int64(10)), gopurs_runtime.Int(int64(5))).IntVal) == (int64(15)) {
					__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Failed"))
			}
		end_branch_0:
			return __t0
		}()
	})
	return cache_Main_main
}

func Call_Main_g(a_0_loop int64, b_1_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	return gopurs_runtime.Apply2(Call_Data_Semiring_add(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt())))), gopurs_runtime.Int(a_0), gopurs_runtime.Int(b_1)).IntVal
}

func Rebox_Main_2826095630_348932501(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.IntVal
	out.V3 = in.V3.IntVal
	return out
}

func Rebox_Main_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Int(in.V2)
	out.V3 = gopurs_runtime.Int(in.V3)
	return out
}
