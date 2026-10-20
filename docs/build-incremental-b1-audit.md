# B1 — audit du planificateur Go

**20 octobre 2026 — audit terminé, implémentation non commencée.** Mesure
privée sur le corpus b8x, avec instrumentation du même code que la production
dans une copie de bootstrap. Aucun changement de production.

## Découpage mesuré

Un passage de `packageFiles` sur les 2 704 modules générés :

| Phase | Temps |
| --- | ---: |
| Lecture/hachage des entrées | 0,17 s |
| Parsing Go | 1,28 s |
| Type-check `go/types` | 4,22 s |
| Graphe de déclarations | 2,34 s |
| Plan de packages | 0,42 s |
| Rendu/écriture et reste | ~2,9 s |
| **Total** | **11,57 s** |

Le type-check et le parsing représentent ~5,5 s, mais le graphe + plan + rendu
représentent ~5,2 s. Un cache qui ne supprimerait que le type-check ne suffirait
pas à garantir 5 s de gain.

## Faisabilité d’un cache par package

Le plan de packages contient six bundles ; une édition de corps ne change en
principe qu’un bundle. Un candidat B1 crédible devrait :

1. persister le plan, la correspondance module → package et les empreintes des
   fichiers rendus ;
2. réutiliser le plan et les bundles des packages dont les **signatures,
   imports et dépendances** sont inchangés ;
3. re-parser/re-typer uniquement le package modifié, avec les export data des
   packages inchangés ;
4. retomber sur le plan complet si les signatures, l’ordre ou la partition
   changent.

Le point difficile est le type-check du package modifié contre des exports
précédents. Le gain plafond est d’environ **8 s** si 5 des 6 bundles sont
réutilisés, mais la sélection, les exports et la reprise peuvent en consommer
une partie. Aucun prototype ne démontre encore ce gain.

## Décision

**B1 est plausible mais reste un chantier de représentation**, pas un cache de
sortie. Il faut un prototype borné avant toute intégration :

- commencer par les changements de corps sans changement de signature ;
- mesurer le type-check partiel et le rendu d’un seul bundle ;
- cible ≥5 s sur feuille, sinon arrêter.

Les chiffres sont conservés dans `htdocs/.gopurs-evidence/a-20261020-b1probe/`.
Le binaire instrumenté n’est pas installé.
