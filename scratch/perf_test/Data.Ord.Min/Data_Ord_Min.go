package Data_Ord_Min

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Min gopurs_runtime.Value
var once_Min sync.Once
func Get_Min() gopurs_runtime.Value {
	once_Min.Do(func() {
		Min = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Min
}

var showMin gopurs_runtime.Value
var once_showMin sync.Once
func Get_showMin() gopurs_runtime.Value {
	once_showMin.Do(func() {
		showMin = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Min ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showMin
}

var semigroupMin gopurs_runtime.Value
var once_semigroupMin sync.Once
func Get_semigroupMin() gopurs_runtime.Value {
	once_semigroupMin.Do(func() {
		semigroupMin = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_1), v1_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = v_1
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = v_1
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = v1_2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})})
})
	})
	return semigroupMin
}

var newtypeMin gopurs_runtime.Value
var once_newtypeMin sync.Once
func Get_newtypeMin() gopurs_runtime.Value {
	once_newtypeMin.Do(func() {
		newtypeMin = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeMin
}

var monoidMin gopurs_runtime.Value
var once_monoidMin sync.Once
func Get_monoidMin() gopurs_runtime.Value {
	once_monoidMin.Do(func() {
		monoidMin = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{})
semigroupMin1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2), v1_3)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = v_2
} else {
if (gopurs_runtime.Bool(v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t3 = v_2
} else {
if (gopurs_runtime.Bool(v_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = v1_3
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t3
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["top"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMin1_2_1
})})
})
	})
	return monoidMin
}

var eqMin gopurs_runtime.Value
var once_eqMin sync.Once
func Get_eqMin() gopurs_runtime.Value {
	once_eqMin.Do(func() {
		eqMin = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqMin
}

var ordMin gopurs_runtime.Value
var once_ordMin sync.Once
func Get_ordMin() gopurs_runtime.Value {
	once_ordMin.Do(func() {
		ordMin = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_2), v1_3)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})})
})
	})
	return ordMin
}


