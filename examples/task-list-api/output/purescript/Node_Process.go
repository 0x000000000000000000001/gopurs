package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Process_go__append gopurs_runtime.Value
var once_Node_Process_go__append sync.Once

func Get_Node_Process_go__append() gopurs_runtime.Value {
	once_Node_Process_go__append.Do(func() {
		cache_Node_Process_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Node_Process_go__append
}

var cache_Node_Process_show gopurs_runtime.Value
var once_Node_Process_show sync.Once

func Get_Node_Process_show() gopurs_runtime.Value {
	once_Node_Process_show.Do(func() {
		cache_Node_Process_show = gopurs_runtime.RecordGet(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("system", 6)
		})), Call_Data_Show_showRecordFieldsConsNil(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("user", 4)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_1636311157_1386611502(Rebox_Node_Process_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_1636311157_1386611502(Rebox_Node_Process_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), "show")
	})
	return cache_Node_Process_show
}

var cache_Node_Process_conj gopurs_runtime.Value
var once_Node_Process_conj sync.Once

func Get_Node_Process_conj() gopurs_runtime.Value {
	once_Node_Process_conj.Do(func() {
		cache_Node_Process_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Node_Process_conj
}

var cache_Node_Process_eq gopurs_runtime.Value
var once_Node_Process_eq sync.Once

func Get_Node_Process_eq() gopurs_runtime.Value {
	once_Node_Process_eq.Do(func() {
		cache_Node_Process_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Node_Process_eq
}

var cache_Node_Process_identity gopurs_runtime.Value
var once_Node_Process_identity sync.Once

func Get_Node_Process_identity() gopurs_runtime.Value {
	once_Node_Process_identity.Do(func() {
		cache_Node_Process_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_identity(x_0_box)
		})
	})
	return cache_Node_Process_identity
}

var cache_Node_Process_go__map gopurs_runtime.Value
var once_Node_Process_go__map sync.Once

func Get_Node_Process_go__map() gopurs_runtime.Value {
	once_Node_Process_go__map.Do(func() {
		cache_Node_Process_go__map = Get_Data_Functor_map__1260122395()
	})
	return cache_Node_Process_go__map
}

var cache_Node_Process_map1 gopurs_runtime.Value
var once_Node_Process_map1 sync.Once

func Get_Node_Process_map1() gopurs_runtime.Value {
	once_Node_Process_map1.Do(func() {
		cache_Node_Process_map1 = Get_Data_Functor_map__4150859291()
	})
	return cache_Node_Process_map1
}

var cache_Node_Process_CpuUsage gopurs_runtime.Value
var once_Node_Process_CpuUsage sync.Once

func Get_Node_Process_CpuUsage() gopurs_runtime.Value {
	once_Node_Process_CpuUsage.Do(func() {
		cache_Node_Process_CpuUsage = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Node_Process_CpuUsage(func() struct {
					system int64
					user   int64
					/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						system int64
						user   int64
						/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
					}{}
					clone.system = gopurs_runtime.RecordGet(orig, "system").IntVal
					clone.user = gopurs_runtime.RecordGet(orig, "user").IntVal
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("system", "user", gopurs_runtime.Int(orig.system), gopurs_runtime.Int(orig.user))
			}()
		})
	})
	return cache_Node_Process_CpuUsage
}

var cache_Node_Process_showCpuUsage gopurs_runtime.Value
var once_Node_Process_showCpuUsage sync.Once

func Get_Node_Process_showCpuUsage() gopurs_runtime.Value {
	once_Node_Process_showCpuUsage.Do(func() {
		cache_Node_Process_showCpuUsage = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_611482891_1386611502((&Constructor_Data_Show_Show[struct {
			system int64
			user   int64
			/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
		}]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(CpuUsage ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.StrWithCodeUnits("system", 6)
			})), Call_Data_Show_showRecordFieldsConsNil(gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.StrWithCodeUnits("user", 4)
			})), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_1636311157_1386611502(Rebox_Node_Process_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_1636311157_1386611502(Rebox_Node_Process_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})), "show"), func() gopurs_runtime.Value {
				orig := func() struct {
					system int64
					user   int64
					/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
				} {
					orig := v_0
					_ = orig
					clone := struct {
						system int64
						user   int64
						/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
					}{}
					clone.system = gopurs_runtime.RecordGet(orig, "system").IntVal
					clone.user = gopurs_runtime.RecordGet(orig, "user").IntVal
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("system", "user", gopurs_runtime.Int(orig.system), gopurs_runtime.Int(orig.user))
			}()).StrVal()), ")"))
		})})))}
	})
	return cache_Node_Process_showCpuUsage
}

var cache_Node_Process_eqCpuUsage gopurs_runtime.Value
var once_Node_Process_eqCpuUsage sync.Once

func Get_Node_Process_eqCpuUsage() gopurs_runtime.Value {
	once_Node_Process_eqCpuUsage.Do(func() {
		cache_Node_Process_eqCpuUsage = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_328390891_3790796878((&Constructor_Data_Eq_Eq[struct {
			system int64
			user   int64
			/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
		}]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(x_0, "system").IntVal) == (gopurs_runtime.RecordGet(y_1, "system").IntVal)) && ((gopurs_runtime.RecordGet(x_0, "user").IntVal) == (gopurs_runtime.RecordGet(y_1, "user").IntVal)))
		})})))}
	})
	return cache_Node_Process_eqCpuUsage
}

var cache_Node_Process_workerH gopurs_runtime.Value
var once_Node_Process_workerH sync.Once

func Get_Node_Process_workerH() gopurs_runtime.Value {
	once_Node_Process_workerH.Do(func() {
		cache_Node_Process_workerH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "worker", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Process_workerH
}

var cache_Node_Process_warningH gopurs_runtime.Value
var once_Node_Process_warningH sync.Once

func Get_Node_Process_warningH() gopurs_runtime.Value {
	once_Node_Process_warningH.Do(func() {
		cache_Node_Process_warningH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "warning", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Process_warningH
}

var cache_Node_Process_unsetEnv gopurs_runtime.Value
var once_Node_Process_unsetEnv sync.Once

func Get_Node_Process_unsetEnv() gopurs_runtime.Value {
	once_Node_Process_unsetEnv.Do(func() {
		cache_Node_Process_unsetEnv = gopurs_runtime.Func(func(key_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_unsetEnv(key_0_box.StrVal())
		})
	})
	return cache_Node_Process_unsetEnv
}

var cache_Node_Process_unsafeSendOptsCb gopurs_runtime.Value
var once_Node_Process_unsafeSendOptsCb sync.Once

func Get_Node_Process_unsafeSendOptsCb() gopurs_runtime.Value {
	once_Node_Process_unsafeSendOptsCb.Do(func() {
		cache_Node_Process_unsafeSendOptsCb = gopurs_runtime.Func5(func(_dollar___unused_0_box gopurs_runtime.Value, msg_1_box gopurs_runtime.Value, handle_2_box gopurs_runtime.Value, opts_3_box gopurs_runtime.Value, cb_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_unsafeSendOptsCb(_dollar___unused_0_box, msg_1_box, handle_2_box, opts_3_box, cb_4_box)
		})
	})
	return cache_Node_Process_unsafeSendOptsCb
}

var cache_Node_Process_unsafeSendOpts gopurs_runtime.Value
var once_Node_Process_unsafeSendOpts sync.Once

func Get_Node_Process_unsafeSendOpts() gopurs_runtime.Value {
	once_Node_Process_unsafeSendOpts.Do(func() {
		cache_Node_Process_unsafeSendOpts = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, msg_1_box gopurs_runtime.Value, handle_2_box gopurs_runtime.Value, opts_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_unsafeSendOpts(_dollar___unused_0_box, msg_1_box, handle_2_box, opts_3_box)
		})
	})
	return cache_Node_Process_unsafeSendOpts
}

var cache_Node_Process_unsafeSendCb gopurs_runtime.Value
var once_Node_Process_unsafeSendCb sync.Once

func Get_Node_Process_unsafeSendCb() gopurs_runtime.Value {
	once_Node_Process_unsafeSendCb.Do(func() {
		cache_Node_Process_unsafeSendCb = gopurs_runtime.Func3(func(msg_0_box gopurs_runtime.Value, handle_1_box gopurs_runtime.Value, cb_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_unsafeSendCb(msg_0_box, handle_1_box, cb_2_box)
		})
	})
	return cache_Node_Process_unsafeSendCb
}

var cache_Node_Process_unsafeSend gopurs_runtime.Value
var once_Node_Process_unsafeSend sync.Once

func Get_Node_Process_unsafeSend() gopurs_runtime.Value {
	once_Node_Process_unsafeSend.Do(func() {
		cache_Node_Process_unsafeSend = gopurs_runtime.Func2(func(msg_0_box gopurs_runtime.Value, handle_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_unsafeSend(msg_0_box, handle_1_box)
		})
	})
	return cache_Node_Process_unsafeSend
}

var cache_Node_Process_unhandledRejectionH gopurs_runtime.Value
var once_Node_Process_unhandledRejectionH sync.Once

func Get_Node_Process_unhandledRejectionH() gopurs_runtime.Value {
	once_Node_Process_unhandledRejectionH.Do(func() {
		cache_Node_Process_unhandledRejectionH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "unhandledRejection", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(cb_0, a_1, __local_var_2), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_Process_unhandledRejectionH
}

var cache_Node_Process_uncaughtExceptionH gopurs_runtime.Value
var once_Node_Process_uncaughtExceptionH sync.Once

func Get_Node_Process_uncaughtExceptionH() gopurs_runtime.Value {
	once_Node_Process_uncaughtExceptionH.Do(func() {
		cache_Node_Process_uncaughtExceptionH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "uncaughtException", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(cb_0, a_1, gopurs_runtime.NormalizeStringValue(__local_var_2)), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_Process_uncaughtExceptionH
}

var cache_Node_Process_setUncaughtExceptionCaptureCallback gopurs_runtime.Value
var once_Node_Process_setUncaughtExceptionCaptureCallback sync.Once

func Get_Node_Process_setUncaughtExceptionCaptureCallback() gopurs_runtime.Value {
	once_Node_Process_setUncaughtExceptionCaptureCallback.Do(func() {
		cache_Node_Process_setUncaughtExceptionCaptureCallback = gopurs_runtime.Func(func(cb_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_setUncaughtExceptionCaptureCallback(cb_0_box)
		})
	})
	return cache_Node_Process_setUncaughtExceptionCaptureCallback
}

var cache_Node_Process_setTitle gopurs_runtime.Value
var once_Node_Process_setTitle sync.Once

func Get_Node_Process_setTitle() gopurs_runtime.Value {
	once_Node_Process_setTitle.Do(func() {
		cache_Node_Process_setTitle = gopurs_runtime.Func(func(newTitle_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_setTitle(newTitle_0_box.StrVal())
		})
	})
	return cache_Node_Process_setTitle
}

var cache_Node_Process_setExitCode gopurs_runtime.Value
var once_Node_Process_setExitCode sync.Once

func Get_Node_Process_setExitCode() gopurs_runtime.Value {
	once_Node_Process_setExitCode.Do(func() {
		cache_Node_Process_setExitCode = gopurs_runtime.Func(func(code_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_setExitCode(code_0_box.IntVal)
		})
	})
	return cache_Node_Process_setExitCode
}

var cache_Node_Process_setEnv gopurs_runtime.Value
var once_Node_Process_setEnv sync.Once

func Get_Node_Process_setEnv() gopurs_runtime.Value {
	once_Node_Process_setEnv.Do(func() {
		cache_Node_Process_setEnv = gopurs_runtime.Func2(func(key_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_setEnv(key_0_box.StrVal(), value_1_box.StrVal())
		})
	})
	return cache_Node_Process_setEnv
}

var cache_Node_Process_rejectionHandledH gopurs_runtime.Value
var once_Node_Process_rejectionHandledH sync.Once

func Get_Node_Process_rejectionHandledH() gopurs_runtime.Value {
	once_Node_Process_rejectionHandledH.Do(func() {
		cache_Node_Process_rejectionHandledH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "rejectionHandled", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Process_rejectionHandledH
}

var cache_Node_Process_platform gopurs_runtime.Value
var once_Node_Process_platform sync.Once

func Get_Node_Process_platform() gopurs_runtime.Value {
	once_Node_Process_platform.Do(func() {
		cache_Node_Process_platform = func() gopurs_runtime.Value {
			var __t0 *Constructor_Data_Maybe_Just[uint32]
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("aix") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(700987045), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("darwin") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(207948338), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("freebsd") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(3079948500), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("linux") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(4148875795), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("openbsd") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(508021620), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("sunos") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(2031616513), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("win32") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(569341156), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				if (Get_Node_Process_platformStr().StrVal()) == ("android") {
					__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 bool
						}{gopurs_runtime.Value{Type: 9, IntVal: int64(2339342990), UnsafePtr: nil}, true}
						if _v.V1 {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
						}
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
					}()))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = Rebox_Node_Process_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 bool
					}{gopurs_runtime.Value{}, false}
					if _v.V1 {
						return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
				}()))
			}
		end_branch_0:
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Node_Process_622082505_3094389156(__t0))}
		}()
	})
	return cache_Node_Process_platform
}

var cache_Node_Process_nextTick_prime_ gopurs_runtime.Value
var once_Node_Process_nextTick_prime_ sync.Once

func Get_Node_Process_nextTick_prime_() gopurs_runtime.Value {
	once_Node_Process_nextTick_prime_.Do(func() {
		cache_Node_Process_nextTick_prime_ = gopurs_runtime.Func2(func(cb_0_box gopurs_runtime.Value, args_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_nextTick_prime_(cb_0_box, args_1_box)
		})
	})
	return cache_Node_Process_nextTick_prime_
}

var cache_Node_Process_nextTick gopurs_runtime.Value
var once_Node_Process_nextTick sync.Once

func Get_Node_Process_nextTick() gopurs_runtime.Value {
	once_Node_Process_nextTick.Do(func() {
		cache_Node_Process_nextTick = gopurs_runtime.Func(func(cb_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_nextTick(cb_0_box)
		})
	})
	return cache_Node_Process_nextTick
}

var cache_Node_Process_mkSignalH_prime_ gopurs_runtime.Value
var once_Node_Process_mkSignalH_prime_ sync.Once

func Get_Node_Process_mkSignalH_prime_() gopurs_runtime.Value {
	once_Node_Process_mkSignalH_prime_.Do(func() {
		cache_Node_Process_mkSignalH_prime_ = gopurs_runtime.Func(func(sig_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_Process_mkSignalH_prime_(sig_0_box.StrVal()))}
		})
	})
	return cache_Node_Process_mkSignalH_prime_
}

var cache_Node_Process_mkSignalH gopurs_runtime.Value
var once_Node_Process_mkSignalH sync.Once

func Get_Node_Process_mkSignalH() gopurs_runtime.Value {
	once_Node_Process_mkSignalH.Do(func() {
		cache_Node_Process_mkSignalH = gopurs_runtime.Func(func(sig_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Node_Process_mkSignalH(uint32(sig_0_box.IntVal)))}
		})
	})
	return cache_Node_Process_mkSignalH
}

var cache_Node_Process_messageH gopurs_runtime.Value
var once_Node_Process_messageH sync.Once

func Get_Node_Process_messageH() gopurs_runtime.Value {
	once_Node_Process_messageH.Do(func() {
		cache_Node_Process_messageH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "message", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(cb_0, a_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), __local_var_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_Process_messageH
}

var cache_Node_Process_lookupEnv gopurs_runtime.Value
var once_Node_Process_lookupEnv sync.Once

func Get_Node_Process_lookupEnv() gopurs_runtime.Value {
	once_Node_Process_lookupEnv.Do(func() {
		cache_Node_Process_lookupEnv = gopurs_runtime.Func(func(k_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_lookupEnv(k_0_box.StrVal())
		})
	})
	return cache_Node_Process_lookupEnv
}

var cache_Node_Process_killStr gopurs_runtime.Value
var once_Node_Process_killStr sync.Once

func Get_Node_Process_killStr() gopurs_runtime.Value {
	once_Node_Process_killStr.Do(func() {
		cache_Node_Process_killStr = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, sig_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_killStr(p_0_box.IntVal, sig_1_box.StrVal())
		})
	})
	return cache_Node_Process_killStr
}

var cache_Node_Process_killInt gopurs_runtime.Value
var once_Node_Process_killInt sync.Once

func Get_Node_Process_killInt() gopurs_runtime.Value {
	once_Node_Process_killInt.Do(func() {
		cache_Node_Process_killInt = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, sig_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_killInt(p_0_box.IntVal, sig_1_box.IntVal)
		})
	})
	return cache_Node_Process_killInt
}

var cache_Node_Process_kill gopurs_runtime.Value
var once_Node_Process_kill sync.Once

func Get_Node_Process_kill() gopurs_runtime.Value {
	once_Node_Process_kill.Do(func() {
		cache_Node_Process_kill = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_kill(p_0_box.IntVal)
		})
	})
	return cache_Node_Process_kill
}

var cache_Node_Process_getUid gopurs_runtime.Value
var once_Node_Process_getUid sync.Once

func Get_Node_Process_getUid() gopurs_runtime.Value {
	once_Node_Process_getUid.Do(func() {
		cache_Node_Process_getUid = Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Data_Nullable_toMaybe(), Get_Node_Process_getUidImpl())
	})
	return cache_Node_Process_getUid
}

var cache_Node_Process_getGid gopurs_runtime.Value
var once_Node_Process_getGid sync.Once

func Get_Node_Process_getGid() gopurs_runtime.Value {
	once_Node_Process_getGid.Do(func() {
		cache_Node_Process_getGid = Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Data_Nullable_toMaybe(), Get_Node_Process_getGidImpl())
	})
	return cache_Node_Process_getGid
}

var cache_Node_Process_getExitCode gopurs_runtime.Value
var once_Node_Process_getExitCode sync.Once

func Get_Node_Process_getExitCode() gopurs_runtime.Value {
	once_Node_Process_getExitCode.Do(func() {
		cache_Node_Process_getExitCode = Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Data_Nullable_toMaybe(), Get_Node_Process_getExitCodeImpl())
	})
	return cache_Node_Process_getExitCode
}

var cache_Node_Process_exitH gopurs_runtime.Value
var once_Node_Process_exitH sync.Once

func Get_Node_Process_exitH() gopurs_runtime.Value {
	once_Node_Process_exitH.Do(func() {
		cache_Node_Process_exitH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "exit", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Process_exitH
}

var cache_Node_Process_exit_prime_ gopurs_runtime.Value
var once_Node_Process_exit_prime_ sync.Once

func Get_Node_Process_exit_prime_() gopurs_runtime.Value {
	once_Node_Process_exit_prime_.Do(func() {
		cache_Node_Process_exit_prime_ = gopurs_runtime.Func(func(code_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_exit_prime_(code_0_box.IntVal)
		})
	})
	return cache_Node_Process_exit_prime_
}

var cache_Node_Process_exit_prime___3720979425 gopurs_runtime.Value
var once_Node_Process_exit_prime___3720979425 sync.Once

func Get_Node_Process_exit_prime___3720979425() gopurs_runtime.Value {
	once_Node_Process_exit_prime___3720979425.Do(func() {
		cache_Node_Process_exit_prime___3720979425 = gopurs_runtime.Func(func(code_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_exit_prime___3720979425(code_0_box.IntVal)
		})
	})
	return cache_Node_Process_exit_prime___3720979425
}

var cache_Node_Process_disconnectH gopurs_runtime.Value
var once_Node_Process_disconnectH sync.Once

func Get_Node_Process_disconnectH() gopurs_runtime.Value {
	once_Node_Process_disconnectH.Do(func() {
		cache_Node_Process_disconnectH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "disconnect", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Process_disconnectH
}

var cache_Node_Process_disconnect gopurs_runtime.Value
var once_Node_Process_disconnect sync.Once

func Get_Node_Process_disconnect() gopurs_runtime.Value {
	once_Node_Process_disconnect.Do(func() {
		cache_Node_Process_disconnect = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_disconnectImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}
	})
	return cache_Node_Process_disconnect
}

var cache_Node_Process_cpuUsageToRecord gopurs_runtime.Value
var once_Node_Process_cpuUsageToRecord sync.Once

func Get_Node_Process_cpuUsageToRecord() gopurs_runtime.Value {
	once_Node_Process_cpuUsageToRecord.Do(func() {
		cache_Node_Process_cpuUsageToRecord = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Node_Process_cpuUsageToRecord(func() struct {
					system int64
					user   int64
					/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
				} {
					orig := v_0_box
					_ = orig
					clone := struct {
						system int64
						user   int64
						/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
					}{}
					clone.system = gopurs_runtime.RecordGet(orig, "system").IntVal
					clone.user = gopurs_runtime.RecordGet(orig, "user").IntVal
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("system", "user", gopurs_runtime.Int(orig.system), gopurs_runtime.Int(orig.user))
			}()
		})
	})
	return cache_Node_Process_cpuUsageToRecord
}

var cache_Node_Process_cpuUsageDiff gopurs_runtime.Value
var once_Node_Process_cpuUsageDiff sync.Once

func Get_Node_Process_cpuUsageDiff() gopurs_runtime.Value {
	once_Node_Process_cpuUsageDiff.Do(func() {
		cache_Node_Process_cpuUsageDiff = gopurs_runtime.Func(func(prev_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_cpuUsageDiff(func() struct {
				system int64
				user   int64
				/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
			} {
				orig := prev_0_box
				_ = orig
				clone := struct {
					system int64
					user   int64
					/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
				}{}
				clone.system = gopurs_runtime.RecordGet(orig, "system").IntVal
				clone.user = gopurs_runtime.RecordGet(orig, "user").IntVal
				return clone
			}())
		})
	})
	return cache_Node_Process_cpuUsageDiff
}

var cache_Node_Process_chdir gopurs_runtime.Value
var once_Node_Process_chdir sync.Once

func Get_Node_Process_chdir() gopurs_runtime.Value {
	once_Node_Process_chdir.Do(func() {
		cache_Node_Process_chdir = gopurs_runtime.Func(func(dir_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Process_chdir(dir_0_box.StrVal())
		})
	})
	return cache_Node_Process_chdir
}

var cache_Node_Process_channelUnref gopurs_runtime.Value
var once_Node_Process_channelUnref sync.Once

func Get_Node_Process_channelUnref() gopurs_runtime.Value {
	once_Node_Process_channelUnref.Do(func() {
		cache_Node_Process_channelUnref = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_channelUnrefImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}
	})
	return cache_Node_Process_channelUnref
}

var cache_Node_Process_channelRef gopurs_runtime.Value
var once_Node_Process_channelRef sync.Once

func Get_Node_Process_channelRef() gopurs_runtime.Value {
	once_Node_Process_channelRef.Do(func() {
		cache_Node_Process_channelRef = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_channelRefImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}
	})
	return cache_Node_Process_channelRef
}

var cache_Node_Process_beforeExitH gopurs_runtime.Value
var once_Node_Process_beforeExitH sync.Once

func Get_Node_Process_beforeExitH() gopurs_runtime.Value {
	once_Node_Process_beforeExitH.Do(func() {
		cache_Node_Process_beforeExitH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "beforeExit", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Process_beforeExitH
}

var cache_Node_Process_abort gopurs_runtime.Value
var once_Node_Process_abort sync.Once

func Get_Node_Process_abort() gopurs_runtime.Value {
	once_Node_Process_abort.Do(func() {
		cache_Node_Process_abort = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), Get_Node_Process_abortImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}
	})
	return cache_Node_Process_abort
}

func Call_Node_Process_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Node_Process_CpuUsage(x_0_loop struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
}) struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
} {
	var x_0 struct {
		system int64
		user   int64
		/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
	} = x_0_loop
	_ = x_0
	return x_0
}

func Call_Node_Process_unsetEnv(key_0_loop string) gopurs_runtime.Value {
	var key_0 string = key_0_loop
	_ = key_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_unsetEnvImpl(), gopurs_runtime.Str(key_0))
	})
}

func Call_Node_Process_unsafeSendOptsCb(_dollar___unused_0_loop gopurs_runtime.Value, msg_1_loop gopurs_runtime.Value, handle_2_loop gopurs_runtime.Value, opts_3_loop gopurs_runtime.Value, cb_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var msg_1 gopurs_runtime.Value = msg_1_loop
	_ = msg_1
	var handle_2 gopurs_runtime.Value = handle_2_loop
	_ = handle_2
	var opts_3 gopurs_runtime.Value = opts_3_loop
	_ = opts_3
	var cb_4 gopurs_runtime.Value = cb_4_loop
	_ = cb_4
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp4(Get_Node_Process_sendOptsCbImpl(), msg_1, handle_2, opts_3, gopurs_runtime.Func(func(err_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Process_unsafeSendOpts(_dollar___unused_0_loop gopurs_runtime.Value, msg_1_loop gopurs_runtime.Value, handle_2_loop gopurs_runtime.Value, opts_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var msg_1 gopurs_runtime.Value = msg_1_loop
	_ = msg_1
	var handle_2 gopurs_runtime.Value = handle_2_loop
	_ = handle_2
	var opts_3 gopurs_runtime.Value = opts_3_loop
	_ = opts_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp3(Get_Node_Process_sendOptsImpl(), msg_1, handle_2, opts_3)
	})
}

func Call_Node_Process_unsafeSendCb(msg_0_loop gopurs_runtime.Value, handle_1_loop gopurs_runtime.Value, cb_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var msg_0 gopurs_runtime.Value = msg_0_loop
	_ = msg_0
	var handle_1 gopurs_runtime.Value = handle_1_loop
	_ = handle_1
	var cb_2 gopurs_runtime.Value = cb_2_loop
	_ = cb_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp3(Get_Node_Process_sendCbImpl(), msg_0, handle_1, gopurs_runtime.Func(func(err_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Process_unsafeSend(msg_0_loop gopurs_runtime.Value, handle_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var msg_0 gopurs_runtime.Value = msg_0_loop
	_ = msg_0
	var handle_1 gopurs_runtime.Value = handle_1_loop
	_ = handle_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Process_sendImpl(), msg_0, handle_1)
	})
}

func Call_Node_Process_setUncaughtExceptionCaptureCallback(cb_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var cb_0 gopurs_runtime.Value = cb_0_loop
	_ = cb_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_setUncaughtExceptionCaptureCallbackImpl(), cb_0)
	})
}

func Call_Node_Process_setTitle(newTitle_0_loop string) gopurs_runtime.Value {
	var newTitle_0 string = newTitle_0_loop
	_ = newTitle_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_setTitleImpl(), gopurs_runtime.Str(newTitle_0))
	})
}

func Call_Node_Process_setExitCode(code_0_loop int64) gopurs_runtime.Value {
	var code_0 int64 = code_0_loop
	_ = code_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_setExitCodeImpl(), gopurs_runtime.Int(code_0))
	})
}

func Call_Node_Process_setEnv(key_0_loop string, value_1_loop string) gopurs_runtime.Value {
	var key_0 string = key_0_loop
	_ = key_0
	var value_1 string = value_1_loop
	_ = value_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Process_setEnvImpl(), gopurs_runtime.Str(key_0), gopurs_runtime.Str(value_1))
	})
}

func Call_Node_Process_nextTick_prime_(cb_0_loop gopurs_runtime.Value, args_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var cb_0 gopurs_runtime.Value = cb_0_loop
	_ = cb_0
	var args_1 gopurs_runtime.Value = args_1_loop
	_ = args_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Process_nextTickCbImpl(), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_0, __local_var_2), gopurs_runtime.Value{})
		}), args_1)
	})
}

func Call_Node_Process_nextTick(cb_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var cb_0 gopurs_runtime.Value = cb_0_loop
	_ = cb_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_nextTickImpl(), cb_0)
	})
}

func Call_Node_Process_mkSignalH_prime_(sig_0_loop string) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var sig_0 string = sig_0_loop
	_ = sig_0
	return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, Data_String_Common_ToUpper(sig_0), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	})})
}

func Call_Node_Process_mkSignalH(sig_0_loop uint32) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var sig_0 uint32 = sig_0_loop
	_ = sig_0
	return (&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_Data_Posix_Signal_toString(sig_0), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	})})
}

func Call_Node_Process_lookupEnv(k_0_loop string) gopurs_runtime.Value {
	var k_0 string = k_0_loop
	_ = k_0
	return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Apply(Get_Foreign_Object_lookup__2830175038(), gopurs_runtime.Str(k_0)), Get_Node_Process_unsafeGetEnv())
}

func Call_Node_Process_killStr(p_0_loop int64, sig_1_loop string) gopurs_runtime.Value {
	var p_0 int64 = p_0_loop
	_ = p_0
	var sig_1 string = sig_1_loop
	_ = sig_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Process_killStrImpl(), gopurs_runtime.Int(p_0), gopurs_runtime.Str(sig_1))
	})
}

func Call_Node_Process_killInt(p_0_loop int64, sig_1_loop int64) gopurs_runtime.Value {
	var p_0 int64 = p_0_loop
	_ = p_0
	var sig_1 int64 = sig_1_loop
	_ = sig_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Process_killIntImpl(), gopurs_runtime.Int(p_0), gopurs_runtime.Int(sig_1))
	})
}

func Call_Node_Process_kill(p_0_loop int64) gopurs_runtime.Value {
	var p_0 int64 = p_0_loop
	_ = p_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_killImpl(), gopurs_runtime.Int(p_0))
	})
}

func Call_Node_Process_exit_prime_(code_0_loop int64) gopurs_runtime.Value {
	var code_0 int64 = code_0_loop
	_ = code_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_exitImpl(), gopurs_runtime.Int(code_0))
	})
}

func Call_Node_Process_exit_prime___3720979425(code_0_loop int64) gopurs_runtime.Value {
exit_prime___3720979425:
	for {
		if false {
			continue exit_prime___3720979425
		}
		var code_0 int64 = code_0_loop
		_ = code_0
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.UncurriedApp(Get_Node_Process_exitImpl(), gopurs_runtime.Int(code_0))
		})
	}
}

func Call_Node_Process_cpuUsageToRecord(v_0_loop struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
}) struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
} {
	var v_0 struct {
		system int64
		user   int64
		/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
	} = v_0_loop
	_ = v_0
	return v_0
}

func Call_Node_Process_cpuUsageDiff(prev_0_loop struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
}) gopurs_runtime.Value {
	var prev_0 struct {
		system int64
		user   int64
		/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
	} = prev_0_loop
	_ = prev_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_cpuUsageDiffImpl(), func() gopurs_runtime.Value {
			orig := prev_0
			_ = orig
			return gopurs_runtime.RecordDict2("system", "user", gopurs_runtime.Int(orig.system), gopurs_runtime.Int(orig.user))
		}())
	})
}

func Call_Node_Process_chdir(dir_0_loop string) gopurs_runtime.Value {
	var dir_0 string = dir_0_loop
	_ = dir_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Process_chdirImpl(), gopurs_runtime.Str(dir_0))
	})
}

func Rebox_Node_Process_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Node_Process_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Node_Process_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[uint32]{}
	out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Node_Process_328390891_3790796878(in *Constructor_Data_Eq_Eq[struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Node_Process_611482891_1386611502(in *Constructor_Data_Show_Show[struct {
	system int64
	user   int64
	/*gopurs:record-v1:2:00730079007300740065006d,0075007300650072*/
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Node_Process_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Get_Node_Process_abortImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_AbortImpl
}

func Get_Node_Process_argv() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Argv
}

func Get_Node_Process_argv0() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Argv0
}

func Get_Node_Process_channelRefImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ChannelRefImpl
}

func Get_Node_Process_channelUnrefImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ChannelUnrefImpl
}

func Get_Node_Process_chdirImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ChdirImpl
}

func Get_Node_Process_clearUncaughtExceptionCaptureCallback() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ClearUncaughtExceptionCaptureCallback
}

func Get_Node_Process_config() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Config
}

func Get_Node_Process_connected() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Connected
}

func Get_Node_Process_cpuUsage() gopurs_runtime.Value {
	return _Gopurs_Node_Process_CpuUsage
}

func Get_Node_Process_cpuUsageDiffImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_CpuUsageDiffImpl
}

func Get_Node_Process_cwd() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Cwd
}

func Get_Node_Process_debugPort() gopurs_runtime.Value {
	return _Gopurs_Node_Process_DebugPort
}

func Get_Node_Process_disconnectImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_DisconnectImpl
}

func Get_Node_Process_execArgv() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ExecArgv
}

func Get_Node_Process_execPath() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ExecPath
}

func Get_Node_Process_exit() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Exit
}

func Get_Node_Process_exitImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ExitImpl
}

func Get_Node_Process_getEnv() gopurs_runtime.Value {
	return _Gopurs_Node_Process_GetEnv
}

func Get_Node_Process_getExitCodeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_GetExitCodeImpl
}

func Get_Node_Process_getGidImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_GetGidImpl
}

func Get_Node_Process_getTitle() gopurs_runtime.Value {
	return _Gopurs_Node_Process_GetTitle
}

func Get_Node_Process_getUidImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_GetUidImpl
}

func Get_Node_Process_hasUncaughtExceptionCaptureCallback() gopurs_runtime.Value {
	return _Gopurs_Node_Process_HasUncaughtExceptionCaptureCallback
}

func Get_Node_Process_killImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_KillImpl
}

func Get_Node_Process_killIntImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_KillIntImpl
}

func Get_Node_Process_killStrImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_KillStrImpl
}

func Get_Node_Process_memoryUsage() gopurs_runtime.Value {
	return _Gopurs_Node_Process_MemoryUsage
}

func Get_Node_Process_memoryUsageRss() gopurs_runtime.Value {
	return _Gopurs_Node_Process_MemoryUsageRss
}

func Get_Node_Process_nextTickCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_NextTickCbImpl
}

func Get_Node_Process_nextTickImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_NextTickImpl
}

func Get_Node_Process_pid() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Pid
}

func Get_Node_Process_platformStr() gopurs_runtime.Value {
	return _Gopurs_Node_Process_PlatformStr
}

func Get_Node_Process_ppid() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Ppid
}

func Get_Node_Process_process() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Process
}

func Get_Node_Process_resourceUsage() gopurs_runtime.Value {
	return _Gopurs_Node_Process_ResourceUsage
}

func Get_Node_Process_sendCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SendCbImpl
}

func Get_Node_Process_sendImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SendImpl
}

func Get_Node_Process_sendOptsCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SendOptsCbImpl
}

func Get_Node_Process_sendOptsImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SendOptsImpl
}

func Get_Node_Process_setEnvImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SetEnvImpl
}

func Get_Node_Process_setExitCodeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SetExitCodeImpl
}

func Get_Node_Process_setTitleImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SetTitleImpl
}

func Get_Node_Process_setUncaughtExceptionCaptureCallbackImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_SetUncaughtExceptionCaptureCallbackImpl
}

func Get_Node_Process_stderr() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Stderr
}

func Get_Node_Process_stderrIsTTY() gopurs_runtime.Value {
	return _Gopurs_Node_Process_StderrIsTTY
}

func Get_Node_Process_stdin() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Stdin
}

func Get_Node_Process_stdinIsTTY() gopurs_runtime.Value {
	return _Gopurs_Node_Process_StdinIsTTY
}

func Get_Node_Process_stdout() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Stdout
}

func Get_Node_Process_stdoutIsTTY() gopurs_runtime.Value {
	return _Gopurs_Node_Process_StdoutIsTTY
}

func Get_Node_Process_unsafeGetEnv() gopurs_runtime.Value {
	return _Gopurs_Node_Process_UnsafeGetEnv
}

func Get_Node_Process_unsetEnvImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Process_UnsetEnvImpl
}

func Get_Node_Process_uptime() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Uptime
}

func Get_Node_Process_version() gopurs_runtime.Value {
	return _Gopurs_Node_Process_Version
}
