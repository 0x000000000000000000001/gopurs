# C2 — backend résident/incrémental

**Statut au 23 octobre 2026 : daemon one-shot existant, activé par
`GOPURS_DAEMON=1`, sans surveillance des sources.** Le protocole reste
séquentiel. Le daemon et le one-shot utilisent le même chemin de préparation
compacte prouvée lorsque le manifeste est actif et que le PBO utilise plus
d'un worker ; un refus de preuve entraîne la préparation complète.

Cette page conserve la conception initiale et les mesures antérieures à la
reprise du 23 octobre 2026. Les mentions de code restant à écrire décrivent
ces étapes historiques. Les anciens deltas ne prouvaient pas les changements
de demandes de spécialisation ; leurs chiffres, notamment les **20–21 s** du
daemon, ne sont pas des mesures du chemin corrigé. Le comportement actuel
est précisé dans la note finale et dans [C2-7](build-c2-persistent.md).

## Mesures historiques de départ

Feuille density 1024, 8 jobs :

| Phase | Temps |
| --- | ---: |
| Lookup + gel | 2,74 s |
| Préparation + monomorphisation | 9,75 s |
| Clés des modules | 5,34 s |
| PBO producer | 9,73 s |
| Go packages | 10,06 s |
| Publication/reste | ~6,05 s |
| **Total** | **43,67 s** |

Le cache disque accélère un build global, mais chaque édition repart d’un
processus neuf et rescane les 2 704 modules.

## Cible

- backend feuille **≤20–25 s** d’abord, puis ≤15 s avec B1/B3 ;
- aucune régression froide >5 % ;
- fallback complet au moindre doute ;
- pas de cache « approximatif » : les dépendances doivent être prouvées.

## Architecture proposée au stade initial

### 1. Processus résident

Un daemon garde :

- le snapshot TAST courant ;
- les modules préparés et monomorphisés ;
- les clés par module ;
- les parties PBO (body, implémentations, directives, workers) et les
  environnements d’émission ;
- le plan de packages et les bundles Go ;
- les fichiers Go précédents et les bridges C1c.

Il reçoit des deltas d’édition et répond par des artefacts ou des diagnostics.

### 2. Delta et fermeture

- hash du module modifié ;
- fermeture de dépendants prouvée par les clés/lectures existantes ;
- seuls les modules de la fermeture sont préparés/optimisés ;
- les autres sont réutilisés sans décodage complet ;
- si une clé ou une lecture manque, fallback full.

### 3. Publication

- staging incrémental par module/package ;
- pack de publication si le groupe change ;
- validation des objets lus seulement ;
- manifeste mis à jour sans rescanner 550 Mio.

### 4. Go

- plan de packages stable (topologique ou densité 1024) ;
- C1c : wrapper public stable, bridge, implémentation séparée ;
- seuls le bridge/implémentation changés et le lien sont recompilés ;
- `-p 14` pour la feuille.

## Prototype borné

1. Un daemon en mémoire qui ne garde que le TAST + les clés des modules.
2. Une édition de corps sans changement de signature.
3. Mesurer : delta TAST, clés, PBO, plan, Go, total.
4. Si le backend ne descend pas sous 30 s sur ce cas, arrêter C2.
5. Ensuite seulement, ajouter PBO/implémentations et publication.

## Risques

- Mémoire : 2 704 modules préparés/PBO peuvent être gros.
- Le cache en mémoire doit être invalidé correctement.
- L’ordre d’init et les ponts C1c doivent rester cohérents.
- Un daemon peut masquer des erreurs de persistance ; le mode one-shot doit
  rester l’oracle.
- La cible 30 s complète dépend aussi du frontend (~2 s) et de Go (~1 s).

## Décision au stade initial

C2 était alors en conception, non codé, avec un point de départ backend
mesuré à **43,7 s**.


## Point d’entrée constaté au stade initial

`src/Main.purs` appelle `Gopurs.Driver.compile` une seule fois. Dans
`src/Gopurs/Driver.purs`, `compile` :

1. lit la configuration ;
2. ouvre une `BuildCache.Session` avec `bracket (BuildCache.begin config) BuildCache.close` ;
3. prépare les modules ;
4. exécute le PBO et l’émission ;
5. publie puis ferme la session.

À ce stade, il n'existait aucune API de session réutilisable : le processus
natif était one-shot et son état mémoire était détruit à la fin. Rappeler
simplement `compile` dans une boucle aurait donc refait le backend.

## Premier refactor C2

Avant tout daemon, il faut extraire :

- un `SessionState` persistant (TAST, modules préparés, clés, plan) ;
- un `compileDelta` qui reçoit des changements et ne traite que la fermeture ;
- une séparation nette entre `close` de session et destruction du cache ;
- un mode one-shot conservé comme oracle.

C'était le prérequis identifié avant le prototype C2.


## Plan d’exécution C2

### C2-0 — baseline et frontière de session

- [x] Identifier le point d’entrée one-shot (`Main.compile`).
- [x] Identifier la session bracketed et les phases recalculées.
- [ ] Mesurer un cycle one-shot de référence et conserver les logs.

### C2-1 — extraction de l’API de build

- séparer `compileWithConfig` de `readConfig` ;
- exposer un `SessionState` minimal (config + input) ;
- conserver `compile` comme wrapper one-shot ;
- vérifier que les sorties et timings sont identiques.

**Gate :** aucune régression de comportement ; build des quatre hôtes non requis à ce stade.

### C2-2 — boucle résidente et cycle de vie

- [x] boucle locale `GOPURS_RESIDENT=1` ;
- [x] trigger `GOPURS_RESIDENT_TRIGGER` ;
- [x] config conservée dans le processus ;
- [x] erreurs de build non fatales, boucle relancée ;
- [ ] état backend réutilisé (C2-3).

**Gate :** pas de fuite d’état, pas de faux hit. Le build reste complet
pour l’instant ; C2-2 ne prétend pas encore gagner du temps.

### C2-3 — TAST/préparation persistants

- cache du TAST chargé et des modules préparés ;
- si l’input est identique, sauter `load TAST` et `prepareModules` ;
- si un module change, mesurer ce qui peut être réutilisé ;
- preuves de dépendances avant toute réutilisation partielle.

**Gate :** hit inchangé ≤1 s ; édition de corps mesurée ; fallback si doute.

### C2-4 — delta de modules

- préparer uniquement la fermeture modifiée ;
- réutiliser les modules préparés inchangés ;
- clés par module avec preuve de contexte ;
- mémo PBO/implémentations/directives.

**Gate :** backend feuille ≤25 s, sorties exactes.

### C2-5 — plan Go et publication

- plan de packages stable ;
- C1c : bridges publics, implémentations séparées ;
- publication et lookup incrémentaux ;
- cache Go des packages inchangés.

**Gate :** backend feuille ≤20 s ; Go ≤5 s ; froid ≤+5 %.

### C2-6 — protocole watch

- daemon local, arrêt propre, rechargement complet ;
- mêmes garanties que le one-shot ;
- tests de concurrence et de panne.

**Gate final :** chaîne complète ≤30 s sur une édition de corps, sans régression froide >5 %.

## Budgets

- C2-1/C2-2 : quelques jours.
- C2-3/C2-4 : le cœur, plusieurs jours.
- C2-5/C2-6 : intégration et validation.
- Aucun merge dans `edge` avant la gate C2-4.

## Risques principaux

- dépendances de spécialisation non prouvées ;
- mémoire des modules préparés/PBO ;
- ordre d’init et bridges C1c ;
- faux hit après édition de type ;
- cold qui régresse.

## Décision

C2 est **assumé comme un mini-projet**. Le premier code est C2-1, pas un
watch.


## C2-3 — mesure des modules préparés

La base et la feuille ont été rejouées avec le natif C2 et
`GOPURS_GO_PACKAGES=0` (les hashes préparés sont indépendants du layout Go ;
l’émission Go packages échoue encore sur `Promise_Internal_NewPromise`).

- 2 704 modules préparés base ;
- 2 704 modules préparés feuille ;
- **2 703 hashes identiques** ;
- **1 seul module modifié : `Inter.Cli.Ping.Main`**.

C’est la preuve qu’une **préparation par module** est viable pour une édition
de corps : les modules inchangés ne sont pas modifiés par la préparation
globale dans ce cas.

Reste à faire :

- extraire un `PreparedModule` par module ;
- préparer uniquement le module modifié + fermeture ;
- réutiliser les 2 703 modules préparés ;
- vérifier les métadonnées globales et les clés.

Cette mesure ne valide pas encore le cache : elle montre seulement que les
données sont réutilisables. Le prochain code est `prepareModulesDelta`.


## C2-3 — delta de préparation validé

Implémentation :

- `PreparedCache` (prepared modules, hashes source, modules source, `MonoCache`) ;
- `prepareModulesWithCache` ;
- `prepareModulesDelta` : un seul module modifié, réutilise les autres, monomorphise le module modifié avec `globalAstMap`/`instantiations` du cache ;
- fallback full si plusieurs modules changent ;
- cache persistant dans `Resident`.

Mesure base → feuille dans le même processus, `GOPURS_GO_PACKAGES=0` :

| Build | Backend | Préparation | Clés | PBO | Publication |
| --- | ---: | ---: | ---: | ---: | ---: |
| base | 106,652 s | 26,569 s | 5,767 s | 43,136 s | 26,342 s |
| feuille delta | **47,372 s** | **absente** | 7,216 s | 11,621 s | 1,321 s |

Validation : **11 136 fichiers identiques** à un build one-shot forcé
(`GOPURS_FORCE_BUILD=1`), `bad=0`.

Le delta de préparation est donc correct et saute bien la phase de 26,6 s
(cas base froide). Le backend feuille reste à 47,4 s : les clés (7,2 s) et le
PBO (11,6 s) dominent encore. C2-4 doit les rendre incrémentaux.


## C2-4 — clés incrémentales validées

Le delta de préparation transmet les modules préparés inchangés à
`ModuleCache.initialize`. Résultat sur le second build résident :

- `prepared fingerprints: computed=1, reused=2703` ;
- clés : **7,216 s → 0,696 s** ;
- backend feuille : **47,372 → 38,183 s** ;
- PBO : 11,204 s ;
- publication : 1,227 s ;
- **11 136 fichiers identiques** au build one-shot forcé (`bad=0`).

Le backend reste à **38,2 s** : le PBO (11,2 s) est maintenant le principal
poste fixe. C2-4 doit encore traiter le PBO/l’émission.


## C2-4 — cache PBO en mémoire

Le processus résident conserve les `BackendModule` du build précédent. Les
modules sans FFI dont les clés correspondent n'ont plus besoin de relire les
parties PBO ; les modules FFI sont réutilisés seulement si leur digest FFI est
identique.

Mesure base → feuille, `GOPURS_GO_PACKAGES=0` :

| Build | Backend | PBO | Clés | Publication | Load TAST |
| --- | ---: | ---: | ---: | ---: | ---: |
| base | 106,015 s | 43,136 s | 5,724 s | 24,535 s | 1,244 s |
| feuille delta | **34,749 s** | **4,992 s** | 0,715 s | 1,444 s | 1,607 s |

Le PBO passe de **6,321 à 4,992 s** avec la couverture FFI (mesure précédente
sans FFI : 6,321 s). Les sorties restent exactes : **11 136 fichiers** ; seule
`.gopurs-ownership.json` diffère entre deux publications.

Le backend feuille reste à **34,7 s** ; la cible C2 ≤25 s n'est pas atteinte.


## C2-4 — résultat final du delta backend

Le hachage structurel `Wire.write` de 2 704 modules source était le coût
résiduel. Il est remplacé par le hash des fichiers `corefn.json` déjà lus.

Mesure base → feuille dans le même processus, `GOPURS_GO_PACKAGES=0` :

| Build | Backend | Préparation | PBO | Clés | Publication | Load TAST |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| base | 89,874 s | 10,022 s | 45,154 s | 5,793 s | 24,352 s | 1,194 s |
| feuille delta | **14,734 s** | **1,807 s** | 4,899 s | 0,731 s | 1,292 s | 1,437 s |

Validation : **11 136 fichiers identiques** au build one-shot forcé
(`GOPURS_FORCE_BUILD=1`, backend 60,682 s), `bad=0`.

**La cible C2 backend ≤25 s est atteinte.** Le backend feuille est à
**14,7 s** dans ce scénario. La chaîne complète dépend encore de C1c/Go et du
frontend ; Go packages reste désactivé pour cette mesure.


## C2-5 — Go packages activés

Le blocage `Promise_Internal_NewPromise` venait d'un symlink `gopurs` placé au
mauvais niveau. Avec `app/gopurs` correct, le layout packages passe :

| Build | Backend | Go packages | PBO | Clés | Préparation |
| --- | ---: | ---: | ---: | ---: | ---: |
| base | 87,355 s | 10,957 s | 46,984 s | 7,073 s | 12,591 s |
| feuille delta | **29,150 s** | **11,536 s** | 5,499 s | 0,791 s | 1,925 s |

Go packages est désormais le principal poste du delta. Le backend packages
reste à **29,15 s**, au-dessus de la cible ≤25 s. Le monolithique reste à
**14,7 s** ; C2-5 doit rendre la planification Go incrémentale.


## C2-5 — overlay C1c branché

`GOPURS_GO_PACKAGE_OVERLAY=1` prépare un résultat de paquets privé dans
`preparePackageSession` lorsque `eligible-body` est prouvé, puis
`writePackageFiles` ne rappelle plus `packageFiles`. Les bundles inchangés sont
réutilisés par lien, le bundle overlay est ajouté, les entrées et les paquets
dépendants sont réécrits. Tout doute (type, méthode, `init`, FFI, import,
référence, plan) retombe sur le planificateur complet.

Mesures du binaire build 7 (Ping, trois éditions chaudes) :

| Édition | Backend overlay | Go packages | Backend full oracle | Go packages full |
| --- | ---: | ---: | ---: | ---: |
| `lot11` | 21,515 s | 0,021 s | — | — |
| `lot12` | **19,773 s** | 0,008 s | — | — |
| `lot13` | **18,960 s** | 0,004 s | 38,539 s | 10,227 s |
| `lot14` (build 8) | **17,323 s** | 0,003 s | 38,376 s | 10,146 s |

Les sorties Ping sont identiques à l'oracle full. Le `go build` de l'entrée
tombe à **0,8–2,1 s** après une compilation full ; la chaîne complète est
estimée à **22–24 s**. Le premier overlay après une republication full peut
rester à **21,5 s** à cause du lookup C2 plus froid.

Limites du lot : un module sans type, un seul module modifié, flag opt-in,
pas de type-check du paquet overlay avant publication (seul `go build` de
l'utilisateur le vérifie). Le fallback complet reste l'oracle.


## C2-6 — daemon one-shot existant, sans watch

`GOPURS_DAEMON=1` transforme `bin/gopurs` en client d'un daemon local. Le
processus et sa configuration restent en mémoire entre les invocations. La
préparation et les digests FFI disposent de références mémoire ; le plan Go et
les objets PBO ont aussi leurs caches disque. Leur réutilisation suit les
preuves et invalidations de la note finale. Le daemon ne surveille pas les
sources : chaque `b` écrit un jeton dans `GOPURS_DAEMON_DIR` et attend le statut
correspondant.

Fonctionnement :

- premier appel : démarrage du daemon, build froid ;
- appels suivants : même processus, build incrémental ;
- redémarrage automatique si le binaire, les arguments ou les options
  (`GOPURS_GO_PACKAGES`, overlay, jobs, force) changent ;
- `GOPURS_DAEMON_DIR` par défaut : `<projet>/.gopurs-daemon` ;
- `GOPURS_DAEMON=0` désactive le client ;
- en mode daemon, `GOPURS_GO_PACKAGES=1` et
  `GOPURS_GO_PACKAGE_OVERLAY=1` sont les défauts (surchargables à `0`).

Le protocole est séquentiel : un appel doit terminer avant le suivant. Plusieurs
`b` concurrents ne sont pas pris en charge. Le nettoyage automatique du daemon
sur `b -c` ne fait pas partie de ce prototype.

Mesures historiques sur Ping, avant la reprise du 23 octobre, même projet,
après édition frontend :

| Invocation | Backend | Remarque |
| --- | ---: | --- |
| premier `b` (daemon froid) | 91,9 s | layout packages, plan complet |
| 2e `b` | 24,9 s | première édition overlay |
| 3e `b` | 21,2 s | |
| 4e `b` | 20,7 s | |
| `b` sans édition | 1,15 s | hit global du cache |
| frontend `spago build` no-op | 2,05 s | mesuré séparément |

Ces mesures décrivent le prototype historique, câblé pour les hôtes natif et
JS, et mesuré uniquement en natif. Aucun nouveau chiffre daemon n'est déduit
des résultats one-shot de la reprise.

### Persistance disque du `PreparedCache` — prototype rejeté

Le plan Go et les parties PBO sont déjà persistés : un processus neuf avec
overlay applique bien le plan (`Go packages 4 ms`, `Go overlay`) et restaure
2 703 backends. Il paie encore préparation **11,18 s**, clés **5,56 s**,
restauration/émission PBO **16,67 s**, lookup/exports/publication ~8,6 s, soit
**40,03 s** de backend.

Un prototype monolithique (`Wire.write` de tous les modules préparés +
`MonoCache` + métadonnées) a ensuite été mesuré :

- snapshot **264 074 602 octets** ;
- encodage **47,8 s**, écriture **0,8 s**, sauvegarde totale **72,3 s** ;
- backend du build qui sauvegarde : **160,7 s**.

Le chargement n'a pas été retenu après ce résultat : le coût d'encodage et la
taille rendent le snapshot monolithique inutilisable. Un cache par module et
un manifeste incrémental seraient nécessaires, mais la restauration/émission
PBO resterait un plancher. Cette expérience avait conduit à retenir le daemon
à court terme ; elle ne décrit pas l'état préparé compact introduit ensuite.

## Note de reprise — 23 octobre 2026

Les validations et mesures de delta ci-dessus sont historiques et limitées
aux corps testés. La reprise C2-7 a reproduit un contre-exemple : un appel
polymorphe qui passe de `Int` à `String` change la préparation de son fournisseur,
sans modifier le source de ce fournisseur ni l'interface de l'appelant.

Le daemon conserve son protocole de jeton/statut et sa boucle séquentielle.
Après une édition, le driver essaie désormais `preparePersistent` lorsque
`config.buildManifest && config.optimizerJobs > 1`, sauf build forcé. C'est
le même chemin compact que le one-shot : l'état précédent doit correspondre
à la génération commitée, et les preuves de forme, de contribution et de
spécialisation doivent autoriser la réutilisation. Il ne s'agit pas du retour
à l'ancien delta fondé uniquement sur les hashes source/interface. Un refus
de preuve déclenche la préparation complète avec le TAST déjà chargé.

Lorsque cette condition de configuration est fausse, une édition repasse par
la préparation complète. Le chemin mémoire peut encore réutiliser une
préparation lorsque l'inventaire et tous les hashes TAST sont inchangés. Un
build forcé prépare toujours l'ensemble des modules. Les headers et objets
PBO ne sont restaurés qu'avec leurs preuves de contexte, d'empreinte préparée,
de spécialisation, de métadonnées et de FFI.

Les protections mémoire restent en place : les backends PBO non associés à
une clé de génération sont vidés à chaque session en miss. Un hit global
vérifié sur disque vide aussi les références préparation/PBO/FFI, car cette
génération peut avoir été publiée par un autre processus. Après une erreur,
les trois références sont vidées ; une erreur signalée par le chemin
incrémental déclenche un retry forcé complet, et les autres erreurs sont
rapportées au client sans arrêter la boucle du daemon.

Les **20–21 s** historiques restent des observations de l'ancien prototype.
Aucune nouvelle mesure de performance daemon n'est présentée ici. Voir
[la reprise C2-7](build-c2-persistent.md) pour le contrat et la validation de
la préparation compacte, ainsi que les mesures one-shot.
