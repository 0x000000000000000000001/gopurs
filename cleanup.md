# Plan de nettoyage — lisibilité et maintenabilité

> **Statut** : proposition, 29 octobre 2026.
> **Base** : `release/0.1.0-nix`, arbre propre, commit `8361c05`.
> **Portée** : noms, découpage, documentation, outillage. Aucun changement de
> comportement, de représentation ou de performance.
> Ce plan succède au plan v1 de maintenabilité, clos le 2 octobre 2026
> (lot 15). Il ne rouvre pas les chantiers de performance déjà tranchés.

**Résumé** : sept lots progressifs (environ 20 jours). D'abord une baseline et
un glossaire, puis le rangement par phase (`Driver`, `Prepare`, `Rewrite`,
`Emit`, `Cache`, `Host`, `Runtime`), ensuite le découpage des gros modules, la
documentation et les garde-fous CI. Chaque lot est mécanique et vérifié par
l'identité octet par octet des sorties Go.

**État au 29 octobre 2026** (branche `cleanup/v2`) :

- **Lot 0 — fait** : baseline consignée dans
  [docs/history/cleanup-baseline.md](docs/history/cleanup-baseline.md),
  `tools/check-output-identity.mjs` et ses tests ajoutés.
- **Lot 1 — partiel** : `.gitattributes` corrigé, `DependencyGraph` supprimé,
  commentaire de `BuildCache` corrigé, index `docs/`, `tests/` et `tools/`
  ajoutés, en-têtes FFI manuscrites ajoutés. L'archivage des journaux
  `build-*` reste groupé avec le lot 5 pour ne pas casser les liens deux fois.
- **Lot 2 — fait** : [docs/naming.md](docs/naming.md),
  [docs/module-map.md](docs/module-map.md), alignement de `README.md`,
  `CONTRIBUTING.md`, `docs/architecture.md`, `docs/architecture-overview.md`
  et `docs/README.md`, plus `tools/check-docs.mjs` et son script npm. La carte
  couvre les **104 modules actuels** : les 105 de la baseline moins
  `DependencyGraph`, supprimé au lot 1 ; elle documente cette soustraction.
- **Lot 3a — préparé sur `cleanup/v2`, intégration sur main autorisée le 1 novembre** :
  `CompilerSnapshot` → `Host.Snapshot`, `Diagnostics` → `Host.Diagnostics`,
  `Metrics` → `Host.Metrics`, `Resident` → `Host.Resident`, `Runtime` →
  `Runtime.Source`. Les compagnons FFI, `tools/embed-runtime.mjs`,
  `.gitignore`, les outils et les documents normatifs suivent. Quatre
  fixtures donnent **1 220 fichiers Go identiques** avant/après. L’intégration
  a été autorisée après la publication de RC7, avant la future
  version stable. Elle conserve les correctifs de release de main ; les
  preuves du 29 octobre restent historiques et la CI rapide de la fusion
  constitue un contrôle distinct.
- **Lots 3b à 7 — à faire**.

### Vérification de l’intégration — 1 novembre 2026

La fusion de `cleanup/v2` (`0220540`) avec le main de publication (`2f071e0`)
conserve les corrections worker/Manifest et les pins de bibliothèques de main.
Le cache temporaire Nix ajouté au dernier commit de cleanup est exclu de
l’arbre intégré et ignoré ; la branche et la copie utilisateur sont préservées.

Deux bundles JavaScript reconstruits sur macOS arm64 ont compilé et exécuté
`ClassRefSyntax` avec le même frontend TAST local de développement et les mêmes
bibliothèques isolées : **332 fichiers de sortie Go identiques**. Ce contrôle
ponctuel ne remplace ni les campagnes de fixtures ni la qualification native
et multiplateforme du futur candidat. Aucun snapshot de référence n’est modifié.

Les contrôles ciblés couvrent l’identité des sorties, les documents, le runtime
embarqué et le snapshot du compilateur. Le validateur documentaire a été corrigé
pour conserver les tirets consécutifs des ancres GitHub après suppression de la
ponctuation (régression rouge puis verte). Le contrôle réel vérifie 37 documents,
104 modules et les fichiers distribués ; le paquet npm contient 59 fichiers.
Les liens de glossaire/carte dans les documents distribués pointent vers GitHub.

Ces quatre suites sont désormais requises par la CI rapide et son reçu
`linux-quick-v2` (24 gates). Les anciens reçus et l’omission d’une de ces suites
sont rejetés. La qualification CI reste distincte des preuves locales ci-dessus.

## 1. Objectif

Réduire le temps nécessaire pour répondre à trois questions :

1. **Où est la responsabilité X ?** (moins de 5 minutes)
2. **Quel est le chemin d'une compilation ?** (une page, pas 1 000 lignes)
3. **Puis-je modifier X sans casser Y ?** (frontières de modules et de tests
   explicites)

Le critère de réussite n'est pas « moins de lignes » : c'est moins
d'indirection mentale pour un mainteneur qui arrive sans contexte.

### Non-objectifs

- Pas de changement de sémantique, de format de sortie ou de cache.
- Pas de changement de représentation Go, d'ordre d'évaluation, de strictness.
- Pas de renommage des identifiants Go générés (ils sont un contrat public).
- Pas de reformatage massif, pas de réécriture opportuniste.
- Pas de traduction des journaux historiques.
- Pas de suppression des hôtes JS, Go ou Rust.

### Ce que « sans sacrifier les perfs » impose

- Aucun `newtype`, wrapper ou champ ajouté dans un chemin chaud.
- Aucune traversée supplémentaire, aucun `Map` intermédiaire ajouté.
- Les extractions de modules gardent les mêmes définitions et le même ordre
  d'appel ; on déplace, on ne réimplémente pas.
- Le code Go généré doit rester **identique octet par octet** sur le corpus de
  référence. Si un déplacement change la sortie, le lot est arrêté, pas le
  snapshot mis à jour.
- Les mesures de compilation sont comparées au bruit près (±2 %), avec le
  protocole existant ; aucune conclusion à partir d'un seul run.

## 2. Diagnostic

### 2.1 Mesures

| Constat | Mesure |
| --- | --- |
| Modules PureScript | **105** (15 486 lignes), dont 63 à la racine de `src/Gopurs` |
| Modules > 200 lignes | **21** ; > 400 lignes : **4** |
| Modules exposant tout implicitement | **2** (`GoAst`, `Printer`) |
| Modules avec FFI en trois hôtes | **15** (dont `Main`), soit 45 compagnons `.go`/`.js`/`.rs` |
| Documentation | **68** fichiers (59 Markdown), **17 987** lignes de Markdown, dont **36** `build-*` |
| Gros documents | `testing.md` 3 858 l., `architecture.md` 1 037 l., `README.md` 850 l. |
| Outils | **148** `.mjs` à la racine de `tools/`, dont **99** tests |
| Snapshots | **405** en Release, **405** en Dev, **9** paquets |
| Références en dur à `src/Gopurs` | **57**, dans **13** outils et **7** documents |

### 2.2 Arborescence plate

Les 63 modules à la racine de `src/Gopurs` mélangent cinq étapes
(`Prepare`, `Rewrite`, `Emit`, `Cache`, `Host`) et des familles de données.
Un nouvel arrivant ne peut pas déduire l'ordre des passes depuis l'arborescence ;
il doit lire `CodeGen` et `Driver` pour reconstituer le pipeline.

### 2.3 Vocabulaire ambigu

| Mot | Sens actuels | Effet |
| --- | --- | --- |
| `Prepare` / `Preparation` | phase de préparation ; moteur de travaux parallèles | on croit que `Preparation` est la phase |
| `Emit` / `Emission` / `Output` / `Build` | production de code ; ordonnanceur de lots ; écriture de fichiers ; construction PBO | quatre mots pour trois étapes différentes |
| `Worker` | fonction Go générée ; unité de parallélisme ; instantané de signatures en cache | impossible de savoir de quel « worker » parle un commentaire |
| `Cache` | protocole de stockage ; session ; cache de modules ; codecs | quatre modules, quatre sens |
| `Analysis` | analyse d'appel ; analyse d'expression ; analyse d'ownership | pas de frontière claire entre fait, preuve et transformation |
| `CodeGen` / `CodegenState` | même concept, deux casses | incohérence de nommage visible |
| `Ffi` / `ADT` / `JSON` | `FfiBridge`, `AdtMetadata`, `Json` | casses d'acronymes non uniformes |
| `*Exprs` | `AdtExprs`, `CallExprs`, `TypedExprs`… (11 modules) | suffixe pluriel qui ne dit pas la responsabilité |

### 2.4 Modules-dieux et état implicite

| Module | Lignes | Symptôme |
| --- | --- | --- |
| `Cache.Modules` | 913 | 26 exports, `ModuleCache` à **27 champs**, `initialize` à **8 arguments positionnels** |
| `Cache.Wire` | 887 | codecs et instances de domaines différents dans un seul fichier |
| `Driver.Prepare.Persistent` | 445 | préchargement, projection et publication mêlés |
| `Printer` | 440 | rendu des expressions, déclarations et types dans un seul module |
| `CodeGen` | 387 | importe presque toutes les familles d'expression et compose sept passes |

`CodeGen` est le point de couture du backend : c'est pratique pour compiler,
mais c'est le pire point d'entrée pour comprendre. Il doit devenir un
orchestrateur lisible, pas un sac d'imports.

### 2.5 Multi-hôte et code généré

- Chaque module à FFI existe en quatre exemplaires : `.purs`, `.go`, `.js`,
  `.rs`. Les trois implémentations doivent rester alignées ; rien ne le
  vérifie automatiquement.
- `.gitattributes` masque **tous** les `.purs` des statistiques GitHub et
  marque **tous** les `.js` comme générés. Un visiteur ne voit donc pas le
  langage réel du projet, et les FFI JS manuscrites sont présentées comme
  générées.
- `Runtime.js/go/rs` est généré par `tools/embed-runtime.mjs` ; les autres
  compagnons sont manuscrits. Cette distinction n'est écrite nulle part.
- 13 outils contiennent des chemins `src/Gopurs/...` en dur ; tout mouvement
  de module doit les mettre à jour.

### 2.6 Documentation

- Pas d'index `docs/README.md` : 68 fichiers, plusieurs points d'entrée
  concurrents.
- `architecture.md` mélange l'architecture normative et le récit de campagnes
  de performance (chiffres, candidats, pourcentages). Le mainteneur ne sait
  pas ce qui est contractuel.
- `testing.md` mélange commandes, contrats de test et journal historique de
  plus de 3 800 lignes.
- Les documents publics sont en anglais, l'architecture interne en français,
  l'historique dans les deux langues. Ce n'est pas un problème en soi, mais
  la règle n'est pas énoncée.
- `package.json` liste des documents un par un dans `files` : tout déplacement
  casse la publication npm sans que les tests le voient.

### 2.7 Dette ponctuelle

- `DependencyGraph` n'est importé par personne.
- `BuildCache` porte un commentaire « when cache reuse is wired » alors que le
  cache est branché.
- `bin/` contient 100 Mo de binaires locaux ignorés par Git ; rien n'indique
  dans le dépôt lesquels sont des artefacts et lesquels sont des lanceurs.

## 3. Vocabulaire cible

Règle directrice : **un concept, un mot ; un mot, un concept**. Les phases sont
des verbes, les données des noms, les preuves des adjectifs.

### 3.1 Glossaire normatif

| Terme | Sens unique retenu | À éviter |
| --- | --- | --- |
| `compile` | exécution complète du pilote | `build`, `run` pour le pilote |
| `driver` | orchestration, options, phases | `main`, `pipeline` |
| `prepare` | chargement TAST, métadonnées, spécialisation, cache de préparation | `preprocess`, `setup` |
| `specialize` | monomorphisation | `instantiate`, `mono` |
| `rewrite` | transformation PureScript/IR avant émission | `optimize` pour nos passes |
| `emit` | production en mémoire des déclarations Go | `generate`, `output` |
| `output` | écriture des fichiers de sortie | `publish`, `emit` |
| `publish` | validation et commit du cache | `output`, `write` |
| `build` | construction PBO d'un `BackendModule` | `compile` |
| `worker` | fonction Go générée (locale ou de module) | unité de parallélisme |
| `job` | unité de travail parallèle | `worker`, `task` |
| `protocol` | format et transactions du stockage | `cache` seul |
| `session` | durée de vie d'un build et son staging | `cache`, `store` |
| `snapshot` | capture immuable à un instant donné | `state` |
| `metadata` | faits immuables dérivés du TAST | `state`, `info` |
| `state` | état mutable d'une passe | `metadata`, `context` |
| `host` | JS, Go ou Rust qui exécute le compilateur | `backend`, `runtime` |
| `runtime` | runtime Go embarqué | `host`, `support` |
| `FFI` | frontière étrangère | `Ffi`, `foreign` |
| `ADT` | type de données algébrique | `Adt` |
| `JSON` | format texte | `Json` |
| `codegen` | génération de code, en un mot | `CodeGen`, `Codegen` |

### 3.2 Règles de nommage

- **Modules** : `Gopurs.<Phase>.<Domaine>` ou `Gopurs.<Phase>.<Domaine>.<Rôle>`.
  Le nom du module doit être exactement le chemin relatif, sans exception.
- **Pas de suffixe `Exprs`** : le dossier `Emit/Expr/` porte déjà le concept.
  On écrit `Emit.Expr.Adt`, pas `Emit.Expr.AdtExprs`.
- **Pas de `Utils`, `Common`, `Helpers`, `Manager`, `Misc`** dans un nouveau
  nom. Si un module n'a pas de domaine, c'est souvent qu'il mélange deux
  responsabilités.
- **`Analysis`** est réservé aux faits en lecture seule. Une transformation
  s'appelle `Rewrite` ou `Pass` ; une vérification s'appelle `Check`.
- **`State`** est réservé au mutable (`Ref`, buffer, cache en mémoire).
  **`Metadata`** est immuable. **`Config`** vient des options.
- **Fonctions** : verbe à l'infinitif (`prepare`, `specialize`, `rewrite`,
  `emit`, `publish`). Les variantes injectables utilisent `With...` ;
  les FFI utilisent `...Impl`.
- **Booléens** : `is*`, `has*`, `should*`. Pas de paramètre positionnel
  booléen dans une nouvelle API publique ; préférer un type de mode.
- **Acronymes** : `FFI`, `ADT`, `AST`, `JSON`, `Go`, `JS`, `Rust`. Jamais
  `Ffi`, `Adt`, `Json` dans un nouveau nom.
- **« Worker »** : réservé aux fonctions Go générées. La concurrence utilise
  `job` ; le cache utilise `signatures` ou `functions`, pas `workers`.

## 4. Arborescence cible

```text
src/
├── Main.purs
└── Gopurs/
    ├── Driver.purs                 # orchestration, options, ordre des phases
    ├── Driver/
    │   ├── Config.purs
    │   ├── Build.purs              # construction PBO, signatures
    │   ├── Scheduler.purs          # lots, fibres, contre-pression (ex-Emission)
    │   ├── Manifest.purs
    │   └── Output.purs             # écriture staging/runtime/entrées
    ├── Prepare.purs                # phase : TAST -> modules préparés
    ├── Prepare/
    │   ├── Jobs.purs               # ex-Preparation
    │   ├── Metadata/               # Adt, Class, Constructor, GlobalTypes, Rebox, Layout
    │   ├── Specialize/             # ex-Monomorphization + ForeignForwarders
    │   └── Persistent/             # ex-Driver.Prepare.Persistent + Projection, Capsule, Proof
    ├── Rewrite.purs                # composition ordonnée des passes (si extraite)
    ├── Rewrite/
    │   ├── ThunkFusion.purs
    │   ├── FunctionFusion.purs
    │   ├── ImmediateApplications.purs (+ Scope)
    │   ├── ClosedDictionaries.purs (+ Admission, Scope)
    │   ├── BorrowedObjects.purs
    │   ├── DecoderSchemas.purs (+ Admission, Programs, Source, Types, Workers)
    │   ├── Ownership.purs (+ Analysis, Candidates, Types, Workers)
    │   ├── NativeRecordArgs.purs (+ Projection, Source)
    │   └── Names.purs              # ex-WorkerNames
    ├── Emit.purs                   # ex-CodeGen : un module -> Go
    ├── Emit/
    │   ├── State.purs              # ex-CodegenState
    │   ├── Module/                 # Bindings, Declarations
    │   ├── Workers/                # Local, Module, Parameters (ex-GoFunctions)
    │   ├── Expr/                   # Context, Analysis, Adt, Binding, Call, ...
    │   ├── Go/                     # Ast, Code, Source, Symbol, Types, Imports,
    │   │                           # Conversions, Printer
    │   └── FFI/                    # Bridge, Types, Support
    ├── Cache/
    │   ├── Session.purs            # ex-Driver.Cache
    │   ├── Protocol.purs           # ex-BuildCache
    │   └── ...                     # Modules, Wire, Pool, TypeMemo, ...
    ├── Host/                       # Snapshot, Diagnostics, Metrics, Resident
    └── Runtime/Source.purs         # ex-Runtime (runtime Go embarqué)
```

Points structurants :

- `Driver` ne contient plus la préparation ni le cache de session : il
  orchestre des phases qui ont leur propre dossier.
- `Prepare` possède le chargement, les métadonnées, la spécialisation et la
  persistance. `Rewrite` possède les passes Go. `Emit` possède la traduction
  et l'impression. Un module ne traverse jamais deux de ces frontières.
- `Emit.purs` reste le point d'entrée lisible ; il ne doit plus importer
  directement les onze familles d'expressions, mais passer par
  `Emit.Expr` (dispatch) et `Rewrite`.
- `Cache` ne dépend plus de `Driver` pour sa session ; `Driver` dépend de
  `Cache`.
- `Host` isole le code dépendant de Node/Go/Rust du backend pur.

## 5. Lots

Chaque lot est indépendant, réversible et peut s'arrêter sans laisser le
dépôt dans un état intermédiaire. Un lot = une PR (ou une petite série de PR
pour le lot 3), jamais un mélange déplacement + logique.

### Lot 0 — Baseline et garde-fous (0,5 j)

**But** : pouvoir prouver que rien ne change.

- Figer un corpus de référence (TAST + sorties Go) et consigner les empreintes
  dans `docs/history/cleanup-baseline.md`.
- Rejouer la compilation avec l'hôte JS, l'hôte natif séquentiel et l'hôte
  natif parallèle ; conserver les trois arbres de sortie.
- Mesurer trois charges de compilation (petite, moyenne, b8x) avec le protocole
  existant ; consigner machine, commandes et médianes.
- Ajouter `tools/check-output-identity.mjs` : compare deux arbres de sortie
  fichiers par fichiers (chemins + SHA-256) et échoue sur toute différence.

**Critère de sortie** : baseline reproductible, aucune modification du dépôt
hors documentation.

### Lot 1 — Hygiène et vérité du dépôt (1 j)

**But** : enlever ce qui trompe le lecteur, sans toucher au code.

- Supprimer `DependencyGraph` (inutilisé) ou documenter pourquoi il reste.
- Corriger le commentaire obsolète de `BuildCache`.
- Réparer `.gitattributes` :
  - ne plus masquer les `.purs` ;
  - marquer comme générés uniquement `bin/gopurs.js`, `output/`,
    `src/Gopurs/Runtime.{js,go,rs}` et les artefacts de build ;
  - laisser les FFI `.go`/`.js`/`.rs` manuscrites visibles.
- Créer `docs/README.md` (index et statut), `tests/README.md`,
  `tools/README.md`.
- Ajouter un en-tête de trois lignes aux FFI manuscrites :
  `// Handwritten FFI for Gopurs.X — keep in sync with the .purs foreign imports.`
  (ou l'équivalent JS/Rust).
- Déplacer les 36 `build-*` et les journaux de campagne dans
  `docs/history/`, avec un index et des liens de redirection depuis les
  anciens chemins si nécessaire.

**Critère de sortie** : `git diff --check` propre, tous les liens locaux
résolus, `npm pack --dry-run` inchangé fonctionnellement.

### Lot 2 — Glossaire et alignement documentaire (1,5 j)

**But** : figer le vocabulaire avant de déplacer le code.

- Créer `docs/naming.md` à partir de la section 3, avec les exceptions
  historiques explicitement listées.
- Mettre à jour `docs/architecture.md`, `docs/architecture-overview.md`,
  `CONTRIBUTING.md` et `README.md` pour employer les termes cibles.
- Ajouter une section « vocabulaire » dans `CONTRIBUTING.md` et un lien vers
  `docs/naming.md`.
- Écrire `docs/module-map.md` : une ligne par module, sa phase, sa
  responsabilité, ses dépendances autorisées. Ce document devient la source
  de vérité pour le lot 3.

**Critère de sortie** : aucun nouveau terme ambigu dans les documents
normatifs ; la carte des modules couvre les 105 modules.

### Lot 3 — Déplacement par phase (6 à 8 j, 5 PR)

**But** : l'arborescence raconte le pipeline.

Découpage des PR, dans cet ordre (du moins dépendant au plus dépendant) :

1. **Host + Runtime** : `CompilerSnapshot` → `Host.Snapshot`,
   `Diagnostics` → `Host.Diagnostics`, `Metrics` → `Host.Metrics`,
   `Resident` → `Host.Resident`, `Runtime` → `Runtime.Source`.
   Mettre à jour `tools/embed-runtime.mjs`, `.gitignore`, les imports de tests
   et les 57 références en dur.
2. **Cache** : `Driver.Cache` → `Cache.Session`, `BuildCache` →
   `Cache.Protocol`. Déplacer les compagnons FFI et mettre à jour
   `tools/build-cache*.test.mjs`, `cache-context-runtime.mjs`.
3. **Prepare** : `Driver.Prepare*` → `Prepare*`, `Preparation` →
   `Prepare.Jobs`, métadonnées → `Prepare.Metadata`, `Monomorphization` →
   `Prepare.Specialize`.
4. **Rewrite** : les neuf passes et `WorkerNames` → `Rewrite.Names`.
5. **Emit** : `CodeGen` → `Emit`, `CodegenState` → `Emit.State`, `Go*` →
   `Emit.Go.*`, `Ffi*` → `Emit.FFI.*`, expressions →
   `Emit.Expr.*`, workers → `Emit.Workers.*`. Cette PR peut être scindée en
   Go / Expr / FFI si elle dépasse 40 fichiers.

Règles pour chaque PR :

- `git mv` du `.purs` **et** de ses trois compagnons FFI dans le même commit.
- Mise à jour du `module` et de tous les imports ; pas de module alias
  temporaire.
- Mise à jour de `docs/module-map.md`, des liens `src/Gopurs/...` et de
  `package.json.files` si un document publié bouge.
- Reconstruction JS **et** native, puis comparaison octet par octet avec la
  baseline du lot 0.
- Exécution des tests ciblés, puis de la suite complète.

**Critère de sortie par PR** :

- sorties Go identiques à la baseline ;
- 401 fixtures Release et 401 Dev, aucun snapshot modifié ;
- `node --test tools/*.test.mjs` vert ;
- compilation native séquentielle/parallèle et JS : mêmes fichiers ;
- aucun chemin `src/Gopurs/<ancien>` restant.

### Lot 4 — Décomposer les modules-dieux (3 à 5 j)

**But** : des modules qu'on peut lire d'une traite, sans changer les API
appelantes.

- `Cache.Modules` : extraire, derrière une façade `Cache.Modules` conservée :
  - `Cache.Modules.Index` : types, clés, `initialize` ;
  - `Cache.Modules.Lookup` : lookups, lectures, preuves ;
  - `Cache.Modules.Emit` : émission et restauration ;
  - `Cache.Modules.Report` : rapport, implémentations, signatures.
- `Cache.Wire` : séparer les codecs génériques des instances de domaine
  (`Wire.Module`, `Wire.Metadata`, `Wire.Syntax`) ; conserver `Wire` comme
  façade et vérifier que les instances restent importées.
- `Printer` : `Printer.Expr`, `Printer.Decl`, `Printer.Type`, avec `Printer`
  comme façade ; `Printer.Builder` inchangé.
- `Driver.Prepare.Persistent` : séparer préchargement, projection et
  publication.
- Optionnel : extraire `Rewrite.purs` de `CodeGen.prepareModule` pour que
  l'ordre des passes soit lisible en un seul endroit. C'est le seul lot avec
  un léger changement d'orchestration ; il doit être validé par la baseline.

**Critère de sortie** : aucun appelant modifié (façades), snapshots
identiques, `npm run test:cache*` vert, taille du bundle compilateur dans le
bruit.

### Lot 5 — Documentation (3 j)

**But** : une page pour comprendre, une page pour chaque contrat, l'histoire
à part.

- `docs/README.md` : tableau `Document | Public | Statut | Dernière
  vérification`, avec quatre entrées : Démarrer, Référence, Interne,
  Histoire.
- Découper `docs/architecture.md` en :
  - `architecture/overview.md` (le pipeline en une page) ;
  - `architecture/pipeline.md` (ordre des passes et provenance) ;
  - `architecture/cache.md` ;
  - `architecture/emission.md` ;
  - `architecture/ffi.md`.
  Déplacer les récits de campagne dans `history/`.
- Découper `docs/testing.md` en `testing/commands.md` (le tableau de
  sélection), `testing/contracts.md` (les contrats) et
  `testing/history.md` (les lots).
- Réduire `README.md` à 200–250 lignes : pitch, statut, installation,
  quickstart, carte de la documentation, développement. Déplacer les
  benchmarks détaillés dans `docs/history/benchmarks.md`.
- Règle de langue : documents publics en anglais ; architecture normative en
  anglais ; journaux historiques conservés tels quels. Ne pas traduire
  l'histoire.
- Ajouter `tools/check-docs.mjs` : liens relatifs, ancres, fichiers
  `package.json.files`, et absence de lien vers un document déplacé.

**Critère de sortie** : un nouveau venu lit `README` puis
`architecture/overview.md` et sait où intervenir ; 0 lien cassé.

### Lot 6 — Tests et outils (2 j)

**But** : retrouver un test ou un script sans lire 148 noms.

- Créer `tools/{build,release,ffi,dev}/` pour les scripts non-test ;
  **garder les `*.test.mjs` à la racine de `tools/`** pour ne pas casser
  `node --test tools/*.test.mjs`.
- Mettre à jour `package.json` : chemins, scripts `test`, `test:fast`,
  `check:docs`, `check:naming` ; conserver les scripts granulaires.
- `tests/README.md` : rôle de `passing/`, `passing-snapshots/`,
  `passing-snapshots-dev/`, `passing-package-snapshots/`, `host/` ; procédure
  de mise à jour d'un snapshot ; convention de nommage
  (`<Fonctionnalité><Variante>`, pas de numéro d'issue pour les nouveaux).
- Ne pas renommer les snapshots existants (405 × 2) : ajouter un index et des
  alias documentés. Renommer les dossiers `*-dev` casserait les scripts pour
  un gain faible.
- Documenter les fichiers générés vs manuscrits dans `tools/README.md`.

**Critère de sortie** : `npm test` existe et passe ; `bin/test` et
`bin/modtest` inchangés ; `npm pack` contient les nouveaux chemins.

### Lot 7 — Garde-fous et onboarding (2 j)

**But** : empêcher la dérive après le nettoyage.

- CI :
  - `check:docs` (liens, `package.json.files`) ;
  - `check:naming` (nouveaux fichiers seulement : `Ffi`, `Adt`, `Json`,
    `Exprs`, `Utils`, `Manager`) ;
  - `check:modules` (chemin = nom de module, module importé ou entrée) ;
  - `check:ffi` (les `foreign import` du `.purs` ont un équivalent dans les
    trois hôtes) ;
  - identité des sorties sur un petit corpus ;
  - suite complète sur Linux, native sur macOS.
- Exercice d'onboarding chronométré, consigné dans
  `docs/history/onboarding.md` :
  1. trouver où est produite une erreur de la CLI ;
  2. ajouter une fixture et son snapshot ;
  3. changer un message de diagnostic sans casser les snapshots.
- Checklist PR (annexe B) ajoutée à `CONTRIBUTING.md`.

**Critère de sortie** : les trois exercices sont réalisables en moins de
30 minutes chacun par un mainteneur qui n'a jamais lu le backend.

## 6. Validation

### Par lot

| Contrôle | Commande | Seuil |
| --- | --- | --- |
| Build JS | `npm run build` | vert, deux fois de suite |
| Build natif | `npm run build:native -- --keep-workspace` | vert |
| Tests unitaires | `node --test tools/*.test.mjs` | 0 échec |
| CLI / cache | `npm run test:cli`, `npm run test:cache*` | 0 échec |
| Fixtures | `./bin/test` | 401 fixtures, 0 snapshot modifié |
| Dev | `./bin/test -d` | 401 fixtures, 0 snapshot modifié |
| Identité | `node tools/check-output-identity.mjs baseline current` | 0 fichier différent |
| Liens | `node tools/check-docs.mjs` | 0 lien cassé |
| Format | `git diff --check` | propre |

### Global (fin du plan)

- `find src/Gopurs -maxdepth 1 -name '*.purs'` ne retourne que les modules
  d'entrée de phase (`Driver.purs`, `Prepare.purs`, `Rewrite.purs`,
  `Emit.purs`) — ou aucun, si les entrées rejoignent leur dossier.
- Tout module est à son chemin, avec un nom qui figure dans
  `docs/module-map.md`.
- Aucun module ne mélange deux phases ; tout module > 400 lignes a un
  en-tête qui justifie ses responsabilités.
- `Cache.Modules` < 20 exports, `ModuleCache` découpé en structures nommées
  (pas 27 champs anonymes).
- Sorties Go identiques à la baseline sur le corpus de référence.
- Benchmarks de compilation dans le bruit (±2 %) sur les trois charges.
- `README` + `architecture/overview.md` suffisent à comprendre le pipeline.
- Les trois exercices d'onboarding passent.

## 7. Risques et parades

| Risque | Parade |
| --- | --- |
| Un déplacement change les sorties Go | Baseline octet par octet ; arrêt du lot, pas de mise à jour de snapshot |
| Les noms de modules changent les clés de cache utilisateur | Planifier après la release 0.1.0 ; sinon incrémenter la version de schéma du cache |
| FFI introuvable après déplacement | Déplacer `.purs`, `.go`, `.js`, `.rs` dans le même commit ; `check:ffi` |
| Bootstrap natif cassé | Reconstruire JS et natif, comparer aux binaires de référence, mettre à jour les checksums |
| Liens et `package.json.files` cassés | `check:docs` avant merge ; `npm pack --dry-run` |
| Régression de perf | Aucune logique modifiée ; mesures avant/après ; pas de wrapper dans les chemins chauds |
| PR géante impossible à relire | Une phase par PR ; plafond indicatif de 40 fichiers ; scinder `Emit` si besoin |
| Dérive opportuniste | Checklist PR : aucun changement de comportement, aucun reformatage |
| Conflits avec la branche release | Branche `cleanup/v2` après le tag ; rebase fréquent ; lots courts |

## 8. Hors périmètre (backlog)

- Suppression d'un hôte (JS, Go ou Rust) : décision produit, pas nettoyage.
- Remplacement d'Argonaut ou du format de cache.
- `newtype` pour les identifiants et les unités.
- Renommage des identifiants Go générés.
- Traduction des journaux historiques.
- Réécriture du système d'effets ou de l'ordonnanceur.

## Annexe A — Correspondance des modules

Préfixe `Gopurs.` omis. Les sous-modules non listés suivent leur parent.

| Cible | Source actuelle |
| --- | --- |
| `Driver` | `Driver` |
| `Driver.Config` | `Driver.Config` |
| `Driver.Build` | `Driver.Build` |
| `Driver.Scheduler` | `Emission` |
| `Driver.Manifest` | `Driver.Manifest` |
| `Driver.Output` | `Driver.Output` |
| `Prepare` | `Driver.Prepare` |
| `Prepare.Jobs` | `Preparation` |
| `Prepare.Persistent` | `Driver.Prepare.Persistent` |
| `Prepare.Persistent.Projection` | `Driver.Prepare.Projection` |
| `Prepare.Persistent.ProjectionCapsule` | `Driver.Prepare.ProjectionCapsule` |
| `Prepare.Persistent.Proof` | `Driver.Prepare.Proof` |
| `Prepare.Metadata.Adt` | `AdtMetadata` |
| `Prepare.Metadata.Class` | `ClassMetadata` |
| `Prepare.Metadata.Constructor` | `ConstructorMetadata` |
| `Prepare.Metadata.ConstructorLayout` | `ConstructorLayout` |
| `Prepare.Metadata.GlobalTypes` | `GlobalTypes` |
| `Prepare.Metadata.Rebox` | `ReboxMetadata` |
| `Prepare.Specialize` | `Monomorphization` |
| `Prepare.Specialize.ForeignForwarders` | `Monomorphization.ForeignForwarders` |
| `Rewrite.ThunkFusion` | `ThunkFusion` |
| `Rewrite.FunctionFusion` | `FunctionFusion` |
| `Rewrite.ImmediateApplications` | `ImmediateApplications` |
| `Rewrite.ClosedDictionaries` | `ClosedDictionaries` |
| `Rewrite.BorrowedObjects` | `BorrowedObjects` |
| `Rewrite.DecoderSchemas` | `DecoderSchemas` |
| `Rewrite.Ownership` | `Ownership` |
| `Rewrite.NativeRecordArgs` | `NativeRecordArgs` |
| `Rewrite.Names` | `WorkerNames` |
| `Emit` | `CodeGen` |
| `Emit.State` | `CodegenState` |
| `Emit.Module.Bindings` | `ModuleBindings` |
| `Emit.Module.Declarations` | `ModuleDeclarations` |
| `Emit.Workers.Local` | `LocalWorkers` |
| `Emit.Workers.Module` | `ModuleWorkers` |
| `Emit.Workers.Parameters` | `GoFunctions` |
| `Emit.Expr.Context` | `ExprContext` |
| `Emit.Expr.Analysis` | `ExprAnalysis` |
| `Emit.Expr.Adt` | `AdtExprs` |
| `Emit.Expr.Binding` | `BindingExprs` |
| `Emit.Expr.Call` | `CallExprs` |
| `Emit.Expr.Call.Analysis` | `CallAnalysis` |
| `Emit.Expr.Call.Arguments` | `CallArguments` |
| `Emit.Expr.Constructor` | `ConstructorExprs` |
| `Emit.Expr.Control` | `ControlExprs` |
| `Emit.Expr.Effect` | `EffectExprs` |
| `Emit.Expr.Function` | `FunctionExprs` |
| `Emit.Expr.Literal` | `LiteralExprs` |
| `Emit.Expr.Primitive` | `PrimitiveExprs` |
| `Emit.Expr.Record` | `RecordExprs` |
| `Emit.Expr.Typed` | `TypedExprs` |
| `Emit.Expr.ArrayIntrinsics` | `ArrayIntrinsics` |
| `Emit.Expr.ArrayTraverse` | `ArrayTraverse` |
| `Emit.Expr.ObjectTraverse` | `ObjectTraverse` |
| `Emit.Expr.Int32Loops` | `Int32Loops` |
| `Emit.Go.Ast` | `GoAst` |
| `Emit.Go.Code` | `GoCode` |
| `Emit.Go.Source` | `GoSource` |
| `Emit.Go.Symbol` | `GoSymbol` |
| `Emit.Go.Types` | `GoTypes` |
| `Emit.Go.Imports` | `GoImports` |
| `Emit.Go.Conversions` | `GoConversions` |
| `Emit.Go.Printer` | `Printer` |
| `Emit.FFI.Bridge` | `FfiBridge` |
| `Emit.FFI.Types` | `FfiTypes` |
| `Emit.FFI.Support` | `FfiSupport` |
| `Cache.Session` | `Driver.Cache` |
| `Cache.Protocol` | `BuildCache` |
| `Cache.Modules` | `Cache.Modules` |
| `Cache.Wire` | `Cache.Wire` |
| `Cache.Pool` | `Cache.Pool` |
| `Cache.TypeMemo` | `Cache.TypeMemo` |
| `Cache.BindingCapsule` | `Cache.BindingCapsule` |
| `Cache.HeaderPreload` | `Cache.HeaderPreload` |
| `Cache.LookupTable` | `Cache.LookupTable` |
| `Host.Snapshot` | `CompilerSnapshot` |
| `Host.Diagnostics` | `Diagnostics` |
| `Host.Metrics` | `Metrics` |
| `Host.Resident` | `Resident` |
| `Runtime.Source` | `Runtime` |
| `Main` | `Main` |

Modules supprimés : `DependencyGraph` (aucun import).

## Annexe B — Checklist de PR de nettoyage

- [ ] Le lot ne change ni comportement, ni sortie, ni format de cache.
- [ ] Un seul type de changement : déplacement, renommage, extraction ou doc.
- [ ] Les FFI `.go`/`.js`/`.rs` ont suivi leur `.purs`.
- [ ] `docs/module-map.md` et `docs/naming.md` sont à jour.
- [ ] `git diff --check` est propre.
- [ ] `npm run build` et `npm run build:native -- --keep-workspace` passent.
- [ ] `node --test tools/*.test.mjs` passe.
- [ ] `./bin/test` et `./bin/test -d` ne changent aucun snapshot.
- [ ] `tools/check-output-identity.mjs` ne voit aucune différence.
- [ ] `npm pack --dry-run` inclut les bons fichiers.
- [ ] Aucune référence à l'ancien chemin ne subsiste
      (`grep -rn 'src/Gopurs/<ancien>'`).
- [ ] Le benchmark de compilation est dans le bruit, ou la mesure est
      expliquée.

## Annexe C — Ordre recommandé

```text
Lot 0  baseline
Lot 1  hygiène
Lot 2  glossaire + carte des modules
Lot 3a Host + Runtime
Lot 3b Cache
Lot 3c Prepare
Lot 3d Rewrite
Lot 3e Emit
Lot 4  modules-dieux
Lot 5  documentation
Lot 6  tests + outils
Lot 7  CI + onboarding
```

Ne pas commencer le lot 3 avant le tag `0.1.0` et avant que le lot 2 ait
stabilisé le vocabulaire. Les lots 5 et 6 peuvent avancer en parallèle du
lot 3, à condition de ne pas modifier les mêmes fichiers.

<!--
Note pour la revue : ce plan privilégie les déplacements mécaniques et les
façades. Toute extraction qui modifierait l'ordre des passes, les clés de
cache ou les représentations doit être traitée comme un changement de
comportement, avec sa propre baseline et sa propre revue.
-->
