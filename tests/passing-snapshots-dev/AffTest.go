package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Aff_launchAff_(), gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Starting Aff test..."))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(0)))), gopurs_runtime.Func(func(ref_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_forkAff(Call_Main_loop(int64(100), Call_Main_incrementTask(ref_1))), gopurs_runtime.Func(func(f1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_forkAff(Call_Main_loop(int64(100), Call_Main_incrementTask(ref_1))), gopurs_runtime.Func(func(f2_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
								return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Ref_read(), ref_1)), gopurs_runtime.Func(func(finalVal_6 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t0 gopurs_runtime.Value
									{
										if (finalVal_6.IntVal) == (int64(200)) {
											__t0 = gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Success: 200")))
											goto end_branch_0
										} else {

										}
									}
									{
										__t0 = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fail: Wrong value"))), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(finalVal_6.IntVal)))
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
			__t0 = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), v1_1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Main_loop(gopurs_runtime.IntSub(v_0, int64(1)), v1_1)
			}))
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_incrementTask(ref_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var ref_0 gopurs_runtime.Value = ref_0_loop
	_ = ref_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), Call_Effect_Aff_delay(1.0), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_2.IntVal, int64(1)))
		}), ref_0))
	}))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
