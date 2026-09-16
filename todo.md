# Gopurs — contrat d’usage et réutilisation des nœuds

Plan terminé le 16 septembre 2026. Le lot 1 a été validé par l’utilisateur. Les lots 2 à 5 sont implémentés et vérifiés pour le périmètre conservateur décrit ci-dessous : arbres monomorphes, fonctions de premier ordre et entrées dont toute la forêt est exclusive.

Le 17 septembre, le contrat est simplifié à la demande de l’utilisateur : seuls `bindingUsage` et `variableUse` sont conservés, sans marqueur de version ou de phase. Les champs historiques `usageCount` et `escapes` sont retirés du producteur et des lecteurs PBO de Gopurs/Purust. Cette migration est distincte des mesures du 16 septembre ; la nouvelle compilation Haskell reste à faire par l’utilisateur. Côté Gopurs, build/bundle sans avertissement, 12 tests du contrat direct, 41 tests PBO et 64 tests outils passent. Le snapshot et l’exécution de `OwnedTrees` passent avec le compilateur TAST installé. Les JSON déjà présents ne sont pas réécrits.

## Objectif et ordre de réalisation

Réduire les allocations du Go généré en réutilisant les cellules dont la modification ne peut affecter aucune autre référence observable. Le premier cas concret est le Red-Black Tree d’altbak.pub. Les optimisations doivent découler des propriétés du programme, sans reconnaître le nom du benchmark.

1. **Haskell :** produire un contrat explicite distinguant usages des bindings et des occurrences.
2. **PBO de Gopurs :** lire ce contrat et maintenir sa validité après transformation.
3. **Analyse sur l’IR transformé :** établir les alias, le partage des champs et les conditions d’exclusivité.
4. **Gopurs :** réutiliser les cellules dans les chemins où ces conditions sont prouvées.
5. **Validation :** contrôler la persistance et les invariants, puis mesurer allocations et temps.

Le TAST enrichi fournit déjà les types, `dataDecls`, `classDecls`, `ForAll`, contraintes, `TypeApp`, champs et queues de rangées. Les nouvelles informations complètent ce typage. Elles doivent distinguer trois propriétés : nombre d’usages d’un binding, dernier usage d’une référence locale, exclusivité de l’objet mémoire.

Un dernier usage peut aider à transférer une référence ou éviter un clone tout en laissant d’autres références vers l’objet. Il ne suffit pas à autoriser une mutation destructive.

## Lot 1 — Contrat d’usage Haskell

**Livraison attendue :** un schéma précis, son calcul conservateur, sa lecture/écriture JSON et des tests des propriétés annoncées. Le générateur Go et les optimisations de Purust ne sont pas modifiés dans ce lot.

**État :** code, documentation et tests livrés ; compilation et tests Haskell pris en charge puis validés par l’utilisateur après correction des avertissements et examen des nouveaux snapshots. Aucun build ou test Haskell supplémentaire lancé par l’agent. Le JSON réel de `.test_modules/TastCoverage/corefn.json`, produit avec le contrat version 1, a ensuite été décodé avec succès par le nouveau lecteur PBO.

### 1.1. Définir les faits et leur portée

- [x] Distinguer les informations du binding de celles de chacune de ses occurrences.
- [x] Définir les usages par **instance du binding**, dans la portée analysée. Un paramètre peut avoir un usage par invocation alors qu’une variable extérieure capturée peut servir à plusieurs appels de closure.
- [x] Représenter explicitement l’inconnu. Un champ absent ne devient jamais zéro usage ou preuve de non-échappement.
- [x] Conserver le caractère contextuel de l’indication d’échappement ; ne pas la présenter comme une analyse des alias mémoire.
- [x] Définir `lastLocalUse` sur les chemins d’exécution, y compris gardes et captures. Un compteur décrémenté pendant un parcours d’AST ne constitue pas à lui seul une preuve suffisante.

### 1.2. Schéma JSON implémenté

Les exemples sont des fragments ; les autres champs TAST de chaque nœud sont conservés. Aucun bloc `usageAnalysis`, numéro de contrat ou champ de phase n’est émis. Les lecteurs interprètent directement les blocs présents sur les annotations.

Sur l’annotation d’un binding local ou du paramètre lié par une abstraction :

```json
{
  "annotation": {
    "bindingUsage": {
      "bindingId": 17,
      "maxUses": 1,
      "hasEscapingUseContext": false
    }
  }
}
```

Sur l’annotation d’une occurrence de cette variable :

```json
{
  "annotation": {
    "variableUse": {
      "bindingId": 17,
      "lastLocalUse": true
    }
  }
}
```

Cas où l’analyse ne fournit pas de borne, par exemple certaines captures réutilisables :

```json
{
  "bindingUsage": {
    "bindingId": 17,
    "maxUses": null,
    "hasEscapingUseContext": true
  }
}
```

| Champ | Contrat |
| --- | --- |
| `bindingId` | Identifiant unique du binding dans le module exporté, attribué avec résolution des portées lexicales. Les occurrences locales désignent cet identifiant. Il ne s’agit pas d’une adresse d’objet ni d’un identifiant persistant entre versions du programme. |
| `maxUses` | Entier positif ou nul : borne supérieure des usages directs de cette instance du binding dans la portée analysée. `null` : borne non établie. Zéro n’est émis que lorsque l’absence d’usage est établie. |
| `hasEscapingUseContext` | `true` : un usage dans un contexte classé potentiellement échappant a été rencontré ; `false` : aucun tel contexte n’a été détecté dans l’analyse de ce binding ; `null` : information inconnue. Aucune valeur ne prouve l’exclusivité de l’objet ou de ses champs. |
| `lastLocalUse` | `true` : après cette occurrence, aucun usage direct ultérieur de cette instance du binding n’est possible sur les chemins concernés. `null` : propriété non établie. Ce champ ne décrit pas les autres alias vers la valeur. |

L’absence d’un bloc ou d’un champ optionnel équivaut à une information inconnue. Un bloc présent exige un identifiant valide ; un champ invalide n’est pas silencieusement accepté. Les références globales/importées conservent leur qualification existante ; aucun compte global n’est déduit d’un comptage local incomplet, notamment pour les exports. Ces faits décrivent le CoreFn source et doivent être invalidés ou recalculés après transformation.

- [x] Formaliser ce schéma et les nœuds auxquels chaque bloc s’applique, notamment paramètres d’abstractions, bindings de `let`, binders de patterns et occurrences `Var`.
- [x] Implémenter la vérification de l’unicité et de la résolution des identifiants, y compris en présence de noms identiques dans des portées distinctes et de groupes récursifs.
- [x] Définir les invariants des valeurs sérialisées et des champs absents ; une borne négative n’est pas une valeur valide de `maxUses`.

### 1.3. Implémentation Haskell

- [x] Dans `CoreFn/Ann.hs`, remplacer le couple opaque d’information d’usage par des types distinguant faits du binding, faits d’occurrence et inconnus, sans refondre les autres annotations.
- [x] Depuis `CoreFn/Usage.hs`, appeler `CoreFn/Usage/Analysis.hs` pour attribuer les identifiants et calculer les faits selon le contrat. Sommes séquentielles, branches exclusives, gardes successives et répétitions ont des règles explicites.
- [x] Tenir compte des captures et de la récursion. Si le nombre de réutilisations n’est pas établi, garder une borne inconnue ; ne pas déduire un dernier usage de l’ordre textuel de définition d’une closure.
- [x] Établir les derniers usages par une analyse arrière des chemins ; garder la propriété inconnue lorsque les informations disponibles sont insuffisantes.
- [x] Adapter `CoreFn/ToJSON.hs`, `CoreFn/FromJSON.hs` et documenter le point d’appel dans `Make/Actions.hs`.
- [x] Ne produire aucun certificat `unique`, `canMutate`, de fraîcheur de résultat ou d’exclusivité transitive à partir de ces seuls faits d’usage.

### 1.4. Migration du format d’usage

- [x] Retirer `usageCount` et `escapes`, initialement conservés pendant la migration additive. Les lecteurs Gopurs/Purust ne les interprètent plus.
- [x] Émettre les nouveaux blocs directement, sans marqueur racine, et conserver les autres champs du TAST à l’identique. Le lecteur Haskell résout les références à `typeTable` et lit les expressions `TypeApp`, nécessaires aux tests d’aller-retour typés.
- [x] Lire les anciens JSON sans inventer de nouveaux certificats à partir des champs historiques. En particulier, un ancien marqueur d’occurrence ne devient pas automatiquement une preuve conforme au nouveau contrat.
- [x] Vérifier que les lecteurs des forks PBO existants acceptent les champs supplémentaires : les artefacts déjà compilés Gopurs/Purust donnent le même module décodé pour un JSON `Test.Fib` enrichi en mémoire. Cela vérifie les champs ignorés, pas le futur consommateur du contrat.
- [x] Couvrir l’aller-retour du nouveau format par le lecteur Haskell et la compatibilité avec le format historique dans les tests confiés à l’utilisateur et validés par lui.
- [x] Ne jamais reprendre pour `maxUses` l’ancien comportement qui convertissait un `usageCount` absent en `0`.

L’implémentation des nouvelles optimisations des backends n’est pas une condition de compatibilité du lot 1. Leur absence doit seulement laisser les informations supplémentaires inutilisées.

### 1.5. Tests du contrat

- [x] Écrire les cas ci-dessous dans `tests/TestCoreFn.hs`, avec tests supplémentaires des annotations périmées, identifiants invalides et bornes négatives/fractionnaires. La suppression des anciens champs et du marqueur racine est également couverte.
- [x] Compilation et tests Haskell validés par l’utilisateur ; JSON réel du compilateur reconstruit inspecté et accepté par le nouveau lecteur PBO.

| Cas témoin | Propriété à vérifier |
| --- | --- |
| Variable inutilisée, utilisée une fois, puis deux fois | Bornes locales respectivement établies, avec distinction entre compte du binding et dernier usage d’une occurrence. |
| Paramètre utilisé une fois dans une fonction appelée plusieurs fois | Compte par instance du paramètre, sans confusion avec un total global. |
| Une utilisation dans chacune de deux branches exclusives | Borne par chemin et derniers usages propres aux branches. |
| Garde qui échoue avant une autre utilisation | Inclusion du chemin qui évalue la garde puis poursuit vers une alternative suivante. |
| Utilisation après un `case` | Une occurrence dans une branche ne devient pas un dernier usage si le binding est réutilisé ensuite. |
| Capture dans une closure réutilisable | Aucun dernier usage déduit de la seule occurrence textuelle ; borne inconnue si la répétition n’est pas établie. |
| Noms identiques dans des portées distinctes | Identifiants distincts et occurrences rattachées au bon binding. |
| Récursion et groupes de bindings récursifs | Traitement conservateur des répétitions et résolution correcte des identifiants. |
| Valeur globale exportée ou référence importée | Aucune borne globale nulle déduite de l’absence d’usage local. |
| Alias local, arbre conservé par l’appelant, sous-arbre retourné | Les faits locaux ne sont jamais présentés comme une preuve d’exclusivité mémoire. |
| JSON historique, nouveau ou incomplet | Les anciens champs sont ignorés ; les nouveaux blocs sont validés sans marqueur racine ; l’inconnu ne devient jamais un fait positif. |

Exemple de limite à conserver dans la documentation et les tests :

```purescript
child t = case t of
  Node x -> x
```

Le scrutinee `t` peut être utilisé une seule fois sans contexte directement classé échappant, tandis que son champ `x` est retourné. Son éventuel `maxUses = 1` et `hasEscapingUseContext = false` ne doivent donc pas autoriser une mutation.

**Critère de fin du lot 1 :** schéma et types explicites, tests des chemins et captures réussis, compatibilité vérifiée, exemples JSON réels inspectés. Aucune mutation destructive n’est activée par ce lot.

## Lot 2 — Lecture et validité après PBO

**État : terminé.** Le [PBO utilisé par Gopurs](/Users/0x1/Documents/htdocs/purescript-backend-optimizer-gopurs/CORE_FN_USAGE.md) lit explicitement le contrat et l’invalide aux frontières de transformation.

- [x] Ajouter `sourceUsage :: Maybe SourceUsage` aux annotations ; lire directement les blocs `bindingUsage` et `variableUse`, sans condition de version ou de phase.
- [x] Garder la provenance module + identifiant lors de la lecture, vérifier les placements, la résolution lexicale et l’unicité, y compris les masquages sans métadonnées.
- [x] Conserver l’inconnu : absence ou `null` ne donnent aucun certificat. Une borne entière dépassant l’intervalle `Int` devient inconnue ; les bornes négatives ou fractionnaires sont rejetées.
- [x] Invalider complètement identités et faits source avant la collecte des corps de monomorphisation de Gopurs, à la sortie du monomorphiseur PBO et à l’entrée de `Convert.toBackendModule`.
- [x] Utiliser les identités lexicales de l’IR final pour les nouvelles preuves. Les niveaux sont interprétés dans leur fonction et leur portée, pas comme des identifiants globaux.
- [x] Vérifier la lecture, l’invalidation immuable et l’absence de ces faits dans le résultat converti ; les tests existants de spécialisation restent passants.

**Choix par rapport au plan initial :** aucun remappage des identifiants source copiés n’est nécessaire : la pipeline Gopurs les efface avant la copie ; le monomorphiseur PBO les invalide aussi en sortie lorsqu’il est appelé directement. Les preuves de possession sont recalculées après transformation. Ce choix évite de faire dépendre la sûreté de l’inlining de faits devenus périmés. Les anciens `usageCount`, `escapes` et les nouveaux derniers usages ne sont pas promus en preuve d’exclusivité.

**Validation :** 10 tests du contrat source et 41 tests PBO existants passent ; le lecteur accepte également un JSON réel produit par le Haskell reconstruit.

## Lot 3 — Analyse du partage et de la possession sur l’IR transformé

**État : terminé pour le sous-ensemble pris en charge**, dans [Gopurs.Ownership](/Users/0x1/Documents/htdocs/gopurs/gopurs/src/Gopurs/Ownership.purs).

- [x] Analyser l’IR après PBO et fusion des thunks. Résoudre chaque alias en racine + chemin de champs dans son environnement lexical.
- [x] Calculer les lectures nécessaires, les sous-arbres conservés dans le résultat ou consommés par un appel, les cellules mortes et les références encore observables dans la continuation.
- [x] Exiger que les sous-arbres transférés soient disjoints : ni chemin dupliqué, ni parent conservé avec son descendant. Une racine fraîche contenant un enfant partagé est refusée.
- [x] Valider par point fixe les familles de fonctions, y compris récursives, qui préservent une forêt exclusive. Les appels inconnus, FFI, closures et effets sont refusés.
- [x] Distinguer les cellules parents mortes des sous-arbres conservés lors des rotations et des appels intermédiaires. Retirer de l’environnement les références vers toutes les cellules consommées ou données à un appel.
- [x] Vérifier dans l’IR et le Go de RBTree que `depth` lit les enfants et retourne un scalaire, sans mutation ni conservation de références. Il reste appelé normalement après la construction ; cette inspection n’est pas une passe générale de résumés d’observateurs.
- [x] Prouver les entrées par des constructions récursivement fraîches ou vides. `nil` peut être partagé ; les paramètres empruntés gardent le chemin persistant.

Le contrat conditionnel est : **forêt d’entrée exclusive et disjointe → forêt de résultat exclusive et disjointe**. L’analyse utilise les types et les corps disponibles après PBO, sans reconnaître un nom de benchmark. Les tests utilisent également un autre ADT et un autre ordre de champs.

**Limites explicites :** ADT monomorphes du module avec un constructeur portant des champs et au plus un constructeur vide ; champs scalaires, enums ou récursifs du même type ; fonctions de premier ordre retournant cet arbre. Pas d’analyse générale des conteneurs opaques, des closures, des résumés importés ou de tous les observateurs. Les cas hors périmètre restent persistants. Les optimisations Records de Purust sont indépendantes de ce travail.

## Lot 4 — Génération Go avec réutilisation des cellules

**État : terminé**, intégré avant la génération des fonctions ordinaires.

- [x] Générer des fonctions internes consommantes qui recyclent les cellules mortes et peuvent transmettre une cellule donneuse au rééquilibrage.
- [x] Matérialiser les lectures avant les mutations, traiter les branches séparément et conserver le contrôle des gardes. Convertir les appels récursifs terminaux en boucles.
- [x] Sélectionner ces fonctions seulement aux entrées fraîches prouvées ; conserver les signatures et la sémantique persistante des fonctions publiques pour les arbres empruntés ou partagés.
- [x] Réserver les noms internes avec leur forme Go, y compris face aux FFI et aux collisions après normalisation des noms.
- [x] Examiner le Go réel : cinq fonctions consommantes pour RBTree (`makeBlack`, `balance`, `ins`, `insert`, `buildTree`) ; `act` sélectionne la construction depuis `nil`.
- [x] Mesurer une médiane de 100 000 allocations pour 100 000 clés, contre 2 483 949 auparavant. Le champ `Rc` n’est jamais utilisé comme preuve d’exclusivité.

Une invocation externe du `buildTree` public reste persistante. La sonde mesure un adaptateur depuis `nil` vers le même point d’entrée consommant que celui choisi dans le `act` généré. Les tests d’anciennes versions passent par les fonctions publiques.

## Lot 5 — Correction et mesures comparables

**État : terminé.** Protocole, sources, binaires, empreintes et sorties conservés dans [PROTOCOL.md](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-adt-reuse-validation-20260916/PROTOCOL.md).

- [x] Vérifier quatre rotations, ordre BST, couleurs, hauteur noire, profondeur, doublons, insertions ascendantes, descendantes et mélangées.
- [x] Vérifier les anciennes versions, 32 snapshots, les enfants retenus et une racine fraîche contenant un enfant partagé.
- [x] Comparer structure et couleurs à la version persistante ; vérifier l’absence de cycles et d’alias ainsi que la construction de 100 000 nœuds.
- [x] Mesurer allocations Runtime et octets cumulés. Ces compteurs incluent le bruit du Runtime ; ils ne constituent pas une instrumentation des seuls constructeurs.
- [x] Comparer avant/après/manuscrit dans trois processus par variante, ordre tournant, sans compilation simultanée : Go 1.27 darwin/arm64, `GOGC=800`, PGO désactivé.
- [x] Valider les 126 résultats numériques des neuf suites et conserver séparément les baselines historiques du README.

| Campagne contrôlée finale | Compilé avant | Compilé après | Manuscrit |
| --- | ---: | ---: | ---: |
| Suite complète, somme des médianes par test | 27,51207 ms | 13,21992 ms | 11,27350 ms |
| RBTree, médiane du harnais | 24,19912 ms | 9,94929 ms | 8,92221 ms |
| Allocations RBTree, médiane de la sonde | 2 483 949 | 100 000 | 100 000 |
| Octets cumulés RBTree, médiane de la sonde | 79 486 352 | 3 200 000 | 3 200 000 |

Les sondes d’allocations comportent 15 échantillons par variante, avec GC préalable. Après optimisation, les valeurs varient entre 100 000 et 100 001 allocations, et entre 3 200 000 et 3 200 016 octets. La profondeur reste 22. Les durées de ces sondes suivent un protocole différent du harnais et ne sont pas mélangées au tableau temporel ci-dessus.

**Vérifications réalisées :** build et bundle sans erreur ni avertissement ; 10 nouveaux tests PBO et 41 existants ; 64 tests outils Gopurs, dont 21 tests de possession ; fixture `OwnedTrees` compilée et exécutée avant/après, puis snapshot contrôlé en mode normal sans mise à jour ; harnais Go de persistance et de consommation. La totalité des autres fixtures de compilation Gopurs n’a pas été relancée.

## Contexte factuel conservé

Dans le [README examiné](/Users/0x1/Documents/htdocs/altbak.pub/README.md:55), le compilé Go affichait 24,02 ms et le manuscrit 11,09 ms, contre 27,44 ms auparavant pour ce dernier. RBTree passe de 24,856 ms à 8,75242 ms dans le manuscrit, tandis que la colonne compilée affiche 20,93562 ms. RBTree représente environ 98,5 % de la baisse historique du total manuscrit.

Le commit altbak.pub `6e85c35b1` du 9 septembre a remplacé un pool de dix millions de cases et la reconstruction des nœuds par des insertions/rotations en place. Le [manuscrit actuel](/Users/0x1/Documents/htdocs/altbak.pub/src/Test/RBTreeFFICheatcode.go:69) crée une cellule par clé distincte. Le Go généré avant ce travail utilisait déjà structs typées et appels directs, mais reconstruisait le chemin d’insertion.

L’[audit historique du 8 septembre](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-allocation-audit-20260908/AUDIT.md:65) comptait environ 2,48 millions d’allocations et 79,5 Mo cumulés pour RBTree compilé, principalement dans `balance`. Ce n’est pas une nouvelle mesure de l’artefact actuel.

Les 11,09 ms ont été publiées le 15 septembre dans `e0eda4141`, en conservant les 24,02 ms compilées. Elles correspondent à un run archivé ; la campagne distincte à trois processus donne environ 11,72 ms. Horloge et protocole ont aussi évolué. Les chiffres historiques ne permettent pas d’attribuer chaque milliseconde à une transformation isolée, et les 11 ms ne sont pas une garantie pour ce plan.

## Livraison et reproduction

Le [contrat Haskell](/Users/0x1/Documents/htdocs/purescript/CORE_FN_USAGE.md), le [contrat lecteur PBO](/Users/0x1/Documents/htdocs/purescript-backend-optimizer-gopurs/CORE_FN_USAGE.md) et la [description de la passe Go](/Users/0x1/Documents/htdocs/gopurs/gopurs/docs/adt-reuse.md) explicitent leurs garanties respectives. Les cinq lots sont livrés dans le périmètre ci-dessus. Le README officiel d’altbak n’a pas été réécrit avec les résultats de cette campagne.

Depuis `gopurs/gopurs`, les vérifications ciblées peuvent être reproduites avec :

```sh
npm run build --silent
GOCACHE=/private/tmp/gopurs-adt-reuse-gocache node --test --test-concurrency=1 tools/*.test.mjs
PATH="$HOME/.local/bin:$PATH" ./bin/test OwnedTrees --keep-workspace
```

Les commandes de compilation et de mesure d’altbak ainsi que l’adaptation des sondes sont décrites dans le protocole archivé. Les extensions à d’autres formes d’ADT, aux closures ou à une analyse générale des observateurs demanderaient des preuves supplémentaires ; elles ne sont pas activées implicitement.

## Essai du 17 septembre — recoloration avec copies conditionnelles

Un prototype limité au Go généré de `insert` et `makeBlack` omet les copies des enfants et de la clé lorsque la cellule réutilisée est précisément le nœud d’origine. Un donneur distinct reçoit toujours tous les champs ; la sélection des cellules, la couleur et `Rc` restent inchangés. Aucun changement correspondant n’est intégré au générateur.

Les huit tests Go passent avant/après, dont deux cas supplémentaires avec donneur distinct. Sur cinq paires de processus alternées, cinq appels mesurés par processus, 100 000 clés et `GOGC=800`, la variation appariée médiane est de **+0,055 %**, avec des paires allant de −1,468 % à +2,179 %. Les médianes des 25 échantillons sont **9,992250 → 10,050541 ms**. Les allocations restent à **100 000 / 3 200 000 octets** : aucun gain temporel reproductible ne justifie cette garde supplémentaire.

Le [prototype et ses mesures](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-field-updates-20260917-prototype/REPORT.md) sont conservés pour éviter de reprendre cette variante sans nouvelle hypothèse. Ces durées de sonde ne remplacent ni le total du README ni les mesures du harnais complet. La simplification de la sélection des cellules et des branches de rééquilibrage reste une piste distincte à mesurer.

## Livraison du 17 septembre — sélection statique des cellules

**État : intégré et validé.** Le générateur distingue les cellules mortes dont la non-nullité est prouvée par les captures de champs des cellules nullable. Il consomme les premières directement, sans balayage ni garde d’allocation, puis garde le repli existant pour les autres. Le stock restant est transmis entre les arguments frères, les constructions et le choix du donneur des appels, y compris récursifs terminaux. Les lectures scalaires conditionnelles ne donnent aucune preuve supplémentaire.

À CoreFn identique (301 fichiers), seul le Go de RBTree change parmi les 388 fichiers produits : les corps consommants de `balance` et `makeBlack`. Le fichier passe de 5 141 à 3 932 lignes. Les types, le runtime, les FFI et les autres modules sont identiques.

| Nouvelle campagne contrôlée | Compilé avant | Compilé après | Manuscrit |
| --- | ---: | ---: | ---: |
| Suite complète, somme des médianes par test | 12,58267 ms | 12,26888 ms | 10,59596 ms |
| RBTree, médiane du harnais | 9,40708 ms | 9,09525 ms | 8,28300 ms |

Trois processus par variante, ordre tournant, 126 résultats numériques validés : le gain total est de **2,49 %**, dont presque toute la baisse vient de RBTree (**3,31 %**). Dans la sonde séparée à cinq paires, quatre sont favorables ; la variation appariée médiane est de **−3,53 %**, et les médianes des 25 appels passent de **9,370458 à 9,080833 ms** (−3,09 %). Les allocations restent à une médiane de **100 000 / 3 200 000 octets**. Les chiffres de sonde et de harnais ne sont pas mélangés.

**Vérifications :** build et bundle sans erreur ni avertissement ; 64 tests outils existants et quatre nouvelles régressions exécutant le Go généré ; fixture `OwnedTrees` compilée et exécutée, snapshot mis à jour après inspection puis contrôlé sans réécriture ; huit tests Go avant et huit après. Les nouvelles régressions couvrent les captures avant réemploi du parent, les arguments frères, le repli nullable et la consommation du donneur d’un appel récursif terminal. Les autres fixtures de compilation n’ont pas été relancées.

Le [rapport, les sources et les mesures brutes](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-static-cells-20260917/integration/REPORT.md) conservent le protocole et les empreintes. Cette campagne ne remplace pas la baseline officielle actuelle du [README d’altbak](/Users/0x1/Documents/htdocs/altbak.pub/README.md:55), **13,01 ms compilées / 11,09 ms manuscrites** ; le README reste inchangé.
