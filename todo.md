# Programme de cleanup de `tools/`

État au 8 novembre 2026 : 172 entrées à la racine de `tools/`, dont 102 tests
`*.test.mjs` et 52 scripts non-test `.mjs` (environ 23 800 lignes de `.mjs`,
dont 17 700 lignes de tests). Le code fonctionne et la couverture est réelle,
mais l'ensemble n'est pas encore lisible ni facilement maintenable.

Objectif : réduire le nombre d'entrées, séparer les tests des outils, supprimer
les études mortes, imposer quelques conventions et rendre l'inventaire fiable.

## Principes

- Réduire avant de reformater : ne pas refactoriser du code qui va disparaître.
- Pas de réécriture globale des 102 tests ni des 24 000 lignes.
- Ne pas changer de langage : bash pour les lanceurs, Node pour l'orchestration,
  Go pour les outils natifs, PureScript pour le compilateur.
- Garder les `*.test.mjs` à plat : `node --test tools/*.test.mjs` reste le
  point d'entrée canonique.
- Un lot = un diff relisible, des tests verts et un retour arrière possible.

## Lot 0 — Baseline et garde-fous

- [ ] Consigner le commit de départ et l'inventaire exact :
  `git rev-parse HEAD`, nombre de tests, nombre de scripts non-test.
- [ ] Vérifier que la campagne rapide et les tests ciblés passent avant de
  toucher à l'organisation.
- [ ] Ajouter une règle : aucun nouveau script non-test sans entrée dans
  `tools/README.md`, `--help` et test de fumée.
- [ ] Interdire les nouveaux scripts d'étude directement dans `tools/` :
  ils vont dans un répertoire d'archives ou de documentation.

## Lot 1 — Réduire le poids mort

### 1.1 Orphelins — terminé

- [x] Supprimer `tools/go-imports-differential.mjs`.
- [x] Supprimer `tools/preparation-native.mjs` et
  `tools/preparation-native.rs`.
- [x] Vérifier qu'aucune référence ne subsiste et committer.

### 1.2 Cluster historique P00/P01 — archivé

- [x] Archiver les 18 fichiers (~2 537 lignes) dans
  `docs/history/preparation-p00/tools/`, avec un README d'archive.
- [x] Remplacer les références `tools/preparation-*` par les chemins d'archive
  dans `docs/build-preparation-baseline.md` et
  `docs/build-preparation-compact.md`.
- [x] Mettre à jour `tools/README.md` pour ne plus présenter ces scripts comme
  des outils actifs.
- [x] Ne pas toucher à `preparation-source-variants.mjs`,
  `preparation-transition-oracles.mjs`, `preparation-native.test.mjs`,
  `preparation-persistent.test.mjs` ni `preparation.test.mjs`, qui sont actifs.
- [x] Vérifier que `tools/` ne contient plus de script d'étude non maintenu.

Critère de sortie : plus aucun script d'étude non maintenu dans `tools/`, et
aucune référence cassée dans `docs/`.

## Lot 2 — Config unique et exports morts — terminé

- [x] Extraire la liste des tests rapides dans `tools/ci/gates.mjs` et l'utiliser
  depuis `release-quality.mjs`, `ci/receipt.mjs` et `ci/receipt.test.mjs`.
- [x] Vérifier que chaque nom de la liste correspond à un fichier existant.
- [x] Supprimer ou rendre privés les exports inutilisés :
  `check-docs.mjs`, `pgo.mjs`, `test-workspace.mjs`.
- [x] Vérifier les exports avec une recherche de type `grep` avant suppression.

Critère de sortie : une seule source pour la sélection rapide, et plus d'export
public sans consommateur.

## Lot 3 — Organisation des fichiers — terminé pour build/release

- [x] Créer `tools/build/` et y déplacer `build-native`, `build-rust`,
  `build-ffi`, `embed-runtime`, `native-workspace`, `prepare-native-output` et
  `pgo`.
- [x] Créer `tools/release/` et y déplacer `release-prepare`, `release-smoke`,
  `release-quality*`, `release-sbom`, `release-benchmark` et
  `release-npm-provenance`.
- [x] Créer `tools/lib/` : primitives partagées (voir lot 4).
- [x] `tools/ci/` : déjà en place.
- [x] Garder les `*.test.mjs` à la racine de `tools/`.
- [x] Déplacer un domaine à la fois : build puis release.
- [x] Mettre à jour `package.json`, les workflows CI, `nix/gopurs.nix`, les
  imports relatifs, les tests et les docs pour build et release.
- [x] Vérifier les anciens chemins avec `git grep` pour build et release.
- [x] Conserver les commandes publiques (`npm run build:native`,
  `npm run release:quality`, etc.).

Critère de sortie : `tools/` racine ne contient plus que les tests, les
répertoires de domaine et les lanceurs explicitement documentés.

## Lot 4 — `tools/lib` et pilote de lisibilité — partiel

- [x] Créer un petit socle partagé, sans framework :
  - `cli.mjs` : `showUsage`, `fail` ;
  - `process.mjs` : `CommandRunner`, `Interrupted` (déplacé depuis
    `command-runner.mjs`) ;
  - `fs.mjs` : `readJson`, `ensureDirectory`, `writeJsonAtomic` ;
  - `repo.mjs` : `repoRoot` ;
  - `log.mjs` : `step`, `info`.
- [x] Refactoriser trois scripts pilotes :
  `release-quality.mjs`, `native-workspace.mjs`, `test-workspace.mjs`.
- [ ] Extraire le bloc d'exécution de `release-quality.mjs` dans un `main()`.
  Les deux autres pilotes sont des modules bibliothèques et n'ont pas de CLI.
- [x] Ne pas changer leur comportement observable : 144 tests ciblés passent.

Critère de sortie : les trois scripts utilisent les primitives, les tests
passent, et il reste seulement l'extraction optionnelle de `main()`.

## Lot 5 — Style et ergonomie — partiel

- [x] Choisir Prettier comme formateur unique pour les outils non-test.
- [x] Ajouter `format` / `format:check` et exécuter `format:check` en CI.
- [x] Documenter les conventions dans `tools/README.md` et ajouter les JSDoc
  des helpers `tools/lib`.
- [ ] Appliquer la charte aux scripts restants : logique racine, lignes à plus
  de deux effets, constantes nommées.
- [x] Ajouter un test de fumée `--help` pour les outils manuels.
- [x] Dépendance Python : `preparation-observe.py` est archivé avec le cluster
  P00/P01, il n'y a plus de dépendance active.

Critère de sortie : `format:check` passe, les outils manuels ont une aide
minimale, et la charte est appliquée aux scripts actifs.

## Lot 6 — Documentation et inventaire — terminé

- [x] Rendre `tools/README.md` exhaustif pour les outils gérés : scripts
  top-level, `lib/`, `build/`, `release/` et `ci/`.
- [x] Inventaire dynamique plutôt qu'un manifeste séparé.
- [x] Ajouter `tools/tools-inventory.test.mjs`, qui échoue si un outil non-test
  géré n'est pas nommé dans `tools/README.md`.
- [x] Vérifier les liens et les commandes de `docs/` après les déplacements
  (`check-docs`).

Critère de sortie : un nouveau mainteneur peut trouver et exécuter chaque outil
sans lire son code.

## Lot 7 — Tests : lisibilité ciblée — partiel

- [x] Ne pas réécrire les 102 tests pour le style.
- [x] Extraire `tools/lib/go.mjs` : répertoire temporaire, module Go, écriture
  de fichiers et exécution de commandes Go. Pilotes : `go-imports`,
  `struct-pointer-boxing`, `runtime-contracts`.
- [ ] Étendre les helpers aux schémas TAST restants.
- [ ] Traiter en priorité les fichiers de plus de 500 lignes :
  `test-runner.test.mjs` (865), `preparation-persistent.test.mjs` (739),
  `owned-trees.test.mjs` (572).
- [ ] Découper uniquement lorsqu'un fichier est modifié pour une raison
  fonctionnelle.
- [ ] Garder les assertions exactes et les snapshots inchangés.

Critère de sortie : les tests restent équivalents, mais les scénarios réutilisent
les mêmes helpers.

## Lot 8 — Validation finale — partiel

- [x] `node tools/check-docs.mjs`
- [x] `npm run test:package`
- [x] `npm run format:check`
- [ ] `node --test tools/*.test.mjs` (campagne complète, avec l'environnement
  natif requis)
- [ ] `npm run build:native` si un script de build a été déplacé
- [x] `git diff --check`
- [x] Vérifier que les commandes publiques et les workflows CI n'ont pas changé
  de contrat.

## Définition de terminé

- Aucun orphelin ni script d'étude non maintenu dans `tools/`.
- Une seule source pour la liste des tests rapides.
- Tous les scripts non-test sont classés et documentés.
- Les `*.test.mjs` restent à plat et la campagne complète passe.
- Les trois scripts pilotes suivent la charte et utilisent `tools/lib`.
- Un test d'inventaire empêche la réapparition du problème.

## Décisions ouvertes

Aucune : Prettier est retenu et build/release sont déplacés.
