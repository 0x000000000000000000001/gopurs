# Plan d'optimisation `gopurs` (phase 3) — Passer `μ/C` et `Q₂/C` sous `1x` et pousser au plafond matériel

> Plan fondé sur le profiling empirique (`pprof` CPU + `alloc_space`, `go tool objdump` ARM64) réalisé sur `HEAD c9f4b1a` (`AppAll` 37/37 cas, `GOGC=800`, `-pgo=off`, Apple M4 Pro).

## Baselines officielles de référence (`altbak.pub/README.md`)

Sur les **37 cas Go Release** publiés dans `/Users/0x1/Documents/htdocs/altbak.pub/README.md` (post-phase 2) :
- **Total WIP / C** : Go `~8.36 ms` vs C `~8.78 ms` (**`0.95x`**)
- **Moyenne arithmétique (`μ/C`)** : **`1.34x`**
- **Médiane (`Q₂/C`)** : **`1.30x`**
- **Moyenne géométrique (`G/C`)** : **`0.79x`**
- **Répartition** : 12 cas ≤ `1.0x`, 25 cas > `1.0x` (16 micro-cas `1.2–3.8 ns` jusqu’à `3.08x`, 9 kernels réels `0.10–16 μs` jusqu’à `1.87x`).

Règles strictes :
- **Généralité** : aucune heuristique nommée par benchmark, aucune constante de workload codée en dur, aucun contournement de `Bench.opaque`.
- **Sémantique PureScript / JS exacte** : Int32 wrapping, arrondi float64 `IntMul` hors plage prouvée, UTF-16/WTF-8 (surrogates isolés et séquences malformées), immuabilité, ordre strict d’évaluation, ABI publique intacte.
- **DX invisible** : aucun flag de build Go, aucun `GOEXPERIMENT`, aucun `cgo`, cross-compilation standard (`GOOS`/`GOARCH`) préservée via build tags et fallback Go pur.
- **Empirisme** : chaque étape est validée par tests unitaires, inspection du Go généré + `objdump`, mesure pairée `AppAll` 37/37 (0 divergence) et comparaison aux baselines officielles `altbak.pub/README.md`.

---

## Lot 1 — Plomberie FFI, inlining des fast-paths & micro-cas (`16 cas`, `3.08x → ~1.0–1.3x`)

Objectif : éliminer l’appel indirect `Bench_Opaque.1` (`9.69%` cum CPU), faire inliner les workers fermés (`StateMonad`, `LazyEvaluation`, `ListOps`, `ArrayOps`), supprimer les frames `128–500 B` sur les micro-cas et corriger `cu-length` (`2.64x`) et `ref-modify` (`1.96x`).

- [x] **1.1 — Descente des fabriques FFI `Effect` sans état (`tools/ffi-gen/parser.go` + `ImmediateFfiEffects.purs`)**
  - Preuve `pprof`/`objdump` : `Bench_Opaque.1` pèse `3.94%` flat / `9.69%` cum CPU sur 11 benchmarks `prepare` (jusqu’à 50 % du temps sur `EffectRefWrite`, `TraverseMaybeFailFirst`, `RecordSetSmall`). PBO (`floatLet`) hisse `act := Bench_Opaque(arg)` hors de la closure interne, transformant un appel direct inlinable en appel indirect `CALL (R0)` + frame `32 B`.
  - Action : détecter dans `tools/ffi-gen/parser.go` (`StatelessEffectFactory bool`) toute fonction Go FFI `func F(params...) func() R` dont le corps est un unique `return func() R { ... }` (aucune instruction ni allocation d’état avant `return`). Dans `ImmediateFfiEffects.bindThunk`, capturer les arguments évalués au site du `Let` (ordre strict conservé) et émettre `F(args...)()` directement au site `EffectBind` pour que Go inline `F`.
  - Validation : tests unitaires Go `tools/ffi-gen` (`go test`), `npm run build:ffi`, nouveaux cas dans `tools/ffi-immediate-effects.test.mjs` (stateless vs stateful avec état capturé avant `return`, ordre strict), `objdump` confirmant la disparition de `Bench_Opaque.1`.

- [x] **1.2 — Outlining systématique des fallbacks froids (`AffineLoops`, `ArrayRangeFold`, `ListFusion`, `RecordConsumer`, `RefLoopRegions`, `CountedFunctionSums`, `ArrayTraverseFold`)**
  - Preuve `pprof`/`objdump` : `Call_Test_StateMonad_runManyTimes` (`0.82%` flat), `Call_Test_LazyEvaluation_runManyTimes`, `Call_Test_ListOps_sumEvens`, `Call_Test_ArrayOps_sumEvens` ne sont pas inlinés (boucle lente inline dans le même corps → coût > 80) ; `RecordSetMedium` a une frame de `320 B`, `RecordSetSmall` et `STRefWrite` `128 B`, `Church` `96 B`.
  - Action : extraire le fallback froid dans un helper `*_slow` / fonction hors chemin chaud afin que les wrappers rapides (`cost ~12`) s’inlinent dans les appelants et que les actions chaudes gardent une frame minimale (`32 B`).
  - Validation : `objdump` confirmant l’inlining de `StateMonad`/`LazyEvaluation`/`ListOps`/`ArrayOps` et la réduction de frame sur `RecordSet*`/`STRefWrite`/`Church` ; tests `affine-loops`, `array-range-fold`, `record-affine-consumer`, `counted-function-sums`, `ref-loops` verts.

- [x] **1.3 — `StringCodeUnitsLength` inlinable en 3 instructions & suppression `NormalizeStringValue` redondant (`2.64x`)**
  - Preuve `objdump` : `Get_Test_StringCodeUnitsLengthUnicode1000_prepare` émet `NormalizeStringValue(__local_var_2_3)` puis un `CALL gopurs_runtime.StringCodeUnitsLength(SB)` non inliné.
  - Action : séparer `StringCodeUnitsLength` dans `runtime/runtime.go` en fast-path minimal (inliné : `Type == TypeString`, `IntVal == 0`, `IntVal < 0`) + `stringCodeUnitsLengthSlow`, et éviter le wrapping `NormalizeStringValue` superflu quand l’argument est déjà un `TypeValue` passé à un FFI prenant `Value`.
  - Validation : `runtime-contracts` (`-race`), `string_metadata_test.go`, `objdump` sans `CALL StringCodeUnitsLength`.

- [x] **1.4 — Identités algébriques Int32 (`x * 1`, `0 + x`, puissance mod $2^{32}$) dans `RefLoopRegions`, `PrimitiveExprs`, `CountedFunctionSums`**
  - Preuve `objdump` : `STRefModify` et `EffectRefModify` (`1.96x`) exécutent 7 instructions (`uint64(n + 1<<53) <= 2<<53` + `math.mod` fallback) pour `0 + n * 1` ; `Church` (`1.54x`) exécute 4 `MUL+SXTW` sériels.
  - Action : simplifier `OpAdd (LitInt 0) x`, `OpMultiply x (LitInt 1)` et `OpMultiply x (LitInt 0)` dans `RefLoopRegions` / `PrimitiveExprs` (avec préservation exacte du narrowing Int32), et réduire la chaîne de produits dans `CountedFunctionSums` (`int32` sans `SXTW` intermédiaires + arbre d’exponentiation).
  - Validation : tests unitaires (y compris entrées larges `int64` hors `Int32` et négatives), `objdump` confirmant la disparition de `math.mod`.

- [x] **1.5 — Gardes & lectures records rapides (`record/set-small/medium/large`, `2.80x–3.08x`)**
  - Preuve `objdump` : `RecordShape2` compare 2 strings (`K0 == "a" && K1 == "b"`) au lieu d’un pointeur partagé ; `RecordData.Vals` (stride 24 B) étale 16 entiers sur 384 B (6 cache lines vs 64 B en C).
  - Action : partager/interner les clés `RecordKeys2..5` pour test pointeur rapide dans `RecordShapeGuard.purs` et ajouter un fast-path compact pour les valeurs entières de `RecordData` (ou stockage contigu `IntVals []int32`/`[]int64`) tout en gardant la compatibilité totale avec `Vals []Value`.
  - Validation : `record_shapes_test.go`, `compact_int_records_test.go`, `record-affine-consumer.test.mjs`, `record-affine-consumer-e2e.test.mjs`.

- [x] **1.6 — Validation & mesure pairée Lot 1**

---

## Lot 2 — Tableaux : fusion zéro-alloc du tri inversé, `traverse/fail-last` & `SIMD` 128 B

Objectif : éliminer `954.85 MB` (`81.5%`) d’allocations heap dans `array/sort-descending-10000` (`1.39x → < 0.5x`), et faire passer `traverse/fail-last` (`1.08x`) et `array/foldl-sum` (`1.07x`) largement sous `1.0x`.

- [x] **2.1 — Fusion zéro-allocation du `SortPlan` décroissant (`ArrayProducerFold.purs`) + détection monotone rapide**
  - Preuve `pprof` (`alloc_space`) : `Get_Test_ArraySortDescending10000_prepare.func1.1.2.1` alloue `954.85 MB` (`81.52%` du total) pour `make([]int32, 10000)` + copie inversée avant `AffineParitySumInt32`.
  - Action : dans `ArrayProducerFold.purs`, lorsque `monoName == true` dans la branche strictement décroissante, consommer directement la source sans allouer ni inverser : pour `AffineParitySumInt32`, inverser les poids pairs/impairs si `len(src) & 1 == 0` ; pour un consommateur générique, parcourir `src` à l’envers. Déplier/vectoriser la boucle de vérification monotone (`isMonotoneInt32` 4x unroll ou noyau SIMD `.s` générique dans `gopurs_runtime`).
  - Validation : `array-producer-fold.test.mjs`, `array-producer-fold-e2e.test.mjs` (longueurs paires, impaires, 0, 1, 2, strictement décroissant, doublons, non-monotone), `pprof -alloc_space` confirmant `0 B` alloués sur entrée décroissante.

- [x] **2.2 — `ArrayTraverseFold.purs` déplié par 2 + comparaison `int32` directe (`traverse/fail-last`, `1.08x`)**
  - Preuve `objdump` : branche `if index & 1 == 0` à chaque élément + élargissement `int64` + `r_.IntAdd`.
  - Action : hisser le comparateur `int32` hors de la boucle (avec court-circuit si hors plage `int32`) et dérouler par paires `(i, i+1)` sans branche de parité.
  - Validation : `array-traverse-fold.test.mjs` (longueurs paires/impaires, échec premier/milieu/dernier, débordement Int32).

- [x] **2.3 — Dépliage 32 `int32` (128 B) de `sumEvenOddInt32Asm` & inlining `SumInt32` (`array/foldl-sum`, `1.07x`)**
  - Preuve `objdump` : `simd_arm64.s` boucle par 16 `int32` (64 B) avec `SUB; CMP; BGE` ; `SumInt32` n’est pas inliné.
  - Action : ajouter `loop32` (128 B = 1 cache line Apple Silicon, `SUBS`) dans `runtime/simd_arm64.s` (et `simd_amd64.s` si pertinent), et réduire le wrapper `SumInt32` pour qu’il s’inline ou appelle directement le noyau.
  - Validation : `tools/simd-kernels.test.mjs` (`arm64`, `amd64`, `riscv64`, `js/wasm`, `-race`).

- [x] **2.4 — Validation & mesure pairée Lot 2**

---

## Lot 3 — Chaînes : `TypeStringArray`, noyau SIMD UTF-8 & `SumCodePoints` sans safe-slice

Objectif : passer `string/cp-take-unicode-1500` (`1.87x`), `string/cp-length-unicode-1000` (`1.53x`) et `string/concat-1000` (`1.30x`) largement sous `1.0x`.

- [x] **3.1 — Représentation compacte `TypeStringArray` (`[]string`) dans `runtime.go` & `StringConcatFold.purs` (`1.30x`)**
  - Preuve `objdump` : `StringConcat1000` parcourt `[]r_.Value` et exécute 14 instructions de `StrVal()` par élément × 1 000.
  - Action : ajouter `TypeStringArray = 30` (`[]string`, 16 B/él.) dans `runtime/runtime.go` avec support complet (`ArrayValues`, `Equal`, `StringConcatFold`, `Data.Array.replicate`, `Data.String.split`, `ForeignData`, etc.) + boucle ASCII optimisée dans `StringConcatFold.purs`.
  - Validation : `string-concat-fold.test.mjs`, `string-split-fold.test.mjs`, `runtime-contracts` (`-race`), `objdump` confirmant 0 appel `StrVal()` dans la boucle chaude.

- [x] **3.2 — Noyau SIMD `.s` `countUTF8CodePointsAsm` (NEON `arm64` 16 B/itér. + SSE2 `amd64`, fallback Go SWAR + WTF-8) (`1.53x`)**
  - Preuve `objdump` : `Data_String_CodePoints_Length` rematérialise 5 constantes 64-bit à chaque mot de 8 octets et sérialise sur `carry`.
  - Action : ajouter `CountUTF8CodePoints` dans `gopurs_runtime` (`runtime/simd_arm64.s`, `runtime/simd_amd64.s`, `runtime/simd_asm.go`, `runtime/simd_generic.go`, validation RFC 3629 exacte 16 B/itér. avec sortie propre vers le fallback WTF-8 sur `0xED` ou octet invalide) et l’appeler depuis `gopurs-strings/src/Data/String/CodePoints.go`.
  - Validation : 12 000+ oracles UTF-8/WTF-8/surrogates/malformés (`tools/strings-kernels.test.mjs`, `tools/simd-kernels.test.mjs`, cross-compile `arm64`/`amd64`/`riscv64`/`js/wasm`, `-race`).

- [x] **3.3 — Noyau `SumCodePoints` (sans safe-slice `rem[k:]`) appelé par `CodePointFold.purs` (`1.87x`)**
  - Preuve `objdump` : `CodePointFold` émet 5 instructions de safe-slice sur chaque `rem = rem[k:]`, 4 `MOVBU` séparés et 3 branches chaînées par code point.
  - Action : ajouter `SumCodePoints(s string, limit int64) (sum int64, taken bool)` dans `gopurs_runtime` (`arm64` `.s` / `amd64` `.s` + fallback Go pur indexé par pointeur/index sans re-slicing, fallback `DecodeStringCodePoint` sur `0xED`/malformé) et l’utiliser dans `CodePointFold.purs` quand `directAdd == true` (bounded et unbounded).
  - Validation : `codepoint-fold.test.mjs`, `codepoint-take-fold.test.mjs` (11 235 cas), `simd-kernels.test.mjs` (`-race` + cross-compile).

- [x] **3.4 — Validation & mesure pairée Lot 3**

---

## Lot 4 — Récursion numérique (`Fibonacci` `1.28x`, `Ackermann` `1.06x`, `Primes` `1.04x`)

- [x] **4.1 — `TailAddLoops.purs` (`Fib` `1.28x`) : abaissement des récurrences linéaires homogènes d’ordre 2 + serrage `int32`**
  - Preuve `objdump` : `Call_Test_Fib_fib` paie une frame `64 B` (5 spills), 4 `r.IntSub` (`NOOP`+`SXTW`) et double test `x == 0`/`x == 1`.
  - Action : reconnaître généralement les récurrences linéaires additives pures d’ordre 2 (`f(0)=c0, f(1)=c1, f(n) = f(n-1) + f(n-2)` pour `0 <= n <= MaxInt32`, avec fallback exact pour `n < 0`) en boucle itérative `int32` $O(n)$, et resserrer le chemin général `TailAddLoops` (`int32` sans `SXTW` inutile).
  - Validation : `tail-add-loops.test.mjs` (cas de base variés, `n = 0..46`, wrapping Int32 pour grands `n`, récurrences générales non-Fib).

- [x] **4.2 — `Ackermann` (`1.06x`) & `Primes` (`1.04x`)**
  - Preuve `objdump` : `Call_Test_Ackermann_ackermann` paie 3 `SXTW` + 2 `NOOP` dans sa boucle ; `Call_Test_Primes___gopurs_list_progression_fold_0` a une frame de `608 B` et une boucle de somme à 8 branches.
  - Action : activer un worker interne `int32` pour les fonctions récursives `Int -> Int` à wrapping prouvé (élimine `SXTW`/`NOOP` sur `Ackermann` et `Fib` général), et déporter le fallback `lower != 2` + simplifier la somme de `ListProgressionFold.purs` (`Primes`).
  - Validation : tests `int32-loops`, `list-progression-fold`, `AppAll` 37/37.

- [x] **4.3 — Validation & mesure pairée Lot 4**

---

## Lot 5 — Rounds 2–5 de `pprof` profond, dépassement du plan (`RBTree`, `AST`, `split`, `filter`, `cp-take`, `append`, `record/set`), campagne officielle & graduation `Stable` / `WIP`

- [x] **5.1 — Re-profiling complet `pprof` + `objdump` post-Lots 1–4 & Rounds 2–5 d’optimisations générales**
  - **Round 2 — `RBTree` & `string/split` & `string/indexOf` & `array/filter`** :
    - `Ownership/BalancedLoop.purs` : réutilisation en place des nœuds sur les 4 rotations RBTree + allocation par slab 512 nœuds + worker `buildTree_fast` (`8.32 ms → 1.91 ms`).
    - `StringSplitFold.purs` : raccourci 0-alloc pour séparateur ASCII 1 octet (`strings.Count` + `SumCodePoints`, `1.85 μs → 0.19 μs`, `0.077x/C`).
    - `gopurs-strings/src/Data/String/CodeUnits.go` + `IndexByteASCII` (`simd_arm64.s`/`simd_amd64.s`) : recherche NEON/SSE2 64 B/itér. en passe unique (`indexOf-hit` `0.30 → 0.17 μs`, `0.87x/JS`, `0.15x/C`).
    - `NestedRecursion.purs` (`Ackermann` `16.09 → 0.43 μs`, `0.028x/C`).
  - **Round 3 — `ConstructorFold` CSE (`AST Evaluation`)** :
    - `ConstructorFold.purs` : CSE pur par branche (`withScope` + purge sur shadowing `Let`) pour éliminer les auto-appels récursifs identiques ; `AST Evaluation` passe de `0.027 μs → 0.011 μs` (`0.11x/C`, `2.6×`).
  - **Round 4 — Noyaux SIMD ARM64/AMD64 (`filter`, `cp-take`) & `RBTree` profond (`1.91 → 1.61 ms`)** :
    - `simd_arm64.s`/`simd_amd64.s` : `filterMaskEvenOddInt32Asm` hybride `VUZP1` + 3×`VADDP` (fast-path `VCMTST`+`VBIC` pour `target == 0`, prefix-XOR SSE2 sans LUT sur `amd64`) → `array/filter-10000` `2.96 → 1.34 μs` (`0.30x/C`).
    - `simd_arm64.s` : `sumCodePointsLimitAsm` & `countUTF8Asm` déroulés 32 B/itér. avec masques télescopiques partagés (`string/cp-take-unicode-1500` `1.20 → 0.45 μs`, `0.70x/C`).
    - `Ownership/BalancedLoop.purs` + `ModuleWorkers.purs` : bitmask registre `__dirMask uint64`, helper froid `_overflow`, attachement feuille direct avec sortie anticipée parent noir, élision d’écritures/barrières GC redondantes, garde `if __nc := arg; __nc != nil` (`guardNullableSelfCalls`) divisant par 2 les appels récursifs sur `depth`.
  - **Round 5 — Outlining froid (`ScalarArrayFold`, `ArrayProducerFold`), constantes UTF-8 contiguës, `RecordConsumer` équilibré & boucle de descente `BalancedLoop` hissée** :
    - `runtime/runtime.go` (`SumIntArraySlow`) + `ScalarArrayFold.purs` + `ArrayProducerFold.purs` : déport des fallbacks froids `TypeIntArray`/`TypeValue` dans des helpers dédiés et fast-path direct `TypeInt32Array && TypeInt32Array` sans `ArrayLength` (`array/append-1000` `0.072 → 0.066 μs`, `0.78x/C`, frame `576 B → 64 B`).
    - `simd_arm64.s` : bloc contigu `utf8_consts<>` (4 insns vs 27) et fenêtre d’accumulation 120 paires à vidage différé dans `sumCodePointsLimitAsm`.
    - `RecordShapeGuard.purs` + `RecordConsumer.purs` : garde stricte par pointeur (`integerReadsStrictShape`) + arbre d’addition binaire équilibré (`record/set-medium` `3.08x → 2.42x`, `record/set-large` `2.80x → 2.09x`, faisant baisser `Q₄/C` global à `2.42x`).
    - `Ownership/BalancedLoop.purs` : hissage de `__path[__depth] = __cur; __depth++; __dirMask <<= 1` avant la comparaison de clé (boucle de descente compactée à 12 insns avec `LDP` fusionné + back-edge `BLT` direct → `RBTree` `1.61 ms`, `0.19x/C`).

- [x] **5.2 — Reconstruction `bin/gopurs.js` + `bin/gopurs-native` + suite complète de tests**
  - `npm run build:ffi && npm run build && npm run build:native` : 0 warning, 0 erreur (`570` modules TAST, `362 411` types).
  - Tests verts : `test:ffi` (39/39), `test:identity` (8/8), `test:cache` (6/6, `-race`), `test:cache-codecs` (15/15), `test:profiles` (12/12 en `js` + `native-parallel`, `dev` + `release`), `tools/ffi-gen` (`go test -race`), `tools/build-cache` (`go test -race`), `gopurs-strings` (`code-points-native.mjs`, 286 entrées), 179 tests unitaires/runtime/SIMD (`-race`, Rosetta 2 `amd64`, cross-compile `linux/{arm64,amd64,riscv64}` + `js/wasm`) et les 6 suites E2E complètes (`array-producer-fold-e2e`, `codepoint-take-fold-e2e`, `constructor-fold`, `record-affine-consumer-e2e`, `tail-add-loop-e2e`, `traverse-maybe-total-e2e`) en `js` + `native-parallel`.

- [x] **5.3 — Campagne officielle `altbak.pub` (`measure.py`, 3 processus) & mise à jour `altbak.pub/README.md` (`Stable` vs `WIP`)**
  - Campagne officielle exécutée sur machine calme (`GOGC=800`, `-pgo=off`, `bin/gopurs-native`, `go-pure` 14 core + `go-pure-all` 23 library, 37/37 oracles validés sur 3 processus, archivée dans `altbak.pub/var/benchmark/go-opt-20261128/campaign/measurements/results.json`).
  - Graduation objective publiée dans `/Users/0x1/Documents/htdocs/altbak.pub/README.md` (vérifiée par `publish.backend_status_rows` + `test/benchmark-ratio-statistics.py` 9/9) :
    - **34 cas passés en `Stable`** (plancher matériel / fermé / `Bench.opaque` vérifié au `objdump`) : Total Stable `~ 0.0039 ms` (**`0.061x/C`**), `μ/C = 0.89x`, `Q₂/C = 0.83x`, `Q₄/C = 2.42x`, `G/C = 0.39x`, `Q₄/JS = 0.87x`.
    - **3 cas maintenus en `WIP`** (pistes d’optimisation générale encore identifiables) : `Red-Black Tree` (`~ 1612.49 μs`, `0.19x/C`), `array/filter-10000` (`~ 1.34 μs`, `0.30x/C`), `string/cp-take-unicode-1500` (`~ 0.45 μs`, `0.70x/C`) → Total WIP `~ 1.61 ms` (**`0.19x/C`**), `μ/C = 0.40x`, `Q₂/C = 0.30x`, `Q₄/C = 0.70x`, `G/C = 0.34x`.
    - **Global combiné 37 cas** : **`1.618 ms`** vs C `8.777 ms` (**`0.18x/C`**, **5.42× plus rapide que C**), **`μ/C = 0.85x`**, **`Q₂/C = 0.76x`**, **`G/C = 0.39x`**, **`Q₄/C = 2.42x`**, **`Q₄/JS = 0.87x`** (strictement plus rapide que JS V8 sur 37/37 cas).

---

## Journal de mesures

| Étape | Total / C (C `8.78 ms`) | `μ/C` | `Q₂/C` | `Q₄/C` | `G/C` | `Q₄/JS` | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Baseline officielle post-phase 2 (`HEAD c9f4b1a`) | `8.36 ms` (`0.95x`) | `1.34x` | `1.30x` | `3.08x` | `0.79x` | `1.55x` | Référence `altbak.pub/README.md` (37 cas WIP, 25/37 > `1x/C`) |
| Lot 1 (pairé 37/37) | — | — | — | — | — | — | `cu-length` `2.05×`, `ref-*` `1.33–1.50×`, `record/set-*` `1.21–1.27×`, `Lazy`/`ArrayOps` `1.13–1.15×` (0 div.) |
| Lot 1+2 (pairé 37/37) | — | — | — | — | — | — | `sort-descending` `9.43×` (`0.16x/C`, 0 alloc), `traverse/fail-last` `10.57×` (`0.11x/C`), `foldl-sum` `1.26×`, `array/map` `1.27×` (0 div.) |
| Lots 1–4 (pairé 37/37) | `8.26 ms` (`0.95x`) | `0.93x` | `0.82x` | `2.65x` | `0.47x` | `1.55x` | `Fib` `23.3×` (`0.05x/C`), `Ackermann` `37.5×` (`0.03x/C`), `Primes` `1.71×`, `cp-length` `3.07×`, `cp-take` `1.93×`, `concat` `2.55×`, 21/21 kernels `< 1x/C` |
| Rounds 2–3 (pairé 37/37) | `1.91 ms` (`0.22x`) | `0.87x` | `0.85x` | `2.43x` | `0.40x` | `0.87x` | `RBTree` `4.36×` (`0.22x/C`), `AST` `2.66×` (`0.11x/C`), `split` `9.8×` (`0.08x/C`), `indexOf-hit` `1.7×` (`0.87x/JS`, 37/37 `< 1x/JS`) |
| Round 4 (pairé 37/37) | `1.65 ms` (`0.19x`) | `0.85x` | `0.75x` | `2.57x` | `0.39x` | `0.87x` | `RBTree` `1.64 ms` (`0.19x/C`), `filter` `1.34 μs` (`0.29x/C`), `cp-take` `0.45 μs` (`0.70x/C`) |
| **Campagne officielle finale (`measure.py`, 3 proc., `gopurs-native`, 37/37)** | **`1.62 ms` (`0.18x`)** | **`0.85x`** | **`0.76x`** | **`2.42x`** | **`0.39x`** | **`0.87x`** | **34 Stable** (`0.0039 ms`, `0.061x/C`, `μ=0.89x`, `Q₂=0.83x`) + **3 WIP** (`1.61 ms`, `0.19x/C`, `μ=0.40x`, `Q₂=0.30x`, `Q₄=0.70x`) ; 37/37 oracles OK |
