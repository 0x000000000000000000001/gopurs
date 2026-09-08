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

- [ ] **1.1 — Identifier les outils réellement utilisés.** Relever chemins et versions de `purs` typé, Spago, Node et Go, ainsi que les révisions de gopurs et du PBO local déclaré dans `spago.yaml`.
- [ ] **1.2 — Vérifier un cycle court.** Exécuter `./bin/test NativeRecordSizes -c` avec le `purs` TAST sur le `PATH`, puis la même fixture sans `-c` ; relever les fichiers produits et comparer les snapshots sans les actualiser.
- [ ] **1.3 — Définir les contrôles par famille.** Associer types/records à `NativeRecordBoxing` et `NativeRecordSizes`, FFI à `FFIIntegerReturns`, appels à `CurriedLambdas`, tableaux à `ArrayRoundtrip`, récursion à `TCO`/`TCOMutRec`, fusion à `ThunkFusion`. Vérifier leur état initial par petits groupes.
- [ ] **1.4 — Consigner les limites initiales.** Distinguer échecs existants, exclusions et contrôles non exécutés. Conserver les sorties nécessaires aux comparaisons suivantes hors des sources de production.

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
- [ ] **3.2 — Extraire l'émission d'un module.** Déplacer le callback `onCodegenModule` du chemin actif dans une fonction nommée, avec ses dépendances explicites ; comparer une fixture avec FFI et une sans FFI.
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
