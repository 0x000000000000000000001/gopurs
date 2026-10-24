package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if Call_Data_Ord_between(Rebox_Main_3308271157_4177771502(Rebox_Main_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))) {
					__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail"))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}
		end_branch_0:
			return __t0
		}()
	})
	return cache_Main_main
}

func Rebox_Main_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}
