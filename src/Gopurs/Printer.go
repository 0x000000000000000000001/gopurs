package Gopurs_Printer

import "strings"

// Preserve the JavaScript printer's literal spelling, including its WTF-8
// encoding of isolated UTF-16 surrogates. Valid UTF-8 is copied byte for byte;
// Go's rune iterator would replace those surrogate bytes with U+FFFD.
func EscapeGoStringImpl(value string) string {
	const hex = "0123456789abcdef"
	var out strings.Builder
	out.Grow(len(value))
	writeByteEscape := func(b byte) {
		out.WriteString("\\x")
		out.WriteByte(hex[b>>4])
		out.WriteByte(hex[b&15])
	}
	for i := 0; i < len(value); i++ {
		b := value[i]
		switch {
		case b == '"' || b == '\\':
			out.WriteByte('\\')
			out.WriteByte(b)
		case b < 32 || b == 127:
			writeByteEscape(b)
		case b == 0xed && i+2 < len(value) && value[i+1] >= 0xa0 && value[i+1] <= 0xbf && value[i+2] >= 0x80 && value[i+2] <= 0xbf:
			writeByteEscape(b)
			writeByteEscape(value[i+1])
			writeByteEscape(value[i+2])
			i += 2
		default:
			out.WriteByte(b)
		}
	}
	return out.String()
}
