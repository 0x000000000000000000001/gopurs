package c1c

import (
	"strings"
	"testing"
)

func TestOverlaySource(t *testing.T) {
	old := `package purescript
func Helper() string { return "old" }
`
	new := `package purescript
func Get() string { return Helper() }
`
	out, err := OverlaySource("p_overlay", "gopurs/output/purescript/p_old", old, new)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, `oldpkg "gopurs/output/purescript/p_old"`) {
		t.Fatalf("missing import: %s", out)
	}
	if !strings.Contains(out, "oldpkg.Helper()") {
		t.Fatalf("missing qualification: %s", out)
	}
	if err := CompileCheck("p_overlay", out); err != nil {
		t.Fatal(err)
	}
}
