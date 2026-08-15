package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_null gopurs_runtime.Value
var once_Main_null sync.Once

func Get_Main_null() gopurs_runtime.Value {
	once_Main_null.Do(func() {
		cache_Main_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_null(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_null
}

var cache_Main_comparing gopurs_runtime.Value
var once_Main_comparing sync.Once

func Get_Main_comparing() gopurs_runtime.Value {
	once_Main_comparing.Do(func() {
		cache_Main_comparing = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_comparing(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_comparing
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordBooleanImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)).IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t4 string
			{
				var __t_tag_1 uint32 = uint32(Get_Main_test().IntVal)
				if uint32(__t_tag_1) == 1527465420 {
					__t4 = "LT"
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(Get_Main_test().IntVal)
				if uint32(__t_tag_2) == 380165415 {
					__t4 = "GT"
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(Get_Main_test().IntVal)
				if uint32(__t_tag_3) == 902936544 {
					__t4 = "EQ"
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
			}
		end_branch_4:
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t4))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_5 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_5
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

var cache_Main_comparing__3506074860 gopurs_runtime.Value
var once_Main_comparing__3506074860 sync.Once

func Get_Main_comparing__3506074860() gopurs_runtime.Value {
	once_Main_comparing__3506074860.Do(func() {
		cache_Main_comparing__3506074860 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_comparing__3506074860(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_comparing__3506074860
}

func Call_Main_null(v_0_loop []gopurs_runtime.Value) bool {
	var v_0 []gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (0)
}

func Call_Main_comparing(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
	var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
	_ = dictOrd_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	var y_3 gopurs_runtime.Value = y_3_loop
	_ = y_3
	return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_Main_comparing__3506074860(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
	var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
	_ = dictOrd_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	var y_3 gopurs_runtime.Value = y_3_loop
	_ = y_3
	return uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}
