# Validation du plan FFI idiomatique

Travail du 12 novembre 2026. Les tests ci-dessous distinguent le compilateur
(hôte JavaScript/WASM ou natif), le Go produit et les services testés.

## Versions exactes

- gopurs : base `7a445c26ce1bd022adfdcad6b9c08945014cd305`, paquet de développement
  `0.1.0-rc.14` avec les modifications de ce travail.
- Frontend PureScript : base `105f5752b39cc0f6c99a40064295338bf57ad223`, avec
  préservation de `EffectFnAff` et résolution des synonymes dans `Desugar.hs`.
- Port Aff : base `37dc3ebb4a332d598ce6d33b01350c44cba4b7dc`, avec arbitrage
  résultat/annulation corrigé.
- b8x : base `c6c6e28b50719a5e6829e9a81bb63dd7ce48e7fa` ; interfaces PureScript
  et JavaScript des premières FFI migrées inchangées. Le complément Router
  conserve son API publique et adapte sa signature FFI privée sur toutes les cibles.
- Go `1.27.0 darwin/arm64`, Node `24.8.0`, Spago `0.93.45`, frontend typé
  `0.15.16`. Le numéro seul ne distingue pas le fork du compilateur officiel.
- SDK : pgx `5.11.0`, amqp091-go `1.15.0`, xxhash/v2 `2.3.0`, pvx
  `ac00bc32f6e7`. PostgreSQL de preuve : `16.14` ; client JS pg `8.16.3`.

## Commandes de reproduction

Depuis gopurs, après construction des hôtes et du parseur :

```sh
npm run build:ffi
npm run build
npm run build:native -- --keep-workspace
(cd tools/ffi-gen && go test ./...)
GO111MODULE=off go test -race ./runtime -count=1
npm run test:ffi
node --test tools/ffi-aff-e2e.test.mjs
node --test tools/ffi-package.test.mjs
npm run test:package
npm run check:docs
```

Le test E2E accepte `GOPURS_TEST_MODES=js` ou `native-parallel`. Il utilise le
port Aff local pour Go et le paquet Aff officiel pour l'oracle JavaScript :
les primitives natives du port ne sont pas toutes implémentées en JavaScript.
Le test de paquet installe le vrai tarball dans un répertoire isolé et vérifie
l'absence de dépendance aux sources du checkout du compilateur.

Depuis le port `gopurs-aff` :

```sh
node --test tools/aff-native.test.mjs
```

Depuis b8x :

```sh
node run/bak/go/tests/ordinary-ffi.mjs
node run/bak/go/tests/effects-and-subscriptions.mjs
B8X_POSTGRES_TEST_URL=... node run/bak/go/tests/postgres-nulls.mjs
B8X_RABBITMQ_TEST_URL=... node run/bak/go/tests/rabbitmq-idiomatic.mjs
```

Les runners services documentent leurs variables et leur portée dans
`b8x/run/bak/go/tests/README.md` et `rabbitmq-idiomatic.md`. Les preuves de ce
travail utilisent des services jetables locaux ; elles ne modifient pas les
bases applicatives. Les conteneurs de preuve sont supprimés après utilisation.
Les sources Go réelles et leurs adaptateurs sont testés sous `-race` ; cela ne
remplace pas une reconstruction ou un test fonctionnel de tout b8x.

## Contrats observés

- Signatures Go complètes identiques via WASM et natif, schéma versionné,
  diagnostics localisés et refus des formes incompatibles.
- Go ordinaire compilable séparément ; actions différées et répétables,
  erreurs, nil typés, panics, callbacks, contexte et arrêt de l'attente.
- Arbitrage résultat/annulation, publication avant acknowledgment, réentrance,
  libération unique des acquisitions tardives, fin des goroutines des fixtures.
- Handles identiques, conversions Foreign sans JSON ni coercition numérique,
  tableaux imbriqués, callbacks retournant des valeurs structurées, records
  fermés et structs locaux, substitution des paramètres des layouts élidés.
- PostgreSQL : référence JavaScript avant/après, valeurs nulles/absentes,
  sérialisation métier, résultats vides, erreurs, count, connexions dédiées,
  transactions et requête réellement active annulée.
- RabbitMQ : connexion, canal, queue, publication, callback persistant après
  inscription, annulation d'abonnement, fermeture et erreurs du SDK réel.
- Hash, Paseto, Unzip, identifiants, timer et signal : tests directs et via les
  adaptateurs, en préservant les politiques applicatives documentées.

## Bilan DX et limites qualifiées

| FFI pilote | Lignes avant → après | Références runtime avant → après |
| --- | ---: | ---: |
| RabbitMQ | 91 → 100 | 24 → 0 |
| PostgreSQL | 269 → 179 | 67 → 0 |
| Hash | 37 → 31 | 11 → 0 |
| Paseto | 96 → 57 | 26 → 0 |
| Unzip | 71 → 43 | 30 → 0 |
| Router | 64 → 22 | 1 → 0 |

Les interfaces Go de RabbitMQ permettent aussi des doublures ; le nombre de
lignes ne suffit donc pas à mesurer sa simplification. Aucun registre de
bindings ni option par fonction n'a été ajouté. Les fonctions de cleanup
retournées sont des closures Go ordinaires exprimant explicitement la propriété
d'une ressource, pas des appels à un helper réservé à gopurs.

Router utilise maintenant `unvariantF` côté PureScript pour extraire le label
et le contenu. Le dispatcher Go reçoit aussi la variante originale, transmise
au fallback sans reconstruction. Il ne dépend plus du runtime. Cette adaptation
conserve l'API publique et ne crée pas de correspondance structurelle implicite
pour les rangées ouvertes ou les variantes polymorphes. Le runner dédié et son
périmètre sont documentés dans `b8x/run/bak/go/tests/ordinary-ffi.md`.

Les newtypes sont instanciés par le frontend ; les ADT à un constructeur et un
champ utilisent le layout élidé déjà établi du backend, avec substitution et
garde récursive testées. Une nouvelle correspondance générale des ADT à
plusieurs constructeurs vers des sommes Go n'est pas annoncée. Les layouts de
structs importés non disponibles et les aliases externes non résolus ne sont
pas devinés. L'import direct d'`Aff` reste une intégration spécialisée du port ;
la voie ordinaire recommandée est `EffectFnAff`.

Les effets de SDK qui ignorent le contexte ne deviennent pas physiquement
interruptibles. Les callbacks déjà en vol peuvent terminer après la fermeture
d'un abonnement. Les divergences antérieures de RabbitMQ (acquittement), Signal
(one-shot) et Unzip (extraction disque non implémentée) sont documentées ; elles
ne sont pas présentées comme corrigées par le pont.

Les chemins générés ont été inspectés : les conversions structurées sont
statiques, les chemins dynamiques n'ajoutent pas de sérialisation JSON, et les
workers directs qui contourneraient une conversion structurelle ne sont pas
publiés. Aucune conclusion de performance applicative n'est tirée de ce travail.
Les baselines du README d'altbak.pub restent la référence pour une telle étude.

## Résultats finaux

Les suites suivantes sont vertes sur les derniers hôtes construits, après le
refus explicite des imports ordinaires `Aff` et des imports contraints :

| Vérification | Résultat |
| --- | --- |
| `npm run test:ffi` | 31 tests, 0 échec ; les cinq snippets Go du guide sont compilés et exécutés sous `-race` |
| Parseur Go et runtime | `go test ./...` et `go test -race ./runtime -count=1` réussis |
| Port Aff | Runner natif réussi, couvrant les 12 tests Go et les sous-cas bracket |
| `npm run test:ffi:e2e` | 3 tests, 0 échec ; hôtes JS et natif, release/dev, oracle JS officiel |
| `npm run test:ffi:package` | Installation offline avec cache npm vide ; deux hôtes, migration depuis le frontend officiel, six étapes par hôte réussies |
| Distribution et documentation | 61 fichiers dans le paquet ; 106 modules et références documentaires vérifiés |
| Frontend Haskell | 78 exemples CoreFn, 0 échec ; export JSON réel avant/après sur trois modules |
| b8x hors réseau | 6 paquets FFI ordinaires/générés et 10 paquets effets/abonnements réussis sous `-race` |
| Foreign | 27 combinaisons représentation/taille et cas imbriqués null/undefined réussis sous `-race` |

Les intégrations services ont été exécutées avant la dernière extension des
diagnostics, sans modification ultérieure des FFI applicatives : PostgreSQL
couvre dix cas comparés à JavaScript, quatre cas directs et deux cas de
ressources/annulation ; RabbitMQ couvre six tests directs, les ponts générés de
ses sept signatures et les 48 vérifications JavaScript existantes.

Le parcours du paquet confirme l'égalité des faits du parseur natif/WASM et les
étapes `official-javascript`, `cold`, `warm`, `signature-change`,
`invalid-signature`, `recovery` pour chacun des deux hôtes. L'erreur de signature
ne remplace pas le cache valide. Le binaire natif construit est installé dans
l'environnement isolé pour ce test ; sa construction depuis les sources reste
une étape distincte de l'installation npm offline.

Les cibles exécutées ici sont les programmes Go natifs `darwin/arm64` et les
oracles JavaScript sous Node. Le parseur exécuté en WASM est un composant du
compilateur : ces preuves ne qualifient pas les SDK PostgreSQL/RabbitMQ sur une
cible d'exécution WASM. Aucun build complet de l'application b8x n'est annoncé.

## Artefacts de cette validation

Le répertoire local conservé est
`/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/gopurs-ffi-package-F0bwy1`.
Il contient le tarball installé, `report.json` (commandes, sorties, versions,
statuts et étapes) et `artifacts.json` (chemins, tailles et empreintes). C'est
une preuve locale temporaire ; les runners du dépôt permettent de la reproduire.
La construction native finale est consignée dans
`/tmp/gopurs-idiomatic-native-final.log`.

| Artefact | SHA-256 |
| --- | --- |
| `gopurs-0.1.0-rc.14.tgz` | `9d9ce2366c570a1bbcbd53f4f0c3901f40efc567f128cd453e405c30444cbbc1` |
| `gopurs.js` | `094207bc6ebdfd6ec65443bdb612051197fc7c1dd47ce08344a080e754ddef82` |
| `gopurs-native` | `b64330f5c5da22cb56568ffe73222b93c5c2b1ff4ac8bb6ad28ef6eb9121ec52` |
| `ffi_gen.wasm` | `2c9ada9b14ba3b37cacbbf7249beca98f87aef107033f2884f5b7313d2912c59` |
| `purs` | `e071fce73bee0f1a97c4a9c1a2bb78956f307f916745196af084cf91ed62546e` |

## Complément Router : suppression de la dernière dépendance applicative

Le complément utilise l'éliminateur existant `unvariantF` dans Router.purs.
Aucune modification supplémentaire de gopurs, PBO, du TAST ou du runtime n'a
été nécessaire. Une recherche de `gopurs_runtime` dans les fichiers Go de
`b8x/src` ne trouve désormais aucune occurrence.

`node run/bak/go/tests/router-idiomatic.mjs` valide onze résultats identiques
entre l'ancien et le nouveau JavaScript, puis les hôtes gopurs JS/natif en
release/dev sous `-race`. Il utilise le vrai Router.purs, les bibliothèques
VariantF/Run et un Proem minimal ; le reste de b8x n'est pas reconstruit. Deux
tests supplémentaires compilent le fichier Router.go exact en Go ordinaire
sous `-race`. Les preuves sont conservées dans `/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/gopurs-tests-dDUpVA`.

Le parcours inclut un effet inconnu relayé par `Run.send` puis interprété plus
loin : la variante originale, son foncteur et la continuation sont préservés.
La signature FFI privée des autres cibles a été alignée. PHP et Rust (sept tests
en RC et sept en ARC) ont été exécutés ; Java et F# ont été relus mais non
exécutés, les environnements JDK/.NET étant absents.

## Correction de l'installation du frontend et parcours applicatif réel

Le premier `b -c` de l'utilisateur a révélé une lacune de déploiement : les tests
isolés utilisaient le frontend reconstruit, tandis que le shell interactif
lançait l'ancien fork installé dans `~/.local/bin/purs` (révision `3c8fcfd7`,
29 septembre). Une fixture reproduit son `EffectFnAff` développé en `EffectFn2`.
Elle confirme aussi qu'un changement entre deux binaires annonçant `0.15.16`
ne suffit pas à invalider le cache frontend.

Le binaire installé a été actualisé avec le frontend reconstruit de ce travail,
après sauvegarde de l'ancien dans
`/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/purs-installed-before-ffi-gtnf_6hl/purs`.
Le parcours exact `b -c` a ensuite réussi : 2 708 modules frontend, génération
Go, `go mod tidy` et sortie 0. Le log est `/tmp/b8x-ffi-real-build.log`. Ce
contrôle complète les fixtures ; il ne comprend ni `b -b` ni l'exécution
fonctionnelle de l'application. Aucun changement des signatures PostgreSQL
ou du code de sélection des effets n'a été nécessaire pour corriger ce défaut.

## Régression `Foreign` révélée par les tests applicatifs

Le premier lancement applicatif après `b -c` a révélé un rejet prématuré de
`undefined` dans `ForeignValue`. Une instrumentation temporaire, retirée après
capture, a identifié exactement `params[13][0].trace.cause`, tag runtime 0.
Ce champ vient de la conversion JSON d'un `Nothing` optionnel de l'EventStore.
Il est valide : le sérialiseur applicatif doit omettre cette propriété.

Le transport utilise désormais un marqueur natif privé distinct de `nil`,
sans supprimer la propriété. `Any` et `Box` restaurent `undefined` au retour
vers PureScript. Le port Foreign reconnaît ce marqueur lors de la conversion
JSON explicite : omission d'un champ objet, `null` pour une position de tableau.
Les fonctions et constructeurs algébriques restent hors de ce transport.

PostgreSQL prépare également les éléments des tableaux SQL via le callback
JSON applicatif. Les objets imbriqués ne sont plus laissés à l'encodeur JSON
implicite de pgx. Le type natif `pgtype.Array` conserve les dimensions des
tableaux multidimensionnels ; les formes irrégulières produisent une erreur.
Aucun import runtime n'a été réintroduit dans la FFI applicative.

Preuves ciblées : suite runtime sous `-race`, 31 tests FFI, port Foreign avec
oracle `JSON.stringify`, six tests Go PostgreSQL ordinaires, dix cas SQL réels
comparés à JavaScript, quatre cas de comptage, acquisitions et annulation.
La preuve PostgreSQL est dans
`/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/b8x-postgres-nulls-ziVy87`.
La preuve Foreign est dans
`/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/gopurs-foreign-json-records-pW6Na9`.

Deux limites SQL préexistantes restent distinctes de ce correctif : un paramètre
SQL directement `undefined` n'est pas normalisé en SQL NULL ; `rows.Values`
aplatit les résultats SQL multidimensionnels. Le test des paramètres vérifie
leurs dimensions côté serveur avec `array_ndims` et `to_jsonb`, sans annoncer
une modification du décodage des résultats.

Validation applicative finale : après reconstruction du compilateur natif et
régénération de b8x, `docker exec -w /var/www/b8x core-api-cli-1 bash -c
'./bin/run Test'` termine avec le code 0 : **287/287 tests passent**. Le log
complet est `/tmp/b8x-undefined-application-tests.log`. Cette exécution ne
lance pas le nettoyage global des bases de test de l'option `t -c`.
Le log de génération est `/tmp/b8x-undefined-build.log` et celui de la
construction du compilateur `/tmp/gopurs-undefined-native-build-final.log`.
