package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Stream gopurs_runtime.Value
var once_Main_Stream sync.Once

func Get_Main_Stream() gopurs_runtime.Value {
	once_Main_Stream.Do(func() {
		cache_Main_Stream = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Stream
}

var cache_Main_IsStream_dollar_Dict gopurs_runtime.Value
var once_Main_IsStream_dollar_Dict sync.Once

func Get_Main_IsStream_dollar_Dict() gopurs_runtime.Value {
	once_Main_IsStream_dollar_Dict.Do(func() {
		cache_Main_IsStream_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 668736856, UnsafePtr: unsafe.Pointer(Call_Main_IsStream_dollar_Dict(func() struct {
				cons   gopurs_runtime.Value
				uncons gopurs_runtime.Value
				/*gopurs:record-v1:2:0063006f006e0073,0075006e0063006f006e0073*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					cons   gopurs_runtime.Value
					uncons gopurs_runtime.Value
					/*gopurs:record-v1:2:0063006f006e0073,0075006e0063006f006e0073*/
				}{}
				clone.cons = gopurs_runtime.RecordGet(orig, "cons")
				clone.uncons = gopurs_runtime.RecordGet(orig, "uncons")
				return clone
			}()))}
		})
	})
	return cache_Main_IsStream_dollar_Dict
}

var cache_Main_uncons gopurs_runtime.Value
var once_Main_uncons sync.Once

func Get_Main_uncons() gopurs_runtime.Value {
	once_Main_uncons.Do(func() {
		cache_Main_uncons = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_uncons(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_uncons
}

var cache_Main_streamIsStream gopurs_runtime.Value
var once_Main_streamIsStream sync.Once

func Get_Main_streamIsStream() gopurs_runtime.Value {
	once_Main_streamIsStream.Do(func() {
		cache_Main_streamIsStream = gopurs_runtime.Value{Type: 9, IntVal: 668736856, UnsafePtr: unsafe.Pointer(Rebox_Main_506827641_2356788234((&Constructor_Main_IsStream[gopurs_runtime.Value, *Constructor_Main_Stream[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer((&Constructor_Main_Stream[gopurs_runtime.Value]{1, x_0, xs_1}))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Main_Stream[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Main_Stream[gopurs_runtime.Value])(v_0.UnsafePtr).V1, Get_Data_Unit_unit()))
		})})))}
	})
	return cache_Main_streamIsStream
}

var cache_Main_uncons__3690923209 gopurs_runtime.Value
var once_Main_uncons__3690923209 sync.Once

func Get_Main_uncons__3690923209() gopurs_runtime.Value {
	once_Main_uncons__3690923209.Do(func() {
		cache_Main_uncons__3690923209 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_uncons__3690923209(Rebox_Main_356935894_759323721(gopurs_runtime.CoerceToStruct[Constructor_Main_Stream[gopurs_runtime.Value]](__eta_norm_0_0_box)))
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(Rebox_Main_759323721_356935894(orig.tail))})
			}()
		})
	})
	return cache_Main_uncons__3690923209
}

var cache_Main_cons gopurs_runtime.Value
var once_Main_cons sync.Once

func Get_Main_cons() gopurs_runtime.Value {
	once_Main_cons.Do(func() {
		cache_Main_cons = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_cons
}

var cache_Main_cons__2394070454 gopurs_runtime.Value
var once_Main_cons__2394070454 sync.Once

func Get_Main_cons__2394070454() gopurs_runtime.Value {
	once_Main_cons__2394070454.Do(func() {
		cache_Main_cons__2394070454 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(Rebox_Main_759323721_356935894(Call_Main_cons__2394070454(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box)))}
		})
	})
	return cache_Main_cons__2394070454
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func2(func(dictIsStream_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value]](dictIsStream_0_box), s_1_box)
		})
	})
	return cache_Main_test
}

var cache_Main_test__483983286 gopurs_runtime.Value
var once_Main_test__483983286 sync.Once

func Get_Main_test__483983286() gopurs_runtime.Value {
	once_Main_test__483983286.Do(func() {
		cache_Main_test__483983286 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(Rebox_Main_759323721_356935894(Call_Main_test__483983286(Rebox_Main_356935894_759323721(gopurs_runtime.CoerceToStruct[Constructor_Main_Stream[gopurs_runtime.Value]](s_0_box)))))}
		})
	})
	return cache_Main_test__483983286
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var dones_0_0_0 *Constructor_Main_Stream[string]
			_ = dones_0_0_0
			var dones_0_0_0_cell **Constructor_Main_Stream[string]
			_ = dones_0_0_0_cell
			// FALLBACK TCO: isLoop=false len=1
			dones_0_0_0 = Rebox_Main_356935894_759323721((&Constructor_Main_Stream[gopurs_runtime.Value]{1, gopurs_runtime.Str("Done"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(Rebox_Main_759323721_356935894((*dones_0_0_0_cell)))}
			})}))
			dones_0_0_0_cell = &dones_0_0_0
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str((Call_Main_test__483983286(dones_0_0_0)).V0).StrVal()))
		}()
	})
	return cache_Main_main
}

type Constructor_Main_Stream[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}

type Constructor_Main_IsStream[T_el any, T_s any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0049007300530074007200650061006d
func init() {
	gopurs_runtime.StructGetters[668736856] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "cons":
			return gopurs_runtime.Box(c.V0)
		case "uncons":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_IsStream: " + key)
		}
	}
}

func Call_Main_IsStream_dollar_Dict(x_0_loop struct {
	cons   gopurs_runtime.Value
	uncons gopurs_runtime.Value
	/*gopurs:record-v1:2:0063006f006e0073,0075006e0063006f006e0073*/
}) *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		cons   gopurs_runtime.Value
		uncons gopurs_runtime.Value
		/*gopurs:record-v1:2:0063006f006e0073,0075006e0063006f006e0073*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("cons", "uncons", orig.cons, orig.uncons)
	}())
}

func Call_Main_uncons(dict_0_loop *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_uncons__3690923209(__eta_norm_0_0_loop *Constructor_Main_Stream[string]) struct {
	head string
	tail *Constructor_Main_Stream[string]
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
} {
uncons__3690923209:
	for {
		if false {
			continue uncons__3690923209
		}
		var __eta_norm_0_0 *Constructor_Main_Stream[string] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return func(record struct {
			head gopurs_runtime.Value
			tail *Constructor_Main_Stream[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}) struct {
			head string
			tail *Constructor_Main_Stream[string]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		} {
			return struct {
				head string
				tail *Constructor_Main_Stream[string]
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}{record.head.StrVal(), Rebox_Main_356935894_759323721(record.tail)}
		}(struct {
			head gopurs_runtime.Value
			tail *Constructor_Main_Stream[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{gopurs_runtime.Str((__eta_norm_0_0).V0), gopurs_runtime.CoerceToStruct[Constructor_Main_Stream[gopurs_runtime.Value]](gopurs_runtime.Apply((__eta_norm_0_0).V1, Get_Data_Unit_unit()))})
	}
}

func Call_Main_cons(dict_0_loop *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_cons__2394070454(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) *Constructor_Main_Stream[string] {
cons__2394070454:
	for {
		if false {
			continue cons__2394070454
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Rebox_Main_356935894_759323721((&Constructor_Main_Stream[gopurs_runtime.Value]{1, gopurs_runtime.Str(__eta_norm_1_0), __eta_norm_0_1}))
	}
}

func Call_Main_test(dictIsStream_0_loop *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value], s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsStream_0 *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value] = dictIsStream_0_loop
	_ = dictIsStream_0
	var s_1 gopurs_runtime.Value = s_1_loop
	_ = s_1
	// TAST (Let): v_2_0 shape=App(Other) bindingType=(Record (Row [head: Any, tail: (TypeVar s$scope12)] Empty))
	v_2_0 := func() struct {
		head gopurs_runtime.Value
		tail gopurs_runtime.Value
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := gopurs_runtime.Apply(dictIsStream_0.V1, s_1)
		_ = orig
		clone := struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head")
		clone.tail = gopurs_runtime.RecordGet(orig, "tail")
		return clone
	}()
	_ = v_2_0
	// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
	__local_var_3_1 := v_2_0.tail
	_ = __local_var_3_1
	return gopurs_runtime.Apply2(dictIsStream_0.V0, v_2_0.head, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return __local_var_3_1
	}))
}

func Call_Main_test__483983286(s_0_loop *Constructor_Main_Stream[string]) *Constructor_Main_Stream[string] {
test__483983286:
	for {
		if false {
			continue test__483983286
		}
		var s_0 *Constructor_Main_Stream[string] = s_0_loop
		_ = s_0
		// TAST (Let): v_1_0 shape=LitRecord bindingType=(Record (Row [head: Any, tail: (TypeVar s$scope12)] Empty))
		v_1_0 := func(record struct {
			head gopurs_runtime.Value
			tail *Constructor_Main_Stream[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}) struct {
			head gopurs_runtime.Value
			tail gopurs_runtime.Value
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		} {
			return struct {
				head gopurs_runtime.Value
				tail gopurs_runtime.Value
				/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
			}{record.head, gopurs_runtime.Value{Type: 9, IntVal: 1020906690, UnsafePtr: unsafe.Pointer(record.tail)}}
		}(struct {
			head gopurs_runtime.Value
			tail *Constructor_Main_Stream[gopurs_runtime.Value]
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{gopurs_runtime.Str((s_0).V0), gopurs_runtime.CoerceToStruct[Constructor_Main_Stream[gopurs_runtime.Value]](gopurs_runtime.Apply((s_0).V1, Get_Data_Unit_unit()))})
		_ = v_1_0
		// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
		__local_var_2_1 := v_1_0.tail
		_ = __local_var_2_1
		return Rebox_Main_356935894_759323721((&Constructor_Main_Stream[gopurs_runtime.Value]{1, v_1_0.head, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return __local_var_2_1
		})}))
	}
}

func Rebox_Main_356935894_759323721(in *Constructor_Main_Stream[gopurs_runtime.Value]) *Constructor_Main_Stream[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Stream[string]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1
	return out
}

func Rebox_Main_506827641_2356788234(in *Constructor_Main_IsStream[gopurs_runtime.Value, *Constructor_Main_Stream[gopurs_runtime.Value]]) *Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_IsStream[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_759323721_356935894(in *Constructor_Main_Stream[string]) *Constructor_Main_Stream[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Stream[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}
