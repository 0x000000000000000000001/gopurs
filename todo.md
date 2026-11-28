# Plan `gopurs` — SIMD généraliste (Go 1.27+) & optimisations compilateur générales

> **État :** **Terminé et validé** (Go 1.27.4 `linux/arm64` dans `b8x` Docker + Go 1.27.0 `darwin/arm64` hôte, campagne officielle 3 processus publiée dans `altbak.pub/README.md`).

## Baselines officielles de référence (`altbak.pub/README.md`, post-Option A — 0 ASM, 0 hack benchmark)

Campagne officielle 3 processus (`altbak.pub/var/benchmark/go-clean-20261128/`, `GOGC=800`, `-pgo=off`) comparée à la référence C11 idiomatique (`3.69 ms`, `clang -O3` sans SIMD manuel) :

| Périmètre Go Release (`gopurs`) | Total Go | Total C11 | `/C` | `μ/C` | `Q₂/C` | `Q₄/C` | `G/C` |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| **Stable (25 cas)** | `1.62 ms` | `3.65 ms` | `0.44x` | `1.27x` | `1.31x` | `2.36x` | `1.06x` |
| **WIP (12 cas)** | `0.33 ms` | `0.041 ms` | `8.24x` | `17178.56x` | `3.57x` | `205317.94x` | `22.85x` |
| **Tous les 37 cas** | `1.96 ms` | `3.69 ms` | `0.53x` | `5572.29x` | `1.50x` | `205317.94x` | `2.87x` |

Cas qui étaient en **WIP (12)** avant exécution :
- **Cœur (6) :** `List Processing` (`0.41 μs`), `Tail Call Optimization` (`63.00 μs`, `1.94x/C`), `Deep Record Updates` (`8.77 μs`, `2.71x/C`), `Church Numerals` (`231.92 μs`), `Prime Sieve` (`16.07 μs`), `Array Processing` (`0.32 μs`).
- **Bibliothèque (6) :** `array/map-10000` (`1.79 μs`, `3.98x/C`), `array/foldl-sum-10000` (`0.90 μs`, `3.12x/C`), `array/sort-descending-10000` (`7.50 μs`, `2.63x/C`), `array/append-1000` (`0.21 μs`, `2.46x/C`), `foldable/traverse-maybe-1000` (`0.18 μs`, `4.17x/C`), `string/cp-length-unicode-1000` (`3.86 μs`, `3.15x/C`).

## Principes non négociables
- **100 % généralité :** aucune passe ni aucun kernel façonné pour un benchmark (zéro `AffineParity*`, zéro `FilterMask*`, zéro `SumCodePoints*`, zéro `SumOddProgression`, zéro forme close de modulo/progression).
- **DX invisible (zéro variable d'env) :** activation SIMD out-of-the-box sur un simple `go build` dès Go 1.27+, sans jamais demander `GOEXPERIMENT=simd`.
- **Portabilité & transition future vers `import "simd"` :**
  - Les `.s` sont gardés par `//go:build (arm64 || amd64) && go1.27` (respectivement `arm64 && go1.27` et `amd64 && go1.27`).
  - Fallback Go pur automatique via `runtime/simd_generic.go` (`//go:build !(arm64 || amd64) || !go1.27`) pour Go ≤ 1.26, `js/wasm`, `riscv64`, etc.
  - Aucun conflit possible au passage Go 1.28+ : le tag `go1.27` continue de s'activer proprement, et le jour où Go sortira `import "simd"` de `GOEXPERIMENT`, seules les implémentations internes de `gopurs_runtime` seront remplacées (zéro changement dans le codegen, `gopurs-strings` ou `b8x`).
- **Sémantique PureScript / JS exacte & empirisme :** vérification par oracles exhaustifs (UTF-8/WTF-8, surrogates, débordements Int32, `minInt32`), `pprof`/`objdump` avant chaque optimisation compilateur, et validation réelle sur `b8x` (2 711 modules, Go 1.27 Docker).

---

## Lot 1 — Micro-socle SIMD généraliste dans `gopurs_runtime` (`arm64` NEON + `amd64` SSE2, gardé `&& go1.27`)

Objectif : réintroduire un petit noyau propre (~400–550 lignes/arch) strictement limité aux primitives universelles de tableaux `[]int32` et chaînes UTF-8/ASCII, avec fallback Go pur garanti.

- [x] **1.1 — Primitives tableaux `[]int32` (universelles)**
  - `SumInt32(xs []int32) int32` — réduction additive pure (`foldl (+) 0`, `array/foldl-sum-10000` : `0.90 → 0.24 μs`, `0.84x/C`).
  - `IsNonIncreasingInt32(xs []int32) bool` & `IsNonDecreasingInt32(xs []int32) bool` — détection de run monotone (`Data.Array.sort`, `array/sort-descending-10000`).
  - `IndexOfInt32(xs []int32, target int32) int` & `ContainsInt32(xs []int32, target int32) bool` — recherche linéaire (`elem`, `elemIndex`).
- [x] **1.2 — Primitives chaînes UTF-8 / ASCII (universelles, fort ROI prod `b8x`)**
  - `CountUTF8CodePoints(s string) int64` — validation RFC 3629 NEON/SSE2 16–32 B/itér. avec repli WTF-8 sur `0xED` ou séquence malformée (`Data.String.CodePoints.length`, `string/cp-length-unicode-1000` : `3.86 → 0.54 μs`, `0.44x/C`).
  - `IndexByteASCII(s string, c byte) (int, bool)` — recherche 1er octet + détection préfixe ASCII en passe unique (`Data.String.CodeUnits.indexOf`, `string/indexOf-hit-10000` : `1.27 → 0.16 μs`, `0.15x/C`; `miss` : `0.21 → 0.17 μs`, `0.14x/C`).
  - `FirstNonASCII(s string) int` & `IsASCII(s string) bool` — scan NEON/SSE2 pour court-circuiter les chemins purement ASCII (`CodeUnits.unitBoundary`).
  - `AdvanceUTF8CodePoints(s string, n int64) (int, bool)` — offset octet du `n`-ième code point (utilisant `FirstNonASCII` + `countUTF8Asm` par blocs sans ASM additionnel, branché dans `gopurs-strings` `CodePoints.{_Take, _CodePointAt}`).
- [x] **1.3 — Plomberie build & cache + tests multi-cibles**
  - Mis à jour `tools/build/embed-runtime.mjs`, `src/Gopurs/Runtime/Simd.purs`, `tools/build-cache/{contract.go,session.go}`, `tools/lib/go.mjs`, `gopurs-strings/src/Data/String/{CodePoints.go,CodeUnits.go}`.
  - Vérifié dans `tools/simd-kernels.test.mjs`, `tools/strings-kernels.test.mjs`, `tools/scalar-array-fold.test.mjs` : oracles scalaires + fuzzing UTF-8/WTF-8/surrogates, `go vet`, `go test -race`, exécution native `arm64` + Rosetta `amd64`, cross-compile `linux/arm64`, `linux/amd64`, `linux/riscv64`, `js/wasm`, et validation du build fallback pur Go (`!go1.27`).

---

## Lot 2 — Peephole compilateur général : `div` / `mod` euclidien par constante positive

Objectif : supprimer le coût du helper `IntMod`/`IntDiv` (garde `d == 0` + `SDIV` + branches) quand le diviseur est une constante entière `d > 0` connue à la compilation — 100 % général pour tout code PureScript.

- [x] **2.1 — Réduction de force dans `PrimitiveExprs` / `Int32Loops` / `CodeGen` / `RecordReadGroups`**
  - Puissance de 2 (`d = 2^k > 0`) : `x mod 2^k` → `int64(int32(x) & (d - 1))` (et `& int32(d-1)` dans `Int32Loops`), `x div 2^k` → `int64(int32(x) >> k)` (exact pour tout `int32` signé, y compris les négatifs et `minInt32`, en sémantique euclidienne PureScript).
  - Constante positive quelconque (`0 < d <= 2147483647`) : helpers inlinés `IntModPos32`/`IntModPos64`/`IntDivPos32`/`IntDivPos64` dans `runtime/runtime.go` avec fast-path `uint32` pour `0 <= a <= MaxInt32` (génère `UMULL`/`LSR` 32 bits sans garde `d == 0`), plus `moduloZeroComparison` (`x mod y == 0` / `!= 0` → `IntModZero` / `(a & (d-1)) == 0` / `(a % d) == 0`).
- [x] **2.2 — Validation exhaustive & mesure**
  - Ajouté `tools/divmod-peephole.test.mjs` couvrant `minInt32`, `maxInt32`, négatifs, multiples exacts, débordements 64 bits, `d ∈ {1, 2, 3, 4, 7, 10, 16, 2^30}`, `d <= 0`.
  - Mesuré sur campagne officielle 3 processus : `Tail Call Optimization` `63.00 → 39.00 μs` (`-38%`, `1.20x/C`), `Deep Record Updates` `8.77 → 7.48 μs` (`-15%`, `2.31x/C`), `Prime Sieve` `16.07 → 15.07 μs` (`-6.2%`).

---

## Lot 3 — Empirisme & optimisations générales sur boucles `Prime Sieve` et boucles fusionnées scalaires

Objectif : attaquer de manière 100 % générale les goulots restants sur `Prime Sieve` et les boucles fusionnées `[]int32`, uniquement guidés par le profiling empirique et BCE.

- [x] **3.1 — Profiling empirique sur `Prime Sieve` (`16.07 μs → 15.07 μs`)**
  - L'inspection réelle de `Test/Primes/Primes.purs` a montré une liste chaînée stricte `data List a = Nil | Cons a (List a)` (et non `STArray`), déjà abaissée par `ListFilter.purs` (partage de suffixe sans copie, 1 nœud par élément copié, contrat vérifié sur 511 masques dans `tools/list-filter.test.mjs` et `tools/list-filter-ir.test.mjs`).
  - Décision retenue : conserver `ListFilter.purs` intact (pas de slab allocateur qui risquerait de retenir des blocs morts en GC et casserait le contrat exact d'allocation) ; `Prime Sieve` bénéficie directement de `IntModZero` (Lot 2 : `16.07 → 15.07 μs`, `-6.2%`).
- [x] **3.2 — BCE & dépliage 8 voies général sur les boucles fusionnées (`IndexedAddLanes`, `ArrayProducerFold`, `ArrayTraverseFold`)**
  - Ajouté `IndexedAddLanes.emitContributionReverse` (8 voies déroulées sans aucun bounds check dans la boucle principale, vérifié via `-gcflags=-d=ssa/check_bce/debug=1` et 2 760 comparaisons exhaustives dans `tools/indexed-add-lanes.test.mjs`) et branché sur la branche décroissante `SortPlan` de `ArrayProducerFold.purs`.
  - Vérifié que `emitContribution` forward (déjà utilisé par `ArrayProducerFold` et `ArrayTraverseFold`) n'a déjà **0 bounds check** dans ses boucles principales 8 voies.
  - Effet combiné NEON `IsNonIncreasingInt32` + `emitContributionReverse` sur `array/sort-descending-10000` : `7.50 → 1.45 μs` (**`5.2×` plus rapide**, `0.51x/C`).

---

## Lot 4 — Intégration finale, vérification `b8x` (Go 1.27) & publication `altbak.pub`

- [x] **4.1 — Bootstrap & suites complètes `gopurs`**
  - `npm run build` (0 warning / 0 erreur), `npm run build:native` (566 modules, 357 444 types), `tools/build-cache` (`go test ./...`), `tools/ffi-gen` (`go test ./...`), `test:identity` (8/8), `test:cache` (6/6 sous `-race`), `test:ffi` (39/39), `simd-kernels`, `strings-kernels`, `divmod-peephole`, `indexed-add-lanes`, `array-producer-fold`, `list-filter-ir`, `scalar-array-fold`, `tail-add-loop-e2e` (`js` + `native-parallel`).
- [x] **4.2 — Vérification réelle sur `b8x` (conteneur Docker Go 1.27)**
  - Build complet `b8x` (2 711 modules, 0 warning / 0 erreur), présence des 5 fichiers dans `output/gopurs_runtime/`, compilation vérifiée dans `core-api-cli-1` (`go1.27.4 linux/arm64`) sur `./gopurs_runtime` et `./Test.Main/main/main.go` (`B8X_GO127_BUILD_OK`).
- [x] **4.3 — Campagne officielle 3 processus `altbak.pub` & mise à jour `README.md`**
  - Campagne 3 processus (`altbak.pub/var/benchmark/go-simd-general-20261128/`, `go1.27.0 darwin/arm64`, sans `GOEXPERIMENT`, `GOGC=800`, `-pgo=off`), 37/37 oracles vérifiés et 5 symboles NEON liés confirmés par `go tool nm`.
  - Mis à jour `altbak.pub/README.md` (**30 Stable / 7 WIP**, 5 graduations vers Stable : `Tail Call Optimization`, `Deep Record Updates`, `array/foldl-sum-10000`, `array/sort-descending-10000`, `string/cp-length-unicode-1000`), `test/benchmark-ratio-statistics.py` 9/9 OK.

---

## Journal de mesures

| Étape | Total 37 cas (`/C`, C `3.69 ms`) | Stable (`G/C`, `Q₂/C`, `Q₄/C`) | WIP (nb cas) | Notes |
| --- | --- | --- | --- | --- |
| Baseline post-Option A (`go-clean-20261128`, 0 ASM) | `1.96 ms` (`0.53x`) | 25 cas (`1.06x`, `1.31x`, `2.36x`) | 12 cas | Référence précédente dans `altbak.pub/README.md` |
| **Lots 1–3 : SIMD généraliste Go 1.27+ & peephole `div`/`mod` + reverse lanes** (`go-simd-general-20261128`) | **`1.91 ms` (`0.52x`, `Q₂/C = 1.31x`, `G/C = 2.30x`, `Q₄/JS = 0.86x`)** | **30 cas (`1.66 ms`, `0.45x/C`, `G/C = 0.95x`, `Q₂/C = 1.29x`, `Q₄/C = 2.35x`, `G/JS = 0.0018x`)** | **7 cas (`0.25 ms`)** | 0 hack benchmark, 0 `GOEXPERIMENT`, 5 cas gradués en Stable, 37/37 cas plus rapides que JS |


---

# Phase 2 — Plan d'optimisation des 10 cas Go WIP (guidé par `pprof` / `objdump`, 100 % général)

> **État :** **Terminé et validé** (Go 1.27.4 `linux/arm64` dans `b8x` Docker + Go 1.27.0 `darwin/arm64` hôte, campagnes officielles 3 processus publiées dans `altbak.pub/README.md`).

## Baselines officielles de référence (`altbak.pub/README.md`, 28 nov. 2026, Go 1.27.0 `darwin/arm64`)

| Périmètre Go Release (`gopurs`) | Total Go | Total C | `/C` | `/JS` |
| --- | ---: | ---: | ---: | ---: |
| **Total Stable (30 cas)** | `1.66 ms` | `3.67 ms` | `0.45x` | `0.015x` |
| **Total WIP (10 cas)** | `29.40 ms` | `17.87 ms` | `1.65x` | `0.25x` |
| **Total (40 cas)** | `31.05 ms` | `21.54 ms` | `1.44x` | `0.14x` |

Les 10 cas WIP à attaquer (par masse temporelle & empirisme `pprof`/`objdump`) :
1. `JSON to Typed AST` (`19733.00 μs`, `1.87x /C`, `0.24x /JS`) — `27.1 MB`/pass, `278k` allocs/pass (`directIndex` `[]directToken` = **38.6 %** des octets, `strings.Clone` = **12.8 %** des objets, `tc_cndJust`/`tc_ntJust` = **17.1 %**, `RecordShape4`/`2` = **16.2 %**, lexer `string`/`space` octet par octet).
2. `Array Indexing` (`8139.67 μs`, `1.24x /C`, `0.32x /JS` vs Imp. Go FFI `3241.50 μs` = `0.49x /C`) — `objdump` confirme dans la boucle chaude (8,39 M itér.) : 3 tests de tag `ArrayAccess` + rechargement du slice header + reboxing `unsafeIndex` (`Int(...).IntVal`) + `IntAdd(..., 0)` + `index_6` resté `int64` faute de suivi du temporaire de branche `__t2`.
3. `JSON Decoding` (`1276.75 μs`, `1.94x /C`, `0.14x /JS`) — `2.23 MB`/pass, `33.7k` allocs/pass (`strings.Clone` = **22.1 %** des objets, `RecordShape3`/`5` = **26.9 %** des octets, croissance slice `cap=4` dans `DecoderSchemas/Stream` = **15.1 %** des octets, `Array`/`Just`/`Rebox`, lexer `argonautStreamScanString` octet par octet).
4. `Church Numerals` (`231.71 μs`, `0.18x /JS`) — **62.5 % flat / 84.7 % cum** dans `gopurs_runtime.Apply` (dispatch de fonctions d'ordre supérieur boxées).
5. `Prime Sieve` (`15.07 μs`, `0.26x /JS`) — **94.3 %** des objets heap sont des nœuds `Cons` (dont **64.1 %** dans `ListFilter`).
6. `array/map-10000` (`1.75 μs`, `3.89x /C`, `0.014x /JS`), `array/append-1000` (`0.21 μs`, `2.43x /C`), `foldable/traverse-maybe-1000` (`0.18 μs`, `4.05x /C`) — `objdump` révèle **16 instructions `NOOP`** (`InlMark` Go sur les 2 closures inlinées × 8 voies) + **8× `SXTW` + 4× `SBFIZ`** dus aux accumulateurs `int64` et aux appels de closures dans `emitContribution`.
7. `List Processing` (`0.40 μs`, `0.047x /JS`) & `Array Processing` (`0.31 μs`, `0.043x /JS`) — déjà >20× plus rapides que JS ; zéro forme close autorisée (`todo.md` principes), gains collatéraux uniquement via passes générales.

---

## Phase 2 — Lot 1 : Passes compilateur générales (LICM `ArrayAccess`, `Int32Loops` temporaires de branche, inlining `int32` des voies scalaires)

- [x] **P2-1.1 — `ArrayIntrinsics/Index.purs` (`unsafeIndex`)** : retourner `converted` avec `exprType: elementType` sans reboxer systématiquement en `TypeValue`.
- [x] **P2-1.2 — `Int32Loops.purs`** :
  - Suivi à point fixe des variables locales `int64` normalisées (temporaires de branche `__tN` dont toutes les écritures sont normalisées ou self-préservantes) pour promouvoir les compteurs conditionnels (`index_6`, `total_6`, etc.).
  - Supprimer l'enrobage `IntAdd(w, int64(0))` dans `narrow` quand `w` est une expression runtime non-constante.
  - **LICM de représentation de tableau `Value`** : dans les boucles pures indexant une variable `TypeValue` invariante via `ArrayAccess`/`ArrayLength` (directement ou via alias local), hisser le dispatch de représentation (`TypeInt32Array`, `TypeIntArray`, `TypeStringArray`, `TypeArray`, fallback) une seule fois avant la boucle.
- [x] **P2-1.3 — Propagation d'amplitude `int32` (`PrimitiveExprs`) & inlining `int32` sans closure dans `IndexedAddLanes` / `ArrayProducerFold` / `ArrayTraverseFold`** :
  - Étendre `PrimitiveExprs.integerMagnitude` (`Let`, `Branch`, `OpIntNegate`, `OpIntBitNot`) et supprimer l'IIFE superflue de `boundedIntProduct`.
  - Dans `emitContribution32` / `emitContributionReverse32`, accumuler les 8 voies en `int32` et inliner directement l'expression scalaire `int32` prouvée (0 `InlMark` `NOOP`, 0 `SXTW` redondant, préservation exacte du seuil `2^53` IEEE-754 et des graines larges).
- [x] **P2-1.4 — Tests & mesure empirique** (`int32-loops`, `array-unsafe-index`, `indexed-add-lanes`, `array-producer-fold`, `array-traverse-fold` ; résultats mesurés : `Array Indexing` `8139.67 → 4762.0 μs` = `0.73x /C` plus rapide que C, `array/map-10000` `1.75 → 1.096 μs`, `traverse-maybe-1000` `0.18 → 0.116 μs`, `array/append-1000` `0.21 → 0.175 μs`, `List Processing` `0.40 → 0.323 μs`).

## Phase 2 — Lot 2 : Réduction d'allocations & inlining des décodeurs JSON (`sync.Pool`, arènes `noscan`, `vsChecker` bitset+slab, inlining feuilles scalaires)

- [x] **P2-2.1 — Étude empirique SIMD lexical JSON (rejeté sur preuves)** : testé `ScanJSONString`/`SkipJSONSpace` NEON/SSE2 sur les deux corpus (`JsonTypedAst` : 99,1 % des chaînes < 16 B, moyenne 6,4 B, 0 run d'espaces ≥ 16 ; `JsonDecoding` : 96,6 % < 16 B, moyenne 7,2 B, 0 run ≥ 16). L'appel cross-package cassait l'inlining de `space()`/`take()` (+220 μs sur `JsonDecoding`). Retiré l'ASM inutile pour conserver zéro code mort dans `gopurs_runtime`.
- [x] **P2-2.2 — `sync.Pool` & arènes `noscan` owned dans `CoreFn.Json.Text`, `CoreFn.Usage` et `gopurs-argonaut-codecs`** :
  - `directTokenPool` + `ntTableScratchPool` (`refs`, `pending`, `state`, `argsSlab`) + `argonautTextTokenPool` + `unquoteScratch` réutilisable.
  - Arène `noscan` 4 KB (`directDocument.stringChunk` + `cloneBytesOwned` dans `tc_cndModuleName`, `argonautStreamParser.stringChunk` avec fast-path inliné) éliminant ~100 % des allocs `strings.Clone` et `strings.Join`/`MakeNoZero` tout en préservant l'isolation inter-appels (`owned-reference`).
  - `CoreFn.Usage` (`vsChecker`) : bitset `[]uint64` + slab `[]vsScope` poolé avec `bindingID` intégré (zéro `origin + "#" + FormatInt`, zéro `map[string]bool`, zéro `&vsIdentity`).
  - `cndQualified` : suppression des closures `cndTryString` redondantes sur le chemin rapide.
- [x] **P2-2.3 — `DecoderSchemas/Stream.purs` & `runtime.go`** : inlining des feuilles scalaires (`Int`, `Number`, `String`, `Boolean`) à `cap=4` (validé A/B vs `cap=8` qui ajoutait +240 KB) et singletons vides zéro-alloc dans `runtime.go`.
- [x] **P2-2.4 — Tests oracles & mesure empirique** : `decoder-stream.test.mjs` (11/11 sous `-race`), `TestPublicModuleParserAndErrorBoundary` (2 025 cas d'erreur sous `-race`), 12 & 17 empreintes SHA-256 vérifiées. Résultats : `JsonTypedAst` **`19733 → 17922 μs`** (`-9,2 %`, allocs régime établi `27,1 → 12,6 MB/pass` = `-53,5 %`), `JsonDecoding` **`1276.75 → 1101.9 μs`** (`-13,7 %`, `0.119x /JS`).

## Phase 2 — Lot 3 : Passes FP générales (HOF unboxés via TAST, composition d'itérations comptées, inlining `ListFilter` & réutilisation linéaire `Cons`)

- [x] **P2-3.1 — HOF monomorphes natifs via TAST (`func(int64) int64`) & composition d'itérations comptées dans `FunctionFusion.purs` / `FunctionExprs.purs` / `ModuleWorkers.purs`** : attachement `WithFunctionData(..., func(int64) int64)` sur les lambdas `Func [Int] Int`, enveloppement `CountedIntIterator` & `ComposeCountedIntIterators` (aplatissement `(g^m)^n = g^(n*m)` avec garde d'overflow `int64` et préservation exacte des effets/panics). Résultat mesuré : `Church Numerals` **`231.71 → 71.99 μs`** (**`3.22×`**, `0.057x /JS`).
- [x] **P2-3.2 — TRMC linéaire, réutilisation de cellules `Cons` & déforestation de pipelines fermés `fold (transform (range ...))` (`LinearList.purs` + `Divisor32`)** : reconnaissance des listes chaînées `Int` locales (monomorphes ou polymorphes instanciées à `Int`), filtres locaux (inversés ou directs), transformations TRMC linéaires (`__gopurs_linear_*` sur arguments prouvés frais uniquement, préservation stricte des listes partagées), `Divisor32` Lemire hissé hors boucle, et déforestation zéro-allocation des pipelines fermés `fold (linearTransform (range lo hi))` (`__gopurs_linear_fold_*` + suppression des `LetRec` purs devenus morts). Résultat mesuré : `Prime Sieve` **`15.07 → 3.52 μs`** (**`4.28×`**, `0.060x /JS`, `0 allocs/op`).
- [x] **P2-3.3 — Tests & validation empirique** (`counted-functions.test.mjs`, `divmod-peephole.test.mjs` incluant 200 000 paires `Divisor32` sous `-race` + Rosetta `amd64`, `linear-list.test.mjs` sous `-race` + `AllocsPerRun == 0`, `list-filter-ir.test.mjs`, `list-filter.test.mjs`, `tools/c1c` `go test ./...`).

## Phase 2 — Lot 4 : Validation finale `b8x` (Go 1.27) & campagne officielle `altbak.pub`

- [x] **P2-4.1 — Suites complètes `gopurs` + bootstrap natif + build réel `b8x` (2 711 modules, Docker Go 1.27)** : `npm run build` (0 warning / 0 erreur), `npm run build:native` propre sans ASM mort (567 modules, 362 487 types), `tools/{build-cache,ffi-gen,c1c}` (`go test ./...`), `test:identity` (8/8), `test:cache` (6/6 `-race`), `test:ffi` OK, 109/109 tests compilateur/runtime/e2e (`js` + `native-parallel`), `TestPublicModuleParserAndErrorBoundary` (2 025 cas `-race`), et build complet `b8x` (2 711 modules) validé dans `core-api-cli-1` (`go1.27.4 linux/arm64`, `B8X_PHASE2_GO127_BUILD_OK`).
- [x] **P2-4.2 — Campagne officielle 3 processus `altbak.pub` (Core/Library + Diagnostics), graduations Stable et mise à jour `altbak.pub/README.md`** : campagne `altbak.pub/var/benchmark/go-phase2-20261128/` (`go1.27.0 darwin/arm64`, `GOGC=800` Core+Lib, `GOGC=100` Diagnostics, `-pgo=off`, zéro `GOEXPERIMENT`), 40/40 oracles & empreintes SHA-256 validés, 4 graduations vers Stable (**34 Stable / 6 WIP** : `Array Indexing`, `array/map-10000`, `array/append-1000`, `foldable/traverse-maybe-1000`), `test/benchmark-ratio-statistics.py` **9/9 OK**.

### Journal de mesures Phase 2 (40 cas, référence `altbak.pub/README.md`)

| Étape | Total 40 cas (`/C`, `/JS`) | Total Stable (`/C`, `G/C`, `Q₂/C`, `Q₄/C`) | Total WIP (nb cas) | Gains clés |
| --- | --- | --- | --- | --- |
| Baseline pré-Phase 2 (`28 nov. 2026`, 30 Stable / 10 WIP) | `31.05 ms` (`1.44x /C`, `0.14x /JS`, `G/C = 2.24x`) | 30 cas (`1.66 ms`, `0.45x /C`, `G/C = 0.95x`, `Q₂/C = 1.29x`, `Q₄/C = 2.35x`) | 10 cas (`29.40 ms`, `1.65x /C`, `0.25x /JS`) | Référence officielle `altbak.pub/README.md` |
| **Phase 2 Lots 1–3 (`go-phase2-20261128`, 34 Stable / 6 WIP)** | **`24.68 ms` (`-20.5%`, `1.15x /C`, `0.11x /JS`, `G/C = 1.99x`, `G/JS = 0.0038x`)** | **34 cas (`6.43 ms`, `0.63x /C`, `G/C = 1.02x`, `Q₂/C = 1.31x`, `Q₄/C = 2.62x`, `0.046x /JS`)** | **6 cas (`18.24 ms`, `1.63x /C`, `0.20x /JS`, `Q₂/JS = 0.058x`)** | `Array Indexing` `-41.2%` (`0.73x /C`), `JSON to Typed AST` `-14.3%` (`1.60x /C`), `Church Numerals` `-69.0%` (`0.057x /JS`), `Prime Sieve` `-77.4%` (`0.059x /JS`, `0 allocs/op`), `array/map-10000` `-37.3%`, `traverse-maybe-1000` `-35.4%`, `List Processing` `-22.6%`, `array/append-1000` `-17.7%`, `array/sort-descending-10000` `-10.7%`, 0 régression |
