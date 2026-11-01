// Handwritten FFI for Gopurs.GoImports; keep in sync with the .purs foreign imports.
// Native array flattening for the Go backend. The JavaScript backend uses a
// single concat so both produce the same multiset of imports.
export const concatStringArrays = arrays => [].concat(...arrays);
