# Gopurs — réduire le temps de compilation de b8x

État au 21 septembre 2026 : comparaisons Char/String, scanner, garde de substitution, clés différées, pipeline PBO/émission, cache local des corps préparés, index des champs de conversion, Array.any générique et **cache borné des instanciations PBO** intégrés. Dernier backend b8x : **268,011 s (4 min 28)**, contre 591,665 s au départ (évolution historique **−54,70 %, 5 min 24 gagnées**). Les essais de garde avant unification et d’ordonnanceur PBO à 91 paires ne sont pas retenus faute de gain établi.

**Dernière sous-étape, point 8 : 319,381 → 268,011 s de backend (−16,08 %), allocations 541,68 → 458,57 Gio (−15,34 %).** `b -c -n` entier : **349,175 s (5 min 49)**. **63 tests JS et 9 tests Go sous `-race` passent ; 2 655 TAST et 2 959 Go strictement identiques.** Gain également observé sur le petit corpus alterné ; un seul nouveau run complet, secondes exactes à consolider. Pic RSS : 6,444 → 6,190 Gio. Sourcemaps encore en échec, statut conf toujours masqué. Points 8/9 : sous-étapes conservées ; prochaine priorité importante : **point 3, les parcours restants de collecte transitive**. Les points 3 à 8 et 10 restent ouverts.

## Points à traiter un par un, ultérieurement

Travailler sur un seul point à la fois. Commencer par une expérience courte, vérifier le comportement et mesurer le gain avant de passer au suivant. Les chiffres du diagnostic ci-dessous sont des observations ; les gains proposés restent à établir.

- [x] **1. Comparaisons Char/String et adaptation FFI.** Deux signatures Go génériques dans gopurs-prelude, arguments génériques transmis directement par le bridge. Suppression des trois allocations/72 octets par comparaison confirmée ; tests JS/natifs et mesure complète b8x terminés. Backend : 591,665 → 493,217 s ; allocations : 1 030,31 → 858,66 Gio. Bilan et limites ci-dessous.
- [x] **2. Scanner d’imports Go.** Prédicat d’identifiant avec une conversion en entier et des conditions explicites : moins de chaînes et de closures. Microbenchmark ×3,04 ; backend b8x 493,217 → 449,638 s (−8,84 %), allocations 858,66 → 789,39 Gio. Les 2 959 Go produits restent identiques. La propagation des imports pour éviter des scans reste une piste distincte, non implémentée. Bilan ci-dessous.
- [ ] **3. Préparation et monomorphisation.** Une garde dans `TypeSubstitution.substitute` est intégrée : elle concerne surtout l’optimisation/émission, pas `Substitute.substituteExprType` de la monomorphisation. Gain backend observé : 449,638 → 433,732 s. Deux gardes évitent maintenant les clés de spécialisation inutiles : préparation 116,209 → 100,035 s, backend 433,732 → 424,522 s. Le cache local du préfixe de `transitiveCollect` est maintenant intégré : préparation 96,746 → 81,535 s, backend 388,011 → 374,255 s ; 2 655 TAST et 2 959 Go identiques. Le cœur du point reste ouvert, notamment les passes `monomorphizeExpr`/`collectExpr` encore répétées ; établir leurs dépendances avant de supprimer un calcul.
- [ ] **4. Parallélisme applicatif au-delà de l’émission.** Première sous-étape intégrée : le producteur PBO séquentiel optimise pendant que le lot précédent émet du Go. Optimisation/émission : 289,795 → 256,177 s (−11,60 %), backend : 424,522 → 388,011 s. Un lot actif, un lot en attente, ordre des signatures et des directives conservé ; tests JS/natifs et détection des courses validés. Pour aller plus loin, identifier les calculs indépendants de préparation/PBO et traiter explicitement la transmission des directives entre modules avant de paralléliser l’optimisation elle-même.
- [ ] **5. Échecs masqués dans le build b8x.** Diagnostiquer puis corriger l’échec des sourcemaps et de `conf`, et rendre leurs statuts visibles. Le glob des sourcemaps ne trouve aucun `.purs` ; la cause précise de l’échec de configuration reste inconnue. Mesurer le coût de la configuration lorsqu’elle réussit.
- [ ] **6. GC, après réduction des allocations.** Reprofiler le même travail ; distinguer GC idle, autres travaux GC, pauses et coût mural. N’évaluer un réglage du GC qu’avec une comparaison contrôlée temps/mémoire. Ne pas convertir les 72,8 % de CPU échantillonné en promesse de gain mural.
- [ ] **7. Bilan complet sur le vrai workflow.** Après validation de chaque changement isolé, répéter `b -c -n` avec les mêmes entrées, paramètres et conditions de cache ; publier le détail des phases, allocations, pic mémoire et statuts. Comparer aux références de compilation b8x ci-dessous ; garder les baselines d’exécution officielles d’altbak.pub pour leur périmètre propre.
- [ ] **8. Coût interne des passes PBO, au-delà du parallélisme.** Garde générale sur les expressions écartée ; `Array.any` générique intégré. **Cache FIFO de 512 instanciations par module désormais intégré**, créé explicitement dans le builder, clés expression/type complets, résultats immuables et directives conservées. La réutilisation des résultats favorise aussi les instanciations suivantes. Backend b8x **319,381 → 268,011 s**, allocations **−83,11 Gio**, 2 655 TAST/2 959 Go identiques. Les autres parcours PBO restent à traiter selon un profil actualisé ; le gros chantier suivant est la collecte transitive du point 3.
- [x] **9. Indexer les champs de conversion Go.** Index immuable construit une fois depuis les layouts TAST, partagé entre émissions. Premier résultat dans l’ordre des clés, priorité constructeur/classe et deux alias conservés. `findReboxFields` passe de **150,56 Gio à 69 Mio** et de 53,78 à 0,04 s CPU échantillonnées hors GC. Backend b8x : **401,157 → 342,863 s**, allocations totales −21,27 %. Les **90 tests**, 4 468 recherches sur les métadonnées b8x et les comparaisons natives passent ; **2 655 TAST et 2 959 Go identiques**. Bilan et limites ci-dessous.
- [ ] **10. Compilation incrémentale entre builds.** Réutiliser les résultats des modules inchangés pour le cycle quotidien. Définir l’invalidation des signatures, directives PBO, spécialisations transitives, FFI et versions du compilateur ; comparer les sorties avec un build propre. Ce chantier est distinct de l’accélération de `b -c -n` : un cache invalidé par `-c` n’améliore pas ce rebuild forcé. Potentiel à mesurer sur une modification réelle de b8x, pas sur la seule commande propre.

Chaque point doit laisser un résultat vérifié, les mesures avant/après et ses limites dans ce document. Consolider les gains cumulés historiques avec des séries de runs comparables.

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
