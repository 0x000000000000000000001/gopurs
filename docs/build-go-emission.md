# Émission de packages Go — G03 validé

**G03 est validé le 14 octobre 2026**, suite G à **4/7**, P à **0/7**. Le
[plan G01](build-go-package-plan.md) est porté dans le résolveur Go. Le pilote
complet acquiert les exports et publie les fragments multipackages via le staging
et le journal existants. Les quatre hôtes passent les contrôles CLI et les
fixtures. b8x compile et exécute son oracle ; le bootstrap multipackage fonctionne
comme compilateur. La parité des deux grands corpus est exacte sur les quatre
hôtes. La restauration incrémentale et les contrats complets de cache relèvent de
G05 ; l'acceptation temporelle et l'activation finale du layout relèvent de G06.

Le [layout G06 installé](build-go-measurements.md) conserve cet algorithme, avec
des ancres **1/8192** retenues après mesures. Il regroupe les fragments ordinaires
en un fichier par package et abrège les alias d'import. Les références chiffrées
et preuves ci-dessous restent celles de G03 à **1/128**.

## Émetteur typé

`tools/ffi-gen/bindings.go` partage son contexte `go/types` avec `packages.go` :
sources possédées, objets d'origine, déclarations et `Info.InitOrder`.
`PackageSources` applique les unions getters/caches/workers, FFI entières,
méthodes/receveurs et références privées. Les membres gardent leurs arêtes
distinctes. Les contraintes d'initialisation précèdent la condensation.

Les SCC sont calculées sans récursion. La file topologique est ordonnée par
l'empreinte des identités, puis les ancres **1/128** ouvrent les intervalles.
La densité 1 est utilisée uniquement par les tests de frontières. Le résultat
`gopurs/go-packages-v1` contient le plan, son empreinte, les entrées et les
sources émises ; il ne sérialise pas chaque référence du graphe typé.

Le rendu qualifie les références libres avec les alias du plan, conserve les
champs/méthodes, retire les imports devenus inutiles et ajoute les dépendances
d'ordre. `init_external` retient les effets des imports externes ; `init_all`
retient la fin de la chaîne d'initialisation, même pour une entrée partielle.
`Driver.Output.writePackageOutput` écrit les fragments et les deux formes
d'entrée existantes, après résolution des symboles d'entrée.

Le pilote CLI sélectionne ce layout avec `GOPURS_GO_PACKAGES=1`. L'opération
`packages` du cache prépare une requête compacte depuis les fichiers du staging.
`writePackageFiles` émet les entrées résolues ; la publication valide l'inventaire,
les empreintes et les dépendances topologiques avant d'utiliser le journal de
propriété. Les fichiers externes et l'inode du verrou sont préservés.

Le cache observe le `go.mod` existant, `go.sum`, les dépendances natives locales
et les remplacements locaux. Il fige ces sources avant d'acquérir les exports
avec `go list`. Les imports sont ceux des fichiers sélectionnés par Go, sans
ajouter les packages inutilisés ni les fichiers de test. La version et
l'environnement de la chaîne Go participent aux clés. `go.mod` est créé seulement
s'il manque ; `go.sum` reste externe.

L'acquisition G06 d'exports allégés utilise, lorsque `GOFLAGS` est vide,
`-p=8 -trimpath -gcflags=all=-N -l -linkobj=<os.DevNull>`. Le cache Go conserve alors les
informations de types sans les objets de lien ni les corps destinés à l'inlining.
Ces options participent aux clés d'action de Go ; une compilation normale
utilisant le même GOCACHE doit reconstruire ses propres objets exécutables. Les
options explicites du demandeur suivent le parcours d'acquisition à `-p=1`.
Le [bilan G06](build-go-measurements.md) vérifie l'émission exacte à partir des
250 exports b8x, puis les quatre hôtes et les 36 chaînes du candidat complet.

`tools/ffi-gen/workspace.go` fournit le transport natif
`gopurs/go-package-files-v1` : chemins et SHA-256 en entrée, inventaire/plan/entrées
en sortie. Il vérifie les octets consommés puis écrit dans une nouvelle
destination privée. JS invoque un helper figé dans le snapshot du compilateur ;
Go appelle le même code directement et Rust emploie l'opération C **4**. Les
opérations 0/1/2/3 gardent leurs rôles, dont le transport pur `Packages` comme
oracle. Le résolveur pur reçoit toujours les exports sous forme d'octets.

## Records et identités d'initialisation

Une annotation `/*gopurs:record-v1:<nombre>:<labels UTF-16 hex>*/` accompagne
chaque `TypeRecord`, y compris lorsqu'il passe par un fragment opaque.
`//gopurs:init-v1:<propriétaire hex>:<identité hex>` conserve l'identité fournie
par `GoInitDecl`. Les annotations sont des commentaires Go, sans changement de
type monolithique. Le rendu multipackage les consomme et les retire.

Les champs explicitement attribués à un record deviennent **`P_` + label ASCII
identifiant**, sinon **`PX_` + unités UTF-16 hexadécimales de largeur quatre**.
Les unités isolées sont préservées. Le nommage intermédiaire historique a été
rendu injectif pour les collisions `type`/`go__type`, `a-b`/`a_minus_b`, les
labels vides et les préfixes numériques ; les noms ASCII ordinaires restent
identiques.

Une égalité de types Go n'attribue pas à elle seule la propriété d'un record :
`struct{V0 int64}` peut représenter un tuple, un record ou un type natif. Le
candidat qui adaptait toutes les formes égales est rejeté. Chaque champ renommé
est lié à son objet Go d'origine et à une annotation explicite. Les FFI qui
énoncent un layout de record doivent fournir cette information elles aussi,
par annotation ou `SourceInput.records` ; leur revue appartient à G04.

Le runtime fournit `RecordFieldName` et `RecordFieldLabel` pour les accès
réfléchis explicites aux records PureScript. La correspondance est bijective,
conserve ordre/layout et WTF-8, et ne s'applique pas implicitement aux champs
natifs, aux tuples ou à `Rc`/`V0` des ADT.

## Transports et observations

`GoSource.packageSources` expose la même opération via WASM, Go direct et
l'opération **3** du C ABI Rust. Les opérations 0/1/2 gardent leurs rôles.
Le transport pur reçoit les mêmes sources et exports dans les trois cas.

Les contrôles ciblés éprouvent :

- génériques, assertions, champs embarqués, interfaces et variables homonymes ;
- cycles sans racine et **95 comparaisons SCC/fermeture transitive** ;
- stabilité du plan pour un changement de corps et permutation des sources ;
- effets externes, ordre variables/`init`, entrée partielle et **64 appels
  concurrents** partageant une cellule lazy ;
- échec du record privé après séparation naïve, puis exécution équivalente avec
  labels exportés ; conservation d'un tuple et d'un type natif de forme égale ;
- production des annotations par le véritable AST/printer, écriture par
  `Driver.Output`, compilation et exécution des deux entrées.

Les premières réponses multipackages sont byte-exactes entre les trois
transports. Les bootstraps Go et Rust sans PGO sont reconstruits dans des
répertoires privés. La campagne de fixtures réelles compare ensuite JS, Go
séquentiel, Go parallèle et Rust sur les mêmes TAST.

La première fixture réelle a exposé un trap WASM dans `runtime.mapaccess2`
lorsque le rendu recherchait un champ à partir d'un identifiant sans objet lié.
La variante qui ignore ces recherches redonne les **1 267 508 octets** natifs
exacts. La table de renommage est maintenant typée par `*types.Var` et consultée
uniquement pour des champs. Le cas complet est conservé ; les petits essais
isolés à 8/32/128/512 champs ne reproduisent pas ce trap.

Les exports de cette fixture occupaient **39 061 418 octets** sous forme
d'archives complètes ; leurs membres `__.PKGDEF` occupent **1 537 762 octets**.
Leur extraction conserve exactement les sorties natives. Cette réduction de
transport n'était pas la correction du trap et n'est pas un gain de compilation
mesuré. `StringEdgeCases` a ensuite dépassé le tampon stdout JS de **10 MiB**
(`ENOBUFS`). Les réponses de programme complet sont maintenant écrites dans un
fichier temporaire privé, lu puis supprimé, ce qui supprime cette limite de pipe.
Le CLI emploie le transport natif compact pour les gros corpus.

La comparaison suivante a révélé des annotations UTF-16 différentes entre JS et
Go : le passage par `fromEnum Char` remplaçait les surrogates par `fffd`. Le
producteur convertit désormais explicitement les points de code en unités
UTF-16. `real-fixtures-5` termine **12/12 passages / 129 commandes** :
`NativeRecordWorkers`, `FFIIntegerReturns` et `StringEdgeCases` sur les quatre
hôtes. Sources monolithiques et multipackages, plans et trois transports sont
exacts ; compilation et exécution réussissent. `StringEdgeCases` couvre
**271 TAST / 33 packages**. Une sélection erronée de l'ancien binaire Rust dans
la tentative 4 est conservée ; le harnais vérifie maintenant son empreinte.

`driver-packages-2` valide le CLI complet sur les quatre hôtes : migration
monolithique → multipackage → monolithique, dépendance native et remplacement
local, modification de signature, deux entrées, build forcé exact, fichiers
externes et inode du verrou préservés. L'identique ne réécrit pas les sorties et
ne compile ni ne lie. Ces contrôles fonctionnels ne constituent pas des paires
de performance. La restauration des objets de modules après une édition et les
campagnes de reprise étendues restent à compléter en G05.

## Grands corpus et rapprochement du plan — 14 octobre 2026

`candidate-5` contient la correction d'identité des helpers : `goTypeIdentity`
et `structTypeIdentity` parcourent les types imbriqués sans les annotations de
transport. Les annotations ne participent donc plus aux noms `Rebox`. Sur
`NativeRecordWorkers`, retirer les seules annotations puis appliquer `gofmt`
redonne le snapshot historique à l'octet près. Les quatre hôtes réussissent
`real-fixtures-6`. Les snapshots examinés sont ensuite mis à jour : **27 lignes
d'annotations** dans `FFIIntegerReturns`, **134** dans `NativeRecordWorkers` ;
le snapshot FFI et `StringEdgeCases` restent identiques. Le rejeu strict des
trois fixtures sur les quatre hôtes termine ses **67 commandes**, sans modifier
les snapshots pendant le rejeu.

Les sources des compagnons sont éprouvées sous **Go 1.22.12**. Le module WASM
reste épinglé à Go 1.27.0 ; la vérification du parseur utilise donc une copie
privée des mêmes sources avec un `go.mod` 1.22. La première tentative dans le
module épinglé est refusée, puis une fixture demandant inutilement Go 1.26 est
corrigée en 1.22. Le parseur et le cache passent ensuite tous leurs tests.

Le premier b8x avait terminé l'acquisition des exports et l'émission, mais la
publication refusait les anciens Go sans propriété vérifiée. Le harnais conserve
ces **3 010 fichiers** dans `legacy-generated/`, après vérification contre
`full-1-inventory.json`, dans une nouvelle copie du corpus. Aucun journal n'est
transplanté. `large-2-b8x-native-parallel` publie **2 704 TAST / 39 613 Go /
525 packages**. Son oracle compilé donne exactement :

```text
camel=""
accent="e"
pong
```

La reconstruction identique ne lance ni compilateur Go ni éditeur de liens.
Ces passages fonctionnels ne sont pas des mesures d'acceptation G06.

Le bootstrap `candidate-4` compile en **157 packages**, puis son binaire passe
le test CLI. Après correction des identités, `candidate-5` produit **512 TAST /
8 389 Go / 150 packages** ; ce second binaire passe également le test CLI complet
(migration, dépendance locale, sélection d'entrée, forcé et identique). La
première tentative de bootstrap, dont les chemins de projet avaient une
profondeur incorrecte et ne trouvaient pas les FFI, reste conservée.

Le véritable `planPackages` reproduit les **15 plans G01** figés (b8x base et
trois éditions, bootstrap ; densités 128/256/512). Le rapprochement des sorties
effectivement émises précise les écarts de corpus :

- b8x conserve les placements de **185 591 déclarations de même nom**. Les
  **76 helpers Rebox renommés** utilisent des labels non ASCII-identifiants,
  notamment `rlist'` et `seo.updatedAt`, dont le nommage a été rendu injectif.
  Leurs signatures, normalisées uniquement pour cette comparaison, établissent
  une bijection. **51** changent de package. Le modèle G01 avec ces identités
  actualisées retrouve les **185 667 placements / 525 packages / 35 596 arêtes**
  exacts ; il y avait 35 594 arêtes avant ces renommages.
- le bootstrap courant contient **65 964 déclarations**, contre 65 145 au relevé
  G00. Son graphe typé est recalculé depuis le monolithe courant. Le relevé G00
  bootstrap attribuait les propriétaires depuis les stems Go ; la comparaison
  courante les relie aux vrais noms PureScript du même inventaire TAST. Le modèle
  retrouve alors exactement les **150 packages / 5 875 arêtes** émis.

Les dépendances comparées sont les imports réels de `go list`. Les modèles
retiennent respectivement **2 637 / 1 103 opérations d'initialisation**. Ces
égalités sont des preuves de partition et de correction, pas des gains temporels.

`large-parity-1` termine **huit comparaisons / 28 commandes** : les deux corpus
sur JS, Go séquentiel, Go parallèle et Rust. Les builds forcés reproduisent les
inventaires, les octets et les dates des fichiers ; les relances identiques sont
des hits, sans compilation ni lien Go. Les TAST, `go.mod`/`go.sum` et les inodes
des verrous restent identiques dans chaque racine. Les quatre objets de chaque
module, ses empreintes préparée/demandes et les inventaires possédés sont exacts
entre hôtes. Les clés d'environnement diffèrent selon l'identité du compilateur,
comme prévu ; la première assertion qui comparait les mémos entiers est conservée.

L'audit des déclarations b8x compare ensuite les tokens Go après les seules
normalisations de qualification, de champs record explicitement exportés et des
76 noms de helpers examinés : **185 667 déclarations équivalentes, zéro écart de
corps**. Les **348 tags d'enregistrement** restent uniques, avec les mêmes
propriétaires. `final-reference/summary.json` identifie les quatre artefacts et
atteste les sources de compilation inchangées depuis `candidate-5` ; seules les
attentes de tests et les adaptations du runner préparées pour G04 ont évolué.

## Preuves et reproducteurs

Racine : `gopurs-go-packages-g03-20261013/` dans le temporaire OpenCode.

- `before/` : **1 258 fichiers inventoriés**, sources G02 et état du travail ;
- `protocol.md` : objectifs et budgets avant les changements ;
- `packages-*.log`, `packages-*/` : programmes Go, graphe et oracles ;
- `record-shape-before.go` et `.log` : candidat refusé sur l'ABI tuple/natif ;
- `record-labels-before.log` et `go-packages-*/before/` : collisions de labels ;
- `record-transports-2.log` : **15 contrôles réussis**, writer et transports ;
- `native-1/`, `candidate-1/`, `gopurs-rust-build-9O2LOw/` : premiers bootstraps ;
- `real-fixtures-1/` : première requête réelle, exports compacts, réponse native,
  trap WASM et comparaison après garde ;
- `wasm-debug-1/` : instrumentation privée, binaire avant/après et traces ;
- `transport-refresh-2/` : reconstruction du bridge corrigé, avec ancien archive
  et binaires conservés ;
- `real-fixtures-2/3/4/` : reproducteurs `ENOBUFS`, divergence UTF-16 et sélection
  du binaire Rust ; `real-fixtures-5/summary.json` : campagne complète exacte ;
- `candidate-3/`, `candidate-4/`, `native-3/4/`, `rust-build-3/4.log` : candidats
  et bootstraps monolithiques figés, Rust sans PGO ;
- `workspace-prototype-1/`, `workspace-files-1.log`, `file-transports-1.log` :
  oracle du transport par fichiers, mutation des entrées et parité JS/Go/C-Rust ;
- `exports-scope-before/`, `exports-scope-before.log`, `exports-scope-after.log` :
  rejet des imports inutilisés/test puis acquisition du graphe Go sélectionné ;
- `driver-packages-1/2.log`, `driver-packages-AvLsBq/` : contrôles CLI quatre hôtes.
- `candidate-5/`, `native-5/`, `rust-build-5.log` : correction Rebox figée,
  bootstrap Rust sans PGO et inventaires des artefacts ;
- `type-identity-after.log`, `real-fixtures-6/`, `snapshot-review-2/`,
  `snapshots-update-1.log`, `strict-replays-1/` : revue puis replays stricts ;
- `bootstrap-order-before.log` / `bootstrap-order-after.log` : attentes du script
  de bootstrap adaptées à la préparation des packages natifs avant l'émission ;
- `parser-go122-1/2/3.log`, `parser-go122-companion-1/`, `cache-go122-1.log` :
  tentatives et validation de compatibilité des compagnons ;
- `large-1-b8x-native-sequential/`, `large-2-b8x-native-parallel/` et son
  `oracle-1/` : publication refusée, reprise avec provenance, compilation/oracle ;
- `large-1/2-bootstrap-native-sequential/`, `large-3-bootstrap-native-parallel/`,
  `bootstrap-execute-1/2.log` : échec initial et deux bootstraps exécutés ;
- `plan-port-1/2/`, `emitted-b8x-1/2/3.json`, `emitted-bootstrap-1.json`,
  `bootstrap-plan-current-1/`, `emitted-plan-reconciliation-1/` : plans figés,
  écarts initiaux, signatures des helpers et rapprochement exact des graphes.
- `large-parity-1/`, `emitted-b8x-tokens-1.json`, `abi-b8x-1.json`,
  `final-reference/` : parité finale, audit des corps/tags et clôture G03.

Les échecs de harnais sont aussi conservés : sortie `go build` heurtant le
répertoire `rejected`, confusion des positions de braces avec celles du
`FieldList` Go, et test du runtime lancé sans `GO111MODULE=off`.

Le nettoyage du cache Go global de l'hôte a été autorisé par l'utilisateur.
Il a libéré environ **30 GiB**, portant l'espace libre à **38 GiB**. Les preuves
et GOCACHE privés de campagnes sont conservés ; les temps de validation après
ce nettoyage ne sont pas des paires de performance.

## Reproduire les contrôles

```sh
npm run build
npm run build:ffi
node --test --test-concurrency=1 tools/go-source.test.mjs \
  tools/go-packages.test.mjs tools/record-tuple-conversions.test.mjs
GOPURS_TEST_MODES=js,native-sequential,native-parallel,rust \
  node --test tools/driver-packages.test.mjs
```

Dans `tools/ffi-gen` : `go test -race -count=1 ./...`.
Dans `runtime` : `GO111MODULE=off go test -race -count=1 -run TestRecordField`.
`GOPURS_TEST_EVIDENCE=<répertoire existant>` conserve les programmes et sorties.
