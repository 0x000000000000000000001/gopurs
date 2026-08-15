package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_partialTCO gopurs_runtime.Value
var once_Main_partialTCO sync.Once

func Get_Main_partialTCO() gopurs_runtime.Value {
	once_Main_partialTCO.Do(func() {
		cache_Main_partialTCO = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_partialTCO(_dollar___unused_0_box, (v_1_box.IntVal) != (0), v1_2_box.IntVal))
		})
	})
	return cache_Main_partialTCO
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_partialTCO(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop bool, v1_2_loop int64) int64 {
partialTCO:
	for {
		if false {
			continue partialTCO
		}
		var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
		_ = _dollar___unused_0
		var v_1 bool = v_1_loop
		_ = v_1
		var v1_2 int64 = v1_2_loop
		_ = v1_2
		var __t1 int64
		{
			if v_1 {
				var __t0 int64
				{
					if (v1_2) == (0) {
						__t0 = 0
						goto end_branch_0
					} else {

					}
				}
				{
					_dollar___unused_0_loop = gopurs_runtime.Value{}
					v_1_loop = true
					v1_2_loop = (v1_2) - (1)
					continue partialTCO
					__t0 = gopurs_runtime.Value{}.IntVal
				}
			end_branch_0:
				__t1 = __t0
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
		}
	end_branch_1:
		return gopurs_runtime.Int(__t1).IntVal
	}
}
