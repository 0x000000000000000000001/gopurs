package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_init gopurs_runtime.Value
var once_Main_init sync.Once

func Get_Main_init() gopurs_runtime.Value {
	once_Main_init.Do(func() {
		cache_Main_init = gopurs_runtime.RecordDict2("bar", "foo", gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(2), gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(3), gopurs_runtime.Int(4))), gopurs_runtime.Int(1))
	})
	return cache_Main_init
}

var cache_Main_updated gopurs_runtime.Value
var once_Main_updated sync.Once

func Get_Main_updated() gopurs_runtime.Value {
	once_Main_updated.Do(func() {
		cache_Main_updated = func() gopurs_runtime.Value {
			origVal := Get_Main_init()
			if origVal.Type != gopurs_runtime.TypeRecord2 {
				return gopurs_runtime.RecordUpdateDict(origVal, []string{"foo", "bar"}, []gopurs_runtime.Value{gopurs_runtime.Int(10), gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(Get_Main_init(), "bar"), "baz", gopurs_runtime.Int(20), "qux", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_init(), "bar"), "qux"), "lhs", gopurs_runtime.Int(30), "rhs", gopurs_runtime.Int(40)))})
			}
			clone := *((*gopurs_runtime.RecordData2)(origVal.UnsafePtr))
			clone.V1 = gopurs_runtime.Int(10)
			clone.V0 = gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(Get_Main_init(), "bar"), "baz", gopurs_runtime.Int(20), "qux", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_init(), "bar"), "qux"), "lhs", gopurs_runtime.Int(30), "rhs", gopurs_runtime.Int(40)))
			return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
		}()
	})
	return cache_Main_updated
}

var cache_Main_expected gopurs_runtime.Value
var once_Main_expected sync.Once

func Get_Main_expected() gopurs_runtime.Value {
	once_Main_expected.Do(func() {
		cache_Main_expected = gopurs_runtime.RecordDict2("bar", "foo", gopurs_runtime.RecordDict2("baz", "qux", gopurs_runtime.Int(20), gopurs_runtime.RecordDict2("lhs", "rhs", gopurs_runtime.Int(30), gopurs_runtime.Int(40))), gopurs_runtime.Int(10))
	})
	return cache_Main_expected
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func6(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value, dictEq2_2_box gopurs_runtime.Value, dictEq3_3_box gopurs_runtime.Value, l_4_box gopurs_runtime.Value, r_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(dictEq_0_box, dictEq1_1_box, dictEq2_2_box, dictEq3_3_box, l_4_box, r_5_box)
		})
	})
	return cache_Main_check
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
				if ((gopurs_runtime.RecordGet(Get_Main_updated(), "foo").IntVal) == (gopurs_runtime.Int(10).IntVal)) && (((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_updated(), "bar"), "baz").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_expected(), "bar"), "baz").IntVal)) && (((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_updated(), "bar"), "qux"), "lhs").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_expected(), "bar"), "qux"), "lhs").IntVal)) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_updated(), "bar"), "qux"), "rhs").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Get_Main_expected(), "bar"), "qux"), "rhs").IntVal)))) {
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

func Call_Main_check(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value, dictEq2_2_loop gopurs_runtime.Value, dictEq3_3_loop gopurs_runtime.Value, l_4_loop gopurs_runtime.Value, r_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	var dictEq2_2 gopurs_runtime.Value = dictEq2_2_loop
	_ = dictEq2_2
	var dictEq3_3 gopurs_runtime.Value = dictEq3_3_loop
	_ = dictEq3_3
	var l_4 gopurs_runtime.Value = l_4_loop
	_ = l_4
	var r_5 gopurs_runtime.Value = r_5_loop
	_ = r_5
	return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(l_4, "foo"), gopurs_runtime.RecordGet(r_5, "foo")).IntVal) != (0)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_4, "bar"), "baz"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_5, "bar"), "baz")).IntVal) != (0)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq2_2, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_4, "bar"), "qux"), "lhs"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_5, "bar"), "qux"), "lhs")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(l_4, "bar"), "qux"), "rhs"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(r_5, "bar"), "qux"), "rhs")).IntVal) != (0)))))
}
