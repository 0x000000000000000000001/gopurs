package purescript

import "gopurs/output/gopurs_runtime"

import (
	"strings"
	"unicode/utf8"
)

// A PureScript String has UTF-16 semantics. A surrogate pair may arrive either
// as one UTF-8 scalar or as two WTF-8 units after concatenating strings.
func decodeCodePoint(s string) (rune, int) {
	r, width := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && width == 1 && len(s) >= 3 && s[0] == 0xED && s[1] >= 0xA0 && s[1] <= 0xBF && s[2]&0xC0 == 0x80 {
		r = rune(int64(s[0]&0x0F)<<12 | int64(s[1]&0x3F)<<6 | int64(s[2]&0x3F))
		width = 3
		if r <= 0xDBFF && len(s) >= 6 && s[3] == 0xED && s[4] >= 0xB0 && s[4] <= 0xBF && s[5]&0xC0 == 0x80 {
			low := rune(int64(s[3]&0x0F)<<12 | int64(s[4]&0x3F)<<6 | int64(s[5]&0x3F))
			return 0x10000 + (r-0xD800)*0x400 + low - 0xDC00, 6
		}
	}
	return r, width
}

func Data_String_CodePoints_Length(s string) int64 {
	count := int64(0)
	for i := 0; i < len(s); count++ {
		_, width := decodeCodePoint(s[i:])
		i += width
	}
	return count
}

func decodeWTF8(s string) []rune {
	var res []rune
	for i := 0; i < len(s); {
		r, size := decodeCodePoint(s[i:])
		res = append(res, r)
		i += size
	}
	return res
}

func encodeWTF8(runes []rune) string {
	var bytes []byte
	buf := make([]byte, 4)
	for _, r := range runes {
		if r >= 0xD800 && r <= 0xDFFF {
			bytes = append(bytes, 0xED, byte(0xA0|((r>>6)&0x3F)), byte(0x80|(r&0x3F)))
		} else {
			n := utf8.EncodeRune(buf, r)
			bytes = append(bytes, buf[:n]...)
		}
	}
	return string(bytes)
}

func Data_String_CodePoints__UnsafeCodePointAt0(fallback interface{}, str string) int64 {
	runes := decodeWTF8(str)
	if len(runes) > 0 {
		return int64(runes[0])
	}
	return 0
}

func Data_String_CodePoints__CodePointAt(fallback interface{}, just func(interface{}) interface{}, nothing interface{}, unsafeCodePointAt0 interface{}, index int64, str string) interface{} {
	runes := decodeWTF8(str)
	if index < 0 || index >= int64(len(runes)) {
		return nothing
	}
	return just(int64(runes[index]))
}

func Data_String_CodePoints__CountPrefix(fallback interface{}, unsafeCodePointAt0 interface{}, pred func(int64) bool, str string) int64 {
	runes := decodeWTF8(str)
	count := int64(0)
	for _, r := range runes {
		if !pred(int64(r)) {
			break
		}
		count++
	}
	return count
}

func Data_String_CodePoints__FromCodePointArray(singleton interface{}, cps []interface{}) string {
	runes := make([]rune, len(cps))
	for i, cp := range cps {
		switch v := cp.(type) {
		case int64:
			runes[i] = rune(v)
		case float64:
			runes[i] = rune(v)
		case int:
			runes[i] = rune(v)
		}
	}
	return encodeWTF8(runes)
}

func Data_String_CodePoints__Singleton(fallback interface{}, cp int64) string {
	return encodeWTF8([]rune{rune(cp)})
}

func Data_String_CodePoints__Take(fallback interface{}, n int64, str string) string {
	if n <= 0 {
		return ""
	}
	offset := 0
	malformed := false
	for count := int64(0); count < n && offset < len(str); count++ {
		r, width := decodeCodePoint(str[offset:])
		malformed = malformed || (r == utf8.RuneError && width == 1)
		offset += width
	}
	if offset == len(str) {
		return str
	}
	if malformed {
		// Preserve the existing replacement-character policy for malformed
		// bytes inside a proper prefix. Bytes beyond the prefix are irrelevant.
		return encodeWTF8(decodeWTF8(str[:offset]))
	}
	// A proper prefix continues to own its storage, avoiding retention of a
	// potentially much larger input string.
	return strings.Clone(str[:offset])
}

func Data_String_CodePoints__ToCodePointArray(fallback interface{}, unsafeCodePointAt0 interface{}, str string) []int64 {
	res := make([]int64, Data_String_CodePoints_Length(str))
	for offset, i := 0, 0; offset < len(str); i++ {
		r, width := decodeCodePoint(str[offset:])
		res[i] = int64(r)
		offset += width
	}
	return res
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_String_CodePoints__CodePointAt = // TAST: (Func [(Func [Int, String] (ADT ["Data","Maybe","Maybe"] [Int])), (ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (Func [String] Int), Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := gopurs_runtime.Unbox[int64](arg4)
	go_arg5 := gopurs_runtime.Unbox[string](arg5)
	go_res := Data_String_CodePoints__CodePointAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodePoints__CountPrefix = // TAST: (Func [(Func [(Func [Int] Boolean), String] Int), (Func [String] Int), (Func [Int] Boolean), String] Int)
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := func(p0_0 int64) bool {
		inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Int(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodePoints__CountPrefix(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Data_String_CodePoints__FromCodePointArray = // TAST: (Func [(Func [Int] String), (Array Int)] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func() []any {
		_source := arg1
		_out := make([]any, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.ArrayAccess(_source, _i)
		}
		return _out
	}()
	go_res := Data_String_CodePoints__FromCodePointArray(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodePoints__Singleton = // TAST: (Func [(Func [Int] String), Int] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := Data_String_CodePoints__Singleton(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodePoints__Take = // TAST: (Func [(Func [Int, String] String), Int, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_CodePoints__Take(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodePoints__ToCodePointArray = // TAST: (Func [(Func [String] (Array Int)), (Func [String] Int), String] (Array Int))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_CodePoints__ToCodePointArray(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})
var _Gopurs_Data_String_CodePoints__UnsafeCodePointAt0 = // TAST: (Func [(Func [String] Int), String] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodePoints__UnsafeCodePointAt0(go_arg0, go_arg1)
	return gopurs_runtime.Int(go_res)
})
var _Gopurs_Data_String_CodePoints_Length = // TAST: (Func [String] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_CodePoints_Length(go_arg0)
	return gopurs_runtime.Int(go_res)
})

// --- Boxed-result FFI workers ---
func Data_String_CodePoints__CodePointAt_nativeWorker(arg0 any, arg1 gopurs_runtime.Value, arg2 any, arg3 any, arg4 int64, arg5 string) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_res := Data_String_CodePoints__CodePointAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodePoints__CountPrefix_nativeWorker(arg0 any, arg1 any, arg2 gopurs_runtime.Value, arg3 string) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := func(p0_0 int64) bool {
		inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Int(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg3 := arg3
	go_res := Data_String_CodePoints__CountPrefix(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}
