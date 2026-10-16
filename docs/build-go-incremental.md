# Cache multipackage — G05 validé

G05 est validé le **14 octobre 2026**, suite G à **6/7**, P à **0/7**.
Le layout est sélectionné par
`GOPURS_GO_PACKAGES=1`. Ses fichiers publics sont déjà publiés sous le journal
commun : fragments par propriétaire, imports d'ordre, initialisation et entrées.

Le [layout retenu en G06](build-go-measurements.md) regroupe ensuite les sources
ordinaires par package, à densité **1/8192**. La restauration utilise toujours les
sources privées avant partition. L'[activation et la migration](build-cache.md#sélectionner-le-découpage-go-g06)
s'effectuent dans la racine de propriété existante.

## Sources privées avant partition

Le mémo historique restaurait les Go `module`/`ffi` depuis les sorties possédées.
Le layout multipackage publie des fragments qualifiés, qui ne permettent pas de
reconstituer directement le module avant partition. G05 ajoute au mémo un
inventaire optionnel `sources` : `module`, puis `ffi` si présente. Chaque source
porte le codec **`gopurs/go-source-gzip-v1`**, l'objet compressé et l'empreinte/taille
des octets bruts.

Le producteur utilise gzip à compression rapide, sans date ni nom de fichier.
Il conserve les quatre objets métier existants. Les anciennes sorties
monolithiques restent restaurables via le manifeste historique. Une source
compressée est un objet privé ; elle ne crée ni chemin Go public supplémentaire
ni droit de supprimer un fichier.

`publishSession` rattache les sources du staging au mémo ; `stageModuleSource`
réutilise un objet déjà vérifié lorsque son contenu brut est identique. Les
barrières de publication restent celles du store commun. `freezeModuleCache`
vérifie l'objet, décompresse dans la taille annoncée, contrôle le contenu brut
et écrit une copie privée en lecture seule. Un objet absent, une corruption,
un mauvais checksum gzip, une taille/empreinte brute incohérente ou un codec
inconnu invalident le candidat. Les fragments publics ne servent jamais de
source de restauration.

Les clés préparée/PBO/émetteur continuent à décider de la réutilisation métier.
Les sources restaurées entrent dans le même résolveur typé, avec les exports
natifs figés du build courant. Ce résolveur recalcule les placements, qualifications
et ABI de records : un ancien placement ne devient pas une autorité du cache.

## Validations

La preuve de départ `module-sources-before.log` échoue avec un index de candidats
vide après publication multipackage. Après changement, les tests avec détection
de races passent, y compris sous **Go 1.22.12** : restauration des octets avant
qualification, copies privées, corruptions, sources manquantes et atomicité.
Les tests couvrent les transitions monolithe/scission/fusion/suppression et la
mort d'un processus après journalisation d'un fragment imbriqué.

Le harnais des sources vérifie maintenant les propriétaires du manifeste plutôt
que des chemins monolithiques. Ses quatorze transitions passent sur JS et Go
séquentiel/parallèle et Rust : **56 comparaisons / 350 commandes**, corps
privés/exportés, directives, constructeurs, records,
dictionnaires, ajout/substitution/rétraction de spécialisations, renommage et
retour exact à l'état initial. Il compare aussi les nouveaux objets compressés et
trace les reconstructions Go identiques sans compilation/lien.

Le premier candidat Rust est rejeté : la copie du cache Cargo conservait un script
de build compilé avec l'ancien `CARGO_MANIFEST_DIR`, donc un chemin d'archive Go
G03. La reconstruction utilise désormais un target Cargo neuf et vérifie le
chemin natif émis par le script. Le candidat 2 passe ensuite les quatorze
transitions sur les quatre hôtes et la reprise du pilote Rust. Le candidat
défectueux et les traces sont gardés.

La revue finale ajoute la récupération d'un flux gzip invalide dont l'objet et
le manifeste ont été rehashés. La publication revalide maintenant les octets
décompressés avant de reprendre un objet antérieur ; sinon elle le réencode.
`source-repair-before.log` conserve l'échec, `source-repair-after.log` la suite
passée avec détection de races sous Go 1.22.12. Le cache corrigé est intégré au
candidat 1 G06, dont les **89 commandes CLI** passent sur les quatre hôtes.

Sur b8x, les **2 997 sources privées** totalisent **151 556 052 octets bruts** et
**12 641 367 octets compressés uniques**. Une édition de corps FFI reprend les
**2 704 backends**, restaure **2 703 émissions** et ne réémet qu'une FFI. Un seul
Go public change. Le forcé et le retour à la base reproduisent les octets, les
quatre parties et les sources compressées ; l'oracle compilé passe. Chaque
reconstruction Go identique suivante ne compile ni ne lie. La campagne conserve
**14 commandes** et l'inode du verrou initial, dans la même racine de propriété.

Le runner b8x passe ses dix contrôles, dont les nouvelles transitions de layout,
les dépendances transitives et deux entrées. Il continue à laisser Go valider le
graphe de dépendances de chaque invocation ; sa clé racine/entrée reste adaptée.

Preuves : `gopurs-go-packages-g05-20261014/`, notamment `before/`,
`module-sources-before.log`, `module-sources-after-1.log`, `cache-go122-1.log`,
`candidate-1-rejection.json`, `source-transitions-1.log`, `cache-recovery-1.log`
et `runner-1.log`. `final-reference/summary.json` rassemble les artefacts, sources,
compteurs et empreintes. Les neuf snapshots G04 sont inchangés à cette clôture.

## Stockage : référence G03 et bilan G06

Le relevé statique du candidat G03 donne **156 309 879 octets de Go** contre
**133 809 190** dans le corpus monolithique de comparaison, soit **+16,82 %**.
Ce relevé dépasse le budget de +10 % et est traité dans les
[mesures G06](build-go-measurements.md), avant leur acceptation.
Il ne constitue ni une mesure de temps de compilation ni le bilan final des
objets gopurs, du GOCACHE ou des exécutables.

Le candidat 7 mesuré sur **36 chaînes** ramène le Go b8x à **122 994 214 octets**
sur la base, soit **−8,08 %** face au monolithe. Les quatre catégories de stockage
passent leur budget sur les six scénarios. La marge la plus faible est celle du
GOCACHE total hôte + runner identique : **+9,997897 %**, environ **23 976 octets**
sous le plafond +10 % du corpus figé.

Après publication multipackage réussie, `retirePackageObjects` reprend le verrou
stable et relit les racines courantes. Il conserve les objets encore référencés ;
une publication incomplète suspend l'entretien. Un échec d'entretien laisse des
objets récupérables sans annuler le commit réussi. Les tests couvrent la reprise
du verrou, la publication concurrente et la conservation de ces racines.
