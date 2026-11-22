package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Regex_go__append gopurs_runtime.Value
var once_Data_String_Regex_go__append sync.Once

func Get_Data_String_Regex_go__append() gopurs_runtime.Value {
	once_Data_String_Regex_go__append.Do(func() {
		cache_Data_String_Regex_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_Regex_go__append
}

var cache_Data_String_Regex_showRegex gopurs_runtime.Value
var once_Data_String_Regex_showRegex sync.Once

func Get_Data_String_Regex_showRegex() gopurs_runtime.Value {
	once_Data_String_Regex_showRegex.Do(func() {
		cache_Data_String_Regex_showRegex = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, Get_Data_String_Regex_showRegexImpl()}))}
	})
	return cache_Data_String_Regex_showRegex
}

var cache_Data_String_Regex_search gopurs_runtime.Value
var once_Data_String_Regex_search sync.Once

func Get_Data_String_Regex_search() gopurs_runtime.Value {
	once_Data_String_Regex_search.Do(func() {
		cache_Data_String_Regex_search = gopurs_runtime.Apply2(Get_Data_String_Regex__search(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_Regex_search
}

var cache_Data_String_Regex_replace_prime_ gopurs_runtime.Value
var once_Data_String_Regex_replace_prime_ sync.Once

func Get_Data_String_Regex_replace_prime_() gopurs_runtime.Value {
	once_Data_String_Regex_replace_prime_.Do(func() {
		cache_Data_String_Regex_replace_prime_ = gopurs_runtime.Apply2(Get_Data_String_Regex__replaceBy(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_Regex_replace_prime_
}

var cache_Data_String_Regex_renderFlags gopurs_runtime.Value
var once_Data_String_Regex_renderFlags sync.Once

func Get_Data_String_Regex_renderFlags() gopurs_runtime.Value {
	once_Data_String_Regex_renderFlags.Do(func() {
		cache_Data_String_Regex_renderFlags = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_Regex_renderFlags(func() struct {
				dotAll     bool
				global     bool
				ignoreCase bool
				multiline  bool
				sticky     bool
				unicode    bool
				/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
			} {
				orig := v_0_box
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
			}()))
		})
	})
	return cache_Data_String_Regex_renderFlags
}

var cache_Data_String_Regex_regex gopurs_runtime.Value
var once_Data_String_Regex_regex sync.Once

func Get_Data_String_Regex_regex() gopurs_runtime.Value {
	once_Data_String_Regex_regex.Do(func() {
		cache_Data_String_Regex_regex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_String_Regex_regex(s_0_box.StrVal(), func() struct {
					dotAll     bool
					global     bool
					ignoreCase bool
					multiline  bool
					sticky     bool
					unicode    bool
					/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
				} {
					orig := f_1_box
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
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_String_Regex_regex
}

var cache_Data_String_Regex_parseFlags gopurs_runtime.Value
var once_Data_String_Regex_parseFlags sync.Once

func Get_Data_String_Regex_parseFlags() gopurs_runtime.Value {
	once_Data_String_Regex_parseFlags.Do(func() {
		cache_Data_String_Regex_parseFlags = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Data_String_Regex_parseFlags(s_0_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.dotAll), gopurs_runtime.Bool(orig.global), gopurs_runtime.Bool(orig.ignoreCase), gopurs_runtime.Bool(orig.multiline), gopurs_runtime.Bool(orig.sticky), gopurs_runtime.Bool(orig.unicode)})
			}()
		})
	})
	return cache_Data_String_Regex_parseFlags
}

var cache_Data_String_Regex_match gopurs_runtime.Value
var once_Data_String_Regex_match sync.Once

func Get_Data_String_Regex_match() gopurs_runtime.Value {
	once_Data_String_Regex_match.Do(func() {
		cache_Data_String_Regex_match = gopurs_runtime.Apply2(Get_Data_String_Regex__match(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))})
	})
	return cache_Data_String_Regex_match
}

var cache_Data_String_Regex_flags gopurs_runtime.Value
var once_Data_String_Regex_flags sync.Once

func Get_Data_String_Regex_flags() gopurs_runtime.Value {
	once_Data_String_Regex_flags.Do(func() {
		cache_Data_String_Regex_flags = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), Get_Data_String_Regex_flagsImpl())
	})
	return cache_Data_String_Regex_flags
}

func Call_Data_String_Regex_renderFlags(v_0_loop struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}) string {
	var v_0 struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	} = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0.global {
			__t0 = "g"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = ""
	}
end_branch_0:
	var __t1 string
	{
		if v_0.ignoreCase {
			__t1 = "i"
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = ""
	}
end_branch_1:
	var __t2 string
	{
		if v_0.multiline {
			__t2 = "m"
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = ""
	}
end_branch_2:
	var __t3 string
	{
		if v_0.dotAll {
			__t3 = "s"
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = ""
	}
end_branch_3:
	var __t4 string
	{
		if v_0.sticky {
			__t4 = "y"
			goto end_branch_4
		} else {

		}
	}
	{
		__t4 = ""
	}
end_branch_4:
	var __t5 string
	{
		if v_0.unicode {
			__t5 = "u"
			goto end_branch_5
		} else {

		}
	}
	{
		__t5 = ""
	}
end_branch_5:
	return gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(__t0, __t1), __t2), __t3), __t4), __t5)
}

func Call_Data_String_Regex_regex(s_0_loop string, f_1_loop struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var s_0 string = s_0_loop
	_ = s_0
	var f_1 struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	} = f_1_loop
	_ = f_1
	var __t0 string
	{
		if f_1.global {
			__t0 = "g"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = ""
	}
end_branch_0:
	var __t1 string
	{
		if f_1.ignoreCase {
			__t1 = "i"
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = ""
	}
end_branch_1:
	var __t2 string
	{
		if f_1.multiline {
			__t2 = "m"
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = ""
	}
end_branch_2:
	var __t3 string
	{
		if f_1.dotAll {
			__t3 = "s"
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = ""
	}
end_branch_3:
	var __t4 string
	{
		if f_1.sticky {
			__t4 = "y"
			goto end_branch_4
		} else {

		}
	}
	{
		__t4 = ""
	}
end_branch_4:
	var __t5 string
	{
		if f_1.unicode {
			__t5 = "u"
			goto end_branch_5
		} else {

		}
	}
	{
		__t5 = ""
	}
end_branch_5:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	} { _v := Data_String_Regex_RegexImpl_nativeWorker(Get_Data_Either_Left(), Get_Data_Either_Right(), s_0, gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(__t0, __t1), __t2), __t3), __t4), __t5)); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
}

func Call_Data_String_Regex_parseFlags(s_0_loop string) struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
} {
	var s_0 string = s_0_loop
	_ = s_0
	return struct {
		dotAll     bool
		global     bool
		ignoreCase bool
		multiline  bool
		sticky     bool
		unicode    bool
		/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
	}{(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.StrWithCodeUnits("s", 1)), gopurs_runtime.Str(s_0)).IntVal) != (0), (gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.StrWithCodeUnits("g", 1)), gopurs_runtime.Str(s_0)).IntVal) != (0), (gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.StrWithCodeUnits("i", 1)), gopurs_runtime.Str(s_0)).IntVal) != (0), (gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.StrWithCodeUnits("m", 1)), gopurs_runtime.Str(s_0)).IntVal) != (0), (gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.StrWithCodeUnits("y", 1)), gopurs_runtime.Str(s_0)).IntVal) != (0), (gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.StrWithCodeUnits("u", 1)), gopurs_runtime.Str(s_0)).IntVal) != (0)}
}

func Get_Data_String_Regex__match() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex__Match
}

func Get_Data_String_Regex__replaceBy() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex__ReplaceBy
}

func Get_Data_String_Regex__search() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex__Search
}

func Get_Data_String_Regex_flagsImpl() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_FlagsImpl
}

func Get_Data_String_Regex_regexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_RegexImpl
}

func Get_Data_String_Regex_replace() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Replace
}

func Get_Data_String_Regex_showRegexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_ShowRegexImpl
}

func Get_Data_String_Regex_source() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Source
}

func Get_Data_String_Regex_split() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Split
}

func Get_Data_String_Regex_test() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Test
}
