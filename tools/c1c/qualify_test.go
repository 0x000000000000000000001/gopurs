package c1c

import (
	"strings"
	"testing"
)

func TestBuildOverlayQualifiesCrossPackageSymbols(t *testing.T) {
	source := `package purescript

import (
	r "gopurs/output/gopurs_runtime"
)

func Get_Main_main() r.Value {
	return Call_Helper_run(r.Str("x"))
}
`
	symbols := map[string]string{
		"Call_Helper_run":    "p_helper",
		"Constructor_Helper": "p_helper",
	}
	out, dependencies, err := BuildOverlaySource("p_overlay", source, symbols, func(pkg string) string {
		return "gopurs/output/purescript/" + pkg
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(dependencies) != 1 || dependencies[0].Package != "p_helper" {
		t.Fatalf("dependencies: %+v", dependencies)
	}
	if !strings.Contains(out, dependencies[0].Alias+".Call_Helper_run(") {
		t.Fatalf("missing qualification: %s", out)
	}
	if !strings.Contains(out, "package p_overlay") {
		t.Fatalf("package not rewritten: %s", out)
	}
	if !strings.Contains(out, `"gopurs/output/purescript/p_helper"`) {
		t.Fatalf("missing import: %s", out)
	}
	if strings.Count(out, "gopurs/output/gopurs_runtime") != 1 {
		t.Fatalf("existing import changed: %s", out)
	}
}

func TestBuildOverlayKeepsLocalDeclarationsAndShadow(t *testing.T) {
	source := `package purescript

func Helper() string { return "local" }

func Get_Main_main() string {
	Helper := func() string { return "shadow" }
	return Helper()
}
`
	out, dependencies, err := BuildOverlaySource("p_overlay", source, map[string]string{"Helper": "p_other"}, func(pkg string) string {
		return "gopurs/output/purescript/" + pkg
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(dependencies) != 0 {
		t.Fatalf("local declaration was qualified: %+v", dependencies)
	}
	if strings.Contains(out, "p_other") {
		t.Fatalf("unexpected package alias in source: %s", out)
	}
}

func TestRewriteUsesFreeAlias(t *testing.T) {
	source := `package p_entry

import old "gopurs/output/purescript/p_old"

var c1c_overlay int

func Run() { old.Get_Changed() }
`
	out, changed, err := RewritePackageReferencesAs(source, "gopurs/output/purescript/p_old", "gopurs/output/purescript/p_new", "c1c_overlay", []string{"Get_Changed"})
	if err != nil || !changed {
		t.Fatalf("rewrite: %v %v", changed, err)
	}
	if !strings.Contains(out, "c1c_overlay.Get_Changed()") {
		t.Fatalf("missing selector: %s", out)
	}
}
