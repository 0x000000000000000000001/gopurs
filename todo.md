# Gopurs — réduire le temps de compilation de b8x

État au 22 septembre 2026 : **records compacts et traversée FFI Value intégrés**. Diagnostic applicatif JSON général : **61,042 → 45,816 ms (−24,94 %)**, décodage seul **−27,14 %**, allocations combinées **−24,84 %**. JS actuel **8,647 ms**, écart restant **×5,30**. Le TAST profite peu de ces changements : **582,022 → 567,609 ms (−2,48 % observés, plages qui se recoupent)**. Baselines README précédentes : Go **57,43 / 573,63 ms** pour JSON général / TAST. Aucun nouveau b8x complet. Détail et suites générales au point 12 ci-dessous.

**Comparaison des pistes, 22 septembre :** 16 variantes × 3 processus, puis sonde de mises à jour immuables. Sur le **vrai Go généré modifié en scratch**, fusionner le décodage des tableaux donne **45,364 → 35,193 ms (−22,42 %)** ; avec bêta-réduction **34,709 ms (−23,49 %)**, JS témoin **9,051 ms**. Dans le moteur manuel partagé, modifier seulement l'accès au dictionnaire gagne **6,02 %** ; enlever aussi les enveloppes `Value` aux appels/résultats accélère le décodage de **×4,35**. Les prototypes natifs complets restent vers **8,3–8,5 ms** avec parsing. **Aucun de ces nouveaux gains n'est intégré.** Priorité immédiate : fusion des tableaux ; chantier architectural : appels/résultats natifs et stockage stable, fonctions polymorphes partagées. Interfaces et descripteurs restent des outils complémentaires ; les chaînes naïves de wrappers se dégradent fortement. Détails dans la première section du point 12.

Étape précédente : **fusion de `resolveArgs` PBO intégrée**, avec gain Go modeste sur le diagnostic JSON/TAST : **590,273 → 573,631 ms (−2,82 %)** et **595,948 → 586,711 Mio (−1,55 %)** par corpus. JS : **83,381 → 75,899 ms**. Erreurs, cycles et AST préservés ; détails ci-dessous.

Étape précédente : **applications multiarguments du runtime corrigées et passerelle JSON allégée**. Sur le diagnostic altbak JSON/TAST (12 modules, 1 P, GOGC=100), Go passe de **720,986 à 599,956 ms (−16,79 %)** et alloue **19,98 % de moins** ; JS contrôle **85,856 ms**, soit encore **×6,99**. Compilateur reconstruit, résultats structurels identiques, détail ci-dessous.

État précédent du compilateur complet : **indexation Go corrigée et chargement TAST natif à 8 workers par défaut**. Sur les 238 modules de gopurs-aff : backend natif **9,167 → 7,8185 s (−14,71 %)** dans les séries mesurées ; dernier contrôle JS **6,9665 s**, soit encore **12,23 % d'écart**. Le vrai `gopurs-aff/bin/test -c` passe entièrement : backend **7,864 s**.

**Sous-étape précédente, point 11 : suppression d'une copie du tableau entier à chaque lecture d'indice**, révélée par pprof dans la liste d'attente du décodeur de types. Le chargement Aff passe de **2,9315 à 1,3380 s** avec les deux changements. Sur b8x, contrôle limité au chargement : **21,1255 → 13,5975 s** avec 1 puis 8 workers (−35,64 %, +0,650 Gio de pic RSS de cette phase). **Pas de nouveau backend b8x complet** : le dernier `b -c` avant ces changements reste **181,220 s de backend / 267,697 s de commande**, référence initiale historique 591,665 s. Les points 3, 4, 6, 7, 8, 10, 11 et 12 restent ouverts.

## Points à traiter un par un, ultérieurement

**Priorité : point 12, intégrer la fusion du décodage des tableaux dont le gain est maintenant mesuré sur le Go généré.** La version scratch conserve chaque appel au décodeur, son ordre et la première erreur : **−22,42 % au total**. Ensuite, établir sur du vrai PureScript une **convention d'appel native pour des fonctions polymorphes partagées**, gardant records et résultats natifs à travers l'appel. Les sondes isolées montrent que changer seulement le dictionnaire est insuffisant ; les enveloppes et intermédiaires coûtent davantage. Comparer interfaces/accesseurs et descripteurs sans dupliquer le corps par record, mesurer allocations et taille. Éviter les chaînes de wrappers : 32 mises à jour font passer une lecture à 525 ns contre 20 ns avec les records actuels. La bêta-réduction ciblée est utile mais son gain total isolé reste modeste. Les sondes manuelles vers 8 ms ne prouvent pas encore une traduction générale par gopurs.

**Point 11 toujours ouvert :** la FFI de traversée est allégée, mais le TAST ne gagne que modestement. Reprofiler ce décodeur avant sa prochaine optimisation propre ; `resolveArgs` et les applications à 6–10 arguments sont déjà corrigés. Le parseur n'a pas été changé. Confirmer ultérieurement le gain global et la mémoire sur le vrai b8x (point 7). Le prochain gros chantier de parallélisme reste le point 4 : ordonnancement PBO entre modules avec transmission des directives conservée. Défauts natifs : chargement 8, émission 8, préparation 2 ; chargement JS 1.

Travailler sur un seul point à la fois. Commencer par une expérience courte, vérifier le comportement et mesurer le gain avant de passer au suivant. Les chiffres du diagnostic ci-dessous sont des observations ; les gains proposés restent à établir.

- [x] **1. Comparaisons Char/String et adaptation FFI.** Deux signatures Go génériques dans gopurs-prelude, arguments génériques transmis directement par le bridge. Suppression des trois allocations/72 octets par comparaison confirmée ; tests JS/natifs et mesure complète b8x terminés. Backend : 591,665 → 493,217 s ; allocations : 1 030,31 → 858,66 Gio. Bilan et limites ci-dessous.
- [x] **2. Prédicat de caractères du scanner d’imports Go.** Prédicat d’identifiant avec une conversion en entier et des conditions explicites : moins de chaînes et de closures. Microbenchmark ×3,04 ; backend b8x 493,217 → 449,638 s (−8,84 %), allocations 858,66 → 789,39 Gio. Les 2 959 Go produits restent identiques. La propagation des imports pour éviter des scans reste une piste distincte, non implémentée. Bilan ci-dessous.
- [ ] **3. Préparation et monomorphisation.** Une garde dans `TypeSubstitution.substitute` est intégrée : elle concerne surtout l’optimisation/émission, pas `Substitute.substituteExprType` de la monomorphisation. Gain backend observé : 449,638 → 433,732 s. Deux gardes évitent maintenant les clés de spécialisation inutiles : préparation 116,209 → 100,035 s, backend 433,732 → 424,522 s. Le cache local du préfixe de `transitiveCollect` est maintenant intégré : préparation 96,746 → 81,535 s, backend 388,011 → 374,255 s ; 2 655 TAST et 2 959 Go identiques. Les contributions des passes `monomorphizeExpr`/`collectExpr` sont maintenant aussi mises en cache avec invalidation des dépendances : préparation **82,013 → 45,096 s**, backend **268,011 → 223,904 s**, allocations −62,93 Gio. 59 tests passent et toutes les sorties b8x sont identiques. Les autres coûts de préparation restent ouverts.
- [ ] **4. Parallélisme applicatif au-delà de l’émission.** Le producteur PBO séquentiel optimise pendant que le lot précédent émet du Go : optimisation/émission 289,795 → 256,177 s, backend 424,522 → 388,011 s. **Collecte transitive désormais parallèle**, deux workers Aff natifs par défaut (`GOPURS_PREPARE_JOBS`, borne 1–8), snapshots immuables, fusion ordonnée et barrière entre tours. Préparation **45,706 → 37,602 s (−17,73 %)**, backend **207,831 → 196,494 s** ; 20 tests ciblés et sorties b8x identiques. Le prototype de petits lots successifs est écarté : copies quadratiques mesurées ; la version retenue partitionne une seule fois par tour. **Chemin partagé de `forceEscape` supprimé**, échappement au tas et concurrence validés ; contrôle rapproché b8x **210,202 → 189,835 s**. **Balayage des workers d’émission terminé : émission 8 par défaut**, sorties identiques et vrai workflow validé ; préparation maintenue à 2. Chargement natif désormais à 8 après les contrôles du point 11. Prochaine étape : traiter explicitement la transmission des directives entre modules avant de paralléliser l’optimisation PBO elle-même.
- [x] **5. Échecs masqués dans le build b8x.** Le second `purs compile` est supprimé : Spago génère déjà les sourcemaps JS. Configuration réservée à JS via un outil Node indépendant de Docker/programme compilé ; PHP/Go utilisent leur environnement à l’exécution. Les erreurs utiles sont visibles et propagées. **12 tests passent**, mini-build réel de 54 modules avec les vrais templates JS (65 placeholders, aucun restant) et lecture de configuration PHP validés. **Le nouveau `b -c -n` Go complet termine sans erreur**, sans second passage sourcemaps ni appel de configuration ; détails et limites ci-dessous.
- [ ] **6. GC, après réduction des allocations.** Dernier profil : **635 cycles, 44,915 ms de pauses globales**, pic RSS **6,203 Gio**. Il reste à évaluer, si nécessaire après réduction des allocations, un réglage du GC par comparaison contrôlée temps/mémoire. Ce chantier est secondaire ; les pourcentages CPU des workers, notamment idle, ne sont pas une promesse de gain mural.
- [ ] **7. Validation continue sur le vrai workflow.** Le dernier bilan complet b8x précède la correction d’indexation et le chargement à 8 ; ces changements sont validés par le workflow Aff complet et le chargement b8x isolé. Refaire le workflow b8x complet pour confirmer leur bilan global et leur pic mémoire. Ce point reste une exigence de validation des prochaines optimisations, pas un gain autonome à rechercher. Après validation de chaque changement isolé, répéter `b -c` (natif par défaut) avec les mêmes entrées, paramètres et conditions de cache ; publier le détail des phases, pic mémoire et statuts, et les allocations lors des campagnes profilées. Le dernier passage chargement est volontairement sans pprof/gctrace pour rester comparable aux mesures de concurrence ; il ne remplace pas le dernier profil GC. Comparer aux références de compilation b8x ci-dessous ; garder les baselines d’exécution officielles d’altbak.pub pour leur périmètre propre.
- [ ] **8. Coût interne des passes PBO et de l’émission Go.** `Array.any` générique et cache FIFO de 512 instanciations par module intégrés ; collecte transitive accélérée aux points 3/4. **Boxage des pointeurs ADT structuré**, texte Go et imports préservés : backend **229,134 → 219,560 s**, scanner **54,77 → 31,62 Gio**, allocations totales **−12,42 Gio** ; 22 tests et sorties b8x identiques. La suppression isolée de Str/StrVal est écartée (aucune allocation économisée). **Cache externe PBO intégré** : backend **219,560 → 207,831 s**, allocations **−7,94 Gio** ; 41 tests et sorties b8x identiques. Dernier profil : `Semantics.build` **45,82 Gio**, `makeExternEvalSpine` **35,67 Gio**, `lookupImplementation` **13,47 Gio**. Sonde de 509 484 constructions sur 304 modules : seulement **0,32 % de répétitions par identité**, 8,38 % d’enveloppes répétées surtout feuilles ; cache général non retenu. Restent les recherches locales et analyses/constructions plus ciblées, à mesurer. Pour poursuivre l’AST structuré sur records/tableaux, surveiller les imports (`exprImports` : 10,48 Gio dans le profil précédent). Le prototype parallèle a aussi révélé des conversions coûteuses dans le bridge `Array.slice` ; leur gain potentiel sur le chemin actuel reste à établir. Coûts cumulés, à ne pas additionner.
- [x] **9. Indexer les champs de conversion Go.** Index immuable construit une fois depuis les layouts TAST, partagé entre émissions. Premier résultat dans l’ordre des clés, priorité constructeur/classe et deux alias conservés. `findReboxFields` passe de **150,56 Gio à 69 Mio** et de 53,78 à 0,04 s CPU échantillonnées hors GC. Backend b8x : **401,157 → 342,863 s**, allocations totales −21,27 %. Les **90 tests**, 4 468 recherches sur les métadonnées b8x et les comparaisons natives passent ; **2 655 TAST et 2 959 Go identiques**. Bilan et limites ci-dessous.
- [ ] **10. Compilation incrémentale entre builds.** Réutiliser les résultats des modules inchangés pour le cycle quotidien. Définir l’invalidation des signatures, directives PBO, spécialisations transitives, FFI et versions du compilateur ; comparer les sorties avec un build propre. Ce chantier est distinct de l’accélération de `b -c -n` : un cache invalidé par `-c` n’améliore pas ce rebuild forcé. Potentiel à mesurer sur une modification réelle de b8x, pas sur la seule commande propre.
- [ ] **11. Écart JS/natif au chargement TAST.** Descriptions des types décodées une seule fois dans PBO, puis correction générale de `OpArrayIndex` dans gopurs : convertir seulement l'élément sélectionné évite les copies répétées du tableau d'indices en attente. Cette ligne représentait **3,15 Go d'allocations** dans le profil Aff. À 1 worker : chargement **2,9315 → 2,4540 s** ; avec le nouveau défaut natif à 8 : **1,3380 s**. JS final : **0,5855 s**. Backend natif **7,8185 s / JS 6,9665 s** ; tests Aff complets et sorties JS/native comparées. Chargement b8x −35,64 % entre 1 et 8 workers, mémoire de phase +0,650 Gio ; validation du backend b8x complet à refaire. Le natif utilise déjà `encoding/json.Unmarshal` ; profiler les coûts restants avant de changer de parseur. Aucun cache entre builds ni perte d'information TAST.

- [ ] **12. Performances générales du code Go, révélées par JSON Decoding.** Corrections runtime/FFI intégrées : −24,94 % sur le JSON général dans leur comparaison contrôlée. **Comparaison des pistes terminée, nouveaux prototypes non intégrés.** Suite ordonnée : **(a)** fusion des tableaux/décodage connu, gain scratch −22,42 % sur le Go réellement généré ; **(b)** convention d'appel native pour fonctions polymorphes partagées, stockage stable et résultats natifs ; **(c)** suppression des closures/ADT immédiatement consommés ; **(d)** adaptation des records par interfaces ou descripteurs selon l'usage, sans chaînes de wrappers. Dictionnaire natif seul : 6,02 % dans la sonde, passage à des arguments/résultats natifs : ×4,35 sur le décodage de ce moteur manuel. Les prototypes complets avoisinent JS, sans démontrer la généralité de toute instance `DecodeJson`. Conserver rangées, immutabilité, ordre/erreurs/effets, fonctions échappantes et ABI FFI ; mesurer taille du code partagé. Gain propre à `Maybe` encore non établi. Tester JSON général puis TAST ; refaire b8x après intégration. Le parseur devient dominant seulement dans les prototypes natifs.

Chaque point doit laisser un résultat vérifié, les mesures avant/après et ses limites dans ce document. Consolider les gains cumulés historiques avec des séries de runs comparables.

## Point 12 — sélection des axes après comparaison, 22 septembre 2026

**16 variantes, 48 processus**, ordre alterné/rotatif, deux chauffes et cinq échantillons par phase, médiane des minima de processus ; 1 P, GOGC=100, PGO désactivé. Aucun build/test/profil concurrent. Baselines officielles README **45,82 ms Go / 8,65 ms JS**, inchangées ; les témoins de cette campagne donnent **45,364 / 9,051 ms**. Les nouvelles modifications restent en scratch.

| Vrai Go généré, modifié en scratch | Décodage (ms) | Parse + decode (ms) | Allocations décodage (Mio) |
|---|---:|---:|---:|
| Témoin | 38,327 | 45,364 | 56,729 |
| Application dans les branches + bêta-réduction | 35,124 | 44,020 | 54,489 |
| Fusion du décodage des tableaux | 26,939 | **35,193** | 42,107 |
| Fusion + bêta-réduction/retour du Left existant | 27,425 | **34,709** | **39,867** |
| JS témoin | 7,307 | 9,051 | — |

**Fusion : −22,42 % au total, −25,77 % d'octets de décodage.** Le décodeur fourni, son ABI, les constructeurs d'erreur et l'ordre restent identiques ; tous les callbacks sont appelés même après une erreur. Le résultat conserve la première erreur. Ce n'est pas une autorisation de court-circuiter un `traverse` Applicative arbitraire. La combinaison réduit les allocations de 29,72 %, mais ne démontre pas un gain de débit supplémentaire face à la fusion seule. Bêta-réduction isolée : −2,96 % au total, plages recoupées. Binaires stables ou plus petits (6 458 962 octets témoin ; 6 420 194 combiné).

**Sondes architecturales séparées**, même moteur partagé et même stockage entre variantes :

- Dictionnaire `Value`/RecordGet/Apply → dictionnaire natif conservant l'ABI Value : **2,835 → 2,664 ms**, −6,02 %, allocations identiques. Puis arguments/résultats natifs : **0,613 ms**, soit **×4,35**, allocations **4,659 → 1,035 Mio**. Ce résultat isole la frontière d'appel ; son tuple boxé n'est pas l'exact layout `Either` généré.
- Tableaux directs contre arbre de concaténations, résultats natifs constants : **1,141 → 0,656 ms** (−42,50 %). Résultats par valeur contre enveloppes allouées, tableaux directs constants : **1,013 → 0,656 ms** (−35,22 %). Ce modèle force les allocations d'enveloppes pour étudier le mécanisme ; il ne mesure pas exactement les ADT de gopurs ni un gain propre à `Maybe`. Ne pas sommer ces pourcentages.
- Reprise interfaces/descripteurs : **0,548 / 0,708 ms de décodage**, **8,501 / 8,324 ms avec parsing**. Structs typées plus sobres ; pas de vainqueur fiable au total dominé par le parsing. Les schémas de ces moteurs restent manuels.
- **Limite des interfaces imbriquées :** après 1/8/32 remplacements immuables du même champ, lecture **52,05 / 153,40 / 525,30 ns**, contre environ **20 ns Value / 2,8 ns slots plats**. Construction exclue ; toutes les versions intermédiaires restent accessibles via les wrappers. Ce défaut vient du chaînage naïf et peut être évité par un autre stockage ; ce n'est pas une limite intrinsèque des interfaces.

**Voie retenue : fonctions polymorphes partagées avec arguments/résultats natifs, records stables et construction fusionnée.** Exploiter les types du TAST ; garder les corps communs, adapter seulement layouts/accesseurs aux formes. Interfaces utiles comme vues, descripteurs utiles pour les rangées ouvertes ; pas de remplacement universel des records par des wrappers. Garder un repli générique pour les valeurs/fonctions échappantes, instances inconnues et FFI. Prochaine intégration concrète : la fusion des tableaux, puis une preuve générée du nouvel appel polymorphe. Le vrai prototype généré combiné reste **environ ×3,8 derrière JS** ; les 8 ms manuelles sont un potentiel architectural, pas un gain acquis.

**Validation :** 17 cas fixes pour toutes les variantes ; 317 différentiels JS pour les moteurs manuels. Les modifications du Go généré sont identiques au Go témoin sur 317 cas, dont 10 écarts JS/Go préexistants sur les Int hors bornes signées 32 bits (contrat natif documenté). Aucun cas chronométré n'est concerné. Ordre des callbacks, erreurs, immutabilité et résultats complets contrôlés. Pas de nouvelle mesure TAST/b8x ; aucune cellule README remplacée par un prototype.

[Rapport détaillé](/Users/0x1/Documents/htdocs/altbak.pub-gopurs/docs/benchmark-results/2026-09-22-json-options.md) · [Archive des résultats et scripts](/Users/0x1/Documents/htdocs/altbak.pub-gopurs/docs/benchmark-results/2026-09-22-json-options.json).

## Point 12 — comparaison architecturale, prototypes du 22 septembre 2026

Première campagne, conservée comme historique ; la comparaison plus large ci-dessus précise les contributions et remplace son ordre de priorité.

Expérience demandée sur **interfaces composées + structs privées** face à **descripteurs de layout + slots**, en conservant un corps de décodeur partagé entre tous les schémas. Le prototype matérialise les résultats ; aucun cache ni retour du JSON brut. Deux mesures distinctes : décodage du corpus général complet, puis opérations polymorphes sur 1/8/32 layouts. **Aucune nouvelle correction de production dans cette étape.**

| Médiane, même corpus (ms) | Go actuel | Prototype interfaces | Prototype descripteurs | JS |
|---|---:|---:|---:|---:|
| Parsing | 6,834 | 6,708 | 6,659 | 1,518 |
| Décodage | 35,191 | **0,525** | **0,602** | 7,007 |
| Parse + decode | 44,385 | **7,871** | **8,064** | 8,485 |

Baselines officielles README inchangées : **45,82 ms Go / 8,65 ms JS**. Les prototypes n'alimentent pas ces cellules. Trois processus par modèle, ordre alterné, deux chauffes/cinq échantillons, 1 P/GOGC=100/PGO off, contrôles hors mesure. Allocations décodage : **56,729 → 1,035 Mio** avec structs/interfaces (−98,17 %), ou 1,676 Mio avec descripteurs. **×5,64 sur parse + decode**, **×67 sur décodage** face au Go actuel dans cette expérience. Le prototype se situe près de JS au total ; son avance de 7,24 % dans cette série ne prouve pas une supériorité générale. **Les deux prototypes partagent exactement le même moteur** : les interfaces gagnent 12,73 % sur le décodage face aux slots. Le gros écart avec gopurs inclut aussi la suppression des dictionnaires/closures, enveloppes ADT et tableaux intermédiaires. Il ne mesure pas l'embedding isolé.

Les **17 cas de l'oracle indépendant** et **317 cas différentiels face au véritable décodeur JS** passent pour chaque modèle : valeurs complètes, erreurs/priorités, limites Int, champs absents, options, variantes et tableaux. Le moteur emploie des schémas manuels ; la généralité des instances arbitraires `DecodeJson`, des opérations de rangées et de l'ABI PureScript n'est pas démontrée. Pas de nouveau TAST ou b8x.

**Sonde records, 32 layouts :** lecture/calcul **12,24 ns Value / 5,41 ns interfaces / 1,89 ns descripteurs** ; construction+lecture **82,73 / 18,01 / 64,19 ns** ; insertion immuable+lecture **143,50 / 57,76 / 51,85 ns**. Les interfaces économisent la construction, les descripteurs gagnent en accès pur. L'extension par interface retient la source : chaînes d'extensions, rétention mémoire, suppressions et remplacement changeant le type restent à étudier. Tests : 32 × 256 valeurs avec tous les champs, immutabilité et GC.

**Taille :** 1/8/32 types, toujours **un seul symbole de fonction de calcul** ; binaire sans symboles **1 571 554 / 1 605 458 / 1 624 994 octets** (+3,40 % de 1 à 32). Les accesseurs/métadonnées croissent ; ce test ne prédit pas la taille complète de gopurs et ne mesure pas le coût d'une monomorphisation totale.

**Décision : poursuivre par une preuve générée depuis PureScript du stockage natif à travers le polymorphisme**, puis traiter conjointement les intermédiaires de contrôle/ADT/tableaux. Le prototype justifie ce chantier architectural ; il ne constitue pas une solution prête à généraliser ni une promesse de gain identique.

[Rapport et archive reproductible](/Users/0x1/Documents/htdocs/altbak.pub-gopurs/docs/benchmark-results/2026-09-22-json-architecture.md) · [Scripts exécutés](/Users/0x1/Documents/htdocs/scratch/json-architecture-20260922).

## Point 12 — records compacts et traversée FFI Value, intégrés le 22 septembre 2026

Le profil initial du décodage général attribuait environ **23 % des octets alloués directement à `RecordToMap`** : insertion des champs et conversion de petits dictionnaires. Trois sondes isolées, puis combinées, ont précédé l'intégration. `CoerceToStruct` évite map et tri pour 0/1 champ ; `RecordSet` insère/remplace sans modifier la source et conserve les représentations compactes ; `TraverseArrayImpl` et le bridge transmettent directement les `Value`, sans conversion `[]any` ni callbacks relais inutiles. Même algorithme de traversée, sans changement de PBO ou du parseur.

| Comparaison alternée, médianes (ms) | Go avant | Go après | JS après |
|---|---:|---:|---:|
| JSON général, décodage | 50,175 | **36,556** | 7,086 |
| JSON général, parse + decode | 61,042 | **45,816** | 8,647 |
| TAST, décodage | 466,219 | 457,856 | 57,565 |
| TAST, parse + decode | 582,022 | **567,609** | 79,059 |

Trois processus par version/runtime, deux chauffes et cinq échantillons par phase, médiane des minima par processus ; 1 P, GOGC=100, aucun build/profil en parallèle. **Allocations combinées JSON : 83,099 → 62,460 Mio (−24,84 %) ; TAST : 586,711 → 569,378 Mio (−2,95 %)**. Face aux cellules historiques README Go **57,43 / 573,63 ms**, gains observés **20,22 % / 1,05 %**. L'amélioration TAST reste faible avec recouvrement des plages ; ne pas extrapoler au chargement complet ni à b8x.

Les 24 processus retrouvent les mêmes oracles (17 cas JSON et 12 modules TAST). Bundle JS du JSON général identique ; pour TAST, le fingerprint avait été fusionné dans le module principal : imports/liste des sources réparés, définitions chronométrées et fingerprint inchangés vérifiés. Runtime complet `-race`, 15 tests FFI/records et 13 vérifications de runners passent. Compilateurs JS et natif reconstruits. Le profil d'allocations final place `RecordToMap` autour de **1,4 % direct** ; restent callbacks/applications, adaptateurs de primitives JSON, constructeurs `Either`/`Maybe`, concaténations et copies d'objets. Pourcentages cumulés non additionnables.

**Prochaine sonde concrète :** `gDecodeJsonCons` génère encore `(case result of Left e -> \_ -> Left e; Right r -> \k -> k r) (\r -> Right (insert field value r))`. Distribuer l'application puis réduire les lambdas peut supprimer deux fonctions intermédiaires. Tester l'évaluation unique du scrutateur, l'ordre, la première erreur/son chemin, l'absence d'insertion sur `Left`, les fonctions échappantes et les applications partielles. La fusion des enveloppes ADT est un axe distinct à mesurer ensuite, puis les traversées spécialisées. **Aucun gain futur chiffré n'est encore établi.**

[Rapport et mesures](/Users/0x1/Documents/htdocs/altbak.pub-gopurs/docs/benchmark-results/2026-09-22-json-general-runtime.md) · [Profils, sondes et builds figés](/Users/0x1/Documents/htdocs/scratch/json-general-20260922).

## Point 11 — fusion de resolveArgs, intégrée le 22 septembre 2026

Une boucle ST remplace `traverse resolveId`, `sequence Maybe`, puis `sequence Either`. Elle conserve tous les arguments dans l’ordre, attend si une référence manque, puis renvoie la première erreur. Une erreur déjà vue ne masque jamais une référence encore en attente. `force` et ses replis `Any` restent inchangés ; le tableau de résultat est local et figé après sa dernière mutation.

Le premier prototype avec `void`, mapping ST final et `unsafePartial` à chaque indice est écarté : allocations Go légèrement augmentées. La version retenue emploie des binds explicites et place `unsafePartial` autour de la boucle bornée, suivant le style du décodeur existant.

| Même corpus altbak, processus alternés | Go avant | Go après | JS avant | JS après |
|---|---:|---:|---:|---:|
| Parsing | 73,697 ms | 77,189 ms | 19,929 ms | 20,116 ms |
| Décodage seul | 492,369 ms | 475,123 ms | 61,034 ms | 58,586 ms |
| Parse + decode | **590,273 ms** | **573,631 ms** | **83,381 ms** | **75,899 ms** |
| Allocations parse + decode | 595,948 Mio | **586,711 Mio** | — | — |

**Go −2,82 % de temps et −1,55 % d’allocations ; JS −8,97 % de temps.** Le décodeur seul gagne 3,50 % en Go, 4,01 % en JS. Trois processus par version/runtime, deux chauffes puis cinq échantillons, médiane des minima ; ordre avant/après inversé au passage central. GOMAXPROCS=1, GOGC=100, PGO désactivé. Le parseur ne change pas ; ses variations de temps ne sont pas attribuées au patch. Le gain temporel est petit, soutenu par 9,237 Mio d’allocations économisées par corpus. L’écart Go/JS ne se referme pas sur cette étape : **×7,56** après, contre ×7,08 dans le témoin rapproché.

La baseline publiée immédiatement précédente était 599,956 ms Go / 85,856 ms JS ; le témoin frais est 590,273 / 83,381 ms. Utiliser ce dernier pour le gain de cette seule étape. Aucun nouveau b8x complet ni profil de son chargement parallèle : ne pas extrapoler ces pourcentages à la commande entière.

**1 000 tables différentielles identiques** (372 succès, 628 erreurs), entrées inchangées ; suite JS enrichie et **22 cas natifs** passent, dont six cas d’erreurs/priorités. Les douze processus finaux concordent avec les empreintes complètes JSON/AST de référence. Compilateurs JS/natif reconstruits pour intégrer la modification.

**Sous-étape `resolveArgs` terminée, point 11 ouvert.** Les parcours Row/contraintes et les conversions/callbacks FFI représentent des cibles plus larges à mesurer séparément ; ne pas prolonger cette micro-optimisation en lui prêtant un gros gain potentiel.

[Rapport et archive](/Users/0x1/Documents/htdocs/altbak.pub-gopurs/docs/benchmark-results/2026-09-22-resolve-args.md) · [Campagne et tests](/Users/0x1/Documents/htdocs/scratch/resolve-args-20260922).

## Point 11 — applications multiarguments et FFI JSON, 22 septembre 2026

Profil isolé du décodeur sur les 12 vrais TAST figés d’altbak : `Apply` représente **26,31 % des allocations directes échantillonnées**, notamment parce que `Apply6`…`Apply10` passent par des chaînes de closures alors que tous les arguments sont déjà disponibles. Le dispatch JSON a sept arguments. Une expérience courte confirme **610 → 509 ms** pour le seul décodage avec cette correction limitée.

Correction générale retenue : chemins directs `Apply2`…`Apply10` pour les fonctions saturées, une seule capture pour une application partielle jusqu’à `Func11`, fallback séquentiel conservé pour les fonctions qui retournent des fonctions. `CaseJsonImpl` dans gopurs-argonaut-core passe aussi directement ses arguments et son retour en `Value`, sans conversions en interfaces. **PBO, le parseur et le TAST restent inchangés.**

| Diagnostic altbak, corpus entier | Go avant | Go après | JS contrôle |
|---|---:|---:|---:|
| Parsing | 76,483 ms | 75,850 ms | 20,219 ms |
| Décodage | 616,670 ms | 507,676 ms | 62,653 ms |
| Parse + decode | **720,986 ms** | **599,956 ms** | **85,856 ms** |
| Allocations parse + decode | 744,781 Mio | **595,948 Mio** | — |

**−16,79 % de temps, −19,98 % d’allocations ; écart Go/JS encore ×6,99.** Face à la baseline publiée de 717,160 ms, le gain est de 16,34 %. Trois processus par version/runtime, deux chauffes et cinq échantillons par phase, médiane des minima de processus. GOMAXPROCS=1, GOGC=100 ; Node conserve ses threads de fond. JS est strictement le même bundle. Les fingerprints restent hors chrono, avec allocations de validation entre passes. Aucun nouveau workflow b8x complet : ces chiffres ne mesurent ni son gain global, ni son chargement parallèle par défaut.

Les neuf processus avant/après concordent avec les empreintes JSON/AST de référence. **501 fichiers Go sur 503 identiques** : seuls runtime et FFI Argonaut changent. Tests runtime optimisés et `-race` : 99 combinaisons d’arités, 27 panics, zéro allocation des applications saturées 2–10, durées de vie et concurrence. Suite Argonaut : 32 assertions/propriétés, 536 cas QuickCheck ; neuf contrôles de runners/publication passent. Compilateurs JS et natif reconstruits, binaire natif installé. Les manifests du diagnostic incluent désormais les sources des bibliothèques natives pour invalider les anciens builds après un changement FFI.

Le profil après correction attribue encore des allocations aux applications partielles, records, `Either`/`Maybe`, traversées et adaptations FFI. `TraverseArrayImpl` représente 19,23 % en cumul (callbacks inclus, pas un gain récupérable garanti). Prochaines sondes : fusionner les traversées du résolveur de types en conservant les erreurs, puis supprimer les conversions de tableaux réellement annulées immédiatement. **Point 11 encore ouvert.**

[Rapport et archive](/Users/0x1/Documents/htdocs/altbak.pub-gopurs/docs/benchmark-results/2026-09-22-json-typed-ast.md) · [Profils et expériences](/Users/0x1/Documents/htdocs/scratch/json-decode-20260922) · [Garanties du runtime](docs/runtime-closures.md).

## Point 11 — indexation sans copie et chargement parallèle, intégrés et mesurés

Le profil du corpus utilisateur **gopurs-aff, 238 modules** attribue 3,15 Go d'allocations à une seule conversion de tableau dans le décodeur de types. La liste d'indices en attente de notre précédente optimisation PBO exposait un défaut général de génération Go : chaque lecture convertissait et copiait tout le conteneur avant d'en sélectionner un élément. La correction indexe la représentation source puis convertit uniquement l'élément, en conservant son type TAST, l'ordre source/indice et leur évaluation unique. Une sonde isolée d'une ligne ramène les allocations totales instrumentées de 20,68 à 17,45 Go.

| Moyennes sur le même corpus Aff (s) | Natif avant, 1 worker | Indexation corrigée, 1 worker | Natif final, 8 workers | JS final |
|---|---:|---:|---:|---:|
| Chargement | 2,9315 | 2,4540 | **1,3380** | 0,5855 |
| Préparation | 1,2510 | 1,2310 | 1,3195 | 0,4805 |
| Optimisation / émission | 4,9790 | 5,0055 | 5,1575 | 5,8965 |
| Backend | 9,1670 | 8,6915 | **7,8185** | **6,9665** |

Deux observations par colonne, sorties neuves et TAST/FFI figés, GOMAXPROCS=14, GC normal, sans instrumentation. Première comparaison alternée ; contrôles finaux après reconstruction. **Indexation seule : chargement −16,29 %, backend −5,19 %, CPU −7,69 %.** Le balayage de workers confirme le gain supplémentaire : quatre workers 1,630/1,641 s, huit 1,277 s ; binaire final huit workers 1,354/1,322 s. L'écart final avec JS reste **12,23 %** ; les variations entre campagnes ne sont pas toutes attribuées aux changements.

**B8x, chargement seulement, ordre 1/8/8/1 : 21,1255 → 13,5975 s (−35,64 %)**, CPU 91,538 → 91,155 s, pic RSS **1,594 → 2,244 Gio**. Copie temporaire du même binaire, arrêt juste après chargement/tri, 2 655 TAST inchangés. Pas d'extrapolation à la commande entière ni au pic mémoire complet. Ce résultat justifie le défaut natif à huit ; `GOPURS_JOBS=1` reste disponible et JS conserve un worker.

**Quatre tests d'indexation/conversion passent**, dont le chemin réel `OpArrayIndex` et une mesure d'octets alloués indépendante de la longueur du tableau. Quinze cas de configuration natifs passent. **Le vrai `gopurs-aff/bin/test -c` passe entièrement**, compilation Go et suite Aff incluses : chargement 1,334 s, préparation 1,299 s, émission 5,230 s, backend **7,864 s**, commande 61,020 s. Le script reconstruit désormais le compilateur sélectionné, natif par défaut ou JS avec `GOPURS_JS=1`. Le binaire initial était déjà récent : ce défaut du script n'expliquait pas à lui seul l'écart utilisateur.

**238 TAST conservés, 293 Go identiques entre JS corrigé et natif corrigé**, workflow compris. Par rapport à l'ancien émetteur, seuls `Data_Array.go` et `Data_Array_Partial.go` changent (indexation et variables temporaires), différences revues et testées. Binaire reconstruit/exécuté/installé identique. Pas de nouveau `-race` complet ni de profil GC b8x. Les essais de validation interrompus par manque d'espace ou utilisant le `purs` standard de Homebrew sont exclus ; le passage réussi utilise bien le fork TAST et retrouve les entrées figées.

**Point 11 reste ouvert**, avec la préparation du point 3 et la validation b8x du point 7 comme suites utiles. Les références historiques b8x ci-dessous précèdent ces changements ; les baselines d'exécution officielles d'altbak.pub restent distinctes.

[Rapport et protocole](/Users/0x1/Documents/htdocs/scratch/gopurs-aff-gap-20260921/rapport.md) · [Mesures finales](/Users/0x1/Documents/htdocs/scratch/gopurs-aff-gap-20260921/final-comparison.json) · [Chargement b8x](/Users/0x1/Documents/htdocs/scratch/gopurs-aff-gap-20260921/b8x-load-workers.json) · [Workflow Aff validé](/Users/0x1/Documents/htdocs/scratch/gopurs-aff-gap-20260921/workflow-validation.json).

## Point 11 — historique du décodage des descriptions de types, avant correction de l'indexation

Le décodeur PBO relisait les mêmes champs JSON à chaque tour de résolution des références. Il prépare maintenant une description typée par entrée, puis ne repasse que sur les indices encore en attente. L'ordre ascendant et la publication immédiate sont conservés ; au point fixe, seul le premier indice restant est forcé comme avant. Les erreurs de champs tardifs restent différées pour conserver leur priorité face aux erreurs de références. Aucun changement du TAST ni état persistant entre builds.

Le candidat limité à la liste d'attente est écarté : chargement natif 3,438 → 3,570 s. La version complète passe la comparaison avec l'ancien décodeur sur **2 655 modules / 741 410 types**, plus les cas ciblés de cycles en cascade, partage, indépendance, indices invalides et priorité des erreurs. Reconstructions du compilateur sans erreur ni avertissement PureScript.

| Petit corpus, deux observations par version après chauffe | Avant | Après |
|---|---:|---:|
| Chargement natif | 3,4105 s | **2,9470 s** |
| Backend natif | 9,5200 s | **9,0670 s** |
| Chargement JS | 0,6545 s | 0,6075 s |
| Backend JS | 8,9780 s | 8,8720 s |

**Chargement natif −13,59 %, backend −4,76 %.** Le ratio total natif/JS vaut environ 1,022, mais le chargement reste 4,85 fois plus lent et le CPU total natif nettement supérieur. L'ancien écart global 6,8 / 12,9 s n'est pas reproduit dans cet état du code.

| B8x, même corpus figé, un passage par version | Avant | Après |
|---|---:|---:|
| Chargement | 31,492 s | **25,435 s** |
| Préparation | 37,075 s | 38,157 s |
| Optimisation / émission | 111,610 s | 112,714 s |
| Backend | 180,180 s | **176,312 s** |
| CPU utilisateur + système | 888,387 s | 859,751 s |
| Pic RSS | 6,076 Gio | 6,143 Gio |

**Chargement −6,057 s / −19,23 %, backend −3,868 s / −2,15 %, CPU −3,22 %, RSS +1,11 %.** Les variations de préparation/émission ne sont pas attribuées au changement. Les mesures courtes confirment le gain de chargement, mais les secondes b8x reposent sur un seul A/B.

Le build propre réel actuel **`b -c`** termine avec le code 0 : chargement **26,249 s**, préparation **39,736 s**, optimisation/émission **115,227 s**, backend **181,220 s**, commande **267,697 s**. Le natif est devenu le défaut et `-n` a été retiré séparément du lanceur ; ce chantier n'a pas modifié ces scripts. Il s'agit d'une validation d'intégration, pas d'une comparaison contrôlée de la commande entière avec les anciens lanceurs.

**304 TAST / 399 Go identiques dans les dix passages courts ; 2 655 TAST / 2 959 Go identiques dans les deux passages b8x et le workflow final.** 582 chemins sources FFI vérifiés inchangés, un seul `purs compile` applicatif, binaire mesuré et installé identique. Sans option bundle, le workflow ne compile pas le binaire applicatif final. Ces mesures ne remplacent pas le dernier pprof GC ; les baselines d'exécution d'altbak.pub restent distinctes.

**Correction conservée ; point 11 reste ouvert.** Prochaine sonde : coût et allocations du décodage AST natif restant. L'ordonnancement PBO entre modules du point 4 reste le gros chantier de parallélisme.

[Rapport et protocole](/Users/0x1/Documents/htdocs/scratch/gopurs-load-20260921/rapport.md) · [Mesures b8x](/Users/0x1/Documents/htdocs/scratch/gopurs-load-20260921/b8x-comparison.json) · [Validation complète](/Users/0x1/Documents/htdocs/scratch/gopurs-load-20260921/workflow-validation.json).

## Point 4 — workers réévalués, émission 8 retenue

Même binaire corrigé pour les balayages isolés, entrées figées, nouveaux dossiers de sortie, GOMAXPROCS=14, chargement=1, pipeline actif, GC par défaut, sans pprof/gctrace. Chaque réglage varie séparément, l’autre reste à 2. Émission=1 désactive aussi le chevauchement avec PBO.

**Petit corpus : une chauffe puis trois passages par combinaison**, ordre alterné, 22 processus au total. Émission 1/2/4/8 : optimisation/émission moyenne **6,915 / 5,438 / 5,517 / 5,115 s**. De 2 à 8 : phase −5,94 %, backend −4,37 %, CPU +0,61 %, pic RSS −3,08 %. Préparation 4/8 : collecte transitive plus rapide, mais gain de préparation totale proche de 2,5 %, sans gain backend. Les 304 TAST et 399 Go sont identiques dans tous les passages.

| b8x, ordre des passages | Préparation / émission | Préparation | Optimisation / émission | Backend |
|---|---|---:|---:|---:|
| Référence isolée | 2 / 2 | 45,191 s | 124,229 s | 203,089 s |
| Préparation à 8 | 8 / 2 | 39,087 s | 124,239 s | 199,523 s |
| Émission à 8 | 2 / 8 | 39,892 s | **113,532 s** | **186,326 s** |
| Retour référence | 2 / 2 | 41,351 s | 132,346 s | 208,951 s |
| Vrai `b -c -n`, nouveaux défauts | 2 / 8 | 42,498 s | **120,331 s** | **197,565 s** |

**Défaut d’émission porté à 8**, préparation conservée à 2. Sur les passages isolés, la phase optimisation/émission baisse de 8,6 à 14,2 % face aux témoins. Le gain total inclut aussi des variations de chargement/préparation : ne pas l’attribuer intégralement au réglage. Le résultat de préparation à 8 reste trop proche des variations à 2 pour modifier ce défaut. Seuls 2 et 8 ont été contrôlés sur b8x ; 1 et 4 l’ont été sur le petit corpus. Aucun optimum universel n’est revendiqué.

**Validation réelle : commande entière 304,738 s (5 min 05)**, backend 197,565 s, CPU 893,291 s, pic RSS 6,234 Gio. Elle reconstruit le compilateur et le TAST, avec aucune surcharge `GOPURS_*` ; c’est une observation supplémentaire à 8 avec un protocole distinct. Les étapes hors backend occupent environ 107 s. Le profil historique précédent donnait 290,815 s pour la commande, avec pprof/gctrace : ce nouveau passage ne démontre pas un gain global de commande. Des changements concomitants du lanceur/bootstrap, extérieurs à cette sous-étape, étaient aussi présents lors du workflow final ; les balayages sur binaire figé les contournent. Leur diff est conservé dans le rapport.

**2 655 TAST et 2 959 Go strictement identiques** dans les quatre passages isolés et le workflow final ; 582 chemins source/FFI surveillés inchangés. Le binaire reconstruit/exécuté est celui installé. Un seul `purs compile` applicatif, toutes les étapes réussissent, **10 tests d’émission passent**. Les garanties du pipeline restent inchangées ; cette sous-étape modifie une constante, sans nouvelle campagne race ni nouveau profil d’allocations.

**Suite : ordonnanceur PBO respectant les directives**, avec mesure avant intégration. L’optimisation des modules PBO demeure séquentielle. L’augmentation des workers existants apporte un gain réel mais limité ; elle ne suffit pas à saturer tout le pipeline. Les baselines d’exécution du README altbak.pub restent distinctes de ces durées de compilation.

[Rapport et mesures reproductibles](/Users/0x1/Documents/htdocs/scratch/gopurs-workers-20260921/rapport.md) · [Validation du vrai workflow](/Users/0x1/Documents/htdocs/scratch/gopurs-workers-20260921/workflow-validation.json).

## Point 4 — forceEscape corrigé, intégré et validé

Le runtime conserve l’échappement au tas par une affectation globale visible pour le compilateur, derrière une variable toujours fausse, selon le principe de `internal/abi.Escape` dans Go. L’écriture et le mutex ne sont plus exécutés. L’analyse du véritable bootstrap confirme l’échappement de `f` ; `forceEscape` et les constructeurs deviennent également inlinables. Aucun changement de PBO, d’Aff, du GC ni des nombres de workers.

**Trois tests Go et 66 sous-cas passent normalement et sous `-race`** : Func–Func11, 55 applications partielles, captures avec objets/pointeurs/tableaux, fonctions et goroutines créatrices terminées, piles sollicitées puis GC répétés, quatre créateurs et quatre appelants concurrents. Les cas réellement exécutés sont vérifiés. Builds JS/natif du compilateur sans erreur ni avertissement PureScript.

Petit corpus alterné après deux chauffes : **optimisation/émission 6,7225 → 5,9050 s (−12,16 %)**, backend **11,8445 → 11,2035 s (−5,41 %)**, CPU −6,09 %, pic RSS +10,44 %. Sur les six passages, 304 TAST et 398 Go sur 399 sont identiques ; seul le runtime attendu change.

Le vrai `b -c -n` profilé réussit : backend **202,215 s**, commande **290,815 s**. Face à la référence historique de 196,494 s / 277,148 s, cela ne démontre pas une accélération murale. CPU −2,32 %, allocations **375,66 → 372,34 Gio (−3,32 Gio)**, pic RSS **6,067 → 6,203 Gio**. Le chargement et la préparation augmentent ; leurs variations restent dans le bilan.

Un contrôle rapproché des deux binaires, ancien puis corrigé, sur les mêmes TAST b8x complète ce profil. Deux répertoires scratch, FFI résolues vers les mêmes fichiers, mêmes paramètres ; sans pprof/gctrace pour les deux variantes. Ce contrôle mesure le backend seul, sans reconstruire le compilateur ni relancer PureScript.

| Contrôle rapproché b8x | Référence | Correction |
|---|---:|---:|
| Chargement | 35,186 s | 33,013 s |
| Préparation | 40,166 s | 39,143 s |
| Optimisation / émission | 134,846 s | **117,677 s** |
| Backend | 210,202 s | **189,835 s** |
| CPU processus | 964,849 s | 902,882 s |
| Pic RSS | 6,108 Gio | 6,201 Gio |

**Backend −20,367 s / −9,69 %, optimisation/émission −12,73 %, CPU −6,42 %, pic RSS +1,51 %.** Les 3,196 s gagnées en chargement/préparation ne sont pas attribuées intégralement au changement. Une paire rapprochée, soutenue par le petit corpus alterné, confirme son intérêt sans garantir ces secondes sur toutes les compilations.

**2 655 TAST et 2 958 Go sur 2 959 identiques**, sans ajout ni retrait ; le runtime émis correspond exactement au fichier canonique corrigé. Les 2 748 fichiers surveillés restent inchangés pendant le profil et le contrôle ; binaires testé/profilé/installé identiques. Tous les processus réussissent, un seul `purs compile` applicatif, aucun appel de configuration JS. Sans option bundle, `b -c -n` ne compile pas le binaire applicatif final ; les baselines d’exécution altbak.pub restent distinctes.

**Correction conservée.** Prochaine étape : mesurer de nouveau les nombres de workers avant de réintroduire un ordonnanceur PBO. La durée complète de la commande reste à consolider ; le gain du contrôle backend n’est pas extrapolé à tout le workflow.

[Tests et petit benchmark](/Users/0x1/Documents/htdocs/scratch/gopurs-escape-fix-20260921/rapport.md) · [Profil complet et contrôle rapproché](/Users/0x1/Documents/htdocs/scratch/b8x-escape-fix-20260921/rapport.md) · [Comparaison rapprochée](/Users/0x1/Documents/htdocs/scratch/gopurs-escape-fix-20260921/b8x-pair-comparison.json) · [Garantie du runtime](/Users/0x1/Documents/htdocs/gopurs/gopurs/docs/runtime-closures.md).

## Point 4 — diagnostic historique de forceEscape, avant correction

Le runtime prend **un mutex global et écrit dans `EscapeSink` à chaque création de closure** (`Func`, `Func2`…`Func11`). Deux optimisations indépendantes sollicitent donc le même état partagé. L’ancien prototype à 91 paires n’est plus présent dans le scratch : la nouvelle sonde isole ce mécanisme sur le compilateur actuel, avec deux copies d’un même module, et ne prétend pas reconstituer exactement l’ancien run.

Trois modules coûteux (`Data.List`, `Data.Map.Internal`, `Data.List.Lazy`), mêmes entrées et purmeta figé avant publication, deux mémos privés neufs par évaluation. Variantes répétées dans les deux ordres, processus successifs, **18 observations par condition**. Le même binaire permet de contourner le bloc mutex/écriture ; le compilateur Go confirme que les closures continuent d’échapper au tas.

| Deux conversions, sans collection GC pendant la paire | Runtime actuel | Chemin global contourné |
|---|---:|---:|
| Séquentiel | 273,6 ms | 272,6 ms |
| Deux goroutines | 214,7 ms | **141,5 ms** |
| Accélération parallèle | ×1,27 | **×1,93** |
| CPU processus en parallèle | 422,2 ms | **282,8 ms** |
| Allocations en parallèle | 666,3077 Mio | 666,3075 Mio |

**−34,09 % de temps parallèle, −33,02 % de CPU**, effet confirmé sur les trois modules. Avec GC normal, comparaison keep/skip/skip/keep : **289,4 → 209,8 ms (−27,50 %)**, CPU −17,41 %. Le GC contribue donc au coût, mais n’est pas l’unique frein. Les variantes gardent les mêmes allocations à quelques Kio près ; aucune collection dans les 72 paires GCoff.

La trace montre les deux workers réellement actifs : environ 296 ms en état Go Running sur 303 ms chacun, environ 2 ms bloqués et 4,6 ms prêts mais non planifiés. Le mutex de `forceEscape` concentre 3,99 ms cumulées d’attente bloquée. L’essentiel du surcoût n’est pas une longue attente endormie : mutex, opérations atomiques et écriture partagée sont retirés ensemble dans la sonde. Leurs coûts internes respectifs restent non séparés.

**22 passages réussis, 304 TAST et 399 Go identiques à chaque fois**, binaire installé inchangé. Les résultats candidats sont utilisés pour poursuivre le build ; la comparaison des fichiers ne valide pas individuellement tous les replays comme IR. Pas de nouveau `b -c -n` complet, de test `-race` de remplacement ni de modification du runtime de production à ce stade. Deux copies d’un même module peuvent favoriser les caches CPU ; ne pas extrapoler ces pourcentages au compilateur entier.

**Étape suivante : remplacer le stockage global exécuté par un mécanisme d’échappement sans état partagé**, selon le principe employé par Go dans `internal/abi.Escape`. Vérifier les durées de vie des closures et la concurrence sous `-race`, puis mesurer le parallélisme déjà intégré et b8x. Réévaluer ensuite l’ordonnancement PBO, plutôt que multiplier immédiatement les workers. Les baselines d’exécution altbak.pub restent distinctes des durées de compilation.

[Rapport et limites](/Users/0x1/Documents/htdocs/scratch/pbo-parallel-cause-20260921/rapport.md) · [Comparaisons](/Users/0x1/Documents/htdocs/scratch/pbo-parallel-cause-20260921/comparison.json) · [Validation](/Users/0x1/Documents/htdocs/scratch/pbo-parallel-cause-20260921/validation.json) · [Runtime concerné](/Users/0x1/Documents/htdocs/gopurs/gopurs/runtime/runtime.go:483).

## Point 4 — collecte transitive parallèle, intégrée et mesurée

La sonde courte `Semantics.build` ne justifie pas un cache général par identité : **509 484 appels, 0,32 % de répétitions par identité et 8,38 % d’enveloppes répétées**, essentiellement des feuilles. Le travail passe donc au parallélisme de préparation. PBO expose un dispatcher monadique et conserve son API pure séquentielle. Chaque tour lit des snapshots immuables ; les contributions sont fusionnées dans l’ordre original avec priorité du premier payload, union des callers et invalidation du cache conservées.

Gopurs répartit une seule fois les travaux du tour en blocs contigus, exécutés via `defer`/Aff par des goroutines natives. **Deux workers par défaut**, réglables de 1 à 8 avec `GOPURS_PREPARE_JOBS`. Le prototype précédent à petits lots est rejeté : redécouper constamment le reste produisait des copies quadratiques, amplifiées par le bridge de `Array.slice`. Ce surcoût appartenait au prototype, pas à la référence de production.

**17 tests JS et 3 tests Go racines sous `-race` passent**, avec 5 sous-cas natifs : point fixe, hits du cache, fusion ordonnée malgré des achèvements désordonnés, indépendance, différé, réexécution, concurrence réelle et borne de workers. Builds JS/natif sans erreur ni avertissement PureScript. Petit corpus alterné après deux chauffes : préparation **1,2805 → 1,1920 s (−6,91 %)**, backend **10,9965 → 11,0850 s (+0,80 %)** ; aucun gain global établi sur ce petit corpus. 304 TAST et 399 Go identiques dans les six passages.

| Vrai `b -c -n` | Avant | Après |
|---|---:|---:|
| Chargement TAST | 33,482 s | 32,099 s |
| Préparation / monomorphisation | 45,706 s | **37,602 s** |
| Optimisation / émission | 128,640 s | 126,789 s |
| Backend | 207,831 s | **196,494 s** |
| Commande entière | 288,509 s | **277,148 s** |
| CPU backend utilisateur + système | 953,616 s | 946,781 s |
| Allocations estimées | 374,42 Gio | 375,66 Gio |
| Pic RSS backend | 5,911 Gio | 6,067 Gio |

**Préparation −8,104 s / −17,73 %, backend −11,337 s / −5,45 %, commande −11,361 s / −3,94 %.** CPU −0,72 %, allocations +0,33 %, pic RSS +2,63 % (+0,155 Gio). Le nouveau compteur transitif vaut **28,773 s**, incluses dans la préparation. Un seul nouveau run complet : les secondes exactes restent à consolider et les variations de chargement/émission ne sont pas attribuées au changement.

**2 655 TAST et 2 959 Go strictement identiques**, aucun ajout ni retrait ; 2 744 fichiers sources/configuration surveillés inchangés. Binaires testé/profilé/installé identiques. Tous les processus mesurés retournent 0, un seul `purs compile` applicatif, aucun appel de configuration JS. Sans option bundle, ce workflow génère le Go et résout ses dépendances, sans compiler le binaire applicatif final. Baselines officielles d’exécution altbak.pub distinctes ; aucun gain d’exécution revendiqué.

**Changement conservé ; point 4 reste ouvert.** L’optimisation/émission occupe encore 126,789 s sur 196,494 s ; les recherches locales PBO restent une piste de sonde au point 8. La transmission des directives entre modules doit être préservée avant toute parallélisation de l’optimisation PBO. L’incrémental entre builds (10) reste distinct du rebuild forcé.

[Sondes et tests](/Users/0x1/Documents/htdocs/scratch/pbo-build-20260921/rapport.md) · [Profil complet et limites](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-parallel-20260921/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-parallel-20260921/comparison.json) · [Validation](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-parallel-20260921/validation.json) · [Configuration](/Users/0x1/Documents/htdocs/gopurs/gopurs/docs/parallel-preparation.md).

## Point 8 — cache des recherches externes PBO, intégré et mesuré

La sonde native sur 304 modules compte **119 411 recherches**, dont **94 658 externes** : **88 180 répétitions (93,16 %)**, maximum 145 noms externes distincts dans un module. Le builder crée maintenant un cache FIFO de **512 paires `(module, identifiant)` par conversion**, absences comprises. La recherche locale reste prioritaire après chaque binding ; le cache ne fige donc pas un résultat négatif quand une définition apparaît. Purmeta reste stable pendant la conversion synchrone et est publié ensuite. Le bridge conserve des clés String natives, comparées par contenu. L’ancienne API `toBackendModule` reste sans cache ; aucun état global supplémentaire.

**31 tests JS et 10 tests Go sous `-race` passent** : équivalence avec la conversion originale, miss forward puis définition locale, priorité locale, publication/remplacement/reset, clés FFI fraîchement boxées, FIFO, indépendance et concurrence. Builds JS/natif réussis. Deux chauffes puis référence/candidat/candidat/référence : **backend 11,011 → 10,824 s (−1,70 %)**, optimisation/émission **6,086 → 5,891 s (−3,21 %)**, CPU **−1,49 %**, 304 TAST / 399 Go identiques dans les six passages.

| Vrai `b -c -n` | Avant | Après |
|---|---:|---:|
| Chargement TAST | 36,362 s | 33,482 s |
| Préparation/monomorphisation | 47,389 s | 45,706 s |
| Optimisation/émission | 135,805 s | **128,640 s** |
| Backend | 219,560 s | **207,831 s** |
| Commande entière | 295,235 s | **288,509 s** |
| CPU utilisateur + système | 982,126 s | **953,616 s** |
| Allocations estimées | 382,36 Gio | **374,42 Gio** |
| Pic RSS | 5,957 Gio | **5,911 Gio** |

**−11,729 s de backend / −5,34 %, −7,94 Gio d’allocations / −2,08 %, CPU −2,90 %.** `lookupImplementation` baisse de **22,42 à 14,09 Gio** et de **10,20 à 5,89 s CPU échantillonnées**. Un seul nouveau run complet ; les **4,563 s** gagnées en chargement/préparation et les variations des étapes externes ne sont pas attribuées à ce cache. Le petit corpus confirme un effet plus modeste. Les coûts pprof imbriqués ne s’additionnent pas.

**2 655 TAST et 2 959 Go strictement identiques**, aucun ajout ni retrait ; 2 740 fichiers sources/configuration surveillés inchangés. Binaire profilé, candidat testé et binaire installé identiques. Tous les processus mesurés retournent 0, un seul `purs compile` applicatif, aucun appel de configuration JS. L’option bundle étant absente, la commande ne compile pas le binaire applicatif final. Baselines officielles d’exécution altbak.pub distinctes de ces temps de compilation.

**Changement conservé ; point 8 reste ouvert.** Prochaine piste : construction d’expressions PBO et recherches locales encore coûteuses, à mesurer avant nouvelle réutilisation. Le parallélisme supplémentaire (4) et l’incrémental entre builds (10) restent des chantiers séparés.

[Sonde et tests](/Users/0x1/Documents/htdocs/scratch/pbo-lookup-20260921/rapport.md) · [Profil b8x](/Users/0x1/Documents/htdocs/scratch/b8x-lookup-cache-20260921/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-lookup-cache-20260921/comparison.json) · [Validation](/Users/0x1/Documents/htdocs/scratch/b8x-lookup-cache-20260921/validation.json).

## Point 8 — boxage des pointeurs ADT structuré, intégré et mesuré

La sonde des 22 conversions `Str(chars[i]).StrVal()` du scanner exact passe 280 cas mais ne réduit aucune allocation, avec moins de 1 % de variation de temps : piste écartée. La mesure suivante sur 304 modules attribue **39,97 % des octets rescannés** au boxage des pointeurs ADT. `GoBoxStructPointer` conserve désormais tag et opérande jusqu’au rendu, sans réimpression/scan intermédiaire. Quatre fichiers de production, mêmes octets Go finaux.

**22 tests passent**, dont tag, identité, nil, évaluation unique et imports imbriqués. Deux chauffes puis référence/candidat/candidat/référence sur le petit corpus : optimisation/émission **−4,73 %**, backend **−2,55 %**, CPU **−3,07 %** ; 304 TAST / 399 Go identiques dans les six passages.

| Vrai `b -c -n` | Avant | Après |
|---|---:|---:|
| Chargement TAST | 34,866 s | 36,362 s |
| Préparation | 47,974 s | 47,389 s |
| Optimisation / émission | 146,288 s | **135,805 s** |
| Backend | 229,134 s | **219,560 s** |
| Commande entière | 301,760 s | **295,235 s** |
| CPU utilisateur + système | 1 021,770 s | 982,126 s |
| Allocations cumulées | 394,78 Gio | **382,36 Gio** |
| Pic RSS | 6,022 Gio | 5,957 Gio |

**Backend −9,574 s / −4,18 %, commande −6,525 s / −2,16 %, allocations −12,42 Gio / −3,15 %.** Le scanner passe de **54,77 à 31,62 Gio (−42,26 %)** et de 18,85 à 11,91 s CPU cumulées. La collecte des imports structurés augmente de **4,18 à 10,36 Gio** : le gain local n’est pas intégralement un gain global. Une seule nouvelle mesure b8x complète ; la référence antérieure plus rapide donnait 223,904 s, soit −4,344 s par rapport à elle. Le petit benchmark alterné et la baisse du CPU/allocations confortent un gain modeste, sans promettre un nombre fixe de secondes.

**2 655 TAST et 2 959 Go identiques**, toutes les FFI incluses ; 2 708 fichiers surveillés inchangés ; binaire testé/reconstruit/installé identique. Aucun échec masqué, ni changement de concurrence ou de GC. Changement conservé. Restent les coûts PBO ci-dessus et les autres fragments opaques, avec une nouvelle sonde avant chaque extension.

[Sondes et tests](/Users/0x1/Documents/htdocs/scratch/gopurs-scanner-next-20260921/rapport.md) · [Profil b8x et limites](/Users/0x1/Documents/htdocs/scratch/b8x-structured-boxing-20260921/rapport.md) · [Comparaison](/Users/0x1/Documents/htdocs/scratch/b8x-structured-boxing-20260921/comparison.json) · [Validation](/Users/0x1/Documents/htdocs/scratch/b8x-structured-boxing-20260921/validation.json).

## Points 6/7/8 — nouveau profil de contrôle après correction du build

Même compilateur natif, **2 655 TAST / 2 959 Go identiques**, 2 707 fichiers surveillés inchangés. Le vrai `b -c -n` retourne 0 ; tous les processus mesurés réussissent, sans second `purs compile` ni appel de configuration Go. Aucun changement d’optimisation pendant ce passage.

| Mesure | Profil précédent | Répétition |
|---|---:|---:|
| Chargement TAST | 33,826 s | 34,866 s |
| Préparation / monomorphisation | 45,096 s | 47,974 s |
| Optimisation / émission | 144,973 s | **146,288 s** |
| Backend | 223,904 s | **229,134 s** |
| Commande entière | 299,516 s | **301,760 s** |
| CPU utilisateur + système | 1 011,879 s | 1 021,770 s |
| Allocations cumulées | 395,64 Gio | **394,78 Gio** |
| Pic RSS | 6,027 Gio | 6,022 Gio |

Backend +2,34 %, CPU +0,98 %, allocations −0,22 % : cette répétition au même binaire ne démontre pas une régression. Le GC représente 68,66 % du CPU échantillonné dans l’union des chemins suivis, dont 44,00 points de workers idle ; **49,223 ms seulement de pauses globales sur 679 cycles**. Les pourcentages CPU ne sont pas des gains muraux promis. Paramètres inchangés : 14 P, chargement 1, émission 2, pipeline actif.

**Cibles confirmées :** scanner `referencedImports` **54,77 Gio / 18,85 s CPU**, PBO `Semantics.build` **48,98 Gio / 18,58 s CPU**, `makeExternEvalSpine` **43,52 Gio / 17,24 s CPU**. Coûts cumulés et imbriqués, à ne pas additionner. Sonde suivante courte : supprimer dans une copie du scanner les conversions générées `Str(chars[i]).StrVal()`, vérifier les cas et mesurer. Si le gain est faible, passer à la suppression de scans via l’AST Go structuré ou aux recherches d’implémentations PBO. Pas de gain chiffré avant mesure.

[Rapport complet et limites](/Users/0x1/Documents/htdocs/scratch/b8x-clean-profile-20260921/rapport.md) · [Comparaison](/Users/0x1/Documents/htdocs/scratch/b8x-clean-profile-20260921/comparison.json) · [Validation](/Users/0x1/Documents/htdocs/scratch/b8x-clean-profile-20260921/validation.json). Ce build émet le Go et résout ses dépendances ; il ne compile pas le binaire applicatif final sans option bundle. Les baselines d’exécution d’altbak.pub restent distinctes.

## Point 5 — JS/PHP également corrigés

La lecture du code et une sonde Spago 1.0.3 confirment que le build JS génère déjà ses sourcemaps. Le passage supplémentaire sur `output/*/*.purs` est supprimé. L’échec immédiat de `conf` provenait du daemon Docker absent. PHP lit lui aussi ses variables directement via `getenv` : seule la cible JS appelle encore la génération.

`bin/conf` utilise un petit outil Node qui charge `env/_shared`, `env/<ENV>` puis l’environnement du processus. Il interpole les templates JS sans lancer le programme compilé ni Docker, encode les valeurs comme chaînes JS et conserve le commit Git. Les suffixes minuscules des noms sont reconnus ; les fallbacks explicites des templates s’appliquent aux variables absentes, notamment POSTGRES_IDLE_TIMEOUT_MS. Les erreurs de génération restent visibles et empêchent un statut de succès.

**12 tests passent**, syntaxe Bash/Node validée. Un mini-build réel de **54 modules** avec les vrais scripts, templates et environnement local de b8x produit des sourcemaps valides et deux configurations JS importables : **65 placeholders, aucun restant**. Les FFI PHP sont vérifiées avec un environnement factice. Les fichiers temporaires de configuration sont supprimés ; cible Go conservée, aucun rebuild complet b8x ni nouveau gain temporel revendiqué. Les chemins watch/bundle historiques ne font pas partie de cette correction.

[Compte-rendu et limites](/Users/0x1/Documents/htdocs/scratch/b8x-js-php-postbuild-20260921/rapport.md) · [Validation](/Users/0x1/Documents/htdocs/scratch/b8x-js-php-postbuild-20260921/validation.json). Les mentions d’échecs dans les profils historiques ci-dessous décrivent les runs antérieurs à cette correction.

## Point 5 — étapes finales corrigées pour Go

Après le profil ci-dessous, `b8x/bin/build` réserve les deux étapes finales historiques aux cibles autres que Go. Le glob des sourcemaps cherchait des `.purs` absents dans `output` ; le générateur `CodeGen Config` ne produit que les templates JS/PHP. Les FFI Go lisent déjà l’environnement au démarrage. L’export `GIT_COMMIT` de `conf`, effectué dans un processus enfant, n’alimentait pas le build parent ni les exécutions Go suivantes.

Validation sans compilation réelle : `bash -n`, puis 7 exécutions du vrai script dans des répertoires temporaires avec Spago/Go/npm/purs/conf simulés. Cas : Go `-c -n`, Go simple, Go watch, préservation des chemins JS et PHP, propagation d’un échec Spago et d’un échec `go mod tidy`. Tous passent ; aucun appel sourcemaps/conf pour Go. Aucun gain temporel supplémentaire n’est revendiqué. Le diagnostic et les corrections JS/PHP restent hors de cette sous-étape.

## Point 3 — contributions de collecte transitive mises en cache

Le cache des préfixes préparés réutilise désormais aussi les contributions de `monomorphizeExpr`/`collectExpr`. Il est local à une invocation et invalide un corps si ses types/arguments/substitutions changent, ou si une spécialisation d'un nom global dont il dépend apparaît. Pendant cette collecte les ensembles de clés ne font que croître : leurs cardinalités permettent une vérification conservatrice. Les noms des arguments statiques, dictionnaires locaux, bindings, cases et guards sont inclus.

Les contributions sont rejouées dans l'ordre existant en **conservant le premier payload et en unissant les callers**. Un test contre l'ancien algorithme a détecté puis validé la correction d'une inversion initiale : `Map.insertWith` transmet l'existant avant le nouveau. Cette première variante est écartée ; les chiffres finaux ci-dessous concernent la correction.

**59 tests ciblés passent**, dont 9 cas de cache également comparés à l'ancien `transitiveCollect`. Builds JS et natif sans erreur ni avertissement. Petit corpus : deux chauffes puis référence/candidat/candidat/référence ; préparation **1,497 → 1,309 s (−12,59 %)**, backend **11,638 → 11,678 s (+0,34 %)**, CPU +0,79 %. Aucun gain total établi sur ce petit corpus. **304 TAST et 399 Go identiques dans les six runs**.

| Vrai `b -c -n` b8x | Avant | Après |
|---|---:|---:|
| Chargement TAST | 34,676 s | 33,826 s |
| **Préparation/monomorphisation** | **82,013 s** | **45,096 s** |
| Optimisation/émission | 151,315 s | 144,973 s |
| **Backend** | **268,011 s** | **223,904 s** |
| **Commande entière** | **349,175 s** | **299,516 s** |
| CPU utilisateur + système | 1 175,454 s | 1 011,879 s |
| Allocations estimées | 458,57 Gio | **395,64 Gio** |
| Pic RSS | 6,190 Gio | 6,027 Gio |

**−36,917 s / −45,01 % de préparation ; −44,107 s / −16,46 % de backend ; −62,93 Gio / −13,72 % d'allocations.** CPU −13,92 %, objets alloués −1,445 milliard / −12,91 %, pic RSS −2,63 %. Le binaire testé, celui reconstruit par `b -c -n` et celui installé sont identiques. Sources surveillées inchangées ; **2 655 TAST et 2 959 Go strictement identiques**, aucun ajout ni retrait.

Un seul nouveau run complet : la réduction cohérente de la préparation, du CPU et des allocations confirme l'intérêt du changement, mais les secondes exactes restent à consolider. Les 6,342 s de baisse de l'optimisation/émission et les autres variations de la commande ne sont pas toutes attribuées au cache. Le pic RSS global ne mesure pas isolément la rétention de ses entrées.

Les sourcemaps retournent toujours 1 ; le hook corrigé capture cette fois **`conf` à 1 également**, toujours masqué par le script. Le code global 0 ne valide donc pas ces deux étapes. La commande n'a pas compilé le binaire final applicatif (option bundle absente).

**Changement conservé.** Point 3 reste ouvert pour les autres coûts de préparation, mais les parcours répétés visés ici sont traités. L'optimisation/émission représente maintenant **145 s sur 224 s de backend** : priorité suivante, identifier les autres parcours et allocations coûteux de PBO et de l'émission (point 8), puis envisager le parallélisme supplémentaire (point 4) sur des tâches indépendantes. Le point 10 reste distinct : un cache nettoyé par `-c` n'accélère pas ce rebuild forcé.

[Implémentation et sondes](/Users/0x1/Documents/htdocs/scratch/gopurs-transitive-20260921/rapport.md) · [Profil b8x](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-after-20260921/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-after-20260921/comparison.json) · [Identité des sorties](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-after-20260921/file-comparison.json).

## Point 8 — cache des instanciations intégré et validé

La sonde native sur 304 modules attribue **38,89 % du temps des instanciations** aux paires expression/type déjà rencontrées dans le module. Un cache réel réutilise aussi les résultats intermédiaires des applications à plusieurs types : le prototype atteint **55 795 hits / 67 011 appels** sur ce corpus, puis **720 602 / 784 523** sur b8x. Ces compteurs appartiennent au prototype, pas à une instrumentation permanente.

Le builder crée par Effect un cache FIFO de **512 entrées maximum par module**, alloué paresseusement. `ConvertEnv` et `Env` transmettent uniquement la fonction d’instanciation. Les clés conservent les références complètes ; les arbres restent immuables. La résolution des implémentations et `InlineNever` précèdent le cache. Les résultats publiés ne contiennent ni son état ni son Env. Le cache Go est protégé par mutex et calcule hors verrou. La borne porte sur les entrées, pas sur un nombre absolu d’octets.

**63 tests JS et 9 tests Go sous `-race` passent** : FIFO, identité, GC, indépendance des Effects, réentrance/concurrence, quantificateurs et capture comparés au calcul original, changement de directive après un hit et remplacement d’un corps portant le même nom. Builds JS/natif sans erreur ni avertissement. Le retour FFI curryfié utilise le bridge existant ; aucun changement du runtime ni du générateur Go dans cette étape.

Comparaison native finale, avec les FFI réelles : deux chauffes puis référence/candidat/candidat/référence. **304 TAST et 399 Go identiques dans six runs**. Optimisation/émission **7,405 → 6,580 s (−11,14 %)**, backend **12,762 → 12,009 s (−5,90 %)**, CPU −6,47 %. Pic RSS moyen +3,17 %. Les premiers scratchs omettaient le lien `../gopurs` des FFI ; leurs durées ne servent pas de référence de compilation complète.

| Vrai `b -c -n` | Avant cache | Après cache |
|---|---:|---:|
| Chargement TAST | 34,616 s | 34,676 s |
| Préparation/monomorphisation | 83,987 s | 82,013 s |
| Optimisation/émission | 200,769 s | **151,315 s** |
| **Backend** | **319,381 s** | **268,011 s** |
| **Commande entière** | **415,464 s** | **349,175 s** |
| CPU utilisateur + système | 1 393,516 s | 1 175,454 s |
| Allocations estimées | 541,68 Gio | **458,57 Gio** |
| Pic RSS | 6,444 Gio | 6,190 Gio |

**−51,370 s de backend / −16,08 %, −83,11 Gio d’allocations / −15,34 %, CPU −15,65 %.** Optimisation/émission −24,63 %. La baisse cohérente du CPU et des allocations conforte le gain mural. Les 66,289 s de baisse de la commande entière incluent aussi les variations des autres étapes. Un seul nouveau run complet : secondes exactes à consolider. Le pic RSS global baisse de 3,93 % ; le poids exclusif des objets retenus par le cache n’est pas isolé.

**2 655 TAST et 2 959 Go identiques**, aucun ajout ni retrait, FFI réelles incluses. Le binaire testé, celui reconstruit par `b -c -n` et celui installé sont identiques. Les sourcemaps retournent toujours 1. Le statut de `conf`, enveloppé dans `|| true`, n’a pas été capturé dans ce passage : le hook visait la ligne 123, l’appel étant ligne 124. Le point 5 reste ouvert.

**Cache conservé ; point 8 encore ouvert.** Prochaine étape importante : mesurer les recalculs `monomorphizeExpr`/`collectExpr` du point 3 et établir leurs invalidations avant réutilisation. Les anciens profils scratch ont été nettoyés avant cette tâche ; la référence historique ci-dessus provient du relevé conservé dans ce todo. Les baselines d’exécution officielles d’altbak.pub restent distinctes de ces temps de compilation.

[Implémentation et sondes](/Users/0x1/Documents/htdocs/scratch/gopurs-neutral-cache-20260921/rapport.md) · [Profil b8x et limites](/Users/0x1/Documents/htdocs/scratch/b8x-neutral-cache-after-20260921/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-neutral-cache-after-20260921/comparison.json) · [Sorties identiques](/Users/0x1/Documents/htdocs/scratch/b8x-neutral-cache-after-20260921/file-comparison.json).

## Point 8 — capture des parcours PBO et Array.any générique

La sonde JS sur les **2 655 TAST b8x** observe **734 669 substitutions racines / 44 779 920 visites de nœuds**. Avec le critère conservateur « aucune référence d’annotation remplacée », seules 1 234 racines / 53 643 visites restent inchangées (**0,12 % des visites**). Une garde à l’entrée des expressions complètes est écartée. Les sous-arbres inchangés sont plus nombreux (87,93 %), mais leur détection a elle-même un coût. Les 223 641 répétitions d’une paire expression/argument par identité sont une piste à mesurer, sans cache encore ajouté.

Le profil révélait **11,94 Gio d’allocations directes** dans la passerelle de `Array.any` : copie du tableau puis emballage de tous ses éléments en interfaces, avant le premier test du prédicat. `AnyImpl[A any]` supprime ces conversions. `FfiBridge` résout désormais récursivement les paramètres génériques dans les callbacks, tableaux et retours, en cohérence avec l’instanciation Go. La bibliothèque n’expose pas la représentation du runtime.

**512 substitutions réelles validées** : −17,27 % d’octets et −19,60 % d’objets sur ce microbenchmark. **46 tests ciblés passent**, dont ordre, court-circuit, tableau vide, immutabilité, callbacks/retours génériques et portée TypeApp. Builds JS/natif sans avertissement ni erreur. Petit corpus natif alterné : backend **12,984 → 12,933 s**, phase optimisation/émission **7,496 → 7,515 s**, soit **pas de gain mural établi** ; 304 TAST et 399 Go identiques dans six runs.

| Vrai `b -c -n`, même instrumentation | Avant | Après |
|---|---:|---:|
| Chargement/tri TAST | 35,275 s | 34,616 s |
| Préparation/monomorphisation | 91,296 s | 83,987 s |
| Optimisation/émission | 216,283 s | 200,769 s |
| **Backend interne** | **342,863 s** | **319,381 s** |
| **Commande entière** | **434,000 s** | **415,464 s** |
| CPU utilisateur + système | 1 424,706 s | 1 393,516 s |
| Allocations cumulées estimées | 554,31 Gio | **541,68 Gio** |
| Pic RSS | 6,290 Gio | 6,444 Gio |

**−12,63 Gio / −2,28 % d’allocations**, −489,6 millions d’objets / −3,44 %. Plus aucune allocation directe échantillonnée dans la passerelle AnyImpl. **2 655 TAST inchangés ; 2 958 Go identiques et uniquement Data_Array_ffi.go modifié**, exactement pour la signature et son bridge. Le binaire du profil est identique au candidat testé et à celui installé. Sources b8x inchangées ; seul `.DS_Store`, métadonnée Finder, change dans le relevé large des fichiers.

**Correction conservée pour la réduction démontrée des allocations. Les 23,482 s de gain mural restent indicatives** : un seul nouveau run complet, petit corpus presque stable, CPU réel −2,19 % et charge variable. Le CPU échantillonné pprof augmente de 1 075,96 à 1 098,46 s et n’est pas interchangeable avec le CPU système. Le pic RSS augmente de 2,44 %. Sourcemaps/conf échouent toujours avec le code 1, masqué par le script ; ils restent au point 5. Les baselines d’exécution du README altbak.pub restent distinctes de ces durées de compilation.

[Rapport, capture et limites](/Users/0x1/Documents/htdocs/scratch/gopurs-neutral-probe-20260921/rapport.md) · [Microbenchmark FFI](/Users/0x1/Documents/htdocs/scratch/gopurs-any-probe-20260921/rapport.md) · [Profil b8x](/Users/0x1/Documents/htdocs/scratch/b8x-any-after-20260921/comparison.json) · [Comparaison des sorties](/Users/0x1/Documents/htdocs/scratch/b8x-any-after-20260921/file-comparison.json).

## Point 9 — index intégré et validé sur b8x le 21 septembre 2026

`Gopurs.ReboxMetadata` indexe les alias `Constructor_…` et `Data_…` à partir des tables originales des constructeurs/classes. Une insertion n’écrase jamais une clé existante : même premier résultat que les anciens scans croissants, et constructeurs prioritaires. `Main.loadAndPrepareModules` construit l’index une seule fois ; les mises à jour de `globalFunctions` le conservent. Aucun cache mutable global ni nouvelle FFI.

La sonde initiale vérifie **4 468 recherches** sur les métadonnées réelles b8x. **90 tests ciblés passent**, dont 10 sur les collisions, alias, ordre des champs, absences et indépendance des métadonnées. Builds JS/natif sans erreur ni avertissement. Sur le petit corpus, deux chauffes puis référence/candidat/candidat/référence : optimisation/émission **8,464 → 7,844 s (−7,32 %)**, backend **14,125 → 13,602 s (−3,71 %)** ; **399 Go et 304 TAST identiques dans les six runs**.

| Vrai `b -c -n`, même instrumentation | Avant index | Après index |
|---|---:|---:|
| Chargement/tri TAST | 34,061 s | 35,275 s |
| Préparation/monomorphisation | 83,683 s | 91,296 s |
| Optimisation/émission | 283,405 s | **216,283 s** |
| **Backend interne** | **401,157 s** | **342,863 s** |
| **Commande entière** | **486,084 s** | **434,000 s** |
| CPU utilisateur + système | 1 806,211 s | 1 424,706 s |
| Allocations cumulées estimées | 704,04 Gio | **554,31 Gio** |
| Pic RSS | 6,215 Gio | 6,290 Gio |

**2 655 TAST et 2 959 Go identiques**, mêmes sources pendant le run, binaire reconstruit identique au candidat du petit benchmark. Les allocations du chemin ciblé passent de **150,56 Gio à 69 Mio**, et le CPU total baisse de **21,12 %**. L’index lui-même représente environ 24 Mio d’allocations échantillonnées ; son coût est inclus dans les totaux.

Le backend gagne **58,294 s** sur cette paire, mais la référence plus rapide du 20 septembre donne **31,392 s / −8,39 %** de gain. Le frontend b8x et la préparation augmentent pendant ce run ; la charge concurrente empêche une attribution exacte de toutes les secondes murales. Un seul nouveau run complet ; pas de gain extrapolé à une compilation sans profilage. Le pic mémoire augmente de **1,21 %**. Les échecs masqués sourcemaps/configuration persistent.

**Index conservé et compilateur natif installé.** Prochaine priorité : point 8 (passes internes PBO), puis point 3 (collecte transitive répétée), selon les prochaines captures.

[Rapport et limites](/Users/0x1/Documents/htdocs/scratch/gopurs-rebox-index-20260921/rapport.md) · [Benchmark natif alterné](/Users/0x1/Documents/htdocs/scratch/gopurs-rebox-index-20260921/native-benchmark.json) · [Mesures b8x](/Users/0x1/Documents/htdocs/scratch/b8x-rebox-after-20260921/comparison.json) · [Sorties identiques](/Users/0x1/Documents/htdocs/scratch/b8x-rebox-after-20260921/file-comparison.json).

## Points 6/7 — profil complet du 21 septembre 2026, avant index

Même binaire exact que le 20 septembre, **2 655 TAST et 2 959 Go identiques** à la référence. Les 1 957 fichiers source suivis par le manifeste n’ont pas changé pendant le run. Aucune optimisation de code appliquée.

| Mesure | 20 septembre | 21 septembre |
|---|---:|---:|
| Chargement/tri TAST | 33,731 s | 34,061 s |
| Préparation/monomorphisation | 81,535 s | 83,683 s |
| Optimisation/émission | 258,985 s | 283,405 s |
| **Backend interne** | **374,255 s** | **401,157 s** |
| **Commande entière** | **458,578 s** | **486,084 s** |
| CPU utilisateur + système | 1 801,035 s | 1 806,211 s |
| Allocations cumulées estimées | 702,92 Gio | 704,04 Gio |
| Pic RSS | 6,358 Gio | 6,215 Gio |

Le temps mural varie de +7,19 %, alors que CPU (+0,29 %), allocations (+0,16 %) et travail produit restent stables. Des services macOS de gestion du stockage étaient actifs vers la fin ; leur contribution à l’écart n’est pas quantifiée. **Cette répétition ne démontre pas une régression**, ni des conditions système parfaitement comparables. Les échecs masqués sourcemaps/configuration restent présents.

Les cibles prioritaires sont désormais l’index de `findReboxFields` (point 9), les passes internes PBO (point 8) et les parcours encore répétés de la collecte transitive (point 3). Le scanner reste secondaire : `referencedImports` représente encore 55,11 Gio et 19,23 s CPU hors GC ; propager les imports pourrait éviter des scans. Tous ces cumuls incluent leurs descendants et peuvent se recouvrir ; ils ne sont pas des gains muraux promis.

L’union GC représente **69,81 % du CPU échantillonné**, dont 43,27 points de workers idle. Les pauses globales ne totalisent que **83,39 ms** sur 921 cycles. Les coûts du profilage lui-même et les limites d’attribution des piles restent présents. Régler le GC ou multiplier les goroutines ne fournit toujours pas de gain garanti.

Le JS est déjà désactivé pour b8x cible Go. Le passage TAST du bootstrap gopurs émet encore du JS ; sa suppression ne pourrait économiser qu’une fraction de cette étape de 9,7 s. C’est une piste secondaire, distincte des gros chantiers ci-dessus.

[Rapport et limites](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260921/rapport.md) · [Mesures comparées](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260921/comparison.json) · [Profils CPU/mémoire et filtres](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260921/profile-totals.json) · [Identité des sorties](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260921/previous-output-comparison.json).

## Point 1 — microbenchmark terminé le 19 septembre 2026

Copies exactes du runtime et des wrappers du compilateur profilé ; seule la signature des paramètres et du retour `Ordering` change de `interface{}` vers `gopurs_runtime.Value`. `Func5`, `Apply5`, `Unbox` et `Box` sont conservés. Go 1.27.0, Apple M4 Pro, un P, cinq répétitions par cas ; médianes :

| Cas | Temps actuel → typé | Allocations/appel | Octets/appel |
|---|---:|---:|---:|
| Char, valeurs String préparées | 46,52 → 8,564 ns | 3 → 0 | 72 → 0 |
| String, valeurs String préparées | 53,98 → 9,488 ns | 3 → 0 | 72 → 0 |
| Char, `Str` à chaque appel | 64,67 → 27,15 ns | 5 → 2 | 104 → 32 |
| String, `Str` à chaque appel | 63,71 → 29,40 ns | 5 → 2 | 104 → 32 |

Les 28 paires testées passent pour les deux variantes : LT/EQ/GT, ASCII, Unicode, chaînes vides, préfixes et chaînes longues. L’analyse d’échappement confirme la fuite des trois arguments LT/EQ/GT au point d’appel FFI, via les interfaces et le retour vers `Box[interface{}]`. Cela explique les allocations de 24 octets vues dans le profil.

L’appel avec création des valeurs String est ×2,17 à ×2,38 plus rapide dans cette expérience. **Ce n’est pas un gain mesuré sur la compilation b8x.** Le profil b8x attribuait environ 171,38 Gio aux deux wrappers ; leur suppression effective dans le compilateur complet reste à vérifier.

La variante explicitement typée a servi de preuve expérimentale. L’intégration utilise ensuite des fonctions FFI Go génériques, comme décrit ci-dessous.

[Rapport du microbenchmark](/Users/0x1/Documents/htdocs/scratch/gopurs-ord-bench-20260919/rapport.md) · [Mesures brutes](/Users/0x1/Documents/htdocs/scratch/gopurs-ord-bench-20260919/bench.txt) · [Code et vérifications](/Users/0x1/Documents/htdocs/scratch/gopurs-ord-bench-20260919/ord_test.go) · [Analyse d’échappement](/Users/0x1/Documents/htdocs/scratch/gopurs-ord-bench-20260919/escape.log).

## Point 1 — intégration validée sur b8x

`OrdCharImpl[T any]` et `OrdStringImpl[T any]` conservent LT/EQ/GT dans le type `T`, avec les mêmes corps de comparaison. Le bridge instancie déjà les fonctions génériques avec `gopurs_runtime.Value` ; il transmet maintenant directement les arguments portant ces paramètres génériques. La FFI ne dépend pas du runtime gopurs. Les changements de **gopurs et gopurs-prelude doivent être livrés ensemble**.

| Mesure | Référence | Après correction | Variation |
|---|---:|---:|---:|
| Chargement/tri TAST | 34,283 s | 34,278 s | stable |
| Préparation/monomorphisation | 131,432 s | 117,134 s | −10,88 % |
| Optimisation/émission | 425,946 s | 341,801 s | −19,75 % |
| **Backend interne** | **591,665 s** | **493,217 s** | **−16,64 %** |
| **`b -c -n` entier** | **676,083 s** | **596,490 s** | **−11,77 %** |
| CPU système du backend | 2 642,935 s | 2 199,249 s | −16,79 % |
| Allocations cumulées | 1 030,31 Gio | 858,66 Gio | −16,66 % |
| Objets alloués, estimés | 32,467 milliards | 24,837 milliards | −23,50 % |
| Pic de RSS | 6,599 Gio | 6,639 Gio | environ stable |

Le backend gagne **1 min 38 s** et la commande entière **1 min 20 s**, malgré environ 20 s supplémentaires pour reconstruire le compilateur. Le temps CPU du GC diminue également d’environ 17 %, avec une part du CPU presque inchangée autour de 72 %.

Validation : build sans erreur/avertissement ; 16 tests FFI ; fixture Comparisons dynamique avec générateurs JS et natif (même snapshot) ; 2 tests du scanner incluant le vrai natif. Les **2 655 TAST sont identiques par SHA-256** avant/après ; parmi les **2 959 Go**, seul `Data_Ord_ffi.go` change. Les autres points n’ont pas été implémentés.

Limites : une référence du matin et un run après correction du soir, même instrumentation et mêmes paramètres, pas une série alternée répétée. Les sourcemaps et `conf` échouent toujours silencieusement : leur correction reste au point 5. Le gain mural d’une configuration réussie n’est pas établi.

[Rapport d’intégration](/Users/0x1/Documents/htdocs/scratch/b8x-ord-after-20260919/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-ord-after-20260919/summary.json) · [Comparaison des fichiers](/Users/0x1/Documents/htdocs/scratch/b8x-ord-after-20260919/file-comparison.json) · [Profil CPU comparé](/Users/0x1/Documents/htdocs/scratch/b8x-ord-after-20260919/cpu-summary.md).

## Point 2 — scanner validé le 20 septembre 2026

`isIdentifierChar` convertit le caractère une fois avec `Data.Char.toCharCode`, puis compare des entiers avec des conditions explicites. Cela supprime les créations de chaînes pour les bornes et les closures booléennes du prédicat. Le scanner conserve sa reconnaissance des identifiants ASCII et de tout caractère supérieur à 127, ainsi que les chaînes, commentaires, imports triés et dédupliqués.

Microbenchmark natif sur 160 fragments réels de b8x, cinq répétitions : **4,332 → 1,423 ms (×3,04)**, −57,4 % d’octets alloués et −72,4 % d’allocations. Les 1 713 cas par variante passent en JS et Go, avec pile Go limitée à 1 Mio. Le test de production enrichi passe aussi sur le compilateur natif reconstruit.

| Mesure | Après point 1 | Après scanner | Variation |
|---|---:|---:|---:|
| Préparation/monomorphisation | 117,134 s | 118,488 s | +1,16 % |
| Optimisation/émission | 341,801 s | 297,364 s | −13,00 % |
| **Backend interne** | **493,217 s** | **449,638 s** | **−8,84 %** |
| **`b -c -n` entier** | **596,490 s** | **550,478 s** | **−7,71 %** |
| Allocations cumulées | 858,66 Gio | 789,39 Gio | −8,07 % |
| Objets alloués estimés | 24,837 milliards | 20,182 milliards | −18,74 % |
| Pic RSS | 6,639 Gio | 6,347 Gio | −4,41 % |

Le CPU cumulé échantillonné de `referencedImports` baisse de **59,80 à 20,60 s (−65,55 %)**. Sur cette paire de mesures, le backend gagne encore **43,579 s** et la commande entière **46,011 s**. Depuis la référence initiale : **2 min 22 s gagnées sur le backend (−24,00 %)**.

Les **2 655 TAST et les 2 959 fichiers Go sont tous identiques** avant/après. Même commande et instrumentation, mêmes paramètres, une paire de runs non alternée. Les sourcemaps et `conf` échouent toujours : point 5 inchangé. Aucun changement du parallélisme, du GC ou de la propagation des imports.

[Microbenchmark](/Users/0x1/Documents/htdocs/scratch/gopurs-scanner-bench-20260920/rapport.md) · [Rapport d’intégration](/Users/0x1/Documents/htdocs/scratch/b8x-scanner-after-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-scanner-after-20260920/comparison.json) · [Comparaison des fichiers](/Users/0x1/Documents/htdocs/scratch/b8x-scanner-after-20260920/file-comparison.json).

## Point 3 — première sous-étape : garde de substitution

Le profil et la capture ont distingué deux chemins : `TypeSubstitution.substitute`, qui allouait environ 101 Gio pendant l’optimisation via `Semantics`, et `Substitute.substituteExprType`, utilisé par la monomorphisation. L’estimation initiale de 20–40 s basée sur la phase préparation était donc trop directe ; le gain observé de cette sous-étape est **15,906 s sur le backend**.

Une capture JS sur les 2 655 TAST, sans émission Go, observe 23,2 millions d’appels et échantillonne 512 cas. La garde peut conserver l’arbre d’origine dans 442 cas : aucune variable n’est une clé de substitution, et aucun `ForAll` n’impose le parcours conservateur. Le renommage sous quantificateur, la substitution simultanée et les rangées ouvertes restent inchangés.

Le microbenchmark natif mesure −12 à −14 % de temps et −24 % d’octets sur ce corpus. Les variantes « types fermés uniquement » et « prédicat à appels directs » régressent et sont écartées. La garde retenue ralentit les cas qui doivent encore être substitués ; sa portée reste locale, sans cache global ni mutation.

| Mesure | Après scanner | Après garde | Variation |
|---|---:|---:|---:|
| Préparation/monomorphisation | 118,488 s | 116,209 s | −1,92 % |
| Optimisation/émission | 297,364 s | 283,705 s | −4,59 % |
| **Backend interne** | **449,638 s** | **433,732 s** | **−3,54 %** |
| **`b -c -n` entier** | **550,478 s** | **537,635 s** | **−2,33 %** |
| Allocations cumulées | 789,39 Gio | 774,71 Gio | −1,86 % |
| Objets alloués estimés | 20,182 milliards | 20,236 milliards | +0,27 % |
| Pic RSS | 6,347 Gio | 6,489 Gio | +2,24 % |

Validation : **30 tests de portée/substitution**, les **512 cas sur le compilateur natif reconstruit**, **2 655 TAST et 2 959 Go tous identiques** par SHA-256. Une seule paire de runs non alternée ; la petite variation de préparation n’est pas attribuée à cette garde. Aucune réduction du nombre d’allocations ou du pic mémoire n’est établie. Les échecs `sourcemaps/conf` restent au point 5.

[Capture et microbenchmark](/Users/0x1/Documents/htdocs/scratch/gopurs-substitution-bench-20260920/rapport.md) · [Rapport d’intégration](/Users/0x1/Documents/htdocs/scratch/b8x-substitution-after-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-substitution-after-20260920/comparison.json) · [Comparaison des fichiers](/Users/0x1/Documents/htdocs/scratch/b8x-substitution-after-20260920/file-comparison.json).

## Point 3 — deuxième sous-étape : clés de spécialisation après les gardes

La capture réelle sur les 2 655 TAST détecte **1 477 535 clés calculées puis rejetées**. `collectExpr` calcule désormais sa clé après les gardes sur les types ; `monomorphizeExpr` vérifie d’abord la présence de la fonction dans la table, puis calcule le hash seulement si la spécialisation existe. Aucun cache ajouté, mêmes clés et mêmes noms.

Les appels récursifs JS à `mangleType` passent de **79,66 à 45,51 millions (−42,87 %)**. Le cache par identité n’a pas été retenu : seulement 2,47 % de répétitions du même objet aux appels racines avant changement. Les noms répétés ne prouvent pas à eux seuls l’égalité des types.

| Mesure | Après garde de substitution | Après clés différées | Variation |
|---|---:|---:|---:|
| Préparation/monomorphisation | 116,209 s | 100,035 s | −13,92 % |
| Optimisation/émission | 283,705 s | 289,795 s | +2,15 % |
| **Backend interne** | **433,732 s** | **424,522 s** | **−2,12 %** |
| **`b -c -n` entier** | **537,635 s** | **528,775 s** | **−1,65 %** |
| Allocations cumulées | 774,71 Gio | 737,44 Gio | −4,81 % |
| Objets alloués estimés | 20,236 milliards | 19,641 milliards | −2,94 % |
| Pic RSS | 6,489 Gio | 6,403 Gio | −1,32 % |

**16,174 s gagnées dans la phase ciblée, 9,210 s au total du backend** : la variation des autres phases empêche d’attribuer tout le gain ciblé à la commande entière. Allocations cumulées de `mangleType` : environ 51,16 → 26,68 Gio. Une paire de runs non alternée ; mêmes TAST, paramètres et instrumentation. **21 tests passent ; 2 655 TAST et 2 959 Go tous identiques**. Échecs sourcemaps/conf inchangés.

Depuis le départ : **591,665 → 424,522 s sur le backend, −28,25 % (2 min 47 gagnées)**.

[Capture](/Users/0x1/Documents/htdocs/scratch/gopurs-mangle-bench-20260920/rapport.md) · [Rapport natif](/Users/0x1/Documents/htdocs/scratch/b8x-mangle-after-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-mangle-after-20260920/comparison.json) · [Comparaison des fichiers](/Users/0x1/Documents/htdocs/scratch/b8x-mangle-after-20260920/file-comparison.json).

## Point 3 — troisième sous-étape : garde avant unification, essai retiré

La capture observe **902 306 substitutions (68,10 %) et 1 193 583 paires d’unification (70,70 %) évitables** dans le chemin `collectExpr / ExprApp`. La garde existante a été avancée avant ces calculs en conservant la collecte des appels imbriqués via `acc2`.

Le vrai `b -c -n` ne confirme pas de gain :

| Mesure | Référence conservée | Essai retiré |
|---|---:|---:|
| Préparation/monomorphisation | 100,035 s | 102,456 s |
| Backend interne | 424,522 s | 433,681 s |
| Commande entière | 528,775 s | 542,073 s |
| Allocations cumulées | 737,44 Gio | 730,99 Gio (−0,87 %) |

**21 tests passent, 2 655 TAST et 2 959 Go identiques.** Le Go saute bien les calculs sans surcoût de génération identifié. Une seule paire non alternée : toutes les grandes phases ralentissent d’environ 2 %, donc le lien causal avec la modification n’est pas établi. Le nombre d’appels évités surestime leur importance si l’on ne considère pas leur coût. La faible baisse mémoire et l’absence de gain de temps démontré conduisent à **retirer uniquement cet essai** et rétablir les sources et exécutables précédents. Échecs sourcemaps/conf inchangés.

[Capture](/Users/0x1/Documents/htdocs/scratch/gopurs-early-guard-20260920/rapport.md) · [Rapport natif](/Users/0x1/Documents/htdocs/scratch/b8x-early-guard-after-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-early-guard-after-20260920/comparison.json).

À l’issue de cet essai, la référence conservée était **424,522 s**. Le point 3 n’est pas déclaré terminé ; la sous-étape suivante porte sur le point 4.

## Point 3 — quatrième sous-étape : cache du préfixe de collecte transitive, prototype prometteur

Après l’absence de gain de l’ordonnanceur à 91 paires, retour au chantier de préparation demandé par l’utilisateur. La capture des 2 655 TAST b8x relève **31 tours, 493 656 traitements pour 16 258 instanciations** : 477 398 répétitions avec les mêmes entrées. Le prototype conserve le corps après `applyStaticArgs → resolveGlobals → rewriteExpr`. Chaque `monomorphizeExpr` et `collectExpr` reste exécuté dans l’ordre, car la table de spécialisations évolue entre les tours.

**Preuve JS complète** : table transitive, 2 655 modules et métadonnées identiques sur quatre préparations alternées. Temps instrumentés : 40,886 → 35,219 s (−13,86 %), distincts d’un benchmark natif.

**Première paire native sur les vrais TAST b8x**, préparation seule, même binaire avec/sans cache :

| Mesure | Référence | Cache |
|---|---:|---:|
| Préparation/monomorphisation | 97,092 s | **78,106 s (−19,55 %)** |
| Collecte transitive | 88,153 s | 68,704 s |
| Reconstructions du préfixe | 20,565 s | 0,673 s |
| Chargement + préparation, processus | 129,910 s | 111,479 s |

**18,986 s gagnées dans la préparation**, 37,48 Gio d’allocations évitées sur chargement + préparation (−17,32 %), pic RSS +2,07 %. Le cache natif retrouve 16 258 calculs et 477 398 hits ; il vérifie l’identité des types/substitutions et de chaque argument, avec recalcul en cas de différence. Cache local, entrées immuables conservées, aucune hypothèse d’exclusivité du TAST.

**16 contrôles d’invalidation passent** ; deux parcours natifs complets du petit corpus produisent **399 Go et 304 TAST identiques**. Leur total reste stable (12,816 → 12,872 s), malgré une préparation plus courte : aucun gain global déduit de ce petit corpus. Les 2 655 contenus préparés b8x ont été comparés en JS ; la sonde native b8x ne vérifie que leur nombre et les entrées.

Le prototype reste en scratch, sources et compilateur installés inchangés. **Prochaine étape : intégration propre puis comparaison des sorties natives b8x et vrai `b -c -n`**, avec mesure de l’effet mémoire sur PBO et répétition avant de consolider le gain. Une seule paire native de préparation ne donne pas le nouveau total : la référence backend complète reste **388,011 s** ; la préparation historique était 96,746 s, cohérente avec le contrôle à 97,092 s.

[Preuve JS](/Users/0x1/Documents/htdocs/scratch/gopurs-transitive-probe-20260920/rapport.md) · [Rapport natif](/Users/0x1/Documents/htdocs/scratch/gopurs-transitive-native-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/gopurs-transitive-native-20260920/summary.json) · [Sorties natives](/Users/0x1/Documents/htdocs/scratch/gopurs-transitive-native-20260920/go-validation.json).

## Point 3 — cinquième sous-étape : cache intégré, vrai build b8x validé

Le cache du préfixe de `transitiveCollect` est intégré en PureScript, avec une garde d’identité privée en JS/Go. Il est local à chaque appel et vérifie types, substitutions, longueurs et éléments des arguments avant réutilisation. Les passes dépendant de la table courante (`monomorphizeExpr` puis `collectExpr`) continuent dans le même ordre. Aucun changement du parallélisme ni mutation des entrées.

| Mesure | Référence après pipeline | Après cache | Variation |
|---|---:|---:|---:|
| Préparation/monomorphisation | 96,746 s | 81,535 s | −15,72 % |
| Optimisation/émission | 256,177 s | 258,985 s | +1,10 % |
| **Backend interne** | **388,011 s** | **374,255 s** | **−3,55 %** |
| Allocations cumulées | 737,44 Gio | 702,92 Gio | −4,68 % |
| Pic RSS | 6,222 Gio | 6,358 Gio | +2,17 % |

**15,211 s gagnées dans la préparation, 13,756 s sur le backend et 34,52 Gio d’allocations évitées.** La commande entière passe de 490,771 à 458,578 s, mais la reconstruction du compilateur gagne aussi 18,885 s avec un cache Go réchauffé par la validation préalable : ne pas attribuer les 32,193 s totales au seul changement.

Validation : **50 tests JS et 13 cas de garde native passent** ; builds sans erreur/avertissement ; petit corpus natif avec 399 Go identiques, puis vrai `b -c -n` avec **2 655 TAST et 2 959 Go tous identiques**. Les sources b8x restent inchangées, et le binaire installé est exactement celui profilé. Les échecs sourcemaps/conf préexistants restent au point 5.

Une seule nouvelle mesure complète face à la référence historique, pas une série alternée répétée. Le cache est conservé : le gain ciblé est cohérent avec le prototype et la baisse des allocations. Le point 3 reste ouvert pour les autres recalculs ; aucun gain supplémentaire n’est présumé. Nouvelle référence complète : **374,255 s**, soit **−36,75 % depuis 591,665 s**.

[Rapport](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-after-20260920/rapport.md) · [Comparaison](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-after-20260920/comparison.json) · [Identité des sorties](/Users/0x1/Documents/htdocs/scratch/b8x-transitive-after-20260920/file-comparison.json).

## Point 4 — première sous-étape : optimisation et émission simultanées

L’optimisation PBO reste séquentielle : le builder transmet les directives exportées par un module au suivant. En revanche, elle peut avancer pendant que le lot précédent émet du Go dans ses goroutines Aff. Le pipeline borne les travaux retenus à un lot actif, un lot en attente et le module en cours d’optimisation. Les lots restent ordonnés, avec leurs barrières de dépendance et leurs snapshots de signatures. Le nettoyage attend naturellement le lot actif, y compris les écritures asynchrones non annulables.

Activé par défaut ; `GOPURS_PIPELINE=0` rétablit le fonctionnement précédent. `GOPURS_EMIT_JOBS` reste à 2 par défaut ; à 1, l’émission reste immédiate et séquentielle. Aucun changement PBO supplémentaire dans cette sous-étape.

| Mesure | Après clés différées | Avec pipeline | Variation |
|---|---:|---:|---:|
| Optimisation/émission | 289,795 s | 256,177 s | −11,60 % |
| **Backend interne** | **424,522 s** | **388,011 s** | **−8,60 %** |
| **`b -c -n` entier** | **528,775 s** | **490,771 s** | **−7,19 %** |
| CPU utilisateur + système | 1 872,318 s | 1 867,965 s | −0,23 % |
| Allocations cumulées | 737,44 Gio | 737,44 Gio | stable à l’arrondi |
| Pic RSS | 6,403 Gio | 6,222 Gio | −2,82 % |

**33,618 s gagnées dans la phase ciblée, 36,511 s sur le backend et 38,004 s sur la commande entière.** Le CPU et les allocations sont quasi inchangés : le résultat est cohérent avec un meilleur chevauchement du même travail. La préparation et le bootstrap ont aussi varié ; leurs gains ne sont pas attribués au pipeline. Une seule paire complète b8x, complétée par quatre runs natifs alternés sur 304 TAST altbak : −6,14 % dans la phase ciblée et −4,24 % sur le backend en moyenne. Ces temps de compilation sont distincts des baselines officielles d’exécution d’altbak.pub.

Validation : **2 655 TAST et 2 959 Go identiques** sur b8x ; **10 tests JS**, **3 tests natifs principaux sous Go -race**, corpus de **304 TAST sous -race sans course signalée et 399 Go identiques**. Après le profil complet, le nettoyage sur échec du producteur a été corrigé pour attendre les callbacks non annulables ; cette branche n’avait pas été exercée par le backend réussi. Le compilateur final est reconstruit, ses tests repassent et deux compilations avec/sans pipeline donnent les mêmes 399 Go. Le profil b8x n’a pas été répété pour cette branche d’échec ; le chemin réussi est inchangé.

Les échecs sourcemaps/conf restent inchangés au point 5. Le point 4 reste ouvert : optimiser plusieurs modules PBO simultanément nécessite d’abord un contrat explicite sur leurs directives et dépendances. Aucun gain supplémentaire n’est chiffré à ce stade.

[Expérience courte et tests](/Users/0x1/Documents/htdocs/scratch/gopurs-pipeline-20260920/rapport.md) · [Rapport b8x](/Users/0x1/Documents/htdocs/scratch/b8x-pipeline-after-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/b8x-pipeline-after-20260920/comparison.json) · [Comparaison des fichiers](/Users/0x1/Documents/htdocs/scratch/b8x-pipeline-after-20260920/file-comparison.json).

## Point 4 — deuxième sous-étape : 2, 4 ou 8 workers, réglage conservé

Neuf compilations natives du même corpus de 304 TAST, pipeline activé, trois passages par réglage dans l’ordre 2/4/8, 8/2/4, 4/8/2. Une chauffe exclue, répertoires frais, même binaire, entrées FFI figées, aucun profilage. Moyennes :

| Workers | Optimisation/émission | Écart vs 2 | Backend total |
|---|---:|---:|---:|
| **2** | **7,482 s** | référence | **12,671 s** |
| 4 | 7,613 s | +1,76 % | 12,751 s |
| 8 | 7,737 s | +3,41 % | 12,886 s |

**Aucun gain : le réglage reste à 2.** Les médianes confirment cette décision. Les **399 Go sont identiques à la référence historique dans les neuf runs**. Le premier essai de chauffe avait détecté neuf FFI modifiées depuis la mesure précédente ; il a été exclu, puis les sources historiques ont été figées dans scratch avant de recommencer. Aucun code de production modifié, aucun nouveau build b8x lancé. Référence b8x conservée : **388,011 s de backend**.

Ces chiffres ne permettent pas de désigner PBO comme unique goulot. La prochaine expérience sur le parallélisme peut mesurer les attentes entre producteur et émission ; paralléliser plusieurs modules PBO nécessite toujours de préserver les directives entre modules.

[Rapport et protocole](/Users/0x1/Documents/htdocs/scratch/gopurs-pipeline-workers-20260920/rapport.md) · [Neuf mesures](/Users/0x1/Documents/htdocs/scratch/gopurs-pipeline-workers-20260920/benchmark.json) · [Statistiques](/Users/0x1/Documents/htdocs/scratch/gopurs-pipeline-workers-20260920/summary.json).

## Point 4 — troisième sous-étape : mesurer les attentes du pipeline

Sondes limitées à une copie native dans scratch, sur les mêmes 304 TAST et FFI figées. Deux runs instrumentés encadrés par deux références, pipeline actif et 2 workers. Moyennes dans la phase optimisation/émission de **7,783 s** :

- **PBO synchrone : 6,554 s (84,2 %)**, dont 2,700 s simultanées avec l’émission.
- **Jonctions du producteur : 1,210 s (15,5 %)**, presque entièrement pendant les appels `enqueue` ; jonction finale de 1,56 ms seulement.
- Un lot d’émission est actif 3,914 s ; **aucun lot actif pendant 49,7 % de la phase**. Cela ne signifie pas que tous les cœurs sont inactifs.

Ces durées se recouvrent, elles ne s’additionnent pas. Les deux captures concordent ; 304 conversions PBO, 304 traductions, 175 lots et jusqu’à 2 traductions simultanées. **399 Go et 304 TAST identiques dans les quatre runs**. La phase instrumentée est 2,64 % plus lente en moyenne que les deux contrôles ; bruit et surcoût ne sont pas séparés. Aucun code de production modifié, aucun build b8x complet relancé et aucune extrapolation chiffrée à b8x.

Le travail séquentiel de PBO devient la cible prioritaire sur ce corpus. Prochaine micro-étape : identifier dans `toBackendModule` une portion coûteuse indépendante des directives/publications précédentes et vérifier cette indépendance avant de paralléliser. Aucun gain supplémentaire n’est encore établi.

[Diagnostic détaillé](/Users/0x1/Documents/htdocs/scratch/gopurs-pipeline-waits-20260920/rapport.md) · [Intervalles et synthèse](/Users/0x1/Documents/htdocs/scratch/gopurs-pipeline-waits-20260920/summary.json).

## Point 4 — quatrième sous-étape : coûts et indépendance dans PBO

Deux runs instrumentés encadrés par deux contrôles sur les 304 TAST figés. Sur **6,352 s de PBO**, l’optimisation des expressions prend **4,896 s (77,1 %)**, la conversion initiale **1,096 s (17,3 %)**. Les préparatifs clairement indépendants — invalidation des annotations d’usage, parsing des directives locales, métadonnées des constructeurs — ne totalisent que **0,036 s (0,57 % de PBO)**. Leur parallélisation seule ne justifie pas un chantier.

La conversion initiale n’est pas indépendante des publications : **12 640 lectures réelles de purmeta par run, dont 11 195 trouvent des implémentations publiées**. `buildM/getCtx/analyze` consulte déjà les implémentations pour l’analyse et les simplifications ; le mutex du cache ne remplace pas l’ordre de disponibilité. L’optimisation consomme ensuite les directives. Le simple déplacement de `toBackendExpr` avant le module précédent n’est donc pas retenu.

**399 Go et 304 TAST identiques dans les quatre runs**, mêmes compteurs sur les deux captures. Écart moyen de la phase instrumentée aux contrôles : +0,43 %, bruit et surcoût non séparés. Code de production et binaire installé inchangés ; aucune extrapolation à b8x et aucun nouveau build complet.

Prochaine expérience à potentiel : rendre explicites les dépendances aux implémentations/directives et tester deux modules indépendants en parallèle, en comparant résultats et publications à l’ordre actuel. Aucun gain n’est encore établi ; cette expérience doit précéder toute modification générale du builder.

[Rapport et dépendances](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-parts-20260920/rapport.md) · [Coûts et compteurs](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-parts-20260920/parts-summary.json).

## Point 4 — cinquième sous-étape : prototype de deux optimisations simultanées

**Preuve positive en natif sur `Data.List` et `Data.Map.Internal`**, adjacents aux indices 247/248 du corpus de 304 TAST. Les 247 publications précédentes sont figées ; les deux actions utilisent les directives disponibles avant `Data.List`. Leurs lectures de cache observées sont compatibles avec cet état commun, sans lecture de l’autre module. Après calcul, les résultats sont publiés dans l’ordre initial.

| Mode | Trois mesures hors chauffe | Moyenne |
|---|---|---:|
| Séquentiel | 392,112 / 391,573 / 396,143 ms | **393,276 ms** |
| Deux goroutines | 343,118 / 346,469 / 350,163 ms | **346,583 ms** |

**46,693 ms gagnées sur cette paire (−11,87 %, facteur 1,135).** Même processus, GC hors chrono, ordre S/P/P/S/S/P après une chauffe par mode. Les actions capturées sont déjà préparées ; le chrono exclut invalidation initiale, comparaison, publication et émission Go. Chaque tâche ralentit individuellement en concurrence, limitant le gain total ; la cause précise n’est pas isolée.

Comparaison des **12 champs de BackendModule**, incluant tous les champs des analyses publiées, la syntaxe, les directives et les implémentations : mêmes résultats que la référence dans tous les replays. Une mutation volontaire de `analysis.externs` est détectée. OptimizationSteps est vérifié vide, comme dans le build habituel. **Le même couple passe sous Go -race sans course signalée.** Les compilations séquentielles préparatoires donnent les mêmes 399 Go ; les résultats parallèles sont comparés comme IR, sans réémission Go.

Le prototype reste en scratch : compilateur installé et ordonnanceur de production inchangés. Le gain de la paire ne s’extrapole pas à b8x. Prochaine étape possible : plusieurs paires, puis détermination des dépendances avant exécution et contrat explicite sur les directives ; cet essai identifie les dépendances lors d’un parcours de référence séquentiel.

[Rapport et limites](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-pair-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-pair-20260920/summary.json) · [Validations](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-pair-20260920/proof.json).

## Point 4 — sixième sous-étape : six paires d’optimisations simultanées

Extension du prototype à **six paires adjacentes et disjointes**, retenues avant les replays parmi 212 candidates selon les durées observées du parcours séquentiel. Aucun échec remplacé : **six paires sur six passent**, soit douze modules. Même cache figé et mêmes directives avant chaque paire, jonction des deux goroutines puis publication dans l’ordre initial.

Trois répétitions par mode après chauffe, ordre S/P/P/S/S/P : la somme des moyennes passe de **1,091643 à 0,990776 s (−9,24 %, 100,867 ms gagnées)**. Les gains par paire vont de **3,28 à 17,72 %**. Ce total additionne six replays isolés ; il exclut notamment publication et émission. Les paires sont choisies pour leur coût, donc ces chiffres ne représentent pas un build complet.

**48 replays natifs valident l’égalité complète des résultats PBO et publications**, analyses incluses ; les mutations volontaires des analyses des douze modules sont détectées. **Les six mêmes paires passent sous Go -race**, avec 12 replays supplémentaires et aucune course signalée. Les deux compilations préparatoires conservent **399 Go et 304 TAST identiques**. Les résultats parallèles sont comparés comme IR, sans réémission Go.

Le résultat positif dépasse désormais la paire initiale. Le prototype reste en scratch, sans modification du compilateur installé ; la référence b8x demeure **388,011 s**. Prochaine étape possible : prototype d’ordonnancement sur le corpus complet, avec dépendances déterminées avant calcul et contrat explicite des directives, puis mesure incluant coordination, publication et émission. Les lectures observées après une compilation séquentielle ne suffisent pas comme ordonnanceur de production.

[Rapport et limites](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-pairs-20260920/rapport.md) · [Mesures des six paires](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-pairs-20260920/summary.json) · [Validations](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-pairs-20260920/proof.json).

## Point 4 — septième sous-étape : ordonnanceur sur tout le corpus, aucun gain démontré

Prototype natif en scratch sur les **304 modules**, avec émission réelle des résultats parallèles. Le plan est calculé **avant PBO**, à partir des imports, réexports et références qualifiées du TAST déjà monomorphisé, puis de leur fermeture transitive. Il retient **29 paires contiguës disjointes, soit 58 modules**, et conserve l’ordre des callbacks/publications du builder.

Le contrat des directives a été précisé : dans le parcours sans cache actuel, la sortie d’un module ne contient que ses propres directives et remplace celles transmises au suivant. A reçoit son environnement réel ; B reçoit des directives initiales vides lorsqu’il ne peut pas référencer A. Le prototype est limité aux callbacks actuels de gopurs, sans cache-hit ni semantics personnalisées.

Deux chauffes exclues, puis trois mesures par variante dans un ordre alterné, mêmes entrées figées et réglages que les expériences précédentes :

| Mesure moyenne | Référence | Prototype | Variation |
|---|---:|---:|---:|
| Optimisation + émission | 7,485 s | 7,541 s | +0,76 % |
| Backend complet | 12,616 s | 12,694 s | +0,62 % |

**Aucun bénéfice mesuré : cette version n’est pas intégrée.** La planification prend 35 à 49 ms et est incluse. La référence historique de 7,482 s sur ce corpus est cohérente avec le nouveau contrôle. Ces petits écarts ne démontrent pas une régression générale de la parallélisation PBO.

Validation : **304 comparaisons complètes des résultats PBO** contre leur calcul séquentiel au même point du parcours, analyses/directives incluses ; **2 906 noms de dépendances audités**, tous couverts par le graphe hors références au module courant. **Le même plan passe sous Go -race**, avec 304 comparaisons supplémentaires et aucune course signalée. Les **399 Go et 304 TAST sont identiques** dans les dix runs réussis, dont deux chauffes et deux validations instrumentées. Cette fois les résultats parallèles sont bien réémis en Go.

Indice concret pour la suite : **quatre des six paires précédemment validées sont interdites par la fermeture conservatrice**, une cinquième est séparée par la sélection gloutonne et une seule est reprise telle quelle. Le gain local des six paires ne s’extrapole donc pas à ce plan global. Prochaine micro-étape possible : préciser les dépendances qui bloquent les modules coûteux, avant d’élargir l’ordonnancement. Aucun nouveau build b8x ; compilateur installé inchangé et référence backend maintenue à **388,011 s**.

[Rapport et limites](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-scheduler-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-scheduler-20260920/summary.json) · [Validations](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-scheduler-20260920/proof.json) · [Couverture des anciennes paires](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-scheduler-20260920/previous-pairs-coverage.json).

## Point 4 — huitième sous-étape : affiner les dépendances, 29 → 91 paires

Le graphe précédent parcourait des modules encore non publiés et utilisait les imports comme s’ils étaient tous lus par PBO. La fermeture est désormais testée avec deux règles : conserver les références aux modules absents comme feuilles, sans parcourir leur corps ; utiliser les références qualifiées du TAST préparé comme racines. Seule la dépendance du second module vers le premier interdit une paire, car la disponibilité du second ne change pas pour le premier.

Sur les mêmes 304 modules, les paires adjacentes admissibles passent de **47 à 128 puis 149** ; la sélection disjointe passe de **29 à 82 puis 91**, soit **182 modules regroupables**. `Data.List / Data.Map.Internal` et `Data.Foldable / Data.FunctorWithIndex` sont libérées et retenues. Les deux autres paires examinées restent bloquées par la granularité des modules intermédiaires ; leurs bindings responsables sont enregistrés.

**Validation sémantique JavaScript : 304 résultats PBO identiques**, analyses et directives incluses, avec un cache en mémoire reproduisant la disponibilité native et un plan figé avant optimisation. Toutes les lectures observées sont couvertes ; aucun second module regroupé ne lit le premier. Contrôle négatif sur `analysis.externs` réussi. Les `Map` sont comparées par leurs entrées après vérification de leurs tailles/hauteurs, sans exiger la même forme d’arbre. Les **304 fermetures initiales natives** et les **304 TAST** correspondent exactement aux références.

**Aucun gain de temps natif encore mesuré.** Ce rejeu utilise des calculs JS successifs sur un cache figé ; il ne teste pas les goroutines ni l’émission. Le patch du planificateur natif est préparé, non appliqué et non compilé. Prochaine étape : validation native et `-race`, puis benchmark incluant l’émission. Compilateur installé inchangé ; référence b8x maintenue à **388,011 s**.

[Rapport](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-dependencies-20260920/rapport.md) · [Plans](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-dependencies-20260920/plans.json) · [Validation](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-dependencies-20260920/validation.json) · [Patch natif à tester](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-dependencies-20260920/native-planner.patch).

## Point 4 — neuvième sous-étape : 91 paires en natif, aucun gain mesuré

Le plan affiné est testé dans une copie native du prototype : **91 paires disjointes**, noms et indices identiques au plan JavaScript préalable, soit 182 modules regroupés sur 304. Les goroutines calculent les résultats PBO sur le cache figé ; les publications et l’émission Go conservent leur ordre.

**Validation native et `-race` réussies** : 304 comparaisons complètes des résultats PBO par parcours, analyses/directives incluses, 2 906 noms de dépendances audités et aucune course signalée. Les **399 Go et 304 TAST restent identiques** dans les dix runs, comprenant deux validations, deux chauffes et six mesures.

| Mesure moyenne, trois runs par variante | Référence | 91 paires | Variation |
|---|---:|---:|---:|
| Optimisation + émission | 7,688 s | 7,877 s | +2,45 % |
| Backend complet | 13,169 s | 13,155 s | −0,11 % |
| CPU utilisateur + système | 60,764 s | 61,835 s | +1,76 % |

La planification prend 19,85 à 20,66 ms, incluses dans la phase. **Aucune accélération démontrée ; pas d’intégration.** La stabilité du total masque les variations de chargement/préparation et ne prouve pas un gain de l’ordonnanceur. La référence historique du corpus était 7,482 s ; la comparaison pertinente reste ici celle des variantes alternées dans la même session. Trois répétitions ne démontrent pas une régression générale du parallélisme PBO.

L’élargissement du plan de 29 à 91 paires ne suffit donc pas. Prochaine question à mesurer : quelles durées des calculs simultanés et quelles attentes de l’émission annulent les gains locaux ? Leur cause précise n’est pas établie ; ne pas élargir encore le graphe sans cette mesure. Aucun nouveau build b8x, sources et compilateur installés inchangés ; référence backend maintenue à **388,011 s**.

[Rapport](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-prefix-native-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-prefix-native-20260920/summary.json) · [Validations](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-prefix-native-20260920/proof.json).

## Point 4 — dixième sous-étape : le chevauchement est annulé par l’allongement des calculs

Quatre parcours instrumentés du même binaire (séquentiel/parallèle/parallèle/séquentiel), mêmes entrées et sorties vérifiées. Avec les 91 paires, la somme des durées PBO passe de **6,377 à 8,115 s**, tandis que leur union reste à **6,377 → 6,441 s**. Pour les seuls 182 modules regroupés : **3,312 → 4,996 s** cumulées. L’attente de l’émission ne baisse pas : **1,264 → 1,327 s**. La phase complète reste **7,685 → 7,814 s**. Ce sont des intervalles muraux, pas du CPU ; leurs chevauchements empêchent de les additionner.

Les **399 Go et 304 TAST sont identiques** dans les quatre runs. Le parallélisme est réel, mais aucun gain exploitable n’est établi. La part du GC, du scheduling et de la contention mémoire dans l’allongement reste inconnue. Conformément à la demande utilisateur, on passe au chantier plus large du **point 3 : réduire les reconstructions répétées de la collecte transitive**, en conservant ce diagnostic. Pas de modification du compilateur installé ni de nouveau build b8x.

[Rapport](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-waits-20260920/rapport.md) · [Mesures](/Users/0x1/Documents/htdocs/scratch/gopurs-pbo-waits-20260920/summary.json).

## Diagnostic de référence — vrai `b -c -n` du 19 septembre 2026

Révisions mesurées : b8x `3ea9c731cbef5c7bb8593b4d2e8a5594a313bb26`, gopurs `fdd3113d08784f5337058a8b22f015379475ae35`, PBO gopurs `87f6d0220aae744a3538513dd4aaeebe8cfcb05c`. Les versions détaillées et profils sont référencés ci-dessous.

**Le backend reste le principal coût : 9 min 52 s de travail interne sur 11 min 16 s au total. Une cible précise ressort : le scanner d’imports Go et les comparaisons Char/String qu’il utilise produisent énormément d’allocations. Le GC domine le CPU échantillonné, mais cela ne signifie pas qu’il explique la même proportion du temps d’attente.**

La génération Go et `go mod tidy` ont réussi. Deux étapes finales ont échoué silencieusement : sourcemaps et configuration. Le script a néanmoins affiché `Done!` et retourné 0. Cette mesure décrit donc le parcours réellement exécuté, avec ces échecs ; elle ne mesure pas une génération de configuration réussie.

## Ce qui a été exécuté

Un seul `b -c -n`, depuis b8x, avec son frontend PureScript habituel et la reconstruction effective du compilateur natif. 2 655 modules TAST, 204,56 Mio de JSON, puis 2 959 fichiers Go produits. Le TAST est enregistré sous le nom `corefn.json` : le nom du fichier n’en fait pas le CoreFn standard.

- Go 1.27.0, darwin/arm64 ; runtime natif configuré avec 14 P.
- Chargement TAST séquentiel et émission jusqu’à 2 modules indépendants, réglages par défaut ; aucune surcharge `GOPURS_JOBS`, `GOPURS_EMIT_JOBS`, `GOGC`, `GOMEMLIMIT` ou `GOMAXPROCS`.
- `PPROF=1` et `GODEBUG=gctrace=1` activés uniquement sur le gopurs natif qui compile b8x.
- Marqueurs temporaires autour des étapes du script ; binaire exact et sources Go du compilateur conservés pour lire les profils.
- Aucun fichier source suivi modifié. Le build a normalement reconstruit les sorties et le binaire ignorés par Git. Les dépôts b8x et gopurs étaient propres à la fin de cette mesure.

Les versions, SHA Git et SHA-256 du binaire sont dans [provenance.json](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/provenance.json) et [events.jsonl](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/events.jsonl).

## Où passent les 11 min 16 s

Les lignes suivantes sont additives, arrondies au centième de seconde.

| Étape | Durée murale |
|---|---:|
| Reconstruction de gopurs natif | 42,40 s |
| Nettoyage b8x | 1,97 s |
| Frontend PureScript de b8x | 27,60 s |
| Gopurs : chargement/tri du TAST | 34,28 s |
| Gopurs : préparation + monomorphisation | 131,43 s |
| Gopurs : optimisation + émission | 425,95 s |
| Reste du processus natif, dont finalisation des profils | 5,12 s |
| `go mod tidy` | 1,04 s |
| Sourcemaps — échec | 0,18 s |
| Configuration — échec | 1,61 s |
| Autres coûts d’orchestration et étapes mineures | 4,51 s |
| **Total** | **676,08 s** |

Le chronomètre interne du backend indique **591,665 s** ; son processus complet dure **596,775 s**, profils compris. L’optimisation/émission représente 72 % du temps interne du backend et 63 % du total. La préparation/monomorphisation en représente respectivement 22 % et 19 %. Le profil n’isole pas leurs sous-étapes murales davantage.

La reconstruction de gopurs comprend 1,2 s de backend Node, 9,5 s de frontend TAST, 29,1 s de génération Go, puis environ 1,8 s pour l’étape de build Go. Les caches globaux habituels restent actifs : `-c` nettoie les sorties du projet, pas tous les caches de la machine.

## Ce que le GC révèle vraiment

Le profil contient 2 023,53 secondes CPU échantillonnées, réparties sur environ 596 secondes murales. Une union des piles GC, sans additionner les cumuls imbriqués, donne :

| CPU échantillonné | Secondes CPU | Part du profil |
|---|---:|---:|
| GC sur workers idle | 921,00 | 45,51 % |
| Autre travail GC | 552,22 | 27,29 % |
| **Ensemble GC** | **1 473,22** | **72,80 %** |

Les workers idle exécutent le GC lorsqu’un P ne trouve pas d’autre travail Go prêt à s’exécuter. Ils peuvent céder lorsque du travail arrive. Cela montre de la capacité disponible pendant ces fenêtres ; cela ne distingue pas à lui seul une partie séquentielle, des dépendances ou une attente. Ce travail peut néanmoins concurrencer le programme en mémoire/cache.

Les traces enregistrent **1 206 cycles**, environ **0,106 s de pauses globales cumulées**, et une pause individuelle maximale de **8,21 ms**. L’assistance au marquage imposée aux allocations ne représente qu’environ **0,67 s CPU estimée**. Ces chiffres ne mesurent pas tous les effets indirects du GC ni le coût de l’allocation et du balayage.

**Je corrige donc la portée du constat précédent : le GC est lourd en CPU sur le vrai b8x, mais ce profil ne prouve pas qu’un réglage du GC diviserait le temps de compilation.** Le pourcentage `gctrace` utilise encore un autre dénominateur — la capacité cumulée des P, avec marquage idle exclu — et ne se compare pas directement à pprof.

La mesure système du processus donne 2 642,93 s CPU, soit environ **4,43 CPU utilisés en moyenne**, GC compris, pour 14 P configurés. Pprof donne environ 3,40 CPU échantillonnés ; cet écart reste inexpliqué ici. Les chiffres pprof servent à localiser les coûts, pas à remplacer la mesure CPU système.

## Le principal problème concret : les allocations

Le profil mémoire estime **1 030,31 Gio alloués cumulativement**, soit environ **1,01 Tio**, et **32,47 milliards d’allocations d’objets**. Ce n’est pas la RAM occupée simultanément : le pic de RSS du processus est **6,60 Gio**. Le profil de fin rapporte environ **1,82 Gio vivants** ; il décrit l’état du dernier GC, pas un pic exact ni un diagnostic de fuite.

| Chemin observé | Octets alloués cumulés | Part des allocations |
|---|---:|---:|
| `GoCode.referencedImports`, appels descendants inclus | 238,87 Gio | 23,18 % |
| Monomorphisation des modules, appels descendants inclus | 209,14 Gio | 20,30 % |
| Wrapper FFI de comparaison `Char`, allocations propres | 114,73 Gio | 11,14 % |
| `gopurs_runtime.Str`, allocations propres | 103,02 Gio | 10,00 % |
| Wrapper FFI de comparaison `String`, allocations propres | 56,65 Gio | 5,50 % |

**Les lignes de ce tableau se recouvrent et ne s’additionnent pas.** Les mesures mémoire sont des estimations extrapolées à partir de l’échantillonnage Go. Les substitutions de types et `mangleType` ressortent aussi : le problème ne se limite pas au scanner.

`referencedImports` coûte **100,37 s CPU échantillonnées hors GC**. Cela représente 52,7 % des échantillons reconnus comme génération Go hors GC, selon une classification par piles ; ce n’est ni une durée murale du scanner ni un gain garanti.

La lecture du code généré explique une partie du coût. Dans [GoCode.purs:36](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/GoCode.purs:36), la reconnaissance d’un caractère d’identifiant utilise sept comparaisons d’ordre. Le [Go réellement exécuté](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/native-source/purescript/Gopurs_GoCode.go:472) appelle pour chacune `Apply5(ordCharImpl, LT, EQ, GT, Str(char), Str(constante))`, avec aussi des closures intermédiaires pour les booléens. Les sept comparaisons sont calculées avant le résultat final. L’accès au caractère est déjà direct : il ne s’agit pas d’un retour au parcours répété des préfixes UTF-16.

Les fonctions anonymes `init.func197` et `init.func200` correspondent aux wrappers `OrdCharImpl` et `OrdStringImpl`. Le profil attribue leurs allocations aux appels FFI [ligne 85](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/native-source/purescript/Data_Ord_ffi.go:85) et [ligne 115](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/native-source/purescript/Data_Ord_ffi.go:115), pas aux lignes `Unbox` ou `Box` voisines. La FFI reçoit les trois valeurs `Ordering` comme `interface{}` : les conversions des `Value` expliquent les objets de 24 octets observés. Le microbenchmark et l’analyse d’échappement ajoutés au point 1 confirment désormais ce mécanisme.

## Conséquences pour le parallélisme et les prochains gains

Aff permet déjà l’émission parallèle de deux modules indépendants. Cela ne rend pas automatiquement parallèles les calculs purs, la monomorphisation et tout l’optimiseur. Le PBO actuel transmet les directives exportées d’un module au suivant ; son ordonnanceur reste séquentiel. La configuration de 14 P ne signifie donc pas 14 compilations applicatives actives.

Il existe bien des chantiers concrets, dans cet ordre de vérification :

1. **Réduire le coût des comparaisons Char/String et du scanner d’imports.** Commencer par un microbenchmark isolé de l’adaptation FFI actuelle contre une variante sans conversions inutiles, avec mesure des octets/allocations par appel et vérification de LT/EQ/GT. Puis examiner la génération de comparaisons natives grâce aux types connus du TAST, et la propagation des imports pour éviter des scans inutiles. Le profil établit le coût ; aucun gain de cette modification n’a encore été mesuré.
2. **Cibler les parcours et substitutions de la monomorphisation.** La phase dure 131 s et alloue environ 209 Gio dans les piles reconnues. Les parcours `transitiveCollect`/`collectExpr` et substitutions fournissent des points de départ précis, avant tout refactoring global.
3. **Élargir le parallélisme applicatif au-delà de l’émission.** La capacité disponible rend la piste pertinente. Il faut d’abord séparer les calculs indépendants de la transmission des directives et mesurer le résultat sur b8x. Augmenter seulement `GOPURS_EMIT_JOBS` n’enlève pas ces dépendances.

Le profil ne permet pas de promettre ×2, ni de convertir les 100 s CPU du scanner en 100 s murales récupérables. Il apporte en revanche une cible beaucoup plus précise qu’« ajouter des goroutines » ou « régler le GC ».

## Comparaison avec les références et limites

La référence b8x documentée dans [parallel-emission.md](/Users/0x1/Documents/htdocs/gopurs/gopurs/docs/parallel-emission.md) est de 581,421 s pour le backend après parallélisation, contre 614,509 s avant. Ici, 591,665 s instrumentées restent du même ordre (+1,8 % par rapport au précédent après-changement). Les entrées ont 2 655 modules contre 2 657 dans cette référence, et le contexte/cache/profilage diffère : ceci n’est pas une nouvelle comparaison contrôlée avant/après.

Les baselines officielles du [README altbak.pub](/Users/0x1/Documents/htdocs/altbak.pub/README.md:53) mesurent l’exécution des programmes compilés, pas cette compilation. Elles ne fournissent donc pas un objectif directement comparable aux 676 s mesurées ici. L’ancien profil GC altbak ne sert plus de substitut à b8x.

Les sorties de sourcemaps sont masquées dans [build:120](/Users/0x1/Documents/htdocs/b8x/bin/build:120) ; le wrapper a enregistré un code 1, et le glob `output/*/*.purs` ne contient aucun fichier. L’échec de `conf` est explicitement absorbé dans [build:124](/Users/0x1/Documents/htdocs/b8x/bin/build:124), avec un code 1 relevé par le marqueur de sortie. Sa cause exacte n’a pas été collectée puisque le script supprime ses sorties. **Le temps d’un éventuel build Go de configuration réussi n’est donc pas établi par ce run.**

Un seul run complet a été effectué. Le coût de finalisation du profil mémoire et les effets du profilage sont présents ; les allocations sont échantillonnées, les piles CPU ne couvrent pas exhaustivement chaque phase. Aucune optimisation ni correction de ces échecs n’a été implémentée dans cette tâche de diagnostic.

## Preuves conservées

- [Synthèse chiffrée JSON](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/summary.json), [journal du build](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/build.log), [chronométrage et GC natifs](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/native.stderr.log).
- [Analyse CPU détaillée](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/cpu-analysis.md), [allocations par site](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/alloc-space-top.txt), [allocations par chemin](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/alloc-space-cum.txt), [mémoire vivante](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/live-space-top.txt).
- [Profil CPU brut](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/cpu.prof), [profil mémoire brut](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/mem.prof), [binaire exact](/Users/0x1/Documents/htdocs/scratch/b8x-profile-20260919/gopurs-native).

Depuis ce dossier, les profils se relisent sans lancer de build :

```sh
go tool pprof -top ./gopurs-native ./cpu.prof
go tool pprof -top -sample_index=alloc_space ./gopurs-native ./mem.prof
go tool pprof -top -sample_index=inuse_space ./gopurs-native ./mem.prof
```
