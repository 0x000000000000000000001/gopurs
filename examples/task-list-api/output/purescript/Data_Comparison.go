package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Comparison_go__append gopurs_runtime.Value
var once_Data_Comparison_go__append sync.Once

func Get_Data_Comparison_go__append() gopurs_runtime.Value {
	once_Data_Comparison_go__append.Do(func() {
		cache_Data_Comparison_go__append = gopurs_runtime.RecordGet(Call_Data_Semigroup_semigroupFn(Call_Data_Semigroup_semigroupFn(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Comparison_1625289059_4179793454(Rebox_Data_Comparison_4179793454_1625289059(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Ordering_semigroupOrdering()))))})), "append")
	})
	return cache_Data_Comparison_go__append
}

var cache_Data_Comparison_Comparison gopurs_runtime.Value
var once_Data_Comparison_Comparison sync.Once

func Get_Data_Comparison_Comparison() gopurs_runtime.Value {
	once_Data_Comparison_Comparison.Do(func() {
		cache_Data_Comparison_Comparison = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Comparison_Comparison(x_0_box)
		})
	})
	return cache_Data_Comparison_Comparison
}

var cache_Data_Comparison_semigroupComparison gopurs_runtime.Value
var once_Data_Comparison_semigroupComparison sync.Once

func Get_Data_Comparison_semigroupComparison() gopurs_runtime.Value {
	once_Data_Comparison_semigroupComparison.Do(func() {
		cache_Data_Comparison_semigroupComparison = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Semigroup_semigroupFn(Call_Data_Semigroup_semigroupFn(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Comparison_1625289059_4179793454(Rebox_Data_Comparison_4179793454_1625289059(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Ordering_semigroupOrdering()))))})), "append"), v_0, v1_1)
		})}))}
	})
	return cache_Data_Comparison_semigroupComparison
}

var cache_Data_Comparison_newtypeComparison gopurs_runtime.Value
var once_Data_Comparison_newtypeComparison sync.Once

func Get_Data_Comparison_newtypeComparison() gopurs_runtime.Value {
	once_Data_Comparison_newtypeComparison.Do(func() {
		cache_Data_Comparison_newtypeComparison = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Comparison_newtypeComparison
}

var cache_Data_Comparison_monoidComparison gopurs_runtime.Value
var once_Data_Comparison_monoidComparison sync.Once

func Get_Data_Comparison_monoidComparison() gopurs_runtime.Value {
	once_Data_Comparison_monoidComparison.Do(func() {
		cache_Data_Comparison_monoidComparison = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Comparison_semigroupComparison()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Comparison_monoidComparison
}

var cache_Data_Comparison_defaultComparison gopurs_runtime.Value
var once_Data_Comparison_defaultComparison sync.Once

func Get_Data_Comparison_defaultComparison() gopurs_runtime.Value {
	once_Data_Comparison_defaultComparison.Do(func() {
		cache_Data_Comparison_defaultComparison = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Comparison_defaultComparison(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Data_Comparison_defaultComparison
}

var cache_Data_Comparison_contravariantComparison gopurs_runtime.Value
var once_Data_Comparison_contravariantComparison sync.Once

func Get_Data_Comparison_contravariantComparison() gopurs_runtime.Value {
	once_Data_Comparison_contravariantComparison.Do(func() {
		cache_Data_Comparison_contravariantComparison = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
		})}))}
	})
	return cache_Data_Comparison_contravariantComparison
}

func Call_Data_Comparison_Comparison(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Comparison_defaultComparison(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	return Call_Data_Ord_compare(dictOrd_0)
}

func Rebox_Data_Comparison_1625289059_4179793454(in *Constructor_Data_Semigroup_Semigroup[uint32]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Comparison_4179793454_1625289059(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[uint32])(unsafe.Pointer(in))
}
