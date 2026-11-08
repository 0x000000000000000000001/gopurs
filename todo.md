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

## Lot 3 — Organisation des fichiers

- [ ] Créer les répertoires de domaine pour les scripts non-test :
  - `tools/build/` : `build-native`, `build-rust`, `build-ffi`,
    `embed-runtime`, `native-workspace`, `prepare-native-output` ;
  - `tools/release/` : `release-prepare`, `release-smoke`,
    `release-quality*`, `release-sbom`, `release-benchmark`,
    `release-npm-provenance` ;
  - `tools/lib/` : primitives partagées (voir lot 4) ;
  - `tools/ci/` : déjà en place.
- [ ] Garder les `*.test.mjs` à la racine de `tools/`.
- [ ] Déplacer un domaine à la fois, avec mise à jour de `package.json`, des
  workflows CI, de `bin/`, des imports relatifs et des docs.
- [ ] Vérifier les anciens chemins avec `git grep` et les scripts npm.
- [ ] Conserver les commandes publiques (`npm run build:native`, etc.).

Critère de sortie : `tools/` racine ne contient plus que les tests, les
répertoires de domaine et les lanceurs explicitement documentés.

## Lot 4 — `tools/lib` et pilote de lisibilité

- [ ] Créer un petit socle partagé, sans framework :
  - `cli.mjs` : `parseArgs`, `--help`, `fail` ;
  - `process.mjs` : `run`, timeouts, signaux — réutiliser
    `command-runner.mjs` ;
  - `fs.mjs` : `readJson`, `writeJson`, `tempDir`, `ensureEmpty` ;
  - `repo.mjs` : racine, chemins, Git ;
  - `log.mjs` : étapes, succès, erreurs.
- [ ] Refactoriser trois scripts pilotes seulement :
  `release-quality.mjs`, `native-workspace.mjs`, `test-workspace.mjs`.
- [ ] Leur donner un `main()`, un `--help`, des fonctions nommées et des
  constantes nommées.
- [ ] Ne pas changer leur comportement observable.

Critère de sortie : les trois scripts sont plus courts, testés, et les
primitives sont réutilisées sans duplication.

## Lot 5 — Style et ergonomie

- [ ] Choisir un formateur unique (Prettier ou Biome) pour `tools/`.
- [ ] Ajouter un `format:check` et l'exécuter en CI.
- [ ] Ajouter une charte courte :
  - pas de logique au niveau racine ;
  - pas de ligne à plus de deux effets ;
  - constantes nommées pour les chemins, limites et listes ;
  - erreurs explicites avec la commande et le fichier concernés ;
  - JSDoc sur les fonctions exportées.
- [ ] Ajouter des tests de fumée `--help` pour les outils manuels.
- [ ] Documenter ou supprimer la dépendance cachée à Python 3
  (`preparation-observe.py`).

Critère de sortie : `format:check` passe, les nouveaux scripts suivent la
charte, et les outils manuels ont une aide minimale.

## Lot 6 — Documentation et inventaire

- [ ] Rendre `tools/README.md` exhaustif : une ligne par script non-test avec
  objectif, usage, prérequis et statut (CI, manuel, historique).
- [ ] Ou générer cet inventaire depuis un manifeste et vérifier qu'il est à jour.
- [ ] Ajouter un test d'inventaire qui échoue si un script non-test n'est
  référencé nulle part, avec une liste explicite pour les outils manuels.
- [ ] Vérifier les liens et les commandes de `docs/` après les déplacements.

Critère de sortie : un nouveau mainteneur peut trouver et exécuter chaque outil
sans lire son code.

## Lot 7 — Tests : lisibilité ciblée

- [ ] Ne pas réécrire les 102 tests pour le style.
- [ ] Extraire des helpers communs pour les schémas répétés :
  création de workspace Go, écriture de TAST, lancement de `go test`,
  capture de sortie.
- [ ] Traiter en priorité les fichiers de plus de 500 lignes :
  `test-runner.test.mjs` (865), `preparation-persistent.test.mjs` (739),
  `owned-trees.test.mjs` (572).
- [ ] Découper uniquement lorsqu'un fichier est modifié pour une raison
  fonctionnelle.
- [ ] Garder les assertions exactes et les snapshots inchangés.

Critère de sortie : les tests restent équivalents, mais les scénarios réutilisent
les mêmes helpers.

## Lot 8 — Validation finale

- [ ] `node tools/check-docs.mjs`
- [ ] `npm run test:package`
- [ ] `npm run format:check`
- [ ] `node --test tools/*.test.mjs` (campagne complète, avec l'environnement
  natif requis)
- [ ] `npm run build:native` si un script de build a été déplacé
- [ ] `git diff --check`
- [ ] Vérifier que les commandes publiques et les workflows CI n'ont pas changé
  de contrat.

## Définition de terminé

- Aucun orphelin ni script d'étude non maintenu dans `tools/`.
- Une seule source pour la liste des tests rapides.
- Tous les scripts non-test sont classés et documentés.
- Les `*.test.mjs` restent à plat et la campagne complète passe.
- Les trois scripts pilotes suivent la charte et utilisent `tools/lib`.
- Un test d'inventaire empêche la réapparition du problème.

## Décisions ouvertes

1. Formateur : Prettier ou Biome ?
2. Organisation : déplacer build/release dès maintenant, ou seulement après
   les lots 2 et 3 ?
