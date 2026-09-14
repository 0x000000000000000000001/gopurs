# Gopurs — deuxième vague de cleanup

Objectif : simplifier les décisions et les représentations internes après le découpage de la première vague. Cinq gros lots de code, dans l'ordre ci-dessous, puis une clôture. Chaque lot est livré comme un ensemble cohérent.

Point de départ vérifié le 14 septembre 2026 : première vague terminée ; compilation complète des 438 modules sans erreur ni avertissement ; altbak validé avec 387 fichiers Go et 14 résultats fonctionnels inchangés ; les 28 assertions d'ArrayRoundtrip et son snapshot passent. L'architecture actuelle est décrite dans [docs/architecture.md](docs/architecture.md). L'ancien journal jusqu'au lot 12 reste consultable avec `git show baa1e071:todo.md` ; la réexécution d'ArrayRoundtrip est consignée dans [docs/array-roundtrip.md](docs/array-roundtrip.md).

## Règles de livraison

- Maintenir le comportement : même représentation des valeurs, mêmes appels, mêmes conversions et même ordre d'évaluation. Les types TAST, `dataDecls`, `classDecls`, `ForAll`, `TypeApp`, contraintes, ordre des champs et queues de rangées restent accessibles aux décisions du backend.
- Prendre une référence avant chaque lot et comparer le Go sur les mêmes entrées TAST et les mêmes FFI. Conserver les chemins nécessaires aux bibliothèques dans les montages isolés.
- Valider chaque gros lot par `bin/go/run -c` depuis `altbak.pub`, puis comparer le Go généré et les résultats fonctionnels. Expliquer tout écart ; une mise à jour de snapshot ne constitue pas une explication.
- Conserver zéro avertissement. À la clôture, le vérifier par une compilation complète sans sorties préexistantes, avec les caches de paquets disponibles.
- Adapter les tests existants si une interface interne change. Ajouter un contrôle ciblé seulement pour une incertitude concrète ; le jalon altbak convenu suffit à la validation transversale de cette vague.
- Consigner quelques lignes de résultat sous le lot terminé. Une fois les six cases cochées, la vague est close ; toute nouvelle optimisation ou correction devient un chantier distinct.

## Lots

- [x] **1 — Séparer les métadonnées immuables de l'état de génération.**

  **Constat :** `CodegenState` contient encore les métadonnées en plus des déclarations, du compteur et des demandes Rebox. `AdtExprs` compte 35 lectures de référence, souvent pour relire les mêmes tables de types ; `CodeGen` en compte 12.

  **Travail :** passer les métadonnées directement dans le contexte de traduction et les helpers de types/conversions. Garder dans l'état mutable les données effectivement produites pendant la traduction. Raccorder en un lot les émetteurs et leurs appelants ; conserver l'ordre des allocations de noms et de l'enregistrement des helpers.

  **Terminé lorsque :** les décisions de typage ne nécessitent plus de lire une référence mutable, les responsabilités restantes de l'état sont explicites et le Go est inchangé.

  **Résultat (14 septembre 2026) :** `ExprContext` et les helpers reçoivent directement `CodegenMetadata`. L'état mutable ne contient plus que `rawDecls`, `globalId` et `reboxPairs` ; le champ `decls`, toujours vide, est retiré. Les lectures `Ref.read` passent de 97 à 4, réservées à ces sorties. Compilation sans avertissement, 11 tests de conversion adaptés réussis, `bin/go/run -c` réussi : mêmes 300 entrées TAST, 387 fichiers Go identiques octet pour octet et 14 résultats fonctionnels inchangés. Architecture mise à jour ; PBO et altbak inchangés. Preuves locales : `/private/tmp/gopurs-wave2-state-t8wc3xll/verification.json`.

- [ ] **2 — Donner un chemin commun à la préparation des constructeurs.**

  **Constat :** `AdtExprs` répète entre définition, construction saturée et accès aux champs la résolution du type ADT, le calcul des arguments génériques et le choix des noms. `TypeStructPointer` transporte quatre informations positionnelles, dont trois chaînes.

  **Travail :** extraire une préparation commune avec des champs nommés pour l'identité PureScript, le constructeur Go, son nom instancié et ses arguments de type. La réutiliser dans les chemins de construction et d'accès ; clarifier la représentation interne des pointeurs et ses consommateurs dans `GoTypes` et `GoConversions`. Préserver les distinctions entre enums, constructeurs éliminés, ADT natifs, pointeurs et dictionnaires de classes.

  **Terminé lorsque :** une même décision de représentation n'est plus reconstruite dans plusieurs branches et que noms, tags, champs, arités et conversions générés restent identiques.

- [ ] **3 — Simplifier les appels et sortir les intrinsics de tableaux.**

  **Constat :** les 697 lignes de `CallExprs` mêlent reconnaissance des intrinsics, traduction des arguments, choix des appels directs et émission des boucles. Les chemins curryfiés et non curryfiés répètent une partie de ces traitements.

  **Travail :** donner aux intrinsics reconnus une représentation explicite et leur propre émetteur pour `map`, `filter` et `foldl`. Factoriser la traduction ordonnée des arguments et les adaptations communes. Garder dans `CallExprs` la sélection entre appels directs, indirects, partiels, surapplications et sauts TCO. Conserver les gardes propres à chaque chemin et le marqueur de filtre frais utilisé par ArrayRoundtrip.

  **Terminé lorsque :** le parcours d'un appel se lit par étapes nommées, les mêmes règles ne sont plus copiées entre les variantes et aucun argument n'est traduit ou évalué deux fois.

- [ ] **4 — Unifier les déclarations Go et le calcul des imports.**

  **Constat :** structs, workers nommés et helpers Rebox sont encore des `rawDecls`. `CodeGen` et `Printer` calculent tous deux des imports à partir du texte rendu. Dans `CodeGen`, `usedPkgNames` parcourt la queue de `parts = [declsStr]`, donc une liste vide.

  **Travail :** représenter les déclarations natives répétées — structs, fonctions nommées et initialisations — avec des nœuds adaptés de `GoAst`, puis migrer ensemble leurs producteurs. Donner au calcul des imports un propriétaire unique, retirer le chemin vide et les rendus intermédiaires devenus inutiles. Déclarer les imports nécessaires aux fragments opaques conservés. Déplacer la préparation du currying encore faite par `Printer` pour `GoFunc` vers les émetteurs, en réutilisant les formes natives existantes.

  **Terminé lorsque :** le printer rend des décisions déjà prises, les déclarations et leurs imports ont un parcours unique, et les déclarations brutes restantes sont identifiées avec leurs dépendances. La frontière visée reste celle de [docs/go-ast-printer.md](docs/go-ast-printer.md), avec le même Go produit.

- [ ] **5 — Clarifier les conversions et l'émission des bridges FFI.**

  **Constat :** `GoConversions` construit encore de nombreuses fonctions Go par concaténation. `FfiBridge` combine analyse des signatures, adaptations d'arguments, callbacks et retours ; `Main.emitModule` assemble lui-même le fichier FFI final.

  **Travail :** utiliser les nœuds de fonctions Go existants pour les enveloppes répétées de conversion. Dans la FFI, séparer l'analyse de la signature et des annotations TAST de l'émission des arguments, callbacks et retours ; mutualiser les prédicats et enveloppes identiques. Regrouper l'assemblage du fichier FFI derrière sa frontière dédiée afin que `Main` coordonne son émission. Préserver les conventions propres à la FFI, aux effets, aux newtypes et aux valeurs opaques.

  **Terminé lorsque :** les conversions répétées sont exprimées une seule fois par famille, les choix de types sont lisibles avant le rendu et les bridges générés gardent leurs signatures, conversions et comportements d'erreur.

- [ ] **6 — Clore la vague et mettre la carte d'architecture à jour.**

  Réactualiser les propriétaires dans la documentation, vérifier les liens et interfaces après les cinq lots, puis enregistrer le dernier contrôle altbak et la compilation complète sans avertissements. Conserver un bilan court des changements et des éventuels écarts expliqués. Les anciens contrôles TCO/fusion et l'extension de couverture ArrayRoundtrip restent des travaux distincts ; ils ne rouvrent pas automatiquement cette vague.
