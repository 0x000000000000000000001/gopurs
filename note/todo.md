# TODO: Zero-Cost FFI (Box/Unbox Optimization)

Avec l'intégration du parseur WASM (`ffi-gen`) qui expose l'AST des fichiers FFI Go sous forme de JSON, le compilateur `gopurs` a maintenant connaissance de la véritable signature native des fonctions FFI.

## Objectif
Actuellement, les appels aux fonctions FFI depuis PureScript sont traités comme des appels génériques opaques (`gopurs_runtime.Value`), ce qui force des appels via `gopurs_runtime.Apply(bridge, arg)` et implique un cycle de boxing/unboxing et des allocations de closures inutiles.

L'objectif est d'implémenter un **Zero-Cost FFI** :
1. Dans `CodeGen.purs`, lors de la traduction des expressions (`translateExprImpl`), détecter les appels aux variables `foreign`.
2. Vérifier dans `ffiDecls` si la fonction FFI correspondante est "saturée" (tous les arguments attendus par la signature native sont fournis).
3. Si oui, court-circuiter complètement le wrapper généré (`var _Gopurs_XXX`) et compiler un appel natif Go direct (ex: `MonImpl(arg1, arg2)`).
4. S'assurer que les arguments natifs (ex: `int`, `string`) sont passés directement unboxés si le TAST indique qu'il s'agit de primitives.

Cela permettra d'éliminer totalement le surcoût de la barrière FFI entre PureScript et Go pour les fonctions correctement typées.
