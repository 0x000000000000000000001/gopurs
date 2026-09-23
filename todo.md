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

**Portage intégré (23 septembre) : la couche ST/FFI est passée en `Value` natif.**
`Control.Monad.ST.Internal` (`map_`, `pure_`, `bind_`, `run`, `while`,
`forImpl`, `foreach`, `newImpl`, `read`, `modifyImpl`, `write`) et
`Control.Monad.ST.Uncurried` (`mkSTFn1..9`, `runSTFn1..9`) déclarent désormais
callbacks et résultats en `gopurs_runtime.Value` ; le pont FFI les passe
directement, sans closures d’adaptation `any` ni boîtes. Décodage TAST :
**−11,88 % en médiane sur 10 paires appariées** (avant 398,36 → après
352,75 ms), allocations **−9,09 %** (442 632 472 → 402 378 424 octets).
Mesure officielle : **Go 358,78 / 396,74 ms**, contrôle JS 60,64 / 81,77 ms
(bundle identique) ; cellules README Go **396,74 ms**, JS **81,77 ms**.
Validation : suite `gopurs-st` verte, bootstrap natif OK, 12 modules conformes.
**Côté compilateur : aucun gain ni régression mural établi** (b8x `b -c` :
250,0 s avant, 260,9 / 258,5 s après ; A/B backend contrôlé et A/B
`gopurs-aff` recoupés, dérive machine non séparée) ; le CPU baisse
systématiquement (~2 %) et le chargement TAST progresse. Question laissée
ouverte pour le point 7.

**Intégration suivante (23 septembre) : la résolution de la table de types passe
en Go natif.** Le chemin `decodeTypeTable` (51,3 % des allocations de décodage,
~11 Ko et ~280 objets par type) est implémenté en FFI Go
(`CoreFn/Json.go`) derrière `foreign import decodeTypeTableImpl` ; le JS
conserve l’algorithme PureScript via `TypeTable.decodeTypeTablePS`. Campagne
appariée (5 paires, deux workspaces régénérés, oracles dans chaque processus) :
décodage **363,41 → 186,06 ms (−48,80 %)**, allocations
**402 378 424 → 213 520 672 (−46,94 %)**, total **418,35 → 228,95 ms
(−45,27 %)**, parsing inchangé (+0,21 %, allocations identiques). Cellules
officielles : **Go 246,45 ms**, contrôle JS **87,60 ms** (algorithme JS
inchangé, bundle désormais différent). Validation sémantique : différentiel
natif vs PureScript sur 20 cas fixés et 400 tables aléatoires (cycles,
références hors borne, erreurs différées, `Row`/`ForAll`/`ConstrainedType`),
12 empreintes conformes, bootstrap natif OK. Rapport :
[2026-09-23-native-type-table.md](../../altbak.pub-gopurs/docs/benchmark-results/2026-09-23-native-type-table.md).

**Lecture GC.** Sur le Go généré figé, à code identique : `GOGC=100` 322,0 ms
contre `GOGC=1200` 164,6 ms ; la variante native 163,1 ms à 100 et 88,3 ms à
1200. La moitié du coût antérieur était de la pression GC transitoire, pas du
travail de décodage. Les variantes natives (table de types seule :
−47,5 % min / −46,9 % octets ; + boucles `decodeArray` directes : −49,4 % /
−50,7 % cumulés) sont consignées dans
`scratch/tast-revolution-20260923/README.md`.

### Déroulement prévu

1. [x] **Référence** compilateur actuel : corpus TAST puis b8x complet.
2. [x] **Profil du décodage seul** (CPU, allocations, vivant), hors
   empreintes, avec oracles dans le processus.
3. [x] **Sonde scratch** : helpers ST `Value` natifs sur le Go généré figé —
   **−13,42 % temps / −9,09 % octets**, oracles identiques.
4. [x] **Portage dans gopurs** : la FFI Go de `Control.Monad.ST.Internal` et
   `Uncurried` déclare callbacks et résultats en `gopurs_runtime.Value` ; le
   pont passe directement (pas d’intrinsèques de codegen nécessaires).
   **−11,88 % de décodage sur 10 paires, −9,09 % d’allocations.**
5. [x] **Validation sur le code réellement régénéré** : suite `gopurs-st`
   verte, bootstrap natif, oracles 12 modules dans chaque processus, A/B
   backend et `gopurs-aff`.
6. [x] **Mesure appariée avant/après** : décodage, total, allocations ;
   cellules officielles et README mises à jour. **b8x : pas de gain ni de
   régression murale établis**, CPU en baisse (~2 %), chargement plus rapide.
7. [x] **Table de types native** : FFI Go derrière `decodeTypeTableImpl`, JS
   délégué à l’algorithme PureScript. **−48,80 % de décodage sur 5 paires
   appariées, −46,94 % d’allocations** ; cellules **246,45 / 87,60 ms**.
8. [ ] **Suite du chantier** : intégrer la boucle `decodeArray` directe
   (−7,1 % d’allocations de plus sur la sonde), puis étendre le décodage natif
   au reste du décodeur (`decodeExpr`, `decodeAnnWithUsage`,
   `decodeSourceUsage`), avant les passerelles `unsafePartial`/
   `Array.unsafeIndex` et l’ABI native complète. Reprendre b8x avec une mesure
   appariée dédiée lorsque la question murale devra être tranchée.

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
- Baselines README (23 septembre 2026, après table de types native) : **Go
  191,59 / 246,45 ms**, **JS 65,67 / 87,60 ms** pour TAST (décodage / total) ;
  JSON général **Go 15,11 / 424,10 ms**, **JS 9,28 / 80,26 ms** (inchangés,
  chemin non concerné). Baselines b8x (23 septembre 2026, avant migration) :
  chargement 9,638 s, préparation 38,991 s, optimisation/émission 111,579 s,
  backend 160,212 s, commande 250,01 s, pic RSS 5,70 Gio. Après migration ST :
  commande 258,5–260,9 s, backend 167,3–170,6 s ; A/B backend contrôlé
  167,5/172,2 s (avant) contre 174,1/174,5 s (après) ; minima `gopurs-aff`
  7,009 s contre 7,242 s.
- Migration ST/FFI : `gopurs/gopurs-st/src/Control/Monad/ST/{Internal,Uncurried}.go`,
  workspace `altbak.pub-gopurs/var/benchmark/json-tast-stf-20260923`, campagnes
  `stf-campaign-20260923` et `stf-campaign-2-20260923`.
- Table de types native : `purescript-backend-optimizer-gopurs/src/PureScript/Backend/Optimizer/CoreFn/{Json.purs,Json.go,Json.js,TypeTable.purs}`,
  workspace `altbak.pub-gopurs/var/benchmark/json-tast-native-tt-20260923`,
  résultats `native-tt-results-20260923`, campagne appariée
  `native-tt-campaign-20260923`, rapport
  [2026-09-23-native-type-table.md](../../altbak.pub-gopurs/docs/benchmark-results/2026-09-23-native-type-table.md).
- Les pourcentages cumulés (`Apply2`, `TraverseArrayImpl`, …) ne mesurent pas
  un coût propre et ne s’additionnent pas.
