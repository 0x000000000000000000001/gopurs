package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_Last_go__append gopurs_runtime.Value
var once_Data_Semigroup_Last_go__append sync.Once

func Get_Data_Semigroup_Last_go__append() gopurs_runtime.Value {
	once_Data_Semigroup_Last_go__append.Do(func() {
		cache_Data_Semigroup_Last_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Semigroup_Last_go__append
}

var cache_Data_Semigroup_Last_Last gopurs_runtime.Value
var once_Data_Semigroup_Last_Last sync.Once

func Get_Data_Semigroup_Last_Last() gopurs_runtime.Value {
	once_Data_Semigroup_Last_Last.Do(func() {
		cache_Data_Semigroup_Last_Last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Last_Last(x_0_box)
		})
	})
	return cache_Data_Semigroup_Last_Last
}

var cache_Data_Semigroup_Last_showLast gopurs_runtime.Value
var once_Data_Semigroup_Last_showLast sync.Once

func Get_Data_Semigroup_Last_showLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_showLast.Do(func() {
		cache_Data_Semigroup_Last_showLast = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Last_showLast(dictShow_0_box)
		})
	})
	return cache_Data_Semigroup_Last_showLast
}

var cache_Data_Semigroup_Last_semigroupLast gopurs_runtime.Value
var once_Data_Semigroup_Last_semigroupLast sync.Once

func Get_Data_Semigroup_Last_semigroupLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_semigroupLast.Do(func() {
		cache_Data_Semigroup_Last_semigroupLast = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_1
		})}))}
	})
	return cache_Data_Semigroup_Last_semigroupLast
}

var cache_Data_Semigroup_Last_ordLast gopurs_runtime.Value
var once_Data_Semigroup_Last_ordLast sync.Once

func Get_Data_Semigroup_Last_ordLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_ordLast.Do(func() {
		cache_Data_Semigroup_Last_ordLast = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Last_ordLast(dictOrd_0_box)
		})
	})
	return cache_Data_Semigroup_Last_ordLast
}

var cache_Data_Semigroup_Last_functorLast gopurs_runtime.Value
var once_Data_Semigroup_Last_functorLast sync.Once

func Get_Data_Semigroup_Last_functorLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_functorLast.Do(func() {
		cache_Data_Semigroup_Last_functorLast = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, m_1)
		})}))}
	})
	return cache_Data_Semigroup_Last_functorLast
}

var cache_Data_Semigroup_Last_eqLast gopurs_runtime.Value
var once_Data_Semigroup_Last_eqLast sync.Once

func Get_Data_Semigroup_Last_eqLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_eqLast.Do(func() {
		cache_Data_Semigroup_Last_eqLast = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Last_eqLast(dictEq_0_box)
		})
	})
	return cache_Data_Semigroup_Last_eqLast
}

var cache_Data_Semigroup_Last_eq1Last gopurs_runtime.Value
var once_Data_Semigroup_Last_eq1Last sync.Once

func Get_Data_Semigroup_Last_eq1Last() gopurs_runtime.Value {
	once_Data_Semigroup_Last_eq1Last.Do(func() {
		cache_Data_Semigroup_Last_eq1Last = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))
		})}))}
	})
	return cache_Data_Semigroup_Last_eq1Last
}

var cache_Data_Semigroup_Last_ord1Last gopurs_runtime.Value
var once_Data_Semigroup_Last_ord1Last sync.Once

func Get_Data_Semigroup_Last_ord1Last() gopurs_runtime.Value {
	once_Data_Semigroup_Last_ord1Last.Do(func() {
		cache_Data_Semigroup_Last_ord1Last = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Semigroup_Last_eq1Last()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
		})}))}
	})
	return cache_Data_Semigroup_Last_ord1Last
}

var cache_Data_Semigroup_Last_boundedLast gopurs_runtime.Value
var once_Data_Semigroup_Last_boundedLast sync.Once

func Get_Data_Semigroup_Last_boundedLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_boundedLast.Do(func() {
		cache_Data_Semigroup_Last_boundedLast = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_Last_boundedLast(dictBounded_0_box)
		})
	})
	return cache_Data_Semigroup_Last_boundedLast
}

var cache_Data_Semigroup_Last_applyLast gopurs_runtime.Value
var once_Data_Semigroup_Last_applyLast sync.Once

func Get_Data_Semigroup_Last_applyLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_applyLast.Do(func() {
		cache_Data_Semigroup_Last_applyLast = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Semigroup_Last_functorLast()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_0, v1_1)
		})}))}
	})
	return cache_Data_Semigroup_Last_applyLast
}

var cache_Data_Semigroup_Last_bindLast gopurs_runtime.Value
var once_Data_Semigroup_Last_bindLast sync.Once

func Get_Data_Semigroup_Last_bindLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_bindLast.Do(func() {
		cache_Data_Semigroup_Last_bindLast = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Semigroup_Last_applyLast()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_1, v_0)
		})}))}
	})
	return cache_Data_Semigroup_Last_bindLast
}

var cache_Data_Semigroup_Last_applicativeLast gopurs_runtime.Value
var once_Data_Semigroup_Last_applicativeLast sync.Once

func Get_Data_Semigroup_Last_applicativeLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_applicativeLast.Do(func() {
		cache_Data_Semigroup_Last_applicativeLast = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Semigroup_Last_applyLast()))}
		}), Get_Data_Semigroup_Last_Last()}))}
	})
	return cache_Data_Semigroup_Last_applicativeLast
}

var cache_Data_Semigroup_Last_monadLast gopurs_runtime.Value
var once_Data_Semigroup_Last_monadLast sync.Once

func Get_Data_Semigroup_Last_monadLast() gopurs_runtime.Value {
	once_Data_Semigroup_Last_monadLast.Do(func() {
		cache_Data_Semigroup_Last_monadLast = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Semigroup_Last_applicativeLast()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Semigroup_Last_bindLast()))}
		})}))}
	})
	return cache_Data_Semigroup_Last_monadLast
}

func Call_Data_Semigroup_Last_Last(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Semigroup_Last_showLast(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Last ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Semigroup_Last_ordLast(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Semigroup_Last_eqLast(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Semigroup_Last_boundedLast(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}
