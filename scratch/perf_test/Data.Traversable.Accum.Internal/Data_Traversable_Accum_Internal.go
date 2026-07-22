package Data_Traversable_Accum_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var StateR gopurs_runtime.Value
var once_StateR sync.Once
func Get_StateR() gopurs_runtime.Value {
	once_StateR.Do(func() {
		StateR = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return StateR
}

var StateL gopurs_runtime.Value
var once_StateL sync.Once
func Get_StateL() gopurs_runtime.Value {
	once_StateL.Do(func() {
		StateL = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return StateL
}

var stateR gopurs_runtime.Value
var once_stateR sync.Once
func Get_stateR() gopurs_runtime.Value {
	once_stateR.Do(func() {
		stateR = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return stateR
}

var stateL gopurs_runtime.Value
var once_stateL sync.Once
func Get_stateL() gopurs_runtime.Value {
	once_stateL.Do(func() {
		stateL = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return stateL
}

var functorStateR gopurs_runtime.Value
var once_functorStateR sync.Once
func Get_functorStateR() gopurs_runtime.Value {
	once_functorStateR.Do(func() {
		functorStateR = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["accum"], "value": gopurs_runtime.Apply(f_0, v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value"])})
})
})
})})
	})
	return functorStateR
}

var functorStateL gopurs_runtime.Value
var once_functorStateL sync.Once
func Get_functorStateL() gopurs_runtime.Value {
	once_functorStateL.Do(func() {
		functorStateL = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["accum"], "value": gopurs_runtime.Apply(f_0, v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value"])})
})
})
})})
	})
	return functorStateL
}

var applyStateR gopurs_runtime.Value
var once_applyStateR sync.Once
func Get_applyStateR() gopurs_runtime.Value {
	once_applyStateR.Do(func() {
		applyStateR = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
v1_4_1 := gopurs_runtime.Apply(f_0, v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["accum"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["accum"], "value": gopurs_runtime.Apply(v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value"], v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value"])})
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateR()
})})
	})
	return applyStateR
}

var applyStateL gopurs_runtime.Value
var once_applyStateL sync.Once
func Get_applyStateL() gopurs_runtime.Value {
	once_applyStateL.Do(func() {
		applyStateL = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
v1_4_1 := gopurs_runtime.Apply(x_1, v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["accum"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["accum"], "value": gopurs_runtime.Apply(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value"], v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value"])})
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateL()
})})
	})
	return applyStateL
}

var applicativeStateR gopurs_runtime.Value
var once_applicativeStateR sync.Once
func Get_applicativeStateR() gopurs_runtime.Value {
	once_applicativeStateR.Do(func() {
		applicativeStateR = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": s_1, "value": a_0})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateR()
})})
	})
	return applicativeStateR
}

var applicativeStateL gopurs_runtime.Value
var once_applicativeStateL sync.Once
func Get_applicativeStateL() gopurs_runtime.Value {
	once_applicativeStateL.Do(func() {
		applicativeStateL = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"accum": s_1, "value": a_0})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateL()
})})
	})
	return applicativeStateL
}


