package Data_Decidable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Decide "gopurs/output/Data.Decide"
	pkg_Data_Divisible "gopurs/output/Data.Divisible"
)

var lose gopurs_runtime.Value
var once_lose sync.Once
func Get_lose() gopurs_runtime.Value {
	once_lose.Do(func() {
		lose = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["lose"]
})
	})
	return lose
}

var lost gopurs_runtime.Value
var once_lost sync.Once
func Get_lost() gopurs_runtime.Value {
	once_lost.Do(func() {
		lost = gopurs_runtime.Func(func(dictDecidable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictDecidable_0.PtrVal.(map[string]gopurs_runtime.Value)["lose"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return lost
}

var decidablePredicate gopurs_runtime.Value
var once_decidablePredicate sync.Once
func Get_decidablePredicate() gopurs_runtime.Value {
	once_decidablePredicate.Do(func() {
		decidablePredicate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_2_0 gopurs_runtime.Value
spin_2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(spin_2_0, v_3)
})
return gopurs_runtime.Apply(spin_2_0, gopurs_runtime.Apply(f_0, a_1))
})
}), "Decide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Decide.Get_choosePredicate()
}), "Divisible1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divisible.Get_divisiblePredicate()
})})
	})
	return decidablePredicate
}

var decidableOp gopurs_runtime.Value
var once_decidableOp sync.Once
func Get_decidableOp() gopurs_runtime.Value {
	once_decidableOp.Do(func() {
		decidableOp = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
chooseOp_1_0 := gopurs_runtime.Apply(pkg_Data_Decide.Get_chooseOp(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
divisibleOp_2_1 := gopurs_runtime.Apply(pkg_Data_Divisible.Get_divisibleOp(), dictMonoid_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lose": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_5_2 gopurs_runtime.Value
spin_5_2 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(spin_5_2, v_6)
})
return gopurs_runtime.Apply(spin_5_2, gopurs_runtime.Apply(f_3, a_4))
})
}), "Decide0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return chooseOp_1_0
}), "Divisible1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return divisibleOp_2_1
})})
})
	})
	return decidableOp
}

var decidableEquivalence gopurs_runtime.Value
var once_decidableEquivalence sync.Once
func Get_decidableEquivalence() gopurs_runtime.Value {
	once_decidableEquivalence.Do(func() {
		decidableEquivalence = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_2_0 gopurs_runtime.Value
spin_2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(spin_2_0, v_3)
})
return gopurs_runtime.Apply(spin_2_0, gopurs_runtime.Apply(f_0, a_1))
})
}), "Decide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Decide.Get_chooseEquivalence()
}), "Divisible1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divisible.Get_divisibleEquivalence()
})})
	})
	return decidableEquivalence
}

var decidableComparison gopurs_runtime.Value
var once_decidableComparison sync.Once
func Get_decidableComparison() gopurs_runtime.Value {
	once_decidableComparison.Do(func() {
		decidableComparison = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lose": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_3_0 gopurs_runtime.Value
spin_3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(spin_3_0, v_4)
})
return gopurs_runtime.Apply(spin_3_0, gopurs_runtime.Apply(f_0, a_1))
})
})
}), "Decide0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Decide.Get_chooseComparison()
}), "Divisible1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divisible.Get_divisibleComparison()
})})
	})
	return decidableComparison
}


