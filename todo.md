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

**Compilateur (b8x, charge réelle, 2 655 modules).** Chargement TAST **2,29 s**
contre 9,64 s publiés (−76 %). Backend : 157,6 s (compilateur d'entrée de
session) → **143,2 s** avec la `Map` et le scanner natifs (−9,1 %), puis
**129,6 s** à `GOGC=300` (−17,8 %) et **113,6 s** à `GOGC=600` (−27,9 %).

**Protocole backend.** Lancer depuis la racine `b8x` **sans `--main`** : les
`modulePath` du TAST sont relatifs et les FFI ne sont résolues que depuis ce
répertoire. Les mesures faites en `cwd=run/bak/go` avec `--main Main` tournaient
sur une charge dégradée (FFI en stubs) et ne valent pas pour le vrai build ;
elles ont été refaites.

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

## Phase 3 — Passes PBO chaudes (profilage fait, gains mesurés)

- [x] **Profilage** (`PPROF=1`, charge réelle) : backend ~154 s dont
      `optimize + emit` ~103 s ; **GC ≈ 57 % du CPU** ; 312 Go alloués sur le
      run. Top : `Data.Map` (`(*Node).clone` 32,6 Go = 10,4 %, `find` et
      comparateurs ~16 Go), concaténation de tableaux 10,6 Go, `RecordDict*`
      ~20 Go, rebox ~24 Go, `Monomorphize` ~30 Go, `Semantics.quote` 19 % cum,
      `emitModule` 33 % cum (dont `referencedImports` 26,6 Go = 8,5 %).
      **Aucune passe unique ne domine** : le coût est le volume d'allocations.
- [x] **`Data.Map` `insertClone`** (~15 lignes natives) : **157,6 → 153,4 s
      (−2,7 %)** sur la charge réelle, RSS stable.
- [x] **Scanner d'imports natif** (`Gopurs.GoCode.referencedImports`) :
      **sortie Go byte-identique** (A/B sur 2 655 modules, 2 960 fichiers) et
      backend **153,4 → 143,2 s (−6,6 %)** ; cumulé avec la `Map` : **−9,1 %**.
- [x] **Balayage GC (charge réelle)** : 143,2 s (100) → 129,6 s (300, −9,5 %)
      → 113,6 s (600, −20,7 %) ; RSS 5,1 → 8,8 → 13,9 Gio. Décider du défaut
      `GOGC` du lanceur (politique mémoire).
- [ ] **Suite** : comparateurs de `Map` (FFI Value-native ou comparateur natif
      `String`/`Ident`), sites `<>` en boucle, rebox PBO, reste du chemin
      d'émission (`printGoExpr` 5,9 Go, `toCharArray`) — chaque correctif :
      sortie Go byte-identique + mesure appariée sur la charge réelle.

## Phase 4 — Généralisation (le bout de la logique)

- [x] **Hissage des dictionnaires clos (première passe codegen généralisée)** :
      `Gopurs.ClosedDictionaries` hie les constructions de dictionnaires de
      classe closes (sans capture locale, sans nœuds d'effet, type de classe
      annoté) en getters `sync.Once` de niveau module, uniquement sous les
      corps de liaisons déjà cachés. JSON général, campagne appariée :
      décodage **−14,5 %**, allocations **−9,2 %**, total **−5,4 %**, parsing
      inchangé ; plans de records reconstruits **91 368 → 18** (copies
      instrumentées) ; le module de test ne gagne qu'un getter (la chaîne
      réellement reconstruite à chaque variante d'événement). TAST : neutre
      (chemin déjà natif, allocations identiques au byte). 13 fixtures
      compilées/exécutées OK ; le seul snapshot modifié (`DerivingFunctor.go`)
      était déjà obsolète depuis le 21/09. Bootstrap du compilateur (459
      modules) OK. Rapport `2026-09-23-closed-dictionary-caching.md`.
- [x] **Déspecialisation des forwarders FFI (fusion, premier morceau)** :
      PBO spécialisait les fins wrappers FFI (`caseJson*`) en forçant un retour
      ADT natif que le wrapper re-boxait (aller-retour boxed→native→boxed,
      2 allocations par appel). `Gopurs.Monomorphization` détecte désormais
      **mécaniquement** les forwarders (lambda eta-expansé appliquant un import
      étranger du même module à ses paramètres, à travers `unsafeCoerce`) et
      les exclut de la spécialisation. JSON : décodage **−17,6 %** (apparié vs
      état précédent), allocations décodage **15,04 → 13,49 Mo (−10,3 %)** ;
      cumulé depuis la baseline : allocations **−18,5 %**, décodage −19,1 %.
      Code généré identique au byte à l'exclusion nommée sur le corpus (562
      spécialisations), oracle JSON exact, 13 fixtures OK. Rapport
      `2026-09-24-accessor-despecialization.md`.
- [ ] **Constructeurs `Right`/`Just`** (~24 % des allocations restantes
      mesurées) : relèvent de l'ABI native (représentation), pas d'un filtre de
      spécialisation.
- [ ] **Suite du hissage** : dictionnaires partiellement appliqués ; mesurer sur
      une charge réelle (b8x ou gopurs-aff) et vérifier les chemins d'erreur.
      Variante élargie (`LitRecord`, `CtorSaturated`) **testée et rejetée** :
      568 getters au lieu de 19, −1,3 % d'allocations seulement, +1,4 % de
      temps au décodage (campagne appariée) → forme étroite conservée.
- [ ] **Prochain gros poste (profil à jour, 500 passes)** : constructeurs
      `Just`/`Right` **20,8 %** des allocations de décodage, adaptateurs
      `Foreign.Object.lookup` **10,1 %**, `decodeForeignObject` 7,0 % (dont
      l'essentiel est la construction des maps `FO.Object`), plan de records
      6,6 %, accesseurs `caseJson*` 7,3 % (encodage CPS qui re-construit des
      `Maybe`/`Either`). Cible recommandée : fusion codegen des résultats
      immédiatement déstructurés (case-of-known-constructor) avant les FFI
      natives ponctuelles.
- [ ] **Politique GC (levier mesuré)** : même binaire, runs intercalés —
      `GOGC=100→300` à 1 P : cycles GC 196→42, décodage **−26,9 %**, total
      −23,3 % ; à 4 P (forme applicative) : **−5,9 %** décodage, −9,6 % total ;
      `GOGC=600` ≈ −28 % à 1 P. Le protocole de campagne reste à `GOGC=100`
      (comparabilité JS/C) ; le défaut des applications générées est une
      décision mémoire (RSS) non tranchée. Détail :
      `2026-09-23-json-general-profiling.md`.
- [ ] **Plan de records — décision de contrat** : précalculer les noms et la
      disposition à la construction du plan éviterait 2 évaluations de symbole
      par champ et par enregistrement, mais les tests figent la fidélité à
      `Record.insert` (double évaluation, disposition recalculée). Gain estimé
      ~2-4 % du décodage ; à trancher avant de toucher au contrat Go/JS.
- [ ] Trancher entre **ABI native** (arguments et résultats typés de bout en
      bout, point 12 du plan historique) et **codegen ciblé** (émettre du
      contrôle direct pour les schémas monadiques connus).
- [ ] Prototyper sur un chemin chaud complet (décodeur de champ → worker
      partagé → stockage), mesurer le potentiel, puis généraliser.
- [ ] Critère de succès : le gain apparaît **sans portage manuel** — c'est le
      compilateur qui le produit.

## Phase 5 — JSON général

- [x] **Référence C mesurée** (rapport altbak
      `2026-09-23-native-c-references.md`, lignes ajoutées à la table C) :
      JSON Decoding **644,96 µs** (simdjson + arène) contre Go 14,67 ms et
      JS 8,94 ms → **×22,7** de marge côté natif (décodage seul ×41,9,
      parsing ×6,3). L'écart est dominé par l'allocation et le GC, pas par le
      parseur. Array Indexing : sur le même noyau, Go (4,21 ms) est 1,6× plus
      rapide que `clang -O3` (6,59 ms) ; la représentation boxed coûte ≤2 % ;
      une reformulation en blocs (vectorisable) descend à ~0,51 ms → le levier
      restant est algorithmique.
- [x] **Baseline et profils par phase (23 sept.)** : baseline re-mesurée
      (`var/benchmark/json-dec-baseline{-results}-20260923`, Go seul) —
      parsing **2 366,2** / décodage **11 232,6** / total **14 294,3 µs**,
      allocations identiques à la campagne publiée (4 612 800 / 16 556 992 /
      21 169 776 B). Le profileur CPU `runtime/pprof` est inexploitable sur
      cette machine (échantillons attribués aux threads `kevent`/`madvise`
      alors que `sys` réel ≈ 0,01 s) → lecture par `sample(1)` + profils
      d'allocation ; rapport `2026-09-23-json-general-profiling.md`. GC ≈ 14 %
      du décodage (`gctrace`, GOGC=100). Postes locaux mesurés : reconstruction
      du plan de records **2 393 appels/pass ≈ 645 µs (6,5 %)** — dictionnaires
      et closures neufs à chaque appel, donc pas de cache possible côté
      runtime ; setup `decodeArray` ≈ 0,6 % ; setup `decodeForeignObject`
      ≈ 1,6 %. Constructeurs `Right`/`Just` ≈ 19,5 % des allocations,
      adaptateurs `Foreign.Object.lookup` ≈ 9 %.
- [ ] **Cible prioritaire** : hoisting codegen des constructions de
      dictionnaires constantes dans les lambdas (le plan de records est
      reconstruit à chaque enregistrement décodé) ; puis chemin `lookup`
      Value-level. Le chemin `DecodeJson` générique reste à ×1,6 du JS
      (15,11 / 9,28 ms) : le levier restant est la construction de
      dictionnaires et la représentation boxed, pas les algorithmes de
      décodage.
- [x] **Référence C pour JSON to Typed AST** (`typed-ast.cc`, simdjson + arène,
      ~1 100 lignes) : les **12 empreintes** du corpus sont reproduites à
      l'identique (validées d'abord textuellement contre la PS, puis par
      l'oracle figé dans chaque processus). Mesure officielle : C
      **4,26 / 4,52 / 8,76 ms** contre Go 27,46 / 15,38 / 56,99 ms et JS
      20,36 / 61,99 / 80,35 ms → **×6,5** au total (×6,4 au parsing, ×3,4 au
      décodage). Seule la passe **pure** de validation d'usage n'est pas
      reproduite : mesurée sur le même build Go, elle pèse **872,8 µs sur
      13 206,6 µs (6,6 %)** → une référence C équivalente serait ~4,8 ms, soit
      encore ×3,2. Le principal écart restant est le **parsing** (×6,4) et non
      le décodeur, déjà natif des deux côtés.

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
