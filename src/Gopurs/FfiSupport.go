// Handwritten FFI for Gopurs.FfiSupport; keep in sync with the .purs foreign imports.
package Gopurs_FfiSupport

import (
	"fmt"
	"gopurs/output/gopurs_ffi_parser"
)

func ExtractFfiAstImpl(content string, _ any) string {
	result, err := gopurs_ffi_parser.Extract(content)
	if err != nil {
		panic(fmt.Errorf("Go FFI parse failed: %w", err))
	}
	return result
}

func PrepareFfiAstImpl(prefix string, content string, _ any) string {
	result, err := gopurs_ffi_parser.Prepare(prefix, content)
	if err != nil {
		panic(fmt.Errorf("Go FFI parse failed: %w", err))
	}
	return result
}

func BindGoSourceImpl(content string, _ any) string {
	result, err := gopurs_ffi_parser.Bind(content)
	if err != nil {
		panic(fmt.Errorf("Go FFI parse failed: %w", err))
	}
	return result
}

func PackageGoSourceImpl(content string, _ any) string {
	result, err := gopurs_ffi_parser.Packages(content)
	if err != nil {
		panic(fmt.Errorf("Go FFI parse failed: %w", err))
	}
	return result
}

func PackageGoFilesImpl(content string, _ any) string {
	result, err := gopurs_ffi_parser.PackageFiles(content)
	if err != nil {
		panic(fmt.Errorf("Go package emission failed: %w", err))
	}
	return result
}
