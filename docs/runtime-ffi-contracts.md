# Runtime and compiler FFI contracts

## Boundary owners

| Boundary | Owner |
| --- | --- |
| PureScript types to Go representations | `GoTypes`, `GoConversions` and `FfiBridge`, using TAST. |
| Compiler operations on JS, Go or Rust hosts | Foreign companions under `src/`, with explicit owners for mutable state. |
| Boxed values and storage | `runtime/runtime.go`: `Value`, constructors, accessors, applications and conversions. |
| Runtime source embedded in the compiler | `tools/build/embed-runtime.mjs` and `Gopurs.Runtime.Source`. |
| Compiler to generated application | `Driver.Output.writeRuntime`, writing the embedded bytes to `output/gopurs_runtime/runtime.go`. |

The runtime is both an API used by library FFI and a source artifact copied
into applications. Tests exercise the canonical source directly; a built
compiler does not read that source file at startup.

## Value storage and lifetime

`Value` contains a `Type` tag, an `IntVal` field and a GC-tracked `UnsafePtr`.
Copying it copies the wrapper, not the pointed-to storage. Its producer must
supply the tag and layout expected by the chosen accessor. Direct FFI access
is not a decoder for arbitrary values.

| Family | Payload and sharing |
| --- | --- |
| Integers, numbers, booleans | `IntVal`; floating-point values use `math.Float64bits`, booleans use 0/1. |
| Strings | Byte length in `IntVal`, immutable byte pointer in `UnsafePtr`. `Str`/`StrValue` own the boundary; substrings retain backing storage. Empty strings have no pointer. |
| Arrays | `Array` allocates a slice wrapper and shares `[]Value` backing storage. Conversions to another representation convert elements. |
| Records | `Record` shares its map; `RecordDict` shares key/value slices. Fixed `RecordDict0` through `RecordDict5` own compact containers. Published PureScript containers are immutable. |
| Opaque values | `Any` retains a Go interface and its referenced storage. Copies preserve identity; the owning FFI determines mutability. |
| Functions | `Func` through `Func11` retain closures and captures on the heap through `forceEscape`. Partial applications retain supplied arguments. `WithFunctionData` attaches metadata; partial application returns an ordinary function. |
| Native ADTs | Code generation selects tags and layouts. `StructGetters` registers getters during Go initialization. `Rc` participates in reuse protocols; the GC still owns memory reclamation. |

Container constructors and `Box`/`Unbox` do not promise deep copies.
`Box([]Value)` shares backing storage; `Box([]any)` creates a new boxed slice.
FFI must distinguish read-only views from privately owned writable buffers.
See [packed strings](packed-strings.md), [closure lifetime](runtime-closures.md),
[ADT reuse](adt-reuse.md) and [representation conversions](architecture.md#representations-and-conversions).
`runtime.KeepAlive` alone does not replace the heap escape enforced by
`forceEscape` for closure captures.

## Primitive semantics

### UTF-16 concatenation

PureScript strings use canonical WTF-8 storage: valid surrogate pairs form
UTF-8 scalars, while isolated surrogates retain their three WTF-8 bytes.
`ConcatString` combines a trailing high surrogate with a leading low surrogate
across its operands, without mutating either operand.

`PrimitiveExprs` emits this helper for `OpStringAppend`; the Go
`Data.Semigroup` FFI also uses it. The native compiler therefore uses the same
semantics when folding constants. Plain Go string addition does not cover
this boundary case.

### Int operations

Runtime helpers follow the reference JS FFI of Prelude and `Data.Int.Bits`:

- `IntNegate`, `IntAdd` and `IntSub` wrap to signed 32-bit results.
- `IntMul` reproduces `(a * b) | 0`: round the product as a `Number`, then
  convert to 32 bits. For example, `top * top` is `0`. Modulo `2^32` precedes
  Go conversion for products beyond `int64` range.
- Bitwise operands are 32-bit. Shifts mask the count by `31`, including
  negative counts; `IntZshr` returns the unsigned result.
- `IntDiv` and `IntMod` implement Prelude's Euclidean behavior and return zero
  for a zero divisor. The reference result of `bottom / -1` is `2147483648`.

Generated primitives and the Go library FFI share these helpers. `Value`'s
`int64` storage can carry unsigned-shift results and the division boundary;
normalization occurs at the relevant operation. Operands are evaluated once
and in order. Admitted [Int32 loops](int32-loops.md) can use signed native
slots after proving their writes and checking entry ranges; worker boundaries
and the general path retain `int64` storage.

### Number formatting

The Go `Data.Show.ShowNumberImpl` follows Prelude's JS `showNumberImpl`:
shortest round-trippable binary64 digits, decimal notation for
`1e-6 <= abs(n) < 1e21`, scientific notation otherwise. Exponents have `e`,
an explicit sign and no padding; decimal integers receive `.0`.
Both zeros print `0.0`; special values print `NaN`, `Infinity` and `-Infinity`.
Stored negative zero remains distinct, and generated literals preserve it
through `runtime.NegativeZero`. The executed JS FFI is the test oracle.

## JSON views and updates

- `ReadJSONObject` borrows foreign maps and compact `JSONObject` storage.
  `ReadJSONObjectValue` materializes maps for record representations.
  A view does not authorize mutation of its source.
- A compact `JSONObject` owns its storage and strings. `UnboxObject` and
  `RecordToMap` materialize independent containers for it; nested values can
  still share storage.
- For opaque values already holding a map, `UnboxObject` directly returns a
  `map[string]any` and `RecordToMap` directly returns a `map[string]Value`.
  Writers must own that map or copy it explicitly.
- `RecordSet` implements immutable insertion/replacement, including foreign
  maps. `RecordUpdateDict` replaces existing fields in PureScript records;
  compact forms share keys and copy values. Its map path inherits
  `RecordToMap` sharing and is not a general copy operation for opaque maps.

## Asynchronous work

`Retain`, `Release` and `EventLoopWait` track outstanding host tasks. Retain
before starting a task and release exactly once on completion. The counter
must be positive when waiting for that generation begins. This protocol
controls program completion; Go references control value lifetime.

## Compiler companion contracts

| Module | Contract |
| --- | --- |
| `Printer` | Identical literal escaping across hosts: control escapes, valid UTF-8 and isolated UTF-16 surrogates represented through WTF-8. The Go scanner works in bytes. |
| `Printer.Builder` | One opaque buffer per render. `pushImpl` mutates and returns the same handle; older handles alias it. `toStringImpl` returns an immutable string. |
| `GoAst` | `memoizeName` owns a cache per partial application, including empty results. The callback is pure; a Go mutex covers lookup and initialization. |
| `GoCode` | JS uses the PureScript fallback; native Go scanning follows the same token boundaries, ignores literals/comments and returns unique sorted imports. |
| `GoImports` | Inputs are borrowed. The returned array is new, preserving order and duplicates while sharing immutable strings; `collectImports` sorts and deduplicates. |
| `FfiSupport` | The same Go parser and JSON protocol over WASM, native Go or a Go archive through C for Rust. Operations `0`/`1` extract/prepare FFI; operation `2` binds typed Go references. Extraction/preparation errors include module and path. |
| `Runtime.Source` | Constant embedded from canonical Go source, with no runtime source-file read. |
| `Host.Metrics` | Monotonic milliseconds with a host-local origin. Memory-profile configuration is effective in Go and neutral in JS, set before jobs begin. |
| `Main` | One diagnostic after cleanup; JS allows output to drain, Go exits after synchronous output. |

`Printer.Builder.withOut` confines mutation to a synchronous callback. Chain
handles in order and do not retain one for another render. The opaque type
does not guarantee linear ownership. Go stores a pointer to `strings.Builder`,
never a copied builder; concurrent renders use different buffers.

`Cache.Pool` exposes its native pointer; `BuildCache` and `Host.Snapshot`
callbacks are native functions. The bridge still supplies required PureScript
boxing. A final `_ any` parameter keeps effectful operations delayed; a
`String -> Effect Unit` callback becomes `func(string) func()`.
`GoImports.concatStringArrays` is generic to preserve element representation.

For pointer-returning workers, the bridge reuses an argument's wrapper when
the result is that same pointer with matching Go and PureScript types; a new
pointer receives a new wrapper. `Cache.TypeMemo` uses comparable generic keys
instantiated with the full `Value` representation of `ExprType`, preserving its
tag and pointer owner. The companion does not depend on the runtime package.

The parser, typed adaptations and diagnostics remain owned by `ffi-gen`,
`FfiBridge` and `FfiSupport`/`Driver`, respectively. Different host transports
share parser and representation policy.

## Verification

After building the compiler, select the relevant suites:

```sh
node --test tools/native-ffi.test.mjs tools/embed-runtime.test.mjs \
  tools/runtime-contracts.test.mjs tools/string-concat.test.mjs \
  tools/closure-lifetime.test.mjs tools/integer-boundaries.test.mjs \
  tools/integer-division.test.mjs tools/number-show.test.mjs \
  tools/apply-arity.test.mjs tools/function-data.test.mjs \
  tools/value-array-unboxing.test.mjs tools/go-imports.test.mjs
```

`native-ffi` compares native companions against JS oracles and exercises
buffers, caches, GC and clocks under the race detector. `runtime-contracts`
runs canonical runtime tests normally and with `-race`, including aliases,
immutable updates and task retention. Primitive suites compare executed Go
against JS for surrogate boundaries, integer operations and binary64 formats.
`embed-runtime` checks embedded bytes, stable timestamps, regeneration and
loading without the source file. The closure and application suites exercise
arities, captures and partial application.

`native-go-code.test.mjs` additionally checks the compiled scanner with a
bounded stack using `GOPURS_NATIVE_OUTPUT` from a retained bootstrap.
See [testing.md](testing.md) for setup and command selection.
