package c1c

import (
	"strings"
	"testing"
)

func TestPartitionSourceQualifiesTopLevelAndKeepsLocals(t *testing.T) {
	source := `package purescript
import old "fmt"
func Main() string { return old.Sprint(Other()) }
func Local() string { Other := "local"; return Other }
func Other() string { return "other" }
`
	placements := map[string]string{"Main": "a", "Local": "a", "Other": "b"}
	out, err := PartitionSource(source, "a", placements, placements,
		map[string]string{"fmt": "fmt", "example/b": "dep"}, func(pkg string) string { return "example/" + pkg })
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "fmt.Sprint(dep.Other())") || !strings.Contains(out, `Other := "local"; return Other`) || strings.Contains(out, "func Other") {
		t.Fatalf("invalid partition:\n%s", out)
	}
}

func TestPartitionSourceRejectsAmbiguousGroupAndShadowedImport(t *testing.T) {
	for _, source := range []string{
		"package purescript\nvar ( A int; B int )\n",
		"package purescript\nfunc A() int { dep := 1; return dep + B() }\nfunc B() int { return 1 }\n",
	} {
		placements := map[string]string{"A": "a", "B": "b"}
		_, err := PartitionSource(source, "a", placements, placements,
			map[string]string{"example/b": "dep"}, func(pkg string) string { return "example/" + pkg })
		if err == nil {
			t.Fatalf("accepted ambiguous source: %s", source)
		}
	}
}
