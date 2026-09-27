# Gopurs — PBO : runtime, ordonnanceur et finitions (27 septembre 2026)

Ce fichier remplace le journal précédent (sauvegardé dans
`scratch/b8x-pbo-visibility-20260926/todo-journal-20260927.md`).

## État acquis

- **Backend** (b8x, 2 683 modules) : **48,5-48,7 s** en 8 workers (défauts du
  lanceur : `GOPURS_PBO_JOBS=8`, `GOPURS_PREPARE_JOBS=8`, `GOGC=off`,
  `GOMEMLIMIT=10GiB`), **74,5-82 s** en séquentiel ; **JS 77,4 s** →
  `/JS = 0,63x` (table `altbak.pub/README.md`). Mesure appariée de l'étape 2 :
  **56,1-56,5 → 48,5-48,7 s** à réglages identiques.
- **Allocations** : 214,2 → 189,1 → **175,8 Go** (profil échantillonné).
  Sept leviers livrés : rebox identité → cast, comparaisons `EvalRef` natives,
  `arrayBind` en `[]Value`, degré B-tree 6, native call lowering (v1→v3),
  imprimante « writer » avec builder natif, comparateurs natifs `Map` +
  `Set.map` incrémental (étape 2).
- **Parité** byte-exacte séquentiel ↔ parallèle ↔ JS (2 974 fichiers) ;
  fixtures validées sans mise à jour des snapshots.
- **Défauts du lanceur** (`gopurs/bin/gopurs`) : ≥ 32 Go → `GOGC=off`,
  `GOMEMLIMIT=10GiB`, `GOPURS_PBO_JOBS=8`, `GOPURS_PREPARE_JOBS=8`
  (surchargeables) ; branche `GOPURS_JS=1` → `node --stack-size=65536`.
- **Profil CPU (8 workers, après étape 2)** : comparateur `Map` **6,5 s**
  (contre 8,2 s), `Set.map` quadratique de `ThunkFusion` disparu ; `Apply`/
  `Apply2` et la machinerie Aff restent les premiers postes ; GC réelle ~2-3 %
  (`gctrace`, 37 cycles) — l'échantillonnage du profil mémoire Go reste
  désactivé sauf `GOPURS_ALLOC_PROFILE`.

## Plan

### 1 — Runtime Go et concurrence du compilateur (terminé)

- [x] Désactiver l'échantillonnage mémoire Go quand aucun profil n'est
      demandé : **~15 % de CPU en moins** (séquentiel 84-87 → **76,6 s** ;
      parallèle neutre car le surcoût y était recouvert).
- [x] `GOPURS_PREPARE_JOBS=8` par défaut (transitive 16,8 → 13,8 s,
      parallèle 57,5 → **53,7 s**).
- [x] Re-profilé (parallèle) : `lock2`/`osyield` ~18 % — dont une part
      d'**artefact** (le profiler échantillonne les threads endormis :
      `usleep`/`pthread_cond_*`) ; GC ~17 % ; `Apply` ~45 % cum ;
      comparateur `Map` ~4 %.
- [x] `GOMAXPROCS` 8/10/12 : **pires** que le défaut (14) → rien à changer
      (60,9 / 62,0 / 57,8 s vs 52,6-54,7 s).
- [x] `GOMEMLIMIT=12GiB` et `GODEBUG=madvdontneed=0` : sans gain → politique
      (off + 10 GiB) inchangée.
- [x] **Pool de workers** : non justifié par les mesures (peu de fibres
      vivantes, création en µs, `GOMAXPROCS` neutre) — abandonné au profit
      des étapes suivantes.

### 2 — Machines `Map` (le plus gros poste gopurs après le lowering)

Profil (parallèle, hors symboles runtime) : `(*Node).foldl` **17,7 s cum**
(callback PS par élément), `(*Node).find` **10,9 s cum**,
`Call_Data_Map_Internal_compareInt` **7,5 s dont 7,2 s dans `Apply2`**
(dispatch du dictionnaire) + `OrdStringImpl` 2,0 s ; `insert` 7,3 s,
`lookup` 6,8 s. **Nos helpers monomorphes** (`compareIdents`) ne pèsent que
**0,83 s** → les variantes monomorphes « côté PBO » ne valent pas le coup.

- [x] **Comparateurs natifs dans le B-tree (fork `gopurs-ordered-collections`)** :
      `lookupNativeImpl`/`insertNativeImpl` prennent un comparateur Go opaque
      (`asTree(m).withCompare(cmp)` par opération, arbre partagé intact), plus
      des helpers Go bruts `LookupNative`/`InsertNative` (valeur + booléen, le
      `Maybe` reste côté appelant). Comparateur `String` = `strings.Compare`
      sur la chaîne déballée. Tests natifs du fork ajoutés
      (`TestNativeComparatorEntryPoints`), suite complète OK.
- [x] **PBO : module `NativeMaps` (`{purs,go,js}`)** — comparateur
      `Qualified Ident` (déballe `Qualified[string]` du CoreFn décodé *et*
      `Qualified[Value]` de l'optimiseur : module puis ident) et comparateur
      `String` ; repli JS sur `Data.Map.Internal`. Routés :
      `BackendImplementations` (`Convert` : 2 insertions par binding,
      `lookupImplementation`, `lookupPurmetaImplementation` ; `Builder` :
      `createRankLookup`), `foreignSemantics`, `Map Ident Level` (`toLevel`),
      `Map ProperName` (`dataTypes`). Séparé de `FfiSupport` pour préserver le
      test FFI natif autonome (`native-ffi-support.mjs`).
- [x] **`Set.map` quadratique de `freshWorker`** (`ThunkFusion`,
      `FunctionFusion`) : les noms assainis étaient reconstruits à *chaque*
      groupe de bindings (`Set.map sanitizeName` sur l'ensemble complet) —
      ~4,8 s CPU pour ThunkFusion seul. Ensemble `emitted :: Set String`
      maintenu incrémentalement.
- Mesures appariées (même session, deux binaires) : 8 workers **60,6 → 54,1 s**
      (avec le `GOPURS_PREPARE_JOBS` par défaut), **56,1-56,5 → 48,5-48,7 s**
      (défauts du lanceur) ; séquentiel **84-91 → 74,5-82 s** ; CPU total **192,6 →
      185,0 s** ; allocations **190,3 → 175,8 Go** ; `compareInt` **8,2 →
      6,5 s**. Parité byte-exacte (0 `.go` modifié sur 2 974).
- [ ] **Reste** : le coût `compareInt` restant (~6,5 s) est une longue traîne —
      `LookupImpl` 3,3 s (aucun site dominant : `foreignSemantics` puis
      `Preparation`, planificateur, `Monomorphize`, `GoTypes`) et
      `unionWithSameOrdering` 2,5 s (`Map.union` des directives `EvalRef` et
      des instantiations). Prochains leviers : comparateur `EvalRef` natif
      (lookup + union), unions natives `String`, maps `Int` du planificateur,
      `isIdentity` de `FunctionFusion` (~1,9 s CPU).
- [ ] **`Map.foldl` chauds** : remplacer par `toUnfoldable` + boucle
      (supprime le callback boxé par élément, `FoldlImpl.func1` ~8,6 s cum).

### 3 — GC et tas vivant (clos)

- [x] `gctrace` : **GC = 2-3 % de CPU** (39 cycles, vivant 1,5-2,7 Go après
      GC, pauses ~0,1-0,3 ms) — les ~17 % vus au profil étaient un artefact
      (échantillonnage des threads endormis). Rien à gagner côté politique ;
      la réduction du tas vivant reste marginale. Politique (off + 10 GiB)
      conservée.

### 4 — Extensions du lowering FFI

- [ ] FFI `Effect` (débloque des builders purs côté FFI) et appels via
      dictionnaires ; `int`/`map`/opaques.

### 5 — `gopurs-*` structurel (protocole élargi : fixtures + runtime b8x)

- [ ] `gopurs-aff` : pool de goroutines pour les fibres (bénéfice compileur
      **et** runtime b8x) — après la variante gopurs de l'étape 1.
- [x] `gopurs-ordered-collections` : comparateurs natifs (livrés à l'étape 2 :
      points d'entrée à comparateur opaque + helpers Go bruts + comparateur
      `String`).

### 6 — Divers

- [ ] Chaînes : derniers `<>` par builders (cf. `memmove` du profil).
- [ ] `mangleType` : mémoïsation par identité (2-4 Go potentiels).
- [ ] Relances du builder parallèle (réduire les ~830 tentatives rejetées).

## Méthode de validation

- Parité byte-exacte (b8x + fixtures), profil d'allocations
  (`GOPURS_ALLOC_PROFILE`) **et** profil CPU (`PPROF=1` → `cpu.prof`).
- Temps appariés, machine au repos ; ne pas superposer builds et campagnes.
- Piège connu : toute structure polymorphe accumulée peut produire des rebox
  quadratiques (incidents `GoImports`, liste de l'imprimante : 6 297 Go) —
  toujours profiler, pas seulement vérifier la parité.
- Pour la table `altbak.pub` : comparer le **backend seul** (hors bootstrap du
  compilateur et hors frontend purs).

## Références

- Campagnes : `scratch/b8x-pbo-visibility-20260926/` (`printer2/`,
  `jsbackend/`, `lowering/`, `verify/`).
- Profils : `jsbackend/cpu.prof`, `printer2/alloc-j1.prof`.
- Code : `gopurs/gopurs/src/Gopurs/` (`Main.purs`, `Emission.purs`,
  `Monomorphization.purs`, `Preparation.purs`, `FfiBridge.purs`)
  et `purescript-backend-optimizer-gopurs/`.
