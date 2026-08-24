# Roadmap Optimisation `gopurs` (TAST v2 / TypeApp)

L'objectif est d'atteindre les performances maximales du compilateur Go natif, en éliminant les reliquats d'allocations dynamiques (`gopurs_runtime.Value`) qui ralentissent l'exécution.

> **Preuve empirique (basée sur `altbak.pub`)** :
> En analysant l'output généré par `gopurs` (`Test_ListOps.go` et `Test_Polymorphism.go`), on constate que la monomorphisation actuelle est incomplète à cause de l'absence passée de `TypeApp` :
> 
> 1. **Polymorphisme des ADTs (ex: ListOps - 44 µs vs 1 µs cheatcode)** : 
>    Dans `Test_ListOps.go`, le constructeur `Cons` est défini comme :
>    `type Constructor_Test_ListOps_Cons struct { Rc uint32; V0 gopurs_runtime.Value; V1 *Constructor_Test_ListOps_Cons }`
>    Le payload `V0` est contraint d'être un `gopurs_runtime.Value` (alloué sur le tas, interface{}) au lieu d'un `int64` natif, provoquant 44x plus de temps d'exécution que le code natif (garbage collector très sollicité).
> 
> 2. **Fonctions d'ordre supérieur (ex: foldl)** : 
>    Dans `Test_ListOps.go`, la signature de `foldl` prend et retourne des `gopurs_runtime.Value` même si le type de l'accumulateur est connu dynamiquement.
>    `func Call_Test_ListOps_foldl(v_0_loop gopurs_runtime.Value, ...) gopurs_runtime.Value`
> 
> 3. **Dictionnaires de Type Class (ex: Polymorphism - 2864 µs)** :
>    Dans `Test_Polymorphism.go`, l'application du dictionnaire se fait via l'API dynamique du runtime, qui boxe inutilement et fait une invocation indirecte : 
>    `v1_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_5, ...)`

## Prochaines étapes

- [ ] **Step 1 (Monomorphisation stricte des ADTs)** :
  - **Action** : Utiliser `TypeApp` pour instancier des structures génériques natives Go (`type Cons[T any] struct { V0 T; V1 *Cons[T] }`) ou des structures spécialisées (`type Cons_Int64 struct { ... }`).
  - **Résultat attendu** : Éradication totale de `gopurs_runtime.Value` pour les payloads des ADTs, éliminant les allocations sur le tas pour les primitives comme `int64`.

- [ ] **Step 2 (Monomorphisation des signatures de fonctions)** :
  - **Action** : Les fonctions d'ordre supérieur et les fonctions monomorphisées via `TypeApp` doivent avoir des signatures Go 100% natives sans `Value`.
  - **Exemple** : `foldl` deviendra `func foldl_Int(f func(int64, int64) int64, acc int64, lst *Cons_Int64) int64`.

- [ ] **Step 3 (Dictionnaires statiques)** :
  - **Action** : Les instances de type class étant maintenant explicitement résolues par `TypeApp` et `Monomorphize.purs`, le code Go peut appeler directement les fonctions spécifiques (ex: `+`) ou passer des structs dont les champs sont des pointeurs de fonctions natifs, sans passer par `gopurs_runtime.Apply2`.
