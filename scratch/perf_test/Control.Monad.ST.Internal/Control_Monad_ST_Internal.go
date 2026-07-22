package Control_Monad_ST_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		modify_prime = Get_modifyImpl()
	})
	return modify_prime
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": s_prime_2_0, "value": s_prime_2_0})
}))
})
	})
	return modify
}

var functorST gopurs_runtime.Value
var once_functorST sync.Once
func Get_functorST() gopurs_runtime.Value {
	once_functorST.Do(func() {
		functorST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": Get_map_()})
	})
	return functorST
}

var void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		void = gopurs_runtime.Apply(Get_map_(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return void
}

var monadST gopurs_runtime.Value
var once_monadST sync.Once
func Get_monadST() gopurs_runtime.Value {
	once_monadST.Do(func() {
		monadST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
})})
	})
	return monadST
}

var bindST gopurs_runtime.Value
var once_bindST sync.Once
func Get_bindST() gopurs_runtime.Value {
	once_bindST.Do(func() {
		bindST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": Get_bind_(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
})})
	})
	return bindST
}

var applyST gopurs_runtime.Value
var once_applyST sync.Once
func Get_applyST() gopurs_runtime.Value {
	once_applyST.Do(func() {
		applyST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), f_0), gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), a_1), gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeST().PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_prime_2, a_prime_3))
}))
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
})})
	})
	return applyST
}

var applicativeST gopurs_runtime.Value
var once_applicativeST sync.Once
func Get_applicativeST() gopurs_runtime.Value {
	once_applicativeST.Do(func() {
		applicativeST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_pure_(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
})})
	})
	return applicativeST
}

var lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		lift2 = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyST().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_map_(), f_0), a_1)), b_2)
})
})
})
	})
	return lift2
}

var semigroupST gopurs_runtime.Value
var once_semigroupST sync.Once
func Get_semigroupST() gopurs_runtime.Value {
	once_semigroupST.Do(func() {
		semigroupST = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_lift2(), dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"])})
})
	})
	return semigroupST
}

var monadRecST gopurs_runtime.Value
var once_monadRecST sync.Once
func Get_monadRecST() gopurs_runtime.Value {
	once_monadRecST.Do(func() {
		monadRecST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
fromDone_2_0 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t1 = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t1
})
}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), gopurs_runtime.Apply(f_0, a_1)), Get_new_())), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_while(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_map_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")
})), gopurs_runtime.Apply(Get_read(), r_3))), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), gopurs_runtime.Apply(Get_read(), r_3)), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bind_(), gopurs_runtime.Apply(f_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_write(), e_5), r_3))
}))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(Get_pure_(), pkg_Data_Unit.Get_unit())
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_map_(), fromDone_2_0), gopurs_runtime.Apply(Get_read(), r_3))
}))
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadST()
})})
	})
	return monadRecST
}

var monoidST gopurs_runtime.Value
var once_monoidST sync.Once
func Get_monoidST() gopurs_runtime.Value {
	once_monoidST.Do(func() {
		monoidST = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(Get_pure_(), dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"])})
})})
})
	})
	return monoidST
}

func Get_bind_() gopurs_runtime.Value {
	return Bind_
}

func Get_for_() gopurs_runtime.Value {
	return For_
}

func Get_foreach() gopurs_runtime.Value {
	return Foreach
}

func Get_map_() gopurs_runtime.Value {
	return Map_
}

func Get_modifyImpl() gopurs_runtime.Value {
	return ModifyImpl
}

func Get_new_() gopurs_runtime.Value {
	return New_
}

func Get_pure_() gopurs_runtime.Value {
	return Pure_
}

func Get_read() gopurs_runtime.Value {
	return Read
}

func Get_run() gopurs_runtime.Value {
	return Run
}

func Get_while() gopurs_runtime.Value {
	return While
}

func Get_write() gopurs_runtime.Value {
	return Write
}
