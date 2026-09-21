# Préparation parallèle

`GOPURS_PREPARE_JOBS` règle le nombre de calculs simultanés des contributions transitives de monomorphisation : **2 par défaut**, borné entre **1 et 8**. Une valeur de 1 exécute les travaux séquentiellement. Le réglage est indépendant du chargement TAST et de l’émission Go.

Après la collecte initiale des instanciations, chaque tour dispose d’une table de spécialisations et d’un cache immuables. Un travail correspond à une paire `(nom qualifié, clé de spécialisation)`. Il retourne sa contribution et une éventuelle mise à jour du cache. Le producteur fusionne ensuite les résultats dans l’ordre des deux parcours de Map d’origine, puis termine le tour avant de commencer le suivant. La priorité du premier payload, l’union des callers et le critère d’arrêt sont conservés.

`Gopurs.Preparation.runPreparationJobs` découpe le tableau une seule fois par tour en blocs contigus. Chaque bloc est évalué dans un `defer` exécuté par Aff ; `parTraverse` restitue les blocs dans leur ordre initial. Le backend natif Aff exécute ces calculs dans des goroutines. Le runtime JavaScript conserve les mêmes résultats, sans parallélisme CPU multithread dans ce chemin.

Le premier prototype découpait répétitivement le tableau restant en petits lots. Le profil b8x a montré des copies quadratiques, amplifiées par le bridge non générique de `Array.slice` ; cette organisation a été retirée. Les blocs contigus évitent ces copies répétées, au prix d’un équilibrage dépendant du coût réel des travaux de chaque bloc.

L’API pure PBO `transitiveCollect` reste disponible via le même moteur exécuté séquentiellement. `transitiveCollectWith` reçoit un dispatcher dont le contrat impose de rendre les résultats dans l’ordre des travaux. Les caches appartiennent à un appel et ne sont pas partagés entre builds.

Vérifications ciblées après construction du compilateur :

```sh
node ../../purescript-backend-optimizer-gopurs/test/monomorphize-transitive.mjs output
node ../../purescript-backend-optimizer-gopurs/test/transitive-parallel.mjs output
GOPURS_NATIVE_OUTPUT=/chemin/bootstrap/output node --test tools/preparation-native.test.mjs
```

Le test natif requiert les marqueurs PASS des cas effectivement exécutés sous `-race`. Il contrôle le différé, la réexécution, l’exécution unique, le chevauchement, la borne de concurrence et l’ordre des résultats. Le fichier Go de test est installé temporairement dans le bootstrap généré puis supprimé.
