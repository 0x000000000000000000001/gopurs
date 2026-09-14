# Frontière entre AST Go et impression

État du 14 septembre 2026, après le lot 4 de la deuxième vague de nettoyage.

Les émetteurs construisent les expressions et statements Go à partir du TAST. Ils décident des types natifs, du boxing, des noms, de l'ordre d'évaluation, des paramètres, des captures et des cibles TCO. `Printer` rend les nœuds obtenus sans lire ni modifier `CodegenState`.

## Inventaire des fragments bruts

L'inventaire porte sur les sites du générateur, pas sur le nombre d'occurrences dans les fichiers Go produits.

| Famille | Propriétaires actuels | Représentation et portée |
| --- | --- | --- |
| Déclarations de module | `ModuleDeclarations`, `ModuleBindings`, `FunctionExprs`, `GoConversions`, `CodeGen` | `GoDecl` représente les valeurs avec cache, structs ADT, fonctions nommées, initialisations et getters FFI. Les corps d'enregistrement des getters de classes et les affectations Rebox restent opaques, avec leurs dépendances. |
| Fonctions anonymes du parcours d'expressions | `BindingExprs`, `ModuleBindings`, `FunctionExprs`, `EffectExprs`, `AdtExprs`, `GoFunctions` | Les fonctions natives, enveloppes curryfiées et IIFE utilisent `GoFuncBlock` ou `GoFuncLit`, composés avec `GoCall`. |
| Autres expressions et statements | `BindingExprs`, `ControlExprs`, `CallExprs`, `ArrayIntrinsics`, `CallArguments`, `AdtExprs`, `EffectExprs`, `PrimitiveExprs`, `CodeGen` | Déclarations locales typées, labels et fragments de contrôle, pointeurs, littéraux et commentaires comportent encore des `GoRaw`. Ils constituent d'autres familles. |
| Conversions | `GoConversions`, avec certains rendus spécialisés dans `Printer` | Les conversions de records, tableaux génériques et ADT unboxés contiennent encore des fonctions Go assemblées en chaînes. Les conversions de tableaux d'entiers restent des nœuds structurés jusqu'au consommateur. |
| Adaptations FFI | `FfiBridge` | Les bridges sont rendus depuis les signatures Go et les métadonnées TAST, avec leurs adaptations propres. Ils ne passent pas par les nœuds de fonctions anonymes du parcours d'expressions. |
| Runtime | `runtime/runtime.go`, `Gopurs.Runtime` | La source Go canonique est incluse au build par `tools/embed-runtime.mjs`. La façade PureScript expose la constante embarquée, sans lecture du fichier Go à l'exécution ; extraction réalisée au lot 9. |

## Fonctions natives structurées

`GoFuncBlock params stmts retType` contient une liste ordonnée de paramètres nommés et typés, un corps de statements et un type résultat. Le printer émet uniquement `func(params) retType { body }`. Les `GoReturn` sont explicites : le corps peut donc contenir une boucle TCO sans recevoir de retour supplémentaire.

`GoFuncLit params stmts retExpr retType` est la forme avec retour final. Son rendu délègue à `GoFuncBlock` après ajout de `GoReturn retExpr`. Les usages existants, notamment les IIFE des intrinsics de tableaux, conservent ainsi leur rendu.

Une enveloppe runtime se construit par `GoCall (GoSelector (GoVar "gopurs_runtime") "FuncN") [ fonction ]`. Une IIFE se construit par `GoCall fonction []`. Les émetteurs choisissent `Func`, `Func2`, etc., les groupes de paramètres et les conversions avant de construire ces nœuds. Le nouveau rendu ne fusionne pas les fonctions imbriquées.

Le nœud historique `GoFunc` est supprimé. `GoFunctions.curriedFunction` prépare ses anciens usages avec les nœuds natifs : un paramètre factice pour une fonction sans argument, un groupe jusqu'à dix paramètres, puis des enveloppes unaires successives au-delà de dix jusqu'à ce que le groupe restant tienne dans un `FuncN`. Ce regroupement conserve le Go antérieur. Les nœuds spécialisés de records et de constructeurs conservent leurs conventions de rendu.

## Déclarations et imports

`GoDecl` distingue `GoCachedValue`, `GoStructDecl`, `GoFunctionDecl`, `GoInitDecl` et `GoForeignGetter`. Les structs portent leurs paramètres génériques et leurs champs typés ; les fonctions nommées portent leurs paramètres, leur résultat et leur corps. `CodeGen` fournit des groupes ordonnés de ces déclarations à `GoFile`, en conservant l'ordre et les séparations entre valeurs avec cache, déclarations natives et helpers Rebox, puis getters FFI.

`GoImports.collectImports` est l'unique propriétaire des imports du fichier. Il parcourt les déclarations, expressions et types structurés, puis trie et déduplique leurs dépendances. `Printer.printGoFile` imprime la liste reçue sans rendre les déclarations une première fois pour y rechercher des noms de packages.

`GoRaw` transporte un `GoCode` contenant `{ text, imports }`. Pour les fragments historiques, `rawGo` reconnaît les références à `gopurs_runtime`, `math`, `sync` et `unsafe` au moment de leur création. Cette détection textuelle locale reste un mécanisme de compatibilité ; un fragment peut aussi fournir explicitement sa liste d'imports. Le collecteur lit ces dépendances sans inspecter à nouveau le texte. Les corps opaques d'initialisation des getters et d'affectation Rebox restent distincts des en-têtes de déclaration désormais structurés.

## Contrat pour les prochains changements

- Conserver les expressions comme des `GoExpr` lorsque leur structure peut encore servir à un consommateur. Éviter `rawGo (printGoExpr ...)` pour une fonction déjà représentable par les nœuds ci-dessus.
- Construire les conversions avec les métadonnées TAST dans l'émetteur. Le printer ne doit pas déduire un type effacé, consulter le contexte de traduction ou enregistrer un helper Rebox.
- Préserver l'ordre des paramètres, des statements et des arguments, ainsi que les frontières des fonctions. L'impression ne doit ni dupliquer une expression, ni déplacer son évaluation dans une closure.
- Fournir des noms Go déjà choisis et des corps valides pour les nœuds de fonctions natives. `GoFuncBlock` ne renomme pas les paramètres, ne rajoute pas de variable temporaire et ne vérifie pas les chemins de retour.
- Garder les fragments bruts restants dans leur famille propriétaire. Un nouveau nœud doit remplacer une structure répétée utile, plutôt que simplement emballer une chaîne Go opaque.
- Déclarer les dépendances des nouveaux fragments opaques et compléter le parcours de `GoImports` lorsqu'un nœud structuré introduit une référence à un package.

## Conversions de tableaux

`GoBoxIntArray`, `GoUnboxIntArray` et `GoFreshFilterArray` conservent l'information nécessaire à `ArrayIntrinsics.normalizeFreshIntArrayRoundtrip`. La décision de normaliser le buffer privé du filtre se prend au site du fold, avant l'impression. `Printer` rend les conversions conservées et imprime le marqueur de fraîcheur de façon transparente ; il ne reconnaît pas ce motif dans du texte Go. Voir [la règle des tableaux d'entiers](array-roundtrip.md).

## Validation du lot 4 de la deuxième vague

La compilation et le bundle passent sans avertissement. Les 18 tests existants ciblant les conversions record/tuple, les fonctions sans argument, les workers importés, les tags natifs et l'initialisation récursive réussissent. Un contrôle différentiel temporaire compare également 18 rendus de currying avec l'ancien printer : arités 0, 1, 2, 9, 10, 11, 12, 20 et 21, avec une expression ou un bloc comme corps.

`bin/go/run -c` depuis `altbak.pub` termine avec statut 0. Les 387 fichiers Go sont identiques octet pour octet à la référence prise avant ce lot, sur les mêmes 300 entrées TAST, et les 14 résultats fonctionnels sont inchangés. Les sources PBO et altbak sont inchangées ; aucune conclusion de performance n'est tirée de ces runs. Preuves locales : `/private/tmp/gopurs-wave2-declarations-q81ed26w/verification.json`.
