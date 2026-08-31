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
- [ ] **Baby Step 4.2 : Gérer le polymorphisme aux frontières**
  - [ ] 4.2.1 : Ajouter `gopurs_runtime.ReboxToStruct[T any](val Value) *T` via la réflexion Go (`reflect`) dans `Runtime.purs` pour faire une copie profonde d'un ADT polymorphe vers sa forme concrète.
  - [ ] 4.2.2 : Modifier `coerceGoExpr` dans `CodeGen.purs` pour appeler `ReboxToStruct[T]` au lieu de `CoerceToStruct[T]` en cas de frontière polymorphe (quand le type d'origine et d'arrivée ne correspondent pas).
- [ ] **Baby Step 4.3 : Optimisation AOT (Générer des fonctions natives de Re-boxing)**
  - Plus tard, remplacer l'utilisation de `reflect` par la génération en dur de fonctions Go (`Rebox_Test_Primes_Cons`) pour éliminer l'overhead de la réflexion.
