# Writing Go FFI

Write a `.go` file next to the PureScript module, with ordinary Go functions.
The exported Go name is the PureScript foreign name with its first letter
capitalized. Imports, package declarations and local type declarations are
allowed. The compiler adapts a copy; the handwritten file remains usable by
`gofmt`, `go test` and Go tooling. No per-function registration is needed.

These conventions describe experimental `0.1.0-rc.18` and its matching typed
frontend. Consult the live [quality record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md)
for qualification and publication status. Published `0.1.0-rc.14` predates
this work. The frontend
must retain the identity of `Effect.Aff.Compat.EffectFnAff`; a matching version
number alone does not establish that capability.

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

Pure functions run when called. An `Effect` implementation runs when its effect
is executed: do not add a `Unit` parameter or return an effect thunk. Existing
explicit-thunk FFI remains supported during migration. The existing `EffectFnN`
calling convention remains available for uncurried functions.

A Go `error` return does not itself choose an effect. Use `EffectFnAff` for the
ordinary error-returning operations shown here. A pure import or `Effect` does
not silently acquire an error-handling policy.

## Deferred operations and errors

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

Constructing `hello name`, or converting it with `fromEffectFnAff`, does not run
the Go function. Every execution starts a fresh operation. Installation returns
the canceler immediately; blocking Go work runs in a worker goroutine.

`nil` error means success. A non-nil error reaches the PureScript error callback
with its Go identity preserved. A typed nil pointer stored in an `error`
interface is non-nil in Go and remains a failure. A successful typed nil handle
remains a nil handle; it is not mistaken for an error.

Panics in these workers remain Go panics. They are not recovered and converted
into business errors. Return `error` for expected failures. Direct ordinary Go
implementations of `foreign import ... :: Aff a` are not qualified by this
contract: use `EffectFnAff` and `fromEffectFnAff`. The existing Aff port also has
its own synchronous-exception behavior, so direct Aff import is not an
interchangeable spelling for this bridge.

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

Import Go's standard `context` package. One leading `context.Context` parameter
is supplied by the adapter; it is not a PureScript argument. Canceling the Aff
cancels this context, with the cancellation error as its cause. The context is
also released when the operation finishes.

Success, error and cancellation have a single winner. Once a result wins,
cancellation waits for its callback publication before acknowledging; the Aff
port does not discard that published value. If cancellation wins, late results
do not invoke the result callbacks. An action canceled before installation does
not start the operation.

An SDK that ignores context can continue after its caller stops waiting.
Cancellation does not forcibly stop a goroutine or undo an SDK side effect.
Do not use this operation context as the lifetime of a subscription that must
continue after registration succeeds: its lifetime ends with registration.
Such a subscription needs its own ordinary Go lifetime and an explicit stop
operation. Synchronizing callbacks with that stop is the subscription's
responsibility.

## Handles and resource ownership

Opaque PureScript foreign data can use native Go pointers or SDK handle types.
The adapter transports them without requiring handwritten unboxing. Handle
identity and nil pointers are retained. The SDK still defines whether a handle
may be shared, closed twice or used after closure.

For an acquisition that could return a resource after cancellation, return a
cleanup function along with the ordinary result:

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

The cleanup function has exactly this ownership contract:

- On success delivered to PureScript, ownership transfers to the caller. The
  adapter does not call cleanup. Use the existing PureScript `bracket` and
  explicit `close` operation for the normal lifetime.
- On error, a non-nil cleanup runs once before the error is reported.
- If cancellation wins, a non-nil cleanup runs once when the operation returns,
  including for a late successful result. Cancellation need not wait for that
  return.

For `EffectFnAff Unit`, `(func(), error)` supplies the same cleanup contract
without a value. A nil cleanup is permitted. The compiler never infers a
destructor from a method named `Close`. Cleanup has no error return: choose
logging or recovery inside it when a cleanup failure needs attention. Panics
from cleanup remain panics.

A plain `(T, error)` acquisition provides no automatic late-result cleanup.
Use a masked acquisition/release bracket or the explicit cleanup result when
ownership matters. The port's bracket masks acquisition, then releases a
resource acquired after cancellation before returning the cancellation error.

## Typed callbacks

```purescript
import Effect (Effect)
import Effect.Uncurried (EffectFn1)

foreign import runCallback :: EffectFn1 String Unit -> Effect Unit
```

```go
func RunCallback(callback func(string)) { callback("typed") }
```

The Go callback is directly callable. `EffectFnN` callbacks use their uncurried
convention. For a callback typed `String -> Effect Unit`, the generated adapter
also executes its returned effect when Go calls `func(string)`.

Legacy callbacks whose Go result already represents an effect thunk retain
that thunk convention. Prefer concrete results for new effect implementations.
An `EffectFnAff` with a bare `any` or `interface{}` Go result is diagnosed as
ambiguous: it could mean an old boxed action or a new dynamic result. Use a
concrete result type, or an explicit `(any, error)` result when dynamic data and
an error result are needed. Existing legacy `Effect` implementations retain
their previous interpretation during migration.

Callbacks retained beyond the call require a defined stop/unsubscribe lifetime.
Generated argument conversion alone does not guarantee that a callback cannot
run after a service is closed.

## Foreign values and records

`Foreign` can cross as `any`, and `Array Foreign` as `[]any`. Values originating
in PureScript are exposed as Go primitives and containers; a Go author need not
inspect backend tags. Already-native opaque payloads retain their identity and
numeric types. Native nil slices remain nil when transported as Foreign. Native
`map[string]any` and `[]any` containers left by legacy FFI are also checked for
nested boxed values. Only branches needing conversion are copied; wholly native
subgraphs retain their identity. Cycles and repeated references to a converted
child are preserved. Opaque structs and unrelated SDK types are not traversed.

Absence of a map key remains distinct from a present nil value. Empty arrays,
false and zero are not collapsed into nil. Undefined travels as an opaque native
marker, distinct from nil, and becomes PureScript undefined again when returned
or passed to a callback. Transport retains keys with undefined values; the
Foreign/JSON library applies omission only when the application asks it to
serialize. The marker is not a normal Go data value for arbitrary SDKs or Go's
JSON encoder: application serialization must run before those boundaries.
SQL formatting, JSON serialization and conversions such as decimal-to-string
remain application code. PureScript functions and algebraic constructors still
require an explicit conversion to native data.

Closed records map to locally declared Go structs, for example:

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

Each field uses the capitalized PureScript label. Both sides must describe all
fields with compatible types. Fields must be exported and non-embedded;
colliding capitalized labels, missing or extra fields are errors. Use concrete
Go field types; an `any` field requires an explicit PureScript `Foreign` field.
Map fields do not yet have a qualified structural conversion. Nested closed
records and arrays are supported. `Nullable` records use `*Struct`; ordinary
non-nullable records use struct values. A structurally converted record is a
value, not an opaque pointer-identity guarantee.

Open rows, polymorphic variants and arbitrary imported struct layouts do not
have an inferred structural mapping. Newtypes and ADT transport must follow
their existing qualified representations; there is no general new automatic
Go sum-type mapping. Use an explicitly supported contract or retain a documented
specialized FFI until a mapping is established.

## Diagnostics and verification

Unsupported signatures report the PureScript module, handwritten Go path,
declaration location, encountered Go/PureScript signatures and a corrective
hint. Multiple results are preserved, including in callbacks; an unsupported
shape is rejected instead of being truncated. Fixed-size arrays, variadic
parameters and multiple-result callbacks currently require a supported boundary
such as slices and single-result callbacks. Type aliases requiring unavailable
external export data are not guessed from their spelling.

The standalone Go implementation and real PureScript imports are exercised by
`tools/ffi-aff-e2e.test.mjs`. It compares an official JavaScript Aff reference
with generated Go for both release and dev, including reused caches. The
JavaScript compiler host uses the WASM Go parser; the native compiler host uses
the same parser core. Neither claim makes network SDKs executable on Go/WASM.

From a development checkout with its local library ports and typed frontend:

```sh
GO111MODULE=off go test -race ./runtime
node --test tools/ffi-aff.test.mjs tools/ffi-values.test.mjs tools/ffi-structured.test.mjs
node --test tools/ffi-docs.test.mjs
GOPURS_TEST_MODES=js node --test tools/ffi-aff-e2e.test.mjs
GOPURS_TEST_MODES=native-parallel node --test tools/ffi-aff-e2e.test.mjs
```

These test runners honor `GOPURS_PURS` when several typed frontends are installed.
Ordinary application `spago build` instead resolves `purs` through `PATH`; make
sure that executable is the rebuilt fork too. Rebuilding gopurs does not install
the PureScript frontend. Older and newer fork builds can both report `0.15.16`,
so check the executable path and development revision, then rebuild the
application's frontend output after changing the TAST producer. Reusing old
externs can retain the expanded `EffectFn2` form of `EffectFnAff` even with the
new compiler. In b8x, `b -c` clears that output after the frontend is installed.

Build the compiler hosts before testing them. Native Aff ownership
and legacy cancellation tests live in the separate `gopurs-aff` checkout:
`node --test tools/aff-native.test.mjs`.

The packaged-consumer check installs an actual tarball with `--ignore-scripts`
and an empty npm cache. It exercises a stock-compiler JavaScript project before
migration, the matching typed frontend, cold and reused caches, a Go signature
change, rejection of an incompatible signature and recovery after that error:

```sh
GOPURS_TEST_MODES=js node --test tools/ffi-package.test.mjs
GOPURS_TEST_MODES=js,native-parallel node --test tools/ffi-package.test.mjs
```

`GOPURS_OFFICIAL_PURS` selects the stock compiler for the migration reference;
`GOPURS_PURS` selects the typed compiler. Set `GOPURS_TEST_KEEP_WORKSPACE=1` to
retain the tarball, reports and logs. The package contains the parser WASM and
native helper sources; the optional native compiler binary is qualified
separately by this test.

These are functional and race checks. They make no application performance
claim; performance comparisons must also use the official altbak.pub baselines.
