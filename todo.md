# Gopurs — plan de maintenabilité

Plan v1 — mise à jour : 30 septembre 2026.

Objectif : un backend clair, lisible et modifiable par responsabilités, avec
une génération Go identique. Périmètre : compilateur, runtime, FFI et outillage
de ce dépôt ; les interfaces avec PBO et les bibliothèques sœurs sont incluses.
Les optimisations de performance restent en pause.

## Avancement : 30 % — 30/100 points, 4/15 lots validés

- **Calcul :** somme des points des cases cochées, sur un total fixe de 100.
  Les poids sont des unités de suivi ; le pourcentage mesure les livrables
  validés de ce plan.
- Un lot peut demander plusieurs passes courtes. Ses points sont acquis
  lorsque son critère de fin et les vérifications communes sont satisfaits.
- À chaque clôture : cocher le lot, actualiser ce compteur et la date, puis
  consigner les preuves dans [docs/testing.md](docs/testing.md).
- Scinder un lot conserve son poids total. Une extension du périmètre donne
  lieu à une nouvelle version du plan.

## Lots — ordre de travail

- [x] **01 — Pilote de compilation (10 pts).** `Main` réduit au lancement ;
  phases, configuration, préparation, build et sorties répartis dans `Driver`.
- [x] **02 — Pipeline et erreurs CLI (10 pts).** Durée de vie de l'émetteur et
  des workers encadrée ; nettoyage, diagnostic unique et statut d'échec validés
  en JS, natif séquentiel et natif parallèle.
- [x] **03 — Outillage de build et de tests (5 pts).** Bootstrap, préparation
  du workspace et gestion des sous-processus séparés ; runner partagé,
  interruptions et publication atomique du binaire validés.
- [x] **04 — Contexte de traduction et annotations (5 pts).** `ExprContext`
  utilisé directement par le dispatcher ; annotations et dictionnaires de
  classes isolés dans `TypedExprs` ; parité vérifiée.
- [ ] **05 — Bridge FFI (10 pts).** Dans `FfiBridge`, séparer les décisions de
  signature, l'adaptation des arguments/résultats et le rendu des wrappers.
  Fin : chaque règle a un emplacement nommé et les chemins de secours sont explicites.
- [ ] **06 — Choix des représentations (5 pts).** Revoir `GoTypes`,
  `ConstructorLayout` et les métadonnées ADT/classes.
  Fin : décisions de layout et d'instanciation regroupées, invariants documentés.
- [ ] **07 — Boxing, conversions et Rebox (10 pts).** Clarifier les cas de
  `GoConversions` et les dépendances entre helpers.
  Fin : choix de conversion, enregistrement et émission transitive lisibles séparément.
- [ ] **08 — Expressions restantes du dispatcher (5 pts).** Extraire de
  `CodeGen` la traduction des littéraux composites et des constructeurs saturés.
  Fin : dispatcher court, ordre d'évaluation et de coercition explicite dans les émetteurs.
- [ ] **09 — Bindings, fonctions et TCO (5 pts).** Revoir `ModuleBindings`,
  `BindingExprs`, `FunctionExprs` et leurs analyses partagées.
  Fin : signatures, captures, initialisation récursive et sauts TCO faciles à suivre.
- [ ] **10 — Fusions et applications immédiates (5 pts).** Revoir
  `ThunkFusion`, `FunctionFusion` et `ImmediateApplications`.
  Fin : reconnaissance, conditions d'admission et réécriture distinguées et nommées.
- [ ] **11 — Intrinsics et traversées (5 pts).** Revoir `ArrayIntrinsics`,
  `ArrayTraverse` et `ObjectTraverse`.
  Fin : reconnaissance, capture ordonnée des arguments et émission des boucles identifiables.
- [ ] **12 — Analyses spécialisées (10 pts).** Revoir successivement
  `Ownership`, `BorrowedObjects`, `ClosedDictionaries`, `NativeRecordArgs`
  et `DecoderSchemas`. Fin : pour chaque passe, preuves d'admission, transformation
  et déclarations produites ont des responsabilités et des contrats explicites.
- [ ] **13 — Frontière TAST/PBO (5 pts).** Revoir `Monomorphization`,
  `Preparation` et les informations typées consommées par le backend.
  Fin : provenance des métadonnées, barrières et ordre des passes documentés au point d'usage.
- [ ] **14 — Runtime et FFI JS/Go du compilateur (5 pts).** Revoir les
  frontières entre code PureScript, FFI et `runtime/runtime.go`.
  Fin : responsabilités, durée de vie des valeurs et contrats communs aux deux backends explicites.
- [ ] **15 — Consolidation finale (5 pts).** Sur une même révision, reconstruire
  les deux compilateurs, exécuter les campagnes fixtures/modules et vérifier la
  parité b8x. Fin : résultats et exclusions justifiés, README et architecture
  cohérents, liens documentaires réparés, statut de l'installation Nix renseigné.

## Vérifications communes pour clôturer un lot

- Responsabilités et noms relus ; code mort et doublons repérés traités.
  Une revue peut confirmer qu'une partie est déjà conforme, avec justification.
- Builds et tests ciblés adaptés aux fichiers modifiés réussis ; snapshots
  stricts et exécution Go pour les comportements concernés.
- Pour une modification de génération : mêmes entrées TAST, comparaison octet
  par octet avec la référence puis entre natif séquentiel, natif parallèle et JS,
  y compris l'inventaire des fichiers. Pour l'outillage : contrats de processus
  et bootstrap réel. Pour le runtime : tests de comportement et de durée de vie.
- Documentation des responsabilités et preuves de validation à jour ;
  `git diff --check` propre.

**Prochaine passe : lot 05 — bridge FFI.** Sa clôture portera l'avancement à **40 %**.
