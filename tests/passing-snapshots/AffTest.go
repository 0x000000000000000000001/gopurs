package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_liftEffect gopurs_runtime.Value
var once_Main_liftEffect sync.Once

func Get_Main_liftEffect() gopurs_runtime.Value {
	once_Main_liftEffect.Do(func() {
		cache_Main_liftEffect = Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()))
	})
	return cache_Main_liftEffect
}

var cache_Main_liftEffect1 gopurs_runtime.Value
var once_Main_liftEffect1 sync.Once

func Get_Main_liftEffect1() gopurs_runtime.Value {
	once_Main_liftEffect1.Do(func() {
		cache_Main_liftEffect1 = Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()))
	})
	return cache_Main_liftEffect1
}

var cache_Main_liftEffect2 gopurs_runtime.Value
var once_Main_liftEffect2 sync.Once

func Get_Main_liftEffect2() gopurs_runtime.Value {
	once_Main_liftEffect2.Do(func() {
		cache_Main_liftEffect2 = Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()))
	})
	return cache_Main_liftEffect2
}

var cache_Main_loop gopurs_runtime.Value
var once_Main_loop sync.Once

func Get_Main_loop() gopurs_runtime.Value {
	once_Main_loop.Do(func() {
		cache_Main_loop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_loop(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_loop
}

var cache_Main_incrementTask gopurs_runtime.Value
var once_Main_incrementTask sync.Once

func Get_Main_incrementTask() gopurs_runtime.Value {
	once_Main_incrementTask.Do(func() {
		cache_Main_incrementTask = gopurs_runtime.Func(func(ref_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_incrementTask(ref_0_box)
		})
	})
	return cache_Main_incrementTask
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Aff_launchAff_(), gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Starting Aff test..."))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(0)))), gopurs_runtime.Func(func(ref_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_forkAff(Call_Main_loop(int64(100), Call_Main_incrementTask(ref_1))), gopurs_runtime.Func(func(f1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_forkAff(Call_Main_loop(int64(100), Call_Main_incrementTask(ref_1))), gopurs_runtime.Func(func(f2_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), Call_Effect_Aff_joinFiber(func() struct {
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
							return clone
						}()), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), Call_Effect_Aff_joinFiber(func() struct {
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
								return clone
							}()), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Ref_read(), ref_1)), gopurs_runtime.Func(func(finalVal_6 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t0 gopurs_runtime.Value
									{
										if (finalVal_6.IntVal) == (int64(200)) {
											__t0 = gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Success: 200")))
											goto end_branch_0
										} else {

										}
									}
									{
										__t0 = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail: Wrong value"))), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), finalVal_6).StrVal())))
										}))
									}
								end_branch_0:
									return __t0
								}))
							}))
						}))
					}))
				}))
			}))
		})))
	})
	return cache_Main_main
}

func Call_Main_loop(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
loop:
	for {
		if false {
			continue loop
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit())
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), v1_1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Main_loop((v_0)-(int64(1)), v1_1)
			}))
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_incrementTask(ref_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var ref_0 gopurs_runtime.Value = ref_0_loop
	_ = ref_0
	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.UncurriedApp2(Get_Effect_Aff__delay(), Get_Data_Either_Right(), gopurs_runtime.Float(1.0)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int((x_2.IntVal) + (int64(1)))
		}), ref_0))
	}))
}
