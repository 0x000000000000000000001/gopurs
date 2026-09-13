package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Balanced_dollar_Dict gopurs_runtime.Value
var once_Main_Balanced_dollar_Dict sync.Once

func Get_Main_Balanced_dollar_Dict() gopurs_runtime.Value {
	once_Main_Balanced_dollar_Dict.Do(func() {
		cache_Main_Balanced_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Balanced_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_Balanced_dollar_Dict
}

var cache_Main_Balanced_dollar_Dict__591156948 gopurs_runtime.Value
var once_Main_Balanced_dollar_Dict__591156948 sync.Once

func Get_Main_Balanced_dollar_Dict__591156948() gopurs_runtime.Value {
	once_Main_Balanced_dollar_Dict__591156948.Do(func() {
		cache_Main_Balanced_dollar_Dict__591156948 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_Balanced_dollar_Dict__591156948(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_Balanced_dollar_Dict__591156948
}

var cache_Main_balanced2 gopurs_runtime.Value
var once_Main_balanced2 sync.Once

func Get_Main_balanced2() gopurs_runtime.Value {
	once_Main_balanced2.Do(func() {
		cache_Main_balanced2 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_balanced2(_dollar___unused_0_box, _dollar___unused_1_box, _dollar___unused_2_box)
		})
	})
	return cache_Main_balanced2
}

var cache_Main_balanced1 gopurs_runtime.Value
var once_Main_balanced1 sync.Once

func Get_Main_balanced1() gopurs_runtime.Value {
	once_Main_balanced1.Do(func() {
		cache_Main_balanced1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_balanced1
}

var cache_Main_balanced gopurs_runtime.Value
var once_Main_balanced sync.Once

func Get_Main_balanced() gopurs_runtime.Value {
	once_Main_balanced.Do(func() {
		cache_Main_balanced = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal)))
		})
	})
	return cache_Main_balanced
}

var cache_Main_balanced__2739615157 gopurs_runtime.Value
var once_Main_balanced__2739615157 sync.Once

func Get_Main_balanced__2739615157() gopurs_runtime.Value {
	once_Main_balanced__2739615157.Do(func() {
		cache_Main_balanced__2739615157 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced__2739615157(uint32(v_unused_0_box.IntVal)))
		})
	})
	return cache_Main_balanced__2739615157
}

var cache_Main_balanced__312822740 gopurs_runtime.Value
var once_Main_balanced__312822740 sync.Once

func Get_Main_balanced__312822740() gopurs_runtime.Value {
	once_Main_balanced__312822740.Do(func() {
		cache_Main_balanced__312822740 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced__312822740(uint32(v_unused_0_box.IntVal)))
		})
	})
	return cache_Main_balanced__312822740
}

var cache_Main_balanced__2179065685 gopurs_runtime.Value
var once_Main_balanced__2179065685 sync.Once

func Get_Main_balanced__2179065685() gopurs_runtime.Value {
	once_Main_balanced__2179065685.Do(func() {
		cache_Main_balanced__2179065685 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced__2179065685(uint32(v_unused_0_box.IntVal)))
		})
	})
	return cache_Main_balanced__2179065685
}

var cache_Main_balanced__2564430452 gopurs_runtime.Value
var once_Main_balanced__2564430452 sync.Once

func Get_Main_balanced__2564430452() gopurs_runtime.Value {
	once_Main_balanced__2564430452.Do(func() {
		cache_Main_balanced__2564430452 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_balanced__2564430452(uint32(v_unused_0_box.IntVal)))
		})
	})
	return cache_Main_balanced__2564430452
}

var cache_Main_b3 gopurs_runtime.Value
var once_Main_b3 sync.Once

func Get_Main_b3() gopurs_runtime.Value {
	once_Main_b3.Do(func() {
		cache_Main_b3 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b3
}

var cache_Main_b2 gopurs_runtime.Value
var once_Main_b2 sync.Once

func Get_Main_b2() gopurs_runtime.Value {
	once_Main_b2.Do(func() {
		cache_Main_b2 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b2
}

var cache_Main_b1 gopurs_runtime.Value
var once_Main_b1 sync.Once

func Get_Main_b1() gopurs_runtime.Value {
	once_Main_b1.Do(func() {
		cache_Main_b1 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b1
}

var cache_Main_b0 gopurs_runtime.Value
var once_Main_b0 sync.Once

func Get_Main_b0() gopurs_runtime.Value {
	once_Main_b0.Do(func() {
		cache_Main_b0 = gopurs_runtime.Str("ok")
	})
	return cache_Main_b0
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("ok")), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Balanced[T_sym any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2947706876] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Balanced[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Balanced: " + key)
		}
	}
}

func Call_Main_Balanced_dollar_Dict(x_0_loop struct {
}) uint32 {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return uint32(func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}().IntVal)
}

func Call_Main_Balanced_dollar_Dict__591156948(x_0_loop struct {
}) uint32 {
Balanced_dollar_Dict__591156948:
	for {
		if false {
			continue Balanced_dollar_Dict__591156948
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return uint32(func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)
	}
}

func Call_Main_balanced2(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
		orig := struct {
		}{}
		_ = orig
		return gopurs_runtime.RecordDict0()
	}().IntVal)), UnsafePtr: nil}
}

func Call_Main_balanced(_dollar___unused_0_loop uint32, v_1_loop uint32) string {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return "ok"
}

func Call_Main_balanced__2739615157(v_unused_0_loop uint32) string {
balanced__2739615157:
	for {
		if false {
			continue balanced__2739615157
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return "ok"
	}
}

func Call_Main_balanced__312822740(v_unused_0_loop uint32) string {
balanced__312822740:
	for {
		if false {
			continue balanced__312822740
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return "ok"
	}
}

func Call_Main_balanced__2179065685(v_unused_0_loop uint32) string {
balanced__2179065685:
	for {
		if false {
			continue balanced__2179065685
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return "ok"
	}
}

func Call_Main_balanced__2564430452(v_unused_0_loop uint32) string {
balanced__2564430452:
	for {
		if false {
			continue balanced__2564430452
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return "ok"
	}
}
