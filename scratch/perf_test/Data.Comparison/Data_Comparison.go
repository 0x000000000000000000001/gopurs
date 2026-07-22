package Data_Comparison

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Comparison gopurs_runtime.Value
var once_Comparison sync.Once
func Get_Comparison() gopurs_runtime.Value {
	once_Comparison.Do(func() {
		Comparison = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Comparison
}

var semigroupComparison gopurs_runtime.Value
var once_semigroupComparison sync.Once
func Get_semigroupComparison() gopurs_runtime.Value {
	once_semigroupComparison.Do(func() {
		semigroupComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(v_0, x_2)
__local_var_4_1 := gopurs_runtime.Apply(v1_1, x_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(__local_var_3_0, x_5)
__local_var_7_3 := gopurs_runtime.Apply(__local_var_4_1, x_5)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(__local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(__local_var_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t4 = __local_var_7_3
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
})
})
})})
	})
	return semigroupComparison
}

var newtypeComparison gopurs_runtime.Value
var once_newtypeComparison sync.Once
func Get_newtypeComparison() gopurs_runtime.Value {
	once_newtypeComparison.Do(func() {
		newtypeComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeComparison
}

var monoidComparison gopurs_runtime.Value
var once_monoidComparison sync.Once
func Get_monoidComparison() gopurs_runtime.Value {
	once_monoidComparison.Do(func() {
		monoidComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupComparison()
})})
	})
	return monoidComparison
}

var defaultComparison gopurs_runtime.Value
var once_defaultComparison sync.Once
func Get_defaultComparison() gopurs_runtime.Value {
	once_defaultComparison.Do(func() {
		defaultComparison = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
})
	})
	return defaultComparison
}

var contravariantComparison gopurs_runtime.Value
var once_contravariantComparison sync.Once
func Get_contravariantComparison() gopurs_runtime.Value {
	once_contravariantComparison.Do(func() {
		contravariantComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cmap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2)), gopurs_runtime.Apply(f_0, y_3))
})
})
})
})})
	})
	return contravariantComparison
}


