package Data_Ord_Down

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Down gopurs_runtime.Value
var once_Down sync.Once
func Get_Down() gopurs_runtime.Value {
	once_Down.Do(func() {
		Down = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Down
}

var showDown gopurs_runtime.Value
var once_showDown sync.Once
func Get_showDown() gopurs_runtime.Value {
	once_showDown.Do(func() {
		showDown = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Down ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showDown
}

var newtypeDown gopurs_runtime.Value
var once_newtypeDown sync.Once
func Get_newtypeDown() gopurs_runtime.Value {
	once_newtypeDown.Do(func() {
		newtypeDown = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeDown
}

var eqDown gopurs_runtime.Value
var once_eqDown sync.Once
func Get_eqDown() gopurs_runtime.Value {
	once_eqDown.Do(func() {
		eqDown = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqDown
}

var ordDown gopurs_runtime.Value
var once_ordDown sync.Once
func Get_ordDown() gopurs_runtime.Value {
	once_ordDown.Do(func() {
		ordDown = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2), v1_3)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
if (gopurs_runtime.Bool(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t2
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})})
})
	})
	return ordDown
}

var boundedDown gopurs_runtime.Value
var once_boundedDown sync.Once
func Get_boundedDown() gopurs_runtime.Value {
	once_boundedDown.Do(func() {
		boundedDown = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
ordDown1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_3), v1_4)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"], "bottom": dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["top"], "Ord0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordDown1_3_2
})})
})
	})
	return boundedDown
}


