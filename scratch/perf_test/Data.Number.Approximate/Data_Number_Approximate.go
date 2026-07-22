package Data_Number_Approximate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
)

var Tolerance gopurs_runtime.Value
var once_Tolerance sync.Once
func Get_Tolerance() gopurs_runtime.Value {
	once_Tolerance.Do(func() {
		Tolerance = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Tolerance
}

var Fraction gopurs_runtime.Value
var once_Fraction sync.Once
func Get_Fraction() gopurs_runtime.Value {
	once_Fraction.Do(func() {
		Fraction = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Fraction
}

var eqRelative gopurs_runtime.Value
var once_eqRelative sync.Once
func Get_eqRelative() gopurs_runtime.Value {
	once_eqRelative.Do(func() {
		eqRelative = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.IntVal == gopurs_runtime.Float(0.0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), v2_2)), v_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
} else {
if (gopurs_runtime.Bool(v2_2.IntVal == gopurs_runtime.Float(0.0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), v1_1)), v_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), v1_1), v2_2))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_numDiv(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), v_0), gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), v1_1), v2_2)))), gopurs_runtime.Float(2.0))).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
}
}
return __t0
})
})
})
	})
	return eqRelative
}

var eqApproximate gopurs_runtime.Value
var once_eqApproximate sync.Once
func Get_eqApproximate() gopurs_runtime.Value {
	once_eqApproximate.Do(func() {
		eqApproximate = gopurs_runtime.Apply(Get_eqRelative(), gopurs_runtime.Float(0.000001))
	})
	return eqApproximate
}

var neqApproximate gopurs_runtime.Value
var once_neqApproximate sync.Once
func Get_neqApproximate() gopurs_runtime.Value {
	once_neqApproximate.Do(func() {
		neqApproximate = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolNot(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqRelative(), gopurs_runtime.Float(0.000001)), x_0), y_1))
})
})
	})
	return neqApproximate
}

var eqAbsolute gopurs_runtime.Value
var once_eqAbsolute sync.Once
func Get_eqAbsolute() gopurs_runtime.Value {
	once_eqAbsolute.Do(func() {
		eqAbsolute = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), x_1), y_2))), v_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
})
})
})
	})
	return eqAbsolute
}


