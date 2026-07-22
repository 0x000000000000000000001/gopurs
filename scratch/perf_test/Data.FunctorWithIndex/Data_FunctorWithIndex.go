package Data_FunctorWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		mapWithIndex = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"]
})
	})
	return mapWithIndex
}

var mapDefault gopurs_runtime.Value
var once_mapDefault sync.Once
func Get_mapDefault() gopurs_runtime.Value {
	once_mapDefault.Do(func() {
		mapDefault = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
})
})
	})
	return mapDefault
}

var functorWithIndexTuple gopurs_runtime.Value
var once_functorWithIndexTuple sync.Once
func Get_functorWithIndexTuple() gopurs_runtime.Value {
	once_functorWithIndexTuple.Do(func() {
		functorWithIndexTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(__local_var_1_0, m_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
})})
	})
	return functorWithIndexTuple
}

var functorWithIndexProduct gopurs_runtime.Value
var once_functorWithIndexProduct sync.Once
func Get_functorWithIndexProduct() gopurs_runtime.Value {
	once_functorWithIndexProduct.Do(func() {
		functorWithIndexProduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictFunctorWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorProduct1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_7}))
})), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctorWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_7}))
})), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_4_2
})})
})
})
	})
	return functorWithIndexProduct
}

var functorWithIndexMultiplicative gopurs_runtime.Value
var once_functorWithIndexMultiplicative sync.Once
func Get_functorWithIndexMultiplicative() gopurs_runtime.Value {
	once_functorWithIndexMultiplicative.Do(func() {
		functorWithIndexMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
})})
	})
	return functorWithIndexMultiplicative
}

var functorWithIndexMaybe gopurs_runtime.Value
var once_functorWithIndexMaybe sync.Once
func Get_functorWithIndexMaybe() gopurs_runtime.Value {
	once_functorWithIndexMaybe.Do(func() {
		functorWithIndexMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_1_0, v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
})})
	})
	return functorWithIndexMaybe
}

var functorWithIndexLast gopurs_runtime.Value
var once_functorWithIndexLast sync.Once
func Get_functorWithIndexLast() gopurs_runtime.Value {
	once_functorWithIndexLast.Do(func() {
		functorWithIndexLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_1_0, v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
})})
	})
	return functorWithIndexLast
}

var functorWithIndexIdentity gopurs_runtime.Value
var once_functorWithIndexIdentity sync.Once
func Get_functorWithIndexIdentity() gopurs_runtime.Value {
	once_functorWithIndexIdentity.Do(func() {
		functorWithIndexIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), v_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
})})
	})
	return functorWithIndexIdentity
}

var functorWithIndexFirst gopurs_runtime.Value
var once_functorWithIndexFirst sync.Once
func Get_functorWithIndexFirst() gopurs_runtime.Value {
	once_functorWithIndexFirst.Do(func() {
		functorWithIndexFirst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(__local_var_1_0, v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
})})
	})
	return functorWithIndexFirst
}

var functorWithIndexEither gopurs_runtime.Value
var once_functorWithIndexEither sync.Once
func Get_functorWithIndexEither() gopurs_runtime.Value {
	once_functorWithIndexEither.Do(func() {
		functorWithIndexEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(__local_var_1_0, m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
})})
	})
	return functorWithIndexEither
}

var functorWithIndexDual gopurs_runtime.Value
var once_functorWithIndexDual sync.Once
func Get_functorWithIndexDual() gopurs_runtime.Value {
	once_functorWithIndexDual.Do(func() {
		functorWithIndexDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
})})
	})
	return functorWithIndexDual
}

var functorWithIndexDisj gopurs_runtime.Value
var once_functorWithIndexDisj sync.Once
func Get_functorWithIndexDisj() gopurs_runtime.Value {
	once_functorWithIndexDisj.Do(func() {
		functorWithIndexDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
})})
	})
	return functorWithIndexDisj
}

var functorWithIndexCoproduct gopurs_runtime.Value
var once_functorWithIndexCoproduct sync.Once
func Get_functorWithIndexCoproduct() gopurs_runtime.Value {
	once_functorWithIndexCoproduct.Do(func() {
		functorWithIndexCoproduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictFunctorWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCoproduct1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
__local_var_7_4 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(__local_var_6_3, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(__local_var_7_4, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_6 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_7}))
}))
__local_var_8_7 := gopurs_runtime.Apply(dictFunctorWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_8}))
}))
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(__local_var_7_6, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(__local_var_8_7, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t8
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_4_2
})})
})
})
	})
	return functorWithIndexCoproduct
}

var functorWithIndexConst gopurs_runtime.Value
var once_functorWithIndexConst sync.Once
func Get_functorWithIndexConst() gopurs_runtime.Value {
	once_functorWithIndexConst.Do(func() {
		functorWithIndexConst = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
})})
	})
	return functorWithIndexConst
}

var functorWithIndexConj gopurs_runtime.Value
var once_functorWithIndexConj sync.Once
func Get_functorWithIndexConj() gopurs_runtime.Value {
	once_functorWithIndexConj.Do(func() {
		functorWithIndexConj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
})})
	})
	return functorWithIndexConj
}

var functorWithIndexCompose gopurs_runtime.Value
var once_functorWithIndexCompose sync.Once
func Get_functorWithIndexCompose() gopurs_runtime.Value {
	once_functorWithIndexCompose.Do(func() {
		functorWithIndexCompose = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictFunctorWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCompose1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)), v_5)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctorWithIndex1_2.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_7, "value1": b_8}))
}))
})), v_6)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_4_2
})})
})
})
	})
	return functorWithIndexCompose
}

var functorWithIndexArray gopurs_runtime.Value
var once_functorWithIndexArray sync.Once
func Get_functorWithIndexArray() gopurs_runtime.Value {
	once_functorWithIndexArray.Do(func() {
		functorWithIndexArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": Get_mapWithIndexArray(), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
})})
	})
	return functorWithIndexArray
}

var functorWithIndexApp gopurs_runtime.Value
var once_functorWithIndexApp sync.Once
func Get_functorWithIndexApp() gopurs_runtime.Value {
	once_functorWithIndexApp.Do(func() {
		functorWithIndexApp = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], f_2), v_3)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})})
})
	})
	return functorWithIndexApp
}

var functorWithIndexAdditive gopurs_runtime.Value
var once_functorWithIndexAdditive sync.Once
func Get_functorWithIndexAdditive() gopurs_runtime.Value {
	once_functorWithIndexAdditive.Do(func() {
		functorWithIndexAdditive = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
})})
	})
	return functorWithIndexAdditive
}

func Get_mapWithIndexArray() gopurs_runtime.Value {
	return MapWithIndexArray
}
