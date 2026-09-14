# Gopurs — compréhension et maintenance par des développeurs PureScript

## Résultat attendu

Un développeur PureScript doit pouvoir installer le projet, suivre la génération d’un programme, trouver le responsable d’un comportement, comprendre son code interne, le modifier et lancer le contrôle approprié sans reconstituer l’historique des investigations de l’IA.

Ce plan couvre **gopurs et les 50 bibliothèques locales `gopurs-*`** : sources PureScript, FFI Go et JavaScript, runtime, parser, tests, scripts, configurations et documentation. Le précédent todo est remplacé ; son historique reste dans Git.

La priorité est la compréhension et la maintenance. **Le nombre de `.mjs` n’est plus un critère de réussite autonome.** Les scripts temporaires disparaissent ; la logique et les tests qui appartiennent à PureScript y sont intégrés ; une petite interface technique peut rester en JavaScript si sa nécessité et son fonctionnement sont clairs. Un changement d’extension ou de dossier ne suffit pas à améliorer le code.

Le présent travail prépare uniquement le plan. Sa mise en œuvre se fera ensemble, lot par lot.

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

- [ ] **1 — Établir la carte et les références des 51 dossiers.**

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

À remplir au fil du travail, en quelques lignes par lot : périmètre revu, décisions expliquées ensemble, changements utiles, éléments conservés volontairement et résultat des contrôles. Aucun lot d’implémentation de ce nouveau plan n’est encore réalisé.
