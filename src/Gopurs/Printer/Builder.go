// Handwritten FFI for Gopurs.Printer.Builder; keep in sync with the .purs foreign imports.
package Gopurs_Printer_Builder

import "strings"

// Each rendering owns one opaque pointer. Copies of the handle alias it;
// the strings.Builder itself must never be copied or shared by writers.
func NewBuilderImpl() *strings.Builder {
	return new(strings.Builder)
}

func PushImpl(builder *strings.Builder, chunk string) *strings.Builder {
	builder.WriteString(chunk)
	return builder
}

// String keeps its immutable backing alive after the handle is released.
func ToStringImpl(builder *strings.Builder) string {
	return builder.String()
}
