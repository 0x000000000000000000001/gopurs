# Gopurs — fiabilité et reproductibilité

Plan v2 — mise à jour : 7 octobre 2026.

Le plan v1 de maintenabilité est clôturé à **100/100 points, 15/15 lots**.
Sa [validation finale](docs/testing.md#consolidation-finale--lot-15-2-octobre-2026)
conserve les preuves et les limites à l'origine de ce nouveau plan.

Objectif : résoudre les écarts de comportement identifiés, compléter la couverture
utile et rendre les campagnes et l'installation reproductibles. Périmètre :
compilateur, runtime, FFI, outillage et bibliothèques sœurs ; interventions dans
PBO et le frontend TAST limitées aux besoins démontrés.

## Avancement : 100 % — 100/100 points, 8/8 lots validés

- **Calcul :** somme des points des cases cochées, sur un total fixe de 100.
- Un lot peut demander plusieurs passes ; aucun point partiel n'est acquis.
  Son critère de fin et les vérifications communes doivent être satisfaits.
- À chaque clôture : cocher le lot, actualiser le compteur et la date, puis
  consigner les preuves dans [docs/testing.md](docs/testing.md), sous le plan v2.
- Scinder un lot conserve son poids total ; étendre le périmètre exige une
  nouvelle version du plan.

## Lots — ordre de travail

- [x] **01 — Campagnes de tests reproductibles (15 pts).** Isoler les temporaires
  Spago et les nettoyages des runners frères ; corriger l'initialisation de
  `spec`, dépendante d'un répertoire ignoré. Fin : démarrage dans une copie
  fraîche, reprise des échecs et bilan complet vérifiés, sans altérer les autres
  checkouts.
- [x] **02 — Décodage des chaînes dans le TAST (15 pts).** Corriger le décodage
  natif des symboles et labels de records exposé par `StringEdgeCases`.
  Fin : fixture réintégrée, exécution et parité JS/natif validées.
- [x] **03 — Sémantique des chaînes UTF-16 (15 pts).** Résoudre la divergence
  de pliage de `StringEscapes`, notamment la concaténation de deux moitiés de
  surrogate. Fin : assertion réactivée, fixture réintégrée et résultat validé
  en JS, natif séquentiel et natif parallèle.
- [x] **04 — Bornes et opérations sur Int (10 pts).** Établir le comportement
  attendu aux bornes 32 bits, puis corriger l'écart confirmé par `2136`.
  Fin : fixture réintégrée, négation et opérations voisines concernées couvertes.
- [x] **05 — Affichage des Number (10 pts).** Comparer `NumberLiterals` au
  comportement JS de référence et résoudre l'écart de `Show Number`.
  Fin : oracle justifié, fixture réintégrée et cas limites pertinents validés.
- [x] **06 — Dérivations rejetées par le frontend (10 pts).** Comparer les quatre
  fixtures exclues au fork et à la référence amont ; corriger les incompatibilités
  confirmées. Fin : pour chaque cas, prise en charge validée ou limitation de
  version établie par une reproduction comparative et documentée.
- [x] **07 — Couverture des bibliothèques (15 pts).** Ajouter une suite Go
  autonome pour l'adaptation QuickCheck ; examiner les trois `pending` de
  `spec`, compléter les lacunes réelles et expliciter les cas intentionnels.
  Fin : contrôles utiles intégrés au parcours standard et exécutés avec succès.
- [x] **08 — Installation et validation finale (10 pts).** Vérifier le parcours
  depuis des checkouts frais et exécuter `nix flake check` et `nix develop` dans
  un environnement équipé. Fin : compilateurs JS et natif Go reconstruits sur
  les mêmes sources, campagnes fixtures/modules et parité b8x validées ;
  résultats, dépendances locales et limites documentés.

## Vérifications communes pour clôturer un lot

- Chaque correction de comportement dispose d'un test reproduisant le défaut
  avant correction, puis réussi après correction.
- Builds et tests ciblés adaptés aux fichiers modifiés réussis ; exécution Go
  et snapshots stricts pour les comportements concernés.
- Après évolution du runtime ou des FFI, reconstruire aussi l'hôte utilisé dans
  le checkout actif et vérifier le build d'une application consommatrice.
- Pour une modification de génération : mêmes entrées TAST, inventaires et
  comparaison octet par octet entre JS, natif Go séquentiel et natif Go parallèle.
  Les écarts avec la référence précédente doivent être expliqués par la correction.
- Examiner les changements de Go généré avant de mettre à jour un snapshot,
  puis le revérifier strictement. Lever une exclusion après validation effective.
- Pour l'outillage : isolation, reprise, statuts et sous-processus vérifiés avec
  les vrais outils. Pour le runtime : comportement et durée de vie vérifiés.
- Preuves identifiées par plan et lot dans `docs/testing.md`, documentation à
  jour et `git diff --check` propre.

**Plan v2 clôturé le 7 octobre 2026.** La
[validation finale](docs/testing.md#plan-v2--lot-08--installation-et-validation-finale-7-octobre-2026)
consigne les 400 fixtures, les 51 runners, la parité b8x et la vérification Nix.
Les optimisations de performance restent en pause ; toute reprise exige un gain
mesuré significatif.
