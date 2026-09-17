package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_131790935_3790796878(Rebox_Main_3790796878_131790935(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_sortEntries gopurs_runtime.Value
var once_Main_sortEntries sync.Once

func Get_Main_sortEntries() gopurs_runtime.Value {
	once_Main_sortEntries.Do(func() {
		cache_Main_sortEntries = gopurs_runtime.Apply(Get_Data_Array_sortBy__485191994(), gopurs_runtime.Apply(Get_Data_Ord_comparing__93829917(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.RecordGet(v_0, "key").IntVal)
		})))
	})
	return cache_Main_sortEntries
}

var cache_Main_defaultContext gopurs_runtime.Value
var once_Main_defaultContext sync.Once

func Get_Main_defaultContext() gopurs_runtime.Value {
	once_Main_defaultContext.Do(func() {
		cache_Main_defaultContext = func() gopurs_runtime.Value {
			orig := struct {
				bound   int64
				depth   int64
				options struct {
					isTail bool
				}
				tcoIdent bool
			}{int64(0), int64(0), struct {
				isTail bool
			}{false}, false}
			_ = orig
			return gopurs_runtime.RecordDict4("bound", "depth", "options", "tcoIdent", gopurs_runtime.Int(orig.bound), gopurs_runtime.Int(orig.depth), func() gopurs_runtime.Value {
				orig := orig.options
				_ = orig
				return gopurs_runtime.RecordDict1("isTail", gopurs_runtime.Bool(orig.isTail))
			}(), gopurs_runtime.Bool(orig.tcoIdent))
		}()
	})
	return cache_Main_defaultContext
}

var cache_Main_updateContext gopurs_runtime.Value
var once_Main_updateContext sync.Once

func Get_Main_updateContext() gopurs_runtime.Value {
	once_Main_updateContext.Do(func() {
		cache_Main_updateContext = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_updateContext(value_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict4("bound", "depth", "options", "tcoIdent", gopurs_runtime.Int(orig.bound), gopurs_runtime.Int(orig.depth), func() gopurs_runtime.Value {
					orig := orig.options
					_ = orig
					return gopurs_runtime.RecordDict1("isTail", gopurs_runtime.Bool(orig.isTail))
				}(), gopurs_runtime.Bool(orig.tcoIdent))
			}()
		})
	})
	return cache_Main_updateContext
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func2(func(name_0_box gopurs_runtime.Value, condition_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(name_0_box.StrVal(), (condition_1_box.IntVal) != (0))
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t0 string
			{
				if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := func() gopurs_runtime.Value {
						arr := Call_Data_Array_sortBy__485191994(gopurs_runtime.Apply(Get_Data_Ord_comparing__93829917(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int(gopurs_runtime.RecordGet(v_0, "key").IntVal)
						})), []struct {
							key  int64
							name string
						}{struct {
							key  int64
							name string
						}{int64(2), "new"}, struct {
							key  int64
							name string
						}{int64(2), "old"}, struct {
							key  int64
							name string
						}{int64(1), "first"}})
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								orig := v
								_ = orig
								return gopurs_runtime.RecordDict2("key", "name", gopurs_runtime.Int(orig.key), gopurs_runtime.Str(orig.name))
							}()
						}
						return gopurs_runtime.Array(boxed)
					}()
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Str(gopurs_runtime.RecordGet(v_0, "name").StrVal())
						}), v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr))), func() gopurs_runtime.Value {
					arr := []string{"first", "new", "old"}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0) {
					__t0 = "stable duplicate keys"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "Fail: stable duplicate keys"
			}
		end_branch_0:
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 string
				{
					if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
						arr_val_arrayMap6 := func() gopurs_runtime.Value {
							arr := []struct {
								key  int64
								name string
							}{struct {
								key  int64
								name string
							}{int64(2), "new"}, struct {
								key  int64
								name string
							}{int64(2), "old"}, struct {
								key  int64
								name string
							}{int64(1), "first"}}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = func() gopurs_runtime.Value {
									orig := v
									_ = orig
									return gopurs_runtime.RecordDict2("key", "name", gopurs_runtime.Int(orig.key), gopurs_runtime.Str(orig.name))
								}()
							}
							return gopurs_runtime.Array(boxed)
						}()
						_ = arr_val_arrayMap6
						arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
						_ = arr_go_arrayMap6
						res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
						_ = res_go_arrayMap6
						for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
							res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Str(gopurs_runtime.RecordGet(v_1, "name").StrVal())
							}), v_arrayMap6)
						}
						return gopurs_runtime.Array(res_go_arrayMap6)
					}()).UnsafePtr))), func() gopurs_runtime.Value {
						arr := []string{"new", "old", "first"}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0) {
						__t1 = "persistent input"
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = "Fail: persistent input"
				}
			end_branch_1:
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t1)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("constant record updates")), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					}))
				}))
			}))
		}()
	})
	return cache_Main_main
}

func Call_Main_updateContext(value_0_loop int64) struct {
	bound   int64
	depth   int64
	options struct {
		isTail bool
	}
	tcoIdent bool
} {
	var value_0 int64 = value_0_loop
	_ = value_0
	return func() struct {
		bound   int64
		depth   int64
		options struct {
			isTail bool
		}
		tcoIdent bool
	} {
		clone := func() struct {
			bound   int64
			depth   int64
			options struct {
				isTail bool
			}
			tcoIdent bool
		} {
			orig := Get_Main_defaultContext()
			_ = orig
			clone := struct {
				bound   int64
				depth   int64
				options struct {
					isTail bool
				}
				tcoIdent bool
			}{}
			clone.bound = gopurs_runtime.RecordGet(orig, "bound").IntVal
			clone.depth = gopurs_runtime.RecordGet(orig, "depth").IntVal
			clone.options = func() struct {
				isTail bool
			} {
				orig := gopurs_runtime.RecordGet(orig, "options")
				_ = orig
				clone := struct {
					isTail bool
				}{}
				clone.isTail = (gopurs_runtime.RecordGet(orig, "isTail").IntVal) != (0)
				return clone
			}()
			clone.tcoIdent = (gopurs_runtime.RecordGet(orig, "tcoIdent").IntVal) != (0)
			return clone
		}()
		clone.tcoIdent = true
		clone.options = struct {
			isTail bool
		}{true}
		clone.bound = value_0
		return clone
	}()
}

func Call_Main_check(name_0_loop string, condition_1_loop bool) gopurs_runtime.Value {
	var name_0 string = name_0_loop
	_ = name_0
	var condition_1 bool = condition_1_loop
	_ = condition_1
	var __t0 string
	{
		if condition_1 {
			__t0 = name_0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = ("Fail: ") + (name_0)
	}
end_branch_0:
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_131790935_3790796878(in *Constructor_Data_Eq_Eq[[]string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_131790935(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]string]{}
	out.V0 = in.V0
	return out
}
