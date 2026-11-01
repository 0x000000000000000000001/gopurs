// Handwritten FFI for Gopurs.GoCode; keep in sync with the .purs foreign imports.
// JavaScript calls the PureScript scanner supplied by the caller; the Go
// backend implements the same scan natively.
export const referencedImportsImpl = fallback => text => fallback(text);
