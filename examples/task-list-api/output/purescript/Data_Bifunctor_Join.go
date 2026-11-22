package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bifunctor_Join_go__append gopurs_runtime.Value
var once_Data_Bifunctor_Join_go__append sync.Once

func Get_Data_Bifunctor_Join_go__append() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_go__append.Do(func() {
		cache_Data_Bifunctor_Join_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Bifunctor_Join_go__append
}

var cache_Data_Bifunctor_Join_Join gopurs_runtime.Value
var once_Data_Bifunctor_Join_Join sync.Once

func Get_Data_Bifunctor_Join_Join() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_Join.Do(func() {
		cache_Data_Bifunctor_Join_Join = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_Join(x_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_Join
}

var cache_Data_Bifunctor_Join_showJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_showJoin sync.Once

func Get_Data_Bifunctor_Join_showJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_showJoin.Do(func() {
		cache_Data_Bifunctor_Join_showJoin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_showJoin(dictShow_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_showJoin
}

var cache_Data_Bifunctor_Join_ordJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_ordJoin sync.Once

func Get_Data_Bifunctor_Join_ordJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_ordJoin.Do(func() {
		cache_Data_Bifunctor_Join_ordJoin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_ordJoin(dictOrd_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_ordJoin
}

var cache_Data_Bifunctor_Join_newtypeJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_newtypeJoin sync.Once

func Get_Data_Bifunctor_Join_newtypeJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_newtypeJoin.Do(func() {
		cache_Data_Bifunctor_Join_newtypeJoin = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Bifunctor_Join_newtypeJoin
}

var cache_Data_Bifunctor_Join_eqJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_eqJoin sync.Once

func Get_Data_Bifunctor_Join_eqJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_eqJoin.Do(func() {
		cache_Data_Bifunctor_Join_eqJoin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_eqJoin(dictEq_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_eqJoin
}

var cache_Data_Bifunctor_Join_bifunctorJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_bifunctorJoin sync.Once

func Get_Data_Bifunctor_Join_bifunctorJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_bifunctorJoin.Do(func() {
		cache_Data_Bifunctor_Join_bifunctorJoin = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_bifunctorJoin(dictBifunctor_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_bifunctorJoin
}

var cache_Data_Bifunctor_Join_biapplyJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_biapplyJoin sync.Once

func Get_Data_Bifunctor_Join_biapplyJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_biapplyJoin.Do(func() {
		cache_Data_Bifunctor_Join_biapplyJoin = gopurs_runtime.Func(func(dictBiapply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_biapplyJoin(dictBiapply_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_biapplyJoin
}

var cache_Data_Bifunctor_Join_biapplicativeJoin gopurs_runtime.Value
var once_Data_Bifunctor_Join_biapplicativeJoin sync.Once

func Get_Data_Bifunctor_Join_biapplicativeJoin() gopurs_runtime.Value {
	once_Data_Bifunctor_Join_biapplicativeJoin.Do(func() {
		cache_Data_Bifunctor_Join_biapplicativeJoin = gopurs_runtime.Func(func(dictBiapplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bifunctor_Join_biapplicativeJoin(dictBiapplicative_0_box)
		})
	})
	return cache_Data_Bifunctor_Join_biapplicativeJoin
}

func Call_Data_Bifunctor_Join_Join(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Bifunctor_Join_showJoin(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Join ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Bifunctor_Join_ordJoin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Bifunctor_Join_eqJoin(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Bifunctor_Join_bifunctorJoin(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
	_ = dictBifunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, f_1, v_2)
	})}))}
}

func Call_Data_Bifunctor_Join_biapplyJoin(dictBiapply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBiapply_0 gopurs_runtime.Value = dictBiapply_0_loop
	_ = dictBiapply_0
	// TAST (Let): bifunctorJoin1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar a), (TypeVar a)]) [(TypeVar p$scope31)])])
	bifunctorJoin1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_Bifunctor_Join_bifunctorJoin(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{})))
	_ = bifunctorJoin1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(bifunctorJoin1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), v_2, v1_3)
	})}))}
}

func Call_Data_Bifunctor_Join_biapplicativeJoin(dictBiapplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBiapplicative_0 gopurs_runtime.Value = dictBiapplicative_0_loop
	_ = dictBiapplicative_0
	// TAST (Let): biapplyJoin1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar a), (TypeVar a)]) [(TypeVar p$scope41)])])
	biapplyJoin1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Bifunctor_Join_biapplyJoin(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapplicative_0, "Biapply0"), gopurs_runtime.Value{})))
	_ = biapplyJoin1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(biapplyJoin1_1_0)}
	}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapplicative_0, "bipure"), a_2, a_2)
	})}))}
}
