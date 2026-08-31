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

### 2. Le bug de `void` (RecordGet sur Func)
Lors de l'investigation sur `void`, on a découvert que le code Go final génère `RecordGet(Func(\_ -> unit), "map")` au lieu d'un appel natif. Ceci est causé par un décalage d'arguments dans l'AST lors du "fallback" de `monomorphizeExpr` (quand `instType` contient encore des variables de type).

- [ ] *Action* : Corriger le filtrage des arguments (`filteredArgs`) dans `Monomorphize.purs` lors d'un fallback `rebuildSpine`, afin que le dictionnaire statique ne soit pas supprimé de l'AST recréé si la fonction cible n'est finalement pas substituée par sa version spécialisée.

### 3. Exploitation des TypeApp et Monomorphisation au Générateur 
  - [ ] S'assurer que `Convert.purs` n'ignore aucun nœud `ExprTypeApp` issu du JSON (TAST).
  - Remplacer les fonctions rigides actuelles (`collectAppSpine` et `collectTypeAppSpine` dans `Monomorphize.purs`) par un `collectSpine` unifié et robuste (tolérant à l'ordre d'imbrication) pour dépiler conjointement les arguments classiques (`ExprApp`) et les arguments de type (`ExprTypeApp`).
- **Action 2 (Générateur Go)** :
  - Amélioration de `exprTypeToGoType` pour préserver le type générique (ex: `T_a`) au lieu de le rabaisser en `Value`.
  - Adaptation du *Wrapper* avec `gopurs_runtime.AnyToValue` et `ValueToAny` pour déballer/remballer aux frontières d'appel du Worker.
