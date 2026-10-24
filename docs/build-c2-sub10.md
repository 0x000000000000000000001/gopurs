# C2-8 — backend one-shot central sous 10 s confirmé

La poursuite [C2-10](build-c2-sub2.md) vise 2 s sans daemon ; L reste
au-dessus de ce seuil. La campagne est close pour impasse pratique ; son
audit strict échoue uniquement sur le seuil de vitesse. La médiane
**8,586 s** ci-dessous demeure le résultat historique propre à C2-8/v4.

La suite [C2-9](build-c2-sub4.md) atteint désormais 3,841 s backend médian
sur trois paires exactes et 8,791 s pour le vrai `b`. Les chiffres ci-dessous
restent les résultats historiques propres à C2-8/v4.

**C2-8 achevé avec v4 : 8,586 s backend médian, trois paires byte-exactes,
chaque incrémental strictement sous 10 s.** Compilation Go, vrai `b`, deux
fallbacks sur corruption et daemons validés ; l'audit de clôture passe ses
25 critères. Le vrai `b` prend **13,651 s au total**, dont **8,977 s backend**.
Ce document décrit les changements qui prolongent [C2-7](build-c2-persistent.md).
Le résultat C2-7 de 16,970 s médian reste une mesure historique distincte.

La portée est une édition de littéral dans le corps de
`src/Util/Type/Variant/Encoding/Test/DecodeJsonWith.purs` de b8x, module
`Util.Type.Variant.Encoding.Encoding.Test.DecodeJsonWith`, avec cache amorcé.
Chaque backend démarre dans un processus neuf, sans daemon ni watch. Le seuil
porte sur le backend ; frontend, oracle de comparaison et compilation Go
sont mesurés séparément. Les contrôles d'intégrité habituels du backend restent
inclus dans son temps. Il ne promet pas le même temps pour un changement de
type, de dictionnaire, d'argument statique, de structure ou de FFI.

## Mesures finales v4

Binaire v4 :
`0394c4fe3aceb577f99480872853c17c44a8d43f309105b4a47909e0242f5fd5`.
L'[état des sources v4](../../../.gopurs-evidence/c2-8-sub10-20261023/v4-source-state.json)
inclut gopurs, PBO et la correction du runtime natif `gopurs-aff`.

La [campagne v4](../../../.gopurs-evidence/c2-8-sub10-20261023/central-validated-v4/results.json)
comprend l'amorçage et trois paires indépendantes sur ce même binaire, dont le
SHA est resté stable.

| Étape | Backend incrémental | Full forcé | Processus incrémental | Frontend séparé | Oracle |
| --- | ---: | ---: | ---: | ---: | --- |
| Édition 18 | 8,586 s | 79,464 s | 9,197 s | 2,477 s | identique |
| Édition 19 | 8,552 s | 79,994 s | 9,151 s | 2,275 s | identique |
| Édition 20 | 8,603 s | 78,698 s | 9,382 s | 2,211 s | identique |
| Médiane v4 | **8,586 s** | **79,464 s** | **9,197 s** | **2,275 s** | **3/3** |

Chacune des trois paires compare les 2 704 modules et les 27 sorties publiées,
octet pour octet. Les trois temps backend et les trois temps processus sont
strictement inférieurs à 10 s. L'amorçage, **104,897 s backend**, est exclu ;
il ne constitue pas une mesure à froid de la machine. Frontend et compilation
Go sont exclus du backend comme du temps processus indiqué.

Le run médian, édition 18, optimise deux modules, en émet un et restaure 2 703
sorties. Lookup et gel prennent 1,309 s ; préparation 1,690 s, dont chargement
TAST 1,217 s et preuve compacte 0,454 s ; clés de cache 0,092 s ; transfert des
directives/workers 0,130 s ; optimisation et émission 2,396 s ; exports Go
0,916 s ; packages 0,003 s ; publication 1,359 s ; fermeture de session 0,438 s.
Le producteur PBO, 2,149 s, est inclus dans l'optimisation et l'émission.

La campagne utilise `--skip-go-build`. La compilation Go finale a réussi ; son
[rapport v4](../../../.gopurs-evidence/c2-8-sub10-20261023/v4-final-go-build-result.json)
rattache `v4-final-test-main` au digest des packages de `full-20` et au SHA du
compilateur. `Test.Main` n'a pas été exécuté.

Le [vrai `b` v4](../../../.gopurs-evidence/c2-8-sub10-20261023/normal-b-defaults-v4/result.json)
a réussi avec les scripts b8x identiques aux originaux et sans variable
`GOPURS_*`, `GOGC` ou `GOMEMLIMIT` transmise par l'appelant. Le processus natif
observé active packages/overlay/manifeste à `1`, utilise huit workers et le
pipeline, avec `GOGC=off` et `GOMEMLIMIT=10GiB`, sans daemon. La configuration
expérimentale est restaurée. Ce contrôle isolé prend **8,977 s backend** et
**13,651 s pour `b`**, frontend et `go mod tidy` inclus. Il ne constitue pas une
quatrième paire, ni une mesure de compilation Go ou d'exécution de l'application.

L'[audit de clôture](../../../.gopurs-evidence/c2-8-sub10-20261023/completion-audit.json)
est réussi : ses 25 critères rattachent les sources et le SHA final aux trois
paires, au vrai `b`, à la compilation Go, aux fallbacks, aux daemons et aux
suites de régression. Aucun contrôle de livraison de ce périmètre ne reste ouvert.

## Mesures v3 — avant correction du fallback

Binaire v3 mesuré :
`3ceb9250c378acf7420c4679e28bf426636a2a0745796ba45878af0485b7dcc2`.
L'[état des sources](../../../.gopurs-evidence/c2-8-sub10-20261023/v3-source-state.json)
conserve les HEAD gopurs/PBO et les hashes des fichiers relevés lors de ce build.

| Étape | Backend incrémental | Full forcé | Processus incrémental | Frontend séparé | Oracle |
| --- | ---: | ---: | ---: | ---: | --- |
| Édition 15 | 8,132 s | 77,040 s | 8,663 s | 2,432 s | identique |
| Édition 16 | 8,867 s | 80,193 s | 9,614 s | 2,224 s | identique |
| Édition 17 | 8,506 s | 78,590 s | 9,125 s | 2,369 s | identique |
| Médiane v3 | **8,506 s** | **78,590 s** | **9,125 s** | **2,369 s** | **3/3** |

Les trois incrémentaux respectent chacun le seuil strict de 10 s backend. Le
temps processus est également inférieur à 10 s pour chaque essai ; il ne
comprend ni le frontend ni la compilation Go. Les 2 704 modules et 27 sorties
publiées sont byte-exacts au full pour chacune des trois paires.

L'amorçage v3 a pris **102,373 s backend**, dont 25,497 s de publication initiale
de la preuve.
Cet amorçage est exclu de la mesure incrémentale ; ce n'est pas une mesure à
froid de la machine.

Les [résultats v3](../../../.gopurs-evidence/c2-8-sub10-20261023/candidate-v3/results.json)
et le [journal de campagne](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/candidate-v3.log)
conservent les sept builds, leurs empreintes et les trois comparaisons complètes.
La campagne utilise `--skip-go-build`. La compilation séparée du dernier
résultat, `go build ./Test.Main/main`, a réussi ; son
[rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/final-go-build-result.json)
rattache le binaire produit au digest des packages de `full-17` et au compilateur
v3. `Test.Main` n'a pas été exécuté.

Le run médian, édition 17, optimise deux modules, en émet un et restaure 2 703
sorties. Ses principaux postes sont : lookup et gel 1,247 s ; préparation
1,727 s, dont chargement TAST 1,233 s et preuve compacte 0,475 s ; clés de cache
0,093 s ; transfert des directives/workers 0,128 s ; optimisation et émission
2,371 s ; exports Go 0,977 s ; packages 0,004 s ; publication 1,332 s ; fermeture
de session 0,369 s. Le producteur PBO, 2,127 s, est inclus dans l'optimisation et
l'émission ; ces sous-métriques ne doivent pas être additionnées une seconde fois.

## Mécanismes

La preuve compacte de C2-7 reste l'autorité de préparation. Elle conserve les
types TAST, les déclarations ADT/classes, les arguments statiques et les corps
de dictionnaires effectivement consultés ; elle est ancrée au contexte et à
la génération TAST commitée. C2-8 réduit les lectures, encodages et fichiers
temporaires autour de cette preuve. Le chargement et le décodage du TAST global
restent présents.

### Préparation et métadonnées

- Le gel du TAST produit `source-hashes.json` à partir des octets privés
  effectivement copiés. La préparation n'accepte cette table que pour un
  inventaire complet, sans doublon ni nom manquant ; sinon elle relit les sources
  gelées pour calculer leurs hashes. Le contrôle après gel reste en place.
- Le hash des métadonnées globales peut être repris lorsque la preuve compacte
  établit leur invariance, que le contexte correspond et que tous les mémos
  courants portent le même hash connu. Une preuve absente ou incohérente impose
  l'encodage habituel. Les tables typées restent disponibles pour la génération.
- Le manifeste source/interface reprend une interface seulement pour le même
  source connu et le même contexte/directives. Ce raccourci ne remplace ni les
  clés préparées ni les preuves PBO.
- Le snapshot de base peut réutiliser les observations du lookup qui vient
  d'avoir lieu, si les noms, types et chemins canoniques de la recette concordent.
  La copie privée et le snapshot après gel continuent de vérifier les entrées.

### Clés et lectures PBO

Les contributions dont les clés préparées, demandes et préfixe canonique sont
exactes sont identifiées avant la fermeture des lecteurs. L'égalité du contexte,
des métadonnées et de tous les FFI reste requise. Ces contributions inchangées
arrêtent la propagation d'invalidation : le graphe restant ne contient que les
modules qui ont encore besoin de la preuve par lectures. Une lecture future
absente reste distincte d'une contribution antérieure en attente.

Les objets de lectures vérifiés sont regroupés dans le fichier privé
`prior/owners.json`, au lieu de créer un répertoire et un fichier par module.
`null` signifie preuve inconnue ; une chaîne vide signifie preuve connue vide.
Les entrées absentes restent inconnues. Un index groupé présent mais malformé
est rejeté ; il ne peut pas être masqué par d'anciens fichiers individuels.
L'ancien format individuel reste lisible si l'index groupé est absent.

L'index privé destiné à PureScript omet les champs `Sources` et `Reads` qu'il
ne consomme pas. Ils restent dans le manifeste privé faisant autorité. Les
directives et workers sont transférés par lots depuis leurs objets vérifiés,
sans matérialiser leurs fichiers temporaires. Ce transfert d'octets n'accorde
aucune visibilité anticipée aux contributions : le résolveur garde l'ordre
canonique et la condition de finalisation. Les consommateurs qui ont besoin
de fichiers conservent leur chemin de matérialisation.

Une lecture positive déjà prouvée sur une contribution restaurée peut éviter
de charger son implémentation si son propriétaire est strictement antérieur
et finalisé dans le snapshot de la tentative. Les absences et dépendances en
attente gardent leurs vérifications. Si un module doit finalement être optimisé,
son AST réellement préparé est matérialisé et contrôlé avant conversion.

### Sources Go et publication

Le calcul des signatures de packages réutilise les identités des sources
commitées inchangées et celles des sources fraîchement produites. Les octets
restent vérifiés lors de leur matérialisation. L'overlay central conserve ses
cinq bundles ; les conditions structurelles et les fallbacks de C2-7 restent
applicables.

Les packs épinglés et le manifeste privé continuent d'isoler une session de la
suppression des objets publics. Les sorties héritées sont vérifiées avant
publication. Une erreur `incremental cache:` provoque une seule relance full
forcée. Une preuve inconnue ne devient jamais une preuve connue vide pour
obtenir un hit. Les nouvelles métriques distinguent également la fermeture de
session, comprise dans le temps backend, de la publication.

### Correction du fallback natif dans v4

Le contrôle de corruption de `Generic/3` a reproduit un défaut réel dans v3 :
l'objet corrompu était rejeté, mais l'exception synchrone remontant par
`Aff.liftEffect` provoquait un panic Go avant la relance full. Le
[résultat d'échec](../../../.gopurs-evidence/c2-8-sub10-20261023/v3-native-read-parts-fallback-fixture/results.json)
est conservé. Les mesures centrales v3 restent valides, mais ne suffisaient
donc pas à autoriser sa livraison.

Dans `gopurs-aff/src/Effect/Aff.go`, `runAffSync` convertit désormais ces
exceptions en erreurs Aff et préserve l'identité des erreurs Go. Le callback
`use` de `GeneralBracket` est également évalué sous cette protection, pour
exécuter la libération des ressources avant le gestionnaire d'erreur. Le
[test pérenne](../../gopurs-aff/tools/aff-native.test.mjs) compile le vrai FFI et
vérifie six tests Go et quatre sous-cas de bracket avec `-race`.

V4 passe les corruptions de la partie 3 lue par `read-parts` et de la partie 1
chargée par le résolveur paresseux : une seule relance full, objets réparés,
aucun staging résiduel, deux modules bruts et cinq fichiers générés byte-exacts
au full indépendant, puis compilation Go réussie.

## Protocole

La campagne utilise la copie b8x sous
`htdocs/.gopurs-evidence/a-20261021-overlay/app/b8x` et le
[script de paires](../../../.gopurs-evidence/c2-8-sub10-20261023/measure-central-pairs.py).
Elle alterne trois incrémentaux et leurs fulls forcés, après un amorçage, sans
autre charge lourde. Le SHA du binaire est contrôlé avant et après chaque build ;
le TAST du module édité est identique pour les deux membres d'une paire.

L'appel direct à `gopurs-native` fixe packages, overlay et manifeste à `1`,
préparation/PBO/émission à huit workers, pipeline activé, `GOGC=off`,
`GOMEMLIMIT=10GiB` et `GOWORK=off`. Les modes daemon/résident, le snapshot
compilateur externe et le profilage sont retirés. Le full ajoute
`GOPURS_FORCE_BUILD=1`, désactive l'overlay et doit optimiser/émettre tous les
modules. Ces options explicites reproduisent le backend mesuré ; elles ne
constituent pas à elles seules un test de `b` et des défauts du lanceur.

Pour chaque paire, l'oracle vérifie hashes et tailles des objets commités,
décompresse les sources Go, compare les **2 704 modules** et les **27 sorties**
packages/runtime/entrées, puis vérifie le littéral dans les sources brutes et
publiées. Il ne se limite pas aux fichiers modifiés. Les essais et leurs échecs
restent conservés dans l'évidence.

Le champ `complete` du script signifie que le nombre de paires demandé a
terminé. Il ne vérifie ni un minimum de trois paires ni le seuil de 10 s : ces
deux critères ont été vérifiés séparément sur les résultats v3 et v4, avec chacun
des trois incrémentaux strictement sous 10 s dans chaque campagne. La compilation
Go est contrôlée séparément. `Test.Main` exécuté n'est pas l'oracle : la campagne C2-7 avait
rencontré un blocage dans un test d'intégration antérieur à `Util.spec`.

### Contrôle du vrai `b` avec v3

Le [contrôle du lanceur](../../../.gopurs-evidence/c2-8-sub10-20261023/normal-b-defaults-v3-fixed/result.json)
a réussi avec v3 sur une nouvelle édition littérale, avec les scripts `b`, `build` et
leurs helpers copiés à l'identique depuis b8x. L'appelant ne transmet aucune
variable `GOPURS_*`, `GOGC` ou `GOMEMLIMIT` ; seul le PATH de la toolchain est
ajouté. Le processus natif observé utilise packages/overlay/manifeste à `1`,
huit workers de préparation/PBO/émission, le pipeline, `GOGC=off` et
`GOMEMLIMIT=10GiB`, sans daemon. La cible Go étant sélectionnée, `b` suffit.

Ce contrôle isolé a pris **8,028 s backend** et **11,879 s pour l'appel `b`**,
frontend et `go mod tidy` inclus. Il valide les défauts du lanceur et n'ajoute
pas une quatrième paire à la campagne de performance. La configuration et le
source expérimentaux ont été restaurés.

Le [premier essai](../../../.gopurs-evidence/c2-8-sub10-20261023/normal-b-defaults-v3/result.json)
a échoué avant le backend : le montage temporaire masquait les packages P00.
Le script a été corrigé pour tous les conserver et ne rediriger que le
compilateur. Cet échec de fixture reste archivé ; aucune correction de production
n'a été nécessaire.

## Validations

Validations v3 acquises :

| Vérification | Résultat | Preuve |
| --- | --- | --- |
| Campagne centrale native | 3 paires byte-exactes, chaque incrémental <10 s | [résultats](../../../.gopurs-evidence/c2-8-sub10-20261023/candidate-v3/results.json) |
| Compilation Go de `full-17` | Succès, sans exécution de `Test.Main` | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/final-go-build-result.json) |
| Vrai `b`, options du lanceur par défaut | Succès, backend 8,028 s, appel complet 11,879 s, sans daemon | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/normal-b-defaults-v3-fixed/result.json) |
| Compilation Spago incrémentale | 0 erreur, 0 avertissement dans cette invocation | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-spago.log) |
| Suites ciblées PureScript exécutées sous Node | 27/27, dont 12 tests de préparation persistante | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-ps-tests.log) |
| Go build-cache, `-race` | Succès, 20,574 s | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-cache-race.log) |
| Go C1c, `-race` | Succès, 2,042 s | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-c1c-race.log) |
| Helper des lectures groupées | Succès, 2,017 s | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-read-owners-helper.log) |

Ces suites couvrent notamment la distinction absent/vide/inconnu, les préfixes
exacts, la finalisation canonique, le rejeu paresseux, les types instanciés,
les dictionnaires à deux niveaux, les arguments statiques, les payloads
corrompus et leur réparation. Les deux tests des lectures groupées ont d'abord
échoué sur l'ancien comportement ; le
[journal rouge](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-read-owners-red.log)
est conservé.

Validations acquises pour la correction et v4 :

| Vérification | Résultat | Preuve |
| --- | --- | --- |
| Campagne centrale v4 | 3 paires byte-exactes, chaque incrémental <10 s, médiane 8,586 s | [résultats](../../../.gopurs-evidence/c2-8-sub10-20261023/central-validated-v4/results.json) |
| Compilation Go de `full-20` | Succès, sans exécution de `Test.Main` | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/v4-final-go-build-result.json) |
| Vrai `b` v4, sans exports de configuration | Succès, backend 8,977 s, total 13,651 s, sans daemon | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/normal-b-defaults-v4/result.json) |
| Corruption `Generic/3`, lecture groupée | Fallback unique, réparation, sorties exactes, compilation Go | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/v4-native-read-parts-fallback-fixture/results.json) |
| Corruption `Generic/1`, implémentation paresseuse | Fallback unique, réparation, sorties exactes, compilation Go | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/v4-native-implementations-fallback-fixture/results.json) |
| Aff natif, `-race` | Six tests Go et quatre sous-cas de bracket réussis | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-aff-native-regression.log) |
| Bridge JS/Go | 1/1 | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-bridge.log) |
| Lanceur | 9/9 | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/v3-launcher.log) |
| Snapshot du compilateur | 3/3 | [journal](../../../.gopurs-evidence/c2-8-sub10-20261023/logs/compiler-snapshot-regression.log) |
| Daemon v4, un worker | Sorties exactes au full, erreur récupérable, compilation Go | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/daemon-result.json) |
| Daemon v4, huit workers | Sorties exactes au full, erreur récupérable, compilation Go | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/daemon-jobs-8-result.json) |
| Audit final | 25 critères réussis, binaire et sources concordants | [rapport](../../../.gopurs-evidence/c2-8-sub10-20261023/completion-audit.json) |

Les deux contrôles daemon enchaînent `Int → String → Int`, puis une édition
littérale. Le TAST invalide échoue tout en conservant le PID du daemon ; les
builds suivants réussissent. Les clients concurrents restent hors portée.
Certains journaux de tests conservent leur préfixe historique `v3` ; les rapports
natifs de fallback et de daemon portent explicitement le SHA v4.

Les validations de livraison v4 sont terminées. Les contrôles v3 restent des
preuves historiques propres à l'ancien binaire.

Aucune validation supplémentaire des hôtes JS/Rust, du froid universel ou des
changements structurels n'est déduite de cette campagne centrale.

## Historique et références

C2-7 conserve sa médiane 16,970 s contre 92,934 s full, sur son propre binaire
et ses trois paires. Les candidats intermédiaires C2-8 restent disponibles :
[v1](../../../.gopurs-evidence/c2-8-sub10-20261023/candidate-v1/results.json),
[v2](../../../.gopurs-evidence/c2-8-sub10-20261023/candidate-v2/results.json), et
[profil v1](../../../.gopurs-evidence/c2-8-sub10-20261023/profile-v1/result.json).
Leurs mesures isolées ne sont pas amalgamées aux campagnes complètes v3 et v4.

La baseline officielle du
[README altbak.pub](../../../altbak.pub-gopurs/README.md), section « Compilation
times to Go », donne b8x JS-host 55,775 s, Go-host 33,037 s et Rust-host 22,404 s.
Son protocole diffère de l'édition avec cache amorcé et packages/overlay. Le
full forcé apparié est l'oracle local de cette campagne ; aucun ratio comparable
n'est déduit de la baseline historique.

Les essais utilisent la copie d'évidence. Les projets principaux b8x et
purescript ne sont pas modifiés par cette campagne. Ce document ne remplace ni
les résultats historiques ni les critères encore ouverts de `todo.md`.
