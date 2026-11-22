package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Number_Approximate_lessThanOrEq gopurs_runtime.Value
var once_Data_Number_Approximate_lessThanOrEq sync.Once

func Get_Data_Number_Approximate_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Number_Approximate_lessThanOrEq.Do(func() {
		cache_Data_Number_Approximate_lessThanOrEq = Get_Data_Ord_lessThanOrEq__756116985()
	})
	return cache_Data_Number_Approximate_lessThanOrEq
}

var cache_Data_Number_Approximate_sub gopurs_runtime.Value
var once_Data_Number_Approximate_sub sync.Once

func Get_Data_Number_Approximate_sub() gopurs_runtime.Value {
	once_Data_Number_Approximate_sub.Do(func() {
		cache_Data_Number_Approximate_sub = Get_Data_Ring_sub__179436318()
	})
	return cache_Data_Number_Approximate_sub
}

var cache_Data_Number_Approximate_div gopurs_runtime.Value
var once_Data_Number_Approximate_div sync.Once

func Get_Data_Number_Approximate_div() gopurs_runtime.Value {
	once_Data_Number_Approximate_div.Do(func() {
		cache_Data_Number_Approximate_div = Get_Data_EuclideanRing_div__266314078()
	})
	return cache_Data_Number_Approximate_div
}

var cache_Data_Number_Approximate_mul gopurs_runtime.Value
var once_Data_Number_Approximate_mul sync.Once

func Get_Data_Number_Approximate_mul() gopurs_runtime.Value {
	once_Data_Number_Approximate_mul.Do(func() {
		cache_Data_Number_Approximate_mul = Get_Data_Semiring_mul__2603109278()
	})
	return cache_Data_Number_Approximate_mul
}

var cache_Data_Number_Approximate_add gopurs_runtime.Value
var once_Data_Number_Approximate_add sync.Once

func Get_Data_Number_Approximate_add() gopurs_runtime.Value {
	once_Data_Number_Approximate_add.Do(func() {
		cache_Data_Number_Approximate_add = Get_Data_Semiring_add__2603109278()
	})
	return cache_Data_Number_Approximate_add
}

var cache_Data_Number_Approximate_not gopurs_runtime.Value
var once_Data_Number_Approximate_not sync.Once

func Get_Data_Number_Approximate_not() gopurs_runtime.Value {
	once_Data_Number_Approximate_not.Do(func() {
		cache_Data_Number_Approximate_not = Get_Data_HeytingAlgebra_not__1990131970()
	})
	return cache_Data_Number_Approximate_not
}

var cache_Data_Number_Approximate_Tolerance gopurs_runtime.Value
var once_Data_Number_Approximate_Tolerance sync.Once

func Get_Data_Number_Approximate_Tolerance() gopurs_runtime.Value {
	once_Data_Number_Approximate_Tolerance.Do(func() {
		cache_Data_Number_Approximate_Tolerance = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Number_Approximate_Tolerance(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Number_Approximate_Tolerance
}

var cache_Data_Number_Approximate_Fraction gopurs_runtime.Value
var once_Data_Number_Approximate_Fraction sync.Once

func Get_Data_Number_Approximate_Fraction() gopurs_runtime.Value {
	once_Data_Number_Approximate_Fraction.Do(func() {
		cache_Data_Number_Approximate_Fraction = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Number_Approximate_Fraction(x_0_box.FloatVal()))
		})
	})
	return cache_Data_Number_Approximate_Fraction
}

var cache_Data_Number_Approximate_eqRelative gopurs_runtime.Value
var once_Data_Number_Approximate_eqRelative sync.Once

func Get_Data_Number_Approximate_eqRelative() gopurs_runtime.Value {
	once_Data_Number_Approximate_eqRelative.Do(func() {
		cache_Data_Number_Approximate_eqRelative = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Number_Approximate_eqRelative(v_0_box.FloatVal(), v1_1_box.FloatVal(), v2_2_box.FloatVal()))
		})
	})
	return cache_Data_Number_Approximate_eqRelative
}

var cache_Data_Number_Approximate_eqApproximate gopurs_runtime.Value
var once_Data_Number_Approximate_eqApproximate sync.Once

func Get_Data_Number_Approximate_eqApproximate() gopurs_runtime.Value {
	once_Data_Number_Approximate_eqApproximate.Do(func() {
		cache_Data_Number_Approximate_eqApproximate = gopurs_runtime.Apply(Get_Data_Number_Approximate_eqRelative(), gopurs_runtime.Float(0.000001))
	})
	return cache_Data_Number_Approximate_eqApproximate
}

var cache_Data_Number_Approximate_neqApproximate gopurs_runtime.Value
var once_Data_Number_Approximate_neqApproximate sync.Once

func Get_Data_Number_Approximate_neqApproximate() gopurs_runtime.Value {
	once_Data_Number_Approximate_neqApproximate.Do(func() {
		cache_Data_Number_Approximate_neqApproximate = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Number_Approximate_neqApproximate(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Data_Number_Approximate_neqApproximate
}

var cache_Data_Number_Approximate_eqAbsolute gopurs_runtime.Value
var once_Data_Number_Approximate_eqAbsolute sync.Once

func Get_Data_Number_Approximate_eqAbsolute() gopurs_runtime.Value {
	once_Data_Number_Approximate_eqAbsolute.Do(func() {
		cache_Data_Number_Approximate_eqAbsolute = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Number_Approximate_eqAbsolute(v_0_box.FloatVal(), x_1_box.FloatVal(), y_2_box.FloatVal()))
		})
	})
	return cache_Data_Number_Approximate_eqAbsolute
}

func Call_Data_Number_Approximate_Tolerance(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Number_Approximate_Fraction(x_0_loop float64) float64 {
	var x_0 float64 = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Number_Approximate_eqRelative(v_0_loop float64, v1_1_loop float64, v2_2_loop float64) bool {
	var v_0 float64 = v_0_loop
	_ = v_0
	var v1_1 float64 = v1_1_loop
	_ = v1_1
	var v2_2 float64 = v2_2_loop
	_ = v2_2
	var __t0 bool
	{
		if (v1_1) == (0.0) {
			__t0 = (Data_Number_Abs(v2_2)) <= (v_0)
			goto end_branch_0
		} else {

		}
	}
	{
		if (v2_2) == (0.0) {
			__t0 = (Data_Number_Abs(v1_1)) <= (v_0)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = (Data_Number_Abs((v1_1) - (v2_2))) <= (((v_0) * (Data_Number_Abs((v1_1) + (v2_2)))) / (2.0))
	}
end_branch_0:
	return __t0
}

func Call_Data_Number_Approximate_neqApproximate(x_0_loop float64, y_1_loop float64) bool {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	return (Call_Data_Number_Approximate_eqRelative(0.000001, x_0, y_1)) != (true)
}

func Call_Data_Number_Approximate_eqAbsolute(v_0_loop float64, x_1_loop float64, y_2_loop float64) bool {
	var v_0 float64 = v_0_loop
	_ = v_0
	var x_1 float64 = x_1_loop
	_ = x_1
	var y_2 float64 = y_2_loop
	_ = y_2
	return (Data_Number_Abs((x_1) - (y_2))) <= (v_0)
}
