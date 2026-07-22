package Type_Data_Boolean

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var reflectBoolean gopurs_runtime.Value
var once_reflectBoolean sync.Once
func Get_reflectBoolean() gopurs_runtime.Value {
	once_reflectBoolean.Do(func() {
		reflectBoolean = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectBoolean"]
})
	})
	return reflectBoolean
}

var orTrue gopurs_runtime.Value
var once_orTrue sync.Once
func Get_orTrue() gopurs_runtime.Value {
	once_orTrue.Do(func() {
		orTrue = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return orTrue
}

var orFalse gopurs_runtime.Value
var once_orFalse sync.Once
func Get_orFalse() gopurs_runtime.Value {
	once_orFalse.Do(func() {
		orFalse = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return orFalse
}

var or gopurs_runtime.Value
var once_or sync.Once
func Get_or() gopurs_runtime.Value {
	once_or.Do(func() {
		or = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
})
	})
	return or
}

var notTrue gopurs_runtime.Value
var once_notTrue sync.Once
func Get_notTrue() gopurs_runtime.Value {
	once_notTrue.Do(func() {
		notTrue = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return notTrue
}

var notFalse gopurs_runtime.Value
var once_notFalse sync.Once
func Get_notFalse() gopurs_runtime.Value {
	once_notFalse.Do(func() {
		notFalse = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return notFalse
}

var not gopurs_runtime.Value
var once_not sync.Once
func Get_not() gopurs_runtime.Value {
	once_not.Do(func() {
		not = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
	})
	return not
}

var isBooleanTrue gopurs_runtime.Value
var once_isBooleanTrue sync.Once
func Get_isBooleanTrue() gopurs_runtime.Value {
	once_isBooleanTrue.Do(func() {
		isBooleanTrue = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectBoolean": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})})
	})
	return isBooleanTrue
}

var isBooleanFalse gopurs_runtime.Value
var once_isBooleanFalse sync.Once
func Get_isBooleanFalse() gopurs_runtime.Value {
	once_isBooleanFalse.Do(func() {
		isBooleanFalse = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectBoolean": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})})
	})
	return isBooleanFalse
}

var reifyBoolean gopurs_runtime.Value
var once_reifyBoolean sync.Once
func Get_reifyBoolean() gopurs_runtime.Value {
	once_reifyBoolean.Do(func() {
		reifyBoolean = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (v_0).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, Get_isBooleanTrue()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
} else {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, Get_isBooleanFalse()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
}
return __t0
})
})
	})
	return reifyBoolean
}

var if_ gopurs_runtime.Value
var once_if_ sync.Once
func Get_if_() gopurs_runtime.Value {
	once_if_.Do(func() {
		if_ = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
})
})
	})
	return if_
}

var ifTrue gopurs_runtime.Value
var once_ifTrue sync.Once
func Get_ifTrue() gopurs_runtime.Value {
	once_ifTrue.Do(func() {
		ifTrue = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return ifTrue
}

var ifFalse gopurs_runtime.Value
var once_ifFalse sync.Once
func Get_ifFalse() gopurs_runtime.Value {
	once_ifFalse.Do(func() {
		ifFalse = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return ifFalse
}

var andTrue gopurs_runtime.Value
var once_andTrue sync.Once
func Get_andTrue() gopurs_runtime.Value {
	once_andTrue.Do(func() {
		andTrue = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return andTrue
}

var andFalse gopurs_runtime.Value
var once_andFalse sync.Once
func Get_andFalse() gopurs_runtime.Value {
	once_andFalse.Do(func() {
		andFalse = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return andFalse
}

var and gopurs_runtime.Value
var once_and sync.Once
func Get_and() gopurs_runtime.Value {
	once_and.Do(func() {
		and = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})
})
	})
	return and
}


