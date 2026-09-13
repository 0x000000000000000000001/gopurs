package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_either gopurs_runtime.Value
var once_Main_either sync.Once

func Get_Main_either() gopurs_runtime.Value {
	once_Main_either.Do(func() {
		cache_Main_either = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_either(v_0_box, v1_1_box, v2_2_box)
		})
	})
	return cache_Main_either
}

var cache_Main_either__1840012924 gopurs_runtime.Value
var once_Main_either__1840012924 sync.Once

func Get_Main_either__1840012924() gopurs_runtime.Value {
	once_Main_either__1840012924.Do(func() {
		cache_Main_either__1840012924 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_either__1840012924(v_unused_0_box, v1_unused_1_box, v2_2_box))
		})
	})
	return cache_Main_either__1840012924
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Call_Main_either__1840012924(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: 1485529257, UnsafePtr: unsafe.Pointer(Rebox_Main_234776603_17902363((&Constructor_Either_Left[string, string]{1, "Done"})))})))
	})
	return cache_Main_main
}

func Call_Main_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var __t0 gopurs_runtime.Value
	{
		if v2_2.Type == 9 && v2_2.IntVal == 1485529257 {
			__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2.Type == 9 && v2_2.IntVal == 3726768370 {
			__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_either__1840012924(v_unused_0_loop gopurs_runtime.Value, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) string {
either__1840012924:
	for {
		if false {
			continue either__1840012924
		}
		var v_unused_0 gopurs_runtime.Value = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
		_ = v1_unused_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if v2_2.Type == 9 && v2_2.IntVal == 1485529257 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Str((*Constructor_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0.StrVal())).StrVal())
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_2.Type == 9 && v2_2.IntVal == 3726768370 {
				__t0 = gopurs_runtime.Str(gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), (*Constructor_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0).StrVal())
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0.StrVal()
	}
}

func Rebox_Main_234776603_17902363(in *Constructor_Either_Left[string, string]) *Constructor_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
