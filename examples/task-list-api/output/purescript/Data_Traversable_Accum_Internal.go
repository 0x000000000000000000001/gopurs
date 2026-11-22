package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Traversable_Accum_Internal_StateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_StateR sync.Once

func Get_Data_Traversable_Accum_Internal_StateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_StateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_StateR = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_Accum_Internal_StateR(x_0_box)
		})
	})
	return cache_Data_Traversable_Accum_Internal_StateR
}

var cache_Data_Traversable_Accum_Internal_StateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_StateL sync.Once

func Get_Data_Traversable_Accum_Internal_StateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_StateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_StateL = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_Accum_Internal_StateL(x_0_box)
		})
	})
	return cache_Data_Traversable_Accum_Internal_StateL
}

var cache_Data_Traversable_Accum_Internal_stateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateR sync.Once

func Get_Data_Traversable_Accum_Internal_stateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateR = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_Accum_Internal_stateR(v_0_box)
		})
	})
	return cache_Data_Traversable_Accum_Internal_stateR
}

var cache_Data_Traversable_Accum_Internal_stateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateL sync.Once

func Get_Data_Traversable_Accum_Internal_stateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateL = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Traversable_Accum_Internal_stateL(v_0_box)
		})
	})
	return cache_Data_Traversable_Accum_Internal_stateL
}

var cache_Data_Traversable_Accum_Internal_functorStateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateR sync.Once

func Get_Data_Traversable_Accum_Internal_functorStateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateR = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_0 shape=App(Other) bindingType=(Record (Row [accum: (TypeVar s$scope9), value: (TypeVar a$scope13)] Empty))
			v_3_0 := func() struct {
				accum gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
			} {
				orig := gopurs_runtime.Apply(k_1, s_2)
				_ = orig
				clone := struct {
					accum gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
				}{}
				clone.accum = gopurs_runtime.RecordGet(orig, "accum")
				clone.value = gopurs_runtime.RecordGet(orig, "value")
				return clone
			}()
			_ = v_3_0
			return gopurs_runtime.RecordDict2("accum", "value", v_3_0.accum, gopurs_runtime.Apply(f_0, v_3_0.value))
		})}))}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateR
}

var cache_Data_Traversable_Accum_Internal_functorStateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateL sync.Once

func Get_Data_Traversable_Accum_Internal_functorStateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateL = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_0 shape=App(Other) bindingType=(Record (Row [accum: (TypeVar s$scope18), value: (TypeVar a$scope22)] Empty))
			v_3_0 := func() struct {
				accum gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
			} {
				orig := gopurs_runtime.Apply(k_1, s_2)
				_ = orig
				clone := struct {
					accum gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
				}{}
				clone.accum = gopurs_runtime.RecordGet(orig, "accum")
				clone.value = gopurs_runtime.RecordGet(orig, "value")
				return clone
			}()
			_ = v_3_0
			return gopurs_runtime.RecordDict2("accum", "value", v_3_0.accum, gopurs_runtime.Apply(f_0, v_3_0.value))
		})}))}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateL
}

var cache_Data_Traversable_Accum_Internal_applyStateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateR sync.Once

func Get_Data_Traversable_Accum_Internal_applyStateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateR = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Traversable_Accum_Internal_functorStateR()
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_0 shape=App(Other) bindingType=(Record (Row [accum: (TypeVar s$scope27), value: (TypeVar a$scope31)] Empty))
			v_3_0 := func() struct {
				accum gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
			} {
				orig := gopurs_runtime.Apply(x_1, s_2)
				_ = orig
				clone := struct {
					accum gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
				}{}
				clone.accum = gopurs_runtime.RecordGet(orig, "accum")
				clone.value = gopurs_runtime.RecordGet(orig, "value")
				return clone
			}()
			_ = v_3_0
			// TAST (Let): v1_4_1 shape=App(Other) bindingType=(Record (Row [accum: (TypeVar s$scope27), value: (Func [(TypeVar a$scope31)] (TypeVar b$scope32))] Empty))
			v1_4_1 := func() struct {
				accum gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
			} {
				orig := gopurs_runtime.Apply(f_0, v_3_0.accum)
				_ = orig
				clone := struct {
					accum gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
				}{}
				clone.accum = gopurs_runtime.RecordGet(orig, "accum")
				clone.value = gopurs_runtime.RecordGet(orig, "value")
				return clone
			}()
			_ = v1_4_1
			return gopurs_runtime.RecordDict2("accum", "value", v1_4_1.accum, gopurs_runtime.Apply(v1_4_1.value, v_3_0.value))
		})}))}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateR
}

var cache_Data_Traversable_Accum_Internal_applyStateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateL sync.Once

func Get_Data_Traversable_Accum_Internal_applyStateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateL = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Traversable_Accum_Internal_functorStateL()
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_3_0 shape=App(Other) bindingType=(Record (Row [accum: (TypeVar s$scope36), value: (Func [(TypeVar a$scope40)] (TypeVar b$scope41))] Empty))
			v_3_0 := func() struct {
				accum gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
			} {
				orig := gopurs_runtime.Apply(f_0, s_2)
				_ = orig
				clone := struct {
					accum gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
				}{}
				clone.accum = gopurs_runtime.RecordGet(orig, "accum")
				clone.value = gopurs_runtime.RecordGet(orig, "value")
				return clone
			}()
			_ = v_3_0
			// TAST (Let): v1_4_1 shape=App(Other) bindingType=(Record (Row [accum: (TypeVar s$scope36), value: (TypeVar a$scope40)] Empty))
			v1_4_1 := func() struct {
				accum gopurs_runtime.Value
				value gopurs_runtime.Value
				/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
			} {
				orig := gopurs_runtime.Apply(x_1, v_3_0.accum)
				_ = orig
				clone := struct {
					accum gopurs_runtime.Value
					value gopurs_runtime.Value
					/*gopurs:record-v1:2:0061006300630075006d,00760061006c00750065*/
				}{}
				clone.accum = gopurs_runtime.RecordGet(orig, "accum")
				clone.value = gopurs_runtime.RecordGet(orig, "value")
				return clone
			}()
			_ = v1_4_1
			return gopurs_runtime.RecordDict2("accum", "value", v1_4_1.accum, gopurs_runtime.Apply(v_3_0.value, v1_4_1.value))
		})}))}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateL
}

var cache_Data_Traversable_Accum_Internal_applicativeStateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateR sync.Once

func Get_Data_Traversable_Accum_Internal_applicativeStateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateR = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Traversable_Accum_Internal_applyStateR()
		}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
		})}))}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateR
}

var cache_Data_Traversable_Accum_Internal_applicativeStateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateL sync.Once

func Get_Data_Traversable_Accum_Internal_applicativeStateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateL = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Traversable_Accum_Internal_applyStateL()
		}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
		})}))}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateL
}

func Call_Data_Traversable_Accum_Internal_StateR(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Traversable_Accum_Internal_StateL(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Traversable_Accum_Internal_stateR(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Data_Traversable_Accum_Internal_stateL(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}
