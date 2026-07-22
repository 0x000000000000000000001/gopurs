package Control_Monad_Gen_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0), y_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = y_1
} else {
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = x_0
} else {
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = x_0
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
	})
	return max
}

var genTuple gopurs_runtime.Value
var once_genTuple sync.Once
func Get_genTuple() gopurs_runtime.Value {
	once_genTuple.Do(func() {
		genTuple = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_Tuple()), a_1)), b_2)
})
})
})
	})
	return genTuple
}

var genNonEmpty gopurs_runtime.Value
var once_genNonEmpty sync.Once
func Get_genNonEmpty() gopurs_runtime.Value {
	once_genNonEmpty.Do(func() {
		genNonEmpty = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
unfoldable1_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0), dictMonadGen_1)
return gopurs_runtime.Func(func(dictUnfoldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
unfoldable2_5_2 := gopurs_runtime.Apply(unfoldable1_3_1, dictUnfoldable_4)
return gopurs_runtime.Func(func(gen_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_NonEmpty.Get_NonEmpty()), gen_6)), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["resize"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_max(), gopurs_runtime.Int(0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), x_7), gopurs_runtime.Int(1)))
})), gopurs_runtime.Apply(unfoldable2_5_2, gen_6)))
})
})
})
})
	})
	return genNonEmpty
}

var genMaybe_prime gopurs_runtime.Value
var once_genMaybe_prime sync.Once
func Get_genMaybe_prime() gopurs_runtime.Value {
	once_genMaybe_prime.Do(func() {
		genMaybe_prime = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
Bind1_2_1 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(bias_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(gen_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseFloat"], gopurs_runtime.Float(0.0)), gopurs_runtime.Float(1.0))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_5), bias_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Maybe.Get_Just()), gen_4)
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
}
return __t2
}))
})
})
})
	})
	return genMaybe_prime
}

var genMaybe gopurs_runtime.Value
var once_genMaybe sync.Once
func Get_genMaybe() gopurs_runtime.Value {
	once_genMaybe.Do(func() {
		genMaybe = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genMaybe_prime(), dictMonadGen_0), gopurs_runtime.Float(0.75))
})
	})
	return genMaybe
}

var genIdentity gopurs_runtime.Value
var once_genIdentity sync.Once
func Get_genIdentity() gopurs_runtime.Value {
	once_genIdentity.Do(func() {
		genIdentity = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Identity.Get_Identity())
})
	})
	return genIdentity
}

var genEither_prime gopurs_runtime.Value
var once_genEither_prime sync.Once
func Get_genEither_prime() gopurs_runtime.Value {
	once_genEither_prime.Do(func() {
		genEither_prime = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(bias_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genA_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genB_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseFloat"], gopurs_runtime.Float(0.0)), gopurs_runtime.Float(1.0))), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_6), bias_3).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Left()), genA_4)
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Right()), genB_5)
}
return __t2
}))
})
})
})
})
	})
	return genEither_prime
}

var genEither gopurs_runtime.Value
var once_genEither sync.Once
func Get_genEither() gopurs_runtime.Value {
	once_genEither.Do(func() {
		genEither = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genEither_prime(), dictMonadGen_0), gopurs_runtime.Float(0.5))
})
	})
	return genEither
}


