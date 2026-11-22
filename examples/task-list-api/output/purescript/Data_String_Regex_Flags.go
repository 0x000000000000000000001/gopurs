package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Regex_Flags_go__append gopurs_runtime.Value
var once_Data_String_Regex_Flags_go__append sync.Once

func Get_Data_String_Regex_Flags_go__append() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_go__append.Do(func() {
		cache_Data_String_Regex_Flags_go__append = Get_Data_Semigroup_append__4083744612()
	})
	return cache_Data_String_Regex_Flags_go__append
}

var cache_Data_String_Regex_Flags_voidLeft gopurs_runtime.Value
var once_Data_String_Regex_Flags_voidLeft sync.Once

func Get_Data_String_Regex_Flags_voidLeft() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_voidLeft.Do(func() {
		cache_Data_String_Regex_Flags_voidLeft = Get_Data_Functor_voidLeft__3176963436()
	})
	return cache_Data_String_Regex_Flags_voidLeft
}

var cache_Data_String_Regex_Flags_guard gopurs_runtime.Value
var once_Data_String_Regex_Flags_guard sync.Once

func Get_Data_String_Regex_Flags_guard() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_guard.Do(func() {
		cache_Data_String_Regex_Flags_guard = Get_Control_Alternative_guard__3323555686()
	})
	return cache_Data_String_Regex_Flags_guard
}

var cache_Data_String_Regex_Flags_eq gopurs_runtime.Value
var once_Data_String_Regex_Flags_eq sync.Once

func Get_Data_String_Regex_Flags_eq() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_eq.Do(func() {
		cache_Data_String_Regex_Flags_eq = Get_Data_Eq_eq__121813084()
	})
	return cache_Data_String_Regex_Flags_eq
}

var cache_Data_String_Regex_Flags_append1 gopurs_runtime.Value
var once_Data_String_Regex_Flags_append1 sync.Once

func Get_Data_String_Regex_Flags_append1() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_append1.Do(func() {
		cache_Data_String_Regex_Flags_append1 = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_Regex_Flags_append1
}

var cache_Data_String_Regex_Flags_disj gopurs_runtime.Value
var once_Data_String_Regex_Flags_disj sync.Once

func Get_Data_String_Regex_Flags_disj() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_disj.Do(func() {
		cache_Data_String_Regex_Flags_disj = Get_Data_HeytingAlgebra_disj__2421783737()
	})
	return cache_Data_String_Regex_Flags_disj
}

var cache_Data_String_Regex_Flags_RegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_RegexFlags sync.Once

func Get_Data_String_Regex_Flags_RegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_RegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_RegexFlags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_Regex_Flags_RegexFlags(func() struct {
					dotAll     bool
					global     bool
					ignoreCase bool
					multiline  bool
					sticky     bool
					unicode    bool
					/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						dotAll     bool
						global     bool
						ignoreCase bool
						multiline  bool
						sticky     bool
						unicode    bool
						/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
					}{}
					clone.dotAll = (gopurs_runtime.RecordGet(orig, "dotAll").IntVal) != (0)
					clone.global = (gopurs_runtime.RecordGet(orig, "global").IntVal) != (0)
					clone.ignoreCase = (gopurs_runtime.RecordGet(orig, "ignoreCase").IntVal) != (0)
					clone.multiline = (gopurs_runtime.RecordGet(orig, "multiline").IntVal) != (0)
					clone.sticky = (gopurs_runtime.RecordGet(orig, "sticky").IntVal) != (0)
					clone.unicode = (gopurs_runtime.RecordGet(orig, "unicode").IntVal) != (0)
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
			}()
		})
	})
	return cache_Data_String_Regex_Flags_RegexFlags
}

var cache_Data_String_Regex_Flags_unicode gopurs_runtime.Value
var once_Data_String_Regex_Flags_unicode sync.Once

func Get_Data_String_Regex_Flags_unicode() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_unicode.Do(func() {
		cache_Data_String_Regex_Flags_unicode = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{false, false, false, false, false, true}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_unicode
}

var cache_Data_String_Regex_Flags_sticky gopurs_runtime.Value
var once_Data_String_Regex_Flags_sticky sync.Once

func Get_Data_String_Regex_Flags_sticky() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_sticky.Do(func() {
		cache_Data_String_Regex_Flags_sticky = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{false, false, false, false, true, false}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_sticky
}

var cache_Data_String_Regex_Flags_showRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_showRegexFlags sync.Once

func Get_Data_String_Regex_Flags_showRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_showRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_showRegexFlags = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_1446562561_1386611502((&Constructor_Data_Show_Show[struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		}]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			arrayIntrinsic_arg_1 := func(v_1 gopurs_runtime.Value) string {
				return "global"
			}
			_ = arrayIntrinsic_arg_1
			_ = arrayIntrinsic_arg_1
			var __t2 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0) {
					__t2 = gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()})
					goto end_branch_2
				} else {

				}
			}
			{
				if ((gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0)) != (true) {
					__t2 = gopurs_runtime.Array([]gopurs_runtime.Value{})
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			arrayIntrinsic_arg_3 := __t2
			_ = arrayIntrinsic_arg_3
			_ = arrayIntrinsic_arg_3
			arrayIntrinsic_arg_4 := func(v_1 gopurs_runtime.Value) string {
				return "ignoreCase"
			}
			_ = arrayIntrinsic_arg_4
			_ = arrayIntrinsic_arg_4
			var __t5 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0) {
					__t5 = gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()})
					goto end_branch_5
				} else {

				}
			}
			{
				if ((gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0)) != (true) {
					__t5 = gopurs_runtime.Array([]gopurs_runtime.Value{})
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_5:
			arrayIntrinsic_arg_6 := __t5
			_ = arrayIntrinsic_arg_6
			_ = arrayIntrinsic_arg_6
			arrayIntrinsic_arg_7 := func(v_1 gopurs_runtime.Value) string {
				return "multiline"
			}
			_ = arrayIntrinsic_arg_7
			_ = arrayIntrinsic_arg_7
			var __t8 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0) {
					__t8 = gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()})
					goto end_branch_8
				} else {

				}
			}
			{
				if ((gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0)) != (true) {
					__t8 = gopurs_runtime.Array([]gopurs_runtime.Value{})
					goto end_branch_8
				} else {

				}
			}
			{
				__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_8:
			arrayIntrinsic_arg_9 := __t8
			_ = arrayIntrinsic_arg_9
			_ = arrayIntrinsic_arg_9
			arrayIntrinsic_arg_10 := func(v_1 gopurs_runtime.Value) string {
				return "dotAll"
			}
			_ = arrayIntrinsic_arg_10
			_ = arrayIntrinsic_arg_10
			var __t11 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0) {
					__t11 = gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()})
					goto end_branch_11
				} else {

				}
			}
			{
				if ((gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0)) != (true) {
					__t11 = gopurs_runtime.Array([]gopurs_runtime.Value{})
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_11:
			arrayIntrinsic_arg_12 := __t11
			_ = arrayIntrinsic_arg_12
			_ = arrayIntrinsic_arg_12
			arrayIntrinsic_arg_13 := func(v_1 gopurs_runtime.Value) string {
				return "sticky"
			}
			_ = arrayIntrinsic_arg_13
			_ = arrayIntrinsic_arg_13
			var __t14 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0) {
					__t14 = gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()})
					goto end_branch_14
				} else {

				}
			}
			{
				if ((gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0)) != (true) {
					__t14 = gopurs_runtime.Array([]gopurs_runtime.Value{})
					goto end_branch_14
				} else {

				}
			}
			{
				__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_14:
			arrayIntrinsic_arg_15 := __t14
			_ = arrayIntrinsic_arg_15
			_ = arrayIntrinsic_arg_15
			arrayIntrinsic_arg_16 := func(v_1 gopurs_runtime.Value) string {
				return "unicode"
			}
			_ = arrayIntrinsic_arg_16
			_ = arrayIntrinsic_arg_16
			var __t17 gopurs_runtime.Value
			{
				if (gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0) {
					__t17 = gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Unit_unit()})
					goto end_branch_17
				} else {

				}
			}
			{
				if ((gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0)) != (true) {
					__t17 = gopurs_runtime.Array([]gopurs_runtime.Value{})
					goto end_branch_17
				} else {

				}
			}
			{
				__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_17:
			arrayIntrinsic_arg_18 := __t17
			_ = arrayIntrinsic_arg_18
			_ = arrayIntrinsic_arg_18
			// TAST (Let): usedFlags_1_0 shape=App(Var) bindingType=(Array String)
			usedFlags_1_0 := func() []string {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{}), func() gopurs_runtime.Value {
					arr := func() []string {
						arr_val_arrayMap9 := arrayIntrinsic_arg_3
						_ = arr_val_arrayMap9
						_ = arr_val_arrayMap9
						arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
						_ = arr_go_arrayMap9
						res_go_arrayMap9 := make([]string, len(*arr_go_arrayMap9))
						_ = res_go_arrayMap9
						for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
							res_go_arrayMap9[i_arrayMap9] = arrayIntrinsic_arg_1(v_arrayMap9)
						}
						return res_go_arrayMap9
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()), func() gopurs_runtime.Value {
					arr := func() []string {
						arr_val_arrayMap8 := arrayIntrinsic_arg_6
						_ = arr_val_arrayMap8
						_ = arr_val_arrayMap8
						arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
						_ = arr_go_arrayMap8
						res_go_arrayMap8 := make([]string, len(*arr_go_arrayMap8))
						_ = res_go_arrayMap8
						for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
							res_go_arrayMap8[i_arrayMap8] = arrayIntrinsic_arg_4(v_arrayMap8)
						}
						return res_go_arrayMap8
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()), func() gopurs_runtime.Value {
					arr := func() []string {
						arr_val_arrayMap7 := arrayIntrinsic_arg_9
						_ = arr_val_arrayMap7
						_ = arr_val_arrayMap7
						arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
						_ = arr_go_arrayMap7
						res_go_arrayMap7 := make([]string, len(*arr_go_arrayMap7))
						_ = res_go_arrayMap7
						for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
							res_go_arrayMap7[i_arrayMap7] = arrayIntrinsic_arg_7(v_arrayMap7)
						}
						return res_go_arrayMap7
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()), func() gopurs_runtime.Value {
					arr := func() []string {
						arr_val_arrayMap6 := arrayIntrinsic_arg_12
						_ = arr_val_arrayMap6
						_ = arr_val_arrayMap6
						arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
						_ = arr_go_arrayMap6
						res_go_arrayMap6 := make([]string, len(*arr_go_arrayMap6))
						_ = res_go_arrayMap6
						for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
							res_go_arrayMap6[i_arrayMap6] = arrayIntrinsic_arg_10(v_arrayMap6)
						}
						return res_go_arrayMap6
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()), func() gopurs_runtime.Value {
					arr := func() []string {
						arr_val_arrayMap5 := arrayIntrinsic_arg_15
						_ = arr_val_arrayMap5
						_ = arr_val_arrayMap5
						arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
						_ = arr_go_arrayMap5
						res_go_arrayMap5 := make([]string, len(*arr_go_arrayMap5))
						_ = res_go_arrayMap5
						for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
							res_go_arrayMap5[i_arrayMap5] = arrayIntrinsic_arg_13(v_arrayMap5)
						}
						return res_go_arrayMap5
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()), func() gopurs_runtime.Value {
					arr := func() []string {
						arr_val_arrayMap4 := arrayIntrinsic_arg_18
						_ = arr_val_arrayMap4
						_ = arr_val_arrayMap4
						arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
						_ = arr_go_arrayMap4
						res_go_arrayMap4 := make([]string, len(*arr_go_arrayMap4))
						_ = res_go_arrayMap4
						for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
							res_go_arrayMap4[i_arrayMap4] = arrayIntrinsic_arg_16(v_arrayMap4)
						}
						return res_go_arrayMap4
					}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr))).UnsafePtr)
				unboxed := make([]string, len(arr))
				for i, v := range arr {
					unboxed[i] = v.StrVal()
				}
				return unboxed
			}()
			_ = usedFlags_1_0
			var __t22 gopurs_runtime.Value
			{
				arrayEq_arg_19 := Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_1140313009_3790796878(Rebox_Data_String_Regex_Flags_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})
				_ = arrayEq_arg_19
				_ = arrayEq_arg_19
				arrayEq_arg_20 := usedFlags_1_0
				_ = arrayEq_arg_20
				_ = arrayEq_arg_20
				arrayEq_arg_21 := []string{}
				_ = arrayEq_arg_21
				_ = arrayEq_arg_21
				if func() bool {
					arrayEq_leftValue := arrayEq_arg_20
					_ = arrayEq_leftValue
					arrayEq_rightValue := arrayEq_arg_21
					_ = arrayEq_rightValue
					arrayEq_left_22 := arrayEq_leftValue
					_ = arrayEq_left_22
					arrayEq_right_22 := arrayEq_rightValue
					_ = arrayEq_right_22
					if (len(arrayEq_left_22)) != (len(arrayEq_right_22)) {
						return false
					} else {

					}
					for arrayEq_index_22 := range arrayEq_left_22 {
						if ((arrayEq_left_22)[arrayEq_index_22]) != ((arrayEq_right_22)[arrayEq_index_22]) {
							return false
						} else {

						}
					}
					return true
				}() {
					__t22 = gopurs_runtime.StrWithCodeUnits("noFlags", 7)
					goto end_branch_22
				} else {

				}
			}
			{
				__t22 = gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(", Data_String_Common_JoinWith(" <> ", usedFlags_1_0).StrVal()), ")"))
			}
		end_branch_22:
			return __t22
		})})))}
	})
	return cache_Data_String_Regex_Flags_showRegexFlags
}

var cache_Data_String_Regex_Flags_semigroupRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_semigroupRegexFlags sync.Once

func Get_Data_String_Regex_Flags_semigroupRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_semigroupRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_semigroupRegexFlags = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2919058177_4179793454((&Constructor_Data_Semigroup_Semigroup[struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		}]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "dotAll").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "global").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "ignoreCase").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "multiline").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "sticky").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "unicode").IntVal) != (0)))})
		})})))}
	})
	return cache_Data_String_Regex_Flags_semigroupRegexFlags
}

var cache_Data_String_Regex_Flags_noFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_noFlags sync.Once

func Get_Data_String_Regex_Flags_noFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_noFlags.Do(func() {
		cache_Data_String_Regex_Flags_noFlags = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{false, false, false, false, false, false}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_noFlags
}

var cache_Data_String_Regex_Flags_newtypeRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_newtypeRegexFlags sync.Once

func Get_Data_String_Regex_Flags_newtypeRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_newtypeRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_newtypeRegexFlags = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2766313800_385277032((&Constructor_Data_Newtype_Newtype[struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		}, struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		}]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_String_Regex_Flags_newtypeRegexFlags
}

var cache_Data_String_Regex_Flags_multiline gopurs_runtime.Value
var once_Data_String_Regex_Flags_multiline sync.Once

func Get_Data_String_Regex_Flags_multiline() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_multiline.Do(func() {
		cache_Data_String_Regex_Flags_multiline = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{false, false, false, true, false, false}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_multiline
}

var cache_Data_String_Regex_Flags_monoidRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_monoidRegexFlags sync.Once

func Get_Data_String_Regex_Flags_monoidRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_monoidRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_monoidRegexFlags = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2860751713_1201789390((&Constructor_Data_Monoid_Monoid[struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		}]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2919058177_4179793454(Rebox_Data_String_Regex_Flags_4179793454_2919058177(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_String_Regex_Flags_semigroupRegexFlags()))))}
		}), func() struct {
			dotAll     bool
			global     bool
			ignoreCase bool
			multiline  bool
			sticky     bool
			unicode    bool
			/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
		} {
			orig := Get_Data_String_Regex_Flags_noFlags()
			_ = orig
			clone := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{}
			clone.dotAll = (gopurs_runtime.RecordGet(orig, "dotAll").IntVal) != (0)
			clone.global = (gopurs_runtime.RecordGet(orig, "global").IntVal) != (0)
			clone.ignoreCase = (gopurs_runtime.RecordGet(orig, "ignoreCase").IntVal) != (0)
			clone.multiline = (gopurs_runtime.RecordGet(orig, "multiline").IntVal) != (0)
			clone.sticky = (gopurs_runtime.RecordGet(orig, "sticky").IntVal) != (0)
			clone.unicode = (gopurs_runtime.RecordGet(orig, "unicode").IntVal) != (0)
			return clone
		}()})))}
	})
	return cache_Data_String_Regex_Flags_monoidRegexFlags
}

var cache_Data_String_Regex_Flags_ignoreCase gopurs_runtime.Value
var once_Data_String_Regex_Flags_ignoreCase sync.Once

func Get_Data_String_Regex_Flags_ignoreCase() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_ignoreCase.Do(func() {
		cache_Data_String_Regex_Flags_ignoreCase = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{false, false, true, false, false, false}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_ignoreCase
}

var cache_Data_String_Regex_Flags_global gopurs_runtime.Value
var once_Data_String_Regex_Flags_global sync.Once

func Get_Data_String_Regex_Flags_global() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_global.Do(func() {
		cache_Data_String_Regex_Flags_global = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{false, true, false, false, false, false}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_global
}

var cache_Data_String_Regex_Flags_eqRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_eqRegexFlags sync.Once

func Get_Data_String_Regex_Flags_eqRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_eqRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_eqRegexFlags = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_646382433_3790796878(Rebox_Data_String_Regex_Flags_3790796878_646382433(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("unicode", 7)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2737952170_3790796878(Rebox_Data_String_Regex_Flags_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("sticky", 6)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2737952170_3790796878(Rebox_Data_String_Regex_Flags_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("multiline", 9)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2737952170_3790796878(Rebox_Data_String_Regex_Flags_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("ignoreCase", 10)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2737952170_3790796878(Rebox_Data_String_Regex_Flags_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("global", 6)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2737952170_3790796878(Rebox_Data_String_Regex_Flags_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.StrWithCodeUnits("dotAll", 6)
		})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Regex_Flags_2737952170_3790796878(Rebox_Data_String_Regex_Flags_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Data_String_Regex_Flags_eqRegexFlags
}

var cache_Data_String_Regex_Flags_dotAll gopurs_runtime.Value
var once_Data_String_Regex_Flags_dotAll sync.Once

func Get_Data_String_Regex_Flags_dotAll() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_dotAll.Do(func() {
		cache_Data_String_Regex_Flags_dotAll = func() gopurs_runtime.Value {
			orig := struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			}{true, false, false, false, false, false}
			_ = orig
			return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
		}()
	})
	return cache_Data_String_Regex_Flags_dotAll
}

func Call_Data_String_Regex_Flags_RegexFlags(x_0_loop struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}) struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
} {
	var x_0 struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	} = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Data_String_Regex_Flags_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_1446562561_1386611502(in *Constructor_Data_Show_Show[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_2766313800_385277032(in *Constructor_Data_Newtype_Newtype[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}, struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_2860751713_1201789390(in *Constructor_Data_Monoid_Monoid[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = func() gopurs_runtime.Value {
		orig := in.V1
		_ = orig
		return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
	}()
	return out
}

func Rebox_Data_String_Regex_Flags_2919058177_4179793454(in *Constructor_Data_Semigroup_Semigroup[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_3790796878_646382433(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_4179793454_2919058177(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Semigroup[struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Data_String_Regex_Flags_646382433_3790796878(in *Constructor_Data_Eq_Eq[struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
