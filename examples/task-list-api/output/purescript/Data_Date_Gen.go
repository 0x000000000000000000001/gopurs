package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_Gen_compose gopurs_runtime.Value
var once_Data_Date_Gen_compose sync.Once

func Get_Data_Date_Gen_compose() gopurs_runtime.Value {
	once_Data_Date_Gen_compose.Do(func() {
		cache_Data_Date_Gen_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Date_Gen_compose
}

var cache_Data_Date_Gen_fromJust gopurs_runtime.Value
var once_Data_Date_Gen_fromJust sync.Once

func Get_Data_Date_Gen_fromJust() gopurs_runtime.Value {
	once_Data_Date_Gen_fromJust.Do(func() {
		cache_Data_Date_Gen_fromJust = Get_Data_Maybe_fromJust__3946112507()
	})
	return cache_Data_Date_Gen_fromJust
}

var cache_Data_Date_Gen_bind gopurs_runtime.Value
var once_Data_Date_Gen_bind sync.Once

func Get_Data_Date_Gen_bind() gopurs_runtime.Value {
	once_Data_Date_Gen_bind.Do(func() {
		cache_Data_Date_Gen_bind = Get_Control_Bind_bind__3973282666()
	})
	return cache_Data_Date_Gen_bind
}

var cache_Data_Date_Gen_bottom gopurs_runtime.Value
var once_Data_Date_Gen_bottom sync.Once

func Get_Data_Date_Gen_bottom() gopurs_runtime.Value {
	once_Data_Date_Gen_bottom.Do(func() {
		cache_Data_Date_Gen_bottom = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
	})
	return cache_Data_Date_Gen_bottom
}

var cache_Data_Date_Gen_bottom1 gopurs_runtime.Value
var once_Data_Date_Gen_bottom1 sync.Once

func Get_Data_Date_Gen_bottom1() gopurs_runtime.Value {
	once_Data_Date_Gen_bottom1.Do(func() {
		cache_Data_Date_Gen_bottom1 = gopurs_runtime.Int(int64(1))
	})
	return cache_Data_Date_Gen_bottom1
}

var cache_Data_Date_Gen_genDate gopurs_runtime.Value
var once_Data_Date_Gen_genDate sync.Once

func Get_Data_Date_Gen_genDate() gopurs_runtime.Value {
	once_Data_Date_Gen_genDate.Do(func() {
		cache_Data_Date_Gen_genDate = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Date_Gen_genDate(dictMonadGen_0_box)
		})
	})
	return cache_Data_Date_Gen_genDate
}

func Call_Data_Date_Gen_genDate(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope1)])
	Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
	_ = Bind1_2_1
	// TAST (Let): pure_3_2 shape=App(Var) bindingType=(Func [(ADT ["Data","Date","Date"] [])] (TypeApp (TypeVar m$scope1) [(ADT ["Data","Date","Date"] [])]))
	pure_3_2 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = pure_3_2
	return gopurs_runtime.Apply2(Bind1_2_1.V1, Call_Data_Date_Component_Gen_genYear(dictMonadGen_0), gopurs_runtime.Func(func(year_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 int64
		{
			if ((gopurs_runtime.IntMod(year_4.IntVal, int64(4))) == (int64(0))) && (((gopurs_runtime.IntMod(year_4.IntVal, int64(400))) == (int64(0))) || (((gopurs_runtime.IntMod(year_4.IntVal, int64(100))) == (int64(0))) != (true))) {
				__t3 = int64(365)
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = int64(364)
		}
	end_branch_3:
		return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_5
		}), Get_Data_Int_toNumber()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(__t3))), gopurs_runtime.Func(func(days_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial(), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_7_4 shape=App(Var) bindingType=Any
				__local_var_7_4 := Call_Data_Date_exactDate(year_4.IntVal, 1908470532, int64(1))
				_ = __local_var_7_4
				var __t6 gopurs_runtime.Value
				{
					if __local_var_7_4.V1 {
						__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Gen_2280409795_3094389156(Rebox_Data_Date_Gen_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := Call_Data_Date_adjust(days_5.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__local_var_7_4.V0))
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))))}
						goto end_branch_6
					} else {

					}
				}
				{
					if !__local_var_7_4.V1 {
						__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))}
						goto end_branch_6
					} else {

					}
				}
				{
					__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
				}
			end_branch_6:
				// TAST (Let): __local_var_8_5 shape=Branch(App(Var), Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope103)])
				__local_var_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t6)
				_ = __local_var_8_5
				var __t7 gopurs_runtime.Value
				{
					if __local_var_8_5 != nil {
						__t7 = (__local_var_8_5).V0
						goto end_branch_7
					} else {

					}
				}
				{
					__t7 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Date_Date { panic("Failed pattern match") }())}
				}
			end_branch_7:
				return __t7
			}))))})
		}))
	}))
}

func Rebox_Data_Date_Gen_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Date_Gen_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}
