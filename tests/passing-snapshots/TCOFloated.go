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
			// TAST (Let): eqRowCons2_0_0 shape=LitRecord expectedFromAst=*Constructor_Data_Eq_EqRecord actual=*Constructor_Data_Eq_EqRecord bindingType=(TypeApp (ADT ["Data","Eq","EqRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row)])
			eqRowCons2_0_0 := (&Constructor_Data_Eq_EqRecord{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.RecordGet(ra_1, "foo").IntVal) == (gopurs_runtime.RecordGet(rb_2, "foo").IntVal))
					})
				})
			})})
			_ = eqRowCons2_0_0
			// TAST (Let): __local_var_1_1 shape=LitRecord expectedFromAst=*Constructor_Data_Ord_OrdRecord actual=*Constructor_Data_Ord_OrdRecord bindingType=(TypeApp (ADT ["Data","Ord","OrdRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row)])
			__local_var_1_1 := (&Constructor_Data_Ord_OrdRecord{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(eqRowCons2_0_0)}
			}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): left_4_2 shape=App(Var) expectedFromAst=uint32 actual=uint32 bindingType=(ADT ["Data","Ordering","Ordering"] [])
						left_4_2 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.RecordGet(ra_2, "foo"), gopurs_runtime.RecordGet(rb_3, "foo")).IntVal)
						_ = left_4_2
						var __t3 uint32
						{
							if (left_4_2 == 902936544) != (true) {
								__t3 = left_4_2
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
			// TAST (Let): eqRec1_2_4 shape=LitRecord expectedFromAst=*Constructor_Data_Eq_Eq actual=*Constructor_Data_Eq_Eq bindingType=(ADT ["Data","Eq","Eq"] [(Record (Row [] (TypeVar row)))])
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
			// TAST (Let): left_1_0 shape=App(Var) expectedFromAst=uint32 actual=uint32 bindingType=(ADT ["Data","Ordering","Ordering"] [])
			left_1_0 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.RecordGet(x_0, "foo"), gopurs_runtime.Int(0)).IntVal)
			_ = left_1_0
			if (((left_1_0 == 902936544) != (true)) && (left_1_0 == 380165415)) != (true) {
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
