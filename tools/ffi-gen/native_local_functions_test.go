package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestNativeLocalFunctionsAdmission(t *testing.T) {
	prefix := `package p
import r "gopurs/output/gopurs_runtime"
type V = r.Value
var global V
`
	tests := []struct {
		name, body string
		accepted   bool
	}{
		{"direct", `func test(){ f:=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)) }`, true},
		{"alias", `func test(){ var f V; f=r.Func2(func(x,y V)V{return x}); g:=f; h:=g; _=r.Apply2(h,r.Int(1),r.Int(2)) }`, true},
		{"unused", `func test(){ var f V; f=r.Func(func(v V)V{panic("unused")}); g:=f; _=g }`, true},
		{"recursive", `func test(){ var f V; f=r.Func(func(v V)V{if v.IntVal==0{return v};return r.Apply(f,r.Int(v.IntVal-1))}); _=r.Apply(f,r.Int(3)) }`, true},
		{"loop slot", `func test(){ var f V; for i:=0;i<3;i++ {f=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1))} }`, true},
		{"join outside assignment block", `func test(b bool){ var f V; if b {f=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)); goto done}; done: _=f }`, true},
		{"jump in wrapper body", `func test(){ f:=r.Func(func(v V)V{goto done; done:return v}); _=r.Apply(f,r.Int(1)) }`, true},
		{"independent function label", `func test(){ unused:=func(){goto done; done:return}; _=unused; f:=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)); done:return }`, true},
		{"captured call", `func test() func()V{ f:=r.Func(func(v V)V{return v}); return func()V{return r.Apply(f,r.Int(1))} }`, true},
		{"return", `func test()V{ f:=r.Func(func(v V)V{return v}); return f }`, false},
		{"alias return", `func test()V{ f:=r.Func(func(v V)V{return v}); g:=f; return g }`, false},
		{"captured value", `func test()func()V{ f:=r.Func(func(v V)V{return v}); return func()V{return f} }`, false},
		{"stored", `func test(){ f:=r.Func(func(v V)V{return v}); _=[]V{f} }`, false},
		{"unknown call", `func test(){ f:=r.Func(func(v V)V{return v}); consume(f) }`, false},
		{"partial", `func test(){ f:=r.Func2(func(x,y V)V{return x}); _=r.Apply(f,r.Int(1)) }`, false},
		{"overapplied", `func test(){ f:=r.Func(func(v V)V{return v}); _=r.Apply2(f,r.Int(1),r.Int(2)) }`, false},
		{"reassigned", `func test(){ f:=r.Func(func(v V)V{return v}); g:=f; f=g; _=r.Apply(f,r.Int(1)) }`, false},
		{"initialized", `func test(){ var f V = r.Int(1); f=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)) }`, false},
		{"declaration group", `func test(){ var f,g V; f=r.Func(func(v V)V{return v}); _=g; _=r.Apply(f,r.Int(1)) }`, false},
		{"available local type", `func test(){ type Local=r.Value; var f r.Value; f=r.Func(func(v Local)Local{return v}); _=r.Apply(f,r.Int(1)) }`, true},
		{"later local type", `func test(){ var f r.Value; type Local=r.Value; f=r.Func(func(v Local)Local{return v}); _=r.Apply(f,r.Int(1)) }`, false},
		{"mixed type availability", `func test(){ type Early=r.Value; var f r.Value; type Late=r.Value; f=r.Func(func(v Late)Early{return v}); _=r.Apply(f,r.Int(1)) }`, false},
		{"nested local type", `func test(){ var f r.Value; {type Local=r.Value; f=r.Func(func(v Local)Local{return v}); _=r.Apply(f,r.Int(1))} }`, false},
		{"parameter", `func test(f V){ f=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)) }`, false},
		{"named result", `func test()(f V){ f=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)); return }`, false},
		{"call before assignment", `func test(){ var f V; _=r.Apply(f,r.Int(1)); f=r.Func(func(v V)V{return v}) }`, false},
		{"alias before assignment", `func test(){ var f V; g:=f; f=r.Func(func(v V)V{return v}); _=r.Apply(g,r.Int(1)) }`, false},
		{"closure before assignment", `func test(){ var f V; g:=func()V{return r.Apply(f,r.Int(1))}; _=g(); f=r.Func(func(v V)V{return v}) }`, false},
		{"conditional assignment", `func test(b bool){ var f V; if b {f=r.Func(func(v V)V{return v})}; _=r.Apply(f,r.Int(1)) }`, false},
		{"nested assignment", `func test(){ var f V; assign:=func(){f=r.Func(func(v V)V{return v})}; assign(); _=r.Apply(f,r.Int(1)) }`, false},
		{"goto", `func test(){ var f V; goto use; f=r.Func(func(v V)V{return v}); use: _=r.Apply(f,r.Int(1)) }`, false},
		{"global", `func test(){ global=r.Func(func(v V)V{return v}); _=r.Apply(global,r.Int(1)) }`, false},
		{"inspect", `func test(){ f:=r.Func(func(v V)V{return v}); _=f.Type }`, false},
		{"shadow runtime", `func test(r CustomRuntime){ f:=r.Func(func(v V)V{return v}); _=r.Apply(f,r.Int(1)) }`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := prefix + tt.body
			result, err := optimizeNativeLocalFunctions(source)
			if err != nil {
				t.Fatal(err)
			}
			if (result != source) != tt.accepted {
				t.Fatalf("admission=%v, want %v:\n%s", result != source, tt.accepted, result)
			}
			again, err := optimizeNativeLocalFunctions(result)
			if err != nil || again != result {
				t.Fatalf("pass not idempotent: %v", err)
			}
			if tt.accepted {
				f, err := parser.ParseFile(token.NewFileSet(), "result.go", result, 0)
				if err != nil {
					t.Fatal(err)
				}
				ast.Inspect(f, func(node ast.Node) bool {
					if call, ok := node.(*ast.CallExpr); ok {
						if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
							if strings.HasPrefix(selector.Sel.Name, "Func") || strings.HasPrefix(selector.Sel.Name, "Apply") {
								t.Errorf("retained runtime wrapper/application: %s", result)
							}
						}
					}
					return true
				})
			}
		})
	}
	for _, source := range []string{
		`package p; import r "other/runtime"; func f(){v:=r.Func(func(x int)int{return x});_=r.Apply(v,1)}`,
		`package p; import . "gopurs/output/gopurs_runtime"; func f(){v:=Func(func(x Value)Value{return x});_=Apply(v,Int(1))}`,
	} {
		got, err := optimizeNativeLocalFunctions(source)
		if err != nil || got != source {
			t.Fatalf("changed unrelated source: %v %s", err, got)
		}
	}
	if _, err := optimizeNativeLocalFunctions("not Go"); err == nil {
		t.Fatal("invalid source accepted")
	}
}

func TestNativeLocalFunctionsLexicalBindings(t *testing.T) {
	source := `package p
import r "gopurs/output/gopurs_runtime"
func f(){
 v:=r.Func(func(x r.Value)r.Value{return x})
 {v:=r.Int(3);_=v.Type}
 _=r.Apply(v,r.Int(1))
}
`
	result, err := optimizeNativeLocalFunctions(source)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(result, "r.Func(") || strings.Contains(result, "r.Apply(") || !strings.Contains(result, "{v:=r.Int(3);_=v.Type}") {
		t.Fatalf("shadowing was not resolved lexically: %s", result)
	}
}

func TestNativeLocalFunctionsExecution(t *testing.T) {
	const source = `package main
import("encoding/json";"os";"runtime";r "gopurs/output/gopurs_runtime")
type V=r.Value
var customBoxes int
func customBox(f func(V)V)V { customBoxes++;return r.Func(f) }
func direct(x int64)int64 { var f V;f=r.Func(func(v V)V{return r.Int(v.IntVal+x)});g:=f;h:=g;return r.Apply(h,r.Int(4)).IntVal }
func unused()int64 { var f V;f=r.Func(func(v V)V{panic("never")});_ = f;return 12 }
func escape(seed int64)V { var f V;f=r.Func(func(v V)V{return r.Int(seed+v.IntVal)});return f }
func stored(seed int64)int64 {var f V;f=r.Func(func(v V)V{return r.Int(seed+v.IntVal)});box:=[]V{f};return r.Apply(box[0],r.Int(5)).IntVal}
func partial()int64 {var f V;f=r.Func2(func(a,b V)V{return r.Int(a.IntVal+b.IntVal)});return r.Apply(r.Apply(f,r.Int(3)),r.Int(9)).IntVal}
func reassigned()int64 {var f V;f=r.Func(func(v V)V{return r.Int(v.IntVal+1)});old:=f;f=r.Func(func(v V)V{return r.Int(v.IntVal+100)});return r.Apply(old,r.Int(1)).IntVal+r.Apply(f,r.Int(1)).IntVal}
func captureValue(seed int64)V {var f V;f=r.Func(func(v V)V{return r.Int(seed+v.IntVal)});return r.Func(func(_ V)V{return f})}
func captureCall(seed int64)V {var f V;f=r.Func(func(v V)V{return r.Int(seed+v.IntVal)});return r.Func(func(v V)V{return r.Apply(f,v)})}
func recursion(n int64)int64 {var f V;f=r.Func(func(v V)V{if v.IntVal==0{return r.Int(7)};return r.Apply(f,r.Int(v.IntVal-1))});return r.Apply(f,r.Int(n)).IntVal}
func ordered()[]int64 {var trace []int64;arg:=func(v int64)V{trace=append(trace,v);return r.Int(v)};var f V;f=r.Func2(func(a,b V)V{trace=append(trace,3);return r.Int(a.IntVal+b.IntVal)});v:=r.Apply2(f,arg(1),arg(2));return append(trace,v.IntVal)}
func shadowed()int64 {r:=struct{Func func(func(V)V)V}{Func:customBox};f:=r.Func(func(v V)V{return v});_ = f;return int64(customBoxes)}
func multiDecl()int64 {var a,b V;a=r.Func(func(v V)V{return r.Int(v.IntVal+2)});_ = b;return r.Apply(a,r.Int(4)).IntVal}
func main(){a:=captureCall(17);b:=captureValue(21);e:=escape(31);for i:=0;i<100;i++{direct(int64(i));runtime.GC()};values:=map[string]any{"direct":direct(10),"unused":unused(),"escaped":r.Apply(e,r.Int(2)).IntVal,"stored":stored(10),"partial":partial(),"reassigned":reassigned(),"captureValue":r.Apply(r.Apply(b,r.Value{}),r.Int(3)).IntVal,"captureCall":r.Apply(a,r.Int(4)).IntVal,"recursive":recursion(50),"ordered":ordered(),"shadowed":shadowed(),"multiDecl":multiDecl()};json.NewEncoder(os.Stdout).Encode(values)}
`
	output, err := optimizeNativeLocalFunctions(source)
	if err != nil {
		t.Fatal(err)
	}
	directory := t.TempDir()
	packagePut(t, directory, "go.mod", "module gopurs/output\n\ngo 1.22\n")
	runtime, err := os.ReadFile("../../runtime/runtime.go")
	if err != nil {
		t.Fatal(err)
	}
	packagePut(t, directory, "gopurs_runtime/runtime.go", string(runtime))
	var baseline string
	for i, text := range []string{source, output} {
		packagePut(t, directory, "main.go", text)
		command := exec.Command("go", "run", ".")
		command.Dir = directory
		command.Env = append(os.Environ(), "GOWORK=off", "GOTOOLCHAIN=local")
		result, err := command.CombinedOutput()
		if err != nil {
			t.Fatalf("execute variant %d: %v\n%s", i, err, result)
		}
		if i == 0 {
			baseline = string(result)
		} else if string(result) != baseline {
			t.Fatalf("changed results/order/GC lifetime: before %s after %s", baseline, result)
		}
	}
	want := `{"captureCall":21,"captureValue":24,"direct":14,"escaped":33,"multiDecl":6,"ordered":[1,2,3,3],"partial":12,"reassigned":103,"recursive":7,"shadowed":1,"stored":15,"unused":12}` + "\n"
	if baseline != want {
		t.Fatalf("unexpected semantic control result: %s", baseline)
	}
}
