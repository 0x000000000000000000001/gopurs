# Tree-cell reuse

[`Gopurs.Ownership`](../src/Gopurs/Ownership.purs) analyzes the transformed
`BackendModule` after PBO and specialized CodeGen passes, before tail-call
analysis and ordinary function generation. It creates workers that consume
exclusively owned trees. Ordinary functions retain persistent semantics.

## Proof and eligibility

A consuming worker receives a forest whose nonempty nodes are exclusive and
disjoint, and returns a forest with the same property. A fresh root containing
a shared child does not qualify. Empty `nil` values may be shared.

Eligible ADTs are monomorphic, declared in the module and represented by a
pointer, with one constructor carrying fields and at most one empty
constructor. Fields contain scalars, enums or the same tree type. Functions
are first-order, accept scalar or tree parameters, and return a tree. Opaque
fields, functions, arrays and representation conversions are excluded.

Local references are resolved lexically. Projection aliases identify canonical
paths consisting of a root and field indices. Retained subtrees and arguments
to consuming workers must be disjoint: no duplicate paths and no path that is
a prefix of another. A result cannot retain both a parent and its descendant,
or place the same child in two fields.

Calls are validated to a fixed point across module functions. Rejecting a
candidate also rejects callers relying on its contract. Unknown calls, FFI,
closures and effects provide no consumption proof. The existing `Rc` field
is not an ownership proof and this pass requires no reference-counting protocol.

## Mutation and evaluation order

Required field reads are captured before mutation. Dead cells are roots or
prefixes no longer retained as result subtrees. Each branch has its own proof;
a guard grants no mutation rights on its failure path. An old root or subtree
still observable by the continuation prevents consumption.

Workers may receive an extra donor cell, allowing an insertion to pass its
former parent to a balancing function. Cells leave the available pool when
consumed. A consuming `let` removes old aliases before introducing the result.
Terminal recursive calls become loops.

Strict prefixes dereferenced during capture form a known-nonnull pool, used
before nullable donors. Conditional scalar reads do not prove nonnullness.
Remaining cells pass between sibling constructions and arguments; an external
or terminal recursive call chooses its donor after its arguments have consumed
their cells. Pools belong to individual plans, branches and continuations.
All addresses are captured before mutation, including when reusing a parent
before a child.

## Responsibilities

- `Ownership.prepare` collects candidates, validates them to a fixed point,
  publishes declarations and signatures, and rewrites proven fresh calls.
- `Ownership.Candidates` checks layouts and signatures. Worker/helper names
  are reserved once in source order, including names of later rejected
  candidates and their normalized Go forms.
- `Ownership.Types` defines paths, environments and tree terms.
  `Ownership.Analysis` proves disjointness and freshness and records
  continuation observations.
- `Ownership.Workers` captures values, plans reusable cells and emits Go.
  Each attempt has its own name counter; a failed attempt rejects the whole
  body. The same generator validates candidates and emits surviving workers.

`snapshot` converts `Keep` paths into already-read `Existing` values after
`planCells` checks disjointness. `reusablePaths` distinguishes dereferenced
prefixes from conditional scalar reads. `emitLet` removes old aliases and
resets the donor before the continuation; `emitTerminal` selects the donor
after emitting arguments.

## Call selection and source annotations

A consuming worker is selected only when every tree argument is a recursively
fresh admissible construction or an empty value. A local variable or borrowed
child does not establish freshness. Recognition never depends on benchmark,
function or type names.

The TAST's optional `bindingUsage` and `variableUse` describe direct local uses,
not exclusivity or alias freedom. Missing facts are unknown. PBO invalidates
source usage identities and facts before monomorphization and conversion.
This pass recomputes its proof on lexical references and paths in the final IR;
`lastLocalUse` and escaping-context flags never authorize mutation.

## Validation

After building:

```sh
node --test tools/owned-trees.test.mjs
bin/test OwnedTrees
```

The tests cover disjointness, aliases, observable continuations, guards,
transitive rejection, recursive candidate groups, naming collisions, capture
before mutation, sibling arguments and nullable donors. The fixture verifies
rotations, key order, colors, black height, duplicates, retained versions and
shared children with names independent of the benchmark.
