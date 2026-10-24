# Native session preload regression

This fixture compiles the current PureScript session controller and FFI, with
the current Go storage helper, through the typed frontend and an explicit
native generator. Private disposable fixtures cover hit, miss, optional decode
failure, forced/allocation builds and a missing recipe. The production controller
also exercises begin failure, cancellation and an immediately completed preload;
all started jobs must finish before use or session closure.

From the project root:

```sh
python3 tools/fixtures/projection-capsule/run-native.py \
  --native-workspace /absolute/path/to/retained-native-workspace \
  --backend /absolute/path/to/gopurs-native \
  --work-dir /absolute/path/to/new-evidence-directory \
  --go /absolute/path/to/go \
  --fixture-dir tools/fixtures/preload-session \
  --pass-marker PRELOAD-SESSION-NATIVE-PASS
```

Executed successfully on 2026-10-24 using candidate J as generator and current
sources. Commands, source hashes and outputs are retained in
`../../.gopurs-evidence/c2-10-sub2-20261023/preload-native-regression`
(relative to the project root). This validates the native callback/controller;
the persistent preparation tests separately verify exact proof hash admission
and rejection after proof changes. It does not establish an end-to-end timing
gain for the newly compiled compiler.

The header extension also compiles the current `HeaderPreload` controller and
codec: workers 1/8, exact owner/inventory checks, late stage/root/context/count
selection, missing/duplicate/malformed parts, and caller error/cancellation or
failed parallel children. All child bracket releases must finish before the
outer session release; an unused optional result is still joined. It prints
`HEADER-PRELOAD-NATIVE-PASS` before the main pass marker. Its reader seam is
explicit and tests verified-read results; storage SHA verification remains the
responsibility of the existing native helper/corruption tests.

This extension passed on 2026-10-24 through candidate K using the current PS
sources. Exact commands, captured sources and five successful native phases are
retained at `../../.gopurs-evidence/c2-10-sub2-20261023/header-preload-native-2`.
The first attempt is retained alongside it and failed on a fixture `Ident` type
annotation before generation. The helper lifecycle and reservoir guards are
tested here; the Driver flags which enable header preload require the final
compiler integration runs and are not claimed as exercised by this fixture.

The two-phase extension passed on 2026-10-24 through candidate K, with the
current PureScript copied and generated afresh. Evidence is retained at
`../../.gopurs-evidence/c2-10-sub2-20261023/header-two-phase-native-1`. It compares
the direct and phased header codecs at jobs 1/8, observes that decoding begins
only at the explicit FFI boundary, and exercises child errors, caller errors
and cancellation in both phases. A failed raw read starts no decoder; both
brackets join before the outer close. The real Driver flag gating remains a
separate integration check on the final candidate.
