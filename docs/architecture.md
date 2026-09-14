# Architecture du backend

État vérifié le 14 septembre 2026. Les noms ci-dessous renvoient aux propriétaires
actuels du code ; les étapes suivent le chemin actif de [Main](../src/Main.purs).

## Entrée typée et ordre des passes

1. Le fork PureScript émet le TAST enrichi dans `output/<Module>/corefn.json`.
   `coreFnModulesFromOutput` du PBO local lit ces fichiers, les décode et trie les
   modules par dépendances. Le nom de fichier ne signifie pas que le backend
   utilise le CoreFn standard : `dataDecls`, `classDecls`, les annotations de
   types et les instanciations `TypeApp` restent accessibles après décodage.
2. `Main.loadAndPrepareModules` collecte les constructeurs éliminés, charge les
   directives, puis construit les types des constructeurs, les types globaux
   et les champs des classes. Il enrichit les modules avec les déclarations de
   données synthétiques des classes.
3. `Monomorphization.monomorphizeModules` collecte les instanciations, propage
   les besoins transitifs, filtre les candidats et spécialise les modules.
   Les métadonnées de représentation des ADT sont calculées depuis les modules
   enrichis avant spécialisation ; les types globaux viennent des modules
   d'origine. Ces tables accompagnent ensuite les modules monomorphisés.
4. `Builder.buildModules` du PBO reçoit ces modules et les convertit en
   `BackendModule` optimisés, avec les directives et `coreForeignSemantics`.
   **La monomorphisation de gopurs précède donc cette optimisation PBO.**
5. `Main.emitModule` transmet chaque module à `CodeGen.translateWithFunctions`.
   Celui-ci applique `ThunkFusion.optimizeThunkProducers`, crée un état local
   de traduction, prépare l'analyse TCO et les signatures des fonctions via
   `ModuleBindings`, puis traduit les déclarations et leurs expressions.
6. Les émetteurs construisent les expressions et les déclarations `GoDecl`,
   dont les helpers de conversion. `GoImports.collectImports` collecte leurs
   dépendances avant que `Printer.printGoFile` les rende en texte Go. Les
   signatures produites sont rendues à `Main`, qui les transporte vers les
   modules suivants pour les appels directs entre modules.
7. Après le code du module, `emitModule` traite sa FFI Go : localisation,
   préparation et décodage des déclarations, puis génération du bridge typé.
   Les entrées exécutables sont écrites après le parcours des modules.

Les `ForAll`, contraintes, types structurels, ordre des champs et queues de
rangées font partie des données typées utilisées par le backend. Une décision
de représentation doit partir de ces informations, pas d'une reconstruction
à partir de chaînes Go déjà imprimées.

## Où modifier une responsabilité

Tous les modules gopurs de ce tableau se trouvent dans [src/Gopurs](../src/Gopurs).

| Responsabilité | Modules |
| --- | --- |
| Types globaux, constructeurs et classes | `GlobalTypes`, `ConstructorMetadata`, `ClassMetadata` |
| Représentations des ADT et spécialisation | `AdtMetadata`, `Monomorphization` |
| Contrats distincts des métadonnées immuables et de l'état mutable | `CodegenState` |
| Dispatcher récursif, assemblage du fichier | `CodeGen` |
| Contexte, résultat et callbacks de traduction | `ExprContext` |
| Fonctions de module, signatures, groupes TCO | `ModuleBindings` |
| Structs ADT et enregistrement des getters de classes | `ModuleDeclarations` |
| Bindings locaux, récursion locale et initialisation | `BindingExprs` |
| Sélection des appels, surapplications et sauts TCO | `CallExprs`, `CallAnalysis` |
| Traduction ordonnée et adaptation des arguments | `CallArguments` |
| Reconnaissance et émission de map, filter et foldl | `ArrayIntrinsics` |
| Abstractions, branches et effets | `FunctionExprs`, `ControlExprs`, `EffectExprs` |
| ADT, records et primitives | `AdtExprs`, `RecordExprs`, `PrimitiveExprs` |
| Identité, champs et arguments génériques des constructeurs | `ConstructorLayout` |
| Analyses d'expressions | `ExprAnalysis` |
| Types Go, boxing et conversions | `GoTypes`, `GoConversions` |
| Préparation des enveloppes curryfiées | `GoFunctions` |
| Dépendances des fragments opaques et imports du module | `GoCode`, `GoImports` |
| Représentation et rendu Go | `GoAst`, `Printer` |
| Frontière FFI et adaptation des signatures | `FfiSupport`, `FfiBridge` |

`ExprContext` permet aux familles d'expressions de rappeler le dispatcher sans
importer `CodeGen`. Il transporte directement les tables immuables de
`CodegenMetadata`, utilisées pour le typage, les signatures et la préparation
des constructeurs et des records. Ces décisions ne lisent aucune référence
mutable.

La référence `CodegenState` contient uniquement les déclarations structurées
produites (`declarations`), le compteur des bindings récursifs (`globalId`) et les
couples de conversion à émettre (`reboxPairs`). Le compteur local `nextId` et
les statements restent transportés dans les résultats ; les déclarations
principales sont renvoyées directement par `ModuleBindings.declarations`.
La génération Rebox consulte les métadonnées directement et relit les couples
accumulés jusqu'à avoir émis les conversions transitives. Le printer ne consulte
ni ne modifie cet état. Le [contrat AST/printer](go-ast-printer.md) décrit les
nœuds structurés et les familles qui restent assemblées en chaînes.

`CodeGen` assemble trois groupes ordonnés de `GoDecl` : valeurs avec cache,
déclarations natives et helpers Rebox, puis getters FFI. Tous empruntent le
même parcours d'imports et de rendu. Les corps opaques encore nécessaires,
notamment les affectations Rebox et l'enregistrement des getters de classes,
portent leur texte et leurs dépendances dans `GoCode`. Le constructeur `rawGo`
reconnaît les imports historiques à la création du fragment ; aucun module
Go imprimé n'est reparcouru pour calculer ses imports.

`ConstructorLayout` partage la préparation entre définitions, constructions
saturées et accès aux champs. Il garde explicites les variantes d'instanciation
des champs et du pointeur. `TypeStructPointer` porte un record nommé : identité
du tag runtime, type PureScript, constructeur Go, nom Go instancié et arguments
de type. `GoAst.structPointer` construit le nom instancié ; les conversions et
l'instanciation générique réutilisent le nom du constructeur sans découper le
texte imprimé.

`CallArguments` traduit chaque argument dans l'ordre, en conservant ses
statements et le compteur de noms. Les intrinsics curryfiés boxent chaque
argument immédiatement ; les autres chemins conservent sa représentation
native jusqu'à l'adaptation de l'appel. `ArrayIntrinsics` reçoit ces arguments
déjà traduits et émet les boucles. `CallExprs` conserve les priorités existantes :
TCO avant les intrinsics pour `App`, intrinsics avant TCO pour `UncurriedApp`.

## Runtime et FFI

[runtime/runtime.go](../runtime/runtime.go) est la source canonique du runtime.
`tools/embed-runtime.mjs` génère le module FFI `Runtime.js` au build ; Spago puis
esbuild embarquent sa constante dans le bundle. `Main` écrit ce texte dans
l'application générée, sans relire la source Go au lancement du backend.

La FFI utilise une autre chaîne : `tools/ffi-gen` analyse le Go avec son AST
natif et expose le contrat JSON par WebAssembly. `tools/ffi-runner.mjs` exécute
le WASM ; `FfiSupport` prépare la source et décode la réponse, puis `FfiBridge`
rapproche les déclarations Go des types TAST. Une erreur de syntaxe, de runner
ou de décodage fait échouer le build ; elle ne devient pas une liste vide de
fonctions. L'absence de fichier FFI suit encore le chemin de bridge de secours
de `Main`, distinct d'un échec d'analyse.

## Sorties et cache

| Chemin dans l'application | Contenu |
| --- | --- |
| `output/purescript/Module_Name.go` | Code du module PureScript |
| `output/purescript/Module_Name_ffi.go` | FFI et bridge, pour les modules concernés |
| `output/gopurs_runtime/runtime.go` | Copie exacte du runtime embarqué |
| `output/go.mod` | Module `gopurs/output`, version Go déclarée 1.22 |
| `output/main/main.go` | Entrée partagée sélectionnée par `--main` |
| `output/Module.Name/main/main.go` | Entrée propre à chaque module ciblé |

`onSkipModule` renvoie toujours `Nothing` : gopurs régénère les modules et leur
FFI à chaque invocation. Aucun `.gopurs-cache.json` n'est lu ou écrit par ce
chemin. Le PBO emploie encore `.purmeta` pour les implémentations nécessaires à
l'optimisation ; cela ne remplace pas la génération Go. Les caches de paquets
et de compilation Spago restent une couche distincte.

Le backend ne purge pas les anciens fichiers de `output` lorsqu'un module ou
une FFI disparaît. Pour comparer deux générateurs, conserver les mêmes entrées
TAST et inventorier les sorties ; pour une fixture, le runner fournit un
workspace neuf. Voir [les vérifications](testing.md).

## Carte des dépôts et des consommateurs

Le [suivi du lot 1](../todo.md#lot-1--carte-et-référence-du-14-septembre-2026)
attribue les **51 dépôts indépendants** et toutes leurs familles de fichiers.
Le répertoire parent `gopurs/` est leur conteneur ; son `output/` et son fichier
IDE `.psc-ide-port` ne constituent pas une bibliothèque supplémentaire.
Les noms publics sont ceux des déclarations `module` et de leurs exports,
indépendamment du préfixe `gopurs-` des répertoires.

| Famille et propriétaire | Entrées → traitement → sorties | Consommateurs et revue |
| --- | --- | --- |
| Configuration de chaque dépôt | `package.json`, Bower/Dhall, Spago et lockfiles → choix des outils, dépendances et sources | npm, Spago, Pulp, CI et éditeurs ; lot 2 |
| Build de gopurs | `runtime/runtime.go` → `embed-runtime.mjs` → `Runtime.js` ; `src/**/*.purs` et FFI JS → Spago/esbuild → `bin/gopurs.js` | npm `prepare`, `bin/gopurs`, runners et applications ; lot 4 |
| Entrée et pipeline gopurs | TAST de l'application → métadonnées → monomorphisation → PBO → émetteurs → AST/printer | Fichiers Go et signatures utilisées par les modules suivants ; lots 5–7 |
| Bibliothèque `gopurs-*` | `src/**/*.purs` → modules publics, réexports, instances et signatures étrangères ; compagnons `.go`/`.js` | Imports des autres paquets, des applications et des tests ; lots 9–14 |
| Bridge gopurs et FFI de bibliothèque | Chemin source du TAST + `.go` → recherche PBO → parser → déclarations JSON → rapprochement avec les types → `<Module>_ffi.go` | Programme Go généré ; lot 7 et lot de la bibliothèque |
| Runtime gopurs | Source Go canonique → chaîne d'embarquement ci-dessus → `output/gopurs_runtime/runtime.go` | Go généré, bridges et FFI des bibliothèques ; lot 8 |
| Parser gopurs | `tools/ffi-gen/{parser,types,main_js_wasm}.go` → build Go `js/wasm` → `ffi_gen.wasm`, avec `wasm_exec.js` du même Go | `ffi-runner.mjs`, puis `FfiSupport` ; lots 4 et 8 |
| Tests de gopurs | Fixtures et compagnons, snapshots, AST construits par les tests Node, cas Go du parser | Runners Node, modules PureScript compilés et Go ; lots 3, 4 et 8 |
| Tests, exemples et benchmarks des bibliothèques | `test/`, `example(s)/`, `bench/`, `benchmark/`, `integration-tests/`, compagnons et données | `bin/test`, commandes npm/Pulp/Spago et CI propres au paquet ; lots 4 et 9–14 |
| Documentation de chaque dépôt | README, guides, licences, images et docs | Utilisateurs et mainteneurs ; lot 15, avec contexte actualisé dans chaque lot |

`bin/pkg` déclare **22 paquets core**, consommés par `bin/setup` et le runner
de fixtures. La liste complète de `bin/modtest` est découverte au lancement :
**49 frères ont un `bin/test` exécutable**, sur 50 bibliothèques. QuickCheck
n'en a pas. Une entrée dans `extraPackages` est un choix de résolution, pas la
preuve qu'un paquet appartient aux dépendances effectivement utilisées.
La compilation du backend emploie son propre graphe Spago : PBO local, plus
les overrides `st`, `unsafe-coerce` et `assert`. Le graphe de l'application
détermine les bibliothèques et FFI Go qu'elle utilise.

## Usages qui échappent à une recherche d'imports

Avant de supprimer un fichier ou un export, suivre aussi ces entrées :

- `Main` découvre les exports `main` quand `--main` est absent. Le parseur PBO
  lit les TAST depuis `output`, sans import de fichier explicite dans gopurs.
- `findFfiFile` du PBO essaie d'abord le compagnon du `modulePath` TAST, puis
  les répertoires FFI et Spago. Une FFI Go peut donc être consommée sans aucun
  import textuel vers son chemin. La FFI JS est sélectionnée par le compilateur
  PureScript pour les outils et parcours JavaScript conservés.
- `FfiSupport.js` résout le runner relativement à la source, au module compilé
  ou au bundle. Le runner instancie le WASM et appelle le global `parseFFI`
  enregistré par `main_js_wasm.go` ; cette fonction n'a pas d'import JS ordinaire.
- Les tests Node importent `output/Gopurs.*/index.js` et construisent les AST
  avec leurs exports. Les supprimer du bundle ou renommer un constructeur
  demande de vérifier ces consommateurs au lot 3.
- Le runner découvre les fixtures par répertoire et lit `@dependencies` et
  `@snapshot-ffi`. Il copie aussi les fichiers et répertoires compagnons.
  Les snapshots et données d'entrée ne sont pas des sorties jetables.
- Le runtime consulte notamment `StructGetters` par tag et emploie `reflect`
  pour des conversions de champs. Les noms et enregistrements générés font
  partie de ce contrat, même sans appel statique visible dans une bibliothèque.
- `gopurs-spec/test/Integration.purs` découvre les sous-répertoires de
  `integration-tests/cases`, leurs sorties attendues et `env-template`.
  `gopurs-console/scripts/test` est appelé par npm et compare la sortie à
  `test/expected_output.txt`. Des commandes npm de benchmarks et de chaînes
  importent également les modules compilés directement.

## Provenance et sources à conserver

L'inventaire du lot 1 sépare fichiers suivis, fichiers locaux non suivis,
artefacts ignorés et dépendances installées. Les attributs Linguist des dépôts
masquent largement le PureScript et le JavaScript : `linguist-generated` ou
`linguist-vendored` n'établit pas à lui seul qu'un fichier est reconstructible.

46 bibliothèques ont un remote `upstream`, mais aucune référence locale
`upstream/*` n'était disponible lors de l'inventaire. Le backend,
`argonaut-core`, `js-bigints` et `js-date` n'ont pas ce remote ; QuickCheck a
directement l'origin PureScript upstream et est détaché sur **v8.0.1**.
Ses fichiers suivis correspondent à ce tag. Ses deux ajouts locaux non suivis,
`spago.yaml` et `src/Test/QuickCheck/Gen.go`, participent à l'adaptation Go et
restent à examiner aux lots 2 et 14.

Une comparaison octet par octet de `src/` avec **33 versions du registre déjà
présentes dans `.spago/p` de gopurs** donne 224 fichiers identiques, 36 fichiers
différents, 79 ajouts locaux et aucun fichier manquant. Ces versions sont des
références locales datées, pas les dernières versions upstream. Parmi les
ajouts, 77 sont des `.go` ; les deux autres sont
`Control/Monad/Free.js` et `Data/Map/Internal.js`. Les différences touchent
aussi le PureScript et le JavaScript de `aff`, `argonaut-core`, `arrays`,
`foldable-traversable`, `foreign`, `foreign-object`, `free`, `lazy`,
`node-event-emitter`, `nullable`, `ordered-collections`, `record`, `refs` et
`strings`. Les versions et chemins exacts sont dans l'inventaire temporaire
mentionné au lot 1 ; les 17 autres bibliothèques n'ont pas été comparées à une
distribution du registre dans ce contrôle. QuickCheck dispose du tag ci-dessus.

Les licences sont maintenues par dépôt. `gopurs-yoga-json` conserve aussi
`LICENCE/simple-json.LICENSE`. `gopurs-js-bigints` n'a actuellement ni README
ni fichier de licence suivi : le lot 12 doit établir ce contexte, sans déduire
sa provenance du seul nom du paquet.

Les artefacts distribués suivis de gopurs sont `tools/ffi_gen.wasm` et
`tools/wasm_exec.js` ; leur reconstruction est `npm run build:ffi`, avec
Go **1.27.0** imposé par `tools/ffi-gen/go.mod`. `bin/gopurs.js` et
`src/Gopurs/Runtime.js` sont générés et ignorés. Les `output/`, `.spago/`,
`.purmeta/`, caches npm et `node_modules/`, y compris ceux des exemples et
environnements d'intégration, se contrôlent par leurs entrées et leur commande
de reconstruction. Ne pas y reporter des modifications de sources.
