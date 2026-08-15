package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_A gopurs_runtime.Value
var once_Main_A sync.Once

func Get_Main_A() gopurs_runtime.Value {
	once_Main_A.Do(func() {
		cache_Main_A = gopurs_runtime.Value{Type: 9, IntVal: int64(4219254943), UnsafePtr: nil}
	})
	return cache_Main_A
}

var cache_Main_parseTest gopurs_runtime.Value
var once_Main_parseTest sync.Once

func Get_Main_parseTest() gopurs_runtime.Value {
	once_Main_parseTest.Do(func() {
		cache_Main_parseTest = gopurs_runtime.Func3(func(dictPartial_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_parseTest(dictPartial_0_box, v_1_box, v1_2_box)
		})
	})
	return cache_Main_parseTest
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_guardsTest gopurs_runtime.Value
var once_Main_guardsTest sync.Once

func Get_Main_guardsTest() gopurs_runtime.Value {
	once_Main_guardsTest.Do(func() {
		cache_Main_guardsTest = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_guardsTest(func() []float64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]float64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.FloatVal()
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Float(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_guardsTest
}

var cache_Main_gcd gopurs_runtime.Value
var once_Main_gcd sync.Once

func Get_Main_gcd() gopurs_runtime.Value {
	once_Main_gcd.Do(func() {
		cache_Main_gcd = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_gcd(v_0_box.FloatVal(), v1_1_box.FloatVal()))
		})
	})
	return cache_Main_gcd
}

type Constructor_Main_A struct {
	Rc uint32
}

func Call_Main_parseTest(dictPartial_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var v1_2 gopurs_runtime.Value = v1_2_loop
	_ = v1_2
	var __t0 float64
	{
		if (v1_2.FloatVal()) == (0.0) {
			__t0 = 0.0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().FloatVal()
	}
end_branch_0:
	return gopurs_runtime.Float(__t0)
}

func Call_Main_guardsTest(v_0_loop []float64) []float64 {
	var v_0 []float64 = v_0_loop
	_ = v_0
	var __t2 []float64
	{
		var __t_and_1 bool = false
		if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (1) {

			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
				arr := v_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Float(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), 0).FloatVal()), gopurs_runtime.Float(0.0))
			__t_and_1 = (uint32(__t_tag_0.IntVal) == 380165415)
		}
		if __t_and_1 {
			__t2 = func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}()
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = v_0
	}
end_branch_2:
	return __t2
}

func Call_Main_gcd(v_0_loop float64, v1_1_loop float64) float64 {
gcd:
	for {
		if false {
			continue gcd
		}
		var v_0 float64 = v_0_loop
		_ = v_0
		var v1_1 float64 = v1_1_loop
		_ = v1_1
		var __t1 float64
		{
			if (v_0) == (0.0) {
				__t1 = v1_1
				goto end_branch_1
			} else {

			}
		}
		{
			if (v1_1) == (0.0) {
				__t1 = v_0
				goto end_branch_1
			} else {

			}
		}
		{
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.Float(v1_1))
			if uint32(__t_tag_0.IntVal) == 380165415 {
				v_0_loop = 0.0
				v1_1_loop = v1_1
				continue gcd
				__t1 = gopurs_runtime.Value{}.FloatVal()
				goto end_branch_1
			} else {

			}
		}
		{
			v_0_loop = 0.0
			v1_1_loop = v_0
			continue gcd
			__t1 = gopurs_runtime.Value{}.FloatVal()
		}
	end_branch_1:
		return __t1
	}
}
