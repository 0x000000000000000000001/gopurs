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
			__local_ref_1 := gopurs_runtime.Float(1.0)
			_ = __local_ref_1
			var __local_iface_1 interface{} = __local_ref_1
			n1_0_0 := gopurs_runtime.Any(&__local_iface_1)
			_ = n1_0_0
			__local_ref_3 := gopurs_runtime.Float(1.0)
			_ = __local_ref_3
			var __local_iface_3 interface{} = __local_ref_3
			n2_1_2 := gopurs_runtime.Any(&__local_iface_3)
			_ = n2_1_2
			_dollar___unused_2_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_while(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_2_5 := (*(n1_0_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
				_ = __local_var_2_5
				var __t_tag_6 gopurs_runtime.Value = Data_Ord_OrdNumberImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, 1000.0, __local_var_2_5.FloatVal())
				_ = __t_tag_6
				return gopurs_runtime.Bool((uint32(__t_tag_6.IntVal) == 380165415))
			}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				n1_prime__2_7 := (*(n1_0_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
				_ = n1_prime__2_7
				n2_prime__3_8 := (*(n2_1_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
				_ = n2_prime__3_8
				*(n2_1_2.PtrVal().(*interface{})) = gopurs_runtime.Float((n1_prime__2_7.FloatVal()) + (n2_prime__3_8.FloatVal()))
				_dollar___unused_4_9 := gopurs_runtime.Float((n1_prime__2_7.FloatVal()) + (n2_prime__3_8.FloatVal()))
				_ = _dollar___unused_4_9
				*(n1_0_0.PtrVal().(*interface{})) = n2_prime__3_8
				return n2_prime__3_8
			})), gopurs_runtime.Value{})
			_ = _dollar___unused_2_4
			return (*(n2_1_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
		})).FloatVal())
	})
	return cache_Main_fib
}
