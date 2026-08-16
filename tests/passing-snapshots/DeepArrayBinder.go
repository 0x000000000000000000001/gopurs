package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons)(nil))}
	})
	return cache_Main_Nil
}

var cache_Main_match2 gopurs_runtime.Value
var once_Main_match2 sync.Once

func Get_Main_match2() gopurs_runtime.Value {
	once_Main_match2.Do(func() {
		cache_Main_match2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_match2(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](v_0_box)))
		})
	})
	return cache_Main_match2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool((Call_Main_match2(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons{1, gopurs_runtime.Float(1.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(2.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(3.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(4.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(5.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(6.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(7.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(8.0), (&Constructor_Main_Cons{1, gopurs_runtime.Float(9.0), (*Constructor_Main_Cons)(nil)})})})})})})})})}))}))) == (100.0)))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Main_Cons
}

type Constructor_Main_Nil struct {
	Rc uint32
}

func Call_Main_match2(v_0_loop *Constructor_Main_Cons) float64 {
match2:
	for {
		if false {
			continue match2
		}
		var v_0 *Constructor_Main_Cons = v_0_loop
		_ = v_0
		var __t2 float64
		{
			var __t_and_1 bool = false
			if v_0 != nil {

				var __t_tag_0 *Constructor_Main_Cons = (v_0).V1
				__t_and_1 = (__t_tag_0 != nil)
			}
			if __t_and_1 {
				__t2 = (((v_0).V0.FloatVal()) * (((v_0).V1).V0.FloatVal())) + (Call_Main_match2(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(((v_0).V1).V1)})))
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = 0.0
		}
	end_branch_2:
		return __t2
	}
}
