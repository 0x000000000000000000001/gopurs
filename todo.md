# Roadmap Optimisation `gopurs` (TAST v3 / TypeApp)

L'objectif est d'atteindre les performances maximales du compilateur Go natif ("cheatcode"), en éliminant les reliquats d'allocations dynamiques (`gopurs_runtime.Value`) qui ralentissent l'exécution.

> **Preuve empirique (basée sur `altbak.pub`)** :
> En analysant l'output généré par `gopurs` (comme `Test_ListOps.go`), on constate que la spécialisation des types est incomplète malgré la puissance du TAST v3, car le générateur de code Go n'exploite pas encore les informations de type exactes :
> 
> 1. **Polymorphisme des ADTs (ex: ListOps - 44 µs vs 1 µs cheatcode)** : 
>    Dans `Test_ListOps.go`, le constructeur `Cons` est défini avec :
>    `type Constructor_Test_ListOps_Cons struct { Rc uint32; V0 gopurs_runtime.Value; V1 *Constructor_Test_ListOps_Cons }`
>    Le payload `V0` est contraint d'être un `gopurs_runtime.Value` (alloué sur le tas, `interface{}`) au lieu d'un `int64` natif, sollicitant massivement le Garbage Collector.
> 
> 2. **Surcharge de boxing (ex: Opérations mathématiques)** : 
>    Une simple soustraction est traduite par `v_3_loop = gopurs_runtime.Int((v_3.IntVal) - (1))`. Bien que l'opération se fasse de manière native, le résultat est immédiatement ré-alloué sur la heap sous la forme d'un `gopurs_runtime.Int`.

## État des lieux (Accompli) :
- [x] **Dictionnaires de Type Class (Monomorphisation)** : L'utilisation de `Apply2(Box(dict.V0), ...)` a été complètement éradiquée. Grâce à l'optimiseur centralisé branché sur le TAST v3, toutes les instances de Type Classes sont résolues statiquement avant la génération de code Go.
- [x] **Unification du pipeline de Build** : Migration vers `spago bundle` et interfaçage avec `purescript-backend-optimizer`.

## Prochaines étapes

- [x] **Step 1 (Monomorphisation stricte des ADTs)** :
  - **Action** : Le TAST v3 fournissant les `TypeApp`, `gopurs` doit générer des structures génériques natives Go 1.18 (`type Cons[T any] struct { V0 T; V1 *Cons[T] }`) ou spécialisées (`type Cons_Int64`).
  - **Résultat attendu** : Éradication totale de `gopurs_runtime.Value` pour les payloads des ADTs. Fini les pointeurs dynamiques pour des primitives. *(Accompli, mais a révélé le besoin vital de l'étape 2)*

- [x] **Step 2 (Pattern Worker/Wrapper pour les fonctions & Instanciation Locale)** :
  - **Problème Actuel** : Les boucles locales (`LetRec`) inlinées par l'optimiseur conservent leur type polymorphe d'origine (`TypeVar "a"`), ce qui provoque un crash mémoire (panic) lorsqu'elles tentent de lire un payload ADT généré nativement (ex: lire un `Value` à la place d'un `int64`).
  - [x] **Phase 1 (Instanciation locale des LetRec)** : Scrutateur de type dans `translateExprImpl_ LetRec`. Inférer le type concret (`int64`) depuis l'appel de la boucle (`GoCall`) pour instancier ses arguments au lieu d'utiliser aveuglément `TypeVar` (qui devient `Value`). *(Accompli : Le routage correct des types attendus via le TAST résout les panics d'exécution liés au mismatch `Value` vs type natif)*
  - [x] **Phase 2 (Workers Génériques Top-Level)** : Restauration des paramètres génériques natifs Go (Go 1.18+).
    - *Action* : Collecte des variables implicites (`TypeVar`) dans l'AST des types pour forcer la signature générique du Worker (`func Call_foldl[T_b any, T_a any]`). Mise à jour des appels directs (`mbDirectCall`) pour injecter explicitement la version opaque (`[gopurs_runtime.Value, ...]`) afin de satisfaire l'inférence du compilateur Go.
    - *Résultat* : Les fonctions génériques sont de retour dans le code Go sans faire planter le compilateur. L'échafaudage pour la Phase 3 est en place.

- [ ] **Step 3 (Monomorphisation Intelligente des Types d'Arguments)** :
  - [] **Action 2 (Générateur Go)** :
    - Amélioration de `exprTypeToGoType` pour préserver le type générique (ex: `T_a`) au lieu de le rabaisser en `Value`.
    - Adaptation du *Wrapper* avec `gopurs_runtime.AnyToValue` et `ValueToAny` pour "déballer" et "remballer" aux frontières d'appel du Worker.
  - **Bénéfice ultime** : Le Worker n'effectue plus aucune conversion dynamique et reçoit des structures 100% natives. Les appels directs depuis PureScript tireront parti de la "monomorphisation par stenciling" de Go (ex: appel direct de la version `..._int` hyper rapide). Zéro-cost abstraction sur le chemin critique !

Tu peux tester si tout roule dans altbak.pub avec `bin/go/run -c`
