package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Strong_Strong_dollar_Dict gopurs_runtime.Value
var once_Data_Profunctor_Strong_Strong_dollar_Dict sync.Once

func Get_Data_Profunctor_Strong_Strong_dollar_Dict() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_Strong_dollar_Dict.Do(func() {
		cache_Data_Profunctor_Strong_Strong_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1323482783, UnsafePtr: unsafe.Pointer(Call_Data_Profunctor_Strong_Strong_dollar_Dict(func() struct {
				Profunctor0 gopurs_runtime.Value
				first       gopurs_runtime.Value
				second      gopurs_runtime.Value
				/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,00660069007200730074,007300650063006f006e0064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Profunctor0 gopurs_runtime.Value
					first       gopurs_runtime.Value
					second      gopurs_runtime.Value
					/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,00660069007200730074,007300650063006f006e0064*/
				}{}
				clone.Profunctor0 = gopurs_runtime.RecordGet(orig, "Profunctor0")
				clone.first = gopurs_runtime.RecordGet(orig, "first")
				clone.second = gopurs_runtime.RecordGet(orig, "second")
				return clone
			}()))}
		})
	})
	return cache_Data_Profunctor_Strong_Strong_dollar_Dict
}

var cache_Data_Profunctor_Strong_strongFn gopurs_runtime.Value
var once_Data_Profunctor_Strong_strongFn sync.Once

func Get_Data_Profunctor_Strong_strongFn() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_strongFn.Do(func() {
		cache_Data_Profunctor_Strong_strongFn = gopurs_runtime.Value{Type: 9, IntVal: 1323482783, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()))}
		}), gopurs_runtime.Func2(func(a2b_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(a2b_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map")}))}
	})
	return cache_Data_Profunctor_Strong_strongFn
}

var cache_Data_Profunctor_Strong_second gopurs_runtime.Value
var once_Data_Profunctor_Strong_second sync.Once

func Get_Data_Profunctor_Strong_second() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_second.Do(func() {
		cache_Data_Profunctor_Strong_second = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Strong_second(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Strong_second
}

var cache_Data_Profunctor_Strong_first gopurs_runtime.Value
var once_Data_Profunctor_Strong_first sync.Once

func Get_Data_Profunctor_Strong_first() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_first.Do(func() {
		cache_Data_Profunctor_Strong_first = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Strong_first(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Strong_first
}

var cache_Data_Profunctor_Strong_splitStrong gopurs_runtime.Value
var once_Data_Profunctor_Strong_splitStrong sync.Once

func Get_Data_Profunctor_Strong_splitStrong() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_splitStrong.Do(func() {
		cache_Data_Profunctor_Strong_splitStrong = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Strong_splitStrong(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box))
		})
	})
	return cache_Data_Profunctor_Strong_splitStrong
}

var cache_Data_Profunctor_Strong_fanout gopurs_runtime.Value
var once_Data_Profunctor_Strong_fanout sync.Once

func Get_Data_Profunctor_Strong_fanout() gopurs_runtime.Value {
	once_Data_Profunctor_Strong_fanout.Do(func() {
		cache_Data_Profunctor_Strong_fanout = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Strong_fanout(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box))
		})
	})
	return cache_Data_Profunctor_Strong_fanout
}

type Constructor_Data_Profunctor_Strong_Strong[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00500072006f00660075006e00630074006f0072005f005300740072006f006e0067:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00500072006f00660075006e00630074006f0072005f005300740072006f006e0067005f005300740072006f006e0067
func init() {
	gopurs_runtime.StructGetters[1323482783] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Profunctor0":
			return gopurs_runtime.Box(c.V0)
		case "first":
			return gopurs_runtime.Box(c.V1)
		case "second":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Profunctor_Strong_Strong: " + key)
		}
	}
}

func Call_Data_Profunctor_Strong_Strong_dollar_Dict(x_0_loop struct {
	Profunctor0 gopurs_runtime.Value
	first       gopurs_runtime.Value
	second      gopurs_runtime.Value
	/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,00660069007200730074,007300650063006f006e0064*/
}) *Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value] {
	var x_0 struct {
		Profunctor0 gopurs_runtime.Value
		first       gopurs_runtime.Value
		second      gopurs_runtime.Value
		/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,00660069007200730074,007300650063006f006e0064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Profunctor0", "first", "second", orig.Profunctor0, orig.first, orig.second)
	}())
}

func Call_Data_Profunctor_Strong_second(dict_0_loop *Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Profunctor_Strong_first(dict_0_loop *Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Profunctor_Strong_splitStrong(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
	_ = dictSemigroupoid_0
	// TAST (Let): composeFlipped_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar p$scope28) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar c$scope31)]), (ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope30), (TypeVar c$scope31)])]), (TypeApp (TypeVar p$scope28) [(ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope30), (TypeVar c$scope31)]), (ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope30), (TypeVar d$scope32)])])] (TypeApp (TypeVar p$scope28) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar c$scope31)]), (ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope30), (TypeVar d$scope32)])]))
	composeFlipped_1_0 := Call_Control_Semigroupoid_composeFlipped(dictSemigroupoid_0)
	_ = composeFlipped_1_0
	return gopurs_runtime.Func3(func(dictStrong_2 gopurs_runtime.Value, l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(composeFlipped_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_2, "first"), l_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_2, "second"), r_4))
	})
}

func Call_Data_Profunctor_Strong_fanout(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
	_ = dictSemigroupoid_0
	// TAST (Let): splitStrong1__3813661248_1_0 shape=App(Var) bindingType=(ConstrainedType (Func [(TypeApp (TypeVar p$scope38) [(TypeVar a$scope39), (TypeVar b$scope40)]), (TypeApp (TypeVar p$scope38) [(TypeVar a$scope39), (TypeVar c$scope41)])] (TypeApp (TypeVar p$scope38) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope39), (TypeVar a$scope39)]), (ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope40), (TypeVar c$scope41)])])))
	splitStrong1__3813661248_1_0 := Call_Data_Profunctor_Strong_splitStrong(dictSemigroupoid_0)
	_ = splitStrong1__3813661248_1_0
	return gopurs_runtime.Func(func(dictStrong_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_2, "Profunctor0"), gopurs_runtime.Value{})
		_ = __local_var_3_1
		// TAST (Let): splitStrong2_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar p$scope38) [(TypeVar a$scope39), (TypeVar b$scope40)]), (TypeApp (TypeVar p$scope38) [(TypeVar a$scope39), (TypeVar c$scope41)])] (TypeApp (TypeVar p$scope38) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope39), (TypeVar a$scope39)]), (ADT ["Data","Tuple","Tuple"] [(TypeVar b$scope40), (TypeVar c$scope41)])]))
		splitStrong2_4_2 := gopurs_runtime.Apply(splitStrong1__3813661248_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1323482783, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]](dictStrong_2))})
		_ = splitStrong2_4_2
		return gopurs_runtime.Func2(func(l_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "dimap"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{a_7, a_7}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()
			}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_7
			}), gopurs_runtime.Apply2(splitStrong2_4_2, l_5, r_6))
		})
	})
}
