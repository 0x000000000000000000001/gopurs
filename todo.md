# Résolution des Bugs AOT (Gopurs)

## Prochaines étapes (Baby steps)

### 1. Typage strict des boucles internes inlinées (Élimination des Rebox)
L'optimiseur inline physiquement certaines boucles locales (ex: la boucle `go` de `reverse` au sein de `filter`). Cependant, la variable de type d'origine de la boucle inlinée (ex: `a`) échappe à la substitution lors de la spécialisation (ex: `a -> int64`). Cela force un fallback vers `gopurs_runtime.Value` et génère des appels coûteux à `Rebox_` (deep-copy O(N)) à chaque itération.

- [x] *Action* : Empêcher la fuite de typage polymorphe lors de l'inlining des fonctions locales.
  - [x] **Baby Step 1.1** : Tracer l'évaluation par `collectLocalExpr` et `rewriteExpr` dans `Monomorphize.purs` pour repérer la perte de substitution sur les `LetRec` inlinés.
  - [x] **Baby Step 1.2** : Forcer l'unification ou propager les types concrets aux variables orphelines pour que les boucles internes conservent leur typage strict (`int64`, etc.).
  - [x] **Baby Step 1.3** : Valider dans `Test_Primes.go` et `Test_ListOps.go` la disparition totale des appels `Rebox_`.

### 1.5 Monomorphisation Profonde (Deep Monomorphization)
Actuellement, la passe de monomorphisation est superficielle. L'analyse des instanciations (`analyzeInstantiations`) ne descend pas dans le corps des futures fonctions spécialisées pour découvrir les instanciations en cascade (ex: `filter__Int64` appelle `reverse` avec `Int64`, ce qui nécessite de générer `reverse__Int64`). Ce fallback sur les appels génériques engendre des Deep Copies O(N) catastrophiques (`Rebox_`).

- [x] *Action* : Rendre la découverte des instanciations itérative ou récursive (jusqu'à un point fixe).
  - [x] **Baby Step 1.5.1** : Analyser le fonctionnement de `analyzeInstantiations` et son dictionnaire `instMap` actuel.
  - [x] **Baby Step 1.5.2** : Modifier l'algorithme pour qu'après avoir substitué les types d'une fonction, il réanalyse son corps pour découvrir les appels polymorphes internes.
  - [x] **Baby Step 1.5.3** : Mettre à jour les références (ast substitution) dans le corps spécialisé pour pointer vers les nouvelles versions `__...` (ex: `reverse__Int64`).
  - [x] **Baby Step 1.5.4** : Vérifier sur `Prime Sieve` la disparition du `Rebox_` final et le retour à un benchmark < 60 μs.

### 1.6 Correction de la Monomorphisation Profonde (Deep Monomorphization Fix)
Bien que les étapes 1.5 aient été implémentées, la boucle interne inlinée de `reverse` au sein de `filter` (dans `Test_Primes.go`) s'exécute toujours avec le type `gopurs_runtime.Value` (Any) au lieu de `int64`. Cela provoque un appel inutile à `Rebox_` de complexité O(N).
Le diagnostic a montré que :
- Soit `Monomorphize.purs` (`rewriteExpr` / `monomorphizeBindLocal`) échoue à propager la substitution de type (`a -> Int64`) jusque dans les annotations `Binding Ann` des boucles locales `LetRec`.
- Soit la fonction `monomorphizeExpr` n'arrive pas à substituer `ExprTypeApp (ExprVar "reverse") Int64` par `reverse__Int64` parce que `genericType` est évalué à `Any` (dépourvu de variable de type, la condition `hasTypeVariables` l'ignore donc et saute le lookup dans `instMap`).
- En conséquence, le compilateur émet un appel à la version polymorphe générique `reverse__1084068251` (List Any -> List Any), qui est inlinée plus tard par `Semantics.purs`.

- [ ] *Action* : Diagnostiquer et corriger le lookup de l'AST lors de la monomorphisation finale pour s'assurer que les appels génériques internes sont effectivement liés à leurs clones spécialisés.
  - [x] **Baby Step 1.6.1** : Tracer `monomorphizeExpr` (cas `ExprTypeApp`) lors de la passe finale (`buildSpecializedBindings`) pour vérifier pourquoi `instMap` n'est pas interrogé ou pourquoi la substitution échoue (vérifier la valeur de `genericType` et `instType`).
  - [ ] **Baby Step 1.6.2** : Corriger la propagation ou la condition de lookup (`if not (hasTypeVariables genericType)...`) pour que l'AST reconstruise bien un appel vers la fonction manglée (ex: `reverse__Int64`).
  - [ ] **Baby Step 1.6.3** : Vérifier que `rewriteExpr` et `monomorphizeBindLocal` appliquent correctement la substitution de type dans les annotations des `LetRec`.
  - [ ] **Baby Step 1.6.4** : Regénérer le code et valider dans `Test_Primes.go` que la sous-boucle inlinée de `reverse` opère désormais sur `*Constructor_Test_Primes_Cons[int64]` (et disparition finale du `Rebox_`). Test empirique de perf: < 60 μs.
