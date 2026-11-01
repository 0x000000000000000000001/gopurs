# Tests et portée des validations

Les commandes de cette page partent de la racine de gopurs, avec Go, Node,
Spago et le `purs` TAST sur `PATH`. Voir [l'installation](../README.md).
Le runner utilise `bin/gopurs`, donc le binaire natif existant par défaut ;
`GOPURS_JS=1` sélectionne le bundle JS et `GOPURS_RUST=1` l'hôte Rust.
`-c` reconstruit une fois l'hôte sélectionné. Après une modification du
compilateur, `npm run build:native -- --keep-workspace`
reconstruit les deux versions et conserve le workspace pour les tests natifs.

## Choisir un contrôle court

| Famille modifiée | Commande ciblée |
| --- | --- |
| Pilote et durée de vie des tâches | `node --test tools/emission.test.mjs`, après le build ; tests natifs ci-dessous |
| Statut de sortie et diagnostics CLI | `npm run test:cli`, après `npm run build:native` |
| Bootstrap natif et gestion des processus | `node --test tools/build-native.test.mjs tools/test-runner.test.mjs` |
| Frontière TAST/PBO, types globaux et spécialisation | `node --test tools/monomorphization.test.mjs tools/global-types.test.mjs tools/preparation.test.mjs tools/foreign-forwarders.test.mjs`, après le build ; préparation native ci-dessous |
| Annotations et représentations natives | `node --test tools/native-record-workers.test.mjs tools/boxed-record-arguments.test.mjs tools/native-sum-results.test.mjs`, après le build |
| Références Go possédées et portées FFI | `npm run build:ffi`, puis `node --test tools/go-source.test.mjs tools/module-codec.test.mjs tools/emitter-worker-keys.test.mjs`, après le build ; `go test -race -count=1 ./...` dans `tools/ffi-gen` et `tools/build-cache` |
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
| Évaluation des arguments et fabriques de callbacks Array | `node --test tools/array-intrinsic-evaluation.test.mjs tools/array-intrinsics.test.mjs`, après le build ; [validation du 24 octobre](build-cold-20261024.md#validation-fonctionnelle) |
| Propriété des arbres et workers consommants | `node --test tools/owned-trees.test.mjs`, après le build ; `./bin/test OwnedTrees RBTree ConstructorReuse` |
| Mise en cache des dictionnaires fermés | `node --test tools/closed-dictionaries.test.mjs`, après le build ; `./bin/test StaticDictionary JsonRecordPlan` |
| Emprunt d'objets et composition des décodeurs | `node --test tools/borrowed-objects.test.mjs tools/closed-dictionaries.test.mjs tools/decoder-schemas.test.mjs`, après le build ; `./bin/test JsonRecordPlan` |
| Contrat du parser Go | `go test ./...` depuis `tools/ffi-gen` |
| Parser WASM et erreurs FFI | `npm run test:ffi`, après `npm run build` |
| Sélection, isolation et erreurs du runner | `npm run test:runner` |
| Cache des exécutables et verrou du runner b8x | Depuis `../../b8x` : `B8X_GO_RUN_NATIVE_TEST=1 node --test test/go-run.mjs` ; [contrats et commande conteneur](../../../b8x/test/go-run.md) |
| Contrat persistant du cache gopurs, intégrité et publication | `npm run test:cache` ; [protocole, propriété et invalidation](build-cache.md) |
| Hit avant décodage, invalidations et publication du pilote | `npm run test:cache-driver`, après le bootstrap ; `GOPURS_TEST_RUST=1` ajoute l'hôte Rust |
| Corps, layouts, demandes de spécialisation et renommages sur cache engagé | `npm run test:cache-sources`, après le bootstrap ; `GOPURS_TEST_RUST=1` ajoute l'hôte Rust |

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

La [campagne du 24 octobre](build-cold-20261024.md) ajoute une validation ciblée
de l'encodage Wire et de l'évaluation des intrinsèques Array : 39 tests de
cache/préparation, 4 tests Array, 24 programmes Go compilés et exécutés, et le
gate driver-cache sur les hôtes JS et natif parallèle. Les quatre éditions
successives de b8x sont comparées à une génération complète. Cette campagne
ne remplace pas une exécution intégrale du corpus ni les validations de l'hôte
Rust ou du mode natif séquentiel.

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
FFI invalide et écriture de l'entrée exécutable. Depuis le lot 03, l'absence de
sortie est détectée pendant `cache lookup + freeze` ; les conflits des destinations
publiques sont détectés pendant `cache publish`, après la génération privée.
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
./bin/test -d FFIIntegerReturns --keep-workspace
```

Sans cible, `bin/test` sélectionne toutes les fixtures non exclues de
`tests/passing`. Les noms explicites gardent leur ordre ; la reprise inclut sa
cible. `--skip-before=NAME` et `skip_before=NAME` restent compatibles. Une cible
inconnue, une option inconnue ou une sélection entièrement exclue échoue avant
le build. `--list` ne compile rien et ne crée pas de workspace.

La vérification de snapshots est le mode par défaut dans les deux profils.
`bin/test` sélectionne Release ; `-d` ou `--dev` sélectionne Dev. Chaque profil
compare son Go généré à ses propres références, puis compile et exécute la
fixture. Les sources et les résultats fonctionnels attendus restent communs.
Un fichier absent ou différent échoue : examiner le diff et la phase responsable
avant de le remplacer.

| Profil | Layout historique (`GOPURS_GO_PACKAGES=0`) | Layout packages (`GOPURS_GO_PACKAGES=1`) |
| --- | --- | --- |
| Release (défaut) | `tests/passing-snapshots` | `tests/passing-package-snapshots` |
| Dev (`-d`, `--dev`) | `tests/passing-snapshots-dev` | `tests/passing-package-snapshots-dev` |

Le runner fixe `GOPURS_GO_PACKAGES=0` par défaut pour utiliser les références
complètes du layout historique. `GOPURS_GO_PACKAGES=1` explicite sélectionne
le layout packages et ses références. Une invocation vérifie seulement le
layout et le profil sélectionnés, pas les deux layouts à la fois.

Les snapshots du layout historique sont nommés `<Fixture>.go`, avec
`<Fixture>_ffi.go` en plus si la source déclare `-- @snapshot-ffi`.
Une mise à jour intentionnelle s'effectue ainsi :

```bash
./bin/test FFIIntegerReturns --update-snapshots
./bin/test -d FFIIntegerReturns --update-snapshots
# Équivalent historique : UPDATE_SNAPSHOTS=1 ./bin/test FFIIntegerReturns
# Profil Dev : UPDATE_SNAPSHOTS=1 ./bin/test -d FFIIntegerReturns
```

Seuls les snapshots du profil sélectionné sont mis à jour, après compilation
et exécution Go réussies de la fixture concernée. Par défaut, la campagne
s'arrête au premier échec ; les mises à jour
des fixtures déjà réussies restent écrites. Le contrôle d'exécution conserve
le contrat historique : statut zéro et absence de `Fail` dans la sortie.

Pour valider toute la suite dans les deux profils, lancer deux campagnes :

```bash
./bin/test --all --keep-going
./bin/test -d --all --keep-going
```

### Validation des profils avant intégration de C2 — 24 octobre 2026

Sur l'hôte Go natif, les 401 fixtures non exclues ont été compilées et
exécutées dans chacun des profils Release et Dev. Chaque profil possède
405 références dans le layout historique, incluant les fichiers FFI, et
9 références validées dans le layout packages. Les 93 tests du runner passent.

Les références Release anciennes ont été rafraîchies après succès d'exécution
et revue des différences : 178 fichiers historiques et 3 fichiers packages.
Les replays stricts dans des workspaces neufs passent pour `Conditional`,
`3329`, `NativeTraverseCallback` et `TCO`, ainsi que `JsonRecordPlan` en
packages. Des oracles JavaScript complètent les contrôles de `Conditional`
(8 cas booléens), `3329` (12 assertions d'injection/projection), `TCO` et des
trois fixtures packages modifiées. Le binaire natif est resté inchangé pendant
cette validation ; aucune correction du compilateur n'a été nécessaire.

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
`bin/modtest` ; son log par cible est `logs/test.log`. Les sorties stdout/stderr
de chaque cible sont diffusées pendant son exécution et conservées dans ce log,
sans être rejouées à la fin. Avec `-d` ou `--dev`, le lanceur privé annonce aussi
le mode Dev à chaque appel du compilateur par une bibliothèque.

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

`tools/dev-loop-state.test.mjs` protège le coût des callbacks génériques de
`tailRec` : une liste croît dans un record pendant 100 000 étapes en Dev, sans
spécialisation. Le test mesure les allocations à 1 000 puis 2 000 étapes avant
le grand cas, vérifie longueur et somme, puis compare le programme complet à
JavaScript. Une croissance quadratique échoue sur les petits cas, sans attendre
les 100 000 étapes. Les deux hôtes JS et natif sont couverts.

`tools/recursive-function-result.test.mjs` vérifie qu'une recherche récursive
peut renvoyer une fonction avant de recevoir ses derniers arguments. Les cas
de comparaison, d'égalité et d'affichage s'exécutent en Release et en Dev et
doivent produire la même sortie que JavaScript. Le type de retour du worker
conserve les arguments non consommés par ses paramètres réels.

## Plan v3 — lot 00 : référence du cycle b8x, 10 octobre 2026

Le lot 00 est validé : **3/7 lots du plan v3**, après les lots 01 et 02. La
[référence détaillée](build-baseline.md) conserve les scénarios, la ventilation,
les outils et les limites des mesures sur un corpus figé de **2 704 TAST /
755 005 types**. Au départ, b8x sélectionne JS ; la campagne construit une copie
avec sa configuration Go et ses dépendances locales.

### Résultats

- **Relance identique : 36,902 s dans le processus gopurs**, médiane de trois
  passages ; les **3 010 fichiers Go sont réécrits**, alors que leur contenu est
  identique. Les médianes internes sont 1,298 s pour chargement/tri, 10,710 s pour
  préparation/spécialisation et 24,486 s pour optimisation/émission. Les lots
  d'émission se chevauchent avec PBO et ne s'ajoutent pas au total.
- **Chaîne Spago/backend → tidy → Docker/runner : 40,808 s de médiane**,
  étendue 38,892–42,134 s, jusqu'à la frontière `exec` de `Test.Main`. Les trois
  passages ne compilent ni ne lient en Go. La suite métier n'entre pas dans
  cette horloge ; les exécutions effectives sont contrôlées par les oracles.
- **Runner seul à chaud : 0,770 s de médiane**, même exécutable à l'octet près,
  sans compilation ni lien. Acquisition du verrou affichée à 0,000 s dans toutes
  les invocations. Les observations avant/après retrouvent des services et
  healthchecks JS ; le verrou de mesure est privé.
- **Feuille / partagée / FFI : 1 / 5 / 1 Go modifiés** par rapport à la référence.
  Chaque nouvelle variante demande **212–220 s** pour compiler le package commun
  `purescript`, avec **2 CPU / `-p 1` / `GOMEMLIMIT=12GiB`**. Un échantillon par
  première construction, dépendances déjà en cache ; pics RSS observés de
  14,5 à 14,9 GiB. Les limites du runner sont celles du parcours habituel.
- La modification partagée recompile **1 375 modules frontend** en **20,685 s**
  de médiane, mais un seul TAST change de contenu. Son Go modifié comprend
  `Control.Semigroupoid`, ce qui confirme la portée des demandes de spécialisation
  provenant des appelants. La modification FFI ne change aucun TAST.

### Vérifications et reproductibilité

- Les treize mesures frontend ont été corrigées après revue du harnais : retirer
  le backend Spago activait JS/sourcemaps. Un backend neutre maintient les options
  réelles `corefn,docs` ; tous les TAST sont byte-identiques aux entrées des
  mesures gopurs/Go déjà réalisées. Les relevés préliminaires restent archivés.
- **Quatre corpus, trois modes** : JS, natif séquentiel et natif parallèle
  produisent les mêmes **3 010 Go par corpus**, runtime, bridges et entrées
  compris, ainsi que le même `go.mod` minimal. Inventaires et octets sont
  également identiques entre les répétitions de chaque scénario.
- **Cinq oracles Go réussis** sous Linux : référence, identique et trois
  modifications, avec les résultats attendus de `caseToCamel`, `RemoveAccents`
  et `Ping`. Les constructions de `Test.Main` réussissent aussi.
- **Trois snapshots revus, six replays stricts réussis** : `FFIIntegerReturns`
  et `NativeRecordWorkers` dans les trois modes, avec compilation/exécution Go.
  Les snapshots et les octets des compilateurs correspondent à la référence.
- Audit final de **4 318 fichiers inventoriés**, **8 112 observations de
  sources/compagnons**, TAST, sorties et binaires publiés ; contrôles de syntaxe,
  formatage de l'oracle et `git diff --check` réussis. Les mutations privées sont
  restaurées et les deux conteneurs dédiés sont supprimés.

La prochaine étape est le raccordement du chemin sans changement au lot 03.
La compilation monolithique Go constitue un coût distinct ; sa granularité reste
une suite mesurée du plan, après la livraison initiale du cache.

Preuves conservées dans
`<tmp>/b8x-incremental-lot00-srr89dtf/` :
`README.md`, `summary.json`, `backend-results.json`, `frontend-corefn-results.json`,
`cycle-results.json`, graphes d'actions de `measurements-go/`, `parity-results.json`,
`snapshot-review.json`, `strict-replays.json` et `final-audit.json`. Les sources,
TAST, Go, compilateurs, ressources FFI et exécutables sont conservés ; le README
fournit le rejeu de parité sur une copie neuve du corpus figé.

## Plan v3 — lot 01 : réutilisation des exécutables b8x, 10 octobre 2026

Le lot 01 est validé : **1/7 lots du plan v3**. `b8x/bin/_go-run` présente à Go
une copie privée, vérifiée par SHA-256, du dernier exécutable publié pour ce
point d'entrée. `go build` valide toujours les sources, dépendances et options ;
il peut ainsi éviter le lien lorsque son build ID est à jour. L'indice `latest`
est publié atomiquement après succès. Les publications valides sont immuables,
et chaque invocation conserve son propre chemin après libération du verrou.

Les diagnostics distinguent l'acquisition du verrou, la validation/build et les
limites `GOMAXPROCS`, `-p` et `GOMEMLIMIT`. Le message d'attente n'apparaît que
sur contention. Les [contrats et commandes de reproduction du runner](../../../b8x/test/go-run.md)
décrivent le cache et les contrôles intégrés.

### Contrôles de correction

- **10/10 contrôles réussis** avec Go **1.26.8**, Node **25.7.0**, Bash **5.2.15**
  et util-linux dans le conteneur API Linux ARM64. Les neuf contrôles de contrats
  réussissent aussi sur macOS avec ses shims `flock`/`setsid`, puis le contrôle
  natif est exécuté séparément avec Go **1.27.0** et Node **24.8.0**.
- Les sept contrôles historiques passent avant modification dans le conteneur.
  Le nouveau reproducteur natif échoue contre l'ancien runner au deuxième
  lancement : il observe encore un lien. Le même reproducteur passe après.
- Sources, dépendance modifiée à date conservée, `go.mod` et options de lien
  sont pris en compte. Un build échoué n'exécute pas l'ancien binaire et conserve
  la publication précédente. Les indices absents/invalides et les exécutables
  absents/altérés sont couverts, y compris une altération conservant le build ID.
- Les tests vérifient la copie privée, les artefacts propres à chaque invocation,
  les applications concurrentes, les annulations du build et du demandeur de
  verrou, les descendants, arguments, statuts, signaux et environnements.
  Le test d'annulation attend désormais une contention effective plutôt qu'un
  délai fixe ; l'ancienne synchronisation avait produit `null !== 143` sur macOS.
- **Trois exécutions réelles de `Ping`** renvoient exactement `pong\n`, statut 0.
  L'entrée Go `Ping` conservée est liée aux packages b8x figés. La dernière
  relance prend **0,405 s**, sans compilation ni lien ; il s'agit d'un contrôle
  fonctionnel, pas d'une seconde campagne statistique.
- À la reprise du 10 octobre, les empreintes du runner et de son harnais sont
  identiques aux copies validées. `bash -n` et `git diff --check` sont vérifiés.

### Mesure à cache chaud

Cinq paires alternées utilisent les mêmes **2 984 fichiers Go, `go.mod` et
`go.sum`**, le même point d'entrée `Test.Main`, la même image API et les mêmes
caches de packages/modules. Les sorties proviennent d'une génération isolée
de **2 688 TAST**. Le cache des exécutables et son verrou sont privés, sur le
même montage partagé que le cache habituel. Paramètres : **Go 1.26.8 Linux ARM64,
`GOMAXPROCS=2`, `-p 1`, `GOMEMLIMIT=12GiB`, `GOWORK=off`, `GOFLAGS=-x`**.

Le chronométrage couvre le runner jusqu'au transfert à l'exécutable : validation,
lien éventuel, copies, empreintes, publication et libération du verrou. Un
harnais externe `BASH_ENV` s'arrête au dernier diagnostic, juste avant `exec`.
`Test.Main` appelle directement `runSpecAndGetResults` et ne traite pas `--help` :
la suite métier n'entre donc pas dans cette mesure. L'exécution réelle est
éprouvée par les tests natifs et par `Ping`.

| Mesure par relance chaude | Avant | Après |
| --- | ---: | ---: |
| Médiane jusqu'à `exec` | 1,812 s | 0,745 s |
| Minimum | 1,599 s | 0,701 s |
| Maximum, tous échantillons conservés | 1,938 s | 11,008 s |
| Compilations Go | 0 | 0 |
| Liens | 1 | 0 |

La réduction médiane est de **58,9 %**, soit **1,067 s** économisée. Les cinq
relances optimisées passent de **0,270 à 0,290 s dans `go build`**. L'acquisition
du verrou y est affichée à `0.000 s`, avec une résolution de 10 ms ; cette campagne
isolée ne mesure pas la contention des services/healthchecks en production.

L'échantillon après à **11,008 s** reste inclus. Son appel Go dure 0,280 s et son
acquisition de verrou est affichée à zéro ; le retard se situe hors de ces deux
temps mesurés, sans cause identifiée. Les quatre autres échantillons après vont
de 0,701 à 0,823 s. Le résultat établit un gain médian, pas une borne de latence.

Les dix invocations publient le **même exécutable de 74 672 311 octets**, SHA-256
`50880fc828e2931d101281b4e0f55b09f57398e0fdc7df6db8e42b5807face06`.
L'inventaire et les octets des 2 986 fichiers Go/module restent identiques ;
l'audit final contrôle également les 2 688 TAST et 3 098 sources/compagnons.

### Préparation et limites de portée

La sortie Go active au début de la campagne échouait déjà avec l'ancien runner
(`Constructor_Data_Maybe_Just` absent). Ses CoreFn annonçaient `0.15.15` sans les
tables TAST. Le corpus de mesure est donc reconstruit dans un workspace isolé,
avec les overrides de bibliothèques de b8x et le fork typé. L'override inutilisé
`gopurs-math`, qui pointait vers un checkout absent, y est omis.

La première tentative est interrompue par un redémarrage du conteneur API
(statut 137). La campagne reprend dans un conteneur dédié de la même image.
Le build Go préparatoire à `12GiB` dépasse 900 s et est annulé par le harnais.
Une préparation ponctuelle à `18GiB` réussit la compilation et le lien, puis
l'application échoue sur sa configuration PostgreSQL absente ; l'invocation
complète dure 889,529 s. L'inspection du runner de tests explique alors pourquoi
`--help` n'évitait pas l'exécution, et conduit au chronométrage avant `exec`.
Ces préparations sont exclues des médianes ; les deux variantes mesurées utilisent
les réglages habituels `2 / 1 / 12GiB`.

Cette campagne établit le bénéfice des **relances Go à chaud**. La ventilation
du cycle PureScript/gopurs et l'incrémentalité du backend restent les objets
des lots 00 et 02–06 du plan v3.

Preuves conservées dans
`<tmp>/b8x-go-run-lot01-aa6h2e9b/` :
`before/`, `after/`, les logs rouge/vert et natifs, `comparison.json`, les traces
`measurements/handoff-*`, `measure-runner.mjs`, `compare-runners.mjs`,
`handoff-only.sh`, `ping-results.json`, `ping-entry-provenance.json`,
`published-binaries/`, `typed-final-inputs.json`, `normalized-module/`,
`final-audit.json` et `README.md`. Les échecs et reprises préparatoires y sont
également conservés ; les fichiers générés et les dépendances sont figés dans
`typed-workspace-v2/`.

## Plan v3 — lot 02 : contrat du cache persistant, 10 octobre 2026

Le lot 02 est validé : **2/7 lots du plan v3**. Le
[contrat versionné](build-cache.md) précise les clés, la recette complète à
fournir, les artefacts par module, le format de valeurs portable, l'identité du
compilateur, les sorties possédées et les règles de publication/invalidation.

Le stockage Go fournit `snapshot`, `lookup` et `publish`, avec SHA-256,
publication atomique des fichiers et du manifeste, verrou noyau et journal de
propriété pour les reprises. Le même code est utilisé par le bridge JS et lié
au bridge Go. `go.mod` est créé seulement en son absence ; les dépendances
existantes et `go.sum` restent la propriété du parcours Go.

### Vérifications effectuées

- `npm run test:cache` réussit : tests du stockage sous race detector et échange
  de manifestes/publications entre la FFI JS et la FFI Go compilée contre le
  package réellement préparé pour le bootstrap.
- **38 cas de stockage**, plus le harnais de sous-processus, réussissent sous
  `go test -race -count=1 -v -timeout 90s ./...`, avec Go **1.27.0 macOS ARM64**
  et Go **1.26.8 Linux ARM64**. La passe Linux utilise un conteneur isolé de la
  même image que l'API, source montée en lecture seule, réseau désactivé,
  `GOMAXPROCS=2`, deux CPU et 4 GiB ; il est supprimé après le test.
- Couverture : edits à date conservée, ajout/retrait/renommage de TAST, FFI
  présente/absente, racines de résolution, directives/options, compilateur et
  ressources, liens symboliques, versions et corruption, sorties abîmées,
  propriété et suppression, écritures identiques, `go.mod`/`go.sum`, requêtes
  périmées, six publications concurrentes, verrou d'un processus tué et
  récupération d'une première publication interrompue.
- La revue a identifié le cas du journal absent avec manifeste précédent valide.
  Le reproducteur conservé laisse une sortie obsolète avec le premier auxiliaire
  (**rouge**), puis passe après récupération de la propriété depuis ce manifeste
  (**vert**). Les deux auxiliaires et les logs sont conservés.
- `npm run build:native -- --keep-workspace` reconstruit le bundle JS et le natif
  final depuis **502 TAST / 288 955 types**. Le package de cache est copié par
  `prepare-native-output.mjs` puis compilé avec le compilateur natif.
- Les anciens compilateurs sont conservés. `FFIIntegerReturns` et
  `NativeRecordWorkers` partagent les **mêmes 121 TAST / 27 950 types cumulés**
  entre avant/après et entre **JS, natif séquentiel et natif parallèle**.
  Chaque mode produit exactement les mêmes **174 fichiers Go et deux `go.mod`**,
  runtime, FFI et points d'entrée compris. Chaque programme Go compile, s'exécute
  avec le statut 0 et affiche `Done`.
- Les **trois snapshots** versionnés correspondent aux sorties formatées lors
  de la revue. **Six replays stricts** réussissent ensuite, deux fixtures dans
  chacun des trois modes. Aucune actualisation de snapshot n'a été nécessaire.
- `go vet ./...` et `git diff --check` réussissent.

### Portée et suite

Le pilote effectue encore un build complet. Le producteur de recette, les codecs
métier PBO et les raccordements de réutilisation relèvent des lots 03–04 ; les
tests du format de valeurs valident la grammaire et la conservation des octets,
pas encore un aller-retour de vrais `BackendModule`. Le contrat distingue
explicitement ces obligations du stockage déjà testé.

La validation de ce lot n'est pas une mesure de performances. Le lot 00 doit
établir les références du backend et du cycle réel avant la mesure du gain du
chemin sans changement. Le changement de politique `go.mod` du protocole sera
appliqué au pilote lors de ce raccordement.

Preuves conservées dans
`<tmp>/gopurs-cache-contract-uynq0oxb/` :
`before.json`, compilateurs `before/` et `after/`,
`verify-codegen.mjs`, `before-codegen.json`, `after-codegen.json`,
`final-codegen.json`, les TAST et sorties de `fixtures/`, les workspaces natifs,
`cache-tests-final.log`, `linux-cache-tests-final.log`, `linux-command.json`,
`repro-missing-journal.mjs`, `journal-{before,after}.log`,
`snapshot-review.json`, `strict-replays.json`, `final-audit.json` et `README.md`.

## Plan v3 — lot 03 : build identique vérifié, 10 octobre 2026

**Lot validé** après les contrôles de correction, les mesures appariées et
l'audit final. Le plan v3 compte **4/7 lots validés**.

### Raccordement et contrats exercés

Le pilote vérifie maintenant les empreintes avant le décodage des TAST. Un hit
saute préparation, monomorphisation, optimisation et émission ; un miss consomme
des copies privées des TAST et FFI, puis publie par le stockage du lot 02.
La résolution FFI commune produit sa propre trace des candidats absents et des
répertoires parcourus. `Host.Snapshot` copie puis réexécute les octets
identifiés, avec les ressources WASM/JS ou l'archive native intégrée.

Le protocole conserve `go.mod` existant et `go.sum`, retire seulement les sorties
obsolètes possédées et élide les écritures byte-identiques. La migration d'une
ancienne sortie exige une correspondance avec la génération fraîche. Un nouvel
output ne peut pas écraser un fichier externe différent, même après la création
du journal. `--force`, `GOPURS_FORCE_BUILD=1` et `GOPURS_ALLOC_PROFILE` demandent
un calcul complet. Les directives sont toujours chargées ; un diagnostic Rebox
marque le manifeste pour imposer un nouveau calcul au lieu de masquer le message.
Le cache de travail `.purmeta` du PBO JS est privé par invocation.

### Vérifications

- Bootstrap JS/Go réussi : **505 modules / 289 658 types**. Les compilateurs
  utilisés sont copiés et inventoriés avant les mesures.
- `npm run test:cache` : stockage sous détecteur de courses, vrais bridges
  JS/Go, remplacement du bundle et du binaire **après leur chargement**, puis
  vérification du code exécuté et de son SHA-256 après réexécution.
- Le stockage et le producteur passent aussi sous détecteur de courses avec
  le **Go 1.26.8 Linux/arm64** de l'image b8x ; le contrôle hôte utilise
  **Go 1.27.0 macOS/arm64**. Les éditions aller-retour consomment les octets
  figés ; une édition persistante refuse la publication.
- `test:cache-driver` : quatre modes **JS / Go séquentiel / Go parallèle / Rust**,
  chaque transition comparée octet par octet à une génération neuve des mêmes
  TAST et aux autres modes. La suite couvre hits sans passes ni changements de
  dates, calcul forcé, tidy, sorties endommagées, edits à date conservée,
  ajouts/retraits de modules et bridges, entrées, priorités et absences FFI,
  collisions externes, TAST/FFI invalides, SIGTERM, concurrence, profils et
  récupération des métadonnées. Les quatre applications finales s'exécutent
  avec Go et rendent `Done`.
- **31 contrôles CLI** réussis, incluant Rust séquentiel/parallèle. Les erreurs
  restent de statut 1 ; les conflits des destinations publiques sont désormais
  signalés à `cache publish`, et les erreurs de TAST incomplet empêchent une
  publication partielle.
- Trois snapshots `FFIIntegerReturns`, son bridge, et `NativeRecordWorkers`
  examinés et **inchangés** ; **six replays stricts** réussis avec compilation
  et exécution Go dans les trois modes demandés. Aucune mise à jour de snapshot.
- Suite PBO `purmeta-build-cache` réussie avec le bundle reconstruit.
- L'hôte Rust optionnel est construit et vérifié dans un candidat O3/ThinLTO
  **sans PGO**, sélectionné explicitement par les tests ; son binaire installé
  n'est pas remplacé par ce candidat de vérification.

Deux reproducteurs sont conservés : le hit sans réécriture échoue avec l'ancien
compilateur puis passe avec le nouveau ; la collision d'une nouvelle sortie avec
un fichier externe échoue avant le correctif de propriété puis passe après.

### Mesure du backend b8x

La copie du corpus du lot 00 contient les mêmes **2 704 TAST / 755 005 types**.
Les **3 010 fichiers Go**, runtime, bridges et entrées compris, sont byte-exacts
avec cette référence dans les trois modes demandés. Les comparaisons de contenu
et de dates sont effectuées hors chronométrage. Les tests de correction sont
terminés avant les mesures ; aucun autre build de cette campagne n'est lancé
pendant ces mesures.

Trois paires alternées, avec le launcher et ses huit workers usuels :

| Processus gopurs | Avant | Hit après raccordement |
| --- | ---: | ---: |
| Échantillons | 35,798 / 35,498 / 35,732 s | 1,257 / 1,249 / 1,521 s |
| Médiane | **35,732 s** | **1,257 s** |
| Fichiers Go réécrits par invocation | 3 010 | **0** |

La durée médiane diminue de **96,5 %** (environ **28,4×**). La médiane interne
de vérification du cache est **684 ms** ; les **1,257 s** mesurent le processus
entier, copie/réexécution du compilateur comprises. Aucun hit ne décode les
TAST ni n'entre dans les passes coûteuses. La référence historique du lot 00
reste **36,902 s** ; le chiffre avant ci-dessus est la mesure appariée du lot 03.

Le coût froid est significatif et conservé explicitement :

| Parcours après raccordement — un passage par ligne | Total | Publication |
| --- | ---: | ---: |
| Première adoption, Go parallèle, sorties Go déjà présentes | **65,054 s** | **25,660 s** |
| Sortie et cache neufs, Go séquentiel | 135,060 s | 53,396 s |
| Sortie et cache neufs, JS | 117,375 s | 53,979 s |
| Retour JS → Go parallèle, objets déjà présents | 39,936 s | 1,371 s |

L'adoption initiale conserve les dates des 3 010 Go identiques, mais crée les
objets durables. Une sortie neuve écrit aussi les fichiers publics : les écritures
et synchronisations du stockage expliquent le coût de publication observé. Les
deux générations neuves ci-dessus servent aussi à la parité ; ce ne sont pas des
médianes. Leurs hits suivants prennent respectivement **1,453 s** et **1,420 s**.
Un miss conserve encore le calcul global ; les gains par module restent au lot 04.

### Parcours Spago → tidy → runner

Une préparation puis trois paires alternées utilisent le même frontend figé,
le Go Linux de b8x, les caches Go chauds et un verrou/exécutable privés. La
fenêtre continue inclut Docker et s'arrête avant `exec` de `Test.Main` ; elle
exclut les inventaires et la suite métier PostgreSQL.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Chaîne — échantillons | 38,145 / 39,290 / 37,611 s | 4,180 / 3,958 / 3,539 s |
| **Chaîne — médiane** | **38,145 s** | **3,958 s** |
| Spago + backend — médiane | 36,022 s | 2,933 s |
| `go mod tidy` — médiane | 1,070 s | 0,256 s |
| Docker + runner jusqu'à `exec` — médiane | 1,051 s | 0,794 s |

Le gain médian de la chaîne est **89,6 %**, environ **9,64×**. Les médianes des
phases ne s'additionnent pas nécessairement à celle de la chaîne. La référence
historique du lot 00 était **40,808 s**. Le gain de tidy accompagne la
conservation de `go.mod`, que l'ancien pilote réinitialisait à chaque passage.

Tous les graphes d'actions mesurés contiennent **zéro compilation et zéro lien**.
Les réglages restent `GOMAXPROCS=2`, `-p 1`, `GOMEMLIMIT=12GiB` ; chaque attente
du verrou privé est affichée à **0,000 s**. L'exécutable de tests conserve le
SHA-256 du lot 00, `9e08ce7474320728fac10c16d721e050b827a495f14eb3c10a70de970a546378`.
Après les horloges, l'oracle rend `camel=""`, `accent="e"`, `pong` et le vrai
`Ping` rend `pong\n`. Le conteneur dédié est supprimé.

Le cache final occupe **137 272 731 octets** (environ **131 MiB**, 2 670 fichiers),
avec un journal de propriété séparé de **675 784 octets**. L'audit contrôle
**4 318 fichiers inventoriés**, la constance des **8 112 observations de
sources/compagnons** du lot 00, les 14 passages backend, les sept passages de chaîne,
les artefacts figés, les sorties finales et les snapshots. La configuration
backend est la seule différence de la copie source ; b8x actif reste en JS.

Le premier harnais de chaîne incluait le pré-inventaire des dates dans l'horloge.
Ses résultats sont conservés dans `cycle-preliminary-inventory-clock/` ; seules
les mesures corrigées, avec cet inventaire hors horloge, alimentent ce bilan.

### Preuves et rejeu

Répertoire local :
`<tmp>/gopurs-nochange-lot03-b33v4ie0/`.
Son `README.md` documente les artefacts `before`, `candidate-1`, `after`, les
reproducteurs, les commandes de campagne et le rejeu sur une copie neuve.
Les workspaces des transitions et des replays stricts sont conservés avec leurs
logs et inventaires ; les compilateurs, ressources et corpus portent leurs
SHA-256. Les essais intermédiaires échoués sont gardés avec les campagnes finales.

<a id="plan-v3--lot-04--incrémentalité-par-module-validation-en-cours-11-octobre-2026"></a>

## Plan v3 — lot 04 : incrémentalité par module, validée (12 octobre 2026)

**Clôturé sur `publicsync`.** Le [bilan livré](build-incremental.md) rassemble les
mesures appariées, les compteurs, les limites et les preuves des lots 04–05.
Face au lot 03 : **−10,37 % feuille / −4,48 % partagé / −32,76 % FFI**, avec
**+20,19 % à froid**, **+0,350 s sur l'identique suivant le froid** et **+413,347 Mo
d'objets**. L'utilisateur confirme la poursuite avec ce compromis. La validation
complète réussit sur les compilateurs installés ; l'historique des essais suit.

Le candidat `allhit` conserve la préparation globale et utilise les clés des
modules spécialisés, le préfixe PBO conservateur et le contexte complet de
l'émetteur. Si toutes les clés PBO correspondent, les backends/workers sont
validés puis transmis directement à l'émetteur canonique ; les contributions PBO
non consommées sont republiées depuis leurs copies privées vérifiées. Un miss
réactive le builder et le décodage métier de ses contributions. Le contrat et
les bridges sont décrits dans [build-cache.md](build-cache.md).

### Gain confirmé sur édition FFI

Cinq paires alternées utilisent les compilateurs figés, le corpus TAST du lot 00,
des caches engagés propres à chaque copie et les mêmes options natives parallèles
(`prepare=8`, `pbo=8`, `emit=8`, pipeline actif). Aucune compilation lourde ne
tourne pendant ces mesures. Les horloges incluent le lancement et la réexécution
du compilateur figé ; elles excluent le frontend, le build Go et les inventaires.

| Passage | Lot 03 | Candidat allhit |
| --- | ---: | ---: |
| 1 | 37,898 s | 30,816 s |
| 2 | 38,332 s | 30,724 s |
| 3 | 38,594 s | 31,156 s |
| 4 | 38,124 s | 30,542 s |
| 5 | 38,525 s | 30,663 s |
| **Médiane** | **38,332 s** | **30,724 s** |

Gain médian : **19,85 %**, soit environ **7,61 s**. Les dix inventaires contiennent
les **3 010 Go** attendus du lot 00 ; chaque édition ne réécrit que
`Util_Type_String_String_ffi.go`. Le candidat réutilise les **2 704 résultats PBO**
et restaure **2 703 émissions**, avec les mêmes empreintes préparées et les mêmes
**397 516 134 octets** des quatre parties métier que les candidats antérieurs.

### Campagne b8x complète : correction acquise, performances insuffisantes

La campagne suivante exécute successivement le candidat et le lot 03 sur des
racines privées neuves, sans objets préchargés. Elle supprime seulement les
3 010 Go connus de la copie du corpus avant le premier passage, puis compare
chaque variante aux inventaires du lot 00 et à une génération forcée sur les
mêmes TAST. Les chronos ci-dessous sont des passages isolés, pas des médianes.

| Scénario | Lot 03 | Candidat allhit | Go réécrits |
| --- | ---: | ---: | ---: |
| Cache et sorties neufs | 93,599 s | 239,367 s | 3 010 |
| Relance identique | 1,471 s | 2,355 s | 0 |
| Édition feuille | 39,505 s | 53,865 s | 1 |
| Fonction partagée | 39,111 s | 88,715 s | 5 |
| Édition FFI | 37,789 s | 32,151 s | 1 |

| Variante incrémentale | PBO recalculés / réutilisés | Émissions nouvelles / restaurées | Candidat forcé |
| --- | ---: | ---: | ---: |
| Feuille | 800 / 1 904 | 1 / 2 703 | 88,826 s |
| Partagée | 2 697 / 7 | 2 697 / 7 | 85,337 s |
| FFI | 0 / 2 704 | 1 / 2 703 | 90,340 s |

Chaque build forcé conserve les dates des sorties byte-identiques. Les quatre
artefacts métier et les empreintes préparées correspondent aux builds
incrémentaux. Pour la feuille, un module préparé change et le préfixe invalide
799 successeurs. Pour la fonction partagée, deux modules préparés changent
(spécialisations incluses) et le préfixe invalide 2 695 successeurs. La clé des
métadonnées globales et les environnements d'émission limitent aussi la
restauration, bien que seuls cinq contenus Go changent réellement.

Sur le premier build neuf, la publication durable prend **152,943 s**, contre
**55,330 s** au lot 03. Les objets occupent respectivement **534 973 714 octets /
10 992 objets** et **135 273 391 octets / 2 666 objets**. Après les variantes,
le candidat conserve **552 370 422 octets / 11 016 objets** ; ces tailles incluent
les sorties et les manifestes adressés par contenu, pas le staging ni le journal
de propriété. Le volume logique des quatre artefacts de référence est toujours
**397 516 134 octets**.

L'audit complet `audit-allhit.json` passe : **2 704 TAST et 3 010 Go**, mêmes
inventaires TAST, sorties et quatre artefacts métier entre JS, Go séquentiel et
Go parallèle sur le corpus original puis l'édition FFI. Les passages JS et Go
séquentiel utilisent des objets préchargés et servent à la parité ; leurs temps
ne sont pas des mesures de cache froid. Les treize transitions plus petites
couvrent en plus Rust sur les mêmes TAST.

**Le lot 04 reste non validé (4/7)** : la réutilisation et la correction sont
vérifiées, mais le gain FFI ne compense pas les régressions des éditions TAST et
du premier build. L'encodage des clés/artefacts et l'invalidation conservatrice
restent à traiter avant de conclure au bénéfice global. Les variantes complètes
conservent les preuves de cette régression ; elles ne sont pas remplacées par
les seuls résultats favorables du chemin FFI.

### Optimisation du rendu après le bilan allhit

Le candidat `wirecalls` remplace les compositions de fonctions du rendu wire
par des appels explicites et construit directement les enveloppes JSON fixes.
Les bridges et les octets persistants restent identiques. Sur trois paires
alternées de bascules de la fonction partagée, la médiane passe de **91,353 s
(`allhit`) à 80,329 s (`wirecalls`), soit −12,07 %**. Les six passages vérifient
les artefacts, les empreintes préparées, les 3 010 Go et exactement cinq
réécritures. Le bootstrap Go, huit tests de codecs et le rejeu JS des
193 248 099 octets d'implémentations PBO passent.

L'essai `exprmemo` corrige le parcours répété des expressions
partagées lors de l'encodage. `encode-sharing-before.log` échoue dès la profondeur
1 sur deux visites de la même feuille ; `encode-sharing-after.log` passe jusqu'à
la profondeur 24 sans développer les 16 777 216 feuilles logiques. Les neuf
contrôles de codecs/bridges passent, y compris rétention des objets après GC,
isolation des artefacts et distinction des tags nuls. Les bootstraps Go/Rust,
le rejeu JS du corpus et les quatre artefacts du probe sont byte-exacts. Le
temps sur édition partagée, **76,695 s contre 77,223 s pour `wirecalls`**, ne
confirme toutefois pas de gain significatif. L'essai et son reproducteur sont
conservés dans les preuves ; les sources, tests et compilateurs actifs sont
revenus à `wirecalls` (`restore-wirecalls-build.log`).

`smallctors` remplace ensuite les codecs génériques des petits constructeurs
fréquents (`Maybe`, `Tuple`, noms, `Qualified`, `Prop`, `Literal`) par des
branches PureScript explicites. Trois paires alternées confirment **76,673 s
(`wirecalls`) → 69,201 s (`smallctors`), soit −9,75 %** sur l'édition partagée.
Les artefacts, les clés préparées et les 3 010 Go sont byte-exacts, avec cinq
réécritures par bascule. Le bootstrap Go, les huit tests de codecs et le rejeu
JS des implémentations passent. `smallctors-wire-audit.json` vérifie aussi les
clés dans les manifestes conservés : le premier bilan avait échoué sur un
fichier de preuve non exporté, puis l'export du harnais a été complété.

Le coût froid fait l'objet d'un essai indépendant, `stagepool` : staging des
objets dans huit tâches au plus, emplacements fixes du manifeste et attente de
toutes les tâches avant toute publication, y compris après une erreur. Le
protocole, les SHA-256 et les synchronisations de chaque objet restent les mêmes.
La suite Go du stockage passe sous race detector sur macOS, notamment le test
des objets partagés, de l'ordre canonique et d'un échec tardif avant publication.
Deux paires alternées de publication froide des mêmes 3 010 sorties et 2 704
mémos, dans des racines neuves, confirment **174,391 → 95,186 s en moyenne,
soit −45,42 %**. Les sorties, mémos et objets sont byte-exacts, puis le lookup
retourne un hit. La mesure isole le stockage : l'inventaire TAST est vide et
aucune conversion ne s'exécute. `stagepool-corpus-files-summary.json` conserve
ce bilan. La campagne initiale par pipes avait dépassé son délai sur la seconde
référence avant toute écriture d'objet ; les quatre mesures retenues utilisent
des fichiers pour stdin/stdout/stderr. `stagepool` est retenu.

L'essai `workerkeys` remplace le préfixe des workers dans la clé d'émission par
les signatures référencées après les transformations locales de `CodeGen`,
avec leurs absences. Le rendu consomme exactement cette vue ; les métadonnées
globales, le résultat PBO complet et la FFI restent dans la clé. Le reproducteur
`workerkeys-before-2.log` échoue sur une invalidation par un worker sans rapport ;
le correctif passe aussi les changements de signature, retraits et ajouts, et
le test compilé des appels natifs importés. Son bootstrap natif réussit.
Le premier probe, combiné à `stagepool`, prend **59,547 s** sur l'édition
partagée, avec **6 émissions / 2 698 restaurations**, cinq réécritures et les
3 010 Go, quatre artefacts et clés préparées byte-exacts. Trois paires alternées
contre `stagepool` seul confirment ensuite **68,087 → 60,866 s de médiane,
soit −10,60 %**, avec cinq réécritures à chaque bascule et les mêmes parties et
empreintes préparées. `workerkeys` est retenu. Le candidat `stagepool` a été
relié depuis une copie du bootstrap `smallctors` ; `stagepool-native-audit.json`
confirme que seul `gopurs_build_cache/store.go` diffère dans le Go compilé.

L'essai suivant, `publishpool`, applique la même borne de huit tâches aux
écritures publiques, après publication durable du journal d'union. Un test
injecte une corruption d'objet après cette barrière, exige la fin des autres
écritures et un pointeur inchangé, puis reprend par le protocole normal. Les
tests passent sous race detector, mais les deux paires froides ne gagnent que
**5,87 %** (**92,655 → 87,215 s** en moyenne), sous le seuil préalable de 10 %.
L'essai est retiré ; `store.go` et `stage_test.go` sont restaurés à l'octet près
depuis `stagepool`, avec leurs versions expérimentales et les mesures conservées.
La condition de gain a arrêté la campagne avant le relien natif et le probe b8x
neuf : aucun résultat b8x froid ne lui est attribué.

`wireids` essaie ensuite de transporter les références du codec sous forme
d'entiers et de rendre les constructeurs/tableaux dans le pool générique.
Le bootstrap Go, onze contrôles de codecs/bridges et d'émission, le rejeu JS
des 193 248 099 octets d'implémentations et les deux corpus du probe sont
byte-exacts. Le probe partagé prend toutefois **62,814 s**, sans gain établi
face à `workerkeys` (**59,547 s** sur son probe et **60,866 s** de médiane sur
les paires précédentes). L'essai est retiré ; les cinq sources et le test restent
dans `candidate-wireids`. Les sources actives reviennent à `workerkeys`, avec
reconstruction JS et restauration atomique de son binaire natif conservé.
Les bridges Rust expérimentaux n'ont pas fait l'objet d'un bootstrap.

Le profil du stockage `stagepool` mesure **88,26 s écoulées / 25,56 s de CPU**,
dont **3,11 s de CPU** pour la validation wire ; les opérations de fichiers
dominent. L'essai `batchsync` conserve le `fsync` de chaque fichier et son
renommage atomique, mais synchronise le répertoire des objets une fois après
jonction de toutes les tâches. La barrière précède le journal de propriété et
s'applique même aux objets déjà présents. Les sorties publiques et le commit
conservent leur synchronisation immédiate. Les tests passent sous race detector
macOS et Linux/arm64 Go 1.26.8. Deux paires froides alternées confirment
**93,644 → 71,312 s en moyenne, soit −23,85 %** ; les 3 010 sorties, 2 704 mémos
et objets sont byte-exacts et les lookups retournent un hit. `batchsync` est
retenu. L'audit du relien natif montre que seul `gopurs_build_cache/store.go`
diffère du bootstrap `workerkeys`.

Le probe b8x sans objets préalimentés ni sorties Go prend **146,488 s**, dont
**76,246 s** de publication ; il écrit les 3 010 Go attendus. La relance identique
prend **2,010 s**, sans réécriture, et l'édition partagée **64,642 s**, avec cinq
réécritures, 6 émissions et 2 698 restaurations. Les sorties, parties métier et
clés préparées sont identiques à `workerkeys`. Ces passages isolés ne sont pas
des paires ; le froid reste au-dessus des **93,599 s** du lot 03.

Le profil séparé des codecs `workerkeys` conserve les allers-retours byte-exacts.
Backend et implémentations allouent respectivement **26,51 / 29,04 Go** en
écriture et **34,35 / 41,43 Go** en lecture, pour **179,41 / 193,25 Mo** de wire.
Le profil d'allocations échantillonné, séparé des mesures de vitesse, attribue
**22,56 %** des octets alloués à la construction de `codecBackendSyntax` et de
ses codecs auxiliaires. L'essai `syntaxfns` extrait les branches dans des
fonctions appelées directement par le parcours récursif. Le bootstrap Go, les
dix contrôles de codecs/workers, le corpus JS et les deux corpus b8x passent.
Les allocations d'encodage backend/implémentations passent à **15,86 / 18,56 Go**
(contre **26,51 / 29,04 Go**), et celles de décodage à **28,38 / 35,62 Go**
(contre **34,35 / 41,43 Go**). Chaque aller-retour natif reproduit les octets.
Le probe partagé prend **59,284 s**, contre **64,642 s** pour `batchsync`, avec
les mêmes parties, clés préparées, 3 010 Go et cinq réécritures. Trois paires
alternées contre `batchsync` confirment **63,502 → 59,211 s de médiane,
soit −6,76 % / −4,29 s**. Chaque paire gagne entre **3,31 et 5,60 s** ; les
artefacts, empreintes et Go restent byte-exacts, avec cinq réécritures par
bascule. `syntaxfns` est retenu, avec les bridges et le schéma wire actuels.

`validate-syntaxfns.py` lance ensuite la campagne complète froide/identique/
feuille/partagée/FFI, trois paires actuelles face au lot 03, la parité b8x JS/Go
séquentiel/parallèle, le bootstrap Rust, les transitions dans les quatre hôtes,
les six replays stricts et les contrôles CLI. Les sources et binaires sont
audités contre `candidate-syntaxfns` avant et après. Les trois snapshots sont
byte-identiques aux références lot 03/allhit déjà examinées, avec le réexamen
conservé dans `snapshot-review-post-batchsync.json`.

La campagne b8x complète passe, avec artefacts et sorties incrémentaux/forcés
byte-exacts et zéro réécriture sur chaque passage forcé après mutation :

| Scénario | Lot 03, référence conservée | `allhit` | `syntaxfns` |
| --- | ---: | ---: | ---: |
| Cache et sorties neufs | 93,599 s | 239,367 s | 133,543 s |
| Identique | 1,471 s | 2,355 s | 2,019 s |
| Feuille | 39,505 s | 53,865 s | 49,502 s |
| Partagée | 39,111 s | 88,715 s | 59,785 s |
| FFI | 37,789 s | 32,151 s | 35,500 s |

Ce sont des passages isolés. Le froid publie en **73,303 s** ; les mutations
réécrivent toujours **1 / 5 / 1 Go**. Feuille et partagée réutilisent respectivement
**1 904 / 7 résultats PBO** et **2 703 / 2 698 émissions**. La FFI réutilise les
2 704 résultats PBO et 2 703 émissions, mais les transformations locales avant
chaque lookup d'émission contribuent encore au coût des hits.

Le harnais s'est ensuite arrêté sur `after-full-prepared.json`, absent de
l'ancienne campagne `allhit` qui n'exportait pas encore ces clés. Le script
initial est conservé dans `validate-syntaxfns-before-recovery.py` ;
`recover-allhit-prepared.py` récupère les clés des manifestes adressés par contenu
en exigeant les mêmes sorties et parties métier, ainsi que l'accord de tous les
manifestes correspondants. La reprise conserve les dix passages réussis et
termine toutes les vérifications :

- **Trois paires actuelles face au lot 03** sur l'édition partagée : médianes
  **37,785 → 58,459 s**, soit une régression de **54,72 %**. Le gain cumulé des
  optimisations du lot 04 ne suffit toujours pas sur ce scénario.
- Parité **JS / Go séquentiel / Go parallèle** sur le corpus original puis FFI :
  mêmes **2 704 TAST, 3 010 Go**, parties métier et empreintes préparées.
- Bootstrap Rust privé, puis treize transitions incrémental/frais dans les
  **quatre hôtes**, avec comparaisons des sorties, parties et clés préparées.
  Workspace conservé : `gopurs-tests-fKXOWc`.
- **Six replays stricts / 35 commandes réussies**, sans mise à jour des trois
  snapshots réexaminés, puis **19 contrôles CLI** verts.
- Audits avant/après de **346 fichiers des sources, outils et binaires**, tous
  identiques au candidat figé `syntaxfns`. Le Rust privé est identifié par
  `gopurs-rust-syntaxfns.json` ; les binaires Rust des anciennes copies ne le
  remplacent pas.

`validation-syntaxfns-summary.json` et `validation-syntaxfns-counts.json` consignent
ce bilan. **Le lot reste non validé**, à cause des régressions des éditions TAST
et du froid.

L'essai suivant, `restoreprep`, anticipe le lookup d'émission à partir des
références externes du backend d'entrée. Cette clé n'est engagée qu'après
vérification, sur un miss, que les références transformées y sont incluses.
Backend, métadonnées, compilateur/options, signatures concernées (absences
comprises) et FFI sont couverts. Une référence introduite hors de cette vue
conserve le chemin post-transformation. Le hit reprend également le digest
vérifié des workers plutôt que de les réencoder. Les bridges et les quatre
schémas métier restent identiques. Le reproducteur `restoreprep-before.log`
échoue sur **deux lectures de métadonnées de préparation au lieu de zéro** pour
un hit. Après application, le reproducteur passe, ainsi que le bootstrap Go et
les dix contrôles de codecs/workers. Les probes prennent **30,485 s en FFI** et
**53,037 s sur l'édition partagée**, avec parties, clés préparées et sorties
byte-exactes ; les réécritures restent **1 / 5 Go**. En FFI, le temps cumulé des
lots d'émission descend de **10,575 à 3,686 s**. Les probes sont préalimentés en
objets et ne constituent pas une mesure froide.

Trois paires FFI alternées face à `syntaxfns` confirment **36,325 → 29,725 s**,
soit **−18,17 %**, au-dessus du seuil fixé avant confirmation de **10 %**.
Les trois paires partagées donnent **58,859 → 54,223 s**, soit **−7,88 %**.
Chaque paire gagne dans les deux scénarios. `restoreprep` est retenu. Le harnais
`paired-modules.mjs` vérifie aussi les 2 704 TAST inchangés lors des bascules FFI,
les compteurs attendus et la réécriture du seul bridge FFI concerné. Un complément
du reproducteur couvre le cas où `BorrowedObjects` introduit `borrowObject`
hors des références d'entrée : le lookup doit alors suivre la préparation et
couvrir la présence, la signature et l'absence de ce helper. Il passe dans
`restoreprep-fallback-1.log`, puis est intégré à `emitter-worker-keys.test.mjs`.

`candidate-restoreprep-final` reprend les sources de production et binaires
mesurés ; seul ce complément du test diffère (`restoreprep-final-freeze.json`).
`validate-restoreprep.py` lance ensuite la campagne froide complète, les paires
actuelles contre le lot 03, la parité b8x des trois hôtes, le bootstrap Rust,
les transitions des quatre hôtes, les replays stricts, la CLI et les audits.
Les tests de stockage macOS/Linux restent rattachés à `batchsync`, avec contrôle
byte-exact des sources du stockage. La campagne est terminée et verte :

| Scénario | `restoreprep`, passage complet |
| --- | ---: |
| Cache et sorties neufs | 136,821 s |
| Identique | 2,045 s |
| Feuille | 45,440 s |
| Partagée | 55,780 s |
| FFI | 30,215 s |

Le froid part sans objet préalimenté ni sortie Go et publie en **74,458 s**.
Les trois mutations réécrivent **1 / 5 / 1 Go** ; leurs contrôles forcés donnent
les mêmes sorties, parties métier et clés préparées, sans réécriture. Les trois
paires actuelles face au lot 03 donnent **37,550 → 51,087 s (+36,05 %)** sur
l'édition partagée. Les gains intermédiaires sont confirmés, mais ce scénario,
la feuille et le froid restent au-dessus de la référence lot 03.

La parité b8x original/FFI est byte-exacte entre JS, Go séquentiel et Go parallèle :
**2 704 TAST / 3 010 Go**, quatre parties et clés préparées identiques. Le
bootstrap Rust privé et les treize transitions des quatre hôtes passent
(`gopurs-tests-1dm4oB`). Les **six replays stricts / 35 commandes**, les trois
snapshots inchangés, les **19 contrôles CLI** et l'audit des **346 fichiers** de
sources/outils/binaires passent également. `validation-restoreprep-summary.json`
et `validation-restoreprep-counts.json` consignent le bilan ; le Rust correspondant
est identifié dans `gopurs-rust-restoreprep.json`. **Le lot 04 reste non validé.**

La piste suivante est diagnostique : `trace-pbo-reads.py` instrumente une copie
privée du bundle JS, avec toutes les valeurs des lookups retournées à l'identique.
Il enregistre les implémentations consultées et absentes, les directives lues et
les contributions par binding, sous empreintes wire v1. Deux builds forcés du
corpus original/partagé doivent rester byte-exacts. La comparaison confronte les
lectures du build original à l'environnement du build partagé, dans l'ordre
canonique, puis vérifie toute réutilisation prédite contre les trois vrais
artefacts PBO. Elle mesure le potentiel d'une preuve plus fine que le préfixe,
sans encore établir sa correction générale ni un gain de vitesse. Les temps de
cette instrumentation sont exclusivement diagnostiques.

La trace termine avec parité byte-exacte du corpus original/partagé. Elle prédit
**2 643 modules réutilisables / 61 invalidations**, contre sept hits avec le
préfixe ; les trois artefacts PBO de chaque réutilisation prédite sont identiques
au build forcé. Elle compte **65 178 lectures d'implémentations**, dont **10 565
absentes**, et **44 624 lectures de directives**. Une dépendance au module entier
invaliderait encore **899 consommateurs**. `pboreads-trace/summary.json` conserve
les détails, et `pbo-read-proof-design.md` la conception préalable.

L'essai `pboreads` ajoute une observation propre à chaque conversion et valide
ses lectures dans la vue de rang courante lorsque le préfixe diffère. La première
partie devient `pbo/backend-v2` (`body` puis preuve optionnelle), avec les trois
autres codecs et la grammaire wire inchangés. Les anciens manifestes v1 restent
lisibles par le stockage pour la propriété/récupération. Le seuil d'acceptation
annoncé avant mesure est **10 % sur trois paires partagées face à `restoreprep`**.

Le reproducteur `pboreads-before.log` échoue sur **« an unrelated predecessor
invalidated optimization »** ; après raccordement il passe, ainsi que les
contrôles de signatures workers. `optimizer-observation.mjs` passe sur les
directives initiales malgré overrides, absences, références tardives, forced
private inlines, tentative différée et publication canonique. Les cinq contrôles
de visibilité/API historiques et les quatre contrôles de lookup passent aussi.
Le contrôle de publication des directives passe ; l'autre cas de l'ancien
`module-directive-range.mjs` appelle `Native.filterRange`, absent également des
sources de `restoreprep-final`. Cet échec de harnais est conservé dans
`pboreads-targeted-js-2.log`, puis seul le contrôle applicable est rejoué.
Le bootstrap Go, le stockage macOS sous race detector, les quatre autres fixtures
sémantiques et les **onze contrôles codecs/workers/cache** passent. Le natif est
conservé dans `gopurs-native-build-1Vpr2c`, avec pointeur
`pboreads-native-output.json`. Le probe partagé prend **43,631 s** : **61 modules
optimisés / 2 643 réutilisés**, **6 émissions / 2 698 restaurations**, cinq sorties
Go réécrites. La partie backend-v2 porte le total des quatre artefacts de
**397 516 134 à 411 148 559 octets** sur l'original. L'initialisation est
préalimentée en objets et ne mesure pas un cache froid.

Les trois paires alternées face à `restoreprep-final` confirment **53,687 →
45,051 s**, soit **−16,09 %**, avec un gain dans chaque paire. Le seuil préalable
est atteint : **`pboreads` est retenu**. Les corps v1 extraits et les trois autres
parties, TAST, clés préparées et sorties Go sont byte-exacts. Les preuves
persistées de l'original contiennent **63 588 références d'implémentations et
32 561 références de directives** : les références locales/au module courant,
présentes dans la trace diagnostique, sont couvertes par le module préparé et
le contexte initial et ne sont pas persistées.

Le probe conserve aussi un écart diagnostique :
64 raisons d'invalidation pour 61 reconstructions, car une tentative différée
devenue hit conserve sa raison spéculative. `pboreads-counter-before.log`
reproduit cet écart ; le hit retire maintenant sa raison devenue caduque. Le
reproducteur enrichi est intégré à `optimizer-read-cache.test.mjs` et passe.
`optimizer-read-build.test.mjs` ajoute **sept transitions** cache/frais, chacune
en séquentiel et LIFO : module sans rapport, binding sans rapport dans un
fournisseur consulté, corps inliné, directive, apparition/disparition d'une
référence et changement de rang. Les sorties Go et les quatre parties, preuves
comprises, sont identiques aux générations fraîches. Le comptage ne conserve
que les misses finaux.

`backend-body-v1.mjs` extrait le préfixe de graphe du champ `body` du v2 et
reconstitue le wire v1 exact, avec contrôle de la table de types. Le harnais
`paired-readproof.mjs` compare ainsi les octets du corps et les trois autres
parties entre candidats, puis les quatre objets complets lors des retours à un
même état pour un même candidat. Toutes ces vérifications, avec les mêmes TAST
et inventaires Go, sont hors des chronométrages.

`candidate-pboreads-final` est reconstruit et figé après la correction du compteur
et les compléments de tests. Le delta vérifié est dans
`pboreads-final-delta.json` : seul `Cache.Modules` change dans les sources de
production. Le bootstrap Go et les **douze contrôles codecs/workers/cache**
passent (`native-build-pboreads-final.log`, `pboreads-final-codecs-workers.log`).
`validate-pboreads.py` termine la campagne froide complète, les paires actuelles
face au lot 03, la parité b8x JS/Go sur original, feuille, partagé et FFI avec
preuves v2, le bootstrap Rust, les transitions des quatre hôtes, le stockage
Linux sous race detector, les replays stricts et les audits. Toutes ces
vérifications passent. **Le lot 04 reste non validé pour ses performances.**

Les dix passages Go parallèles passent,
avec **135,630 s froid / 1,738 s identique / 42,671 s feuille / 44,590 s partagée /
30,768 s FFI**. Le froid publie en **71,822 s**. La feuille réoptimise un module
et en restaure 2 703 ; le partagé réoptimise 61 modules et en restaure 2 643.
Les raisons d'invalidation sont maintenant cohérentes avec ces comptes :
**59 lectures modifiées + deux modules préparés** sur le partagé. Les mutations
réécrivent **1 / 5 / 1 Go** ; les contrôles forcés donnent les mêmes quatre parties
et sorties, sans réécriture supplémentaire.

Les trois paires actuelles face au lot 03 donnent **39,404 → 46,093 s**, soit
**+16,97 %** sur le partagé. La régression se réduit, mais demeure. Les campagnes
JS et Go séquentiel passent également : mêmes **2 704 TAST / 3 010 Go**, clés
préparées et quatre parties, preuves de lecture comprises, dans les trois hôtes
sur les états original, feuille, partagé et FFI. Le stockage Linux/arm64 passe
sous race detector avec Go **1.26.8**, sur les mêmes sources que le test macOS.

Le bootstrap Rust privé sans PGO et les **treize transitions dans les quatre
hôtes** passent. Le workspace des transitions est `gopurs-tests-NYbvFq` ; le
binaire effectivement testé est `gopurs-rust-pboreads`, SHA-256
`010c32157b43e86cf3a055e38bdc053652253a3834d6d2489225852947b1f22e`.
Les six replays stricts comportent **35 commandes réussies**, puis les **19
contrôles CLI** passent. Les trois snapshots examinés restent byte-identiques.
L'audit final couvre **411 fichiers des sources/outils et binaires**, égaux au
candidat figé ; son périmètre inclut aussi les tests PBO. Les pointeurs et comptes
sont dans `validation-pboreads-summary.json` et `validation-pboreads-counts.json`.

Le profil partagé suivant (`probe-b8x-pboreads-shared-profile`) conserve les
mêmes octets et compteurs. `Wire.read` cumule **19,70 s CPU** ; le profil met
aussi en évidence le coût du GC et des allocations. Le harnais lié au vrai
package Go `gopurs-native-build-Zo3sdv` mesure **32,210 / 36,017 / 1,748 /
4,283 Go d'allocations cumulées** pour décoder les quatre parties. Il conserve
ses profils et résultats dans `codec-pboreads-decode-2`. La première tentative
avait sélectionné un mauvais getter généré pour les workers ; son échec après
les trois autres parties reste dans `codec-pboreads-decode.log`.

L'essai suivant, `scalarhex`, remplace le tableau temporaire de caractères du
prédicat `hex` par un parcours direct dans les bridges scalaires JS/Go/Rust.
L'ensemble accepté reste `[0-9a-f]*`, chaîne vide comprise ; les contrôles de
longueur et de réencodage restent dans le codec partagé. Le reproducteur parcourt
les **65 536 unités UTF-16** via le vrai décodeur, avec empreinte avant/après.
Le seuil fixé avant mesure est **5 % sur trois paires partagées** face à
`pboreads-final`, avec égalité des quatre parties complètes.

Le bootstrap Go et les **douze contrôles codecs/workers/cache** passent. Les
vecteurs avant/après donnent la même empreinte
`08995d9410ca41459f341113acbe9d544865b67135ba62f50ecfc0a0ce6a74f6` ; le bridge
Go vérifie aussi toutes les unités UTF-16, les grandes chaînes et les suffixes
invalides sous race detector. Le probe partagé prend **39,454 s**, avec
**61 optimisés / 2 643 réutilisés / 6 émissions / 2 698 restaurations** et cinq
Go réécrits. Les quatre parties complètes, clés préparées et TAST sont égaux à
`pboreads-final` sur les états original/partagé.

Les trois paires confirment **45,346 → 40,791 s**, soit **−10,05 %**. Chaque
paire gagne et le seuil de 5 % est atteint : **`scalarhex` est retenu**. Le
candidat est figé dans `candidate-scalarhex`, avec seulement les quatre fichiers
du codec scalaire et son test modifiés depuis `pboreads-final`. La compilation
native est conservée dans `gopurs-native-build-V7HgZQ`.
`validate-scalarhex.py` termine le froid complet, les paires actuelles
face au lot 03, la parité b8x des trois hôtes sur les quatre états, le bootstrap
Rust privé, les transitions des quatre hôtes, les replays stricts, la CLI et les
audits. Le stockage commun est byte-identique aux sources déjà testées sous race
detector macOS/Linux pour `pboreads`. **La validation complète passe ; le lot 04
reste non validé pour les performances.**

Les dix passages Go parallèles sont réussis, avec
**131,050 s froid / 1,982 s identique / 37,781 s feuille / 39,396 s partagée /
27,924 s FFI**. Les mutations réécrivent **1 / 5 / 1 Go** et les builds forcés
retrouvent les mêmes parties/sorties, sans réécriture. Les trois paires actuelles
face au lot 03 donnent **37,404 → 40,019 s (+6,99 %)** sur le partagé : le coût
supplémentaire se réduit encore, sans établir de gain global sur ce scénario.
Les campagnes JS et Go séquentiel sont terminées ; la parité des trois hôtes est
byte-exacte sur les quatre états, preuves v2 comprises. Le bootstrap Rust privé
passe, avec SHA-256
`0a3d9cbcc0b25fd8d6b83128da9c9241bd417c0c01d196549b14575d953d9bef`.
Les **treize transitions des quatre hôtes** passent dans `gopurs-tests-S1crHw`,
avec ce Rust privé sans PGO. Les six replays stricts comptent **35 commandes
réussies** ; les **19 contrôles CLI** et l'audit de **411 fichiers** passent.
Les trois snapshots examinés restent byte-identiques. Les pointeurs et comptes
finaux sont dans `validation-scalarhex-summary.json` et
`validation-scalarhex-counts.json`. Le diagnostic d'allocations suivant relie
le même harnais aux codecs du package natif `scalarhex` conservé.

Ce diagnostic confirme **74,258 → 47,813 Go alloués (−35,61 %)** pour le décodage
des quatre parties. Les décompositions JSON et lectures des références restent
coûteuses. L'essai suivant, `nodeviews`, remplace les nœuds JSON bruts du contexte
de décodage par des vues typées et des arêtes `Int` déjà vérifiées. Les scalaires
et noms non consommés gardent leur comportement antérieur ; les arêtes vers
l'avant, cycliques ou non canoniques restent rejetées avant décodage métier.
Le wire et les bridges restent identiques.

`nodeviews-before.log` conserve le reproducteur rouge : **8 195 lectures de
champs JSON bruts** pour **4 096 références au même scalaire**. Le nouveau
décodeur doit analyser cette forme une seule fois. Le JS précédent est conservé
dans `nodeviews-before-output` pour la comparaison différentielle des vecteurs
valides/malformés et de leurs réencodages. Le seuil annoncé avant mesure est
**10 % sur trois paires partagées face à `scalarhex`**. Bootstrap, tests ciblés,
comparaisons des quatre artefacts et paires sont lancés par `build-nodeviews.py`.
Le reproducteur passe à **quatre lectures brutes**. Le différentiel compare
**4 136 vecteurs × 15 codecs = 62 040 résultats**, dont 372 acceptés, avec les
mêmes acceptations/refus et réencodages. Les quatorze contrôles ciblés, le
bootstrap Go et la parité b8x original/partagé passent. Cependant les trois
paires donnent **41,029 → 41,001 s (−0,07 %)** : le seuil de 10 % n'est pas
atteint. **`nodeviews` est rejeté**, candidat et preuves conservés.

Les deux fichiers expérimentaux sont rétablis depuis `scalarhex`, le JS est
reconstruit et le natif restauré atomiquement. `restore-scalarhex-audit.json`
confirme les **411 fichiers** identiques au candidat pleinement validé. Ce
retour conserve l'optimisation `scalarhex` et retire seulement `nodeviews`.

L'essai suivant, `inttables`, cible les tables `Map Int ExprType` et
`Map Int NeutralExpr` du décodeur. Le profil attribue **4,87 Gio d'allocations
propres** au wrapper `Data_Ord_OrdIntImpl`. Les quatre insertions/lookups de ces
tables passent maintenant par les helpers `NativeMaps` existants à comparaison
entière directe ; le JS conserve `Data.Map` avec `Ord Int`. Seul `Wire.purs`
change, sans nouveau bridge ni changement des octets wire. Le différentiel de
référence est conservé dans `inttables-before-differential.json`. Le seuil
préalable est **5 % sur trois paires partagées face à `scalarhex`**. Le bootstrap,
les contrôles, le différentiel et le probe byte-exact passent. Les trois paires
donnent cependant **41,110 → 40,348 s (−1,85 %)** : **`inttables` est rejeté**.
Le patch inverse rétablit `Wire.purs`, puis `restore-after-inttables.py`
reconstruit le JS et restaure le natif atomiquement. L'audit
`restore-after-inttables-audit.json` confirme les **411 fichiers** identiques au
candidat `scalarhex` validé. Le candidat et les mesures de l'essai sont conservés.

Le diagnostic suivant mesure les copies internes au staging sur les **13 813
fichiers / 544 896 118 octets** d'un manifeste réel, avec huit tâches. Les trois
paires donnent **1,762 s de médiane par copie contre 3,205 s par lien physique**.
Les octets sont identiques et les fichiers privés indépendants du cache vivant,
mais cette piste est écartée avant modification du pilote. Ces mesures isolées
ne sont pas des temps de build.

`profile-prepared-keys.py` relie ensuite un harnais au package natif `scalarhex`
validé pour profiler l'encodage des modules réellement préparés et des métadonnées
globales. Chaque empreinte de module doit correspondre à la campagne b8x validée.
La préparation précède le profil d'allocations ; ce diagnostic est distinct des
paires de performance. Le harnais termine avec les **2 704 empreintes exactes** :
**64,829 Go alloués** pour **282,025 Mo** de wire des modules, puis **3,418 Go**
pour les **20,134 Mo** des métadonnées. `ordExprType` représente environ **41 %**
des allocations échantillonnées des modules (coût cumulé, non additionnable aux
autres lignes du profil).

L'essai `typeintern` conserve le raccourci d'identité `TypeMemo`, puis remplace
la table `Map ExprType Int` par l'internage des références canoniques de types
dans un second `Pool`. Il évite les comparaisons structurelles récursives tout
en préservant l'ordre d'encodage des enfants et la localité des tables. L'ancien
`typepool` rejeté n'avait pas ce raccourci d'identité et redéployait les sous-types
partagés. Les bridges existants sont réutilisés ; seuls `Wire.purs` et son test
changent. Les octets wire et les clés doivent rester identiques.

Le reproducteur préalable `typeintern-before.log` est rouge sur **44 594
comparaisons structurelles**, pour 27 racines et 1 774 octets wire. Les vecteurs
comparent les 17 constructeurs, les types égaux-distincts/partagés, l'ordre des
indices et l'isolation des artefacts. `build-typeintern.py` termine les vérifications,
le bootstrap Go, le probe byte-exact et trois paires partagées. Le reproducteur
passe à **zéro comparaison structurelle**, avec les mêmes 1 774 octets et SHA-256
`227056a7cb411debe5935a6f1f0748e89c0f2e307d7ed5ff05beba4f91cd527e`.
Les **1 007 descripteurs de types / 17 constructeurs**, leurs variantes de
partage/ordre et allers-retours, les **62 040 comparaisons différentielles** et
les **13 contrôles ciblés** passent. Le bootstrap est conservé dans
`gopurs-native-build-zPeOcl` ; les quatre parties complètes, TAST, empreintes
préparées et Go du probe original/partagé restent byte-exacts.

Les trois paires confirment **43,923 → 40,838 s**, soit **−7,03 % / −3,086 s**,
avec un gain dans chaque paire et cinq réécritures à chaque bascule. Le calcul
des clés passe de **8,119 à 5,341 s de médiane**. Le seuil préalable de 5 % est
atteint : **`typeintern` est retenu**, figé dans `candidate-typeintern`.
`validate-typeintern.py` termine ensuite la campagne réellement froide complète,
les paires actuelles face au lot 03, la parité b8x des trois hôtes sur les quatre
états, le bootstrap Rust privé, les treize transitions des quatre hôtes, les
replays stricts, la CLI et les audits. Le stockage est byte-identique à la version
déjà testée sous race detector macOS/Linux. **La validation complète passe le
12 octobre 2026 ; le lot 04 reste ouvert pour le bilan de performances.**

Les dix passages Go parallèles sont réussis : **127,687 s froid / 1,900 s
identique / 38,493 s feuille / 37,654 s partagée / 29,235 s FFI**. Les mutations
réécrivent **1 / 5 / 1 Go** ; les builds forcés retrouvent les mêmes sorties et
parties, sans réécriture. La feuille réutilise **2 703 optimisations/émissions** ;
le partagé réutilise **2 643 optimisations / 2 698 émissions** ; la FFI réutilise
**2 704 optimisations / 2 703 émissions**. Les quatre parties complètes, les
empreintes préparées et les TAST restent égaux à `scalarhex-final` sur les
quatre états. La parité JS/Go séquentiel/Go parallèle retrouve les **3 010 Go**
byte-exacts par état.

Les trois paires actuelles face au lot 03 donnent **39,273 → 38,480 s**, soit
**−2,02 %** sur le partagé. Ce faible gain remplace la régression observée lors
des paires précédentes, mais le froid reste à **127,687 s** contre les
**93,599 s historiques** du lot 03. Sa publication prend **66,593 s** ; elle
produit **10 992 objets / 548 620 453 octets**, puis le cache atteint **11 016
objets / 566 151 451 octets** après les mutations. Ces passages froids ne sont
pas une nouvelle comparaison par paires. Le coût de publication reste à traiter.

Le bootstrap Rust privé sans PGO est identifié par SHA-256
`544c15ba0c2b7a83151a25167d5a774638212513d3bea062dda6819c025efbb3`.
Les **treize transitions des quatre hôtes** passent dans `gopurs-tests-MDbt3V`,
avec ce binaire Rust. Les six replays stricts comptent **35 commandes réussies**,
les **19 contrôles CLI** passent et l'audit retrouve les **411 fichiers** du
candidat figé. Les trois snapshots examinés sont inchangés. Les pointeurs,
comptes et statuts sont dans `validation-typeintern-summary.json`,
`validation-typeintern-counts.json` et `validation-typeintern-records.json` ;
`typeintern-decision.json` porte désormais `validationComplete: true`.

Le diagnostic `profile-storage-typeintern.py` rejoue ensuite la publication des
octets validés dans une racine privée neuve, avec une copie instrumentée du
stockage. Les sorties, mémos et objets sont exacts, puis le lookup réussit. Sur
**65,919 s** de publication isolée, le staging prend **38,676 s** et les écritures
publiques **27,032 s**. Ces dernières comprennent **12,041 s** de synchronisation
des fichiers et **11,441 s** des répertoires. Les temps cumulés des workers du
staging se recouvrent ; ce diagnostic ne constitue pas une paire de performances.

L'essai privé `publicpool2` reprend la publication publique à huit tâches sur
cette référence qui intègre `batchsync`. Chaque fichier conserve `atomicWrite`,
sa synchronisation et celle du répertoire ; toutes les tâches sont jointes avant
retour, même sur erreur, et avant suppression/commit. Le protocole et les bridges
restent identiques. L'ancien `publishpool`, rejeté à **−5,87 %** avant `batchsync`,
garde sa décision. Les seuils du nouvel essai sont fixés à **10 % sur trois paires
de publication froide**, puis **5 % sur trois paires de build froid complet**.
`measure-publicpool2.py` termine les tests macOS sous race detector et les trois
paires de stockage : **72,155 → 59,773 s de médiane (−17,16 %)**, avec un gain dans
chaque paire. Les **3 010 sorties / 2 704 modules / 10 991 objets de contenu
uniques** sont byte-exacts et les six lookups réussissent. Le premier seuil est
atteint ; la confirmation est exécutée avec un candidat privé.
`build-publicpool2-cold.py` termine le stockage Linux sous race detector, le relien
depuis les sources Go du bootstrap `typeintern`, puis trois paires réellement
froides. L'audit confirme une seule différence de code Go,
`gopurs_build_cache/store.go`. Chaque passage contrôle aussi la relance identique
et une édition feuille, partagée ou FFI. Go **1.26.8 linux/arm64** exécute la suite
de stockage avec l'image b8x identifiée par digest.

Les paires froides complètes donnent **127,163 → 120,354 s de médiane**, soit
**−5,35 % / −6,809 s**, avec un gain dans chaque paire. Les six passages démarrent
sans cache ni Go généré et réécrivent **3 010 Go**. Les hits identiques ne
réécrivent rien ; les éditions réécrivent **1 / 5 / 1 Go**, avec les mêmes TAST,
clés préparées, quatre parties et sorties que `typeintern-final`. Le seuil
préalable de 5 % est atteint : **`publicpool2` est retenu pour validation complète**.

`install-publicpool2.py` vérifie les **411 fichiers**, applique le natif par
renommage atomique et confirme les sources/outils face au candidat figé. Seuls
`tools/build-cache/store.go`, `tools/build-cache/stage_test.go` et le natif
diffèrent de `typeintern` ; le SHA-256 natif est
`b152c19bca501a71d0a0f3749728b1342faee8c7fa99bccfffb80804f6a05c43`.
`validate-publicpool2.py` lance le froid complet avec builds forcés, trois paires
actuelles face au lot 03 **pour chacun des scénarios feuille/partagé/FFI**, la
parité b8x des trois hôtes sur les quatre états, le bootstrap Rust privé, les
treize transitions des quatre hôtes, les replays stricts, la CLI et les audits.
Le stockage audité reprend ses preuves macOS/Linux. **La validation complète
passe le 12 octobre ; le lot 04 reste ouvert pour les performances.**

La campagne Go parallèle est terminée : **121,921 s froid / 1,785 s identique /
41,478 s feuille / 43,920 s partagée / 29,764 s FFI**. Ce sont des passages
uniques. Les mutations réécrivent **1 / 5 / 1 Go** ; les trois contrôles forcés
retrouvent les mêmes octets sans réécriture. Les quatre parties métier et les
empreintes préparées restent égales à `typeintern-final`. Les trois paires
actuelles face au lot 03 donnent **42,186 → 40,256 s (−4,57 %)** sur la feuille
et **43,255 → 44,457 s (+2,78 %)** sur le partagé : le gain ciblé au froid ne
suffit donc pas à établir un gain global significatif.

La première comparaison FFI s'arrête sur une assertion du harnais : elle exige
du lot 03 les compteurs de réutilisation ajoutés au lot 04. Le build concerné a
réussi et ses **3 010 Go** sont byte-exacts ; son inventaire de mémos est vide,
conformément à cette référence. L'essai et les scripts antérieurs sont conservés.
La correction fige la capacité attendue depuis le probe initial de chaque
compilateur, impose l'inventaire correspondant à chaque passage et conserve les
compteurs stricts pour `publicpool2`, ainsi que la vérification du seul chemin FFI
réécrit pour les deux variantes. `audit-publicpool2-ffi-recovery.py` retrouve les
**411 fichiers** inchangés et contrôle les résultats déjà réussis. La validation
reprend avec `paired-ffi-lot03-publicpool2-retry1.log` et un nouveau répertoire de
mesures ; `publicpool2-ffi-recovery.json` consigne la reprise auditée.

La reprise FFI réussit les trois paires : **41,195 → 30,960 s (−24,85 %)** face
au lot 03. Les campagnes b8x JS/Go séquentiel/Go parallèle retrouvent les mêmes
TAST, empreintes préparées, quatre parties et **3 010 Go** sur les quatre états.
Le bootstrap Rust privé sans PGO est réellement construit puis utilisé dans les
treize transitions des quatre hôtes (`gopurs-tests-JLVF5c`) ; SHA-256
`8f9d04580b0882550abdf13c5adbff95bec37655211f6102af989e480c16bdd3`.
Les **35 commandes strictes**, les **19 contrôles CLI** et l'audit des **411
fichiers** passent. Les trois snapshots examinés sont inchangés.
`validation-publicpool2-{summary,counts,records}.json` conserve les preuves,
y compris l'échec initial du harnais et sa reprise ; `publicpool2-decision.json`
porte `validationComplete: true`. `publicpool2` devient la référence entièrement
validée pour la suite des essais.

Le profil des clés est ensuite rejoué avec le paquet Go validé de `publicpool2`
et les points d'entrée réexaminés. Il retrouve les **2 704 clés**, les
**282 024 665 octets wire des modules** et les **20 134 333 octets des métadonnées**,
avec le même digest de métadonnées. Face au profil `scalarhex`, l'encodeur
`typeintern` réduit les allocations des modules de **64,829 à 45,466 Go
(−29,87 %)** et celles des métadonnées de **3,418 à 1,259 Go (−63,18 %)**. Les
allocations échantillonnées attribuent encore **17,85 %** au chemin de suppression
des champs record (`Record.Unsafe.unsafeDelete`, coût cumulé incluant
`RecordToMap`). Les temps de ce profil séquentiel instrumenté ne sont pas des
paires de performances. `prepared-keys-publicpool2-{summary,input-audit}.json`
et les profils complets conservent ces mesures.

L'essai `recordfields` ajoute une traversée typée des champs qui lit le record
initial sans le copier après chaque champ. L'ordre canonique et le décodeur
restent ceux du codec ; les bridges existants servent les lectures. Le seuil
préalable est **5 % sur trois paires partagées face à `publicpool2-final`**.
Le reproducteur observe **72 → 0 énumérations/copies des records d'entrée**,
avec les mêmes **5 891 octets** et le SHA-256
`d1dcfc64da2303b28b8eab428529bc621f5ae5a0ccce42b262c95d9066759de3`.
Les **200 vecteurs de records**, les **62 040 comparaisons différentielles** et
les **14 contrôles ciblés** passent. Le bootstrap Go réussit dans
`gopurs-native-build-eVHfKt` (**510 modules / 309 930 types**).
Le harnais s'arrête ensuite sur un libellé de journal obsolète ; la reprise
vérifie les **425 fichiers**, les cinq commandes réussies et l'identité du
binaire installé avec celui du bootstrap avant de reprendre les probes et
paires. `recordfields-workspace-recovery.json` conserve l'audit et le script
antérieur. Les probes retrouvent les mêmes TAST, clés, quatre parties et Go.
Les trois paires donnent pourtant **37,437 → 38,824 s (+3,71 %)** sur le build
partagé, avec une régression dans chaque paire. Les clés préparées passent de
**4,572 à 4,207 s** de médiane, mais ce gain local ne suffit pas. **L'essai est
rejeté.** Le retour au codec et aux binaires `publicpool2` est exact sur les
**425 fichiers**, avec les tests de sources du lot 05 préservés ; le bundle est
reconstruit à l'identique et les 13 contrôles ciblés passent.
`recordfields-rejection-audit-{before,after}.json` et `recordfields-decision.json`
conservent le rejet ; le candidat et tous les résultats restent disponibles.

L'essai suivant `backrefs` porte sur les tables du décodeur. Les deux
`Map Int` des types et expressions deviennent des tables indexées locales,
avec vues bornées au préfixe déjà décodé. De petits bridges JS/Go/Rust conservent
les valeurs opaques et appliquent les bornes ; la validation du schéma reste en
PureScript. Les types invalides inutilisés rejettent toujours l'artefact et les
expressions invalides restent absentes. Le seuil préalable est **5 % sur trois
paires partagées face à `publicpool2-final`**. Le reproducteur de référence
compte **411 comparaisons de types / 287 comparaisons d'expressions** pour
**33 types / 100 nœuds d'expression**, dont un partage de profondeur 24. Les
**8 895 octets** et le SHA-256
`50cf809afaf3b7709aee67e74615584e4362e113f7b1eacd9e81afb0957ca057`
sont conservés avant changement. Le reproducteur après changement observe zéro
comparaison, avec les mêmes octets ; les **1 007 descripteurs de types**, les
**62 040 cas différentiels**, les **16 contrôles ciblés** et le bootstrap Go
passent. Le premier bootstrap Rust échoue sur `RefCell` en mode multithread ;
le remplacement par des accès courts `Mutex`, sans verrou pendant les callbacks,
permet le bootstrap réel et son smoke test.

La revue du Rust généré révèle ensuite que le retour fonctionnel de `types` et
`expressions` reconstruit la table à chaque lecture. Le reproducteur exécuté sur
ce crate compte **256 callbacks au lieu de 32 pour huit lectures**, dans chacune
des deux tables. La tentative de paires est interrompue et conservée après audit
des **430 fichiers** ; elle ne sert pas à une décision de performances.
Le correctif `backrefs-table` retourne une table opaque, conservée dans le
contexte puis lue par `lookup`. Les bridges de stockage conviennent déjà à cette
interface. La nouvelle campagne réussit les **17 contrôles ciblés**, les deux
bootstraps et le reproducteur Rust (**256 → 32 callbacks**). Les probes
retrouvent les mêmes TAST, clés, quatre parties et Go.
Les trois paires donnent **35,982 → 34,995 s (−2,74 %)**, sous le seuil fixé de
5 % : **essai rejeté**. Deux paires gagnent, la deuxième régresse de 8,49 % ;
la baisse locale du producteur PBO (**11,421 → 9,963 s** de médiane) ne suffit
pas à justifier la conservation. Le retour au codec et aux binaires `publicpool2`
est exact sur les **425 fichiers**, tests de sources du lot 05 compris ; le
bundle est reconstruit à l'identique et les 13 contrôles ciblés passent.
`backrefs-table-decision.json` et `backrefs-table-rejection-audit-after.json`
conservent la décision et l'audit. Les candidats et essais restent disponibles.

L'essai privé suivant `publicsync` regroupe les synchronisations de répertoires
des sorties publiques. `writeOutputs` conserve chaque écriture temporaire, la
synchronisation du fichier et le renommage, puis joint les huit writers avant
une barrière par répertoire et ancêtre, de bas en haut. Cette barrière couvre
aussi les sorties identiques d'une reprise et précède suppression et commit.
Les bridges et les octets wire restent identiques. Le reproducteur passe de
**34 synchronisations de répertoire à 4** pour 34 sorties ; le chemin de reprise
retrouve les quatre barrières même sans réécriture. Une erreur de barrière garde
le pointeur précédent, les sorties obsolètes et le journal d'union ; la reprise
restaure une sortie perdue sans réécrire les autres et conserve l'inode du verrou.
Les tests du reproducteur passent sous race detector. Seuils préalables face à
`publicpool2` : **10 % sur trois paires de publication froide isolée**, puis
**5 % sur trois paires froides complètes**. Les trois paires de stockage donnent
**59,636 → 50,999 s (−14,48 %)**, avec un gain dans chaque paire, les **3 010 Go /
2 704 modules / 10 991 objets** exacts et les lookups réussis. Le race detector
macOS passe aussi sur le stockage complet. Le SIGKILL avant la première barrière
est ensuite éprouvé : récupération de 33 sorties, fichiers survivants et inode
du verrou conservés. Les tests sous race detector passent avec **Go 1.26.8
linux/arm64**. Le relien natif ne change que `gopurs_build_cache/store.go` dans
les sources générées. Les trois paires froides complètes donnent **118,971 →
112,857 s (−5,14 %)**, avec un gain dans chaque paire, les quatre parties et Go
exacts, et **0 / 1 / 5 / 1** réécritures pour identique/feuille/partagé/FFI.
Les trois paires de non-régression donnent **34,587 → 34,341 s (−0,71 %) feuille /
38,374 → 36,384 s (−5,19 %) partagé / 25,228 → 25,659 s (+1,71 %) FFI**. Le budget
de 5 % passe dans les trois cas. L'identique passe de **1,658 à 1,612 s**, dans
son budget de `max(5 %, 0,100 s)`. Le candidat mesuré est installé atomiquement
après audit ; **430 fichiers** identifient les sources/outils/fixtures/binaires.
Le natif a pour SHA-256
`841646465c860666e56e37e5064c9db734c31c1570b1707d601286a208739462`.

Les reproducteurs sont intégrés dans `tools/build-cache-output.test.mjs` et
`tools/build-cache-fixtures/`, puis ajoutés à `npm run test:cache`. Ils
instrumentent des copies privées du vrai stockage pour observer les barrières
et provoquer l'échec/SIGKILL. `validate-publicsync.py` réussit leur validation
intégrée, les bridges, la campagne b8x, les paires actuelles face au lot 03,
la parité des trois hôtes sur les mêmes TAST, le bootstrap Rust, les transitions
pilote/sources/contexte dans les quatre hôtes et les replays stricts. Les trois
snapshots examinés sont inchangés. **La validation complète de `publicsync`
est réussie** ; les paires froides actuelles face au lot 03 complètent le bilan.

Les barrières intégrées passent dans `gopurs-cache-output-Oek4EG`, ainsi que les
quatre contrôles du stockage, des bridges et de l'identité du compilateur. La
campagne native complète réussit, avec **116,835 s froid / 1,739 s identique /
33,888 s feuille / 35,994 s partagé / 25,625 s FFI**, et contrôles forcés à octets exacts sans
réécriture. Les trois paires actuelles face au lot 03 donnent :

| Édition — temps gopurs | Lot 03 | `publicsync` | Gain |
| --- | ---: | ---: | ---: |
| Feuille | 37,958 s | 34,020 s | 10,37 % |
| Partagée | 37,832 s | 36,138 s | 4,48 % |
| FFI | 37,722 s | 25,364 s | 32,76 % |

Les campagnes JS, Go séquentiel et Go parallèle passent sur les quatre états,
avec mêmes TAST, clés, parties et Go. Le bootstrap Rust privé sans PGO réussit ;
son SHA-256 est
`155d915ca48698e7546788be0e79c0c825aff66cca42b7fc5f195ab50ac995c4`.
Les transitions du pilote passent dans les quatre hôtes (`gopurs-tests-YyIejA`).
Les sources passent dans `gopurs-tests-9VGNGK` (**56 comparaisons / 294 commandes**)
et le contexte dans `gopurs-tests-7uwezQ` (**80 comparaisons / 401 commandes**),
avec variantes de compilateur/runtime réellement reconstruites, mêmes TAST,
quatre parties, clés et Go, hits sans réécriture et oracles exécutés.

Les **35 commandes de replays stricts**, les **19 contrôles CLI** et l'audit final
des **430 fichiers** passent. `validation-publicsync-summary.json` et
`validation-publicsync-counts.json` identifient ces résultats ;
`publicsync-decision.json` porte `validationComplete: true`. Le gain partagé
reste modeste. `measure-publicsync-lot03-cold.py` termine les trois paires froides
actuelles face au lot 03, avec caches et sorties neufs : **94,191 → 113,213 s
(+20,19 %)**. L'identique suivant ces froids passe de **1,366 à 1,716 s** et les
objets de **135 326 356 à 548 673 418 octets**. Les six échantillons gardent les
3 010 Go et 2 704 TAST exacts ; l'audit final des compilateurs et des sources est
inchangé. `paired-cold-lot03-publicsync/summary.json` conserve les mesures et les
phases ; le [bilan](build-incremental.md) détaille les limites et la dispersion.

L'utilisateur confirme la poursuite avec `publicsync` après explicitation du
surcoût froid. Ce candidat fixe la fin des essais d'optimisation du lot 04 ;
les paires froides actuelles et leur bilan étant terminés, le lot est clos.
Le retour à environ 30 s à froid n'est pas un résultat établi.

Les validations complètes ci-dessous conservent les preuves historiques de `allhit`.

### Vérifications déjà terminées sur le candidat allhit

- Treize transitions comparées à une génération fraîche dans les quatre modes
  **JS, Go séquentiel, Go parallèle et Rust**, sur les mêmes TAST. Les inventaires
  Go, les quatre artefacts métier et les empreintes préparées sont identiques
  entre hôtes. Le changement de retour FFI Go `int64` → `any`, à TAST inchangés,
  modifie le module, son bridge et le consommateur `Extra.go` ; le programme
  compilé s'exécute. La corruption d'un backend interdit le raccourci global et
  déclenche sa réparation par le builder normal.
- Bootstrap Rust reconstruit avec `GOPURS_RUST_PGO=0`, archive Go de support
  unique, compilation et exécution du programme Go de contrôle réussies. Le
  binaire privé est identifié par `gopurs-rust-allhit.json`.
- Les trois snapshots ont été réexaminés sans mise à jour. Les **six replays
  stricts** de `FFIIntegerReturns` et `NativeRecordWorkers` passent en JS,
  natif séquentiel et natif parallèle : inventaires Go byte-exacts, comparaison
  stricte après `gofmt`, compilation et exécution.
- **19 contrôles CLI** passent, y compris les erreurs de TAST, FFI et écritures,
  avec conservation des diagnostics et des sorties engagées.
- Stockage sous race detector : macOS après ajout de `sameTast`, puis campagne
  Linux/arm64 du code courant avec le Go 1.26.8 de l'image b8x.

### Preuves locales

`<tmp>/gopurs-module-cache-lot04-nj_1suxg/`
conserve les candidats et leurs SHA-256, les reproducteurs rouges, les profils,
les essais écartés et les workspaces isolés. Voir notamment `candidate-allhit.json`,
`allhit-before.log`, `driver-allhit-four-hosts.log`, `paired-ffi-allhit/`,
`snapshot-review-inputproof.json`, `strict-replays-allhit/records.json`,
`cli-allhit.log` et `storage-allhit-linux-race.log`. Le `README.md` retrace les
transformations et distingue les probes avec objets préchargés des caches froids.
La campagne multi-hôtes conserve son workspace dans
`<tmp>/gopurs-tests-wEqaby/`.
`allhit-source-audit.json` vérifie que 221 fichiers des sources gopurs/PBO et du
stockage correspondent au candidat figé testé.
`final-campaign-allhit.json` identifie les quatre campagnes successives ;
`probe-b8x-allhit-final/`, `probe-b8x-lot03-final/`, `probe-b8x-allhit-js/` et
`probe-b8x-allhit-sequential/` conservent les chronos, compteurs et inventaires.

<a id="plan-v3--lot-05--transitions-de-sources-campagne-en-cours-12-octobre-2026"></a>

## Plan v3 — lot 05 : correction incrémental/complet, validée (12 octobre 2026)

**Clôturé sur `publicsync`.** La campagne finale décrite ci-dessous satisfait la
[matrice de couverture](build-incremental.md#validation-de-correction--lot-05),
avec comparaisons incrémental/frais et multi-hôtes, compilation/exécution des
oracles, hits sans réécriture et audit du candidat livré.

Le complément `tests/host/CacheTransitions/`, construit par
`tools/cache-source-variants.mjs`, fait évoluer un programme sur **14 états** :
initial, corps privé/exporté, directive d'inlining, constructeur, record, type de
champ, dictionnaire, demandes transitives ajoutées/remplacées/retirées, module
renommé/restauré et retour global. Chaque état TAST est construit une fois puis
consommé par les hôtes. Les layouts traversent des appels conservés avec des
directives `never` ; les données d'exécution proviennent d'une `Effect.Ref`.

La campagne privée `lot05-source-transitions/run-4uQJYd` passe **56 comparaisons**
dans **JS, Go séquentiel, Go parallèle et Rust** : inventaires et octets Go,
quatre parties du cache, empreintes préparées, reconstruction dans une sortie
neuve, hit identique sans réécriture et compilation/exécution de chaque oracle.
Les fournisseurs `Cache.Poly` et `Cache.Relay` conservent leurs TAST lors des
changements d'appels, mais leurs clés et sorties changent ; le retrait retrouve
les clés initiales. Le renommage retire les anciennes sorties du module, de sa
FFI et de son entrée secondaire. Les changements de layouts invalident
conservativement les 66 émissions, avec sorties égales au frais.

Le premier essai `run-hKAAde` est conservé : l'inventaire textuel échouait sur
l'alias macOS `main/main.go` / `Main/main.go`. L'audit retrouve les **93 fichiers
Go exacts** et un inode distinct pour chaque sortie possédée. Le harnais exige
une bijection manifeste/fichiers physiques, avec hashes et tailles contrôlés ;
les inventaires physiques restent comparés entre hôtes et au build frais.

Le parcours est intégré dans `tools/driver-source-cache.test.mjs` et lancé par
`npm run test:cache-sources`. Les 14 variantes et oracles du dépôt sont exacts
par rapport au reproducteur privé. Le contrôle d'intégration des quatre hôtes
passe avec `publicpool2` et son Rust privé validé dans `gopurs-tests-S2M1Cn` :
**294 commandes réussies**, dont **56 reconstructions fraîches, 56 hits sans
réécriture et 56 compilations/exécutions**. L'audit retrouve les sources du test
et des fixtures exactes, ainsi que les binaires figés et les objets finaux du
cache. `lot05-source-transitions/integration-publicpool2.json` fixe les preuves.
`GOPURS_TEST_MODES` permet
de choisir les hôtes ; `GOPURS_TEST_COMPILER` et `GOPURS_TEST_RUST_BINARY` fixent
les exécutables. `GOPURS_TEST_KEEP_WORKSPACE=1` conserve sources, TAST, commandes,
résultats, binaires des oracles et caches finaux ; un échec conserve toujours
les preuves. Cette intégration est la référence historique `publicpool2` ; le
rejeu final `publicsync` est décrit en fin de section.

Le complément `lot05-context-transitions/` prépare maintenant **20 états ×
4 hôtes** : limite de réécriture, workers de préparation/PBO/émission, pipeline,
racine FFI, entrées, version de compilateur, runtime et retours à l'état initial.
Le runtime de référence est embarqué dans des copies privées JS/Go/Rust ; une
variante y ajoute un commentaire distinctif, avec reconstructions réelles Go et
Rust. Le harnais exige les miss de contexte, les mêmes sorties au frais, les
hits sans réécriture, la propriété exacte des sorties et les oracles exécutables.
`build-lot05-context.py` lance cette campagne sur les versions validées
`typeintern`/`publicpool2` et le runtime privé. Les constructions privées Go/Rust
réussissent. La première tentative valide 13 états JS puis s'arrête sur une
attente erronée du harnais : `--main` conserve l'entrée du module choisi ainsi
que l'entrée commune. L'audit retrouve les **92 Go attendus**, les **67 commandes
réussies** et les **425 fichiers de production inchangés**. La tentative est
conservée dans `run-5U2bRd` ; le retry exige les deux chemins et reprend dans une
nouvelle racine. La reprise `run-OfOWca` réussit les **80 comparaisons / 401
commandes** : quatre parties, clés et Go exacts au frais et entre les quatre
hôtes, hits sans réécriture et oracles exécutés. Le changement de runtime et
son retour réécrivent uniquement `gopurs_runtime/runtime.go` ; les changements
d'options et de version de compilateur ne réécrivent aucun Go identique.
Les **425 fichiers de production restent inchangés**.

Le test est maintenant intégré dans `tools/driver-context-cache.test.mjs` :

```sh
npm run test:cache-context
```

Le parcours courant couvre **16 états par hôte** dans JS/Go séquentiel/Go
parallèle. `GOPURS_TEST_MODES`, `GOPURS_TEST_COMPILER`, `GOPURS_TEST_RUST_BINARY`
et `GOPURS_TEST_KEEP_WORKSPACE` suivent les conventions du test de sources.
`GOPURS_TEST_RUST=1` ajoute le quatrième hôte. Une erreur conserve son workspace,
les TAST, les résultats, les commandes, les caches et les exécutables des oracles.

Pour éprouver aussi deux versions de compilateur et le runtime, construire le
catalogue avec `tools/cache-context-runtime.mjs`. Sa configuration JSON indique
deux compilateurs figés et les sorties retenues des bootstraps Go/Rust de la
version `baseline` :

```json
{
  "baseline": { "compiler": "/compiler-a/bin/gopurs", "rust": "/compiler-a/gopurs-rust" },
  "alternate": { "compiler": "/compiler-b/bin/gopurs", "rust": "/compiler-b/gopurs-rust" },
  "nativeOutput": "/retained-go-bootstrap/output",
  "rustOutput": "/retained-rust-bootstrap/rust",
  "destination": "/new-private-runtime-variants"
}
```

```sh
node tools/cache-context-runtime.mjs context-variants.json
GOPURS_TEST_CONTEXT_VARIANTS=/new-private-runtime-variants/variants.json \
GOPURS_TEST_MODES=js,native-sequential,native-parallel,rust \
GOPURS_TEST_KEEP_WORKSPACE=1 npm run test:cache-context
```

L'outil copie les ressources et les sources générées, substitue uniquement le
runtime embarqué, reconstruit les binaires Go/Rust et fige les hashes des trois
variantes. Le test étend alors la matrice à **20 états par hôte** et contrôle les
changements effectifs d'identité. La campagne d'intégration reconstruit ces
variantes avec l'outil du dépôt et rejoue les quatre hôtes, puis le parcours
sans catalogue en JS. Cette intégration réussit dans `gopurs-tests-v2nL09`
(**80 comparaisons / 401 commandes**, quatre hôtes), puis `gopurs-tests-O8cZFe`
(**16 comparaisons / 81 commandes**, JS sans catalogue). Les variantes sont
construites par l'outil du dépôt ; les 96 oracles sont compilés et exécutés.

L'audit croisé final a d'abord exigé à tort l'égalité brute des TAST avec la
campagne privée, située à une autre profondeur. Les **57 différences** se limitent
aux `modulePath` relatifs des bibliothèques : leurs chemins résolus, leurs sources
et tous les autres champs TAST sont exacts. Les TAST intégrés sont byte-exacts
avec l'état initial de `gopurs-tests-S2M1Cn`, ainsi que toutes leurs clés préparées.
Les parties et Go restent byte-exacts avec la campagne privée. Le validateur
initial et son échec sont conservés ; `audit-integration-recovery.py` termine
l'audit, y compris les objets finaux et les **427 fichiers** sources/outils/FFI/
binaires. `integration-completed.json` conserve les résultats et les chemins
audités.

### Validation sur le candidat de clôture `publicsync`

Le rejeu sur `publicsync` réussit les **14 états sources × 4 hôtes**, soit **56
comparaisons et 294 commandes**, dans `gopurs-tests-9VGNGK`. Le catalogue de
contexte est reconstruit avec `publicsync`, `publicpool2` et une variante runtime
privée ; les **20 états × 4 hôtes**, soit **80 comparaisons et 401 commandes**,
passent dans `gopurs-tests-7uwezQ`. Les quatre hôtes partagent les mêmes TAST ;
les résultats incrémentaux/frais ont les mêmes parties, clés et Go, avec hits
sans réécriture et oracles compilés/exécutés. Les changements de runtime et
retours réécrivent uniquement `gopurs_runtime/runtime.go`.

Les treize transitions du pilote, les contrôles de corruption/propriété/
concurrence/interruption, les barrières publiques avec SIGKILL, les **35 commandes
strictes** et les **19 contrôles CLI** passent également. L'audit final couvre
**430 fichiers**, les trois snapshots restent inchangés et Rust est réellement
reconstruit sans PGO. `publicsync-source-transitions.json`,
`publicsync-context-transitions.json` et `validation-publicsync-summary.json`
fixent les preuves. La revue `lot05-coverage-review.md` relie chaque exigence au
test livré. Le bilan du lot 04 et cette revue étant terminés, **le lot 05 est
validé ; le plan v3 passe à 6/7 lots**.

<a id="plan-v3--lot-06--mesures-du-parcours-complet-en-cours-12-octobre-2026"></a>

## Plan v3 — lot 06 : mesures du parcours complet validées (12 octobre 2026)

**Lot validé ; plan v3 clôturé à 7/7.** La campagne complète et son audit
réussissent : **24 chaînes, 24 oracles**, quatre scénarios de trois paires,
inventaires TAST/Go/modules exacts, même binaire par scénario, **430 fichiers
de production**, caches Go de base et inodes de verrou inchangés. Les trois
snapshots déjà examinés restent exacts ; les 35 replays stricts du lot 05
portaient sur ce même compilateur, sans changement de code depuis.
Le [bilan final](build-incremental.md#parcours-complet--lot-06) et les
[24 mesures individuelles](build-cycle-lot06.csv) publient les résultats.

Les compilateurs et le corpus des lots 04–05 sont figés. Le harnais privé
`measure-lot06-cycle.py` chronomètre directement `spago build` avec le backend
réel, `go mod tidy`, puis le runner b8x jusqu'à la frontière `exec Test.Main`.
Un wrapper mesure l'invocation réelle de gopurs à l'intérieur de Spago ; les
traces Go identifient compilations et liens. Les ressources restent fixées à
**GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB**.

Chaque générateur prépare son GOCACHE privé sur la base seulement ; chaque
échantillon en reçoit une copie neuve, avec les indices du binaire de base du
runner. Trois paires alternées sont prévues pour identique/feuille/partagé/FFI.
Les journals gopurs restent dans leur racine ; un vrai build restaure la base
avant chaque mutation. Les 3 010 Go, les 2 704 TAST et les fichiers de modules
sont vérifiés, puis l'oracle de la mutation est compilé et exécuté hors chrono.
Le premier lancement est conservé dans la preuve privée
`<tmp>/gopurs-cycle-lot06-publicsync/`.
Il échoue au contrôle d'inventaire : le `PATH` sélectionnait Spago **0.93.45** du
bootstrap gopurs au lieu du **1.0.3** du lot 00. Trois modules sont ajoutés et
`Spago.Generated.BuildInfo` change ; ces temps sont exclus des résultats.
`failure-audit.json`, l'ancien harnais et le workspace restent conservés.

La correction invoque le Spago du lot 00 par son chemin explicite, avec contrôle
de version et de SHA-256. Le contrôle préalable
`gopurs-lot06-frontend-preflight/` passe **7 transitions**, chacune avec les
**2 704 TAST byte-exacts** et le lockfile inchangé. La deuxième tentative dans
`gopurs-cycle-lot06-publicsync-spago103/` reproduit les **3 010 Go et 2 704 TAST**,
mais `tidy` échoue sur la résolution DNS de `proxy.golang.org` : des dépendances
de test manquaient au cache de modules copié. Ce deuxième échec et son audit
restent conservés, avec des temps exclus des résultats.

`gopurs-lot06-offline-modules/` prépare une copie privée complète des
téléchargements, vérifie `tidy`, `go mod verify` et une seconde relance `tidy`
hors ligne, puis fige l'inventaire des **3 094 fichiers**. `go.mod`, `go.sum`, les
Go et les TAST restent exacts. Les **7 transitions frontend** passent aussi avec
`spago build --offline`, dans `gopurs-lot06-frontend-offline-preflight/`.
La campagne est relancée dans `gopurs-cycle-lot06-publicsync-offline/`, avec
`GOPROXY=off` et `GOSUMDB=off` et des GOCACHE toujours construits séparément sur
la base. Le protocole précise les passages
froids isolés, leurs états de cache et la durée séparée de construction/exécution
des oracles. Les configurations Docker effectives sont conservées avec les
traces et les inventaires.

### Point intermédiaire et reprise des mesures

La campagne hors ligne valide **16 chaînes sur 24** : trois paires identiques,
trois paires feuille et deux paires partagées. Tous ces échantillons ont leurs
inventaires byte-exacts et leurs oracles réussis. Le passage `shared-3-lot03`
est rejeté : `docker exec` termine en statut 0 après **53,760 s**, mais sans
marqueur `exec` ni rapport du runner ; la trace s'arrête pendant la compilation.
Le harnais refuse correctement ce passage. La cause de l'interruption n'est
pas établie par les traces conservées ; `interruption-audit.json` en fixe les
faits. Les **430 fichiers de production** et les toolchains sont inchangés.

| Scénario complet (3 paires) | Mesure | Lot 03 | `publicsync` | Variation |
| --- | --- | ---: | ---: | ---: |
| Identique | gopurs | 1,123 s | 1,559 s | +0,436 s / +38,81 % |
| Identique | chaîne jusqu'à `exec` | 3,559 s | 3,962 s | +0,403 s / +11,33 % |
| Feuille | gopurs | 37,457 s | 35,030 s | −6,48 % |
| Feuille | chaîne jusqu'à `exec` | 247,982 s | 244,403 s | −1,44 % |

Ce sont les médianes observées, avec leur dispersion conservée : les trois
paires feuille gagnent **7,62 / 5,30 / 7,08 %** sur gopurs, mais la chaîne varie
de **−2,59 / −1,93 / +1,34 %**. La compilation Go reste autour de **207 s** ;
le faible écart de chaîne ne démontre pas un gain robuste du cycle complet.
L'identique ne réécrit aucun Go et ne compile/lie rien.

`measure-lot06-cycle.py --resume PREVIOUS NEW_DESTINATION` reprend les huit
chaînes manquantes dans `gopurs-cycle-lot06-publicsync-resume1/`. Le corps des
mesures/oracles et sept fonctions du harnais sont AST-identiques à la version
figée. Les résultats acquis sont relus avec leurs inventaires/hashes/oracles ;
les projets restent à leurs racines d'origine. Des clones forensiques de leur
état interrompu sont conservés sans être exécutés. Les seeds Go sont inventoriés
et recontrôlés sur des copies privées : zéro compilation/lien et binaire de base
exact requis avant la reprise. Chaque mesure reçoit ensuite sa propre copie du
seed. Les signaux reçus et l'état du conteneur lors d'un échec sont désormais
consignés ; le journal de lancement est conservé dans la preuve principale.
La reprise se termine avec succès. `summary.json` et `results.json` dans
`gopurs-cycle-lot06-publicsync-resume1/` couvrent les **24 échantillons** ;
`sample-evidence.json` conserve leur provenance. `analyze-lot06-cycle.py`
recontrôle chaque inventaire, oracle, action Go et hash de binaire, les audits
avant/après et les snapshots. Sa synthèse est conservée dans
`gopurs-cycle-lot06-analysis-final/`.

Les médianes finales de chaîne sont **3,559 → 3,962 s** sur l'identique,
**247,982 → 244,403 s** sur la feuille, **265,895 → 268,585 s** sur le partagé,
et **250,273 → 236,510 s** sur la FFI. Les variations sont donc **+11,33 %
(+0,403 s) / −1,44 % / +1,01 % / −5,50 %**, avec toutes les paires conservées.
La chaîne partagée ne gagne pas en médiane et la dispersion de la feuille
empêche d'attribuer un gain robuste au cycle complet. Le backend gagne
respectivement **6,48 / 1,40 / 32,70 %** après édition. Les temps Go de
**206,820 / 209,650 / 208,470 s**, soit environ **85 / 78 / 88 %** de la chaîne
livrée, confirment **G avant P**. Les observations froides de chaîne restent
identifiées comme passages uniques, et le surcoût froid backend apparié
**+20,19 %** est conservé au bilan.

Le bloc de réinitialisation publié dans `build-cache.md` est exécuté tel quel
sur une fixture privée : attente d'un `flock` réellement détenu, résolution
d'une sortie symbolique, répétition de la commande, conservation byte-exacte du
journal, des TAST, des Go et des fichiers de modules, avec dates et inode du
verrou inchangés. Preuve : `gopurs-lot06-reset-doc-check/summary.json`.

## Suite G — G00 : graphe typé et références validés (13 octobre 2026)

Le [rapport G00](build-go-packages.md) reprend le compilateur livré après v3/06.
L'analyse des **2 997 Go du package b8x** résout tous les corps avec `go/types`
et les exports réels de Go **1.26.8 linux/arm64**. La base et les trois éditions
ont chacune leur graphe typé et leurs inventaires exacts. Le bootstrap conservé
est vérifié de la même manière : **625 Go / 115 FFI**, sources exactes face à
l'audit du binaire livré.

Le graphe par module conserve **94,53 %** du code dans une seule composante
cyclique. Le graphe par déclaration relève **185 667 unités**, avec contraintes
de méthodes et champs distinguées. Les fermetures statiques des éditions
représentent **0,0006 / 6,12 / 11,48 %** des octets de déclarations, avant
regroupement, contraintes d'ABI et initialisation : ces tailles ne sont pas
des gains de temps. Les **348 tags `StructGetters`** sont distincts ; les usages
FFI interunités et les chemins réflexifs sont inventoriés.

L'audit en lecture seule vérifie les **430 fichiers de production**, les hashes
des cinq graphes, le seed Go original et les binaires `Test.Main`, puis relève
le stockage des douze échantillons `publicsync` de v3/06. Les budgets avant
candidats sont fixés à **+5 % mémoire / +10 % par catégorie de stockage**, avec
les seuils de temps déjà inscrits au plan.

La première calibration mémoire est **rejetée** : le harnais avait oublié les
bibliothèques sœurs `workspace/gopurs`, rendant certaines FFI introuvables.
L'inventaire exact échoue ; la tentative et ses stubs sont conservés, temps/RSS
exclus. Après correction et préflight des **2 704 chemins source/FFI**, le froid
passe : **113,687 s / 9,841 GiB de RSS maximal**, **3 010 Go** exacts et
**10 992 objets**.

Le harnais s'arrête ensuite avant l'identique, faute de marge disque : **7,9 GiB
libres**, sous le plancher de **8 GiB**. Le froid reste validé. Après libération
d'espace par l'utilisateur (**19,1 GiB disponibles**), la reprise termine les
six autres passages dans la même racine/journal. Les sept inventaires sont
exacts ; zéro réécriture sur l'identique, **1 / 5 / 1** sur les éditions, inode
du verrou inchangé. La commande chronométrée est AST-identique au harnais figé.
Les sept reproducteurs de contrats Go et le contrôle de réflexion passent,
dont quatre erreurs de compilation attendues. Le froid applicatif Go prend
**255,780 s**, avec deux compilations/un lien, **15 433 168 KiB** de pic RSS et
le binaire exact de v3/06. La relance identique ne compile ni ne lie ; l'oracle
passe. Ce passage unique est plus lent que l'ancien froid intégral de
**221,830 s** ; les observations ne sont pas appariées, l'écart est conservé
sans l'attribuer au préchauffage. **G00 est validé, suite G à 1/7.**

Preuve : `gopurs-go-packages-g00-20261012/`, notamment
`declaration-analysis-1/`, `bootstrap-typed-graph-1/`, `reference-audit-1/`,
`backend-memory-reference-1/rejection.json` et
`backend-memory-reference-2/pause.json`, `backend-memory-resume1/` et
`record-contracts-1/`, `go-application-cold-1/` et `final-reference/summary.json`.
L'audit final confirme de nouveau les 430 fichiers de production et les preuves
requises ; la calibration mémoire et le détail des reprises sont publiés dans
le rapport G00.

## Suite G — G01 : plan canonique validé (13 octobre 2026)

Le [plan de placement](build-go-package-plan.md) conserve les FFI entières,
les méthodes/receveurs, les cellules des getters et l'ordre d'initialisation Go.
Les identités des inscriptions sont ancrées à leur constructeur, celles des
méthodes à leur type receveur. Les intervalles topologiques de densité 1/128
donnent **525 packages b8x / 158 pour le bootstrap**. Les densités 1/256 et 1/512
sont également conservées.

Un contrôle indépendant vérifie la couverture exacte des déclarations, les
arêtes du quotient et le DAG des **douze plans b8x et trois plans du bootstrap**.
La permutation des 185 667 identifiants et de l'ordre d'énumération conserve
les empreintes de placement. Les contrats d'identité couvrent insertion d'une
inscription, offsets/corps, nom local du receveur, remplacement à cardinalité
égale, rétraction et collision forcée. **95 graphes** sont comparés à un oracle
de fermeture transitive pour les SCC.

**Sept programmes Go** comparent l'initialisation commune, la séparation naïve,
le regroupement ordonné, la perte d'une inscription, un initialiseur central
et deux cas de RHS intercalées. Les deux écarts attendus sont conservés. Les
cas corrigés préservent les effets d'un package externe importé par du code
inutilisé, les inscriptions d'une entrée partielle, et une seule cellule lazy
pour 64 appels concurrents depuis deux packages.

La tentative `initialization-contracts-3/` est rejetée : l'ordre écrit à la main
pour les RHS intercalées différait de Go. La reprise `-4` utilise
`go/types.Info.InitOrder` et passe, y compris groupe `var`, résultats multiples
et variable `_`. Une incompatibilité Python 3.9 (`int.bit_count`) est aussi
conservée dans `consumer-set-models-1/`, puis corrigée avant le relevé `-2`.

Les modèles non retenus restent archivés. Les fermetures du plan canonique
représentent **38 792 440 / 63 388 578 / 65 252 780 octets** après les trois
éditions ; le partagé déplace **1 818 déclarations inchangées**, contre zéro
pour les modifications de corps feuille/FFI. Il s'agit de tailles statiques ;
le Go multipackage n'est pas encore émis ni chronométré.

Preuve : `gopurs-go-packages-g01-20261013/`, notamment
`canonical-interval-models-1/`, `canonical-verification-1/`,
`identity-contracts-1/`, `initialization-contracts-4/` et `final-plan/`.
L'audit final confirme les **430 fichiers de production** de G00.
**G01 est validé, suite G à 2/7 ; G02 porte les références jusqu'au rendu.**

## Suite G — G02 : références qualifiables validées (13 octobre 2026)

Le [contrat G02](build-go-references.md) porte les symboles des getters,
workers, constructeurs et helpers jusqu'au rendu. Les workers locaux sont
distingués des symboles publiés ; les types nommés portent leur propriétaire.
`GoSource.bindSources` résout les unités complètes avec `go/types`, corps et
exports réels compris. `printBoundGoCode` qualifie selon le consommateur en
préservant les membres et portées locales.

La campagne initiale fournit **12 comparaisons b8x** : base/feuille/partagé/FFI,
JS/Go séquentiel/Go parallèle, **2 704 TAST / 3 010 Go** par passage, exacts
face à v3. Les objets `workers-v2` des quatre états sont également exacts.
La campagne finale ajoute **six passages** : base native, identique sans
réécriture, feuille incrémentale/forcée, puis base JS/Rust. Elle conserve la
même racine et l'inode **286740996**. Ces durées de validation ne remplacent
pas les paires de mesures G03/G06.

Les cinq corpus G00 passent la liaison typée. Base/feuille/FFI : **185 671
symboles**, partagé : **185 672**, bootstrap : **65 338**. Le contrôle vérifie
les cibles internes et tous les spans ; le rendu monolithique restitue les noms
et octets historiques. Les métadonnées exhaustives de diagnostic restent hors
du cache de production ; leur taille impose un transport compact en G03.

Les contrôles ciblés couvrent AST, signatures et codecs, records/tuples,
Rebox, arbres possédés, tags, workers et contrats des bridges. Les replays
stricts finaux exécutent **46 commandes**, deux fixtures et quatre hôtes ;
les trois snapshots examinés par empreinte ne changent pas. Le résolveur rend
des JSON identiques via WASM, appel Go direct et bridge C/Rust, avec génériques,
callbacks, assertions, champs embarqués, interfaces et exports Go réels.
Les tests des packages parseur et cache passent sous `-race` ; les compagnons
natifs passent aussi dans leur module de compatibilité Go 1.22.

La migration `workers-v1` → `workers-v2` est éprouvée sur **61 modules** :
source modifiée, reprise depuis le seul manifeste, identique, forcé byte-exact,
compilation/exécution, **dix commandes**, inode **287016135** conservé. Le
premier candidat reste disponible comme oracle négatif de la reprise.

Tentatives conservées : classification initiale de `init`, méthodes
d'interface et champs embarqués ; assertion Rebox historique ; vecteur de
codec avec labels de champs non encodés ; première hypothèse de migration
avec journal encore valide ; usage de `utf16.RuneLen` refusé par le vet Go 1.22.
La dernière correction ne touche que le comptage UTF-16 du résolveur.
Les bootstraps et transports finaux sont identifiés séparément des campagnes
b8x dans `gopurs-go-packages-g02-20261013/final-reference/summary.json`.

**G02 est validé, suite G à 3/7 ; P reste à 0/7.** L'émission multipackage,
les changements d'ABI/réflexion, la publication et les gains mesurés relèvent
des quatre lots G restants.

## Suite G — G03 : émission multipackage validée (14 octobre 2026)

L'[émetteur typé](build-go-emission.md) est raccordé au CLI avec
`GOPURS_GO_PACKAGES=1`. Les sources, packages, entrées et exports figés sont
transportés en WASM, Go direct et C/Rust ; le CLI utilise le transport compact
par fichiers. Les tests couvrent scopes, cycles, records privés/exportés,
génériques, ordre des effets et cellule lazy partagée par 64 appels concurrents.

Les quatre hôtes réussissent les fixtures réelles et les migrations CLI. Après
revue des annotations et correction de l'identité Rebox, deux snapshots reçoivent
uniquement **27 / 134 lignes d'annotations**, puis **67 commandes strictes**
rejouent trois fixtures sur les quatre hôtes.

b8x publie **2 704 TAST / 39 613 Go / 525 packages** ; son oracle compile et rend
`camel=""`, `accent="e"`, `pong`. Le bootstrap courant publie **512 TAST /
8 389 Go / 150 packages** ; son binaire compile et passe lui-même les contrôles
CLI. La campagne finale termine **huit comparaisons / 28 commandes** avec
inventaires Go, quatre objets par module, TAST et dates byte-exacts entre hôtes.
Les hits identiques ne réécrivent aucun Go et Go ne compile ni ne lie.

Les **15 plans G01** figés sont reproduits. Les plans réellement émis retrouvent
exactement les modèles actualisés : **35 596 arêtes b8x / 5 875 bootstrap**.
Les 76 renommages de helpers, leurs 51 déplacements et les différences de corpus
bootstrap sont documentés. Les **348 tags runtime** b8x sont uniques et exacts ;
les **185 667 déclarations** sont syntaxiquement équivalentes après les seuls
renommages prévus. Le parseur et le cache passent aussi sous Go 1.22.12.

Les reproducteurs du trap WASM, du tampon JS, des annotations UTF-16, des noms
Rebox, de la propriété des anciennes sorties b8x et des échecs de harnais sont
conservés. `gopurs-go-packages-g03-20261013/final-reference/summary.json` identifie
les candidats et l'audit final. **G03 est clos, suite G à 4/7 ; P à 0/7.** Ces
campagnes valident la correction ; les seuils temporels restent à vérifier en G06.

## Suite G — G04 : consommateurs validés (14 octobre 2026)

Les [consommateurs multipackages](build-go-consumers.md) passent sur JS, Go
séquentiel/parallèle et Rust : **36 passages / 207 commandes strictes**, mêmes
TAST, sources exactes et exécution égale au monolithe. La revue des neuf nouveaux
snapshots couvre **592 déclarations / 57 fragments**, les wrappers FFI, chemins,
imports d'ordre et entrées. Les références sont créées après compilation et
exécution, puis restent inchangées pendant le rejeu strict.

Les **89 commandes** de `driver-packages` vérifient aussi les FFI absentes,
erreurs de syntaxe, reprises, signatures et dépendances natives. Les suites
réelles `foreign`/`js-promise`, onze probes réfléchies b8x et le bootstrap courant
passent. `gopurs-go-packages-g04-20261014/final-reference/summary.json` conserve
les empreintes et compteurs ; les tentatives préalables restent disponibles.
**G04 est clos, suite G à 5/7 ; P reste à 0/7.**

## Suite G — G05 : cache et runner validés (14 octobre 2026)

Le [cache multipackage](build-go-incremental.md) passe **56 comparaisons /
350 commandes** de transitions sur les quatre hôtes. Les sources avant partition
sont des objets gzip privés, vérifiés en forme compressée et décompressée ; les
quatre objets métier conservent leurs codecs. Les corruptions, récupérations,
migrations, morts de processus et verrous passent, avec détection de races et
Go 1.22.12. Le candidat Rust à archive périmée reste conservé et le remplacement
est validé. Les **89 commandes CLI** du candidat final embarquent aussi le
contrôle de récupération après un flux gzip invalide rehashé.

b8x ne réécrit qu'un Go après l'édition FFI, avec **2 704 backends repris /
2 703 émissions restaurées**, exactitude face au forcé puis au retour à la base,
oracle et identique sans compilation/lien. Ses sources compressées uniques
occupent **12 641 367 octets**. Le runner b8x passe dix contrôles, dont migration,
dépendance transitive et deux entrées. Preuves :
`gopurs-go-packages-g05-20261014/final-reference/summary.json`.
**G05 est clos, suite G à 6/7 ; P reste à 0/7.**

## Suite G — G06 : candidat regroupé, correction validée (14 octobre 2026)

Le candidat 2 de `gopurs-go-packages-g06-20261014` regroupe les sources par
package avec des imports résolus par `go/types`. Les directives exigeant des
frontières de fichiers gardent le rendu fragmenté, avec le même résolveur.

- Suites du résolveur et du cache passées, dont cache avec détection de races
  sous Go 1.22.12, migrations fragmenté/regroupé et ordre des initialiseurs FFI.
- **56 comparaisons / 350 commandes** de transitions, récupération du cache et
  **89 commandes CLI** sur JS, Go séquentiel/parallèle et Rust.
- **612 déclarations Main/FFI** examinées dans neuf snapshots, avant mise à jour
  après compilation/exécution ; **36 passages / 207 commandes** de rejeu strict,
  avec empreintes inchangées pendant le rejeu.
- Transports purs WASM/Go/C-Rust et writer éprouvés ; le WASM du dépôt est
  byte-exact face au candidat figé.

Preuve : `gopurs-go-packages-g06-20261014/functional-candidate-2/summary.json`.
Le [bilan G06](build-go-measurements.md) conserve les deux pilotes rejetés pour
l'acceptation. Le pilote 2 passe ses quatre oracles mesurés et les deux seeds,
avec **−84,71 % Go / −66,46 % chaîne** sur une paire feuille, mais dépasse les
budgets GOCACHE de base, RSS backend, gopurs froid et identique. Les diagnostics
suivants conservent les octets Go et situent le pic RSS pendant la publication.
Le candidat 3 privé, à ancres 1/8192 avec acquisition d'exports allégés, passe
ensuite le résolveur, **89 commandes CLI et 56 comparaisons de transitions**
sur les quatre hôtes (`functional-private-3/summary.json`). Son pilote de
performance est en cours ; les snapshots du dépôt restent ceux du candidat 2.
**À ce stade du 14 octobre, G06 reste ouvert : G à 6/7, P à 0/7.**

## Suite G — G06 : livraison validée, pause avant P (16 octobre 2026)

Le **candidat 7** est retenu et installé à densité **1/8192**, avec ses binaires
JS/Go/Rust et son WASM identiques aux artefacts mesurés. Le découpage conserve
l'algorithme canonique G01, les contraintes d'initialisation et les bridges ;
b8x produit **6 packages / 15 arêtes / 27 Go**, soit **122 994 214 octets**.

### Acceptation mesurée

Les **36 chaînes / 36 oracles** représentent trois paires alternées pour chacun
des six scénarios. Les **39 critères de temps et de stockage passent**, avec
ressources du runner inchangées. Le [bilan ventilé](build-go-measurements.md) et
les [36 lignes brutes](build-go-cycle-g06.csv) publient médianes, plages et paires.

| Scénario | Go | Chaîne jusqu'à `exec` | Backend |
| --- | ---: | ---: | ---: |
| Feuille | −78,73 % | −62,88 % | **+26,88 %** |
| Partagé | −65,95 % | −48,36 % | **+34,37 %** |
| FFI | −66,98 % | −54,50 % | **+42,62 %** |
| Froid applicatif | −49,14 % | −31,89 % | +2,15 % |
| Froid complet | −46,20 % | −28,16 % | +2,23 % |

L'identique ne réécrit, ne compile et ne lie rien : **4,160 → 4,050 s de chaîne**.
Le GOCACHE identique est à **+9,997897 %**, à environ **23 976 octets** du plafond.
Le RSS backend augmente de **1,09–11,21 %** selon le scénario, celui du compilateur
Go baisse de **43,41–47,01 %** sur les builds qui compilent. Le RSS reste un
indicateur secondaire selon la décision utilisateur du 14 octobre.

**La référence immédiate `publicsync` était déjà ralentie.** Le backend reste
à **~118 s à froid**, loin des **~30 s initialement évoquées**, et à **48,5 s sur
la feuille**. La livraison valide les gains Go/chaîne de G ; elle conserve le
surcoût de préparation et de cache accumulé avant G. Aucun gain futur de P
n'entre dans ce verdict. Le temps d'exécution métier/PostgreSQL est hors du
chronomètre jusqu'à `exec`.

### Correction et installation

- **Huit comparaisons / 28 commandes** des grands corpus : b8x **2 704 TAST /
  27 Go**, bootstrap **512 TAST / 23 Go**, byte-exacts entre JS, Go séquentiel,
  Go parallèle et Rust ; forcé/hit exacts, dates conservées, reconstruction
  sans compilation/lien et oracle b8x exécuté.
- **56 comparaisons / 350 commandes de transitions** sur le candidat 7 : mêmes
  TAST, parties métier/préparées, sources privées et Go exécutés dans les quatre
  hôtes ; les **89 commandes CLI** passent avant et après installation.
- **612 déclarations Main/FFI** examinées dans neuf snapshots. Leurs changements
  de chemins, imports et qualifications sont revus avant mise à jour. Les
  **207 commandes strictes sur l'installation** passent, avec références et
  snapshots monolithiques inchangés pendant le rejeu.
- **82 tests ciblés** des transports, writer, lanceurs et runners de tests ;
  intégrité, publication interrompue, récupération et transitions du pilote dans
  les quatre hôtes. Les suites Go avec détection de races et compatibilité
  **Go 1.22.12** sont conservées dans les validations du candidat.
- Bootstrap régénéré depuis les sources courantes : **512 TAST / 23 Go exacts**,
  binaire reproduit à l'octet près, puis contrôle CLI séquentiel/parallèle.
- Runner b8x : **10 contrôles sur macOS et 10 sur Linux**, dont migration,
  dépendance transitive, deux entrées, concurrence, annulation et réutilisation.
- Audit des artefacts installés et de **295 fichiers source/outils** ; deux
  fixtures ajoutent des identités d'ancre fixes pour conserver leurs frontières
  avec la nouvelle densité. `git diff --check` passe dans gopurs et b8x.

Les échecs de harnais sont conservés : fillers insuffisants à la nouvelle densité,
vrai compilateur transmis aux mocks, mauvais Spago/profondeur de workspace et
PATH Go absent dans le conteneur. Les reprises gardent les contrôles stricts et
les preuves réussies ; aucun changement du compilateur mesuré n'est introduit.

Preuves : `gopurs-go-packages-g06-20261014/final-reference/summary.json`,
`delivery-7/replays-2/`, `delivery-7/replays-3/`, `current-bootstrap-7/`,
`final-private-7/`, `snapshot-review-7.json`, `full-2/combined/`.
L'[usage et la migration](build-cache.md#sélectionner-le-découpage-go-g06) sont
documentés dans gopurs et `b8x/test/go-run.md`, avec sélection
`GOPURS_GO_PACKAGES=1` et retour monolithique explicite.
**G est clos à 7/7. P reste à 0/7 : pause effective avant P00.**

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
`<tmp>/gopurs-v2-campaigns-koy5tymo/`.
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
`<tmp>/gopurs-v2-tast-strings-5p8w6c8y/`,
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
`<tmp>/gopurs-v2-utf16-aaiyq81c/`,
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
`<tmp>/gopurs-v2-int32-d3hedy_d/`,
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
`<tmp>/gopurs-v2-local-install-7cj5t5s4/`,
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
`<tmp>/gopurs-v2-number-show-mw16tagz/`,
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
`<tmp>/gopurs-v2-deriving-bgLIyO/`,
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
`<tmp>/gopurs-v2-libraries-FnSmNh/`,
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
`<tmp>/gopurs-v2-final-DFj3Sa/`,
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
`<tmp>/gopurs-final-consolidation-xy_5fmmn/`.
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
`<tmp>/gopurs-runtime-ffi-cleanup-6052azh5/`.
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
