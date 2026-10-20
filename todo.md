# Gopurs — état et plan actif

**Mise à jour : 20 octobre 2026.** Le plan v3 (lots 00–06) est clos et la suite
G00–G06 est livrée. P00 est terminé, C01 et B2 sont rejetés. Le chantier C
(architecture incrémentale et découpage Go SCC-aware) est ouvert après
l'exploration de granularité. L'historique complet des lots, seuils, échecs et
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
| A — pack de publication froide | **Activé par défaut**, opt-out `GOPURS_PACK_PUBLICATION=0` |
| Granularité Go / C | **Exploration privée terminée**, non installée |
| C1 — découpage SCC-aware | **Prototype échoué**, refactor codegen requis |
| C1b — refactor SCC codegen | **Échoué** : topo 46,7 s, noinline 25,6 s |
| C1c — bridge ABI/export | **Scalaires + type validés** : FFI 1,02 s |
| C2 — backend résident | **C2-3 mesuré**, 1 module change |

Références : [bilan P00](docs/build-preparation-baseline.md),
[étude C01](docs/build-preparation-compact.md),
[audit des coûts](docs/build-cost-audit.md),
[profil étape 0](docs/build-cost-profile-step0.md),
[données étape 0](docs/build-cost-profile-step0.json),
[pack de publication](docs/build-publication-pack.md),
[B3 lookup](docs/build-incremental-b3.md),
[B2 rejeté](docs/build-incremental-b2.md),
[audit B1](docs/build-incremental-b1-audit.md),
[granularité et architecture](docs/build-granularity-architecture.md),
[conception C1b](docs/build-c1b-package-graph.md),
[conception C2](docs/build-c2-resident.md),
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
- **A (défaut)** : trois paires froides donnent **122,419 → 83,155 s** de backend
  médian, soit **−41,816 s (−34,16 %)** ; la publication médiane passe de
  **43,611 à 5,067 s**. Le replay b8x JS/Go/Rust est exact (27/27) avec des
  hits à **0,88–1,12 s**. Le compilateur installé donne **76,163 s** à froid,
  **0,914 s** au hit et **45,421 s** après `DecodeWithJson`.
  `GOPURS_PACK_PUBLICATION=0` restaure l'ancien chemin.
- **Granularité Go (C0)** : à density 1024 (68 packages), le build Go de
  feuille passe de 30,2 à 9,6 s en `-p 1` et à **2,67 s** en `-p 14` pour le
  module Ping ; le froid Go `-p 14` est 46,7 s. Le `planHash` est identique
  base/feuille. Mais FFI reste à **32,9 s** et son isolation crée un cycle
  d'imports ; shared reste à 31,2 s. Le backend reste à 43,7 s. Voir
  [granularité](docs/build-granularity-architecture.md).

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

**Décision : A d'abord, B après A.** Le profil a dépassé le seuil de **≥5 s**.
A est activé par défaut après validation JS/Go/Rust ; la compaction et la
validation de base passent. B peut commencer après le stress multi-éditions.

### Chantier A — publication froide regroupée

**Activé par défaut ; `GOPURS_PACK_PUBLICATION=0` conserve l'ancien chemin.**
Résultats, code et limites : [pack de publication](docs/build-publication-pack.md).

**Constat mesuré :** 13 840 écritures durables, 11 005 objets uniques, 550 Mio,
45–47 s ; 97,8 % des échantillons actifs de publication sont des opérations
filesystem.

- [x] **Définir l'enveloppe de publication** : 64 packs bruts, index privé,
  engagement atomique, sorties publiques individuelles, sans restauration
  globale. Les octets `Wire` restent l'oracle.
- [x] **Préserver l'intégrité et la reprise** : pack absent/corrompu, index
  corrompu, invalidation et reconstruction sont couverts par les tests.
- [x] **Mesurer trois paires froides** : backend **122,419 → 83,155 s** en
  médiane (**−41,816 s, −34,16 %**) ; publication **43,611 → 5,067 s**. Le hit
  pack-on reste à **1,00–1,60 s**. 64 packs, 10 979 entrées, ~427,5 Mio.
- [x] **Édition feuille** : backend **45,911 s** contre **45,352 s** pour P00,
  publication 1,928 s, hit 1,562 s, 27/27 sorties exactes. L'édition ajoute
  3 packs avant compaction.
- [x] **Compaction par groupe** : après `shared` puis `ffi`, retour à
  **64 packs / 64 fichiers**, **427,49 Mio**, sorties `ffi` exactes ; 2 643
  modules réutilisés, 61 optimisés. Un stress multi-éditions reste à faire.
- [x] **Critères d'arrêt :** gain froid ≥20 s **atteint** ; feuille dans le
  bruit, compaction validée, JS/Go/Rust exacts. Un stress multi-éditions et les
  autres scénarios restent à documenter. RSS pack-on +0,4 à +0,8 Gio.
- [x] **Bridges :** tests JS/Go packés, fixture Rust (301 entrées, hit
  0,117 s, `Done`) et replay b8x JS/Go/Rust (27/27, hits 0,88–1,12 s).
  Go séquentiel reste non rejoué séparément ; le Rust installé n'a pas été
  reconstruit avec le pack, seul le binaire privé l'a été.

**Note :** ce chantier vise le froid. Il ne doit pas rouvrir la restauration
globale de C01 ni promettre seul les ~60 s. Le `go build` de la chaîne complète
reste à ~48,5 s après feuille.

### Chantier B — incrémental ciblé

**Constat mesuré :** 28,9 s de coûts fixes après édition, dont Go packages
10,6 s, PBO 9,9 s, clés 5,4 s et lookup 3,0 s. Le module modifié ne coûte que
quelques millisecondes.

Après A, choisir **un seul** sous-chantier, dans cet ordre :

- [~] **B1 — planification Go incrémentale.** Audit terminé : sur 11,57 s,
  parsing 1,28, type-check 4,22, graphe 2,34, plan 0,42, rendu/reste ~2,9.
  Le type-check seul ne suffit pas ; un cache de plan par package avec export
  data est nécessaire. Aucun prototype ni changement de production. Voir
  [audit B1](docs/build-incremental-b1-audit.md).
- [~] **B2 — réutilisation PBO sans décoder 2 703 modules.** Le prototype
  réduit le PBO de **9,9 à 7,9 s**, mais le backend feuille reste à **45,8 s**
  (pas de gain total) ; la transition FFI a exigé un digest émetteur et a été
  corrigée. Candidat **rejeté et retiré** du chemin installé. Voir
  [B2](docs/build-incremental-b2.md).
- [~] **B3 — clés et lookup sélectifs.** Le lookup packé avec descripteurs
  ouverts et validation parallèle donne **0,914 → 0,655 s** au hit et
  **45,421 → 43,210 s** sur une feuille, pas les 5 s visés. Les clés préparées
  (~5,3 s) restent entières et exigent les preuves P03/P04. Voir
  [B3](docs/build-incremental-b3.md).

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

### Chantier C — architecture incrémentale et Go SCC-aware

**Ouvert le 20 octobre 2026 après l'exploration C0.** Le `go build` n'est plus
le seul mur : la granularité peut descendre à 1–3 s pour un module séparable,
mais le backend reste à 43,7 s et les modules dans un SCC/cycle gardent un
plancher Go d'environ 33 s. Aucune densité statique ni isolation naïve ne donne
un build court universel.

- [x] **C0 — exploration de granularité (privée).** Densités 512–8192,
  `-p 14`, isolation dynamique, multi-édition. Résultats et limites dans
  [granularité](docs/build-granularity-architecture.md). Aucun changement de
  production.
- [~] **C1 — découpage Go SCC-aware.** Prototype terminé, **rejeté** : Ping
  isole 3 déclarations et tombe à 1,16 s, mais FFI a un SCC de 57 composants /
  4 447 déclarations ; même la fermeture de dépendances (180 déclarations)
  recrée un cycle package. La cible FFI ≤15 s n'est pas atteinte. Voir
  [granularité](docs/build-granularity-architecture.md).
- [~] **C1b — refactor codegen SCC.** Trois variantes testées : isolation SCC,
  partition topologique stable et `-gcflags=all=-l`. Résultats : FFI
  **46,7 s** en topo, **25,6 s** sans inlining ; la cible ≤15 s n’est pas
  atteinte. Ping reste à 1,16 s. **C1b rejeté**, le plancher FFI reste
  structurel. Voir [C1b](docs/build-c1b-package-graph.md).
- [x] **C1c — bridge ABI/export (preuve de concept).** Trois fonctions
  scalaires et le type `IntervalId` validés : `p_types`, champs exportés,
  alias public, implémentation séparée. Go base 52,96 s, **feuille 1,02 s**,
  `SetInterval`/`ClearInterval` OK. Reste : automatisation de l’extraction,
  init/valeurs globales, autres types. Voir
  [C1b/C1c](docs/build-c1b-package-graph.md).
- [~] **C2 — backend résident/incrémental.** C2-1/C2-2 codés. C2-3 mesuré :
  base/feuille **2 703 modules préparés identiques**, seul
  `Inter.Cli.Ping.Main` change. Le cache par module est viable ; il reste à
  écrire `prepareModulesDelta`. Voir [C2](docs/build-c2-resident.md). Cible :
  backend feuille ≤20–25 s, puis ≤15 s avec B1/B3. Fallback complet au moindre
  doute.
- [ ] **C3 — froid, conditionnel.** Le froid actuel (backend ~77 s + Go 46,7 s)
  ne peut pas viser 60 s par un cache chaud. À ouvrir après C1/C2 avec un budget
  séparé, pas comme sous-produit.

**Critères d'arrêt C :** trois paires alternées, sorties exactes, aucune
régression froide >5 %. C1/C1b sont arrêtés : FFI reste 25,6–46,7 s. C2 ne
doit pas être lancé seul pour promettre un build universel ; il ne réduit que
le backend.

**Dépendances :** B1 (plan Go), B3 (clés) et B2 (descripteur PBO) deviennent des
composants de C2, pas des candidats autonomes. B2 reste rejeté dans l'ancien
chemin ; B3 reste partiel et exige les preuves P03/P04.

## Budgets et protocole

- **Persistance de préparation** : les budgets P00 restent la référence
  (≤320 Mio courants, ≤640 Mio avec précédent/staging, ≤6,5 s de restauration,
  ≤11 s de parcours complet, froid ≤+5 %). C01 ne les respecte pas.
- **Chantier A** : gain froid ≥20 s ; pas de régression >5 % ; stockage et
  reprise mesurés. Un pack lent à restaurer ou à réparer est un rejet.
- **Chantier B** : ≥5 s par sous-chantier ; feuille sous ~35 s après le premier
  gain accepté ; 30 s seulement après cumul et validation.
- **Chantier C** : C1b est rejeté ; C1c doit viser un Go feuille ≤5 s pour les
  fonctions à corps modifié, avec fallback si type/signature change. C2 backend
  ≤20–25 s puis ≤15 s ; aucune régression froide >5 %.
- **Mesures** : trois paires alternées minimum, sans campagne lourde
  concurrente ; conserver chaque essai et chaque échec ; comparer les octets
  JS, Go séquentiel/parallèle et Rust sur les mêmes TAST ; examiner les snapshots
  avant mise à jour, puis rejouer strictement.
- **Relance identique** : budget `max(5 %, 0,100 s)`. **Autres scénarios** :
  régression ≤5 %. Le temps prime ; le RSS est un indicateur secondaire.
- **Publication** : les preuves durables restent dans
  `htdocs/.gopurs-evidence/` ; les exports vérifiables dans `docs/`.

## Worktrees de travail

- **gopurs** : branche `c1b` créée depuis `edge` (HEAD `944f0e9`), sans
  worktree séparé. Les changements B3/docs non commités restent dans la
  worktree `gopurs/gopurs` sur `edge`.
- **b8x** : worktree `/Users/0x1/Documents/htdocs/b8x-c1b`, branche `c1b`,
  depuis `master` `67b6799f`. La worktree principale `b8x` garde les
  changements utilisateur et n'est pas utilisée par les essais.
- **purescript** : worktree `/Users/0x1/Documents/htdocs/purescript-c1b`,
  branche `c1b`, depuis `edge` `105f575`.
- **PBO** : worktree existant
  `/Users/0x1/Documents/htdocs/purescript-backend-optimizer-gopurs`,
  branche `edge-gopurs`.
- **altbak.pub** : worktree existant
  `/Users/0x1/Documents/htdocs/altbak.pub-gopurs`, branche `edge-gopurs`.

## Prochaine action

1. **A est clos** : pack par défaut, froid installé 76,163 s, hit 0,914 s,
   feuille 45,421 s, sorties exactes.
2. **C0 est terminé** : meilleur cas Go 2,67 s (Ping, density 1024, `-p 14`),
   mais FFI à 32,9 s, isolation FFI impossible (cycle), backend 43,7 s.
3. **C1c est validé** sur scalaires et type (`IntervalId`, Go feuille 1,02 s).
   La prochaine priorité est **C2** : backend résident, cible ≤25 s. C1c doit
   encore être automatisé pour les types/init, mais Go n’est plus le mur.
4. **B1/B2/B3 sont des composants de C2** ; B2 reste rejeté dans l'ancien
   chemin, B3 reste partiel et attend P03/P04.
5. Garder Go séquentiel, le stress multi-éditions et les trois jeux de sorties
   froides comme contrôles de régression du pack.
