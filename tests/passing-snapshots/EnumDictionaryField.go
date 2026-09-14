package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Fast gopurs_runtime.Value
var once_Main_Fast sync.Once

func Get_Main_Fast() gopurs_runtime.Value {
	once_Main_Fast.Do(func() {
		cache_Main_Fast = gopurs_runtime.Value{Type: 9, IntVal: int64(4067170750), UnsafePtr: nil}
	})
	return cache_Main_Fast
}

var cache_Main_Safe gopurs_runtime.Value
var once_Main_Safe sync.Once

func Get_Main_Safe() gopurs_runtime.Value {
	once_Main_Safe.Do(func() {
		cache_Main_Safe = gopurs_runtime.Value{Type: 9, IntVal: int64(1097047119), UnsafePtr: nil}
	})
	return cache_Main_Safe
}

var cache_Main_Protected_dollar_Dict gopurs_runtime.Value
var once_Main_Protected_dollar_Dict sync.Once

func Get_Main_Protected_dollar_Dict() gopurs_runtime.Value {
	once_Main_Protected_dollar_Dict.Do(func() {
		cache_Main_Protected_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Call_Main_Protected_dollar_Dict(func() struct {
				priority uint32
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					priority uint32
				}{}
				clone.priority = uint32(gopurs_runtime.RecordGet(orig, "priority").IntVal)
				return clone
			}()))}
		})
	})
	return cache_Main_Protected_dollar_Dict
}

var cache_Main_Protected_dollar_Dict__3961917678 gopurs_runtime.Value
var once_Main_Protected_dollar_Dict__3961917678 sync.Once

func Get_Main_Protected_dollar_Dict__3961917678() gopurs_runtime.Value {
	once_Main_Protected_dollar_Dict__3961917678.Do(func() {
		cache_Main_Protected_dollar_Dict__3961917678 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_3555759771_48729792(Call_Main_Protected_dollar_Dict__3961917678(func() struct {
				priority uint32
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					priority uint32
				}{}
				clone.priority = uint32(gopurs_runtime.RecordGet(orig, "priority").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Main_Protected_dollar_Dict__3961917678
}

var cache_Main_Protected_dollar_Dict__62804680 gopurs_runtime.Value
var once_Main_Protected_dollar_Dict__62804680 sync.Once

func Get_Main_Protected_dollar_Dict__62804680() gopurs_runtime.Value {
	once_Main_Protected_dollar_Dict__62804680.Do(func() {
		cache_Main_Protected_dollar_Dict__62804680 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_2202314751_48729792(Call_Main_Protected_dollar_Dict__62804680(func() struct {
				priority uint32
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					priority uint32
				}{}
				clone.priority = uint32(gopurs_runtime.RecordGet(orig, "priority").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Main_Protected_dollar_Dict__62804680
}

var cache_Main_eqPriority gopurs_runtime.Value
var once_Main_eqPriority sync.Once

func Get_Main_eqPriority() gopurs_runtime.Value {
	once_Main_eqPriority.Do(func() {
		cache_Main_eqPriority = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 bool
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 4067170750 {
					var __t_tag_4 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_4
					__t5 = (uint32(__t_tag_4) == 4067170750)
					goto end_branch_5
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 1097047119 {

					var __t_tag_1 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 1097047119)
				}
				__t5 = __t_and_2
			}
		end_branch_5:
			return gopurs_runtime.Bool(__t5)
		})})))}
	})
	return cache_Main_eqPriority
}

var cache_Main_protectedString gopurs_runtime.Value
var once_Main_protectedString sync.Once

func Get_Main_protectedString() gopurs_runtime.Value {
	once_Main_protectedString.Do(func() {
		cache_Main_protectedString = gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_2202314751_48729792((&Constructor_Main_Protected[string]{1, 4067170750})))}
	})
	return cache_Main_protectedString
}

var cache_Main_protectedInt gopurs_runtime.Value
var once_Main_protectedInt sync.Once

func Get_Main_protectedInt() gopurs_runtime.Value {
	once_Main_protectedInt.Do(func() {
		cache_Main_protectedInt = gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_3555759771_48729792((&Constructor_Main_Protected[int64]{1, 1097047119})))}
	})
	return cache_Main_protectedInt
}

var cache_Main_priority gopurs_runtime.Value
var once_Main_priority sync.Once

func Get_Main_priority() gopurs_runtime.Value {
	once_Main_priority.Do(func() {
		cache_Main_priority = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_priority(dict_0_box)
		})
	})
	return cache_Main_priority
}

var cache_Main_priority__3373190335 gopurs_runtime.Value
var once_Main_priority__3373190335 sync.Once

func Get_Main_priority__3373190335() gopurs_runtime.Value {
	once_Main_priority__3373190335.Do(func() {
		cache_Main_priority__3373190335 = gopurs_runtime.Value{Type: 9, IntVal: int64(1097047119), UnsafePtr: nil}
	})
	return cache_Main_priority__3373190335
}

var cache_Main_priority__3023229849 gopurs_runtime.Value
var once_Main_priority__3023229849 sync.Once

func Get_Main_priority__3023229849() gopurs_runtime.Value {
	once_Main_priority__3023229849.Do(func() {
		cache_Main_priority__3023229849 = gopurs_runtime.Value{Type: 9, IntVal: int64(4067170750), UnsafePtr: nil}
	})
	return cache_Main_priority__3023229849
}

var cache_Main_priority__4145524652 gopurs_runtime.Value
var once_Main_priority__4145524652 sync.Once

func Get_Main_priority__4145524652() gopurs_runtime.Value {
	once_Main_priority__4145524652.Do(func() {
		cache_Main_priority__4145524652 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_priority__4145524652(dict_0_box)
		})
	})
	return cache_Main_priority__4145524652
}

var cache_Main_isSafe gopurs_runtime.Value
var once_Main_isSafe sync.Once

func Get_Main_isSafe() gopurs_runtime.Value {
	once_Main_isSafe.Do(func() {
		cache_Main_isSafe = gopurs_runtime.Func(func(dictProtected_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_isSafe(dictProtected_0_box)
		})
	})
	return cache_Main_isSafe
}

var cache_Main_isSafe__735247059 gopurs_runtime.Value
var once_Main_isSafe__735247059 sync.Once

func Get_Main_isSafe__735247059() gopurs_runtime.Value {
	once_Main_isSafe__735247059.Do(func() {
		cache_Main_isSafe__735247059 = func() gopurs_runtime.Value {
			var __t_tag_0 gopurs_runtime.Value = Call_Main_priority(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_3555759771_48729792(Rebox_Main_48729792_3555759771(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](Get_Main_protectedInt()))))})
			_ = __t_tag_0
			return gopurs_runtime.Bool((uint32(__t_tag_0.IntVal) == 1097047119))
		}()
	})
	return cache_Main_isSafe__735247059
}

var cache_Main_isSafe__916403509 gopurs_runtime.Value
var once_Main_isSafe__916403509 sync.Once

func Get_Main_isSafe__916403509() gopurs_runtime.Value {
	once_Main_isSafe__916403509.Do(func() {
		cache_Main_isSafe__916403509 = func() gopurs_runtime.Value {
			var __t_tag_0 gopurs_runtime.Value = Call_Main_priority(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_2202314751_48729792(Rebox_Main_48729792_2202314751(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](Get_Main_protectedString()))))})
			_ = __t_tag_0
			return gopurs_runtime.Bool((uint32(__t_tag_0.IntVal) == 1097047119))
		}()
	})
	return cache_Main_isSafe__916403509
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
			actual   bool
			expected bool
		}{(Call_Main_isSafe(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_3555759771_48729792(Rebox_Main_48729792_3555759771(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](Get_Main_protectedInt()))))}).IntVal) != (0), true}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(Call_Main_isSafe(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(Rebox_Main_2202314751_48729792(Rebox_Main_48729792_2202314751(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](Get_Main_protectedString()))))}).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Fast struct {
	Rc uint32
}

type Constructor_Main_Safe struct {
	Rc uint32
}

type Constructor_Main_Protected[T_a any] struct {
	Rc uint32
	V0 uint32
}

func init() {
	gopurs_runtime.StructGetters[2242194644] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Protected[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "priority":
			return gopurs_runtime.Value{Type: 9, IntVal: int64(c.V0)}
		default:
			panic("Key not found in dictionary Constructor_Main_Protected: " + key)
		}
	}
}

func Call_Main_Protected_dollar_Dict(x_0_loop struct {
	priority uint32
}) *Constructor_Main_Protected[gopurs_runtime.Value] {
	var x_0 struct {
		priority uint32
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("priority", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.priority), UnsafePtr: nil})
	}())
}

func Call_Main_Protected_dollar_Dict__3961917678(x_0_loop struct {
	priority uint32
}) *Constructor_Main_Protected[int64] {
Protected_dollar_Dict__3961917678:
	for {
		if false {
			continue Protected_dollar_Dict__3961917678
		}
		var x_0 struct {
			priority uint32
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[int64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("priority", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.priority), UnsafePtr: nil})
		}())
	}
}

func Call_Main_Protected_dollar_Dict__62804680(x_0_loop struct {
	priority uint32
}) *Constructor_Main_Protected[string] {
Protected_dollar_Dict__62804680:
	for {
		if false {
			continue Protected_dollar_Dict__62804680
		}
		var x_0 struct {
			priority uint32
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("priority", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.priority), UnsafePtr: nil})
		}())
	}
}

func Call_Main_priority(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(dict_0, "priority").IntVal)), UnsafePtr: nil}
}

func Call_Main_priority__4145524652(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
priority__4145524652:
	for {
		if false {
			continue priority__4145524652
		}
		var dict_0 gopurs_runtime.Value = dict_0_loop
		_ = dict_0
		return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(dict_0, "priority").IntVal)), UnsafePtr: nil}
	}
}

func Call_Main_isSafe(dictProtected_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProtected_0 gopurs_runtime.Value = dictProtected_0_loop
	_ = dictProtected_0
	var __t_tag_0 gopurs_runtime.Value = Call_Main_priority(gopurs_runtime.Value{Type: 9, IntVal: 2242194644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Protected[gopurs_runtime.Value]](dictProtected_0))})
	_ = __t_tag_0
	return gopurs_runtime.Bool((uint32(__t_tag_0.IntVal) == 1097047119))
}

func Rebox_Main_2202314751_48729792(in *Constructor_Main_Protected[string]) *Constructor_Main_Protected[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Protected[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3555759771_48729792(in *Constructor_Main_Protected[int64]) *Constructor_Main_Protected[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Protected[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_48729792_2202314751(in *Constructor_Main_Protected[gopurs_runtime.Value]) *Constructor_Main_Protected[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Protected[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_48729792_3555759771(in *Constructor_Main_Protected[gopurs_runtime.Value]) *Constructor_Main_Protected[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Protected[int64]{}
	out.V0 = in.V0
	return out
}
