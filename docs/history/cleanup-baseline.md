# Baseline de nettoyage — 29 octobre 2026

Cette baseline sert au plan [cleanup.md](../../cleanup.md). Elle fige ce qui ne
doit pas changer pendant les lots de structure : les snapshots de référence,
les binaires du compilateur et les versions d'outils.

## Référence

| Élément | Valeur |
| --- | --- |
| Commit de base | `8361c05f0a0517d55be7ce8be7600cec198c33ee` |
| Branche de base | `release/0.1.0-nix` |
| Branche de travail | `cleanup/v2` |
| Node | `v24.8.0` |
| npm | `11.6.0` |
| Go | `go1.27.0 darwin/arm64` |
| Spago local | `0.93.45` |
| `purs` local | `0.15.16` |

Les binaires listés ci-dessous sont des artefacts locaux non versionnés. Ils
ne sont pas une référence de publication ; ils documentent l'état qui a servi
à la vérification initiale.

## Empreintes

Les snapshots sont versionnés : leur empreinte agrégée est reproductible depuis
le commit de base.

| Référence | Fichiers | SHA-256 agrégé |
| --- | ---: | --- |
| `tests/passing-snapshots` | 405 | `b045108a31c7164efb7b346a5737c3bd38c382274b2ee59556b093d51f5e27b6` |
| `tests/passing-snapshots-dev` | 405 | `137b183d6693ab58b10b35d082d373bc022323e9b7f59f96d7c0dee4c3e0a4b7` |
| `tests/passing-package-snapshots` | 9 | `afc1022e6561b37271bb390e17cf89f27d9b090d37b3f7c695a1566dc6bbbd74` |
| `bin/gopurs.js` | 1 | `e7c761c44aef7276b335265459c03cfe6551b7ed8ec00f76f8233aa6fbe30b3b` |
| `bin/gopurs-native` | 1 | `eb97922721f15c3e742ff14335a8294b6005ca321ae2e9a57bb8847776e46298` |
| `bin/gopurs-rust` | 1 | `aa64a8939964b736ce2736b022d2925d8ed969eb110782b1c938c056257a4983` |
| `tools/ffi_gen.wasm` | 1 | `d87a0cad6a4ca7dcbcd67d07ccb72283eb5150b7aa17a6cef619573b5e935d3e` |
| `tools/wasm_exec.js` | 1 | `0c949f4996f9a89698e4b5c586de32249c3b69b7baadb64d220073cc04acba14` |
| `package-lock.json` | 1 | `4e98ddb1c45abb67459c49f51936b9cf94449753c3132c781c2dfb90ae93e248` |
| `spago.lock` | 1 | `cd5871691818bf2bf788a4eb2f3fa1e4705f1e858ceaf7e80e4b914fb8b5770b` |

Commande d'agrégation, exécutée depuis la racine :

```bash
find <dir> -type f -print0 \
  | LC_ALL=C sort -z \
  | xargs -0 shasum -a 256 \
  | shasum -a 256
```

La dernière ligne donne le SHA-256 agrégé ; la liste intermédiaire n'est pas
conservée, car les snapshots versionnés en sont la source.

## Contrôle fonctionnel

La fixture `ClassRefSyntax` a été compilée et exécutée le 29 octobre 2026 :

```bash
./bin/test ClassRefSyntax
```

Résultat : 1 fixture réussie, snapshot inchangé, exécution Go réussie.
Ce contrôle prouve que la chaîne Spago → gopurs → Go fonctionne dans
l'environnement de la baseline. Il ne remplace pas la campagne complète.

## Garde-fou d'identité

`tools/check-output-identity.mjs` compare deux arbres de sortie ou un arbre
avec un manifeste :

```bash
# Créer un manifeste à partir d'un arbre conservé
node tools/check-output-identity.mjs snapshot /tmp/run-a/output/purescript /tmp/run-a.manifest.json

# Comparer un second run au manifeste
node tools/check-output-identity.mjs compare /tmp/run-b/output/purescript /tmp/run-a.manifest.json

# Comparer deux arbres directement
node tools/check-output-identity.mjs diff /tmp/run-a/output/purescript /tmp/run-b/output/purescript
```

Options : `--ignore <glob>` (répétable) et `--max-report <n>`. Le script ne lit
pas le Go : il compare les chemins, les contenus et les liens symboliques. Un
manifeste est du JSON trié, donc comparable entre runs et entre machines.

Le contrôle est volontairement strict : un fichier manquant, supplémentaire ou
modifié fait échouer la commande avec le code 1. Les différences sont listées
par type (`missing`, `extra`, `changed`).

## Limite connue au 29 octobre 2026

`node --test tools/driver-cache.test.mjs` échoue avec l'environnement par
défaut : le lanceur `bin/gopurs` active `GOPURS_GO_PACKAGES=1` pour l'hôte
natif, alors que le premier mode du test (JS) produit un layout plat et que le
test compare les octets entre hôtes. Le même test passe avec
`GOPURS_GO_PACKAGES=0` :

```bash
GOPURS_GO_PACKAGES=0 node --test tools/driver-cache.test.mjs
```

Cette limite est indépendante du nettoyage : aucun code de layout, de cache ou
de lanceur n'a été modifié dans les lots 0 et 1. Elle est consignée ici pour ne
pas être confondue avec une régression de déplacement de modules.

## Ce que la baseline ne couvre pas

- Les mesures de performance complètes : elles restent dans le protocole
  [benchmark-protocol.md](../benchmark-protocol.md) et les campagnes
  historiques.
- Les builds reproductibles des binaires natifs sur une autre machine.
- La matrice CI : elle est vérifiée séparément.
