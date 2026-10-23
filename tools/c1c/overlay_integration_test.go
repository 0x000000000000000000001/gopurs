package c1c

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestOverlayEndToEnd(t *testing.T) {
	root := t.TempDir()
	write := func(path, content string) {
		target := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	write("go.mod", "module example.com/x\n\ngo 1.22\n")
	oldSource := `package old
func Helper() string { return "new" }
func Get() string { return "old" }
`
	write("old/old.go", oldSource)
	newSource := `package purescript
func Get() string { return Helper() }
`
	overlay, err := OverlaySource("overlay", "example.com/x/old", oldSource, newSource)
	if err != nil {
		t.Fatal(err)
	}
	write("overlay/overlay.go", overlay)
	entry := `package main
import (
 "fmt"
 old "example.com/x/old"
)
func main() { fmt.Print(old.Get()) }
`
	rewritten, changed, err := RewritePackageReferences(entry, "example.com/x/old", "example.com/x/overlay", []string{"Get"})
	if err != nil {
		t.Fatal(err)
	}
	if !changed {
		t.Fatal("entry was not rewritten")
	}
	write("entry/main.go", rewritten)
	command := exec.Command("go", "run", "./entry")
	command.Dir = root
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("go run: %v: %s", err, output)
	}
	if string(output) != "new" {
		t.Fatalf("output %q", output)
	}
}
