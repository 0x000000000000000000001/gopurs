# Références Go qualifiables — G02

**Validé le 13 octobre 2026, suite G à 3/7.** Les identités et le rendu préparés
ici prolongent le [plan G01](build-go-package-plan.md).
Le pilote conserve le layout monolithique ; son raccordement au plan de packages
est G03. Les preuves sont conservées sous
`gopurs-go-packages-g02-20261013/` dans le répertoire temporaire OpenCode.

## Identités portées par le générateur

`GoSymbol` contient `{ owner, name }`. Le propriétaire est une unité source
`generated/<préfixe-module>` ou `ffi/<préfixe-module>`, fournie par le producteur.
Le placement lui associera un chemin de package et un nom visible. Il ne se
déduit pas du découpage textuel d'un nom de fonction.

`GoReference` distingue `LocalReference`, pour les workers lexicaux, et
`SymbolReference`, pour les workers publiés, getters et helpers possédés.
`FunctionInfo.reference` remplace `fullName`. Les appels directs, signatures FFI,
workers de module, workers de consommation et Rebox utilisent ce contrat.

Les déclarations `GoCachedValue`, `GoStructDecl`, `GoFunctionDecl` et
`GoForeignGetter` portent leurs symboles. Une inscription `GoInitDecl` porte
`register/<constructeur>` ; l'ajout d'une autre inscription ne renumérote pas
cette identité. `Printer.declaredSymbols` expose les déclarations nommées,
y compris les trois noms `cache_`, `once_` et `Get_` d'une valeur paresseuse.
`Printer.initializerSymbols` expose les identités d'inscription dans l'ordre source.

Les constructeurs et leurs accès possèdent le symbole du type Go unique.
`StructPointer` conserve séparément :

| Donnée | Identité |
| --- | --- |
| `fullName` | type PureScript qualifié |
| `baseStructName` | identité utilisée pour calculer le tag runtime |
| `symbol` | déclaration nommée Go et son propriétaire |
| `structName`, `fullPath`, `typeArgs` | nom historique, instanciation rendue et arguments structurés |

`TypeNamed`, `TypePointer`, `GoTypeApply` et l'assertion avec `GoType` évitent
d'effacer ces références aux frontières des types. `TypeRecord` et
`GoRecordFieldAccess` conservent le label PureScript. Les noms monolithiques,
tags, ordre des champs et représentations runtime sont préservés. La convention
exportée `P_`/`PX_` et sa correspondance réflexive restent le contrat G03/G04
décrit dans G01.

## Liaison des fragments et FFI

Un fragment opaque ne peut être résolu isolément : le même nom peut être global,
paramètre local, champ de struct ou clé de map. `GoRaw` conserve son texte et ses
imports de compatibilité ; `GoSource.bindSources` lie les unités complètes, une
fois leurs portées disponibles, avec le parser et **`go/types`**, corps compris.

Entrée `BindRequest` : chemin de package, version Go, architecture cible,
sources possédées avec identités d'initialisation, exports `gc` des dépendances
en octets base64. Le résolveur ne lit pas le disque et ne lance pas Go. L'audit
des corpus utilise les exports réels figés par G00.

Sortie `gopurs/go-source-v1` : `BoundGoCode`, dans l'ordre canonique des chemins,
avec texte exact, imports/alias, identifiants réservés, déclarations et
références attribuées. Les positions sont en unités UTF-16, calculées en une
passe sur les positions UTF-8 ; elles ne servent pas d'identité aux déclarations
nommées. Les directives `//line` ne déplacent pas ces positions physiques.

- Les références de package sont qualifiables ; les champs et méthodes restent
  des membres et conservent un lien vers leur déclaration propriétaire.
- Une méthode générique utilise l'objet d'origine et le type receveur, pas le
  nom de sa variable receveur. Les méthodes d'interface appartiennent à
  l'interface ; elles ne créent pas de déclaration `func` autonome.
- Un champ embarqué définit un champ implicite et référence un type au même
  identifiant. Le type est qualifié dans la déclaration ; les sélections du
  champ restent des sélections.
- Les `init` générés exigent leur identité explicite. Les `init` FFI sont
  ordonnés dans l'unité FFI entière ; les initialisations `_` ont un jeton
  d'ordre, sans nom référençable.
- Les variables locales, paramètres génériques, labels et pré-déclarés Go
  restent locaux. Les imports, y compris `_`, sont conservés.

`Printer.printBoundGoCode` vérifie les spans, consulte le placement et rend les
références selon le consommateur. Une référence ou un alias manquant, une
définition placée ailleurs ou un alias rencontrant un identifiant réservé
produit une erreur. `GoImports.boundImports` déduit les nouveaux imports des
références réellement qualifiées. L'assemblage des fragments de packages et
des clauses d'import appartient à G03.

## Bridges et cache

Le contrat de liaison est implémenté une fois dans `tools/ffi-gen/bindings.go` :
WASM l'expose par `bindGoSource`, Go par `gopurs_ffi_parser.Bind`, Rust par
l'opération `2` du même C ABI. Les opérations d'extraction/préparation restent
`0`/`1`. `prepare-native-output.mjs` embarque aussi `bindings.go`.

Les tests comparent les JSON complets entre les trois transports, avec Unicode,
types génériques, interfaces, champs embarqués, vrais exports Go et import
d'effet seul. Le test Rust compile le bridge de production et son archive Go,
avec une petite enveloppe `Value`/effet. Le bootstrap Rust complet est également
reconstruit sans PGO et rejoué sur les fixtures.

Les signatures utilisent **`gopurs/workers-v2`**. Le codec portable transporte
référence, propriétaire et types structurels. La clé d'émission change si le
propriétaire d'un worker référencé change, même à nom monolithique égal.
`readWorkers` exige le codec v2 et une référence globale non vide. Les manifests
v1 restent lisibles pour la propriété et la reprise ; leurs signatures ne
sont pas acceptées comme données v2. Voir le [contrat du cache](build-cache.md).

## Preuves et portée

- Campagne initiale : **12 comparaisons b8x**, quatre états sur les mêmes
  **2 704 TAST**, JS/Go séquentiel/Go parallèle. Les **3 010 fichiers Go** et
  les objets workers de chaque état sont byte-exacts entre hôtes ; le Go est
  aussi exact face à la référence v3.
- Liaison exhaustive de la base : **2 997 unités**, **185 671 symboles**,
  **832 129 références ordinaires** ; bootstrap : **625 unités**,
  **65 338 symboles**, **278 461 références ordinaires**. Les champs, méthodes,
  imports et références natives ont des compteurs distincts. Ces nombres de
  symboles ne sont pas les nombres de groupes de déclarations G01.
- Les trois éditions sont aussi résolues avec leurs sources et exports figés.
  Toutes les cibles générées et tous les spans rapportés sont contrôlés.
- La fixture de qualification produit `42` après séparation en deux packages ;
  la substitution textuelle sans portée échoue et reste conservée.
- Les trois snapshots de `FFIIntegerReturns`/`NativeRecordWorkers` sont examinés
  par empreinte et rejoués sans mise à jour, avec compilation/exécution Go, sur
  JS, Go séquentiel, Go parallèle et Rust.
- La campagne finale b8x ajoute **six passages exacts** : natif forcé, identique
  sans réécriture, édition feuille incrémentale/forcée, puis base JS et Rust.
  L'inode du verrou reste **286740996**. Les signatures v2 complètes sont aussi
  byte-exactes entre ces trois hôtes.
- La migration réelle v1 → v2 couvre **61 modules**, dix commandes, modification
  de source, reprise depuis le manifeste après retrait privé du journal,
  identique, forcé et exécution. Le premier candidat échoue sur cette reprise ;
  le candidat corrigé reconstruit la propriété et conserve l'inode **287016135**.

La campagne b8x finale et la migration utilisent `candidate-2`. Le dernier
ajustement, limité au comptage UTF-16 de `bindings.go`, remplace `utf16.RuneLen`
par son calcul pour les runes valides : le contrôle des modules déclarant Go
1.22 refusait cet appel de bibliothèque ajouté en 1.23. `candidate-3` conserve
les mêmes sources PureScript et les mêmes fonctions d'extraction/préparation.
Ses transports sont reconstruits puis vérifiés par les tests JSON complets,
le race detector et **46 commandes de replay strict** sur les quatre hôtes.
Les empreintes et l'audit exact des sources figurent dans `final-reference/`.

Les premières tentatives ont révélé et conservent trois cas de liaison :
`init` traité comme un nom ordinaire, méthode d'interface traitée comme une
fonction autonome, puis référence de type effacée par un champ embarqué.
Les assertions Rebox et le vecteur de codec ont également nécessité une
adaptation de leurs données de test ; les logs initiaux restent disponibles.
La première tentative de migration attendait à tort un échec avec un journal
encore valide : celui-ci permettait déjà la publication. La deuxième cible
explicitement la reprise à partir du seul manifeste ancien.

Le JSON exhaustif de diagnostic représente **1 064 194 252 octets** pour b8x
et **463 386 924 octets** pour le bootstrap. L'audit ne le publie pas dans le
cache : il en conserve empreintes et compteurs par fichier. G03 doit exploiter
le graphe côté Go et transporter un plan compact lors du raccordement au
pilote. Le transport JSON intégral d'un corpus n'est pas validé via WASM/Rust.
Les budgets mémoire/stockage et seuils temporels G restent ceux de G00 ; ces
validations fonctionnelles, lancées avec d'autres contrôles, ne constituent pas
des mesures de performance appariées.

## Reproduire les contrôles ciblés

```sh
npm run build
npm run build:ffi
node --test --test-concurrency=1 tools/go-source.test.mjs \
  tools/module-codec.test.mjs tools/emitter-worker-keys.test.mjs
```

Dans `tools/ffi-gen` puis `tools/build-cache` : `go test -race -count=1 ./...`.
`GOPURS_TEST_EVIDENCE=<répertoire existant>` conserve les programmes avant/après
et les réponses des trois transports produits par `go-source.test.mjs`.
