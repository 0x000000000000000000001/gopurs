package Data_Monoid

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var monoidUnit gopurs_runtime.Value
var once_monoidUnit sync.Once
func Get_monoidUnit() gopurs_runtime.Value {
	once_monoidUnit.Do(func() {
		monoidUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": pkg_Data_Unit.Get_unit(), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupUnit()
})})
	})
	return monoidUnit
}

var monoidString gopurs_runtime.Value
var once_monoidString sync.Once
func Get_monoidString() gopurs_runtime.Value {
	once_monoidString.Do(func() {
		monoidString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Str(""), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupString()
})})
	})
	return monoidString
}

var monoidRecordNil gopurs_runtime.Value
var once_monoidRecordNil sync.Once
func Get_monoidRecordNil() gopurs_runtime.Value {
	once_monoidRecordNil.Do(func() {
		monoidRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"memptyRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
}), "SemigroupRecord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupRecordNil()
})})
	})
	return monoidRecordNil
}

var monoidOrdering gopurs_runtime.Value
var once_monoidOrdering sync.Once
func Get_monoidOrdering() gopurs_runtime.Value {
	once_monoidOrdering.Do(func() {
		monoidOrdering = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ordering.Get_semigroupOrdering()
})})
	})
	return monoidOrdering
}

var monoidArray gopurs_runtime.Value
var once_monoidArray sync.Once
func Get_monoidArray() gopurs_runtime.Value {
	once_monoidArray.Do(func() {
		monoidArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Array([]gopurs_runtime.Value{}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupArray()
})})
	})
	return monoidArray
}

var memptyRecord gopurs_runtime.Value
var once_memptyRecord sync.Once
func Get_memptyRecord() gopurs_runtime.Value {
	once_memptyRecord.Do(func() {
		memptyRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["memptyRecord"]
})
	})
	return memptyRecord
}

var monoidRecord gopurs_runtime.Value
var once_monoidRecord sync.Once
func Get_monoidRecord() gopurs_runtime.Value {
	once_monoidRecord.Do(func() {
		monoidRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoidRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRecord1_2_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoidRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["SemigroupRecord0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["appendRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(dictMonoidRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["memptyRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRecord1_2_0
})})
})
})
	})
	return monoidRecord
}

var mempty gopurs_runtime.Value
var once_mempty sync.Once
func Get_mempty() gopurs_runtime.Value {
	once_mempty.Do(func() {
		mempty = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
})
	})
	return mempty
}

var monoidFn gopurs_runtime.Value
var once_monoidFn sync.Once
func Get_monoidFn() gopurs_runtime.Value {
	once_monoidFn.Do(func() {
		monoidFn = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
__local_var_2_1 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
semigroupFn_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_3, x_5)), gopurs_runtime.Apply(g_4, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty1_1_0
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_3_2
})})
})
	})
	return monoidFn
}

var monoidRecordCons gopurs_runtime.Value
var once_monoidRecordCons sync.Once
func Get_monoidRecordCons() gopurs_runtime.Value {
	once_monoidRecordCons.Do(func() {
		monoidRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_2_0 := dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
Semigroup0_3_1 := gopurs_runtime.Apply(dictMonoid_1.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoidRecord_5 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRecordCons1_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_semigroupRecordCons(), dictIsSymbol_0), gopurs_runtime.Value{}), gopurs_runtime.Apply(dictMonoidRecord_5.PtrVal.(map[string]gopurs_runtime.Value)["SemigroupRecord0"], gopurs_runtime.Value{})), Semigroup0_3_1)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"memptyRecord": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), mempty1_2_0), gopurs_runtime.Apply(dictMonoidRecord_5.PtrVal.(map[string]gopurs_runtime.Value)["memptyRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})))
}), "SemigroupRecord0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRecordCons1_6_2
})})
})
})
})
})
	})
	return monoidRecordCons
}

var power gopurs_runtime.Value
var once_power sync.Once
func Get_power() gopurs_runtime.Value {
	once_power.Do(func() {
		power = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
__local_var_2_1 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], p_5), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t4 = mempty1_1_0
} else {
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), p_5), gopurs_runtime.Int(1))).IntVal != 0 {
__t4 = x_3
} else {
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), p_5), gopurs_runtime.Int(2))), gopurs_runtime.Int(0))).IntVal != 0 {
x_prime_6_5 := gopurs_runtime.Apply(go__4_2, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intDiv(), p_5), gopurs_runtime.Int(2)))
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], x_prime_6_5), x_prime_6_5)
} else {
x_prime_6_3 := gopurs_runtime.Apply(go__4_2, gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intDiv(), p_5), gopurs_runtime.Int(2)))
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], x_prime_6_3), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], x_prime_6_3), x_3))
}
}
}
return __t4
})
return go__4_2
})
})
	})
	return power
}

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (v_2).IntVal != 0 {
__t1 = v1_3
} else {
__t1 = mempty1_1_0
}
return __t1
})
})
})
	})
	return guard
}


