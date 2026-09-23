# Gopurs — plan : du décodage natif à la généralisation

## État mesuré (23 septembre 2026)

Le décodage TAST est intégralement natif en Go derrière la FFI existante
(cinq `foreign import` dans `CoreFn/{Json,TypeTable,Usage}.purs`, repli
PureScript pour le JS). Cellules officielles du diagnostic TAST :

| Cellule | Go | JS | Ratio |
|---|---:|---:|---:|
| Décodage | **15,41 ms** | 59,54 ms | **×0,26** |
| Total | **48,00 ms** | 78,19 ms | **×0,61** |

Campagne cumulée appariée (5 paires, workspaces régénérés, oracles dans chaque
processus) : décodage **336,61 → 15,16 ms (−95,50 %)**, allocations
**402 378 424 → 27 106 040 octets (−93,26 %)**, total **386,41 → 48,13 ms
(−87,54 %)** ; parsing inchangé. Détail, provenance et limites :
[2026-09-23-tast-native-decoding.md](../../altbak.pub-gopurs/docs/benchmark-results/2026-09-23-tast-native-decoding.md).

**Lecture honnête.** Ces chiffres opposent du Go *écrit à la main* à du JS
*généré*. Le levier démontré est « contrôle direct + moins de transitoires »,
pas « le compilateur Go bat le compilateur JS ». La logique n'est pas terminée :
voir les phases 3–5.

## Phase 1 — Finir le décodeur (borné, rapide)

- [x] **Différentiel des chemins d'erreur** : **2 210 mutations sur 6 graines
      → 1 493 cas d'erreur + 717 décodages réussis**, messages et empreintes
      identiques au PureScript, testé sur la sortie intégrée du compilateur. Il
      a fait apparaître 5 divergences réelles (messages `Object`/`Array`,
      wrapping `AtKey` de `getFieldOptional'`, `AtIndex` sur
      `decodeModuleName`, message `Array` de la table de types, règle
      d'emballage par clé de `decodeReExports`), toutes corrigées.
- [x] **Transitoires réduits** : entrée native de la table de types (plus
      d'aller-retour `Box`), suppression de la closure par appel dans `cndExpr`,
      scopes persistants (chaîne) dans la validation d'usage.
      Allocations décodage : **28 724 952 → 27 106 040 B (−5,6 %)** ; campagne
      cumulée appariée : décodage **−95,50 %**, total **−87,54 %**, allocations
      **−93,26 %** ; cellules officielles **15,41 / 48,00 ms**.
- [x] **Couplages aux noms générés documentés** (`Call_…_decodeSourceSpan`,
      `Call_Data_Map_Internal_fromFoldable__…`) : commentaires explicites aux
      deux sites d’appel (chemin froid, échec bruyant à la compilation si le
      nom change). Un remplacement stable est possible mais différé.

## Phase 2 — Effet compilateur (à trancher)

- [ ] Reprendre b8x quand la question murale sera prioritaire : mesure appariée
      (`b -c` avant/après, phases, RSS, CPU) et `gopurs-aff` (238 modules)
      comme corpus de contrôle. Le décodeur n'est qu'une partie du chargement
      TAST.

## Phase 3 — Passes PBO chaudes (profilage fait, correctifs en cours)

- [x] **Profilage du backend b8x** (`PPROF=1`, 237 modules) : backend 139,3 s
      dont `optimize + emit` 103,1 s. **GC ≈ 57 % du CPU** (`gcBgMarkWorker`
      244 s sur 430 s d'échantillons, `mProf` du profilage inclus) ; le CPU
      applicatif est dispersé. Profil d'allocations (312 Go sur le run) :
      `Data.Map` (`(*Node).clone` 32,6 Go = 10,4 %, `(*Node).find` et
      comparateurs ~16 Go), concaténation de tableaux 10,6 Go, `RecordDict*`
      ~20 Go, rebox PBO ~24 Go, `Monomorphize` ~30 Go, `Semantics.quote` 19 %
      cum, `emitModule` 33 % cum (codegen + imprimante + `referencedImports`).
      Conclusion : **aucune passe unique ne domine** — le coût est le volume
      d'allocations et le GC, réparti entre Map, représentations/rebox,
      tableaux et émission.
- [x] **Premier correctif ciblé** : `Data.Map` `insertClone` (capacité `len+1`
      au lieu de `maxDegree` sur le chemin d'insertion) → backend
      **−6,5 % / −14,0 %** sur deux paires appariées (médianes 148,6 → 133,3 s),
      RSS légèrement en baisse ; oracles altbak validés après reconstruction.
- [x] **Balayage GC** (backend, `GOGC`) : 140,2 s (100) → 125,2 s (300,
      −10,7 %) → 117,7 s (600, −16,1 %), RSS 5,3 → 9,0 → 14,9 Gio. Un défaut
      `GOGC` dans le lanceur est une décision de politique (mémoire) à trancher.
- [ ] **Suite** : traiter les postes identifiés — comparateurs de `Map`
      (FFI Value-native, ou comparateur natif pour les clés `String`/`Ident`),
      concaténations de tableaux (sites `<>` en boucle), rebox PBO, chemin
      d'émission (`referencedImports`, `printGoExpr`, `toCharArray`). Chaque
      correctif : mesure appariée sur le backend + oracles altbak.

## Phase 4 — Généralisation (le bout de la logique)

- [ ] Trancher entre **ABI native** (arguments et résultats typés de bout en
      bout, point 12 du plan historique) et **codegen ciblé** (émettre du
      contrôle direct pour les schémas monadiques connus).
- [ ] Prototyper sur un chemin chaud complet (décodeur de champ → worker
      partagé → stockage), mesurer le potentiel, puis généraliser.
- [ ] Critère de succès : le gain apparaît **sans portage manuel** — c'est le
      compilateur qui le produit.

## Phase 5 — JSON général

- [ ] Le chemin `DecodeJson` générique reste à ×1,6 du JS (15,11 / 9,28 ms).
      Appliquer la méthode aux primitives de décodage plutôt qu'aux instances,
      ou assumer l'écart.

## Décisions ouvertes

- **Comparaison** : « Go manuel vs JS généré ». Si une comparaison symétrique
  est voulue, porter aussi un chemin JS natif pour mesurer l'écart réel des
  compilateurs.
- **Maintenance** du Go manuel dans PBO (taille, revue, couplage aux noms
  générés et aux spécialisations).
- **Chemins d'erreur** : non couverts par les 12 modules valides du corpus ;
  seul le différentiel de la table de types les exerce.

## Références

- Rapport : [2026-09-23-tast-native-decoding.md](../../altbak.pub-gopurs/docs/benchmark-results/2026-09-23-tast-native-decoding.md)
- Sources : `purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/CoreFn/{Json,TypeTable,Usage}.{purs,go,js}`
- Workspaces et campagnes : `altbak.pub-gopurs/var/benchmark/json-tast-native-{tt,arr,usage2,ann,dec,dec5}-20260923`,
  `native-*-campaign-20260923`, `native-dec-cumulative-20260923`
- Différentiels, tests et variantes : `scratch/tast-revolution-20260923/`
- Protocole : `altbak.pub-gopurs/bin/benchmark/json-diagnostic.py`
  (GOMAXPROCS=1, GOGC=100, PGO désactivé, médiane des minima de processus) ;
  campagnes appariées : `scratch/tast-decode-20260923/campaign.py`.
