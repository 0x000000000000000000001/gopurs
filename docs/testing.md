# Tests et portée des validations

Les commandes de cette page partent de la racine de gopurs, avec Go, Node,
Spago et le `purs` TAST sur `PATH`. Voir [l'installation](../README.md).
Le runner utilise le bundle existant ; `-c` le reconstruit une fois via npm.

## Choisir un contrôle court

| Famille modifiée | Commande ciblée |
| --- | --- |
| Types et records | `./bin/test NativeRecordBoxing NativeRecordSizes -c` |
| Bridge FFI | `./bin/test FFIIntegerReturns -c` |
| Appels et fonctions | `./bin/test CurriedLambdas -c` |
| Conversions de tableaux | `./bin/test ArrayRoundtrip -c` |
| Récursion | `./bin/test TCO TCOMutRec -c` |
| Fusion de thunks | `./bin/test ThunkFusion -c` |
| Contrat du parser Go | `go test ./...` depuis `tools/ffi-gen` |
| Parser WASM et erreurs FFI | `npm run test:ffi`, après `npm run build` |
| Sélection, isolation et erreurs du runner | `npm run test:runner` |

Cette table indique quel contrôle choisir, pas que chaque fixture possède un
snapshot validé avec le dernier générateur. Les limites connues figurent plus
bas. Le parser ne doit être reconstruit avec `npm run build:ffi` que si ses
sources changent ; le WASM et son runtime JavaScript doivent alors rester
appariés, comme décrit dans le README.

## Sélection et snapshots

```bash
./bin/test --list
./bin/test TCOMutRec ThunkFusion --list
./bin/test --skip-before ThunkFusion --list
./bin/test FFIIntegerReturns --keep-workspace
```

Sans cible, `bin/test` sélectionne toutes les fixtures non exclues de
`tests/passing`. Les noms explicites gardent leur ordre ; la reprise inclut sa
cible. `--skip-before=NAME` et `skip_before=NAME` restent compatibles. Une cible
inconnue, une option inconnue ou une sélection entièrement exclue échoue avant
le build. `--list` ne compile rien et ne crée pas de workspace.

La vérification de snapshots est le mode par défaut. Un fichier absent ou
différent échoue : examiner le diff et la phase responsable avant de le
remplacer. Les snapshots sont `tests/passing-snapshots/<Fixture>.go`, avec
`<Fixture>_ffi.go` en plus si la source déclare `-- @snapshot-ffi`.
Une mise à jour intentionnelle s'effectue ainsi :

```bash
./bin/test FFIIntegerReturns --update-snapshots
# Équivalent historique : UPDATE_SNAPSHOTS=1 ./bin/test FFIIntegerReturns
```

Les snapshots ne sont écrits qu'après compilation et exécution Go réussies de
la fixture concernée. La campagne s'arrête au premier échec ; les mises à jour
des fixtures déjà réussies restent écrites. Le contrôle d'exécution conserve
le contrat historique : statut zéro et absence de `Fail` dans la sortie.

## Isolation, logs et caches

Chaque fixture reçoit son propre répertoire temporaire : sources, configuration
Spago, lockfile, `.spago` et `output`. `tests/runner` n'est ni lu ni modifié.
Les `.go`, `.js` et répertoires compagnons d'une fixture sont copiés avec sa
source. `-- @dependencies: assert prelude effect console` peut limiter ses
dépendances ; sans directive, la liste de `bin/pkg` est utilisée. Les checkouts
core restent requis et les paquets utilisent le cache global de Spago.

Les logs distinguent compilation PureScript, génération Go, formatage,
snapshots, dépendances Go, compilation Go et exécution. Les workspaces réussis
sont supprimés, sauf avec `--keep-workspace`. Échecs et interruptions conservent
le workspace et affichent son chemin. SIGINT/SIGTERM sont transmis à la
commande active et ses sous-processus. Une compilation PureScript échouée
n'est pas relancée automatiquement.

`-c` ne réinitialise pas les caches globaux : il reconstruit gopurs. Les sorties
de fixture sont neuves avec ou sans cette option. À l'intérieur du backend,
chaque lancement régénère le Go et relit la FFI ; changer seulement une FFI Go
ne demande pas de recompiler le PureScript si les entrées TAST sont inchangées.

## Exclusions et modules frères

Les neuf exclusions historiques sont conservées dans
[tools/test-selection.mjs](../tools/test-selection.mjs) :

| Fixtures | Motif enregistré dans le runner |
| --- | --- |
| `DerivingClause`, `DerivingContravariant`, `DerivingFunctorFromBi`, `DerivingFunctorFromPro`, `DerivingProfunctor` | Fonctionnalités de compilateur plus récentes que celles prises en charge par ces fixtures |
| `NumberLiterals` | Différences de sérialisation IEEE-754 |
| `StringEdgeCases`, `StringEscapes` | Surrogates isolés et chaînes Go UTF-8 |
| `2136` | Débordement aux bornes 32 bits, avec les entiers natifs 64 bits de gopurs |

Ces motifs décrivent les exclusions existantes, pas une nouvelle vérification
de chacune. `bin/modtest` sélectionne les checkouts frères `gopurs-*` possédant
un `bin/test` exécutable :

```bash
./bin/modtest --all --list
./bin/modtest --skip-before strings --list
./bin/modtest prelude strings
```

La sélection complète est le défaut. Les noms avec ou sans `gopurs-` sont
acceptés ; `-c` reconstruit le backend depuis ce checkout. Chaque script frère
gère encore ses propres sorties et nettoyages ; l'isolation des fixtures de
`bin/test` ne s'étend pas automatiquement à ces scripts.

## Bilan du nettoyage au 14 septembre 2026

Les lots récents utilisent, à la demande de l'utilisateur, ce jalon transversal
depuis **altbak.pub** :

```bash
bin/go/run -c
```

Les lots 7.4–7.7, 8, 9 et 10 ont chacun été validés par ce parcours : rebuild du
backend et du bundle, compilation de l'application, génération puis compilation
Go et exécution des 14 cas du mode `pure`. Les comparaisons ont conservé les
mêmes 300 entrées CoreFn, les 387 fichiers Go identiques octet par octet et les
14 résultats fonctionnels. Les durées ne servent pas à conclure sur les
performances ; leurs baselines restent celles du README d'altbak.

Le lot runtime a aussi vérifié la propagation d'une modification de la source
Go dans le bundle et l'exécution d'un paquet déplacé sans la source runtime.
Le lot runner a passé sept contrats avec commandes de compilation simulées,
puis deux fixtures minimales avec les vrais outils dans des workspaces
séparés. Ces checks ne constituent pas une campagne complète du compilateur.

Restent ouverts : la référence des snapshots TCO/TCOMutRec et le contrôle
ciblé de fusion du bilan de la première vague, la validation étendue de
la [règle ArrayRoundtrip](array-roundtrip.md) prévue en 3.4 de ce chantier, et
une campagne complète `passing` / modules frères. Les contrôles ciblés plus
anciens et leurs écarts préexistants sont datés dans le todo. L'ancienne
affirmation « 100 % des tests officiels verts » ne décrit pas cette validation.

Le lot documentaire a reconstruit le backend sans artefacts compilés, installé
son archive dans un projet npm vide et exécuté l'exemple du README via les deux
backends : mêmes 82 fichiers Go et même sortie. Les dépendances installées,
checkouts frères et caches ont été réutilisés ; le téléchargement de tous les
prérequis n'a pas été rejoué. Ce build avait révélé **85 avertissements
préexistants (73 sources, 12 dépendances)**. Le lot 12 les a supprimés : deux
compilations de référence sans sorties préexistantes ont recompilé chacune les
438 modules, passant de 85 à **zéro avertissement et zéro erreur**. Un build
incrémental silencieux ne suffit pas à établir ce résultat. Les preuves
détaillées jusqu’au lot 12 restent consultables avec `git show baa1e071:todo.md`.
Le [todo actuel](../todo.md) décrit la deuxième vague de nettoyage.

Le contrôle `ArrayRoundtrip -c --keep-workspace` du 14 septembre confirme les
28 assertions existantes et le snapshot inchangé. Le résultat incorrect du
singleton pair consigné le 9 septembre ne se reproduit plus : le résultat est
`8`. Cette vérification n'a nécessité aucune modification du compilateur ;
elle ne désigne pas la cause ni la correction de l'ancien échec.
