# Gopurs — Optimisations restantes

> **⚠️ ATTENTION : La commande de validation des performances est `./bin/go/run -c` dans le dossier `altbak.pub`. N'utilisez jamais `./bin/run go` sans argument car cela exécute tous les backends et fausse les métriques. Les tests ciblés de correction sont documentés séparément.**

Ce plan intègre l'audit en lecture seule du 7 septembre 2026 : backend gopurs, dépendance locale `purescript-backend-optimizer` et Go généré dans `altbak.pub/run/bak/go/output/purescript`. La fusion ciblée des thunks a ensuite été implémentée et mesurée le même jour (5.4a). Les autres gains proposés restent à mesurer.

## État actuel et priorités

La TCO, les appels directs, la spécialisation des opérations primitives, les records fermés et certaines structures récursives sont déjà bien optimisés. La boucle mesurée de `Test_Polymorphism` travaille notamment en `int64`, sans dictionnaire ni `Apply` dans ses dix millions d'itérations.

Les coûts restants concernent surtout les fonctions transmises comme valeurs, les conversions entre représentations typées et génériques, les allocations et les structures intermédiaires. `gopurs_runtime.Value` est une structure taguée : sa présence ne prouve pas à elle seule une allocation sur le tas. Un `FALLBACK TCO` ou un `Rebox` présent dans un fichier n'est pas nécessairement exécuté.

Ordre de travail conseillé, en conservant les numéros des chantiers historiques :

1. **Exécution :** closures typées (5.1–5.3), puis préservation des records State (2) et tableaux typés (6).
2. **Exécution et taille du code :** DCE élargie (1), dictionnaires réellement utilisés (3), fusion des parcours (7), représentation mémoire (8).
3. **Compilateur :** cache incrémental (9) et extraction FFI (10), avant l'auto-hébergement (4).
4. **Recherche :** élargissement de l'élimination des thunks au-delà du premier cas implémenté et simplification du calcul (5.4–5.6), sans promettre les performances du cheatcode.

## Performances du code généré

### 1. DCE après spécialisation et génération des wrappers

PBO possède déjà une DCE locale et une DCE de module. Deux limites subsistent : `Monomorphize.purs` conserve les déclarations originales et leurs spécialisations puis exporte tous leurs identifiants ; `Convert.purs` considère ces exports comme vivants. Ensuite, `CodeGen.purs` ajoute des wrappers, temporaires et fonctions de repli, dont certains ne servent plus.

Les wrappers inutilisés ne sont pas forcément gratuits : `Runtime.purs` fait passer chaque `Func`/`FuncN` par `forceEscape`, qui écrit dans la globale `EscapeSink`. Leur construction peut donc encore coûter à l'exécution. La DCE vise aussi les allocations, pas seulement la taille du source et le temps de compilation.

- [ ] **1.1** : Analyser l'atteignabilité après monomorphisation et conserver les exports nécessaires aux appels intermodules, plutôt que toutes les déclarations originales et spécialisées.
- [ ] **1.2** : Analyser les utilisations des wrappers et temporaires introduits par CodeGen. Les représenter dans un AST analysable ou conserver leurs dépendances explicitement : un compteur sur PBO seul ne voit pas les déclarations ajoutées ensuite, notamment via `GoRaw`.
- [ ] **1.3** : Éviter de générer les wrappers et fallbacks dont la valeur n'est jamais utilisée. Distinguer les usages réels des `_ = x`, traiter les groupes récursifs et préserver les effets des expressions évaluées ; le nom `FALLBACK TCO` n'est pas une preuve de code mort.
- [ ] **1.4** : Supprimer les paramètres morts après spécialisation et leurs arguments lorsque leur évaluation peut être retirée. Exemple : le `foldl` spécialisé de `Test_ListOps` additionne nativement mais recharge encore `Get_Data_Semiring_intAdd()` à chaque tour pour un paramètre inutilisé.
- [ ] **1.5** : Valider sur `Test_Polymorphism` et `Test_ListOps` la disparition des constructions inutiles, la conservation des résultats et l'évolution de la taille générée, des allocations et du temps. Ne pas confondre coût de montage d'un appel et coût par itération.

### 2. Records typés de bout en bout, notamment StateMonad

Dans le chemin exécuté de `Test_StateMonad.chainModifications`, le résultat passe par `record dynamique → struct typée → RecordDict → struct générique → RecordDict`. Les structs sont donc parfois déjà connues, puis reperdues aux frontières de fonctions en `func(Value) Value`.

- [ ] **2.1** : Tracer les conversions dans `chainModifications`, depuis les annotations TAST jusqu'à `boxGoExprImpl`, `unboxGoExpr` et aux signatures des fonctions locales dans `CodeGen.purs`.
- [ ] **2.2** : Préserver les types fermés des records dans les paramètres et retours des closures typées, en lien avec le chantier 5. Ne pas limiter la correction à l'annotation du bloc de retour.
- [ ] **2.3** : Éliminer les conversions aller-retour équivalentes et conserver les accès directs aux champs dans les maillons de State.
- [ ] **2.4** : Pour les frontières qui doivent rester dynamiques, étudier l'usage des `RecordDict1` à `RecordDict5` déjà disponibles, afin d'éviter les slices du `RecordDict` général lorsque la forme est connue.
- [ ] **2.5** : Valider sur le chemin appelé par `runManyTimes` la réduction des `RecordDict`, `RecordGet` et allocations, avec résultats inchangés. Établir une nouvelle référence mesurée plutôt que conserver 108 μs comme cible fixe.

### 3. Méthodes typées dans les dictionnaires de classes

Le dictionnaire `Monoidish Int` conserve une méthode en `Value`, mais le benchmark actuel de polymorphisme a déjà éliminé ce dictionnaire de sa boucle. Le `Rebox` de `Get_Test_Polymorphism_intMonoidish` appartient à un chemin sans appel trouvé dans le dossier généré examiné : sa disparition seule ne démontrerait aucun gain sur les dix millions d'itérations.

- [ ] **3.1** : Identifier les pertes de types des champs fonctionnels et des superclasses entre `Main.classDeclsFields`, PBO et `CodeGen.structFieldGoType`. La substitution profonde de PBO traverse déjà les types `Func`, records et ADT ; localiser la perte effective plutôt que réimplémenter cette substitution.
- [ ] **3.2** : Conserver les signatures spécialisées des méthodes, par exemple `func(int64, int64) int64` ou leur équivalent curryfié, en s'appuyant sur les closures typées du chantier 5.
- [ ] **3.3** : Limiter les adaptateurs et `Rebox` aux véritables frontières génériques ; conserver une représentation compatible avec les accès runtime/FFI qui subsistent.
- [ ] **3.4** : Ajouter au harnais un cas où un dictionnaire spécialisé circule réellement comme valeur et où sa méthode reste appelée. Vérifier le chemin généré, les résultats et les allocations, puis mesurer. Garder `Test_Polymorphism` comme contrôle de la spécialisation déjà acquise.

### 5. Closures typées, puis élimination des thunks

Deux étapes distinctes : préserver les signatures des fonctions de première classe, puis étudier si certaines closures peuvent disparaître. `CodeGen.exprTypeToGoType` retombe actuellement sur `TypeValue` pour les types flèches, alors que des appels directs et fonctions locales bénéficient déjà de signatures natives.

`Test_LazyEvaluation` définit son propre `newtype Lazy a = Lazy (Unit -> a)`, sans mémoïsation. Le vrai `Data.Lazy` passe par une FFI avec `sync.Once`. Main exclut actuellement les FFI de la monomorphisation : améliorer CodeGen seul ne supprimera pas toutes ces frontières.

- [ ] **5.1** : Préserver les types flèches lors de leur conversion en types Go, y compris dans les champs de structures, paramètres et valeurs de retour.
- [ ] **5.2** : Émettre et appeler des closures natives typées lorsque leur signature est connue ; générer les adaptateurs vers `Value` uniquement aux frontières qui l'exigent. Couvrir capture, application partielle et retour de fonction.
- [ ] **5.3** : Comparer le chemin Lazy généré et la FFI à closures natives, en mesurant appels dynamiques et allocations. Traiter séparément les FFI de `Data.Lazy` et leur éventuelle spécialisation. Examiner aussi `forceEscape` et la durée de vie des captures ; sa suppression doit respecter les invariants de représentation `unsafe` du runtime.
- [ ] **5.4** : Étudier une analyse de demande/stricte évaluation pour supprimer les thunks lorsque c'est sémantiquement valide. Préserver les branches non demandées, la terminaison et le partage/mémoïsation de `Data.Lazy` ; le caractère déterministe seul ne suffit pas.
- [x] **5.4a** : Première fusion producteur/consommateur dans `Gopurs.ThunkFusion`, avant la TCO. Reconnaître une récursion construisant un accumulateur `Unit -> Int`, chaque maillon demandant son prédécesseur exactement une fois et inconditionnellement, avec uniquement des opérations entières totales autorisées. Remplacer les consommations immédiates à seed totale par une fonction auxiliaire à accumulateur `Int`. Conserver le producteur original pour les autres usages ; les appels inconnus, effets, divisions et demandes conditionnelles restent exclus. Cette passe ne dépend ni du nom du benchmark ni de ses constantes, et ne modifie pas `Data.Lazy`.
- [ ] **5.5** : Étudier séparément le calcul constant, la simplification des récurrences et la sortie des calculs invariants des boucles. Le cheatcode remplace toute la chaîne de mille thunks par `acc += 1000` : la suppression des closures seule peut laisser un million d'additions au lieu de mille.
- [ ] **5.6** : Valider séparément les gains du typage, de l'élimination des thunks et de la simplification du calcul. Les anciens chiffres de 17,3 ms, 11 ms et 0,25 ms sont des références historiques non revalidées par l'audit, pas des garanties ni des critères de réussite généraux.

**Mesures de 5.4a, le 7 septembre 2026 :** sur Apple M4 Pro, sans PGO, `GOGC=800`, la commande complète `./bin/go/run -c` passe. Lazy passe de **16,814 ms à 0,229 ms** au minimum de dix mesures (environ **73×**), avec le résultat `1000000` inchangé. Les 14 sorties de benchmarks correspondent à la référence ; seul `Test_LazyEvaluation.go` change parmi les fichiers Go générés du dossier `purescript`. Une mesure Go séparée, sur trois répétitions d'une seconde et avec résultat consommé, donne une médiane de **16,774 ms contre 0,244 ms** (environ **69×**), **1 000 000 contre 0 allocation**, et **32 Mo contre 0 octet alloué** par calcul. Ces octets concernent les allocations du calcul, pas la mémoire totale du processus.

Le programme PureScript du benchmark et le runtime sont inchangés : la passe enlève les closures intermédiaires, tout en conservant le million d'additions. Les tests ciblés sont décrits dans [`tests/thunk-fusion/README.md`](tests/thunk-fusion/README.md) : profondeurs et seeds variables, captures, ordre non commutatif, paramètres simultanés, collisions de noms Go, closures échappées, demandes conditionnelles et effets. Le typage général des closures (5.1–5.3), la mémoïsation de `Data.Lazy` et la simplification algébrique (5.5) restent des chantiers distincts.

**Périmètre sémantique :** les blocs locaux `LetRec` sont entièrement exclus de la fusion. Une annotation `Int` ne suffit pas à prouver qu'une valeur récursive est déjà initialisée ; cette barrière empêche d'avancer sa lecture et de dépendre du contrôle d'initialisation actuel du backend. Un test direct sur l'IR vérifie cette exclusion, ainsi que le maintien de l'optimisation pour les paramètres ordinaires déjà évalués.

### 6. Tableaux natifs et suppression des copies de conversion

Le chemin exécuté de `Test_ArrayOps.sumEvens` convertit le tableau filtré de `[]Value` vers `[]int64`, puis de nouveau vers `[]Value`, avant un `foldl` utilisant `Apply2(intAdd, ...)`. Les intrinsics de tableaux dans `CodeGen.purs` continuent à boxer leurs arguments et à appeler leurs fonctions via le runtime.

- [ ] **6.1** : Conserver les types d'éléments à travers les intrinsics `range`, `map`, `filter` et `foldl`, ainsi qu'aux frontières FFI concernées.
- [ ] **6.2** : Supprimer les conversions aller-retour inutiles et générer des appels typés ou des opérations primitives dans les boucles.
- [ ] **6.3** : Adapter la capacité des tableaux résultats lorsqu'une borne utile est connue, en mesurant le compromis entre réallocations et mémoire réservée.
- [ ] **6.4** : Valider que `Test_ArrayOps.sumEvens` ne reconstruit plus des tableaux pour changer uniquement de représentation ; mesurer allocations, octets alloués et temps.

### 7. Fusion des producteurs et consommateurs

Même typé et converti en boucles, `Test_ListOps.sumEvens` construit une liste initiale, une liste filtrée puis la parcourt pour sommer. Éviter ces structures intermédiaires demande une fusion des parcours, au-delà du déboxing.

- [ ] **7.1** : Identifier les compositions producteur/filtre/réduction dont les structures intermédiaires ne sont pas utilisées ailleurs.
- [ ] **7.2** : Fusionner progressivement ces compositions en conservant résultats, ordre d'évaluation observable et comportement de terminaison. Appliquer ensuite les mêmes principes aux tableaux.
- [ ] **7.3** : Valider sur ListOps et ArrayOps la disparition des structures intermédiaires visées et la baisse des allocations. Ne pas prendre un changement d'algorithme propre au cheatcode comme preuve d'une optimisation générale.

### 8. Représentation mémoire des ADT et allocations

RBTree et les listes bénéficient déjà de pointeurs typés. Certains ADT à plusieurs constructeurs non vides, comme AstTree, conservent des enfants en `Value`. Par ailleurs, CodeGen ajoute `Rc uint32` à tous les constructeurs, sans accès `.Rc` ni appel `IncRef` trouvé dans le dossier `purescript` généré examiné.

- [ ] **8.1** : Mesurer les tailles des structures et les allocations des listes/arbres, puis étudier des représentations typées des sommes récursives encore génériques.
- [ ] **8.2** : Vérifier l'utilisation réelle de `Rc`, y compris dans le runtime et les FFI, puis éviter le champ et son padding là où aucun mécanisme de réutilisation/comptage n'en dépend.
- [ ] **8.3** : Examiner les `Rebox` de structures récursives : certains recopient toute une liste. Éviter ces frontières sur les chemins fréquents plutôt que traiter toute coercion comme une opération scalaire.
- [ ] **8.4** : Valider les gains en mémoire, allocations et temps sur les chemins effectivement utilisés, en conservant les conventions de représentation partagées avec le runtime.

## Performances du compilateur

### 9. Compilation incrémentale et coût de la monomorphisation

Dans les deux chemins de construction de `Main.purs`, `onSkipModule` impose `res <- pure Nothing`, alors que des caches sont encore écrits. La réutilisation des modules est donc désactivée. La collecte transitive des spécialisations effectue aussi des parcours répétés à point fixe ; leur poids reste à profiler.

- [ ] **9.1** : Mesurer séparément chargement, collecte/spécialisation, optimisation PBO, CodeGen, FFI et compilation Go.
- [ ] **9.2** : Restaurer une réutilisation du cache avec invalidation sur les entrées CoreFn/TAST, versions/options du compilateur, directives, FFI et dépendances de spécialisation/inlining. Une nouvelle instanciation dans un appelant peut invalider le module définissant la fonction, même si son source n'a pas changé ; un contrôle des dates seul ne suffit pas.
- [ ] **9.3** : Restaurer les métadonnées nécessaires à l'optimisation intermodules et éviter les écritures d'outputs inchangés.
- [ ] **9.4** : Établir un protocole incrémental ciblé distinct de la validation propre `./bin/go/run -c`, qui efface les outputs. Vérifier les cas sans changement, modification d'un appelant, changement de spécialisation et modification FFI, en comparant leurs résultats à la construction propre.
- [ ] **9.5** : Si le profil le justifie, remplacer les rescans des instanciations/appelants dans `transitiveCollect` par une file des nouveaux travaux et une mémoïsation des résultats.

### 10. Extraction FFI persistante et mise en cache

`Gopurs/FfiSupport.js` lance synchroniquement un nouveau processus Node, charge et instancie le parser WASM pour chaque module FFI traité par Main.

- [ ] **10.1** : Mesurer la part du démarrage des processus et du chargement WASM dans le temps de compilation.
- [ ] **10.2** : Réutiliser un parser persistant ou natif, puis mettre en cache l'AST FFI selon le contenu, la version du parser et les options pertinentes.
- [ ] **10.3** : Préserver les diagnostics d'erreur et la cohérence des wrappers ; comparer construction propre et réutilisation sans changement.

### 4. Auto-hébergement : compiler gopurs avec gopurs

Obtenir un compilateur natif est un objectif de portabilité et de parité fonctionnelle, avec un bénéfice de performance à mesurer après les chantiers 9 et 10. Cela n'accélère pas automatiquement le code des programmes produits.

- [x] **4.1** : Première tentative de compilation du projet gopurs avec la version JS de gopurs lancée. Ce jalon historique ne signifie pas qu'un binaire natif fonctionnel a été obtenu.
- [ ] **4.2** : Porter les FFI de l'écosystème et celles du compilateur/PBO : filesystem et processus, `v8.serialize`/`deserialize` et restauration des prototypes du cache, `WeakMap` des variables libres, échappement des chaînes et extraction FFI Node/WASM.
- [ ] **4.3** : Obtenir un binaire utilisable et vérifier la parité des diagnostics, résultats et outputs avec le compilateur JS, notamment sur `altbak.pub`.
- [ ] **4.4** : Comparer temps et mémoire du compilateur natif et du compilateur JS à entrées, options et état de cache identiques. Documenter le gain constaté, sans présumer qu'il sera important.

## Validation et interprétation des mesures

- [ ] Conserver la commande de validation propre indiquée en tête. Un protocole incrémental devra être défini explicitement pour le chantier 9 ; le nettoyage `-c` ne permet pas de mesurer les hits du cache.
- [ ] Séparer temps d'exécution du programme, temps du backend et compilation Go ; distinguer constructions propres et incrémentales.
- [ ] Fixer machine, entrées, versions, options PGO et `GOGC` pour chaque comparaison. Le runner examiné utilise `GOGC=800` hors PGO et `GOGC=1000` avec PGO : les résultats doivent préciser ces paramètres.
- [ ] Compléter le minimum de dix mesures actuel par la distribution des temps, les allocations, les octets alloués et la mémoire maximale, en précisant le périmètre mesuré.
- [ ] Vérifier les résultats sur plusieurs entrées et inspecter le chemin réellement exécuté. La disparition textuelle de `Value`, `Apply` ou `Rebox` ne démontre pas à elle seule une accélération.
- [ ] Séparer FFI comparable et cheatcode. Le cheatcode Lazy simplifie le calcul ; celui de StateMonad ignore son argument `limit` et utilise une taille fixe. Ils ne constituent pas des équivalents généraux des fonctions PureScript pour toutes les entrées.
