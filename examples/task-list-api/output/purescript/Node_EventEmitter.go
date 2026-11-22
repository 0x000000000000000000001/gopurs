package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_EventEmitter_discard gopurs_runtime.Value
var once_Node_EventEmitter_discard sync.Once

func Get_Node_EventEmitter_discard() gopurs_runtime.Value {
	once_Node_EventEmitter_discard.Do(func() {
		cache_Node_EventEmitter_discard = Get_Control_Bind_discard__1714184905()
	})
	return cache_Node_EventEmitter_discard
}

var cache_Node_EventEmitter_pure gopurs_runtime.Value
var once_Node_EventEmitter_pure sync.Once

func Get_Node_EventEmitter_pure() gopurs_runtime.Value {
	once_Node_EventEmitter_pure.Do(func() {
		cache_Node_EventEmitter_pure = Get_Effect_pureE()
	})
	return cache_Node_EventEmitter_pure
}

var cache_Node_EventEmitter_go__map gopurs_runtime.Value
var once_Node_EventEmitter_go__map sync.Once

func Get_Node_EventEmitter_go__map() gopurs_runtime.Value {
	once_Node_EventEmitter_go__map.Do(func() {
		cache_Node_EventEmitter_go__map = Get_Data_Functor_map__4008632948()
	})
	return cache_Node_EventEmitter_go__map
}

var cache_Node_EventEmitter_EventHandle gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle sync.Once

func Get_Node_EventEmitter_EventHandle() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle.Do(func() {
		cache_Node_EventEmitter_EventHandle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.StrVal(), value1}))}
			})
		})
	})
	return cache_Node_EventEmitter_EventHandle
}

var cache_Node_EventEmitter_EventHandle__1927327469 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__1927327469 sync.Once

func Get_Node_EventEmitter_EventHandle__1927327469() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__1927327469.Do(func() {
		cache_Node_EventEmitter_EventHandle__1927327469 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__1927327469(__eta_norm_1_0_box.StrVal(), __eta_norm_0_unused_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__1927327469
}

var cache_Node_EventEmitter_EventHandle__4236249709 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__4236249709 sync.Once

func Get_Node_EventEmitter_EventHandle__4236249709() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__4236249709.Do(func() {
		cache_Node_EventEmitter_EventHandle__4236249709 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_EventHandle__4236249709(__eta_norm_0_0_box.StrVal())
		})
	})
	return cache_Node_EventEmitter_EventHandle__4236249709
}

var cache_Node_EventEmitter_EventHandle__1332576874 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__1332576874 sync.Once

func Get_Node_EventEmitter_EventHandle__1332576874() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__1332576874.Do(func() {
		cache_Node_EventEmitter_EventHandle__1332576874 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__1332576874(__eta_norm_1_0_box.StrVal(), __eta_norm_0_unused_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__1332576874
}

var cache_Node_EventEmitter_EventHandle__1063446879 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__1063446879 sync.Once

func Get_Node_EventEmitter_EventHandle__1063446879() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__1063446879.Do(func() {
		cache_Node_EventEmitter_EventHandle__1063446879 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__1063446879(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__1063446879
}

var cache_Node_EventEmitter_EventHandle__1917583647 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__1917583647 sync.Once

func Get_Node_EventEmitter_EventHandle__1917583647() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__1917583647.Do(func() {
		cache_Node_EventEmitter_EventHandle__1917583647 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__1917583647(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__1917583647
}

var cache_Node_EventEmitter_EventHandle__1264074026 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__1264074026 sync.Once

func Get_Node_EventEmitter_EventHandle__1264074026() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__1264074026.Do(func() {
		cache_Node_EventEmitter_EventHandle__1264074026 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__1264074026(__eta_norm_1_0_box.StrVal(), __eta_norm_0_unused_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__1264074026
}

var cache_Node_EventEmitter_EventHandle__2860140895 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__2860140895 sync.Once

func Get_Node_EventEmitter_EventHandle__2860140895() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__2860140895.Do(func() {
		cache_Node_EventEmitter_EventHandle__2860140895 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__2860140895(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__2860140895
}

var cache_Node_EventEmitter_EventHandle__4174212650 gopurs_runtime.Value
var once_Node_EventEmitter_EventHandle__4174212650 sync.Once

func Get_Node_EventEmitter_EventHandle__4174212650() gopurs_runtime.Value {
	once_Node_EventEmitter_EventHandle__4174212650.Do(func() {
		cache_Node_EventEmitter_EventHandle__4174212650 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_EventEmitter_EventHandle__4174212650(__eta_norm_1_0_box.StrVal(), __eta_norm_0_unused_1_box))}
		})
	})
	return cache_Node_EventEmitter_EventHandle__4174212650
}

var cache_Node_EventEmitter_subscribeSameFunction gopurs_runtime.Value
var once_Node_EventEmitter_subscribeSameFunction sync.Once

func Get_Node_EventEmitter_subscribeSameFunction() gopurs_runtime.Value {
	once_Node_EventEmitter_subscribeSameFunction.Do(func() {
		cache_Node_EventEmitter_subscribeSameFunction = gopurs_runtime.Func4(func(onXFn_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_4_0 := gopurs_runtime.UncurriedApp3(onXFn_0, __local_var_1, __local_var_2, __local_var_3)
				_ = __local_var_4_0
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_unsafeOff(), __local_var_1, __local_var_2, __local_var_3)
				})
			}), gopurs_runtime.Value{})
		})
	})
	return cache_Node_EventEmitter_subscribeSameFunction
}

var cache_Node_EventEmitter_setMaxListeners gopurs_runtime.Value
var once_Node_EventEmitter_setMaxListeners sync.Once

func Get_Node_EventEmitter_setMaxListeners() gopurs_runtime.Value {
	once_Node_EventEmitter_setMaxListeners.Do(func() {
		cache_Node_EventEmitter_setMaxListeners = gopurs_runtime.Func2(func(max_0_box gopurs_runtime.Value, emitter_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_setMaxListeners(max_0_box.IntVal, emitter_1_box)
		})
	})
	return cache_Node_EventEmitter_setMaxListeners
}

var cache_Node_EventEmitter_setUnlimitedListeners gopurs_runtime.Value
var once_Node_EventEmitter_setUnlimitedListeners sync.Once

func Get_Node_EventEmitter_setUnlimitedListeners() gopurs_runtime.Value {
	once_Node_EventEmitter_setUnlimitedListeners.Do(func() {
		cache_Node_EventEmitter_setUnlimitedListeners = gopurs_runtime.Apply(Get_Node_EventEmitter_setMaxListeners(), gopurs_runtime.Int(int64(0)))
	})
	return cache_Node_EventEmitter_setUnlimitedListeners
}

var cache_Node_EventEmitter_removeListenerH gopurs_runtime.Value
var once_Node_EventEmitter_removeListenerH sync.Once

func Get_Node_EventEmitter_removeListenerH() gopurs_runtime.Value {
	once_Node_EventEmitter_removeListenerH.Do(func() {
		cache_Node_EventEmitter_removeListenerH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "removeListener", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(jsSymbol_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_0, gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_symbolOrStr(), Get_Data_Either_Left(), Get_Data_Either_Right(), jsSymbol_1)), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_EventEmitter_removeListenerH
}

var cache_Node_EventEmitter_prependOnceListener_ gopurs_runtime.Value
var once_Node_EventEmitter_prependOnceListener_ sync.Once

func Get_Node_EventEmitter_prependOnceListener_() gopurs_runtime.Value {
	once_Node_EventEmitter_prependOnceListener_.Do(func() {
		cache_Node_EventEmitter_prependOnceListener_ = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_prependOnceListener_(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_prependOnceListener_
}

var cache_Node_EventEmitter_prependOnceListener gopurs_runtime.Value
var once_Node_EventEmitter_prependOnceListener sync.Once

func Get_Node_EventEmitter_prependOnceListener() gopurs_runtime.Value {
	once_Node_EventEmitter_prependOnceListener.Do(func() {
		cache_Node_EventEmitter_prependOnceListener = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_prependOnceListener(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_prependOnceListener
}

var cache_Node_EventEmitter_prependListener_ gopurs_runtime.Value
var once_Node_EventEmitter_prependListener_ sync.Once

func Get_Node_EventEmitter_prependListener_() gopurs_runtime.Value {
	once_Node_EventEmitter_prependListener_.Do(func() {
		cache_Node_EventEmitter_prependListener_ = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_prependListener_(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_prependListener_
}

var cache_Node_EventEmitter_prependListener gopurs_runtime.Value
var once_Node_EventEmitter_prependListener sync.Once

func Get_Node_EventEmitter_prependListener() gopurs_runtime.Value {
	once_Node_EventEmitter_prependListener.Do(func() {
		cache_Node_EventEmitter_prependListener = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_prependListener(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_prependListener
}

var cache_Node_EventEmitter_once_ gopurs_runtime.Value
var once_Node_EventEmitter_once_ sync.Once

func Get_Node_EventEmitter_once_() gopurs_runtime.Value {
	once_Node_EventEmitter_once_.Do(func() {
		cache_Node_EventEmitter_once_ = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_once_(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_once_
}

var cache_Node_EventEmitter_once gopurs_runtime.Value
var once_Node_EventEmitter_once sync.Once

func Get_Node_EventEmitter_once() gopurs_runtime.Value {
	once_Node_EventEmitter_once.Do(func() {
		cache_Node_EventEmitter_once = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_once(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_once
}

var cache_Node_EventEmitter_on_ gopurs_runtime.Value
var once_Node_EventEmitter_on_ sync.Once

func Get_Node_EventEmitter_on_() gopurs_runtime.Value {
	once_Node_EventEmitter_on_.Do(func() {
		cache_Node_EventEmitter_on_ = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_on_(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_on_
}

var cache_Node_EventEmitter_on gopurs_runtime.Value
var once_Node_EventEmitter_on sync.Once

func Get_Node_EventEmitter_on() gopurs_runtime.Value {
	once_Node_EventEmitter_on.Do(func() {
		cache_Node_EventEmitter_on = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, psCb_1_box gopurs_runtime.Value, eventEmitter_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_on(gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box), psCb_1_box, eventEmitter_2_box)
		})
	})
	return cache_Node_EventEmitter_on
}

var cache_Node_EventEmitter_newListenerH gopurs_runtime.Value
var once_Node_EventEmitter_newListenerH sync.Once

func Get_Node_EventEmitter_newListenerH() gopurs_runtime.Value {
	once_Node_EventEmitter_newListenerH.Do(func() {
		cache_Node_EventEmitter_newListenerH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "newListener", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(jsSymbol_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_0, gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_symbolOrStr(), Get_Data_Either_Left(), Get_Data_Either_Right(), jsSymbol_1)), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_EventEmitter_newListenerH
}

var cache_Node_EventEmitter_listenerCount gopurs_runtime.Value
var once_Node_EventEmitter_listenerCount sync.Once

func Get_Node_EventEmitter_listenerCount() gopurs_runtime.Value {
	once_Node_EventEmitter_listenerCount.Do(func() {
		cache_Node_EventEmitter_listenerCount = gopurs_runtime.Func2(func(emitter_0_box gopurs_runtime.Value, eventName_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_EventEmitter_listenerCount(emitter_0_box, eventName_1_box.StrVal())
		})
	})
	return cache_Node_EventEmitter_listenerCount
}

var cache_Node_EventEmitter_getMaxListeners gopurs_runtime.Value
var once_Node_EventEmitter_getMaxListeners sync.Once

func Get_Node_EventEmitter_getMaxListeners() gopurs_runtime.Value {
	once_Node_EventEmitter_getMaxListeners.Do(func() {
		cache_Node_EventEmitter_getMaxListeners = gopurs_runtime.Apply(Get_Effect_Uncurried_runEffectFn1(), Get_Node_EventEmitter_getMaxListenersImpl())
	})
	return cache_Node_EventEmitter_getMaxListeners
}

var cache_Node_EventEmitter_eventNames gopurs_runtime.Value
var once_Node_EventEmitter_eventNames sync.Once

func Get_Node_EventEmitter_eventNames() gopurs_runtime.Value {
	once_Node_EventEmitter_eventNames.Do(func() {
		cache_Node_EventEmitter_eventNames = gopurs_runtime.Func(func(ee_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Node_EventEmitter_eventNames(ee_0_box))
		})
	})
	return cache_Node_EventEmitter_eventNames
}

type Constructor_Node_EventEmitter_EventHandle[T_emitterType any, T_pureScriptCallback any, T_javaScriptCallback any] struct {
	Rc uint32
	V0 string
	V1 gopurs_runtime.Value
}

func Call_Node_EventEmitter_EventHandle__1927327469(__eta_norm_1_0_loop string, __eta_norm_0_unused_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__1927327469:
	for {
		if false {
			continue EventHandle__1927327469
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, Get_Node_Process_identity()})
	}
}

func Call_Node_EventEmitter_EventHandle__4236249709(__eta_norm_0_0_loop string) gopurs_runtime.Value {
EventHandle__4236249709:
	for {
		if false {
			continue EventHandle__4236249709
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.StrVal(), value1}))}
			})
		}), gopurs_runtime.Str(__eta_norm_0_0))
	}
}

func Call_Node_EventEmitter_EventHandle__1332576874(__eta_norm_1_0_loop string, __eta_norm_0_unused_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__1332576874:
	for {
		if false {
			continue EventHandle__1332576874
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, Get_Effect_Uncurried_mkEffectFn1()})
	}
}

func Call_Node_EventEmitter_EventHandle__1063446879(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__1063446879:
	for {
		if false {
			continue EventHandle__1063446879
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Node_EventEmitter_EventHandle__1917583647(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__1917583647:
	for {
		if false {
			continue EventHandle__1917583647
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Node_EventEmitter_EventHandle__1264074026(__eta_norm_1_0_loop string, __eta_norm_0_unused_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__1264074026:
	for {
		if false {
			continue EventHandle__1264074026
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, Get_Effect_Uncurried_mkEffectFn1()})
	}
}

func Call_Node_EventEmitter_EventHandle__2860140895(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__2860140895:
	for {
		if false {
			continue EventHandle__2860140895
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Node_EventEmitter_EventHandle__4174212650(__eta_norm_1_0_loop string, __eta_norm_0_unused_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
EventHandle__4174212650:
	for {
		if false {
			continue EventHandle__4174212650
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta_norm_1_0, Get_Effect_Uncurried_mkEffectFn1()})
	}
}

func Call_Node_EventEmitter_setMaxListeners(max_0_loop int64, emitter_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var max_0 int64 = max_0_loop
	_ = max_0
	var emitter_1 gopurs_runtime.Value = emitter_1_loop
	_ = emitter_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_EventEmitter_setMaxListenersImpl(), emitter_1, gopurs_runtime.Int(max_0))
	})
}

func Call_Node_EventEmitter_prependOnceListener_(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_unsafePrependOnceListener(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_prependOnceListener(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp4(Get_Node_EventEmitter_subscribeSameFunction(), Get_Node_EventEmitter_unsafePrependOnceListener(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_prependListener_(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_unsafePrependListener(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_prependListener(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp4(Get_Node_EventEmitter_subscribeSameFunction(), Get_Node_EventEmitter_unsafePrependListener(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_once_(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_unsafeOnce(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_once(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp4(Get_Node_EventEmitter_subscribeSameFunction(), Get_Node_EventEmitter_unsafeOnce(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_on_(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_unsafeOn(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_on(v_0_loop *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], psCb_1_loop gopurs_runtime.Value, eventEmitter_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
	_ = v_0
	var psCb_1 gopurs_runtime.Value = psCb_1_loop
	_ = psCb_1
	var eventEmitter_2 gopurs_runtime.Value = eventEmitter_2_loop
	_ = eventEmitter_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
		__local_var_3_0 := (v_0).V0
		_ = __local_var_3_0
		// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
		__local_var_4_1 := gopurs_runtime.Apply((v_0).V1, psCb_1)
		_ = __local_var_4_1
		return gopurs_runtime.UncurriedApp4(Get_Node_EventEmitter_subscribeSameFunction(), Get_Node_EventEmitter_unsafeOn(), eventEmitter_2, gopurs_runtime.Str(__local_var_3_0), __local_var_4_1)
	})
}

func Call_Node_EventEmitter_listenerCount(emitter_0_loop gopurs_runtime.Value, eventName_1_loop string) gopurs_runtime.Value {
	var emitter_0 gopurs_runtime.Value = emitter_0_loop
	_ = emitter_0
	var eventName_1 string = eventName_1_loop
	_ = eventName_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_EventEmitter_listenerCountImpl(), emitter_0, gopurs_runtime.Str(eventName_1))
	})
}

func Call_Node_EventEmitter_eventNames(ee_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var ee_0 gopurs_runtime.Value = ee_0_loop
	_ = ee_0
	arrayIntrinsic_arg_0 := func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp3(Get_Node_EventEmitter_symbolOrStr(), Get_Data_Either_Left(), Get_Data_Either_Right(), x_1)
	}
	_ = arrayIntrinsic_arg_0
	_ = arrayIntrinsic_arg_0
	arrayIntrinsic_arg_1 := gopurs_runtime.Apply(Get_Node_EventEmitter_eventNamesImpl(), ee_0)
	_ = arrayIntrinsic_arg_1
	_ = arrayIntrinsic_arg_1
	return func() []gopurs_runtime.Value {
		arr_val_arrayMap0 := arrayIntrinsic_arg_1
		_ = arr_val_arrayMap0
		_ = arr_val_arrayMap0
		arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
		_ = arr_go_arrayMap0
		res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
		_ = res_go_arrayMap0
		for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
			res_go_arrayMap0[i_arrayMap0] = arrayIntrinsic_arg_0(v_arrayMap0)
		}
		return res_go_arrayMap0
	}()
}

func Get_Node_EventEmitter_eventNamesImpl() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_EventNamesImpl
}

func Get_Node_EventEmitter_getMaxListenersImpl() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_GetMaxListenersImpl
}

func Get_Node_EventEmitter_listenerCountImpl() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_ListenerCountImpl
}

func Get_Node_EventEmitter_go__new() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_Go__new
}

func Get_Node_EventEmitter_setMaxListenersImpl() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_SetMaxListenersImpl
}

func Get_Node_EventEmitter_symbolOrStr() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_SymbolOrStr
}

func Get_Node_EventEmitter_unsafeEmitFn() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_UnsafeEmitFn
}

func Get_Node_EventEmitter_unsafeOff() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_UnsafeOff
}

func Get_Node_EventEmitter_unsafeOn() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_UnsafeOn
}

func Get_Node_EventEmitter_unsafeOnce() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_UnsafeOnce
}

func Get_Node_EventEmitter_unsafePrependListener() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_UnsafePrependListener
}

func Get_Node_EventEmitter_unsafePrependOnceListener() gopurs_runtime.Value {
	return _Gopurs_Node_EventEmitter_UnsafePrependOnceListener
}
