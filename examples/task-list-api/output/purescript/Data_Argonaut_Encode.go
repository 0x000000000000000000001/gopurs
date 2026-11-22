package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Encode_composeFlipped gopurs_runtime.Value
var once_Data_Argonaut_Encode_composeFlipped sync.Once

func Get_Data_Argonaut_Encode_composeFlipped() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_composeFlipped.Do(func() {
		cache_Data_Argonaut_Encode_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Argonaut_Encode_composeFlipped
}

var cache_Data_Argonaut_Encode_composeFlipped__2837117588 gopurs_runtime.Value
var once_Data_Argonaut_Encode_composeFlipped__2837117588 sync.Once

func Get_Data_Argonaut_Encode_composeFlipped__2837117588() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_composeFlipped__2837117588.Do(func() {
		cache_Data_Argonaut_Encode_composeFlipped__2837117588 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_composeFlipped__2837117588(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Argonaut_Encode_composeFlipped__2837117588
}

var cache_Data_Argonaut_Encode_composeFlipped__2170759508 gopurs_runtime.Value
var once_Data_Argonaut_Encode_composeFlipped__2170759508 sync.Once

func Get_Data_Argonaut_Encode_composeFlipped__2170759508() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_composeFlipped__2170759508.Do(func() {
		cache_Data_Argonaut_Encode_composeFlipped__2170759508 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_composeFlipped__2170759508(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Argonaut_Encode_composeFlipped__2170759508
}

var cache_Data_Argonaut_Encode_composeFlipped__1050801172 gopurs_runtime.Value
var once_Data_Argonaut_Encode_composeFlipped__1050801172 sync.Once

func Get_Data_Argonaut_Encode_composeFlipped__1050801172() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_composeFlipped__1050801172.Do(func() {
		cache_Data_Argonaut_Encode_composeFlipped__1050801172 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_composeFlipped__1050801172(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Argonaut_Encode_composeFlipped__1050801172
}

var cache_Data_Argonaut_Encode_composeFlipped__3683799540 gopurs_runtime.Value
var once_Data_Argonaut_Encode_composeFlipped__3683799540 sync.Once

func Get_Data_Argonaut_Encode_composeFlipped__3683799540() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_composeFlipped__3683799540.Do(func() {
		cache_Data_Argonaut_Encode_composeFlipped__3683799540 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_composeFlipped__3683799540(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Data_Argonaut_Encode_composeFlipped__3683799540
}

var cache_Data_Argonaut_Encode_toJsonString gopurs_runtime.Value
var once_Data_Argonaut_Encode_toJsonString sync.Once

func Get_Data_Argonaut_Encode_toJsonString() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_toJsonString.Do(func() {
		cache_Data_Argonaut_Encode_toJsonString = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_toJsonString(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Encode_toJsonString
}

var cache_Data_Argonaut_Encode_toJsonString__2078790066 gopurs_runtime.Value
var once_Data_Argonaut_Encode_toJsonString__2078790066 sync.Once

func Get_Data_Argonaut_Encode_toJsonString__2078790066() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_toJsonString__2078790066.Do(func() {
		cache_Data_Argonaut_Encode_toJsonString__2078790066 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Argonaut_Encode_toJsonString__2078790066(func() []struct {
				completed bool
				id        string
				title     string
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
			} {
				arr := *(*[]gopurs_runtime.Value)(__eta0_0_box.UnsafePtr)
				unboxed := make([]struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				}, len(arr))
				for i, v := range arr {
					unboxed[i] = func() struct {
						completed bool
						id        string
						title     string
						/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
					} {
						orig := v
						_ = orig
						clone := struct {
							completed bool
							id        string
							title     string
							/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
						}{}
						clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
						clone.id = gopurs_runtime.RecordGet(orig, "id").StrVal()
						clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
						return clone
					}()
				}
				return unboxed
			}()))
		})
	})
	return cache_Data_Argonaut_Encode_toJsonString__2078790066
}

var cache_Data_Argonaut_Encode_toJsonString__938966123 gopurs_runtime.Value
var once_Data_Argonaut_Encode_toJsonString__938966123 sync.Once

func Get_Data_Argonaut_Encode_toJsonString__938966123() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_toJsonString__938966123.Do(func() {
		cache_Data_Argonaut_Encode_toJsonString__938966123 = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_toJsonString__938966123(Rebox_Data_Argonaut_Encode_2831479053_614949298(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box)))
		})
	})
	return cache_Data_Argonaut_Encode_toJsonString__938966123
}

var cache_Data_Argonaut_Encode_toJsonString__2285955122 gopurs_runtime.Value
var once_Data_Argonaut_Encode_toJsonString__2285955122 sync.Once

func Get_Data_Argonaut_Encode_toJsonString__2285955122() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_toJsonString__2285955122.Do(func() {
		cache_Data_Argonaut_Encode_toJsonString__2285955122 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Argonaut_Encode_toJsonString__2285955122(func() struct {
				completed bool
				id        string
				title     string
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
			} {
				orig := __eta0_0_box
				_ = orig
				clone := struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				}{}
				clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
				clone.id = gopurs_runtime.RecordGet(orig, "id").StrVal()
				clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
				return clone
			}()))
		})
	})
	return cache_Data_Argonaut_Encode_toJsonString__2285955122
}

var cache_Data_Argonaut_Encode_toJsonString__1304219019 gopurs_runtime.Value
var once_Data_Argonaut_Encode_toJsonString__1304219019 sync.Once

func Get_Data_Argonaut_Encode_toJsonString__1304219019() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_toJsonString__1304219019.Do(func() {
		cache_Data_Argonaut_Encode_toJsonString__1304219019 = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_toJsonString__1304219019(Rebox_Data_Argonaut_Encode_2831479053_1331040862(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box)))
		})
	})
	return cache_Data_Argonaut_Encode_toJsonString__1304219019
}

func Call_Data_Argonaut_Encode_composeFlipped__2837117588(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__2837117588:
	for {
		if false {
			continue composeFlipped__2837117588
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Data_Argonaut_Encode_4242125399_2831479053(Rebox_Data_Argonaut_Encode_2831479053_4242125399(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Call_Data_Argonaut_Encode_Class_encodeJsonArray(gopurs_runtime.Apply(Get_Main_encodeRecord(), gopurs_runtime.Value{})))))), Get_Data_Argonaut_Core_stringify())
	}
}

func Call_Data_Argonaut_Encode_composeFlipped__2170759508(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__2170759508:
	for {
		if false {
			continue composeFlipped__2170759508
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Argonaut_Core_stringify())
	}
}

func Call_Data_Argonaut_Encode_composeFlipped__1050801172(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__1050801172:
	for {
		if false {
			continue composeFlipped__1050801172
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Data_Argonaut_Encode_2037062001_2831479053(Rebox_Data_Argonaut_Encode_2831479053_2037062001(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_encodeRecord(), gopurs_runtime.Value{}))))), Get_Data_Argonaut_Core_stringify())
	}
}

func Call_Data_Argonaut_Encode_composeFlipped__3683799540(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeFlipped__3683799540:
	for {
		if false {
			continue composeFlipped__3683799540
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), __eta_norm_1_0, Get_Data_Argonaut_Core_stringify())
	}
}

func Call_Data_Argonaut_Encode_toJsonString(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] = dictEncodeJson_0_loop
	_ = dictEncodeJson_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(dictEncodeJson_0), Get_Data_Argonaut_Core_stringify())
}

func Call_Data_Argonaut_Encode_toJsonString__2078790066(__eta0_0_loop []struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}) string {
toJsonString__2078790066:
	for {
		if false {
			continue toJsonString__2078790066
		}
		var __eta0_0 []struct {
			completed bool
			id        string
			title     string
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
		} = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(Call_Data_Argonaut_Encode_composeFlipped__2837117588(Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Data_Argonaut_Encode_4242125399_2831479053(Rebox_Data_Argonaut_Encode_2831479053_4242125399(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](Call_Data_Argonaut_Encode_Class_encodeJsonArray(gopurs_runtime.Apply(Get_Main_encodeRecord(), gopurs_runtime.Value{})))))), Get_Data_Argonaut_Core_stringify()), func() gopurs_runtime.Value {
			arr := __eta0_0
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					orig := v
					_ = orig
					return gopurs_runtime.RecordDict3("completed", "id", "title", gopurs_runtime.Bool(orig.completed), gopurs_runtime.Str(orig.id), gopurs_runtime.Str(orig.title))
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()).StrVal()
	}
}

func Call_Data_Argonaut_Encode_toJsonString__938966123(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	error string
	/*gopurs:record-v1:1:006500720072006f0072*/
}]) gopurs_runtime.Value {
toJsonString__938966123:
	for {
		if false {
			continue toJsonString__938966123
		}
		var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
			error string
			/*gopurs:record-v1:1:006500720072006f0072*/
		}] = dictEncodeJson_0_loop
		_ = dictEncodeJson_0
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Data_Argonaut_Encode_614949298_2831479053(dictEncodeJson_0)), Get_Data_Argonaut_Core_stringify())
	}
}

func Call_Data_Argonaut_Encode_toJsonString__2285955122(__eta0_0_loop struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}) string {
toJsonString__2285955122:
	for {
		if false {
			continue toJsonString__2285955122
		}
		var __eta0_0 struct {
			completed bool
			id        string
			title     string
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
		} = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply3(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Data_Argonaut_Encode_2037062001_2831479053(Rebox_Data_Argonaut_Encode_2831479053_2037062001(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_encodeRecord(), gopurs_runtime.Value{}))))), Get_Data_Argonaut_Core_stringify(), func() gopurs_runtime.Value {
			orig := __eta0_0
			_ = orig
			return gopurs_runtime.RecordDict3("completed", "id", "title", gopurs_runtime.Bool(orig.completed), gopurs_runtime.Str(orig.id), gopurs_runtime.Str(orig.title))
		}()).StrVal()
	}
}

func Call_Data_Argonaut_Encode_toJsonString__1304219019(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	status string
	/*gopurs:record-v1:1:007300740061007400750073*/
}]) gopurs_runtime.Value {
toJsonString__1304219019:
	for {
		if false {
			continue toJsonString__1304219019
		}
		var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
			status string
			/*gopurs:record-v1:1:007300740061007400750073*/
		}] = dictEncodeJson_0_loop
		_ = dictEncodeJson_0
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Argonaut_Encode_Class_encodeJson(Rebox_Data_Argonaut_Encode_1331040862_2831479053(dictEncodeJson_0)), Get_Data_Argonaut_Core_stringify())
	}
}

func Rebox_Data_Argonaut_Encode_1331040862_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	status string
	/*gopurs:record-v1:1:007300740061007400750073*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Encode_2037062001_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
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

func Rebox_Data_Argonaut_Encode_2831479053_1331040862(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	status string
	/*gopurs:record-v1:1:007300740061007400750073*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
		status string
		/*gopurs:record-v1:1:007300740061007400750073*/
	}])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Encode_2831479053_2037062001(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
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

func Rebox_Data_Argonaut_Encode_2831479053_4242125399(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[[]struct {
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

func Rebox_Data_Argonaut_Encode_2831479053_614949298(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	error string
	/*gopurs:record-v1:1:006500720072006f0072*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
		error string
		/*gopurs:record-v1:1:006500720072006f0072*/
	}])(unsafe.Pointer(in))
}

func Rebox_Data_Argonaut_Encode_4242125399_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[[]struct {
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

func Rebox_Data_Argonaut_Encode_614949298_2831479053(in *Constructor_Data_Argonaut_Encode_Class_EncodeJson[struct {
	error string
	/*gopurs:record-v1:1:006500720072006f0072*/
}]) *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}
