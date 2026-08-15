package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_test(x_0_box))
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_isDesc gopurs_runtime.Value
var once_Main_isDesc sync.Once

func Get_Main_isDesc() gopurs_runtime.Value {
	once_Main_isDesc.Do(func() {
		cache_Main_isDesc = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_isDesc(func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_isDesc
}

var cache_Main_h gopurs_runtime.Value
var once_Main_h sync.Once

func Get_Main_h() gopurs_runtime.Value {
	once_Main_h.Do(func() {
		cache_Main_h = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_h(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(o_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_h
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_f(o_0_box))
		})
	})
	return cache_Main_f
}

func Call_Main_test(x_0_loop gopurs_runtime.Value) bool {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var __t0 bool
	{
		if (gopurs_runtime.RecordGet(x_0, "bool").IntVal) != (0) {
			__t0 = ((gopurs_runtime.RecordGet(x_0, "str").StrVal()) == ("Foo")) || (((gopurs_runtime.RecordGet(x_0, "str").StrVal()) == ("Bar")) && ((gopurs_runtime.RecordGet(x_0, "bool").IntVal) != (0)))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = ((gopurs_runtime.RecordGet(x_0, "str").StrVal()) == ("Bar")) && ((gopurs_runtime.RecordGet(x_0, "bool").IntVal) != (0))
	}
end_branch_0:
	return __t0
}

func Call_Main_isDesc(v_0_loop []float64) bool {
	var v_0 []float64 = v_0_loop
	_ = v_0
	var __t_and_1 bool = false
	if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (2) {

		var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
			arr := v_0
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), 0).FloatVal()), gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
			arr := v_0
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), 1).FloatVal()))
		__t_and_1 = (uint32(__t_tag_0.IntVal) == 380165415)
	}
	return __t_and_1
}

func Call_Main_h(o_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var o_0 []gopurs_runtime.Value = o_0_loop
	_ = o_0
	var __t0 []gopurs_runtime.Value
	{
		if (gopurs_runtime.Int(int64(len(o_0))).IntVal) == (3) {
			__t0 = o_0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()
	}
end_branch_0:
	return __t0
}

func Call_Main_f(o_0_loop gopurs_runtime.Value) int64 {
	var o_0 gopurs_runtime.Value = o_0_loop
	_ = o_0
	var __t0 int64
	{
		if (gopurs_runtime.RecordGet(o_0, "foo").StrVal()) == ("Foo") {
			__t0 = gopurs_runtime.RecordGet(o_0, "bar").IntVal
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = 0
	}
end_branch_0:
	return __t0
}
