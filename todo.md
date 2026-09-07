# Gopurs — Priorités utiles

> **Benchmark : `./bin/go/run -c` dans `altbak.pub`. Ne pas utiliser `./bin/run go` sans argument : il exécute tous les backends.**

Les chiffres marqués « Prototype » viennent de prototypes Go ; les mesures après intégration du regroupement figurent en 1.7. Intégrer et mesurer chaque point séparément. Les régressions doivent rester dans `tests/passing`, avec leurs snapshots Go ; pas de projet de test séparé ni de script de test `.mjs`.

## 2. Spécialiser le boxing des entiers à la frontière FFI

Prototype Arrays, 900 éléments : **15,4 → 9,5 µs**, **1 095 → 15 allocations**, en remplaçant seulement `Box(int64)` par `Int(int64)` dans les wrappers de range et d'addition.

- [ ] **2.1** Ajouter une fixture dans `passing` avec une FFI Go retournant `int64`, `int`, `[]int64` et `[]int`. Couvrir zéro, valeurs négatives, limites supportées et tableaux vides ; inclure un retour dynamique pour protéger le fallback.
- [ ] **2.2** Dans `wrapReturn` de `CodeGen.purs`, émettre `gopurs_runtime.Int(v)` pour un type Go explicitement `int64`, et `gopurs_runtime.Int(int64(v))` pour `int`. Conserver `Box` pour `any`, les types inconnus et les types nommés non reconnus.
- [ ] **2.3** Réutiliser cette sélection pour les éléments des tableaux `[]int64` et `[]int`. Conserver la représentation `[]Value` et le traitement actuel des autres types d'éléments.
- [ ] **2.4** Vérifier les allers-retours FFI de la fixture et ses snapshots : `Int` sur les retours connus, fallback inchangé ailleurs. Ajouter un cas où le résultat est consommé via une fonction passée en argument.
- [ ] **2.5** Faire passer la fixture seulement (pas la suite entière `passing`). Régénérer altbak.pub et vérifier les wrappers de `Data.Array.rangeImpl` et `Data.Semiring.intAdd`, sans modifier leur logique ni les conversions d'entrée.
- [ ] **2.6** Relancer `./bin/go/run -c`, puis comparer Arrays à 900 et 90 000 éléments avec temps, octets et allocations par opération. Mesurer séparément de l'étape 1 et vérifier les autres résultats du benchmark.

Se limiter à ces deux changements : le pipeline Arrays entièrement typé, la fusion des opérations et la réutilisation des nœuds RBTree restent hors de cette liste. Reporter uniquement les gains mesurés après intégration dans le compilateur.
