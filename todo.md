# Gopurs — PBO : état et chantiers (27 septembre 2026)

## Acquis (b8x, 2 683 modules)

- **Backend** : **46,0-46,4 s** en 8 workers (défauts du lanceur),
  **69,3-69,5 s** en séquentiel ; **JS 77,4 s** → `/JS = 0,60x`
  (table `altbak.pub/README.md`, cellule b8x ~46,4 s).
- **Allocations** : 214,2 → **174,8 Go** (profil échantillonné) ; CPU total
  8 workers : 192,6 → **178,4 s** ; `compareInt` 8,2 → **4,3 s**.
- **Parité** byte-exacte séquentiel ↔ parallèle ↔ JS (2 974 fichiers `.go`).
- **Lanceur** (`gopurs/bin/gopurs`) : ≥ 32 Go → `GOGC=off`,
  `GOMEMLIMIT=10GiB`, `GOPURS_PBO_JOBS=8`, `GOPURS_PREPARE_JOBS=8`
  (surchargeables) ; `GOPURS_JS=1` → `node --stack-size=65536`.
- **Leviers livrés** : rebox identité → cast, comparaisons `EvalRef` natives,
  `arrayBind` en `[]Value`, degré B-tree 6, native call lowering (v1→v3),
  imprimante « writer » avec builder natif, comparateurs natifs `Map`
  (`NativeMaps` + fork `ordered-collections`) et `Set.map` incrémental.
- **Écartés** (seuil ci-dessous) : reste de `compareInt` (4,3 s), unions/maps
  résiduels, `FunctionFusion.isIdentity` (1,8 s intrinsèque), `mangleType`,
  politique GC/tas (GC réelle 2-3 %).

## Chantiers (par impact murale)

1. [ ] **Relances du builder parallèle** — ~830 tentatives rejetées à
   8 workers, +40 % d'allocations. Indices d'ordonnancement profondeur 2,
   reprise plus fine que la reconversion complète. Le plus gros levier restant
   (seule piste crédible sous ~45 s).
   - **Vue vivante testée puis écartée** (27/09) : miroir `Ref`
     implémentations + directives, utilisable quand la contribution de
     directives du prédécesseur est vide. Résultat : **−41 relances sur 844**
     seulement (mesure sous charge, compteurs indicatifs) — les rejeux viennent
     de courses réelles (le prédécesseur n'est pas finalisé au moment du
     lookup), pas de l'instantané. Sous le seuil : sources et binaire restaurés.
   - Restes possibles mais incertains : ordonnancement strict (moins de
     parallélisme), priorité aux relances, checkpoint de conversion — à
     n'ouvrir que si une mesure au repos démontre ≥ ~2 s.
2. [ ] **Chaînes / `memmove`** (~15 s CPU) — généraliser le writer natif de
   l'imprimante aux derniers `<>`/`joinWith` chauds (émission, préparation).
3. [ ] **Lowering FFI v4** — FFI `Effect` (builders purs côté FFI), appels
   via dictionnaires, `int`/`map`/opaques ; cible la masse `Apply`/`Apply2`
   (~90/70 s cumulés).
4. [ ] **`Map.foldl` chauds** (~8,6 s cum, callback currifié par élément) —
   entrée native à callback non currifié ou sites convertis en boucles ;
   uniquement si le gain murale est démontré.
5. [ ] **`gopurs-aff` (runtime b8x)** — pool de goroutines pour les fibres,
   pour les programmes compilés (et la machinerie Aff du compilateur).
   - **Bug connu (fixtures)** : sur les petits corpus, une compilation
     parallèle peut finir en `all goroutines are asleep - deadlock!`
     (fibres forkées non relâchées après une erreur de conversion).
     Repro ~1/600, non reproductible en boucle serrée ; `bin/test` force
     désormais `GOPURS_PBO_JOBS=1`/`GOPURS_PREPARE_JOBS=1` (sortie identique,
     validée sur b8x) pour rester déterministe. À élucider avec le pool.

## Méthode

- Parité byte-exacte (b8x + fixtures) ; profils d'allocations
  (`GOPURS_ALLOC_PROFILE`) **et** CPU (`PPROF=1` → `cpu.prof`) après chaque
  changement.
- Temps appariés, machine au repos ; ne pas superposer builds et campagnes ;
  table `altbak.pub` = **backend seul** (hors bootstrap et hors frontend purs).
- Piège : jamais de structure polymorphe accumulée (rebox quadratique —
  incidents `GoImports`, imprimante à 6 297 Go) ; profiler, pas seulement
  comparer la parité.
- **Seuil d'effort** : ne traiter que les leviers ≥ ~2-3 s CPU ou à effet
  murale mesuré ; documenter puis écarter les micro-gains.

## Références

- Campagnes, profils et rapport : `scratch/b8x-pbo-visibility-20260926/`
  (`mapnative/rapport.md`, `printer2/`, `jsbackend/`, `verify/`).
- Code : `gopurs/gopurs/src/Gopurs/`,
  `purescript-backend-optimizer-gopurs/`,
  `gopurs/gopurs-ordered-collections/`.
