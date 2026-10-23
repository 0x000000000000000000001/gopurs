package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// File transport for the same pure package emitter. The driver freezes these
// inputs and owns publication; no graph or Go source crosses the host ABI.
type PackageSourceFile struct {
	Owner        string         `json:"owner"`
	Module       string         `json:"moduleName"`
	Path         string         `json:"path"`
	Source       string         `json:"source"`
	SHA256       string         `json:"sha256"`
	Initializers []SourceSymbol `json:"initializers"`
	Records      []SourceRecord `json:"records"`
}

type PackageExportFile struct {
	Path   string `json:"path"`
	Source string `json:"source"`
	SHA256 string `json:"sha256"`
}

type PackageFilesRequest struct {
	Format      string              `json:"format"`
	PackagePath string              `json:"packagePath"`
	GoVersion   string              `json:"goVersion"`
	Arch        string              `json:"arch"`
	Files       []PackageSourceFile `json:"files"`
	Exports     []PackageExportFile `json:"exports"`
	Entries     []SourceSymbol      `json:"entries"`
	Signatures  map[string]string   `json:"signatures,omitempty"`
	Output      string              `json:"output"`
}

type EmittedPackageFile struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Module string `json:"moduleName"`
	SHA256 string `json:"sha256"`
	Bytes  int    `json:"bytes"`
}

type PackageFilesResult struct {
	Format   string               `json:"format"`
	PlanHash string               `json:"planHash"`
	Packages []GoPackage          `json:"packages"`
	Files    []EmittedPackageFile `json:"files"`
	Entries  []SymbolPlacement    `json:"entries"`
}

func readPackageInput(source, digest string) ([]byte, error) {
	if !filepath.IsAbs(source) || len(digest) != 64 {
		return nil, fmt.Errorf("package input requires an absolute path and SHA-256: %s", source)
	}
	data, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}
	if sourceBytesHash(data) != digest {
		return nil, fmt.Errorf("package input changed: %s", source)
	}
	return data, nil
}

func EmitPackageFiles(request PackageFilesRequest) (*PackageFilesResult, error) {
	if request.Format != "gopurs/go-package-files-v1" || !filepath.IsAbs(request.Output) {
		return nil, fmt.Errorf("invalid Go package file request")
	}
	pure := PackageRequest{BindRequest: BindRequest{PackagePath: request.PackagePath,
		GoVersion: request.GoVersion, Arch: request.Arch, Exports: map[string][]byte{}}, Entries: request.Entries}
	owners := map[string]PackageSourceFile{}
	for _, file := range request.Files {
		if _, exists := owners[file.Path]; exists {
			return nil, fmt.Errorf("duplicate Go source: %s", file.Path)
		}
		data, err := readPackageInput(file.Source, file.SHA256)
		if err != nil {
			return nil, err
		}
		owners[file.Path] = file
		pure.Files = append(pure.Files, SourceInput{Owner: file.Owner, Module: file.Module, Path: file.Path,
			Content: string(data), Initializers: file.Initializers, Records: file.Records})
	}
	for _, file := range request.Exports {
		if _, exists := pure.Exports[file.Path]; exists || file.Path == "" {
			return nil, fmt.Errorf("duplicate or empty Go export path: %s", file.Path)
		}
		data, err := readPackageInput(file.Source, file.SHA256)
		if err != nil {
			return nil, err
		}
		pure.Exports[file.Path] = data
	}
	emitted, err := PackageSources(pure)
	if err != nil {
		return nil, err
	}
	result := &PackageFilesResult{Format: request.Format, PlanHash: emitted.PlanHash,
		Packages: emitted.Packages, Entries: emitted.Entries, Files: []EmittedPackageFile{}}
	// A caller supplies a new private destination. Never overwrite an existing
	// directory, including an interrupted generation or a public output tree.
	if err := os.Mkdir(request.Output, 0700); err != nil {
		return nil, err
	}
	complete := false
	defer func() {
		if !complete {
			os.RemoveAll(request.Output)
		}
	}()
	for _, file := range emitted.Files {
		directory, base, ok := strings.Cut(file.Path, "/")
		if !ok || directory == "" || directory == "." || directory == ".." ||
			base == "" || path.Base(base) != base || strings.ContainsAny(file.Path, "\\\x00") {
			return nil, fmt.Errorf("invalid emitted Go package path: %s", file.Path)
		}
		kind, module := "order", ""
		if directory == "init_all" || directory == "init_external" {
			kind = "package-init"
		} else if base == packageBundleName {
			kind = "package"
		} else if owner, found := owners[base]; found {
			kind, module = "package", owner.Module
			if strings.HasPrefix(owner.Owner, "ffi/") {
				kind = "ffi-package"
			}
		} else if base != "zz_gopurs_order.go" {
			return nil, fmt.Errorf("emitted Go file has no producer: %s", file.Path)
		}
		if err := os.MkdirAll(filepath.Join(request.Output, directory), 0700); err != nil {
			return nil, err
		}
		data := []byte(file.Content)
		if err := os.WriteFile(filepath.Join(request.Output, filepath.FromSlash(file.Path)), data, 0600); err != nil {
			return nil, err
		}
		result.Files = append(result.Files, EmittedPackageFile{file.Path, kind, module, sourceBytesHash(data), len(data)})
	}
	complete = true
	return result, nil
}

func PackageFiles(content string) (string, error) {
	var request PackageFilesRequest
	decoder := json.NewDecoder(strings.NewReader(content))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return "", err
	}
	if decoder.Decode(new(any)) != io.EOF {
		return "", fmt.Errorf("trailing Go package file request")
	}
	result, err := EmitPackageFiles(request)
	if err != nil {
		return "", err
	}
	var encoded bytes.Buffer
	encoder := json.NewEncoder(&encoded)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(result)
	return strings.TrimSuffix(encoded.String(), "\n"), err
}
