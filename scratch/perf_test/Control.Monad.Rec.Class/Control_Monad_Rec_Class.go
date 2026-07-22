package Control_Monad_Rec_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Ref "gopurs/output/Effect.Ref"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var Loop gopurs_runtime.Value
var once_Loop sync.Once
func Get_Loop() gopurs_runtime.Value {
	once_Loop.Do(func() {
		Loop = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": value0})
})
	})
	return Loop
}

var Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		Done = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": value0})
})
	})
	return Done
}

var tailRecM gopurs_runtime.Value
var once_tailRecM sync.Once
func Get_tailRecM() gopurs_runtime.Value {
	once_tailRecM.Do(func() {
		tailRecM = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"]
})
	})
	return tailRecM
}

var tailRecM2 gopurs_runtime.Value
var once_tailRecM2 sync.Once
func Get_tailRecM2() gopurs_runtime.Value {
	once_tailRecM2.Do(func() {
		tailRecM2 = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, o_4.PtrVal.(map[string]gopurs_runtime.Value)["a"]), o_4.PtrVal.(map[string]gopurs_runtime.Value)["b"])
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": a_2, "b": b_3}))
})
})
})
})
	})
	return tailRecM2
}

var tailRecM3 gopurs_runtime.Value
var once_tailRecM3 sync.Once
func Get_tailRecM3() gopurs_runtime.Value {
	once_tailRecM3.Do(func() {
		tailRecM3 = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(o_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, o_5.PtrVal.(map[string]gopurs_runtime.Value)["a"]), o_5.PtrVal.(map[string]gopurs_runtime.Value)["b"]), o_5.PtrVal.(map[string]gopurs_runtime.Value)["c"])
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": a_2, "b": b_3, "c": c_4}))
})
})
})
})
})
	})
	return tailRecM3
}

var untilJust gopurs_runtime.Value
var once_untilJust sync.Once
func Get_untilJust() gopurs_runtime.Value {
	once_untilJust.Do(func() {
		untilJust = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": pkg_Data_Unit.Get_unit()})
} else {
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})), m_2)
})), pkg_Data_Unit.Get_unit())
})
})
	})
	return untilJust
}

var whileJust gopurs_runtime.Value
var once_whileJust sync.Once
func Get_whileJust() gopurs_runtime.Value {
	once_whileJust.Do(func() {
		whileJust = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(dictMonadRec_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_2.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_2.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": v_5})
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], v_5), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})), m_4)
})), mempty_1_0)
})
})
})
	})
	return whileJust
}

var tailRec gopurs_runtime.Value
var once_tailRec sync.Once
func Get_tailRec() gopurs_runtime.Value {
	once_tailRec.Do(func() {
		tailRec = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t1 = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(f_0, x_2))
})
})
	})
	return tailRec
}

var tailRec2 gopurs_runtime.Value
var once_tailRec2 sync.Once
func Get_tailRec2() gopurs_runtime.Value {
	once_tailRec2.Do(func() {
		tailRec2 = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["a"]), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["b"]))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t1 = v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
return gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a_1), b_2))
})
})
})
	})
	return tailRec2
}

var tailRec3 gopurs_runtime.Value
var once_tailRec3 sync.Once
func Get_tailRec3() gopurs_runtime.Value {
	once_tailRec3.Do(func() {
		tailRec3 = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_0 gopurs_runtime.Value
go__4_0 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__4_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["a"]), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["b"]), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["c"]))
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t1 = v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
return gopurs_runtime.Apply(go__4_0, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a_1), b_2), c_3))
})
})
})
})
	})
	return tailRec3
}

var monadRecMaybe gopurs_runtime.Value
var once_monadRecMaybe sync.Once
func Get_monadRecMaybe() gopurs_runtime.Value {
	once_monadRecMaybe.Do(func() {
		monadRecMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Value{PtrVal: func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
}}
var go__3_3 gopurs_runtime.Value
go__3_3 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(go__3_3, __local_var_2_0.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t4 = v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
return gopurs_runtime.Apply(go__3_3, __local_var_2_0.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, a0_1)))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_monadMaybe()
})})
	})
	return monadRecMaybe
}

var monadRecIdentity gopurs_runtime.Value
var once_monadRecIdentity sync.Once
func Get_monadRecIdentity() gopurs_runtime.Value {
	once_monadRecIdentity.Do(func() {
		monadRecIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t1 = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Apply(f_0, x_2))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_monadIdentity()
})})
	})
	return monadRecIdentity
}

var monadRecFunction gopurs_runtime.Value
var once_monadRecFunction sync.Once
func Get_monadRecFunction() gopurs_runtime.Value {
	once_monadRecFunction.Do(func() {
		monadRecFunction = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), e_2))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t1 = v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
return gopurs_runtime.Apply(go__3_0, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, a0_1), e_2))
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
})})
	})
	return monadRecFunction
}

var monadRecEither gopurs_runtime.Value
var once_monadRecEither sync.Once
func Get_monadRecEither() gopurs_runtime.Value {
	once_monadRecEither.Do(func() {
		monadRecEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Value{PtrVal: func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
}}
var go__3_3 gopurs_runtime.Value
go__3_3 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(go__3_3, __local_var_2_0.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t4 = v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
return gopurs_runtime.Apply(go__3_3, __local_var_2_0.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, a0_1)))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_monadEither()
})})
	})
	return monadRecEither
}

var monadRecEffect gopurs_runtime.Value
var once_monadRecEffect sync.Once
func Get_monadRecEffect() gopurs_runtime.Value {
	once_monadRecEffect.Do(func() {
		monadRecEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(f_0, a_1)), pkg_Effect_Ref.Get__new())), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(pkg_Effect.Get_untilE(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), r_3)), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(f_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Ref.Get_write(), e_5), r_3)), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect.Get_pureE(), gopurs_runtime.Bool(false))
}))
}))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(pkg_Effect.Get_pureE(), gopurs_runtime.Bool(true))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_applyEffect().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(pkg_Effect.Get_pureE(), fromDone_2_0)), gopurs_runtime.Apply(pkg_Effect_Ref.Get_read(), r_3))
}))
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
})})
	})
	return monadRecEffect
}

var loop3 gopurs_runtime.Value
var once_loop3 sync.Once
func Get_loop3() gopurs_runtime.Value {
	once_loop3.Do(func() {
		loop3 = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": a_0, "b": b_1, "c": c_2})})
})
})
})
	})
	return loop3
}

var loop2 gopurs_runtime.Value
var once_loop2 sync.Once
func Get_loop2() gopurs_runtime.Value {
	once_loop2.Do(func() {
		loop2 = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": a_0, "b": b_1})})
})
})
	})
	return loop2
}

var functorStep gopurs_runtime.Value
var once_functorStep sync.Once
func Get_functorStep() gopurs_runtime.Value {
	once_functorStep.Do(func() {
		functorStep = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": m_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Apply(f_0, m_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})})
	})
	return functorStep
}

var forever gopurs_runtime.Value
var once_forever sync.Once
func Get_forever() gopurs_runtime.Value {
	once_forever.Do(func() {
		forever = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(ma_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": u_3})
})), ma_2)
})), pkg_Data_Unit.Get_unit())
})
})
	})
	return forever
}

var bifunctorStep gopurs_runtime.Value
var once_bifunctorStep sync.Once
func Get_bifunctorStep() gopurs_runtime.Value {
	once_bifunctorStep.Do(func() {
		bifunctorStep = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})})
	})
	return bifunctorStep
}


