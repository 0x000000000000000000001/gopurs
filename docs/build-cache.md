# Contrat du cache de build — version 1

Le lot 02 fournit le stockage exécutable et son protocole commun aux hôtes JS et
Go. Le lot 03 raccorde la vérification du build identique avant le décodage des
TAST, les entrées figées et la publication des sorties. Le lot 04 livré
réutilise l'optimisation et l'émission des modules après une préparation globale ;
sa validation et ses coûts figurent dans le [bilan des lots 04–05](build-incremental.md).
La [référence avant raccordement](build-baseline.md) mesure **36,902 s**
de médiane pour une relance gopurs identique.

## Usage quotidien et maintenance

Exécuter les commandes depuis la racine de l'application, avec le compilateur
TAST et le Spago de cette application sur le `PATH`. Les commandes npm du dépôt
gopurs sélectionnent sa propre toolchain de bootstrap ; vérifier les chemins avec
`command -v purs spago`. La référence b8x utilise Spago **1.0.3**, distinct du
Spago **0.93.45** installé pour construire le backend.

Lorsque `workspace.backend.cmd` de `spago.yaml` désigne gopurs, `spago build`
construit les TAST puis lance le backend. Le cache est actif par défaut :

```sh
spago build
go -C output mod tidy
# Pour une entrée sélectionnée explicitement par --main Main :
go -C output run ./main
```

Sans backend configuré dans Spago, lancer gopurs après la compilation frontend,
avec les mêmes options d'entrée/FFI à chaque invocation, comme dans le
[parcours d'installation](../README.md#getting-started). Dans b8x, le runner
habituel valide le cache Go avant de lancer le binaire ; ses ressources de
référence sont **GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB**.

### Sélectionner le découpage Go G06

Le layout mesuré utilise des ancres **1/8192**, avec un fichier ordinaire regroupé
par package. Il est disponible dans les trois compilateurs reconstruits.
Le lanceur `bin/gopurs` active désormais packages, overlay et manifeste par défaut
pour l'hôte Go natif, sans daemon. Les autres hôtes et l'appel direct du binaire
gardent une sélection explicite :

```sh
export GOPURS_GO_PACKAGES=1
spago build
go -C output mod tidy
```

Sans backend Spago configuré, appliquer le même environnement à l'invocation
habituelle de `bin/gopurs`, après production des TAST. Dans b8x, depuis la racine
du projet sur la cible Go :

```sh
b
bin/test
```

`bin/_gopurs` sélectionne le checkout voisin `gopurs/gopurs/bin/gopurs` ; le
launcher utilise le natif par défaut, JS avec `GOPURS_JS=1`, Rust avec
`GOPURS_RUST=1`. La [livraison G06](build-go-measurements.md) identifie les quatre
artefacts installés, y compris le WASM JS. Pour reconstruire depuis les sources,
utiliser `npm run build:native`, ou `npm run build` pour JS et
`GOPURS_RUST_PGO=0 npm run build:rust` pour le Rust non-PGO mesuré. Une modification
du résolveur Go demande aussi `npm run build:ffi` avec Go 1.27.0 avant le build JS.

La première invocation après changement de compilateur/layout est un miss de
cache. La publication migre les sorties possédées dans leur racine actuelle et
retire les anciens chemins par le journal ; elle conserve les fichiers externes,
`go.mod`, `go.sum` et l'inode du verrou. Le retour monolithique utilise
`GOPURS_GO_PACKAGES=0 spago build`. Les entrées `output/main/main.go` et
`output/<Module>/main/main.go` gardent leur rôle ; le runner suit leurs imports
sans connaître les noms de packages générés.

L'overlay C1c, activé par défaut par le lanceur natif, réutilise le plan Go
publié et ne rend que le module modifié si les signatures, les arêtes de
déclaration, les imports et l'init sont prouvés inchangés. Sinon le
planificateur complet reste l'oracle. Voir
[C2-5](build-c2-go-packages.md). Les valeurs explicites, dont
`GOPURS_GO_PACKAGE_OVERLAY=0` et `GOPURS_BUILD_MANIFEST=0`, restent prioritaires.
Le [chemin C2-7](build-c2-persistent.md) décrit la préparation compacte et les
mesures one-shot sans daemon.

### Daemon one-shot (sans watch)

Pour garder le `PreparedCache`, les backends PBO et le plan entre deux `b`,
activer le client daemon :

```sh
export GOPURS_DAEMON=1
export GOPURS_GO_PACKAGES=1
export GOPURS_GO_PACKAGE_OVERLAY=1
b
```

Le premier `b` démarre le daemon et paie le build froid ; les suivants
réutilisent le processus. Aucune surveillance de fichiers : `b` envoie un
jeton et attend le statut. Le daemon vit dans `<projet>/.gopurs-daemon/`
(`pid`, `log`, `status`) et redémarre si le compilateur, les arguments ou les
options changent. `GOPURS_DAEMON=0` ou l'arrêt du PID le désactivent.

### Pack de publication expérimental

Le pack regroupe par défaut les artefacts de modules dans 64 packs privés et un
index ; les sorties publiques restent individuelles. `GOPURS_PACK_PUBLICATION=0`
restaure le format individuel. Le gain froid mesuré est de **−34,2 %** sur trois
paires. La compaction par groupe ramène le cache à 64 packs après édition ; JS,
Go et Rust ont rejoué le corpus b8x packé avec des sorties exactes. Voir
[pack de publication](build-publication-pack.md).

Sur le corpus b8x figé, le gain jusqu'à `exec` est de **48–63 % après édition**.
Le backend seul prend davantage de temps après édition et reste à **~118 s à
froid** ; les [mesures ventilées](build-go-measurements.md) détaillent ce coût et
la faible marge du GOCACHE. La préparation incrémentale P est en pause.

### Forcer la génération complète

Avec le backend configuré dans Spago :

```sh
GOPURS_FORCE_BUILD=1 spago build
```

Pour appeler directement un backend déjà construit, depuis l'application :

```sh
/chemin/vers/gopurs/bin/gopurs --force --main Main
```

Conserver les options du build habituel. Le forçage refait préparation,
optimisation et émission ; il publie le résultat et conserve les dates des Go
dont le contenu est identique. Il ne force pas la recompilation du frontend ni
celle des packages Go déjà valides.

### Réinitialiser le cache réutilisable du projet

Sur Linux ou macOS avec Python 3, arrêter les builds de ce projet, puis exécuter
ce bloc depuis sa racine. `output` est résolu physiquement : un lien vers la
sortie Go de b8x est donc accepté.

```sh
python3 - <<'PY'
from pathlib import Path
import fcntl
import shutil

cache = Path("output").resolve(strict=True) / ".gopurs-cache"
with (cache / "write.lock").open("a+b") as lock:
    fcntl.flock(lock, fcntl.LOCK_EX)
    for name in ("current.json", "recipe.json"):
        (cache / name).unlink(missing_ok=True)
    objects = cache / "objects"
    if objects.exists():
        shutil.rmtree(objects)
PY
spago build
```

Cette procédure enlève le pointeur, la recette et les objets réutilisables. Elle
conserve `write.lock`, `output/.gopurs-ownership.json`, les TAST, les Go publiés,
`go.mod` et `go.sum`. Le prochain build reconstruit les objets et réutilise le
journal pour gérer les sorties obsolètes. Les éventuels stagings d'invocations
interrompues restent distincts du cache réutilisable.

Le journal appartient à sa racine de sortie ; il ne doit pas être transplanté
dans une autre racine. Pour une sortie historique sans journal, suivre les
[règles d'adoption et de reprise](#invalidation-et-reprise).

### Lire les diagnostics

- Un hit global termine avant le décodage des TAST et annonce zéro module
  optimisé/émis. Aucun Go n'est réécrit.
- Après édition, les compteurs `optimized`/`reused` et `emitted`/`restored`
  distinguent les calculs des restaurations ; la préparation reste globale.
- Les motifs d'invalidation indiquent l'entrée, la preuve ou l'artefact qui
  impose un recalcul. Un objet manquant/corrompu provoque un miss ; un conflit
  avec une sortie externe différente provoque une erreur.
- Les temps de passes du pipeline peuvent se chevaucher. Le total backend
  exclut Spago, le démarrage du processus et `go build` ; mesurer la chaîne
  directement pour évaluer le cycle utilisateur.

## Implémentation et frontière d'hôte

- `src/Gopurs/BuildCache.purs` expose `exchange` pour le stockage et
  `sessionExchange` pour le producteur du pilote, tous deux `String -> Effect String`.
- `tools/build-cache/{contract,store,packs}.go` possède les clés, validations,
  lectures, verrouillage, publications et le pack de publication par défaut.
- `BuildCache.go` appelle directement ce code. `prepare-native-output.mjs` le
  copie en package `gopurs_build_cache` lors du bootstrap, comme le parseur FFI.
- `BuildCache.js` appelle `tools/build-cache.mjs`. L'auxiliaire Go est compilé à la
  première utilisation, puis réutilisé selon le contenu de ses sources et la
  plateforme dans `node_modules/.cache/gopurs-build-cache/`. Son lancement JS
  requiert Go lors de cette première compilation. Le protocole de verrouillage
  est prévu pour Linux et macOS.
- `tools/build-cache/session.go` possède la recette FFI, les copies privées et
  l'inventaire complet ; `Driver.Cache` est son adaptateur PureScript.
- L'hôte Rust optionnel relie ce même package dans son archive Go de support FFI.
  L'archive contient un seul runtime Go pour le parseur et le cache.

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

Le résolveur de `session.go`, commun aux hôtes, produit cette trace en même temps
qu'il lit les octets FFI à consommer. Sa priorité reprend celle du résolveur PBO :
compagnon `.purs` → `.go`, puis `.spago`, `spago.d`, packages et versions `v*`
dans l'ordre lexical, répertoire `--ffi`, racine locale. Pour chaque racine, les
candidats sont `src/A/B.go`, `src/A.B.go`, puis `A.B.go`. Les chemins inspectés
qui ne sont pas des répertoires sont enregistrés aussi : une cible de symlink
peut devenir un répertoire sans changer les entrées de son parent.

La liste des observations est conservée dans `.gopurs-cache/recipe.json`. Au
prochain lancement, elle permet de vérifier le contenu avant tout décodage TAST.
Cette recette est un indice : seule une clé correspondant au manifeste vérifié
autorise un hit. Son absence, sa corruption ou sa désynchronisation après deux
publications concurrentes provoquent au plus un miss.

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
du compilateur exécuté. `CompilerSnapshot` copie l'exécutable ou le bundle et ses
ressources dans un répertoire privé, puis **réexécute cette copie avant le cache**.
Une installation remplacée entre le chargement du premier processus et la copie
fait ainsi exécuter la nouvelle copie ; les anciens objets mémoire ne sont pas
attribués au nouveau fichier. Les noms temporaires sont exclus de la clé du
compilateur ; seuls les noms logiques et les contenus comptent. JS inclut aussi
`process.versions` dans les options. Le processus final libère sa copie sur
succès ou erreur traitée ; un arrêt brutal peut laisser une copie orpheline,
qui n'est jamais réutilisée.

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

| Codec | Données couvertes par le codec du lot 04 |
| --- | --- |
| `pbo/backend-v2` (`v1` antérieur) | `BackendModule` : nom, commentaires, imports, types ADT, groupes de bindings récursifs ou non et leurs `NeutralExpr`, exports/réexports, déclarations ADT/classes, imports FFI et types ; le v2 ajoute l'enveloppe `body` et les lectures PBO optionnelles |
| `pbo/implementations-v1` | Map qualifiée vers `BackendAnalysis` et `ExternImpl` : usages, tailles, complexité, args, dépendances, résultats et corps/dictionnaires/constructeurs complets |
| `pbo/directives-v1` | Contributions `InlineDirectiveMap`, avec références, accesseurs et directives explicites |
| `gopurs/workers-v2` (`v1` antérieur) | `ModuleFunctions` : `reference` (symbole possédé), `fArgs`, `fRet`, `arity`, avec les `GoType` structurels et propriétaires des types nommés |

Depuis [G02](build-go-references.md), les anciens manifests `workers-v1`
restent lisibles pour la propriété des sorties et la reprise. Le producteur
exige `workers-v2` pour restaurer une signature et refuse les références
lexicales parmi les workers publiés. Un changement de propriétaire fait partie
de la clé d'émission, même lorsque le nom Go monolithique reste identique.

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

Le stockage valide une grammaire JSON explicite ; les codecs métier valident en
plus les constructeurs, arités, champs et invariants avant consommation des
valeurs. **La validation du stockage ne remplace pas ces codecs.** Le chemin
sans conversion PBO conserve les contributions non consommées sous forme
d'octets vérifiés ; une conversion qui en a besoin exige leur décodage métier.

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

## Parcours du pilote — lots 03 et 04

<a id="lot-04--raccordement-en-cours-de-validation"></a>

### Lot 04 — raccordement livré

Le lot 04 ajoute `Gopurs.Cache.Modules` et les codecs typés de
`Gopurs.Cache.Wire`. Les lots 04–05 sont validés sur `publicsync` ; les mesures,
les compromis et les preuves sont publiés dans le [bilan](build-incremental.md).

- `begin` copie les candidats d'un manifeste engagé dans `stage/prior`, après
  vérification de leurs objets. Les candidats restent privés et immuables pendant
  le calcul ; les sorties publiques ne servent pas à reconstituer leur contenu.
- La préparation reste globale. `prepared` hache le module réellement spécialisé,
  y compris les arguments statiques introduits par les appelants. La première
  clé `optimizerEnvironment` est amorcée par les directives et le contexte
  compilateur/options ; chaque module ajoute son empreinte au préfixe suivant.
  Cette preuve couvre les corps, contributions, absences et rangs, avec une
  invalidation conservatrice de tout le suffixe après un changement.
- L'optimisation `pboreads` conserve cette preuve et ajoute un repli plus précis lorsque
  seul le préfixe diffère : même contexte initial, même module préparé et mêmes
  demandes des appelants, puis validation des lectures PBO antérieures dans la
  vue courante. Une implémentation lue couvre sa présence/absence et, si présente,
  le couple complet analyse/corps. Les valeurs attendues viennent des anciennes
  contributions privées ; leurs empreintes par binding sont mémorisées pour le
  build. Un fournisseur déjà restauré prouve directement l'égalité de sa table,
  après résolution de sa visibilité et de la présence de la référence.
- Les directives lues sont celles de l'environnement initial de la conversion,
  incluant les forced private inlines. Les overrides locaux sont déterminés par
  l'entrée et les mêmes lectures. Une référence à un prédécesseur non finalisé
  met la tentative en attente, même si seule sa directive est consultée. Une
  absence ne peut pas remplacer cet état d'attente. Les observations sont
  propres à la tentative et ne sont transmises qu'à la publication canonique
  d'un résultat final. Une preuve absente ou refusée entraîne la reconstruction.
  Son seuil de performance face à `restoreprep` est atteint et les validations
  complètes multi-hôtes passent. Le bilan distingue ce gain local des mesures
  finales face au lot 03.
- Les empreintes préparées sont calculées dans un pool borné par `prepareJobs`,
  avec un contexte d'encodage privé par module, puis assemblées dans l'ordre
  canonique. Une édition FFI peut conserver ces empreintes lorsque `sameTast`
  prouve l'identité de l'inventaire TAST complet et du compilateur, et que le
  premier préfixe engagé correspond au contexte options/directives courant.
  La préparation globale s'exécute toujours. Toute modification de TAST, même
  limitée à un appelant, impose le recalcul des empreintes préparées.
- `buildModulesWithReuse` et `buildModulesParallelWithReuse` ajoutent un callback
  aux barrières canoniques de publication. L'API historique conserve sa sémantique.
  Les résultats réutilisés alimentent le même pipeline d'émission que les résultats
  recalculés ; l'ordre de fin des tâches ne détermine aucune publication.
- Si toutes les clés PBO correspondent, le pilote valide les backends et workers
  dans un pool borné, puis alimente directement l'émetteur en ordre canonique.
  Aucun optimiseur n'est démarré : les implémentations/directives, qui ne sont
  consommées par aucune conversion, sont republiées depuis les copies privées
  vérifiées. Leurs digests restent dans la clé d'émission. Une seule clé absente
  ou un échec de décodage des backends/workers rétablit le builder normal, qui
  décode les contributions avant de les rendre visibles aux conversions.
  Une nouvelle signature FFI peut toujours changer les workers et invalider
  l'émission des consommateurs, même si tous les TAST sont identiques.
- La clé d'émission couvre le résultat PBO complet, les métadonnées globales et
  layouts, les signatures de workers utilisées et la FFI résolue (chemin,
  présence, contenu), ainsi que le contexte compilateur/options.
  `workerkeys` remplace le préfixe de toutes les signatures visibles par la vue
  des références qualifiées externes de l'IR après les transformations locales
  de `CodeGen.prepareModule`. Chaque référence porte `Maybe FunctionInfo` dans
  le snapshot du lot : les absences font partie de la clé. Le rendu reçoit
  exactement les signatures présentes dans cette vue. Les workers restaurés
  sont fusionnés dans l'ordre canonique.
- `restoreprep` permet un lookup avant les transformations locales. Une clé du
  domaine `input-workers-contained-v1` couvre les références externes du backend
  d'entrée et ne peut être engagée qu'après avoir vérifié, sur un miss, que toutes
  les références transformées y sont incluses. Les transformations sont
  déterminées par le backend, les métadonnées source et le compilateur, qui sont
  couverts par cette clé ; elles ne consomment pas les signatures importées.
  Chaque référence d'entrée porte sa signature présente ou absente. Si une passe
  introduit une référence supplémentaire, le domaine `prepared-workers-v1`
  impose un lookup après préparation. Le rendu consomme dans les deux cas la vue
  transformée exacte. Le digest des workers restaurés vient de leur artefact
  privé vérifié, sans réencodage. La décision de performance
  figure dans [le bilan](testing.md#optimisation-du-rendu-après-le-bilan-allhit).
- Les quatre parties sont publiées avec les sorties Go via le stockage v1.
  `--force`, le profilage d'allocations et le marqueur de diagnostics désactivent également la
  réutilisation par module.

Depuis `pboreads`, la première partie est `pbo/backend-v2` : un record
`{ body, reads :: Maybe Observation.Reads }` encodé dans le même wire v1. Le corps
est encodé en premier et garde ses octets v1 après retrait de l'enveloppe ; les
trois autres parties conservent leur codec. Il reste quatre objets durables par
module. Le stockage accepte les manifestes `backend-v1` antérieurs pour préserver
la propriété des sorties et la récupération ; le producteur typé ne les restaure
pas comme des corps v2. La migration est éprouvée avec une sortie détenue modifiée
et conservation du même fichier de verrou.

Le format métier utilise des constructeurs explicites `WithGraph`,
`WithTypes` et `TypeRef`, dans la grammaire wire v1. Les références d'objets sont
des indices wire `Int`, distincts des nœuds dans la table. Les références d'objets et de
types sont locales à un artefact ; les arêtes d'objets doivent aller vers un indice
antérieur, les types sont reconstruits dans l'ordre avec leurs seuls prédécesseurs.
Les références absentes, cycliques ou vers l'avant produisent un miss.
Le prédicat scalaire `hex` parcourt directement la chaîne dans les bridges
JS/Go/Rust et accepte exactement `[0-9a-f]*`. La longueur des unités UTF-16 et des
bits IEEE-754, ainsi que leur réencodage exact, restent vérifiés par le codec
partagé. Le bilan `scalarhex` consigne les vecteurs et mesures de ce changement.

L'internage travaille sur les fragments JSON canoniques, pas sur la représentation
d'un hôte. Les Maps/Sets sont triés explicitement selon l'ordre PureScript, y
compris l'ordre UTF-16 des chaînes.

La table de types dispose d'un raccourci d'identité privé (`TypeMemo`) : il retient
les objets immuables pendant l'encodage d'un artefact et retrouve un indice déjà
attribué. `typeintern` conserve ce lookup, puis encode un type absent et interne
sa référence de graphe canonique dans un second `Pool`. Un type nouveau reçoit
son indice après ses enfants, dans l'ordre antérieur. Deux objets distincts de
même structure retrouvent les mêmes références et le même indice, sans ajouter
de nœud ; aucun ordre de comparaison ni identité d'hôte n'entre dans les octets
persistants. Le raccourci d'identité évite de redéployer les sous-types partagés.
Les bridges JS/Go/Rust existants ne définissent pas le schéma métier. La décision
de performance et la validation complète réussie figurent dans le bilan du lot 04,
désormais clos avec le surcoût de publication froide documenté.
À la lecture, une table typée
PureScript reconstruit chaque `NeutralExpr` une seule fois, dans l'ordre des
références antérieures, pour conserver le partage du graphe.

Les branches de `BackendSyntax` sont communes à l'instance `Codec` et aux
fonctions directes du parcours `NeutralExpr`. Cette organisation évite de
reconstruire le dictionnaire complet du codec à chaque nœud récursif, tout en
conservant les noms, arités, contrôles et l'ordre d'internage du format v1.

La grammaire est validée à la publication d'un contenu nouveau. Un objet déjà
présent dans le manifeste engagé, avec même codec, taille et SHA-256, conserve
cette preuve : le stockage relit et vérifie ses octets, sans reparcourir sa
grammaire. Le décodage métier reste obligatoire avant sa réutilisation par PBO.

Les compteurs indiquent `optimized`, `reused`, `emitted` et `restored`, avec les
raisons `no-entry`, `prepared-or-callers`, `optimizer-prefix`, `optimizer-reads`
ou `business-codec`.
Le compteur `prepared fingerprints` distingue les empreintes calculées et
réutilisées. La préparation, les clés des modules, la clé des métadonnées globales
et le traitement des artefacts sont chronométrés séparément.

### Parcours commun

1. `Main` entre dans le compilateur figé. `Driver.Config` lit les options, puis
   `loadDirectives` conserve ses diagnostics habituels, y compris sur un hit.
2. `Driver.Cache.begin` vérifie la recette, les blobs et les sorties sous verrou.
   Un hit retourne immédiatement : zéro décodage, préparation, optimisation ou
   émission. Les mesures et le message de hit rendent ce chemin observable.
3. Sur un miss, `begin` crée `.gopurs-cache/stages/build-*` et copie les TAST
   inventoriés. Leur digest combiné doit correspondre à l'empreinte initiale ;
   les entrées sont recontrôlées. PBO décode exclusivement ces copies. Un TAST
   rejeté par le décodeur fait échouer le build entier et ne peut produire un hit.
4. La préparation globale fournit les noms, chemins et imports étrangers des
   modules source. `plan` résout et copie les FFI ; ses empreintes sont calculées
   à partir des mêmes lectures que ces copies. Il complète la recette et vérifie
   qu'elle est toujours actuelle. L'émetteur lit uniquement les copies FFI,
   avec le chemin source canonique pour contextualiser les erreurs.
5. Le runtime, les modules, bridges et entrées sont générés dans le staging.
   `publish` construit leur inventaire, exige la clé prévue et applique la
   transaction ci-dessous. Le bracket du pilote libère ensuite son staging.

Une édition aller-retour des sources pendant le calcul ne peut donc attribuer
des octets intermédiaires à une clé différente : les octets consommés sont les
octets figés. Un changement encore présent à la publication fait échouer le
build avec une demande de relance. Les espaces `.purmeta` JS sont également
privés par invocation, via `GOPURS_PURMETA_DIRECTORY` dans le PBO.

### Calcul forcé et diagnostics

```sh
./bin/gopurs --force          # calcul complet, publication et écritures minimales
GOPURS_FORCE_BUILD=1 ./bin/gopurs
```

`GOPURS_ALLOC_PROFILE` force également le calcul, même lorsque les entrées et
sorties sont valides. Les compteurs et horloges de passes ne sont pas rejoués
sur un hit : le message annonce explicitement zéro module optimisé/émis.

Le diagnostic sémantique `Rebox missing` est marqué par `Gopurs.Diagnostics`.
Un manifeste marqué force le prochain calcul pour reproduire le diagnostic ;
le marqueur wire `gopurs/diagnostics-v1` contient `["unit"]`, et n'est pas un
texte à rejouer. Les erreurs de directives sont recalculées avant le lookup.
Les erreurs de décodage, FFI ou émission empêchent toute nouvelle publication.
Les options de parallélisme restent dans la clé, avec une invalidation
conservatrice lors d'un changement de configuration.

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
La suppression de `go.mod` produit un miss. `Driver.Output.writeRuntime` écrit
seulement le runtime dans le staging ; la création éventuelle de `go.mod` est
désormais confiée au stockage lors de la publication.

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
   atomiquement. Un contenu identique conserve sa date. Après jonction des
   writers, les répertoires des sorties et leurs ancêtres jusqu'à la racine de
   sortie sont synchronisés de bas en haut, y compris pour les fichiers déjà
   identiques lors d'une reprise. Les fichiers obsolètes du journal sont ensuite
   retirés seulement si leurs octets correspondent à une version possédée ; une
   modification externe d'un fichier à retirer est un conflit.
6. Après un dernier contrôle des entrées, le nouveau manifeste puis le pointeur
   sont publiés atomiquement. Le journal est ensuite réduit à l'inventaire final.

Le staging des objets utilise au plus huit tâches, chacune écrivant son
emplacement fixe du manifeste ; les octets identiques peuvent viser le même
digest. Toutes ces tâches sont terminées avant de publier le journal ou de
rendre une erreur. `publicpool2` utilise ensuite le même pool borné à huit tâches
pour les sorties publiques, après vérification des conflits et engagement durable
du journal d'union. Chaque tâche revérifie sa destination unique et les octets de
son objet ; toutes sont jointes avant suppression, contrôle final des entrées,
commit ou retour sur erreur. Une publication partielle reste récupérable depuis
le journal.
`batchsync` conserve la synchronisation de chaque fichier objet, puis
regroupe celle du répertoire `objects` après jonction de toutes les tâches et
validation du manifeste, avant le journal. Cette barrière s'exécute aussi lorsque
tous les objets existaient déjà, afin de couvrir les renommages laissés par un
staging interrompu. `publicsync` regroupe aussi les barrières des répertoires
publics, après jonction des writers et avant toute suppression ou commit. Chaque
fichier conserve sa synchronisation avant renommage ; chaque répertoire et ses
ancêtres sont synchronisés une fois. Un échec conserve le journal d'union et le
pointeur précédent. Le manifeste et le pointeur utilisent toujours `atomicWrite`,
avec leur synchronisation immédiate du répertoire. Le pointeur est engagé en
dernier, sous le même verrou stable.

`tools/build-cache-output.test.mjs`, inclus dans `npm run test:cache`, éprouve
les barrières et leur ordre, une erreur de synchronisation puis un SIGKILL avant
la première barrière publique. Il utilise des copies instrumentées du stockage
réel et conserve les workspaces en cas d'échec, ou avec
`GOPURS_TEST_KEEP_WORKSPACE=1`. Les tests macOS/Linux sous race detector et les
paires froides de `publicsync` passent ; le candidat mesuré est installé et sa
validation complète multi-hôtes réussit. Les lots 04–05 sont clos ; leur
[bilan](build-incremental.md) documente aussi le surcoût froid face au lot 03.

Le verrou utilise `flock` sur `.gopurs-cache/write.lock`. Son inode reste stable ;
il est fermé et libéré par le noyau après la mort d'un processus. Les projets
ayant des sorties distinctes ne partagent pas ce verrou. Pour réinitialiser un
cache en conservant les sorties, supprimer `current.json` et les objets sous
verrou, en conservant le fichier de verrou et le journal de propriété.

La transaction rend les publications cohérentes pour les clients de ce
protocole ; elle n'est pas un renommage atomique de tout le répertoire Go. Un
consommateur externe lisant les fichiers pendant la publication doit être
coordonné au niveau du parcours de build. Les dates/hashes avant et après ne
remplacent pas une vue stable des sources : le pilote décode les copies des
octets inventoriés, afin de couvrir les edits aller-retour pendant le calcul.

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
- Première migration d'une sortie historique : le pilote compare les fichiers
  des espaces générés à une génération fraîche. Il adopte seulement les chemins
  et octets identiques. Un fichier ancien inconnu ou différent provoque une
  erreur explicite et demande une sortie contenant seulement les artefacts du
  frontend. Les assets externes sont conservés. Une fois le journal établi,
  seuls ses fichiers obsolètes peuvent être retirés ; les fichiers externes
  ajoutés ensuite restent hors inventaire.
  Si un nouveau chemin généré rencontre un fichier externe différent, la
  publication échoue avant toute écriture publique (`unowned output differs`).

La recette, les codecs métier et l'inventaire complet sont des obligations du
producteur. Le stockage garantit leur intégrité et leur publication ; la preuve
que ces données décrivent bien le calcul est exercée lors des raccordements des
lots 03–05. La collecte des objets orphelins est différée ; leur présence ne
constitue jamais un hit.

**Candidat G06 multipackage :** après une publication réussie, le pilote effectue
un entretien des objets sous le verrou stable. Il relit le manifeste courant,
qui peut être celui d'un autre build, et exige un journal de propriété égal à
ses sorties. Il conserve le manifeste, les sorties, les artefacts métier et de
diagnostic ainsi que les sources compressées G05. Seuls les fichiers réguliers
de `objects` nommés par un SHA-256 et absents de cet ensemble sont retirés ;
les stages privés et les fichiers de contrôle restent en place. Une publication
incomplète suspend cet entretien. Les lectures de candidats sont protégées par
le même verrou et copient leurs données dans le staging avant de le rendre.
Un échec d'entretien conserve des objets jetables sans invalider la publication
réussie. Le parcours monolithique conserve la politique des lots précédents.
Les mesures d'acceptation de ce candidat sont suivies dans
[G06](build-go-measurements.md#diagnostics-suivant-le-pilote-3).

## Vérifications

```sh
npm run test:cache
npm run test:cache-driver # après reconstruction des compilateurs
npm run test:cache-sources
npm run test:cache-context
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

`compiler-snapshot.test.mjs` remplace réellement le bundle/binaire après son
chargement et vérifie le marqueur exécuté et son SHA-256 après réexécution.
`driver-cache.test.mjs` compare les transitions à une génération neuve sur les
mêmes TAST et entre JS, Go séquentiel et Go parallèle : hits sans passes, dates,
edits à date conservée, ajouts/retraits, priorités FFI, points d'entrée, fichiers
altérés, erreurs, SIGTERM, publications concurrentes, profils et corruption.
`GOPURS_TEST_RUST=1` ajoute l'hôte Rust ; `GOPURS_TEST_KEEP_WORKSPACE=1` conserve
les logs et inventaires. Les tests Go du producteur couvrent aussi les edits
aller-retour, les diagnostics et l'adoption historique.

`driver-source-cache.test.mjs` couvre les corps privés/exportés, directives,
layouts et demandes transitives de spécialisation sur 14 états de sources.
`driver-context-cache.test.mjs` vérifie 16 états d'options/ordonnancement/entrées,
étendus à 20 avec un catalogue de versions réelles de compilateur et de runtime.
Chaque état compare l'incrémental au frais sur les mêmes TAST, les quatre parties,
les clés, les octets Go, le hit sans réécriture et l'exécution de l'oracle.
`tools/cache-context-runtime.mjs` construit les variantes privées ; la
[procédure et les preuves](testing.md#plan-v3--lot-05--transitions-de-sources-campagne-en-cours-12-octobre-2026)
décrivent le catalogue, les quatre hôtes et la conservation des workspaces.

Le [bilan du lot 02](testing.md#plan-v3--lot-02--contrat-du-cache-persistant-10-octobre-2026)
conserve les vérifications du bootstrap, des sorties applicatives et des snapshots.
Le [bilan du lot 03](testing.md#plan-v3--lot-03--build-identique-vérifié-10-octobre-2026)
consigne les transitions, la parité b8x, les gains appariés (**35,732 → 1,257 s**
pour gopurs ; **38,145 → 3,958 s** pour la chaîne) et le coût de publication initiale.
