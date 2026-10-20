# C1b — conception du découpage Go SCC-aware

**Statut : conception ouverte, aucun code de production.** Ce document fixe la
cible et l’algorithme après l’échec des isolations locales (C1).

## Constat mesuré

- Ping : SCC de 3 déclarations, isolation réussie, Go feuille **1,16 s**.
- FFI : SCC de 57 composants / 4 447 déclarations.
- Fermeture globale FFI : **2 510 déclarations / 205 unités / 10,9 Mo**
  (programme total : 185 667 déclarations / 123 Mo).
- Après extraction de FFI + `Util.Html.Clean.Clean` :
  - `p_iso` garde **12 imports directs sur cycle** ;
  - **51 packages** peuvent l’atteindre ;
  - une arête `p_02c4 → p_iso` porte **51 références** de symboles.
- Le package FFI d’origine fait 11 Mo et reste compilé après extraction.

Conclusion : une extraction locale ne suffit pas. C1b doit traiter le graphe
package complet et les cycles, pas une arête.

## Objectif

Rendre l’isolation d’un module valide **pour tous les modules**, en conservant :

- l’ordre d’initialisation ;
- les symboles privés et helpers partagés ;
- l’ABI des types et constructeurs générés ;
- les sorties et le comportement.

Cible Go feuille :

- module séparable : **≤5 s** ;
- module d’un SCC large : **≤15 s** ;
- arrêt si FFI reste **>25 s**.

## Algorithme prévu

### 1. Graphe package

- construire le graphe des déclarations (déjà disponible dans le planificateur) ;
- agréger en packages et en SCC packages ;
- mesurer la taille (déclarations, octets) et les dépendances de chaque SCC ;
- détecter les arêtes de rétroaction (feedback edge set).

### 2. Choix de coupe

Pour le module édité, choisir un ensemble déplacé `S` :

- fermé vers le bas si possible : toutes les dépendances internes de `S` sont
  dans `S` ;
- sinon, pour chaque arête sortante `S → A` qui revient vers `S`, choisir une
  inversion ;
- minimiser la taille de `S` et le nombre d’arêtes à inverser.

### 3. Inversion de dépendance

Trois mécanismes, du moins cher au plus cher :

1. **Déplacement de helper/type** vers un package bas neutre ;
2. **Callback/registre** : une variable de fonction dans un package neutre,
   initialisée par le fournisseur ;
3. **Interface/ABI de valeur** pour les constructeurs et types qui ne peuvent
   pas être déplacés.

Chaque inversion doit préserver :

- l’ordre d’init (les registres sont remplis avant l’exécution) ;
- les types Go (pas de conversion sauvage) ;
- les exports/imports et les entrées.

### 4. Rendu et validation

- rendre seulement les packages touchés ;
- compiler le `main` réel ;
- exécuter les tests et comparer les sorties ;
- replays JS/Go/Rust sur les mêmes TAST ;
- mesurer froid, feuille, RSS et taille des packages.

## Premier prototype borné

1. Sur FFI, construire le graphe package et lister les 12 imports sur cycle.
2. Choisir **une** arête à inverser par callback (`Util.Html.Clean.Clean` ou
   une fonction pure) sans toucher aux types.
3. Vérifier que le cycle diminue ; répéter jusqu’à un DAG.
4. Si le nombre d’inversions ou la taille de `S` explose, arrêter : C1b n’est
   pas borné pour ce module.

Le prototype doit rester privé et ne pas modifier gopurs.

## Risques

- Les types/constructeurs générés ne sont pas des valeurs dynamiques ; les
  callbacks ne suffisent pas toujours.
- L’init Go peut devenir un graphe de registres fragile.
- Le package d’origine reste gros s’il n’est pas lui-même découpé.
- Le froid et le lien peuvent régresser.
- C1b peut ne pas être borné pour les SCC de 4 447 déclarations.

## Décision actuelle

C1b est **arrêté** après la mesure du SCC : 49 packages / 92,2 Mo pour FFI,
contre 1 package / 2 Ko pour Ping. Aucun code de production, aucun commit.
Le plancher FFI actuel reste **~33 s**, le backend **43,7 s**.


## Mesure finale du SCC package

Analyse du graphe package après extraction FFI + helper :

| Cas | SCC contenant le module | Taille du SCC | Verdict |
| --- | ---: | ---: | --- |
| Ping | 1 package | 2 Ko | isolation propre |
| FFI | **49 packages** | **92,2 Mo** | SCC massif |

Le programme généré fait 72 packages / 125 Mo. Isoler FFI ne crée pas un petit
cycle : cela fusionne 49 packages dans un SCC de 92 Mo. Il n’existe pas de
petite coupe locale ; C1b complet exigerait de casser presque tout le graphe.

**Verdict : C1b n’est pas borné pour ce programme.** Le plancher Go FFI reste
~33 s. Une refonte complète du codegen serait nécessaire, sans garantie de
gain, et ne doit pas être présentée comme une optimisation.


## C1b — partition topologique SCC

Une seconde refonte a été testée : ordonner les SCC et les regrouper en paquets
de taille bornée, avec des noms stables par hash d’ancres.

Résultats :

| Variante | Packages | Plan stable | Go base | Go feuille |
| --- | ---: | --- | ---: | ---: |
| topo max 2 500 | 74 | oui | 49,4 s | **46,7 s** |
| topo max 5 000 | 38 | oui | 47,3 s | **44,7 s** |
| topo 2 500 + `-gcflags=all=-l` | 74 | oui | 26,8 s | **25,6 s** |

Le plan est stable et le package FFI modifié est petit, mais les dépendants
recompilent massivement. Même sans inlining Go, FFI reste à **25,6 s**. La
cible C1b de ≤15 s n’est pas atteinte.

**Verdict C1b : rejeté.** Ni l’isolation SCC, ni la partition topologique, ni
la désactivation de l’inline ne donnent un build FFI court. Le plancher FFI
reste structurel avec le modèle de packages actuel.


## C1c — bridge ABI/export (preuve de concept)

Un bridge de dépendance a été testé sur la fonction FFI
`Util_Type_String_String_RemoveAccents` :

- package public stable `p_t_2d34…` ;
- package bridge stable `p_bridge` avec `var RemoveAccents func(string) string` ;
- package d’implémentation `p_ffi_impl` qui enregistre la fonction dans `init` ;
- main importe l’implémentation pour l’initialisation ;
- wrapper public identique base/FFI.

Résultats :

| Mesure | Base | FFI |
| --- | ---: | ---: |
| Go build du `main` | 51,17 s | **1,01 s** |
| Valeur `RemoveAccents("é")` | `"e"` | **`"e!"`** |

Le package public ne dépend plus de l’implémentation modifiée : les dépendants
ne recompilent plus. C’est la première mécanique qui donne un Go feuille court
sur FFI.

### Limites

- Seules les fonctions à paramètres/retours simples ont été testées.
- Les types/constructeurs générés devront être extraits dans un package de
  types stable pour que les signatures contenant ces types fonctionnent.
- L’ordre d’init, les valeurs globales et les méthodes doivent être couverts.
- Il faut automatiser l’extraction des corps et la génération du bridge.
- Le backend reste à 43,7 s ; C1c ne règle que le Go.

### C1c multi-fonctions

Trois fonctions du même module ont été bridgées :
`RemoveAccents`, `UpperCaseFirst`, `LowerCaseFirst`.

| Mesure | Base | FFI |
| --- | ---: | ---: |
| Go build `main` | 50,17 s | **1,04 s** |
| Valeurs | `"e"`, `"Abc"`, `"aBC"` | `"e!"`, `"Abc"`, `"aBC"` |

Le mécanisme tient donc sur plusieurs fonctions. Le verrou restant est le
transfert des **types/constructeurs générés** et l’automatisation complète.

**Verdict : C1c est une piste sérieuse, pas encore une solution complète.**


### C1c et types générés

Test suivant : `Util_Timer__ClearInterval(id *IntervalId, …)` avec un type
généré. Le bridge `any` et l’implémentation séparée sont compilés, mais
l’implémentation ne peut pas accéder au champ privé `id.quit` :

```
purescript/p_ffi_impl/...: id.quit undefined (cannot refer to unexported field quit)
```

C1c scalaire fonctionne (**1,04 s**), mais les types générés avec champs privés
exigent une étape d’ABI supplémentaire :

- déplacer le type dans un package stable ;
- exposer les champs ou des accesseurs ;
- ou garder l’implémentation dans le package public pour ces types ;
- propager l’alias/ABI aux dépendants.

**Verdict initial : C1c n’est pas encore général aux types générés.** Le verrou
est l’ABI des champs privés, pas le bridge de fonction.

### C1c — ABI de type validée

Le type `IntervalId` a ensuite été déplacé dans `p_types` :

- champs exportés `Ticker` / `Quit` ;
- alias public `type IntervalId = p_types.IntervalId` ;
- implémentation séparée `clearIntervalImpl(*p_types.IntervalId, …)` ;
- bridge `any` + wrapper public stable.

Résultats :

| Mesure | Base | FFI |
| --- | ---: | ---: |
| Go build `main` | 52,96 s | **1,02 s** |
| `SetInterval` / `ClearInterval` | `*p_types.IntervalId` | `clear-ok` |

C1c couvre maintenant les scalaires **et** un type généré à champs privés.
Reste à automatiser l’extraction des types/fonctions et à traiter l’init/les
valeurs globales.
