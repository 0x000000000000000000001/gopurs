package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_f2 gopurs_runtime.Value
var once_Main_f2 sync.Once

func Get_Main_f2() gopurs_runtime.Value {
	once_Main_f2.Do(func() {
		cache_Main_f2 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f2(dict_0_box)
		})
	})
	return cache_Main_f2
}

var cache_Main_f1 gopurs_runtime.Value
var once_Main_f1 sync.Once

func Get_Main_f1() gopurs_runtime.Value {
	once_Main_f1.Do(func() {
		cache_Main_f1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f1(dict_0_box)
		})
	})
	return cache_Main_f1
}

var cache_Main_f_prime_ gopurs_runtime.Value
var once_Main_f_prime_ sync.Once

func Get_Main_f_prime_() gopurs_runtime.Value {
	once_Main_f_prime_.Do(func() {
		cache_Main_f_prime_ = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f_prime_(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_f_prime_
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
		})
	})
	return cache_Main_f
}

type Constructor_Main_Monad2[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d006f006e006100640032
func init() {
	gopurs_runtime.StructGetters[518617989] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Monad2[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "f2":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Monad2: " + key)
		}
	}
}

type Constructor_Main_Monad1[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004d006f006e006100640031
func init() {
	gopurs_runtime.StructGetters[4280338790] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Monad1[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "f1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Monad1: " + key)
		}
	}
}

func Call_Main_f2(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "f2")
}

func Call_Main_f1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "f1")
}

func Call_Main_f_prime_(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope3)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
	_ = Bind1_1_0
	// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope3)])
	Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
	_ = Applicative0_2_1
	return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Bind_bind(Bind1_1_0), gopurs_runtime.Apply(Call_Control_Applicative_pure(Applicative0_2_1), gopurs_runtime.Int(n_3.IntVal)), gopurs_runtime.Func(func(n_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Control_Applicative_pure(Applicative0_2_1), gopurs_runtime.Int(n_prime__4.IntVal))
		}))
	})
}

func Call_Main_f(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope5)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
	_ = Bind1_1_0
	// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope5)])
	Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
	_ = Applicative0_2_1
	return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Bind_bind(Bind1_1_0), gopurs_runtime.Apply(Call_Control_Applicative_pure(Applicative0_2_1), gopurs_runtime.Int(n_3.IntVal)), gopurs_runtime.Func(func(n_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Call_Control_Applicative_pure(Applicative0_2_1), gopurs_runtime.Int(n_prime__4.IntVal))
		}))
	})
}
