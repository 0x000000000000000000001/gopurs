package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Store_go__map gopurs_runtime.Value
var once_Control_Comonad_Store_go__map sync.Once

func Get_Control_Comonad_Store_go__map() gopurs_runtime.Value {
	once_Control_Comonad_Store_go__map.Do(func() {
		cache_Control_Comonad_Store_go__map = gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map")
	})
	return cache_Control_Comonad_Store_go__map
}

var cache_Control_Comonad_Store_unwrap gopurs_runtime.Value
var once_Control_Comonad_Store_unwrap sync.Once

func Get_Control_Comonad_Store_unwrap() gopurs_runtime.Value {
	once_Control_Comonad_Store_unwrap.Do(func() {
		cache_Control_Comonad_Store_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Comonad_Store_unwrap
}

var cache_Control_Comonad_Store_store gopurs_runtime.Value
var once_Control_Comonad_Store_store sync.Once

func Get_Control_Comonad_Store_store() gopurs_runtime.Value {
	once_Control_Comonad_Store_store.Do(func() {
		cache_Control_Comonad_Store_store = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Store_store(f_0_box, x_1_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Store_store
}

var cache_Control_Comonad_Store_runStore gopurs_runtime.Value
var once_Control_Comonad_Store_runStore sync.Once

func Get_Control_Comonad_Store_runStore() gopurs_runtime.Value {
	once_Control_Comonad_Store_runStore.Do(func() {
		cache_Control_Comonad_Store_runStore = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Store_runStore(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Comonad_Store_runStore
}

func Call_Control_Comonad_Store_store(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{gopurs_runtime.Value{}, gopurs_runtime.Value{}}
}

func Call_Control_Comonad_Store_runStore(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	// TAST (Let): __local_var_1_0 shape=Other bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar s$scope12), (Func [(TypeVar s$scope12)] (TypeVar a$scope13))])
	__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{(v_0).V1, (v_0).V0}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}())
	_ = __local_var_1_0
	// TAST (Let): __local_var_2_1 shape=Other bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar s$scope12), (Func [(TypeVar s$scope12)] (TypeVar a$scope13))])
	__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{(__local_var_1_0).V0, gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), (__local_var_1_0).V1)}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}())
	_ = __local_var_2_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
		}{(__local_var_2_1).V1, (__local_var_2_1).V0}
		return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
	}()))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}
