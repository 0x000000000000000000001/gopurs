package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_producer gopurs_runtime.Value
var once_Main_producer sync.Once

func Get_Main_producer() gopurs_runtime.Value {
	once_Main_producer.Do(func() {
		cache_Main_producer = gopurs_runtime.Func(func(avar_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_producer(avar_0_box)
		})
	})
	return cache_Main_producer
}

var cache_Main_consumer gopurs_runtime.Value
var once_Main_consumer sync.Once

func Get_Main_consumer() gopurs_runtime.Value {
	once_Main_consumer.Do(func() {
		cache_Main_consumer = gopurs_runtime.Func(func(avar_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_consumer(avar_0_box)
		})
	})
	return cache_Main_consumer
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Aff_launchAff_(), gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Starting AVar test..."))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Effect_AVar_empty()), gopurs_runtime.Func(func(avar_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_forkAff(Call_Main_consumer(avar_1)), gopurs_runtime.Func(func(f1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_forkAff(Call_Main_producer(avar_1)), gopurs_runtime.Func(func(f2_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_joinFiber(func() struct {
							isSuspended gopurs_runtime.Value
							join        gopurs_runtime.Value
							kill        gopurs_runtime.Value
							onComplete  gopurs_runtime.Value
							run         gopurs_runtime.Value
							/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
						} {
							orig := f1_2
							_ = orig
							clone := struct {
								isSuspended gopurs_runtime.Value
								join        gopurs_runtime.Value
								kill        gopurs_runtime.Value
								onComplete  gopurs_runtime.Value
								run         gopurs_runtime.Value
								/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
							}{}
							clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
							clone.join = gopurs_runtime.RecordGet(orig, "join")
							clone.kill = gopurs_runtime.RecordGet(orig, "kill")
							clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
							clone.run = gopurs_runtime.RecordGet(orig, "run")
							return clone
						}()), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_joinFiber(func() struct {
								isSuspended gopurs_runtime.Value
								join        gopurs_runtime.Value
								kill        gopurs_runtime.Value
								onComplete  gopurs_runtime.Value
								run         gopurs_runtime.Value
								/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
							} {
								orig := f2_3
								_ = orig
								clone := struct {
									isSuspended gopurs_runtime.Value
									join        gopurs_runtime.Value
									kill        gopurs_runtime.Value
									onComplete  gopurs_runtime.Value
									run         gopurs_runtime.Value
									/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
								}{}
								clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
								clone.join = gopurs_runtime.RecordGet(orig, "join")
								clone.kill = gopurs_runtime.RecordGet(orig, "kill")
								clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
								clone.run = gopurs_runtime.RecordGet(orig, "run")
								return clone
							}()), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Test completed.")))
							}))
						}))
					}))
				}))
			}))
		})))
	})
	return cache_Main_main
}

func Call_Main_producer(avar_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var avar_0 gopurs_runtime.Value = avar_0_loop
	_ = avar_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_delay(5.0), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Producer: Putting value..."))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_AVar_put(gopurs_runtime.Str("Go + PureScript AVar"), avar_0), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Producer: Done!")))
			}))
		}))
	}))
}

func Call_Main_consumer(avar_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var avar_0 gopurs_runtime.Value = avar_0_loop
	_ = avar_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Consumer: Waiting for value..."))), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_AVar_take(avar_0), gopurs_runtime.Func(func(val_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Consumer: Got value!"))), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 gopurs_runtime.Value
				{
					if (val_2.StrVal()) == ("Go + PureScript AVar") {
						__t0 = gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Success: AVar value matches")))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail: Wrong AVar value"))), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Call_Effect_Console_logShow(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))), gopurs_runtime.Str(val_2.StrVal())))
					}))
				}
			end_branch_0:
				return __t0
			}))
		}))
	}))
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
