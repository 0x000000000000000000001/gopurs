# Gopurs — état et plan actif

**Mise à jour : 19 octobre 2026.** Le plan v3 (lots 00–06) est clos et la suite
G00–G06 est livrée. P00 est terminé et C01 a été rejeté. Aucun candidat
d'optimisation n'est actif. L'historique complet des lots, seuils, échecs et
décisions est conservé dans
[`docs/plan-archive-20261019.md`](docs/plan-archive-20261019.md), copie figée
de l'ancien `todo.md` (SHA-256
`3c28923ef2543d657ceeb5f3525c872aba30e789d72b45d71866d091877a9ccd`).

## Objectif

Réduire le cycle **modifier → construire → tester**. Les repères historiques
restent **~30 s de backend après modification** et **~60 s de backend à froid**.
Ils ne sont pas encore atteints et ne doivent pas être utilisés comme critère
d'acceptation d'un premier lot.

Deux chronomètres restent distincts :

| Périmètre | Feuille | Partagé | FFI | Froid applicatif | Froid complet |
| --- | ---: | ---: | ---: | ---: | ---: |
| Backend P00 | 45,352 s | 47,047 s | 35,056 s | 115,692 s | 118,837 s |
| Chaîne G06 jusqu'à `exec` | 99,772 s | 149,661 s | 120,738 s | 238,438 s | 276,664 s |

Le `go build` seul vaut **48,490 s** après feuille. Une cible de 30 s pour la
chaîne complète exige donc aussi une évolution du découpage Go, pas seulement
du backend.

## État courant

| Élément | État |
| --- | --- |
| Plan v3 (00–06) | **7/7**, clôturé |
| Suite G (G00–G06) | **7/7**, layout packages livré |
| P00 — audit des coûts de préparation | **Terminé**, P à **1/7** |
| C01 — objets compacts / publication regroupée | **Rejeté en l'état** |
| P01–P06 | **Archivés, non financés par les mesures actuelles** |
| Prochain candidat | **A — pack de publication froide**, aucun code engagé |

Références : [bilan P00](docs/build-preparation-baseline.md),
[étude C01](docs/build-preparation-compact.md),
[audit des coûts](docs/build-cost-audit.md),
[profil étape 0](docs/build-cost-profile-step0.md),
[données étape 0](docs/build-cost-profile-step0.json),
[mesures G06](docs/build-go-measurements.md).

### Ce que les mesures ont établi

- **P00** : le backend reste à 45–47 s après édition et 115–119 s à froid. Même
  avec préparation et clés gratuites, il resterait 28,679 / 30,905 s après
  feuille/partagé et 99,612 / 102,562 s aux froids. P seul ne peut pas atteindre
  les repères.
- **C01** : le regroupement publie beaucoup plus vite, mais la restauration
  coûte **11,264 s** (budget 6,5 s), les coûts connus après édition
  **11,784 / 11,964 s** (budget 11 s) et le surcoût froid favorable
  **7,349 / 7,155 s** (budget 5,785 / 5,942 s). Le candidat est rejeté, ses
  preuves sont conservées.
- **Après édition**, quatre coûts fixes totalisent **28,9 s sur 45,4 s** :
  planification Go 10,582 s, restauration PBO 9,930 s, clés préparées 5,432 s,
  lookup/gel 3,002 s. Sur un passage feuille, 2 703 modules sont restaurés et
  un seul est optimisé ; les tentatives d'optimisation prennent ~22 ms.
- **À froid**, publication **46,993 s** et `optimize + emit` **44,458 s**
  représentent environ 78 % du backend. La publication écrit 13 840 objets
  durables pour 550 Mio ; C01 montre qu'un regroupement en 193 groupes peut
  ramener ce type de publication de ~82 s à ~3,3 s.
- **Go packages** coûte ~10 s à chaque build : lecture/parsing/type-check/
  planification des 14 fichiers générés (123 Mio), même si un seul module change.
- **Étape 0** : la publication froide est à **97,8 %** composée d'opérations
  filesystem hors attente (`fsync`, rename, open, lstat) ; le profil PBO est
  dominé par le runtime/GC et l'émission, pas par un optimiseur unique ; la
  réutilisation feuille passe explicitement par `readPart`/`Wire.read` pour
  2 703 modules.

## Décisions

1. **Conserver G06 comme état livré.** Aucune variation face à G06 n'est
   attribuée à P sans campagne appariée.
2. **Ne pas poursuivre P02–P06 tels quels.** Leur ordre supposait que la
   persistance de préparation financerait les repères ; les mesures ne le
   montrent pas. Ils restent archivés, pas supprimés.
3. **Profilage court terminé ; A est prioritaire.** Le pack de publication
   attaque le poste froid le plus mesuré. B1–B3 restent après A ; aucun nouveau
   cache de préparation n'est lancé.
4. **Priorité au temps.** Le RSS est mesuré et documenté, sans plafond
   automatique ; une aggravation démesurée reste un motif d'arrêt.
5. **Un seul candidat à la fois**, avec trois paires alternées, sorties
   byte-exactes et seuil d'arrêt fixé avant l'essai.
6. **Aucun commit/push de l'assistant.** Les commits ordinaires restent au bot
   de l'utilisateur ; les preuves et échecs sont conservés.

## Plan actif

### Étape 0 — profilage court, sans code — **terminée le 19 octobre**

[Rapport](docs/build-cost-profile-step0.md) et
[données vérifiables](docs/build-cost-profile-step0.json). Une session bornée,
sans campagne lourde concurrente ; l'échec de configuration initial est conservé
dans `logs/run1-failed/`.

- [x] **Découper PBO à froid** : optimiseur, coordinateur, attente, émission et
  contre-pression. Les compteurs P00 donnent 39 s de temps mur pour ~10 s de
  tentatives cumulées ; il faut savoir où passe le reste.
- [x] **Découper la planification Go** : lecture/hachage, parsing, type-check,
  planification et écriture des 14 fichiers de 123 Mio.
- [x] **Découper la publication froide** : staging, écriture, `fsync`, rename,
  manifeste et maintenance. Confirmer le coût par objet et le potentiel d'un
  pack.
- [x] **Découper la réutilisation PBO** : 2 703 décodages, coordinateur et
  émission, séparément du calcul d'optimisation.

**Résultat :** froid **121,055 s** avec échantillonnage (publication 44,630 s,
planification Go 11,731 s) ; feuille **52,151 s** (2 703 modules restaurés,
1 optimisé, 5,788 s de clés, 13,516 s de PBO sous profil). La publication est
à **97,8 %** filesystem hors attente ; le profil de réutilisation montre
`readPart` → `Wire.read` → décodage JSON. Le pack de publication est le candidat
au gain plafond le plus élevé (**~40 s** à froid).

**Décision : A d'abord, B après A.** Le profil a dépassé le seuil de **≥5 s** :
A est retenu pour une conception détaillée, sans code avant présentation.

### Chantier A — publication froide regroupée

**Retenu après l'étape 0.** Conception à présenter avant le premier changement :
parcours actuel `BuildCache/store/session`, enveloppe de pack, index, engagement
atomique, reprise et bridges JS/Go/Rust. Pas de restauration globale.

**Constat mesuré :** 13 840 écritures durables, 11 005 objets uniques, 550 Mio,
45–47 s ; 97,8 % des échantillons actifs de publication sont des opérations
filesystem. C01 donne l'ordre de grandeur : 82,4 → 3,34 s en 193 groupes.

- [ ] **Définir l'enveloppe de publication** : pack de sorties/cache, index et
  engagement atomique, sans imposer une restauration globale. Les octets
  individuels `Wire` restent l'oracle ; le pack est un transport de publication.
- [ ] **Préserver l'intégrité et la reprise** : objet absent/corrompu, échec
  avant engagement, index corrompu, reprise après interruption.
- [ ] **Mesurer trois paires froides** (`cold-app`, `cold-complete`) avec les
  mêmes TAST et sorties byte-exactes, puis vérifier la chaleur du cache.
- [ ] **Critères d'arrêt :** gain médian **≥20 s** sur le froid applicatif,
  aucune régression >5 % ailleurs, stockage et récupération acceptables. Si le
  gain net est **<10 s** ou si le pack casse un contrat de reprise, abandonner.
- [ ] **Bridges :** `BuildCache`/store/session Go commun, `Wire` inchangé pour
  les objets individuels, validation JS/Go/Rust avant intégration.

**Note :** ce chantier vise le froid. Il ne doit pas rouvrir la restauration
globale de C01 ni promettre seul les ~60 s.

### Chantier B — incrémental ciblé

**Constat mesuré :** 28,9 s de coûts fixes après édition, dont Go packages
10,6 s, PBO 9,9 s, clés 5,4 s et lookup 3,0 s. Le module modifié ne coûte que
quelques millisecondes.

Après A, choisir **un seul** sous-chantier, dans cet ordre :

- [ ] **B1 — planification Go incrémentale.** Réutiliser le plan/type-check des
  packages dont les sources n'ont pas changé ; mesurer le gain sur feuille et
  vérifier que le froid ne régresse pas.
- [ ] **B2 — réutilisation PBO sans décoder 2 703 modules.** Réutiliser les
  signatures/sorties nécessaires et éviter la restauration complète des modules
  inchangés ; conserver les preuves de dépendances existantes.
- [ ] **B3 — clés et lookup sélectifs.** Éviter de re-hacher 2 704 modules et de
  revalider 550 Mio quand un seul module change, sous preuve complète.

**Critères d'arrêt :** **≥5 s** de gain backend sur le scénario annoncé, aucune
régression >5 % ailleurs, trois paires alternées. Si le gain est **<2 s**,
arrêter le sous-chantier. Un premier candidat accepté doit ramener la feuille
sous **~35 s** ; le repère de 30 s n'est évalué qu'après cumul.

**Bridges :** ne pas rouvrir P02–P06 sans nécessité prouvée. Toute modification
de `PreparedInstantiation`, des clés ou du cache module doit être présentée
avant code, avec les transports JS/Go/Rust concernés.

Le profil feuille montre `readPart`/`Wire.read` dans la réutilisation et
`Wire.write`/hash dans les clés ; B2 et B3 sont donc fondés sur des chemins
observés, pas sur une hypothèse de stockage.

### Chantier C — chaîne complète, conditionnel

**À n'ouvrir que si l'objectif est 30 s de bout en bout.** Le `go build` après
feuille vaut 48,490 s et compile un package généré de 39,8 Mio. Revoir la
densité/le découpage des packages Go, le lien et les actions du runner ; mesurer
séparément backend et Go. Aucun travail n'est engagé dans ce chantier avant A
ou B.

## Budgets et protocole

- **Persistance de préparation** : les budgets P00 restent la référence
  (≤320 Mio courants, ≤640 Mio avec précédent/staging, ≤6,5 s de restauration,
  ≤11 s de parcours complet, froid ≤+5 %). C01 ne les respecte pas.
- **Chantier A** : gain froid ≥20 s ; pas de régression >5 % ; stockage et
  reprise mesurés. Un pack lent à restaurer ou à réparer est un rejet.
- **Chantier B** : ≥5 s par sous-chantier ; feuille sous ~35 s après le premier
  gain accepté ; 30 s seulement après cumul et validation.
- **Mesures** : trois paires alternées minimum, sans campagne lourde
  concurrente ; conserver chaque essai et chaque échec ; comparer les octets
  JS, Go séquentiel/parallèle et Rust sur les mêmes TAST ; examiner les snapshots
  avant mise à jour, puis rejouer strictement.
- **Relance identique** : budget `max(5 %, 0,100 s)`. **Autres scénarios** :
  régression ≤5 %. Le temps prime ; le RSS est un indicateur secondaire.
- **Publication** : les preuves durables restent dans
  `htdocs/.gopurs-evidence/` ; les exports vérifiables dans `docs/`.

## Prochaine action

1. Présenter le code actuel de publication (`BuildCache`, store, session), la
   transformation de pack et les bridges avant le premier changement.
2. Après accord, implémenter **A** seul et mesurer trois paires froides ; arrêter
   si le gain net est <10 s ou si la reprise est cassée.
3. Passer ensuite à **B1–B3**, un sous-chantier à la fois, avec le même seuil de
   5 s et les validations quatre hôtes.
4. Revenir à l'utilisateur avec un bilan chiffré avant toute intégration ou
   réouverture de P.
