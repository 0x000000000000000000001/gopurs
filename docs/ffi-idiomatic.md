# Writing Go FFI

Use this guide to implement a PureScript `foreign import` in Go. FFI means
foreign function interface: the boundary where PureScript calls code written
in another language.

Write a `.go` file next to the PureScript module, with ordinary Go functions.
The exported Go name is the PureScript foreign name with its first letter
capitalized. Imports, package declarations and local type declarations are
allowed. The compiler adapts a copy; the handwritten file remains usable by
`gofmt`, `go test` and Go tooling. No per-function registration is needed.

For the complete path from an empty workspace to execution, follow the
[quickstart](quickstart.md), then the [migration guide's first foreign
import](migration.md#port-the-applications-foreign-functions). Once your project
is configured, choose a signature from the table below and implement the
matching Go function.

Read [deferred operations and errors](#deferred-operations-and-errors) for SDK
calls that can fail, then [resource ownership](#handles-and-resource-ownership)
if they acquire files, connections or other resources. The later sections cover
[external dependencies](#external-go-modules), [generic functions](#generic-go-functions)
and [data conversion](#foreign-values-and-records). Use
[verification](#diagnostics-and-verification) when changing or testing a port.

These conventions describe `0.1.1` and its matching typed
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
is executed: do not add a `Unit` parameter or return an effect thunk, meaning
another function that must be called to run the effect. Existing FFI that
explicitly returns such a function remains supported during migration.
`EffectFnN` remains available for uncurried functions, whose arguments are
passed together in one call.

A Go `error` return does not itself choose an effect. Use `EffectFnAff` for the
ordinary error-returning operations shown here. A pure import or `Effect` does
not silently acquire an error-handling policy.

## Deferred operations and errors

`EffectFnAff` is the foreign representation of an asynchronous operation.
Wrap it with `fromEffectFnAff` to expose an ordinary `Aff` action to PureScript
callers:

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
the Go function. Every execution starts a fresh operation. When the Aff action
starts, the bridge immediately returns a canceler to the Aff runtime; blocking
Go work runs in a worker goroutine.

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
cancellation waits for the result to be published through its callback before
acknowledging cancellation; the Aff port does not discard that published value.
If cancellation wins, late results do not invoke the result callbacks. An action
canceled before the bridge is installed does not start the operation.

An SDK that ignores context can continue after its caller stops waiting.
Cancellation does not forcibly stop a goroutine or undo an SDK side effect.
Do not use this operation context as the lifetime of a subscription that must
continue after registration succeeds: its lifetime ends with registration.
Such a subscription needs its own ordinary Go lifetime and an explicit stop
operation. Synchronizing callbacks with that stop is the subscription's
responsibility.

## Handles and resource ownership

An opaque handle gives PureScript a reference to a resource without exposing
its Go fields. Declare it with `foreign import data`, and use a native Go
pointer or SDK handle type in the implementation.
The adapter transports these handles without requiring handwritten unboxing.
Handle identity and nil pointers are retained. The SDK still defines whether
a handle may be shared, closed twice or used after closure.

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
Here, masking means postponing cancellation while acquisition or release runs,
so cancellation cannot leave a resource between those two responsibilities.

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

## External Go modules

Go dependencies are separate from Spago dependencies. Declare them before
generation: the native compiler's package planner can inspect imported Go
packages with `-mod=readonly`. Running `go mod tidy` only after generation
is too late for that path.

A library port can publish `go-dependencies.json` at its root:

```json
{
  "schema": 1,
  "require": { "github.com/dlclark/regexp2": "v1.11.5" }
}
```

Prepare an application's output module explicitly, passing the relevant
ports, including transitive ones with Go dependencies:

```sh
node /path/to/gopurs/tools/library-go-dependencies.mjs output ../gopurs-strings
gopurs --main Main
(cd output && GOWORK=off GOPROXY=off GOTOOLCHAIN=local go build -mod=readonly ./main)
```

For an npm installation, the helper is at
`node_modules/gopurs/tools/library-go-dependencies.mjs`; in a source checkout,
use its `tools/` directory.

The preparation command creates `module gopurs/output` with Go 1.22 when
needed, installs the exact declared versions and downloads their transitive
graph. It preserves existing application requirements, replacements and
checksums. Conflicting port pins, application pins or required toolchain
changes stop preparation and restore `go.mod`/`go.sum`; reconcile them
explicitly. Replacements already chosen by the application remain in effect.
Keep both files when rebuilding frontend output. The compiler does not
implicitly install modules or select `latest`.

Repository fixture runners perform this preparation for their selected local
port overrides. A standalone application or port runner must invoke it after
any cleanup of `output/`, before gopurs. The
[dependency preparation tests](https://github.com/0x000000000000000000001/gopurs/blob/main/tools/library-go-dependencies.test.mjs)
cover merging, conflicts and rollback without network access.

## Generic Go functions

Unconstrained Go type parameters can preserve a PureScript function's
polymorphism, including its callbacks and arrays:

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

The element types remain opaque to this implementation; it forwards each
element to the supplied callback and allocates a separate result slice.
The callback runs once per element, in input order. Use a concrete parameter
such as `string` or `int64` when the PureScript contract supplies that type.

The general generated bridge instantiates these Go type parameters with
`gopurs_runtime.Value`, including occurrences inside callbacks and slices.
`Value` is the runtime's common representation for values that need a tag and
payload. The handwritten implementation does not need to import the runtime.
This avoids converting every polymorphic element to `any`; it does not promise
native scalar specialization at every call or zero allocations. A Go
constraint requiring arithmetic, ordering or SDK methods is not established
by a PureScript type variable alone. Keep such operations in concrete
functions or pass the required operation as a typed callback.

Production examples include `OrdCharImpl`/`OrdStringImpl` in `gopurs-prelude`
and `AnyImpl` in `gopurs-arrays`. The
[generic FFI tests](https://github.com/0x000000000000000000001/gopurs/blob/main/tools/ffi-generics.test.mjs)
exercise their actual sources, several type parameters, callback results,
short-circuiting and preservation of input arrays.

## Foreign values and records

`Foreign` can cross as `any`, and `Array Foreign` as `[]any`. Values originating
in PureScript are exposed as Go primitives and containers; a Go author need not
inspect backend tags. Already-native opaque payloads retain their identity and
numeric types. Native nil slices remain nil when transported as Foreign.

Native `map[string]any` and `[]any` containers left by legacy FFI are also checked
for nested boxed values. Only branches needing conversion are copied; wholly native
subgraphs retain their identity. Cycles and repeated references to a converted
child are preserved. Opaque structs and unrelated SDK types are not traversed.

Absence of a map key remains distinct from a present nil value. Empty arrays,
false and zero are not collapsed into nil. Undefined travels as an opaque native
marker, distinct from nil, and becomes PureScript undefined again when returned
or passed to a callback. Transport retains keys with undefined values; the
Foreign/JSON library applies omission only when the application asks it to
serialize.

The marker is not a normal Go data value for arbitrary SDKs or Go's
JSON encoder: application serialization must run before those boundaries.
SQL formatting, JSON serialization and conversions such as decimal-to-string
remain application code. PureScript functions and algebraic constructors still
require an explicit conversion to native data.

A closed record has a complete, known set of fields. It maps to a locally
declared Go struct, for example:

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

Choose the boundary according to the meaning of the value:

| Value or operation | Supported approach and limit |
| --- | --- |
| Polymorphic value forwarded without inspection | An unconstrained Go type parameter and typed callbacks, as above. The general bridge can still use `Value`. |
| `Foreign` data consumed by an SDK | `any` or `[]any` with the native conversion described above; validate the shape before asserting SDK-specific types. |
| Open-row record helper | An explicit map-based helper can copy or forward opaque entries. Keep its Go parameters together; see [curried record helpers](migration.md#curried-record-helper-signatures). This does not establish an arbitrary struct layout. |
| ADT or newtype | Use an existing qualified representation, or explicit PureScript construction/deconstruction around the boundary. An available TAST type does not by itself define a public Go sum-type ABI. |
| Imported SDK struct | Use an opaque foreign handle, or a local checked struct and explicit SDK conversion. The bridge does not guess unavailable imported field layouts. |

`any` and `interface{}` are the same Go type. Prefer `any` in new code, but
changing its spelling does not remove boxing. A legacy polymorphic `any`
parameter may carry a backend value; it is not interchangeable with a
`Foreign` parameter that requests native conversion. Do not inspect backend
tags in ordinary application FFI. Document a specialized runtime boundary
when a library primitive needs one.

## Copying, aliases and mutation

Treat PureScript arrays, records and polymorphic inputs as borrowed values:
the caller retains ownership, and a pure FFI implementation must preserve them.
Aliases are references to the same underlying storage. Do not infer exclusive
ownership from a native signature or from a conversion happening in one
generated path: direct native calls and generic slices can retain aliases.
Allocate a result before updating an input; a slice or map copy is shallow
and does not give ownership of its nested handles or containers.

Closed-record structural conversion reconstructs the checked fields and
arrays, while opaque handles retain identity. Native `Foreign` subgraphs may
also be shared; only branches requiring conversion are copied. Consequently,
neither structural conversion nor `Foreign` transport promises a deep clone
of a resource graph. A nullable struct record has value semantics, distinct
from an opaque pointer handle with a shared lifetime.

For an effect that intentionally mutates a native resource, state which
operation creates it, which aliases may observe updates, whether concurrent
calls are allowed, and who closes it. Retaining a callback requires an
explicit stop operation. Use the acquisition cleanup contract above for
failure or cancellation, and the caller's explicit release after success.
The [structured-value tests](https://github.com/0x000000000000000000001/gopurs/blob/main/tools/ffi-structured.test.mjs)
and [Foreign transport tests](https://github.com/0x000000000000000000001/gopurs/blob/main/runtime/ffi_foreign_test.go)
cover conversion, nested values and identity boundaries.

## Diagnostics and verification

Unsupported signatures report the PureScript module, handwritten Go path,
declaration location, encountered Go/PureScript signatures and a corrective
hint. Signature analysis preserves every declared Go result, including callback
results. A shape the adapter cannot support is rejected instead of being
truncated. Fixed-size arrays, variadic parameters and multiple-result callbacks
currently require a supported boundary such as slices and single-result
callbacks. Type aliases requiring unavailable external export data are not
guessed from their spelling.

The standalone Go implementation and real PureScript imports are exercised by
`tools/ffi-aff-e2e.test.mjs`. It compares an official JavaScript Aff reference
with generated Go for both release and dev, including reused caches. The
JavaScript compiler host uses the WASM Go parser; the native compiler host uses
the same parser core. Neither claim makes network SDKs executable on Go/WASM.

### Tests in a development checkout

From a development checkout with its local library ports and typed frontend:

```sh
GO111MODULE=off go test -race ./runtime
node --test tools/ffi-aff.test.mjs tools/ffi-values.test.mjs tools/ffi-structured.test.mjs
node --test tools/ffi-docs.test.mjs
GOPURS_TEST_MODES=js node --test tools/ffi-aff-e2e.test.mjs
GOPURS_TEST_MODES=native-parallel node --test tools/ffi-aff-e2e.test.mjs
```

### Select the frontend

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

### Test the packaged compiler

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
