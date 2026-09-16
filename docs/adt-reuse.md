# Réutilisation des cellules d’arbres

La passe `Gopurs.Ownership` analyse le `BackendModule` après PBO et la fusion des
thunks, avant l’analyse TCO et la génération des fonctions ordinaires. Elle
produit des fonctions spécialisées qui consomment un arbre possédé exclusivement.
Les fonctions ordinaires gardent leur sémantique persistante.

## Propriété établie

L’entrée d’une fonction spécialisée est une forêt dont les nœuds non vides sont
exclusifs et disjoints. Le résultat conserve cette propriété. Une racine neuve
contenant un enfant partagé ne remplit pas cette condition. La représentation
vide `nil` peut être partagée.

La première implémentation traite les ADT monomorphes définis dans le module,
représentés par un pointeur, avec un constructeur portant des champs et au plus
un constructeur vide. Les champs portent des scalaires, des enums, ou le même
type d’arbre. Les fonctions acceptées sont de premier ordre, avec des paramètres
scalaires ou arbres et un résultat arbre. Les champs opaques, fonctions, tableaux
et conversions de représentation restent hors de cette analyse.

Chaque référence locale de l’IR est résolue dans sa portée. Les alias de
projections désignent un chemin canonique : une racine et une suite d’indices de
champs. Les sous-arbres conservés dans un résultat ou passés à d’autres fonctions
spécialisées doivent former des chemins disjoints : aucun doublon et aucun
chemin préfixe d’un autre. Cela interdit de conserver ensemble un parent et son
descendant, ou de placer le même enfant dans deux champs du résultat.

Les appels sont validés par un point fixe sur les fonctions du module. Retirer
une fonction non prouvée invalide les fonctions qui dépendaient de son contrat.
Les appels inconnus, les FFI, les closures et les effets ne constituent pas des
preuves de consommation. Une fonction peut donc rester persistante même si son
algorithme permettrait une analyse plus précise.

## Ordre des opérations et cellules disponibles

Les lectures de champs nécessaires au résultat sont matérialisées avant les
mutations. Les cellules mortes correspondent aux racines ou aux préfixes de
chemins qui ne font plus partie des sous-arbres conservés. Les branches disposent
de leurs propres preuves ; une garde ne donne pas de droit de mutation sur son
chemin d’échec.

Chaque invocation reçoit éventuellement une cellule supplémentaire à réutiliser.
Cela permet à une insertion de transmettre son ancien parent à la fonction de
rééquilibrage, même lorsque cette dernière reçoit seulement ses champs. Les
cellules disponibles sont retirées du stock lorsqu’elles servent. Lors d’un
`let` consommant, l’environnement perd aussi les anciennes références vers les
cellules données au calcul, avant d’introduire la nouvelle racine du résultat.

Une ancienne racine ou un sous-arbre encore observable dans la continuation
empêche la consommation correspondante. Les appels terminaux récursifs des
fonctions spécialisées deviennent des boucles. Le champ `Rc` existant ne sert
pas de preuve : aucun protocole complet de comptage de références n’est requis
par cette passe.

## Choix des appels

Le générateur peut sélectionner une fonction spécialisée lorsque tous ses
arguments arbres sont des constructions fraîches admissibles, composées
récursivement de telles constructions et de valeurs vides. Une variable locale
ou un enfant emprunté ne suffit pas à établir cette fraîcheur.

Les noms des types, des fonctions et des benchmarks n’interviennent pas dans la
reconnaissance. Les noms générés sont réservés avec leur forme Go et celle de
leurs fonctions auxiliaires, en tenant compte des fonctions et des FFI du module.

## Contrat source et transformations

Le lecteur PBO reconnaît `usageAnalysis` version 1, phase `corefn`, et conserve
la provenance module + identifiant. Un champ absent reste inconnu. Les bornes
supérieures à l’intervalle `Int` du lecteur deviennent inconnues, sans troncature.

Avant la monomorphisation et la conversion PBO, ces identités et faits source sont
invalidés. Les copies et spécialisations ne peuvent donc pas transporter de
certificat périmé. La passe Go recalcule ses preuves sur les références lexicales
et les chemins de l’IR final ; elle ne transforme pas un ancien `usageCount`, un
`lastLocalUse` ou un marqueur d’échappement en preuve d’exclusivité.

## Vérification

Les tests `tools/owned-trees.test.mjs` construisent des IR ciblés pour vérifier
l’acceptation des chemins disjoints et le refus des alias, des continuations
encore observables, des gardes et des collisions de noms. La fixture
`tests/passing/OwnedTrees.purs` vérifie quatre rotations, l’ordre des clés, les
couleurs, la hauteur noire, les doublons, les anciennes versions et les enfants
partagés, avec des noms différents du benchmark.

Le build et le bundle passent sans avertissement. Les 64 tests des outils
Gopurs passent, dont 21 tests de possession. La fixture `OwnedTrees` compile et
s’exécute ; le contrôle final du snapshot passe sans le réécrire. Les 10 tests
du contrat source et les 41 tests PBO existants passent également.

Les mesures avant/après utilisent le [protocole et les artefacts archivés](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-adt-reuse-validation-20260916/PROTOCOL.md).
Pour 100 000 clés, les médianes passent de 2 483 949 à 100 000 allocations et
de 79 486 352 à 3 200 000 octets cumulés. Dans la campagne à trois processus
par variante, RBTree passe de 24,20 à 9,95 ms, contre 8,92 ms pour le manuscrit.
Les baselines historiques du README d’altbak restent distinctes des mesures
contrôlées de cette campagne. Le [bilan complet](/Users/0x1/Documents/htdocs/gopurs/gopurs/todo.md)
précise les limites et les commandes de reproduction.
