package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Inject_dollar_Dict gopurs_runtime.Value
var once_Main_Inject_dollar_Dict sync.Once

func Get_Main_Inject_dollar_Dict() gopurs_runtime.Value {
	once_Main_Inject_dollar_Dict.Do(func() {
		cache_Main_Inject_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Inject_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Inject_dollar_Dict
}

var cache_Main_prj gopurs_runtime.Value
var once_Main_prj sync.Once

func Get_Main_prj() gopurs_runtime.Value {
	once_Main_prj.Do(func() {
		cache_Main_prj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_prj(gopurs_runtime.CoerceToStruct[Constructor_Main_Inject](dict_0_box))
		})
	})
	return cache_Main_prj
}

var cache_Main_injectRefl gopurs_runtime.Value
var once_Main_injectRefl sync.Once

func Get_Main_injectRefl() gopurs_runtime.Value {
	once_Main_injectRefl.Do(func() {
		cache_Main_injectRefl = gopurs_runtime.Value{Type: 9, IntVal: 930202529, UnsafePtr: unsafe.Pointer(&Constructor_Main_Inject{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_0})}
		})})}
	})
	return cache_Main_injectRefl
}

var cache_Main_injectLeft gopurs_runtime.Value
var once_Main_injectLeft sync.Once

func Get_Main_injectLeft() gopurs_runtime.Value {
	once_Main_injectLeft.Do(func() {
		cache_Main_injectLeft = gopurs_runtime.Value{Type: 9, IntVal: 930202529, UnsafePtr: unsafe.Pointer(&Constructor_Main_Inject{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_0})}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 *Constructor_Data_Maybe_Just
			{
				if v_0.Type == 9 && v_0.IntVal == 3711209382 {
					__t0 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Either_Left)(v_0.UnsafePtr).V0}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (*Constructor_Data_Maybe_Just)(nil)
			}
		end_branch_0:
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
		})})}
	})
	return cache_Main_injectLeft
}

var cache_Main_inj gopurs_runtime.Value
var once_Main_inj sync.Once

func Get_Main_inj() gopurs_runtime.Value {
	once_Main_inj.Do(func() {
		cache_Main_inj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inj(gopurs_runtime.CoerceToStruct[Constructor_Main_Inject](dict_0_box))
		})
	})
	return cache_Main_inj
}

var cache_Main_injL gopurs_runtime.Value
var once_Main_injL sync.Once

func Get_Main_injL() gopurs_runtime.Value {
	once_Main_injL.Do(func() {
		cache_Main_injL = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_injL(x_0_box)
		})
	})
	return cache_Main_injL
}

var cache_Main_injectRight gopurs_runtime.Value
var once_Main_injectRight sync.Once

func Get_Main_injectRight() gopurs_runtime.Value {
	once_Main_injectRight.Do(func() {
		cache_Main_injectRight = gopurs_runtime.Func(func(dictInject_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_injectRight(dictInject_0_box)
		})
	})
	return cache_Main_injectRight
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_inj__2246686172 gopurs_runtime.Value
var once_Main_inj__2246686172 sync.Once

func Get_Main_inj__2246686172() gopurs_runtime.Value {
	once_Main_inj__2246686172.Do(func() {
		cache_Main_inj__2246686172 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_inj__2246686172(gopurs_runtime.CoerceToStruct[Constructor_Main_Inject](dict_0_box))
		})
	})
	return cache_Main_inj__2246686172
}

var cache_Main_injectLeft__4013048200 gopurs_runtime.Value
var once_Main_injectLeft__4013048200 sync.Once

func Get_Main_injectLeft__4013048200() gopurs_runtime.Value {
	once_Main_injectLeft__4013048200.Do(func() {
		cache_Main_injectLeft__4013048200 = gopurs_runtime.Value{Type: 9, IntVal: 930202529, UnsafePtr: unsafe.Pointer(&Constructor_Main_Inject{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_0})}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 *Constructor_Data_Maybe_Just
			{
				if v_0.Type == 9 && v_0.IntVal == 3711209382 {
					__t0 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Either_Left)(v_0.UnsafePtr).V0}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (*Constructor_Data_Maybe_Just)(nil)
			}
		end_branch_0:
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
		})})}
	})
	return cache_Main_injectLeft__4013048200
}

var cache_Main_prj__1257140861 gopurs_runtime.Value
var once_Main_prj__1257140861 sync.Once

func Get_Main_prj__1257140861() gopurs_runtime.Value {
	once_Main_prj__1257140861.Do(func() {
		cache_Main_prj__1257140861 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_prj__1257140861(gopurs_runtime.CoerceToStruct[Constructor_Main_Inject](dict_0_box))
		})
	})
	return cache_Main_prj__1257140861
}

type Constructor_Main_Inject struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[930202529] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Inject)(ptr)
		_ = c
		switch key {
		case "inj":
			return gopurs_runtime.Box(c.V0)
		case "prj":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Inject: " + key)
		}
	}
}

func Call_Main_Inject_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_prj(dict_0_loop *Constructor_Main_Inject) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Inject = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_inj(dict_0_loop *Constructor_Main_Inject) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Inject = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_injL(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_0})}
}

func Call_Main_injectRight(dictInject_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictInject_0 gopurs_runtime.Value = dictInject_0_loop
	_ = dictInject_0
	return gopurs_runtime.Value{Type: 9, IntVal: 930202529, UnsafePtr: unsafe.Pointer(&Constructor_Main_Inject{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "inj"), x_1)})}
	}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 *Constructor_Data_Maybe_Just
		{
			if v_1.Type == 9 && v_1.IntVal == 2465973597 {
				__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "prj"), (*Constructor_Data_Either_Right)(v_1.UnsafePtr).V0))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (*Constructor_Data_Maybe_Just)(nil)
		}
	end_branch_0:
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
	})})}
}

func Call_Main_inj__2246686172(dict_0_loop *Constructor_Main_Inject) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Inject = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_prj__1257140861(dict_0_loop *Constructor_Main_Inject) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Inject = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}
