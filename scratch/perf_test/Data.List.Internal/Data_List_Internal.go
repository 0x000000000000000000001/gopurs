package Data_List_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		Leaf = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
	})
	return Leaf
}

var Two gopurs_runtime.Value
var once_Two sync.Once
func Get_Two() gopurs_runtime.Value {
	once_Two.Do(func() {
		Two = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": value0, "value1": value1, "value2": value2})
})
})
})
	})
	return Two
}

var Three gopurs_runtime.Value
var once_Three sync.Once
func Get_Three() gopurs_runtime.Value {
	once_Three.Do(func() {
		Three = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Three"), "value0": value0, "value1": value1, "value2": value2, "value3": value3, "value4": value4})
})
})
})
})
})
	})
	return Three
}

var TwoLeft gopurs_runtime.Value
var once_TwoLeft sync.Once
func Get_TwoLeft() gopurs_runtime.Value {
	once_TwoLeft.Do(func() {
		TwoLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("TwoLeft"), "value0": value0, "value1": value1})
})
})
	})
	return TwoLeft
}

var TwoRight gopurs_runtime.Value
var once_TwoRight sync.Once
func Get_TwoRight() gopurs_runtime.Value {
	once_TwoRight.Do(func() {
		TwoRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("TwoRight"), "value0": value0, "value1": value1})
})
})
	})
	return TwoRight
}

var ThreeLeft gopurs_runtime.Value
var once_ThreeLeft sync.Once
func Get_ThreeLeft() gopurs_runtime.Value {
	once_ThreeLeft.Do(func() {
		ThreeLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ThreeLeft"), "value0": value0, "value1": value1, "value2": value2, "value3": value3})
})
})
})
})
	})
	return ThreeLeft
}

var ThreeMiddle gopurs_runtime.Value
var once_ThreeMiddle sync.Once
func Get_ThreeMiddle() gopurs_runtime.Value {
	once_ThreeMiddle.Do(func() {
		ThreeMiddle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ThreeMiddle"), "value0": value0, "value1": value1, "value2": value2, "value3": value3})
})
})
})
})
	})
	return ThreeMiddle
}

var ThreeRight gopurs_runtime.Value
var once_ThreeRight sync.Once
func Get_ThreeRight() gopurs_runtime.Value {
	once_ThreeRight.Do(func() {
		ThreeRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ThreeRight"), "value0": value0, "value1": value1, "value2": value2, "value3": value3})
})
})
})
})
	})
	return ThreeRight
}

var KickUp gopurs_runtime.Value
var once_KickUp sync.Once
func Get_KickUp() gopurs_runtime.Value {
	once_KickUp.Do(func() {
		KickUp = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("KickUp"), "value0": value0, "value1": value1, "value2": value2})
})
})
})
	})
	return KickUp
}

var fromZipper gopurs_runtime.Value
var once_fromZipper sync.Once
func Get_fromZipper() gopurs_runtime.Value {
	once_fromZipper.Do(func() {
		fromZipper = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = v1_1
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "TwoLeft")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v1_1, "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "TwoRight")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_1}))
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ThreeLeft")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Three"), "value0": v1_1, "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value3": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value4": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]}))
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ThreeMiddle")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Three"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_1, "value3": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value4": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]}))
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ThreeRight")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Three"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value4": v1_1}))
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
}
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}
}()
})
})
	})
	return fromZipper
}

var insertAndLookupBy gopurs_runtime.Value
var once_insertAndLookupBy sync.Once
func Get_insertAndLookupBy() gopurs_runtime.Value {
	once_insertAndLookupBy.Do(func() {
		insertAndLookupBy = gopurs_runtime.Func(func(comp_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(orig_2 gopurs_runtime.Value) gopurs_runtime.Value {
var up_3_0 gopurs_runtime.Value
up_3_0 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "TwoLeft")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Three"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value4": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "TwoRight")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_fromZipper(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Three"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value4": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ThreeLeft")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(up_3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("KickUp"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}), "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value2": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})}))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ThreeMiddle")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(up_3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("KickUp"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}), "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value2": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})}))
} else {
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "ThreeRight")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(up_3_0, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("KickUp"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"]}), "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Two"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})}))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
}
}
__t1 = __t2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
var down_4_3 gopurs_runtime.Value
down_4_3 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"found": gopurs_runtime.Bool(false), "result": gopurs_runtime.Apply(gopurs_runtime.Apply(up_3_0, v_5), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("KickUp"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value1": k_1, "value2": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})}))})
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Two")).IntVal != 0 {
v2_7_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(comp_0, k_1), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t10 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_7_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"found": gopurs_runtime.Bool(true), "result": orig_2})
} else {
if (gopurs_runtime.Bool(v2_7_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Apply(down_4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("TwoLeft"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}), "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Apply(down_4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("TwoRight"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
}
}
__t4 = __t10
} else {
if (gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Three")).IntVal != 0 {
v2_7_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(comp_0, k_1), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_7_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"found": gopurs_runtime.Bool(true), "result": orig_2})
} else {
v3_8_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(comp_0, k_1), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(v3_8_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"found": gopurs_runtime.Bool(true), "result": orig_2})
} else {
if (gopurs_runtime.Bool(v2_7_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(down_4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ThreeLeft"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value2": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value3": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value4"]}), "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v2_7_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal != 0 && gopurs_runtime.Bool(v3_8_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal != 0)).IntVal != 0 {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(down_4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ThreeMiddle"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value3": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value4"]}), "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
} else {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Apply(down_4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("ThreeRight"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value3"]}), "value1": v_5})), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value4"])
}
}
}
__t8 = __t7
}
__t4 = __t8
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(down_4_3, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})), orig_2)
})
})
})
	})
	return insertAndLookupBy
}

var emptySet gopurs_runtime.Value
var once_emptySet sync.Once
func Get_emptySet() gopurs_runtime.Value {
	once_emptySet.Do(func() {
		emptySet = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
	})
	return emptySet
}


