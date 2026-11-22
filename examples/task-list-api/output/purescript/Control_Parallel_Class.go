package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Parallel_Class_compose gopurs_runtime.Value
var once_Control_Parallel_Class_compose sync.Once

func Get_Control_Parallel_Class_compose() gopurs_runtime.Value {
	once_Control_Parallel_Class_compose.Do(func() {
		cache_Control_Parallel_Class_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_compose
}

var cache_Control_Parallel_Class_compose1 gopurs_runtime.Value
var once_Control_Parallel_Class_compose1 sync.Once

func Get_Control_Parallel_Class_compose1() gopurs_runtime.Value {
	once_Control_Parallel_Class_compose1.Do(func() {
		cache_Control_Parallel_Class_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_compose1
}

var cache_Control_Parallel_Class_composeFlipped gopurs_runtime.Value
var once_Control_Parallel_Class_composeFlipped sync.Once

func Get_Control_Parallel_Class_composeFlipped() gopurs_runtime.Value {
	once_Control_Parallel_Class_composeFlipped.Do(func() {
		cache_Control_Parallel_Class_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_composeFlipped
}

var cache_Control_Parallel_Class_composeFlipped1 gopurs_runtime.Value
var once_Control_Parallel_Class_composeFlipped1 sync.Once

func Get_Control_Parallel_Class_composeFlipped1() gopurs_runtime.Value {
	once_Control_Parallel_Class_composeFlipped1.Do(func() {
		cache_Control_Parallel_Class_composeFlipped1 = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_composeFlipped1
}

var cache_Control_Parallel_Class_compose2 gopurs_runtime.Value
var once_Control_Parallel_Class_compose2 sync.Once

func Get_Control_Parallel_Class_compose2() gopurs_runtime.Value {
	once_Control_Parallel_Class_compose2.Do(func() {
		cache_Control_Parallel_Class_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_compose2
}

var cache_Control_Parallel_Class_compose3 gopurs_runtime.Value
var once_Control_Parallel_Class_compose3 sync.Once

func Get_Control_Parallel_Class_compose3() gopurs_runtime.Value {
	once_Control_Parallel_Class_compose3.Do(func() {
		cache_Control_Parallel_Class_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_compose3
}

var cache_Control_Parallel_Class_discard gopurs_runtime.Value
var once_Control_Parallel_Class_discard sync.Once

func Get_Control_Parallel_Class_discard() gopurs_runtime.Value {
	once_Control_Parallel_Class_discard.Do(func() {
		cache_Control_Parallel_Class_discard = gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard[gopurs_runtime.Value]](Get_Control_Bind_discardUnit()).V0
	})
	return cache_Control_Parallel_Class_discard
}

var cache_Control_Parallel_Class_compose4 gopurs_runtime.Value
var once_Control_Parallel_Class_compose4 sync.Once

func Get_Control_Parallel_Class_compose4() gopurs_runtime.Value {
	once_Control_Parallel_Class_compose4.Do(func() {
		cache_Control_Parallel_Class_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Parallel_Class_compose4
}

var cache_Control_Parallel_Class_discard1 gopurs_runtime.Value
var once_Control_Parallel_Class_discard1 sync.Once

func Get_Control_Parallel_Class_discard1() gopurs_runtime.Value {
	once_Control_Parallel_Class_discard1.Do(func() {
		cache_Control_Parallel_Class_discard1 = gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard[gopurs_runtime.Value]](Get_Control_Bind_discardUnit()).V0
	})
	return cache_Control_Parallel_Class_discard1
}

var cache_Control_Parallel_Class_Parallel_dollar_Dict gopurs_runtime.Value
var once_Control_Parallel_Class_Parallel_dollar_Dict sync.Once

func Get_Control_Parallel_Class_Parallel_dollar_Dict() gopurs_runtime.Value {
	once_Control_Parallel_Class_Parallel_dollar_Dict.Do(func() {
		cache_Control_Parallel_Class_Parallel_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(Call_Control_Parallel_Class_Parallel_dollar_Dict(func() struct {
				Apply0     gopurs_runtime.Value
				Apply1     gopurs_runtime.Value
				parallel   gopurs_runtime.Value
				sequential gopurs_runtime.Value
				/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Apply0     gopurs_runtime.Value
					Apply1     gopurs_runtime.Value
					parallel   gopurs_runtime.Value
					sequential gopurs_runtime.Value
					/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
				}{}
				clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
				clone.Apply1 = gopurs_runtime.RecordGet(orig, "Apply1")
				clone.parallel = gopurs_runtime.RecordGet(orig, "parallel")
				clone.sequential = gopurs_runtime.RecordGet(orig, "sequential")
				return clone
			}()))}
		})
	})
	return cache_Control_Parallel_Class_Parallel_dollar_Dict
}

var cache_Control_Parallel_Class_Parallel_dollar_Dict__189190355 gopurs_runtime.Value
var once_Control_Parallel_Class_Parallel_dollar_Dict__189190355 sync.Once

func Get_Control_Parallel_Class_Parallel_dollar_Dict__189190355() gopurs_runtime.Value {
	once_Control_Parallel_Class_Parallel_dollar_Dict__189190355.Do(func() {
		cache_Control_Parallel_Class_Parallel_dollar_Dict__189190355 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(Call_Control_Parallel_Class_Parallel_dollar_Dict__189190355(func() struct {
				Apply0     gopurs_runtime.Value
				Apply1     gopurs_runtime.Value
				parallel   gopurs_runtime.Value
				sequential gopurs_runtime.Value
				/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Apply0     gopurs_runtime.Value
					Apply1     gopurs_runtime.Value
					parallel   gopurs_runtime.Value
					sequential gopurs_runtime.Value
					/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
				}{}
				clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
				clone.Apply1 = gopurs_runtime.RecordGet(orig, "Apply1")
				clone.parallel = gopurs_runtime.RecordGet(orig, "parallel")
				clone.sequential = gopurs_runtime.RecordGet(orig, "sequential")
				return clone
			}()))}
		})
	})
	return cache_Control_Parallel_Class_Parallel_dollar_Dict__189190355
}

var cache_Control_Parallel_Class_ParCont gopurs_runtime.Value
var once_Control_Parallel_Class_ParCont sync.Once

func Get_Control_Parallel_Class_ParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_ParCont.Do(func() {
		cache_Control_Parallel_Class_ParCont = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_ParCont(x_0_box)
		})
	})
	return cache_Control_Parallel_Class_ParCont
}

var cache_Control_Parallel_Class_sequential gopurs_runtime.Value
var once_Control_Parallel_Class_sequential sync.Once

func Get_Control_Parallel_Class_sequential() gopurs_runtime.Value {
	once_Control_Parallel_Class_sequential.Do(func() {
		cache_Control_Parallel_Class_sequential = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Parallel_Class_sequential
}

var cache_Control_Parallel_Class_parallel gopurs_runtime.Value
var once_Control_Parallel_Class_parallel sync.Once

func Get_Control_Parallel_Class_parallel() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel.Do(func() {
		cache_Control_Parallel_Class_parallel = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Parallel_Class_parallel
}

var cache_Control_Parallel_Class_newtypeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_newtypeParCont sync.Once

func Get_Control_Parallel_Class_newtypeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_newtypeParCont.Do(func() {
		cache_Control_Parallel_Class_newtypeParCont = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Parallel_Class_newtypeParCont
}

var cache_Control_Parallel_Class_monadParWriterT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParWriterT sync.Once

func Get_Control_Parallel_Class_monadParWriterT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParWriterT.Do(func() {
		cache_Control_Parallel_Class_monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParWriterT(dictMonoid_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParWriterT
}

var cache_Control_Parallel_Class_monadParStar gopurs_runtime.Value
var once_Control_Parallel_Class_monadParStar sync.Once

func Get_Control_Parallel_Class_monadParStar() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParStar.Do(func() {
		cache_Control_Parallel_Class_monadParStar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParStar(dictParallel_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParStar
}

var cache_Control_Parallel_Class_monadParReaderT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParReaderT sync.Once

func Get_Control_Parallel_Class_monadParReaderT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParReaderT.Do(func() {
		cache_Control_Parallel_Class_monadParReaderT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParReaderT(dictParallel_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParReaderT
}

var cache_Control_Parallel_Class_monadParMaybeT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParMaybeT sync.Once

func Get_Control_Parallel_Class_monadParMaybeT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParMaybeT.Do(func() {
		cache_Control_Parallel_Class_monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParMaybeT(dictParallel_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParMaybeT
}

var cache_Control_Parallel_Class_monadParExceptT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParExceptT sync.Once

func Get_Control_Parallel_Class_monadParExceptT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParExceptT.Do(func() {
		cache_Control_Parallel_Class_monadParExceptT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParExceptT(dictParallel_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParExceptT
}

var cache_Control_Parallel_Class_monadParCostar gopurs_runtime.Value
var once_Control_Parallel_Class_monadParCostar sync.Once

func Get_Control_Parallel_Class_monadParCostar() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParCostar.Do(func() {
		cache_Control_Parallel_Class_monadParCostar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParCostar(dictParallel_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParCostar
}

var cache_Control_Parallel_Class_monadParParCont gopurs_runtime.Value
var once_Control_Parallel_Class_monadParParCont sync.Once

func Get_Control_Parallel_Class_monadParParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParParCont.Do(func() {
		cache_Control_Parallel_Class_monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_monadParParCont
}

var cache_Control_Parallel_Class_functorParCont gopurs_runtime.Value
var once_Control_Parallel_Class_functorParCont sync.Once

func Get_Control_Parallel_Class_functorParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_functorParCont.Do(func() {
		cache_Control_Parallel_Class_functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_functorParCont
}

var cache_Control_Parallel_Class_applyParCont gopurs_runtime.Value
var once_Control_Parallel_Class_applyParCont sync.Once

func Get_Control_Parallel_Class_applyParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_applyParCont.Do(func() {
		cache_Control_Parallel_Class_applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_applyParCont
}

var cache_Control_Parallel_Class_applicativeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_applicativeParCont sync.Once

func Get_Control_Parallel_Class_applicativeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_applicativeParCont.Do(func() {
		cache_Control_Parallel_Class_applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_applicativeParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_applicativeParCont
}

var cache_Control_Parallel_Class_altParCont gopurs_runtime.Value
var once_Control_Parallel_Class_altParCont sync.Once

func Get_Control_Parallel_Class_altParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_altParCont.Do(func() {
		cache_Control_Parallel_Class_altParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_altParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_altParCont
}

var cache_Control_Parallel_Class_plusParCont gopurs_runtime.Value
var once_Control_Parallel_Class_plusParCont sync.Once

func Get_Control_Parallel_Class_plusParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_plusParCont.Do(func() {
		cache_Control_Parallel_Class_plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_plusParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_plusParCont
}

var cache_Control_Parallel_Class_alternativeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_alternativeParCont sync.Once

func Get_Control_Parallel_Class_alternativeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_alternativeParCont.Do(func() {
		cache_Control_Parallel_Class_alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Parallel_Class_alternativeParCont(dictMonadEffect_0_box)
		})
	})
	return cache_Control_Parallel_Class_alternativeParCont
}

type Constructor_Control_Parallel_Class_Parallel[T_f any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0050006100720061006c006c0065006c005f0043006c006100730073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0050006100720061006c006c0065006c005f0043006c006100730073005f0050006100720061006c006c0065006c
func init() {
	gopurs_runtime.StructGetters[327692956] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Apply0":
			return gopurs_runtime.Box(c.V0)
		case "Apply1":
			return gopurs_runtime.Box(c.V1)
		case "parallel":
			return gopurs_runtime.Box(c.V2)
		case "sequential":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Control_Parallel_Class_Parallel: " + key)
		}
	}
}

func Call_Control_Parallel_Class_Parallel_dollar_Dict(x_0_loop struct {
	Apply0     gopurs_runtime.Value
	Apply1     gopurs_runtime.Value
	parallel   gopurs_runtime.Value
	sequential gopurs_runtime.Value
	/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
}) *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Apply0     gopurs_runtime.Value
		Apply1     gopurs_runtime.Value
		parallel   gopurs_runtime.Value
		sequential gopurs_runtime.Value
		/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", orig.Apply0, orig.Apply1, orig.parallel, orig.sequential)
	}())
}

func Call_Control_Parallel_Class_Parallel_dollar_Dict__189190355(x_0_loop struct {
	Apply0     gopurs_runtime.Value
	Apply1     gopurs_runtime.Value
	parallel   gopurs_runtime.Value
	sequential gopurs_runtime.Value
	/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
}) *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] {
Parallel_dollar_Dict__189190355:
	for {
		if false {
			continue Parallel_dollar_Dict__189190355
		}
		var x_0 struct {
			Apply0     gopurs_runtime.Value
			Apply1     gopurs_runtime.Value
			parallel   gopurs_runtime.Value
			sequential gopurs_runtime.Value
			/*gopurs:record-v1:4:004100700070006c00790030,004100700070006c00790031,0070006100720061006c006c0065006c,00730065007100750065006e007400690061006c*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", orig.Apply0, orig.Apply1, orig.parallel, orig.sequential)
		}())
	}
}

func Call_Control_Parallel_Class_ParCont(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Parallel_Class_sequential(dict_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V3
}

func Call_Control_Parallel_Class_parallel(dict_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Control_Parallel_Class_monadParWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): applyWriterT_1_0 shape=App(Var) bindingType=Any
	applyWriterT_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
	_ = applyWriterT_1_0
	return gopurs_runtime.Func(func(dictParallel_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyWriterT1_3_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m$scope29) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope27)])])])
		applyWriterT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply0"), gopurs_runtime.Value{})))
		_ = applyWriterT1_3_1
		// TAST (Let): applyWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar f$scope28) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w$scope27)])])])
		applyWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply1"), gopurs_runtime.Value{})))
		_ = applyWriterT2_4_2
		// TAST (Let): __local_var_5_3 shape=App(Var) bindingType=Any
		__local_var_5_3 := Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_2))
		_ = __local_var_5_3
		// TAST (Let): __local_var_5_4 shape=App(Var) bindingType=Any
		__local_var_5_4 := Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_2))
		_ = __local_var_5_4
		return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT1_3_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_2)}
		}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_5_3, v_6)
		}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_5_4, v_6)
		})}))}
	})
}

func Call_Control_Parallel_Class_monadParStar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): parallel1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope40) [(TypeVar a$scope46)])] (TypeApp (TypeVar f$scope39) [(TypeVar a$scope46)]))
	parallel1_1_0 := Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0))
	_ = parallel1_1_0
	// TAST (Let): sequential1_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope39) [(TypeVar a$scope52)])] (TypeApp (TypeVar m$scope40) [(TypeVar a$scope52)]))
	sequential1_2_1 := Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0))
	_ = sequential1_2_1
	// TAST (Let): applyStar_3_2 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar m$scope40), (TypeVar a$scope41)])])
	applyStar_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Profunctor_Star_applyStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})))
	_ = applyStar_3_2
	// TAST (Let): applyStar1_4_3 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f$scope39), (TypeVar a$scope41)])])
	applyStar1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Profunctor_Star_applyStar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})))
	_ = applyStar1_4_3
	return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar_3_2)}
	}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_4_3)}
	}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel1_1_0, v_5)
	}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential1_2_1, v_5)
	})}))}
}

func Call_Control_Parallel_Class_monadParReaderT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): applyReaderT_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e$scope63), (TypeVar m$scope62)])])
	applyReaderT_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Reader_Trans_applyReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})))
	_ = applyReaderT_1_0
	// TAST (Let): applyReaderT1_2_1 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e$scope63), (TypeVar f$scope61)])])
	applyReaderT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Reader_Trans_applyReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})))
	_ = applyReaderT1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_2_1)}
	}), gopurs_runtime.Apply(Get_Control_Monad_Reader_Trans_mapReaderT(), Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0))), gopurs_runtime.Apply(Get_Control_Monad_Reader_Trans_mapReaderT(), Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0)))}))}
}

func Call_Control_Parallel_Class_monadParMaybeT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): applyCompose_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope86), (ADT ["Data","Maybe","Maybe"] [])])])
	applyCompose_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Functor_Compose_applyCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Parallel_Class_3552963512_3741347833(Rebox_Control_Parallel_Class_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe()))))}))
	_ = applyCompose_1_0
	return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyMaybeT_3_1 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m$scope87) [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])])
		applyMaybeT_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_2))
		_ = applyMaybeT_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyMaybeT_3_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose_1_0)}
		}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
		}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
		})}))}
	})
}

func Call_Control_Parallel_Class_monadParExceptT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): applyCompose_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope104), (ADT ["Data","Either","Either"] [(TypeVar e$scope106)])])])
	applyCompose_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Functor_Compose_applyCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})), Get_Data_Either_applyEither()))
	_ = applyCompose_1_0
	return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyExceptT_3_1 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m$scope105) [(ADT ["Data","Either","Either"] [(TypeVar e$scope106), (TypeVar a)])])])
		applyExceptT_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_2))
		_ = applyExceptT_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyExceptT_3_1)}
		}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose_1_0)}
		}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
		}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
		})}))}
	})
}

func Call_Control_Parallel_Class_monadParCostar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
	_ = dictParallel_0
	// TAST (Let): sequential1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f$scope124) [(TypeVar a$scope126)])] (TypeApp (TypeVar m$scope125) [(TypeVar a$scope126)]))
	sequential1_1_0 := Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0))
	_ = sequential1_1_0
	// TAST (Let): parallel1_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope125) [(TypeVar a$scope126)])] (TypeApp (TypeVar f$scope124) [(TypeVar a$scope126)]))
	parallel1_2_1 := Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dictParallel_0))
	_ = parallel1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Functor_Costar_applyCostar()
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Functor_Costar_applyCostar()
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), sequential1_1_0, v_3)
	}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), parallel1_2_1, v_3)
	})}))}
}

func Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): applyContT_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m$scope75)])])
	applyContT_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_applyContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{})))
	_ = applyContT_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0)
	}), Get_Control_Parallel_Class_ParCont(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return v_2
	})}))}
}

func Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0))), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Cont_Trans_functorContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})), "map"), f_1), Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0)))))
	})}))}
}

func Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope151)])
	Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_1_0
	// TAST (Let): discard2_2_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope151) [Unit]), (Func [Unit] (TypeApp (TypeVar m$scope151) [Unit]))] (TypeApp (TypeVar m$scope151) [Unit]))
	discard2_2_1 := Call_Control_Bind_bind(Bind1_1_0)
	_ = discard2_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0)
	}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(discard2_2_1, Call_Control_Monad_Cont_Trans_runContT(v_3, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_7)), gopurs_runtime.Func(func(mb_9 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t2 gopurs_runtime.Value
						{
							if mb_9.Type == 9 && mb_9.IntVal == 930809136 && mb_9.UnsafePtr == nil {
								__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{a_8, true}
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()))}, ra_6))
								goto end_branch_2
							} else {

							}
						}
						{
							if mb_9.Type == 9 && mb_9.IntVal == 930809136 && mb_9.UnsafePtr != nil {
								__t2 = gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(a_8, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mb_9.UnsafePtr).V0))
								goto end_branch_2
							} else {

							}
						}
						{
							__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_2:
						return __t2
					}))
				})), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Control_Monad_Cont_Trans_runContT(v1_4, gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_6)), gopurs_runtime.Func(func(ma_10 gopurs_runtime.Value) gopurs_runtime.Value {
							var __t3 gopurs_runtime.Value
							{
								if ma_10.Type == 9 && ma_10.IntVal == 930809136 && ma_10.UnsafePtr == nil {
									__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
										_v := struct {
											V0 gopurs_runtime.Value
											V1 bool
										}{b_9, true}
										if _v.V1 {
											return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
										}
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
									}()))}, rb_7))
									goto end_branch_3
								} else {

								}
							}
							{
								if ma_10.Type == 9 && ma_10.IntVal == 930809136 && ma_10.UnsafePtr != nil {
									__t3 = gopurs_runtime.Apply(k_5, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ma_10.UnsafePtr).V0, b_9))
									goto end_branch_3
								} else {

								}
							}
							{
								__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
							}
						end_branch_3:
							return __t3
						}))
					}))
				}))
			}))
		}))
	})}))}
}

func Call_Control_Parallel_Class_applicativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): applyParCont1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m$scope174), (TypeVar a)])])
	applyParCont1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0))
	_ = applyParCont1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyParCont1_1_0)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Cont_Trans_applicativeContT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{})))))}))}
}

func Call_Control_Parallel_Class_altParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope183)])
	Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_2_1
	// TAST (Let): discard2_3_2 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope183) [Unit]), (Func [Unit] (TypeApp (TypeVar m$scope183) [Unit]))] (TypeApp (TypeVar m$scope183) [Unit]))
	discard2_3_2 := Call_Control_Bind_bind(Bind1_2_1)
	_ = discard2_3_2
	// TAST (Let): __local_var_4_3 shape=App(Other) bindingType=Any
	__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
	_ = __local_var_4_3
	// TAST (Let): functorParCont1_5_4 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m$scope183), (TypeVar a)])])
	functorParCont1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0))
	_ = functorParCont1_5_4
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorParCont1_5_4)}
	}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(done_9 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(discard2_3_2, Call_Control_Monad_Cont_Trans_runContT(v_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_9)), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t5 gopurs_runtime.Value
					{
						if (b_11.IntVal) != (0) {
							__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), Get_Data_Unit_unit())
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = gopurs_runtime.Apply2(discard2_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_9)), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(k_8, a_10)
						}))
					}
				end_branch_5:
					return __t5
				}))
			})), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Control_Monad_Cont_Trans_runContT(v1_7, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_9)), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t6 gopurs_runtime.Value
						{
							if (b_12.IntVal) != (0) {
								__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), Get_Data_Unit_unit())
								goto end_branch_6
							} else {

							}
						}
						{
							__t6 = gopurs_runtime.Apply2(discard2_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_9)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(k_8, a_11)
							}))
						}
					end_branch_6:
						return __t6
					}))
				}))
			}))
		}))
	})}))}
}

func Call_Control_Parallel_Class_plusParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): altParCont1_1_0 shape=App(Var) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m$scope4), (TypeVar a)])])
	altParCont1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Call_Control_Parallel_Class_altParCont(dictMonadEffect_0))
	_ = altParCont1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altParCont1_1_0)}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), Get_Data_Unit_unit())
	})}))}
}

func Call_Control_Parallel_Class_alternativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
	_ = dictMonadEffect_0
	// TAST (Let): applicativeParCont1_1_0 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m$scope180), (TypeVar a)])])
	applicativeParCont1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Parallel_Class_applicativeParCont(dictMonadEffect_0))
	_ = applicativeParCont1_1_0
	// TAST (Let): plusParCont1_2_1 shape=App(Var) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m$scope180), (TypeVar a)])])
	plusParCont1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Call_Control_Parallel_Class_plusParCont(dictMonadEffect_0))
	_ = plusParCont1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeParCont1_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusParCont1_2_1)}
	})}))}
}

func Rebox_Control_Parallel_Class_3552963512_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Parallel_Class_3741347833_3552963512(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}
