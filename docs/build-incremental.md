# Bilan du cache incrémental — lots 04 à 06

Le [bilan du 24 octobre](build-cold-20261024.md) décrit les optimisations et
mesures ultérieures. Les chiffres et le comportement ci-dessous restent ceux
de la campagne historique du 12 octobre.

Lots clôturés le **12 octobre 2026**, sur le candidat **`publicsync`**. Le plan v3
compte **7/7 lots validés**. Les mesures du parcours complet confirment la priorité
du découpage des packages Go (suite G), avant la préparation incrémentale (P).
L'historique des essais et leurs validations est conservé dans
[testing.md](testing.md#plan-v3--lot-04--incrémentalité-par-module-validée-12-octobre-2026).

## Résultat livré

Une modification déclenche toujours la préparation globale des métadonnées et
des spécialisations. Le pilote réutilise ensuite les résultats PBO, les workers
et les émissions dont les entrées et les lectures restent valides. Le builder
republie leurs contributions aux barrières canoniques. Les compteurs indiquent
les réutilisations et les causes d'invalidation.

Sur le corpus b8x figé de **2 704 TAST / 3 010 Go** :

| Édition | Modules optimisés / réutilisés | Émissions calculées / restaurées | Go réécrits |
| --- | ---: | ---: | ---: |
| Feuille | 1 / 2 703 | 1 / 2 703 | 1 |
| Partagée | 61 / 2 643 | 6 / 2 698 | 5 |
| FFI | 0 / 2 704 | 1 / 2 703 | 1 |

Le build identique termine avant le décodage/préparation/optimisation/émission et
ne réécrit aucun Go. Un build forcé reproduit les sorties incrémentales à l'octet
près, en conservant aussi les dates des contenus identiques.

## Performances et compromis

Mesures **gopurs seul**, sans profilage, en natif parallèle avec huit workers de
préparation/PBO/émission. Chaque scénario compare trois paires alternées au lot
03 sur les mêmes TAST et FFI. Les campagnes froides partent de caches et sorties
Go réellement neufs ; leurs relances identiques suivent immédiatement.

| Scénario | Lot 03 | `publicsync` | Variation du temps |
| --- | ---: | ---: | ---: |
| Froid backend, cache et Go neufs | 94,191 s | 113,213 s | **+20,19 % / +19,022 s** |
| Relance identique après ce froid | 1,366 s | 1,716 s | **+25,63 % / +0,350 s** |
| Édition feuille | 37,958 s | 34,020 s | **−10,37 %** |
| Édition partagée | 37,832 s | 36,138 s | **−4,48 %** |
| Édition FFI | 37,722 s | 25,364 s | **−32,76 %** |

Les trois paires froides régressent de **18,37 % / 20,19 % / 23,00 %**. Les hits
suivant le froid sont dispersés : l'un des trois passages du lot 03 prend
**4,043 s** ; les valeurs ci-dessus sont les médianes, avec tous les passages
conservés. Le lot 06 remesure l'identique dans la chaîne complète ci-dessous.

Le volume des objets à froid passe de **135 326 356 à 548 673 418 octets**, soit
**+413 347 062 octets**, et de **2 666 à 10 992 objets**. Les quatre parties
intermédiaires des modules représentent **411 148 559 octets**. Les objets
orphelins sont conservés ; leur collecte reste différée.

Face au lot 03, le surcoût froid se concentre dans `optimize + emit`, qui inclut
le calcul des clés et l'encodage du cache : **23,483 → 42,992 s** de médiane.
La préparation reste proche (**9,884 → 9,811 s**) ; la publication durable
baisse légèrement (**56,241 → 54,026 s**). Les compteurs de temps du pipeline
se chevauchent et ne doivent pas être additionnés entre eux.

Les optimisations de stockage réduisent le coût du candidat précédent :
`publicsync` gagne **14,48 % sur la publication isolée** et **5,14 % sur le froid
complet** face à `publicpool2`, au-dessus des seuils préalables de 10 % et 5 %.
Ses budgets de non-régression face à ce prédécesseur passent. Ces gains locaux
ne constituent pas un retour au coût d'avant le cache par module.

**Décision :** après présentation du surcoût froid, l'utilisateur confirme la
poursuite avec ce compromis. `publicsync` clôt les essais d'optimisation du lot
04. Le gain partagé est modeste et le retour à environ 30 s à froid n'est pas
établi. Les temps ci-dessus excluent `go build` ; le lot 06 mesure maintenant
son poids dans la chaîne directe jusqu'à `exec`.

## Parcours complet — lot 06

**24 chaînes validées**, soit trois paires alternées pour identique, feuille,
partagé et FFI. Les [mesures individuelles](build-cycle-lot06.csv) sont publiées
avec ce bilan. Référence : compilateur du lot 03 ; candidat : `publicsync`.
Spago **1.0.3** et le frontend TAST sont figés par hash, en mode hors ligne.
Le backend natif utilise huit workers de préparation/PBO/émission. `tidy`
s'exécute avec Go **1.27.0** sur macOS ; le runner utilise Go **1.26.8 linux/arm64**,
**GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB**, dans une image identifiée par digest.

Chaque générateur a son projet et son journal. Un vrai build remet la base
avant chaque édition. Chaque mesure reçoit un **GOCACHE privé préchauffé sur
la base seulement**, avec le binaire de base du runner ; les variantes éditées
ne sont pas précompilées. Le chronomètre de chaîne va du lancement de Spago au
marqueur précédant `exec Test.Main`. Les oracles sont compilés et exécutés
séparément, hors de ce chronomètre ; la suite métier/PostgreSQL de `Test.Main`
n'entre pas dans ces mesures, comme au lot 00.

| Scénario | gopurs lot 03 | gopurs livré | Variation | Chaîne lot 03 | Chaîne livrée | Variation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| Identique | 1,123 s | 1,559 s | +38,81 % / +0,436 s | 3,559 s | 3,962 s | +11,33 % / +0,403 s |
| Feuille | 37,457 s | 35,030 s | −6,48 % | 247,982 s | 244,403 s | −1,44 % |
| Partagée | 37,244 s | 36,721 s | −1,40 % | 265,895 s | 268,585 s | +1,01 % |
| FFI | 37,969 s | 25,554 s | −32,70 % | 250,273 s | 236,510 s | −5,50 % |

Ce sont les médianes d'invocations complètes, sans addition de médianes de
phases. La dispersion fait partie du résultat :

| Édition | Variation des trois paires gopurs | Variation des trois paires de chaîne |
| --- | --- | --- |
| Feuille | −7,62 / −5,30 / −7,08 % | −2,59 / −1,93 / +1,34 % |
| Partagée | −1,30 / −1,40 / −1,24 % | +1,52 / −3,93 / +1,75 % |
| FFI | −33,45 / −32,04 / −25,07 % | −5,82 / −4,24 / −0,60 % |

Le gain backend est réel sur ces paires, surtout en FFI. La chaîne feuille ne
montre qu'un faible écart observé et la chaîne partagée régresse en médiane ;
ces mesures n'établissent pas un gain robuste du cycle complet sur ces deux
éditions. L'identique conserve zéro Go réécrit et zéro compilation/lien, avec
le surcoût absolu indiqué ci-dessus.

### Où passe le temps après édition

Médianes du candidat livré ; ces colonnes décrivent des horloges distinctes :

| Édition | Frontend/Spago hors backend | `tidy` | `go build` | Commandes de l'action `purescript` | Lien | Pic RSS du compilateur Go |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| Feuille | 1,622 s | 0,399 s | 206,820 s | 202,788 s | 1,133 s | 14,673 GiB |
| Partagée | 21,295 s | 0,495 s | 209,650 s | 205,196 s | 1,152 s | 14,697 GiB |
| FFI | 1,658 s | 0,481 s | 208,470 s | 204,576 s | 1,148 s | 14,732 GiB |

Les actions Go exécutent **deux compilations** (`purescript` et l'entrée) et
**un lien** pour chaque édition, bien que seuls **1 / 5 / 1 Go** changent.
Le RSS est le maximum `VmHWM` observé par un échantillonnage à une seconde ;
`GOMEMLIMIT` reste une cible souple de GC. Aucune contention du verrou n'est
observée. Les oracles du candidat prennent **1,376 / 1,356 / 1,381 s** de médiane,
construction et exécution comprises, hors chaîne.

Go représente environ **85 % / 78 % / 88 %** de la chaîne livrée. Le package
commun `purescript` absorbe presque tout ce temps. La préparation gopurs reste
globale, autour de 9–10 s après édition : ces mesures confirment **G avant P**.
Les seuils déjà fixés pour ces suites restent inchangés ; le lot 06 livre la
référence et les limites, sans leur attribuer des gains encore à obtenir.
Le [relevé G00](build-go-packages.md) prolonge cette référence avec le graphe
du Go réellement généré, les contraintes des bridges et les budgets de ressources.

### Froid, volume et intégrité

La préparation du runner mesure un passage froid par générateur : cache gopurs
et sorties Go absents, GOCACHE vide, modules téléchargés préremplis, frontend/TAST
conservés. Ces **observations uniques**, non appariées, donnent **322,723 →
339,051 s** de chaîne et **95,905 → 114,657 s** de backend ; chacune exécute
**242 compilations et un lien**. Les trois paires froides du lot 04 ci-dessus
restent la référence du surcoût backend ; il ne s'agit pas d'un froid frontend.

Au terme de la séquence d'éditions, les objets conservés occupent **140 784 950
octets** pour le lot 03 et **566 140 963 octets** pour le candidat, objets
orphelins inclus. Ces volumes après plusieurs transitions se distinguent du
volume initial mesuré sur les froids appariés.

Les **24 inventaires** reproduisent les **2 704 TAST / 3 010 Go**, ainsi que
`go.mod` et `go.sum`. Les **24 oracles** passent ; le binaire de `Test.Main` est
byte-exact entre générateurs et répétitions pour chaque scénario. Les audits
finaux confirment les **430 fichiers de production**, les caches Go de base et
les inodes de verrou inchangés. Le même compilateur avait déjà réussi la parité
JS/Go séquentiel/Go parallèle, le bootstrap Rust et les 35 commandes strictes
du lot 05 ; les trois snapshots examinés sont toujours inchangés.

Les deux échecs de préparation et l'interruption d'un runner sont conservés.
La reprise a gardé les 16 échantillons complets et rejoué les huit manquants,
avec provenance par échantillon et mêmes racines de journal. Le passage sans
marqueur `exec` est exclu. Voir l'[historique du lot 06](testing.md#plan-v3--lot-06--mesures-du-parcours-complet-validées-12-octobre-2026).

### Utilisation livrée

Le cache est actif par défaut. Les commandes de build habituel, de calcul forcé
(`--force` ou `GOPURS_FORCE_BUILD=1`) et de réinitialisation sous verrou figurent
dans [usage et maintenance](build-cache.md#usage-quotidien-et-maintenance), avec
les diagnostics de réutilisation/invalidation. Le reset documentaire est exécuté
sur une fixture privée : verrou réellement détenu, sortie symbolique, répétition,
contenus/dates préservés et inode du verrou stable.

## Validation de correction — lot 05

| Couverture | Vérification livrée |
| --- | --- |
| Corps privé/exporté/inliné, directives, types et layouts | 14 états de sources, dont constructeur, record, type de champ et dictionnaire |
| Spécialisations transitives | Ajout, remplacement et retrait dans `Calls → Relay → Poly`, avec TAST des fournisseurs inchangés et préparations rétractées |
| Modules, FFI et entrées | Suppression/ajout, renommage/restauration, priorité/absence FFI et signature Go `int64 → any` consommée depuis un autre module |
| Contexte compilateur/runtime/options | 20 états, versions réelles de compilateur et variantes de runtime reconstruites en JS/Go/Rust |
| Intégrité et propriété | Backend corrompu, sorties manquantes/altérées, recette/manifeste invalides, `go.mod`/`go.sum`, collisions externes |
| Échec, interruption et concurrence | Diagnostics CLI, pilote interrompu, publications concurrentes, erreurs de barrière, SIGKILL réel et reprise, inode du verrou stable |

- **JS / Go séquentiel / Go parallèle** : parité b8x sur les quatre états,
  mêmes TAST, quatre parties, clés et Go ; contrôles incrémental/forcé exacts.
- **Rust** : compilateur privé réellement reconstruit sans PGO et testé.
- **Quatre hôtes** : treize transitions du pilote ; sources **56 comparaisons /
  294 commandes** ; contexte **80 comparaisons / 401 commandes**. Les tests
  sources/contexte compilent et exécutent chaque oracle, comparent au frais et
  vérifient chaque hit sans réécriture.
- **Stockage** : race detectors macOS et Go 1.26.8 linux/arm64 ; reproducteurs
  avant/après, erreurs injectées et interruption réelle conservés.
- **35 commandes de replays stricts**, **19 contrôles CLI**, **430 fichiers
  audités**, trois snapshots examinés et inchangés.

Commandes reproductibles et configuration des variantes :
[tests du cache](build-cache.md#vérifications) et
[campagnes sources/contexte](testing.md#plan-v3--lot-05--correction-incrémentalcomplet-validée-12-octobre-2026).

## Identités et preuves

```text
Bundle JS : 6d3113118958ffa7b762094406fd130a98a5940771f5aafcc25381d535b1f2e6
Natif Go : 841646465c860666e56e37e5064c9db734c31c1570b1707d601286a208739462
Rust privé : 155d915ca48698e7546788be0e79c0c825aff66cca42b7fc5f195ab50ac995c4
```

Le bundle JS utilise le helper de cache compilé depuis les sources auditées ;
son identité d'exécution couvre aussi ce helper et les autres ressources.

Les preuves locales sont conservées dans
`/private/var/folders/w9/l8bnb22d6c75c401f71djbt00000gn/T/opencode/gopurs-module-cache-lot04-nj_1suxg/`.
Leurs fichiers de synthèse ont les SHA-256 suivants :

```text
validation-publicsync-summary.json
  2bef9220b5c0cb8bca029625c0039cd529fbdf6ed4ef4a662e9fa7f8561df06d
validation-publicsync-counts.json
  01884d35bd649e09a43a48038cf3a64655b5005d1ff1a8e1f4ea538bc5b22707
publicsync-validation-audit-after.json
  20d113282ae5d923337b0c3e5633f52ee79bd394b6394bbb50d86cedfee43a74
paired-cold-lot03-publicsync/summary.json
  654a533318f3b75aa1e813063a50816c8d0d48b5907a696e26967254287b3f9e
```

Les candidats retenus/rejetés, les tentatives échouées, les mesures individuelles,
les journaux, les workspaces et les audits sont conservés avec ces synthèses.

Preuves du lot 06, dans les répertoires frères `gopurs-cycle-lot06-publicsync-offline/`,
`gopurs-cycle-lot06-publicsync-resume1/` et `gopurs-cycle-lot06-analysis-final/` :

```text
gopurs-cycle-lot06-publicsync-resume1/summary.json
  d10d53ded77ddc192bc1e403a26a51be513e75b1e077ce06ea47ec75d0ecaef4
gopurs-cycle-lot06-publicsync-resume1/results.json
  d49c9e65707b5da0575dfbfe990f35190f1565c6fc300f0a0c2e031c35c03b64
gopurs-cycle-lot06-publicsync-resume1/protocol.json
  bae33ecb7606fd6fea94e22b5e466d2e98e9a0c229b5039fe1bcd5aba3dcd2ad
gopurs-cycle-lot06-analysis-final/summary.json
  9e3e475e0af2d5503b5b9440d5502f862403ce8ecc362cfa6c9aac0b6b4d871c
```
