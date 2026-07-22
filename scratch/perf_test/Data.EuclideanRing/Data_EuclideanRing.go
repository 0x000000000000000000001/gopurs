package Data_EuclideanRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_CommutativeRing "gopurs/output/Data.CommutativeRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var mod gopurs_runtime.Value
var once_mod sync.Once
func Get_mod() gopurs_runtime.Value {
	once_mod.Do(func() {
		mod = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mod"]
})
	})
	return mod
}

var gcd gopurs_runtime.Value
var once_gcd sync.Once
func Get_gcd() gopurs_runtime.Value {
	once_gcd.Do(func() {
		gcd = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEuclideanRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
zero_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_1.PtrVal.(map[string]gopurs_runtime.Value)["CommutativeRing0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["zero"]
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], b_4), zero_2_0)).IntVal != 0 {
__t1 = a_3
} else {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_gcd(), dictEq_0), dictEuclideanRing_1), b_4), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_1.PtrVal.(map[string]gopurs_runtime.Value)["mod"], a_3), b_4))
}
return __t1
})
})
}
}()
})
})
	})
	return gcd
}

var euclideanRingNumber gopurs_runtime.Value
var once_euclideanRingNumber sync.Once
func Get_euclideanRingNumber() gopurs_runtime.Value {
	once_euclideanRingNumber.Do(func() {
		euclideanRingNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"degree": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(1)
}), "div": Get_numDiv(), "mod": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(0.0)
})
}), "CommutativeRing0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_CommutativeRing.Get_commutativeRingNumber()
})})
	})
	return euclideanRingNumber
}

var euclideanRingInt gopurs_runtime.Value
var once_euclideanRingInt sync.Once
func Get_euclideanRingInt() gopurs_runtime.Value {
	once_euclideanRingInt.Do(func() {
		euclideanRingInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"degree": Get_intDegree(), "div": Get_intDiv(), "mod": Get_intMod(), "CommutativeRing0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_CommutativeRing.Get_commutativeRingInt()
})})
	})
	return euclideanRingInt
}

var div gopurs_runtime.Value
var once_div sync.Once
func Get_div() gopurs_runtime.Value {
	once_div.Do(func() {
		div = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["div"]
})
	})
	return div
}

var lcm gopurs_runtime.Value
var once_lcm sync.Once
func Get_lcm() gopurs_runtime.Value {
	once_lcm.Do(func() {
		lcm = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEuclideanRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
Semiring0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_1.PtrVal.(map[string]gopurs_runtime.Value)["CommutativeRing0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{})
zero_3_1 := Semiring0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
gcd2_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_gcd(), dictEq_0), dictEuclideanRing_1)
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolDisj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], a_5), zero_3_1)), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], b_6), zero_3_1))).IntVal != 0 {
__t3 = zero_3_1
} else {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictEuclideanRing_1.PtrVal.(map[string]gopurs_runtime.Value)["div"], gopurs_runtime.Apply(gopurs_runtime.Apply(Semiring0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], a_5), b_6)), gopurs_runtime.Apply(gopurs_runtime.Apply(gcd2_4_2, a_5), b_6))
}
return __t3
})
})
})
})
	})
	return lcm
}

var degree gopurs_runtime.Value
var once_degree sync.Once
func Get_degree() gopurs_runtime.Value {
	once_degree.Do(func() {
		degree = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["degree"]
})
	})
	return degree
}

func Get_intDegree() gopurs_runtime.Value {
	return IntDegree
}

func Get_intDiv() gopurs_runtime.Value {
	return IntDiv
}

func Get_intMod() gopurs_runtime.Value {
	return IntMod
}

func Get_numDiv() gopurs_runtime.Value {
	return NumDiv
}
