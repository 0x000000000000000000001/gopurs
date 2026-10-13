package Gopurs_Cache_Wire

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

func Hash(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

func Hex(text string) bool {
	for i := 0; i < len(text); i++ {
		c := text[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}

func CompareStrings(a, b string) int64 {
	// UTF-8 and UTF-16 have the same order until a supplementary scalar occurs.
	wide := func(s string) bool {
		for i := 0; i < len(s); i++ {
			if s[i] >= 0xf0 {
				return true
			}
		}
		return false
	}
	if wide(a) || wide(b) {
		a, b = StringUnits(a), StringUnits(b)
	}
	return int64(strings.Compare(a, b))
}

func StringUnits(text string) string {
	const digits = "0123456789abcdef"
	var b strings.Builder
	unit := func(v uint16) {
		b.WriteByte(digits[v>>12])
		b.WriteByte(digits[v>>8&15])
		b.WriteByte(digits[v>>4&15])
		b.WriteByte(digits[v&15])
	}
	for len(text) > 0 {
		r, n := utf8.DecodeRuneInString(text)
		if len(text) >= 3 && text[0] == 0xed && text[1] >= 0xa0 && text[1] <= 0xbf && text[2]&0xc0 == 0x80 {
			r, n = rune(text[0]&15)<<12|rune(text[1]&63)<<6|rune(text[2]&63), 3
		}
		if r > 0xffff {
			hi, lo := utf16.EncodeRune(r)
			unit(uint16(hi))
			unit(uint16(lo))
		} else {
			unit(uint16(r))
		}
		text = text[n:]
	}
	return b.String()
}

func UnitsString(text string) string {
	var b strings.Builder
	for i := 0; i < len(text); i += 4 {
		v, _ := strconv.ParseUint(text[i:i+4], 16, 16)
		if v >= 0xd800 && v <= 0xdbff && i+8 <= len(text) {
			lo, _ := strconv.ParseUint(text[i+4:i+8], 16, 16)
			if lo >= 0xdc00 && lo <= 0xdfff {
				b.WriteRune(utf16.DecodeRune(rune(v), rune(lo)))
				i += 4
				continue
			}
		}
		if v >= 0xd800 && v <= 0xdfff {
			b.WriteByte(0xed)
			b.WriteByte(0x80 | byte(v>>6&63))
			b.WriteByte(0x80 | byte(v&63))
		} else {
			b.WriteRune(rune(v))
		}
	}
	return b.String()
}

func NumberBits(value float64) string {
	text := strconv.FormatUint(math.Float64bits(value), 16)
	return strings.Repeat("0", 16-len(text)) + text
}

func BitsNumber(text string) float64 {
	bits, _ := strconv.ParseUint(text, 16, 64)
	return math.Float64frombits(bits)
}
