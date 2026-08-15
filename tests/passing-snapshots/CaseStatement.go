package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_N gopurs_runtime.Value
var once_Main_N sync.Once

func Get_Main_N() gopurs_runtime.Value {
	once_Main_N.Do(func() {
		cache_Main_N = gopurs_runtime.Value{Type: 9, IntVal: 2406916180, UnsafePtr: unsafe.Pointer((*Constructor_Main_J)(nil))}
	})
	return cache_Main_N
}

var cache_Main_J gopurs_runtime.Value
var once_Main_J sync.Once

func Get_Main_J() gopurs_runtime.Value {
	once_Main_J.Do(func() {
		cache_Main_J = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2406916180, UnsafePtr: unsafe.Pointer(&Constructor_Main_J{1, value0})}
		})
	})
	return cache_Main_J
}

var cache_Main_A gopurs_runtime.Value
var once_Main_A sync.Once

func Get_Main_A() gopurs_runtime.Value {
	once_Main_A.Do(func() {
		cache_Main_A = gopurs_runtime.Value{Type: 9, IntVal: int64(4219254943), UnsafePtr: nil}
	})
	return cache_Main_A
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Value{Type: 9, IntVal: int64(4250879068), UnsafePtr: nil}
	})
	return cache_Main_B
}

var cache_Main_C gopurs_runtime.Value
var once_Main_C sync.Once

func Get_Main_C() gopurs_runtime.Value {
	once_Main_C.Do(func() {
		cache_Main_C = gopurs_runtime.Value{Type: 9, IntVal: int64(2167983901), UnsafePtr: nil}
	})
	return cache_Main_C
}

var cache_Main_h gopurs_runtime.Value
var once_Main_h sync.Once

func Get_Main_h() gopurs_runtime.Value {
	once_Main_h.Do(func() {
		cache_Main_h = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2406916180, UnsafePtr: unsafe.Pointer(Call_Main_h(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Main_J](v1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Main_J](v2_2_box)))}
		})
	})
	return cache_Main_h
}

var cache_Main_g gopurs_runtime.Value
var once_Main_g sync.Once

func Get_Main_g() gopurs_runtime.Value {
	once_Main_g.Do(func() {
		cache_Main_g = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_g(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal), uint32(v2_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f(v_0_box.StrVal(), v1_1_box.StrVal(), uint32(v2_2_box.IntVal)))
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_N struct {
	Rc uint32
}

type Constructor_Main_J struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_A struct {
	Rc uint32
}

type Constructor_Main_B struct {
	Rc uint32
}

type Constructor_Main_C struct {
	Rc uint32
}

func Call_Main_h(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Main_J, v2_2_loop *Constructor_Main_J) *Constructor_Main_J {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Main_J = v1_1_loop
	_ = v1_1
	var v2_2 *Constructor_Main_J = v2_2_loop
	_ = v2_2
	var __t0 *Constructor_Main_J
	{
		if v1_1 == nil {
			__t0 = v2_2
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2 == nil {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		if (v1_1 != nil) && (v2_2 != nil) {
			__t0 = &Constructor_Main_J{1, gopurs_runtime.Apply2(v_0, (v1_1).V0, (v2_2).V0)}
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_J](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_0:
	return __t0
}

func Call_Main_g(v_0_loop uint32, v1_1_loop uint32, v2_2_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 uint32 = v1_1_loop
	_ = v1_1
	var v2_2 uint32 = v2_2_loop
	_ = v2_2
	var __t0 uint32
	{
		if v2_2 == 4219254943 {
			__t0 = v_0
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2 == 4250879068 {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2 == 2167983901 {
			__t0 = 2167983901
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
	}
end_branch_0:
	return __t0
}

func Call_Main_f(v_0_loop string, v1_1_loop string, v2_2_loop uint32) string {
	var v_0 string = v_0_loop
	_ = v_0
	var v1_1 string = v1_1_loop
	_ = v1_1
	var v2_2 uint32 = v2_2_loop
	_ = v2_2
	var __t0 string
	{
		if v2_2 == 4219254943 {
			__t0 = v_0
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2 == 4250879068 {
			__t0 = v1_1
			goto end_branch_0
		} else {

		}
	}
	{
		if v2_2 == 2167983901 {
			__t0 = "Done"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
	}
end_branch_0:
	return __t0
}
