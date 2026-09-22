# Fused Either traversals

Two compiler changes reduce allocations in ordinary PureScript JSON decoders.
They do not introduce a schema-specific decoder or change the parser.

## Foreign.Object

`ObjectTraverse` recognizes `Data.TraversableWithIndex.traverseWithIndex` with
the standard `Foreign.Object.traversableWithIndexObject` and
`Data.Either.applicativeEither` dictionaries. It supports calls supplying two,
three or four arguments, including reusable partially applied traversals.

The ordinary object traversal folds repeated immutable insertions. Its Go
implementation copies the accumulating map at each successful insertion. The
specialized traversal allocates one private output map and fills it directly.
The input remains unchanged, and each invocation owns a fresh output map.

Keys use `sort.Strings`, matching the existing Go `_foldM` implementation.
Every callback still runs after a `Left`, and the first error is retained.
`traverseWithIndexDefault` is deliberately excluded: its `mapWithIndex` path
uses a different callback iteration order in the existing Go FFI. Unknown
dictionaries and dictionary accesses already reduced by PBO keep their generic
path.

## Array traversals and callbacks

Array/Either fusion also covers nonindexed `Data.Traversable.traverse`, the
statically known `traversableArray.traverse` dictionary field, and
`traverseArrayImpl` supplied with the exact standard Either `apply`, `map`,
`pure` and Array `append` methods. These paths occur in the PBO TAST decoder.
Unknown dictionaries and other Applicatives retain their original code.
Arguments are captured in order at their original partial-application stage.

The existing Array/Either fusion can now consume an adjacent two-parameter
lambda, or a unary lambda for the nonindexed traversal, as a native Go function.
The callback's result retains its generated
representation, including a native Either and native record success payload.
The loop calls that function directly and inspects the native result, avoiding
the previous callback boxing, `Apply2`, and immediate result unboxing.

The callback is captured at the original application stage. Computation between
the two lambdas keeps the original curried path. All callbacks still run in
index order; only the first error is returned. The output array remains
`[]Value`, so boxing the final element payload is still required. This is a
local consumer optimization, not a complete native dictionary/callback ABI.

The same campaign changes PBO's own `getField` and `getFieldOptional'` helpers
to construct `AtKey` wrappers only after a decoding failure. Missing fields,
`null`, nested errors and the number of decoder calls retain their existing
semantics. The parser and TypeTable helpers are unchanged.

## Validation and measurements

- `NativeTraverseCallback` covers native record results, first-error priority,
  calls after failure, empty arrays, reusable captures and staged lambdas.
- `ArrayTraverseEither` covers the nonindexed forms, partial applications,
  strict callback evaluation after failure, argument order, and unknown/Maybe
  Applicative fallbacks.
- `ObjectTraverseEither` exercises the actual generated object loop, success,
  first-error priority, all callbacks, empty input, reusable captures and
  unchanged input.
- The object traversal generator tests cover supplied argument counts,
  dictionary/default fallbacks, and 270 comparisons against the existing
  traversal, including numeric and non-ASCII keys and output independence.
- Real JSON and TAST diagnostic measurements use the original PureScript
  programs and independent fixed oracles.
- PBO's `json-fields` tests cover both changed helpers, including missing/null
  fields, complete error paths and exact decoder call counts.

The paired campaign, current numbers, binary sizes and limitations are recorded
in [the altbak report](../../../altbak.pub-gopurs/docs/benchmark-results/2026-09-22-fused-json-traversals.md).
