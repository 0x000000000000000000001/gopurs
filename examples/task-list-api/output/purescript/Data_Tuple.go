package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Tuple_go__append gopurs_runtime.Value
var once_Data_Tuple_go__append sync.Once

func Get_Data_Tuple_go__append() gopurs_runtime.Value {
	once_Data_Tuple_go__append.Do(func() {
		cache_Data_Tuple_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Tuple_go__append
}

var cache_Data_Tuple_conj gopurs_runtime.Value
var once_Data_Tuple_conj sync.Once

func Get_Data_Tuple_conj() gopurs_runtime.Value {
	once_Data_Tuple_conj.Do(func() {
		cache_Data_Tuple_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Data_Tuple_conj
}

var cache_Data_Tuple_Tuple gopurs_runtime.Value
var once_Data_Tuple_Tuple sync.Once

func Get_Data_Tuple_Tuple() gopurs_runtime.Value {
	once_Data_Tuple_Tuple.Do(func() {
		cache_Data_Tuple_Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Data_Tuple_Tuple
}

var cache_Data_Tuple_Tuple__2229328841 gopurs_runtime.Value
var once_Data_Tuple_Tuple__2229328841 sync.Once

func Get_Data_Tuple_Tuple__2229328841() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__2229328841.Do(func() {
		cache_Data_Tuple_Tuple__2229328841 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_Tuple__2229328841(Rebox_Data_Tuple_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box.IntVal)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_Tuple__2229328841
}

var cache_Data_Tuple_Tuple__2441629065 gopurs_runtime.Value
var once_Data_Tuple_Tuple__2441629065 sync.Once

func Get_Data_Tuple_Tuple__2441629065() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__2441629065.Do(func() {
		cache_Data_Tuple_Tuple__2441629065 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_Tuple__2441629065(__eta_norm_1_0_box.IntVal, Rebox_Data_Tuple_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_Tuple__2441629065
}

var cache_Data_Tuple_Tuple__137183657 gopurs_runtime.Value
var once_Data_Tuple_Tuple__137183657 sync.Once

func Get_Data_Tuple_Tuple__137183657() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__137183657.Do(func() {
		cache_Data_Tuple_Tuple__137183657 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_Tuple__137183657(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_Tuple__137183657
}

var cache_Data_Tuple_Tuple__1777676937 gopurs_runtime.Value
var once_Data_Tuple_Tuple__1777676937 sync.Once

func Get_Data_Tuple_Tuple__1777676937() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__1777676937.Do(func() {
		cache_Data_Tuple_Tuple__1777676937 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_Tuple__1777676937(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.StrVal())
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_Tuple__1777676937
}

var cache_Data_Tuple_Tuple__2248427721 gopurs_runtime.Value
var once_Data_Tuple_Tuple__2248427721 sync.Once

func Get_Data_Tuple_Tuple__2248427721() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__2248427721.Do(func() {
		cache_Data_Tuple_Tuple__2248427721 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_Tuple__2248427721(__eta_norm_1_0_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](__eta_norm_0_1_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_Tuple__2248427721
}

var cache_Data_Tuple_Tuple__2018102267 gopurs_runtime.Value
var once_Data_Tuple_Tuple__2018102267 sync.Once

func Get_Data_Tuple_Tuple__2018102267() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__2018102267.Do(func() {
		cache_Data_Tuple_Tuple__2018102267 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_Tuple__2018102267(__eta_norm_0_0_box.StrVal())
		})
	})
	return cache_Data_Tuple_Tuple__2018102267
}

var cache_Data_Tuple_Tuple__3062481225 gopurs_runtime.Value
var once_Data_Tuple_Tuple__3062481225 sync.Once

func Get_Data_Tuple_Tuple__3062481225() gopurs_runtime.Value {
	once_Data_Tuple_Tuple__3062481225.Do(func() {
		cache_Data_Tuple_Tuple__3062481225 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_Tuple__3062481225(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_Tuple__3062481225
}

var cache_Data_Tuple_uncurry gopurs_runtime.Value
var once_Data_Tuple_uncurry sync.Once

func Get_Data_Tuple_uncurry() gopurs_runtime.Value {
	once_Data_Tuple_uncurry.Do(func() {
		cache_Data_Tuple_uncurry = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_uncurry(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
		})
	})
	return cache_Data_Tuple_uncurry
}

var cache_Data_Tuple_swap gopurs_runtime.Value
var once_Data_Tuple_swap sync.Once

func Get_Data_Tuple_swap() gopurs_runtime.Value {
	once_Data_Tuple_swap.Do(func() {
		cache_Data_Tuple_swap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Tuple_swap(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Data_Tuple_swap
}

var cache_Data_Tuple_snd gopurs_runtime.Value
var once_Data_Tuple_snd sync.Once

func Get_Data_Tuple_snd() gopurs_runtime.Value {
	once_Data_Tuple_snd.Do(func() {
		cache_Data_Tuple_snd = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_snd(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Tuple_snd
}

var cache_Data_Tuple_showTuple gopurs_runtime.Value
var once_Data_Tuple_showTuple sync.Once

func Get_Data_Tuple_showTuple() gopurs_runtime.Value {
	once_Data_Tuple_showTuple.Do(func() {
		cache_Data_Tuple_showTuple = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_showTuple(dictShow_0_box, dictShow1_1_box)
		})
	})
	return cache_Data_Tuple_showTuple
}

var cache_Data_Tuple_semiringTuple gopurs_runtime.Value
var once_Data_Tuple_semiringTuple sync.Once

func Get_Data_Tuple_semiringTuple() gopurs_runtime.Value {
	once_Data_Tuple_semiringTuple.Do(func() {
		cache_Data_Tuple_semiringTuple = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_semiringTuple(dictSemiring_0_box)
		})
	})
	return cache_Data_Tuple_semiringTuple
}

var cache_Data_Tuple_semigroupoidTuple gopurs_runtime.Value
var once_Data_Tuple_semigroupoidTuple sync.Once

func Get_Data_Tuple_semigroupoidTuple() gopurs_runtime.Value {
	once_Data_Tuple_semigroupoidTuple.Do(func() {
		cache_Data_Tuple_semigroupoidTuple = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_866528276_2173614649((&Constructor_Control_Semigroupoid_Semigroupoid[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
	return cache_Data_Tuple_semigroupoidTuple
}

var cache_Data_Tuple_semigroupTuple gopurs_runtime.Value
var once_Data_Tuple_semigroupTuple sync.Once

func Get_Data_Tuple_semigroupTuple() gopurs_runtime.Value {
	once_Data_Tuple_semigroupTuple.Do(func() {
		cache_Data_Tuple_semigroupTuple = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictSemigroup1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_semigroupTuple(dictSemigroup_0_box, dictSemigroup1_1_box)
		})
	})
	return cache_Data_Tuple_semigroupTuple
}

var cache_Data_Tuple_ringTuple gopurs_runtime.Value
var once_Data_Tuple_ringTuple sync.Once

func Get_Data_Tuple_ringTuple() gopurs_runtime.Value {
	once_Data_Tuple_ringTuple.Do(func() {
		cache_Data_Tuple_ringTuple = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_ringTuple(dictRing_0_box)
		})
	})
	return cache_Data_Tuple_ringTuple
}

var cache_Data_Tuple_monoidTuple gopurs_runtime.Value
var once_Data_Tuple_monoidTuple sync.Once

func Get_Data_Tuple_monoidTuple() gopurs_runtime.Value {
	once_Data_Tuple_monoidTuple.Do(func() {
		cache_Data_Tuple_monoidTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_monoidTuple(dictMonoid_0_box)
		})
	})
	return cache_Data_Tuple_monoidTuple
}

var cache_Data_Tuple_heytingAlgebraTuple gopurs_runtime.Value
var once_Data_Tuple_heytingAlgebraTuple sync.Once

func Get_Data_Tuple_heytingAlgebraTuple() gopurs_runtime.Value {
	once_Data_Tuple_heytingAlgebraTuple.Do(func() {
		cache_Data_Tuple_heytingAlgebraTuple = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_heytingAlgebraTuple(dictHeytingAlgebra_0_box)
		})
	})
	return cache_Data_Tuple_heytingAlgebraTuple
}

var cache_Data_Tuple_genericTuple gopurs_runtime.Value
var once_Data_Tuple_genericTuple sync.Once

func Get_Data_Tuple_genericTuple() gopurs_runtime.Value {
	once_Data_Tuple_genericTuple.Do(func() {
		cache_Data_Tuple_genericTuple = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_4041189842_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V1}))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{(*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
	return cache_Data_Tuple_genericTuple
}

var cache_Data_Tuple_functorTuple gopurs_runtime.Value
var once_Data_Tuple_functorTuple sync.Once

func Get_Data_Tuple_functorTuple() gopurs_runtime.Value {
	once_Data_Tuple_functorTuple.Do(func() {
		cache_Data_Tuple_functorTuple = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2363162019_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
	return cache_Data_Tuple_functorTuple
}

var cache_Data_Tuple_invariantTuple gopurs_runtime.Value
var once_Data_Tuple_invariantTuple sync.Once

func Get_Data_Tuple_invariantTuple() gopurs_runtime.Value {
	once_Data_Tuple_invariantTuple.Do(func() {
		cache_Data_Tuple_invariantTuple = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2883191915_2241343270((&Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), Get_Data_Tuple_functorTuple())})))}
	})
	return cache_Data_Tuple_invariantTuple
}

var cache_Data_Tuple_fst gopurs_runtime.Value
var once_Data_Tuple_fst sync.Once

func Get_Data_Tuple_fst() gopurs_runtime.Value {
	once_Data_Tuple_fst.Do(func() {
		cache_Data_Tuple_fst = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_fst(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
		})
	})
	return cache_Data_Tuple_fst
}

var cache_Data_Tuple_lazyTuple gopurs_runtime.Value
var once_Data_Tuple_lazyTuple sync.Once

func Get_Data_Tuple_lazyTuple() gopurs_runtime.Value {
	once_Data_Tuple_lazyTuple.Do(func() {
		cache_Data_Tuple_lazyTuple = gopurs_runtime.Func(func(dictLazy_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_lazyTuple(dictLazy_0_box)
		})
	})
	return cache_Data_Tuple_lazyTuple
}

var cache_Data_Tuple_extendTuple gopurs_runtime.Value
var once_Data_Tuple_extendTuple sync.Once

func Get_Data_Tuple_extendTuple() gopurs_runtime.Value {
	once_Data_Tuple_extendTuple.Do(func() {
		cache_Data_Tuple_extendTuple = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3108778612_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Tuple_functorTuple()
		}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))})}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
	return cache_Data_Tuple_extendTuple
}

var cache_Data_Tuple_eqTuple gopurs_runtime.Value
var once_Data_Tuple_eqTuple sync.Once

func Get_Data_Tuple_eqTuple() gopurs_runtime.Value {
	once_Data_Tuple_eqTuple.Do(func() {
		cache_Data_Tuple_eqTuple = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_eqTuple(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Data_Tuple_eqTuple
}

var cache_Data_Tuple_ordTuple gopurs_runtime.Value
var once_Data_Tuple_ordTuple sync.Once

func Get_Data_Tuple_ordTuple() gopurs_runtime.Value {
	once_Data_Tuple_ordTuple.Do(func() {
		cache_Data_Tuple_ordTuple = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_ordTuple(dictOrd_0_box)
		})
	})
	return cache_Data_Tuple_ordTuple
}

var cache_Data_Tuple_eq1Tuple gopurs_runtime.Value
var once_Data_Tuple_eq1Tuple sync.Once

func Get_Data_Tuple_eq1Tuple() gopurs_runtime.Value {
	once_Data_Tuple_eq1Tuple.Do(func() {
		cache_Data_Tuple_eq1Tuple = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_eq1Tuple(dictEq_0_box)
		})
	})
	return cache_Data_Tuple_eq1Tuple
}

var cache_Data_Tuple_ord1Tuple gopurs_runtime.Value
var once_Data_Tuple_ord1Tuple sync.Once

func Get_Data_Tuple_ord1Tuple() gopurs_runtime.Value {
	once_Data_Tuple_ord1Tuple.Do(func() {
		cache_Data_Tuple_ord1Tuple = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_ord1Tuple(dictOrd_0_box)
		})
	})
	return cache_Data_Tuple_ord1Tuple
}

var cache_Data_Tuple_curry gopurs_runtime.Value
var once_Data_Tuple_curry sync.Once

func Get_Data_Tuple_curry() gopurs_runtime.Value {
	once_Data_Tuple_curry.Do(func() {
		cache_Data_Tuple_curry = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_curry(f_0_box, a_1_box, b_2_box)
		})
	})
	return cache_Data_Tuple_curry
}

var cache_Data_Tuple_comonadTuple gopurs_runtime.Value
var once_Data_Tuple_comonadTuple sync.Once

func Get_Data_Tuple_comonadTuple() gopurs_runtime.Value {
	once_Data_Tuple_comonadTuple.Do(func() {
		cache_Data_Tuple_comonadTuple = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3056445460_2550391993((&Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Tuple_extendTuple()
		}), Get_Data_Tuple_snd()})))}
	})
	return cache_Data_Tuple_comonadTuple
}

var cache_Data_Tuple_commutativeRingTuple gopurs_runtime.Value
var once_Data_Tuple_commutativeRingTuple sync.Once

func Get_Data_Tuple_commutativeRingTuple() gopurs_runtime.Value {
	once_Data_Tuple_commutativeRingTuple.Do(func() {
		cache_Data_Tuple_commutativeRingTuple = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_commutativeRingTuple(dictCommutativeRing_0_box)
		})
	})
	return cache_Data_Tuple_commutativeRingTuple
}

var cache_Data_Tuple_boundedTuple gopurs_runtime.Value
var once_Data_Tuple_boundedTuple sync.Once

func Get_Data_Tuple_boundedTuple() gopurs_runtime.Value {
	once_Data_Tuple_boundedTuple.Do(func() {
		cache_Data_Tuple_boundedTuple = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_boundedTuple(dictBounded_0_box)
		})
	})
	return cache_Data_Tuple_boundedTuple
}

var cache_Data_Tuple_booleanAlgebraTuple gopurs_runtime.Value
var once_Data_Tuple_booleanAlgebraTuple sync.Once

func Get_Data_Tuple_booleanAlgebraTuple() gopurs_runtime.Value {
	once_Data_Tuple_booleanAlgebraTuple.Do(func() {
		cache_Data_Tuple_booleanAlgebraTuple = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_booleanAlgebraTuple(dictBooleanAlgebra_0_box)
		})
	})
	return cache_Data_Tuple_booleanAlgebraTuple
}

var cache_Data_Tuple_applyTuple gopurs_runtime.Value
var once_Data_Tuple_applyTuple sync.Once

func Get_Data_Tuple_applyTuple() gopurs_runtime.Value {
	once_Data_Tuple_applyTuple.Do(func() {
		cache_Data_Tuple_applyTuple = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_applyTuple(dictSemigroup_0_box)
		})
	})
	return cache_Data_Tuple_applyTuple
}

var cache_Data_Tuple_bindTuple gopurs_runtime.Value
var once_Data_Tuple_bindTuple sync.Once

func Get_Data_Tuple_bindTuple() gopurs_runtime.Value {
	once_Data_Tuple_bindTuple.Do(func() {
		cache_Data_Tuple_bindTuple = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_bindTuple(dictSemigroup_0_box)
		})
	})
	return cache_Data_Tuple_bindTuple
}

var cache_Data_Tuple_applicativeTuple gopurs_runtime.Value
var once_Data_Tuple_applicativeTuple sync.Once

func Get_Data_Tuple_applicativeTuple() gopurs_runtime.Value {
	once_Data_Tuple_applicativeTuple.Do(func() {
		cache_Data_Tuple_applicativeTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_applicativeTuple(dictMonoid_0_box)
		})
	})
	return cache_Data_Tuple_applicativeTuple
}

var cache_Data_Tuple_monadTuple gopurs_runtime.Value
var once_Data_Tuple_monadTuple sync.Once

func Get_Data_Tuple_monadTuple() gopurs_runtime.Value {
	once_Data_Tuple_monadTuple.Do(func() {
		cache_Data_Tuple_monadTuple = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Tuple_monadTuple(dictMonoid_0_box)
		})
	})
	return cache_Data_Tuple_monadTuple
}

type Constructor_Data_Tuple_Tuple[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}

func Call_Data_Tuple_Tuple__2229328841(__eta_norm_1_0_loop *Constructor_Data_Tuple_Tuple[uint32, float64], __eta_norm_0_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
Tuple__2229328841:
	for {
		if false {
			continue Tuple__2229328841
		}
		var __eta_norm_1_0 *Constructor_Data_Tuple_Tuple[uint32, float64] = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1923280011_138441832(Rebox_Data_Tuple_138441832_1923280011(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3132786365_138441832(__eta_norm_1_0))}, gopurs_runtime.Int(__eta_norm_0_1)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
	}
}

func Call_Data_Tuple_Tuple__2441629065(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
Tuple__2441629065:
	for {
		if false {
			continue Tuple__2441629065
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2874766729_138441832(Rebox_Data_Tuple_138441832_2874766729(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Int(__eta_norm_1_0), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1170268447_3094389156(__eta_norm_0_1))}}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
	}
}

func Call_Data_Tuple_Tuple__137183657(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
Tuple__137183657:
	for {
		if false {
			continue Tuple__137183657
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3363075976_138441832(Rebox_Data_Tuple_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Int(__eta_norm_1_0), gopurs_runtime.Int(__eta_norm_0_1)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
	}
}

func Call_Data_Tuple_Tuple__1777676937(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
Tuple__1777676937:
	for {
		if false {
			continue Tuple__1777676937
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2169642284_138441832(Rebox_Data_Tuple_138441832_2169642284(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Int(__eta_norm_1_0), gopurs_runtime.Str(__eta_norm_0_1)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
	}
}

func Call_Data_Tuple_Tuple__2248427721(__eta_norm_1_0_loop float64, __eta_norm_0_1_loop *Constructor_Data_Time_Time) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
Tuple__2248427721:
	for {
		if false {
			continue Tuple__2248427721
		}
		var __eta_norm_1_0 float64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Data_Time_Time = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_4002750455_138441832(Rebox_Data_Tuple_138441832_4002750455(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Float(__eta_norm_1_0), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(__eta_norm_0_1)}}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
	}
}

func Call_Data_Tuple_Tuple__2018102267(__eta_norm_0_0_loop string) gopurs_runtime.Value {
Tuple__2018102267:
	for {
		if false {
			continue Tuple__2018102267
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3103076375_138441832((&Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{1, value0.StrVal(), value1})))}
			})
		}), gopurs_runtime.Str(__eta_norm_0_0))
	}
}

func Call_Data_Tuple_Tuple__3062481225(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
Tuple__3062481225:
	for {
		if false {
			continue Tuple__3062481225
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3103076375_138441832(Rebox_Data_Tuple_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Str(__eta_norm_1_0), __eta_norm_0_1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()))))}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
	}
}

func Call_Data_Tuple_uncurry(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(f_0, (v_1).V0, (v_1).V1)
}

func Call_Data_Tuple_swap(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{(v_0).V1, (v_0).V0}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Data_Tuple_snd(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0).V1
}

func Call_Data_Tuple_showTuple(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
	_ = dictShow1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1829442467_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Tuple ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), " "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), ")"))
	})})))}
}

func Call_Data_Tuple_semiringTuple(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
	_ = dictSemiring_0
	// TAST (Let): one_1_0 shape=Other bindingType=(TypeVar a$scope23)
	one_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "one")
	_ = one_1_0
	// TAST (Let): zero_2_1 shape=Other bindingType=(TypeVar a$scope23)
	zero_2_1 := gopurs_runtime.RecordGet(dictSemiring_0, "zero")
	_ = zero_2_1
	return gopurs_runtime.Func(func(dictSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2944453955_2826095630((&Constructor_Data_Semiring_Semiring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_3, "add"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring1_3, "mul"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{one_1_0, gopurs_runtime.RecordGet(dictSemiring1_3, "one")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{zero_2_1, gopurs_runtime.RecordGet(dictSemiring1_3, "zero")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}())})))}
	})
}

func Call_Data_Tuple_semigroupTuple(dictSemigroup_0_loop gopurs_runtime.Value, dictSemigroup1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	var dictSemigroup1_1 gopurs_runtime.Value = dictSemigroup1_1_loop
	_ = dictSemigroup1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3430395043_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup1_1, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})})))}
}

func Call_Data_Tuple_ringTuple(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
	_ = dictRing_0
	// TAST (Let): semiringTuple1_1_0 shape=App(Var) bindingType=Any
	semiringTuple1_1_0 := Call_Data_Tuple_semiringTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{}))
	_ = semiringTuple1_1_0
	return gopurs_runtime.Func(func(dictRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semiringTuple2_3_1 shape=App(Other) bindingType=(ADT ["Data","Semiring","Semiring"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope53), (TypeVar b$scope54)])])
		semiringTuple2_3_1 := Rebox_Data_Tuple_2826095630_2944453955(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(semiringTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing1_2, "Semiring0"), gopurs_runtime.Value{}))))
		_ = semiringTuple2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3935494019_3060961102((&Constructor_Data_Ring_Ring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2944453955_2826095630(semiringTuple2_3_1))}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing1_2, "sub"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
}

func Call_Data_Tuple_monoidTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar a$scope70)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	// TAST (Let): semigroupTuple1_2_1 shape=App(Var) bindingType=Any
	semigroupTuple1_2_1 := gopurs_runtime.Apply(Get_Data_Tuple_semigroupTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
	_ = semigroupTuple1_2_1
	return gopurs_runtime.Func(func(dictMonoid1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupTuple2_4_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope70), (TypeVar b$scope71)])])
		semigroupTuple2_4_2 := Rebox_Data_Tuple_4179793454_3430395043(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupTuple1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_3, "Semigroup0"), gopurs_runtime.Value{}))))
		_ = semigroupTuple2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3350710019_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3430395043_4179793454(semigroupTuple2_4_2))}
		}), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{mempty_1_0, gopurs_runtime.RecordGet(dictMonoid1_3, "mempty")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}())})))}
	})
}

func Call_Data_Tuple_heytingAlgebraTuple(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
	_ = dictHeytingAlgebra_0
	// TAST (Let): tt_1_0 shape=Other bindingType=(TypeVar a$scope87)
	tt_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
	_ = tt_1_0
	// TAST (Let): ff_2_1 shape=Other bindingType=(TypeVar a$scope87)
	ff_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
	_ = ff_2_1
	return gopurs_runtime.Func(func(dictHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_248497731_2663347022((&Constructor_Data_HeytingAlgebra_HeytingAlgebra[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "conj"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "disj"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{ff_2_1, gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "ff")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "implies"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "not"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		}), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{tt_1_0, gopurs_runtime.RecordGet(dictHeytingAlgebra1_3, "tt")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}())})))}
	})
}

func Call_Data_Tuple_fst(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	return (v_0).V0
}

func Call_Data_Tuple_lazyTuple(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
	_ = dictLazy_0
	// TAST (Let): go__defer_1_0 shape=App(Var) bindingType=(Func [(Func [Unit] (TypeVar a$scope131))] (TypeVar a$scope131))
	go__defer_1_0 := Call_Control_Lazy_go__defer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0))
	_ = go__defer_1_0
	return gopurs_runtime.Func(func(dictLazy1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): defer1_3_1 shape=App(Var) bindingType=(Func [(Func [Unit] (TypeVar b$scope132))] (TypeVar b$scope132))
		defer1_3_1 := Call_Control_Lazy_go__defer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy1_2))
		_ = defer1_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1697203252_690570009((&Constructor_Control_Lazy_Lazy[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Apply(go__defer_1_0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_4, Get_Data_Unit_unit()).UnsafePtr).V0
				})), gopurs_runtime.Apply(defer1_3_1, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_4, Get_Data_Unit_unit()).UnsafePtr).V1
				}))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})})))}
	})
}

func Call_Data_Tuple_eqTuple(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2450374787_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
	})})))}
}

func Call_Data_Tuple_ordTuple(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): eqTuple1_1_0 shape=App(Var) bindingType=Any
	eqTuple1_1_0 := gopurs_runtime.Apply(Get_Data_Tuple_eqTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
	_ = eqTuple1_1_0
	return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): eqTuple2_3_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope60), (TypeVar b$scope61)])])
		eqTuple2_3_1 := Rebox_Data_Tuple_3790796878_2450374787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))))
		_ = eqTuple2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1535415139_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2450374787_3790796878(eqTuple2_3_1))}
		}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
			v_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
			_ = v_6_2
			var __t3 gopurs_runtime.Value
			{
				if v_6_2 == 1527465420 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				if v_6_2 == 380165415 {
					__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1)
			}
		end_branch_3:
			return __t3
		})})))}
	})
}

func Call_Data_Tuple_eq1Tuple(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2909557234_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Tuple_eqTuple(dictEq_0, dictEq1_1)))
	})})))}
}

func Call_Data_Tuple_ord1Tuple(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): ordTuple1_1_0 shape=App(Var) bindingType=Any
	ordTuple1_1_0 := Call_Data_Tuple_ordTuple(dictOrd_0)
	_ = ordTuple1_1_0
	// TAST (Let): eq1Tuple1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq1"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope64)])])
	eq1Tuple1_2_1 := Rebox_Data_Tuple_1766074591_2909557234(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_Tuple_eq1Tuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))))
	_ = eq1Tuple1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1697277074_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2909557234_1766074591(eq1Tuple1_2_1))}
	}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordTuple1_1_0, dictOrd1_3)))
	})})))}
}

func Call_Data_Tuple_curry(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{a_1, b_2}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))})
}

func Call_Data_Tuple_commutativeRingTuple(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
	_ = dictCommutativeRing_0
	// TAST (Let): ringTuple1_1_0 shape=App(Var) bindingType=Any
	ringTuple1_1_0 := Call_Data_Tuple_ringTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
	_ = ringTuple1_1_0
	return gopurs_runtime.Func(func(dictCommutativeRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): ringTuple2_3_1 shape=App(Other) bindingType=(ADT ["Data","Ring","Ring"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope170), (TypeVar b$scope171)])])
		ringTuple2_3_1 := Rebox_Data_Tuple_3060961102_3935494019(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](gopurs_runtime.Apply(ringTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing1_2, "Ring0"), gopurs_runtime.Value{}))))
		_ = ringTuple2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_569417955_1073849710((&Constructor_Data_CommutativeRing_CommutativeRing[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3935494019_3060961102(ringTuple2_3_1))}
		})})))}
	})
}

func Call_Data_Tuple_boundedTuple(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	// TAST (Let): top_1_0 shape=Other bindingType=(TypeVar a$scope175)
	top_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "top")
	_ = top_1_0
	// TAST (Let): bottom_2_1 shape=Other bindingType=(TypeVar a$scope175)
	bottom_2_1 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
	_ = bottom_2_1
	// TAST (Let): ordTuple1_3_2 shape=App(Var) bindingType=Any
	ordTuple1_3_2 := Call_Data_Tuple_ordTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
	_ = ordTuple1_3_2
	return gopurs_runtime.Func(func(dictBounded1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): ordTuple2_5_3 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope175), (TypeVar b$scope176)])])
		ordTuple2_5_3 := Rebox_Data_Tuple_4177771502_1535415139(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordTuple1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_4, "Ord0"), gopurs_runtime.Value{}))))
		_ = ordTuple2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_6519395_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1535415139_4177771502(ordTuple2_5_3))}
		}), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{bottom_2_1, gopurs_runtime.RecordGet(dictBounded1_4, "bottom")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{top_1_0, gopurs_runtime.RecordGet(dictBounded1_4, "top")}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}())})))}
	})
}

func Call_Data_Tuple_booleanAlgebraTuple(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
	_ = dictBooleanAlgebra_0
	// TAST (Let): heytingAlgebraTuple1_1_0 shape=App(Var) bindingType=Any
	heytingAlgebraTuple1_1_0 := Call_Data_Tuple_heytingAlgebraTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
	_ = heytingAlgebraTuple1_1_0
	return gopurs_runtime.Func(func(dictBooleanAlgebra1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): heytingAlgebraTuple2_3_1 shape=App(Other) bindingType=(ADT ["Data","HeytingAlgebra","HeytingAlgebra"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope184), (TypeVar b$scope185)])])
		heytingAlgebraTuple2_3_1 := Rebox_Data_Tuple_2663347022_248497731(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](gopurs_runtime.Apply(heytingAlgebraTuple1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra1_2, "HeytingAlgebra0"), gopurs_runtime.Value{}))))
		_ = heytingAlgebraTuple2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1224846915_72969294((&Constructor_Data_BooleanAlgebra_BooleanAlgebra[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_248497731_2663347022(heytingAlgebraTuple2_3_1))}
		})})))}
	})
}

func Call_Data_Tuple_applyTuple(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2708463828_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Tuple_functorTuple()
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0), gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})})))}
}

func Call_Data_Tuple_bindTuple(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	// TAST (Let): applyTuple1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope188)])])
	applyTuple1_1_0 := Rebox_Data_Tuple_3741347833_2708463828(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Tuple_applyTuple(dictSemigroup_0)))
	_ = applyTuple1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2495815764_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2708463828_3741347833(applyTuple1_1_0))}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope188), (TypeVar b$scope193)])
		v1_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
		_ = v1_4_1
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (v1_4_1).V0), (v1_4_1).V1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})})))}
}

func Call_Data_Tuple_applicativeTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): applyTuple1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope206)])])
	applyTuple1_1_0 := Rebox_Data_Tuple_3741347833_2708463828(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Tuple_applyTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))))
	_ = applyTuple1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_894579924_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2708463828_3741347833(applyTuple1_1_0))}
	}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))})))}
}

func Call_Data_Tuple_monadTuple(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): applicativeTuple1_1_0 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope76)])])
	applicativeTuple1_1_0 := Rebox_Data_Tuple_1439734649_894579924(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Data_Tuple_applicativeTuple(dictMonoid_0)))
	_ = applicativeTuple1_1_0
	// TAST (Let): bindTuple1_2_1 shape=App(Var) bindingType=(ADT ["Control","Bind","Bind"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope76)])])
	bindTuple1_2_1 := Rebox_Data_Tuple_2748095225_2495815764(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Data_Tuple_bindTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))))
	_ = bindTuple1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2157788756_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_894579924_1439734649(applicativeTuple1_1_0))}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_2495815764_2748095225(bindTuple1_2_1))}
	})})))}
}

func Rebox_Data_Tuple_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Tuple_1224846915_72969294(in *Constructor_Data_BooleanAlgebra_BooleanAlgebra[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_138441832_1923280011(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], int64]{}
	out.V0 = Rebox_Data_Tuple_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Tuple_138441832_2169642284(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, string]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Data_Tuple_138441832_2874766729(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_Maybe_Just[int64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_Maybe_Just[int64]]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Data_Tuple_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1))
	return out
}

func Rebox_Data_Tuple_138441832_3103076375(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1
	return out
}

func Rebox_Data_Tuple_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Tuple_138441832_3363075976(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Tuple_138441832_4002750455(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[float64, *Constructor_Data_Time_Time] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[float64, *Constructor_Data_Time_Time]{}
	out.V0 = in.V0.FloatVal()
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V1)
	return out
}

func Rebox_Data_Tuple_1439734649_894579924(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_1535415139_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_1697203252_690570009(in *Constructor_Control_Lazy_Lazy[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Lazy_Lazy[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_1697277074_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_1766074591_2909557234(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_1829442467_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_1923280011_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[uint32, float64], int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_3132786365_138441832(in.V0))}
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Tuple_2157788756_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Monad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2169642284_138441832(in *Constructor_Data_Tuple_Tuple[int64, string]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Data_Tuple_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2450374787_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_248497731_2663347022(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V2)}
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V5)}
	return out
}

func Rebox_Data_Tuple_2495815764_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2663347022_248497731(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V2)
	out.V3 = in.V3
	out.V4 = in.V4
	out.V5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V5)
	return out
}

func Rebox_Data_Tuple_2708463828_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2748095225_2495815764(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2826095630_2944453955(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V2)
	out.V3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V3)
	return out
}

func Rebox_Data_Tuple_2874766729_138441832(in *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Tuple_1170268447_3094389156(in.V1))}
	return out
}

func Rebox_Data_Tuple_2883191915_2241343270(in *Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2909557234_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_2944453955_2826095630(in *Constructor_Data_Semiring_Semiring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V2)}
	out.V3 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V3)}
	return out
}

func Rebox_Data_Tuple_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_3060961102_3935494019(in *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) *Constructor_Data_Ring_Ring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Tuple_3103076375_138441832(in *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Data_Tuple_3108778612_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Tuple_3350710019_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Tuple_3363075976_138441832(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Tuple_3430395043_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_3741347833_2708463828(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_3790796878_2450374787(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_3935494019_3060961102(in *Constructor_Data_Ring_Ring[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_4002750455_138441832(in *Constructor_Data_Tuple_Tuple[float64, *Constructor_Data_Time_Time]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Tuple_4041189842_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_4177771502_1535415139(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_4179793454_3430395043(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_569417955_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_6519395_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Tuple_866528276_2173614649(in *Constructor_Control_Semigroupoid_Semigroupoid[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Tuple_894579924_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(unsafe.Pointer(in))
}
