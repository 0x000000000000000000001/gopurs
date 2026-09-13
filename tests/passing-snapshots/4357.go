package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_monoidAdditive gopurs_runtime.Value
var once_Main_monoidAdditive sync.Once

func Get_Main_monoidAdditive() gopurs_runtime.Value {
	once_Main_monoidAdditive.Do(func() {
		cache_Main_monoidAdditive = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_2095395797_1201789390(Rebox_Main_1201789390_2095395797(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Additive_monoidAdditive(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))})))))}
	})
	return cache_Main_monoidAdditive
}

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, value0.IntVal}))}
		})
	})
	return cache_Main_Foo
}

var cache_Main_Bar gopurs_runtime.Value
var once_Main_Bar sync.Once

func Get_Main_Bar() gopurs_runtime.Value {
	once_Main_Bar.Do(func() {
		cache_Main_Bar = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2512729583, UnsafePtr: unsafe.Pointer((&Constructor_Main_Bar{1, value0.IntVal}))}
		})
	})
	return cache_Main_Bar
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test(Rebox_Main_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))
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

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_g(v_0_box))
		})
	})
	return cache_Main_g
}

type Constructor_Main_Foo struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_Bar struct {
	Rc uint32
	V0 int64
}

func Call_Main_test(v_0_loop *Constructor_Data_Maybe_Just[int64]) int64 {
	var v_0 *Constructor_Data_Maybe_Just[int64] = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if v_0 != nil {
			__t0 = (v_0).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), Rebox_Main_2095395797_1201789390(Rebox_Main_1201789390_2095395797(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Additive_monoidAdditive(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Main_348932501_2826095630(Rebox_Main_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}))))), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Array([]gopurs_runtime.Value{})).IntVal
	}
end_branch_0:
	return __t0
}

func Call_Main_g(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 int64
	{
		if v_0.Type == 9 && v_0.IntVal == 2512729583 {
			__t0 = (*Constructor_Main_Bar)(v_0.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 2763139640 {
			__t0 = (*Constructor_Main_Foo)(v_0.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = int64(42)
	}
end_branch_0:
	return __t0
}

func Rebox_Main_1201789390_2095395797(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_2095395797_1201789390(in *Constructor_Data_Monoid_Monoid[int64]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_2826095630_348932501(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.IntVal
	out.V3 = in.V3.IntVal
	return out
}

func Rebox_Main_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Main_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Int(in.V2)
	out.V3 = gopurs_runtime.Int(in.V3)
	return out
}
