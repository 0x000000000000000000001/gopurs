package ForeignKinds_Lib

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var NatProxy gopurs_runtime.Value
var once_NatProxy sync.Once
func Get_NatProxy() gopurs_runtime.Value {
	once_NatProxy.Do(func() {
		NatProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NatProxy")})
	})
	return NatProxy
}

var proxy2 gopurs_runtime.Value
var once_proxy2 sync.Once
func Get_proxy2() gopurs_runtime.Value {
	once_proxy2.Do(func() {
		proxy2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NatProxy")})
	})
	return proxy2
}

var proxy1 gopurs_runtime.Value
var once_proxy1 sync.Once
func Get_proxy1() gopurs_runtime.Value {
	once_proxy1.Do(func() {
		proxy1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NatProxy")})
	})
	return proxy1
}

var addNatZero gopurs_runtime.Value
var once_addNatZero sync.Once
func Get_addNatZero() gopurs_runtime.Value {
	once_addNatZero.Do(func() {
		addNatZero = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return addNatZero
}

var addNatSucc gopurs_runtime.Value
var once_addNatSucc sync.Once
func Get_addNatSucc() gopurs_runtime.Value {
	once_addNatSucc.Do(func() {
		addNatSucc = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
	})
	return addNatSucc
}

var addNat gopurs_runtime.Value
var once_addNat sync.Once
func Get_addNat() gopurs_runtime.Value {
	once_addNat.Do(func() {
		addNat = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NatProxy")})
})
})
})
	})
	return addNat
}


