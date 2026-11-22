package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Semigroupoid_Semigroupoid_dollar_Dict gopurs_runtime.Value
var once_Control_Semigroupoid_Semigroupoid_dollar_Dict sync.Once

func Get_Control_Semigroupoid_Semigroupoid_dollar_Dict() gopurs_runtime.Value {
	once_Control_Semigroupoid_Semigroupoid_dollar_Dict.Do(func() {
		cache_Control_Semigroupoid_Semigroupoid_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(Call_Control_Semigroupoid_Semigroupoid_dollar_Dict(func() struct {
				compose gopurs_runtime.Value
				/*gopurs:record-v1:1:0063006f006d0070006f00730065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					compose gopurs_runtime.Value
					/*gopurs:record-v1:1:0063006f006d0070006f00730065*/
				}{}
				clone.compose = gopurs_runtime.RecordGet(orig, "compose")
				return clone
			}()))}
		})
	})
	return cache_Control_Semigroupoid_Semigroupoid_dollar_Dict
}

var cache_Control_Semigroupoid_semigroupoidFn gopurs_runtime.Value
var once_Control_Semigroupoid_semigroupoidFn sync.Once

func Get_Control_Semigroupoid_semigroupoidFn() gopurs_runtime.Value {
	once_Control_Semigroupoid_semigroupoidFn.Do(func() {
		cache_Control_Semigroupoid_semigroupoidFn = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer((&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
		})}))}
	})
	return cache_Control_Semigroupoid_semigroupoidFn
}

var cache_Control_Semigroupoid_compose gopurs_runtime.Value
var once_Control_Semigroupoid_compose sync.Once

func Get_Control_Semigroupoid_compose() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose.Do(func() {
		cache_Control_Semigroupoid_compose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Semigroupoid_compose
}

var cache_Control_Semigroupoid_compose__3872452778 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3872452778 sync.Once

func Get_Control_Semigroupoid_compose__3872452778() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3872452778.Do(func() {
		cache_Control_Semigroupoid_compose__3872452778 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3872452778(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3872452778
}

var cache_Control_Semigroupoid_compose__1323742214 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1323742214 sync.Once

func Get_Control_Semigroupoid_compose__1323742214() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1323742214.Do(func() {
		cache_Control_Semigroupoid_compose__1323742214 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1323742214(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1323742214
}

var cache_Control_Semigroupoid_compose__251463616 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__251463616 sync.Once

func Get_Control_Semigroupoid_compose__251463616() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__251463616.Do(func() {
		cache_Control_Semigroupoid_compose__251463616 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__251463616(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__251463616
}

var cache_Control_Semigroupoid_compose__3629354669 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3629354669 sync.Once

func Get_Control_Semigroupoid_compose__3629354669() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3629354669.Do(func() {
		cache_Control_Semigroupoid_compose__3629354669 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3629354669(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3629354669
}

var cache_Control_Semigroupoid_compose__3537576912 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3537576912 sync.Once

func Get_Control_Semigroupoid_compose__3537576912() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3537576912.Do(func() {
		cache_Control_Semigroupoid_compose__3537576912 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3537576912(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3537576912
}

var cache_Control_Semigroupoid_compose__953525693 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__953525693 sync.Once

func Get_Control_Semigroupoid_compose__953525693() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__953525693.Do(func() {
		cache_Control_Semigroupoid_compose__953525693 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__953525693(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__953525693
}

var cache_Control_Semigroupoid_compose__1578359809 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1578359809 sync.Once

func Get_Control_Semigroupoid_compose__1578359809() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1578359809.Do(func() {
		cache_Control_Semigroupoid_compose__1578359809 = gopurs_runtime.Func3(func(__eta_norm_2_unused_0_box gopurs_runtime.Value, __eta_norm_1_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1578359809(__eta_norm_2_unused_0_box, __eta_norm_1_unused_1_box, __eta_norm_0_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1578359809
}

var cache_Control_Semigroupoid_compose__1084262850 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1084262850 sync.Once

func Get_Control_Semigroupoid_compose__1084262850() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1084262850.Do(func() {
		cache_Control_Semigroupoid_compose__1084262850 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1084262850(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1084262850
}

var cache_Control_Semigroupoid_compose__2472072002 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2472072002 sync.Once

func Get_Control_Semigroupoid_compose__2472072002() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2472072002.Do(func() {
		cache_Control_Semigroupoid_compose__2472072002 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2472072002(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2472072002
}

var cache_Control_Semigroupoid_compose__635039573 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__635039573 sync.Once

func Get_Control_Semigroupoid_compose__635039573() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__635039573.Do(func() {
		cache_Control_Semigroupoid_compose__635039573 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__635039573(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__635039573
}

var cache_Control_Semigroupoid_compose__817328085 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__817328085 sync.Once

func Get_Control_Semigroupoid_compose__817328085() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__817328085.Do(func() {
		cache_Control_Semigroupoid_compose__817328085 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__817328085(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__817328085
}

var cache_Control_Semigroupoid_compose__1053824693 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1053824693 sync.Once

func Get_Control_Semigroupoid_compose__1053824693() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1053824693.Do(func() {
		cache_Control_Semigroupoid_compose__1053824693 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1053824693(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1053824693
}

var cache_Control_Semigroupoid_compose__706442655 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__706442655 sync.Once

func Get_Control_Semigroupoid_compose__706442655() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__706442655.Do(func() {
		cache_Control_Semigroupoid_compose__706442655 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__706442655(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__706442655
}

var cache_Control_Semigroupoid_compose__2752544656 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2752544656 sync.Once

func Get_Control_Semigroupoid_compose__2752544656() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2752544656.Do(func() {
		cache_Control_Semigroupoid_compose__2752544656 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2752544656(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2752544656
}

var cache_Control_Semigroupoid_compose__2473888545 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2473888545 sync.Once

func Get_Control_Semigroupoid_compose__2473888545() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2473888545.Do(func() {
		cache_Control_Semigroupoid_compose__2473888545 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2473888545(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2473888545
}

var cache_Control_Semigroupoid_compose__81161766 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__81161766 sync.Once

func Get_Control_Semigroupoid_compose__81161766() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__81161766.Do(func() {
		cache_Control_Semigroupoid_compose__81161766 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__81161766(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__81161766
}

var cache_Control_Semigroupoid_compose__18355393 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__18355393 sync.Once

func Get_Control_Semigroupoid_compose__18355393() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__18355393.Do(func() {
		cache_Control_Semigroupoid_compose__18355393 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__18355393(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__18355393
}

var cache_Control_Semigroupoid_compose__297303379 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__297303379 sync.Once

func Get_Control_Semigroupoid_compose__297303379() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__297303379.Do(func() {
		cache_Control_Semigroupoid_compose__297303379 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__297303379(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__297303379
}

var cache_Control_Semigroupoid_compose__1680806490 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1680806490 sync.Once

func Get_Control_Semigroupoid_compose__1680806490() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1680806490.Do(func() {
		cache_Control_Semigroupoid_compose__1680806490 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1680806490(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1680806490
}

var cache_Control_Semigroupoid_compose__262436915 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__262436915 sync.Once

func Get_Control_Semigroupoid_compose__262436915() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__262436915.Do(func() {
		cache_Control_Semigroupoid_compose__262436915 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__262436915(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__262436915
}

var cache_Control_Semigroupoid_compose__2920693720 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2920693720 sync.Once

func Get_Control_Semigroupoid_compose__2920693720() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2920693720.Do(func() {
		cache_Control_Semigroupoid_compose__2920693720 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2920693720(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2920693720
}

var cache_Control_Semigroupoid_compose__2990151384 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2990151384 sync.Once

func Get_Control_Semigroupoid_compose__2990151384() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2990151384.Do(func() {
		cache_Control_Semigroupoid_compose__2990151384 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2990151384(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2990151384
}

var cache_Control_Semigroupoid_compose__3715450552 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3715450552 sync.Once

func Get_Control_Semigroupoid_compose__3715450552() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3715450552.Do(func() {
		cache_Control_Semigroupoid_compose__3715450552 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3715450552(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3715450552
}

var cache_Control_Semigroupoid_compose__2907274104 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2907274104 sync.Once

func Get_Control_Semigroupoid_compose__2907274104() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2907274104.Do(func() {
		cache_Control_Semigroupoid_compose__2907274104 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2907274104(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2907274104
}

var cache_Control_Semigroupoid_compose__13732245 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__13732245 sync.Once

func Get_Control_Semigroupoid_compose__13732245() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__13732245.Do(func() {
		cache_Control_Semigroupoid_compose__13732245 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__13732245(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__13732245
}

var cache_Control_Semigroupoid_compose__447393272 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__447393272 sync.Once

func Get_Control_Semigroupoid_compose__447393272() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__447393272.Do(func() {
		cache_Control_Semigroupoid_compose__447393272 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__447393272(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__447393272
}

var cache_Control_Semigroupoid_compose__985198819 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__985198819 sync.Once

func Get_Control_Semigroupoid_compose__985198819() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__985198819.Do(func() {
		cache_Control_Semigroupoid_compose__985198819 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__985198819(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__985198819
}

var cache_Control_Semigroupoid_compose__1195436277 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1195436277 sync.Once

func Get_Control_Semigroupoid_compose__1195436277() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1195436277.Do(func() {
		cache_Control_Semigroupoid_compose__1195436277 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1195436277(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1195436277
}

var cache_Control_Semigroupoid_compose__2013581781 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2013581781 sync.Once

func Get_Control_Semigroupoid_compose__2013581781() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2013581781.Do(func() {
		cache_Control_Semigroupoid_compose__2013581781 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2013581781(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2013581781
}

var cache_Control_Semigroupoid_compose__3089340516 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3089340516 sync.Once

func Get_Control_Semigroupoid_compose__3089340516() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3089340516.Do(func() {
		cache_Control_Semigroupoid_compose__3089340516 = gopurs_runtime.Func3(func(__eta_norm_2_unused_0_box gopurs_runtime.Value, __eta_norm_1_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3089340516(__eta_norm_2_unused_0_box, __eta_norm_1_unused_1_box, __eta_norm_0_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3089340516
}

var cache_Control_Semigroupoid_compose__3958277077 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3958277077 sync.Once

func Get_Control_Semigroupoid_compose__3958277077() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3958277077.Do(func() {
		cache_Control_Semigroupoid_compose__3958277077 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3958277077(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3958277077
}

var cache_Control_Semigroupoid_compose__3383039189 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3383039189 sync.Once

func Get_Control_Semigroupoid_compose__3383039189() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3383039189.Do(func() {
		cache_Control_Semigroupoid_compose__3383039189 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3383039189(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3383039189
}

var cache_Control_Semigroupoid_compose__2488896981 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2488896981 sync.Once

func Get_Control_Semigroupoid_compose__2488896981() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2488896981.Do(func() {
		cache_Control_Semigroupoid_compose__2488896981 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2488896981(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2488896981
}

var cache_Control_Semigroupoid_compose__3939066901 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3939066901 sync.Once

func Get_Control_Semigroupoid_compose__3939066901() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3939066901.Do(func() {
		cache_Control_Semigroupoid_compose__3939066901 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3939066901(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3939066901
}

var cache_Control_Semigroupoid_compose__2813039818 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2813039818 sync.Once

func Get_Control_Semigroupoid_compose__2813039818() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2813039818.Do(func() {
		cache_Control_Semigroupoid_compose__2813039818 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2813039818(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2813039818
}

var cache_Control_Semigroupoid_compose__4004832778 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4004832778 sync.Once

func Get_Control_Semigroupoid_compose__4004832778() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4004832778.Do(func() {
		cache_Control_Semigroupoid_compose__4004832778 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__4004832778(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__4004832778
}

var cache_Control_Semigroupoid_compose__2889719291 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2889719291 sync.Once

func Get_Control_Semigroupoid_compose__2889719291() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2889719291.Do(func() {
		cache_Control_Semigroupoid_compose__2889719291 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2889719291(__eta_norm_1_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2889719291
}

var cache_Control_Semigroupoid_compose__1982962581 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1982962581 sync.Once

func Get_Control_Semigroupoid_compose__1982962581() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1982962581.Do(func() {
		cache_Control_Semigroupoid_compose__1982962581 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1982962581(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1982962581
}

var cache_Control_Semigroupoid_compose__3606255701 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3606255701 sync.Once

func Get_Control_Semigroupoid_compose__3606255701() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3606255701.Do(func() {
		cache_Control_Semigroupoid_compose__3606255701 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3606255701(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3606255701
}

var cache_Control_Semigroupoid_compose__2491846907 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2491846907 sync.Once

func Get_Control_Semigroupoid_compose__2491846907() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2491846907.Do(func() {
		cache_Control_Semigroupoid_compose__2491846907 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2491846907(__eta_norm_1_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2491846907
}

var cache_Control_Semigroupoid_compose__580204084 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__580204084 sync.Once

func Get_Control_Semigroupoid_compose__580204084() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__580204084.Do(func() {
		cache_Control_Semigroupoid_compose__580204084 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__580204084(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__580204084
}

var cache_Control_Semigroupoid_compose__3042204923 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3042204923 sync.Once

func Get_Control_Semigroupoid_compose__3042204923() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3042204923.Do(func() {
		cache_Control_Semigroupoid_compose__3042204923 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3042204923(__eta_norm_1_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3042204923
}

var cache_Control_Semigroupoid_compose__323228359 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__323228359 sync.Once

func Get_Control_Semigroupoid_compose__323228359() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__323228359.Do(func() {
		cache_Control_Semigroupoid_compose__323228359 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__323228359(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__323228359
}

var cache_Control_Semigroupoid_compose__788005840 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__788005840 sync.Once

func Get_Control_Semigroupoid_compose__788005840() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__788005840.Do(func() {
		cache_Control_Semigroupoid_compose__788005840 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__788005840(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__788005840
}

var cache_Control_Semigroupoid_compose__1772354132 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1772354132 sync.Once

func Get_Control_Semigroupoid_compose__1772354132() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1772354132.Do(func() {
		cache_Control_Semigroupoid_compose__1772354132 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1772354132(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1772354132
}

var cache_Control_Semigroupoid_compose__50666552 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__50666552 sync.Once

func Get_Control_Semigroupoid_compose__50666552() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__50666552.Do(func() {
		cache_Control_Semigroupoid_compose__50666552 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__50666552(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__50666552
}

var cache_Control_Semigroupoid_compose__2152657429 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2152657429 sync.Once

func Get_Control_Semigroupoid_compose__2152657429() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2152657429.Do(func() {
		cache_Control_Semigroupoid_compose__2152657429 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2152657429(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2152657429
}

var cache_Control_Semigroupoid_compose__2655475989 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2655475989 sync.Once

func Get_Control_Semigroupoid_compose__2655475989() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2655475989.Do(func() {
		cache_Control_Semigroupoid_compose__2655475989 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2655475989(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2655475989
}

var cache_Control_Semigroupoid_compose__2593021333 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2593021333 sync.Once

func Get_Control_Semigroupoid_compose__2593021333() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2593021333.Do(func() {
		cache_Control_Semigroupoid_compose__2593021333 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2593021333(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2593021333
}

var cache_Control_Semigroupoid_compose__2246386360 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2246386360 sync.Once

func Get_Control_Semigroupoid_compose__2246386360() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2246386360.Do(func() {
		cache_Control_Semigroupoid_compose__2246386360 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2246386360(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2246386360
}

var cache_Control_Semigroupoid_compose__3326469365 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3326469365 sync.Once

func Get_Control_Semigroupoid_compose__3326469365() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3326469365.Do(func() {
		cache_Control_Semigroupoid_compose__3326469365 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3326469365(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3326469365
}

var cache_Control_Semigroupoid_compose__3055166288 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3055166288 sync.Once

func Get_Control_Semigroupoid_compose__3055166288() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3055166288.Do(func() {
		cache_Control_Semigroupoid_compose__3055166288 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3055166288(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3055166288
}

var cache_Control_Semigroupoid_compose__1433444296 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1433444296 sync.Once

func Get_Control_Semigroupoid_compose__1433444296() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1433444296.Do(func() {
		cache_Control_Semigroupoid_compose__1433444296 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1433444296(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1433444296
}

var cache_Control_Semigroupoid_compose__2650337094 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2650337094 sync.Once

func Get_Control_Semigroupoid_compose__2650337094() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2650337094.Do(func() {
		cache_Control_Semigroupoid_compose__2650337094 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2650337094(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2650337094
}

var cache_Control_Semigroupoid_compose__661099861 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__661099861 sync.Once

func Get_Control_Semigroupoid_compose__661099861() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__661099861.Do(func() {
		cache_Control_Semigroupoid_compose__661099861 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__661099861(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__661099861
}

var cache_Control_Semigroupoid_compose__2074229983 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2074229983 sync.Once

func Get_Control_Semigroupoid_compose__2074229983() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2074229983.Do(func() {
		cache_Control_Semigroupoid_compose__2074229983 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2074229983(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2074229983
}

var cache_Control_Semigroupoid_compose__3307342101 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3307342101 sync.Once

func Get_Control_Semigroupoid_compose__3307342101() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3307342101.Do(func() {
		cache_Control_Semigroupoid_compose__3307342101 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3307342101(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3307342101
}

var cache_Control_Semigroupoid_compose__419544780 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__419544780 sync.Once

func Get_Control_Semigroupoid_compose__419544780() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__419544780.Do(func() {
		cache_Control_Semigroupoid_compose__419544780 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__419544780(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__419544780
}

var cache_Control_Semigroupoid_compose__1934312767 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1934312767 sync.Once

func Get_Control_Semigroupoid_compose__1934312767() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1934312767.Do(func() {
		cache_Control_Semigroupoid_compose__1934312767 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1934312767(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1934312767
}

var cache_Control_Semigroupoid_compose__1157135857 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1157135857 sync.Once

func Get_Control_Semigroupoid_compose__1157135857() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1157135857.Do(func() {
		cache_Control_Semigroupoid_compose__1157135857 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1157135857(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1157135857
}

var cache_Control_Semigroupoid_compose__1739002828 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1739002828 sync.Once

func Get_Control_Semigroupoid_compose__1739002828() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1739002828.Do(func() {
		cache_Control_Semigroupoid_compose__1739002828 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1739002828(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1739002828
}

var cache_Control_Semigroupoid_compose__3898604607 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3898604607 sync.Once

func Get_Control_Semigroupoid_compose__3898604607() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3898604607.Do(func() {
		cache_Control_Semigroupoid_compose__3898604607 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3898604607(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3898604607
}

var cache_Control_Semigroupoid_compose__585867329 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__585867329 sync.Once

func Get_Control_Semigroupoid_compose__585867329() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__585867329.Do(func() {
		cache_Control_Semigroupoid_compose__585867329 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__585867329(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__585867329
}

var cache_Control_Semigroupoid_compose__2240476445 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2240476445 sync.Once

func Get_Control_Semigroupoid_compose__2240476445() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2240476445.Do(func() {
		cache_Control_Semigroupoid_compose__2240476445 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2240476445(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2240476445
}

var cache_Control_Semigroupoid_compose__1284278492 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1284278492 sync.Once

func Get_Control_Semigroupoid_compose__1284278492() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1284278492.Do(func() {
		cache_Control_Semigroupoid_compose__1284278492 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1284278492(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1284278492
}

var cache_Control_Semigroupoid_compose__3320522428 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3320522428 sync.Once

func Get_Control_Semigroupoid_compose__3320522428() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3320522428.Do(func() {
		cache_Control_Semigroupoid_compose__3320522428 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3320522428(__eta_norm_1_unused_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3320522428
}

var cache_Control_Semigroupoid_compose__1048546465 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1048546465 sync.Once

func Get_Control_Semigroupoid_compose__1048546465() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1048546465.Do(func() {
		cache_Control_Semigroupoid_compose__1048546465 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1048546465(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1048546465
}

var cache_Control_Semigroupoid_compose__2993739570 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2993739570 sync.Once

func Get_Control_Semigroupoid_compose__2993739570() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2993739570.Do(func() {
		cache_Control_Semigroupoid_compose__2993739570 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2993739570(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2993739570
}

var cache_Control_Semigroupoid_compose__2865838076 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2865838076 sync.Once

func Get_Control_Semigroupoid_compose__2865838076() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2865838076.Do(func() {
		cache_Control_Semigroupoid_compose__2865838076 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2865838076(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2865838076
}

var cache_Control_Semigroupoid_compose__13873697 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__13873697 sync.Once

func Get_Control_Semigroupoid_compose__13873697() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__13873697.Do(func() {
		cache_Control_Semigroupoid_compose__13873697 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__13873697(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__13873697
}

var cache_Control_Semigroupoid_compose__4254564146 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4254564146 sync.Once

func Get_Control_Semigroupoid_compose__4254564146() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4254564146.Do(func() {
		cache_Control_Semigroupoid_compose__4254564146 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__4254564146(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__4254564146
}

var cache_Control_Semigroupoid_compose__1304014924 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1304014924 sync.Once

func Get_Control_Semigroupoid_compose__1304014924() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1304014924.Do(func() {
		cache_Control_Semigroupoid_compose__1304014924 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1304014924(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1304014924
}

var cache_Control_Semigroupoid_compose__494499605 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__494499605 sync.Once

func Get_Control_Semigroupoid_compose__494499605() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__494499605.Do(func() {
		cache_Control_Semigroupoid_compose__494499605 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__494499605(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__494499605
}

var cache_Control_Semigroupoid_compose__3460990597 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3460990597 sync.Once

func Get_Control_Semigroupoid_compose__3460990597() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3460990597.Do(func() {
		cache_Control_Semigroupoid_compose__3460990597 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3460990597(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3460990597
}

var cache_Control_Semigroupoid_compose__178943130 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__178943130 sync.Once

func Get_Control_Semigroupoid_compose__178943130() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__178943130.Do(func() {
		cache_Control_Semigroupoid_compose__178943130 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__178943130(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__178943130
}

var cache_Control_Semigroupoid_compose__100839156 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__100839156 sync.Once

func Get_Control_Semigroupoid_compose__100839156() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__100839156.Do(func() {
		cache_Control_Semigroupoid_compose__100839156 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__100839156(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__100839156
}

var cache_Control_Semigroupoid_compose__1866182613 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1866182613 sync.Once

func Get_Control_Semigroupoid_compose__1866182613() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1866182613.Do(func() {
		cache_Control_Semigroupoid_compose__1866182613 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1866182613(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1866182613
}

var cache_Control_Semigroupoid_compose__2640382690 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2640382690 sync.Once

func Get_Control_Semigroupoid_compose__2640382690() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2640382690.Do(func() {
		cache_Control_Semigroupoid_compose__2640382690 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2640382690(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2640382690
}

var cache_Control_Semigroupoid_compose__529030880 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__529030880 sync.Once

func Get_Control_Semigroupoid_compose__529030880() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__529030880.Do(func() {
		cache_Control_Semigroupoid_compose__529030880 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__529030880(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__529030880
}

var cache_Control_Semigroupoid_compose__532200073 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__532200073 sync.Once

func Get_Control_Semigroupoid_compose__532200073() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__532200073.Do(func() {
		cache_Control_Semigroupoid_compose__532200073 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__532200073(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__532200073
}

var cache_Control_Semigroupoid_compose__880864056 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__880864056 sync.Once

func Get_Control_Semigroupoid_compose__880864056() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__880864056.Do(func() {
		cache_Control_Semigroupoid_compose__880864056 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__880864056(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_compose__880864056
}

var cache_Control_Semigroupoid_compose__172418801 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__172418801 sync.Once

func Get_Control_Semigroupoid_compose__172418801() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__172418801.Do(func() {
		cache_Control_Semigroupoid_compose__172418801 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__172418801(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__172418801
}

var cache_Control_Semigroupoid_compose__1247429674 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1247429674 sync.Once

func Get_Control_Semigroupoid_compose__1247429674() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1247429674.Do(func() {
		cache_Control_Semigroupoid_compose__1247429674 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1247429674(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1247429674
}

var cache_Control_Semigroupoid_compose__446754736 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__446754736 sync.Once

func Get_Control_Semigroupoid_compose__446754736() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__446754736.Do(func() {
		cache_Control_Semigroupoid_compose__446754736 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__446754736(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__446754736
}

var cache_Control_Semigroupoid_compose__281216780 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__281216780 sync.Once

func Get_Control_Semigroupoid_compose__281216780() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__281216780.Do(func() {
		cache_Control_Semigroupoid_compose__281216780 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__281216780(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__281216780
}

var cache_Control_Semigroupoid_compose__204472163 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__204472163 sync.Once

func Get_Control_Semigroupoid_compose__204472163() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__204472163.Do(func() {
		cache_Control_Semigroupoid_compose__204472163 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__204472163(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__204472163
}

var cache_Control_Semigroupoid_compose__4153383075 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4153383075 sync.Once

func Get_Control_Semigroupoid_compose__4153383075() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4153383075.Do(func() {
		cache_Control_Semigroupoid_compose__4153383075 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__4153383075(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__4153383075
}

var cache_Control_Semigroupoid_compose__330160021 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__330160021 sync.Once

func Get_Control_Semigroupoid_compose__330160021() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__330160021.Do(func() {
		cache_Control_Semigroupoid_compose__330160021 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__330160021(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__330160021
}

var cache_Control_Semigroupoid_compose__931517511 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__931517511 sync.Once

func Get_Control_Semigroupoid_compose__931517511() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__931517511.Do(func() {
		cache_Control_Semigroupoid_compose__931517511 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__931517511(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__931517511
}

var cache_Control_Semigroupoid_compose__1719608277 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1719608277 sync.Once

func Get_Control_Semigroupoid_compose__1719608277() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1719608277.Do(func() {
		cache_Control_Semigroupoid_compose__1719608277 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1719608277(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1719608277
}

var cache_Control_Semigroupoid_compose__2381855605 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2381855605 sync.Once

func Get_Control_Semigroupoid_compose__2381855605() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2381855605.Do(func() {
		cache_Control_Semigroupoid_compose__2381855605 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2381855605(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2381855605
}

var cache_Control_Semigroupoid_compose__623785531 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__623785531 sync.Once

func Get_Control_Semigroupoid_compose__623785531() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__623785531.Do(func() {
		cache_Control_Semigroupoid_compose__623785531 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__623785531(__eta_norm_1_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__623785531
}

var cache_Control_Semigroupoid_compose__3925319035 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3925319035 sync.Once

func Get_Control_Semigroupoid_compose__3925319035() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3925319035.Do(func() {
		cache_Control_Semigroupoid_compose__3925319035 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3925319035(__eta_norm_1_0_box, __eta_norm_0_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3925319035
}

var cache_Control_Semigroupoid_compose__1420521365 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1420521365 sync.Once

func Get_Control_Semigroupoid_compose__1420521365() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1420521365.Do(func() {
		cache_Control_Semigroupoid_compose__1420521365 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1420521365(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1420521365
}

var cache_Control_Semigroupoid_compose__3404027605 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3404027605 sync.Once

func Get_Control_Semigroupoid_compose__3404027605() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3404027605.Do(func() {
		cache_Control_Semigroupoid_compose__3404027605 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3404027605(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3404027605
}

var cache_Control_Semigroupoid_compose__674691509 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__674691509 sync.Once

func Get_Control_Semigroupoid_compose__674691509() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__674691509.Do(func() {
		cache_Control_Semigroupoid_compose__674691509 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__674691509(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__674691509
}

var cache_Control_Semigroupoid_compose__2621543011 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2621543011 sync.Once

func Get_Control_Semigroupoid_compose__2621543011() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2621543011.Do(func() {
		cache_Control_Semigroupoid_compose__2621543011 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2621543011(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2621543011
}

var cache_Control_Semigroupoid_compose__1650034709 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1650034709 sync.Once

func Get_Control_Semigroupoid_compose__1650034709() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1650034709.Do(func() {
		cache_Control_Semigroupoid_compose__1650034709 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1650034709(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1650034709
}

var cache_Control_Semigroupoid_compose__2328161164 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2328161164 sync.Once

func Get_Control_Semigroupoid_compose__2328161164() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2328161164.Do(func() {
		cache_Control_Semigroupoid_compose__2328161164 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2328161164(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2328161164
}

var cache_Control_Semigroupoid_compose__1111374466 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1111374466 sync.Once

func Get_Control_Semigroupoid_compose__1111374466() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1111374466.Do(func() {
		cache_Control_Semigroupoid_compose__1111374466 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1111374466(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1111374466
}

var cache_Control_Semigroupoid_compose__1993852626 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1993852626 sync.Once

func Get_Control_Semigroupoid_compose__1993852626() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1993852626.Do(func() {
		cache_Control_Semigroupoid_compose__1993852626 = gopurs_runtime.Func3(func(__eta_norm_2_unused_0_box gopurs_runtime.Value, __eta_norm_1_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1993852626(__eta_norm_2_unused_0_box, __eta_norm_1_unused_1_box, __eta_norm_0_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1993852626
}

var cache_Control_Semigroupoid_compose__992929588 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__992929588 sync.Once

func Get_Control_Semigroupoid_compose__992929588() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__992929588.Do(func() {
		cache_Control_Semigroupoid_compose__992929588 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__992929588(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__992929588
}

var cache_Control_Semigroupoid_compose__3330123079 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3330123079 sync.Once

func Get_Control_Semigroupoid_compose__3330123079() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3330123079.Do(func() {
		cache_Control_Semigroupoid_compose__3330123079 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3330123079(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3330123079
}

var cache_Control_Semigroupoid_compose__2717487349 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2717487349 sync.Once

func Get_Control_Semigroupoid_compose__2717487349() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2717487349.Do(func() {
		cache_Control_Semigroupoid_compose__2717487349 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2717487349(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2717487349
}

var cache_Control_Semigroupoid_compose__1662793621 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1662793621 sync.Once

func Get_Control_Semigroupoid_compose__1662793621() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1662793621.Do(func() {
		cache_Control_Semigroupoid_compose__1662793621 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1662793621(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1662793621
}

var cache_Control_Semigroupoid_compose__891439676 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__891439676 sync.Once

func Get_Control_Semigroupoid_compose__891439676() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__891439676.Do(func() {
		cache_Control_Semigroupoid_compose__891439676 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__891439676(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__891439676
}

var cache_Control_Semigroupoid_compose__2251324598 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2251324598 sync.Once

func Get_Control_Semigroupoid_compose__2251324598() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2251324598.Do(func() {
		cache_Control_Semigroupoid_compose__2251324598 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2251324598(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2251324598
}

var cache_Control_Semigroupoid_compose__1684629916 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1684629916 sync.Once

func Get_Control_Semigroupoid_compose__1684629916() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1684629916.Do(func() {
		cache_Control_Semigroupoid_compose__1684629916 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1684629916(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1684629916
}

var cache_Control_Semigroupoid_compose__2626215702 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2626215702 sync.Once

func Get_Control_Semigroupoid_compose__2626215702() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2626215702.Do(func() {
		cache_Control_Semigroupoid_compose__2626215702 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2626215702(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2626215702
}

var cache_Control_Semigroupoid_compose__2260825623 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2260825623 sync.Once

func Get_Control_Semigroupoid_compose__2260825623() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2260825623.Do(func() {
		cache_Control_Semigroupoid_compose__2260825623 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2260825623(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2260825623
}

var cache_Control_Semigroupoid_compose__1131435178 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1131435178 sync.Once

func Get_Control_Semigroupoid_compose__1131435178() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1131435178.Do(func() {
		cache_Control_Semigroupoid_compose__1131435178 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1131435178(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1131435178
}

var cache_Control_Semigroupoid_compose__2759654789 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2759654789 sync.Once

func Get_Control_Semigroupoid_compose__2759654789() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2759654789.Do(func() {
		cache_Control_Semigroupoid_compose__2759654789 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__2759654789(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__2759654789
}

var cache_Control_Semigroupoid_compose__3870335061 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3870335061 sync.Once

func Get_Control_Semigroupoid_compose__3870335061() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3870335061.Do(func() {
		cache_Control_Semigroupoid_compose__3870335061 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3870335061(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3870335061
}

var cache_Control_Semigroupoid_compose__17554965 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__17554965 sync.Once

func Get_Control_Semigroupoid_compose__17554965() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__17554965.Do(func() {
		cache_Control_Semigroupoid_compose__17554965 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__17554965(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__17554965
}

var cache_Control_Semigroupoid_compose__1102073438 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1102073438 sync.Once

func Get_Control_Semigroupoid_compose__1102073438() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1102073438.Do(func() {
		cache_Control_Semigroupoid_compose__1102073438 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__1102073438(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__1102073438
}

var cache_Control_Semigroupoid_compose__3918397002 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3918397002 sync.Once

func Get_Control_Semigroupoid_compose__3918397002() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3918397002.Do(func() {
		cache_Control_Semigroupoid_compose__3918397002 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3918397002(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3918397002
}

var cache_Control_Semigroupoid_compose__3798708373 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3798708373 sync.Once

func Get_Control_Semigroupoid_compose__3798708373() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3798708373.Do(func() {
		cache_Control_Semigroupoid_compose__3798708373 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3798708373(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3798708373
}

var cache_Control_Semigroupoid_compose__3874304042 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3874304042 sync.Once

func Get_Control_Semigroupoid_compose__3874304042() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3874304042.Do(func() {
		cache_Control_Semigroupoid_compose__3874304042 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3874304042(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3874304042
}

var cache_Control_Semigroupoid_compose__827489973 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__827489973 sync.Once

func Get_Control_Semigroupoid_compose__827489973() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__827489973.Do(func() {
		cache_Control_Semigroupoid_compose__827489973 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__827489973(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__827489973
}

var cache_Control_Semigroupoid_compose__474161516 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__474161516 sync.Once

func Get_Control_Semigroupoid_compose__474161516() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__474161516.Do(func() {
		cache_Control_Semigroupoid_compose__474161516 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__474161516(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__474161516
}

var cache_Control_Semigroupoid_compose__3222558869 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3222558869 sync.Once

func Get_Control_Semigroupoid_compose__3222558869() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3222558869.Do(func() {
		cache_Control_Semigroupoid_compose__3222558869 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_compose__3222558869(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Control_Semigroupoid_compose__3222558869
}

var cache_Control_Semigroupoid_composeFlipped gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped sync.Once

func Get_Control_Semigroupoid_composeFlipped() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped.Do(func() {
		cache_Control_Semigroupoid_composeFlipped = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box))
		})
	})
	return cache_Control_Semigroupoid_composeFlipped
}

var cache_Control_Semigroupoid_composeFlipped__768198773 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__768198773 sync.Once

func Get_Control_Semigroupoid_composeFlipped__768198773() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__768198773.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__768198773 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__768198773(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__768198773
}

var cache_Control_Semigroupoid_composeFlipped__2126304935 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__2126304935 sync.Once

func Get_Control_Semigroupoid_composeFlipped__2126304935() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__2126304935.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__2126304935 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__2126304935(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__2126304935
}

var cache_Control_Semigroupoid_composeFlipped__2502192415 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__2502192415 sync.Once

func Get_Control_Semigroupoid_composeFlipped__2502192415() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__2502192415.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__2502192415 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__2502192415(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__2502192415
}

var cache_Control_Semigroupoid_composeFlipped__4266205525 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__4266205525 sync.Once

func Get_Control_Semigroupoid_composeFlipped__4266205525() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__4266205525.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__4266205525 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__4266205525(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__4266205525
}

var cache_Control_Semigroupoid_composeFlipped__3966794215 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__3966794215 sync.Once

func Get_Control_Semigroupoid_composeFlipped__3966794215() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__3966794215.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__3966794215 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__3966794215(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__3966794215
}

var cache_Control_Semigroupoid_composeFlipped__325980775 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__325980775 sync.Once

func Get_Control_Semigroupoid_composeFlipped__325980775() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__325980775.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__325980775 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__325980775(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__325980775
}

var cache_Control_Semigroupoid_composeFlipped__3929334311 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__3929334311 sync.Once

func Get_Control_Semigroupoid_composeFlipped__3929334311() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__3929334311.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__3929334311 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Semigroupoid_composeFlipped__3929334311(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Semigroupoid_composeFlipped__3929334311
}

type Constructor_Control_Semigroupoid_Semigroupoid[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f00530065006d006900670072006f00750070006f00690064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f00530065006d006900670072006f00750070006f00690064005f00530065006d006900670072006f00750070006f00690064
func init() {
	gopurs_runtime.StructGetters[350442445] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "compose":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Control_Semigroupoid_Semigroupoid: " + key)
		}
	}
}

func Call_Control_Semigroupoid_Semigroupoid_dollar_Dict(x_0_loop struct {
	compose gopurs_runtime.Value
	/*gopurs:record-v1:1:0063006f006d0070006f00730065*/
}) *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] {
	var x_0 struct {
		compose gopurs_runtime.Value
		/*gopurs:record-v1:1:0063006f006d0070006f00730065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("compose", orig.compose)
	}())
}

func Call_Control_Semigroupoid_compose(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Control_Semigroupoid_compose__3872452778(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3872452778:
	for {
		if false {
			continue compose__3872452778
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Argonaut_Core_toObject(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1323742214(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1323742214:
	for {
		if false {
			continue compose__1323742214
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Argonaut_Core_toArray(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__251463616(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__251463616:
	for {
		if false {
			continue compose__251463616
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_4021906832_123048125(Rebox_Control_Semigroupoid_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__3629354669(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3629354669:
	for {
		if false {
			continue compose__3629354669
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_4021906832_123048125(Rebox_Control_Semigroupoid_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__3537576912(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3537576912:
	for {
		if false {
			continue compose__3537576912
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_4021906832_123048125(Rebox_Control_Semigroupoid_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumWeekday()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__953525693(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__953525693:
	for {
		if false {
			continue compose__953525693
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_4021906832_123048125(Rebox_Control_Semigroupoid_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumWeekday()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__1578359809(__eta_norm_2_unused_0_loop gopurs_runtime.Value, __eta_norm_1_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1578359809:
	for {
		if false {
			continue compose__1578359809
		}
		var __eta_norm_2_unused_0 gopurs_runtime.Value = __eta_norm_2_unused_0_loop
		_ = __eta_norm_2_unused_0
		var __eta_norm_1_unused_1 gopurs_runtime.Value = __eta_norm_1_unused_1_loop
		_ = __eta_norm_1_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return gopurs_runtime.Apply2(Get_Data_Newtype_unwrap(), gopurs_runtime.Value{}, gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_runExceptT(), __eta_norm_0_2))
	}
}

func Call_Control_Semigroupoid_compose__1084262850(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1084262850:
	for {
		if false {
			continue compose__1084262850
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_liftEffect(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2472072002(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2472072002:
	for {
		if false {
			continue compose__2472072002
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_liftEffect(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__635039573(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__635039573:
	for {
		if false {
			continue compose__635039573
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__817328085(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__817328085:
	for {
		if false {
			continue compose__817328085
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1053824693(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1053824693:
	for {
		if false {
			continue compose__1053824693
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__706442655(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__706442655:
	for {
		if false {
			continue compose__706442655
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_Regex_Flags_RegexFlags(), gopurs_runtime.Apply(Get_Data_String_Regex_flagsImpl(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2752544656(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2752544656:
	for {
		if false {
			continue compose__2752544656
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_fromEffectFnAff(), gopurs_runtime.Apply(Get_Database_SQLite3_beginImpl(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2473888545(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2473888545:
	for {
		if false {
			continue compose__2473888545
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_fromEffectFnAff(), gopurs_runtime.Apply(Get_Database_SQLite3_closeImpl(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__81161766(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__81161766:
	for {
		if false {
			continue compose__81161766
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_fromEffectFnAff(), gopurs_runtime.Apply(Get_Database_SQLite3_commitImpl(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__18355393(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__18355393:
	for {
		if false {
			continue compose__18355393
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_fromEffectFnAff(), gopurs_runtime.Apply(Get_Database_SQLite3_rollbackImpl(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__297303379(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__297303379:
	for {
		if false {
			continue compose__297303379
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_void(), gopurs_runtime.Apply(Get_Effect_Aff_launchAff(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1680806490(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1680806490:
	for {
		if false {
			continue compose__1680806490
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Function_go__const(), gopurs_runtime.Apply(Get_Effect_Aff_liftEffect(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__262436915(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__262436915:
	for {
		if false {
			continue compose__262436915
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_Canceler(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2920693720(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2920693720:
	for {
		if false {
			continue compose__2920693720
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Left(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2990151384(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2990151384:
	for {
		if false {
			continue compose__2990151384
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Left(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3715450552(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3715450552:
	for {
		if false {
			continue compose__3715450552
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Left(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2907274104(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2907274104:
	for {
		if false {
			continue compose__2907274104
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Left(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__13732245(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__13732245:
	for {
		if false {
			continue compose__13732245
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__447393272(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__447393272:
	for {
		if false {
			continue compose__447393272
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Left(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__985198819(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__985198819:
	for {
		if false {
			continue compose__985198819
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Right(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1195436277(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1195436277:
	for {
		if false {
			continue compose__1195436277
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2013581781(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2013581781:
	for {
		if false {
			continue compose__2013581781
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3089340516(__eta_norm_2_unused_0_loop gopurs_runtime.Value, __eta_norm_1_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3089340516:
	for {
		if false {
			continue compose__3089340516
		}
		var __eta_norm_2_unused_0 gopurs_runtime.Value = __eta_norm_2_unused_0_loop
		_ = __eta_norm_2_unused_0
		var __eta_norm_1_unused_1 gopurs_runtime.Value = __eta_norm_1_unused_1_loop
		_ = __eta_norm_1_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return gopurs_runtime.Apply(Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Apply2(Get_Data_NonEmpty_singleton(), gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_2173899317_3706288089(Rebox_Control_Semigroupoid_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList()))))}, __eta_norm_0_2))
	}
}

func Call_Control_Semigroupoid_compose__3958277077(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3958277077:
	for {
		if false {
			continue compose__3958277077
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3383039189(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3383039189:
	for {
		if false {
			continue compose__3383039189
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2488896981(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2488896981:
	for {
		if false {
			continue compose__2488896981
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3939066901(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3939066901:
	for {
		if false {
			continue compose__3939066901
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2813039818(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2813039818:
	for {
		if false {
			continue compose__2813039818
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Control_Applicative_pure(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, x_3)
		})
	}
}

func Call_Control_Semigroupoid_compose__4004832778(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__4004832778:
	for {
		if false {
			continue compose__4004832778
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Control_Applicative_pure(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, x_3)
		})
	}
}

func Call_Control_Semigroupoid_compose__2889719291(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2889719291:
	for {
		if false {
			continue compose__2889719291
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1982962581(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1982962581:
	for {
		if false {
			continue compose__1982962581
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3606255701(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3606255701:
	for {
		if false {
			continue compose__3606255701
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2491846907(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2491846907:
	for {
		if false {
			continue compose__2491846907
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__580204084(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__580204084:
	for {
		if false {
			continue compose__580204084
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_fromCharArray(), gopurs_runtime.Apply(Get_Data_Array_NonEmpty_toArray(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3042204923(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3042204923:
	for {
		if false {
			continue compose__3042204923
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__323228359(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__323228359:
	for {
		if false {
			continue compose__323228359
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_fromCodePointArray(), gopurs_runtime.Apply(Get_Data_Array_NonEmpty_toArray(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__788005840(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__788005840:
	for {
		if false {
			continue compose__788005840
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_encodeString(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1772354132(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1772354132:
	for {
		if false {
			continue compose__1772354132
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_toNonEmptyString(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__50666552(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__50666552:
	for {
		if false {
			continue compose__50666552
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_countPrefix(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2152657429(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2152657429:
	for {
		if false {
			continue compose__2152657429
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2655475989(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2655475989:
	for {
		if false {
			continue compose__2655475989
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2593021333(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2593021333:
	for {
		if false {
			continue compose__2593021333
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2246386360(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2246386360:
	for {
		if false {
			continue compose__2246386360
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_liftS(), gopurs_runtime.Apply(Get_Data_String_CodePoints_countPrefix(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3326469365(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3326469365:
	for {
		if false {
			continue compose__3326469365
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3055166288(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3055166288:
	for {
		if false {
			continue compose__3055166288
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__4060784579(), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottom__2335111300()), gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_top__2335111300()))
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__1433444296(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1433444296:
	for {
		if false {
			continue compose__1433444296
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2650337094(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2650337094:
	for {
		if false {
			continue compose__2650337094
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_charAt(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__661099861(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__661099861:
	for {
		if false {
			continue compose__661099861
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2074229983(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2074229983:
	for {
		if false {
			continue compose__2074229983
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_liftS(), gopurs_runtime.Apply(Get_Data_String_CodePoints_codePointAt(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3307342101(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3307342101:
	for {
		if false {
			continue compose__3307342101
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__419544780(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__419544780:
	for {
		if false {
			continue compose__419544780
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__1934312767(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1934312767:
	for {
		if false {
			continue compose__1934312767
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__1157135857(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1157135857:
	for {
		if false {
			continue compose__1157135857
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__1739002828(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1739002828:
	for {
		if false {
			continue compose__1739002828
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__3898604607(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3898604607:
	for {
		if false {
			continue compose__3898604607
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__585867329(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__585867329:
	for {
		if false {
			continue compose__585867329
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_toEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__2240476445(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2240476445:
	for {
		if false {
			continue compose__2240476445
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Number_Format_Exponential(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1284278492(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1284278492:
	for {
		if false {
			continue compose__1284278492
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Number_Format_Fixed(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3320522428(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3320522428:
	for {
		if false {
			continue compose__3320522428
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Number_Format_Precision(), gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1048546465(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1048546465:
	for {
		if false {
			continue compose__1048546465
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__2993739570(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2993739570:
	for {
		if false {
			continue compose__2993739570
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__2865838076(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2865838076:
	for {
		if false {
			continue compose__2865838076
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__13873697(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__13873697:
	for {
		if false {
			continue compose__13873697
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__4254564146(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__4254564146:
	for {
		if false {
			continue compose__4254564146
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__1304014924(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1304014924:
	for {
		if false {
			continue compose__1304014924
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Enum_fromEnum(), gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_1306125126_123048125(Rebox_Control_Semigroupoid_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond()))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__494499605(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__494499605:
	for {
		if false {
			continue compose__494499605
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3460990597(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3460990597:
	for {
		if false {
			continue compose__3460990597
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromNumber(), gopurs_runtime.Apply(Get_Data_Int_toNumber(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__178943130(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__178943130:
	for {
		if false {
			continue compose__178943130
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_encodeString(), gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__100839156(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__100839156:
	for {
		if false {
			continue compose__100839156
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_toNonEmptyString(), gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1866182613(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1866182613:
	for {
		if false {
			continue compose__1866182613
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2640382690(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2640382690:
	for {
		if false {
			continue compose__2640382690
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Int_fromNumber(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__529030880(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__529030880:
	for {
		if false {
			continue compose__529030880
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Time_Duration_over(), Get_Data_Time_Duration_Milliseconds(), Get_Data_Time_Duration_negate())
		_ = __local_var_2_0
		// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Time_Duration_fromDuration(), gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_3362292604_1457896548(Rebox_Control_Semigroupoid_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds()))))})
		_ = __local_var_3_1
		return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
		})
	}
}

func Call_Control_Semigroupoid_compose__532200073(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__532200073:
	for {
		if false {
			continue compose__532200073
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Interval_Duration_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Interval_Duration_Second().IntVal)), UnsafePtr: nil})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__eta_norm_0_1, x_3))
		})
	}
}

func Call_Control_Semigroupoid_compose__880864056(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__880864056:
	for {
		if false {
			continue compose__880864056
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Time_Duration_toDuration(), gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_3362292604_1457896548(Rebox_Control_Semigroupoid_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds()))))})
		_ = __local_var_2_0
		// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=Any
		__local_var_3_1 := gopurs_runtime.Apply2(Get_Data_Time_Duration_compose1__1079398611(), gopurs_runtime.Apply2(Get_Data_Time_Duration_over(), Get_Data_Time_Duration_Milliseconds(), Get_Data_Time_Duration_negate()), gopurs_runtime.Apply(Get_Data_Time_Duration_fromDuration(), gopurs_runtime.Value{Type: 9, IntVal: 487663984, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_3362292604_1457896548(Rebox_Control_Semigroupoid_1457896548_3362292604(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](Get_Data_Time_Duration_durationMilliseconds()))))}))
		_ = __local_var_3_1
		return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
		})
	}
}

func Call_Control_Semigroupoid_compose__172418801(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__172418801:
	for {
		if false {
			continue compose__172418801
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_ceil(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1247429674(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1247429674:
	for {
		if false {
			continue compose__1247429674
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_floor(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__446754736(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__446754736:
	for {
		if false {
			continue compose__446754736
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_round(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__281216780(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__281216780:
	for {
		if false {
			continue compose__281216780
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_trunc(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__204472163(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__204472163:
	for {
		if false {
			continue compose__204472163
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Right(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__4153383075(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__4153383075:
	for {
		if false {
			continue compose__4153383075
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Either_Right(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__330160021(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__330160021:
	for {
		if false {
			continue compose__330160021
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__931517511(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__931517511:
	for {
		if false {
			continue compose__931517511
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(Get_Data_Argonaut_Parser_jsonParser(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1719608277(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1719608277:
	for {
		if false {
			continue compose__1719608277
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2381855605(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2381855605:
	for {
		if false {
			continue compose__2381855605
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__623785531(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__623785531:
	for {
		if false {
			continue compose__623785531
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3925319035(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3925319035:
	for {
		if false {
			continue compose__3925319035
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta_norm_1_0, gopurs_runtime.Apply(__eta_norm_0_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1420521365(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1420521365:
	for {
		if false {
			continue compose__1420521365
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__3404027605(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3404027605:
	for {
		if false {
			continue compose__3404027605
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__674691509(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__674691509:
	for {
		if false {
			continue compose__674691509
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2621543011(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2621543011:
	for {
		if false {
			continue compose__2621543011
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_fromEffectFnAff(), gopurs_runtime.Apply(Get_Database_SQLite3_openImpl(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1650034709(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1650034709:
	for {
		if false {
			continue compose__1650034709
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__2328161164(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2328161164:
	for {
		if false {
			continue compose__2328161164
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Exception_Unsafe_unsafeThrowException(), gopurs_runtime.Apply(Get_Effect_Exception_error(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1111374466(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1111374466:
	for {
		if false {
			continue compose__1111374466
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Effect_Exception_throwException(), gopurs_runtime.Apply(Get_Effect_Exception_error(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1993852626(__eta_norm_2_unused_0_loop gopurs_runtime.Value, __eta_norm_1_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1993852626:
	for {
		if false {
			continue compose__1993852626
		}
		var __eta_norm_2_unused_0 gopurs_runtime.Value = __eta_norm_2_unused_0_loop
		_ = __eta_norm_2_unused_0
		var __eta_norm_1_unused_1 gopurs_runtime.Value = __eta_norm_1_unused_1_loop
		_ = __eta_norm_1_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return gopurs_runtime.Apply(Get_Effect_Exception_throwException(), gopurs_runtime.Apply(Get_Effect_Exception_error(), __eta_norm_0_2))
	}
}

func Call_Control_Semigroupoid_compose__992929588(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__992929588:
	for {
		if false {
			continue compose__992929588
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_fromArray(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_toCharArray(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3330123079(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3330123079:
	for {
		if false {
			continue compose__3330123079
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_fromArray(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_toCodePointArray(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2717487349(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2717487349:
	for {
		if false {
			continue compose__2717487349
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1662793621(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1662793621:
	for {
		if false {
			continue compose__1662793621
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__891439676(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__891439676:
	for {
		if false {
			continue compose__891439676
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_liftS(), gopurs_runtime.Apply(Get_Data_String_CodePoints_indexOf(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2251324598(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2251324598:
	for {
		if false {
			continue compose__2251324598
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_liftS(), gopurs_runtime.Apply(Get_Data_String_CodePoints_lastIndexOf(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1684629916(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1684629916:
	for {
		if false {
			continue compose__1684629916
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2626215702(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2626215702:
	for {
		if false {
			continue compose__2626215702
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_lastIndexOf(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2260825623(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2260825623:
	for {
		if false {
			continue compose__2260825623
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_Internal_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_contains(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__1131435178(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1131435178:
	for {
		if false {
			continue compose__1131435178
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_Internal_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_endsWith(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__2759654789(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__2759654789:
	for {
		if false {
			continue compose__2759654789
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_NonEmpty_Internal_liftS(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_startsWith(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3870335061(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3870335061:
	for {
		if false {
			continue compose__3870335061
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__17554965(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__17554965:
	for {
		if false {
			continue compose__17554965
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_compose__1102073438(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__1102073438:
	for {
		if false {
			continue compose__1102073438
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromString(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_Internal_toString(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3918397002(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3918397002:
	for {
		if false {
			continue compose__3918397002
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_CodePoints_length(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3798708373(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3798708373:
	for {
		if false {
			continue compose__3798708373
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3874304042(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3874304042:
	for {
		if false {
			continue compose__3874304042
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__827489973(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__827489973:
	for {
		if false {
			continue compose__827489973
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_CodeUnits_toChar(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__474161516(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__474161516:
	for {
		if false {
			continue compose__474161516
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(Get_Data_String_CodeUnits_toCharArray(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(), x_2))
	}
}

func Call_Control_Semigroupoid_compose__3222558869(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
compose__3222558869:
	for {
		if false {
			continue compose__3222558869
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Control_Semigroupoid_composeFlipped(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
	_ = dictSemigroupoid_0
	// TAST (Let): compose1_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar a$scope16) [(TypeVar c$scope18), (TypeVar d$scope19)]), (TypeApp (TypeVar a$scope16) [(TypeVar b$scope17), (TypeVar c$scope18)])] (TypeApp (TypeVar a$scope16) [(TypeVar b$scope17), (TypeVar d$scope19)]))
	compose1_1_0 := Call_Control_Semigroupoid_compose(dictSemigroupoid_0)
	_ = compose1_1_0
	return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(compose1_1_0, g_3, f_2)
	})
}

func Call_Control_Semigroupoid_composeFlipped__768198773(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__768198773:
	for {
		if false {
			continue composeFlipped__768198773
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta1_1, __eta0_0)
	}
}

func Call_Control_Semigroupoid_composeFlipped__2126304935(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__2126304935:
	for {
		if false {
			continue composeFlipped__2126304935
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_stringify(), gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Class_encodeJson(), gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_4242125399_2831479053(Rebox_Control_Semigroupoid_2831479053_4242125399(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Class_encodeJsonArray(), gopurs_runtime.Apply(Get_Main_encodeRecord(), gopurs_runtime.Value{}))))))}))
	}
}

func Call_Control_Semigroupoid_composeFlipped__2502192415(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__2502192415:
	for {
		if false {
			continue composeFlipped__2502192415
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_CodePoints_CodePoint(), Get_Data_String_CodePoints_fromEnum())
	}
}

func Call_Control_Semigroupoid_composeFlipped__4266205525(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__4266205525:
	for {
		if false {
			continue composeFlipped__4266205525
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta1_1, __eta0_0)
	}
}

func Call_Control_Semigroupoid_composeFlipped__3966794215(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__3966794215:
	for {
		if false {
			continue composeFlipped__3966794215
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_stringify(), __eta_norm_1_0)
	}
}

func Call_Control_Semigroupoid_composeFlipped__325980775(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__325980775:
	for {
		if false {
			continue composeFlipped__325980775
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_stringify(), gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Class_encodeJson(), gopurs_runtime.Value{Type: 9, IntVal: 3228574649, UnsafePtr: unsafe.Pointer(Rebox_Control_Semigroupoid_2037062001_2831479053(Rebox_Control_Semigroupoid_2831479053_2037062001(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_encodeRecord(), gopurs_runtime.Value{})))))}))
	}
}

func Call_Control_Semigroupoid_composeFlipped__3929334311(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__3929334311:
	for {
		if false {
			continue composeFlipped__3929334311
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Argonaut_Core_stringify(), __eta_norm_1_0)
	}
}

func Rebox_Control_Semigroupoid_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[int64])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_123048125_4021906832(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[uint32])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_1457896548_3362292604(in *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) *Constructor_Data_Time_Duration_Duration[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Time_Duration_Duration[float64])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_2037062001_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_2173899317_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_2831479053_2037062001(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
		completed bool
		id        string
		title     string
		/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_2831479053_4242125399(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[[]struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[[]struct {
		completed bool
		id        string
		title     string
		/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_3362292604_1457896548(in *Constructor_Data_Time_Duration_Duration[float64]) *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_3706288089_2173899317(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Semigroupoid_4242125399_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[[]struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}
