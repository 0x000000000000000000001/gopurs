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
			return Call_Main_mkRecord(v_0_box, v1_1_box)
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
			a_prime__0_1 := gopurs_runtime.Bool(true)
			_ = a_prime__0_1
			obj_0_0 := gopurs_runtime.RecordDict1("value", gopurs_runtime.Bool((a_prime__0_1.IntVal) != (0)))
			_ = obj_0_0
			var __t3 string
			{
				if (gopurs_runtime.Bool((gopurs_runtime.RecordGet(obj_0_0, "value").IntVal) != (0)).IntVal) != (0) {
					__t3 = "true"
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = "false"
			}
		end_branch_3:
			_dollar___unused_1_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t3)), gopurs_runtime.Value{})
			_ = _dollar___unused_1_2
			a_prime__2_5 := gopurs_runtime.Float(2.0)
			_ = a_prime__2_5
			point_2_4 := gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Float(a_prime__2_5.FloatVal()), gopurs_runtime.Float(1.0))
			_ = point_2_4
			_dollar___unused_3_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.RecordGet(point_2_4, "x").FloatVal()) == (2.0))), gopurs_runtime.Value{})
			_ = _dollar___unused_3_6
			_dollar___unused_4_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.RecordGet(point_2_4, "y").FloatVal()) == (1.0))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_mkRecord(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	return gopurs_runtime.RecordDict3("bar", "baz", "foo", v1_1, gopurs_runtime.Str("baz"), v_0)
}
