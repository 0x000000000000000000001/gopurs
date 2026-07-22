package Data_Unfoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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

var unfoldr gopurs_runtime.Value
var once_unfoldr sync.Once
func Get_unfoldr() gopurs_runtime.Value {
	once_unfoldr.Do(func() {
		unfoldr = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"]
})
	})
	return unfoldr
}

var unfoldableMaybe gopurs_runtime.Value
var once_unfoldableMaybe sync.Once
func Get_unfoldableMaybe() gopurs_runtime.Value {
	once_unfoldableMaybe.Do(func() {
		unfoldableMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, b_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
}), "Unfoldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Maybe()
})})
	})
	return unfoldableMaybe
}

var unfoldableArray gopurs_runtime.Value
var once_unfoldableArray sync.Once
func Get_unfoldableArray() gopurs_runtime.Value {
	once_unfoldableArray.Do(func() {
		unfoldableArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"unfoldr": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing()), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
}))), pkg_Data_Tuple.Get_fst()), pkg_Data_Tuple.Get_snd()), "Unfoldable10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
})})
	})
	return unfoldableArray
}

var replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		replicate = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], i_3), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_2, "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), i_3), gopurs_runtime.Int(1))})})
}
return __t0
})), n_1)
})
})
})
	})
	return replicate
}

var replicateA gopurs_runtime.Value
var once_replicateA sync.Once
func Get_replicateA() gopurs_runtime.Value {
	once_replicateA.Do(func() {
		replicateA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictUnfoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTraversable_2 gopurs_runtime.Value) gopurs_runtime.Value {
sequence_3_0 := gopurs_runtime.Apply(dictTraversable_2.PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_0)
return gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence_3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_replicate(), dictUnfoldable_1), n_4), m_5))
})
})
})
})
})
	})
	return replicateA
}

var none gopurs_runtime.Value
var once_none sync.Once
func Get_none() gopurs_runtime.Value {
	once_none.Do(func() {
		none = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
})), pkg_Data_Unit.Get_unit())
})
	})
	return none
}

var fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		fromMaybe = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(b_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": b_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})})})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}))
})
	})
	return fromMaybe
}

func Get_unfoldrArrayImpl() gopurs_runtime.Value {
	return UnfoldrArrayImpl
}
