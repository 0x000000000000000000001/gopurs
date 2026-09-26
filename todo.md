# Gopurs — PBO : réduction des allocations et finitions (26 septembre 2026)

Ce fichier remplace le journal de la session précédente (sauvegardé dans
`scratch/b8x-pbo-visibility-20260926/todo-journal-20260926.md`).

## État acquis

- **Parité byte-exacte** séquentiel ↔ parallèle sur b8x (2 987 fichiers Go),
  vérifiée pour `jobs=1/2/4/8/64` et à chaque passe de validation.
- **Backend : 136 s (référence corrigée) → 68,7 s stable** (`jobs=8` +
  `GOGC=off` + `GOMEMLIMIT=10GiB` ; 4 passages entre 67,6 et 69,4 s).
- **`b -c` de bout en bout sur b8x** : 146,5 s au total, dont **backend
  69,9 s** (chargement 1,3 + préparation 26,8 + optim+émission 41,7), le
  reste étant le bootstrap du compilateur (~36,5 s), le frontend PureScript
  et `go mod tidy`.
- **Défauts du lanceur** (`gopurs/bin/gopurs`) : sur une machine ≥ 32 Go,
  `GOGC=off`, `GOMEMLIMIT=10GiB` et `GOPURS_PBO_JOBS=8` si l'utilisateur n'a
  rien défini ; réglages explicites respectés ; petites machines inchangées.
- **Tests** : émission 10/10 ; suite PBO 140/140 ; tests natifs du `Map`
  (ABI comparateur) passants.
- **Correctif d'affichage** : la progression est imprimée dans
  `onCodegenModule` (un appel par module finalisé, ordre canonique) — **à
  vérifier au prochain build du compilateur**.
- Profils, rapports et campagnes : `scratch/b8x-pbo-visibility-20260926/`
  (`rapport.md`, `pbo-parallel-rapport.md`, `verify/`, `cont/`, `e2e/`).

## Mesures d'allocations actuelles (pprof, binaire `751d2ab4`)

| Famille | Séquentiel | 8 workers | Nature |
|---|---:|---:|---|
| **`Rebox_*`** (647 fonctions) | **41,3 Go** | **70,3 Go** | coercions de dictionnaires (codegen) |
| Dictionnaires `RecordDict*` | 18,9 Go | 24,1 Go | dictionnaires construits à l'exécution |
| Callbacks `Map` (insert/lookup/union) | 23,0 Go (~50 cumulés) | 33,3 Go | pont FFI, boxage par comparaison |
| `insert`/`clone` B-tree | 17,6 Go | 25,2 Go | copies de chemins des `Map` |
| `Array.bind` (`concatMap`) | ~24,5 Go cumulés | — | concaténations de tableaux |
| PBO Semantics | 21,4 Go | 35,0 Go | optimiseur |
| `mangleType` | 8,6 à plat / 13,5 cumulés | 8,8 Go | préparation |
| Imprimante + chaînes | ~11 Go | ~11 Go | impression Go |
| `Apply` / runtime | ~21 Go | ~27 Go | application générique, boxage |
| **Total** | **214,2 Go** | **300,7 Go** | +40 % en parallèle (relances) |

Les constructeurs `Maybe`/`Either` sont marginaux ici (~0,7 Go), contrairement
au décodage JSON d'altbak.

> Après les deux passes d'optimisation (rebox identité → cast, comparaison
> native `EvalRef`) : total **204,7 Go** séquentiel et backend **64,2 s**
> (8 workers, 63,8-64,8 s sur 3 passages).

## Enseignements

1. **L'hypothèse « perfs du code compilé gopurs » est confirmée** : le
   reboxing et les dictionnaires pèsent ~60 Go en séquentiel et ~94 Go en
   parallèle, la même famille que les gains d'altbak (dictionnaires clos,
   déspecialisation des forwarders). L'estimation antérieure « reboxing
   ~10 Go » était très sous-évaluée.
2. **Nouvelle cible** : les `bind` de tableaux (`concatMap`), même classe que
   les folds d'imports déjà corrigés.
3. Les **callbacks `Map`** restent ~50 Go cumulés : le pont boxe encore chaque
   comparaison malgré la fusion `compareInt`.
4. Le mode 8 workers **alloue +40 %** (3 508 tentatives contre 2 683) :
   réduire les relances améliore temps et mémoire.

## Plan

### 1 — Reboxing et dictionnaires (le plus gros poste)

- [x] **Diagnostic (26/09)** : le reboxage est partout (tous les modules PBO) et
      vient d'instanciations distinctes de structs génériques dont les
      paramètres de type sont **fantômes** (aucun champ ne les utilise). Sur le
      Go généré de b8x : 12 764 des 20 579 rebox sont des copies identiques
      champ à champ (structs à 1 champ le plus souvent).
- [x] **Passe codegen** : `renderReboxFunction` émet un cast
      `(*Dest)(unsafe.Pointer(in))` quand toutes les affectations sont
      identiques (champs immuables, `Rc` jamais lu dans le Go généré).
      Validation : diff exhaustif de la sortie b8x — **12 764 conversions,
      0 autre changement** (déclarations et imports compris), 4 fixtures
      exécutées OK, `b -c` OK.
      Mesures : rebox **41,3 → 31,4 Go**, total **214,2 → 208,8 Go**,
      backend **69,9 → 64,3 s** (`b -c`), **68,8 → 64,7 s** (8 workers),
      **101,8 → 87,6 s** (séquentiel). Parité séquentiel ↔ parallèle
      re-vérifiée (0 divergence) sur le nouveau compilateur.
- [x] **Deuxième passe (26/09) — comparaison native des clés `EvalRef`** :
      `Ord EvalRef`/`Eq EvalRef` comparaient `Qualified Ident` via l'instance
      polymorphe `ordQualified`, compilée déspecialisée en `Value` par le
      backend : chaque comparaison construisait un dictionnaire
      `Ord (Qualified Value)` et reboxait les deux idents (2,46 Go, le plus
      gros rebox restant à lui seul). Les instances sont maintenant manuelles
      et monomorphes (`compareQualifiedIdent`/`eqQualifiedIdent` dans
      `CoreFn`, `compareMaybeIdents`/`compareLevels` dans `Semantics`, FFI
      natives `compareStringImpl`/`compareIntImpl` dans `FfiSupport`).
      Validation : sortie b8x **identique au bit** (0 fichier modifié), tests
      PBO 140/140, symbole chaud disparu du profil.
      Mesures : total **208,8 → 204,7 Go**, séquentiel **87,6 → 86,8 s**,
      8 workers **64,7 → 64,2 s** (63,8-64,8 sur 3 passages).
- [ ] **Limite identifiée — le pont FFI boxe** : toute FFI PS↔Go passe par un
      wrapper `_Gopurs_*` `Value` (unbox/box par appel). Une passe de
      *native call lowering* pour les FFI connues (arguments concrets) retirerait
      ces boîtes — 2 par comparaison de chaîne ici, mais le levier est général
      (hashString, comparateurs, etc.).
- [ ] Rebox restants (29,3 Go) : conversions réelles, par ex.
      `Tuple[string, Value]` → `Tuple[Value, Value]` et boxage d'arrays
      élément par élément (`Value{… UnsafePtr: Rebox_…}`). Pistes : construire
      directement le type cible dans les coercions d'arrays (éviter box/unbox
      par élément), mutualiser les paires identiques entre modules,
      canonicaliser les instanciations à paramètres fantômes.
- [ ] Dictionnaires `RecordDict*` (19,3 Go) : identifier les constructions
      répétées et les hisser.
- [x] Objectif indicatif : −30 Go séquentiel → **atteint** (214,2 → 204,7 Go
      avec les deux passes, dont −9,9 Go de rebox identité sur le compilateur
      et −2,1 Go de rebox `Qualified` ; les 12 764 call sites du code
      utilisateur neuf n'allouent plus).

### 2 — `Array.bind` / `concatMap` (~24,5 Go cumulés)

- [ ] Localiser les sites via `-peek 'Control_Bind_ArrayBind'` (appelants par
      fonction).
- [ ] Remplacer les accumulations quadratiques par une accumulation linéaire
      (liste) ou une concaténation native, comme `concatStringArrays` dans
      `Gopurs.GoImports`.
- [ ] Vérifier l'absence de tempête de reboxage (leçon de l'incident
      `GoImports` : toute modification du code gopurs doit être validée par un
      profil d'allocations, pas seulement par la parité).

### 3 — Callbacks `Map` (~50 Go cumulés)

- [ ] Documenter le coût résiduel : boxage des clés par comparaison dans le
      pont (`InsertImpl`/`LookupImpl`/`UnionWithImpl`).
- [ ] Étudier une spécialisation native par type de clé (chemins dédiés pour
      `String`, `Qualified Ident`, `EvalRef`) évitant la fonction PS.
- [ ] Mesurer sur le corpus et vérifier la parité.

### 4 — `mangleType` (13,5 Go cumulés)

- [ ] Mémoïsation bornée par identité aux sites `collectExpr` et
      `specializationKey` (réutiliser `BoundedMemo`/`SameIdentity`).
- [ ] Vérifier que les types réutilisés sont partagés (sinon le mémo ne sert
      à rien) et mesurer taux de succès.

### 5 — B-tree `insert` (17,6 / 25,2 Go)

- [ ] Réduire le nombre d'insertions dans les chemins chauds (accumulateurs,
      `Map.insertWith`/`alter` évités, structures dédiées).
- [ ] Option : étudier le degré du B-tree (compromis allocations)
      comparaisons) — expérience d'une ligne, à mesurer.

### 6 — Relances du builder parallèle (+40 % d'allocations)

- [ ] Mesurer les tentatives rejetées (825-850 à 8 workers) et leurs causes
      (`pendingDeps` par module).
- [ ] Renforcer les indices d'ordonnancement (profondeur 2 des références,
      imports utiles) pour éviter les tentatives perdues.
- [ ] Si les rejets restent nombreux, envisager une reprise à granularité
      plus fine (par groupes de bindings) plutôt qu'une reconversion entière.

### 7 — Imprimante et chaînes (~11 Go)

- [ ] `printGoExpr`/`printGoFile` : pré-dimensionner les `strings.Builder`
      ou réutiliser des buffers ; mesurer `MakeNoZero`.

### 8 — Finitions

- [ ] Vérifier l'affichage de progression corrigé au prochain build
      (numéros croissants, sans doublons).
- [ ] Documenter les variables du lanceur (`GOGC`, `GOMEMLIMIT`,
      `GOPURS_PBO_JOBS`, `GOPURS_EMIT_JOBS`, `GOPURS_PREPARE_JOBS`,
      `GOPURS_PIPELINE`, `GOPURS_ALLOC_PROFILE`).
- [ ] Envisager un contrôle de non-régression d'allocations (profil pprof
      sur un corpus réduit) dans la validation.
- [ ] Réévaluer le bootstrap du compilateur (~36 s dans `b -c`) : seule
      partie incompressible du flux `-c`.

## Méthode de validation

- **Parité byte-exacte** : manifester les sorties dans des répertoires
  propres, comparer séquentiel ↔ parallèle à chaque changement.
- **Allocations** : profil `GOPURS_ALLOC_PROFILE` avant/après, analyse par
  familles (`scratch/b8x-pbo-visibility-20260926/verify/analyze.py`).
- **Tests** : `node --test tools/emission.test.mjs` (émission) et
  `node --test test/*.mjs` (PBO, 140 tests), tests natifs du `Map`.
- **Temps** : machine au repos, répétitions (médiane), ne pas superposer
  builds et campagnes ; noter la charge système.
- **Piège connu** : une réécriture en `List` de folds gopurs a produit
  +240 Go de reboxage (`Rebox_Gopurs_GoImports_*`). Toujours mesurer les
  allocations après une modification du code gopurs lui-même.

## Références

- Rapports : `scratch/b8x-pbo-visibility-20260926/rapport.md`,
  `pbo-parallel-rapport.md`.
- Profils et analyseur : `scratch/b8x-pbo-visibility-20260926/verify/`.
- Dernières campagnes : `cont/` (dispatch continu), `e2e/b-c.log`.
- Code : `purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/`
  (`Builder.purs`, `Convert.purs`, `Semantics.purs`, `Cache.*`),
  `gopurs/gopurs/src/Gopurs/` (`Emission.purs`, `GoImports.*`, …),
  `gopurs/gopurs-ordered-collections/src/Data/Map/Internal.*`.
