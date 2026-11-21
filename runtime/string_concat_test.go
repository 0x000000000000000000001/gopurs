package gopurs_runtime

import (
	"math/rand"
	"testing"
)

func TestConcatStringsMatchesStrictAppend(t *testing.T) {
	fragments := []string{"", "a", "😀é", "\xed\xa0\x80", "\xed\xaf\xbf", "\xed\xb0\x80", "\xed\xbf\xbf", "\xff", "\xed\xa0", "\xed\xa0\x80\xed\xb0\x80"}
	check := func(seed string, parts []string) {
		t.Helper()
		want := seed
		boxed := make([]Value, len(parts))
		for i, part := range parts {
			want = ConcatString(want, part)
			boxed[i] = Str(part)
		}
		if got := ConcatStrings(seed, parts, func(s string) string { return s }); got != want {
			t.Fatalf("seed %q parts %q: got %q, want %q", seed, parts, got, want)
		}
		if got := ConcatStrings(seed, boxed, Value.StrVal); got != want {
			t.Fatalf("boxed seed %q parts %q: got %q, want %q", seed, parts, got, want)
		}
	}
	for _, seed := range fragments {
		check(seed, nil)
		for _, a := range fragments {
			for _, b := range fragments {
				check(seed, []string{a, "", b})
			}
		}
	}
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 10000; i++ {
		parts := make([]string, rng.Intn(30))
		for j := range parts {
			if rng.Intn(2) == 0 {
				parts[j] = fragments[rng.Intn(len(fragments))]
			} else {
				buf := make([]byte, rng.Intn(20))
				rng.Read(buf)
				parts[j] = string(buf)
			}
		}
		check(fragments[rng.Intn(len(fragments))], parts)
	}
}
