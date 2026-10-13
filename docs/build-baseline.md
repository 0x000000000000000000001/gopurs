# Référence du cycle b8x — lot 00

Campagne du 10 octobre 2026. Le but est de disposer d'une référence reproductible
avant le raccordement du cache persistant, pour les builds identiques et les
modifications localisées. Les mesures couvrent le frontend, gopurs, la résolution
Go et le runner jusqu'à la frontière d'exécution de `Test.Main`.

**Comparaison après raccordement :** le
[lot 03](testing.md#plan-v3--lot-03--build-identique-vérifié-10-octobre-2026)
reprend ce corpus à l'octet près. Ses mesures appariées donnent **35,732 → 1,257 s**
pour gopurs et **38,145 → 3,958 s** pour la chaîne identique. Il documente aussi
le coût de première publication du cache. Les valeurs ci-dessous restent la
référence historique du lot 00.

Le [bilan des lots 04–06](build-incremental.md) reprend ensuite les mêmes TAST et
Go : gains après édition face au lot 03, surcoût froid mesuré, volume du cache
et preuves de correction. Le lot 06 mesure le parcours complet jusqu'à `exec` :
**−1,44 % feuille / +1,01 % partagé / −5,50 % FFI**, avec dispersion publiée,
et **+0,403 s** sur l'identique. Le package Go commun reste le coût dominant.

## État observé et outils

Au début de la campagne, les liens `spago.yaml`, `.spago` et `output` de b8x
sélectionnent **JS**. Le worker API tourne sous Node ; ses cinq derniers
healthchecks `Ping` renvoient `pong\n`, sans appel au verrou Go. La configuration
prévoit un healthcheck toutes les 30 s, avec un timeout de 10 s. Cette observation
porte sur une fenêtre d'exécution JS, pas sur la contention d'une production Go.

La sortie Go conservée contient **2 688 TAST typés**, version `0.15.16`. Ses
fichiers Go communs sont byte-identiques au corpus du lot 01 ; elle possède dix
points d'entrée supplémentaires et des fichiers `go.mod`/`go.sum` différents.
L'ancien constat d'une sortie Go issue de CoreFn non typé ne s'applique donc plus.

Le workspace de mesure copie les sources de travail et les dépendances locales.
Il reprend les overrides de `run/bak/go/spago.go.yaml`, avec omission explicite
de l'override `math` absent et inutilisé. Le frontend et le backend sont invoqués
séparément afin de chronométrer leurs frontières, puis réunis pour les relances
de bout en bout. La génération initiale produit **2 704 TAST / 755 005 types** ;
le delta avec la sortie Go conservée provient de ce nouveau snapshot applicatif.
Les sources et options restent figées pendant la campagne.

| Phase | Outils et réglages |
| --- | --- |
| Frontend | Spago 1.0.3, copie par octets du `purs` actif 0.15.16 typé ; codegen `corefn,docs` |
| Backend | Copie du launcher et du natif gopurs après le lot 02 ; réglages du launcher sur cet hôte : préparation/PBO/émission à 8, pipeline actif, `GOGC=off`, `GOMEMLIMIT=10GiB` |
| Hôte | macOS ARM64, 48 GiB, 14 CPU logiques ; Go 1.27.0 pour `go mod tidy` |
| Runner Go | Image API figée, Go 1.26.8 Linux ARM64 ; valeurs usuelles `GOMAXPROCS=2`, `-p 1`, `GOMEMLIMIT=12GiB` |
| VM Docker | Environ 23,5 GiB visibles, 14 CPU ; les conteneurs API observés n'ont pas de limite CPU/mémoire individuelle |

Le conteneur de mesure est dédié. Le répertoire Go y occupe le chemin habituel
`/var/www/b8x/run/bak/go/output`. Son verrou et ses exécutables publiés sont privés ;
les caches Go de packages et modules utilisent les volumes persistants habituels.
L'état chaud/froid est déterminé par les commandes réellement exécutées dans le
graphe d'actions, et non par la seule existence de ces répertoires.

## Scénarios

| Scénario | Transition mesurée |
| --- | --- |
| Complet | Sortie frontend vide, dépendances déjà disponibles ; première génération Go |
| Identique | Trois relances, mêmes sources et mêmes octets TAST/Go |
| Feuille | `Inter.Cli.Ping.Main` : `log "pong"` devient `log "pong-lot00"` |
| Partagée | `caseToCamel` de `Util.Type.String.String` ajoute `"-lot00"` au résultat, sans changer son type |
| FFI | `RemoveAccents` ajoute `"!"` au résultat, sans changer sa signature Go |

Chaque variante localisée part du même snapshot frontend de référence. Le
frontend, le backend et `go mod tidy` sont mesurés trois fois par variante ;
les inventaires et octets doivent être identiques entre ses répétitions.
Les changements applicatifs sont des sondes de mesure dans la copie privée.

Le runner construit le vrai point d'entrée `Test.Main`. Le hook externe
`BASH_ENV` s'arrête après publication du binaire et libération du verrou, juste
avant `exec`. `Test.Main` ne traite pas `--help` : la suite métier et PostgreSQL
ne font donc pas partie de ce chronométrage. Un oracle Go séparé exerce ensuite
les trois chemins modifiés, dont le vrai `Ping`, et vérifie leurs résultats.

## Lecture des horloges

- **Frontend / gopurs / tidy / runner** : durées murales mesurées aux frontières
  des processus. Le runner est chronométré dans le conteneur ; l'ouverture de
  `docker exec` depuis l'hôte n'est pas comprise dans cette durée.
- **Chargement et préparation** : compteurs existants de `Gopurs.Metrics`.
  `transitive specializations` est inclus dans `prepare + monomorphize`.
- **Optimisation et émission** : le pipeline fait se chevaucher ces tâches.
  `PBO producer`, `attemptMillis` des workers et `generation + writes` cumulé
  ne sont pas des composantes additives de `backend total`. Un passage
  séquentiel de diagnostic fournit une ventilation sans ce chevauchement.
- **Verrou** : acquisition mesurée par `_go-run`, résolution de 10 ms sous Linux ;
  le message d'attente indique une contention effective.
- **Compilation et lien Go** : `GOFLAGS=-x -debug-actiongraph=...` conserve les
  actions, identités et durées des outils. Une sonde avec le Go du conteneur
  vérifie que cette instrumentation conserve un binaire à jour avec **zéro
  invocation d'outil**. Les durées `CmdReal` sont cumulées par action ; les
  actions de cache sans commande ne sont pas comptées comme des compilations.
- **Mémoire** : échantillons `/proc` à une seconde, avec `VmHWM` des processus
  `compile`. Ce relevé est une observation de ressources, pas une modification
  des limites du runner.

Une modification connue des sources peut retrouver un objet Go déjà calculé
lors d'un passage précédent. Les premières constructions des variantes doivent
donc être distinguées de leurs répétitions à cache chaud ; elles ne sont pas
fusionnées dans une médiane unique.

## Résultats frontend et backend

Les médianes ci-dessous portent sur les durées des processus. Le build complet
possède un seul échantillon de première génération ; les autres scénarios en
possèdent trois, tous conservés avec leur minimum et maximum.

| Scénario | n | Frontend (s) | Processus gopurs (s) | `go mod tidy` (s) | Modules recompilés par purs | TAST changés | Go changés |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Complet | 1 | 28,429 | 38,693 | 2,069 | 2 704 | Nouvelle sortie | 3 010 |
| Identique | 3 | 1,228 | **36,902** | 0,946 | 0 | 0 | **0** |
| Feuille | 3 | 1,275 | 36,187 | 0,992 | 1 | 1 | 1 |
| Partagée | 3 | **20,685** | 35,612 | 1,619 | **1 375** | 1 | 5 |
| FFI | 3 | 1,226 | 37,257 | 0,817 | 0 | 0 | 1 |

La revue du premier harnais a identifié un biais : supprimer le backend de la
configuration Spago activait aussi `js,sourcemaps`. Les treize mesures frontend
ont été rejouées avec un backend neutre `/usr/bin/true`, qui conserve les options
réelles `corefn,docs`. **Chaque TAST est byte-identique** au corpus déjà utilisé
pour mesurer gopurs et Go. Les relevés préliminaires restent conservés, mais les
temps frontend du tableau viennent exclusivement de cette campagne corrigée.

À chaque invocation gopurs, les **3 010 fichiers Go sont réécrits**, même lors
des relances byte-identiques. La référence représente **133 809 190 octets de
Go**. Les trois relances identiques prennent 37,291 / 35,971 / 36,902 s ; leur
médiane ne masque donc pas un chemin de réutilisation déjà actif.

L'horloge interne `backend total` de la première génération mesure 36,251 s,
contre 38,693 s autour du processus. Cet écart initial de 2,442 s est hors de
l'horloge interne ; aucune cause n'est attribuée sans profil supplémentaire.
Sur les relances identiques, les médianes sont respectivement 36,804 et 36,902 s.

`Driver.Output.writeRuntime` réécrit également `go.mod` sous sa forme minimale
(`module gopurs/output`, `go 1.22`). Chaque `go mod tidy` recherche donc à nouveau
les dépendances ; il restitue les mêmes `go.mod`/`go.sum` normalisés, avec une
directive Go 1.26.0. Le raccordement du lot 03 devra appliquer la politique du
contrat de cache : créer `go.mod` seulement lorsqu'il manque.

### Ventilation du backend

| Horloge interne | Identique, médiane (s) | Diagnostic séquentiel (s), un passage |
| --- | ---: | ---: |
| Chargement TAST + tri | 1,298 | 1,170 |
| Préparation + monomorphisation | 10,710 | 16,665 |
| Dont spécialisations transitives | 7,276 | 11,518 |
| Optimisation + émission | 24,486 | 53,619 |
| Dont émission/écritures, lots cumulés | 21,177 | 19,267 |
| Backend total | 36,804 | 71,456 |

Le diagnostic fixe préparation/PBO/émission à un worker et désactive le pipeline.
Dans ce mode, la différence entre optimisation/émission et lots d'émission vaut
**34,352 s** : optimisation et orchestration hors émission. Cette soustraction
n'est pas applicable au pipeline parallèle. Les sorties du passage séquentiel
sont byte-identiques à la référence parallèle.

### Portée des changements

- **Feuille** : seul `purescript/Inter_Cli_Ping_Main.go` change.
- **FFI** : seul `purescript/Util_Type_String_String_ffi.go` change ; les TAST sont
  identiques, ce qui interdit de fonder le cache sur leurs seules empreintes.
- **Partagée** : `Util_Type_String_String.go`, `Core_Message_Field_Payload.go`,
  `Util_Style_Classname.go`, `Util_Type_String_Test_CaseToCamel.go` et
  **`Control_Semigroupoid.go`** changent. Ce dernier illustre les contributions
  de spécialisation provenant des appelants : la validité ne se réduit pas à
  « source du module inchangée ». Le frontend recompile 1 375 modules, mais un
  seul TAST change effectivement de contenu.

Ces mesures justifient le chemin sans changement du lot 03, puis la réutilisation
par module avec les environnements complets du lot 04. Elles donnent aussi deux
coûts distincts à suivre : la préparation globale qui subsiste après une édition,
et l'invalidation du frontend sur une fonction partagée. Elles n'établissent pas
encore un gain du cache : celui-ci devra être mesuré sur les mêmes transitions.

## Construction et lancement Go

Les dépendances Go et le runtime sont déjà disponibles dans le cache ; le
package applicatif de chaque nouvelle variante doit être compilé. Le graphe
confirme **deux compilations** (`purescript` et le point d'entrée), puis un lien.
L'ordre est complet → identique → feuille → partagée → FFI. Les variantes sont
définies relativement au corpus complet ; le dernier exécutable présenté à Go
provient du scénario précédent. Chaque variante est nouvelle pour le cache.

| Scénario | n | Runner jusqu'à `exec` (s) | `go build` (s) | Compilation cumulée (s) | Dont `purescript` (s) | Lien (s) | Pic RSS du compilateur (GiB) |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Complet | 1 | 219,238 | 219,080 | 215,103 | 215,015 | 1,155 | 14,497 |
| Identique, médiane | 3 | **0,770** | 0,420 | **0** | **0** | **0** | — |
| Feuille | 1 | 218,070 | 217,710 | 214,015 | 213,870 | 1,138 | 14,926 |
| Partagée | 1 | 216,502 | 216,130 | 212,304 | 212,214 | 1,206 | 14,503 |
| FFI | 1 | 224,711 | 224,280 | 220,453 | 220,360 | 1,200 | 14,840 |

Toutes les invocations réussissent avec **`2 / 1 / 12GiB`**, sans augmentation
ponctuelle de mémoire. `GOMEMLIMIT` est une cible souple du GC : le RSS observé
peut dépasser 12 GiB. Ces constructions initiales ont un seul échantillon par
scénario ; les écarts de quelques secondes ne constituent pas une comparaison
statistique entre variantes.

Les trois relances identiques durent 0,838 / 0,762 / 0,770 s, sans compilation ni
lien. Elles conservent le même exécutable, SHA-256
`9e08ce7474320728fac10c16d721e050b827a495f14eb3c10a70de970a546378`.
Les cinq oracles d'exécution réussissent, avec les résultats attendus de
`caseToCamel`, `RemoveAccents` et `Ping` pour chaque scénario. Les inventaires
Go, `go.mod` et `go.sum` restent identiques après les appels du runner.

L'acquisition du verrou est affichée à **0,000 s pour chaque invocation** ; aucun
message d'attente effective n'est observé. Le verrou de la campagne est privé.
Les observations avant/après retrouvent le worker Node et des healthchecks JS
réussis : cette campagne n'évalue pas la contention d'un déploiement Go.

La compilation du package monolithique domine donc largement le parcours après
une édition, même avec seulement un fichier modifié. Le cache gopurs pourra
économiser son propre travail ; le coût des **212–220 s** de compilation de
`purescript` demande une évolution distincte de la granularité des packages,
prévue parmi les suites conditionnées par les mesures du plan v3.

## Relance complète sans changement

La chaîne **`spago build` avec le vrai backend → `go mod tidy` → runner de
`Test.Main` jusqu'à `exec`** est ensuite chronométrée d'une traite sur le même
workspace, après un passage de préparation. Cette horloge comprend l'ouverture
de `docker exec`, contrairement au tableau du runner seul. Les wrappers de
sélection de commande et l'exécution de la suite métier ne sont pas inclus.

| Relance | Spago + gopurs (s) | Tidy (s) | Docker + runner (s) | Total (s) |
| --- | ---: | ---: | ---: | ---: |
| 1 | 39,760 | 1,017 | 1,355 | 42,134 |
| 2 | 37,054 | 0,876 | 0,961 | 38,892 |
| 3 | 38,785 | 1,005 | 1,016 | 40,808 |
| **Médiane** | **38,785** | **1,005** | **1,016** | **40,808** |

Les trois passages réussissent avec **zéro commande de compilation ou de lien
Go**, zéro attente mesurée du verrou, et des TAST/Go/fichiers de module identiques
à la référence. Le passage préparatoire de 41,481 s retrouve le package complet
en cache et refait uniquement le lien, car l'indice du dernier binaire désignait
alors la variante FFI ; il est conservé séparément.

Cette mesure constitue la référence de bout en bout du futur chemin sans
changement. Pour les builds modifiés, les phases ont été mesurées séparément :
une somme de leurs médianes ne doit pas être présentée comme un chronométrage
de bout en bout observé.

## Parité et contrôles de correction

Les quatre corpus distincts — référence, feuille, partagée et FFI — partagent
chacun les mêmes TAST entre **JS, natif séquentiel et natif parallèle**. Leurs
**3 010 fichiers Go par corpus**, runtime, bridges et points d'entrée compris,
sont byte-identiques dans les trois modes. Le `go.mod` minimal produit est aussi
identique ; les fichiers de module normalisés sont conservés avec les sorties
compilées sous Linux. Les trois répétitions de chaque variante native parallèle
produisent déjà le même inventaire et les mêmes octets.

Les passes JS de correction utilisent un worker et
`NODE_OPTIONS=--max-old-space-size=24576`, avec la pile du launcher à 64 MiB.
Leurs durées sont archivées comme diagnostics, hors des mesures natives du
tableau de performances. Les Go byte-identiques sont ceux compilés et exécutés
par les cinq oracles de la campagne Linux.

Les trois snapshots de `FFIIntegerReturns` et `NativeRecordWorkers` sont revus
contre les sorties formatées conservées du lot 02 : leurs octets correspondent.
**Six replays stricts réussissent ensuite**, deux fixtures dans chacun des modes
JS, natif séquentiel et natif parallèle, avec compilation puis exécution Go.
Les empreintes des compilateurs utilisés correspondent aux copies de la campagne.

L'audit final vérifie **4 318 fichiers inventoriés**, dont seule la configuration
Spago privée a changé comme prévu, et **8 112 observations de sources/compagnons
présents ou absents**. Les sources des TAST restent dans le workspace figé ; ses
mutations de mesure sont restaurées. L'exécutable de référence conservé fait
**75 779 980 octets**. Les contrôles de syntaxe JS/Python/Bash, de formatage de
l'oracle Go et `git diff --check` réussissent. Les changements applicatifs
concurrents sont consignés séparément dans l'audit de provenance.

## Reproducteurs et preuves

Répertoire de campagne :
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/b8x-incremental-lot00-srr89dtf/`.

- `initial-state.json`, `source-inventory-before.json`, `live-before.json` :
  révisions, changements concurrents, cibles, outils, healthchecks et processus ;
- `prepare.mjs`, `workspace.json`, `frozen-source-inventory.json`, `workspace/` :
  copie applicative, configuration Go adaptée et dépendances figées ;
- `toolchain/` : frontend, launcher, natif, bundle JS, ressources FFI et runner ;
- `measure-backend.mjs`, `mutations.json`, `measurements-backend/`,
  `backend-results.json`, `scenarios/` : transitions, chronomètres et sorties ;
- `measure-frontend.mjs`, `frontend-corefn-results.json`,
  `measurements-frontend-corefn/` : correction des options frontend et vérification
  de l'identité des TAST pour chaque scénario ;
- `setup-container.py`, `container.json`, `probe/` : image, montages et vérification
  du graphe d'actions ;
- `measure-go.mjs`, `handoff-only.sh`, `oracle.go`, `measurements-go/` : parcours
  du runner, graphes d'actions, mémoire et contrôles d'exécution ;
- `measure-cycle.py`, `cycle-results.json`, `measurements-cycle/` : chaîne
  complète mesurée après préparation, avec vérification des octets hors horloge ;
- `verify-b8x-modes.mjs`, `parity-results.json`, `parity/` : comparaison des quatre
  corpus dans les trois modes, avec sorties de correction conservées ;
- `snapshot-review.json`, `strict-replays.py`, `strict-replays.json` : revue des
  snapshots et six replays stricts ;
- `summary.json`, `summary.md`, `final-audit.json`, `README.md` : échantillons
  agrégés, audit final et commandes de reproduction sur une copie neuve ;
- `container-cleanup.json` : suppression des deux conteneurs dédiés après les
  mesures ; les sources, sorties et exécutables de la campagne sont conservés.

Les scripts refusent d'écraser un répertoire de campagne existant. Le conteneur
de mesure peut être recréé avec la commande enregistrée dans `container.json`.
Le rejeu doit utiliser des destinations neuves, les mêmes outils/entrées et un
état de cache explicite ; les caches Go partagés ne doivent pas être vidés pour
obtenir artificiellement un build froid.
