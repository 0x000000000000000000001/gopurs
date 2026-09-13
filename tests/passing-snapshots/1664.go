package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))
	})
	return cache_Main_pure
}

var cache_Main_Identity gopurs_runtime.Value
var once_Main_Identity sync.Once

func Get_Main_Identity() gopurs_runtime.Value {
	once_Main_Identity.Do(func() {
		cache_Main_Identity = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Identity
}

var cache_Main_Identity__3362365398 gopurs_runtime.Value
var once_Main_Identity__3362365398 sync.Once

func Get_Main_Identity__3362365398() gopurs_runtime.Value {
	once_Main_Identity__3362365398.Do(func() {
		cache_Main_Identity__3362365398 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Identity__3362365398(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Main_Identity__3362365398
}

var cache_Main_IdentityEff gopurs_runtime.Value
var once_Main_IdentityEff sync.Once

func Get_Main_IdentityEff() gopurs_runtime.Value {
	once_Main_IdentityEff.Do(func() {
		cache_Main_IdentityEff = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_IdentityEff(x_0_box)
		})
	})
	return cache_Main_IdentityEff
}

var cache_Main_IdentityEff__3171000602 gopurs_runtime.Value
var once_Main_IdentityEff__3171000602 sync.Once

func Get_Main_IdentityEff__3171000602() gopurs_runtime.Value {
	once_Main_IdentityEff__3171000602.Do(func() {
		cache_Main_IdentityEff__3171000602 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_IdentityEff__3171000602(x_0_box)
		})
	})
	return cache_Main_IdentityEff__3171000602
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(v_0_box)
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Identity[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_Identity__3362365398(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Identity__3362365398:
	for {
		if false {
			continue Identity__3362365398
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Get_Data_Unit_unit()
	}
}

func Call_Main_IdentityEff(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_IdentityEff__3171000602(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
IdentityEff__3171000602:
	for {
		if false {
			continue IdentityEff__3171000602
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_test(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		v1_1_0 := gopurs_runtime.Apply(v_0, gopurs_runtime.Value{})
		_ = v1_1_0
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), Get_Data_Unit_unit()), gopurs_runtime.Value{})
	})
}
