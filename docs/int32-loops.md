# Native signed-32-bit loops

[`Gopurs.Int32Loops`](../src/Gopurs/Int32Loops.purs) keeps eligible loop slots
and iteration operations in Go `int32`, avoiding repeated sign extension inside
the loop. Public function, wrapper and FFI signatures remain `int64`.

## Transformation

The fast path checks the initial values before narrowing them. For example:

```go
if n == int64(int32(n)) && acc == int64(int32(acc)) {
    return func() int64 {
        n32, acc32 := int32(n), int32(acc)
        for n32 != 0 {
            n32--
            acc32++
        }
        return int64(acc32)
    }()
}
// The original loop handles inputs outside the signed 32-bit range.
```

Signed overflow follows PureScript Int semantics. Immediate returns of wide
input values retain the original path, including when the loop runs zero times.
The accumulator is widened on exit rather than after every operation.

The proof uses the structured Go AST before printing. Each slot is selected
independently only if all recognized writes normalize its value. Comparisons
are narrowed only when both operands are proven signed. Captures, shadowing,
nested loops, opaque references and unknown forms reject the affected admission.
See [bindings and tail calls](architecture.md#bindings-captures-and-tail-calls).

## Validation

After building:

```sh
node --test tools/int32-loops.test.mjs
bin/test Int32Loops
```

The tool tests compile and execute generated before/after Go under the race
detector against a JavaScript oracle. Cases cover overflow, wide entry and exit
values, zero iterations, branches, simultaneous updates, effectful operand
order, and rejected AST forms. Performance measurements follow the
[benchmark protocol](benchmark-protocol.md); old scratch binaries and timings
are not the current implementation contract.
