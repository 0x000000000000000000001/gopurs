package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ordRecord gopurs_runtime.Value
var once_Main_ordRecord sync.Once

func Get_Main_ordRecord() gopurs_runtime.Value {
	once_Main_ordRecord.Do(func() {
		cache_Main_ordRecord = func() gopurs_runtime.Value {
			// TAST (Let): eqRowCons2_0_0 -> *Constructor_Data_Eq_EqRecord
			eqRowCons2_0_0 := (&Constructor_Data_Eq_EqRecord{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.RecordGet(ra_1, "foo").IntVal) == (gopurs_runtime.RecordGet(rb_2, "foo").IntVal))
					})
				})
			})})
			_ = eqRowCons2_0_0
			// TAST (Let): __local_var_1_1 -> *Constructor_Data_Ord_OrdRecord
			__local_var_1_1 := (&Constructor_Data_Ord_OrdRecord{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(eqRowCons2_0_0)}
			}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): left_4_2 -> gopurs_runtime.Value
						left_4_2 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.RecordGet(ra_2, "foo"), gopurs_runtime.RecordGet(rb_3, "foo"))
						_ = left_4_2
						var __t3 uint32
						{
							if (uint32(left_4_2.IntVal) == 902936544) != (true) {
								__t3 = uint32(left_4_2.IntVal)
								goto end_branch_3
							} else {

							}
						}
						{
							__t3 = 902936544
						}
					end_branch_3:
						return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
					})
				})
			})})
			_ = __local_var_1_1
			// TAST (Let): eqRec1_2_4 -> *Constructor_Data_Eq_Eq
			eqRec1_2_4 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_1.V0), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
			_ = eqRec1_2_4
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqRec1_2_4)}
			}), gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_1.V1), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
		}()
	})
	return cache_Main_ordRecord
}

var cache_Main_looper gopurs_runtime.Value
var once_Main_looper sync.Once

func Get_Main_looper() gopurs_runtime.Value {
	once_Main_looper.Do(func() {
		cache_Main_looper = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_looper(x_0_box))
		})
	})
	return cache_Main_looper
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_looper(gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(100000)))))
	})
	return cache_Main_main
}

func Call_Main_looper(x_0_loop gopurs_runtime.Value) string {
looper:
	for {
		if false {
			continue looper
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		var __t1 string
		{
			// TAST (Let): left_1_0 -> gopurs_runtime.Value
			left_1_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.RecordGet(x_0, "foo"), gopurs_runtime.Int(0))
			_ = left_1_0
			if (((uint32(left_1_0.IntVal) == 902936544) != (true)) && (uint32(left_1_0.IntVal) == 380165415)) != (true) {
				__t1 = "Done"
				goto end_branch_1
			} else {

			}
		}
		{
			x_0_loop = gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int((gopurs_runtime.RecordGet(x_0, "foo").IntVal)-(1)))
			continue looper
			__t1 = gopurs_runtime.Value{}.StrVal()
		}
	end_branch_1:
		return __t1
	}
}
