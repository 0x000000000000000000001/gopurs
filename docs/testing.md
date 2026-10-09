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
| Frontière TAST/PBO, types globaux et spécialisation | `node --test tools/monomorphization.test.mjs tools/global-types.test.mjs tools/preparation.test.mjs tools/foreign-forwarders.test.mjs`, après le build ; préparation native ci-dessous |
| Annotations et représentations natives | `node --test tools/native-record-workers.test.mjs tools/boxed-record-arguments.test.mjs tools/native-sum-results.test.mjs`, après le build |
| Preuves source/TCO des arguments record projetés | `node --test tools/native-record-args.test.mjs tools/native-record-workers.test.mjs`, après le build ; `./bin/test NativeRecordWorkers NativeRecordReturns` |
| Layouts, métadonnées et instanciation | `node --test tools/representation-contract.test.mjs tools/mixed-constructor-tags.test.mjs tools/elided-constructor-payloads.test.mjs tools/record-tuple-conversions.test.mjs`, après le build |
| Boxing et émission transitive Rebox | `node --test tools/rebox-generation.test.mjs tools/rebox-metadata.test.mjs tools/struct-pointer-boxing.test.mjs tools/value-array-unboxing.test.mjs`, après le build |
| Littéraux composites et constructeurs | `node --test tools/composite-expressions.test.mjs tools/elided-constructor-payloads.test.mjs tools/nullary-constructor-bindings.test.mjs`, après le build ; `./bin/test ConstructorReuse NativeArrayReboxing` |
| Types et records | `./bin/test NativeRecordBoxing NativeRecordSizes -c` |
| Dérivations et portée des instances | `./bin/test DerivingContravariant DerivingFunctorFromBi DerivingFunctorFromPro DerivingProfunctor` ; [comparaison des frontends](#plan-v2--lot-06--dérivations-et-imports-6-octobre-2026) |
| Chaînes TAST et labels de records | `./bin/test StringEdgeCases CompilerHostStrings` ; `node --test tools/record-tuple-conversions.test.mjs`, après le build ; décodeurs PBO ci-dessous |
| Bridge FFI | `node --test tools/ffi-bridge.test.mjs tools/ffi-generics.test.mjs`, après le build ; `./bin/test FFIIntegerReturns -c` |
| FFI JS/Go du compilateur et embarquement | `node --test tools/native-ffi.test.mjs tools/embed-runtime.test.mjs tools/go-imports.test.mjs`, après le build ; `native-go-code.test.mjs` avec `GOPURS_NATIVE_OUTPUT` |
| Stockage et durée de vie du runtime | `node --test tools/runtime-contracts.test.mjs tools/closure-lifetime.test.mjs tools/apply-arity.test.mjs tools/function-data.test.mjs` ; [contrats](runtime-ffi-contracts.md) |
| Appels et fonctions | `./bin/test CurriedLambdas -c` |
| Conversions de tableaux | `./bin/test ArrayRoundtrip -c` |
| Récursion | `./bin/test TCO TCOMutRec -c` |
| Boucles scalaires en `int32` | `node --test tools/int32-loops.test.mjs tools/integer-boundaries.test.mjs tools/binding-contracts.test.mjs`, après le build ; `./bin/test Int32Loops TCO ThunkFusion 2136` ; [mesures et campagne à trois modes](int32-loops.md) |
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

## Contrat de la frontière TAST/PBO

`monomorphization.test.mjs` contrôle les entrées publiques de l'adaptateur gopurs :
invalidation des faits source avant la collecte, conservation des corps utiles,
distinction entre barrière intrinsèque précoce et exclusions tardives, admission
par les types d'origine et ordre des modules. Il compare aussi la sortie pure
avec le chemin `Aff` utilisant réellement le collecteur PBO et `Preparation`.

`global-types.test.mjs` fixe les priorités binding/expression/FFI, le maintien
d'un `Any` explicite, les qualifications et le repli limité sur les applications.
Les contrats de représentations et de Rebox contrôlent les tables originales
et enrichies. `foreign-forwarders.test.mjs` vérifie la reconnaissance exacte
des wrappers, indépendamment de l'admission finale de leurs spécialisations.

`preparation.test.mjs` vérifie le différé, la réexécution, l'unicité des appels et
l'ordre des résultats en JS. Les tests natifs exercent le même `Preparation`
compilé, avec des travaux bloqués par des canaux et terminés hors ordre :

```sh
GOPURS_NATIVE_OUTPUT=/chemin/bootstrap/output node --test tools/preparation-native.test.mjs
```

Ils exigent les marqueurs PASS du différé, de la réexécution, du chevauchement,
de la borne à huit et des modes séquentiels sous `go test -race`. Le délai de
compilation à froid du gros package bootstrap est distinct du délai de 30 s
fixé pour exécuter les tests Go.

Les suites PBO `monomorphize-transitive`, `transitive-parallel`,
`monomorphize-callsite`, `monomorphize-cache` et `source-usage` s'exécutent avec
`node ../../purescript-backend-optimizer-gopurs/test/<suite>.mjs output`. Elles
contrôlent le moteur consommé par cet adaptateur : propagation, cache, priorité
des contributions, arités et portée des annotations ; elles ne reconstruisent
pas une autre copie du backend.

### Chaînes et labels du TAST

Le frontend encode les `PSString` contenant des surrogates isolés en tableaux
d'unités UTF-16. Les champs `TypeLevelString.value` et `Row.fields.label` doivent
accepter ces tableaux, recomposer les paires valides et conserver les surrogates
isolés en WTF-8 côté Go. La référence PureScript et les deux chemins natifs
(JSON déjà analysé et texte JSON indexé) disposent de régressions PBO :

```sh
node ../../purescript-backend-optimizer-gopurs/test/type-table-strings.mjs output
# GOPURS_NATIVE_OUTPUT désigne le output du bootstrap conservé.
cp ../../purescript-backend-optimizer-gopurs/test/type-table-strings_test.go \
  "$GOPURS_NATIVE_OUTPUT/purescript/type_table_strings_test.go"
go -C "$GOPURS_NATIVE_OUTPUT" test -race ./purescript -run '^TestTypeTablePSString' -count=1 -v
```

Ces tests couvrent les valeurs et labels, les paires valides/inversées/incomplètes,
les unités invalides et les chemins d'erreur. Après modification de `Json.go`,
exécuter `python3 bin/json-text/generate.py`, puis `--check`, depuis PBO.
`record-tuple-conversions.test.mjs` compile et exécute le Go produit pour les
labels inhabituels : champs natifs, boxing/unboxing, accès, mises à jour, tailles
de dictionnaires 1 à 7 et collision avec le préfixe d'encodage réservé.

## Contrat de sortie CLI

`npm run test:cli` compile une fixture réelle dans un workspace temporaire,
puis appelle `bin/gopurs` en JS, natif séquentiel et natif parallèle. Chaque mode
vérifie le succès (statut 0 et Go identique) et cinq erreurs réelles (statut 1) :
chargement du TAST, écriture du runtime, écriture d'un module avec workers PBO,
FFI invalide et écriture de l'entrée exécutable.
Le frontend de la fixture est sélectionné comme pour le bootstrap natif :
`GOPURS_PURS` si renseigné, sinon le fork local sous `../../purescript`.

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

## Contrat des schémas de décodeurs

`decoder-schemas.test.mjs` exerce la façade sur les formes exactes d'applications,
les ABI, les proxies et symboles, les limites de taille/profondeur, les alias et
les barrières récursives. Les tests contrôlent les annotations, la réservation de
toute la famille de noms et la publication des sources dans l'ordre de visite.

Les corps personnalisés sont contrôlés avec leurs vraies lectures et branches :
provenance de l'objet, distinction absent/null, schémas complets, transmission du
payload `Left`, masquage des niveaux, constructeurs annotés et captures triées.
Deux programmes Go exécutent les workers émis avec les helpers Argonaut réels
et `go test -race`. Ils vérifient les labels dupliqués, la propriété des chaînes,
le repli texte sur une méthode opaque, les choix ordonnés, les lectures imbriquées
et les erreurs exactes. Les getters de construction du second programme sont
produits par le générateur Go ordinaire ; les tests du helper `CompiledSchema`
sont également exécutés dans ces workspaces isolés.

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
la fixture concernée. Par défaut, la campagne s'arrête au premier échec ; les mises à jour
des fixtures déjà réussies restent écrites. Le contrôle d'exécution conserve
le contrat historique : statut zéro et absence de `Fail` dans la sortie.

## Isolation, logs et caches

Chaque fixture reçoit son propre répertoire temporaire : sources, configuration
Spago, lockfile, `.spago`, `output` et `tmp`. `TMPDIR`, `TMP` et `TEMP` désignent
ce dernier pour toutes ses commandes, afin d'isoler les extractions Spago entre
campagnes concurrentes. `tests/runner` n'est ni lu ni modifié.
Les `.go`, `.js` et répertoires compagnons d'une fixture sont copiés avec sa
source. `-- @dependencies: assert prelude effect console` peut limiter ses
dépendances ; sans directive, la liste de `bin/pkg` est utilisée. Les checkouts
core restent requis et les paquets utilisent le cache global de Spago.

Les logs distinguent compilation PureScript, génération Go, formatage,
snapshots, dépendances Go, compilation Go et exécution. Les rapports et logs sont
toujours conservés ; les autres fichiers d'une cible réussie sont supprimés,
sauf avec `--keep-workspace`. Échecs et interruptions conservent
le workspace et affichent son chemin. SIGINT/SIGTERM sont transmis à la
commande active et ses sous-processus. Une compilation PureScript échouée
n'est pas relancée automatiquement.

Les deux runners affichent `Results: /…/results.json`. Le rapport versionné est
remplacé atomiquement avant/après chaque cible ; il contient toute la sélection,
les états `pending`, `running`, `passed`, `failed` ou `interrupted`, les chemins,
les dates et les diagnostics. `--keep-going` continue après un échec individuel
et rend un statut non nul dès qu'une cible échoue. Sans cette option, les cibles
suivantes restent `pending` dans le bilan.

```bash
./bin/test --all --keep-going
./bin/test --resume-failed /chemin/results.json --list
./bin/test --resume-failed /chemin/results.json --keep-going
```

`--resume-failed` sélectionne les cibles non réussies du rapport, y compris un
`running` laissé par un arrêt brutal. Il vérifie le type de campagne et la racine
du dépôt, puis crée une nouvelle campagne avec `resumedFrom`. Les anciens résultats
restent intacts ; un succès antérieur n'est pas présenté comme un succès sur les
sources actuelles. Les noms, `--all` et `--skip-before` ne se combinent pas avec
cette option. La vérification stricte des snapshots reste le défaut de la reprise.

`-c` ne réinitialise pas les caches globaux : il reconstruit gopurs. Les sorties
de fixture sont neuves avec ou sans cette option. À l'intérieur du backend,
chaque lancement régénère le Go et relit la FFI ; changer seulement une FFI Go
ne demande pas de recompiler le PureScript si les entrées TAST sont inchangées.

## Exclusions et modules frères

Les exclusions de [tools/test-selection.mjs](../tools/test-selection.mjs) ont été
réexécutées le **2 octobre 2026**, avec le même compilateur que la campagne
finale. Après réintégration de `StringEdgeCases` le **3 octobre**, de
`StringEscapes`, `2136` et `NumberLiterals` le **5 octobre**, puis des quatre
fixtures de dérivation le **6 octobre**, **aucune exclusion ne subsiste**.

`DerivingContravariant`, `DerivingFunctorFromBi`, `DerivingFunctorFromPro` et
`DerivingProfunctor` échouaient aussi avec le frontend amont : leurs imports
ne rendaient pas toutes les instances nécessaires disponibles en compilation
isolée. Les imports explicites corrigent les quatre cas ; leurs méthodes
dérivées sont maintenant exécutées et vérifiées dans les trois modes. Voir le
[lot 06 du plan v2](#plan-v2--lot-06--dérivations-et-imports-6-octobre-2026).

`DerivingClause` réintègre la sélection : compilation et exécution Go réussies,
**334 fichiers Go identiques** entre natif séquentiel, natif parallèle et JS
sur les mêmes entrées TAST. Son snapshot, auparavant absent, a été créé puis
revérifié en mode strict. Les snapshots existants n'ont pas été remplacés.
Cette campagne comptait **391 fixtures sélectionnables**, distinctes des huit exclusions.

`StringEdgeCases` réintègre aussi la sélection après correction du décodage TAST
et des labels de records : exécution dans les trois modes, parité octet par octet,
création puis vérification stricte du snapshot. Les preuves figurent au
[lot 02 du plan v2](#plan-v2--lot-02--chaînes-tast-et-records-3-octobre-2026).

`StringEscapes` réintègre la sélection avec l'assertion de concaténation des
surrogates réactivée et huit cas maintenus à l'exécution par `Effect.Ref`.
Le pliage et l'exécution respectent l'oracle JS dans les trois modes ; son
snapshot est créé puis vérifié strictement au
[lot 03 du plan v2](#plan-v2--lot-03--concaténation-utf-16-5-octobre-2026).

`2136` réintègre la sélection après correction des opérations entières aux bornes
32 bits. Son prédicat original est conservé comme assertion, avec des contrôles
de constantes et d'exécution ; les trois modes rendent l'oracle JS et leur Go
est identique. Le snapshot revu passe strictement au
[lot 04 du plan v2](#plan-v2--lot-04--bornes-et-opérations-int-5-octobre-2026).

`NumberLiterals` réintègre la sélection avec l'oracle du Prelude JS actuel et
les corrections Go de zéro signé et d'exposants. Les trois modes produisent
le même Go et exécutent les 41 cas littéraux et dynamiques ; le nouveau snapshot
est revu puis vérifié strictement au
[lot 05 du plan v2](#plan-v2--lot-05--affichage-des-number-5-octobre-2026).

`bin/modtest` sélectionne les checkouts frères `gopurs-*` possédant un
`bin/test` exécutable :

```bash
./bin/modtest --all --list
./bin/modtest --skip-before strings --list
./bin/modtest prelude strings
```

La sélection complète est le défaut. Les noms avec ou sans `gopurs-` sont
acceptés ; `-c` reconstruit le backend depuis ce checkout. Chaque cible reçoit
sa propre copie des sources de toute la famille de bibliothèques, sans `.git`,
`node_modules`, `.spago`, `.cache` ni `output`. Les liens relatifs de configuration
restent relatifs. Le compilateur préconstruit est partagé par lien et les scripts
s'exécutent dans les copies, avec des temporaires privés. Les runners frères
nettoient uniquement leur propre checkout lors d'une invocation directe.
Les options de bilan, reprise et conservation ci-dessus s'appliquent aussi à
`bin/modtest` ; son log par cible est `logs/test.log`.

Dans `spec`, l'environnement d'intégration copie le template en lecture seule,
ignore ses éventuels `node_modules`/`output`, puis remplace `SPEC_REPO_PATH`.
Il utilise `spago` sur le `PATH` et le lanceur `bin/gopurs` avec les variables du
mode appelant, y compris pour les neuf cas imbriqués. Une initialisation échouée
ou annulée détruit sa copie avant une nouvelle tentative. Son `bin/test` exécute
aussi les deux régressions de `test/integration-environment.mjs`, contre le vrai
programme Go `Test.IntegrationEnvironment` et des commandes externes simulées.

QuickCheck possède désormais `package.test.main: Test.Go` et un `bin/test`
exécutable, sélectionné par `./bin/modtest quickcheck`. La suite utilise une
graine fixe et compare les sorties JavaScript/Go. Les trois `pending` de `spec`
sont des fixtures intentionnelles ; leurs contrats d'exécution et d'affichage
sont vérifiés dans le parcours standard. Voir le
[lot 07 du plan v2](#plan-v2--lot-07--couverture-des-bibliothèques-6-octobre-2026).

## Plan v2 — lot 01 : campagnes reproductibles, 3 octobre 2026

Les runners partagent désormais les rapports/reprises et les temporaires privés.
`modtest` copie les bibliothèques par cible ; les **46 scripts** qui nettoyaient
leurs voisins ont été corrigés. `spec` initialise une copie autonome du template
et ne publie sa référence qu'une fois la préparation terminée.

Lot validé : **15/100 points, 1/8 lots** du plan v2. Vérifications terminées :

- **70 tests Node** de runners/bootstrap, dont les nettoyages directs des
  **50 scripts** installés, les rapports complets, les reprises, les snapshots
  stricts et les signaux aux descendants. Les quatre nouveaux scénarios de
  campagne et les 46 nettoyages défectueux ont d'abord échoué sur la référence.
- **Deux régressions natives de `spec`** échouent avant correction : lecture de
  `node_modules` absent et répertoire incomplet laissé après initialisation
  échouée. Elles passent après correction, dont une nouvelle tentative sur le
  même environnement. Elles sont intégrées à son `bin/test`.
- Les quatre fixtures touchées par les erreurs d'extraction de la campagne v1
  (`OperatorAliasElsewhere`, `PendingConflictingImports2`, `PolykindBindingGroup1`,
  `PolykindInstantiatedInstance`) passent en **deux campagnes concurrentes**,
  avec snapshots stricts et exécution Go. `TCOMutRec` et `ThunkFusion` passent
  aussi pendant les contrôles de reprise.
- Erreur PureScript réelle puis reprise : rapport `failed/passed`, nouvelle
  campagne limitée à la cible corrigée et rapport précédent préservé. SIGTERM
  pendant la compilation par le vrai `purs` : statut **143**, descendant arrêté,
  rapport `interrupted/pending`, puis reprise **2/2** réussie.
- Erreur YAML réelle dans une copie de `prelude` : rapport `failed/pending`,
  puis reprise de `prelude` et `strings` **2/2** après correction.
- Campagne complète sur des copies fraîches : **50/50 runners réussis**, dont
  49 exécutent du Go et `assert` vérifie sa compilation. `spec` réussit ses deux
  nouvelles régressions, puis **70 tests, dont huit intégrations imbriquées** ; ses
  **trois pending** restent consignés pour le lot 07.
- L'espace libre étant descendu sous 1 Gio, la campagne a été arrêtée par
  SIGTERM avant épuisement du disque : **41 réussites, une interruption de
  `spec`, huit cibles en attente**. Après `go clean -cache`, la commande
  `--resume-failed …/results.json --keep-going` réussit les **neuf cibles**,
  sans modifier le premier rapport. `real-modules-results.json` relie les deux
  bilans sur les mêmes sources et artefacts figés.
- Parité de `spec` sur **390 entrées TAST figées** : **490 fichiers Go et
  `go.mod` identiques octet par octet** entre JS, natif Go entièrement séquentiel
  (tous les workers à 1, pipeline désactivé) et parallèle (workers à 4).
- Préservation vérifiée des **5 506 fichiers/liens** de la famille source
  copiée et des **2 160 entrées** de caches/sorties des checkouts d'origine.
  Les sources exécutées correspondent aux fichiers livrés. L'audit des liens
  locaux et ancres, la syntaxe des scripts modifiés et `git diff --check` passent.

Les preuves avant/après, scripts de reproduction, rapports, empreintes des
artefacts et journaux sont conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-campaigns-koy5tymo/`.
Toolchain : Node **24.8.0**, Go **1.27.0**, Spago **1.0.3**, même frontend TAST
**0.15.16 development** que la clôture v1 ; `toolchain.json` conserve les SHA-256
du frontend et des deux compilateurs utilisés.

## Plan v2 — lot 02 : chaînes TAST et records, 3 octobre 2026

`StringEdgeCases` est réintégrée. Lot validé : **30/100 points, 2/8 lots** du
plan v2.

La reproduction avant correction rejetait `Symbols.typeTable.value` et
`Records.typeTable.fields.label` en natif ; **six fichiers** divergeaient du
backend JS. L'exécution Go échouait aussi côté JS : champs `�`/emoji invalides et
clés de records émises sans échappement. L'exécution directe du frontend JS
réussissait et fournit l'oracle `Done\nDone\n`.

Le décodeur natif PBO accepte maintenant les chaînes ou tableaux d'unités UTF-16
`0..65535`, recompose les paires valides, conserve les surrogates isolés en WTF-8
et garde les chemins d'erreur. `Json/Text.go` est régénéré depuis `Json.go` ; le
générateur reproduit également la politique existante des littéraux entiers et
le parcours linéaire des tableaux de littéraux. Côté gopurs, `recordFieldName`
encode les labels inhabituels avec le préfixe réservé `gopurs_field_`, et les clés
passent par l'échappement des littéraux Go dans toutes les conversions et
opérations de records.

Vérifications terminées sur les copies figées :

- Régressions avant/après : les deux décodeurs natifs échouent sur les tableaux
  UTF-16 avant correction ; le test exécutable des labels échoue à la compilation
  Go. Tous passent après correction.
- **33 tests Node** sur les records et représentations, dont le nouveau test
  compilant/exécutant les labels isolés, astraux, guillemets, antislash et saut de
  ligne, les tailles 1 à 7, les mises à jour et la réservation du préfixe.
- **Neuf tests Node PBO**, plus la suite `type-table.mjs` : référence PSString,
  valeurs, labels, erreurs, références et ordre de résolution.
- **Cinq tests Go PBO sous `-race`**, couvrant les deux chemins PSString, les
  arguments/résolutions de types et **2 025 cas** de frontière parser/schéma.
- Bootstrap JS et natif réussi : **500 modules, 287 718 types**. La vérification
  `python3 bin/json-text/generate.py --check` passe sur les sources livrées.
- **271 entrées TAST identiques** avant/après et entre les trois modes :
  **336 fichiers Go et `go.mod` identiques octet par octet** entre JS, natif Go
  séquentiel (workers à 1, pipeline désactivé) et parallèle (workers à 4).
  Les trois exécutables produisent exactement l'oracle JS.
- Revue complète du diff généré : seul **`Records.go`** change par rapport à
  l'ancien JS, pour les noms de champs et l'échappement des clés. Les six écarts
  de l'ancien natif disparaissent après décodage complet de `Records`/`Symbols`.
- Snapshot `StringEdgeCases.go` créé après revue/exécution, puis **12 fixtures**
  réussies en mode strict sur le compilateur final : `StringEdgeCases`,
  `NativeRecordBoxing`, `NativeRecordSizes`, `NativeRecordReturns`,
  `NativeRecordWorkers`, `RecordTypeChangingUpdate`, `NestedRecordUpdate`,
  `NewtypeWithRecordUpdate`, `PolyLabels`, `CompilerHostStrings`,
  `BlockStringEdgeCases` et `CompactRecordConsumers`.
- Audit des **184 entrées source** figées du compilateur : les six fichiers
  source attendus changent. Les corrections, dépendances et tests livrés
  correspondent aux copies validées ; liens locaux et `git diff --check` passent.

Preuves, empreintes, sorties avant/après et rapports conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-tast-strings-5p8w6c8y/`,
notamment `fixture-before.json`, `fixture-after.json`, `generated-go-review.diff`,
`bootstrap-after-final.log`, `pbo-native-after.log`, `fixtures-strict.log` et
`final-results.json`. Toolchain : Node **24.8.0**, Go **1.27.0**, Spago **1.0.3**,
frontend TAST **0.15.16 development**. Le nettoyage préalable demandé a libéré
**1,07 Gio** de caches et anciens workspaces de tests ; `gopurs-cleanup.json`
confirme des statuts Git identiques avant/après suppression.

## Plan v2 — lot 03 : concaténation UTF-16, 5 octobre 2026

`StringEscapes` est réintégrée. Lot validé : **45/100 points, 3/8 lots** du
plan v2.

Avant correction, le natif pliait `loneSurrogates` à `false`, contre `true` en
JS : seul `Main.go` divergeait. L'assertion réactivée échouait en natif ; le Go
produit par JS échouait sur le nouveau contrôle de concaténation à l'exécution.
Le test isolé échouait aussi pour le code émis et la FFI Prelude :
`ed a0 80` suivi de `ed b0 80` restait six octets au lieu de `f0 90 80 80`.

`runtime.ConcatString` recompose maintenant la paire pouvant apparaître à la
jonction de deux chaînes WTF-8 canoniques. Les autres octets, dont les surrogates
isolés, sont préservés. `PrimitiveExprs` émet ce helper pour `OpStringAppend` et
`gopurs-prelude/src/Data/Semigroup.go` l'utilise également. Le bootstrap donne
ainsi au pliage du compilateur natif la même sémantique qu'aux programmes produits.

Vérifications terminées sur les copies figées :

- **801 couples UTF-16 et quatre cas d'associativité**, comparés à l'oracle JS
  pour les deux chemins, sous `-race`. Ils couvrent chaînes vides, caractères
  BMP/astraux, bornes des surrogates, ordre inversé, préfixes et suffixes.
- **69 tests Node réussis** : quatre contrôles runtime/concaténation, dont les
  contrats de durée de vie normalement et sous `-race`, et 65 contrôles des
  runners et de l'embarquement exact du runtime.
- Bootstrap JS et natif réussi : **499 modules, 287 679 types**.
- **269 entrées TAST identiques** avant/après et entre JS, natif Go séquentiel
  (workers à 1, pipeline désactivé) et parallèle (workers à 4, pipeline activé).
  Les **334 fichiers Go et `go.mod` sont identiques octet par octet** ; les trois
  exécutables produisent exactement l'oracle du frontend JS, `Done\n`.
- Revue de tous les écarts avec l'ancien JS : **283 remplacements de
  concaténation dans 69 fichiers**, l'ajout du helper au runtime et l'appel
  correspondant dans la FFI Prelude. Les **264 autres fichiers** sont inchangés.
  L'analyse Go inverse les seuls appels au helper et retrouve exactement
  l'ancien code après gofmt, en conservant l'ordre et le parenthésage.
- **150 fixtures compilées et exécutées** pour examiner les snapshots contenant
  une addition et les cas voisins de chaînes. **24 snapshots** sont actualisés
  après revue de leurs **157 remplacements de concaténation** ; le nouveau
  `StringEscapes.go` correspond à la sortie exécutée dans les trois modes.
  Les **128 autres snapshots examinés** restent identiques.
- Après mise à jour, **28 fixtures réussies avec snapshots stricts et exécution
  Go** : les 25 fixtures actualisées, plus `StringEdgeCases`,
  `CompilerHostStrings` et `BlockStringEdgeCases`.
- Les runners **`prelude` et `strings` réussissent tous les deux** dans leurs
  copies isolées. La première collecte de snapshots a rencontré `ENOSPC` après
  78 réussites ; les **72 cibles restantes** ont réussi à la reprise avec un
  cache Go privé par shard, réduit entre les cibles. Les premiers rapports et
  diagnostics sont conservés ; les caches privés de collecte sont supprimés.
- Audit des **2 940 entrées source archivées** : seuls les correctifs attendus
  et les snapshots revus changent. Les six fichiers de correction/test/sélection
  et les 25 snapshots livrés correspondent aux copies validées. Liens locaux,
  ancres, syntaxe du nouveau test et `git diff --check` vérifiés.

Les **53 dépôts archivés** sont identifiés dans `source-heads.json`, dont gopurs
`8581703`, PBO `157a544` et Prelude `82e63fe`. Preuves, scripts, rapports et
empreintes conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-utf16-aaiyq81c/`,
notamment `REPRODUCTION.md`, `fixture-before.json`, `fixture-after.json`,
`generated-go-review.diff`, `snapshot-review.diff`, `fixtures-strict.log` et
`final-results.json`. Toolchain : Node **24.8.0**, Go **1.27.0**, Spago **1.0.3**,
frontend TAST **0.15.16 development**.

## Plan v2 — lot 04 : bornes et opérations Int, 5 octobre 2026

`2136` est réintégrée. Lot validé : **55/100 points, 4/8 lots** du plan v2.

La référence est constituée des FFI JS de Prelude et `Data.Int.Bits`, exécutées
directement comme oracles. Avant correction, onze opérations échouent sur les
deux chemins Go testés ; division et modulo passent déjà. Le Go produit par
chacun des trois compilateurs échoue sur le prédicat original de `2136`, et
seul `Main.go` diffère entre JS et natif lors du pliage des constantes.

Les helpers du runtime sont partagés par `PrimitiveExprs` et les FFI Go de
`Data.Semiring`, `Data.Ring` et `Data.Int.Bits` : négation/addition/soustraction
avec débordement signé sur 32 bits, opérations binaires sur 32 bits et comptes
de décalage masqués par `31`. La multiplication reproduit **`(a * b) | 0`**,
avec l'arrondi `Number` avant conversion : `top * top` vaut donc **0**.
Le stockage `int64` conserve les résultats non signés de `zshr` et le quotient
`bottom / -1 == 2147483648`. Ces contrats sont détaillés dans
[runtime-ffi-contracts.md](runtime-ffi-contracts.md#opérations-sur-int).
Le bootstrap corrige aussi le pliage natif avec ces mêmes primitives.

Vérifications terminées sur les copies figées :

- **5 978 cas par chemin, soit 11 956 comparaisons sur 13 opérations**, sous
  `-race`, contre les quatre vraies FFI JS. Les chemins code émis et FFI Go
  couvrent bornes signées/non signées, comptes négatifs ou supérieurs à 31,
  couples pseudo-aléatoires, ordre et évaluation unique des opérandes.
  Les quatre oracles sont identiques à l'octet près aux `foreign.js` exécutés
  par le frontend.
- `2136` garde son prédicat original sous une assertion, ajoute cinq contrôles
  avec opérandes issus de littéraux et quinze contrôles dynamiques via
  `Effect.Ref`. Les trois modes rendent exactement l'oracle JS **`Done\n`**.
- **114 tests Node réussis** : sept contrôles entiers/division/runtime et
  embarquement, puis 107 tests voisins de génération. Les runners **`prelude`
  et `integers` réussissent tous les deux** dans leurs copies isolées.
- Bootstrap JS et natif réussi : **499 modules, 287 679 types**.
- **269 entrées TAST identiques** avant/après et entre JS, natif Go séquentiel
  (workers à 1, pipeline désactivé) et parallèle (workers à 4, pipeline activé).
  Les **334 fichiers Go et `go.mod` sont identiques octet par octet**.
- Revue complète du Go produit : **221 remplacements de primitives dans
  31 fichiers**, plus le runtime et les trois FFI revus séparément ; les
  **300 autres fichiers** restent identiques à l'ancien JS. L'inversion des
  seuls appels aux helpers retrouve exactement la référence après gofmt.
- **82 fixtures compilées et exécutées**, avec **83 snapshots examinés** :
  40 identiques et 43 installés après revue. Parmi ces derniers, 36 comportent
  uniquement les **3 991 remplacements Int**, et le nouveau `2136.go` correspond
  exactement à la sortie exécutée dans les trois modes.
- Les six autres snapshots (`ArrayRoundtrip`, `DerivingClause`, `FieldConsPuns`,
  `FieldPuns`, `QualifiedDo`, `test-int`) rattrapent **12 concaténations du lot 03**.
  Son filtre textuel ` + ` avait manqué les opérateurs sans espaces. L'analyse
  Go confirme que ces seuls appels à `ConcatString` expliquent leurs écarts.
- Après installation, **47 fixtures réussies avec snapshots stricts et
  exécution Go**, via le vrai `bin/test` et `UPDATE_SNAPSHOTS=0`. Cette passe
  couvre tous les snapshots actualisés, `NegativeIntInRange`, `SolvingAddInt`,
  `SolvingMulInt`, `FFIIntegerReturns`, `test-int` et `StringEscapes`.
- Audit des **2 942 entrées source archivées** : les huit fichiers livrés de
  correction/test/sélection et les 43 snapshots correspondent aux copies
  validées. Liens locaux, ancres, syntaxe du nouveau test, réintégration effective
  de `2136` et `git diff --check` vérifiés.

Les **53 dépôts archivés** et les **34 fichiers du lot 03 superposés** sont
identifiés dans `source-heads.json` et `lot03-overlay.json` : notamment gopurs
`8581703`, PBO `157a544`, Prelude `82e63fe` et integers `37b9dbe`.
Les preuves, rapports, scripts et empreintes sont conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-int32-d3hedy_d/`,
notamment `REPRODUCTION.md`, `integers-red.log`, `integers-green.log`,
`fixture-before.json`, `fixture-after.json`, `generated-review.json`,
`snapshot-review.json`, `snapshot-review.diff`, `fixtures-strict.log` et
`final-results.json`. Toolchain : Node **24.8.0**, Go **1.27.0**, Spago **1.0.3**,
frontend TAST **0.15.16 development**.

### Suivi du lot 04 — installation locale et `rung`

Le signalement `rung` a révélé que le natif du checkout actif datait encore du
5 octobre à 03 h 12 : les copies figées étaient validées, mais ce binaire local
embarquait l'ancien runtime. Les FFI actualisées appelaient donc des helpers
absents. Le runner `altbak.pub-gopurs/bin/native/driver.py` reconstruisait seulement
le JS avec `--clean`, puis lançait le natif par défaut.

Le natif local a été reconstruit. Le runner choisit maintenant `build:native`
par défaut, `build` pour `GOPURS_JS=1` et `build:rust` pour `GOPURS_RUST=1`.
La régression reproduit le runtime périmé avant correction et réussit après :
**sept tests du runner et douze tests de dispatch réussis**. La fixture isolée
du test d'empreintes JSON fournit aussi son fichier frontend factice manquant.
La commande réelle `./bin/go/run --clean` réussit la reconstruction, le build
Go et la validation des **14 résultats du benchmark**. Le runtime produit est
identique à la source canonique ; les **831 fichiers source/configuration**
inventoriés du compilateur, de PBO et des bibliothèques sont préservés.

Journaux avant/après, anciens binaires, empreintes et résultats conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-local-install-7cj5t5s4/`,
notamment `runner-red.log`, `runner-green.log`, `bootstrap.log`, `rung-after.log`
et `final-results.json`. Le plan reste à **55/100, 4/8 lots**.

## Plan v2 — lot 05 : affichage des Number, 5 octobre 2026

`NumberLiterals` est réintégrée. Lot validé : **65/100 points, 5/8 lots** du
plan v2.

La fixture initiale échoue aussi en JS : son ancien oracle à 14 chiffres attend
`0.25996181067142` pour le littéral `0.25996181067141905`. Les valeurs attendues
sont alignées sur la FFI JS Prelude réellement exécutée, également employées par
la fixture actuelle du frontend local. L'identité octet par octet de `Show.js`
et du `Data.Show/foreign.js` exécuté est enregistrée.

Deux écarts Go sont reproduits séparément : `show (-0.0)` rendait `-0.0` au lieu
de `0.0`, et les exposants étaient complétés par un zéro (`e-08` contre `e-8`).
La FFI `Data/Show.go` conserve les chiffres les plus courts retrouvant le même
binary64, les seuils JS de notation décimale `1e-6`/`1e21`, le suffixe `.0` des
entiers décimaux et les valeurs spéciales. Elle normalise maintenant les zéros
signés et les exposants. Le bootstrap applique aussi ce contrat à l'écriture
des littéraux par le compilateur natif ; le signe des valeurs reste préservé.
Voir les [contrats Number](runtime-ffi-contracts.md#affichage-des-number).

Vérifications terminées :

- **32 527 cas binary64 comparés à la vraie FFI JS sous `-race`** : zéros signés,
  NaN/infinis, chacun des exposants binaires finis, sous-normaux, voisins des
  puissances de dix, frontières de notation et 4 096 motifs pseudo-aléatoires.
  Le test échoue avant la correction et réussit après.
- **41 cas dans `NumberLiterals`, chacun vérifié deux fois**, depuis son littéral
  puis via `Effect.Ref` : 82 comparaisons exécutées. La version renforcée échoue
  avec les trois anciens compilateurs ; après correction, tous rendent l'oracle
  frontend JS **`Done\n`**.
- Bootstrap JS et natif avant/après réussi : **499 modules, 287 679 types**.
- **269 entrées TAST identiques** avant/après et entre JS, natif Go séquentiel
  (workers à 1, pipeline désactivé) et parallèle (workers à 4, pipeline activé).
  Les **334 fichiers Go et `go.mod` sont identiques octet par octet**. Le seul
  changement par rapport à l'ancien JS est **`Data_Show_ffi.go`**, revu en entier ;
  l'ancien écart natif de notation des littéraux dans `Main.go` disparaît.
- `CompilerHostNumbers` conserve le signe des zéros pliés et dynamiques :
  exécution réussie dans les trois modes, sur **60 entrées TAST**, avec
  **86 fichiers Go et `go.mod` identiques**.
- **11 tests Node réussis** : la régression numérique et dix contrôles des
  bridges FFI, génériques, callbacks et forwarders.
- Nouveau snapshot **`NumberLiterals.go` revu avant installation**, correspondant
  exactement au Main exécuté dans les trois modes après gofmt. **27 fixtures
  passent avec snapshots stricts et exécution Go**, dont tous les snapshots
  appelant Show Number, ainsi que `2136` et `StringEscapes`. Les 28 snapshots
  existants, dont deux FFI, restent identiques.
- Les runners **`prelude` et `numbers` réussissent** dans leurs copies isolées.
- Le parcours réel **`altbak ./bin/go/run --clean` réussit** : reconstruction du
  natif actif, compilation Go et **14 résultats du benchmark validés**. Cette
  reconstruction locale couvre 500 modules et 288 603 types ; ses 847 fichiers
  source/configuration inventoriés, dont les travaux PBO concurrents, sont
  préservés. Sa FFI Show générée est identique à celle de la validation figée.
- Audit des **2 944 fichiers archivés** : les quatre fichiers de correction,
  test et sélection, ainsi que le nouveau snapshot, correspondent aux copies
  validées. Liens et ancres, syntaxe du test, sélection effective de la fixture
  et `git diff --check` vérifiés.

Les **53 dépôts archivés** sont identifiés par leurs empreintes et HEAD,
notamment gopurs `141d6fc`, PBO `157a544` et Prelude `3ee1343`.
Preuves, scripts et rapports conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-number-show-mw16tagz/`,
notamment `REPRODUCTION.md`, `original-js.log`, `show-red.log`, `show-green.log`,
`fixture-before.json`, `fixture-after.json`, `generated-go-review.diff`,
`snapshot-review.json`, `fixtures-strict.log`, `host-results.json`,
`local-application.log` et `final-results.json`. Toolchain : Node **24.8.0**,
Go **1.27.0**, Spago **1.0.3**, frontend TAST **0.15.16 development**.

## Plan v2 — lot 06 : dérivations et imports, 6 octobre 2026

Les quatre fixtures de dérivation sont réintégrées. Lot validé :
**75/100 points, 6/8 lots** du plan v2.

Les originaux sont identiques aux fixtures du frontend local. Leur compilation
isolée échoue avec `CannotDeriveInvalidConstructorArg` aussi bien sur le fork
TAST **0.15.16 development** que sur la référence amont **0.15.15**, avec les
mêmes sources de dépendances du package set **77.10.1**. Les modules
`bifunctors` et `profunctor` sont déjà présents dans les globs de compilation.
Le problème est la portée des instances : les modules qui les définissent
doivent appartenir à la fermeture des imports du module testé. Le harnais du
frontend local appelle `rebuildModule` avec tous les `supportExterns` pour ses
fixtures à module unique ; ce contexte masque les imports manquants.

La reproduction minimale ajoute uniquement les imports suivants aux originaux.
Les **huit compilations** réussissent alors, quatre par frontend :

| Fixture | Modules d'instances ajoutés | Assertions exécutées par la fixture renforcée |
| --- | --- | --- |
| `DerivingContravariant` | `Data.Bifunctor`, `Data.Profunctor` | 7 : tous les constructeurs, triple contravariance, fonctions, tuples, records imbriqués et argument quantifié |
| `DerivingFunctorFromBi` | `Data.Bifunctor`, `Data.Bifoldable`, `Data.Bitraversable` | 24 : `map`, `foldl`, `foldr`, `foldMap`, `traverse`, `sequence` et ordre des effets sur les trois constructeurs |
| `DerivingFunctorFromPro` | `Data.Profunctor` | 2 : double contravariance des fonctions, tableaux et records imbriqués |
| `DerivingProfunctor` | `Data.Bifunctor` | 6 : tous les constructeurs, transformations gauche/droite, champs constants et quantifiés, tuples et records imbriqués |

Les **39 assertions** utilisent une entrée lue via `Effect.Ref` pour conserver
les chemins dynamiques ; les transformations changent aussi les types des
paramètres. Les deux frontends signalent encore ces imports d'instances seuls
comme `UnusedImport` : les commentaires des fixtures expliquent leur nécessité.

Vérifications terminées :

- Comparaison avant/après sur chaque cas : **huit rejets initiaux**, puis
  **huit succès** après les seuls imports. Les diagnostics complets et les
  empreintes des originaux sont conservés.
- Fixtures renforcées compilées puis exécutées en JavaScript par **les deux
  frontends**, soit huit exécutions réussies avec exactement **`Done\n`**.
- Pour **chacune des quatre fixtures**, **268 entrées TAST figées** et
  **333 fichiers Go plus `go.mod` identiques octet par octet** entre le backend
  JS, le natif Go séquentiel (tous les workers à 1, pipeline désactivé) et le
  natif Go parallèle (workers à 4, pipeline activé). Les **douze compilations
  et exécutions Go** réussissent avec le même résultat que les frontends JS.
- Les quatre nouveaux snapshots sont revus : dictionnaires et branches dérivés,
  variance, conversions de tuples/records, préservation des champs constants
  et quantifiés, assertions et ordre des effets. Ils correspondent exactement
  aux `Main.go` exécutés dans les trois modes après gofmt.
- Création par le vrai `bin/test`, puis **4/4 fixtures réussies en mode strict**
  avec `UPDATE_SNAPSHOTS=0`, compilation et exécution Go.
- **63 tests Node du runner réussis**, zéro échec et zéro saut, après retrait
  des quatre exclusions et adaptation de sa fixture de sélection.
- `./bin/test --list` sélectionne les **400 fixtures présentes**, sans exclusion.
  Ce comptage est une vérification de sélection ; les validations d'exécution
  de cette passe portent sur les quatre fixtures ci-dessus.
- Empreintes des sources, dépendances et compilateurs vérifiées ; snapshots
  installés identiques aux références validées, liens locaux et `git diff --check`
  propres.

Preuves, scripts de reproduction, sources originales, inventaires TAST/Go et
journaux conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-deriving-bgLIyO/`,
notamment `REPRODUCTION.md`, `comparison-before.json`, `comparison-imports.json`,
`validation.json`, `snapshot-review.json`, `fixtures-strict.log`,
`runner-tests.log` et `final-results.json`. Toolchain : Node **24.8.0**, Go
**1.27.0**, Spago **1.0.3**. `validation.json` identifie les deux frontends et
les deux hôtes gopurs par leurs SHA-256 ; `sources.json` identifie les checkouts,
dont gopurs **`475c954`** et frontend **`b4a7fb1`**. Le binaire frontend utilisé
annonce le commit **`3c8fcfd… DIRTY`** : son empreinte, plutôt que le seul HEAD
du checkout, désigne la version effectivement testée.

Les workspaces volumineux de ce lot sont conservés dans `workspaces.tar.gz` ;
`archive.json` donne son empreinte et les **33 058 fichiers vérifiés** avant
suppression des copies non compressées. Les journaux et manifests restent
accessibles directement à la racine des preuves.

## Plan v2 — lot 07 : couverture des bibliothèques, 6 octobre 2026

QuickCheck dispose d'une suite Go autonome, intégrée au runner standard ; les
trois `pending` de `spec` sont examinés, documentés et couverts par des tests
actifs. Lot validé : **90/100 points, 7/8 lots** du plan v2.

### QuickCheck

`gopurs-quickcheck/bin/test` compile `Test.Go`, exécute l'oracle JavaScript puis
le programme Go et compare les deux transcriptions octet par octet. La
configuration Spago déclare les dépendances de test et les overrides Go du
graphe résolu ; `npm run test:go` expose le même parcours. La graine **12345**,
lue via `Effect.Ref`, conserve des entrées dynamiques et un résultat reproductible.

La suite couvre l'état et le rejeu des générateurs, la restauration de taille,
les intervalles Int/Number, les combinators, les records imbriqués, les sommes
génériques et les fonctions arbitraires. **Seize cas de perturbation Float32**
exercent la FFI, dont les zéros signés, l'arrondi, les sous-normaux, les débordements,
les infinis et NaN. Une propriété mixte donne **33 succès et 31 échecs attendus
sur 64 essais** : chaque échec est rejoué depuis sa graine et l'exception du
premier essai est vérifiée. Un vecteur de **100 000 éléments** et **100 000 essais
de propriété** contrôlent aussi les parcours longs.

### Spec

Les trois feuilles de `ParallelSpec` sont intentionnelles : `g.3` dans un groupe
séquentiel, `ppp.1` dans un groupe imbriqué ne contenant que du pending et `z.3`
dans un groupe parallèle. Les anciens tests de collecte annonçaient du pending
dans leur nom sans en contenir ; leurs noms décrivent maintenant leur vrai contenu.

Les trois tests actifs de `PendingSpec` vérifient les corps et hooks par test
ignorés, l'arbre des résultats, les chemins et comptes des événements, le bilan,
un rendez-vous entre pairs parallèles via AVar et l'ordre séquentiel imbriqué.
Ce dernier vérifie aussi le mode annoncé par les événements, pour qu'un
ordonnancement parallèle rapide ne puisse satisfaire accidentellement le test.
Le neuvième cas d'intégration, `pending-mixed-contexts`, vérifie le rendu attendu.
Son action séquentielle sépare les groupes adjacents et stabilise l'ordre des
titres du golden.

Vérifications terminées :

- **QuickCheck réussit via le vrai `bin/modtest`**, dans une copie fraîche, avec
  zéro erreur ni avertissement frontend et une sortie Go identique à son oracle JS.
- Sur **234 entrées TAST figées**, les **292 fichiers Go et `go.mod` sont
  identiques octet par octet** entre l'hôte JS, le natif Go séquentiel (workers
  à 1, pipeline désactivé) et le natif Go parallèle (workers à 4, pipeline activé).
  Les trois programmes Go compilent et rendent la même transcription que l'oracle JS.
- Une mutation isolée de sa FFI, lisant les bits Float32 comme un entier non signé,
  est rejetée par le vrai `bin/test` : le diff des graines et fonctions générées
  propage un statut d'échec. La FFI originale est ensuite restaurée.
- **`spec` réussit avec 74 tests et 3 pending intentionnels**, dont les neuf cas
  d'intégration Go ; ses **deux tests d'initialisation** réussissent également.
  La version livrée de `PendingSpec` est incluse dans cette campagne.
- Les **trois nouveaux tests et le golden** passent aussi en JavaScript avec les
  dépendances de référence du package set **77.7.0**. Ce contrôle emploie les
  paquets du registre : les adaptations locales Aff/AVar ont des contrats FFI
  natifs et ne constituent pas un environnement JS exécutable pour cette suite.
- Dans cette copie JS, trois mutations indépendantes de `Test.Spec` sont
  détectées : exécuter le corps de `pending'`, désactiver `parallel`, ou ignorer
  `sequential`. Les trois tests repassent après restauration.
- **64 tests Node des runners réussissent**, zéro échec ni saut, dont le contrôle
  du nettoyage local du nouveau script QuickCheck. La sélection courante compte
  **51 runners** ; les exécutions de bibliothèques de cette passe concernent
  QuickCheck et `spec`.
- Les **4 515 entrées source/artefact initiales** sont auditées : les changements
  se limitent aux fichiers du lot, et les travaux du lot 06 sont préservés.
  **63 liens locaux et ancres**, la syntaxe du script, les manifestes JSON et
  `git diff --check` sont vérifiés.

Les incidents restent consignés : le premier golden supposait un ordre fixe
entre deux groupes parallélisables ; la fixture ajoute une frontière séquentielle
avant de conserver cet ordre. Une relance ultérieure a rencontré `ENOSPC` après
les nouveaux tests actifs ; sa reprise complète via `--resume-failed` réussit
une fois l'espace disponible. Les rapports d'échec et de reprise sont conservés.

Preuves et scripts dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-libraries-FnSmNh/`,
notamment `sources-before.json`, `toolchain.json`, `quickcheck-first.log`,
`quickcheck-validation.py`, `quickcheck-validation.json`, `spec-focused-validation.py`,
`spec-focused-validation.json`, `spec-validated.log`, `runner-tests.log` et
`final-results.json`. Toolchain : Node **24.8.0**, Go **1.27.0**, Spago **1.0.3**,
frontend TAST **0.15.16 development** ; les bibliothèques utilisent le package
set **77.7.0**, les intégrations imbriquées de `spec` le **75.0.0**.

## Plan v2 — lot 08 : installation et validation finale, 7 octobre 2026

Le lot est validé : **100/100 points, 8/8 lots clôturés**. La passe des
**6 et 7 octobre 2026** couvre les installations, les reconstructions et les
campagnes complètes sur les sources inventoriées.

Toolchain local : macOS ARM64, Node **24.8.0**, npm **11.6.0**, Go **1.27.0** et
Spago **1.0.3** pour les applications. Le build JS utilise le toolchain npm
verrouillé, notamment Spago **0.93.45** et esbuild **0.28.1**. Les frontends et
artefacts sont identifiés par SHA-256 dans `toolchain.json`.

Vérifications terminées :

- **Nix vérifié à l'exécution sur aarch64-linux** : `nix flake check`,
  `nix develop` et `nix-instantiate shell.nix` réussissent avec Nix **2.31.2**.
  Le conteneur Docker utilise un magasin Nix en mémoire et les trois fichiers
  Nix copiés en lecture seule ; leurs empreintes et le lockfile restent identiques.
  Le shell fournit Node **24.18.0**, npm **11.16.0**, Go **1.26.4**, Spago **1.0.4**,
  esbuild **0.27.2**, purs-tidy **0.11.1**, PLS **0.18.5**, Python **3.14.6** et
  Git **2.54.0**. Il ne fournit pas `purs`. Les shells Darwin et x86_64-linux
  n'ont pas été exécutés ; les builds du compilateur utilisent le toolchain local
  non-Nix, notamment Go **1.27.0** et le frontend TAST.
- **Installation dans des checkouts neufs** : clonages réels via `bin/setup
  --core`, puis `--all`, avec l'adaptation locale QuickCheck fournie séparément.
  La sélection fraîche ne contenait que **50 des 51 runners** du checkout actif :
  `js-uri` manquait dans `ADDITIONAL_PACKAGES`. Après correction, le vrai setup
  clone ce dépôt et les deux inventaires des **51 runners** coïncident.
- Les bibliothèques clonées sont ensuite placées aux références locales
  inventoriées, avec les modifications non commitées des lots 06/07 superposées.
  Le paquet local `gopurs-argonaut-codecs`, **21 fichiers hors dépôt Git**, est
  également copié et identifié par empreintes pour reproduire le graphe natif.
- **`npm ci` puis bootstrap natif réussis** : **500 modules et 288 603 types**.
  Les deux reconstructions du bundle JS dans le checkout neuf sont identiques.
  Le natif Go reconstruit est identique à celui du checkout actif.
  Le frontend npm frais annonce **`87c821c… DIRTY`**, tandis que le frontend TAST
  local annonce **`3c8fcfd… DIRTY`** ; les deux sont identifiés par leur binaire.
  Le bundle JS frais diffère de l'ancien bundle, dont le frontend local est
  sélectionné par un lien npm vers le frontend TAST. La parité de génération
  vérifiée ci-dessous confirme les mêmes résultats Go.
- **b8x : 2 683 entrées TAST figées et 2 987 fichiers Go plus `go.mod` identiques
  octet par octet**, entre l'ancien bundle JS, le nouveau JS, le natif séquentiel
  et le natif parallèle. Les entrées et compagnons FFI sont vérifiés avant/après
  chaque génération. Ce contrôle concerne la génération de b8x.
- **Archive npm installée offline dans une application vide**, puis compilation
  TAST, génération par le binaire installé et exécution Go réussies avec exactement
  `packaged backend ok`. Le paquet est construit depuis le checkout neuf.
- La campagne Node révèle que `test:cli` sélectionnait implicitement le frontend
  npm, qui n'émet pas le TAST attendu dans une installation fraîche. Le test
  utilise désormais `findTypedCompiler`, comme le bootstrap natif, et respecte
  `GOPURS_PURS` ; **ses 19 contrôles passent après correction**, avec les trois hôtes.
- **435 contrôles Node réussis, zéro échec final ni saut**, sur les 60 fichiers
  de tests des hôtes JS et natif Go. Le fichier Rust optionnel, activé séparément
  par `npm run test:rust`, n'appartient pas à cette campagne. La préparation native
  sous `-race` a rencontré `ENOSPC` à la compilation sur macOS ; le même harnais,
  avec les **615 fichiers Go du bootstrap**, passe dans un conteneur Linux ARM64
  en mémoire avec Go **1.27.0** et Node **25.7.0**. Ses trois tests Go et cinq
  sous-cas valident le différé, la réexécution, le chevauchement, les bornes et
  l'ordre des résultats. Les autres contrôles s'exécutent sous Node **24.8.0**
  sur macOS ARM64.
- **51/51 runners de bibliothèques réussis**, dont le contrôle `assert` en
  compilation seule et les 50 autres avec exécution Go. QuickCheck compare ses
  sorties JS/Go ; `spec` passe **74 tests, 3 pending intentionnels**, ses neuf
  intégrations et ses deux tests d'initialisation. Après deux tentatives limitées
  par `ENOSPC`, la reprise isolée de `spec` avec `--resume-failed` réussit ; le
  rapport final est `campaign-modules/gopurs-tests-ijAzEq/results.json`.
- **400/400 fixtures réussies**, avec snapshots stricts, compilation et exécution
  Go, sans exclusion ni mise à jour des snapshots. La comparaison des inventaires
  vérifie chaque nom entre la sélection active et celle du checkout neuf.
  Les agrégats référencent les **262 rapports de fixtures** et les **53 rapports
  de bibliothèques**, reprises et tentatives initiales comprises.

Les campagnes utilisent des temporaires et caches privés. Deux saturations
disque ont interrompu les fixtures après **119**, puis **171 succès** préservés
dans les rapports individuels. Les bilans agrégés, arrêtés respectivement à
115 et 170 succès, sont reconstruits depuis ces rapports atomiques.
`--resume-failed` reprend réellement `DeepArrayBinder`, puis
`FunctionalDependencies`, `Functions`, `Functions2` et `Generalization1` : les
cinq cas passent avec les mêmes sources et binaires. Les **225 dernières
fixtures** passent ensuite une par une, avec contrôle d'espace avant chaque
cible et réduction du cache Go avant l'écriture du bilan. Aucun succès antérieur
n'est réattribué à des sources différentes ; les rapports initiaux restent intacts.

Les premiers journaux conservent aussi les incidents du harnais de validation :
le premier conteneur Nix n'exportait pas `USER`, empêchant l'activation du PATH ;
la reprise corrige son environnement. La première compilation d'intégration en
mode offline manquait du checkout Git `spec-node` ; son téléchargement permet
la compilation. Le chemin `js-uri` inutilisé dans ce graphe n'empêche pas ce
build : la correction du setup porte sur l'inventaire des runners.

L'audit couvre **4 928 entrées source dans 54 dépôts**, plus les **21 fichiers**
du paquet local `argonaut-codecs`. Les différences de cette passe se limitent
à l'inventaire de setup, au test CLI et à la documentation ; les sources des
bibliothèques, du frontend et de PBO gardent leurs empreintes. Les liens locaux,
les scripts modifiés et `git diff --check` sont vérifiés. Le lien vers le rapport
scratch disparu du 21 septembre est remplacé par sa provenance textuelle dans
`parallel-emission.md`.

Preuves conservées dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-v2-final-DFj3Sa/`,
notamment `sources-before.json`, `installed-remote-heads.json`,
`setup-selection-before.json`, `fresh-checkouts.json`, `unversioned-argonaut-codecs.json`,
`nix-results.json`, `nix-second.log`, `build-results.json`, `b8x-results.json`,
`package-results.json`, `node-cli-fixed.log`, `campaign-fixtures-results.json`,
`campaign-modules-results.json`, `source-audit.json`, `REPRODUCTION.md` et
`final-results.json`. Les workspaces b8x sont archivés
dans `b8x-workspaces.tar.gz`, après vérification des **11 554 fichiers** conservés.
`node-final-results.json` agrège les reprises CLI et `-race` ; le bootstrap est
conservé dans `bootstrap-workspace.tar.gz`, dont les **4 106 fichiers** sont vérifiés.
Les workspaces du lot 07 sont désormais archivés dans son `workspaces.tar.gz` ;
son `archive.json` enregistre les **30 744 fichiers vérifiés** avant suppression
des copies non compressées.

## Consolidation finale — lot 15, 2 octobre 2026

La référence est une arborescence compilateur/runtime figée, contrôlée par
empreintes avant et après reconstruction et génération. Le bootstrap couvre
**500 modules et 287 550 types** ; `bin/gopurs.js` et `bin/gopurs-native`
reconstruits sont identiques octet par octet aux artefacts sauvegardés. Le
manifeste des **111 fichiers compilateur/runtime** est stable. Les modifications
du lot portent sur la documentation, le harnais de tests natifs, la sélection
des fixtures et le nouveau snapshot `DerivingClause`.

Toolchain : Node **24.8.0**, Go **1.27.0**, Spago **1.0.3** et fork `purs`
**0.15.16 development**, commit affiché `3c8fcfd7a3d440bba487fe9fe059284cffc6e908`
avec arbre modifié. L'empreinte du binaire frontend est conservée dans
`toolchain.json` ; PBO est à `0f41544464ec0f42e6cb0dd77b206852813f904f`.
Le build JS de gopurs emploie son toolchain npm verrouillé, distinct de celui
des applications et du bootstrap TAST.

Contrôles terminés :

- `npm run build:native -- --keep-workspace`, puis campagne
  `node --test tools/*.test.mjs` avec `GOPURS_NATIVE_OUTPUT` renseigné :
  **359 tests réussis, zéro échec et zéro saut**, y compris la préparation
  native sous `-race` et les deux contrôles du scanner natif complet.
  `preparation-native.test.mjs` sépare désormais le budget de compilation à
  froid de 900 s du délai Go d'exécution de 30 s et du budget hôte de 60 s ;
- **44 tests Node PBO réussis** : les cinq suites du lot 13 et les trois tests
  `parallel-load`. L'oracle autonome `type-table` et les tests Go du parser FFI
  réussissent également ; ils ne sont pas comptés comme tests Node ;
- pipeline natif : **quatre tests sous `-race`**, puis contrôle des bornes et
  chargement séquentiel/parallèle d'un échantillon figé de **133 modules b8x** ;
- b8x : **2 683 entrées TAST figées**, **2 987 fichiers Go identiques octet par
  octet** à la référence et entre les compilateurs natif séquentiel, natif
  parallèle et JS. Inventaires, runtime, bridges, entrées et `go.mod` sont
  comparés. Une passe complémentaire avec `GOPURS_JOBS=1`, préparation/PBO/
  émission à 1 et `GOPURS_PIPELINE=0` confirme la même identité. Ce contrôle
  porte sur la génération, pas sur l'exécution de b8x ;
- `argonaut-codecs/test/typed-plans.mjs`, dans une copie isolée : compilation
  PureScript, exécution JS, génération native, exécution Go sous `-race` et
  contrôle de propriété du texte réussis ;
- compilation frontend du paquet QuickCheck isolé : **231 modules**, zéro
  erreur ni avertissement. Le paquet n'a pas de runner Go autonome ; ce build
  n'est pas compté comme une suite Go ;
- les neuf exclusions historiques ont été réexécutées. `DerivingClause`
  réintègre le runner après exécution, parité des trois modes et snapshot strict ;
  les huit motifs restants sont détaillés ci-dessus.

La campagne longue utilise une copie de **1 365 fichiers source/configuration
de 52 répertoires frères**, car leurs scripts peuvent supprimer les sorties et
caches des autres paquets. Chaque cible possède son journal et son statut ;
l'échec d'une cible n'empêche pas les autres tentatives.

**Modules : 50/50 runners réussis**, dont `assert` en compilation seule ; les
49 autres exécutent leur programme Go, notamment `node-net`. `spec` réussit
avec **70 tests et 3 pending**, y compris ses huit intégrations imbriquées.
La première tentative avait échoué à initialiser le `node_modules` temporaire,
puis produit sept erreurs Go en cascade ; la suivante a rencontré `ENOSPC`.
Après restauration du répertoire vide ignoré `env-template/node_modules`
présent dans le checkout d'origine et libération d'espace, le runner complet
réussit sans changement de source. Ses intégrations invoquent `npx spago` et
le bundle JS, même quand le runner extérieur utilise le compilateur natif.

**Fixtures : 391/391 réussies**, avec snapshots stricts, compilation et exécution
Go. La sélection initiale de 390 a été exécutée avec préparation/PBO/émission
à huit workers ; `DerivingClause`, ajoutée après son examen, a passé séparément
le runner courant. Aucun snapshot existant n'a changé.

Les incidents de campagne restent dans les journaux : `ENOSPC` a interrompu les
shards après 266 succès ; les 124 cibles sans résultat valide ont été reprises
après purge du cache Go. Cette reprise a donné 120 succès et quatre erreurs
d'extraction Spago (`ENOENT` dans le répertoire temporaire partagé, ou module
de dépendance absent). Les quatre cibles — `OperatorAliasElsewhere`,
`PendingConflictingImports2`, `PolykindBindingGroup1`,
`PolykindInstantiatedInstance` — réussissent en relance séquentielle, dans un
`TMPDIR` dédié, avec les mêmes sources et compilateurs. L'agrégation vérifie
l'inventaire complet de la sélection courante, pas seulement le nombre de succès.

**Nix : non vérifié à l'exécution.** L'exécutable `nix` est absent. La revue de
`flake.nix` et `shell.nix` confirme un environnement de développement, sans
`packages` ni `apps`, qui laisse le frontend TAST et les checkouts frères à
fournir localement. Ni `nix flake check` ni `nix develop` n'ont été exécutés.
Les builds réussis ci-dessus utilisent le toolchain local non-Nix.

Le README, la carte d'architecture et les renvois historiques ont été consolidés.
L'audit des **17 documents et 90 liens locaux** ne relève aucun fichier ni ancre
manquants. Les **1 365 fichiers frères** et les **111 fichiers compilateur/runtime**
conservent leurs empreintes de référence ; `git diff --check` est propre.

Les preuves, références, manifests et journaux sont conservés sous
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-final-consolidation-xy_5fmmn/`.
Le bootstrap natif est dans `gopurs-native-build-Hezybl/` sous le même parent ;
le workspace `typed-plans` est `gopurs-tests-Qhgufu/`. `final-results.json`
agrège les résultats, et `final.diff` conserve le diff complet du lot par rapport
à l'arborescence figée, y compris les changements déjà repris par le bot de commit.

Le lot 15 est validé : **100/100 points, 15/15 lots** du plan v1 clôturés.
Cette clôture couvre la maintenabilité et les validations décrites : les huit
exclusions, les trois `pending` de `spec`, l'absence de suite Go autonome de
QuickCheck et l'exécution Nix non vérifiée restent des limites explicites.
Aucun gain de performance n'est déduit de ces validations.

## Runtime et FFI du compilateur — lot 14, 2 octobre 2026

`Printer.Builder` possède maintenant le buffer mutable, les opérations FFI
JS/Go et le confinement par `withOut`. `Printer` conserve le rendu et
l'échappement. La FFI `GoCode` sépare l'adaptation des valeurs boxées du scanner
`scanReferencedImports` ; `GoImports` explicite l'emprunt des entrées et la
propriété du résultat. Les imports runtime de la FFI native sont explicites.

La revue des autres compagnons confirme leur responsabilité unique : cache de
noms, transport du parser, constante runtime, horloge/configuration du profil
et sortie du processus. Le runtime canonique conserve ses octets et son ABI
partagée avec les bibliothèques ; ses familles, alias, copies, captures et
retenues sont décrits dans [runtime-ffi-contracts.md](runtime-ffi-contracts.md).

Vérifications effectuées :

- reconstruction JS et native via `npm run build:native -- --keep-workspace` ;
  bootstrap vérifié sur **500 modules et 287 550 types**, zéro avertissement JS
  et les mêmes quatre avertissements TAST de PBO ;
- **42 tests ciblés avant et après**, sans saut : `native-ffi`,
  `runtime-contracts`, `embed-runtime`, `closure-lifetime`, `apply-arity`,
  `function-data`, `value-array-unboxing`, `go-imports`, `ffi-errors`,
  `ffi-bridge` et `ffi-generics`. Le premier état du harnais natif échouait sur
  l'import runtime absent de sa copie autonome de `Printer.go` ; ce harnais a
  été réparé avant d'établir le passage vert de référence ;
- ces suites exécutent **13 tests Go du runtime**, normalement et sous `-race`,
  et **8 tests Go des compagnons FFI sous `-race`**. Chaînes packées et WTF-8,
  sous-chaînes, buffers d'impression, graphes de conteneurs après sortie du
  créateur, GC forcé, caches synchronisés, copies de maps et travail retenu
  sont couverts, avec les matrices de closures/applications déjà présentes ;
- **2 contrats de scan natif complet avant et après**, sans saut, à partir des
  workspaces bootstrap. Le lancement après extraction a d'abord dépassé le
  budget global de 60 s de `go run`. Le harnais distingue désormais compilation
  (900 s maximum) et exécution (60 s), en gardant la pile bornée à 1 Mio ; le
  contrôle complet réussit ;
- **8 fixtures**, snapshots stricts puis compilation et exécution Go avec
  `GOPURS_PBO_JOBS=8 GOPURS_PREPARE_JOBS=8` : `FFIIntegerReturns`,
  `CurriedLambdas`, `NativeArrayReboxing`, `ObjectUpdate2`, `ConstructorReuse`,
  `JsonRecordPlan`, `FFIConstraintWorkaround`, `StaticDictionary`.
  `StringEdgeCases` et `StringEscapes`, demandées initialement, restent filtrées
  par les exclusions du runner et ne sont pas comptées comme validées. Les
  surrogates sont exercés ici par les contrats FFI/runtime ;
- b8x : référence produite avec le compilateur vivant sauvegardé avant la passe,
  sur **2 683 entrées TAST figées**, puis **2 987 fichiers Go identiques octet
  par octet** en natif parallèle, natif séquentiel et JS. Inventaires, runtime,
  bridges, entrées exécutables et `go.mod` sont inclus. Le manifeste des
  **111 fichiers sources et artefacts compilateur/runtime** reste stable.

Le lot 14 est validé : **95/100 points, 14 lots sur 15** clôturés. Le contrôle
b8x porte sur la génération ; les tests natifs et fixtures vérifient l'exécution.
`git diff --check` est propre. Les preuves locales, la référence et le diff
complet de cette passe sont conservés dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-runtime-ffi-cleanup-6052azh5/`.
Le workspace bootstrap est conservé sous le même parent dans
`gopurs-native-build-TLO4c1/`. Le lot 15, consolidation des campagnes et de la
documentation avec statut explicite de Nix, constituait alors la prochaine passe.

## Frontière TAST/PBO — lot 13, 2 octobre 2026

`Monomorphization` sépare la reconnaissance des wrappers dans `ForeignForwarders`
et nomme les exclusions tardives avec `SpecializationBarriers`. L'invalidation
source et l'exception intrinsèque précèdent la collecte ; l'admission par les
types d'origine suit le point fixe. La provenance des tables est explicitée
dans `Driver.Prepare`, `GlobalTypes`, `CodegenState` et l'architecture.

La revue confirme que `Preparation` possède déjà une seule responsabilité dans
un petit module : exécuter un tour différé en conservant l'ordre. PBO possède
le cache, la fusion et la barrière entre tours. Le contrat est précisé au point
d'appel et vérifié sur les deux runtimes ; les fonctions de métadonnées gardent
leurs modules et leurs politiques distinctes.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **499 modules et
  287 541 types**. Le build JS est sans avertissement ; les quatre avertissements
  TAST de PBO sont identiques à la passe précédente ;
- **46 tests gopurs avant et après**, dont treize nouveaux contrats dans
  `monomorphization`, `global-types` et `preparation`. Les suites voisines sont
  `foreign-forwarders`, `native-record-args`, `representation-contract` et
  `rebox-metadata` ;
- **41 tests PBO avant et après**, sans saut : `monomorphize-transitive` (12),
  `transitive-parallel` (5), `monomorphize-callsite` (3), `monomorphize-cache` (9)
  et `source-usage` (12), exécutés sur le même `output` compilé que gopurs ;
- tests natifs de préparation **avant et après sous `-race`** : trois tests
  principaux et cinq sous-cas, avec tous les marqueurs PASS attendus. Le premier
  lancement du wrapper Node a dépassé son budget global de 120 s ; la relance Go
  directe autorise une compilation froide plus longue et conserve le délai
  d'exécution de 30 s. Différé, réexécution, unicité, chevauchement, borne à huit
  et restitution ordonnée sont validés ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `NativeRecordWorkers`, `StaticDictionary`, `TypeClassMemberOrderChange`,
  `VisibleTypeApplications`, `Rank2Types`, `FFIConstraintWorkaround`,
  `ConstructorReuse`, `MutRec` ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **107 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

Le lot 13 est validé : **90/100 points, 13 lots sur 15** clôturés. Le contrôle
b8x porte sur la génération ; les tests natifs et fixtures vérifient l'exécution
Go. `git diff --check` est propre. Le runtime et les FFI JS/Go du compilateur,
au lot 14, constituaient alors la prochaine passe.

## Analyses spécialisées — lot 12, passe DecoderSchemas, 2 octobre 2026

`DecoderSchemas` conserve l'orchestration dans une façade de 85 lignes, contre
595 auparavant. `Source` possède la résolution des dictionnaires, `Admission`
les critères de spécialisation, `Programs` la preuve des corps personnalisés,
`Types` les capacités du schéma et `Workers` l'émission DOM/texte. Les sources de
constructeurs restent compilées par le générateur ordinaire et partagées entre
les deux modes. L'argument de préfixe inutilisé de l'émetteur a été retiré ; ses
choix de mode sont désormais nommés.

Vérifications effectuées :

- reconstruction JS et native ; bootstrap TAST vérifié sur **498 modules et
  287 336 types**. Le build JS est sans avertissement ; le build TAST ne conserve
  que les quatre avertissements préexistants de PBO. L'extraction de la
  construction des records supprime le masquage de nom de `DecoderSchemas` ;
- **105 tests ciblés avant et après** : `decoder-schemas`, `borrowed-objects`,
  `closed-dictionaries`, `owned-trees`, `native-record-args` et
  `native-record-workers`. Les treize nouveaux contrats couvrent les gardes ABI,
  les formes exactes, les budgets, les barrières, la réservation des noms, les
  lectures prouvées, les erreurs, les branches et les captures de constructeurs.
  Les deux programmes Go de schémas et les tests du helper réel `CompiledSchema`
  passent avec le détecteur de courses ;
- **8 fixtures**, snapshots stricts, compilation et exécution Go avec 8 workers :
  `JsonRecordPlan`, `StaticDictionary`, `DuplicateProperties`,
  `NativeRecordReturns`, `ObjectTraverseEither`, `FunctionScope`, `MutRec`,
  `OneConstructor` ;
- programme Argonaut `test/typed-plans.purs` dans un workspace isolé : **263
  entrées TAST figées**, **330 fichiers Go et `go.mod` identiques** entre le
  compilateur de référence et le nouveau natif, puis compilation et exécution
  réussies. Ce contrôle exerce les plans imbriqués et les workers texte ; les
  corps personnalisés prouvés sont couverts par les contrats d'IR ci-dessus ;
- b8x : référence régénérée avec le compilateur précédent sur **2 683 entrées
  TAST figées**, puis **2 987 fichiers Go identiques octet par octet** en natif
  parallèle, natif séquentiel et JS, sans ajout ni suppression. Runtime, bridges
  FFI, entrées exécutables et `go.mod` sont inclus. Le manifeste des **106 fichiers
  sources et artefacts compilateur/runtime** est stable pendant ces comparaisons.

La validation des cinq analyses du lot 12 a porté le score à **85/100**,
avec **12 lots sur 15** clôturés. Le contrôle b8x porte sur la génération, les
tests et fixtures sur la compilation et l'exécution Go ; `git diff --check` est
propre. La frontière TAST/PBO du lot 13 constituait alors la prochaine passe.

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

La passe `NativeRecordArgs` a été validée avec `DecoderSchemas` encore à revoir
dans le lot 12. Le score est resté à **75/100**, avec onze lots clôturés. Le contrôle
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

À l'issue de cette vague historique restaient ouverts : les snapshots TCO/TCOMutRec et le contrôle
ciblé de fusion du bilan de la première vague, la validation étendue de
la [règle ArrayRoundtrip](array-roundtrip.md) prévue en 3.4 de ce chantier, et
une campagne complète `passing` / modules frères. La consolidation du lot 15
ci-dessus donne l'état actuel ; les constats de cette section sont datés. L'ancienne
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
Le [todo actuel](../todo.md) décrit le plan v2 de fiabilité et reproductibilité.

Le contrôle `ArrayRoundtrip -c --keep-workspace` du 14 septembre confirme les
28 assertions existantes et le snapshot inchangé. Le résultat incorrect du
singleton pair consigné le 9 septembre ne se reproduit plus : le résultat est
`8`. Cette vérification n'a nécessité aucune modification du compilateur ;
elle ne désigne pas la cause ni la correction de l'ancien échec.

## Référence du lot 1 de maintenance — 14 septembre 2026

Cette référence historique accompagne la [carte des dépôts](architecture.md#carte-des-dépôts-et-des-consommateurs).
Elle vérifiait les points d'entrée et les comportements ci-dessous ; la revue
interne et la campagne des bibliothèques étaient encore à faire à cette date.

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

Cette section conserve le constat du **14 septembre 2026**. Les numéros de lots
renvoient à l'ancien journal de maintenance ; la campagne du lot 15 du plan v1
donne le bilan consolidé des fixtures et bibliothèques au 2 octobre 2026.

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
restaient ouverts à cette date. Le lot 1 n'établit ni un build intégral de chaque bibliothèque,
ni une validation réseau/FS/Aff, ni un build sans caches de dépendances.

## Lot 2 — installation et configurations

Cette section décrit la vague de maintenance du **14 septembre 2026**, antérieure
au plan v1 de maintenabilité.

Le [guide local](../README.md#develop-one-library-locally) décrit désormais le
parcours de chaque bibliothèque. Le lot 2 a examiné les **404 fichiers de
configuration** inventoriés : **91 modifiés, 313 conservés**, ainsi que
`bin/setup` et les imports/dépendances Go. Les fichiers Bower, Dhall, CI,
formatage, lint et règles d'exclusion gardent leurs rôles existants. Les
configurations des exemples et le template d'intégration de `spec` sont aussi
identifiés ; `SPEC_REPO_PATH` y est remplacé par le runner, et leur revue de
tests était attribuée au lot 14 de cette ancienne vague.

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
