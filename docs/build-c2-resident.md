# C2 — backend résident/incrémental

**Statut : conception ouverte, aucun code de production.** C1c règle le Go
pour les cas testés (scalaire + type), mais le backend reste à **43,7 s** sur
feuille. C2 doit supprimer les scans globaux.

## Mesures de départ

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

## Architecture proposée

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

## Décision actuelle

C2 est **en conception**, non codé. Le plancher backend reste **43,7 s**.


## Point d’entrée constaté

`src/Main.purs` appelle `Gopurs.Driver.compile` une seule fois. Dans
`src/Gopurs/Driver.purs`, `compile` :

1. lit la configuration ;
2. ouvre une `BuildCache.Session` avec `bracket (BuildCache.begin config) BuildCache.close` ;
3. prépare les modules ;
4. exécute le PBO et l’émission ;
5. publie puis ferme la session.

Il n’existe aujourd’hui **aucune API de session réutilisable** : le processus
natif est one-shot, et l’état de préparation/PBO/plan est détruit à la fin.
Un simple `--watch` qui rappelle `compile` referait donc tout le backend.

## Premier refactor C2

Avant tout daemon, il faut extraire :

- un `SessionState` persistant (TAST, modules préparés, clés, plan) ;
- un `compileDelta` qui reçoit des changements et ne traite que la fermeture ;
- une séparation nette entre `close` de session et destruction du cache ;
- un mode one-shot conservé comme oracle.

C’est le prérequis du prototype C2 ; aucun code de production n’est encore
engagé.


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
