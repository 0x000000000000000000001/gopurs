# Tests et portée des validations

Les commandes de cette page partent de la racine de gopurs, avec Go, Node,
Spago et le `purs` TAST sur `PATH`. Voir [l'installation](../README.md).
Le runner utilise `bin/gopurs`, donc le binaire natif existant par défaut ;
`GOPURS_JS=1` sélectionne le bundle JS. `-c` reconstruit le bundle JS une fois via
npm. Après une modification du compilateur, `npm run build:native -- --keep-workspace`
reconstruit les deux versions et conserve le workspace pour les tests natifs.

## Choisir un contrôle court

| Famille modifiée | Commande ciblée |
| --- | --- |
| Pilote et durée de vie des tâches | `node --test tools/emission.test.mjs`, après le build ; tests natifs ci-dessous |
| Statut de sortie et diagnostics CLI | `npm run test:cli`, après `npm run build:native` |
| Bootstrap natif et gestion des processus | `node --test tools/build-native.test.mjs tools/test-runner.test.mjs` |
| Annotations et représentations natives | `node --test tools/native-record-workers.test.mjs tools/boxed-record-arguments.test.mjs tools/native-sum-results.test.mjs`, après le build |
| Preuves source/TCO des arguments record projetés | `node --test tools/native-record-args.test.mjs tools/native-record-workers.test.mjs`, après le build ; `./bin/test NativeRecordWorkers NativeRecordReturns` |
| Layouts, métadonnées et instanciation | `node --test tools/representation-contract.test.mjs tools/mixed-constructor-tags.test.mjs tools/elided-constructor-payloads.test.mjs tools/record-tuple-conversions.test.mjs`, après le build |
| Boxing et émission transitive Rebox | `node --test tools/rebox-generation.test.mjs tools/rebox-metadata.test.mjs tools/struct-pointer-boxing.test.mjs tools/value-array-unboxing.test.mjs`, après le build |
| Littéraux composites et constructeurs | `node --test tools/composite-expressions.test.mjs tools/elided-constructor-payloads.test.mjs tools/nullary-constructor-bindings.test.mjs`, après le build ; `./bin/test ConstructorReuse NativeArrayReboxing` |
| Types et records | `./bin/test NativeRecordBoxing NativeRecordSizes -c` |
| Bridge FFI | `node --test tools/ffi-bridge.test.mjs tools/ffi-generics.test.mjs`, après le build ; `./bin/test FFIIntegerReturns -c` |
| Appels et fonctions | `./bin/test CurriedLambdas -c` |
| Conversions de tableaux | `./bin/test ArrayRoundtrip -c` |
| Récursion | `./bin/test TCO TCOMutRec -c` |
| Bindings, captures et signatures de workers | `node --test tools/binding-contracts.test.mjs tools/recursive-initialization.test.mjs tools/local-native-returns.test.mjs tools/zero-arity-functions.test.mjs`, après le build |
| Fusion de thunks | `./bin/test ThunkFusion -c` |
| Admission des fusions et applications immédiates | `node --test tools/thunk-fusion.test.mjs tools/counted-functions.test.mjs tools/immediate-applications.test.mjs`, après le build |
| Intrinsics, indexation et traversées Either | `node --test tools/array-intrinsics.test.mjs tools/array-safe-index.test.mjs tools/array-unsafe-index.test.mjs tools/array-traverse-either.test.mjs tools/object-traverse-either.test.mjs tools/value-array-unboxing.test.mjs`, après le build |
| Propriété des arbres et workers consommants | `node --test tools/owned-trees.test.mjs`, après le build ; `./bin/test OwnedTrees RBTree ConstructorReuse` |
| Mise en cache des dictionnaires fermés | `node --test tools/closed-dictionaries.test.mjs`, après le build ; `./bin/test StaticDictionary JsonRecordPlan` |
| Emprunt d'objets et composition des décodeurs | `node --test tools/borrowed-objects.test.mjs tools/closed-dictionaries.test.mjs tools/decoder-schemas.test.mjs`, après le build ; `./bin/test JsonRecordPlan` |
| Contrat du parser Go | `go test ./...` depuis `tools/ffi-gen` |
| Parser WASM et erreurs FFI | `npm run test:ffi`, après `npm run build` |
| Sélection, isolation et erreurs du runner | `npm run test:runner` |

Cette table indique quel contrôle choisir, pas que chaque fixture possède un
snapshot validé avec le dernier générateur. Les limites connues figurent plus
bas. Le parser ne doit être reconstruit avec `npm run build:ffi` que si ses
sources changent ; le WASM et son runtime JavaScript doivent alors rester
appariés, comme décrit dans le README.

Pour le pilote, comparer également le Go généré sur les **mêmes entrées TAST**
avec l'ancien compilateur, le nouveau natif séquentiel, le natif parallèle et
le bundle JS. Comparer les fichiers octet par octet, y compris runtime, bridges
FFI et entrées exécutables. Les tests de
[l'émetteur natif](../tools/emission-native_test.go) se copient dans le répertoire
`output/purescript` du workspace conservé ; depuis son répertoire `output` :

```bash
go test -race -run '^TestPipelineNative' -count=1 -timeout 30s ./purescript
```

## Bootstrap natif et processus des runners

Les tests `build-native.test.mjs` remplacent les commandes de compilation par
des outils simulés. Ils vérifient la sélection du fork typé le plus récent ou
de `GOPURS_PURS`, son isolation du toolchain npm, le rejet d'un TAST incomplet,
la conservation du binaire précédent sur échec, et la politique de nettoyage
du workspace. Les interruptions SIGINT/SIGTERM doivent arrêter les descendants,
conserver les logs et rendre les statuts 130/143.

`test-runner.test.mjs` exerce aussi la gestion partagée des processus à travers
les campagnes de fixtures et de bibliothèques. Pour valider un changement du
bootstrap avec les vrais outils, exécuter `npm run build:native -- --keep-workspace`.
Un refactoring de cet outillage à sources et toolchains identiques peut ensuite
être contrôlé par comparaison des artefacts `bin/gopurs.js` et `bin/gopurs-native`.

La passe du 30 septembre 2026 a validé les **15 tests d'outillage**, puis un
bootstrap réel : les deux artefacts reconstruits avaient les mêmes empreintes
SHA-256 qu'avant le refactoring. `FFIIntegerReturns` et `ObjectUpdate2` ont aussi
passé leurs snapshots stricts, leur compilation Go et leur exécution avec le
runner partagé.

## Contrat de sortie CLI

`npm run test:cli` compile une fixture réelle dans un workspace temporaire,
puis appelle `bin/gopurs` en JS, natif séquentiel et natif parallèle. Chaque mode
vérifie le succès (statut 0 et Go identique) et cinq erreurs réelles (statut 1) :
chargement du TAST, écriture du runtime, écriture d'un module avec workers PBO,
FFI invalide et écriture de l'entrée exécutable.

Les erreurs doivent conserver leur message d'origine sur stderr, avec un seul
préfixe `[gopurs] error:`, et terminer avant le délai du test. Le résultat est
traité par `Main` après la sortie des brackets et de la supervision du pilote.
Les tests de l'émetteur ci-dessus vérifient séparément que les workers ont bien
terminé leur nettoyage à cette frontière.

## Contrat des représentations

`representation-contract.test.mjs` croise les métadonnées d'origine et enrichies,
les déclarations Go et les champs de dictionnaires instanciés. Il fixe aussi les
différences d'arité entre valeurs, champs génériques, définitions et constructions
saturées, l'ordre des `TypeApp`, la résolution `$Dict`, les priorités d'élimination
et l'identité des alternatives `nil` importées. Il utilise les entrées publiques
déjà présentes avant le refactoring.

Les tests `mixed-constructor-tags`, `elided-constructor-payloads` et
`record-tuple-conversions` complètent ce contrat par compilation et exécution
du Go produit : tags distincts, payloads polymorphes, records et champs de classes.

## Contrat des arguments record projetés

`native-record-args.test.mjs` exerce les trois entrées publiques de l'analyse.
Les contrats fixent les champs visibles et leurs collisions Go, la priorité des
annotations, les quantificateurs de ligne et les paramètres monomorphes associés.
Ils distinguent les portées source des identités TCO, y compris les initialisations,
les patterns imbriqués, les gardes et les captures différées.

La vérification finale rejette une preuve source devenue invalide et choisit
chaque argument indépendamment. Les résultats natifs `Maybe`/`Either`/`Tuple`
sont contrôlés aux deux étapes. `native-record-workers` compile et exécute les
workers et leurs appelants Go : layouts plus larges, wrappers dynamiques/partiels,
préservation des champs supplémentaires et replis. `NativeRecordReturns` complète
ces contrats sur le chemin TAST/PBO avec un payload record transporté par `Either`.

## Contrat des dictionnaires fermés

`closed-dictionaries.test.mjs` distingue la preuve de réutilisation sans
annotation de celle du déplacement d'un site annoté. Les tests vérifient la
priorité du premier binding admissible, les qualifications et formes exactes
d'applications, la substitution entre arguments et le refus des types ambigus,
dynamiques ou de rang supérieur.

Les contrats de réécriture couvrent les racines protégées par leurs annotations,
les captures lexicales, les initialisations de `let`, les barrières récursives,
les effets et leurs enfants indépendants. Ils fixent aussi l'ordre des nouveaux
bindings, la réservation des noms et le maintien de sites de cache distincts.
Un programme Go généré, exécuté avec `go test -race`, vérifie l'évaluation
paresseuse, le partage du binding d'origine, une construction par site déplacé
et la réévaluation des constructions capturant les arguments de chaque appel.

## Contrat de l'emprunt d'objets

`borrowed-objects.test.mjs` couvre les alias de dictionnaires, les cycles, les
qualifications et arités exactes, les annotations préservées et la descente
distincte dans les producteurs admis ou rejetés. Les lecteurs de champs sont
contrôlés par nom, suffixe numérique et position de l'argument emprunté. Les
témoins de portée vérifient les niveaux locaux, leur masquage et les neuf formes
de closure, récursion ou effet, avec des contrôles positifs indépendants.

La suite exécute aussi les deux tests natifs du helper réel dans un workspace
isolé, avec `go test -race` : emprunt identité, copie du conteneur sur le chemin
possédé et replis exacts sur entrée invalide ou décodeur inconnu. Les suites
`closed-dictionaries` et `decoder-schemas` contrôlent les passes voisines ; la
fixture `JsonRecordPlan` vérifie l'ordre des callbacks, les erreurs et les
résultats persistants sur le chemin TAST/PBO.

## Contrat des workers consommants

`owned-trees.test.mjs` vérifie les chemins disjoints, les continuations encore
observables, les gardes, les captures avant mutation et les collisions des noms
worker/helper, y compris avec les FFI. Ses témoins de point fixe contrôlent le
rejet transitif des appelants et d'une famille mutuellement récursive dépendant
d'un worker rejeté, avec une famille indépendante comme contrôle positif.

Les programmes Go générés vérifient le réemploi des cellules sans alias ni cycle,
les cellules connues et nullables, le donneur des appels terminaux et le sous-arbre
conservé par une famille mutuellement récursive admise. `OwnedTrees` complète ces
contrats sur le chemin TAST/PBO : rotations, invariants de l'arbre, anciennes
versions encore observables et enfants partagés.

## Contrat des intrinsics et traversées

`array-intrinsics.test.mjs` vérifie les noms, qualifications et arités admis selon
la convention, ainsi que les limites de la normalisation du buffer entier frais.
Son programme Go compare les chemins curryfié, boxé, slice native et worker natif
de map/filter/fold : tableaux vides, ordre des callbacks, fold gauche à accumulateur
boxé et absence d'alias entre entrées et résultats réutilisés.

Les tests d'indexation contrôlent l'ordre des arguments, les bornes, les types des
éléments et la conversion d'un seul élément. `array-traverse-either` et
`object-traverse-either` rejettent les dictionnaires et formes inconnus avant
toute traduction ; leurs programmes Go vérifient respectivement **11 055** et
**270 comparaisons** avec les traversées strictes de référence. Ils contrôlent
les captures au bon stade, tous les callbacks après le premier `Left`, l'ordre
des clés Go et le stockage neuf des closures réutilisées.

La fixture `NativeTraverseCallback` complète ces contrats avec les captures
lexicales des callbacks natifs et le repli lorsqu'un calcul sépare deux lambdas.
`ArrayTraverseEither`, `ObjectTraverseEither` et `ArrayRoundtrip` exercent les
formes produites par le chemin TAST/PBO, avec snapshots stricts et exécution Go.

## Contrat des fusions et applications immédiates

Les 21 tests de `thunk-fusion.test.mjs` contrôlent la conservation du producteur,
l'omission des workers inutilisés, le forçage unique, les opérations entières
admises, les scopes récursifs et les collisions de noms source/Go/FFI.
`counted-functions.test.mjs` vérifie le motif récursif complet, le chemin négatif,
les annotations et les collisions ; ses programmes Go contrôlent les applications
partielles, les valeurs réutilisées, l'ordre et les échecs des callbacks.

`immediate-applications.test.mjs` vérifie l'admission de tous les résultats d'une
branche, le budget de duplication cumulé, la conservation des évaluations et les
scopes récursifs. Son programme Go vérifie les captures après transplantation,
les effets, les échecs et la réutilisation. Les **63 tests** de ces trois suites
ont réussi avant et après le refactoring du lot 10.

Les fixtures `ThunkFusion`, `ThunkFusionNewtype` et `CountedFunctions` complètent
ces contrats sur le chemin TAST/PBO avec snapshots stricts et exécution Go.

## Contrat des bindings et des fonctions

`binding-contracts.test.mjs` vérifie la publication préalable des signatures
récursives locales, leur affinement dans l'ordre source et la déclaration de tout
le groupe avant ses affectations. Un programme Go généré vérifie les permutations
d'arguments lors des sauts TCO et les captures de closures propres à chaque
itération. Ces deux tests ont aussi réussi avec le compilateur JS antérieur à
l'extraction de `LocalWorkers` et `ModuleWorkers`.

`recursive-initialization.test.mjs` contrôle les lectures anticipées, les cellules
publiées après initialisation et les références différées. `local-native-returns`
et `zero-arity-functions` exercent les frontières entre lambdas, les résultats
natifs et l'évaluation différée des fonctions sans argument. Les fixtures TCO et
de portée complètent ces contrats par snapshots stricts et exécution Go.

## Contrat des expressions composites

`composite-expressions.test.mjs` observe les contextes enfants, les statements,
les compteurs de noms et l'état Rebox aux frontières des émetteurs. Il distingue
la traduction complète d'un tableau avant adaptation de la conversion immédiate
de chaque champ de record ou de constructeur. Il contrôle aussi les tableaux
vides/hétérogènes, les paramètres de callbacks curryfiés et non curryfiés, leur
portée locale et la réservation du nom lors d'une réutilisation de constructeur.

Les fixtures `ConstructorReuse`, `NativeArrayReboxing`, `NativeRecordBoxing` et
les tests de constructeurs/sommes natifs vérifient le raccordement au dispatcher,
les snapshots et le comportement du Go généré.

## Contrat Rebox

`rebox-generation.test.mjs` part des appels publics de conversion, puis émet
leurs helpers. Il vérifie la fermeture transitive dans les deux sens, la
convergence des champs récursifs, la déduplication et l'ordre indépendant de
l'insertion des demandes. Le Go généré est compilé et exécuté pour contrôler
les payloads imbriqués, `nil`, l'identité des pointeurs à paramètres fantômes,
les tableaux natifs et l'évaluation unique des opérandes. Un cas sans métadonnées
fixe également le diagnostic et l'omission historiques du helper.

`rebox-metadata.test.mjs` couvre les collisions et priorités de l'index de champs.
`struct-pointer-boxing.test.mjs`, `value-array-unboxing.test.mjs` et les tests de
records/sommes natifs complètent les passages par `Value` et leurs imports.

## Contrat du bridge FFI

`ffi-bridge.test.mjs` exerce les trois entrées publiques de `FfiBridge` avec des
déclarations analysées par le vrai parser. Il vérifie que les noms et signatures
publiés correspondent aux workers émis, et distingue les workers générés des
signatures admissibles pour un appel direct. Le Go produit est compilé et exécuté
pour contrôler les callbacks, les records, les retours sans valeur, les effets
différés, les variables étrangères et les replis de noms ou de déclaration.

`ffi-generics.test.mjs` complète ce contrat avec les paramètres génériques,
callbacks de tableaux, fonctions renvoyées et valeurs Applicative. Les erreurs
de source et de transport restent couvertes par `ffi-errors.test.mjs`.

## Sélection et snapshots

```bash
./bin/test --list
./bin/test TCOMutRec ThunkFusion --list
./bin/test --skip-before ThunkFusion --list
./bin/test FFIIntegerReturns --keep-workspace
```

Sans cible, `bin/test` sélectionne toutes les fixtures non exclues de
`tests/passing`. Les noms explicites gardent leur ordre ; la reprise inclut sa
cible. `--skip-before=NAME` et `skip_before=NAME` restent compatibles. Une cible
inconnue, une option inconnue ou une sélection entièrement exclue échoue avant
le build. `--list` ne compile rien et ne crée pas de workspace.

La vérification de snapshots est le mode par défaut. Un fichier absent ou
différent échoue : examiner le diff et la phase responsable avant de le
remplacer. Les snapshots sont `tests/passing-snapshots/<Fixture>.go`, avec
`<Fixture>_ffi.go` en plus si la source déclare `-- @snapshot-ffi`.
Une mise à jour intentionnelle s'effectue ainsi :

```bash
./bin/test FFIIntegerReturns --update-snapshots
# Équivalent historique : UPDATE_SNAPSHOTS=1 ./bin/test FFIIntegerReturns
```

Les snapshots ne sont écrits qu'après compilation et exécution Go réussies de
la fixture concernée. La campagne s'arrête au premier échec ; les mises à jour
des fixtures déjà réussies restent écrites. Le contrôle d'exécution conserve
le contrat historique : statut zéro et absence de `Fail` dans la sortie.

## Isolation, logs et caches

Chaque fixture reçoit son propre répertoire temporaire : sources, configuration
Spago, lockfile, `.spago` et `output`. `tests/runner` n'est ni lu ni modifié.
Les `.go`, `.js` et répertoires compagnons d'une fixture sont copiés avec sa
source. `-- @dependencies: assert prelude effect console` peut limiter ses
dépendances ; sans directive, la liste de `bin/pkg` est utilisée. Les checkouts
core restent requis et les paquets utilisent le cache global de Spago.

Les logs distinguent compilation PureScript, génération Go, formatage,
snapshots, dépendances Go, compilation Go et exécution. Les workspaces réussis
sont supprimés, sauf avec `--keep-workspace`. Échecs et interruptions conservent
le workspace et affichent son chemin. SIGINT/SIGTERM sont transmis à la
commande active et ses sous-processus. Une compilation PureScript échouée
n'est pas relancée automatiquement.

`-c` ne réinitialise pas les caches globaux : il reconstruit gopurs. Les sorties
de fixture sont neuves avec ou sans cette option. À l'intérieur du backend,
chaque lancement régénère le Go et relit la FFI ; changer seulement une FFI Go
ne demande pas de recompiler le PureScript si les entrées TAST sont inchangées.

## Exclusions et modules frères

Les neuf exclusions historiques sont conservées dans
[tools/test-selection.mjs](../tools/test-selection.mjs) :

| Fixtures | Motif enregistré dans le runner |
| --- | --- |
| `DerivingClause`, `DerivingContravariant`, `DerivingFunctorFromBi`, `DerivingFunctorFromPro`, `DerivingProfunctor` | Fonctionnalités de compilateur plus récentes que celles prises en charge par ces fixtures |
| `NumberLiterals` | Différences de sérialisation IEEE-754 |
| `StringEdgeCases`, `StringEscapes` | Surrogates isolés et chaînes Go UTF-8 |
| `2136` | Débordement aux bornes 32 bits, avec les entiers natifs 64 bits de gopurs |

Ces motifs décrivent les exclusions existantes, pas une nouvelle vérification
de chacune. `bin/modtest` sélectionne les checkouts frères `gopurs-*` possédant
un `bin/test` exécutable :

```bash
./bin/modtest --all --list
./bin/modtest --skip-before strings --list
./bin/modtest prelude strings
```

La sélection complète est le défaut. Les noms avec ou sans `gopurs-` sont
acceptés ; `-c` reconstruit le backend depuis ce checkout. Chaque script frère
gère encore ses propres sorties et nettoyages ; l'isolation des fixtures de
`bin/test` ne s'étend pas automatiquement à ces scripts.

## Analyses spécialisées — lot 12, passe NativeRecordArgs, 2 octobre 2026

`NativeRecordArgs` conserve l'API et la preuve TCO qui sélectionne chaque argument
du worker. `Source` possède l'admission avant monomorphisation et les portées
CoreFn ; `Projection` possède les champs visibles, les labels Go et la politique
des résultats natifs. La sélection finale réutilise les champs déjà reconnus.
Le contrat documentaire inclut désormais les résultats `Maybe`/`Either`/`Tuple`
déjà admis par le code.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **493 modules et
  286 672 types**. Le build JS est sans avertissement ; les cinq avertissements
  du build TAST dans `DecoderSchemas` et PBO sont identiques à la passe précédente ;
- **36 tests ciblés avant et après** : `native-record-args`, `native-record-workers`,
  `boxed-record-arguments`, `native-sum-results`, `record-tuple-conversions` et
  `foreign-forwarders`. Les douze nouveaux contrats couvrent les priorités
  d'annotations, les signatures source, les portées, les labels et la vérification
  indépendante du TCO. Les programmes Go contrôlent les appels natifs et boxés,
  les résultats natifs, les captures et la conservation des records complets ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `NativeRecordWorkers`, `NativeRecordReturns`, `NativeRecordBoxing`,
  `NativeRecordSizes`, `CompactRecordConsumers`, `DuplicateProperties`,
  `FunctionScope`, `StaticDictionary` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **101 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

La passe `NativeRecordArgs` est validée ; seule `DecoderSchemas` reste à revoir
dans le lot 12. Le score reste à **75/100**, avec onze lots clôturés. Le contrôle
b8x porte sur la génération, les tests et fixtures sur la compilation et
l'exécution Go ; `git diff --check` est propre.

## Analyses spécialisées — lot 12, passe ClosedDictionaries, 2 octobre 2026

`ClosedDictionaries` conserve la réécriture et la publication des bindings.
`Admission` possède les preuves distinctes de réutilisation et de déplacement ;
`Scope` porte la fermeture lexicale et les exclusions d'effets/récursion. L'index
des bindings partagés est immuable, séparé de l'état des nouveaux bindings.
La descente commune protège les racines annotées ; le parcours des références
libres utilise le pli structurel du backend, avec les portées de binding explicites.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **491 modules et
  286 441 types**. Le build JS est sans avertissement ; les cinq avertissements
  du build TAST dans `DecoderSchemas` et PBO sont identiques à la passe précédente ;
- **49 tests ciblés avant et après** : `closed-dictionaries`, `borrowed-objects`
  et `decoder-schemas`. Les neuf nouveaux tests fixent les politiques d'admission,
  l'ordre source, les annotations et les portées. Le programme Go généré vérifie
  avec `go test -race` la paresse des getters, une construction par site et les
  captures réévaluées à chaque appel ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `StaticDictionary`, `JsonRecordPlan`, `EnumDictionaryField`,
  `TypeClassMemberOrderChange`, `EmptyDicts`, `3558-UpToDateDictsForHigherOrderFns`,
  `FunctionScope`, `MutRec` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **99 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

La passe `ClosedDictionaries` a été validée avec deux analyses encore à revoir
dans le lot 12. Le score est resté à **75/100**, avec onze lots clôturés.
Le contrôle b8x porte sur la génération, les tests et fixtures sur la
compilation et l'exécution Go ; `git diff --check` est propre.

## Analyses spécialisées — lot 12, passe BorrowedObjects, 2 octobre 2026

`BorrowedObjects` reste dans un seul module : `admitBorrowing` sépare l'admission
du producteur et la preuve des usages de la construction de l'appel helper.
`BorrowContext` porte les définitions admissibles et la signature du helper ;
`BorrowScope` distingue l'enveloppe `Either` de ses alias d'objet. Les annotations,
les barrières de portée et l'ordre de descente de la réécriture sont explicites.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **489 modules et
  286 220 types**. Le build JS est sans avertissement ; les cinq avertissements
  du build TAST dans `DecoderSchemas` et PBO sont identiques à la passe précédente ;
- **40 tests ciblés avant et après** : `borrowed-objects`, `closed-dictionaries`
  et `decoder-schemas`. Les 19 nouveaux tests d'emprunt couvrent les alias,
  annotations, qualifications, arités, scopes et replis ; deux tests Go du helper
  réel sont exécutés avec le détecteur de courses ;
- **4 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `JsonRecordPlan`, `ObjectTraverseEither`, `ObjectUpdate2`, `FunctionScope` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **97 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

La passe `BorrowedObjects` a été validée avec trois analyses encore à revoir dans
le lot 12. Le score est resté à **75/100**, avec onze lots clôturés.
Le contrôle b8x porte sur la génération, les tests et
fixtures sur la compilation et l'exécution Go ; `git diff --check` est propre.

## Analyses spécialisées — lot 12, passe Ownership, 2 octobre 2026

`Ownership` conserve l'orchestration du point fixe et la sélection des appels
frais. `Candidates` possède l'admission des layouts/signatures et la réservation
des paires de noms ; `Analysis` et `Types` portent le langage de preuve et les
chemins canoniques ; `Workers` possède les captures, le plan de cellules, le
retrait des alias et les déclarations Go. La construction redondante de noms
provisoires et le contexte inutilisé du plan ont été retirés.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **489 modules et
  286 160 types**. Le build JS est sans avertissement ; les cinq avertissements
  du build TAST dans `DecoderSchemas` et PBO sont les mêmes qu'au lot 11 ;
- **29 tests de possession avant et après**, dont quatre nouveaux contrats :
  rejet transitif des appelants, cycle dépendant d'un worker rejeté, collision
  sur un helper consommant FFI et exécution d'une famille mutuellement récursive
  admise. Les programmes Go vérifient aussi les captures, les stocks de cellules,
  les donneurs et l'absence de cycles ou d'alias introduits par le réemploi ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `OwnedTrees`, `RBTree`, `ConstructorReuse`, `Recursion`, `TCO`, `TCOMutRec`,
  `ShadowedTCOLet`, `OneConstructor` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **97 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

La passe `Ownership` a été validée avec quatre analyses encore à revoir dans le
lot 12. Le score est resté à **75/100**, avec onze lots clôturés. Le contrôle b8x
porte sur la génération, les tests et fixtures sur la compilation et l'exécution
Go ; `git diff --check` est propre.

## Validation des intrinsics et traversées — lot 11, 2 octobre 2026

`ArrayTraverse` et `ObjectTraverse` séparent reconnaissance, capture des arguments
et émission de la boucle. `CallArguments.capture` partage les liaisons ordonnées
en laissant au consommateur le choix de représentation. `ArrayIntrinsics.Index`
isole l'indexation et la conversion de l'élément sélectionné ; `Source` partage
l'accès aux buffers. Map/filter/fold distinguent la boucle, l'adaptation des
callbacks connus et les arguments restants.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **485 modules et
  285 099 types**. Le build JS est sans avertissement ; le build TAST neuf signale
  cinq avertissements dans `DecoderSchemas` et PBO, hors des modules modifiés ;
- **12 tests ciblés avant et après** : `array-intrinsics`, `array-safe-index`,
  `array-unsafe-index`, `array-traverse-either`, `object-traverse-either`,
  `value-array-unboxing`. Les trois nouveaux contrats d'intrinsics ont aussi été
  exécutés avec le générateur précédent. Les programmes Go couvrent notamment
  les **11 055 comparaisons de tableaux** et **270 comparaisons d'objets** ;
- **12 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `ArrayRoundtrip`, `ArrayTraverseEither`, `ObjectTraverseEither`,
  `NativeTraverseCallback`, `NativeArrayReboxing`, `ArrayType`, `DerivingFoldable`,
  `DerivingFunctor`, `CurriedLambdas`, `FunctionScope`, `JsonRecordPlan`,
  `FFIIntegerReturns` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **93 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

Le contrôle b8x porte sur la génération ; les tests et fixtures ci-dessus valident
la compilation et l'exécution Go. Le score du plan est passé à **75/100**, avec
onze lots clôturés ; `git diff --check` est propre.

## Validation des fusions — lot 10, 1er octobre 2026

`ThunkFusion` nomme la reconnaissance des paramètres, des corps suspendus et des
arguments récursifs et initiaux. `FunctionFusion` distingue les preuves du cas zéro,
de la décrémentation et de la composition du callback. `WorkerNames` partage la
réservation des identifiants source et Go. `ImmediateApplications` produit un plan
d'admission consommé par la réécriture ; `Scope` possède le comptage des occurrences,
la substitution et le renommage des locaux.

Vérifications effectuées :

- reconstruction JS et native sans avertissement ; bootstrap avec TAST vérifié
  sur **483 modules et 284 959 types** ;
- **63 tests ciblés** : `thunk-fusion`, `counted-functions`,
  `immediate-applications`. Les trois suites, dont les 21 nouveaux tests de thunks
  et les programmes Go des callbacks/captures, ont réussi avant et après le
  refactoring ;
- **12 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `ThunkFusion`, `ThunkFusionNewtype`, `CountedFunctions`, `CurriedLambdas`,
  `FunctionScope`, `PartialFunction`, `PartialTCO`, `TCO`, `TCOMutRec`,
  `ShadowedTCOLet`, `ArrayTraverseEither`, `ObjectTraverseEither` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **69 fichiers
  sources et artefacts compilateur** est stable avant et après ces comparaisons.

Le contrôle b8x porte sur la génération ; les tests et fixtures ci-dessus valident
la compilation et l'exécution Go. Le score du plan est passé à **70/100**, avec
dix lots clôturés ; `git diff --check` est propre.

## Validation des bindings et fonctions — lot 09, 1er octobre 2026

`ModuleBindings` publie les signatures et confie l'émission à `ModuleWorkers`.
`BindingExprs` distingue allocation des noms, publication des signatures locales
et initialisation des cellules ; `LocalWorkers` partage l'émission native et
curryfiée. Les environnements capturés, les paramètres d'itération et les types
résiduels des fonctions ont des helpers communs. Les champs inutilisés de
`LoopTarget` ont été retirés ; sélection et émission des sauts restent nommées
dans `CallExprs`.

Vérifications effectuées :

- reconstruction JS et native sans avertissement ; bootstrap avec TAST vérifié
  sur **481 modules et 284 456 types** ;
- **32 tests ciblés** : `binding-contracts`, `recursive-initialization`,
  `local-native-returns`, `zero-arity-functions`, `imported-workers`,
  `native-record-workers`, `native-sum-results`, `boxed-record-arguments`,
  `composite-expressions`, `representation-contract`, `elided-constructor-payloads`,
  `nullary-constructor-bindings`. Les deux nouveaux tests, dont la compilation
  et l'exécution Go des captures TCO, ont aussi réussi avant le refactoring ;
- **15 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `TCO`, `TCOCase`, `TCOFloated`, `TCOMutRec`, `ShadowedTCO`, `ShadowedTCOLet`,
  `PartialTCO`, `MutRec`, `MutRec2`, `MutRec3`, `CurriedLambdas`, `FunctionScope`,
  `BigFunction`, `NativeRecordWorkers`, `NativeRecordReturns` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **67 fichiers
  sources et artefacts compilateur** est stable avant et après ces comparaisons.

Le contrôle b8x porte sur la génération ; les tests et fixtures ci-dessus valident
la compilation et l'exécution Go. Le score du plan est passé à **65/100**, avec
neuf lots clôturés ; `git diff --check` est propre.

## Validation du dispatcher — lot 08, 1er octobre 2026

Les tableaux et records littéraux passent par `LiteralExprs` ; définitions et
constructions saturées passent par `ConstructorExprs`. L'ordre de traduction,
les conversions immédiates ou différées et la réutilisation des constructeurs
sont explicites dans ces émetteurs. Le typage des paramètres de callbacks utilise
le helper commun déjà employé par les records.

Vérifications effectuées :

- reconstruction des compilateurs JS et natif ; bootstrap avec TAST vérifié sur
  **479 modules et 284 264 types** ;
- **39 tests ciblés** : `composite-expressions`, `record-tuple-conversions`,
  `elided-constructor-payloads`, `nullary-constructor-bindings`,
  `native-sum-results`, `native-record-workers`, `boxed-record-arguments`,
  `local-native-returns`, `representation-contract`, `rebox-generation`,
  `native-constructor-tags`. Les six nouveaux tests contrôlent notamment l'état
  Rebox visible au champ suivant, la portée des paramètres et le compteur de noms ;
- **12 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `ConstructorReuse`, `NativeArrayReboxing`, `ArrayRoundtrip`, `NativeRecordBoxing`,
  `NativeRecordSizes`, `NativeRecordReturns`, `CompactRecordConsumers`,
  `TypeClassMemberOrderChange`, `EnumDictionaryField`, `PartiallyAppliedMaybe`,
  `DuplicateProperties`, `RBTree` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus.

Les builds, les 39 tests ciblés et les trois comparaisons b8x ont été reconfirmés
après le nettoyage final des imports et noms locaux. Le manifeste des **65 fichiers
sources et artefacts compilateur** est stable avant et après ces comparaisons.
Le contrôle b8x porte sur la génération ; les tests et fixtures ci-dessus valident
la compilation et l'exécution Go. Le score du plan est passé à **60/100**, avec
huit lots clôturés ; `git diff --check` est propre.

## Validation des conversions et de Rebox — lot 07, 1er octobre 2026

`GoConversions` porte le choix des conversions et la normalisation des passages
par `Value`. Les layouts de `Maybe`/`Either`/`Tuple sont isolés dans `NativeAdts` ;
`Rebox` possède les demandes, la traduction des champs et l'émission transitive.
Les helpers de records/tableaux sont nommés et la branche de tableau source natif
inaccessible après normalisation a été retirée.

Vérifications effectuées :

- reconstruction des compilateurs JS et natif ;
- **43 tests ciblés** : `rebox-generation`, `rebox-metadata`,
  `struct-pointer-boxing`, `value-array-unboxing`, `record-tuple-conversions`,
  `native-sum-results`, `native-record-workers`, `boxed-record-arguments`,
  `array-traverse-either`, `object-traverse-either`, `elided-constructor-payloads`.
  Les trois nouveaux tests Rebox, dont la compilation/exécution Go, ont aussi
  réussi avec les modules JS précédents avant reconstruction ;
- **12 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `NativeArrayReboxing`, `ArrayRoundtrip`, `NativeRecordBoxing`, `NativeRecordSizes`,
  `NativeRecordWorkers`, `NativeRecordReturns`, `CompactRecordConsumers`,
  `EnumDictionaryField`, `ArrayTraverseEither`, `ObjectTraverseEither`,
  `MaybeFfiRoundtrip`, `RBTree` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des sources et
  des deux binaires est identique avant et après ces trois comparaisons.

Ce contrôle b8x porte sur la génération ; les tests et fixtures ci-dessus
valident séparément la compilation et l'exécution Go.

## Validation des représentations — lot 06, 1er octobre 2026

La lecture des `TypeApp`, le nommage des constructeurs et les règles
d'instanciation sont regroupés dans `GoTypes`, `GoAst` et `ConstructorLayout`.
Les tables ADT/classes portent leurs types et conventions de clés. La revue des
classes a confirmé que leur ordre de champs était déjà partagé ; le commentaire
de collecte des enums a été corrigé pour refléter les déclarations d'origine.

Vérifications effectuées, puis builds et parité reconfirmés sur les artefacts
finaux :

- reconstruction des compilateurs JS et natif ;
- **43 tests ciblés** : `representation-contract`, `mixed-constructor-tags`,
  `elided-constructor-payloads`, `record-tuple-conversions`,
  `native-constructor-tags`, `boxed-constructor-tags`,
  `nullary-constructor-bindings`, `rebox-metadata` et `closed-dictionaries`.
  Les contrats d'instanciation ont aussi été contrôlés avec les modules JS
  précédents avant reconstruction ;
- **10 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `OneConstructor`, `PartiallyAppliedMaybe`, `PrimedTypeName`, `ConstructorReuse`,
  `RBTree`, `TypeClassMemberOrderChange`, `EnumDictionaryField`,
  `InheritMultipleSuperClasses`, `NativeRecordBoxing`, `NativeArrayReboxing` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Les bridges,
  le runtime, les entrées exécutables et `go.mod` sont inclus. Les TAST viennent
  de `b8x/run/bak/go/output`, le lien `b8x/output` ciblant alors le backend Rust.

La parité b8x contrôle la génération ; la compilation et l'exécution Go sont
couvertes par les tests ciblés et les fixtures ci-dessus.

## Validation du bridge FFI — lot 05, 30 septembre 2026

Le découpage de `FfiBridge` en façade, `Signatures`, `TypeSupport`, `Values`
et `Render` a été validé par :

- la reconstruction des compilateurs JS et natif ;
- les **19 tests** de `ffi-bridge.test.mjs`, `ffi-generics.test.mjs` et
  `ffi-errors.test.mjs`. Les deux nouveaux tests du contrat public ont aussi
  été exécutés avec les modules JS précédents avant leur reconstruction ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `FFIIntegerReturns`, `FFIConstraintWorkaround`, `CompactRecordConsumers`,
  `NativeRecordSizes`, `EffFn`, `FFIDefaultESExport`, `ESFFIFunctionFunction`,
  `ESFFIValueVar`. Les trois fixtures déclarant `@snapshot-ffi` vérifient aussi
  le texte de leur bridge ;
- b8x : référence régénérée avec le binaire précédent, **2 684 entrées TAST
  figées**, **2 989 fichiers Go identiques octet par octet** en natif parallèle,
  natif séquentiel et JS, sans ajout ni suppression. Les bridges, le runtime,
  les entrées exécutables et `go.mod` sont inclus dans la comparaison.

Ce contrôle b8x porte sur la génération. Les tests du bridge et les fixtures
ci-dessus couvrent séparément la compilation et l'exécution du Go produit.

## Validation du contexte de traduction au 30 septembre 2026

Le dispatcher `CodeGen` utilise directement `ExprContext`, dont `childContext`
définit le contexte des opérandes ordinaires. Le traitement des annotations et
des dictionnaires de classes est isolé dans `TypedExprs`. Cette passe a validé :

- la reconstruction des compilateurs JS et natif ;
- **28 tests ciblés** : records boxés et natifs, sommes natives, retours de
  fonctions locales, initialisation récursive, dictionnaires fermés, tests de
  constructeurs et traversées Array/Object ; le helper de métadonnées des tests
  fournit désormais aussi la table `ffiFunctions` requise par le générateur ;
- **10 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `TypeClassMemberOrderChange`, `EnumDictionaryField`, `StaticDictionary`,
  `NativeRecordWorkers`, `NativeRecordBoxing`, `ArrayRoundtrip`, `TCO`,
  `TCOMutRec`, `LetInInstance`, `PolykindBindingGroup2` ;
- b8x : une référence régénérée avec le binaire précédent sur une copie figée
  des **2 684 entrées TAST**, puis **2 989 fichiers Go identiques octet par octet**
  en natif parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime,
  bridges FFI et entrées exécutables sont inclus ; `go.mod` est également
  identique. Ce dernier contrôle porte sur la génération de code.

## Validation du pilote au 30 septembre 2026

Le découpage `Driver` / `Config` / `Prepare` / `Build` / `Output` et le scope
`Emission.withEmitter` ont été validés par :

- la reconstruction des compilateurs JS et natif ;
- les **19 tests JS** de l'émetteur et la suite native `TestPipelineNative`
  sous `go test -race`, notamment l'arrêt des workers sur erreur ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `FFIIntegerReturns`, `ArrayRoundtrip`, `ObjectUpdate2`, `NativeRecordBoxing`,
  `TCOMutRec`, `ThunkFusion`, `1664`, `test-int` ;
- b8x : **2 684 entrées TAST identiques**, **2 989 fichiers Go identiques octet
  par octet** entre l'ancien compilateur et les nouveaux natif parallèle,
  natif séquentiel et JS. Ce contrôle porte sur la génération de code ;
- une erreur d'écriture réelle, en remplaçant temporairement `Control_Bind.go`
  par un répertoire dans un workspace de fixture : arrêt en environ une seconde,
  sans deadlock. Le code de sortie zéro observé avec `launchAff_` a ensuite été
  corrigé par le traitement du résultat via `runAff_` et le contrat CLI ci-dessus ;
- les **18 scénarios CLI** de succès et d'échec sur les trois modes. La FFI Go
  invalide, qui déclenchait un `panic` natif, est désormais convertie en erreur
  Aff et rend le statut 1 avec le diagnostic contextualisé.

## Bilan du nettoyage au 14 septembre 2026

Les lots récents utilisent, à la demande de l'utilisateur, ce jalon transversal
depuis **altbak.pub** :

```bash
bin/go/run -c
```

Les lots 7.4–7.7, 8, 9 et 10 ont chacun été validés par ce parcours : rebuild du
backend et du bundle, compilation de l'application, génération puis compilation
Go et exécution des 14 cas du mode `pure`. Les comparaisons ont conservé les
mêmes 300 entrées CoreFn, les 387 fichiers Go identiques octet par octet et les
14 résultats fonctionnels. Les durées ne servent pas à conclure sur les
performances ; leurs baselines restent celles du README d'altbak.

Le lot runtime a aussi vérifié la propagation d'une modification de la source
Go dans le bundle et l'exécution d'un paquet déplacé sans la source runtime.
Le lot runner a passé sept contrats avec commandes de compilation simulées,
puis deux fixtures minimales avec les vrais outils dans des workspaces
séparés. Ces checks ne constituent pas une campagne complète du compilateur.

Restent ouverts : la référence des snapshots TCO/TCOMutRec et le contrôle
ciblé de fusion du bilan de la première vague, la validation étendue de
la [règle ArrayRoundtrip](array-roundtrip.md) prévue en 3.4 de ce chantier, et
une campagne complète `passing` / modules frères. Les contrôles ciblés plus
anciens et leurs écarts préexistants sont datés dans le todo. L'ancienne
affirmation « 100 % des tests officiels verts » ne décrit pas cette validation.

Le lot documentaire a reconstruit le backend sans artefacts compilés, installé
son archive dans un projet npm vide et exécuté l'exemple du README via les deux
backends : mêmes 82 fichiers Go et même sortie. Les dépendances installées,
checkouts frères et caches ont été réutilisés ; le téléchargement de tous les
prérequis n'a pas été rejoué. Ce build avait révélé **85 avertissements
préexistants (73 sources, 12 dépendances)**. Le lot 12 les a supprimés : deux
compilations de référence sans sorties préexistantes ont recompilé chacune les
438 modules, passant de 85 à **zéro avertissement et zéro erreur**. Un build
incrémental silencieux ne suffit pas à établir ce résultat. Les preuves
détaillées jusqu’au lot 12 restent consultables avec `git show baa1e071:todo.md`.
Le [todo actuel](../todo.md) décrit la deuxième vague de nettoyage.

Le contrôle `ArrayRoundtrip -c --keep-workspace` du 14 septembre confirme les
28 assertions existantes et le snapshot inchangé. Le résultat incorrect du
singleton pair consigné le 9 septembre ne se reproduit plus : le résultat est
`8`. Cette vérification n'a nécessité aucune modification du compilateur ;
elle ne désigne pas la cause ni la correction de l'ancien échec.

## Référence du lot 1 de maintenance — 14 septembre 2026

Cette référence accompagne la [carte des 51 dépôts](../todo.md#lot-1--carte-et-référence-du-14-septembre-2026).
Elle vérifie les points d'entrée et les comportements ci-dessous ; la revue
interne de tous les fichiers et la campagne des bibliothèques restent à faire.

| Contrôle exécuté | Résultat et portée |
| --- | --- |
| `node --test tools/*.test.mjs` | Les 11 fichiers passent : **43 tests**, aucun échec ni cas ignoré ; utilise les modules PureScript compilés de gopurs. |
| `go test -count=1 ./...` dans `tools/ffi-gen` | Réussite des tests natifs du parser, sans réutiliser un résultat de test en cache. |
| `./bin/test --list` | **373 fixtures** sélectionnées ; les neuf exclusions restent appliquées. Ce contrôle ne compile ni n'exécute ces fixtures. |
| `./bin/modtest --all --list` | **49 bibliothèques** sélectionnées ; QuickCheck est absent faute de `bin/test`. Les scripts frères n'ont pas été exécutés. |
| `bin/go/run -c` depuis altbak.pub | Build du backend et du bundle, compilation de l'application, génération et compilation Go, puis **14 cas pure réussis**. |
| Comparaison altbak avant/après | **300 TAST et 387 fichiers Go identiques octet par octet**, chemins inclus. Empreintes prises avant le nettoyage de `-c`. |
| `npm pack --dry-run --ignore-scripts --json` | Le manifeste npm contient le bundle, le lanceur, le runner FFI, le WASM et `wasm_exec.js`. L'installation autonome n'a pas été rejouée dans ce lot. |
| `spago ls deps --offline --json` sur une copie temporaire de la configuration Aff | Résolution réussie avec Spago 1.0.4, sans lockfile initial ; les chemins locaux ont été rendus absolus pour préserver leur destination. Pas de compilation d'Aff. |

Les 14 sorties fonctionnelles d'altbak, dans l'ordre de `src/App.purs`, sont :

```text
AST=7, Fibonacci=55, List=202950, TCO=100000, Records=20000,
Ackermann=125, Church=100000, Primes=21536, RBTree=22,
Polymorphism=10000000, State=1200, Lazy=1000000, Array=202950, RowToList=5
```

Les sources de référence sont gopurs `f1fa7ef6`, PBO Go `574c72e6`, fork
PureScript `40840b3e` et altbak `1d141427`. Le binaire TAST effectivement appelé
par altbak annonce `0.15.16 [development build; commit: a6a9864… DIRTY]` : ce
contrôle ne prouve pas qu'il a été reconstruit depuis le HEAD du fork.
Node est en **24.8.0** et Go en **1.27.0**. Le build npm de gopurs sélectionne
Spago **0.93.45** et `purs` **0.15.16** dans `node_modules/.bin`. Le script altbak
sélectionne Spago **1.0.4** et le binaire TAST via `run/bak/js/node_modules/.bin`.
Le `PATH` interactif seul désignait Spago **1.0.3** et `purs` **0.15.15** : ne
pas assimiler les commandes d'un shell, de npm et d'altbak.

L'environnement restreint empêchait initialement l'écriture dans les caches
Go et SQLite de Spago. Les contrôles Go ont utilisé un `GOCACHE` dans `/tmp` ;
Spago a ensuite été exécuté avec accès à son cache habituel. Ces premiers
arrêts sont des limites d'accès de l'environnement, pas des régressions du
compilateur. Les preuves et journaux sont temporaires dans
`/tmp/gopurs-lot1-20260914/`. Les résultats fonctionnels ci-dessus sont la
référence durable ; les performances restent à comparer aux baselines du
README d'altbak, sans conclusion tirée de ce seul run.

## Limites de configuration et de couverture relevées au lot 1

- **45 des 49 `bin/test` frères** nettoient aussi les `output`, `.spago` et
  `.cache` des autres `gopurs-*`. Les quatre nettoyages limités au paquet sont
  ceux de `functions`, `lazy`, `js-bigints` et `strings-extra`. Ce dernier
  lance aussi `go get github.com/iancoleman/strcase`. Revue au lot 4.
- `gopurs-assert/bin/test` vérifie la compilation avec `go build ./...`, sans
  suite exécutable. Les 48 autres runners frères ciblent `Test.Main`.
  `gopurs/spago.yaml` déclare ce module sans source de suite correspondante ;
  les fixtures de `tests/passing` ne sont pas une suite Spago `Test.Main`.
  `node-net` possède `test/Main.purs`, mais sa configuration ne déclare pas
  `package.test`. La configuration locale de QuickCheck n'en déclare pas non
  plus. Attribution aux lots 2, 3, 13 et 14, selon le propriétaire.
- **42 bibliothèques** ont un `spago.yaml` suivi qui pointe vers
  `spago.go.yaml`. Les fichiers Bower, Dhall, npm, CI et les compagnons JS
  décrivent encore d'autres parcours. Une commande `npm test` peut lancer
  Pulp/JavaScript, et n'est pas interchangeable avec `bin/test`.
- Les configurations actives citent six dossiers absents : `js-uri` et
  `simple-json` dans 41 paquets chacun, `test` dans 40, `safe-coerce` dans
  `effect` et `functions`, `math` et `starter` dans `unfoldable`. **Le contrôle
  Aff ci-dessus réussit avec ses trois entrées absentes** : une entrée inutilisée
  d'`extraPackages` n'est pas un blocage démontré. Le lot 2 doit vérifier la
  résolution utile à chaque paquet avant de modifier ces configurations.
- Les noms Spago de `js-promise`, `js-promise-aff`, `node-path` et
  `node-process` contiennent le préfixe `gopurs-`, alors que les overrides
  utilisent aussi les noms sans préfixe. Leurs interfaces et chemins sont
  recensés ; leur harmonisation éventuelle appartient au lot 2.
- Cinq modules de bibliothèque ont des imports étrangers avec un compagnon
  JS et aucun compagnon Go adjacent : `Foreign.Keys`,
  `Foreign.Object.ST.Unsafe`, `Foreign.Object.Unsafe`, `Node.Encoding` et
  `Node.Symbol`. La recherche FFI et les traitements intrinsèques doivent être
  suivis avant de conclure à un défaut ; leur couverture Go n'est pas établie
  par ce lot. Revue aux lots 7, 10 et 13.
- Sept lockfiles Spago ignorés sont présents dans `aff`, `argonaut-core`,
  `avar`, `js-date`, `now`, `nullable` et `strings-extra`. Trois lockfiles npm
  ignorés sont présents dans `node-fs`, `node-http` et `node-process`. Ils
  appartiennent à l'état local observé, sans garantie qu'un clone neuf les
  reconstruise à l'identique ; revue au lot 2.

Les écarts de snapshots et les campagnes non exécutées décrits plus haut
restent ouverts. Le lot 1 n'établit ni un build intégral de chaque bibliothèque,
ni une validation réseau/FS/Aff, ni un build sans caches de dépendances.

## Lot 2 — installation et configurations

Le [guide local](../README.md#develop-one-library-locally) décrit désormais le
parcours de chaque bibliothèque. Le lot 2 a examiné les **404 fichiers de
configuration** inventoriés : **91 modifiés, 313 conservés**, ainsi que
`bin/setup` et les imports/dépendances Go. Les fichiers Bower, Dhall, CI,
formatage, lint et règles d'exclusion gardent leurs rôles existants. Les
configurations des exemples et le template d'intégration de `spec` sont aussi
identifiés ; `SPEC_REPO_PATH` y est remplacé par le runner, et leur revue de
tests reste au lot 14.

**1 106 overrides inutilisés ont été retirés de 46 configurations Spago
principales**, y compris les références aux six noms de dossiers absents.
Les graphes complets avant/après ont été résolus avec Spago 1.0.4 dans deux
copies des 51 dépôts : mêmes paquets, versions, chemins locaux et dépendances
de tests. Les lockfiles Spago suivis ont été synchronisés ; les sept lockfiles
locaux ignorés n'ont pas servi à construire ces copies. La politique de
lockfiles et les 42 liens suivis sont préservés.

Quatre lockfiles avaient aussi des métadonnées locales périmées : `random`
dans `foreign-object` et `ordered-collections`, `js-promise` dans
`js-promise-aff`, et `free` dans `run`. Leur liste de dépendances reflète
maintenant les manifestes locaux, sans changer de chemin ni de version.
Le backend garde son Spago 0.93.45 et son graphe de paquets verrouillé ; ne pas
confondre ce graphe avec celui recalculé par Spago 1.0.4 pour la comparaison.

| Contrôle du lot 2 | Résultat |
| --- | --- |
| `spago ls deps --offline --transitive --json`, avant/après | **51/51 graphes identiques** ; aucune résolution en échec. |
| `bin/setup`, avec Git simulé dans `/tmp` | Sélections **25/50**, options invalides, prévalidation, arrêt au premier échec de clone, réexécution sans recloner, chemins avec espaces et préservation des fichiers existants vérifiés. |
| `./bin/setup --all` dans l'arborescence existante | **50 checkouts conservés**, aucun clone ni changement de leurs fichiers. |
| Build du backend dans une copie sans `output`, `.spago` ou `Runtime.js` généré | **445 modules compilés**, zéro avertissement, zéro erreur, bundle produit ; dépendances installées et caches de paquets réutilisés. |
| `gopurs-refs` dans une copie contenant uniquement les 25 checkouts core | **161 modules compilés**, zéro avertissement/erreur ; génération Go et tests réussis (`All tests passed!`). Les runners frères et leurs nettoyages n'ont pas été appelés. |
| Archive du backend → projet npm vide, cache npm vide, installation hors ligne avec `--ignore-scripts` | **Un seul paquet installé**, aucune dépendance npm d'exécution ; exemple TAST/Go réussi, **82 fichiers Go**, sortie `Hello from gopurs`, runtime identique à sa source canonique. |
| Archive reconstruite depuis les sources, appliquée aux mêmes TAST d'altbak | **387 fichiers Go identiques** à la référence, dans un répertoire de sortie neuf. |
| `node --test tools/*.test.mjs` | **43/43 tests réussis** après les changements. |
| `bin/go/run -c` depuis altbak | **300 TAST et 387 fichiers Go identiques octet par octet**, mêmes **14 résultats fonctionnels** que le lot 1. |
| Lockfiles npm suivis | Les cinq correspondent aux dépendances directes déclarées ; `npm ci --dry-run --ignore-scripts --offline` réussit pour le lockfile Promise/Aff corrigé. |

Le backend ES inutilisé a été retiré du manifeste npm de gopurs ; esbuild est
une dépendance de développement, et l'archive n'a plus de dépendance npm
d'exécution. Dans Promise/Aff, le lockfile enregistre maintenant le compilateur
forké déjà déclaré et retire 99 entrées devenues inutiles, notamment celles
d'ESLint. Les versions des entrées conservées n'ont pas changé. Le lockfile
Yoga JSON est aligné sur son manifeste. Sa commande d'éditeur utilise désormais
`spago build --json-errors`, à la place d'un `test.dhall` absent.

Le profil core installe 22 bibliothèques et trois supports ; il n'ajoute pas
ces supports aux dépendances par défaut des fixtures. Le profil complet exige
le checkout QuickCheck avec ses deux fichiers d'adaptation Go. Leur publication
dans un fork installable n'a pas été réalisée : l'installateur vérifie ce
prérequis avant tout clone. Les anciennes limites des tests, des noms de paquets
et des FFI restent attribuées aux lots suivants ; cette validation ne vaut pas
campagne intégrale des bibliothèques ou exécution de leurs CI JavaScript.

Les détails, copies avant/après, revue par fichier et journaux sont temporaires
dans `/tmp/gopurs-lot2-20260914/`. Aucune mesure de performance n'est déduite de
ce lot ; les baselines restent celles du README d'altbak.
