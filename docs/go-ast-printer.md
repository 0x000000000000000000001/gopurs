# Frontière entre AST Go et impression

État du 14 septembre 2026, après le lot 8 du nettoyage.

Les émetteurs construisent les expressions et statements Go à partir du TAST. Ils décident des types natifs, du boxing, des noms, de l'ordre d'évaluation, des paramètres, des captures et des cibles TCO. `Printer` rend les nœuds obtenus sans lire ni modifier `CodegenState`.

## Inventaire des fragments bruts

L'inventaire porte sur les sites du générateur, pas sur le nombre d'occurrences dans les fichiers Go produits.

| Famille | Propriétaires actuels | Représentation et portée |
| --- | --- | --- |
| Déclarations de module | `CodeGen`, `ModuleBindings`, `FunctionExprs`, `GoConversions` | Structs ADT, getters de classes, workers nommés et helpers Rebox sont encore assemblés en chaînes. `rawDecls` conserve leur ordre ; `GoDecl` représente les globals avec cache et `sync.Once`. |
| Fonctions anonymes du parcours d'expressions | `BindingExprs`, `ModuleBindings`, `FunctionExprs`, `EffectExprs`, `AdtExprs` | Les 16 sites qui assemblaient une fonction native, une enveloppe curryfiée ou une IIFE utilisent désormais `GoFuncBlock` ou `GoFuncLit`, composés avec `GoCall`. |
| Autres expressions et statements | `BindingExprs`, `ControlExprs`, `CallExprs`, `ArrayIntrinsics`, `CallArguments`, `AdtExprs`, `EffectExprs`, `PrimitiveExprs`, `CodeGen` | Déclarations locales typées, labels et fragments de contrôle, pointeurs, littéraux et commentaires comportent encore des `GoRaw`. Ils constituent d'autres familles. |
| Conversions | `GoConversions`, avec certains rendus spécialisés dans `Printer` | Les conversions de records, tableaux génériques et ADT unboxés contiennent encore des fonctions Go assemblées en chaînes. Les conversions de tableaux d'entiers restent des nœuds structurés jusqu'au consommateur. |
| Adaptations FFI | `FfiBridge` | Les bridges sont rendus depuis les signatures Go et les métadonnées TAST, avec leurs adaptations propres. Ils ne passent pas par les nœuds de fonctions anonymes du parcours d'expressions. |
| Runtime | `runtime/runtime.go`, `Gopurs.Runtime` | La source Go canonique est incluse au build par `tools/embed-runtime.mjs`. La façade PureScript expose la constante embarquée, sans lecture du fichier Go à l'exécution ; extraction réalisée au lot 9. |

## Fonctions natives structurées

`GoFuncBlock params stmts retType` contient une liste ordonnée de paramètres nommés et typés, un corps de statements et un type résultat. Le printer émet uniquement `func(params) retType { body }`. Les `GoReturn` sont explicites : le corps peut donc contenir une boucle TCO sans recevoir de retour supplémentaire.

`GoFuncLit params stmts retExpr retType` est la forme avec retour final. Son rendu délègue à `GoFuncBlock` après ajout de `GoReturn retExpr`. Les usages existants, notamment les IIFE des intrinsics de tableaux, conservent ainsi leur rendu.

Une enveloppe runtime se construit par `GoCall (GoSelector (GoVar "gopurs_runtime") "FuncN") [ fonction ]`. Une IIFE se construit par `GoCall fonction []`. Les émetteurs choisissent `Func`, `Func2`, etc., les groupes de paramètres et les conversions avant de construire ces nœuds. Le nouveau rendu ne fusionne pas les fonctions imbriquées.

Le nœud historique `GoFunc` conserve son rendu spécialisé, dont le regroupement de paramètres et les enveloppes runtime. Les nœuds spécialisés de records et de constructeurs conservent également leurs conventions de rendu. Leur simplification éventuelle demande un lot distinct ; elle ne découle pas de l'ajout de `GoFuncBlock`.

## Contrat pour les prochains changements

- Conserver les expressions comme des `GoExpr` lorsque leur structure peut encore servir à un consommateur. Éviter `GoRaw (printGoExpr ...)` pour une fonction déjà représentable par les nœuds ci-dessus.
- Construire les conversions avec les métadonnées TAST dans l'émetteur. Le printer ne doit pas déduire un type effacé, consulter le contexte de traduction ou enregistrer un helper Rebox.
- Préserver l'ordre des paramètres, des statements et des arguments, ainsi que les frontières des fonctions. L'impression ne doit ni dupliquer une expression, ni déplacer son évaluation dans une closure.
- Fournir des noms Go déjà choisis et des corps valides pour les nœuds de fonctions natives. `GoFuncBlock` ne renomme pas les paramètres, ne rajoute pas de variable temporaire et ne vérifie pas les chemins de retour.
- Garder les fragments bruts restants dans leur famille propriétaire. Un nouveau nœud doit remplacer une structure répétée utile, plutôt que simplement emballer une chaîne Go opaque.

`printGoFile` conserve actuellement la détection textuelle de certains imports (`unsafe`, `math`, runtime) après rendu des déclarations. Ce mécanisme reste inchangé dans le lot 8.

## Conversions de tableaux

`GoBoxIntArray`, `GoUnboxIntArray` et `GoFreshFilterArray` conservent l'information nécessaire à `ArrayIntrinsics.normalizeFreshIntArrayRoundtrip`. La décision de normaliser le buffer privé du filtre se prend au site du fold, avant l'impression. `Printer` rend les conversions conservées et imprime le marqueur de fraîcheur de façon transparente ; il ne reconnaît pas ce motif dans du texte Go. Voir [la règle des tableaux d'entiers](array-roundtrip.md).

## Validation du lot 8

Après compilation du nouveau nœud et migration d'un premier site dans `EffectExprs`, les occurrences équivalentes des cinq modules ont été migrées ensemble. `altbak.pub/bin/go/run -c` reconstruit le backend et le bundle, compile les 300 modules PureScript puis le Go, et termine les 14 benchmarks du mode `pure` avec statut 0.

Les 387 fichiers Go restent identiques octet pour octet à ceux du lot 7.6, avec les mêmes entrées CoreFn et résultats fonctionnels. Les snapshots de gopurs et les fichiers suivis d'altbak sont inchangés. Cette validation d'intégration remplace la campagne supplémentaire de fixtures à la demande de l'utilisateur ; aucune conclusion de performance n'est tirée de ces runs.
