package Data_Decide

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Divide "gopurs/output/Data.Divide"
)

var choosePredicate gopurs_runtime.Value
var once_choosePredicate sync.Once
func Get_choosePredicate() gopurs_runtime.Value {
	once_choosePredicate.Do(func() {
		choosePredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"choose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(f_0, x_3)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(v_1, __local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(__local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(v1_2, __local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_dividePredicate()
})})
	})
	return choosePredicate
}

var chooseOp gopurs_runtime.Value
var once_chooseOp sync.Once
func Get_chooseOp() gopurs_runtime.Value {
	once_chooseOp.Do(func() {
		chooseOp = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
divideOp_1_0 := gopurs_runtime.Apply(pkg_Data_Divide.Get_divideOp(), dictSemigroup_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"choose": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(f_2, x_5)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(v_3, __local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(__local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(v1_4, __local_var_6_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
})
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return divideOp_1_0
})})
})
	})
	return chooseOp
}

var chooseEquivalence gopurs_runtime.Value
var once_chooseEquivalence sync.Once
func Get_chooseEquivalence() gopurs_runtime.Value {
	once_chooseEquivalence.Do(func() {
		chooseEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"choose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
v3_6_4 := gopurs_runtime.Apply(f_0, b_4)
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v3_6_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v3_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Bool(false)
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t5
} else {
if (gopurs_runtime.Bool(v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
v3_6_2 := gopurs_runtime.Apply(f_0, b_4)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(v3_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_2, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v3_6_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t3
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
})
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideEquivalence()
})})
	})
	return chooseEquivalence
}

var chooseComparison gopurs_runtime.Value
var once_chooseComparison sync.Once
func Get_chooseComparison() gopurs_runtime.Value {
	once_chooseComparison.Do(func() {
		chooseComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"choose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
v2_5_0 := gopurs_runtime.Apply(f_0, a_3)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
v3_6_4 := gopurs_runtime.Apply(f_0, b_4)
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v3_6_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v3_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t5
} else {
if (gopurs_runtime.Bool(v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
v3_6_2 := gopurs_runtime.Apply(f_0, b_4)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(v3_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_2, v2_5_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v3_6_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t3
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
})
})
}), "Divide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideComparison()
})})
	})
	return chooseComparison
}

var choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		choose = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["choose"]
})
	})
	return choose
}

var chosen gopurs_runtime.Value
var once_chosen sync.Once
func Get_chosen() gopurs_runtime.Value {
	once_chosen.Do(func() {
		chosen = gopurs_runtime.Func(func(dictDecide_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictDecide_0.PtrVal.(map[string]gopurs_runtime.Value)["choose"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return chosen
}


