# Gopurs — compréhension et maintenance par des développeurs PureScript

## Résultat attendu

Un développeur PureScript doit pouvoir installer le projet, suivre la génération d’un programme, trouver le responsable d’un comportement, comprendre son code interne, le modifier et lancer le contrôle approprié sans reconstituer l’historique des investigations de l’IA.

Ce plan couvre **gopurs et les 50 bibliothèques locales `gopurs-*`** : sources PureScript, FFI Go et JavaScript, runtime, parser, tests, scripts, configurations et documentation. Le précédent todo est remplacé ; son historique reste dans Git.

La priorité est la compréhension et la maintenance. **Le nombre de `.mjs` n’est plus un critère de réussite autonome.** Les scripts temporaires disparaissent ; la logique et les tests qui appartiennent à PureScript y sont intégrés ; une petite interface technique peut rester en JavaScript si sa nécessité et son fonctionnement sont clairs. Un changement d’extension ou de dossier ne suffit pas à améliorer le code.

La mise en œuvre se fait ensemble, lot par lot. Le bilan en fin de document suit les lots réalisés.

## Manière de travailler

1. Avant chaque lot, expliquer son rôle avec un exemple concret : entrée, traitement, sortie, consommateurs. Présenter les choix de langage, d’architecture ou de commande à décider ensemble avant de les figer.
2. Prendre une référence des comportements concernés et distinguer les défauts déjà présents. En cas d’incertitude, faire une expérience courte dans `/tmp` avant d’engager une refonte.
3. Livrer un bloc cohérent et compréhensible. À l’intérieur du lot, avancer sans demander une confirmation pour chaque renommage ou extraction de fonction.
4. Vérifier le comportement, expliquer le résultat et mettre à jour ce todo avant de choisir ensemble le lot suivant. Une étape d’explication ne vaut pas autorisation d’exécuter tout le plan.
5. Préserver les modifications en cours de l’utilisateur. Chaque bibliothèque reste un dépôt avec ses propres responsabilités ; ce plan n’impose ni fusion des dépôts ni nouveau framework.
6. Garder les preuves et expériences temporaires hors des dépôts. Réutiliser les documents existants et consigner ici un bilan court par lot ; éviter une collection de rapports, scripts et journaux permanents.

## Ce que la revue doit vérifier dans chaque fichier

Cette grille s’applique à tous les fichiers maintenus, y compris les petits fichiers, les helpers et les compagnons FFI. Les dépendances installées sont hors périmètre. Les fichiers générés ou distribués depuis un outil tiers sont contrôlés par leur source canonique, leur provenance et leur procédure de reconstruction.

| Aspect | Critère concret |
| --- | --- |
| Rôle et emplacement | On comprend pourquoi le fichier existe, qui l’appelle et à quelle phase il intervient : build, compilation, programme généré ou tests. |
| Organisation interne | Le parcours principal se lit dans l’ordre ; les opérations secondaires ont des noms qui expriment leur rôle. Les commentaires longs ne compensent pas un flux confus. |
| Fonctions | Chaque fonction a une responsabilité identifiable. Extraire une fonction lorsqu’elle nomme une étape utile ou partage une règle réelle ; garder ensemble ce qui se comprend ensemble. Aucun quota de lignes ou de fonctions. |
| Noms et signatures | Noms du domaine, unités et variantes explicites, signatures PureScript claires. Remplacer les suites d’arguments positionnels ambigus et les drapeaux mal nommés par des contrats adaptés lorsque cela simplifie les appels. |
| Types et données | Utiliser ADT, records nommés et résultats typés pour exprimer les variantes utiles. Rendre visibles les invariants ; éviter de reconstruire des informations disponibles à partir de chaînes ou de structures non typées. |
| Effets et état | Distinguer calcul et effets, localiser la mutation, rendre l’ordre d’évaluation visible. Examiner chaque coercition, référence mutable et `unsafePerformEffect` ; garder les usages nécessaires avec leur justification. |
| Erreurs et cas limites | Distinguer absence normale, entrée invalide et panne. Conserver le contexte des erreurs ; examiner les exceptions absorbées, fallbacks silencieux, panics, stubs et diagnostics de debug. |
| Duplication | Partager les règles identiques et la préparation répétée. Garder explicites les variantes sémantiques ; éviter les helpers universels qui cachent leurs différences. |
| Génération de code | Séparer préparation typée et impression. Réutiliser les nœuds Go existants pour les structures répétées ; rendre explicites les fragments opaques nécessaires et leurs dépendances. |
| Lisibilité locale | Mise en forme cohérente, imports utiles, expressions décomposées lorsqu’elles masquent les étapes, imbrication raisonnable. Les commentaires expliquent les contraintes et les invariants plutôt que l’historique des essais. |
| Tests | Le nom annonce un comportement, la préparation reste courte et partageable, l’assertion vérifie une propriété utile. Les tests ne recopient pas l’implémentation pour s’approuver eux-mêmes. |
| API et provenance | Conserver les interfaces publiques et les licences. Distinguer code upstream, adaptation Go et ajout local ; limiter les divergences upstream aux améliorations justifiées. |

La taille d’un fichier sert à choisir où regarder, pas à décider de le découper. Un fichier revu peut rester inchangé si son rôle et son organisation sont satisfaisants.

## Invariants à préserver

- Le backend utilise le TAST enrichi du fork PureScript : `dataDecls`, `classDecls`, annotations de types, `ForAll`, contraintes, `TypeApp`, ordre des labels et queues de rangées. Préserver ces informations et l’ordre des passes, notamment la monomorphisation avant l’optimisation PBO.
- Préserver les représentations natives, tags, boxing, conversions, conventions FFI, arités, currying, effets, TCO, fusion et ordre d’évaluation. Une correction fonctionnelle découverte pendant la revue doit être explicitée et validée comme telle.
- Conserver le runtime Go canonique et un chemin de distribution fonctionnel du compilateur. Ne pas ajouter par inadvertance un besoin de lire le checkout source au lancement du bundle distribué.
- Pour les bibliothèques, préserver signatures publiques, sémantique PureScript, callbacks, annulation, gestion des ressources et compatibilité avec les dépendances utilisées.
- Garder le compilateur PureScript forké, le PBO et altbak comme dépendances et points de validation. Toute modification nécessaire de leurs sources constitue un périmètre distinct à discuter.

## Éléments observés à reprendre

- `tools/` contient 20 `.mjs`. Neuf tests spécialisés construisent des AST via le JavaScript compilé de PureScript ; leur préparation et leurs programmes Go sont largement répétés. Les onze tests Node représentent 43 cas réussis lors du dernier contrôle ; les tests Go du parser passent également. Cela constitue une référence, sans imposer de conserver chaque fichier ou assertion.
- Plusieurs `bin/test` frères dupliquent la chaîne de compilation et nettoient aussi `../gopurs-*/output`, `.spago` et `.cache`. La portée des effets dépasse donc le paquet lancé.
- Des configurations, notamment celle de `gopurs-aff`, référencent `gopurs-js-uri`, `gopurs-simple-json` et `gopurs-test`, absents de l’inventaire local actuel. Vérifier leur utilité et leur résolution avant modification.
- `gopurs-js-bigints` n’a pas de README et `gopurs-quickcheck` n’a pas de `bin/test` dans l’état inspecté. Clarifier leur parcours réel plutôt que leur appliquer un modèle mécaniquement.
- Les frontières actuelles sont décrites dans [docs/architecture.md](docs/architecture.md), [docs/go-ast-printer.md](docs/go-ast-printer.md) et [docs/testing.md](docs/testing.md). Les précédents cleanups ne dispensent pas d’examiner l’intérieur des fonctions.

## Lots

- [x] **1 — Établir la carte et les références des 51 dossiers.**

  Inventorier les fichiers suivis et non suivis pertinents, les points d’entrée, les exports, les compagnons FFI, les configurations, les dépendances et les sorties générées. Identifier le propriétaire et les consommateurs de chaque famille. Repérer les usages dynamiques et les différences upstream/local avant toute suppression.

  Pour chaque dossier, noter ici son rôle, ses commandes actuelles et le lot responsable de sa revue. Tenir un suivi compact des fichiers restant à examiner ; une famille entière ne peut pas être déclarée revue sur la base de quelques fichiers représentatifs. Conserver les listes détaillées temporaires dans `/tmp` et noter les exceptions durables dans la documentation existante.

  **Terminé lorsque :** les 51 dossiers sont attribués, le chemin build → TAST → compilateur → Go → exécution est explicable, les références et les limites connues sont identifiées. Les constats sont vérifiés, sans transformer chaque soupçon en refactor obligatoire.

- [ ] **2 — Rendre installation, configurations et dépendances compréhensibles.**

  Examiner `package.json`, fichiers Spago et lockfiles, liens symboliques, `bin/setup`, `bin/pkg`, fichiers Go et règles d’exclusion dans tous les dépôts. Clarifier la liste core et la liste complète des paquets, les versions d’outils, les chemins locaux et la façon de travailler sur un paquet isolé.

  Traiter les références absentes ou obsolètes après vérification. Réduire la duplication des configurations lorsque le partage reste compatible avec des dépôts indépendants. Retirer les fichiers de configuration dont l’absence d’usage est établie ; préserver les configurations encore utilisées par les workflows upstream ou JavaScript.

  **Terminé lorsque :** les dépendances et commandes de démarrage ont une source claire, chaque paquet dispose d’un parcours documenté, et un build ne dépend pas d’un lien ou fichier temporaire créé au hasard d’une session précédente.

- [ ] **3 — Installer une base de tests maintenable en PureScript.**

  Répartir les contrôles entre fixtures PureScript de bout en bout, tests internes PureScript et tests Go du runtime/parser. Rendre chaque catégorie accessible par une commande claire. Les déclarations de tests dans les configurations doivent correspondre à des entrées qui existent et fonctionnent.

  Reprendre les neuf tests spécialisés de `tools/` : records boxés, payloads de constructeurs éliminés, workers importés, division entière, tags mixtes et natifs, conversions record/tuple, initialisation récursive et fonctions sans argument. Réutiliser les fixtures existantes ; construire les AST nécessaires dans une suite PureScript typée. Partager les quelques helpers de préparation utiles et supprimer les montages JavaScript remplacés.

  Réexaminer les fixtures, snapshots, compagnons et exclusions existants. Préserver une régression pertinente ; supprimer un diagnostic ponctuel ou une assertion redondante avec un motif explicite. Aucun portage systématique des 43 cas pour atteindre un nombre arbitraire de tests.

  **Terminé lorsque :** un développeur PureScript peut lire, ajouter et exécuter un test du compilateur dans le langage du projet, les cas retenus ont une correspondance explicite, et les anciens scripts remplacés sont retirés.

- [ ] **4 — Simplifier le build et les runners de tous les dépôts.**

  Expliquer puis décider ensemble le sort de chaque outil : `embed-runtime.mjs`, `build-ffi.mjs`, `ffi-runner.mjs`, runners, sélection, processus, workspaces et snapshots. Pour chacun : phase d’exécution, entrées, sorties, consommateurs et raison éventuelle d’une interface JavaScript.

  Garder la logique évolutive et testable accessible aux développeurs PureScript. Conserver uniquement les interfaces de bas niveau et l’amorçage proportionnés à leur besoin, avec une organisation explicite. Réduire la duplication entre `bin/test`, `bin/modtest` et les scripts frères sans créer une nouvelle collection de petits wrappers.

  Limiter le nettoyage aux sorties du run concerné. Préserver sélection, reprise, journaux utiles, arrêts sur erreur, interruption des processus enfants et mise à jour intentionnelle des snapshots. Rendre visibles les effets sur les configurations et caches ; conserver les chemins nécessaires aux dépendances locales dans les workspaces temporaires.

  **Terminé lorsque :** les commandes sont simples à expliquer et à maintenir, leurs effets restent locaux, leurs contrats utiles sont testés et chaque script restant a une nécessité établie. Le bundle et la reconstruction du parser fonctionnent sans montage implicite.

- [ ] **5 — Rendre le pipeline et les métadonnées du compilateur lisibles.**

  Revoir `Main`, `CodeGen`, `CodegenState`, `ExprContext`, `GlobalTypes`, `ConstructorMetadata`, `ClassMetadata`, `AdtMetadata`, `ConstructorLayout` et `Monomorphization`, jusqu’aux fonctions auxiliaires.

  Nommer les étapes de préparation, spécialisation, optimisation et émission. Rendre claires les représentations transportées, leurs producteurs et leurs consommateurs. Préserver la séparation entre métadonnées immuables et état produit pendant l’émission ; rendre les mises à jour et dépendances d’ordre explicites.

  Simplifier les paramètres ambigus, traitements imbriqués, parcours répétés et fonctions qui mélangent collecte, décision et effets. Garder les variantes de constructeurs, classes et spécialisation explicites au lieu de les cacher dans une abstraction générique.

  **Terminé lorsque :** on peut suivre un module TAST jusqu’à ses déclarations Go et localiser chaque décision de représentation, sans remonter une chaîne de helpers opaques ni consulter l’historique du projet.

- [ ] **6 — Revoir l’intérieur des émetteurs et des passes d’optimisation.**

  Couvrir `ModuleBindings`, `BindingExprs`, `FunctionExprs`, `CallExprs`, `CallArguments`, `CallAnalysis`, `ArrayIntrinsics`, `AdtExprs`, `RecordExprs`, `PrimitiveExprs`, `ControlExprs`, `EffectExprs`, `ExprAnalysis` et `ThunkFusion`.

  Pour chaque famille, séparer clairement reconnaissance du cas, préparation typée, choix de représentation et émission. Réduire les longues expressions et blocs imbriqués lorsque des étapes nommées rendent leur ordre visible. Examiner les gardes répétées, noms générés, chemins de secours et usages de chaînes Go.

  Préserver les priorités entre appels, intrinsics et TCO, l’évaluation unique des arguments, les frontières des closures et les marqueurs de tableaux frais. Documenter localement les préconditions des optimisations conservées et les relier aux tests appropriés.

  **Terminé lorsque :** chaque famille a un parcours principal lisible, ses variantes sont identifiables et ses invariants vérifiables ; les extractions améliorent la lecture au lieu de simplement réduire la taille des fichiers.

- [ ] **7 — Clarifier les conversions, l’AST Go et les bridges FFI.**

  Revoir `GoTypes`, `GoConversions`, `GoAst`, `GoCode`, `GoFunctions`, `GoImports`, `ModuleDeclarations`, `Printer`, `FfiTypes`, `FfiSupport` et `FfiBridge`, avec leurs compagnons FFI.

  Donner des contrats explicites aux conversions : type source, type cible, champs concernés, boxing et helpers Rebox requis. Mutualiser les enveloppes réellement identiques, préserver l’ordre des demandes transitives et utiliser les nœuds Go existants pour les structures répétées.

  Dans la FFI, distinguer signature analysée, rapprochement avec les annotations TAST, adaptation des arguments/callbacks/retours et émission du fichier. Centraliser les règles communes par famille ; préserver les conventions des effets, newtypes et valeurs opaques.

  Examiner les traces de debug, panics et bridges de secours : distinguer diagnostic utile, contrat non pris en charge et reste d’investigation. Conserver un chemin d’erreur explicite, sans supprimer un contrôle nécessaire uniquement pour nettoyer le code.

  **Terminé lorsque :** les décisions sont compréhensibles avant le rendu, l’impression ne redécouvre pas le typage, les erreurs portent leur contexte et les fragments Go opaques restants ont une raison précise.

- [ ] **8 — Rendre le runtime Go et le parser accessibles aux mainteneurs.**

  Revoir toutes les fonctions du runtime canonique et du parser, ainsi que leur interface avec PureScript. Organiser les responsabilités : valeurs et tags, records, tableaux, fonctions, conversions, initialisation, analyse et renommage des déclarations FFI.

  Expliciter les contrats de mémoire et de représentation : pointeurs, `unsafe`, durée de vie, ownership, mutations et références. Nommer les étapes complexes et partager les traitements répétés lorsque cela reste idiomatique en Go. Conserver les protections nécessaires et des erreurs exploitables.

  Garder les sources canoniques identifiées. Vérifier l’embarquement du runtime, l’appariement WASM/runtime JavaScript et la reconstruction des artefacts distribués. Les fichiers tiers ou générés ne sont pas réécrits manuellement pour uniformiser le style.

  **Terminé lorsque :** les interfaces utiles sont expliquées depuis PureScript, les invariants internes du Go sont localisés et testés selon le besoin, et la distribution du compilateur reste autonome.

- [ ] **9 — Revoir les bibliothèques de base, jusque dans chaque fonction.**

  Dossiers : `gopurs-prelude`, `gopurs-effect`, `gopurs-console`, `gopurs-assert`, `gopurs-exceptions`, `gopurs-partial`, `gopurs-unsafe-coerce`, `gopurs-functions`, `gopurs-refs`, `gopurs-st`, `gopurs-lazy`, `gopurs-integers`, `gopurs-numbers`, `gopurs-enums`, `gopurs-strings`, `gopurs-strings-extra`, `gopurs-unfoldable`, `gopurs-foldable-traversable`, `gopurs-arrays`.

  Appliquer la grille à chaque source, compagnon et test. Examiner particulièrement les conversions numériques, indices et bornes, encodages de chaînes, conventions curryfiées/non curryfiées, références, mutations et évaluation différée. Rendre explicites les correspondances entre signatures PureScript et implémentations Go.

  Conserver la documentation et la structure upstream lorsqu’elles sont déjà lisibles. Éviter une réécriture stylistique de bibliothèques stables ou une généralisation qui complique leur synchronisation.

  **Terminé lorsque :** les 19 paquets ont été revus intégralement, leurs adaptations locales sont compréhensibles et les changements sont vérifiés par les tests pertinents des paquets concernés.

- [ ] **10 — Revoir collections, records, valeurs étrangères et JSON.**

  Dossiers : `gopurs-catenable-lists`, `gopurs-ordered-collections`, `gopurs-foreign`, `gopurs-foreign-object`, `gopurs-record`, `gopurs-nullable`, `gopurs-variant`, `gopurs-argonaut-core`, `gopurs-yoga-json`.

  Examiner représentation des collections et records, distinction absence/null/valeur, ordre des champs, conversions récursives, encodage/décodage et erreurs. Rendre lisibles les parcours génériques et leurs cas particuliers ; isoler les interfaces de bas niveau nécessaires à la FFI.

  Reprendre les fichiers denses, dont `Yoga.JSON`, par responsabilités et contrats, sans morceler les définitions qui se comprennent ensemble. Vérifier les tests de données imbriquées et de cas limites réellement touchés.

  **Terminé lorsque :** les neuf paquets ont une correspondance claire entre API PureScript, représentation Go et comportement d’erreur ; chaque fichier maintenu a été examiné.

- [ ] **11 — Revoir effets asynchrones, annulation et interpréteurs.**

  Dossiers : `gopurs-aff`, `gopurs-avar`, `gopurs-free`, `gopurs-run`, `gopurs-js-promise`, `gopurs-js-promise-aff`.

  Rendre explicites les états, transitions, files d’attente, callbacks, finaliseurs, propriété des ressources et règles d’annulation. Dans les implémentations Go, clarifier l’usage des goroutines, verrous et canaux ; dans les modules PureScript, séparer description des opérations et interprétation lorsque cette frontière existe déjà.

  Préserver les garanties de complétion, propagation d’erreur et exécution unique des callbacks. Examiner les implémentations JavaScript conservées selon leur provenance et leur usage réel. Utiliser les tests ciblés existants pour départager toute incertitude de concurrence avant de changer la structure.

  **Terminé lorsque :** les six paquets sont revus, leurs cycles de vie sont explicables et les changements ont des vérifications observables, sans faire reposer la validation uniquement sur les benchmarks purs d’altbak.

- [ ] **12 — Revoir nombres spécialisés, dates, horloges et génération.**

  Dossiers : `gopurs-js-bigints`, `gopurs-datetime`, `gopurs-js-date`, `gopurs-now`, `gopurs-random`, `gopurs-uuid`.

  Clarifier unités, précision, bornes, conversion des dates, source d’horloge et source d’aléa. Distinguer calcul pur et accès à l’environnement. Examiner les adaptations qui portent un nom historique JavaScript mais ciblent Go, et documenter leur contrat réel.

  Fournir le contexte de maintenance manquant, notamment pour `gopurs-js-bigints`, avec exemples et commande de test proportionnés. Garder les tests déterministes lorsque le contrat le permet.

  **Terminé lorsque :** les six paquets ont des types, unités, effets et limites compréhensibles, leurs fichiers sont revus et leur parcours de développement est clair.

- [ ] **13 — Revoir les huit bibliothèques Node et leurs ressources.**

  Dossiers : `gopurs-node-buffer`, `gopurs-node-event-emitter`, `gopurs-node-fs`, `gopurs-node-http`, `gopurs-node-net`, `gopurs-node-path`, `gopurs-node-process`, `gopurs-node-streams`.

  Suivre chaque API jusqu’à sa FFI : buffers et encodages, fichiers et descripteurs, événements et listeners, connexions et réponses, flux, fermeture, annulation et erreurs. Rendre explicites les limites de compatibilité des implémentations Go.

  Partager les conversions et traductions d’erreurs répétées avec des helpers locaux compréhensibles. Revoir les corps longs, callbacks imbriqués, interfaces trop générales et conventions implicites. Préserver les différences nécessaires entre variantes synchrones, asynchrones et Aff.

  Les tests doivent isoler fichiers, ports et processus et libérer leurs ressources. Retirer les sorties de tests accumulées ; garder les fixtures d’entrée et exemples utiles.

  **Terminé lorsque :** les huit paquets sont revus intégralement, leurs ressources ont un cycle de vie explicite et les tests concernés valident les comportements que le mode pure d’altbak ne couvre pas.

- [ ] **14 — Revoir les bibliothèques de test elles-mêmes.**

  Dossiers : `gopurs-quickcheck` et `gopurs-spec`.

  Clarifier générateurs, graines, assertions, arbre de tests, reporters, exécution, attentes et timeouts. Examiner les interfaces PureScript/Go, les exemples, fixtures et tests d’intégration. Vérifier les chemins permettant de développer et tester ces paquets sans dépendre d’un bricolage local.

  Conserver des helpers de test proportionnés et des diagnostics qui permettent de comprendre l’échec. Une bibliothèque de test doit simplifier les tests de gopurs sans créer une seconde infrastructure opaque.

  **Terminé lorsque :** les deux paquets et tous leurs fichiers maintenus sont revus, leurs commandes fonctionnent et leurs usages recommandés sont cohérents avec la suite retenue au lot 3.

- [ ] **15 — Vérifier la compréhension complète et clore le chantier.**

  Reprendre la liste des 51 dossiers et vérifier qu’aucun fichier maintenu n’a été oublié. Pour chacun, enregistrer un état final : revu et conservé, simplifié, remplacé ou supprimé. Les fichiers reconnus lisibles n’ont pas à être modifiés pour cocher la revue.

  Actualiser les README et documents existants : démarrage, responsabilités, phases du build, sources canoniques, FFI, commandes de tests et ajout d’un cas. Vérifier les chemins et supprimer les mentions obsolètes. Chaque interface JavaScript ou outil particulier restant doit avoir un rôle explicable et un propriétaire identifiable.

  Faire un parcours de maintenance représentatif : retrouver une règle de génération, lire sa fonction et ses invariants, trouver son test, comprendre une FFI de bibliothèque et déterminer le contrôle à lancer. Corriger les détours inutiles révélés par ce parcours.

  Réaliser les validations finales ci-dessous, noter le bilan et les limites résiduelles explicitement discutées. Une exception non comprise ou un paquet non revu empêche de déclarer l’objectif complet atteint. Les nouvelles fonctionnalités ou optimisations deviennent un chantier distinct.

## Validation et critères de clôture

- **Compilateur et build :** après chaque bloc qui les touche, `bin/go/run -c` depuis `altbak.pub` reste le jalon transversal convenu. Comparer le Go généré et les résultats sur les mêmes TAST/FFI. La dernière référence connue contient 300 entrées TAST, 387 fichiers Go et 14 résultats ; reprendre une référence fraîche lors de l’exécution du plan.
- **Bibliothèques :** lancer les contrôles pertinents des paquets modifiés. Le mode pure d’altbak ne prouve pas les comportements réseau, filesystem ou asynchrones. Regrouper les validations par famille ; ne pas répéter toute la campagne des 50 bibliothèques après chaque petite modification.
- **Tests et outils migrés :** vérifier les garanties retenues, notamment erreurs, isolation, interruptions et snapshots. Un snapshot modifié ou un test supprimé demande une explication de son effet sur la couverture.
- **Build final :** compiler les sources maintenues sans sorties compilées préexistantes, avec les dépendances et caches de paquets disponibles. Corriger les avertissements du code maintenu ; distinguer les éventuelles limites externes constatées. Contrôler aussi la distribution et les entrées réellement modifiées.
- **Performance :** aucune conclusion tirée d’un run isolé. Si un changement touche une représentation ou une allocation, comparer avec les baselines du README d’altbak et mesurer le chemin concerné.
- **Lisibilité :** revue de chaque fichier et de son organisation interne, commandes accessibles aux développeurs PureScript, usages de FFI et de mutation justifiés, absence de montages temporaires abandonnés. Ni le nombre de lignes, ni le nombre de fichiers, ni l’extension ne remplacent ce critère.
- **Fin :** les 15 lots et les 51 dossiers sont couverts, les comportements ont les validations appropriées et la documentation décrit le code final. Le chantier se ferme sans nouvelle vague implicite.

## Bilan des lots

Le lot 1 est réalisé ci-dessous. Les bilans suivants préciseront le périmètre réellement revu, les décisions, les changements conservés et le résultat des contrôles.

### Lot 1 — carte et référence du 14 septembre 2026

**Terminé pour l'inventaire et l'attribution.** Aucun fichier source n'est
déclaré revu en profondeur à ce stade. Les choix de langage, de découpage et de
commandes des lots suivants restent à examiner ensemble. Ce lot ajoute la
carte ci-dessous et le contexte durable dans [l'architecture](docs/architecture.md#carte-des-dépôts-et-des-consommateurs)
et [les validations](docs/testing.md#référence-du-lot-1-de-maintenance--14-septembre-2026).

#### Périmètre et suivi des fichiers

Les **51 dépôts** contiennent **2 255 fichiers suivis et 2 fichiers non suivis**
pertinents, soit **2 257 entrées attribuées individuellement**. Les deux ajouts
locaux sont `gopurs-quickcheck/spago.yaml` et
`gopurs-quickcheck/src/Test/QuickCheck/Gen.go` ; ils sont conservés. Les 14
fichiers ignorés présents hors répertoires de dépendances/sorties comprennent
les deux artefacts de build gopurs, dix lockfiles locaux et deux fichiers REPL.
Les dépendances installées et les sorties générées ont été distinguées des
sources maintenues ; elles ne constituent pas des fichiers source revus.

Dans la table, **S/T/O/C/D** compte les fichiers restant à revoir : sources
(compagnons inclus), tests/exemples/benchmarks/données, outils, configurations,
documentation/licences. Les liens symboliques suivis comptent comme une entrée.
Les dossiers d'intégration de `spec` appartiennent à T, y compris leurs
configurations et sorties attendues. Toutes les colonnes restent ouvertes
pour la revue interne prévue par la grille ; cartographier n'en retire aucun.

- **S et T** : lot indiqué pour la bibliothèque. Pour gopurs, S se répartit
  entre lots **5 : 10**, **6 : 14**, **7 : 13**, **8 : 6** fichiers ; T contient
  **490** fixtures/compagnons, **376** snapshots et **11** tests Node, au lot 3.
  Les tests Go du parser sont comptés dans S/lot 8 avec leur source canonique.
- **O : lot 4**, **C : lot 2**, **D : lot 15** dans chaque dépôt, avec mise à
  jour du contexte dans le lot concerné. `bin/pkg` appartient à C. La revue des
  interfaces des outils et du parser associe les lots 4, 7 et 8.
- **G : 2 fichiers distribués suivis**, hors S/T/O/C/D : `ffi_gen.wasm` et
  `wasm_exec.js` de gopurs, à vérifier par leur reconstruction au lot 8.

Le solde global est **S 516 / T 1 098 / O 63 / C 404 / D 174 / G 2**.
Une clôture partielle devra indiquer les chemins réellement examinés et
actualiser ce solde ; elle ne peut pas valider une famille par échantillonnage.
Le manifeste temporaire fournit chaque chemin, son état Git, son lot, son
type de fichier et son empreinte. Si `/tmp` a été purgé, repartir de
`git ls-files --cached --others --exclude-standard` dans chacun des 51 dépôts,
inventorier séparément les fichiers ignorés utiles, puis appliquer les familles
ci-dessus et les listes de modules des lots 5–8. Les fichiers nouvellement
ajoutés restent eux aussi à attribuer avant le bilan final du lot 15.

#### Dossiers, API et commandes actuelles

Chaque dépôt est propriétaire des API et compagnons mentionnés sur sa ligne.
Leur consommation passe par les imports PureScript et la résolution Spago,
puis par les bridges Go ou la FFI JavaScript selon le parcours. Les quatre
bibliothèques sans compagnon dans `src/` sont `catenable-lists`,
`js-promise-aff`, `run` et `variant` ; elles consomment les représentations de
leurs dépendances. Les relations d'imports, de réexports et d'overrides sont
inventoriées dans le détail temporaire ; les consommateurs par phase et les
usages dynamiques sont expliqués dans l'architecture.

Les commandes sont à lancer depuis le dépôt de la ligne :

- **B** : `npm run build` ; `./bin/test [fixture]`, `./bin/modtest [paquet]`,
  `npm run test:runner`, `npm run test:ffi` ; parser : `go test ./...` dans
  `tools/ffi-gen`. `build:runtime` et `build:ffi` reconstruisent leurs artefacts.
- **G** : `./bin/test` ; `-c` reconstruit aussi gopurs. Le script compile avec
  Spago, appelle `../gopurs/bin/gopurs --main Test.Main`, compile Go et lance
  `output/go_test_app`. **†** indique un nettoyage qui touche aussi les frères.
- **A†** : `./bin/test` d'assert ; build PureScript, génération Go sans `--main`,
  puis `go build ./...`, sans exécution d'une suite.
- **Q** : `npm run build` et `npm test` déclarent Pulp/JavaScript ; aucun
  `bin/test` Go. La configuration Spago locale déclare le paquet et ses
  dépendances, sans déclarer de suite `package.test`.

Ces commandes décrivent les entrées existantes, pas 51 parcours tous validés.
Les autres scripts npm/Pulp/Spago, benchmarks et workflows CI sont recensés
avec leurs configurations. Leur revue appartient aux lots 2 et 4 ; ils ne sont
pas remplacés par une nouvelle commande commune dans ce lot.

| Dépôt | Rôle et API principales | Commandes | S/T/O/C/D à revoir | Lot S/T |
| --- | --- | --- | --- | --- |
| `gopurs` | Compilateur TAST → Go ; `Main`, `Gopurs.*` | B | 43/877/13/8/7 | 3, 5–8 |
| `gopurs-aff` | Fibres, annulation, finaliseurs ; `Effect.Aff.*` | G† | 5/3/1/10/5 | 11 |
| `gopurs-argonaut-core` | Valeurs et parser JSON ; `Data.Argonaut.*` | G† | 7/3/1/10/5 | 10 |
| `gopurs-arrays` | Tableaux purs, non vides et ST ; `Data.Array.*` | G† | 15/11/1/7/3 | 9 |
| `gopurs-assert` | Assertions ; `Test.Assert` | A† | 3/0/1/6/3 | 9 |
| `gopurs-avar` | Variables asynchrones ; `Effect.AVar`, `Effect.Aff.AVar` | G† | 4/1/1/11/4 | 11 |
| `gopurs-catenable-lists` | Listes et files concaténables ; `Data.CatList`, `Data.CatQueue` | G† | 2/21/1/8/3 | 10 |
| `gopurs-console` | Sortie console ; `Effect.Console`, `Effect.Class.Console` | G† | 4/2/2/7/3 | 9 |
| `gopurs-datetime` | Dates, instants, heures, intervalles et générateurs ; `Data.Date*`, `Data.Time*`, `Data.Interval*` | G† | 22/1/1/7/3 | 12 |
| `gopurs-effect` | Effets synchrones et non curryfiés ; `Effect.*` | G† | 10/1/1/6/2 | 9 |
| `gopurs-enums` | Énumération, bornes et génération ; `Data.Enum.*` | G† | 5/3/1/6/3 | 9 |
| `gopurs-exceptions` | Exceptions et variantes partielles ; `Effect.Exception.*` | G† | 4/1/1/6/3 | 9 |
| `gopurs-foldable-traversable` | Plis et traversées indexées ou bifonctorielles ; `Data.Foldable*`, `Data.Traversable*` | G† | 17/3/1/7/3 | 9 |
| `gopurs-foreign` | Valeurs étrangères, index et clés ; `Foreign.*` | G† | 8/12/1/6/3 | 10 |
| `gopurs-foreign-object` | Objets et variantes ST ; `Foreign.Object.*` | G† | 12/4/1/7/3 | 10 |
| `gopurs-free` | Free/Cofree, trampoline, Yoneda ; `Control.Monad.Free`, `Control.Comonad.Cofree` | G† | 9/17/1/8/4 | 11 |
| `gopurs-functions` | Appels non curryfiés ; `Data.Function.Uncurried` | G | 3/1/1/6/3 | 9 |
| `gopurs-integers` | Entiers et opérations binaires ; `Data.Int.*` | G† | 6/2/1/7/3 | 9 |
| `gopurs-js-bigints` | Grands entiers ; `JS.BigInt` | G | 3/1/1/4/0 | 12 |
| `gopurs-js-date` | Interface date historique JavaScript ; `Data.JSDate` | G† | 3/3/1/15/4 | 12 |
| `gopurs-js-promise` | Promesses, rejet et variante différée ; `Promise.*` | G† | 8/3/1/7/3 | 11 |
| `gopurs-js-promise-aff` | Conversion Promise/Aff ; `Promise.Aff` | G† | 1/3/1/11/2 | 11 |
| `gopurs-lazy` | Évaluation différée ; `Data.Lazy` | G | 3/1/1/6/3 | 9 |
| `gopurs-node-buffer` | Buffers mutables, immuables, ST et encodages ; `Node.Buffer.*`, `Node.Encoding` | G† | 14/7/1/9/3 | 13 |
| `gopurs-node-event-emitter` | Événements et symboles ; `Node.EventEmitter.*`, `Node.Symbol` | G† | 6/2/1/8/2 | 13 |
| `gopurs-node-fs` | Fichiers sync/async/Aff et métadonnées ; `Node.FS.*` | G† | 18/6/1/7/3 | 13 |
| `gopurs-node-http` | HTTP/HTTPS, requêtes, réponses et serveurs ; `Node.HTTP.*`, `Node.HTTPS` | G† | 22/2/1/7/3 | 13 |
| `gopurs-node-net` | Sockets, adresses et serveurs ; `Node.Net.*` | G† | 16/1/1/7/3 | 13 |
| `gopurs-node-path` | Chemins ; `Node.Path` | G† | 3/1/1/7/3 | 13 |
| `gopurs-node-process` | Processus et plateforme ; `Node.Process`, `Node.Platform` | G† | 4/1/1/7/3 | 13 |
| `gopurs-node-streams` | Flux et adaptation Aff ; `Node.Stream.*` | G† | 4/16/1/8/3 | 13 |
| `gopurs-now` | Horloge ; `Effect.Now` | G† | 3/1/1/10/5 | 12 |
| `gopurs-nullable` | Valeur nullable ; `Data.Nullable` | G† | 3/4/1/11/5 | 10 |
| `gopurs-numbers` | Nombres, approximation et formatage ; `Data.Number.*` | G† | 7/1/1/6/3 | 9 |
| `gopurs-ordered-collections` | Maps et ensembles ordonnés ; `Data.Map.*`, `Data.Set.*` | G† | 7/7/1/6/3 | 10 |
| `gopurs-partial` | Calculs partiels ; `Partial.*` | G† | 6/1/1/7/3 | 9 |
| `gopurs-prelude` | Classes, primitives, instances et réexports ; `Prelude`, `Control.*`, `Data.*` | G† | 85/7/1/7/3 | 9 |
| `gopurs-quickcheck` | Tests de propriétés, générateurs et instances ; `Test.QuickCheck.*` | Q | 5/1/0/7/4 | 14 |
| `gopurs-random` | Aléa avec effets ; `Effect.Random` | G† | 3/1/1/7/3 | 12 |
| `gopurs-record` | Construction, copie et union de records ; `Record.*` | G† | 7/2/1/7/3 | 10 |
| `gopurs-refs` | Références mutables ; `Effect.Ref` | G† | 3/1/1/7/3 | 9 |
| `gopurs-run` | Effets extensibles et interpréteurs ; `Run.*` | G† | 7/7/1/7/2 | 11 |
| `gopurs-spec` | Arbre de tests, reporters et runner ; `Test.Spec.*` | G† | 29/30/1/11/12 | 14 |
| `gopurs-st` | Mutation locale, références et appels non curryfiés ; `Control.Monad.ST.*` | G† | 10/1/1/8/3 | 9 |
| `gopurs-strings` | Chaînes, points/unités de code et regex ; `Data.String.*`, `Data.Char.*` | G† | 28/10/1/8/3 | 9 |
| `gopurs-strings-extra` | Opérations de chaînes supplémentaires ; `Data.String.Extra` | G | 3/1/1/15/5 | 9 |
| `gopurs-unfoldable` | Dépliages ; `Data.Unfoldable`, `Data.Unfoldable1` | G† | 6/1/1/6/3 | 9 |
| `gopurs-unsafe-coerce` | Coercition non vérifiée ; `Unsafe.Coerce` | G† | 3/1/1/8/3 | 9 |
| `gopurs-uuid` | Identifiants UUID ; `Data.UUID` | G† | 3/1/1/11/3 | 12 |
| `gopurs-variant` | Sommes extensibles ; `Data.Variant.*`, `Data.Functor.Variant` | G† | 3/4/1/9/2 | 10 |
| `gopurs-yoga-json` | Encodage/décodage JSON générique ; `Yoga.JSON.*` | G† | 11/6/1/10/5 | 10 |

#### Référence, constats et limites

Les 51 dossiers sont attribués ; le parcours build → TAST → métadonnées →
monomorphisation → PBO → génération/bridges → Go → exécution est relié à ses
propriétaires. Le TAST du fork et ses types enrichis restent le contrat
d'entrée. Les **22 paquets core**, les **49 runners frères** et la liste des
**50 bibliothèques locales** sont trois ensembles distincts.

Référence fraîche : **43 tests Node réussis**, tests Go du parser réussis,
**373 fixtures listées** et **49 runners listés**. Le jalon
`altbak.pub/bin/go/run -c` réussit : **300 TAST**, **387 fichiers Go identiques
octet par octet à l'état d'entrée**, **14 résultats fonctionnels**. Le manifeste
npm inclut les artefacts distribués attendus. Les versions d'outils, commits,
sorties fonctionnelles et limites d'accès aux caches sont consignés dans
[docs/testing.md](docs/testing.md#référence-du-lot-1-de-maintenance--14-septembre-2026).
Aucun résultat de performance n'est déduit de cette campagne.

Les exceptions durables sont documentées : deux ajouts QuickCheck non suivis,
45 runners nettoyant les frères, six noms de dossiers absents référencés dans
les configurations, différences de noms Spago, entrées de tests à clarifier,
cinq modules étrangers sans compagnon Go adjacent et absence de README/licence
suivie dans `js-bigints`. La résolution Aff hors ligne réussit malgré ses trois
entrées absentes ; ces constats n'imposent donc pas tous une correction.

Les usages dynamiques sont repérés et les sources de 33 bibliothèques ont été
comparées aux versions du registre déjà en cache : **224 fichiers identiques,
36 différents et 79 ajouts**, avec versions et chemins conservés temporairement.
QuickCheck correspond au tag upstream v8.0.1 pour ses fichiers suivis. Aucune
comparaison aux derniers HEAD upstream ni revue interne complète n'est revendiquée.

Preuves et listes détaillées : `/tmp/gopurs-lot1-20260914/` contient
`inventory.json`, `files.tsv`, `dynamic-uses.tsv`, `registry-comparison.json`,
les listes de fixtures/runners, les manifestes et les logs des validations.
Ce sont des données temporaires ; la carte, les responsabilités, les résultats
et les exceptions utiles sont conservés dans ce todo et les documents existants.
Les lots **2 à 15 restent ouverts** ; aucun refactoring, changement de
configuration, suppression de fichier ni commit n'a été réalisé par ce lot.
