# Counted function producers

[`Gopurs.FunctionFusion`](../src/Gopurs/FunctionFusion.purs) runs after PBO and
thunk fusion, before ownership analysis and tail-call optimization. It replaces
a recursive chain of functions with a closure capturing its counter and calling
a loop worker.

## Recognized shape

The producer belongs to a singleton recursive group, has the explicit type
`Int -> (Int -> Int) -> Int -> Int`, and follows this structure:

```purescript
build 0 = identity
build n =
  let previous = build (n - 1)
  in \f x -> f (previous f x)
```

The identity body is checked directly or through a name in the same module.
Grouped arrows and a polymorphic identity are accepted. Incompatible
annotations, other decrements, mutual recursion, opaque base cases and extra
work retain the original implementation. Recognition does not depend on
application or benchmark names.

## Guarantees

For `n >= 0`, the returned function calls a worker that applies `f` exactly
`n` times. Existing tail-call optimization turns the worker into a loop. The
captured count stays immutable and each invocation has its own working counter,
preserving saved functions, reuse and partial applications. Zero invokes no
callback.

The negative branch evaluates the original producer body at production time.
The guard precedes the returned closure, preserving the public arity. Worker
names are reserved against bindings and FFI names after Go normalization.

Counter and accumulator parameters use `int64`; callbacks retain their `Value`
ABI and `Apply` calls. The proof uses typed transformed bodies, independently
of source usage annotations.

## Validation

After building, run:

```sh
node --test tools/counted-functions.test.mjs
bin/test CountedFunctions
```

The tests cover admission and rejection, annotations, the negative branch, name
collisions, callback order, partial applications, reuse, and a failing callback
followed by another invocation. The fixture includes a producer capturing the
counter that must retain the generic path.
