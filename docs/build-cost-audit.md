# Audit des coûts de build — 19 octobre 2026

**Portée :** synthèse en lecture seule des campagnes P00, G06 et C01. Cette
synthèse n'ajoute pas de mesure et ne change pas le code. Les temps backend
viennent de P00 ; la chaîne complète vient de G06. Les deux périmètres ne se
substituent pas.

**Profil complémentaire :** l'[étape 0](build-cost-profile-step0.md) découpe les
coûts non attribués sur un build froid privé et une édition feuille. Elle
confirme le poids des opérations filesystem dans la publication, du
parsing/type-check Go et du décodage `Wire.read` à la réutilisation.

**Suite A :** le [pack de publication](build-publication-pack.md), désormais par
défaut, donne **122,4 → 83,2 s** de backend médian sur trois paires froides
(−34,2 %) ; la feuille reste dans le bruit, la compaction ramène à 64 packs et
JS/Go/Rust rejouent b8x avec 27/27 sorties exactes. Un stress multi-éditions
reste à documenter.

**B3 :** le lookup packé parallèle donne **0,914 → 0,655 s** au hit et
**45,4 → 43,2 s** sur une feuille, sous le seuil de 5 s. Voir
[B3](build-incremental-b3.md).

**B2 :** le descripteur PBO sans corps réduit le PBO de **9,9 à 7,9 s**, mais
le backend feuille reste à **45,8 s** ; candidat rejeté et retiré. Voir
[B2](build-incremental-b2.md).

**B1 :** audit du planificateur Go : **11,57 s** au total, dont type-check 4,22
et graphe/rendu ~5,2. Un cache par package avec export data est nécessaire ;
aucun prototype. Voir [audit B1](build-incremental-b1-audit.md).

**C0 — granularité Go :** à density 1024 (68 packages), le build feuille passe
de 30,2 à 9,6 s en `-p 1` et à **2,67 s** en `-p 14` pour Ping ; froid Go
46,7 s. Le `planHash` est identique base/feuille. Mais FFI reste à 32,9 s,
son isolation crée un cycle, et le backend reste à 43,7 s. Voir
[granularité et architecture](build-granularity-architecture.md).

**C1 — isolation SCC :** Ping (SCC de 3 déclarations) passe à **1,16 s**.
FFI a un SCC de **57 composants / 4 447 déclarations** ; l’isolation du SCC
et la fermeture de dépendances (180 déclarations) recréent un cycle package.
C1 est rejeté ; C1b (refactor codegen) est requis pour un plancher universel.
La première tentative C1b (extraction de FFI + `Util.Html.Clean.Clean`) laisse
encore un cycle via plusieurs packages. La fermeture globale FFI fait
**10,9 Mo / 2 510 déclarations**, presque le package complet. L’analyse finale
montre un SCC package de **49 packages / 92,2 Mo** pour FFI (Ping : 1 package /
2 Ko). C1b n’est pas borné pour ce programme ; il exigerait une refonte de
presque tout le graphe.

## Deux chronomètres à ne pas confondre

| Périmètre | Feuille | Partagé | FFI | Froid applicatif | Froid complet |
| --- | ---: | ---: | ---: | ---: | ---: |
| Backend P00 | 45,352 s | 47,047 s | 35,056 s | 115,692 s | 118,837 s |
| Chaîne G06 jusqu'à `exec` | 99,772 s | 149,661 s | 120,738 s | 238,438 s | 276,664 s |

Les repères **~30 s après édition / ~60 s à froid** sont des repères backend.
La chaîne complète est environ deux fois plus longue. Si l'objectif est la chaîne
complète, le seul `go build` après feuille vaut déjà **48,490 s** : il constitue
un plancher indépendant du backend.

## Où passe le backend après édition

Médianes P00 pour la feuille, décomposées sans compter deux fois les phases
imbriquées. Les clés et PBO sont dans `optimize + emit`.

| Poste | Secondes | Part | Contenu observé |
| --- | ---: | ---: | --- |
| Planification des packages Go | 10,582 | 23,3 % | Lecture des 14 fichiers générés (123 Mio), parsing/type-check/plan complet |
| PBO et émission | 9,930 | 21,9 % | 1 module optimisé, 2 703 restaurés ; très peu de calcul d'optimisation |
| Préparation + monomorphisation | 9,905 | 21,8 % | Demandes directes 2,054 ; transitif 6,769 ; réécriture 0,597 |
| Clés des modules préparés | 5,432 | 12,0 % | 2 704 modules hachés ; `computed=2704, reused=0` |
| Lookup et gel du cache | 3,002 | 6,6 % | Validation des sorties/objets précédents et copie des TAST |
| Résidu backend | 2,731 | 6,0 % | Non attribué par les compteurs |
| Publication du cache | 1,565 | 3,5 % | Réutilisation des objets précédents |
| Chargement/tri TAST | 1,210 | 2,7 % | 2 704 modules |
| Exports Go | 0,631 | 1,4 % | Acquisition des exports externes |
| Enveloppe processus | 0,746 | 1,6 % | Lancement/fermeture hors backend |

Le détail PBO d'un passage feuille est frappant : **22 ms** de tentatives
d'optimisation, **6,1 s** d'attente, **1,4 s** de coordinateur et **1,0 s**
d'émission, pour **2 703 modules restaurés** et un seul émis. Le temps après
édition ne mesure donc presque pas le travail d'optimisation du module modifié.

## Où passe le backend à froid

| Poste | Froid applicatif | Part | Froid complet | Part |
| --- | ---: | ---: | ---: | ---: |
| Publication du cache | 46,993 s | 40,6 % | 45,167 s | 38,0 % |
| PBO (hors clés) | 38,956 s | 33,7 % | 41,041 s | 34,5 % |
| Planification des packages Go | 10,001 s | 8,6 % | 10,364 s | 8,7 % |
| Préparation + monomorphisation | 9,727 s | 8,4 % | 9,859 s | 8,3 % |
| Clés des modules | 5,288 s | 4,6 % | 5,432 s | 4,6 % |
| Lookup et gel | 0,878 s | 0,8 % | 1,005 s | 0,8 % |
| Chargement/tri TAST | 1,155 s | 1,0 % | 1,201 s | 1,0 % |
| Exports Go | 0,637 s | 0,6 % | 3,104 s | 2,6 % |
| Résidu + enveloppe | 2,414 s | 2,1 % | 2,511 s | 2,1 % |

`optimize + emit`, clés comprises, vaut **44,458 / 46,640 s**. La publication
et cet ensemble représentent à eux deux **environ 78 %** du froid. La
préparation, cible principale de P, en représente moins de 9 %.

## Quatre causes mesurées

### 1. La publication froide écrit des milliers de petits objets durables

Le manifeste P00 contient **27 sorties**, **10 816 parties de modules** et
**2 997 sources de modules**, soit **13 840 écritures** pour **11 005 objets
uniques** et **550 Mio**. Le stage lit, hache, écrit chaque objet dans un fichier
temporaire, synchronise, renomme, puis synchronise les répertoires. À
**45–47 s**, cela correspond à environ **3,2 ms par écriture** : le coût est
d'abord celui du nombre d'objets et des barrières de durabilité, pas celui des
octets.

C01 fournit un ordre de grandeur : le même type de jeu publié en **193 groupes
compressés** prend **3,34 s** (complet) et **2,89 s** (compact), contre
**82,41 s** et **90,97 s** en objets individuels. C'est la preuve qu'un
regroupement de publication peut supprimer l'essentiel de ce poste. C01 ne
prouve pas qu'une intégration de production serait gratuite, et sa restauration
globale était trop chère ; le regroupement devrait rester **une enveloppe de
publication**, pas une obligation de tout restaurer.

### 2. PBO/émission à froid n'est pas « 39 s d'optimisation »

Les compteurs d'un passage froid donnent, pour **2 704 modules** :

- tentatives cumulées **81,8 s** sur huit workers, soit environ **10,2 s** de
  temps mur d'optimisation ;
- coordinateur **28,9 s** ;
- attente **7,8 s** ;
- émission **28,5 s** ;
- génération/écritures cumulées **37,5 s**.

Ces compteurs se recouvrent et ne s'additionnent pas. Ils montrent surtout que
le temps mur de **39 s** n'est pas passé dans le seul optimiseur. Une partie
importante vient de l'ordonnancement, de l'émission et de la contre-pression.
Il faut un profil CPU avant d'attribuer ce poste à l'optimisation elle-même.

### 3. Après édition, le cache paie un coût fixe proche de 28 s

Sur une feuille, les quatre postes suivants totalisent **28,9 s sur 45,4 s** :

- planification complète des packages Go : **10,6 s** ;
- restauration/décodage des 2 703 modules réutilisés : **9,9 s** ;
- hachage des 2 704 modules préparés : **5,4 s** ;
- validation du manifeste et des 550 Mio d'objets précédents : **3,0 s**.

Le module effectivement modifié ne coûte que quelques millisecondes à optimiser.
C'est la raison pour laquelle une édition d'une ligne conserve un temps proche
d'un build complet. Réduire uniquement la préparation ne déplacera donc pas
beaucoup ce plancher.

### 4. La préparation n'est pas le mur principal

La préparation et sa monomorphisation coûtent **9,7–10,1 s** après édition comme
à froid. Même en la rendant entièrement gratuite, le backend P00 laisserait
**28,7 s** sur feuille et **99,6 s** à froid applicatif. Elle ne peut donc pas,
à elle seule, atteindre les repères de 30 s et 60 s.

## Ce que cela implique pour les objectifs

- **30 s backend après édition :** il manque environ **15–17 s**. Le gisement
  principal est la planification Go, la restauration PBO et les clés/lookup
  (28,9 s de coûts fixes), pas la préparation seule. Un ou deux de ces postes
  doivent être supprimés en grande partie, ce qui demande une modification
  structurelle.
- **60 s backend à froid :** il manque **55–59 s**. Le regroupement de
  publication peut plausiblement en retirer **~40 s**, mais PBO/émission reste
  à **39–46 s**, la planification Go à **10 s** et la préparation à **10 s**.
  Un seul changement ne suffit pas ; même le regroupement laisserait environ
  **75 s**.
- **30 s de chaîne complète :** le `go build` après feuille vaut déjà **48,5 s**
  et compile une package générée de **39,8 Mio**. Il faut soit modifier le
  découpage/densité des packages Go, soit changer de stratégie de compilation ;
  le backend seul ne peut pas atteindre cet objectif.
- **RSS :** le pic reste de l'ordre de **10–11 Gio** avec `GOGC=off`. Ce n'est
  pas la contrainte prioritaire, mais toute nouvelle structure doit être
  chiffrée avant intégration.

## Prochaine instrumentation utile, avant tout nouveau candidat

1. **Profil CPU du froid PBO/émission** pour séparer optimiseur, coordinateur,
   émission et attente. Les compteurs actuels ne suffisent pas à attribuer les
   39 s.
2. **Découpage interne de la planification Go** : lecture/hachage, parsing,
   type-check, planification et écriture, sur les 14 fichiers de 123 Mio.
3. **Découpage de la publication froide** : stage, écriture, `fsync`, rename et
   manifeste, pour confirmer le coût par objet et le gain d'un pack.
4. **Coût de la restauration PBO réutilisée** : mesurer separately les 2 703
   décodages et le travail de coordinateur, plutôt que de les laisser dans un
   seul compteur `PBO producer`.

Ces mesures peuvent être courtes et read-only. Elles ne nécessitent pas une
nouvelle campagne de préparation ni un nouveau candidat de cache.

## Preuves

- `docs/build-preparation-p00-timings.csv` et
  `.gopurs-evidence/p00-20261019/analysis/timings.json` : phases P00.
- `.gopurs-evidence/p00-20261019/measurement/logs/` : compteurs PBO, cache et
  publications par passage.
- `docs/build-go-cycle-g06.csv` et `docs/build-go-measurements.md` : chaîne G06.
- `.gopurs-evidence/p01-compact-20261019-c01/packs/results.json` : publication
  raw/regroupée C01.
- `.gopurs-evidence/p00-20261019/tree/b8x/output/.gopurs-cache/` : manifeste
  de 13 840 écritures et 550 Mio d'objets.
