# Integer-array round-trip elimination

A freshly generated filter buffer can avoid the two copies in an immediate
`Value → []int64 → Value` round trip consumed by an integer fold. The generator
normalizes the private buffer in place before the fold. Range generation,
filtering, fold iteration and callback dispatch retain their normal behavior.

## Recognition

[GoConversions](../src/Gopurs/GoConversions.purs) preserves integer-array
conversions as `GoUnboxIntArray` and `GoBoxIntArray` until printing.
[ArrayIntrinsics](../src/Gopurs/ArrayIntrinsics.purs) marks its own two-argument
`Data.Array.filterImpl` result with `GoFreshFilterArray` only when it creates a
fresh `[]Value` through `make` and `append`. The marker itself prints transparently.

The ordinary `App` path recognizes this exact expression after argument boxing:

```text
foldInt(step, seed,
  GoBoxIntArray(GoUnboxIntArray(Array(GoFreshFilterArray(filterExpr)))))
```

The callee must be the qualified `Data.Foldable.foldlArray` with exactly three
arguments, the callback must have the concrete `(Int, Int) -> Int` signature,
and the seed must be an integer literal. Dynamic seeds, unknown callbacks,
other scalar types and unqualified or different callees keep the copies.

The proof comes from the structured conversions and the fresh-buffer marker.
It does not search through variables, remaining lets, branches, calls, storage
or raw Go, and does not propagate freshness through a binding table. Outer
`Typed` wrappers may describe enclosing applications; those annotations alone
do not establish the runtime representation.

## Transformation and guarantees

Only the array expression is replaced. Its simplified Go shape is:

```go
func() Value {
    source := Array(filterExpr())
    items := *(*[]Value)(source.UnsafePtr)
    for i, item := range items {
        items[i] = Int(item.IntVal)
    }
    return source
}()
```

The producer runs once, before fresh IIFE-local temporaries enter scope.
Translation statements, evaluation order, the seed, `Apply2` calls and element
order are preserved. Normalization uses exactly `.IntVal` followed by `Int`:
no additional tag check or signed-32-bit conversion is introduced. This resets
each element's tag and pointer exactly as the original copies did.

Normalization finishes before the first callback, clearing retained pointers
at the same stage as the original round trip. The predicate cannot access the
fresh result buffer; this local fact permits mutation without a general
ownership analysis. Upstream FFI calls, including range generation, still run.

## Fallbacks and validation

Stored or opaque arrays, references, ST/FFI-produced arrays without the marker,
other consumers, partial or excess applications, and the `UncurriedApp` fold
path retain their normal conversions. Empty filter results use the same rule
with zero normalization iterations and an unchanged accumulator. A valid empty
array header is distinct from a `Value` with a nil `UnsafePtr`.

After building:

```sh
node --test tools/array-intrinsics.test.mjs
bin/test ArrayRoundtrip
```

The tests verify the exact marker and conversion shape, strict admission and
fallbacks. The [generated fixture](../tests/passing-snapshots/ArrayRoundtrip.go)
shows the normalization in an actual compiled program. Use the
[benchmark protocol](benchmark-protocol.md) for performance comparisons.
