package gopurs_runtime

import (
	"strings"
	"testing"
)

func TestASCIIWordScanRetainsUnicodeAndMalformedBoundaries(t *testing.T) {
	markers := []string{
		"", "é", "漢", "😀", "\xed\xa0\x80", "\xed\xbf\xbf",
		"\xed\xa0\x80\xed\xb0\x80", "\xff", "\xe0\x9f\xbf", "\xf0\x9f", "\xf4\x90\x80\x80",
	}
	for prefix := 0; prefix <= 512; prefix++ {
		for _, marker := range markers {
			// Slices at each alignment also exercise word loads on unaligned
			// backing storage. No scan may skip the first non-ASCII byte.
			text := strings.Repeat("a", 16+prefix) + marker + strings.Repeat("b", 257)
			for offset := 0; offset < 16; offset++ {
				view := text[offset:]
				wantPrefix := 0
				for wantPrefix < len(view) && view[wantPrefix] < 128 {
					wantPrefix++
				}
				if got := stringASCIIUnitPrefix(view); got != wantPrefix {
					t.Fatalf("prefix=%d offset=%d marker=%x: ASCII prefix %d, want %d", prefix, offset, marker, got, wantPrefix)
				}
				if got, want := CountStringCodeUnits(view), referenceCodeUnits(view); got != want {
					t.Fatalf("prefix=%d offset=%d marker=%x: UTF-16 length %d, want %d", prefix, offset, marker, got, want)
				}
			}
		}
	}
}

func TestASCIIWordScanExactSliceLengthsAndHighBytes(t *testing.T) {
	for _, length := range []int{0, 1, 7, 8, 9, 31, 32, 33, 127, 128, 129, 255, 256, 257, 511, 512, 513, 1024} {
		ascii := strings.Repeat("a", length)
		if stringASCIIUnitPrefix(ascii) != length || CountStringCodeUnits(ascii) != length {
			t.Fatalf("ASCII slice length %d", length)
		}
		for _, position := range []int{0, 7, 8, 31, 32, 127, 128, 255, 256, length / 2, length - 1} {
			if position < 0 || position >= length {
				continue
			}
			for byteValue := 128; byteValue < 256; byteValue++ {
				bytes := []byte(ascii)
				bytes[position] = byte(byteValue)
				view := string(bytes)
				if got := stringASCIIUnitPrefix(view); got != position {
					t.Fatalf("length=%d position=%d byte=%x: ASCII prefix %d", length, position, byteValue, got)
				}
				if got, want := CountStringCodeUnits(view), referenceCodeUnits(view); got != want {
					t.Fatalf("length=%d position=%d byte=%x: UTF-16 length %d, want %d", length, position, byteValue, got, want)
				}
			}
		}
	}
}

var scannedStringUnits int

func BenchmarkCountStringCodeUnitsRunWidths(b *testing.B) {
	for _, input := range []struct{ name, text string }{
		{"ASCII", strings.Repeat("a", 10000)},
		{"Unicode", strings.Repeat("a😀é", 1000)},
		{"ShortASCII", strings.Repeat("aaaaaé", 1000)},
		{"MediumASCII", strings.Repeat(strings.Repeat("a", 40)+"é", 1000)},
		{"Mixed", strings.Repeat("a", 5000) + "😀" + strings.Repeat("a", 5000)},
	} {
		b.Run(input.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				scannedStringUnits = CountStringCodeUnits(input.text)
			}
		})
	}
}
