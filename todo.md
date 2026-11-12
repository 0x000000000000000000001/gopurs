# FFI Go idiomatique — plan de travail

Plan établi le 12 novembre 2026. Les cases ci-dessous décrivent le travail à
réaliser ; ce document ne constitue pas une validation des capacités proposées.

## Objectif final

Offrir une DX exemplaire pour écrire une FFI PureScript en Go : du code naturel,
lisible par un développeur Go, utilisant directement les types et les conventions
des bibliothèques appelées.

Dans les FFI applicatives courantes, viser **zéro import de
`gopurs/output/gopurs_runtime` et zéro manipulation de sa représentation interne**.
La présence de `Value`, `Box`, `Apply`, `FuncN`, de tags ou d'annuleurs internes
doit être prise en charge par les adaptateurs générés. Une intégration spécialisée
peut conserver un accès explicite au runtime, avec une justification documentée.

L'objectif porte également sur les conventions : remplacer ces appels par des
helpers propriétaires, des annotations répétitives ou une nouvelle configuration
à maintenir ne suffirait pas. La FFI doit demander le moins possible au
développeur de connaître les mécanismes de gopurs.

## Principes de conception

- Le développeur écrit des fonctions Go ordinaires, avec leurs imports, types,
  pointeurs, callbacks, résultats et erreurs.
- Les cas courants ne demandent ni registre de bindings, ni DSL, ni fichier de
  correspondance supplémentaire, ni option de compilation par fonction.
- Le type du `foreign import` PureScript et la signature Go fournissent ensemble
  le contrat. La présence d'un retour `error` ne suffit pas à choisir une
  sémantique d'effet ou d'asynchronisme.
- Le TAST conserve les types PureScript, leurs instanciations, les déclarations
  de données et de classes. Exploiter ces informations avant de recourir à une
  représentation dynamique ou à la réflexion.
- Les transformations AST et les adaptateurs s'appliquent aux copies destinées
  à la compilation. Les fichiers Go manuscrits restent les sources de référence,
  utilisables avec les outils Go habituels.
- Le noyau du parseur Go est partagé par WASM et les hôtes natifs. Une règle de
  transformation doit avoir une seule implémentation commune.
- Réutiliser le parseur, le modèle de types et le générateur existants. Chaque
  extension doit répondre à un cas concret ; éviter un second framework FFI.
- Préserver les FFI existantes pendant la migration et les interfaces PureScript
  et JavaScript des applications pilotes.
- Préserver le sens des valeurs : pas de conversion implicite supplémentaire en
  JSON, de changement silencieux de représentation numérique ou de copie
  profonde systématique.
- Limiter les allocations, copies, closures, goroutines et passages par
  `any` à ceux nécessaires au contrat. Une signature idiomatique ne garantit pas
  à elle seule un appel natif optimisé : vérifier séparément les deux aspects.

## Point de départ constaté

- [Le parseur Go](tools/ffi-gen/parser.go) analyse les déclarations et réécrit
  déjà leurs noms et références pour l'intégration dans le code généré.
- [Son modèle JSON](tools/ffi-gen/types.go) et
  [le modèle PureScript](src/Gopurs/FfiTypes.purs) ne représentent actuellement
  qu'un seul résultat par fonction. Les retours Go multiples perdent de
  l'information avant la génération.
- [FfiBridge](src/Gopurs/FfiBridge.purs) et ses modules
  [Render](src/Gopurs/FfiBridge/Render.purs),
  [Values](src/Gopurs/FfiBridge/Values.purs) et
  [TypeSupport](src/Gopurs/FfiBridge/TypeSupport.purs) génèrent déjà des
  adaptations pour des scalaires, tableaux, callbacks et handles.
- L'adaptation générale d'une fonction Go retournant `(T, error)` vers
  `EffectFnAff a`, avec gestion du contexte et de l'annulation, reste à construire.
- Dans b8x, RabbitMQ, PostgreSQL, Paseto, Unzip et le hash xxhash64 exposent
  particulièrement cette mécanique. Les FFI simples donnent déjà des exemples
  de signatures Go naturelles.

Le parseur fournit les faits Go ; le TAST fournit les faits PureScript ; le
générateur vérifie leur compatibilité et produit la liaison. Les transformations
syntaxiques seules ne prouvent pas la compatibilité des types.

## Lot 0 — Définir le contrat et les preuves minimales

- [ ] Inventorier les signatures réellement utilisées dans les FFI manuscrites
  de b8x et les ports concernés : effets, retours multiples, callbacks,
  handles, valeurs dynamiques, records et ADT.
- [ ] Séparer les conversions propres à l'application des conversions imposées
  par la représentation du backend.
- [ ] Consigner une baseline : imports du runtime, wrappers manuscrits, casts,
  helpers spécifiques, instructions de configuration et résultats observables.
- [ ] Choisir une fixture sans réseau ni SDK externe : une fonction Go qui
  retourne `(string, error)`, associée à un import `EffectFnAff String`.
- [ ] Définir ses cas succès, erreur, exécution différée et exécutions répétées.
  Prévoir une reproduction ciblée réalisable en moins de cinq minutes.
- [ ] Définir le contrat cible ci-dessous, en indiquant pour chaque forme son
  statut : existante, à étendre ou non prise en charge.

| Contrat PureScript | Forme Go visée | Adaptation attendue |
| --- | --- | --- |
| Fonction pure | Paramètres et résultat Go ordinaires | Conversion des valeurs |
| `Effect a` / `Effect Unit` | Fonction Go ordinaire retournant une valeur / aucun résultat | Exécution au déclenchement de l'effet |
| `EffectFnN` | Fonction ou callback Go typé | Adaptation de l'appel non curryfié |
| `EffectFnAff a` | Résultat `T` ou `(T, error)` | Exécution différée, succès et échec |
| `EffectFnAff Unit` | Aucun résultat ou un `error` | Succès sans valeur et échec |
| Opération asynchrone annulable | `context.Context` et résultats ordinaires | Propagation de l'annulation selon le contrat de l'opération |
| Handle opaque PureScript | Pointeur ou type natif du SDK | Transport du handle sans déballage manuel |
| `Aff a` directement importé | Formes Go ordinaires équivalentes | Extension à qualifier après le premier pont `EffectFnAff` |

**Critère de sortie :** une première signature précise, des comportements
observables à vérifier et une liste de limites explicites. Aucune migration
générale avant cette preuve.

## Lot 1 — Conserver les signatures Go complètes

- [ ] Représenter tous les résultats dans le modèle partagé, y compris pour
  les callbacks, les résultats nommés et les déclarations groupées.
- [ ] Conserver les informations nécessaires sur pointeurs, types nommés,
  alias, imports, interfaces et fonctions. Distinguer notamment slices,
  tableaux de taille fixe et paramètres variadiques.
- [ ] Identifier `error` et `context.Context` par leur identité de type, en
  tenant compte des alias et des noms masqués ; éviter une simple comparaison
  avec le nom d'un paramètre ou un alias d'import.
- [ ] Utiliser les mécanismes existants de résolution des types lorsque
  l'analyse syntaxique ne suffit pas. Définir les formes initialement acceptées.
- [ ] Mettre à jour le décodage PureScript et les consommateurs du contrat.
  Une forme non prise en charge ne doit plus être réduite silencieusement.
- [ ] Garder les réécritures AST cohérentes avec les scopes, imports et
  références ; préserver commentaires, littéraux et provenance des diagnostics.
- [ ] Versionner le contrat et invalider les caches concernés. Mettre à jour
  l'identité des helpers, les listes de sources embarquées et le paquet distribué.

**Validation :** tests du parseur sur signatures complètes et diagnostics ;
mêmes faits de types via le parseur WASM et natif ; non-régression des FFI
existantes. Pas de migration applicative dans ce lot.

## Lot 2 — Générer le premier pont idiomatique

- [ ] Croiser la signature Go avec le type TAST instancié du foreign import.
  Reconnaître explicitement `EffectFnAff a` avant son éventuel dépliage ;
  vérifier la substitution des paramètres de ses newtypes et callbacks.
- [ ] Générer le pont `(T, error)` vers `EffectFnAff a`, puis le cas
  `error` vers `EffectFnAff Unit`.
- [ ] Prendre en charge les opérations sans erreur Go selon leur type
  PureScript, sans imposer un retour `error` artificiel au développeur.
- [ ] Conserver le démarrage différé : construire une action ou appliquer ses
  arguments ne doit pas exécuter prématurément l'opération.
- [ ] Une nouvelle exécution de l'action doit lancer une nouvelle opération.
  Ne pas introduire de mémorisation involontaire.
- [ ] Définir la transmission d'`error`, de `nil` et des valeurs nulles typées.
  Traiter explicitement la politique des panics, sans les transformer
  silencieusement en erreurs métier.
- [ ] Conserver la compatibilité avec les signatures utilisant déjà le runtime.
  Toute ambiguïté entre ancien et nouveau contrat doit être diagnostiquée.

**Validation :** la fixture Go ne contient aucune référence au runtime et peut
être compilée/testée dans un module Go ordinaire. Son import PureScript passe
les cas succès, erreur, démarrage différé et répétition avec les hôtes JS et
natif. Inspecter aussi le Go généré pour vérifier le chemin réellement utilisé.

## Lot 3 — Fixer le cycle de vie asynchrone

- [ ] Faire revenir l'installation de l'opération avec son annuleur sans
  attendre une opération bloquante. Sinon l'annulation ne devient observable
  qu'après la fin du travail.
- [ ] Définir un arbitrage unique entre résultat, erreur et annulation :
  une seule terminaison observable, y compris en présence de courses.
- [ ] Relier un `context.Context` Go au cycle de vie de l'action PureScript,
  avec propagation et libération du contexte au bon moment.
- [ ] Distinguer une opération effectivement interruptible d'une opération
  qui continue après l'arrêt de l'attente. Ne pas annoncer une annulation
  physique que le SDK ne fournit pas.
- [ ] Fixer la propriété des ressources lors d'une acquisition concurrente
  avec une annulation : qui reçoit ou libère un résultat tardif, et quand.
  S'appuyer sur un contrat explicite d'acquisition/libération ; ne pas déduire
  automatiquement un destructeur d'une méthode nommée `Close`.
- [ ] Couvrir callbacks et abonnements de longue durée : durée de validité,
  désinscription, appels après arrêt et synchronisation avec la fermeture.
- [ ] Privilégier les mécanismes Go usuels, dont `context`, et les abstractions
  PureScript existantes. Toute convention supplémentaire doit être nécessaire
  et expliquée par un cas concret.

**Validation :** fixtures déterministes sans services externes pour annulation
avant démarrage, pendant l'opération, course avec succès/erreur, résultat tardif,
libération unique et callback après arrêt. Vérifier les goroutines restantes
et les courses avec les outils Go adaptés avant de migrer les acquisitions
RabbitMQ ou PostgreSQL.

## Lot 4 — Rendre les paramètres et résultats naturels

- [ ] Adapter les callbacks typés afin que le code Go puisse appeler une
  fonction ordinaire sans `Apply` ni fabrication manuelle d'un effet.
- [ ] Transporter les handles par leurs types Go natifs. Vérifier aller-retour,
  identité, `nil`, durée de validité et usages après fermeture selon le contrat
  du SDK.
- [ ] Définir le passage de `Foreign`, `Array Foreign` et des données
  dynamiques vers les valeurs Go correspondantes, notamment pour PostgreSQL.
  Un `any` naturel ne doit pas masquer une obligation de déballer un `Value`.
- [ ] Préserver les distinctions pertinentes : valeur absente, `nil`, tableau
  vide, zéro, faux, taille et précision numériques.
- [ ] Garder les conversions SQL, JSON et les choix de sérialisation sous
  le contrôle de l'application.
- [ ] Compléter les conventions `Effect` et `EffectFnN` pour éviter les thunks
  et arguments `Unit` artificiels dans le Go manuscrit lorsque le type
  PureScript permet de les générer sans ambiguïté.
- [ ] Qualifier ensuite l'import direct d'`Aff a` : réutiliser le contrat
  asynchrone commun sans imposer une migration des imports `EffectFnAff`
  existants ni exposer une nouvelle famille de helpers au développeur.

**Critère de sortie :** des signatures natives couvrent les besoins des premiers
pilotes, avec erreurs précises sur les formes incompatibles. Le typage et le
comportement priment sur une acceptation permissive de toutes les signatures.

## Lot 5 — Migrer les pilotes b8x

### RabbitMQ

- [ ] Migrer une opération simple, puis les acquisitions, fermetures et
  abonnements lorsque leurs contrats de cycle de vie sont validés.
- [ ] Remplacer `rabbitAff` et les manipulations du runtime par les fonctions,
  types et retours ordinaires du SDK.
- [ ] Conserver les interfaces PureScript/JavaScript et vérifier les callbacks
  de messages, erreurs de connexion et fermetures.
- [ ] Vérifier les scénarios annulables et non annulables réellement proposés
  par le SDK ; documenter les limites restantes.

### PostgreSQL

- [ ] Migrer d'abord une requête simple, puis les paramètres/résultats
  dynamiques, le comptage, les connexions dédiées et les transactions.
- [ ] Retirer des FFI les inspections de tags, `Box`, `Apply` et les essais
  de déballage manuel des handles.
- [ ] Utiliser le contexte transmis pour les opérations qui le prennent
  en charge, avec fermeture des rows et libération des connexions.
- [ ] Préserver exactement les conversions applicatives et les comportements
  des erreurs, résultats vides, valeurs nulles et transactions.

**Validation de chaque pilote :** tests Go directs avec doublures lorsque
possible, puis tests d'intégration représentatifs. Comparer les comportements
PureScript avant/après et les interfaces partagées avec JavaScript. La seule
compilation ne prouve pas le bon fonctionnement de la FFI.

**Critère DX :** comprendre les opérations Go courantes en connaissant Go et
le SDK concerné, sans apprendre la représentation du runtime PureScript.

## Lot 6 — Étendre aux données structurées et aux autres FFI

- [ ] Ajouter les correspondances records fermés vers structs Go à partir du
  TAST et des déclarations Go : champs, types, noms, présence et valeurs nulles.
  Définir les divergences acceptées et les diagnostics des autres cas.
- [ ] Exploiter les déclarations TAST pour les newtypes et ADT. Garder les tags,
  layouts et détails de boxing dans les adaptateurs générés.
- [ ] Traiter les rangées ouvertes, variantes polymorphes, conteneurs imbriqués
  et contraintes uniquement avec des règles explicites et des preuves ciblées.
- [ ] Si une information ou une garantie manque, identifier précisément le
  besoin avant de proposer une évolution du TAST/PBO ; ne pas supposer que les
  types ont été effacés.
- [ ] Migrer ensuite les FFI b8x restantes, notamment Paseto, Unzip, Hash et
  Router, en conservant les politiques métier et les comportements des cibles.
- [ ] Recenser chaque dépendance résiduelle au runtime dans une FFI manuscrite,
  sa raison et le travail nécessaire pour la retirer.

**Critère de sortie :** les FFI applicatives ordinaires n'exposent plus le
runtime. Les exceptions sont limitées, identifiées et justifiées techniquement.

## Lot 7 — Vérifier la DX de bout en bout et documenter

- [ ] Documenter quelques exemples Go naturels : fonction pure, effet, opération
  avec erreur, callback, handle, contexte et données structurées.
- [ ] Présenter une seule voie recommandée pour chaque besoin courant.
  Décrire les conventions indispensables et les limites sans demander au
  lecteur de connaître l'ABI du compilateur.
- [ ] Situer les diagnostics dans le module PureScript et le fichier Go
  manuscrit : signature rencontrée, forme attendue et correction possible.
- [ ] Refaire le parcours depuis un projet neuf et une migration depuis le
  compilateur/JavaScript officiels, en installant le paquet assemblé.
- [ ] Vérifier les sorties avec cache froid, réutilisation du cache, changement
  de signature FFI et récupération après erreur.
- [ ] Vérifier analyse et génération avec les hôtes JS/WASM et natif.
  Qualifier séparément l'exécution des bibliothèques sur les cibles supportées :
  le parseur WASM ne rend pas un SDK réseau compatible avec une cible WASM.
- [ ] Vérifier que le paquet distribue le parseur et les helpers correspondant
  au nouveau contrat, sans dépendance implicite au checkout du développeur.
- [ ] Comparer la FFI avant/après : code de liaison manuscrit, imports, casts,
  configuration et connaissances spécifiques requises. Vérifier aussi les
  allocations et conversions introduites lorsque le chemin le justifie.
- [ ] Pour toute conclusion de performance applicative, utiliser les baselines
  officielles du README d'altbak.pub. Les microbenchmarks ciblés complètent
  ces baselines et ne les remplacent pas.

## Définition de terminé

- [ ] Les FFI courantes s'écrivent et se testent comme du Go ordinaire.
- [ ] Les pilotes RabbitMQ et PostgreSQL, puis les autres FFI ordinaires de b8x,
  n'importent plus le runtime et ne manipulent plus sa représentation interne.
- [ ] Aucun helper propriétaire ne remplace artificiellement le boilerplate
  supprimé ; les conventions et fichiers à maintenir restent minimaux.
- [ ] Les effets, erreurs, callbacks, annulations et ressources ont un contrat
  observable couvert par des tests.
- [ ] Les conversions préservent les valeurs et les comportements applicatifs.
- [ ] Les FFI existantes restent utilisables pendant la migration.
- [ ] Les signatures non prises en charge échouent avec un diagnostic utile.
- [ ] Les exemples de documentation et les parcours d'installation/migration
  ont été exécutés avec les versions exactes consignées.
- [ ] Les vérifications distinguent les hôtes du compilateur, les cibles
  d'exécution et les bibliothèques réellement testées.
- [ ] Les limites restantes sont explicites et ne sont pas présentées comme
  des cas validés.

## Première étape à entreprendre

Commencer par le lot 0, puis préserver les retours multiples dans le parseur
pour la fixture `(string, error) → EffectFnAff String`. Obtenir une preuve
courte de génération et d'exécution avant d'élargir le contrat ou de migrer
b8x. À chaque étape, conserver la commande de reproduction et le résultat
observé ; un test échoué doit orienter la prochaine petite modification.
