# Architecture du backend

Les étapes suivent le chemin actif de [Driver](../src/Gopurs/Driver.purs), appelé
par [Main](../src/Main.purs). Le pilote et ses frontières sont décrits ci-dessous.

## Entrée typée et ordre des passes

1. Le fork PureScript émet le TAST enrichi dans `output/<Module>/corefn.json`.
   `coreFnModulesFromOutput` du PBO local lit ces fichiers, les décode et trie les
   modules par dépendances. Le nom de fichier ne signifie pas que le backend
   utilise le CoreFn standard : `dataDecls`, `classDecls`, les annotations de
   types et les instanciations `TypeApp` restent accessibles après décodage.
2. `Driver.Prepare.prepareModules` collecte les constructeurs éliminés, charge les
   directives, puis construit les types des constructeurs, les types globaux
   et les champs des classes. Il enrichit les modules avec les déclarations de
   données synthétiques des classes.
3. `Monomorphization.monomorphizeModulesWith` collecte les instanciations, propage
   les besoins transitifs, filtre les candidats et spécialise les modules.
   Les métadonnées pointeurs viennent des modules enrichis avant spécialisation ;
   les enums et types globaux viennent des modules d'origine. Ces tables
   accompagnent ensuite les modules monomorphisés.
4. `Driver.Build` choisit `Builder.buildModules` ou `buildModulesParallel` du PBO.
   Ces builders convertissent les modules en `BackendModule` optimisés, avec les
   directives et `coreForeignSemantics`.
   **La monomorphisation de gopurs précède donc cette optimisation PBO.**
5. `Driver.Output.emitModule` lit les signatures FFI natives, puis transmet
   chaque module à `CodeGen.translateWithFunctions`.
   Celui-ci applique `ThunkFusion.optimizeThunkProducers`, crée un état local
   de traduction, prépare l'analyse TCO et les signatures des fonctions via
   `ModuleBindings`, puis traduit les déclarations et leurs expressions.
6. Les émetteurs construisent les expressions et les déclarations `GoDecl`,
   dont les helpers de conversion. `GoImports.collectImports` collecte leurs
   dépendances avant que `Printer.printGoFile` les rende en texte Go. Les
   signatures produites sont rendues à `Driver.Build`. Chaque lot lit une vue
   immuable ; ses résultats sont fusionnés dans l'ordre des modules avant que
   le lot suivant puisse les consulter pour les appels directs entre modules.
7. `Driver.Output` écrit le code du module et son bridge FFI typé. Il fournit
   aussi l'écriture du runtime, de `go.mod` et des entrées exécutables. `Driver`
   écrit le runtime avant le build, puis les entrées après sa réussite.

Les `ForAll`, contraintes, types structurels, ordre des champs et queues de
rangées font partie des données typées utilisées par le backend. Une décision
de représentation doit partir de ces informations, pas d'une reconstruction
à partir de chaînes Go déjà imprimées.

## Pilote de compilation

`Main` lance l'effet avec `runAff_` et traite son résultat après le nettoyage du
pilote. Sur erreur, il écrit le diagnostic original sur stderr puis demande une
sortie de statut 1. La petite FFI `Main.js` laisse Node vider ses écritures ;
`Main.go` termine le processus après le diagnostic synchrone. Le point d'entrée
Go ne consulte pas le code stocké par `Node.Process.setExitCode`.

À la frontière FFI, `Driver.Output` convertit explicitement les exceptions de
préparation `Effect` en erreurs `Aff` : le `liftEffect` natif ne les intercepte
pas. Elles suivent ainsi le même nettoyage et le même diagnostic CLI que les
erreurs des opérations de fichiers asynchrones.

`Driver.compile` rend visible l'ordre des phases. Les options sont lues une
fois par `Driver.Config`, avec les valeurs par défaut et les bornes existantes.
`Driver.Prepare` rend un `PreparedModules` contenant les
modules spécialisés, les directives, les modules exécutables et un champ
`metadata :: CodegenMetadata` distinct de ces données d'orchestration.

`Driver.Build` possède les références mutables de la compilation : signatures
déjà publiées et compteurs de progression. Son coordinateur est le seul
producteur de la file d'émission, même avec le builder PBO parallèle.
`Emission.withEmitter` encadre **tout** le build : producteur, workers PBO et
émetteurs. Il termine le dernier lot sur succès et supervise les fibres avant
la fermeture de l'émetteur. Un worker peut être suspendu en publiant son
résultat dans l'AVar ; il doit aussi être arrêté si le coordinateur échoue.
Voir le [contrat du pipeline](parallel-emission.md) et ses tests JS/natifs.

## Où modifier une responsabilité

Tous les modules gopurs de ce tableau se trouvent dans [src/Gopurs](../src/Gopurs).

| Responsabilité | Modules |
| --- | --- |
| Ordre des phases et mesures | `Driver` |
| Arguments et variables d'environnement du pilote | `Driver.Config` |
| Chargement du TAST et assemblage des métadonnées | `Driver.Prepare` |
| Choix du builder, ordonnanceur Aff et publication des signatures | `Driver.Build` |
| Lots, dépendances, contre-pression et durée de vie des fibres | `Emission` |
| Traduction d'un module, assemblage FFI et fichiers de sortie | `Driver.Output` |
| Types globaux, constructeurs et classes | `GlobalTypes`, `ConstructorMetadata`, `ClassMetadata` |
| Représentations des ADT et spécialisation | `AdtMetadata`, `Monomorphization` |
| Contrats distincts des métadonnées immuables et de l'état mutable | `CodegenState` |
| Dispatcher récursif, assemblage du fichier | `CodeGen` |
| Contexte, résultat et callbacks de traduction | `ExprContext` |
| Annotations typées, propagation et dictionnaires de classes | `TypedExprs` |
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
| Analyse du Go FFI et façade du bridge | `FfiSupport`, `FfiBridge` |
| Appariement des déclarations et signatures d'appel natif | `FfiBridge.Signatures` |
| Types du bridge et instanciation générique | `FfiBridge.TypeSupport` |
| Adaptation des arguments, callbacks et résultats FFI | `FfiBridge.Values` |
| Rendu des wrappers, workers et déclarations manquantes | `FfiBridge.Render` |

`ExprContext` permet aux familles d'expressions de rappeler le dispatcher sans
importer `CodeGen`. Il transporte directement les tables immuables de
`CodegenMetadata`, utilisées pour le typage, les signatures et la préparation
des constructeurs et des records. Ces décisions ne lisent aucune référence
mutable.

Le dispatcher reçoit lui-même un `ExprContext`, comme les émetteurs spécialisés.
`ExprContext.childContext` prépare les opérandes ordinaires : profondeur suivante,
hors position terminale et hors bloc d'effet, sans cible TCO ni boucle héritée.
Le type attendu est passé explicitement. Les corps de bindings, branches et
fonctions conservent leurs propres règles de propagation du contexte.

`TypedExprs` traite les annotations `Typed` : il propage le type vers le résultat
des `Let`/`LetRec`, construit les dictionnaires dans l'ordre des champs de la
classe et adapte les autres résultats. Il garde explicites les cas où le boxing
doit être évité : records déjà boxés destinés à un consommateur dynamique, records
natifs projetés, sommes natives, pointeurs et closures. Chaque champ de dictionnaire
est traduit puis converti avant le suivant, car la conversion peut enregistrer
des helpers Rebox dans l'état du module.

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

`CallArguments` traduit chaque argument dans l'ordre, en conservant ses
statements et le compteur de noms. Les intrinsics curryfiés boxent chaque
argument immédiatement ; les autres chemins conservent sa représentation
native jusqu'à l'adaptation de l'appel. `ArrayIntrinsics` reçoit ces arguments
déjà traduits et émet les boucles. `CallExprs` conserve les priorités existantes :
TCO avant les intrinsics pour `App`, intrinsics avant TCO pour `UncurriedApp`.

## Choix des représentations

Les tables assemblées par `Driver.Prepare` ont des contrats distincts, repris
par les alias de types de `CodegenMetadata` :

| Table | Entrée et identité |
| --- | --- |
| `ConstructorMetadata.ConstructorTypes` | ADT d'origine ; clé `<module_avec_underscores>.<constructeur_brut>`, champs dans l'ordre de déclaration |
| `elidedCtors` | ADT d'origine avec un seul constructeur et un seul champ ; noms `Constructor_…`, avec les points du module historiquement conservés |
| `ClassMetadata.ClassFields` | Classes d'origine ; clé PureScript qualifiée avec points, méthodes et superclasses triées ensemble par label |
| `AdtMetadata.PointerAdtPaths` | ADT enrichis des dictionnaires synthétiques ; clé du type PureScript qualifié, constructeur à payload et arité déclarée |
| `pointerAdtNodes`, `pointerAdtLeaves`, `enumCtors` | Identités runtime `Data_<module_avec_underscores>_<constructeur_sanitisé>` |
| `enumAdts` | Types PureScript qualifiés, depuis les ADT d'origine |

Un ADT pointeur possède exactement un constructeur à payload et au plus une
alternative sans champ, représentable par `nil`. Plusieurs alternatives sans
champ doivent conserver leurs tags distincts. Un enum possède au moins un
constructeur, tous sans champ ; un dictionnaire de classe vide reste exclu des
enums. Les dictionnaires synthétiques sont ajoutés après
la collecte des constructeurs éliminés : une classe à une méthode conserve son
wrapper. `ClassMetadata.sortedClassFields` fournit déjà le même ordre à la table
nommée et au constructeur synthétique ; `ModuleDeclarations` s'en sert pour les
champs `V0…` et les getters.

`GoTypes.appliedAdt` rassemble les arguments d'un ADT puis de ses `TypeApp`, de
l'intérieur vers l'extérieur. Il ne retire ni `ForAll` ni les contraintes.
La recherche d'un pointeur préfère le nom exact, puis son alias `$Dict`.
L'instanciation conserve les politiques suivantes :

| Contexte | Arguments Go |
| --- | --- |
| Type de valeur | Préfixe fourni, tronqué à l'arité ; positions manquantes complétées par `Value` |
| Champ générique | Arguments fournis si leur nombre est exact ; sinon tous les paramètres déclarés si leur nombre convient ; sinon tous `Value` |
| Constructeur annoté par un ADT simple | Arité explicite du layout prioritaire ; si des arguments manquent, tous deviennent `Value`, sinon le préfixe est conservé |
| Définition annotée par un `TypeApp` | Préfixe des arguments appliqués, limité aux variables du layout, sans compléter les positions manquantes |
| Construction saturée annotée par un `TypeApp` | Variables du layout effacées en `Value` |

Les arités nulles ne produisent aucun argument générique. Les types de valeur
testent l'élimination sur le nom porté par le chemin ADT avant de choisir enum
ou pointeur ; les champs génériques choisissent l'enum d'abord, puis testent
l'élimination sur le constructeur à payload résolu. Ces priorités sont explicites
dans `GoTypes`. Les tableaux génériques gardent la traduction de valeur de leurs
éléments. Un record natif exige une queue de rangée absente ; le premier champ
de chaque label gagne, puis les labels sont triés.

`ConstructorLayout` résout l'identité avant les champs : le type pointeur attendu
peut fournir le module de définition ; un accès aux champs garde son qualificateur
explicite. Les champs ADT priment sur ceux d'une classe. La préparation choisit
une fois les arguments pour le pointeur et ses champs. Pour une alternative
`nil`, une définition résout le worker dans le module courant, une construction
saturée dans le module résolu ; l'identité runtime reste celle du payload.

`GoAst.constructorNames` partage les noms du constructeur Go et du tag runtime,
en laissant le préfixe de module à l'appelant. `TypeStructPointer` conserve ces noms,
le type PureScript qualifié et les arguments. `GoAst.structPointer` est le point
de construction du nom Go instancié. `instantiateGenericGoType` le reconstruit
après substitution, sans découper de texte imprimé ; une variable absente de
l'environnement devient `Value`.

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
de `Driver.Output`, distinct d'un échec d'analyse.

`FfiBridge` expose trois opérations : enveloppes boxées, signatures publiées
et workers à résultat boxé. `Signatures` centralise l'appariement des noms
(`<Module>_<Nom>`, puis sa variante suffixée `_`) et l'admission d'un
`NativeCandidate`. Ce candidat porte la déclaration Go, son type TAST, le besoin
d'un worker et le nom à appeler. La publication d'une `FunctionInfo` vérifie
ensuite les types de tous les arguments et du résultat. Cette étape est plus
stricte que l'émission des workers : certains workers générés restent inutilisés
par les appels directs, notamment avec un retour record ou un argument natif
non représentable par cette interface.

`TypeSupport` instancie les paramètres génériques des wrappers jusque dans les
callbacks et les conteneurs. `Values` possède les adaptations dans les deux sens
et partage le traitement des callbacks entre wrappers et workers. `Render`
assemble leurs déclarations Go. Les valeurs étrangères, les fonctions sans
argument et les signatures non admissibles conservent leur enveloppe boxée.
Une déclaration Go absente produit l'enveloppe qui signale « FFI not implemented »
à l'utilisation ; une source Go mal formée reste une erreur de compilation.

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

Pendant un build, le PBO conserve les implémentations `.purmeta` récemment
utilisées dans un cache LRU. À la frontière de chaque module, il ramène le
poids des entrées retenues à 64 Mio de données sérialisées. Cette limite ne
mesure pas le tas JavaScript décodé ; les imports d'un module peuvent aussi
dépasser ce budget pendant son traitement. Le début du build suivant purge
le cache et les validations des modules, afin de ne jamais reprendre une
spécialisation périmée. `clearPurmetaCache` permet toujours une purge explicite.

Le backend ne purge pas les anciens fichiers de `output` lorsqu'un module ou
une FFI disparaît. Pour comparer deux générateurs, conserver les mêmes entrées
TAST et inventorier les sorties ; pour une fixture, le runner fournit un
workspace neuf. Voir [les vérifications](testing.md).

## Outils de build et de validation

Les scripts de [tools](../tools) ont des responsabilités distinctes :

| Responsabilité | Fichier |
| --- | --- |
| Ordre des étapes du bootstrap, logs et publication atomique du binaire | `build-native.mjs` |
| Configuration Spago native, sélection du fork `purs`, vérification du TAST | `native-workspace.mjs` |
| Sous-processus, groupes de processus, signaux et descripteurs des logs | `command-runner.mjs` |
| Présentation des commandes et erreurs des campagnes de tests | `test-process.mjs` |
| Sélection des fixtures, snapshots et workspaces | `test-runner.mjs` et ses helpers `test-*` |
| Parcours des bibliothèques sœurs | `modtest-runner.mjs` |

`CommandRunner` exécute une commande à la fois. Il rend son statut de sortie ;
le client interprète les échecs après avoir affiché les logs et appelé
`checkInterrupted`. SIGINT/SIGTERM sont transmis au groupe actif, descendants
compris. Chaque propriétaire appelle `dispose` dans un `finally` pour retirer
les handlers de signaux. Le build natif et `TestProcesses` utilisent ce même
contrat, avec leur propre présentation des étapes.

Le bootstrap conserve son workspace sur échec et installe le nouveau binaire
par renommage d'un fichier préparé sur le même système de fichiers. La
vérification du TAST et la compilation Go précèdent cette publication.

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

`bin/pkg` déclare **22 paquets core** pour le runner de fixtures, **trois
checkouts de support** pour leur développement et **25 autres bibliothèques**.
`bin/setup --core` installe les 25 premiers checkouts et `--all` couvre les 50,
avec l'adaptation QuickCheck locale fournie au préalable. `--list` décrit la
sélection sans modifier de fichiers. La liste de `bin/modtest` est, elle,
découverte au lancement :
**49 frères ont un `bin/test` exécutable**, sur 50 bibliothèques. QuickCheck
n'en a pas. Une entrée dans `extraPackages` est un choix de résolution, pas la
preuve qu'un paquet appartient aux dépendances effectivement utilisées.
La compilation du backend emploie son propre graphe Spago : PBO local, plus
les overrides `st` et `unsafe-coerce`. L'override `assert`, inutilisé par ce
graphe, a été retiré au lot 2. Le graphe de l'application
détermine les bibliothèques et FFI Go qu'elle utilise.

Les configurations restent propres à chaque dépôt. Le lot 2 a retiré les
overrides inutilisés après comparaison des graphes résolus, en conservant les
overrides transitifs nécessaires et les 42 liens Spago suivis. Le
[guide de développement local](../README.md#develop-one-library-locally)
décrit les profils d'installation, les outils et la politique des lockfiles.
Les dépendances npm de gopurs servent uniquement au build : le bundle distribué
ne dépend ni d'esbuild ni du backend ES au lancement.

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
