# Carte des modules

Cette carte est la **source de vérité du lot 3** : une ligne par module, sa
phase cible, sa responsabilité et les modules dont il dépend réellement
(imports `Gopurs.*` du `.purs`, doublons retirés). Le nom cible et le lot
proviennent de `cleanup.md` ; le glossaire est dans [naming.md](naming.md).

## Couverture

| Référence | Modules |
| --- | ---: |
| Baseline du 29 octobre 2026 | 105 |
| Supprimé au lot 1 (`DependencyGraph`, aucun import) | 1 |
| **Présents au 29 octobre 2026** | **104** |
| dont à la racine de `src/Gopurs` | 62 |
| dont imbriqués sous `src/Gopurs` | 41 |
| dont `src/Main.purs` | 1 |

La baseline comptait 63 modules à la racine, dont `DependencyGraph` ; sa
suppression au lot 1 ramène le total à 104. La carte couvre donc les 105
modules de la baseline : 104 lignes actuelles et la ligne de suppression en
annexe. `node tools/check-docs.mjs` vérifie que chaque `.purs` de `src/` a
exactement une ligne et qu'aucune ligne ne pointe vers un module disparu.

## Phases et dépendances autorisées

Ordre cible de lecture : `Main` → `Driver` → `Prepare` → `Rewrite` →
`Emit` ; `Cache` sert les trois dernières ; `Host` et `Runtime` sont des
feuilles. Une phase ne dépend que d'elle-même et des phases situées en aval
dans cet ordre, sauf les exceptions assumées listées plus bas.

| Phase | Dépendances autorisées |
| --- | --- |
| Driver | Toutes les phases : le pilote est le seul orchestrateur. |
| Prepare | Prepare, Cache (codecs et preuves de lookup), Host, `Emit.Go` (représentations Go des métadonnées) et `Emit.State` (contrat des métadonnées). Ni Driver, ni Rewrite, ni Emit hors `Emit.Go` et `Emit.State`. |
| Rewrite | Rewrite, Prepare.Metadata, Emit.Go et Emit.State (IR et métadonnées partagés), Cache.LookupTable, Host. Ni Driver ni le dispatcher `Emit`. |
| Emit | Emit, Rewrite (composition des passes et preuves d'ABI partagées), Prepare.Metadata, Cache.LookupTable, Host. Ni Driver. |
| Cache | Cache, Emit (types et entrée d'émission), Prepare, Host. Ni Driver. |
| Host | Host pur pour Snapshot, Diagnostics et Metrics : aucune dépendance Gopurs. Resident est une entrée de service : il peut dépendre de `Driver`, `Driver.Config` et `Driver.Prepare`. |
| Runtime | Aucune dépendance Gopurs : le texte est généré par `tools/embed-runtime.mjs`. |
| Entrée | Driver et Host uniquement. |

Les dépendances réelles sont celles du code au 29 octobre 2026. Quand une
ligne viole la règle ci-dessus, elle figure aussi dans la table des écarts.

## Driver — orchestration

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.Driver`<br>`src/Gopurs/Driver.purs` | `Gopurs.Driver` | Driver | Ordonne les phases de compilation et expose l'entrée `compile`. | `Cache.HeaderPreload`, `Cache.Modules`, `Diagnostics`, `Driver.Build`, `Driver.Cache`, `Driver.Config`, `Driver.Manifest`, `Driver.Output`, `Driver.Prepare`, `Driver.Prepare.Persistent`, `Metrics` |
| `Gopurs.Driver.Build`<br>`src/Gopurs/Driver/Build.purs` | `Gopurs.Driver.Build` | Driver | Choisit le builder PBO, publie les signatures et alimente la file d'émission. | `Cache.HeaderPreload`, `Cache.Modules`, `Cache.Wire`, `Driver.Cache`, `Driver.Config`, `Driver.Prepare`, `Emission`, `Metrics` |
| `Gopurs.Driver.Config`<br>`src/Gopurs/Driver/Config.purs` | `Gopurs.Driver.Config` | Driver | Lit les options CLI et l'environnement du pilote, avec leurs bornes. | `Emission` |
| `Gopurs.Driver.Manifest`<br>`src/Gopurs/Driver/Manifest.purs` | `Gopurs.Driver.Manifest` | Driver | Construit, lit et compare le manifeste des entrées et sorties. | `Cache.Wire`, `ExprContext` |
| `Gopurs.Driver.Output`<br>`src/Gopurs/Driver/Output.purs` | `Gopurs.Driver.Output` | Driver | Traduit un module préparé et écrit modules, runtime et points d'entrée. | `CodeGen`, `CodegenState`, `Driver.Cache`, `ExprContext`, `FfiBridge`, `FfiSupport`, `FfiTypes`, `GoSource`, `GoSymbol`, `Runtime` |
| `Gopurs.Emission`<br>`src/Gopurs/Emission.purs` | `Gopurs.Driver.Scheduler` | Driver | Ordonne les lots d'émission, leurs dépendances et la contre-pression des fibres. | `Metrics` |

## Prepare — chargement, métadonnées et spécialisation

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.AdtMetadata`<br>`src/Gopurs/AdtMetadata.purs` | `Gopurs.Prepare.Metadata.Adt` | Prepare | Décrit les représentations ADT (pointeurs, énumérations) dérivées du TAST. | `GoAst` |
| `Gopurs.ClassMetadata`<br>`src/Gopurs/ClassMetadata.purs` | `Gopurs.Prepare.Metadata.Class` | Prepare | Indexe les classes, leurs superclasses et leurs méthodes par label. | — |
| `Gopurs.ConstructorLayout`<br>`src/Gopurs/ConstructorLayout.purs` | `Gopurs.Prepare.Metadata.ConstructorLayout` | Prepare | Calcule l'identité, les champs et les arguments génériques des constructeurs. | `CodegenState`, `ConstructorMetadata`, `GoAst`, `GoSymbol`, `GoTypes` |
| `Gopurs.ConstructorMetadata`<br>`src/Gopurs/ConstructorMetadata.purs` | `Gopurs.Prepare.Metadata.Constructor` | Prepare | Décrit les constructeurs, leurs champs et leurs identités. | `GoAst` |
| `Gopurs.Driver.Prepare`<br>`src/Gopurs/Driver/Prepare.purs` | `Gopurs.Prepare` | Prepare | Charge le TAST, assemble les métadonnées et prépare les modules spécialisés. | `AdtMetadata`, `Cache.LookupTable`, `Cache.Wire`, `ClassMetadata`, `CodegenState`, `ConstructorMetadata`, `Driver.Cache`, `Driver.Prepare.Projection`, `Driver.Prepare.Proof`, `GlobalTypes`, `Metrics`, `Monomorphization`, `Preparation`, `ReboxMetadata` |
| `Gopurs.Driver.Prepare.Persistent`<br>`src/Gopurs/Driver/Prepare/Persistent.purs` | `Gopurs.Prepare.Persistent` | Prepare | Précharge, projette et publie les artefacts persistants de préparation. | `Cache.LookupTable`, `Cache.Wire`, `ClassMetadata`, `Driver.Cache`, `Driver.Manifest`, `Driver.Prepare`, `Driver.Prepare.Projection`, `Driver.Prepare.ProjectionCapsule`, `Driver.Prepare.Proof`, `Metrics`, `Monomorphization`, `Preparation` |
| `Gopurs.Driver.Prepare.Projection`<br>`src/Gopurs/Driver/Prepare/Projection.purs` | `Gopurs.Prepare.Persistent.Projection` | Prepare | Découpe et fusionne les projections de modules préparés. | `Cache.LookupTable`, `CodegenState` |
| `Gopurs.Driver.Prepare.ProjectionCapsule`<br>`src/Gopurs/Driver/Prepare/ProjectionCapsule.purs` | `Gopurs.Prepare.Persistent.ProjectionCapsule` | Prepare | Emballe et rouvre une projection en fragments vérifiés. | `Cache.LookupTable`, `CodegenState`, `Driver.Prepare.Projection` |
| `Gopurs.Driver.Prepare.Proof`<br>`src/Gopurs/Driver/Prepare/Proof.purs` | `Gopurs.Prepare.Persistent.Proof` | Prepare | Calcule les preuves d'identité et de réutilisation d'une préparation. | `Cache.Wire` |
| `Gopurs.GlobalTypes`<br>`src/Gopurs/GlobalTypes.purs` | `Gopurs.Prepare.Metadata.GlobalTypes` | Prepare | Résout les types globaux des bindings et des imports FFI. | — |
| `Gopurs.Monomorphization`<br>`src/Gopurs/Monomorphization.purs` | `Gopurs.Prepare.Specialize` | Prepare | Collecte, propage et applique les spécialisations monomorphes. | `Monomorphization.ForeignForwarders`, `NativeRecordArgs` |
| `Gopurs.Monomorphization.ForeignForwarders`<br>`src/Gopurs/Monomorphization/ForeignForwarders.purs` | `Gopurs.Prepare.Specialize.ForeignForwarders` | Prepare | Reconnaît les wrappers FFI exclus des copies spécialisées. | — |
| `Gopurs.Preparation`<br>`src/Gopurs/Preparation.purs` | `Gopurs.Prepare.Jobs` | Prepare | Exécute un tour de travaux de préparation en `Aff`, ordonné et borné. | — |
| `Gopurs.ReboxMetadata`<br>`src/Gopurs/ReboxMetadata.purs` | `Gopurs.Prepare.Metadata.Rebox` | Prepare | Indexe les champs à convertir entre représentations natives et boxées. | `ClassMetadata`, `ConstructorMetadata`, `GoAst` |

## Rewrite — passes sur l'IR

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.BorrowedObjects`<br>`src/Gopurs/BorrowedObjects.purs` | `Gopurs.Rewrite.BorrowedObjects` | Rewrite | Prouve et réécrit les emprunts d'objets JSON en lecture seule. | `Cache.LookupTable`, `CodegenState` |
| `Gopurs.ClosedDictionaries`<br>`src/Gopurs/ClosedDictionaries.purs` | `Gopurs.Rewrite.ClosedDictionaries` | Rewrite | Réécrit et met en cache les dictionnaires fermés. | `ClosedDictionaries.Admission`, `CodegenState` |
| `Gopurs.ClosedDictionaries.Admission`<br>`src/Gopurs/ClosedDictionaries/Admission.purs` | `Gopurs.Rewrite.ClosedDictionaries.Admission` | Rewrite | Sépare la preuve de réutilisation de la nouvelle mise en cache. | `Cache.LookupTable`, `ClosedDictionaries.Scope`, `CodegenState` |
| `Gopurs.ClosedDictionaries.Scope`<br>`src/Gopurs/ClosedDictionaries/Scope.purs` | `Gopurs.Rewrite.ClosedDictionaries.Scope` | Rewrite | Vérifie la fermeture lexicale et l'absence d'effets ou de récursion. | — |
| `Gopurs.DecoderSchemas`<br>`src/Gopurs/DecoderSchemas.purs` | `Gopurs.Rewrite.DecoderSchemas` | Rewrite | Spécialise les schémas de décodeurs JSON et leurs workers. | `Cache.LookupTable`, `CodegenState`, `DecoderSchemas.Admission`, `DecoderSchemas.Source`, `DecoderSchemas.Workers`, `GoAst` |
| `Gopurs.DecoderSchemas.Admission`<br>`src/Gopurs/DecoderSchemas/Admission.purs` | `Gopurs.Rewrite.DecoderSchemas.Admission` | Rewrite | Admet un schéma de décodeur et ses programmes. | `DecoderSchemas.Programs`, `DecoderSchemas.Source`, `DecoderSchemas.Types` |
| `Gopurs.DecoderSchemas.Programs`<br>`src/Gopurs/DecoderSchemas/Programs.purs` | `Gopurs.Rewrite.DecoderSchemas.Programs` | Rewrite | Dérive les programmes de dictionnaires d'un schéma. | `DecoderSchemas.Source`, `DecoderSchemas.Types` |
| `Gopurs.DecoderSchemas.Source`<br>`src/Gopurs/DecoderSchemas/Source.purs` | `Gopurs.Rewrite.DecoderSchemas.Source` | Rewrite | Résout les dictionnaires sources et les symboles des schémas. | — |
| `Gopurs.DecoderSchemas.Types`<br>`src/Gopurs/DecoderSchemas/Types.purs` | `Gopurs.Rewrite.DecoderSchemas.Types` | Rewrite | Définit les schémas, programmes et poids d'admission. | — |
| `Gopurs.DecoderSchemas.Workers`<br>`src/Gopurs/DecoderSchemas/Workers.purs` | `Gopurs.Rewrite.DecoderSchemas.Workers` | Rewrite | Construit les workers des schémas de décodeurs. | `DecoderSchemas.Types`, `GoAst`, `GoSymbol`, `Printer` |
| `Gopurs.FunctionFusion`<br>`src/Gopurs/FunctionFusion.purs` | `Gopurs.Rewrite.FunctionFusion` | Rewrite | Admet et réécrit les producteurs de fonctions comptés. | `WorkerNames` |
| `Gopurs.ImmediateApplications`<br>`src/Gopurs/ImmediateApplications.purs` | `Gopurs.Rewrite.ImmediateApplications` | Rewrite | Admet et réduit les applications immédiates sûres. | `ImmediateApplications.Scope` |
| `Gopurs.ImmediateApplications.Scope`<br>`src/Gopurs/ImmediateApplications/Scope.purs` | `Gopurs.Rewrite.ImmediateApplications.Scope` | Rewrite | Calcule occurrences, substitutions et renommages des niveaux locaux. | — |
| `Gopurs.NativeRecordArgs`<br>`src/Gopurs/NativeRecordArgs.purs` | `Gopurs.Rewrite.NativeRecordArgs` | Rewrite | Vérifie TCO et choisit l'ABI des arguments record. | `GoAst`, `NativeRecordArgs.Projection`, `NativeRecordArgs.Source` |
| `Gopurs.NativeRecordArgs.Projection`<br>`src/Gopurs/NativeRecordArgs/Projection.purs` | `Gopurs.Rewrite.NativeRecordArgs.Projection` | Rewrite | Projette les champs et admet les résultats natifs. | `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.NativeRecordArgs.Source`<br>`src/Gopurs/NativeRecordArgs/Source.purs` | `Gopurs.Rewrite.NativeRecordArgs.Source` | Rewrite | Prouve la source des lecteurs partagés avant monomorphisation. | `NativeRecordArgs.Projection` |
| `Gopurs.Ownership`<br>`src/Gopurs/Ownership.purs` | `Gopurs.Rewrite.Ownership` | Rewrite | Calcule le point fixe des workers consommants et sélectionne les appels frais. | `CodegenState`, `GoAst`, `GoSymbol`, `Ownership.Analysis`, `Ownership.Candidates`, `Ownership.Types`, `Ownership.Workers` |
| `Gopurs.Ownership.Analysis`<br>`src/Gopurs/Ownership/Analysis.purs` | `Gopurs.Rewrite.Ownership.Analysis` | Rewrite | Analyse chemins, scalaires, termes d'arbres et fraîcheur. | `CodegenState`, `GoAst`, `Ownership.Types`, `PrimitiveExprs` |
| `Gopurs.Ownership.Candidates`<br>`src/Gopurs/Ownership/Candidates.purs` | `Gopurs.Rewrite.Ownership.Candidates` | Rewrite | Admet les layouts et signatures, et réserve les paires de noms. | `CodegenState`, `GoAst`, `GoTypes`, `Ownership.Analysis`, `Ownership.Types` |
| `Gopurs.Ownership.Types`<br>`src/Gopurs/Ownership/Types.purs` | `Gopurs.Rewrite.Ownership.Types` | Rewrite | Définit les types de l'analyse d'ownership. | `CodegenState`, `GoAst` |
| `Gopurs.Ownership.Workers`<br>`src/Gopurs/Ownership/Workers.purs` | `Gopurs.Rewrite.Ownership.Workers` | Rewrite | Émet captures, cellules réutilisables et déclarations consommantes. | `CodegenState`, `GoAst`, `GoSymbol`, `Ownership.Analysis`, `Ownership.Types` |
| `Gopurs.ThunkFusion`<br>`src/Gopurs/ThunkFusion.purs` | `Gopurs.Rewrite.ThunkFusion` | Rewrite | Admet et réécrit les producteurs de thunks en workers comptés. | `WorkerNames` |
| `Gopurs.WorkerNames`<br>`src/Gopurs/WorkerNames.purs` | `Gopurs.Rewrite.Names` | Rewrite | Réserve les noms des workers produits par les fusions. | `GoAst` |

## Emit — traduction et rendu Go

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.AdtExprs`<br>`src/Gopurs/AdtExprs.purs` | `Gopurs.Emit.Expr.Adt` | Emit | Assemble les représentations ADT et la réutilisation de constructeurs. | `CodegenState`, `ConstructorLayout`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `Printer` |
| `Gopurs.ArrayIntrinsics`<br>`src/Gopurs/ArrayIntrinsics.purs` | `Gopurs.Emit.Expr.ArrayIntrinsics` | Emit | Reconnaît et émet map, filter et foldl. | `ArrayIntrinsics.Index`, `ArrayIntrinsics.Source`, `CallAnalysis`, `CallArguments`, `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.ArrayIntrinsics.Index`<br>`src/Gopurs/ArrayIntrinsics/Index.purs` | `Gopurs.Emit.Expr.ArrayIntrinsics.Index` | Emit | Admet l'indexation et convertit l'élément sélectionné. | `ArrayIntrinsics.Source`, `CallAnalysis`, `CallArguments`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.ArrayIntrinsics.Source`<br>`src/Gopurs/ArrayIntrinsics/Source.purs` | `Gopurs.Emit.Expr.ArrayIntrinsics.Source` | Emit | Lie un tableau déjà évalué sans copier ses éléments. | `GoAst` |
| `Gopurs.ArrayTraverse`<br>`src/Gopurs/ArrayTraverse.purs` | `Gopurs.Emit.Expr.ArrayTraverse` | Emit | Émet les boucles des traversées Either sur tableaux. | `AdtExprs`, `CallAnalysis`, `CallArguments`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.BindingExprs`<br>`src/Gopurs/BindingExprs.purs` | `Gopurs.Emit.Expr.Binding` | Emit | Émet les bindings locaux, la récursion et l'initialisation. | `CallAnalysis`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoTypes`, `LocalWorkers`, `Printer` |
| `Gopurs.CallAnalysis`<br>`src/Gopurs/CallAnalysis.purs` | `Gopurs.Emit.Expr.Call.Analysis` | Emit | Analyse les cibles d'appel, les spines et les closures. | `Cache.LookupTable`, `ExprAnalysis`, `ExprContext`, `GoAst` |
| `Gopurs.CallArguments`<br>`src/Gopurs/CallArguments.purs` | `Gopurs.Emit.Expr.Call.Arguments` | Emit | Traduit, capture et adapte les arguments d'appel. | `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.CallExprs`<br>`src/Gopurs/CallExprs.purs` | `Gopurs.Emit.Expr.Call` | Emit | Sélectionne appels, surapplications et sauts TCO. | `ArrayIntrinsics`, `ArrayTraverse`, `CallAnalysis`, `CallArguments`, `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `ObjectTraverse` |
| `Gopurs.CodeGen`<br>`src/Gopurs/CodeGen.purs` | `Gopurs.Emit` | Emit | Dispatcher récursif et assemblage du fichier Go d'un module. | `AdtExprs`, `ArrayIntrinsics`, `BindingExprs`, `BorrowedObjects`, `Cache.LookupTable`, `CallExprs`, `ClosedDictionaries`, `CodegenState`, `ConstructorExprs`, `ControlExprs`, `DecoderSchemas`, `EffectExprs`, `ExprAnalysis`, `ExprContext`, `FunctionExprs`, `FunctionFusion`, `GoAst`, `GoConversions`, `GoImports`, `GoSymbol`, `GoTypes`, `ImmediateApplications`, `LiteralExprs`, `ModuleBindings`, `ModuleDeclarations`, `Ownership`, `PrimitiveExprs`, `Printer`, `RecordExprs`, `ThunkFusion`, `TypedExprs` |
| `Gopurs.CodegenState`<br>`src/Gopurs/CodegenState.purs` | `Gopurs.Emit.State` | Emit | Décrit les métadonnées immuables et l'état mutable d'une émission. | `AdtMetadata`, `Cache.LookupTable`, `ClassMetadata`, `ConstructorMetadata`, `GoAst`, `GoSymbol`, `ReboxMetadata` |
| `Gopurs.ConstructorExprs`<br>`src/Gopurs/ConstructorExprs.purs` | `Gopurs.Emit.Expr.Constructor` | Emit | Traduit les définitions et constructions saturées de constructeurs. | `AdtExprs`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.ControlExprs`<br>`src/Gopurs/ControlExprs.purs` | `Gopurs.Emit.Expr.Control` | Emit | Traduit branches, alternatives et structures de contrôle. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `Printer` |
| `Gopurs.EffectExprs`<br>`src/Gopurs/EffectExprs.purs` | `Gopurs.Emit.Expr.Effect` | Emit | Traduit les expressions d'effet et leurs enveloppes. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `Printer` |
| `Gopurs.ExprAnalysis`<br>`src/Gopurs/ExprAnalysis.purs` | `Gopurs.Emit.Expr.Analysis` | Emit | Fournit les analyses d'expressions en lecture seule. | `ExprContext`, `GoAst`, `GoTypes` |
| `Gopurs.ExprContext`<br>`src/Gopurs/ExprContext.purs` | `Gopurs.Emit.Expr.Context` | Emit | Transporte le contexte, le résultat et les callbacks de traduction. | `CodegenState`, `GoAst` |
| `Gopurs.FfiBridge`<br>`src/Gopurs/FfiBridge.purs` | `Gopurs.Emit.FFI.Bridge` | Emit | Façade d'assemblage du bridge FFI typé. | `CodegenState`, `FfiBridge.Render`, `FfiBridge.Signatures`, `FfiTypes`, `GoAst` |
| `Gopurs.FfiBridge.Render`<br>`src/Gopurs/FfiBridge/Render.purs` | `Gopurs.Emit.FFI.Bridge.Render` | Emit | Rend wrappers, workers et déclarations FFI manquantes. | `FfiBridge.Signatures`, `FfiBridge.TypeSupport`, `FfiBridge.Values`, `FfiTypes`, `GoTypes` |
| `Gopurs.FfiBridge.Signatures`<br>`src/Gopurs/FfiBridge/Signatures.purs` | `Gopurs.Emit.FFI.Bridge.Signatures` | Emit | Apparie déclarations étrangères et signatures d'appel natif. | `CodegenState`, `FfiTypes`, `GoAst`, `GoSymbol` |
| `Gopurs.FfiBridge.TypeSupport`<br>`src/Gopurs/FfiBridge/TypeSupport.purs` | `Gopurs.Emit.FFI.Bridge.TypeSupport` | Emit | Instancie les types du bridge et leur générique. | `FfiTypes` |
| `Gopurs.FfiBridge.Values`<br>`src/Gopurs/FfiBridge/Values.purs` | `Gopurs.Emit.FFI.Bridge.Values` | Emit | Adapte arguments, callbacks et résultats FFI. | `FfiBridge.TypeSupport`, `FfiTypes`, `GoTypes` |
| `Gopurs.FfiSupport`<br>`src/Gopurs/FfiSupport.purs` | `Gopurs.Emit.FFI.Support` | Emit | Analyse la FFI Go et expose les opérations de support. | `FfiTypes` |
| `Gopurs.FfiTypes`<br>`src/Gopurs/FfiTypes.purs` | `Gopurs.Emit.FFI.Types` | Emit | Définit les types du bridge FFI. | — |
| `Gopurs.FunctionExprs`<br>`src/Gopurs/FunctionExprs.purs` | `Gopurs.Emit.Expr.Function` | Emit | Traduit abstractions, applications et fonctions. | `CallAnalysis`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoSymbol` |
| `Gopurs.GoAst`<br>`src/Gopurs/GoAst.purs` | `Gopurs.Emit.Go.Ast` | Emit | Définit les nœuds d'expression, de type et de déclaration Go. | `GoCode`, `GoSymbol` |
| `Gopurs.GoCode`<br>`src/Gopurs/GoCode.purs` | `Gopurs.Emit.Go.Code` | Emit | Porte les fragments opaques et leurs dépendances d'imports. | `GoSymbol` |
| `Gopurs.GoConversions`<br>`src/Gopurs/GoConversions.purs` | `Gopurs.Emit.Go.Conversions` | Emit | Choisit les conversions, le boxing et la lecture de Value. | `CodegenState`, `GoAst`, `GoConversions.NativeAdts`, `GoConversions.Rebox`, `Printer` |
| `Gopurs.GoConversions.NativeAdts`<br>`src/Gopurs/GoConversions/NativeAdts.purs` | `Gopurs.Emit.Go.Conversions.NativeAdts` | Emit | Décrit les layouts natifs de Maybe, Either et Tuple. | `GoAst`, `Printer` |
| `Gopurs.GoConversions.Rebox`<br>`src/Gopurs/GoConversions/Rebox.purs` | `Gopurs.Emit.Go.Conversions.Rebox` | Emit | Émet les helpers de conversion et les demandes Rebox. | `CodegenState`, `Diagnostics`, `GoAst`, `GoSymbol`, `GoTypes`, `Printer`, `ReboxMetadata` |
| `Gopurs.GoFunctions`<br>`src/Gopurs/GoFunctions.purs` | `Gopurs.Emit.Workers.Parameters` | Emit | Construit paramètres, variables d'itération et enveloppes curryfiées. | `GoAst` |
| `Gopurs.GoImports`<br>`src/Gopurs/GoImports.purs` | `Gopurs.Emit.Go.Imports` | Emit | Collecte et rend les imports d'un module Go. | `GoAst`, `GoCode`, `GoSymbol` |
| `Gopurs.GoSource`<br>`src/Gopurs/GoSource.purs` | `Gopurs.Emit.Go.Source` | Emit | Lie les sources Go aux symboles et assemble les fichiers de package. | `FfiSupport`, `GoCode`, `GoSymbol` |
| `Gopurs.GoSymbol`<br>`src/Gopurs/GoSymbol.purs` | `Gopurs.Emit.Go.Symbol` | Emit | Définit et compare les symboles Go. | — |
| `Gopurs.GoTypes`<br>`src/Gopurs/GoTypes.purs` | `Gopurs.Emit.Go.Types` | Emit | Traduit les annotations PureScript en types Go. | `AdtMetadata`, `GoAst`, `GoSymbol` |
| `Gopurs.Int32Loops`<br>`src/Gopurs/Int32Loops.purs` | `Gopurs.Emit.Expr.Int32Loops` | Emit | Prouve et versionne les boucles scalaires à état signé 32 bits. | `GoAst`, `GoFunctions` |
| `Gopurs.LiteralExprs`<br>`src/Gopurs/LiteralExprs.purs` | `Gopurs.Emit.Expr.Literal` | Emit | Traduit les littéraux de tableaux et de records, avec leurs conversions. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes`, `Printer`, `RecordExprs` |
| `Gopurs.LocalWorkers`<br>`src/Gopurs/LocalWorkers.purs` | `Gopurs.Emit.Workers.Local` | Emit | Émet les workers locaux, leurs paramètres et signatures. | `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoSymbol`, `GoTypes`, `Int32Loops` |
| `Gopurs.ModuleBindings`<br>`src/Gopurs/ModuleBindings.purs` | `Gopurs.Emit.Module.Bindings` | Emit | Émet les fonctions de module, leurs signatures et groupes TCO. | `CallAnalysis`, `CodegenState`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoSymbol`, `GoTypes`, `ModuleWorkers`, `NativeRecordArgs` |
| `Gopurs.ModuleDeclarations`<br>`src/Gopurs/ModuleDeclarations.purs` | `Gopurs.Emit.Module.Declarations` | Emit | Émet les structs ADT et l'enregistrement des getters de classes. | `CodegenState`, `GoAst`, `GoSymbol`, `GoTypes` |
| `Gopurs.ModuleWorkers`<br>`src/Gopurs/ModuleWorkers.purs` | `Gopurs.Emit.Workers.Module` | Emit | Émet les workers de module et leurs wrappers. | `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoFunctions`, `GoSymbol`, `Int32Loops` |
| `Gopurs.ObjectTraverse`<br>`src/Gopurs/ObjectTraverse.purs` | `Gopurs.Emit.Expr.ObjectTraverse` | Emit | Émet les boucles des traversées Either sur objets. | `AdtExprs`, `CallAnalysis`, `CallArguments`, `ExprContext`, `GoAst`, `GoConversions` |
| `Gopurs.PrimitiveExprs`<br>`src/Gopurs/PrimitiveExprs.purs` | `Gopurs.Emit.Expr.Primitive` | Emit | Traduit les primitives et leurs conversions. | `CodegenState`, `GoAst`, `GoConversions` |
| `Gopurs.Printer`<br>`src/Gopurs/Printer.purs` | `Gopurs.Emit.Go.Printer` | Emit | Rend les déclarations Go en texte. | `GoAst`, `GoCode`, `GoSymbol`, `Printer.Builder` |
| `Gopurs.Printer.Builder`<br>`src/Gopurs/Printer/Builder.purs` | `Gopurs.Emit.Go.Printer.Builder` | Emit | Tient le buffer mutable privé d'un rendu. | — |
| `Gopurs.RecordExprs`<br>`src/Gopurs/RecordExprs.purs` | `Gopurs.Emit.Expr.Record` | Emit | Traduit les records et choisit leur représentation. | `CodegenState`, `GoAst`, `GoConversions`, `GoTypes` |
| `Gopurs.TypedExprs`<br>`src/Gopurs/TypedExprs.purs` | `Gopurs.Emit.Expr.Typed` | Emit | Propage les annotations typées et construit les dictionnaires de classes. | `CallAnalysis`, `ExprAnalysis`, `ExprContext`, `GoAst`, `GoConversions`, `GoTypes` |

## Cache — protocole et réutilisation

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.BuildCache`<br>`src/Gopurs/BuildCache.purs` | `Gopurs.Cache.Protocol` | Cache | Définit le protocole de stockage versionné et ses échanges JSON. | — |
| `Gopurs.Cache.BindingCapsule`<br>`src/Gopurs/Cache/BindingCapsule.purs` | `Gopurs.Cache.BindingCapsule` | Cache | Encode et sélectionne les capsules de bindings. | `Cache.Wire` |
| `Gopurs.Cache.HeaderPreload`<br>`src/Gopurs/Cache/HeaderPreload.purs` | `Gopurs.Cache.HeaderPreload` | Cache | Précharge et décode les en-têtes de cache. | `Cache.Wire`, `Driver.Cache`, `Metrics`, `Preparation` |
| `Gopurs.Cache.LookupTable`<br>`src/Gopurs/Cache/LookupTable.purs` | `Gopurs.Cache.LookupTable` | Cache | Représente les lectures de cache prouvées et complètes. | `Cache.Wire` |
| `Gopurs.Cache.Modules`<br>`src/Gopurs/Cache/Modules.purs` | `Gopurs.Cache.Modules` | Cache | Gère les clés et la réutilisation par module (préparé, PBO, émission). | `Cache.BindingCapsule`, `Cache.HeaderPreload`, `Cache.Wire`, `CodeGen`, `CodegenState`, `Driver.Cache`, `Driver.Output`, `Driver.Prepare`, `ExprContext`, `GoSymbol`, `Metrics`, `Preparation` |
| `Gopurs.Cache.Pool`<br>`src/Gopurs/Cache/Pool.purs` | `Gopurs.Cache.Pool` | Cache | Intern les fragments JSON canoniques d'un artefact. | — |
| `Gopurs.Cache.TypeMemo`<br>`src/Gopurs/Cache/TypeMemo.purs` | `Gopurs.Cache.TypeMemo` | Cache | Mémoïse les types immuables par index canonique. | — |
| `Gopurs.Cache.Wire`<br>`src/Gopurs/Cache/Wire.purs` | `Gopurs.Cache.Wire` | Cache | Fournit les codecs portables et les tables partagées. | `Cache.Pool`, `Cache.TypeMemo`, `GoAst`, `GoSymbol` |
| `Gopurs.Driver.Cache`<br>`src/Gopurs/Driver/Cache.purs` | `Gopurs.Cache.Session` | Cache | Possède la session de stockage, le staging et la recette FFI. | `BuildCache`, `Driver.Config`, `Metrics` |

## Host — services dépendants de l'hôte

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.CompilerSnapshot`<br>`src/Gopurs/CompilerSnapshot.purs` | `Gopurs.Host.Snapshot` | Host | Copie et réexécute le compilateur depuis un emplacement privé. | — |
| `Gopurs.Diagnostics`<br>`src/Gopurs/Diagnostics.purs` | `Gopurs.Host.Diagnostics` | Host | Écrit les diagnostics et signale leur non-rejouabilité. | — |
| `Gopurs.Metrics`<br>`src/Gopurs/Metrics.purs` | `Gopurs.Host.Metrics` | Host | Mesure les phases et accumule les durées d'émission. | — |
| `Gopurs.Resident`<br>`src/Gopurs/Resident.purs` | `Gopurs.Host.Resident` | Host | Boucle résidente déclenchée par fichier, réutilisant la session de compilation. | `Driver`, `Driver.Config`, `Driver.Prepare` |

## Runtime — runtime embarqué

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Gopurs.Runtime`<br>`src/Gopurs/Runtime.purs` | `Gopurs.Runtime.Source` | Runtime | Fournit le texte Go du runtime embarqué. | — |

## Entrée

| Module actuel | Cible (lot 3) | Phase | Responsabilité | Dépend de |
| --- | --- | --- | --- | --- |
| `Main`<br>`src/Main.purs` | `Main` | Entrée | Point d'entrée CLI : snapshot, options, exécution et code de sortie. | `CompilerSnapshot`, `Driver`, `Driver.Config`, `Resident` |

## Écarts de frontière actuels

Ces imports traversent une frontière interdite par la table ci-dessus. Ils
sont conservés pendant le lot 2, puis retirés par le lot qui possède le module
appelé. Aucun n'est introduit par un nouveau code.

| Module | Importe | Pourquoi c'est un écart | Retrait |
| --- | --- | --- | --- |
| `Cache.Modules` | `Driver.Output` | Le cache appelle directement l'émetteur du pilote ; la cible fait dépendre `Driver` de `Cache`. | 3b |
| `Driver.Cache` | `Driver.Config` | La session lit encore les options par le module du pilote ; la cible reçoit une valeur de configuration. | 3b |
| `Driver.Prepare.Persistent` | `Driver.Manifest` | La persistance de préparation emprunte le manifeste du pilote au lieu de son contrat propre. | 3c |
| `Prepare` (`Monomorphization`) | `Rewrite.NativeRecordArgs` | La barrière de spécialisation consulte une preuve de partage des lecteurs ; la preuve doit rejoindre `Prepare.Metadata` ou le cache. | 3c–3d |
| `Rewrite.Ownership.Analysis` | `Emit.Expr.Primitive` | L'analyse d'ownership reconnaît une primitive avec l'émetteur ; la reconnaissance doit rejoindre l'IR partagé. | 3d |
| `Prepare` (`Driver.Prepare`, `ConstructorLayout`, `Projection`, `ProjectionCapsule`) | `Emit.State` (`CodegenState`) | `CodegenState` porte le contrat des métadonnées produites par `Prepare` ; la dépendance est assumée jusqu'au découpage du lot 4. | 4 |
| `Rewrite` (passes) | `Emit.State` (`CodegenState`) | Même contrat de métadonnées partagé ; les passes ne doivent pas dépendre du dispatcher `Emit`. | 4 |

## Annexes

### Module supprimé

| Module | Ancien chemin | Raison |
| --- | --- | --- |
| `Gopurs.DependencyGraph` | `src/Gopurs/DependencyGraph.purs` | Aucun import ; supprimé au lot 1. |

### Entretien de la carte

- Une PR de déplacement met à jour la ligne du module, son chemin et ses
  dépendances dans le même commit que le `git mv`.
- Une extraction ou un changement d'import met à jour la colonne « Dépend de »
  du module appelant et, si nécessaire, la table des écarts.
- Le nom cible ne devient le nom actuel qu'après le merge de la PR du lot.
  Tant que le déplacement n'est pas mergé, la colonne « Module actuel » reste
  la référence pour les liens et les outils.
- `node tools/check-docs.mjs` échoue si un module de `src/` manque, si une
  ligne référence un module absent, ou si deux lignes portent le même nom.
