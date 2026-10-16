# Consommateurs multipackages — G04 validé

G04 est validé le **14 octobre 2026**, suite G à **5/7**, P à **0/7**.
Le candidat figé est `gopurs-go-packages-g03-20261013/candidate-5`, sans PGO.
Les preuves G04 se trouvent dans le répertoire de campagne
`gopurs-go-packages-g04-20261014` sous le temporaire OpenCode.

## Fixtures et snapshots

`tools/driver-package-consumers.test.mjs` exécute neuf consommateurs :
`NativeRecordWorkers`, `FFIIntegerReturns`, `StringEdgeCases`,
`NativeRecordBoxing`, `NativeRecordReturns`, `NativeArrayReboxing`,
`NativeTraverseCallback`, `JsonRecordPlan` et `B8xNativeGeneration`.
Pour chacun, le frontend produit un seul inventaire TAST. Le monolithe compilé
fournit l'oracle d'exécution ; JS, Go séquentiel, Go parallèle et Rust émettent
ensuite les mêmes octets multipackages, puis chaque sortie compile et s'exécute.

La première campagne termine les **36 passages**, mais refuse les neuf snapshots
absents. Leur revue retrouve **592 déclarations Main dans 57 fragments**,
équivalentes après retrait des qualifications et décodage des seuls noms de
champs record exportés. Les wrappers et workers FFI sont examinés également.
Les neuf références créées représentent **362 422 octets**.

`tools/test-snapshots.mjs` découvre les fragments `Main.go`, le `Main_ffi.go`
demandé par `@snapshot-ffi`, leurs fichiers d'ordre, les deux packages
d'initialisation et l'entrée principale. Le snapshot texte indique chaque chemin
avant son contenu Go formaté. `tools/test-runner.mjs` sélectionne ce format quand
`GOPURS_GO_PACKAGES=1` ; `GOPURS_TEST_COMPILER` permet de sélectionner un artefact
figé. L'écriture des snapshots reste postérieure à la compilation et à l'exécution.
Le test du runner couvre aussi un déplacement de package, le rejeu strict et le
refus de remplacer un snapshot après un échec de compilation.

Dans le [candidat G06 regroupé](build-go-measurements.md), le lecteur retrouve
Main et la FFI demandée à travers les repères source de `0_gopurs_package.go`.
La vue garde les imports du fichier réel et les sections sélectionnées, en
distinguant commentaires et chaînes Go. Les neuf références sont réexaminées
pour ce rendu, puis rejouées strictement sur les quatre hôtes.

Après création avec `bin/test --update-snapshots`, les empreintes correspondent
aux fichiers examinés. **207 commandes** rejouent strictement les neuf fixtures
sur les quatre hôtes ; aucune empreinte de snapshot ne change pendant le rejeu.

```sh
GOPURS_GO_PACKAGES=1 ./bin/test NativeRecordWorkers FFIIntegerReturns
GOPURS_TEST_MODES=js,native-sequential,native-parallel,rust \
  node --test tools/driver-package-consumers.test.mjs
```

## FFI et diagnostics

`tools/driver-packages.test.mjs` couvre désormais la FFI absente, une erreur de
syntaxe, le retour de la FFI, les dépendances Go locales/transitives et un
changement de signature native. Les **89 commandes** passent sur les quatre
hôtes. L'erreur de syntaxe conserve le manifeste courant et les sorties publiques,
avec un diagnostic indiquant le module et le chemin d'origine ; la réparation
retrouve le hit antérieur. Une FFI absente produit le stub prévu : la compilation
réussit, puis l'appel signale `FFI not implemented`. Son retour restaure les
octets et l'exécution précédents. Les fichiers externes et l'inode du verrou sont
vérifiés pendant ces transitions.

Les suites réelles de `gopurs-foreign` et `gopurs-js-promise` passent en Go
parallèle dans des copies isolées. Le harnais sélectionne explicitement le
candidat, après constat que le binaire installé n'avait pas son empreinte.

| FFI réfléchie de b8x | Contrat vérifié avant/après |
| --- | --- |
| `Foreign` | Nom Go `Date`, tags primitifs, conversion JSON des labels `type`/`value`. |
| `Promise.Rejection` | `Value.AnyVal`, erreurs natives/boxées, absence d'erreur. |
| `Util.Cookie.Cookie` | Clés de map `maxAge`, `path`, `httpOnly`, `secure`, `sameSite`, et wrapper réellement émis. |
| `Util.Json.TaggedSum` | Ajout de `value` manquant dans une map, sans dépendance au nom des champs Go. |
| `Util.Run.Router` | Labels de variant et callbacks, chemin natif et wrapper boxé. |

Les **onze probes** ont une sortie exacte entre monolithe et multipackage. Le
wrapper Cookie recevant une `Value` boxée ne développe déjà pas les options dans
la référence ; cette limite préexistante est conservée dans la preuve, tandis que
l'entrée native sur map vérifie les options. Les FFI examinées n'ont pas nécessité
de modification de production.

## Bootstrap et portée

Les bootstraps Go de G03, dont le candidat final en **150 packages**, compilent
et passent comme compilateurs les contrôles CLI. `prepare-native-output.mjs`
copie les sources du parseur et du cache avant la génération, afin que les exports
natifs soient disponibles pour la résolution typée. `build-native.mjs` et
`build-rust.mjs` utilisent ces mêmes compagnons ; Rust transporte aussi les plans
et génère les mêmes packages que les trois autres hôtes.

Le graphe typé et les compilations du bootstrap couvrent les références générées
des FFI PBO (`NativeMaps`, `CoreFn.Json`, `Json.Text`, `Usage`) et des helpers
Argonaut. Les codecs, arités, tags et noms PureScript gardent leurs contrats.
Les campagnes G04 valident les consommateurs ; G05 porte la restauration après
édition et G06 les seuils temporels/mémoire/stockage.

## Preuves conservées

- `before/`, `runner-packages-before.log` / `runner-packages-after.log` : ancienne
  hypothèse monolithique et validation du runner adapté ;
- `package-consumers-ob1lXL/`, `snapshots-review-1.json`, `snapshot-review.json` :
  première campagne, revue syntaxique et empreintes préalables ;
- `snapshots-update-1.log`, `gopurs-tests-MiEsMR/` : création après exécution ;
- `consumers-strict-1.log`, `package-consumers-pLueES/` : rejeu strict quatre hôtes ;
- `driver-diagnostics-1/2.log`, `driver-packages-YbTkOL/` : première assertion de
  diagnostic incorrecte, isolation des hôtes et validation finale ;
- `libraries-1/`, `reflection-1/` : suites réelles, sources et onze probes appariées ;
- `final-reference/summary.json` : artefacts, compteurs et empreintes finales.
