package Data_Profunctor_Split

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var SplitF gopurs_runtime.Value
var once_SplitF sync.Once
func Get_SplitF() gopurs_runtime.Value {
	once_SplitF.Do(func() {
		SplitF = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitF"), "value0": value0, "value1": value1, "value2": value2})
})
})
})
	})
	return SplitF
}

var unSplit gopurs_runtime.Value
var once_unSplit sync.Once
func Get_unSplit() gopurs_runtime.Value {
	once_unSplit.Do(func() {
		unSplit = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
})), v_1)
})
})
	})
	return unSplit
}

var split gopurs_runtime.Value
var once_split sync.Once
func Get_split() gopurs_runtime.Value {
	once_split.Do(func() {
		split = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fx_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitF"), "value0": f_0, "value1": g_1, "value2": fx_2}))
})
})
})
	})
	return split
}

var profunctorSplit gopurs_runtime.Value
var once_profunctorSplit sync.Once
func Get_profunctorSplit() gopurs_runtime.Value {
	once_profunctorSplit.Do(func() {
		profunctorSplit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unSplit(), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_split(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_2, gopurs_runtime.Apply(f_0, x_4))
})), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(i_3, x_4))
}))
})
}))
})
})})
	})
	return profunctorSplit
}

var lowerSplit gopurs_runtime.Value
var once_lowerSplit sync.Once
func Get_lowerSplit() gopurs_runtime.Value {
	once_lowerSplit.Do(func() {
		lowerSplit = gopurs_runtime.Func(func(dictInvariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unSplit(), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictInvariant_0.PtrVal.(map[string]gopurs_runtime.Value)["imap"], a_2), b_1)
})
}))
})
	})
	return lowerSplit
}

var liftSplit gopurs_runtime.Value
var once_liftSplit sync.Once
func Get_liftSplit() gopurs_runtime.Value {
	once_liftSplit.Do(func() {
		liftSplit = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_split(), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
	})
	return liftSplit
}

var hoistSplit gopurs_runtime.Value
var once_hoistSplit sync.Once
func Get_hoistSplit() gopurs_runtime.Value {
	once_hoistSplit.Do(func() {
		hoistSplit = gopurs_runtime.Func(func(nat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unSplit(), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitF"), "value0": f_1, "value1": g_2, "value2": gopurs_runtime.Apply(nat_0, x_3)}))
})
})
}))
})
	})
	return hoistSplit
}

var functorSplit gopurs_runtime.Value
var once_functorSplit sync.Once
func Get_functorSplit() gopurs_runtime.Value {
	once_functorSplit.Do(func() {
		functorSplit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unSplit(), gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fx_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("SplitF"), "value0": g_1, "value1": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(h_2, x_4))
}), "value2": fx_3}))
})
})
}))
})})
	})
	return functorSplit
}


