package Effect_Ref

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var new_ gopurs_runtime.Value
var once_new_ sync.Once
func Get_new_() gopurs_runtime.Value {
	once_new_.Do(func() {
		new_ = Get__new()
	})
	return new_
}

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

var modify_ gopurs_runtime.Value
var once_modify_ sync.Once
func Get_modify_() gopurs_runtime.Value {
	once_modify_.Do(func() {
		modify_ = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_applyEffect().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(pkg_Effect.Get_pureE(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_3_0 := gopurs_runtime.Apply(f_0, s_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": s_prime_3_0, "value": s_prime_3_0})
})), s_1))
})
})
	})
	return modify_
}

func Get__new() gopurs_runtime.Value {
	return X_New
}

func Get_modifyImpl() gopurs_runtime.Value {
	return ModifyImpl
}

func Get_newWithSelf() gopurs_runtime.Value {
	return NewWithSelf
}

func Get_read() gopurs_runtime.Value {
	return Read
}

func Get_write() gopurs_runtime.Value {
	return Write
}
