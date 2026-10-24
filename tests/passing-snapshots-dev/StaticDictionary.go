package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
			__local_var_0_0 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_callRun(int64(7)), int64(8)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}())
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_callUse(int64(7)), int64(9)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
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
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{__t5, "alpha"}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_callRun(int64(-1)), int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_4_6
			_dollar___unused_5_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_callUse(int64(-1)), int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_5_7
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

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}
