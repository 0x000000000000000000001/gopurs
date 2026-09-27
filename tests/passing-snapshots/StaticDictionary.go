package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_callUse gopurs_runtime.Value
var once_Main_callUse sync.Once

func Get_Main_callUse() gopurs_runtime.Value {
	once_Main_callUse.Do(func() {
		cache_Main_callUse = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_callUse(n_0_box.IntVal))
		})
	})
	return cache_Main_callUse
}

var cache_Main_callRun gopurs_runtime.Value
var once_Main_callRun sync.Once

func Get_Main_callRun() gopurs_runtime.Value {
	once_Main_callRun.Do(func() {
		cache_Main_callRun = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_callRun(n_0_box.IntVal))
		})
	})
	return cache_Main_callRun
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_callRun(int64(7)), int64(8)})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_callUse(int64(7)), int64(9)}), gopurs_runtime.Value{})
			_ = __local_var_2_2
			// TAST (Let): v_3_4 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
			v_3_4 := Call_Worker_runSummary(int64(7))
			_ = v_3_4
			var __t5 string
			{
				if v_3_4.V2 {
					__t5 = v_3_4.V1.name
					goto end_branch_5
				} else {

				}
			}
			{
				if !v_3_4.V2 {
					__t5 = ""
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() string { panic("Failed pattern match") }()
			}
		end_branch_5:
			__local_var_3_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___4176622598("", struct {
				actual   string
				expected string
			}{__t5, "alpha"}), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_callRun(int64(-1)), int64(-1)}), gopurs_runtime.Value{})
			_ = __local_var_4_6
			__local_var_5_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_callUse(int64(-1)), int64(-1)}), gopurs_runtime.Value{})
			_ = __local_var_5_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_callUse(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_useSummary(n_0)
	_ = v_1_0
	var __t1 int64
	{
		if !v_1_0.V2 {
			__t1 = int64(-1)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_1_0.V2 {
			__t1 = v_1_0.V1.id
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_callRun(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_runSummary(n_0)
	_ = v_1_0
	var __t1 int64
	{
		if !v_1_0.V2 {
			__t1 = int64(-1)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_1_0.V2 {
			__t1 = v_1_0.V1.id
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}
