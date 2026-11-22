package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Coproduct_Inject_compose gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_compose sync.Once

func Get_Data_Functor_Coproduct_Inject_compose() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_compose.Do(func() {
		cache_Data_Functor_Coproduct_Inject_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Coproduct_Inject_compose
}

var cache_Data_Functor_Coproduct_Inject_compose1 gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_compose1 sync.Once

func Get_Data_Functor_Coproduct_Inject_compose1() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_compose1.Do(func() {
		cache_Data_Functor_Coproduct_Inject_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Coproduct_Inject_compose1
}

var cache_Data_Functor_Coproduct_Inject_Inject_dollar_Dict gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_Inject_dollar_Dict sync.Once

func Get_Data_Functor_Coproduct_Inject_Inject_dollar_Dict() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_Inject_dollar_Dict.Do(func() {
		cache_Data_Functor_Coproduct_Inject_Inject_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Coproduct_Inject_Inject_dollar_Dict(func() struct {
				inj gopurs_runtime.Value
				prj gopurs_runtime.Value
				/*gopurs:record-v1:2:0069006e006a,00700072006a*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					inj gopurs_runtime.Value
					prj gopurs_runtime.Value
					/*gopurs:record-v1:2:0069006e006a,00700072006a*/
				}{}
				clone.inj = gopurs_runtime.RecordGet(orig, "inj")
				clone.prj = gopurs_runtime.RecordGet(orig, "prj")
				return clone
			}()))}
		})
	})
	return cache_Data_Functor_Coproduct_Inject_Inject_dollar_Dict
}

var cache_Data_Functor_Coproduct_Inject_prj gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_prj sync.Once

func Get_Data_Functor_Coproduct_Inject_prj() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_prj.Do(func() {
		cache_Data_Functor_Coproduct_Inject_prj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Coproduct_Inject_prj(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Functor_Coproduct_Inject_prj
}

var cache_Data_Functor_Coproduct_Inject_injectReflexive gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_injectReflexive sync.Once

func Get_Data_Functor_Coproduct_Inject_injectReflexive() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_injectReflexive.Do(func() {
		cache_Data_Functor_Coproduct_Inject_injectReflexive = gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), Get_Data_Maybe_Just()}))}
	})
	return cache_Data_Functor_Coproduct_Inject_injectReflexive
}

var cache_Data_Functor_Coproduct_Inject_injectLeft gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_injectLeft sync.Once

func Get_Data_Functor_Coproduct_Inject_injectLeft() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_injectLeft.Do(func() {
		cache_Data_Functor_Coproduct_Inject_injectLeft = gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), Get_Data_Either_Left()), gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), Get_Data_Maybe_Just(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
		}))}))}
	})
	return cache_Data_Functor_Coproduct_Inject_injectLeft
}

var cache_Data_Functor_Coproduct_Inject_inj gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_inj sync.Once

func Get_Data_Functor_Coproduct_Inject_inj() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_inj.Do(func() {
		cache_Data_Functor_Coproduct_Inject_inj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Coproduct_Inject_inj(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Functor_Coproduct_Inject_inj
}

var cache_Data_Functor_Coproduct_Inject_injectRight gopurs_runtime.Value
var once_Data_Functor_Coproduct_Inject_injectRight sync.Once

func Get_Data_Functor_Coproduct_Inject_injectRight() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Inject_injectRight.Do(func() {
		cache_Data_Functor_Coproduct_Inject_injectRight = gopurs_runtime.Func(func(dictInject_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Coproduct_Inject_injectRight(dictInject_0_box)
		})
	})
	return cache_Data_Functor_Coproduct_Inject_injectRight
}

type Constructor_Data_Functor_Coproduct_Inject_Inject[T_f any, T_g any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00460075006e00630074006f0072005f0043006f00700072006f0064007500630074005f0049006e006a006500630074:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00460075006e00630074006f0072005f0043006f00700072006f0064007500630074005f0049006e006a006500630074005f0049006e006a006500630074
func init() {
	gopurs_runtime.StructGetters[3819687178] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "inj":
			return gopurs_runtime.Box(c.V0)
		case "prj":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Functor_Coproduct_Inject_Inject: " + key)
		}
	}
}

func Call_Data_Functor_Coproduct_Inject_Inject_dollar_Dict(x_0_loop struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
	/*gopurs:record-v1:2:0069006e006a,00700072006a*/
}) *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		inj gopurs_runtime.Value
		prj gopurs_runtime.Value
		/*gopurs:record-v1:2:0069006e006a,00700072006a*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("inj", "prj", orig.inj, orig.prj)
	}())
}

func Call_Data_Functor_Coproduct_Inject_prj(dict_0_loop *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Functor_Coproduct_Inject_inj(dict_0_loop *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Functor_Coproduct_Inject_injectRight(dictInject_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictInject_0 gopurs_runtime.Value = dictInject_0_loop
	_ = dictInject_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3819687178, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Coproduct_Coproduct(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Either_Right(), Call_Data_Functor_Coproduct_Inject_inj(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dictInject_0)))), gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
	}), Call_Data_Functor_Coproduct_Inject_prj(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Coproduct_Inject_Inject[gopurs_runtime.Value, gopurs_runtime.Value]](dictInject_0)))}))}
}
