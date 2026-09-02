# Résolution des Bugs AOT (Gopurs)

## Prochaines étapes (Baby steps)

### 1. Typage strict des boucles internes inlinées (Élimination des Rebox)
L'optimiseur inline physiquement certaines boucles locales (ex: la boucle `go` de `reverse` au sein de `filter`). Cependant, la variable de type d'origine de la boucle inlinée (ex: `a`) échappe à la substitution lors de la spécialisation (ex: `a -> int64`). Cela force un fallback vers `gopurs_runtime.Value` et génère des appels coûteux à `Rebox_` (deep-copy O(N)) à chaque itération.

- [ ] *Action* : Empêcher la fuite de typage polymorphe lors de l'inlining des fonctions locales.
  - [ ] **Baby Step 1.1** : Tracer l'évaluation par `collectLocalExpr` et `rewriteExpr` dans `Monomorphize.purs` pour repérer la perte de substitution sur les `LetRec` inlinés.
  - [ ] **Baby Step 1.2** : Forcer l'unification ou propager les types concrets aux variables orphelines pour que les boucles internes conservent leur typage strict (`int64`, etc.).
  - [ ] **Baby Step 1.3** : Valider dans `Test_Primes.go` et `Test_ListOps.go` la disparition totale des appels `Rebox_`.

### 2. Optimisation des appels FFI (Suppression de l'overhead de boxing)
Actuellement, les appels aux fonctions externes (FFI) sont traités dynamiquement via `gopurs_runtime.Apply(...)` et forcent le boxing des arguments et de la valeur de retour (via `gopurs_runtime.Int` ou les wrappers `_Gopurs_...`). Étant donné que le type exact de la fonction importée est connu dans le TAST (ex: `Func [Int] Int`), le compilateur devrait émettre un appel statique natif direct.

- [ ] *Action* : Shunter le dispatcher `Apply` pour les appels FFI dont la signature est connue.
  - [ ] **Baby Step 2.1** : Identifier dans `CodeGen.purs` où les applications de fonctions FFI sont générées. Intercepter l'appel si la cible est une référence FFI.
  - [ ] **Baby Step 2.2** : Extraire l'arité et les types concrets de la fonction depuis `dataDecls` ou `globalTypes`.
  - [ ] **Baby Step 2.3** : Modifier l'émission Go pour bypasser `gopurs_runtime.Apply` et générer un appel direct natif (ex: `Test_AckermannFFI_RunAckermannFFI(int64(...))`).
  - [ ] **Baby Step 2.4** : Adapter la génération du wrapper FFI (`_ffi.go`) et valider les gains de performance empiriquement sur les benchmarks `*FFI`.
