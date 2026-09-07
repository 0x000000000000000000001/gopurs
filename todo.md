# Résolution des Bugs AOT (Gopurs) - V2 (Optimisations Avancées)

> **⚠️ ATTENTION : Le seul test à faire est `./bin/go/run -c` dans le dossier `altbak.pub`. N'utilisez jamais `./bin/run go` sans argument car cela exécute tous les backends et fausse les métriques.**

## Prochaines étapes (Baby steps)

### 1. Dead Code Elimination (DCE) de bas niveau
L'objectif est de nettoyer le code généré. Actuellement, la passe de défondamentalisation des closures génère des versions "de secours" (ex: `FALLBACK TCO`) et des variables intermédiaires (les `__tX`) qui finissent par ne jamais être utilisées dans le code final. Le compilateur Go les nettoie, mais cela gonfle la taille du code généré et rallonge la compilation.
- [ ] *Action* : Ajouter une passe d'analyse juste avant l'émission du code Go pour purger tout ce qui n'a aucune référence.
  - [ ] **Baby Step 1.1** : Implémenter un compteur de références sur les identifiants générés dans l'AST interne de PBO avant la phase de CodeGen Go.
  - [ ] **Baby Step 1.2** : Conditionner l'émission (print) des variables locales et des fonctions de repli (`FALLBACK TCO`) au fait que leur compteur d'utilisation soit strictement supérieur à 0.
  - [ ] **Baby Step 1.3** : Valider via `Test_Polymorphism.go` que les fonctions inlinées non appelées ne sont plus générées, tout en vérifiant que le code compile et s'exécute avec les mêmes performances.

### 2. Aplatissement garanti des retours de Records (Type Propagation)
L'objectif est de ne plus générer de `RecordDict` dynamique pour les structures de données (comme la Monade d'État), même dans des contextes complexes où la fonction est générée dynamiquement (ex: l'état final emballé dans `gopurs_runtime.RecordDict`).
- [ ] *Action* : Affiner la propagation des types depuis le TAST jusqu'au backend pour garantir qu'il renvoie toujours une `struct` native anonyme si le type est connu statiquement.
  - [ ] **Baby Step 2.1** : Identifier dans `CodeGen.purs` où le fallback vers `RecordDict` est appelé (ex: retour de la Monade d'État dans `Test_StateMonad.go`).
  - [ ] **Baby Step 2.2** : S'assurer que le TAST propage l'information de type fermée (`type: Record`) jusqu'aux blocs de retour des fonctions dynamiques, et déclencher la création d'une `struct` anonyme à la place du dictionnaire.
  - [ ] **Baby Step 2.3** : Valider sur `Test_StateMonad.go` que le record renvoyé par la chaîne d'état ne subit plus d'overhead de coercion, améliorant potentiellement les temps d'exécution actuels (108 μs).

### 3. Monomorphisation complète des Dictionnaires de Type Class
L'objectif est d'éradiquer les derniers `gopurs_runtime.Value` dans le code fortement polymorphe. Actuellement, un dictionnaire de type class monomorphisé stocke ses méthodes comme des interfaces dynamiques, ce qui force un boxing/unboxing (ex: `Rebox_...`).
- [ ] *Action* : Propager la spécialisation *à l'intérieur* de la signature du dictionnaire, pour qu'un dictionnaire `Monoidish Int` ne contienne plus que des méthodes typées `int64`.
  - [ ] **Baby Step 3.1** : Tracer la création des dictionnaires lors de l'instanciation de Type Class dans PBO et repérer pourquoi le typage interne conserve `gopurs_runtime.Value`.
  - [ ] **Baby Step 3.2** : Appliquer la même logique de substitution profonde (Deep Monomorphization) aux champs de la structure du dictionnaire.
  - [ ] **Baby Step 3.3** : Mettre à jour `CodeGen` pour générer une structure de dictionnaire Go totalement typée (ex: les champs `mappend_` utiliseront la signature primitive).
  - [ ] **Baby Step 3.4** : Vérifier sur `Test_Polymorphism.go` que la coercion (`Rebox_...`) disparaît totalement, optimisant les appels aux méthodes de classe.

### 4. Self-Hosting (Auto-hébergement) : Compiler `gopurs` avec `gopurs`
Le but ultime pour avoir un compilateur extrêmement rapide : compiler le code source PureScript de `gopurs` avec le backend `gopurs` pour obtenir un binaire natif Go (`gopurs.go`), plutôt que de tourner via Node/JS.
- [ ] *Action* : Amorcer la chaîne de cross-compilation et valider la parité des fonctionnalités.
  - [ ] **Baby Step 4.1** : Lancer la compilation du projet `gopurs` en utilisant le backend `gopurs` (la version JS actuelle).
  - [ ] **Baby Step 4.2** : Identifier et résoudre les éventuelles APIs FFI manquantes dans le portage de l'écosystème PureScript -> Go (ex: FileSystem, ChildProcess) nécessaires au compilateur.
  - [ ] **Baby Step 4.3** : Exécuter une compilation (ex: `altbak.pub`) avec le nouveau binaire `gopurs.go` et valider que l'output est identique (et mesurer le gain foudroyant de temps de compilation !).

### 5. Optimisation massive de Lazy Evaluation (Thunk Elimination & Typed Closures)
Le benchmark sur l'évaluation paresseuse (1 million de thunks forcés) montre que `gopurs` (~17.3 ms) est pénalisé par le boxing des closures (`gopurs_runtime.Func`) et l'application via `gopurs_runtime.Apply`. En monomorphisant les types flèches (Typed Closures), le temps tombe à ~11.0 ms. En poussant jusqu'à l'élimination totale du thunk via l'analyse de sévérité (Strictness Analysis), il tombe à 0.25 ms avec 0 allocation.
- [ ] *Action* : Étendre la monomorphisation profonde aux closures (arrow types) et implémenter l'analyse de sévérité pour éliminer les thunks déterministes.
  - [ ] **Baby Step 5.1** : Identifier dans le `purescript-backend-optimizer` (probablement dans `Monomorphize.purs` et `Convert.purs`) où la monomorphisation s'arrête actuellement face aux fonctions de première classe (comme dans le type `Lazy`).
  - [ ] **Baby Step 5.2** : Permettre au CodeGen d'émettre des closures purement typées (ex: `func() int64`) au lieu d'envelopper systématiquement les retours et les passages de fonctions dans le type boîte `gopurs_runtime.Value` ou `gopurs_runtime.Func`.
  - [ ] **Baby Step 5.3** : Évaluer la faisabilité d'une passe de *Strictness Analysis* pour détecter si un thunk est forcé de manière inconditionnelle (Thunk Elimination) et émettre le cas échéant une évaluation stricte immédiate.
  - [ ] **Baby Step 5.4** : Valider via le script de benchmark que `gopurs` rejoint les performances du cheatcode (vers 0.25 ms pour ce test extrême).
