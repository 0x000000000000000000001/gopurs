# A — pack de publication froide

**19 octobre 2026 — activé par défaut après validation.** Le pack est écrit
normalement ; `GOPURS_PACK_PUBLICATION=0` restaure le chemin des objets
individuels.

## Implémentation

- `tools/build-cache/packs.go` ajoute un index privé
  `packs/index.json` et des packs bruts de 64 groupes : `sha256 → {pack, offset,
  bytes}`. Les objets de modules (parties PBO, sources compressées, diagnostics)
  sont regroupés ; les 27 sorties publiques restent des fichiers individuels.
- `store.stage` collecte les nouveaux artefacts et écrit les packs avant de
  committer le manifeste. Un objet déjà présent, individuel ou packé, n'est pas
  réécrit.
- `store.readObject` lit la tranche du pack et vérifie le SHA-256 de l'objet ;
  à défaut, il retombe sur le fichier individuel.
- `freezeModuleCache` et `lookup` utilisent le même `readObject`.
- `retirePackageObjects` élague les entrées et les packs non référencés.
- Un pack absent/corrompu invalide son entrée d'index ; la publication suivante
  le reconstruit.
- `prepare-native-output.mjs`, `build-rust.mjs` et `build-cache.mjs` incluent
  `packs.go`.

## Mesures froides appariées

Même binaire, même corpus, caches privés, trois paires alternées. Les temps sont
le **backend total**.

| Paire | Pack-off | Pack-on | Écart | Publication off/on |
| --- | ---: | ---: | ---: | ---: |
| 1 | 122,419 s | 80,603 s | **−41,816 s (−34,16 %)** | 43,611 / 3,873 s |
| 2 | 130,139 s | 83,155 s | **−46,984 s (−36,10 %)** | 51,693 / 5,517 s |
| 3 | 113,581 s | 85,156 s | **−28,425 s (−25,03 %)** | 41,070 / 5,067 s |
| **Médiane** | **122,419 s** | **83,155 s** | **−41,816 s (−34,16 %)** | **43,611 / 5,067 s** |

La publication médiane baisse de **38,5 s**. Un passage de contrôle du binaire
final après correctif donne **79,914 s**, dans la plage des paires.

Le cache hit pack-on est **1,00–1,60 s** (contre 1,62 s pour la référence P00),
donc la lecture des packs ne pénalise pas la relance identique.

### Édition feuille

Sur le cache froid packé, une édition feuille P00 donne **45,911 s** de backend
(publication 1,928 s, lookup 3,979 s), contre **45,352 s** pour P00
(publication 1,565 s, lookup 3,002 s). Les 27 sorties sont byte-exactes. Le hit
après édition est **1,562 s**.

L'édition ajoute **3 packs** avant compaction. La compaction par groupe réécrit
seulement les groupes qui ont plusieurs packs ou des octets morts : après les
passages `shared` puis `ffi`, on revient à **64 packs / 64 fichiers**,
**427,49 Mio** de packs et un index de ~1,9 Mio. Les 27 sorties `ffi` restent
byte-exactes ; le passage `ffi` réutilise 2 643 modules et en optimise 61.

### Replay b8x

| Hôte | Froid | Hit | Entrées | Sorties exactes |
| --- | ---: | ---: | ---: | ---: |
| JS | 120,680 s | 1,120 s | 10 979 | 27/27 |
| Go natif | 78,448 s | 0,946 s | 10 979 | 27/27 |
| Rust non-PGO | 67,073 s | 0,882 s | 10 979 | 27/27 |

Les logs de hit sont écrits hors du workspace : la racine du projet est une
entrée de répertoire, et un fichier de log ajouté avant le lookup provoque un
`inputs-changed`. Go séquentiel n'a pas été rejoué séparément.

## Intégrité et stockage

- 27/27 sorties publiques byte-exactes face à la référence P00 pour les passages
  conservés (paires 1 et 3, passage final).
- **64 packs**, **10 979 entrées**, **427,5 Mio** de packs, **126,8 Mio**
  d'objets individuels et **1,9 Mio** d'index. Le total est proche du CAS
  antérieur (554,3 Mio), avec environ +1,9 Mio d'index.
- RSS maximal : **10,8–11,0 Gio** pack-off, **11,25–11,67 Gio** pack-on, soit
  environ **+0,4 à +0,8 Gio** sur ces passages. Le temps reste prioritaire ; le
  coût mémoire est documenté.
- La paire 2 a réussi et ses phases sont conservées, mais ses 27 sorties ont été
  écrasées par la paire 3 avant copie. Une passe de reproduction avec sorties
  conservées reste nécessaire pour clore formellement les trois paires.

## Validation

- `go test -race -count=1 -timeout=120s ./...` : passe.
- `node tools/build-cache.test.mjs` : passe, avec un cache réellement packé
  échangé entre le helper JS et le pont Go.
- Rust non-PGO privé : fixture `FFIIntegerReturns`, 60 modules, **301 entrées**
  packées, premier build 2,439 s, hit 0,117 s, exécution `Done`. Binaire
  `gopurs-rust-pack` SHA-256
  `9e965ff2f9e646047f0aef28b88cfe17d237080f2c2a1504898b659c7909f1ad`.
- Tests dédiés : aller-retour pack, source de module gelée, corruption de pack,
  invalidation de l'index et chemin désactivé.
- Échec conservé : le premier build pack-on a révélé une course sur `pending`
  (`concurrent map read and map write`), corrigée par un verrou dans le writer.
- Les sorties publiques, `Wire`, le manifeste et les clés de cache ne changent
  pas.

## Limites

- Le pack est **actif par défaut** ; `GOPURS_PACK_PUBLICATION=0` conserve le
  format individuel.
- JS/Go et Rust ont exercé un cache packé réel. Le replay b8x doit écrire ses
  logs **hors du workspace** : la racine du projet est une entrée de répertoire
  et un nouveau fichier `hit.stderr` invalide le cache avant le lookup.
- La feuille et deux transitions (`shared`, `ffi`) sont mesurées ; un stress de
  nombreuses éditions et de gros changements reste à documenter.
- La compaction par groupe est validée sur deux éditions successives ; son
  comportement sous forte fragmentation reste à éprouver.
- Le gain de 34 % porte sur le backend, pas sur la chaîne complète : le
  `go build` de G06 reste à ~48,5 s après feuille.
- A seul ne promet pas 60 s à froid : après ce gain, le backend reste vers
  **~83 s** sur ces passages.

## Décision

**A est activé par défaut.** Le gain froid dépasse largement le seuil de 20 s
et la publication est bien la cause. JS, Go et Rust ont produit et relu un cache
b8x packé avec 27/27 sorties exactes. `GOPURS_PACK_PUBLICATION=0` reste
l'échappatoire ; Go séquentiel et un stress multi-éditions restent à documenter.

## Preuves

- `htdocs/.gopurs-evidence/a-20261019/logs/` : paires, hit, échec initial.
- `htdocs/.gopurs-evidence/a-20261019/run-pairs.sh` : protocole des paires.
- `docs/build-cost-profile-step0.md` : profil initial.
- Binaire privé :
  `p00-20261019/diagnostic/gopurs/gopurs/bin/gopurs-native-pack`,
  SHA-256 final `8ba4d7c1e577164dce0fab5ab6b40e61284222395b84c5f7db353c3e4f2d6620`.
