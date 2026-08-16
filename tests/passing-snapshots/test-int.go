package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=*Constructor_Data_Maybe_Just actual=*Constructor_Data_Maybe_Just bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
			__local_var_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(gopurs_runtime.Float(-2147483648.0).FloatVal())))
			_ = __local_var_0_0
			var __t1 gopurs_runtime.Value
			{
				if __local_var_0_0 != nil {
					__t1 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("Just ")+(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((__local_var_0_0).V0.IntVal)).StrVal())))
					goto end_branch_1
				} else {

				}
			}
			{
				if __local_var_0_0 == nil {
					__t1 = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Nothing"))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_1:
			return __t1
		}()
	})
	return cache_Main_main
}
