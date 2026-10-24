# C2-10 — réduction du backend central, cible de 2 s non atteinte

**Mesures finales du 24 octobre 2026.** Le candidat L donne **2,335 / 2,347 /
2,401 s backend**, soit une médiane de **2,347 s**, contre **3,841 s pour
D/C2-9** : **1,494 s de moins, soit 38,90 %**. Les trois paires sont identiques
octet pour octet à leur full forcé sur 2 704 modules et 27 sorties publiées.
L'étendue des mesures est de 66 ms, nettoyage inclus. **Le seuil strict de 2 s
n'est pas atteint.** Le vrai `b` prend **6,836 s**, dont **2,277 s backend**,
sans daemon ni exports manuels. La campagne est **close pour impasse pratique
documentée** : les sept suites et quatorze gates de l’audit passent ; seul le
gate de vitesse échoue. L’audit global reste **failed, code de sortie 1**.

La portée reste l'édition d'un littéral du corps de
`Util.Type.Variant.Encoding.Encoding.Test.DecodeJsonWith`, avec cache disque
amorcé, dans un processus neuf. Le temps backend inclut vérifications,
publication et nettoyage synchrone du staging. Le frontend PureScript, la
comparaison externe et la compilation Go sont séparés. Les changements de
types, de structure ou de FFI gardent leurs chemins de validation et de repli ;
ils ne bénéficient pas d'une promesse de temps identique.

## Utilisation et identité

Le `b` habituel utilise le lanceur Go natif, qui active déjà packages, overlay
et manifeste sans exports manuels. Les valeurs explicites, dont `0`, restent
respectées. Le daemon demeure opt-in ; aucun daemon ni processus résident
n'intervient dans ces mesures. La politique existante de la machine de 48 Gio
reste huit workers et `GOGC=off`, `GOMEMLIMIT=10GiB`. C2-10 ne change pas ces
défauts et n'active pas ces chemins par défaut pour les hôtes JS/Rust.

Binaire L :
`69c8e28afed6fa2d435e4ca4e06d38d9ab3baaa0c5c5d12c8338789af564831d`.
Le [bootstrap audité](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-build.json)
relie le binaire installé à sa copie conservée et aux
[sources capturées](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-source-state.json).
Les 67 fichiers modifiés ou nouveaux de gopurs et le fichier PBO concordent
après construction ; Aff est inchangé par rapport au commit capturé.
Les éditions documentaires ultérieures ne font pas partie du compilateur.

## Protocole et mesures finales

Trois paires alternées incrémental/full utilisent le même binaire et trois
TAST distincts, sans autre charge lourde. Chaque full force la préparation et
l'émission de 2 704 modules, sans réutilisation ni overlay. Les comparaisons
portent sur les SHA et longueurs vérifiés des sources Go brutes, FFI comprise,
et des 27 fichiers Go publiés. Les capsules facultatives et leurs objets ne
remplacent jamais cet oracle de sortie.

| Édition | Backend incrémental | Full forcé | Processus incrémental | Frontend séparé | Oracle |
| --- | ---: | ---: | ---: | ---: | --- |
| 66 | 2,335 s | 76,269 s | 2,813 s | 1,944 s | 2 704 + 27 exacts |
| 67 | 2,347 s | 76,724 s | 2,977 s | 1,924 s | 2 704 + 27 exacts |
| 68 | 2,401 s | 76,882 s | 2,921 s | 1,998 s | 2 704 + 27 exacts |
| Médiane | **2,347 s** | **76,724 s** | **2,921 s** | **1,944 s** | Trois paires exactes |

Preuves : [paire 66](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-pair66/results.json)
et [campagne finale 67/68](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-final67-68/results.json).
L'amorçage, la première création des capsules et l'édition de chauffe 65 sont
séparés des trois mesures. Le full intercalé doit conserver les capsules
encore valides ; la mesure suivante reste un processus neuf.

Le profilage, le daemon, le résident et `GOPURS_KEEP_STAGE` sont désactivés dans
les paires. Les inventaires avant/après doivent être identiques : les stages
anciens conservés pour les profils ne sont pas pris pour des fuites du run.
Les trois paires satisfont ce contrôle. Les profils qui conservent leur stage
excluent son nettoyage ; ils servent à l'attribution, pas à remplacer les
chronomètres des paires.

La baseline officielle du [README altbak.pub](../../../altbak.pub/README.md),
section « Compilation times … to Go », demeure **33,037 s** pour b8x sur
l'hôte Go. Son protocole diffère de cette édition avec cache amorcé ; aucun
ratio comparable n'en est déduit. La comparaison de campagne pertinente
reste [D/C2-9](build-c2-sub4.md), avec ses trois paires et sa médiane de 3,841 s.

## Changements depuis D

- **Moins de copies pendant les vérifications Go.** Les lectures qui ne
  consomment qu'un hash travaillent en flux ; les lectures utilisées par le
  gel gardent les octets nécessaires. Les snapshots complets et leurs erreurs
  restent observés. Empreintes du compilateur et des entrées se chevauchent,
  avec jointure avant résultat et priorité d'erreur déterministe. Les écritures
  du pack privé réservent des tranches disjointes sous verrou puis utilisent
  `WriteAt` hors verrou ; tout le lot finit avant index, et tout échec détruit
  le stage incomplet.
- **Préparation typée à la demande.** Le schéma 5 ajoute une capsule de lectures
  de `globalTypes`, avec couverture explicite et ancrages de la projection.
  Les onze autres tables, les interfaces et les en-têtes restent complets.
  Une clé hors couverture force la table complète issue des huit fragments
  vérifiés ; « pas encore chargé » ne devient jamais « absent ». Codecs,
  monomorphisation et publication qui exigent le complet le forcent. Les
  informations de types, ADT et classes du TAST sont conservées.
- **Moins de travail PBO sur les résultats restaurés.** Les en-têtes prouvés
  sont admis au même point dans l'ordre du builder ; dépendances et enfants
  sont construits pour les modules encore incertains. Les capsules par lecteur
  ancrent les résultats de lectures aux SHA des parties 1 et au contexte.
  Une nouvelle lecture non couverte utilise le résolveur complet ; une valeur
  fraîche prime sur une capsule. Les quatre parties obligatoires restent
  inchangées. Les workers antérieurs ne sont décodés que pour les propriétaires
  visibles demandés dans le snapshot du lot.
- **Chevauchements privés et bornés.** Après un vrai miss natif, la projection
  peut être chargée pendant le gel ; elle n'est admise qu'après les contrôles
  de la session finale. Les octets de partie 2 sont lus pendant la préparation,
  puis décodés au début du gel FFI. Les gardes tardives de session, de sources,
  de callers et de FFI restent décisives. Chaque fibre et ses workers sont
  joints avant fermeture, annulation achevée ou retry. Les voies non admissibles
  et les hôtes JS/Rust conservent leur comportement séquentiel.
- **Publication et exports Go.** Le parsing des imports utilise les octets déjà
  vérifiés ; les lectures restent bornées et ordonnées. L'index de packs privé
  réutilise la structure qu'il vient d'écrire. Le compactage prépare ses groupes
  en parallèle, joint toutes les tâches et installe l'index durable avant la
  nouvelle map ; corruption et réparation ne permettent pas d'éviction sur un
  ancien descripteur. Enfin, les exports gardent leur contrôle des fichiers,
  mais la publication transporte leurs références attendues privées jusqu'à
  `stage` : le SHA et la longueur des octets réellement stockés sont vérifiés
  avant installation, au lieu d'une lecture de contrôle supplémentaire.

Les capsules sont facultatives, liées à des références committées et à leurs
ancrages. Leur absence ou corruption entraîne un miss local ; une partie
obligatoire corrompue reste une erreur qui impose le repli complet. Un full
intercalé ne conserve une capsule qu'avec ses ancrages toujours valides.
Les snapshots d'entrée, les barrières avant/après sorties publiques, le journal
d'appartenance et le pointeur committé en dernier ne sont pas supprimés.

## Impasse pratique du périmètre actuel

Le run médian consacre 503 ms au lookup/gel, 156 ms à la préparation, 121 ms au
FFI, 535 ms à optimisation/émission dont 385 ms au producteur PBO, 164 ms aux
exports, 668 ms à la publication et 71 ms à la fermeture. Les mesures incluses
ou chevauchées ne s'additionnent pas. Une lecture/indexation à 18 ms ne permet
notamment pas de présenter les 164 ms d'exports comme du travail anticipable.

Les sondes ont servi à sélectionner des changements limités, puis à rejeter
les pistes sans plafond suffisant. Elles ne constituent pas des gains backend
additionnables :

| Piste restante | Observation et décision |
| --- | --- |
| GC explicite à deux frontières | Trois paires privées : médiane 2,410 → 2,620 s, six oracles exacts. Rejet ; politique par défaut conservée. [Preuve](../../../.gopurs-evidence/c2-10-sub2-20261023/gc-boundaries-k/attempt2/results.json) |
| Ordonnancement dynamique des lectures | Lookup/gel H : 419 ms avec répartition fixe à 8, 428 ms en dynamique 8, 475 ms en dynamique 16. Aucun changement de scheduler. [Preuve](../../../.gopurs-evidence/c2-10-sub2-20261023/scheduling-compact-h-probe/scheduling-h48-collected.json) |
| Vérification groupée des objets hérités | La phase modules du staging K ne prend que 34,9 ms, sous le seuil de sonde convenu de 40 ms. Groupement non exécuté et non intégré. [Attribution](../../../.gopurs-evidence/c2-10-sub2-20261023/stage-inherited-k-probe/RESULTS.md) |
| Nettoyage parallèle final | Sur clones privés de 391 entrées, seulement environ 3 ms gagnées ; pas de changement du nettoyage synchrone. [Preuve](../../../.gopurs-evidence/c2-10-sub2-20261023/cleanup-k-probe/result.json) |
| Publication | Replay K : environ 606 ms, dont trois snapshots 138 ms, stage 149 ms et retrait 112 ms, postes imbriqués. La seule double lecture de packages supprimée a gagné 22–38 ms sur trois paires isolées ; son gain backend propre n'est pas établi. [Attribution](../../../.gopurs-evidence/c2-10-sub2-20261023/publication-k-probe/RESULTS.md), [Expected](../../../.gopurs-evidence/c2-10-sub2-20261023/expected-outputs-k-probe/expected-outputs-k64-v2.json) |
| Réutilisation du parse de manifeste | Les deux lectures encadrent un commit et portent sur des SHA différents. Les deux décodages totalisent environ 25 ms ; réutiliser le premier n'aide pas le delta. Réutiliser le nouveau manifeste demanderait une preuve supplémentaire pour un seul décodage : priorité basse, aucune sonde ni intégration. |

Le [bilan d’impasse local](../../../.gopurs-evidence/c2-10-sub2-20261023/impasse-assessment.md)
relie aussi les attributions PBO aux 347 ms encore manquantes. Ces plafonds
étayent une impasse pratique pour de nouvelles petites retouches
sans réduire les garanties. Ils ne démontrent pas un plancher théorique, ni
l'impossibilité d'une autre architecture. Un pack d'entrées persistant, une
modification des barrières ou un daemon seraient des chantiers distincts ;
ils ne sont pas présentés comme la suite automatique de ces résultats.

## Validations finales et verdict

| Contrôle | État à cette révision |
| --- | --- |
| Bootstrap et liaison sources/binaire L | Réussis, rapport d'identité ci-dessus |
| Trois paires centrales L | Réussies : 2 704 sources brutes + 27 sorties exactes pour chacune, trois TAST distincts, aucun stage ajouté |
| Admission HeaderPreload dans le vrai Driver L | 14 branches, cinq oracles, hit/force/jobs1/sans manifeste et corruption pertinente ou non : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-header-gating-fixture/results.json) |
| Contrôleur deux phases | Spago sans erreur/avertissement, 87 tests JS, vrai parcours PS→TAST→Go avec annulation/erreur et workers 1/8 : [commandes](../../../.gopurs-evidence/c2-10-sub2-20261023/header-two-phase-validation-commands.json) |
| Publication Expected | Ciblés et suite cache Go `-race` réussis ; octets consommés, inventaire privé, symlinks, retry, jointure et JSON inchangé : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/expected-outputs-production-review/validation-results.json) |
| Suites finales liées au candidat L | Sept suites réussies : Spago, 87 tests PS, cache et C1c `-race`, bridge, launcher et compiler-snapshot ; [rapport final de reprise](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-validation-fixture-resume/validation-commands.json) |
| Partie 1 corrompue L | Repli complet, sources et sorties exactes, stages vides : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-implementations-fixture/results.json) |
| Partie 2 corrompue L | Lecture des headers et repli complet vérifiés : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-headers-v2-fixture/results.json) |
| Partie 3 corrompue L | Lecture effective de `Generic/3`, repli, réparation et sorties exactes avec overlay désactivé dans toute cette fixture : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-read-parts-v5-fixture/results.json) |
| Projection absente/corrompue L | Préparation complète, sources brutes exactes et valeurs publiques exécutées ; différences de packages conservées : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-projection-fixture/results.json) |
| Capsule metadata et nouvelle demande L | Full intercalé, absence/corruption, nouvelle demande et extension de couverture sans préparation globale ; sources et valeurs publiques exactes : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-metadata-capsule-fixture/results.json) |
| Capsule binding L | Création, full intercalé, perte/corruption, couverture réduite injectée et nouvelles lectures : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-binding-capsule-fixture/results.json) |
| Daemons 1/8 workers | Contrôles de compatibilité réussis, sans processus laissé actif : [1 worker](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-daemon-1-result.json), [8 workers](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-daemon-8-result.json), [provenance d’exécution](../../../.gopurs-evidence/c2-10-sub2-20261023/validation-l/remaining-execution-v7.json) |
| Compilation Go finale | `go build ./Test.Main/main` réussi en 21,942 s sur le full 68, entrées inchangées ; compilation seule : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-final-go-build-result.json) |
| Vrai `b` L | 6,836 s mural, 2,277 s backend ; options automatiques observées, aucun daemon, configuration restaurée : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-normal-b/result.json) |
| Audit final | **14 gates PASS, seul `three_strict_pairs_and_variability` FAILED** car aucun run ne passe sous 2 000 ms ; global **failed / exit 1** : [rapport](../../../.gopurs-evidence/c2-10-sub2-20261023/completion-audit.json), [exécution et SHA](../../../.gopurs-evidence/c2-10-sub2-20261023/final-audit-execution.json) |

Le contrôle metadata L observe `Generic.alternate` absent puis présent dans
la couverture, la valeur publique attendue et les sources exactes, sans
préparation globale. Il ne trace pas directement le lookup ni le forçage dans
CodeGen. Le comportement d’une clé hors couverture est démontré séparément
par les [tests `LookupTable` sur le vrai parcours PS→Go](../../../.gopurs-evidence/c2-10-sub2-20261023/globaltypes-lookup-prototype/native-ps/result.json).

La [première exécution des suites](../../../.gopurs-evidence/c2-10-sub2-20261023/candidate-l-validation-final/validation-commands.json)
conserve quatre succès — Spago, 87 tests PS, cache Go et C1c `-race` — puis
l’échec du bridge : sa fixture Go ne copiait pas le runtime requis par le FFI
actuel. La correction ajoute cette copie dans `tools/build-cache.test.mjs`,
qui n’est pas embarqué dans le compilateur. Le bootstrap L, sa capture
immuable, le binaire et les trois paires ne changent pas.

Le [manifeste d’écart du seul test](../../../.gopurs-evidence/c2-10-sub2-20261023/bridge-test-runtime-fix/test-only-override.json)
fixe les identités avant/après et le diff ; son statut `prepared` décrit cette
définition immuable, pas le verdict d’exécution. Le rapport final ci-dessus
reprend les quatre succès après vérification exacte de leurs preuves, puis
valide bridge, launcher et compiler-snapshot en 1,494 s, 5,227 s et 13,235 s
respectivement. Sources strictement vérifiées avant/après, unique écart de
test explicite et binaire inchangé : le lot de sept suites est **passed**.
Le [verdict de cette correction de fixture](../../../.gopurs-evidence/c2-10-sub2-20261023/bridge-test-runtime-fix/validation.json)
est **passed** ; le rapport initial en échec reste conservé.

Le test de partie 3 utilise un corpus monomorphe, une fonction `never` et un
label modifié pour imposer la lecture du worker corrompu. L'overlay y vaut `0`
dès le prime et dans tous les runs : le contrôle porte sur lecture/repli,
avec égalité des sources brutes **et** des sorties publiées. Le réglage par défaut avec overlay
activé reste couvert par les trois paires centrales et le gating HeaderPreload.
Les essais rouges de fixtures sont conservés ; les ajustements du corpus et
de l'oracle n'ont demandé aucune correction de production.

L'[outil de réduction de couverture L](../../../.gopurs-evidence/c2-10-sub2-20261023/capsule-narrower-l/build-report.json)
est validé avec les vrais codecs natifs du candidat. Il prépare une injection
explicite de capsule valide à couverture réduite. Le contrôle binding L utilise
cette injection pour prouver le traitement d’une nouvelle lecture ; il ne
prétend pas qu’une édition naturelle produit cette couverture réduite.

Le vrai `b` a été lancé dans une copie isolée avec les scripts identiques,
sans variables d’activation fournies par l’appelant. Son temps mural comprend
le frontend et les helpers ; la compilation Go de l’application reste séparée.
Ce contrôle est distinct des trois paires finales : il ne constitue pas une
quatrième mesure et ne change pas leur médiane de 2,347 s. Les durées des
contrôles binding/daemon sont des temps de validation, pas des benchmarks.

Les petites fixtures peuvent publier des dispositions de packages différentes
entre overlay et full. Leur oracle conserve alors les différences de fichiers,
exige les mêmes sources Go brutes et exécute les valeurs publiques via le plan
vérifié. Il ne déclare pas les packages byte-exacts. Les paires centrales,
elles, exigent les deux égalités. `Test.Main` n'est pas un oracle d'exécution
réussi de cette campagne ; la compilation Go et les fixtures exécutables sont
des contrôles distincts.

L’[audit strict de clôture](../../../.gopurs-evidence/c2-10-sub2-20261023/completion-audit.json)
a été exécuté avec son seuil inchangé. Les quatorze gates de qualité et de
provenance sont **PASS** ; `three_strict_pairs_and_variability` est **FAILED**
car les trois temps dépassent 2 000 ms. Le rapport global est **failed** et
l’exécution retourne **1**. L’écart-type d’échantillon est de 35,157 ms,
le coefficient de variation de 1,489 %, sur une étendue de 66 ms.

Aucune validation ne reste en attente ni aucun défaut fonctionnel identifié
non résolu dans ce périmètre. La clôture pour impasse pratique s’appuie sur
les preuves de qualité et le [bilan des plafonds](../../../.gopurs-evidence/c2-10-sub2-20261023/impasse-assessment.md) ;
elle ne transforme pas l’objectif de 2 s en succès et ne prouve pas un plancher
théorique. Les preuves restent sous
`.gopurs-evidence/c2-10-sub2-20261023/` depuis `htdocs`.
Les essais n'ont pas modifié les projets b8x/purescript principaux. Aucun
commit, push ou merge n'a été effectué par l'assistant.

Lors du dernier nettoyage après inventaire, seuls les répertoires `fixture` et
`tmp` du diagnostic PBO `complete2` ont été retirés. Rapports, logs, sources,
binaire et métadonnées archivées sont conservés ; la conservation vérifiée
est décrite dans [le relevé disque](../../../.gopurs-evidence/c2-10-sub2-20261023/disk-recovery-pbo-final-20261024/after.json).
Ce nettoyage n'est pas une nouvelle mesure ni une suppression de preuve rouge.
