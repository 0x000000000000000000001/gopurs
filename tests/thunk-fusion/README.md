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

The application benchmark remains `./bin/go/run -c` from `altbak.pub`, with
its PureScript source unchanged. Allocation measurements should call the
generated `Call_Test_LazyEvaluation_runManyTimes(1000, 0)` using the same
Go benchmark harness and `GOGC=800` for both revisions, consuming its result.
Measure `ns/op`, `B/op` and `allocs/op`, and check other depths/seeds separately.
