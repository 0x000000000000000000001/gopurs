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
