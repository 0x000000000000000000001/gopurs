# Nettoyage de tools/

Audit du 8 novembre 2026 : 175 entrées, dont 102 tests et 54 scripts non-test.
Le but est de réduire le bruit sans toucher aux tests actifs ni aux outils
manuels documentés.

## Plan

- [x] **1. Supprimer les orphelins.**
  - `tools/go-imports-differential.mjs` : aucun appelant, ne finit pas par
    `.test.mjs`, donc jamais exécuté par `node --test tools/*.test.mjs`.
  - `tools/preparation-native.mjs` et `tools/preparation-native.rs` : aucune
    référence, aucune documentation, et collision de nom avec le test actif
    `preparation-native.test.mjs`.
  - Vérifier qu'aucune référence ne subsiste, puis committer.

- [ ] **2. Traiter le cluster historique P00/P01.**
  - Décider du sort des 18 fichiers (~2 537 lignes) actuellement référencés
    uniquement par `docs/build-preparation-baseline.md`,
    `docs/build-preparation-compact.md` et `docs/plan-archive-20261019.md`.
  - Option A : les archiver dans un répertoire d'études clairement identifié.
    Option B : les supprimer et conserver uniquement les preuves JSON/CSV déjà
    publiées dans `docs/`.
  - Ne pas toucher à `preparation-source-variants.mjs`,
    `preparation-transition-oracles.mjs`, `preparation-native.test.mjs` ni
    `preparation-persistent.test.mjs`, qui sont actifs.
  - Mettre à jour les documents qui citent les scripts déplacés ou supprimés.

- [ ] **3. Centraliser la liste des tests rapides.**
  - Extraire la liste dupliquée de `release-quality.mjs`, `ci/receipt.mjs` et
    `ci/receipt.test.mjs` dans une constante unique.
  - Éviter qu'une modification de la sélection rapide passe inaperçue parce que
    les trois copies ont divergé.

- [ ] **4. Rendre `tools/README.md` exhaustif.**
  - Nommer les 54 scripts non-test, ou générer l'inventaire depuis le disque.
  - Ajouter un contrôle qui échoue si un script non-test n'est référencé nulle
    part, avec une liste explicite pour les outils manuels légitimes.

## Règles

- Les 102 fichiers `*.test.mjs` sont actifs : la campagne complète les exécute
  même sans référence textuelle.
- Les outils manuels documentés (`test-native-pbo.mjs`,
  `cache-context-runtime.mjs`, `release-benchmark.mjs`, `pgo.mjs`) restent en
  place.
- Chaque suppression doit être vérifiée par une recherche de références et un
  `git diff --check`.
