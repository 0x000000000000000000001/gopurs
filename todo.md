# Résolution des Bugs AOT (Gopurs)

## Prochaines étapes (Baby steps)

### 1. Optimisation des appels FFI (Suppression de l'overhead de boxing)
Actuellement, les appels aux fonctions externes (FFI) sont traités dynamiquement via `gopurs_runtime.Apply(...)` et forcent le boxing des arguments et de la valeur de retour (via `gopurs_runtime.Int` ou les wrappers `_Gopurs_...`). Étant donné que le type exact de la fonction importée est connu dans le TAST (ex: `Func [Int] Int`), le compilateur devrait émettre un appel statique natif direct.

- [ ] *Action* : Shunter le dispatcher `Apply` pour les appels FFI dont la signature est connue.
  - [ ] **Baby Step 1.1** : Identifier dans `CodeGen.purs` où les applications de fonctions FFI sont générées. Intercepter l'appel si la cible est une référence FFI.
  - [ ] **Baby Step 1.2** : Extraire l'arité et les types concrets de la fonction depuis `dataDecls` ou `globalTypes`.
  - [ ] **Baby Step 1.3** : Modifier l'émission Go pour bypasser `gopurs_runtime.Apply` et générer un appel direct natif (ex: `Test_AckermannFFI_RunAckermannFFI(int64(...))`).
  - [ ] **Baby Step 1.4** : Adapter la génération du wrapper FFI (`_ffi.go`) et valider les gains de performance empiriquement sur les benchmarks `*FFI`.
