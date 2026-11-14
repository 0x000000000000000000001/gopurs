# Go AST and printer boundary

Emitters choose native types, boxing, names, evaluation order, parameters,
captures and tail-call targets from the typed IR. `Printer` renders their
nodes without reading or changing `CodegenState`.

`Printer.Builder` owns a mutable buffer per rendering operation. `withOut`
confines its handle to a synchronous callback; `emit` and `emitMany` chain
writes in order. Only the completed string leaves the operation. See the
[runtime and FFI contracts](runtime-ffi-contracts.md).

## Structured functions and declarations

`GoFuncBlock params stmts retType` contains ordered, named and typed
parameters, statements and a result type. Its `GoReturn` nodes are explicit,
so a body containing a tail-call loop does not receive an extra return.
`GoFuncLit params stmts retExpr retType` delegates to that representation
with a final `GoReturn retExpr`.

Emitters construct runtime wrappers by calling `gopurs_runtime.FuncN` with
a function node, and immediately invoked functions by calling the function
node with an empty argument list. They choose grouping and conversions before
printing. `GoFunctions.curriedFunction` supplies a dummy parameter for a
zero-argument function, groups up to ten arguments, and wraps excess arguments
in unary functions. The printer preserves these function boundaries.

`GoDecl` distinguishes cached values, structs, named functions, initialization
blocks and foreign getters. Structs retain generic parameters and typed fields;
functions retain parameter and result types and bodies. `CodeGen` supplies
ordered declaration groups to `GoFile`.

## Opaque fragments and imports

Some source remains opaque: local declaration/control fragments, portions of
record and ADT conversions, class-getter registration bodies and Rebox
assignments. The owning emitter remains responsible for their semantics.
FFI bridges are rendered from Go signatures and TAST metadata by `FfiBridge`.
The runtime is canonical Go source embedded from `runtime/runtime.go`.

`GoRaw` carries a `GoCode` with `{ text, imports }`. `rawGo` can recognize
references to `gopurs_runtime`, `math`, `sync` and `unsafe` when a fragment is
created; new fragments can supply dependencies explicitly. The import
collector does not scan the fragment again.

`GoImports.collectImports` owns the complete file import list. It traverses
structured declarations, expressions and types, collects opaque dependencies,
then sorts and deduplicates them. `Printer.printGoFile` prints the supplied
list without rendering declarations once to discover imports and again to
produce output.

## Owned references and qualification

`GoRef` distinguishes lexical references from owned symbols. Declarations,
constructors, pointers and worker signatures retain their producer's owner.
`GoTypeApply` retains type arguments; `GoTypeAssertion` takes a `GoType`.
`GoRecordFieldAccess` keeps the PureScript label separately from its Go field
name.

`GoSource.bindSources` resolves complete sources, including opaque fragments
and FFI, with the Go parser and `go/types`. `BoundGoCode` retains spans and
targets, distinguishing ordinary references, fields and methods.
`printBoundGoCode` and `boundImports` apply the consumer's package placement.
The cache and package-publication contract is described in
[build-cache.md](build-cache.md).

## Rules for changes

- Keep expressions structured while a downstream consumer can use their
  structure. Avoid printing a representable function just to wrap it in `rawGo`.
- Derive conversions from TAST in the emitter. The printer does not infer
  missing types, consult translation context or register Rebox helpers.
- Preserve parameter, statement and argument order and function boundaries.
  Printing must not duplicate evaluation or move it into a closure.
- Supply selected names and valid bodies to function nodes. `GoFuncBlock`
  does not rename parameters, invent temporaries or validate return paths.
- Declare dependencies for opaque fragments and extend `GoImports` when a
  structured node introduces a package reference.

`GoBoxIntArray`, `GoUnboxIntArray` and `GoFreshFilterArray` retain information
for `ArrayIntrinsics.normalizeFreshIntArrayRoundtrip`. The fold emitter
chooses whether a private filter buffer can be normalized before printing.
The printer renders surviving conversions and treats the freshness marker
transparently. See [array-roundtrip.md](array-roundtrip.md).

Relevant checks include `tools/go-imports.test.mjs`,
`tools/native-ffi.test.mjs` and the code-generation fixtures described in
[testing.md](testing.md).
