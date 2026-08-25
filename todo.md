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

- [ ] **Step 1 (Monomorphisation stricte des ADTs)** :
  - **Action** : Le TAST v3 fournissant les `TypeApp`, `gopurs` doit générer des structures génériques natives Go 1.18 (`type Cons[T any] struct { V0 T; V1 *Cons[T] }`) ou spécialisées (`type Cons_Int64`).
  - **Résultat attendu** : Éradication totale de `gopurs_runtime.Value` pour les payloads des ADTs. Fini les pointeurs dynamiques pour des primitives.

- [ ] **Step 2 (Pattern Worker/Wrapper pour les fonctions)** :
  - **Action** : Les fonctions d'ordre supérieur (comme `foldl`) ou les boucles récursives doivent être scindées en deux :
    - *Worker* : Le moteur interne avec une signature 100% native (ex: `func foldl_worker(acc int64, lst *Cons[int64]) int64`), exécutant la boucle de manière itérative sur la stack sans allouer de `Value`.
    - *Wrapper* : Une surcouche maintenant la signature d'origine acceptant/retournant des `gopurs_runtime.Value` pour garantir la compatibilité si la fonction est exportée ou passée comme closure dynamique.
  - **Bénéfice** : Zéro-cost abstraction sur le chemin critique, tout en gardant une API compatible avec le système de types dynamiques du runtime. 

Tu peux tester si tout roule dans altbak.pub avec `bin/go/run -c`