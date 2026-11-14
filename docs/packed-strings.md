# Packed String payloads

[`runtime/runtime.go`](../runtime/runtime.go) stores `TypeString` directly in a
`Value`, without allocating a separate `*string` header:

| Field | Meaning |
| --- | --- |
| `IntVal` | Byte length |
| `UnsafePtr` | Immutable string bytes from `unsafe.StringData` |
| Nil pointer | Empty string |

Write through `Str` and read through `StrValue`:

```go
func StrValue(v Value) string {
    if v.UnsafePtr == nil {
        return ""
    }
    return unsafe.String((*byte)(v.UnsafePtr), int(v.IntVal))
}
```

## Contract

- Referenced bytes must remain immutable for the entire lifetime of the value,
  as required by `unsafe.String`. FFI code must copy mutable buffers before
  exposing them as strings.
- A substring retains its backing allocation.
- An empty string has no pointer. Always use `StrValue`; do not dereference
  `UnsafePtr` or reinterpret it as `*string`.
- The `Value` size is unchanged. Consumers must use the shared helpers rather
  than depending on the old representation.

Runtime changes must also regenerate `src/Gopurs/Runtime/Source.go` and
`Source.js` using `npm run build:runtime`. The generated sources are checked by
[embed-runtime.test.mjs](../tools/embed-runtime.test.mjs).

## Validation

[ffi-generics.test.mjs](../tools/ffi-generics.test.mjs) checks the generic bridge,
including its allocation contract. Library regression runs should include
`argonaut-core`, `foreign`, `foreign-object`, `node-buffer` and `strings-extra`.
Performance comparisons use the [benchmark protocol](benchmark-protocol.md)
and the official altbak baselines.
