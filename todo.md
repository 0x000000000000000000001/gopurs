# Rendre gopurs maintenable

Objectif : pouvoir modifier une responsabilité à la fois, comprendre ses entrées et ses sorties, et vérifier rapidement le résultat. Les points ci-dessous sont ordonnés ; chaque case correspond à une modification ou une vérification indépendante.

Constats au 8 septembre 2026 : `CodeGen.purs` compte 3 546 lignes, `Main.purs` 524, `Runtime.purs` 1 008 et `bin/test` 317. Ces tailles orientent le découpage ; elles ne prouvent pas de défaut de comportement.

## Principes de travail

- Commencer chaque chantier par une preuve courte : références, fixture existante, sortie générée ou reproduction isolée. Découper toute investigation qui ne peut pas donner un premier résultat en moins de cinq minutes.
- Pour une extraction, conserver le comportement et comparer le Go généré. Traiter séparément toute correction ou optimisation découverte en chemin.
- Préserver le TAST enrichi : `dataDecls`, `classDecls`, `ann.type`, `TypeApp`, `ForAll`, contraintes, ordre des champs et queues de rangées. Les types disponibles doivent rester accessibles aux décisions de représentation native.
- Valider avec les fixtures concernées et leurs snapshots ; ajouter un cas seulement si un invariant utile manque. Réserver la suite complète aux jalons qui touchent plusieurs familles.
- Une extraction doit conserver les snapshots existants. Toute différence doit être expliquée avant leur mise à jour ; vérifier aussi les sorties absentes, que le runner actuel crée même avec `UPDATE_SNAPSHOTS=0`.
- Pour un changement susceptible de modifier les performances, comparer avant/après sur les mêmes entrées et se référer aux baselines officielles de `../../altbak.pub/README.md`. Aucun gain déduit d'un lancement isolé.
- Cocher une étape avec une preuve concise et la portée réelle de la validation.

## 1. Établir une référence de travail reproductible

- [x] **1.1 — Identifier les outils réellement utilisés.** Chemins résolus, versions exécutées et révisions Git relevés le 8 septembre 2026 ; voir le relevé ci-dessous. Le shell, le build npm et altbak sélectionnent des outils différents. Aucun build ni test exécuté.
- [x] **1.2 — Vérifier un cycle court.** Le 8 septembre 2026, `./bin/test NativeRecordSizes -c`, puis la même fixture sans `-c`, réussissent avec le `purs` TAST sur le `PATH` et `UPDATE_SNAPSHOTS=0`. Les 30 assertions passent à chaque run ; les deux snapshots restent identiques aux références. Voir le relevé ci-dessous.
- [ ] **1.3 — Définir les contrôles par famille.** Associer types/records à `NativeRecordBoxing` et `NativeRecordSizes`, FFI à `FFIIntegerReturns`, appels à `CurriedLambdas`, tableaux à `ArrayRoundtrip`, récursion à `TCO`/`TCOMutRec`, fusion à `ThunkFusion`. Vérifier leur état initial par petits groupes. **Types/records, FFI, appels et ArrayRoundtrip vérifiés ; sortie de TCO corrigée, snapshots TCO/TCOMutRec encore non validés ; fusion encore à exécuter.**
- [ ] **1.4 — Consigner les limites initiales.** Distinguer échecs existants, exclusions et contrôles non exécutés. Conserver les sorties nécessaires aux comparaisons suivantes hors des sources de production.

### Relevé 1.1 — Outils et révisions

Les chemins ci-dessous sont résolus après suivi des liens symboliques. Versions obtenues avec `--version`, ou `go version`.

| Outil et contexte | Version observée | Exécutable résolu |
| --- | --- | --- |
| `purs` du shell | `0.15.15` | `/opt/homebrew/lib/node_modules/purescript/purs.bin` |
| `purs` local de gopurs, prioritaire pendant `npm run build` | `0.15.16` ; paquet `purescript-npm` : `0.15.16-0x1` | `/Users/0x1/Documents/htdocs/gopurs/gopurs/node_modules/purescript-npm/purs` |
| `purs` du fork TAST sélectionné par altbak | `0.15.16 [development build; commit: 59b80f2a183647c6d4e99f06b5658ff550f1cb68 DIRTY]` | `/Users/0x1/Documents/htdocs/purescript/.stack-work/dist/aarch64-osx/ghc-9.8.4/build/purs/purs` |
| Spago du shell | `1.0.3` | `/opt/homebrew/lib/node_modules/spago/bin/bundle.js` |
| Spago sélectionné par altbak | `1.0.4` | `/Users/0x1/Documents/htdocs/altbak.pub/run/bak/js/node_modules/spago/bin/bundle.js` |
| Node, commun aux contextes relevés | `v24.8.0` | `/opt/homebrew/Cellar/node/24.8.0/bin/node` |
| Go, commun aux contextes relevés | `go1.27.0 darwin/arm64` | `/opt/homebrew/Cellar/go/1.27.0/libexec/bin/go` |

Le shell résout ces commandes via `/opt/homebrew/bin`. `bin/test` hérite du `PATH` ; son option `-c` appelle `npm run build`, où le `node_modules/.bin/purs` de gopurs devient prioritaire. Il n'existe pas de Spago local dans gopurs : celui du `PATH` reste sélectionné. `altbak.pub/bin/go/run` préfixe le `PATH` avec `/Users/0x1/Documents/htdocs/altbak.pub/run/bak/js/node_modules/.bin`, qui fournit le fork TAST et Spago 1.0.4. Le build npm de gopurs conserve toutefois son propre `purs` local en priorité. Ces résolutions ont été contrôlées avec les chemins correspondants et des appels de version uniquement.

Pour l'étape 1.2, le `purs` TAST à sélectionner est celui exposé par le répertoire `.bin` d'altbak ci-dessus ; le `purs` 0.15.15 du shell n'est pas ce binaire. Aucun `PATH` persistant n'a été modifié. Le suffixe `DIRTY` est celui rapporté par le binaire du fork, pas une mesure de l'état actuel de ses sources. Son SHA-256 est `a9162efcb4256de3e1564c6ab766b0a2638795c711e21865fd477135c8f4020b` ; celui du `purs` local de gopurs est `ae313bebeb0c150dd71d94a95e6654966806e4911602dfc1e5b6329162aaa9b9`.

| Dépôt | Racine réelle | Révision HEAD | Branche | État au relevé |
| --- | --- | --- | --- | --- |
| gopurs | `/Users/0x1/Documents/htdocs/gopurs/gopurs` | `324b9326e47ef271a9d9b93c1c385415a55f1167` | `edge` | Seul `todo.md` modifié |
| PBO déclaré dans `spago.yaml` | `/Users/0x1/Documents/htdocs/purescript-backend-optimizer-gopurs` | `67ba2151e3717b27a13e95c25d3b25c8bdc645cc` | `edge-gopurs` | Propre |

Le chemin PBO déclaré est bien `../../purescript-backend-optimizer-gopurs`. Vérification par résolution du chemin, `git rev-parse --show-toplevel HEAD`, `git symbolic-ref -q --short HEAD` et `git status --short` dans chaque dépôt.

### Relevé 1.2 — Cycle ciblé NativeRecordSizes

Commandes exécutées depuis la racine de gopurs, avec le `PATH` limité à ces processus :

```bash
export PATH="/Users/0x1/Documents/htdocs/altbak.pub/run/bak/js/node_modules/.bin:$PATH"
export UPDATE_SNAPSHOTS=0
./bin/test NativeRecordSizes -c
./bin/test NativeRecordSizes
```

Le premier run reconstruit `bin/gopurs.js` puis nettoie les caches du runner. Le build rapporte zéro erreur et zéro avertissement. Chaque run compare `Main.go` et `Main_ffi.go` aux snapshots existants, compile le Go et exécute les 30 assertions : 31 lignes utiles, terminées par `Done`, puis `1 passed, 0 failed`. Les sorties sont identiques octet pour octet.

Produits relevés : bundle `bin/gopurs.js`, 196 fichiers Go sous `tests/runner/output/` (dont `purescript/Main.go`, `purescript/Main_ffi.go`, `gopurs_runtime/runtime.go`, `main/main.go` et `Main/main/main.go`), `go.mod` et exécutable `gopurs_main`. Les empreintes SHA-256 des 196 fichiers Go et du bundle sont identiques entre les deux runs. Les deux snapshots, `tests/runner/spago.yaml`, son lockfile et tous les fichiers suivis hors `todo.md` conservent leurs empreintes initiales.

La première tentative s'est arrêtée avant compilation sur l'accès au cache SQLite de Spago hors sandbox ; les deux runs réussis ont utilisé l'accès autorisé aux caches Spago et Go. Aucun changement du compilateur, du runner ou des fixtures. Validation limitée à cette fixture en Go ; aucune exécution JavaScript ni suite complète. Logs, sorties et inventaires conservés dans `/private/tmp/gopurs-step-1-2-c6z31szu/` (`clean-retry.log`, `repeat.log`, `first-output-files.txt`, empreintes avant/après).

### Relevé 1.3 — Types et records

Le 8 septembre 2026, `./bin/test NativeRecordBoxing` réussit sans `-c`, avec le même `PATH` qu'en 1.2 et `UPDATE_SNAPSHOTS=0`. Le bundle est identique à celui vérifié en 1.2. Le Go généré correspond octet pour octet à `tests/passing-snapshots/NativeRecordBoxing.go`, puis les 10 assertions passent : 11 lignes utiles terminées par `Done`, bilan `1 passed, 0 failed`.

Le snapshot vérifié conserve le chemin ciblé : `Call_Main_consumeEntry` reçoit un record Go natif, le boxe avec `RecordDict2`, puis le transmet au consommateur opaque via `Apply`. La fixture couvre les mises à jour des champs et la conservation des versions précédentes. Avec les 30 assertions de `NativeRecordSizes` validées en 1.2, la référence types/records est établie.

Snapshot, configuration et lockfile du runner, bundle et fichiers suivis inchangés après exécution. Seul le présent compte rendu est modifié. Aucune autre fixture ni exécution JavaScript lancée. Preuves conservées dans `/private/tmp/gopurs-step-1-3-boxing-3g6_jrno/` : `run.log`, `stdout.txt`, `Main.go` et `verification.json`. L'étape 1.3 reste ouverte pour les autres familles.

### Relevé 1.3 — FFI

Le 8 septembre 2026, `./bin/test FFIIntegerReturns` réussit sans `-c`, avec le même `PATH` qu'en 1.2 et `UPDATE_SNAPSHOTS=0`. Le bundle conserve l'empreinte vérifiée en 1.2. Les deux fichiers Go générés correspondent octet pour octet à `FFIIntegerReturns.go` et `FFIIntegerReturns_ffi.go`, puis les 27 assertions passent : sortie `Done`, bilan `1 passed, 0 failed`.

La fixture couvre les retours `int64`/`int`, leurs tableaux vides et non vides, les bornes Int PureScript, les fallbacks `any` numériques/String/Boolean et un consommateur opaque lu depuis une `Ref`. Les wrappers vérifiés utilisent `Int` pour les entiers natifs et chaque élément des tableaux, et conservent `Box` pour les retours `any`.

Snapshots, configuration et lockfile du runner, bundle et fichiers suivis inchangés après exécution. Seul le présent compte rendu est modifié. Aucune autre fixture ni exécution JavaScript lancée. Preuves dans `/private/tmp/gopurs-step-1-3-ffi-gn1z8kpv/` : `run.log`, `stdout.txt`, `Main.go`, `Main_ffi.go` et `verification.json`.

### Relevé 1.3 — Appels

Le 8 septembre 2026, `./bin/test CurriedLambdas` réussit sans `-c`, avec le même `PATH` qu'en 1.2 et `UPDATE_SNAPSHOTS=0`. Le bundle conserve l'empreinte vérifiée en 1.2. Le Go généré correspond octet pour octet à `tests/passing-snapshots/CurriedLambdas.go`, puis les 52 assertions passent : sortie `Done`, bilan `1 passed, 0 failed`.

La fixture couvre les applications partielles réutilisées, l'ordre des arguments jusqu'à six, l'ordre et l'exécution différée des effets, les frontières `let`/branche/récursion/appel/`Fn2`/`TypeApp` et la capture d'état lors de l'exécution de l'effet.

Snapshot, configuration et lockfile du runner, bundle et fichiers suivis inchangés après exécution. Seul le présent compte rendu est modifié. Aucune autre fixture ni exécution JavaScript lancée. Preuves dans `/private/tmp/gopurs-step-1-3-calls-g9w2b5n5/` : `run.log`, `stdout.txt`, `Main.go` et `verification.json`.

### Relevé 1.3 — Tableaux

Le 8 septembre 2026, `./bin/test ArrayRoundtrip` réussit sans `-c`, avec le même `PATH` qu'en 1.2 et `UPDATE_SNAPSHOTS=0`. Le bundle conserve l'empreinte vérifiée en 1.2. Le Go généré correspond octet pour octet à `tests/passing-snapshots/ArrayRoundtrip.go`, puis les 28 assertions passent : 29 lignes utiles terminées par `Done`, bilan `1 passed, 0 failed`.

La fixture couvre 14 cas de tableaux (vide, singletons, parité, signes, doublons, bornes Int32), 9 plages inclusives ascendantes ou descendantes et 5 cas du pipeline `range → filter → fold`, dont `n=900 → 202950`. Les sommes intermédiaires restent dans Int32. Ce contrôle établit la référence de cette fixture ; il ne mesure pas les performances et ne constitue pas une validation étendue de toutes les optimisations de tableaux.

Snapshot, configuration et lockfile du runner, bundle et fichiers suivis conservent leurs empreintes d'avant ce run. Seul le présent compte rendu est modifié. Aucune autre fixture ni exécution JavaScript lancée. Preuves dans `/private/tmp/gopurs-step-1-3-arrays-oa2sxhuy/` : `run.log`, `stdout.txt`, `Main.go` et `verification.json`.

### Relevé 1.3 — Récursion : dépendances isolées, référence non validée

Le 8 septembre 2026, les premiers essais de `./bin/test TCO` et `./bin/test TCOMutRec` échouent avant l'appel à gopurs sur `EscapedSkolem` dans `gopurs-foreign-object/src/Foreign/Object.purs:203:14`. Sans directive `@dependencies`, ces fixtures utilisaient la liste générale du runner. Logs initiaux conservés dans `/private/tmp/gopurs-step-1-3-recursion-207hcm10/`.

Après accord pour isoler leurs dépendances, ajout des seules directives suivantes, sans changement des algorithmes ni du runner :

- `TCO` : `-- @dependencies: prelude effect console arrays tailrec`.
- `TCOMutRec` : `-- @dependencies: assert prelude effect console`.

Les deux commandes sont relancées séparément sans `-c`, avec le même `PATH` qu'en 1.2 et `UPDATE_SNAPSHOTS=0`. La compilation PureScript et la génération Go réussissent désormais, mais chaque commande s'arrête sur un écart de snapshot (`0 passed, 1 failed`). Les snapshots sont conservés. Pour `TCO`, le diff comprend notamment des fonctions spécialisées et des changements de représentation ; il ne se limite pas à des renommages.

Le Go de chaque fixture est ensuite compilé et exécuté séparément, sans actualiser les snapshots, pour vérifier le comportement :

- **`TCO` : sortie incorrecte.** Sortie obtenue `0, 1, 2, 3, 4, 0, 42, Done`, contre `0, 1, 2, 3, 4, 10000, 42, Done` attendu. `length (span (\_ -> true) (1..10000)).init` donne `0` au lieu de `10000`. Le processus termine avec le code 0, cette fixture affichant ses résultats sans assertions ; la comparaison explicite détecte l'échec. Cause localisée dans la conversion de `Nothing` à la frontière FFI, voir l'isolation ci-dessous.
- **`TCOMutRec` : comportement vérifié, snapshot non conforme.** Les 8 assertions passent, sortie `Done`, code de sortie 0. Les tests de débordement de pile restent commentés ; ce succès ne prouve pas une pile constante pour tous les cas.

Les empreintes des deux snapshots, du bundle (identique à celui de 1.2), de la configuration et du lockfile du runner sont inchangées. Les seuls changements de sources sont les deux directives de dépendances ; le présent compte rendu est également mis à jour. Aucune correction du générateur ni exécution JavaScript. Preuves conservées dans `/private/tmp/gopurs-step-1-3-recursion-deps-w6w0mjua/` : logs du runner, sous-dossiers `TCO/` et `TCOMutRec/` avec Go généré, diffs, binaires et sorties, puis `verification.json`. La référence récursion reste non validée.

### Relevé 1.3 — Isolation de findIndex / span

Le 8 septembre 2026, une reproduction indépendante compare JavaScript et Go sur des tableaux de 0, 1, 3 et 10000 éléments. Le défaut apparaît déjà avec `findIndex (\_ -> false) []` : JavaScript retourne `Nothing`, Go retourne `Just 0`. Les tailles des entrées sont correctes. `span (\_ -> true) [1]` donne ainsi les longueurs `0,1` en Go contre `1,0` en JavaScript.

Deux traces, ajoutées uniquement au Go généré de cette copie, montrent que la valeur fournie comme `nothing` arrive au helper FFI avec le tag 9, l'identifiant ADT 930809136 et un pointeur non nul. Le helper ne trouve aucun élément et renvoie cette valeur ; `span` la traite comme `Just 0`. La génération convertit `Box(&Constructor_Data_Maybe_Nothing[Value]{})` via `CoerceToStruct[Constructor_Data_Maybe_Just[Value]]`, alors que le consommateur attend un pointeur nul pour `Nothing`. L'échec est localisé dans la conversion entre représentations de Maybe avant l'entrée FFI.

Le bundle, les sources, les snapshots et la configuration du runner sont inchangés. Aucun correctif du générateur ni benchmark appliqué. Les sorties Go restent identiques avec et sans traces. Reproduction, comparaison, traces et points de génération à examiner conservés dans `/private/tmp/gopurs-span-isolation-2f_emf1y/RESULTS.md`. La prochaine correction doit préserver la distinction `Nothing` / `Just 0` ; la référence récursion reste non validée.

### Relevé 1.3 — Correction des conversions de Maybe

Le 8 septembre 2026, correction limitée à `CodeGen.purs` : la conversion de la forme valeur de Maybe émet la représentation canonique (identifiant `hashString "Data_Data_Maybe_Just"`, pointeur nul pour Nothing, pointeur vers Just avec son contenu sinon). La conversion inverse utilise le même identifiant. Le passage à un pointeur typé réutilise le reboxing existant pour convertir le contenu de Just depuis `Value` vers son type natif.

La nouvelle fixture `MaybeFfiRoundtrip` vérifie 24 assertions : Nothing, Just 0, Just 7, allers-retours via Ref, recherche sans résultat ou avec résultat, et span vide, complet, partiel et sur 10000 éléments. Un producteur non inliné et des entrées lues via Ref exercent aussi les conversions de retour de fonction. La première version de la fixture échoue avec le bundle antérieur (`Nothing` devient `Just 0`). Les cas supplémentaires ont exposé l'identifiant erroné et une réinterprétation du contenu (`Just 0` devenait `Just 6` après correction du seul identifiant). La version finale donne les mêmes sorties en Go et JavaScript, sans instrumentation.

Validation finale : `npm run build` réussit sans avertissement ; `bin/test MaybeFfiRoundtrip -c` passe avec son nouveau snapshot, puis `NativeRecordBoxing` (10 assertions), `ArrayRoundtrip` (28) et `FFIIntegerReturns` (27) passent avec leurs références inchangées. `UPDATE_SNAPSHOTS=0` est conservé. Le nettoyage `-c` est nécessaire ici : un premier essai sur le runner chaud référençait un alias Show d'un module résiduel (`Data.Interval`) dans le nouveau snapshot, absent de la reproduction isolée.

`bin/test TCO` reste en échec sur son snapshot antérieur. Son Go est compilé et exécuté séparément : sortie exacte `0, 1, 2, 3, 4, 10000, 42, Done`, désormais correcte. Aucun snapshot existant n'est modifié ; seul celui de `MaybeFfiRoundtrip` est ajouté. Configuration et lockfile du runner préservés. Les snapshots de récursion restent à examiner ; TCOMutRec et ThunkFusion ne sont pas exécutés durant cette correction. Aucun benchmark exécuté.

Preuves, expériences intermédiaires, sorties JavaScript/Go, logs finaux et empreintes conservés dans `/private/tmp/gopurs-maybe-fix-8wgjs445/`. L'étape 1.3 reste ouverte.

## 2. Retrouver un arbre de sources lisible

Nettoyage effectué le 8 septembre 2026 dans gopurs et le worktree PBO utilisé par son build ; voir le relevé ci-dessous.

- [x] **2.1 — Classer les fichiers suspects.** Imports, appels, scripts, FFI et points d’entrée examinés dans gopurs, PBO-Go et leurs consommateurs locaux. Les candidats sans usage établi ont été retirés.
- [x] **2.2 — Sauver les expériences utiles.** Les fixtures actives et les outils FFI restent en place ; aucun essai orphelin examiné ne nécessite de promotion en nouvel outil. Les anciens essais retirés restent disponibles dans l’historique Git.
- [x] **2.3 — Retirer les reliquats prouvés inutiles.** Sauvegardes, rejets, patchs de diagnostic, essais isolés et anciennes sorties `output-test` retirés après recherche de leurs consommateurs. Les tests et snapshots actifs sont préservés.
- [x] **2.4 — Clarifier les chemins alternatifs.** `CodeGenBackend`, `OptimizeTAST`, `QuoteTAST`, `UsageAnalysis` et `FBIP` supprimés : aucun import ni point d’entrée actif ne les utilisait. `QuoteTAST` était consommé uniquement par `OptimizeTAST`.
- [x] **2.5 — Prévenir le retour des artefacts.** Exclusions ciblées ajoutées aux deux `.gitignore` pour les sauvegardes, rejets et sorties identifiées. Imports et calculs locaux purs inutilisés retirés des modules concernés.

### Relevé — Nettoyage gopurs et PBO-Go

Périmètre : `gopurs/gopurs` et `purescript-backend-optimizer-gopurs`, worktree désigné par son `spago.yaml`. Les autres worktrees PBO ne sont pas modifiés.

- **gopurs :** 892 lignes nettes de sources et d’essais retirées, 24 fichiers supprimés. Cela inclut les chemins et modules sans appelant, les helpers et calculs purs inutilisés, les références globales jamais lues, les expériences isolées et les sauvegardes. Les dépendances directes `exceptions` et `parallel` étaient redondantes avec celles d’`aff` ; les 69 packages résolus gardent les mêmes versions et sources. Le lockfile actualise aussi les dépendances déjà déclarées par le PBO local.
- **PBO :** 229 lignes nettes de sources et d’essais retirées, dont le module Debug orphelin, trois décodeurs privés inutilisés de `CoreFn.Json`, des helpers de diagnostic et deux branches TypeApp dupliquées. Les 993 fichiers générés d’`output-test` (environ 20 Mo), ainsi que les anciens patchs et sauvegardes, n’avaient aucun consommateur dans les sources, configurations ou scripts.
- **Usages particuliers vérifiés :** les outils FFI Go/Wasm, les entrées ES et Rust, les tests configurés, les imports dynamiques de `test/typeapp.mjs` et les constructeurs enregistrés par réflexion dans `Cache.js` sont conservés. Les calculs de contexte `LetRec` qui lisent des `Ref` via `unsafePerformEffect` ne sont pas supprimés sur le seul critère d’un résultat ignoré.

Validation : gopurs compile, ainsi que les deux packages PBO (`backend-optimizer`, `backend-es`) avec le compilateur npm utilisé par gopurs. Des avertissements de paramètres ignorés et de noms masqués restent présents. `NativeRecordBoxing`, `FFIIntegerReturns` et `ArrayRoundtrip` passent avec leurs snapshots inchangés ; les 173 fichiers Go des deux premières fixtures sont identiques aux références avant nettoyage. Les 18 tests TypeApp passent sur le PBO compilé par gopurs.

`CurriedLambdas` et `MaybeFfiRoundtrip` s’arrêtent sur des écarts de snapshots, y compris avec `-c`. Leur Go est compilé et exécuté séparément : les assertions passent. Un bundle reconstruit à partir des HEAD avant nettoyage produit exactement les mêmes 195 fichiers Go sur les mêmes entrées pour chacune des deux fixtures. Ces écarts ne sont donc pas introduits par le nettoyage ; aucun snapshot n’est actualisé. Ils restent à prendre en compte pour terminer la référence 1.3.

Preuves, manifestes des suppressions, builds, comparaisons et sorties : `/private/tmp/gopurs-pbo-cleanup-ucl3t3ja/`.

## 3. Ramener Main à l'orchestration du build

Le build dispose maintenant d’un seul chemin actif : `main` appelle `emitModule`. La préparation des métadonnées et le rôle du cache restent à clarifier.

- [x] **3.1 — Déterminer le chemin de référence.** La CLI construit et exécute `Main.main`. Aucun appelant de `runBuild` trouvé ; son helper `takeMonomorphizedModules`, qui mutait le record préparé, était utilisé uniquement par ce chemin abandonné.
- [x] **3.2 — Extraire l'émission d'un module.** Le 8 septembre 2026, extraction du callback de `main` dans `emitModule`, avec les métadonnées préparées, le dossier FFI et les deux modules en paramètres. Les alias devenus inutiles dans `main` sont retirés. Compilation sans avertissement ; `NativeRecordBoxing` (10 assertions) et `FFIIntegerReturns` (27) passent avant et après. Tous les fichiers Go générés sont identiques octet par octet (86 et 87 fichiers respectivement), ainsi que les trois snapshots existants. `runBuild` reste inchangé. Preuves : `/private/tmp/gopurs-emit-module-d_1b62n0/`.
- [x] **3.3 — Unifier le build.** `runBuild`, son import FFI et `Main.js` supprimés. Le chemin actif passe par `main` et `emitModule`. Les sorties Go, y compris les points d’entrée, restent identiques sur les comparaisons avant/après.
- [x] **3.4 — Extraire la préparation des métadonnées.** Les préparations des types globaux, des constructeurs, des classes, des représentations ADT et de la monomorphisation ont leurs modules dédiés (3.4.1–3.4.5). `Main` coordonne leurs appels.
- [x] **3.4.1 — Clarifier la construction des types globaux.** Le 9 septembre 2026, `Gopurs.GlobalTypes` expose uniquement `buildGlobalTypes`. Le parcours des modules et groupes de définitions, le choix des annotations et l'ajout des FFI ont des fonctions nommées ; les helpers restent privés. `Main` importe la construction de la table. Build sans avertissement ; `FFIIntegerReturns` passe avant/après (27 assertions). Les 60 entrées CoreFn et les 87 fichiers Go générés sont identiques, ainsi que les snapshots et la configuration du runner. Preuves : `/private/tmp/gopurs-global-types-7ao6ks2h/`.
- [x] **3.4.2 — Clarifier la préparation des constructeurs.** Le 9 septembre 2026, `Gopurs.ConstructorMetadata` expose `buildConstructorTypes`, `collectElidedConstructors` et l'alias `ConstructorTypes`. Les parcours des modules, déclarations et constructeurs ont des helpers privés ; `Main` appelle les deux fonctions avant l'ajout des classes synthétiques. Les champs TAST, l'ordre des parcours et les deux conventions de nommage restent identiques. Build sans avertissement. Sur `GenericsRep`, les anciens et nouveaux bundles relancés sur les mêmes 58 entrées CoreFn produisent les mêmes 82 fichiers Go ; compilation Go réussie et sortie d'exécution identique. La première comparaison avait un écart dans `Data_Eq.go`, également reproduit par l'ancien bundle : les deux régénérations concordent. Le runner complet rencontre déjà `EscapedSkolem` dans `Foreign.Object` avant refactorisation ; la comparaison utilise une copie temporaire de la fixture avec les seules dépendances `prelude effect console`. Son snapshot présente déjà un écart avant refactorisation, identique après ; aucun snapshot ni configuration du runner n'est modifié. Preuves : `/private/tmp/gopurs-constructor-metadata-pgxun7l1/`.
- [x] **3.4.3 — Clarifier la préparation des classes.** Le 9 septembre 2026, `Gopurs.ClassMetadata` expose `buildClassFields`, `addClassDataDeclarations` et l'alias `ClassFields`. Un helper privé construit et trie les champs des superclasses et méthodes pour les deux représentations ; noms, types, variables et ordre restent identiques. `Main` appelle ces fonctions au même stade de préparation. La table inutilisée `classDeclsMap` et l'argument ignoré correspondant de `translate` sont supprimés. Build réussi ; 89 avertissements dans des lignes inchangées de `CodeGen`, aucun dans le nouveau module. Pour chacune des fixtures `Superclasses1` et `TypeClassMemberOrderChange`, les 58 entrées CoreFn et 82 fichiers Go sont identiques avant/après ; le Go compile et affiche respectivement `21.0` et `true`, puis `Done`. Les deux snapshots ont déjà un écart avant refactorisation, strictement identique après. Comparaison via des copies temporaires avec les dépendances `prelude effect console` ; snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-class-metadata-asnbjm2m/`.
- [x] **3.4.4 — Clarifier les représentations des ADT.** Le 9 septembre 2026, `Gopurs.AdtMetadata` expose `buildPointerAdtMetadata`, `buildEnumAdtMetadata` et leurs types de sortie. Les critères de sélection et le nommage des structs ont des helpers privés ; les motifs de constructeur unique remplacent les accès `unsafeIndex`, et l'absence de feuille utilise `Maybe`. `Main` reçoit les cinq tables depuis les modules enrichis avec les classes synthétiques. Règles, ordre des parcours et collisions conservés. Build sans avertissement. Pour chacune des fixtures `CaseStatement` et `RBTree`, les 58 entrées CoreFn et 82 fichiers Go sont identiques avant/après ; le Go compile et affiche respectivement `Done` et une profondeur de `22` pour 100 000 insertions. Les écarts avec les snapshots existaient avant la modification et restent identiques. Comparaison via des copies temporaires avec les dépendances `prelude effect console` ; snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-adt-metadata-80se99et/`.
- [x] **3.4.5 — Isoler la préparation de la monomorphisation.** Le 9 septembre 2026, `Gopurs.Monomorphization` expose uniquement `monomorphizeModules`. L'index des définitions, la collecte des FFI, le filtrage et `hasTypeVariables` ont des helpers privés ; l'ordre collecte, propagation transitive, filtrage puis monomorphisation reste identique. `Main` fournit les types globaux d'origine et les modules enrichis avec les classes. L'appel inutilisé à `collectAllTypes`, la table `adtTypes` et le transport d'`instantiations` vers `CodeGen` sont supprimés ; les spécialisations restent locales à la passe. Build réussi, 89 avertissements sur des lignes inchangées de `CodeGen`, aucun dans le nouveau module. `VisibleTypeApplications` et `FFIIntegerReturns` produisent respectivement les mêmes 82 et 87 fichiers Go sur 58 et 60 entrées CoreFn identiques avant/après. Le Go compile et les exécutions aboutissent à `Done`, dont les 27 assertions FFI. Les écarts avec les snapshots existent avant la modification et restent identiques. `VisibleTypeApplications` utilise une copie temporaire avec les seules dépendances `prelude effect console` ; la fixture FFI utilise ses dépendances déclarées. Snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-monomorphization-9qvd255w/`.
- [x] **3.5 — Rendre le cache compréhensible.** Le 9 septembre 2026, l'écriture des `.gopurs-cache.json`, sans lecteur actif dans gopurs et les scripts altbak examinés, est retirée de `Main`, avec son import et `cacheVersion`. Le commentaire d'`onSkipModule` précise la régénération des modules et de leur FFI à chaque lancement ; les API partagées du PBO restent intactes. Build sans avertissement. Sur `FFIIntegerReturns`, les 60 entrées CoreFn et 87 fichiers Go sont identiques avant/après, les 27 assertions passent, et aucun cache n'est créé après modification (60 auparavant). L'écart de snapshot préexistant reste identique. Une modification temporaire de la seule copie FFI du runner (`ReturnInt64` renvoie `value + 1`) est prise en compte par une relance directe de gopurs puis Go sans nettoyage ni Spago : l'assertion échoue avec `Expected: 0 / Actual: 1`. Après restauration exacte et nouvelle génération, le programme affiche `Done` et les 87 fichiers Go sont identiques à la référence ; les contenus et dates des 60 CoreFn n'ont pas changé. Snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-remove-cache-5xwslkjs/`.

## 4. Donner des noms aux contextes et borner l'état du codegen

Les structures partagées ont des alias, `translate` reçoit les métadonnées et le module, et les options du traducteur sont nommées. Chaque invocation de `translate` possède maintenant son propre ensemble de conversions Rebox dans `CodegenState`.

- [x] **4.1 — Nommer les structures existantes.** Les métadonnées, l'environnement local, l'état de génération, les informations des fonctions, les contextes de boucle et le résultat d'expression ont des alias ; leur contenu reste identique (4.1.1–4.1.6).
- [x] **4.1.1 — Nommer l'état de génération.** Le 9 septembre 2026, l'alias `CodegenState` centralise les 12 champs répétés dans `translateExprImpl_` et `translateExprImpl__` ; leurs signatures et la création de `helpersRef` utilisent `Ref CodegenState`. Build réussi avec les 89 avertissements déjà observés dans `CodeGen`. Sur `RBTree`, les 58 entrées CoreFn et les 82 fichiers Go sont identiques avant/après ; le Go compile et renvoie une profondeur de `22` pour 100 000 insertions. L'écart de snapshot préexistant reste identique. Comparaison via une copie temporaire de la fixture avec les dépendances `prelude effect console`. Preuves : `/private/tmp/gopurs-codegen-state-9au_tywe/`.
- [x] **4.1.2 — Nommer l'environnement local.** Le 9 septembre 2026, `LocalBinding` décrit le nom Go et son type, et `LocalEnv` associe chaque `localId` d'origine à cette liaison. Les trois signatures de `constructorReuse`, `translateExprImpl_` et `translateExprImpl__` utilisent cet alias ; un commentaire distingue la clé d'origine du nom Go éventuellement renommé. Build réussi, avec les mêmes 89 avertissements de `CodeGen`. Sur `ShadowedRename`, les 59 entrées CoreFn et les 84 fichiers Go sont identiques avant/après ; le Go compile, l'assertion passe et l'exécution affiche `Done`. L'écart de snapshot préexistant reste identique. Comparaison via une copie temporaire avec les dépendances `prelude effect console assert`. Preuves : `/private/tmp/gopurs-local-env-od0tbfpj/`.
- [x] **4.1.3 — Nommer le résultat d'expression.** `ExprResult` centralise les quatre champs renvoyés par `translateExprImpl_` et `translateExprImpl__`.
- [x] **4.1.4 — Nommer les informations des fonctions.** `FunctionInfo` et `ModuleFunctions` remplacent les records répétés dans `unwrapFunc`, `moduleArities` et les signatures du traducteur.
- [x] **4.1.5 — Nommer les contextes de boucle.** `LoopTarget` et `LoopContext` décrivent les cibles TCO ; les signatures du traducteur et la création du contexte dans `translate` utilisent ces alias.
- [x] **4.1.6 — Centraliser les métadonnées du codegen.** `CodegenMetadataRow` définit les neuf tables communes ; `CodegenMetadata`, `CodegenState` et `PreparedData` partagent cette rangée fermée. Les champs et la création de l'état restent identiques.
- [x] **4.2 — Regrouper les métadonnées de translate.** `Main` projette explicitement les neuf tables dans `CodegenMetadata`. `translate` reçoit ce record et le module : deux arguments au lieu de onze. L'argument d'importations ignoré, son calcul et les imports devenus inutiles sont retirés. L'ordre de préparation et les branches de traduction sont conservés.

Validation commune du lot 4.1.3–4.2, le 9 septembre 2026 : build réussi avec les mêmes 89 avertissements de `CodeGen`, aucun dans `Main`. Avant/après, `RBTree`, `TCOMutRec` et `TypeClassMemberOrderChange` produisent respectivement les mêmes 82, 84 et 82 fichiers Go sur 58, 59 et 58 entrées CoreFn identiques. Le Go compile et les trois exécutions réussissent : profondeur `22`, huit assertions TCO puis `Done`, et `true` puis `Done`. Les trois écarts de snapshot préexistants restent identiques. Les fixtures ADT et classes utilisent des copies temporaires avec les dépendances `prelude effect console` ; `TCOMutRec` conserve ses dépendances déclarées. Snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-context-batch-esbqxzmy/`.

- [x] **4.3 — Expliciter les options du traducteur.** Les deux variantes existantes ont des noms explicites et reçoivent `ExprOptions` ; les références à l'état et aux tables de fonctions décrivent leur contenu. Les branches de traduction restent en place.
- [x] **4.3.1 — Nommer les options d'expression.** `ExprOptions` contient `isTail` et `inEffectBlock`. Les 55 appels conservent exactement leurs deux valeurs ; neuf transmettent directement les options reçues.
- [x] **4.3.2 — Nommer l'entrée sans type attendu.** `translateExprImpl_` devient `translateExpr` et continue de transmettre `Nothing` à la variante avec type attendu.
- [x] **4.3.3 — Nommer l'entrée avec type attendu.** `translateExprImpl__` devient `translateExprWithExpectedType` ; son paramètre `Maybe ExprType` et sa propagation restent identiques.
- [x] **4.3.4 — Nommer la référence à l'état.** `helpersRef` devient `codegenStateRef` ; sa création, ses lectures et ses écritures sont conservées. La rangée ouverte d'`isClosureNode` reste inchangée.
- [x] **4.3.5 — Nommer les tables de fonctions.** `moduleArities` devient `moduleFunctions` ; ses variantes locales et le champ de l'accumulateur récursif suivent ce nommage, sans changement de contenu.

Validation commune du lot 4.3, le 9 septembre 2026 : build réussi avec les mêmes 89 avertissements de `CodeGen`. Les 84 fichiers Go et 59 entrées CoreFn de `TCOMutRec`, ainsi que les 88 fichiers Go et 61 entrées CoreFn de `CurriedLambdas`, sont identiques avant/après. Le Go compile dans les deux cas ; les huit assertions TCO passent et affichent `Done`. `CurriedLambdas` échoue déjà avant modification dans `checkBoundaries` avec `Key 'left' not found in record` (code de sortie 2) ; après modification, le même panic et la même pile sont reproduits, adresses mémoire normalisées. Les deux écarts de snapshot préexistants restent identiques. Les copies temporaires des fixtures conservent leurs dépendances déclarées ; snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-expr-options-kj810gza/`.

- [x] **4.4 — Caractériser et regrouper les usages de l'état global.** Les conversions Rebox ont des fonctions d'enregistrement, de recherche des champs et de génération. La séquence A → B → A confirme une sortie stable pour A et la conservation des conversions de B dans la référence globale.
- [x] **4.4.1 — Observer A → B → A.** Un harness temporaire capture les 58 vrais modules transmis à `translate` par `GenericsRep`. Avec A = `Data.Ord` et B = `Data.Bounded`, après remise à zéro explicite de la seule Ref avant chaque scénario de mesure, A seul, A initial et A après B produisent les mêmes 63 257 octets Go. Les paires retenues passent de 13 à 24 puis restent à 24 ; aucun reset implicite entre les modules. Les rapports et fichiers Go des scénarios sont identiques avant/après extraction.
- [x] **4.4.2 — Extraire l'enregistrement des conversions.** `registerReboxPair` centralise les deux sites existants ; clé de module, dédoublonnage et écriture dans la référence globale restent identiques.
- [x] **4.4.3 — Extraire la recherche des champs.** `findReboxFields` conserve la priorité constructeur puis classe, les conventions de nommage, l'ordre des champs et le diagnostic de métadonnées absentes. `ReboxFields` nomme son résultat.
- [x] **4.4.4 — Extraire le rendu d'une conversion.** `renderReboxFunction` mutualise le rendu identique des constructeurs et des classes ; noms, hashes et conversions des champs sont conservés.
- [x] **4.4.5 — Extraire la génération des fonctions Rebox.** `generateReboxFunctions` remplace le bloc de 79 lignes de `translate` par un appel. La boucle relit les paires après le rendu, qui peut demander de nouvelles conversions ; dédoublonnage, terminaison et ordre des fonctions restent identiques.

Validation commune du lot 4.4, le 9 septembre 2026 : compilation réussie avec 86 avertissements connus, contre 89 auparavant (deux noms inutilisés et un masquage disparaissent avec l'extraction, aucun nouvel avertissement). `GenericsRep` et `TypeClassMemberOrderChange` produisent chacun les mêmes 82 fichiers Go sur 58 entrées CoreFn identiques avant/après ; le Go compile et les deux exécutions réussissent. Les écarts de snapshots préexistants restent identiques. Copies temporaires avec les dépendances `prelude effect console` ; snapshots et configuration du runner inchangés. La caractérisation A → B → A est rejouée sur les deux bundles, dans un espace temporaire isolé du runner. Preuves : `/private/tmp/gopurs-rebox-extraction-lpnkrwyc/`.

- [x] **4.5 — Localiser l'état nécessaire.** Les conversions Rebox utilisent la référence locale de génération ; `globalReboxPairs` est supprimée. La séquence A → B → A confirme des états distincts, initialement vides, avec un Go inchangé.
- [x] **4.5.1 — Créer un ensemble local de conversions.** `reboxPairs` appartient à `CodegenState` et est initialisé à `Set.empty` à chaque invocation de `translate` ; les métadonnées partagées ne contiennent pas cet état.
- [x] **4.5.2 — Enregistrer dans l'état reçu.** `registerReboxPair` reçoit la référence locale, dédoublonne les paires et modifie uniquement son champ `reboxPairs`.
- [x] **4.5.3 — Expliciter l'état des coercions.** `coerceGoExpr` reçoit et transmet la référence locale, y compris lors des conversions récursives.
- [x] **4.5.4 — Expliciter l'état du boxing.** `boxGoExpr` et `boxGoExprImpl` transmettent la même référence ; les branches qui émettent directement ignorent ce paramètre.
- [x] **4.5.5 — Expliciter l'état de l'unboxing.** `unboxGoExpr` conserve la référence dans ses appels récursifs et ses conversions de champs et d'éléments.
- [x] **4.5.6 — Générer depuis l'état local.** `renderReboxFunction` transmet la référence aux coercions des champs. `generateReboxFunctions` la relit à chaque tour pour découvrir les nouvelles paires ; la référence globale et sa table par nom de module sont retirées.

## 5. Séparer les types Go et les conversions de valeurs

- [x] **5.1 — Extraire la traduction des types.** `Gopurs.GoTypes` expose `exprTypeToGoType` et `isClosedRowTail`. Leurs corps restent identiques ; les relais dans `CodeGen` conservent l'API existante, y compris pour les usages FFI du helper de rangées.
- [x] **5.2 — Déplacer les variantes génériques.** Les variantes partagent le même module pur, sans dépendance à `CodeGen`, avec des relais conservés.
- [x] **5.2.1 — Déplacer la traduction générique.** `exprTypeToGenericGoType` rejoint `GoTypes` sans changement des règles pour les paramètres, applications de types, ADT et rangées.
- [x] **5.2.2 — Déplacer les types des champs.** `structFieldGoType` rejoint `GoTypes` avec le même traitement des interfaces.
- [x] **5.2.3 — Déplacer l'instanciation des types Go.** `instantiateGenericGoType` rejoint `GoTypes` avec les mêmes substitutions et reconstructions.
- [x] **5.3 — Rendre les effets des coercions visibles.** Le passage explicite de la référence aux quatre fonctions de conversion et à l'enregistrement est réalisé aux points 4.5.2–4.5.6 ; les conversions restent dans `CodeGen` avant leur extraction.

Validation commune du lot de dix étapes, le 9 septembre 2026 : compilation réussie avec les mêmes 86 avertissements connus, dont ceux des fonctions déplacées. Les corps des cinq fonctions de types sont conservés littéralement. `GenericsRep`, `TypeClassMemberOrderChange`, `NativeRecordBoxing` et `VisibleTypeApplications` produisent les mêmes 82, 82, 86 et 84 fichiers Go sur 58, 58, 60 et 59 entrées CoreFn identiques avant/après ; le Go compile et les quatre exécutions réussissent, dont les dix assertions de boxing. Les écarts de snapshots préexistants restent identiques. Le harness rejoue A seul puis A → B → A sans reset après modification : quatre références distinctes, toutes initialement vides, avec 13 puis 13 → 11 → 13 paires. Les quatre sorties rejouées, les 82 fichiers Go de la génération capturée et les conversions émises sont identiques avant/après. Registre global absent. Copies temporaires des fixtures ; snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-local-rebox-types-tx4bqq91/`.

- [x] **5.4 — Extraire les conversions.** `Gopurs.GoConversions` regroupe le catalogue des ADT natifs, le boxing, l'unboxing, les coercions et la génération Rebox. Le groupe de fonctions qui s'appellent mutuellement est déplacé ensemble ; les points d'entrée de `CodeGen` conservent leurs signatures et délèguent au nouveau module.
- [x] **5.4.1 — Partager les types d'état.** `CodegenMetadataRow`, `CodegenMetadata` et `CodegenState` sont définis dans `Gopurs.CodegenState`, avec des alias conservés dans `CodeGen`. Champs et initialisation restent identiques ; les nouveaux modules ne dépendent pas de `CodeGen`.
- [x] **5.4.2 — Déplacer le catalogue des représentations natives.** `UnboxedADT`, `unboxableADTs` et `getUnboxedADT` rejoignent `GoConversions` ; les règles de `Maybe`, `Tuple` et `Either` sont conservées.
- [x] **5.4.3 — Déplacer le boxing.** `boxGoExpr` et `boxGoExprImpl` gardent leurs rôles respectifs, leur récursion et les conversions des records et tableaux.
- [x] **5.4.4 — Déplacer l'unboxing.** `unboxGoExpr` conserve les branches et le traitement de chaque représentation cible.
- [x] **5.4.5 — Déplacer les coercions.** `coerceGoExpr` conserve l'ordre des clauses, les appels et l'enregistrement des conversions. Les traces et la clause redondante préexistantes restent inchangées.
- [x] **5.4.6 — Déplacer l'enregistrement Rebox.** `registerReboxPair` continue de modifier uniquement l'ensemble de la référence locale reçue.
- [x] **5.4.7 — Déplacer la recherche des champs.** `ReboxFields` et `findReboxFields` conservent la priorité constructeur puis classe, les noms reconnus et le diagnostic de métadonnées absentes.
- [x] **5.4.8 — Déplacer le rendu Rebox.** `renderReboxFunction` appelle directement les deux helpers de `GoTypes`, avec les mêmes substitutions, noms hachés et chaînes Go.
- [x] **5.4.9 — Déplacer la génération transitive.** `generateReboxFunctions` relit la même référence après chaque rendu et termine lorsqu'aucune fonction supplémentaire n'est découverte.
- [x] **5.5 — Documenter les choix de représentation.** Des commentaires près des fonctions expliquent les records fermés et les autres queues de rangées, les paramètres génériques, les tableaux, la normalisation des pointeurs avant boxing et le rôle distinct des signatures Go à la frontière FFI.

Validation commune du lot de dix étapes, le 9 septembre 2026 : compilation réussie ; les 80 avertissements émis par `CodeGen` et `GoConversions` correspondent exactement aux diagnostics déjà connus des fonctions déplacées ou conservées. `GoTypes`, dont seuls les commentaires changent, était en cache lors du build. Les corps déplacés sont identiques après les seules qualifications de trois appels vers deux helpers de `GoTypes` ; le reste de `CodeGen` est conservé hors imports et relais. `GenericsRep`, `TypeClassMemberOrderChange`, `NativeRecordBoxing` et `ArrayRoundtrip` produisent les mêmes 82, 82, 86 et 195 fichiers Go sur 58, 58, 60 et 153 entrées CoreFn identiques avant/après. Les quatre exécutables Go réussissent avec les mêmes résultats, dont dix assertions de records avec callback opaque et vingt-huit cas de tableaux. Les quatre écarts de snapshots préexistants restent identiques. La caractérisation A seul puis A → B → A utilise les mêmes 58 CoreFn et aucune remise à zéro : références distinctes, ensembles initialement vides, 13 puis 13 → 11 → 13 paires dans les deux versions. Les quatre sorties rejouées et les 82 fichiers Go de la génération capturée sont identiques. Fixtures copiées temporairement ; snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-conversions-ur3eh9hh/`.

## 6. Donner au pont FFI sa propre frontière

La génération des wrappers est isolée dans `Gopurs.FfiBridge`. Le transport de `FfiSupport.js`, qui construit un programme Node et une commande shell imbriqués, relève des étapes suivantes.

- [x] **6.1 — Extraire les helpers FFI purs.** Le nouveau module rassemble la lecture des signatures Go et TAST et la résolution des newtypes. Les helpers partagés rejoignent `GoAst` et `GoTypes` ; les signatures et relais dans `CodeGen` sont conservés.
- [x] **6.1.1 — Partager le nommage Go.** `capitalize` rejoint `GoAst` près de `sanitizeName`, avec le même traitement des underscores. La fonction locale distincte de `Printer` reste inchangée.
- [x] **6.1.2 — Partager le rendu diagnostique TAST.** `printExprType` rejoint `GoTypes`, avec un export explicite et le même texte pour les diagnostics du traducteur et les commentaires FFI.
- [x] **6.1.3 — Déplacer la lecture des signatures Go.** `printTypeNode` et `isStandardPursFunc` rejoignent `FfiBridge`, avec les mêmes règles pour fonctions, tableaux, maps et types opaques.
- [x] **6.1.4 — Déplacer les helpers TAST du pont.** `getTastArgType`, `getTastReturnType`, `exprTypeToDummyTypeNode` et `flattenFuncArgs` conservent leurs résultats et leurs appels récursifs.
- [x] **6.1.5 — Déplacer la résolution des newtypes.** `resolveNewtype` conserve sa recherche dans `dataDecls`, ses critères et son parcours des annotations enrichies.
- [x] **6.2 — Extraire les wrappers.** `FfiBridge` contient les conversions FFI et leurs générateurs ; il dépend directement des helpers de `GoAst` et `GoTypes`, sans dépendance vers `CodeGen`.
- [x] **6.2.1 — Déplacer le boxing FFI simple.** `boxFfiValue` conserve le traitement de `int64`, `int` et des autres types.
- [x] **6.2.2 — Déplacer ensemble callbacks et retours.** `unwrapValueToFunc` et `wrapReturn`, mutuellement récursifs, restent dans le même module avec les mêmes branches, conversions et chaînes Go.
- [x] **6.2.3 — Déplacer le générateur de wrapper.** `generateWrapperFunc` conserve l'adaptation des arguments, des retours, des fonctions génériques et des variables étrangères.
- [x] **6.2.4 — Déplacer l'assemblage du pont.** `generateFfiBridge` conserve les noms exportés, l'ordre des correspondances et le diagnostic des déclarations absentes.
- [x] **6.2.5 — Brancher Main directement.** Les deux sites d'émission utilisent `FfiBridge.generateFfiBridge` ; les points d'entrée publics de `CodeGen` restent des relais. Les exports du nouveau module sont explicites. `hasTypeVars` reste auprès du traducteur, qui en est le seul consommateur.

Validation commune du lot de dix étapes, le 9 septembre 2026 : compilation réussie. Les 14 blocs déplacés, signatures et corps compris, sont identiques octet pour octet ; une transformation inverse restitue exactement les quatre fichiers sources initiaux. `CodeGen` passe de 3 151 à 2 762 lignes. Les 94 avertissements émis comprennent les mêmes 86 diagnostics déjà connus et huit masquages du nom `expr` dans `Printer`, recompilé avec un contenu inchangé. `FFIIntegerReturns`, `FFIConstraintWorkaround`, `NativeRecordSizes` et `ESFFIFunctionFunction` produisent les mêmes 87, 87, 196 et 194 fichiers Go sur 60, 60, 153 et 152 entrées CoreFn identiques avant/après, wrappers compris. Les trois premières fixtures conservent leur écart préexistant de snapshot principal ; le snapshot principal d'`ESFFIFunctionFunction` et les deux snapshots FFI déclarés correspondent aux références. Les exécutions d'`FFIIntegerReturns`, `FFIConstraintWorkaround` (avec appel réel du callback) et `ESFFIFunctionFunction` réussissent avec les mêmes sorties ; ce dernier exécute seulement son `main` affichant `Done`.

Limite préexistante observée avant toute modification : `NativeRecordSizes` réussit 29 assertions puis plante pendant la vérification de la trace d'effets (`tests/passing/NativeRecordSizes.purs:147`), avec `fatal error: fault`, un code de sortie 2 et une pile passant par `Test_Assert.go:425` puis `Value.StrVal` à `runtime.go:91`. Deux répétitions indépendantes du binaire initial confirment le comportement. Après extraction, les 29 lignes de sortie, le signal et les onze frames de la pile principale sont identiques, adresses et offsets normalisés. Ce cas n'est pas compté comme réussi. Fixtures copiées temporairement avec leurs compagnons FFI ; snapshots et configuration du runner inchangés. Preuves : `/private/tmp/gopurs-ffi-bridge-r0wbn890/`.

- [x] **6.3 — Séparer le transport Node.** `FfiSupport.js` appelle un runner autonome ; entrée standard, sortie JSON et comportement des erreurs sont conservés.
- [x] **6.3.1 — Sortir le runner WASM.** `tools/ffi-runner.mjs` charge le runtime voisin, instancie le WASM, lit le contenu Go sur stdin et écrit la réponse. Le traitement des rejets asynchrones reste identique.
- [x] **6.3.2 — Passer les arguments sans shell.** `execFileSync(process.execPath, [runner])` remplace le programme et la commande imbriqués, avec les mêmes options UTF-8 et limite de 10 Mio.
- [x] **6.3.3 — Résoudre les chemins des trois emplacements.** Le runner est trouvé depuis les sources, la sortie Spago et le bundle installé, indépendamment du répertoire courant ; espaces, accents, apostrophes, guillemets et caractères de shell sont couverts.
- [x] **6.4 — Reproduire la fabrication du WASM.** La commande explicite `npm run build:ffi` compile le parseur et récupère le chargeur du même Go. Les étapes ordinaires `build` et `prepare` continuent d'utiliser les artefacts versionnés.
- [x] **6.4.1 — Reconstruire la paire d'artefacts.** `tools/build-ffi.mjs` exige la version exacte Go 1.27.0 de `go.mod`, désactive la sélection automatique du toolchain et utilise `-trimpath`, `-buildvcs=false` et un build ID vide. Compilation et copie du runtime sont préparées avant remplacement ; un échec de compilation conserve les artefacts existants. Le chargeur fourni par Go 1.27.0 est identique au fichier déjà versionné.
- [x] **6.4.2 — Documenter et vérifier la distribution.** Le README décrit la reconstruction, les tests natifs et les fichiers livrés. Le paquet npm produit contient le bundle, le runner, le WASM, son runtime, le builder et les sources/tests du parseur ; son bundle retrouve les outils depuis un autre répertoire.
- [x] **6.5 — Tester le parseur Go isolément.** L'analyse des sources Go est séparée de l'interface JavaScript et se teste nativement.
- [x] **6.5.1 — Isoler le module Go.** `tools/ffi-gen/go.mod` définit un module sans dépendance externe, avec la version du générateur épinglée.
- [x] **6.5.2 — Regrouper le contrat JSON.** `types.go` conserve exactement les structures `TypeNode`, `FFIDecl` et leurs tags.
- [x] **6.5.3 — Extraire l'analyse.** `parser.go` expose localement `parseFFI(content string) string` et contient les helpers AST ; le FileSet partagé, les règles et les retours `[]` sont conservés.
- [x] **6.5.4 — Réduire l'adaptateur WASM.** `main_js_wasm.go` contient seulement l'adaptation des arguments JavaScript, l'enregistrement de la fonction et l'attente. Le suffixe de fichier exclut `syscall/js` des tests natifs.
- [x] **6.5.5 — Vérifier les contrats.** Onze cas natifs couvrent fonctions, variables, génériques, callbacks imbriqués, types inconnus, filtrage, contenu vide/invalide et coupure avant wrappers, avec les mêmes distinctions entre champs omis, `null` et `[]`.

Validation commune du lot de dix étapes, le 9 septembre 2026 : tests Go et build du bundle réussis. Les quatorze réponses du corpus, JSON brut compris, sont identiques avec l'ancien transport, le nouveau runner sur l'ancien WASM, puis le WASM reconstruit en Go 1.27.0. Sources, sortie Spago, bundle et bundle extrait du paquet npm donnent les mêmes réponses ; les trois emplacements sont aussi vérifiés dans un chemin contenant espaces et caractères spéciaux. Deux reconstructions depuis des répertoires courants différents produisent les mêmes SHA-256 ; le runtime correspond au GOROOT de compilation. Les scénarios de mauvaise version Go et de source invalide échouent en conservant les artefacts et en nettoyant le dossier temporaire. Les tests de WASM absent/corrompu reproduisent les catégories d'erreur antérieures, notamment le rejet asynchrone journalisé avec retour vide et code de sortie zéro, traité ultérieurement en 6.6.

`FFIIntegerReturns` et `FFIConstraintWorkaround` conservent chacun leurs 87 fichiers Go et 60 entrées CoreFn, wrappers compris. Les deux exécutables réussissent avec les mêmes sorties et le callback FFI est appelé. Les écarts préexistants de snapshots principaux restent identiques ; snapshots et configuration du runner sont conservés. Preuves : `/private/tmp/gopurs-ffi-transport-lnutc6ta/`.

- [x] **6.6 — Clarifier les erreurs FFI.** Les échecs du parseur, du runner et du décodage sont explicites ; `Main` reçoit les déclarations typées ou une exception contextualisée.
- [x] **6.6.1 — Reproduire les erreurs initiales.** Le parseur acceptait le Go invalide comme `[]` avec statut zéro ; un WASM corrompu produisait une sortie vide avec statut zéro. Le décodage de `Main` remplaçait aussi une réponse JSON invalide par `[]` après journalisation.
- [x] **6.6.2 — Distinguer erreur et résultat vide dans Go.** `parseFFI` retourne `(string, error)`, rejette les déclarations partielles et vérifie la sérialisation. Un contenu valide vide ou sans déclaration retenue conserve `[]`. Le préfixe de package préserve les positions de diagnostic et les directives `//line` de la source.
- [x] **6.6.3 — Adapter la frontière WASM.** `main_js_wasm.go` expose un résultat interne `{json, error}` ; le JSON des déclarations valides reste inchangé.
- [x] **6.6.4 — Terminer le runner avec un statut fiable.** Le runner vérifie le résultat du WASM, écrit les diagnostics sur stderr et termine avec le statut 1 après vidange. Il observe les rejets de `go.run` sans attendre le programme Go, qui reste en attente des appels.
- [x] **6.6.5 — Propager le diagnostic du processus enfant.** Le transport Node capture stderr et lève une erreur sans journalisation supplémentaire. `FfiSupport` ajoute une seule fois le module PureScript et le chemin du fichier FFI.
- [x] **6.6.6 — Isoler le décodage pur.** `decodeFfiDecls` distingue JSON illisible et déclarations invalides dans un `Either`, avec le décodeur existant de `FfiDecl`.
- [x] **6.6.7 — Exposer les déclarations typées.** `extractFfiDecls` reçoit le contexte du fichier et retourne `Effect (Array FfiDecl)` ; le JSON brut reste derrière l'import étranger.
- [x] **6.6.8 — Simplifier Main.** Son unique appel FFI consomme cette API. Le décodage Argonaut, le repli `[]`, la journalisation et l'import `unsafePerformEffect` correspondants sont retirés.
- [x] **6.6.9 — Vérifier les diagnostics.** Treize cas Go et onze tests Node passent. `npm run test:ffi` couvre aussi les réponses JSON invalides via le vrai module compilé isolé, les outils absents/corrompus et l'absence de double journalisation. Cinq scénarios sur le bundle complet terminent avec statut 1 : Go invalide, JSON illisible, schéma invalide, WASM absent et WASM corrompu.
- [x] **6.6.10 — Comparer les sorties valides.** Douze réponses valides du corpus restent identiques octet pour octet ; les deux entrées invalides auparavant converties en `[]` échouent désormais. Les deux fixtures FFI compilent en Go et s'exécutent avec les mêmes sorties, sans changement du Go généré.

Validation du lot, le 9 septembre 2026 : reconstruction du WASM et du bundle réussie, compilation PureScript sans erreur ni avertissement. Chaque fixture utilise une sortie générée propre : `FFIIntegerReturns` conserve ses 87 fichiers Go et 60 entrées CoreFn ; `FFIConstraintWorkaround`, avec ses seules dépendances utiles dans une copie temporaire, conserve 85 fichiers Go et 59 entrées CoreFn. Les deux exécutables terminent avec statut zéro. Les différences de snapshots principaux étaient déjà présentes et restent identiques ; les snapshots et la configuration du runner sont conservés. L'essai initial avec toutes les dépendances avait échoué sur `EscapedSkolem` dans `gopurs-foreign-object`, avant génération Go ; ce problème externe n'est pas corrigé par ce lot.

Le comportement change pour les erreurs d'une FFI trouvée. Une FFI absente conserve le chemin de secours existant. Le compilateur peut avoir écrit certains fichiers avant d'échouer : ce lot ne rend pas la génération transactionnelle. Revue indépendante sans constat actionnable. Preuves : `/private/tmp/gopurs-ffi-errors-iaj25xd5/`.

## 7. Découper le traducteur d'expressions par responsabilité

Constat : `translateExprImpl__` couvre environ 1 960 lignes. Le contexte du point 4 doit permettre des extractions sans créer de dépendances circulaires.

- [ ] **7.1 — Cartographier les branches.** Associer chaque famille à ses entrées, effets sur l'état et fixtures ; repérer les helpers réellement partagés.
- [ ] **7.2 — Extraire une première branche simple.** Choisir un littéral ou une primitive, comparer son résultat et conserver le dispatch central lisible.
- [ ] **7.3 — Extraire les records.** Déplacer successivement construction, accès et mise à jour ; vérifier ordre des champs, évaluation unique et conservation des anciennes valeurs.
- [ ] **7.4 — Extraire les constructeurs ADT.** Déplacer une opération à la fois, en préservant les décisions fondées sur `dataDecls` et les représentations natives.
- [ ] **7.5 — Extraire les appels.** Commencer par les helpers de spine et d'arité, puis les appels directs ; traiter ensuite closures et appels indirects avec `CurriedLambdas` et les fixtures FFI.
- [ ] **7.6 — Extraire le contrôle et la récursion.** Traiter séparément branches, boucles et TCO ; vérifier `TCO`, `TCOMutRec` et `ThunkFusion` avant de valider cette famille.
- [ ] **7.7 — Réduire les API.** Définir des exports explicites pour les modules extraits et retirer les relais devenus inutiles ; conserver les helpers propres à une famille dans son module.

## 8. Clarifier la frontière entre AST Go et impression

Constat : `Printer.purs` reste court (235 lignes), mais du Go est aussi assemblé via `GoRaw`, `rawDecls` et `printGoExpr` dans le codegen.

- [ ] **8.1 — Inventorier les fragments bruts.** Séparer déclarations, expressions, conversions et adaptations FFI ; choisir une famille répétée dont la structuration simplifie réellement le code.
- [ ] **8.2 — Structurer un seul motif.** Ajouter le nœud GoAst nécessaire et son rendu, puis migrer un site ; comparer le Go après `gofmt` et exécuter sa fixture.
- [ ] **8.3 — Migrer les occurrences équivalentes.** Réutiliser ce nœud pour les autres sites de la même famille, sans étendre la règle de génération.
- [ ] **8.4 — Fixer le contrat du printer.** Documenter ce qui relève du rendu et ce qui doit être décidé avant celui-ci, notamment les conversions de tableaux déjà structurées dans GoAst.

## 9. Éditer le runtime comme du Go

Constat : `Runtime.purs` contient essentiellement une grande chaîne de code Go.

- [ ] **9.1 — Définir la source canonique.** Prévoir un fichier Go éditable et une inclusion compatible avec le build npm, le bundle et les chemins d'installation ; vérifier ce mécanisme sur un exemple isolé.
- [ ] **9.2 — Déplacer le texte sans réécriture.** Extraire le runtime et vérifier l'identité du `runtime.go` émis avant toute modification de présentation ou de structure.
- [ ] **9.3 — Valider l'assemblage.** Reconstruire le bundle, générer une fixture, compiler son Go et vérifier qu'une modification du fichier canonique est bien prise en compte.
- [ ] **9.4 — Séparer une famille si utile.** Commencer par les helpers de records ; vérifier `NativeRecordSizes` et `CompactRecordConsumers`. Traiter appels/closures et event loop dans des étapes ultérieures distinctes avec leurs propres fixtures.

## 10. Rendre le runner de tests prévisible

Constat : `bin/test` mélange sélection, préparation Spago, nettoyage, compilation, snapshots et exécution dans un workspace partagé. `bin/modtest` démarre par défaut à `strings`.

- [ ] **10.1 — Extraire la sélection.** En faire une fonction et ajouter un mode affichant les fixtures choisies sans compiler ; vérifier cas unique, liste, reprise et nom inconnu.
- [ ] **10.2 — Séparer vérification et mise à jour des snapshots.** Extraire leur traitement, puis rendre explicite la création d'un snapshot absent ; vérifier concordance, différence et absence dans un petit répertoire temporaire.
- [ ] **10.3 — Extraire compilation et exécution.** Donner à chaque phase une fonction, un diagnostic et un statut de sortie ; contrôler un succès et un échec provoqué.
- [ ] **10.4 — Caractériser le workspace partagé.** Exécuter deux fixtures aux modules distincts successivement et relever les fichiers conservés avant de modifier le nettoyage.
- [ ] **10.5 — Isoler les fichiers d'une exécution.** Remplacer progressivement le nettoyage par noms de modules par un espace appartenant au run ; préserver les dépendances nécessaires et vérifier restauration de configuration après échec et interruption.
- [ ] **10.6 — Expliciter les campagnes de modules.** Vérifier la portée actuelle de `modtest`, puis distinguer campagne complète et reprise demandée ; rendre la sélection visible avant exécution.

## 11. Documenter le parcours contributeur et valider les jalons

- [ ] **11.1 — Documenter l'architecture effectivement obtenue.** Ajouter le trajet TAST → préparation → optimisation/monomorphisation → codegen → GoAst → impression, avec le rôle du runtime et de la FFI ; vérifier l'ordre réel des passes dans le code.
- [ ] **11.2 — Corriger les instructions locales.** Décrire les checkouts requis, dont `../../purescript-backend-optimizer-gopurs`, le `purs` typé, le build du bundle et celui du WASM. Vérifier les chemins d'exécution Go des exemples.
- [ ] **11.3 — Documenter les vérifications courtes.** Donner les commandes par famille, les modes de snapshots, la politique de cache et les exclusions de tests ; contextualiser les affirmations de couverture du README.
- [ ] **11.4 — Vérifier chaque jalon transversal.** Après un ensemble cohérent d'extractions, lancer la suite `passing`, comparer les échecs à l'état initial et examiner les diffs générés ; consigner exactement ce qui a été exécuté.
- [ ] **11.5 — Vérifier le parcours final.** Rejouer installation locale, build et fixture ciblée depuis un environnement propre ; corriger la documentation au vu du résultat et consigner les limites restantes.
