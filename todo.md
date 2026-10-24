# Gopurs — état et plan actif

**Mise à jour : 24 octobre 2026.** Le plan v3 (lots 00–06) est clos et la suite
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
Ils restent des repères généraux. C2-9 confirme moins de 4 s backend avec D
sur le cas central borné décrit ci-dessous. Cela ne clôt pas l'objectif à froid
ni celui de la chaîne complète. C2-10 est clos pour impasse pratique : trois
paires L exactes à 2,347 s médian, quatorze gates qualité verts et seul le
seuil strict de 2 s en échec. Ce seuil reste non atteint.

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
| C1c — bridge ABI/export | **Intégré** dans le cache Go ; overlay par défaut via le lanceur natif |
| C2 — backend résident | Mesures historiques ; chemin corrigé et daemon conservés |
| C2-5 — Go incrémental | **Overlay 3.5 branché**, Go packages 0,004–0,021 s |
| C2-6 — daemon one-shot | Vérifié byte-exact avec 1 et 8 workers, y compris après erreur |
| C2-7 — one-shot persistant | **16,970 s médiane** sur édition littérale centrale ; 3 paires byte-exactes, cible ≤20 s atteinte |
| C2-8 — one-shot central <10 s | **Achevé**, v4 : 8,586 s backend médian, trois paires byte-exactes ; vrai `b` 13,651 s total ; audit 25/25 |
| C2-9 — one-shot central <4 s | **Achevé**, D : 3,841 s backend médian, trois paires byte-exactes ; vrai `b` 8,791 s total ; cache disque, sans daemon |
| C2-10 — cible centrale <2 s | **Clos pour impasse pratique**, L : 2,347 s médiane, trois paires exactes ; vrai `b` 6,836 s total ; cible non atteinte |

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
[plan C2-5 Go](docs/build-c2-go-packages.md),
[build one-shot persistant](docs/build-c2-persistent.md),
[backend central sous 10 s](docs/build-c2-sub10.md),
[backend central sous 4 s](docs/build-c2-sub4.md),
[poursuite vers 2 s et limites](docs/build-c2-sub2.md),
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
  scalaires et le type `IntervalId` validés manuellement. Composants
  automatiques dans `tools/c1c` : manifeste ABI, `OverlaySource`,
  qualification par table symbole→paquet, réécriture des références, test
  end-to-end `go run`. **Branchés opt-in** dans le pipeline Go.
  Voir [C2-5](docs/build-c2-go-packages.md).
- [x] **C2-3 — delta de préparation historique.** Base 106,652 s,
  feuille **47,372 s**, préparation sautée ; **11 136 fichiers identiques** au
  build one-shot forcé sur le cas testé. Le skip source/interface de cette
  version est remplacé par la preuve de préparation C2-7.
- [x] **C2-4 — clés/PBO incrémentaux historiques.** Delta monolithique : préparation
  **1,807 s**, clés **0,731 s**, PBO **4,899 s**, backend feuille **14,734 s**,
  one-shot forcé 60,682 s, **11 136 fichiers identiques**.
- [~] **C2-5 — Go packages.** Layout packages activé puis overlay C1c branché
  derrière `GOPURS_GO_PACKAGE_OVERLAY=1` : `Go packages` passe de **11,5 s** à
  **0,003–0,021 s** ; backend overlay **17,323–21,515 s** (médiane 19,773 s)
  sur les éditions Ping, sorties identiques à l'oracle full, `go build`
  0,8–2,1 s. Cache de plan public, signatures réutilisées, fallback complet sur
  type/méthode/init/FFI/import/référence. **Flag opt-in ; pas de type-check
  overlay avant publication.**
  Voir [C2-5](docs/build-c2-go-packages.md).
- [~] **C2-6 — daemon one-shot (sans watch).** Prototype `GOPURS_DAEMON=1`
  dans `bin/gopurs` : démarrage automatique, jeton/statut fichier, redémarrage
  sur changement de binaire/options. Mesures historiques : premier build
  91,9 s, puis **24,9 / 21,2 /
  20,7 s** backend Ping, hit no-op 1,15 s, frontend no-op 2,05 s. Non validé :
  concurrence, `b -c`, Rust.
  Voir [C2-6](docs/build-c2-resident.md).
- [x] **C2-7 — one-shot persistant, cas central borné.** Préparation compacte
  prouvée, contributions PBO à la demande, clés exactes ou fermeture des lectures,
  overlay conservant les cinq bundles de `DecodeJsonWith`. Trois paires natives
  finales : **16,970 / 17,293 / 16,946 s**, médiane **16,970 s** contre
  **92,934 s** full. Les 2 704 modules et 27 sorties de packages sont
  byte-exacts pour chaque paire. Frontend et Go exclus ; cache amorcé.
  Le test porte sur une édition littérale de corps : changement structurel,
  de type ou de spécialisation hors preuve revient à la préparation complète.
  Packs corrompus rejetés puis réparés lors du full ; suites ciblées et Go
  `-race` vertes. Daemon séquentiel vérifié avec 1 et 8 workers, `Int → String
  → Int`, édition littérale et échec récupérable. Activé par défaut dans le
  lanceur Go natif pour permettre `b` sans exports ; autres hôtes opt-in.
  Voir [C2-7](docs/build-c2-persistent.md).
- [x] **C2-8 — backend central sous 10 s.** Trois paires natives v4 :
  **8,586 / 8,552 / 8,603 s**, médiane **8,586 s**, contre **79,464 s** full
  médian. Les 2 704 modules et 27 sorties sont byte-exacts pour chaque paire.
  Chaque incrémental respecte le seuil strict ; frontend et compilation Go
  exclus, cache amorcé, édition littérale de corps uniquement. Suites ciblées
  27/27 et Go `-race` vertes. La campagne v3 antérieure et son `b` restent
  historiques : une corruption avait ensuite révélé un panic Aff empêchant
  le fallback. V4 corrige la propagation des exceptions et
  la libération des brackets ; deux corruptions natives et les daemons à un/huit
  workers sont verts. Compilation Go finale et vrai `b` v4 réussis : **8,977 s
  backend**, **13,651 s au total** pour ce dernier, sans daemon. Audit de clôture
  **25/25**, SHA et sources concordants ; C2-8 achevé sur ce périmètre.
  Voir [C2-8](docs/build-c2-sub10.md).
- [x] **C2-9 — backend central sous 4 s.** Trois paires natives D :
  **3,872 / 3,841 / 3,817 s**, médiane **3,841 s**, contre **78,111 s** full
  médian. Les 2 704 modules et 27 sorties sont byte-exacts pour chaque paire.
  Préparation à la demande, projection typée fragmentée, gel privé groupé,
  contributions PBO/workers paresseuses et scan d'overlay réduit sous preuve.
  Go final compilé, 54 tests ciblés et Go `-race` verts ; replis sur parties
  corrompues et projection absente/corrompue validés, daemons 1/8 conservés.
  Le vrai `b` utilise les défauts : **3,826 s backend**, **8,791 s total**,
  sans daemon. Le seuil porte sur une édition littérale centrale avec cache
  amorcé ; frontend et compilation Go restent séparés. Le contrôle projection
  conserve les différences de disposition des packages et prouve les valeurs
  publiques à l'exécution en plus de l'égalité des sources Go brutes.
  Voir [C2-9 et son audit](docs/build-c2-sub4.md).
- [x] **C2-10 — exploration close pour impasse, cible de 2 s non atteinte.** L est lié à
  sa capture de sources. Trois paires finales : **2,335 / 2,347 / 2,401 s**,
  médiane **2,347 s backend**, full **76,724 s** médian, 2 704 modules et
  27 sorties exactes dans chaque paire. Gain de **38,90 %** face à D. Vrai `b`
  **6,836 s total / 2,277 s backend**, replis/capsules et daemons validés.
  Les sept suites de régression sont passées. Audit strict : **14 gates PASS**,
  seul le seuil de vitesse **FAILED** ; résultat global **failed / exit 1**.
  Les sondes résiduelles étayent une impasse pratique des petites retouches,
  sans établir de plancher théorique. Voir [C2-10](docs/build-c2-sub2.md).
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

- **gopurs** : worktree principale sur branche **`c2`** ; `c1b` existe aussi.
  Changements C2/C1c/docs non commités. Aucun commit de l'assistant.
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

**C2-10 clos pour impasse pratique avec L (`69c8e28a…`).** Trois paires exactes
donnent **2,347 s médian** ; vrai `b` **6,836 s total / 2,277 s backend**.
Les sept suites et les quatorze gates de qualité/provenance sont verts ; seul
le seuil de 2 s échoue, donc audit global **failed / exit 1**. Aucun gate
n’est en attente. Le [bilan](docs/build-c2-sub2.md) distingue cette impasse
d’un plancher théorique. Les points et identités D ci-dessous restent
l’historique C2-9 ; toute nouvelle architecture exige un périmètre distinct.

1. **A est clos** : pack par défaut, froid installé 76,163 s, hit 0,914 s,
   feuille 45,421 s, sorties exactes.
2. **C2-4 historique** : backend monolithique **14,734 s** sur le cas testé ;
   la preuve compacte vient de C2-7 ; les mesures closes sont celles de C2-9, C2-10 restant en validation.
3. **C2-5 est branché** (désormais par défaut via le lanceur Go natif) : overlay C1c dans
   `preparePackageSession`/`writePackageFiles`, entrées et imports réécrits,
   fallback full. Go packages **0,003–0,021 s**, backend **17,323–21,515 s**
   (médiane 19,773 s), sorties Ping identiques au full.
4. **C2-7 central validé** : médiane **16,970 s** après édition littérale,
   trois paires byte-exactes. Extension aux changements structurels/types/FFI
   et campagne quatre hôtes restent hors de cette validation. À la demande de
   l'utilisateur, le lanceur Go natif active désormais le chemin validé par
   défaut ; les autres hôtes gardent leur activation explicite.
5. **C2-9 achevé avec D.** Trois paires centrales exactes,
   chaque incrémental sous 4 s ; replis sur corruption/absence et daemons
   un/huit workers validés, compilation Go du dernier résultat réussie. Le vrai
   `b` sans exports fonctionne : 3,826 s backend, 8,791 s total. Le seuil reste
   celui du backend, après édition littérale centrale et avec cache amorcé.
   L'objectif de poursuite des optimisations atteint son seuil d'arrêt autorisé.
6. **HANDOFF détaillé** : branche `c2`, binaire final validé D
   `13160b1ab6689bee0bf930e4292e841b67d0e4c7ed4aab2ce538bc04ae2ab742`.
   Évidence : `.gopurs-evidence/c2-9-sub4-20261023/` depuis `htdocs`, campagnes
   `candidate-d/` et `candidate-d-final/`, sources `candidate-d-source-state.json`,
   liaison binaire `candidate-d-build.json`, audit `completion-audit.json` et
   détails dans `docs/build-c2-sub4.md`. Médianes D : backend **3,841 s**,
   full **78,111 s**, processus **4,432 s**. C2-8 v4 (`0394c4fe…`) reste conservé
   sous `.gopurs-evidence/c2-8-sub10-20261023/central-validated-v4/`, avec
   backend **8,586 s**, full **79,464 s**, processus **9,197 s** et correction Aff.
   La campagne v3 (`3ceb9250…`) reste
   conservée sous `candidate-v3/`, médiane backend **8,506 s**, full **78,590 s**.
   Le binaire C2-7 `457bfcf3…` et sa campagne restent dans
   `.gopurs-evidence/c2-7-reprise-20261023/central-validated-pairs/`.
   Ancien binaire `33f8ba03…` ; daemon historique `cebd9c75…` ; overlay
   build 7 `102f3d3e…` / build 8 `0dbbfcf4…`. Aucun commit de l'assistant. Le code C2/C1c
   n'est pas installé dans l'arbre P00. La validation C2-9 est achevée ; aucun
   merge n'a été effectué par l'assistant.
   - Fichiers C2 : `src/Gopurs/Driver.purs`,
     `src/Gopurs/Driver/Build.purs`, `src/Gopurs/Driver/Prepare.purs`,
     `src/Gopurs/Driver/Output.purs`, `src/Gopurs/Cache/Modules.purs`,
     `src/Gopurs/Monomorphization.purs`, `src/Gopurs/Resident.purs`.
   - Fichiers Go : `tools/build-cache/packages.go` (cache de plan, signatures),
     `tools/build-cache/packages_test.go`, `tools/build-cache/session.go`
     (`GOPURS_KEEP_STAGE`).
   - C1c : `tools/c1c` (manifeste, overlay, réécriture, tests E2E).
   - Docs : `docs/build-c2-resident.md`, `docs/build-c2-go-packages.md`.
   - Tests : `go test -race ./...` dans `tools/build-cache` et `tools/c1c`.
   - Go packages exige le bon symlink `app/gopurs` vers
     `p00-20261019/tree/gopurs` ; sinon `Promise_Internal_NewPromise` échoue.
   - Preuves : `.gopurs-evidence/a-20261020-c2-*`, `a-20261021-*`,
     `a-20261020-partition`.
7. Garder Go séquentiel, le stress multi-éditions et les trois jeux de sorties
   froides comme contrôles de régression du pack.
