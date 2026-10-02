package Gopurs_Printer_Builder

import (
	"strings"

	"gopurs/output/gopurs_runtime"
)

// Each rendering owns one opaque pointer. Copying the Value aliases that
// handle; the strings.Builder itself must never be copied or shared by writers.
func NewBuilderImpl(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Any(new(strings.Builder))
}

func PushImpl(builder gopurs_runtime.Value, chunk string) gopurs_runtime.Value {
	builder.AnyVal().(*strings.Builder).WriteString(chunk)
	return builder
}

// String keeps its immutable backing alive after the handle is released.
func ToStringImpl(builder gopurs_runtime.Value) string {
	return builder.AnyVal().(*strings.Builder).String()
}
