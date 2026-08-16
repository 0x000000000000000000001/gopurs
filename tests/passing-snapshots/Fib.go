package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fib gopurs_runtime.Value
var once_Main_fib sync.Once

func Get_Main_fib() gopurs_runtime.Value {
	once_Main_fib.Do(func() {
		cache_Main_fib = gopurs_runtime.Float(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Float(1.0))
			_ = __local_var_0_0
			n1_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = n1_1_1
			n2_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Float(1.0)), gopurs_runtime.Value{})
			_ = n2_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_3_4 := (*(n1_1_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
				_ = __local_var_3_4
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1000.0), __local_var_3_4)
				return gopurs_runtime.Bool((uint32(__t_tag_5.IntVal) == 380165415))
			}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				n1_prime__3_6 := (*(n1_1_1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
				_ = n1_prime__3_6
				n2_prime__4_7 := (*(n2_2_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
				_ = n2_prime__4_7
				*(n2_2_2.PtrVal().(*interface{})) = gopurs_runtime.Float((n1_prime__3_6.FloatVal()) + (n2_prime__4_7.FloatVal()))
				_dollar___unused_5_8 := gopurs_runtime.Float((n1_prime__3_6.FloatVal()) + (n2_prime__4_7.FloatVal()))
				_ = _dollar___unused_5_8
				*(n1_1_1.PtrVal().(*interface{})) = n2_prime__4_7
				return n2_prime__4_7
			})), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			return (*(n2_2_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
		})).FloatVal())
	})
	return cache_Main_fib
}
