package c1c

import (
	"strings"
	"testing"
)

func TestBodyLiteralsKeepPublishedFieldABI(t *testing.T) {
	previous := `package purescript
func First(x struct { value string /*gopurs:record-v1:00760061006c00750065*/ }) string { return x.value + "old" }
func Second() int { return 12 }
`
	current := strings.Replace(strings.Replace(previous, `"old"`, `"new\nvalue"`, 1), "return 12", "return 42", 1)
	changes, err := CompareBodyLiterals(previous, current)
	if err != nil {
		t.Fatal(err)
	}
	for published, expected := range map[string]string{
		`package p_one
func First(x struct { P_value string }) string { return x.P_value + "old" }
`: `package p_one
func First(x struct { P_value string }) string { return x.P_value + "new\nvalue" }
`,
		"package p_two\nfunc Second() int { return 12 }\n": "package p_two\nfunc Second() int { return 42 }\n",
	} {
		actual, err := changes.Apply(published)
		if err != nil || actual != expected {
			t.Fatalf("literal patch=%q want=%q error=%v", actual, expected, err)
		}
	}
}

func TestBodyLiteralsRejectGraphTypesMetadataAndStalePlan(t *testing.T) {
	previous := "package purescript\n// producer metadata\nconst K = 1\nfunc Main(x int) string { return Other(x) + \"old\" }\n"
	for _, current := range []string{
		strings.Replace(previous, "Other(x)", "Changed(x)", 1),
		strings.Replace(previous, "x int", "x int64", 1),
		strings.Replace(previous, "K = 1", "K = 2", 1),
		strings.Replace(previous, "producer metadata", "different record metadata", 1),
	} {
		if _, err := CompareBodyLiterals(previous, current); err == nil {
			t.Fatalf("accepted nonliteral edit: %s", current)
		}
	}
	changes, err := CompareBodyLiterals(previous, strings.Replace(previous, `"old"`, `"new"`, 1))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := changes.Apply(strings.Replace(previous, `"old"`, `"stale"`, 1)); err == nil {
		t.Fatal("accepted stale published literal")
	}
}
