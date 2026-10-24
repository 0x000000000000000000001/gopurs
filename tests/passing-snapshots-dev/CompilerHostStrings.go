package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func(func(expected_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(expected_0_box.StrVal())
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_check("")
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Main_check("quotes\" slash\\"), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Main_check("\x00\x08\x09\x0a\x0d\x1f\x7f"), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(Call_Main_check("é漢€💻𝄞  "), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(Call_Main_check("𝌆"), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(Call_Main_check("\xed\xa0\x80"), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Main_check("\xed\xbf\xbf"), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(Call_Main_check("a\xed\xa0\x80z"), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(Call_Main_check("\xed\xa0\x80𐀀\xed\xbf\xbf"), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Main_check("math.Abs(1); unsafe.Pointer(nil); sync.Once{}; gopurs_runtime.Value{}"), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Main_check("/* sync.Once */ // unsafe.Pointer\x0a\"math.Abs\""), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_check(expected_0_loop string) gopurs_runtime.Value {
	var expected_0 string = expected_0_loop
	_ = expected_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Str(expected_0))
		_ = __local_var_1_0
		reference_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = reference_2_1
		actual_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), reference_2_1), gopurs_runtime.Value{})
		_ = actual_3_2
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{actual_3_2.StrVal(), expected_0}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
	})
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

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}
