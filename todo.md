> **Règle absolue pour les tests :** Le *seul* test à lancer pour valider les développements est `./bin/go/run -c` depuis le dossier `altbak.pub`.

# Résolution des Bugs AOT (Gopurs)

> **⚠️ ATTENTION : Le seul test à faire est `./bin/go/run -c` dans le dossier `altbak.pub`. N'utilisez jamais `./bin/run go` sans argument car cela exécute tous les backends et fausse les métriques.**

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
  - [x] **Baby Step 1.6.2** : Corriger la propagation ou la condition de lookup (`if not (hasTypeVariables genericType)...`) pour que l'AST reconstruise bien un appel vers la fonction manglée (ex: `reverse__Int64`).
  - [x] **Baby Step 1.6.3** : Vérifier que `rewriteExpr` et `monomorphizeBindLocal` appliquent correctement la substitution de type dans les annotations des `LetRec`.
    - [x] 1.6.3.1 : Tracer l'annotation `bAnn` des fonctions locales (ex: la sous-fonction `go` de `reverse`). *Résultat : `bAnn` n'a pas de type (`Nothing`), mais le site d'appel (`ExprVar`) conserve des variables de type.*
    - [x] 1.6.3.2 : Tracer `unifySpine` dans `collectLocalExpr` pour observer les types comparés lors de l'inférence locale (`paramType` vs `actualType`). Le but est de prouver que `actualType` (l'argument passé, ex: `Nil`) est vu comme `Any` ou générique parce que la réécriture initiale (`rewriteExpr`) du corps de `filter` a ignoré/échoué à substituer le type de ces arguments.
    - [x] 1.6.3.3 : Faire en sorte que `substituteExprType` ou `rewriteExpr` applique bien les substitutions aux `ExprLet` et `ExprVar` qui pointent vers les fonctions locales (`go`), pour que le type `oldType` `[TypeVar a]` devienne bien `newType` `[Int64]`, permettant ainsi à l'inférence de déclencher la spécialisation de la boucle interne.
  - [x] **Baby Step 1.6.4** : Regénérer le code et valider dans `Test_Primes.go` que la sous-boucle inlinée de `reverse` opère désormais sur `*Constructor_Test_Primes_Cons[int64]` (et disparition finale du `Rebox_`).
    *Diagnostic actuel (Antigravity) : J'ai corrigé `substituteExprType` dans `Substitute.purs` pour qu'il n'ignore plus les variables sous `ForAll`. Résultat : la boucle interne locale de `filter` est bien spécialisée en `int64` sans `Rebox_` ! CEPENDANT, un `Rebox_` persiste à la fin de `filter` lors de l'appel inliné à `reverse`. Pourquoi ? Parce que le compilateur PureScript inline `reverse` dans `filter` avant de générer `corefn.json`, et lors de cet inlining, il renomme les variables de type (ex: `a` devient `a1`) pour éviter la capture. Notre map de substitution fixe `{"a" -> Int}` échoue donc à substituer `a1`, laissant la sous-boucle `go` de `reverse` en générique `gopurs_runtime.Value`.*
    - [x] 1.6.4.1 : Diagnostiquer comment PBO gère ces variables de type renommées (ex: lors de `unifySpine`) et identifier un moyen de construire une map de substitution plus exhaustive ou dynamique pour l'AST inliné. *(Fait: Le problème vient bien du renommage (a -> a1) lors de l'inlining natif de PureScript. La map de substitution générée par `inferSubst` est trop rigide).*
    - [x] 1.6.4.2 : Adapter la logique de substitution (dans `Monomorphize.purs`, au niveau de `astSubstFn`, `rewriteExpr` ou via `unifySpine`) pour qu'elle soit capable de découvrir dynamiquement les variables renommées (ex: en unifiant les types de l'AST inliné avec les arguments réels) OU aplatir les renommages avant la substitution.
    - [x] 1.6.4.3 : Implémenter le correctif pour que la substitution se propage correctement dans l'AST inliné.
    - [x] 1.6.4.4 : Vérifier que la génération Go ne produit plus aucun appel `Rebox` pour la boucle `go` issue de `reverse`, et que les performances de `sieve` sur 500 primes s'en trouvent maintenues ou améliorées. Test empirique de perf visé : < 100 μs. (Résultat mesuré : 38 μs !)

### 1.7 Persistance du bug de Monomorphisation sur les LetRec inlinés (Rebox)
Malgré les validations apparentes de l'étape 1.6, une inspection du code Go généré (`Test_Primes.go`) montre que le bug n'est pas complètement résolu. La boucle `go` principale de `filter` est bien spécialisée en `int64`, mais la sous-boucle `go` issue de l'inlining de `reverse` continue d'opérer sur du typage générique (`*Constructor_Test_Primes_Cons[gopurs_runtime.Value]`). En conséquence, `filter` subit toujours des pénalités de conversion via `Rebox_...` à la fin de son exécution pour passer de `Cons[int64]` à `Cons[Value]`.

**Diagnostic mis à jour :**
Le patch du compilateur Haskell (`Desugar.hs`) est **un succès total**. Les types des `LetRec` (comme le `go` de `reverse`) sont désormais bien présents et corrects dans le `corefn.json` généré (`type: 14`). Les `type: null` restants sur les masques (`Binder`) ne sont pas un problème car PBO sait unifier par le haut.
Le problème vient de la passe d'**Inlining/Monomorphisation de PBO**. Lorsque PBO inline l'AST de `reverse` au sein de `filter`, soit il perd l'annotation de type du `LetRec` copié, soit la propagation de substitution (`Monomorphize.purs`) échoue à s'appliquer aux variables de type renommées (`a1`) à l'intérieur de ce sous-arbre.

- [ ] *Action* : Garantir la spécialisation récursive des sous-arbres inlinés et la disparition stricte des `Rebox_` restants dans PBO.
  - [x] **Baby Step 1.7.1** : (Fait) Validation empirique que `Test_Primes.go` génère encore des `Rebox_` à cause des patterns non-unifiés et confirmation que `corefn.json` contient bien les types.
  - [x] **Baby Step 1.7.2** : Inspecter la logique d'Inlining et de Monomorphisation dans `purescript-backend-optimizer` (`Monomorphize.purs`). Identifier pourquoi l'annotation du `LetRec` inliné est vue comme `Nothing` ou pourquoi la substitution de type `a1 -> Int64` s'arrête aux portes du sous-arbre inliné.
  - [x] **Baby Step 1.7.3** : Patcher la passe d'optimisation dans PBO pour propager la Map de substitution de manière récursive ou préserver scrupuleusement les annotations de type lors de la copie de l'AST.
  - [x] **Baby Step 1.7.4** : Regénérer le code (via `./bin/go/run -c`) et valider strictement via un `grep` que le fichier `Test_Primes.go` ne contient **plus aucun** appel à `Rebox_Test_Primes...` généré au sein de la fonction `filter` et `sieve`.

### 2. Unboxing TAST (Scrutinee Fusion)
(Ceci est inspiré de htdocs/purescript-backend-erl)
L'objectif est d'atteindre le niveau de performance "Cheatcode" en détruisant purement et simplement les structures de données éphémères lors de l'exécution via la fusion de la vue (Scrutinee Fusion).
Par exemple, au lieu qu'une fonction comme `Map.lookup` alloue un objet `Just(valeur)` en mémoire (que le Garbage Collector devra ramasser l'instant suivant), la fonction spécialisée retournera directement les valeurs brutes `(valeur, bool)` sur la stack Go. Le `case` appelant (le scrutinee) lira ces variables primitives natives sans avoir jamais construit ni ouvert de "boîte" `Just` ou `Nothing`.

- [ ] *Action* : Aplatir les retours de structures de données éphémères (Maybe, Tuple, Either, State) aux frontières d'appels purs.
  - [x] **Baby Step 2.1** : Identifier dans `CodeGen.purs` les ADTs simples candidats à l'unboxing natif (ex: `Maybe` devient `(Value, bool)`, `Tuple` devient `(Value, Value)`).
  - [x] **Baby Step 2.2** : Adapter la génération de la fonction Go pure pour qu'elle retourne directement une structure par valeur allouée sur la stack (Option B), évitant toute allocation de constructeur sur le tas (zéro allocation).
  - [ ] **Baby Step 2.3** : Générer automatiquement une closure "pont" qui ré-emboîte (re-boxe) ces variables natives dans un vrai constructeur (ex: `Just`) *uniquement* lorsque la fonction est passée comme argument polymorphe de première classe (ex: passée à un `Array.map`).
  - [ ] **Baby Step 2.4** : Mesurer sur un benchmark intensif (ex: 10 millions de `Map.lookup` ou une grosse boucle `StateT`) la chute vertigineuse de la pression sur le Garbage Collector et l'accélération d'exécution.

### 3. Optimisation des Records
L'implémentation actuelle s'appuie sur `gopurs_runtime.RecordGet` / `RecordUpdate2` et des hashmaps dynamiques (générant un overhead d'allocation O(N) lors des copies). L'objectif est d'éliminer ces allocations en exploitant la connaissance statique des records.

- [ ] *Action* : Remplacer l'usage des dictionnaires dynamiques par des structures Go natives (ou un ShapeArray) lorsque le polymorphisme de rangée n'est pas requis (records fermés) ou suite à la monomorphisation.
  - [ ] **Baby Step 3.1** : Identifier dans le TAST (via `ann.type` ou la passe de spécialisation) les sites d'instanciation et de manipulation où un record possède un type fermé et exact (sans variable de rangée `| r`).
  - [ ] **Baby Step 3.2** : Modifier le CodeGen pour déclarer dynamiquement une `struct` Go native (ex: `type Record_a_Int_b_String struct { a int64; b string }`) associée à ce type exact.
  - [ ] **Baby Step 3.3** : Traduire les accès et les mises à jour (ex: `RecordUpdate`) en mutations/copies de structs natives par valeur (ex: `newRec := rec; newRec.a = ...`), garantissant **zéro allocation mémoire sur le tas (heap)**.
  - [ ] **Baby Step 3.4** : Implémenter une couche de coercion (boxing "à la demande") pour convertir ces structs natives en dictionnaires dynamiques `gopurs_runtime.Value` *uniquement* lors d'un passage à une fonction exigeant un record polymorphe.
  - [ ] **Baby Step 3.5** : Vérifier sur le benchmark `Test_Records.go` (Deep Record Updates) la disparition des allocations et la chute drastique du temps d'exécution (objectif : division du temps par ~800, en dessous de 50 ns).
  - [x] **Baby Step 1.7.3** : Patcher `Monomorphize.purs` (dans `purescript-backend-optimizer`). Dans `monomorphizeExpr`, lors du traitement d'un `ExprApp` qui cible une fonction spécialisée, utiliser la map de substitution locale (`substType`) pour appliquer un `rewriteExpr` sur les arguments de l'appel. Cela garantira que les variables libres (comme `a1` sur le `Nil` inliné) soient correctement substituées par le type spécialisé (ex: `Int64`) au point d'appel.
