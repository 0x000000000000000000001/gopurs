# Shared representations for higher-order containers

A native list can become expensive when it crosses the generic function ABI:
converting its element representation reconstructs the list, its tuples and
their payloads. Gopurs can keep one representation for an eligible product and
the lists containing it, so those boundaries do not require recursive copies.

The source analysis records concrete products used by callbacks or functions
stored as values. A product qualifies for the current representation policy
when it is a `Data.Tuple.Tuple` pairing a nonempty closed record of primitive
scalars with a known, nonparameterized enum. No application module or record
label is recognized specially. A direct monomorphic use alone does not select
the product: keeping its native fields avoids extra boxing in that case.

Selected tuples use the existing pointer constructor with `Value` payloads;
their `Data.List.Types.List` containers use the existing `Cons[Value]` layout.
The generic definitions of `Tuple a b` and `List a` retain their type parameters.
Unknown types, open records, composite record fields, incomplete constructor
arities and elided constructors retain their ordinary representations.

`SharedContainerUsage` collects source evidence before signatures are published.
The selected pairs travel in pointer metadata, including preparation snapshots
and projections. `SharedContainerLayout` checks the declaration metadata and
selects the same arguments for value types, generic fields and constructor
creation. Record field order is normalized for selection, while field names
and types remain significant.

This is a source-based representation heuristic. Later optimization may remove
some higher-order boundaries; source evidence is not a proof of a runtime gain
for every program. It authorizes neither mutation nor object reuse. No new
runtime tag, conversion cache or allocation arena is introduced, and the
existing record and FFI conversion contracts remain in force.

## Validation

The `shared-container-layout` and `shared-container-usage` tool tests cover
selection, rejection, type applications and agreement between type and
constructor selection. `SharedListPayloads` exercises callbacks, retained
functions, list versions and FFI boundaries. `SharedListNative` controls the
direct monomorphic case, which must retain its native layout.

Performance comparisons use the application benchmark in
`bench-candidates/cst-bench`, alongside the published `altbak.pub` README.
Profile-guided optimization by the Go compiler is a separate build setting;
ordinary gopurs output does not imply a PGO training run.
