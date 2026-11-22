package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Aff_Compat_bind gopurs_runtime.Value
var once_Effect_Aff_Compat_bind sync.Once

func Get_Effect_Aff_Compat_bind() gopurs_runtime.Value {
	once_Effect_Aff_Compat_bind.Do(func() {
		cache_Effect_Aff_Compat_bind = Get_Control_Bind_bind__406760756()
	})
	return cache_Effect_Aff_Compat_bind
}

var cache_Effect_Aff_Compat_compose gopurs_runtime.Value
var once_Effect_Aff_Compat_compose sync.Once

func Get_Effect_Aff_Compat_compose() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose.Do(func() {
		cache_Effect_Aff_Compat_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Compat_compose
}

var cache_Effect_Aff_Compat_compose__2421196171 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose__2421196171 sync.Once

func Get_Effect_Aff_Compat_compose__2421196171() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose__2421196171.Do(func() {
		cache_Effect_Aff_Compat_compose__2421196171 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_compose__2421196171(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Effect_Aff_Compat_compose__2421196171
}

var cache_Effect_Aff_Compat_compose__3820756491 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose__3820756491 sync.Once

func Get_Effect_Aff_Compat_compose__3820756491() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose__3820756491.Do(func() {
		cache_Effect_Aff_Compat_compose__3820756491 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_compose__3820756491(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Effect_Aff_Compat_compose__3820756491
}

var cache_Effect_Aff_Compat_compose__1905136619 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose__1905136619 sync.Once

func Get_Effect_Aff_Compat_compose__1905136619() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose__1905136619.Do(func() {
		cache_Effect_Aff_Compat_compose__1905136619 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_compose__1905136619(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Effect_Aff_Compat_compose__1905136619
}

var cache_Effect_Aff_Compat_compose1 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose1 sync.Once

func Get_Effect_Aff_Compat_compose1() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose1.Do(func() {
		cache_Effect_Aff_Compat_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Compat_compose1
}

var cache_Effect_Aff_Compat_compose1__3536704176 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose1__3536704176 sync.Once

func Get_Effect_Aff_Compat_compose1__3536704176() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose1__3536704176.Do(func() {
		cache_Effect_Aff_Compat_compose1__3536704176 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_compose1__3536704176(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Effect_Aff_Compat_compose1__3536704176
}

var cache_Effect_Aff_Compat_compose1__3782928176 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose1__3782928176 sync.Once

func Get_Effect_Aff_Compat_compose1__3782928176() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose1__3782928176.Do(func() {
		cache_Effect_Aff_Compat_compose1__3782928176 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_compose1__3782928176(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Effect_Aff_Compat_compose1__3782928176
}

var cache_Effect_Aff_Compat_compose1__3718257072 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose1__3718257072 sync.Once

func Get_Effect_Aff_Compat_compose1__3718257072() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose1__3718257072.Do(func() {
		cache_Effect_Aff_Compat_compose1__3718257072 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_compose1__3718257072(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Effect_Aff_Compat_compose1__3718257072
}

var cache_Effect_Aff_Compat_pure gopurs_runtime.Value
var once_Effect_Aff_Compat_pure sync.Once

func Get_Effect_Aff_Compat_pure() gopurs_runtime.Value {
	once_Effect_Aff_Compat_pure.Do(func() {
		cache_Effect_Aff_Compat_pure = Get_Effect_pureE()
	})
	return cache_Effect_Aff_Compat_pure
}

var cache_Effect_Aff_Compat_discard gopurs_runtime.Value
var once_Effect_Aff_Compat_discard sync.Once

func Get_Effect_Aff_Compat_discard() gopurs_runtime.Value {
	once_Effect_Aff_Compat_discard.Do(func() {
		cache_Effect_Aff_Compat_discard = Get_Control_Bind_discard__1192935977()
	})
	return cache_Effect_Aff_Compat_discard
}

var cache_Effect_Aff_Compat_compose2 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose2 sync.Once

func Get_Effect_Aff_Compat_compose2() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose2.Do(func() {
		cache_Effect_Aff_Compat_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Compat_compose2
}

var cache_Effect_Aff_Compat_compose3 gopurs_runtime.Value
var once_Effect_Aff_Compat_compose3 sync.Once

func Get_Effect_Aff_Compat_compose3() gopurs_runtime.Value {
	once_Effect_Aff_Compat_compose3.Do(func() {
		cache_Effect_Aff_Compat_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Effect_Aff_Compat_compose3
}

var cache_Effect_Aff_Compat_pure1 gopurs_runtime.Value
var once_Effect_Aff_Compat_pure1 sync.Once

func Get_Effect_Aff_Compat_pure1() gopurs_runtime.Value {
	once_Effect_Aff_Compat_pure1.Do(func() {
		cache_Effect_Aff_Compat_pure1 = Get_Control_Applicative_pure__3156000668()
	})
	return cache_Effect_Aff_Compat_pure1
}

var cache_Effect_Aff_Compat_EffectFnCanceler gopurs_runtime.Value
var once_Effect_Aff_Compat_EffectFnCanceler sync.Once

func Get_Effect_Aff_Compat_EffectFnCanceler() gopurs_runtime.Value {
	once_Effect_Aff_Compat_EffectFnCanceler.Do(func() {
		cache_Effect_Aff_Compat_EffectFnCanceler = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_EffectFnCanceler(x_0_box)
		})
	})
	return cache_Effect_Aff_Compat_EffectFnCanceler
}

var cache_Effect_Aff_Compat_EffectFnAff gopurs_runtime.Value
var once_Effect_Aff_Compat_EffectFnAff sync.Once

func Get_Effect_Aff_Compat_EffectFnAff() gopurs_runtime.Value {
	once_Effect_Aff_Compat_EffectFnAff.Do(func() {
		cache_Effect_Aff_Compat_EffectFnAff = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_EffectFnAff(x_0_box)
		})
	})
	return cache_Effect_Aff_Compat_EffectFnAff
}

var cache_Effect_Aff_Compat_fromEffectFnAff gopurs_runtime.Value
var once_Effect_Aff_Compat_fromEffectFnAff sync.Once

func Get_Effect_Aff_Compat_fromEffectFnAff() gopurs_runtime.Value {
	once_Effect_Aff_Compat_fromEffectFnAff.Do(func() {
		cache_Effect_Aff_Compat_fromEffectFnAff = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_fromEffectFnAff(v_0_box)
		})
	})
	return cache_Effect_Aff_Compat_fromEffectFnAff
}

var cache_Effect_Aff_Compat_fromEffectFnAff__2688986814 gopurs_runtime.Value
var once_Effect_Aff_Compat_fromEffectFnAff__2688986814 sync.Once

func Get_Effect_Aff_Compat_fromEffectFnAff__2688986814() gopurs_runtime.Value {
	once_Effect_Aff_Compat_fromEffectFnAff__2688986814.Do(func() {
		cache_Effect_Aff_Compat_fromEffectFnAff__2688986814 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_fromEffectFnAff__2688986814(v_0_box)
		})
	})
	return cache_Effect_Aff_Compat_fromEffectFnAff__2688986814
}

var cache_Effect_Aff_Compat_fromEffectFnAff__560502078 gopurs_runtime.Value
var once_Effect_Aff_Compat_fromEffectFnAff__560502078 sync.Once

func Get_Effect_Aff_Compat_fromEffectFnAff__560502078() gopurs_runtime.Value {
	once_Effect_Aff_Compat_fromEffectFnAff__560502078.Do(func() {
		cache_Effect_Aff_Compat_fromEffectFnAff__560502078 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_fromEffectFnAff__560502078(v_0_box)
		})
	})
	return cache_Effect_Aff_Compat_fromEffectFnAff__560502078
}

var cache_Effect_Aff_Compat_fromEffectFnAff__2546930558 gopurs_runtime.Value
var once_Effect_Aff_Compat_fromEffectFnAff__2546930558 sync.Once

func Get_Effect_Aff_Compat_fromEffectFnAff__2546930558() gopurs_runtime.Value {
	once_Effect_Aff_Compat_fromEffectFnAff__2546930558.Do(func() {
		cache_Effect_Aff_Compat_fromEffectFnAff__2546930558 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Effect_Aff_Compat_fromEffectFnAff__2546930558(v_0_box)
		})
	})
	return cache_Effect_Aff_Compat_fromEffectFnAff__2546930558
}

func Call_Effect_Aff_Compat_compose__2421196171(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2421196171:
	for {
		if false {
			continue compose__2421196171
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Either_Left())
	}
}

func Call_Effect_Aff_Compat_compose__3820756491(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3820756491:
	for {
		if false {
			continue compose__3820756491
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Either_Left())
	}
}

func Call_Effect_Aff_Compat_compose__1905136619(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1905136619:
	for {
		if false {
			continue compose__1905136619
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Either_Left())
	}
}

func Call_Effect_Aff_Compat_compose1__3536704176(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose1__3536704176:
	for {
		if false {
			continue compose1__3536704176
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Either_Right())
	}
}

func Call_Effect_Aff_Compat_compose1__3782928176(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose1__3782928176:
	for {
		if false {
			continue compose1__3782928176
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Either_Right())
	}
}

func Call_Effect_Aff_Compat_compose1__3718257072(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose1__3718257072:
	for {
		if false {
			continue compose1__3718257072
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Either_Right())
	}
}

func Call_Effect_Aff_Compat_EffectFnCanceler(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Effect_Aff_Compat_EffectFnAff(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Effect_Aff_Compat_fromEffectFnAff(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Call_Effect_Aff_makeAff(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_2_0 := gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Left(), __local_var_2), gopurs_runtime.Value{})
			}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Right(), __local_var_2), gopurs_runtime.Value{})
			}))
			_ = __local_var_2_0
			return gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_5_1 := gopurs_runtime.UncurriedApp3(__local_var_2_0, e_3, gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Left(), __local_var_5), gopurs_runtime.Value{})
						}), gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Right(), __local_var_5), gopurs_runtime.Value{})
						}))
						_ = __local_var_5_1
						return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
					})
				}))
			})
		})
	}))
}

func Call_Effect_Aff_Compat_fromEffectFnAff__2688986814(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEffectFnAff__2688986814:
	for {
		if false {
			continue fromEffectFnAff__2688986814
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return Call_Effect_Aff_makeAff__1714332378(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_2_0 := gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Left(), __local_var_2), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Right(), __local_var_2), gopurs_runtime.Value{})
				}))
				_ = __local_var_2_0
				return gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							__local_var_5_1 := gopurs_runtime.UncurriedApp3(__local_var_2_0, e_3, gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Left(), __local_var_5), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Right(), __local_var_5), gopurs_runtime.Value{})
							}))
							_ = __local_var_5_1
							return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
						})
					}))
				})
			})
		}))
	}
}

func Call_Effect_Aff_Compat_fromEffectFnAff__560502078(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEffectFnAff__560502078:
	for {
		if false {
			continue fromEffectFnAff__560502078
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return Call_Effect_Aff_makeAff__2246750170(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_2_0 := gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Left(), __local_var_2), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Right(), __local_var_2), gopurs_runtime.Value{})
				}))
				_ = __local_var_2_0
				return gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							__local_var_5_1 := gopurs_runtime.UncurriedApp3(__local_var_2_0, e_3, gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Left(), __local_var_5), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Right(), __local_var_5), gopurs_runtime.Value{})
							}))
							_ = __local_var_5_1
							return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
						})
					}))
				})
			})
		}))
	}
}

func Call_Effect_Aff_Compat_fromEffectFnAff__2546930558(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEffectFnAff__2546930558:
	for {
		if false {
			continue fromEffectFnAff__2546930558
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return Call_Effect_Aff_makeAff__158558778(gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_2_0 := gopurs_runtime.UncurriedApp2(v_0, gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Left(), __local_var_2), gopurs_runtime.Value{})
				}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k_1, Get_Data_Either_Right(), __local_var_2), gopurs_runtime.Value{})
				}))
				_ = __local_var_2_0
				return gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return Call_Effect_Aff_makeAff__925755130(gopurs_runtime.Func(func(k2_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							__local_var_5_1 := gopurs_runtime.UncurriedApp3(__local_var_2_0, e_3, gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Left(), __local_var_5), gopurs_runtime.Value{})
							}), gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), k2_4, Get_Data_Either_Right(), __local_var_5), gopurs_runtime.Value{})
							}))
							_ = __local_var_5_1
							return gopurs_runtime.Apply(Get_Data_Function_const__4207166068(), gopurs_runtime.Apply(Get_Effect_Aff__pure(), Get_Data_Unit_unit()))
						})
					}))
				})
			})
		}))
	}
}
