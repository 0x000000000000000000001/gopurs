package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_First_go__append gopurs_runtime.Value
var once_Data_Semigroup_First_go__append sync.Once

func Get_Data_Semigroup_First_go__append() gopurs_runtime.Value {
	once_Data_Semigroup_First_go__append.Do(func() {
		cache_Data_Semigroup_First_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Semigroup_First_go__append
}

var cache_Data_Semigroup_First_First gopurs_runtime.Value
var once_Data_Semigroup_First_First sync.Once

func Get_Data_Semigroup_First_First() gopurs_runtime.Value {
	once_Data_Semigroup_First_First.Do(func() {
		cache_Data_Semigroup_First_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_First_First(x_0_box)
		})
	})
	return cache_Data_Semigroup_First_First
}

var cache_Data_Semigroup_First_showFirst gopurs_runtime.Value
var once_Data_Semigroup_First_showFirst sync.Once

func Get_Data_Semigroup_First_showFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_showFirst.Do(func() {
		cache_Data_Semigroup_First_showFirst = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_First_showFirst(dictShow_0_box)
		})
	})
	return cache_Data_Semigroup_First_showFirst
}

var cache_Data_Semigroup_First_semigroupFirst gopurs_runtime.Value
var once_Data_Semigroup_First_semigroupFirst sync.Once

func Get_Data_Semigroup_First_semigroupFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_semigroupFirst.Do(func() {
		cache_Data_Semigroup_First_semigroupFirst = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Data_Semigroup_First_semigroupFirst
}

var cache_Data_Semigroup_First_ordFirst gopurs_runtime.Value
var once_Data_Semigroup_First_ordFirst sync.Once

func Get_Data_Semigroup_First_ordFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_ordFirst.Do(func() {
		cache_Data_Semigroup_First_ordFirst = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_First_ordFirst(dictOrd_0_box)
		})
	})
	return cache_Data_Semigroup_First_ordFirst
}

var cache_Data_Semigroup_First_functorFirst gopurs_runtime.Value
var once_Data_Semigroup_First_functorFirst sync.Once

func Get_Data_Semigroup_First_functorFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_functorFirst.Do(func() {
		cache_Data_Semigroup_First_functorFirst = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, m_1)
		})}))}
	})
	return cache_Data_Semigroup_First_functorFirst
}

var cache_Data_Semigroup_First_eqFirst gopurs_runtime.Value
var once_Data_Semigroup_First_eqFirst sync.Once

func Get_Data_Semigroup_First_eqFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_eqFirst.Do(func() {
		cache_Data_Semigroup_First_eqFirst = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_First_eqFirst(dictEq_0_box)
		})
	})
	return cache_Data_Semigroup_First_eqFirst
}

var cache_Data_Semigroup_First_eq1First gopurs_runtime.Value
var once_Data_Semigroup_First_eq1First sync.Once

func Get_Data_Semigroup_First_eq1First() gopurs_runtime.Value {
	once_Data_Semigroup_First_eq1First.Do(func() {
		cache_Data_Semigroup_First_eq1First = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))
		})}))}
	})
	return cache_Data_Semigroup_First_eq1First
}

var cache_Data_Semigroup_First_ord1First gopurs_runtime.Value
var once_Data_Semigroup_First_ord1First sync.Once

func Get_Data_Semigroup_First_ord1First() gopurs_runtime.Value {
	once_Data_Semigroup_First_ord1First.Do(func() {
		cache_Data_Semigroup_First_ord1First = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Semigroup_First_eq1First()))}
		}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))
		})}))}
	})
	return cache_Data_Semigroup_First_ord1First
}

var cache_Data_Semigroup_First_boundedFirst gopurs_runtime.Value
var once_Data_Semigroup_First_boundedFirst sync.Once

func Get_Data_Semigroup_First_boundedFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_boundedFirst.Do(func() {
		cache_Data_Semigroup_First_boundedFirst = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semigroup_First_boundedFirst(dictBounded_0_box)
		})
	})
	return cache_Data_Semigroup_First_boundedFirst
}

var cache_Data_Semigroup_First_applyFirst gopurs_runtime.Value
var once_Data_Semigroup_First_applyFirst sync.Once

func Get_Data_Semigroup_First_applyFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_applyFirst.Do(func() {
		cache_Data_Semigroup_First_applyFirst = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Semigroup_First_functorFirst()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v_0, v1_1)
		})}))}
	})
	return cache_Data_Semigroup_First_applyFirst
}

var cache_Data_Semigroup_First_bindFirst gopurs_runtime.Value
var once_Data_Semigroup_First_bindFirst sync.Once

func Get_Data_Semigroup_First_bindFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_bindFirst.Do(func() {
		cache_Data_Semigroup_First_bindFirst = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Semigroup_First_applyFirst()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_1, v_0)
		})}))}
	})
	return cache_Data_Semigroup_First_bindFirst
}

var cache_Data_Semigroup_First_applicativeFirst gopurs_runtime.Value
var once_Data_Semigroup_First_applicativeFirst sync.Once

func Get_Data_Semigroup_First_applicativeFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_applicativeFirst.Do(func() {
		cache_Data_Semigroup_First_applicativeFirst = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Semigroup_First_applyFirst()))}
		}), Get_Data_Semigroup_First_First()}))}
	})
	return cache_Data_Semigroup_First_applicativeFirst
}

var cache_Data_Semigroup_First_monadFirst gopurs_runtime.Value
var once_Data_Semigroup_First_monadFirst sync.Once

func Get_Data_Semigroup_First_monadFirst() gopurs_runtime.Value {
	once_Data_Semigroup_First_monadFirst.Do(func() {
		cache_Data_Semigroup_First_monadFirst = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Semigroup_First_applicativeFirst()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Semigroup_First_bindFirst()))}
		})}))}
	})
	return cache_Data_Semigroup_First_monadFirst
}

func Call_Data_Semigroup_First_First(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Semigroup_First_showFirst(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(First ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Semigroup_First_ordFirst(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Semigroup_First_eqFirst(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Semigroup_First_boundedFirst(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}
