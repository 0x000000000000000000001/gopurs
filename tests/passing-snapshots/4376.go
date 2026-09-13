package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_v1 gopurs_runtime.Value
var once_Main_v1 sync.Once

func Get_Main_v1() gopurs_runtime.Value {
	once_Main_v1.Do(func() {
		cache_Main_v1 = func() gopurs_runtime.Value {
			orig := struct {
				a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
			}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{Get_Data_Unit_unit(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())}
			_ = orig
			return gopurs_runtime.RecordDict1("a", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.a)})
		}()
	})
	return cache_Main_v1
}

var cache_Main_v2 gopurs_runtime.Value
var once_Main_v2 sync.Once

func Get_Main_v2() gopurs_runtime.Value {
	once_Main_v2.Do(func() {
		cache_Main_v2 = func() gopurs_runtime.Value {
			orig := func() struct {
				a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
			} {
				clone := func() struct {
					a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				} {
					orig := Get_Main_v1()
					_ = orig
					clone := struct {
						a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					}{}
					clone.a = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "a"))
					return clone
				}()
				clone.a = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(Call_Data_Maybe_monoidMaybe(gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupUnit()))}), "mempty"))
				return clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict1("a", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.a)})
		}()
	})
	return cache_Main_v2
}

var cache_Main_union gopurs_runtime.Value
var once_Main_union sync.Once

func Get_Main_union() gopurs_runtime.Value {
	once_Main_union.Do(func() {
		cache_Main_union = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v3_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_union(_dollar___unused_0_box, v_1_box, v3_2_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_union
}

var cache_Main_shouldSolve gopurs_runtime.Value
var once_Main_shouldSolve sync.Once

func Get_Main_shouldSolve() gopurs_runtime.Value {
	once_Main_shouldSolve.Do(func() {
		cache_Main_shouldSolve = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_shouldSolve
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_asNothing gopurs_runtime.Value
var once_Main_asNothing sync.Once

func Get_Main_asNothing() gopurs_runtime.Value {
	once_Main_asNothing.Do(func() {
		cache_Main_asNothing = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_asNothing(func() struct {
					a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
				} {
					orig := v_0_box
					_ = orig
					clone := struct {
						a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
					}{}
					clone.a = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "a"))
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict1("a", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.a)})
			}()
		})
	})
	return cache_Main_asNothing
}

func Call_Main_union(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v3_2_loop gopurs_runtime.Value) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var v3_2 gopurs_runtime.Value = v3_2_loop
	_ = v3_2
	return 513803634
}

func Call_Main_asNothing(v_0_loop struct {
	a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
}) struct {
	a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
} {
	var v_0 struct {
		a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	} = v_0_loop
	_ = v_0
	return func() struct {
		a *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	} {
		clone := v_0
		clone.a = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}())
		return clone
	}()
}
