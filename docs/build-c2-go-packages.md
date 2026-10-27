# C2-5 — planification Go incrémentale

**Statut au 23 octobre 2026 : overlay activé par défaut par le lanceur Go natif,
opt-in pour JS/Rust, avec relocalisation d'un module simple ou conservation de
ses sections dans plusieurs paquets.** Le cas
central validé porte sur une modification de libellé dans
`Util.Type.Variant.Encoding.Encoding.Test.DecodeJsonWith`. L'acceptation d'une
édition dépend des preuves décrites en fin de document ; un corps modifié
n'est pas à lui seul une preuve de réutilisation.

Les anciens tableaux Ping ci-dessous décrivent leurs campagnes historiques.
Ils ne valident pas les performances de la reprise actuelle du cache de
préparation. Les résultats backend de cette reprise et la cible ≤20 s sont
suivis dans [build-c2-persistent.md](build-c2-persistent.md).

## Principe

Ne pas recopier les anciens fichiers Go « en espérant ». Le cache doit reposer
sur :

- un plan de packages stable (`planHash`) ;
- une correspondance module → bundle ;
- des signatures/types inchangés ;
- C1c pour le module modifié ;
- un fallback full dès qu'une preuve manque.

## Phase 1 — Cache de plan

- persister `package-result.json`, `planHash`, les bundles Go et la table
  module → bundle ;
- stocker les hashes des sources Go générées ;
- au build suivant, comparer sources et planHash.

**Gate :** hit de plan sans modification → bundles identiques, aucune
re-planification.

## Phase 2 — Détection delta

- identifier les modules Go modifiés ;
- vérifier que leurs déclarations/types/exports sont inchangés ;
- refuser le delta si signature, type, import ou plan change.

**Gate :** une édition de corps n'est acceptée qu'après les vérifications de
signatures, de références et de placement. Une preuve manquante entraîne le
planificateur complet.

## Phase 3 — Rendu ciblé

- réutiliser les bundles des packages inchangés ;
- ne rendre que le package modifié ;
- intégrer C1c : wrapper public stable, package d'implémentation, bridge ;
- ne jamais modifier le package public pour un corps seul.

**Gate :** sorties identiques au full pour Ping et FFI.

## Phase 4 — Intégration driver

- brancher le cache dans `writePackageFiles`/`BuildCache.packages` ;
- conserver le one-shot comme oracle ;
- fallback explicite et journalisé.

**Gate :** backend packages ≤20 s sur la feuille, sans faux hit.

## Phase 5 — Validation

- trois paires alternées ;
- scénarios Ping, FFI, shared, multi-éditions ;
- quatre hôtes et sorties byte-exactes ;
- froid non prioritaire pour l'instant, mais non régressé.

## Critères d'arrêt

- gain <2 s → arrêt ;
- sortie différente du full → arrêt ;
- planHash instable → arrêt ;
- cache qui ne peut pas être invalidé → arrêt.

## État actuel

Les phases 1 et 3.1–3.5 sont codées. Le cache de plan persiste sous
`.gopurs-cache/go-packages/plan-v1/` et l'overlay C1c est branché derrière
`GOPURS_GO_PACKAGE_OVERLAY=1`, avec `GOPURS_GO_PACKAGES=1`. Le planificateur
complet reste l'oracle et le fallback par défaut. Le chemin `retained-N` ajouté
le 23 octobre conserve les paquets existants d'une source répartie entre
plusieurs bundles. Il complète le premier chemin de relocalisation décrit en
phase 3.5 ; il ne reprend pas la conception API/bridge historique ci-dessous.


## Contrat de cache — phase 1

Un hit de plan ne peut être déclaré que si **toutes** les conditions suivantes
tiennent.

### Clé

- `planHash` du planificateur ;
- hash de chaque source Go générée, par module ;
- hash des exports/imports et des signatures de déclarations générées ;
- digest FFI par module ;
- version du format de cache ;
- options de layout et version Go.

### Données persistées

Sous `.gopurs-cache/go-packages/plan-v1/` :

- `package-result.json` (planHash, packages, files, entries) ;
- les bundles Go, nom + SHA-256 ;
- la table `module -> package/bundle` ;
- les hashes de sources et de signatures ;
- les hashes des fichiers de sortie.

Écriture atomique : staging privé, `fsync`, puis rename. Aucun fichier public
n’est modifié pendant l’écriture.

### Invalidation

Le cache est **miss** si :

- `planHash` change ;
- un module change de signature, type, export, import ou digest FFI ;
- un bundle manque ou son hash diffère ;
- la table module → bundle est incomplète ;
- la version/options/Go changent ;
- un fichier de sortie est absent ou corrompu.

Un simple changement de corps **ne suffit pas** à lui seul : la phase 1 doit
encore vérifier les signatures. Le rendu ciblé arrive en phase 3.

### API

- `writePlanCache` après publication réussie ;
- `readPlanCache` au début de `preparePackageSession` ;
- mode one-shot conservé comme oracle ;
- aucun skip du planificateur en phase 1 : on valide seulement le contrat.

### Tests phase 1

- base/feuille : planHash stable, un seul bundle modifié ;
- cache valide relu sans erreur ;
- cache corrompu → miss ;
- signature modifiée → miss ;
- options/Go modifiés → miss ;
- sorties du build complet inchangées.

### Gate phase 1

Le cache est relu et validé, mais **aucun gain de temps n’est attendu** avant
la phase 3. Tout hit non prouvé est un échec.

### État d’implémentation

- [x] Persistance de `request.json` et `result.json` sous
  `.gopurs-cache/go-packages/plan-v1/`.
- [x] Lecture/validation des options/sources/entrées au build suivant
  (`validated`, `miss-sources`, `miss-entries`, `miss-result`).
- [x] Test unitaire du contrat de base (`TestPlanCacheStatus`).
- [x] Hash de signatures Go par fichier (déclarations/types, corps exclus).
- [x] Invalidation sur signatures modifiées (`miss-signatures`).
- [x] Test unitaire signature changée / cache corrompu.
- [x] Statut `eligible-body` : source modifiée, signatures identiques.
- [x] Skip du planificateur et rendu ciblé (phase 3.5, opt-in
  `GOPURS_GO_PACKAGE_OVERLAY=1`, validation des sorties en cours).


## Phase 3 — conception historique (C1c automatique)

Cette proposition API/implémentation/bridge n'a pas été réalisée telle quelle.
Les sections 3.1–3.5 retracent les étapes du premier overlay ; le comportement
actuel des modules répartis sur plusieurs paquets est décrit dans la reprise
du 23 octobre. Les mentions « reste à faire » ci-dessous correspondent à ces
étapes historiques.

### Entrées

- `package-result.json` précédent (planHash, bundles, entries) ;
- table `module -> package` reconstruite depuis les marqueurs
  `gopurs-source-v1` des bundles ;
- sources Go générées du build courant ;
- signatures Go par module (phase 1) ;
- digest FFI par module.

### Conditions d’éligibilité

Le premier lot phase 3 est codé : `planCacheStatus` distingue
`validated`, `eligible-body` (corps seul) et `miss-signatures`. Aucun skip
n’est encore actif.

Le delta n’est accepté que si :

- `planHash` inchangé ;
- seuls des modules dont les signatures/types/exports/imports sont inchangés
  ont changé ;
- aucun constructor/type/class/init global n’est modifié ;
- aucun bundle manquant ou hash invalide.

Sinon : fallback full, immédiat.

### Découpage API / implémentation

Pour chaque module modifié :

- **package public stable** : conserve les noms, types, signatures et exports ;
- **package d’implémentation** : contient les corps modifiés ;
- **bridge stable** : variables de fonctions typées, registre d’initialisation ;
- **package de types stable** : types/constructeurs partagés, champs exportés
  ou accesseurs ;
- wrapper public : appelle le bridge, jamais l’implémentation directement.

Les types restent définis une seule fois ; le public et l’implémentation les
importent depuis le package de types.

### Initialisation

- l’implémentation s’enregistre dans `init` ;
- le `main` importe le package d’implémentation pour garantir l’init ;
- l’ordre d’init est explicite et testé ;
- pas de wrapper appelé avant l’enregistrement.

### Réutilisation des bundles

- copier les bundles des packages inchangés ;
- ne remplacer que le bundle contenant le module modifié ;
- ne jamais copier un bundle dont le hash ou la signature ne correspond pas ;
- conserver les entrées et le `init_all`.

### Fallback

Full planner si :

- type/constructor/signature/import/FFI change ;
- plusieurs packages sont affectés ;
- méthode, valeur globale ou init complexe non couverte ;
- bridge impossible à typer ;
- test de sortie différent.

### Tests

- Ping body-only : plan réutilisé, Go build réduit ;
- FFI : bridge + type public ;
- type/constructor changé : fallback full ;
- cache corrompu/planHash différent : fallback ;
- 27 sorties exactes, trois paires, quatre hôtes.

### Gate

Backend packages **≤20 s**, puis chaîne complète ≤30 s avec Go/C1c. Si le delta
n’est pas exact ou si le plan n’est pas réutilisable, arrêt.


## Phase 3.1 — manifeste C1c (fait)

Premier sous-lot codé dans `tools/c1c` :

- extraction Go AST des imports, types, champs, méthodes, fonctions, valeurs ;
- signatures sans corps ;
- manifeste JSON ;
- tests `TestExtract` et `TestSignatureExcludesBody` ;
- `go test -race` passe.

Ce manifeste est l’entrée du générateur API/implémentation/bridge. Aucun skip
du planificateur n’est encore actif.


## Phase 3.2 — overlay generator (fait)

`tools/c1c/overlay.go` :

- extrait les symboles exportés de l’ancien bundle ;
- extrait les déclarations locales du nouveau module ;
- qualifie les références au bundle avec un alias `oldpkg` ;
- réécrit le `package` et ajoute l’import ;
- vérifie la syntaxe générée ;
- tests `TestOverlaySource` et `CompileCheck` passent.

Le générateur produit une base d’overlay compilable, mais n’est pas encore branché
au compilateur ni au cache de packages. Le skip et le patch des entrées restent
à faire.


## Phase 3.3 — réécriture des références (fait)

`tools/c1c/rewrite.go` :

- repère l’import de l’ancien bundle ;
- remplace `ancien.Symbole` par `nouveau.Symbole` pour les symboles du module ;
- ajoute l’import du package overlay ;
- conserve les autres symboles de l’ancien bundle ;
- test `TestRewritePackageReferences` et validation syntaxique.

Composants C1c disponibles : manifeste, overlay, réécriture. Le branchement
dans le compilateur, le skip et le calcul de plan incrémental restent à faire.


## Phase 3.4 — overlay end-to-end (fait)

Test `TestOverlayEndToEnd` :

- ancien bundle `old` ;
- nouveau module ;
- génération overlay ;
- réécriture de l’entrée `old.Get` → `overlay.Get` ;
- suppression de l’import ancien devenu inutilisé ;
- `go run` réel et sortie vérifiée.

Les composants C1c sont donc validés ensemble sur un module simple. Il reste
le branchement dans le pipeline b8x et le skip du planificateur.


## Phase 3.5 — overlay branché dans le pipeline Go

Implémentation de ce lot historique (alors opt-in ; défaut actuel en tête de page) :

- `GOPURS_GO_PACKAGE_OVERLAY=1` active le chemin overlay ;
- `preparePackageSession` lit le `request.json`/`result.json` précédent,
  vérifie les fichiers publiés, puis prépare un `package-result.json` privé
  avant le planificateur complet ;
- `writePackageFiles` détecte ce résultat déjà présent dans le stage et
  n'appelle plus `packageFiles` ;
- `request.json`, `result.json` et `overlay.json` sont publiés sous
  `output/.gopurs-cache/go-packages/plan-v1/` par rename atomique depuis le
  stage (jamais depuis `session.output`, qui est privé) ;
- les bundles inchangés sont réutilisés par lien privé (copie si le lien est
  impossible), seuls le bundle overlay et ses dépendants réécrits changent ;
- `plan-v1/overlay.json` conserve le module, la source, le paquet actif et les
  symboles redirigés pour l'édition suivante ;
- le cache de plan vérifie désormais les hashes des fichiers publiés
  (`miss-output`).

### Preuves d'éligibilité du premier chemin de relocalisation

Le delta est refusé si un seul des points suivants manque :

- signatures Go identiques (`eligible-body`) ;
- une seule source générée modifiée, jamais un fichier `ffi/` ;
- déclarations de la section identiques, sans type, méthode, `init` ni
  variable de paquet initialisée ;
- mêmes arêtes de graphe déclaration → déclaration (références par
  déclaration) ;
- mêmes paquets externes consommés par le module ;
- aucune référence privée inter-paquets ;
- aucun appel du module édité depuis un autre module du même bundle ;
- aucun import point, aucun paquet dépendant ambigu ;
- résultat validé par le même contrat que `packageOutputs`.

Sinon : `Go overlay fallback: <raison>` sur stderr et planificateur complet.

### Tests

- `TestPackageOverlayExecutesBodyEdit` : overlay + `go run` + oracle full ;
- `TestPackageOverlayRepeatedBodyEdit` : deuxième édition sur le même paquet ;
- `TestPackageOverlayRejectsNewDeclarationReference` : arête de graphe
  modifiée → full ;
- `TestPackageOverlayRejectsTypeDeclaration` : type → full ;
- `TestPlanCacheStatus` : `miss-output` sur fichier publié modifié ;
- `TestBuildOverlayQualifiesCrossPackageSymbols` et
  `TestRewriteUsesFreeAlias` côté `tools/c1c`.

### Limites du premier lot

- un module sans type déclaré ; les types restent dans le paquet d'origine ;
- un seul module modifié par build ;
- les éditions consécutives du même module sont supportées (overlay record) ;
  une édition d'un autre module retombe full ;
- FFI modifiée → full ; FFI inchangée référencée par le corps → overlay ;
- aucune comparaison implicite au full dans le chemin chaud. Dans ce lot, le
  flag était désactivé par défaut dans l'attente des autres campagnes.

### Mesures historiques build 6 — trois paires (Ping)

Machine Darwin arm64, `GOPURS_GO_PACKAGES=1`,
`GOPURS_GO_PACKAGE_OVERLAY=1`, mêmes TAST pour chaque paire. Le backend full
est un oracle one-shot (cache module chaud, planificateur complet).

| Paire | Backend overlay | Go exports* | Go packages | Backend full | Go packages full | Sortie overlay/full |
| --- | ---: | ---: | ---: | ---: | ---: | --- |
| `lot6` | **21,200 s** | 1,429 s | **0,005 s** | 52,476 s | 11,283 s | `pong-lot6` / identique |
| `lot7` | **20,375 s** | 1,315 s | **0,005 s** | 52,683 s | 11,535 s | `pong-lot7` / identique |
| `lot8` | **19,524 s** | 1,293 s | **0,008 s** | 51,295 s | 11,271 s | `pong-lot8` / identique |

\* `Go exports` inclut la construction de la requête, la réutilisation des
signatures et la préparation overlay ; la phase `Go packages` ne rend plus que
les entrées et le résultat privé.

Éditions consécutives sans oracle : `lot9` 21,067 s, `lot10` 19,805 s. Le
`go build` de l'entrée Ping tombe à **0,81–1,13 s** quand le paquet full a déjà
été compilé ; la chaîne complète reste donc sous 30 s. Le premier `go build`
après une republication full recompile le gros paquet et coûte ~20 s : il doit
être mesuré séparément.

Le `planHash` de l'overlay change à chaque corps, mais la preuve de graphe
(arêtes par déclaration) est stable. Aucune sortie Ping ne diffère du full sur
les trois paires.

### Mesures historiques build 7 — binaire mesuré

Le build 7 déplace la lecture des imports Go après la preuve overlay : un
corps seul ne parse plus les en-têtes des 2 704 sources. Trois éditions
consécutives, cache Go Ping déjà chaud :

| Édition | Backend overlay | Go exports | Go packages | `go build` Ping | Sortie |
| --- | ---: | ---: | ---: | ---: | --- |
| `lot11` (après full) | 21,515 s | 1,388 s | 0,021 s | 0,878 s | `pong-lot11` |
| `lot12` | **19,773 s** | 1,187 s | 0,008 s | 0,795 s | `pong-lot12` |
| `lot13` | **18,960 s** | 1,189 s | 0,004 s | 0,801 s | `pong-lot13` |

Oracle full sur le même TAST `lot13` : backend **38,539 s**, Go packages
**10,227 s**, sortie `pong-lot13` identique. Le premier overlay qui suit une
publication full paie le lookup C2 plus froid (21,5 s) ; les suivants passent
sous 20 s. Médiane des trois : **19,773 s**. Chaîne complète mesurée :
backend 19,0–21,5 s + `go build` 0,8–0,9 s + frontend ~2 s, soit **~22–24 s**.

### Mesures historiques build 8 — binaire de cette campagne

Le build 8 ajoute uniquement la validation des chemins du plan publié. Une
paire supplémentaire sur le binaire final donne :

| Édition | Backend overlay | Go exports | Go packages | Backend full | Go packages full | `go build` overlay/full | Sortie |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| `lot14` | **17,323 s** | 1,207 s | **0,003 s** | 38,376 s | 10,146 s | 2,054 / 20,198 s | `pong-lot14` / identique |

Cette paire historique passait sous 20 s, avec exécution Ping identique au
full. Elle ne constitue pas une validation du cache de préparation corrigé
lors de la reprise du 23 octobre.

### Reprise du 23 octobre — modules répartis sur plusieurs paquets

Le module `Util.Type.Variant.Encoding.Encoding.Test.DecodeJsonWith` est réparti
sur cinq bundles. Le chemin `retained-5` conserve les placements, les imports,
les caches privés et les fichiers d'ordre publiés. Il remplace seulement les
sections dont les octets changent ; les autres fichiers sont liés au stage.
Il ne crée pas de paquet overlay et n'écrit pas de record de relocalisation.
`retained-5` désigne les cinq paquets contenant des déclarations du module,
pas nécessairement cinq fichiers modifiés. Les entrées conservent leur paquet.
Le résultat reçoit un nouveau `planHash`, dérivé du plan précédent et du hash
de la source courante, sans recalculer le placement des déclarations.

Ce chemin est essayé avant la lecture des exports Go et le planificateur
complet lorsque `planCacheStatus` vaut `eligible-body`. Il exige exactement
une source Go générée modifiée, répartie dans plusieurs sections publiées,
avec des signatures brutes inchangées. Une édition PureScript peut produire
plusieurs sources Go différentes après optimisation ; ce cas n'est pas
couvert. Un ancien record de relocalisation d'un autre module entraîne aussi
un fallback. Le résultat préparé est publié par le même contrat que les plans
complets ; aucune sortie publique n'est modifiée pendant la préparation.

La preuve vérifie les hashes du plan publié, l'unicité des placements, les
signatures, les références par déclaration et les imports externes. Types,
méthodes, `init`, variables globales initialisées et graphe ambigu restent des
motifs de fallback. Les références comparées sont celles résolues par les
tables de symboles et d'imports des sections ; cette vérification ne remplace
pas un type-check Go complet de corps arbitraires.

Le planificateur renomme aussi les champs de records, par exemple `value` en
`P_value`. Pour une édition de littéral dans un corps, l'ancien Go brut est
matérialisé depuis le cache et son hash doit correspondre à la requête du plan
publié. Une comparaison de tous les tokens bruts exige que seuls des littéraux
de corps aient changé ; les commentaires, dont les annotations de records,
font partie de cette preuve. Les suites complètes de littéraux de chaque
fonction publiée doivent ensuite correspondre à l'ancien brut. On peut alors
modifier leurs seules positions, en conservant exactement les transformations
de types du planificateur. Identifiants, opérateurs, types, imports et
annotations changés ne satisfont pas cette preuve. Le cas central exercé est
le remplacement du texte d'un `it` dans `spec`.

Lorsque cette preuve de littéraux n'est pas disponible, une seconde tentative
repartit les déclarations de la source courante selon leurs placements
publiés. Elle qualifie les références avec les imports déjà présents, puis
compare signatures, références et imports avec chaque ancienne section.
Elle refuse les références privées entre paquets, les imports nouveaux,
les alias masqués, les groupes de déclarations répartis entre plusieurs
paquets et les clés de littéraux composites ambiguës. Elle ne reconstruit
pas le renommage des champs de records effectué par le planificateur : une
signature différente après réécriture entraîne donc le full.

Ces deux tentatives ne garantissent pas l'accélération d'une édition de corps
quelconque. Le chemin chaud ne lance pas implicitement un build full de
comparaison. À chaque refus, `Go overlay fallback: <raison>` est journalisé
et le planificateur complet reprend. Le lanceur Go natif active désormais ces
chemins par défaut ; les autres hôtes restent opt-in. La validation générale
des scénarios FFI et des quatre hôtes n'est pas acquise.

Validation locale : tests Go `-race` complets build-cache et C1c ; oracle
d'exécution d'un module réparti sur deux paquets contre son rendu monolithique,
puis deuxième édition ; rejets de types, init, variables initialisées, arêtes,
signatures, métadonnées et placements ambigus. La sonde sur les cinq bundles
réels a accepté l'édition centrale en **0,836 s**, puis **0,920 s** dans une copie
figée. Ces durées couvrent uniquement `preparePackageOverlay`, sans frontend,
backend PBO ni compilation Go ; elles ne valident pas la cible backend globale.

Une première paire native centrale a également comparé les **2 704 modules**
(sources Go brutes, y compris FFI) et les **27 fichiers Go publiés** (paquets,
ordre, initialisation, entrées et runtime) au full forcé : aucun écart de
bytes. Le script vérifie les objets compressés, leurs hashes bruts et les
fichiers réellement publiés. Cette preuve figure dans
`central-probe-pair/results.json` ; elle porte sur l'édition de libellé de
cette paire, sans prétendre valider les performances des trois paires finales.

La campagne finale `central-validated-pairs/results.json` a ensuite validé les
trois paires complètes : **16,970 / 17,293 / 16,946 s** de backend one-shot,
avec les 2 704 modules et 27 sorties byte-exacts à chaque fois. Compilation Go
finale réussie ; frontend et Go exclus de ces durées. Voir le protocole et les
limites dans [C2-7](build-c2-persistent.md).

Le helper JS inclut `overlay_sections.go`, `partition.go` et `literals.go` dans
sa clé de contenu. Un test en copie modifie chacun de ces fichiers et vérifie
qu'un nouveau helper est capturé. Les snapshots JS/natif, le pont de cache
JS↔Go et les suites Go `-race` sont verts ; les listes de sources embarquées
par les hôtes natif et Rust incluent aussi ces fichiers. Cela vérifie leur
présence et les chemins testés, sans constituer une campagne Rust complète.

La copie figée, son test Go via `-overlay` et le journal sont conservés sous
`htdocs/.gopurs-evidence/c2-7-reprise-20261023/overlay-partition-proof/`.
Le script `measure-central-pairs.py` adjacent prépare trois paires natives avec
vérification des sources Go brutes et de tous les fichiers Go publiés contre
le full. Ses mesures de campagne ne sont pas incluses dans cette sonde.
