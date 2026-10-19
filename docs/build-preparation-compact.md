# P01 — étude resserrée : objets compacts et publication regroupée

**Étude autorisée et terminée le 19 octobre 2026. C01 rejeté en l'état.**
La publication regroupée gagne fortement, mais le décodage/restauration et le
surcoût froid dépassent les budgets. P reste à **1/7** pour l'audit P00 ; une
nouvelle décision est nécessaire pour recadrer P avant l'intégration de P01.
Référence : [P00](build-preparation-baseline.md), mêmes entrées et artefacts
conservés dans `htdocs/.gopurs-evidence/p00-20261019/`.

Exports vérifiables : [résultats, modèles et empreintes](build-preparation-compact-c01.json)
et [96 mesures brutes de stockage](build-preparation-compact-c01.csv).

## Périmètre et seuils fixés avant les essais

Un candidat **C01** : 64 groupes stables SHA-256 par famille, gzip niveau rapide,
octets Wire individuels inchangés, index d'objets et engagement atomique via les
primitives du stockage figé. Les familles sont préparés, métadonnées, demandes
par appelant et contributions transitives. La variante compacte des dernières
retire le corps substitué et remplace les entrées/observations par des empreintes.

- **≤320 Mio** supplémentaires courants, index compris ; **≤640 Mio** avec
  état précédent/staging, conservation bornée à éprouver avant intégration.
- **≤6,5 s** de restauration/intégrité ; **≤11 s** pour le nouveau parcours
  complet après édition, preuve/recalcul/sérialisation/publication compris.
- Froid **≤+5 %** (5,785 / 5,942 s sur P00), identique **≤+0,100 s**, autres
  scénarios **≤+5 %** ; minimum backend feuille/partagé **−10 %**, et **−5 %**
  pour chaque optimisation autonome. RSS observé, temps prioritaire.
- Trois paires alternées raw/regroupé, bases privées restaurées et mesures
  sérielles sous surveillance de concurrence. Tous les essais sont conservés.
- Un seul candidat avant bilan : un coût bloquant ou une preuve insuffisante
  revient à la décision ; pas de recherche automatique de multiples layouts.

Les coûts de regroupement, compression, vérification, extraction et publication
sont inclus. Les encodages déjà payés pour les clés préparées sont distingués
d'un encodage supplémentaire. Un gain local de stockage ne vaut pas un gain
mesuré du backend intégré.

## Contrat de preuve du prototype

Le `PreparedInstantiation` actuel contient `info`, `expr`, `dependencies` et
`contribution`. Les observations `dependencies` ne prouvent pas toutes les
lectures externes entre builds. La projection compacte conserve la contribution
et deux empreintes d'audit (`info`, `dependencies`), sans prétendre remplacer ces
lectures manquantes.

L'index porte une **clé globale des entrées du contexte figé** : inventaire et
contenu de tous les TAST, identité du compilateur, options et provenance. Le
harnais fournit cette clé à partir des inventaires conservés ; cela n'implémente
pas un collecteur de dépendances de production. Une édition TAST
invalide donc ce snapshot, même quand les octets d'une contribution ne changent
pas. Une édition FFI seule peut conserver la preuve de préparation. Ajout,
retrait, renommage et changement du contenu des TAST appartiennent à l'identité ;
l'inventaire est trié canoniquement, les ordres internes aux TAST restent hashés.
La contribution compacte n'est pas un état de reprise du cache monotone PBO :
un résultat non prouvé doit être recalculé depuis les sources.

Le coût et la taille d'une preuve sélective P03 restent un poste distinct. Ils
ne sont pas remplacés par une comparaison de cardinalités ou d'empreintes de
résultat. Les fixtures P00 servent à éprouver les invalidations conservatrices.

## Bridges

`Wire` reste l'oracle des objets individuels. Le regroupement introduit une
enveloppe privée dans une sonde liée au stockage Go commun : gzip, index, hashes,
verrou et engagement. L'extraction restitue les octets originaux aux codecs
typés ; un contrat de session de production ne sera défini qu'après ce bilan.
La projection compacte utilise un record portable et les codecs existants,
dans une copie privée du compilateur.

## Campagne et mesures locales

Preuves durables : `htdocs/.gopurs-evidence/p01-compact-20261019-c01/`.
La sonde et les dépendances sont copiées depuis P00. Les deux constructions
privées réussissent les contrôles de préparation et de métadonnées : la seconde
ajoute la restauration complète au même exécutable pour la comparaison
appariée ; la première reste dans `compact-v1-projection-only/`.

La projection compacte de la base est byte-exacte entre JS et Go avec huit
workers de codec. Sur les quatre états b8x projetés en Go, elle contient
16 452 entrées ; la FFI seule reproduit intégralement la base.
Le remplacement partagé retire une ancienne clé et en introduit une nouvelle.

| Phase | Médiane, trois répétitions |
| --- | ---: |
| Lecture/décodage de l'état transitif complet (même sonde) | 3,851 s |
| Lecture/décodage compact | 1,084 s |
| Construction des deux empreintes d'audit | 1,739 s |
| Encodage/écriture compact | 1,102 s |

Les octets uniques transitifs passent de **330,01 à 131,78 Mio**, avant
compression et index de regroupement. Les gains locaux sont d'environ **60 %
en octets et 72 % en restauration**, avec **2,841 s** de construction/encodage
au premier calcul. Les phases de roundtrip de contrôle sont séparées ; leur
temps n'est pas présenté comme un coût de restauration utile.

La projection de la sonde part d'objets complets déjà décodés. Son compteur de
lecture des anciens objets reste hors du coût proposé au froid : le collecteur
normal dispose déjà de ces valeurs en mémoire. Les effets de layout/GC dans
un backend intégré devront encore être mesurés.

Le helper privé compare les layouts raw et regroupé sur les jeux complet et
compact, trois paires alternées chacun, avec base, feuille, partagé et FFI.
Chaque édition repart des objets de base ; l'inode du verrou est conservé.
Le froid du stockage signifie **CAS vide**, avec les pages d'entrée du système
chaudes comme en P00. Les 48 restitutions sont byte-exactes face aux oracles
Wire P00 et aux projections compactes. Les manifests/blobs regroupés sont
déterministes entre les trois répétitions.

Les contrôles de protocole précèdent les mesures longues : preuve périmée,
échec avant engagement, corruption/absence de blob, index corrompu et reprise.
Les sept états des fixtures P00 éprouvent la clé globale avec leurs véritables
inventaires TAST ; ce contrôle de clé utilise de petits objets témoins. Les
oracles métier quatre hôtes sont ceux de P00 (build/hit/forcé sur mêmes TAST),
tandis que la nouvelle projection compacte est comparée entre JS et Go.
La nouvelle projection n'a pas été exécutée en Rust ni en Go séquentiel. Une
intégration requerrait toujours ces validations et les replays stricts ; le
rejet temporel de C01 n'est pas une acceptation de bridge de production.

### Publication regroupée

« Complet » désigne ici préparés + métadonnées + demandes par appelant +
`PreparedInstantiation` ; « compact » remplace la dernière famille par sa
projection. Cela ne mesure pas tous les modèles de P00 sous un nouveau nom.
Les deux modèles ont chacun leurs paires raw/regroupé ; leur comparaison
mutuelle n'est pas une campagne appariée de backend intégré.

| Mesure à la base | Complet raw → regroupé | Compact raw → regroupé |
| --- | ---: | ---: |
| Publication, processus médian | 82,753 → **3,729 s** | 90,972 → **2,947 s** |
| Plage des trois publications regroupées | 3,344–3,792 s | 2,888–3,224 s |
| Réduction de publication par paire | 95,45–95,94 % | 96,51–96,83 % |
| CAS courant, index compris | 824,83 → **121,37 Mio** | 626,38 → **91,93 Mio** |
| Groupes publiés | 21 861 → 193 | 21 861 → 193 |
| Extraction regroupée, médiane [min–max] | 1,655 [1,563–1,897] s | 1,781 [1,330–1,826] s |

Les trois familles multi-objets utilisent chacune 64 groupes ; les métadonnées
forment un groupe. Après feuille/partagé/FFI, la publication regroupée compacte
coûte **0,519 / 0,562 / 0,326 s** en médiane. La variation de l'extraction reste
visible dans le CSV. Le compact raw est plus lent à publier que le complet raw
dans cette campagne malgré sa taille moindre : les octets persistés ne sont
pas utilisés comme substitut à un chronomètre.

Ces chronomètres incluent lancement/surveillance/fermeture du processus, comme
en P00. Le JSON et le CSV gardent aussi les temps internes. La restauration
typée inclut lecture, intégrité et décodage ; les roundtrips de contrôle sont
chronométrés séparément.

## Coût complet comptable et critères

Le modèle additionne les restaurations typées préparés/métadonnées/appelants
de P00, la restauration transitive mesurée avec la même sonde complète/compacte,
l'extraction C01 et les encodages/publications des changements. Il ne suppose
aucun recouvrement de travaux. Ce sont des **coûts connus combinés**, pas des
temps de backend incrémental livré. Les résiduels sont calculés pour chaque
passage P00 avant la médiane.

| Critère | Budget fixé | Complet regroupé | Compact regroupé |
| --- | ---: | ---: | ---: |
| Décodage typé seul | inclus dans restauration | 12,250 s | **9,483 s** |
| Extraction + restauration/intégrité | ≤6,5 s | 13,905 s | **11,264 s — dépassement** |
| Coûts connus après feuille / partagé | ≤11 s, preuves/recalcul compris | 14,373 / 14,460 s | **11,784 / 11,964 s — dépassement avant preuves** |
| Surcoût froid applicatif / complet, crédit maximal des clés | ≤5,785 / 5,942 s | 10,417 / 10,223 s | **7,349 / 7,155 s — dépassement** |
| CAS courant supplémentaire | ≤320 Mio | 121,37 Mio | **91,93 Mio — passe** |
| CAS base/édition + extraction temporaire, maximum connu | ≤640 Mio | 952,68 Mio | **725,18 Mio — dépassement** |

Le **décodage seul dépasse déjà 6,5 s** : supprimer entièrement l'étape
d'extraction ne suffirait pas. Les 621,36 Mio de fichiers Wire extraits du
compact sont comptés avec le CAS et son état précédent ; ce total ne prétend
pas être un pic échantillonné de chaque écriture filesystem. La conservation
bornée en production reste à définir. Les archives de campagne sont distinctes
des besoins d'une transaction.

Au froid, le crédit favorable soustrait **5,227 / 5,421 s** déjà consacrées aux
clés préparées/métadonnées. Cela suppose que leurs octets encodés puissent être
retenus et réutilisés. Sans ce crédit, le compact ajoute **12,576 s** ; même
avec lui, il ajoute environ **6,35 % / 6,02 %** aux froids P00. Les effets de
session, de layout et de GC ne sont pas mesurés dans un backend intégré.

### Gain net hypothétique après édition

| Modèle | Feuille : gain / résiduel | Partagé : gain / résiduel |
| --- | ---: | ---: |
| Complet regroupé | 1,816 s (4,06 %) / 43,052 s | 1,710 s (3,70 %) / 45,365 s |
| Compact regroupé | **4,405 s (9,85 %) / 40,464 s** | **4,206 s (9,10 %) / 42,869 s** |

Ces marges supposent toute la préparation actuelle évitée, puis déduisent les
coûts connus. **Preuves sélectives, index inverses, rétractions et recalcul
affecté restent à financer.** La clé globale réellement implémentée invalide
toute édition TAST ; elle ne réalise donc pas ces gains feuille/partagé.

Les trois échantillons du modèle compact vont de **9,10 à 10,78 %** sur feuille
et de **8,88 à 9,18 %** sur partagé. Le résultat feuille est proche du minimum
de 10 %, mais cette dispersion n'efface ni l'absence de marge pour les preuves,
ni les dépassements indépendants de restauration et de froid. La marge médiane
restante pour conserver −10 % est déjà **−0,067 / −0,414 s**. En FFI seule, le
modèle compact est presque neutre (**−0,073 s de gain**, soit +0,20 % de temps).
La relance identique intégrée n'a pas été mesurée pour C01.

Les repères restent **~30 s après modification / ~60 s à froid**. Même les
résiduels hypothétiques compacts restent vers **40–43 s** après édition ; le
modèle froid serait vers **123–126 s**, contre 115,692–118,837 s de référence.

## Bilan et recommandation

**C01 est rejeté en l'état pour les temps, avant même le dépassement de
stockage temporaire.** Le regroupement/compression résout le coût local de
publication, mais la restauration globale de ces familles consomme trop du
temps qu'on cherche à économiser. Aucune optimisation autonome n'est acceptée
sur le temps backend à partir de ces mesures isolées.

Recommandation : **recadrer P avant de poursuivre**, en exigeant une réduction
du volume décodé/restauré et un chemin complet finançant les preuves sélectives
et le recalcul. Le regroupement peut servir de brique à ce futur cadrage ; son
gain local ne justifie pas à lui seul P02–P06. Ce potentiel reste indéterminé.
L'étude autorisée s'arrête à ce bilan, avec un nouveau point de décision.

### Preuves et clôture

- **22 actions de projection/restauration**, dont trois paires de restauration,
  une projection JS byte-exacte et neuf projections des jeux changés.
- **96 mesures de publication/extraction**, trois paires alternées par modèle
  et quatre états ; 48 restitutions byte-exactes et manifests déterministes.
- **Quatre rejets attendus** : interruption avant engagement, blob corrompu,
  blob absent, index corrompu. Preuve périmée traitée en miss ; restauration
  réussie après interruption et après remplacement du blob.
- Sept états de clés globales : ajouts/remplacements/rétractions cycliques,
  fournisseur de dictionnaire, FFI et retour initial ; seule la transition FFI
  conserve la clé précédente.
- **145 traces de commandes complètes et sans compilation externe détectée**,
  avec logs et empreintes dans le JSON publié. Les deux builds de sonde privés
  et leurs contrôles de préparation sont conservés séparément.
- RSS observé : pic de projection **10,08 Gio**, helper de stockage **9,51 Gio**
  avec `GOGC=off` / `GOMEMLIMIT=10GiB`. Les restaurations complètes/compactes
  sont également proches de cette limite ; le rejet repose sur le temps.
- Première sonde conservée dans `compact-v1-projection-only/` ; première analyse
  et ses exports conservés dans `analysis-v1-without-evidence-manifest/`.
  La seconde analyse ajoute le manifeste des preuves et le périmètre précis
  des validations, sans remplacer les mesures.
- Clôture : **589 empreintes de fichiers vérifiées**, exports CSV/JSON identiques
  aux analyses conservées, sources des sondes confrontées aux copies compilées,
  quatre scripts JS contrôlés par `node --check`, `gofmt` et `git diff --check`
  passants. Le patch privé s'applique à la référence ; empreintes installées
  vérifiées et dépôts b8x/PBO propres. Relevé dans `closure.json`.

### Reproducteurs

- `tools/preparation-compact-build.mjs P00 NEW_STUDY` : copie auditée et build
  privé, avec `preparation-probe-build.mjs … compact` pour une reprise archivée.
- `tools/preparation-compact-measure.mjs project STUDY` : projection, identité
  JS/Go, trois paires de restauration et deltas des éditions.
- `tools/preparation-compact-measure.mjs packs STUDY` : helper figé, contrôles de
  protocole, trois paires de publication/extraction et validation des octets.
- `tools/preparation-compact-analyze.mjs STUDY` : distributions, empreintes,
  coûts complets comptables et exports ; aucune acceptation d'optimisation
  intégrée n'est déduite de ces seules micro-mesures.
