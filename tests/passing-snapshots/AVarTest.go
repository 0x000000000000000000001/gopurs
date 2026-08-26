package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_liftEffect gopurs_runtime.Value
var once_Main_liftEffect sync.Once

func Get_Main_liftEffect() gopurs_runtime.Value {
	once_Main_liftEffect.Do(func() {
		cache_Main_liftEffect = Get_Effect_Aff__liftEffect()
	})
	return cache_Main_liftEffect
}

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
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff_forkAff(), Call_Main_consumer(avar_1)), gopurs_runtime.Func(func(f1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff_forkAff(), Call_Main_producer(avar_1)), gopurs_runtime.Func(func(f2_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff_joinFiber(), func() gopurs_runtime.Value {
							orig := func() *struct {
								isSuspended gopurs_runtime.Value
								join        gopurs_runtime.Value
								kill        gopurs_runtime.Value
								onComplete  gopurs_runtime.Value
								run         gopurs_runtime.Value
							} {
								orig := f1_2
								_ = orig
								clone := struct {
									isSuspended gopurs_runtime.Value
									join        gopurs_runtime.Value
									kill        gopurs_runtime.Value
									onComplete  gopurs_runtime.Value
									run         gopurs_runtime.Value
								}{}
								clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
								clone.join = gopurs_runtime.RecordGet(orig, "join")
								clone.kill = gopurs_runtime.RecordGet(orig, "kill")
								clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
								clone.run = gopurs_runtime.RecordGet(orig, "run")
								return &clone
							}()
							_ = orig
							return gopurs_runtime.RecordDict([]string{"isSuspended", "join", "kill", "onComplete", "run"}, []gopurs_runtime.Value{orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run})
						}()), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff_joinFiber(), func() gopurs_runtime.Value {
								orig := func() *struct {
									isSuspended gopurs_runtime.Value
									join        gopurs_runtime.Value
									kill        gopurs_runtime.Value
									onComplete  gopurs_runtime.Value
									run         gopurs_runtime.Value
								} {
									orig := f2_3
									_ = orig
									clone := struct {
										isSuspended gopurs_runtime.Value
										join        gopurs_runtime.Value
										kill        gopurs_runtime.Value
										onComplete  gopurs_runtime.Value
										run         gopurs_runtime.Value
									}{}
									clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
									clone.join = gopurs_runtime.RecordGet(orig, "join")
									clone.kill = gopurs_runtime.RecordGet(orig, "kill")
									clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
									clone.run = gopurs_runtime.RecordGet(orig, "run")
									return &clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict([]string{"isSuspended", "join", "kill", "onComplete", "run"}, []gopurs_runtime.Value{orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run})
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
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.UncurriedApp2(Get_Effect_Aff__delay(), Get_Data_Either_Right(), gopurs_runtime.Float(5.0)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Producer: Putting value..."))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply2(Get_Effect_Aff_AVar_put(), gopurs_runtime.Str("Go + PureScript AVar"), avar_0), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Producer: Done!")))
			}))
		}))
	}))
}

func Call_Main_consumer(avar_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var avar_0 gopurs_runtime.Value = avar_0_loop
	_ = avar_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Consumer: Waiting for value..."))), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff_AVar_take(), avar_0), gopurs_runtime.Func(func(val_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
						return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(val_2.StrVal())).StrVal())))
					}))
				}
			end_branch_0:
				return __t0
			}))
		}))
	}))
}
