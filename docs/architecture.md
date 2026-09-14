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
6. Les émetteurs construisent le `GoAst`, les déclarations brutes encore
   nécessaires et les helpers de conversion. `Printer.printGoFile` les rend en
   texte Go. Les signatures produites sont rendues à `Main`, qui les transporte
   vers les modules suivants pour les appels directs entre modules.
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
| Métadonnées et état mutable par traduction | `CodegenState` |
| Dispatcher récursif, assemblage du fichier | `CodeGen` |
| Contexte, résultat et callbacks de traduction | `ExprContext` |
| Fonctions de module, signatures, groupes TCO | `ModuleBindings` |
| Bindings locaux, récursion locale et initialisation | `BindingExprs` |
| Applications, appels directs et intrinsics | `CallExprs`, `CallAnalysis` |
| Abstractions, branches et effets | `FunctionExprs`, `ControlExprs`, `EffectExprs` |
| ADT, records et primitives | `AdtExprs`, `RecordExprs`, `PrimitiveExprs` |
| Analyses d'expressions | `ExprAnalysis` |
| Types Go, boxing et conversions | `GoTypes`, `GoConversions` |
| Représentation et rendu Go | `GoAst`, `Printer` |
| Frontière FFI et adaptation des signatures | `FfiSupport`, `FfiBridge` |

`ExprContext` permet aux familles d'expressions de rappeler le dispatcher sans
importer `CodeGen`. Le compteur de noms et les statements produits sont
transportés dans les résultats ; `CodegenState` conserve notamment les
déclarations et les couples de conversion à émettre. Le printer ne consulte
ni ne modifie cet état. Le [contrat AST/printer](go-ast-printer.md) décrit les
nœuds structurés et les familles qui restent assemblées en chaînes.

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
