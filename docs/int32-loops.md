# Boucles natives `int32` — 9 octobre 2026

## Résultat mesuré sur `rung`

Sur les mêmes **306 entrées TAST** du benchmark à 14 cas, la médiane du total
passe de **15,219493 ms à 12,500117 ms**, soit **−17,9 %**. Les deux versions
respectent le débordement signé sur 32 bits. La référence avant optimisation
utilise le compilateur `14bdfcf9e68d17d1c4816deab2a12f1aff7af26c` conservé avant
les modifications.

| Mesure | Avant | Après |
| --- | ---: | ---: |
| Polymorphism, médiane par ligne | 4,575208 ms | 2,232927 ms |
| LazyEvaluation, médiane par ligne | 0,458858 ms | 0,230134 ms |
| RBTree, médiane par ligne | 9,662438 ms | 9,527208 ms |
| Total, médiane des processus | **15,219493 ms** | **12,500117 ms** |
| Somme des médianes par ligne | 15,215416 ms | 12,500440 ms |
| Étendue des totaux | 15,177754–15,367258 ms | 12,441795–12,757503 ms |

Protocole : Go **1.27.0 darwin/arm64**, `GOGC=800`, `-pgo=off`, `GOFLAGS` et
`GOEXPERIMENT` vides. Cinq processus par variante, exécutés séquentiellement en
inversant l'ordre A/B à chaque paire, après la fin des builds et des validations.
Chaque processus utilise la calibration et le meilleur de dix batches du runner
existant. Les 14 résultats numériques sont vérifiés à chaque exécution.

Les binaires Go produits par les hôtes JS et natif parallèle ont aussi été
exécutés : **12,535517 ms** et **12,864618 ms** sur leurs contrôles individuels.
Le lancement habituel `./bin/go/run --run-only`, après reconstruction du workspace
actif, donne **12,720502 ms**. Ses **393 fichiers Go et son `go.mod`** sont
identiques à ceux de la campagne figée.

La baisse de Polymorphism et LazyEvaluation est reproductible et correspond au
changement de code machine. RBTree conserve exactement son Go antérieur ; son
petit écart de mesure ne constitue pas un gain attribué à cette optimisation.
Les diagnostics ArrayIndexing et WorkloadArrayInt sont chargés dans ce graphe,
mais leurs fonctions ne font pas partie du total historique à 14 cas.

## Transformation

La correction des entiers du 5 octobre a ajouté deux extensions de signe par
itération à certaines boucles. Un changement de parenthésage dans les helpers,
puis un stockage `int32` avec des temporaires `int64`, ont produit les mêmes
conversions dans le corps de boucle. Ces deux prototypes sont conservés avec
leurs désassemblages.

`Int32Loops` conserve également les paramètres d'itération et leurs opérations
en `int32`, après une garde sur les valeurs initiales. Forme simplifiée :

```go
// Avant : la normalisation est répétée à chaque tour.
for n != 0 {
    n = gopurs_runtime.IntSub(n, 1)
    acc = gopurs_runtime.IntAdd(acc, 1)
}
return acc
```

```go
// Après : chemin natif admis uniquement pour des entrées signées.
if n == int64(int32(n)) && acc == int64(int32(acc)) {
    return func() int64 {
        n32, acc32 := int32(n), int32(acc)
        for n32 != 0 {
            n32--
            acc32++
        }
        return int64(acc32)
    }()
}
// La boucle précédente traite les autres entrées.
```

Sur ARM64, le corps chaud de Polymorphism contient l'addition, la soustraction
et le branchement `CBNZW`. L'extension de signe de l'accumulateur se fait à la
sortie. Le type public reste `int64`, de même que les signatures des wrappers
et bridges. Les retours immédiats de valeurs non signées sont préservés.

La preuve porte sur l'AST Go structuré avant impression. Un slot est choisi
indépendamment si toutes ses écritures reconnues sont normalisées. Les
comparaisons ne se réduisent que lorsque leurs deux opérandes sont prouvés
signés. Les captures, masquages, boucles imbriquées, références opaques et formes
non reconnues font échouer l'admission concernée. Voir
[le contrat d'architecture](architecture.md#bindings-captures-et-tco).

## Validation et reproduction

Le test persistant construit le Go avant/après, le compile et l'exécute sous
`go test -race` contre **5 616 cas oracle JS par version**, soit **11 232
comparaisons** : débordements, valeurs initiales larges, sorties sans itération,
branchements, mises à jour simultanées, ordre/valeurs des opérandes à effets,
division, décalages non signés et constantes pliées. Une première version du
test reproduit l'erreur de conversion de `4294967295` en constante `int32` ;
le refus de cette fausse preuve et les conversions via helper la corrigent.

```sh
npm run build:native -- --keep-workspace
node --test --test-concurrency=1 tools/int32-loops.test.mjs \
  tools/integer-boundaries.test.mjs tools/binding-contracts.test.mjs \
  tools/local-native-returns.test.mjs tools/zero-arity-functions.test.mjs
UPDATE_SNAPSHOTS=0 ./bin/test Int32Loops TCOMutRec PartialTCO \
  ThunkFusion ThunkFusionNewtype CountedFunctions CurriedLambdas \
  2288 3957 InferRecFunWithConstrainedArgument
```

La campagne fige les TAST et les sources FFI avant de lancer l'ancien compilateur,
le nouveau JS, le natif séquentiel et le natif parallèle. Seuls les chemins
`modulePath` sont normalisés une fois vers les sources archivées ; ces mêmes
octets sont ensuite utilisés dans chaque mode. Les programmes des fixtures sont
comparés à leur exécution JavaScript, les fichiers Go et `go.mod` à l'octet près.

Au total, **18 fixtures et le benchmark** ont passé cette comparaison dans les
trois modes : **3 262 entrées TAST cumulées** sur ces 19 graphes indépendants,
et **4 183 fichiers Go/`go.mod` cumulés par mode**. Ces comptes comprennent les
modules communs à plusieurs graphes. Les 18 fixtures ont aussi passé leur
exécution contre l'oracle JS avant optimisation, puis dans chacun des trois
modes après optimisation.

Fixtures : `Int32Loops`, `TCO`, `TCOFloated`, `TCOMutRec`, `PartialTCO`,
`ThunkFusion`, `ThunkFusionNewtype`, `CountedFunctions`, `2136`,
`FFIIntegerReturns`, `LocalNativeReturns`, `CurriedLambdas`, `2288`, `2689`,
`3957`, `InferRecFunWithConstrainedArgument`, `OwnedTrees` et `RBTree`.

La revue utilise le parseur Go pour retirer uniquement les copies rapides
reconnues : elle retrouve exactement les fichiers avant optimisation après
`gofmt` : 43 copies rapides dans 17 fichiers générés sur l'ensemble des graphes.
Les dix snapshots concernés sont installés après cette revue. Les 18 fixtures
sont ensuite rejouées strictement sur les entrées figées ; les dix snapshots
mis à jour passent aussi le vrai `bin/test` avec recompilation PureScript.

Les scripts, sources/binaires avant-après, empreintes, TAST, logs, différences,
désassemblages et mesures individuelles sont conservés dans :

```text
/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/rung-int32-optimization-z3ic1dwe/
```

Fichiers principaux : `campaign.mjs`, `review.go`, `review/report.json`,
`parity-report.json`, `remaining-parity-report.json`, `strict-report.json`,
`remaining-strict-report.json`, `measure-final.py`,
`measurements/report.json`, `literal-range-red.log`, `targeted-tests.log` et
`live-results.json`. Pour une nouvelle mesure des artefacts figés depuis ce
dossier, avec un nom de sortie encore inexistant :

```sh
python3 -B measure-final.py --output measurements-replay
```
