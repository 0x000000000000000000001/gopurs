# Rendre les FFI Go de b8x idiomatiques

L’audit du 6 novembre 2026 a porté sur les 21 fichiers FFI Go de `b8x/src`. Les quatre premiers cas ont été reproduits en exécutant isolément les wrappers Go générés de b8x. Les autres entrées décrivent des dépendances visibles dans les sources.

## Défauts reproduits

- [x] **[Cookie.go](../../b8x/src/Util/Cookie/Cookie.go#L28) — transmettre les six options séparément.** La réflexion sur une `gopurs_runtime.Value` faisait disparaître les attributs du cookie. PureScript déstructure désormais `SerializeOptions` avant l’appel aux FFI Go, JS et PHP. Le Go compilé produit les mêmes en-têtes que la bibliothèque JS pour la session, la déconnexion (`Max-Age=0`) et les options désactivées. [Frontière PureScript](../../b8x/src/Util/Cookie/Cookie.purs#L20) · [Tests de session et déconnexion](../../b8x/test/ffi/Native.purs#L25)

- [ ] **[Paseto.go](../../b8x/src/Util/Crypto/Paseto/Paseto.go#L18) — le chiffrement perd le contenu du jeton.** Après chiffrement puis déchiffrement du payload de connexion, seul `exp` est présent. L’assertion vers `map[string]interface{}` échoue sur la valeur enveloppée et le chiffrement continue sans `userId` ni `csrfToken`.
  Fournir à PASETO un objet de claims natif via la frontière `Foreign`, puis transmettre le résultat Go ou l’erreur via l’adaptateur `Aff` commun.
  Vérifier l’aller-retour des deux champs et l’échec explicite pour un payload incompatible. [Appel de connexion](../../b8x/src/Inter/Api/Membership/Auth.purs#L151)

- [ ] **[RabbitMq.go](../../b8x/src/Infra/Client/RabbitMq/RabbitMq.go#L43) — la connexion opaque provoque un panic.** À la création du canal, le type dynamique reçu est `gopurs_runtime.Value` et l’assertion exige `*amqp.Connection`. Les opérations suivantes supposent de même un `*amqp.Channel`.
  Essayer des paramètres Go typés et faire préserver l’identité des pointeurs par le bridge, à l’aller comme au retour. Adapter aussi `onMessage` en `func(string)` pour que la boucle de consommation appelle directement le callback.
  Vérifier la séquence connexion → canal → livraison d’un message → fermeture.

- [ ] **[TaggedSum.go](../../b8x/src/Util/Json/TaggedSum.go#L3) — la normalisation JSON ne s’applique pas.** En passant `{"type":"Given"}` par le wrapper, `value` reste absent. La réflexion examine l’enveloppe. Sur une map Go reçue directement, la fonction modifie de plus l’objet fourni.
  Envisager la normalisation en PureScript avec les opérations de lecture et d’insertion sur les objets étrangers.
  Vérifier que le nouveau résultat décode `Given Nothing`, que l’objet d’origine garde ses champs et qu’un `value` existant est préservé. [Contrat du décodeur](../../b8x/src/Util/Json/TaggedSum.purs#L10)

## Dépendances au runtime dans le code métier

- [ ] **[Postgres.go](../../b8x/src/Infra/Client/Postgres/Postgres.go#L49) — le client SQL interprète la représentation de `Value`.** `_Query` et `_QueryCount` dupliquent la lecture des tags, de `IntVal` et des chaînes. Elles appellent directement le symbole généré `Foreign_UnboxForJSON`. Cinq opérations déballent le handle et `_Query` réemballe chaque ligne puis le tableau.
  Séparer la conversion `Foreign` de la politique SQL. Fournir au client des paramètres Go, garder dans b8x les choix d’encodage attendus par pgx, et adapter les handles typés à la frontière.
  Vérifier que scalaires, NULL, objets JSON et tableaux conservent leur comportement dans les deux chemins. [Deuxième conversion](../../b8x/src/Infra/Client/Postgres/Postgres.go#L129)

- [ ] **[Router.go](../../b8x/src/Util/Run/Router.go#L18) — le routeur connaît la représentation de `VariantF`.** `Build` appelle `ExtractVariant`, imprime le variant entier en cas d’échec, puis essaie pointeur, map, struct ou chaîne et répète la recherche du handler.
  Essayer d’extraire l’étiquette et la charge côté PureScript avec `unvariantF`. Le Go pourrait alors se limiter à la table de handlers.
  Vérifier qu’une étiquette connue transmet sa charge et qu’une inconnue remet le variant original au fallback. [Éliminateur existant](../gopurs-variant/src/Data/Functor/Variant.purs#L472)

- [ ] **[Hash.go](../../b8x/src/Util/Crypto/Hash.go#L29) — un hash synchrone reconstruit tout le protocole de `EffectFnAff`.** Dans `_Xxhash64`, une ligne calcule le résultat. `Func2`, `Box`, `Apply` et `Func3` servent à le livrer et à acquitter l’annulation.
  Garder le calcul Go sous forme de fonction `string → string` et l’exposer via l’adaptateur `Aff` commun. Préserver l’API PureScript `String -> Aff String`, également utilisée par le backend JS qui charge le module de hash de façon asynchrone.

- [ ] **[Unzip.go](../../b8x/src/Util/File/Unzip.go#L11) — chaque sortie ZIP manipule des callbacks boxés.** L’ouverture, la recherche du fichier et la lecture appellent chacune `Apply(onError, Box(err))`.
  L’extraction Go pourrait retourner `(string, error)` et laisser l’adaptateur `Aff` commun publier le résultat.
  Vérifier aussi le passage de l’`ArrayBuffer` PureScript au paramètre Go `zipContent string`.
  Essayer une entrée `[]byte` sur un vrai ZIP binaire, puis vérifier les erreurs pour une archive illisible et un fichier absent. `_UnzipToDirectory` doit continuer à signaler explicitement son état non implémenté.

## Conventions d’appel visibles dans de petites FFI

- [ ] **[Signal.go](../../b8x/src/Util/Signal.go#L9)** — un handler `Effect Unit` devient `func(interface{}) interface{}` et s’appelle avec `fn(nil)`.
  Faire adapter le handler en `func()` et vérifier que l’effet s’exécute à la réception du signal.
- [ ] **[Timer.go](../../b8x/src/Util/Timer.go#L12)** — chaque tick appelle aussi `fn(nil)`. La création retourne `interface{}` alors que l’arrêt attend un `*IntervalId`.
  Réutiliser l’adaptation vers `func()` et faire transiter le handle opaque sous son type Go.
  Vérifier qu’on peut arrêter le ticker avec le handle obtenu à sa création.
- [ ] **[EventStore.go](../../b8x/src/Infra/EventStore/Postgres/EventStore.go#L21)** — `GenerateLockKeyStr(_ interface{}) interface{}` produit toujours une chaîne.
  Essayer une fonction Go sans argument retournant `string`, en laissant le bridge différer l’appel jusqu’à l’exécution de l’`Effect String`.
  Deux exécutions doivent avancer le compteur deux fois.
- [ ] **[Ulid.go](../../b8x/src/Util/Type/Ulid.go#L7)** — `_GenerateUlid` reçoit un argument factice et efface le type de sa chaîne de retour.
  Essayer la même adaptation de `Effect String` que pour l’EventStore et vérifier qu’un identifiant est généré à chaque exécution.
- [ ] **[Uuid.go](../../b8x/src/Util/Type/Uuid.go#L3)** — `_GenerateV7Uuid` renvoie une chaîne via `interface{}` pour un import déclaré `Effect UUID`.
  Envisager un import privé `Effect String` et construire le `Uuid` applicatif côté PureScript.
  Préserver l’API publique et vérifier la version 7 après passage par PureScript.
- [ ] **[KeepAlive.go](../../b8x/src/Inter/Cli/_/Util/KeepAlive.go#L1)** — `KeepMainAlive` déclare un résultat `interface{}` et un `return nil` après un `select {}` qui bloque indéfiniment.
  Essayer une fonction Go sans résultat pour cet `Effect Unit` et vérifier que le blocage commence uniquement lorsque l’effet est exécuté.

## Où placer les adaptations

1. **Dans le bridge de gopurs**, exploiter le TAST et la signature Go pour les records fermés, les handles opaques, les callbacks et les effets. Procéder par petites fixtures passant par un véritable `foreign import`. Préserver les appelants existants qui utilisent `any` pour recevoir une valeur enveloppée. [Conversion actuelle des arguments](src/Gopurs/FfiBridge/Values.purs#L40)
2. **Dans `gopurs-foreign`**, réutiliser ou compléter les conversions de données réellement dynamiques. Les exposer en PureScript pour pouvoir les fournir à la FFI au besoin. Le traitement des cookies, des claims ou du SQL reste dans b8x. [Conversion JSON existante](../gopurs-foreign/src/Foreign.purs#L192)
3. **Pour `EffectFnAff`**, essayer un seul adaptateur partagé, à partir du cas simple de `_Xxhash64`, puis du couple succès/erreur de l’extraction ZIP. Étendre l’essai à PostgreSQL, PASETO et RabbitMQ en préservant pour chaque opération le mode d’exécution et le contrat d’annulation. [Adaptateur local de RabbitMQ](../../b8x/src/Infra/Client/RabbitMq/RabbitMq.go#L13)
