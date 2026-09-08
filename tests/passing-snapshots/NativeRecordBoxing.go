package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_makeEntry gopurs_runtime.Value
var once_Main_makeEntry sync.Once

func Get_Main_makeEntry() gopurs_runtime.Value {
	once_Main_makeEntry.Do(func() {
		cache_Main_makeEntry = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, label_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeEntry(count_0_box.IntVal, label_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"count", "label"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label)})
			}()
		})
	})
	return cache_Main_makeEntry
}

var cache_Main_describeEntry gopurs_runtime.Value
var once_Main_describeEntry sync.Once

func Get_Main_describeEntry() gopurs_runtime.Value {
	once_Main_describeEntry.Do(func() {
		cache_Main_describeEntry = gopurs_runtime.Func(func(entry_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_describeEntry(func() struct {
				count int64
				label string
			} {
				orig := entry_0_box
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}()))
		})
	})
	return cache_Main_describeEntry
}

var cache_Main_consumeEntry gopurs_runtime.Value
var once_Main_consumeEntry sync.Once

func Get_Main_consumeEntry() gopurs_runtime.Value {
	once_Main_consumeEntry.Do(func() {
		cache_Main_consumeEntry = gopurs_runtime.Func2(func(consume_0_box gopurs_runtime.Value, entry_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_consumeEntry(consume_0_box, func() struct {
				count int64
				label string
			} {
				orig := entry_1_box
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}()))
		})
	})
	return cache_Main_consumeEntry
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, actual_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(label_0_box.StrVal(), expected_1_box.StrVal(), actual_2_box.StrVal())
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Main_describeEntry())
			_ = __local_var_0_0
			consumerRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = consumerRef_1_1
			consume_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), consumerRef_1_1), gopurs_runtime.Value{})
			_ = consume_2_2
			originalRef_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := Call_Main_makeEntry(int64(5), "alpha")
				_ = orig
				return gopurs_runtime.RecordDict([]string{"count", "label"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label)})
			}()), gopurs_runtime.Value{})
			_ = originalRef_3_3
			original_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), originalRef_3_3), gopurs_runtime.Value{})
			_ = original_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Main_check("original", "5:alpha", Call_Main_consumeEntry(consume_2_2, func() struct {
				count int64
				label string
			} {
				orig := original_4_4
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_5_5
			// TAST (Let): changedCount_6_6 shape=Other bindingType=(Record (Row [count: Int, label: String] Any))
			changedCount_6_6 := func() struct {
				count int64
				label string
			} {
				orig := gopurs_runtime.RecordUpdate1(original_4_4, "count", gopurs_runtime.Int(int64(-7)))
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}()
			_ = changedCount_6_6
			countRef_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := changedCount_6_6
				_ = orig
				return gopurs_runtime.RecordDict([]string{"count", "label"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label)})
			}()), gopurs_runtime.Value{})
			_ = countRef_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_check("count updated", "-7:alpha", Call_Main_consumeEntry(consume_2_2, changedCount_6_6)), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_check("original after count update", "5:alpha", Call_Main_consumeEntry(consume_2_2, func() struct {
				count int64
				label string
			} {
				orig := original_4_4
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_9_9
			// TAST (Let): changedLabel_10_10 shape=Other bindingType=(Record (Row [label: String, count: Int] Any))
			changedLabel_10_10 := func() struct {
				count int64
				label string
			} {
				clone := changedCount_6_6
				clone.label = "beta"
				return clone
			}()
			_ = changedLabel_10_10
			labelRef_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := changedLabel_10_10
				_ = orig
				return gopurs_runtime.RecordDict([]string{"count", "label"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label)})
			}()), gopurs_runtime.Value{})
			_ = labelRef_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_check("label updated", "-7:beta", Call_Main_consumeEntry(consume_2_2, changedLabel_10_10)), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_check("count version after label update", "-7:alpha", Call_Main_consumeEntry(consume_2_2, changedCount_6_6)), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Main_check("both fields updated", "0:", Call_Main_consumeEntry(consume_2_2, func() struct {
				count int64
				label string
			} {
				clone := changedLabel_10_10
				clone.count = int64(0)
				clone.label = ""
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Main_check("label version after both updates", "-7:beta", Call_Main_consumeEntry(consume_2_2, changedLabel_10_10)), gopurs_runtime.Value{})
			_ = __local_var_15_15
			savedOriginal_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), originalRef_3_3), gopurs_runtime.Value{})
			_ = savedOriginal_16_16
			savedCount_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), countRef_7_7), gopurs_runtime.Value{})
			_ = savedCount_17_17
			savedLabel_18_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), labelRef_11_11), gopurs_runtime.Value{})
			_ = savedLabel_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_check("retained original", "5:alpha", Call_Main_consumeEntry(consume_2_2, func() struct {
				count int64
				label string
			} {
				orig := savedOriginal_16_16
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_check("retained count version", "-7:alpha", Call_Main_consumeEntry(consume_2_2, func() struct {
				count int64
				label string
			} {
				orig := savedCount_17_17
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Main_check("retained label version", "-7:beta", Call_Main_consumeEntry(consume_2_2, func() struct {
				count int64
				label string
			} {
				orig := savedLabel_18_18
				_ = orig
				clone := struct {
					count int64
					label string
				}{}
				clone.count = gopurs_runtime.RecordGet(orig, "count").IntVal
				clone.label = gopurs_runtime.RecordGet(orig, "label").StrVal()
				return clone
			}())), gopurs_runtime.Value{})
			_ = __local_var_21_21
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_makeEntry(count_0_loop int64, label_1_loop string) struct {
	count int64
	label string
} {
	var count_0 int64 = count_0_loop
	_ = count_0
	var label_1 string = label_1_loop
	_ = label_1
	return struct {
		count int64
		label string
	}{count_0, label_1}
}

func Call_Main_describeEntry(entry_0_loop struct {
	count int64
	label string
}) string {
	var entry_0 struct {
		count int64
		label string
	} = entry_0_loop
	_ = entry_0
	return ((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(entry_0.count)).StrVal()) + (":")) + (entry_0.label)
}

func Call_Main_consumeEntry(consume_0_loop gopurs_runtime.Value, entry_1_loop struct {
	count int64
	label string
}) string {
	var consume_0 gopurs_runtime.Value = consume_0_loop
	_ = consume_0
	var entry_1 struct {
		count int64
		label string
	} = entry_1_loop
	_ = entry_1
	return gopurs_runtime.Apply(consume_0, func() gopurs_runtime.Value {
		orig := entry_1
		_ = orig
		return gopurs_runtime.RecordDict([]string{"count", "label"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.count), gopurs_runtime.Str(orig.label)})
	}()).StrVal()
}

func Call_Main_check(label_0_loop string, expected_1_loop string, actual_2_loop string) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 string = expected_1_loop
	_ = expected_1
	var actual_2 string = actual_2_loop
	_ = actual_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(actual_2), gopurs_runtime.Str(expected_1)))
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(((label_0)+(": "))+(actual_2))), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
