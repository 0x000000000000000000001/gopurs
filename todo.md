# Gopurs — Priorités utiles

> **Benchmark : `./bin/go/run -c` dans `altbak.pub`. Ne pas utiliser `./bin/run go` sans argument : il exécute tous les backends.**

Les gains ci-dessous viennent de prototypes Go, pas encore du compilateur. Intégrer et mesurer chaque point séparément. Les régressions doivent rester dans `tests/passing`, avec leurs snapshots Go ; pas de projet de test séparé ni de script de test `.mjs`.

## 1. Regrouper les lambdas adjacentes avec `Func2`

Prototype Church : **1,83 → 0,47 ms**, **100 260 → 157 allocations**. Le calcul et les applications partielles sont conservés.

- [x] **1.1** Fixture `tests/passing/CurriedLambdas.purs` ajoutée : deux producteurs annotés, entrées variées et neuf assertions. Le snapshot initial confirmait les deux `Func` imbriqués après le PBO, avec reproduction à l'identique avant optimisation.
- [x] **1.2** Collecteur `collectCurriedAbs` ajouté dans `CodeGen.purs` : seuls les `Abs` adjacents, éventuellement séparés par `Typed`, sont regroupés. Ordre et niveaux des paramètres, annotations du corps terminal conservés.
- [x] **1.3** Collecteur branché dans `Abs`, avec réutilisation de `buildFunc`. Le type du corps tient compte des seuls paramètres collectés et reste fonctionnel si d'autres arguments subsistent. Boxing final et runtime d'application conservés. `./bin/test CurriedLambdas -c` passe ; snapshot mis à jour et validé avec trois paires de `Func` remplacées par `Func2`, corps des calculs identiques.
- [x] **1.4** Fixture étendue à 30 assertions : appels saturés, partielles réellement réutilisées via `Ref`, callback non linéaire, traces des effets et absence d'exécution anticipée. Arités trois et six couvertes, avec partielles avant et après cinq arguments ; snapshot confirmé en `Func3` et `Func5` retournant `Func`. `./bin/test CurriedLambdas -c` et l'exécution du JavaScript PureScript passent.
- [ ] **1.5** Ajouter les cas où la fusion doit s'arrêter : calcul ou `Let` entre deux lambdas, branche et récursion locale. Le collecteur ne doit franchir ni `LetRec`, `App`, effet, `UncurriedAbs` ou `TypeApp`. Vérifier les résultats et les frontières conservées dans les snapshots.
- [ ] **1.6** Faire passer la fixture puis la suite `passing`, relire les différences de snapshots et vérifier que Church génère bien les applications regroupées depuis le compilateur.
- [ ] **1.7** Relancer `./bin/go/run -c` et mesurer Church avant/après sur le Go régénéré : temps, octets et allocations par opération. Utiliser la même charge et les mêmes réglages, sans mesures concurrentes ; confirmer le gain avec `GOGC=800` et `GOGC=100`.

## 2. Spécialiser le boxing des entiers à la frontière FFI

Prototype Arrays, 900 éléments : **15,4 → 9,5 µs**, **1 095 → 15 allocations**, en remplaçant seulement `Box(int64)` par `Int(int64)` dans les wrappers de range et d'addition.

- [ ] **2.1** Ajouter une fixture dans `passing` avec une FFI Go retournant `int64`, `int`, `[]int64` et `[]int`. Couvrir zéro, valeurs négatives, limites supportées et tableaux vides ; inclure un retour dynamique pour protéger le fallback.
- [ ] **2.2** Dans `wrapReturn` de `CodeGen.purs`, émettre `gopurs_runtime.Int(v)` pour un type Go explicitement `int64`, et `gopurs_runtime.Int(int64(v))` pour `int`. Conserver `Box` pour `any`, les types inconnus et les types nommés non reconnus.
- [ ] **2.3** Réutiliser cette sélection pour les éléments des tableaux `[]int64` et `[]int`. Conserver la représentation `[]Value` et le traitement actuel des autres types d'éléments.
- [ ] **2.4** Vérifier les allers-retours FFI de la fixture et ses snapshots : `Int` sur les retours connus, fallback inchangé ailleurs. Ajouter un cas où le résultat est consommé via une fonction passée en argument.
- [ ] **2.5** Faire passer la fixture puis la suite `passing`. Régénérer altbak.pub et vérifier les wrappers de `Data.Array.rangeImpl` et `Data.Semiring.intAdd`, sans modifier leur logique ni les conversions d'entrée.
- [ ] **2.6** Relancer `./bin/go/run -c`, puis comparer Arrays à 900 et 90 000 éléments avec temps, octets et allocations par opération. Mesurer séparément de l'étape 1 et vérifier les autres résultats du benchmark.

Se limiter à ces deux changements : le pipeline Arrays entièrement typé, la fusion des opérations et la réutilisation des nœuds RBTree restent hors de cette liste. Reporter uniquement les gains mesurés après intégration dans le compilateur.
