# Résolution des Bugs AOT (Gopurs)

## Prochaines étapes (Baby steps)

### 1. Optimisation de ListOps (Monomorphisation et TCO)
Le code de `ListOps` s'est amélioré grâce à la restauration du typage et de l'inlining, mais il reste des lacunes d'optimisation (notamment sur la monomorphisation et la TCO) :

- [x] **Le cas de la récursion (TCO cassée)** : La boucle TCO est détruite pour la version spécialisée de `foldl`, car l'appel récursif retombe sur la version générique polymorphe.
  - [x] *Action* : Inspecter pourquoi, lors de la monomorphisation d'une fonction récursive, `Monomorphize.purs` "oublie" de remplacer l'identifiant polymorphe par l'identifiant spécialisé dans le corps de la fonction (le `LetRec`).
    - [x] **Baby Step 1.2** : Corriger le renommage récursif lors de la spécialisation. *(Terminé : Le renommage via `rewriteExpr` est maintenant correctement appelé et substitue bien la variable `ExprVar` par la version monomorphisée)*.
      - 1.2.1 : Localiser dans `Monomorphize.purs` (probablement dans `monomorphizeExpr` lors de la création d'une nouvelle spécialisation) l'endroit où le corps de la fonction générique est cloné/réécrit.
      - 1.2.2 : Injecter une substitution (`Map Ident Ident` ou via `localDicts`) pour que l'ancien nom de la fonction (ex: `foldl`) soit remplacé par le nouveau nom manglé (ex: `foldl__12345`) dans le corps de la fonction clonée, afin que l'appel récursif pointe bien sur la version spécialisée.
    - [x] **Baby Step 1.3** : Vérifier que l'output Go généré contient bien la boucle native (le label de TCO). *(Terminé : La fonction injectée par `Monomorphize.purs` est maintenant encapsulée dans un constructeur `Rec` au lieu de `NonRec`, ce qui permet à l'analyseur `TCO.purs` de l'optimiser en boucle `for`)*.
- [ ] **Le cas du DPE (Dictionary Passing Elimination)** : L'inlining des primitives "Higher-Order" échoue (ex: `v_0` qui est en réalité `intAdd`). 
  - [ ] *Action* : Vérifier si le DPE est bien exécuté pour `mod` ou `intAdd`, et pourquoi il s'arrête à un appel de closure (`Apply2`) au lieu de se réduire à l'opérateur primitif `ExternPrimOp`.
    - [x] **Baby Step 2.0** : Valider empiriquement que la réparation de `wrapUnused` (retrait des dicts du wrapper) permet de compiler `Test.Ackermann` sans erreur `undefined`. *(Terminé : Les erreurs `__eta_norm` ont totalement disparu du build !)*
    - [x] **Baby Step 2.1** : Tracer l'évaluation de `add semiringInt` avant qu'elle ne soit passée à `foldl`. *(Terminé : Le compilateur PureScript extrait `add semiringInt` sous forme d'une déclaration globale `Test.ListOps.add`. PBO la voit donc comme une variable globale)*.
    - [x] **Baby Step 2.2** : Inspecter la fonction `isStatic` dans `Monomorphize.purs`. Vérifier si elle retourne `True` pour l'argument `intAdd`. *(Terminé : `isStatic` renvoie bien `True` pour `Test.ListOps.add` ET pour `0` (ExprLit), ce qui engendre un effet de bord inattendu sur la récursivité)*.
    - [x] **Baby Step 2.3** : Si `isStatic` bloque la spécialisation, ajuster le filtrage pour que les fonctions pures (comme les opérateurs mathématiques) soient systématiquement substituées dans le corps des fonctions de haut niveau. *(Terminé : J'ai modifié `isStatic` pour ignorer les valeurs de données (`ExprLit`, `ExprConstructor`), ce qui a réparé le bug de l'accumulateur écrasé. Le DPE fonctionne désormais parfaitement, `intAdd` est inliné, et le benchmark Go `Test.ListOps` s'exécute avec succès en **15 µs** !)*
### 2. Le bug de `void` (RecordGet sur Func)
Lors de l'investigation sur `void`, on a découvert que le code Go final génère `RecordGet(Func(\_ -> unit), "map")` au lieu d'un appel natif. Ceci est causé par un décalage d'arguments dans l'AST lors du "fallback" de `monomorphizeExpr` (quand `instType` contient encore des variables de type).

- [ ] *Action* : Corriger le filtrage des arguments (`filteredArgs`) dans `Monomorphize.purs` lors d'un fallback `rebuildSpine`, afin que le dictionnaire statique ne soit pas supprimé de l'AST recréé si la fonction cible n'est finalement pas substituée par sa version spécialisée.

### 3. Exploitation des TypeApp et Monomorphisation au Générateur 
  - [x] S'assurer que `Convert.purs` n'ignore aucun nœud `ExprTypeApp` issu du JSON (TAST). *(Déjà implémenté dans l'optimiseur)*
  - [x] Remplacer les fonctions rigides actuelles (`collectAppSpine` et `collectTypeAppSpine` dans `Monomorphize.purs`) par un `collectSpine` unifié et robuste (tolérant à l'ordre d'imbrication) pour dépiler conjointement les arguments classiques (`ExprApp`) et les arguments de type (`ExprTypeApp`). *(Déjà implémenté dans l'optimiseur)*
- **Action 2 (Générateur Go)** :
  - [ ] **Baby Step 3.1 : Capter et transmettre les TypeApp dans CodeGen.purs**
    - [ ] 3.1.1 : Créer un type de donnée `GoSpineArg = GoSpineApp (Array TcoExpr) | GoSpineTypeApp ExprType`.
    - [x] 3.1.2 : Remplacer `flattenApp` par `collectGoSpine` pour dépiler les `App`, `UncurriedApp` et `TypeApp`.
    - [x] 3.1.3 : Remplacer l'appel à `flattenApp` dans `translateExprImpl__` par `collectGoSpine`.
    - [x] 3.1.4 : Mettre à jour la logique d'application des `App`/`TypeApp` pour filtrer les types (à sauvegarder) et les valeurs (à évaluer), et supprimer l'ancien traitement isolé de `TypeApp`.
  - [ ] **Baby Step 3.2 : Amélioration de `exprTypeToGoType`**
    - [x] 3.2.1 : Modifier `exprTypeToGoType` pour préserver le type générique (ex: `T_a`) au lieu de le rabaisser systématiquement en `gopurs_runtime.Value` (TypeAny/TypeValue).
    - [x] 3.2.2 : Propager les déclarations génériques en Go (structs et func) avec `[T_a any]` pour satisfaire le compilateur Go.
  - [ ] **Baby Step 3.3 : Adaptation du Wrapper d'appels natifs**
    - [x] 3.3.1 : Utiliser `gopurs_runtime.AnyToValue` et `ValueToAny` pour déballer/remballer aux frontières d'appel du Worker (adaptation de `generateWrapperFunc` si nécessaire).

### 4. Unboxing total des ADT (Génération de structs natives)
L'objectif est d'éliminer les allocations de `gopurs_runtime.Value` pour les constructeurs de données (comme `Tree a` dans `Red-Black Tree`), afin de diviser par deux ou trois les temps d'exécution.
*(Prérequis : Avoir accompli le point 3 pour que `exprTypeToGoType` puisse lire les `TypeApp` et générer des pointeurs typés comme `*Node[int64]` au lieu de génériques polymorphes).*

- [x] **Baby Step 4.1 : Changer la signature des structs générées** (Déjà fait par l'étape 3.2.1)
  - Modifier `CodeGen.purs` (lors de `generateStructs`) pour utiliser les informations de `dataDecls` (du TAST) afin de générer les champs avec leur vrai type (`V0 T_a`, `V1 int64`) au lieu de `gopurs_runtime.Value`.
- [x] **Baby Step 4.2 : Gérer le polymorphisme aux frontières**
  - [x] 4.2.1 : Ajouter `gopurs_runtime.ReboxToStruct[T any](val Value) *T` via la réflexion Go (`reflect`) dans `Runtime.purs` pour faire une copie profonde d'un ADT polymorphe vers sa forme concrète.
  - [x] 4.2.2 : Modifier `coerceGoExpr` dans `CodeGen.purs` pour appeler `ReboxToStruct[T]` au lieu de `CoerceToStruct[T]` en cas de frontière polymorphe (quand le type d'origine et d'arrivée ne correspondent pas).
- [x] **Baby Step 4.3 : Optimisation AOT (Générer des fonctions natives de Re-boxing)**
  - [x] 4.3.1 : Dans `CodeGen.purs`, utiliser l'état global (ex: `Ref`) pour collecter toutes les paires `(SrcType, DestType)` de structs nécessitant un reboxing.
  - [x] 4.3.2 : Modifier `coerceGoExpr`, `boxGoExpr` et `unboxGoExpr` pour émettre un appel à la fonction native (ex: `Rebox_HashSrc_HashDest(expr)`).
  - [x] 4.3.3 : Lors de la génération finale du module, émettre le code Go des fonctions `Rebox_HashSrc_HashDest` avec des affectations statiques champ par champ.
- [x] **Baby Step 4.4 : Validation empirique du Re-boxing et résolution du Segfault (Panic)**
  - [x] 4.4.1 : Découverte de la cause racine du panic : PBO perdait le type des variables synthétiques de `ExprCase` (`v`, `v1`) car le TAST (JSON) omet l'annotation de type dessus.
  - [x] 4.4.2 : Patch de `Convert.purs` (dans PBO) pour forcer l'héritage du type depuis les `Binders` environnants. Les types sont désormais restaurés dans le backend Go.
  - [x] 4.4.3 : **Régression bloquante (Effet de bord du cache)** : La purge du cache `.purmeta` (pour valider 4.4.2) a fait réapparaître un bug masqué de `gopurs` : `Get_Control_Semigroupoid_composeImpl` est généré mais jamais défini. *(Terminé : En relançant la compilation complète de gopurs via `npm run build`, l'erreur a disparu, confirmant que le binaire `gopurs.js` utilisé par `bin/go/run` n'était pas à jour avec mes patchs précédents)*.
  - [x] 4.4.4 : **Investigation du panic (Nil pointer dereference)** : Lors de l'exécution de `./bin/go/run`, le code Go généré panique sur un "invalid memory address or nil pointer dereference" dans `Call_Test_ListOps_sumEvens.func3`. *(Terminé : La cause était `LetRec` qui inférait les types des paramètres via `paramTypes` sur le corps de la fonction. Comme le TAST n'a pas les types des variables de pattern matching dans les closures, les paramètres devenaient des `TypeValue` (boxes). J'ai modifié `LetRec` pour qu'il extraie les types des paramètres directement depuis le type AST de la fonction (via `extractExprFuncType`), ce qui préserve le typage strict des arguments natifs. Les listes sont unboxées avec succès, et `Test.ListOps` (900 elements) s'exécute en **10 ms** !)*
  - [x] 4.4.5 : **Investigation du panic (Nil pointer dereference) persistant dans Test.Primes** : Le panic `SIGSEGV` persistait dans `Test.Primes` au niveau de `Rebox_Test_Primes...`. *(Terminé : La cause était `GetCtorField` qui forçait la coercition du type du champ natif vers son type générique `[Value]` non-instancié. Cette fausse déclaration de type forçait `GoApp` (l'appelant) à réappliquer un `unboxGoExpr`, ce qui émettait un `CoerceToStruct` par-dessus un `Rebox`, créant une corruption de la disposition mémoire. En supprimant cette coercition redondante, le type instancié exact `[int64]` est préservé et transmis correctement. `Test.Primes` compile et s'exécute désormais avec succès en **7 µs** !)*

plus tard : ajouter des asserts aux tests pour vérifier que les calculs sont justes (e.g. prime sieve, additoon de valeurs pas de leurs adresses mémoire)

### 5. Monomorphisation des fonctions locales (LetRec)
L'optimiseur PBO actuel (`Monomorphize.purs`) spécialise très bien les fonctions globales polymorphes, mais ignore les fonctions définies localement (dans des blocs `let` ou `where` / `LetRec`). Lorsqu'une fonction globale (ex: `filter`) est inlinée dans un contexte monomorphe (ex: `sieve` travaillant sur des `Int`), sa boucle interne (`go`) reste polymorphe. Cela force le backend Go à insérer des appels massifs de deep-copy (`Rebox_...`) à chaque itération (complexité d'allocation `O(N^2)`).

- **Action** : Étendre le parcours du monomorphiseur (dans `purescript-backend-optimizer`) pour qu'il inspecte et spécialise les déclarations `LetRec` locales.
  - [ ] **Baby Step 5.1 : Traversée des LetRec dans le Monomorphiseur** : Ajouter la visite récursive des expressions de type `ExprLet` et `ExprLetRec` dans `Monomorphize.purs` de PBO pour collecter les signatures d'appel des fonctions locales.
  - [ ] **Baby Step 5.2 : Clonage et Substitution locale** : Si une fonction locale polymorphe (ex: `go`) est appelée avec un type spécifique (ex: `Int`), générer une version spécialisée (`go_Int`) dans les `Binders` du `LetRec` et substituer les appels internes.
  - [ ] **Baby Step 5.3 : Validation des types TypesApp locaux** : S'assurer que le passage des variables de type (via `TypeApp`) aux fonctions locales est correctement résolu et effacé une fois la fonction spécialisée.
  - [ ] **Baby Step 5.4 : Validation sur Test.Primes** : Compiler `Test.Primes` et vérifier dans le Go généré que `go` et `reverse` sont devenus monomorphes (`int64`), sans aucun appel `Rebox_...` dans la boucle. Comparer les performances avec le benchmark de référence (viser moins de 10 µs).