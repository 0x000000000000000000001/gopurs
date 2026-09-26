# Gopurs — parallélisme déterministe de PBO et accélération de b8x

Date : 26 septembre 2026.

Les chemins de sources ci-dessous sont relatifs à `htdocs/`.

> **Journal (26 septembre, soirée).**
>
> - **Parité byte-exacte atteinte** : binaire `1860e443` (visibilité par rang,
>   relances, directives accumulées + propriété testée). Sur b8x, `jobs=1/2/4/8`
>   produisent les mêmes 2 987 fichiers Go (0 ajouté, 0 retiré, 0 modifié),
>   vérifié aussi pour `jobs=64`.
> - **Gains mesurés** (passages individuels) : `jobs=1` 144,4 s de backend
>   (105,4 s d'optim+émission) ; `jobs=8` 133,9 s (100,1 s) ; `jobs=64` 128,7 s
>   (95,5 s). Tentatives : 2 683 / 3 269 (+22 %) / 4 194 (+56 %). Le plateau
>   vient du chemin critique de conversion, pas du nombre de workers.
> - **Correction notable** : `foldlWithIndex` de `Data.Map` passe
>   `(clé, accumulateur, élément)` ; deux types identiques avaient masqué
>   l'inversion. Remplacé par `foldrWithIndex`, avec test de propriété
>   `test/directives-removal.mjs` (300 cas aléatoires × tous les indices).
> - **Comparateurs manuels** (`Qualified`, `EvalRef`) : évite le boxage
>   d'`ordMaybe` ; ordre vérifié identique (400 cas) ; **−11 % partout**
>   (séquentiel 128,3 s ; `jobs=8` 120,5 s ; `jobs=8` + `GOGC=600` 100,4 s),
>   parité 0 divergence.
> - **Profil d'allocations pprof** ajouté au compilateur natif
>   (`GOPURS_ALLOC_PROFILE`) : **242,9 Go** au total. Postes : callbacks du
>   `Map` natif ~53 Go (cumulé), `insertClone` 16,9 Go, dictionnaires ~19 Go,
>   `mangleType` 13,2 Go, reboxing ~10 Go, `printGoExpr` ~7 Go.
> - **`Map` natif** (`gopurs-ordered-collections`) : les deux callbacks
>   `compare` + `fromOrdering` fusionnés en un comparateur natif
>   `k -> k -> Int` pour les sept opérations — une frontière FFI et un boxage
>   en moins par comparaison. Validation parité/temps en cours.
> - **Bilan campagne finale** (binaire `5d1c88fe`, parité 0 divergence partout) :
>   `jobs=1` 116,4 s ; `2` 119,6 s ; `4` 113,7 s ; `8` 110,4 s ;
>   `8 + GOGC=600` **91,7 s** (RSS 14,2 Go). Cumulé sur la journée :
>   séquentiel 144,4 → 116,4 s, `jobs=8` 133,9 → 110,4 s, meilleure config
>   115,6 → 91,7 s. Référence corrigée du matin : 136 s → **91,7 s (−33 %)**.
> - **Postes d'allocation restants** (profil 231 Go) : `insertClone` 16,9 Go,
>   callbacks `Map` ~26 Go cumulés, dictionnaires `RecordDict*` ~19 Go,
>   `mangleType` 13,2 Go, reboxing ~10 Go, `printGoExpr` ~7 Go.
> - **`GOMEMLIMIT`** (parité 0 divergence) : `GOGC=off` + limites 8/10/12 GiB →
>   97,9 / **95,6** / 95,5 s, pics 10,0 / 12,0 / 14,5 Go. Genou à 10 GiB ; le
>   pic dépasse la limite d'environ 25 %. Politique recommandée : 8 GiB (CI) ou
>   10 GiB (station), au lieu de `GOGC=600` non borné (91,7 s, 14,2 Go).
> - **Profil du réglage cible** (`jobs=8`, `GOGC=off`, `GOMEMLIMIT=10GiB`) :
>   70 % des échantillons actifs en GC/allocateur, émission ~0,8 % — le mur
>   reste la mémoire, pas l'ordonnanceur.
> - **B-tree** : le chemin « clé trouvée » ne copie plus la liste d'enfants
>   (nœuds immuables) ; tests natifs du Map réécrits pour la nouvelle ABI et
>   passants. Mesure d'allocations : **neutre** (`insert` 17,31 vs 17,33 Go) ;
>   parité 0 divergence. Les temps de la campagne B-tree ont tourné pendant un
>   `rustc` extérieur (load > 20) et ne sont pas exploitables.
> - **Piège documenté — reboxage** : une première réécriture de
>   `Gopurs.GoImports` (accumulation en `List`) a produit **+240 Go
>   d'allocations** (`Rebox_Gopurs_GoImports_*`) : les dictionnaires des folds
>   sont recoercés par élément. Changement reverté. Leçon : toute modification
>   du code gopurs lui-même doit être validée par un profil d'allocations, pas
>   seulement par la parité.
> - **Imports (2e tentative, retenue)** : concaténation native
>   `Array (Array String) -> Array String` (FFI Go + JS) à la place des folds
>   `foldMap` quadratiques de `collectImports`, puis des 31 folds récursifs de
>   `declImports`/`typeImports`/`exprImports`. Parité **0 divergence** ;
>   allocations **236,7 → 223,9 Go** (−5,4 %) ; aucun reboxage ;
>   `init.func230` 15,7 → 2,7 Go cumulés. Temps à confirmer machine au repos
>   (campagnes `rustc` extérieures à répétition).
> - **Reste à cibler** : `foldMapDefaultR` résiduel 8,5 Go cumulés (folds PBO
>   `TypeSubstitution`, `mapAccumL`), mémo `mangleType` 13,3 Go,
>   dictionnaires/reboxing ~19 Go, `insert` 17,3 Go, callbacks `Map` ~26 Go
>   cumulés, `printGoExpr` 7,2 Go.
> - **Temps finaux** (binaire `38dfd114`, parité 0 divergence partout) :
>   `jobs=8` + `GOGC=off` + `GOMEMLIMIT=10GiB` → **89,4 / 92,7 s** (moyenne
>   91,0 s, optim+émission 62-65 s, CPU user ~210 s) ; `jobs=8` par défaut
>   106-126 s (variance machine) ; référence corrigée du matin 136 s →
>   **−33 %**. Allocations cumulées : 274 Go (profil initial) → 223,9 Go.
> - **Politique recommandée** : `GOGC=off` + `GOMEMLIMIT=10GiB` (pic ~12 Go),
>   ou 8 GiB (pic ~10 Go) sur CI ; activer le parallèle par défaut après
>   stabilisation de l'ordonnanceur.
> - **Différé, non fait** : reprise des conversions rejetées ; degré du
>   B-tree ; réduction des dictionnaires et du reboxing côté codegen.
> - **Suite à trancher selon les compteurs** : réduire le coût des relances
>   (reprise à granularité fine) ou alléger la coordination/allocations.

> Le contenu précédent de ce fichier (phases décodeur TAST / JSON général)
> reste dans l'historique Git — dernier commit `92c55b9` touchant ce fichier.
> Les rapports associés sont dans `altbak.pub-gopurs/docs/benchmark-results/`.

## Objectif et critères de réussite

Accélérer le backend natif sur le corpus réel de b8x en rendant explicites les
informations partagées entre optimisations, puis en exploitant les cœurs
disponibles et en réduisant les allocations des chemins dominants.

- **Premier objectif : environ 80 s de backend**, contre le passage séquentiel
  corrigé à environ 136 s. **Environ 60 s serait un très bon résultat** ;
  45–65 s constitue une ambition ultérieure incluant la préparation et la mémoire.
- Ces temps sont des objectifs de travail, à confirmer par des mesures appariées.
- L'oracle de cette refonte est le **builder séquentiel avec accumulation correcte
  des directives**, conservé comme référence indépendante.
- Même corpus et mêmes options : couverture complète et Go byte-identique entre
  séquentiel et parallèle, pour 1/2/4/8 workers, puis compilation et exécution
  des tests pertinents.
- Les résultats de performance doivent préciser le CPU, les allocations, le RSS
  et la variabilité, en plus du temps mural.
- Le benchmark backend part de TAST existants. Le gain sur `b -c` sera mesuré
  séparément, avec le frontend PureScript et la compilation Go.

## Point de départ

Corpus historique : 2 683 modules TAST, 2 987 fichiers Go, environ 217,8 Mo de
TAST, M4 Pro à 14 cœurs. Relever à nouveau ces nombres et les empreintes au
démarrage de la campagne.

| Configuration historique | Optimisation + émission | Backend |
|---|---:|---:|
| Ancien binaire, passage de confirmation | 89,2 s | 120,4 s |
| Séquentiel corrigé | 103,3 s | 136,0 s |
| Prototype corrigé, 8 workers PBO | 72,4 s | 106,5 s |

Les deux derniers résultats sont des passages individuels. Le prototype
parallèle diverge encore dans `Data_Either.go`. La sortie séquentielle corrigée
diffère elle-même de l'ancien compilateur dans 945 fichiers ; sa validation à
l'exécution doit être complétée.

Le profil initial estime environ 274 Go d'allocations cumulées et 57 % du CPU
échantillonné dans le GC. Ce profil concerne l'ancien binaire et l'ensemble de
l'invocation ; il faudra le renouveler sur le candidat.

La simulation pondérée par les octets TAST et le graphe approximatif extrait du
Go généré ne permettent pas de déterminer le chemin critique réel de PBO.

## Principes retenus, inspirés de TypeScript

- Entrées partagées stables, état de travail privé, résultats immuables.
- Distinguer une référence à un module d'une dépendance envers un résultat
  d'optimisation de ce module.
- Déterminer la visibilité par une règle explicite, indépendamment de l'heure
  de fin des workers.
- Réutiliser les caches selon leur véritable contexte de validité.
- Employer l'affinité entre modules pour améliorer la localité une fois les
  dépendances et les coûts mesurés.

Références locales :

- `TypeScript/tsc/internal/compiler/checkerpool.go` : partitionnement équilibré,
  affinité, caches par checker, ordre de parcours stable.
- `TypeScript/tsc/internal/checker/checker.go` : calcul des informations à la
  demande dans l'état propre au checker.
- `TypeScript/tsc/internal/checker/utilities.go` : ordre des types et symboles.
- `TypeScript/tsc/internal/compiler/program.go` : émission parallèle et collecte
  ordonnée des résultats.
- `TypeScript/tsc/internal/core/{arena,linkstore}.go` : allocations groupées et
  tables d'informations associées aux nœuds.

## 0 — Figer une référence reproductible

- [ ] Relever les révisions, modifications locales, versions d'outils et SHA-256
      des sources et binaires utilisés, y compris le bootstrap et les FFI.
- [ ] Produire des artefacts distincts pour la référence séquentielle corrigée
      et le candidat, à partir d'états de sources figés pendant chaque bootstrap.
      Le script actuel `tools/build-native.mjs` publie le binaire installé :
      prévoir une destination de candidat explicite pour les campagnes.
- [ ] Fixer l'accumulation correcte des directives comme contrat de référence
      pendant les comparaisons ; vérifier le probe A/B/C et les tests applicatifs.
- [ ] Préparer des sorties de benchmark vierges, isolées des sorties usuelles de
      b8x, tout en préservant les chemins relatifs réels des FFI. Chaque invocation
      doit produire tous les fichiers attendus.
- [ ] Refaire trois passages de référence, sans campagnes concurrentes, avec
      manifeste TAST/FFI et manifeste des sorties complètes.
- [ ] Rectifier les conclusions trop fortes du rapport historique et des
      commentaires du prototype : impossibilité byte-exacte non démontrée,
      facteur d'accélération distinct de l'occupation des workers, temps PBO
      distinct de la phase combinée optimisation/émission.

Sources : `scratch/b8x-profile-20260925/`,
`gopurs/gopurs/tools/build-native.mjs` et
`purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/Builder.purs`.

## 1 — Mesurer le travail et les attentes réels

- [ ] Instrumenter la conversion de chaque module, la publication des résultats,
      la préparation du codegen, la génération des corps et l'impression/écriture.
- [ ] Relever les durées murales par module ; utiliser les profils/labels et
      traces natifs pour l'attribution CPU, le GC et la contention.
- [ ] Compter les tâches prêtes, actives et bloquées, les barrières, les lectures
      d'implémentations et de directives, les hits, absences et dépendances en attente.
- [ ] Identifier les lectures transitives introduites par l'inlining et les
      spécialisations, notamment le cas `Data.Either` / `Effect.applicativeEffect`.
- [ ] Construire un graphe des dépendances effectivement rencontrées, pondéré
      avec les temps observés ; identifier les modules lourds et les longues chaînes.
- [ ] Garder les métriques désactivables et comparer les performances sans
      instrumentation de diagnostic.

Livrable : une attribution du temps suffisante pour distinguer manque de travail
prêt, déséquilibre, publication, codegen et pression mémoire.

## 2 — Expliciter l'environnement d'optimisation

### Contrat de visibilité

Attribuer à chaque module son rang dans l'ordre canonique du builder séquentiel.
Pour le consommateur de rang `i`, après consultation de ses implémentations locales :

1. Un module externe de rang `j >= i` reste invisible, même si son calcul est fini.
2. Un module de rang `j < i` fournit son résultat final lorsqu'il est prêt.
3. Si ce prédécesseur n'est pas prêt, la lecture signale une dépendance en attente.
4. Un identifiant absent d'un résultat final constitue une absence définitive.
5. Un module extérieur au build suit le contrat de la référence ; les métadonnées
   d'un ancien build ne deviennent pas visibles par accident.

Cette règle permet de calculer un module tardif sans modifier les décisions des
modules qui le précèdent. Les attentes entre modules suivent des rangs strictement
décroissants ; leur profondeur réelle reste à mesurer.

- [ ] Introduire un contexte de build explicite : identité du build, index des
      modules et emplacements de résultats finaux immuables.
- [ ] Faire passer toutes les lectures externes par une interface contrôlée,
      en distinguant valeur disponible, absence définitive et résultat en attente.
- [ ] Remplacer les lectures implicites du magasin global dans la conversion
      par une vue propre au consommateur.
- [ ] Appliquer ce contrat aux directives externes inférées, tout en reproduisant
      exactement la priorité des directives de configuration, des directives
      locales/exportées et des règles inférées.
- [ ] Auditer aussi les opérations autres que `lookup` : fusion, export des
      directives du module et `addStop`. Les arrêts d'inlining restent locaux
      à l'évaluation ; leur insertion préserve les autres accessors de l'entrée.
- [ ] Garder les implémentations locales évolutives au fil des bindings, avec
      le même traitement des groupes récursifs que la référence.
- [ ] Rendre la validité des mémos explicite : les absences dépendent de la vue
      du consommateur ; une dépendance en attente ne devient jamais un miss définitif.
- [ ] Comparer ce nouveau chemin en mode séquentiel à la référence indépendante
      avant d'ajouter l'ordonnancement parallèle.

Sources principales dans `purescript-backend-optimizer-gopurs/` :
`src/PureScript/Backend/Optimizer/{Builder,Convert,Semantics,Cache,BoundedMemo}`.

## 3 — Remplacer les lots par un ordonnancement continu

Première mise en œuvre : conversions sur vues fixes avec découverte des
dépendances manquantes et relances contrôlées. Cela permet de conserver le cœur
de l'évaluateur synchrone et de mesurer le coût des relances.

- [ ] Préparer les entrées et isoler les hooks ayant des effets des tentatives
      de conversion ; un résultat provisoire n'est ni émis ni publié.
- [ ] Donner à chaque tentative une vue fixe des résultats disponibles et des
      mémos privés. Journaliser toute lecture d'un prédécesseur encore manquant.
- [ ] Accepter un résultat uniquement si toutes ses lectures sont définitives.
      Sinon, écarter ce résultat et remettre le module en attente de ses dépendances.
- [ ] Amorcer les dépendances avec les imports antérieurs connus, en excluant
      auto-imports et modules absents du corpus. Compléter avec les consultations
      réelles ; les références tardives sont traitées par la visibilité.
- [ ] Maintenir une file prête et un nombre borné de conversions actives ;
      réapprovisionner les workers à chaque achèvement utile.
- [ ] Donner priorité à la progression des prédécesseurs requis, avec un ordre
      de départ reproductible. Les tâches en attente libèrent leur place d'exécution.
- [ ] Publier chaque résultat final dans son emplacement ; les autres conversions
      y accèdent selon leur vue, indépendamment du curseur de sortie du coordinateur.
- [ ] Alimenter initialement le codegen dans l'ordre canonique, avec les snapshots
      attendus par l'émetteur existant, afin d'isoler la validation de PBO.
- [ ] Remplacer le fallback actuel qui contourne les dépendances par un diagnostic
      d'invariant précis si des modules restent bloqués sans travail possible.
- [ ] Mesurer le nombre et le coût des tentatives rejetées, l'attente et le RSS.
      Si les relances dominent, affiner la reprise aux groupes de bindings avant
      d'augmenter le nombre de workers.
- [ ] Après validation, tester une priorité pondérée par les coûts observés et
      une affinité de worker pour les caches réutilisables, à capacité mémoire bornée.

Sources : `Builder.purs`, `Cache.{purs,js,go}`, `BoundedMemo.{purs,js,go}`,
`gopurs/gopurs/src/Main.purs` et la FFI d'ordonnancement strictement nécessaire.

Livrable : parité complète séquentiel/parallèle, notamment pour `Data_Either.go`,
et première campagne de scaling 1/2/4/8 workers.

## 4 — Avancer la publication des signatures utilisées par l'émission

Le codegen consulte `metadata.globalFunctions` pour choisir certains appels
directs. Cette visibilité doit être traitée aussi soigneusement que `purmeta`.

- [ ] Mesurer le temps de préparation des signatures et celui de l'émission des
      corps, puis leur impact sur l'attente du producteur.
- [ ] Extraire une préparation de codegen réutilisable contenant le module
      transformé, les bindings TCO, les signatures et les auxiliaires nécessaires.
      Inclure les transformations précédant `ModuleBindings.prepare`, notamment
      les schémas de décodeurs et l'ownership, ainsi que `owned.functions`.
- [ ] Calculer une seule fois cette préparation et la consommer lors de la
      génération des corps.
- [ ] Construire les vues de signatures correspondant à l'ordre logique de
      référence ; rendre ces vues disponibles dès que leurs informations sont prêtes.
- [ ] Autoriser le chevauchement de générations de corps qui disposent déjà de
      leurs vues complètes, en conservant la parité des choix d'appels et de l'ABI.
- [ ] Vérifier les appels importés, récursifs, partiels et les wrappers FFI.

Sources : `gopurs/gopurs/src/Gopurs/{CodeGen,ModuleBindings,CallExprs,Emission}.purs`
et `gopurs/gopurs/src/Main.purs`.

## 5 — Réduire les allocations et accélérer la préparation

- [ ] Reprofiler le candidat déterministe pour sélectionner les deux ou trois
      postes les plus coûteux, avec attribution non additive des profils cumulés.
- [ ] Améliorer les représentations/clés des tables chaudes : identifiants
      compacts, comparateurs spécialisés et suppression des conversions intermédiaires.
- [ ] Employer des structures mutables à propriétaire unique pour l'état de
      travail local quand le contrat le permet ; figer les résultats partagés.
- [ ] Évaluer la réutilisation bornée des mémos et buffers ; mesurer leur taux
      de succès, le travail économisé et les données retenues.
- [ ] Appliquer les enseignements de répartition à la préparation et aux tours
      de spécialisations transitives, en conservant leurs snapshots et fusions
      déterministes. Mesurer la fusion et le point fixe séparément des workers.
- [ ] Pour chaque changement, mesurer son effet isolément et conserver la parité
      avec la référence. Les réglages GC restent constants dans les comparaisons.

L'objectif est que les gains apparaissent dans le compilateur généré et soient
réutilisables par les autres programmes compilés lorsque la transformation est générale.

## 6 — Validation et campagne finale

### Régressions ciblées nécessaires

- [ ] Corpus vide, auto-imports, imports externes et couverture de tous les modules.
- [ ] Directive A/B/C, directives par défaut, locales, exportées, inférées et arrêts
      d'inlining, avec vérification des priorités et des accessors.
- [ ] Référence transitive par inlining et référence introduite par spécialisation.
- [ ] Module tardif déjà calculé mais invisible ; prédécesseur manquant puis
      disponible ; identifiant réellement absent ; isolation de deux builds successifs.
- [ ] Achèvements volontairement désordonnés, relances, erreurs et annulation :
      publication et émission uniques des résultats finaux.
- [ ] Reproduction réduite du cas `Data.Either` / `Effect`, comparée entre JS
      et natif, puis exécutée pour vérifier son comportement.
- [ ] Tests natifs avec `-race` sur les nouveaux magasins, workers et mémos,
      plus un parcours intégré du compilateur reconstruit avec le détecteur.

### Vérifications existantes à réutiliser

- [ ] Construire le JS et le natif (`npm run build`, `npm run build:native` dans
      le checkout de candidat), puis vérifier la provenance du binaire exécuté.
- [ ] Exécuter les régressions PBO pertinentes : `implementation-lookup`,
      `binding-order`, `bounded-memo`, `purmeta-build-cache`, `monomorphize-*`
      et `transitive-parallel`, puis la suite PBO appropriée.
- [ ] Résoudre l'échec de dépendance `esbuild` précédemment signalé pour
      `monomorphize-cache.mjs` afin d'obtenir un bilan complet.
- [ ] Exécuter `node --test tools/emission.test.mjs`, les tests natifs de
      préparation et les tests de workers importés/ABI après les modifications
      correspondantes.
- [ ] Utiliser `./bin/test` pour compiler et exécuter les fixtures concernées,
      puis la sélection complète lors de la validation finale.
- [ ] Comparer les sorties b8x dans des destinations vierges : ensemble des
      chemins, nombres de fichiers, contenu et SHA-256. Un répertoire contenant
      des fichiers d'un ancien passage ne constitue pas un oracle de couverture.
- [ ] Exécuter `go build ./...` sur la sortie b8x candidate et les tests
      d'application pertinents ; distinguer compilation, exécution et performances.

### Mesures

- [ ] Utiliser des TAST et FFI figés ; préserver le cwd/résolution de b8x et
      l'invocation backend sans `--main` pour la comparabilité historique.
- [ ] Comparer 1/2/4/8 workers PBO sur le même binaire candidat ; fixer les autres
      paramètres, notamment préparation, émission, pipeline, GC et processeurs.
- [ ] Faire au moins trois passages appariés pour les configurations finales,
      sans chevauchement avec un bootstrap ou une autre campagne ; publier
      médiane et dispersion. Profiler séparément des passages chronométrés.
- [ ] Publier les temps de phases, CPU, RSS, allocations, GC, couverture,
      divergences, relances et occupation utile mesurée.
- [ ] Mesurer ensuite `b -c` de bout en bout avec un protocole explicite de caches.
- [ ] Choisir le nombre de workers par défaut à partir des résultats de correction,
      de temps et de mémoire, puis documenter `GOPURS_PBO_JOBS`.

## Jalons

1. **Référence et diagnostic** : baseline corrigée validée, lectures et temps mesurés.
2. **Parité** : environnement explicite et conversions parallèles byte-identiques.
3. **Premier gain confirmé** : scaling apparié, cible de travail autour de 80 s.
4. **Optimisation étendue** : émission, allocations et préparation, cible autour de 60 s.
5. **Intégration** : tests complets, mesure de `b -c`, choix du défaut et documentation.

Les temps observés et le profil des attentes déterminent les priorités après
chaque jalon ; les objectifs de durée sont réévalués à partir de ces mesures.
