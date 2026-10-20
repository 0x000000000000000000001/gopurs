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
