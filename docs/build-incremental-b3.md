# B3 — lookup packé et validation parallèle

**20 octobre 2026 — partiel, conservé.** Objectif annoncé : ≥5 s sur le
backend après édition. Résultat : **~0,26 s au hit** et **~2,2 s sur une
feuille** ; le seuil n'est pas atteint. Les clés préparées sélectives ne sont
pas implémentées.

## Changement

- `packDB` garde ouverts les descripteurs des packs pendant un lookup.
- Les artefacts du manifeste sont validés en parallèle avec `boundedEach`.
- Une erreur d'artefact invalide tous les packs de façon conservatrice ; aucun
  pack corrompu n'est réutilisé par une publication ultérieure.
- Le chemin `GOPURS_PACK_PUBLICATION=0` et les sorties publiques sont inchangés.

## Mesures

| Passage | Avant | Après | Gain |
| --- | ---: | ---: | ---: |
| Hit packé, lookup + gel | 0,872–0,914 s | **0,624–0,655 s** | ~0,26 s |
| Feuille `DecodeWithJson` | 45,421 s | **43,210 s** | ~2,2 s |
| Froid | 76,163 s | 76,815 s | bruit |

La feuille est une mesure unique, pas trois paires ; les phases hors lookup
varient aussi. Les 27 sorties restent byte-exactes. Le cold reste dans la plage
du pack A.

## Validation

- `go test -race` passe.
- `node tools/build-cache.test.mjs` passe.
- Le cache corrompu est invalidé puis reconstruit ; le hit normal ne réécrit pas
  l'index.
- Le binaire privé est `gopurs-native-b3`, SHA-256
  `f9940312ad8fd7e1d194978afbde69748d5e8fcc34b12fcdac5f3d2602902e31`.

## Limites

- Le gain est inférieur au seuil B3 de 5 s.
- Le coût des **clés préparées** (~5,3 s) reste entier ; le rendre sélectif
  exige les preuves de lectures P03/P04, pas un simple cache de hash.
- La mesure feuille n'est pas appariée trois fois.
- L'artefact installé n'a pas été reconstruit avec ce correctif ; le binaire
  privé a servi aux mesures.

## Suite

Ne pas retenir B3 comme optimisation autonome suffisante. Les prochains
gisements restent B1 (planification Go globale) et B2 (décodage des 2 703
modules PBO réutilisés). B1 demande un cache de plan par package avec export
data ; B2 demande un descripteur de module réutilisé sans décoder le corps.
