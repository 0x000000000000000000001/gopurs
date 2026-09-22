# Shared construction plans for JSON records

The native Argonaut record decoder constructs the final record once. Its
standard `GDecodeJson` Nil/Cons dictionaries compose an immutable description of
the fields. A single Go loop executes their real PureScript field decoders and
stores the successful values in one fresh buffer, rather than allocating a
larger immutable record at every recursive return.

The public `GDecodeJson` class and method types are unchanged. The implementation
lives in the local `gopurs-argonaut-codecs` package, in
`Data.Argonaut.Decode.Internal.Record`. JavaScript uses the original recursive
implementation through the same internal interface. A custom tail dictionary
without a recognized plan also keeps the original recursive implementation.
Its returned record is never mutated by the optimized path.

## Contract

- Field decoders execute in the original row-list order, stopping on the first
  error. Missing fields, optional fields, null, and `AtKey` errors retain their
  PureScript implementation.
- Symbol callbacks run before their field and again in reverse order after
  all fields succeed, as in the original calls to `Record.insert`. The latter
  names determine the output fields; duplicate names preserve insertion order
  and replacement priority.
- The field plan is immutable. Every decode owns its output buffers; recursive,
  concurrent and repeated invocations cannot mutate earlier results.
- Case analysis of `Maybe`/`Either` remains in PureScript. The FFI does not
  assume generated constructor tags, names, field offsets, or error layouts.
  It invokes the internal success extractor only after the supplied `isRight`
  predicate succeeds; reaching its `Left` branch would be an internal error.
- Only the standard instance constructors create plans. Arbitrary dictionaries
  retain the fallback. No schema-specific decoder body is generated.

## Runtime support

`WithFunctionData(function, metadata)` returns an ordinarily callable `Value`
with a distinct tag and a GC-visible owner holding its function and metadata.
`FunctionData[T]` recognizes only an exact metadata type. Reattaching metadata
replaces the association, rather than adding a wrapper chain. Applying a
function partially returns an ordinary partial function without metadata
intended for the complete call.

All `Apply` variants support these functions. Foreign function classification
and Promise callbacks also support them; no client needs to decode a raw Go
closure pointer. The runtime has no metadata registry or shared mutable cache.
Metadata owners must keep shared data immutable or provide their own
synchronization; the record plan itself is immutable.

## Scope and validation

Payloads and output slots still use `Value`. This optimization removes
intermediate record construction and recursive result wrappers; it is not a
fully typed record/array ABI and establishes no ceiling for later work.

The tests cover runtime arities, partial and excess arguments, allocation-free
saturated dispatch, GC and races; record callback order, first error, custom
fallback, duplicate names, retained records, reentrancy and concurrency; and a
real PureScript fixture with custom `DecodeJson` and `GDecodeJson` instances.
Official performance numbers must come from paired generated-application runs
with the complete JSON and TAST oracles, not the scratch probes.
