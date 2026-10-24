# Backend froid : campagne du 24 octobre 2026

Les changements d'encodage des records et d'évaluation des intrinsèques Array
réduisent le temps du backend froid sur b8x. Trois paires alternées donnent une
médiane des gains appariés de **5,951 s**, sans daemon. Les dernières sondes
supplémentaires restent sous le seuil de 2 s et ne sont pas intégrées.

## Changements conservés

- **Encodage Wire des records (F).** `Gopurs.Cache.Wire` parcourt les champs du
  record complet, sans construire à chaque étape un record amputé par
  `Record.delete`. Le format, l'ordre des champs et de l'internement, le
  décodage et les règles d'invalidation sont conservés.
- **Arguments des intrinsèques Array (G, incluant F).** `Gopurs.CallExprs` et
  `Gopurs.ArrayIntrinsics` capturent les arguments une fois, dans leur ordre
  d'évaluation, avant la boucle, y compris pour un tableau vide. Cela évite
  notamment de reconstruire la fabrique du callback à chaque itération. Le
  fold natif conserve le type de sa valeur initiale ; les arguments de
  surapplication sont évalués après le résultat de l'intrinsèque.

Les adaptations de `driver-cache.test.mjs` corrigent des attentes obsolètes sur
les compteurs de réutilisation et ciblent l'objet faisant réellement autorité
dans le pack pour tester sa corruption. Les oracles de sortie et les scénarios
d'invalidation sont conservés. Ces adaptations ne sont pas des optimisations.

## Protocole et résultats

Le corpus b8x contient **2 704 TAST** et **10 816 parties métier en cache**.
Chaque essai réinitialise le cache backend et le Go généré dans une copie
isolée, à partir des mêmes entrées figées. Le temps inclut la préparation,
l'optimisation, l'émission, les preuves et la publication/fermeture des caches.
Le frontend PureScript, la compilation finale de l'exécutable Go et le calcul
des oracles sont exclus. Le cache de fichiers de l'OS et les caches des outils
Go ne sont pas purgés.

Les mesures natives utilisent un Apple M4 Pro, macOS arm64, Go 1.27.0, 48 GiB de RAM,
`GOGC=off`, `GOMEMLIMIT=16GiB` et huit workers de préparation/PBO/émission.
`GOMAXPROCS` garde la sélection du runtime. Les essais sont séquentiels, sans
autre compilation ou benchmark de cette campagne en concurrence. Chaque froid
est suivi d'une relance sans changement.

A est le compilateur gelé au début de cette recherche structurelle ; G est le
compilateur reconstruit depuis les sources conservées. Les sorties de G
changent intentionnellement avec les nouvelles captures : chaque version a
son propre oracle Go. Les six essais vérifient les Go bruts/regroupés et les
10 816 parties de cache contre leur référence, avec contrôle des entrées,
sources et exécutables.

| Paire | Ordre | A, backend froid | G, backend froid | Gain A − G |
| --- | --- | ---: | ---: | ---: |
| 1 | A puis G | 87,565 s | 79,396 s | 8,169 s |
| 2 | G puis A | 84,518 s | 78,567 s | 5,951 s |
| 3 | A puis G | 84,061 s | 78,859 s | 5,202 s |

La médiane des durées passe de **84,518 à 78,859 s**, soit **5,659 s / 6,7 %**.
La médiane des différences calculées paire par paire est **5,951 s** ; elle
diffère de la différence des deux médianes. La variation observée est conservée
dans le tableau et dans les [six mesures CSV](build-cold-20261024.csv).

Les trois no-op de G mesurent **339 / 336 / 342 ms de backend**, soit une médiane
de **339 ms**. La validation incrémentale distincte mesure le backend à **2,783 s** pour la
première édition, puis **2,125 / 2,129 / 2,131 s**. La génération complète après
la quatrième édition reproduit exactement les sorties incrémentales. Ces
mesures ne démontrent pas un incrémental inférieur à 2 s.

## Validation fonctionnelle

- 39 tests de cache, codec et préparation passent, ainsi que les 4 tests ciblés
  sur les intrinsèques Array et leurs fabriques de callbacks.
- 24 programmes Go issus des snapshots sont compilés et exécutés : 11 avec
  chacun des hôtes JS et natif en monolithe, puis un avec chacun en packages.
  Les fixtures de callbacks et `ArrayRoundtrip` sont aussi exécutées dans les
  deux hôtes. Les snapshots actualisés comprennent des commentaires de
  métadonnées que le générateur produisait déjà avant ces changements.
- Le gate driver-cache passe sur JS et natif parallèle : transitions entre
  hôtes, éditions TAST/FFI, entrées, suppressions/réintroductions de modules,
  corruption des objets engagés, échecs, interruptions, concurrence et
  récupération sont confrontés à des générations fraîches.
- Les quatre éditions b8x et leur oracle complet passent. Le Go de `Test.Main`
  compile ; l'application b8x elle-même n'est pas exécutée dans cette campagne.

Ces contrôles sont ciblés : ils ne remplacent pas une campagne intégrale du
corpus, ni une validation supplémentaire de l'hôte Rust ou du natif séquentiel.

## Pistes examinées sans nouvelle intégration

| Sonde privée | Résultat observé et décision |
| --- | --- |
| Mémoïsation d'expressions du codec | Gain froid de 0,505 s sur la comparaison complète, malgré un microbenchmark favorable. |
| Regroupement des directives, projections Convert, workers natifs, concaténations Wire | Chaque sonde reste sous 2 s. Leur combinaison de code donne un gain médian de 1,3275 s. |
| Autre GC disponible dans le toolchain local | Gain isolé de 1,076 s ; combinaison avec le code précédent : 0,843 s face aux références. Les gains ne s'additionnent pas. |
| Émission à quatre workers ou ordonnanceur limité à huit | Régressions observées ; réglages non retenus. |
| Réduction du travail du checker Go de packaging | Plafond diagnostique de 2,780 s avant le coût des métadonnées nécessaires ; la sonde à corps supprimés ne produit pas un compilateur valide. |
| Concaténations étendues, avec preuve de stabilité des lectures | 583 appels évités dans 40 fichiers. Séquence G/sonde/G : 80,510 / 79,546 / 79,667 s, soit 0,5425 s face à la médiane des références. Non intégrée. |

La première transformation globale des concaténations a été rejetée avant
compilation complète : un aplatissement peut déplacer une lecture de variable
après un appel qui la modifie. La variante mesurée conserve les frontières
ASCII non vides et prouve la stabilité des identifiants locaux. Son oracle
exécute le véritable transformeur, vérifie les cas effectivement transformés
et les refus attendus, les bytes WTF-8 et les traces d'évaluation.

Le profil d'allocations ancien sert au diagnostic des pistes ; le profil
mémoire courant ne couvre que le démarrage et ne quantifie pas les allocations
restantes du compilateur. Les temps CPU cumulatifs se recouvrent et ne sont
pas des secondes murales récupérables que l'on peut additionner.

Cette campagne n'a plus de piste restante étayée à plus de 2 s de gain froid.
Cela n'établit pas une limite théorique du compilateur : une nouvelle hypothèse
structurelle devra d'abord produire une preuve mesurée.

## Provenance et limites

Exécutables natifs gelés, SHA-256 :

```text
A aaf0f971b2e87a68111c51def9ee1e8d474da4c1f0e260cee178171ac9cfb903
G 2c14439288f541c9f6d945fd028cb595cb1bb73a3b91f07d2dad5b9e52560e83
```

Les journaux locaux sont conservés dans les campagnes
`cold-full-20261024` et `cold-structural-20261024` de `.gopurs-evidence`, dont
`final-pairs.json`, `final-pairs-source-state.json`, les rapports de bootstrap,
les oracles des six essais et `release-validation.json`. Les sources compilées
de G correspondent au bootstrap ; seul le test driver-cache a été adapté
ensuite, avec son identité enregistrée séparément. Les sondes privées n'ont
remplacé ni ces sources ni le binaire installé.

La référence officielle historique du
[README d'altbak.pub](https://github.com/0x000000000000000000001/altbak.pub#compilation-times)
reste **33,037 s** pour b8x avec l'hôte Go natif. Son protocole diffère de celui
décrit ici : aucun ratio de vitesse entre ces chiffres n'est pertinent.
Le [bilan du 12 octobre](build-incremental.md) et la référence A antérieure
restent des historiques distincts.

Ces résultats ne mesurent ni le frontend, ni la première compilation Go de
`t -c`, ni l'exécution après « Running Go application ». Aucun gain sur ces
étapes ni passage du build complet sous 5 s n'est établi par cette campagne.
