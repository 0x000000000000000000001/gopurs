# JSON decoder code generation

These general compiler rules reduce intermediate representations in composed
PureScript decoders. They run after PBO and do not encode a JSON schema,
benchmark name or special TAST decoder. Parsing and decoder validation retain
their existing semantics.

## Array/Either traversal

[ArrayTraverse](../src/Gopurs/ArrayTraverse.purs) recognizes qualified indexed
traversals with the standard Array and Either dictionaries. It replaces the
intermediate mapped-result array and repeated reconstruction with one private
output buffer. Two supplied arguments create a reusable binary worker, three
capture the callback, and four execute the traversal.

Arguments are captured once, in order and at their original partial-application
stage. Every callback runs in index order, including after the first failure:
the original strict mapping evaluated them all. Only the first `Left` is
returned. Inputs stay unchanged and repeated invocations own separate buffers.
Unknown dictionaries, other Applicatives and unsupported arities retain the
generic path. Constructor tags and fields come from shared ADT metadata.

Nonindexed forms and direct native callbacks are described in
[fused Either traversals](json-traversals.md). The output buffer remains
`[]Value`; callback payloads may still require boxing even when the callback
itself returns a native result.

## Immediate applications

[ImmediateApplications](../src/Gopurs/ImmediateApplications.purs) reduces a
known unary lambda application after function/thunk fusion and before ownership
analysis and tail-call optimization. It can move an application into a `Let`
or branches only when every result permits eliminating that application.

The substituted argument must already be a value: a local, scalar literal or
lambda without `LetRec`. Locals and scalar literals may replace repeated uses;
a lambda cannot replace multiple uses within one body. When distributing into
branches, the total duplicated syntax must fit a 128-node growth budget,
counting all substitutions rather than only the largest branch. Zero or one
substitution imposes no size limit. Branches have at most four explicit
alternatives and must also have a reducible default; an explicit failure is
also admitted without evaluating the callback.

Calls, global getters, field accesses, constructors and arrays are not movable
arguments. Recursive scopes, opaque function results and partial applications
of multi-parameter lambdas retain their calls. Transplanted bound variables
receive fresh levels; free variables keep their original bindings. Recognition
passes through `Typed` and `TypeApp` without discarding annotations globally.
Callback effects and divergence stay deferred until invocation.

## Native Maybe, Either and Tuple results

[TypedExprs](../src/Gopurs/TypedExprs.purs) can preserve a `TypeStructValue`
when the expected representation is boxed but its ADT identity matches the
recognized native result. An annotation alone then does not force an
intermediate `Value` wrapper.

| Native value layout | Payload and discriminant slots |
| --- | --- |
| `Maybe` | Payload `V0`, presence `V1` |
| `Either` | Left `V0`, right `V1`, branch `V2` |
| `Tuple` | First `V0`, second `V1` |

For these value layouts, payload slots are `Value` and discriminants are
`bool`. More precise native pointer layouts retain their own typed payloads
and conversion rules. Shared field metadata selects the correct slot, notably
`V1` for `Right`. Dynamic boundaries and public `Get_…` wrappers still box when
required. Preserving one native envelope does not make every nested payload
native or generalize the layout to arbitrary ADTs.

## Validation

After building:

```sh
node --test tools/array-traverse-either.test.mjs tools/immediate-applications.test.mjs tools/native-sum-results.test.mjs
```

The tests execute generated Go against explicit oracles, covering callback and
argument order, calls after failure, first-error priority, untouched inputs,
independent buffers, partial applications, lexical capture, effect timing,
rejected shapes, native fields and dynamic boundaries. Real application
performance comparisons follow the [benchmark protocol](benchmark-protocol.md).
These rules still consume the existing JSON tree; they are not a byte-stream
decoder and do not alter the typed `tcorefn` contract.
