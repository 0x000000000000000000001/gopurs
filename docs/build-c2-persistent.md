# C2-7 — build incrémental one-shot sans daemon

Suite : [C2-8 — backend central sous 10 s](build-c2-sub10.md), médiane v4 de
8,586 s sur trois paires byte-exactes, fallback natif corrigé et livraison
validée. Le vrai `b` v4 prend 13,651 s au total, dont 8,977 s backend.
Les mesures C2-7 ci-dessous conservent leur binaire et leur protocole historiques.

**Cible ≤20 s atteinte sur les trois paires natives du 23 octobre 2026 :
16,970 s de backend médian, contre 92,934 s en full forcé.** Préparation
persistante compacte, contributions PBO à la demande et overlay multi-bundles
sont implémentés. Le lanceur `bin/gopurs` les active désormais par défaut pour
l'hôte Go natif ; JS et Rust restent opt-in. Les trois oracles sont byte-exacts.

Le cas mesuré est une édition de littéral dans le corps de
`src/Util/Type/Variant/Encoding/Test/DecodeJsonWith.purs` de b8x, dont le module
s'appelle `Util.Type.Variant.Encoding.Encoding.Test.DecodeJsonWith`.
Ce n'est pas une garantie de temps pour toute modification de ce module : les
changements structurels, de type ou de spécialisation peuvent exiger une
préparation complète. Le chemin séquentiel et le daemon restent disponibles.

## Mesures natives finales

Campagne : `central-validated-pairs/results.json` dans l'évidence ci-dessous.
Binaire : `457bfcf3c5b0913113518c94066d10fdc0625f804f66f5c7ed42ba515da802d0`.

| Cas | Incrémental packages/overlay | Full forcé | Oracle complet |
| --- | ---: | ---: | --- |
| Édition 6 | 16,970 s | 92,934 s | identique |
| Édition 7 | 17,293 s | 90,813 s | identique |
| Édition 8 | 16,946 s | 94,393 s | identique |
| Médiane | **16,970 s** | **92,934 s** | **3/3** |

Chaque backend est un processus neuf, sans daemon ni watch. Le frontend prend
2,040–2,142 s et reste hors du temps backend, comme la compilation Go.
L'amorçage du nouveau binaire, **108,679 s**, est exclu ; il inclut 27,213 s de
publication initiale de la preuve compacte. Ce n'est pas un résultat à froid de
la machine. Le chemin rapide mesuré suppose ce cache amorcé.

Chaque incrémental vérifie `reads decoded=2704/2704`, prépare un module,
en optimise deux, en émet un et restaure 2 703 sorties. Les cinq bundles restent
en place. Le run médian consacre 2,058 s à la préparation, chargement compris ;
1,630 s aux clés/preuves et 3,483 s au producteur PBO, inclus dans les 5,139 s
`optimize + emit`. La matérialisation d'un seul lecteur coûte 5 ms ; exports,
restauration et overlay Go prennent 3,048 s ; publication 1,606 s.

Les trois paires finales utilisent `--skip-go-build`, puis une compilation Go
séparée du dernier résultat, réussie en **22,940 s** (`go build ./Test.Main/main`).
Deux paires intermédiaires avaient aussi compilé
chaque sortie. La campagne finale conserve le même binaire pour ses sept builds.

Le full utilise `GOPURS_FORCE_BUILD=1`, désactive l'overlay et optimise/émet les
2 704 modules. L'oracle vérifie les objets du manifeste commité, leurs hashes,
la décompression des sources Go, puis compare les 2 704 modules et les 27
fichiers de packages/runtime/entrées octet pour octet. Il vérifie aussi que le
nouveau littéral apparaît dans le Go brut et dans les packages publiés.

`Test.Main` n'est pas utilisé comme oracle d'exécution : le lancement complet
s'est bloqué dans un test d'intégration antérieur à `Util.spec`. Les compilations
Go et la comparaison exhaustive sont des vérifications distinctes.

## Utilisation habituelle et reproduction de la mesure

Dans b8x sur la cible Go, **`b` suffit** : il passe par `bin/gopurs`, qui active
packages, overlay et manifeste sans daemon. Les valeurs explicites de
`GOPURS_GO_PACKAGES`, `GOPURS_GO_PACKAGE_OVERLAY` et `GOPURS_BUILD_MANIFEST`
restent prioritaires, notamment `0` pour désactiver une option. À partir de
32 Gio de RAM, le lanceur choisit déjà 8 workers de préparation/PBO et la
politique GC utilisée lors des mesures, sauf surcharge explicite.

Pour reproduire les mesures en appelant directement le binaire natif, après
production du TAST, les options restent explicites puisque cet appel contourne
le lanceur :

```sh
export GOPURS_GO_PACKAGES=1
export GOPURS_GO_PACKAGE_OVERLAY=1
export GOPURS_BUILD_MANIFEST=1
export GOPURS_PREPARE_JOBS=8
export GOPURS_PBO_JOBS=8
export GOPURS_EMIT_JOBS=8
export GOPURS_PIPELINE=1
export GOGC=off
export GOMEMLIMIT=10GiB
export GOWORK=off
unset GOPURS_DAEMON GOPURS_DAEMON_SERVER GOPURS_RESIDENT GOPURS_RESIDENT_TRIGGER
unset GOPURS_COMPILER_SNAPSHOT GOPURS_FORCE_BUILD
htdocs/gopurs/gopurs/bin/gopurs-native
```

Une préparation complète publie l'état de référence. Un changement de
compilateur/contexte ou un delta non couvert impose une nouvelle préparation
complète. La préparation compacte exige le manifeste et plus d'un
worker PBO. Avec un seul worker, le chemin de préparation complet est conservé.
L'activation par défaut concerne uniquement l'hôte Go natif du lanceur ; elle
ne constitue pas une validation supplémentaire des autres hôtes.

## Preuve de préparation

Les seuls hashes source/interface ne suffisent pas. Le contre-exemple corrigé
est `Generic.ignore :: forall a. a -> Int` : passer son appel de `ignore 1` à
`ignore "hello"` change la spécialisation de `Generic`, même si son source et
l'interface de l'appelant sont inchangés. Des lectures PBO vides ne prouvent pas
la préparation qui les précède.

`Gopurs.Driver.Prepare.Proof` et `Persistent` séparent désormais :

- les clés de spécialisation, nécessaires à la réécriture des appels ;
- les payloads propres à chaque module, nécessaires à ses clones ;
- les corps réellement consultés par la spécialisation transitive, la génération
  des clones et la résolution des dictionnaires, y compris à travers leurs alias.

La forme vérifiée conserve types, structure, noms, imports/exports, déclarations
ADT/classes, annotations utiles et littéraux des patterns. Les positions source
et annotations d'usage sont retirées de cette comparaison ; parmi les valeurs,
seules celles des expressions scalaires dynamiques sont effacées. Les arguments
dynamiques des payloads persistés sont normalisés en placeholders conservant
leur annotation de type d'origine, même absente, leur position et leur nombre.
Les arguments statiques ne sont pas normalisés.

La preuve est liée au contexte/directives et à la génération TAST du manifeste
commité. Un delta d'un seul module n'est admis que si cette forme reste identique
et qu'aucun corps modifié n'a été consommé par la spécialisation. Le module changé
est réellement préparé ; les autres empreintes sont réutilisées sous cette
preuve. Si un lecteur doit finalement être optimisé, son vrai module préparé
est matérialisé et son hash comparé avant de poursuivre. Un placeholder source
n'est jamais optimisé comme s'il s'agissait du module préparé.

Le cache `output/.gopurs-cache/preparation-literals-v1` utilise une preuve
schema 2, des clés hashées et des payloads propres adressés par contenu. Sur b8x :
preuve 3,45 Mo, dont index 1,47 Mo ; payloads 79,87 Mo ; 6 780 empreintes de corps
consultés. L'oracle de normalisation a comparé les 2 704 modules préparés au full,
octet pour octet. Un index vide conserve les exports d'origine et le caractère
privé des bindings non exportés.

Un échec de preuve revient à la préparation complète avec le TAST déjà chargé.
Un payload absent/corrompu est rejeté et réparé lors de la publication complète.
Le full forcé recalcule toujours la préparation ; il peut seulement éviter de
réencoder une preuve déjà validée pour la même génération, après vérification
de tous ses payloads.

## Réutilisation PBO et émission

Une restauration rapide nécessite un contexte identique, ainsi que l'égalité
des métadonnées globales et de tous les FFI avec les preuves commitées, plus soit
les clés complètes exactes (préparation, demandes et préfixe d'optimisation), soit
la preuve par lectures. Cette dernière invalide
la fermeture inverse des lecteurs des modules préparés modifiés, en incluant les
observations inconnues. Sa traversée visite chaque module une fois.

Les lectures autoritaires proviennent des objets vérifiés du manifeste commité.
Le miroir plat `reads-v1/<module>.json` reste diagnostique. Les observations
absentes diffèrent des observations connues vides ; le codec
`gopurs/reads-owners-v1` conserve cette distinction.

En parallèle, le cache fige le manifeste/index et épingle les packs par des liens
privés. Il charge directives, implémentations, workers et Go seulement quand le
consommateur en a besoin. La suppression des packs publics par un GC concurrent
ne détruit pas cette génération privée. Les objets hérités sont vérifiés avant
publication ; ils ne sont pas réencodés un par un dans le staging.

Le résolveur PBO garde l'ordre canonique : seules les contributions déjà
finalisées et antérieures sont visibles. Les workers d'émission sont validés une
fois par module, puis réintroduits dans l'ordre original des lots. Aucun seed
global arbitraire de `globalFunctions` n'est utilisé.

Une erreur signalée comme `incremental cache:` annule la session et relance une
seule fois un full forcé. La publication vérifie aussi les objets déjà indexés
avant réemploi ; un objet packé corrompu est réécrit avec les données fraîches,
même si une copie individuelle valide existe. Si le pack est remplacé au même
chemin, son ancien descripteur est fermé après remplacement atomique.

## Overlay central

Le module central est réparti dans cinq bundles. Le chemin `retained-5` conserve
ce découpage et réécrit les sections concernées. Il vérifie source brute engagée,
signatures, imports, arêtes et tokens, en préservant les renommages de records du
planificateur. Il accepte les modifications de littéraux couvertes par cette
preuve ; toute ambiguïté ou modification non admise revient au plan complet.
Voir [le contrat multi-bundles](build-c2-go-packages.md).

## Validation et limites

- Build Spago réussi, sans nouvelle erreur ni nouvel avertissement du cache ;
  les 36 avertissements préexistants restent présents.
- Suites ciblées : préparation, index/monomorphisation, préfixes exacts,
  contributions paresseuses, lectures PBO, workers et snapshots compilateur.
- Préparation persistante : 10 tests, dont `Int → String`, arguments
  statiques, dictionnaires à deux niveaux, payload corrompu/réparé et index vide
  avec binding privé non typé.
- Suites Go `-race` pour build-cache et C1c. La corruption de pack avec édition
  TAST, le fallback complet puis le hit suivant sont couverts ; les défauts de
  réparation ont été reproduits avant correction.
- Daemon final vérifié avec 1 et 8 workers : `Int → String → Int`, puis édition
  littérale, sorties byte-exactes au full et compilation Go. Une entrée TAST
  invalide échoue sans tuer le daemon ; le build suivant réussit. Les daemons de
  test sont arrêtés. Les appels clients concurrents restent non supportés.

Les preuves détaillées sont sous :
`htdocs/.gopurs-evidence/c2-7-reprise-20261023/`.
Scripts principaux : `measure-central-pairs.py`, `check-daemon.py` ; résultats
`central-validated-pairs/results.json`, `daemon-result.json`,
`daemon-jobs-8-result.json`, `compact-normalized-oracle.json`.

## Historique et références

| Étape | Backend / résultat | Portée |
| --- | --- | --- |
| Ancien Ping 29,3 s | preuve de préparation insuffisante | invalidé par l'appel polymorphe |
| Reprise sûre Ping, `7a325ed1…` | 37,293 / 44,877 / 36,888 s | 3 oracles d'exécution Ping, cible non atteinte |
| Central avant préparation compacte | 71,765 s | mesure unique de contrôle, pas une médiane |
| Premier central compact, `104b36b3…` | 68,660 s ; full 134,363 s | 2 704 modules + 27 sorties byte-exacts |
| Version avant réparation des packs | 17,680 / 18,115 / 16,947 s | 2 paires complètes exactes ; campagne interrompue avant le dernier full |

Les expériences rejetées restent rejetées : snapshot complet `PreparedCache`
264 Mo/160,7 s, réutilisation fondée seulement sur source/interface, slice dont
le hash incluait les scalaires dynamiques, et seed global des workers. Le nouvel
état compact repose sur un contrat distinct et un oracle exhaustif.

La baseline officielle `altbak.pub/README.md`, section « Compilation times to
Go », donne b8x JS-host 55,775 s, Go-host 33,037 s et Rust-host 22,404 s. Son
protocole diffère de l'édition one-shot packages/overlay ; aucun ratio comparable
n'est déduit de ces chiffres historiques.

Aucun commit, push ou merge. Aucun changement dans les worktrees principaux b8x
et purescript ; les essais utilisent une copie d'évidence. Les changements PBO
restent dans `purescript-backend-optimizer-gopurs`.
