# Plan d'optimisation `gopurs` (phase 2) — Rapprochement `μ/C`, `Q₂/C`, `G/C` vers `1x`

> Plan fondé sur le profiling empirique (`pprof` CPU + `alloc_space`, `go tool objdump` ARM64 et désassemblage `clang -O3 -S` de `altbak.pub/tmp/c-library/arrays.c`) réalisé sur le build figé `go-opt-20261127` (`go1.27.0 darwin/arm64`, `GOGC=800`, `-pgo=off`).

## Baselines officielles de référence (`altbak.pub/README.md`)

Sur les **37 cas Go Release** publiés dans `/Users/0x1/Documents/htdocs/altbak.pub/README.md` (post-campagne `go-opt-20261127`) :
- **Total WIP / C** : Go `~8.57 ms` vs C `~8.78 ms` (**`0.98x`**)
- **Moyenne arithmétique (`μ/C`)** : **`2.60x`**
- **Médiane (`Q₂/C`)** : **`2.46x`**
- **Moyenne géométrique (`G/C`)** : **`1.43x`**

Règles strictes :
- **Généralité** : aucune heuristique nommée par benchmark, aucune constante de workload codée en dur, aucun contournement des barrières d'entrées opaques (`Bench.opaque`).
- **Sémantique PureScript / JS exacte** : Int32 wrapping, arrondi float64 `IntMul` hors plage prouvée, UTF-16/WTF-8 (surrogates isolés et séquences invalides), immuabilité, ordre strict d'évaluation, ABI publique intacte.
- **DX invisible** : aucun flag de build, aucun `GOEXPERIMENT`, aucun `cgo`, cross-compilation standard Go (`GOOS`/`GOARCH`) intégralement préservée via build tags et fallback Go pur.
- **Empirisme** : chaque étape est validée par tests unitaires, inspection du Go généré + `objdump`/BCE, `pprof` isolé et comparaison aux baselines officielles `altbak.pub/README.md`.

---

## Lot 1 — Overhead d'appel `Effect` / FFI & micro-coûts `2–6.6 ns` (16 cas sur 37)

Objectif : attaquer le bloc massif des 16 cas ultra-courts (`2.01x` à `5.22x` : `record/set-*`, `st/ref-*`, `effect/ref-*`, `string/cu-length`, `traverse/fail-first`, `RowToList`, `Polymorphism`, `StateMonad`, `LazyEvaluation`, `Church`, `ListOps`, `ArrayOps`) où **60 % à 85 %** du temps CPU est consommé par la plomberie `Effect`/FFI.

- [x] **1.1 — Élimination du trampoline `func() int` dans `FfiBridge/Values.purs` & `EffectExprs.purs` (`6.10%` flat CPU)**
  - Preuve `pprof`/`objdump` : `Bench_MeasureBatch_runEffectWorker.func1.func1` pèse `6.10%` flat ; `EffectExprs.wrapIntAction` attache `func() int64` dans `WithFunctionData`, mais tout FFI Go prenant `func() int` se voit générer un wrapper `func() int { return int(__gopurs_native_action()) }` (1 frame + 1 `CMP RSP` + 1 branche indirecte supplémentaire à chaque appel).
  - Action : permettre à `runtime.go` (`functionData`) / `EffectExprs.purs` / `FfiBridge/Values.purs` d'exposer directement l'entrée `func() int` native sans closure trampoline (en préservant l'entrée `func() int64` et l'ABI boxée).
  - Validation : `native-effect-callback.test.mjs`, `runtime-contracts` (`-race`), `objdump` confirmant la disparition de la closure trampoline. *(Fait via `gopurs_runtime.IntActionFromInt64` + test `-race -d=checkptr=2`.)*

- [x] **1.2 — Liaison native des actions FFI `Effect` hissées par `Let` (`7.67%` flat `r_.Apply` + `3.14%` `init.func4.1`)**
  - Preuve `pprof` (`-peek=Apply`) : `r_.Apply` est le **#1 flat CPU** de la suite, directement appelé à chaque itération de `EffectRefWrite/Modify`, `STRefWrite/Modify`, `RecordSetSmall/Medium/Large`, `TraverseMaybeFailFirst`, `StringCodeUnitsLength`, `ArrayAppend`, etc. Quand `let act = ffi arg` est hissé hors du bloc `Effect` interne (`pbo` `floatLet` ou code utilisateur), `ImmediateFfiEffects` ne le reconnaît plus et génère `r_.Apply(Get_*, arg)` → `r_.Func` → `r_.Apply(act, r_.Value{})` + `r_.Box`.
  - Action : dans `BindingExprs.purs`, `EffectExprs.purs` et `ImmediateFfiEffects.purs`, détecter un `Let` liant une application saturée d'une fabrique FFI `Effect` éligible (`func(A...) func() R`), conserver son thunk natif `func() R` dans l'environnement local (sans casser les usages de première classe si la variable échappe) et l'invoquer directement dans `EffectBind` sans passer par `r_.Apply` ni `r_.Func`.
  - Validation : nouveaux tests unitaires (action `Effect` liée en `Let`, réexécutée plusieurs fois, échappée/capturée, ordre strict des arguments), `ffi-immediate-effects.test.mjs`, `pprof -peek=Apply` à `~0%` dans les boucles chaudes.

- [x] **1.3 — Inlining de `*_runEffectWorker` & chemin rapide `r_.Box` (`3.31%` + `2.96%` flat CPU)**
  - Preuve `objdump` : `Bench_Opaque_runEffectWorker` inline à la fois `Bench_Opaque` (`CALL opaqueValue`, coût 57) et `r_.Box` (`CALL boxSlow`, coût 57) ⇒ coût > 114 > budget Go (80) ⇒ le worker n'est jamais inliné dans `RowToList`, `Church`, `StateMonad`, `LazyEvaluation`, `Polymorphism`, `ListOps`, `ArrayOps`, `RecordSet*`.
  - Action : dans `ImmediateFfiEffects.purs` / `FfiBridge/ImmediateEffects.purs`, réduire le coût d'inlining des appels FFI `Effect` immédiats (émission directe au site d'appel ou séparation du chemin lent de `Box`) pour que le compilateur Go appelle la cible FFI sans frame intermédiaire.
  - Validation : `objdump` de `Get_Test_RowToList_act.func1.2` confirmant l'appel direct sans frame `*_runEffectWorker` ; `ffi-immediate-effects.test.mjs` OK.

- [x] **1.4 — Nettoyages ciblés sur les micro-cas core (`StateMonad`, `ArrayOps`, `Polymorphism`, `Church`)**
  - Preuve `pprof`/Go généré :
    - `StateMonad` : `_ = Get_Data_Unit_unit()` exécute un `sync.Once` mort à chaque appel. *(Vérifié par `objdump` : Go DCE déjà cet appel — 12 instrs sans frame.)*
    - `ArrayOps` (`0.11s` cum) : `ArrayRangeFold.purs` capture `Get_Data_Semiring_intAdd()` (`sync.Once`) et la closure prédicat même dans le fast-path fermé (`constantSum`). *(Outlining du `slow` fallback pour réduire le cadre de pile.)*
    - `Polymorphism` (`0.09s` cum) : les workers locaux récursifs sans capture (`Call_local_*`) sont déclarés/assignés dans le corps chaud de l'action `Effect` au lieu d'être hissés en fonctions top-level. *(Élagage des pairs récursifs morts + `var f func(...) = func(...)` → dévirtualisé & inliné à 100 % par Go !)*
    - `Church` : 6 gardes redondantes `x >= 0` et 4 multiplications découpées 16 bits dans `CountedFunctionSums.purs` alors qu'il s'agit de multiplications exactes mod $2^{32}$ (`int64(int32(a * b))`). *(Gardes dédupliquées + `PrimitiveExprs.modularIntProduct` → 23 `NOOP` supprimés + 4 `MUL`/`SXTW` !)*
  - Action : éliminer ces évaluations mortes (`ArrayRangeFold`, `Unit` pur), hisser les workers locaux sans capture (ou éviter leur réallocation/réinitialisation), et simplifier les multiplications exactes mod $2^{32}$ et gardes dupliquées dans `CountedFunctionSums.purs`.
  - Validation : tests unitaires associés (`array-range-fold`, `counted-function-sums`, `local-workers`, `closed-scalar`) OK.

- [x] **1.5 — Validation & mesure Lot 1**
  - Suite de tests `gopurs` + mesure pairée `AppAll` 37/37 cas (0 divergence de sortie). *(Pairé 3 runs : `μ/C` `2.49x → 1.88x`, `Q₂/C` `2.34x → 1.71x`, `G/C` `1.37x → 1.08x` !)*

---

## Lot 2 — Tableaux compacts `[]int32`, zéro-copie pure, boucles sans `NOOP`/BCE & 2 noyaux SIMD `.s` invisibles (6 cas `2.13x`–`6.96x`)

Objectif : attaquer les 6 cas tableaux (`array/append-1000` `6.96x`, `traverse/maybe-1000` `4.46x`, `array/map-10000` `4.18x`, `array/sort-descending-10000` `2.99x`, `array/foldl-sum-10000` `2.74x`, `array/eq-neq-last-10000` `2.13x`) pour les aligner sur les temps C (`clang -O3` NEON).

- [x] **2.1 — Emprunt zéro-copie `BorrowIntArray` pour les appels PureScript purs + fusion `SortPlan` (`array/sort`)**
  - Fait : `CallArguments.coercePrefixBorrow` + `gopurs_runtime.BorrowIntArray` pour les workers purs `Call_*` (`UnboxIntArray` conservé aux frontières FFI) ; `SortPlan` ajouté dans `ArrayProducerFold.purs` (`sort`/`sortBy`/`sortBy__`/`sortByImpl` → 1 seul `[]int32` 40 Ko sans `Rebox(ordInt)` ni double copie, `array/sort-descending-10000` `15.17 μs → 6.84 μs`, `1.47x`).

- [x] **2.2 — Représentation compacte `TypeInt32Array` (`[]int32`, 4 octets/élément) avec fallback `TypeIntArray` (`[]int64`)**
  - Fait : `TypeInt32Array = 29` dans `runtime.go` avec compaction automatique dans `IntArray([]int64)` (fallback `TypeIntArray` hors `int32`), `EqualInt32Array` (`bytes.Equal` 40 Ko, `array/eq-neq-last-10000` `1.34 μs → 0.49 μs`, `0.78x`), support complet dans `ArrayEquality`, `ScalarArrayFold`, `ArrayIndexedFold`, `ArrayProducerFold`, `ArrayTraverseFold`, `ArraySort`, `gopurs-st` et `gopurs-node-buffer`.

- [x] **2.3 — Correction BCE, suppression des `NOOP` / gardes `2^52` et réduction `int32` unique**
  - Fait : `ScalarArrayFold.purs` (`index <= len-8`), `ArrayIndexedFold.purs` (`TypeInt32Array`/`TypeIntArray` sans copie `ArrayValues`), `emitAffineStep32` en `int32` brut sans garde `2^52` ni closure dans `FilterPlan` (`array/filter-10000` `5.03 μs → 3.14 μs`, `0.69x`) et `ArrayTraverseFold.emitFallibleLoop`.

- [x] **2.4 — Deux noyaux SIMD `.s` génériques et 100 % invisibles dans `gopurs_runtime` (`arm64` NEON + `amd64` SSE2 + fallback Go pur)**
  - Fait : `runtime/{simd_asm.go,simd_generic.go,simd_arm64.s,simd_amd64.s}` (noyau `sumEvenOddInt32Asm` 16 `int32`/itér. sans multiplication dans la boucle, 0 flag, 0 `cgo`, auto-émis via `Gopurs.Runtime.Simd` et `Driver/Output.purs`, intégré dans `tools/build-cache/`).
  - Reconnaissance algébrique générale `IndexedAddLanes.extractAffineParity` prouvée exacte sur $\mathbb{Z}/2^{32}\mathbb{Z}$ (`|c| <= 2^21`).
  - Validé : `tools/simd-kernels.test.mjs` (`go vet`, `-race`, `arm64`, `amd64`, `riscv64`, `js/wasm`), `tools/build-cache` `go test -race`, `driver-cache.test.mjs` et `array-producer-fold-e2e.test.mjs`.

- [x] **2.5 — Validation & mesure Lot 2**
  - Mesure pairée `AppAll` 37/37 identiques (3 runs entrelacés, `GOGC=800`) vs `go-opt-20261127` : **`μ/C = 1.50x`** (vs `2.64x`), **`Q₂/C = 1.47x`** (vs `2.47x`), **`G/C = 0.90x`** (vs `1.47x`, sous `1.0x` !). Gains : `array/map` `0.73x` (5.75×), `array/filter` `0.69x` (1.60×), `array/foldl-sum` `1.13x` (2.44×), `array/sort-descending` `1.47x` (2.22×), `array/eq-neq-last` `0.78x` (2.71×), `array/append` `0.89x` (7.90×), `traverse-maybe` `0.80x` (5.60×).

---

## Lot 3 — Chaînes Unicode & recherche de sous-chaîne (4 cas `1.36x`–`2.57x`)

Objectif : amener `string/cp-length` (`2.57x`), `string/cp-take` (`2.42x`), `string/indexOf-miss` (`1.96x`) et `string/concat` (`1.36x`) autour ou sous `1.0x`.

- [x] **3.1 — `gopurs-strings` `CodePoints.go` : BCE + comptage SWAR 8 octets (`string/cp-length` `2.57x`, `2.82 μs`)**
  - Fait : validateur/compteur SWAR 8 octets en blocs de 240 octets dans `gopurs-strings/src/Data/String/CodePoints.go` (popcount entier `(contSum*0x0101...)>>56` sans franchissement GPR↔FPR, court-circuit 3-octets quand `lead34 == lead4`, rembobinage exact à la frontière de caractère vers fallback BCE `i <= n-4` WTF-8). Validé par 12 000+ oracles UTF-8/WTF-8/surrogates/malformés (`tools/strings-kernels.test.mjs`, `-race`) ; `string/cp-length-unicode-1000` `2.61 μs → 1.65 μs` (`1.51x`, 1.58× plus rapide, Go 100 % pur).

- [x] **3.2 — `CodePointFold.purs` : boucle BCE `len(rem) >= 4` + accumulation additive `int64` (`string/cp-take` `2.42x`, `1.56 μs`)**
  - Fait : tranche décroissante `codePointFold_rem` (`len(rem) >= 4` → 0 `panicBounds`/`panicSlice` dans `objdump`) + détection `directAdd` sur l'AST Go traduit + accumulation `int64` brute avec réduction `int32` unique (préservation exacte des seeds larges hors `Int32` sur entrée vide). Validé : `codepoint-fold.test.mjs`, `codepoint-take-fold.test.mjs` (11 235 cas) ; `string/cp-take-unicode-1500` `1.44 μs → 1.12–1.22 μs` (`1.73–1.89x`).

- [x] **3.3 — `gopurs-strings` `CodeUnits.go` : `isBMPText` sur octets + recherche anti-pathologie (`string/indexOf-miss` `1.96x`, `2.42 μs`)**
  - Fait : `isBMPText` sur octets purs (`b < 0xF0 && b != 0xED`) + `indexBytesAdaptive` (pivote de `x[0]` vers `x[m-1]` via `strings.IndexByte` NEON après 4 faux positifs, puis fallback `strings.Index` après 8 échecs d'ancre pour garantie $O(n)$ pire cas). Validé : `tools/strings-kernels.test.mjs` ; `string/indexOf-miss-10000` `2.35 μs → 0.20 μs` (`0.16x`, **11.63× plus rapide**) et `indexOf-hit` intact (`0.28 μs`, `0.26x`).

- [x] **3.4 — `StringConcatFold.purs` : accumulation `int64` brute + `concatFold_finish` sans capture mutable (`string/concat` `1.36x`, `1.26 μs`)**
  - Fait : passage par valeur de `concatFold_acc` à `concatFold_finish` (zéro spill pile), suppression du suivi `concatFold_any` lorsque le seed constant est déjà dans `Int32`, plus inlining restauré d'`ArrayValues` et simplification de `StrValue` dans `runtime/runtime.go`. Validé : `string-concat-fold.test.mjs` (13 254 cas), `runtime-contracts.test.mjs` (`-race`) ; `string/concat-1000` `1.23 μs → 1.18 μs` (`1.28x`).

- [x] **3.5 — Validation & mesure Lot 3**
  - Suite `strings-kernels`, `codepoint-fold`, `codepoint-take-fold`, `string-concat-fold`, `runtime-contracts` 100 % verte + mesure pairée 37/37 identiques.

---

## Lot 4 — `TailAddLoops.purs` (`Fibonacci` `1.65x`, `0.13 μs`) & campagne officielle `altbak.pub`

Objectif : éliminer les derniers overheads sur la récursion additive (`Fib`) puis valider le compilateur natif (`bin/gopurs-native`), la compilation multi-packages et lancer la campagne officielle `altbak.pub`.

- [x] **4.1 — Simplification & spécialisation de `TailAddLoops.purs` (`Fib` `1.65x`, `0.13 μs` vs `0.079 μs`)**
  - Fait : accumulation `acc` en `int64` brut avec réduction `IntAdd` unique au `Finish`, et dépliage d'un cran avec gardes de cas de base purs (`extractBaseCases` + `extractPureStep`) sur les récursions binaires additives pures (réduit `fib(10)` de 88 à 33 appels tout en préservant l'ordre gauche-droite, le wrapping Int32 et les valeurs de base `Int64` larges). Validé : nouveau test oracle binaire dans `tools/tail-add-loops.test.mjs` (7/7) ; `Fibonacci` `0.129 μs → 0.092–0.102 μs` (`1.16–1.29x`, 1.26–1.38× plus rapide).

- [x] **4.2 — Reconstruction `bin/gopurs-native`, tests complets & campagne officielle `altbak.pub`**
  - Reconstruit `bin/gopurs.js` + `bin/gopurs-native` (`npm run build && npm run build:native`, 568 modules, 356 797 types, 0 erreur / 0 warning).
  - Suites validées : `test:ffi` (39/39), `test:cache` (6/6 dont `-race`), `test:cache-driver` (`js`/`native-sequential`/`native-parallel`), `test:cache-sources` (`js`/`native-sequential`/`native-parallel`), `test:cache-context` (tous hôtes), `test:cache-codecs` (15/15), `test:identity` (8/8), 37/37 tests ciblés (dont `simd-kernels` arm64/amd64/generic/`-race`, `strings-kernels` 12 000+ oracles, `runtime-contracts` `-race`), `tools/build-cache` `go test -race ./...`.
  - Campagne officielle `altbak.pub` (`driver.py` avec `bin/gopurs-native` par défaut sans flag + `measure.py`, 3 processus indépendants, `GOGC=800`, `-pgo=off`, 37/37 sorties validées, 0 régression sur 37 cas) :
    - **37 cas (`go-pure` 14 + `go-pure-all.library` 23)** : Total `8.574 ms (0.98x) → 8.359 ms (0.95x)`, **`μ/C` `2.60x → 1.34x`**, **`Q₂/C` `2.46x → 1.30x`**, **`G/C` `1.43x → 0.79x`**.
    - **14 cas core (`go-pure`)** : Total `8.539 ms (0.97x) → 8.341 ms (0.95x)`, **`μ/C` `1.61x → 1.06x`**, **`Q₂/C` `1.83x → 1.19x`**, **`G/C` `0.53x → 0.38x`**.
    - **23 cas bibliothèque WIP** : Total `0.0343 ms (1.86x) → 0.0175 ms (0.95x)`, **`μ/C` `3.20x → 1.50x`**, **`Q₂/C` `3.65x → 1.39x`**, **`G/C` `2.63x → 1.22x`**.

---

## Journal de mesures

| Étape | Total WIP / C (C `8.78 ms`) | `μ/C` | `Q₂/C` | `G/C` | Notes |
| --- | --- | --- | --- | --- | --- |
| Baseline officielle `go-opt-20261127` | `8.57 ms` (`0.98x`) | `2.60x` | `2.46x` | `1.43x` | Référence `altbak.pub/README.md` |
| Lot 1 (Effect/FFI + micro-cas core) | `8.31 ms` (`0.95x`) | `1.88x` | `1.71x` | `1.08x` | Pairé vs Old `2.49x`/`2.34x`/`1.37x` (37/37 identiques) |
| Lot 2 (`[]int32`, Borrow, BCE/NOOP, SIMD `.s`) | `9.04 ms` (`1.03x`) | `1.50x` | `1.47x` | `0.90x` | Pairé vs Old `2.64x`/`2.47x`/`1.47x` (`G/C < 1x` ! 37/37 identiques) |
| Lot 3 + 4.1 (SWAR UTF-8, `cp-take`, `indexOf`, `concat`, `TailAddLoops`) | `8.63 ms` (`0.98x`) | `1.33x` | `1.29x` | `0.78x` | Pairé vs Old `2.55x`/`2.40x`/`1.41x` (37/37 identiques, 0 régression) |
| **Lot 4.2 — Campagne officielle `measure.py` (3 processus)** | **`8.36 ms` (`0.95x`)** | **`1.34x`** | **`1.30x`** | **`0.79x`** | **`bin/gopurs-native` par défaut, 0 flag, 37/37 validés, 0 régression** |
