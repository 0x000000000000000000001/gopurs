package Type_Data_Ordering

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var reflectOrdering gopurs_runtime.Value
var once_reflectOrdering sync.Once
func Get_reflectOrdering() gopurs_runtime.Value {
	once_reflectOrdering.Do(func() {
		reflectOrdering = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectOrdering"]
})
	})
	return reflectOrdering
}

var isOrderingLT gopurs_runtime.Value
var once_isOrderingLT sync.Once
func Get_isOrderingLT() gopurs_runtime.Value {
	once_isOrderingLT.Do(func() {
		isOrderingLT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectOrdering": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
})})
	})
	return isOrderingLT
}

var isOrderingGT gopurs_runtime.Value
var once_isOrderingGT sync.Once
func Get_isOrderingGT() gopurs_runtime.Value {
	once_isOrderingGT.Do(func() {
		isOrderingGT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectOrdering": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
})})
	})
	return isOrderingGT
}

var isOrderingEQ gopurs_runtime.Value
var once_isOrderingEQ sync.Once
func Get_isOrderingEQ() gopurs_runtime.Value {
	once_isOrderingEQ.Do(func() {
		isOrderingEQ = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectOrdering": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})})
	})
	return isOrderingEQ
}

var reifyOrdering gopurs_runtime.Value
var once_reifyOrdering sync.Once
func Get_reifyOrdering() gopurs_runtime.Value {
	once_reifyOrdering.Do(func() {
		reifyOrdering = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, Get_isOrderingLT()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, Get_isOrderingEQ()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, Get_isOrderingGT()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
})
})
	})
	return reifyOrdering
}

var invertOrderingLT gopurs_runtime.Value
var once_invertOrderingLT sync.Once
func Get_invertOrderingLT() gopurs_runtime.Value {
	once_invertOrderingLT.Do(func() {
		invertOrderingLT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return invertOrderingLT
}

var invertOrderingGT gopurs_runtime.Value
var once_invertOrderingGT sync.Once
func Get_invertOrderingGT() gopurs_runtime.Value {
	once_invertOrderingGT.Do(func() {
		invertOrderingGT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return invertOrderingGT
}

var invertOrderingEQ gopurs_runtime.Value
var once_invertOrderingEQ sync.Once
func Get_invertOrderingEQ() gopurs_runtime.Value {
	once_invertOrderingEQ.Do(func() {
		invertOrderingEQ = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return invertOrderingEQ
}

var invert gopurs_runtime.Value
var once_invert sync.Once
func Get_invert() gopurs_runtime.Value {
	once_invert.Do(func() {
		invert = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
	})
	return invert
}

var equalsLTLT gopurs_runtime.Value
var once_equalsLTLT sync.Once
func Get_equalsLTLT() gopurs_runtime.Value {
	once_equalsLTLT.Do(func() {
		equalsLTLT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsLTLT
}

var equalsLTGT gopurs_runtime.Value
var once_equalsLTGT sync.Once
func Get_equalsLTGT() gopurs_runtime.Value {
	once_equalsLTGT.Do(func() {
		equalsLTGT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsLTGT
}

var equalsLTEQ gopurs_runtime.Value
var once_equalsLTEQ sync.Once
func Get_equalsLTEQ() gopurs_runtime.Value {
	once_equalsLTEQ.Do(func() {
		equalsLTEQ = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsLTEQ
}

var equalsGTLT gopurs_runtime.Value
var once_equalsGTLT sync.Once
func Get_equalsGTLT() gopurs_runtime.Value {
	once_equalsGTLT.Do(func() {
		equalsGTLT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsGTLT
}

var equalsGTGT gopurs_runtime.Value
var once_equalsGTGT sync.Once
func Get_equalsGTGT() gopurs_runtime.Value {
	once_equalsGTGT.Do(func() {
		equalsGTGT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsGTGT
}

var equalsGTEQ gopurs_runtime.Value
var once_equalsGTEQ sync.Once
func Get_equalsGTEQ() gopurs_runtime.Value {
	once_equalsGTEQ.Do(func() {
		equalsGTEQ = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsGTEQ
}

var equalsEQLT gopurs_runtime.Value
var once_equalsEQLT sync.Once
func Get_equalsEQLT() gopurs_runtime.Value {
	once_equalsEQLT.Do(func() {
		equalsEQLT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsEQLT
}

var equalsEQGT gopurs_runtime.Value
var once_equalsEQGT sync.Once
func Get_equalsEQGT() gopurs_runtime.Value {
	once_equalsEQGT.Do(func() {
		equalsEQGT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsEQGT
}

var equalsEQEQ gopurs_runtime.Value
var once_equalsEQEQ sync.Once
func Get_equalsEQEQ() gopurs_runtime.Value {
	once_equalsEQEQ.Do(func() {
		equalsEQEQ = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return equalsEQEQ
}

var equals gopurs_runtime.Value
var once_equals sync.Once
func Get_equals() gopurs_runtime.Value {
	once_equals.Do(func() {
		equals = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
})
	})
	return equals
}

var appendOrderingLT gopurs_runtime.Value
var once_appendOrderingLT sync.Once
func Get_appendOrderingLT() gopurs_runtime.Value {
	once_appendOrderingLT.Do(func() {
		appendOrderingLT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return appendOrderingLT
}

var appendOrderingGT gopurs_runtime.Value
var once_appendOrderingGT sync.Once
func Get_appendOrderingGT() gopurs_runtime.Value {
	once_appendOrderingGT.Do(func() {
		appendOrderingGT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return appendOrderingGT
}

var appendOrderingEQ gopurs_runtime.Value
var once_appendOrderingEQ sync.Once
func Get_appendOrderingEQ() gopurs_runtime.Value {
	once_appendOrderingEQ.Do(func() {
		appendOrderingEQ = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return appendOrderingEQ
}

var append_ gopurs_runtime.Value
var once_append_ sync.Once
func Get_append_() gopurs_runtime.Value {
	once_append_.Do(func() {
		append_ = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
})
	})
	return append_
}


