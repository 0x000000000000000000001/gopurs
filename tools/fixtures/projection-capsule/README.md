# Native projection capsule regression

`Main.purs` exercises the current `ProjectionCapsule.create`, Wire codec and
`ProjectionCapsule.open` through the enriched TAST and native Go backend. It
checks covered and new present/absent global types, complete metadata bytes,
capsule round-trip bytes, interfaces and main modules. It also runs the current
`mapPreparationJobs` with 1 and 8 workers: exact capsule bytes and order,
repeated executions, an injected encoding failure, joined bracket releases and
an exact retry. The failure case checks actual maximum concurrency and zero
active jobs before the retry.

Use a workspace retained by `npm run build:native -- --keep-workspace` to supply
the typed compiler and dependency inventory. The runner copies current project
sources and the required dependency closure into a **new** directory; it does
not modify the retained workspace or install a compiler.

From the gopurs project root:

```sh
python3 tools/fixtures/projection-capsule/run-native.py \
  --native-workspace /absolute/path/to/gopurs-native-build-WORKSPACE \
  --backend /absolute/path/to/gopurs-native \
  --work-dir /absolute/path/to/new-projection-capsule-test \
  --go /absolute/path/to/go
```

The runner executes the typed frontend, native support preparation, explicit
native backend, `go build`, and the generated executable. Success prints
`PROJECTION-CAPSULE-NATIVE-PASS` in `run.stdout`; `result.json` records executable
and source hashes, all commands and their exit status. No JavaScript backend is
used to generate the regression executable.

This test was executed on 2026-10-23 with the retained candidate I backend
(`8d4ea8a29d2ee3da19c310affe403c7c24bbcb0bbbf21e1f169cb38eb06f4744`)
and the corrected current source. Evidence is retained at
`../../.gopurs-evidence/c2-10-sub2-20261023/capsule-record-repro/durable-native`
relative to the project root.

The regression specifically uses the eleven-field shared metadata record and
the twelve-field complete record. Native lowering of generic field insertion
into compact records remains a separate known issue; this capsule constructs
both record shapes explicitly. This test does not replace the complete native
Driver checks for capsule carry, missing/corrupt objects and cache fallback.

The shared effectful dispatcher checks were executed on 2026-10-23 using
candidate J (`5ea417545f80e0e15ebb3214ee423a3124cdadb90262b99ec555e01fe2fb486a`)
as the native generator and the current PureScript sources. The exact commands,
source capture and successful native result are retained at
`../../.gopurs-evidence/c2-10-sub2-20261023/priming-native-regression`.
Both failure cases reported `started == finished` and `active=0`, with measured
maximum concurrency equal to the configured 1 or 8 workers.
