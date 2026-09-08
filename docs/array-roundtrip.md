**Règle minimale pour les tableaux d'entiers — étapes 3.2 et 3.3**

Décision de conception du 8 septembre 2026, intégrée au générateur en 3.3 le même jour. Les validations ciblées de l'implémentation sont consignées en section 7, les mesures de 3.5 en section 8. Les régressions étendues restent prévues en 3.4.

L'objectif est de supprimer les deux buffers de conversion lorsqu'un **résultat frais de filtre**, converti en tableau natif d'entiers puis reboxé, est immédiatement consommé par le fold intrinsèque. Pour conserver exactement la normalisation des éléments, les deux boucles de copie deviennent **une passe de normalisation sur le buffer privé du filtre, avant le fold**. Le range, le filtre, le fold et le dispatch `Apply2` restent présents.

**1. Où le motif apparaît et où l'information disparaît**

Le témoin `sumEvens` de 3.1 suivait ce chemin avant optimisation. Son [snapshot actuel](/Users/0x1/Documents/htdocs/gopurs/gopurs/tests/passing-snapshots/ArrayRoundtrip.go:328) montre désormais la normalisation décrite en section 4. Les liens ci-dessous pointent vers les sites actuels du générateur :

| Étape | Représentation et site du générateur |
|---|---|
| Range | L'appel existant à `rangeImpl` fournit ici un `Value` contenant un tableau. Son bridge FFI reste exécuté. |
| Filtre | Le chemin `UncurriedApp Data.Array.filterImpl` génère un buffer neuf par `make([]Value, 0)` puis `append`. Son résultat Go est `TypeNativeArray TypeValue`. [CodeGen](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1910). |
| Contrainte `Array Int` | Le traitement de `Typed` impose `TypeNativeArray TypeInt64`, via `coerceGoExpr`. [CodeGen](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1240). |
| Première conversion | Emballage du header par `Array`, puis copie des `.IntVal` vers `[]int64`. Auparavant imprimée immédiatement en `GoRaw`, elle reste désormais structurée en `GoUnboxIntArray`. [CodeGen](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:361). |
| Deuxième conversion | `boxGoExprImpl` recrée un `[]Value` avec `Int(v)`, désormais représenté par `GoBoxIntArray`. [CodeGen](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:187). |
| Fold | Dans ce témoin, le fold emprunte le chemin **App**, qui boxe ses arguments puis lit un `[]Value` avec `Apply2`. Ce n'est pas le chemin `UncurriedApp foldlArray`. [Arguments](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1513), [boucle](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1542). |

Le boxing final peut être provoqué par un `Typed` extérieur ou par le boxing des arguments de `App`. Le point de reconnaissance doit donc examiner l'argument **après son boxing normal**, tout en conservant les conversions structurées jusque-là.

**2. Représentation conservée en 3.3**

Trois formes internes étroites ont été ajoutées à `GoExpr` :

- `GoUnboxIntArray sourceValue` : conversion d'un `Value` vers `[]int64`.
- `GoBoxIntArray sourceNative` : conversion de `[]int64` vers `Value`.
- `GoFreshFilterArray expression` : marqueur transparent attestant que le générateur vient de construire lui-même le buffer résultat du filtre intrinsèque.

Les deux premières formes sont produites aux branches de conversion `TypeInt64`, avant tout appel à `printGoExpr`. Leur impression sans optimisation reproduit les conversions précédentes. Les autres types continuent d'emprunter les chemins existants. Le marqueur ne change pas le code imprimé ; il est créé au site `make/append` du `filterImpl` reconnu, avec exactement deux arguments et uniquement pour son résultat `[]Value`, jamais à partir d'un nom de variable ou d'une ressemblance entre IIFE.

Cela évite d'ajouter une provenance à tous les résultats de traduction ou de la déduire par recherche textuelle dans du Go déjà imprimé. La preuve reste attachée à une expression locale et ne se propage pas dans la table des variables liées.

**3. Garde d'activation**

Dans le chemin `App` du fold, reconnaître uniquement l'argument tableau d'index 2, après le boxing normal, si toutes les conditions suivantes sont réunies :

1. Le symbole est explicitement qualifié `Data.Foldable.foldlArray`, avec exactement trois arguments. Établir le tableau entier par la conversion réellement produite vers `[]int64`, issue du `Typed (Array Int)` du candidat. Établir l'accumulateur/résultat par la signature concrète du callback `(Int, Int) → Int` et un seed littéral entier, notamment le zéro des témoins. En 3.3, les seeds dynamiques et les symboles locaux ou non qualifiés sont refusés, sans analyse supplémentaire.
2. L'argument est exactement `GoBoxIntArray (GoUnboxIntArray sourceValue)` : le type intermédiaire est réellement `[]int64`, sans conversion supplémentaire.
3. `sourceValue` est le simple emballage `Array(GoFreshFilterArray expression)`, produit par le générateur pour `Data.Array.filterImpl`. Le résultat frais est directement imbriqué dans les conversions et le fold.
4. Aucun stockage, retour, capture, passage FFI/ST ou autre consommateur ne sépare le filtre du fold. Le reconnaisseur ne traverse ni `GoVar`, ni `Let` restant après optimisation amont, ni appel ordinaire, branche ou `GoRaw` pour retrouver une origine supposée.
5. Les instructions associées à la traduction et leur ordre sont conservés. Le remplacement de l'expression tableau n'entraîne aucune nouvelle traduction ni duplication du producteur.

Ne pas utiliser seulement `getExprType` sur le fold ou son argument : dans le backend actuel de `sumEvens`, le fold est enveloppé par `Typed Int → Typed Int → Typed b`, et une annotation de fonction entoure l'argument tableau avant son `Typed (Array Int)`. Le seed possède lui aussi une annotation de fonction extérieure. Ces annotations intermédiaires ne décrivent pas nécessairement la représentation de la valeur à cet endroit. Le reconnaisseur doit s'appuyer sur les conversions structurées, la signature réelle du callback et la valeur du seed ; il ne doit ni rejeter le témoin sur la seule annotation extérieure ni supprimer globalement les `Typed`.

Le consommateur reçoit les éléments, jamais le slice complet. La fraîcheur du buffer découle du `make/append` émis par le filtre : le prédicat reçoit les valeurs d'entrée et n'a pas accès au buffer résultat. Cette preuve locale autorise sa normalisation sur place ; elle ne constitue pas une analyse générale d'ownership.

Une frontière FFI en amont du filtre, comme le bridge de `rangeImpl`, reste exécutée avec ses conversions actuelles. La règle ne la franchit pas : son segment commence au buffer privé produit par le filtre. Un tableau fourni directement par une FFI, une `Ref`, ST ou une fonction opaque ne porte pas le marqueur et conserve ses copies.

**4. Transformation retenue**

Forme reconnue, en pseudo-AST :

```text
foldInt(step, seed,
  GoBoxIntArray(GoUnboxIntArray(Array(GoFreshFilterArray(filterExpr)))))
```

Seule l'expression tableau est remplacée. Pseudo-Go de son évaluation :

```go
func() Value {
    source := Array(filterExpr()) // producteur et filtre complets, une seule fois
    items := *(*[]Value)(source.UnsafePtr)
    for i, item := range items {
        items[i] = Int(item.IntVal)
    }
    return source
}()
```

Les noms temporaires doivent être frais et locaux à l'IIFE, selon les mécanismes du générateur. Le code du fold reçoit cette expression à la même position qu'auparavant ; initialisation de l'accumulateur, appels `Apply2` et ordre des éléments restent identiques. Aucun backing array supplémentaire n'est créé pour normaliser les éléments.

Conserver précisément `.IntVal` puis `Int`, sans conversion Int32, contrôle de tag ajouté ou changement de calcul. L'aller-retour actuel transforme chaque élément en `Value` de type entier avec pointeur nul. Le simple remplacement par la `Value` d'origine ne préserverait pas cette propriété pour des valeurs non canoniques.

La normalisation se termine **avant le premier appel du fold**. La déplacer seulement dans l'argument de chaque `Apply2` pourrait prolonger la rétention des pointeurs présents dans le buffer source. La passe sur place les efface avant consommation, comme le font les copies actuelles. Elle n'exige donc pas de nouvelle analyse des signatures Go des bridges FFI, actuellement traitées après `translate` dans Main.

**5. Fallbacks et limites**

| Situation | Décision pour cette première règle |
|---|---|
| `sumEvens` et `sumRangeEvens` du snapshot 3.1 | Candidats attendus : paire entière, filtre frais et fold scalaire. |
| Filtre qui ne sélectionne aucun élément | Même règle ; header de tableau valide, normalisation vide, accumulateur inchangé. |
| `sumArrayEvens` du snapshot 3.1 | Pas de paire correspondante : conserver sa conversion d'entrée. |
| Paire sur un tableau conservé ailleurs, une variable ou une valeur opaque | Fallback, sans chercher à prouver l'unicité de cette valeur. |
| Tableau sans type Int établi, autres scalaires, records, ADT ou tableaux imbriqués | Fallback. |
| Retour, stockage, appel FFI/ST ou consommateur ordinaire après les conversions | Fallback, même si le producteur paraît frais. |
| Fold de l'autre chemin `UncurriedApp`, application partielle ou surapplication | Hors de cette première intégration. |
| Copie identité générale `Value → []Value`, fusion range/filter/fold, préallocation | Hors périmètre. |

Un tableau vide possède un header valide. Ne pas confondre `Array(nil)` et une `Value` dont `UnsafePtr` est nul ; ne pas ajouter de traitement spécial qui modifierait les erreurs existantes. Dans le candidat retenu, le filtre crée lui-même le tableau vide.

La capacité excédentaire du buffer filtré peut rester vivante pendant le fold, alors que le dernier buffer précédent était dimensionné à sa longueur. Les octets alloués diminuent en retirant les deux buffers ; cela ne garantit pas une baisse générale du pic RSS. La passe de normalisation conservée distingue aussi cette règle du prototype qui retirait les deux passes : **les mesures propres à cette implémentation figurent en section 8**.

**6. Critères de validation des étapes suivantes**

En 3.3, les deux `make` de conversion doivent disparaître des noyaux `sumEvens` et `sumRangeEvens`, avec une normalisation sur place avant le fold. Le range, le filtre avec sa politique d'append et le dispatch existant restent reconnaissables. Les formes exclues doivent continuer d'imprimer le code de conversion actuel.

En 3.4, conserver les 28 assertions Go/JS de 3.1 et ajouter les témoins nécessaires pour les refus, l'évaluation unique, l'ordre producteur → filtre → fold et la conservation des entrées accessibles ailleurs. Les tests doivent porter sur le comportement et le Go effectivement généré, dans `tests/passing` et ses snapshots. Vérifier notamment qu'un callback reçoit encore des entiers normalisés et que le marqueur ne survit pas à une frontière de stockage ou d'appel.

**7. État de l'intégration 3.3**

La règle est implémentée dans [CodeGen](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/CodeGen.purs:1052), avec les trois nœuds de GoAst et leur impression dans Printer. Le remplacement intervient uniquement sur l'argument tableau déjà traduit et boxé ; les instructions associées et le compteur `nextId` sont conservés. Les noms de la normalisation sont locaux à une IIFE, suffixés par la profondeur du fold et le compteur courant ; aucune expression du programme n'est évaluée sous leur portée.

`./bin/test ArrayRoundtrip -c` réussit avec le `purs` typé utilisé par altbak. Les 28 assertions passent aussi en JavaScript, avec sortie identique octet pour octet à celle du Go. Une seconde exécution sans mise à jour reproduit le snapshot. La comparaison avant/après ne change que `Call_Main_sumEvens` et `Call_Main_sumRangeEvens` : deux buffers et leurs boucles de copie sont remplacés par une normalisation sur place dans chaque fonction. Chaque noyau conserve son unique range, son filtre avec `append` et son `Apply2`. Toutes les autres fonctions du snapshot, dont `sumArrayEvens` et ses conversions d'entrée, sont identiques.

Le contrôle ciblé `./bin/test FFIIntegerReturns` réussit également sans mise à jour de ses deux snapshots (`Main.go` et FFI). Ses 27 assertions passent aussi en JavaScript. Ce contrôle protège notamment l'impression des conversions d'entiers conservées hors du motif.

La suite complète `passing` et les nouveaux témoins de refus, d'effets et de partage restent à traiter en 3.4. Aucune nouvelle mesure de performance ou de pic RSS n'a été réalisée en 3.3 ; les chiffres du prototype ne sont pas ceux de cette implémentation.

**8. Mesures de l'étape 3.5**

Le [bilan sur le Go régénéré](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-array-roundtrip-3-5-20260908/RESULTS.md) conserve le protocole, les dix paires de chaque série et la dispersion. Les économies d'allocation sont confirmées : −19,44 % d'octets/op à 900, −15,57 % à 90 000, deux allocations supprimées dans les deux cas. Les temps médians du noyau baissent de 2,47 % et 0,86 %, sans gain global d'altbak démontré. Le RSS de la sonde répétée reste stable à 900 et baisse à 90 000 ; ce n'est pas une garantie de RAM pour tout programme. Un petit ralentissement de Fib est reproduit par un contrôle isolé (+3 % environ), malgré son Go inchangé, avec cause non établie. La validation étendue de 3.4 reste ouverte.
