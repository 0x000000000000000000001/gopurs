package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Zero gopurs_runtime.Value
var once_Main_Zero sync.Once

func Get_Main_Zero() gopurs_runtime.Value {
	once_Main_Zero.Do(func() {
		cache_Main_Zero = gopurs_runtime.Value{Type: 9, IntVal: 1285566470, UnsafePtr: unsafe.Pointer((*Constructor_Main_Even)(nil))}
	})
	return cache_Main_Zero
}

var cache_Main_Even gopurs_runtime.Value
var once_Main_Even sync.Once

func Get_Main_Even() gopurs_runtime.Value {
	once_Main_Even.Do(func() {
		cache_Main_Even = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1285566470, UnsafePtr: unsafe.Pointer((&Constructor_Main_Even{1, value0}))}
		})
	})
	return cache_Main_Even
}

var cache_Main_Odd gopurs_runtime.Value
var once_Main_Odd sync.Once

func Get_Main_Odd() gopurs_runtime.Value {
	once_Main_Odd.Do(func() {
		cache_Main_Odd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Odd
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
		cache_Main_g = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_g(x_0_box.FloatVal()))
		})
	})
	return cache_Main_g
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_f(v_0_box.FloatVal()))
		})
	})
	return cache_Main_f
}

var cache_Main_oddToNumber gopurs_runtime.Value
var once_Main_oddToNumber sync.Once

func Get_Main_oddToNumber() gopurs_runtime.Value {
	once_Main_oddToNumber.Do(func() {
		cache_Main_oddToNumber = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_oddToNumber(v_0_box))
		})
	})
	return cache_Main_oddToNumber
}

var cache_Main_evenToNumber gopurs_runtime.Value
var once_Main_evenToNumber sync.Once

func Get_Main_evenToNumber() gopurs_runtime.Value {
	once_Main_evenToNumber.Do(func() {
		cache_Main_evenToNumber = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_evenToNumber(gopurs_runtime.CoerceToStruct[Constructor_Main_Even](v_0_box)))
		})
	})
	return cache_Main_evenToNumber
}

type Constructor_Main_Zero struct {
	Rc uint32
}

type Constructor_Main_Even struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Odd struct {
	Rc uint32
	V0 *Constructor_Main_Even
}

func Call_Main_g(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return Call_Main_f((x_0) / (0.0))
}

func Call_Main_f(v_0_loop float64) float64 {
	var v_0 float64 = v_0_loop
	_ = v_0
	var __t0 float64
	{
		if (v_0) == (0.0) {
			__t0 = 0.0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = (Call_Main_f((v_0) / (0.0))) + (0.0)
	}
end_branch_0:
	return __t0
}

func Call_Main_oddToNumber(v_0_loop gopurs_runtime.Value) float64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (Call_Main_evenToNumber(gopurs_runtime.CoerceToStruct[Constructor_Main_Even](v_0))) + (0.0)
}

func Call_Main_evenToNumber(v_0_loop *Constructor_Main_Even) float64 {
	var v_0 *Constructor_Main_Even = v_0_loop
	_ = v_0
	var __t0 float64
	{
		if v_0 == nil {
			__t0 = 0.0
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = (Call_Main_evenToNumber(gopurs_runtime.CoerceToStruct[Constructor_Main_Even]((v_0).V0))) + (0.0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().FloatVal()
	}
end_branch_0:
	return __t0
}
