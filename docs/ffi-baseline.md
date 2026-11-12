# FFI idiomatique : état initial et contrat de preuve

Inventaire du 12 novembre 2026, avant migration applicative. Il accompagne le
[plan](../todo.md) et distingue les comportements observés des extensions à
valider. Les chiffres décrivent les sources manuscrites, pas le Go généré.

## Versions et périmètre

- gopurs : `7a445c26ce1bd022adfdcad6b9c08945014cd305`, paquet `0.1.0-rc.14`.
- b8x : `c6c6e28b50719a5e6829e9a81bb63dd7ce48e7fa`.
- Go : `go1.27.0 darwin/arm64` ; Node : `v24.8.0`.
- SDK déclarés dans `b8x/run/bak/go/module/go.mod` : pgx `v5.11.0`,
  amqp091-go `v1.15.0`, xxhash/v2 `v2.3.0`, pvx
  `v0.0.0-20210912195928-ac00bc32f6e7`.
- `b8x/src` contient 22 fichiers Go, dont la sonde autonome `test_pgx.go`,
  et donc 21 fichiers de FFI applicative. Six importent le runtime.

Les ports Aff, Foreign, Foreign.Index, Variant et Buffer constituent aussi
des dépendances pertinentes. Leur code définit parfois la représentation
du langage lui-même ; son accès au runtime n'a pas la même justification
que celui d'une opération applicative RabbitMQ ou PostgreSQL.

## Baseline chiffrée

Les chemins ci-dessous sont relatifs à `b8x/src`. Les nombres comptent les
occurrences textuelles qualifiées `gopurs_runtime.X`, y compris les types de
paramètres et de résultats. Les assertions de type comprennent les casts
de SDK légitimes : leur total n'est pas un objectif de suppression aveugle.

| FFI | Lignes | Références runtime | `Value` | `Box` | `Apply` | `FuncN` | Assertions Go | Tests de tags |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `Infra/Client/RabbitMq/RabbitMq.go` | 91 | 24 | 15 | 3 | 4 | 2 | 6 | 0 |
| `Infra/Client/Postgres/Postgres.go` | 269 | 67 | 31 | 9 | 13 | 6 | 19 | 8 |
| `Util/Crypto/Hash.go` | 37 | 11 | 6 | 1 | 2 | 2 | 0 | 0 |
| `Util/Crypto/Paseto/Paseto.go` | 96 | 26 | 12 | 4 | 6 | 4 | 1 | 0 |
| `Util/File/Unzip.go` | 71 | 30 | 12 | 6 | 8 | 4 | 0 | 0 |
| `Util/Run/Router.go` | 64 | 1 | 0 | 0 | 0 | 0 | 4 | 0 |
| **Total** | **628** | **159** | **76** | **23** | **33** | **18** | **30** | **8** |

Il existe 17 imports `EffectFnAff` dans ces FFI : sept RabbitMQ, cinq
PostgreSQL, deux Paseto, deux Unzip et un Hash. Les 18 constructions `FuncN`
se répartissent en 11 inscriptions de callbacks et sept annuleurs. RabbitMQ
partage `rabbitAff`, PostgreSQL partage `pgNonCanceler`. PostgreSQL appelle
aussi deux fois `Foreign_UnboxForJSON`, sans import runtime supplémentaire.

La configuration Go actuelle consiste en un module de dépendances partagé,
copié par `bin/_gopurs` dans l'output avant analyse des imports. Aucun registre
par fonction ni fichier de correspondance n'est présent. Les exports Go et
PureScript suivent la convention de nommage déjà reconnue par le parseur.

Reproduction des comptes depuis b8x :

```sh
python3 - <<'PY'
from pathlib import Path
import re
files = sorted(Path('src').rglob('*.go'))
print('Go files:', len(files))
for path in files:
    source = path.read_text()
    if '"gopurs/output/gopurs_runtime"' not in source:
        continue
    patterns = [r'gopurs_runtime\.', r'gopurs_runtime\.Value',
                r'gopurs_runtime\.Box\(', r'gopurs_runtime\.Apply\(',
                r'gopurs_runtime\.Func\d?\(', r'\.\([\w*]', r'\.Type\s*==']
    print(path, len(source.splitlines()),
          *(len(re.findall(pattern, source)) for pattern in patterns))
PY
```

## Signatures et décisions applicatives

| Domaine | Contrats actuels | Adaptation imputable au backend | Politique à conserver dans l'application |
| --- | --- | --- | --- |
| RabbitMQ | Handles opaques `Connection`/`Channel`, `EffectFnAff`, callback `EffectFn1 String Unit` | `any` puis assertions de pointeurs, `rabbitAff`, `Apply` de callback | URL, vhost, queue durable, message persistant, acquittements, fermeture du canal |
| PostgreSQL | `Effect` pool, handles paramétrés, callback `Foreign -> String`, `Array Foreign`, `EffectFnAff` | Thunk pool, goroutines et callbacks Aff, déballage conditionnel du handle, tags des paramètres, boxing des rows | Paramètres SQL, sérialisation objet par le callback, choix pool/connexion/transaction, fermeture rows, libération connexion |
| Paseto | `Foreign`, résultat `String` ou `Foreign` sous `EffectFnAff` | Goroutines et callbacks Aff, boxing | Clé hexadécimale, copie des claims, expiration h/d/m en RFC3339, validation du token |
| Unzip | `ArrayBuffer`, `EffectFnAff String/Unit` | Annuleur et appels de callbacks | Recherche HTML, textes des erreurs, fermeture du reader ; extraction disque actuellement non implémentée |
| Hash | Trois fonctions pures et `String -> EffectFnAff String` | Pont Aff pour xxhash sans erreur métier | Encodage hexadécimal fixe sur 16 caractères |
| Router | Handles polymorphes, callbacks, `VariantF`, `Run` | `ExtractVariant` et essais de réflexion sur la représentation | Dictionnaire de handlers, priorité du handler sur fallback |
| JSON Native | Callbacks construisant `Either`, `Foreign` | Callbacks typés déjà reconnus | `encoding/json` et son choix numérique `float64` |
| TaggedSum | `Foreign -> Foreign` | Réflexion sur une représentation dynamique | Ajout de `value: null` si la clé manque, conservation d'une valeur présente |
| Timer / Signal | `EffectFnN`, callback `Effect Unit`, handle `*IntervalId` | Callback avec argument Unit et résultat `interface{}` artificiels | Ticker, signal choisi, arrêt/désinscription |
| UUID / ULID / EventStore | `Effect String` ou newtype UUID | Argument Unit artificiel et retour `interface{}` | Génération et format des identifiants |
| Config, Cookie, HTML, String, Path | Valeurs, maps, tableaux et scalaires | Conversion des records/maps et tableaux | Configuration centralisée, normalisation, échappement, formats |

PostgreSQL convertit actuellement `Int` et `Number` en chaînes décimales,
`Boolean` en `"t"`/`"f"`, les objets avec le callback `encodeJson`, et laisse
les autres valeurs déballées. Transporter naturellement `[]any` ne doit pas
supprimer ces décisions SQL. Il faut distinguer null, propriété absente,
tableau vide, zéro et faux. Les objets JSON restent sous contrôle applicatif ;
le pont ne doit pas ajouter une sérialisation JSON universelle.

Les fonctions à résultats Go multiples sont nombreuses dans les SDK, mais
les FFI les consomment à l'intérieur de leurs wrappers : `Channel`, `Dial`,
`Query`, `Exec`, `Acquire`, `Values`, `Encrypt`, `ScanClaims`, `Open`, etc.
Leur absence des signatures exportées est précisément une conséquence de
la limitation actuelle du pont.

## Contrat initial et statut

| Forme PureScript | Forme Go cible | Statut à la baseline |
| --- | --- | --- |
| Fonction pure scalaire/tableau/callback | Fonction Go typée | Existante, preuve ciblée exécutée |
| `Effect a` | Fonction ordinaire retournant `a` | À étendre ; les thunks/arguments Unit existent déjà |
| `Effect Unit` | Fonction sans résultat | À vérifier explicitement pour différer l'appel |
| `EffectFnN` | Fonction Go non curryfiée et callbacks typés | Existante partiellement ; suppression de Unit à étendre |
| `EffectFnAff a` | `T` ou `(T, error)` | À construire |
| `EffectFnAff Unit` | Aucun résultat ou `error` | À construire |
| Opération annulable | Premier argument `context.Context` | À construire avec arbitrage et propriété explicites |
| Handle opaque | Pointeur/interface SDK | Partiellement existant ; identité, nil et durée de vie à prouver |
| `Foreign` / `Array Foreign` | `any` / `[]any` sans `Value` caché | À étendre et vérifier |
| Record fermé | Struct Go | À construire depuis les deux déclarations typées |
| Newtype | Type Go sous-jacent | Existant partiellement ; substitution des paramètres à corriger/valider |
| ADT / rangée ouverte / variante | Types Go avec convention explicite | Pas de correspondance générale validée |
| `Aff a` directement importé | Contrat asynchrone commun | Non pris en charge par le premier pont ; qualification ultérieure |

Un retour Go `error` n'autorise pas à lui seul la création d'un effet. Le type
TAST de l'import décide si le pont construit une action, puis la signature Go
doit être compatible. Les anciennes signatures runtime demeurent une voie
de compatibilité explicite pendant la migration.

## Première fixture, sans réseau

Le contrat de preuve est `foreign import probe :: Boolean -> EffectFnAff String`
face à `func Probe(fail bool) (string, error)`. Le corps incrémente un compteur
atomique, renvoie `"ready"` en succès et `errors.New("probe failure")` en échec.
Un getter ordinaire permet d'observer le compteur. Le Go manuscrit ne doit
importer que la bibliothèque standard et être testable dans son propre module.

Les vérifications sont : compteur nul après construction/application des
arguments, `"ready"` après une exécution réussie, message d'erreur inchangé,
exactement une notification par exécution, compteur incrémenté de nouveau
quand la même action est rejouée. Les variantes suivantes ajoutent `error`
seul, aucun résultat, résultat sans erreur, nil de pointeur et `error` contenant
un pointeur typé nil. Ce dernier suit Go : l'interface `error` non nil est un
échec ; aucune réflexion ne doit le transformer en succès. Un panic reste
un panic Go, sans conversion silencieuse en erreur métier.

Avant extension, reproduction exécutée depuis gopurs :

```sh
node --input-type=module <<'JS'
import { prepareFfi } from './output/Gopurs.FfiSupport/index.js';
const result = prepareFfi({moduleName:'IdiomaticAffProbe',path:'IdiomaticAffProbe.go'})
  ('IdiomaticAffProbe_')
  ('func Probe(fail bool) (string, error) { return "ready", nil }')();
console.log(JSON.stringify(result.decls, null, 2));
JS
node --test tools/ffi-bridge.test.mjs
```

Résultat observé : le modèle décodé conserve `ret: string` et perd `error`.
Les cinq tests existants du pont réussissent en 1,94 s : scalaires `int32`,
tableaux, callbacks, records, effets sous forme historique et fallback des
déclarations. Ces résultats sont une baseline, pas une validation d'Aff.

### Information TAST corrigée pendant la première preuve

Une compilation fraîche a aussi montré que `simplifyType` déroulait le
newtype `EffectFnAff` en `EffectFn2`, et qu'un synonyme `EffectFnCb` présent
dans le corps d'un newtype devenait `TypeApp Any`. Le compilateur possédait
les informations dans son environnement ; elles se perdaient à l'export.

La correction ciblée dans `purescript/src/Language/PureScript/CoreFn/Desugar.hs`
conserve le type qualifié `Effect.Aff.Compat.EffectFnAff` avec son argument
instancié, et résout les synonymes avec substitution et garde récursive.
Elle n'ajoute pas tous les newtypes à `dataDecls` : une évolution générale
demanderait de distinguer identité et représentation, pour éviter de traiter
un newtype effacé comme un ADT alloué.

Validation du fork : compilation de trois petits modules indépendants avant
et après, assertions sur le JSON exporté, puis 78 tests CoreFn réussis :

```sh
cd ../../purescript
stack test purescript:tests --test-arguments='--match=corefn'
```

Le binaire de tests nouvellement construit a effectivement été exécuté avec
`--match corefn` : 78 exemples, zéro échec en 1,54 s. Les quatre nouveaux tests
couvrent l'identité qualifiée, ses arguments, les alias, les callbacks dans
les newtypes et la garde récursive. La fixture fraîche produit maintenant
`String -> EffectFnAff String`, et `Wrapped Int` contenant `EffectFnCb Int`
devient bien `EffectFn1 Int Unit`. Cette preuve de contrat TAST ne remplace
pas l'exécution du pont asynchrone généré.

## Cycle de vie à prouver avant acquisitions

L'inscription du pont doit rendre immédiatement un annuleur, même si le corps
Go bloque. Chaque inscription possède son propre contexte et son propre
arbitrage ; succès, erreur et annulation ne doivent pas terminer deux fois.
Annuler le contexte n'implique pas qu'un SDK l'écoute. Sans contexte accepté
par l'opération, l'annulation ne peut arrêter physiquement le travail.

Une acquisition qui gagne la course après annulation ne peut simplement être
oubliée. La première politique sûre est de garder ces acquisitions non
annulables jusqu'à un contrat explicite d'acquisition/libération, ou de prévoir
un transfert/une libération explicite testé. Déduire automatiquement `Close`
serait incorrect : une connexion pgx dédiée doit être `Release`, et une
transaction a encore un autre contrat.

Le code du SDK installé confirme deux distinctions pour amqp091-go `v1.15.0` :
`PublishWithContext` rejette un contexte déjà annulé, mais n'interrompt pas les
I/O d'une publication déjà commencée (`channel.go`, documentation à partir de
la ligne 1852). `ConsumeWithContext` utilise le contexte pour toute la durée
de l'abonnement et demande de vider le canal des livraisons jusqu'à fermeture
(ligne 1449). Lui transmettre un contexte libéré dès le succès de l'inscription
arrêterait immédiatement l'abonnement ; ces durées de vie doivent rester
distinctes.

RabbitMQ conserve ses callbacks de messages jusqu'à fermeture du canal. Le
succès de l'inscription ne marque donc pas la fin de validité du callback.
L'annuleur actuel n'annule pas l'abonnement. Timer et Signal présentent aussi
une durée de vie supérieure à l'appel ; désinscription, double fermeture et
callback concurrent avec l'arrêt demandent des tests dédiés.

Les fixtures de cycle de vie doivent utiliser des canaux de synchronisation,
pas des délais arbitraires : blocage avant résultat, annulation, autorisation
du résultat tardif, observation de la notification/libération. Exécuter les
courses répétées sous `go test -race`, avec un compteur de ressources et une
attente déterministe de fin des goroutines.

## Intégration b8x et limites initiales

- `bin/build` suit la cible sélectionnée par `spago.yaml`. À l'inventaire,
  `spago.yaml -> run/bak/js/spago.js.yaml` et `output -> run/bak/js/output` :
  le lancer tel quel ne teste pas le Go. Les entrées Go sont
  `run/bak/go/spago.go.yaml` et `bin/_gopurs`.
- `bin/test` lance la suite applicative via le conteneur `api-cli` hors Docker.
  Sa variante `--clean-dbs` détruit les bases de test et ne constitue pas un
  prérequis de cette preuve locale.
- `run/bak/go/tests/postgres-nulls.mjs` possède un oracle JavaScript et une
  exécution Go sous `-race`, mais son test actuel attend encore un retour
  Promise et trois arguments de `_query`. La FFI actuelle reçoit quatre
  arguments, dont `encodeJson`, et retourne `EffectFnAff` : remettre ce runner
  en accord avant de le considérer comme validation de la migration.
- `node ../gopurs/gopurs-foreign/test/json-records.mjs` couvre des records
  dynamiques et nulls, indépendamment du réseau PostgreSQL. Son exécution
  actuelle reste à consigner lors de la modification du transport Foreign.
- Le Go RabbitMQ actuel utilise `autoAck=true`; le JavaScript appelle `ack`
  après le callback. Cette divergence préexistante n'est pas corrigée
  implicitement par le pont. Tester/documenter les politiques réellement
  conservées sur chaque cible.
- `_unzipToDirectory` échoue explicitement avec « Not implemented ». Retirer
  son wrapper ne signifie pas implémenter l'extraction disque.
- Le Router manipule une variante polymorphe et `Run`, pas un record fermé
  ordinaire. Son exception runtime doit rester justifiée tant que le contrat
  de transport de `VariantF` n'est pas établi.

Aucun test de service PostgreSQL/RabbitMQ n'a été exécuté pour cet inventaire.
Les succès d'un parseur WASM concernent l'hôte du compilateur ; ils ne prouvent
pas que ces SDK réseau fonctionnent comme cible Go/WASM. Aucune conclusion de
performance applicative n'est tirée de ces contrôles fonctionnels ; toute
comparaison future doit utiliser les baselines du README d'altbak.pub.
