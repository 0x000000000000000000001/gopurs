# Writing Go FFI

Implement a PureScript `foreign import` with an ordinary Go function in a `.go`
file beside its `.purs` module. Capitalize the foreign name: `hello` becomes
`Hello`. Imports, packages and local types are allowed; gopurs adapts a copy,
so the original stays usable with Go tooling. No registration is needed.

New project? Start with the [quickstart](quickstart.md) or
[migration example](migration.md#port-the-applications-foreign-functions).
For an existing project, choose a signature below. Read the ownership section
before acquiring resources or retaining callbacks.

These contracts require `0.1.1` and its matching typed frontend. RC14 predates
them. The frontend must preserve `EffectFnAff` identity; its version number
alone does not prove that capability. See the [quality record](quality-0.1.0.md).

## Choose the PureScript contract first

| PureScript import | Ordinary Go implementation |
| --- | --- |
| `String -> String` | `func Echo(value string) string` |
| `Effect String` | `func Label() string` |
| `Effect Unit` | `func Fire()` |
| `String -> EffectFnAff String` | `func Hello(name string) (string, error)` |
| `EffectFnAff String` | `func NoError() string` |
| `EffectFnAff Unit` | `func Notify() error` or `func Fire()` |
| `EffectFn1 String Unit -> Effect Unit` | `func RunCallback(callback func(string))` |
| `EffectFnAff String`, with cancellation | `func Pending(ctx context.Context) (string, error)` |

Pure functions run when called; `Effect` functions run when the effect executes.
Do not add a `Unit` parameter or another function to delay new effect bindings.
Legacy effect thunks remain supported. `EffectFnN` uses uncurried arguments.

For the ordinary error-returning operations below, use `EffectFnAff`. A Go
`error` return does not automatically give a pure import or `Effect` an error policy.

## Deferred operations and errors

Expose `EffectFnAff` through `fromEffectFnAff`:

```purescript
import Effect.Aff (Aff)
import Effect.Aff.Compat (EffectFnAff, fromEffectFnAff)

foreign import hello :: String -> EffectFnAff String

greet :: String -> Aff String
greet name = fromEffectFnAff (hello name)
```

```go
package fixture

import "errors"

func Hello(name string) (string, error) {
    if name == "error" {
        return "", errors.New("expected error")
    }
    return "hello " + name, nil
}
```

Construction does not run the Go function. Each execution starts fresh work in
a goroutine and immediately supplies a canceler to Aff.

- `nil` error means success; a non-nil error retains its Go identity.
- A typed nil inside an `error` interface remains a failure. A successful nil
  handle remains a nil handle.
- Worker panics remain Go panics. Return `error` for expected failures.

Use `EffectFnAff`, not a direct `Aff a` foreign import: the latter is not
qualified by this bridge contract and has different exception behavior.

## Context and cancellation

```purescript
foreign import pending :: EffectFnAff String
```

```go
func Pending(ctx context.Context) (string, error) {
    <-ctx.Done()
    return "", ctx.Err()
}
```

Import Go's `context` package. The bridge supplies one leading `context.Context`;
it is not a PureScript argument. Cancellation sets its cause, and completion
releases it.

Success, error and cancellation have one winner. A winning result is published
before cancellation is acknowledged. If cancellation wins, late results do not
call the result callbacks. Cancellation before bridge installation prevents startup.

Context cannot stop an uncooperative SDK, undo effects or kill a goroutine.
A subscription needs a separate lifetime and explicit stop operation: this
context ends when registration completes. The subscription must synchronize
callbacks with its stop operation.

## Handles and resource ownership

Use `foreign import data` for opaque resources and a Go pointer or SDK handle
in the implementation. Identity and nil pointers are preserved; the SDK defines
sharing, repeated close and use-after-close behavior.

For an acquisition that may finish after cancellation, return a cleanup function:

```go
import "os"

func Open(path string) (*os.File, func(), error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, nil, err
    }
    return file, func() { _ = file.Close() }, nil
}

func Close(file *os.File) error { return file.Close() }
```

```purescript
foreign import data Handle :: Type
foreign import open :: String -> EffectFnAff Handle
foreign import close :: Handle -> EffectFnAff Unit
```

| Outcome | Ownership and cleanup |
| --- | --- |
| Success delivered | The caller owns the resource; use PureScript `bracket` and explicit `close`. The adapter does not call cleanup. |
| Error | A non-nil cleanup runs once before reporting the error. |
| Cancellation wins | A non-nil cleanup runs once when the operation returns, including a late success. Cancellation need not wait for that return. |

For `EffectFnAff Unit`, `(func(), error)` has the same contract without a value.
Nil cleanup is allowed. No destructor is inferred from a `Close` method.
Cleanup returns no error; handle failures inside it. Panics remain panics.

A plain `(T, error)` gives no automatic late-result cleanup. Use the explicit
cleanup result or a masked acquisition/release bracket. The port's bracket
postpones cancellation during acquisition/release and releases a resource
acquired after cancellation before returning the cancellation error.

## Typed callbacks

```purescript
import Effect (Effect)
import Effect.Uncurried (EffectFn1)

foreign import runCallback :: EffectFn1 String Unit -> Effect Unit
```

```go
func RunCallback(callback func(string)) { callback("typed") }
```

Go calls the callback directly. `EffectFnN` uses uncurried arguments; a
`String -> Effect Unit` callback also has its returned effect executed.
Legacy callbacks returning an effect thunk keep that convention.

A bare `any`/`interface{}` result from `EffectFnAff` is ambiguous between a legacy
boxed action and dynamic data. Use a concrete result, or explicit `(any, error)`.
Legacy `Effect` retains its previous interpretation.

Retained callbacks need an explicit stop/unsubscribe lifetime. Conversion alone
does not prevent callbacks after a service closes.

## External Go modules

Go dependencies are separate from Spago's. Prepare them **before generation**:
the package planner may inspect imports with `-mod=readonly`, so a later
`go mod tidy` is too late.

A port declares exact versions in `go-dependencies.json`. Pass every relevant
port, including transitive ports, to the preparation helper:

```json
{
  "schema": 1,
  "require": { "github.com/dlclark/regexp2": "v1.11.5" }
}
```

```sh
node /path/to/gopurs/tools/library-go-dependencies.mjs output ../gopurs-strings
gopurs --main Main
(cd output && GOWORK=off GOPROXY=off GOTOOLCHAIN=local go build -mod=readonly ./main)
```

For npm installations, find the helper under `node_modules/gopurs/tools/`.
It creates `module gopurs/output` with Go 1.22 if needed, installs exact versions
and downloads their dependencies. Existing application requirements,
replacements and checksums are preserved. Pin or toolchain conflicts fail and
restore `go.mod`/`go.sum`; resolve them explicitly. The compiler never selects
`latest` or installs dependencies implicitly.

Run preparation again after deleting `output/`; keep `go.mod` and `go.sum`
between builds. Repository fixtures prepare local overrides automatically;
standalone applications and port runners must do so themselves.

## Generic Go functions

Unconstrained Go parameters can forward polymorphic values through callbacks:

```purescript
foreign import transform :: forall a b. (a -> b) -> Array a -> Array b
```

```go
func Transform[A, B any](f func(A) B, values []A) []B {
    result := make([]B, len(values))
    for i, value := range values {
        result[i] = f(value)
    }
    return result
}
```

This implementation calls the callback once per element, in order, and allocates
a separate result slice. Use concrete Go types when the PureScript contract
provides them.

The general bridge instantiates these parameters with `gopurs_runtime.Value`,
including inside callbacks and slices; handwritten code need not import the
runtime. This does not promise scalar specialization or zero allocation.
A PureScript type variable does not establish Go arithmetic, ordering or SDK
constraints: use a concrete function or pass the operation as a typed callback.

Examples are `OrdCharImpl`/`OrdStringImpl` in `gopurs-prelude` and `AnyImpl` in
`gopurs-arrays`; see [generic FFI tests](../tools/ffi-generics.test.mjs).

## Foreign values and records

`Foreign` crosses as `any`, and `Array Foreign` as `[]any`. PureScript data is
exposed as native primitives/containers; already-native opaque payloads retain
identity and numeric types. Nil slices remain nil.

Legacy maps/slices containing boxed values are converted only where needed.
Native subgraphs retain identity; cycles and repeated converted references are
preserved. Opaque structs and unrelated SDK types are not traversed.

Missing keys, present nil, empty arrays, false and zero stay distinct.
Undefined uses an opaque marker and round-trips to PureScript undefined.
Transport preserves its keys; omission happens only during application JSON
serialization. Serialize or convert before passing that marker to an arbitrary
SDK or Go's JSON encoder. SQL formatting, numeric conversions, PureScript
functions and algebraic constructors also need appropriate application conversion.

Closed records map to locally declared Go structs:

```purescript
type Child = { active :: Boolean, label :: String }
foreign import children :: Array Child -> Array Child
```

```go
type Child struct {
    Active bool
    Label  string
}

func Children(values []Child) []Child { return values }
```

Match every field with its capitalized label and compatible concrete type.
Fields must be exported and non-embedded; missing, extra or colliding fields
are errors. An `any` field requires PureScript `Foreign`. Nested closed records
and arrays work; `Nullable` records use `*Struct`, ordinary records use values.
Map fields have no qualified structural conversion. Converted records have
value semantics, not opaque pointer identity.

| Boundary | Use |
| --- | --- |
| Polymorphic value forwarded unchanged | Unconstrained type parameters and typed callbacks |
| Dynamic SDK data | `Foreign` → `any`/`[]any`; validate its shape |
| Open-row helper | An explicit map helper; keep Go parameters together ([example](migration.md#curried-record-helper-signatures)) |
| ADT/newtype | Existing qualified representation or explicit PureScript construction/deconstruction |
| Imported SDK struct | Opaque handle, or a checked local struct with explicit SDK conversion |

Open rows, polymorphic variants and arbitrary imported structs have no inferred
struct layout. TAST types alone do not define a public Go sum-type ABI.

`any` and `interface{}` are identical spellings. A legacy polymorphic `any` may
carry a backend value; `Foreign` explicitly requests native conversion.
Ordinary application FFI should not inspect backend tags. Document specialized
runtime boundaries used by library primitives.

## Copying, aliases and mutation

Treat arrays, records and polymorphic inputs as borrowed: pure FFI must preserve
them. A native signature or conversion does not prove exclusive ownership.
Allocate a result before updating input; a shallow slice/map copy does not own
nested handles or containers.

Structural record conversion reconstructs fields/arrays, but opaque handles
retain identity and native `Foreign` subgraphs may stay shared. Neither is a
deep-clone guarantee. Nullable struct records still have value semantics.

For intentional resource mutation, define creation, observable aliases,
concurrency and closure. Use explicit stop operations for retained callbacks,
acquisition cleanup on failure/cancellation and caller release on success.
See [structured-value tests](../tools/ffi-structured.test.mjs) and
[Foreign transport tests](../runtime/ffi_foreign_test.go).

## Diagnostics and verification

Unsupported signatures report the module, Go path/location, both signatures
and a corrective hint. Results are never silently truncated. Fixed arrays,
variadics, multiple-result callbacks and aliases needing unavailable export
data require an explicitly supported boundary.

### Tests in a development checkout

```sh
GO111MODULE=off go test -race ./runtime
node --test tools/ffi-aff.test.mjs tools/ffi-values.test.mjs tools/ffi-structured.test.mjs
node --test tools/ffi-docs.test.mjs
GOPURS_TEST_MODES=js node --test tools/ffi-aff-e2e.test.mjs
GOPURS_TEST_MODES=native-parallel node --test tools/ffi-aff-e2e.test.mjs
```

The Aff fixture compares official JavaScript with generated Go in Release/Dev
and reused caches. JS/WASM and native hosts share the Go parser; this does not
make network SDKs executable on Go/WASM. Rebuild each compiler host first.
Additional Aff ownership/cancellation checks live in `gopurs-aff`:
`node --test tools/aff-native.test.mjs`.

### Select the frontend

Tests honor `GOPURS_PURS`; application `spago build` resolves `purs` through
`PATH`. Rebuilding gopurs does not install the frontend. Different fork builds
may both report `0.15.16`: verify the executable/revision and rebuild frontend
output after changing it. Old externs can retain expanded `EffectFn2` instead
of `EffectFnAff`. In b8x, `b -c` clears that output after frontend installation.

### Test the packaged compiler

```sh
GOPURS_TEST_MODES=js node --test tools/ffi-package.test.mjs
GOPURS_TEST_MODES=js,native-parallel node --test tools/ffi-package.test.mjs
```

These checks install a real tarball with an empty npm cache and test migration,
cold/reused caches, changed signatures, errors and recovery. `GOPURS_OFFICIAL_PURS`
selects stock PureScript; `GOPURS_PURS` selects the typed fork.
`GOPURS_TEST_KEEP_WORKSPACE=1` retains reports and logs. The native binary is
qualified separately from the package's parser/helper sources.

Functional/race checks do not establish performance gains; compare those with
the official altbak.pub baselines.
