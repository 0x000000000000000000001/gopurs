package c1c

import (
	"strings"
	"testing"
)

const missingFixture = `package purescript
import rt "gopurs/output/gopurs_runtime"
//gopurs:missing-ffi-v1:{"binding":"Missing","module":"App.Foreign","symbol":"prefix","path":"src/App/Foreign.go"}
var Missing = rt.Func(func(_ rt.Value) rt.Value { panic("missing") })
func Get_App_Foreign_prefix() rt.Value { return Missing }
`

func TestMissingFFIReachability(t *testing.T) {
	for _, test := range []struct {
		name   string
		source string
		fails  bool
	}{
		{"direct", `func Get_Main_main() { Apply(Get_App_Foreign_prefix()) }`, true},
		{"transitive", `func Get_Main_main() { helper() }; func helper() { Apply(Missing) }`, true},
		{"callback", `func Get_Main_main() { callback(func() { Apply(Missing) }) }`, true},
		{"returned closure", `func Get_Main_main() func() { return func() { Apply(Missing) } }`, true},
		{"unused function", `func unused() { Apply(Missing) }; func Get_Main_main() {}`, false},
		{"shadowed local", `func Get_Main_main() { Missing := func() {}; Missing() }`, false},
		{"shadowed parameter", `func helper(Missing func()) { Missing() }; func Get_Main_main() { helper(func(){}) }`, false},
		{"selector field", `type T struct { Missing int }; func Get_Main_main() { var t T; consume(t.Missing) }`, false},
		{"literal field", `type T struct { Missing int }; func Get_Main_main() { consume(T{Missing: 1}) }`, false},
		{"init call", `func init() { Apply(Missing) }; func Get_Main_main() {}`, true},
		{"initialized effect", `var boot = Apply(Missing); func Get_Main_main() {}`, true},
		{"stored closure", `var dormant = func() { Apply(Missing) }; func Get_Main_main() {}`, false},
		{"used stored closure", `var dormant = func() { Apply(Missing) }; func Get_Main_main() { dormant() }`, true},
		{"closure factory", `var dormant = factory(); func factory() func() { return func() { Apply(Missing) } }; func Get_Main_main() {}`, false},
		{"used closure factory", `var dormant = factory(); func factory() func() { return func() { Apply(Missing) } }; func Get_Main_main() { dormant() }`, true},
		{"runtime wrapper", `var dormant = rt.Func(func(_ rt.Value) rt.Value { return Missing }); func Get_Main_main() {}`, false},
		{"used runtime wrapper", `var dormant = rt.Func(func(_ rt.Value) rt.Value { return Missing }); func Get_Main_main() { Apply(dormant) }`, true},
		{"eager callback", `var boot = callback(func() { Apply(Missing) }); func Get_Main_main() {}`, true},
		{"immediate closure", `var boot = func() int { Apply(Missing); return 1 }(); func Get_Main_main() {}`, true},
		{"parenthesized closure", `var boot = (func() int { Apply(Missing); return 1 })(); func Get_Main_main() {}`, true},
		{"generic initializer", `var boot = initialize[int](); func initialize[T any]() int { Apply(Missing); return 1 }; func Get_Main_main() {}`, true},
		{"blank initializers", `var _ = Apply(Missing); var _ = 1; func Get_Main_main() {}`, true},
		{"intrinsic eliminated", `func Get_Main_main() int { return 1 + 2 }`, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			files := map[string]string{"ffi.go": missingFixture, "main.go": "package purescript\nimport rt \"gopurs/output/gopurs_runtime\"\n" + test.source}
			err := CheckMissingFFI(files, []string{"Get_Main_main"})
			if (err != nil) != test.fails {
				t.Fatalf("failure=%v, expected %v", err, test.fails)
			}
			if err != nil {
				for _, context := range []string{"App.Foreign", "src/App/Foreign.go", "prefix", "--ffi", "workspace.extraPackages"} {
					if !strings.Contains(err.Error(), context) {
						t.Fatalf("missing %q in %v", context, err)
					}
				}
			}
		})
	}
}

func TestMissingFFISelectedEntries(t *testing.T) {
	files := map[string]string{"ffi.go": missingFixture, "main.go": `package purescript
func Get_Server_main() {}
func Get_Client_main() { Apply(Missing) }
`}
	if err := CheckMissingFFI(files, []string{"Get_Server_main"}); err != nil {
		t.Fatal(err)
	}
	if err := CheckMissingFFI(files, []string{"Get_Server_main", "Get_Client_main"}); err == nil {
		t.Fatal("missing client FFI accepted")
	}
	if err := CheckMissingFFI(files, nil); err != nil {
		t.Fatal("library with no selected executable", err)
	}
}

func TestMissingFFIRuntimeStorageConstructors(t *testing.T) {
	for _, test := range []struct {
		name       string
		expression string
	}{
		{"Func11", `rt.Func11(func(a,b,c,d,e,f,g,h,i,j,k rt.Value) rt.Value { return Missing })`},
		{"FuncAny", `rt.FuncAny(func(_ rt.Value) rt.Value { return Missing })`},
		{"Any", `rt.Any(func() { Apply(Missing) })`},
		{"AnyToValue", `rt.AnyToValue(func() { Apply(Missing) })`},
		{"generic Box", `rt.Box[func()](func() { Apply(Missing) })`},
		{"generic Wrap", `rt.Wrap0[rt.Value](func() rt.Value { return Missing })`},
		{"Record", `rt.Record(map[string]rt.Value{"f": rt.Func(func(_ rt.Value) rt.Value { return Missing })})`},
		{"RecordDict", `rt.RecordDict([]string{"f"}, []rt.Value{Missing})`},
		{"RecordDict5", `rt.RecordDict5("a", "b", "c", "d", "e", Missing, Missing, Missing, Missing, Missing)`},
		{"Array", `rt.Array([]rt.Value{Missing})`},
		{"Constructor", `rt.Constructor("Callback", []rt.Value{Missing})`},
		{"Constructor5", `rt.Constructor5("Callback", Missing, Missing, Missing, Missing, Missing)`},
		{"WithFunctionData", `rt.WithFunctionData(rt.Func(func(_ rt.Value) rt.Value { return Missing }), nil)`},
	} {
		t.Run(test.name, func(t *testing.T) {
			for _, used := range []bool{false, true} {
				body := ""
				if used {
					body = "Apply(stored)"
				}
				files := map[string]string{
					"ffi.go": missingFixture,
					"main.go": "package purescript\nimport rt \"gopurs/output/gopurs_runtime\"\nvar stored = " + test.expression +
						"\nfunc Get_Main_main() { " + body + " }",
				}
				err := CheckMissingFFI(files, []string{"Get_Main_main"})
				if (err != nil) != used {
					t.Fatalf("used=%v, error=%v", used, err)
				}
			}
		})
	}
}

func TestMissingFFIRuntimeStorageStillEvaluatesArguments(t *testing.T) {
	files := map[string]string{
		"ffi.go": missingFixture,
		"main.go": `package purescript
import rt "gopurs/output/gopurs_runtime"
var stored = rt.Record(map[string]rt.Value{"f": rt.Apply(Missing, rt.Value{})})
func Get_Main_main() {}
`,
	}
	if err := CheckMissingFFI(files, []string{"Get_Main_main"}); err == nil {
		t.Fatal("eager missing FFI in stored value accepted")
	}
}

func TestMissingFFIDoesNotParseSourcesWithoutMarkers(t *testing.T) {
	if err := CheckMissingFFI(map[string]string{"unknown.go": "not parsed"}, []string{"Get_Main_main"}); err != nil {
		t.Fatal(err)
	}
}

func TestMissingFFIAnnotationIsValidated(t *testing.T) {
	files := map[string]string{"ffi.go": "package p\n" + MissingFFIMarker + "{}\nfunc Get_Main_main() {}"}
	if err := CheckMissingFFI(files, []string{"Get_Main_main"}); err == nil || !strings.Contains(err.Error(), "annotation") {
		t.Fatal(err)
	}
}
