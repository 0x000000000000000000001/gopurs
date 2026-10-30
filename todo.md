# Plan de publication officielle de gopurs

Ce plan suit la préparation de `0.1.0`, sa qualification puis sa publication.
Une case de préparation cochée ne vaut pas validation de la release : les
résultats doivent porter sur les sources et artefacts exacts qui seront publiés.
Le bilan des campagnes est dans `docs/quality-0.1.0.md`.

État RC7 au 30 octobre : la qualification des trois plateformes et l'audit des
artefacts originaux passent. Le workflow reste en échec à la préparation du
draft ; une reprise opérateur a été explicitement autorisée pour ces artefacts,
sans modifier le tag. Les 45 assets sont publiés sur GitHub en prerelease ;
l'archive npm publique et sa provenance sont vérifiées. npm a aussi attribué
`latest` à RC7 malgré `--tag next` et refusé sa suppression ; cette exception
a été acceptée explicitement. L'installation npm et le quickstart public passent
sur Linux arm64 avec des caches applicatifs neufs. La compilation b8x depuis
les artefacts publics passe aussi : 17 étapes, 2 707 modules TAST et 12 exécutables
liés, sans exécution des programmes. La publication RC7 est clôturée avec ces
exceptions ; les cases ouvertes ci-dessous restent des travaux pour la stable.
L'intégration ultérieure de `bin/modtest` concerne les futures campagnes
complètes ; elle n'ajoute aucune validation à cette qualification RC7.

## 1. Figer le périmètre et le contrat public

- [x] Choisir la version de départ et le schéma de versionnement.
- [x] Lister les surfaces stables : CLI, backend Go, formats intermédiaires, bibliothèques `gopurs-*`.
- [x] Lister explicitement ce qui reste expérimental.
- [x] Fixer les plateformes supportées : OS, architectures, versions de Go et de PureScript.
- [x] Auditer les licences du compilateur, du fork PureScript et de PBO ; consigner séparément les écarts des ports sources dans [l'audit des bibliothèques](docs/library-source-audit.md).
- [x] Ajouter les attributions et notices tierces nécessaires.
- [x] Définir la politique de support et de dépréciation.

## 2. Garantir l'installation depuis zéro

- [x] Tester un clone propre sur une machine sans cache ni binaire préexistant.
- [x] Rendre le build JavaScript indépendant des checkouts voisins.
- [x] Rendre le bootstrap natif et les applications reproductibles depuis des entrées publiques épinglées (fork TAST, PBO, bibliothèques et adaptations Go).
- [x] Épingler les dépendances du build JavaScript.
- [x] Figer et vérifier les révisions du frontend TAST et de tous les ports `gopurs-*` nécessaires à la release.
- [x] Documenter la chaîne de bootstrap : binaire JS, binaire Go précédent, artefacts de référence.
- [x] Publier les checksums des binaires de bootstrap de RC7 avec les artefacts signés.
- [x] Vérifier la reproductibilité des sorties, ou documenter précisément les différences.
- [x] Mettre en place une CI Linux et macOS sur les architectures supportées.
- [ ] Vérifier que la CI part d'un environnement vierge, sans cache préchauffé obligatoire.

## 2 bis. Nixifier le projet

- [x] Fournir des sorties `packages`, `apps`, `checks` et `devShells` sur le modèle de phpurs, sur les trois systèmes annoncés.
- [x] Construire un paquet gopurs autonome depuis les sources et les lockfiles, sans dépendance à un checkout voisin.
- [x] Fournir un frontend TAST épinglé et un parcours documenté pour compiler une application avec les ports Go.
- [x] Vérifier localement sur Linux arm64 `nix build`, `nix run` et `nix flake check`, avec un vrai programme PureScript/FFI compilé puis exécuté en Go ; les validations de la matrice CI sont consignées ci-dessous.
- [x] Vérifier l'évaluation des trois systèmes sans IFD et archiver les preuves locales Nix ; cette évaluation ne remplace pas leurs builds en CI.
- [x] Exécuter les validations Nix sur Linux amd64/arm64 et macOS arm64 en CI : les trois passent dans la campagne initiale du 27 octobre ; la validation du tag reste distincte.
- [x] Documenter les commandes de développement, de build et d'installation Nix et leurs entrées exactes.

## 3. Nettoyer le dépôt

- [x] Retirer les artefacts générés, caches, logs, binaires et sauvegardes du suivi Git.
- [x] Retirer les preuves de benchmark et fichiers d'archive internes des surfaces publiées.
- [x] Rechercher et supprimer les secrets, jetons, URLs privées et chemins personnels.
- [x] Enlever les mentions `WIP` des surfaces officielles ou les convertir en limitations documentées.
- [x] Supprimer le code mort et les options expérimentales visibles.
- [x] Harmoniser le nommage `gopurs` et les intitulés des paquets.
- [x] Vérifier que les fichiers publiés dans les paquets sont limités au nécessaire.
- [x] Éviter tout refactor massif non nécessaire avant la première release.

## 4. Finaliser la documentation

- [x] Écrire un quickstart testé de bout en bout.
- [x] Documenter l'installation, les prérequis et la référence CLI.
- [x] Documenter l'architecture, les formats et le fonctionnement du backend Go.
- [x] Documenter les limitations connues et les différences avec le compilateur officiel.
- [x] Documenter la migration et la compatibilité entre versions.
- [x] Ajouter un guide de dépannage et une FAQ.
- [x] Ajouter `CONTRIBUTING`, `SECURITY`, `CHANGELOG` et le code de conduite.
- [x] Reprendre les benchmarks avec protocole, machine, commandes et scripts reproductibles.
- [x] Vérifier qu'aucune promesse de performance ou de compatibilité n'est ambiguë.

## 5. Passer les portes de qualité

- [x] Exécuter la suite de qualification complète de RC7 dans un environnement propre sur les trois plateformes : 267 portes chacune, avec l'hôte Rust expérimental exclu ; ce résultat ne vaut pas CI globale verte.
- [x] Exécuter les 401 fixtures en Release et les 401 en Dev avec le compilateur natif frais : 802/802 passent, snapshots inchangés et exécution Go validée ; preuves archivées dans le bilan qualité.
- [x] Vérifier la conformité avec le compilateur officiel sur un corpus représentatif.
- [x] Vérifier les golden tests, la parité FFI et les formats générés.
- [x] Tester les erreurs, annulations, bornes de configuration mémoire et cas de concurrence, avec les tests Go `-race` et l'émission native.
- [x] Vérifier les performances à froid et à chaud avec le protocole publié : 51 corpus, 1 224 invocations, comparaison aux baselines officielles et preuves archivées.
- [ ] Faire tourner une release candidate et corriger les régressions bloquantes.
- [ ] Exiger une CI verte sur le tag de release : le run RC7 reste en échec au draft, avec une exception opérateur documentée pour sa publication.
- [ ] Corriger la reprise CI pour distinguer les jobs réellement réexécutés des copies de métadonnées GitHub, et la comparaison de provenance npm pour le seul ajout d'un `keyid` DSSE vide.
- [x] Configurer la CI sur les tags pour couvrir le frontend TAST, le bootstrap natif et l'exécution d'applications, en plus du paquet JavaScript ; son exécution verte reste une porte distincte.

## 6. Préparer les artefacts et la publication

- [x] Produire les binaires RC7 par plateforme avec les flags de release documentés.
- [x] Générer les checksums, la provenance et les SBOM de RC7 ; artefacts originaux audités.
- [x] Signer le tag RC7 et les attestations de ses artefacts ; vérifications cryptographiques passées.
- [x] Rédiger les notes de version et le changelog de RC7 ; la version stable reste à préparer.
- [x] Tester l'installation npm de RC7 depuis les artefacts publics sur Linux arm64, avec vérification des octets et de la provenance.
- [ ] Vérifier les ports sources `gopurs-*` : commits compatibles épinglés, dépendances résolues et notices complètes ; appliquer versions, bornes et allowlists aux paquets effectivement publiés, sans versions communes imposées ([audit et écarts restants](docs/library-source-audit.md)).
- [x] Faire un `npm publish --dry-run` sur l'archive locale préparée.
- [x] Vérifier la provenance de la publication npm réelle de RC7 et l'identité exacte des octets téléchargés.
- [ ] Préparer la procédure de rollback et conserver les artefacts précédents.
- [x] Aligner le contrat de support avec les parcours réellement distribués : hôte JavaScript npm/Nix et binaires Go ; hôte Rust expérimental.
- [ ] Enregistrer dans Git toutes les sources, configurations et notices nécessaires avant de créer la RC ; vérifier un checkout du commit publié.
- [x] Effectuer le smoke test public sur Linux arm64 dans un nouveau projet, avec caches npm, Spago et Go neufs et prérequis de toolchain déjà préparés.

## 7. Clôturer la publication

- [x] Vérifier le parcours du README et de son quickstart lié : installation publique, compilation PureScript/Go et sortie exacte `Hello from gopurs`, avec les prérequis documentés.
- [x] Vérifier la compilation b8x complète depuis les artefacts publics RC7 sur Linux arm64 : 17 étapes et 12 exécutables liés ; aucun programme exécuté.
- [x] Publier les limites et les travaux futurs de RC7 sans les présenter comme terminés.
- [x] Publier les notes de RC7, les plateformes supportées et le canal de support dans la release et la documentation.
- [x] Archiver les preuves RC7, les résultats de la CI et les contrôles publics, avec leurs limites et exceptions.

## Critères de sortie

- [ ] Clone propre, installation et premier build réussis sans intervention manuelle.
- [ ] CI verte sur toutes les plateformes annoncées.
- [ ] Aucun secret, chemin privé ou artefact interne dans la publication.
- [ ] Licences et attributions complètes.
- [ ] Artefacts reproductibles ou procédure de vérification explicite.
- [x] Release candidate vérifiée depuis les artefacts publics : quickstart exécuté et b8x compilé sur Linux arm64 ; exceptions de publication documentées.
- [ ] Rollback testé.
- [ ] Documentation et benchmarks conformes à ce qui est réellement mesuré.
- [ ] Installation et application Nix validées sur la matrice supportée.
