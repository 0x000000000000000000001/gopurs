package Data_Unfoldable1

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
})
	})
	return fromJust
}

var unfoldr1 gopurs_runtime.Value
var once_unfoldr1 sync.Once
func Get_unfoldr1() gopurs_runtime.Value {
	once_unfoldr1.Do(func() {
		unfoldr1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"]
})
	})
	return unfoldr1
}

var unfoldable1Maybe gopurs_runtime.Value
var once_unfoldable1Maybe sync.Once
func Get_unfoldable1Maybe() gopurs_runtime.Value {
	once_unfoldable1Maybe.Do(func() {
		unfoldable1Maybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr1": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(f_0, b_1).PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
})
})})
	})
	return unfoldable1Maybe
}

var unfoldable1Array gopurs_runtime.Value
var once_unfoldable1Array sync.Once
func Get_unfoldable1Array() gopurs_runtime.Value {
	once_unfoldable1Array.Do(func() {
		unfoldable1Array = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr1": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing()), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
}))), pkg_Data_Tuple.Get_fst()), pkg_Data_Tuple.Get_snd())})
	})
	return unfoldable1Array
}

var replicate1 gopurs_runtime.Value
var once_replicate1 sync.Once
func Get_replicate1() gopurs_runtime.Value {
	once_replicate1.Do(func() {
		replicate1 = gopurs_runtime.Func(func(dictUnfoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], i_3), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_2, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_2, "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_3), gopurs_runtime.Int(1))})})
}
return __t0
})), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), n_1), gopurs_runtime.Int(1)))
})
})
})
	})
	return replicate1
}

var replicate1A gopurs_runtime.Value
var once_replicate1A sync.Once
func Get_replicate1A() gopurs_runtime.Value {
	once_replicate1A.Do(func() {
		replicate1A = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictUnfoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTraversable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_3_0 := gopurs_runtime.Apply(dictTraversable1_2.PtrVal.(map[string]gopurs_runtime.Value)["sequence1"], dictApply_0)
return gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_replicate1(), dictUnfoldable1_1), n_4), m_5))
})
})
})
})
})
	})
	return replicate1A
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(dictUnfoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_replicate1(), dictUnfoldable1_0), gopurs_runtime.Int(1))
})
	})
	return singleton
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func(func(dictUnfoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(start_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(end_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], end_2), start_1).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Int(0)), gopurs_runtime.Int(1))
}
__local_var_3_0 := __t1
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
i_prime_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), i_4), __local_var_3_0)
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), i_4), end_2)).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": i_prime_5_2})
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": i_4, "value1": __t3})
})), start_1)
})
})
})
	})
	return range_
}

var iterateN gopurs_runtime.Value
var once_iterateN sync.Once
func Get_iterateN() gopurs_runtime.Value {
	once_iterateN.Do(func() {
		iterateN = gopurs_runtime.Func(func(dictUnfoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_2, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Int(1))})})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __t0})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": s_3, "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), n_1), gopurs_runtime.Int(1))}))
})
})
})
})
	})
	return iterateN
}

func Get_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return Unfoldr1ArrayImpl
}
