# Optimisations du code généré pour les décodeurs JSON

Les changements du 22 septembre 2026 réduisent les représentations
intermédiaires produites par la composition de décodeurs. Ils interviennent dans
le générateur Go, après PBO, et s'appliquent à des formes générales de programme.
Aucun schéma JSON, module de benchmark ou traitement particulier du TAST n'est
encodé dans ces règles. Le parseur JSON et la validation réalisée par les
décodeurs restent les mêmes.

Les mesures finales et les ablations sont consignées séparément dans le
[bilan altbak des optimisations JSON générées](../../../altbak.pub-gopurs/docs/benchmark-results/2026-09-22-generated-json-optimizations.md).
Cette note décrit le périmètre de l'implémentation et
ses garanties ; elle ne reprend pas les chiffres des prototypes scratch.

## Fusion de la traversée indexée Array/Either

[ArrayTraverse](../src/Gopurs/ArrayTraverse.purs), appelé depuis
[CallExprs](../src/Gopurs/CallExprs.purs), reconnaît les applications qualifiées de
`Data.TraversableWithIndex.traverseWithIndexDefault` ou `traverseWithIndex` lorsque
les deux premiers arguments sont explicitement les dictionnaires standards
`Data.TraversableWithIndex.traversableWithIndexArray` et
`Data.Either.applicativeEither`.

L'implémentation générique compose `mapWithIndex` et `sequence`. La règle émet une
boucle avec un buffer résultat privé, rempli par index, qui remplace le tableau
intermédiaire des résultats et la reconstruction par sous-tableaux/concaténations
du parcours générique. Elle conserve le premier `Left` et retourne un `Right`
contenant le tableau lorsque tous les éléments réussissent.

Les gardes et contraintes sont les suivantes :

- Trois arguments créent une fermeture attendant le tableau. Quatre arguments
  exécutent la traversée. Les autres arités conservent le chemin ordinaire.
- Les symboles homonymes d'autres modules, les dictionnaires locaux ou opaques et
  les autres `Applicative` ne déclenchent pas la règle.
- Chaque argument est capturé une fois, dans l'ordre, avant les instructions du
  suivant. La création du callback reste au même stade dans la forme partielle.
- Le callback reçoit chaque index et chaque élément dans l'ordre, **y compris
  après le premier échec**. La boucle ne fait pas de sortie anticipée : le
  `mapWithIndex` strict d'origine évaluait tous les callbacks avant `sequence`.
- L'entrée n'est pas modifiée. Chaque invocation, y compris d'une fermeture
  partielle réutilisée, possède son propre buffer résultat.
- Les tests de tags, accès aux champs et constructions d'ADT utilisent
  [AdtExprs](../src/Gopurs/AdtExprs.purs) et les métadonnées de constructeurs. La
  règle ne contient ni tags numériques ni disposition d'erreur Argonaut.

Cette transformation ne constitue pas une réécriture générale des traversées
`Applicative`. En particulier, conserver seulement le premier résultat en erreur
n'autorise pas à supprimer les appels suivants.

## Réduction des applications immédiates

[ImmediateApplications](../src/Gopurs/ImmediateApplications.purs) réduit une
application immédiatement connue à une lambda unaire. La passe intervient après
les fusions de fonctions et de thunks, avant l'analyse d'ownership et le TCO. Elle
peut également déplacer l'application sous un `Let` ou dans les branches d'un
producteur de fonctions, lorsque chaque résultat possible permet effectivement
d'éliminer l'application.

La substitution est limitée à un argument qui est déjà une valeur : local,
lambda sans `LetRec`, ou littéral scalaire. L'argument compte au plus 128 nœuds ;
le paramètre remplacé apparaît au plus une fois dans le corps. Les branches
acceptées ont au plus quatre alternatives explicites, et leur branche par défaut
doit aussi être réductible. Une branche retournant une fonction opaque reste
inchangée.

Les appels, getters globaux, accès aux champs, constructeurs et tableaux ne sont
pas des arguments déplaçables. Les scopes `LetRec` restent intacts. Une lambda à
plusieurs paramètres partiellement appliquée garde son appel. Les variables liées
de la valeur transplantée reçoivent des niveaux frais, supérieurs aux niveaux de
l'expression initiale, pour éviter une capture entre scopes frères. Les variables
libres continuent de désigner leur liaison d'origine.

La reconnaissance traverse les enveloppes `Typed` et `TypeApp` de l'expression
fonctionnelle. Elle réduit une application connue ; elle ne supprime pas
globalement les annotations de type. Les effets et la divergence contenus dans
un callback restent différés jusqu'à son appel. Les tests protègent notamment les
branches qui ignorent ce callback, les effets précédant la branche sélectionnée
et la réutilisation de la fonction produite.

## Conservation des résultats natifs Maybe, Either et Tuple

Le traitement de `Typed` dans [CodeGen](../src/Gopurs/CodeGen.purs) conserve un
résultat `TypeStructValue` lorsque l'annotation attend `TypeValue` et que son identité d'ADT correspond à celle reconnue
par `getUnboxedADT` dans [GoConversions](../src/Gopurs/GoConversions.purs).
L'annotation décrit alors une valeur déjà représentée par la structure native
attendue ; elle ne provoque plus à elle seule un emballage intermédiaire en
`gopurs_runtime.Value`.

La règle utilise les représentations existantes :

| Type | Champs de la représentation native |
| --- | --- |
| `Maybe` | Valeur en `V0`, présence en `V1` |
| `Either` | Valeur gauche en `V0`, valeur droite en `V1`, choix de branche en `V2` |
| `Tuple` | Première valeur en `V0`, seconde en `V1` |

Les champs de données restent des `Value` et les indicateurs des `bool`. La table
`fieldIndex` partagée avec les conversions permet à `AdtExprs.getField` de lire
directement le bon champ natif, notamment `V1` pour le contenu de `Right`.

La garde porte sur l'identité du résultat réellement traduit et celle de l'ADT
annoté, et exige une conversion attendue vers `TypeValue`. Un layout natif de pointeur déjà plus précis conserve sa conversion et ses payloads typés. Une frontière dynamique nécessitant un `Value` conserve son boxing, ainsi
que les wrappers publics `Get_…`. Les tests de tags et les conventions des
constructeurs restent partagés avec les autres chemins du générateur. Ces
changements ne créent pas de nouveau `GoType` et ne généralisent pas cette
représentation à tous les ADT.

## Validation ciblée

Les tests exécutent du Go généré, en plus des assertions structurelles :

| Test | Contrat vérifié |
| --- | --- |
| [array-traverse-either.test.mjs](../tools/array-traverse-either.test.mjs) | Reconnaissance stricte ; 11 055 comparaisons contre le `TraverseArrayImpl` FFI courant, avec tailles de 0 à 65 et toutes les positions du premier échec ; ordre des arguments et callbacks, premier `Left`, entrée intacte, buffers indépendants, formes natives/boxées/partielles. Le fixture CodeGen comprend la forme à trois arguments rencontrée dans `decodeArray`. |
| [immediate-applications.test.mjs](../tools/immediate-applications.test.mjs) | Réductions permises, refus des arguments évaluables et des branches opaques, absence de duplication d'un paramètre utilisé plusieurs fois, scopes récursifs conservés ; oracle Go de capture, effets, échec et réutilisation. |
| [native-sum-results.test.mjs](../tools/native-sum-results.test.mjs) | Champs natifs de `Maybe`, `Either` et `Tuple`, absence de reboxing lors des lectures de résultats annotés, tags et payloads corrects aux frontières dynamiques, wrappers et évaluation unique des producteurs. |

Ces tests se lancent depuis la racine de gopurs, après compilation du backend JS :

```sh
node --test tools/array-traverse-either.test.mjs tools/immediate-applications.test.mjs tools/native-sum-results.test.mjs
```

Le test d'applications immédiates compare l'exécution optimisée à un oracle
explicite. Il ne présente pas deux appels à `CodeGen.translate`, qui inclut déjà
la passe, comme une comparaison avant/après. Les empreintes des corpus complets,
la présence des règles dans les vrais décodeurs régénérés et les ablations de
performance relèvent du bilan altbak lié plus haut.

## Limites actuelles

La fusion de tableaux utilise encore `Apply2` pour le callback. Son résultat
intermédiaire est un `Value`, le buffer de sortie est un `[]Value` et l'Either
final franchit une frontière boxée. Elle ne supprime donc pas les boîtes créées
par le callback lui-même. La spécialisation de sa signature et le maintien d'un
résultat natif jusque dans la boucle restent des pistes distinctes à valider.

De même, conserver la structure native de `Maybe`, `Either` ou `Tuple` élimine
une enveloppe sans rendre leurs payloads récursivement natifs. Une évolution vers
des payloads typés doit préserver les accords entre signatures de workers,
conversions, constructeurs, captures et interfaces FFI.

Le décodage continue de consommer l'arbre JSON existant. Ces règles ne constituent
ni un parseur spécialisé ni un décodeur direct depuis les octets. Elles ne
modifient pas le contrat `tcorefn` : informations de type, déclarations, rangées,
quantificateurs et validations du décodeur restent dans le périmètre de
correction. Un gain sur ces diagnostics n'établit pas, à lui seul, le gain d'une
compilation complète ou d'un chargement parallèle.
