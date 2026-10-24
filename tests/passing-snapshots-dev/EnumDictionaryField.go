package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Fast gopurs_runtime.Value
var once_Main_Fast sync.Once

func Get_Main_Fast() gopurs_runtime.Value {
	once_Main_Fast.Do(func() {
		cache_Main_Fast = gopurs_runtime.Value{Type: 9, IntVal: int64(4067170750), UnsafePtr: nil}
	})
	return cache_Main_Fast
}

var cache_Main_Safe gopurs_runtime.Value
var once_Main_Safe sync.Once

func Get_Main_Safe() gopurs_runtime.Value {
	once_Main_Safe.Do(func() {
		cache_Main_Safe = gopurs_runtime.Value{Type: 9, IntVal: int64(1097047119), UnsafePtr: nil}
	})
	return cache_Main_Safe
}

var cache_Main_eqPriority gopurs_runtime.Value
var once_Main_eqPriority sync.Once

func Get_Main_eqPriority() gopurs_runtime.Value {
	once_Main_eqPriority.Do(func() {
		cache_Main_eqPriority = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 bool
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 4067170750 {
					var __t_tag_4 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_4
					__t5 = (uint32(__t_tag_4) == 4067170750)
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 1097047119 {

					var __t_tag_1 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 1097047119)
				}
				__t5 = __t_and_2
			}
		end_branch_5:
			return gopurs_runtime.Bool(__t5)
		})})))}
	})
	return cache_Main_eqPriority
}

var cache_Main_protectedString gopurs_runtime.Value
var once_Main_protectedString sync.Once

func Get_Main_protectedString() gopurs_runtime.Value {
	once_Main_protectedString.Do(func() {
		cache_Main_protectedString = gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_2202314751_48729792((&Constructor_Main_Protected[string]{1, 4067170750})))}
	})
	return cache_Main_protectedString
}

var cache_Main_protectedInt gopurs_runtime.Value
var once_Main_protectedInt sync.Once

func Get_Main_protectedInt() gopurs_runtime.Value {
	once_Main_protectedInt.Do(func() {
		cache_Main_protectedInt = gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_3555759771_48729792((&Constructor_Main_Protected[int64]{1, 1097047119})))}
	})
	return cache_Main_protectedInt
}

var cache_Main_priority gopurs_runtime.Value
var once_Main_priority sync.Once

func Get_Main_priority() gopurs_runtime.Value {
	once_Main_priority.Do(func() {
		cache_Main_priority = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_priority(dict_0_box)
		})
	})
	return cache_Main_priority
}

var cache_Main_isSafe gopurs_runtime.Value
var once_Main_isSafe sync.Once

func Get_Main_isSafe() gopurs_runtime.Value {
	once_Main_isSafe.Do(func() {
		cache_Main_isSafe = gopurs_runtime.Func(func(dictProtected_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_isSafe(dictProtected_0_box)
		})
	})
	return cache_Main_isSafe
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(Call_Main_isSafe(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_3555759771_48729792(Rebox_Main_48729792_3555759771(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](Get_Main_protectedInt()))))}).IntVal) != (0), true}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}())
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{(Call_Main_isSafe(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_2202314751_48729792(Rebox_Main_48729792_2202314751(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](Get_Main_protectedString()))))}).IntVal) != (0), false}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Fast struct {
	Rc uint32
}

type Constructor_Main_Safe struct {
	Rc uint32
}

type Constructor_Main_Protected[T_a any] struct {
	Rc uint32
	V0 uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00500072006f007400650063007400650064
func init() {
	gopurs_runtime.StructGetters[2242194644] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Protected[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "priority":
			return gopurs_runtime.Value{Type: 9, IntVal: int64(c.V0)}
		default:
			panic("Key not found in dictionary Constructor_Main_Protected: " + key)
		}
	}
}

func Call_Main_priority(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(dict_0, "priority").IntVal)), UnsafePtr: nil}
}

func Call_Main_isSafe(dictProtected_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProtected_0 gopurs_runtime.Value = dictProtected_0_loop
	_ = dictProtected_0
	var __t_tag_0 uint32 = uint32(Call_Main_priority(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](dictProtected_0))}).IntVal)
	_ = __t_tag_0
	return gopurs_runtime.Bool((uint32(__t_tag_0) == 1097047119))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_2202314751_48729792(in *Constructor_Main_Protected[string]) *Constructor_Main_Protected[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Protected[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3555759771_48729792(in *Constructor_Main_Protected[int64]) *Constructor_Main_Protected[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Protected[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_48729792_2202314751(in *Constructor_Main_Protected[gopurs_runtime.Value]) *Constructor_Main_Protected[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Protected[string])(unsafe.Pointer(in))
}

func Rebox_Main_48729792_3555759771(in *Constructor_Main_Protected[gopurs_runtime.Value]) *Constructor_Main_Protected[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Protected[int64])(unsafe.Pointer(in))
}
