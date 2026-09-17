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

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1953100407_1386611502(Rebox_Main_1386611502_1953100407(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_Binding gopurs_runtime.Value
var once_Main_Binding sync.Once

func Get_Main_Binding() gopurs_runtime.Value {
	once_Main_Binding.Do(func() {
		cache_Main_Binding = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer((&Constructor_Main_Binding[gopurs_runtime.Value]{1, value0, value1.StrVal()}))}
			})
		})
	})
	return cache_Main_Binding
}

var cache_Main_Binding__267875142 gopurs_runtime.Value
var once_Main_Binding__267875142 sync.Once

func Get_Main_Binding__267875142() gopurs_runtime.Value {
	once_Main_Binding__267875142.Do(func() {
		cache_Main_Binding__267875142 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer(Rebox_Main_3242760310_4122662347(Call_Main_Binding__267875142(func() struct {
				label string
				line  int64
			} {
				orig := __eta_norm_1_0_box
				_ = orig
				clone := struct {
					label string
					line  int64
				}{}
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
				return clone
			}(), __eta_norm_0_1_box.StrVal())))}
		})
	})
	return cache_Main_Binding__267875142
}

var cache_Main_NonRec gopurs_runtime.Value
var once_Main_NonRec sync.Once

func Get_Main_NonRec() gopurs_runtime.Value {
	once_Main_NonRec.Do(func() {
		cache_Main_NonRec = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3480592549, UnsafePtr: unsafe.Pointer((&Constructor_Main_NonRec[gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](value0)}))}
		})
	})
	return cache_Main_NonRec
}

var cache_Main_NonRec__4285879642 gopurs_runtime.Value
var once_Main_NonRec__4285879642 sync.Once

func Get_Main_NonRec__4285879642() gopurs_runtime.Value {
	once_Main_NonRec__4285879642.Do(func() {
		cache_Main_NonRec__4285879642 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_NonRec__4285879642(Rebox_Main_4122662347_3242760310(gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_NonRec__4285879642
}

var cache_Main_Rec gopurs_runtime.Value
var once_Main_Rec sync.Once

func Get_Main_Rec() gopurs_runtime.Value {
	once_Main_Rec.Do(func() {
		cache_Main_Rec = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4196596074, UnsafePtr: unsafe.Pointer((&Constructor_Main_Rec[gopurs_runtime.Value]{1, func() []*Constructor_Main_Binding[gopurs_runtime.Value] {
				arr := *(*[]gopurs_runtime.Value)(value0.UnsafePtr)
				unboxed := make([]*Constructor_Main_Binding[gopurs_runtime.Value], len(arr))
				for i, v := range arr {
					unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](v)
				}
				return unboxed
			}()}))}
		})
	})
	return cache_Main_Rec
}

var cache_Main_Rec__1133290652 gopurs_runtime.Value
var once_Main_Rec__1133290652 sync.Once

func Get_Main_Rec__1133290652() gopurs_runtime.Value {
	once_Main_Rec__1133290652.Do(func() {
		cache_Main_Rec__1133290652 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Rec__1133290652(func() []*Constructor_Main_Binding[struct {
				label string
				line  int64
			}] {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)
				unboxed := make([]*Constructor_Main_Binding[struct {
					label string
					line  int64
				}], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Main_4122662347_3242760310(gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](v))
				}
				return unboxed
			}())
		})
	})
	return cache_Main_Rec__1133290652
}

var cache_Main_flatten gopurs_runtime.Value
var once_Main_flatten sync.Once

func Get_Main_flatten() gopurs_runtime.Value {
	once_Main_flatten.Do(func() {
		cache_Main_flatten = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Main_flatten(v_0_box)
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer(Rebox_Main_3242760310_4122662347(v))}
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Main_flatten
}

var cache_Main_describe gopurs_runtime.Value
var once_Main_describe sync.Once

func Get_Main_describe() gopurs_runtime.Value {
	once_Main_describe.Do(func() {
		cache_Main_describe = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_describe(Rebox_Main_4122662347_3242760310(gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Main_describe
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(ADT ["Main","Bind"] [(Record (Row [line: Int, label: String] Empty))])])])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 3480592549, UnsafePtr: unsafe.Pointer(Rebox_Main_863915020_2606613137((&Constructor_Main_NonRec[struct {
				label string
				line  int64
			}]{1, (&Constructor_Main_Binding[struct {
				label string
				line  int64
			}]{1, struct {
				label string
				line  int64
			}{"typed", int64(42)}, "x"})})))})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___2499841393("", struct {
				actual   []string
				expected []string
			}{func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr_val_arrayMap6 := func() gopurs_runtime.Value {
							arr := Call_Main_flatten(__local_var_2_2)
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer(Rebox_Main_3242760310_4122662347(v))}
							}
							return gopurs_runtime.Array(boxed)
						}()
						_ = arr_val_arrayMap6
						arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
						_ = arr_go_arrayMap6
						res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
						_ = res_go_arrayMap6
						for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
							res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(Get_Main_describe(), v_arrayMap6)
						}
						return gopurs_runtime.Array(res_go_arrayMap6)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}(), []string{"x:typed:42"}}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{gopurs_runtime.Int(int64(len(Call_Main_flatten(gopurs_runtime.Value{Type: 9, IntVal: 4196596074, UnsafePtr: unsafe.Pointer(Rebox_Main_2297685059_1811419774((&Constructor_Main_Rec[struct {
					label string
					line  int64
				}]{1, func() []*Constructor_Main_Binding[gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_Binding[struct {
							label string
							line  int64
						}]{}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer(Rebox_Main_3242760310_4122662347(v))}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]*Constructor_Main_Binding[gopurs_runtime.Value], len(arr))
					for i, v := range arr {
						unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](v)
					}
					return unboxed
				}()})))})))).IntVal, int64(0)}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
				}))
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Binding[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 string
}

type Constructor_Main_NonRec[T_a any] struct {
	Rc uint32
	V0 *Constructor_Main_Binding[T_a]
}

type Constructor_Main_Rec[T_a any] struct {
	Rc uint32
	V0 []*Constructor_Main_Binding[gopurs_runtime.Value]
}

func Call_Main_Binding__267875142(__eta_norm_1_0_loop struct {
	label string
	line  int64
}, __eta_norm_0_1_loop string) *Constructor_Main_Binding[struct {
	label string
	line  int64
}] {
Binding__267875142:
	for {
		if false {
			continue Binding__267875142
		}
		var __eta_norm_1_0 struct {
			label string
			line  int64
		} = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 string = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Binding[struct {
			label string
			line  int64
		}]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_NonRec__4285879642(__eta_norm_0_0_loop *Constructor_Main_Binding[struct {
	label string
	line  int64
}]) gopurs_runtime.Value {
NonRec__4285879642:
	for {
		if false {
			continue NonRec__4285879642
		}
		var __eta_norm_0_0 *Constructor_Main_Binding[struct {
			label string
			line  int64
		}] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 3480592549, UnsafePtr: unsafe.Pointer(Rebox_Main_863915020_2606613137((&Constructor_Main_NonRec[struct {
			label string
			line  int64
		}]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Rec__1133290652(__eta_norm_0_0_loop []*Constructor_Main_Binding[struct {
	label string
	line  int64
}]) gopurs_runtime.Value {
Rec__1133290652:
	for {
		if false {
			continue Rec__1133290652
		}
		var __eta_norm_0_0 []*Constructor_Main_Binding[struct {
			label string
			line  int64
		}] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 4196596074, UnsafePtr: unsafe.Pointer(Rebox_Main_2297685059_1811419774((&Constructor_Main_Rec[struct {
			label string
			line  int64
		}]{1, func() []*Constructor_Main_Binding[gopurs_runtime.Value] {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := __eta_norm_0_0
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer(Rebox_Main_3242760310_4122662347(v))}
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]*Constructor_Main_Binding[gopurs_runtime.Value], len(arr))
			for i, v := range arr {
				unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](v)
			}
			return unboxed
		}()})))}
	}
}

func Call_Main_flatten(v_0_loop gopurs_runtime.Value) []*Constructor_Main_Binding[struct {
	label string
	line  int64
}] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 []*Constructor_Main_Binding[struct {
		label string
		line  int64
	}]
	{
		if v_0.Type == 9 && v_0.IntVal == 3480592549 {
			__t0 = []*Constructor_Main_Binding[struct {
				label string
				line  int64
			}]{Rebox_Main_4122662347_3242760310((*Constructor_Main_NonRec[gopurs_runtime.Value])(v_0.UnsafePtr).V0)}
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 4196596074 {
			__t0 = func() []*Constructor_Main_Binding[struct {
				label string
				line  int64
			}] {
				arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := (*Constructor_Main_Rec[gopurs_runtime.Value])(v_0.UnsafePtr).V0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3177043775, UnsafePtr: unsafe.Pointer(v)}
					}
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
				unboxed := make([]*Constructor_Main_Binding[struct {
					label string
					line  int64
				}], len(arr))
				for i, v := range arr {
					unboxed[i] = Rebox_Main_4122662347_3242760310(gopurs_runtime.CoerceToStruct[Constructor_Main_Binding[gopurs_runtime.Value]](v))
				}
				return unboxed
			}()
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() []*Constructor_Main_Binding[struct {
			label string
			line  int64
		}] {
			panic("Failed pattern match")
		}()
	}
end_branch_0:
	return __t0
}

func Call_Main_describe(v_0_loop *Constructor_Main_Binding[struct {
	label string
	line  int64
}]) string {
	var v_0 *Constructor_Main_Binding[struct {
		label string
		line  int64
	}] = v_0_loop
	_ = v_0
	return (((((v_0).V1) + (":")) + ((v_0).V0.label)) + (":")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((v_0).V0.line)).StrVal())
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

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1953100407(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1953100407_1386611502(in *Constructor_Data_Show_Show[[]string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2297685059_1811419774(in *Constructor_Main_Rec[struct {
	label string
	line  int64
}]) *Constructor_Main_Rec[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Rec[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3242760310_4122662347(in *Constructor_Main_Binding[struct {
	label string
	line  int64
}]) *Constructor_Main_Binding[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Binding[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("label", "line", gopurs_runtime.Str(orig.label), gopurs_runtime.Int(orig.line))
	}()
	out.V1 = in.V1
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

func Rebox_Main_4122662347_3242760310(in *Constructor_Main_Binding[gopurs_runtime.Value]) *Constructor_Main_Binding[struct {
	label string
	line  int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Binding[struct {
		label string
		line  int64
	}]{}
	out.V0 = func() struct {
		label string
		line  int64
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			label string
			line  int64
		}{}
		clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
		clone.line = gopurs_runtime.RecordGet(orig, "line").IntVal
		return clone
	}()
	out.V1 = in.V1
	return out
}

func Rebox_Main_863915020_2606613137(in *Constructor_Main_NonRec[struct {
	label string
	line  int64
}]) *Constructor_Main_NonRec[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_NonRec[gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_3242760310_4122662347(in.V0)
	return out
}
