# Module map

This map describes the current source tree: one row per PureScript module,
its conceptual phase, responsibility and direct `Gopurs.*` imports. Phases
classify responsibilities; they do not prescribe a future directory layout.
See [architecture.md](architecture.md) for execution order and
[naming.md](naming.md) for terminology.

The map covers **104 modules**. `node tools/check-docs.mjs` verifies source
paths, duplicate rows and complete coverage against `src/`. Import lists below
are sorted and deduplicated from the module source.

## Driver — orchestration

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.Driver`<br>`src/Gopurs/Driver.purs` | Driver | Orders compilation phases and exposes `compile`. | `Cache.HeaderPreload`, `Cache.Modules`, `Driver.Build`, `Driver.Cache`, `Driver.Config`, `Driver.Manifest`, `Driver.Output`, `Driver.Prepare`, `Driver.Prepare.Persistent`, `Host.Diagnostics`, `Host.Metrics` |
| `Gopurs.Driver.Build`<br>`src/Gopurs/Driver/Build.purs` | Driver | Selects the PBO builder, publishes signatures and feeds emission batches. | `Cache.HeaderPreload`, `Cache.Modules`, `Cache.Wire`, `Driver.Cache`, `Driver.Config`, `Driver.Prepare`, `Emission`, `Host.Metrics` |
| `Gopurs.Driver.Config`<br>`src/Gopurs/Driver/Config.purs` | Driver | Reads CLI options and environment settings with their bounds. | `Emission` |
| `Gopurs.Driver.Manifest`<br>`src/Gopurs/Driver/Manifest.purs` | Driver | Constructs, reads and compares input/output manifests. | `Cache.Wire`, `ExprContext` |
| `Gopurs.Driver.Output`<br>`src/Gopurs/Driver/Output.purs` | Driver | Emits prepared modules, runtime source and entry points. | `CodeGen`, `CodegenState`, `Driver.Cache`, `ExprContext`, `FfiBridge`, `FfiSupport`, `FfiTypes`, `GoSource`, `GoSymbol`, `Runtime.Source` |
| `Gopurs.Emission`<br>`src/Gopurs/Emission.purs` | Driver | Schedules emission batches, dependencies and fiber backpressure. | `Host.Metrics` |

## Prepare — loading, metadata and specialization

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.AdtMetadata`<br>`src/Gopurs/AdtMetadata.purs` | Prepare | Derives ADT pointer and enum representations from TAST. | `GoAst` |
| `Gopurs.ClassMetadata`<br>`src/Gopurs/ClassMetadata.purs` | Prepare | Indexes classes, superclasses and methods by label. | — |
| `Gopurs.ConstructorLayout`<br>`src/Gopurs/ConstructorLayout.purs` | Prepare | Resolves constructor identity, fields and generic arguments. | `CodegenState`, `ConstructorMetadata`, `GoAst`, `GoSymbol`, `GoTypes` |
| `Gopurs.ConstructorMetadata`<br>`src/Gopurs/ConstructorMetadata.purs` | Prepare | Describes constructors, fields and identities. | `GoAst` |
| `Gopurs.Driver.Prepare`<br>`src/Gopurs/Driver/Prepare.purs` | Prepare | Loads TAST, assembles metadata and prepares specialized modules. | `AdtMetadata`, `Cache.LookupTable`, `Cache.Wire`, `ClassMetadata`, `CodegenState`, `ConstructorMetadata`, `Driver.Cache`, `Driver.Prepare.Projection`, `Driver.Prepare.Proof`, `GlobalTypes`, `Host.Metrics`, `Monomorphization`, `Preparation`, `ReboxMetadata` |
| `Gopurs.Driver.Prepare.Persistent`<br>`src/Gopurs/Driver/Prepare/Persistent.purs` | Prepare | Preloads, projects and publishes persistent preparation artifacts. | `Cache.LookupTable`, `Cache.Wire`, `ClassMetadata`, `Driver.Cache`, `Driver.Manifest`, `Driver.Prepare`, `Driver.Prepare.Projection`, `Driver.Prepare.ProjectionCapsule`, `Driver.Prepare.Proof`, `Host.Metrics`, `Monomorphization`, `Preparation` |
| `Gopurs.Driver.Prepare.Projection`<br>`src/Gopurs/Driver/Prepare/Projection.purs` | Prepare | Splits and merges prepared-module projections. | `Cache.LookupTable`, `CodegenState` |
| `Gopurs.Driver.Prepare.ProjectionCapsule`<br>`src/Gopurs/Driver/Prepare/ProjectionCapsule.purs` | Prepare | Packs and reopens projections as verified fragments. | `Cache.LookupTable`, `CodegenState`, `Driver.Prepare.Projection` |
| `Gopurs.Driver.Prepare.Proof`<br>`src/Gopurs/Driver/Prepare/Proof.purs` | Prepare | Computes preparation identity and reuse proofs. | `Cache.Wire` |
| `Gopurs.GlobalTypes`<br>`src/Gopurs/GlobalTypes.purs` | Prepare | Resolves global types of bindings and FFI imports. | — |
| `Gopurs.Monomorphization`<br>`src/Gopurs/Monomorphization.purs` | Prepare | Collects, propagates and applies monomorphic specializations. | `Monomorphization.ForeignForwarders`, `NativeRecordArgs` |
| `Gopurs.Monomorphization.ForeignForwarders`<br>`src/Gopurs/Monomorphization/ForeignForwarders.purs` | Prepare | Recognizes FFI wrappers excluded from specialized copies. | — |
| `Gopurs.Preparation`<br>`src/Gopurs/Preparation.purs` | Prepare | Runs one ordered, bounded round of preparation jobs in `Aff`. | — |
| `Gopurs.ReboxMetadata`<br>`src/Gopurs/ReboxMetadata.purs` | Prepare | Indexes fields converted between native and boxed representations. | `ClassMetadata`, `ConstructorMetadata`, `GoAst` |

## Rewrite — backend IR passes

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.BorrowedObjects`<br>`src/Gopurs/BorrowedObjects.purs` | Rewrite | Proves and rewrites read-only borrowing of JSON objects. | `Cache.LookupTable`, `CodegenState` |
| `Gopurs.ClosedDictionaries`<br>`src/Gopurs/ClosedDictionaries.purs` | Rewrite | Rewrites and caches closed dictionaries. | `ClosedDictionaries.Admission`, `CodegenState` |
| `Gopurs.ClosedDictionaries.Admission`<br>`src/Gopurs/ClosedDictionaries/Admission.purs` | Rewrite | Separates existing-binding reuse from new cache admission. | `Cache.LookupTable`, `ClosedDictionaries.Scope`, `CodegenState` |
| `Gopurs.ClosedDictionaries.Scope`<br>`src/Gopurs/ClosedDictionaries/Scope.purs` | Rewrite | Checks lexical closure and absence of effects or recursion. | — |
| `Gopurs.DecoderSchemas`<br>`src/Gopurs/DecoderSchemas.purs` | Rewrite | Specializes JSON decoder schemas and workers. | `Cache.LookupTable`, `CodegenState`, `DecoderSchemas.Admission`, `DecoderSchemas.Source`, `DecoderSchemas.Workers`, `GoAst` |
| `Gopurs.DecoderSchemas.Admission`<br>`src/Gopurs/DecoderSchemas/Admission.purs` | Rewrite | Admits decoder schemas and their programs. | `DecoderSchemas.Programs`, `DecoderSchemas.Source`, `DecoderSchemas.Types` |
| `Gopurs.DecoderSchemas.Programs`<br>`src/Gopurs/DecoderSchemas/Programs.purs` | Rewrite | Derives dictionary programs from schemas. | `DecoderSchemas.Source`, `DecoderSchemas.Types` |
| `Gopurs.DecoderSchemas.Source`<br>`src/Gopurs/DecoderSchemas/Source.purs` | Rewrite | Resolves source dictionaries and schema symbols. | — |
| `Gopurs.DecoderSchemas.Types`<br>`src/Gopurs/DecoderSchemas/Types.purs` | Rewrite | Defines schemas, programs and admission weights. | — |
| `Gopurs.DecoderSchemas.Workers`<br>`src/Gopurs/DecoderSchemas/Workers.purs` | Rewrite | Constructs decoder-schema workers. | `DecoderSchemas.Types`, `GoAst`, `GoSymbol`, `Printer` |
| `Gopurs.FunctionFusion`<br>`src/Gopurs/FunctionFusion.purs` | Rewrite | Admits and rewrites counted function producers. | `WorkerNames` |
| `Gopurs.ImmediateApplications`<br>`src/Gopurs/ImmediateApplications.purs` | Rewrite | Admits and reduces safe immediate applications. | `ImmediateApplications.Scope` |
| `Gopurs.ImmediateApplications.Scope`<br>`src/Gopurs/ImmediateApplications/Scope.purs` | Rewrite | Computes occurrences, substitutions and local-level renaming. | — |
| `Gopurs.NativeRecordArgs`<br>`src/Gopurs/NativeRecordArgs.purs` | Rewrite | Checks tail-call behavior and chooses record argument ABIs. | `GoAst`, `NativeRecordArgs.Projection`, `NativeRecordArgs.Source` |
| `Gopurs.NativeRecordArgs.Projection`<br>`src/Gopurs/NativeRecordArgs/Projection.purs` | Rewrite | Projects fields and admits native result layouts. | `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.NativeRecordArgs.Source`<br>`src/Gopurs/NativeRecordArgs/Source.purs` | Rewrite | Proves shared reader sources before monomorphization. | `NativeRecordArgs.Projection` |
| `Gopurs.Ownership`<br>`src/Gopurs/Ownership.purs` | Rewrite | Computes consuming-worker fixed points and selects fresh calls. | `CodegenState`, `GoAst`, `GoSymbol`, `Ownership.Analysis`, `Ownership.Candidates`, `Ownership.Types`, `Ownership.Workers` |
| `Gopurs.Ownership.Analysis`<br>`src/Gopurs/Ownership/Analysis.purs` | Rewrite | Analyzes paths, scalar terms, tree terms and freshness. | `CodegenState`, `GoAst`, `Ownership.Types`, `PrimitiveExprs` |
| `Gopurs.Ownership.Candidates`<br>`src/Gopurs/Ownership/Candidates.purs` | Rewrite | Admits layouts/signatures and reserves worker/helper names. | `CodegenState`, `GoAst`, `GoTypes`, `Ownership.Analysis`, `Ownership.Types` |
| `Gopurs.Ownership.Types`<br>`src/Gopurs/Ownership/Types.purs` | Rewrite | Defines ownership-analysis types. | `CodegenState`, `GoAst` |
| `Gopurs.Ownership.Workers`<br>`src/Gopurs/Ownership/Workers.purs` | Rewrite | Emits captures, reusable cells and consuming declarations. | `CodegenState`, `GoAst`, `GoSymbol`, `Ownership.Analysis`, `Ownership.Types` |
| `Gopurs.ThunkFusion`<br>`src/Gopurs/ThunkFusion.purs` | Rewrite | Admits and rewrites thunk producers into counted workers. | `WorkerNames` |
| `Gopurs.WorkerNames`<br>`src/Gopurs/WorkerNames.purs` | Rewrite | Reserves worker names for fusion passes. | `GoAst` |

## Emit — Go translation and rendering

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.AdtExprs`<br>`src/Gopurs/AdtExprs.purs` | Emit | Constructs ADT representations and constructor reuse. | `CodegenState`, `ConstructorLayout`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `Printer` |
| `Gopurs.ArrayIntrinsics`<br>`src/Gopurs/ArrayIntrinsics.purs` | Emit | Recognizes and emits map, filter and foldl. | `ArrayIntrinsics.Index`, `ArrayIntrinsics.Source`, `CallAnalysis`, `CallArguments`, `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.ArrayIntrinsics.Index`<br>`src/Gopurs/ArrayIntrinsics/Index.purs` | Emit | Admits indexing and converts selected elements. | `ArrayIntrinsics.Source`, `CallAnalysis`, `CallArguments`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.ArrayIntrinsics.Source`<br>`src/Gopurs/ArrayIntrinsics/Source.purs` | Emit | Binds evaluated arrays without copying their elements. | `GoAst` |
| `Gopurs.ArrayTraverse`<br>`src/Gopurs/ArrayTraverse.purs` | Emit | Emits Either traversal loops over arrays. | `AdtExprs`, `CallAnalysis`, `CallArguments`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.BindingExprs`<br>`src/Gopurs/BindingExprs.purs` | Emit | Emits local bindings, recursion and initialization. | `CallAnalysis`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoTypes`, `LocalWorkers`, `Printer` |
| `Gopurs.CallAnalysis`<br>`src/Gopurs/CallAnalysis.purs` | Emit | Analyzes call targets, application spines and closures. | `Cache.LookupTable`, `ExprAnalysis`, `ExprContext`, `GoAst` |
| `Gopurs.CallArguments`<br>`src/Gopurs/CallArguments.purs` | Emit | Translates, captures and adapts call arguments. | `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.CallExprs`<br>`src/Gopurs/CallExprs.purs` | Emit | Selects calls, overapplication and tail jumps. | `ArrayIntrinsics`, `ArrayTraverse`, `CallAnalysis`, `CallArguments`, `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `ObjectTraverse` |
| `Gopurs.CodeGen`<br>`src/Gopurs/CodeGen.purs` | Emit | Dispatches expression translation and assembles a module's Go file. | `AdtExprs`, `ArrayIntrinsics`, `BindingExprs`, `BorrowedObjects`, `Cache.LookupTable`, `CallExprs`, `ClosedDictionaries`, `CodegenState`, `ConstructorExprs`, `ControlExprs`, `DecoderSchemas`, `EffectExprs`, `ExprAnalysis`, `ExprContext`, `FunctionExprs`, `FunctionFusion`, `GoAst`, `GoConversions`, `GoImports`, `GoSymbol`, `GoTypes`, `ImmediateApplications`, `LiteralExprs`, `ModuleBindings`, `ModuleDeclarations`, `Ownership`, `PrimitiveExprs`, `Printer`, `RecordExprs`, `ThunkFusion`, `TypedExprs` |
| `Gopurs.CodegenState`<br>`src/Gopurs/CodegenState.purs` | Emit | Defines immutable emission metadata and mutable translation state. | `AdtMetadata`, `Cache.LookupTable`, `ClassMetadata`, `ConstructorMetadata`, `GoAst`, `GoSymbol`, `ReboxMetadata` |
| `Gopurs.ConstructorExprs`<br>`src/Gopurs/ConstructorExprs.purs` | Emit | Translates constructor definitions and saturated construction. | `AdtExprs`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.ControlExprs`<br>`src/Gopurs/ControlExprs.purs` | Emit | Translates branches, alternatives and control flow. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `Printer` |
| `Gopurs.EffectExprs`<br>`src/Gopurs/EffectExprs.purs` | Emit | Translates effect expressions and their wrappers. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `Printer` |
| `Gopurs.ExprAnalysis`<br>`src/Gopurs/ExprAnalysis.purs` | Emit | Provides read-only expression analyses. | `ExprContext`, `GoAst`, `GoTypes` |
| `Gopurs.ExprContext`<br>`src/Gopurs/ExprContext.purs` | Emit | Carries translation context, results and callbacks. | `CodegenState`, `GoAst` |
| `Gopurs.FfiBridge`<br>`src/Gopurs/FfiBridge.purs` | Emit | Assembles typed FFI bridges. | `CodegenState`, `FfiBridge.Render`, `FfiBridge.Signatures`, `FfiTypes`, `GoAst` |
| `Gopurs.FfiBridge.Render`<br>`src/Gopurs/FfiBridge/Render.purs` | Emit | Renders wrappers, workers and missing-FFI declarations. | `FfiBridge.Signatures`, `FfiBridge.TypeSupport`, `FfiBridge.Values`, `FfiTypes`, `GoTypes` |
| `Gopurs.FfiBridge.Signatures`<br>`src/Gopurs/FfiBridge/Signatures.purs` | Emit | Matches foreign declarations to native call signatures. | `CodegenState`, `FfiBridge.TypeSupport`, `FfiTypes`, `GoAst`, `GoSymbol` |
| `Gopurs.FfiBridge.TypeSupport`<br>`src/Gopurs/FfiBridge/TypeSupport.purs` | Emit | Instantiates bridge types and generic arguments. | `FfiTypes` |
| `Gopurs.FfiBridge.Values`<br>`src/Gopurs/FfiBridge/Values.purs` | Emit | Adapts FFI arguments, callbacks and results. | `FfiBridge.TypeSupport`, `FfiTypes`, `GoTypes` |
| `Gopurs.FfiSupport`<br>`src/Gopurs/FfiSupport.purs` | Emit | Parses Go FFI and exposes support operations. | `FfiTypes` |
| `Gopurs.FfiTypes`<br>`src/Gopurs/FfiTypes.purs` | Emit | Defines FFI bridge types. | — |
| `Gopurs.FunctionExprs`<br>`src/Gopurs/FunctionExprs.purs` | Emit | Translates abstractions, applications and functions. | `CallAnalysis`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoSymbol` |
| `Gopurs.GoAst`<br>`src/Gopurs/GoAst.purs` | Emit | Defines Go expression, type and declaration nodes. | `GoCode`, `GoSymbol` |
| `Gopurs.GoCode`<br>`src/Gopurs/GoCode.purs` | Emit | Carries opaque fragments and import dependencies. | `GoSymbol` |
| `Gopurs.GoConversions`<br>`src/Gopurs/GoConversions.purs` | Emit | Selects conversions, boxing and Value access. | `CodegenState`, `GoAst`, `GoConversions.NativeAdts`, `GoConversions.Rebox`, `Printer` |
| `Gopurs.GoConversions.NativeAdts`<br>`src/Gopurs/GoConversions/NativeAdts.purs` | Emit | Describes native Maybe, Either and Tuple layouts. | `GoAst`, `Printer` |
| `Gopurs.GoConversions.Rebox`<br>`src/Gopurs/GoConversions/Rebox.purs` | Emit | Emits conversion helpers and tracks Rebox requests. | `CodegenState`, `GoAst`, `GoSymbol`, `GoTypes`, `Host.Diagnostics`, `Printer`, `ReboxMetadata` |
| `Gopurs.GoFunctions`<br>`src/Gopurs/GoFunctions.purs` | Emit | Constructs parameters, iteration bindings and curried wrappers. | `GoAst` |
| `Gopurs.GoImports`<br>`src/Gopurs/GoImports.purs` | Emit | Collects and renders Go module imports. | `GoAst`, `GoCode`, `GoSymbol` |
| `Gopurs.GoSource`<br>`src/Gopurs/GoSource.purs` | Emit | Binds Go sources to symbols and assembles package files. | `FfiSupport`, `GoCode`, `GoSymbol` |
| `Gopurs.GoSymbol`<br>`src/Gopurs/GoSymbol.purs` | Emit | Defines and compares Go symbols. | — |
| `Gopurs.GoTypes`<br>`src/Gopurs/GoTypes.purs` | Emit | Translates PureScript annotations into Go types. | `AdtMetadata`, `GoAst`, `GoSymbol` |
| `Gopurs.Int32Loops`<br>`src/Gopurs/Int32Loops.purs` | Emit | Proves and versions scalar loops with signed 32-bit state. | `GoAst`, `GoFunctions` |
| `Gopurs.LiteralExprs`<br>`src/Gopurs/LiteralExprs.purs` | Emit | Translates array and record literals with their conversions. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `Printer`, `RecordExprs` |
| `Gopurs.LocalWorkers`<br>`src/Gopurs/LocalWorkers.purs` | Emit | Emits local workers, parameters and signatures. | `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoSymbol`, `GoTypes`, `Int32Loops` |
| `Gopurs.ModuleBindings`<br>`src/Gopurs/ModuleBindings.purs` | Emit | Emits module functions, signatures and tail-call groups. | `CallAnalysis`, `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoSymbol`, `GoTypes`, `ModuleWorkers`, `NativeRecordArgs` |
| `Gopurs.ModuleDeclarations`<br>`src/Gopurs/ModuleDeclarations.purs` | Emit | Emits ADT structs and class-getter registration. | `CodegenState`, `GoAst`, `GoSymbol`, `GoTypes` |
| `Gopurs.ModuleWorkers`<br>`src/Gopurs/ModuleWorkers.purs` | Emit | Emits module workers and wrappers. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoSymbol`, `Int32Loops` |
| `Gopurs.ObjectTraverse`<br>`src/Gopurs/ObjectTraverse.purs` | Emit | Emits Either traversal loops over objects. | `AdtExprs`, `CallAnalysis`, `CallArguments`, `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.PrimitiveExprs`<br>`src/Gopurs/PrimitiveExprs.purs` | Emit | Translates primitives and their conversions. | `CodegenState`, `GoAst`, `GoConversions` |
| `Gopurs.Printer`<br>`src/Gopurs/Printer.purs` | Emit | Renders Go declarations as text. | `GoAst`, `GoCode`, `GoSymbol`, `Printer.Builder` |
| `Gopurs.Printer.Builder`<br>`src/Gopurs/Printer/Builder.purs` | Emit | Owns the private mutable buffer for one rendering operation. | — |
| `Gopurs.RecordExprs`<br>`src/Gopurs/RecordExprs.purs` | Emit | Translates records and selects their representations. | `CodegenState`, `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.TypedExprs`<br>`src/Gopurs/TypedExprs.purs` | Emit | Propagates type annotations and constructs class dictionaries. | `CallAnalysis`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes` |

## Cache — protocol and reuse

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.BuildCache`<br>`src/Gopurs/BuildCache.purs` | Cache | Defines the versioned storage protocol and JSON exchanges. | — |
| `Gopurs.Cache.BindingCapsule`<br>`src/Gopurs/Cache/BindingCapsule.purs` | Cache | Encodes and selects binding capsules. | `Cache.Wire` |
| `Gopurs.Cache.HeaderPreload`<br>`src/Gopurs/Cache/HeaderPreload.purs` | Cache | Preloads and decodes cache headers. | `Cache.Wire`, `Driver.Cache`, `Host.Metrics`, `Preparation` |
| `Gopurs.Cache.LookupTable`<br>`src/Gopurs/Cache/LookupTable.purs` | Cache | Represents proven, complete cache lookups. | `Cache.Wire` |
| `Gopurs.Cache.Modules`<br>`src/Gopurs/Cache/Modules.purs` | Cache | Manages prepared, PBO and emission keys and per-module reuse. | `Cache.BindingCapsule`, `Cache.HeaderPreload`, `Cache.Wire`, `CodeGen`, `CodegenState`, `Driver.Cache`, `Driver.Output`, `Driver.Prepare`, `ExprContext`, `GoSymbol`, `Host.Metrics`, `Preparation` |
| `Gopurs.Cache.Pool`<br>`src/Gopurs/Cache/Pool.purs` | Cache | Interns canonical JSON fragments within an artifact. | — |
| `Gopurs.Cache.TypeMemo`<br>`src/Gopurs/Cache/TypeMemo.purs` | Cache | Memoizes immutable types by canonical index. | — |
| `Gopurs.Cache.Wire`<br>`src/Gopurs/Cache/Wire.purs` | Cache | Provides portable codecs and shared tables. | `Cache.Pool`, `Cache.TypeMemo`, `GoAst`, `GoSymbol` |
| `Gopurs.Driver.Cache`<br>`src/Gopurs/Driver/Cache.purs` | Cache | Owns the storage session, staging and FFI recipe. | `BuildCache`, `Driver.Config`, `Host.Metrics` |

## Host — compiler services

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.Host.Diagnostics`<br>`src/Gopurs/Host/Diagnostics.purs` | Host | Writes diagnostics and marks them as non-replayable. | — |
| `Gopurs.Host.Metrics`<br>`src/Gopurs/Host/Metrics.purs` | Host | Measures phases and accumulates emission timings. | — |
| `Gopurs.Host.Resident`<br>`src/Gopurs/Host/Resident.purs` | Host | Runs a file-triggered resident compiler session. | `Driver`, `Driver.Config`, `Driver.Prepare` |
| `Gopurs.Host.Snapshot`<br>`src/Gopurs/Host/Snapshot.purs` | Host | Copies and re-executes the compiler from private storage. | — |

## Runtime — embedded source

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Gopurs.Runtime.Source`<br>`src/Gopurs/Runtime/Source.purs` | Runtime | Exposes embedded Go runtime source. | — |

## Entry point

| Module and source | Phase | Responsibility | Direct dependencies |
| --- | --- | --- | --- |
| `Main`<br>`src/Main.purs` | Entry | Enters the CLI: snapshot, options, execution and exit status. | `Driver`, `Driver.Config`, `Host.Resident`, `Host.Snapshot` |

## Shared boundaries

The present imports include deliberate sharing across conceptual phases:
`Driver.Prepare` and rewrite passes use metadata types from `CodegenState`;
`Cache.Modules` calls `Driver.Output`; `Monomorphization` consults record-reader
proofs from `NativeRecordArgs`. `Host.Resident` is a service entry point and
depends on the driver. The table records these actual relationships.

Update a module row when its declaration, path or imports change. Add or
remove rows with the corresponding source changes, then run the documentation
checker. Foreign companions and dynamically loaded resources are not
PureScript import edges; consult the [FFI contracts](runtime-ffi-contracts.md)
and [development guide](development.md) before removing them.
