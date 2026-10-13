# Plan de packages — G01

**Plan canonique vérifié le 13 octobre 2026.** La référence et les budgets
figurent dans [G00](build-go-packages.md). G01 définit le placement ; les
qualifications, l'émission compilable et les mesures Go sont les étapes G02–G06.
Les tailles de fermeture ci-dessous sont des estimations statiques, sans gain
de temps attribué au compilateur.

## Décision et portée

Le premier plan à porter jusqu'à l'émission emploie des **intervalles d'un ordre
topologique canonique**, délimités par des ancres logiques de densité **1/128**.
Il conserve les FFI entières et l'ordre Go d'initialisation par des contraintes
de dépendance. Sur le corpus b8x, cela donne **525 packages**, **35 594 arêtes**,
et au plus **7 461 745 octets de déclarations** dans un package. Runtime,
packages natifs externes, entrées et raccordement d'initialisation s'ajoutent à
ces 525 unités générées.

Le même algorithme donne **158 packages / 6 105 arêtes** sur les **65 145
déclarations** du bootstrap livré ; son plus gros package contient **4 483 254
octets**. Les graphes des quatre états b8x et du bootstrap sont vérifiés par
un second contrôle, indépendant du calcul de SCC/placement.

Les densités **1/256** et **1/512** sont aussi conservées : respectivement
**249 / 123 packages** b8x, avec des maxima de **18 081 522 / 18 427 455 octets**.
La densité 1/128 est le point de départ fixé avant le relevé canonique. Son
acceptation finale dépend toujours des seuils G : **−30 % Go / −20 % chaîne**
sur chaque édition, avec les budgets de froid, mémoire et stockage de G00.

## Identités et propriété

Aujourd'hui, les noms internes sont surtout des chaînes dans `GoAst`,
`GoCode`, les signatures des workers et les FFI. `Driver.Output` écrit tout
dans `purescript`. Le plan introduit une identité logique avant le placement :

- **déclaration nommée** : version, namespace généré/FFI, module PureScript,
  catégorie et noms déclarés ; les groupes `var`/`const`/`type` restent indivisibles ;
- **méthode** : identité du type receveur et nom de méthode, indépendamment du
  nom de la variable receveur et de sa position source ;
- **inscription `StructGetters`** : identité du constructeur enregistré,
  indépendamment d'un numéro `init#N` ;
- **`init` FFI** : ordre local explicite dans l'atome du fichier FFI. Ajouter une
  inscription générée ailleurs ne renumérote pas son identité de placement ;
- **spécialisation et helper** : identité fournie par leur producteur, comprenant
  les arguments de type normalisés lorsque ceux-ci distinguent la déclaration.
  Les helpers Rebox portent notamment la paire ordonnée source/destination.

Le prototype rapproche ces identités des noms Go finaux déjà uniques. G02 doit
les transporter depuis leurs producteurs jusqu'au rendu, sans les reconstruire
à partir de positions, de tailles ou d'octets de corps. Une collision d'identité
ou de nom de package produit une erreur avec les identités concernées ; elle
ne fusionne jamais silencieusement des déclarations.

Chaque déclaration a exactement un propriétaire. Les contraintes suivantes
sont appliquées avant de condenser les cycles :

1. un getter, son `cache_…`, son `once_…` et son worker `Call_…` associé ;
2. les déclarations d'un même fichier FFI ;
3. chaque méthode et son type receveur ;
4. un usage ordinaire de symbole privé et la déclaration de ce symbole ;
5. les arêtes de valeur/type et, conservativement, les contraintes de champs ;
6. l'ordre des initialisations décrit ci-dessous.

Les champs ne sont pas des identifiants à qualifier comme des fonctions :
leurs arêtes restent distinguées des références ordinaires.

## Construction déterministe du DAG

Après les unions de propriété, les SCC sont calculées. Une SCC a pour ancre le
SHA-256 de l'ensemble trié de ses identités logiques, encodé canoniquement. Les
SCC sans fournisseur sont sorties d'une file de priorité triée par cette ancre ;
un consommateur devient disponible lorsque tous ses fournisseurs sont sortis.

Dans cet ordre fournisseurs-avant-consommateurs, une ancre dont les 64 premiers
bits sont divisibles par **128** ouvre un package. Le premier intervalle sans
ancre prend `p_root`, les autres `p_<128 premiers bits de l'ancre>`. Une
collision de nom est détectée avant émission. Grouper des intervalles contigus
préserve l'acyclicité : tout import externe à l'intervalle pointe vers un
intervalle antérieur.

Chemin prévu : `gopurs/output/purescript/<package>`. Dans chaque package, les
fragments gardent le nom du fichier Go d'origine et l'ordre des déclarations ;
cela conserve notamment l'ordre lexical utilisé par Go pour l'initialisation.
L'alias d'import commence par `g` et les six premiers chiffres hexadécimaux du
SHA-256 du chemin complet du propriétaire. Il est allongé jusqu'à être unique
parmi les imports requis et les identifiants du fichier ; une collision avec
un identifiant local après les 64 chiffres est résolue par des `_` supplémentaires.
Une collision de hash complet entre propriétaires est une erreur explicite.
Ce calcul porte sur l'ensemble des imports, sans compteur dépendant de l'ordre
des workers. Les clés d'émission G05 incluront les propriétaires et alias utilisés.

Une modification de corps à graphe identique conserve exactement le placement.
Une fusion/scission ou un ajout/retrait recalcule le plan canonique depuis les
identités présentes ; aucun ancien propriétaire n'est réutilisé par inertie.
Le mouvement des déclarations inchangées est mesuré, pas supposé nul.

La permutation des **185 667 identifiants numériques** du graphe de base, avec
inversion de l'ordre des arêtes et méthodes, redonne les mêmes empreintes de
placement pour les trois densités. Les fixtures vérifient aussi l'insertion
d'une inscription, le renommage local du receveur, la substitution à cardinalité
égale, la rétraction et une collision forcée. **95 graphes** sont comparés à
un oracle de fermeture transitive pour les SCC, dont des cycles sans racine.

## Exports, layouts et bridges

Le plan fournit le propriétaire d'une déclaration ; G02 résout chaque référence
à son symbole, son rôle et son propriétaire. La qualification dépend ensuite
du package consommateur. En mode monolithique, le nom Go historique est rendu
tel quel : le critère G02 reste l'inventaire byte-exact sur les mêmes TAST.

- Les valeurs, types et fonctions traversant une frontière utilisent leur nom
  exporté. Les références ordinaires privées imposent la colocalisation ci-dessus.
  Les sélections de méthode gardent leur syntaxe ; leur receveur appartient au
  même package que la méthode.
- Une déclaration de type nommé reste unique. Les instanciations génériques,
  assertions, casts de constructeurs et helpers Rebox pointent vers cette même
  déclaration. Les types et helpers sont placés par leurs vraies dépendances,
  sans dupliquer un type ni extraire globalement tous les layouts.
- Les records anonymes privés ne peuvent pas simplement franchir un package.
  Pour le layout multipackage, la règle proposée est un nom de champ **exporté
  et injectif par label** : `P_` + label ASCII identifiant ; pour les autres
  labels, `PX_` + unités UTF-16 hexadécimales de largeur fixe. Les deux domaines
  sont disjoints. Cela couvre aussi les labels vides, Unicode, mots-clés et `_`.
  L'ordre des champs reste celui du layout d'origine. Les champs natifs externes
  gardent leur contrat Go et une correspondance explicite avec les labels.
- Les tags runtime des constructeurs, labels PureScript et noms Go sont trois
  identités distinctes. `Rc`/`V0` des constructeurs et les champs de tuples ne
  sont pas renommés comme des labels de records. Les conversions doivent
  conserver représentation, ordre, boxing et types effectifs.
- Les accès par réflexion exigent une correspondance explicite label/champ.
  Le reproducteur G00 confirme qu'un tag de struct ne change pas `FieldByName`.
  Les cinq FFI applicatives recensées, les conversions réflexives du runtime et
  les casts directs PBO sont des validations obligatoires G03/G04.
- Les fragments opaques doivent porter leurs références générées, en plus des
  imports externes. Les références FFI sont réécrites par AST avec leur portée,
  leurs alias et leurs liaisons, jamais par remplacement textuel global.

Ces règles constituent les obligations du rendu G02/G03. Les plans privés ne
constituent pas encore un programme Go multipackage compilé.

## Initialisation : contrainte et raccordement des entrées

Le graphe de référence fournit **2 288 opérations d'initialisation de variables**
et **349 `init`**, dont **348 inscriptions** de dictionnaires. Le plan ajoute
les arêtes d'ordre entre opérations successives, puis entre le dernier RHS et
le premier `init`, et enfin entre les `init` dans leur ordre source Go. La
condensation absorbe les cycles que la colocalisation FFI impose à cet ordre.
Une unité peut apparaître plusieurs fois dans la séquence : les groupes de
déclarations ne permettent pas de supposer un ordre d'évaluation contigu.

Le rendu conserve les RHS et types Go actuels, ainsi que l'ordre relatif des
déclarations/fichiers portant des effets à l'intérieur d'un package. Les entrées
doivent retenir la fin de cette chaîne par un package de raccordement commun,
y compris une entrée qui ne référence qu'une partie des fonctions applicatives.

Les imports externes de l'ancien package commun sont aussi une dépendance
d'initialisation : une FFI devenue inaccessible peut être la seule à importer
un package à effets. Un raccordement externe commun doit conserver cette union
d'imports et précéder les unités générées à effets. Les dépendances Go externes
restent inchangées ; ce raccordement conserve leur présence et leur ordre avant
les effets applicatifs.

Les reproducteurs exécutés avec Go **1.26.8 linux/arm64** vérifient :

- ordre variables/`init` avant et après regroupement ;
- inscription conservée pour une entrée partielle et effets d'un import externe
  référencé uniquement par une fonction inutilisée ;
- **64 appels concurrents** depuis deux packages au même getter : une évaluation,
  un pointeur de stockage unique ;
- ordre Go exact des RHS intercalées, initialisations multiples et variable `_`.

Une séparation naïve et une inscription supprimée produisent les écarts attendus.
Un essai d'initialiseur central a aussi exposé une erreur du harnais : l'ordre
écrit à la main initialisait une dépendance trop tôt. Sa trace est conservée ;
le reproducteur corrigé tire son ordre de `go/types.Info.InitOrder` et passe.
Le premier plan conserve l'initialisation Go dans les packages.

## Portée estimée des trois éditions

Pour le plan canonique 1/128, les fermetures comprennent les packages modifiés,
les anciens/nouveaux placements et leurs consommateurs statiques :

| Édition | Packages directement touchés après édition | Fermeture après édition | Octets de déclarations | Déclarations inchangées déplacées |
| --- | ---: | ---: | ---: | ---: |
| Feuille | 1 | 35 | 38 792 440 | 0 |
| Partagée | 34 | 121 | 63 388 578 | 1 818 |
| FFI | 1 | 129 | 65 252 780 | 0 |

L'édition partagée modifie le graphe et déplace donc des déclarations inchangées.
Le prototype initial à identités ordinales en déplaçait 316 à cette densité ;
le plan canonique en déplace **1 818**. Cette amplification est conservée dans
le bilan. Les identités sémantiques des inscriptions priment sur ce meilleur
résultat ponctuel de l'ancien nommage.

Les plus grosses unités contraintes expliquent une partie du coût résiduel :
le getter de test PostgreSQL avec son cache/once (**6 968 195 octets**), les
tables Unicode (**3 249 712** et **1 274 937 octets**), et le groupe d'ordre des
inscriptions (**4 277 déclarations / 873 866 octets**). Le regroupement final
peut encore élargir les régions affectées. Go peut inversement éviter de
recompiler un consommateur dont les exports importés sont inchangés.

Ces chiffres excluent les nouveaux imports/qualifications et les petites unités
de raccordement. Ils ne remplacent ni une compilation ni les budgets de stockage
et mémoire. Les actions réellement recompilées seront tracées en G03/G06.

## Modèles conservés et preuves

Les expériences privées comprennent les intervalles initiaux, les hauteurs
depuis les fournisseurs, les hauteurs depuis les consommateurs, les bandes
par origine, les FFI par déclaration, l'initialiseur central et les ensembles
de consommateurs terminaux. Les variantes non retenues gardent leurs plans,
fermetures et déplacements ; leurs tailles ne sont pas des gains Go mesurés.

Racine locale :

```text
/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/
  gopurs-go-packages-g01-20261013/
```

- `protocol.md` : paramètres fixés avant chaque famille et décisions ;
- `canonical-interval-models-1/` : quatre états, identités et douze plans ;
- `canonical-verification-1/` : quotient exact, deuxième contrôle de DAG,
  replay avec permutation et trois plans du bootstrap ;
- `identity-contracts-1/` : identités, rétractions, collision et 95 oracles SCC ;
- `initialization-contracts-1/` à `-4/` : programmes et traces ; l'essai `-3`
  rejeté reste documenté, `-4` utilise l'ordre du vérificateur de types ;
- `consumer-set-models-1/rejection.json` : incompatibilité `int.bit_count` avec
  Python 3.9, corrigée par un comptage équivalent avant la reprise `-2`.
- `final-plan/summary.json` : audit de clôture, choix du plan et empreintes.

La suite G02 doit porter ces identités et dépendances dans l'AST, les signatures
et les fragments Go, en gardant le rendu monolithique byte-exact. Le plan ne
devient le layout livré qu'après compilation, parité multi-hôtes, migration du
cache et validation des seuils G.
