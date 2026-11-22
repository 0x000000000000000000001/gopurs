package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_Last_go__append gopurs_runtime.Value
var once_Data_Maybe_Last_go__append sync.Once

func Get_Data_Maybe_Last_go__append() gopurs_runtime.Value {
	once_Data_Maybe_Last_go__append.Do(func() {
		cache_Data_Maybe_Last_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Maybe_Last_go__append
}

var cache_Data_Maybe_Last_Last gopurs_runtime.Value
var once_Data_Maybe_Last_Last sync.Once

func Get_Data_Maybe_Last_Last() gopurs_runtime.Value {
	once_Data_Maybe_Last_Last.Do(func() {
		cache_Data_Maybe_Last_Last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Last_Last(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Maybe_Last_Last
}

var cache_Data_Maybe_Last_showLast gopurs_runtime.Value
var once_Data_Maybe_Last_showLast sync.Once

func Get_Data_Maybe_Last_showLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_showLast.Do(func() {
		cache_Data_Maybe_Last_showLast = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Maybe_Last_showLast(dictShow_0_box)
		})
	})
	return cache_Data_Maybe_Last_showLast
}

var cache_Data_Maybe_Last_semigroupLast gopurs_runtime.Value
var once_Data_Maybe_Last_semigroupLast sync.Once

func Get_Data_Maybe_Last_semigroupLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_semigroupLast.Do(func() {
		cache_Data_Maybe_Last_semigroupLast = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_125234255_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil {
					__t0 = v1_1
					goto end_branch_0
				} else {

				}
			}
			{
				if v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr == nil {
					__t0 = v_0
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_0:
			return __t0
		})})))}
	})
	return cache_Data_Maybe_Last_semigroupLast
}

var cache_Data_Maybe_Last_ordLast gopurs_runtime.Value
var once_Data_Maybe_Last_ordLast sync.Once

func Get_Data_Maybe_Last_ordLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_ordLast.Do(func() {
		cache_Data_Maybe_Last_ordLast = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Maybe_Last_ordLast(dictOrd_0_box)
		})
	})
	return cache_Data_Maybe_Last_ordLast
}

var cache_Data_Maybe_Last_ord1Last gopurs_runtime.Value
var once_Data_Maybe_Last_ord1Last sync.Once

func Get_Data_Maybe_Last_ord1Last() gopurs_runtime.Value {
	once_Data_Maybe_Last_ord1Last.Do(func() {
		cache_Data_Maybe_Last_ord1Last = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_170825214_3985601471(Rebox_Data_Maybe_Last_3985601471_170825214(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_Maybe_ord1Maybe()))))}
	})
	return cache_Data_Maybe_Last_ord1Last
}

var cache_Data_Maybe_Last_newtypeLast gopurs_runtime.Value
var once_Data_Maybe_Last_newtypeLast sync.Once

func Get_Data_Maybe_Last_newtypeLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_newtypeLast.Do(func() {
		cache_Data_Maybe_Last_newtypeLast = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_1892488872_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_Maybe_Last_newtypeLast
}

var cache_Data_Maybe_Last_monoidLast gopurs_runtime.Value
var once_Data_Maybe_Last_monoidLast sync.Once

func Get_Data_Maybe_Last_monoidLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_monoidLast.Do(func() {
		cache_Data_Maybe_Last_monoidLast = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_1089607855_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Maybe_Last_semigroupLast()
		}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())})))}
	})
	return cache_Data_Maybe_Last_monoidLast
}

var cache_Data_Maybe_Last_monadLast gopurs_runtime.Value
var once_Data_Maybe_Last_monadLast sync.Once

func Get_Data_Maybe_Last_monadLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_monadLast.Do(func() {
		cache_Data_Maybe_Last_monadLast = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_1642601656_2568689657(Rebox_Data_Maybe_Last_2568689657_1642601656(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_Maybe_monadMaybe()))))}
	})
	return cache_Data_Maybe_Last_monadLast
}

var cache_Data_Maybe_Last_invariantLast gopurs_runtime.Value
var once_Data_Maybe_Last_invariantLast sync.Once

func Get_Data_Maybe_Last_invariantLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_invariantLast.Do(func() {
		cache_Data_Maybe_Last_invariantLast = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3818752775_2241343270(Rebox_Data_Maybe_Last_2241343270_3818752775(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]](Get_Data_Maybe_invariantMaybe()))))}
	})
	return cache_Data_Maybe_Last_invariantLast
}

var cache_Data_Maybe_Last_functorLast gopurs_runtime.Value
var once_Data_Maybe_Last_functorLast sync.Once

func Get_Data_Maybe_Last_functorLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_functorLast.Do(func() {
		cache_Data_Maybe_Last_functorLast = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3689823567_2812149806(Rebox_Data_Maybe_Last_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
	})
	return cache_Data_Maybe_Last_functorLast
}

var cache_Data_Maybe_Last_extendLast gopurs_runtime.Value
var once_Data_Maybe_Last_extendLast sync.Once

func Get_Data_Maybe_Last_extendLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_extendLast.Do(func() {
		cache_Data_Maybe_Last_extendLast = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_570732504_3290176857(Rebox_Data_Maybe_Last_3290176857_570732504(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Get_Data_Maybe_extendMaybe()))))}
	})
	return cache_Data_Maybe_Last_extendLast
}

var cache_Data_Maybe_Last_eqLast gopurs_runtime.Value
var once_Data_Maybe_Last_eqLast sync.Once

func Get_Data_Maybe_Last_eqLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_eqLast.Do(func() {
		cache_Data_Maybe_Last_eqLast = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Maybe_Last_eqLast(dictEq_0_box)
		})
	})
	return cache_Data_Maybe_Last_eqLast
}

var cache_Data_Maybe_Last_eq1Last gopurs_runtime.Value
var once_Data_Maybe_Last_eq1Last sync.Once

func Get_Data_Maybe_Last_eq1Last() gopurs_runtime.Value {
	once_Data_Maybe_Last_eq1Last.Do(func() {
		cache_Data_Maybe_Last_eq1Last = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_1662389854_1766074591(Rebox_Data_Maybe_Last_1766074591_1662389854(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Maybe_eq1Maybe()))))}
	})
	return cache_Data_Maybe_Last_eq1Last
}

var cache_Data_Maybe_Last_boundedLast gopurs_runtime.Value
var once_Data_Maybe_Last_boundedLast sync.Once

func Get_Data_Maybe_Last_boundedLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_boundedLast.Do(func() {
		cache_Data_Maybe_Last_boundedLast = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Maybe_Last_boundedLast(dictBounded_0_box)
		})
	})
	return cache_Data_Maybe_Last_boundedLast
}

var cache_Data_Maybe_Last_bindLast gopurs_runtime.Value
var once_Data_Maybe_Last_bindLast sync.Once

func Get_Data_Maybe_Last_bindLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_bindLast.Do(func() {
		cache_Data_Maybe_Last_bindLast = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3887487416_2748095225(Rebox_Data_Maybe_Last_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe()))))}
	})
	return cache_Data_Maybe_Last_bindLast
}

var cache_Data_Maybe_Last_applyLast gopurs_runtime.Value
var once_Data_Maybe_Last_applyLast sync.Once

func Get_Data_Maybe_Last_applyLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_applyLast.Do(func() {
		cache_Data_Maybe_Last_applyLast = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3552963512_3741347833(Rebox_Data_Maybe_Last_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe()))))}
	})
	return cache_Data_Maybe_Last_applyLast
}

var cache_Data_Maybe_Last_applicativeLast gopurs_runtime.Value
var once_Data_Maybe_Last_applicativeLast sync.Once

func Get_Data_Maybe_Last_applicativeLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_applicativeLast.Do(func() {
		cache_Data_Maybe_Last_applicativeLast = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_649684152_1439734649(Rebox_Data_Maybe_Last_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}
	})
	return cache_Data_Maybe_Last_applicativeLast
}

var cache_Data_Maybe_Last_altLast gopurs_runtime.Value
var once_Data_Maybe_Last_altLast sync.Once

func Get_Data_Maybe_Last_altLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_altLast.Do(func() {
		cache_Data_Maybe_Last_altLast = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3904200120_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3689823567_2812149806(Rebox_Data_Maybe_Last_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
		}), gopurs_runtime.RecordGet(Get_Data_Maybe_Last_semigroupLast(), "append")})))}
	})
	return cache_Data_Maybe_Last_altLast
}

var cache_Data_Maybe_Last_plusLast gopurs_runtime.Value
var once_Data_Maybe_Last_plusLast sync.Once

func Get_Data_Maybe_Last_plusLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_plusLast.Do(func() {
		cache_Data_Maybe_Last_plusLast = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_267397720_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3904200120_3421983481(Rebox_Data_Maybe_Last_3421983481_3904200120(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Data_Maybe_Last_altLast()))))}
		}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))}})))}
	})
	return cache_Data_Maybe_Last_plusLast
}

var cache_Data_Maybe_Last_alternativeLast gopurs_runtime.Value
var once_Data_Maybe_Last_alternativeLast sync.Once

func Get_Data_Maybe_Last_alternativeLast() gopurs_runtime.Value {
	once_Data_Maybe_Last_alternativeLast.Do(func() {
		cache_Data_Maybe_Last_alternativeLast = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_673668088_2307501113((&Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_649684152_1439734649(Rebox_Data_Maybe_Last_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_267397720_3706288089(Rebox_Data_Maybe_Last_3706288089_267397720(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_Maybe_Last_plusLast()))))}
		})})))}
	})
	return cache_Data_Maybe_Last_alternativeLast
}

func Call_Data_Maybe_Last_Last(x_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var x_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_0_loop
	_ = x_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Data_Maybe_Last_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_345859663_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Last ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Maybe_showMaybe(dictShow_0), "show"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1))}).StrVal()), ")"))
	})})))}
}

func Call_Data_Maybe_Last_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_2155612431_4177771502(Rebox_Data_Maybe_Last_4177771502_2155612431(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(dictOrd_0)))))}
}

func Call_Data_Maybe_Last_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_3508461103_3790796878(Rebox_Data_Maybe_Last_3790796878_3508461103(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(dictEq_0)))))}
}

func Call_Data_Maybe_Last_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_Last_591333647_2094947566(Rebox_Data_Maybe_Last_2094947566_591333647(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Call_Data_Maybe_boundedMaybe(dictBounded_0)))))}
}

func Rebox_Data_Maybe_Last_1089607855_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Maybe_Last_125234255_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_1439734649_649684152(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Monad[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_1662389854_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_170825214_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord1[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_1766074591_1662389854(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_1892488872_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_2094947566_591333647(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V2)
	return out
}

func Rebox_Data_Maybe_Last_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_2241343270_3818752775(in *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]) *Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_2568689657_1642601656(in *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_267397720_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3290176857_570732504(in *Constructor_Control_Extend_Extend[gopurs_runtime.Value]) *Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3421983481_3904200120(in *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) *Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_345859663_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3552963512_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3706288089_267397720(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3741347833_3552963512(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3790796878_3508461103(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3818752775_2241343270(in *Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3904200120_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Alt_Alt[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_3985601471_170825214(in *Constructor_Data_Ord_Ord1[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_4177771502_2155612431(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_570732504_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Extend_Extend[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_591333647_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Maybe_Last_649684152_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Maybe_Last_673668088_2307501113(in *Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Alternative_Alternative[gopurs_runtime.Value])(unsafe.Pointer(in))
}
