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
