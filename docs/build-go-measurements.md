# Mesures du layout Go — G06 validé

**Livré le 16 octobre 2026 : G à 7/7, P à 0/7. La pause avant P00 est effective.**
Le candidat 7 mesuré est installé, avec sélection explicite du layout par
`GOPURS_GO_PACKAGES=1`. Les validations finales et l'audit sont conservés dans
`gopurs-go-packages-g06-20261014/final-reference/summary.json`.

Les seuils temporels de `todo.md` restent ceux fixés avant G : **−30 % de `go build` et
−20 % jusqu'à `exec`** sur chacune des éditions feuille, partagée et FFI ;
régression temporelle ailleurs ≤5 %, dont le froid ; budget de l'identique
`max(5 % de la référence, 0,100 s)`. Le budget RSS initial était de **+5 %**
maximum pour le backend et le compilateur Go séparément. Go généré, objets gopurs, GOCACHE et exécutables : **+10 %**
maximum par catégorie, à scénario et périmètre identiques.

**Priorité révisée par l'utilisateur pendant le pilote 3, le 14 octobre 2026 :**
le temps prime ; le RSS est secondaire sauf aggravation démesurée. Le seuil RSS
historique de +5 % devient donc indicatif pour le verdict G06. Le **+13,4 %**
observé sur le premier seed froid du candidat 3 est un coût à documenter, pas un
motif de rejet à lui seul. Aucun nouveau plafond numérique n'a été fixé. Les
mesures RSS backend/Go, les protocoles et les analyses historiques sont conservés ;
une lecture complémentaire `analysis-time-priority.json` distingue les critères
temps/stockage encore bloquants. Les décisions antérieures restent datées de
leur protocole initial. Les seuils temporels, le stockage et les ressources de
mesure sont maintenus.

## Bilan du candidat 7 — mesures terminées le 16 octobre 2026

**36 chaînes / 36 oracles passent**, avec trois paires alternées pour chacun des
six scénarios. `full-2/combined/` réunit les 30 nouvelles chaînes et les six froides
de `cold-complete-7/`, sur le même candidat et les mêmes helpers/environnement
figés. **Les 39 critères bloquants de temps et de stockage passent.** Les preuves
du froid initial restent byte-identiques. Les validations finales des grands
corpus, du bootstrap, des snapshots, du cache et du runner passent : **G06 est clos**.

### Deux périmètres de temps à lire ensemble

La référence immédiate est `publicsync`, déjà doté du cache v3. La chaîne couvre
Spago, gopurs, tidy et le runner jusqu'à `exec` ; l'exécution de la suite métier
reste hors de ce temps. Les valeurs ci-dessous sont les médianes en secondes.

| Scénario | Backend référence → candidat | Variation backend | Go référence → candidat | Variation Go | Chaîne référence → candidat | Variation chaîne |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| Identique | 1,579 → 1,571 | −0,48 % | 0,440 → 0,300 | −31,82 % | 4,160 → 4,050 | **−2,64 %** |
| Feuille | 38,214 → 48,487 | **+26,88 %** | 228,000 → 48,490 | −78,73 % | 268,816 → 99,772 | **−62,88 %** |
| Partagé | 36,967 → 49,672 | **+34,37 %** | 224,690 → 76,510 | −65,95 % | 289,840 → 149,661 | **−48,36 %** |
| FFI | 28,048 → 40,001 | **+42,62 %** | 234,180 → 77,320 | −66,98 % | 265,350 → 120,738 | **−54,50 %** |
| Froid applicatif | 115,346 → 117,830 | +2,15 % | 232,090 → 118,040 | −49,14 % | 350,100 → 238,438 | **−31,89 %** |
| Froid complet | 115,177 → 117,740 | +2,23 % | 239,500 → 128,850 | −46,20 % | 385,109 → 276,664 | **−28,16 %** |

**Le bénéfice après édition vient de Go ; le backend gopurs ralentit.** Les
objectifs G des éditions portent sur Go et la chaîne. Les budgets froid/identique
évaluent séparément backend, Go et chaîne. Le froid applicatif garde les TAST et
les dépendances Go externes préchauffées ; le froid complet recrée aussi le
frontend et part de GOCACHE privés vides.

Cette acceptation face à `publicsync` **ne valide pas un retour aux ~30 secondes
de backend évoquées au début du chantier**. Le backend reste à environ **118 s à
froid** et **48,5 s sur la feuille**. Le coût du cache persistant accumulé avant G,
notamment sa publication froide, reste présent. Le lot 04 avait déjà documenté
**94,191 → 113,213 s (+20,19 %)** face au lot 03, comparaison appariée distincte.
Les anciens ~30 s et les centaines de secondes de chaîne ne constituent pas le
même chronomètre. P reste à 0/7 et ne démarre pas avant la pause convenue ; aucun
gain futur de P n'entre dans le verdict G.

### Dispersion et marges

| Scénario | Plage de chaîne référence | Plage candidat | Variations des trois paires |
| --- | ---: | ---: | --- |
| Identique | 4,087–4,165 s | 3,978–4,053 s | −0,84 / −2,76 / −4,38 % |
| Feuille | 254,104–277,411 s | 96,969–104,227 s | −61,84 / −61,23 / −64,04 % |
| Partagé | 284,930–300,077 s | 149,099–155,113 s | −47,67 / −46,48 / −50,13 % |
| FFI | 253,894–269,822 s | 119,789–121,111 s | −55,12 / −54,50 / −52,82 % |
| Froid applicatif | 337,686–351,396 s | 234,835–238,649 s | −32,15 / −31,83 / −30,46 % |
| Froid complet | 372,982–389,313 s | 269,612–282,369 s | −27,71 / −28,94 / −26,68 % |

Les paires backend froid applicatif sont à **+1,26 / +2,15 / +0,26 %**, avec
**3,283 s** de marge médiane au budget ; froid complet : **+4,10 / +3,45 / +1,95 %**,
marge **3,196 s**. L'identique ne réécrit aucun Go et ne compile ni ne lie. Toutes
les invocations conservent les TAST attendus et les octets Go propres à leur
layout ; les quatre états du candidat reproduisent l'inventaire déjà figé.

### Stockage et mémoire

Le Go généré baisse d'environ **8,08 %**. La variation des objets va de **−2,04 %
à +1,084 %** selon le scénario ; celle des exécutables de **+4,846 % à +5,240 %**.
Le GOCACHE total hôte + runner baisse après édition (**−22,50 % feuille /
−10,67 % partagé et FFI**), mais reste extrêmement proche du plafond sur
l'identique : **1 140 143 579 → 1 254 133 961 octets, +9,997897 %**, soit environ
**23 976 octets de marge** au seuil +10 %. Cette marge concerne le corpus figé,
pas une garantie pour tout projet. Les froids sont à **+9,8998 % / +9,9914 %**.

Le RSS backend augmente de **1,09 % à 11,21 %** suivant le scénario. Le pic RSS du
compilateur Go baisse de **43,41 % à 47,01 %** sur les scénarios qui compilent.
Ces variations sont publiées comme coût secondaire selon la priorité utilisateur ;
les hausses backend observées restent de cet ordre, sans multiplication du RSS.

Les [36 lignes de mesures](build-go-cycle-g06.csv) conservent les temps ventilés,
ressources, volumes bruts avant archivage, réécritures, oracles et empreintes des
binaires. Les 15 archives par paire couvrent les 30 nouvelles chaînes, avec
restauration vérifiée et journaux/pointeurs/inodes préservés ; le froid complet
conserve sa couverture antérieure.

### Layout installé et actions Go observées

Le 16 octobre, la densité **1/8192** et les artefacts du candidat 7 sont installés
après audit des **295 fichiers source/outils**. La revue des neuf snapshots
retrouve **612 déclarations Main/FFI**, les entrées et les imports d'initialisation
externes identiques après normalisation des qualifications de packages.
Les **207 commandes strictes** sur l'installation et ses **89 commandes CLI**
passent. Le bootstrap courant reproduit les **512 TAST, 23 Go et le binaire**
de la référence validée, puis passe le contrôle CLI. Les **10 contrôles du runner
passent sur macOS et Linux**, avec migration, dépendances, concurrence, annulation
et absence de compilation/lien à l'identique. L'audit final retrouve les mêmes
artefacts et les 295 fichiers de l'inventaire, avec les deux fixtures d'ancres
explicitement adaptées après installation.
Les [commandes d'activation et de migration](build-cache.md#sélectionner-le-découpage-go-g06)
emploient `GOPURS_GO_PACKAGES=1` ; la sélection du layout reste explicite.

Le graphe b8x vérifié par `go list` comporte six packages générés. Chaque ligne
importe les lignes précédentes, soit **15 arêtes**, plus ses imports externes :

| Ordre | Package sous `gopurs/output/purescript/` | Dépendances générées |
| --- | --- | --- |
| A | `p_root` | aucune |
| B | `p_4a7cb91a03c02000599e1c8c5e7db74a` | A |
| C | `p_cdbf3019a3a300007f3d219dbc66ccf3` | A, B |
| D | `p_6962c9722c1020003fbdca2a3f0977a9` | A, B, C |
| E | `p_fb6aa4931c9c800080c6393704c5a9fd` | A, B, C, D |
| F | `p_f0be974f136b400019e98aceab298411` | A, B, C, D, E |

`init_external` et `init_all` assurent l'initialisation ; runtime et entrées
s'ajoutent à ces six packages. Les **27 Go** totalisent **122 994 214 octets**.
Ce volume statique est distinct des gains temporels mesurés.

| Scénario | Go réécrits | Packages générés compilés | Actions compile / link du runner |
| --- | ---: | --- | ---: |
| Identique | 0 | aucun | 0 / 0 |
| Feuille | 1, dans F | F | 3 / 1 |
| Partagé | 2, dans E et F | E, F | 4 / 1 |
| FFI | 1, dans E | E, F | 4 / 1 |
| Froid applicatif | 27 | A à F | 9 / 1 |
| Froid complet | 27 | A à F, runtime et dépendances externes froides | 249 / 1 |

Ces actions sont identiques dans les trois répétitions de chaque scénario.
Après édition, `init_all` et l'entrée `command-line-arguments` expliquent les
deux compilations supplémentaires ; en FFI, F est recompilé par dépendance à E
sans réécriture de son Go. Les packages A à D sont réutilisés dans les trois
éditions. Go décide des actions à partir de ses clés de contenu et de dépendances.
La référence monolithique compile deux actions après édition et en froid
applicatif, mais l'une porte sur tout `purescript` ; compter les actions ne
remplace donc pas leur chronométrage.

## Historique — candidat initial de stockage

Les fichiers G03 représentaient **156 309 879 octets**, contre **133 809 190**
dans le relevé monolithique correspondant (**+16,82 %**). Parmi les fragments
générés, le qualificateur `gopurs_runtime.` apparaît **1 403 720 fois**.

Le candidat utilise un alias court de cet import, choisi après collecte des noms
réservés du package et du fragment. Seuls les objets `types.PkgName` du runtime
sont réécrits. Les imports pointés, imports anonymes, sélecteurs locaux et chaînes
gardent leur sens. Les chemins et identités de packages, noms de symboles, champs
record, ordre d'initialisation et tags restent gouvernés par le même plan.

Le reproducteur conserve une erreur initiale de typage dans le harnais, puis
l'échec de volume avec le code antérieur (**2 095 → 2 142 octets** pour son
fragment). La suite du résolveur passe avec le candidat ; elle compile les
versions monolithique et multipackage et vérifie le masquage des identifiants.
Le premier pilote confirme ce gain de stockage, mais rejette le candidat pour
ses coûts froids. Ses snapshots candidats sont examinés, sans remplacement des
références G04 à ce stade.

## Pilote 1 : alias court, candidat rejeté

Une paire exploratoire, sans autre campagne lourde concurrente, donne :

| Édition feuille | `publicsync` | Packages | Variation |
| --- | ---: | ---: | ---: |
| gopurs | 35,193 s | 62,727 s | +78,24 % |
| `go build` | 217,08 s | 35,53 s | −83,63 % |
| Chaîne jusqu'à `exec` | 254,702 s | 101,328 s | −60,22 % |
| Go généré | 133 809 196 octets | 137 412 898 octets | +2,69 % |
| RSS backend | 11 129 536 512 octets | 11 529 224 192 octets | +3,59 % |
| RSS compilateur Go | 15 833 804 800 octets | 999 821 312 octets | −93,69 % |

Les deux oracles passent. La feuille recompile 29 packages, contre 2 actions de
compilation monolithiques. Cette paire n'est pas une validation des seuils.

La revue du harnais relève ensuite deux écarts de périmètre : les volumes ont
été capturés avant l'oracle et le cache d'exports hôte était commun aux deux
générateurs. Ces mesures restent identifiées comme telles. La version suivante
mesure après l'oracle, clone les seeds avant cet oracle et isole aussi le cache
d'exports par générateur et par échantillon.

La préparation des seeds, avec backend et GOCACHE vides mais TAST conservés,
observe **347,542 → 636,399 s** de chaîne et **111,536 → 488,491 s** de gopurs.
La publication seule occupe **399,129 s** côté packages. Le RSS backend froid
passe de **10 566 778 880 à 11 622 252 544 octets** ; le GOCACHE de ce même état
de **1 136 605 702 à 1 644 772 137 octets**. Ces dépassements interdisent de
retenir le candidat 1. Ils sont conservés avec les bons résultats feuille.

## Regroupement de fichiers en validation

Le candidat suivant regroupe les sections source dans
`purescript/<package>/0_gopurs_package.go`. Il garde le plan de dépendances et
ordonne les sections par nom de fichier source, puis les déclarations dans leur
ordre initial. Les imports nommés et pointés sont résolus par les objets typés ;
les noms locaux restent réservés. Un préfixe numérique évite les collisions avec
les noms de modules pendant une migration sur système insensible à la casse.

Les directives de fichier et les commentaires réservés aux repères de provenance
gardent les fragments séparés, en utilisant le même résolveur d'imports. Le
reproducteur couvre notamment l'import simultané d'un même package avec `.` et
avec un alias, ainsi que l'ordre d'initialisation de deux FFI. Son échec initial
de harnais et le défaut observé dans le repli fragmenté restent conservés.

Le cache publie le fichier regroupé comme sortie `package` sans propriétaire
module unique ; les sources privées G05 conservent la provenance par module.
Les vues de snapshots sélectionnent les sections Main et FFI demandées avec les
imports du fichier réel. Leur lecture distingue les commentaires des chaînes Go.
Les migrations fragmenté/regroupé et les protections des fichiers externes sont
vérifiées dans le store commun. Le regroupement n'a pas encore de mesure acceptée.

Le candidat 2 est reconstruit sans PGO avec un target Rust neuf. Les tests du
résolveur, y compris le repli fragmenté corrigé, passent ; les tests de cache
passent avec détection de races sous Go 1.22.12. Le test de projection vérifie
aussi qu'un faux repère dans une chaîne Go brute ne crée pas une section source.

Les **56 comparaisons / 350 commandes** de transitions, la récupération du cache
sur les quatre hôtes et les **89 commandes CLI** passent. La revue indépendante
des neuf snapshots retrouve **612 déclarations Main/FFI identiques**, après la
seule normalisation des alias d'import. Les fichiers d'ordre, d'initialisation
et les entrées restent exacts. La mise à jour après compilation/exécution est
suivie du rejeu strict **36 passages / 207 commandes**, sans modification des
empreintes pendant le rejeu. Les transports WASM, Go natif et C/Rust donnent
également les mêmes octets. Le WASM reconstruit dans le dépôt égale celui du
candidat figé.

`functional-candidate-2/summary.json` atteste cette validation de correction.

## Pilote 2 : regroupement correct, budgets non atteints

Le pilote 2 termine sans autre campagne lourde concurrente, avec une paire
exploratoire par scénario. Les quatre oracles mesurés et les deux oracles des
seeds passent. Le relevé des volumes suit désormais l'oracle, avec caches
d'exports hôte et caches du runner privés par générateur et échantillon.

| Mesure | `publicsync` | Candidat 2 | Variation |
| --- | ---: | ---: | ---: |
| Chaîne feuille | 262,842 s | 88,153 s | −66,46 % |
| `go build` feuille | 224,79 s | 34,37 s | −84,71 % |
| gopurs feuille | 35,751 s | 51,190 s | +43,19 % |
| Chaîne identique | 4,159 s | 4,884 s | +0,724 s |
| gopurs identique | 1,625 s | 1,752 s | +0,126 s |
| `go build` identique | 0,42 s | 0,78 s | +0,36 s |
| Chaîne du seed froid | 360,489 s | 282,595 s | −21,61 % |
| `go build` du seed froid | 245,53 s | 142,56 s | −41,94 % |
| gopurs du seed froid | 111,789 s | 136,992 s | +22,54 % |

Les deux relances identiques ne réécrivent aucun Go et ne compilent ni ne lient.
La feuille déclenche 29 compilations de packages contre 2 dans la référence.
Le froid des seeds conserve les TAST ; il ne remplace pas les scénarios
`cold-application` et `cold-complete` prévus pour la campagne finale.

| Ressource après oracle, base froide | `publicsync` | Candidat 2 | Variation |
| --- | ---: | ---: | ---: |
| Go généré | 133 809 190 octets | 131 614 774 octets | −1,64 % |
| Objets gopurs | 548 621 658 octets | 563 400 927 octets | +2,69 % |
| GOCACHE runner | 1 136 866 936 octets | 1 616 620 711 octets | +42,20 % |
| GOCACHE hôte | 3 280 696 octets | 174 187 319 octets | +170 906 623 octets |
| GOCACHE total | 1 140 147 632 octets | 1 790 808 030 octets | +57,07 % |
| Exécutables | 95 678 320 octets | 102 126 492 octets | +6,74 % |
| RSS backend | 10 493 444 096 octets | 11 766 104 064 octets | +12,13 % |
| RSS compilateur Go | 15 688 208 384 octets | 996 937 728 octets | −93,65 % |

Le GOCACHE total de la feuille baisse de **2 136 748 416 à 2 087 769 035 octets
(−2,29 %)** : le seed et la nouvelle version monolithique coexistent dans ce
scénario. Ce résultat ne compense pas le dépassement sur la base/identique.
Le RSS backend feuille augmente aussi de **11,24 %**. Le candidat 2 est donc
**rejeté pour l'acceptation G06**, malgré sa correction validée et ses gains Go.
`pilot-2/analysis.json` conserve les décisions par métrique ; une seule paire ne
permet pas encore d'estimer la dispersion ni de valider les seuils.

### Diagnostics suivant le pilote 2

La lecture des archives applicatives réellement compilées donne **996 532 268
→ 1 472 406 670 octets**. Les exports `__.PKGDEF` passent de **23 020 816 à
74 943 157 octets** ; les tables de chaînes des objets de **159 052 077 à
340 791 855 octets**. Parmi les noms de symboles reconnus, **127 663 035 octets**
sont répétés entre objets du candidat. Le regroupement de sources réduit le
nombre de fichiers, mais les objets de packages conservent des duplications de
types, métadonnées et code. Ces volumes ne sont pas des mesures de temps.

Un rejeu forcé avec profil d'allocations et traces GC conserve tous les octets Go.
Un second diagnostic échantillonne le RSS toutes les 200 ms : **11 050 926 080
octets** au pic de la référence contre **12 209 405 952 octets** pour le candidat.
Le candidat atteint **11 108 024 320 octets** pendant la partition, puis son pic
pendant la publication. Ces invocations instrumentées ne servent pas de temps
d'acceptation. L'essai de restitution des arènes temporaires à la frontière
`PackageFiles` conserve les octets, mais culmine encore à **11 742 101 504
octets (+6,25 %)** face à cette référence instrumentée. Il est rejeté et
`workspace.go` est restauré à l'octet près. Le profil suivant examine le partage
des chaînes d'identifiants du parseur, avec le même budget RSS de +5 % et un
maximum de +5 % de temps face au candidat 2. Son premier build, échoué sur une
accolade manquante, est conservé avant la reprise corrigée. Celle-ci conserve
aussi les octets Go, mais culmine à **11 707 400 192 octets (+5,94 %)**. Cet essai
est également rejeté ; `bindings.go` revient exactement au candidat 2. Ces
diagnostics uniques ne sont pas promus en gains confirmés ni en estimations de
dispersion.

Un probe privé prépare aussi la densité **1/512**, déjà conservée en G01, à
partir des **2 997 sources G05 de base**. Son témoin 1/128 reproduit exactement
les sorties du candidat 2 : **525 packages / 35 596 arêtes / 131 614 774 octets**.
La variante 1/512 produit **123 packages / 5 261 arêtes / 126 354 274 octets** ;
ces tailles statiques ne valident aucun seuil temporel. L'arbitrage porte
sur la duplication des objets, le coût à chaud et les recompilations plus
larges ; les seuils G restent inchangés. La première commande de compilation
échoue avant Go car l'image n'a pas `/usr/bin/time`. La reprise conserve cet
échec et utilise le chronomètre intégré à Bash. La variante 1/512 compile, passe
l'oracle de base et trois relances sans action de compilation/lien. Son GOCACHE
runner après oracle est de **1 397 640 250 octets**, encore **+22,94 %** de la
référence monolithique du pilote 2. Les temps directs observés sont **120,607 s**
à froid puis **0,279 / 0,203 / 0,189 s** à chaud ; ce parcours direct et ces
invocations uniques ne remplacent pas les paires du runner réel. La variante
privée suivante 1/2048 produit **39 packages / 713 arêtes / 124 811 178 octets**.
Elle compile en **123,077 s**, passe l'oracle et trois relances sans action
(**0,231 / 0,161 / 0,169 s**). Son GOCACHE runner après oracle reste à
**1 325 667 660 octets (+16,61 %)** de la référence monolithique. Les archives
applicatives représentent **1 185 930 274 octets**, dont **255 631 186 octets** de
tables de chaînes. Ce probe reste donc hors budget de stockage, avant même les
exports hôte. Le probe 1/8192 réduit le plan à **6 packages / 15 arêtes /
122 994 214 octets**. Il compile en **124,831 s**, passe l'oracle et trois
relances sans compilation/lien (**0,330 / 0,154 / 0,153 s**). Son GOCACHE runner
après oracle représente **1 245 896 235 octets (+9,59 %)** du monolithe : la
marge jusqu'à +10 % reste étroite, avant ajout du cache hôte. Ses deux
exécutables occupent **100 314 910 octets**.

Le plancher de 8 GiB libres a arrêté ce dernier probe après le build froid.
Le harnais appliquait aussi ce contrôle au nettoyage du conteneur ; ce défaut
est corrigé et le conteneur arrêté explicitement. La reprise vérifie les
empreintes de toutes les sources et conserve le build froid. Les caches achevés
1/512 et 1/2048 sont archivés après reconstruction et vérification : **4 332
fichiers / 2 723 307 910 octets logiques → 400 535 552 octets d'archive**.

Le **candidat 3 privé** réunit 1/8192 et les exports allégés ci-dessous pour les
quatre hôtes. Sa reconstruction passe ; l'archive Go liée par Rust est contrôlée
dans un target neuf. L'audit des sources ne trouve que les changements
d'acquisition d'exports, de leurs tests, de densité privée et de WASM privé.
La densité du dépôt reste à 1/128. Le premier lancement du script
de reconstruction échoue avant toute construction sur une indentation ; sa
sortie est conservée et la reprise utilise le script corrigé. L'acceptation
demande toujours le cache total hôte + runner, les trois éditions, les scénarios
froids, le RSS et les paires du parcours complet.

La suite du résolveur passe pour ce candidat privé. Les **89 commandes CLI**
et les **56 comparaisons / 350 commandes de transitions** passent sur JS, Go
séquentiel/parallèle et Rust, avec mêmes TAST, mêmes objets métier/préparés et
mêmes Go exécutables. `functional-private-3/summary.json` conserve ce contrôle ;
les snapshots de référence restent ceux du candidat 2. Le **pilote 3** termine
sur identique, feuille, partagé et FFI, une paire exploratoire par scénario.
Les caches du runner partent de la base seule ; les volumes hôte/runner sont
relevés après oracle selon le périmètre corrigé du pilote 2.

### Pilote 3 : objectifs des éditions dépassés, acceptation encore ouverte

Les huit oracles des mesures et les deux des seeds passent. Une seule paire ne
permet pas encore de conclure sur la dispersion ou l'acceptation finale.

| Mesure | `publicsync` | Candidat 3 | Variation |
| --- | ---: | ---: | ---: |
| Chaîne identique | 4,215 s | 3,952 s | −6,23 % |
| Chaîne feuille | 259,300 s | 96,466 s | −62,80 % |
| `go build` feuille | 222,72 s | 45,06 s | −79,77 % |
| gopurs feuille | 34,188 s | 49,172 s | +43,83 % |
| Chaîne partagée | 283,712 s | 156,083 s | −44,99 % |
| `go build` partagé | 223,75 s | 79,82 s | −64,33 % |
| gopurs partagé | 36,617 s | 52,469 s | +43,29 % |
| Chaîne FFI | 261,076 s | 123,012 s | −52,88 % |
| `go build` FFI | 229,35 s | 78,08 s | −65,96 % |
| gopurs FFI | 28,749 s | 42,069 s | +46,33 % |
| Chaîne du seed froid | 340,480 s | 250,014 s | −26,57 % |
| gopurs du seed froid | 112,483 s | 123,740 s | +10,01 % |

L'identique ne réécrit, ne compile et ne lie rien. Les éditions feuille,
partagée et FFI déclenchent respectivement **3 / 4 / 4 compilations**, contre
**2 / 2 / 2** dans la référence, avec un lien dans chaque mesure.

Le GOCACHE total de base/identique atteint **1 254 221 155 octets**, contre
**1 140 147 632** : **+10,005 %**, soit environ 59 ko au-delà du budget maintenu.
Les objets gopurs dépassent aussi le budget après partagé (**562 168 328 →
680 702 881 octets, +21,09 %**) et FFI (**566 156 271 → 714 586 193, +26,22 %**).
Les Go générés baissent d'environ **8,08 %** ; les exécutables restent dans le
budget. Les objets, le cache Go et les exécutables sont des volumes, pas des gains
de temps. Le RSS backend froid augmente de **13,36 %** ; le pic du compilateur Go
passe de **15 846 334 464 à 8 355 500 032 octets**. Le RSS est traité selon la
priorité révisée, sans effacer le verdict historique.

`analysis-time-priority.json` conserve les trois échecs de stockage. Le surcoût
backend du seed froid est observé séparément ; les scénarios `cold-application`
et `cold-complete`, ainsi que les trois paires alternées, restent nécessaires.

### Diagnostics suivant le pilote 3

Le relevé de portée des objets retrouve **126 120 311 / 160 018 756 octets**
non référencés par le manifeste courant après partagé/FFI. Les objets utiles
restent autour de **554,6 Mo**, contre **548,6 Mo** dans le monolithe. Les gros
fichiers regroupés amplifient l'accumulation des générations précédentes.

Un entretien après publication est en validation dans le cache commun : il
reprend le verrou, relit le manifeste courant et son journal de propriété,
conserve ses sorties, ses quatre objets métier, ses sources G05 et ses
diagnostics. Une publication incomplète suspend cet entretien ; les lecteurs
ont déjà figé leurs entrées sous le même verrou dans leur staging privé. Le
reproducteur échoue avant correction ; la suite cache avec détection de races
passe sous Go 1.27.0 et 1.22.12 après correction. Son temps reste à mesurer dans
la chaîne.

Le test concurrent supplémentaire alterne deux publications de sources
distinctes, huit gels de candidats et huit entretiens. Chaque lecteur retrouve
une version complète de sa source privée et le manifeste final donne un hit.
Il passe avec détection de races sous Go 1.27.0 et 1.22.12. Le rejeu depuis un
staging antérieur recrée aussi un objet supprimé entre-temps.

Le profil CPU autonome de `PackageFiles` conserve le plan et tous les octets Go.
Sur cet hôte, **77 % des échantillons CPU** sont attribués à `runtime.madvise`.
Trois paires alternées comparent GC désactivé et `GOGC=100`, tous deux sous
`GOMEMLIMIT=10GiB` : médianes **11,337 → 11,744 s (+3,59 %)**, plages
**11,230–12,302 / 11,518–12,036 s**. Le RSS passe de **8 021 164 032 à
4 111 712 256 octets**, mais le seuil local de **5 % de gain de temps** échoue.
Cet essai est **rejeté**, conformément à la priorité au temps. Le profil et ces
mesures autonomes ne sont pas des temps d'acceptation de chaîne.

Le candidat suivant vise les copies du rendu : écriture des modifications
directement dans le fichier final, réserve de capacité, résolution des imports
par leurs objets typés et hachage des buffers déjà lus. La suite du résolveur
avec détection de races passe. Trois paires autonomes donnent **11,960 →
10,757 s (−10,06 %)**, avec gain dans chaque paire, pour des plages de
**11,715–11,991 / 10,743–11,382 s**. Le RSS médian passe de **8 021 098 496 à
6 271 516 672 octets**. Le plan et les fichiers sont byte-exacts. Ce gain local
dépasse le seuil préalable de 5 % ; il reste à vérifier dans le backend complet
et sur les quatre hôtes.

Le dépassement de GOCACHE est ciblé par `-trimpath` sur l'acquisition d'exports
allégés, uniquement quand `GOFLAGS` est vide. Le probe conserve les mêmes
**250 exports de packages**, le même résultat de partition et les mêmes Go.
Le cache isolé passe de **7 373 709 à 7 287 995 octets (−85 714)** ; les archives
d'exports passent de **5 271 792 à 5 186 870 octets**. Ce volume ne constitue
pas un gain temporel ni une validation du cache total de la prochaine campagne.
Les suites cache/race, dont compilation normale après acquisition des exports
et respect des options explicites, passent sous Go 1.27.0 et 1.22.12.

Le **candidat 4 privé** est reconstruit pour les quatre hôtes, à densité
1/8192 ; Rust utilise un target neuf, avec contrôle du chemin de l'archive liée.
L'audit retrouve uniquement le cache, ses tests, le rendu, son transport fichier
et le WASM recompilé. Sa validation fonctionnelle passe : **89 commandes CLI**
et **56 comparaisons / 350 commandes de transitions**, sur JS, Go séquentiel,
Go parallèle et Rust, avec mêmes TAST, objets métier/préparés et fichiers Go
exécutés. `functional-private-4/summary.json` conserve cette preuve.

Un contrôle ciblé termine sur **trois paires alternées de backend froid**,
avec TAST conservés mais sorties
Go, cache gopurs et GOCACHE hôte absents dans chaque échantillon. Ce contrôle
compare les fichiers complets à ceux du pilote 3 et garde le budget backend
froid de +5 % ; il ne remplace pas la campagne du runner jusqu'à `exec`.

| Paire froide | `publicsync` | Candidat 4 |
| --- | ---: | ---: |
| 1 | 113,034 s | 123,868 s |
| 2, ordre inversé | 112,674 s | 120,893 s |
| 3 | 107,511 s | 118,115 s |
| Médiane | **112,674 s** | **120,893 s** |

Le surcoût des médianes est de **7,29 %** ; chacune des trois paires dépasse
aussi le budget de +5 %. Le candidat 4 est donc **refusé au contrôle froid**.
Il manque **2,586 s** à sa médiane pour respecter la limite. Les six inventaires
TAST/Go restent byte-exacts face aux états du pilote 3. Les plages de temps sont
**107,511–113,034 / 118,115–123,868 s** ; le RSS médian est de
**10 566 811 648 / 11 652 235 264 octets**, conservé comme indicateur secondaire.

Les phases médianes du candidat occupent **7,032 s d'exports**, **11,358 s de
partition** et **42,478 s de publication**, contre **54,935 s de publication**
dans le monolithe. Les phases imbriquées/pipelinées ne s'additionnent pas pour
reconstruire le temps total. Les caches d'objets froids restent à **548 656 101 /
554 601 923 octets**. Le GOCACHE hôte vaut **0 / 7 296 196 octets** dans ce
contrôle sans `tidy`, runner ni oracle ; ce périmètre ne remplace pas le cache
total hôte + runner du pilote.

### Réutilisation des compresseurs G05 : essai rejeté

Chaque source nouvelle créait un writer gzip `BestSpeed`. L'essai suivant
réutilise ses buffers au moyen d'un pool et de `Reset`, qui réinitialise aussi
le flux et l'en-tête. Les writers sont privés pendant chaque opération ; le
codec, les objets et les bridges gardent leur contrat. Le test intercale des
sources vides, répétitives et binaires sur huit tâches, compare chaque objet
aux octets d'un writer neuf et vérifie sa décompression. Les suites cache/race
passent sous Go 1.27.0 et 1.22.12.

Trois paires alternées terminent sur le stockage effectif des **2 997 sources
b8x**, avec lecture, hash, compression, écriture et synchronisation dans de
nouveaux répertoires privés. Le seuil local préalable est de **5 % de gain** ;
les objets restent exacts face au manifeste froid du candidat 4. Les médianes
**10,418 → 10,478 s (+0,57 %)** échouent au seuil, avec des plages de
**9,261–10,465 / 9,691–10,896 s**. Les allocations passent de **2 641 056 720 à
210 843 816 octets**, et le RSS médian de **1 006 632 960 à 231 669 760 octets**,
mais ces baisses ne suffisent pas à retenir le changement. Le code et son test
sont conservés dans les preuves ; `modules.go` et `modules_test.go` reviennent
byte-exactement à leur état avant l'essai.

### Acquisition d'exports à quatre processus en validation

L'acquisition de types imposait `go list -p=1` pour ses dépendances. Trois
paires froides comparent cette limite à **quatre processus**, avec le même
hôte, `GOGC=off`, `GOMEMLIMIT=10GiB`, parallélisme runtime hôte par défaut et un
GOCACHE neuf par échantillon. La médiane passe de **6,269 à 2,482 s (−60,41 %)** ;
les plages sont **6,229–6,283 / 2,471–2,504 s**, avec gain dans chaque paire.
Les **250 archives d'exports sont byte-exactes** dans les six mesures ; le
cache vaut **7 287 995 octets** dans chaque cas. Le rejeu du résolveur donne
le même plan et les mêmes fichiers Go.

Le raccordement utilise `-p=4` sur le parcours d'exports allégés avec `GOFLAGS`
vide ; les options explicites suivent le parcours à `-p=1`. Les limites runtime
de l'hôte sont héritées par les compilations enfants. Il s'agit de parallélisme
interne à l'acquisition du backend ; le protocole du runner reste
`GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB`. Les suites cache/race passent sous Go
1.27.0 et 1.22.12. Le **candidat 5 privé est reconstruit et audité** : le seul
changement de production depuis le candidat 4 est cette limite de processus.
L'essai gzip rejeté est absent ; le WASM est identique au candidat 4. Les
**89 commandes CLI et 56 comparaisons / 350 commandes de transitions** passent
sur les quatre hôtes, avec les mêmes TAST, objets métier/préparés et fichiers
Go exécutés. Le contrôle de **trois paires froides** termine dans
`cold-backend-5/` :

| Paire froide | `publicsync` | Candidat 5 |
| --- | ---: | ---: |
| 1 | 111,196 s | 118,779 s |
| 2, ordre inversé | 107,918 s | 116,026 s |
| 3 | 110,519 s | 117,662 s |
| Médiane | **110,519 s** | **117,662 s** |

Le candidat 5 est **refusé au contrôle froid** : **+6,46 %**, avec dépassement
du budget de +5 % dans chacune des trois paires. Il manque **1,616 s** à la
médiane pour respecter la limite. Les six inventaires TAST/Go sont exacts ; les
plages sont **107,918–111,196 / 116,026–118,779 s**. Le RSS médian est de
**10 570 924 032 / 11 691 933 696 octets**, indicateur secondaire. Les objets
gopurs et le GOCACHE hôte ont les mêmes volumes que dans le contrôle froid 4.

Les phases médianes du candidat sont **3,068 s d'exports**, **10,898 s de
partition**, **41,970 s de publication** et **45,791 s d'optimisation/émission**.
La référence relève **51,942 s de publication** et **43,033 s
d'optimisation/émission**. Ces phases imbriquées ne s'additionnent pas pour
reconstituer le temps total.

### Profils réellement froids du candidat 5

Les deux diagnostics CPU/allocations/GC terminent avec TAST et Go byte-exacts,
sur racines neuves et caches vides. **Leurs temps instrumentés ne sont pas des
mesures d'acceptation.** Le profil du candidat attribue **6,58 s CPU** au parcours
`EmitPackageFiles`, dont **3,31 s** au résolveur, **1,43 s** au graphe et **1,53 s**
au rendu ; les coûts cumulatifs peuvent se recouvrir. Les allocations
échantillonnées de ce parcours représentent **5,34 Gio**, dont **3,50 Gio** dans
le résolveur. Le CPU global attribué à `runtime.madvise` passe de **11,83 à
16,00 s** ; ces échantillons ne donnent pas directement un gain mural réalisable.

La lecture du Go natif confirme que les continuations d'exports, de partition
et de publication capturent encore `prepared_7` pour relire `mainModules`,
alors que les modules préparés et métadonnées ne sont plus consommés après
l'émission. Cette rétention est une piste à mesurer, sans gain établi.

Un probe natif privé extrait ce tableau à l'entrée de la continuation de
préparation et remplace ses quatre relectures. Le changement porte sur les
captures des continuations ; les appels de bridges gardent leurs valeurs et
leur contrat. Le probe est construit depuis le candidat 5, avec le Go généré
avant/après conservé et une provenance propre. Son seuil préalable est de
**5 % de gain sur les phases séquentielles après émission** (exports,
partition, publication), sans régression du backend total, sur **trois paires
alternées** à froid. Les paires terminent avec six inventaires exacts, mais le
probe est **rejeté** : médianes après émission **58,274 → 56,603 s (−2,87 %)**,
plages **56,953–59,313 / 54,914–56,611 s**. Le backend total passe de **120,824 à
116,712 s (−3,40 %)**, avec plages **116,545–121,330 / 116,673–117,309 s** et
une légère régression dans la deuxième paire. Le RSS médian passe de
**11 549 491 200 à 11 857 330 176 octets**. Le gain local requis n'est pas
atteint ; l'artefact natif privé et ses sources avant/après sont conservés.

### Acquisition d'exports à huit processus en validation

Trois paires froides comparent ensuite le cap de quatre processus à huit,
avec les mêmes ressources runtime et les mêmes options d'exports allégés.
Le seuil local préalable est de 5 % de gain. Les médianes sont **2,555 →
2,166 s (−15,22 %)**, les plages **2,443–2,606 / 2,098–2,233 s**, avec gain
dans chaque paire. Les **250 archives d'exports sont byte-exactes**, le cache
vaut **7 287 995 octets** dans les six échantillons et le rejeu du résolveur
conserve le plan et les Go. Le raccordement passe le parcours `GOFLAGS` vide
à `-p=8` ; le parcours d'options explicites utilise `-p=1`. Le gain local
d'environ 0,39 s reste à vérifier dans un candidat complet reconstruit.
`validate-export-cap-8.py` termine séquentiellement les suites cache/race Go
1.27.0/1.22.12, la reconstruction du candidat 6 sur les quatre hôtes, l'audit,
la validation CLI/transitions puis les trois paires froides. Les **89 commandes
CLI / 56 comparaisons de transitions** passent avec mêmes TAST, objets métier,
clés préparées et Go exécutés. Le seul changement de production depuis le
candidat 5 est le cap d'exports quatre → huit ; les essais rejetés de capture
des entrées et de pool gzip sont absents.

| Paire froide | `publicsync` | Candidat 6 | Variation |
| --- | ---: | ---: | ---: |
| 1 | 112,299 s | 117,524 s | +4,65 % |
| 2, ordre inversé | 110,432 s | 116,155 s | **+5,18 %** |
| 3 | 112,877 s | 117,026 s | +3,68 % |
| Médiane | **112,299 s** | **117,026 s** | **+4,21 %** |

Le **contrôle froid passe sur les médianes**, avec six inventaires exacts,
mais la deuxième paire dépasse légèrement +5 %. Les plages sont
**110,432–112,877 / 116,155–117,524 s**. La marge médiane n'est que de
**0,888 s** ; la campagne de bout en bout reste nécessaire. La référence
mesurée est plus lente que pendant le contrôle 5 : la baisse du surcoût relatif
ne peut donc pas être entièrement attribuée au changement de cap.

Les phases médianes du candidat sont **2,724 s d'exports**, **11,228 s de
partition**, **42,973 s de publication** et **44,326 s d'optimisation/émission**,
contre **54,350 / 42,362 s** de publication/optimisation-émission dans la
référence. Le RSS médian est de **10 565 730 304 / 11 609 423 872 octets**,
indicateur secondaire. Les volumes d'objets et de GOCACHE hôte sont identiques
aux contrôles 4–5, toujours dans le périmètre backend seul.

### Campagne complète du candidat 6 — 36 chaînes terminées

La campagne `full-1/` est lancée avec **trois paires alternées sur chacun des
six scénarios** : identique, feuille, partagé, FFI, froid applicatif et froid
complet, soit **36 chaînes et 36 oracles attendus**, en plus des seeds de base.
Le harnais vérifié, les empreintes des quatre hôtes et les preuves du contrôle
froid sont recontrôlés dans `full-1-preflight.json`. Le runner conserve
**GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB** ; le backend conserve les huit workers
et les ressources hôte de G00/v3. L'archivage vérifié entre scénarios est activé.

Le lancement dispose de **26 510 938 112 octets libres** après l'archivage des
anciennes campagnes terminées. Le script `run-full-cycle-1.py` enchaîne les
analyses historique et priorité au temps une fois toutes les chaînes et tous
les oracles achevés. G06 reste ouvert jusqu'aux mesures complètes, aux parités
des grands corpus, au bootstrap et aux replays stricts.

La campagne est interrompue le **15 octobre à 12 h 11**, après **28 chaînes /
28 oracles passants**. Le frontend de `cold-application-3-reference` reçoit
**SIGTERM** ; le backend et le harnais conservent ce signal, et le conteneur est
arrêté normalement. L'émetteur du signal est inconnu. Les journaux d'alimentation
ne montrent aucune transition veille/réveil dans la plage locale 09 h–12 h 59.
L'invocation incomplète est exclue des résultats et conservée avec ses fichiers.

Les quatre scénarios chauds terminent leurs **trois paires** :

| Scénario | Chaîne référence → candidat 6 | Variation chaîne | Go référence → candidat 6 | Variation Go |
| --- | ---: | ---: | ---: | ---: |
| Identique | 4,224 → 4,132 s | −2,18 % | 0,410 → 0,280 s | −31,71 % |
| Feuille | 268,991 → 101,791 s | **−62,16 %** | 229,850 → 49,130 s | **−78,63 %** |
| Partagé | 290,427 → 154,048 s | **−46,96 %** | 230,260 → 78,940 s | **−65,72 %** |
| FFI | 245,445 → 114,843 s | **−53,21 %** | 216,200 → 72,560 s | **−66,44 %** |

Le backend seul augmente sur les éditions : **39,128 → 50,303 s** feuille,
**36,817 → 51,516 s** partagé et **26,815 → 39,421 s** FFI. Le gain Go domine
ce coût dans chaque chaîne. L'identique ne réécrit ni ne compile ni ne lie rien.
Son GOCACHE total est à **+9,9979 %**, juste sous +10 % ; les trois éditions
réduisent ce volume de **22,50 / 10,67 / 10,67 %**. Les objets gopurs sont à
**+1,08 / +0,35 / −1,35 / −2,04 %** sur identique/feuille/partagé/FFI :
l'entretien corrige l'accumulation du pilote 3. Les autres catégories de
stockage de ces scénarios restent dans leurs budgets.

Le froid applicatif n'a que **deux paires complètes** : chaîne
**327,383 → 225,338 s (−31,17 %)**, backend **111,174 → 114,714 s (+3,18 %)**,
Go **213,415 → 108,065 s (−49,36 %)**. Ces résultats sont provisoires.
La première reprise des **huit mesures restantes** est lancée avec le même candidat et les
mêmes seeds privés ; les quinze fonctions de mesure sont chargées directement
depuis le harnais figé et leur empreinte est contrôlée. Les deux générateurs de
la troisième paire froide applicative repartent sur des racines neuves. Les
28 résultats et toutes les commandes antérieures ont une copie préalable.
Le contrôle préalable compare aussi l'AST de l'environnement de mesure et les
volumes des seeds hôte/externes à leurs relevés initiaux. La séquence
`continue-full-cycle-1.py` est protégée contre la veille automatique par
`caffeinate` et lance l'entretien disque avant les mesures.

L'espace disponible impose d'archiver les quatre échantillons froids déjà
achevés avant la reprise. Cette exception à l'archivage en fin de scénario est
hors chronométrages, après l'arrêt de l'invocation interrompue. Les sources,
mesures, journaux, pointeurs et inodes des verrous restent à leur emplacement ;
chaque fichier archivé est restauré/vérifié. Les deux derniers échantillons
froids applicatifs seront archivés après leur oracle ; un contrôle de couverture
vérifie que les deux archives réunissent exactement les six échantillons et
leurs volumes mesurés. Ce contrôle de couverture passe après la première reprise.

#### Deuxième reprise : garde-fou disque après 33 chaînes

La première reprise termine cinq chaînes supplémentaires, toutes avec oracle
passant, puis le garde-fou de **8 GiB libres** arrête le parcours avant le
`tidy` de `cold-complete-2-reference`. Le frontend/backend de cette invocation
avait terminé avec succès. L'espace relevé après l'arrêt est de
**8 418 267 136 octets libres**. Ce résultat de chaîne incomplet est exclu ;
le backend terminé est conservé comme preuve séparée, TAST/Go exacts vérifiés.

Le froid applicatif a désormais **trois paires** : chaîne **328,718 →
225,789 s (−31,31 %)**, Go **214,250 → 108,160 s (−49,52 %)**, backend
**113,315 → 115,048 s (+1,53 %)**. Il passe les budgets temporels et de
stockage sur les médianes. Le backend de la troisième paire est à **+5,63 %** ;
les plages backend sont **109,032–116,550 / 114,381–123,112 s**.

La première paire complète de froid complet donne **401,521 → 291,413 s
(−27,42 %)** de chaîne et **255,860 → 137,160 s (−46,39 %)** de Go, mais le
backend est à **117,121 → 126,187 s (+7,74 %)**, au-dessus du budget de +5 %.
Une paire seule ne tranche pas ce scénario. Le côté candidat déjà terminé de
la deuxième paire est conservé, puis écarté du jeu apparié : les deux côtés
repartent ensemble sur des racines neuves après l'interruption.

`continue-full-cycle-2.py` lance séquentiellement l'archivage des trois
échantillons froids complets terminés, puis les **quatre mesures des paires 2
et 3**. Le jeu retenu repart donc de **32 chaînes**, avec sauvegarde des
**33 chaînes complètes**, dont une écartée du jeu apparié. L'archivage du froid
complet intervient désormais après chaque paire terminée, hors chronométrages.
Les fonctions de mesure et l'AST de l'environnement restent ceux du harnais
figé. Le contrôle final couvrira les six mesures retenues et l'échantillon
écarté, tous restaurables, sans masquer le premier arrêt disque.

#### Bilan des trois paires et extension du froid complet

La deuxième reprise termine avec **36 chaînes / 36 oracles passants**, TAST et
Go byte-exacts. Les 28 résultats d'avant le SIGTERM et les 32 résultats retenus
avant la deuxième reprise sont conservés dans le jeu final. La sauvegarde des
33 chaînes contient aussi le côté isolé de la paire interrompue, explicitement
écarté puis rejoué. La couverture des archives vérifie les **six échantillons
froids complets retenus et cet échantillon isolé** ; celle du froid applicatif
couvre ses six échantillons. Les volumes d'acceptation restent les volumes bruts.

| Scénario froid, trois paires | Référence | Candidat 6 | Variation |
| --- | ---: | ---: | ---: |
| Chaîne applicative | 328,718 s | 225,789 s | −31,31 % |
| Go applicatif | 214,250 s | 108,160 s | −49,52 % |
| Backend applicatif | 113,315 s | 115,048 s | +1,53 % |
| Chaîne complète | 401,521 s | 291,413 s | −27,42 % |
| Go complet | 253,360 s | 137,160 s | −45,86 % |
| **Backend complet** | **117,121 s** | **125,117 s** | **+6,83 %** |

L'analyse priorité au temps ne trouve **qu'un critère bloquant en échec** : le
backend froid complet dépasse de **2,139 s** son budget de +5 %. Le gain de la
chaîne ne compense pas ce dépassement. Les trois paires backend sont :

| Paire | Référence | Candidat 6 | Variation |
| --- | ---: | ---: | ---: |
| 1 | 117,121 s | 126,187 s | +7,74 % |
| 2, ordre inversé, rejouée | 116,268 s | 121,965 s | +4,90 % |
| 3, rejouée | 121,593 s | 125,117 s | +2,90 % |

Les plages backend sont **116,268–121,593 / 121,965–126,187 s**. La dispersion
traverse le seuil par paire alors que le ratio des médianes le dépasse. Le
protocole prévoit d'augmenter l'échantillon dans ce cas : **trois paires froides
complètes supplémentaires, numérotées 4 à 6, sont fixées avant leur lancement**.
La décision suivante portera sur **les six paires réunies**, avec toutes les
mesures actuelles, les mêmes seuils et le même candidat 6. Les médianes globales,
les plages et chaque écart apparié seront publiés. G06 reste ouvert.

`extend-cold-cycle-6.py` charge directement les **15 fonctions et les quatre
instructions d'environnement du harnais figé par AST**. Les paires supplémentaires
alternent candidat/référence, référence/candidat, candidat/référence. Chacune
repart sur des racines neuves, sans sortie frontend sauf `go.mod`/`go.sum`, avec
GOCACHE hôte/runner vide et téléchargements hors ligne disponibles. Le runner
conserve ses ressources, le backend ses huit workers. L'archivage restauré/vérifié
intervient après chaque paire et ses oracles, hors chronométrages. La réserve
disque préalable est portée à **16 GiB pour les trois paires** ; le plancher
pendant les commandes reste **8 GiB**. Les résultats et analyses des 36 chaînes
restent à leurs chemins ; le bilan cumulé de **42 chaînes attendues** sera dans
`cold-extension-6/combined/`.

L'extension est lancée en arrière-plan sous `caffeinate` après vérification des
archives et contrôle d'absence de campagne lourde concurrente. Le dernier relevé
préalable donne **17 493 106 688 octets libres**. Les 41 racines historiques
archivées pendant cette préparation ont leurs pointeurs, journaux et inodes
revérifiés avant lancement.

Le stockage des trois paires froides passe : GOCACHE total **+9,8998 % applicatif /
+9,9920 % complet**, objets **+1,0836 %**, exécutables **+4,8460 %**, Go généré
**−8,0824 %**. La marge GOCACHE reste étroite. Le RSS médian backend passe de
**10 568 679 424 à 11 725 324 288 octets (+10,94 %)** en froid applicatif et de
**10 565 943 296 à 11 655 020 544 octets (+10,31 %)** en froid complet. Le RSS du
compilateur Go baisse respectivement de **46,90 % et 47,62 %**. Ces coûts mémoire
restent documentés comme indicateurs secondaires selon la priorité utilisateur.

Les journaux déjà acquis situent les phases froides complètes du candidat à
**3,633 s d'exports**, **12,369 s de packages** et **43,477 s de publication**
en médiane ; la publication de référence vaut **56,880 s**. L'optimisation/émission
est à **46,658 → 47,800 s**. Ces phases peuvent se recouvrir : leurs médianes
ne s'additionnent pas en une attribution du surcoût. Le verdict continue à
utiliser le chronomètre backend externe. `full-1-cold-phase-observations.json`
conserve les valeurs par paire et les empreintes des journaux, sans nouveau run.

#### Verdict à six paires : candidat 6 hors budget froid

L'extension termine : **42 chaînes / 42 oracles passants**, dont **six paires
froides complètes**. Les 36 résultats et fichiers de preuve initiaux sont
revérifiés à l'octet près. Les six échantillons supplémentaires sont archivés
après leurs paires, avec couverture et restauration vérifiées.

| Froid complet, six paires réunies | Référence | Candidat 6 | Variation |
| --- | ---: | ---: | ---: |
| Chaîne jusqu'à `exec` | 398,610 s | 292,306 s | −26,67 % |
| Compilation Go | 247,900 s | 137,115 s | −44,69 % |
| **Backend** | **117,733 s** | **124,544 s** | **+5,79 %** |

Les écarts backend des paires 4–6 sont **+2,71 / +4,69 / +8,71 %**, après
**+7,74 / +4,90 / +2,90 %** pour les paires 1–3. Les plages réunies sont
**114,035–123,843 / 121,965–127,196 s**. Le ratio des médianes dépasse encore le
budget de +5 %, de **0,925 s** : **le candidat 6 est refusé pour l'acceptation
G06**. Les résultats des trois premières paires font intégralement partie de
cette décision ; les seuils ne changent pas.

Les autres critères temps/stockage passent. En froid complet, le GOCACHE total
est à **+9,9916 %**, les objets à **+1,0836 %**, les exécutables à **+4,8460 %**
et le Go généré à **−8,0824 %**. Le RSS backend médian est à **+10,35 %**
et celui du compilateur Go à **−46,34 %**, indicateurs secondaires. Les phases
journalisées sur six paires situent les packages à **12,072 s**, les exports à
**3,636 s**, la publication à **57,463 → 44,098 s** et l'optimisation/émission à
**44,564 → 47,588 s**. Elles restent non additionnables ; le temps externe
du backend porte le critère.

#### Probe privé de rendu en une passe

`renderPackages` parcourt chaque déclaration deux fois : une première passe
collecte imports/noms réservés et références à qualifier, puis une seconde
refait les lectures typées pour construire les réécritures. Le probe privé
prépare les réécritures dans la première passe et reporte seulement la résolution
de leurs alias jusqu'à l'écriture. Le graphe et les identités restent les mêmes.
Les requêtes/résultats de `PackageFiles` empruntent les mêmes bridges JS/WASM,
Go direct et C/Rust ; leurs fichiers doivent être byte-exacts.

Seuil fixé avant mesure : **au moins 5 % de gain sur le parcours autonome
`PackageFiles` de b8x, trois paires alternées**, mêmes sources/exports et
**GOGC=off / GOMEMLIMIT=10GiB / GOMAXPROCS hôte par défaut**. Les suites du
résolveur avant/après passent. Les trois paires donnent **11,463 → 10,761 s
(−6,12 %)**, avec gains **7,52 / 5,43 / 6,12 %**. Les plages sont
**11,269–11,770 / 10,657–10,885 s**. Les six plans et inventaires Go sont exacts ;
le RSS médian passe de **6 273 761 280 à 6 363 856 896 octets**. Le seuil local
est atteint. Ce gain d'environ **0,702 s** est propre au parcours autonome ;
il reste à mesurer après reconstruction du candidat complet, notamment avec la
préparation encore vivante en mémoire. `render-single-pass-probe/` conserve
le protocole préalable, les sources avant/après, les binaires et les six sorties.

Le rendu du probe est reporté exactement dans `tools/ffi-gen/packages.go`, avec
la densité du dépôt à 1/128. La séquence `validate-single-pass-7.py` est lancée
sous `caffeinate` : archivage des objets de neuf anciennes racines G02/G03
terminées, tests du résolveur sous race Go 1.27.0 et dans un companion Go
1.22.12, reconstruction privée du **candidat 7 à 1/8192**, audit puis contrôle
CLI/transitions sur les quatre hôtes. Les caches Go de cette construction sont
privés et seront archivés après les validations terminées.

La séquence de reconstruction/validation termine avec succès. Les deux race
detectors passent, ainsi que **89 commandes CLI / 56 comparaisons de transitions**
sur JS, Go séquentiel/parallèle et Rust : mêmes TAST, objets métier/préparés,
mêmes Go exécutés. L'audit retrouve uniquement `tools/ffi-gen/packages.go` et
son WASM reconstruit parmi les changements depuis le candidat 6.

Les artefacts du candidat 7 sont identifiés par SHA-256 :

- Go natif : `719c61d98117bf9098f3f8d97efe55cb620d25af5046cf68caf24b65b7b5836d` ;
- JS : `199cbc2bba023222f90ad86c025531e514ad5f8856bf18494f40e999881c95ed` ;
- Rust : `aa64a8939964b736ce2736b022d2925d8ed969eb110782b1c938c056257a4983` ;
- WASM : `d87a0cad6a4ca7dcbcd67d07ccb72283eb5150b7aa17a6cef619573b5e935d3e`.

Le contrôle froid suivant utilise `measure-cold-complete-7.py` :
**trois paires de chaînes réellement froides complètes**, avec Spago réel,
sortie frontend absente sauf `go.mod`/`go.sum`, caches backend/hôte/runner vides,
puis compilation Go et oracle. Il charge les mêmes 15 fonctions et les mêmes
quatre instructions d'environnement du harnais figé. Le budget backend/Go/chaîne
reste **+5 % séparément** ; les critères de stockage restent ceux de G. Les
résultats du candidat 7 formeront leur propre jeu, avec archivage après chaque
paire et une réserve préalable de 16 GiB. Les autres scénarios et la livraison
G06 restent à valider pour ce candidat.

`continue-cold-complete-7.py` est lancé sous `caffeinate`. Il vérifie d'abord les
sorties encore publiées de quatre anciennes racines, puis archive leurs objets
et lance les paires si la réserve de 16 GiB est atteinte. La racine du vieux
probe `compact-retry` conserve une invocation interrompue par SIGPROF ; son
dernier build réussi et son inventaire publié doivent être exacts avant
archivage. Les journaux de cette interruption sont conservés.

#### Candidat 7 : froid complet validé sur trois paires

Le contrôle termine avec **six chaînes / six oracles passants**, les mêmes TAST
et les mêmes Go que la référence de chaque layout. L'archivage après chaque paire
est restauré/vérifié, et la couverture réunit les six échantillons. Tous les
critères temporels et de stockage de ce scénario passent.

| Froid complet, candidat 7 | Référence | Candidat 7 | Variation |
| --- | ---: | ---: | ---: |
| Backend | 115,177 s | 117,740 s | **+2,23 %** |
| Compilation Go | 239,500 s | 128,850 s | **−46,20 %** |
| Chaîne jusqu'à `exec` | 385,109 s | 276,664 s | **−28,16 %** |

Les trois paires backend donnent **+4,10 / +3,45 / +1,95 %**, toutes sous +5 %.
Les plages sont **113,103–115,293 / 117,543–119,148 s**. La marge médiane jusqu'au
budget est de **3,196 s**. Ce nouveau contrôle apparié valide le candidat 7 pour
le froid complet ; la différence de temps avec la campagne 6 ne constitue pas
à elle seule une mesure du gain du correctif entre candidats.

Le GOCACHE total reste proche de son budget, à **+9,9914 %**. Les objets sont à
**+1,0836 %**, les exécutables à **+4,8460 %**, le Go généré à **−8,0824 %**. Le
RSS backend médian augmente de **10,45 %**, de **10 562 027 520 à 11 665 506 304
octets**, et celui du compilateur Go baisse de **47,01 %**. Ces indicateurs
mémoire restent secondaires selon la priorité utilisateur.

Ce résultat porte sur le seul scénario froid complet. Le candidat 7 doit encore
passer identique, feuille, partagé, FFI et froid applicatif sur trois paires,
puis les grands corpus quatre hôtes, le bootstrap courant et la revue/reprise
stricte des snapshots. Les **30 autres chaînes** seront propres au candidat 7 ;
les mesures du candidat 6 restent des preuves historiques. G06 reste ouvert.

#### Préparation des cinq scénarios restants du candidat 7

L'autorisation d'archivage est étendue aux targets Cargo terminés **5–7** et aux
sorties frontend de campagnes terminées. Les trois targets représentent
**7 626 fichiers / 3 417 704 410 octets logiques**, dédupliqués en **854 523 904
octets de base d'archive**. Les binaires JS/Go/Rust, le WASM et les preuves des
constructions réussies sont revérifiés avant/après.

Pour les sorties frontend, un adaptateur par liste JSON alimente le codec
d'archivage existant. Son aller-retour vérifie octets, modes, dates et répertoires
vides. La sélection porte sur les seuls répertoires de modules contenant
`corefn.json`, `docs.json` et `externs.cbor` ; les onze modules contenant aussi
une entrée Go restent sur place. Sur **26 échantillons terminés**, les **70 018
répertoires / 210 054 fichiers / 8 134 448 452 octets logiques** sont conservés
dans une base dédupliquée de **36 966 400 octets**, plus index et contrôles.
Chaque fichier est restauré et vérifié avant retrait. Les inventaires TAST et
Go sont contrôlés, ainsi que les inodes des verrous, les journaux de propriété,
les pointeurs courants et les résultats originaux. L'espace libre observé passe
de **12 441 772 032 à 21 148 631 040 octets** pendant ces deux opérations ; cette
variation hôte ne constitue ni un gain temporel ni une mesure de stockage du
candidat. Le complément archive les objets de **12 racines terminées** des
anciennes paires de stockage `publicpool2`/`publicsync` : **131 904 fichiers /
6 573 619 140 octets logiques**, avec une base dédupliquée de **100 757 504
octets**. Les 36 commandes snapshot/publication/lookup réussies et les **3 010
Go par racine** sont revérifiés. Les résultats, exécutables, sources Go publiées,
journaux, pointeurs et inodes des verrous sont préservés. L'espace libre observé
passe de **19 509 039 104 à 26 274 283 520 octets** pendant ce complément.

`measure-remaining-cycle-7.py` réutilise les **15 fonctions et quatre instructions
d'environnement** du harnais figé, vérifiées par AST contre le froid du candidat 7.
Il prévoit **30 chaînes**, soit trois paires alternées pour identique, feuille,
partagé, FFI et froid applicatif, puis leur réunion avec les six froides complètes
déjà acquises. Les caches privés partent de la base seulement. L'archivage intervient
après chaque paire terminée, puis sur les caches de travail chauds après les
24 mesures chaudes ; les seeds nécessaires au froid applicatif restent présents.
Réserve de départ : **20 GiB**, plancher des commandes : **8 GiB**.

La revue corrige une collision de noms entre les journaux du lanceur et du codec
d'archivage. Le contrôle sur petits workspaces passe ensuite par le vrai lanceur
figé : archives chaudes/froides, refus après oracle échoué, maintien des caches
encore actifs, contrôle des verrous/journaux et restauration des **16 racines**.
L'essai initial arrêté sur une indentation, avant toute mesure ou archive, est
conservé dans `full-2-harness-rehearsal/` ; le rejeu réussi est dans
`full-2-harness-rehearsal-2/`.

La campagne `full-2/`, lancée en arrière-plan sous `caffeinate`, termine ses
30 chaînes avec oracles passants. Le bilan combiné avec `cold-complete-7/`
réunit **36 chaînes et 39 critères temps/stockage passants**, détaillés en tête
de ce document. Les validations finales réussies sont réunies dans
`final-reference/summary.json` ; la livraison et la pause avant P sont effectives.

**Borne confirmée par l'utilisateur le 16 octobre : terminer G, puis pause avant
P.** Si les critères du candidat 7 passent, la suite couvre les validations
finales et la livraison de G06. Un critère bloquant est présenté avant toute
nouvelle série de candidats ; la campagne n'entraîne pas automatiquement une
nouvelle recherche d'optimisation. Le bilan de livraison doit garder visibles
les coûts du backend seul et de la chaîne, ainsi que la régression froide
accumulée depuis le point de départ, au-delà de la comparaison immédiate à
`publicsync`.

### Historique du raccordement des exports allégés

L'acquisition G03 extrayait `__.PKGDEF` pour le transport, mais conservait les
archives Go complètes dans GOCACHE. Un probe repart d'un cache vide et utilise
`-gcflags=all=-N -l -linkobj=<os.DevNull>` pour cette seule acquisition. Les
**250 exports** occupent **5 271 792 octets** ; le GOCACHE du probe contient
**7 373 709 octets**. Le résultat du résolveur, plan et inventaire compris, ainsi
que les **131 549 639 octets de sources de packages** sont byte-exacts au témoin
1/128. Les entrées et le runtime ne font pas partie de ce dernier total. Les
**9,003 s** observées pour l'acquisition sont un diagnostic unique, pas un gain
temporel en paires ; ce cache n'inclut pas les autres phases du backend.

Le raccordement dans `tools/build-cache/packages.go` est limité à un `GOFLAGS`
vide. La suite du cache passe sous Go **1.27.0 et 1.22.12** avec détection de races,
dont un build
exécutable normal utilisant ensuite le même GOCACHE et la conservation d'une
erreur provenant d'options explicites du demandeur. Les bridges et les mesures du
candidat complet sont ensuite validés dans la clôture du candidat 7.

## Protocole du cycle

La campagne `gopurs-go-packages-g06-20261014/measure-cycle.py` reprend le corpus
b8x isolé, les quatre états TAST, les mutations et le compilateur `publicsync`
identifié par v3/06. Chaque générateur conserve sa racine et son journal de
propriété. Spago et les téléchargements Go sont figés ; les invocations sont
hors ligne. Une édition part d'un vrai retour à la base dans cette même racine.

Chaque mesure reçoit un GOCACHE privé et le binaire du runner, préchauffés
uniquement avec la base. Le runner réel utilise Go 1.26.8 linux/arm64 dans l'image
historique, avec **GOMAXPROCS=2 / -p 1 / GOMEMLIMIT=12GiB**. Le backend garde les
huit workers et les ressources hôte de G00. Le relevé fonctionnel G03/G05,
souvent concurrent et limité à deux processeurs, n'est pas une référence temporelle.

Le chronomètre direct part de Spago et s'arrête au marqueur précédant `exec`.
Le relevé conserve frontend hors backend, gopurs, tidy, verrou, Go, actions de
compilation/lien, RSS et volumes logiques. Le RSS backend inclut les processus
enfants de son invocation ; celui du compilateur Go est le maximum des `VmHWM`
échantillonnés à une seconde. Les oracles sont construits et exécutés hors de la
chaîne mesurée. La suite métier/PostgreSQL de `Test.Main` n'est pas dans ce temps,
conformément au périmètre G00/v3.

À partir du pilote 2, les volumes sont relevés **après l'oracle**. Les caches Go
des exports hôte et du runner sont privés, clonés depuis leur état de base et
publiés séparément ainsi qu'en total dans la catégorie GOCACHE. Le cache utilisé
pour le retour à la base n'est pas celui de l'échantillon. Le seed du runner ne
contient que `Test.Main` sur la base, comme en G00 ; l'oracle est construit ensuite
dans chaque échantillon, toujours au même chemin conteneur.

Un passage exploratoire à une paire peut détecter un obstacle avant la campagne
complète. Il est explicitement marqué `exploratory` et ne valide aucun seuil.
La décision finale requiert au moins trois paires alternées, les contrôles froids,
la parité des hôtes et les replays stricts après revue des snapshots.

### Archivage entre scénarios

L'utilisateur autorise aussi l'archivage sans perte des répertoires `objects`
des caches gopurs de campagnes et d'échantillons terminés. Le verrou conserve son
inode et le journal de propriété reste à son emplacement. L'option privée
`G06_ARCHIVE_COMPLETED=1` archive les GOCACHE hôte/runner des échantillons après
**toutes les paires et tous les oracles d'un scénario**. Pour les seuls scénarios
froids, leurs objets gopurs sont aussi archivés : ces racines ne servent à aucun
retour à la base ultérieur. Les seeds et les racines chaudes encore utilisées
restent disponibles. Les volumes publiés sont ceux relevés avant l'archivage.

La restauration de chaque fichier est vérifiée avant retrait. L'archive
revérifie les nombres de fichiers et octets mesurés, puis les empreintes des
journaux/pointeurs et les inodes des verrous. Un test sur petits workspaces
exerce les scénarios chaud/froid, l'arrêt en cas d'oracle échoué et la
conservation d'un cache chaud actif. Il compare aussi les AST des fonctions
chronométrées et de la boucle de mesures avec leur état antérieur : ils sont
identiques après retrait de l'appel d'entretien entre scénarios. Les deux
échecs initiaux du harnais, dus à l'indentation avant tout lancement de mesure,
sont conservés, puis le troisième contrôle passe.

## Preuves

Répertoire temporaire OpenCode `gopurs-go-packages-g06-20261014/` :

- `final-reference/summary.json` : **acceptation de livraison**, artefacts
  installés, inventaire final de 295 fichiers source/outils, snapshots, empreintes
  des preuves/documents, **G 7/7, P 0/7 et pause avant P00** ;
- `full-2/combined-summary.json`, `full-2/combined/analysis-time-priority.json` :
  **36 chaînes / 36 oracles**, **39 critères bloquants passants** ;
  `full-2/archive-coverage.json` couvre les 30 nouvelles mesures et renvoie à la
  couverture froide initiale. `docs/build-go-cycle-g06.csv` publie les 36 lignes,
  SHA-256 `d47ec36b85396dfb7fafec631454ada5f0dc426294b8f59a6b657a5a4ebc650b` ;
- `validate-final-private-7.py`, `final-private-7-launch.log`, `final-private-7/` :
  validation privée réussie : **huit comparaisons de grands corpus / 28 commandes**,
  b8x **2 704 TAST / 27 Go**, bootstrap **512 TAST / 23 Go**, hits sans réécriture
  et reconstructions sans compilation/lien ; bootstrap émis éprouvé comme
  compilateur, puis **neuf fixtures / 207 commandes** dans les quatre hôtes ;
- `review-layout-7.go`, `review-layout-7-self-test.log`,
  `review-layout-7-reference-check.json` : comparateur structurel préparé pour
  la revue des déplacements de packages ; cinq contrôles positifs/négatifs et
  les **612 déclarations des neuf références existantes** passent.
  `snapshot-review-7.json` et `snapshot-7-*.diff` conservent la revue réussie des
  neuf candidats, des chemins/imports et des déclarations avant mise à jour ;
- `delivery-7/installation.json`, `delivery-7/before/` : installation atomique des
  artefacts exacts et conservation des anciens binaires/WASM/snapshots ; audit
  initial de **295 fichiers source/outils**. `delivery-7/performance-audit.json`
  recalcule les **39 critères** depuis les lignes brutes et vérifie le CSV ;
- `delivery-7/fixture-adjustments/`, `delivery-7/replays/` : échecs conservés de
  deux fixtures à 512 fillers qui n'ouvraient plus de frontière à 1/8192, et du
  passage accidentel du vrai compilateur aux tests de faux compilateur. Les
  fixtures ajoutent chacune deux identités d'ancre fixes, avec assertions
  conservées ; `validate-delivery-7.py` reprend dans `replays-2/` avec environnement
  corrigé. Ces deux fichiers de tests sont les seules différences source/outils
  admises depuis l'installation ; les artefacts mesurés restent exacts ;
- `delivery-7/replays-2/` : **82 tests ciblés**, intégrité/récupération du cache,
  transitions du pilote dans les quatre hôtes, mise à jour des neuf snapshots
  revus puis **207 commandes strictes**, et **89 commandes CLI** passent.
  Le bootstrap suivant s'arrête sur une erreur de sélection du harnais : Spago
  0.93.45 au lieu du 1.0.3 figé et une racine d'une profondeur différente ;
  `delivery-7/bootstrap-preflight-review/` vérifie les 512 TAST et ne retrouve
  que **231 chemins relatifs résolvant aux mêmes fichiers**, plus la version
  Spago et ses spans dans `Spago.Generated.BuildInfo`. L'échec est conservé ;
  la reprise `replays-3/` sélectionne le Spago vérifié et une racine de même
  profondeur, en conservant les assertions d'égalité brute TAST/Go ;
- `current-bootstrap-7/` : reconstruction depuis les sources courantes, **512 TAST
  et 23 Go byte-exacts**, binaire SHA-256
  `8d276f3fba42a8ea4c88ddfbda04640e5eedd44daafff3b70fbc11839b631ab4`, identique
  au bootstrap émis validé auparavant ; contrôle CLI séquentiel/parallèle passant.
  Les runners macOS et Linux réussissent chacun dix contrôles. Le premier runner Linux en réussit neuf
  puis détecte l'absence de Go dans le `PATH` du harnais Docker ; sa reprise
  conserve l'échec et utilise `/usr/local/go/bin:/usr/local/bin:/usr/bin:/bin`,
  comme le protocole figé. `delivery-7/replays-3/runner-linux-retry.log` conserve
  les dix succès ; `delivery-7/replays-3/summary.json` réunit les contrôles réussis,
  avec les preuves antérieures et les empreintes des artefacts inchangées ;
- `delivery-7-reserve-review-2/`, `completed-delivery-7-storage-objects/` : huit
  anciens probes terminés, **87 936 fichiers / 4 273 351 792 octets logiques**
  restaurés/vérifiés, base d'archive **93 941 760 octets** ; journaux, pointeurs,
  verrous, binaires et Go publiés conservés. Le premier contrôle arrêté avant
  archivage faute de journal de commande sur un probe plus ancien est conservé ;
- `before/`, `before.json` : état avant le changement d'alias ;
- `runtime-alias-before.log`, `runtime-alias-before-2.log`,
  `runtime-alias-after-1.log` : harnais corrigé, échec antérieur et suite passée ;
- `candidate-*/`, `build-candidate.py` : artefacts et sources figés, Rust construit
  dans un target neuf et chemin de son archive native contrôlé ;
- `review-runtime-alias.go` : comparaison des fichiers complets de snapshots
  après normalisation syntaxique du seul import runtime ;
- `measure-cycle.py`, `backend-timer.py` : commandes du parcours et compteurs ;
- `global-cache-clean.json` : purge autorisée du seul cache Go global après les
  tests actifs, avec **8,75 → 20,26 Go** d'espace libre mesuré en octets décimaux.
- `pilot-1/` : états, commandes, sorties, mesures et analyse du candidat rejeté ;
- `pilot-2/` : seconde paire exploratoire, caches séparés, résultats après oracle,
  rejet détaillé et `archive-analysis.json` ; `pilot-2-archive-analysis.log`
  conserve une erreur initiale du lecteur d'actions (`Cmd: null`), corrigée avant
  le rapport `pilot-2-archive-analysis-2.log` ;
- `memory-diagnostic-2/`, `memory-phases-2/` : allocations, GC, RSS par phase et
  vérifications des sorties/inodes ; `memory-boundary-before/` conserve le code
  et le dépassement avant l'essai de restitution mémoire ;
- `memory-boundary-after/`, `candidate-memory-boundary/` : essai natif rejeté,
  code, artefact et décision ; `intern-identifiers-before/` fixe les seuils du
  probe suivant ; `intern-identifiers-after/` conserve son rejeu exact et son
  rejet ; `build-native-probe.py` identifie les variantes natives privées ;
- `density-probe/`, `probe-density.py`, `emit-density.py` : sources G05 vérifiées,
  témoin 1/128 exact, variantes privées et compilations/oracles ;
- `thin-exports-before/`, `thin-exports-probe/` : code antérieur, protocole,
  exports et comparaison exacte ; `thin-exports-cache-tests.log` et
  `thin-exports-cache-go122.log` : suites race Go 1.27.0 et 1.22.12 ;
- `candidate-3-provenance.json`, `candidate-3-diff-audit.json` : reconstruction
  des quatre hôtes et inventaire limité des changements du candidat privé ;
- `functional-private-3/`, `candidate-3-resolver-tests.log` : contrôles CLI,
  transitions et résolveur avant le pilote 3 ; `pilot-3/` et `pilot-3.log` :
  campagne exploratoire achevée, analyses historique et priorité temps ;
- `pilot-3/object-reachability.json`, `package-retention-before/`,
  `package-retention-before.log`, `package-retention-after.log`,
  `package-retention-go122.log`, `package-retention-concurrency*.log` : portée
  des objets, code antérieur, échec du reproducteur puis suites race passées ;
- `package-cpu-3/`, `package-gc-pairs-3/` : profil CPU et essai GC rejeté, trois
  paires alternées, programmes figés et vérifications exactes ;
- `render-buffers-before/`, `compare-render-buffers.py`,
  `render-buffers-pairs/`, `render-buffers-resolver-tests.log` : code antérieur,
  seuil préalable, trois paires locales et suite du résolveur après réduction
  des copies ;
- `trimmed-exports-before/`, `trimmed-exports-probe/` : code antérieur,
  protocole, exportation sans chemins locaux et égalité des sorties ;
  `trimmed-exports-tests*.log` : suites cache/race des deux versions de Go ;
- `build-candidate-4.log`, `candidate-4-provenance.json`,
  `candidate-4-diff-audit.json` : reconstruction et audit du candidat suivant ;
  `functional-private-4/` : CLI et transitions sur les quatre hôtes ;
  `measure-cold-backend.py`, `cold-backend-4/` : trois paires froides, inventaires
  exacts, rejet et relevé des phases ;
- `source-gzip-reuse-before/`, `compare-source-gzip.py`,
  `source-gzip-reuse-tests*.log`, `source-gzip-reuse-pairs/` : état antérieur,
  protocole local, tests et mesures de réutilisation des compresseurs ;
  `source-gzip-reuse-rejected/` conserve le code, les tests et le rejet ;
- `parallel-exports-before/`, `compare-parallel-exports.py`,
  `parallel-exports-pairs/` : code antérieur, seuil fixé et six acquisitions
  froides avec comparaison exacte des exports et des sorties du résolveur ;
  `parallel-exports-tests*.log` : suites cache/race des deux versions de Go ;
- `build-candidate-5.log`, `candidate-5-provenance.json`,
  `candidate-5-diff-audit.json` : reconstruction et audit de la variante à quatre
  processus ; `functional-private-5/` : validation multi-hôtes passée ;
  `cold-backend-5/` : trois paires froides, inventaires exacts et rejet ;
  `profile-cold-backend.py`, `cold-profile-5/` : profils CPU/allocations/GC
  réellement froids, diagnostic uniquement, TAST/Go exacts ;
  `cold-profile-5/gc-phase-analysis-2.json` corrige le premier découpage, qui
  prenait une ligne de progression pour la fin de l'émission ; les deux sont
  conservés. `probe-entry-capture.py`, `entry-capture-probe/`,
  `measure-entry-capture.py`, `entry-capture-pairs/` : seuil préalable, variante
  native privée et paires de capture des seules entrées ;
- `compare-export-cap.py`, `export-cap-8-pairs/` : code antérieur, seuil fixé,
  trois paires froides à quatre/huit processus, exports et Go exacts ;
  `validate-export-cap-8.py`, `export-cap-8-validation.log` : séquence de
  validation terminée du candidat 6 ; `candidate-6-diff-audit.json`,
  `functional-private-6/`, `cold-backend-6/` : audit, parités fonctionnelles et
  contrôle froid à trois paires ;
  `completed-entry-capture-objects/` : archivage terminé des objets des six
  échantillons du probe rejeté, **66 036 fichiers restaurés/vérifiés /
  3 327 669 033 octets logiques → 107 597 824 octets d'archive**, avec journaux,
  pointeurs et inodes des verrous vérifiés. **17 103 654 912 octets libres**
  sont observés après retrait ;
- `completed-cold-6-objects/` : objets des six échantillons froids terminés,
  **65 994 fichiers restaurés/vérifiés / 3 309 774 072 octets logiques →
  123 506 688 octets d'archive** ; journaux, pointeurs et inodes des verrous
  vérifiés ; **14 682 370 048 octets libres** observés après retrait ;
- `archive-reviewed-objects.py`, `reviewed-cold-object-roots.json`,
  `completed-historical-cold-objects/` : quatorze anciennes racines terminées,
  dont les deux diagnostics froids 5, archivées avec restauration vérifiée :
  **129 026 fichiers / 6 519 532 513 octets logiques → 143 634 432 octets
  d'archive**. Le premier contrôle, arrêté avant archivage, supposait une édition
  feuille pour chaque ancienne paire ; la reprise reconnaît aussi partagé/FFI,
  après examen des preuves. Le contrôle des journaux/pointeurs/inodes passe ;
- `global-cache-clean-6.json` : purge autorisée du seul GOCACHE global
  régénérable après les constructions/tests terminés, **1 322 909 336 octets
  logiques** ; **21 926 551 552 octets libres** observés après purge ;
- `full-1-preflight.json`, `run-full-cycle-1.py` : empreintes du candidat 6 et
  harnais d'archivage vérifié, campagne complète de 36 chaînes dans `full-1/`,
  interrompue à 28 chaînes ;
  `full-1-resume-preparation/`, `prepare-full-1-resume.py`,
  `resume-full-cycle-1.py` : copie des 28 résultats et commandes, preuve du
  SIGTERM, contrôle des événements d'alimentation et reprise des huit mesures
  restantes ; `continue-full-cycle-1.py`, `full-1-resume-launch.log` : séquence
  d'entretien/reprise lancée ; `archive-full-1-cold-rows.py` : archivage vérifié
  des échantillons froids achevés de part et d'autre de l'interruption ;
- `full-1-resume-2-preparation/`, `prepare-full-1-resume-2.py`,
  `resume-full-cycle-2.py`, `continue-full-cycle-2.py` : arrêt sur le garde-fou
  disque, sauvegarde des 33 chaînes, mise à part du côté candidat de la paire
  interrompue et deuxième reprise. `archive-full-1-complete-rows.py` conserve
  les échantillons froids complets terminés avant reprise puis après chaque
   paire ; les mesures et l'échantillon écarté sont tous comptabilisés ;
- `full-1/resume-2-summary.json`, `full-1/analysis-time-priority.json` : bilan des
  **36 chaînes / 36 oracles**, seul échec temporel/stockage sur le backend froid
  complet ; `full-1/archive-cold-complete-coverage.json` vérifie les sept
  échantillons physiques. Ses trois archives couvrent **12 301 777 952 octets
  logiques / 137 039 fichiers restaurés/vérifiés → 1 399 865 344 octets d'archive** ;
- `extend-cold-cycle-6.py`, `archive-cold-extension-pair.py`,
  `cold-extension-6-script-check.json`, `cold-extension-6-script-check-2.json` : extension fixée à trois paires
  supplémentaires ; fonctions/environnement figés et adaptation de la seule
  racine de campagne dans l'archiveur contrôlés par AST ; le second contrôle
  fige la réserve de 16 GiB et les empreintes des scripts lancés.
  `cold-extension-6-launch.log` conserve le lancement ;
- `cold-extension-6/summary.json`, `cold-extension-6/combined/` : bilan cumulé
  de **42 chaînes / 42 oracles** et refus froid à six paires ;
  `original-proof-hashes.json`, `archive-coverage.json` et `phase-observations.json`
  conservent la vérification des preuves initiales, la couverture et les phases.
  Les trois archives supplémentaires représentent **116 895 fichiers
  restaurés/vérifiés / 10 493 267 550 octets logiques → 1 397 006 336 octets
  d'archive** ; **12 417 388 544 octets libres** à la fin du harnais ;
- `render-single-pass-probe/`, `measure-single-pass-render.py` : code, seuil
  préalable, tests du résolveur et trois paires exactes du rendu en une passe ;
  seuil local passé, validation du candidat complet à suivre ;
- `validate-single-pass-7.py`, `single-pass-7-validation-launch.log` : séquence
  de reconstruction/validation du candidat 7 terminée ;
  `single-pass-7-validation-summary.json`, `candidate-7-diff-audit.json` et
  `functional-private-7/` en conservent les résultats et empreintes ;
  `review-candidate-7-object-roots.py`
  vérifie les commandes et les preuves des neuf anciennes racines ;
  `measure-cold-complete-7.py`, `archive-cold-7-pair.py` : contrôle froid complet
  préparé avec helpers figés et archivage par paire ;
- `completed-candidate-7-historical-objects/` : **47 680 fichiers restaurés/vérifiés /
  3 108 661 484 octets logiques → 155 734 016 octets d'archive**, journaux,
  pointeurs et inodes préservés pour les neuf racines G02/G03 terminées ;
- `candidate-7-build-cache-archived/` : cache privé de la construction et des
  validations achevées, **7 003 fichiers restaurés/vérifiés / 1 250 538 919 octets
  logiques → 278 642 688 octets d'archive** ; **12 509 724 672 octets libres**
  au bilan de validation ;
- `continue-cold-complete-7.py`, `cold-complete-7-launch.log`,
  `review-cold-7-historical-objects.py` : séquence de contrôle des anciens
  inventaires publiés, entretien des objets terminés et paires froides du
  candidat 7 ;
- `completed-cold-7-historical-objects/` : objets des quatre racines anciennes
  dont les Go publiés ont été recontrôlés, **35 672 fichiers restaurés/vérifiés /
  5 301 458 737 octets logiques → 324 653 056 octets d'archive** ; journaux,
  pointeurs et inodes conservés ; **17 595 547 648 octets libres** après retrait ;
- `cold-complete-7/summary.json`, `analysis-time-priority.json`,
  `archive-coverage.json` : contrôle froid complet passant, six chaînes exactes
  et couverture des six échantillons. Les trois archives représentent
  **116 895 fichiers restaurés/vérifiés / 10 493 228 169 octets logiques →
  1 397 174 272 octets d'archive** ; **12 602 638 336 octets libres** au bilan ;
- `archive-roots-file.py`, `frontend-archive-rehearsal/`, `full-2-space-steps.json` :
  transport des racines par liste JSON, aller-retour vérifié et entretien
  séquentiel Cargo/frontend ; `cargo-targets-5-7/` et
  `completed-cold-frontends-controls/summary.json` conservent les volumes et
  contrôles des trois targets et des 26 sorties frontend terminées ;
- `prepare-full-2-storage-reserve.py`, `full-2-storage-reserve-review/`,
  `completed-full-2-storage-objects/` : revue des douze racines historiques,
  restauration vérifiée de leurs objets et contrôles avant/après ;
- `check-full-2-harness.py`, `full-2-harness-rehearsal/failure.json`,
  `full-2-harness-rehearsal-2/summary.json` : échec initial conservé et contrôle
  réussi des helpers figés, du vrai lanceur et des deux archiveurs ;
- `measure-remaining-cycle-7.py`, `archive-full-2-pair.py`,
  `archive-full-2-warm-work.py`, `run-remaining-cycle-7.py` : continuation des
  cinq scénarios avec le candidat 7 ; `remaining-cycle-7-launch.log`,
  `remaining-cycle-7.log` et `full-2/` conservent la campagne lancée ;
- `reviewed-extension-object-roots.json`, `completed-extension-historical-objects/` :
  **22 anciens probes terminés**, **233 658 fichiers restaurés/vérifiés /
  11 988 264 506 octets logiques → 194 080 768 octets d'archive**, avec journaux,
  pointeurs et inodes des verrous vérifiés ; **13 214 932 992 octets libres**
  observés après retrait ;
- `review-extension-final-probe-objects.py`, `reviewed-extension-final-probes/`,
  `completed-extension-final-probe-objects/` : **19 probes historiques complets /
  190 commandes réussies** vérifiés avant archivage de leurs objets,
  **201 004 fichiers restaurés/vérifiés / 10 449 700 421 octets logiques →
  203 436 032 octets d'archive**. Journaux, pointeurs et inodes vérifiés ;
  **17 506 934 784 octets libres** observés après retrait ;
- `full-1/archives/cold-application-before-resume/` : **89 366 fichiers
  restaurés/vérifiés / 6 999 633 482 octets logiques → 471 601 152 octets
  d'archive** ; après reprise, `cold-application-after-resume/` conserve
  **3 499 889 198 octets logiques → 467 488 768 octets d'archive**.
  `full-1/archive-cold-application-coverage.json` vérifie la couverture des six
  mesures et les volumes bruts ;
- `completed-full-1-warm-resources/` : seeds/caches de retour à la base et objets
  des seuls scénarios chauds terminés, **42 372 fichiers restaurés/vérifiés /
  5 897 221 246 octets logiques → 469 946 368 octets d'archive**. Les seeds hôte
  et externes nécessaires au froid restent disponibles ; journaux/pointeurs et
  inodes des verrous sont vérifiés ;
- `reviewed-resume-object-roots.json`, `completed-resume-historical-objects/` :
  vingt anciens probes terminés, **223 069 fichiers restaurés/vérifiés /
  11 511 914 130 octets logiques → 193 658 880 octets d'archive**, avec journaux,
  pointeurs et inodes des verrous vérifiés. Le volume logique archivé ne mesure
  pas l'espace physique libéré ; le relevé après retrait est de
  **11 727 245 312 octets libres**, d'où l'entretien supplémentaire avant reprise ;
- `reviewed-second-object-roots.json`, `completed-historical-second-objects/` :
  onze anciennes racines terminées, **150 806 fichiers restaurés/vérifiés /
  5 176 699 659 octets logiques → 236 109 824 octets d'archive**, avec journaux,
  pointeurs et inodes des verrous vérifiés. **26 516 660 224 octets libres**
  sont observés immédiatement après retrait ;
- `global-cache-clean-5.json` : purge autorisée du cache Go global après les
  constructions/tests achevés, **1 327 854 209 octets logiques**, avec
  **13 023 977 472 → 14 380 912 640 octets libres** avant les mesures froides ;
- `completed-archive-policy.json`, `measure-cycle-before-completed-archives.py.saved`,
  `check-completed-archive-harness.py`, `completed-archive-harness-check*.log`,
  `completed-archive-harness-check-3/` : politique et contrôle de l'entretien
  entre scénarios ; `archive-completed-objects.py`, `completed-objects-1/` :
  archivage terminé de **18 racines**, sous leurs verrous d'origine,
  **238 678 fichiers restaurés/vérifiés / 10 175 907 138 octets logiques →
  225 255 424 octets d'archive**. Les journaux, pointeurs et inodes des verrous
  sont vérifiés ; **20 562 853 888 octets libres** sont observés après retrait.
  Les anciens chemins `objects` portent des pointeurs de restauration ;
- `global-cache-clean-4.json` : purge autorisée du cache Go global régénérable
  après la reconstruction et les tests, **1 384 608 554 octets logiques**, avec
  **14 083 264 512 → 15 500 914 688 octets libres** avant les paires froides ;
- `global-cache-clean-3.json` : purge autorisée du cache Go global régénérable
  après ces tests, **2 316 884 516 octets logiques**, avec **20 908 732 416 →
  23 267 434 496 octets libres** avant le pilote ;
- `completed-probes-go-caches/` : huit GOCACHE de diagnostics achevés,
  **52 619 fichiers reconstruits/vérifiés / 2 221 628 767 octets logiques**, dans
  **264 724 480 octets d'archive** ;
- `pilot-1-go-caches/` et `archive-go-caches.py` : archivage sans perte autorisé,
  avec déduplication, **25 374 fichiers reconstruits et vérifiés**, soit
  **7 052 084 118 octets logiques** dans **681 504 768 octets d'archive**. Les
  volumes d'acceptation restent les volumes bruts mesurés, avant archivage.
- `pilot-2-go-caches/` : même archivage des huit caches du runner terminés,
  **26 055 fichiers reconstruits et vérifiés**, soit **12 305 421 514 octets
  logiques** dans **625 188 864 octets d'archive** ; les caches hôte restent
  disponibles pour les diagnostics suivants.
- `pilot-3-go-caches/` : douze caches privés du runner achevé archivés et
  restaurés/vérifiés, **20 440 fichiers / 18 997 079 879 octets logiques →
  1 014 153 216 octets d'archive**. Chaque ancien chemin contient un pointeur de
  restauration ; les preuves et les mesures brutes sont conservées. Le relevé
  après retrait indique **17 684 180 992 octets libres**.
- `g00-g01-g05-go-caches/` : dix caches privés des lots clos G00/G01/G05,
  **19 327 fichiers restaurés/vérifiés / 10 493 652 809 octets logiques →
  212 881 408 octets d'archive**. Les caches d'origine portent leurs pointeurs
  de restauration ; le relevé après retrait donne **12 899 565 568 octets libres**.
- `cargo-targets-archive-inputs.json`, `cargo-targets-1-4/` : extension de
  l'archivage autorisée explicitement par l'utilisateur aux quatre targets Cargo
  terminés. Le même format sans perte conserve **10 168 fichiers restaurés et
  vérifiés / 4 556 780 824 octets logiques**, dans **1 127 735 296 octets
  d'archive**. Les binaires des candidats sont disponibles séparément ; les
  anciens chemins `rust-{1,2,3,4}/target` portent un pointeur, restaurable avec
  `archive-go-caches.py restore` au chemin original. Le relevé après retrait
  indique **15 073 353 728 octets libres**.
- `v3-lot06-resume-go-caches/` : même archivage des dix caches privés terminés de
  la reprise v3/06, **16 830 fichiers vérifiés / 19 340 958 118 octets logiques**,
  dans **477 765 632 octets d'archive**. Les mesures et les autres artefacts v3
  restent à leurs chemins d'origine ; chaque ancien GOCACHE porte le pointeur
  nécessaire à sa restauration.
- `v3-original-go-caches/` : les dix-neuf caches du runner de la campagne v3/06
  initiale, clôturée par la reprise ci-dessus, sont aussi archivés sans perte :
  **31 928 fichiers reconstruits/vérifiés / 31 565 717 888 octets logiques →
  477 827 072 octets d'archive**. Le relevé après retrait indique
  **21 218 127 872 octets libres**. Les exécutables, commandes et inventaires
  restent dans la campagne initiale.
