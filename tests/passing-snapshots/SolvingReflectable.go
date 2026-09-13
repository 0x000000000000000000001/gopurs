package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_refString gopurs_runtime.Value
var once_Main_refString sync.Once

func Get_Main_refString() gopurs_runtime.Value {
	once_Main_refString.Do(func() {
		cache_Main_refString = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refString
}

var cache_Main_refStringPass gopurs_runtime.Value
var once_Main_refStringPass sync.Once

func Get_Main_refStringPass() gopurs_runtime.Value {
	once_Main_refStringPass.Do(func() {
		cache_Main_refStringPass = gopurs_runtime.Bool((gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_42398615_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("PureScript")
		})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()) == ("PureScript"))
	})
	return cache_Main_refStringPass
}

var cache_Main_refOrderingLT gopurs_runtime.Value
var once_Main_refOrderingLT sync.Once

func Get_Main_refOrderingLT() gopurs_runtime.Value {
	once_Main_refOrderingLT.Do(func() {
		cache_Main_refOrderingLT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refOrderingLT
}

var cache_Main_refOrderingGT gopurs_runtime.Value
var once_Main_refOrderingGT sync.Once

func Get_Main_refOrderingGT() gopurs_runtime.Value {
	once_Main_refOrderingGT.Do(func() {
		cache_Main_refOrderingGT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refOrderingGT
}

var cache_Main_refOrderingEQ gopurs_runtime.Value
var once_Main_refOrderingEQ sync.Once

func Get_Main_refOrderingEQ() gopurs_runtime.Value {
	once_Main_refOrderingEQ.Do(func() {
		cache_Main_refOrderingEQ = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refOrderingEQ
}

var cache_Main_refOrderingPass gopurs_runtime.Value
var once_Main_refOrderingPass sync.Once

func Get_Main_refOrderingPass() gopurs_runtime.Value {
	once_Main_refOrderingPass.Do(func() {
		cache_Main_refOrderingPass = func() gopurs_runtime.Value {
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_3507070629_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
			})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
			_ = __t_tag_0
			var __t_and_4 bool = false
			if uint32(__t_tag_0.IntVal) == 1527465420 {

				var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_3507070629_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
				})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
				_ = __t_tag_1
				var __t_and_3 bool = false
				if uint32(__t_tag_1.IntVal) == 902936544 {

					var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_3507070629_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
					_ = __t_tag_2
					__t_and_3 = (uint32(__t_tag_2.IntVal) == 380165415)
				}
				__t_and_4 = __t_and_3
			}
			return gopurs_runtime.Bool(__t_and_4)
		}()
	})
	return cache_Main_refOrderingPass
}

var cache_Main_refInt gopurs_runtime.Value
var once_Main_refInt sync.Once

func Get_Main_refInt() gopurs_runtime.Value {
	once_Main_refInt.Do(func() {
		cache_Main_refInt = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refInt
}

var cache_Main_refIntPass gopurs_runtime.Value
var once_Main_refIntPass sync.Once

func Get_Main_refIntPass() gopurs_runtime.Value {
	once_Main_refIntPass.Do(func() {
		cache_Main_refIntPass = gopurs_runtime.Bool((gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_308661683_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(42))
		})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).IntVal) == (int64(42)))
	})
	return cache_Main_refIntPass
}

var cache_Main_refBooleanT gopurs_runtime.Value
var once_Main_refBooleanT sync.Once

func Get_Main_refBooleanT() gopurs_runtime.Value {
	once_Main_refBooleanT.Do(func() {
		cache_Main_refBooleanT = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refBooleanT
}

var cache_Main_refBooleanF gopurs_runtime.Value
var once_Main_refBooleanF sync.Once

func Get_Main_refBooleanF() gopurs_runtime.Value {
	once_Main_refBooleanF.Do(func() {
		cache_Main_refBooleanF = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_refBooleanF
}

var cache_Main_refBooleanPass gopurs_runtime.Value
var once_Main_refBooleanPass sync.Once

func Get_Main_refBooleanPass() gopurs_runtime.Value {
	once_Main_refBooleanPass.Do(func() {
		cache_Main_refBooleanPass = gopurs_runtime.Bool(((gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_3689533068_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).IntVal) != (0)) && (((gopurs_runtime.Apply(Call_Data_Reflectable_reflectType(Rebox_Main_3689533068_465586088((&Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(false)
		})}))), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).IntVal) != (0)) != (true)))
	})
	return cache_Main_refBooleanPass
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			_ = __local_var_0_0
			var __t1 gopurs_runtime.Value
			{
				if ((Get_Main_refIntPass().IntVal) != (0)) && (((Get_Main_refStringPass().IntVal) != (0)) && (((Get_Main_refBooleanPass().IntVal) != (0)) && ((Get_Main_refOrderingPass().IntVal) != (0)))) {
					__t1 = __local_var_0_0
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return Get_Data_Unit_unit()
				})
			}
		end_branch_1:
			return __t1
		}()
	})
	return cache_Main_main
}

func Rebox_Main_308661683_465586088(in *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, int64]) *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3507070629_465586088(in *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, uint32]) *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3689533068_465586088(in *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, bool]) *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_42398615_465586088(in *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, string]) *Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Reflectable_Reflectable[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
