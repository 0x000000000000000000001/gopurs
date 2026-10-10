# Référence du cycle b8x — lot 00

Campagne du 10 octobre 2026. Le but est de disposer d'une référence reproductible
avant le raccordement du cache persistant, pour les builds identiques et les
modifications localisées. Les mesures couvrent le frontend, gopurs, la résolution
Go et le runner jusqu'à la frontière d'exécution de `Test.Main`.

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
séparément afin de chronométrer leurs frontières. La génération initiale produit
**2 704 TAST / 755 005 types** ; le delta avec la sortie Go conservée provient de ce nouveau
snapshot applicatif. Les sources et options restent figées pendant la campagne.

| Phase | Outils et réglages |
| --- | --- |
| Frontend | Spago 1.0.3, copie par octets du `purs` actif 0.15.16 typé |
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

Les médianes ci-dessous portent sur les durées des processus. **Les colonnes
frontend sont celles du contrôle préliminaire avec génération JS/sourcemaps ;
elles seront remplacées par le relevé `corefn,docs` du parcours Go.** Le build complet
possède un seul échantillon de première génération ; les autres scénarios en
possèdent trois, tous conservés avec leur minimum et maximum.

| Scénario | n | Frontend (s) | Processus gopurs (s) | `go mod tidy` (s) | Modules recompilés par purs | TAST changés | Go changés |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Complet | 1 | 34,772 | 38,693 | 2,069 | 2 704 | Nouvelle sortie | 3 010 |
| Identique | 3 | 1,269 | **36,902** | 0,946 | 0 | 0 | **0** |
| Feuille | 3 | 1,299 | 36,187 | 0,992 | 1 | 1 | 1 |
| Partagée | 3 | **24,903** | 35,612 | 1,619 | **1 375** | 1 | 5 |
| FFI | 3 | 1,332 | 37,257 | 0,817 | 0 | 0 | 1 |

À chaque invocation gopurs, les **3 010 fichiers Go sont réécrits**, même lors
des relances byte-identiques. La référence représente **133 809 190 octets de
Go**. Les trois relances identiques prennent 37,291 / 35,971 / 36,902 s ; leur
médiane ne masque donc pas un chemin de réutilisation déjà actif.

L'horloge interne `backend total` de la première génération mesure 36,251 s,
contre 38,693 s autour du processus. Cet écart initial de 2,442 s est hors de
l'horloge interne ; aucune cause n'est attribuée sans profil supplémentaire.
Sur les relances identiques, les médianes sont respectivement 36,804 et 36,902 s.

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
- `setup-container.py`, `container.json`, `probe/` : image, montages et vérification
  du graphe d'actions ;
- `measure-go.mjs`, `handoff-only.sh`, `oracle.go`, `measurements-go/` : parcours
  du runner, graphes d'actions, mémoire et contrôles d'exécution.

Les scripts refusent d'écraser un répertoire de campagne existant. Le conteneur
de mesure peut être recréé avec la commande enregistrée dans `container.json`.
Le rejeu doit utiliser des destinations neuves, les mêmes outils/entrées et un
état de cache explicite ; les caches Go partagés ne doivent pas être vidés pour
obtenir artificiellement un build froid.
