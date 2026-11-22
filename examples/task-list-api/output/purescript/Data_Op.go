package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Op_compose gopurs_runtime.Value
var once_Data_Op_compose sync.Once

func Get_Data_Op_compose() gopurs_runtime.Value {
	once_Data_Op_compose.Do(func() {
		cache_Data_Op_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Op_compose
}

var cache_Data_Op_compose1 gopurs_runtime.Value
var once_Data_Op_compose1 sync.Once

func Get_Data_Op_compose1() gopurs_runtime.Value {
	once_Data_Op_compose1.Do(func() {
		cache_Data_Op_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Op_compose1
}

var cache_Data_Op_Op gopurs_runtime.Value
var once_Data_Op_Op sync.Once

func Get_Data_Op_Op() gopurs_runtime.Value {
	once_Data_Op_Op.Do(func() {
		cache_Data_Op_Op = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Op_Op(x_0_box)
		})
	})
	return cache_Data_Op_Op
}

var cache_Data_Op_semigroupoidOp gopurs_runtime.Value
var once_Data_Op_semigroupoidOp sync.Once

func Get_Data_Op_semigroupoidOp() gopurs_runtime.Value {
	once_Data_Op_semigroupoidOp.Do(func() {
		cache_Data_Op_semigroupoidOp = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer((&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v1_1, v_0)
		})}))}
	})
	return cache_Data_Op_semigroupoidOp
}

var cache_Data_Op_semigroupOp gopurs_runtime.Value
var once_Data_Op_semigroupOp sync.Once

func Get_Data_Op_semigroupOp() gopurs_runtime.Value {
	once_Data_Op_semigroupOp.Do(func() {
		cache_Data_Op_semigroupOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Op_semigroupOp(dictSemigroup_0_box)
		})
	})
	return cache_Data_Op_semigroupOp
}

var cache_Data_Op_newtypeOp gopurs_runtime.Value
var once_Data_Op_newtypeOp sync.Once

func Get_Data_Op_newtypeOp() gopurs_runtime.Value {
	once_Data_Op_newtypeOp.Do(func() {
		cache_Data_Op_newtypeOp = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Op_newtypeOp
}

var cache_Data_Op_monoidOp gopurs_runtime.Value
var once_Data_Op_monoidOp sync.Once

func Get_Data_Op_monoidOp() gopurs_runtime.Value {
	once_Data_Op_monoidOp.Do(func() {
		cache_Data_Op_monoidOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Op_monoidOp(dictMonoid_0_box)
		})
	})
	return cache_Data_Op_monoidOp
}

var cache_Data_Op_contravariantOp gopurs_runtime.Value
var once_Data_Op_contravariantOp sync.Once

func Get_Data_Op_contravariantOp() gopurs_runtime.Value {
	once_Data_Op_contravariantOp.Do(func() {
		cache_Data_Op_contravariantOp = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_1, f_0)
		})}))}
	})
	return cache_Data_Op_contravariantOp
}

var cache_Data_Op_categoryOp gopurs_runtime.Value
var once_Data_Op_categoryOp sync.Once

func Get_Data_Op_categoryOp() gopurs_runtime.Value {
	once_Data_Op_categoryOp.Do(func() {
		cache_Data_Op_categoryOp = gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer((&Constructor_Control_Category_Category[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Data_Op_semigroupoidOp()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Data_Op_categoryOp
}

func Call_Data_Op_Op(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Op_semigroupOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Semigroup_semigroupFn(dictSemigroup_0)))}
}

func Call_Data_Op_monoidOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_monoidFn(dictMonoid_0)))}
}
