# Thunk producer/consumer fusion regression

From the gopurs repository, run:

```sh
GOPURS_TEST_TOOL_DIR="$PWD/../../altbak.pub/run/bak/js/node_modules/.bin" \
  bash tests/thunk-fusion/run
```

Select the same TAST-capable PureScript frontend as `altbak.pub` using
`GOPURS_TEST_TOOL_DIR` (the example assumes the sibling checkout layout).
Without that setting, the runner uses gopurs's `node_modules/.bin`; this must
contain a frontend emitting the required types and constructor declarations.
The runner builds the backend, compiles the fixture, checks the generated Go
structure, and executes its semantic assertions.
It uses only the fixture's dependencies. The general passing-test runner's
larger package set currently also pulls in an unrelated `Foreign.Object`
compilation failure with the typed compiler.

Assertions cover runtime depths and seeds, captured increments, noncommutative
updates, simultaneous loop arguments, collisions after Go name sanitization,
escaping thunks, undemanded failing seeds, and effects that must run only when
executed. Recognized producers are renamed relative to the performance
benchmark. The runner checks that `runAdds`, `runOrder`, `runVary` and `runClash`
call their generated strict workers, directly or through their generated
getters. It does not pin numeric worker suffixes or keep a full Go snapshot.

Separate refusal cases use only Int parameters and total seeds, so an unrelated
type or seed restriction cannot make them pass: a condition based on the loop
counter demands the predecessor conditionally, an overwrite ignores it, and a
third producer forces it twice. The runner verifies that these consumers still
exercise their original producers and that no strict workers were emitted for
them. Disabling fusion therefore fails the positive structural checks, while
weakening the exactly-once demand guard fails the negative checks. The failing
seed and effect timing assertions remain as semantic regression coverage.

The direct IR regression in `check-recursive-seed.mjs` checks another boundary:
an Int annotation alone does not prove that a recursively bound value is already
initialized. The pass conservatively leaves every `LetRec` scope untouched,
including both its bindings and its body. The regression checks that ordinary
function parameters still fuse, while recursive scopes retain delayed reads.
It fails against the earlier pass without this barrier.

`check-semantics.mjs` runs the actual pass against an independent interpreter of
its IR subset: 7,188 differential comparisons across 11 formulas and three thunk
positions, including wrapping int64 boundary values, plus 15 rejection cases.
It also checks bounded executions of nonterminating recursions and the LetRec
barrier. These finite checks supplement the compiled Go fixture; they are not a
formal proof or a test of the frontend and Go emitter. The bounded executions
do not by themselves prove divergence.

The correctness argument is an invariant: at every producer step, the worker's
integer accumulator equals the result of forcing the corresponding accumulated
thunk. Each admitted update preserves the same expression and captured argument
values, substituting the integer for the single unconditional predecessor call.
The seed and moved operations are total. Branches and next control arguments do
not inspect the thunk, so the original producer and worker follow the same
recursion. A finite run therefore returns the same integer; an infinite producer
still never returns. The consumer stays in its original branch or closure, and
the pass does not hoist work out of a surrounding delayed computation.

This is not a conversion of arbitrary Lazy values to eager evaluation. A stream
that returns a constructor before recursing does not match the producer grammar;
neither does a memoized `Data.Lazy` value. Escaping producers retain their original
implementation. The LetRec barrier additionally prevents speculative reads of
recursively initialized values.

The application benchmark remains `./bin/go/run -c` from `altbak.pub`, with
its PureScript source unchanged. Allocation measurements should call the
generated `Call_Test_LazyEvaluation_runManyTimes(1000, 0)` using the same
Go benchmark harness and `GOGC=800` for both revisions, consuming its result.
Measure `ns/op`, `B/op` and `allocs/op`, and check other depths/seeds separately.
