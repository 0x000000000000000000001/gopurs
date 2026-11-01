# Vocabulaire et nommage

Ce document est normatif. Il fixe le sens des mots employés dans le code, les
documents d'architecture et les revues. La carte des modules
[module-map.md](module-map.md) applique ce vocabulaire à chaque module ; le plan
[cleanup.md](../cleanup.md) décrit les lots qui font passer le dépôt des noms
historiques aux noms cibles.

Tant qu'un lot de déplacement n'est pas mergé, le nom historique reste le nom
du code. Les documents normatifs emploient le terme cible dans la prose et
citent le nom historique entre accents graves quand ils parlent du fichier
actuel. Les journaux historiques ne sont pas réécrits.

## Glossaire normatif

Règle directrice : **un concept, un mot ; un mot, un concept**. Les phases sont
des verbes, les données des noms, les preuves des adjectifs.

| Terme | Sens unique retenu | À éviter |
| --- | --- | --- |
| `compile` | exécution complète du pilote, de la lecture des options à la publication | `build`, `run` pour le pilote |
| `driver` | orchestration : options, ordre des phases, résultat final | `main`, `pipeline` |
| `prepare` | chargement du TAST, métadonnées, spécialisation et cache de préparation | `preprocess`, `setup` |
| `specialize` | monomorphisation d'un module ou d'un binding | `instantiate`, `mono` |
| `rewrite` | transformation PureScript/IR avant l'émission | `optimize` pour nos passes |
| `emit` | production en mémoire des déclarations Go | `generate`, `output` |
| `output` | écriture des fichiers de sortie dans le staging ou le public | `publish`, `emit` |
| `publish` | validation et commit du cache, y compris le retrait des sorties possédées | `output`, `write` |
| `build` | construction PBO d'un `BackendModule` | `compile` |
| `worker` | fonction Go générée, locale ou de module | unité de parallélisme |
| `job` | unité de travail parallèle ordonnée par le pilote ou PBO | `worker`, `task` |
| `protocol` | format et transactions du stockage persistant | `cache` seul |
| `session` | durée de vie d'une compilation et son staging | `cache`, `store` |
| `snapshot` | capture immuable à un instant donné | `state` |
| `metadata` | faits immuables dérivés du TAST | `state`, `info` |
| `state` | état mutable d'une passe | `metadata`, `context` |
| `host` | JS, Go ou Rust qui exécute le compilateur | `backend`, `runtime` |
| `runtime` | runtime Go embarqué dans la sortie | `host`, `support` |
| `FFI` | frontière étrangère | `Ffi`, `foreign` |
| `ADT` | type de données algébrique | `Adt` |
| `JSON` | format texte | `Json` |
| `codegen` | génération de code, en un mot | `CodeGen`, `Codegen` |

Précisions :

- **`build` n'est pas `compile`.** `Driver.Build` construit des
  `BackendModule` avec PBO ; `Driver.compile` exécute toutes les phases.
- **`emit` n'écrit rien.** L'émission produit un `GoFile` en mémoire ;
  `Driver.Output` l'écrit, et la session de cache le publie.
- **`worker` n'est pas `job`.** Un worker est une fonction Go générée ; un job
  est une unité de travail parallèle. Le cache conserve des signatures de
  fonctions ou de workers, jamais des « jobs ».
- **`protocol` n'est pas `cache`.** Le protocole est le format et les
  transactions ; le cache est l'ensemble des décisions de réutilisation.
- **`host` n'est pas `runtime`.** Le host exécute le compilateur ; le runtime
  est du Go embarqué dans l'application générée.
- **`metadata` n'est pas `state`.** Les métadonnées viennent du TAST et ne
  changent pas ; l'état est mutable pendant une passe.

## Règles de nommage

### Modules et fichiers

- Le nom d'un module est exactement son chemin relatif sous `src/`, sans
  `.purs` : `src/Gopurs/Host/Snapshot.purs` déclare `Gopurs.Host.Snapshot`.
  Aucune exception, aucun alias de module temporaire.
- Un module cible s'écrit `Gopurs.<Phase>.<Domaine>` ou
  `Gopurs.<Phase>.<Domaine>.<Rôle>`. Les phases sont `Driver`, `Prepare`,
  `Rewrite`, `Emit`, `Cache`, `Host` et `Runtime`.
- Les fichiers et dossiers sont en `PascalCase` et suivent les segments du
  module. Un module à FFI garde ses compagnons `.go`, `.js` et `.rs` à côté du
  `.purs`, avec le même nom de base.
- Un module qui n'a pas de domaine clair mélange probablement deux
  responsabilités : le découper plutôt que l'appeler `Common`.

### Suffixes et mots interdits

- **Pas de suffixe `Exprs`.** Le dossier `Emit/Expr/` porte déjà le concept :
  `Emit.Expr.Adt`, pas `Emit.Expr.AdtExprs`.
- **Pas de `Utils`, `Common`, `Helpers`, `Manager` ni `Misc`** dans un nouveau
  nom. Ces mots cachent une responsabilité au lieu de la nommer.
- **`Analysis`** est réservé aux faits en lecture seule. Une transformation
  s'appelle `Rewrite` ou `Pass` ; une vérification s'appelle `Check`.
- **`State`** est réservé au mutable (`Ref`, buffer, cache en mémoire).
  **`Metadata`** est immuable. **`Config`** vient des options.
- **`Snapshot`** est une capture immuable ; **`Session`** est une durée de vie
  avec staging ; **`Protocol`** est un format de stockage.
- Un nouveau nom de module développe les mots (`Constructor`, `Expression`,
  `Declaration`) plutôt que d'inventer une abréviation.

### Identifiants

- **Fonctions** : verbe à l'infinitif (`prepare`, `specialize`, `rewrite`,
  `emit`, `publish`). Les variantes injectables utilisent `With...` ; les FFI
  utilisent `...Impl`.
- **Booléens** : `is*`, `has*`, `should*`. Pas de paramètre positionnel booléen
  dans une nouvelle API publique ; préférer un type de mode.
- **Preuves** : nommer la propriété établie (`covered`, `complete`,
  `permitsLiteralEdit`, `sameTast`), pas le mécanisme qui la calcule.
- **Acronymes** : `FFI`, `ADT`, `AST`, `JSON`, `Go`, `JS`, `Rust` en
  majuscules. Jamais `Ffi`, `Adt` ou `Json` dans un nouveau nom.
- **« Worker »** est réservé aux fonctions Go générées. La concurrence utilise
  `job` ; le cache utilise `signatures` ou `functions`, pas `workers`.
- **Identifiants Go générés** : jamais renommés. Ils sont un contrat public de
  la sortie, pas un détail interne du compilateur.

## Noms historiques conservés jusqu'au lot 3

Les formes suivantes restent dans le code jusqu'à la PR de leur lot, soit
parce qu'elles violent une règle ci-dessus, soit parce que leur chemin change.
Elles restent valides dans les tests et les documents normatifs tant que le
déplacement n'est pas mergé. **Aucun nouveau module ne doit reprendre ces
formes.**

| Forme historique | Modules actuels | Cible | Lot de retrait |
| --- | --- | --- | --- |
| `Ffi*` | `FfiBridge`, `FfiBridge.Render`, `FfiBridge.Signatures`, `FfiBridge.TypeSupport`, `FfiBridge.Values`, `FfiTypes`, `FfiSupport` | `Emit.FFI.*` | 3e |
| `Adt*` | `AdtMetadata`, `AdtExprs` | `Prepare.Metadata.Adt`, `Emit.Expr.Adt` | 3c, 3e |
| `*Exprs` | `AdtExprs`, `BindingExprs`, `CallExprs`, `ConstructorExprs`, `ControlExprs`, `EffectExprs`, `FunctionExprs`, `LiteralExprs`, `PrimitiveExprs`, `RecordExprs`, `TypedExprs` | `Emit.Expr.*` | 3e |
| `CodeGen`, `CodegenState` | `CodeGen`, `CodegenState` (et le type `CodegenMetadata`) | `Emit`, `Emit.State` | 3e |
| `WorkerNames` | `WorkerNames` | `Rewrite.Names` | 3d |
| `Go*` | `GoAst`, `GoCode`, `GoSource`, `GoSymbol`, `GoTypes`, `GoImports`, `GoConversions`, `GoFunctions` | `Emit.Go.*`, `Emit.Workers.Parameters` | 3e |
| `Printer*` | `Printer`, `Printer.Builder` | `Emit.Go.Printer*` | 3e |
| `*Workers` | `LocalWorkers`, `ModuleWorkers` | `Emit.Workers.*` | 3e |
| `Preparation` | `Preparation` | `Prepare.Jobs` | 3c |
| `Monomorphization*` | `Monomorphization`, `Monomorphization.ForeignForwarders` | `Prepare.Specialize*` | 3c |
| `Emission` | `Emission` | `Driver.Scheduler` | 3 (pilote) |
| `BuildCache` | `BuildCache` | `Cache.Protocol` | 3b |
| `Driver.Cache` | `Driver.Cache` | `Cache.Session` | 3b |
| `Driver.Prepare*` | `Driver.Prepare`, `Driver.Prepare.Persistent`, `Driver.Prepare.Projection`, `Driver.Prepare.ProjectionCapsule`, `Driver.Prepare.Proof` | `Prepare*` | 3c |

Le lot 3a a déjà déplacé `CompilerSnapshot` en `Host.Snapshot`, `Diagnostics`
en `Host.Diagnostics`, `Metrics` en `Host.Metrics`, `Resident` en
`Host.Resident` et `Runtime` en `Runtime.Source`. Ces cinq noms ne sont plus
des exceptions : ils suivent les règles ci-dessus.

Le plan ne fixe pas la PR qui déplace `Emission` : sa cible est
`Driver.Scheduler`, et elle peut rejoindre une PR du pilote ou une PR dédiée
sans changer son comportement. Les autres lots sont ceux de l'annexe C de
[cleanup.md](../cleanup.md).

Les modules déjà conformes conservent leur nom quand ils changent de dossier :

- `Driver`, `Driver.Config`, `Driver.Build`, `Driver.Manifest`,
  `Driver.Output` ;
- `Cache.Modules`, `Cache.Wire`, `Cache.Pool`, `Cache.TypeMemo`,
  `Cache.BindingCapsule`, `Cache.HeaderPreload`, `Cache.LookupTable` ;
- `Host.Snapshot`, `Host.Diagnostics`, `Host.Metrics`, `Host.Resident`,
  `Runtime.Source` (déplacés au lot 3a) ;
- `ThunkFusion`, `FunctionFusion`, `ImmediateApplications`,
  `ClosedDictionaries`, `BorrowedObjects`, `DecoderSchemas`, `Ownership`,
  `NativeRecordArgs`, `GlobalTypes`, `ClassMetadata`, `ConstructorMetadata`,
  `ConstructorLayout`, `ReboxMetadata`, `ModuleBindings`,
  `ModuleDeclarations`, `ExprContext`, `ExprAnalysis`, `ArrayIntrinsics`,
  `ArrayTraverse`, `ObjectTraverse`, `Int32Loops`, `Main`.

`DependencyGraph` a été supprimé au lot 1 ; ce n'est pas un nom historique à
réintroduire. La correspondance complète, module par module, est dans
[module-map.md](module-map.md).

## Application aux documents

- **Normatif** : `README.md`, `CONTRIBUTING.md`, `docs/README.md`,
  `docs/architecture-overview.md`, `docs/architecture.md`, `docs/naming.md`,
  `docs/module-map.md` et les contrats d'architecture. Ils emploient le
  vocabulaire cible et mettent à jour leurs liens quand un module bouge.
- **Référence** : les documents d'une optimisation peuvent citer le nom
  historique tant que le module n'a pas bougé, avec le nom cible en regard.
- **Historique** : `docs/history/`, `build-*.md` et les journaux de campagne
  sont figés. Ils gardent leurs noms et leurs liens d'origine.

Les documents publics sont en anglais ; l'architecture interne et ce glossaire
sont en français. Ne pas mélanger les deux langues dans une même page, sauf
citation d'un nom de code.
