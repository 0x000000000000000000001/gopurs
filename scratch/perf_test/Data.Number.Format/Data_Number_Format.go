package Data_Number_Format

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var clamp gopurs_runtime.Value
var once_clamp sync.Once
func Get_clamp() gopurs_runtime.Value {
	once_clamp.Do(func() {
		clamp = gopurs_runtime.Func(func(low_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(hi_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], low_0), x_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = x_2
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = low_0
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = low_0
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
__local_var_4_1 := __t2
v_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], hi_1), __local_var_4_1)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = hi_1
} else {
if (gopurs_runtime.Bool(v_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t4 = hi_1
} else {
if (gopurs_runtime.Bool(v_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = __local_var_4_1
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
})
})
	})
	return clamp
}

var Precision gopurs_runtime.Value
var once_Precision sync.Once
func Get_Precision() gopurs_runtime.Value {
	once_Precision.Do(func() {
		Precision = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Precision"), "value0": value0})
})
	})
	return Precision
}

var Fixed gopurs_runtime.Value
var once_Fixed sync.Once
func Get_Fixed() gopurs_runtime.Value {
	once_Fixed.Do(func() {
		Fixed = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Fixed"), "value0": value0})
})
	})
	return Fixed
}

var Exponential gopurs_runtime.Value
var once_Exponential sync.Once
func Get_Exponential() gopurs_runtime.Value {
	once_Exponential.Do(func() {
		Exponential = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Exponential"), "value0": value0})
})
	})
	return Exponential
}

var toStringWith gopurs_runtime.Value
var once_toStringWith sync.Once
func Get_toStringWith() gopurs_runtime.Value {
	once_toStringWith.Do(func() {
		toStringWith = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Precision")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_toPrecisionNative(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Fixed")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_toFixedNative(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Exponential")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_toExponentialNative(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
})
	})
	return toStringWith
}

var precision gopurs_runtime.Value
var once_precision sync.Once
func Get_precision() gopurs_runtime.Value {
	once_precision.Do(func() {
		precision = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Precision"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_clamp(), gopurs_runtime.Int(1)), gopurs_runtime.Int(21)), x_0)})
})
	})
	return precision
}

var fixed gopurs_runtime.Value
var once_fixed sync.Once
func Get_fixed() gopurs_runtime.Value {
	once_fixed.Do(func() {
		fixed = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Fixed"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_clamp(), gopurs_runtime.Int(0)), gopurs_runtime.Int(20)), x_0)})
})
	})
	return fixed
}

var exponential gopurs_runtime.Value
var once_exponential sync.Once
func Get_exponential() gopurs_runtime.Value {
	once_exponential.Do(func() {
		exponential = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Exponential"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_clamp(), gopurs_runtime.Int(0)), gopurs_runtime.Int(20)), x_0)})
})
	})
	return exponential
}

func Get_toExponentialNative() gopurs_runtime.Value {
	return ToExponentialNative
}

func Get_toFixedNative() gopurs_runtime.Value {
	return ToFixedNative
}

func Get_toPrecisionNative() gopurs_runtime.Value {
	return ToPrecisionNative
}

func Get_toString() gopurs_runtime.Value {
	return ToString
}
