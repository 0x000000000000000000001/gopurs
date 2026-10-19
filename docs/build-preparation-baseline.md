# Préparation incrémentale — P00

**P00 terminé le 19 octobre 2026 ; P : 1/7 (audit), G : 7/7.**
Le bilan recommande de **revoir le périmètre avant P01** : la persistance des
états complets sondés ne satisfait pas le minimum de gain, et son coût froid est
trop élevé. Une représentation plus compacte reste à évaluer, sans gain acquis.
Une [étude resserrée de P01](build-preparation-compact.md) a ensuite été
autorisée et terminée le 19 octobre : **C01 rejeté en l'état**, publication
améliorée mais restauration et surcoût froid au-delà des budgets. Une nouvelle
décision précède la suite. Les repères restent
**~30 s de backend après modification / ~60 s à froid**.

Résultats exportés : [36 passages bruts](build-preparation-p00-timings.csv),
[coûts et modèles](build-preparation-p00-costs.json),
[empreintes et validations](build-preparation-p00-evidence.json).

**Reprise du 19 octobre après coupure :** le répertoire temporaire de P00 est
absent, comme les anciennes preuves G06 l'étaient au démarrage. Les scripts et
ce document sont présents. Les 36 résultats du 16 avaient passé la comparaison
byte-exacte, mais leurs froids étaient trop dispersés ; ils ne valent pas une
référence temporelle acceptée. Les valeurs rappelées ci-dessous proviennent des
sorties d'outils conservées dans la conversation, sans journaux bruts encore
disponibles.

La reprise écrit ses preuves dans **`htdocs/.gopurs-evidence/p00-20261019/`**,
hors des répertoires temporaires. Les quatre artefacts ont les mêmes SHA-256
qu'au démarrage. Le b8x de travail est passé à `67b6799`, mais le corpus privé
reprend `e6a3469` par archive Git ; ses caches viennent explicitement de
`run/bak/go/`. Les 2 704 CoreFn actuellement disponibles dans cette sortie sont
non typés (`0.15.15`) : ils sont inventoriés comme tels ; les TAST privés ont été
régénérés avec le fork typé `0.15.16`. Leur contenu historique ne sert pas
d'oracle de préparation.

## Provenance et référence à reconstituer

Le contrôle initial retrouve les sources gopurs propres au commit `0d597f5`,
PBO au commit `40d759e` et b8x au commit `e6a3469`. Les répertoires temporaires
des anciennes campagnes ne sont plus présents au chemin enregistré. Le CSV
G06 versionné conserve ses 36 échantillons, mais pas leur ventilation interne.

Les empreintes JS, Rust et WASM correspondent aux artefacts G06 documentés.
Le natif installé est différent :

| Artefact | SHA-256 au début de P00 |
| --- | --- |
| Natif | `15dcf6904f8c75f67bda248e74d6e0a7758ca80442319ea641631da9c98b1072` |
| JS | `199cbc2bba023222f90ad86c025531e514ad5f8856bf18494f40e999881c95ed` |
| Rust | `aa64a8939964b736ce2736b022d2925d8ed969eb110782b1c938c056257a4983` |
| WASM | `d87a0cad6a4ca7dcbcd67d07ccb72283eb5150b7aa17a6cef619573b5e935d3e` |

Au démarrage du 16, le corpus b8x disponible contient **2 704 TAST / 755 005 types**. L'égalité de
ces comptes avec G06 ne prouve pas l'identité des octets. Une copie privée et
ses inventaires constituent donc une **référence P00 reconstituée**, à valider
entre hôtes ; aucune variation face aux anciens temps G06 ne sera attribuée à
une optimisation de P.

## Protocole fixé avant les mesures

- Figer sources, dépendances, TAST, artefacts, commandes et environnement ;
  conserver les échecs et chaque échantillon. Les sources actives servent
  seulement à produire la copie privée.
- Layout explicite `GOPURS_GO_PACKAGES=1`, préparation/PBO/émission à huit
  workers, pipeline actif, `GOGC=off`, `GOMEMLIMIT=10GiB` pour le natif.
  Runner : image G06, `GOMAXPROCS=2`, `-p 1`, `GOMEMLIMIT=12GiB`.
- Reprendre original, feuille (`pong-lot00`), partagé (suffixe `-lot00` de
  `caseToCamel`), FFI (suffixe `!` de `RemoveAccents`), identique et les deux
  froids. Chaque édition part de la base ; caches privés et préchauffage de la
  base seulement. Aucun profil CPU/allocation dans les mesures d'acceptation.
- Mesurer d'abord les frontières existantes. Les phases imbriquées et le temps
  cumulé des workers ne s'ajoutent pas au total. Conserver le résidu non attribué.
  Une instrumentation plus fine devra reproduire les octets et son surcoût sera
  évalué sur **trois paires alternées** ; au-delà de 5 % d'écart absolu, ses temps
  seront traités comme diagnostics et non comme ventilation représentative du
  binaire témoin.
- Seuil P futur : **−10 % du backend total** sur feuille et partagé ; **−5 %**
  pour chaque optimisation autonome annoncée. Budgets temporels ailleurs :
  +5 %, identique `max(5 %, 0,100 s)`. Le temps est prioritaire ; RSS mesuré,
  sans réintroduire de plafond numérique automatique.
- Chiffrer le stockage/restauration des objets candidats avant de retenir des
  unités persistées. Le bilan P00 proposera un budget de stockage explicite à
  la décision avant P01 ; un coût de cache ne sera jamais soustrait du gain brut
  sans mesure ou hypothèse identifiée.
- Comparer mêmes TAST et sorties byte-exactes en JS, Go séquentiel/parallèle et
  Rust, puis cache/build forcé. Les fixtures ciblent notamment retraits,
  substitutions à cardinalité égale, cycles, fournisseurs et FFI seules.
- Une seule campagne lourde, exécutée en arrière-plan. Publication dans ce
  dépôt du protocole, du harnais et du bilan ; preuves volumineuses privées
  inventoriées séparément.
- À la reprise du 19 : `preparation-observe.py` enregistre les processus et
  leurs liens de parenté toutes les deux secondes, en distinguant le travail du
  harnais d'un compilateur externe. Une compilation lourde déjà présente diffère
  le lancement ; une concurrence détectée pendant le passage invalide son temps
  et arrête la séquence après conservation des logs. Un écart important négatif
  du binaire instrumenté ne suffit pas davantage à démontrer sa représentativité.

## Code actuel et périmètre de l'observation

`Driver.Prepare` recharge tous les TAST et reconstruit les métadonnées source.
`Monomorphization` invalide les usages source, construit l'index des corps et
les barrières, collecte les demandes directes, appelle le point fixe PBO,
filtre l'éligibilité, puis réécrit les modules. `Cache.Modules.initialize`
recalcule les empreintes préparées si la preuve globale `sameTast` échoue.

Les compteurs existants permettent déjà de séparer chargement, préparation,
clés, lookup PBO, optimisation/émission, exports Go, partition et publication.
Le premier harnais observe ces compteurs. **Impact bridges : aucun nouveau
contrat de transport à ce stade.**

### Graphe des lectures — audit du code de référence

| Unité | Entrées effectivement consommées | Condition à prouver entre builds |
| --- | --- | --- |
| Chargement/tri (`Driver.Prepare`, PBO `App`) | Inventaire et contenu TAST, imports, types, annotations, chemins | Ajouts/retraits et ordre canonique inclus ; mêmes diagnostics de décodage |
| Types globaux (`GlobalTypes`) | Noms qualifiés, annotations de bindings puis d'expressions, repli sur résultat d'application, types foreign | Préserver les priorités de lecture, le manque d'annotation et `Any` explicite |
| Constructeurs/élisions (`ConstructorMetadata`) | ADT source originaux, variables, ordre et arité des champs | Une copie spécialisée ou une classe synthétique ne remplace pas l'ADT source |
| Champs/classes (`ClassMetadata`) | Méthodes, superclasses, variables ; tri des labels ; enrichissement des déclarations | Conserver `Any` pour les superclasses synthétiques et l'ordre de fusion |
| Rebox (`ReboxMetadata`) | Tables constructeurs/classes, noms générés | Premier nom trié gagnant, constructeurs prioritaires sur classes ; invalider les collisions |
| ADT pointeurs/enums (`AdtMetadata`) | Déclarations enrichies pour les pointeurs ; originales pour les enums/élisions | Ne pas transformer un dictionnaire vide en enum ; conserver tags et layouts |
| Index des corps/barrières (`Monomorphization`) | Corps et usages invalidés, foreign, forwarders, candidats workers record, types globaux | `Data.Ring.negate` est exclu de l'index tôt ; les autres barrières filtrent après la collecte |
| Demandes directes (PBO `collectInstantiations`) | Corps de l'appelant, annotations, types/présence des fournisseurs, arguments statiques et substitutions | L'appelant inchangé peut relire un type fournisseur différent ; conserver payload et provenance |
| Contribution transitive (PBO `collectEntry`) | Corps fournisseur, arguments/substitutions, dictionnaires externes consultés par `rewriteExpr.resolveDict`, ensemble des spécialisations | Versionner contenu et absences, pas seulement nombre d'entrées ou identité mémoire |
| Fusion/point fixe (PBO `transitiveCollectWith`) | Contributions ordonnées, choix du premier payload, union des appelants | Rétraction et cycles nécessitent une justification depuis les demandes encore présentes |
| Réécriture finale (PBO `monomorphize`) | Module source, corps/dictionnaires externes, spécialisations admissibles, barrières | Un fournisseur au TAST inchangé peut recevoir/perdre des copies ; ordre des groupes et usages réanalysés |
| Clés préparées (`Cache.Modules.initialize`) | Module spécialisé complet, directives/contexte, préfixe canonique ; métadonnées globales | La preuve `sameTast` est globale ; une édition TAST fait actuellement re-hacher tous les modules |

**Limite précise du cache transitif actuel :** `PreparedInstantiation` stocke
un corps substitué, une contribution et des lectures de spécialisations. Le
raccourci `sameDependencySizes` est valable dans la croissance monotone d'un
calcul à corps fixes. `SpecializationReads` mémorise surtout les absences qui
peuvent changer au tour suivant ; il ne constitue pas une trace complète des
lectures de corps/dictionnaires entre builds. La réutilisation persistante
exigerait donc des observations supplémentaires et une rétraction correcte.

### Frontières de transport à couvrir si P est poursuivi

- `Preparation.runPreparationJobs` : thunks purs, snapshots immuables, résultats
  rendus dans l'ordre des entrées ; la fusion reste côté PBO. La concurrence
  repose sur les implémentations hôtes de `Aff`, pas sur un codec de tâches.
- `Wire` : représentation portable partagée, tables de types et graphe à arêtes
  arrière. Les identités mémoire n'entrent pas dans les preuves persistées.
  Un `Module Ann` est déjà encodable, mais son coût de restauration reste à
  mesurer avant de retenir cet objet.
- `BuildCache.sessionExchange` : sessions, gel privé, intégrité, publication et
  récupération communs aux hôtes ; les objets de préparation devront emprunter
  ce stockage. Une économie de calcul n'autorise pas à omettre son coût durable.
- Métadonnées source : l'ABI que consomment Rebox, les wrappers FFI et l'émetteur
  reste dérivée des déclarations d'origine, avant copies spécialisées.

### Lecture attendue du potentiel

Pour un scénario, la borne optimiste est le temps des phases effectivement
supprimables, **sans additionner les horloges imbriquées**. Le gain net estimé
retire lecture/décodage, validation, recalcul de la région affectée et nouvelle
publication. Un coût non mesuré reste une hypothèse explicite. La réduction du
volume d'IR ou du nombre de demandes ne vaut pas une réduction de temps.

Les clés préparées se trouvent dans `optimize + emit`, et non dans le compteur
`prepare + monomorphize` : leur part ne doit être ni oubliée, ni comptée deux fois.
En FFI seule, la preuve `sameTast` peut déjà réutiliser les empreintes ; son gain
accessible diffère donc de celui des éditions feuille et partagée.

## Référence durable du 19 — 36 passages terminés

Les 36 passages backend sont byte-exacts entre témoin et instrumenté, avec les
mêmes TAST, clés préparées, parties métier et sources privées. Les **48 traces
de ressources** de préparation/mesure sont complètes, sans compilation externe
détectée. Données : `measurement/results.json`, `analysis/timings.json` et
`analysis/timings.csv` dans la campagne durable.

| Scénario | Témoin médian | Min–max témoin | Écart instrumenté médian | Résiduel à coût de préparation nul |
| --- | ---: | ---: | ---: | ---: |
| Identique | 1,621 s | 1,575–1,671 s | −5,13 % | 1,621 s |
| Feuille | 45,352 s | 44,717–45,795 s | +1,90 % | 28,679 s |
| Partagé | 47,047 s | 46,205–50,234 s | −0,26 % | 30,905 s |
| FFI | 35,056 s | 34,560–36,328 s | +2,34 % | 23,480 s |
| Froid applicatif | 115,692 s | 115,471–116,008 s | −0,02 % | 99,612 s |
| Froid complet | 118,837 s | 116,717–119,269 s | +1,34 % | 102,562 s |

L'écart d'instrumentation est le rapport des médianes. Les trois deltas appariés
restent dans le JSON : première paire feuille **+7,28 %**, première paire partagé
**−6,59 %**, deuxième paire FFI **+6,01 %**. Le seuil de représentativité médiane
de 5 % passe sur les scénarios qui exécutent la préparation ; il échoue de peu
sur l'identique (**−0,083 s**), qui n'exécute aucune de ces passes. Ce dernier
résultat ne valide pas une ventilation instrumentée de l'identique.

La dernière colonne est une **borne comptable irréalisable telle quelle** :
chargement, préparation et toutes les clés sont soustraits à chaque passage,
puis les résiduels sont résumés. Elle exclut même les coûts nécessaires de
restauration, preuve et recalcul. Les médianes de phases ne s'additionnent pas
nécessairement à la médiane du total.

### Ventilation du témoin (médianes, secondes)

| Frontière | Feuille | Partagé | FFI | Froid applicatif | Froid complet |
| --- | ---: | ---: | ---: | ---: | ---: |
| Lookup et gel des entrées | 2,962 | 2,869 | 2,976 | 0,883 | 1,022 |
| Chargement et tri TAST | 1,186 | 1,162 | 1,178 | 1,143 | 1,160 |
| Préparation et monomorphisation | 9,851 | 9,759 | 9,656 | 9,650 | 9,690 |
| Clés de cache module (dans optimiser/émettre) | 5,210 | 5,322 | 0,684 | 5,239 | 5,434 |
| Producteur PBO (chevauche l'émission) | 9,873 | 11,669 | — | 39,170 | 39,292 |
| Optimiser/émettre, clés incluses | 15,122 | 17,020 | 6,009 | 44,638 | 44,846 |
| Exports Go | 0,842 | 0,648 | 0,640 | 0,659 | 2,883 |
| Analyse et partition Go | 10,382 | 10,557 | 9,784 | 9,983 | 9,925 |
| Publication du cache | 1,456 | 1,400 | 1,312 | 46,581 | 45,149 |
| Résidu backend non attribué | 2,569 | 2,600 | 2,587 | 1,596 | 1,620 |
| Enveloppe de processus hors backend | 0,680 | 0,749 | 0,708 | 0,658 | 0,796 |

Le gel FFI ajoute environ **0,23 s**, le runtime moins de **0,002 s**. En FFI,
le producteur de modules restaurés coûte **2,600 s**, le lookup all-hit **2,616 s** ;
les empreintes préparées sont déjà réutilisées (**0,007 s**, contre **4,546 /
4,615 s** après feuille/partagé). La clé globale des métadonnées coûte encore
**0,643 / 0,682 / 0,650 s**. Ces sous-phases ne s'ajoutent pas une seconde fois à
`optimize + emit`.

Le détail instrumenté situe les demandes directes à **2,054 / 2,032 s**, le
transitif à **6,769 / 6,886 s**, et la réécriture finale à **0,597 / 0,618 s**
sur feuille/partagé. Le point fixe comporte **31 tours** dans tous les passages.
Ses jobs prennent **2,992 / 2,976 s** au total ; les **3,777 / 3,910 s** restantes
couvrent création/contrôle des tours et fusions, sans attribution plus fine.
Les passes source de métadonnées (types, constructeurs, classes, Rebox,
élisions, pointeurs, enums) totalisent environ **0,09 s** ; l'enrichissement
des classes ajoute **0,01 s**, et l'invalidation des usages **0,16 s**.
Le résidu interne de préparation est **0,024 s**.

Conséquences avant mesure du cache candidat :

- Le chargement source seul, même gratuit, économiserait environ **2,6 %** du
  backend feuille : il ne peut pas satisfaire seul le minimum autonome de 5 %.
- Les métadonnées seules, même en incluant leur clé globale, sont sous ce même
  minimum. Une persistance éventuelle doit justifier son rôle dans un chemin
  plus large, plutôt qu'un gain autonome de cette taille.
- Le plafond brut chargé/préparé/clés est **16,189 / 16,170 s** après
  feuille/partagé et **11,516 s** après FFI. Son coût réel reste à déduire.
- Au vrai froid, les objets réutilisables n'existent pas : P doit les produire.
  Même la soustraction fictive de toute la préparation laisserait environ
  **100–103 s**, bien au-dessus des **~60 s**. La publication froide (**45–47 s**)
  et PBO/émission restent une dette distincte ; le lot 04 terminé n'est pas
  rouvert par ce constat.

Le RSS témoin après édition est d'environ **10,9–11,0 Gio**, celui de
l'instrumenté du même ordre. Le cache objet courant contient **528,66 Mio**,
les Go générés **117,30 Mio** ; ce sont les repères de stockage pour chiffrer
les objets supplémentaires. Ils ne constituent pas des gains temporels.

## Coûts des unités persistables

Les sept catégories passent leurs contrôles d'intégrité et aller-retour
canonique. Les modules préparés reproduisent les hashes du backend témoin.
**45 actions d'encodage/restauration**, **63 réencodages de jeux changés** et
**84 publications** ont réussi, avec **193 traces de ressources** sans
concurrence détectée, construction de la sonde de stockage comprise.

Les durées ci-dessous sont les médianes de trois répétitions. La restauration
inclut lecture d'index, lecture des objets, taille/hash et décodage typé.
Le réencodage canonique de contrôle est mesuré séparément et n'est pas ajouté
à cette durée : les processus de sonde complets incluent aussi ce contrôle et
leur fin de vie. Les coûts d'interaction GC avec un backend complet ne sont pas
mesurés par ces actions isolées.

| Objet | Nombre | Octets uniques (Mio) | Encodage + écriture | Restauration | Publication initiale interne |
| --- | ---: | ---: | ---: | ---: | ---: |
| Modules source décodés | 2 704 | 343,36 | 3,233 s | 5,382 s | 10,107 s |
| Modules préparés | 2 704 | 268,96 | 4,540 s | 5,136 s | 9,983 s |
| Métadonnées globales | 1 | 19,20 | 0,686 s | 1,157 s | 0,393 s |
| Demandes directes par fournisseur | 14 650 | 122,91 | 1,329 s | 1,703 s | 52,356 s |
| Demandes par appelant | 2 704 | 191,14 | 1,562 s | 2,106 s | 8,931 s |
| Instanciations finales par fournisseur | 20 483 | 158,30 | 2,135 s | 1,978 s | 75,543 s |
| États transitifs `PreparedInstantiation` | 16 452 | 330,01 | 5,127 s | 3,700 s | 72,377 s |

La sonde de publication utilise les primitives durables actuelles et un objet
CAS par unité, plus un index atomique. Le coût initial très élevé des milliers
de petites unités est un résultat de **cette granularité et de ce protocole**.
Un regroupement éventuel serait une autre représentation à mesurer. Les temps
ne permettent pas d'attribuer précisément le coût entre validation, écritures
et synchronisations sans instrumentation supplémentaire.

Dispersion de restauration : modules préparés **5,107–5,340 s**, métadonnées
**1,121–1,179 s**, contributions transitives **3,661–3,873 s**. Les sondes
source/préparé/contributions culminent autour de **9,8 Gio RSS**, métadonnées
autour de **2,9 Gio**. Ce RSS reste celui des sondes, pas un delta intégré au
backend ; le temps demeure le critère principal.

### Deltas effectivement observés

- Feuille : **1 module préparé** changé, **13 231 octets** ; 1 contribution
  d'appelant, 2 états transitifs ; aucune métadonnée globale changée.
- Partagé : **2 modules préparés**, `Util.Type.String.String` et
  `Control.Semigroupoid`, **1 293 245 octets** ; 1 contribution d'appelant,
  2 états transitifs changés et **1 retiré**. Le fournisseur
  `Control.Semigroupoid` change sans édition de son TAST.
- FFI seule : aucune différence dans les sept catégories.
- Réencoder les modules préparés changés coûte **0,001 / 0,127 s** sur
  feuille/partagé. Publier ce jeu dans un CAS déjà peuplé coûte **0,087 /
  0,103 s** en interne ; le processus de stockage complet prend **0,342 /
  0,347 s**, vérifications préalables comprises.

Ces deltas constituent des oracles de sortie. Ils ne prouvent pas quels objets
peuvent être réutilisés : il faut découvrir les lectures invalidées avant de
connaître la région à recalculer. Le petit nombre de changements n'est donc pas
converti en pourcentage de temps économisé.

## Gain net accessible : mesures, bornes et inconnues

Le calcul retire au temps brut supprimable la restauration mesurée, le
réencodage des seuls objets changés et les processus de stockage marginaux.
Les coûts de familles sont additionnés comme **hypothèse comptable**, puis le
résiduel est calculé pour chaque passage. La fusion dans une transaction unique
pourrait amortir une partie de ces coûts ; les preuves, les index inverses et
le recalcul affecté restent à ajouter. Ces modèles ne sont pas des campagnes
d'acceptation d'un compilateur optimisé.

| Modèle | Gain feuille avant preuves/recalcul | Gain partagé avant preuves/recalcul | Temps résiduel feuille / partagé |
| --- | ---: | ---: | ---: |
| Modules source seuls | **−4,621 s (−10,19 %)** | **−4,640 s (−9,88 %)** | 49,973 / 51,695 s |
| Modules préparés + métadonnées | **9,418 s (21,06 %)** | **9,271 s (19,65 %)** | 35,450 / 37,804 s |
| Préparés + métadonnées + appelants + états transitifs | **2,313 s (5,17 %)** | **2,233 s (4,83 %)** | 42,556 / 44,842 s |
| Appelants + états transitifs seuls | **1,717 s (3,79 %)** | **1,979 s (3,97 %)** | 43,635 / 45,068 s |
| Instanciations finales seules | **5,388 s (11,88 %)** | **5,914 s (11,85 %)** | 39,964 / 41,133 s |

Les pourcentages et résiduels sont les médianes des calculs par passage ; ils ne
se déduisent pas nécessairement par soustraction/division des médianes du tableau.

**Lecture pour la décision :**

1. **Rejeter la persistance intégrale du source décodé.** Sa restauration coûte
   plus de quatre fois le chargement TAST actuel, pour 343 Mio supplémentaires.
2. **Ne pas engager P01–P06 avec les états complets sondés.** Le jeu préparé +
   métadonnées + appelants + états transitifs ajoute **809,32 Mio** et demande
   **12,099 s de restauration**. Il reste sous le minimum de −10 % sur les deux
   éditions avant même d'implémenter les nouvelles preuves. Rendre gratuitement
   publication et réencodage ne suffirait toujours pas sur les médianes : la
   restauration seule absorbe trop du plafond brut d'environ 16,2 s.
3. **Préparés + métadonnées n'est qu'un plafond partiel.** Ces 288,16 Mio
   demandent 6,293 s de restauration. Il reste **4,946 / 4,650 s** pour financer
   toutes les preuves, leurs objets/index et le recalcul, si l'on veut encore
   atteindre −10 %. Le parcours intégral P03–P04 ne tient pas dans ce reliquat
   avec les objets actuels. La faisabilité d'une preuve compacte reste
   **indéterminée**, et le gain de 9,4/9,3 s n'est pas une prévision livrable.
4. **Les instanciations finales ne suffisent pas à rétracter.** Elles ne portent
   pas la provenance nécessaire après édition ; leurs marges pour conserver
   −10 % sont seulement **0,853 / 0,904 s** avant preuves/recalcul. Leur
   publication froide, à cette granularité, les disqualifie aussi.

En FFI seule, la preuve globale de TAST identiques permet d'envisager une voie
plus simple : le modèle préparés/métadonnées laisse **4,748 s** (13,42 %) et
un résiduel de **30,248 s** ; les instanciations finales seules laissent
**6,020 s** (16,57 %) et **28,889 s**. Ces voies exigent encore leur intégration
et un coût froid acceptable. Elles ne satisferaient pas à elles seules le
minimum demandé sur **feuille et partagé**.

### Dette froide et budgets proposés pour une éventuelle reprise

Le budget de régression froide de +5 % vaut **5,785 s / 5,942 s** sur cette
référence. Aux granularités sondées, la seule publication initiale coûte :

- préparés + métadonnées : **10,377 s** ; encodage + processus de publication
  isolés : **15,845 s** ;
- préparés + métadonnées + appelants + états transitifs : **91,684 s** ;
  encodage + processus de publication isolés : **104,151 s**.

Au froid, les clés actuelles paient déjà un encodage préparé/métadonnées.
Conserver leurs octets pourrait en éviter la duplication. Même en supposant
cette duplication entièrement supprimée, la publication additive du layout
actuel dépasse le budget. Une intégration recouvrant ces travaux serait une
autre hypothèse à éprouver, pas un gain comptabilisé ici. Aucun de ces résultats
ne réduit la dette des **~60 s** au vrai froid.

Si l'utilisateur autorise une reprise ajustée, proposer avant le premier
candidat les limites suivantes sur ce corpus :

| Budget proposé | Limite et justification |
| --- | --- |
| Coût incrémental complet de préparation persistante | **≤11 s** après feuille/partagé : restauration, intégrité, preuves, recalcul, sérialisation, publication et surcoûts de session inclus ; laisse une marge vers le minimum de −10 % |
| Restauration, intégrité comprise | **≤6,5 s**, à inclure dans les 11 s ; le jeu préparé/métadonnées est déjà à 6,293 s |
| Stockage supplémentaire courant | **≤320 Mio**, objets et index compris ; environ **+60,5 %** du CAS de 528,66 Mio, proposition explicite propre à P |
| Conservation et staging supplémentaires | **≤640 Mio** pour courant + précédent/staging, avec politique bornée à définir ; aucune croissance illimitée implicite |
| Froid | **≤+5 %**, soit 5,785 / 5,942 s ; coûts de première publication inclus |
| Identique | **≤+0,100 s** ici, conformément à `max(5 %, 0,100 s)` |
| Autres scénarios | **≤+5 %** ; chaque optimisation autonome annoncée doit gagner **≥5 %** |

Au bilan P00, ce sont des **budgets soumis à décision**, pas des résultats
atteints ou une autorisation d'optimiser ; ils ont ensuite été repris pour
l'étude C01 autorisée. Ils conservent le minimum de 10 %, sans promettre
~30 s : avec la restauration déjà mesurée, même le modèle partiel optimiste
laisse environ **35–38 s** sur feuille/partagé avant les preuves.

### Unités et effort : recommandation avant P01

- **Écarter** les snapshots source décodés et la persistance unitaire brute des
  `PreparedInstantiation`. Leur intégration est concevable, mais les mesures
  ne justifient pas son coût.
- **Garder comme candidats à une étude bornée** les modules préparés et les
  métadonnées, avec preuves/index nettement plus compacts et publication
  regroupée à mesurer. Il faudrait d'abord réviser P01 pour éprouver stockage,
  restauration et coût froid, avant les quatre lots fonctionnels P02–P05.
- **Effort moyen** pour un snapshot global réutilisé uniquement sous preuve
  complète `sameTast` : raccordement session, sélection des entrées/diagnostics,
  codec et repli, plus validation des quatre hôtes. Cela cible surtout FFI.
- **Effort élevé** pour la préparation générale après édition : observations
  de lectures dans PBO, provenance des contributions, index inverses, rétraction
  des cycles, ordre des fusions et réécriture sélective. Les raccourcis actuels
  ne fournissent pas ces preuves ; l'effort couvre toujours P02–P06 et dépend
  d'une nouvelle représentation validée en amont.

**Bilan P00 :** potentiel insuffisant pour les objets complets sondés ; potentiel
d'une représentation compacte encore indéterminé. **Recommandation : suspendre
le P01 actuel et décider d'un périmètre révisé avant tout nouveau candidat.**

## Validation et clôture

- Référence b8x : **2 704 TAST / 755 005 types**, mêmes octets dans tous les
  hôtes ; **36 passages appariés** et objets métier byte-exacts.
- Validation b8x : **7 couples hôte/état**, chacun avec build, hit et forcé,
  soit **21 passages backend** ; base dans les quatre hôtes, trois éditions en
  Go parallèle ; **4 exécutions** avec les oracles attendus et le runner borné.
- Fixtures : **28 transitions exécutées**, mêmes Go, parties de cache et clés
  préparées entre JS, Go séquentiel/parallèle et Rust, ainsi qu'avec un build neuf.
- Sondes : sept catégories, empreintes préparées confrontées au backend,
  aller-retour canonique, deltas FFI nuls, sérialisation et publication répétées.
- Les snapshots de référence n'ont pas été mis à jour ; les comparaisons sont
  strictes. Les échecs de harnais du 19 restent archivés. Les preuves du 16
  perdues sont distinguées des résultats durables du 19.

Les mesures concernent le **processus backend**, pas la chaîne jusqu'à `exec`.
Les exécutions runner sont des validations séparées. Les chiffres G06 et v3/06
restent des références historiques ; aucun delta face à eux n'est attribué à P.

P00 clôt l'audit et porte P à **1/7**. Le compilateur optimisé de préparation
n'est pas livré par ce lot. L'étude C01 autorisée ensuite est close avec rejet
du candidat en l'état ; son bilan renouvelle le point de décision avant la suite.

## Premiers résultats du 16 octobre — preuves temporaires perdues

Le frontend privé a reproduit **les 2 704 TAST / 755 005 types à l'octet près**
face à la sortie b8x disponible, après correction du harnais. Le passage de
reconstitution prend 26,636 s ; ce n'est pas une médiane d'acceptation.
La copie initiale inventorie 4 408 fichiers, complétés par les reprises tracées.
Le seed natif réussi prend **119,852 s**, suivi d'un hit à **1,636 s** ; ces deux
passages préparent la campagne et ne sont pas ses médianes. La copie instrumentée
est construite depuis **512 TAST**, avec 28 contrôles de préparation/métadonnées
passants. Sur b8x, elle reproduit les 27 Go, les clés préparées, les quatre parties
de cache par module et les sources privées à l'octet près.

La campagne appariée comporte **36 passages du processus backend**, sans runner
applicatif dans ces chronomètres. Les validations d'exécution séparées conservent
les limites du runner G06. L'enveloppe hors compteur interne (`backend total`)
est publiée séparément : lancement, gel du compilateur, fin du processus et
instrumentation `/usr/bin/time` en font partie.
Dans la reprise du 19, l'enveloppe comprend aussi le lancement et les relevés
initial/final du surveillant de ressources. Ce coût commun aux deux membres
d'une paire est conservé ; les anciennes mesures sans surveillant ne servent
pas à attribuer un gain.

La chaîne utilisée pour produire le natif disponible et résoudre les exports est
**Go 1.27.0 darwin/arm64**. Le runner d'exécution reste **Go 1.26.8 linux/arm64**
dans l'image G06. Cette référence reconstituée n'isole donc pas un delta à G06.
Les valeurs historiques G06 restent **48,487 / 49,672 / 40,001 s** après
feuille/partagé/FFI et **117,830 / 117,740 s** aux froids applicatif/complet.
Elles ne constituent pas encore des mesures P00.

Les 36 passages du 16 ont terminé avec égalité des sorties. Les médianes de
processus témoin étaient **1,464 / 49,194 / 50,943 / 46,411 s** pour identique,
feuille, partagé et FFI. Les froids étaient trop dispersés : **147,579–426,383 s**
applicatif et **177,054–332,559 s** complet côté témoin. Une compilation externe
dans b8x, commencée à 17 h 27, recouvrait une partie du froid complet ; elle
n'explique pas le passage lent antérieur. La surveillance observait environ
5 cœurs utilisés, 21 Gio résidents et 1,7 Gio de swap système après la campagne.
Ces valeurs sont des diagnostics historiques ; les phases exactes et les coûts
d'objets doivent être remesurés dans la reprise durable.

### Sondes et reproducteurs

- `tools/preparation-baseline.mjs` fige les entrées, reconstruit le frontend et
  vérifie le seed. Les reprises sont des opérations distinctes et les logs des
  tentatives précédentes restent présents.
- `tools/preparation-probe/instrumentation.patch` détaille les frontières dans
  une copie privée de `Driver.Prepare` et `Monomorphization`. Le builder privé
  est `tools/preparation-probe-build.mjs` ; les compilateurs installés ne servent
  que de générateurs témoins.
- `tools/preparation-measure.mjs` conserve les quatre états TAST, les bases de
  cache et les sorties. La restauration revient au chemin d'origine de chaque
  copie privée ; l'inode de son verrou est conservé. Les archives de journaux ne
  sont jamais activées dans leur répertoire d'archive.
- `tools/preparation-analyze.mjs` calcule les frontières additives, les résidus,
  les tours transitifs, les médianes/min/max et les écarts par paire. Le cumul
  d'émission chevauche le producteur ; il n'est pas ajouté au total.
- `tools/preparation-probe/Probe.purs` encode et restaure source, préparation,
  métadonnées, demandes directes/transitives et contributions. Les demandes
  directes fusionnées par fournisseur et les contributions de chaque appelant
  sont deux objets distincts, tous deux mesurés. Chaque
  restauration vérifie le hash et refait l'encodage canonique hors de son
  chronomètre de lecture. Trois répétitions alternent encodage/restauration,
  avec un processus neuf par action et le cache de pages du système chaud.
- `tools/preparation-probe/store.go` lie les fonctions du stockage figé :
  `stageArtifact`, `putStagedObject`, `boundedEach`, barrière de répertoire et
  publication atomique d'un index. Le compteur interne exclut la vérification
  préalable des objets déjà présents ; la durée du processus l'inclut. C'est
  le **coût marginal du stockage**, pas une transaction de build complète avec
  de nouvelles preuves de dépendances. Les deux durées sont conservées.
- `tools/preparation-codec-measure.mjs` conserve octets, empreintes, deltas par
  édition, roundtrips et coûts de publication. Le fait qu'un objet se décode
  correctement ne prouve pas qu'il est encore réutilisable après une édition.
- `tools/preparation-costs.mjs` soustrait les coûts médians des objets aux phases
  observées, puis calcule le résiduel de chaque passage avant la médiane. Le
  budget restant appartient aux preuves et au recalcul affecté, non implémentés
  dans P00. Le stockage mesuré reprend la lecture de tous les objets par le
  stockage actuel ; une future transaction partagée pourrait amortir une partie
  de ce coût. Au froid, les clés existantes encodent déjà modules préparés et
  métadonnées : réutiliser ces octets pourrait éviter un deuxième encodage, mais
  le coût de cette intégration n'est pas mesuré par une sonde séparée.
- `tools/preparation-publish.mjs` exporte temps bruts, analyse des coûts et
  empreintes des preuves vers `docs/` après vérification de la complétude des
  mesures, de l'absence de concurrence détectée et des validations.
- `tests/host/PreparationTransitions` et le pilote de transitions ajoutent les
  cas de cycle, de cardinalité constante, de fournisseur de dictionnaire et de
  FFI seule. Les observations PBO complètent la comparaison des Go exécutés.
- `tools/preparation-validate.mjs` reprend b8x dans les quatre hôtes sur la base,
  puis les trois éditions en natif parallèle, avec cache/forcé/hit et quatre
  oracles d'exécution. Les petites transitions couvrent toutes les éditions
  dans tous les hôtes.

### Échecs de reconstitution du 16 — journaux temporaires perdus

- `frontend-base` : checkout Git `html-parser-halogen` absent du cache Spago
  privé ; reconstitué depuis le dépôt copié, révision vérifiée contre le lock.
- Le premier helper de reprise cherchait `extraPackages` au lieu de
  `extra_packages` dans le lock JSON ; la correction conserve le même contrôle
  de révision.
- `frontend-base-attempt-2` : copie initiale incomplète, répertoire `b8x/test/`
  omis ; 46 imports de `Test.Util.Assert` échouent. La reprise ajoute ces sources
  depuis la révision figée et contrôle les chemins source référencés par les TAST.
- `native-base` : `go.mod`/`go.sum` applicatifs omis ; les exports externes en
  `-mod=readonly` échouent après préparation/émission. Passage incomplet conservé,
  puis copie des fichiers de module avec empreintes et sauvegarde de l'état
  précédent. Ses 59,492 s internes ne sont pas un temps de build réussi.

### Reprise durable du 19

- `p00-20261019-freeze1-incomplete/` conserve la première copie interrompue par
  le contrôle de typage du CoreFn disponible. Le harnais autorise désormais
  l'inventaire d'une sortie historique non typée, tout en exigeant les TAST typés
  après reconstruction frontend.
- La nouvelle copie contient 4 416 fichiers inventoriés. Les sources applicatives
  viennent de l'archive Git de la révision P00 initiale ; le travail courant reste
  dans sa révision propre.
- Au contrôle de reprise : aucune compilation lourde, **0 swap**, 48 Gio de RAM
  physique. La reconstruction des références et des sondes a été menée sous
  surveillance de concurrence.
- Le frontend privé régénère **2 704 TAST / 755 005 types** en 27,411 s. Le seed
  natif passe en **116,277 s**, avec 27 sorties vérifiées. La sonde de temps
  reproduit ses empreintes du 16 : JS `3cc5798724a28a2b7344eb63be88dbedfe6bc6f1b80657a80f0a3d3293753c70`,
  natif `a78e70042ad55b317cfcf25b3c5f7e0d31d3cd3ec071aa8bb830732273e9a0c7`.
- `codec-v1-failed/` conserve l'erreur de compilation de la première sonde de
  codecs : l'appel `Process.exit 1` doit être `Process.exit' 1`. Ce défaut du
  harnais n'est pas un échec du codec. La version suivante utilise les petits
  lots dynamiques du cache et mesure aussi le réencodage des objets changés.
- La sonde de codecs compile depuis **513 TAST / 313 300 types**, après les
  contrôles de préparation/métadonnées. Son natif porte l'empreinte
  `6999a2685a9ac88fa3bcf35877dd541a5c195014b3b070c7efddfcc58e205eed` ; son source
  est vérifié identique à `tools/preparation-probe/Probe.purs` avant mesure.
- Les nouvelles fixtures passent : **7 états × 4 hôtes = 28 comparaisons et
  exécutions**. Chaque transition compare sorties, quatre parties de cache,
  clés préparées et sources privées avec un build neuf ; la relance identique
  préserve les dates. Les observations PBO vérifient aussi le remplacement à
  cardinalité constante et la disparition du soutien cyclique après rétraction.
  Preuves : `run-3/workspaces/gopurs-tests-gfsQP6/`, sans mise à jour de snapshot.
- Contrôle final : 14 scripts JS passent `node --check`, le surveillant passe
  son contrôle d'attribution, la sonde Go respecte `gofmt`, le patch privé
  s'applique à la référence et les empreintes exportées correspondent aux
  analyses. `git diff --check` passe ; les artefacts installés correspondent
  encore aux empreintes validées. Les dépôts b8x et PBO sont propres.
