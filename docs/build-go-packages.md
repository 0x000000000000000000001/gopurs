# Granularité des packages Go — référence G00

**G00 validé le 13 octobre 2026**, après les relevés commencés le 12. Le plan v3
est clos à **7/7** ; G00, le [plan G01](build-go-package-plan.md) et les
[références G02](build-go-references.md) portent la suite G à **3/7 lots validés**. Le
[bilan v3/06](build-incremental.md#parcours-complet--lot-06) donne la référence
du parcours complet. Les mesures ci-dessous portent sur les sorties du
compilateur livré, `publicsync`.

## Pourquoi descendre sous le module

`Driver.Output` écrit les modules et leurs FFI dans `output/purescript/`.
Chaque entrée importe ce package et appelle `purescript.Get_…_main`. Le
package commun recompilé représente presque tout le coût Go après édition :

| Édition | `go build` | Action de compilation `purescript` | Chaîne jusqu'à `exec` | Pic RSS du compilateur Go |
| --- | ---: | ---: | ---: | ---: |
| Feuille | 206,820 s | 202,788 s | 244,403 s | 14,673 GiB |
| Partagée | 209,650 s | 205,196 s | 268,585 s | 14,697 GiB |
| FFI | 208,470 s | 204,576 s | 236,510 s | 14,732 GiB |

Ce sont les médianes des trois passages `publicsync` de v3/06. Les
[24 mesures](build-cycle-lot06.csv) conservent aussi la référence lot 03 et
la dispersion. Les ressources du runner restent **2 / -p 1 / 12GiB**.

L'analyseur privé parse le Go final, puis vérifie **tous les corps** avec
`go/types`, Go **1.26.8 linux/arm64**, en utilisant les exports réels des
dépendances. Les fichiers sont rattachés aux propriétaires du manifeste.
Cette analyse voit aussi les références contenues dans les FFI et les
fragments opaques déjà rendus. Elle ne se limite pas aux imports PureScript.

Sur **2 704 modules**, le package contient **2 997 fichiers**, soit
**133 747 559 octets**. Le graphe agrégé comporte **26 143 arêtes**, **951
composantes fortement connexes**, dont une seule non triviale : **1 754
modules / 126 426 172 octets, soit 94,53 %** du code du package.
Sans les arêtes de champs, ce bloc contient encore **94,36 %** du code ; en
ne gardant que les appels de fonctions, **92,72 %**. Les spécialisations de
bibliothèques peuvent appeler des getters applicatifs : déplacer simplement
chaque module dans son propre package ne résout donc pas le problème.

## Graphe par déclaration

Le deuxième relevé descend aux déclarations Go : fonctions, types, groupes
`var`/`const`, et `init`. Les groupes restent indivisibles pour conserver les
initialisations multiples et les dépendances `iota`. Les **26 méthodes** sont
liées à leur type receveur. Les champs sont identifiés séparément : une sélection
de champ exprime une contrainte de layout/visibilité, et non un identifiant à
préfixer par un nom de package.

La base compte **185 667 déclarations**, **385 401 arêtes distinctes** avant
l'ajout des contraintes de méthode, et **133 173 560 octets de déclarations**.
Avec les arêtes de champs conservées, elle donne **185 193 composantes**, dont
**214 non triviales**. Le DAG condensé est vérifié par un second algorithme.
Le plus gros bloc, **6 967 987 octets / 5,23 %**, est une seule fonction :
`Get_Infra_Projection_Postgres_Finder_Test_Integration_HandleProjectionReadFind_spec`.

Les trois variantes passent également le contrôle de types. Les identités de
déclaration sont rapprochées avant/après ; les octets des corps sont comparés.
Les fermetures inverses incluent les appels, les références de types et les
contraintes de champs recensées :

| Édition | Déclarations modifiées / ajoutées / retirées | Déclarations dans la fermeture après édition | Octets | Part des octets de déclarations |
| --- | ---: | ---: | ---: | ---: |
| Feuille | 1 / 0 / 0 | 1 | 800 | 0,0006 % |
| Partagée | 4 / 5 / 4 | 966 | 8 147 525 | 6,12 % |
| FFI | 1 / 0 / 0 | 457 | 15 291 979 | 11,48 % |

Ces tailles montrent un potentiel de partition. **Ce ne sont ni des temps Go
prédits, ni les régions finales de recompilation.** Le regroupement en packages,
les layouts partagés et l'initialisation ajouteront des contraintes. Go peut
aussi conserver un dépendant lorsque les exports de son fournisseur restent
identiques. À l'inverse, un package regroupé peut inclure des déclarations hors
de cette fermeture. Émettre 185 000 packages n'est pas le plan proposé.

## Contraintes de bridges relevées

### Identité des types et visibilité

Le rendu actuel produit des records anonymes à champs souvent privés :

```go
func Make() struct{ value int64 }
func Consume(x struct{ value int64 })
```

Dans deux packages distincts, ces signatures n'utilisent plus le même type,
car l'identité d'un champ privé inclut son package. Exporter un champ exige
aussi de préserver la correspondance avec son label PureScript, sans collision.
Un type nommé ne peut pas être dupliqué pour résoudre ce problème : sa
déclaration propriétaire fait partie de son identité. Une méthode doit être
déclarée dans le package de son type receveur.

`GoAst.StructPointer` distingue déjà nom Go, nom PureScript et identité de tag
runtime. La qualification doit préserver cette séparation. Les helpers Rebox,
arguments de types et signatures de workers doivent référencer la même
déclaration de type ; les casts `unsafe.Pointer` des FFI dépendent du layout.

### Initialisation et accès dynamiques

- Le corpus comporte **349 fonctions `init`**, dont **348 enregistrements
  `StructGetters`**, avec **348 tags distincts**. `ModuleDeclarations` produit
  ces fonctions et leurs fragments Go ; `runtime.RecordGet` consulte le registre.
  Une entrée doit conserver les enregistrements requis même si elle ne référence
  pas directement leur module d'origine.
- L'autre `init`, dans `Infra.EventStore.Postgres.EventStore`, initialise un
  identifiant de processus et le générateur aléatoire. Les dépendances de
  valeurs ne suffisent pas à fixer son placement ou son ordre d'exécution.
- **2 288 initialisations de variables** sont relevées dans l'ordre établi par
  `go/types`. Les getters générés emploient `cache_…` et `once_…` ; leur stockage
  doit rester unique, avec la même évaluation paresseuse.
- Cinq modules utilisent directement la réflexion : `Foreign`,
  `Promise.Rejection`, `Util.Cookie.Cookie`, `Util.Json.TaggedSum`,
  `Util.Run.Router`. Les accès `FieldByName("maxAge")`, `FieldByName("type")`,
  `FieldByName("value")` et `MethodByName` nécessitent une revue explicite.
  Un tag de struct ne modifie pas le comportement de `FieldByName`.
- Le runtime a aussi des chemins de conversion réflexive par **indice de
  champ** (`CoerceToStruct`, `ReboxStruct`, `copyReflectField`) et des accès
  `unsafe`. L'ordre et la représentation des champs sont des contraintes d'ABI.

### FFI applicatives et bootstrap

Le contrôle typé de b8x recense **187 usages FFI intermodules**, répartis sur
**17 arêtes** : notamment les promesses, les streams, HTTP et les event emitters.
Les références sont résolues dans le Go final, y compris les types de signatures.

Le même analyseur vérifie le bootstrap Go conservé de `publicsync` : **625
fichiers, dont 115 FFI**, **65 145 déclarations**, **404 usages FFI entre unités
générées / 28 arêtes**. Les corps sont vérifiés avec les exports réels du
runtime, du parser et du cache natifs ; le package compilateur n'a pas besoin
d'être recompilé pour cet inventaire. Les sources égalent l'audit du compilateur
livré. Les accès directs aux constructeurs concernent notamment
`CoreFn.Json`, `CoreFn.Json.Text`, `CoreFn.Usage` et `NativeMaps` dans PBO ;
les parseurs Argonaut partagent aussi des helpers, types et champs privés.

Le contrat de `ffi-gen/parser.go` actuel décrit les signatures et renomme les
déclarations FFI en respectant leur portée ; il ne transporte pas encore la
propriété interpackage de toutes les références. G02 doit couvrir ces données
dans le même parser pour WASM/JS, appel Go direct et archive C liée au bootstrap
Rust. `prepare-native-output.mjs`, `build-native.mjs` et `build-rust.mjs`
constituent les points de raccordement. Les tests multi-hôtes complets restent
ceux de G04/G06, après la transformation effective.

## Budgets fixés avant les candidats

- Gain G : **≥30 % sur `go build` et ≥20 % jusqu'à `exec`**, pour chacune des
  éditions feuille/partagée/FFI, sur au moins trois paires alternées.
- Régression : **≤5 %** ailleurs, froid compris ; identique :
  **max(5 % de la référence, 0,100 s)**.
- Mémoire : **+5 % maximum**, séparément pour le backend et le compilateur Go,
  par scénario et avec la même méthode de mesure.
- Stockage : **+10 % maximum par catégorie** — Go généré, objets gopurs,
  GOCACHE, exécutables conservés — à état/scénario égal. Le volume logique et
  le nombre de fichiers sont relevés séparément ; les copies de preuves sont
  exclues du stockage produit.

Le froid Go de préparation v3/06 fournit déjà une observation complète :
**221,830 s**, **242 compilations / un lien**, **15 421 356 KiB** de pic RSS Go.
Ce passage unique n'est pas une médiane appariée.

Le froid backend supplémentaire est byte-exact : **113,687 s**, **10 566 516 736
octets de RSS maximal (9,841 GiB)**, relevé par `/usr/bin/time -l` sur macOS.
Les **3 010 Go** totalisent **133 809 190 octets** ; le cache neuf conserve
**10 992 objets / 548 614 004 octets**. Ce relevé unique calibre la mémoire ;
il ne remplace pas les temps appariés de v3/04–06.

La calibration mémoire complète comporte **sept passages exacts**, dont les
deux retours à la base. Observations uniques du backend livré :

| Scénario | RSS maximal en octets | RSS en GiB | Temps observé |
| --- | ---: | ---: | ---: |
| Froid | 10 566 516 736 | 9,841 | 113,687 s |
| Identique | 1 419 460 608 | 1,322 | 3,299 s |
| Feuille | 10 897 637 376 | 10,149 | 38,092 s |
| Partagée | 11 136 319 488 | 10,372 | 40,017 s |
| FFI | 10 966 827 008 | 10,214 | 28,452 s |

Les temps de cette calibration sont conservés, y compris l'identique plus
lent ; ils ne constituent pas une nouvelle comparaison appariée de performance.
Les futures décisions G remesureront référence et candidat avec la même méthode.

La première tentative mémoire backend est rejetée : le harnais avait omis les
bibliothèques sœurs `workspace/gopurs`, référencées par les chemins TAST. Les
stubs FFI produits ont fait échouer l'inventaire byte-exact. La tentative, son
cache, ses journaux et ses ressources sont conservés, temps/RSS exclus. Le
harnais corrigé vérifie les **2 704 chemins source/FFI** avant le build.

### Stockage de référence

L'audit en lecture seule des douze échantillons `publicsync` de v3/06 relève
les médianes suivantes, en **octets logiques**. Le GOCACHE et les exécutables
sont inventoriés **après l'oracle de validation** : ils incluent ce dernier et
la base conservée. Cette portée doit rester identique lors de la comparaison G.

| Scénario | Go généré | Objets gopurs | GOCACHE après oracle | Exécutables conservés après oracle |
| --- | ---: | ---: | ---: | ---: |
| Identique | 133 809 190 | 548 617 831 | 1 136 866 936 | 95 678 320 |
| Feuille | 133 809 196 | 552 623 923 | 2 133 465 620 | 171 457 948 |
| Partagée | 133 811 204 | 562 156 847 | 2 133 476 274 | 171 467 154 |
| FFI | 133 809 196 | 566 140 963 | 2 133 465 636 | 171 454 876 |

Les **3 010 Go** sont présents dans chaque état. Le GOCACHE contient **1 676
fichiers** sur l'identique et **1 686** après édition ; les exécutables conservés
sont respectivement **deux** et **trois**. Le cache de base original est revérifié
par hash, ainsi que les binaires `Test.Main` de chaque échantillon. Ces chiffres
ne mesurent pas l'espace physique des clones APFS utilisés pour les preuves.

### Pause disque et reprise du 13 octobre

Après le froid valide, le harnais s'arrête **avant l'échantillon suivant** :
**7,9 GiB libres**, sous le plancher fixé à **8 GiB**. Après libération d'espace
par l'utilisateur, **19,1 GiB** sont disponibles et les six passages restants
réussissent. La reprise conserve le résultat froid, ses preuves et la racine du
journal ; elle revalide les inventaires, les ressources et l'inode du verrou.
La commande chronométrée est AST-identique à celle du harnais figé. Les sept
résultats conservent leur provenance ; l'inode du verrou reste inchangé.

Les **sept reproducteurs Go** réussissent, dont **quatre erreurs de compilation
attendues** : identité des champs privés entre packages, accès privé, duplication
de types nommés et méthode sur type non local. Les trois cas légaux sont le
package commun, les champs structurels exportés et l'alias vers un type nommé
partagé. Le contrôle de réflexion confirme qu'un tag n'adapte pas `FieldByName`.

### Froid applicatif Go

La mesure complémentaire préchauffe **uniquement les dépendances** (runtime
compris), avec application et entrée absentes du GOCACHE. Les traces confirment
ensuite **deux compilations (`purescript`, entrée) et un lien**. Le passage prend
**255,780 s de `go build`**, avec **15 433 168 KiB** de pic RSS du compilateur.
Le binaire est byte-exact face à v3/06. La relance identique prend **0,440 s de
validation Go**, sans compilation/lien ; l'oracle `camel / accent / Ping` passe.

Ce froid applicatif est plus lent que le froid intégral antérieur (**221,830 s**).
Ces observations uniques, non appariées, ne démontrent pas un effet du
préchauffage ; l'écart reste publié et les futures comparaisons G reprendront
des paires alternées avec leur dispersion. Après le build, avant l'oracle,
le GOCACHE contient **1 774 fichiers / 1 136 733 038 octets** ; l'unique binaire
publié occupe **75 779 628 octets**.

L'audit final réunit les cinq graphes typés, les sept passages mémoire avec
provenance, les contrats Go, les deux états du runner et l'oracle, ainsi que
les **430 fichiers de production** toujours exacts. **G00 est clos à 1/7**.

## Preuves et suite

Racine locale conservée :

```text
<tmp>/
  gopurs-go-packages-g00-20261012/
```

Principaux artefacts :

- `protocol.md`, `graph.go`, `run-reference.py`, `run-variants.py` ;
- `base-typed-graph-1/` : premier inventaire par module ;
- `base-declaration-graph-1/`, `variant-declaration-graphs-1/` et
  `declaration-analysis-1/` : graphes, inventaires, fermetures et comparaisons ;
- `bootstrap-typed-graph-1/` : inventaire typé du bootstrap livré ;
- `reference-audit-1/` : audit des **430 fichiers de production**, bridges,
  registres, dépendances Go de base et stockage des douze échantillons ;
- `record-contracts-1/` : sept reproducteurs d'identité/visibilité des types Go,
  plus le contrôle de réflexion, exécutés avec le Go épinglé ;
- `backend-memory-reference-1/rejection.json` et
  `backend-memory-reference-2/pause.json` : tentative rejetée et froid valide.
- `backend-memory-resume1/` : sept ressources/inventaires exacts avec provenance ;
- `go-application-cold-1/` : froid applicatif, identique et oracle réussis ;
- `final-reference/summary.json` : audit final et clôture de G00.

Les quatre hashes des graphes b8x et ceux du bilan v3/06 ont été revérifiés à la
reprise après coupure réseau. Les analyses emploient les outils et dépendances
locaux figés. Le [plan G01](build-go-package-plan.md) définit les identités et le
regroupement canonique, ses contraintes de visibilité/initialisation et la
portée statique des éditions après regroupement.
