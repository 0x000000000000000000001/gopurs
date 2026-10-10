# Contrat du cache de build — version 1

Le lot 02 fournit le stockage exécutable et son protocole commun aux hôtes JS et
Go. Le raccordement au pilote appartient aux lots 03 (build identique) et 04
(réutilisation par module). À ce stade, le pilote exécute encore un build complet ;
aucun gain sur les quelque 33 secondes de gopurs n'est attribué à ce lot.

## Implémentation et frontière d'hôte

- `src/Gopurs/BuildCache.purs` expose `exchange :: String -> Effect String`.
- `tools/build-cache/{contract,store}.go` possède les clés, validations, lectures,
  verrouillage et publications.
- `BuildCache.go` appelle directement ce code. `prepare-native-output.mjs` le
  copie en package `gopurs_build_cache` lors du bootstrap, comme le parseur FFI.
- `BuildCache.js` appelle `tools/build-cache.mjs`. L'auxiliaire Go est compilé à la
  première utilisation, puis réutilisé selon le contenu de ses sources et la
  plateforme dans `node_modules/.cache/gopurs-build-cache/`. Son lancement JS
  requiert Go lors de cette première compilation. Le protocole de verrouillage
  est prévu pour Linux et macOS.

La requête et la réponse sont du JSON UTF-8. La requête porte `schema: 1` et une
opération `snapshot`, `lookup` ou `publish`. Les réponses portent `schema: 1` et
un statut `snapshot`, `hit`, `miss`, `published` ou `error`. Une erreur de protocole
reste une réponse JSON ; une erreur de lancement du processus est une exception
`Effect`, à convertir en erreur `Aff` à la frontière du pilote.

Le bridge du stockage transmet des chaînes JSON. Les bridges applicatifs et les
représentations des valeurs PureScript sont vérifiés séparément par la comparaison
des générations JS/Go séquentiel/Go parallèle sur les mêmes TAST.

## Recette et clés

Exemple minimal de requête de stockage — un vrai pilote doit fournir **toutes**
les ressources et observations de résolution décrites ci-dessous :

```json
{
  "schema": 1,
  "operation": "snapshot",
  "spec": {
    "workspace": "/absolute/project",
    "output": "output",
    "compiler": [
      { "name": "compiler", "kind": "file", "path": "/immutable/gopurs-native" }
    ],
    "options": [
      { "name": "main", "value": "Main" },
      { "name": "rewriteLimit", "value": "10000" }
    ],
    "inputs": [
      { "name": "tast", "kind": "tast", "path": "output" },
      { "name": "ffi/Main/0", "kind": "file", "path": "src/Main.go" }
    ]
  }
}
```

Les noms d'options et d'observations sont uniques et triés par octets UTF-8. Les
chemins de workspace et de sortie sont absolus et résolus ; chaque entrée
conserve aussi son chemin demandé et sa destination de lien symbolique. Les
dates et tailles ne constituent jamais une preuve de validité des entrées.

| Observation | Contenu couvert |
| --- | --- |
| `file` | Présence/absence, chemin résolu et SHA-256 des octets |
| `directory` | Présence/absence, chemin résolu, noms/type/cible des entrées immédiates, dans l'ordre lexical |
| `tast` | Inventaire exact des `<module>/corefn.json` immédiats : nom, chemin résolu et SHA-256 ; sans décodage JSON |

`directory` n'est pas un hash récursif. Le producteur doit enregistrer chaque
répertoire parcouru par le résolveur : `.spago`, `spago.d`, packages et versions,
répertoire FFI optionnel et racines locales. Pour chaque module, il doit enregistrer
les candidats absents précédant le candidat retenu, ainsi que le fichier retenu
et l'ordre de priorité de la recherche. En cas de résolution entièrement absente,
tous les candidats consultés font partie de la recette. L'apparition d'une FFI
plus prioritaire ou d'un nouveau répertoire doit invalider la recette.

Le lot 03 devra produire cette trace depuis le résolveur réellement utilisé,
puis la vérifier avant le décodage complet des TAST. Le stockage ne déduit pas
les chemins FFI en analysant les modules. Une trace incomplète ne peut pas servir
de preuve de réutilisation.

Les options sémantiques comprennent la sélection des points d'entrée, le
répertoire FFI, la limite de réécriture, les directives et tout mode changeant
les sorties ou les diagnostics. Les directives actuelles de `App.loadDirectives`
sont embarquées dans le compilateur ; un futur fichier de directives devra
figurer comme entrée, y compris lorsqu'il est absent. Les options d'observation
telles qu'un profil d'allocations doivent forcer le parcours mesuré, ou posséder
un contrat explicite de rejeu. Le parallélisme n'est exclu d'une clé sémantique
qu'après preuve d'équivalence des sorties concernées.

### Identité du compilateur

Un fichier nommé `compiler` est obligatoire. L'identité combine les noms logiques
et les SHA-256 des ressources de `spec.compiler`, toutes présentes :

- natif : l'exécutable effectivement utilisé contient PBO, gopurs, runtime et
  parseur FFI natif ;
- JS : le bundle exécuté, ses éventuels modules externes, `ffi-runner.mjs`,
  `ffi_gen.wasm`, `wasm_exec.js` et l'auxiliaire de stockage effectivement utilisé.

Les changements locaux reconstruits sont donc couverts, même sans nouveau
commit Git. Une modification source non reconstruite n'est pas un changement
du compilateur exécuté. Le pilote doit figer l'identité de l'artefact qu'il
exécute : un chemin remplacé pendant le build ne prouve pas l'identité du code
déjà chargé. Des artefacts immuables ou une identité embarquée couvrant toutes
leurs sources/ressources devront assurer cette association lors du raccordement.

Les hôtes lisent le même format. Leurs identités d'exécutables distinctes causent
une invalidation conservatrice ; un format portable ne constitue pas à lui seul
une preuve que deux compilateurs différents peuvent partager leurs résultats.

### Construction des empreintes

`Key(domain, parts...)` hache avec SHA-256 la concaténation de chaque composant
précédé de sa longueur **en octets UTF-8**, en décimal ASCII, suivie de `:`.
Les premiers composants sont `gopurs-build-cache`, `1`, puis le domaine.

- `compiler` : paires nom logique / digest des ressources, triées par nom ;
- `options` : paires nom / valeur, triées par nom ;
- `directory` et `tast` : observations ordonnées décrites plus haut ;
- `build` : workspace, sortie, identité du compilateur, clé des options, puis
  nom/kind/path/resolved/state/digest de chaque observation triée.

Ce cadrage distingue les frontières entre composants et les valeurs vides. Les
blobs sont identifiés par le SHA-256 de leurs octets et leur longueur. Les clés
ne reposent pas sur l'ordre des propriétés JSON ni sur la représentation mémoire
d'un `Map` PureScript.

## Manifeste et données par module

`output/.gopurs-cache/current.json` est un pointeur versionné vers un objet
manifeste. Le manifeste contient :

- `format`, `schema`, `snapshot` (clés et observations) ;
- `outputs`, triées par chemin, avec `kind`, `module` et référence d'objet ;
- `modules`, dans l'ordre canonique de dépendances fourni par le producteur ;
- `diagnostics`, objet facultatif destiné aux diagnostics sémantiques rejouables ;
- `goMod: "create-if-absent"` et `goSum: "external"`.

Chaque mémo de module est indivisible : `name`, quatre clés (`prepared`,
`optimizerEnvironment`, `emitterEnvironment`, `callerDemands`) et les quatre
parties suivantes, dans cet ordre :

| Codec réservé | Données que le codec du lot 04 devra couvrir |
| --- | --- |
| `pbo/backend-v1` | `BackendModule` : nom, commentaires, imports, types ADT, groupes de bindings récursifs ou non et leurs `NeutralExpr`, exports/réexports, déclarations ADT/classes, imports FFI et types |
| `pbo/implementations-v1` | Map qualifiée vers `BackendAnalysis` et `ExternImpl` : usages, tailles, complexité, args, dépendances, résultats et corps/dictionnaires/constructeurs complets |
| `pbo/directives-v1` | Contributions `InlineDirectiveMap`, avec références, accesseurs et directives explicites |
| `gopurs/workers-v1` | `ModuleFunctions` : `fullName`, `fArgs`, `fRet`, `arity`, avec les `GoType` structurels |

`prepared` couvre le module après spécialisation réellement consommé par PBO.
`optimizerEnvironment` couvre les implémentations **et leurs absences**, les
corps inlinés, directives accumulées et paramètres d'analyse. `emitterEnvironment`
couvre les métadonnées globales, layouts/ABI, signatures des workers visibles et
signatures/contenus FFI. `callerDemands` couvre les demandes de spécialisation
transitives provenant des appelants : un fournisseur inchangé peut être invalidé
par un appelant modifié. Un hash global conservateur de chaque environnement est
admis avant d'établir des dépendances plus fines.

Le builder PBO possède déjà la republication des implémentations/directives sur
son chemin `onSkipModule`. Son chemin de skip ne demande pas de codegen : le
raccordement doit donc republier aussi les workers au niveau de gopurs, aux mêmes
barrières et dans le même ordre que l'émission normale. Une restauration en ordre
de fin des workers parallèles est incorrecte.

### Encodage portable des valeurs

Le stockage valide une grammaire JSON explicite ; les codecs métier devront en
plus valider les constructeurs, arités, champs et invariants avant toute
réutilisation. **La validation du stockage ne remplace pas ces codecs**, dont
l'implémentation et les vrais allers-retours PBO relèvent du lot 04.

| Valeur | Forme JSON |
| --- | --- |
| Unit | `["unit"]` |
| Boolean | `["boolean", true]` |
| Int | `["int", "-42"]`, décimal canonique signé sur 64 bits |
| Number | `["number", "8000000000000000"]`, bits IEEE-754 sur 16 chiffres hexadécimaux minuscules |
| String | `["string", "0061d800"]`, unités UTF-16 sur quatre chiffres hexadécimaux chacune |
| Array | `["array", [value, ...]]` |
| Constructeur | `["ctor", "NomStable", [argument, ...]]` |
| Record | `["record", [["labelUTF16hex", value], ...]]`, labels uniques triés |

Les codecs de Maps/Sets doivent encoder des listes ordonnées selon l'ordre
PureScript du type de clé, en rejetant les doublons. Les ADT utilisent des noms
stables et des champs documentés ; jamais les tags numériques d'un hôte. Les
producteurs rendent cette grammaire en JSON compact, sans espaces ni newline
final, pour obtenir les mêmes octets avec les deux hôtes. Les
chaînes conservent les surrogates isolés, les nombres conservent `-0`, infinis et
payloads NaN. Une valeur non représentable par un hôte fait échouer son décodage
de cache. Les erreurs de version, de shape ou d'invariant sont des cache misses.
Changer la grammaire ou les schémas métier exige une nouvelle version.

## Propriété et publication

Les chemins autorisés correspondent aux sorties du pilote :

| `kind` | Chemin relatif sous `output` |
| --- | --- |
| `module` | `purescript/<Module_avec_underscores>.go` |
| `ffi` | `purescript/<Module_avec_underscores>_ffi.go` |
| `runtime` | `gopurs_runtime/runtime.go` |
| `entry` | `main/main.go` ou `<Module>/main/main.go` |

La requête de publication doit inventorier **toutes** les sorties du build ; le
stockage ne connaît pas à lui seul les modules atteignables ni les entrées
sélectionnées. Les chemins d'objets, les sorties absolues/hors racine, les
doublons et les liens symboliques dans les destinations sont rejetés. Les
fichiers externes ne sont pas découverts par un balayage destructeur de `*.go`.

`go.mod` est créé avec `module gopurs/output` et `go 1.22` uniquement s'il manque.
Un fichier existant est conservé, y compris les `require`, `replace`, `toolchain`
et changements de version effectués par Go. `go.sum` est externe. Les deux
fichiers sont exclus du hash des sorties Go ; `go build` reste responsable de
valider le graphe de dépendances. Un hit gopurs n'autorise pas à éviter ce build.
La suppression de `go.mod` produit un miss. Ces règles sont celles du protocole ;
le lot 03 adaptera `Driver.Output.writeRuntime`, qui écrit actuellement le modèle
minimal à chaque build.

### Transaction

1. `snapshot` et `lookup` acquièrent le verrou du répertoire de sortie canonique,
   puis hachent les entrées. `lookup` vérifie le pointeur, le manifeste, les blobs,
   les sorties publiques et les entrées une seconde fois avant de rendre un hit.
2. En cas de miss, le producteur construit dans un **staging privé**. Les builds
   concurrents peuvent calculer en parallèle ; aucun ne doit émettre dans les
   sorties publiques avant `publish`.
3. `publish` reprend le verrou, recalcule la recette et exige `expectedKey` égal
   à la clé observée avant le calcul. Il copie les données du staging en objets
   vérifiés, valide le manifeste et recontrôle les entrées avant publication.
4. Le journal `output/.gopurs-ownership.json`, vérifié par SHA-256 et versionné,
   reçoit l'union des anciennes sorties et des octets qui vont être publiés.
   Il survit à une réinitialisation du cache d'objets.
5. Chaque fichier modifié est écrit temporairement, synchronisé et renommé
   atomiquement. Un contenu identique conserve sa date. Les fichiers obsolètes
   du journal sont retirés seulement si leurs octets correspondent à une version
   possédée ; une modification externe d'un fichier à retirer est un conflit.
6. Après un dernier contrôle des entrées, le nouveau manifeste puis le pointeur
   sont publiés atomiquement. Le journal est ensuite réduit à l'inventaire final.

Le verrou utilise `flock` sur `.gopurs-cache/write.lock`. Son inode reste stable ;
il est fermé et libéré par le noyau après la mort d'un processus. Les projets
ayant des sorties distinctes ne partagent pas ce verrou. Pour réinitialiser un
cache en conservant les sorties, supprimer `current.json` et les objets sous
verrou, en conservant le fichier de verrou et le journal de propriété.

La transaction rend les publications cohérentes pour les clients de ce
protocole ; elle n'est pas un renommage atomique de tout le répertoire Go. Un
consommateur externe lisant les fichiers pendant la publication doit être
coordonné au niveau du parcours de build. Les dates/hashes avant et après ne
remplacent pas une vue stable des sources : le futur pilote doit décoder les
octets inventoriés ou coordonner la génération TAST, afin d'éviter les edits
aller-retour pendant une compilation.

## Invalidation et reprise

- Cache absent, version inconnue, JSON invalide, objet absent/corrompu, entrée ou
  sortie modifiée : `miss`, avec une raison explicite et reconstruction.
- Échec du build avant publication : aucun nouveau manifeste valide.
- Arrêt pendant publication, y compris le premier build : le journal conserve
  les fichiers partiels ; un ancien pointeur ne peut être utilisé si les sorties
  ou le journal ne correspondent plus. Le build suivant répare/retire ces fichiers.
- Cache d'objets supprimé : reconstruction, avec propriété conservée par le journal.
- Journal absent mais manifeste précédent vérifiable : miss, puis récupération
  de l'inventaire engagé lors de la reconstruction. Les sorties d'une transaction
  interrompue exigent toujours son journal ; celui-ci doit être conservé.
- Journal de propriété corrompu : erreur explicite, sans suppression devinée de
  fichiers. Il faut retrouver un inventaire fiable ou reconstruire dans une
  sortie neuve. Ce journal est un état de propriété, distinct du cache jetable.
- Première migration d'une sortie historique : le pilote doit établir un
  inventaire fiable ou utiliser une sortie neuve ; il ne peut pas attribuer
  arbitrairement à gopurs tous les fichiers Go préexistants.

La recette, les codecs métier et l'inventaire complet sont des obligations du
producteur. Le stockage garantit leur intégrité et leur publication ; la preuve
que ces données décrivent bien le calcul est exercée lors des raccordements des
lots 03–05. La collecte des objets orphelins est différée ; leur présence ne
constitue jamais un hit.

## Vérifications

```sh
npm run test:cache
# Stockage seul :
go -C tools/build-cache test -race -count=1 -v -timeout 90s ./...
```

Les tests contrôlent les edits à date conservée, les inventaires TAST, les FFI
présentes/absentes et leurs racines, les ressources du compilateur, les options,
les liens symboliques, l'intégrité/version, les outputs obsolètes et externes,
`go.mod`/`go.sum`, les requêtes périmées, les valeurs wire limites, six éditeurs
concurrents et la reprise après mort réelle d'un détenteur du verrou. Le test
d'hôtes compile la FFI Go réelle contre le package préparé pour le bootstrap et
échange des publications JS → Go puis Go → JS.

Le [bilan du lot 02](testing.md#plan-v3--lot-02--contrat-du-cache-persistant-10-octobre-2026)
conserve les vérifications du bootstrap, des sorties applicatives et des snapshots.
