# Gopurs — réduire le temps de compilation de b8x

État au 20 septembre 2026 : points 1 et 2 intégrés, puis une première sous-étape de substitution validée sur le vrai `b -c -n`. Le backend passe désormais de la référence initiale de 591,665 à 433,732 s (−26,69 %). La garde de substitution ajoute un gain observé de 15,906 s (−3,54 %) après le scanner. Le cœur préparation/monomorphisation du point 3 et les points 4 à 7 restent ouverts. Ce document remplace intégralement l’ancienne liste.

## Points à traiter un par un, ultérieurement

Travailler sur un seul point à la fois. Commencer par une expérience courte, vérifier le comportement et mesurer le gain avant de passer au suivant. Les chiffres du diagnostic ci-dessous sont des observations ; les gains proposés restent à établir.

- [x] **1. Comparaisons Char/String et adaptation FFI.** Deux signatures Go génériques dans gopurs-prelude, arguments génériques transmis directement par le bridge. Suppression des trois allocations/72 octets par comparaison confirmée ; tests JS/natifs et mesure complète b8x terminés. Backend : 591,665 → 493,217 s ; allocations : 1 030,31 → 858,66 Gio. Bilan et limites ci-dessous.
- [x] **2. Scanner d’imports Go.** Prédicat d’identifiant avec une conversion en entier et des conditions explicites : moins de chaînes et de closures. Microbenchmark ×3,04 ; backend b8x 493,217 → 449,638 s (−8,84 %), allocations 858,66 → 789,39 Gio. Les 2 959 Go produits restent identiques. La propagation des imports pour éviter des scans reste une piste distincte, non implémentée. Bilan ci-dessous.
- [ ] **3. Préparation et monomorphisation.** Une garde dans `TypeSubstitution.substitute` est intégrée : elle concerne surtout l’optimisation/émission, pas `Substitute.substituteExprType` de la monomorphisation. Gain backend observé : 449,638 → 433,732 s. Le cœur du point reste ouvert : cibler `transitiveCollect`, `collectExpr`, `mangleType` et les substitutions de la monomorphisation ; vérifier un recalcul précis avant toute transformation globale.
- [ ] **4. Parallélisme applicatif au-delà de l’émission.** Identifier les calculs indépendants de la préparation et de l’optimisation PBO ; expliciter la transmission des directives entre modules avant de modifier l’ordonnancement. Mesurer sur b8x, conserver le déterminisme du résultat et vérifier les accès partagés. Augmenter seulement le nombre de workers d’émission ne résout pas les dépendances séquentielles.
- [ ] **5. Échecs masqués dans le build b8x.** Diagnostiquer puis corriger l’échec des sourcemaps et de `conf`, et rendre leurs statuts visibles. Le glob des sourcemaps ne trouve aucun `.purs` ; la cause précise de l’échec de configuration reste inconnue. Mesurer le coût de la configuration lorsqu’elle réussit.
- [ ] **6. GC, après réduction des allocations.** Reprofiler le même travail ; distinguer GC idle, autres travaux GC, pauses et coût mural. N’évaluer un réglage du GC qu’avec une comparaison contrôlée temps/mémoire. Ne pas convertir les 72,8 % de CPU échantillonné en promesse de gain mural.
- [ ] **7. Bilan complet sur le vrai workflow.** Après validation de chaque changement isolé, répéter `b -c -n` avec les mêmes entrées, paramètres et conditions de cache ; publier le détail des phases, allocations, pic mémoire et statuts. Comparer aux références de compilation b8x ci-dessous ; garder les baselines d’exécution officielles d’altbak.pub pour leur périmètre propre.

Chaque point doit laisser un résultat vérifié, les mesures avant/après et ses limites dans ce document. Aucun gain ×2 n’est établi par le profil actuel.

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

Prochaine micro-étape du **point 3** : mesurer les répétitions de `mangleType` ou de `transitiveCollect` dans la préparation/monomorphisation. Référence actuelle du backend : **433,732 s**.

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
