package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Class_MonadEffect_dollar_Dict gopurs_runtime.Value
var once_Effect_Class_MonadEffect_dollar_Dict sync.Once

func Get_Effect_Class_MonadEffect_dollar_Dict() gopurs_runtime.Value {
	once_Effect_Class_MonadEffect_dollar_Dict.Do(func() {
		cache_Effect_Class_MonadEffect_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(Call_Effect_Class_MonadEffect_dollar_Dict(func() struct {
				Monad0     gopurs_runtime.Value
				liftEffect gopurs_runtime.Value
				/*gopurs:record-v1:2:004d006f006e006100640030,006c006900660074004500660066006500630074*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0     gopurs_runtime.Value
					liftEffect gopurs_runtime.Value
					/*gopurs:record-v1:2:004d006f006e006100640030,006c006900660074004500660066006500630074*/
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.liftEffect = gopurs_runtime.RecordGet(orig, "liftEffect")
				return clone
			}()))}
		})
	})
	return cache_Effect_Class_MonadEffect_dollar_Dict
}

var cache_Effect_Class_monadEffectEffect gopurs_runtime.Value
var once_Effect_Class_monadEffectEffect sync.Once

func Get_Effect_Class_monadEffectEffect() gopurs_runtime.Value {
	once_Effect_Class_monadEffectEffect.Do(func() {
		cache_Effect_Class_monadEffectEffect = gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Effect_Class_monadEffectEffect
}

var cache_Effect_Class_liftEffect gopurs_runtime.Value
var once_Effect_Class_liftEffect sync.Once

func Get_Effect_Class_liftEffect() gopurs_runtime.Value {
	once_Effect_Class_liftEffect.Do(func() {
		cache_Effect_Class_liftEffect = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Effect_Class_liftEffect
}

var cache_Effect_Class_liftEffect__3747263901 gopurs_runtime.Value
var once_Effect_Class_liftEffect__3747263901 sync.Once

func Get_Effect_Class_liftEffect__3747263901() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__3747263901.Do(func() {
		cache_Effect_Class_liftEffect__3747263901 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Class_liftEffect__3747263901(__eta0_0_box)
		})
	})
	return cache_Effect_Class_liftEffect__3747263901
}

var cache_Effect_Class_liftEffect__3733620925 gopurs_runtime.Value
var once_Effect_Class_liftEffect__3733620925 sync.Once

func Get_Effect_Class_liftEffect__3733620925() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__3733620925.Do(func() {
		cache_Effect_Class_liftEffect__3733620925 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Class_liftEffect__3733620925(__eta0_0_box)
		})
	})
	return cache_Effect_Class_liftEffect__3733620925
}

type Constructor_Effect_Class_MonadEffect[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004500660066006500630074005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004500660066006500630074005f0043006c006100730073005f004d006f006e00610064004500660066006500630074
func init() {
	gopurs_runtime.StructGetters[2217729261] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "liftEffect":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Effect_Class_MonadEffect: " + key)
		}
	}
}

func Call_Effect_Class_MonadEffect_dollar_Dict(x_0_loop struct {
	Monad0     gopurs_runtime.Value
	liftEffect gopurs_runtime.Value
	/*gopurs:record-v1:2:004d006f006e006100640030,006c006900660074004500660066006500630074*/
}) *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value] {
	var x_0 struct {
		Monad0     gopurs_runtime.Value
		liftEffect gopurs_runtime.Value
		/*gopurs:record-v1:2:004d006f006e006100640030,006c006900660074004500660066006500630074*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "liftEffect", orig.Monad0, orig.liftEffect)
	}())
}

func Call_Effect_Class_liftEffect(dict_0_loop *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Effect_Class_liftEffect__3747263901(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
liftEffect__3747263901:
	for {
		if false {
			continue liftEffect__3747263901
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()).V1, __eta0_0)
	}
}

func Call_Effect_Class_liftEffect__3733620925(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
liftEffect__3733620925:
	for {
		if false {
			continue liftEffect__3733620925
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()).V1, __eta0_0)
	}
}
