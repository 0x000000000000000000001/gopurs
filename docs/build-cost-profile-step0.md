# Étape 0 — profil des coûts de build

**19 octobre 2026 — terminée.** Profil diagnostique sur un workspace privé P00,
sans modification du code de production. Les temps incluent l'échantillonnage :
ce ne sont pas des mesures d'acceptation.

## Protocole

- 2 704 TAST de base, cache applicatif vide pour le froid, cache Go de base.
- Compilateur natif P00 instrumenté, huit workers préparation/PBO/émission,
  pipeline actif, `GOPURS_GO_PACKAGES=1`, `GOGC=off`, `GOMEMLIMIT=10GiB`.
- Un build froid, puis une édition feuille sur le cache produit, avec des
  fenêtres `sample` de 20 ms (10 ms pour les fenêtres courtes).
- Les fenêtres sont alignées sur les compteurs de phases ; elles peuvent
  chevaucher deux phases. Les parts de piles sont des diagnostics, pas une
  attribution au temps mur.

La première tentative a échoué avant la publication : la copie privée n'avait
pas le répertoire voisin `../gopurs/gopurs-js-promise`, requis par le TAST de
`Promise.Internal`. L'échec est conservé dans `logs/run1-failed/`. Après
restauration de cette dépendance, le build froid et la feuille ont réussi.

## Build froid

| Mesure | Valeur |
| --- | ---: |
| Durée du processus | 123,010 s |
| Backend | 121,055 s |
| RSS maximal | 10,9 Gio |
| Cache lookup + gel | 1,172 s |
| Chargement/tri TAST | 1,237 s |
| Préparation + monomorphisation | 10,846 s |
| Clés des modules | 5,456 s |
| PBO | 42,532 s |
| `optimize + emit` | 48,141 s |
| Exports Go | 0,856 s |
| Planification des packages Go | 11,731 s |
| Publication du cache | 44,630 s |

Compteurs PBO : **2 704 modules optimisés**, tentatives cumulées **92,246 s**
sur huit workers, coordinateur **31,125 s**, attente **8,995 s**, émission
**30,695 s**, génération/écritures cumulées **40,930 s**.

### Profil de publication

La fenêtre de publication contient **97,8 % d'appels filesystem** parmi les
échantillons hors attente :

| Appel | Échantillons |
| --- | ---: |
| `__fcntl` (dont synchronisation) | 4 148 |
| `__rename` | 2 675 |
| `__open` | 2 526 |
| `lstat` | 2 190 |
| `__unlink` | 771 |
| `write` | 507 |

Le JSON et SHA-256 sont marginaux. La publication froide est donc bien dominée
par les **opérations de fichiers durables** (création, synchronisation, rename,
suppression), pas par l'encodage. C'est cohérent avec les **13 840 écritures**
et les **11 005 objets uniques** du manifeste.

### Profil PBO

Les piles sont majoritairement des attentes de goroutines. En retirant les
attentes, les feuilles actives sont surtout :

- runtime/GC et allocation : **79 %** des échantillons actifs ;
- code gopurs généré (interpréteur, conversion, émission) : **18 %** ;
- filesystem : **1,3 %**.

Les piles contiennent `Convert`, `CodeGen` et `Builder`, mais aucun optimiseur
unique n'écrase le profil. Le temps PBO de 42,5 s est donc un mélange de
conversion/émission, d'allocation et de coordination, pas « 42 s de calcul
d'optimisation ». Les compteurs le confirment : les tentatives cumulées ne
représentent que ~92 s sur huit workers, tandis que coordinateur et émission
pèsent ~31 s chacun.

### Profil de planification Go

La fenêtre de planification contient **1 159 frames `go/types`** et **655 frames
`go/ast`**, avec `renderPackages`, `packageGraph` et `go/ast.Walk` visibles.
Les feuilles actives sont réparties entre parcours AST, rendu de packages,
allocations runtime et lectures/écritures de fichiers. Cela confirme que les
11,7 s correspondent au **parsing/type-check/plan complet des 14 fichiers
générés (123 Mio)**, et non à un `go build` externe.

## Édition feuille

| Mesure | Valeur |
| --- | ---: |
| Durée du processus | 52,844 s |
| Backend | 52,151 s |
| RSS maximal | 11,1 Gio |
| Cache lookup + gel | 3,238 s |
| Chargement/tri TAST | 1,303 s |
| Préparation + monomorphisation | 10,980 s |
| Clés des modules | 5,788 s |
| PBO | 13,516 s |
| `optimize + emit` | 19,342 s |
| Exports Go | 0,941 s |
| Planification des packages Go | 11,098 s |
| Publication du cache | 1,805 s |

**2 703 modules restaurés, un seul optimisé.** Les tentatives d'optimisation
prennent **12 ms** ; l'attente PBO vaut **8,169 s**, le coordinateur **1,816 s**
et l'émission **1,254 s**.

Les fenêtres montrent :

- les clés passent par `Wire.write`, `genericEncode`, les codecs de record et
  d'array ; les feuilles actives sont surtout runtime/allocation ;
- la réutilisation passe explicitement par `readPart`, `Wire.read`,
  `decodeExpression` et `decodeBackendSyntax` ; le filesystem ne représente
  qu'environ **2 %** des échantillons actifs. Le coût est donc le décodage JSON +
  allocation, pas la lecture disque ;
- la fenêtre packages/publish contient `go/types` et `go/ast`, puis la
  publication feuille reste courte (1,8 s).

Le profil feuille est plus lent que P00 (**52,2 s contre 45,4 s**) à cause de
l'échantillonnage et de la fenêtre d'observation. Il ne remplace pas les
mesures appariées.

## Classement des candidats

| Candidat | Coût mesuré | Gain plafond raisonnable | Verdict |
| --- | ---: | ---: | --- |
| **A — pack de publication froide** | 44,6 s | **~40 s** (C01 : 82,4 → 3,34 s) | **Priorité 1** |
| B1 — planification Go incrémentale | 11,1–11,7 s | ≥5 s après édition, possible au froid | Après A |
| B2 — réutilisation PBO sans décoder 2 703 modules | 9,9 s (13,5 sous profil) | ~6–8 s après édition | Après A |
| B3 — clés/lookup sélectifs | 9,0 s combinés | ~4–5 s après édition | Après A |
| Préparation | 10,8 s | plafond insuffisant seul | Pas en premier |

**Décision : commencer par A.** C'est le gain froid le plus grand, le plus
directement mesuré et le plus borné : il ne demande pas de prouver une
réutilisation sélective de la préparation. C01 fournit déjà l'ordre de grandeur
du pack, mais A doit rester une **enveloppe de publication** : pas de
restauration globale, les objets `Wire` individuels restent l'oracle.

A seul ne promet pas 60 s : après suppression de ~40 s de publication, le froid
resterait autour de ~75 s. B1/B2/B3 restent nécessaires pour le cycle après
édition et pour la suite du froid.

## Limites

- Une seule mesure froide et une seule feuille, avec échantillonnage ; pas de
  trois paires alternées.
- Les parts de piles sont des **top-of-stack** ; les attentes de goroutines
  dominent et ont été exclues des parts actives.
- Les fenêtres d'échantillonnage chevauchent parfois deux phases ; les compteurs
  de phases restent la référence.
- Le build feuille a été exécuté sur le cache froid produit, avec une édition
  feuille P00 ; ce n'est pas un nouveau corpus.
- Aucune conclusion d'acceptation ni attribution de gain de production n'est
  tirée de ces profils.

## Preuves

- `docs/build-cost-profile-step0.json` : phases, compteurs, RSS, empreintes des
  échantillons et décision.
- `htdocs/.gopurs-evidence/step0-20261019/logs/run2/` : build froid réussi.
- `htdocs/.gopurs-evidence/step0-20261019/logs/run3-leaf/` : édition feuille.
- `htdocs/.gopurs-evidence/step0-20261019/logs/run1-failed/` : échec de
  configuration conservé.
- `docs/build-cost-audit.md` : synthèse P00/G06/C01.
- `docs/build-preparation-p00-timings.csv`,
  `docs/build-go-cycle-g06.csv` et les analyses associées : références.
