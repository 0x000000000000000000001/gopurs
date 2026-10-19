> **Archive figée — 19 octobre 2026.** Copie de `todo.md` avant le passage au
> plan actif court. Conservée pour l'historique des lots 00–06, G00–G06,
> P00–P06, des seuils, échecs et décisions. Ne pas mettre à jour ce document.
> Le plan actif est dans [`todo.md`](../todo.md).
>
> Empreinte SHA-256 de la copie d'origine :
> `3c28923ef2543d657ceeb5f3525c872aba30e789d72b45d71866d091877a9ccd`.

# Gopurs — build incrémental et lancement Go

Plan v3 et suites — mise à jour : 16 octobre 2026.

## Objectif et périmètre

Réduire le cycle **modifier un fichier → construire → exécuter les tests** en
réutilisant les résultats encore valides. La phase gopurs complète, signalée à
environ **33 secondes**, est mesurée à **36,902 s de médiane sans changement**
au lot 00 ; la chaîne jusqu'à l'exécution prend **40,808 s**. La
[référence ventilée](docs/build-baseline.md) fixe les conditions de comparaison.
Le lot 03 ramène la médiane gopurs identique à **1,257 s** et la chaîne à
**3,958 s**, avec zéro sortie Go réécrite et zéro compilation/lien Go.

Les repères initiaux rappelés par l'utilisateur restent visibles dans les bilans :
**~60 s de backend à froid et ~30 s après modification**. Ce sont des repères
historiques, à distinguer des mesures appariées de chaque lot. Les références
intermédiaires v3/06 (`publicsync`) et G06 servent à attribuer les gains ; le
retour vers ces temps initiaux reste l'objectif de performance du backend.

La première livraison doit couvrir les builds sans changement et les
modifications localisées. Elle concerne gopurs, les adaptations nécessaires du
PBO et le runner Go de `../../b8x`. Le résultat incrémental doit être identique
au résultat d'un build complet.

Le plan v2 est clôturé ; sa
[validation finale](docs/testing.md#plan-v2--lot-08--installation-et-validation-finale-7-octobre-2026)
conserve les preuves historiques.

## Avancement : 7/7 lots validés — plan v3 clôturé

Cocher un lot uniquement après satisfaction de son critère de fin et des
vérifications communes. Actualiser ce compteur et consigner les résultats par
plan et lot dans [docs/testing.md](docs/testing.md).

## Constats de départ

- `Driver.Build.onSkipModule` renvoie toujours `Nothing` : chaque invocation
  refait l'optimisation et l'émission de tous les modules.
- `Driver.Prepare` recharge les TAST et refait les métadonnées globales ainsi
  que la collecte transitive des spécialisations.
- `.purmeta` est un stockage interne au build courant. Le natif le conserve en
  mémoire ; le format V8 du chemin JS ne fournit pas un cache persistant commun.
- `Driver.Output` réécrit les sorties et ne retire pas les anciens fichiers
  lorsqu'un module ou une FFI disparaît.
- Les modules générés sont réunis dans **un seul package Go `purescript`**.
  Une modification réelle d'un fichier peut donc invalider ce gros package.
- `b8x/bin/_go-run` demande un nouveau binaire temporaire à chaque invocation.
  La déduplication des exécutables intervient après `go build`.
- Le runner affiche « Waiting for the Go build lock… » avant toute acquisition,
  même immédiate. Son verrou est partagé entre les conteneurs API ; les
  healthchecks `Ping` utilisent également ce parcours toutes les 30 secondes.
- Ses valeurs par défaut sont `GOMAXPROCS=2` et `go build -p 1`. Les caches Go
  sont placés dans le volume persistant `var/api/cache/go`.

## Lots — ordre de travail

### 00 — Mesurer le cycle réel

- [x] **Distinguer les coûts et conserver la référence avant changement.**
  - Relever les versions, cibles, options, chemins et états des caches réellement
    utilisés sur l'hôte et dans le conteneur.
  - Mesurer chargement TAST, préparation/spécialisation, optimisation, émission,
    attente effective du verrou, compilation Go et édition de liens.
  - Vérifier les concurrents du runner, notamment les lancements des services
    et les healthchecks, sans attribuer une contention au seul message affiché.
  - Comparer build complet, relance identique, modification d'un module feuille,
    modification d'une fonction partagée et modification d'une FFI.
  - **Fin :** références reproductibles, sources et artefacts identifiés, temps
    ventilés ; le coût du verrou est séparé du coût du build Go.

  **Validé le 10 octobre 2026 :** corpus figé de **2 704 TAST / 755 005 types**.
  La relance identique réécrit **3 010 Go byte-identiques** et prend **36,902 s**
  dans gopurs ; la chaîne Spago/tidy/runner jusqu'à `exec` prend **40,808 s** de
  médiane. Le runner seul prend **0,770 s**, sans compilation ni lien. Une édition
  de 1 à 5 fichiers Go déclenche **212–220 s** de compilation du package commun.
  Verrou mesuré séparément, cinq oracles réussis, parité des quatre corpus dans
  les trois modes, six replays stricts et audit des sources/artefacts réussis.
  Voir la [référence détaillée](docs/build-baseline.md) et le
  [bilan de validation](docs/testing.md#plan-v3--lot-00--référence-du-cycle-b8x-10-octobre-2026).

### 01 — Réutiliser le binaire validé par Go dans b8x

- [x] **Éviter les reconstructions d'exécutables déjà à jour.**
  - Faire valider le dernier artefact par `go build`, au lieu de lui présenter
    systématiquement une cible temporaire vide. Éprouver cette stratégie avec
    le vrai toolchain du conteneur.
  - Conserver la publication d'artefacts immuables : chaque invocation exécute
    son propre résultat après libération du verrou.
  - Préserver les erreurs, les annulations et la terminaison des sous-processus ;
    un build échoué ne doit jamais déclencher l'ancien exécutable.
  - Rendre les messages explicites sur l'attente mesurée et sur les deux limites
    distinctes : ressources du processus et parallélisme entre packages.
  - Mesurer les limites de ressources avant de modifier leurs valeurs.
  - **Fin :** relance identique sans compilation ni lien superflus, changement
    pris en compte, concurrence et annulation vérifiées ; gains mesurés dans le
    runner réel et contrats de `b8x/test/go-run.mjs` préservés.

  **Validé le 10 octobre 2026 :** dix contrôles réussis avec le Go du conteneur,
  reproducteur natif rouge avant changement puis vert après, et lancement réel
  de `Ping` réussi. Sur cinq paires alternées, la médiane jusqu'à `exec` du binaire
  de tests à chaud passe de **1,812 s à 0,745 s (−58,9 %)** ; les cinq relances
  optimisées ne compilent ni ne lient, avec le même exécutable à l'octet près.
  Voir le [bilan, les limites et les preuves](docs/testing.md#plan-v3--lot-01--réutilisation-des-exécutables-b8x-10-octobre-2026).

### 02 — Définir le cache persistant et la propriété des sorties

- [x] **Introduire un contrat de cache versionné et vérifiable.**
  - Identifier les entrées par contenu : inventaire des TAST, FFI résolues et
    leur présence/absence, chemins pertinents, directives et options sémantiques.
  - Inclure une identité du compilateur couvrant PBO, runtime et parseur FFI,
    y compris les modifications locales ; le seul commit Git ne suffit pas.
  - Définir les données persistées par module et un encodage exploitable par les
    hôtes JS et Go ; ne pas sérialiser directement leur représentation mémoire.
  - Définir l'inventaire des sorties possédées par gopurs et leur intégrité.
    Traiter explicitement `go.mod`/`go.sum` et les modifications normales de
    `go mod tidy`, afin de préserver les dépendances et la validité du cache.
  - Publier les entrées et le manifeste atomiquement ; coordonner les builds
    concurrents d'un même workspace et isoler les autres projets.
  - Une entrée absente, incompatible ou corrompue déclenche une reconstruction.
  - **Fin :** clés, schéma, publication et invalidation documentés, avec tests
    ciblés d'intégrité, de version et d'isolation.

  **Validé le 10 octobre 2026 :** [contrat v1](docs/build-cache.md), stockage
  partagé JS/Go et journal de propriété ; tests sous race detector sur macOS et
  Linux, reprise après interruption, parité avant/après sur 121 TAST dans les
  trois modes et six replays stricts réussis. Le pilote sera raccordé aux lots
  03–04 ; la mesure de référence du lot 00 précède l'évaluation de son gain.
  Voir le [bilan et les preuves](docs/testing.md#plan-v3--lot-02--contrat-du-cache-persistant-10-octobre-2026).

### 03 — Build sans changement et écritures minimales

- [x] **Terminer rapidement lorsque les entrées sont inchangées.**
  - Vérifier le manifeste avant le décodage complet des TAST et les passes
    d'optimisation. Les seules dates de modification ne prouvent pas la validité.
  - Réutiliser les sorties valides ; restaurer ou régénérer les fichiers absents
    ou altérés.
  - Écrire uniquement les contenus modifiés et retirer les sorties obsolètes
    appartenant au manifeste, en préservant les fichiers externes à cet inventaire.
  - Prendre en compte la sélection des points d'entrée et la résolution des FFI.
  - **Fin :** une relance identique n'exécute pas les passes coûteuses et conserve
    les fichiers inchangés ; suppressions et sorties endommagées sont traitées
    correctement. Le coût de vérification du cache est mesuré.

  **Validé le 10 octobre 2026 :** sur trois paires alternées du corpus b8x,
  gopurs passe de **35,732 s à 1,257 s (−96,5 %)** ; la chaîne jusqu'à `exec`
  passe de **38,145 s à 3,958 s (−89,6 %)**. Les **3 010 Go** sont byte-exacts
  dans les trois modes ; les hits ne décodent ni ne préparent les TAST et ne
  réécrivent aucun Go. Suppressions, réparations, priorités FFI, erreurs,
  interruption, concurrence et collisions externes sont vérifiées ; six replays
  stricts passent sans modification des snapshots. La première adoption coûte
  **65,054 s**, dont **25,660 s** de publication durable. Voir le
  [bilan complet et les limites](docs/testing.md#plan-v3--lot-03--build-identique-vérifié-10-octobre-2026).

### 04 — Incrémentalité par module après modification

- [x] **Réutiliser l'optimisation et l'émission des modules encore valides.**
  - Conserver initialement la préparation globale des métadonnées et des
    spécialisations ; comparer les entrées préparées réellement consommées.
  - Persister les résultats optimisés, les implémentations/directives PBO, les
    signatures des workers et les sorties Go/FFI nécessaires à leur réutilisation.
  - Raccorder `onSkipModule` et republier les contributions des modules réutilisés
    dans l'ordre canonique, y compris dans le pipeline d'émission parallèle.
  - Inclure l'environnement d'optimisation et de génération dans la preuve de
    validité. Prévoir une invalidation conservatrice lorsque les dépendances
    exactes ne sont pas encore établies.
  - Couvrir les corps inlinés, les lectures d'implémentations et leurs absences,
    les directives accumulées, les layouts/ABI et les signatures FFI.
  - Couvrir les demandes de spécialisation : un appelant modifié peut changer le
    module qui fournit une spécialisation, même si la source de ce dernier est
    inchangée. Les seuls imports et signatures publiques ne suffisent pas.
  - **Fin :** une modification localisée réutilise effectivement des modules,
    avec compteurs et raisons d'invalidation lisibles ; résultats identiques au
    build complet pour les scénarios d'invalidation.

  **Validé le 12 octobre 2026 sur `publicsync`, avec compromis explicite :**
  trois paires face au lot 03 donnent **−10,37 % feuille / −4,48 % partagé /
  −32,76 % FFI**. Le froid passe de **94,191 à 113,213 s (+20,19 %)** ; l'identique
  suivant ces froids passe de **1,366 à 1,716 s (+0,350 s)**. Les objets à froid
  passent de **135,326 à 548,673 Mo**. Les validations de correction complètes
  passent ; l'utilisateur confirme la poursuite après présentation du surcoût.
  Le [bilan livré](docs/build-incremental.md) conserve méthode, limites, compteurs
  et identités. Le lot 06 mesure encore le parcours complet.

  **Historique — candidat `allhit` non validé au 11 octobre :** il passe les
  treize transitions dans les quatre hôtes, la parité b8x JS/Go et les six
  replays stricts. Sur cinq paires, l'édition FFI gagne **19,85 %**
  (**38,332 → 30,724 s**). La campagne complète révèle cependant des
  régressions : feuille **39,505 → 53,865 s**, partagée **39,111 → 88,715 s**,
  cache/sorties neufs **93,599 → 239,367 s**. Le préfixe conservateur réutilise
  respectivement 1 904 et 7 modules ; les octets restent égaux aux builds forcés.
  Les optimisations retenues `wirecalls` et `smallctors` gagnent respectivement
  **12,07 %** et **9,75 %** face à leur prédécesseur sur les paires partagées.
  `stagepool` réduit la publication froide isolée de **174,391 à 95,186 s
  (−45,42 %)**. `workerkeys` restaure 2 698 émissions sur le probe partagé et
  gagne **10,60 %** face à `stagepool` seul sur trois paires (**68,087 → 60,866 s**).
  `batchsync` gagne ensuite **23,85 %** sur le stockage froid isolé
  (**93,644 → 71,312 s**), avec tests macOS/Linux sous race detector. Son probe
  b8x froid prend **146,488 s** ; il reste plus lent que le lot 03. `syntaxfns`
  gagne **6,76 % / 4,29 s** sur trois paires partagées (**63,502 → 59,211 s**),
  avec baisse d'allocations et octets identiques. Sa campagne complète et les
  validations multi-hôtes passent, mais les trois paires actuelles face au lot 03
  donnent **37,785 → 58,459 s (+54,72 %)** sur l'édition partagée. Le froid est à
  **133,543 s**, la feuille à **49,502 s** et la FFI à **35,500 s**. L'essai
  `restoreprep` gagne ensuite **18,17 %** en FFI (**36,325 → 29,725 s**) et
  **7,88 %** sur le partagé (**58,859 → 54,223 s**) sur trois paires de chaque
  scénario. Il est retenu, avec reproducteurs de restauration anticipée et de
  repli sur une référence introduite. Sa campagne complète et ses validations
  multi-hôtes passent : **136,821 s froid / 2,045 s identique / 45,440 s feuille /
  55,780 s partagée / 30,215 s FFI**. Trois paires actuelles face au lot 03
  confirment encore **37,550 → 51,087 s (+36,05 %)** sur le partagé. Une trace
  isolée des lectures PBO prédit **2 643 réutilisations / 61 invalidations**, avec
  artefacts vérifiés contre les builds forcés, au lieu des **7 réutilisations**
  du préfixe. L'essai `pboreads` intègre ces preuves dans `pbo/backend-v2`, en
  gardant quatre objets par module. Le reproducteur rouge/vert, les contrôles PBO,
  le bootstrap Go, le stockage macOS sous race detector et les onze contrôles
  codecs/workers/cache passent. Le probe partagé donne **43,631 s**, avec les
  **61 / 2 643** optimisations/réutilisations prévues et cinq Go réécrits.
  Les trois paires partagées donnent **53,687 → 45,051 s (−16,09 %)** ; le candidat
  est retenu. Quatorze transitions cache/frais supplémentaires passent, ainsi que
  le rouge/vert d'une raison spéculative restée après un hit. Le candidat final
  est reconstruit et figé, avec douze contrôles ciblés verts. Sa campagne Go
  parallèle passe : **135,630 s froid / 1,738 s identique / 42,671 s feuille /
  44,590 s partagée / 30,768 s FFI**. Les trois paires actuelles donnent encore
  **39,404 → 46,093 s (+16,97 %)** face au lot 03 sur le partagé. La parité JS/Go
  séquentiel/Go parallèle passe sur les quatre états, ainsi que le stockage Linux
  sous race detector. Le bootstrap Rust privé, les treize transitions des quatre
  hôtes, les six replays stricts (**35 commandes**), les **19 contrôles CLI** et
  l'audit de **411 fichiers** passent ; aucun snapshot n'est modifié. Le lot reste
  ouvert pour les performances.
  Seuil fixé avant confirmation : **au moins 10 %** face à `restoreprep` sur
  trois paires partagées. Le profil des allocations motive ensuite l'essai
  `scalarhex` : validation hexadécimale directe dans les bridges scalaires.
  Les trois paires confirment **45,346 → 40,791 s (−10,05 %)**, au-dessus du seuil
  préalable de **5 %** face à `pboreads-final`, avec les mêmes octets. Le candidat
  est retenu et figé après le bootstrap Go et douze contrôles ciblés verts. Sa
  campagne complète donne **131,050 s froid / 1,982 s identique / 37,781 s feuille /
  39,396 s partagée / 27,924 s FFI**, avec parité des trois hôtes sur les quatre
  états et bootstrap Rust réussi. Les paires actuelles face au lot 03 donnent
  encore **37,404 → 40,019 s (+6,99 %)** sur le partagé. Les treize transitions
  des quatre hôtes, les six replays stricts (**35 commandes**), les **19 contrôles
  CLI** et l'audit de **411 fichiers** passent ; les snapshots restent inchangés.
  Le lot reste ouvert pour les performances. Le profil confirme ensuite une
  baisse des allocations de décodage de **35,61 %** avec `scalarhex`. L'essai
  `nodeviews` analyse une seule fois les formes JSON et les références du graphe,
  mais ne gagne que **0,07 %** sur les trois paires face à `scalarhex`, sous son
  seuil préalable de **10 %**. Il est rejeté malgré les tests verts ; les
  **411 fichiers** des sources/outils/binaires sont revenus exactement au candidat
  `scalarhex` validé, avec preuves expérimentales conservées. L'essai `inttables`
  raccorde ensuite les comparateurs entiers existants aux deux tables du codec,
  mais ne gagne que **1,85 %**, sous le seuil préalable de **5 % sur trois paires
  partagées**. Il est également rejeté ; le retour exact à `scalarhex` est audité.
  Le diagnostic des transferts privés écarte les liens physiques, plus lents que
  les copies (**3,205 contre 1,762 s**). Le profil des empreintes préparées retrouve
  les 2 704 clés exactes et **64,829 Go alloués**, dont environ 41 % des allocations
  échantillonnées dans les comparaisons structurelles de types. L'essai
  `typeintern` combine le raccourci d'identité conservé et l'internage canonique
  des types, avec bridges existants et wire identique. Il est retenu à **−7,03 %**
  (**43,923 → 40,838 s**) sur trois paires partagées face à `scalarhex`, au-dessus
  du seuil préalable de **5 %**, avec un gain dans chaque paire. Les 13 contrôles
  ciblés, le différentiel et la parité des artefacts passent. La validation complète
  termine le 12 octobre : **127,687 s froid / 1,900 s identique / 38,493 s feuille /
  37,654 s partagée / 29,235 s FFI**, avec sorties incrémentales/forcées byte-exactes.
  Les trois paires actuelles face au lot 03 donnent **39,273 → 38,480 s (−2,02 %)**
  sur le partagé. La parité b8x des trois hôtes, le bootstrap Rust privé, les
  treize transitions des quatre hôtes, les **35 commandes strictes**, les **19
  contrôles CLI** et l'audit des **411 fichiers** passent ; snapshots inchangés.
  Le lot reste ouvert : le gain partagé est modeste et le froid, dont **66,593 s
  de publication**, reste au-dessus des **93,599 s historiques** du lot 03.
  Le diagnostic du stockage mesure **38,676 s de staging / 27,032 s de publication
  publique** sur **65,919 s** isolées. L'essai privé `publicpool2` reprend la
  publication bornée à huit tâches après `batchsync`, avec les synchronisations
  existantes. Seuils préalables : **10 % sur trois paires de publication froide**,
  puis **5 % sur trois paires de build froid complet**. Les trois paires de
  stockage donnent **72,155 → 59,773 s (−17,16 %)**, avec octets exacts, lookups et
  tests macOS sous race detector verts. Le stockage Linux et l'audit du relien
  passent ; les trois paires froides complètes donnent **127,163 → 120,354 s
  (−5,35 %)**, au-dessus du seuil de 5 %, avec un gain dans chaque paire et les
  mêmes octets. `publicpool2` est retenu puis installé après audit de **411
  fichiers**. Sa validation complète est lancée, avec trois paires actuelles
  face au lot 03 pour chacun des scénarios feuille/partagé/FFI. La campagne Go
  parallèle termine à **121,921 s froid / 1,785 s identique / 41,478 s feuille /
  43,920 s partagée / 29,764 s FFI**, avec contrôles forcés byte-exacts sans
  réécriture. Les paires donnent **−4,57 % feuille / +2,78 % partagée** face au
  lot 03. La comparaison FFI reprend dans une nouvelle tentative après correction
  d'une assertion du harnais qui exigeait les compteurs du lot 04 sur le lot 03 ;
  les **411 fichiers** sont inchangés et le build de référence est byte-exact.
  La reprise réussit les paires FFI à **41,195 → 30,960 s (−24,85 %)**. La parité
  b8x des trois hôtes, le bootstrap Rust privé sans PGO, les treize transitions
  des quatre hôtes, les **35 commandes strictes**, les **19 contrôles CLI** et
  l'audit des **411 fichiers** passent ; snapshots inchangés. `publicpool2` est
  entièrement validé. Le lot reste ouvert pour le bilan global de performances.
  Le profil actuel des clés confirme **−29,87 % d'allocations des modules /
  −63,18 % des métadonnées** depuis `scalarhex`, à octets identiques. Il reste
  **45,466 Go alloués** pour les modules, dont **17,85 %** échantillonnés dans la
  suppression successive des champs record. L'essai `recordfields` lit les
  champs sur le record initial, avec les bridges existants et les mêmes octets.
  Seuil préalable : **5 % sur trois paires partagées face à `publicpool2`**.
  Le reproducteur passe de **72 à 0 copies observées**, avec **5 891 octets
  identiques** ; les **200 vecteurs**, **62 040 comparaisons différentielles**,
  **14 contrôles ciblés** et le bootstrap Go passent. Les probes restent
  byte-exacts, mais les trois paires donnent **37,437 → 38,824 s (+3,71 %)**.
  L'essai est rejeté et le retour à `publicpool2` est audité sur **425 fichiers**,
  avec tests du lot 05 préservés. La piste suivante `backrefs` remplace les deux
  tables arborescentes du décodeur par des tables indexées à préfixes bornés,
  avec bridges JS/Go/Rust de transport opaque et validation PureScript. Seuil
  préalable : **5 % sur trois paires partagées face à `publicpool2`**. Le
  reproducteur passe de **698 à 0 comparaisons d'indices**, à octets identiques ;
  les 16 contrôles ciblés et le bootstrap Go passent. Le bootstrap Rust réussit
  après correction `RefCell` → `Mutex`. Sa revue révèle ensuite une construction
  à chaque lookup (**256 callbacks au lieu de 32** dans le reproducteur compilé).
  La tentative de paires est interrompue et auditée sur **430 fichiers**.
  `backrefs-table` renvoie une table opaque, stockée dans le contexte puis lue
  séparément. Ses 17 contrôles, bootstraps et reproducteur Rust passent, mais
  les trois paires donnent **35,982 → 34,995 s (−2,74 %)**, sous le seuil de 5 %.
  L'essai est rejeté ; les **425 fichiers** de `publicpool2` et des tests de
  sources sont restaurés exactement, avec contrôles ciblés verts.
  `publicpool2` reste la référence pleinement validée.
  L'essai privé `publicsync` regroupe ensuite les barrières de répertoires après
  jonction des writers, avec fichiers synchronisés individuellement, journal
  d'union durable et reprise des sorties identiques. Le reproducteur passe de
  **34 à 4 synchronisations**, avec échec/reprise sous race detector vérifiés.
  Seuils fixés avant mesure : **10 % sur trois paires de publication froide
  isolée**, puis **5 % sur trois paires froides complètes**. Le stockage gagne
  **14,48 % (59,636 → 50,999 s)**, dans chacune des trois paires, à octets exacts.
  SIGKILL, race detector macOS/Linux et relien natif audité passent. Les trois
  paires froides complètes gagnent **5,14 % (118,971 → 112,857 s)**, avec un gain
  dans chaque paire et les mêmes octets. Les budgets de non-régression passent :
  **−0,71 % feuille / −5,19 % partagé / +1,71 % FFI**, avec l'identique à
  **1,658 → 1,612 s**. Le candidat mesuré est installé atomiquement et figé sur
  **430 fichiers** ; les reproducteurs d'échec/SIGKILL sont intégrés à
  `test:cache`. La validation complète b8x/multi-hôtes/Rust/transitions/replays
  stricts réussit, avec les paires actuelles face au lot 03.
  Les trois paires actuelles passent à **−10,37 % feuille / −4,48 % partagé /
  −32,76 % FFI** face au lot 03. Les barrières intégrées, les bridges et les
  campagnes b8x des trois hôtes passent. Les paires froides actuelles face au
  lot 03 sont lancées ; le lot reste ouvert jusqu'au bilan global.
  **Décision de périmètre :** `publicsync` est le candidat de clôture du lot 04.
  Après explicitation du coût du premier build, l'utilisateur confirme la
  poursuite avec ce compromis ; le retour à 30 s à froid n'est pas un critère
  atteint ni une condition ajoutée pour cette clôture. La comparaison froide
  actuelle et le bilan du surcoût restent à produire. Les campagnes b8x dans
  les trois hôtes, le bootstrap Rust, les transitions du pilote et les **56
  comparaisons de sources / 294 commandes** passent. La matrice de contexte
  réussit ses **80 comparaisons / 401 commandes**, puis les **35 commandes
  strictes**, les **19 contrôles CLI** et l'audit des **430 fichiers** passent.
  Les snapshots restent inchangés ; `validationComplete: true` est enregistré.
  Les trois paires froides actuelles terminent à **94,191 → 113,213 s (+20,19 %)**,
  avec Go/TAST exacts, cache neuf par échantillon et audit final inchangé. Le
  bilan documente aussi **+0,350 s** sur l'identique et **+413,347 Mo** d'objets.
  Voir le
  [bilan et les preuves](docs/testing.md#plan-v3--lot-04--incrémentalité-par-module-validation-en-cours-11-octobre-2026).

### 05 — Campagne de correction incrémental/complet

- [x] **Éprouver les transitions de build sur des reproducteurs persistants.**
  - Relance identique ; modification privée ; corps exporté à signature stable ;
    corps inliné ; changement de type, constructeur, record ou dictionnaire.
  - Ajout/retrait d'une demande de spécialisation et changements transitifs.
  - Modification du corps ou de la signature d'une FFI ; ajout/retrait d'une FFI.
  - Ajout, renommage et suppression de modules dans les entrées TAST.
  - Changement de compilateur, runtime, options et point d'entrée.
  - Cache absent/corrompu, sortie manquante/altérée, build interrompu ou échoué,
    reprise et builds concurrents.
  - **Fin :** pour chaque transition, inventaire et octets du Go généré égaux
    entre build incrémental et build complet ; parité JS/natif séquentiel/natif
    parallèle sur les mêmes TAST, compilation Go et exécution des régressions.

  **Validé le 12 octobre 2026 sur `publicsync` :** les transitions sources,
  contexte, pilote et stockage couvrent le critère de fin dans les quatre hôtes,
  avec compilation/exécution, replays stricts et audit final. Voir la
  [matrice de couverture et les preuves](docs/build-incremental.md#validation-de-correction--lot-05).

  **Historique de préparation du 12 octobre 2026 :** le brouillon privé
  `lot05-source-transitions/` complète la revue de couverture avec **14 états de
  sources** : corps privé/exporté, directive, constructeur, record/type de champ,
  dictionnaire, demandes de spécialisation ajoutées/remplacées/retirées et module
  renommé avec sa FFI et son entrée secondaire. Le harnais prépare chaque TAST
  une fois, puis prévoit les comparaisons incrémental/frais et multi-hôtes, les
  hits identiques et les oracles exécutables. La syntaxe JS et la construction
  des 14 variantes passent. Après la campagne `publicpool2`, les **14 états
  PureScript compilent**. Le premier build JS révèle une assertion trop stricte
  de l'inventaire : sur macOS, `main/main.go` apparaît sous `Main/main.go`, dans
  le répertoire TAST existant. Les **93 sorties** sont exactes ; le harnais exige
  désormais une bijection manifeste/fichiers par identité de fichier, avec
  contrôle des hashes et tailles. L'essai initial est conservé. La reprise
  `run-4uQJYd` passe les **56 comparaisons (14 états × 4 hôtes)**, avec octets,
  parties et clés exacts, hits sans réécriture et oracles Go exécutés. Le test
  est intégré au dépôt sous `tools/driver-source-cache.test.mjs`, avec ses fixtures
  et `npm run test:cache-sources` ; sa campagne d'intégration réussit dans les
  quatre hôtes (`gopurs-tests-S2M1Cn`) : **294 commandes**, dont **56 builds
  frais / 56 hits sans réécriture / 56 compilations et exécutions**. Sources,
  fixtures et compilateurs sont audités. Le lot reste ouvert pour les transitions
  de contexte compilateur/runtime/options et la campagne finale.
  Le complément privé prépare **20 états × 4 hôtes**, avec versions validées
  `typeintern`/`publicpool2`, changements des options et des entrées, et runtime
  distinctif réellement embarqué dans des compilateurs privés JS/Go/Rust.
  Les constructions privées réussissent. Le harnais valide 13 états JS avant
  une attente erronée sur les deux chemins de l'entrée sélectionnée ; ses
  92 Go sont exacts et les 425 fichiers de production inchangés. La tentative
  est conservée et l'attente corrigée. La reprise `run-OfOWca` réussit les
  **80 comparaisons / 401 commandes**, avec le seul Go du runtime réécrit lors
  de son changement et retour, et les 425 fichiers de production inchangés.
  `npm run test:cache-context` pérennise le test ; `tools/cache-context-runtime.mjs`
  reconstruit les variantes privées depuis des bootstraps retenus. La campagne
  d'intégration passe dans les quatre hôtes (`gopurs-tests-v2nL09`, **80
  comparaisons / 401 commandes**), puis sans catalogue en JS
  (`gopurs-tests-O8cZFe`, **16 comparaisons / 81 commandes**). L'audit retrouve
  les TAST et clés exacts de la campagne de sources, les parties et Go exacts de
  la campagne privée, et les **427 fichiers** inchangés. Son attente initiale
  sur 57 `modulePath` relatifs entre workspaces de profondeurs différentes est
  corrigée après vérification des chemins résolus ; l'échec est conservé.
  Le candidat de clôture `publicsync` réussit le rejeu dans les quatre hôtes :
  **56 comparaisons / 294 commandes de sources** (`gopurs-tests-9VGNGK`), puis
  **80 comparaisons / 401 commandes de contexte** (`gopurs-tests-7uwezQ`). Le
  bootstrap Rust, les transitions du pilote, la reprise SIGKILL, les replays
  stricts, la CLI et l'audit final des **430 fichiers** passent. La couverture
  du lot 05 est établie ; le bilan froid du lot 04 étant terminé, le lot est clos.

### 06 — Mesurer le gain et livrer le parcours utilisateur

- [x] **Valider le bénéfice sur b8x et documenter l'usage.**
  - Reconstruire les compilateurs JS et natif, puis vérifier l'application
    consommatrice avec les artefacts effectivement utilisés.
  - Rejouer les scénarios du lot 00 avec les mêmes entrées, options et états de
    cache ; mesurer séparément gopurs et le parcours complet jusqu'aux tests.
  - Mesurer aussi le surcoût du premier build, la taille du cache et les modules
    recalculés après chaque type de modification.
  - Fournir un moyen explicite de forcer le build complet et de réinitialiser le
    cache du projet, ainsi que les diagnostics de réutilisation/invalidation.
  - **Fin :** gains significatifs et limites établis, snapshots revus puis
    rejoués strictement, documentation et preuves publiées dans le dépôt.

  **Validé le 12 octobre 2026 :** **24 chaînes / 24 oracles** passent, avec
  trois paires alternées par scénario, frontend réel, timer backend, tidy et
  runner jusqu'à `exec`. Chaque GOCACHE privé part de la base seulement, à
  **2 / -p 1 / 12GiB**. Les **2 704 TAST / 3 010 Go**, les fichiers de modules
  et les binaires par scénario sont byte-exacts. L'audit final conserve les
  **430 fichiers de production**, les seeds et les inodes de verrou. Les trois
  snapshots restent inchangés ; les 35 replays stricts du lot 05 ont éprouvé
  ce même compilateur, également reconstruit/testé en JS, Go et Rust.

  Face au lot 03, les médianes gopurs gagnent **6,48 % feuille / 1,40 % partagé /
  32,70 % FFI**. Les chaînes varient de **−1,44 % / +1,01 % / −5,50 %**, avec la
  dispersion par paire publiée : le gain du cycle feuille/partagé n'est pas
  robuste. L'identique passe de **3,559 à 3,962 s** (**+0,403 s**). Go absorbe
  **78–88 %** du temps après édition, ce qui confirme la priorité **G puis P**.
  Le froid gopurs apparié reste à **+20,19 %** ; les observations froides de
  chaîne, non appariées, sont **322,723 / 339,051 s**. Le volume et les limites
  sont explicités dans `docs/build-incremental.md`, avec les 24 mesures dans
  `docs/build-cycle-lot06.csv` et les commandes utilisateur dans `docs/build-cache.md`.

  Les échecs Spago/DNS et le passage interrompu restent conservés ; la reprise
  des huit derniers échantillons garde les racines/journaux et les seeds de
  base. Preuves finales : `gopurs-cycle-lot06-publicsync-resume1/` et
  `gopurs-cycle-lot06-analysis-final/`. Le reset documentaire est exécuté avec
  contrôle du verrou, des contenus et des dates. Le plan v3 est clos à **7/7**.

## Suites planifiées — préparation incrémentale et packages Go

**Mise à jour du 19 octobre 2026 : 1/7 lot P (audit P00) et 7/7 lots G validés.** Ces
quatorze lots ont leurs propres critères de fin ; le compteur du plan v3 reste
celui des lots 00–06. La référence v3/06 est conservée ; G06 fournit désormais
la référence immédiate pour une éventuelle reprise de P.

### Priorité et dépendances

Ordre confirmé par v3/06, avec point de décision avant P01 :
**G00 → G01…G06 → P00 → décision → P01…P06**.
**Borne de travail confirmée le 16 octobre 2026 : finir G, puis faire une pause
avant P. G a été livré, puis la pause respectée.** La campagne
du candidat 7, son bilan et les validations nécessaires à G06 ont réussi.
La borne fixée était de terminer les validations
nécessaires à G06 si les critères passent. Un critère bloquant doit être présenté
avant toute nouvelle série de candidats ; aucune nouvelle recherche
d'optimisation n'est engagée automatiquement.
**Reprise explicitement autorisée le 16 octobre : P00 seul, terminé le 19.** Son
[bilan chiffré](docs/build-preparation-baseline.md) recommande de revoir le
périmètre : les états complets sondés ne justifient pas P01–P06. L'étude
resserrée de P01 autorisée le 19 est terminée : **C01 rejeté en l'état**, avec
[bilan](docs/build-preparation-compact.md) et nouvelle décision avant la suite.

La priorité donnée au découpage Go s'appuyait sur v3/06 : **206,820 / 209,650 /
208,470 s** de `go build` après feuille/partagé/FFI, dont **202,788 / 205,196 /
204,576 s** dans les commandes de l'action du package `purescript`. Go
représentait **78–88 %** de la chaîne, contre environ **9–10 s** de préparation
globale gopurs. Ces références figurent dans `docs/build-incremental.md`.

Après G06, le backend mesuré en layout packages (`GOPURS_GO_PACKAGES=1`) prend
**48,487 / 49,672 / 40,001 s** après feuille/partagé/FFI et **117,830 / 117,740 s**
aux froids applicatif/complet. Le [bilan G06](docs/build-go-measurements.md)
conserve les paires ; P00 devra déterminer quelle part de ces temps la
préparation incrémentale peut réellement supprimer, net du coût de son cache.

- **G00 → G01 → G02 → G03 → G04 → G05 → G06** : graphe réel, partition,
  références qualifiées, émission, FFI/bootstrap, cache/runner, livraison.
- **P00 → P01 → P02 → P03 → P04 → P05 → P06** : référence, entrées persistantes,
  métadonnées, demandes directes, point fixe, préparation sélective, livraison.
- Chaque suite s'appuie sur le cache et les validations de v3. G n'exige pas P,
  et P fonctionne avec le layout Go retenu lors de son démarrage. Si G est livré
  en premier, P00 prend G06 comme référence immédiate et conserve celle de v3/06.
- Chaque lot produit un candidat identifié et ses preuves. Les changements
  préparatoires d'architecture ont un critère de correction ; l'activation d'un
  chemin optimisé exige en plus le seuil de performances de la suite.

### Protocole de mesure et seuils d'acceptation

- Figer le corpus b8x, les compilateurs, les commandes, les ressources et les
  états de cache. Reprendre original, feuille, partagée et FFI ; ajouter les
  transitions ciblées décrites ci-dessous. Conserver chaque essai et son retour
  à la référence en cas de rejet.
- Mesurer **au moins trois paires alternées**, sans profilage ni autre campagne
  lourde concurrente. Comparer les médianes, conserver chaque paire et augmenter
  l'échantillon si la dispersion empêche de conclure. Les longues campagnes
  tournent en arrière-plan.
- **Suite P :** minimum d'acceptation d'au moins **10 % de réduction du temps
  gopurs total** sur chacune des éditions feuille et partagée. Chaque optimisation
  autonome doit gagner au moins **5 %** sur son scénario annoncé avant l'essai.
  Les coûts de lecture, décodage, preuve et publication font partie de ce temps.
  L'objectif reste un retour vers **~30 s après modification** : depuis G06,
  cela demande environ **38–40 %** sur feuille/partagé. Le minimum de 10 %
  laisserait encore environ **44–45 s** ; son atteinte ne suffit donc pas à
  déclarer cet objectif rempli.
- **Dette froide du backend :** conserver explicitement le repère **~60 s**
  face aux **~118 s** de G06. P00 doit chiffrer la contribution possible de P
  à cette réduction ; la réutilisation après édition ne garantit pas un gain
  au premier build. Le respect du budget de régression au froid laisse cette
  dette ouverte tant que le retour vers le repère initial n'est pas établi.
- **Suite G :** objectif d'acceptation d'au moins **30 % de réduction du temps
  `go build`** et **20 % sur la chaîne mesurée jusqu'à `exec`**, pour chacune des
  éditions feuille, partagée et FFI. Garder initialement les ressources du runner
  à **`GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB`** pour attribuer le gain au découpage.
- Pour chaque suite, budget de régression de **5 %** sur les autres scénarios,
  notamment le froid. Pour la relance identique, budget absolu/plancher de
  **max(5 % de la référence, 0,100 s)**. Enregistrer en P00/G00 les budgets de
  mémoire et de stockage avant les essais ; une hausse doit être chiffrée dans
  le bilan d'acceptation. Ces seuils sont des critères, pas des gains prédits.
- Pour G, chaque échantillon part d'un `GOCACHE` privé préchauffé sur l'état de
  base seulement ; la variante modifiée doit encore demander sa compilation.
  Préparer cet état séparément pour chaque générateur et conserver les traces
  des compilations/liens réellement exécutés. Mesurer explicitement le froid
  applicatif et le froid complet ; conserver les caches de travail de b8x.
- Chronométrer directement la chaîne jusqu'à `exec`, puis l'exécution des tests.
  Relever aussi frontend, gopurs, tidy, attente du verrou, compilation, lien,
  RSS, allocations, octets persistés, fichiers réécrits et packages recompilés.
- **Priorité utilisateur révisée le 14 octobre, pendant le pilote G06/3 :** le
  temps prime ; le RSS est secondaire sauf aggravation démesurée. Le seuil
  historique de +5 % RSS sert désormais d'indicateur, sans rejet automatique.
  Continuer à mesurer backend et compilateur Go séparément et documenter le coût
  mémoire. Aucun nouveau plafond numérique n'a été fixé. Conserver les anciens
  protocoles/résultats ; les seuils temporels et de stockage restent applicables.

### Suite P — Préparation incrémentale : 1/7 lot validé (audit P00)

**Point de départ :** `Driver.Prepare.prepareModules` charge tous les TAST,
reconstruit les métadonnées source puis appelle
`Monomorphization.monomorphizeModulesWithDispatch`. Le PBO dispose d'un cache
de travail interne à `transitiveCollectWith`, limité à son point fixe courant.
Le lot 04 reprend ensuite les résultats d'optimisation/émission et certaines
empreintes préparées. La suite P doit produire les mêmes modules préparés et
les mêmes métadonnées en réévaluant seulement les contributions invalidées.

#### P00 — Référence et dépendances de préparation

**Terminé le 19 octobre 2026 :** [protocole, mesures et décision proposée](docs/build-preparation-baseline.md).
Les anciennes campagnes temporaires sont absentes ; JS/Rust/WASM correspondent
à G06, mais le natif installé a changé d'empreinte. La référence P00 est
reconstituée et vérifiée dans une copie privée avant toute attribution de gain.
**Reprise autorisée le 19 octobre après coupure :** les preuves temporaires P00
ont disparu. Les scripts restent présents ; la reconstitution utilise désormais
`htdocs/.gopurs-evidence/p00-20261019/`, la révision b8x initiale et un relevé de
concurrence. Les froids du premier passage étaient trop dispersés pour être
retenus. La reprise durable termine 36 passages appariés, les sondes codecs/
publication, 28 transitions quatre hôtes et la validation b8x cache/forcé/hit.

**Bilan :** référence backend **45,352 / 47,047 / 35,056 s** après
feuille/partagé/FFI, **115,692 / 118,837 s** aux froids. Le jeu préparé +
métadonnées + appelants + états transitifs ajoute **809,32 Mio**, restaure en
**12,099 s** et ne laisse que **2,313 / 2,233 s** de gain avant preuves/recalcul,
donc moins que le minimum de 10 %. Sa publication froide coûte **91,684 s**
isolément. Le modèle partiel préparés/métadonnées laisse **9,418 / 9,271 s**
avant preuves, mais ne finance pas le graphe complet actuel. Le potentiel d'une
preuve compacte reste indéterminé ; même ce modèle partiel laisse **35,450 /
37,804 s** avant preuves, et P seul ne résout pas le froid vers ~60 s.
**Recommandation P00 : revoir P01 avant tout nouveau candidat.** Budgets proposés,
unités écartées et effort figurent dans le bilan ; l'étude resserrée décrite
ci-dessous a ensuite été autorisée par l'utilisateur.

- [x] **Délimiter les unités réutilisables, leur coût complet et le gain accessible.**
  - Établir la référence sur le compilateur disponible après G06 (référence
    reconstituée, empreinte native différente et provenance explicite), avec layout packages
    explicite et protocole figé : froid applicatif/complet, identique, feuille,
    partagée et FFI. Conserver les références v3/06 et les repères initiaux.
  - Ventiler `Driver.Prepare`, les constructeurs de métadonnées,
    `Gopurs.Monomorphization` et le collecteur transitif du PBO : chargement,
    enrichissement des classes, index des corps/types, demandes directes,
    tours transitifs, réécriture et calcul des clés.
  - Rapprocher cette ventilation du temps backend total, en isolant aussi les
    coûts du cache, de l'émission et de l'analyse/partition Go. Éviter les doubles
    comptes et conserver les coûts non attribués ; les anciens 9–10 s de
    préparation ne suffisent pas à expliquer les 40–50 s actuelles après édition.
  - Inventorier les lectures de corps, types, dictionnaires, arguments statiques,
    présences/absences et barrières d'éligibilité de chaque passe. Conserver les
    contributions intermédiaires et leurs empreintes comme oracles.
  - Chiffrer par scénario le temps évitable brut, les coûts de restauration,
    validation, sérialisation et publication des unités envisagées, puis le
    **gain net plausible en secondes et en pourcentage** et le **temps résiduel
    estimé**. Distinguer mesures, bornes et hypothèses, avec leur incertitude ;
    confronter ce potentiel au minimum de 10 %, à l'objectif ~30 s après édition
    et à la dette froide ~60 s.
  - Ajouter au corpus les retraits de demandes, substitutions à cardinalité
    égale, cycles, changements de fournisseur et transitions FFI seules.
  - **Bridges :** recenser les transports de `Preparation`, `Wire`, du cache
    commun et des observations PBO nécessaires à ces données.
  - **Fin :** référence reproductible, graphe des lectures, budget de stockage/
    restauration et bilan du gain net accessible ; choix des unités persistées
    motivé par les mesures et l'effort d'implémentation attendu. Présenter ce
    bilan à l'utilisateur pour une **décision explicite avant P01** : poursuivre,
    ajuster le périmètre ou s'arrêter. Si le potentiel est insuffisant ou reste
    indéterminé, exposer ce constat avant toute nouvelle série de candidats.

#### P01 — Entrées et objets de préparation persistants

**Étude resserrée autorisée et terminée le 19 octobre après P00 :**
[objets compacts et publication regroupée](docs/build-preparation-compact.md).
Un candidat **C01** (64 groupes stables par famille, gzip rapide), 96 mesures
appariées de stockage et clé globale conservatrice. Les seuils temps/stockage
de P00 ont été repris avant les essais. **C01 rejeté en l'état ; P reste à 1/7.**

Le compact regroupé publie initialement en **2,947 s** au lieu de **90,972 s**
et occupe **91,93 Mio** courants. Mais extraction/décodage coûtent **11,264 s**
(budget 6,5 s), les coûts connus après feuille/partagé **11,784 / 11,964 s**
(budget 11 s avant même preuves/recalcul), et le surcoût froid favorable
**7,349 / 7,155 s** (budget 5,785 / 5,942 s). Le modèle ne laisse que
**4,405 / 4,206 s** de gain avant preuves, soit **9,85 / 9,10 %**, avec résiduels
**40,464 / 42,869 s**, loin du repère ~30 s. Le froid vers ~60 s reste ouvert.

**Recommandation : recadrer le volume restauré et le chemin complet avant
intégration ou nouveau candidat.** Le gain local de publication est conservé
comme résultat de l'étude ; les preuves sélectives et les replays quatre hôtes
de la nouvelle représentation restent requis pour une future intégration.
Nouvelle décision utilisateur avant la suite.

- [ ] **Identifier et charger les entrées au niveau du module.**
  - Prolonger l'inventaire de contenu TAST avec les identités et versions des
    passes, les options/directives consommées, les chemins et les annotations
    utiles aux diagnostics. Traiter ajout, retrait, renommage et ordre canonique.
  - Définir les objets source/préparation et leurs preuves dans le stockage
    commun : codecs portables, intégrité, staging privé, engagement atomique et
    reprise. Permettre aux contributions validées d'éviter le décodage source.
  - Mesurer la restauration avant de retenir un objet représentant le TAST
    décodé : les coûts d'encodage, de lecture et de GC peuvent dépasser ceux du
    TAST original. Garder le chargement complet comme oracle et repli explicite.
  - **Bridges :** étendre les contrats `BuildCache` et `Wire` seulement pour les
    objets retenus, avec octets/protocoles identiques en JS, Go et Rust.
  - **Fin :** inventaire différentiel exact, codecs éprouvés, erreurs de décodage
    et diagnostics conservés ; objets absents/incompatibles/corrompus recalculés.

#### P02 — Métadonnées source par contributions

- [ ] **Remplacer les reconstructions globales par des contributions remplaçables.**
  - Extraire par module types globaux, constructeurs, champs de classes,
    élisions, enums et ADT pointeurs ; dériver les index Rebox depuis les bonnes
    déclarations. Préserver les priorités et l'ordre des fusions actuelles.
  - Garder la provenance source de l'ABI : enrichissement des classes avant les
    dictionnaires, ADT originaux pour enums/élisions et `Any` explicite conservé.
    Une copie spécialisée ne devient pas la déclaration de référence.
  - Remplacer ou retirer la contribution de chaque module changé ; invalider
    les dérivations qui lisent une entrée modifiée, ajoutée ou disparue.
  - **Bridges :** encoder les nouvelles contributions avec les codecs communs ;
    contrôler les layouts que consomment génération, Rebox et FFI.
  - **Fin :** toutes les tables et leurs empreintes égalent la reconstruction
    globale, y compris sur changement/suppression de type, classe ou module.

#### P03 — Demandes directes et preuves de lecture

- [ ] **Rendre remplaçable la contribution de chaque appelant.**
  - Faire produire à `collectInstantiations` les demandes directes par module,
    leur provenance et les lectures qui les justifient : corps externes,
    absences, types, dictionnaires et arguments statiques.
  - Réévaluer aussi un appelant inchangé si une lecture de fournisseur change.
    Indexer les lecteurs et retirer l'ancienne contribution avant la fusion
    canonique ; préserver la sélection du payload et l'union des appelants.
  - Inclure les dépendances des barrières FFI, forwarders, intrinsics et workers
    record partagés ; conserver leur position dans la pipeline actuelle.
  - **Bridges :** raccordement des observations dans le PBO et codec des preuves ;
    l'ordonnancement parallèle publie les résultats dans l'ordre canonique.
  - **Fin :** demandes directes et diagnostics égaux au calcul complet pour
    changement de corps à type stable, ajout/retrait et présence/absence.

#### P04 — Point fixe transitif incrémental

- [ ] **Recalculer les spécialisations affectées, avec rétraction correcte.**
  - Persister les contributions des couples définition/spécialisation, les
    entrées effectivement lues et leurs dépendances inverses ; propager une
    file de travail jusqu'au même point fixe que `transitiveCollectWith`.
  - Traiter les retraits et cycles en réévaluant la région affectée depuis des
    demandes encore justifiées. Éliminer les spécialisations qui ne subsistent
    que par soutien mutuel après la disparition de leur dernier appelant.
  - Les raccourcis actuels de cardinalité reposent sur la croissance monotone
    d'un même calcul à corps fixes. Entre builds, vérifier le contenu et les
    lectures nécessaires, y compris une substitution à cardinalité identique.
  - Conserver snapshots immuables par tour, barrières et fusion déterministe ;
    un cas dont la preuve est insuffisante repasse par le calcul complet.
  - **Bridges :** adapter le collecteur et ses observations PBO ; éprouver les
    transports et l'évaluation des jobs dans les quatre hôtes.
  - **Fin :** point fixe, payloads et appelants exacts sur les transitions
    directes/transitives, cycliques et séquentielles/parallèles.

#### P05 — Réécriture sélective et raccordement au lot 04

- [ ] **Restaurer les modules préparés dont toutes les entrées sont prouvées.**
  - Cléer la réécriture par source, demandes applicables, corps/types lus,
    barrières et contexte ; reprendre les modules spécialisés et leurs clés
    seulement sous cette preuve. Une demande nouvelle peut changer un fournisseur
    dont le TAST est identique, comme `Control.Semigroupoid` dans le corpus b8x.
  - Assembler `PreparedModules` dans l'ordre canonique et fournir au lot 04 les
    mêmes clés préparées, métadonnées et contributions que le parcours complet.
  - Prévoir la reprise de toute la préparation lorsque ses entrées sont
    identiques, notamment pour une édition limitée au corps d'une FFI. Vérifier
    séparément les entrées qui déterminent les points d'entrée et les diagnostics.
  - Intégrer `--force`, les motifs d'invalidation et les compteurs par passe ;
    comptabiliser lecture, validation et publication dans la mesure du gain.
  - **Bridges :** raccorder session/cache/préparation en JS, Go et Rust ; réutiliser
    les barrières de publication et preuves d'émission du lot 04.
  - **Fin :** préparation, quatre parties du cache par module et sorties Go
    byte-exactes face au build forcé ; seuils P atteints sur les paires annoncées.

#### P06 — Robustesse, mesures et livraison

- [ ] **Valider la préparation incrémentale dans le parcours complet.**
  - Rejouer v3/05 avec les transitions spécifiques P00–P05, les corruptions,
    interruptions, reprises et accès concurrents aux nouveaux objets.
  - Reconstruire et tester les compilateurs réellement installés : mêmes TAST
    pour JS, Go séquentiel/parallèle et le bootstrap Rust ; comparer les modules
    préparés, sorties, diagnostics, puis compiler et exécuter les régressions.
  - Mesurer froid, identique, feuille, partagée et FFI, taille du cache, mémoire,
    lectures et recalculs. Vérifier les seuils face à P00/G06 et le bilan cumulé
    face à v3/06 ; publier aussi l'écart restant aux repères ~30 s après édition
    et ~60 s à froid. Examiner les snapshots puis effectuer les replays stricts.
  - **Fin :** gains confirmés, minimum d'acceptation et objectif utilisateur
    évalués séparément, dette froide restante explicitée, preuves conservées et
    documentation de préparation, cache, invalidations et récupération publiée.

### Suite G — Granularité des packages Go : 7/7 lots validés

**Point de départ :** `Driver.Output` écrit modules et FFI dans `purescript` et
les entrées appellent `purescript.Get_…_main`. `GoAst`, `GoImports` et `Printer`
connaissent déjà les imports externes, mais de nombreuses références internes,
noms de types et fragments Go supposent un espace de noms commun. Le découpage
doit couvrir les dépendances après spécialisation/PBO, les layouts, helpers,
FFI et enregistrements du runtime. Les seuls imports PureScript sont insuffisants.

#### G00 — Graphe généré et coût réel de recompilation

**Validé le 13 octobre 2026 :** le [relevé typé](docs/build-go-packages.md)
couvre la base b8x, les trois éditions et le bootstrap livré. Le découpage par
module laisse **94,53 %** du code dans un bloc cyclique. Le graphe par déclaration
ramène les fermetures statiques des éditions à **0,0006 / 6,12 / 11,48 %** des
octets, avant regroupement et contraintes d'ABI/initialisation. Ces tailles ne
sont pas des gains de compilation. Budgets fixés avant les candidats : **+5 %
mémoire / +10 % par catégorie de stockage**. Le froid backend exact est relevé
à **113,687 s / 9,841 GiB de RSS** ; après une pause disque sous 8 GiB et
libération d'espace, les sept passages mémoire sont exacts avec la même
racine/journal. Les sept contrats Go et la réflexion passent. Le stockage est
audité. Le froid applicatif Go prend **255,780 s**, avec deux compilations/un
lien et le binaire exact ; l'identique suivant ne compile ni ne lie, et l'oracle
passe. Ce passage isolé, plus lent que l'ancien froid intégral de **221,830 s**,
ne remplace pas les comparaisons appariées. Audit final : **430 fichiers de
production inchangés**, preuve `gopurs-go-packages-g00-20261012/final-reference/`.

- [x] **Identifier les dépendances et les unités coûteuses sur le Go final.**
  - Relever déclarations, références de valeurs/types, workers, getters, helpers
    Rebox, initialisations et FFI, y compris dans `GoRaw` et les fragments opaques.
    Rapprocher l'inventaire de l'AST Go et de sa résolution de symboles.
  - Mesurer composantes cycliques, fan-out, taille des groupes et portée réelle
    des trois éditions b8x ; relever les dépendances introduites par l'inlining
    et les spécialisations. Tracer les actions de compilation et de lien Go.
  - **Bridges :** inventorier les références Go des FFI applicatives, du runtime,
    des bibliothèques et du bootstrap ; relever les usages de `StructGetters`,
    des tags et de la réflexion qui échappent aux appels statiques.
  - **Fin :** graphe audité, référence temporelle actuelle et budgets mémoire/
    stockage ; potentiel de partition étayé par la taille des régions affectées.

#### G01 — Propriété des symboles et partition déterministe

**Validé le 13 octobre 2026 :** le [plan canonique](docs/build-go-package-plan.md)
retient les intervalles topologiques à ancres logiques **1/128**, les FFI
entières et les contraintes d'ordre Go. Il donne **525 packages b8x / 158 pour
le bootstrap**, avec **15 DAG** vérifiés indépendamment, permutation des
identifiants stable, **95 oracles SCC** et **sept programmes** d'initialisation.
Les corps feuille/FFI conservent le placement ; le partagé déplace **1 818
déclarations inchangées**. Les fermetures après regroupement représentent
**38 792 440 / 63 388 578 / 65 252 780 octets**, sans gain Go attribué.
Les huit familles exploratoires, l'initialiseur central et ses erreurs de
harnais restent conservés. Le rendu effectif et l'acceptation temporelle restent
G02–G06. Audit final : **430 fichiers de production inchangés**, preuve
`gopurs-go-packages-g01-20261013/final-plan/`.

- [x] **Définir un plan de packages acyclique et stable.**
  - Attribuer un propriétaire aux valeurs, types, spécialisations et helpers ;
    partir des dépendances réellement générées et regrouper les composantes
    fortement connexes. Mesurer le coût des groupes et de leurs dépendants.
  - Définir les exports requis, la localisation des layouts partagés et les
    règles des helpers qui dépendent de plusieurs types. Toute extraction de
    déclarations partagées doit conserver l'identité des types et réduire le
    coût mesuré de la région de recompilation.
  - Stabiliser noms de packages, chemins et alias à partir des identités logiques.
    Définir fusion/scission et collisions ; une modification de corps sans
    changement de graphe doit conserver le plan de placement.
  - **Bridges :** inclure les arêtes et contraintes de visibilité FFI/runtime
    dans ce plan, ainsi que les appels aux getters et workers natifs.
  - **Fin :** plan canonique, DAG vérifié et estimation documentée de la portée
    de recompilation ; les grosses composantes résiduelles sont expliquées.

#### G02 — Références et types Go qualifiables

**Validé le 13 octobre 2026 :** les [références possédées](docs/build-go-references.md)
couvrent AST, workers, types, rendu et liaison typée des fragments/FFI. Les cinq
corpus G00 sont résolus ; b8x et bootstrap représentent **2 997 / 625 unités**.
Les **18 passages b8x** sont byte-exacts sur les mêmes TAST, Rust inclus dans
la campagne finale. Les objets workers v2 sont exacts entre hôtes. Les trois
transports du résolveur, les race detectors, les replays stricts et la reprise
v1 → v2 à journal absent passent. Les trois snapshots restent inchangés.
Le relevé JSON exhaustif n'est pas publié au cache ; G03 doit raccorder un plan
compact au pilote. Aucun gain temporel n'est revendiqué. Preuve :
`gopurs-go-packages-g02-20261013/final-reference/`.

- [x] **Porter la propriété des références jusqu'au rendu.**
  - Introduire les identités propriétaire/symbole nécessaires dans `GoAst`, les
    métadonnées et signatures de workers ; qualifier au rendu selon le package
    consommateur. Mettre à jour `GoImports`, `Printer` et leurs producteurs.
  - Couvrir aussi arguments de types, assertions, constructeurs, noms de champs,
    callbacks, helpers et fragments opaques. Rendre explicites leurs dépendances
    générées ; traiter le Go FFI avec son parseur et la portée des symboles.
  - Distinguer identité PureScript, identité du type Go et tag runtime ; les
    conventions de labels/record et la réflexion font partie de ce contrat.
  - **Bridges :** adapter le contrat du parseur FFI si l'inventaire l'exige, avec
    parité WASM/appel Go direct et transport Rust ; auditer les fragments des FFI
    du compilateur qui embarquent eux-mêmes des noms générés.
  - **Fin :** toutes les références sont attribuées ; le rendu du layout
    monolithique reste byte-exact sur le corpus et les fixtures existants.

#### G03 — Émission multipackage, ABI et initialisation

**Validé le 14 octobre 2026 :** l'[émetteur typé et le writer](docs/build-go-emission.md) portent
le plan 1/128, les annotations de records/initialiseurs et les trois transports.
Les fixtures compilent et s'exécutent, avec oracles SCC, ordre des effets et
cellule lazy partagée. La campagne réelle termine **12/12 passages exacts** sur
les quatre hôtes, y compris `StringEdgeCases` (**271 TAST / 33 packages**). Les
reproducteurs du trap WASM, du tampon JS et des annotations UTF-16 sont conservés.
Le pilote complet acquiert les exports compacts et publie par le journal ; les
quatre hôtes passent migration, entrées, forcé, dépendance locale et identique
sans compilation/lien. b8x publie **39 613 Go / 525 packages** et son oracle
compilé passe, avec reconstruction identique sans compilation/lien. Le bootstrap
courant produit **150 packages**, compile et passe le test CLI comme compilateur.
Le rapprochement des graphes émis retrouve **185 667 placements / 35 596 arêtes**
b8x et **65 964 placements / 5 875 arêtes** bootstrap ; les renommages de helpers
et différences de corpus sont expliqués. Après revue, les deux snapshots ne
reçoivent que **27 / 134 lignes d'annotations**, puis les **67 commandes strictes**
des trois fixtures passent sur les quatre hôtes. Les **huit comparaisons / 28
commandes** des grands corpus sont byte-exactes sur les quatre hôtes, avec quatre
objets de modules exacts, hits sans réécriture et reconstructions sans compilation
ni lien. Les **348 tags** b8x et les **185 667 corps de déclarations**, normalisés
pour les seuls renommages prévus, restent équivalents. Les sources du compilateur
figé sont auditées dans `gopurs-go-packages-g03-20261013/final-reference/`.

- [x] **Émettre un programme compilable selon le plan G01.**
  - Raccorder les chemins, imports, alias et entrées de `Driver.Output` au plan ;
    choisir une visibilité légale pour chaque valeur, type et champ traversant
    une frontière Go. Couvrir notamment les records à champs non exportés et
    leur identité de type, les génériques, ADT pointeurs et conversions Rebox.
  - Conserver la sémantique des getters paresseux/`sync.Once`, la récursion et les
    enregistrements `StructGetters`. Rendre explicites les dépendances d'init,
    y compris celles nécessaires quand une entrée ne référence qu'une partie
    des modules ; vérifier collisions et ordre des enregistrements.
  - Émettre déterministement les cycles regroupés et les références introduites
    par l'optimiseur. Vérifier que le placement des helpers respecte encore le
    plan, puis contrôler le graphe avec `go list` et `go build`.
  - **Bridges :** mettre à jour qualifications et accès des wrappers générés,
    tout en vérifiant tags, boxing/unboxing, callbacks et ABI des workers.
  - **Fin :** les fixtures de dépendances, cycles et types interpackages compilent ;
    exécution, diagnostics et initialisation sont équivalents à la référence,
    avec octets identiques entre hôtes pour le layout multipackage.

#### G04 — FFI réelles, bibliothèques et bootstrap

**Validé le 14 octobre 2026 :** les [consommateurs réels](docs/build-go-consumers.md)
passent **36 passages / 207 commandes strictes** sur les quatre hôtes, avec les
mêmes TAST et les résultats du monolithe. Les neuf snapshots multipackages ont
été examinés (**592 déclarations / 57 fragments**), créés après exécution puis
rejoués sans modification. Les **89 commandes** CLI couvrent FFI absente,
syntaxe invalide, retour de FFI et signatures/dépendances natives. Les suites
`foreign`/`js-promise`, les **onze probes réfléchies** b8x et le bootstrap courant
sont verts. Preuve : `gopurs-go-packages-g04-20261014/final-reference/`.

- [x] **Faire fonctionner les consommateurs avec les frontières de packages.**
  - Placer chaque FFI avec son propriétaire et réécrire par AST les références
    qui doivent traverser un package, en respectant variables locales, imports,
    noms exportés, helpers et alias. Vérifier corps/signatures, FFI absentes,
    erreurs de syntaxe et dépendances natives externes.
  - Rejouer les fixtures et bibliothèques utilisant workers natifs, callbacks,
    tableaux, records, types paramétrés et réflexion ; adapter les seules FFI
    dont le contrat dépend effectivement du package monolithique.
  - Raccorder `build-native.mjs`, `prepare-native-output.mjs` et les outils de
    tests aux nouveaux chemins ; éprouver le bootstrap Go et le compilateur
    Rust qui génère ces mêmes packages, avec artefacts privés identifiés.
  - **Bridges :** audit avant/après des bridges gopurs, bibliothèques et b8x,
    ainsi que des packages natifs `gopurs_ffi_parser` et `gopurs_build_cache`.
  - **Fin :** bootstraps et consommateurs réels verts, sources FFI et sorties
    appariées, aucune référence générée non résolue ou collision de noms.

#### G05 — Cache, migration des sorties et runner b8x

**Validé le 14 octobre 2026 :** les [sources avant partition](docs/build-go-incremental.md)
sont conservées sous `gopurs/go-source-gzip-v1`. Restauration, intégrité,
récupération et migrations passent avec détection de races, Go 1.22 inclus.
Les **56 comparaisons / 350 commandes** de transitions réelles passent sur les
quatre hôtes. L'artefact Rust incorrect reste conservé et son remplacement est
validé. b8x restaure **2 703 émissions** après une édition FFI et ne réécrit qu'un
Go, avec forcé/retour exacts, oracle et identique sans compilation/lien. Les dix
contrôles du runner passent, avec migration de layout, dépendance transitive et
deux entrées. La récupération finale des objets décodés est intégrée au candidat
1 G06, qui passe **89 commandes CLI** sur les quatre hôtes. Preuve :
`gopurs-go-packages-g05-20261014/final-reference/`.

- [x] **Rendre le nouveau layout incrémental et récupérable.**
  - Versionner le layout dans le contrat de cache ; couvrir dans les clés
    d'émission le propriétaire du module, les placements/ABI référencés et les
    règles de nommage. Une édition isolée doit garder les sorties des packages
    indépendants byte-identiques et éviter une invalidation globale de placement.
  - Publier les fichiers possédés et retirer les anciennes sorties au travers
    du staging/journal existants. Éprouver migration monolithique/multipackage,
    fusion/scission de groupe, suppression, interruption et reprise ; conserver
    les fichiers externes, l'inode du verrou et la propriété de `go.mod`/`go.sum`.
  - Raccorder sélection des mains, clés du runner et artefacts exécutables b8x ;
    vérifier verrou, publication atomique, annulation et exécutions concurrentes.
  - **Bridges :** adapter les inventaires et transports du cache aux nouveaux
    chemins ; contrôler les FFI déplacées avec leurs modules et signatures.
  - **Fin :** incrémental et complet donnent le même inventaire/octet par layout ;
    zéro compilation/lien à la relance identique et recompilation limitée aux
    actions réellement invalidées par Go, tracées sur chaque transition.

#### G06 — Mesures finales, snapshots et livraison

**Validé le 16 octobre 2026 :** le candidat 7 à densité **1/8192** est installé
avec ses artefacts exacts. Les **36 chaînes / 36 oracles** et **39 critères
temps/stockage** passent. La chaîne gagne **62,88 % feuille / 48,36 % partagé /
54,50 % FFI**, malgré les surcoûts backend de **26,88 / 34,37 / 42,62 %**. Le
backend reste à **~118 s à froid / 40–50 s après édition**, loin des repères
initiaux **~60 s / ~30 s** respectivement ; face à `publicsync`, le froid complet
est à **+2,23 %**, dans le budget. Le GOCACHE identique est à
**+9,997897 %**. Le [bilan complet](docs/build-go-measurements.md) publie les
paires, plages, ressources, graphes, actions Go et commandes de migration.

La parité des deux grands corpus passe sur les quatre hôtes, comme les **56
comparaisons de transitions**. Les **612 déclarations** des neuf snapshots sont
revues avant mise à jour, puis les **207 commandes strictes** et **89 commandes
CLI** passent sur l'installation. Intégrité/récupération du cache, bootstrap
courant (**512 TAST / 23 Go / binaire exact**) et **10 contrôles du runner sur
chacun de macOS et Linux** passent. L'audit final porte sur les artefacts mesurés,
295 fichiers source/outils, les snapshots et les preuves conservées :
`gopurs-go-packages-g06-20261014/final-reference/summary.json`.
**G est clos à 7/7 ; P reste à 0/7, en pause.**

**Historique des candidats et reprises, 14–16 octobre :** un alias runtime court,
résolu par `go/types` avec contrôle des
noms réservés, vise le dépassement de stockage relevé sur G03 (**+16,82 %**).
Le résolveur et les quatre hôtes CLI passent. Le
[protocole G06](docs/build-go-measurements.md) conserve les seuils, la référence
`publicsync`, les ressources du runner et les GOCACHE privés de base.
Le pilote 1 gagne **83,63 % de Go / 60,22 % de chaîne** sur la feuille (une
paire seulement), mais son candidat est rejeté : froid backend, RSS froid et
GOCACHE dépassent les budgets. Le candidat 2 regroupe les fichiers par package,
avec contrôle typé des imports et maintien des frontières requises par les
directives. Il passe **56 comparaisons de transitions / 89 commandes CLI** sur
les quatre hôtes. Les **612 déclarations Main/FFI** des neuf snapshots sont
examinées avant mise à jour, puis **207 commandes** les rejouent strictement sur
les quatre hôtes. Preuve : `functional-candidate-2/summary.json`. Le pilote 2
termine avec les volumes après oracle et les caches d'exports isolés : feuille
**−84,71 % Go / −66,46 % chaîne**, mais GOCACHE de base **+57,07 %**, RSS backend
froid **+12,13 %**, gopurs froid **+22,54 %** et chaîne identique **+0,724 s**.
Le candidat 2 est rejeté pour l'acceptation. Les diagnostics de mémoire et
d'archives sont conservés ; la restitution mémoire à la frontière du résolveur
est rejetée à **+6,25 % de RSS** dans le diagnostic, puis le partage des
identifiants à **+5,94 %**, avec sorties exactes et retour audité au candidat 2.
Les probes 1/512 et 1/2048 passent leurs oracles, mais dépassent encore le budget
GOCACHE runner (**+22,94 % / +16,61 %**). Le probe 1/8192 passe l'oracle avec
**6 packages et +9,59 % de GOCACHE runner** ; le cache total reste à mesurer.
Les exports allégés reproduisent les sources exactement et passent les suites
cache/race Go 1.27.0 et 1.22.12. Le candidat 3 privé réunit ces deux essais et
passe le résolveur, **89 commandes CLI / 56 comparaisons de transitions** sur
les quatre hôtes. Son pilote termine avec **10 oracles passants** : chaîne
feuille **−62,80 %**, partagée **−44,99 %**, FFI **−52,88 %**, identique
**−6,23 %**, sur une seule paire. Restent hors budget le backend du seed froid
(**+10,01 %**), le GOCACHE total de base (**+10,005 %**) et les objets gopurs après
partagé/FFI (**+21,09 % / +26,22 %**). Un entretien des objets non référencés après
publication et une réduction des copies du rendu sont en validation. L'essai
GC à 100 % est rejeté : **+3,59 % de temps** sur trois paires locales malgré la
baisse de RSS. La réduction des copies gagne **10,06 %** sur trois paires de
partition seule, avec sorties exactes ; les exports sans chemins locaux
économisent **85 714 octets** de cache dans le probe. Les suites cache/race
passent sous Go 1.27.0 et 1.22.12. Le candidat 4 privé est reconstruit et audité ;
il passe **89 commandes CLI / 56 comparaisons de transitions** sur les quatre
hôtes. Les trois paires du contrôle froid ciblé donnent **112,674 → 120,893 s
(+7,29 %)**, au-dessus du budget dans chaque paire, avec six inventaires exacts :
le candidat 4 est refusé au contrôle froid. La réutilisation des compresseurs de
sources G05 est rejetée (**+0,57 % de temps** local malgré moins d'allocations),
avec retour byte-exact au code antérieur. L'acquisition d'exports à quatre
processus gagne **60,41 %** sur trois paires froides locales, avec mêmes exports,
cache et Go ; les suites cache/race passent sous Go 1.27.0 et 1.22.12.
Le candidat 5 privé est reconstruit, audité et passe **89 commandes CLI /
56 comparaisons de transitions** sur les quatre hôtes. Ses trois paires froides
donnent **110,519 → 117,662 s (+6,46 %)**, avec dépassement dans chaque paire et
six inventaires exacts : il est refusé au contrôle froid, à **1,616 s** de la
limite médiane. Les profils CPU/allocations/GC réellement froids terminent avec
sorties exactes ; ils servent au diagnostic. Le Go du driver conserve encore
la préparation entière après émission pour relire les seules entrées : cette
piste est testée par un probe natif privé, construit avec avant/après conservés.
Il est **rejeté** après trois paires : **−2,87 % après émission**, sous le seuil
préalable de 5 %, malgré **−3,40 % de backend total** en médiane ; six inventaires
exacts, une légère régression totale dans la deuxième paire. Le probe d'exports
à huit processus gagne ensuite **15,22 %** sur trois paires froides locales,
avec mêmes exports, cache et Go ; son raccordement est en validation. La séquence
du candidat 6 termine : suites cache/race des deux Go, reconstruction des
quatre hôtes, audit et **89 commandes CLI / 56 comparaisons de transitions**
passants. Le contrôle froid à trois paires passe sur les médianes,
**112,299 → 117,026 s (+4,21 %)**, avec six inventaires exacts ; la deuxième paire
reste à **+5,18 %**, et la marge médiane n'est que de **0,888 s**. La campagne
complète `full-1/` de **36 chaînes** est lancée : trois paires alternées par
scénario, archivage vérifié entre scénarios et ressources G00/v3 conservées.
Elle est interrompue par **SIGTERM à 28/36 chaînes**, toutes exactes avec leurs
oracles, pendant le troisième froid applicatif. Sur trois paires complètes,
la chaîne gagne **62,16 % feuille / 46,96 % partagé / 53,21 % FFI**, malgré un
surcoût backend ; les seuils Go/chaîne et le stockage chauds passent. Le GOCACHE
identique reste juste sous la limite, à **+9,9979 %**. Les deux paires froides
applicatives donnent **−31,17 % de chaîne / +3,18 % backend**, encore provisoires.
La première reprise des huit mesures restantes utilise les fonctions du
harnais figé, racines neuves pour la paire interrompue et sauvegarde des
28 mesures acquises. L'entretien préalable archive les quatre échantillons
froids terminés, hors chronométrages ; la couverture des six échantillons sera
vérifiée entre les archives avant/après reprise. `caffeinate` protège la
séquence contre la veille automatique. Elle termine à **33 chaînes avec oracles
passants**, puis le garde-fou disque arrête la suivante avant `tidy`. Le froid
applicatif passe sur trois paires : **−31,31 % chaîne / −49,52 % Go / +1,53 %
backend** en médiane. La première paire de froid complet donne **−27,42 % de
chaîne**, mais **+7,74 % de backend** : les deux autres paires restent nécessaires.
La deuxième reprise est lancée avec archivage après chaque paire froide. Les
33 chaînes complètes sont sauvegardées ; le côté candidat de la paire
interrompue est mis à part pour rejouer cette paire entièrement. Le jeu retenu
repart de **32 chaînes**, puis termine à **36 chaînes / 36 oracles passants**.
Les TAST/Go sont exacts ; les 28 résultats initiaux, les 32 retenus avant reprise
et le côté isolé sauvegardé sont vérifiés. Les archives couvrent les six
échantillons retenus et le côté isolé. Le froid complet donne **−27,42 % chaîne /
−45,86 % Go**, mais **117,121 → 125,117 s (+6,83 %) de backend**, seul critère
temps/stockage en échec. Les écarts backend par paire sont **+7,74 / +4,90 /
+2,90 %** : l'extension prévue en cas de dispersion est fixée à **trois paires
supplémentaires**, avec décision sur **les six réunies** et mêmes seuils.
Le harnais réutilise ses 15 fonctions et son environnement par AST, avec racines
neuves et archivage après chaque paire. Les 36 résultats/analyses sont conservés ;
le bilan cumulé attend 42 chaînes. Réserve préalable de **16 GiB**, plancher de
**8 GiB** pendant les commandes. L'extension est lancée sous `caffeinate`, après
archivage restauré/vérifié de **41 racines historiques supplémentaires**, avec
**17,49 Go libres** au contrôle préalable. Elle termine à **42 chaînes / 42
oracles**, avec les 36 mesures initiales et toutes les preuves préservées. Les
six paires froides réunies donnent **398,610 → 292,306 s (−26,67 %) de chaîne**,
**247,900 → 137,115 s (−44,69 %) de Go** et **117,733 → 124,544 s (+5,79 %)
de backend**. Le candidat 6 est **refusé au froid**, à **0,925 s** du budget ;
le stockage passe. Les paires supplémentaires backend sont à **+2,71 / +4,69 /
+8,71 %**. Un probe privé supprime le deuxième parcours AST du rendu : seuil
préalable de **5 % sur trois paires `PackageFiles`**. Il gagne **6,12 %
(11,463 → 10,761 s)**, avec gain dans chaque paire, suites du résolveur et six
plans/Go exacts. Ce gain local est retenu pour validation dans un candidat
complet reconstruit. Le rendu est reporté à l'identique dans le dépôt, puis la
séquence du **candidat 7** est lancée sous `caffeinate` : archivage de neuf racines
historiques terminées, race detectors Go 1.27.0/1.22.12, reconstruction quatre
hôtes et contrôles CLI/transitions. Son contrôle suivant est préparé sur
**trois paires froides complètes, frontend compris**, avec les fonctions du
harnais figé et mêmes seuils. La reconstruction, les deux race detectors et les
**89 commandes CLI / 56 comparaisons** terminent avec succès : mêmes TAST,
objets métier/préparés et Go exécutés sur les quatre hôtes. L'audit confirme
le rendu en une passe et son WASM reconstruit comme seuls changements depuis 6.
Le cache privé de construction est archivé/restauré/vérifié ; **12,50 Go libres**
restent avant le contrôle froid. La séquence d'entretien complémentaire puis
des paires froides est lancée sous `caffeinate`, avec contrôle préalable des
inventaires publiés et réserve de 16 GiB. Le contrôle termine avec **six chaînes /
six oracles passants** et tous ses critères temps/stockage validés :
**115,177 → 117,740 s (+2,23 %) de backend**, **239,500 → 128,850 s (−46,20 %)
de Go**, **385,109 → 276,664 s (−28,16 %) de chaîne**. Les trois paires backend
sont à **+4,10 / +3,45 / +1,95 %**, toutes sous +5 %, avec **3,196 s** de marge
médiane. Le GOCACHE total reste à **+9,9914 %** ; les six archives/volumes et
TAST/Go sont vérifiés. Restent les **30 chaînes des cinq autres scénarios** du
candidat 7, les grands corpus quatre hôtes, le bootstrap courant et les
snapshots examinés/rejoués strictement ; G06 reste ouvert.
L'archivage supplémentaire autorisé des targets Cargo 5–7 et des sorties frontend
de **26 échantillons terminés** est achevé : **7 626 fichiers Cargo / 210 054
fichiers frontend** restaurés/vérifiés avant retrait, binaires, journaux, pointeurs
et inodes des verrous préservés. Le harnais des 30 chaînes reprend les 15 helpers
et quatre instructions d'environnement figés, avec archivage après chaque paire
et libération des caches chauds après leurs 24 mesures. Le contrôle du lanceur et
de l'archivage passe, avec **16 racines restaurées** ; réserve de départ **20 GiB**.
Douze anciennes racines de stockage supplémentaires sont archivées avec
restauration vérifiée (**131 904 fichiers**), portant l'espace libre observé à
**26,27 Go**. Les **30 chaînes** du candidat 7 sont lancées sous `caffeinate`
dans `full-2/` ; leur bilan sera réuni aux six froides de `cold-complete-7/`.
La campagne termine le **16 octobre** : **36 chaînes / 36 oracles** réunis,
**39 critères temps/stockage passants**, données froides initiales inchangées.
Les médianes de chaîne gagnent **62,88 % feuille / 48,36 % partagé / 54,50 % FFI**,
avec **26,88 / 34,37 / 42,62 % de surcoût backend** respectivement. Le froid
applicatif est à **−31,89 % de chaîne / +2,15 % de backend** ; le froid complet
reste à **−28,16 % / +2,23 %**. Le GOCACHE identique est à **+9,997897 %**,
extrêmement proche du plafond. Le backend reste à **~118 s à froid** : ce verdict
face à `publicsync` ne vaut pas retour aux repères initiaux de **~60 s à froid /
~30 s après modification**.
Les 36 lignes sont publiées dans `docs/build-go-cycle-g06.csv`. La validation
finale privée passe : **huit comparaisons / 28 commandes** des grands corpus
quatre hôtes, bootstrap émis fonctionnel et **neuf fixtures / 207 commandes**.
La revue des nouveaux snapshots retrouve **612 déclarations Main/FFI**, entrées
et imports d'initialisation externes exacts après normalisation des qualifications.
La densité **1/8192** et les artefacts mesurés sont installés après audit des
**295 fichiers source/outils**. Les replays sur l'installation et le bootstrap
courant passent ; deux fixtures reçoivent des identités d'ancres fixes pour
éprouver les frontières à cette densité. Les erreurs de sélection de Spago,
profondeur de workspace et PATH Go du harnais final sont conservées et corrigées
dans des reprises distinctes. Les commandes réussies antérieures sont préservées.
Les dix contrôles du runner passent sur macOS et Linux. G06 est livré, puis pause
avant P.
Les objets de **55 racines terminées** sont
archivés avec restauration vérifiée, journaux et inodes des verrous conservés ;
**26,5 Go libres** sont observés au lancement historique du candidat 6. La densité
du dépôt, restée 1/128 pendant les essais, passe à **1/8192** à l'installation du 7.
G est clos à **7/7**, P reste à **0/7, en pause**.

- [x] **Confirmer le gain de bout en bout et livrer le layout retenu.**
  - Exécuter le protocole de paires sur les trois éditions et mesurer aussi
    froid/identique, frontend, gopurs, compilation, lien, ressources et suite de
    tests. Vérifier les seuils G avec le runner réel, à ressources constantes.
  - Comparer les sorties byte-exactes entre JS, Go séquentiel, Go parallèle et
    Rust, puis entre cache et build forcé, sur les mêmes TAST et le même layout.
    Comparer à la référence monolithique les résultats d'exécution, diagnostics,
    tags et contrats FFI ; examiner les changements de chemins/imports des
    snapshots avant leur mise à jour, puis lancer les replays stricts.
  - Rejouer robustesse du stockage, migrations et contrats du runner ; publier
    commandes, empreintes, graphe, raisons de recompilation et limites observées.
  - **Fin :** seuils atteints, campagne complète verte et parcours d'installation/
    migration/récupération documenté dans b8x et gopurs.

## Vérifications communes

- Présenter le code actuel, la transformation proposée et l'impact éventuel sur
  les bridges avant chaque changement ciblé.
- Préserver le travail concurrent et conserver les reproducteurs avant/après.
- Un cache ne doit changer ni le résultat ni les diagnostics pertinents d'un
  build. Expliquer toute différence avec la référence précédente.
- Comparer les mêmes TAST et les inventaires complets de sorties, y compris
  runtime, bridges et entrées ; distinguer les posttraitements normaux de Go.
- Examiner les snapshots avant leur mise à jour, puis les rejouer strictement.
- Exécuter les tests adaptés au lot, conserver les commandes, empreintes,
  mesures et limites ; vérifier `git diff --check`.
- Les commits et pushes ordinaires restent gérés par le bot de l'utilisateur.
