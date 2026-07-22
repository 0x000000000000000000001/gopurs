package Data_Profunctor_Choice

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
)

var right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		right = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["right"]
})
	})
	return right
}

var left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		left = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["left"]
})
	})
	return left
}

var splitChoice gopurs_runtime.Value
var once_splitChoice sync.Once
func Get_splitChoice() gopurs_runtime.Value {
	once_splitChoice.Do(func() {
		splitChoice = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictChoice_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], gopurs_runtime.Apply(dictChoice_1.PtrVal.(map[string]gopurs_runtime.Value)["right"], r_3)), gopurs_runtime.Apply(dictChoice_1.PtrVal.(map[string]gopurs_runtime.Value)["left"], l_2))
})
})
})
})
	})
	return splitChoice
}

var fanin gopurs_runtime.Value
var once_fanin sync.Once
func Get_fanin() gopurs_runtime.Value {
	once_fanin.Do(func() {
		fanin = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictChoice_1 gopurs_runtime.Value) gopurs_runtime.Value {
rmap_2_0 := gopurs_runtime.Apply(pkg_Data_Profunctor.Get_rmap(), gopurs_runtime.Apply(dictChoice_1.PtrVal.(map[string]gopurs_runtime.Value)["Profunctor0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(rmap_2_0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupoid_0.PtrVal.(map[string]gopurs_runtime.Value)["compose"], gopurs_runtime.Apply(dictChoice_1.PtrVal.(map[string]gopurs_runtime.Value)["right"], r_4)), gopurs_runtime.Apply(dictChoice_1.PtrVal.(map[string]gopurs_runtime.Value)["left"], l_3)))
})
})
})
})
	})
	return fanin
}

var choiceFn gopurs_runtime.Value
var once_choiceFn sync.Once
func Get_choiceFn() gopurs_runtime.Value {
	once_choiceFn.Do(func() {
		choiceFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"left": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(v_0, v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "right": pkg_Data_Either.Get_functorEither().PtrVal.(map[string]gopurs_runtime.Value)["map"], "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
})})
	})
	return choiceFn
}


