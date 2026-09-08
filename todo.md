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

Constat : Git suit des `.bak`, `.orig`, `.rej`, des modules de scratch et des dumps dans `src/` et `tests/runner/`.

- [ ] **2.1 — Classer les fichiers suspects.** Rechercher leurs imports, appels et usages dans les scripts ; examiner notamment `Scratch`, `TestJson`, `TestTrace`, `ShowTco`, `DumpHeap` et les fichiers expérimentaux du runner.
- [ ] **2.2 — Sauver les expériences utiles.** Transformer une expérience encore pertinente en fixture ou outil nommé avec sa commande d'exécution, une expérience à la fois.
- [ ] **2.3 — Retirer les reliquats prouvés inutiles.** Supprimer les sauvegardes, rejets et sorties sans usage ni contenu unique utile, par petits lots ; vérifier le diff et les références restantes.
- [ ] **2.4 — Clarifier les chemins alternatifs.** Vérifier les usages de `CodeGenBackend`, `OptimizeTAST`, `QuoteTAST`, `UsageAnalysis` et `FBIP` ; documenter leur rôle ou retirer ceux dont l'inutilité est établie.
- [ ] **2.5 — Prévenir le retour des artefacts.** Ajouter des exclusions précises pour les sorties identifiées et nettoyer les imports des seuls modules touchés.

## 3. Ramener Main à l'orchestration du build

Constat : `runBuild` et `main` dupliquent la préparation des sorties, les callbacks de `buildModules` et le traitement FFI. Ils diffèrent notamment sur la consommation des modules et le fichier de debug.

- [ ] **3.1 — Déterminer le chemin de référence.** Rechercher les appelants de `runBuild` et caractériser les différences avec `main`, y compris `takeMonomorphizedModules` dans `Main.js`, qui modifie le record reçu.
- [x] **3.2 — Extraire l'émission d'un module.** Le 8 septembre 2026, extraction du callback de `main` dans `emitModule`, avec les métadonnées préparées, le dossier FFI et les deux modules en paramètres. Les alias devenus inutiles dans `main` sont retirés. Compilation sans avertissement ; `NativeRecordBoxing` (10 assertions) et `FFIIntegerReturns` (27) passent avant et après. Tous les fichiers Go générés sont identiques octet par octet (86 et 87 fichiers respectivement), ainsi que les trois snapshots existants. `runBuild` reste inchangé. Preuves : `/private/tmp/gopurs-emit-module-d_1b62n0/`.
- [ ] **3.3 — Unifier le build.** Faire utiliser le même pipeline aux points d'entrée nécessaires, ou retirer `runBuild` si aucun usage n'est établi. Conserver les comportements du chemin actif, notamment les points d'entrée Go émis.
- [ ] **3.4 — Extraire la préparation des métadonnées.** Commencer par `buildGlobalTypes`, puis déplacer séparément les tables de constructeurs, les tables de classes et la préparation de la monomorphisation ; comparer les données et sorties à chaque déplacement.
- [ ] **3.5 — Rendre le cache compréhensible.** Caractériser `onSkipModule` et son `res <- pure Nothing`, ainsi que les écritures de cache ; retirer les branches inaccessibles prouvées ou documenter le fonctionnement réellement conservé. Vérifier deux builds successifs et une modification FFI.

## 4. Donner des noms aux contextes et borner l'état du codegen

Constat : `translate` prend une longue liste d'arguments et `translateExprImpl`, `translateExprImpl_`, `translateExprImpl__` répètent de gros records. Des `Ref` existent au niveau global.

- [ ] **4.1 — Nommer les structures existantes.** Introduire des alias pour les métadonnées, l'environnement local, l'état de génération et le résultat d'expression, sans changer leur contenu.
- [ ] **4.2 — Regrouper les métadonnées de translate.** Remplacer sa liste d'arguments par un contexte nommé ; comparer le Go généré sur une fixture avec ADT et une avec classes.
- [ ] **4.3 — Expliciter les options du traducteur.** Remplacer les booléens positionnels par des champs nommés, puis renommer les trois variantes selon leur responsabilité ; procéder séparément du déplacement des branches.
- [ ] **4.4 — Caractériser la durée de vie des Ref.** Tracer création, remise à zéro et usages de `globalReboxPairs`, `globalRecordStructs`, `globalRecordDecls`. Comparer A → B → A dans un même processus avec A exécuté seul avant de conclure à une fuite d'état.
- [ ] **4.5 — Localiser l'état nécessaire.** Déplacer une référence effectivement utilisée dans l'état de compilation, puis vérifier la même séquence ; retirer séparément les références sans usage démontré.

## 5. Séparer les types Go et les conversions de valeurs

- [ ] **5.1 — Extraire la traduction des types.** Déplacer `exprTypeToGoType` et ses dépendances minimales vers un module dédié, en gardant d'abord l'API publique existante comme relais.
- [ ] **5.2 — Déplacer les variantes génériques.** Y regrouper `exprTypeToGenericGoType`, `structFieldGoType` et `instantiateGenericGoType` ; contrôler applications de types, ADT et classes avec les fixtures existantes.
- [ ] **5.3 — Rendre les effets des coercions visibles.** Identifier les écritures effectuées par `coerceGoExpr`, boxing et unboxing ; leur transmettre explicitement l'état défini au point 4 avant toute extraction.
- [ ] **5.4 — Extraire les conversions.** Déplacer boxing, puis unboxing/coercions dans des étapes séparées ; vérifier records natifs, tableaux et callbacks avec snapshots identiques.
- [ ] **5.5 — Documenter les choix de représentation.** Décrire près du code les conditions du passage en natif et du fallback `Value`, notamment rangées ouvertes, polymorphisme et frontières FFI.

## 6. Donner au pont FFI sa propre frontière

Constat : les helpers et générateurs FFI occupent la fin de `CodeGen.purs` ; `FfiSupport.js` construit un programme Node et une commande shell imbriqués.

- [ ] **6.1 — Extraire les helpers FFI purs.** Commencer par `printTypeNode`, la résolution des newtypes et la lecture des signatures TAST, en conservant exactement les résultats.
- [ ] **6.2 — Extraire les wrappers.** Déplacer les conversions de retour, puis `generateWrapperFunc` et `generateFfiBridge` vers un module FFI dédié ; vérifier `FFIIntegerReturns`, `ESFFIFunctionFunction` et leurs snapshots.
- [ ] **6.3 — Séparer le transport Node.** Sortir le runner WASM dans un fichier et transmettre les arguments sans assembler une commande shell ; comparer le JSON pour fonction, variable, générique et fichier vide.
- [ ] **6.4 — Reproduire la fabrication du WASM.** Ajouter une commande explicite de génération de `ffi_gen.wasm` et de récupération du `wasm_exec.js` correspondant ; documenter la version Go et vérifier le parsing avec les artefacts reconstruits.
- [ ] **6.5 — Tester le parseur Go isolément.** Extraire le parsing de `tools/ffi-gen/main.go` hors de l'adaptateur `syscall/js`, avec quelques cas de contrat ciblés.
- [ ] **6.6 — Clarifier les erreurs FFI.** Reproduire un fichier Go invalide et une réponse JSON invalide, puis traiter séparément les retours silencieux `[]` et le logging via `unsafePerformEffect`. Vérifier un diagnostic contextualisé et un échec observable.

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
