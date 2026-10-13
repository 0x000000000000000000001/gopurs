package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
	"syscall"
)

const compilerSnapshotEnv = "GOPURS_COMPILER_SNAPSHOT"

type compilerSnapshot struct {
	Schema    int     `json:"schema"`
	Host      string  `json:"host"`
	Resources []Input `json:"resources"`
}

func readCompilerSnapshot() (string, compilerSnapshot, error) {
	root := os.Getenv(compilerSnapshotEnv)
	var manifest compilerSnapshot
	if root == "" || !filepath.IsAbs(root) {
		return "", manifest, fmt.Errorf("compiler snapshot is required")
	}
	data, err := os.ReadFile(filepath.Join(root, "compiler.json"))
	if err != nil {
		return "", manifest, err
	}
	if err := strictJSON(data, &manifest); err != nil {
		return "", manifest, err
	}
	if manifest.Schema != 1 || (manifest.Host != "go" && manifest.Host != "js" && manifest.Host != "rust") {
		return "", manifest, fmt.Errorf("compiler snapshot version/host")
	}
	for _, input := range manifest.Resources {
		if input.Kind != "file" || filepath.IsAbs(input.Path) || filepath.Clean(input.Path) != input.Path || input.Path == ".." || len(input.Path) >= 3 && input.Path[:3] == "../" {
			return "", manifest, fmt.Errorf("invalid compiler resource")
		}
	}
	return root, manifest, nil
}

func CompilerInputs() ([]Input, error) {
	root, manifest, err := readCompilerSnapshot()
	if err != nil {
		return nil, err
	}
	inputs := append([]Input{}, manifest.Resources...)
	for i := range inputs {
		inputs[i].Path = filepath.Join(root, inputs[i].Path)
	}
	return inputs, nil
}

// Re-exec BEFORE any cache decision. Even when the installation was replaced
// after this process loaded, the next process executes exactly the copied bytes.
func FreezeNativeCompiler() error {
	return FreezeCompiler("go")
}

func FreezeCompiler(host string) error {
	name := "gopurs-native"
	if host == "rust" {
		name = "gopurs-rust"
	} else if host != "go" {
		return fmt.Errorf("unknown native compiler host")
	}
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	executable, err = filepath.EvalSymlinks(executable)
	if err != nil {
		return err
	}
	if os.Getenv(compilerSnapshotEnv) != "" {
		root, manifest, err := readCompilerSnapshot()
		if err != nil {
			return err
		}
		if manifest.Host != host || executable != filepath.Join(root, name) {
			return fmt.Errorf("compiler is not executing its frozen artifact")
		}
		return nil
	}
	bytes, err := os.ReadFile(executable)
	if err != nil {
		return err
	}
	root, err := os.MkdirTemp("", "gopurs-compiler-")
	if err != nil {
		return err
	}
	canonical, err := filepath.EvalSymlinks(root)
	if err != nil {
		os.RemoveAll(root)
		return err
	}
	root = canonical
	target := filepath.Join(root, name)
	if err := os.WriteFile(target, bytes, 0500); err != nil {
		os.RemoveAll(root)
		return err
	}
	manifest, _ := json.Marshal(compilerSnapshot{Schema: 1, Host: host, Resources: []Input{{Name: "compiler", Kind: "file", Path: name}}})
	if err := os.WriteFile(filepath.Join(root, "compiler.json"), manifest, 0400); err != nil {
		os.RemoveAll(root)
		return err
	}
	// The generated entry point may start PPROF before reaching Main. Its
	// interval timer survives exec while signal handlers do not; a pending
	// SIGPROF can therefore kill the new image before the Go runtime starts.
	// The frozen entry point starts its own profile after this timer is stopped.
	pprof.StopCPUProfile()
	err = syscall.Exec(target, append([]string{target}, os.Args[1:]...), append(os.Environ(), compilerSnapshotEnv+"="+root))
	os.RemoveAll(root)
	return err
}

func ReleaseCompilerSnapshot() {
	root, _, err := readCompilerSnapshot()
	if err == nil {
		os.RemoveAll(root)
	}
}
