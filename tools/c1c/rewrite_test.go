package c1c

import "strings"
import "testing"

func TestRewritePackageReferences(t *testing.T) {
	source := `package p_entry

import old "gopurs/output/purescript/p_old"

func Run() { old.Get_Changed() ; old.Other() }
`
	out, changed, err := RewritePackageReferences(source, "gopurs/output/purescript/p_old", "gopurs/output/purescript/p_new", []string{"Get_Changed"})
	if err != nil {
		t.Fatal(err)
	}
	if !changed {
		t.Fatal("expected change")
	}
	if !strings.Contains(out, `c1c_overlay "gopurs/output/purescript/p_new"`) {
		t.Fatalf("missing import: %s", out)
	}
	if !strings.Contains(out, "c1c_overlay.Get_Changed()") || strings.Contains(out, "c1c_overlay.Other()") {
		t.Fatalf("bad rewrite: %s", out)
	}
	if err := ValidateRewritten(out); err != nil {
		t.Fatal(err)
	}
}
