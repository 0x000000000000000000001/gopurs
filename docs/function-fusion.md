# Producteurs de fonctions comptées

`Gopurs.FunctionFusion` s’exécute après PBO et la fusion des thunks, avant
l’analyse de possession et la TCO. Elle remplace une chaîne récursive de
fonctions par une closure capturant son compteur et appelant un worker bouclé.

## Forme reconnue

Le producteur appartient à un groupe récursif singleton, avec le type explicite
`Int -> (Int -> Int) -> Int -> Int`, et suit exactement cette structure :

```purescript
build 0 = identity
build n =
  let previous = build (n - 1)
  in \f x -> f (previous f x)
```

L’identité est vérifiée dans son corps, directement ou via un nom du même module.
Le regroupement des flèches et une identité polymorphe sont admis. Les annotations
incompatibles, les décréments différents, les groupes mutuellement récursifs,
les bases opaques et le travail supplémentaire font conserver le code initial.
Les noms du programme ou du benchmark n’interviennent pas dans la reconnaissance.

## Code produit et garanties

Pour `n >= 0`, le producteur retourne une fonction `\f x -> worker n f x`.
Le worker applique `f` exactement `n` fois à son accumulateur ; la TCO existante
le transforme en boucle. Le compteur capturé reste immuable : chaque invocation
possède son propre compteur de travail. Fonctions sauvegardées, réutilisations
et applications partielles conservent ainsi leur comportement.

La branche négative conserve le corps initial, évalué au moment de la production.
Le garde placé avant la fonction retournée préserve l’arité publique du
producteur. Pour zéro, aucun callback n’est invoqué. Les noms des workers sont
réservés contre les bindings et FFI existants, y compris après normalisation Go.

Les paramètres compteur/accumulateur deviennent `int64`, tandis que le callback
conserve son ABI `Value` et ses appels `Apply`. La passe n’ajoute ni formule
arithmétique pour remplacer une composition ni représentation native générale
des fonctions. Elle utilise les types et les corps du TAST transformé ; les
annotations d’usage source ne constituent pas une preuve nécessaire à ce motif.

## Validation du 17 septembre 2026

Le build et le bundle passent sans erreur ni avertissement. Les 68 tests outils
existants et 29 nouveaux passent. Deux tests exécutent le Go généré et vérifient
les callbacks, les applications partielles, les réutilisations et une exception
au troisième callback suivie d’une nouvelle invocation. Les autres contrôlent
les refus, les annotations, le chemin négatif et les collisions de noms.

La nouvelle fixture TAST `CountedFunctions` s’exécute avant et après changement.
Elle couvre également un producteur qui capture le compteur, donc refusé par la
passe. Son nouveau snapshot et celui de `ThunkFusion` ont ensuite été vérifiés
sans mise à jour, avec compilation et exécution Go réussies. La totalité des
autres fixtures de compilation n’a pas été relancée.

Dans la [campagne archivée](/Users/0x1/Documents/htdocs/altbak.pub/scratch/gopurs-counted-functions-20260917/REPORT.md),
301 CoreFn identiques produisent un changement limité à `Test_Church.go`.
La sonde indépendante passe de 490,800 à 236,358 µs, avec cinq paires favorables,
157 à 112 allocations et 6 800 à 5 280 octets par calcul. Le harnais complet
confirme Church à 486,04 → 233,92 µs, puis 480,54 → 236,29 µs lors d’une seconde
série. La baisse du total n’est pas établie : les autres lignes fluctuent
suffisamment pour masquer ce gain. La baseline officielle du README reste
distincte de ces campagnes et n’a pas été réécrite.
