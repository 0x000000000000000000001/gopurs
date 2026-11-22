package purescript

import "gopurs/output/gopurs_runtime"

import "unicode/utf8"

// This module's FFI is emitted independently of Data.String.CodeUnits.
func unsafeUnitAt(str string, idx int) (string, bool) {
	if idx < 0 {
		return "", false
	}
	for i, unit := 0, 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		if size == 1 && i+2 < len(str) && str[i] == 0xed && str[i+1] >= 0xa0 && str[i+1] <= 0xbf && str[i+2]&0xc0 == 0x80 {
			r, size = rune(str[i]&0x0f)<<12|rune(str[i+1]&0x3f)<<6|rune(str[i+2]&0x3f), 3
		}
		if r > 0xffff {
			if idx == unit || idx == unit+1 {
				c := uint16(0xd800 + ((r - 0x10000) >> 10))
				if idx == unit+1 {
					c = uint16(0xdc00 + ((r - 0x10000) & 0x3ff))
				}
				return string([]byte{0xed, 0x80 | byte(c>>6&0x3f), 0x80 | byte(c&0x3f)}), true
			}
			unit++
		} else if idx == unit {
			return str[i : i+size], true
		}
		unit++
		i += size
	}
	return "", false
}

func Data_String_Unsafe_CharAt(i interface{}) interface{} {
	return func(s interface{}) interface{} {
		str := gopurs_runtime.Unbox[string](s)
		idx := gopurs_runtime.Unbox[int](i)
		if char, ok := unsafeUnitAt(str, idx); ok {
			return char
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	}
}

func Data_String_Unsafe_Char(s interface{}) interface{} {
	str := gopurs_runtime.Unbox[string](s)
	if char, ok := unsafeUnitAt(str, 0); ok && char == str {
		return char
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_String_Unsafe_Char = // TAST: (Func [String] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_Unsafe_Char(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Unsafe_CharAt = // TAST: (Func [Int, String] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_Unsafe_CharAt(go_arg0)
	return gopurs_runtime.Func(func(p0 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			if fn, ok := go_res.(func(any) any); ok {
				return gopurs_runtime.Box(fn(gopurs_runtime.Unbox[any](p0)))
			}
			return gopurs_runtime.Apply(gopurs_runtime.Box(go_res), p0)
		}()
	})
})

// --- Boxed-result FFI workers ---
func Data_String_Unsafe_Char_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_Unsafe_Char(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Data_String_Unsafe_CharAt_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_Unsafe_CharAt(go_arg0)
	return gopurs_runtime.Box(go_res)
}
