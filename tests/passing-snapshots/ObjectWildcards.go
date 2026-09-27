package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
			__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := struct {
						value bool
					}{(v_0.IntVal) != (0)}
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Bool(orig.value))
				}()
			}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})), gopurs_runtime.Value{})
			_ = __local_var_0_0
			var __t2 string
			{
				if (gopurs_runtime.RecordGet(__local_var_0_0, "value").IntVal) != (0) {
					__t2 = "true"
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = "false"
			}
		end_branch_2:
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t2)), gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					orig := struct {
						x float64
						y float64
					}{v_2.FloatVal(), 1.0}
					_ = orig
					return gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Float(orig.x), gopurs_runtime.Float(orig.y))
				}()
			}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float(2.0)
			})), gopurs_runtime.Value{})
			_ = __local_var_2_3
			__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.RecordGet(__local_var_2_3, "x").FloatVal()) == (2.0))), gopurs_runtime.Value{})
			_ = __local_var_3_4
			__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.RecordGet(__local_var_2_3, "y").FloatVal()) == (1.0))), gopurs_runtime.Value{})
			_ = __local_var_4_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_mkRecord(gopurs_runtime.Float(1.0), gopurs_runtime.Str("Done")).bar.StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_mkRecord(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) struct {
	bar gopurs_runtime.Value
	baz string
	foo gopurs_runtime.Value
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	return struct {
		bar gopurs_runtime.Value
		baz string
		foo gopurs_runtime.Value
	}{v1_1, "baz", v_0}
}
