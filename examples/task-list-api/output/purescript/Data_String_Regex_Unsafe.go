package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Regex_Unsafe_identity gopurs_runtime.Value
var once_Data_String_Regex_Unsafe_identity sync.Once

func Get_Data_String_Regex_Unsafe_identity() gopurs_runtime.Value {
	once_Data_String_Regex_Unsafe_identity.Do(func() {
		cache_Data_String_Regex_Unsafe_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Regex_Unsafe_identity(x_0_box)
		})
	})
	return cache_Data_String_Regex_Unsafe_identity
}

var cache_Data_String_Regex_Unsafe_unsafeRegex gopurs_runtime.Value
var once_Data_String_Regex_Unsafe_unsafeRegex sync.Once

func Get_Data_String_Regex_Unsafe_unsafeRegex() gopurs_runtime.Value {
	once_Data_String_Regex_Unsafe_unsafeRegex.Do(func() {
		cache_Data_String_Regex_Unsafe_unsafeRegex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Regex_Unsafe_unsafeRegex(s_0_box.StrVal(), func() struct {
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
		})
	})
	return cache_Data_String_Regex_Unsafe_unsafeRegex
}

func Call_Data_String_Regex_Unsafe_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_String_Regex_Unsafe_unsafeRegex(s_0_loop string, f_1_loop struct {
	dotAll     bool
	global     bool
	ignoreCase bool
	multiline  bool
	sticky     bool
	unicode    bool
	/*gopurs:record-v1:6:0064006f00740041006c006c,0067006c006f00620061006c,00690067006e006f007200650043006100730065,006d0075006c00740069006c0069006e0065,0073007400690063006b0079,0075006e00690063006f00640065*/
}) gopurs_runtime.Value {
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
	return Call_Data_Either_either__3820936092(Get_Partial_Unsafe_unsafeCrashWith(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}), func() gopurs_runtime.Value {
		_v := Call_Data_String_Regex_regex(s_0, f_1)
		if _v.V2 {
			return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
	}())
}
