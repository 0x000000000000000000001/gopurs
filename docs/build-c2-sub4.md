# C2-9 — backend one-shot central sous 4 s

Le candidat D atteint **3,841 s backend médian**, avec trois mesures
**3,872 / 3,841 / 3,817 s**. Chaque paire est identique octet pour octet à son
full forcé sur les 2 704 modules et les 27 fichiers publiés. Le vrai `b` prend
**8,791 s au total**, dont **3,826 s backend**, avec les options automatiques et
sans daemon. Le binaire validé est installé dans `bin/gopurs-native`.

La portée reste une édition de littéral dans le corps du module central
`Util.Type.Variant.Encoding.Encoding.Test.DecodeJsonWith`, avec cache amorcé,
dans un processus neuf. Le backend inclut vérifications, publication et
fermeture synchrone du staging. Frontend PureScript, comparaison externe et
compilation Go sont mesurés séparément. Ce résultat ne promet pas un `b`
complet sous 4 s, ni le même temps après changement de type, structure ou FFI.

## Mesures et identité

Binaire D :
`13160b1ab6689bee0bf930e4292e841b67d0e4c7ed4aab2ce538bc04ae2ab742`.
Le [rapport de bootstrap](../../../.gopurs-evidence/c2-9-sub4-20261023/candidate-d-build.json)
relie le binaire installé, sa copie conservée et les
[sources capturées](../../../.gopurs-evidence/c2-9-sub4-20261023/candidate-d-source-state.json).
Les 29 fichiers modifiés ou nouveaux de gopurs et le fichier PBO ont été
vérifiés après construction ; Aff est inchangé par rapport au commit de base.
Les changements documentaires ultérieurs ne font pas partie du compilateur.

| Édition | Backend incrémental | Full forcé | Processus incrémental | Frontend séparé |
| --- | ---: | ---: | ---: | ---: |
| 31 | 3,872 s | 78,896 s | 4,434 s | 2,518 s |
| 32 | 3,841 s | 77,158 s | 4,374 s | 1,904 s |
| 33 | 3,817 s | 78,111 s | 4,432 s | 2,026 s |
| Médiane | **3,841 s** | **78,111 s** | **4,432 s** | **2,026 s** |

Les campagnes [première paire](../../../.gopurs-evidence/c2-9-sub4-20261023/candidate-d/results.json)
et [deux paires suivantes](../../../.gopurs-evidence/c2-9-sub4-20261023/candidate-d-final/results.json)
utilisent le même SHA. Elles alternent incrémental/full, sans autre charge
lourde. Les trois TAST édités sont distincts et chaque full optimise et émet
2 704 modules sans réutilisation. L'étendue des temps incrémentaux est de
55 ms. L'amorçage de 103,590 s est exclu ; il ne représente pas un froid universel.

Le protocole conserve huit workers de préparation/PBO/émission, pipeline,
packages/overlay/manifeste, `GOGC=off` et `GOMEMLIMIT=10GiB`. Profilage, daemon,
résident et conservation du staging sont désactivés. Les paires 32/33 vérifient
aussi les inventaires de stages avant/après : aucun nouveau stage ne subsiste,
les anciens profils conservés restant intacts. La paire 31 enregistre la
fermeture réussie dans ses métriques, sans ce relevé supplémentaire.

Le run médian consacre 481 ms au lookup/gel, 694 ms à la préparation,
83 ms aux clés, 121 ms aux en-têtes, 1 312 ms à l'optimisation/émission,
234 ms aux exports, 4 ms aux packages, 753 ms à la publication et 62 ms à la
fermeture. Le producteur PBO de 1 082 ms est inclus dans l'optimisation/émission.

Par rapport aux trois paires C2-8, la médiane backend baisse de **8,586 s à
3,841 s**, soit environ **55 %**. Le full forcé médian est de 78,111 s contre
79,464 s auparavant ; aucune régression de ce full apparié n'est observée.
Cela ne constitue pas une nouvelle campagne de froid complet.

La baseline officielle du [README altbak.pub](../../../altbak.pub-gopurs/README.md),
section « Compilation times to Go », reste b8x JS-host 55,775 s, Go-host
33,037 s et Rust-host 22,404 s. Son protocole diffère de cette édition avec
cache amorcé ; aucun ratio de performance comparable n'en est déduit.

## Changements

- Le calcul des empreintes et le gel des entrées sont parallélisés avec une
  limite fixe. Le TAST est conservé dans un pack privé vérifié, puis matérialisé
  par propriétaire à la demande. Les clients anciens et les replis complets
  disposent toujours du chemin de matérialisation intégrale.
- La préparation commence par les hashes exacts des sources gelées. Elle
  recharge les corps nécessaires à la demande et rejoue les demandes jusqu'à
  stabilisation. Une projection compacte conserve les métadonnées typées TAST,
  les interfaces, les layouts ADT/classes et les autres informations requises.
  Elle utilise huit fragments au maximum, décodés en parallèle et fusionnés
  dans l'ordre. Le schéma 4 vérifie hashes, inventaire, ordre et absence de
  collisions ; une preuve absente ou invalide impose la préparation complète.
- Les signatures des workers sont chargées par propriétaire dans un snapshot
  immuable de lot, seulement lorsqu'elles sont nécessaires. Le cas central
  décode 49 contributions sur 2 704, avec les contrôles de visibilité conservés.
- Les lectures PBO antérieures servent d'indices d'ordonnancement pour éviter
  des tentatives prématurées. Après un miss confirmé, les implémentations
  antérieures admissibles sont décodées en parallèle. Les octets sont vérifiés
  avant lancement et tout le lot est validé avant insertion dans le mémo privé.
  Rang, finalisation et visibilité restent contrôlés par le builder ; les
  nouveaux accès utilisent le résolveur paresseux habituel.
- L'overlay réutilise ses octets vérifiés et évite les copies de grandes tables
  de noms. Après preuve exacte d'un changement limité aux littéraux et
  concordance avec le manifeste gelé, il limite le scan des noms au module
  modifié. Les autres cas utilisent l'index complet. Hashes, imports, sections,
  placement et contrôle avant publication restent vérifiés.

Les expériences isolées ne servent pas de résultats finaux. Les candidats A,
B et C et leurs profils restent dans l'évidence. Le passage de C à D est
confirmé sur le backend : premier essai C 4,433 s, puis les trois paires D
ci-dessus. Aucun daemon supplémentaire ni nouvelle option utilisateur n'est
nécessaire.

## Validations

| Contrôle | Résultat et preuve |
| --- | --- |
| Trois paires centrales | 2 704 modules et 27 sorties byte-exacts, chaque backend <4 s ; campagnes ci-dessus |
| Go du dernier full | Compilation réussie, entrées vérifiées avant/après et SHA du binaire conservé : [rapport](../../../.gopurs-evidence/c2-9-sub4-20261023/candidate-d-final-go-build-result.json) |
| Vrai `b` | Options automatiques, sans daemon, backend 3,826 s / total 8,791 s, configuration restaurée : [rapport](../../../.gopurs-evidence/c2-9-sub4-20261023/candidate-d-normal-b/result.json) |
| Partie 3 corrompue | Une relance full, réparation, stages vides, sources brutes et packages exacts, Go compilé : [rapport](../../../.gopurs-evidence/c2-9-sub4-20261023/d-native-read-parts-fallback-fixture/results.json) |
| Partie 1 corrompue | Même validation sur le résolveur d'implémentations : [rapport](../../../.gopurs-evidence/c2-9-sub4-20261023/d-native-implementations-fallback-fixture/results.json) |
| Projection absente puis corrompue | Préparation complète, fragments réparés, sources brutes exactes, valeurs publiques exécutées et Go compilé : [rapport](../../../.gopurs-evidence/c2-9-sub4-20261023/d-native-projection-fallback-execution-fixture/results.json) |
| Daemon, 1 et 8 workers | `Int → String → Int`, littéral et récupération après TAST invalide, sorties exactes et Go compilé : [1 worker](../../../.gopurs-evidence/c2-9-sub4-20261023/d-daemon-1-result.json), [8 workers](../../../.gopurs-evidence/c2-9-sub4-20261023/d-daemon-8-result.json) |
| Suites de régression | Spago 0 erreur / 36 avertissements préexistants ; 54 tests ciblés ; cache Go et C1c avec `-race` ; bridge 1, lanceur 9, snapshot 3 tests réussis : [commandes et hashes des logs](../../../.gopurs-evidence/c2-9-sub4-20261023/validation-commands.json) |

Le premier [contrôle de projection](../../../.gopurs-evidence/c2-9-sub4-20261023/d-native-projection-fallback-fixture/results.json)
a échoué sur l'exigence d'égalité des packages : l'overlay autorisé déplace
`Caller`, alors que le full le regroupe dans `p_root`. Les deux sources Go
brutes étaient déjà identiques. Le contrôle corrigé conserve l'overlay et les
différences de fichiers dans son rapport (`generated_go_byte_exact=false`).
Il sélectionne le package actif à partir du plan et de son overlay vérifiés,
puis exécute `answer=42` et le label attendu dans warm, chacun des replis et
chacun des fulls. `public_values_exact=true` constitue l'oracle complémentaire ;
aucune égalité des packages n'est revendiquée pour cette petite fixture.
Les trois campagnes centrales restent strictement byte-exactes sur les packages.

`b` a été exécuté sans variables `GOPURS_*`, `GOGC` ou `GOMEMLIMIT` fournies par
l'appelant. Son temps total comprend le frontend et les helpers, notamment
`go mod tidy` ; ce contrôle n'inclut pas la compilation Go de l'application.
La compilation Go finale est une vérification séparée. `Test.Main` n'a pas été
exécuté : le blocage antérieur d'un test d'intégration n'est pas présenté comme
un succès d'exécution de l'application.

L'[audit de clôture](../../../.gopurs-evidence/c2-9-sub4-20261023/completion-audit.json)
relie sources courantes, binaire, snapshots, trois paires, Go final, vrai `b`,
replis, daemons et suites. Les résultats C2-8 restent historiques dans
[leur document](build-c2-sub10.md). La campagne ne modifie pas les projets
b8x/purescript principaux ; les essais utilisent des copies d'évidence. Aucun commit, push ou
merge n'a été effectué par l'assistant.
