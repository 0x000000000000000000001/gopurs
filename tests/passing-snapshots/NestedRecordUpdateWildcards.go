package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_update gopurs_runtime.Value
var once_Main_update sync.Once

func Get_Main_update() gopurs_runtime.Value {
	once_Main_update.Do(func() {
		cache_Main_update = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_update(v_0_box, v1_1_box, v2_2_box, v3_3_box)
		})
	})
	return cache_Main_update
}

var cache_Main_init gopurs_runtime.Value
var once_Main_init sync.Once

func Get_Main_init() gopurs_runtime.Value {
	once_Main_init.Do(func() {
		cache_Main_init = gopurs_runtime.RecordDict2("bar", "foo", gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(2), gopurs_runtime.Int(3)), gopurs_runtime.Int(1))
	})
	return cache_Main_init
}

var cache_Main_expected gopurs_runtime.Value
var once_Main_expected sync.Once

func Get_Main_expected() gopurs_runtime.Value {
	once_Main_expected.Do(func() {
		cache_Main_expected = gopurs_runtime.RecordDict2("bar", "foo", gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(20), gopurs_runtime.Int(30)), gopurs_runtime.Int(10))
	})
	return cache_Main_expected
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func5(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value, dictEq2_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, r_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(dictEq_0_box, dictEq1_1_box, dictEq2_2_box, l_3_box, r_4_box)
		})
	})
	return cache_Main_check
}

var cache_Main_after gopurs_runtime.Value
var once_Main_after sync.Once

func Get_Main_after() gopurs_runtime.Value {
	once_Main_after.Do(func() {
		cache_Main_after = func() gopurs_runtime.Value {
			origVal := Get_Main_init()
			if origVal.Type != gopurs_runtime.TypeRecord2 {
				return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo", "bar"}, []gopurs_runtime.Value{gopurs_runtime.Int(10), gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(Get_Main_init(), "bar"), "baz", gopurs_runtime.Int(20), "qux", gopurs_runtime.Int(30))})
			}
			clone := *((*gopurs_runtime.RecordData2)(origVal.UnsafePtr))
			clone.V1 = gopurs_runtime.Int(10)
			clone.V0 = gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(Get_Main_init(), "bar"), "baz", gopurs_runtime.Int(20), "qux", gopurs_runtime.Int(30))
			return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
		}()
	})
	return cache_Main_after
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			_ = __local_var_0_0
			var __t1 gopurs_runtime.Value
			{
				if ((gopurs_runtime.RecordGet(Get_Main_after(), "foo").IntVal) == (gopurs_runtime.Int(10).IntVal)) && (((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_after(), "bar"), "baz").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_expected(), "bar"), "baz").IntVal)) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_after(), "bar"), "qux").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_expected(), "bar"), "qux").IntVal))) {
					__t1 = __local_var_0_0
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_1:
			return __t1
		}()
	})
	return cache_Main_main
}

func Call_Main_update(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value, v3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	var v2_2 gopurs_runtime.Value = v2_2_loop
	_ = v2_2
	var v3_3 gopurs_runtime.Value = v3_3_loop
	_ = v3_3
	return func() gopurs_runtime.Value {
		origVal := v_0
		if origVal.Type != gopurs_runtime.TypeRecord2 {
			return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo", "bar"}, []gopurs_runtime.Value{v1_1, gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v_0, "bar"), "baz", v2_2, "qux", v3_3)})
		}
		clone := *((*gopurs_runtime.RecordData2)(origVal.UnsafePtr))
		clone.V1 = v1_1
		clone.V0 = gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v_0, "bar"), "baz", v2_2, "qux", v3_3)
		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
	}()
}

func Call_Main_check(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value, dictEq2_2_loop gopurs_runtime.Value, l_3_loop gopurs_runtime.Value, r_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	var dictEq2_2 gopurs_runtime.Value = dictEq2_2_loop
	_ = dictEq2_2
	var l_3 gopurs_runtime.Value = l_3_loop
	_ = l_3
	var r_4 gopurs_runtime.Value = r_4_loop
	_ = r_4
	return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(l_3, "foo"), gopurs_runtime.RecordGet(r_4, "foo")).IntVal) != (0)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_3, "bar"), "baz"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_4, "bar"), "baz")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq2_2, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_3, "bar"), "qux"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_4, "bar"), "qux")).IntVal) != (0))))
}
