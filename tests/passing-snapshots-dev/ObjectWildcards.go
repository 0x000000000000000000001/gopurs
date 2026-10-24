package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_mkRecord gopurs_runtime.Value
var once_Main_mkRecord sync.Once

func Get_Main_mkRecord() gopurs_runtime.Value {
	once_Main_mkRecord.Do(func() {
		cache_Main_mkRecord = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_mkRecord(v_0_box, v1_1_box)
				_ = orig
				return gopurs_runtime.RecordDict3("bar", "baz", "foo", orig.bar, gopurs_runtime.Str(orig.baz), orig.foo)
			}()
		})
	})
	return cache_Main_mkRecord
}

var cache_Main_getValue gopurs_runtime.Value
var once_Main_getValue sync.Once

func Get_Main_getValue() gopurs_runtime.Value {
	once_Main_getValue.Do(func() {
		cache_Main_getValue = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})
	})
	return cache_Main_getValue
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply2(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict1("value", v_0)
			}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			}))
			_ = __local_var_0_0
			obj_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = obj_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.RecordGet(obj_1_1, "value").IntVal) != (0))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			point_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("x", "y", v_3, gopurs_runtime.Float(1.0))
			}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float(2.0)
			})), gopurs_runtime.Value{})
			_ = point_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.RecordGet(point_3_3, "x").FloatVal()) == (2.0))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.RecordGet(point_3_3, "y").FloatVal()) == (1.0))), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_mkRecord(gopurs_runtime.Float(1.0), gopurs_runtime.Str("Done")).bar.StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_mkRecord(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) struct {
	bar gopurs_runtime.Value
	baz string
	foo gopurs_runtime.Value
	/*gopurs:record-v1:3:006200610072,00620061007a,0066006f006f*/
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	return struct {
		bar gopurs_runtime.Value
		baz string
		foo gopurs_runtime.Value
		/*gopurs:record-v1:3:006200610072,00620061007a,0066006f006f*/
	}{v1_1, "baz", v_0}
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
