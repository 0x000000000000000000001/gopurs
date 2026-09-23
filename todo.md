# Gopurs — réduire l’écart JSON/TAST et accélérer la compilation

## Campagne du 23 septembre — référence et profil attribué du décodage TAST

**Référence reproduite avec le compilateur actuel.** TAST (12 modules,
21 574 types, 5,55 Mo de JSON) : Go **25,24 / 371,00 / 423,14 ms**
(parse / décodage / total), JS **20,41 / 56,37 / 79,50 ms** ; allocations
décodage **442 632 472 octets**, total **498 663 744**, identiques à la
campagne de référence. Ratio décodage **×6,58**, total **×5,32**.

**Référence b8x complète** (`b -c`, compilateur natif reconstruit par le
script) : code 0, **250,01 s** mural, **160,21 s** de backend, pic RSS
**5,70 Gio** ; phases `[gopurs]` : chargement 9,64 s, préparation 38,99 s
(dont spécialisations transitives 30,01 s), optimisation/émission 111,58 s.
C’est la nouvelle base de comparaison (l’ancien bilan publié était
181,22 s / 267,70 s).

**Profil du décodage seul** (copie scratch du Go généré, corpus parsé une
fois, oracles 12 modules revérifiés dans le processus ; binaire
`decode-profile`, profils CPU, allocations et vivant dans
`scratch/tast-decode-20260923`) :

- **442,6 Mo et 10,88 M d’objets par passe**, pour un AST décodé d’environ
  **15 Mo retenus** : ~96 % des octets et objets sont transitoires.
- **Table de types : 54,8 % des allocations** (21 574 types, ~11 Ko et
  ~280 objets par type). Le bouclage n’est **pas quadratique** : simulation
  exacte de l’algorithme → **2,47 tentatives de résolution par type**
  (10–12 tours, 53 189 tentatives).
- Le code généré du décodeur lui-même ne représente que **~5 % des objets
  alloués**. Le reste est de la glue runtime/FFI : `gopurs_runtime` 19,6 %,
  adaptateurs ST **16,2 %**, accesseurs `Either`/`Maybe` 8,9 %,
  `Foreign.Object` 4,8 %, `StateT` 3,8 %, `Data.Ord` 3,0 %, `Record_Unsafe`
  2,1 %, `Partial.Unsafe` 1,3 %. À plat : `Any` 6,2 %, `Right` 4,3 %,
  `Just` 4,1 %, `Apply`/`Apply2` 6,8 %, adaptateurs ST uncurryés ~4,7 %.
- Le CPU est dominé par le GC (~74 % des échantillons à plat), pas par le
  calcul.

**Conséquence pour la stratégie.** Réécrire la source PureScript du décodeur a
un plafond faible (son code généré est ~5 % des objets) ; l’essentiel est dans
la **couche d’appel runtime/FFI**. Chaque opération `ST` traverse le pont
étranger, qui recrée des closures d’adaptation `any` et encadre arguments et
résultats (`bind_` : deux closures et quatre boîtes d’interface par appel).

**Première expérience mesurée (sonde scratch, 23 septembre) : la couche
ST/FFI.** Sur le Go généré figé, les wrappers `Control.Monad.ST.Internal`
(`bind_`, `map_`, `pure_`, `run`, `while`, `forImpl`, `foreach`) et
`runSTFn1..9` remplacés par des helpers `Value` natifs donnent, en paires
alternées : décodage **385,91 → 334,12 ms (−13,42 %)** au minimum,
**−11,27 %** en médiane, **−12,06 %** en mode allocations, et
**442 632 766 → 402 378 712 octets (−9,09 %, −40,25 Mo)** par passe, oracles
12 modules identiques. La sonde ne touche ni les implémentations
`Data.Array.ST` (conversions `[]Value` ↔ `[]any` restantes) ni le décodeur :
c’est un **plafond de la couche ST/FFI**, pas encore un résultat du
compilateur. Le rapport complet, les protocoles et la provenance sont dans
[2026-09-23-tast-decode-profile.md](../../altbak.pub-gopurs/docs/benchmark-results/2026-09-23-tast-decode-profile.md).

### Déroulement prévu

1. [x] **Référence** compilateur actuel : corpus TAST puis b8x complet.
2. [x] **Profil du décodage seul** (CPU, allocations, vivant), hors
   empreintes, avec oracles dans le processus.
3. [x] **Sonde scratch** : helpers ST `Value` natifs sur le Go généré figé —
   **−13,42 % temps / −9,09 % octets**, oracles identiques. Le gain est
   démontré ; la prochaine étape est de l’obtenir depuis gopurs lui-même
   (intrinsèques de codegen pour `Control.Monad.ST` ou ABI `Value`),
   puis de mesurer la campagne appariée et le b8x complet.
4. [ ] **Portage dans gopurs** : intrinsèques de codegen pour les opérations
   ST (ou ABI `Value` pour `Control.Monad.ST`), en gardant le comportement des
   dictionnaires, des applications partielles et des ordres d’évaluation.
   Étendre ensuite aux implémentations `Data.Array.ST` si le gain se
   confirme, puis mesurer.
5. [ ] **Validation sur le code réellement régénéré** : AST complets, erreurs
   et leur priorité, références de types et cycles, ordre des callbacks,
   immutabilité.
6. [ ] **Mesure appariée avant/après** : décodage, total, allocations, taille
   du binaire, puis phases et pic mémoire du workflow b8x.
7. [ ] **Décision d’intégration** fondée sur les chiffres ; documenter le gain
   et ses limites ici.

**Validation obligatoire à chaque étape : répéter `b -c` sur le vrai b8x**
(mêmes entrées, paramètres et état de cache), publier le détail des phases, le
pic mémoire et les statuts, et rejeter toute attribution non appuyée par une
campagne appariée. Les 8–9 ms des moteurs manuels restent un potentiel
architectural, pas un résultat acquis.

## Ensuite : modèle d’appel natif de bout en bout (point 12)

Conserver arguments et résultats typés depuis le décodeur de champ, à travers
le worker partagé, jusqu’au stockage final :

- callbacks stockés dans les dictionnaires/getters, avec arguments et succès
  natifs à travers les workers ;
- records retournés et rangées ouvertes complètes, champs supplémentaires
  compris ;
- tableaux finaux encore en `[]Value` et allocations ADT qui échappent au
  flux natif ;
- interfaces/descripteurs plats selon les besoins, sans chaînes de wrappers.

Contraintes : garder immutabilité, ordre, première erreur, captures, appels
partiels et ABI FFI ; mesurer temps, allocations et taille.

**Ne pas répéter la sonde des seules enveloppes `Either`/`Maybe`** (aucun gain
de débit démontré malgré **−6,09 % d’octets**) et ne pas présenter la
construction unique des records comme une ABI native complète.

## Accélération de la compilation

- [ ] **Optimisation PBO parallèle.** Chargement, émission et collecte
  transitive utilisent déjà plusieurs workers ; **l’optimisation PBO entre
  modules reste séquentielle**. Prochaine étape : définir les dépendances et
  la transmission des directives entre modules, puis ordonnancer les modules
  indépendants en conservant les résultats consultables par les suivants.
  Remesurer le débit et la mémoire avec le runtime corrigé. Défauts natifs :
  chargement 8, émission 8, préparation 2 ; chargement JS 1.
- [ ] **Coûts internes des passes PBO.** Reprofiler les recherches et
  reconstructions répétées après les corrections récentes (cache FIFO, cache
  externe, collecte transitive parallèle, boxage ADT). Surveiller
  `exprImports` et les conversions du bridge `Array.slice` révélées par le
  prototype parallèle.
- [ ] **Compilation incrémentale persistante.** Réutiliser les modules
  inchangés dans le cycle quotidien. Définir l’invalidation (signatures,
  directives PBO, spécialisations transitives, FFI, version du compilateur),
  comparer les sorties avec un build propre, et mesurer sur une modification
  réelle de b8x — pas sur le seul rebuild forcé `-c`.
- [ ] **Réglage du GC**, seulement après réduction des allocations dominantes.
  Le profil b8x actuel n’a pas refait la trace GC complète ; dernier relevé
  connu : 635 cycles, 44,915 ms de pauses globales, pic RSS 6,203 Go (à
  actualiser). Comparer temps et mémoire de façon contrôlée ; chantier
  secondaire.

## Références de mesure

- Compilateur : `gopurs` natif (`bin/gopurs-native`), fork TAST local pour
  `purs`, PBO dans `purescript-backend-optimizer-gopurs`.
- Diagnostic : `altbak.pub-gopurs/bin/benchmark/json-diagnostic.py`
  (`--suite JsonTypedAst` ou `JsonDecoding`), GOMAXPROCS=1, GOGC=100, PGO
  désactivé, cinq échantillons, médiane des minima de processus.
- Sonde TAST : `scratch/tast-decode-20260923` (`build-profile.sh`,
  `run-profile.sh cpu|alloc|validate|live`), corpus
  `ef5ed1bae6ae52d084a4b3fd9d1e9e90a5da2c3a2d1e6e3aa56122007e0b223a`.
- Baselines README (23 septembre 2026) : **Go 15,11 / 424,10 ms**,
  **JS 9,28 / 80,26 ms**. Baselines de compilation b8x (23 septembre 2026) :
  **chargement 9,638 s, préparation 38,991 s, optimisation/émission
  111,579 s, backend 160,212 s, commande 250,01 s, pic RSS 5,70 Gio**.
- Les pourcentages cumulés (`Apply2`, `TraverseArrayImpl`, …) ne mesurent pas
  un coût propre et ne s’additionnent pas.
