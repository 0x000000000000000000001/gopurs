package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// PureScript Char and CodeUnits indices denote UTF-16 units. Native strings
// use UTF-8, with WTF-8 for isolated surrogates (including halves of a slice).
func decodeUnitRune(s string) (rune, int) {
	if len(s) >= 3 && s[0] == 0xed && s[1] >= 0xa0 && s[1] <= 0xbf && s[2]&0xc0 == 0x80 {
		return rune(s[0]&0x0f)<<12 | rune(s[1]&0x3f)<<6 | rune(s[2]&0x3f), 3
	}
	return utf8.DecodeRuneInString(s)
}

func unitString(c uint16) string {
	if c >= 0xd800 && c <= 0xdfff {
		return string([]byte{0xed, 0x80 | byte(c>>6&0x3f), 0x80 | byte(c&0x3f)})
	}
	return string(rune(c))
}

func stringUnits(s string) []uint16 {
	units := make([]uint16, 0, gopurs_runtime.CountStringCodeUnits(s))
	for i := 0; i < len(s); {
		r, size := decodeUnitRune(s[i:])
		if r > 0xffff {
			hi, lo := utf16.EncodeRune(r)
			units = append(units, uint16(hi), uint16(lo))
		} else {
			units = append(units, uint16(r))
		}
		i += size
	}
	return units
}

func Data_String_CodeUnits_FromCharArray(a []string) string {
	var units []uint16
	for _, c := range a {
		units = append(units, stringUnits(c)...)
	}
	var b strings.Builder
	for i := 0; i < len(units); i++ {
		c := units[i]
		if c >= 0xd800 && c <= 0xdbff && i+1 < len(units) && units[i+1] >= 0xdc00 && units[i+1] <= 0xdfff {
			b.WriteRune(utf16.DecodeRune(rune(c), rune(units[i+1])))
			i++
		} else {
			b.WriteString(unitString(c))
		}
	}
	return b.String()
}

func Data_String_CodeUnits_ToCharArray(str string) []string {
	arr := make([]string, 0, gopurs_runtime.CountStringCodeUnits(str))
	for i := 0; i < len(str); {
		r, size := decodeUnitRune(str[i:])
		if r > 0xffff {
			hi, lo := utf16.EncodeRune(r)
			arr = append(arr, unitString(uint16(hi)), unitString(uint16(lo)))
		} else {
			arr = append(arr, str[i:i+size])
		}
		i += size
	}
	return arr
}

func Data_String_CodeUnits_Singleton(c string) string {
	return c
}

// Return the byte boundary for a UTF-16 index. split marks the midpoint of a
// supplementary scalar, where slicing must produce an isolated surrogate.
func unitBoundary(s string, index int) (offset int, split bool) {
	if index <= 0 {
		return 0, false
	}
	for i, unit := 0, 0; i < len(s); {
		if unit == index {
			return i, false
		}
		r, size := decodeUnitRune(s[i:])
		if r > 0xffff {
			if unit+1 == index {
				return i, true
			}
			unit++
		}
		unit++
		i += size
	}
	return len(s), false
}

func Data_String_CodeUnits__CharAt(just func(string) interface{}, nothing interface{}, idx int, str string) interface{} {
	if idx < 0 {
		return nothing
	}
	offset, split := unitBoundary(str, idx)
	if offset == len(str) {
		return nothing
	}
	r, size := decodeUnitRune(str[offset:])
	if r > 0xffff {
		hi, lo := utf16.EncodeRune(r)
		if split {
			return just(unitString(uint16(lo)))
		}
		return just(unitString(uint16(hi)))
	}
	return just(str[offset : offset+size])
}

func Data_String_CodeUnits__ToChar(just func(string) interface{}, nothing interface{}, str string) interface{} {
	if len(str) > 0 {
		r, size := decodeUnitRune(str)
		if size == len(str) && r <= 0xffff {
			return just(str)
		}
	}
	return nothing
}

// The boxed string may carry immutable UTF-16 length metadata. Native strings
// without metadata retain the same UTF-8/WTF-8 scan.
func Data_String_CodeUnits_Length(value gopurs_runtime.Value) int {
	return gopurs_runtime.StringCodeUnitsLength(value)
}

func Data_String_CodeUnits_CountPrefix(p func(string) bool, str string) int {
	count := 0
	for i := 0; i < len(str); {
		r, size := decodeUnitRune(str[i:])
		if r > 0xffff {
			hi, lo := utf16.EncodeRune(r)
			if !p(unitString(uint16(hi))) {
				return count
			}
			count++
			if !p(unitString(uint16(lo))) {
				return count
			}
		} else if !p(str[i : i+size]) {
			return count
		}
		count++
		i += size
	}
	return count
}

func findUnits(pattern, input []uint16, start int, backwards bool) int {
	if start > len(input)-len(pattern) {
		start = len(input) - len(pattern)
	}
	for i := start; i >= 0 && i+len(pattern) <= len(input); {
		match := true
		for j, c := range pattern {
			if input[i+j] != c {
				match = false
				break
			}
		}
		if match {
			return i
		}
		if backwards {
			i--
		} else {
			i++
		}
	}
	return -1
}

func Data_String_CodeUnits__IndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	return Data_String_CodeUnits__IndexOfStartingAt(just, nothing, x, 0, s)
}

// A pattern made only of BMP, non-surrogate code points is encoded the same way
// in UTF-8/WTF-8 wherever its code units match, so a byte search is exact for
// it. Supplementary scalars and isolated surrogates must use the code-unit path
// because the same UTF-16 units can have two encodings.
func isBMPText(s string) bool {
	for i := 0; i < len(s); {
		r, size := decodeUnitRune(s[i:])
		if r > 0xffff || (r >= 0xd800 && r <= 0xdfff) {
			return false
		}
		i += size
	}
	return true
}

func Data_String_CodeUnits__IndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx >= 0 {
		begin, split := unitBoundary(str, startIdx)
		if !split && isBMPText(x) {
			if startIdx > 0 && begin == len(str) && startIdx != gopurs_runtime.CountStringCodeUnits(str) {
				return nothing
			}
			off := strings.Index(str[begin:], x)
			if off < 0 {
				return nothing
			}
			return just(startIdx + gopurs_runtime.CountStringCodeUnits(str[begin:begin+off]))
		}
	}
	input, pattern := stringUnits(str), stringUnits(x)
	if startIdx < 0 || startIdx > len(input) || startIdx+len(pattern) > len(input) {
		return nothing
	}
	idx := findUnits(pattern, input, startIdx, false)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Data_String_CodeUnits__LastIndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	input, pattern := stringUnits(s), stringUnits(x)
	idx := findUnits(pattern, input, len(input), true)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Data_String_CodeUnits__LastIndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	input, pattern := stringUnits(str), stringUnits(x)
	if startIdx < 0 {
		startIdx = 0
	}
	idx := findUnits(pattern, input, startIdx, true)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func sliceUnits(str string, start, end int) string {
	if start >= end {
		return ""
	}
	begin, splitBegin := unitBoundary(str, start)
	finish, splitEnd := unitBoundary(str, end)
	if !splitBegin && !splitEnd {
		return str[begin:finish]
	}
	var b strings.Builder
	if splitBegin {
		r, size := decodeUnitRune(str[begin:])
		_, lo := utf16.EncodeRune(r)
		b.WriteString(unitString(uint16(lo)))
		begin += size
	}
	b.WriteString(str[begin:finish])
	if splitEnd {
		r, _ := decodeUnitRune(str[finish:])
		hi, _ := utf16.EncodeRune(r)
		b.WriteString(unitString(uint16(hi)))
	}
	return b.String()
}

func Data_String_CodeUnits_Take(idx int, str string) string {
	if idx <= 0 {
		return ""
	}
	return sliceUnits(str, 0, idx)
}

func Data_String_CodeUnits_Drop(idx int, str string) string {
	if idx <= 0 {
		return str
	}
	// Dropping a prefix only needs the boundary of the cut: computing
	// gopurs_runtime.CountStringCodeUnits(str) here made every drop O(len(str)), which turned token-by-token
	// lexing into quadratic work. unitBoundary already handles a cut inside a
	// supplementary scalar, which must keep the low surrogate.
	begin, split := unitBoundary(str, idx)
	if begin >= len(str) {
		return ""
	}
	if !split {
		return str[begin:]
	}
	r, size := decodeUnitRune(str[begin:])
	_, lo := utf16.EncodeRune(r)
	return unitString(uint16(lo)) + str[begin+size:]
}

func Data_String_CodeUnits_Slice(start int, end int, str string) string {
	length := gopurs_runtime.CountStringCodeUnits(str)
	if start < 0 {
		start += length
	}
	if end < 0 {
		end += length
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}
	return sliceUnits(str, start, end)
}

func Data_String_CodeUnits_SplitAt(idx int, str string) map[string]interface{} {
	return map[string]interface{}{"before": Data_String_CodeUnits_Take(idx, str), "after": Data_String_CodeUnits_Drop(idx, str)}
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_String_CodeUnits__CharAt = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), Int, String] (ADT ["Data","Maybe","Maybe"] [Char]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodeUnits__CharAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__IndexOf = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodeUnits__IndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__IndexOfStartingAt = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_String_CodeUnits__IndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__LastIndexOf = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodeUnits__LastIndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__LastIndexOfStartingAt = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_String_CodeUnits__LastIndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__ToChar = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String] (ADT ["Data","Maybe","Maybe"] [Char]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_CodeUnits__ToChar(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_CountPrefix = // TAST: (Func [(Func [Char] Boolean), String] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) bool {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_CountPrefix(go_arg0, go_arg1)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_String_CodeUnits_Drop = // TAST: (Func [Int, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_Drop(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_FromCharArray = // TAST: (Func [(Array Char)] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() []string {
		_source := arg0
		_out := make([]string, gopurs_runtime.ArrayLength(_source))
		for _i := range _out {
			_out[_i] = gopurs_runtime.Unbox[string](gopurs_runtime.ArrayAccess(_source, _i))
		}
		return _out
	}()
	go_res := Data_String_CodeUnits_FromCharArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_Length = // TAST: (Func [String] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_CodeUnits_Length(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})
var _Gopurs_Data_String_CodeUnits_Singleton = // TAST: (Func [Char] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_CodeUnits_Singleton(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_Slice = // TAST: (Func [Int, Int, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_CodeUnits_Slice(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_SplitAt = // TAST: (Func [Int, String] (Record (Row [before: String, after: String] Empty)))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_SplitAt(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
		_raw := go_res
		res_map := make(map[string]gopurs_runtime.Value)
		res_map["before"] = gopurs_runtime.Box(_raw["before"])
		res_map["after"] = gopurs_runtime.Box(_raw["after"])
		return gopurs_runtime.Record(res_map)
	}()
})
var _Gopurs_Data_String_CodeUnits_Take = // TAST: (Func [Int, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_Take(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_ToCharArray = // TAST: (Func [String] (Array Char))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_CodeUnits_ToCharArray(go_arg0)
	return func() gopurs_runtime.Value {
		res_arr := make([]gopurs_runtime.Value, len(go_res))
		for i, v := range go_res {
			res_arr[i] = gopurs_runtime.Box(v)
		}
		return gopurs_runtime.Array(res_arr)
	}()
})

// --- Boxed-result FFI workers ---
func Data_String_CodeUnits__CharAt_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 int, arg3 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_String_CodeUnits__CharAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodeUnits__IndexOf_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 string, arg3 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_String_CodeUnits__IndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodeUnits__IndexOfStartingAt_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 string, arg3 int, arg4 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_String_CodeUnits__IndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodeUnits__LastIndexOf_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 string, arg3 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Data_String_CodeUnits__LastIndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodeUnits__LastIndexOfStartingAt_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 string, arg3 int, arg4 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Int(int64(p0_0)))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_String_CodeUnits__LastIndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodeUnits__ToChar_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_String_CodeUnits__ToChar(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}

func Data_String_CodeUnits_CountPrefix_nativeWorker(arg0 gopurs_runtime.Value, arg1 string) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) bool {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return gopurs_runtime.Unbox[bool](inner_res0)
	}
	go_arg1 := arg1
	go_res := Data_String_CodeUnits_CountPrefix(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
