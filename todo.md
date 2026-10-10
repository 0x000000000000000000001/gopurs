# Gopurs — build incrémental et lancement Go

Plan v3 — mise à jour : 10 octobre 2026.

## Objectif et périmètre

Réduire le cycle **modifier un fichier → construire → exécuter les tests** en
réutilisant les résultats encore valides. La phase gopurs complète prend environ
**33 secondes d'après le relevé utilisateur** ; établir sa ventilation avant de
fixer un objectif chiffré.

La première livraison doit couvrir les builds sans changement et les
modifications localisées. Elle concerne gopurs, les adaptations nécessaires du
PBO et le runner Go de `../../b8x`. Le résultat incrémental doit être identique
au résultat d'un build complet.

Le plan v2 est clôturé ; sa
[validation finale](docs/testing.md#plan-v2--lot-08--installation-et-validation-finale-7-octobre-2026)
conserve les preuves historiques.

## Avancement : 1/7 lots validés

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

- [ ] **Distinguer les coûts et conserver la référence avant changement.**
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

- [ ] **Introduire un contrat de cache versionné et vérifiable.**
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

### 03 — Build sans changement et écritures minimales

- [ ] **Terminer rapidement lorsque les entrées sont inchangées.**
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

### 04 — Incrémentalité par module après modification

- [ ] **Réutiliser l'optimisation et l'émission des modules encore valides.**
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

### 05 — Campagne de correction incrémental/complet

- [ ] **Éprouver les transitions de build sur des reproducteurs persistants.**
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

### 06 — Mesurer le gain et livrer le parcours utilisateur

- [ ] **Valider le bénéfice sur b8x et documenter l'usage.**
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

## Suites conditionnées par les mesures

- **Préparation incrémentale :** si la préparation globale domine encore les
  builds modifiés, rendre incrémentaux le chargement/décodage, les métadonnées et
  la collecte transitive des spécialisations, avec la même preuve d'équivalence.
- **Granularité des packages Go :** si le package monolithique `purescript` domine
  encore la compilation, étudier un découpage compatible avec les dépendances
  générées, les spécialisations, les helpers partagés et l'absence de cycles
  d'import Go. L'incrémentalité gopurs seule ne résout pas ce coût.

Ces suites seront dimensionnées après le lot 06 ; elles ne sont pas comptées
dans les sept lots de la première livraison.

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
