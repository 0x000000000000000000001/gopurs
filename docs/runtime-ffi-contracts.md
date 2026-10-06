# Runtime et FFI du compilateur

## Propriétaires des frontières

| Frontière | Responsabilité |
| --- | --- |
| Types PureScript → représentations Go | `GoTypeOf`, `GoConversions` et `FfiBridge` choisissent le layout et les adaptations à partir du TAST. |
| Appels du compilateur → hôte JS ou Go | Les compagnons FFI de `src/` implémentent la même opération observable ; leurs caches et buffers ont un propriétaire explicite. |
| Valeurs Go boxées → stockage | `runtime/runtime.go` définit l'ABI de `Value`, les constructeurs, accès, applications et conversions génériques. |
| Source du runtime → compilateur construit | `tools/embed-runtime.mjs` produit `Runtime.js` et `Runtime.go` ; `Gopurs.Runtime` expose leur constante. |
| Compilateur → application produite | `Driver.Output.writeRuntime` écrit la constante embarquée dans `output/gopurs_runtime/runtime.go`. |

Le runtime est à la fois une API utilisée par les bibliothèques FFI et un
**artefact source copié à l'octet près** dans chaque application. La revue du
lot 14 garde ce fichier canonique unique : ses familles d'opérations sont
cartographiées ci-dessous et leurs contrats exécutés directement sur cette
source. Un découpage physique ou une réorganisation de ses déclarations
modifierait cet artefact. Les extractions de ce lot portent sur la frontière
interne du compilateur, où les responsabilités étaient mélangées.

## Stockage et durée de vie de `Value`

`Value` contient un tag `Type`, un champ `IntVal` et un `UnsafePtr` suivi par le
GC Go. Copier un `Value` copie cette enveloppe ; cela ne clone pas le stockage
pointé. Le producteur garantit le tag et le layout attendus par l'accès choisi.
Les accès directs de la FFI ne constituent pas un décodeur de valeurs arbitraires.

| Famille | Charge utile et propriété |
| --- | --- |
| Entiers, nombres, booléens | Charge dans `IntVal` ; les nombres utilisent `math.Float64bits`, les booléens 0/1. Les lectures typées passent notamment par `FloatVal`, `BoolVal` et `Unbox`. |
| Chaînes | Longueur **en octets** dans `IntVal`, pointeur sur les octets immuables dans `UnsafePtr`. `Str`/`StrValue` sont la frontière de lecture/écriture ; une sous-chaîne retient son support. La chaîne vide est canonique sans pointeur. |
| Tableaux | `Array` alloue l'enveloppe du slice et partage son support `[]Value`. La lecture d'un tableau de `Value` peut partager ce support ; une conversion vers une autre représentation convertit ses éléments. |
| Records | `Record` partage la map fournie ; `RecordDict` partage les slices clés/valeurs. Les formes fixes `RecordDict0`…`5` possèdent leur petit conteneur. Les conteneurs publiés comme valeurs PureScript restent immuables. |
| Valeurs opaques | `Any` retient une interface Go, donc aussi son objet, ses pointeurs et ses slices. Copier l'enveloppe conserve l'identité de cet objet ; la FFI propriétaire décide de sa mutabilité. |
| Fonctions | `Func`…`Func11` conservent la closure et ses captures sur le tas via `forceEscape`. Les applications partielles retiennent les arguments fournis. `WithFunctionData` attache des métadonnées à la fonction ; une application partielle restitue une fonction ordinaire. |
| ADT natifs | Le tag et le layout sont décidés par le générateur. `StructGetters` reçoit les getters pendant les `init` Go, avant leurs lectures concurrentes. Le champ `Rc` des ADT concernés sert au protocole de réutilisation ; la mémoire reste gérée par le GC. |

Les constructeurs de conteneurs n'introduisent pas de copie profonde. Le passage
FFI doit donc distinguer une **vue en lecture** d'un buffer privé à modifier.
Les conversions `Box`/`Unbox` ne sont pas non plus une promesse générale de copie :
`Box([]Value)` partage le support, tandis que `Box([]any)` construit les valeurs
boxées dans un nouveau slice. Les conversions typées du générateur possèdent
leurs propres règles, décrites dans [l'architecture](architecture.md).

Les [chaînes packées](packed-strings.md), les [closures](runtime-closures.md)
et la [réutilisation d'ADT](adt-reuse.md) détaillent ces protocoles. En
particulier, `runtime.KeepAlive` seul ne remplace pas l'échappement sur le tas
des captures imposé par `forceEscape`.

### Concaténation UTF-16

Les chaînes PureScript sont stockées en WTF-8 canonique : les paires de
surrogates valides forment un scalaire UTF-8, et les surrogates isolés gardent
leurs trois octets WTF-8. `ConcatString` recompose la paire pouvant apparaître
entre un surrogate haut en fin d'opérande gauche et un surrogate bas en début
d'opérande droit. Les opérandes restent immuables.

`PrimitiveExprs` émet ce helper pour `OpStringAppend` ; la FFI Go
`Data.Semigroup.concatString` l'utilise aussi. Le compilateur natif emploie ainsi
la même sémantique pendant le pliage des constantes. Une addition Go brute de
deux chaînes ne suffit pas à respecter ce contrat UTF-16.

### Opérations sur `Int`

Les primitives suivent les FFI JS de référence de Prelude et `Data.Int.Bits` :

- `IntNegate`, `IntAdd` et `IntSub` rendent un entier signé sur 32 bits avec
  débordement circulaire ; `negate bottom` vaut donc `bottom`.
- `IntMul` reproduit `(a * b) | 0` : le produit est d'abord arrondi comme un
  `Number`, puis converti sur 32 bits. Par exemple, `top * top` vaut `0` dans
  cette référence. Le modulo `2^32` précède la conversion Go pour accepter les
  produits dépassant la capacité d'un `int64`.
- `IntBitNot`, `IntBitAnd`, `IntBitOr` et `IntBitXor` utilisent des opérandes
  sur 32 bits. `IntShl`, `IntShr` et `IntZshr` masquent le compte par `31`, y
  compris pour les comptes négatifs. `IntZshr` rend le résultat non signé.
- `IntDiv` et `IntMod` respectent la division euclidienne de Prelude, avec zéro
  pour un diviseur nul. La division `bottom / -1` rend `2147483648` ; le test
  conserve explicitement ce comportement de la référence.

`PrimitiveExprs` et les FFI Go de `Data.Semiring`, `Data.Ring` et
`Data.Int.Bits` partagent ces helpers. Les wrappers binaires boxés délèguent aux
mêmes opérations. Le stockage `int64` de `Value` porte aussi les résultats
positifs de `zshr` et le quotient frontière : la conversion sur 32 bits se fait
aux opérations concernées. Le compilateur natif utilise ces contrats pendant
le pliage, et chaque opérande émis est évalué une seule fois, dans l'ordre.

### Affichage des `Number`

La FFI Go `Data.Show.ShowNumberImpl` suit la FFI JS Prelude `showNumberImpl` :
chiffres les plus courts permettant de retrouver le même binary64, notation
décimale pour `1e-6 <= abs(n) < 1e21`, notation scientifique en dehors de cet
intervalle. Les exposants utilisent `e`, un signe explicite et aucun zéro de
remplissage (`1e-7`, `1e+21`). Un entier en notation décimale reçoit `.0`.

Les deux zéros s'affichent `0.0` ; les valeurs spéciales sont `NaN`, `Infinity`
et `-Infinity`. Cela ne modifie pas la valeur stockée : le générateur préserve
les littéraux `-0` avec `runtime.NegativeZero`. Le compilateur natif utilise
aussi cette FFI pour écrire les littéraux numériques Go.

L'oracle est la FFI JS exécutée, sans arrondi arbitraire à 14 chiffres.
`NumberLiterals` vérifie les mêmes valeurs depuis leurs littéraux et après
lecture d'une `Effect.Ref`, dont les sous-normaux et les limites de notation.

### Vues JSON et mises à jour

- `ReadJSONObject` emprunte les maps étrangères et les `JSONObject` compacts.
  `ReadJSONObjectValue` matérialise une map pour les représentations record.
  Une vue ne donne pas l'autorisation de modifier le conteneur source.
- Un `JSONObject` compact possède son stockage et ses chaînes.
  `UnboxObject` et `RecordToMap` en matérialisent un conteneur indépendant,
  avec partage possible des valeurs imbriquées.
- Pour une valeur opaque contenant déjà une map, le contrat diffère :
  `UnboxObject` rend directement une `map[string]any`, et `RecordToMap` rend
  directement une `map[string]Value`. Les appelants qui écrivent doivent
  posséder la map ou en faire une copie explicite.
- `RecordSet` possède l'insertion et le remplacement immuables, y compris pour
  les maps étrangères. `RecordUpdateDict` vise le remplacement de champs déjà
  présents dans les records PureScript ; ses formes compactes partagent les
  clés et copient les valeurs. Son chemin de map hérite du contrat de
  `RecordToMap` : ce n'est pas une primitive de copie pour une map opaque.

### Travail asynchrone

`Retain`, `Release` et `EventLoopWait` comptent les tâches hôtes encore actives.
Le producteur retient une tâche avant de la lancer ; celle-ci libère exactement
une fois sa retenue à la fin. Le compteur doit déjà être positif quand l'attente
d'une génération commence. Ce protocole supervise la fin du programme ; la
durée de vie mémoire des valeurs reste assurée par leurs références Go.

## Contrats des compagnons JS/Go

| Module | Contrat commun et propriétaire de l'état |
| --- | --- |
| `Printer` | Échappement identique des littéraux : contrôles en `\xNN`, UTF-8 valide conservé, surrogates UTF-16 isolés représentés en WTF-8 puis échappés. Le scanner Go travaille en octets, sans itération qui remplacerait ces surrogates. |
| `Printer.Builder` | Un buffer opaque par rendu. `pushImpl` modifie le buffer et retourne la même poignée ; les anciennes poignées sont des alias. Le résultat de `toStringImpl` est une chaîne immuable qui survit à la poignée. |
| `GoAst` | `memoizeName` crée un cache par application partielle, y compris pour le résultat vide. Le callback est pur. En Go, le mutex encadre lecture et initialisation afin d'exécuter le callback une seule fois par clé. |
| `GoCode` | JS appelle le fallback PureScript fourni. Go adapte les `Value` à `scanReferencedImports`, puis boxe son résultat. Les deux reconnaissent les mêmes frontières de tokens, ignorent commentaires/littéraux et rendent des imports uniques triés. |
| `GoImports` | Les entrées sont empruntées en lecture. Le résultat possède un nouveau tableau, garde l'ordre et les doublons et partage seulement les chaînes immuables. `collectImports` possède le tri et la déduplication. |
| `FfiSupport` | Même parser Go et même protocole JSON. JS appelle le runner WASM ; Go appelle directement `gopurs_ffi_parser`, préparé par `prepare-native-output.mjs`. La façade PureScript décode la réponse et ajoute module/chemin aux erreurs. |
| `Runtime` | Constante identique à la source canonique, embarquée dans les deux compilateurs. Aucune lecture de `runtime/runtime.go` au lancement du compilateur construit. |
| `Metrics` | Temps monotone en millisecondes ; l'origine n'est pas commune aux deux hôtes. Le réglage du profil mémoire est effectif en Go et neutre en JS. Le pilote le configure avant les workers. |
| `Main` | Après nettoyage et diagnostic unique, JS fixe le statut en laissant Node vider ses sorties ; Go termine avec le statut 1 après l'écriture synchrone. |

`Printer.Builder.withOut` confine la mutation à un callback synchrone et
n'exporte que `Out`, `emit`, `emitMany` et `withOut`. Le callback chaîne les
poignées dans l'ordre ; il ne doit pas capturer la poignée pour un autre rendu.
Le type masque le constructeur sans prétendre assurer une propriété linéaire.
La FFI Go conserve un **pointeur** de `strings.Builder`, jamais une copie de ce
builder. Les rendus parallèles utilisent des buffers distincts.

Les autres petites FFI gardent une responsabilité unique et ne demandent pas de
dispatcher supplémentaire. Le parser, les adaptations typées des bibliothèques
et les diagnostics restent respectivement dans `ffi-gen`, `FfiBridge` et
`FfiSupport`/`Driver`. Les différences de transport n'ajoutent pas un deuxième
parser ni une deuxième politique de représentation.

## Vérifications reproductibles

Après `npm run build` :

```sh
node --test tools/native-ffi.test.mjs tools/embed-runtime.test.mjs \
  tools/runtime-contracts.test.mjs tools/string-concat.test.mjs tools/closure-lifetime.test.mjs \
  tools/integer-boundaries.test.mjs tools/integer-division.test.mjs tools/number-show.test.mjs \
  tools/apply-arity.test.mjs tools/function-data.test.mjs \
  tools/value-array-unboxing.test.mjs tools/go-imports.test.mjs
```

- `native-ffi` exécute les compagnons natifs réels sous `-race` avec le runtime
  canonique et le parser natif. Il compare échappement, scan et concaténation
  aux oracles JS, puis contrôle caches, buffers, GC et horloge. Les imports Go
  du runtime sont explicites, comme dans les autres FFI compilables isolément.
- `runtime-contracts` exécute tous les tests de `runtime/`, normalement puis
  sous `-race` : chaînes/sous-chaînes après GC, graphes de conteneurs après
  sortie du créateur, lectures concurrentes, alias de stockage, maps JSON,
  mises à jour immuables et retenues asynchrones.
- `string-concat` compare le Go émis par `PrimitiveExprs` et la FFI Prelude à
  l'oracle JS, sous `-race` : 801 couples UTF-16, bornes des surrogates, chaînes
  vides, caractères BMP/astraux, préfixes/suffixes et associativité.
- `integer-boundaries` importe les vraies FFI JS comme oracles : **5 978 cas**
  sur 13 opérations, exécutés sur le Go émis et sur les FFI Go sous `-race`,
  avec vérification de l'ordre et du nombre d'évaluations. Les bornes signées,
  résultats non signés, comptes de décalage et couples pseudo-aléatoires sont
  couverts. `integer-division` conserve ses 24 cas de lois euclidiennes.
- `number-show` compare **32 527 cas binary64** à la vraie FFI JS sous `-race` :
  zéros signés, NaN/infinis, chacun des exposants binaires finis, voisins des
  puissances de dix, sous-normaux et 4 096 motifs de bits pseudo-aléatoires.
- `embed-runtime` vérifie les octets embarqués, la résolution relative au script,
  les timestamps stables, la régénération après changement et le chargement JS
  sans le fichier source du runtime.
- Les suites closures/applications/métadonnées conservent leurs matrices
  d'arités, captures et applications partielles. `native-go-code.test.mjs`
  vérifie en plus le chemin compilé complet avec une pile bornée, à partir de
  `GOPURS_NATIVE_OUTPUT` issu d'un bootstrap conservé.

Les résultats datés, snapshots stricts et comparaisons des trois compilateurs
sont consignés dans [testing.md](testing.md).
